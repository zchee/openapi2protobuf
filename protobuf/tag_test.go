// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"math"
	"testing"
)

const (
	// MaxNormalTag is the maximum allowed tag number for a field in a normal message.
	maxNormalTag = 536870911 // 2^29 - 1

	// MaxMessageSetTag is the maximum allowed tag number of a field in a message that
	// uses the message set wire format.
	maxMessageSetTag = math.MaxInt32 - 1

	// MaxTag is the maximum allowed tag number. (It is the same as MaxMessageSetTag
	// since that is the absolute highest allowed.)
	maxTag = MaxMessageSetTag

	// SpecialReservedStart is the first tag in a range that is reserved and not
	// allowed for use in message definitions.
	specialReservedStart = 19000
	// SpecialReservedEnd is the last tag in a range that is reserved and not
	// allowed for use in message definitions.
	specialReservedEnd = 19999

	// NB: It would be nice to use constants from generated code instead of
	// hard-coding these here. But code-gen does not emit these as constants
	// anywhere. The only places they appear in generated code are struct tags
	// on fields of the generated descriptor protos.

	// File_packageTag is the tag number of the package element in a file
	// descriptor proto.
	file_packageTag = 2
	// File_dependencyTag is the tag number of the dependencies element in a
	// file descriptor proto.
	file_dependencyTag = 3
	// File_messagesTag is the tag number of the messages element in a file
	// descriptor proto.
	file_messagesTag = 4
	// File_enumsTag is the tag number of the enums element in a file descriptor
	// proto.
	file_enumsTag = 5
	// File_servicesTag is the tag number of the services element in a file
	// descriptor proto.
	file_servicesTag = 6
	// File_extensionsTag is the tag number of the extensions element in a file
	// descriptor proto.
	file_extensionsTag = 7
	// File_optionsTag is the tag number of the options element in a file
	// descriptor proto.
	file_optionsTag = 8
	// File_syntaxTag is the tag number of the syntax element in a file
	// descriptor proto.
	file_syntaxTag = 12

	// Message_nameTag is the tag number of the name element in a message
	// descriptor proto.
	message_nameTag = 1
	// Message_fieldsTag is the tag number of the fields element in a message
	// descriptor proto.
	message_fieldsTag = 2
	// Message_nestedMessagesTag is the tag number of the nested messages
	// element in a message descriptor proto.
	message_nestedMessagesTag = 3
	// Message_enumsTag is the tag number of the enums element in a message
	// descriptor proto.
	message_enumsTag = 4
	// Message_extensionRangeTag is the tag number of the extension ranges
	// element in a message descriptor proto.
	message_extensionRangeTag = 5
	// Message_extensionsTag is the tag number of the extensions element in a
	// message descriptor proto.
	message_extensionsTag = 6
	// Message_optionsTag is the tag number of the options element in a message
	// descriptor proto.
	message_optionsTag = 7
	// Message_oneOfsTag is the tag number of the one-ofs element in a message
	// descriptor proto.
	message_oneOfsTag = 8
	// Message_reservedRangeTag is the tag number of the reserved ranges element
	// in a message descriptor proto.
	message_reservedRangeTag = 9
	// Message_reservedNameTag is the tag number of the reserved names element
	// in a message descriptor proto.
	message_reservedNameTag = 10

	// ExtensionRange_startTag is the tag number of the start index in an
	// extension range proto.
	extensionRange_startTag = 1
	// ExtensionRange_endTag is the tag number of the end index in an
	// extension range proto.
	extensionRange_endTag = 2
	// ExtensionRange_optionsTag is the tag number of the options element in an
	// extension range proto.
	extensionRange_optionsTag = 3

	// ReservedRange_startTag is the tag number of the start index in a reserved
	// range proto.
	reservedRange_startTag = 1
	// ReservedRange_endTag is the tag number of the end index in a reserved
	// range proto.
	reservedRange_endTag = 2

	// Field_nameTag is the tag number of the name element in a field descriptor
	// proto.
	field_nameTag = 1
	// Field_extendeeTag is the tag number of the extendee element in a field
	// descriptor proto.
	field_extendeeTag = 2
	// Field_numberTag is the tag number of the number element in a field
	// descriptor proto.
	field_numberTag = 3
	// Field_labelTag is the tag number of the label element in a field
	// descriptor proto.
	field_labelTag = 4
	// Field_typeTag is the tag number of the type element in a field descriptor
	// proto.
	field_typeTag = 5
	// Field_typeNameTag is the tag number of the type name element in a field
	// descriptor proto.
	field_typeNameTag = 6
	// Field_defaultTag is the tag number of the default value element in a
	// field descriptor proto.
	field_defaultTag = 7
	// Field_optionsTag is the tag number of the options element in a field
	// descriptor proto.
	field_optionsTag = 8
	// Field_jsonNameTag is the tag number of the JSON name element in a field
	// descriptor proto.
	field_jsonNameTag = 10
	// Field_proto3OptionalTag is the tag number of the proto3_optional element
	// in a descriptor proto.
	field_proto3OptionalTag = 17

	// OneOf_nameTag is the tag number of the name element in a one-of
	// descriptor proto.
	oneOf_nameTag = 1
	// OneOf_optionsTag is the tag number of the options element in a one-of
	// descriptor proto.
	oneOf_optionsTag = 2
	// Enum_nameTag is the tag number of the name element in an enum descriptor
	// proto.
	enum_nameTag = 1
	// Enum_valuesTag is the tag number of the values element in an enum
	// descriptor proto.
	enum_valuesTag = 2
	// Enum_optionsTag is the tag number of the options element in an enum
	// descriptor proto.
	enum_optionsTag = 3
	// Enum_reservedRangeTag is the tag number of the reserved ranges element in
	// an enum descriptor proto.
	enum_reservedRangeTag = 4
	// Enum_reservedNameTag is the tag number of the reserved names element in
	// an enum descriptor proto.
	enum_reservedNameTag = 5
	// EnumVal_nameTag is the tag number of the name element in an enum value
	// descriptor proto.
	enumVal_nameTag = 1
	// EnumVal_numberTag is the tag number of the number element in an enum
	// value descriptor proto.
	enumVal_numberTag = 2
	// EnumVal_optionsTag is the tag number of the options element in an enum
	// value descriptor proto.
	enumVal_optionsTag = 3
	// Service_nameTag is the tag number of the name element in a service
	// descriptor proto.
	service_nameTag = 1
	// Service_methodsTag is the tag number of the methods element in a service
	// descriptor proto.
	service_methodsTag = 2
	// Service_optionsTag is the tag number of the options element in a service
	// descriptor proto.
	service_optionsTag = 3
	// Method_nameTag is the tag number of the name element in a method
	// descriptor proto.
	method_nameTag = 1
	// Method_inputTag is the tag number of the input type element in a method
	// descriptor proto.
	method_inputTag = 2
	// Method_outputTag is the tag number of the output type element in a method
	// descriptor proto.
	method_outputTag = 3
	// Method_optionsTag is the tag number of the options element in a method
	// descriptor proto.
	method_optionsTag = 4
	// Method_inputStreamTag is the tag number of the input stream flag in a
	// method descriptor proto.
	method_inputStreamTag = 5
	// Method_outputStreamTag is the tag number of the output stream flag in a
	// method descriptor proto.
	method_outputStreamTag = 6

	// UninterpretedOptionsTag is the tag number of the uninterpreted options
	// element. All *Options messages use the same tag for the field that stores
	// uninterpreted options.
	uninterpretedOptionsTag = 999

	// Uninterpreted_nameTag is the tag number of the name element in an
	// uninterpreted options proto.
	uninterpreted_nameTag = 2
	// Uninterpreted_identTag is the tag number of the identifier value in an
	// uninterpreted options proto.
	uninterpreted_identTag = 3
	// Uninterpreted_posIntTag is the tag number of the positive int value in an
	// uninterpreted options proto.
	uninterpreted_posIntTag = 4
	// Uninterpreted_negIntTag is the tag number of the negative int value in an
	// uninterpreted options proto.
	uninterpreted_negIntTag = 5
	// Uninterpreted_doubleTag is the tag number of the double value in an
	// uninterpreted options proto.
	uninterpreted_doubleTag = 6
	// Uninterpreted_stringTag is the tag number of the string value in an
	// uninterpreted options proto.
	uninterpreted_stringTag = 7
	// Uninterpreted_aggregateTag is the tag number of the aggregate value in an
	// uninterpreted options proto.
	uninterpreted_aggregateTag = 8
	// UninterpretedName_nameTag is the tag number of the name element in an
	// uninterpreted option name proto.
	uninterpretedName_nameTag = 1
)

func TestMaxNormalTag(t *testing.T) {
	tag := File_packageTag
	if tag != file_packageTag {
		t.Fatalf("not equal to tags: file_packageTag: %d, File_packageTag: %d", file_packageTag, File_packageTag)
	}
}
