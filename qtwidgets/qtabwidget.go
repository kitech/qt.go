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

// void showEvent(class QShowEvent *)
func (this *QTabWidget) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QTabWidget) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QTabWidget) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QTabWidget) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void setTabBar(class QTabBar *)
func (this *QTabWidget) InheritSetTabBar(f func(arg0 *QTabBar /*777 QTabBar **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setTabBar", f)
}

// void changeEvent(class QEvent *)
func (this *QTabWidget) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// bool event(class QEvent *)
func (this *QTabWidget) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(class QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InheritInitStyleOption(f func(option *QStyleOptionTabWidgetFrame /*777 QStyleOptionTabWidgetFrame **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QTabWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabWidget(QWidget *)
func NewQTabWidget(parent *QWidget /*777 QWidget **/) *QTabWidget {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTabWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTabWidget()
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
func (this *QTabWidget) AddTab(widget *QWidget /*777 QWidget **/, arg1 string) int {
	var convArg0 = widget.GetCthis()
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
func (this *QTabWidget) AddTab_1(widget *QWidget /*777 QWidget **/, icon *qtgui.QIcon, label string) int {
	var convArg0 = widget.GetCthis()
	var convArg1 = icon.GetCthis()
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
func (this *QTabWidget) InsertTab(index int, widget *QWidget /*777 QWidget **/, arg2 string) int {
	var convArg1 = widget.GetCthis()
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
func (this *QTabWidget) InsertTab_1(index int, widget *QWidget /*777 QWidget **/, icon *qtgui.QIcon, label string) int {
	var convArg1 = widget.GetCthis()
	var convArg2 = icon.GetCthis()
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
func (this *QTabWidget) RemoveTab(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9removeTabEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTabEnabled(int)
func (this *QTabWidget) IsTabEnabled(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12isTabEnabledEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, _Bool)
func (this *QTabWidget) SetTabEnabled(index int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13setTabEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabText(int)
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
func (this *QTabWidget) SetTabText(index int, arg1 string) {
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10setTabTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon tabIcon(int)
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
func (this *QTabWidget) SetTabIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10setTabIconEiRK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabToolTip(int, const QString &)
func (this *QTabWidget) SetTabToolTip(index int, tip string) {
	var tmpArg1 = qtcore.NewQString_5(tip)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13setTabToolTipEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabToolTip(int)
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
func (this *QTabWidget) SetTabWhatsThis(index int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setTabWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int)
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
// [4] int currentIndex()
func (this *QTabWidget) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget()
func (this *QTabWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget13currentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int)
func (this *QTabWidget) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *)
func (this *QTabWidget) IndexOf(widget *QWidget /*777 QWidget **/) int {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QTabWidget) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition()
func (this *QTabWidget) TabPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget11tabPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(enum QTabWidget::TabPosition)
func (this *QTabWidget) SetTabPosition(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget14setTabPositionENS_11TabPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable()
func (this *QTabWidget) TabsClosable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12tabsClosableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(_Bool)
func (this *QTabWidget) SetTabsClosable(closeable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setTabsClosableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), closeable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable()
func (this *QTabWidget) IsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget9isMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(_Bool)
func (this *QTabWidget) SetMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10setMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape()
func (this *QTabWidget) TabShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget8tabShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(enum QTabWidget::TabShape)
func (this *QTabWidget) SetTabShape(s int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11setTabShapeENS_8TabShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
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
// [8] QSize minimumSizeHint()
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
// [4] int heightForWidth(int)
func (this *QTabWidget) HeightForWidth(width int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QTabWidget) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *, Qt::Corner)
func (this *QTabWidget) SetCornerWidget(w *QWidget /*777 QWidget **/, corner int) {
	var convArg0 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setCornerWidgetEP7QWidgetN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget(Qt::Corner)
func (this *QTabWidget) CornerWidget(corner int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12cornerWidgetEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode elideMode()
func (this *QTabWidget) ElideMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget9elideModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElideMode(Qt::TextElideMode)
func (this *QTabWidget) SetElideMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget12setElideModeEN2Qt13TextElideModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
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
func (this *QTabWidget) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool usesScrollButtons()
func (this *QTabWidget) UsesScrollButtons() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget17usesScrollButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUsesScrollButtons(_Bool)
func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget20setUsesScrollButtonsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), useButtons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode()
func (this *QTabWidget) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(_Bool)
func (this *QTabWidget) SetDocumentMode(set bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), set)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabBarAutoHide()
func (this *QTabWidget) TabBarAutoHide() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget14tabBarAutoHideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabBarAutoHide(_Bool)
func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget17setTabBarAutoHideEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QTabWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QTabBar * tabBar()
func (this *QTabWidget) TabBar() *QTabBar /*777 QTabBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget6tabBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTabBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtabwidget.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QTabWidget) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)
func (this *QTabWidget) SetCurrentWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget16setCurrentWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)
func (this *QTabWidget) CurrentChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget14currentChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)
func (this *QTabWidget) TabCloseRequested(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget17tabCloseRequestedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)
func (this *QTabWidget) TabBarClicked(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13tabBarClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)
func (this *QTabWidget) TabBarDoubleClicked(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget19tabBarDoubleClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:161
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabInserted(int)
func (this *QTabWidget) TabInserted(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11tabInsertedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabRemoved(int)
func (this *QTabWidget) TabRemoved(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10tabRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:164
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QTabWidget) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:165
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QTabWidget) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTabWidget) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:167
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTabWidget) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:168
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setTabBar(QTabBar *)
func (this *QTabWidget) SetTabBar(arg0 *QTabBar /*777 QTabBar **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget9setTabBarEP7QTabBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QTabWidget) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:170
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTabWidget) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTabWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:171
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InitStyleOption(option *QStyleOptionTabWidgetFrame /*777 QStyleOptionTabWidgetFrame **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QTabWidget__TabPosition = int

const QTabWidget__North QTabWidget__TabPosition = 0
const QTabWidget__South QTabWidget__TabPosition = 1
const QTabWidget__West QTabWidget__TabPosition = 2
const QTabWidget__East QTabWidget__TabPosition = 3

type QTabWidget__TabShape = int

const QTabWidget__Rounded QTabWidget__TabShape = 0
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
