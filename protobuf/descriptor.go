// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type FileDescriptorProto struct {
	desc *descriptorpb.FileDescriptorProto
}

func NewFileDescriptorProto(name string) *FileDescriptorProto {
	return &FileDescriptorProto{
		desc: &descriptorpb.FileDescriptorProto{
			Name: proto.String(name),
			Options: &descriptorpb.FileOptions{
				OptimizeFor:    descriptorpb.FileOptions_SPEED.Enum(),
				GoPackage:      proto.String(name),
				CcEnableArenas: proto.Bool(true),
			},
			Syntax: proto.String(protoreflect.Proto3.String()),
		},
	}
}

func (fd *FileDescriptorProto) GetName() string {
	return fd.desc.GetName()
}

func (fd *FileDescriptorProto) AddDependency(fdp *FileDescriptorProto) *FileDescriptorProto {
	fd.desc.Dependency = append(fd.desc.Dependency, fdp.desc.GetName())

	return fd
}

func (fd *FileDescriptorProto) AddMessage(msgs ...*MessageDescriptorProto) *FileDescriptorProto {
	ms := make([]*descriptorpb.DescriptorProto, len(msgs))
	for i, msg := range msgs {
		ms[i] = msg.Build()
	}
	fd.desc.MessageType = append(fd.desc.MessageType, ms...)

	return fd
}

func (fd *FileDescriptorProto) AddEnum(enums ...*EnumDescriptorProto) *FileDescriptorProto {
	es := make([]*descriptorpb.EnumDescriptorProto, len(enums))
	for i, enum := range enums {
		es[i] = enum.Build()
	}
	fd.desc.EnumType = append(fd.desc.EnumType, es...)

	return fd
}

func (fd *FileDescriptorProto) AddService(servicses ...*ServiceDescriptorProto) *FileDescriptorProto {
	svcs := make([]*descriptorpb.ServiceDescriptorProto, len(servicses))
	for i, servicse := range servicses {
		svcs[i] = servicse.Build()
	}
	fd.desc.Service = append(fd.desc.Service, svcs...)

	return fd
}

func (fd *FileDescriptorProto) AddExtension(exts ...*FieldDescriptorProto) *FileDescriptorProto {
	es := make([]*descriptorpb.FieldDescriptorProto, len(exts))
	for i, ext := range exts {
		es[i] = ext.Build()
	}
	fd.desc.Extension = append(fd.desc.Extension, es...)

	return fd
}

func (fd *FileDescriptorProto) Build() *descriptorpb.FileDescriptorProto {
	return fd.desc
}
