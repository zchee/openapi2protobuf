// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package prototype

import (
	"google.golang.org/protobuf/reflect/protodesc"
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

const (
	AnyProto        = "google/protobuf/any.proto"
	DurationProto   = "google/protobuf/duration.proto"
	EmptyProto      = "google/protobuf/empty.proto"
	StructProto     = "google/protobuf/struct.proto"
	DescriptorProto = "google/protobuf/descriptor.proto"
	TimestampProto  = "google/protobuf/timestamp.proto"
	WrappersProto   = "google/protobuf/wrappers.proto"
)

var Imports = map[string]string{
	Any:           AnyProto,
	Duration:      DurationProto,
	Empty:         EmptyProto,
	ListValue:     StructProto,
	MethodOptions: DescriptorProto,
	NullValue:     StructProto,
	Struct:        StructProto,
	Timestamp:     TimestampProto,
	DoubleValue:   WrappersProto,
	FloatValue:    WrappersProto,
	Int64Value:    WrappersProto,
	UInt64Value:   WrappersProto,
	Int32Value:    WrappersProto,
	UInt32Value:   WrappersProto,
	BoolValue:     WrappersProto,
	StringValue:   WrappersProto,
	BytesValue:    WrappersProto,
}

func AnyDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(anypb.File_google_protobuf_any_proto)
}

func APIDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(apipb.File_google_protobuf_api_proto)
}

func DescriptorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(descriptorpb.File_google_protobuf_descriptor_proto)
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

var Descriptor = map[string]*descriptorpb.FileDescriptorProto{
	AnyProto:        AnyDescriptor(),
	DurationProto:   DurationDescriptor(),
	EmptyProto:      EmptyDescriptor(),
	StructProto:     StructDescriptor(),
	DescriptorProto: DescriptorDescriptor(),
	TimestampProto:  TimestampDescriptor(),
	WrappersProto:   WrappersDescriptor(),
}
