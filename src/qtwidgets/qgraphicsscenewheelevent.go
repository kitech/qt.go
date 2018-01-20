//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
	return this.QGraphicsSceneEvent.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:139
// index:0
// void QGraphicsSceneWheelEvent(enum QEvent::Type)
func NewQGraphicsSceneWheelEvent(type_ int) *QGraphicsSceneWheelEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneWheelEventFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneWheelEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneWheelEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneWheelEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:140
// index:0
// virtual
// void ~QGraphicsSceneWheelEvent()
func DeleteQGraphicsSceneWheelEvent(*QGraphicsSceneWheelEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:142
// index:0
// QPointF pos()
func (this *QGraphicsSceneWheelEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:143
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsSceneWheelEvent) SetPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:145
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneWheelEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:146
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneWheelEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:148
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneWheelEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:149
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneWheelEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:151
// index:0
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneWheelEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:152
// index:0
// void setButtons(Qt::MouseButtons)
func (this *QGraphicsSceneWheelEvent) SetButtons(buttons int) {
	// 0: (, QFlags<Qt::MouseButton> buttons), (&buttons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:154
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneWheelEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:155
// index:0
// void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneWheelEvent) SetModifiers(modifiers int) {
	// 0: (, QFlags<Qt::KeyboardModifier> modifiers), (&modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:157
// index:0
// int delta()
func (this *QGraphicsSceneWheelEvent) Delta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent5deltaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:158
// index:0
// void setDelta(int)
func (this *QGraphicsSceneWheelEvent) SetDelta(delta int) {
	// 0: (, delta int), (&delta)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent8setDeltaEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &delta)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:160
// index:0
// Qt::Orientation orientation()
func (this *QGraphicsSceneWheelEvent) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:161
// index:0
// void setOrientation(Qt::Orientation)
func (this *QGraphicsSceneWheelEvent) SetOrientation(orientation int) {
	// 0: (, orientation Qt::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

//  body block end
