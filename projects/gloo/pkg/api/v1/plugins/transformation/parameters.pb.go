// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins/transformation/parameters.proto

package transformation // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Parameters struct {
	// headers that will be used to extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         headers:
	//           x-user-id: { userId }
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// part of the (or the entire) path that will be used extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         path: /users/{ userId }
	Path                 *types.StringValue `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Parameters) Reset()         { *m = Parameters{} }
func (m *Parameters) String() string { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()    {}
func (*Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_parameters_2a2cb281daf0f95f, []int{0}
}
func (m *Parameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameters.Unmarshal(m, b)
}
func (m *Parameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameters.Marshal(b, m, deterministic)
}
func (dst *Parameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameters.Merge(dst, src)
}
func (m *Parameters) XXX_Size() int {
	return xxx_messageInfo_Parameters.Size(m)
}
func (m *Parameters) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameters.DiscardUnknown(m)
}

var xxx_messageInfo_Parameters proto.InternalMessageInfo

func (m *Parameters) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Parameters) GetPath() *types.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "transformation.plugins.gloo.solo.io.Parameters")
	proto.RegisterMapType((map[string]string)(nil), "transformation.plugins.gloo.solo.io.Parameters.HeadersEntry")
}
func (this *Parameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Parameters)
	if !ok {
		that2, ok := that.(Parameters)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !this.Path.Equal(that1.Path) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("plugins/transformation/parameters.proto", fileDescriptor_parameters_2a2cb281daf0f95f)
}

var fileDescriptor_parameters_2a2cb281daf0f95f = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xc9, 0xe6, 0x1f, 0x96, 0x79, 0x90, 0xb2, 0xc3, 0x28, 0x32, 0x8a, 0x1e, 0xec, 0xc5,
	0x37, 0x3a, 0x2f, 0x32, 0x3c, 0x09, 0xa2, 0x17, 0x41, 0x2a, 0xec, 0xe0, 0x2d, 0x9d, 0x59, 0x1a,
	0xd7, 0xf6, 0x0d, 0x49, 0x3a, 0xd9, 0x37, 0xf2, 0xf3, 0xf8, 0x11, 0xfc, 0x24, 0x92, 0x74, 0x9b,
	0x0a, 0x1e, 0xbc, 0x3d, 0x6f, 0x79, 0x9f, 0xdf, 0xfb, 0xa3, 0xa1, 0xa7, 0xba, 0x6c, 0xa4, 0xaa,
	0x2d, 0x73, 0x86, 0xd7, 0x76, 0x8e, 0xa6, 0xe2, 0x4e, 0x61, 0xcd, 0x34, 0x37, 0xbc, 0x12, 0x4e,
	0x18, 0x0b, 0xda, 0xa0, 0xc3, 0xe8, 0xe4, 0xf7, 0x02, 0xac, 0x7b, 0x20, 0x4b, 0x44, 0xb0, 0x58,
	0x22, 0x28, 0x8c, 0x47, 0x12, 0x51, 0x96, 0x82, 0x85, 0x4a, 0xde, 0xcc, 0xd9, 0x9b, 0xe1, 0x5a,
	0x6f, 0x21, 0xf1, 0x40, 0xa2, 0xc4, 0x10, 0x99, 0x4f, 0xed, 0xd7, 0xe3, 0x0f, 0x42, 0xe9, 0xe3,
	0xf6, 0x5e, 0x34, 0xa5, 0xfb, 0x85, 0xe0, 0x2f, 0xc2, 0xd8, 0x21, 0x49, 0xba, 0x69, 0x7f, 0x7c,
	0x0d, 0xff, 0xb8, 0x0d, 0xdf, 0x04, 0xb8, 0x6f, 0xeb, 0xb7, 0xb5, 0x33, 0xab, 0x6c, 0x03, 0x8b,
	0xce, 0xe9, 0x8e, 0xe6, 0xae, 0x18, 0x76, 0x12, 0x92, 0xf6, 0xc7, 0x47, 0xd0, 0xba, 0xc2, 0xc6,
	0x15, 0x9e, 0x9c, 0x51, 0xb5, 0x9c, 0xf2, 0xb2, 0x11, 0x59, 0xd8, 0x8c, 0x27, 0xf4, 0xe0, 0x27,
	0x2a, 0x3a, 0xa4, 0xdd, 0x85, 0x58, 0x0d, 0x49, 0x42, 0xd2, 0x5e, 0xe6, 0x63, 0x34, 0xa0, 0xbb,
	0x4b, 0x5f, 0x08, 0xd0, 0x5e, 0xd6, 0x0e, 0x93, 0xce, 0x15, 0xb9, 0x79, 0x78, 0xff, 0x1c, 0x91,
	0xe7, 0x3b, 0xa9, 0x5c, 0xd1, 0xe4, 0x30, 0xc3, 0x8a, 0x79, 0xd1, 0x33, 0x85, 0xcc, 0x5b, 0xfb,
	0x3f, 0xf4, 0x2a, 0x66, 0xce, 0xae, 0xa7, 0x85, 0x64, 0x5c, 0x2b, 0xb6, 0xbc, 0x60, 0x7f, 0xbf,
	0x47, 0xbe, 0x17, 0x34, 0x2f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x25, 0xa3, 0x6a, 0xb0,
	0x01, 0x00, 0x00,
}
