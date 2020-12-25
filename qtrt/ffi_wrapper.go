package qtrt

/*
#include <stdint.h>
#include <ffi.h>

extern void ffi_call_ex(void*fn, int retype, uint64_t* retval, int argc, uint8_t* argtys, uint64_t* argvals);
extern void ffi_call_ex3(void* fn, void* retype, uint64_t* retval,
                  int argc, uintptr_t* argtys, uintptr_t* argvals);
extern void ffi_call_ex3_asmcc(); // just function name is fine, ignore parameters

extern ffi_closure*
make_cppvm_ffi_closure(ffi_cif* cif, void** closfnaddr, void* capdata,
                       void* retype, int argc, uintptr_t* argtys);
extern void cppvm_ffi_closure_callback(void* cthis, void** argvals);
extern void test_call_empty_closure(void* closfn);
*/
import "C"
import (
	"log"
	"reflect"
	"runtime"

	"github.com/kitech/dl/asmcgocall"
	dl "github.com/kitech/dl/dl2"
)

// Library is a dl-opened library holding the corresponding dl.Handle
type FFILibrary struct {
	name   string
	handle dl.Handle
}

// NewFFILibrary takes the library filename and returns a handle towards it.
func NewFFILibrary(libname string) (lib FFILibrary, err error) {
	//libname = get_lib_arch_name(libname)
	lib.name = libname
	lib.handle, err = dl.Open(libname, dl.Now)
	return
}

func (lib FFILibrary) Name() string { return lib.name }
func (lib FFILibrary) Close() error {
	return lib.handle.Close()
}

func (lib FFILibrary) Symbol(fctname string) (Voidptr, error) {
	//println("Fct(",fctname,")...")
	sym, err := lib.handle.Symbol(fctname)
	if err != nil {
		return nil, err
	}

	addr := Voidptr(sym)
	return addr, nil
}

/// above indeed just libdl wrapper
// also as callback data
type CppvmClosure struct {
	// callback for ffi
	cbfn    Voidptr // must first
	retype  Voidptr
	retval  C.uint64_t
	argc    int       // = len(argtys)
	argtys  []Voidptr // go pointer to c error
	argtys2 uintptr   // convert and save here
	clsname string
	mthname string

	upcbfn interface{} // go scope function

	cif         C.ffi_cif
	clos        *C.ffi_closure
	bound_fnptr Voidptr
}

func (this *CppvmClosure) Cif() Voidptr      { return Voidptr(this.clos.cif) }
func (this *CppvmClosure) OrigFunc() Voidptr { return Voidptr(this.clos.fun) }
func (this *CppvmClosure) UserData() Voidptr { return this.clos.user_data }
func (this *CppvmClosure) Func() Voidptr     { return this.bound_fnptr }

// TODO
func (this *CppvmClosure) Call(args ...interface{}) {
	if false {
		C.test_call_empty_closure(this.bound_fnptr)
		return
	}
	if true {
		var retval C.uint64_t = 1230
		argvals := []Voidptr{Voidptr(&retval), Voidptr(&retval)}
		log.Println("argaddr", Voidptr(&retval), this.bound_fnptr)
		if this.argc != len(argvals) {
			log.Panicln("only support 2 arg for test")
		}
		C.ffi_call_ex3(this.bound_fnptr, FFITO_POINTER, &retval,
			C.int(this.argc), //
			(*C.uintptr_t)((Voidptr(this.argtys2))),
			(*C.uintptr_t)((Voidptr(&argvals[0]))))
		log.Println("hhh")
	}

	// would callback to go, cannot use asmcgocall
	if false {
		var retval C.uint64_t = 0
		var argv = struct {
			addr    Voidptr
			retype  Voidptr
			retval  *C.uint64_t
			argc    C.int
			argtys  *C.uintptr_t
			argvals *C.uintptr_t
		}{this.bound_fnptr, FFITO_VOID, &retval, C.int(0), nil, nil}
		asmcgocall.Asmcc(C.ffi_call_ex3_asmcc, Voidptr(&argv))
	}
}

//export cppvm_ffi_closure_callback
func cppvm_ffi_closure_callback(cthis Voidptr, argvals *Voidptr) {
	log.Println("hehe", cthis, argvals)
	this := (*CppvmClosure)(cthis)
	log.Println(this.clsname, this.mthname, len(this.argtys), this.upcbfn)
	p2 := (*[1 << 30]Voidptr)(Voidptr(argvals))
	argc := this.argc
	for i := 0; i < argc; i++ {
		val := (*Voidptr)(p2[i])
		log.Println("arg", i, p2[i], *val, uintptr(*val))
	}
	log.Println("reflect.Call", this.upcbfn, reflect.ValueOf(this.upcbfn).String())
}

func NewCppvmClosure(clsname, mthname string, upcbfn interface{}) *CppvmClosure {
	this := &CppvmClosure{}
	this.clsname = clsname
	this.mthname = mthname
	this.upcbfn = upcbfn

	this.cbfn = C.cppvm_ffi_closure_callback

	// make_cppvm_ffi_closure(ffi_cif* cif, void** closfnaddr, void* capdata,
	// ffi_type* retype, int argc, ffi_type** argtys) {
	goproto2ffitype(upcbfn)
	argtys := []Voidptr{(FFITO_POINTER), (FFITO_POINTER)}
	argtys2 := (*C.uintptr_t)(Voidptr(&argtys[0]))
	this.argtys2 = (uintptr(Voidptr(argtys2)))
	this.argc = len(argtys)

	this.clos = C.make_cppvm_ffi_closure(&this.cif, &this.bound_fnptr, Voidptr(this),
		FFITO_POINTER, C.int(this.argc), argtys2)
	log.Println(clsname, mthname, this.clos)

	runtime.SetFinalizer(this, releaseCppvmClosure)
	return this
}

func goproto2ffitype(fn interface{}) (argtys []Voidptr) {
	fnty := reflect.TypeOf(fn)
	argc := fnty.NumIn()
	log.Println("argc", argc)
	for i := 0; i < argc; i++ {
		argty := fnty.In(i)
		log.Println("arg", i, argty)
		rety := FFITO_POINTER
		switch argty.Kind() {
		case reflect.Ptr:
		case reflect.Int:
			rety = FFITO_INT
		case reflect.Float64:
			rety = FFITO_DOUBLE
		default:
			log.Panicln("unknown", i, argty)
		}
		argtys = append(argtys, rety)
	}
	return
}

func releaseCppvmClosure(obj interface{}) {

}
