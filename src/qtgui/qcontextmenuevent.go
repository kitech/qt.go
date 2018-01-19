//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QContextMenuEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:513
// index:0
// void QContextMenuEvent(enum QContextMenuEvent::Reason, const class QPoint &, const class QPoint &)
func NewQContextMenuEvent(reason int, pos unsafe.Pointer, globalPos unsafe.Pointer) *QContextMenuEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPointS3_", ffiqt.FFI_TYPE_VOID, cthis, &reason, pos, globalPos)
	gopp.ErrPrint(err, rv)
	return &QContextMenuEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:514
// index:1
// void QContextMenuEvent(enum QContextMenuEvent::Reason, const class QPoint &)
func NewQContextMenuEvent_1(reason int, pos unsafe.Pointer) *QContextMenuEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, &reason, pos)
	gopp.ErrPrint(err, rv)
	return &QContextMenuEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:515
// index:0
// virtual
// void ~QContextMenuEvent()
func DeleteQContextMenuEvent(*QContextMenuEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QContextMenuEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:517
// index:0
// inline
// int x()
func (this *QContextMenuEvent) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:518
// index:0
// inline
// int y()
func (this *QContextMenuEvent) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:519
// index:0
// inline
// int globalX()
func (this *QContextMenuEvent) GlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent7globalXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:520
// index:0
// inline
// int globalY()
func (this *QContextMenuEvent) GlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent7globalYEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:522
// index:0
// inline
// const QPoint & pos()
func (this *QContextMenuEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:523
// index:0
// inline
// const QPoint & globalPos()
func (this *QContextMenuEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:525
// index:0
// inline
// QContextMenuEvent::Reason reason()
func (this *QContextMenuEvent) Reason() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent6reasonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
