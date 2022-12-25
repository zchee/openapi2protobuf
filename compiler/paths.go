// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"net/http"
	pathpkg "path"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

// CompilePaths compiles paths object.
func (c *compiler) CompilePaths(paths openapi3.Paths) error {
	svc := protobuf.NewServiceDescriptorProto(conv.NormalizeMessageName(c.opt.packageName) + "Service")

	sorted := make([]string, len(paths))
	i := 0
	for path := range paths {
		sorted[i] = path
		i++
	}
	sort.Strings(sorted)

	for _, path := range sorted {
		item := paths[path]
		if item == nil {
			continue
		}

		name := path // do not change the original path variable

		// cut query template, use only before word
		before, _, ok := strings.Cut(name, "/{")
		if ok {
			name = before
		}

		// remove all `/` separators and convert to UpperCamelCase based on that
		ss := strings.Split(name, "/")
		for i, s := range ss {
			ss[i] = conv.NormalizeMessageName(s)
		}
		name = strings.Join(ss, "")

		for meth, op := range item.Operations() {
			// prepend the http method name to the RPC method name
			methName := conv.NormalizeMessageName(meth) + name

			inputMsgName := methName + "Request"
			outputMsgName := methName + "Response"

			method := &descriptorpb.MethodDescriptorProto{
				Name:       proto.String(methName),
				InputType:  proto.String(inputMsgName),
				OutputType: proto.String(outputMsgName),
			}

			inputMsg := protobuf.NewMessageDescriptorProto(inputMsgName)

			// first, check whether the op has parameters and defines proto message fields
			if params := op.Parameters; len(params) > 0 {
				for _, param := range params {
					var pname string
					switch {
					case param.Ref != "":
						pname = param.Ref
					case param.Value != nil:
						pname = param.Value.Name
					}

					pname = pathpkg.Base(pname)
					p, ok := c.components.Parameters[pname]
					if !ok {
						continue
					}

					var fieldType *descriptorpb.FieldDescriptorProto_Type
					pv := p.Value.Schema.Value
					switch pv.Type {
					case openapi3.TypeBoolean:
						fieldType = protobuf.FieldTypeBool()

					case openapi3.TypeInteger:
						fieldType = IntegerFieldType(pv.Format)

					case openapi3.TypeNumber:
						fieldType = NumberFieldType(pv.Format)

					case openapi3.TypeString:
						fieldType = StringFieldType(pv.Format)
					}

					field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(pname), fieldType)
					inputMsg.AddField(field)
				}
			}

			// parse RequestBody
			if reqBody := op.RequestBody; reqBody != nil {
				var val *openapi3.RequestBody
				switch {
				case reqBody.Value != nil:
					val = reqBody.Value
				case reqBody.Ref != "":
					val = c.components.RequestBodies[reqBody.Ref].Value
				}

				content, ok := val.Content["application/json"]
				if !ok {
					continue
				}

				var fieldVal *openapi3.Schema
				switch {
				case content.Schema.Value != nil:
					fieldVal = content.Schema.Value
				case content.Schema.Ref != "":
					fieldVal = c.components.Schemas[content.Schema.Ref].Value
				}

				fieldName := fieldVal.Title
				var fieldType *descriptorpb.FieldDescriptorProto_Type
				switch content.Schema.Value.Type {
				case openapi3.TypeObject:
					fieldType = protobuf.FieldTypeMessage()
				}

				field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(fieldName), fieldType)
				field.SetTypeName(fieldName)
				inputMsg.AddField(field)

				if description := val.Description; description != "" {
					inputMsg.AddLeadingComment(inputMsg.GetName(), description)
				}
			}
			c.fdesc.AddMessage(inputMsg)

			// parse Responses
			outputMsg := protobuf.NewMessageDescriptorProto(outputMsgName)
			for status, resp := range op.Responses {
				st, _ := strconv.ParseInt(status, 10, 64)

				// TODO(zchee): handle other than 200(http.StatusOK) status
				switch st {
				case http.StatusOK:
					outputType := pathpkg.Base(resp.Ref)

					if !strings.HasSuffix(outputType, "Response") {
						outputType += "Response"
					}

					if description := *resp.Value.Description; description != "" {
						outputMsg.AddLeadingComment(outputMsg.GetName(), description)
					}

					// TODO(zchee): parses and adds resp to outputMsg's field
					// TODO(zchee): parse #/components/responses/* on components.go
				}
			}
			c.fdesc.AddMessage(outputMsg)

			svc.AddMethod(method, op.Description)
		}
	}

	c.fdesc.AddService(svc)

	return nil
}
