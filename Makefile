GO=go

.PHONY: mimecheckforupdates test

all: protos

protos: web.pb.go mime.pb.go

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. $<

mimetypesall.csv: mimetypescommon.csv

# this target will fail if IANA list includes MIME types not in mimetypes.csv.
# run `make mimecheckforupdates` to print the required updates
mime.proto: mimetypes.csv cmd/mimetool/mimetool.go
	$(GO) run cmd/mimetool/mimetool.go \
		-mimetypes=mimetypes.csv \
		> $@

test: all
	$(GO) test

dist: test protocheck all

# manually trigger check for updates to mime types from IANA list
mimecheckforupdates:
	$(GO) run cmd/mimetool/mimetool.go \
		-mimetypes=mimetypes.csv \
		-checkremote \
		-quiet

protocheck:
	prototool break check
