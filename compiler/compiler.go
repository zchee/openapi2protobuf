// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/openapi"
	"go.lsp.dev/openapi2protobuf/protobuf"
	"go.lsp.dev/openapi2protobuf/protobuf/printer"
)

var _ = jsonpointer.GetForToken
var _ = descriptorpb.Default_EnumOptions_Deprecated
var _ desc.Descriptor
var _ printer.Printer
var _ = prototext.Format

var (
	// UnsafeEnabled specifies whether package unsafe can be used.
	_ = protoimpl.UnsafeEnabled

	// Types used by generated code in init functions.
	_ protoimpl.DescBuilder
	_ protoimpl.TypeBuilder

	// Types used by generated code to implement EnumType, MessageType, and ExtensionType.
	_ protoimpl.EnumInfo
	_ protoimpl.MessageInfo
	_ protoimpl.ExtensionInfo

	// Types embedded in generated messages.
	_ protoimpl.MessageState
	_ protoimpl.SizeCache
	_ protoimpl.WeakFields
	_ protoimpl.UnknownFields
	_ protoimpl.ExtensionFields
	_ protoimpl.ExtensionFieldV1

	_ protoimpl.Pointer

	_ = protoimpl.X
)

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
	opt := new(option)
	for _, o := range options {
		o(opt)
	}

	pkgname := opt.packageName
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

	// lspAnyMsg := protobuf.NewMessageDescriptorProto("LSPAny")
	// field := protobuf.NewFieldDescriptorProto(normalizeFieldName(lspAnyMsg.GetName()), protobuf.FieldTypeMessage())
	// field.SetTypeName(lspAnyMsg.GetName())
	// lspAnyMsg.AddField(field)
	// lspAnyOneMsg := lspAnyMsg.Clone()
	// lspAnyOneMsg.SetName("LSPAny1")

	// compile all component objects
	// if err := c.CompileComponents(spec.Components, lspAnyMsg, lspAnyOneMsg); err != nil {
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
	fmt.Fprintln(os.Stdout, prototext.Format(fd))
	// dumpFileDescriptor(fd)
	fdesc, err := desc.CreateFileDescriptor(fd)
	if err != nil {
		return nil, fmt.Errorf("could not convert to desc: %w", err)
	}

	p := printer.Printer{}
	var sb strings.Builder
	if err := p.PrintProtoFile(fdesc, &sb); err != nil {
		return nil, fmt.Errorf("could not print proto: %w", err)
	}
	fmt.Fprintln(os.Stdout, sb.String())

	return fd, nil
}

func dumpFileDescriptor(fd *descriptorpb.FileDescriptorProto) {
	var sb strings.Builder
	sb.WriteString("\n")

	if len(fd.MessageType) > 0 {
		sb.WriteString("MessageType:\n")
		for _, msg := range fd.MessageType {
			sb.WriteString(fmt.Sprintf("%s\n", msg.GetName()))
			for _, field := range msg.GetField() {
				typeName := field.GetTypeName()
				if typeName == "" {
					typeName = field.GetType().String()
				}
				sb.WriteString(fmt.Sprintf("Field: %s -> %s\n", field.GetName(), typeName))
			}
			if len(msg.GetEnumType()) > 0 {
				for _, enum := range msg.EnumType {
					sb.WriteString(fmt.Sprintf("%s\n", enum.GetName()))
					for _, value := range enum.GetValue() {
						sb.WriteString(fmt.Sprintf("Value: %s\n", value))
					}
				}
			}
			if len(msg.GetNestedType()) > 0 {
				for _, nested := range msg.GetNestedType() {
					sb.WriteString(fmt.Sprintf("Field: %s\n", nested.GetName()))
					for _, field := range nested.GetField() {
						typeName := field.GetTypeName()
						if typeName == "" {
							typeName = field.GetType().String()
						}
						sb.WriteString(fmt.Sprintf("Nested: %s -> %s\n", field.GetName(), typeName))
					}
				}
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	if len(fd.EnumType) > 0 {
		sb.WriteString("EnumType:\n")
		for _, enum := range fd.EnumType {
			sb.WriteString(fmt.Sprintf("%s\n", enum.GetName()))
			for _, value := range enum.GetValue() {
				sb.WriteString(fmt.Sprintf("Value: %s\n", value))
			}
		}
		sb.WriteString("\n")
	}

	if len(fd.Service) > 0 {
		sb.WriteString("Service:\n")
		for _, service := range fd.Service {
			sb.WriteString(fmt.Sprintf("%#v\n", service.GetName()))
		}
		sb.WriteString("\n")
	}

	if len(fd.Dependency) > 0 {
		sb.WriteString("Dependency:\n")
		for _, dep := range fd.Dependency {
			sb.WriteString(fmt.Sprintf("%#v\n", dep))
		}
		sb.WriteString("\n")
	}

	if len(fd.Extension) > 0 {
		sb.WriteString("Extension:\n")
		for _, ext := range fd.Extension {
			sb.WriteString(fmt.Sprintf("%#v\n", ext.GetName()))
		}
		sb.WriteString("\n")
	}

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

// CompileComponents compiles all component objects.
func (c *compiler) CompileComponents(components openapi3.Components, additionalMessages ...*protobuf.MessageDescriptorProto) error {
	schemasLookupFunc := c.schemasLookupFunc
	c.schemasLookupFunc = components.Schemas.JSONLookup
	defer func() {
		c.schemasLookupFunc = schemasLookupFunc
	}()

	for name, schema := range components.Schemas {
		msg, err := c.compileSchemaRef(name, schema)
		if err != nil {
			return err
		}
		if skipMessage(msg) {
			continue
		}
		c.fdesc.AddMessage(msg)
	}

	for _, amsg := range additionalMessages {
		c.fdesc.AddMessage(amsg)
	}

	return nil
}

var enumMessage = protobuf.NewMessageDescriptorProto("enum")

func skipMessage(msg *protobuf.MessageDescriptorProto) bool {
	return msg == enumMessage || msg == nil || msg.IsEmptyField()
}

// compileSchemaRef compiles schema reference.
func (c *compiler) compileSchemaRef(name string, schemaRef *openapi3.SchemaRef) (*protobuf.MessageDescriptorProto, error) {
	if val := schemaRef.Value; val != nil {
		// Enum, OneOf, AnyOf, AllOf
		switch {
		case isEnum(val):
			enum := c.CompileEnum(name, val)
			if enum != nil && enum.GetName() != "" {
				c.fdesc.AddEnum(enum)
			}
			return enumMessage, nil

		case isOneOf(val):
			oneof, err := c.CompileOneOf(name, val)
			if err != nil {
				return nil, err
			}
			return oneof, nil

		case isAnyOf(val):
			anyof, err := c.CompileAnyOf(name, val)
			if err != nil {
				return nil, err
			}
			return anyof, nil

		case isAllOf(val):
			return nil, nil
		}

		switch val.Type {
		case openapi3.TypeBoolean:
			return c.compileBuiltin(val, protobuf.FieldTypeBool())

		case openapi3.TypeInteger:
			return c.compileBuiltin(val, IntegerFieldType(val.Format))

		case openapi3.TypeNumber:
			return c.compileBuiltin(val, NumberFieldType(val.Format))

		case openapi3.TypeString:
			return c.compileBuiltin(val, StringFieldType(val.Format))

		case openapi3.TypeArray:
			return c.compileArray(val)

		case openapi3.TypeObject:
			return c.compileObject(val)
		}
	}

	return nil, nil
}

func isEnum(schema *openapi3.Schema) bool { return schema.Enum != nil }

func isOneOf(schema *openapi3.Schema) bool { return schema.OneOf != nil }

func isAnyOf(schema *openapi3.Schema) bool { return schema.AnyOf != nil }

func isAllOf(schema *openapi3.Schema) bool { return schema.AllOf != nil }

func (c *compiler) compileBuiltin(schema *openapi3.Schema, fieldType *descriptorpb.FieldDescriptorProto_Type) (*protobuf.MessageDescriptorProto, error) {
	if fieldType == nil {
		return nil, errors.New("should fieldType is non-nil")
	}

	fieldMsg := protobuf.NewMessageDescriptorProto(normalizeMessageName(schema.Title))
	field := protobuf.NewFieldDescriptorProto(normalizeFieldName(fieldMsg.GetName()), fieldType)
	fieldMsg.AddField(field)

	return fieldMsg, nil
}

func (c *compiler) compileArray(array *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	arrayMsg := protobuf.NewMessageDescriptorProto(normalizeMessageName(array.Title))

	if ref := array.Items.Ref; ref != "" {
		refBase := path.Base(ref)
		refObj, err := c.schemasLookupFunc(refBase)
		if err != nil {
			return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
		}

		switch refObj := refObj.(type) {
		case *openapi3.Schema:
			refMsg, err := c.compileSchemaRef(normalizeMessageName(refObj.Title), array.Items)
			if err != nil {
				return nil, fmt.Errorf("compile array items: %w", err)
			}
			if skipMessage(refMsg) {
				return arrayMsg, nil
			}
			// fmt.Fprintf(os.Stderr, "%s: normalizeFieldName(refObj.Title): %s\n", unwind.FuncName(), normalizeFieldName(refObj.Title))

			fieldType := refMsg.GetFieldType()
			field := protobuf.NewFieldDescriptorProto(normalizeFieldName(refObj.Title), fieldType)
			field.SetRepeated()

			switch fieldType.Number() {
			case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
				arrayMsg.AddNestedMessage(refMsg) // add nested message only MESSAGE type
			}
			arrayMsg.AddField(field)

		default:
			fmt.Fprintf(os.Stderr, "refObj: %T: %#v\n", refObj, refObj)
		}

		return arrayMsg, nil
	}

	itemsMsg, err := c.compileSchemaRef(normalizeMessageName(arrayMsg.GetName()), array.Items)
	if err != nil {
		return nil, fmt.Errorf("compile array items: %w", err)
	}
	if skipMessage(itemsMsg) {
		return arrayMsg, nil
	}

	fieldType := itemsMsg.GetFieldType()
	field := protobuf.NewFieldDescriptorProto(normalizeFieldName(arrayMsg.GetName()), fieldType)
	field.SetNumber()
	field.SetRepeated()

	switch fieldType.Number() {
	case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
		arrayMsg.AddNestedMessage(itemsMsg) // add nested message only MESSAGE type
	}
	arrayMsg.AddField(field)

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
				if skipMessage(refMsg) {
					continue
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
		if skipMessage(propMsg) {
			continue
		}

		fieldType := propMsg.GetFieldType()
		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(propName), fieldType)
		field.SetNumber()
		if prop.Value.Type == openapi3.TypeArray {
			field.SetRepeated()
		}

		switch fieldType.Number() {
		case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
			objMsg.AddNestedMessage(propMsg) // add nested message only MESSAGE type
			field.SetTypeName(propMsg.GetName())

		default:
			field.SetTypeName(fieldType.String())
		}

		objMsg.AddField(field)
	}

	return objMsg, nil
}

// CompileEnum compiles enum objects.
func (c *compiler) CompileEnum(name string, enum *openapi3.Schema) *protobuf.EnumDescriptorProto {
	if enum.Title != "" {
		name = enum.Title
	}

	eb := protobuf.NewEnumDescriptorProto(normalizeMessageName(name))

	for i, e := range enum.Enum {
		var enumValName string
		switch e := e.(type) {
		case string:
			enumValName = normalizeMessageName(e)
		case uint64:
			enumValName = strconv.Itoa(int(e))
		case int64:
			enumValName = strconv.Itoa(int(e))
		case float64:
			enumValName = strconv.Itoa(int(e))
		}
		enumVal := protobuf.NewEnumValueDescriptorProto(eb.GetName()+"_"+enumValName, int32(i+1))
		eb.AddValue(enumVal)
	}

	return eb
}

// CompileOneOf compiles oneOf objects.
func (c *compiler) CompileOneOf(name string, oneOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if oneOf.Title != "" {
		name = oneOf.Title
	}

	msg := protobuf.NewMessageDescriptorProto(normalizeMessageName(name))
	ob := protobuf.NewOneofDescriptorProto(normalizeFieldName(name))
	msg.AddOneof(ob)

	for i, ref := range oneOf.OneOf {
		nestedMsgName := ref.Value.Title
		if nestedMsgName == "" {
			nestedMsgName = name + "_" + strconv.Itoa(i+1)
		}
		nestedMsg, err := c.compileSchemaRef(nestedMsgName, ref)
		if err != nil {
			return nil, fmt.Errorf("compile oneOf ref: %w", err)
		}
		if skipMessage(nestedMsg) {
			continue
		}

		nestedMsg.SetName(name + "_" + strconv.Itoa(i+1))
		msg.AddNestedMessage(nestedMsg)

		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(nestedMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetNumber()
		field.SetOneofIndex(msg.GetOneofIndex())
		field.SetTypeName(nestedMsg.GetName())
		msg.AddField(field)
	}

	return msg, nil
}

// CompileAnyOf compiles anyOf objects.
//
// TODO(zchee): implements correctly.
func (c *compiler) CompileAnyOf(name string, anyOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if anyOf.Title != "" {
		name = anyOf.Title
	}

	msg := protobuf.NewMessageDescriptorProto(normalizeMessageName(name))
	ob := protobuf.NewOneofDescriptorProto(normalizeFieldName(name))
	msg.AddOneof(ob)

	for i, ref := range anyOf.AnyOf {
		anyOfMsgName := ref.Value.Title
		if anyOfMsgName == "" {
			anyOfMsgName = name + "_" + strconv.Itoa(i)
		}
		anyOfMsg, err := c.compileSchemaRef(anyOfMsgName, ref)
		if err != nil {
			return nil, fmt.Errorf("compile anyOf ref: %w", err)
		}
		if skipMessage(anyOfMsg) {
			continue
		}

		// anyOfMsg.SetName(name + "_" + strconv.Itoa(i))
		// if anyOfMsg.GetName() == "" {
		// 	anyOfMsg.SetName(name + "_" + strconv.Itoa(i))
		// }
		msg.AddNestedMessage(anyOfMsg)

		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(anyOfMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetNumber()
		field.SetOneofIndex(msg.GetOneofIndex())
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
