package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/vsekhar/protoweb/internal/naming"
)

const defaultIANAURL = "https://www.iana.org/assignments/character-sets/character-sets.xml"

// struct matching the format of the IANA XML charset list
type xRef struct {
	Type string `xml:"type,attr"`
	Data string `xml:"data,attr"`
}
type record struct {
	Name        string `xml:"name"`
	XRefs       []xRef `xml:"xref"`
	Value       uint64 `xml:"value"`
	Description struct {
		XRefs []xRef `xml:"xref"`
		Data  string `xml:",chardata"`
	} `xml:"description"`
	PreferredAlias string   `xml:"preferred_alias"`
	Aliases        []string `xml:"alias"`
	ProtoName      string   `xml:"-"`
	ProtoTag       uint64   `xml:"-"`
}
type records []record

func (e *records) Len() int           { return len(*e) }
func (e *records) Less(i, j int) bool { return (*e)[i].ProtoTag < (*e)[j].ProtoTag }
func (e *records) Swap(i, j int)      { (*e)[i], (*e)[j] = (*e)[j], (*e)[i] }

type registry struct {
	ID         string     `xml:"id,attr"`
	Registries []registry `xml:"registry"`
	Records    []record   `xml:"record"`
}

func main() {
	log.SetOutput(os.Stderr) // stdout often piped to a file
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	resp, err := http.Get(defaultIANAURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)
	reg := &registry{}
	if err := dec.Decode(reg); err != nil {
		log.Fatal(err)
	}
	var records = new(records)
	*records = make([]record, 0)
	for _, charsetreg := range reg.Registries {
		for _, record := range charsetreg.Records {
			if record.Name == "UTF-8" {
				record.ProtoTag = 1
			} else {
				record.ProtoTag = record.Value
			}
			record.ProtoName = naming.ProtoEnumName(record.Name)
			for i := range record.Aliases {
				// Amiga-1251 and friends...
				record.Aliases[i] = strings.ReplaceAll(record.Aliases[i], "\n", " ")
			}
			*records = append(*records, record)
		}
	}
	sort.Sort(records)
	t := template.New("charsetproto")
	if _, err := t.Parse(protoTemplate); err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, records)
	if err != nil {
		log.Fatal(err)
	}
}

const protoTemplate = `//
// DO NOT EDIT: generated file
//
// Update this file by running: make proto/charset.proto
//
syntax = "proto3";

package web;

import "google/protobuf/descriptor.proto";

message CharsetDescriptor {
  string http_name = 1;
  uint64 mibenum = 2;
  string preferred_alias = 3;
  repeated string aliases = 4;
}

extend google.protobuf.EnumValueOptions {
  CharsetDescriptor charset_descriptor = 7982317;
}

enum Charsets {
  CHARSET_UNSPECIFIED = 0;
{{ range $_, $record := . }}  {{ $record.ProtoName }} = {{ $record.ProtoTag }} [
    (charset_descriptor) = {
      http_name: "{{ $record.Name }}"
      preferred_alias: "{{ $record.PreferredAlias }}"
{{ range $_, $alias := $record.Aliases }}      aliases: "{{ $alias }}"
{{ end }}    }
  ];
{{ end }}}

message Charset {
  oneof Charset {
    Charsets charset = 1;
    string custom = 2;
  }
}
`
