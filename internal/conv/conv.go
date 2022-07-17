// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package conv

import (
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

func NormalizeMessageName(s string) string {
	return strcase.ToCamel(s)
}

func NormalizeFieldName(s string) string {
	return strcase.ToSnake(s)
}

// NormalizeComment normalizes title and description to Go style comment.
//
// This function returns the string that assumes inserting after the "//" token.
func NormalizeComment(title, description string) string {
	var sb strings.Builder

	sb.WriteString(" ") // add space after "//"
	sb.WriteString(NormalizeMessageName(title))
	sb.WriteString(" ") // add space after title

	// ToLower the first letter of the description
	sb.WriteByte(byte(unicode.ToLower(rune(description[0]))))
	// replaces all newline with space for after "//"
	sb.WriteString(strings.ReplaceAll(description[1:], "\n", "\n "))
	if description[len(description)-1] != '.' {
		sb.WriteString(".")
	}

	return sb.String()
}
