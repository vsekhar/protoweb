package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	web "github.com/vsekhar/protoweb"
)

func req2Proto(req *http.Request) (*web.Request, error) {
	ret := new(web.Request)
	method, ok := web.Request_Method_value[req.Method]
	if !ok {
		return nil, fmt.Errorf("bad method: %s", req.Method)
	}
	ret.Method = web.Request_Method(method)
	if req.URL.String() == "*" {
		ret.URI = &web.Request_URI_Wildcard{}
	} else {
		ret.URI = &web.Request_URI_String{URI_String: req.URL.String()}
	}
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

func resp2Proto(resp *http.Response) (*web.Response, error) {
	ret := new(web.Response)
	if _, ok := web.Response_Code_Value_name[int32(resp.StatusCode)]; !ok {
		return nil, fmt.Errorf("unknown response code: %d", resp.StatusCode)
	}
	ret.Code = web.Response_Code_Value(resp.StatusCode)
	ret.Headers = new(web.Response_Headers)
	for header, values := range resp.Header {
		header = strings.ToLower(header)
		lastvalue := ""
		if len(values) > 0 {
			lastvalue = values[len(values)-1]
		}
		switch header {
		case "date":
			date, err := time.Parse(time.RFC1123, lastvalue)
			if err != nil {
				return nil, fmt.Errorf("bad date: %s", values)
			}
			datepb, err := ptypes.TimestampProto(date)
			if err != nil {
				return nil, fmt.Errorf("unable to create timestamp proto: %s", err)
			}
			ret.Headers.Date = datepb
		case "server":
			ret.Headers.Server = lastvalue
		case "x-xss-protection":
			ret.Headers.X_XSS_Protection = lastvalue
		case "x-frame-options":
			value := strings.ToUpper(lastvalue)
			optionnumber, ok := web.Response_Headers_X_Frame_Options_Value_value[value]
			if !ok {
				return nil, fmt.Errorf("unknown value for X-Frame-Options: %s", lastvalue)
			}
			ret.Headers.X_Frame_Options = web.Response_Headers_X_Frame_Options_Value(optionnumber)
		case "expires":
			date, err := time.Parse(time.RFC1123, lastvalue)
			if err != nil {
				datepb, err := ptypes.TimestampProto(date)
				if err != nil {
					return nil, fmt.Errorf("unable to create timestamp proto: %s", err)
				}
				ret.Headers.Expires = &web.Response_Headers_Expires_Date{Expires_Date: datepb}
			} else {
				ret.Headers.Expires = &web.Response_Headers_Expires_Already{}
			}
		case "content-type":
			// TODO: parse type/subtype;parameter=value
			// TODO: enumerate MIME types: https://www.iana.org/assignments/media-types/media-types.xhtml
			parts := strings.Split(lastvalue, ";")
			if len(parts) == 0 || len(parts) > 2 {
				return nil, fmt.Errorf("bad Content-Type: %s", lastvalue)
			}
			contentType := strings.TrimSpace(parts[0])
			ret.Headers.Content_Type = &web.Response_Headers_Content_Type_Message{}
			mimeTypeNumber, ok := web.Response_Headers_Content_Type_Message_Common_MIME_Types_value[contentType]
			if ok {
				ret.Headers.Content_Type.Content_Type_Message = &web.Response_Headers_Content_Type_Message_Common_MIME_Type{
					Common_MIME_Type: web.Response_Headers_Content_Type_Message_Common_MIME_Types(mimeTypeNumber),
				}
			} else {
				ret.Headers.Content_Type.Content_Type_Message = &web.Response_Headers_Content_Type_Message_Other{
					Other: contentType,
				}
			}
			if len(parts) > 1 {
				nameValue := strings.Split(parts[1], "=")
				if len(nameValue) > 0 {
					ret.Headers.Content_Type.Parameter = &web.KeyValue{
						Key: nameValue[0],
					}
					if len(nameValue) > 1 {
						ret.Headers.Content_Type.Parameter.Value = nameValue[1]
					}
				}
			}
		case "set-cookie":
			setcookie := &web.Response_Headers_Set_Cookie_Message{}
			_ = setcookie
			for _, s := range values {
				log.Println(s)
			}
		case "cache-control":

		case "p3p":
			// ignored
		default:
			return nil, fmt.Errorf("unknown header: %s:%s", header, values)
		}
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}
	ret.Body = bodybytes

	return ret, nil
}

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
	if *dumprawresponse {
		log.Printf("Response Headers:")
		for k, v := range resp.Header {
			log.Printf("%s: %s", k, v)
		}
	}
	protoresp, err := resp2Proto(resp)
	if err != nil {
		log.Fatal(err)
	}
	protorespbytes, err := proto.Marshal(protoresp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response bytes - proto:", len(protorespbytes))
}
