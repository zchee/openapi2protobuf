// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"github.com/getkin/kin-openapi/openapi3"

	"go.lsp.dev/openapi2protobuf/internal/conv"
)

// CompileInfo compiles info object.
func (c *compiler) CompileInfo(info *openapi3.Info) error {
	if info == nil {
		return nil
	}

	if title := info.Title; title != "" {
		c.fdesc.SetPackage(conv.NormalizeFieldName(title))
	}

	if description := info.Description; description != "" {
		// c.fb.PackageComments.TrailingComment = description
	}

	if version := info.Version; version != "" {
		// c.fb.PackageComments.LeadingComment += " " + version
	}

	return nil
}
