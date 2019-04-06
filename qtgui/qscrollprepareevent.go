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
type QScrollPrepareEvent struct {
	*qtcore.QEvent
}
type QScrollPrepareEvent_ITF interface {
	qtcore.QEvent_ITF
	QScrollPrepareEvent_PTR() *QScrollPrepareEvent
}

func (ptr *QScrollPrepareEvent) QScrollPrepareEvent_PTR() *QScrollPrepareEvent { return ptr }

func (this *QScrollPrepareEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QScrollPrepareEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQScrollPrepareEventFromPointer(cthis unsafe.Pointer) *QScrollPrepareEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QScrollPrepareEvent{bcthis0}
}
func (*QScrollPrepareEvent) NewFromPointer(cthis unsafe.Pointer) *QScrollPrepareEvent {
	return NewQScrollPrepareEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:995
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollPrepareEvent(const QPointF &)

/*

 */
func (*QScrollPrepareEvent) NewForInherit(startPos qtcore.QPointF_ITF) *QScrollPrepareEvent {
	return NewQScrollPrepareEvent(startPos)
}
func NewQScrollPrepareEvent(startPos qtcore.QPointF_ITF) *QScrollPrepareEvent {
	var convArg0 unsafe.Pointer
	if startPos != nil && startPos.QPointF_PTR() != nil {
		convArg0 = startPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollPrepareEventC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollPrepareEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQScrollPrepareEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:996
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScrollPrepareEvent()

/*

 */
func DeleteQScrollPrepareEvent(this *QScrollPrepareEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollPrepareEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 112)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:998
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF startPos() const

/*

 */
func (this *QScrollPrepareEvent) StartPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent8startPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1000
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF viewportSize() const

/*

 */
func (this *QScrollPrepareEvent) ViewportSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent12viewportSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1001
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF contentPosRange() const

/*

 */
func (this *QScrollPrepareEvent) ContentPosRange() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent15contentPosRangeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1002
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF contentPos() const

/*

 */
func (this *QScrollPrepareEvent) ContentPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent10contentPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1004
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewportSize(const QSizeF &)

/*

 */
func (this *QScrollPrepareEvent) SetViewportSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1005
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentPosRange(const QRectF &)

/*

 */
func (this *QScrollPrepareEvent) SetContentPosRange(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1006
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentPos(const QPointF &)

/*

 */
func (this *QScrollPrepareEvent) SetContentPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollPrepareEvent13setContentPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10713() {
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
