//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QScrollPrepareEvent struct {
	*qtcore.QEvent
}

func (this *QScrollPrepareEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:990
// index:0
// void QScrollPrepareEvent(const class QPointF &)
func NewQScrollPrepareEvent(startPos unsafe.Pointer) *QScrollPrepareEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEventC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, startPos)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollPrepareEventFromPointer(cthis)
	return gothis
}
func NewQScrollPrepareEventFromPointer(cthis unsafe.Pointer) *QScrollPrepareEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QScrollPrepareEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:991
// index:0
// virtual
// void ~QScrollPrepareEvent()
func DeleteQScrollPrepareEvent(*QScrollPrepareEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:993
// index:0
// QPointF startPos()
func (this *QScrollPrepareEvent) StartPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent8startPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:995
// index:0
// QSizeF viewportSize()
func (this *QScrollPrepareEvent) ViewportSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent12viewportSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:996
// index:0
// QRectF contentPosRange()
func (this *QScrollPrepareEvent) ContentPosRange() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent15contentPosRangeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:997
// index:0
// QPointF contentPos()
func (this *QScrollPrepareEvent) ContentPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollPrepareEvent10contentPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:999
// index:0
// void setViewportSize(const class QSizeF &)
func (this *QScrollPrepareEvent) SetViewportSize(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1000
// index:0
// void setContentPosRange(const class QRectF &)
func (this *QScrollPrepareEvent) SetContentPosRange(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1001
// index:0
// void setContentPos(const class QPointF &)
func (this *QScrollPrepareEvent) SetContentPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollPrepareEvent13setContentPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
