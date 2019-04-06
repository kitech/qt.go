//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtabbar.h
// #include <qtabbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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

// /usr/include/qt/QtWidgets/qtabbar.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabWhatsThis(int, const QString &)

/*
Sets the What's This help text of the tab at position index to text.

This function was introduced in  Qt 4.1.

See also tabWhatsThis().
*/
func (this *QTabBar) SetTabWhatsThis(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15setTabWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int) const

/*
Returns the What's This help text of the tab at position index, or an empty string if index is out of range.

This function was introduced in  Qt 4.1.

See also setTabWhatsThis().
*/
func (this *QTabBar) TabWhatsThis(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12tabWhatsThisEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabbar.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QTabBar) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10996() {
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
