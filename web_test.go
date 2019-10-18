package web

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	proto "github.com/golang/protobuf/proto"
)

func makeRequest() *http.Request {
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Add something non-default
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	return req
}

func makeRequestProto() []byte {
	req := makeRequest()
	protoreq, err := Req2Proto(req)
	if err != nil {
		log.Fatal(err)
	}
	protoreqbytes, err := proto.Marshal(protoreq)
	if err != nil {
		log.Fatal(err)
	}
	return protoreqbytes
}

// TODO: makeResponse, makeProtoResponse, capture actual server response
// bytes?

func BenchmarkRequestSerialization(b *testing.B) {
	req := makeRequest()
	// serialize once to pre-allocate buffer
	buf := &bytes.Buffer{}
	if err := req.Write(buf); err != nil {
		b.Fatal(err)
	}
	b.ReportMetric(float64(buf.Len()), "wirebytes/op")
	buf.Reset()
	for i := 0; i < b.N; i++ {
		if err := req.Write(buf); err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

func BenchmarkProtoRequestSerialization(b *testing.B) {
	req := makeRequest()
	reqproto, err := Req2Proto(req)
	if err != nil {
		b.Fatal(err)
	}
	// serialize once to pre-allocate buffer
	buf := proto.NewBuffer([]byte{})
	if err := buf.Marshal(reqproto); err != nil {
		b.Fatal(err)
	}
	bufbytes := buf.Bytes()
	b.ReportMetric(float64(len(bufbytes)), "wirebytes/op")
	buf.Reset()
	for i := 0; i < b.N; i++ {
		if err := buf.Marshal(reqproto); err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}
