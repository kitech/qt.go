//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qlistview.h
// #include <qlistview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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

// /usr/include/qt/QtWidgets/qlistview.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QAbstractScrollArea::wheelEvent().
*/
func (this *QListView) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QAbstractItemView::dragMoveEvent().
*/
func (this *QListView) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:167
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QAbstractItemView::dragLeaveEvent().
*/
func (this *QListView) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:168
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QAbstractItemView::dropEvent().
*/
func (this *QListView) DropEvent(e qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDropEvent_PTR() != nil {
		convArg0 = e.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void startDrag(Qt::DropActions)

/*
Reimplemented from QAbstractItemView::startDrag().
*/
func (this *QListView) StartDrag(supportedActions int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView9startDragE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11268() {
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
