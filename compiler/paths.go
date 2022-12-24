// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

const refPrefix = "#/components/schemas/"

// CompilePaths compiles paths object.
func (c *compiler) CompilePaths(paths openapi3.Paths) error {
	svc := protobuf.NewServiceDescriptorProto(c.opt.packageName)

	sorted := make([]string, len(paths))
	i := 0
	for path := range paths {
		sorted[i] = path
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
			methName := string(meth[0]) + strings.ToLower(meth[1:]) + name
			fmt.Printf("Name: %s\n", methName)

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

					pname = strings.TrimPrefix(pname, refPrefix)
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
						fieldType = integerFieldType(pv.Format)

					case openapi3.TypeNumber:
						fieldType = numberFieldType(pv.Format)

					case openapi3.TypeString:
						fieldType = stringFieldType(pv.Format)

					case openapi3.TypeArray:
						// TODO(zchee): handle

					case openapi3.TypeObject:
						// TODO(zchee): handle
					}

					field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(pname), fieldType)
					inputMsg.AddField(field)
				}
			}

			var val *openapi3.Schema
			// parse RequestBody
			if reqBody := op.RequestBody; reqBody != nil && reqBody.Value != nil {
				content, ok := reqBody.Value.Content["application/json"]
				if !ok {
					continue
				}

				if content.Schema.Ref != "" {
					val = content.Schema.Value
					inputType := strings.TrimPrefix(content.Schema.Ref, refPrefix)
					if !strings.HasSuffix(inputType, "Request") {
						inputType += "Request"
					}
					fmt.Printf("InputType: %s\n", inputType)
					method.InputType = proto.String(inputType)
					continue
				}

				var fieldType *descriptorpb.FieldDescriptorProto_Type
				switch content.Schema.Value.Type {
				case openapi3.TypeBoolean:
					fieldType = protobuf.FieldTypeBool()

				case openapi3.TypeInteger:
					fieldType = integerFieldType(content.Schema.Value.Format)

				case openapi3.TypeNumber:
					fieldType = numberFieldType(content.Schema.Value.Format)

				case openapi3.TypeString:
					fieldType = stringFieldType(content.Schema.Value.Format)

				case openapi3.TypeArray:
					// TODO(zchee): handle

				case openapi3.TypeObject:
					// TODO(zchee): handle
				}
				field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(content.Schema.Value.Title), fieldType)
				inputMsg.AddField(field)
			}

			// parse output type
			for status, resp := range op.Responses {
				st, _ := strconv.ParseInt(status, 10, 64)
				// TODO(zchee): handle other than 200(http.StatusOK) status
				switch st {
				case http.StatusOK:
					outputType := strings.TrimPrefix(resp.Ref, refPrefix)
					if idx := strings.LastIndex(outputType, "/"); idx > 0 {
						outputType = outputType[idx+1:]
					}
					if !strings.HasSuffix(outputType, "Response") {
						outputType += "Response"
					}
					fmt.Printf("OutputType: %s\n", outputType)
					method.OutputType = proto.String(outputType)
				}
			}

			fmt.Printf("RequestBody.Content: %#v\n\n", val)
			svc.AddMethod(method)
		}
	}

	c.fdesc.AddService(svc)

	return nil
}

// Name: GetV2NftCategories
// OutputType: GetNFTCategoriesResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: GetV2Nfts
// OutputType: GetUtilitySoundResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: GetV2Users
// OutputType: GetUserNFTsWithoutUtilityResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: GetV1Assets
// OutputType: GetNFTMetadataResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV1TokensUtilities
// InputType: UpdateReconciliationImageByManualParamRequest
// OutputType: UpdateReconciliationImageByManualResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"消込ユーティリティの画像を手動で更新する", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"userId", "tokenId", "did", "communityId"}, Properties:openapi3.Schemas{"communityId":(*openapi3.SchemaRef)(0xc0006f1980), "did":(*openapi3.SchemaRef)(0xc0006f1a10), "tokenId":(*openapi3.SchemaRef)(0xc0006f1aa0), "userId":(*openapi3.SchemaRef)(0xc0006f1b30)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2AdminNftCollections
// OutputType: GetNFTCollectionsResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PutV2AdminNftCollections
// InputType: UpdateNFTCollectionsParamRequest
// OutputType: UpdateNFTCollectionsResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"collectionId", "name", "image", "description", "isRating", "isInvisible"}, Properties:openapi3.Schemas{"collectionId":(*openapi3.SchemaRef)(0xc0006f1608), "description":(*openapi3.SchemaRef)(0xc0006f1668), "eventId":(*openapi3.SchemaRef)(0xc0006f16c8), "image":(*openapi3.SchemaRef)(0xc0006f1728), "isInvisible":(*openapi3.SchemaRef)(0xc0006f1788), "isRating":(*openapi3.SchemaRef)(0xc0006f17e8), "name":(*openapi3.SchemaRef)(0xc0006f1848)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2AdminUtilities
// OutputType: GetUtilitiesResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV2AdminUtilities
// InputType: CreateUtilityParamRequest
// OutputType: CreateUtilityResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"ユーティリティのinputモデル", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"type"}, Properties:openapi3.Schemas{"image":(*openapi3.SchemaRef)(0xc000615140), "title":(*openapi3.SchemaRef)(0xc0006151a0), "type":(*openapi3.SchemaRef)(0xc000615200), "url":(*openapi3.SchemaRef)(0xc0006152c0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2NftsPrivateMint
// InputType: PrivateMintParamRequest
// OutputType: PrivateMintResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"nftId"}, Properties:openapi3.Schemas{"nftId":(*openapi3.SchemaRef)(0xc0006f0a38)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2Users
// OutputType: GetUserNFTResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV2AdminBaseNftsBulk
// InputType: BulkCreateBaseNFTParamRequest
// OutputType: BulkCreateBaseNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"baseNFTs"}, Properties:openapi3.Schemas{"baseNFTs":(*openapi3.SchemaRef)(0xc0003c5248)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2AdminNftCategories
// InputType: CreateNFTCategoryRelationsParamRequest
// OutputType: CreateNFTCategoryRelationsResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"create relations between category and NFTs", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"baseNFTIds"}, Properties:openapi3.Schemas{"baseNFTIds":(*openapi3.SchemaRef)(0xc000614978)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2AdminPrivateContracts
// InputType: CreatePrivateContractParamRequest
// OutputType: CreatePrivateContractResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"create private contract parameter", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"contractAddress"}, Properties:openapi3.Schemas{"contractAddress":(*openapi3.SchemaRef)(0xc000615020)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2Collections
// OutputType: GetNFTCollectionsResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: GetV2AdminBaseNfts
// OutputType: GetBaseNFTsResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: GetV2BaseNfts
// OutputType: GetBaseNFTsResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV2RecoverMintNft
// InputType: RecoverMintNFTParamRequest
// OutputType: RecoverMintNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"tokenId", "contractAddress", "ownerAddress"}, Properties:openapi3.Schemas{"contractAddress":(*openapi3.SchemaRef)(0xc0006f0b70), "ownerAddress":(*openapi3.SchemaRef)(0xc0006f0c00), "tokenId":(*openapi3.SchemaRef)(0xc0006f0c90)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2Collections
// OutputType: GetNFTCollectionResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PutV2MintPrepare
// InputType: MintPrepareParamRequest
// OutputType: MintPrepareResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"tokenId"}, Properties:openapi3.Schemas{"tokenId":(*openapi3.SchemaRef)(0xc0006158c0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2SupportedNetworks
// OutputType: GetSupportedNetworksResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PutV2MintPrepareRollback
// InputType: MintPrepareParamRequest
// OutputType: MintPrepareRollbackResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"tokenId"}, Properties:openapi3.Schemas{"tokenId":(*openapi3.SchemaRef)(0xc0006158c0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2NftsSyncAll
// InputType: SyncAllNFTParamRequest
// OutputType: SyncAllNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"walletAddress"}, Properties:openapi3.Schemas{"walletAddress":(*openapi3.SchemaRef)(0xc0006f0dc8)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2Users
// OutputType: GetUserNFTsResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: GetV2Collections
// OutputType: GetNFTCollectionNFTsResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV2AdminBaseNft
// InputType: CreateBaseNFTParamRequest
// OutputType: CreateBaseNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"BaseNFT作成", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"name", "image", "backImage"}, Properties:openapi3.Schemas{"backImage":(*openapi3.SchemaRef)(0xc000614300), "conditionLink":(*openapi3.SchemaRef)(0xc000614390), "conditionText":(*openapi3.SchemaRef)(0xc000614420), "id":(*openapi3.SchemaRef)(0xc0006144b0), "image":(*openapi3.SchemaRef)(0xc000614540), "imageType":(*openapi3.SchemaRef)(0xc0006145d0), "listType":(*openapi3.SchemaRef)(0xc000614660), "name":(*openapi3.SchemaRef)(0xc0006146f0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PutV2AdminBaseNft
// InputType: UpdateBaseNFTParamRequest
// OutputType: UpdateBaseNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"BaseNFT更新", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"id"}, Properties:openapi3.Schemas{"backImage":(*openapi3.SchemaRef)(0xc0006f1230), "conditionLink":(*openapi3.SchemaRef)(0xc0006f12c0), "conditionText":(*openapi3.SchemaRef)(0xc0006f1338), "id":(*openapi3.SchemaRef)(0xc0006f13b0), "image":(*openapi3.SchemaRef)(0xc0006f1440), "name":(*openapi3.SchemaRef)(0xc0006f14d0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2AdminNftCollectionsBulk
// InputType: CreateNFTCollectionsParamRequest
// OutputType: CreateNFTCollectionsResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"collections"}, Properties:openapi3.Schemas{"collections":(*openapi3.SchemaRef)(0xc000614c78)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2AdminUtilities
// InputType: LinkUtilityParamRequest
// OutputType: LinkUtilityResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"baseNFTIds"}, Properties:openapi3.Schemas{"baseNFTIds":(*openapi3.SchemaRef)(0xc000615758)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2BaseNfts
// OutputType: GetBaseNFTResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV2AdminNftCollections
// InputType: CreateNFTCollectionItemsParamRequest
// OutputType: CreateNFTCollectionItemsResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"collectionId", "items"}, Properties:openapi3.Schemas{"collectionId":(*openapi3.SchemaRef)(0xc000614ae0), "items":(*openapi3.SchemaRef)(0xc000614b40)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2Nfts
// InputType: CreateNFTParamRequest
// OutputType: CreateNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"did", "baseNFTId"}, Properties:openapi3.Schemas{"baseNFTId":(*openapi3.SchemaRef)(0xc000614db0), "did":(*openapi3.SchemaRef)(0xc000614e40), "image":(*openapi3.SchemaRef)(0xc000614ed0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2Nfts
// OutputType: GetUtilityBookResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
//
// Name: PostV2AdminNftCategoriesBulk
// InputType: CreateNFTCategoriesParamRequest
// OutputType: CreateNFTCategoriesResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"categories"}, Properties:openapi3.Schemas{"categories":(*openapi3.SchemaRef)(0xc000614828)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: PostV2SubscribeSyncNft
// InputType: SyncNFTParamRequest
// OutputType: SyncNFTResponse
// RequestBody.Content: &openapi3.Schema{ExtensionProps:openapi3.ExtensionProps{Extensions:map[string]interface {}{}}, OneOf:openapi3.SchemaRefs(nil), AnyOf:openapi3.SchemaRefs(nil), AllOf:openapi3.SchemaRefs(nil), Not:(*openapi3.SchemaRef)(nil), Type:"object", Title:"", Format:"", Description:"", Enum:[]interface {}(nil), Default:interface {}(nil), Example:interface {}(nil), ExternalDocs:(*openapi3.ExternalDocs)(nil), UniqueItems:false, ExclusiveMin:false, ExclusiveMax:false, Nullable:false, ReadOnly:false, WriteOnly:false, AllowEmptyValue:false, Deprecated:false, XML:(*openapi3.XML)(nil), Min:(*float64)(nil), Max:(*float64)(nil), MultipleOf:(*float64)(nil), MinLength:0x0, MaxLength:(*uint64)(nil), Pattern:"", compiledPattern:(*regexp.Regexp)(nil), MinItems:0x0, MaxItems:(*uint64)(nil), Items:(*openapi3.SchemaRef)(nil), Required:[]string{"message", "subscription"}, Properties:openapi3.Schemas{"message":(*openapi3.SchemaRef)(0xc0006f0f00), "subscription":(*openapi3.SchemaRef)(0xc0006f10e0)}, MinProps:0x0, MaxProps:(*uint64)(nil), AdditionalPropertiesAllowed:(*bool)(nil), AdditionalProperties:(*openapi3.SchemaRef)(nil), Discriminator:(*openapi3.Discriminator)(nil)}
//
// Name: GetV2Users
// OutputType: GetUserNFTDiaryMessageResponse
// RequestBody.Content: (*openapi3.Schema)(nil)
