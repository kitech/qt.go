//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QAccessibleTextInsertEvent struct {
	*QAccessibleTextCursorEvent
}

func (this *QAccessibleTextInsertEvent) GetCthis() unsafe.Pointer {
	return this.QAccessibleTextCursorEvent.GetCthis()
}

// /usr/include/qt/QtGui/qaccessible.h:804
// index:0
// inline
// void QAccessibleTextInsertEvent(class QObject *, int, const class QString &)
func NewQAccessibleTextInsertEvent(obj unsafe.Pointer, position int, text unsafe.Pointer) *QAccessibleTextInsertEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString", ffiqt.FFI_TYPE_VOID, cthis, obj, &position, text)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextInsertEventFromPointer(cthis)
	return gothis
}
func NewQAccessibleTextInsertEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextInsertEvent {
	bcthis0 := NewQAccessibleTextCursorEventFromPointer(cthis)
	return &QAccessibleTextInsertEvent{bcthis0}
}

// /usr/include/qt/QtGui/qaccessible.h:810
// index:1
// inline
// void QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
func NewQAccessibleTextInsertEvent_1(iface unsafe.Pointer, position int, text unsafe.Pointer) *QAccessibleTextInsertEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString", ffiqt.FFI_TYPE_VOID, cthis, iface, &position, text)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextInsertEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:817
// index:0
// virtual
// void ~QAccessibleTextInsertEvent()
func DeleteQAccessibleTextInsertEvent(*QAccessibleTextInsertEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:819
// index:0
// inline
// QString textInserted()
func (this *QAccessibleTextInsertEvent) TextInserted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextInsertEvent12textInsertedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:822
// index:0
// inline
// int changePosition()
func (this *QAccessibleTextInsertEvent) ChangePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextInsertEvent14changePositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
