package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

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
import "mkuse/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.QGesture.GetCthis()
	}
}
func (this *QPinchGesture) SetCthis(cthis unsafe.Pointer) {
	this.QGesture = NewQGestureFromPointer(cthis)
}
func NewQPinchGestureFromPointer(cthis unsafe.Pointer) *QPinchGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QPinchGesture{bcthis0}
}
func (*QPinchGesture) NewFromPointer(cthis unsafe.Pointer) *QPinchGesture {
	return NewQPinchGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:136
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPinchGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:165
// index:0
// Public
// void QPinchGesture(class QObject *)
func NewQPinchGesture(parent *qtcore.QObject /*777 QObject **/) *QPinchGesture {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QPinchGesture) TotalChangeFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16totalChangeFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:169
// index:0
// Public
// void setTotalChangeFlags(QPinchGesture::ChangeFlags)
func (this *QPinchGesture) SetTotalChangeFlags(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalChangeFlagsE6QFlagsINS_10ChangeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:171
// index:0
// Public
// QPinchGesture::ChangeFlags changeFlags()
func (this *QPinchGesture) ChangeFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11changeFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:172
// index:0
// Public
// void setChangeFlags(QPinchGesture::ChangeFlags)
func (this *QPinchGesture) SetChangeFlags(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setChangeFlagsE6QFlagsINS_10ChangeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:174
// index:0
// Public
// QPointF startCenterPoint()
func (this *QPinchGesture) StartCenterPoint() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16startCenterPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:175
// index:0
// Public
// QPointF lastCenterPoint()
func (this *QPinchGesture) LastCenterPoint() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture15lastCenterPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:176
// index:0
// Public
// QPointF centerPoint()
func (this *QPinchGesture) CenterPoint() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11centerPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QPinchGesture) TotalScaleFactor() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16totalScaleFactorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:182
// index:0
// Public
// qreal lastScaleFactor()
func (this *QPinchGesture) LastScaleFactor() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture15lastScaleFactorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:183
// index:0
// Public
// qreal scaleFactor()
func (this *QPinchGesture) ScaleFactor() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11scaleFactorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:184
// index:0
// Public
// void setTotalScaleFactor(qreal)
func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalScaleFactorEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:185
// index:0
// Public
// void setLastScaleFactor(qreal)
func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture18setLastScaleFactorEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:186
// index:0
// Public
// void setScaleFactor(qreal)
func (this *QPinchGesture) SetScaleFactor(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setScaleFactorEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:188
// index:0
// Public
// qreal totalRotationAngle()
func (this *QPinchGesture) TotalRotationAngle() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture18totalRotationAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:189
// index:0
// Public
// qreal lastRotationAngle()
func (this *QPinchGesture) LastRotationAngle() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture17lastRotationAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:190
// index:0
// Public
// qreal rotationAngle()
func (this *QPinchGesture) RotationAngle() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture13rotationAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:191
// index:0
// Public
// void setTotalRotationAngle(qreal)
func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture21setTotalRotationAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:192
// index:0
// Public
// void setLastRotationAngle(qreal)
func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture20setLastRotationAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:193
// index:0
// Public
// void setRotationAngle(qreal)
func (this *QPinchGesture) SetRotationAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture16setRotationAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

type QPinchGesture__ChangeFlag = int

const QPinchGesture__ScaleFactorChanged QPinchGesture__ChangeFlag = 1
const QPinchGesture__RotationAngleChanged QPinchGesture__ChangeFlag = 2
const QPinchGesture__CenterPointChanged QPinchGesture__ChangeFlag = 4

//  body block end
