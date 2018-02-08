package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QPanGesture struct {
	*QGesture
}

func (this *QPanGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGesture.GetCthis()
	}
}
func (this *QPanGesture) SetCthis(cthis unsafe.Pointer) {
	this.QGesture = NewQGestureFromPointer(cthis)
}
func NewQPanGestureFromPointer(cthis unsafe.Pointer) *QPanGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QPanGesture{bcthis0}
}
func (*QPanGesture) NewFromPointer(cthis unsafe.Pointer) *QPanGesture {
	return NewQPanGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QPanGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPanGesture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesture.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPanGesture(QObject *)
func NewQPanGesture(parent *qtcore.QObject /*777 QObject **/) *QPanGesture {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPanGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPanGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPanGesture()
func DeleteQPanGesture(this *QPanGesture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPanGestureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:120
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastOffset()
func (this *QPanGesture) LastOffset() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPanGesture10lastOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:121
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF offset()
func (this *QPanGesture) Offset() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPanGesture6offsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:122
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF delta()
func (this *QPanGesture) Delta() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPanGesture5deltaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal acceleration()
func (this *QPanGesture) Acceleration() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPanGesture12accelerationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastOffset(const QPointF &)
func (this *QPanGesture) SetLastOffset(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPanGesture13setLastOffsetERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffset(const QPointF &)
func (this *QPanGesture) SetOffset(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPanGesture9setOffsetERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceleration(qreal)
func (this *QPanGesture) SetAcceleration(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPanGesture15setAccelerationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

//  body block end
