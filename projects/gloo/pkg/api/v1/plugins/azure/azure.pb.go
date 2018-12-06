// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins/azure/azure.proto

package azure // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/azure"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

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

type UpstreamSpec_FunctionSpec_AuthLevel int32

const (
	UpstreamSpec_FunctionSpec_Anonymous UpstreamSpec_FunctionSpec_AuthLevel = 0
	UpstreamSpec_FunctionSpec_Function  UpstreamSpec_FunctionSpec_AuthLevel = 1
	UpstreamSpec_FunctionSpec_Admin     UpstreamSpec_FunctionSpec_AuthLevel = 2
)

var UpstreamSpec_FunctionSpec_AuthLevel_name = map[int32]string{
	0: "Anonymous",
	1: "Function",
	2: "Admin",
}
var UpstreamSpec_FunctionSpec_AuthLevel_value = map[string]int32{
	"Anonymous": 0,
	"Function":  1,
	"Admin":     2,
}

func (x UpstreamSpec_FunctionSpec_AuthLevel) String() string {
	return proto.EnumName(UpstreamSpec_FunctionSpec_AuthLevel_name, int32(x))
}
func (UpstreamSpec_FunctionSpec_AuthLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_azure_50cd292bf6e2386b, []int{0, 0, 0}
}

// Upstream Spec for Azure Functions Upstreams
// Azure Upstreams represent a collection of Azure Functions for a particular Azure Account
// within a particular Function App
type UpstreamSpec struct {
	// The Name of the Azure Function App where the functions are grouped
	FunctionAppName string `protobuf:"bytes,1,opt,name=function_app_name,json=functionAppName,proto3" json:"function_app_name,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an [Azure Publish Profile JSON file](https://azure.microsoft.com/en-us/downloads/publishing-profile-overview/).
	// {{ hide_not_implemented "Azure Secrets can be created with `glooctl secret create azure ...`" }}
	// Note that this secret is not required unless Function Discovery is enabled
	SecretRef            core.ResourceRef             `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef" json:"secret_ref"`
	Functions            []*UpstreamSpec_FunctionSpec `protobuf:"bytes,3,rep,name=functions" json:"functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_azure_50cd292bf6e2386b, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (dst *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(dst, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetFunctionAppName() string {
	if m != nil {
		return m.FunctionAppName
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return core.ResourceRef{}
}

func (m *UpstreamSpec) GetFunctions() []*UpstreamSpec_FunctionSpec {
	if m != nil {
		return m.Functions
	}
	return nil
}

// Function Spec for Functions on Azure Functions Upstreams
// The Function Spec contains data necessary for Gloo to invoke Azure functions
type UpstreamSpec_FunctionSpec struct {
	// The Name of the Azure Function as it appears in the Azure Functions Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// Auth Level can bve either "anonymous" "function" or "admin"
	// See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details
	AuthLevel            UpstreamSpec_FunctionSpec_AuthLevel `protobuf:"varint,2,opt,name=auth_level,json=authLevel,proto3,enum=azure.plugins.gloo.solo.io.UpstreamSpec_FunctionSpec_AuthLevel" json:"auth_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *UpstreamSpec_FunctionSpec) Reset()         { *m = UpstreamSpec_FunctionSpec{} }
func (m *UpstreamSpec_FunctionSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec_FunctionSpec) ProtoMessage()    {}
func (*UpstreamSpec_FunctionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_azure_50cd292bf6e2386b, []int{0, 0}
}
func (m *UpstreamSpec_FunctionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec_FunctionSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec_FunctionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec_FunctionSpec.Marshal(b, m, deterministic)
}
func (dst *UpstreamSpec_FunctionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec_FunctionSpec.Merge(dst, src)
}
func (m *UpstreamSpec_FunctionSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec_FunctionSpec.Size(m)
}
func (m *UpstreamSpec_FunctionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec_FunctionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec_FunctionSpec proto.InternalMessageInfo

func (m *UpstreamSpec_FunctionSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *UpstreamSpec_FunctionSpec) GetAuthLevel() UpstreamSpec_FunctionSpec_AuthLevel {
	if m != nil {
		return m.AuthLevel
	}
	return UpstreamSpec_FunctionSpec_Anonymous
}

type DestinationSpec struct {
	// The Fucntion Name of the FunctionSpec to be invoked.
	FunctionName         string   `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_azure_50cd292bf6e2386b, []int{1}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (dst *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(dst, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "azure.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*UpstreamSpec_FunctionSpec)(nil), "azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec")
	proto.RegisterType((*DestinationSpec)(nil), "azure.plugins.gloo.solo.io.DestinationSpec")
	proto.RegisterEnum("azure.plugins.gloo.solo.io.UpstreamSpec_FunctionSpec_AuthLevel", UpstreamSpec_FunctionSpec_AuthLevel_name, UpstreamSpec_FunctionSpec_AuthLevel_value)
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if this.FunctionAppName != that1.FunctionAppName {
		return false
	}
	if !this.SecretRef.Equal(&that1.SecretRef) {
		return false
	}
	if len(this.Functions) != len(that1.Functions) {
		return false
	}
	for i := range this.Functions {
		if !this.Functions[i].Equal(that1.Functions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSpec_FunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_FunctionSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec_FunctionSpec)
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.AuthLevel != that1.AuthLevel {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("plugins/azure/azure.proto", fileDescriptor_azure_50cd292bf6e2386b) }

var fileDescriptor_azure_50cd292bf6e2386b = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x6f, 0x2f, 0x6a, 0xec, 0xdc, 0x22, 0x38, 0x71, 0x01, 0x5d, 0x28, 0xc1, 0x0d, 0x31,
	0x3a, 0x8d, 0x10, 0x5d, 0x62, 0x4a, 0x88, 0x2b, 0xe3, 0xa2, 0xc4, 0x8d, 0x0b, 0x9b, 0xa1, 0x9e,
	0x96, 0x91, 0x76, 0xce, 0x64, 0xfe, 0x90, 0xe8, 0xa3, 0xf8, 0x04, 0x3e, 0x0a, 0x4f, 0xe1, 0xc2,
	0x27, 0x31, 0x6d, 0x29, 0x62, 0xa2, 0x0b, 0xdd, 0xb4, 0x73, 0xce, 0xf9, 0xce, 0x97, 0xf3, 0x4b,
	0x3e, 0x32, 0x56, 0xa5, 0x2b, 0x84, 0x34, 0x11, 0xff, 0xe2, 0x34, 0xb4, 0x5f, 0xa6, 0x34, 0x5a,
	0xa4, 0xe1, 0xa9, 0x68, 0x05, 0xac, 0x28, 0x11, 0x99, 0xc1, 0x12, 0x99, 0xc0, 0xf0, 0x41, 0x81,
	0x05, 0x36, 0xb2, 0xa8, 0x7e, 0xb5, 0x1b, 0xe1, 0xd3, 0x42, 0xd8, 0x9d, 0xdb, 0xb2, 0x0c, 0xab,
	0xa8, 0x56, 0x3e, 0x13, 0xd8, 0xfe, 0xf7, 0xc2, 0x46, 0x5c, 0x89, 0xe8, 0xf0, 0x3c, 0xd2, 0x90,
	0xb7, 0xea, 0xe9, 0xd7, 0x1e, 0x09, 0xde, 0x29, 0x63, 0x35, 0xf0, 0x6a, 0xa3, 0x20, 0xa3, 0x4f,
	0xc8, 0xfd, 0xdc, 0xc9, 0xcc, 0x0a, 0x94, 0x29, 0x57, 0x2a, 0x95, 0xbc, 0x82, 0x91, 0x37, 0xf1,
	0x66, 0x7e, 0x32, 0xe8, 0x06, 0xb1, 0x52, 0x6f, 0x79, 0x05, 0x74, 0x49, 0x88, 0x81, 0x4c, 0x83,
	0x4d, 0x35, 0xe4, 0xa3, 0xeb, 0x89, 0x37, 0xbb, 0x99, 0x8f, 0x59, 0x86, 0x1a, 0xba, 0x1b, 0x59,
	0x02, 0x06, 0x9d, 0xce, 0x20, 0x81, 0x7c, 0x75, 0xeb, 0xf8, 0xfd, 0xd1, 0x55, 0xe2, 0xb7, 0x2b,
	0x09, 0xe4, 0x74, 0x43, 0xfc, 0xce, 0xd2, 0x8c, 0x7a, 0x93, 0xde, 0xec, 0x66, 0xfe, 0x82, 0xfd,
	0x1d, 0x98, 0x5d, 0x1e, 0xca, 0x5e, 0x9f, 0x36, 0xeb, 0x22, 0xf9, 0xe5, 0x13, 0x1e, 0x3d, 0x12,
	0x5c, 0xce, 0xe8, 0x63, 0xd2, 0x3f, 0x13, 0x5d, 0xd0, 0x04, 0x5d, 0xb3, 0x41, 0xf9, 0x40, 0x08,
	0x77, 0x76, 0x97, 0x96, 0x70, 0x80, 0xb2, 0x41, 0xb9, 0x37, 0x7f, 0xf5, 0x5f, 0xb7, 0xb0, 0xd8,
	0xd9, 0xdd, 0x9b, 0xda, 0x26, 0xf1, 0x79, 0xf7, 0x9c, 0x2e, 0x88, 0x7f, 0xee, 0xd3, 0x3e, 0xf1,
	0x63, 0x89, 0xf2, 0x73, 0x85, 0xce, 0x0c, 0xaf, 0x68, 0x40, 0xee, 0x76, 0x06, 0x43, 0x8f, 0xfa,
	0xe4, 0x76, 0xfc, 0xb1, 0x12, 0x72, 0x78, 0x3d, 0x7d, 0x49, 0x06, 0x6b, 0x30, 0x56, 0x48, 0xfe,
	0x4f, 0x30, 0xab, 0xf5, 0xb7, 0x1f, 0x0f, 0xbd, 0xf7, 0xcb, 0x3f, 0x04, 0xa1, 0x46, 0x88, 0x94,
	0xc6, 0x4f, 0x90, 0x59, 0x73, 0xaa, 0xf6, 0x45, 0x17, 0x8b, 0xdf, 0x62, 0xb8, 0xbd, 0xd3, 0x24,
	0x64, 0xf1, 0x33, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x90, 0x0b, 0x71, 0x9e, 0x02, 0x00, 0x00,
}
