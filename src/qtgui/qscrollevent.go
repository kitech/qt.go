//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
type QScrollEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:1022
// index:0
// void QScrollEvent(const class QPointF &, const class QPointF &, enum QScrollEvent::ScrollState)
func NewQScrollEvent(contentPos unsafe.Pointer, overshoot unsafe.Pointer, scrollState int) *QScrollEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QScrollEventC2ERK7QPointFS2_NS_11ScrollStateE", ffiqt.FFI_TYPE_VOID, cthis, contentPos, overshoot, &scrollState)
	gopp.ErrPrint(err, rv)
	return &QScrollEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:1023
// index:0
// virtual
// void ~QScrollEvent()
func DeleteQScrollEvent(*QScrollEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QScrollEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1025
// index:0
// QPointF contentPos()
func (this *QScrollEvent) ContentPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QScrollEvent10contentPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1026
// index:0
// QPointF overshootDistance()
func (this *QScrollEvent) OvershootDistance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QScrollEvent17overshootDistanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1027
// index:0
// QScrollEvent::ScrollState scrollState()
func (this *QScrollEvent) ScrollState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QScrollEvent11scrollStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
