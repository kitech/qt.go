//  header block begin

// +build !minimal

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
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtGui/qevent.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabletEvent(QEvent::Type, const QPointF &, const QPointF &, int, int, qreal, int, int, qreal, qreal, int, Qt::KeyboardModifiers, qint64)

/*

 */
func (*QTabletEvent) NewForInherit(t int, pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, device int, pointerType int, pressure float64, xTilt int, yTilt int, tangentialPressure float64, rotation float64, z int, keyState int, uniqueID int64) *QTabletEvent {
	return NewQTabletEvent(t, pos, globalPos, device, pointerType, pressure, xTilt, yTilt, tangentialPressure, rotation, z, keyState, uniqueID)
}
func NewQTabletEvent(t int, pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, device int, pointerType int, pressure float64, xTilt int, yTilt int, tangentialPressure float64, rotation float64, z int, keyState int, uniqueID int64) *QTabletEvent {
	var convArg1 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg1 = pos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg2 = globalPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTabletEventC2EN6QEvent4TypeERK7QPointFS4_iidiiddi6QFlagsIN2Qt16KeyboardModifierEEx", qtrt.FFI_TYPE_POINTER, t, convArg1, convArg2, device, pointerType, pressure, xTilt, yTilt, tangentialPressure, rotation, z, keyState, uniqueID)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabletEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTabletEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:259
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTabletEvent(QEvent::Type, const QPointF &, const QPointF &, int, int, qreal, int, int, qreal, qreal, int, Qt::KeyboardModifiers, qint64, Qt::MouseButton, Qt::MouseButtons)

/*

 */
func (*QTabletEvent) NewForInherit1(t int, pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, device int, pointerType int, pressure float64, xTilt int, yTilt int, tangentialPressure float64, rotation float64, z int, keyState int, uniqueID int64, button int, buttons int) *QTabletEvent {
	return NewQTabletEvent1(t, pos, globalPos, device, pointerType, pressure, xTilt, yTilt, tangentialPressure, rotation, z, keyState, uniqueID, button, buttons)
}
func NewQTabletEvent1(t int, pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, device int, pointerType int, pressure float64, xTilt int, yTilt int, tangentialPressure float64, rotation float64, z int, keyState int, uniqueID int64, button int, buttons int) *QTabletEvent {
	var convArg1 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg1 = pos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg2 = globalPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTabletEventC2EN6QEvent4TypeERK7QPointFS4_iidiiddi6QFlagsIN2Qt16KeyboardModifierEExNS6_11MouseButtonES5_IS9_E", qtrt.FFI_TYPE_POINTER, t, convArg1, convArg2, device, pointerType, pressure, xTilt, yTilt, tangentialPressure, rotation, z, keyState, uniqueID, button, buttons)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabletEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTabletEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:264
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTabletEvent()

/*

 */
func DeleteQTabletEvent(this *QTabletEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTabletEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 128)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:266
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos() const

/*

 */
func (this *QTabletEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:267
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint globalPos() const

/*

 */
func (this *QTabletEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:272
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF() const

/*

 */
func (this *QTabletEvent) PosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent4posFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & globalPosF() const

/*

 */
func (this *QTabletEvent) GlobalPosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent10globalPosFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:275
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x() const

/*

 */
func (this *QTabletEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:276
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*

 */
func (this *QTabletEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:277
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX() const

/*

 */
func (this *QTabletEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:278
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY() const

/*

 */
func (this *QTabletEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:279
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal hiResGlobalX() const

/*

 */
func (this *QTabletEvent) HiResGlobalX() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent12hiResGlobalXEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:280
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal hiResGlobalY() const

/*

 */
func (this *QTabletEvent) HiResGlobalY() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent12hiResGlobalYEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTabletEvent::TabletDevice device() const

/*

 */
func (this *QTabletEvent) Device() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:282
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTabletEvent::PointerType pointerType() const

/*

 */
func (this *QTabletEvent) PointerType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent11pointerTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:283
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 uniqueId() const

/*

 */
func (this *QTabletEvent) UniqueId() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent8uniqueIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:284
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal pressure() const

/*

 */
func (this *QTabletEvent) Pressure() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent8pressureEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:285
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int z() const

/*

 */
func (this *QTabletEvent) Z() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent1zEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:286
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal tangentialPressure() const

/*

 */
func (this *QTabletEvent) TangentialPressure() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent18tangentialPressureEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:287
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rotation() const

/*

 */
func (this *QTabletEvent) Rotation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent8rotationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:288
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int xTilt() const

/*

 */
func (this *QTabletEvent) XTilt() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent5xTiltEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:289
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int yTilt() const

/*

 */
func (this *QTabletEvent) YTilt() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent5yTiltEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:290
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButton button() const

/*

 */
func (this *QTabletEvent) Button() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent6buttonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:291
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons() const

/*

 */
func (this *QTabletEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTabletEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_10654() {
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
}

//  keep block end
