// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"github.com/iancoleman/strcase"
)

func normalizeMessageName(s string) string {
	return strcase.ToCamel(s)
}

func normalizeFieldName(s string) string {
	return strcase.ToSnake(s)
}
