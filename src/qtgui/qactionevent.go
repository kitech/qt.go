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
func NewQActionEventFromPointer(cthis unsafe.Pointer) *QActionEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QActionEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:727
// index:0
// Public
// void QActionEvent(int, class QAction *, class QAction *)
func NewQActionEvent(type_ int, action unsafe.Pointer /*666*/, before unsafe.Pointer /*666*/) *QActionEvent {
	cthis := qtrt.Calloc(1, 256) // 40
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionEventC2EiP7QActionS1_", ffiqt.FFI_TYPE_VOID, cthis, &type_, action, before)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:728
// index:0
// Public virtual
// void ~QActionEvent()
func DeleteQActionEvent(*QActionEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:730
// index:0
// Public inline
// QAction * action()
func (this *QActionEvent) Action() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionEvent6actionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qevent.h:731
// index:0
// Public inline
// QAction * before()
func (this *QActionEvent) Before() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionEvent6beforeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

//  body block end
