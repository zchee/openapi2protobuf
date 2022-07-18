// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf/prototag"
)

type MessageDescriptorProto struct {
	desc    *descriptorpb.DescriptorProto
	comment Comment

	field          map[string]bool
	fieldNumber    int32
	fieldLocations map[int32]*descriptorpb.SourceCodeInfo_Location
	enum           map[string]bool
	enumLocations  []*descriptorpb.SourceCodeInfo_Location
	nested         map[string]bool
}

func NewMessageDescriptorProto(name string) *MessageDescriptorProto {
	return &MessageDescriptorProto{
		desc: &descriptorpb.DescriptorProto{
			Name: proto.String(name),
		},
		field:          make(map[string]bool),
		fieldLocations: make(map[int32]*descriptorpb.SourceCodeInfo_Location),
		enum:           make(map[string]bool),
		nested:         make(map[string]bool),
	}
}

func (md *MessageDescriptorProto) GetName() string {
	return md.desc.GetName()
}

func (md *MessageDescriptorProto) SetName(name string) *MessageDescriptorProto {
	md.desc.Name = proto.String(name)
	return md
}

func (md *MessageDescriptorProto) AddLeadingComment(fn, leading string) *MessageDescriptorProto {
	md.comment.LeadingComments = conv.NormalizeComment(fn, leading)

	return md
}

func (md *MessageDescriptorProto) AddTrailingComment(trailing string) *MessageDescriptorProto {
	md.comment.TrailingComments = trailing

	return md
}

func (md *MessageDescriptorProto) AddLeadingDetachedComment(leadingDetached []string) *MessageDescriptorProto {
	md.comment.LeadingDetachedComments = leadingDetached

	return md
}

func (md *MessageDescriptorProto) GetComment() Comment {
	return md.comment
}

func (md *MessageDescriptorProto) AddField(field *FieldDescriptorProto) *MessageDescriptorProto {
	if md.field[field.GetName()] {
		return md
	}

	md.fieldNumber++
	field.desc.Number = proto.Int32(md.fieldNumber)
	md.field[field.GetName()] = true

	comments := field.GetComment()
	loc := &descriptorpb.SourceCodeInfo_Location{
		LeadingComments:         proto.String(comments.LeadingComments),
		TrailingComments:        proto.String(comments.TrailingComments),
		LeadingDetachedComments: comments.LeadingDetachedComments,
		Path:                    []int32{prototag.MessageFields, md.fieldNumber - 1},
	}
	md.fieldLocations[md.fieldNumber-1] = loc
	md.desc.Field = append(md.desc.Field, field.Build())

	return md
}

func (md *MessageDescriptorProto) GetFieldLocations() []*descriptorpb.SourceCodeInfo_Location {
	fieldLocations := make([]*descriptorpb.SourceCodeInfo_Location, len(md.fieldLocations))
	i := 0
	for _, loc := range md.fieldLocations {
		fieldLocations[i] = loc
		i++
	}

	return append(fieldLocations, md.enumLocations...)
}

func (md *MessageDescriptorProto) SortField(order []string) *MessageDescriptorProto {
	propOrder := make([]string, len(order))
	for i, s := range order {
		propOrder[i] = conv.NormalizeFieldName(s)
	}

	n := len(md.desc.Field)
	fdescFields := make([]*descriptorpb.FieldDescriptorProto, n)
	copy(fdescFields, md.desc.Field)
	md.desc.Field = md.desc.Field[0:] // reset

	i := 0
	for _, field := range propOrder {
		for _, msgfield := range fdescFields {
			if msgfield.GetName() == field {
				md.fieldLocations[msgfield.GetNumber()-1].Path[1] = int32(i)

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
	comments := enum.GetComment()
	loc := &descriptorpb.SourceCodeInfo_Location{
		LeadingComments:         proto.String(comments.LeadingComments),
		TrailingComments:        proto.String(comments.TrailingComments),
		LeadingDetachedComments: comments.LeadingDetachedComments,
		Path:                    []int32{prototag.MessageEnums, md.fieldNumber - 1},
	}
	md.enumLocations = append(md.enumLocations, loc)
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
