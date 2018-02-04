package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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

type QActionEvent struct {
	*qtcore.QEvent
}

func (this *QActionEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QActionEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQActionEventFromPointer(cthis unsafe.Pointer) *QActionEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QActionEvent{bcthis0}
}
func (*QActionEvent) NewFromPointer(cthis unsafe.Pointer) *QActionEvent {
	return NewQActionEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:727
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QActionEvent(int, QAction *, QAction *)
func NewQActionEvent(type_ int, action unsafe.Pointer /*666*/, before unsafe.Pointer /*666*/) *QActionEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionEventC2EiP7QActionS1_", ffiqt.FFI_TYPE_POINTER, type_, action, before)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQActionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:728
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QActionEvent()
func DeleteQActionEvent(this *QActionEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:730
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAction * action()
func (this *QActionEvent) Action() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionEvent6actionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qevent.h:731
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QAction * before()
func (this *QActionEvent) Before() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionEvent6beforeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

//  body block end
