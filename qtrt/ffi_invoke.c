#include <stdio.h>
#include <stdint.h>
#include "ffi.h"

//
static ffi_type* itype2stype(int itype){
    switch (itype) {
    case FFI_TYPE_VOID:
        return &ffi_type_void;
    case FFI_TYPE_POINTER:
        return &ffi_type_pointer;
    case FFI_TYPE_INT:
        return &ffi_type_sint;
    case FFI_TYPE_FLOAT:
        return &ffi_type_float;
    case FFI_TYPE_DOUBLE:
        return &ffi_type_double;
    case FFI_TYPE_SINT8:
        return &ffi_type_sint8;
    case FFI_TYPE_UINT8:
        return &ffi_type_uint8;
    case FFI_TYPE_SINT16:
        return &ffi_type_sint16;
    case FFI_TYPE_UINT16:
        return &ffi_type_uint16;
    case FFI_TYPE_SINT32:
        return &ffi_type_sint32;
    case FFI_TYPE_UINT32:
        return &ffi_type_uint32;
    case FFI_TYPE_SINT64:
        return &ffi_type_sint64;
    case FFI_TYPE_UINT64:
        return &ffi_type_uint64;
    default:
        break;
    }
    return &ffi_type_void;
}

/*
  argtypes int[20]
  argvals uint64_t[20] it's should be arguments's address but store in uint64_t
 */
void ffi_call_ex_impl(void*fn, int retype, uint64_t* retval, int argc, uint8_t* argtys, uint64_t* argvals) {
    ffi_cif cif;
    ffi_type *ffitys[20] = {0};
    void *ffivals[20] = {0};

    for (int i = 0; i < argc; i++) {
        ffitys[i] = itype2stype(argtys[i]);
        ffivals[i] = (void*)argvals[i];
    }

    ffi_type* retyp = itype2stype(retype);
    if (ffi_prep_cif(&cif, FFI_DEFAULT_ABI, argc, retyp, ffitys) == FFI_OK) {
        ffi_call(&cif, fn, retval, ffivals);
    }
}

/*
  for variadic function call
  argtypes int[20]
  argvals uint64_t[20] it's should be arguments's address but store in uint64_t
*/
void ffi_call_var_ex_impl(void*fn, int retype, uint64_t* retval, int fixedargc, int totalargc,
                     uint8_t* argtys, uint64_t* argvals) {
    ffi_cif cif;
    ffi_type *ffitys[20] = {0};
    void *ffivals[20] = {0};

    for (int i = 0; i < totalargc; i++) {
        ffitys[i] = itype2stype(argtys[i]);
        ffivals[i] = (void*)argvals[i];
    }

    ffi_type* retyp = itype2stype(retype);
    if (ffi_prep_cif_var(&cif, FFI_DEFAULT_ABI, fixedargc, totalargc, retyp, ffitys) == FFI_OK) {
        ffi_call(&cif, fn, retval, ffivals);
    }
}

//
void (*ffi_call_ex_fnptr)(void*fn, int retype, uint64_t* retval, int argc,
                       uint8_t* argtys, uint64_t* argvals) = &ffi_call_ex_impl;
void (*ffi_call_ex3_fnptr)(void*fn, void* retype, uint64_t* retval, int argc,
                           void** argtys, void** argvals) = 0;
void (*ffi_call_var_ex_fnptr)(void*fn, int retype, uint64_t* retval, int fixedargc, int totalargc,
                           uint8_t* argtys, uint64_t* argvals) = &ffi_call_var_ex_impl;

void set_so_ffi_call_ex(void* ex_fnptr, void* varex_fnptr, void* ex3_fnptr) {
    if (ex_fnptr != 0 && varex_fnptr != 0) {
        ffi_call_ex_fnptr = ex_fnptr;
        ffi_call_var_ex_fnptr = varex_fnptr;
        ffi_call_ex3_fnptr = ex3_fnptr;
    }
}

void ffi_call_ex(void*fn, int retype, uint64_t* retval, int argc, uint8_t* argtys, uint64_t* argvals) {
    ffi_call_ex_fnptr(fn, retype, retval, argc, argtys, argvals);
}
void ffi_call_var_ex(void*fn, int retype, uint64_t* retval, int fixedargc, int totalargc,
                          uint8_t* argtys, uint64_t* argvals) {
    ffi_call_var_ex_fnptr(fn, retype, retval, fixedargc, totalargc, argtys, argvals);
}

void ffi_call_ex_asmcc(struct {void* fn; int retype; uint64_t* retval;
    int argc; uint8_t* argtys; uint64_t* argvals;} *ax) {
    //printf("[0] %d %d %lu\n", ax->argtys[0], ax->argtys[7], *((uint64_t*)ax->argtys));
    ffi_call_ex_fnptr(ax->fn, ax->retype, ax->retval,
                      ax->argc, ax->argtys, ax->argvals);
}

void ffi_call_ex3(void* fn, void* retype, uint64_t* retval,
                  int argc, void** argtys, void** argvals) {
    //printf("[0] fn %p %p %d\n", ax->fn, ax->retype, ax->argc);
    ffi_call_ex3_fnptr(fn, retype, retval, argc, argtys, argvals);
}

void ffi_call_ex3_asmcc(struct {void* fn; void* retype; uint64_t* retval;
    int argc; void** argtys; void** argvals;} *ax) {
    //printf("[0] fn %p %p %d\n", ax->fn, ax->retype, ax->argc);
    ffi_call_ex3_fnptr(ax->fn, ax->retype, ax->retval,
                       ax->argc, ax->argtys, ax->argvals);
}

void ffi_call_var_ex_asmcc(struct {void*fn; int retype; uint64_t* retval;
    int fixedargc; int totalargc; uint8_t* argtys; uint64_t* argvals;} *ax) {
    ffi_call_var_ex_fnptr(ax->fn, ax->retype, ax->retval,
                          ax->fixedargc, ax->totalargc, ax->argtys, ax->argvals);
}

///// ffi closure
typedef struct { void* cbfn; } cppvm_ffi_closure_header;
// can only binding one extra data, if need more, use struct
static void cppvm_ffi_closure_tmplfn(ffi_cif *cif, void *ret, void* args[],
                                     void* capdata) {
    printf("%s:%d: ok1 %p\n", __FILE__, __LINE__, capdata);
    cppvm_ffi_closure_header* cbhdr = (cppvm_ffi_closure_header*)capdata;
    void (*cbfn)(void*, void**) = cbhdr->cbfn;
    printf("%s:%d: ok1 %p\n", __FILE__, __LINE__, cbfn);
    cbfn(capdata, args);
    printf("%s:%d: ok1 %p\n", __FILE__, __LINE__, capdata);
}

// callback func: void (*)(void* capdata, void**argvals)
ffi_closure*
make_cppvm_ffi_closure(ffi_cif* cif, void** closfnaddr, void* capdata,
                       ffi_type* retype, int argc, ffi_type** argtys) {

    ffi_closure* closure = ffi_closure_alloc(sizeof(ffi_closure), closfnaddr);
    int prepok = ffi_prep_cif(cif, FFI_DEFAULT_ABI, argc, retype, argtys);
    int prepok2 = ffi_prep_closure_loc(closure, cif, cppvm_ffi_closure_tmplfn,
                                          capdata, *closfnaddr);
    printf("%s:%d: ok1 %d, ok2 %d\n", __FILE__, __LINE__, prepok, prepok2);
    return closure;
}

void release_ffi_closure(ffi_closure* clos) {
    ffi_closure_free(clos);
}

void test_call_empty_closure(void* closfn) {
    printf("%s:%d: ok1 %p\n", __FILE__, __LINE__, closfn);
    int val = 5678;
    ((void(*)())closfn)(&val, &val);
}
