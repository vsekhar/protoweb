# Protoweb

Protoweb is a thought experiment to evaluate a QUIC + Protobuf implementation of web standards.

## But why?

Protobufs have several advantages over custom string- and byte-based protocols like HTTP:

  * Versionable protocol: clients and servers implementing different versions of the protocol can still communicate
  * Machine readable specification: clients, servers and other tools can analyze the protocol as well as how it is evolving (for example, to check for backward compatibility)
  * Efficient: common values can be compactly encoded on the wire, with translations available in the protocol specification (kind of like a perfect Huffman code)

## Roadmap

  * Capture some real-ish traffic for replay (Puppeteer?)
  * Define protocol buffers to handle captured traffic
  * Implement counting transport for benchmarking
  * Benchmark traffic replays with HTTP, HTTP/2 and protoweb
  * Implement Go HTTP server and client using protoweb (with HTTP/2 and HTTP fallback)

## Prerequisites for development

The library can be imported as is via Go programs.

For development, testing and distribution packaging, some tools need to be installed.

Install `protoc`. For Debian/Ubuntu:

    $ sudo apt install protobuf-compiler

Install `prototool` (do this from outside the repo directory to avoid mangling go.mod):

    $ go get github.com/uber/prototool/cmd/prototool
    