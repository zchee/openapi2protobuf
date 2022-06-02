// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package lsp

import (
	"go.lsp.dev/openapi2protobuf/compiler"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

func init() {
	compiler.AdditionalMessages = AdditialMessages()
	compiler.DependencyProto = []string{"google.protobuf.Any"}
}

func AdditialMessages() []*protobuf.MessageDescriptorProto {
	lspAnyMsg := protobuf.NewMessageDescriptorProto("LSPAny")
	lspAnyField := protobuf.NewFieldDescriptorProto("any", protobuf.FieldTypeMessage()).SetTypeName(protobuf.Any).SetNumber()
	lspAnyMsg.AddField(lspAnyField)

	lspArrayMsg := protobuf.NewMessageDescriptorProto("LSPArray")
	lspArrayField := protobuf.NewFieldDescriptorProto("array", protobuf.FieldTypeMessage()).SetTypeName(protobuf.Any).SetNumber().SetRepeated()
	lspArrayMsg.AddField(lspArrayField)

	lspObjectMsg := protobuf.NewMessageDescriptorProto("LSPObject").SetMapEntry(true)
	lspObjectMapFieldMsg := protobuf.NewMessageDescriptorProto("LSPObjectMapField").SetMapEntry(true)
	lspObjectKeyField := protobuf.NewFieldDescriptorProto("key", protobuf.FieldTypeString()).SetProto3Optional()
	lspObjectMapFieldMsg.AddField(lspObjectKeyField)
	lspObjectValueField := protobuf.NewFieldDescriptorProto("value", protobuf.FieldTypeMessage()).SetTypeName(protobuf.Any).SetProto3Optional()
	lspObjectMapFieldMsg.AddField(lspObjectValueField)
	lspObjectMsg.AddNestedMessage(lspObjectMapFieldMsg)
	lspObjectField := protobuf.NewFieldDescriptorProto("object", protobuf.FieldTypeMessage()).SetTypeName("LSPObjectMapField").SetNumber().SetRepeated()
	lspObjectMsg.AddField(lspObjectField)

	return []*protobuf.MessageDescriptorProto{
		lspAnyMsg,
		lspArrayMsg,
		lspObjectMsg,
	}
}
