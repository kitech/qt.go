//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	return this.QGesture.GetCthis()
}
func NewQPanGestureFromPointer(cthis unsafe.Pointer) *QPanGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QPanGesture{bcthis0}
}

// /usr/include/qt/QtWidgets/qgesture.h:106
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPanGesture) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:117
// index:0
// Public
// void QPanGesture(class QObject *)
func NewQPanGesture(parent unsafe.Pointer) *QPanGesture {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPanGestureFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:118
// index:0
// Public virtual
// void ~QPanGesture()
func DeleteQPanGesture(*QPanGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:120
// index:0
// Public
// QPointF lastOffset()
func (this *QPanGesture) LastOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture10lastOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:121
// index:0
// Public
// QPointF offset()
func (this *QPanGesture) Offset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture6offsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:122
// index:0
// Public
// QPointF delta()
func (this *QPanGesture) Delta() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture5deltaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:123
// index:0
// Public
// qreal acceleration()
func (this *QPanGesture) Acceleration() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture12accelerationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:125
// index:0
// Public
// void setLastOffset(const class QPointF &)
func (this *QPanGesture) SetLastOffset(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGesture13setLastOffsetERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:126
// index:0
// Public
// void setOffset(const class QPointF &)
func (this *QPanGesture) SetOffset(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGesture9setOffsetERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:127
// index:0
// Public
// void setAcceleration(qreal)
func (this *QPanGesture) SetAcceleration(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGesture15setAccelerationEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
