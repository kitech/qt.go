//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QAccessibleTextCursorEvent struct {
	*QAccessibleEvent
}

func (this *QAccessibleTextCursorEvent) GetCthis() unsafe.Pointer {
	return this.QAccessibleEvent.GetCthis()
}

// /usr/include/qt/QtGui/qaccessible.h:747
// index:0
// inline
// void QAccessibleTextCursorEvent(class QObject *, int)
func NewQAccessibleTextCursorEvent(obj unsafe.Pointer, cursorPos int) *QAccessibleTextCursorEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventC2EP7QObjecti", ffiqt.FFI_TYPE_VOID, cthis, obj, &cursorPos)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextCursorEventFromPointer(cthis)
	return gothis
}
func NewQAccessibleTextCursorEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextCursorEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleTextCursorEvent{bcthis0}
}

// /usr/include/qt/QtGui/qaccessible.h:753
// index:1
// inline
// void QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
func NewQAccessibleTextCursorEvent_1(iface unsafe.Pointer, cursorPos int) *QAccessibleTextCursorEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei", ffiqt.FFI_TYPE_VOID, cthis, iface, &cursorPos)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextCursorEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:760
// index:0
// virtual
// void ~QAccessibleTextCursorEvent()
func DeleteQAccessibleTextCursorEvent(*QAccessibleTextCursorEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:762
// index:0
// inline
// void setCursorPosition(int)
func (this *QAccessibleTextCursorEvent) SetCursorPosition(position int) {
	// 0: (, position int), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEvent17setCursorPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:763
// index:0
// inline
// int cursorPosition()
func (this *QAccessibleTextCursorEvent) CursorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextCursorEvent14cursorPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
