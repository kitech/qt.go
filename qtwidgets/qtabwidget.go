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
// extern C begin: 81
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

// void tabInserted(int)
func (this *QTabWidget) InheritTabInserted(f func(index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabInserted", f)
}

// void tabRemoved(int)
func (this *QTabWidget) InheritTabRemoved(f func(index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabRemoved", f)
}

// void showEvent(QShowEvent *)
func (this *QTabWidget) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QTabWidget) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QTabWidget) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QTabWidget) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void setTabBar(QTabBar *)
func (this *QTabWidget) InheritSetTabBar(f func(arg0 *QTabBar /*777 QTabBar **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setTabBar", f)
}

// void changeEvent(QEvent *)
func (this *QTabWidget) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// bool event(QEvent *)
func (this *QTabWidget) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InheritInitStyleOption(f func(option *QStyleOptionTabWidgetFrame /*777 QStyleOptionTabWidgetFrame **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QTabWidget struct {
	*QWidget
}
type QTabWidget_ITF interface {
	QWidget_ITF
	QTabWidget_PTR() *QTabWidget
}

func (ptr *QTabWidget) QTabWidget_PTR() *QTabWidget { return ptr }

func (this *QTabWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QTabWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQTabWidgetFromPointer(cthis unsafe.Pointer) *QTabWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QTabWidget{bcthis0}
}
func (*QTabWidget) NewFromPointer(cthis unsafe.Pointer) *QTabWidget {
	return NewQTabWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTabWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabWidget(QWidget *)

/*
Constructs a tabbed widget with parent parent.
*/
func NewQTabWidget(parent QWidget_ITF /*777 QWidget **/) *QTabWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTabWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabWidget(QWidget *)

/*
Constructs a tabbed widget with parent parent.
*/
func NewQTabWidget__() *QTabWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTabWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTabWidget()

/*

 */
func DeleteQTabWidget(this *QTabWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int addTab(QWidget *, const QString &)

/*
Adds a tab with the given page and label to the tab widget, and returns the index of the tab in the tab bar.

If the tab's label contains an ampersand, the letter following the ampersand is used as a shortcut for the tab, e.g. if the label is "Bro&wse" then Alt+W becomes a shortcut which will move the focus to this tab.

Note: If you call addTab() after show(), the layout system will try to adjust to the changes in its widgets hierarchy and may cause flicker. To prevent this, you can set the QWidget::updatesEnabled property to false prior to changes; remember to set the property to true when the changes are done, making the widget receive paint events again.

See also insertTab().
*/
func (this *QTabWidget) AddTab(widget QWidget_ITF /*777 QWidget **/, arg1 string) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:75
// index:1
// Public Visibility=Default Availability=Available
// [4] int addTab(QWidget *, const QIcon &, const QString &)

/*
Adds a tab with the given page and label to the tab widget, and returns the index of the tab in the tab bar.

If the tab's label contains an ampersand, the letter following the ampersand is used as a shortcut for the tab, e.g. if the label is "Bro&wse" then Alt+W becomes a shortcut which will move the focus to this tab.

Note: If you call addTab() after show(), the layout system will try to adjust to the changes in its widgets hierarchy and may cause flicker. To prevent this, you can set the QWidget::updatesEnabled property to false prior to changes; remember to set the property to true when the changes are done, making the widget receive paint events again.

See also insertTab().
*/
func (this *QTabWidget) AddTab_1(widget QWidget_ITF /*777 QWidget **/, icon qtgui.QIcon_ITF, label string) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, QWidget *, const QString &)

/*
Inserts a tab with the given label and page into the tab widget at the specified index, and returns the index of the inserted tab in the tab bar.

The label is displayed in the tab and may vary in appearance depending on the configuration of the tab widget.

If the tab's label contains an ampersand, the letter following the ampersand is used as a shortcut for the tab, e.g. if the label is "Bro&wse" then Alt+W becomes a shortcut which will move the focus to this tab.

If index is out of range, the tab is simply appended. Otherwise it is inserted at the specified position.

If the QTabWidget was empty before this function is called, the new page becomes the current page. Inserting a new tab at an index less than or equal to the current index will increment the current index, but keep the current page.

Note: If you call insertTab() after show(), the layout system will try to adjust to the changes in its widgets hierarchy and may cause flicker. To prevent this, you can set the QWidget::updatesEnabled property to false prior to changes; remember to set the property to true when the changes are done, making the widget receive paint events again.

See also addTab().
*/
func (this *QTabWidget) InsertTab(index int, widget QWidget_ITF /*777 QWidget **/, arg2 string) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(arg2)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:78
// index:1
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, QWidget *, const QIcon &, const QString &)

/*
Inserts a tab with the given label and page into the tab widget at the specified index, and returns the index of the inserted tab in the tab bar.

The label is displayed in the tab and may vary in appearance depending on the configuration of the tab widget.

If the tab's label contains an ampersand, the letter following the ampersand is used as a shortcut for the tab, e.g. if the label is "Bro&wse" then Alt+W becomes a shortcut which will move the focus to this tab.

If index is out of range, the tab is simply appended. Otherwise it is inserted at the specified position.

If the QTabWidget was empty before this function is called, the new page becomes the current page. Inserting a new tab at an index less than or equal to the current index will increment the current index, but keep the current page.

Note: If you call insertTab() after show(), the layout system will try to adjust to the changes in its widgets hierarchy and may cause flicker. To prevent this, you can set the QWidget::updatesEnabled property to false prior to changes; remember to set the property to true when the changes are done, making the widget receive paint events again.

See also addTab().
*/
func (this *QTabWidget) InsertTab_1(index int, widget QWidget_ITF /*777 QWidget **/, icon qtgui.QIcon_ITF, label string) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg2 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(label)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTab(int)

/*
Removes the tab at position index from this stack of widgets. The page widget itself is not deleted.

See also addTab() and insertTab().
*/
func (this *QTabWidget) RemoveTab(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9removeTabEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTabEnabled(int) const

/*
Returns true if the page at position index is enabled; otherwise returns false.

See also setTabEnabled() and QWidget::isEnabled().
*/
func (this *QTabWidget) IsTabEnabled(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12isTabEnabledEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, bool)

/*
If enable is true, the page at position index is enabled; otherwise the page at position index is disabled. The page's tab is redrawn appropriately.

QTabWidget uses QWidget::setEnabled() internally, rather than keeping a separate flag.

Note that even a disabled tab/page may be visible. If the page is visible already, QTabWidget will not hide it; if all the pages are disabled, QTabWidget will show one of them.

See also isTabEnabled() and QWidget::setEnabled().
*/
func (this *QTabWidget) SetTabEnabled(index int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13setTabEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabText(int) const

/*
Returns the label text for the tab on the page at position index.

See also setTabText().
*/
func (this *QTabWidget) TabText(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget7tabTextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabwidget.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabText(int, const QString &)

/*
Defines a new label for the page at position index's tab.

If the provided text contains an ampersand character ('&'), a shortcut is automatically created for it. The character that follows the '&' will be used as the shortcut key. Any previous shortcut will be overwritten, or cleared if no shortcut is defined by the text. See the QShortcut documentation for details (to display an actual ampersand, use '&&').

See also tabText().
*/
func (this *QTabWidget) SetTabText(index int, arg1 string) {
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10setTabTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon tabIcon(int) const

/*
Returns the icon for the tab on the page at position index.

See also setTabIcon().
*/
func (this *QTabWidget) TabIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget7tabIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabIcon(int, const QIcon &)

/*
This is an overloaded function.

Sets the icon for the tab at position index.

See also tabIcon().
*/
func (this *QTabWidget) SetTabIcon(index int, icon qtgui.QIcon_ITF) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10setTabIconEiRK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabToolTip(int, const QString &)

/*
Sets the tab tool tip for the page at position index to tip.

See also tabToolTip().
*/
func (this *QTabWidget) SetTabToolTip(index int, tip string) {
	var tmpArg1 = qtcore.NewQString_5(tip)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13setTabToolTipEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabToolTip(int) const

/*
Returns the tab tool tip for the page at position index or an empty string if no tool tip has been set.

See also setTabToolTip().
*/
func (this *QTabWidget) TabToolTip(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget10tabToolTipEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabwidget.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabWhatsThis(int, const QString &)

/*
Sets the What's This help text for the page at position index to text.

This function was introduced in  Qt 4.1.

See also tabWhatsThis().
*/
func (this *QTabWidget) SetTabWhatsThis(index int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setTabWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int) const

/*
Returns the What's This help text for the page at position index, or an empty string if no help text has been set.

This function was introduced in  Qt 4.1.

See also setTabWhatsThis().
*/
func (this *QTabWidget) TabWhatsThis(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12tabWhatsThisEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtabwidget.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex() const

/*

 */
func (this *QTabWidget) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget() const

/*
Returns a pointer to the page currently being displayed by the tab dialog. The tab dialog does its best to make sure that this value is never 0 (but if you try hard enough, it can be).

See also currentIndex() and setCurrentWidget().
*/
func (this *QTabWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget13currentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int) const

/*
Returns the tab page at index position index or 0 if the index is out of range.
*/
func (this *QTabWidget) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
Returns the index position of the page occupied by the widget w, or -1 if the widget cannot be found.
*/
func (this *QTabWidget) IndexOf(widget QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QTabWidget) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition() const

/*

 */
func (this *QTabWidget) TabPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget11tabPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(QTabWidget::TabPosition)

/*

 */
func (this *QTabWidget) SetTabPosition(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget14setTabPositionENS_11TabPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable() const

/*

 */
func (this *QTabWidget) TabsClosable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12tabsClosableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(bool)

/*

 */
func (this *QTabWidget) SetTabsClosable(closeable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setTabsClosableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), closeable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable() const

/*

 */
func (this *QTabWidget) IsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget9isMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(bool)

/*

 */
func (this *QTabWidget) SetMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10setMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape() const

/*

 */
func (this *QTabWidget) TabShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget8tabShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(QTabWidget::TabShape)

/*

 */
func (this *QTabWidget) SetTabShape(s int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11setTabShapeENS_8TabShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QTabWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().

Returns a suitable minimum size for the tab widget.
*/
func (this *QTabWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Reimplemented from QWidget::heightForWidth().
*/
func (this *QTabWidget) HeightForWidth(width int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Reimplemented from QWidget::hasHeightForWidth().
*/
func (this *QTabWidget) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *, Qt::Corner)

/*
Sets the given widget to be shown in the specified corner of the tab widget. The geometry of the widget is determined based on the widget's sizeHint() and the style().

Only the horizontal element of the corner will be used.

Passing 0 shows no widget in the corner.

Any previously set corner widget is hidden.

All widgets set here will be deleted by the tab widget when it is destroyed unless you separately reparent the widget after setting some other corner widget (or 0).

Note: Corner widgets are designed for North and South tab positions; other orientations are known to not work properly.

See also cornerWidget() and setTabPosition().
*/
func (this *QTabWidget) SetCornerWidget(w QWidget_ITF /*777 QWidget **/, corner int) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setCornerWidgetEP7QWidgetN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *, Qt::Corner)

/*
Sets the given widget to be shown in the specified corner of the tab widget. The geometry of the widget is determined based on the widget's sizeHint() and the style().

Only the horizontal element of the corner will be used.

Passing 0 shows no widget in the corner.

Any previously set corner widget is hidden.

All widgets set here will be deleted by the tab widget when it is destroyed unless you separately reparent the widget after setting some other corner widget (or 0).

Note: Corner widgets are designed for North and South tab positions; other orientations are known to not work properly.

See also cornerWidget() and setTabPosition().
*/
func (this *QTabWidget) SetCornerWidget__(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::Corner=Elaborated, Qt::Corner=Enum, , Invalid
	corner := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setCornerWidgetEP7QWidgetN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget(Qt::Corner) const

/*
Returns the widget shown in the corner of the tab widget or 0.

See also setCornerWidget().
*/
func (this *QTabWidget) CornerWidget(corner int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12cornerWidgetEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget(Qt::Corner) const

/*
Returns the widget shown in the corner of the tab widget or 0.

See also setCornerWidget().
*/
func (this *QTabWidget) CornerWidget__() *QWidget /*777 QWidget **/ {
	// arg: 0, Qt::Corner=Elaborated, Qt::Corner=Enum, , Invalid
	corner := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12cornerWidgetEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode elideMode() const

/*

 */
func (this *QTabWidget) ElideMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget9elideModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElideMode(Qt::TextElideMode)

/*

 */
func (this *QTabWidget) SetElideMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget12setElideModeEN2Qt13TextElideModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QTabWidget) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QTabWidget) SetIconSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool usesScrollButtons() const

/*

 */
func (this *QTabWidget) UsesScrollButtons() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget17usesScrollButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUsesScrollButtons(bool)

/*

 */
func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget20setUsesScrollButtonsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), useButtons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode() const

/*

 */
func (this *QTabWidget) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(bool)

/*

 */
func (this *QTabWidget) SetDocumentMode(set bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), set)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabBarAutoHide() const

/*

 */
func (this *QTabWidget) TabBarAutoHide() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget14tabBarAutoHideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabBarAutoHide(bool)

/*

 */
func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget17setTabBarAutoHideEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all the pages, but does not delete them. Calling this function is equivalent to calling removeTab() until the tab widget is empty.
*/
func (this *QTabWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QTabBar * tabBar() const

/*
Returns the current QTabBar.

See also setTabBar().
*/
func (this *QTabWidget) TabBar() *QTabBar /*777 QTabBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget6tabBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTabBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*

 */
func (this *QTabWidget) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)

/*
Makes widget the current widget. The widget used must be a page in this tab widget.

See also addTab(), setCurrentIndex(), and currentWidget().
*/
func (this *QTabWidget) SetCurrentWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget16setCurrentWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
This signal is emitted whenever the current page index changes. The parameter is the new current page index position, or -1 if there isn't a new one (for example, if there are no widgets in the QTabWidget)

Note: Notifier signal for property currentIndex.

See also currentWidget() and currentIndex.
*/
func (this *QTabWidget) CurrentChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget14currentChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)

/*
This signal is emitted when the close button on a tab is clicked. The index is the index that should be removed.

This function was introduced in  Qt 4.5.

See also setTabsClosable().
*/
func (this *QTabWidget) TabCloseRequested(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget17tabCloseRequestedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)

/*
This signal is emitted when user clicks on a tab at an index.

index refers to the tab clicked, or -1 if no tab is under the cursor.

This function was introduced in  Qt 5.2.
*/
func (this *QTabWidget) TabBarClicked(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13tabBarClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)

/*
This signal is emitted when the user double clicks on a tab at an index.

index is the index of a clicked tab, or -1 if no tab is under the cursor.

This function was introduced in  Qt 5.2.
*/
func (this *QTabWidget) TabBarDoubleClicked(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget19tabBarDoubleClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:161
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabInserted(int)

/*
This virtual handler is called after a new tab was added or inserted at position index.

See also tabRemoved().
*/
func (this *QTabWidget) TabInserted(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11tabInsertedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabRemoved(int)

/*
This virtual handler is called after a tab was removed from position index.

See also tabInserted().
*/
func (this *QTabWidget) TabRemoved(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10tabRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:164
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QTabWidget) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:165
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QTabWidget) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QTabWidget) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:167
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().

Paints the tab widget's tab bar in response to the paint event.
*/
func (this *QTabWidget) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:168
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setTabBar(QTabBar *)

/*
Replaces the dialog's QTabBar heading with the tab bar tb. Note that this must be called before any tabs have been added, or the behavior is undefined.

See also tabBar().
*/
func (this *QTabWidget) SetTabBar(arg0 QTabBar_ITF /*777 QTabBar **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTabBar_PTR() != nil {
		convArg0 = arg0.QTabBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9setTabBarEP7QTabBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QTabWidget) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:170
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QTabWidget) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:171
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionTabWidgetFrame *) const

/*
Initialize option with the values from this QTabWidget. This method is useful for subclasses when they need a QStyleOptionTabWidgetFrame, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom() and QTabBar::initStyleOption().
*/
func (this *QTabWidget) InitStyleOption(option QStyleOptionTabWidgetFrame_ITF /*777 QStyleOptionTabWidgetFrame **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionTabWidgetFrame_PTR() != nil {
		convArg0 = option.QStyleOptionTabWidgetFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum type defines where QTabWidget draws the tab row:


*/
type QTabWidget__TabPosition = int

// The tabs are drawn above the pages.
const QTabWidget__North QTabWidget__TabPosition = 0

// The tabs are drawn below the pages.
const QTabWidget__South QTabWidget__TabPosition = 1

// The tabs are drawn to the left of the pages.
const QTabWidget__West QTabWidget__TabPosition = 2

// The tabs are drawn to the right of the pages.
const QTabWidget__East QTabWidget__TabPosition = 3

/*
This enum type defines the shape of the tabs:


*/
type QTabWidget__TabShape = int

// The tabs are drawn with a rounded look. This is the default shape.
const QTabWidget__Rounded QTabWidget__TabShape = 0

// The tabs are drawn with a triangular look.
const QTabWidget__Triangular QTabWidget__TabShape = 1

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
