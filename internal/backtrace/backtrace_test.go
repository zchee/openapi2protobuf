package backtrace

import (
	"testing"
)

type thisTest struct{}

//go:noinline
func (thisTest) method() string { return This() }

//go:noinline
func (*thisTest) pmethod() string { return This() }

func (thisTest) method2() string { return thisTest{}.method() }
func (thisTest) method3() string { return thisTest{}.method2() }

func TestThis(t *testing.T) {
	if this := FuncName(); this != "go.lsp.dev/openapi2protobuf/internal/unwind.TestThis" {
		t.Errorf("got %v but want %v", FuncName(), "go.lsp.dev/openapi2protobuf/internal/unwind.TestThis")
	}

	var tt thisTest
	if method1 := tt.method(); method1 != "go.lsp.dev/openapi2protobuf/internal/unwind.method" {
		t.Errorf("got %v but want %v", method1, "go.lsp.dev/openapi2protobuf/internal/unwind.method")
	}

	if method2 := tt.method2(); method2 != "go.lsp.dev/openapi2protobuf/internal/unwind.method" {
		t.Errorf("got %v but want %v", method2, "go.lsp.dev/openapi2protobuf/internal/unwind.method")
	}

	if method3 := tt.method2(); method3 != "go.lsp.dev/openapi2protobuf/internal/unwind.method" {
		t.Errorf("got %v but want %v", method3, "go.lsp.dev/openapi2protobuf/internal/unwind.method")
	}

	if pmethod := new(thisTest).pmethod(); pmethod != "go.lsp.dev/openapi2protobuf/internal/unwind.(*thisTest).pmethod" {
		t.Errorf("got %v but want %v", pmethod, "go.lsp.dev/openapi2protobuf/internal/unwind.(*thisTest).pmethod")
	}
}

func TestThisN(t *testing.T) {
	if ttn := FuncNameN(0); ttn != "go.lsp.dev/openapi2protobuf/internal/unwind.TestThisN" {
		t.Errorf("got %v but want %v", ttn, "go.lsp.dev/openapi2protobuf/internal/unwind.TestThisN")
	}

	if ttn := FuncNameN(1); ttn != "testing.tRunner" {
		t.Errorf("got %v but want %v", ttn, "go.lsp.dev/openapi2protobuf/internal/unwind.TestThisN")
	}
}

func BenchmarkThis(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FuncName()
		}
	})

	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			thisTest{}.method()
		}
	})

	b.Run("InlinedTwice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			thisTest{}.method2()
		}
	})

	b.Run("InlinedThrice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			thisTest{}.method3()
		}
	})
}

func BenchmarkThisN(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FuncNameN(1)
		}
	})
}
