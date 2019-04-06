//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
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

// /usr/include/qt/QtWidgets/qmdisubwindow.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSystemMenu(QMenu *)

/*
Sets systemMenu as the current system menu for this subwindow.

By default, each QMdiSubWindow has a standard system menu.

QActions for the system menu created by QMdiSubWindow will automatically be updated depending on the current window state; e.g., the minimize action will be disabled after the window is minimized.

QActions added by the user are not updated by QMdiSubWindow.

QMdiSubWindow takes ownership of systemMenu; you do not have to delete it. Any existing menus will be deleted.

See also systemMenu() and showSystemMenu().
*/
func (this *QMdiSubWindow) SetSystemMenu(systemMenu QMenu_ITF /*777 QMenu **/) {
	var convArg0 unsafe.Pointer
	if systemMenu != nil && systemMenu.QMenu_PTR() != nil {
		convArg0 = systemMenu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13setSystemMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * systemMenu() const

/*
Returns a pointer to the current system menu, or zero if no system menu is set. QMdiSubWindow provides a default system menu, but you can also set the menu with setSystemMenu().

See also setSystemMenu() and showSystemMenu().
*/
func (this *QMdiSubWindow) SystemMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10systemMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showSystemMenu()

/*
Shows the system menu below the system menu icon in the title bar.

See also setSystemMenu() and systemMenu().
*/
func (this *QMdiSubWindow) ShowSystemMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow14showSystemMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11278() {
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
