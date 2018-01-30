package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 96
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
type QGraphicsSceneEvent struct {
	*qtcore.QEvent
}

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
// [-2] void QGraphicsSceneEvent(enum QEvent::Type)
func NewQGraphicsSceneEvent(type_ int) *QGraphicsSceneEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneEvent()
func DeleteQGraphicsSceneEvent(*QGraphicsSceneEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QGraphicsSceneEvent) Widget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsSceneEvent6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)
func (this *QGraphicsSceneEvent) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEvent9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
