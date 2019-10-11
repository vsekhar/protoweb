package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/golang/protobuf/proto"

	web "github.com/vsekhar/protoweb"
)

func req2Proto(req *http.Request) (*web.Request, error) {
	ret := new(web.Request)
	method, ok := web.Request_Method_value[req.Method]
	if !ok {
		return nil, fmt.Errorf("bad method: %s", req.Method)
	}
	ret.Method = web.Request_Method(method)
	ret.Uri = req.URL.String()
	ret.Headers = new(web.Request_Headers)
	for header, values := range req.Header {
		header = strings.ToLower(header)
		lastvalue := ""
		if len(values) > 0 {
			lastvalue = values[len(values)-1]
		}
		switch header {
		case "host":
			ret.Headers.Host = lastvalue
		case "user-agent":
			ret.Headers.User_Agent = lastvalue
		case "if-none-match":
			ret.Headers.If_None_Match = lastvalue
		case "accept-encoding":
			ret.Headers.Accept_Encoding = lastvalue
		default:
			return nil, fmt.Errorf("unknown header: %s:%s", header, values)
		}
	}
	return ret, nil
}

func main() {
	var (
		dumprawrequest   = flag.Bool("dumprawrequest", false, "dump raw request")
		dumpprotorequest = flag.Bool("dumpprotorequest", false, "dump proto request")
	)
	flag.Parse()

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
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
	protoreq, err := req2Proto(req)
	if err != nil {
		log.Fatal(err)
	}
	protoreqbytes, err := proto.Marshal(protoreq)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: why is this empty?
	if *dumpprotorequest {
		log.Printf("Proto request: %x", protoreqbytes)
	}
	log.Println("Request bytes - proto:", len(protoreqbytes))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	rawresp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response bytes - raw:", len(rawresp))
}
