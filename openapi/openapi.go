// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package openapi

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

// Schema represents a root of an OpenAPI v3 document.
type Schema struct {
	*openapi3.T
}

// RootOption represents a Protocol Buffers root options.
type RootOption openapi3.ExtensionProps

// Extension represents a Protocol Buffers extension from within OpenAPI spec.
//
// The key name must be "x-extension".
type Extension struct {
	Extend string            `json:"extend" yaml:"extend"`
	Fields []*ExtensionField `json:"fields" yaml:"fields"`
}

// ExtensionField defines the field to be added to the extend message type.
type ExtensionField struct {
	Name   string `yaml:"name" json:"name"`
	Type   string `yaml:"type" json:"type"`
	Number int    `yaml:"number" json:"number"`
}

// LoadFile loads f OpenAPI file and returns the new *Schema.
func LoadFile(ctx context.Context, f string) (*Schema, error) {
	loader := &openapi3.Loader{
		IsExternalRefsAllowed: true,
		Context:               ctx,
	}
	schema, err := loader.LoadFromFile(f)
	if err != nil {
		return nil, fmt.Errorf("could not open %s file: %w", f, err)
	}

	schema.InternalizeRefs(ctx, openapi3.DefaultRefNameResolver)

	return &Schema{T: schema}, nil
}

// The following tokens are OpenAPI Specification Header Object field names.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#headerObject
const (
	TokenName            = "name"
	TokenIn              = "in"
	TokenAllowEmptyValue = "allowEmptyValue"
)

// The following tokens are OpenAPI Specification MediaType Object field names.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#mediaTypeObject
const (
	TokenEncoding = "encoding"
)

// The following tokens are OpenAPI Specification Operation Object field names.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#operation-object
const (
	TokenTags        = "tags"
	TokenSummary     = "summary"
	TokenOperationID = "operationId"
	TokenParameters  = "parameters"
	TokenRequestBody = "requestBody"
	TokenResponses   = "responses"
	TokenCallbacks   = "callbacks"
	TokenSecurity    = "security"
	TokenServers     = "servers"
)

// The following tokens are OpenAPI Specification Paramerter Object field names.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#parameter-object
const (
	TokenStyle         = "style"
	TokenExplode       = "explode"
	TokenAllowReserved = "allowReserved"
	TokenSchema        = "schema"
	TokenExamples      = "examples"
	TokenContent       = "content"
)

// The following tokens are OpenAPI Specification Reference Object field names.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#referenceObject
const (
	TokenRef = "$ref"
)

// The following tokens are taken directly from the JSON Schema definition.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#schema-object
const (
	TokenTitle            = "title"
	TokenMultipleOf       = "multipleOf"
	TokenMaximum          = "maximum"
	TokenExclusiveMaximum = "exclusiveMaximum"
	TokenMinimum          = "minimum"
	TokenExclusiveMinimum = "exclusiveMinimum"
	TokenMaxLength        = "maxLength"
	TokenMinLength        = "minLength"
	TokenPattern          = "pattern"
	TokenMaxItems         = "maxItems"
	TokenMinItems         = "minItems"
	TokenUniqueItems      = "uniqueItems"
	TokenMaxProperties    = "maxProperties"
	TokenMinProperties    = "minProperties"
	TokenRequired         = "required"
	TokenEnum             = "enum"
)

// The following tokens are taken from the JSON Schema definition but their definitions were adjusted to the OpenAPI Specification.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#schema-object
const (
	TokenType                 = "type"
	TokenAllOf                = "allOf"
	TokenOneOf                = "oneOf"
	TokenAnyOf                = "anyOf"
	TokenNot                  = "not"
	TokenItems                = "items"
	TokenProperties           = "properties"
	TokenAdditionalProperties = "additionalProperties"
	TokenDescription          = "description"
	TokenFormat               = "format"
	TokenDefault              = "default"
)

// The following tokens are JSON Schema subset fields, definitions from the OpenAPI Specification.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#schema-object
const (
	TokenNullable      = "nullable"
	TokenDiscriminator = "discriminator"
	TokenReadOnly      = "readOnly"
	TokenWriteOnly     = "writeOnly"
	TokenXml           = "xml"
	TokenExternalDocs  = "externalDocs"
	TokenExample       = "example"
	TokenDeprecated    = "deprecated"
)
