// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"go.lsp.dev/openapi2protobuf/internal/conv"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type FieldDescriptorProto struct {
	desc    *descriptorpb.FieldDescriptorProto
	number  int32
	comment *Comment

	fieldOrder []string
}

func NewFieldDescriptorProto(name string, fieldType *descriptorpb.FieldDescriptorProto_Type) *FieldDescriptorProto {
	fid := &FieldDescriptorProto{
		desc: &descriptorpb.FieldDescriptorProto{
			Name: proto.String(name),
			Type: fieldType,
		},
		comment: &Comment{},
	}

	return fid
}

func (fid *FieldDescriptorProto) GetName() string {
	return fid.desc.GetName()
}

func (fid *FieldDescriptorProto) GetNumber() int32 {
	return fid.desc.GetNumber()
}

func (fid *FieldDescriptorProto) SetNumber() *FieldDescriptorProto {
	fid.number++
	fid.desc.Number = proto.Int32(fid.number)

	return fid
}

func (fid *FieldDescriptorProto) GetTypeName() *string {
	return fid.desc.TypeName
}

func (fid *FieldDescriptorProto) SetTypeName(name string) *FieldDescriptorProto {
	fid.desc.TypeName = proto.String(name)

	return fid
}

func (fid *FieldDescriptorProto) SetOneofIndex(idx int32) *FieldDescriptorProto {
	fid.desc.OneofIndex = proto.Int32(idx)

	return fid
}

func (fid *FieldDescriptorProto) SetRepeated() *FieldDescriptorProto {
	fid.desc.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()

	return fid
}

func (fid *FieldDescriptorProto) AddLeadingComment(fn, leading string) *FieldDescriptorProto {
	fid.comment.LeadingComments = conv.NormalizeComment(fn, leading)

	return fid
}

func (fid *FieldDescriptorProto) AddTrailingComment(trailing string) *FieldDescriptorProto {
	fid.comment.TrailingComments = trailing

	return fid
}

func (fid *FieldDescriptorProto) AddLeadingDetachedComment(leadingDetached []string) *FieldDescriptorProto {
	fid.comment.LeadingDetachedComments = leadingDetached

	return fid
}

func (fid *FieldDescriptorProto) HasComment() bool {
	comment := fid.comment
	return comment.LeadingComments != "" || comment.TrailingComments != "" || len(comment.LeadingDetachedComments) != 0
}

func (fid *FieldDescriptorProto) GetComment() *Comment {
	return fid.comment
}

func (fid *FieldDescriptorProto) SetJsonName(jsonName string) *FieldDescriptorProto {
	fid.desc.JsonName = proto.String(jsonName)

	return fid
}

func (fid *FieldDescriptorProto) SetFieldOption(fieldOptions *descriptorpb.FieldOptions) *FieldDescriptorProto {
	fid.desc.Options = fieldOptions

	return fid
}

func (fid *FieldDescriptorProto) SetProto3Optional() *FieldDescriptorProto {
	fid.desc.Proto3Optional = proto.Bool(true)

	return fid
}

func (fid *FieldDescriptorProto) Build() *descriptorpb.FieldDescriptorProto {
	return fid.desc
}

func FieldTypeDouble() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum()
}

func FieldTypeFloat() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_FLOAT.Enum()
}

func FieldTypeInt64() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum()
}

func FieldTypeUint64() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum()
}

func FieldTypeInt32() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()
}

func FieldTypeFixed64() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_FIXED64.Enum()
}

func FieldTypeFixed32() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_FIXED32.Enum()
}

func FieldTypeBool() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()
}

func FieldTypeString() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
}

func FieldTypeMessage() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
}

func FieldTypeBytes() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_BYTES.Enum()
}

func FieldTypeUint32() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum()
}

func FieldTypeEnum() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum()
}

func FieldTypeSfixed32() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_SFIXED32.Enum()
}

func FieldTypeSfixed64() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_SFIXED64.Enum()
}

func FieldTypeSint32() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_SINT32.Enum()
}

func FieldTypeSint64() *descriptorpb.FieldDescriptorProto_Type {
	return descriptorpb.FieldDescriptorProto_TYPE_SINT64.Enum()
}
