// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type FileDescriptorProto struct {
	*descriptorpb.FileDescriptorProto
}

func NewFileDescriptorProto(name string) *FileDescriptorProto {
	return &FileDescriptorProto{
		FileDescriptorProto: &descriptorpb.FileDescriptorProto{
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

func (fd *FileDescriptorProto) AddDependency(fdp *FileDescriptorProto) *FileDescriptorProto {
	fd.Dependency = append(fd.Dependency, fdp.GetName())

	return fd
}

func (fd *FileDescriptorProto) AddMessage(msgs ...*MessageDescriptorProto) *FileDescriptorProto {
	ms := make([]*descriptorpb.DescriptorProto, len(msgs))
	for i, msg := range msgs {
		ms[i] = msg.Descriptor()
	}
	fd.MessageType = append(fd.MessageType, ms...)

	return fd
}

func (fd *FileDescriptorProto) AddEnum(enums ...*EnumDescriptorProto) *FileDescriptorProto {
	es := make([]*descriptorpb.EnumDescriptorProto, len(enums))
	for i, enum := range enums {
		es[i] = enum.Descriptor()
	}
	fd.EnumType = append(fd.EnumType, es...)

	return fd
}

func (fd *FileDescriptorProto) AddService(servicses ...*ServiceDescriptorProto) *FileDescriptorProto {
	svcs := make([]*descriptorpb.ServiceDescriptorProto, len(servicses))
	for i, servicse := range servicses {
		svcs[i] = servicse.Descriptor()
	}
	fd.Service = append(fd.Service, svcs...)

	return fd
}

func (fd *FileDescriptorProto) AddExtension(exts ...*FieldDescriptorProto) *FileDescriptorProto {
	es := make([]*descriptorpb.FieldDescriptorProto, len(exts))
	for i, ext := range exts {
		es[i] = ext.Descriptor()
	}
	fd.Extension = append(fd.Extension, es...)

	return fd
}
