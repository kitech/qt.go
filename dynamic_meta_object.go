package qtmeta

/*
extern    void* DMetaObject_new(void* modat, void* castfn, void* callfn);
extern    void DMetaObject_delete(void* this_);

*/
import "C"
import "unsafe"

func DMetaObject_new(mod *QMetaObjectData, metacastfn, metacallfn unsafe.Pointer) unsafe.Pointer {
	rv := C.DMetaObject_new(unsafe.Pointer(mod), metacastfn, metacallfn)
	return rv
}

func DMetaObject_delete(mdo unsafe.Pointer) {
	C.DMetaObject_delete(mdo)
}
