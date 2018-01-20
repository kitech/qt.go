//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

// /usr/include/qt/QtWidgets/qgesture.h:106
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPanGesture) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:117
// index:0
// void QPanGesture(class QObject *)
func NewQPanGesture(parent unsafe.Pointer) *QPanGesture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPanGestureFromPointer(cthis)
	return gothis
}
func NewQPanGestureFromPointer(cthis unsafe.Pointer) *QPanGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QPanGesture{bcthis0}
}

// /usr/include/qt/QtWidgets/qgesture.h:118
// index:0
// virtual
// void ~QPanGesture()
func DeleteQPanGesture(*QPanGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:120
// index:0
// QPointF lastOffset()
func (this *QPanGesture) LastOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture10lastOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:121
// index:0
// QPointF offset()
func (this *QPanGesture) Offset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture6offsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:122
// index:0
// QPointF delta()
func (this *QPanGesture) Delta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture5deltaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:123
// index:0
// qreal acceleration()
func (this *QPanGesture) Acceleration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPanGesture12accelerationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:125
// index:0
// void setLastOffset(const class QPointF &)
func (this *QPanGesture) SetLastOffset(value unsafe.Pointer) {
	// 0: (, value const QPointF &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGesture13setLastOffsetERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:126
// index:0
// void setOffset(const class QPointF &)
func (this *QPanGesture) SetOffset(value unsafe.Pointer) {
	// 0: (, value const QPointF &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGesture9setOffsetERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:127
// index:0
// void setAcceleration(qreal)
func (this *QPanGesture) SetAcceleration(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPanGesture15setAccelerationEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
