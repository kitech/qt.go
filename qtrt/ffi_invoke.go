package qtrt

/*
///////////
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

#include "ffi.h"

void *lib = 0;

static void ffi_call_0(void*fn) {
  ffi_cif cif;
  ffi_type *args[10];
  void *values[10];
  char *s;
  ffi_arg rc;

   args[0] = &ffi_type_pointer;
   values[0] = &s;

   args[1] = &ffi_type_pointer;
   s = calloc(1, 256);
   strcpy(s, "dbcdefg");

   int a0 = 3;
   int* a1 = &a0;
   values[1] = &a1;

   args[2] = &ffi_type_pointer;
   char *a2[] = {"testprog", "hjdskkk"};
   void* a20 = (void*)a2;
   values[2] = &a20; // how too pointer of pointer ???

   args[3] = &ffi_type_sint;
   values[3] = &a0;

   if (ffi_prep_cif(&cif, FFI_DEFAULT_ABI, 4, &ffi_type_void, args) == FFI_OK) {
       printf("hehehhee: %p\n", fn);
       int n = 0;
       printf("finish: %d, %d, %p, \n", rc, n, a2);
       printf("finish: %d, %d, %p, %s\n", rc, n, a2, a2[0]);
       // n = ((int (*)(int, int, int, int))(fn))(s, a1, &a2, a1); // ok
       ffi_call(&cif, fn, &rc, values);
       printf("finish: %d, %p\n", rc, n);
    }
}

extern void ffi_call_ex(void*fn, int retype, uint64_t *rc, int argc, uint8_t* argtys, uint64_t* argvals);

static void ffi_call_1(void*fn) {

    uint8_t argtys[20];
    uint64_t argvals[20];

    argtys[0] = FFI_TYPE_POINTER;
    void* o = calloc(1, 256);
    argvals[0] = (uint64_t)(&o);

    int argc = 2;
    argtys[1] = FFI_TYPE_POINTER;
    argvals[1] = (uint64_t)(&argc);

    char *a2[] = {"testprog", "hjdskkk"};
    argtys[2] = FFI_TYPE_POINTER;
    argvals[2] = (uint64_t)(void*)(a2);

    int flag = 0;
    argtys[3] = FFI_TYPE_INT;
    argvals[3] = (uint64_t)(&flag);

    uint64_t retval;
    ffi_call_ex(fn, FFI_TYPE_VOID, &retval, 4, argtys, argvals);
}

*/
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/gonuts/ffi"
)

func itype2stype(itype byte) *C.ffi_type {
	switch itype {
	case FFI_TYPE_VOID:
		return &C.ffi_type_void
	case FFI_TYPE_POINTER:
		return &C.ffi_type_pointer
	case FFI_TYPE_INT:
		return &C.ffi_type_sint
	case FFI_TYPE_FLOAT:
		return &C.ffi_type_float
	case FFI_TYPE_DOUBLE:
		return &C.ffi_type_double
	case FFI_TYPE_SINT16:
		return &C.ffi_type_sint16
	case FFI_TYPE_SINT32:
		return &C.ffi_type_sint32
	case FFI_TYPE_SINT64:
		return &C.ffi_type_sint64
	default:
		log.Println("unknown type:", itype)
		break
	}
	return &C.ffi_type_void
}

/*
TODO
  argtypes int[20]
  argvals uint64_t[20]
*/
func ffi_call_ex(fn unsafe.Pointer, retype int, rc *uint64, argc int, argtys []int, argvals []uint64) {

	var cif C.ffi_cif
	var ffitys = make([]*C.ffi_type, 20)
	var ffivals = make([]unsafe.Pointer, 20)
	var ffirc C.ffi_arg
	_, _, _, _ = cif, ffitys, ffivals, ffirc

	for i := 0; i < argc; i++ {
		switch argtys[i] {
		case C.FFI_TYPE_VOID:
			ffitys[i] = &C.ffi_type_void
			ffivals[i] = nil
		}
	}

	// C.ffi_call(&cif, fn, &ffirc, ffivals)
}

func InvokeQtFuncByName(symname string, args []uint64, types []int) uint64 {
	return 0
}

////////

type Type = ffi.Type
type VRetype = uint64 // interface{}

var debugFFICall = false
var libs = map[string]ffi.Library{}

func SetDebugFFICall(on bool) { debugFFICall = on }
func init() {
	init_invoke()
	init_callack_inherit()
}

func init_invoke() {
	loadModule := func(libpath string, modname string) error {
		var err error
		var lib ffi.Library
		lib, err = ffi.NewLibrary(libpath)
		ErrPrint(err, lib)
		// log.Println(lib)
		if err == nil {
			libs[modname] = lib
		}
		return err
	}

	dirp := "/usr/lib"
	switch runtime.GOOS {
	case "android":
		for i := 1; i < 9; i++ {
			d := fmt.Sprintf("/data/app/org.qtproject.example.go-%d/lib/arm", i)
			if FileExist(d) {
				dirp = d
				break
			}
		}
	}

	for _, modname := range []string{"Core", "Gui", "Widgets", "Network", "Inline"} {
		libpath := fmt.Sprintf("%s/libQt5%s.so", dirp, modname)
		loadModule(libpath, modname)
	}

	// loadModule("./tsym.so", "tsym")
}

func deinit() {}

func InvokeQtFunc(symname string, retype byte, types []byte, args ...interface{}) (VRetype, error) {

	for modname, lib := range libs {
		addr, err := lib.Symbol(symname)
		ErrPrint(err)
		if err != nil {
			continue
		}

		log.Println("FFI Call:", modname, symname, addr)
		// C.ffi_call_0(addr)
		C.ffi_call_1(addr)
		return 0, nil
	}
	return 0, fmt.Errorf("Symbol not found: %s", symname)
}

func InvokeQtFunc5(symname string, retype byte, argc int, types []byte, args []uint64) (VRetype, error) {
	addr := GetQtSymAddr(symname)
	log.Println("FFI Call:", symname, addr)

	var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), &retval, C.int(argc),
		(*C.uint8_t)(&types[0]), (*C.uint64_t)(&args[0]))

	return uint64(retval), fmt.Errorf("Symbol not found: %s", symname)
}

func InvokeQtFunc6(symname string, retype byte, args ...interface{}) (VRetype, error) {
	rawname := symname
	addr := GetQtSymAddr(symname)
	if debugFFICall {
		log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))
	}

	argtys, argvals := convArgs(args...)
	var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), &retval, C.int(len(argtys)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0]))

	onCtorAlloc(rawname)
	return uint64(retval), nil
}

// fix return QSize like pure record, RVO
func InvokeQtFunc7(symname string, args ...interface{}) (VRetype, error) {
	addr := GetQtSymAddr(symname)
	var retype byte = FFI_TYPE_POINTER
	log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))

	var retval unsafe.Pointer = C.calloc(1, 256)
	argtys, argvals := convArgs(args...)
	// var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), (*C.uint64_t)(retval), C.int(len(argtys)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0]))

	return uint64(uintptr(retval)), nil
}

func isUndefinedSymbolErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), ": undefined symbol: ")
}

// 直接使用封装的C++ symbols。好像在这设置没有用啊，符号不同，因为参数表的处理也不同，还是要改生成的调用代码。
var UseWrapSymbols bool = true // see also qtrt.UseCppSymbols TODO merge

func refmtSymbolName(symname string) string {
	return IfElseStr(UseWrapSymbols && strings.HasPrefix(symname, "_Z"), "C"+symname, symname)
}

func GetQtSymAddr(symname string) unsafe.Pointer {
	symname = refmtSymbolName(symname)
	for _, lib := range libs {
		addr, err := lib.Symbol(symname)
		if !isUndefinedSymbolErr(err) {
			ErrPrint(err, "")
		}
		if err != nil {
			continue
		}
		return addr
	}
	log.Println(fmt.Errorf("Symbol not found: %s", symname))
	return nil
}

// TODO
// get method symbol via virtual table offset
// ptr is class instance ptr
// midx is virtual method offset
// bidx is virtual base class offset
// return is the virtual method function pointer
func getSymByVTable(ptr unsafe.Pointer, midx int, bidx int) unsafe.Pointer {
	return ptr
}

func convArgs(args ...interface{}) (argtys []byte, argvals []uint64) {
	argtys = make([]byte, 20)
	argvals = make([]uint64, 20)
	for i, argx := range args {
		argty, argval := convArg(argx)
		argtys[i], argvals[i] = argty, argval
	}
	return
}

func convArg(argx interface{}) (argty byte, argval uint64) {

	av := reflect.ValueOf(argx)
	aty := av.Type()

	switch aty.Kind() {
	case reflect.Uint64:
		argty = FFI_TYPE_UINT64
		argval = uint64(argx.(uint64))
	case reflect.Int:
		argty = FFI_TYPE_INT
		argval = uint64(argx.(int))
	case reflect.Uint:
		argty = FFI_TYPE_UINT32
		argval = uint64(argx.(uint))
	case reflect.Bool:
		argty = FFI_TYPE_INT
		argval = uint64(IfElseInt(argx.(bool), 1, 0))
	case reflect.Ptr:
		argty = FFI_TYPE_POINTER
		argval = uint64(av.Pointer())
	case reflect.Float64:
		argty = FFI_TYPE_DOUBLE
		tv := uint64(0)
		*(*float64)(unsafe.Pointer(&tv)) = argx.(float64)
		argval = tv
	case reflect.UnsafePointer:
		argty = FFI_TYPE_POINTER
		argval = uint64(av.Pointer())
	default:
		log.Println("Unknown type:", argty, argval, aty.String(), argx)
	}

	return
}

func convRetval(retype byte, retval interface{}) interface{} {
	refv := reflect.ValueOf(retval)
	switch retype {
	case FFI_TYPE_VOID:
	case FFI_TYPE_INT:
		return refv.Convert(IntTy).Interface()
	case FFI_TYPE_UINT8:
		return refv.Convert(Uint8Ty).Interface()
	default:
		log.Println("Unknown type:", refv.Type().String())
	}
	return retval
}

const (
	FFI_TYPE_VOID       = byte(C.FFI_TYPE_VOID)
	FFI_TYPE_INT        = byte(C.FFI_TYPE_INT)
	FFI_TYPE_FLOAT      = byte(C.FFI_TYPE_FLOAT)
	FFI_TYPE_DOUBLE     = byte(C.FFI_TYPE_DOUBLE)
	FFI_TYPE_LONGDOUBLE = byte(C.FFI_TYPE_LONGDOUBLE)
	FFI_TYPE_UINT8      = byte(C.FFI_TYPE_UINT8)
	FFI_TYPE_SINT8      = byte(C.FFI_TYPE_SINT8)
	FFI_TYPE_UINT16     = byte(C.FFI_TYPE_UINT16)
	FFI_TYPE_SINT16     = byte(C.FFI_TYPE_SINT16)
	FFI_TYPE_UINT32     = byte(C.FFI_TYPE_UINT32)
	FFI_TYPE_SINT32     = byte(C.FFI_TYPE_SINT32)
	FFI_TYPE_UINT64     = byte(C.FFI_TYPE_UINT64)
	FFI_TYPE_SINT64     = byte(C.FFI_TYPE_SINT64)
	FFI_TYPE_STRUCT     = byte(C.FFI_TYPE_STRUCT)
	FFI_TYPE_POINTER    = byte(C.FFI_TYPE_POINTER)
	FFI_TYPE_COMPLEX    = byte(C.FFI_TYPE_COMPLEX)
)

// func KeepMe() {}

var ctorAllocStacks = map[string][]uintptr{}
var ctorAllocStacksMu sync.Mutex

func onCtorAlloc(symname string) {
	f := func(clsname string) {
		var pc [16]uintptr
		n := runtime.Callers(2, pc[:])
		_ = n
		ctorAllocStacksMu.Lock()
		ctorAllocStacks[clsname] = pc[:]
		ctorAllocStacksMu.Unlock()
	}

	if strings.Index(symname, "C2") > 0 {
		tmp1 := strings.Split(symname, "C2")[0]
		if strings.Index(tmp1, "Q") > 0 {
			tmp2 := strings.Split(tmp1, "Q")[1]
			clsname := "Q" + tmp2
			_ = clsname
			// log.Println("ctor alloc:", clsname)

			f(clsname)
		}
	}
}

// 奇怪了，正则怎么就让程序乱了呢？
func onCtorAlloc1(symname string) {
	reg := `_ZN(\d+)(Q.+)C2.*`
	exp := regexp.MustCompile(reg)
	mats := exp.FindAllStringSubmatch(symname, -1)
	if len(mats) > 0 {
		// var pc [16]uintptr
		// n := runtime.Callers(2, pc[:])
		// _ = n
		// log.Println("fill elems:", n, symname)
		// ctorAllocStacksMu.Lock()
		// ctorAllocStacks[mats[0][2]] = pc[:]
		// ctorAllocStacksMu.Unlock()
	} else {
		// log.Println("not match ctor: ", symname)
	}
}

func GetCtorAllocStack(clsname string) []uintptr {
	// ctorAllocStacksMu.Lock()
	// defer ctorAllocStacksMu.Unlock()
	if stk, ok := ctorAllocStacks[clsname]; ok {
		return stk
	}
	return nil
}

///
func test() {
	var ret VRetype
	var err error
	ret, err = InvokeQtFunc("_Z5qrandv", 0, nil)
	log.Println(ret, err)
	// ret, err = InvokeQtFunc("_Z6qsrandj", nil, nil)
	// log.Println(ret, err)
}
