//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QMoveEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:421
// index:0
// void QMoveEvent(const class QPoint &, const class QPoint &)
func NewQMoveEvent(pos unsafe.Pointer, oldPos unsafe.Pointer) *QMoveEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMoveEventC2ERK6QPointS2_", ffiqt.FFI_TYPE_VOID, cthis, pos, oldPos)
	gopp.ErrPrint(err, rv)
	return &QMoveEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:422
// index:0
// virtual
// void ~QMoveEvent()
func DeleteQMoveEvent(*QMoveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QMoveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:424
// index:0
// inline
// const QPoint & pos()
func (this *QMoveEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMoveEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:425
// index:0
// inline
// const QPoint & oldPos()
func (this *QMoveEvent) OldPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QMoveEvent6oldPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
