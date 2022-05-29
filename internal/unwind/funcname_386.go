//go:build go1.11 && 386

package unwind

import _ "unsafe"

//go:noescape
//go:linkname callers runtime.callers
func callers(skip int, pcbuf []uintptr) int

// FuncName returns the package/function name of the caller.
func FuncName() string

// FuncName returns the package/function n levels below the caller.
func FuncNameN(n int) string {
	var pcbuf [1]uintptr
	callers(1+n, pcbuf[:])
	return Name(pcbuf[0])
}
