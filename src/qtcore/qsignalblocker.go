//  header block begin
// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>
package qtcore

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
type QSignalBlocker struct {
	*qtrt.CObject
}

func (this *QSignalBlocker) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSignalBlockerFromPointer(cthis unsafe.Pointer) *QSignalBlocker {
	return &QSignalBlocker{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qobject.h:547
// index:0
// Public inline
// void QSignalBlocker(class QObject *)
func NewQSignalBlocker(o unsafe.Pointer) *QSignalBlocker {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlockerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, o)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalBlockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:548
// index:1
// Public inline
// void QSignalBlocker(class QObject &)
func NewQSignalBlocker_1(o *QObject) *QSignalBlocker {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = o.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlockerC2ER7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSignalBlockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:549
// index:0
// Public inline
// void ~QSignalBlocker()
func DeleteQSignalBlocker(*QSignalBlocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlockerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:556
// index:0
// Public inline
// void reblock()
func (this *QSignalBlocker) Reblock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlocker7reblockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:557
// index:0
// Public inline
// void unblock()
func (this *QSignalBlocker) Unblock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSignalBlocker7unblockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
