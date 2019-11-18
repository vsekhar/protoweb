# Protoweb

[![Build Status](https://travis-ci.org/vsekhar/protoweb.svg?branch=master)](https://travis-ci.org/vsekhar/protoweb)

Protoweb is a thought experiment to evaluate a QUIC + Protobuf implementation of web standards.

## But why?

Protobufs have several advantages over string- and byte-based protocols like HTTP:

  * Versionable: clients and servers implementing different versions of the protocol can still communicate
  * Schema'd: an external schema allows 
  * Machine readable: clients, servers and other tools can analyze the protocol as well as how it is evolving (for example, to check for backward compatibility)
  * Efficient: common values can be compactly encoded on the wire, with translations available in the schema (kind of like a perfect Huffman code)

## Roadmap

  * Capture some real-ish traffic for replay (Puppeteer?)
  * Define protocol buffers to handle captured traffic
  * Implement counting transport for benchmarking
  * Benchmark traffic replays with HTTP, HTTP/2 and protoweb
  * Implement Go HTTP server and client using protoweb (with HTTP/2 and HTTP fallback)

## Prerequisites for development

The library can be imported as is via Go programs.

For development, testing and distribution packaging, some tools need to be installed.

Install the protocol buffer compiler. For Debian/Ubuntu:

    $ sudo apt install protobuf-compiler

Install some Go protobuf tools:

    $ go get github.com/golang/protobuf/protoc-gen-go
    $ go get google.golang.org/grpc
    $ go get github.com/uber/prototool/cmd/prototool

To recapture header testdata, [install Google Chrome](https://www.google.com/chrome/).

## Discussion

### What about [CBOR](https://cbor.io/)

CBOR is a binary encoding format for serializing key-value pairs. It can be thought of as a more efficient JSON. It is schema-less, self-describing, and non-human readable;.

A key claimed feature of CBOR, which it shares with JSON, is that a decoder does not need access to a schema to decode a CBOR message. This "feature" is suspect. While it is true a CBOR/JSON _library_ does not need a schema, the library at best produces an in-memory set of keys, types and values. The _application_ does need to know what to do with those keys, what each key represents, how keys and values interact (e.g. a flag at one key that changes the interpretation of another), etc.

What CBOR and JSON really mean when they say they are schema-less is that they are punting the work of the schema out from the serialization layer and up into the application layer. Applications on both ends of a connection are left to figure out, for example, the difference between a field called `name` and a field called `canonical_name`, when to use one vs. the other, and how to interpret the values they contain or what it means when an encoder starts including the second or stops including the first.

Weakly-structured messages like CBOR, and JSON before it, are attractive at the start of a development cycle but can lead to challenges over time.

In contrast, Protocol Buffers have features that at first seem not perfectly but over time 

### What about zero-copy?

Many newer encoding formats offer zero-copy semantics, useful for very high speed communications involving many messages containing many fields. In the case of the web