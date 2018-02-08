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
import "gopp"
import "qt.go/qtrt"

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
func (this *QObjectCleanupHandler) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQObjectCleanupHandlerFromPointer(cthis unsafe.Pointer) *QObjectCleanupHandler {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QObjectCleanupHandler{bcthis0}
}
func (*QObjectCleanupHandler) NewFromPointer(cthis unsafe.Pointer) *QObjectCleanupHandler {
	return NewQObjectCleanupHandlerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:50
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QObjectCleanupHandler) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QObjectCleanupHandler10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QObjectCleanupHandler()
func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandlerC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQObjectCleanupHandlerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QObjectCleanupHandler()
func DeleteQObjectCleanupHandler(this *QObjectCleanupHandler) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandlerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:56
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * add(QObject *)
func (this *QObjectCleanupHandler) Add(object *QObject /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandler3addEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(QObject *)
func (this *QObjectCleanupHandler) Remove(object *QObject /*777 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandler6removeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:58
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QObjectCleanupHandler) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QObjectCleanupHandler7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QObjectCleanupHandler) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandler5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
