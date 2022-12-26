// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package conv

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gobuffalo/flect"
)

// RegisterAcronyms registers acronym list to flect package.
func RegisterAcronyms(ss []string) {
	var buf bytes.Buffer

	buf.WriteByte('[')
	for i, s := range ss {
		buf.WriteString(fmt.Sprintf("%q", s))
		if i != len(ss)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')

	if err := flect.LoadAcronyms(bytes.NewReader(buf.Bytes())); err != nil {
		panic(err)
	}
}

// NormalizeMessageName normalizes s to the proto message name.
func NormalizeMessageName(s string) string {
	s = flect.Pascalize(s)
	if strings.Contains(s, "IDS") {
		s = strings.ReplaceAll(s, "IDS", "IDs")
	}

	return s
}

// NormalizeFieldName normalizes s to the proto field name.
func NormalizeFieldName(s string) string {
	return flect.Underscore(s)
}

func ToSingularize(s string) string {
	return flect.Singularize(s)
}

// NormalizeComment normalizes the title and description to Go style comment.
//
// This function returns the string that assumes inserting after the `//` token.
func NormalizeComment(title, description string) string {
	var sb strings.Builder

	sb.WriteString(" ") // add space after "//"
	sb.WriteString(NormalizeMessageName(title))
	sb.WriteString(" is the") // add godoc style words
	sb.WriteString(" ")       // add space after title

	// parse first word
	idx := strings.Index(description, " ")
	switch idx {
	case -1:
		sb.WriteString(flect.Camelize(description))
	default:
		sb.WriteString(flect.Camelize(strings.ToLower(description[:idx])))
		// replaces all newline with space for after "//"
		sb.WriteString(strings.ReplaceAll(description[idx:], "\n", "\n "))
	}
	if description[len(description)-1] != '.' {
		sb.WriteString(".")
	}

	return sb.String()
}

// IsVowel returns whether the r is vowel letter.
func IsVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
