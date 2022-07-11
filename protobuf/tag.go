// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"math"
	"unicode"
	"unicode/utf8"
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
	SpecialReservedStart int32 = 19000
	// SpecialReservedEnd is the last tag in a range that is reserved and not
	// allowed for use in message definitions.
	SpecialReservedEnd int32 = 19999

	// File_packageTag is the tag number of the package element in a file
	// descriptor proto.
	File_packageTag int32 = 2
	// File_dependencyTag is the tag number of the dependencies element in a
	// file descriptor proto.
	File_dependencyTag int32 = 3
	// File_messagesTag is the tag number of the messages element in a file
	// descriptor proto.
	File_messageTypeTag int32 = 4
	// File_enumsTag is the tag number of the enums element in a file descriptor
	// proto.
	File_enumTypeTag int32 = 5
	// File_servicesTag is the tag number of the services element in a file
	// descriptor proto.
	File_servicesTag int32 = 6
	// File_extensionsTag is the tag number of the extensions element in a file
	// descriptor proto.
	File_extensionsTag int32 = 7
	// File_optionsTag is the tag number of the options element in a file
	// descriptor proto.
	File_optionsTag int32 = 8
	// File_syntaxTag is the tag number of the syntax element in a file
	// descriptor proto.
	File_syntaxTag int32 = 12

	// Message_nameTag is the tag number of the name element in a message
	// descriptor proto.
	Message_nameTag int32 = 1
	// Message_fieldsTag is the tag number of the fields element in a message
	// descriptor proto.
	Message_fieldsTag int32 = 2
	// Message_nestedMessagesTag is the tag number of the nested messages
	// element in a message descriptor proto.
	Message_nestedMessagesTag int32 = 3
	// Message_enumsTag is the tag number of the enums element in a message
	// descriptor proto.
	Message_enumsTag int32 = 4
	// Message_extensionRangeTag is the tag number of the extension ranges
	// element in a message descriptor proto.
	Message_extensionRangeTag int32 = 5
	// Message_extensionsTag is the tag number of the extensions element in a
	// message descriptor proto.
	Message_extensionsTag int32 = 6
	// Message_optionsTag is the tag number of the options element in a message
	// descriptor proto.
	Message_optionsTag int32 = 7
	// Message_oneOfsTag is the tag number of the one-ofs element in a message
	// descriptor proto.
	Message_oneOfsTag int32 = 8
	// Message_reservedRangeTag is the tag number of the reserved ranges element
	// in a message descriptor proto.
	Message_reservedRangeTag int32 = 9
	// Message_reservedNameTag is the tag number of the reserved names element
	// in a message descriptor proto.
	Message_reservedNameTag int32 = 10

	// ExtensionRange_startTag is the tag number of the start index in an
	// extension range proto.
	ExtensionRange_startTag int32 = 1
	// ExtensionRange_endTag is the tag number of the end index in an
	// extension range proto.
	ExtensionRange_endTag int32 = 2
	// ExtensionRange_optionsTag is the tag number of the options element in an
	// extension range proto.
	ExtensionRange_optionsTag int32 = 3

	// ReservedRange_startTag is the tag number of the start index in a reserved
	// range proto.
	ReservedRange_startTag int32 = 1
	// ReservedRange_endTag is the tag number of the end index in a reserved
	// range proto.
	ReservedRange_endTag int32 = 2

	// Field_nameTag is the tag number of the name element in a field descriptor
	// proto.
	Field_nameTag int32 = 1
	// Field_extendeeTag is the tag number of the extendee element in a field
	// descriptor proto.
	Field_extendeeTag int32 = 2
	// Field_numberTag is the tag number of the number element in a field
	// descriptor proto.
	Field_numberTag int32 = 3
	// Field_labelTag is the tag number of the label element in a field
	// descriptor proto.
	Field_labelTag int32 = 4
	// Field_typeTag is the tag number of the type element in a field descriptor
	// proto.
	Field_typeTag int32 = 5
	// Field_typeNameTag is the tag number of the type name element in a field
	// descriptor proto.
	Field_typeNameTag int32 = 6
	// Field_defaultTag is the tag number of the default value element in a
	// field descriptor proto.
	Field_defaultTag int32 = 7
	// Field_optionsTag is the tag number of the options element in a field
	// descriptor proto.
	Field_optionsTag int32 = 8
	// Field_jsonNameTag is the tag number of the JSON name element in a field
	// descriptor proto.
	Field_jsonNameTag int32 = 10
	// Field_proto3OptionalTag is the tag number of the proto3_optional element
	// in a descriptor proto.
	Field_proto3OptionalTag int32 = 17

	// OneOf_nameTag is the tag number of the name element in a one-of
	// descriptor proto.
	OneOf_nameTag int32 = 1
	// OneOf_optionsTag is the tag number of the options element in a one-of
	// descriptor proto.
	OneOf_optionsTag int32 = 2

	// Enum_nameTag is the tag number of the name element in an enum descriptor
	// proto.
	Enum_nameTag int32 = 1
	// Enum_valuesTag is the tag number of the values element in an enum
	// descriptor proto.
	Enum_valuesTag int32 = 2
	// Enum_optionsTag is the tag number of the options element in an enum
	// descriptor proto.
	Enum_optionsTag int32 = 3
	// Enum_reservedRangeTag is the tag number of the reserved ranges element in
	// an enum descriptor proto.
	Enum_reservedRangeTag int32 = 4
	// Enum_reservedNameTag is the tag number of the reserved names element in
	// an enum descriptor proto.
	Enum_reservedNameTag int32 = 5

	// EnumVal_nameTag is the tag number of the name element in an enum value
	// descriptor proto.
	EnumVal_nameTag int32 = 1
	// EnumVal_numberTag is the tag number of the number element in an enum
	// value descriptor proto.
	EnumVal_numberTag int32 = 2
	// EnumVal_optionsTag is the tag number of the options element in an enum
	// value descriptor proto.
	EnumVal_optionsTag int32 = 3

	// Service_nameTag is the tag number of the name element in a service
	// descriptor proto.
	Service_nameTag int32 = 1
	// Service_methodsTag is the tag number of the methods element in a service
	// descriptor proto.
	Service_methodsTag int32 = 2
	// Service_optionsTag is the tag number of the options element in a service
	// descriptor proto.
	Service_optionsTag int32 = 3

	// Method_nameTag is the tag number of the name element in a method
	// descriptor proto.
	Method_nameTag int32 = 1
	// Method_inputTag is the tag number of the input type element in a method
	// descriptor proto.
	Method_inputTag int32 = 2
	// Method_outputTag is the tag number of the output type element in a method
	// descriptor proto.
	Method_outputTag int32 = 3
	// Method_optionsTag is the tag number of the options element in a method
	// descriptor proto.
	Method_optionsTag int32 = 4
	// Method_inputStreamTag is the tag number of the input stream flag in a
	// method descriptor proto.
	Method_inputStreamTag int32 = 5
	// Method_outputStreamTag is the tag number of the output stream flag in a
	// method descriptor proto.
	Method_outputStreamTag int32 = 6

	// UninterpretedOptionsTag is the tag number of the uninterpreted options
	// element. All *Options messages use the same tag for the field that stores
	// uninterpreted options.
	UninterpretedOptionsTag = int32(999)

	// Uninterpreted_nameTag is the tag number of the name element in an
	// uninterpreted options proto.
	Uninterpreted_nameTag = 2
	// Uninterpreted_identTag is the tag number of the identifier value in an
	// uninterpreted options proto.
	Uninterpreted_identTag = 3
	// Uninterpreted_posIntTag is the tag number of the positive int value in an
	// uninterpreted options proto.
	Uninterpreted_posIntTag = 4
	// Uninterpreted_negIntTag is the tag number of the negative int value in an
	// uninterpreted options proto.
	Uninterpreted_negIntTag = 5
	// Uninterpreted_doubleTag is the tag number of the double value in an
	// uninterpreted options proto.
	Uninterpreted_doubleTag = 6
	// Uninterpreted_stringTag is the tag number of the string value in an
	// uninterpreted options proto.
	Uninterpreted_stringTag = 7
	// Uninterpreted_aggregateTag is the tag number of the aggregate value in an
	// uninterpreted options proto.
	Uninterpreted_aggregateTag = 8
	// UninterpretedName_nameTag is the tag number of the name element in an
	// uninterpreted option name proto.
	UninterpretedName_nameTag = 1
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
