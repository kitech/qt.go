//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qpushbutton.h
// #include <qpushbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 96
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

// /usr/include/qt/QtWidgets/qpushbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)

/*
Associates the popup menu menu with this push button. This turns the button into a menu button, which in some styles will produce a small triangle to the right of the button's text.

Ownership of the menu is not transferred to the push button.



A push button with popup menus shown in the Fusion widget style.

See also menu().
*/
func (this *QPushButton) SetMenu(menu QMenu_ITF /*777 QMenu **/) {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu() const

/*
Returns the button's associated popup menu or 0 if no popup menu has been set.

See also setMenu().
*/
func (this *QPushButton) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qpushbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMenu()

/*
Shows (pops up) the associated popup menu. If there is no such menu, this function does nothing. This function does not return until the popup menu has been closed by the user.
*/
func (this *QPushButton) ShowMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton8showMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11106() {
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
