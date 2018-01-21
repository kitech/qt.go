package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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
	*QAccessibleTextCursorEvent
}

func (this *QAccessibleTextUpdateEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleTextCursorEvent.GetCthis()
	}
}
func NewQAccessibleTextUpdateEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextUpdateEvent {
	bcthis0 := NewQAccessibleTextCursorEventFromPointer(cthis)
	return &QAccessibleTextUpdateEvent{bcthis0}
}

// /usr/include/qt/QtGui/qaccessible.h:864
// index:0
// Public inline
// void QAccessibleTextUpdateEvent(class QObject *, int, const class QString &, const class QString &)
func NewQAccessibleTextUpdateEvent(obj *qtcore.QObject /*444 QObject **/, position int, oldText *qtcore.QString, text *qtcore.QString) *QAccessibleTextUpdateEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	var convArg0 = obj.GetCthis()
	var convArg2 = oldText.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &position, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextUpdateEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:870
// index:1
// Public inline
// void QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
func NewQAccessibleTextUpdateEvent_1(iface *QAccessibleInterface /*444 QAccessibleInterface **/, position int, oldText *qtcore.QString, text *qtcore.QString) *QAccessibleTextUpdateEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	var convArg0 = iface.GetCthis()
	var convArg2 = oldText.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &position, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextUpdateEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:877
// index:0
// Public virtual
// void ~QAccessibleTextUpdateEvent()
func DeleteQAccessibleTextUpdateEvent(*QAccessibleTextUpdateEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:879
// index:0
// Public inline
// QString textRemoved()
func (this *QAccessibleTextUpdateEvent) TextRemoved() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent11textRemovedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:882
// index:0
// Public inline
// QString textInserted()
func (this *QAccessibleTextUpdateEvent) TextInserted() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent12textInsertedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:885
// index:0
// Public inline
// int changePosition()
func (this *QAccessibleTextUpdateEvent) ChangePosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent14changePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
