package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QScrollEvent struct {
	*qtcore.QEvent
}

func (this *QScrollEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QScrollEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQScrollEventFromPointer(cthis unsafe.Pointer) *QScrollEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QScrollEvent{bcthis0}
}
func (*QScrollEvent) NewFromPointer(cthis unsafe.Pointer) *QScrollEvent {
	return NewQScrollEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:1022
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollEvent(const QPointF &, const QPointF &, enum QScrollEvent::ScrollState)
func NewQScrollEvent(contentPos *qtcore.QPointF, overshoot *qtcore.QPointF, scrollState int) *QScrollEvent {
	var convArg0 = contentPos.GetCthis()
	var convArg1 = overshoot.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QScrollEventC2ERK7QPointFS2_NS_11ScrollStateE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, scrollState)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQScrollEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:1023
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScrollEvent()
func DeleteQScrollEvent(this *QScrollEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QScrollEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:1025
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF contentPos()
func (this *QScrollEvent) ContentPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QScrollEvent10contentPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1026
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF overshootDistance()
func (this *QScrollEvent) OvershootDistance() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QScrollEvent17overshootDistanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1027
// index:0
// Public Visibility=Default Availability=Available
// [4] QScrollEvent::ScrollState scrollState()
func (this *QScrollEvent) ScrollState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QScrollEvent11scrollStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

type QScrollEvent__ScrollState = int

const QScrollEvent__ScrollStarted QScrollEvent__ScrollState = 0
const QScrollEvent__ScrollUpdated QScrollEvent__ScrollState = 1
const QScrollEvent__ScrollFinished QScrollEvent__ScrollState = 2

//  body block end
