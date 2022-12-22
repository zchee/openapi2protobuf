// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package backtrace

import (
	"testing"
)

type funcNameTest struct{}

//go:noinline
func (funcNameTest) method() string { return FuncName() }

//go:noinline
func (*funcNameTest) pmethod() string { return FuncName() }

func (funcNameTest) method2() string { return funcNameTest{}.method() }
func (funcNameTest) method3() string { return funcNameTest{}.method2() }

func TestFuncName(t *testing.T) {
	if this := FuncName(); this != "go.lsp.dev/openapi2protobuf/internal/backtrace.TestFuncName" {
		t.Errorf("got %v but want %v", FuncName(), "go.lsp.dev/openapi2protobuf/internal/backtrace.TestFuncName")
	}

	var tt funcNameTest
	if method1 := tt.method(); method1 != "go.lsp.dev/openapi2protobuf/internal/backtrace.funcNameTest.method" {
		t.Errorf("got %v but want %v", method1, "go.lsp.dev/openapi2protobuf/internal/backtrace.funcNameTest.method")
	}

	if method2 := tt.method2(); method2 != "go.lsp.dev/openapi2protobuf/internal/backtrace.funcNameTest.method" {
		t.Errorf("got %v but want %v", method2, "go.lsp.dev/openapi2protobuf/internal/backtrace.funcNameTest.method")
	}

	if method3 := tt.method2(); method3 != "go.lsp.dev/openapi2protobuf/internal/backtrace.funcNameTest.method" {
		t.Errorf("got %v but want %v", method3, "go.lsp.dev/openapi2protobuf/internal/backtrace.funcNameTest.method")
	}

	if pmethod := new(funcNameTest).pmethod(); pmethod != "go.lsp.dev/openapi2protobuf/internal/backtrace.(*funcNameTest).pmethod" {
		t.Errorf("got %v but want %v", pmethod, "go.lsp.dev/openapi2protobuf/internal/backtrace.(*funcNameTest).pmethod")
	}
}

func TestFuncNameN(t *testing.T) {
	if ttn := FuncNameN(0); ttn != "go.lsp.dev/openapi2protobuf/internal/backtrace.TestFuncNameN" {
		t.Errorf("got %v but want %v", ttn, "go.lsp.dev/openapi2protobuf/internal/backtrace.TestFuncNameN")
	}

	if ttn := FuncNameN(1); ttn != "testing.tRunner" {
		t.Errorf("got %v but want %v", ttn, "go.lsp.dev/openapi2protobuf/internal/backtrace.TestFuncNameN")
	}
}

func BenchmarkFuncName(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FuncName()
		}
	})

	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			funcNameTest{}.method()
		}
	})

	b.Run("InlinedTwice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			funcNameTest{}.method2()
		}
	})

	b.Run("InlinedThrice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			funcNameTest{}.method3()
		}
	})
}

func BenchmarkFuncNameN(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FuncNameN(1)
		}
	})
}
