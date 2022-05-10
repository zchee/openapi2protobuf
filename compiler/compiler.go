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
	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal"
	"go.lsp.dev/openapi2protobuf/openapi"
)

var _ = jsonpointer.GetForToken

var _ = descriptorpb.Default_EnumOptions_Deprecated

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
	opt *option
	fb  *builder.FileBuilder

	schemasLookupFunc lookupFunc
	pathLookupFunc    lookupFunc
}

// Compile takes an OpenAPI spec and compiles it into a protobuf file descriptor.
func Compile(ctx context.Context, spec *openapi.Schema, options ...Option) (*desc.FileDescriptor, error) {
	spec.InternalizeRefs(ctx, openapi3.DefaultRefNameResolver)

	opt := new(option)
	for _, o := range options {
		o(opt)
	}

	pkgname := opt.packageName
	if pkgname == "" && spec.Info != nil && spec.Info.Title != "" {
		pkgname = spec.Info.Title
	}
	pkgname = strcase.ToSnake(pkgname)

	c := &compiler{
		opt: opt,
		fb:  builder.NewFile(pkgname + ".proto").SetName(pkgname).SetPackageName(pkgname),
	}
	c.fb.SetProto3(true) // forces compiling to proto3 syntax

	// compile info object
	if err := c.compileInfo(spec.Info); err != nil {
		return nil, fmt.Errorf("could not compile info object: %w", err)
	}

	// compile servers object
	if err := c.compileServers(spec.Servers); err != nil {
		return nil, fmt.Errorf("could not compile servers object: %w", err)
	}

	// compile paths object
	if err := c.compilePaths(spec.Paths); err != nil {
		return nil, fmt.Errorf("could not compile paths object: %w", err)
	}

	// compile all component objects
	if err := c.compileComponents(spec.Components); err != nil {
		return nil, fmt.Errorf("could not compile component objects: %w", err)
	}

	// compile security object
	if err := c.compileSecurity(spec.Security); err != nil {
		return nil, fmt.Errorf("could not compile security object: %w", err)
	}

	// compile tags object
	if err := c.compileTags(spec.Tags); err != nil {
		return nil, fmt.Errorf("could not compile tags object: %w", err)
	}

	// compile external documentation object
	if err := c.compileExternalDocs(spec.ExternalDocs); err != nil {
		return nil, fmt.Errorf("could not compile external documentation object: %w", err)
	}

	fdesc, err := c.fb.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build a file descriptor: %w", err)
	}

	p := protoprint.Printer{}
	var sb strings.Builder
	if err := p.PrintProtoFile(fdesc, &sb); err != nil {
		return nil, fmt.Errorf("could not print proto: %w", err)
	}
	fmt.Fprintln(os.Stdout, sb.String())

	return fdesc, nil
}

func (c *compiler) compileInfo(info *openapi3.Info) error {
	if info == nil {
		return nil
	}

	if title := info.Title; title != "" {
		c.fb.SetPackageName(strcase.ToSnake(title))
	}

	if description := info.Description; description != "" {
		c.fb.PackageComments.TrailingComment = description
	}

	if version := info.Version; version != "" {
		c.fb.PackageComments.LeadingComment += " " + version
	}

	return nil
}

func (c *compiler) compileServers(info openapi3.Servers) error { return nil }

func (c *compiler) compilePaths(paths openapi3.Paths) error { return nil }

func (c *compiler) compileComponents(components openapi3.Components) error {
	c.schemasLookupFunc = components.Schemas.JSONLookup

	for name, schema := range components.Schemas {
		msg, err := c.compileSchemaRef(name, schema)
		if err != nil {
			return err
		}
		c.fb.AddMessage(msg)
	}

	return nil
}

func (c *compiler) compileSchemaRef(name string, schemaRef *openapi3.SchemaRef) (*builder.MessageBuilder, error) {
	msg := builder.NewMessage(normalizeMessageName(name))

	if val := schemaRef.Value; val != nil {
		if isEnum(val) {
			return c.CompileEnum(msg, val)
		}

		switch val.Type {
		case openapi3.TypeBoolean:
			return c.CompileBuiltin(msg, val, builder.FieldTypeBool()), nil

		case openapi3.TypeInteger:
			return c.CompileBuiltin(msg, val, IntegerFieldType(val.Format)), nil

		case openapi3.TypeNumber:
			return c.CompileBuiltin(msg, val, NumberFieldType(val.Format)), nil

		case openapi3.TypeString:
			return c.CompileBuiltin(msg, val, StringFieldType(val.Format)), nil

		case openapi3.TypeArray:
			return c.CompileArray(msg, val)

		case openapi3.TypeObject:
			return c.CompileObject(msg, val)

		default: // OneOf, AnyOf, AllOf
			switch {
			case isOneOf(val):
				return c.CompileOneOf(msg, name, val)

			case isAnyOf(val):
				return c.CompileAnyOf(msg, name, val)

			case isAllOf(val):

			default:
				internal.Dump("val.Type", val.Type, "\nname", name)
			}
		}
	}

	return nil, errors.New("unreachable")
}

func isEnum(schema *openapi3.Schema) bool { return len(schema.Enum) > 0 }

func isOneOf(schema *openapi3.Schema) bool { return len(schema.OneOf) > 0 }

func isAnyOf(schema *openapi3.Schema) bool { return len(schema.AnyOf) > 0 }

func isAllOf(schema *openapi3.Schema) bool { return len(schema.AllOf) > 0 }

func (c *compiler) CompileBuiltin(msg *builder.MessageBuilder, schema *openapi3.Schema, fieldType *builder.FieldType) *builder.MessageBuilder {
	field := builder.NewField(normalizeFieldName(schema.Title), fieldType)
	if description := schema.Description; description != "" {
		msg.SetComments(normalizeComment(schema.Title, description))
	}
	msg.AddField(field)

	return msg
}

func (c *compiler) CompileEnum(msg *builder.MessageBuilder, enum *openapi3.Schema) (*builder.MessageBuilder, error) {
	msgName := normalizeMessageName(enum.Title)

	eb := builder.NewEnum(msgName)
	if description := enum.Description; description != "" {
		msg.SetComments(normalizeComment(enum.Title, description))
	}
	for i, e := range enum.Enum {
		enumVal := builder.NewEnumValue(msgName + "_" + strconv.Itoa(int(e.(float64))))
		enumVal.SetNumber(int32(i))
		eb.AddValue(enumVal)
	}
	msg.AddNestedEnum(eb)

	field := builder.NewField(normalizeFieldName(enum.Title), builder.FieldTypeEnum(eb))
	msg.AddField(field)

	return msg, nil
}

func (c *compiler) CompileArray(msg *builder.MessageBuilder, array *openapi3.Schema) (*builder.MessageBuilder, error) {
	if ref := array.Items.Ref; ref != "" {
		refBase := path.Base(ref)
		refObj, err := c.schemasLookupFunc(refBase)
		if err != nil {
			return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
		}

		switch refObj := refObj.(type) {
		case *openapi3.Schema:
			refMsg := builder.NewMessage(normalizeMessageName(refObj.Title))
			field := builder.NewField(normalizeFieldName(refObj.Title), builder.FieldTypeMessage(refMsg))
			msg.AddField(field)
		}

		return msg, nil
	}

	arrayMsg, err := c.compileSchemaRef(normalizeMessageName(array.Title), array.Items)
	if err != nil {
		return nil, fmt.Errorf("compile array items: %w", err)
	}
	field := builder.NewField(normalizeFieldName(array.Title), builder.FieldTypeMessage(arrayMsg))
	msg.AddField(field)

	return msg, nil
}

func (c *compiler) CompileObject(msg *builder.MessageBuilder, object *openapi3.Schema) (*builder.MessageBuilder, error) {
	for propName, prop := range object.Properties {
		if ref := prop.Ref; ref != "" {
			refBase := path.Base(ref)
			refObj, err := c.schemasLookupFunc(refBase)
			if err != nil {
				return nil, fmt.Errorf("not found %s ref: %w", ref, err)
			}

			switch refObj := refObj.(type) {
			case *openapi3.Schema:
				refMsg := builder.NewMessage(normalizeMessageName(refObj.Title))
				field := builder.NewField(normalizeFieldName(propName), builder.FieldTypeMessage(refMsg))
				msg.AddField(field)
			}
			continue
		}

		propMsg, err := c.compileSchemaRef(normalizeMessageName(propName), prop)
		if err != nil {
			return nil, fmt.Errorf("compile object items: %w", err)
		}
		if nested := msg.GetNestedMessage(normalizeMessageName(propMsg.GetName())); nested == nil {
			msg.AddNestedMessage(propMsg)
		}
		field := builder.NewField(normalizeFieldName(propName), builder.FieldTypeMessage(propMsg))
		if prop.Value.Type == openapi3.TypeArray {
			field.SetRepeated()
		}
		msg.AddField(field)
	}

	return msg, nil
}

func (c *compiler) CompileOneOf(msg *builder.MessageBuilder, name string, oneof *openapi3.Schema) (*builder.MessageBuilder, error) {
	ob := builder.NewOneOf(normalizeFieldName(name))
	for i, ref := range oneof.OneOf {
		oneOfMsg, err := c.compileSchemaRef(name+"_"+strconv.Itoa(i), ref)
		if err != nil {
			return nil, fmt.Errorf("compile oneof ref: %w", err)
		}
		oneOfMsg.SetName(name + "_" + strconv.Itoa(i))
		msg.AddNestedMessage(oneOfMsg)
		field := builder.NewField(normalizeFieldName(name+"_"+strconv.Itoa(i)), builder.FieldTypeMessage(oneOfMsg))
		ob.AddChoice(field)
	}
	msg.AddOneOf(ob)

	return msg, nil
}

// CompileAnyOf compiles the AnyOf
func (c *compiler) CompileAnyOf(msg *builder.MessageBuilder, name string, anyOf *openapi3.Schema) (*builder.MessageBuilder, error) {
	for i, ref := range anyOf.AnyOf {
		anyOfMsg, err := c.compileSchemaRef(normalizeMessageName(name+"_"+strconv.Itoa(i)), ref)
		if err != nil {
			return nil, fmt.Errorf("compile anyOf ref: %w", err)
		}
		anyOfMsg.SetName(normalizeMessageName(name + "_" + strconv.Itoa(i)))
		var field *builder.FieldBuilder
		if nested := msg.GetNestedMessage(normalizeMessageName(anyOfMsg.GetName())); nested != nil {
			field = builder.NewField(normalizeFieldName(nested.GetName()), builder.FieldTypeMessage(nested))
		} else {
			msg.AddNestedMessage(anyOfMsg)
			field = builder.NewField(normalizeFieldName(anyOfMsg.GetName()), builder.FieldTypeMessage(anyOfMsg))
		}
		msg.AddField(field)
	}

	return msg, nil
}

func (c *compiler) compileSecurity(security openapi3.SecurityRequirements) error { return nil }

func (c *compiler) compileTags(tags openapi3.Tags) error { return nil }

func (c *compiler) compileExternalDocs(docs *openapi3.ExternalDocs) error { return nil }

// IntegerFieldType returns the FieldType of the underlying type of integer from the format.
func IntegerFieldType(format string) *builder.FieldType {
	switch format {
	case "", "int32":
		return builder.FieldTypeInt32()
	case "int64":
		return builder.FieldTypeInt64()
	default:
		return builder.FieldTypeInt64()
	}
}

// NumberFieldType returns the FieldType of the underlying type of number from the format.
func NumberFieldType(format string) *builder.FieldType {
	switch format {
	case "", "double":
		return builder.FieldTypeDouble()
	case "int64", "long":
		return builder.FieldTypeInt64()
	case "integer", "int32":
		return builder.FieldTypeInt32()
	default:
		return builder.FieldTypeFloat()
	}
}

// StringFieldType returns the FieldType of the underlying type of string from the format.
func StringFieldType(format string) *builder.FieldType {
	switch format {
	case "byte":
		return builder.FieldTypeBytes()
	default:
		return builder.FieldTypeString()
	}
}
