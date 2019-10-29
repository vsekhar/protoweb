package web_test // https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c

import (
	"bytes"
	context "context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"net/http"
	"testing"

	proto "github.com/golang/protobuf/proto"
	quic "github.com/lucas-clemente/quic-go"
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
	if !equalStringSlices(reqproto.Header.If_None_Match, header["If-None-Match"]) {
		t.Errorf("bad header If-None-Match: %v", reqproto.Header.If_None_Match)
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

func quicSelfConnect() {

}

func startQUICServer(t *testing.T) {
	l, err := quic.ListenAddr("localhost:", generateTLSConfig(), nil)
	if err != nil {
		t.Fatal(err)
	}
	sess, err := l.Accept(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_ = stream
}

func TestQUIC(t *testing.T) {
	go startQUICServer(t)
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
