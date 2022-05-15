// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"context"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal"
	"go.lsp.dev/openapi2protobuf/openapi"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

var _ = jsonpointer.GetForToken
var _ = descriptorpb.Default_EnumOptions_Deprecated
var _ desc.Descriptor
var _ protoprint.Printer

// Option represents an idiomatic functional option pattern to compile the Protocol Buffers structure from the OpenAPI schema.
type Option func(o *option)

// option holds an option to compile the Protocol Buffers from the OpenAPI schema.
type option struct {
	packageName       string
	useAnnotation     bool
	skipRPC           bool
	skipDeprecatedRPC bool
	usePrefixEnum     bool
	wrapPrimitives    bool
}

// WithPackageName specifies the package name when compiling the Protocol Buffers.
func WithPackageName(packageName string) Option {
	return func(o *option) { o.packageName = packageName }
}

// WithAnnotation sets whether the add "google.api.http" annotation to the compiled Protocol Buffers.
func WithAnnotation(useAnnotation bool) Option {
	return func(o *option) { o.useAnnotation = useAnnotation }
}

// WithSkipRPC creates a new Option to specify if we should
// generate services and rpcs in addition to messages.
func WithSkipRPC(skipRPC bool) Option {
	return func(o *option) { o.skipRPC = skipRPC }
}

// WithSkipDeprecatedRpcs creates a new Option to specify if we should.
//
// Skips generating rpcs for endpoints marked as deprecated.
func WithSkipDeprecatedRPC(skipDeprecatedRPC bool) Option {
	return func(o *option) { o.skipDeprecatedRPC = skipDeprecatedRPC }
}

// WithPrefixEnums prefix enum values with their enum name to prevent protobuf namespacing issues.
func WithPrefixEnums(usePrefixEnum bool) Option {
	return func(o *option) { o.usePrefixEnum = usePrefixEnum }
}

// WithWrapPrimitives wraps primitive types with their wrapper message types.
//
// See:
// - https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
// - https://developers.google.com/protocol-buffers/docs/proto3#default
func WithWrapPrimitives(wrapPrimitives bool) Option {
	return func(o *option) { o.wrapPrimitives = wrapPrimitives }
}

type lookupFunc func(token string) (interface{}, error)

type compiler struct {
	fdesc *protobuf.FileDescriptorProto
	opt   *option

	schemasLookupFunc lookupFunc
	pathLookupFunc    lookupFunc
}

// Compile takes an OpenAPI spec and compiles it into a protobuf file descriptor.
func Compile(ctx context.Context, spec *openapi.Schema, options ...Option) (*descriptorpb.FileDescriptorProto, error) {
	spec.InternalizeRefs(ctx, openapi3.DefaultRefNameResolver)

	opt := new(option)
	for _, o := range options {
		o(opt)
	}

	pkgname := opt.packageName
	if pkgname == "" && spec.Info != nil && spec.Info.Title != "" {
		pkgname = spec.Info.Title
	}

	c := &compiler{
		fdesc: protobuf.NewFileDescriptorProto(pkgname),
		opt:   opt,
	}

	// compile info object
	if err := c.CompileInfo(spec.Info); err != nil {
		return nil, fmt.Errorf("could not compile info object: %w", err)
	}

	// compile servers object
	if err := c.CompileServers(spec.Servers); err != nil {
		return nil, fmt.Errorf("could not compile servers object: %w", err)
	}

	// compile paths object
	if err := c.CompilePaths(spec.Paths); err != nil {
		return nil, fmt.Errorf("could not compile paths object: %w", err)
	}

	// compile all component objects
	if err := c.CompileComponents(spec.Components); err != nil {
		return nil, fmt.Errorf("could not compile component objects: %w", err)
	}

	// compile security object
	if err := c.CompileSecurity(spec.Security); err != nil {
		return nil, fmt.Errorf("could not compile security object: %w", err)
	}

	// compile tags object
	if err := c.CompileTags(spec.Tags); err != nil {
		return nil, fmt.Errorf("could not compile tags object: %w", err)
	}

	// compile external documentation object
	if err := c.CompileExternalDocs(spec.ExternalDocs); err != nil {
		return nil, fmt.Errorf("could not compile external documentation object: %w", err)
	}

	fd := c.fdesc.Build()
	// dumpFileDescriptor(fd)

	fdesc, err := desc.CreateFileDescriptor(fd)
	if err != nil {
		return nil, fmt.Errorf("could not convert to desc: %w", err)
	}

	p := protoprint.Printer{}
	var sb strings.Builder
	if err := p.PrintProtoFile(fdesc, &sb); err != nil {
		return nil, fmt.Errorf("could not print proto: %w", err)
	}
	fmt.Fprintln(os.Stdout, sb.String())

	return fd, nil
}

func dumpFileDescriptor(fd *descriptorpb.FileDescriptorProto) {
	var sb strings.Builder

	sb.WriteString("Dependency:\n")
	for _, dep := range fd.Dependency {
		sb.WriteString(fmt.Sprintf("%#v\n", dep))
	}
	sb.WriteString("\n")

	sb.WriteString("MessageType:\n")
	for _, msg := range fd.MessageType {
		sb.WriteString(fmt.Sprintf("%#v\n", msg.GetName()))
		if msg.GetName() == "TextDocument" {
			sb.WriteString(spew.Sdump(msg) + "\n")
		}
	}
	sb.WriteString("\n")

	sb.WriteString("EnumType:\n")
	for _, enum := range fd.EnumType {
		sb.WriteString(fmt.Sprintf("%#v\n", enum.GetName()))
	}
	sb.WriteString("\n")

	sb.WriteString("Service:\n")
	for _, service := range fd.Service {
		sb.WriteString(fmt.Sprintf("%#v\n", service.GetName()))
	}
	sb.WriteString("\n")

	sb.WriteString("Extension:\n")
	for _, ext := range fd.Extension {
		sb.WriteString(fmt.Sprintf("%#v\n", ext.GetName()))
	}
	sb.WriteString("\n")

	fmt.Fprintln(os.Stderr, sb.String())
}

// CompileInfo compiles info object.
func (c *compiler) CompileInfo(info *openapi3.Info) error {
	if info == nil {
		return nil
	}

	if title := info.Title; title != "" {
		c.fdesc.SetPackage(normalizeFieldName(title))
	}

	if description := info.Description; description != "" {
		// c.fb.PackageComments.TrailingComment = description
	}

	if version := info.Version; version != "" {
		// c.fb.PackageComments.LeadingComment += " " + version
	}

	return nil
}

// CompileServers compiles servers object.
func (c *compiler) CompileServers(info openapi3.Servers) error { return nil }

// CompilePaths compiles paths object.
func (c *compiler) CompilePaths(paths openapi3.Paths) error { return nil }

// CompileComponents all component objects.
func (c *compiler) CompileComponents(components openapi3.Components) error {
	oldSchemasLookupFunc := c.schemasLookupFunc
	c.schemasLookupFunc = components.Schemas.JSONLookup
	defer func() {
		c.schemasLookupFunc = oldSchemasLookupFunc
	}()

	for name, schema := range components.Schemas {
		msg, err := c.compileSchemaRef(name, schema)
		if err != nil {
			return err
		}
		if len(msg.GetFields()) > 0 {
			c.fdesc.AddMessage(msg)
		}
	}

	return nil
}

// compileSchemaRef compiles schema reference.
func (c *compiler) compileSchemaRef(name string, schemaRef *openapi3.SchemaRef) (*protobuf.MessageDescriptorProto, error) {
	msg := protobuf.NewMessageDescriptorProto(name)

	if val := schemaRef.Value; val != nil {
		// Enum, OneOf, AnyOf, AllOf
		switch {
		case isEnum(val):
			return msg, c.CompileEnum(val)

		case isOneOf(val):
			return msg, c.CompileOneOf(name, val)

		case isAnyOf(val):
			// return c.CompileAnyOf(msg, name, val)
			return msg, nil

		case isAllOf(val):
			return msg, nil
		}

		switch val.Type {
		case openapi3.TypeBoolean:
			return c.compileBuiltin(val, protobuf.FieldTypeBool()), nil

		case openapi3.TypeInteger:
			return c.compileBuiltin(val, IntegerFieldType(val.Format)), nil

		case openapi3.TypeNumber:
			return c.compileBuiltin(val, NumberFieldType(val.Format)), nil

		case openapi3.TypeString:
			return c.compileBuiltin(val, StringFieldType(val.Format)), nil

		case openapi3.TypeArray:
			return c.compileArray(val)

		case openapi3.TypeObject:
			return c.compileObject(val)

		default:
			internal.Dump("val.Type", val.Type, "\nname", name)
		}
	}

	return nil, fmt.Errorf("unreachable: %s -> %s", schemaRef.Value.Type, schemaRef.Value.Title)
}

func isEnum(schema *openapi3.Schema) bool { return schema.Enum != nil }

func isOneOf(schema *openapi3.Schema) bool { return schema.OneOf != nil }

func isAnyOf(schema *openapi3.Schema) bool { return schema.AnyOf != nil }

func isAllOf(schema *openapi3.Schema) bool { return schema.AllOf != nil }

func (c *compiler) compileBuiltin(schema *openapi3.Schema, fieldType *descriptorpb.FieldDescriptorProto_Type) *protobuf.MessageDescriptorProto {
	fieldMsg := protobuf.NewMessageDescriptorProto(normalizeMessageName(schema.Title))
	field := protobuf.NewFieldDescriptorProto(normalizeFieldName(fieldMsg.GetName()), fieldType)
	fieldMsg.AddField(field)

	return fieldMsg
}

func (c *compiler) compileArray(array *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	arrayMsg := protobuf.NewMessageDescriptorProto(normalizeMessageName(array.Title))

	if ref := array.Items.Ref; ref != "" {
		fmt.Println("compileArray: in the array.Items.Ref")
		refBase := path.Base(ref)
		refObj, err := c.schemasLookupFunc(refBase)
		if err != nil {
			return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
		}

		switch refObj := refObj.(type) {
		case *openapi3.Schema:
			field := protobuf.NewFieldDescriptorProto(normalizeFieldName(refObj.Title), protobuf.FieldTypeMessage())
			field.SetRepeated()
			arrayMsg.AddField(field)
		}

		return arrayMsg, nil
	}

	items, err := c.compileSchemaRef(normalizeMessageName(arrayMsg.GetName()), array.Items)
	if err != nil {
		return nil, fmt.Errorf("compile array items: %w", err)
	}

	fieldType := items.GetFieldType()
	field := protobuf.NewFieldDescriptorProto(normalizeFieldName(arrayMsg.GetName()), fieldType)
	field.SetNumber()
	field.SetRepeated()

	switch fieldType {
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum():
		arrayMsg.AddNestedMessage(items) // add nested message only MESSAGE type
		field.SetTypeName(arrayMsg.GetName())

	default:
		field.SetTypeName(fieldType.String())
	}
	items.AddField(field)

	// fmt.Fprintf(os.Stderr, "items: %s\n", items.GetName())
	// arrayMsg.AddNestedMessage(items)

	return arrayMsg, nil
}

func (c *compiler) compileObject(object *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	objMsg := protobuf.NewMessageDescriptorProto(normalizeMessageName(object.Title))

	for propName, prop := range object.Properties {
		if ref := prop.Ref; ref != "" {
			refBase := path.Base(ref)
			refObj, err := c.schemasLookupFunc(refBase)
			if err != nil {
				return nil, fmt.Errorf("not found %s ref: %w", ref, err)
			}

			switch refObj := refObj.(type) {
			case *openapi3.Schema:
				refMsg, err := c.compileSchemaRef(normalizeMessageName(refObj.Title), prop)
				if err != nil {
					return nil, fmt.Errorf("compile object items: %w", err)
				}

				field := protobuf.NewFieldDescriptorProto(normalizeFieldName(propName), protobuf.FieldTypeMessage())
				field.SetTypeName(refMsg.GetName())
				field.SetNumber()
				objMsg.AddField(field)
			default:
				fmt.Fprintf(os.Stderr, "refObj: %T: %#v\n", refObj, refObj)
			}

			continue
		}

		propMsg, err := c.compileSchemaRef(normalizeMessageName(propName), prop)
		if err != nil {
			return nil, fmt.Errorf("compile object items: %w", err)
		}

		fieldType := propMsg.GetFieldType()
		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(propName), fieldType)
		field.SetNumber()

		switch fieldType {
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum():
			objMsg.AddNestedMessage(propMsg) // add nested message only MESSAGE type
			field.SetTypeName(propMsg.GetName())
			// fmt.Fprintf(os.Stderr, "field: %s, propMsg: %#v\n", field.GetName(), propMsg.GetName())

		default:
			fmt.Fprintf(os.Stderr, "compileObject: fieldType: %s, propMsg.GetName(): %s\n", fieldType, propMsg.GetName())
			// objMsg.AddNestedMessage(propMsg) // add nested message only MESSAGE type
			typeName := fieldType.String()

			if typeName == "TYPE_MESSAGE" { // repeated
				field2 := protobuf.NewFieldDescriptorProto(normalizeFieldName(propName), propMsg.GetFieldType())
				// fmt.Fprintf(os.Stderr, "propName: %s, field: %s, field2.GeTType\n", normalizeMessageName(propName), *field2.GetTypeName())
				field2.SetNumber()
				field2.SetTypeName(propMsg.GetName())
				propMsg.AddField(field2)

				objMsg.AddNestedMessage(propMsg)
				field.SetTypeName(propMsg.GetName())
				field.SetRepeated()
				objMsg.AddField(field)

				continue // Array
			}

			field.SetTypeName(typeName)
		}

		objMsg.AddField(field)
	}

	return objMsg, nil
}

func (c *compiler) CompileEnum(enum *openapi3.Schema) error {
	msgName := normalizeMessageName(enum.Title)

	eb := protobuf.NewEnumDescriptorProto(msgName)
	for i, e := range enum.Enum {
		enumVal := protobuf.NewEnumValueDescriptorProto(msgName+"_"+strconv.Itoa(int(e.(float64))), int32(i+1))
		eb.AddValue(enumVal)
	}
	c.fdesc.AddEnum(eb)

	return nil
}

// TODO(zchee): implements correctly
// oneof document_changes {
//   TextDocumentEdits text_document_edits = 2;
//
//   CreateFiles create_files = 3;
//
//   RenameFiles rename_files = 4;
//
//   DeleteFiles delete_files = 5;
// }
//
// message TextDocumentContentChangeEvent {
//   TextDocumentContentChangeEvent_0 text_document_content_change_event_0 = 1;
//
//   TextDocumentContentChangeEvent_1 text_document_content_change_event_1 = 2;
//
//   message TextDocumentContentChangeEvent_0 {
//     string text = 1;
//
//     Range range = 2;
//
//     int32 range_length = 3;
//   }
//
//   message TextDocumentContentChangeEvent_1 {
//     string text = 1;
//   }
// }
//
// oneof text_document_content_change_event {
//   TextDocumentContentChangeEvent_0 text_document_content_change_event_0 = 1;
//   TextDocumentContentChangeEvent_1 text_document_content_change_event_1 = 1;
//
//   message TextDocumentContentChangeEvent_0 {
//     string text = 1;
//
//     Range range = 2;
//
//     int32 range_length = 3;
//   }
//
//   message TextDocumentContentChangeEvent_1 {
//     string text = 1;
//   }
// }
func (c *compiler) CompileOneOf(name string, oneof *openapi3.Schema) error {
	msg := protobuf.NewMessageDescriptorProto(normalizeMessageName(name))

	ob := protobuf.NewOneofDescriptorProto(normalizeFieldName(name))
	for i, ref := range oneof.OneOf {
		oneOfMsg, err := c.compileSchemaRef(name+"_"+strconv.Itoa(i), ref)
		if err != nil {
			return fmt.Errorf("compile oneof ref: %w", err)
		}

		oneOfMsg.SetName(name + "_" + strconv.Itoa(i))
		msg.AddNestedMessage(oneOfMsg)

		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(oneOfMsg.GetName()), protobuf.FieldTypeMessage())
		// field.SetOneofIndex(int32(i))
		field.SetTypeName(oneOfMsg.GetName())
		msg.AddField(field)
	}
	msg.AddOneof(ob)

	// field := protobuf.NewFieldDescriptorProto(normalizeFieldName(name), protobuf.FieldTypeMessage(msg))
	// field.SetOneofIndex(int32(len(oneof.OneOf)))
	// msg.AddField(field)

	c.fdesc.AddMessage(msg)

	return nil
}

// TODO(zchee): implements correctly
func (c *compiler) CompileAnyOf(msg *protobuf.MessageDescriptorProto, name string, anyOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	for i, ref := range anyOf.AnyOf {
		anyOfMsg, err := c.compileSchemaRef(normalizeMessageName(name+"_"+strconv.Itoa(i)), ref)
		if err != nil {
			return nil, fmt.Errorf("compile anyOf ref: %w", err)
		}

		anyOfMsg.SetName(normalizeMessageName(name + "_" + strconv.Itoa(i)))
		msg.AddNestedMessage(anyOfMsg)
		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(anyOfMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetTypeName(anyOfMsg.GetName())
		msg.AddField(field)
	}

	return msg, nil
}

// CompileSecurity compiles security object.
func (c *compiler) CompileSecurity(security openapi3.SecurityRequirements) error { return nil }

// CompileTags compiles tags object.
func (c *compiler) CompileTags(tags openapi3.Tags) error { return nil }

// CompileExternalDocs compiles externalDocs object.
func (c *compiler) CompileExternalDocs(docs *openapi3.ExternalDocs) error { return nil }

// IntegerFieldType returns the FieldType of the underlying type of integer from the format.
func IntegerFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "", "int32":
		return protobuf.FieldTypeInt32()

	case "int64":
		return protobuf.FieldTypeInt64()

	default:
		return protobuf.FieldTypeInt64()
	}
}

// NumberFieldType returns the FieldType of the underlying type of number from the format.
func NumberFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "", "double":
		return protobuf.FieldTypeDouble()

	case "int64", "long":
		return protobuf.FieldTypeInt64()

	case "integer", "int32":
		return protobuf.FieldTypeInt64()

	default:
		return protobuf.FieldTypeFloat()
	}
}

// StringFieldType returns the FieldType of the underlying type of string from the format.
func StringFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "byte":
		return protobuf.FieldTypeBytes()

	default:
		return protobuf.FieldTypeString()
	}
}
