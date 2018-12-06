// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter.proto

package aws // import "github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// AWS Lambda contains the configuration necessary to perform transform regular http calls to
// AWS Lambda invocations.
type LambdaPerRoute struct {
	// The name of the function
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The qualifier of the function (defualts to $LATEST if not specified)
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Invocation type - async or regular.
	Async                bool     `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LambdaPerRoute) Reset()         { *m = LambdaPerRoute{} }
func (m *LambdaPerRoute) String() string { return proto.CompactTextString(m) }
func (*LambdaPerRoute) ProtoMessage()    {}
func (*LambdaPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_6f49f9dd19797cac, []int{0}
}
func (m *LambdaPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaPerRoute.Unmarshal(m, b)
}
func (m *LambdaPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaPerRoute.Marshal(b, m, deterministic)
}
func (dst *LambdaPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaPerRoute.Merge(dst, src)
}
func (m *LambdaPerRoute) XXX_Size() int {
	return xxx_messageInfo_LambdaPerRoute.Size(m)
}
func (m *LambdaPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_LambdaPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_LambdaPerRoute proto.InternalMessageInfo

func (m *LambdaPerRoute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LambdaPerRoute) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func (m *LambdaPerRoute) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

type LambdaProtocolExtension struct {
	// The host header for AWS this cluster
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The region for this cluster
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The access_key for AWS this cluster
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// The secret_key for AWS this cluster
	SecretKey            string   `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LambdaProtocolExtension) Reset()         { *m = LambdaProtocolExtension{} }
func (m *LambdaProtocolExtension) String() string { return proto.CompactTextString(m) }
func (*LambdaProtocolExtension) ProtoMessage()    {}
func (*LambdaProtocolExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_6f49f9dd19797cac, []int{1}
}
func (m *LambdaProtocolExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaProtocolExtension.Unmarshal(m, b)
}
func (m *LambdaProtocolExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaProtocolExtension.Marshal(b, m, deterministic)
}
func (dst *LambdaProtocolExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaProtocolExtension.Merge(dst, src)
}
func (m *LambdaProtocolExtension) XXX_Size() int {
	return xxx_messageInfo_LambdaProtocolExtension.Size(m)
}
func (m *LambdaProtocolExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_LambdaProtocolExtension.DiscardUnknown(m)
}

var xxx_messageInfo_LambdaProtocolExtension proto.InternalMessageInfo

func (m *LambdaProtocolExtension) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *LambdaProtocolExtension) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *LambdaProtocolExtension) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *LambdaProtocolExtension) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func init() {
	proto.RegisterType((*LambdaPerRoute)(nil), "envoy.config.filter.http.aws.v2.LambdaPerRoute")
	proto.RegisterType((*LambdaProtocolExtension)(nil), "envoy.config.filter.http.aws.v2.LambdaProtocolExtension")
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_filter_6f49f9dd19797cac) }

var fileDescriptor_filter_6f49f9dd19797cac = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x33, 0x31,
	0x18, 0x85, 0x99, 0xb6, 0x5f, 0xf9, 0x26, 0x88, 0x8b, 0x20, 0xb4, 0x88, 0x62, 0xed, 0xaa, 0x1b,
	0x13, 0xd0, 0xbd, 0x48, 0xc1, 0x55, 0x5d, 0xc8, 0x2c, 0xdd, 0x48, 0x9a, 0xbe, 0x4d, 0x63, 0x33,
	0x79, 0xc7, 0x24, 0x33, 0x75, 0xae, 0xca, 0xbd, 0x2b, 0x6f, 0xc7, 0xbb, 0x90, 0x99, 0x8c, 0x7f,
	0xe8, 0x2a, 0x27, 0xe7, 0x3c, 0xe1, 0x81, 0x90, 0xbd, 0xb5, 0x36, 0x01, 0x1c, 0x2b, 0x1c, 0x06,
	0xa4, 0x27, 0x60, 0x2b, 0xac, 0x99, 0x44, 0xbb, 0xd6, 0x8a, 0x75, 0xd3, 0x26, 0x84, 0x82, 0x89,
	0x9d, 0x67, 0xd5, 0xf9, 0xe1, 0xa8, 0x12, 0x46, 0xaf, 0x44, 0x00, 0xfe, 0x11, 0xe2, 0xcb, 0xa9,
	0x24, 0xfb, 0x37, 0x22, 0x5f, 0xae, 0xc4, 0x2d, 0xb8, 0x0c, 0xcb, 0x00, 0xf4, 0x98, 0x0c, 0xac,
	0xc8, 0x61, 0x9c, 0x4c, 0x92, 0x59, 0x3a, 0x4f, 0x5f, 0xde, 0x5e, 0xfb, 0x03, 0xd7, 0x9b, 0x24,
	0x59, 0x5b, 0xd3, 0x23, 0x92, 0x3e, 0x96, 0xc2, 0xe8, 0xb5, 0x06, 0x37, 0xee, 0x35, 0x4c, 0xf6,
	0x55, 0xd0, 0x03, 0xf2, 0x4f, 0xf8, 0xda, 0xca, 0x71, 0x7f, 0x92, 0xcc, 0xfe, 0x67, 0xf1, 0x32,
	0x7d, 0x4e, 0xc8, 0xa8, 0xb3, 0x34, 0x52, 0x89, 0xe6, 0xfa, 0x29, 0x80, 0xf5, 0x1a, 0x6d, 0xa3,
	0xdb, 0xa0, 0x0f, 0x7f, 0xe8, 0x9a, 0x9a, 0x9e, 0x92, 0xa1, 0x03, 0xa5, 0xd1, 0x46, 0xd7, 0x77,
	0xa0, 0x1b, 0xe8, 0x8c, 0x10, 0x21, 0x25, 0x78, 0x7f, 0xbf, 0x85, 0xba, 0x15, 0xff, 0xc0, 0xd2,
	0x38, 0x2e, 0xa0, 0x6e, 0x48, 0x0f, 0xd2, 0x41, 0x68, 0xc9, 0xc1, 0x2f, 0x32, 0x8e, 0x0b, 0xa8,
	0xe7, 0x57, 0x77, 0x97, 0x4a, 0x87, 0x4d, 0xb9, 0x64, 0x12, 0x73, 0xee, 0xd1, 0xe0, 0x99, 0xc6,
	0x78, 0x16, 0x0e, 0x1f, 0x40, 0x06, 0xcf, 0x3f, 0x83, 0x32, 0x88, 0xbc, 0xd8, 0x2a, 0x5e, 0x98,
	0x52, 0x69, 0xeb, 0xb9, 0xd8, 0xf9, 0xe5, 0xb0, 0xfd, 0xdf, 0x8b, 0xf7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x60, 0xf1, 0x89, 0x3d, 0xa9, 0x01, 0x00, 0x00,
}
