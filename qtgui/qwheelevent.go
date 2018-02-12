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
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QWheelEvent struct {
	*QInputEvent
}

func (this *QWheelEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QWheelEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQWheelEventFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QWheelEvent{bcthis0}
}
func (*QWheelEvent) NewFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	return NewQWheelEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, int, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::Orientation)
func NewQWheelEvent(pos *qtcore.QPointF, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFi6QFlagsIN2Qt11MouseButtonEES3_INS4_16KeyboardModifierEENS4_11OrientationE", qtrt.FFI_TYPE_POINTER, convArg0, delta, buttons, modifiers, orient)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:181
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, int, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::Orientation)
func NewQWheelEvent_1(pos *qtcore.QPointF, globalPos *qtcore.QPointF, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	var convArg0 = pos.GetCthis()
	var convArg1 = globalPos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_i6QFlagsIN2Qt11MouseButtonEES3_INS4_16KeyboardModifierEENS4_11OrientationE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, delta, buttons, modifiers, orient)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:184
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, QPoint, QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers)
func NewQWheelEvent_2(pos *qtcore.QPointF, globalPos *qtcore.QPointF, pixelDelta *qtcore.QPoint /*123*/, angleDelta *qtcore.QPoint /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int) *QWheelEvent {
	var convArg0 = pos.GetCthis()
	var convArg1 = globalPos.GetCthis()
	var convArg2 = pixelDelta.GetCthis()
	var convArg3 = angleDelta.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, qt4Delta, qt4Orientation, buttons, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:187
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, QPoint, QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase)
func NewQWheelEvent_3(pos *qtcore.QPointF, globalPos *qtcore.QPointF, pixelDelta *qtcore.QPoint /*123*/, angleDelta *qtcore.QPoint /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int) *QWheelEvent {
	var convArg0 = pos.GetCthis()
	var convArg1 = globalPos.GetCthis()
	var convArg2 = pixelDelta.GetCthis()
	var convArg3 = angleDelta.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, qt4Delta, qt4Orientation, buttons, modifiers, phase)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:190
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, QPoint, QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase, Qt::MouseEventSource)
func NewQWheelEvent_4(pos *qtcore.QPointF, globalPos *qtcore.QPointF, pixelDelta *qtcore.QPoint /*123*/, angleDelta *qtcore.QPoint /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int) *QWheelEvent {
	var convArg0 = pos.GetCthis()
	var convArg1 = globalPos.GetCthis()
	var convArg2 = pixelDelta.GetCthis()
	var convArg3 = angleDelta.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseENS4_16MouseEventSourceE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, qt4Delta, qt4Orientation, buttons, modifiers, phase, source)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:193
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, QPoint, QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase, Qt::MouseEventSource, _Bool)
func NewQWheelEvent_5(pos *qtcore.QPointF, globalPos *qtcore.QPointF, pixelDelta *qtcore.QPoint /*123*/, angleDelta *qtcore.QPoint /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int, inverted bool) *QWheelEvent {
	var convArg0 = pos.GetCthis()
	var convArg1 = globalPos.GetCthis()
	var convArg2 = pixelDelta.GetCthis()
	var convArg3 = angleDelta.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseENS4_16MouseEventSourceEb", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, qt4Delta, qt4Orientation, buttons, modifiers, phase, source, inverted)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:196
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWheelEvent()
func DeleteQWheelEvent(this *QWheelEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 96)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:199
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pixelDelta()
func (this *QWheelEvent) PixelDelta() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent10pixelDeltaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:200
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint angleDelta()
func (this *QWheelEvent) AngleDelta() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent10angleDeltaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int delta()
func (this *QWheelEvent) Delta() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent5deltaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QWheelEvent) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QWheelEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint globalPos()
func (this *QWheelEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QWheelEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QWheelEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX()
func (this *QWheelEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY()
func (this *QWheelEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF()
func (this *QWheelEvent) PosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent4posFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & globalPosF()
func (this *QWheelEvent) GlobalPosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent10globalPosFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons()
func (this *QWheelEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ScrollPhase phase()
func (this *QWheelEvent) Phase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent5phaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool inverted()
func (this *QWheelEvent) Inverted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent8invertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseEventSource source()
func (this *QWheelEvent) Source() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

type QWheelEvent__ = int

const QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120

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
