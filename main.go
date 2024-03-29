// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

// Command openapi2protobuf generates Protocol Buffers v3 and gRPC services definitions from the OpenAPI/Swagger schema.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.lsp.dev/openapi2protobuf/compiler"
	"go.lsp.dev/openapi2protobuf/internal/conv"
	"go.lsp.dev/openapi2protobuf/openapi"
)

func init() {
	conv.RegisterAcronyms(
		[]string{
			"NFT",
			"DID",
			"IDs",
		},
	)
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	f := args[0]
	pkgname := args[1]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	schema, err := openapi.LoadFile(ctx, f)
	if err != nil {
		return fmt.Errorf("could not load %s OpenAPI file: %w", f, err)
	}

	if _, err = compiler.Compile(ctx, schema, compiler.WithPackageName(pkgname)); err != nil {
		return fmt.Errorf("could not compile file descriptor: %w", err)
	}

	return nil
}
