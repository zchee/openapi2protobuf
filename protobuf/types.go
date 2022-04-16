// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

var builtinTypes = map[string]protoreflect.Kind{
	"bytes":   protoreflect.BytesKind,
	"string":  protoreflect.StringKind,
	"integer": protoreflect.Int32Kind,
	"float":   protoreflect.FloatKind,
	"number":  protoreflect.Int64Kind,
	"boolean": protoreflect.BoolKind,
}

var knownImports = map[string]string{
	"google.protobuf.Any":           "google/protobuf/any.proto",
	"google.protobuf.Duration":      "google/protobuf/duration.proto",
	"google.protobuf.Empty":         "google/protobuf/empty.proto",
	"google.protobuf.ListValue":     "google/protobuf/struct.proto",
	"google.protobuf.MethodOptions": "google/protobuf/descriptor.proto",
	"google.protobuf.NullValue":     "google/protobuf/struct.proto",
	"google.protobuf.Struct":        "google/protobuf/struct.proto",
	"google.protobuf.Timestamp":     "google/protobuf/timestamp.proto",
	"google.protobuf.DoubleValue":   "google/protobuf/wrappers.proto",
	"google.protobuf.FloatValue":    "google/protobuf/wrappers.proto",
	"google.protobuf.Int64Value":    "google/protobuf/wrappers.proto",
	"google.protobuf.UInt64Value":   "google/protobuf/wrappers.proto",
	"google.protobuf.Int32Value":    "google/protobuf/wrappers.proto",
	"google.protobuf.UInt32Value":   "google/protobuf/wrappers.proto",
	"google.protobuf.BoolValue":     "google/protobuf/wrappers.proto",
	"google.protobuf.StringValue":   "google/protobuf/wrappers.proto",
	"google.protobuf.BytesValue":    "google/protobuf/wrappers.proto",
}
