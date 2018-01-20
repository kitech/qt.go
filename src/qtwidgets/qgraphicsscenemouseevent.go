//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:85
// index:0
// void QGraphicsSceneMouseEvent(enum QEvent::Type)
func NewQGraphicsSceneMouseEvent(type_ int) *QGraphicsSceneMouseEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneMouseEventFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneMouseEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMouseEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneMouseEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:86
// index:0
// virtual
// void ~QGraphicsSceneMouseEvent()
func DeleteQGraphicsSceneMouseEvent(*QGraphicsSceneMouseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:88
// index:0
// QPointF pos()
func (this *QGraphicsSceneMouseEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:89
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:91
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneMouseEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:92
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:94
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneMouseEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:95
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneMouseEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:97
// index:0
// QPointF buttonDownPos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownPos(button int) {
	// 0: (, button Qt::MouseButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent13buttonDownPosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:98
// index:0
// void setButtonDownPos(Qt::MouseButton, const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownPos(button int, pos unsafe.Pointer) {
	// 0: (, button Qt::MouseButton, pos const QPointF &), (&button, pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent16setButtonDownPosEN2Qt11MouseButtonERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:100
// index:0
// QPointF buttonDownScenePos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownScenePos(button int) {
	// 0: (, button Qt::MouseButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent18buttonDownScenePosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:101
// index:0
// void setButtonDownScenePos(Qt::MouseButton, const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownScenePos(button int, pos unsafe.Pointer) {
	// 0: (, button Qt::MouseButton, pos const QPointF &), (&button, pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent21setButtonDownScenePosEN2Qt11MouseButtonERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:103
// index:0
// QPoint buttonDownScreenPos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button int) {
	// 0: (, button Qt::MouseButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent19buttonDownScreenPosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:104
// index:0
// void setButtonDownScreenPos(Qt::MouseButton, const class QPoint &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownScreenPos(button int, pos unsafe.Pointer) {
	// 0: (, button Qt::MouseButton, pos const QPoint &), (&button, pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent22setButtonDownScreenPosEN2Qt11MouseButtonERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:106
// index:0
// QPointF lastPos()
func (this *QGraphicsSceneMouseEvent) LastPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent7lastPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:107
// index:0
// void setLastPos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetLastPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:109
// index:0
// QPointF lastScenePos()
func (this *QGraphicsSceneMouseEvent) LastScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent12lastScenePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:110
// index:0
// void setLastScenePos(const class QPointF &)
func (this *QGraphicsSceneMouseEvent) SetLastScenePos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:112
// index:0
// QPoint lastScreenPos()
func (this *QGraphicsSceneMouseEvent) LastScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:113
// index:0
// void setLastScreenPos(const class QPoint &)
func (this *QGraphicsSceneMouseEvent) SetLastScreenPos(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:115
// index:0
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneMouseEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:116
// index:0
// void setButtons(Qt::MouseButtons)
func (this *QGraphicsSceneMouseEvent) SetButtons(buttons int) {
	// 0: (, QFlags<Qt::MouseButton> buttons), (&buttons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:118
// index:0
// Qt::MouseButton button()
func (this *QGraphicsSceneMouseEvent) Button() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent6buttonEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:119
// index:0
// void setButton(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) SetButton(button int) {
	// 0: (, button Qt::MouseButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent9setButtonEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:121
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneMouseEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:122
// index:0
// void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneMouseEvent) SetModifiers(modifiers int) {
	// 0: (, QFlags<Qt::KeyboardModifier> modifiers), (&modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:124
// index:0
// Qt::MouseEventSource source()
func (this *QGraphicsSceneMouseEvent) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:125
// index:0
// void setSource(Qt::MouseEventSource)
func (this *QGraphicsSceneMouseEvent) SetSource(source int) {
	// 0: (, source Qt::MouseEventSource), (&source)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent9setSourceEN2Qt16MouseEventSourceE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:127
// index:0
// Qt::MouseEventFlags flags()
func (this *QGraphicsSceneMouseEvent) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent5flagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:128
// index:0
// void setFlags(Qt::MouseEventFlags)
func (this *QGraphicsSceneMouseEvent) SetFlags(arg0 int) {
	// 0: (, Qt::MouseEventFlags arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent8setFlagsE6QFlagsIN2Qt14MouseEventFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
