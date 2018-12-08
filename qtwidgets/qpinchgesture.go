package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QPinchGesture struct {
	*QGesture
}
type QPinchGesture_ITF interface {
	QGesture_ITF
	QPinchGesture_PTR() *QPinchGesture
}

func (ptr *QPinchGesture) QPinchGesture_PTR() *QPinchGesture { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPinchGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesture.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPinchGesture(QObject *)

/*

 */
func (*QPinchGesture) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QPinchGesture {
	return NewQPinchGesture(parent)
}
func NewQPinchGesture(parent qtcore.QObject_ITF /*777 QObject **/) *QPinchGesture {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPinchGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPinchGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPinchGesture(QObject *)

/*

 */
func (*QPinchGesture) NewForInheritp() *QPinchGesture {
	return NewQPinchGesturep()
}
func NewQPinchGesturep() *QPinchGesture {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPinchGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPinchGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPinchGesture()

/*

 */
func DeleteQPinchGesture(this *QPinchGesture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGestureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:168
// index:0
// Public Visibility=Default Availability=Available
// [4] QPinchGesture::ChangeFlags totalChangeFlags() const

/*

 */
func (this *QPinchGesture) TotalChangeFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture16totalChangeFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTotalChangeFlags(QPinchGesture::ChangeFlags)

/*

 */
func (this *QPinchGesture) SetTotalChangeFlags(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalChangeFlagsE6QFlagsINS_10ChangeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:171
// index:0
// Public Visibility=Default Availability=Available
// [4] QPinchGesture::ChangeFlags changeFlags() const

/*

 */
func (this *QPinchGesture) ChangeFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture11changeFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChangeFlags(QPinchGesture::ChangeFlags)

/*

 */
func (this *QPinchGesture) SetChangeFlags(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture14setChangeFlagsE6QFlagsINS_10ChangeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF startCenterPoint() const

/*

 */
func (this *QPinchGesture) StartCenterPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture16startCenterPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:175
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastCenterPoint() const

/*

 */
func (this *QPinchGesture) LastCenterPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture15lastCenterPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:176
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF centerPoint() const

/*

 */
func (this *QPinchGesture) CenterPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture11centerPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartCenterPoint(const QPointF &)

/*

 */
func (this *QPinchGesture) SetStartCenterPoint(value qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QPointF_PTR() != nil {
		convArg0 = value.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture19setStartCenterPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastCenterPoint(const QPointF &)

/*

 */
func (this *QPinchGesture) SetLastCenterPoint(value qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QPointF_PTR() != nil {
		convArg0 = value.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture18setLastCenterPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenterPoint(const QPointF &)

/*

 */
func (this *QPinchGesture) SetCenterPoint(value qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QPointF_PTR() != nil {
		convArg0 = value.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture14setCenterPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal totalScaleFactor() const

/*

 */
func (this *QPinchGesture) TotalScaleFactor() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture16totalScaleFactorEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lastScaleFactor() const

/*

 */
func (this *QPinchGesture) LastScaleFactor() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture15lastScaleFactorEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal scaleFactor() const

/*

 */
func (this *QPinchGesture) ScaleFactor() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture11scaleFactorEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTotalScaleFactor(qreal)

/*

 */
func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture19setTotalScaleFactorEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScaleFactor(qreal)

/*

 */
func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture18setLastScaleFactorEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaleFactor(qreal)

/*

 */
func (this *QPinchGesture) SetScaleFactor(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture14setScaleFactorEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal totalRotationAngle() const

/*

 */
func (this *QPinchGesture) TotalRotationAngle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture18totalRotationAngleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lastRotationAngle() const

/*

 */
func (this *QPinchGesture) LastRotationAngle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture17lastRotationAngleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:190
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rotationAngle() const

/*

 */
func (this *QPinchGesture) RotationAngle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPinchGesture13rotationAngleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTotalRotationAngle(qreal)

/*

 */
func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture21setTotalRotationAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastRotationAngle(qreal)

/*

 */
func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture20setLastRotationAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRotationAngle(qreal)

/*

 */
func (this *QPinchGesture) SetRotationAngle(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPinchGesture16setRotationAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QPinchGesture__ChangeFlag = int

//
const QPinchGesture__ScaleFactorChanged QPinchGesture__ChangeFlag = 1

//
const QPinchGesture__RotationAngleChanged QPinchGesture__ChangeFlag = 2

//
const QPinchGesture__CenterPointChanged QPinchGesture__ChangeFlag = 4

func (this *QPinchGesture) ChangeFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QPinchGesture_ChangeFlagItemName(val int) string {
	var nilthis *QPinchGesture
	return nilthis.ChangeFlagItemName(val)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
