// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	
	"go.uber.org/zap"
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type ApiSnapshot struct {
	Gateways GatewaysByNamespace
	VirtualServices VirtualServicesByNamespace
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Gateways: s.Gateways.Clone(),
		VirtualServices: s.VirtualServices.Clone(),
	}
}

func (s ApiSnapshot) snapshotToHash() ApiSnapshot {
	snapshotForHashing := s.Clone()
	for _, gateway := range snapshotForHashing.Gateways.List() {
		resources.UpdateMetadata(gateway, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		gateway.SetStatus(core.Status{})
	}
	for _, virtualService := range snapshotForHashing.VirtualServices.List() {
		resources.UpdateMetadata(virtualService, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		virtualService.SetStatus(core.Status{})
	}

	return snapshotForHashing
}

func (s ApiSnapshot) Hash() uint64 {
	return s.hashStruct(s.snapshotToHash())
 }

 func (s ApiSnapshot) HashFields() []zap.Field {
	snapshotForHashing := s.snapshotToHash()
	var fields []zap.Field
	gateways := s.hashStruct(snapshotForHashing.Gateways.List())
	fields = append(fields, zap.Uint64("gateways", gateways ))
	virtualServices := s.hashStruct(snapshotForHashing.VirtualServices.List())
	fields = append(fields, zap.Uint64("virtualServices", virtualServices ))

	return append(fields, zap.Uint64("snapshotHash",  s.hashStruct(snapshotForHashing)))
 }
 
func (s ApiSnapshot) hashStruct(v interface{}) uint64 {
	h, err := hashstructure.Hash(v, nil)
	 if err != nil {
		 panic(err)
	 }
	 return h
 }


