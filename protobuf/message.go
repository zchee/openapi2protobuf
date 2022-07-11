// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
)

type MessageDescriptorProto struct {
	desc *descriptorpb.DescriptorProto

	field       map[string]bool
	fieldMumber int32
	enum        map[string]bool
	nested      map[string]bool

	locations []*descriptorpb.SourceCodeInfo_Location
	path      []int32
}

func NewMessageDescriptorProto(name string) *MessageDescriptorProto {
	return &MessageDescriptorProto{
		desc: &descriptorpb.DescriptorProto{
			Name: proto.String(name),
		},
		field:  make(map[string]bool),
		enum:   make(map[string]bool),
		nested: make(map[string]bool),
		path:   []int32{},
	}
}

func (md *MessageDescriptorProto) Clone() *MessageDescriptorProto {
	desc := &descriptorpb.DescriptorProto{}
	*desc = *md.desc

	mdesc := &MessageDescriptorProto{
		desc:        desc,
		field:       md.field,
		fieldMumber: md.fieldMumber,
		nested:      md.nested,
	}

	return mdesc
}

func (md *MessageDescriptorProto) GetName() string {
	return md.desc.GetName()
}

func (md *MessageDescriptorProto) SetName(name string) *MessageDescriptorProto {
	md.desc.Name = proto.String(name)
	return md
}

func (md *MessageDescriptorProto) AddComment(leading, trailing string, leadingDetached []string) *MessageDescriptorProto {
	loc := &descriptorpb.SourceCodeInfo_Location{
		Path: append(md.path, File_messageTypeTag),
		Span: []int32{0, 0, 0},
	}
	if leading != "" {
		loc.LeadingComments = proto.String(leading)
	}
	if trailing != "" {
		loc.TrailingComments = proto.String(trailing)
	}
	if leadingDetached != nil {
		loc.LeadingDetachedComments = leadingDetached
	}

	md.locations = append(md.locations, loc)

	return md
}

func (md *MessageDescriptorProto) GetLocation() []*descriptorpb.SourceCodeInfo_Location {
	return md.locations
}

func (md *MessageDescriptorProto) AddField(field *FieldDescriptorProto) *MessageDescriptorProto {
	if md.field[field.GetName()] {
		return md
	}

	md.fieldMumber++
	field.desc.Number = proto.Int32(md.fieldMumber)
	md.field[field.GetName()] = true
	md.desc.Field = append(md.desc.Field, field.Build())

	return md
}

func (md *MessageDescriptorProto) SortField(order []string) *MessageDescriptorProto {
	propOrder := make([]string, len(order))
	for i, s := range order {
		propOrder[i] = conv.NormalizeFieldName(s)
	}

	n := len(md.desc.Field)
	fdescFields := make([]*descriptorpb.FieldDescriptorProto, n)
	copy(fdescFields, md.desc.Field)
	md.desc.Field = make([]*descriptorpb.FieldDescriptorProto, n) // reset

	i := 0
	for _, field := range propOrder {
		for _, msgfield := range fdescFields {
			if msgfield.GetName() == field {
				msgfield.Number = proto.Int32(int32(i + 1))
				md.desc.Field[i] = msgfield
				i++
				break
			}
		}
	}

	return md
}

func (md *MessageDescriptorProto) GetFieldByName(name string) *FieldDescriptorProto {
	for _, field := range md.desc.Field {
		if field.GetName() == name {
			return &FieldDescriptorProto{
				desc: field,
			}
		}
	}

	return nil
}

func (md *MessageDescriptorProto) GetFieldType() *descriptorpb.FieldDescriptorProto_Type {
	if len(md.desc.Field) == 1 {
		return md.desc.Field[0].GetType().Enum()
	}

	if len(md.desc.Field) > 1 {
		return descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
	}

	if len(md.desc.EnumType) > 1 {
		return descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum()
	}

	return descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
}

func (md *MessageDescriptorProto) IsEmptyField() bool {
	return len(md.desc.Field) == 0 && len(md.desc.EnumType) == 0 && len(md.desc.NestedType) == 0
}

func (md *MessageDescriptorProto) AddExtension(ext *FieldDescriptorProto) *MessageDescriptorProto {
	md.desc.Extension = append(md.desc.Extension, ext.Build())
	return md
}

func (md *MessageDescriptorProto) GetNestedMessages() []string {
	nesteds := make([]string, len(md.nested))
	i := 0
	for nested := range md.nested {
		nesteds[i] = nested
		i++
	}

	return nesteds
}

func (md *MessageDescriptorProto) HasNestedMessage(nested string) bool {
	return md.nested[nested]
}

func (md *MessageDescriptorProto) AddNestedMessage(nested *MessageDescriptorProto) *MessageDescriptorProto {
	if md.nested[nested.GetName()] {
		return md
	}
	md.desc.NestedType = append(md.desc.NestedType, nested.Build())
	md.nested[nested.GetName()] = true

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

func (md *MessageDescriptorProto) GetOneofIndex() int32 {
	return int32(len(md.desc.OneofDecl) - 1)
}

func (md *MessageDescriptorProto) SetMessageSetWireFormat(messageSetWireFormat bool) *MessageDescriptorProto {
	if md.desc.Options == nil {
		md.desc.Options = &descriptorpb.MessageOptions{}
	}
	md.desc.Options.MessageSetWireFormat = proto.Bool(messageSetWireFormat)

	return md
}

func (md *MessageDescriptorProto) SetNoStandardDescriptorAccessor(noStandardDescriptorAccessor bool) *MessageDescriptorProto {
	if md.desc.Options == nil {
		md.desc.Options = &descriptorpb.MessageOptions{}
	}
	md.desc.Options.NoStandardDescriptorAccessor = proto.Bool(noStandardDescriptorAccessor)

	return md
}

func (md *MessageDescriptorProto) SetDeprecated(deprecated bool) *MessageDescriptorProto {
	if md.desc.Options == nil {
		md.desc.Options = &descriptorpb.MessageOptions{}
	}
	md.desc.Options.Deprecated = proto.Bool(deprecated)

	return md
}

func (md *MessageDescriptorProto) SetMapEntry(mapEntry bool) *MessageDescriptorProto {
	if md.desc.Options == nil {
		md.desc.Options = &descriptorpb.MessageOptions{}
	}
	md.desc.Options.MapEntry = proto.Bool(mapEntry)

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
