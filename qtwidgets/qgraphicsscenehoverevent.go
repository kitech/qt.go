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
// extern C begin: 12
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

type QGraphicsSceneHoverEvent struct {
	*QGraphicsSceneEvent
}
type QGraphicsSceneHoverEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneHoverEvent_PTR() *QGraphicsSceneHoverEvent
}

func (ptr *QGraphicsSceneHoverEvent) QGraphicsSceneHoverEvent_PTR() *QGraphicsSceneHoverEvent {
	return ptr
}

func (this *QGraphicsSceneHoverEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneHoverEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneHoverEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHoverEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneHoverEvent{bcthis0}
}
func (*QGraphicsSceneHoverEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHoverEvent {
	return NewQGraphicsSceneHoverEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneHoverEvent(enum QEvent::Type)
func NewQGraphicsSceneHoverEvent(type_ int) *QGraphicsSceneHoverEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHoverEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneHoverEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneHoverEvent(enum QEvent::Type)
func NewQGraphicsSceneHoverEvent__() *QGraphicsSceneHoverEvent {
	// arg: 0, QEvent::Type=Enum, QEvent::Type=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHoverEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneHoverEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:202
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneHoverEvent()
func DeleteQGraphicsSceneHoverEvent(this *QGraphicsSceneHoverEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:204
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos() const
func (this *QGraphicsSceneHoverEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:207
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos() const
func (this *QGraphicsSceneHoverEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent8scenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetScenePos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:210
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos() const
func (this *QGraphicsSceneHoverEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)
func (this *QGraphicsSceneHoverEvent) SetScreenPos(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:213
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastPos() const
func (this *QGraphicsSceneHoverEvent) LastPos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent7lastPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastPos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetLastPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:216
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastScenePos() const
func (this *QGraphicsSceneHoverEvent) LastScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent12lastScenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScenePos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetLastScenePos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint lastScreenPos() const
func (this *QGraphicsSceneHoverEvent) LastScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScreenPos(const QPoint &)
func (this *QGraphicsSceneHoverEvent) SetLastScreenPos(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:222
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers() const
func (this *QGraphicsSceneHoverEvent) Modifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent9modifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneHoverEvent) SetModifiers(modifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
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
