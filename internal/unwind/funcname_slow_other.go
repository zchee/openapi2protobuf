//go:build !go1.11

package unwind

import "runtime"

// FuncName returns the package/function name of the caller.
func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return Name(pc)
}

// FuncNameN returns the package/function n levels below the caller.
func FuncNameN(n int) string {
	pc, _, _, _ := runtime.Caller(1 + n)
	return Name(pc)
}
