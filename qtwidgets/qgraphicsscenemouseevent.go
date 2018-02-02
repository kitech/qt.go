package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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

type QGraphicsSceneMouseEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneMouseEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneMouseEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneMouseEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMouseEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneMouseEvent{bcthis0}
}
func (*QGraphicsSceneMouseEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMouseEvent {
	return NewQGraphicsSceneMouseEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneMouseEvent(enum QEvent::Type)
func NewQGraphicsSceneMouseEvent(type_ int) *QGraphicsSceneMouseEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneMouseEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneMouseEvent()
func DeleteQGraphicsSceneMouseEvent(this *QGraphicsSceneMouseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:88
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos()
func (this *QGraphicsSceneMouseEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)
func (this *QGraphicsSceneMouseEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:91
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos()
func (this *QGraphicsSceneMouseEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)
func (this *QGraphicsSceneMouseEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos()
func (this *QGraphicsSceneMouseEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)
func (this *QGraphicsSceneMouseEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:97
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF buttonDownPos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownPos(button int) *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent13buttonDownPosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonDownPos(Qt::MouseButton, const QPointF &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownPos(button int, pos *qtcore.QPointF) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent16setButtonDownPosEN2Qt11MouseButtonERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:100
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF buttonDownScenePos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownScenePos(button int) *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent18buttonDownScenePosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonDownScenePos(Qt::MouseButton, const QPointF &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownScenePos(button int, pos *qtcore.QPointF) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent21setButtonDownScenePosEN2Qt11MouseButtonERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint buttonDownScreenPos(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button int) *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent19buttonDownScreenPosEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonDownScreenPos(Qt::MouseButton, const QPoint &)
func (this *QGraphicsSceneMouseEvent) SetButtonDownScreenPos(button int, pos *qtcore.QPoint) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent22setButtonDownScreenPosEN2Qt11MouseButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:106
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastPos()
func (this *QGraphicsSceneMouseEvent) LastPos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent7lastPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastPos(const QPointF &)
func (this *QGraphicsSceneMouseEvent) SetLastPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent10setLastPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:109
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastScenePos()
func (this *QGraphicsSceneMouseEvent) LastScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent12lastScenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScenePos(const QPointF &)
func (this *QGraphicsSceneMouseEvent) SetLastScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent15setLastScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint lastScreenPos()
func (this *QGraphicsSceneMouseEvent) LastScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent13lastScreenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScreenPos(const QPoint &)
func (this *QGraphicsSceneMouseEvent) SetLastScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent16setLastScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons()
func (this *QGraphicsSceneMouseEvent) Buttons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtons(Qt::MouseButtons)
func (this *QGraphicsSceneMouseEvent) SetButtons(buttons int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButton button()
func (this *QGraphicsSceneMouseEvent) Button() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButton(Qt::MouseButton)
func (this *QGraphicsSceneMouseEvent) SetButton(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent9setButtonEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneMouseEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneMouseEvent) SetModifiers(modifiers int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseEventSource source()
func (this *QGraphicsSceneMouseEvent) Source() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSource(Qt::MouseEventSource)
func (this *QGraphicsSceneMouseEvent) SetSource(source int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent9setSourceEN2Qt16MouseEventSourceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseEventFlags flags()
func (this *QGraphicsSceneMouseEvent) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneMouseEvent5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::MouseEventFlags)
func (this *QGraphicsSceneMouseEvent) SetFlags(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneMouseEvent8setFlagsE6QFlagsIN2Qt14MouseEventFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
