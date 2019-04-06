//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtabwidget.h
// #include <qtabwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
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

// /usr/include/qt/QtWidgets/qtabwidget.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabWhatsThis(int, const QString &)

/*
Sets the What's This help text for the page at position index to text.

This function was introduced in  Qt 4.1.

See also tabWhatsThis().
*/
func (this *QTabWidget) SetTabWhatsThis(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setTabWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int) const

/*
Returns the What's This help text for the page at position index, or an empty string if no help text has been set.

This function was introduced in  Qt 4.1.

See also setTabWhatsThis().
*/
func (this *QTabWidget) TabWhatsThis(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12tabWhatsThisEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end

//  keep block begin

func init_unused_10998() {
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
