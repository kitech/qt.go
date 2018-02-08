package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QFocusEvent struct {
	*qtcore.QEvent
}

func (this *QFocusEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QFocusEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQFocusEventFromPointer(cthis unsafe.Pointer) *QFocusEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QFocusEvent{bcthis0}
}
func (*QFocusEvent) NewFromPointer(cthis unsafe.Pointer) *QFocusEvent {
	return NewQFocusEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:389
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFocusEvent(enum QEvent::Type, Qt::FocusReason)
func NewQFocusEvent(type_ int, reason int) *QFocusEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusEventC2EN6QEvent4TypeEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, type_, reason)
	gopp.ErrPrint(err, rv)
	gothis := NewQFocusEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFocusEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:390
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFocusEvent()
func DeleteQFocusEvent(this *QFocusEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:392
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool gotFocus()
func (this *QFocusEvent) GotFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusEvent8gotFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:393
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool lostFocus()
func (this *QFocusEvent) LostFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusEvent9lostFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:395
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FocusReason reason()
func (this *QFocusEvent) Reason() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusEvent6reasonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

//  body block end
