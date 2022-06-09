// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package compiler

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/openapi"
	"go.lsp.dev/openapi2protobuf/protobuf"
	"go.lsp.dev/openapi2protobuf/protobuf/types"
)

var additionalMessages []*protobuf.MessageDescriptorProto

func RegisterAdditionalMessages(descs ...*protobuf.MessageDescriptorProto) {
	additionalMessages = append(additionalMessages, descs...)
}

var dependencyProto []string

func RegisterDependencyProto(deps string) {
	dependencyProto = append(dependencyProto, deps)
}

// Option represents an idiomatic functional option pattern to compile the Protocol Buffers structure from the OpenAPI schema.
type Option func(o *option)

// option holds an option to compile the Protocol Buffers from the OpenAPI schema.
type option struct {
	packageName        string
	useAnnotation      bool
	skipRPC            bool
	skipDeprecatedRPC  bool
	usePrefixEnum      bool
	wrapPrimitives     bool
	additionalMessages []*protobuf.MessageDescriptorProto
}

// WithPackageName specifies the package name when compiling the Protocol Buffers.
func WithPackageName(packageName string) Option {
	return func(o *option) { o.packageName = packageName }
}

// WithAnnotation sets whether the add "google.api.http" annotation to the compiled Protocol Buffers.
func WithAnnotation(useAnnotation bool) Option {
	return func(o *option) { o.useAnnotation = useAnnotation }
}

// WithSkipRPC creates a new Option to specify if we should
// generate services and rpcs in addition to messages.
func WithSkipRPC(skipRPC bool) Option {
	return func(o *option) { o.skipRPC = skipRPC }
}

// WithSkipDeprecatedRpcs creates a new Option to specify if we should.
//
// Skips generating rpcs for endpoints marked as deprecated.
func WithSkipDeprecatedRPC(skipDeprecatedRPC bool) Option {
	return func(o *option) { o.skipDeprecatedRPC = skipDeprecatedRPC }
}

// WithPrefixEnums prefix enum values with their enum name to prevent protobuf namespacing issues.
func WithPrefixEnums(usePrefixEnum bool) Option {
	return func(o *option) { o.usePrefixEnum = usePrefixEnum }
}

// WithWrapPrimitives wraps primitive types with their wrapper message types.
//
// See:
// - https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
// - https://developers.google.com/protocol-buffers/docs/proto3#default
func WithWrapPrimitives(wrapPrimitives bool) Option {
	return func(o *option) { o.wrapPrimitives = wrapPrimitives }
}

// WithAdditionalMessages adds additional messages.
func WithAdditionalMessages(additionalMessages []*protobuf.MessageDescriptorProto) Option {
	return func(o *option) { o.additionalMessages = append(o.additionalMessages, additionalMessages...) }
}

type lookupFunc func(token string) (interface{}, error)

type compiler struct {
	fdesc *protobuf.FileDescriptorProto
	opt   *option

	schemasLookupFunc lookupFunc
	pathLookupFunc    lookupFunc
}

// Compile takes an OpenAPI spec and compiles it into a protobuf file descriptor.
func Compile(ctx context.Context, spec *openapi.Schema, options ...Option) (*descriptorpb.FileDescriptorProto, error) {
	opt := &option{
		additionalMessages: additionalMessages,
	}
	for _, o := range options {
		o(opt)
	}

	pkgname := opt.packageName
	c := &compiler{
		fdesc: protobuf.NewFileDescriptorProto(pkgname),
		opt:   opt,
	}

	// append additional messages
	for _, msg := range c.opt.additionalMessages {
		c.fdesc.AddMessage(msg)
	}
	for _, deps := range dependencyProto {
		c.fdesc.AddDependency(deps)
	}

	// compile info object
	if err := c.CompileInfo(spec.Info); err != nil {
		return nil, fmt.Errorf("could not compile info object: %w", err)
	}

	// compile servers object
	if err := c.CompileServers(spec.Servers); err != nil {
		return nil, fmt.Errorf("could not compile servers object: %w", err)
	}

	// compile paths object
	if err := c.CompilePaths(spec.Paths); err != nil {
		return nil, fmt.Errorf("could not compile paths object: %w", err)
	}

	// compile all component objects
	if err := c.CompileComponents(spec.Components); err != nil {
		return nil, fmt.Errorf("could not compile component objects: %w", err)
	}

	// compile security object
	if err := c.CompileSecurity(spec.Security); err != nil {
		return nil, fmt.Errorf("could not compile security object: %w", err)
	}

	// compile tags object
	if err := c.CompileTags(spec.Tags); err != nil {
		return nil, fmt.Errorf("could not compile tags object: %w", err)
	}

	// compile external documentation object
	if err := c.CompileExternalDocs(spec.ExternalDocs); err != nil {
		return nil, fmt.Errorf("could not compile external documentation object: %w", err)
	}

	fd := c.fdesc.Build()

	// add dependency proto
	depsFileDescriptor := make([]*desc.FileDescriptor, 0, len(dependencyProto))
	for _, deps := range c.fdesc.GetDependency() {
		if depDesc, ok := types.Descriptor[deps]; ok {
			knownDesc, err := desc.CreateFileDescriptor(depDesc)
			if err != nil {
				return nil, fmt.Errorf("could not create %s descriptor: %w", depDesc.GetName(), err)
			}
			depsFileDescriptor = append(depsFileDescriptor, knownDesc)
			c.fdesc.AddDependency(deps)
		}
	}

	fdesc, err := desc.CreateFileDescriptor(fd, depsFileDescriptor...)
	if err != nil {
		return nil, fmt.Errorf("could not convert to desc: %w", err)
	}

	p := protoprint.Printer{}
	var sb strings.Builder
	if err := p.PrintProtoFile(fdesc, &sb); err != nil {
		return nil, fmt.Errorf("could not print proto: %w", err)
	}
	fmt.Fprintln(os.Stdout, sb.String())

	return fd, nil
}
