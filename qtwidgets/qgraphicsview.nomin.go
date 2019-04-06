//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsview.h
// #include <qgraphicsview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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

// /usr/include/qt/QtWidgets/qgraphicsview.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemSelectionMode rubberBandSelectionMode() const

/*

 */
func (this *QGraphicsView) RubberBandSelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView23rubberBandSelectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRubberBandSelectionMode(Qt::ItemSelectionMode)

/*

 */
func (this *QGraphicsView) SetRubberBandSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView26setRubberBandSelectionModeEN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:149
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rubberBandRect() const

/*
This functions returns the current rubber band area (in viewport coordinates) if the user is currently doing an itemselection with rubber band. When the user is not using the rubber band this functions returns (a null) QRectF().

Notice that part of this QRect can be outise the visual viewport. It can e.g contain negative values.

This function was introduced in  Qt 5.1.

See also rubberBandSelectionMode and rubberBandChanged().
*/
func (this *QGraphicsView) RubberBandRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsView14rubberBandRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rubberBandChanged(QRect, QPointF, QPointF)

/*
This signal is emitted when the rubber band rect is changed. The viewport Rect is specified by rubberBandRect. The drag start position and drag end position are provided in scene points with fromScenePoint and toScenePoint.

When rubberband selection ends this signal will be emitted with null vales.

This function was introduced in  Qt 5.1.

See also rubberBandRect().
*/
func (this *QGraphicsView) RubberBandChanged(viewportRect qtcore.QRect_ITF /*123*/, fromScenePoint qtcore.QPointF_ITF /*123*/, toScenePoint qtcore.QPointF_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if viewportRect != nil && viewportRect.QRect_PTR() != nil {
		convArg0 = viewportRect.QRect_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if fromScenePoint != nil && fromScenePoint.QPointF_PTR() != nil {
		convArg1 = fromScenePoint.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if toScenePoint != nil && toScenePoint.QPointF_PTR() != nil {
		convArg2 = toScenePoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView17rubberBandChangedE5QRect7QPointFS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:247
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QAbstractScrollArea::dragEnterEvent().
*/
func (this *QGraphicsView) DragEnterEvent(event qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragEnterEvent_PTR() != nil {
		convArg0 = event.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QAbstractScrollArea::dragLeaveEvent().
*/
func (this *QGraphicsView) DragLeaveEvent(event qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragLeaveEvent_PTR() != nil {
		convArg0 = event.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:249
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QAbstractScrollArea::dragMoveEvent().
*/
func (this *QGraphicsView) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragMoveEvent_PTR() != nil {
		convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:250
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QAbstractScrollArea::dropEvent().
*/
func (this *QGraphicsView) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:262
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QAbstractScrollArea::wheelEvent().
*/
func (this *QGraphicsView) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsView10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11244() {
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
