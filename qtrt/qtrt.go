package qtrt

/*
#cgo CFLAGS: -std=c11

#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"unsafe"
)

func CString(s string) unsafe.Pointer  { return unsafe.Pointer(C.CString(s)) }
func GoString(p unsafe.Pointer) string { return C.GoString((*C.char)(p)) }
func GoStringI(p uint64) string        { return GoString(unsafe.Pointer(uintptr(p))) }

// 所有的Qt绑定类必须继承自这个
type CObject struct {
	Cthis unsafe.Pointer
}

type GetCthiser interface {
	GetCthis() unsafe.Pointer
}

var DebugFinal bool = false

func objectFinalBefore(o interface{}) {
	if DebugFinal {
		log.Println(o, fmt.Sprintf("%#v", o), o.(GetCthiser).GetCthis())
	}
}
func objectFinalAfter(o interface{}) {
	if DebugFinal {
		log.Println(o, fmt.Sprintf("%#v", o), o.(GetCthiser).GetCthis())
	}
}
func SetFinalizer(obj interface{}, finalizer interface{}) {
	runtime.SetFinalizer(obj, func(o interface{}) {
		objectFinalBefore(o)
		ov := reflect.ValueOf(o)
		fv := reflect.ValueOf(finalizer)
		fv.Call([]reflect.Value{ov})
		objectFinalAfter(o)
	})
}

// 直接使用C++ symbols
var UseCppSymbols bool = false

/////////

func BoolTy(pointer bool) reflect.Type {
	if pointer {
		var v = true
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(true)
}

func ByteTy(pointer bool) reflect.Type {
	if pointer {
		var v = byte(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(byte(0))
}

func StringTy(pointer bool) reflect.Type {
	var s = "foo"
	if pointer {
		return reflect.TypeOf(&s)
	}
	return reflect.TypeOf(s)
}

func RuneTy(pointer bool) reflect.Type {
	var s rune = '\000'
	if pointer {
		return reflect.TypeOf(&s)
	}
	return reflect.TypeOf(s)
}

func Int16Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int16(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int16(0))
}

func UInt16Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint16(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint16(0))
}

func Int32Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int32(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int32(0))
}

func UInt32Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint32(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint32(0))
}

func Int64Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int64(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int64(0))
}

func UInt64Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint64(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint64(0))
}

func FloatTy(pointer bool) reflect.Type {
	if pointer {
		var v = float32(0.0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(float32(0.0))
}
func DoubleTy(pointer bool) reflect.Type {
	if pointer {
		var v = float64(0.0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(float64(0.0))
}

func VoidpTy() reflect.Type {
	var v unsafe.Pointer = nil
	return reflect.TypeOf(v)
}
