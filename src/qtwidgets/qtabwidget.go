//  header block begin
// /usr/include/qt/QtWidgets/qtabwidget.h
// #include <qtabwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTabWidget struct {
	*QWidget
}

func (this *QTabWidget) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQTabWidgetFromPointer(cthis unsafe.Pointer) *QTabWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QTabWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qtabwidget.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTabWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public
// void QTabWidget(class QWidget *)
func NewQTabWidget(parent unsafe.Pointer) *QTabWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTabWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:72
// index:0
// Public virtual
// void ~QTabWidget()
func DeleteQTabWidget(*QTabWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:74
// index:0
// Public
// int addTab(class QWidget *, const class QString &)
func (this *QTabWidget) AddTab(widget unsafe.Pointer, arg1 *qtcore.QString) interface{} {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:75
// index:1
// Public
// int addTab(class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) AddTab_1(widget unsafe.Pointer, icon *qtgui.QIcon, label *qtcore.QString) interface{} {
	var convArg1 = icon.GetCthis()
	var convArg2 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:77
// index:0
// Public
// int insertTab(int, class QWidget *, const class QString &)
func (this *QTabWidget) InsertTab(index int, widget unsafe.Pointer, arg2 *qtcore.QString) interface{} {
	var convArg2 = arg2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:78
// index:1
// Public
// int insertTab(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) InsertTab_1(index int, widget unsafe.Pointer, icon *qtgui.QIcon, label *qtcore.QString) interface{} {
	var convArg2 = icon.GetCthis()
	var convArg3 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:80
// index:0
// Public
// void removeTab(int)
func (this *QTabWidget) RemoveTab(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9removeTabEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:82
// index:0
// Public
// bool isTabEnabled(int)
func (this *QTabWidget) IsTabEnabled(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12isTabEnabledEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:83
// index:0
// Public
// void setTabEnabled(int, _Bool)
func (this *QTabWidget) SetTabEnabled(index int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13setTabEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:85
// index:0
// Public
// QString tabText(int)
func (this *QTabWidget) TabText(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7tabTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:86
// index:0
// Public
// void setTabText(int, const class QString &)
func (this *QTabWidget) SetTabText(index int, arg1 *qtcore.QString) {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setTabTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:88
// index:0
// Public
// QIcon tabIcon(int)
func (this *QTabWidget) TabIcon(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7tabIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:89
// index:0
// Public
// void setTabIcon(int, const class QIcon &)
func (this *QTabWidget) SetTabIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setTabIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:92
// index:0
// Public
// void setTabToolTip(int, const class QString &)
func (this *QTabWidget) SetTabToolTip(index int, tip *qtcore.QString) {
	var convArg1 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13setTabToolTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:93
// index:0
// Public
// QString tabToolTip(int)
func (this *QTabWidget) TabToolTip(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget10tabToolTipEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:97
// index:0
// Public
// void setTabWhatsThis(int, const class QString &)
func (this *QTabWidget) SetTabWhatsThis(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setTabWhatsThisEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:98
// index:0
// Public
// QString tabWhatsThis(int)
func (this *QTabWidget) TabWhatsThis(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12tabWhatsThisEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:101
// index:0
// Public
// int currentIndex()
func (this *QTabWidget) CurrentIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:102
// index:0
// Public
// QWidget * currentWidget()
func (this *QTabWidget) CurrentWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:103
// index:0
// Public
// QWidget * widget(int)
func (this *QTabWidget) Widget(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:104
// index:0
// Public
// int indexOf(class QWidget *)
func (this *QTabWidget) IndexOf(widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:105
// index:0
// Public
// int count()
func (this *QTabWidget) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:109
// index:0
// Public
// QTabWidget::TabPosition tabPosition()
func (this *QTabWidget) TabPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget11tabPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:110
// index:0
// Public
// void setTabPosition(enum QTabWidget::TabPosition)
func (this *QTabWidget) SetTabPosition(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget14setTabPositionENS_11TabPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:112
// index:0
// Public
// bool tabsClosable()
func (this *QTabWidget) TabsClosable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12tabsClosableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:113
// index:0
// Public
// void setTabsClosable(_Bool)
func (this *QTabWidget) SetTabsClosable(closeable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setTabsClosableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &closeable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:115
// index:0
// Public
// bool isMovable()
func (this *QTabWidget) IsMovable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget9isMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:116
// index:0
// Public
// void setMovable(_Bool)
func (this *QTabWidget) SetMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:120
// index:0
// Public
// QTabWidget::TabShape tabShape()
func (this *QTabWidget) TabShape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8tabShapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:121
// index:0
// Public
// void setTabShape(enum QTabWidget::TabShape)
func (this *QTabWidget) SetTabShape(s int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11setTabShapeENS_8TabShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:123
// index:0
// Public virtual
// QSize sizeHint()
func (this *QTabWidget) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:124
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QTabWidget) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:125
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QTabWidget) HeightForWidth(width int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:126
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QTabWidget) HasHeightForWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:128
// index:0
// Public
// void setCornerWidget(class QWidget *, Qt::Corner)
func (this *QTabWidget) SetCornerWidget(w unsafe.Pointer, corner int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setCornerWidgetEP7QWidgetN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:129
// index:0
// Public
// QWidget * cornerWidget(Qt::Corner)
func (this *QTabWidget) CornerWidget(corner int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12cornerWidgetEN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &corner)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:131
// index:0
// Public
// Qt::TextElideMode elideMode()
func (this *QTabWidget) ElideMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget9elideModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:132
// index:0
// Public
// void setElideMode(Qt::TextElideMode)
func (this *QTabWidget) SetElideMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget12setElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:134
// index:0
// Public
// QSize iconSize()
func (this *QTabWidget) IconSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:135
// index:0
// Public
// void setIconSize(const class QSize &)
func (this *QTabWidget) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:137
// index:0
// Public
// bool usesScrollButtons()
func (this *QTabWidget) UsesScrollButtons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget17usesScrollButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:138
// index:0
// Public
// void setUsesScrollButtons(_Bool)
func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget20setUsesScrollButtonsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &useButtons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:140
// index:0
// Public
// bool documentMode()
func (this *QTabWidget) DocumentMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12documentModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:141
// index:0
// Public
// void setDocumentMode(_Bool)
func (this *QTabWidget) SetDocumentMode(set bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setDocumentModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:143
// index:0
// Public
// bool tabBarAutoHide()
func (this *QTabWidget) TabBarAutoHide() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget14tabBarAutoHideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:144
// index:0
// Public
// void setTabBarAutoHide(_Bool)
func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget17setTabBarAutoHideEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:146
// index:0
// Public
// void clear()
func (this *QTabWidget) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:148
// index:0
// Public
// QTabBar * tabBar()
func (this *QTabWidget) TabBar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget6tabBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:151
// index:0
// Public
// void setCurrentIndex(int)
func (this *QTabWidget) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:152
// index:0
// Public
// void setCurrentWidget(class QWidget *)
func (this *QTabWidget) SetCurrentWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:155
// index:0
// Public
// void currentChanged(int)
func (this *QTabWidget) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:156
// index:0
// Public
// void tabCloseRequested(int)
func (this *QTabWidget) TabCloseRequested(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget17tabCloseRequestedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:157
// index:0
// Public
// void tabBarClicked(int)
func (this *QTabWidget) TabBarClicked(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13tabBarClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:158
// index:0
// Public
// void tabBarDoubleClicked(int)
func (this *QTabWidget) TabBarDoubleClicked(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget19tabBarDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:161
// index:0
// Protected virtual
// void tabInserted(int)
func (this *QTabWidget) TabInserted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11tabInsertedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:162
// index:0
// Protected virtual
// void tabRemoved(int)
func (this *QTabWidget) TabRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10tabRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:164
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QTabWidget) ShowEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:165
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QTabWidget) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:166
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTabWidget) KeyPressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:167
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QTabWidget) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:168
// index:0
// Protected
// void setTabBar(class QTabBar *)
func (this *QTabWidget) SetTabBar(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9setTabBarEP7QTabBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:169
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QTabWidget) ChangeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:170
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QTabWidget) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtabwidget.h:171
// index:0
// Protected
// void initStyleOption(class QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
