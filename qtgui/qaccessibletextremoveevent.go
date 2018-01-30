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
func NewQAccessibleTextRemoveEvent(obj *qtcore.QObject /*777 QObject **/, position int, text *qtcore.QString) *QAccessibleTextRemoveEvent {
	var convArg0 = obj.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextRemoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:840
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextRemoveEvent(QAccessibleInterface *, int, const QString &)
func NewQAccessibleTextRemoveEvent_1(iface *QAccessibleInterface /*777 QAccessibleInterface **/, position int, text *qtcore.QString) *QAccessibleTextRemoveEvent {
	var convArg0 = iface.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString", ffiqt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextRemoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:847
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextRemoveEvent()
func DeleteQAccessibleTextRemoveEvent(*QAccessibleTextRemoveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:849
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textRemoved()
func (this *QAccessibleTextRemoveEvent) TextRemoved() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextRemoveEvent11textRemovedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:852
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int changePosition()
func (this *QAccessibleTextRemoveEvent) ChangePosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleTextRemoveEvent14changePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
