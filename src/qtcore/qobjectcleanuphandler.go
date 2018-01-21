package qtcore

// /usr/include/qt/QtCore/qobjectcleanuphandler.h
// #include <qobjectcleanuphandler.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QObjectCleanupHandler struct {
	*QObject
}

func (this *QObjectCleanupHandler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func NewQObjectCleanupHandlerFromPointer(cthis unsafe.Pointer) *QObjectCleanupHandler {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QObjectCleanupHandler{bcthis0}
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:50
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QObjectCleanupHandler) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QObjectCleanupHandler10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:53
// index:0
// Public
// void QObjectCleanupHandler()
func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QObjectCleanupHandlerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQObjectCleanupHandlerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:54
// index:0
// Public virtual
// void ~QObjectCleanupHandler()
func DeleteQObjectCleanupHandler(*QObjectCleanupHandler) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QObjectCleanupHandlerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:56
// index:0
// Public
// QObject * add(class QObject *)
func (this *QObjectCleanupHandler) Add(object *QObject /*444 QObject **/) *QObject /*444 QObject **/ {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QObjectCleanupHandler3addEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:57
// index:0
// Public
// void remove(class QObject *)
func (this *QObjectCleanupHandler) Remove(object *QObject /*444 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QObjectCleanupHandler6removeEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:58
// index:0
// Public
// bool isEmpty()
func (this *QObjectCleanupHandler) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QObjectCleanupHandler7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:59
// index:0
// Public
// void clear()
func (this *QObjectCleanupHandler) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QObjectCleanupHandler5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
