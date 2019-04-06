//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtableview.h
// #include <qtableview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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

// /usr/include/qt/QtWidgets/qtableview.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerButtonEnabled(bool)

/*

 */
func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTableView22setCornerButtonEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtableview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCornerButtonEnabled() const

/*

 */
func (this *QTableView) IsCornerButtonEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTableView21isCornerButtonEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_11338() {
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
