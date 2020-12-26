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
// extern C begin: 14
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
// ignore GetCthis for 1 base
func QMenuBarFromptr(cthis Voidptr) *QMenuBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QMenuBar{bcthis0}
}
func (*QMenuBar) Fromptr(cthis Voidptr) *QMenuBar {
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
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(4217600900, "_ZN8QMenuBarC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
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
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(4217600900, "_ZN8QMenuBarC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QMenuBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenuBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenubar.h:67
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)

/*
 */
func (this *QMenuBar) AddAction(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(524583219, "_ZN8QMenuBar9addActionERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QActionFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:96
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * addMenu(QMenu *)

/*
 */
func (this *QMenuBar) AddMenu(menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 Voidptr
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1432639305, "_ZN8QMenuBar7addMenuEP5QMenu", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QActionFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:97
// index:1
// Public Direct Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QString &)

/*
 */
func (this *QMenuBar) AddMenu1(title string) *QMenu /*777 QMenu **/ {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3209398129, "_ZN8QMenuBar7addMenuERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QMenuFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:98
// index:2
// Public Direct Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QIcon &, const QString &)

/*
 */
func (this *QMenuBar) AddMenu2(icon qtgui.QIcon_ITF, title string) *QMenu /*777 QMenu **/ {
	var convArg0 Voidptr
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(title)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc3(3402422444, "_ZN8QMenuBar7addMenuERK5QIconRK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QMenuFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:101
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * addSeparator()

/*
 */
func (this *QMenuBar) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := qtrt.Qtcc3(2911490074, "_ZN8QMenuBar12addSeparatorEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QActionFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:102
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * insertSeparator(QAction *)

/*
 */
func (this *QMenuBar) InsertSeparator(before QAction_ITF /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 Voidptr
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(136532056, "_ZN8QMenuBar15insertSeparatorEP7QAction", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QActionFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:104
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * insertMenu(QAction *, QMenu *)

/*
 */
func (this *QMenuBar) InsertMenu(before QAction_ITF /*777 QAction **/, menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 Voidptr
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 Voidptr
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg1 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2471833135, "_ZN8QMenuBar10insertMenuEP7QActionP5QMenu", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QActionFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenubar.h:106
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QMenuBar) Clear() {
	rv, err := qtrt.Qtcc3(1047224492, "_ZN8QMenuBar5clearEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:132
// index:0
// Public virtual Ignore Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
 */
func (this *QMenuBar) SetVisible(visible bool) {
	rv, err := qtrt.Qtcc3(1756387468, "_ZN8QMenuBar10setVisibleEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&visible))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:135
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void triggered(QAction *)

/*
 */
func (this *QMenuBar) Triggered(action QAction_ITF /*777 QAction **/) {
	var convArg0 Voidptr
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(961109853, "_ZN8QMenuBar9triggeredEP7QAction", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:136
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void hovered(QAction *)

/*
 */
func (this *QMenuBar) Hovered(action QAction_ITF /*777 QAction **/) {
	var convArg0 Voidptr
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3004656425, "_ZN8QMenuBar7hoveredEP7QAction", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQMenuBar(this *QMenuBar) {
	rv, err := qtrt.Qtcc3(4257493948, "_ZN8QMenuBarD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10227() {
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
