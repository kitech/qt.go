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
import "mkuse/cffiqt"
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
type QScrollPrepareEvent struct {
	*qtcore.QEvent
}

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

// /usr/include/qt/QtGui/qevent.h:990
// index:0
// Public
// void QScrollPrepareEvent(const class QPointF &)
func NewQScrollPrepareEvent(startPos *qtcore.QPointF) *QScrollPrepareEvent {
	cthis := qtrt.Calloc(1, 256) // 112
	var convArg0 = startPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEventC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollPrepareEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:991
// index:0
// Public virtual
// void ~QScrollPrepareEvent()
func DeleteQScrollPrepareEvent(*QScrollPrepareEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:993
// index:0
// Public
// QPointF startPos()
func (this *QScrollPrepareEvent) StartPos() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent8startPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:995
// index:0
// Public
// QSizeF viewportSize()
func (this *QScrollPrepareEvent) ViewportSize() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent12viewportSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:996
// index:0
// Public
// QRectF contentPosRange()
func (this *QScrollPrepareEvent) ContentPosRange() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent15contentPosRangeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:997
// index:0
// Public
// QPointF contentPos()
func (this *QScrollPrepareEvent) ContentPos() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent10contentPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:999
// index:0
// Public
// void setViewportSize(const class QSizeF &)
func (this *QScrollPrepareEvent) SetViewportSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1000
// index:0
// Public
// void setContentPosRange(const class QRectF &)
func (this *QScrollPrepareEvent) SetContentPosRange(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1001
// index:0
// Public
// void setContentPos(const class QPointF &)
func (this *QScrollPrepareEvent) SetContentPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEvent13setContentPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
