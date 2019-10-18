package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/golang/protobuf/proto"
	web "github.com/vsekhar/protoweb"
)

func main() {
	var (
		dumprawrequest   = flag.Bool("dumprawrequest", false, "dump raw request")
		dumpprotorequest = flag.Bool("dumpprotorequest", false, "dump proto request")
		dumprawresponse  = flag.Bool("dumprawresponse", false, "dump raw response")
	)
	flag.Parse()

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Add something non-default
	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	// Get size of raw representation
	rawreq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Request bytes - raw:", len(rawreq))
	if *dumprawrequest {
		log.Printf("%s", rawreq)
	}

	// Get size of proto representation
	protoreq, err := web.Req2Proto(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Unparsed headers (included in Other): %v", protoreq.Headers.Other)
	protoreqbytes, err := proto.Marshal(protoreq)
	if err != nil {
		log.Fatal(err)
	}
	if *dumpprotorequest {
		log.Printf("Proto request: %x", protoreqbytes)
	}
	log.Println("Request bytes - proto:", len(protoreqbytes))

	log.Printf("--- Fetching ---")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// strip (large) internal headers
	for header := range resp.Header {
		header = strings.ToLower(header)
		if strings.HasPrefix(header, "x-google-") {
			resp.Header.Del(header)
		}
	}
	rawresp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response bytes - raw:", len(rawresp))
	if *dumprawresponse {
		log.Printf("Response Headers:")
		for k, v := range resp.Header {
			log.Printf("%s: %s", k, v)
		}
	}
	protoresp, err := web.Resp2Proto(resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Unparsed headers (included in Other): %v", protoresp.Headers.Other)
	protorespbytes, err := proto.Marshal(protoresp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response bytes - proto:", len(protorespbytes))
}
