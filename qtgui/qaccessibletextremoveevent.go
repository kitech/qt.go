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
// extern C begin: 5
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

type QAccessibleTextRemoveEvent struct {
	*QAccessibleTextCursorEvent
}
type QAccessibleTextRemoveEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextRemoveEvent_PTR() *QAccessibleTextRemoveEvent
}

func (ptr *QAccessibleTextRemoveEvent) QAccessibleTextRemoveEvent_PTR() *QAccessibleTextRemoveEvent {
	return ptr
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
func NewQAccessibleTextRemoveEvent(obj qtcore.QObject_ITF /*777 QObject **/, position int, text string) *QAccessibleTextRemoveEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextRemoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextRemoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:840
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextRemoveEvent(QAccessibleInterface *, int, const QString &)
func NewQAccessibleTextRemoveEvent_1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, position int, text string) *QAccessibleTextRemoveEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:849
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textRemoved() const
func (this *QAccessibleTextRemoveEvent) TextRemoved() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextRemoveEvent11textRemovedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:852
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int changePosition() const
func (this *QAccessibleTextRemoveEvent) ChangePosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextRemoveEvent14changePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
