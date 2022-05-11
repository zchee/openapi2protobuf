// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ServiceDescriptorProto struct {
	desc *descriptorpb.ServiceDescriptorProto
}

func NewServiceDescriptorProto(name string) *ServiceDescriptorProto {
	return &ServiceDescriptorProto{
		desc: &descriptorpb.ServiceDescriptorProto{
			Name: proto.String(name),
		},
	}
}

func (sd *ServiceDescriptorProto) SetServiceOptions(options *descriptorpb.ServiceOptions) *ServiceDescriptorProto {
	sd.desc.Options = options
	return sd
}

func (sd *ServiceDescriptorProto) Build() *descriptorpb.ServiceDescriptorProto {
	return sd.desc
}
