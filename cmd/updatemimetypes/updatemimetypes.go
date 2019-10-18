package main

import (
	"flag"
	"log"
	"text/template"
)

const mimeTemplate = `
syntax = "proto3";

package web;

// https://tools.ietf.org/html/rfc2046#section-2
enum Top_Level_MIME_Types {
    TOP_LEVEL_MIME_TYPE_UNUSED = 0;
    TEXT = 1;
    IMAGE = 2;
    AUDIO = 3;
    VIDEO = 4;
    APPLICATION = 5;
}

enum MIME_Types {
    MIME_TYPE_UNUSED = 0;
{{ range $name, $value := . }}
    {{ $name }} = {{ $value }};
{{ end }}
}

message MIME_Type {
    Top_Level_MIME_Types Top_Level = 1;
    MIME_Types Type = 2;
}
`

var outFile = flag.String("out", "", "file to output MIME type declarations")

func main() {
	t := template.New("mimetypes")
	if _, err := t.Parse(mimeTemplate); err != nil {
		log.Fatal(err)
	}
}

var commonMIMETypes = map[string]int{
	"text/html": 1,
	"text/css":  2,
}

// TODO: make number assignments stable in this repo. create a list and store
// it here, hit the network to check all official ones are on the list.
// TODO: choose 127 most common MIME types to optimize varint usage

// two lists: common mime types (max 127), other mime types
// few types are promoted to common, spots reserved for future popular ones
// *all* mime types appear in other mime types
// for a non-common mime type: server must encode other
// for a common mime type: server must encode common
// protobuf includes oneof for common or other
// mime types that start uncommon but are promoted to common appear in *both* lists
// clients that don't recognize a mime type can't be expected to handle it
