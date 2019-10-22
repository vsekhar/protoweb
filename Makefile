GO=go
UMT=cmd/updatemimetypes

.PHONY: mimeregenerate

mimeregenerate:
	$(GO) run $(UMT)/updatemimetypes.go generateproto \
		-protofile=mime.proto \
		-force

mimeproto: mimecheck
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=$(UMT)/mimetypes.csv \
		-protofile=mime2.proto \
		generateproto

mimecheck: cmd/updatemimetypes/mimetypes.csv
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=$(UMT)/mimetypes.csv \
		updatetypes

mime.proto: mimecheck
	echo generate mime.proto