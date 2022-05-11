// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"reflect"
	"unsafe"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageDescriptorProto struct {
	desc           *descriptorpb.DescriptorProto
	nestedMessages map[string]bool
	field          map[string]bool
	number         int32
}

func NewMessageDescriptorProto(name string) *MessageDescriptorProto {
	return &MessageDescriptorProto{
		desc: &descriptorpb.DescriptorProto{
			Name: proto.String(name),
		},
		nestedMessages: make(map[string]bool),
		field:          make(map[string]bool),
		number:         int32(1),
	}
}

func (md *MessageDescriptorProto) GetName() string {
	return md.desc.GetName()
}

func (md *MessageDescriptorProto) SetName(name string) *MessageDescriptorProto {
	md.desc.Name = proto.String(name)
	return md
}

func (md *MessageDescriptorProto) AddField(field *FieldDescriptorProto) *MessageDescriptorProto {
	if md.field[field.GetName()] {
		return md
	}
	field.desc.Number = proto.Int32(md.number)
	md.desc.Field = append(md.desc.Field, field.Build())
	md.number++
	md.field[field.GetName()] = true // cast to string is allocation free
	return md
}

func (md *MessageDescriptorProto) AddExtension(ext *FieldDescriptorProto) *MessageDescriptorProto {
	md.desc.Extension = append(md.desc.Extension, ext.Build())
	return md
}

func (md *MessageDescriptorProto) AddNestedMessage(nested *MessageDescriptorProto) *MessageDescriptorProto {
	if md.nestedMessages[nested.GetName()] {
		return md
	}
	md.desc.NestedType = append(md.desc.NestedType, nested.Build())
	md.nestedMessages[nested.GetName()] = true // cast to string is allocation free
	return md
}

func (md *MessageDescriptorProto) AddEnumType(enum *EnumDescriptorProto) *MessageDescriptorProto {
	md.desc.EnumType = append(md.desc.EnumType, enum.Build())
	return md
}

func (md *MessageDescriptorProto) AddOneof(oneof *OneofDescriptorProto) *MessageDescriptorProto {
	md.desc.OneofDecl = append(md.desc.OneofDecl, oneof.Build())
	return md
}

func (md *MessageDescriptorProto) SetMessageOptions(options *descriptorpb.MessageOptions) *MessageDescriptorProto {
	md.desc.Options = options
	return md
}

func (md *MessageDescriptorProto) SetReservedRange(reservedRange ...*descriptorpb.DescriptorProto_ReservedRange) *MessageDescriptorProto {
	md.desc.ReservedRange = append(md.desc.ReservedRange, reservedRange...)
	return md
}

func (md *MessageDescriptorProto) SetExtensionRange(ranges ...*descriptorpb.DescriptorProto_ExtensionRange) *MessageDescriptorProto {
	md.desc.ExtensionRange = append(md.desc.ExtensionRange, ranges...)
	return md
}

func (md *MessageDescriptorProto) Build() *descriptorpb.DescriptorProto {
	return md.desc
}

// packStruct returns a byte slice view of a struct.
func packStruct(s interface{}) []byte {
	v := reflect.ValueOf(s)
	sz := int(v.Elem().Type().Size())

	return unsafe.Slice((*byte)(unsafe.Pointer(v.Pointer())), sz)
}
