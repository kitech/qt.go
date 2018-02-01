package ffiqt

import (
	"fmt"
	"unsafe"
)

/*
 */

func SetCallback(name string, fnptr unsafe.Pointer) {
	symname := fmt.Sprintf("set_callback%s", name)
	InvokeQtFunc6(symname, FFI_TYPE_VOID, fnptr)
}
