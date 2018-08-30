package qtmeta

/*
#include <string.h>
#include <stdlib.h>

#include "dynamic_meta_object.h"

extern    void* DMetaObject_new(void* modat, void* castfn, void* callfn);
extern    void DMetaObject_delete(void* this_);

*/
import "C"
import "unsafe"

func DMetaObject_new(mod *QMetaObjectData2, metacastfn, metacallfn unsafe.Pointer) unsafe.Pointer {
	rv := C.DMetaObject_new(unsafe.Pointer(mod.ptr), metacastfn, metacallfn)
	return rv
}

func DMetaObject_delete(mdo unsafe.Pointer) {
	C.DMetaObject_delete(mdo)
}

type QMetaObjectData2 struct {
	ptr *C.QMetaObjectData
}

func NewQMetaObjectData2() *QMetaObjectData2 {
	mod2 := &QMetaObjectData2{}
	mod2.ptr = (*C.QMetaObjectData)(C.calloc(1, C.sizeof_QMetaObjectData))
	mod2.ptr.superdata = unsafe.Pointer(nil)
	mod2.ptr.stringdata = unsafe.Pointer(nil)
	mod2.ptr.data = unsafe.Pointer(nil)
	mod2.ptr.static_metacall = unsafe.Pointer(nil)

	return mod2
}

func (this *QMetaObjectData2) SetData(superdata, stringdata, data, static_metacall unsafe.Pointer) *QMetaObjectData2 {
	mod2 := this
	mod2.ptr.superdata = superdata
	mod2.ptr.stringdata = stringdata
	mod2.ptr.data = data
	mod2.ptr.static_metacall = static_metacall

	return this // usage: NewQMetaObjectData2().SetData()
}

func (this *QMetaObjectData2) VoidPtr() unsafe.Pointer { return unsafe.Pointer(this.ptr) }
