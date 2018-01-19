//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QFocusEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:389
// index:0
// void QFocusEvent(enum QEvent::Type, Qt::FocusReason)
func NewQFocusEvent(type_ int, reason int) *QFocusEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusEventC2EN6QEvent4TypeEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, cthis, &type_, &reason)
	gopp.ErrPrint(err, rv)
	return &QFocusEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:390
// index:0
// virtual
// void ~QFocusEvent()
func DeleteQFocusEvent(*QFocusEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFocusEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:392
// index:0
// inline
// bool gotFocus()
func (this *QFocusEvent) GotFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusEvent8gotFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:393
// index:0
// inline
// bool lostFocus()
func (this *QFocusEvent) LostFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusEvent9lostFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:395
// index:0
// Qt::FocusReason reason()
func (this *QFocusEvent) Reason() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFocusEvent6reasonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
