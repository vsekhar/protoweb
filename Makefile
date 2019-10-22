GO=go
UMT=cmd/updatemimetypes

.PHONY: mimeregenerate

mimeregenerate:
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv \
		-protofile=mime.proto \
		-force

mimeproto: mimecheck
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv \
		-protofile=mime2.proto

mimecheck: cmd/updatemimetypes/mimetypes.csv
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv

mime.proto: mimecheck
	echo generate mime.proto