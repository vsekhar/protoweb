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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	urls := make(chan urlEntry, 10000)
	headers := make(chan []string, 1000)
	enqueueHeaders := func(url string, hmap map[string]interface{}) {
		for h, vs := range hmap {
			if strings.HasPrefix(strings.ToLower(h), "x-google") {
				continue
			}
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
		log.Printf("enqueueing %s", q.url)
		urls <- q
	}

	fetcher := func(i int) {
		defer wg.Done()
		defer log.Printf("terminating fetcher")

		browser, closeBrowser := chromedp.NewContext(
			context.Background(),
			chromedp.WithLogf(log.Printf),
		)
		defer closeBrowser()
		// start the browser so that tabs (not browsers) are created below
		if err := chromedp.Run(browser); err != nil {
			log.Fatal(err)
		}

		loadURL := func(u urlEntry) {
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

			log.Printf("%d: fetching %s @ depth %d", i, u.url, u.depth)
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
						log.Printf("%d loaded: %s", i, u.url)
						if u.depth < *depth {
							body, err := page.GetResourceContent(frameID, u.url).Do(ctx)
							if err != nil {
								return err
							}
							// requestID := x.RequestID
							// c := chromedp.FromContext(tab)
							// https://github.com/chromedp/chromedp/issues/326#issuecomment-495788351
							// body, err := network.GetResponseBody(requestID).Do(cdp.WithExecutor(tab, c.Target))
							// if err != nil {
							//	log.Printf("%d terminating: %s", i, err)
							//	return
							// }
							n, err := html.Parse(bytes.NewReader(body))
							if err != nil {
								return err
							}
							var f func(n *html.Node)
							f = func(n *html.Node) {
								if n.Type == html.ElementNode && n.Data == "a" {
									for _, a := range n.Attr {
										if a.Key == "href" && (strings.HasPrefix(a.Val, "http://") || strings.HasPrefix(a.Val, "https://")) {
											log.Printf("%d enqueueing: %s", i, a.Val)
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
			if u.depth < *depth {
				// add links to urls with url.depth+1
			}
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
