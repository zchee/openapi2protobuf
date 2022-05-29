//go:build !go1.12

package unwind

import "runtime"

// Name returns the function name for the given pc.
func Name(pc uintptr) string {
	return runtime.FuncForPC(pc).Name()
}
