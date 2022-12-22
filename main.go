// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

// Command openapi2protobuf generates Protocol Buffers v3 and gRPC services definitions from the OpenAPI/Swagger schema.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "go.lsp.dev/openapi2protobuf/internal/lsp"

	"go.lsp.dev/openapi2protobuf/compiler"
	"go.lsp.dev/openapi2protobuf/openapi"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	f := args[0]
	// f := "testdata/lsp/3.17/types-3.17.0-next.8.openapi.yaml"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	schema, err := openapi.LoadFile(ctx, f)
	if err != nil {
		return fmt.Errorf("could not load %s OpenAPI file: %w", f, err)
	}

	if _, err = compiler.Compile(ctx, schema); err != nil {
		return fmt.Errorf("could not compile file descriptor: %w", err)
	}

	return nil
}
