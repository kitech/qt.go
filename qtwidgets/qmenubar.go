// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmenubar.h
// #include <qmenubar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QMenuBar struct {
	*QWidget
}
type QMenuBar_ITF interface {
	QWidget_ITF
	QMenuBar_PTR() *QMenuBar
}

func (ptr *QMenuBar) QMenuBar_PTR() *QMenuBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QMenuBarFromptr(cthis unsafe.Pointer) *QMenuBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QMenuBar{bcthis0}
}
func (*QMenuBar) Fromptr(cthis unsafe.Pointer) *QMenuBar {
	return QMenuBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenuBar(QWidget *)

/*
 */
func (*QMenuBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QMenuBar {
	return NewQMenuBar(parent)
}
func NewQMenuBar(parent QWidget_ITF /*777 QWidget **/) *QMenuBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(4217600900, "_ZN8QMenuBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QMenuBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenuBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenuBar(QWidget *)

/*
 */
func (*QMenuBar) NewForInheritp() *QMenuBar {
	return NewQMenuBarp()
}
func NewQMenuBarp() *QMenuBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(4217600900, "_ZN8QMenuBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QMenuBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenuBar")
	return gothis
}

func DeleteQMenuBar(this *QMenuBar) {
	rv, err := qtrt.Qtcc1(0, "_ZN8QMenuBarD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10137() {
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
