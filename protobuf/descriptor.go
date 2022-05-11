// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type FileDescriptorProto struct {
	desc *descriptorpb.FileDescriptorProto
}

func NewFileDescriptorProto(fqn string) *FileDescriptorProto {
	return &FileDescriptorProto{
		desc: &descriptorpb.FileDescriptorProto{
			Name:    proto.String(strcase.ToSnake(splitByLastDot(fqn))),
			Package: proto.String(fqn),
			Options: &descriptorpb.FileOptions{
				GoPackage:          proto.String(goPackage(fqn)),
				JavaPackage:        proto.String(javaPackage(fqn)),
				JavaOuterClassname: proto.String(strcase.ToCamel(splitByLastDot(fqn))),
				CsharpNamespace:    proto.String(csharpNamespace(fqn)),
				// ObjcClassPrefix:    proto.String(""), // TODO(zchee): split by dot and join only first letter
				CcEnableArenas: proto.Bool(true),
			},
			Syntax: proto.String(protoreflect.Proto3.String()),
		},
	}
}

func splitByDot(fqn string) []string {
	ss := strings.Split(fqn, ".")
	if len(ss) < 2 {
		return []string{fqn}
	}

	return ss
}

func splitByLastDot(fqn string) string {
	idx := strings.LastIndex(fqn, ".")
	if idx == -1 {
		return fqn
	}

	return fqn[idx+1:]
}

func goPackage(fqn string) string {
	return fmt.Sprintf("%s;%s", fqn, splitByLastDot(fqn))
}

func javaPackage(fqn string) string {
	ss := splitByDot(fqn)
	ss = ss[:len(ss)-1]
	for lh, rh := 0, len(ss)-1; lh < rh; lh, rh = lh+1, rh-1 {
		ss[lh], ss[rh] = ss[rh], ss[lh]
	}

	return strings.Join(ss, ".")
}

func csharpNamespace(fqn string) string {
	ss := splitByDot(fqn)
	for i, s := range ss {
		ss[i] = strcase.ToCamel(s)
	}

	return strings.Join(ss, ".")
}

func (fd *FileDescriptorProto) GetName() string {
	return fd.desc.GetName()
}

func (fd *FileDescriptorProto) SetName(name string) {
	fd.desc.Name = proto.String(name)
}

func (fd *FileDescriptorProto) SetPackage(fqn string) {
	fd.desc.Package = proto.String(fqn)
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
