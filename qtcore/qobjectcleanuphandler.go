package qtcore

// /usr/include/qt/QtCore/qobjectcleanuphandler.h
// #include <qobjectcleanuphandler.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QObjectCleanupHandler struct {
	*QObject
}
type QObjectCleanupHandler_ITF interface {
	QObject_ITF
	QObjectCleanupHandler_PTR() *QObjectCleanupHandler
}

func (ptr *QObjectCleanupHandler) QObjectCleanupHandler_PTR() *QObjectCleanupHandler { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QObjectCleanupHandler) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QObjectCleanupHandler10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QObjectCleanupHandler()

/*
Constructs an empty QObjectCleanupHandler.
*/
func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandlerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQObjectCleanupHandlerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QObjectCleanupHandler")
	return gothis
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QObjectCleanupHandler()

/*

 */
func DeleteQObjectCleanupHandler(this *QObjectCleanupHandler) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandlerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:56
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * add(QObject *)

/*
Adds object to this cleanup handler and returns the pointer to the object.

See also remove().
*/
func (this *QObjectCleanupHandler) Add(object QObject_ITF /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandler3addEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(QObject *)

/*
Removes the object from this cleanup handler. The object will not be destroyed.

See also add().
*/
func (this *QObjectCleanupHandler) Remove(object QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandler6removeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:58
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if this cleanup handler is empty or if all objects in this cleanup handler have been destroyed; otherwise return false.

See also add(), remove(), and clear().
*/
func (this *QObjectCleanupHandler) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QObjectCleanupHandler7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectcleanuphandler.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Deletes all objects in this cleanup handler. The cleanup handler becomes empty.

See also isEmpty().
*/
func (this *QObjectCleanupHandler) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QObjectCleanupHandler5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
