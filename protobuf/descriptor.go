// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.lsp.dev/openapi2protobuf/protobuf/prototag"
)

type FileDescriptorProto struct {
	desc       *descriptorpb.FileDescriptorProto
	components map[string]bool
	msgs       map[string]bool
	enums      map[string]bool
	deps       map[string]bool
}

func NewFileDescriptorProto(fqn string) *FileDescriptorProto {
	return &FileDescriptorProto{
		desc: &descriptorpb.FileDescriptorProto{
			Name:    proto.String(strcase.ToSnake(splitByLastDot(fqn))),
			Package: proto.String(fqn),
			Options: &descriptorpb.FileOptions{
				GoPackage: proto.String(goPackage(fqn)),
				// JavaPackage:        proto.String(javaPackage(fqn)),
				// JavaOuterClassname: proto.String(strcase.ToCamel(splitByLastDot(fqn))),
				// JavaMultipleFiles:  proto.Bool(true),
				// CsharpNamespace:    proto.String(csharpNamespace(fqn)),
				// ObjcClassPrefix:    proto.String(objcClassPrefix(fqn)),
				// CcEnableArenas: proto.Bool(true),
			},
			Syntax:         proto.String(protoreflect.Proto3.String()),
			SourceCodeInfo: new(descriptorpb.SourceCodeInfo),
		},
		components: make(map[string]bool),
		msgs:       make(map[string]bool),
		enums:      make(map[string]bool),
		deps:       make(map[string]bool),
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
	return fmt.Sprintf("%s", fqn)
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

func (fd *FileDescriptorProto) AddComponent(name string) {
	fd.components[name] = true
}

func (fd *FileDescriptorProto) HasComponent(name string) bool {
	return fd.components[name]
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

func (fd *FileDescriptorProto) GetDependency() []string {
	return fd.desc.Dependency
}

func (fd *FileDescriptorProto) AddDependency(deps string) *FileDescriptorProto {
	if fd.deps[deps] {
		return fd
	}

	fd.deps[deps] = true
	fd.desc.Dependency = append(fd.desc.Dependency, deps)

	return fd
}

func (fd *FileDescriptorProto) AddMessage(msg *MessageDescriptorProto) *FileDescriptorProto {
	if fd.msgs[msg.GetName()] {
		return fd
	}

	for _, nested := range msg.GetNestedMessages() {
		if nested == msg.GetName() {
			return fd
		}
	}

	fd.msgs[msg.GetName()] = true

	comments := msg.GetComments()
	loc := &descriptorpb.SourceCodeInfo_Location{
		LeadingComments:         proto.String(comments.LeadingComments),
		TrailingComments:        proto.String(comments.TrailingComments),
		LeadingDetachedComments: comments.LeadingDetachedComments,
		Path:                    []int32{prototag.FileMessageType, int32(len(fd.msgs)) - 1},
	}
	fd.desc.SourceCodeInfo.Location = append(fd.desc.SourceCodeInfo.Location, loc)
	fd.desc.MessageType = append(fd.desc.MessageType, msg.Build())

	return fd
}

func (fd *FileDescriptorProto) AddMessageDescriptor(desc *descriptorpb.DescriptorProto) *FileDescriptorProto {
	fd.desc.MessageType = append(fd.desc.MessageType, desc)

	return fd
}

func (fd *FileDescriptorProto) AddEnum(enum *EnumDescriptorProto) *FileDescriptorProto {
	if fd.enums[enum.GetName()] {
		return fd
	}

	fd.enums[enum.GetName()] = true
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

func (fd *FileDescriptorProto) AddSourceCodeInfoLocation(loc *descriptorpb.SourceCodeInfo_Location) *FileDescriptorProto {
	fd.desc.SourceCodeInfo.Location = append(fd.desc.SourceCodeInfo.Location, loc)

	return fd
}

func (fd *FileDescriptorProto) Build() *descriptorpb.FileDescriptorProto {
	sort.Slice(fd.desc.Dependency, func(i, j int) bool { return fd.desc.Dependency[i] < fd.desc.Dependency[j] })
	sort.Slice(fd.desc.EnumType, func(i, j int) bool { return fd.desc.EnumType[i].GetName() < fd.desc.EnumType[j].GetName() })
	sort.Slice(fd.desc.Service, func(i, j int) bool { return fd.desc.Service[i].GetName() < fd.desc.Service[j].GetName() })
	sort.Slice(fd.desc.Extension, func(i, j int) bool { return fd.desc.Extension[i].GetName() < fd.desc.Extension[j].GetName() })

	return fd.desc
}
