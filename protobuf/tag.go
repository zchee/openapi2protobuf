// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"math"
	"unicode"
	"unicode/utf8"

	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	// MaxNormalTag is the maximum allowed tag number for a field in a normal message.
	MaxNormalTag = 536870911 // 2^29 - 1

	// MaxMessageSetTag is the maximum allowed tag number of a field in a message that
	// uses the message set wire format.
	MaxMessageSetTag = math.MaxInt32 - 1

	// MaxTag is the maximum allowed tag number. (It is the same as MaxMessageSetTag
	// since that is the absolute highest allowed.)
	MaxTag = MaxMessageSetTag

	// SpecialReservedStart is the first tag in a range that is reserved and not
	// allowed for use in message definitions.
	SpecialReservedStart = 19000
	// SpecialReservedEnd is the last tag in a range that is reserved and not
	// allowed for use in message definitions.
	SpecialReservedEnd = 19999
)

var (
	file = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("FileDescriptorProto")

	// File_packageTag is the tag number of the package element in a file
	// descriptor proto.
	File_packageTag = int32(file.Fields().ByName("package").Number())
	// File_dependencyTag is the tag number of the dependencies element in a
	// file descriptor proto.
	File_dependencyTag = int32(file.Fields().ByName("dependency").Number())
	// File_messagesTag is the tag number of the messages element in a file
	// descriptor proto.
	File_messageTypeTag = int32(file.Fields().ByName("message_type").Number())
	// File_enumsTag is the tag number of the enums element in a file descriptor
	// proto.
	File_enumTypeTag = int32(file.Fields().ByName("enum_type").Number())
	// File_servicesTag is the tag number of the services element in a file
	// descriptor proto.
	File_servicesTag = int32(file.Fields().ByName("service").Number())
	// File_extensionsTag is the tag number of the extensions element in a file
	// descriptor proto.
	File_extensionsTag = int32(file.Fields().ByName("extension").Number())
	// File_optionsTag is the tag number of the options element in a file
	// descriptor proto.
	File_optionsTag = int32(file.Fields().ByName("options").Number())
	// File_syntaxTag is the tag number of the syntax element in a file
	// descriptor proto.
	File_syntaxTag = int32(file.Fields().ByName("syntax").Number())

	message = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("DescriptorProto")

	// Message_nameTag is the tag number of the name element in a message
	// descriptor proto.
	Message_nameTag = int32(message.Fields().ByName("name").Number())
	// Message_fieldsTag is the tag number of the fields element in a message
	// descriptor proto.
	Message_fieldsTag = int32(message.Fields().ByName("field").Number())
	// Message_nestedMessagesTag is the tag number of the nested messages
	// element in a message descriptor proto.
	Message_nestedMessagesTag = int32(message.Fields().ByName("nested_type").Number())
	// Message_enumsTag is the tag number of the enums element in a message
	// descriptor proto.
	Message_enumsTag = int32(message.Fields().ByName("enum_type").Number())
	// Message_extensionRangeTag is the tag number of the extension ranges
	// element in a message descriptor proto.
	Message_extensionRangeTag = int32(message.Fields().ByName("extension_range").Number())
	// Message_extensionsTag is the tag number of the extensions element in a
	// message descriptor proto.
	Message_extensionsTag = int32(message.Fields().ByName("extension").Number())
	// Message_optionsTag is the tag number of the options element in a message
	// descriptor proto.
	Message_optionsTag = int32(message.Fields().ByName("options").Number())
	// Message_oneOfsTag is the tag number of the one-ofs element in a message
	// descriptor proto.
	Message_oneOfsTag = int32(message.Fields().ByName("oneof_decl").Number())
	// Message_reservedRangeTag is the tag number of the reserved ranges element
	// in a message descriptor proto.
	Message_reservedRangeTag = int32(message.Fields().ByName("reserved_range").Number())
	// Message_reservedNameTag is the tag number of the reserved names element
	// in a message descriptor proto.
	Message_reservedNameTag = int32(message.Fields().ByName("reserved_name").Number())

	extensionRange = message.Messages().ByName("ExtensionRange")

	// ExtensionRange_startTag is the tag number of the start index in an
	// extension range proto.
	ExtensionRange_startTag = int32(extensionRange.Fields().ByName("start").Number())
	// ExtensionRange_endTag is the tag number of the end index in an
	// extension range proto.
	ExtensionRange_endTag = int32(extensionRange.Fields().ByName("end").Number())
	// ExtensionRange_optionsTag is the tag number of the options element in an
	// extension range proto.
	ExtensionRange_optionsTag = int32(extensionRange.Fields().ByName("options").Number())

	reservedRange = message.Messages().ByName("ReservedRange")

	// ReservedRange_startTag is the tag number of the start index in a reserved
	// range proto.
	ReservedRange_startTag = int32(reservedRange.Fields().ByName("start").Number())
	// ReservedRange_endTag is the tag number of the end index in a reserved
	// range proto.
	ReservedRange_endTag = int32(reservedRange.Fields().ByName("end").Number())

	field = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("FieldDescriptorProto")

	// Field_nameTag is the tag number of the name element in a field descriptor
	// proto.
	Field_nameTag = int32(field.Fields().ByName("name").Number())
	// Field_extendeeTag is the tag number of the extendee element in a field
	// descriptor proto.
	Field_extendeeTag = int32(field.Fields().ByName("extendee").Number())
	// Field_numberTag is the tag number of the number element in a field
	// descriptor proto.
	Field_numberTag = int32(field.Fields().ByName("number").Number())
	// Field_labelTag is the tag number of the label element in a field
	// descriptor proto.
	Field_labelTag = int32(field.Fields().ByName("label").Number())
	// Field_typeTag is the tag number of the type element in a field descriptor
	// proto.
	Field_typeTag = int32(field.Fields().ByName("type").Number())
	// Field_typeNameTag is the tag number of the type name element in a field
	// descriptor proto.
	Field_typeNameTag = int32(field.Fields().ByName("type_name").Number())
	// Field_defaultTag is the tag number of the default value element in a
	// field descriptor proto.
	Field_defaultTag = int32(field.Fields().ByName("default_value").Number())
	// Field_optionsTag is the tag number of the options element in a field
	// descriptor proto.
	Field_optionsTag = int32(field.Fields().ByName("options").Number())
	// Field_jsonNameTag is the tag number of the JSON name element in a field
	// descriptor proto.
	Field_jsonNameTag = int32(field.Fields().ByName("json_name").Number())
	// Field_proto3OptionalTag is the tag number of the proto3_optional element
	// in a descriptor proto.
	Field_proto3OptionalTag = int32(field.Fields().ByName("proto3_optional").Number())

	oneOf = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("OneOfDescriptorProto")

	// OneOf_nameTag is the tag number of the name element in a one-of
	// descriptor proto.
	OneOf_nameTag = int32(field.Fields().ByName("name").Number())
	// OneOf_optionsTag is the tag number of the options element in a one-of
	// descriptor proto.
	OneOf_optionsTag = int32(field.Fields().ByName("options").Number())

	enum = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("EnumDescriptorProto")

	// Enum_nameTag is the tag number of the name element in an enum descriptor
	// proto.
	Enum_nameTag = int32(enum.Fields().ByName("name").Number())
	// Enum_valuesTag is the tag number of the values element in an enum
	// descriptor proto.
	Enum_valuesTag = int32(enum.Fields().ByName("value").Number())
	// Enum_optionsTag is the tag number of the options element in an enum
	// descriptor proto.
	Enum_optionsTag = int32(enum.Fields().ByName("options").Number())
	// Enum_reservedRangeTag is the tag number of the reserved ranges element in
	// an enum descriptor proto.
	Enum_reservedRangeTag = int32(enum.Fields().ByName("reserved_range").Number())
	// Enum_reservedNameTag is the tag number of the reserved names element in
	// an enum descriptor proto.
	Enum_reservedNameTag = int32(enum.Fields().ByName("reserved_name").Number())

	enumValue = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("EnumValueDescriptorProto")

	// EnumVal_nameTag is the tag number of the name element in an enum value
	// descriptor proto.
	EnumVal_nameTag = int32(enumValue.Fields().ByName("name").Number())
	// EnumVal_numberTag is the tag number of the number element in an enum
	// value descriptor proto.
	EnumVal_numberTag = int32(enumValue.Fields().ByName("number").Number())
	// EnumVal_optionsTag is the tag number of the options element in an enum
	// value descriptor proto.
	EnumVal_optionsTag = int32(enumValue.Fields().ByName("options").Number())

	service = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("ServiceDescriptorProto")

	// Service_nameTag is the tag number of the name element in a service
	// descriptor proto.
	Service_nameTag = int32(service.Fields().ByName("name").Number())
	// Service_methodsTag is the tag number of the methods element in a service
	// descriptor proto.
	Service_methodsTag = int32(service.Fields().ByName("method").Number())
	// Service_optionsTag is the tag number of the options element in a service
	// descriptor proto.
	Service_optionsTag = int32(service.Fields().ByName("options").Number())

	method = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("MethodDescriptorProto")

	// Method_nameTag is the tag number of the name element in a method
	// descriptor proto.
	Method_nameTag = int32(method.Fields().ByName("name").Number())
	// Method_inputTag is the tag number of the input type element in a method
	// descriptor proto.
	Method_inputTag = int32(method.Fields().ByName("input_type").Number())
	// Method_outputTag is the tag number of the output type element in a method
	// descriptor proto.
	Method_outputTag = int32(method.Fields().ByName("output_type").Number())
	// Method_optionsTag is the tag number of the options element in a method
	// descriptor proto.
	Method_optionsTag = int32(method.Fields().ByName("options").Number())
	// Method_inputStreamTag is the tag number of the input stream flag in a
	// method descriptor proto.
	Method_inputStreamTag = int32(method.Fields().ByName("client_streaming").Number())
	// Method_outputStreamTag is the tag number of the output stream flag in a
	// method descriptor proto.
	Method_outputStreamTag = int32(method.Fields().ByName("server_streaming").Number())

	uninterpretedOptions = descriptorpb.File_google_protobuf_descriptor_proto.Messages().ByName("UninterpretedOption")

	// UninterpretedOptionsTag is the tag number of the uninterpreted options
	// element. All *Options messages use the same tag for the field that stores
	// uninterpreted options.
	UninterpretedOptionsTag = int32(999)

	// Uninterpreted_nameTag is the tag number of the name element in an
	// uninterpreted options proto.
	Uninterpreted_nameTag = int32(uninterpretedOptions.Fields().ByName("name").Number())
	// Uninterpreted_identTag is the tag number of the identifier value in an
	// uninterpreted options proto.
	Uninterpreted_identTag = int32(uninterpretedOptions.Fields().ByName("identifier_value").Number())
	// Uninterpreted_posIntTag is the tag number of the positive int value in an
	// uninterpreted options proto.
	Uninterpreted_posIntTag = int32(uninterpretedOptions.Fields().ByName("positive_int_value").Number())
	// Uninterpreted_negIntTag is the tag number of the negative int value in an
	// uninterpreted options proto.
	Uninterpreted_negIntTag = int32(uninterpretedOptions.Fields().ByName("negative_int_value").Number())
	// Uninterpreted_doubleTag is the tag number of the double value in an
	// uninterpreted options proto.
	Uninterpreted_doubleTag = int32(uninterpretedOptions.Fields().ByName("double_value").Number())
	// Uninterpreted_stringTag is the tag number of the string value in an
	// uninterpreted options proto.
	Uninterpreted_stringTag = int32(uninterpretedOptions.Fields().ByName("string_value").Number())
	// Uninterpreted_aggregateTag is the tag number of the aggregate value in an
	// uninterpreted options proto.
	Uninterpreted_aggregateTag = int32(uninterpretedOptions.Fields().ByName("aggregate_value").Number())
	// UninterpretedName_nameTag is the tag number of the name element in an
	// uninterpreted option name proto.
	UninterpretedName_nameTag = int32(uninterpretedOptions.Messages().ByName("NamePart").Fields().ByName("name_part").Number())
	// UninterpretedName_nameTag is the tag number of the name element in an
	// uninterpreted option name proto.
	UninterpretedName_isExtensionTag = int32(uninterpretedOptions.Messages().ByName("NamePart").Fields().ByName("is_extension").Number())
)

// JsonName returns the default JSON name for a field with the given name.
func JsonName(name string) string {
	var js []rune
	nextUpper := false
	for i, r := range name {
		if r == '_' {
			nextUpper = true
			continue
		}
		if i == 0 {
			js = append(js, r)
		} else if nextUpper {
			nextUpper = false
			js = append(js, unicode.ToUpper(r))
		} else {
			js = append(js, r)
		}
	}
	return string(js)
}

// InitCap returns the given field name, but with the first letter capitalized.
func InitCap(name string) string {
	r, sz := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[sz:]
}

// CreatePrefixList returns a list of package prefixes to search when resolving
// a symbol name. If the given package is blank, it returns only the empty
// string. If the given package contains only one token, e.g. "foo", it returns
// that token and the empty string, e.g. ["foo", ""]. Otherwise, it returns
// successively shorter prefixes of the package and then the empty string. For
// example, for a package named "foo.bar.baz" it will return the following list:
//
//	["foo.bar.baz", "foo.bar", "foo", ""]
func CreatePrefixList(pkg string) []string {
	if pkg == "" {
		return []string{""}
	}

	numDots := 0
	// one pass to pre-allocate the returned slice
	for i := 0; i < len(pkg); i++ {
		if pkg[i] == '.' {
			numDots++
		}
	}
	if numDots == 0 {
		return []string{pkg, ""}
	}

	prefixes := make([]string, numDots+2)
	// second pass to fill in returned slice
	for i := 0; i < len(pkg); i++ {
		if pkg[i] == '.' {
			prefixes[numDots] = pkg[:i]
			numDots--
		}
	}
	prefixes[0] = pkg

	return prefixes
}

// GetMaxTag returns the max tag number allowed, based on whether a message uses
// message set wire format or not.
func GetMaxTag(isMessageSet bool) int32 {
	if isMessageSet {
		return MaxMessageSetTag
	}
	return MaxNormalTag
}
