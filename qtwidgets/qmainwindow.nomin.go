//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmainwindow.h
// #include <qmainwindow.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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

// /usr/include/qt/QtWidgets/qmainwindow.h:126
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QMenuBar * menuBar() const

/*
 */
func (this *QMainWindow) MenuBar() *QMenuBar /*777 QMenuBar **/ {
	rv, err := qtrt.Qtcc3(3966983917, "_ZNK11QMainWindow7menuBarEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QMenuBarFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:127
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMenuBar(QMenuBar *)

/*
 */
func (this *QMainWindow) SetMenuBar(menubar QMenuBar_ITF /*777 QMenuBar **/) {
	var convArg0 Voidptr
	if menubar != nil && menubar.QMenuBar_PTR() != nil {
		convArg0 = menubar.QMenuBar_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2414944809, "_ZN11QMainWindow10setMenuBarEP8QMenuBar", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:129
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * menuWidget() const

/*
 */
func (this *QMainWindow) MenuWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(350857657, "_ZNK11QMainWindow10menuWidgetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:130
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMenuWidget(QWidget *)

/*
 */
func (this *QMainWindow) SetMenuWidget(menubar QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if menubar != nil && menubar.QWidget_PTR() != nil {
		convArg0 = menubar.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1658462860, "_ZN11QMainWindow13setMenuWidgetEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:134
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QStatusBar * statusBar() const

/*
 */
func (this *QMainWindow) StatusBar() *QStatusBar /*777 QStatusBar **/ {
	rv, err := qtrt.Qtcc3(1994287522, "_ZNK11QMainWindow9statusBarEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QStatusBarFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:135
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStatusBar(QStatusBar *)

/*
 */
func (this *QMainWindow) SetStatusBar(statusbar QStatusBar_ITF /*777 QStatusBar **/) {
	var convArg0 Voidptr
	if statusbar != nil && statusbar.QStatusBar_PTR() != nil {
		convArg0 = statusbar.QStatusBar_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2354640415, "_ZN11QMainWindow12setStatusBarEP10QStatusBar", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:154
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QToolBar * addToolBar(const QString &)

/*
 */
func (this *QMainWindow) AddToolBar(title string) *QToolBar /*777 QToolBar **/ {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(4007182669, "_ZN11QMainWindow10addToolBarERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QToolBarFromptr(Voidptr(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10224() {
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
