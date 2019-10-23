GO=go

.PHONY: mimecheckforupdates test

all: webproto

webproto: web.proto mime.proto
	protoc --go_out=. web.proto mime.proto

mimetypesall.csv: mimetypescommon.csv

# this target will fail if IANA list includes MIME types not in mimetypes.csv.
# run `make mimecheckforupdates` to print the required updates
mime.proto: mimetypes.csv cmd/mimetool/mimetool.go
	$(GO) run cmd/mimetool/mimetool.go \
		-mimetypes=mimetypes.csv \
		> $@

test:
	$(GO) test

dist: test webproto

# manually trigger check for updates to mime types from IANA list
mimecheckforupdates:
	$(GO) run cmd/mimetool/mimetool.go \
		-mimetypes=mimetypes.csv \
		-checkremote \
		-quiet
