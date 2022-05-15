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

func NewEnumDescriptorProto(name string) *EnumDescriptorProto {
	ed := &EnumDescriptorProto{
		desc: &descriptorpb.EnumDescriptorProto{
			Name: proto.String(name),
		},
	}

	return ed
}

func (ed *EnumDescriptorProto) GetName() string {
	return *ed.desc.Name
}

func (ed *EnumDescriptorProto) AddValue(value *EnumValueDescriptorProto) *EnumDescriptorProto {
	ed.desc.Value = append(ed.desc.Value, value.Build())
	return ed
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
