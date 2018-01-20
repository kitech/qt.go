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

// /usr/include/qt/QtWidgets/qgesture.h:136
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPinchGesture) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:165
// index:0
// void QPinchGesture(class QObject *)
func NewQPinchGesture(parent unsafe.Pointer) *QPinchGesture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPinchGestureFromPointer(cthis)
	return gothis
}
func NewQPinchGestureFromPointer(cthis unsafe.Pointer) *QPinchGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QPinchGesture{bcthis0}
}

// /usr/include/qt/QtWidgets/qgesture.h:166
// index:0
// virtual
// void ~QPinchGesture()
func DeleteQPinchGesture(*QPinchGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:168
// index:0
// QPinchGesture::ChangeFlags totalChangeFlags()
func (this *QPinchGesture) TotalChangeFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16totalChangeFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:169
// index:0
// void setTotalChangeFlags(QPinchGesture::ChangeFlags)
func (this *QPinchGesture) SetTotalChangeFlags(value int) {
	// 0: (, QFlags<QPinchGesture::ChangeFlag> value), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalChangeFlagsE6QFlagsINS_10ChangeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:171
// index:0
// QPinchGesture::ChangeFlags changeFlags()
func (this *QPinchGesture) ChangeFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11changeFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:172
// index:0
// void setChangeFlags(QPinchGesture::ChangeFlags)
func (this *QPinchGesture) SetChangeFlags(value int) {
	// 0: (, QFlags<QPinchGesture::ChangeFlag> value), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setChangeFlagsE6QFlagsINS_10ChangeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:174
// index:0
// QPointF startCenterPoint()
func (this *QPinchGesture) StartCenterPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16startCenterPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:175
// index:0
// QPointF lastCenterPoint()
func (this *QPinchGesture) LastCenterPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture15lastCenterPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:176
// index:0
// QPointF centerPoint()
func (this *QPinchGesture) CenterPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11centerPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:177
// index:0
// void setStartCenterPoint(const class QPointF &)
func (this *QPinchGesture) SetStartCenterPoint(value unsafe.Pointer) {
	// 0: (, value const QPointF &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setStartCenterPointERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:178
// index:0
// void setLastCenterPoint(const class QPointF &)
func (this *QPinchGesture) SetLastCenterPoint(value unsafe.Pointer) {
	// 0: (, value const QPointF &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture18setLastCenterPointERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:179
// index:0
// void setCenterPoint(const class QPointF &)
func (this *QPinchGesture) SetCenterPoint(value unsafe.Pointer) {
	// 0: (, value const QPointF &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setCenterPointERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:181
// index:0
// qreal totalScaleFactor()
func (this *QPinchGesture) TotalScaleFactor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture16totalScaleFactorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:182
// index:0
// qreal lastScaleFactor()
func (this *QPinchGesture) LastScaleFactor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture15lastScaleFactorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:183
// index:0
// qreal scaleFactor()
func (this *QPinchGesture) ScaleFactor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture11scaleFactorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:184
// index:0
// void setTotalScaleFactor(qreal)
func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalScaleFactorEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:185
// index:0
// void setLastScaleFactor(qreal)
func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture18setLastScaleFactorEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:186
// index:0
// void setScaleFactor(qreal)
func (this *QPinchGesture) SetScaleFactor(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture14setScaleFactorEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:188
// index:0
// qreal totalRotationAngle()
func (this *QPinchGesture) TotalRotationAngle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture18totalRotationAngleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:189
// index:0
// qreal lastRotationAngle()
func (this *QPinchGesture) LastRotationAngle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture17lastRotationAngleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:190
// index:0
// qreal rotationAngle()
func (this *QPinchGesture) RotationAngle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPinchGesture13rotationAngleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:191
// index:0
// void setTotalRotationAngle(qreal)
func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture21setTotalRotationAngleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:192
// index:0
// void setLastRotationAngle(qreal)
func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture20setLastRotationAngleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:193
// index:0
// void setRotationAngle(qreal)
func (this *QPinchGesture) SetRotationAngle(value float64) {
	// 0: (, value qreal), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPinchGesture16setRotationAngleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
