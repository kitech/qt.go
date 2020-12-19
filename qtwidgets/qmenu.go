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

// /usr/include/qt/QtWidgets/qmenu.h:79
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)

/*
 */
func (this *QMenu) AddAction(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(1698034204, "_ZN5QMenu9addActionERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ QActionFromptr(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:152
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * addMenu(QMenu *)

/*
 */
func (this *QMenu) AddMenu(menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(4017241996, "_ZN5QMenu7addMenuEPS_", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ QActionFromptr(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:153
// index:1
// Public Direct Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QString &)

/*
 */
func (this *QMenu) AddMenu1(title string) *QMenu /*777 QMenu **/ {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(2540271396, "_ZN5QMenu7addMenuERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ QMenuFromptr(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:156
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAction * addSeparator()

/*
 */
func (this *QMenu) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := qtrt.Qtcc1(3779268459, "_ZN5QMenu12addSeparatorEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ QActionFromptr(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:166
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
 */
func (this *QMenu) IsEmpty() bool {
	rv, err := qtrt.Qtcc1(1363387044, "_ZNK5QMenu7isEmptyEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:167
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QMenu) Clear() {
	rv, err := qtrt.Qtcc1(3862211293, "_ZN5QMenu5clearEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:200
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString title() const

/*
 */
func (this *QMenu) Title() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(2610552010, "_ZNK5QMenu5titleEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qmenu.h:201
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*
 */
func (this *QMenu) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(841200333, "_ZN5QMenu8setTitleERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:222
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void aboutToShow()

/*
 */
func (this *QMenu) AboutToShow() {
	rv, err := qtrt.Qtcc1(1566663078, "_ZN5QMenu11aboutToShowEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:223
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void aboutToHide()

/*
 */
func (this *QMenu) AboutToHide() {
	rv, err := qtrt.Qtcc1(3249752657, "_ZN5QMenu11aboutToHideEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:224
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void triggered(QAction *)

/*
 */
func (this *QMenu) Triggered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(576772310, "_ZN5QMenu9triggeredEP7QAction", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:225
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void hovered(QAction *)

/*
 */
func (this *QMenu) Hovered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(2472517649, "_ZN5QMenu7hoveredEP7QAction", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQMenu(this *QMenu) {
	rv, err := qtrt.Qtcc1(0, "_ZN5QMenuD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10139() {
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
