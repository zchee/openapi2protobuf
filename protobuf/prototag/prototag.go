// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package prototag

import (
	"math"

	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	// MaxNormalTag is the maximum allowed tag number for a field in a normal message.
	MaxNormalTag = 536870911 // 2^29 - 1

	// MaxMessageSet is the maximum allowed tag number of a field in a message that
	// uses the message set wire format.
	MaxMessageSet = math.MaxInt32 - 1

	// SpecialReservedStart is the first tag in a range that is reserved and not
	// allowed for use in message definitions.
	SpecialReservedStart = 19000
	// SpecialReservedEnd is the last tag in a range that is reserved and not
	// allowed for use in message definitions.
	SpecialReservedEnd = 19999
)

var (
	fileFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("FileDescriptorProto").Fields()

	// FilePackage is the tag number of the package element in a file
	// descriptor proto.
	FilePackage = int32(fileFields.ByName("package").Number())

	// FileDependency is the tag number of the dependencies element in a
	// file descriptor proto.
	FileDependency = int32(fileFields.ByName("dependency").Number())

	// FileMessageType is the tag number of the messages element in a file
	// descriptor proto.
	FileMessageType = int32(fileFields.ByName("message_type").Number())

	// FileEnumType is the tag number of the enums element in a file descriptor
	// proto.
	FileEnumType = int32(fileFields.ByName("enum_type").Number())

	// FileServices is the tag number of the services element in a file
	// descriptor proto.
	FileServices = int32(fileFields.ByName("service").Number())

	// FileExtensions is the tag number of the extensions element in a file
	// descriptor proto.
	FileExtensions = int32(fileFields.ByName("extension").Number())

	// FileOptions is the tag number of the options element in a file
	// descriptor proto.
	FileOptions = int32(fileFields.ByName("options").Number())

	// FileSyntax is the tag number of the syntax element in a file
	// descriptor proto.
	FileSyntax = int32(fileFields.ByName("syntax").Number())
)

var (
	message       = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("DescriptorProto")
	messageFields = message.Fields()

	// MessageName is the tag number of the name element in a message
	// descriptor proto.
	MessageName = int32(messageFields.ByName("name").Number())

	// MessageFields is the tag number of the fields element in a message
	// descriptor proto.
	MessageFields = int32(messageFields.ByName("field").Number())

	// MessageNestedMessages is the tag number of the nested messages
	// element in a message descriptor proto.
	MessageNestedMessages = int32(messageFields.ByName("nested_type").Number())

	// MessageEnums is the tag number of the enums element in a message
	// descriptor proto.
	MessageEnums = int32(messageFields.ByName("enum_type").Number())

	// MessageExtensionRange is the tag number of the extension ranges
	// element in a message descriptor proto.
	MessageExtensionRange = int32(messageFields.ByName("extension_range").Number())

	// MessageExtensions is the tag number of the extensions element in a
	// message descriptor proto.
	MessageExtensions = int32(messageFields.ByName("extension").Number())

	// MessageOptions is the tag number of the options element in a message
	// descriptor proto.
	MessageOptions = int32(messageFields.ByName("options").Number())

	// MessageOneOfs is the tag number of the one-ofs element in a message
	// descriptor proto.
	MessageOneOfs = int32(messageFields.ByName("oneof_decl").Number())

	// MessageReservedRange is the tag number of the reserved ranges element
	// in a message descriptor proto.
	MessageReservedRange = int32(messageFields.ByName("reserved_range").Number())

	// MessageReservedName is the tag number of the reserved names element
	// in a message descriptor proto.
	MessageReservedName = int32(messageFields.ByName("reserved_name").Number())
)

var (
	extensionRange = message.Messages().ByName("ExtensionRange")

	// ExtensionRangeStart is the tag number of the start index in an
	// extension range proto.
	ExtensionRangeStart = int32(extensionRange.Fields().ByName("start").Number())

	// ExtensionRangeEnd is the tag number of the end index in an
	// extension range proto.
	ExtensionRangeEnd = int32(extensionRange.Fields().ByName("end").Number())

	// ExtensionRangeOptions is the tag number of the options element in an
	// extension range proto.
	ExtensionRangeOptions = int32(extensionRange.Fields().ByName("options").Number())
)

var (
	reservedRange = message.Messages().ByName("ReservedRange")

	// ReservedRangeStart is the tag number of the start index in a reserved
	// range proto.
	ReservedRangeStart = int32(reservedRange.Fields().ByName("start").Number())

	// ReservedRangeEnd is the tag number of the end index in a reserved
	// range proto.
	ReservedRangeEnd = int32(reservedRange.Fields().ByName("end").Number())
)

var (
	fieldFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("FieldDescriptorProto").Fields()

	// FieldName is the tag number of the name element in a field descriptor
	// proto.
	FieldName = int32(fieldFields.ByName("name").Number())

	// FieldExtendee is the tag number of the extendee element in a field
	// descriptor proto.
	FieldExtendee = int32(fieldFields.ByName("extendee").Number())

	// FieldNumber is the tag number of the number element in a field
	// descriptor proto.
	FieldNumber = int32(fieldFields.ByName("number").Number())

	// FieldLabel is the tag number of the label element in a field
	// descriptor proto.
	FieldLabel = int32(fieldFields.ByName("label").Number())

	// FieldType is the tag number of the type element in a field descriptor
	// proto.
	FieldType = int32(fieldFields.ByName("type").Number())

	// FieldTypeName is the tag number of the type name element in a field
	// descriptor proto.
	FieldTypeName = int32(fieldFields.ByName("type_name").Number())

	// FieldDefault is the tag number of the default value element in a
	// field descriptor proto.
	FieldDefault = int32(fieldFields.ByName("default_value").Number())

	// FieldOptions is the tag number of the options element in a field
	// descriptor proto.
	FieldOptions = int32(fieldFields.ByName("options").Number())

	// FieldJsonName is the tag number of the JSON name element in a field
	// descriptor proto.
	FieldJsonName = int32(fieldFields.ByName("json_name").Number())

	// FieldProto3Optional is the tag number of the proto3_optional element
	// in a descriptor proto.
	FieldProto3Optional = int32(fieldFields.ByName("proto3_optional").Number())
)

var (
	oneofFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("OneofDescriptorProto").Fields()

	// OneofNameTag is the tag number of the name element in a one-of
	// descriptor proto.
	OneofNameTag = int32(oneofFields.ByName("name").Number())

	// OneofOptionsTag is the tag number of the options element in a one-of
	// descriptor proto.
	OneofOptionsTag = int32(oneofFields.ByName("options").Number())
)

var (
	enumFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("EnumDescriptorProto").Fields()

	// EnumName is the tag number of the name element in an enum descriptor
	// proto.
	EnumName = int32(enumFields.ByName("name").Number())

	// EnumValues is the tag number of the values element in an enum
	// descriptor proto.
	EnumValues = int32(enumFields.ByName("value").Number())

	// EnumOptions is the tag number of the options element in an enum
	// descriptor proto.
	EnumOptions = int32(enumFields.ByName("options").Number())

	// EnumReservedRange is the tag number of the reserved ranges element in
	// an enum descriptor proto.
	EnumReservedRange = int32(enumFields.ByName("reserved_range").Number())

	// EnumReservedName is the tag number of the reserved names element in
	// an enum descriptor proto.
	EnumReservedName = int32(enumFields.ByName("reserved_name").Number())
)

var (
	enumValueFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("EnumValueDescriptorProto").Fields()

	// EnumValName is the tag number of the name element in an enum value
	// descriptor proto.
	EnumValName = int32(enumValueFields.ByName("name").Number())

	// EnumValNumber is the tag number of the number element in an enum
	// value descriptor proto.
	EnumValNumber = int32(enumValueFields.ByName("number").Number())

	// EnumValOptions is the tag number of the options element in an enum
	// value descriptor proto.
	EnumValOptions = int32(enumValueFields.ByName("options").Number())
)

var (
	serviceFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("ServiceDescriptorProto").Fields()

	// ServiceName is the tag number of the name element in a service
	// descriptor proto.
	ServiceName = int32(serviceFields.ByName("name").Number())

	// ServiceMethods is the tag number of the methods element in a service
	// descriptor proto.
	ServiceMethods = int32(serviceFields.ByName("method").Number())

	// ServiceOptions is the tag number of the options element in a service
	// descriptor proto.
	ServiceOptions = int32(serviceFields.ByName("options").Number())
)

var (
	methodFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("MethodDescriptorProto").Fields()

	// MethodName is the tag number of the name element in a method
	// descriptor proto.
	MethodName = int32(methodFields.ByName("name").Number())

	// MethodInput is the tag number of the input type element in a method
	// descriptor proto.
	MethodInput = int32(methodFields.ByName("input_type").Number())

	// MethodOutput is the tag number of the output type element in a method
	// descriptor proto.
	MethodOutput = int32(methodFields.ByName("output_type").Number())

	// MethodOptions is the tag number of the options element in a method
	// descriptor proto.
	MethodOptions = int32(methodFields.ByName("options").Number())

	// MethodInputStream is the tag number of the input stream flag in a
	// method descriptor proto.
	MethodInputStream = int32(methodFields.ByName("client_streaming").Number())

	// MethodOutputStream is the tag number of the output stream flag in a
	// method descriptor proto.
	MethodOutputStream = int32(methodFields.ByName("server_streaming").Number())
)

var (
	uninterpretedOptions       = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("UninterpretedOption")
	uninterpretedOptionsFields = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("UninterpretedOption").Fields()

	// UninterpretedOptions is the tag number of the uninterpreted options
	// element. All *Options messages use the same tag for the field that stores
	// uninterpreted options.
	UninterpretedOptions = int32(999)

	// UninterpretedName is the tag number of the name element in an
	// uninterpreted options proto.
	UninterpretedName = int32(uninterpretedOptionsFields.ByName("name").Number())

	// UninterpretedIdent is the tag number of the identifier value in an
	// uninterpreted options proto.
	UninterpretedIdent = int32(uninterpretedOptionsFields.ByName("identifier_value").Number())

	// UninterpretedPosInt is the tag number of the positive int value in an
	// uninterpreted options proto.
	UninterpretedPosInt = int32(uninterpretedOptionsFields.ByName("positive_int_value").Number())

	// UninterpretedNegInt is the tag number of the negative int value in an
	// uninterpreted options proto.
	UninterpretedNegInt = int32(uninterpretedOptionsFields.ByName("negative_int_value").Number())

	// UninterpretedDouble is the tag number of the double value in an
	// uninterpreted options proto.
	UninterpretedDouble = int32(uninterpretedOptionsFields.ByName("double_value").Number())

	// UninterpretedString is the tag number of the string value in an
	// uninterpreted options proto.
	UninterpretedString = int32(uninterpretedOptionsFields.ByName("string_value").Number())

	// UninterpretedAggregate is the tag number of the aggregate value in an
	// uninterpreted options proto.
	UninterpretedAggregate = int32(uninterpretedOptionsFields.ByName("aggregate_value").Number())

	// UninterpretedNameName is the tag number of the name element in an
	// uninterpreted option name proto.
	UninterpretedNameName = int32(uninterpretedOptions.Messages().ByName("NamePart").Fields().ByName("name_part").Number())

	// UninterpretedNameIsExtension is the tag number of the name element in an
	// uninterpreted option name proto.
	UninterpretedNameIsExtension = int32(uninterpretedOptions.Messages().ByName("NamePart").Fields().ByName("is_extension").Number())
)
