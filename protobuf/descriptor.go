// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	// UnsafeEnabled specifies whether package unsafe can be used.
	_ = protoimpl.UnsafeEnabled

	// Types used by generated code in init functions.
	_ protoimpl.DescBuilder
	_ protoimpl.TypeBuilder

	// Types used by generated code to implement EnumType, MessageType, and ExtensionType.
	_ protoimpl.EnumInfo
	_ protoimpl.MessageInfo
	_ protoimpl.ExtensionInfo

	// Types embedded in generated messages.
	_ protoimpl.MessageState
	_ protoimpl.SizeCache
	_ protoimpl.WeakFields
	_ protoimpl.UnknownFields
	_ protoimpl.ExtensionFields
	_ protoimpl.ExtensionFieldV1

	_ protoimpl.Pointer

	_ = protoimpl.X
)

type FileDescriptorProto struct {
	desc *descriptorpb.FileDescriptorProto
	msg  map[string]bool
	enum map[string]bool
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
				JavaMultipleFiles:  proto.Bool(true),
				CsharpNamespace:    proto.String(csharpNamespace(fqn)),
				// ObjcClassPrefix:    proto.String(objcClassPrefix(fqn)),
				CcEnableArenas: proto.Bool(true),
			},
			Syntax: proto.String(protoreflect.Proto3.String()),
		},
		msg:  make(map[string]bool),
		enum: make(map[string]bool),
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

func objcClassPrefix(fqn string) string {
	isUpper := func(s string) bool {
		return "A" <= s && s >= "Z"
	}
	_ = isUpper

	var b strings.Builder
	ss := strings.Split(strcase.ToSnake(fqn), "_")
	for _, s := range ss {
		b.WriteByte(s[0])
	}

	return b.String()
}

func (fd *FileDescriptorProto) GetName() string {
	return fd.desc.GetName()
	// return string(fd.desc.ProtoReflect().Descriptor().Name())
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

func (fd *FileDescriptorProto) AddMessage(msg *MessageDescriptorProto) *FileDescriptorProto {
	if fd.msg[msg.GetName()] {
		return fd
	}

	for _, nested := range msg.GetNestedMessages() {
		if nested == msg.GetName() {
			return fd
		}
	}

	fd.msg[msg.GetName()] = true
	fd.desc.MessageType = append(fd.desc.MessageType, msg.Build())

	return fd
}

func (fd *FileDescriptorProto) AddEnum(enum *EnumDescriptorProto) *FileDescriptorProto {
	if fd.enum[enum.GetName()] {
		return fd
	}

	fd.enum[enum.GetName()] = true
	fd.desc.EnumType = append(fd.desc.EnumType, enum.Build())

	return fd
}

func (fd *FileDescriptorProto) AddService(servicse *ServiceDescriptorProto) *FileDescriptorProto {
	fd.desc.Service = append(fd.desc.Service, servicse.Build())

	return fd
}

func (fd *FileDescriptorProto) AddExtension(ext *FieldDescriptorProto) *FileDescriptorProto {
	fd.desc.Extension = append(fd.desc.Extension, ext.Build())

	return fd
}

func (fd *FileDescriptorProto) Build() *descriptorpb.FileDescriptorProto {
	return fd.desc
}
