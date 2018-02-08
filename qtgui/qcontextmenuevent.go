package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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

type QContextMenuEvent struct {
	*QInputEvent
}

func (this *QContextMenuEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QContextMenuEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQContextMenuEventFromPointer(cthis unsafe.Pointer) *QContextMenuEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QContextMenuEvent{bcthis0}
}
func (*QContextMenuEvent) NewFromPointer(cthis unsafe.Pointer) *QContextMenuEvent {
	return NewQContextMenuEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:511
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QContextMenuEvent(enum QContextMenuEvent::Reason, const QPoint &, const QPoint &, Qt::KeyboardModifiers)
func NewQContextMenuEvent(reason int, pos *qtcore.QPoint, globalPos *qtcore.QPoint, modifiers int) *QContextMenuEvent {
	var convArg1 = pos.GetCthis()
	var convArg2 = globalPos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPointS3_6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, reason, convArg1, convArg2, modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQContextMenuEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQContextMenuEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:513
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QContextMenuEvent(enum QContextMenuEvent::Reason, const QPoint &, const QPoint &)
func NewQContextMenuEvent_1(reason int, pos *qtcore.QPoint, globalPos *qtcore.QPoint) *QContextMenuEvent {
	var convArg1 = pos.GetCthis()
	var convArg2 = globalPos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPointS3_", qtrt.FFI_TYPE_POINTER, reason, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQContextMenuEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQContextMenuEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:514
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QContextMenuEvent(enum QContextMenuEvent::Reason, const QPoint &)
func NewQContextMenuEvent_2(reason int, pos *qtcore.QPoint) *QContextMenuEvent {
	var convArg1 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QContextMenuEventC2ENS_6ReasonERK6QPoint", qtrt.FFI_TYPE_POINTER, reason, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQContextMenuEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQContextMenuEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:515
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QContextMenuEvent()
func DeleteQContextMenuEvent(this *QContextMenuEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QContextMenuEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:517
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QContextMenuEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:518
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QContextMenuEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:519
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX()
func (this *QContextMenuEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:520
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY()
func (this *QContextMenuEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:522
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & pos()
func (this *QContextMenuEvent) Pos() *qtcore.QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:523
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & globalPos()
func (this *QContextMenuEvent) GlobalPos() *qtcore.QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:525
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QContextMenuEvent::Reason reason()
func (this *QContextMenuEvent) Reason() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QContextMenuEvent6reasonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

type QContextMenuEvent__Reason = int

const QContextMenuEvent__Mouse QContextMenuEvent__Reason = 0
const QContextMenuEvent__Keyboard QContextMenuEvent__Reason = 1
const QContextMenuEvent__Other QContextMenuEvent__Reason = 2

//  body block end
