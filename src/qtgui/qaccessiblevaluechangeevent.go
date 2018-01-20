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
type QAccessibleValueChangeEvent struct {
	*QAccessibleEvent
}

func (this *QAccessibleValueChangeEvent) GetCthis() unsafe.Pointer {
	return this.QAccessibleEvent.GetCthis()
}

// /usr/include/qt/QtGui/qaccessible.h:898
// index:0
// inline
// void QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
func NewQAccessibleValueChangeEvent(obj unsafe.Pointer, val unsafe.Pointer) *QAccessibleValueChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEventC2EP7QObjectRK8QVariant", ffiqt.FFI_TYPE_VOID, cthis, obj, val)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleValueChangeEventFromPointer(cthis)
	return gothis
}
func NewQAccessibleValueChangeEventFromPointer(cthis unsafe.Pointer) *QAccessibleValueChangeEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleValueChangeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qaccessible.h:904
// index:1
// inline
// void QAccessibleValueChangeEvent(class QAccessibleInterface *, const class QVariant &)
func NewQAccessibleValueChangeEvent_1(iface unsafe.Pointer, val unsafe.Pointer) *QAccessibleValueChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEventC2EP20QAccessibleInterfaceRK8QVariant", ffiqt.FFI_TYPE_VOID, cthis, iface, val)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleValueChangeEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:911
// index:0
// virtual
// void ~QAccessibleValueChangeEvent()
func DeleteQAccessibleValueChangeEvent(*QAccessibleValueChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:913
// index:0
// inline
// void setValue(const class QVariant &)
func (this *QAccessibleValueChangeEvent) SetValue(val unsafe.Pointer) {
	// 0: (, val const QVariant &), (val)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:914
// index:0
// inline
// QVariant value()
func (this *QAccessibleValueChangeEvent) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAccessibleValueChangeEvent5valueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
