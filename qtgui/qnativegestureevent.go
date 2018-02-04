package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	*QInputEvent
}

func (this *QNativeGestureEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QNativeGestureEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQNativeGestureEventFromPointer(cthis unsafe.Pointer) *QNativeGestureEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QNativeGestureEvent{bcthis0}
}
func (*QNativeGestureEvent) NewFromPointer(cthis unsafe.Pointer) *QNativeGestureEvent {
	return NewQNativeGestureEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNativeGestureEvent(Qt::NativeGestureType, const QPointF &, const QPointF &, const QPointF &, qreal, ulong, quint64)
func NewQNativeGestureEvent(type_ int, localPos *qtcore.QPointF, windowPos *qtcore.QPointF, screenPos *qtcore.QPointF, value float64, sequenceId uint, intArgument uint64) *QNativeGestureEvent {
	var convArg1 = localPos.GetCthis()
	var convArg2 = windowPos.GetCthis()
	var convArg3 = screenPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QNativeGestureEventC2EN2Qt17NativeGestureTypeERK7QPointFS4_S4_dmy", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, value, sequenceId, intArgument)
	gopp.ErrPrint(err, rv)
	gothis := NewQNativeGestureEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNativeGestureEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:308
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNativeGestureEvent(Qt::NativeGestureType, const QTouchDevice *, const QPointF &, const QPointF &, const QPointF &, qreal, ulong, quint64)
func NewQNativeGestureEvent_1(type_ int, dev *QTouchDevice /*777 const QTouchDevice **/, localPos *qtcore.QPointF, windowPos *qtcore.QPointF, screenPos *qtcore.QPointF, value float64, sequenceId uint, intArgument uint64) *QNativeGestureEvent {
	var convArg1 = dev.GetCthis()
	var convArg2 = localPos.GetCthis()
	var convArg3 = windowPos.GetCthis()
	var convArg4 = screenPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QNativeGestureEventC2EN2Qt17NativeGestureTypeEPK12QTouchDeviceRK7QPointFS7_S7_dmy", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, convArg4, value, sequenceId, intArgument)
	gopp.ErrPrint(err, rv)
	gothis := NewQNativeGestureEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNativeGestureEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:310
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNativeGestureEvent()
func DeleteQNativeGestureEvent(this *QNativeGestureEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QNativeGestureEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 112)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:311
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::NativeGestureType gestureType()
func (this *QNativeGestureEvent) GestureType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent11gestureTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:312
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal value()
func (this *QNativeGestureEvent) Value() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent5valueEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:315
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint pos()
func (this *QNativeGestureEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:316
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint globalPos()
func (this *QNativeGestureEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:318
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & localPos()
func (this *QNativeGestureEvent) LocalPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent8localPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:319
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & windowPos()
func (this *QNativeGestureEvent) WindowPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9windowPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:320
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & screenPos()
func (this *QNativeGestureEvent) ScreenPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:322
// index:0
// Public Visibility=Default Availability=Available
// [8] const QTouchDevice * device()
func (this *QNativeGestureEvent) Device() *QTouchDevice /*777 const QTouchDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QNativeGestureEvent6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTouchDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
