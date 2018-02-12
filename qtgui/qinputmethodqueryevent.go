package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QInputMethodQueryEvent struct {
	*qtcore.QEvent
}

func (this *QInputMethodQueryEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QInputMethodQueryEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQInputMethodQueryEventFromPointer(cthis unsafe.Pointer) *QInputMethodQueryEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputMethodQueryEvent{bcthis0}
}
func (*QInputMethodQueryEvent) NewFromPointer(cthis unsafe.Pointer) *QInputMethodQueryEvent {
	return NewQInputMethodQueryEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:581
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputMethodQueryEvent(Qt::InputMethodQueries)
func NewQInputMethodQueryEvent(queries int) *QInputMethodQueryEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QInputMethodQueryEventC2E6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, queries)
	qtrt.ErrPrint(err, rv)
	gothis := NewQInputMethodQueryEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQInputMethodQueryEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:582
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QInputMethodQueryEvent()
func DeleteQInputMethodQueryEvent(this *QInputMethodQueryEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QInputMethodQueryEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:584
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::InputMethodQueries queries()
func (this *QInputMethodQueryEvent) Queries() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QInputMethodQueryEvent7queriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:586
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(Qt::InputMethodQuery, const QVariant &)
func (this *QInputMethodQueryEvent) SetValue(query int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QInputMethodQueryEvent8setValueEN2Qt16InputMethodQueryERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:587
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(Qt::InputMethodQuery)
func (this *QInputMethodQueryEvent) Value(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QInputMethodQueryEvent5valueEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
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
