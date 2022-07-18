// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
)

type EnumDescriptorProto struct {
	desc    *descriptorpb.EnumDescriptorProto
	comment Comment
}

func NewEnumDescriptorProto(name string) *EnumDescriptorProto {
	ed := &EnumDescriptorProto{
		desc: &descriptorpb.EnumDescriptorProto{
			Name: proto.String(name),
		},
	}

	return ed
}

func (ed *EnumDescriptorProto) GetName() string {
	return ed.desc.GetName()
}

func (ed *EnumDescriptorProto) AddValue(value *EnumValueDescriptorProto) *EnumDescriptorProto {
	ed.desc.Value = append(ed.desc.Value, value.Build())
	return ed
}

func (ed *EnumDescriptorProto) GetValue() []*descriptorpb.EnumValueDescriptorProto {
	return ed.desc.Value
}

func (ed *EnumDescriptorProto) AddLeadingComment(fn, leading string) *EnumDescriptorProto {
	ed.comment.LeadingComments = conv.NormalizeComment(fn, leading)

	return ed
}

func (ed *EnumDescriptorProto) AddTrailingComment(trailing string) *EnumDescriptorProto {
	ed.comment.TrailingComments = trailing

	return ed
}

func (ed *EnumDescriptorProto) AddLeadingDetachedComment(leadingDetached []string) *EnumDescriptorProto {
	ed.comment.LeadingDetachedComments = leadingDetached

	return ed
}

func (ed *EnumDescriptorProto) GetComment() Comment {
	return ed.comment
}

func (ed *EnumDescriptorProto) Build() *descriptorpb.EnumDescriptorProto {
	return ed.desc
}

type EnumValueDescriptorProto struct {
	desc *descriptorpb.EnumValueDescriptorProto
}

func NewEnumValueDescriptorProto(name string, number int32) *EnumValueDescriptorProto {
	return &EnumValueDescriptorProto{
		desc: &descriptorpb.EnumValueDescriptorProto{
			Name:   proto.String(name),
			Number: proto.Int32(number),
		},
	}
}

func (evd *EnumValueDescriptorProto) GetName() string {
	return evd.desc.GetName()
}

func (evd *EnumValueDescriptorProto) SetDeprecated() *EnumValueDescriptorProto {
	evd.desc.Options.Deprecated = proto.Bool(true)
	return evd
}

func (evd *EnumValueDescriptorProto) Build() *descriptorpb.EnumValueDescriptorProto {
	return evd.desc
}

type EnumValues []*EnumValueDescriptorProto

func (evs EnumValues) Build() []*descriptorpb.EnumValueDescriptorProto {
	values := make([]*descriptorpb.EnumValueDescriptorProto, len(evs))
	for i, ev := range evs {
		values[i] = ev.Build()
	}

	return values
}
