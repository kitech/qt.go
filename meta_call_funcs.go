package qtmeta

/*
#include <stdint.h>

extern void qt_static_meta_call_cgo(void*, int, int, void*);
extern void* qt_meta_cast_cgo(void*, char*);
extern int qt_meta_call_cgo(void*, int, int, void*);
*/
import "C"
import (
	"log"
	"unsafe"
)

func staticMetaCallFn() unsafe.Pointer { return C.qt_static_meta_call_cgo }
func metaCastFn() unsafe.Pointer       { return C.qt_meta_cast_cgo }
func metaCallFn() unsafe.Pointer       { return C.qt_meta_call_cgo }

//export qt_static_meta_call_cgo
func qt_static_meta_call_cgo(_o unsafe.Pointer, _c C.int, _id C.int, _a unsafe.Pointer) {
	log.Println(_o, _c, _id, _a)
}

//export qt_meta_cast_cgo
func qt_meta_cast_cgo(_o unsafe.Pointer, _clname *C.char) unsafe.Pointer {
	log.Println(_o, _clname)
	return nil
}

//export qt_meta_call_cgo
func qt_meta_call_cgo(_o unsafe.Pointer, _c C.int, _id C.int, _a unsafe.Pointer) C.int {
	log.Println(_o, _c, _id, _a)
	return 0
}
