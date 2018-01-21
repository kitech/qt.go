package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QGraphicsSceneWheelEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneWheelEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func NewQGraphicsSceneWheelEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneWheelEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneWheelEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:139
// index:0
// Public
// void QGraphicsSceneWheelEvent(enum QEvent::Type)
func NewQGraphicsSceneWheelEvent(type_ int) *QGraphicsSceneWheelEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneWheelEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:140
// index:0
// Public virtual
// void ~QGraphicsSceneWheelEvent()
func DeleteQGraphicsSceneWheelEvent(*QGraphicsSceneWheelEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:142
// index:0
// Public
// QPointF pos()
func (this *QGraphicsSceneWheelEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:143
// index:0
// Public
// void setPos(const class QPointF &)
func (this *QGraphicsSceneWheelEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:145
// index:0
// Public
// QPointF scenePos()
func (this *QGraphicsSceneWheelEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:146
// index:0
// Public
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneWheelEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:148
// index:0
// Public
// QPoint screenPos()
func (this *QGraphicsSceneWheelEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:149
// index:0
// Public
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneWheelEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:151
// index:0
// Public
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneWheelEvent) Buttons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:154
// index:0
// Public
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneWheelEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:157
// index:0
// Public
// int delta()
func (this *QGraphicsSceneWheelEvent) Delta() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent5deltaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:158
// index:0
// Public
// void setDelta(int)
func (this *QGraphicsSceneWheelEvent) SetDelta(delta int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent8setDeltaEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &delta)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:160
// index:0
// Public
// Qt::Orientation orientation()
func (this *QGraphicsSceneWheelEvent) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:161
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QGraphicsSceneWheelEvent) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

//  body block end
