//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QPinchGesture struct {
	*QGesture
}

func (this *QPinchGesture) GetCthis() unsafe.Pointer {
	return this.QGesture.GetCthis()
}
func NewQPinchGestureFromPointer(cthis unsafe.Pointer) *QPinchGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QPinchGesture{bcthis0}
}

// /usr/include/qt/QtWidgets/qgesture.h:136
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPinchGesture) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:165
// index:0
// Public
// void QPinchGesture(class QObject *)
func NewQPinchGesture(parent unsafe.Pointer) *QPinchGesture {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPinchGestureFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:166
// index:0
// Public virtual
// void ~QPinchGesture()
func DeleteQPinchGesture(*QPinchGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:168
// index:0
// Public
// QPinchGesture::ChangeFlags totalChangeFlags()
func (this *QPinchGesture) TotalChangeFlags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16totalChangeFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:171
// index:0
// Public
// QPinchGesture::ChangeFlags changeFlags()
func (this *QPinchGesture) ChangeFlags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11changeFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:174
// index:0
// Public
// QPointF startCenterPoint()
func (this *QPinchGesture) StartCenterPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16startCenterPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:175
// index:0
// Public
// QPointF lastCenterPoint()
func (this *QPinchGesture) LastCenterPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture15lastCenterPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:176
// index:0
// Public
// QPointF centerPoint()
func (this *QPinchGesture) CenterPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11centerPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:177
// index:0
// Public
// void setStartCenterPoint(const class QPointF &)
func (this *QPinchGesture) SetStartCenterPoint(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setStartCenterPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:178
// index:0
// Public
// void setLastCenterPoint(const class QPointF &)
func (this *QPinchGesture) SetLastCenterPoint(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture18setLastCenterPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:179
// index:0
// Public
// void setCenterPoint(const class QPointF &)
func (this *QPinchGesture) SetCenterPoint(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setCenterPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:181
// index:0
// Public
// qreal totalScaleFactor()
func (this *QPinchGesture) TotalScaleFactor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16totalScaleFactorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:182
// index:0
// Public
// qreal lastScaleFactor()
func (this *QPinchGesture) LastScaleFactor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture15lastScaleFactorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:183
// index:0
// Public
// qreal scaleFactor()
func (this *QPinchGesture) ScaleFactor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11scaleFactorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:184
// index:0
// Public
// void setTotalScaleFactor(qreal)
func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalScaleFactorEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:185
// index:0
// Public
// void setLastScaleFactor(qreal)
func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture18setLastScaleFactorEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:186
// index:0
// Public
// void setScaleFactor(qreal)
func (this *QPinchGesture) SetScaleFactor(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setScaleFactorEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:188
// index:0
// Public
// qreal totalRotationAngle()
func (this *QPinchGesture) TotalRotationAngle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture18totalRotationAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:189
// index:0
// Public
// qreal lastRotationAngle()
func (this *QPinchGesture) LastRotationAngle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture17lastRotationAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:190
// index:0
// Public
// qreal rotationAngle()
func (this *QPinchGesture) RotationAngle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture13rotationAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:191
// index:0
// Public
// void setTotalRotationAngle(qreal)
func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture21setTotalRotationAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:192
// index:0
// Public
// void setLastRotationAngle(qreal)
func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture20setLastRotationAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:193
// index:0
// Public
// void setRotationAngle(qreal)
func (this *QPinchGesture) SetRotationAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture16setRotationAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
