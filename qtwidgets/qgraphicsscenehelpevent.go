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
// extern C begin: 16
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
type QGraphicsSceneHelpEvent struct {
	*QGraphicsSceneEvent
}
type QGraphicsSceneHelpEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneHelpEvent_PTR() *QGraphicsSceneHelpEvent
}

func (ptr *QGraphicsSceneHelpEvent) QGraphicsSceneHelpEvent_PTR() *QGraphicsSceneHelpEvent { return ptr }

func (this *QGraphicsSceneHelpEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneHelpEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneHelpEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHelpEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneHelpEvent{bcthis0}
}
func (*QGraphicsSceneHelpEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHelpEvent {
	return NewQGraphicsSceneHelpEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneHelpEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneHelpEvent) NewForInherit(type_ int) *QGraphicsSceneHelpEvent {
	return NewQGraphicsSceneHelpEvent(type_)
}
func NewQGraphicsSceneHelpEvent(type_ int) *QGraphicsSceneHelpEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHelpEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneHelpEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneHelpEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneHelpEvent) NewForInheritp() *QGraphicsSceneHelpEvent {
	return NewQGraphicsSceneHelpEventp()
}
func NewQGraphicsSceneHelpEventp() *QGraphicsSceneHelpEvent {
	// arg: 0, QEvent::Type=Enum, QEvent::Type=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHelpEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneHelpEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneHelpEvent()

/*

 */
func DeleteQGraphicsSceneHelpEvent(this *QGraphicsSceneHelpEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:238
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos() const

/*

 */
func (this *QGraphicsSceneHelpEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSceneHelpEvent8scenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)

/*

 */
func (this *QGraphicsSceneHelpEvent) SetScenePos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:241
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos() const

/*

 */
func (this *QGraphicsSceneHelpEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSceneHelpEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)

/*

 */
func (this *QGraphicsSceneHelpEvent) SetScreenPos(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
