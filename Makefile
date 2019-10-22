GO=go
UMT=cmd/updatemimetypes

all: webproto

webproto: web.proto mime.proto
	protoc --go_out=. web.proto mime.proto

# this target will fail if IANA list includes MIME types not in mimetypes.csv.
# run `make mimecheckforupdates` to print the required updates
mime.proto: mimetypes.csv
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv \
		-proto_out=mime.proto

dist: mime.proto

.PHONY: mimeregenerate mimecheckforupdates

# manually regenerate mimetypes.csv and mime.proto (possibly renumbering types)
mimeregenerate:
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv \
		-proto_out=mime.proto \
		-force

# manually trigger check for updates to mime types from IANA list
mimecheckforupdates:
	$(GO) run $(UMT)/updatemimetypes.go \
		-mimefile=mimetypes.csv

