package main

import (
	"context"
	"flag"
	"log"

	"github.com/chromedp/chromedp"
)

var jobs = flag.Uint("jobs", 1, "number of simultaneous jobs to start for a task")
var sitesFilename = flag.String("sitesfile", "", "path and filename containing seed URLs, one per line")
var skippedFilename = flag.String("skippedfile", "", "path and filename to write skipped URLs")
var outputFilename = flag.String("o", "", "output CSV: {url},{headername},{headervalue}")

func main() {
	flag.Parse()

	// load sites.txt
	// create queue by URL
	// send a TLD to a given task
	// task reads all URLs for that TLD until no more, then gets a new TLD
	// task should read it's own writes (enqueued URLs from a given TLD)
	// tld := publicsuffix.EffectiveTLDPlusOne(domain)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	if err := chromedp.Run(ctx, chromedp.Navigate("https://google.com")); err != nil {
		log.Fatal(err)
	}

	// tldChan is a channel that returns channels on which URLs will be sent
	tldChan := make(chan chan string)

	// urlQueueChan is a channel to enqueue new URLs
	urlQueueChan := make(chan string)

	_, _ = tldChan, urlQueueChan
}
