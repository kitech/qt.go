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
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAccessibleTextRemoveEvent struct {
	*QAccessibleTextCursorEvent
}

func (this *QAccessibleTextRemoveEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleTextCursorEvent.GetCthis()
	}
}
func (this *QAccessibleTextRemoveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleTextCursorEvent = NewQAccessibleTextCursorEventFromPointer(cthis)
}
func NewQAccessibleTextRemoveEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextRemoveEvent {
	bcthis0 := NewQAccessibleTextCursorEventFromPointer(cthis)
	return &QAccessibleTextRemoveEvent{bcthis0}
}
func (*QAccessibleTextRemoveEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTextRemoveEvent {
	return NewQAccessibleTextRemoveEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:834
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextRemoveEvent(QObject *, int, const QString &)
func NewQAccessibleTextRemoveEvent(obj *qtcore.QObject /*777 QObject **/, position int, text string) *QAccessibleTextRemoveEvent {
	var convArg0 = obj.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextRemoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextRemoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:840
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextRemoveEvent(QAccessibleInterface *, int, const QString &)
func NewQAccessibleTextRemoveEvent_1(iface *QAccessibleInterface /*777 QAccessibleInterface **/, position int, text string) *QAccessibleTextRemoveEvent {
	var convArg0 = iface.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextRemoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextRemoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:847
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextRemoveEvent()
func DeleteQAccessibleTextRemoveEvent(this *QAccessibleTextRemoveEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:849
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textRemoved()
func (this *QAccessibleTextRemoveEvent) TextRemoved() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextRemoveEvent11textRemovedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:852
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int changePosition()
func (this *QAccessibleTextRemoveEvent) ChangePosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextRemoveEvent14changePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
