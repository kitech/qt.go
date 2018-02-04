#include <stdint.h>
#include "ffi.h"

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
    case FFI_TYPE_SINT16:
        return &ffi_type_sint16;
    case FFI_TYPE_SINT32:
        return &ffi_type_sint32;
    case FFI_TYPE_SINT64:
        return &ffi_type_sint64;
    default:
        break;
    }
    return &ffi_type_void;
}

/*
  argtypes int[20]
  argvals uint64_t[20]
 */
void ffi_call_ex(void*fn, int retype, uint64_t *retval, int argc, uint8_t* argtys, uint64_t* argvals) {
    ffi_cif cif;
    ffi_type *ffitys[20];
    void *ffivals[20];

    for (int i = 0; i < argc; i++) {
        ffitys[i] = itype2stype(argtys[i]);
        ffivals[i] = (void*)&argvals[i];
    }

    ffi_type* retyp = itype2stype(retype);
    if (ffi_prep_cif(&cif, FFI_DEFAULT_ABI, argc, retyp, ffitys) == FFI_OK) {
        ffi_call(&cif, fn, retval, ffivals);
    }
}
