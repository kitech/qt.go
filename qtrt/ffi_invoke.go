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
       printf("finish: %d, %lld, %p, \n", (int)rc, n, a2);
       printf("finish: %d, %lld, %p, %s\n", (int)rc, n, a2, a2[0]);
       // n = ((int (*)(int, int, int, int))(fn))(s, a1, &a2, a1); // ok
       ffi_call(&cif, fn, &rc, values);
       printf("finish: %d, %p\n", (int)rc, (void*)n);
    }
}

extern void ffi_call_ex(void*fn, int retype, uint64_t *rc, int argc, uint8_t* argtys, uint64_t* argvals);
extern void ffi_call_ex3(void*fn, int retype, uint64_t *rc, int argc, void** argtys, void** argvals);
extern void ffi_call_var_ex(void*fn, int retype, uint64_t *rc, int fixedargc, int totalargc, uint8_t* argtys, uint64_t* argvals);
extern void set_so_ffi_call_ex(void* ex_fnptr, void* varex_fnptr, void* ex3_fnptr);
extern void ffi_call_ex_asmcc(); // just function name is fine, ignore parameters
extern void ffi_call_ex3_asmcc(); // just function name is fine, ignore parameters
extern void ffi_call_var_ex_asmcc();

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
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/kitech/dl/asmcgocall"
)

//////// TODO seems define as uint64 is better
type VRetype = uint64 // interface{}

var debugFFICall = false
var qtlibs = map[string]FFILibrary{}
var allsubmods = map[string]int{"Inline": 1, "Core": 1, "Gui": 1, "Widgets": 1, "Network": 1, "Qml": 1, "Quick": 1, "QuickControls2": 1, "QuickWidgets": 1}

// TODO move to internal package
// call in qtcore/qtgui/... init() function
// 这样就能知道import了哪些qt包了，选择性加载对应的so文件，而不用在这写死加载哪些so文件
func RegisterSubPackage(name string) {
	init_ffi_invoke(name)
}
func SetDebugFFICall(on bool) { debugFFICall = on }

func init() {
	dbgval := os.Getenv("QTGO_DEBUG_FFI_CALL")
	if strings.ToLower(dbgval) == "true" || dbgval == "1" {
		debugFFICall = true
	}
}
func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	init_ffi_invoke("")
	init_so_ffi_call()

	// TODO maybe run when first qtcall
	init_destroyedDynSlot()
	init_callack_inherit()
}

// 可用的参数值, "", Inline/Core/Gui/Widgets/Network/...
func init_ffi_invoke(libname string) {
	if libname == "" {
		libname = "Inline"
	}
	if _, ok := allsubmods[libname]; !ok {
		log.Fatalln("Unknown submod", libname)
	}
	if _, ok := qtlibs[libname]; ok {
		log.Println("Already loaded???", libname)
		return
	}
	// log.Println("Loading...", libname, len(qtlibs))

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
			sepos := strings.Index(appdir, ":") // for service process name
			if sepos != -1 {
				appdir = appdir[:sepos]
			}

			for i := 0; i < 9; i++ {
				d := fmt.Sprintf("/data/app/%s-%s/lib/%s/", appdir,
					IfElseStr(i == 0, "", fmt.Sprintf("%d", i)), archs[runtime.GOARCH])
				if FileExist(d) {
					return d
				}
			}
			dirs, err := filepath.Glob(fmt.Sprintf("/data/app/%s-*", appdir))
			ErrPrint(err)
			if len(dirs) > 0 {
				return dirs[0] + fmt.Sprintf("/lib/%s/", archs[runtime.GOARCH])
			}
			if FileExist(fmt.Sprintf("/data/data/%s/lib/", appdir)) {
				return fmt.Sprintf("/data/data/%s/lib/", appdir)
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

	loadModule := func(libpath string, modname string) error {
		var err error
		var lib FFILibrary
		lib, err = NewFFILibrary(libpath)
		ErrPrint(err, lib, libpath, modname)
		if err == nil {
			qtlibs[modname] = lib
		}
		return err
	}

	mods := []string{libname}
	// TODO auto check static and omit load other module
	if false && !UseWrapSymbols { // raw c++ symbols
		mods = append([]string{"Core", "Gui", "Widgets", "Network", "Qml", "Quick", "QuickControls2", "QuickWidgets"}, mods...)
	}

	for _, modname := range mods {
		libpath := getLibFile(getLibDirp(), modname)
		loadModule(libpath, modname)
	}
}

// load ffi call wrapper from libQt5Inline.so
func init_so_ffi_call() {
	ex_fnptr := GetQtSymAddrRaw("ffi_call_ex")
	ex3_fnptr := GetQtSymAddrRaw("ffi_call_ex3")
	varex_fnptr := GetQtSymAddrRaw("ffi_call_var_ex")
	if ex_fnptr != nil && varex_fnptr != nil {
		C.set_so_ffi_call_ex(ex_fnptr, varex_fnptr, ex3_fnptr)
	}
}

func deinit() {}

// deprecated
func InvokeQtFunc(symname string, retype byte, types []byte, args ...interface{}) (VRetype, error) {
	panic("deprecated " + symname)
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

// deprecated
func InvokeQtFunc5(symname string, retype byte, argc int, types []byte, args []uint64) (VRetype, error) {
	panic("deprecated " + symname)
	addr := GetQtSymAddr(symname)
	log.Println("FFI Call:", symname, addr)

	var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), &retval, C.int(argc),
		(*C.uint8_t)(&types[0]), (*C.uint64_t)(&args[0]))

	return uint64(retval), fmt.Errorf("Symbol not found: %s", symname)
}

// lagecy
func InvokeQtFunc6(symname string, retype byte, args ...interface{}) (VRetype, error) {
	addr := GetQtSymAddr(symname)
	if debugFFICall {
		log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))
	}

	return cc0byaddr(symname, addr, retype, args...)
}

// lagecy
// for variadic function call
func InvokeQtFunc6Var(symname string, retype byte, fixedargc int, args ...interface{}) (VRetype, error) {
	addr := GetQtSymAddr(symname)
	if debugFFICall {
		log.Println("FFI Call:", symname, addr, "retype=", retype, "fixedargc=", fixedargc, "totalargc=", len(args))
	}

	argtys, argvals, argrefps := convArgs(args...)
	_ = argrefps
	var retval C.uint64_t = 0
	_, cok := C.ffi_call_var_ex(addr, C.int(retype), &retval, C.int(fixedargc), C.int(len(args)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0]))
	if debugFFICall {
		ErrPrint(cok, symname, retype, len(args))
	}

	onCtorAlloc(symname)
	return uint64(retval), nil
}

// fix return QSize like pure record, RVO
func InvokeQtFunc7(symname string, args ...interface{}) (VRetype, error) {
	addr := GetQtSymAddr(symname)
	var retype byte = FFITY_POINTER
	log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))

	var retval Voidptr = C.calloc(1, 256)
	argtys, argvals, argrefps := convArgs(args...)
	_ = argrefps
	// var retval C.uint64_t = 0
	C.ffi_call_ex(addr, C.int(retype), (*C.uint64_t)(retval), C.int(len(args)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0]))
	return uint64(uintptr(retval)), nil
}

// dont use C_xxx, high level process sret, this,
func Qtcc0(symname string, retype byte, args ...interface{}) (VRetype, error) {
	addr := GetQtSymAddrRaw(symname)
	if debugFFICall {
		log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))
	}

	return cc0byaddr2(symname, addr, retype, args...)
}

// with symbol cache version
// dont use C_xxx, high level process sret, this,
func Qtcc1(symcrc uint32, symname string, retype byte, args ...interface{}) (VRetype, error) {
	addr := getSymAddrRawCached(symcrc, symname)
	if debugFFICall {
		log.Println("FFI Call:", symname, addr, "retype=", retype, "argc=", len(args))
	}
	//

	return cc0byaddr2(symname, addr, retype, args...)
}

// args, half is ffi_type**, half is argvals
// 直接传递类型对象，而非类型常量，不需要再做任何转换
// Go的slice 的数据指针和C兼容
func Qtcc3(symcrc uint32, symname string, retype Voidptr, args ...Voidptr) (VRetype, error) {
	addr := getSymAddrRawCached(symcrc, symname)
	if debugFFICall {
		log.Println("FFI Call:", symcrc, symname, addr, "retype=", retype, "argc=", len(args)/2)
	}

	if len(args)%2 != 0 {
		log.Println("Invalid parameters", len(args))
	}
	argc := len(args) / 2
	var argtys *C.uintptr_t = nil
	var argvals *C.uintptr_t = nil
	if argc > 0 {
		argtys = (*C.uintptr_t)(Voidptr(&args[0]))
		argvals = (*C.uintptr_t)(Voidptr(&args[argc]))
	}
	var retval C.uint64_t = 0
	var argv = struct {
		addr    Voidptr
		retype  Voidptr
		retval  *C.uint64_t
		argc    C.int
		argtys  *C.uintptr_t
		argvals *C.uintptr_t
	}{addr, retype, &retval, C.int(argc), argtys, argvals}
	asmcgocall.Asmcc(C.ffi_call_ex3_asmcc, Voidptr(&argv))

	onCtorAlloc(symname)

	return uint64(retval), nil
}

func getSymAddrRawCached(symcrc uint32, symname string) Voidptr {
	var addr Voidptr
	addrx, ok := symcache.Load(symcrc)
	if ok {
		addr = addrx.(Voidptr)
	} else {
		addr = GetQtSymAddrRaw(symname)
		if addr != nil {
			symcache.Store(symcrc, addr)
		}
	}
	if addr == nil {
		log.Panicln("Cannot get symbol addr", symcrc, symname)
	}
	return addr
}

// lagecy
// compatiable with old version
// symname just for debug
func cc0byaddr(symname string, symaddr Voidptr, retype byte, args ...interface{}) (VRetype, error) {
	addr := symaddr

	argtys, argvals, argrefps := convArgs(args...)
	_ = argrefps
	var retval C.uint64_t = 0
	_, cok := C.ffi_call_ex(addr, C.int(retype), &retval, C.int(len(args)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0]))
	if debugFFICall {
		ErrPrint(cok, symname, retype, len(args))
	}

	onCtorAlloc(symname)
	return uint64(retval), nil
}

// this version is 1.3x faster than cc0byaddr
// symname just for debug
func cc0byaddr2(symname string, symaddr Voidptr, retype byte, args ...interface{}) (VRetype, error) {
	addr := symaddr

	// 600ns/op => 400ns/op
	// argtys, argvals, argrefps := convArgs(args...)
	ap := argpp.Get()
	if ap == nil {
		log.Panicln("empty arg pack pool")
	}
	convArgs2(ap, args...)
	argtys, argvals, argrefps := ap.argtys, ap.argvals, ap.argrefps

	_ = argrefps
	var retval C.uint64_t = 0
	var argv = struct {
		addr    Voidptr
		retype  C.int
		retval  *C.uint64_t
		argc    C.int
		argtys  *C.uint8_t
		argvals *C.uint64_t
	}{addr, C.int(retype), &retval, C.int(len(args)),
		(*C.uint8_t)(&argtys[0]), (*C.uint64_t)(&argvals[0])}
	asmcgocall.Asmcc(C.ffi_call_ex_asmcc, Voidptr(&argv))

	onCtorAlloc(symname)
	argpp.Put(ap)
	return uint64(retval), nil
}

// TODO resolve ffi parameters and then forward to C scope execute
// C scope receiver func: void(void*fnptr, uint64_t*retval, void* argtys, void* argvals)
func ForwardFFIFunc(pxysymname string, symname string, args ...interface{}) (VRetype, error) {
	return 0, nil
}

func isUndefinedSymbolErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), ": undefined symbol: ")
}
func isNotfoundSymbolErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Symbol not found:")
}

// 直接使用封装的C++ symbols。好像在这设置没有用啊，符号不同，因为参数表的处理也不同，还是要改生成的调用代码。
var UseWrapSymbols bool = false // see also qtrt.UseCppSymbols TODO merge

func refmtSymbolName(symname string) string {
	return IfElseStr(UseWrapSymbols && strings.HasPrefix(symname, "_Z"), "C"+symname, symname)
}

// lagecy
func GetQtSymAddr(symname string) Voidptr {
	symname = refmtSymbolName(symname)
	return GetQtSymAddrRaw(symname)
}

func GetQtSymAddrRaw(symname string) Voidptr {
	for _, lib := range qtlibs {
		addr, err := lib.Symbol(symname)
		if !isUndefinedSymbolErr(err) && !isNotfoundSymbolErr(err) {
			ErrPrint(err, lib.Name(), symname)
		}
		if err != nil {
			continue
		}
		return addr
	}
	log.Fatalln(fmt.Errorf("Symbol not found: %s in %d libs(s)", symname, len(qtlibs)))
	return nil
}

// TODO
// get method symbol via virtual table offset
// ptr is class instance ptr
// midx is virtual method offset
// bidx is virtual base class offset
// return is the virtual method function pointer
func getSymByVTable(ptr Voidptr, midx int, bidx int) Voidptr {
	return ptr
}

func convArgs(args ...interface{}) (argtys []byte, argvals []uint64, argrefps []*reflect.Value) {
	/*
		argtys = make([]byte, 20)
		argvals = make([]uint64, 20)
		argrefps = make([]*reflect.Value, 20)
	*/
	//*
	argtys2 := [20]byte{}
	argvals2 := [20]uint64{}
	argrefps2 := [20]*reflect.Value{}
	argtys = argtys2[:]
	argvals = argvals2[:]
	argrefps = argrefps2[:]
	//*/
	for i, argx := range args {
		argty, argval, argrefp := convArg(i, argx)
		argtys[i], argvals[i], argrefps[i] = argty, argval, &argrefp
	}
	return
}

func convArgs2(ap *argPack, args ...interface{}) {
	for i, argx := range args {
		argty, argval, argrefp := convArg(i, argx)
		ap.argtys[i], ap.argvals[i], ap.argrefps[i] = argty, argval, &argrefp
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
		argty = FFITY_POINTER
		argrefp = reflect.New(av.Type())
		argrefp.Elem().Set(av)
		argval = refvaluint64(&argrefp)

	case reflect.UnsafePointer:
		argty = FFITY_POINTER
		argrefp = reflect.New(VoidpTy())
		argrefp.Elem().Set(av)
		argval = refvaluint64(&argrefp)

	case reflect.String:
		argty = FFITY_POINTER
		argpv := Voidptr(C.CString(argx.(string))) // TODO free memory
		argrefp = reflect.New(VoidpTy())
		argrefp.Elem().Set(reflect.ValueOf(argpv))
		argval = refvaluint64(&argrefp)
		//

	default:
		log.Println("Unknown type:", argty, argval, idx, aty.String(), argx)
	}

	return
}

// emulate reflect.Value
type emuValue struct {
	typ *reflect.Value // placeholder struct pointer field
	ptr Voidptr
	uint8
}

// hacked replacement of flaged depcreated  reflect.Value.Unsafe.Pointer() and reflect.Value.Pointer()
func refvalptr(vp *reflect.Value) Voidptr      { return (*emuValue)(Voidptr(vp)).ptr }
func refvaluintptr(vp *reflect.Value) uintptr  { return uintptr(refvalptr(vp)) }
func refvaluint64(vp *reflect.Value) uint64    { return uint64(refvaluintptr(vp)) }
func refvalptr_(vp *reflect.Value) Voidptr     { return Voidptr(vp.UnsafeAddr()) }
func refvaluintptr_(vp *reflect.Value) uintptr { return uintptr(refvalptr(vp)) }
func refvaluint64_(vp *reflect.Value) uint64   { return uint64(refvaluintptr(vp)) }

func convRetval(retype byte, retval interface{}) interface{} {
	refv := reflect.ValueOf(retval)
	switch retype {
	case FFITY_VOID:
	case FFI_TYPE_INT:
		return refv.Convert(IntTy).Interface()
	case FFI_TYPE_UINT8:
		return refv.Convert(Uint8Ty).Interface()
	default:
		log.Println("Unknown type:", refv.Type().String())
	}
	return retval
}

// TODO deprecated
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
const (
	FFITY_VOID    = byte(C.FFI_TYPE_VOID)
	FFITY_POINTER = byte(C.FFI_TYPE_POINTER)
)

// 由于FFI调用实际使用的是对象类型，可以方便地直接使用，不需要转换
var (
	FFITO_VOID    = Voidptr(&C.ffi_type_void)
	FFITO_POINTER = Voidptr(&C.ffi_type_pointer)
	FFITO_INT     = Voidptr(&C.ffi_type_sint32)
	FFITO_FLOAT   = Voidptr(&C.ffi_type_float)
	FFITO_DOUBLE  = Voidptr(&C.ffi_type_double)
	FFITO_SINT16  = Voidptr(&C.ffi_type_sint16)
	FFITO_UINT16  = Voidptr(&C.ffi_type_uint16)
	FFITO_SINT32  = Voidptr(&C.ffi_type_sint32)
	FFITO_UINT32  = Voidptr(&C.ffi_type_uint32)
	FFITO_SINT64  = Voidptr(&C.ffi_type_sint64)
	FFITO_UINT64  = Voidptr(&C.ffi_type_uint64)
	//FFITO_STRUCT  = Voidptr(&C.ffi_type_struct)
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
