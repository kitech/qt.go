package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
	if this == nil {
		return nil
	} else {
		return this.QAccessibleEvent.GetCthis()
	}
}
func NewQAccessibleTextCursorEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextCursorEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleTextCursorEvent{bcthis0}
}

// /usr/include/qt/QtGui/qaccessible.h:747
// index:0
// Public inline
// void QAccessibleTextCursorEvent(class QObject *, int)
func NewQAccessibleTextCursorEvent(obj *qtcore.QObject /*444 QObject **/, cursorPos int) *QAccessibleTextCursorEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventC2EP7QObjecti", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &cursorPos)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextCursorEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:753
// index:1
// Public inline
// void QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
func NewQAccessibleTextCursorEvent_1(iface *QAccessibleInterface /*444 QAccessibleInterface **/, cursorPos int) *QAccessibleTextCursorEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = iface.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &cursorPos)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextCursorEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:760
// index:0
// Public virtual
// void ~QAccessibleTextCursorEvent()
func DeleteQAccessibleTextCursorEvent(*QAccessibleTextCursorEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:762
// index:0
// Public inline
// void setCursorPosition(int)
func (this *QAccessibleTextCursorEvent) SetCursorPosition(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEvent17setCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:763
// index:0
// Public inline
// int cursorPosition()
func (this *QAccessibleTextCursorEvent) CursorPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextCursorEvent14cursorPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
