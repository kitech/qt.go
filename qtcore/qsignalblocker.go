package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSignalBlocker struct {
	*qtrt.CObject
}

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
func NewQSignalBlocker(o *QObject /*777 QObject **/) *QSignalBlocker {
	var convArg0 = o.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlockerC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalBlockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSignalBlocker)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:548
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSignalBlocker(QObject &)
func NewQSignalBlocker_1(o *QObject) *QSignalBlocker {
	var convArg0 = o.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlockerC2ER7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalBlockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSignalBlocker)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:549
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QSignalBlocker()
func DeleteQSignalBlocker(this *QSignalBlocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlockerD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobject.h:556
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void reblock()
func (this *QSignalBlocker) Reblock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlocker7reblockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:557
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unblock()
func (this *QSignalBlocker) Unblock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlocker7unblockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
