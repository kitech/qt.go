package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"github.com/gonuts/ffi"
)

type QWidget struct {
}

func NewQWidget() *QWidget {
	o := &QWidget{}
	InvokeQtFunc("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", ffi.C_pointer, nil)
	return o
}

type QApplication struct {
}

func NewQApplication() *QApplication {
	o := &QApplication{}
	t := C.calloc(1, 128)
	// args := C.CString("testprog")
	InvokeQtFunc("_ZN12QApplicationC2ERiPPci", nil,
		// InvokeQtFunc("_ZN3AbcC2ERiPPci", nil,
		// InvokeQtFunc("_ZN3AbcC2Eiiii", nil,
		[]Type{ffi.C_uint64, ffi.C_int},
		uint64(uintptr(t)), 0)

	return o
}
