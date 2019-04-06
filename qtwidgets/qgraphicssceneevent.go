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
// extern C begin: 108
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
type QGraphicsSceneEvent struct {
	*qtcore.QEvent
}
type QGraphicsSceneEvent_ITF interface {
	qtcore.QEvent_ITF
	QGraphicsSceneEvent_PTR() *QGraphicsSceneEvent
}

func (ptr *QGraphicsSceneEvent) QGraphicsSceneEvent_PTR() *QGraphicsSceneEvent { return ptr }

func (this *QGraphicsSceneEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QGraphicsSceneEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQGraphicsSceneEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QGraphicsSceneEvent{bcthis0}
}
func (*QGraphicsSceneEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneEvent {
	return NewQGraphicsSceneEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneEvent) NewForInherit(type_ int) *QGraphicsSceneEvent {
	return NewQGraphicsSceneEvent(type_)
}
func NewQGraphicsSceneEvent(type_ int) *QGraphicsSceneEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsSceneEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneEvent()

/*

 */
func DeleteQGraphicsSceneEvent(this *QGraphicsSceneEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsSceneEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the widget where the event originated, or 0 if the event originates from another application.
*/
func (this *QGraphicsSceneEvent) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsSceneEvent6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*

 */
func (this *QGraphicsSceneEvent) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsSceneEvent9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11217() {
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
