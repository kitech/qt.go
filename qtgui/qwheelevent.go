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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QWheelEvent struct {
	*QInputEvent
}
type QWheelEvent_ITF interface {
	QInputEvent_ITF
	QWheelEvent_PTR() *QWheelEvent
}

func (ptr *QWheelEvent) QWheelEvent_PTR() *QWheelEvent { return ptr }

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

/*

 */
func (*QWheelEvent) NewForInherit(pos qtcore.QPointF_ITF, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	return NewQWheelEvent(pos, delta, buttons, modifiers, orient)
}
func NewQWheelEvent(pos qtcore.QPointF_ITF, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFi6QFlagsIN2Qt11MouseButtonEES3_INS4_16KeyboardModifierEENS4_11OrientationE", qtrt.FFI_TYPE_POINTER, convArg0, delta, buttons, modifiers, orient)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, int, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::Orientation)

/*

 */
func (*QWheelEvent) NewForInheritp(pos qtcore.QPointF_ITF, delta int, buttons int, modifiers int) *QWheelEvent {
	return NewQWheelEventp(pos, delta, buttons, modifiers)
}
func NewQWheelEventp(pos qtcore.QPointF_ITF, delta int, buttons int, modifiers int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	// arg: 4, Qt::Orientation=Elaborated, Qt::Orientation=Enum, , Invalid
	orient := 0
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

/*

 */
func (*QWheelEvent) NewForInherit1(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	return NewQWheelEvent1(pos, globalPos, delta, buttons, modifiers, orient)
}
func NewQWheelEvent1(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg1 = globalPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_i6QFlagsIN2Qt11MouseButtonEES3_INS4_16KeyboardModifierEENS4_11OrientationE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, delta, buttons, modifiers, orient)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:181
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, int, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::Orientation)

/*

 */
func (*QWheelEvent) NewForInherit1p(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, delta int, buttons int, modifiers int) *QWheelEvent {
	return NewQWheelEvent1p(pos, globalPos, delta, buttons, modifiers)
}
func NewQWheelEvent1p(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, delta int, buttons int, modifiers int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg1 = globalPos.QPointF_PTR().GetCthis()
	}
	// arg: 5, Qt::Orientation=Elaborated, Qt::Orientation=Enum, , Invalid
	orient := 0
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

/*

 */
func (*QWheelEvent) NewForInherit2(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int) *QWheelEvent {
	return NewQWheelEvent2(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers)
}
func NewQWheelEvent2(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg1 = globalPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pixelDelta != nil && pixelDelta.QPoint_PTR() != nil {
		convArg2 = pixelDelta.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if angleDelta != nil && angleDelta.QPoint_PTR() != nil {
		convArg3 = angleDelta.QPoint_PTR().GetCthis()
	}
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

/*

 */
func (*QWheelEvent) NewForInherit3(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int) *QWheelEvent {
	return NewQWheelEvent3(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers, phase)
}
func NewQWheelEvent3(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg1 = globalPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pixelDelta != nil && pixelDelta.QPoint_PTR() != nil {
		convArg2 = pixelDelta.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if angleDelta != nil && angleDelta.QPoint_PTR() != nil {
		convArg3 = angleDelta.QPoint_PTR().GetCthis()
	}
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

/*

 */
func (*QWheelEvent) NewForInherit4(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int) *QWheelEvent {
	return NewQWheelEvent4(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers, phase, source)
}
func NewQWheelEvent4(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg1 = globalPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pixelDelta != nil && pixelDelta.QPoint_PTR() != nil {
		convArg2 = pixelDelta.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if angleDelta != nil && angleDelta.QPoint_PTR() != nil {
		convArg3 = angleDelta.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseENS4_16MouseEventSourceE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, qt4Delta, qt4Orientation, buttons, modifiers, phase, source)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWheelEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:193
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QWheelEvent(const QPointF &, const QPointF &, QPoint, QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase, Qt::MouseEventSource, bool)

/*

 */
func (*QWheelEvent) NewForInherit5(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int, inverted bool) *QWheelEvent {
	return NewQWheelEvent5(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers, phase, source, inverted)
}
func NewQWheelEvent5(pos qtcore.QPointF_ITF, globalPos qtcore.QPointF_ITF, pixelDelta qtcore.QPoint_ITF /*123*/, angleDelta qtcore.QPoint_ITF /*123*/, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int, inverted bool) *QWheelEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if globalPos != nil && globalPos.QPointF_PTR() != nil {
		convArg1 = globalPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pixelDelta != nil && pixelDelta.QPoint_PTR() != nil {
		convArg2 = pixelDelta.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if angleDelta != nil && angleDelta.QPoint_PTR() != nil {
		convArg3 = angleDelta.QPoint_PTR().GetCthis()
	}
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

/*

 */
func DeleteQWheelEvent(this *QWheelEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWheelEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 96)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:199
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pixelDelta() const

/*

 */
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
// [8] QPoint angleDelta() const

/*

 */
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
// [4] int delta() const

/*

 */
func (this *QWheelEvent) Delta() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent5deltaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QWheelEvent) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos() const

/*

 */
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
// [8] QPoint globalPos() const

/*

 */
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
// [4] int x() const

/*

 */
func (this *QWheelEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*

 */
func (this *QWheelEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX() const

/*

 */
func (this *QWheelEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY() const

/*

 */
func (this *QWheelEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF() const

/*

 */
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
// [16] const QPointF & globalPosF() const

/*

 */
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
// [4] Qt::MouseButtons buttons() const

/*

 */
func (this *QWheelEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ScrollPhase phase() const

/*

 */
func (this *QWheelEvent) Phase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent5phaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool inverted() const

/*

 */
func (this *QWheelEvent) Inverted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent8invertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseEventSource source() const

/*

 */
func (this *QWheelEvent) Source() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*


 */
type QWheelEvent__ = int

//
const QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120

func (this *QWheelEvent) ItemName(val int) string {
	switch val {
	case QWheelEvent__DefaultDeltasPerStep: // 120
		return "DefaultDeltasPerStep"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWheelEvent_ItemName(val int) string {
	var nilthis *QWheelEvent
	return nilthis.ItemName(val)
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
