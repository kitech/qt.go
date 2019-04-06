//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractspinbox.h
// #include <qabstractspinbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
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

// /usr/include/qt/QtWidgets/qabstractspinbox.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QAbstractSpinBox) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QAbstractSpinBox) ContextMenuEvent(event qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QContextMenuEvent_PTR() != nil {
		convArg0 = event.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10988() {
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
