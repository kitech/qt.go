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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessibleValueChangeEvent struct {
	*QAccessibleEvent
}
type QAccessibleValueChangeEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleValueChangeEvent_PTR() *QAccessibleValueChangeEvent
}

func (ptr *QAccessibleValueChangeEvent) QAccessibleValueChangeEvent_PTR() *QAccessibleValueChangeEvent {
	return ptr
}

func (this *QAccessibleValueChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleEvent.GetCthis()
	}
}
func (this *QAccessibleValueChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleEvent = NewQAccessibleEventFromPointer(cthis)
}
func NewQAccessibleValueChangeEventFromPointer(cthis unsafe.Pointer) *QAccessibleValueChangeEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleValueChangeEvent{bcthis0}
}
func (*QAccessibleValueChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleValueChangeEvent {
	return NewQAccessibleValueChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:898
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleValueChangeEvent(QObject *, const QVariant &)
func NewQAccessibleValueChangeEvent(obj qtcore.QObject_ITF /*777 QObject **/, val qtcore.QVariant_ITF) *QAccessibleValueChangeEvent {
	var convArg0 = obj.QObject_PTR().GetCthis()
	var convArg1 = val.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEventC2EP7QObjectRK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleValueChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleValueChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:904
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleValueChangeEvent(QAccessibleInterface *, const QVariant &)
func NewQAccessibleValueChangeEvent_1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, val qtcore.QVariant_ITF) *QAccessibleValueChangeEvent {
	var convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	var convArg1 = val.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEventC2EP20QAccessibleInterfaceRK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleValueChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleValueChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:911
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleValueChangeEvent()
func DeleteQAccessibleValueChangeEvent(this *QAccessibleValueChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:913
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setValue(const QVariant &)
func (this *QAccessibleValueChangeEvent) SetValue(val qtcore.QVariant_ITF) {
	var convArg0 = val.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:914
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVariant value()
func (this *QAccessibleValueChangeEvent) Value() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAccessibleValueChangeEvent5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
