// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbar.h
// #include <qtoolbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QToolBar struct {
	*QWidget
}
type QToolBar_ITF interface {
	QWidget_ITF
	QToolBar_PTR() *QToolBar
}

func (ptr *QToolBar) QToolBar_PTR() *QToolBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QToolBarFromptr(cthis Voidptr) *QToolBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QToolBar{bcthis0}
}
func (*QToolBar) Fromptr(cthis Voidptr) *QToolBar {
	return QToolBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(const QString &, QWidget *)

/*
 */
func (*QToolBar) NewForInherit(title string, parent QWidget_ITF /*777 QWidget **/) *QToolBar {
	return NewQToolBar(title, parent)
}
func NewQToolBar(title string, parent QWidget_ITF /*777 QWidget **/) *QToolBar {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(115476381, "_ZN8QToolBarC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	gothis := QToolBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(const QString &, QWidget *)

/*
 */
func (*QToolBar) NewForInheritp(title string) *QToolBar {
	return NewQToolBarp(title)
}
func NewQToolBarp(title string) *QToolBar {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(115476381, "_ZN8QToolBarC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	gothis := QToolBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:73
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(QWidget *)

/*
 */
func (*QToolBar) NewForInherit1(parent QWidget_ITF /*777 QWidget **/) *QToolBar {
	return NewQToolBar1(parent)
}
func NewQToolBar1(parent QWidget_ITF /*777 QWidget **/) *QToolBar {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(2552095711, "_ZN8QToolBarC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QToolBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:73
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(QWidget *)

/*
 */
func (*QToolBar) NewForInherit1p() *QToolBar {
	return NewQToolBar1p()
}
func NewQToolBar1p() *QToolBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(2552095711, "_ZN8QToolBarC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QToolBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:76
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMovable(bool)

/*
 */
func (this *QToolBar) SetMovable(movable bool) {
	rv, err := qtrt.Qtcc3(2579160314, "_ZN8QToolBar10setMovableEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&movable))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:77
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isMovable() const

/*
 */
func (this *QToolBar) IsMovable() bool {
	rv, err := qtrt.Qtcc3(1223047420, "_ZNK8QToolBar9isMovableEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

func DeleteQToolBar(this *QToolBar) {
	rv, err := qtrt.Qtcc3(1272080248, "_ZN8QToolBarD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10255() {
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
