//go:build go1.11 && 386
// +build go1.11,386

#include "textflag.h"
#include "funcdata.h"

// This code taken from inserting the following into
// the runtime package, compiling it, and inspecting it
// with go tool objdump -s foo
//
//    //go:noinline
//    func bar(pc uintptr) string {
//           return ""
//    }
//
//    //go:noinline
//    func foo() string {
//           return bar(getcallerpc())
//    }

TEXT ·FuncName(SB),0,$12-8
	// This is a lie, but the pointers are to readonly data
	NO_LOCAL_POINTERS

	MOVL addr-4(FP), AX
	MOVL AX, 0(SP)
	CALL ·Name(SB)
	MOVL 8(SP), AX
	MOVL 4(SP), CX
	MOVL CX, ret_base+0(FP)
	MOVL AX, ret_len+4(FP)
	RET
