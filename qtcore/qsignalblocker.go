package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QSignalBlocker struct {
	*qtrt.CObject
}
type QSignalBlocker_ITF interface {
	QSignalBlocker_PTR() *QSignalBlocker
}

func (ptr *QSignalBlocker) QSignalBlocker_PTR() *QSignalBlocker { return ptr }

func (this *QSignalBlocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSignalBlocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSignalBlockerFromPointer(cthis unsafe.Pointer) *QSignalBlocker {
	return &QSignalBlocker{&qtrt.CObject{cthis}}
}
func (*QSignalBlocker) NewFromPointer(cthis unsafe.Pointer) *QSignalBlocker {
	return NewQSignalBlockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobject.h:547
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSignalBlocker(QObject *)

/*

 */
func (*QSignalBlocker) NewForInherit(o QObject_ITF /*777 QObject **/) *QSignalBlocker {
	return NewQSignalBlocker(o)
}
func NewQSignalBlocker(o QObject_ITF /*777 QObject **/) *QSignalBlocker {
	var convArg0 unsafe.Pointer
	if o != nil && o.QObject_PTR() != nil {
		convArg0 = o.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSignalBlockerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalBlockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSignalBlocker)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:548
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSignalBlocker(QObject &)

/*

 */
func (*QSignalBlocker) NewForInherit1(o QObject_ITF) *QSignalBlocker {
	return NewQSignalBlocker1(o)
}
func NewQSignalBlocker1(o QObject_ITF) *QSignalBlocker {
	var convArg0 unsafe.Pointer
	if o != nil && o.QObject_PTR() != nil {
		convArg0 = o.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSignalBlockerC2ER7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSignalBlockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSignalBlocker)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:549
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QSignalBlocker()

/*

 */
func DeleteQSignalBlocker(this *QSignalBlocker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSignalBlockerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobject.h:556
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void reblock()

/*

 */
func (this *QSignalBlocker) Reblock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSignalBlocker7reblockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:557
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unblock()

/*

 */
func (this *QSignalBlocker) Unblock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSignalBlocker7unblockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
