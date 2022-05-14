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
	opt   *option
	fdesc *protobuf.FileDescriptorProto

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
		opt:   opt,
		fdesc: protobuf.NewFileDescriptorProto(pkgname),
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
			return c.CompileEnum(msg, val)

		case isOneOf(val):
			// return c.CompileOneOf(msg, name, val)
			return msg, nil

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
			return c.compileArray(msg, val)

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
	// if description := schema.Description; description != "" {
	// 	msg.SetComments(normalizeComment(schema.Title, description))
	// }
	field := protobuf.NewFieldDescriptorProto(normalizeFieldName(schema.Title), fieldType)
	fieldMsg.AddField(field)

	return fieldMsg
}

func (c *compiler) compileArray(msg *protobuf.MessageDescriptorProto, array *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	arrayMsg := protobuf.NewMessageDescriptorProto(normalizeMessageName(array.Title))

	if ref := array.Items.Ref; ref != "" {
		refBase := path.Base(ref)
		refObj, err := c.schemasLookupFunc(refBase)
		if err != nil {
			return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
		}

		switch refObj := refObj.(type) {
		case *openapi3.Schema:
			// refMsg := c.compileSchemaRef(normalizeMessageName(refObj.Title), refObj.Properties)
			field := protobuf.NewFieldDescriptorProto(normalizeFieldName(refObj.Title), protobuf.FieldTypeMessage(nil))
			field.SetRepeated()
			arrayMsg.AddField(field)
		}

		return arrayMsg, nil
	}

	m, err := c.compileSchemaRef(normalizeMessageName(array.Title), array.Items)
	if err != nil {
		return nil, fmt.Errorf("compile array items: %w", err)
	}
	arrayMsg.AddNestedMessage(m)

	field := protobuf.NewFieldDescriptorProto(normalizeFieldName(arrayMsg.GetName()), protobuf.FieldTypeMessage(arrayMsg))
	field.SetRepeated()
	field.SetTypeName(normalizeFieldName(m.GetName()))
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

				field := protobuf.NewFieldDescriptorProto(propName, protobuf.FieldTypeMessage(refMsg))
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
		objMsg.AddNestedMessage(propMsg)

		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(propName), protobuf.FieldTypeMessage(propMsg))
		field.SetTypeName(propMsg.GetName())

		objMsg.AddField(field)
	}

	return objMsg, nil
}

func (c *compiler) CompileEnum(msg *protobuf.MessageDescriptorProto, enum *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	msgName := normalizeMessageName(enum.Title)

	eb := protobuf.NewEnumDescriptorProto(msgName)
	// if description := enum.Description; description != "" {
	// 	msg.SetComments(normalizeComment(enum.Title, description))
	// }
	for i, e := range enum.Enum {
		enumVal := protobuf.NewEnumValueDescriptorProto(msgName+"_"+strconv.Itoa(int(e.(float64))), int32(i+1))
		eb.AddValue(enumVal)
	}
	c.fdesc.AddEnum(eb)

	// field := protobuf.NewFieldDescriptorProto(normalizeMessageName(enum.Title), protobuf.FieldTypeEnum())
	// field.SetTypeName(eb.GetName())
	// msg.AddField(field)

	return msg, nil
}

// TODO(zchee): implements correctly
func (c *compiler) CompileOneOf(msg *protobuf.MessageDescriptorProto, name string, oneof *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	ob := protobuf.NewOneofDescriptorProto(normalizeFieldName(name))

	for i, ref := range oneof.OneOf {
		oneOfMsg, err := c.compileSchemaRef(name+"_"+strconv.Itoa(i), ref)
		if err != nil {
			return nil, fmt.Errorf("compile oneof ref: %w", err)
		}
		oneOfMsg.SetName(name + "_" + strconv.Itoa(i))
		msg.AddNestedMessage(oneOfMsg)
	}
	msg.AddOneof(ob)
	// msg.AddNestedMessage(oneofMsgRoot)

	return msg, nil
}

func (c *compiler) CompileAnyOf(msg *protobuf.MessageDescriptorProto, name string, anyOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	for i, ref := range anyOf.AnyOf {
		anyOfMsg, err := c.compileSchemaRef(normalizeMessageName(name+"_"+strconv.Itoa(i)), ref)
		if err != nil {
			return nil, fmt.Errorf("compile anyOf ref: %w", err)
		}
		anyOfMsg.SetName(normalizeMessageName(name + "_" + strconv.Itoa(i)))
		msg.AddNestedMessage(anyOfMsg)
		field := protobuf.NewFieldDescriptorProto(normalizeFieldName(anyOfMsg.GetName()), protobuf.FieldTypeMessage(anyOfMsg))
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
