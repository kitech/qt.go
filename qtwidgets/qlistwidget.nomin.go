//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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

// /usr/include/qt/QtWidgets/qlistwidget.h:260
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QListView::dropEvent().
*/
func (this *QListWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:290
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(int, const QMimeData *, Qt::DropAction)

/*
Handles data supplied by an external drag and drop operation that ended with the given action in the given index. Returns true if data and action can be handled by the model; otherwise returns false.

See also supportedDropActions().
*/
func (this *QListWidget) DropMimeData(index int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg1 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, action)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:291
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Returns the drop actions supported by this view.

See also Qt::DropActions.
*/
func (this *QListWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QListWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_11272() {
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
