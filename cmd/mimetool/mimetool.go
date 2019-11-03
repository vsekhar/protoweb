// mimetool converts a listing of MIME types and tag numbers into a
// .proto definition file and outputs it to stdout.
package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/vsekhar/protoweb/internal/naming"
)

var mimefile = flag.String("mimetypes", "", "path to file containing MIME types and numbers")
var checkremote = flag.Bool("checkremote", false, "check remote IANA listing for new MIME types")
var url = flag.String("url", "", "url of remote IANA-formatted XML list of MIME types (default: IANA website)")
var quiet = flag.Bool("quiet", false, "don't output proto (used to check for updates)")

const defaultIANAURL = "https://www.iana.org/assignments/media-types/media-types.xml"

// struct matching the format of the IANA XML MIME types list
type registry struct {
	Title      string     `xml:"title"`
	Registries []registry `xml:"registry"`
	Records    []struct {
		Name  string `xml:"name"`
		XRefs []struct {
			Type string `xml:"type,attr"`
			Data string `xml:"data,attr"`
		} `xml:"xref"`
		Template string `xml:"file"`
	} `xml:"record"`
}

func main() {
	log.SetOutput(os.Stderr) // stdout often piped to a file
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	typesByName := make(map[string]uint32)
	var nextTagNo uint32

	if *mimefile == "" {
		log.Print("no input file, use -mimetypes")
		flag.Usage()
		os.Exit(1)
	}

	fto, err := os.Open(*mimefile)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fto)
	r.Comment = '#'
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(records) == 0 {
		log.Fatal("input file contains no records")
	}
	sane := true
	for i, record := range records {
		name := record[0]
		tagno64, err := strconv.ParseUint(record[1], 10, 32)
		if err != nil {
			log.Fatalf("%s:%d bad tag number: %s - %s", *mimefile, i+1, record[1], err)
		}
		tagno := uint32(tagno64)

		// sanity check (single entry per type, unique tag numbers)
		if existingTag, ok := typesByName[name]; ok {
			if existingTag == tagno {
				log.Printf("duplicate entry for MIME type '%s'", name)
				sane = false
			} else {
				log.Printf("multiple tags for '%s' (%d and %d)", name, existingTag, tagno)
				sane = false
			}
		}

		typesByName[name] = tagno
		if tagno > nextTagNo {
			nextTagNo = tagno
		}
	}
	nextTagNo++

	// sanity check unique tag numbers
	typesByTag := make(map[uint32]string)
	for n, t := range typesByName {
		if existingName, ok := typesByTag[t]; ok {
			log.Printf("duplicate use of tag %d (%s and %s)", t, n, existingName)
			sane = false
		}
	}

	if !sane {
		os.Exit(1)
	}

	if *checkremote {
		// load complete listing from remote server
		if *url == "" {
			*url = defaultIANAURL
		}
		resp, err := http.Get(*url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		dec := xml.NewDecoder(resp.Body)
		reg := &registry{}
		if err := dec.Decode(reg); err != nil {
			log.Fatal(err)
		}
		n := 0
		for _, typereg := range reg.Registries {
			for _, record := range typereg.Records {
				name := typereg.Title + "/" + record.Name
				name = strings.Split(name, " ")[0]
				if _, ok := typesByName[name]; !ok {
					log.Printf("%s,%d", name, nextTagNo)
					nextTagNo++
					sane = false
				}
				n++
			}
		}
		if sane {
			log.Printf("%d remote entries OK", n)
		} else {
			os.Exit(1)
		}
	}

	// produce proto file
	var e = new(entries)
	for n, t := range typesByName {
		enumName := naming.ProtoEnumName(n)
		*e = append(*e, entry{enumName, t, n})
	}
	sort.Sort(e)
	t := template.New("mimeproto")
	if _, err := t.Parse(protoTemplate); err != nil {
		log.Fatal(err)
	}
	var out io.Writer
	if *quiet {
		out = ioutil.Discard
	} else {
		out = os.Stdout
	}
	err = t.Execute(out, e)
	if err != nil {
		log.Fatal(err)
	}
}

type entry struct {
	EnumName string
	Tag      uint32
	HTTPName string
}
type entries []entry

func (e *entries) Len() int {
	return len(*e)
}

func (e *entries) Less(i, j int) bool {
	return (*e)[i].EnumName < (*e)[j].EnumName
}

func (e *entries) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

const protoTemplate = `//
// DO NOT EDIT: generated file
//
// Update this file by running: make proto/mime.proto
//
syntax = "proto3";

package web;

import "google/protobuf/descriptor.proto";

import "charset.proto";

message MIMETypeDescriptor {
  string http_string = 1;
}

extend google.protobuf.EnumValueOptions {
  MIMETypeDescriptor mime_descriptor = 7987671;
}

enum MIMETypes {
  MIME_TYPE_UNUSED = 0 [(mime_descriptor).http_string=""];
{{ range $_, $entry := . }}  {{ $entry.EnumName }} = {{ $entry.Tag }} [(mime_descriptor).http_string="{{ $entry.HTTPName }}"];
{{ end }}}

message MIMEType {
  oneof MIMEType {
    MIMETypes type = 1;
    string custom = 2;
  }
  Charsets charset = 3;
}
`
