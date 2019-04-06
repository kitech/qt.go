//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbutton.h
// #include <qtoolbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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

// /usr/include/qt/QtWidgets/qtoolbutton.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)

/*
Associates the given menu with this tool button.

The menu will be shown according to the button's popupMode.

Ownership of the menu is not transferred to the tool button.

See also menu().
*/
func (this *QToolButton) SetMenu(menu QMenu_ITF /*777 QMenu **/) {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu() const

/*
Returns the associated menu, or 0 if no menu has been defined.

See also setMenu().
*/
func (this *QToolButton) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPopupMode(QToolButton::ToolButtonPopupMode)

/*

 */
func (this *QToolButton) SetPopupMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton12setPopupModeENS_19ToolButtonPopupModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QToolButton::ToolButtonPopupMode popupMode() const

/*

 */
func (this *QToolButton) PopupMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton9popupModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMenu()

/*
Shows (pops up) the associated popup menu. If there is no such menu, this function does nothing. This function does not return until the popup menu has been closed by the user.
*/
func (this *QToolButton) ShowMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton8showMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11352() {
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
