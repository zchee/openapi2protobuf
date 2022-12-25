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

func (sd *ServiceDescriptorProto) AddMethod(method *MethodDescriptorProto) *ServiceDescriptorProto {
	if sd.seen[method.GetName()] {
		return sd
	}
	sd.seen[method.GetName()] = true

	sd.numMethod++
	comments := method.GetComment()
	if comments != nil {
		loc := &descriptorpb.SourceCodeInfo_Location{
			LeadingComments:         proto.String(comments.LeadingComments),
			TrailingComments:        proto.String(comments.TrailingComments),
			LeadingDetachedComments: comments.LeadingDetachedComments,
			Path:                    []int32{prototag.MessageFields, sd.numMethod - 1},
		}
		sd.methodLocations[sd.numMethod-1] = loc
	}
	sd.desc.Method = append(sd.desc.Method, method.Build())

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

type MethodDescriptorProto struct {
	desc    *descriptorpb.MethodDescriptorProto
	comment *Comment
}

func NewMethodDescriptorProto(name, input, output string) *MethodDescriptorProto {
	return &MethodDescriptorProto{
		desc: &descriptorpb.MethodDescriptorProto{
			Name:       proto.String(name),
			InputType:  proto.String(input),
			OutputType: proto.String(output),
		},
		comment: &Comment{},
	}
}

func (sd *MethodDescriptorProto) GetName() string {
	return sd.desc.GetName()
}

func (sd *MethodDescriptorProto) SetMethodOptions(options *descriptorpb.MethodOptions) *MethodDescriptorProto {
	sd.desc.Options = options
	return sd
}

func (sd *MethodDescriptorProto) AddLeadingComment(fn, leading string) *MethodDescriptorProto {
	sd.comment.LeadingComments = conv.NormalizeComment(fn, leading)

	return sd
}

func (sd *MethodDescriptorProto) AddTrailingComment(trailing string) *MethodDescriptorProto {
	sd.comment.TrailingComments = trailing

	return sd
}

func (sd *MethodDescriptorProto) AddLeadingDetachedComment(leadingDetached []string) *MethodDescriptorProto {
	sd.comment.LeadingDetachedComments = leadingDetached

	return sd
}

func (sd *MethodDescriptorProto) GetComment() *Comment {
	return sd.comment
}

func (sd *MethodDescriptorProto) Build() *descriptorpb.MethodDescriptorProto {
	return sd.desc
}
