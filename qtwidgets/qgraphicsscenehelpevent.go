package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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

type QGraphicsSceneHelpEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneHelpEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneHelpEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneHelpEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHelpEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneHelpEvent{bcthis0}
}
func (*QGraphicsSceneHelpEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHelpEvent {
	return NewQGraphicsSceneHelpEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneHelpEvent(enum QEvent::Type)
func NewQGraphicsSceneHelpEvent(type_ int) *QGraphicsSceneHelpEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHelpEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneHelpEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneHelpEvent()
func DeleteQGraphicsSceneHelpEvent(this *QGraphicsSceneHelpEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:237
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos()
func (this *QGraphicsSceneHelpEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSceneHelpEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)
func (this *QGraphicsSceneHelpEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:240
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos()
func (this *QGraphicsSceneHelpEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSceneHelpEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)
func (this *QGraphicsSceneHelpEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
