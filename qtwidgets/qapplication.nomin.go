//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qapplication.h
// #include <qapplication.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

// /usr/include/qt/QtWidgets/qapplication.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setWheelScrollLines(int)

/*

 */
func (this *QApplication) SetWheelScrollLines(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication19setWheelScrollLinesEi", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QApplication_SetWheelScrollLines(arg0 int) {
	var nilthis *QApplication
	nilthis.SetWheelScrollLines(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:159
// index:0
// Public static Visibility=Default Availability=Available
// [4] int wheelScrollLines()

/*

 */
func (this *QApplication) WheelScrollLines() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication16wheelScrollLinesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_WheelScrollLines() int {
	var nilthis *QApplication
	rv := nilthis.WheelScrollLines()
	return rv
}

//  body block end

//  keep block begin

func init_unused_11070() {
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
