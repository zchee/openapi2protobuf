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

var queryRe = regexp.MustCompile(`/{(\w+)}`)

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

		// trim query template if needed
		queries := queryRe.FindStringSubmatch(name)
		name = queryRe.ReplaceAllString(name, "")

		// remove all `/` separators and convert to UpperCamelCase based on that
		ss := strings.Split(name, "/")
		name = "" // reset
		for _, s := range ss {
			name += conv.NormalizeMessageName(s)
		}
		if len(queries) > 0 {
			for i := 1; i < len(queries); i++ {
				sep := "And"
				if i == 1 {
					sep = "By"
				}
				name += sep + conv.NormalizeMessageName(queries[i])
			}
		}

		methodOrder := []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		}
		for _, meth := range methodOrder {
			op := item.GetOperation(meth)
			if op == nil {
				continue
			}

			// prepend the http method name to the RPC method name
			methName := conv.NormalizeMessageName(meth) + name

			inputMsgName := methName + "Request"
			outputMsgName := methName + "Response"

			method := protobuf.NewMethodDescriptorProto(methName, inputMsgName, outputMsgName)

			var fieldOrder []string // for keep parameters order
			inputMsg := protobuf.NewMessageDescriptorProto(inputMsgName)
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
						pname = param.Value.Name
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
					if desc := paramVal.Description; desc != "" {
						field.AddLeadingComment(field.GetName(), desc)
					}

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
				case rb.Ref != "":
					reqBody = c.components.RequestBodies[pathpkg.Base(rb.Ref)].Value
				case rb.Value != nil:
					reqBody = rb.Value
				}

				content, ok := reqBody.Content["application/json"]
				if !ok {
					continue
				}

				var fieldVal *openapi3.Schema
				switch {
				case content.Schema.Ref != "":
					fieldVal = c.components.Schemas[pathpkg.Base(content.Schema.Ref)].Value
				case content.Schema.Value != nil:
					fieldVal = content.Schema.Value
				}

				fieldName := fieldVal.Title
				fieldType := protobuf.FieldTypeMessage()

				field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(fieldName), fieldType)
				field.SetTypeName(fieldName)
				if desc := fieldVal.Description; desc != "" {
					field.AddLeadingComment(field.GetName(), desc)
				}

				fieldOrder = append(fieldOrder, field.GetName())
				inputMsg.AddField(field)
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
					var val *openapi3.Schema

					v := c.components.Schemas[pathpkg.Base(resp.Ref)]
					if v != nil {
						switch {
						case v.Ref != "":
							val = c.components.Schemas[pathpkg.Base(v.Ref)].Value
						case v.Value != nil:
							val = v.Value
						}
					}

					if val == nil {
						var content *openapi3.Response
						switch {
						case resp.Ref != "":
							content = c.components.Responses[pathpkg.Base(resp.Ref)].Value
						case resp.Value != nil:
							content = resp.Value
						}

						switch {
						case content.Content != nil:
							switch {
							case content.Content["application/json"].Schema.Ref != "":
								val = c.components.Schemas[pathpkg.Base(content.Content["application/json"].Schema.Ref)].Value
							case content.Content["application/json"].Schema.Value != nil:
								val = content.Content["application/json"].Schema.Value
							}
						default:
							continue
						}
					}

					var fieldType *descriptorpb.FieldDescriptorProto_Type
					switch val.Type {
					case openapi3.TypeBoolean:
						fieldType = protobuf.FieldTypeBool()

					case openapi3.TypeInteger:
						fieldType = IntegerFieldType(val.Format)

					case openapi3.TypeNumber:
						fieldType = NumberFieldType(val.Format)

					case openapi3.TypeString:
						fieldType = StringFieldType(val.Format)

					case openapi3.TypeArray:
						fieldType = protobuf.FieldTypeMessage()

					case openapi3.TypeObject:
						fieldType = protobuf.FieldTypeMessage()
					}
					field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(val.Title), fieldType)
					if val.Type == openapi3.TypeObject {
						field.SetTypeName(val.Title)
					}

					switch {
					case val.Description != "":
						if decs := val.Description; decs != "" {
							field.AddLeadingComment(field.GetName(), decs)
						}
					case val.Title != "":
						if title := val.Title; title != "" {
							field.AddLeadingComment(field.GetName(), title)
						}
					}
					outputMsg.AddField(field)

					allOfs := val.AllOf
					if allOfs != nil {
						for _, allOf := range allOfs {
							var valsOrder []string // for keep allOf field order
							vals := make(map[string]*openapi3.Schema)

							// parse properties
							if properties := allOf.Value.Properties; properties != nil {
								for name, prop := range properties {
									val := prop.Value
									if val == nil {
										val = c.components.Schemas[pathpkg.Base(prop.Ref)].Value
										name = val.Title
									}
									valsOrder = append(valsOrder, name)
									vals[name] = val
								}
							}
							sort.Strings(valsOrder)

							// parse allOf
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
									if desc := val.Description; desc != "" {
										field.AddLeadingComment(field.GetName(), desc)
									}

								case openapi3.TypeObject:
									field = protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(name), protobuf.FieldTypeMessage())
									field.SetTypeName(name)
									if desc := val.Description; desc != "" {
										field.AddLeadingComment(field.GetName(), desc)
									}
								}

								outputMsg.AddField(field)
							}
						}
					}

					if description := val.Title; description != "" {
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
