// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qstatusbar.h
// #include <qstatusbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

/*
 */
// size 48
type QStatusBar struct {
	*QWidget
}
type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QStatusBarFromptr(cthis unsafe.Pointer) *QStatusBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QStatusBar{bcthis0}
}
func (*QStatusBar) Fromptr(cthis unsafe.Pointer) *QStatusBar {
	return QStatusBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)

/*
 */
func (*QStatusBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QStatusBar {
	return NewQStatusBar(parent)
}
func NewQStatusBar(parent QWidget_ITF /*777 QWidget **/) *QStatusBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2740779997, "_ZN10QStatusBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QStatusBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QStatusBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)

/*
 */
func (*QStatusBar) NewForInheritp() *QStatusBar {
	return NewQStatusBarp()
}
func NewQStatusBarp() *QStatusBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2740779997, "_ZN10QStatusBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QStatusBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QStatusBar")
	return gothis
}

func DeleteQStatusBar(this *QStatusBar) {
	rv, err := qtrt.Qtcc1(0, "_ZN10QStatusBarD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10151() {
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
