// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"net/http"
	pathpkg "path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

var queryRe = regexp.MustCompile(`/{\w+}`)

// CompilePaths compiles paths object.
func (c *compiler) CompilePaths(serviceName string, paths openapi3.Paths) error {
	svc := protobuf.NewServiceDescriptorProto(conv.NormalizeMessageName(serviceName) + "Service")

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

		// trim query template
		name = queryRe.ReplaceAllString(name, "")

		// remove all `/` separators and convert to UpperCamelCase based on that
		ss := strings.Split(name, "/")
		name = ""
		for _, s := range ss {
			name += conv.NormalizeMessageName(s)
		}

		for meth, op := range item.Operations() {
			// prepend the http method name to the RPC method name
			methName := conv.NormalizeMessageName(meth) + name

			inputMsgName := methName + "Request"
			outputMsgName := methName + "Response"

			method := protobuf.NewMethodDescriptorProto(methName, inputMsgName, outputMsgName)

			inputMsg := protobuf.NewMessageDescriptorProto(inputMsgName)

			var fieldOrder []string // for keep parameters order
			// first, check whether the op has parameters and defines proto message fields
			if params := op.Parameters; len(params) > 0 {
				for _, param := range params {
					var pname string
					var paramVal *openapi3.Parameter

					switch {
					case param.Ref != "":
						pname = pathpkg.Base(param.Ref)
						paramVal = c.components.Parameters[pname].Value
					case param.Value != nil:
						pname = pathpkg.Base(param.Value.Name)
						paramVal = param.Value
					}

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

					fieldName := conv.NormalizeFieldName(pname)
					// trim parameter in type name from field name
					fieldName = strings.ReplaceAll(fieldName, "_"+conv.NormalizeFieldName(paramVal.In), "")
					field := protobuf.NewFieldDescriptorProto(fieldName, fieldType)
					fieldOrder = append(fieldOrder, field.GetName())
					inputMsg.AddField(field)
				}
				sort.Strings(fieldOrder)
				inputMsg.SortField(fieldOrder)
			}

			// parse RequestBody for inputMsg
			if rb := op.RequestBody; rb != nil {
				var reqBody *openapi3.RequestBody
				switch {
				case rb.Value != nil:
					reqBody = rb.Value
				case rb.Ref != "":
					reqBody = c.components.RequestBodies[rb.Ref].Value
				}

				content, ok := reqBody.Content["application/json"]
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
				fieldType := protobuf.FieldTypeMessage()

				field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(fieldName), fieldType)
				field.SetTypeName(fieldName)
				fieldOrder = append(fieldOrder, field.GetName())
				inputMsg.AddField(field)

				if description := reqBody.Description; description != "" {
					inputMsg.AddLeadingComment(inputMsg.GetName(), description)
				}
			}
			inputMsg.SortField(fieldOrder)
			c.fdesc.AddMessage(inputMsg)

			// parse Responses for outputMsg
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

					var content *openapi3.MediaType
					switch {
					case resp.Value != nil:
						content = resp.Value.Content["application/json"]
					case resp.Ref != "":
						content = c.components.Responses[resp.Ref].Value.Content["application/json"]
					}

					for _, allOf := range content.Schema.Value.AllOf {
						var valsOrder []string // for keep allOf field order
						vals := make(map[string]*openapi3.Schema)

						if properties := allOf.Value.Properties; properties != nil {
							for name, prop := range properties {
								val := prop.Value
								if val == nil {
									val = c.components.Schemas[prop.Ref].Value
									name = val.Title
								}
								valsOrder = append(valsOrder, name)
								vals[name] = val
							}
						}
						sort.Strings(valsOrder)

						if allOfRef := allOf.Ref; allOfRef != "" {
							name := pathpkg.Base(allOfRef)
							val := c.components.Schemas[name].Value

							valsOrder = append(valsOrder, name)
							vals[name] = val
						}

						for _, name := range valsOrder {
							val := vals[name]

							var field *protobuf.FieldDescriptorProto
							switch val.Type {
							case openapi3.TypeBoolean:
								field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), protobuf.FieldTypeBool())
								if title := val.Title; title != "" {
									field.AddLeadingComment(field.GetName(), title)
								}

							case openapi3.TypeInteger:
								field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), IntegerFieldType(val.Format))
								if title := val.Title; title != "" {
									field.AddLeadingComment(field.GetName(), title)
								}

							case openapi3.TypeNumber:
								field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), NumberFieldType(val.Format))
								if title := val.Title; title != "" {
									field.AddLeadingComment(field.GetName(), title)
								}

							case openapi3.TypeString:
								field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), StringFieldType(val.Format))
								if title := val.Title; title != "" {
									field.AddLeadingComment(field.GetName(), title)
								}

							case openapi3.TypeArray:
								field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), protobuf.FieldTypeMessage())
								if description := val.Description; description != "" {
									field.AddLeadingComment(field.GetName(), description)
								}

							case openapi3.TypeObject:
								field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), protobuf.FieldTypeMessage())
								field.SetTypeName(name)
								if description := val.Description; description != "" {
									field.AddLeadingComment(field.GetName(), description)
								}
							}

							outputMsg.AddField(field)
						}
					}

					if description := content.Schema.Value.Title; description != "" {
						outputMsg.AddLeadingComment(outputMsg.GetName(), description)
					}
				}
			}
			c.fdesc.AddMessage(outputMsg)

			svc.AddMethod(method)
		}
	}

	c.fdesc.AddService(svc)

	return nil
}
