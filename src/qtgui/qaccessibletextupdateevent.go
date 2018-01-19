//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QAccessibleTextUpdateEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qaccessible.h:864
// index:0
// inline
// void QAccessibleTextUpdateEvent(class QObject *, int, const class QString &, const class QString &)
func NewQAccessibleTextUpdateEvent(obj unsafe.Pointer, position int, oldText unsafe.Pointer, text unsafe.Pointer) *QAccessibleTextUpdateEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_", ffiqt.FFI_TYPE_VOID, cthis, obj, &position, oldText, text)
	gopp.ErrPrint(err, rv)
	return &QAccessibleTextUpdateEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:870
// index:1
// inline
// void QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
func NewQAccessibleTextUpdateEvent_1(iface unsafe.Pointer, position int, oldText unsafe.Pointer, text unsafe.Pointer) *QAccessibleTextUpdateEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_", ffiqt.FFI_TYPE_VOID, cthis, iface, &position, oldText, text)
	gopp.ErrPrint(err, rv)
	return &QAccessibleTextUpdateEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:877
// index:0
// virtual
// void ~QAccessibleTextUpdateEvent()
func DeleteQAccessibleTextUpdateEvent(*QAccessibleTextUpdateEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:879
// index:0
// inline
// QString textRemoved()
func (this *QAccessibleTextUpdateEvent) TextRemoved() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent11textRemovedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:882
// index:0
// inline
// QString textInserted()
func (this *QAccessibleTextUpdateEvent) TextInserted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent12textInsertedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:885
// index:0
// inline
// int changePosition()
func (this *QAccessibleTextUpdateEvent) ChangePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent14changePositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
