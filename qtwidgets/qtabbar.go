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
// extern C begin: 34
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

// QSize tabSizeHint(int)
func (this *QTabBar) InheritTabSizeHint(f func(index int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "tabSizeHint", f)
}

// QSize minimumTabSizeHint(int)
func (this *QTabBar) InheritMinimumTabSizeHint(f func(index int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "minimumTabSizeHint", f)
}

// void tabInserted(int)
func (this *QTabBar) InheritTabInserted(f func(index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabInserted", f)
}

// void tabRemoved(int)
func (this *QTabBar) InheritTabRemoved(f func(index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabRemoved", f)
}

// void tabLayoutChange()
func (this *QTabBar) InheritTabLayoutChange(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabLayoutChange", f)
}

// bool event(QEvent *)
func (this *QTabBar) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QTabBar) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QTabBar) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QTabBar) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QTabBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QTabBar) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QTabBar) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QTabBar) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QTabBar) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QTabBar) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void changeEvent(QEvent *)
func (this *QTabBar) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QTabBar) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void initStyleOption(QStyleOptionTab *, int)
func (this *QTabBar) InheritInitStyleOption(f func(option *QStyleOptionTab /*777 QStyleOptionTab **/, tabIndex int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QTabBar struct {
	*QWidget
}
type QTabBar_ITF interface {
	QWidget_ITF
	QTabBar_PTR() *QTabBar
}

func (ptr *QTabBar) QTabBar_PTR() *QTabBar { return ptr }

func (this *QTabBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QTabBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQTabBarFromPointer(cthis unsafe.Pointer) *QTabBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QTabBar{bcthis0}
}
func (*QTabBar) NewFromPointer(cthis unsafe.Pointer) *QTabBar {
	return NewQTabBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtabbar.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTabBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabBar(QWidget *)

/*
Creates a new tab bar with the given parent.
*/
func (*QTabBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTabBar {
	return NewQTabBar(parent)
}
func NewQTabBar(parent QWidget_ITF /*777 QWidget **/) *QTabBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTabBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabBar(QWidget *)

/*
Creates a new tab bar with the given parent.
*/
func (*QTabBar) NewForInheritp() *QTabBar {
	return NewQTabBarp()
}
func NewQTabBarp() *QTabBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTabBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabbar.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTabBar()

/*

 */
func DeleteQTabBar(this *QTabBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtabbar.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabBar::Shape shape() const

/*

 */
func (this *QTabBar) Shape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShape(QTabBar::Shape)

/*

 */
func (this *QTabBar) SetShape(shape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar8setShapeENS_5ShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shape)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int addTab(const QString &)

/*
Adds a new tab with text text. Returns the new tab's index.
*/
func (this *QTabBar) AddTab(text string) int {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar6addTabERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:97
// index:1
// Public Visibility=Default Availability=Available
// [4] int addTab(const QIcon &, const QString &)

/*
Adds a new tab with text text. Returns the new tab's index.
*/
func (this *QTabBar) AddTab1(icon qtgui.QIcon_ITF, text string) int {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar6addTabERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, const QString &)

/*
Inserts a new tab with text text at position index. If index is out of range, the new tab is appened. Returns the new tab's index.
*/
func (this *QTabBar) InsertTab(index int, text string) int {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar9insertTabEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:100
// index:1
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, const QIcon &, const QString &)

/*
Inserts a new tab with text text at position index. If index is out of range, the new tab is appened. Returns the new tab's index.
*/
func (this *QTabBar) InsertTab1(index int, icon qtgui.QIcon_ITF, text string) int {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar9insertTabEiRK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTab(int)

/*
Removes the tab at position index.

See also SelectionBehavior.
*/
func (this *QTabBar) RemoveTab(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar9removeTabEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveTab(int, int)

/*
Moves the item at index position from to index position to.

This function was introduced in  Qt 4.5.

See also tabMoved() and tabLayoutChange().
*/
func (this *QTabBar) MoveTab(from int, to int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar7moveTabEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTabEnabled(int) const

/*
Returns true if the tab at position index is enabled; otherwise returns false.
*/
func (this *QTabBar) IsTabEnabled(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12isTabEnabledEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, bool)

/*
If enabled is true then the tab at position index is enabled; otherwise the item at position index is disabled.

See also isTabEnabled().
*/
func (this *QTabBar) SetTabEnabled(index int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar13setTabEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabText(int) const

/*
Returns the text of the tab at position index, or an empty string if index is out of range.

See also setTabText().
*/
func (this *QTabBar) TabText(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar7tabTextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabbar.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabText(int, const QString &)

/*
Sets the text of the tab at position index to text.

See also tabText().
*/
func (this *QTabBar) SetTabText(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10setTabTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:111
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor tabTextColor(int) const

/*
Returns the text color of the tab with the given index, or a invalid color if index is out of range.

See also setTabTextColor().
*/
func (this *QTabBar) TabTextColor(index int) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12tabTextColorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabTextColor(int, const QColor &)

/*
Sets the color of the text in the tab with the given index to the specified color.

If an invalid color is specified, the tab will use the QTabBar foreground role instead.

See also tabTextColor().
*/
func (this *QTabBar) SetTabTextColor(index int, color qtgui.QColor_ITF) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15setTabTextColorEiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon tabIcon(int) const

/*
Returns the icon of the tab at position index, or a null icon if index is out of range.

See also setTabIcon().
*/
func (this *QTabBar) TabIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar7tabIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabIcon(int, const QIcon &)

/*
Sets the icon of the tab at position index to icon.

See also tabIcon().
*/
func (this *QTabBar) SetTabIcon(index int, icon qtgui.QIcon_ITF) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10setTabIconEiRK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode elideMode() const

/*

 */
func (this *QTabBar) ElideMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar9elideModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElideMode(Qt::TextElideMode)

/*

 */
func (this *QTabBar) SetElideMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar12setElideModeEN2Qt13TextElideModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabToolTip(int, const QString &)

/*
Sets the tool tip of the tab at position index to tip.

See also tabToolTip().
*/
func (this *QTabBar) SetTabToolTip(index int, tip string) {
	var tmpArg1 = qtcore.NewQString5(tip)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar13setTabToolTipEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabToolTip(int) const

/*
Returns the tool tip of the tab at position index, or an empty string if index is out of range.

See also setTabToolTip().
*/
func (this *QTabBar) TabToolTip(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar10tabToolTipEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabbar.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabWhatsThis(int, const QString &)

/*
Sets the What's This help text of the tab at position index to text.

This function was introduced in  Qt 4.1.

See also tabWhatsThis().
*/
func (this *QTabBar) SetTabWhatsThis(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15setTabWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int) const

/*
Returns the What's This help text of the tab at position index, or an empty string if index is out of range.

This function was introduced in  Qt 4.1.

See also setTabWhatsThis().
*/
func (this *QTabBar) TabWhatsThis(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12tabWhatsThisEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabbar.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabData(int, const QVariant &)

/*
Sets the data of the tab at position index to data.

See also tabData().
*/
func (this *QTabBar) SetTabData(index int, data qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if data != nil && data.QVariant_PTR() != nil {
		convArg1 = data.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10setTabDataEiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:131
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant tabData(int) const

/*
Returns the data of the tab at position index, or a null variant if index is out of range.

See also setTabData().
*/
func (this *QTabBar) TabData(index int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar7tabDataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:133
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect tabRect(int) const

/*
Returns the visual rectangle of the tab at position index, or a null rectangle if index is out of range.
*/
func (this *QTabBar) TabRect(index int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar7tabRectEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int tabAt(const QPoint &) const

/*
Returns the index of the tab that covers position or -1 if no tab covers position;

This function was introduced in  Qt 4.3.
*/
func (this *QTabBar) TabAt(pos qtcore.QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar5tabAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex() const

/*

 */
func (this *QTabBar) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QTabBar) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QTabBar) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QTabBar) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDrawBase(bool)

/*

 */
func (this *QTabBar) SetDrawBase(drawTheBase bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar11setDrawBaseEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), drawTheBase)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool drawBase() const

/*

 */
func (this *QTabBar) DrawBase() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar8drawBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QTabBar) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QTabBar) SetIconSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool usesScrollButtons() const

/*

 */
func (this *QTabBar) UsesScrollButtons() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar17usesScrollButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUsesScrollButtons(bool)

/*

 */
func (this *QTabBar) SetUsesScrollButtons(useButtons bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar20setUsesScrollButtonsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), useButtons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable() const

/*

 */
func (this *QTabBar) TabsClosable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12tabsClosableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(bool)

/*

 */
func (this *QTabBar) SetTabsClosable(closable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15setTabsClosableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), closable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabButton(int, QTabBar::ButtonPosition, QWidget *)

/*
Sets widget on the tab index. The widget is placed on the left or right hand side depending upon the position.

Any previously set widget in position is hidden.

The tab bar will take ownership of the widget and so all widgets set here will be deleted by the tab bar when it is destroyed unless you separately reparent the widget after setting some other widget (or 0).

This function was introduced in  Qt 4.5.

See also tabButton() and tabsClosable().
*/
func (this *QTabBar) SetTabButton(index int, position int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar12setTabButtonEiNS_14ButtonPositionEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, position, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * tabButton(int, QTabBar::ButtonPosition) const

/*
Returns the widget set a tab index and position or 0 if one is not set.

See also setTabButton().
*/
func (this *QTabBar) TabButton(index int, position int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar9tabButtonEiNS_14ButtonPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, position)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabbar.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabBar::SelectionBehavior selectionBehaviorOnRemove() const

/*

 */
func (this *QTabBar) SelectionBehaviorOnRemove() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar25selectionBehaviorOnRemoveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionBehaviorOnRemove(QTabBar::SelectionBehavior)

/*

 */
func (this *QTabBar) SetSelectionBehaviorOnRemove(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar28setSelectionBehaviorOnRemoveENS_17SelectionBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool expanding() const

/*

 */
func (this *QTabBar) Expanding() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar9expandingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpanding(bool)

/*

 */
func (this *QTabBar) SetExpanding(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar12setExpandingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable() const

/*

 */
func (this *QTabBar) IsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar9isMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(bool)

/*

 */
func (this *QTabBar) SetMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10setMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode() const

/*

 */
func (this *QTabBar) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(bool)

/*

 */
func (this *QTabBar) SetDocumentMode(set bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), set)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoHide() const

/*

 */
func (this *QTabBar) AutoHide() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar8autoHideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoHide(bool)

/*

 */
func (this *QTabBar) SetAutoHide(hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar11setAutoHideEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:172
// index:0
// Public Visibility=Default Availability=Available
// [1] bool changeCurrentOnDrag() const

/*

 */
func (this *QTabBar) ChangeCurrentOnDrag() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar19changeCurrentOnDragEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChangeCurrentOnDrag(bool)

/*

 */
func (this *QTabBar) SetChangeCurrentOnDrag(change bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar22setChangeCurrentOnDragEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QString accessibleTabName(int) const

/*
Returns the accessibleName of the tab at position index, or an empty string if index is out of range.

See also setAccessibleTabName().
*/
func (this *QTabBar) AccessibleTabName(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar17accessibleTabNameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabbar.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccessibleTabName(int, const QString &)

/*
Sets the accessibleName of the tab at position index to name.

See also accessibleTabName().
*/
func (this *QTabBar) SetAccessibleTabName(index int, name string) {
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar20setAccessibleTabNameEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*

 */
func (this *QTabBar) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
This signal is emitted when the tab bar's current tab changes. The new current has the given index, or -1 if there isn't a new one (for example, if there are no tab in the QTabBar)

Note: Notifier signal for property currentIndex.
*/
func (this *QTabBar) CurrentChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar14currentChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)

/*
This signal is emitted when the close button on a tab is clicked. The index is the index that should be removed.

This function was introduced in  Qt 4.5.

See also setTabsClosable().
*/
func (this *QTabBar) TabCloseRequested(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar17tabCloseRequestedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabMoved(int, int)

/*
This signal is emitted when the tab has moved the tab at index position from to index position to.

note: QTabWidget will automatically move the page when this signal is emitted from its tab bar.

This function was introduced in  Qt 4.5.

See also moveTab().
*/
func (this *QTabBar) TabMoved(from int, to int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar8tabMovedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)

/*
This signal is emitted when user clicks on a tab at an index.

index is the index of a clicked tab, or -1 if no tab is under the cursor.

This function was introduced in  Qt 5.2.
*/
func (this *QTabBar) TabBarClicked(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar13tabBarClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)

/*
This signal is emitted when the user double clicks on a tab at index.

index refers to the tab clicked, or -1 if no tab is under the cursor.

This function was introduced in  Qt 5.2.
*/
func (this *QTabBar) TabBarDoubleClicked(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar19tabBarDoubleClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize tabSizeHint(int) const

/*
Returns the size hint for the tab at position index.
*/
func (this *QTabBar) TabSizeHint(index int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar11tabSizeHintEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize minimumTabSizeHint(int) const

/*
Returns the minimum tab size hint for the tab at position index.

This function was introduced in  Qt 5.0.
*/
func (this *QTabBar) MinimumTabSizeHint(index int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar18minimumTabSizeHintEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:193
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabInserted(int)

/*
This virtual handler is called after a new tab was added or inserted at position index.

See also tabRemoved().
*/
func (this *QTabBar) TabInserted(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar11tabInsertedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabRemoved(int)

/*
This virtual handler is called after a tab was removed from position index.

See also tabInserted().
*/
func (this *QTabBar) TabRemoved(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10tabRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:195
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabLayoutChange()

/*
This virtual handler is called whenever the tab layout changes.

See also tabRect().
*/
func (this *QTabBar) TabLayoutChange() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15tabLayoutChangeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:197
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QTabBar) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:198
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QTabBar) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:199
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QTabBar) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QTabBar) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:201
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QTabBar) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QTabBar) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:203
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QTabBar) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:204
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QTabBar) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QTabBar) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QTabBar) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:209
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QTabBar) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QTabBar) TimerEvent(event qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QTabBar10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:211
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionTab *, int) const

/*
Initialize option with the values from the tab at tabIndex. This method is useful for subclasses when they need a QStyleOptionTab, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom() and QTabWidget::initStyleOption().
*/
func (this *QTabBar) InitStyleOption(option QStyleOptionTab_ITF /*777 QStyleOptionTab **/, tabIndex int) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionTab_PTR() != nil {
		convArg0 = option.QStyleOptionTab_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, tabIndex)
	qtrt.ErrPrint(err, rv)
}

/*
This enum type lists the built-in shapes supported by QTabBar. Treat these as hints as some styles may not render some of the shapes. However, position should be honored.


*/
type QTabBar__Shape = int

// The normal rounded look above the pages
const QTabBar__RoundedNorth QTabBar__Shape = 0

// The normal rounded look below the pages
const QTabBar__RoundedSouth QTabBar__Shape = 1

// The normal rounded look on the left side of the pages
const QTabBar__RoundedWest QTabBar__Shape = 2

// The normal rounded look on the right side the pages
const QTabBar__RoundedEast QTabBar__Shape = 3

// Triangular tabs above the pages.
const QTabBar__TriangularNorth QTabBar__Shape = 4

// Triangular tabs similar to those used in the Excel spreadsheet, for example
const QTabBar__TriangularSouth QTabBar__Shape = 5

// Triangular tabs on the left of the pages.
const QTabBar__TriangularWest QTabBar__Shape = 6

// Triangular tabs on the right of the pages.
const QTabBar__TriangularEast QTabBar__Shape = 7

func (this *QTabBar) ShapeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabBar_ShapeItemName(val int) string {
	var nilthis *QTabBar
	return nilthis.ShapeItemName(val)
}

/*
This enum type lists the location of the widget on a tab.



This enum was introduced or modified in  Qt 4.5.

*/
type QTabBar__ButtonPosition = int

// Left side of the tab.
const QTabBar__LeftSide QTabBar__ButtonPosition = 0

// Right side of the tab.
const QTabBar__RightSide QTabBar__ButtonPosition = 1

func (this *QTabBar) ButtonPositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QTabBar_ButtonPositionItemName(val int) string {
	var nilthis *QTabBar
	return nilthis.ButtonPositionItemName(val)
}

/*
This enum type lists the behavior of QTabBar when a tab is removed and the tab being removed is also the current tab.



This enum was introduced or modified in  Qt 4.5.

*/
type QTabBar__SelectionBehavior = int

// Select the tab to the left of the one being removed.
const QTabBar__SelectLeftTab QTabBar__SelectionBehavior = 0

// Select the tab to the right of the one being removed.
const QTabBar__SelectRightTab QTabBar__SelectionBehavior = 1

// Select the previously selected tab.
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

func init() {
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
