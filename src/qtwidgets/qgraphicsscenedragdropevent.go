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
type QGraphicsSceneDragDropEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneDragDropEvent) GetCthis() unsafe.Pointer {
	return this.QGraphicsSceneEvent.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:252
// index:0
// void QGraphicsSceneDragDropEvent(enum QEvent::Type)
func NewQGraphicsSceneDragDropEvent(type_ int) *QGraphicsSceneDragDropEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneDragDropEventFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneDragDropEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneDragDropEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneDragDropEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:253
// index:0
// virtual
// void ~QGraphicsSceneDragDropEvent()
func DeleteQGraphicsSceneDragDropEvent(*QGraphicsSceneDragDropEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:255
// index:0
// QPointF pos()
func (this *QGraphicsSceneDragDropEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:256
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:258
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneDragDropEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:259
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:261
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneDragDropEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:262
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneDragDropEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:264
// index:0
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneDragDropEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:265
// index:0
// void setButtons(Qt::MouseButtons)
func (this *QGraphicsSceneDragDropEvent) SetButtons(buttons int) {
	// 0: (, QFlags<Qt::MouseButton> buttons), (&buttons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:267
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneDragDropEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:268
// index:0
// void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneDragDropEvent) SetModifiers(modifiers int) {
	// 0: (, QFlags<Qt::KeyboardModifier> modifiers), (&modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:270
// index:0
// Qt::DropActions possibleActions()
func (this *QGraphicsSceneDragDropEvent) PossibleActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:271
// index:0
// void setPossibleActions(Qt::DropActions)
func (this *QGraphicsSceneDragDropEvent) SetPossibleActions(actions int) {
	// 0: (, QFlags<Qt::DropAction> actions), (&actions)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent18setPossibleActionsE6QFlagsIN2Qt10DropActionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &actions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:273
// index:0
// Qt::DropAction proposedAction()
func (this *QGraphicsSceneDragDropEvent) ProposedAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent14proposedActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:274
// index:0
// void setProposedAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetProposedAction(action int) {
	// 0: (, action Qt::DropAction), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent17setProposedActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:275
// index:0
// void acceptProposedAction()
func (this *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:277
// index:0
// Qt::DropAction dropAction()
func (this *QGraphicsSceneDragDropEvent) DropAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent10dropActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:278
// index:0
// void setDropAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetDropAction(action int) {
	// 0: (, action Qt::DropAction), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent13setDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:280
// index:0
// QWidget * source()
func (this *QGraphicsSceneDragDropEvent) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:281
// index:0
// void setSource(class QWidget *)
func (this *QGraphicsSceneDragDropEvent) SetSource(source unsafe.Pointer) {
	// 0: (, source QWidget *), (source)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:283
// index:0
// const QMimeData * mimeData()
func (this *QGraphicsSceneDragDropEvent) MimeData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8mimeDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:284
// index:0
// void setMimeData(const class QMimeData *)
func (this *QGraphicsSceneDragDropEvent) SetMimeData(data unsafe.Pointer) {
	// 0: (, data const QMimeData *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

//  body block end
