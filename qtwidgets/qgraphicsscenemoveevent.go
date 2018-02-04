package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QGraphicsSceneMoveEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneMoveEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneMoveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneMoveEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMoveEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneMoveEvent{bcthis0}
}
func (*QGraphicsSceneMoveEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMoveEvent {
	return NewQGraphicsSceneMoveEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:313
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneMoveEvent()
func NewQGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEventC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneMoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneMoveEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:314
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneMoveEvent()
func DeleteQGraphicsSceneMoveEvent(this *QGraphicsSceneMoveEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:316
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF oldPos()
func (this *QGraphicsSceneMoveEvent) OldPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSceneMoveEvent6oldPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:317
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOldPos(const QPointF &)
func (this *QGraphicsSceneMoveEvent) SetOldPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEvent9setOldPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:319
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF newPos()
func (this *QGraphicsSceneMoveEvent) NewPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSceneMoveEvent6newPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNewPos(const QPointF &)
func (this *QGraphicsSceneMoveEvent) SetNewPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEvent9setNewPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
