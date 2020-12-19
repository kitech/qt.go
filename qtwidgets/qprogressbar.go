// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qprogressbar.h
// #include <qprogressbar.h>
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
type QProgressBar struct {
	*QWidget
}
type QProgressBar_ITF interface {
	QWidget_ITF
	QProgressBar_PTR() *QProgressBar
}

func (ptr *QProgressBar) QProgressBar_PTR() *QProgressBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QProgressBarFromptr(cthis unsafe.Pointer) *QProgressBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QProgressBar{bcthis0}
}
func (*QProgressBar) Fromptr(cthis unsafe.Pointer) *QProgressBar {
	return QProgressBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressBar(QWidget *)

/*
 */
func (*QProgressBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QProgressBar {
	return NewQProgressBar(parent)
}
func NewQProgressBar(parent QWidget_ITF /*777 QWidget **/) *QProgressBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(572852642, "_ZN12QProgressBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QProgressBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QProgressBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressBar(QWidget *)

/*
 */
func (*QProgressBar) NewForInheritp() *QProgressBar {
	return NewQProgressBarp()
}
func NewQProgressBarp() *QProgressBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(572852642, "_ZN12QProgressBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QProgressBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QProgressBar")
	return gothis
}

func DeleteQProgressBar(this *QProgressBar) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QProgressBarD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QProgressBar__Direction = int

//
const QProgressBar__TopToBottom QProgressBar__Direction = 0

//
const QProgressBar__BottomToTop QProgressBar__Direction = 1

func (this *QProgressBar) DirectionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProgressBar_DirectionItemName(val int) string {
	var nilthis *QProgressBar
	return nilthis.DirectionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10141() {
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
