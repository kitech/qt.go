//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:727
// index:0
// void QActionEvent(int, class QAction *, class QAction *)
func NewQActionEvent(type_ int, action unsafe.Pointer, before unsafe.Pointer) *QActionEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionEventC2EiP7QActionS1_", ffiqt.FFI_TYPE_VOID, cthis, &type_, action, before)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionEventFromPointer(cthis)
	return gothis
}
func NewQActionEventFromPointer(cthis unsafe.Pointer) *QActionEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QActionEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:728
// index:0
// virtual
// void ~QActionEvent()
func DeleteQActionEvent(*QActionEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:730
// index:0
// inline
// QAction * action()
func (this *QActionEvent) Action() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionEvent6actionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:731
// index:0
// inline
// QAction * before()
func (this *QActionEvent) Before() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionEvent6beforeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
