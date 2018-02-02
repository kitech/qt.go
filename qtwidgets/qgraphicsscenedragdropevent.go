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

type QGraphicsSceneDragDropEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneDragDropEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneDragDropEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneDragDropEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneDragDropEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneDragDropEvent{bcthis0}
}
func (*QGraphicsSceneDragDropEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneDragDropEvent {
	return NewQGraphicsSceneDragDropEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneDragDropEvent(enum QEvent::Type)
func NewQGraphicsSceneDragDropEvent(type_ int) *QGraphicsSceneDragDropEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneDragDropEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneDragDropEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:253
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneDragDropEvent()
func DeleteQGraphicsSceneDragDropEvent(this *QGraphicsSceneDragDropEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:255
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos()
func (this *QGraphicsSceneDragDropEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:258
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos()
func (this *QGraphicsSceneDragDropEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:261
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos()
func (this *QGraphicsSceneDragDropEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)
func (this *QGraphicsSceneDragDropEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:264
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons()
func (this *QGraphicsSceneDragDropEvent) Buttons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtons(Qt::MouseButtons)
func (this *QGraphicsSceneDragDropEvent) SetButtons(buttons int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneDragDropEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneDragDropEvent) SetModifiers(modifiers int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:270
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropActions possibleActions()
func (this *QGraphicsSceneDragDropEvent) PossibleActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPossibleActions(Qt::DropActions)
func (this *QGraphicsSceneDragDropEvent) SetPossibleActions(actions int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent18setPossibleActionsE6QFlagsIN2Qt10DropActionEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), actions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:273
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction proposedAction()
func (this *QGraphicsSceneDragDropEvent) ProposedAction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent14proposedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProposedAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetProposedAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent17setProposedActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acceptProposedAction()
func (this *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:277
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction dropAction()
func (this *QGraphicsSceneDragDropEvent) DropAction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent10dropActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetDropAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent13setDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:280
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * source()
func (this *QGraphicsSceneDragDropEvent) Source() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSource(QWidget *)
func (this *QGraphicsSceneDragDropEvent) SetSource(source *QWidget /*777 QWidget **/) {
	var convArg0 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:283
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMimeData * mimeData()
func (this *QGraphicsSceneDragDropEvent) MimeData() *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8mimeDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeData(const QMimeData *)
func (this *QGraphicsSceneDragDropEvent) SetMimeData(data *qtcore.QMimeData /*777 const QMimeData **/) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
