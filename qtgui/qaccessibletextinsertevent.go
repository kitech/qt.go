package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessibleTextInsertEvent struct {
	*QAccessibleTextCursorEvent
}
type QAccessibleTextInsertEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextInsertEvent_PTR() *QAccessibleTextInsertEvent
}

func (ptr *QAccessibleTextInsertEvent) QAccessibleTextInsertEvent_PTR() *QAccessibleTextInsertEvent {
	return ptr
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
func NewQAccessibleTextInsertEvent(obj qtcore.QObject_ITF /*777 QObject **/, position int, text string) *QAccessibleTextInsertEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextInsertEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextInsertEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:810
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextInsertEvent(QAccessibleInterface *, int, const QString &)
func NewQAccessibleTextInsertEvent_1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, position int, text string) *QAccessibleTextInsertEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextInsertEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextInsertEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:817
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextInsertEvent()
func DeleteQAccessibleTextInsertEvent(this *QAccessibleTextInsertEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextInsertEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:819
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textInserted()
func (this *QAccessibleTextInsertEvent) TextInserted() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextInsertEvent12textInsertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:822
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int changePosition()
func (this *QAccessibleTextInsertEvent) ChangePosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextInsertEvent14changePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
