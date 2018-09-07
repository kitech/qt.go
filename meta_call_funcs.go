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
	qt_static_meta_call_go(_o, int(_c), int(_id), _a)
}
func qt_static_meta_call_go(_o unsafe.Pointer, _c int, _id int, _a unsafe.Pointer) {
	log.Println(_o, _c, _id, _a)
	switch _c {
		case InvokeMetaMethod:
	}
}

//export qt_meta_cast_cgo
func qt_meta_cast_cgo(_o unsafe.Pointer, _clname *C.char) unsafe.Pointer {
	return qt_meta_cast_go(_o, C.GoString(_clname))
}
func qt_meta_cast_go(_o unsafe.Pointer, _clname string) unsafe.Pointer {
	log.Println(_o, _clname)
	return nil
}

//export qt_meta_call_cgo
func qt_meta_call_cgo(_o unsafe.Pointer, _c C.int, _id C.int, _a unsafe.Pointer) C.int {
	return C.int(qt_meta_call_go(_o, int(_c), int(_id), _a))
}
func qt_meta_call_go(_o unsafe.Pointer, _c int, _id int, _a unsafe.Pointer) C.int {
	log.Println(_o, _c, _id, _a)
	switch _c {
	case InvokeMetaMethod:
	}
	return 0
}

/////
type _CallType = int
const (
	InvokeMetaMethod _CallType = iota
	ReadProperty
	WriteProperty
	ResetProperty
	QueryPropertyDesignable
	QueryPropertyScriptable
	QueryPropertyStored
	QueryPropertyEditable
	QueryPropertyUser
	CreateInstance
	IndexOfMethod
	RegisterPropertyMetaType
	RegisterMethodArgumentMetaType
)
