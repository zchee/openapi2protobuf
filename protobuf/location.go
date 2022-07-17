// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

type Comment struct {
	LeadingDetachedComments []string
	LeadingComments         string
	TrailingComments        string
}
