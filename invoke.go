package main

/*
#cgo CFLAGS: -I/usr/lib/libffi-3.2.1/include/
#cgo LDFLAGS: -lffi

///////////
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

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

extern void ffi_call_ex(void*fn, int retype, uint64_t *rc, int argc, int* argtys, uint64_t* argvals);

static void ffi_call_1(void*fn) {

    int argtys[20];
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
	"gopp"
	"log"
	"unsafe"

	"github.com/gonuts/ffi"
)

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
type VRetype = interface{}

var libs = map[string]ffi.Library{}

func init() {
	loadModule := func(libpath string, modname string) {
		var err error
		var lib ffi.Library
		lib, err = ffi.NewLibrary(libpath)
		gopp.ErrPrint(err, lib)
		// log.Println(lib)
		if err == nil {
			libs[modname] = lib
		}
	}

	for _, modname := range []string{"Core", "Gui", "Widgets", "Network", "Inline"} {
		libpath := fmt.Sprintf("/usr/lib/libQt5%s.so", modname)
		loadModule(libpath, modname)
	}

	loadModule("./tsym.so", "tsym")
}

func deinit() {}

func InvokeQtFunc(symname string, retype Type, types []Type, args ...interface{}) (VRetype, error) {
	if retype == nil {
		retype = ffi.C_void
		retype = ffi.C_uint64
	}
	for modname, lib := range libs {
		addr, err := lib.Symbol(symname)
		gopp.ErrPrint(err)
		if err != nil {
			continue
		}

		log.Println("FFI Call:", modname, symname, addr)
		// C.ffi_call_0(addr)
		C.ffi_call_1(addr)
		return nil, nil
	}
	return nil, fmt.Errorf("Symbol not found: %s", symname)
}

func test() {
	var ret VRetype
	var err error
	ret, err = InvokeQtFunc("_Z5qrandv", nil, nil)
	log.Println(ret, err)
	// ret, err = InvokeQtFunc("_Z6qsrandj", nil, nil)
	// log.Println(ret, err)
}
