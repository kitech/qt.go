//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qaction.h
// #include <qaction.h>
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

// /usr/include/qt/QtWidgets/qaction.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu() const

/*
Returns the menu contained by this action. Actions that contain menus can be used to create menu items with submenus, or inserted into toolbars to create buttons with popup menus.

See also setMenu() and QMenu::addAction().
*/
func (this *QAction) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)

/*
Sets the menu contained by this action to the specified menu.

See also menu().
*/
func (this *QAction) SetMenu(menu QMenu_ITF /*777 QMenu **/) {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11066() {
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
