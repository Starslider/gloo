// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins/consul/consul.proto

package consul // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/consul"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"

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

// Upstream Spec for Consul Upstreams
// consul Upstreams represent a set of one or more addressable pods for a consul Service
// the Gloo consul Upstream maps to a single service port. Because consul Services support mulitple ports,
// Gloo requires that a different upstream be created for each port
// consul Upstreams are typically generated automatically by Gloo from the consul API
type UpstreamSpec struct {
	// The name of the Consul Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The list of service tags Gloo should search for on a service instance
	// before deciding whether or not to include the instance as part of this
	// upstream
	ServiceTags []string `protobuf:"bytes,2,rep,name=service_tags,json=serviceTags" json:"service_tags,omitempty"`
	//     An optional Service Spec describing the service listening at this address
	ServiceSpec *plugins.ServiceSpec `protobuf:"bytes,3,opt,name=service_spec,json=serviceSpec" json:"service_spec,omitempty"`
	// is this consul service connect enabled.
	ConnectEnabled       bool     `protobuf:"varint,4,opt,name=connect_enabled,json=connectEnabled,proto3" json:"connect_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_consul_6015bcfcbfd55d36, []int{0}
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

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceTags() []string {
	if m != nil {
		return m.ServiceTags
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *plugins.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func (m *UpstreamSpec) GetConnectEnabled() bool {
	if m != nil {
		return m.ConnectEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "consul.plugins.gloo.solo.io.UpstreamSpec")
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
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if len(this.ServiceTags) != len(that1.ServiceTags) {
		return false
	}
	for i := range this.ServiceTags {
		if this.ServiceTags[i] != that1.ServiceTags[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if this.ConnectEnabled != that1.ConnectEnabled {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("plugins/consul/consul.proto", fileDescriptor_consul_6015bcfcbfd55d36) }

var fileDescriptor_consul_6015bcfcbfd55d36 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0x2b, 0xa2, 0xdd, 0x45, 0xa1, 0x78, 0x28, 0xbb, 0x20, 0x5d, 0x2f, 0xf6, 0x62,
	0x82, 0xfa, 0x00, 0x82, 0xb8, 0x78, 0xf3, 0xb0, 0xab, 0x17, 0x2f, 0x4b, 0x1a, 0x87, 0x18, 0x6d,
	0x33, 0xa1, 0x93, 0xee, 0x33, 0xf9, 0x1e, 0xbe, 0x89, 0x4f, 0x22, 0x69, 0x5a, 0x5d, 0xc4, 0x83,
	0xa7, 0x64, 0x7e, 0xbe, 0x99, 0x3f, 0xf9, 0x27, 0x99, 0xb9, 0xaa, 0xd5, 0xc6, 0x92, 0x50, 0x68,
	0xa9, 0xad, 0xfa, 0x83, 0xbb, 0x06, 0x3d, 0xa6, 0xb3, 0xa1, 0x8a, 0x0c, 0xd7, 0x15, 0x22, 0x27,
	0xac, 0x90, 0x1b, 0x9c, 0x1e, 0x6b, 0xd4, 0xd8, 0x71, 0x22, 0xdc, 0x62, 0xcb, 0xf4, 0x4e, 0x1b,
	0xff, 0xd2, 0x96, 0x5c, 0x61, 0x2d, 0x02, 0x79, 0x6e, 0x50, 0x84, 0x36, 0xe1, 0x1a, 0x7c, 0x05,
	0xe5, 0x29, 0x56, 0xd2, 0x19, 0xb1, 0xb9, 0x10, 0x83, 0x39, 0x41, 0xb3, 0x31, 0x0a, 0xd6, 0xe4,
	0x40, 0xc5, 0x41, 0xa7, 0x1f, 0x2c, 0x99, 0x3c, 0x3a, 0xf2, 0x0d, 0xc8, 0x7a, 0xe5, 0x40, 0xa5,
	0xf3, 0x64, 0x32, 0x60, 0x56, 0xd6, 0x90, 0xb1, 0x9c, 0x15, 0x07, 0xcb, 0x71, 0xaf, 0xdd, 0xcb,
	0x1a, 0xb6, 0x11, 0x2f, 0x35, 0x65, 0x3b, 0xf9, 0x68, 0x0b, 0x79, 0x90, 0x9a, 0xd2, 0xdb, 0x1f,
	0x24, 0x98, 0x65, 0xa3, 0x9c, 0x15, 0xe3, 0xcb, 0xf9, 0x9f, 0x5f, 0xe4, 0xab, 0x48, 0x06, 0xfb,
	0xef, 0x29, 0xdd, 0x5b, 0xce, 0x92, 0x23, 0x85, 0xd6, 0x82, 0xf2, 0x6b, 0xb0, 0xb2, 0xac, 0xe0,
	0x39, 0xdb, 0xcd, 0x59, 0xb1, 0xbf, 0x3c, 0xec, 0xe5, 0x45, 0x54, 0x6f, 0x16, 0xef, 0x9f, 0x27,
	0xec, 0xe9, 0xfa, 0x7f, 0xa1, 0xb8, 0x37, 0xfd, 0x3b, 0x98, 0xb8, 0x80, 0x72, 0xaf, 0xcb, 0xe4,
	0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x67, 0xb4, 0xcb, 0xae, 0x01, 0x00, 0x00,
}
