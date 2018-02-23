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
       int64_t n = 0;
       printf("finish: %d, %ld, %p, \n", (int)rc, n, a2);
       printf("finish: %d, %ld, %p, %s\n", (int)rc, n, a2, a2[0]);
       // n = ((int (*)(int, int, int, int))(fn))(s, a1, &a2, a1); // ok
       ffi_call(&cif, fn, &rc, values);
       printf("finish: %d, %p\n", (int)rc, (void*)n);
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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

func itype2stype(itype byte) *C.ffi_type {
	switch itype {
	case FFI_TYPE_VOID:
		return &C.ffi_type_void
	case FFI_TYPE_POINTER:
		return &C.ffi_type_pointer
	case FFI_TYPE_INT:
		return &C.ffi_type_sint32
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
type VRetype = uint64 // interface{}

var debugFFICall = false
var qtlibs = map[string]FFILibrary{}

func SetDebugFFICall(on bool) { debugFFICall = on }
func init() {
	init_ffi_invoke()

	// TODO maybe run when first qtcall
	init_destroyedDynSlot()
	init_callack_inherit()
}

func init_ffi_invoke() {
	loadModule := func(libpath string, modname string) error {
		var err error
		var lib FFILibrary
		lib, err = NewFFILibrary(libpath)
		ErrPrint(err, lib)
		// log.Println(lib)
		if err == nil {
			qtlibs[modname] = lib
		}
		return err
	}

	// lib dir prefix
	// go arch name => android lib name
	archs := map[string]string{"386": "x86", "amd64": "x86_64", "arm": "arm", "mips": "mips"}
	oslibexts := map[string]string{"linux": "so", "darwin": "dylib", "windows": "dll"}

	getLibDirp := func() string {
		switch runtime.GOOS {
		case "android":
			bcc, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", os.Getpid()))
			ErrPrint(err)
			appdir := string(bcc[:bytes.IndexByte(bcc, 0)])
			for i := 0; i < 9; i++ {
				d := fmt.Sprintf("/data/app/%s%s/lib/%s/", appdir,
					IfElseStr(i == 0, "", fmt.Sprintf("-%d", i)), archs[runtime.GOARCH])
				if FileExist(d) {
					return d
				}
			}
		}
		return ""
	}
	// dirp must endsWith / or ""
	getLibFile := func(dirp, modname string) string {
		switch runtime.GOOS {
		case "darwin":
			return fmt.Sprintf("%slibQt5%s.%s", dirp, modname, oslibexts[runtime.GOOS])
		case "windows": // best put libs in current directory
			return fmt.Sprintf("%sQt5%s.%s", dirp, modname, oslibexts[runtime.GOOS])
		}
		// case "linux", "freebsd", "netbsd", "openbsd", "android", ...:
		return fmt.Sprintf("%slibQt5%s.%s", dirp, modname, oslibexts["linux"])
	}

	for _, modname := range []string{"Core", "Gui", "Widgets", "Network", "Inline",
		"Qml", "Quick", "QuickControls2", "QuickWidgets"} {
		libpath := getLibFile(getLibDirp(), modname)
		loadModule(libpath, modname)
	}
}

func deinit() {}

func InvokeQtFunc(symname string, retype byte, types []byte, args ...interface{}) (VRetype, error) {
	for modname, lib := range qtlibs {
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
	addr := GetQtSymAddr(symname)
	if debugFFICall {
		log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))
	}

	argtys, argvals, argrefps := convArgs(args...)
	_ = argrefps
	var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), &retval, C.int(len(args)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0]))

	onCtorAlloc(symname)
	return uint64(retval), nil
}

// fix return QSize like pure record, RVO
func InvokeQtFunc7(symname string, args ...interface{}) (VRetype, error) {
	addr := GetQtSymAddr(symname)
	var retype byte = FFI_TYPE_POINTER
	log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))

	var retval unsafe.Pointer = C.calloc(1, 256)
	argtys, argvals, argrefps := convArgs(args...)
	_ = argrefps
	// var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), (*C.uint64_t)(retval), C.int(len(args)),
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
	for _, lib := range qtlibs {
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

func convArgs(args ...interface{}) (argtys []byte, argvals []uint64, argrefps []*reflect.Value) {
	argtys = make([]byte, 20)
	argvals = make([]uint64, 20)
	argrefps = make([]*reflect.Value, 20)
	for i, argx := range args {
		argty, argval, argrefp := convArg(i, argx)
		argtys[i], argvals[i], argrefps[i] = argty, argval, &argrefp
	}
	return
}

var tyconvmap = map[reflect.Kind]byte{
	reflect.Uint64: FFI_TYPE_UINT64, reflect.Int64: FFI_TYPE_SINT64,
	reflect.Uint32: FFI_TYPE_UINT32, reflect.Int32: FFI_TYPE_SINT32,
	reflect.Uint: FFI_TYPE_UINT32, reflect.Int: FFI_TYPE_INT,
	reflect.Uint16: FFI_TYPE_UINT16, reflect.Int16: FFI_TYPE_SINT16,
	reflect.Uint8: FFI_TYPE_UINT8, reflect.Int8: FFI_TYPE_SINT8,
}

// argval should be the value's valid address
//   for non-addressable primitive type, a temporary var is created and it's address is returned
// argrefp for hold the temporary created var's address's reference, prevent gc for a while
func convArg(idx int, argx interface{}) (argty byte, argval uint64, argrefp reflect.Value) {

	av := reflect.ValueOf(argx)
	aty := av.Type()

	switch aty.Kind() {
	case reflect.Uint64, reflect.Int64, reflect.Uint32, reflect.Int32,
		reflect.Int, reflect.Uint, reflect.Uint16, reflect.Int16,
		reflect.Uint8, reflect.Int8:
		argty = tyconvmap[aty.Kind()]
		if av.CanAddr() {
			argrefp = av
			argval = refvaluint64(&argrefp)
		} else {
			argrefp = reflect.New(aty)
			argrefp.Elem().Set(av)
			argval = refvaluint64(&argrefp)
		}
	case reflect.Bool:
		argty = FFI_TYPE_INT
		argrefp = reflect.New(IntTy)
		argrefp.Elem().Set(reflect.ValueOf(IfElseInt(argx.(bool), 1, 0)))
		argval = refvaluint64(&argrefp)
	case reflect.Float64:
		argty = FFI_TYPE_DOUBLE
		if av.CanAddr() {
			argrefp = av
			argval = refvaluint64(&argrefp)
		} else {
			argrefp = reflect.New(Float64Ty)
			argrefp.Elem().Set(av)
			argval = refvaluint64(&argrefp)
		}

	case reflect.Float32:
		argty = FFI_TYPE_FLOAT
		if av.CanAddr() {
			argrefp = av
			argval = refvaluint64(&argrefp)
		} else {
			argrefp = reflect.New(Float32Ty)
			argrefp.Elem().Set(av)
			argval = refvaluint64(&argrefp)
		}

	case reflect.Ptr:
		argty = FFI_TYPE_POINTER
		argrefp = reflect.New(av.Type())
		argrefp.Elem().Set(av)
		argval = refvaluint64(&argrefp)

	case reflect.UnsafePointer:
		argty = FFI_TYPE_POINTER
		argrefp = reflect.New(VoidpTy())
		argrefp.Elem().Set(av)
		argval = refvaluint64(&argrefp)

	default:
		log.Println("Unknown type:", argty, argval, aty.String(), argx)
	}

	return
}

// emulate reflect.Value
type emuValue struct {
	typ *reflect.Value // placeholder struct pointer field
	ptr unsafe.Pointer
	uint8
}

// hacked replacement of flaged depcreated  reflect.Value.Unsafe.Pointer() and reflect.Value.Pointer()
func refvalptr(vp *reflect.Value) unsafe.Pointer  { return (*emuValue)(unsafe.Pointer(vp)).ptr }
func refvaluintptr(vp *reflect.Value) uintptr     { return uintptr(refvalptr(vp)) }
func refvaluint64(vp *reflect.Value) uint64       { return uint64(refvaluintptr(vp)) }
func refvalptr_(vp *reflect.Value) unsafe.Pointer { return unsafe.Pointer(vp.UnsafeAddr()) }
func refvaluintptr_(vp *reflect.Value) uintptr    { return uintptr(refvalptr(vp)) }
func refvaluint64_(vp *reflect.Value) uint64      { return uint64(refvaluintptr(vp)) }

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
	// FFI_TYPE_COMPLEX    = byte(C.FFI_TYPE_COMPLEX)
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
	ctorAllocStacksMu.Lock()
	defer ctorAllocStacksMu.Unlock()
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
