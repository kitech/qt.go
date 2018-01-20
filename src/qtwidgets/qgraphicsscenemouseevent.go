//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

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
type QGraphicsSceneMouseEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneMouseEvent) GetCthis() unsafe.Pointer {
	return this.QGraphicsSceneEvent.GetCthis()
}
func NewQGraphicsSceneMouseEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMouseEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneMouseEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:85
// index:0
// Public
// void QGraphicsSceneMouseEvent(enum QEvent::Type)
func NewQGraphicsSceneMouseEvent(type_ int) *QGraphicsSceneMouseEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneMouseEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:86
// index:0
// Public virtual
// void ~QGraphicsSceneMouseEvent()
func DeleteQGraphicsSceneMouseEvent(*QGraphicsSceneMouseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:88
// index:0
// Public
// QPointF pos()
func (this *QGraphicsSceneMouseEvent) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:89
// index:0
// Public
// void setPos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:91
// index:0
// Public
// QPointF scenePos()
func (this *QGraphicsSceneMouseEvent) ScenePos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:92
// index:0
// Public
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:94
// index:0
// Public
// QPoint screenPos()
func (this *QGraphicsSceneMouseEvent) ScreenPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:95
// index:0
// Public
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneMouseEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:97
// index:0
// Public
// QPointF buttonDownPos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownPos(button int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent13buttonDownPosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:98
// index:0
// Public
// void setButtonDownPos(Qt::MouseButton, const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownPos(button int, pos *qtcore.QPointF) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent16setButtonDownPosEN2Qt11MouseButtonERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:100
// index:0
// Public
// QPointF buttonDownScenePos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownScenePos(button int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent18buttonDownScenePosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:101
// index:0
// Public
// void setButtonDownScenePos(Qt::MouseButton, const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownScenePos(button int, pos *qtcore.QPointF) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent21setButtonDownScenePosEN2Qt11MouseButtonERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:103
// index:0
// Public
// QPoint buttonDownScreenPos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent19buttonDownScreenPosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:104
// index:0
// Public
// void setButtonDownScreenPos(Qt::MouseButton, const class QPoint &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownScreenPos(button int, pos *qtcore.QPoint) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent22setButtonDownScreenPosEN2Qt11MouseButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:106
// index:0
// Public
// QPointF lastPos()
func (this *QGraphicsSceneMouseEvent) LastPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent7lastPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:107
// index:0
// Public
// void setLastPos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetLastPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:109
// index:0
// Public
// QPointF lastScenePos()
func (this *QGraphicsSceneMouseEvent) LastScenePos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent12lastScenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:110
// index:0
// Public
// void setLastScenePos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetLastScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:112
// index:0
// Public
// QPoint lastScreenPos()
func (this *QGraphicsSceneMouseEvent) LastScreenPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:113
// index:0
// Public
// void setLastScreenPos(const class QPoint &)
func (this *QGraphicsSceneMouseEvent) SetLastScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:115
// index:0
// Public
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneMouseEvent) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:118
// index:0
// Public
// Qt::MouseButton button()
func (this *QGraphicsSceneMouseEvent) Button() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:119
// index:0
// Public
// void setButton(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) SetButton(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent9setButtonEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:121
// index:0
// Public
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneMouseEvent) Modifiers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:124
// index:0
// Public
// Qt::MouseEventSource source()
func (this *QGraphicsSceneMouseEvent) Source() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:125
// index:0
// Public
// void setSource(Qt::MouseEventSource)
func (this *QGraphicsSceneMouseEvent) SetSource(source int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent9setSourceEN2Qt16MouseEventSourceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:127
// index:0
// Public
// Qt::MouseEventFlags flags()
func (this *QGraphicsSceneMouseEvent) Flags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
