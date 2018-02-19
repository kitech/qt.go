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

type QNativeGestureEvent struct {
	*QInputEvent
}
type QNativeGestureEvent_ITF interface {
	QInputEvent_ITF
	QNativeGestureEvent_PTR() *QNativeGestureEvent
}

func (ptr *QNativeGestureEvent) QNativeGestureEvent_PTR() *QNativeGestureEvent { return ptr }

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
func NewQNativeGestureEvent(type_ int, localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, value float64, sequenceId uint, intArgument uint64) *QNativeGestureEvent {
	var convArg1 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg1 = localPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if windowPos != nil && windowPos.QPointF_PTR() != nil {
		convArg2 = windowPos.QPointF_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if screenPos != nil && screenPos.QPointF_PTR() != nil {
		convArg3 = screenPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QNativeGestureEventC2EN2Qt17NativeGestureTypeERK7QPointFS4_S4_dmy", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, value, sequenceId, intArgument)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNativeGestureEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNativeGestureEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:308
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNativeGestureEvent(Qt::NativeGestureType, const QTouchDevice *, const QPointF &, const QPointF &, const QPointF &, qreal, ulong, quint64)
func NewQNativeGestureEvent_1(type_ int, dev QTouchDevice_ITF /*777 const QTouchDevice **/, localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, value float64, sequenceId uint, intArgument uint64) *QNativeGestureEvent {
	var convArg1 unsafe.Pointer
	if dev != nil && dev.QTouchDevice_PTR() != nil {
		convArg1 = dev.QTouchDevice_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg2 = localPos.QPointF_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if windowPos != nil && windowPos.QPointF_PTR() != nil {
		convArg3 = windowPos.QPointF_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if screenPos != nil && screenPos.QPointF_PTR() != nil {
		convArg4 = screenPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QNativeGestureEventC2EN2Qt17NativeGestureTypeEPK12QTouchDeviceRK7QPointFS7_S7_dmy", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, convArg4, value, sequenceId, intArgument)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNativeGestureEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNativeGestureEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:310
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNativeGestureEvent()
func DeleteQNativeGestureEvent(this *QNativeGestureEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QNativeGestureEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 112)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:311
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::NativeGestureType gestureType() const
func (this *QNativeGestureEvent) GestureType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent11gestureTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:312
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal value() const
func (this *QNativeGestureEvent) Value() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent5valueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qevent.h:315
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint pos() const
func (this *QNativeGestureEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:316
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint globalPos() const
func (this *QNativeGestureEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:318
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & localPos() const
func (this *QNativeGestureEvent) LocalPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent8localPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:319
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & windowPos() const
func (this *QNativeGestureEvent) WindowPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9windowPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:320
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & screenPos() const
func (this *QNativeGestureEvent) ScreenPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:322
// index:0
// Public Visibility=Default Availability=Available
// [8] const QTouchDevice * device() const
func (this *QNativeGestureEvent) Device() *QTouchDevice /*777 const QTouchDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QNativeGestureEvent6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTouchDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
