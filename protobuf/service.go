// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf/prototag"
)

type ServiceDescriptorProto struct {
	desc    *descriptorpb.ServiceDescriptorProto
	comment *Comment

	seen            map[string]bool
	numMethod       int32
	methodLocations map[int32]*descriptorpb.SourceCodeInfo_Location
}

func NewServiceDescriptorProto(name string) *ServiceDescriptorProto {
	return &ServiceDescriptorProto{
		desc: &descriptorpb.ServiceDescriptorProto{
			Name: proto.String(name),
		},
		comment:         &Comment{},
		seen:            make(map[string]bool),
		methodLocations: make(map[int32]*descriptorpb.SourceCodeInfo_Location),
	}
}

func (sd *ServiceDescriptorProto) GetName() string {
	return sd.desc.GetName()
}

func (sd *ServiceDescriptorProto) AddMethod(method *descriptorpb.MethodDescriptorProto, description string) *ServiceDescriptorProto {
	if sd.seen[method.GetName()] {
		return sd
	}
	sd.seen[method.GetName()] = true

	sd.numMethod++
	comments := description
	if comments != "" {
		loc := &descriptorpb.SourceCodeInfo_Location{
			LeadingComments: proto.String(comments),
			Path:            []int32{prototag.ServiceMethods, sd.numMethod - 1},
		}
		sd.methodLocations[sd.numMethod-1] = loc
		sd.AddLeadingComment(method.GetName(), comments)
	}
	sd.desc.Method = append(sd.desc.Method, method)

	return sd
}

func (sd *ServiceDescriptorProto) SetServiceOptions(options *descriptorpb.ServiceOptions) *ServiceDescriptorProto {
	sd.desc.Options = options
	return sd
}

func (sd *ServiceDescriptorProto) AddLeadingComment(fn, leading string) *ServiceDescriptorProto {
	sd.comment.LeadingComments = conv.NormalizeComment(fn, leading)

	return sd
}

func (sd *ServiceDescriptorProto) AddTrailingComment(trailing string) *ServiceDescriptorProto {
	sd.comment.TrailingComments = trailing

	return sd
}

func (sd *ServiceDescriptorProto) AddLeadingDetachedComment(leadingDetached []string) *ServiceDescriptorProto {
	sd.comment.LeadingDetachedComments = leadingDetached

	return sd
}

func (sd *ServiceDescriptorProto) GetComment() *Comment {
	return sd.comment
}

func (sd *ServiceDescriptorProto) Build() *descriptorpb.ServiceDescriptorProto {
	return sd.desc
}
