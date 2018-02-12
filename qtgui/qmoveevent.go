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
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QMoveEvent struct {
	*qtcore.QEvent
}

func (this *QMoveEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QMoveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQMoveEventFromPointer(cthis unsafe.Pointer) *QMoveEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QMoveEvent{bcthis0}
}
func (*QMoveEvent) NewFromPointer(cthis unsafe.Pointer) *QMoveEvent {
	return NewQMoveEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:421
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMoveEvent(const QPoint &, const QPoint &)
func NewQMoveEvent(pos *qtcore.QPoint, oldPos *qtcore.QPoint) *QMoveEvent {
	var convArg0 = pos.GetCthis()
	var convArg1 = oldPos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMoveEventC2ERK6QPointS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:422
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMoveEvent()
func DeleteQMoveEvent(this *QMoveEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QMoveEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:424
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & pos()
func (this *QMoveEvent) Pos() *qtcore.QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMoveEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:425
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & oldPos()
func (this *QMoveEvent) OldPos() *qtcore.QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QMoveEvent6oldPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
