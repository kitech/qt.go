// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmenu.h
// #include <qmenu.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QMenu struct {
	*QWidget
}
type QMenu_ITF interface {
	QWidget_ITF
	QMenu_PTR() *QMenu
}

func (ptr *QMenu) QMenu_PTR() *QMenu { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QMenuFromptr(cthis unsafe.Pointer) *QMenu {
	bcthis0 := QWidgetFromptr(cthis)
	return &QMenu{bcthis0}
}
func (*QMenu) Fromptr(cthis unsafe.Pointer) *QMenu {
	return QMenuFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenu(QWidget *)

/*
 */
func (*QMenu) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QMenu {
	return NewQMenu(parent)
}
func NewQMenu(parent QWidget_ITF /*777 QWidget **/) *QMenu {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(32748694, "_ZN5QMenuC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QMenuFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenu(QWidget *)

/*
 */
func (*QMenu) NewForInheritp() *QMenu {
	return NewQMenup()
}
func NewQMenup() *QMenu {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(32748694, "_ZN5QMenuC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QMenuFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMenu(const QString &, QWidget *)

/*
 */
func (*QMenu) NewForInherit1(title string, parent QWidget_ITF /*777 QWidget **/) *QMenu {
	return NewQMenu1(title, parent)
}
func NewQMenu1(title string, parent QWidget_ITF /*777 QWidget **/) *QMenu {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2658403316, "_ZN5QMenuC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QMenuFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMenu(const QString &, QWidget *)

/*
 */
func (*QMenu) NewForInherit1p(title string) *QMenu {
	return NewQMenu1p(title)
}
func NewQMenu1p(title string) *QMenu {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2658403316, "_ZN5QMenuC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QMenuFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

func DeleteQMenu(this *QMenu) {
	rv, err := qtrt.Qtcc1(0, "_ZN5QMenuD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10135() {
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
