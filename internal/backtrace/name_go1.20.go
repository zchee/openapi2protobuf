//go:build go1.20

package backtrace

import (
	"runtime"
	"unsafe"
)

// taken from runtime/symtab.go
type funcInfo struct {
	*runtime.Func
	_ unsafe.Pointer
}

// taken from runtime/symtab.go
type inlinedCall struct {
	_  uint8
	_  [3]byte
	fn int32
	_  int32
}

//go:linkname findfunc runtime.findfunc
func findfunc(pc uintptr) funcInfo

//go:linkname funcname runtime.funcname
func funcname(f funcInfo) string

//go:linkname funcdata runtime.funcdata
func funcdata(f funcInfo, i uint8) unsafe.Pointer

//go:linkname pcdatavalue runtime.pcdatavalue
func pcdatavalue(f funcInfo, table int32, targetpc uintptr, cache unsafe.Pointer) int32

//go:linkname funcnameFromNameOff runtime.funcnameFromNameOff
func funcnameFromNameOff(f funcInfo, nameOff int32) string

// Name returns the function name for the given pc.
func Name(pc uintptr) string {
	info := findfunc(pc)

	// adjust pc if necessary
	if pc > info.Entry() {
		pc--
	}

	// attempt to determine name, walking inlining data
	name := funcname(info)
	inldata := funcdata(info, 4)
	if inldata == nil {
		return name
	}

	inltree := (*[1 << 20]inlinedCall)(inldata)
	ix := pcdatavalue(info, 2, pc, nil)
	if ix < 0 {
		return name
	}

	return funcnameFromNameOff(info, inltree[ix].fn)
}
