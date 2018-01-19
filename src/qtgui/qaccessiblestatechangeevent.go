//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QAccessibleStateChangeEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qaccessible.h:723
// index:0
// inline
// void QAccessibleStateChangeEvent(class QObject *, class QAccessible::State)
func NewQAccessibleStateChangeEvent(obj unsafe.Pointer, state int) *QAccessibleStateChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleStateChangeEventC2EP7QObjectN11QAccessible5StateE", ffiqt.FFI_TYPE_VOID, cthis, obj, &state)
	gopp.ErrPrint(err, rv)
	return &QAccessibleStateChangeEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:728
// index:1
// inline
// void QAccessibleStateChangeEvent(class QAccessibleInterface *, class QAccessible::State)
func NewQAccessibleStateChangeEvent_1(iface unsafe.Pointer, state int) *QAccessibleStateChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleStateChangeEventC2EP20QAccessibleInterfaceN11QAccessible5StateE", ffiqt.FFI_TYPE_VOID, cthis, iface, &state)
	gopp.ErrPrint(err, rv)
	return &QAccessibleStateChangeEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:733
// index:0
// virtual
// void ~QAccessibleStateChangeEvent()
func DeleteQAccessibleStateChangeEvent(*QAccessibleStateChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleStateChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:735
// index:0
// inline
// QAccessible::State changedStates()
func (this *QAccessibleStateChangeEvent) ChangedStates() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAccessibleStateChangeEvent13changedStatesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
