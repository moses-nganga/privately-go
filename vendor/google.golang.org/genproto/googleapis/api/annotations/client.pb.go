// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/client.proto

package annotations

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_MethodSignature = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         1051,
	Name:          "google.api.method_signature",
	Tag:           "bytes,1051,rep,name=method_signature,json=methodSignature",
	Filename:      "google/api/client.proto",
}

var E_DefaultHost = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1049,
	Name:          "google.api.default_host",
	Tag:           "bytes,1049,opt,name=default_host,json=defaultHost",
	Filename:      "google/api/client.proto",
}

var E_OauthScopes = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1050,
	Name:          "google.api.oauth_scopes",
	Tag:           "bytes,1050,opt,name=oauth_scopes,json=oauthScopes",
	Filename:      "google/api/client.proto",
}

func init() {
	proto.RegisterExtension(E_MethodSignature)
	proto.RegisterExtension(E_DefaultHost)
	proto.RegisterExtension(E_OauthScopes)
}

func init() { proto.RegisterFile("google/api/client.proto", fileDescriptor_client_1608614df476619f) }

var fileDescriptor_client_1608614df476619f = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x55, 0x40, 0xa8, 0x75, 0x11, 0xa0, 0x2c, 0x20, 0x06, 0xc8, 0xd8, 0xc9, 0x1e, 0xd8,
	0xca, 0xd4, 0x76, 0xe0, 0x8f, 0x84, 0x88, 0x9a, 0x8d, 0x25, 0x72, 0x9d, 0xab, 0x63, 0x29, 0xf5,
	0x59, 0xf6, 0x85, 0xef, 0x02, 0x6c, 0x7c, 0x52, 0x54, 0xc7, 0x11, 0x48, 0x0c, 0x6c, 0x27, 0xbd,
	0xf7, 0xfb, 0x9d, 0xf4, 0xd8, 0x85, 0x46, 0xd4, 0x2d, 0x08, 0xe9, 0x8c, 0x50, 0xad, 0x01, 0x4b,
	0xdc, 0x79, 0x24, 0xcc, 0x58, 0x1f, 0x70, 0xe9, 0xcc, 0x55, 0x9e, 0x4a, 0x31, 0xd9, 0x74, 0x5b,
	0x51, 0x43, 0x50, 0xde, 0x38, 0x42, 0xdf, 0xb7, 0xe7, 0x4f, 0xec, 0x7c, 0x07, 0xd4, 0x60, 0x5d,
	0x05, 0xa3, 0xad, 0xa4, 0xce, 0x43, 0x76, 0xcd, 0x93, 0x62, 0xc0, 0xf8, 0x73, 0xac, 0xbc, 0x38,
	0x32, 0x68, 0xc3, 0xe5, 0xe7, 0x38, 0x3f, 0x9c, 0x4d, 0xd6, 0x67, 0x3d, 0x58, 0x0e, 0xdc, 0x7c,
	0xc5, 0x4e, 0x6a, 0xd8, 0xca, 0xae, 0xa5, 0xaa, 0xc1, 0x40, 0xd9, 0xcd, 0x1f, 0x4f, 0x09, 0xfe,
	0xcd, 0x28, 0x18, 0x44, 0xef, 0xe3, 0x7c, 0x34, 0x9b, 0xac, 0xa7, 0x89, 0x7a, 0xc0, 0x40, 0x7b,
	0x09, 0xca, 0x8e, 0x9a, 0x2a, 0x28, 0x74, 0x10, 0xfe, 0x97, 0x7c, 0x24, 0x49, 0xa4, 0xca, 0x08,
	0x2d, 0x0d, 0x3b, 0x55, 0xb8, 0xe3, 0x3f, 0x4b, 0x2c, 0xa7, 0xab, 0xb8, 0x51, 0xb1, 0x97, 0x14,
	0xa3, 0xd7, 0x45, 0x8a, 0x34, 0xb6, 0xd2, 0x6a, 0x8e, 0x5e, 0x0b, 0x0d, 0x36, 0xbe, 0x10, 0x7d,
	0x24, 0x9d, 0x09, 0x71, 0x5c, 0x69, 0x2d, 0x92, 0x8c, 0xbf, 0xee, 0x7e, 0xdd, 0x5f, 0x07, 0x47,
	0xf7, 0x8b, 0xe2, 0x71, 0x73, 0x1c, 0xa1, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xc2,
	0xcf, 0x71, 0x90, 0x01, 0x00, 0x00,
}