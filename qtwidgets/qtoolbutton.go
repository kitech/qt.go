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
func QToolButtonFromptr(cthis unsafe.Pointer) *QToolButton {
	bcthis0 := QAbstractButtonFromptr(cthis)
	return &QToolButton{bcthis0}
}
func (*QToolButton) Fromptr(cthis unsafe.Pointer) *QToolButton {
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
	var convArg0 unsafe.Pointer
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
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(719486503, "_ZN11QToolButtonC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QToolButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolButton")
	return gothis
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

func init_unused_10157() {
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
