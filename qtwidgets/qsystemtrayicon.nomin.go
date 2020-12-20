//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qsystemtrayicon.h
// #include <qsystemtrayicon.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:82
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setContextMenu(QMenu *)

/*
 */
func (this *QSystemTrayIcon) SetContextMenu(menu QMenu_ITF /*777 QMenu **/) {
	var convArg0 Voidptr
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(2313850013, "_ZN15QSystemTrayIcon14setContextMenuEP5QMenu", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:83
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QMenu * contextMenu() const

/*
 */
func (this *QSystemTrayIcon) ContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.Qtcc1(1821247260, "_ZNK15QSystemTrayIcon11contextMenuEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ QMenuFromptr(Voidptr(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10158() {
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
