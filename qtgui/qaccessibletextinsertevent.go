package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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

type QAccessibleTextInsertEvent struct {
	*QAccessibleTextCursorEvent
}

func (this *QAccessibleTextInsertEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleTextCursorEvent.GetCthis()
	}
}
func (this *QAccessibleTextInsertEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleTextCursorEvent = NewQAccessibleTextCursorEventFromPointer(cthis)
}
func NewQAccessibleTextInsertEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextInsertEvent {
	bcthis0 := NewQAccessibleTextCursorEventFromPointer(cthis)
	return &QAccessibleTextInsertEvent{bcthis0}
}
func (*QAccessibleTextInsertEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTextInsertEvent {
	return NewQAccessibleTextInsertEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:804
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextInsertEvent(QObject *, int, const QString &)
func NewQAccessibleTextInsertEvent(obj *qtcore.QObject /*777 QObject **/, position int, text *qtcore.QString) *QAccessibleTextInsertEvent {
	var convArg0 = obj.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextInsertEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextInsertEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:810
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextInsertEvent(QAccessibleInterface *, int, const QString &)
func NewQAccessibleTextInsertEvent_1(iface *QAccessibleInterface /*777 QAccessibleInterface **/, position int, text *qtcore.QString) *QAccessibleTextInsertEvent {
	var convArg0 = iface.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextInsertEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextInsertEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:817
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextInsertEvent()
func DeleteQAccessibleTextInsertEvent(this *QAccessibleTextInsertEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:819
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textInserted()
func (this *QAccessibleTextInsertEvent) TextInserted() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextInsertEvent12textInsertedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:822
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int changePosition()
func (this *QAccessibleTextInsertEvent) ChangePosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextInsertEvent14changePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
