// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
