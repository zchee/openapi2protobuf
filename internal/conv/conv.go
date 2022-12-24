// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package conv

import (
	"net/http"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

var UpperCaseAcronym = map[string]string{
	http.MethodGet:     "Get",
	http.MethodHead:    "Head",
	http.MethodPost:    "Post",
	http.MethodPut:     "Put",
	http.MethodPatch:   "Patch", // RFC 5789
	http.MethodDelete:  "Delete",
	http.MethodConnect: "Connect",
	http.MethodOptions: "Options",
	http.MethodTrace:   "Trace",
}

func NormalizeMessageName(s string) string {
	camel := strcase.ToCamel(s)
	for acronym, replace := range UpperCaseAcronym {
		if strings.Contains(camel, acronym) {
			camel = strings.ReplaceAll(camel, acronym, replace)
		}
	}

	return camel
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
	sb.WriteString(" is the") // add godoc style words
	sb.WriteString(" ")       // add space after title

	// hasAcronym := false
	// for acronym := range UpperCaseAcronym {
	// 	if strings.Contains(description, acronym) {
	// 		hasAcronym = true
	// 	}
	// }
	// if !hasAcronym {
	// // ToLower the first letter of the description
	// sb.WriteByte(byte(unicode.ToLower(rune(description[0]))))
	// }

	// ToLower the first letter of the description
	sb.WriteByte(byte(unicode.ToLower(rune(description[0]))))
	// replaces all newline with space for after "//"
	sb.WriteString(strings.ReplaceAll(description[1:], "\n", "\n "))
	if description[len(description)-1] != '.' {
		sb.WriteString(".")
	}

	return sb.String()
}
