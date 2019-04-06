//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmdiarea.h
// #include <qmdiarea.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
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

// /usr/include/qt/QtWidgets/qmdiarea.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode() const

/*

 */
func (this *QMdiArea) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(bool)

/*

 */
func (this *QMdiArea) SetDocumentMode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(bool)

/*

 */
func (this *QMdiArea) SetTabsClosable(closable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea15setTabsClosableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), closable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable() const

/*

 */
func (this *QMdiArea) TabsClosable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea12tabsClosableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsMovable(bool)

/*

 */
func (this *QMdiArea) SetTabsMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea14setTabsMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsMovable() const

/*

 */
func (this *QMdiArea) TabsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea11tabsMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(QTabWidget::TabShape)

/*

 */
func (this *QMdiArea) SetTabShape(shape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea11setTabShapeEN10QTabWidget8TabShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shape)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape() const

/*

 */
func (this *QMdiArea) TabShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea8tabShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(QTabWidget::TabPosition)

/*

 */
func (this *QMdiArea) SetTabPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea14setTabPositionEN10QTabWidget11TabPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition() const

/*

 */
func (this *QMdiArea) TabPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea11tabPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_11276() {
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
