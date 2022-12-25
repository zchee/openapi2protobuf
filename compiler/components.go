// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/backtrace"
	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

// CompileComponents compiles all component objects.
func (c *compiler) CompileComponents(components openapi3.Components) error {
	c.schemasLookupFunc = components.Schemas.JSONLookup
	c.parametersLookupFunc = components.Parameters.JSONLookup
	c.requestBodiesLookupFunc = components.RequestBodies.JSONLookup

	schemaNames := make([]string, len(components.Schemas))
	i := 0
	for name := range components.Schemas {
		schemaNames[i] = name
		i++
	}
	sort.Strings(schemaNames)
	for _, name := range schemaNames {
		c.fdesc.AddComponent(conv.NormalizeMessageName(name))
	}

	parameterNames := make([]string, len(components.Parameters))
	i = 0
	for name := range components.Parameters {
		parameterNames[i] = name
		i++
	}
	sort.Strings(parameterNames)
	for _, name := range parameterNames {
		c.fdesc.AddComponent(conv.NormalizeMessageName(name))
	}

	requestBodyNames := make([]string, len(components.RequestBodies))
	i = 0
	for name := range components.RequestBodies {
		requestBodyNames[i] = name
		i++
	}
	sort.Strings(requestBodyNames)
	for _, name := range requestBodyNames {
		c.fdesc.AddComponent(conv.NormalizeMessageName(name))
	}

	for _, name := range schemaNames {
		schemaRef, ok := components.Schemas[name]
		if !ok {
			continue
		}

		msg, err := c.CompileSchemaRef(name, schemaRef)
		if err != nil {
			return err
		}
		if skipMessage(msg) {
			continue
		}

		propOrder, ok := schemaRef.Value.Extensions["x-propertyOrder"].(json.RawMessage)
		if ok && propOrder != nil {
			var xPropertyOrder []string
			if err := json.Unmarshal(propOrder, &xPropertyOrder); err != nil {
				return err
			}
			msg.SortField(xPropertyOrder)
		}

		c.fdesc.AddMessage(msg)
	}

	// TODO(zchee): unnecessary? because Parameters always hasn't schema
	// for _, name := range parameterNames {
	// 	schemaRef, ok := components.Parameters[name]
	// 	if !ok {
	// 		continue
	// 	}
	//
	// 	msg, err := c.compileSchemaRef(name, schemaRef.Value.Schema)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if skipMessage(msg) {
	// 		continue
	// 	}
	//
	// 	propOrder, ok := schemaRef.Value.Extensions["x-propertyOrder"].(json.RawMessage)
	// 	if ok && propOrder != nil {
	// 		var xPropertyOrder []string
	// 		if err := json.Unmarshal(propOrder, &xPropertyOrder); err != nil {
	// 			return err
	// 		}
	// 		msg.SortField(xPropertyOrder)
	// 	}
	//
	// 	c.fdesc.AddMessage(msg)
	// }

	for _, name := range requestBodyNames {
		schemaRef, ok := components.RequestBodies[name]
		if !ok {
			continue
		}

		msg, err := c.CompileRequestBody(name, schemaRef.Value)
		if err != nil {
			return err
		}
		if skipMessage(msg) {
			continue
		}

		propOrder, ok := schemaRef.Value.Extensions["x-propertyOrder"].(json.RawMessage)
		if ok && propOrder != nil {
			var xPropertyOrder []string
			if err := json.Unmarshal(propOrder, &xPropertyOrder); err != nil {
				return err
			}
			msg.SortField(xPropertyOrder)
		}

		c.fdesc.AddMessage(msg)
	}

	return nil
}

// skipMessage reports whether the msg should skip.
func skipMessage(msg *protobuf.MessageDescriptorProto) bool {
	return msg == nil || msg.IsEmptyField()
}

// CompileSchemaRef compiles schema reference.
func (c *compiler) CompileSchemaRef(name string, schemaRef *openapi3.SchemaRef) (*protobuf.MessageDescriptorProto, error) {
	if schemaRef == nil {
		return nil, errors.New("schemaRef must be non-nil")
	}

	if val := schemaRef.Value; val != nil && val.AdditionalProperties != nil {
		additionalProps := val.AdditionalProperties
		if additionalProps != nil {
			if additionalProps.Ref == "" {
				fmt.Fprintf(os.Stderr, "%s\nadditionalProps.Value.Items: %#v\n", backtrace.FuncNameN(1), additionalProps.Value.AnyOf)
			}
			return c.CompileSchemaRef("additionalProperties", additionalProps)
		}
	}

	if val := schemaRef.Value; val != nil {
		// Enum, OneOf, AnyOf, AllOf
		switch {
		case isEnum(val):
			return c.CompileEnum(name, val), nil

		case isOneOf(val):
			return c.CompileOneof(name, val)

		case isAnyOf(val):
			return c.CompileAnyOf(name, val)

		case isAllOf(val):
			return c.CompileAllOf(name, val)
		}

		switch val.Type {
		case openapi3.TypeBoolean:
			return c.CompileBuiltin(name, val, protobuf.FieldTypeBool())

		case openapi3.TypeInteger:
			return c.CompileBuiltin(name, val, IntegerFieldType(val.Format))

		case openapi3.TypeNumber:
			return c.CompileBuiltin(name, val, NumberFieldType(val.Format))

		case openapi3.TypeString:
			return c.CompileBuiltin(name, val, StringFieldType(val.Format))

		case openapi3.TypeArray:
			return c.CompileArray(name, val)

		case openapi3.TypeObject:
			return c.CompileObject(name, val)
		}
	}

	return nil, nil
}

// isEnum reports whether the schema is enum.
func isEnum(schema *openapi3.Schema) bool { return schema.Enum != nil }

// isOneOf reports whether the schema is oneOf.
func isOneOf(schema *openapi3.Schema) bool { return schema.OneOf != nil }

// isAnyOf reports whether the schema is anyOf.
func isAnyOf(schema *openapi3.Schema) bool { return schema.AnyOf != nil }

// isAllOf reports whether the schema is allOf.
func isAllOf(schema *openapi3.Schema) bool { return schema.AllOf != nil }

func (c *compiler) CompileBuiltin(name string, schema *openapi3.Schema, fieldType *descriptorpb.FieldDescriptorProto_Type) (*protobuf.MessageDescriptorProto, error) {
	if fieldType == nil {
		return nil, errors.New("should fieldType is non-nil")
	}

	if schema.Title != "" {
		name = schema.Title
	}
	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), fieldType)
	msg.AddField(field)
	if description := schema.Description; description != "" {
		msg.AddLeadingComment(msg.GetName(), description)
	}

	return msg, nil
}

func (c *compiler) CompileArray(name string, array *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if array.Title != "" {
		name = array.Title
	}
	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))

	if ref := array.Items.Ref; ref != "" {
		refBase := path.Base(ref)

		obj := c.components.Schemas[refBase]
		if obj != nil {
			refMsg, err := c.CompileObject(refBase, obj.Value)
			if err != nil {
				return nil, fmt.Errorf("compile refObj.Items: %w", err)
			}
			if !skipMessage(refMsg) {
				c.fdesc.AddMessage(refMsg)
			}
		}

		refObj, err := c.schemasLookupFunc(refBase)
		if err != nil {
			return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
		}
		if refObj == nil {
			refObj, err = c.parametersLookupFunc(refBase)
			if err != nil {
				return nil, fmt.Errorf("%s: not found %s ref: %w", openapi3.TypeArray, ref, err)
			}
		}

		switch refObj := refObj.(type) {
		case *openapi3.Schema:
			if refObj.Items != nil {
				objMsg, err := c.CompileSchemaRef(conv.NormalizeMessageName(refBase), refObj.Items)
				if err != nil {
					return nil, fmt.Errorf("compile refObj.Items: %w", err)
				}
				c.fdesc.AddMessage(objMsg)
			}

			typename := refObj.Title
			if typename == "" {
				typename = refBase
			}
			field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(typename), protobuf.FieldTypeMessage())
			field.SetTypeName(typename)
			msg.AddField(field)
			if description := array.Description; description != "" {
				msg.AddLeadingComment(msg.GetName(), description)
			}

		default:
			fmt.Fprintf(os.Stderr, "compileArray: refObj: %T: %#v\n", refObj, refObj)
		}

		return msg, nil
	}

	itemsMsg, err := c.CompileSchemaRef(conv.NormalizeMessageName(msg.GetName()), array.Items)
	if err != nil {
		return nil, fmt.Errorf("compile array items: %w", err)
	}
	if skipMessage(itemsMsg) {
		return msg, nil
	}

	fieldType := msg.GetFieldType()
	field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(msg.GetName()), fieldType)
	field.SetTypeName(itemsMsg.GetName())

	switch protoreflect.EnumNumber(*fieldType) {
	case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
		if !c.fdesc.HasComponent(itemsMsg.GetName()) {
			msg.AddNestedMessage(itemsMsg) // add nested message only MESSAGE type
		}
		field.SetTypeName(itemsMsg.GetName())

	default:
		field.SetTypeName(fieldType.String())
	}

	if description := array.Items.Value.Description; description != "" {
		field.AddLeadingComment(field.GetName(), description)
	}
	msg.AddField(field)
	if description := array.Description; description != "" {
		msg.AddLeadingComment(msg.GetName(), description)
	}

	return msg, nil
}

func (c *compiler) CompileObject(name string, object *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
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
				refMsg, err := c.CompileSchemaRef(conv.NormalizeMessageName(refObj.Title), prop)
				if err != nil {
					return nil, fmt.Errorf("compile object items: %w", err)
				}
				if skipMessage(refMsg) {
					continue
				}

				field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(propName), protobuf.FieldTypeMessage())
				field.SetTypeName(refMsg.GetName())
				msg.AddField(field)
				if description := object.Description; description != "" {
					msg.AddLeadingComment(msg.GetName(), description)
				}

			default:
				fmt.Fprintf(os.Stderr, "compileObject: refObj: %T: %#v\n", refObj, refObj)
			}

			continue
		}

		propMsg, err := c.CompileSchemaRef(conv.NormalizeMessageName(propName), prop)
		if err != nil {
			return nil, fmt.Errorf("compile object items: %w", err)
		}
		if skipMessage(propMsg) {
			continue
		}

		fieldType := propMsg.GetFieldType()
		field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(propName), fieldType)
		if prop.Value.Type == openapi3.TypeArray {
			field.SetRepeated()
		}

		switch protoreflect.EnumNumber(*fieldType) {
		case protoreflect.EnumNumber(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE):
			if !c.fdesc.HasComponent(propMsg.GetName()) {
				msg.AddNestedMessage(propMsg) // add nested message only MESSAGE type
			}
			msg.AddNestedMessage(propMsg) // add nested message only MESSAGE type
			field.SetTypeName(propMsg.GetName())

		default:
			field.SetTypeName(fieldType.String())
		}

		if description := prop.Value.Description; description != "" {
			field.AddLeadingComment(field.GetName(), description)
		}
		msg.AddField(field)
		if description := object.Description; description != "" {
			msg.AddLeadingComment(msg.GetName(), description)
		}
	}

	return msg, nil
}

func (c *compiler) CompileRequestBody(name string, requestBody *openapi3.RequestBody) (*protobuf.MessageDescriptorProto, error) {
	content := requestBody.Content["application/json"]

	return c.CompileObject(name, content.Schema.Value)
}

// CompileEnum compiles enum objects.
func (c *compiler) CompileEnum(name string, enum *openapi3.Schema) *protobuf.MessageDescriptorProto {
	if enum.Title != "" {
		name = enum.Title
	}

	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	eb := protobuf.NewEnumDescriptorProto(conv.NormalizeMessageName(name))

	enmuPrefix := conv.NormalizeFieldName(eb.GetName())

	// add _UNSPECIFIED to first enum value
	unspecified := protobuf.NewEnumValueDescriptorProto(strings.ToUpper(enmuPrefix+"_UNSPECIFIED"), int32(0))
	eb.AddValue(unspecified)

	for i, e := range enum.Enum {
		var enumValName string
		switch e := e.(type) {
		case string:
			enumValName = conv.NormalizeFieldName(e)
		case uint64:
			enumValName = strconv.Itoa(int(e))
		case int64:
			enumValName = strconv.Itoa(int(e))
		case float64:
			enumValName = strconv.FormatFloat(float64(e), 'g', -1, 64)
		default:
			fmt.Fprintf(os.Stderr, "%s: enumValName: %T -> %s\n", backtrace.FuncName(), e, e)
		}

		enumVal := protobuf.NewEnumValueDescriptorProto(strings.ToUpper(enmuPrefix+"_"+enumValName), int32(i+1))
		eb.AddValue(enumVal)
	}

	if description := enum.Description; description != "" {
		eb.AddLeadingComment(eb.GetName(), description)
	}
	msg.AddEnumType(eb)
	if description := enum.Description; description != "" {
		msg.AddLeadingComment(msg.GetName(), description)
	}

	return msg
}

// CompileOneof compiles oneof objects.
func (c *compiler) CompileOneof(name string, oneOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	if oneOf.Title != "" {
		name = oneOf.Title
	}

	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))
	ob := protobuf.NewOneofDescriptorProto(conv.NormalizeFieldName(name))
	msg.AddOneof(ob)
	if description := oneOf.Description; description != "" {
		msg.AddLeadingComment(msg.GetName(), description)
	}

	for i, ref := range oneOf.OneOf {
		nestedMsgName := ref.Value.Title
		if nestedMsgName == "" {
			nestedMsgName = name + "_" + strconv.Itoa(i+1)
		}
		nestedMsg, err := c.CompileSchemaRef(nestedMsgName, ref)
		if err != nil {
			return nil, fmt.Errorf("compile oneof ref: %w", err)
		}
		if skipMessage(nestedMsg) {
			continue
		}

		if nestedMsg.GetName() == "" {
			nestedMsg.SetName(name + "_" + strconv.Itoa(i+1))
		}

		if !c.fdesc.HasComponent(nestedMsg.GetName()) {
			msg.AddNestedMessage(nestedMsg)
		}
		field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(nestedMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetOneofIndex(msg.GetOneofIndex())
		field.SetTypeName(nestedMsg.GetName())
		if description := ref.Value.Description; description != "" {
			field.AddLeadingComment(field.GetName(), description)
		}
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
	if description := anyOf.Description; description != "" {
		msg.AddLeadingComment(msg.GetName(), description)
	}

	for i, ref := range anyOf.AnyOf {
		anyOfMsgName := ref.Value.Title
		if anyOfMsgName == "" {
			anyOfMsgName = name + "_" + strconv.Itoa(i+1)
		}
		anyOfMsg, err := c.CompileSchemaRef(anyOfMsgName, ref)
		if err != nil {
			return nil, fmt.Errorf("compile anyOf ref: %w", err)
		}
		if skipMessage(anyOfMsg) {
			continue
		}

		if anyOfMsg.GetName() == "" {
			anyOfMsg.SetName(name + "_" + strconv.Itoa(i+1))
		}
		if !c.fdesc.HasComponent(anyOfMsg.GetName()) {
			msg.AddNestedMessage(anyOfMsg)
		}

		field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(anyOfMsg.GetName()), protobuf.FieldTypeMessage())
		field.SetOneofIndex(msg.GetOneofIndex())
		field.SetTypeName(anyOfMsg.GetName())
		if description := ref.Value.Description; description != "" {
			field.AddLeadingComment(field.GetName(), description)
		}
		msg.AddField(field)
	}

	return msg, nil
}

func (c *compiler) CompileAllOf(name string, allOf *openapi3.Schema) (*protobuf.MessageDescriptorProto, error) {
	msg := protobuf.NewMessageDescriptorProto(conv.NormalizeMessageName(name))

	for i, ref := range allOf.AllOf {
		_ = i
		_ = ref
	}

	return msg, nil
}

// IntegerFieldType returns the FieldType of the underlying type of integer from the format.
func IntegerFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "int32":
		return protobuf.FieldTypeInt32()

	case "int64":
		return protobuf.FieldTypeInt64()

	default:
		return protobuf.FieldTypeInt32()
	}
}

// NumberFieldType returns the FieldType of the underlying type of number from the format.
func NumberFieldType(format string) *descriptorpb.FieldDescriptorProto_Type {
	switch format {
	case "double":
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
