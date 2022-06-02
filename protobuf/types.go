// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/sourcecontextpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/typepb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var builtinTypes = map[string]protoreflect.Kind{
	"bytes":   protoreflect.BytesKind,
	"string":  protoreflect.StringKind,
	"integer": protoreflect.Int32Kind,
	"float":   protoreflect.FloatKind,
	"number":  protoreflect.Int64Kind,
	"boolean": protoreflect.BoolKind,
}

func AnyDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(anypb.File_google_protobuf_any_proto)
}

func APIDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(apipb.File_google_protobuf_api_proto)
}

func DurationDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(durationpb.File_google_protobuf_duration_proto)
}

func EmptyDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(emptypb.File_google_protobuf_empty_proto)
}

func FieldmaskDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(fieldmaskpb.File_google_protobuf_field_mask_proto)
}

func SourceContextDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(sourcecontextpb.File_google_protobuf_source_context_proto)
}

func StructDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(structpb.File_google_protobuf_struct_proto)
}

func TimestampDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(timestamppb.File_google_protobuf_timestamp_proto)
}

func TypeDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(typepb.File_google_protobuf_type_proto)
}

func WrappersDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(wrapperspb.File_google_protobuf_wrappers_proto)
}

const (
	Any           = "google.protobuf.Any"
	Duration      = "google.protobuf.Duration"
	Empty         = "google.protobuf.Empty"
	ListValue     = "google.protobuf.ListValue"
	MethodOptions = "google.protobuf.MethodOptions"
	NullValue     = "google.protobuf.NullValue"
	Struct        = "google.protobuf.Struct"
	Timestamp     = "google.protobuf.Timestamp"
	DoubleValue   = "google.protobuf.DoubleValue"
	FloatValue    = "google.protobuf.FloatValue"
	Int64Value    = "google.protobuf.Int64Value"
	UInt64Value   = "google.protobuf.UInt64Value"
	Int32Value    = "google.protobuf.Int32Value"
	UInt32Value   = "google.protobuf.UInt32Value"
	BoolValue     = "google.protobuf.BoolValue"
	StringValue   = "google.protobuf.StringValue"
	BytesValue    = "google.protobuf.BytesValue"
)

var KnownImports = map[string]string{
	Any:           "google/protobuf/any.proto",
	Duration:      "google/protobuf/duration.proto",
	Empty:         "google/protobuf/empty.proto",
	ListValue:     "google/protobuf/struct.proto",
	MethodOptions: "google/protobuf/descriptor.proto",
	NullValue:     "google/protobuf/struct.proto",
	Struct:        "google/protobuf/struct.proto",
	Timestamp:     "google/protobuf/timestamp.proto",
	DoubleValue:   "google/protobuf/wrappers.proto",
	FloatValue:    "google/protobuf/wrappers.proto",
	Int64Value:    "google/protobuf/wrappers.proto",
	UInt64Value:   "google/protobuf/wrappers.proto",
	Int32Value:    "google/protobuf/wrappers.proto",
	UInt32Value:   "google/protobuf/wrappers.proto",
	BoolValue:     "google/protobuf/wrappers.proto",
	StringValue:   "google/protobuf/wrappers.proto",
	BytesValue:    "google/protobuf/wrappers.proto",
}

var FieldTypes = map[string]protoreflect.Kind{
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}
