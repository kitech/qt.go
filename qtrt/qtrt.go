package qtrt

/*
#cgo CFLAGS: -std=c11

#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func CString(s string) Voidptr {
	return Voidptr(C.CString(s))
}

// string null terminated
func Stringnt(s *string) *string {
	*s = *s + "\x00"
	return s
}

func Stringzt(s string) string { return s + "\x00" }

// note: s will change
// ch = ch + "\x00", and it works
func CStringRef(s *string) Voidptr {
	*s = *s + "\x00" // copy once, but still fast than C.CString/C.free
	strhdr := (*reflect.StringHeader)(Voidptr(s))
	if strhdr.Len == 0 {
		return nil
	}
	return Voidptr(strhdr.Data)
}

// caller add null tail
func CStringRefRaw(s *string) Voidptr {
	// *s = *s + "\x00" // copy once, but still fast than C.CString/C.free
	strhdr := (*reflect.StringHeader)(Voidptr(s))
	if strhdr.Len == 0 {
		return nil
	}
	return Voidptr(strhdr.Data)
}
func GoString(p Voidptr) string         { return C.GoString((*C.char)(p)) }
func GoStringN(p Voidptr, n int) string { return C.GoStringN((*C.char)(p), C.int(n)) }
func GoStringI(p uint64) string         { return GoString(Voidptr(uintptr(p))) }
func GoStringIN(p uint64, n int) string { return GoStringN(Voidptr(uintptr(p)), n) }

var zeromem = C.calloc(1, 8096)

func Cmemset(p Voidptr, c int, n int) {
	if p != nil {
		// C.memset(p, C.int(c), C.size_t(n))
	}
	if false { // 在这里也不能set了，回收的的内存有可能重新分配使用了
		r1 := C.memcmp(p, zeromem, C.size_t(n))
		C.memcpy(p, zeromem, C.size_t(n)) // 采用溢出副作用，而memset则会导致无效内存访问。
		r2 := C.memcmp(p, zeromem, C.size_t(n))
		log.Println("iszero after dtor:", r1, r2)
	}
}

// 所有的Qt绑定类必须继承自这个
type CObject struct {
	Cthis Voidptr
}
type GetCthiser interface {
	GetCthis() Voidptr
}

func (this *CObject) GetCthis() Voidptr {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *CObject) SetCthis(cthis Voidptr) {
	this.Cthis = cthis
}
func (this *CObject) Addr() Voidptr { return Voidptr(&this.Cthis) }

// some fixed internal names
const GetCthisName = "GetCthis"
const SetCthisName = "SetCthis"
const DtorCthisName = "DtorCthis"
const NewCthisName = "NewCthis"
const SetInitStObjName = "C_%s_init_staticMetaObject"

func GetCthis(obj interface{}) Voidptr {
	objval := reflect.ValueOf(obj).Elem()
	retx := objval.MethodByName(GetCthisName).Call(nil)
	return retx[0].Interface().(Voidptr)
}

///////////
var DebugFinalize bool = false
var FinalizeProxyFn = func(f func()) { f() } //
func init() {
	dbgval := os.Getenv("QTGO_DEBUG_FINALIZE")
	if strings.ToLower(dbgval) == "true" || dbgval == "1" {
		DebugFinalize = true
	}
}

func objectFinalBefore(o interface{}) {
	if DebugFinalize {
		log.Println(o, fmt.Sprintf("%#v", o), o.(GetCthiser).GetCthis())
	}
}
func objectFinalAfter(o interface{}) {
	if DebugFinalize {
		log.Println(o, fmt.Sprintf("%#v", o), o.(GetCthiser).GetCthis())
	}
}
func SetFinalizer(obj interface{}, finalizer interface{}) {
	runtime.SetFinalizer(obj, func(o interface{}) {
		FinalizeProxyFn(func() {
			objectFinalBefore(o)
			ov := reflect.ValueOf(o)
			fv := reflect.ValueOf(finalizer)
			insure := false
			if finalizerObjectFilterFn != nil {
				insure = finalizerObjectFilterFn(ov)
			}
			if insure {
				fv.Call([]reflect.Value{ov})
			}
			objectFinalAfter(o)
		})
	})
}
func UnsetFinalizer(obj interface{})   { runtime.SetFinalizer(obj, nil) }
func ReleaseOwnerToQt(obj interface{}) { UnsetFinalizer(obj) }

// 返回true表示确定
var finalizerObjectFilterFn func(reflect.Value) bool = func(ov reflect.Value) bool { return true }

func SetFinalizerObjectFilter(f func(reflect.Value) bool) {
	finalizerObjectFilterFn = f
}

// 直接使用C++ symbols
const UseCppSymbols = true

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
	var v Voidptr = nil
	return reflect.TypeOf(v)
}
