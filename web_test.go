package web_test // https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c

import (
	"bytes"
	"net/http"
	"testing"

	proto "github.com/golang/protobuf/proto"
	web "github.com/vsekhar/protoweb"
)

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestReq2Proto(t *testing.T) {
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	header := map[string][]string{
		"If-None-Match": {`W/"67ab43"`, "54ed21", "7892dd"},
	}
	for n, vs := range header {
		for _, v := range vs {
			req.Header.Add(n, v)
		}
	}
	reqproto, err := web.Req2Proto(req)
	if err != nil {
		t.Fatal(err)
	}
	if reqproto.Header.Other != nil {
		t.Errorf("unparsed headers: %v", reqproto.Header.Other)
	}
	if !equalStringSlices(reqproto.Header.If_None_Match.Strings, header["If-None-Match"]) {
		t.Errorf("bad header If-None-Match: %v", reqproto.Header.If_None_Match.Strings)
	}
}

// TODO: makeResponse, makeProtoResponse, capture actual server response
// bytes?

func BenchmarkRequestSerialization(b *testing.B) {
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		b.Fatal(err)
	}
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
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		b.Fatal(err)
	}
	reqproto, err := web.Req2Proto(req)
	if err != nil {
		b.Fatal(err)
	}
	// serialize once to pre-allocate buffer
	buf := proto.NewBuffer([]byte{})
	if err := buf.Marshal(reqproto); err != nil {
		b.Fatal(err)
	}
	b.ReportMetric(float64(len(buf.Bytes())), "wirebytes/op")
	buf.Reset()
	for i := 0; i < b.N; i++ {
		if err := buf.Marshal(reqproto); err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}
