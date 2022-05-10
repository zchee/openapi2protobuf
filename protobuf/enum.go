// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type EnumDescriptorProto struct {
	desc *descriptorpb.EnumDescriptorProto
}

func NewEnumDescriptorProto(name string, values ...*EnumValueDescriptorProto) *EnumDescriptorProto {
	ed := &EnumDescriptorProto{
		desc: &descriptorpb.EnumDescriptorProto{
			Name: proto.String(name),
		},
	}
	if len(values) > 0 {
		ed.desc.Value = EnumValues(values).Build()
	}

	return ed
}

func (ed *EnumDescriptorProto) AddValue(value *EnumValueDescriptorProto) *EnumDescriptorProto {
	ed.desc.Value = append(ed.desc.Value, value.Descriptor())
	return ed
}

func (ed *EnumDescriptorProto) Descriptor() *descriptorpb.EnumDescriptorProto {
	return ed.desc
}

type EnumValueDescriptorProto struct {
	desc *descriptorpb.EnumValueDescriptorProto
}

func NewEnumValueDescriptorProto(name string, number int32, isDeprecated bool) *EnumValueDescriptorProto {
	return &EnumValueDescriptorProto{
		desc: &descriptorpb.EnumValueDescriptorProto{
			Name:   proto.String(name),
			Number: proto.Int32(number),
			Options: &descriptorpb.EnumValueOptions{
				Deprecated: proto.Bool(isDeprecated),
			},
		},
	}
}

func (evd *EnumValueDescriptorProto) Descriptor() *descriptorpb.EnumValueDescriptorProto {
	return evd.desc
}

type EnumValues []*EnumValueDescriptorProto

func (evs EnumValues) Build() []*descriptorpb.EnumValueDescriptorProto {
	values := make([]*descriptorpb.EnumValueDescriptorProto, len(evs))
	for i, ev := range evs {
		values[i] = ev.Descriptor()
	}

	return values
}
