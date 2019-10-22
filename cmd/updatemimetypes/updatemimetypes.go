// updatemimetypes outputs new MIME types with their numbers
package main

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

var mimeFilename = flag.String("mimefile", "", "mime file containing existing types & numbers")
var inURL = flag.String("url", "", "url of IANA-formatted XML list of MIME types (default is list at IANA website)")
var force = flag.Bool("force", false, "generate files from downloaded list, ignoring existing files and numbering")
var protoFile = flag.String("protofile", "", "proto definition file to output")

const defaultIANAURL = "https://www.iana.org/assignments/media-types/media-types.xml"

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

// MIME types known at time of creation to be common are encoded with tags
// 1-127. This permits one-byte varint encoding of their values. Other MIME
// types (up to tag 16,383) are encoded with two bytes, so the savings is
// small.
var commonMIMETypes = map[string]uint32{
	"application/javascript":   4,
	"application/octet-stream": 5,
	"application/xhtml+xml":    11,
	"application/xml":          15,
	"application/zip":          7,

	"image/bmp":     8,
	"image/gif":     9,
	"image/jpeg":    10,
	"image/svg+xml": 12,
	"image/tiff":    14,

	"text/css":      3,
	"text/html":     2,
	"text/plain":    1,
	"text/richtext": 13,
}

func main() {
	log.SetOutput(os.Stderr) // tool output often piped to a file
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	// internal sanity checks
	if len(commonMIMETypes) > 127 {
		log.Fatalf("too many common types: %d", len(commonMIMETypes))
	}
	typesByTag := make(map[uint32]string)
	for name, tag := range commonMIMETypes {
		// no duplicate common types or tag numbers
		if name2, ok := typesByTag[tag]; !ok {
			typesByTag[tag] = name
		} else {
			log.Fatalf("duplicate use of tag %d ('%s' and '%s')", tag, name, name2)
		}
		// no tags greater than 127 (to ensure one-byte varint encoding)
		if tag > 127 {
			log.Fatalf("common MIME type with tag >127: %s = %d", name, tag)
		}
	}

	// generate update
	buf := &bytes.Buffer{}
	csvbuf := csv.NewWriter(buf)
	n := updateTypes(csvbuf, *force)
	csvbuf.Flush()
	if n > 0 {
		// existing mimetypes file needs to be updated
		if *force {
			if *mimeFilename != "" {
				// dump entire update to the file
				mf, err := os.Create(*mimeFilename)
				if err != nil {
					log.Fatal(err)
				}
				defer mf.Close()
				if _, err := io.Copy(mf, buf); err != nil {
					log.Fatal(err)
				}
			}
		} else {
			// existing mimetypes file needs to be updated,
			// print the required updates and exit
			if _, err := io.Copy(os.Stdout, buf); err != nil {
				log.Fatal(err)
			}
			os.Exit(1)
		}
	}

	if *protoFile != "" {
		// generate protofile
		pf, err := os.Create(*protoFile)
		if err != nil {
			log.Fatal(err)
		}
		defer pf.Close()
		generateProto(pf)
	}
}

func loadMimeFile(filename string) (uint32, map[string]uint32) {
	typesByName := make(map[string]uint32)
	var nextTagNo uint32

	// collect existing mime types and tag numbers
	if *mimeFilename != "" && !*force {
		mimeFile, err := os.Open(filename)
		if err != nil {
			log.Fatalf("could not open %s: %s", *mimeFilename, err)
		}
		r := csv.NewReader(mimeFile)
		r.Comment = '#'
		r.FieldsPerRecord = 2
		r.TrimLeadingSpace = true
		r.ReuseRecord = true
		records, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		if len(records) == 0 {
			log.Print("warning: no records in MIME file")
		}
		for i, record := range records {
			name := record[0]
			tagno64, err := strconv.ParseUint(record[1], 10, 32)
			tagno := uint32(tagno64)
			if err != nil {
				log.Fatalf("%s:%d bad tag number: %s - %s", *mimeFilename, i+1, record[1], err)
			}

			// ensure names are unique
			if existingTag, ok := typesByName[name]; ok {
				if existingTag != tagno {
					log.Fatalf("tag mismatch in MIME file for '%s' (expected %d, got %d)", name, existingTag, tagno)
				}
			}

			typesByName[name] = tagno
			if tagno > nextTagNo {
				nextTagNo = tagno
			}
		}
	}
	nextTagNo++
	if nextTagNo < 128 {
		nextTagNo = 128
	}
	return nextTagNo, typesByName
}

type recordWriter interface {
	Write([]string) error
}

func updateTypes(rw recordWriter, force bool) int {
	if *mimeFilename == "" && !force {
		log.Print("existing MIME file required, specify with -mimefile or regenerate with -force")
		flag.Usage()
		os.Exit(1)
	}
	if force {
		log.Printf("*** regenerating MIME file, numbering may change ***")
	}

	if *inURL == "" {
		*inURL = defaultIANAURL
	}
	resp, err := http.Get(*inURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)
	reg := &registry{}
	if err := dec.Decode(reg); err != nil {
		log.Fatal(err)
	}

	var nextTagNo uint32
	typesByName := make(map[string]uint32)
	if *mimeFilename != "" {
		nextTagNo, typesByName = loadMimeFile(*mimeFilename)
	}

	newTypes := 0
	// ensure common types match existing file, or output new common type
	for name, tag := range commonMIMETypes {
		if mimeFileTag, ok := typesByName[name]; ok {
			if tag != mimeFileTag {
				log.Fatalf("tag mismatch in MIME file for '%s' (expected %d, got %d)", name, tag, mimeFileTag)
			}
		} else {
			// new common MIME type
			rw.Write([]string{name, fmt.Sprint(tag)})
			newTypes++
		}
	}

	// output new types from download
	for _, typereg := range reg.Registries {
		for _, record := range typereg.Records {
			name := typereg.Title + "/" + record.Name
			if _, ok := commonMIMETypes[name]; ok {
				continue
			}
			if _, ok := typesByName[name]; ok {
				continue
			}
			rw.Write([]string{name, fmt.Sprint(nextTagNo)})
			newTypes++
			nextTagNo++
		}
	}
	return newTypes
}

func generateProto(w io.Writer) {
	t := template.New("mimetypes")
	if _, err := t.Parse(protoTemplate); err != nil {
		log.Fatal(err)
	}
	// TODO: generate proto file and write to w
}

type mimeTypeDescriptor struct {
	Httpstring string
	Tag        uint32
	Aliases    []uint32
}

const protoTemplate = `
//
// DO NOT EDIT: generated file
//
// Update this file by running: make mimeproto
//
syntax = "proto3";

package web;

import "google/protobuf/descriptor.proto";

// HTTP_MIME_Type represents information about how a MIME type is represented
// on the wire in HTTP.
message MIME_Type_Descriptor {
    string http_string = 1;
    repeated uint32 aliases = 2;
}

extend google.protobuf.EnumValueOptions {
    MIME_Type_Descriptor mime_descriptor = 7987671;
}

enum MIME_Types {
    MIME_TYPE_UNUSED = 0 [(mime_descriptor)={http_string: ""; aliases: [0]}];

    // Common MIME types: tag numbers 1-127 are reserved for the most common MIME
    // types to allow them to use one-byte varint encoding.
{{ range $name, $descriptor := .CommonTypes }}
    {{ $name }} = {{ $descriptor.Tag }};
{{ end }}
    // reserved n to 127;

    // Uncommon MIME types: all other MIME types start with tag numbers 128 and up.
    //
    // A new MIME type is added here first. If it gains lots of usage, it is aliased
    // into the common list above with the following steps:
    //
    //   1. Reserve a tag number above for the promoted MIME type. E.g.
    //           reserved 42; // for "APPLICATION_AWESOME_APP"
    //   2. Add a aliases option to the entry in the uncommon list. E.g.
    //           APPLICATION_AWESOME_APP = 7376 [(mime_descriptor).aliases=42];
    //   3. Wait for this definition version to propagate
    //   4. Add the type to the common list at the reserved tag number. E.g.
    //           APPLICATION_AWESOME_APP = 42;
    //
    // Some clients may know of the type only in the uncommon list, but can use the
    // aliases value to interpret unknown values in the common list until their
    // definitions are upgraded.
    //
    // Some very old clients may know of the type only in the uncommon list and may
    // also have no knowledge the type is eligible to be promoted. Clients that see
    // a MIME type with tag <128 and do not understand it should definitely update
    // themselves.
    MIME_TYPE_UNUSED_UNCOMMON = 128 [(mime_descriptor)={http_string: ""; aliases: [0]}];
{{ range $name, $descriptor := .UncommonTypes }}
{{ end }}
}

message MIME_Type {
    oneof MIME_Type {
        MIME_Types Type = 1;
        string Other = 2;
    }
}
`
