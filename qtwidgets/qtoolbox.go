package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbox.h
// #include <qtoolbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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

// bool event(QEvent *)
func (this *QToolBox) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void itemInserted(int)
func (this *QToolBox) InheritItemInserted(f func(index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "itemInserted", f)
}

// void itemRemoved(int)
func (this *QToolBox) InheritItemRemoved(f func(index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "itemRemoved", f)
}

// void showEvent(QShowEvent *)
func (this *QToolBox) InheritShowEvent(f func(e *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void changeEvent(QEvent *)
func (this *QToolBox) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

/*

 */
type QToolBox struct {
	*QFrame
}
type QToolBox_ITF interface {
	QFrame_ITF
	QToolBox_PTR() *QToolBox
}

func (ptr *QToolBox) QToolBox_PTR() *QToolBox { return ptr }

func (this *QToolBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QToolBox) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQToolBoxFromPointer(cthis unsafe.Pointer) *QToolBox {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QToolBox{bcthis0}
}
func (*QToolBox) NewFromPointer(cthis unsafe.Pointer) *QToolBox {
	return NewQToolBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QToolBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbox.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBox(QWidget *, Qt::WindowFlags)

/*
Constructs a new toolbox with the given parent and the flags, f.
*/
func NewQToolBox(parent QWidget_ITF /*777 QWidget **/, f int) *QToolBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBoxC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbox.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBox(QWidget *, Qt::WindowFlags)

/*
Constructs a new toolbox with the given parent and the flags, f.
*/
func NewQToolBox__() *QToolBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBoxC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbox.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBox(QWidget *, Qt::WindowFlags)

/*
Constructs a new toolbox with the given parent and the flags, f.
*/
func NewQToolBox__1(parent QWidget_ITF /*777 QWidget **/) *QToolBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBoxC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbox.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QToolBox()

/*

 */
func DeleteQToolBox(this *QToolBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int addItem(QWidget *, const QString &)

/*
Adds the widget in a new tab at bottom of the toolbox. The new tab's text is set to text, and the iconSet is displayed to the left of the text. Returns the new tab's index.
*/
func (this *QToolBox) AddItem(widget QWidget_ITF /*777 QWidget **/, text string) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:64
// index:1
// Public Visibility=Default Availability=Available
// [4] int addItem(QWidget *, const QIcon &, const QString &)

/*
Adds the widget in a new tab at bottom of the toolbox. The new tab's text is set to text, and the iconSet is displayed to the left of the text. Returns the new tab's index.
*/
func (this *QToolBox) AddItem_1(widget QWidget_ITF /*777 QWidget **/, icon qtgui.QIcon_ITF, text string) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertItem(int, QWidget *, const QString &)

/*
Inserts the widget at position index, or at the bottom of the toolbox if index is out of range. The new item's text is set to text, and the icon is displayed to the left of the text. Returns the new item's index.
*/
func (this *QToolBox) InsertItem(index int, widget QWidget_ITF /*777 QWidget **/, text string) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:66
// index:1
// Public Visibility=Default Availability=Available
// [4] int insertItem(int, QWidget *, const QIcon &, const QString &)

/*
Inserts the widget at position index, or at the bottom of the toolbox if index is out of range. The new item's text is set to text, and the icon is displayed to the left of the text. Returns the new item's index.
*/
func (this *QToolBox) InsertItem_1(index int, widget QWidget_ITF /*777 QWidget **/, icon qtgui.QIcon_ITF, text string) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg2 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(int)

/*
Removes the item at position index from the toolbox. Note that the widget is not deleted.
*/
func (this *QToolBox) RemoveItem(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox10removeItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEnabled(int, bool)

/*
If enabled is true then the item at position index is enabled; otherwise the item at position index is disabled.

See also isItemEnabled().
*/
func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox14setItemEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemEnabled(int) const

/*
Returns true if the item at position index is enabled; otherwise returns false.
*/
func (this *QToolBox) IsItemEnabled(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox13isItemEnabledEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbox.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemText(int, const QString &)

/*
Sets the text of the item at position index to text.

If the provided text contains an ampersand character ('&'), a mnemonic is automatically created for it. The character that follows the '&' will be used as the shortcut key. Any previous mnemonic will be overwritten, or cleared if no mnemonic is defined by the text. See the QShortcut documentation for details (to display an actual ampersand, use '&&').

See also itemText().
*/
func (this *QToolBox) SetItemText(index int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox11setItemTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString itemText(int) const

/*
Returns the text of the item at position index, or an empty string if index is out of range.

See also setItemText().
*/
func (this *QToolBox) ItemText(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox8itemTextEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtoolbox.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemIcon(int, const QIcon &)

/*
Sets the icon of the item at position index to icon.

See also itemIcon().
*/
func (this *QToolBox) SetItemIcon(index int, icon qtgui.QIcon_ITF) {
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox11setItemIconEiRK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon itemIcon(int) const

/*
Returns the icon of the item at position index, or a null icon if index is out of range.

See also setItemIcon().
*/
func (this *QToolBox) ItemIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox8itemIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbox.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemToolTip(int, const QString &)

/*
Sets the tooltip of the item at position index to toolTip.

See also itemToolTip().
*/
func (this *QToolBox) SetItemToolTip(index int, toolTip string) {
	var tmpArg1 = qtcore.NewQString_5(toolTip)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox14setItemToolTipEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QString itemToolTip(int) const

/*
Returns the tooltip of the item at position index, or an empty string if index is out of range.

See also setItemToolTip().
*/
func (this *QToolBox) ItemToolTip(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox11itemToolTipEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtoolbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex() const

/*

 */
func (this *QToolBox) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget() const

/*
Returns a pointer to the current widget, or 0 if there is no such item.

See also currentIndex() and setCurrentWidget().
*/
func (this *QToolBox) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox13currentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbox.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int) const

/*
Returns the widget at position index, or 0 if there is no such item.
*/
func (this *QToolBox) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbox.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
Returns the index of widget, or -1 if the item does not exist.
*/
func (this *QToolBox) IndexOf(widget QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QToolBox) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBox5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtoolbox.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*

 */
func (this *QToolBox) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)

/*
Makeswidget the current widget. The widget must be an item in this tool box.

See also addItem(), setCurrentIndex(), and currentWidget().
*/
func (this *QToolBox) SetCurrentWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox16setCurrentWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)

/*
This signal is emitted when the current item is changed. The new current item's index is passed in index, or -1 if there is no current item.

Note: Notifier signal for property currentIndex.
*/
func (this *QToolBox) CurrentChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox14currentChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QToolBox) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbox.h:99
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void itemInserted(int)

/*
This virtual handler is called after a new item was added or inserted at position index.

See also itemRemoved().
*/
func (this *QToolBox) ItemInserted(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox12itemInsertedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void itemRemoved(int)

/*
This virtual handler is called after an item was removed from position index.

See also itemInserted().
*/
func (this *QToolBox) ItemRemoved(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox11itemRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QToolBox) ShowEvent(e qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QShowEvent_PTR() != nil {
		convArg0 = e.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:102
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QToolBox) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
