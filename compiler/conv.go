// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc/builder"
)

func normalizeMessageName(s string) string {
	return strcase.ToCamel(s)
}

func normalizeFieldName(s string) string {
	return strcase.ToSnake(s)
}

func normalizeComment(title, description string) builder.Comments {
	var sb strings.Builder

	sb.WriteString(" ") // add space after "//"
	sb.WriteString(normalizeMessageName(title))
	sb.WriteString(" ") // add space before description

	// toLower to first charactor of description
	sb.WriteByte(byte(unicode.ToLower(rune(description[0]))))
	// replaces all newline with space for after "//"
	sb.WriteString(strings.ReplaceAll(description[1:], "\n", "\n "))

	comments := builder.Comments{
		LeadingComment: sb.String(),
	}

	return comments
}
