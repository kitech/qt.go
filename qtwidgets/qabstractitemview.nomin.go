//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractitemview.h
// #include <qabstractitemview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
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

// /usr/include/qt/QtWidgets/qabstractitemview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropIndicatorShown(bool)

/*

 */
func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setDropIndicatorShownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showDropIndicator() const

/*

 */
func (this *QAbstractItemView) ShowDropIndicator() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17showDropIndicatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(bool)

/*

 */
func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setDragEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragEnabled() const

/*

 */
func (this *QAbstractItemView) DragEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView11dragEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragDropOverwriteMode(bool)

/*

 */
func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView24setDragDropOverwriteModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overwrite)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragDropOverwriteMode() const

/*

 */
func (this *QAbstractItemView) DragDropOverwriteMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21dragDropOverwriteModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragDropMode(QAbstractItemView::DragDropMode)

/*

 */
func (this *QAbstractItemView) SetDragDropMode(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setDragDropModeENS_12DragDropModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:188
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::DragDropMode dragDropMode() const

/*

 */
func (this *QAbstractItemView) DragDropMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12dragDropModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultDropAction(Qt::DropAction)

/*

 */
func (this *QAbstractItemView) SetDefaultDropAction(dropAction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20setDefaultDropActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dropAction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction defaultDropAction() const

/*

 */
func (this *QAbstractItemView) DefaultDropAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17defaultDropActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:301
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void startDrag(Qt::DropActions)

/*
Starts a drag by calling drag->exec() using the given supportedActions.
*/
func (this *QAbstractItemView) StartDrag(supportedActions int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9startDragE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:338
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QAbstractScrollArea::dragEnterEvent().

This function is called with the given event when a drag and drop operation enters the widget. If the drag is over a valid dropping place (e.g. over an item that accepts drops), the event is accepted; otherwise it is ignored.

See also dropEvent() and startDrag().
*/
func (this *QAbstractItemView) DragEnterEvent(event qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragEnterEvent_PTR() != nil {
		convArg0 = event.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:339
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QAbstractScrollArea::dragMoveEvent().

This function is called continuously with the given event during a drag and drop operation over the widget. It can cause the view to scroll if, for example, the user drags a selection to view's right or bottom edge. In this case, the event will be accepted; otherwise it will be ignored.

See also dropEvent() and startDrag().
*/
func (this *QAbstractItemView) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragMoveEvent_PTR() != nil {
		convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:340
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QAbstractScrollArea::dragLeaveEvent().

This function is called when the item being dragged leaves the view. The event describes the state of the drag and drop operation.
*/
func (this *QAbstractItemView) DragLeaveEvent(event qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragLeaveEvent_PTR() != nil {
		convArg0 = event.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:341
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QAbstractScrollArea::dropEvent().

This function is called with the given event when a drop event occurs over the widget. If the model accepts the even position the drop event is accepted; otherwise it is ignored.

See also startDrag().
*/
func (this *QAbstractItemView) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:353
// index:0
// Protected Visibility=Default Availability=Available
// [4] QAbstractItemView::DropIndicatorPosition dropIndicatorPosition() const

/*
Returns the position of the drop indicator in relation to the closest item.

This function was introduced in  Qt 4.1.
*/
func (this *QAbstractItemView) DropIndicatorPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21dropIndicatorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_11062() {
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
