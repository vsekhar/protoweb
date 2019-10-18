package web

import (
	fmt "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// Req2Proto converts an http.Request into a Request protobuffer or
// returns an error.
func Req2Proto(req *http.Request) (*Request, error) {
	ret := new(Request)
	parsedFields := []string{}
	parsed := func(s string) { parsedFields = append(parsedFields, s) }
	method, ok := Request_Method_value[req.Method]
	if !ok {
		return nil, fmt.Errorf("bad method: %s", req.Method)
	}
	ret.Method = Request_Method(method)
	parsed("method")
	if req.URL.String() == "*" {
		ret.URI = &Request_URI_Wildcard{}
	} else {
		ret.URI = &Request_URI_String{URI_String: req.URL.String()}
	}
	parsed("uri")
	ret.Headers = new(Request_Headers)
	for header, values := range req.Header {
		header = strings.ToLower(header)
		lastvalue := ""
		if len(values) > 0 {
			lastvalue = values[len(values)-1]
		}
		knownHeader := true
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
			knownHeader = false
			for _, v := range values {
				ret.Headers.Other = append(ret.Headers.Other, &KeyValue{Key: header, Value: v})
			}
		}
		if knownHeader {
			parsed(header)
		}
	}
	return ret, nil
}

// Resp2Proto converts an http.Response into a Response protobuffer
// or returns an error.
func Resp2Proto(resp *http.Response) (*Response, error) {
	ret := new(Response)
	parsedFields := []string{}
	parsed := func(s string) { parsedFields = append(parsedFields, s) }
	if _, ok := Response_Code_Value_name[int32(resp.StatusCode)]; !ok {
		return nil, fmt.Errorf("unknown response code: %d", resp.StatusCode)
	}
	ret.Code = Response_Code_Value(resp.StatusCode)
	parsed("code")
	ret.Headers = new(Response_Headers)
	for header, values := range resp.Header {
		header = strings.ToLower(header)
		lastvalue := ""
		if len(values) > 0 {
			lastvalue = values[len(values)-1]
		}
		knownHeader := true
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
			optionnumber, ok := Response_Headers_X_Frame_Options_Value_value[value]
			if !ok {
				return nil, fmt.Errorf("unknown value for X-Frame-Options: %s", lastvalue)
			}
			ret.Headers.X_Frame_Options = Response_Headers_X_Frame_Options_Value(optionnumber)
		case "expires":
			date, err := time.Parse(time.RFC1123, lastvalue)
			if err != nil {
				datepb, err := ptypes.TimestampProto(date)
				if err != nil {
					return nil, fmt.Errorf("unable to create timestamp proto: %s", err)
				}
				ret.Headers.Expires = &Response_Headers_Expires_Date{Expires_Date: datepb}
			} else {
				ret.Headers.Expires = &Response_Headers_Expires_Already{}
			}
		case "content-type":
			// TODO: parse type/subtype;parameter=value
			// TODO: enumerate MIME types: https://www.iana.org/assignments/media-types/media-types.xhtml
			parts := strings.Split(lastvalue, ";")
			if len(parts) == 0 || len(parts) > 2 {
				return nil, fmt.Errorf("bad Content-Type: %s", lastvalue)
			}
			contentType := strings.TrimSpace(parts[0])
			ret.Headers.Content_Type = &Response_Headers_Content_Type_Message{}
			mimeTypeNumber, ok := Response_Headers_Content_Type_Message_Common_MIME_Types_value[contentType]
			if ok {
				ret.Headers.Content_Type.Content_Type_Message = &Response_Headers_Content_Type_Message_Common_MIME_Type{
					Common_MIME_Type: Response_Headers_Content_Type_Message_Common_MIME_Types(mimeTypeNumber),
				}
			} else {
				ret.Headers.Content_Type.Content_Type_Message = &Response_Headers_Content_Type_Message_Other{
					Other: contentType,
				}
			}
			if len(parts) > 1 {
				nameValue := strings.Split(parts[1], "=")
				if len(nameValue) > 0 {
					ret.Headers.Content_Type.Parameter = &KeyValue{
						Key: nameValue[0],
					}
					if len(nameValue) > 1 {
						ret.Headers.Content_Type.Parameter.Value = nameValue[1]
					}
				}
			}
		case "set-cookie":
			setcookie := &Response_Headers_Set_Cookie_Message{}
			_ = setcookie
			for _, s := range values {
				log.Println(s)
			}
		default:
			knownHeader = false
			for _, v := range values {
				ret.Headers.Other = append(ret.Headers.Other, &KeyValue{Key: header, Value: v})
			}
		}
		if knownHeader {
			parsed(header)
		}
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}
	ret.Body = bodybytes

	return ret, nil
}
