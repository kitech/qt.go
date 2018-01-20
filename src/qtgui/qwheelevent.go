//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QWheelEvent struct {
	*QInputEvent
}

func (this *QWheelEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:178
// index:0
// void QWheelEvent(const class QPointF &, int, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::Orientation)
func NewQWheelEvent(pos unsafe.Pointer, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFi6QFlagsIN2Qt11MouseButtonEES3_INS4_16KeyboardModifierEENS4_11OrientationE", ffiqt.FFI_TYPE_VOID, cthis, pos, &delta, &buttons, &modifiers, &orient)
	gopp.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(cthis)
	return gothis
}
func NewQWheelEventFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QWheelEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:181
// index:1
// void QWheelEvent(const class QPointF &, const class QPointF &, int, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::Orientation)
func NewQWheelEvent_1(pos unsafe.Pointer, globalPos unsafe.Pointer, delta int, buttons int, modifiers int, orient int) *QWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_i6QFlagsIN2Qt11MouseButtonEES3_INS4_16KeyboardModifierEENS4_11OrientationE", ffiqt.FFI_TYPE_VOID, cthis, pos, globalPos, &delta, &buttons, &modifiers, &orient)
	gopp.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:184
// index:2
// void QWheelEvent(const class QPointF &, const class QPointF &, class QPoint, class QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers)
func NewQWheelEvent_2(pos unsafe.Pointer, globalPos unsafe.Pointer, pixelDelta unsafe.Pointer, angleDelta unsafe.Pointer, qt4Delta int, qt4Orientation int, buttons int, modifiers int) *QWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, cthis, pos, globalPos, pixelDelta, angleDelta, &qt4Delta, &qt4Orientation, &buttons, &modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:187
// index:3
// void QWheelEvent(const class QPointF &, const class QPointF &, class QPoint, class QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase)
func NewQWheelEvent_3(pos unsafe.Pointer, globalPos unsafe.Pointer, pixelDelta unsafe.Pointer, angleDelta unsafe.Pointer, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int) *QWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseE", ffiqt.FFI_TYPE_VOID, cthis, pos, globalPos, pixelDelta, angleDelta, &qt4Delta, &qt4Orientation, &buttons, &modifiers, &phase)
	gopp.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:190
// index:4
// void QWheelEvent(const class QPointF &, const class QPointF &, class QPoint, class QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase, Qt::MouseEventSource)
func NewQWheelEvent_4(pos unsafe.Pointer, globalPos unsafe.Pointer, pixelDelta unsafe.Pointer, angleDelta unsafe.Pointer, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int) *QWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseENS4_16MouseEventSourceE", ffiqt.FFI_TYPE_VOID, cthis, pos, globalPos, pixelDelta, angleDelta, &qt4Delta, &qt4Orientation, &buttons, &modifiers, &phase, &source)
	gopp.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:193
// index:5
// void QWheelEvent(const class QPointF &, const class QPointF &, class QPoint, class QPoint, int, Qt::Orientation, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::ScrollPhase, Qt::MouseEventSource, _Bool)
func NewQWheelEvent_5(pos unsafe.Pointer, globalPos unsafe.Pointer, pixelDelta unsafe.Pointer, angleDelta unsafe.Pointer, qt4Delta int, qt4Orientation int, buttons int, modifiers int, phase int, source int, inverted bool) *QWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventC2ERK7QPointFS2_6QPointS3_iN2Qt11OrientationE6QFlagsINS4_11MouseButtonEES6_INS4_16KeyboardModifierEENS4_11ScrollPhaseENS4_16MouseEventSourceEb", ffiqt.FFI_TYPE_VOID, cthis, pos, globalPos, pixelDelta, angleDelta, &qt4Delta, &qt4Orientation, &buttons, &modifiers, &phase, &source, &inverted)
	gopp.ErrPrint(err, rv)
	gothis := NewQWheelEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:196
// index:0
// virtual
// void ~QWheelEvent()
func DeleteQWheelEvent(*QWheelEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:199
// index:0
// inline
// QPoint pixelDelta()
func (this *QWheelEvent) PixelDelta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10pixelDeltaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:200
// index:0
// inline
// QPoint angleDelta()
func (this *QWheelEvent) AngleDelta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10angleDeltaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:202
// index:0
// inline
// int delta()
func (this *QWheelEvent) Delta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent5deltaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:203
// index:0
// inline
// Qt::Orientation orientation()
func (this *QWheelEvent) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:206
// index:0
// inline
// QPoint pos()
func (this *QWheelEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:207
// index:0
// inline
// QPoint globalPos()
func (this *QWheelEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:208
// index:0
// inline
// int x()
func (this *QWheelEvent) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent1xEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:209
// index:0
// inline
// int y()
func (this *QWheelEvent) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent1yEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:210
// index:0
// inline
// int globalX()
func (this *QWheelEvent) GlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7globalXEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:211
// index:0
// inline
// int globalY()
func (this *QWheelEvent) GlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7globalYEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:213
// index:0
// inline
// const QPointF & posF()
func (this *QWheelEvent) PosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent4posFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:214
// index:0
// inline
// const QPointF & globalPosF()
func (this *QWheelEvent) GlobalPosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10globalPosFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:216
// index:0
// inline
// Qt::MouseButtons buttons()
func (this *QWheelEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:218
// index:0
// inline
// Qt::ScrollPhase phase()
func (this *QWheelEvent) Phase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent5phaseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:219
// index:0
// inline
// bool inverted()
func (this *QWheelEvent) Inverted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent8invertedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:221
// index:0
// inline
// Qt::MouseEventSource source()
func (this *QWheelEvent) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
