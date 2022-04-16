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
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/jsonpointer"
	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
	descbuilder "github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"

	"go.lsp.dev/openapi2protobuf/openapi"
)

var _ = jsonpointer.GetForToken

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

// WithPackagePath specifies the package name when compiling the Protocol Buffers.
func WithPackagePath(packageName string) Option {
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

type compiler struct {
	opt *option
	fb  *descbuilder.FileBuilder
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
		fb:  descbuilder.NewFile(pkgname + ".proto").SetName(pkgname).SetPackageName(pkgname),
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
	for name, schema := range components.Schemas {
		msg, err := c.compileSchemaRef(name, schema, components.Schemas.JSONLookup)
		if err != nil {
			return err
		}
		if msg == nil {
			continue
		}

		c.fb.AddMessage(msg)
	}

	return nil
}

func (c *compiler) compileSecurity(security openapi3.SecurityRequirements) error { return nil }

func (c *compiler) compileTags(tags openapi3.Tags) error { return nil }

func (c *compiler) compileExternalDocs(docs *openapi3.ExternalDocs) error { return nil }

func (c *compiler) compileSchemaRef(name string, schemaRef *openapi3.SchemaRef, lookupfn func(token string) (interface{}, error)) (*descbuilder.MessageBuilder, error) {
	msg := descbuilder.NewMessage(normalizeMessageName(name))

	if val := schemaRef.Value; val != nil {
		// compile enum
		if len(val.Enum) > 0 {
			msgName := normalizeMessageName(val.Title)
			enum := descbuilder.NewEnum(msgName)
			if description := val.Description; description != "" {
				msg.SetComments(normalizeComment(val.Title, description))
			}

			for i, e := range val.Enum {
				enumVal := descbuilder.NewEnumValue(msgName + "_" + strconv.Itoa(int(e.(float64))))
				enumVal.SetNumber(int32(i))
				enum.AddValue(enumVal)
			}
			msg.AddNestedEnum(enum)
			field := descbuilder.NewField(normalizeFieldName(val.Title), descbuilder.FieldTypeEnum(enum))
			msg.AddField(field)

			return msg, nil
		}

		switch val.Type {
		case openapi3.TypeBoolean:
			field := c.compileBooleanField(val.Title)
			if description := val.Description; description != "" {
				msg.SetComments(normalizeComment(val.Title, description))
			}
			msg.AddField(field)

		case openapi3.TypeInteger:
			field := c.compileIntegerField(val.Title, val.Format)
			if description := val.Description; description != "" {
				msg.SetComments(normalizeComment(val.Title, description))
			}
			msg.AddField(field)

		case openapi3.TypeNumber:
			field := c.compileNumberField(val.Title, val.Format)
			if description := val.Description; description != "" {
				msg.SetComments(normalizeComment(val.Title, description))
			}
			msg.AddField(field)

		case openapi3.TypeString:
			field := c.compileStringField(val.Title, val.Format)
			if description := val.Description; description != "" {
				msg.SetComments(normalizeComment(val.Title, description))
			}
			msg.AddField(field)

		case openapi3.TypeArray:
			if ref := val.Items.Ref; ref != "" {
				refBase := path.Base(ref)
				refObj, err := lookupfn(refBase)
				if err != nil {
					return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
				}

				switch refObj := refObj.(type) {
				case *openapi3.Schema:
					refMsg := descbuilder.NewMessage(normalizeMessageName(refObj.Title))
					field := descbuilder.NewField(normalizeFieldName(refObj.Title), descbuilder.FieldTypeMessage(refMsg))
					msg.AddField(field)
				}

				return msg, nil
			}

			arrayMsg, err := c.compileSchemaRef(normalizeMessageName(val.Title), val.Items, lookupfn)
			if err != nil {
				return nil, fmt.Errorf("compile array items: %w", err)
			}
			field := descbuilder.NewField(normalizeFieldName(val.Title), descbuilder.FieldTypeMessage(arrayMsg))
			msg.AddField(field)

		case openapi3.TypeObject:
			for propName, prop := range val.Properties {
				if ref := prop.Ref; ref != "" {
					refBase := path.Base(ref)
					refObj, err := lookupfn(refBase)
					if err != nil {
						return nil, fmt.Errorf("not found %s ref: %w", ref, err)
					}

					switch refObj := refObj.(type) {
					case *openapi3.Schema:
						refMsg := descbuilder.NewMessage(normalizeMessageName(refObj.Title))
						field := descbuilder.NewField(normalizeFieldName(propName), descbuilder.FieldTypeMessage(refMsg))
						msg.AddField(field)
					}
					continue
				}

				propMsg, err := c.compileSchemaRef(normalizeMessageName(propName), prop, lookupfn)
				if err != nil {
					return nil, fmt.Errorf("compile object items: %w", err)
				}
				if nested := msg.GetNestedMessage(normalizeMessageName(propMsg.GetName())); nested == nil {
					msg.AddNestedMessage(propMsg)
				}
				field := descbuilder.NewField(normalizeFieldName(propName), descbuilder.FieldTypeMessage(propMsg))
				if prop.Value.Type == openapi3.TypeArray {
					field.SetRepeated()
				}
				msg.AddField(field)
			}

		default: // OneOf, AnyOf, AllOf
			switch {
			case len(val.OneOf) > 0:
				oneOf := descbuilder.NewOneOf(normalizeFieldName(name))
				for i, ref := range val.OneOf {
					oneOfMsg, err := c.compileSchemaRef(name+"_"+strconv.Itoa(i), ref, lookupfn)
					if err != nil {
						return nil, fmt.Errorf("compile oneof ref: %w", err)
					}
					oneOfMsg.SetName(name + "_" + strconv.Itoa(i))
					msg.AddNestedMessage(oneOfMsg)
					field := descbuilder.NewField(normalizeFieldName(name+"_"+strconv.Itoa(i)), descbuilder.FieldTypeMessage(oneOfMsg))
					oneOf.AddChoice(field)
				}
				msg.AddOneOf(oneOf)

			case len(val.AnyOf) > 0:
				for i, ref := range val.AnyOf {
					anyOfMsg, err := c.compileSchemaRef(normalizeMessageName(name+"_"+strconv.Itoa(i)), ref, lookupfn)
					if err != nil {
						return nil, fmt.Errorf("compile anyOf ref: %w", err)
					}
					anyOfMsg.SetName(normalizeMessageName(name + "_" + strconv.Itoa(i)))
					var field *descbuilder.FieldBuilder
					if nested := msg.GetNestedMessage(normalizeMessageName(anyOfMsg.GetName())); nested != nil {
						field = descbuilder.NewField(normalizeFieldName(nested.GetName()), descbuilder.FieldTypeMessage(nested))
					} else {
						msg.AddNestedMessage(anyOfMsg)
						field = descbuilder.NewField(normalizeFieldName(anyOfMsg.GetName()), descbuilder.FieldTypeMessage(anyOfMsg))
					}
					msg.AddField(field)
				}

			case len(val.AllOf) > 0:
			}
		}
	}

	return msg, nil
}

func normalizeComment(title, description string) descbuilder.Comments {
	var sb strings.Builder
	sb.WriteString(" ")
	sb.WriteString(normalizeMessageName(title))
	sb.WriteString(" ")
	sb.WriteByte(byte(unicode.ToLower(rune(description[0]))))
	sb.WriteString(strings.ReplaceAll(description[1:], "\n", " "))
	return descbuilder.Comments{
		LeadingComment: sb.String(),
	}
}

func (c *compiler) compileBooleanField(title string) *descbuilder.FieldBuilder {
	return descbuilder.NewField(normalizeFieldName(title), descbuilder.FieldTypeBool())
}

func (c *compiler) compileIntegerField(title, format string) *descbuilder.FieldBuilder {
	fieldType := integerFieldType(format)
	return descbuilder.NewField(normalizeFieldName(title), fieldType)
}

func (c *compiler) compileNumberField(title, format string) *descbuilder.FieldBuilder {
	fieldType := numberFieldType(format)
	return descbuilder.NewField(normalizeFieldName(title), fieldType)
}

func (c *compiler) compileStringField(title, format string) *descbuilder.FieldBuilder {
	return descbuilder.NewField(normalizeFieldName(title), stringFieldType(format))
}

func integerFieldType(format string) *descbuilder.FieldType {
	switch format {
	case "", "int32":
		return descbuilder.FieldTypeInt32()
	case "int64":
		return descbuilder.FieldTypeInt64()
	default:
		return descbuilder.FieldTypeInt64()
	}
}

func numberFieldType(format string) *descbuilder.FieldType {
	switch format {
	case "", "double":
		return descbuilder.FieldTypeDouble()
	case "int64", "long":
		return descbuilder.FieldTypeInt64()
	case "integer", "int32":
		return descbuilder.FieldTypeInt32()
	default:
		return descbuilder.FieldTypeFloat()
	}
}

func stringFieldType(format string) *descbuilder.FieldType {
	switch format {
	case "byte":
		return descbuilder.FieldTypeBytes()
	default:
		return descbuilder.FieldTypeString()
	}
}
