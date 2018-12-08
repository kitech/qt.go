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
// extern C begin: 23
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

/*

 */
type QTabletEvent struct {
	*QInputEvent
}
type QTabletEvent_ITF interface {
	QInputEvent_ITF
	QTabletEvent_PTR() *QTabletEvent
}

func (ptr *QTabletEvent) QTabletEvent_PTR() *QTabletEvent { return ptr }

func (this *QTabletEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QTabletEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQTabletEventFromPointer(cthis unsafe.Pointer) *QTabletEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QTabletEvent{bcthis0}
}
func (*QTabletEvent) NewFromPointer(cthis unsafe.Pointer) *QTabletEvent {
	return NewQTabletEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:250
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

// /usr/include/qt/QtGui/qevent.h:254
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

// /usr/include/qt/QtGui/qevent.h:259
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

// /usr/include/qt/QtGui/qevent.h:261
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

// /usr/include/qt/QtGui/qevent.h:262
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

// /usr/include/qt/QtGui/qevent.h:267
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

// /usr/include/qt/QtGui/qevent.h:268
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

// /usr/include/qt/QtGui/qevent.h:270
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

// /usr/include/qt/QtGui/qevent.h:271
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

// /usr/include/qt/QtGui/qevent.h:272
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

// /usr/include/qt/QtGui/qevent.h:273
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

// /usr/include/qt/QtGui/qevent.h:274
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

// /usr/include/qt/QtGui/qevent.h:275
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

// /usr/include/qt/QtGui/qevent.h:276
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

// /usr/include/qt/QtGui/qevent.h:277
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

// /usr/include/qt/QtGui/qevent.h:278
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

// /usr/include/qt/QtGui/qevent.h:279
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

// /usr/include/qt/QtGui/qevent.h:280
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

// /usr/include/qt/QtGui/qevent.h:281
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

// /usr/include/qt/QtGui/qevent.h:282
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

// /usr/include/qt/QtGui/qevent.h:283
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

// /usr/include/qt/QtGui/qevent.h:284
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

// /usr/include/qt/QtGui/qevent.h:285
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

// /usr/include/qt/QtGui/qevent.h:286
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

/*


 */
type QTabletEvent__TabletDevice = int

//
const QTabletEvent__NoDevice QTabletEvent__TabletDevice = 0

//
const QTabletEvent__Puck QTabletEvent__TabletDevice = 1

//
const QTabletEvent__Stylus QTabletEvent__TabletDevice = 2

//
const QTabletEvent__Airbrush QTabletEvent__TabletDevice = 3

//
const QTabletEvent__FourDMouse QTabletEvent__TabletDevice = 4

//
const QTabletEvent__XFreeEraser QTabletEvent__TabletDevice = 5

//
const QTabletEvent__RotationStylus QTabletEvent__TabletDevice = 6

func (this *QTabletEvent) TabletDeviceItemName(val int) string {
	switch val {
	case QTabletEvent__NoDevice: // 0
		return "NoDevice"
	case QTabletEvent__Puck: // 1
		return "Puck"
	case QTabletEvent__Stylus: // 2
		return "Stylus"
	case QTabletEvent__Airbrush: // 3
		return "Airbrush"
	case QTabletEvent__FourDMouse: // 4
		return "FourDMouse"
	case QTabletEvent__XFreeEraser: // 5
		return "XFreeEraser"
	case QTabletEvent__RotationStylus: // 6
		return "RotationStylus"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTabletEvent_TabletDeviceItemName(val int) string {
	var nilthis *QTabletEvent
	return nilthis.TabletDeviceItemName(val)
}

/*


 */
type QTabletEvent__PointerType = int

//
const QTabletEvent__UnknownPointer QTabletEvent__PointerType = 0

//
const QTabletEvent__Pen QTabletEvent__PointerType = 1

//
const QTabletEvent__Cursor QTabletEvent__PointerType = 2

//
const QTabletEvent__Eraser QTabletEvent__PointerType = 3

func (this *QTabletEvent) PointerTypeItemName(val int) string {
	switch val {
	case QTabletEvent__UnknownPointer: // 0
		return "UnknownPointer"
	case QTabletEvent__Pen: // 1
		return "Pen"
	case QTabletEvent__Cursor: // 2
		return "Cursor"
	case QTabletEvent__Eraser: // 3
		return "Eraser"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTabletEvent_PointerTypeItemName(val int) string {
	var nilthis *QTabletEvent
	return nilthis.PointerTypeItemName(val)
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
}

//  keep block end
