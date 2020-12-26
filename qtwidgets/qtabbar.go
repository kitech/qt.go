// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtabbar.h
// #include <qtabbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QTabBar struct {
	*QWidget
}
type QTabBar_ITF interface {
	QWidget_ITF
	QTabBar_PTR() *QTabBar
}

func (ptr *QTabBar) QTabBar_PTR() *QTabBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QTabBarFromptr(cthis Voidptr) *QTabBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QTabBar{bcthis0}
}
func (*QTabBar) Fromptr(cthis Voidptr) *QTabBar {
	return QTabBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qtabbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabBar(QWidget *)

/*
 */
func (*QTabBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTabBar {
	return NewQTabBar(parent)
}
func NewQTabBar(parent QWidget_ITF /*777 QWidget **/) *QTabBar {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(963516406, "_ZN7QTabBarC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QTabBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTabBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabBar(QWidget *)

/*
 */
func (*QTabBar) NewForInheritp() *QTabBar {
	return NewQTabBarp()
}
func NewQTabBarp() *QTabBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(963516406, "_ZN7QTabBarC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QTabBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTabBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabbar.h:96
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int addTab(const QString &)

/*
 */
func (this *QTabBar) AddTab(text string) int {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3484311115, "_ZN7QTabBar6addTabERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:99
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertTab(int, const QString &)

/*
 */
func (this *QTabBar) InsertTab(index int, text string) int {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc3(998099698, "_ZN7QTabBar9insertTabEiRK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:102
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void removeTab(int)

/*
 */
func (this *QTabBar) RemoveTab(index int) {
	rv, err := qtrt.Qtcc3(3929345163, "_ZN7QTabBar9removeTabEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:103
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void moveTab(int, int)

/*
 */
func (this *QTabBar) MoveTab(from int, to int) {
	rv, err := qtrt.Qtcc3(4289803665, "_ZN7QTabBar7moveTabEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&from), Voidptr(&to))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:105
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTabEnabled(int) const

/*
 */
func (this *QTabBar) IsTabEnabled(index int) bool {
	rv, err := qtrt.Qtcc3(1651733622, "_ZNK7QTabBar12isTabEnabledEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:106
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, bool)

/*
 */
func (this *QTabBar) SetTabEnabled(index int, enabled bool) {
	rv, err := qtrt.Qtcc3(3388050708, "_ZN7QTabBar13setTabEnabledEib", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&enabled))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:108
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTabVisible(int) const

/*
 */
func (this *QTabBar) IsTabVisible(index int) bool {
	rv, err := qtrt.Qtcc3(2760369443, "_ZNK7QTabBar12isTabVisibleEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabVisible(int, bool)

/*
 */
func (this *QTabBar) SetTabVisible(index int, visible bool) {
	rv, err := qtrt.Qtcc3(3526810878, "_ZN7QTabBar13setTabVisibleEib", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&visible))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:111
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString tabText(int) const

/*
 */
func (this *QTabBar) TabText(index int) string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(438043380, "_ZNK7QTabBar7tabTextEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabbar.h:112
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabText(int, const QString &)

/*
 */
func (this *QTabBar) SetTabText(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc3(794153379, "_ZN7QTabBar10setTabTextEiRK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&index), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:139
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int currentIndex() const

/*
 */
func (this *QTabBar) CurrentIndex() int {
	rv, err := qtrt.Qtcc3(4235078822, "_ZNK7QTabBar12currentIndexEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:140
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QTabBar) Count() int {
	rv, err := qtrt.Qtcc3(2981223135, "_ZNK7QTabBar5countEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:166
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isMovable() const

/*
 */
func (this *QTabBar) IsMovable() bool {
	rv, err := qtrt.Qtcc3(2296874387, "_ZNK7QTabBar9isMovableEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:167
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMovable(bool)

/*
 */
func (this *QTabBar) SetMovable(movable bool) {
	rv, err := qtrt.Qtcc3(1730991207, "_ZN7QTabBar10setMovableEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&movable))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:172
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool autoHide() const

/*
 */
func (this *QTabBar) AutoHide() bool {
	rv, err := qtrt.Qtcc3(3583682459, "_ZNK7QTabBar8autoHideEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:173
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoHide(bool)

/*
 */
func (this *QTabBar) SetAutoHide(hide bool) {
	rv, err := qtrt.Qtcc3(975943713, "_ZN7QTabBar11setAutoHideEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&hide))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:184
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*
 */
func (this *QTabBar) SetCurrentIndex(index int) {
	rv, err := qtrt.Qtcc3(2237102313, "_ZN7QTabBar15setCurrentIndexEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:187
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
 */
func (this *QTabBar) CurrentChanged(index int) {
	rv, err := qtrt.Qtcc3(1804836883, "_ZN7QTabBar14currentChangedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:188
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)

/*
 */
func (this *QTabBar) TabCloseRequested(index int) {
	rv, err := qtrt.Qtcc3(1827543677, "_ZN7QTabBar17tabCloseRequestedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:189
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabMoved(int, int)

/*
 */
func (this *QTabBar) TabMoved(from int, to int) {
	rv, err := qtrt.Qtcc3(747402176, "_ZN7QTabBar8tabMovedEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&from), Voidptr(&to))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:190
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)

/*
 */
func (this *QTabBar) TabBarClicked(index int) {
	rv, err := qtrt.Qtcc3(3685740000, "_ZN7QTabBar13tabBarClickedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:191
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)

/*
 */
func (this *QTabBar) TabBarDoubleClicked(index int) {
	rv, err := qtrt.Qtcc3(2502640271, "_ZN7QTabBar19tabBarDoubleClickedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQTabBar(this *QTabBar) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QTabBarD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTabBar__Shape = int

//
const QTabBar__RoundedNorth QTabBar__Shape = 0

//
const QTabBar__RoundedSouth QTabBar__Shape = 1

//
const QTabBar__RoundedWest QTabBar__Shape = 2

//
const QTabBar__RoundedEast QTabBar__Shape = 3

//
const QTabBar__TriangularNorth QTabBar__Shape = 4

//
const QTabBar__TriangularSouth QTabBar__Shape = 5

//
const QTabBar__TriangularWest QTabBar__Shape = 6

//
const QTabBar__TriangularEast QTabBar__Shape = 7

func (this *QTabBar) ShapeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabBar_ShapeItemName(val int) string {
	var nilthis *QTabBar
	return nilthis.ShapeItemName(val)
}

/*


 */
type QTabBar__ButtonPosition = int

//
const QTabBar__LeftSide QTabBar__ButtonPosition = 0

//
const QTabBar__RightSide QTabBar__ButtonPosition = 1

func (this *QTabBar) ButtonPositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabBar_ButtonPositionItemName(val int) string {
	var nilthis *QTabBar
	return nilthis.ButtonPositionItemName(val)
}

/*


 */
type QTabBar__SelectionBehavior = int

//
const QTabBar__SelectLeftTab QTabBar__SelectionBehavior = 0

//
const QTabBar__SelectRightTab QTabBar__SelectionBehavior = 1

//
const QTabBar__SelectPreviousTab QTabBar__SelectionBehavior = 2

func (this *QTabBar) SelectionBehaviorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabBar_SelectionBehaviorItemName(val int) string {
	var nilthis *QTabBar
	return nilthis.SelectionBehaviorItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10175() {
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
