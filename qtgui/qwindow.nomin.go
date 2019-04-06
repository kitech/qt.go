//  header block begin

// +build !minimal

package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtGui/qwindow.h:358
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Override this to handle mouse wheel or other wheel events (ev).
*/
func (this *QWindow) WheelEvent(arg0 QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:362
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabletEvent(QTabletEvent *)

/*
Override this to handle tablet press, move, and release events (ev).

Proximity enter and leave events are not sent to windows, they are delivered to the application instance.
*/
func (this *QWindow) TabletEvent(arg0 QTabletEvent_ITF /*777 QTabletEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTabletEvent_PTR() != nil {
		convArg0 = arg0.QTabletEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11tabletEventEP12QTabletEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10852() {
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
}

//  keep block end
