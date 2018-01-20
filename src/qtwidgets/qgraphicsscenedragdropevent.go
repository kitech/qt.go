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
func NewQGraphicsSceneDragDropEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneDragDropEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneDragDropEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:252
// index:0
// Public
// void QGraphicsSceneDragDropEvent(enum QEvent::Type)
func NewQGraphicsSceneDragDropEvent(type_ int) *QGraphicsSceneDragDropEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneDragDropEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:253
// index:0
// Public virtual
// void ~QGraphicsSceneDragDropEvent()
func DeleteQGraphicsSceneDragDropEvent(*QGraphicsSceneDragDropEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:255
// index:0
// Public
// QPointF pos()
func (this *QGraphicsSceneDragDropEvent) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:256
// index:0
// Public
// void setPos(const class QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:258
// index:0
// Public
// QPointF scenePos()
func (this *QGraphicsSceneDragDropEvent) ScenePos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:259
// index:0
// Public
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:261
// index:0
// Public
// QPoint screenPos()
func (this *QGraphicsSceneDragDropEvent) ScreenPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:262
// index:0
// Public
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneDragDropEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:264
// index:0
// Public
// Qt::MouseButtons buttons()
func (this *QGraphicsSceneDragDropEvent) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:267
// index:0
// Public
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneDragDropEvent) Modifiers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:270
// index:0
// Public
// Qt::DropActions possibleActions()
func (this *QGraphicsSceneDragDropEvent) PossibleActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:273
// index:0
// Public
// Qt::DropAction proposedAction()
func (this *QGraphicsSceneDragDropEvent) ProposedAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent14proposedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:274
// index:0
// Public
// void setProposedAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetProposedAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent17setProposedActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:275
// index:0
// Public
// void acceptProposedAction()
func (this *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:277
// index:0
// Public
// Qt::DropAction dropAction()
func (this *QGraphicsSceneDragDropEvent) DropAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent10dropActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:278
// index:0
// Public
// void setDropAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetDropAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent13setDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:280
// index:0
// Public
// QWidget * source()
func (this *QGraphicsSceneDragDropEvent) Source() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:281
// index:0
// Public
// void setSource(class QWidget *)
func (this *QGraphicsSceneDragDropEvent) SetSource(source unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:283
// index:0
// Public
// const QMimeData * mimeData()
func (this *QGraphicsSceneDragDropEvent) MimeData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8mimeDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:284
// index:0
// Public
// void setMimeData(const class QMimeData *)
func (this *QGraphicsSceneDragDropEvent) SetMimeData(data unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

//  body block end
