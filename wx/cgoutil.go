package wx

import (
	"unsafe"
)

// #include "cgoutil.h"
import "C"

func init() {
	C.g_CgoReady = cBool(true)
}

//export x_newGoString
func x_newGoString(p *C.char, n int) *string {
	str := C.GoStringN(p, C.int(n))
	return &str
}

func goString(ptr C.String) string {
	// The C.String is a *string.
	return *(*string)(unsafe.Pointer(ptr))
}

//BUG: Is this safe from GC??
func cString(str *string) C.String {
	// The C.String is a *string.
	return C.String(str)
}

func cBool(b bool) C.BOOL {
	if b {
		return 1
	}
	return 0
}

func goBool(b C.BOOL) bool {
	return b != 0
}

//export x_remoteObjectPtrFromeObjectTable
func x_remoteObjectPtrFromeObjectTable(ptr unsafe.Pointer) {
	globalObjectTable.Remove(ptr)
}
