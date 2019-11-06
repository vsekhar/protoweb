package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/chromedp/cdproto"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

var sitesFilename = flag.String("sitesfile", "", "file containing seed URLs, one per line")
var headersFilename = flag.String("headersfile", "", "CSV file to write headers to: {URL},{headername},{headervalue}")
var jobs = flag.Uint("jobs", 1, "number of simultaneous jobs (Chrome processes) to use to fetch URLs")
var depth = flag.Uint("depth", 1, "depth to traverse (1=seads only, 0=unlimited)")
var timeout = flag.Uint("timeout", 10, "timeout per URL in seconds")
var verbose = flag.Bool("verbose", false, "verbose output")
var progress = flag.Uint("progress", 0, "report progress every n URLs (0=disable)")

type urlEntry struct {
	url   string
	depth uint
}

func eTLDPlusOne(domain string) (string, error) {
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") || strings.Contains(domain, "..") {
		return "", fmt.Errorf("publicsuffix: empty label in domain %q", domain)
	}
	suffix, _ := publicsuffix.PublicSuffix(domain)
	if len(domain) == len(suffix) {
		// handles s3.amazonaws.com and friends which are public suffixes that assign directories rather than subdomains
		return domain, nil
	}
	if len(domain) < len(suffix) {
		return "", fmt.Errorf("publicsuffix: cannot derive eTLD+1 for domain %q", domain)
	}
	i := len(domain) - len(suffix) - 1
	if domain[i] != '.' {
		return "", fmt.Errorf("publicsuffix: invalid public suffix %q for domain %q", suffix, domain)
	}
	return domain[1+strings.LastIndex(domain[:i], "."):], nil
}

type printFunc func(string, ...interface{})

type progressReporter struct {
	every  uint
	print  printFunc
	prefix string

	mu sync.Mutex
	n  uint
}

func newProgressReporter(every uint, print printFunc, prefix string) *progressReporter {
	return &progressReporter{every: every, print: print}
}

func (pr *progressReporter) Done(url string) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	pr.n++
	if pr.every > 0 && pr.n%pr.every == 0 {
		pr.print("...%s: %d @ '%s'", pr.prefix, pr.n, url)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	flag.Parse()
	pr := newProgressReporter(*progress, log.Printf, "")

	urls := make(chan urlEntry, 10000)
	headers := make(chan []string, 1000)
	enqueueHeaders := func(url string, hmap map[string]interface{}) {
		for h, vs := range hmap {
			if strings.HasPrefix(strings.ToLower(h), "x-google") {
				continue
			}
			// TODO: handle multiple headers better (is vs an array?)
			headers <- []string{url, strings.ToLower(h), vs.(string)}
		}
	}
	var wg sync.WaitGroup

	if *sitesFilename == "" {
		log.Fatal("sitesFilename required")
	}
	sitesFile, err := os.Open(*sitesFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer sitesFile.Close()
	scanner := bufio.NewScanner(sitesFile)
	for scanner.Scan() {
		q := urlEntry{url: scanner.Text(), depth: 1}
		if *verbose {
			log.Printf("enqueueing seed %s", q.url)
		}
		urls <- q
	}

	fetcher := func(i int) {
		defer wg.Done()
		if *verbose {
			defer log.Printf("terminating fetcher")
		}

		logFunc := func(string, ...interface{}) {}
		if *verbose {
			logFunc = log.Printf
		}
		browser, closeBrowser := chromedp.NewContext(
			context.Background(),
			chromedp.WithLogf(logFunc),
		)
		defer closeBrowser()
		// start the browser so that tabs (not browsers) are created below
		if err := chromedp.Run(browser); err != nil {
			log.Fatal(err)
		}

		loadURL := func(u urlEntry) {
			if *verbose {
				log.Printf("%d: fetching %s @ depth %d", i, u.url, u.depth)
			}
			tab, closeTab := chromedp.NewContext(browser)
			defer closeTab()

			// Listen and handle events
			loadedChan := make(chan struct{}, 100)
			chromedp.ListenTarget(tab, func(event interface{}) {
				switch x := event.(type) {
				case *network.EventResponseReceived:
					// Enqueue all the headers for output
					if x.Response.Status == 200 {
						url := x.Response.URL
						enqueueHeaders(url, x.Response.RequestHeaders)
						enqueueHeaders(url, x.Response.Headers)
					}
				case *page.EventLoadEventFired:
					loadedChan <- struct{}{}
				}
			})

			err := chromedp.Run(tab,
				network.Enable(),
				// chromedp.Navigate(url.url),
				chromedp.ActionFunc(func(ctx context.Context) error {
					// Start navigation
					frameID, _, _, err := page.Navigate(u.url).Do(ctx)
					if err != nil {
						return err
					}

					// Wait for load event or timeout
					select {
					case <-loadedChan:
						// If we have depth left, enqueue all the URLs on the page
						if *verbose {
							log.Printf("%d loaded: %s", i, u.url)
						}
						if u.depth < *depth {
							body, err := page.GetResourceContent(frameID, u.url).Do(ctx)
							if err != nil {
								if x, ok := err.(*cdproto.Error); ok && x.Code == -32000 {
									// "No resource with given URL found (-32000)"
									break
								}
								return err
							}
							n, err := html.Parse(bytes.NewReader(body))
							if err != nil {
								return err
							}
							var f func(n *html.Node)
							f = func(n *html.Node) {
								if n.Type == html.ElementNode && n.Data == "a" {
									for _, a := range n.Attr {
										if a.Key == "href" && (strings.HasPrefix(a.Val, "http://") || strings.HasPrefix(a.Val, "https://")) {
											if *verbose {
												log.Printf("%d enqueueing: %s", i, a.Val)
											}
											urls <- urlEntry{url: a.Val, depth: u.depth + 1}
										}
									}
								}
								for c := n.FirstChild; c != nil; c = c.NextSibling {
									f(c)
								}
							}
							f(n)
						} // if depth

					case <-ctx.Done():
						return ctx.Err()
					}
					return nil
				}),
			)
			if err != nil {
				log.Print(err)
				return
			}
			pr.Done(u.url)
		}

		for {
			select {
			case url := <-urls:
				loadURL(url)
			default:
				return
			}
		}
	} // func fetcher
	wg.Add(int(*jobs))
	for i := uint(0); i < *jobs; i++ {
		go fetcher(int(i))
	}
	go func() {
		wg.Wait()
		close(headers)
	}()

	var headerWriter io.Writer
	if *headersFilename != "" {
		headerFile, err := os.Create(*headersFilename)
		if err != nil {
			log.Fatal(err)
		}
		defer headerFile.Close()
		headerWriter = headerFile
	} else {
		headerWriter = ioutil.Discard
	}
	headerCsv := csv.NewWriter(headerWriter)
	defer func() {
		headerCsv.Flush()
		if err := headerCsv.Error(); err != nil {
			log.Print(err)
		}
	}()
	for headerline := range headers {
		requestURL := headerline[0]
		if strings.HasPrefix(requestURL, "data:") {
			continue
		}
		u, err := url.ParseRequestURI(requestURL)
		if err != nil {
			log.Print(err)
			continue
		}
		tldp1, err := eTLDPlusOne(u.Hostname())
		if err != nil {
			log.Printf("error parsing url '%s': %s", requestURL, err)
			continue
		}
		values := append([]string{tldp1}, headerline...)
		if err := headerCsv.Write(values); err != nil {
			log.Fatal(err)
		}
	}
}
