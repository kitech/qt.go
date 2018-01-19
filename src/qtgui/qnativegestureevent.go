//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QNativeGestureEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:305
// index:0
// void QNativeGestureEvent(Qt::NativeGestureType, const class QPointF &, const class QPointF &, const class QPointF &, qreal, ulong, quint64)
func NewQNativeGestureEvent(type_ int, localPos unsafe.Pointer, windowPos unsafe.Pointer, screenPos unsafe.Pointer, value float64, sequenceId uint, intArgument uint64) *QNativeGestureEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QNativeGestureEventC2EN2Qt17NativeGestureTypeERK7QPointFS4_S4_dmy", ffiqt.FFI_TYPE_VOID, cthis, &type_, localPos, windowPos, screenPos, &value, &sequenceId, &intArgument)
	gopp.ErrPrint(err, rv)
	return &QNativeGestureEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:308
// index:1
// void QNativeGestureEvent(Qt::NativeGestureType, const class QTouchDevice *, const class QPointF &, const class QPointF &, const class QPointF &, qreal, ulong, quint64)
func NewQNativeGestureEvent_1(type_ int, dev unsafe.Pointer, localPos unsafe.Pointer, windowPos unsafe.Pointer, screenPos unsafe.Pointer, value float64, sequenceId uint, intArgument uint64) *QNativeGestureEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QNativeGestureEventC2EN2Qt17NativeGestureTypeEPK12QTouchDeviceRK7QPointFS7_S7_dmy", ffiqt.FFI_TYPE_VOID, cthis, &type_, dev, localPos, windowPos, screenPos, &value, &sequenceId, &intArgument)
	gopp.ErrPrint(err, rv)
	return &QNativeGestureEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:310
// index:0
// virtual
// void ~QNativeGestureEvent()
func DeleteQNativeGestureEvent(*QNativeGestureEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QNativeGestureEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:311
// index:0
// inline
// Qt::NativeGestureType gestureType()
func (this *QNativeGestureEvent) GestureType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent11gestureTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:312
// index:0
// inline
// qreal value()
func (this *QNativeGestureEvent) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:315
// index:0
// inline
// const QPoint pos()
func (this *QNativeGestureEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:316
// index:0
// inline
// const QPoint globalPos()
func (this *QNativeGestureEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:318
// index:0
// inline
// const QPointF & localPos()
func (this *QNativeGestureEvent) LocalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent8localPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:319
// index:0
// inline
// const QPointF & windowPos()
func (this *QNativeGestureEvent) WindowPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9windowPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:320
// index:0
// inline
// const QPointF & screenPos()
func (this *QNativeGestureEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:322
// index:0
// const QTouchDevice * device()
func (this *QNativeGestureEvent) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent6deviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
