package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QGraphicsSceneDragDropEvent struct {
	*QGraphicsSceneEvent
}
type QGraphicsSceneDragDropEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneDragDropEvent_PTR() *QGraphicsSceneDragDropEvent
}

func (ptr *QGraphicsSceneDragDropEvent) QGraphicsSceneDragDropEvent_PTR() *QGraphicsSceneDragDropEvent {
	return ptr
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
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneDragDropEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneDragDropEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneDragDropEvent(enum QEvent::Type)
func NewQGraphicsSceneDragDropEvent__() *QGraphicsSceneDragDropEvent {
	// arg: 0, QEvent::Type=Enum, QEvent::Type=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneDragDropEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneDragDropEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:253
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneDragDropEvent()
func DeleteQGraphicsSceneDragDropEvent(this *QGraphicsSceneDragDropEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:255
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos() const
func (this *QGraphicsSceneDragDropEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent6setPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:258
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos() const
func (this *QGraphicsSceneDragDropEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8scenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)
func (this *QGraphicsSceneDragDropEvent) SetScenePos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setScenePosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:261
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos() const
func (this *QGraphicsSceneDragDropEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)
func (this *QGraphicsSceneDragDropEvent) SetScreenPos(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setScreenPosERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:264
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons() const
func (this *QGraphicsSceneDragDropEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtons(Qt::MouseButtons)
func (this *QGraphicsSceneDragDropEvent) SetButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers() const
func (this *QGraphicsSceneDragDropEvent) Modifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent9modifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneDragDropEvent) SetModifiers(modifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:270
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropActions possibleActions() const
func (this *QGraphicsSceneDragDropEvent) PossibleActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent15possibleActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPossibleActions(Qt::DropActions)
func (this *QGraphicsSceneDragDropEvent) SetPossibleActions(actions int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent18setPossibleActionsE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), actions)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:273
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction proposedAction() const
func (this *QGraphicsSceneDragDropEvent) ProposedAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent14proposedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProposedAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetProposedAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent17setProposedActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acceptProposedAction()
func (this *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent20acceptProposedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:277
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction dropAction() const
func (this *QGraphicsSceneDragDropEvent) DropAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent10dropActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropAction(Qt::DropAction)
func (this *QGraphicsSceneDragDropEvent) SetDropAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent13setDropActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:280
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * source() const
func (this *QGraphicsSceneDragDropEvent) Source() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSource(QWidget *)
func (this *QGraphicsSceneDragDropEvent) SetSource(source QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if source != nil && source.QWidget_PTR() != nil {
		convArg0 = source.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent9setSourceEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:283
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMimeData * mimeData() const
func (this *QGraphicsSceneDragDropEvent) MimeData() *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QGraphicsSceneDragDropEvent8mimeDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeData(const QMimeData *)
func (this *QGraphicsSceneDragDropEvent) SetMimeData(data qtcore.QMimeData_ITF /*777 const QMimeData **/) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QGraphicsSceneDragDropEvent11setMimeDataEPK9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
