// +build !minimal

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
// extern C begin: 30
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

/*

 */
type QGraphicsSceneWheelEvent struct {
	*QGraphicsSceneEvent
}
type QGraphicsSceneWheelEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneWheelEvent_PTR() *QGraphicsSceneWheelEvent
}

func (ptr *QGraphicsSceneWheelEvent) QGraphicsSceneWheelEvent_PTR() *QGraphicsSceneWheelEvent {
	return ptr
}

func (this *QGraphicsSceneWheelEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneWheelEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneWheelEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneWheelEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneWheelEvent{bcthis0}
}
func (*QGraphicsSceneWheelEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneWheelEvent {
	return NewQGraphicsSceneWheelEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneWheelEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneWheelEvent) NewForInherit(type_ int) *QGraphicsSceneWheelEvent {
	return NewQGraphicsSceneWheelEvent(type_)
}
func NewQGraphicsSceneWheelEvent(type_ int) *QGraphicsSceneWheelEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneWheelEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneWheelEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneWheelEvent) NewForInheritp() *QGraphicsSceneWheelEvent {
	return NewQGraphicsSceneWheelEventp()
}
func NewQGraphicsSceneWheelEventp() *QGraphicsSceneWheelEvent {
	// arg: 0, QEvent::Type=Enum, QEvent::Type=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneWheelEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneWheelEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneWheelEvent()

/*

 */
func DeleteQGraphicsSceneWheelEvent(this *QGraphicsSceneWheelEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:142
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos() const

/*

 */
func (this *QGraphicsSceneWheelEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent6setPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:145
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos() const

/*

 */
func (this *QGraphicsSceneWheelEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent8scenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetScenePos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent11setScenePosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos() const

/*

 */
func (this *QGraphicsSceneWheelEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetScreenPos(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent12setScreenPosERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:151
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons() const

/*

 */
func (this *QGraphicsSceneWheelEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtons(Qt::MouseButtons)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent10setButtonsE6QFlagsIN2Qt11MouseButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers() const

/*

 */
func (this *QGraphicsSceneWheelEvent) Modifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent9modifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetModifiers(modifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int delta() const

/*

 */
func (this *QGraphicsSceneWheelEvent) Delta() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent5deltaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDelta(int)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetDelta(delta int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent8setDeltaEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), delta)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:160
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QGraphicsSceneWheelEvent) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QGraphicsSceneWheelEvent11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*

 */
func (this *QGraphicsSceneWheelEvent) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QGraphicsSceneWheelEvent14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
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
