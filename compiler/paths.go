// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	pathpkg "path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/protobuf"
)

// CompilePaths compiles paths object.
func (c *compiler) CompilePaths(paths openapi3.Paths) error {
	svc := protobuf.NewServiceDescriptorProto(c.opt.packageName)

	for path, item := range paths {
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
			// fmt.Printf("Name: %s\n", methName)

			inputMsgName := methName + "Request"
			outputMsgName := methName + "Response"

			method := &descriptorpb.MethodDescriptorProto{
				Name:       proto.String(methName),
				InputType:  proto.String(inputMsgName),
				OutputType: proto.String(outputMsgName),
			}

			// parse RequestBody
			var inputMsg *protobuf.MessageDescriptorProto
			if body := op.RequestBody; body != nil && body.Value != nil {
				// fmt.Printf("body: %#v\n", body.Value)
				body, ok := body.Value.Content["application/json"]
				if !ok {
					continue
				}

				if body.Schema != nil {
					var err error
					inputMsg, err = c.compileSchemaRef(inputMsgName, body.Schema)
					if err != nil {
						return err
					}
					// fmt.Printf("inputMsg: %#v\n", inputMsg)
				}
			}

			// first, check whether the op has parameters and defines proto message fields
			if params := op.Parameters; len(params) > 0 {
				for _, param := range params {
					var pname string
					switch {
					case param.Ref != "":
						pname = pathpkg.Base(param.Ref)
					case param.Value != nil:
						pname = param.Value.Name
					}
					// fmt.Printf("pname: %s, %t\n", pname, pname == "NFTCategoryInputModel")

					p, ok := c.components.Parameters[pname]
					if !ok {
						continue
					}

					switch p.Value.In {
					case openapi3.ParameterInPath:
					case openapi3.ParameterInQuery:
					case openapi3.ParameterInHeader:
					case openapi3.ParameterInCookie:
						continue // skip
					}

					if inputMsg == nil {
						inputMsg = protobuf.NewMessageDescriptorProto(inputMsgName)
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
						array, err := c.compileArray(name, pv)
						if err != nil {
							return err
						}
						inputMsg.AddNestedMessage(array)
						continue
					}

					field := protobuf.NewFieldDescriptorProto(conv.NormalizeFieldName(pname), fieldType)
					if inputMsg == nil {
						inputMsg = protobuf.NewMessageDescriptorProto(inputMsgName)
					}
					inputMsg.AddField(field)
				}
			}
			c.fdesc.AddMessage(inputMsg)

			// parse output type
			// for status, resp := range op.Responses {
			// 	// outputMsg := protobuf.NewMessageDescriptorProto(outputMsgName)
			//
			// 	st, _ := strconv.ParseInt(status, 10, 64)
			// 	// TODO(zchee): handle other than 200(http.StatusOK) status
			// 	switch st {
			// 	case http.StatusOK:
			// 		content, ok := c.components.Responses[resp.Ref]
			// 		if !ok {
			// 			continue
			// 		}
			// 		body, ok := content.Value.Content["application/json"]
			// 		if !ok {
			// 			continue
			// 		}
			//
			// 		outputMsg, err := c.compileSchemaRef(outputMsgName, body.Schema)
			// 		if err != nil {
			// 			continue
			// 		}
			// 		c.fdesc.AddMessage(outputMsg)
			// 	}
			// }

			svc.AddMethod(method)
		}
	}

	c.fdesc.AddService(svc)

	return nil
}

// Name: GetV2Users
// Name: PostV2AdminNftCollectionsBulk
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00060cc00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"collections":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00060b680)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Collections":true}}
// Name: PostV2AdminNftCollections
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00060ce00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"collectionid":true, "items":true}, fieldNumber:2, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00060bb80), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc00060be00)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Items":true}}
// Name: PostV2AdminPrivateContracts
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00060d200), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:" Postv2Adminprivatecontractsrequest create private contract parameter.", TrailingComments:""}, field:map[string]bool{"contractaddress":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000738400)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: GetV2NftCategories
// Name: PostV2Nfts
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00060d500), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"basenftid":true, "did":true, "image":true}, fieldNumber:3, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000738c00), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc000738e80), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc000739100)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: GetV2Nfts
// Name: GetV2Users
// Name: PostV2AdminBaseNftsBulk
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00060db00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"basenfts":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000573480)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Basenfts":true}}
// Name: PostV2AdminNftCategories
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000744600), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:" Postv2Adminnftcategoriesrequest create relations between category and NFTs.", TrailingComments:""}, field:map[string]bool{"basenftids":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000573c00)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Basenftids":true}}
// Name: GetV2AdminNftCollections
// Name: PutV2AdminNftCollections
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000744b00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"collectionid":true, "description":true, "eventid":true, "image":true, "isinvisible":true, "israting":true, "name":true}, fieldNumber:7, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc0000f3400), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc0000f3800), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc0000f3a80), 3:(*descriptorpb.SourceCodeInfo_Location)(0xc0000f3d00), 4:(*descriptorpb.SourceCodeInfo_Location)(0xc0000f3f80), 5:(*descriptorpb.SourceCodeInfo_Location)(0xc00057e200), 6:(*descriptorpb.SourceCodeInfo_Location)(0xc00057e480)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: GetV2Collections
// Name: PutV2MintPrepareRollback
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000745400), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"tokenid":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00057eb80)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: PostV2NftsSyncAll
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000745600), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"walletaddress":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00057ef80)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: GetV2Users
// Name: GetV1Assets
// Name: PostV1TokensUtilities
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000745c00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:" Postv1Tokensutilitiesrequest 消込ユーティリティの画像を手動で更新する.", TrailingComments:""}, field:map[string]bool{"communityid":true, "did":true, "tokenid":true, "userid":true}, fieldNumber:4, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00057fe00), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc000582080), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc000582300), 3:(*descriptorpb.SourceCodeInfo_Location)(0xc000582580)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: GetV2Collections
// Name: PostV2SubscribeSyncNft
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc0000f1800), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"message":true, "subscription":true}, fieldNumber:2, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000582e00), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc000583480)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Message":true}}
// Name: GetV2Users
// Name: PostV2AdminBaseNft
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000586200), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:" Postv2Adminbasenftrequest baseNFT作成.", TrailingComments:""}, field:map[string]bool{"backimage":true, "conditionlink":true, "conditiontext":true, "id":true, "image":true, "imagetype":true, "listtype":true, "name":true}, fieldNumber:8, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000583c00), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc000583e80), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc00058a100), 3:(*descriptorpb.SourceCodeInfo_Location)(0xc00058a380), 4:(*descriptorpb.SourceCodeInfo_Location)(0xc00058a600), 5:(*descriptorpb.SourceCodeInfo_Location)(0xc00058a880), 6:(*descriptorpb.SourceCodeInfo_Location)(0xc00058ab00), 7:(*descriptorpb.SourceCodeInfo_Location)(0xc00058ad80)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: PutV2AdminBaseNft
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000586b00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:" Putv2Adminbasenftrequest baseNFT更新.", TrailingComments:""}, field:map[string]bool{"backimage":true, "conditionlink":true, "conditiontext":true, "id":true, "image":true, "name":true}, fieldNumber:6, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00058b300), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc00058b580), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc00058b800), 3:(*descriptorpb.SourceCodeInfo_Location)(0xc00058ba80), 4:(*descriptorpb.SourceCodeInfo_Location)(0xc00058bd00), 5:(*descriptorpb.SourceCodeInfo_Location)(0xc00058bf80)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: GetV2BaseNfts
// Name: GetV2Nfts
// Name: GetV2AdminBaseNfts
// Name: PostV2NftsPrivateMint
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000587500), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"nftid":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00058ef00)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: PostV2AdminUtilities
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000587700), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"basenftids":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00058f480)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Basenftids":true}}
// Name: PostV2RecoverMintNft
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000587a00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"contractaddress":true, "owneraddress":true, "tokenid":true}, fieldNumber:3, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc00058fa80), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc00058fd00), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc00058ff80)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
// Name: PostV2AdminNftCategoriesBulk
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc000587e00), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"categories":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000596480)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Categories":true}}
// Name: GetV2AdminUtilities
// Name: PostV2AdminUtilities
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00059a100), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:" Postv2Adminutilitiesrequest ユーティリティのinputモデル.", TrailingComments:""}, field:map[string]bool{"image":true, "title":true, "type":true, "url":true}, fieldNumber:4, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000596c80), 1:(*descriptorpb.SourceCodeInfo_Location)(0xc000596f00), 2:(*descriptorpb.SourceCodeInfo_Location)(0xc000597180), 3:(*descriptorpb.SourceCodeInfo_Location)(0xc000597400)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{"Type":true}}
// Name: GetV2SupportedNetworks
// Name: GetV2BaseNfts
// Name: GetV2Collections
// Name: PutV2MintPrepare
// inputMsg: &protobuf.MessageDescriptorProto{desc:(*descriptorpb.DescriptorProto)(0xc00059a900), comment:protobuf.Comment{LeadingDetachedComments:[]string(nil), LeadingComments:"", TrailingComments:""}, field:map[string]bool{"tokenid":true}, fieldNumber:1, fieldLocations:map[int32]*descriptorpb.SourceCodeInfo_Location{0:(*descriptorpb.SourceCodeInfo_Location)(0xc000597f80)}, enum:map[string]bool{}, enumLocations:[]*descriptorpb.SourceCodeInfo_Location(nil), nested:map[string]bool{}}
