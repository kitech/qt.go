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
// extern C begin: 9
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
type QScrollEvent struct {
	*qtcore.QEvent
}
type QScrollEvent_ITF interface {
	qtcore.QEvent_ITF
	QScrollEvent_PTR() *QScrollEvent
}

func (ptr *QScrollEvent) QScrollEvent_PTR() *QScrollEvent { return ptr }

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
// [-2] void QScrollEvent(const QPointF &, const QPointF &, QScrollEvent::ScrollState)

/*

 */
func (*QScrollEvent) NewForInherit(contentPos qtcore.QPointF_ITF, overshoot qtcore.QPointF_ITF, scrollState int) *QScrollEvent {
	return NewQScrollEvent(contentPos, overshoot, scrollState)
}
func NewQScrollEvent(contentPos qtcore.QPointF_ITF, overshoot qtcore.QPointF_ITF, scrollState int) *QScrollEvent {
	var convArg0 unsafe.Pointer
	if contentPos != nil && contentPos.QPointF_PTR() != nil {
		convArg0 = contentPos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if overshoot != nil && overshoot.QPointF_PTR() != nil {
		convArg1 = overshoot.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QScrollEventC2ERK7QPointFS2_NS_11ScrollStateE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, scrollState)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQScrollEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:1023
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScrollEvent()

/*

 */
func DeleteQScrollEvent(this *QScrollEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QScrollEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:1025
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF contentPos() const

/*

 */
func (this *QScrollEvent) ContentPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QScrollEvent10contentPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1026
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF overshootDistance() const

/*

 */
func (this *QScrollEvent) OvershootDistance() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QScrollEvent17overshootDistanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1027
// index:0
// Public Visibility=Default Availability=Available
// [4] QScrollEvent::ScrollState scrollState() const

/*

 */
func (this *QScrollEvent) ScrollState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QScrollEvent11scrollStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*


 */
type QScrollEvent__ScrollState = int

//
const QScrollEvent__ScrollStarted QScrollEvent__ScrollState = 0

//
const QScrollEvent__ScrollUpdated QScrollEvent__ScrollState = 1

//
const QScrollEvent__ScrollFinished QScrollEvent__ScrollState = 2

func (this *QScrollEvent) ScrollStateItemName(val int) string {
	switch val {
	case QScrollEvent__ScrollStarted: // 0
		return "ScrollStarted"
	case QScrollEvent__ScrollUpdated: // 1
		return "ScrollUpdated"
	case QScrollEvent__ScrollFinished: // 2
		return "ScrollFinished"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QScrollEvent_ScrollStateItemName(val int) string {
	var nilthis *QScrollEvent
	return nilthis.ScrollStateItemName(val)
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
