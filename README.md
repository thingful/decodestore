# encrypted datastore

Simple placeholder implementation of the sketched out datastore implementation.

Implemented using a library called Twirp, which is very like simplified GRPC.
It starts from a protobuf definition like GRPC, but is simpler, and runs over
HTTP/1.1 rather than having a hard requirement on HTTP/2.

## Requirements

* Go 1.8+
* retool
* protoc 3.5.0+

## Using retool

Experimenting here with using retool to handle vendored build time binaries
(like dep and the protobuf stuff). We add tools using the `retool` binary which
adds a `_tools` folder which I've also vendored. After that the required
approach is that for tools within that dir we run them using `retool do
<command>` which ensures everyone is using exactly the same versions.

## Make tasks

Currently v.minimal - I just have a `gen` task which uses the retooled `protoc`
binary to rebuild the interface implementations from the protobuf file.
