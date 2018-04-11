// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/datastore/service.proto

/*
Package datastore is a generated protocol buffer package.

It is generated from these files:
	pkg/datastore/service.proto

It has these top-level messages:
	WriteRequest
	WriteResponse
	Timestamp
	ReadRequest
	ReadResponse
*/
package datastore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// WriteRequest is the message we send to the store in order to write encrypted
// data for particular user.
type WriteRequest struct {
	// The public key for a consumer for which the included data has been
	// encrypted.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// The encrypted data comprising the event to be stored. This data is
	// completely opaque to the data store, so is represented as just an
	// arbitrary sequence of bytes.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WriteRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *WriteRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// WriteResponse is a placeholder response message returned from the call to
// the service. Currently it is empty as we have not identified any fields that
// should be returned, however we keep as a type under our control such that
// should such a need be identified in the future, this type can be extended to
// include those fields.
type WriteResponse struct {
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Timestamp represents a point in time independent of any time zone or
// calendar, represented as seconds and fractions of seconds at nanosecond
// resolution in UTC Epoch time.
type Timestamp struct {
	// Contains the UTC Epoch time in seconds expressed as a 64 bit integer.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Contains the fractional part of UTC Epoch time in nanoseconds expressed as
	// a 64 bit integer.
	Nanoseconds int64 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanoseconds() int64 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

// ReadRequest is a message sent by a consumer to the service by which
// encrypted data can be requested for a specific user, identified by their
// public key.
type ReadRequest struct {
	// The public key for a consumer for which encrypted data has been stored.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// The start of an interval for which data is being requested.
	StartTime *Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The end of an interval for which data is being requested, represented as
	// Timestamp message. This field is optional and if omitted defaults to 'now'
	EndTime *Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// A pagination cursor returned from the call to ReadData that indicates from
	// where data access should continue. This field is optional, if not present
	// it means start from the beginning of the dataset.
	PageCursor string `protobuf:"bytes,4,opt,name=page_cursor,json=pageCursor" json:"page_cursor,omitempty"`
	// Integer value containing the maximum desired number of results within each
	// page of data.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *ReadRequest) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ReadRequest) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ReadRequest) GetPageCursor() string {
	if m != nil {
		return m.PageCursor
	}
	return ""
}

func (m *ReadRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// ReadResponse is the message returned from our service in response to a
// ReadRequest. Comprises the public key of the user for which the data was
// requested, a repeated list of EncryptedEvent messages comprising the data,
// but also a cursor by which the client can paginate through a large dataset.
// Note we do not attempt to use any streaming functionalities here to simplify
// things, rather we will just allow clients to paginate through datasets to
// obtain the data they need.
type ReadResponse struct {
	// The public of the consumer for which this response struct contains data.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// A string containing a cursor pointing at the next page of encrypted
	// events. Can be used by a client when constructing the next ReadRequest to
	// allow the client to easily consume all results within a time window. If
	// this value is an empty string, then no further pages are available to be
	// requested.
	NextPageCursor string `protobuf:"bytes,2,opt,name=next_page_cursor,json=nextPageCursor" json:"next_page_cursor,omitempty"`
	// A list of EncryptedEvent messages comprising a single page of results
	// within a larger requested dataset.
	Events []*ReadResponse_EncryptedEvent `protobuf:"bytes,3,rep,name=events" json:"events,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *ReadResponse) GetNextPageCursor() string {
	if m != nil {
		return m.NextPageCursor
	}
	return ""
}

func (m *ReadResponse) GetEvents() []*ReadResponse_EncryptedEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

// EncryptedEvent contains a single encrypted event message stored for a
// client. This comprises a timestamp which is the instance the data was
// received and written to storage, along with the encrypted list of bytes
// comprising the message. This message is returned
type ReadResponse_EncryptedEvent struct {
	// The time at which this event was recorded expressed in UTC Epoch time.
	Timestamp *Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// The list of bytes comprising a chunk of encrypted data for a user.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ReadResponse_EncryptedEvent) Reset()                    { *m = ReadResponse_EncryptedEvent{} }
func (m *ReadResponse_EncryptedEvent) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse_EncryptedEvent) ProtoMessage()               {}
func (*ReadResponse_EncryptedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *ReadResponse_EncryptedEvent) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ReadResponse_EncryptedEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteRequest)(nil), "decode.thingful.datastore.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "decode.thingful.datastore.WriteResponse")
	proto.RegisterType((*Timestamp)(nil), "decode.thingful.datastore.Timestamp")
	proto.RegisterType((*ReadRequest)(nil), "decode.thingful.datastore.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "decode.thingful.datastore.ReadResponse")
	proto.RegisterType((*ReadResponse_EncryptedEvent)(nil), "decode.thingful.datastore.ReadResponse.EncryptedEvent")
}

func init() { proto.RegisterFile("pkg/datastore/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0x9b, 0x7e, 0xec, 0xdc, 0xac, 0x55, 0xe6, 0x29, 0xb6, 0x88, 0x21, 0x88, 0xcd, 0x53,
	0x0a, 0x2b, 0xf8, 0x2a, 0xb6, 0x16, 0x1f, 0x04, 0x91, 0xa9, 0x20, 0x28, 0x18, 0xa6, 0x99, 0xeb,
	0x76, 0x68, 0x77, 0x26, 0xce, 0x4c, 0x8a, 0xdb, 0x7f, 0xe1, 0x0f, 0xf3, 0xbf, 0xf8, 0x13, 0x64,
	0x26, 0xbb, 0x31, 0x82, 0xee, 0xae, 0x6f, 0x99, 0x73, 0xef, 0x39, 0xf7, 0xdc, 0xc3, 0x0d, 0x1c,
	0x35, 0xd7, 0xb3, 0x13, 0xc1, 0x1d, 0xb7, 0x4e, 0x1b, 0x3c, 0xb1, 0x68, 0x6e, 0x65, 0x8d, 0x65,
	0x63, 0xb4, 0xd3, 0xf4, 0xa1, 0xc0, 0x5a, 0x0b, 0x2c, 0xdd, 0x95, 0x54, 0xb3, 0x2f, 0xed, 0x4d,
	0xd9, 0x37, 0xe6, 0x2f, 0x61, 0xf2, 0xc1, 0x48, 0x87, 0x0c, 0xbf, 0xb6, 0x68, 0x1d, 0x7d, 0x04,
	0xd0, 0xb4, 0x97, 0x37, 0xb2, 0xae, 0xae, 0x71, 0x91, 0x46, 0x59, 0x54, 0x10, 0x46, 0x3a, 0xe4,
	0x0d, 0x2e, 0x28, 0x85, 0x1d, 0xcf, 0x4d, 0x47, 0x59, 0x54, 0x4c, 0x58, 0xf8, 0xce, 0xef, 0xc3,
	0xbd, 0xa5, 0x84, 0x6d, 0xb4, 0xb2, 0x98, 0xbf, 0x06, 0xf2, 0x5e, 0xce, 0xd1, 0x3a, 0x3e, 0x6f,
	0x68, 0x0a, 0xfb, 0x16, 0x6b, 0xad, 0x84, 0x0d, 0x6a, 0x31, 0x5b, 0x3d, 0x69, 0x06, 0x89, 0xe2,
	0x4a, 0xaf, 0xaa, 0xa3, 0x50, 0x1d, 0x42, 0xf9, 0xcf, 0x08, 0x12, 0x86, 0x5c, 0x6c, 0x69, 0xee,
	0x0c, 0xc0, 0x3a, 0x6e, 0x5c, 0xe5, 0xe4, 0x1c, 0x83, 0x5e, 0x32, 0x7d, 0x52, 0xfe, 0x73, 0xf7,
	0xb2, 0x37, 0xc9, 0x48, 0xe0, 0xf9, 0x37, 0x7d, 0x01, 0x63, 0x54, 0xa2, 0x93, 0x88, 0xff, 0x43,
	0x62, 0x1f, 0x95, 0x08, 0x02, 0x8f, 0x21, 0x69, 0xf8, 0x0c, 0xab, 0xba, 0x35, 0x56, 0x9b, 0x74,
	0x27, 0xb8, 0x04, 0x0f, 0x9d, 0x05, 0x84, 0x1e, 0x01, 0x09, 0x0d, 0x56, 0xde, 0x61, 0xba, 0x9b,
	0x45, 0xc5, 0x2e, 0x1b, 0x7b, 0xe0, 0x42, 0xde, 0x61, 0xfe, 0x7d, 0x04, 0x93, 0x6e, 0xe5, 0x2e,
	0xcc, 0x4d, 0x3b, 0x17, 0xf0, 0x40, 0xe1, 0x37, 0x57, 0x0d, 0x47, 0x8e, 0x42, 0xd3, 0x81, 0xc7,
	0xdf, 0xfd, 0x1e, 0xfb, 0x16, 0xf6, 0xf0, 0x16, 0x95, 0xb3, 0x69, 0x9c, 0xc5, 0x45, 0x32, 0x7d,
	0xbe, 0x66, 0xad, 0xa1, 0x83, 0xf2, 0x5c, 0xd5, 0x66, 0xd1, 0x38, 0x14, 0xe7, 0x9e, 0xce, 0x96,
	0x2a, 0x87, 0x57, 0x70, 0xf0, 0x67, 0x85, 0x9e, 0x02, 0x71, 0xab, 0x3c, 0x82, 0xd3, 0xad, 0xe3,
	0xef, 0x69, 0x7f, 0x3b, 0xb0, 0xe9, 0x8f, 0x68, 0x30, 0xea, 0xc2, 0x73, 0xe9, 0x67, 0x20, 0xe1,
	0xe6, 0x5e, 0x71, 0xc7, 0xe9, 0xf1, 0x9a, 0x21, 0xc3, 0xe3, 0x3e, 0x2c, 0x36, 0x37, 0x2e, 0x53,
	0xff, 0x04, 0x63, 0x9f, 0x41, 0x90, 0x7f, 0xba, 0x31, 0xa8, 0x4e, 0xfd, 0x78, 0xcb, 0x40, 0x4f,
	0x93, 0x8f, 0xa4, 0xaf, 0x5c, 0xee, 0x85, 0x5f, 0xf4, 0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe1, 0xc1, 0x34, 0x22, 0xc1, 0x03, 0x00, 0x00,
}
