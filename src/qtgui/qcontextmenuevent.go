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
	*QInputEvent
}

func (this *QContextMenuEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}
func NewQContextMenuEventFromPointer(cthis unsafe.Pointer) *QContextMenuEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QContextMenuEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:513
// index:0
// Public
// void QContextMenuEvent(enum QContextMenuEvent::Reason, const class QPoint &, const class QPoint &)
func NewQContextMenuEvent(reason int, pos *qtcore.QPoint, globalPos *qtcore.QPoint) *QContextMenuEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	var convArg1 = pos.GetCthis()
	var convArg2 = globalPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPointS3_", ffiqt.FFI_TYPE_VOID, cthis, &reason, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQContextMenuEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:514
// index:1
// Public
// void QContextMenuEvent(enum QContextMenuEvent::Reason, const class QPoint &)
func NewQContextMenuEvent_1(reason int, pos *qtcore.QPoint) *QContextMenuEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, &reason, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQContextMenuEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:515
// index:0
// Public virtual
// void ~QContextMenuEvent()
func DeleteQContextMenuEvent(*QContextMenuEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QContextMenuEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:517
// index:0
// Public inline
// int x()
func (this *QContextMenuEvent) X() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:518
// index:0
// Public inline
// int y()
func (this *QContextMenuEvent) Y() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:519
// index:0
// Public inline
// int globalX()
func (this *QContextMenuEvent) GlobalX() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:520
// index:0
// Public inline
// int globalY()
func (this *QContextMenuEvent) GlobalY() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:522
// index:0
// Public inline
// const QPoint & pos()
func (this *QContextMenuEvent) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:523
// index:0
// Public inline
// const QPoint & globalPos()
func (this *QContextMenuEvent) GlobalPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:525
// index:0
// Public inline
// QContextMenuEvent::Reason reason()
func (this *QContextMenuEvent) Reason() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QContextMenuEvent6reasonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
