package qtmeta

/*
#include <stdint.h>

extern void qt_static_meta_call_cgo(void*, int, int, void*);
extern void* qt_meta_cast_cgo(void*, char*);
extern int qt_meta_call_cgo(void*, int, int, void*);
*/
import "C"
import (
	"fmt"
	"log"
	"reflect"
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
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
	objx, found := mdcache.GetC2go(_o)
	if !found {
		log.Println("coresponding goobj not found:", _o)
		return
	}

	tymeta := mdcache.GetByObj(objx)
	sigcnt := int(tymeta.SignalCount)

	switch _c {
	case InvokeMetaMethod:
		// _id 为method index，可能是signal，也可能是slot。signal序号在前，slot序号在后。
		if _id < sigcnt {
			sigfn := mdcache.GetSignal(objx, _id)
			retsx := reflect.ValueOf(sigfn).Call([]reflect.Value{})
			log.Println(retsx)
		} else {
			slotfn := mdcache.GetSlot(objx, _id-sigcnt)
			retsx := reflect.ValueOf(slotfn).Call([]reflect.Value{})
			log.Println(retsx)
		}

	case RegisterMethodArgumentMetaType:
		log.Println("not supported meta call:", _c, _CallTypeName(_c))

	case IndexOfMethod:
		log.Println("not supported meta call:", _c, _CallTypeName(_c))
		// arg0, result, 返回方法的索引序号
		// arg1, 方法指针。好像在这没法用。
	default:
		log.Println("not supported meta call:", _c, _CallTypeName(_c))
	}
}

//export qt_meta_cast_cgo
func qt_meta_cast_cgo(_o unsafe.Pointer, _clname *C.char) unsafe.Pointer {
	if _clname == nil {
		return nil
	}

	_clnamego := C.GoString(_clname)
	log.Println(_o, _clnamego)
	objx, found := mdcache.GetC2go(_o)
	if !found {
		log.Println("coresponding goobj not found:", _o)
		return nil
	}

	tymeta := mdcache.GetByObj(objx)
	if _clnamego == tymeta.ClassName {
		return _o
	}
	// return QThread::qt_metacast(_clname);
	bclsname := "xxxxx"
	symname := fmt.Sprintf("C_ZN%d%s11qt_metacastEPKc", len(bclsname), bclsname)
	rv, err := qtrt.InvokeQtFunc6(symname, qtrt.FFI_TYPE_POINTER, _o, _clname)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

//export qt_meta_call_cgo
func qt_meta_call_cgo(_o unsafe.Pointer, _c C.int, _id C.int, _a unsafe.Pointer) C.int {
	return C.int(qt_meta_call_go(_o, int(_c), int(_id), _a))
}
func qt_meta_call_go(_o unsafe.Pointer, _c int, _id int, _a unsafe.Pointer) C.int {
	log.Println(_o, _c, _id, _a)
	objx, found := mdcache.GetC2go(_o)
	if !found {
		log.Println("coresponding goobj not found:", _o)
		return 0
	}

	tymeta := mdcache.GetByObj(objx)
	mthcnt := int(tymeta.MethodCount)

	switch _c {
	case InvokeMetaMethod:
		// _id 为method index，可能是signal，也可能是slot。signal序号在前，slot序号在后。
		if _id < mthcnt {
			qt_static_meta_call_go(_o, _c, _id, _a)
		}
		_id -= mthcnt
	case RegisterMethodArgumentMetaType:
		if _id < mthcnt {
			qt_static_meta_call_go(_o, _c, _id, _a)
		}
		_id -= mthcnt
	default:
		log.Println("not supported meta call:", _c, _CallTypeName(_c))
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

func _CallTypeName(_c int) string {
	_CallTypes := map[int]string{
		InvokeMetaMethod:               "InvokeMetaMethod",
		ReadProperty:                   "ReadProperty",
		WriteProperty:                  "WriteProperty",
		ResetProperty:                  "ResetProperty",
		QueryPropertyDesignable:        "QueryPropertyDesignable",
		QueryPropertyScriptable:        "QueryPropertyScriptable",
		QueryPropertyStored:            "QueryPropertyStored",
		QueryPropertyEditable:          "QueryPropertyEditable",
		QueryPropertyUser:              "QueryPropertyUser",
		CreateInstance:                 "CreateInstance",
		IndexOfMethod:                  "IndexOfMethod",
		RegisterPropertyMetaType:       "RegisterPropertyMetaType",
		RegisterMethodArgumentMetaType: "RegisterMethodArgumentMetaType",
	}
	if name, ok := _CallTypes[_c]; ok {
		return name
	}
	return fmt.Sprintf("Unknown call: %d", _c)
}
