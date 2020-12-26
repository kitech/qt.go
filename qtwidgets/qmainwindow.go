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

/*
 */
// size 48
type QMainWindow struct {
	*QWidget
}
type QMainWindow_ITF interface {
	QWidget_ITF
	QMainWindow_PTR() *QMainWindow
}

func (ptr *QMainWindow) QMainWindow_PTR() *QMainWindow { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QMainWindowFromptr(cthis Voidptr) *QMainWindow {
	bcthis0 := QWidgetFromptr(cthis)
	return &QMainWindow{bcthis0}
}
func (*QMainWindow) Fromptr(cthis Voidptr) *QMainWindow {
	return QMainWindowFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)

/*
 */
func (*QMainWindow) NewForInherit(parent QWidget_ITF /*777 QWidget **/, flags int) *QMainWindow {
	return NewQMainWindow(parent, flags)
}
func NewQMainWindow(parent QWidget_ITF /*777 QWidget **/, flags int) *QMainWindow {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(1132463102, "_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&flags))
	qtrt.ErrPrint2(err, rv)
	gothis := QMainWindowFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMainWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)

/*
 */
func (*QMainWindow) NewForInheritp() *QMainWindow {
	return NewQMainWindowp()
}
func NewQMainWindowp() *QMainWindow {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(1132463102, "_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&flags))
	qtrt.ErrPrint2(err, rv)
	gothis := QMainWindowFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMainWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)

/*
 */
func (*QMainWindow) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QMainWindow {
	return NewQMainWindowp1(parent)
}
func NewQMainWindowp1(parent QWidget_ITF /*777 QWidget **/) *QMainWindow {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(1132463102, "_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&flags))
	qtrt.ErrPrint2(err, rv)
	gothis := QMainWindowFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMainWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:97
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*
 */
func (this *QMainWindow) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(69313091, "_ZNK11QMainWindow8iconSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:98
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*
 */
func (this *QMainWindow) SetIconSize(iconSize qtcore.QSize_ITF) {
	var convArg0 Voidptr
	if iconSize != nil && iconSize.QSize_PTR() != nil {
		convArg0 = iconSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(4268406733, "_ZN11QMainWindow11setIconSizeERK5QSize", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:138
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * centralWidget() const

/*
 */
func (this *QMainWindow) CentralWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(1283551286, "_ZNK11QMainWindow13centralWidgetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:139
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCentralWidget(QWidget *)

/*
 */
func (this *QMainWindow) SetCentralWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(422271864, "_ZN11QMainWindow16setCentralWidgetEP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQMainWindow(this *QMainWindow) {
	rv, err := qtrt.Qtcc3(3330525862, "_ZN11QMainWindowD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QMainWindow__DockOption = int

//
const QMainWindow__AnimatedDocks QMainWindow__DockOption = 1

//
const QMainWindow__AllowNestedDocks QMainWindow__DockOption = 2

//
const QMainWindow__AllowTabbedDocks QMainWindow__DockOption = 4

//
const QMainWindow__ForceTabbedDocks QMainWindow__DockOption = 8

//
const QMainWindow__VerticalTabs QMainWindow__DockOption = 16

//
const QMainWindow__GroupedDragging QMainWindow__DockOption = 32

func (this *QMainWindow) DockOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMainWindow_DockOptionItemName(val int) string {
	var nilthis *QMainWindow
	return nilthis.DockOptionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10223() {
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
