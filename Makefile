GO=go
PROTOS=$(wildcard proto/*.proto)

UNAME_S := $(uname -s)
CHROME :=
ifeq ($(UNAME_S),Darwin)
	CHROME += /Applications/Google Chrome.app/Contents/MacOS/Google Chrome
endif
ifeq ($(UNAME_S),Linux)
	CHROME += /opt/google/chrome/chrome
endif

.PHONY: mimecheckforupdates test

all: protos

protos: $(PROTOS)
	prototool generate

protolint: $(PROTOS)
	prototool lint

protoformat: $(PROTOS)
	prototool format -w

# this target will fail if IANA list includes MIME types not in mimetypes.csv.
# run `make mimecheckforupdates` to print the required updates
mime.proto: mimetypes.csv cmd/mimetool/mimetool.go
	$(GO) run cmd/mimetool/mimetool.go \
		-mimetypes=mimetypes.csv \
		> $@

test: all
	$(GO) test

bench: all
	$(GO) test -benchmem  -run=^$ github.com/vsekhar/protoweb -bench .

dist: all test protolint protocheck protoformat

# manually trigger check for updates to mime types from IANA list
mimecheckforupdates:
	$(GO) run cmd/mimetool/mimetool.go \
		-mimetypes=mimetypes.csv \
		-checkremote \
		-quiet

protocheck:
	prototool break check

captureheaders: testdata/capture/capture.js
	echo $(UNAME_S) && \
	cd testdata && \
	node capture/capture.js --chromepath=$(CHROME)
