// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/getkin/kin-openapi/openapi3"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/backtrace"
	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

// CompileComponents compiles all component objects.
func (c *compiler) CompileComponents(components openapi3.Components) error {
	schemasLookupFunc := c.schemasLookupFunc
	c.schemasLookupFunc = components.Schemas.JSONLookup
	defer func() {
		c.schemasLookupFunc = schemasLookupFunc
	}()

	for name, schemaRef := range components.Schemas {
		msg, err := c.compileSchemaRef(name, schemaRef)
		if err != nil {
			return err
		}
		if skipMessage(msg, nil) {
			continue
		}

		rawPropertyOrder, ok := schemaRef.Value.Extensions["x-propertyOrder"].(json.RawMessage)
		if ok && rawPropertyOrder != nil {
			var xPropertyOrder []string
			if err := sonic.Unmarshal(rawPropertyOrder, &xPropertyOrder); err != nil {
				return err
			}
			msg.SortField(xPropertyOrder)
		}

		c.fdesc.AddMessage(msg)
	}

	c.fdesc.CleanupMessage()

	return nil
}

func skipMessage(msg, parent *protobuf.MessageDescriptorProto) bool {
	return msg == nil || msg.IsEmptyField()
}

// compileSchemaRef compiles schema reference.
func (c *compiler) compileSchemaRef(name string, schemaRef *openapi3.SchemaRef) (*protobuf.MessageDescriptorProto, error) {
	if additionalProps := schemaRef.Value.AdditionalProperties; additionalProps != nil {
		if additionalProps.Ref == "" {
			fmt.Fprintf(os.Stderr, "%s\nadditionalProps.Value.Items: %#v\n", backtrace.FuncNameN(1), additionalProps.Value.AnyOf)
		}
		return c.compileSchemaRef("additionalProperties", additionalProps)
	}

	if val := schemaRef.Value; val != nil {
		// Enum, OneOf, AnyOf, AllOf
		switch {
		case isEnum(val):
			return c.CompileEnum(name, val), nil

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
			return c.compileBuiltin(name, val, protobuf.FieldTypeBool())

		case openapi3.TypeInteger:
			return c.compileBuiltin(name, val, integerFieldType(val.Format))

		case openapi3.TypeNumber:
			return c.compileBuiltin(name, val, numberFieldType(val.Format))

		case openapi3.TypeString:
			return c.compileBuiltin(name, val, stringFieldType(val.Format))

		case openapi3.TypeArray:
			return c.compileArray(name, val)

		case openapi3.TypeObject:
			return c.compileObject(name, val)
		}
	}

	return nil, nil
}

func isEnum(schema *openapi3.Schema) bool { return schema.Enum != nil }

func isOneOf(schema *openapi3.Schema) bool { return schema.OneOf != nil }

func isAnyOf(schema *openapi3.Schema) bool { return schema.AnyOf != nil }

func isAllOf(schema *openapi3.Schema) bool { return schema.AllOf != nil }

func (c *compiler) compileBuiltin(name string, schema *openapi3.Schema, fieldType *descriptorpb.FieldDescriptorProto_Type) (*protobuf.MessageDescriptorProto, error) {
	if fieldType == nil {
		return nil, errors.New("should fieldType is non-nil")
	}

	if schema.Title != "" {
		name = schema.Title
	}
	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), fieldType)
	msg.AddField(field)

	return msg, nil
}

func (c *compiler) compileArray(name string, array *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if array.Title != "" {
		name = array.Title
	}
	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))

	if ref := array.Items.Ref; ref != "" {
		refBase := path.Base(ref)
		refObj, err := c.schemasLookupFunc(refBase)
		if err != nil {
			return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
		}

		switch refObj := refObj.(type) {
		case *openapi3.Schema:
			field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(refObj.Title), protobuf.FieldTypeMessage())
			field.SetNumber()
			field.SetTypeName(refObj.Title)
			msg.AddField(field)

		default:
			fmt.Fprintf(os.Stderr, "refObj: %T: %#v\n", refObj, refObj)
		}

		return msg, nil
	}

	itemsMsg, err := c.compileSchemaRef(conv.NormalizeMessageName(msg.GetName()), array.Items)
	if err != nil {
		return nil, fmt.Errorf("compile array items: %w", err)
	}
	if skipMessage(itemsMsg, msg) {
		return msg, nil
	}

	fieldType := msg.GetFieldType()
	field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(msg.GetName()), fieldType)
	field.SetNumber()
	field.SetTypeName(itemsMsg.GetName())

	switch protoreflect.EnumNumber(*fieldType) {
	case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
		msg.AddNestedMessage(itemsMsg) // add nested message only MESSAGE type
		field.SetTypeName(itemsMsg.GetName())

	default:
		field.SetTypeName(fieldType.String())
	}

	msg.AddField(field)

	return msg, nil
}

func (c *compiler) compileObject(name string, object *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if object.Title != "" {
		name = object.Title
	}
	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))

	for propName, prop := range object.Properties {
		if ref := prop.Ref; ref != "" {
			refBase := path.Base(ref)
			refObj, err := c.schemasLookupFunc(refBase)
			if err != nil {
				return nil, fmt.Errorf("not found %s ref: %w", ref, err)
			}

			switch refObj := refObj.(type) {
			case *openapi3.Schema:
				refMsg, err := c.compileSchemaRef(conv.NormalizeMessageName(refObj.Title), prop)
				if err != nil {
					return nil, fmt.Errorf("compile object items: %w", err)
				}
				if skipMessage(refMsg, msg) {
					continue
				}

				field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(propName), protobuf.FieldTypeMessage())
				field.SetTypeName(refMsg.GetName())
				field.SetNumber()
				msg.AddField(field)

			default:
				fmt.Fprintf(os.Stderr, "refObj: %T: %#v\n", refObj, refObj)
			}

			continue
		}

		propMsg, err := c.compileSchemaRef(conv.NormalizeMessageName(propName), prop)
		if err != nil {
			return nil, fmt.Errorf("compile object items: %w", err)
		}
		if skipMessage(propMsg, msg) {
			continue
		}

		fieldType := propMsg.GetFieldType()
		field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(propName), fieldType)
		field.SetNumber()
		if prop.Value.Type == openapi3.TypeArray {
			field.SetRepeated()
		}

		switch protoreflect.EnumNumber(*fieldType) {
		case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
			msg.AddNestedMessage(propMsg) // add nested message only MESSAGE type
			field.SetTypeName(propMsg.GetName())

		default:
			field.SetTypeName(fieldType.String())
		}

		msg.AddField(field)
	}

	return msg, nil
}

// CompileEnum compiles enum objects.
func (c *compiler) CompileEnum(name string, enum *openapi3.Schema) *protobuf.MessageDescriptorProto {
	if enum.Title != "" {
		name = enum.Title
	}

	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	eb := protobuf.NewEnumDescriptorProto(conv.NormalizeMessageName(name))

	for i, e := range enum.Enum {
		var enumValName string
		switch e := e.(type) {
		case string:
			enumValName = conv.NormalizeMessageName(e)
		case uint64:
			enumValName = strconv.Itoa(int(e))
		case int64:
			enumValName = strconv.Itoa(int(e))
		case float64:
			enumValName = strconv.FormatFloat(float64(e), 'g', -1, 64)
		default:
			fmt.Fprintf(os.Stderr, "%s: enumValName: %T -> %s\n", backtrace.FuncName(), e, e)
		}

		enumVal := protobuf.NewEnumValueDescriptorProto(eb.GetName()+"_"+enumValName, int32(i+1))
		eb.AddValue(enumVal)
	}

	msg.AddEnumType(eb)

	return msg
}

// CompileOneOf compiles oneOf objects.
func (c *compiler) CompileOneOf(name string, oneOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if oneOf.Title != "" {
		name = oneOf.Title
	}

	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	ob := protobuf.NewOneofDescriptorProto(conv.NormalizeFieldName(name))
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
		if skipMessage(nestedMsg, msg) {
			continue
		}

		if nestedMsg.GetName() == "" {
			nestedMsg.SetName(name + "_" + strconv.Itoa(i+1))
		}
		msg.AddNestedMessage(nestedMsg)

		field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(nestedMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetNumber()
		field.SetOneofIndex(msg.GetOneofIndex())
		field.SetTypeName(nestedMsg.GetName())
		msg.AddField(field)
	}

	return msg, nil
}

// CompileAnyOf compiles anyOf objects.
func (c *compiler) CompileAnyOf(name string, anyOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if anyOf.Title != "" {
		name = anyOf.Title
	}

	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	ob := protobuf.NewOneofDescriptorProto(conv.NormalizeFieldName(name))
	msg.AddOneof(ob)

	for i, ref := range anyOf.AnyOf {
		anyOfMsgName := ref.Value.Title
		if anyOfMsgName == "" {
			anyOfMsgName = name + "_" + strconv.Itoa(i+1)
		}
		anyOfMsg, err := c.compileSchemaRef(anyOfMsgName, ref)
		if err != nil {
			return nil, fmt.Errorf("compile anyOf ref: %w", err)
		}
		if skipMessage(anyOfMsg, msg) {
			continue
		}

		if anyOfMsg.GetName() == "" {
			anyOfMsg.SetName(name + "_" + strconv.Itoa(i+1))
		}
		msg.AddNestedMessage(anyOfMsg)

		field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(anyOfMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetNumber()
		field.SetOneofIndex(msg.GetOneofIndex())
		field.SetTypeName(anyOfMsg.GetName())
		msg.AddField(field)
	}

	return msg, nil
}

// integerFieldType returns the FieldType of the underlying type of integer from the format.
func integerFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "", "int32":
		return protobuf.FieldTypeInt32()

	case "int64":
		return protobuf.FieldTypeInt64()

	default:
		return protobuf.FieldTypeInt64()
	}
}

// numberFieldType returns the FieldType of the underlying type of number from the format.
func numberFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
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

// stringFieldType returns the FieldType of the underlying type of string from the format.
func stringFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "byte":
		return protobuf.FieldTypeBytes()

	default:
		return protobuf.FieldTypeString()
	}
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
