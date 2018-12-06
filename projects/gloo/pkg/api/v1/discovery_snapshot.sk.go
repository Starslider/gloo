// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Secrets   SecretsByNamespace
	Upstreams UpstreamsByNamespace
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s DiscoverySnapshot) snapshotToHash() DiscoverySnapshot {
	snapshotForHashing := s.Clone()
	for _, secret := range snapshotForHashing.Secrets.List() {
		resources.UpdateMetadata(secret, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}
	for _, upstream := range snapshotForHashing.Upstreams.List() {
		resources.UpdateMetadata(upstream, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		upstream.SetStatus(core.Status{})
	}

	return snapshotForHashing
}

func (s DiscoverySnapshot) Hash() uint64 {
	return s.hashStruct(s.snapshotToHash())
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	snapshotForHashing := s.snapshotToHash()
	var fields []zap.Field
	secrets := s.hashStruct(snapshotForHashing.Secrets.List())
	fields = append(fields, zap.Uint64("secrets", secrets))
	upstreams := s.hashStruct(snapshotForHashing.Upstreams.List())
	fields = append(fields, zap.Uint64("upstreams", upstreams))

	return append(fields, zap.Uint64("snapshotHash", s.hashStruct(snapshotForHashing)))
}

func (s DiscoverySnapshot) hashStruct(v interface{}) uint64 {
	h, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}
	return h
}
