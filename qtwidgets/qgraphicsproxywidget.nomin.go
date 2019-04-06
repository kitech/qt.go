//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h
// #include <qgraphicsproxywidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 82
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

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dragEnterEvent().
*/
func (this *QGraphicsProxyWidget) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dragLeaveEvent().
*/
func (this *QGraphicsProxyWidget) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dragMoveEvent().
*/
func (this *QGraphicsProxyWidget) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dropEvent().
*/
func (this *QGraphicsProxyWidget) DropEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QGraphicsSceneWheelEvent *)

/*
Reimplemented from QGraphicsItem::wheelEvent().
*/
func (this *QGraphicsProxyWidget) WheelEvent(event QGraphicsSceneWheelEvent_ITF /*777 QGraphicsSceneWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneWheelEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11214() {
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
