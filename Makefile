GO=go
UMT=cmd/updatemimetypes

.PHONY: mimeregenerate mimecheck dist

all: webproto

webproto: web.proto mime.proto
	protoc --go_out=. web.proto mime.proto

mimeregenerate:
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv \
		-proto_out=mime.proto \
		-force

mimecheck:
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv

mime.proto: mimetypes.csv
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv \
		-proto_out=mime.proto

dist: mime.proto
