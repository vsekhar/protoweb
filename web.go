package web

import (
	fmt "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vsekhar/protoweb/internal/naming"
)

// Req2Proto converts an http.Request into a Request protobuffer or
// returns an error.
func Req2Proto(req *http.Request) (*Request, error) {
	ret := new(Request)
	method, ok := Method_value[req.Method]
	if !ok {
		return nil, fmt.Errorf("bad method: %s", req.Method)
	}
	ret.Method = Method(method)
	if req.URL.String() == "*" {
		ret.URI = &Request_Wildcard{Wildcard: true}
	} else {
		ret.URI = &Request_UriString{UriString: req.URL.String()}
	}
	ret.Header = new(RequestHeaders)
	for header, values := range req.Header {
		header = strings.ToLower(header)
		lastvalue := ""
		if len(values) > 0 {
			lastvalue = values[len(values)-1]
		}
		switch header {
		case "host":
			ret.Header.Host = lastvalue
		case "user-agent":
			ret.Header.UserAgent = lastvalue
		case "if-none-match":
			if ret.Header.IfNoneMatch == nil {
				ret.Header.IfNoneMatch = make([]string, 0, len(values))
			}
			ret.Header.IfNoneMatch = append(ret.Header.IfNoneMatch, values...)
		case "accept-encoding":
			if ret.Header.Accept == nil {
				ret.Header.Accept = new(RequestHeaders_AcceptHeaders)
			}
			for _, val := range values {
				log.Printf("accept-encoding: %s", val)
				parts := strings.Split(val, ";")
				if len(parts) > 2 {
					log.Fatalf("bad accept-encoding value: %s", val)
				}
				enc := &RequestHeaders_AcceptHeaders_Encoding{}
				if parts[0] == "*" {
					enc.Encoding = &RequestHeaders_AcceptHeaders_Encoding_Wildcard{Wildcard: true}
				} else {
					parts[0] = naming.ProtoEnumName(parts[0])
					if encenum, ok := Encodings_value[parts[0]]; ok {
						enc.Encoding = &RequestHeaders_AcceptHeaders_Encoding_Value{
							Value: Encodings(encenum),
						}
					} else {
						log.Fatalf("unknown accept-encoding value: %s", val)
					}
				}
				if len(parts) > 1 {
					parts[1] = strings.ToLower(parts[1])
					n, err := fmt.Sscanf(parts[1], "q=%f", &enc.Q)
					if err != nil || n != 1 {
						log.Fatalf("bad accept-encoding value: %s", val)
					}
				}
				ret.Header.Accept.Encoding = append(ret.Header.Accept.Encoding, enc)
			}
		default:
			for _, v := range values {
				ret.Header.Other = append(ret.Header.Other, &KeyValue{Key: header, Value: v})
			}
		}
	}
	return ret, nil
}

// Resp2Proto converts an http.Response into a Response protobuffer
// or returns an error.
func Resp2Proto(resp *http.Response) (*Response, error) {
	ret := new(Response)
	if _, ok := StatusCodes_name[int32(resp.StatusCode)]; !ok {
		return nil, fmt.Errorf("unknown response code: %d", resp.StatusCode)
	}
	ret.Status = StatusCodes(resp.StatusCode)
	ret.Header = new(Response_Headers)
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
			ret.Header.Date = datepb
		case "server":
			ret.Header.Server = lastvalue
		case "x-xss-protection":
			ret.Header.XXssProtection = lastvalue
		case "x-frame-options":
			value := strings.ToUpper(lastvalue)
			optionnumber, ok := Response_Headers_XFrameOptionsValue_value[value]
			if !ok {
				return nil, fmt.Errorf("unknown value for X-Frame-Options: %s", lastvalue)
			}
			ret.Header.XFrameOptions = Response_Headers_XFrameOptionsValue(optionnumber)
		case "expires":
			date, err := time.Parse(time.RFC1123, lastvalue)
			if err != nil {
				datepb, err := ptypes.TimestampProto(date)
				if err != nil {
					return nil, fmt.Errorf("unable to create timestamp proto: %s", err)
				}
				ret.Header.Expires = &Response_Headers_ExpiresDate{ExpiresDate: datepb}
			} else {
				ret.Header.Expires = &Response_Headers_ExpiresAlready{}
			}
		case "content-type":
			// TODO: parse type/subtype;parameter=value
			parts := strings.Split(lastvalue, ";")
			if len(parts) == 0 || len(parts) > 2 {
				return nil, fmt.Errorf("bad Content-Type: %s", lastvalue)
			}
			contentType := strings.TrimSpace(parts[0])
			ret.Header.ContentType = &MIMEType{}
			mimeTypeNumber, ok := MIMETypes_value[contentType]
			if ok {
				ret.Header.ContentType.MIMEType = &MIMEType_Type{MIMETypes(mimeTypeNumber)}
			} else {
				ret.Header.ContentType.MIMEType = &MIMEType_Custom{contentType}
			}
			if len(parts) > 1 {
				log.Fatalf("TODO: Content-type parameter is unimplemented")
			}
		case "set-cookie":
			setcookie := &Response_Headers_SetCookieMessage{}
			_ = setcookie
			for _, s := range values {
				log.Println(s)
			}
		default:
			for _, v := range values {
				ret.Header.Other = append(ret.Header.Other, &KeyValue{Key: header, Value: v})
			}
		}
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}
	ret.Body = bodybytes

	return ret, nil
}
