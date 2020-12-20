// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtabwidget.h
// #include <qtabwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QTabWidget struct {
	*QWidget
}
type QTabWidget_ITF interface {
	QWidget_ITF
	QTabWidget_PTR() *QTabWidget
}

func (ptr *QTabWidget) QTabWidget_PTR() *QTabWidget { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QTabWidgetFromptr(cthis Voidptr) *QTabWidget {
	bcthis0 := QWidgetFromptr(cthis)
	return &QTabWidget{bcthis0}
}
func (*QTabWidget) Fromptr(cthis Voidptr) *QTabWidget {
	return QTabWidgetFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabWidget(QWidget *)

/*
 */
func (*QTabWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTabWidget {
	return NewQTabWidget(parent)
}
func NewQTabWidget(parent QWidget_ITF /*777 QWidget **/) *QTabWidget {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2398945872, "_ZN10QTabWidgetC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QTabWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTabWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabWidget(QWidget *)

/*
 */
func (*QTabWidget) NewForInheritp() *QTabWidget {
	return NewQTabWidgetp()
}
func NewQTabWidgetp() *QTabWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2398945872, "_ZN10QTabWidgetC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QTabWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QTabWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:74
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int addTab(QWidget *, const QString &)

/*
 */
func (this *QTabWidget) AddTab(widget QWidget_ITF /*777 QWidget **/, arg1 string) int {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc1(1498629996, "_ZN10QTabWidget6addTabEP7QWidgetRK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:77
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertTab(int, QWidget *, const QString &)

/*
 */
func (this *QTabWidget) InsertTab(index int, widget QWidget_ITF /*777 QWidget **/, arg2 string) int {
	var convArg1 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(arg2)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.Qtcc1(3896844051, "_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:80
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void removeTab(int)

/*
 */
func (this *QTabWidget) RemoveTab(index int) {
	rv, err := qtrt.Qtcc1(1571166509, "_ZN10QTabWidget9removeTabEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:82
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTabEnabled(int) const

/*
 */
func (this *QTabWidget) IsTabEnabled(index int) bool {
	rv, err := qtrt.Qtcc1(3901106487, "_ZNK10QTabWidget12isTabEnabledEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:83
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, bool)

/*
 */
func (this *QTabWidget) SetTabEnabled(index int, enabled bool) {
	rv, err := qtrt.Qtcc1(3243295178, "_ZN10QTabWidget13setTabEnabledEib", qtrt.FFITY_POINTER, this.GetCthis(), index, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:85
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTabVisible(int) const

/*
 */
func (this *QTabWidget) IsTabVisible(index int) bool {
	rv, err := qtrt.Qtcc1(779270242, "_ZNK10QTabWidget12isTabVisibleEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:86
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabVisible(int, bool)

/*
 */
func (this *QTabWidget) SetTabVisible(index int, visible bool) {
	rv, err := qtrt.Qtcc1(3667365920, "_ZN10QTabWidget13setTabVisibleEib", qtrt.FFITY_POINTER, this.GetCthis(), index, visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:88
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString tabText(int) const

/*
 */
func (this *QTabWidget) TabText(index int) string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(3660912747, "_ZNK10QTabWidget7tabTextEi", qtrt.FFITY_POINTER, sretobj, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabwidget.h:89
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabText(int, const QString &)

/*
 */
func (this *QTabWidget) SetTabText(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc1(1638282860, "_ZN10QTabWidget10setTabTextEiRK7QString", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:104
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int currentIndex() const

/*
 */
func (this *QTabWidget) CurrentIndex() int {
	rv, err := qtrt.Qtcc1(1989898727, "_ZNK10QTabWidget12currentIndexEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:105
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * currentWidget() const

/*
 */
func (this *QTabWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc1(1593011613, "_ZNK10QTabWidget13currentWidgetEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:106
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * widget(int) const

/*
 */
func (this *QTabWidget) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc1(1033250205, "_ZNK10QTabWidget6widgetEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:107
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
 */
func (this *QTabWidget) IndexOf(widget QWidget_ITF /*777 QWidget **/) int {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(1793686560, "_ZNK10QTabWidget7indexOfEP7QWidget", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:108
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QTabWidget) Count() int {
	rv, err := qtrt.Qtcc1(3464247378, "_ZNK10QTabWidget5countEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:115
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool tabsClosable() const

/*
 */
func (this *QTabWidget) TabsClosable() bool {
	rv, err := qtrt.Qtcc1(2589699479, "_ZNK10QTabWidget12tabsClosableEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:116
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTabsClosable(bool)

/*
 */
func (this *QTabWidget) SetTabsClosable(closeable bool) {
	rv, err := qtrt.Qtcc1(2670076814, "_ZN10QTabWidget15setTabsClosableEb", qtrt.FFITY_POINTER, this.GetCthis(), closeable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:118
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isMovable() const

/*
 */
func (this *QTabWidget) IsMovable() bool {
	rv, err := qtrt.Qtcc1(2203855381, "_ZNK10QTabWidget9isMovableEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:119
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMovable(bool)

/*
 */
func (this *QTabWidget) SetMovable(movable bool) {
	rv, err := qtrt.Qtcc1(4071398512, "_ZN10QTabWidget10setMovableEb", qtrt.FFITY_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:154
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*
 */
func (this *QTabWidget) SetCurrentIndex(index int) {
	rv, err := qtrt.Qtcc1(3828622966, "_ZN10QTabWidget15setCurrentIndexEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:155
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)

/*
 */
func (this *QTabWidget) SetCurrentWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(1617009302, "_ZN10QTabWidget16setCurrentWidgetEP7QWidget", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:158
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
 */
func (this *QTabWidget) CurrentChanged(index int) {
	rv, err := qtrt.Qtcc1(1664281805, "_ZN10QTabWidget14currentChangedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:159
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)

/*
 */
func (this *QTabWidget) TabCloseRequested(index int) {
	rv, err := qtrt.Qtcc1(384078136, "_ZN10QTabWidget17tabCloseRequestedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:160
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)

/*
 */
func (this *QTabWidget) TabBarClicked(index int) {
	rv, err := qtrt.Qtcc1(400432813, "_ZN10QTabWidget13tabBarClickedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:161
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)

/*
 */
func (this *QTabWidget) TabBarDoubleClicked(index int) {
	rv, err := qtrt.Qtcc1(1649400173, "_ZN10QTabWidget19tabBarDoubleClickedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

func DeleteQTabWidget(this *QTabWidget) {
	rv, err := qtrt.Qtcc1(0, "_ZN10QTabWidgetD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTabWidget__TabPosition = int

//
const QTabWidget__North QTabWidget__TabPosition = 0

//
const QTabWidget__South QTabWidget__TabPosition = 1

//
const QTabWidget__West QTabWidget__TabPosition = 2

//
const QTabWidget__East QTabWidget__TabPosition = 3

func (this *QTabWidget) TabPositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabWidget_TabPositionItemName(val int) string {
	var nilthis *QTabWidget
	return nilthis.TabPositionItemName(val)
}

/*


 */
type QTabWidget__TabShape = int

//
const QTabWidget__Rounded QTabWidget__TabShape = 0

//
const QTabWidget__Triangular QTabWidget__TabShape = 1

func (this *QTabWidget) TabShapeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabWidget_TabShapeItemName(val int) string {
	var nilthis *QTabWidget
	return nilthis.TabShapeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10091() {
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
