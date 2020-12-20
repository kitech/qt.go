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
// extern C begin: 4
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
type QToolButton struct {
	*QAbstractButton
}
type QToolButton_ITF interface {
	QAbstractButton_ITF
	QToolButton_PTR() *QToolButton
}

func (ptr *QToolButton) QToolButton_PTR() *QToolButton { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QToolButtonFromptr(cthis Voidptr) *QToolButton {
	bcthis0 := QAbstractButtonFromptr(cthis)
	return &QToolButton{bcthis0}
}
func (*QToolButton) Fromptr(cthis Voidptr) *QToolButton {
	return QToolButtonFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolButton(QWidget *)

/*
 */
func (*QToolButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QToolButton {
	return NewQToolButton(parent)
}
func NewQToolButton(parent QWidget_ITF /*777 QWidget **/) *QToolButton {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(719486503, "_ZN11QToolButtonC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QToolButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolButton(QWidget *)

/*
 */
func (*QToolButton) NewForInheritp() *QToolButton {
	return NewQToolButtonp()
}
func NewQToolButtonp() *QToolButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(719486503, "_ZN11QToolButtonC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QToolButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:95
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoRaise(bool)

/*
 */
func (this *QToolButton) SetAutoRaise(enable bool) {
	rv, err := qtrt.Qtcc1(2791407414, "_ZN11QToolButton12setAutoRaiseEb", qtrt.FFITY_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:106
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void triggered(QAction *)

/*
 */
func (this *QToolButton) Triggered(arg0 QAction_ITF /*777 QAction **/) {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QAction_PTR() != nil {
		convArg0 = arg0.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(784990465, "_ZN11QToolButton9triggeredEP7QAction", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQToolButton(this *QToolButton) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QToolButtonD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QToolButton__ToolButtonPopupMode = int

//
const QToolButton__DelayedPopup QToolButton__ToolButtonPopupMode = 0

//
const QToolButton__MenuButtonPopup QToolButton__ToolButtonPopupMode = 1

//
const QToolButton__InstantPopup QToolButton__ToolButtonPopupMode = 2

func (this *QToolButton) ToolButtonPopupModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QToolButton_ToolButtonPopupModeItemName(val int) string {
	var nilthis *QToolButton
	return nilthis.ToolButtonPopupModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10161() {
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
