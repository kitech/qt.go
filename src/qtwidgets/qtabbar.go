//  header block begin
// /usr/include/qt/QtWidgets/qtabbar.h
// #include <qtabbar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QTabBar struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtabbar.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTabBar) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:74
// index:0
// void QTabBar(class QWidget *)
func NewQTabBar(parent unsafe.Pointer) *QTabBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTabBar{cthis}
}

// /usr/include/qt/QtWidgets/qtabbar.h:75
// index:0
// virtual
// void ~QTabBar()
func DeleteQTabBar(*QTabBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:93
// index:0
// QTabBar::Shape shape()
func (this *QTabBar) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:94
// index:0
// void setShape(enum QTabBar::Shape)
func (this *QTabBar) SetShape(shape int) {
	// 0: (, QTabBar::Shape shape), (&shape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar8setShapeENS_5ShapeE", ffiqt.FFI_TYPE_VOID, this.cthis, &shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:96
// index:0
// int addTab(const class QString &)
func (this *QTabBar) AddTab(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar6addTabERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:97
// index:1
// int addTab(const class QIcon &, const class QString &)
func (this *QTabBar) AddTab_1(icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, const QIcon & icon, const QString & text), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar6addTabERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:99
// index:0
// int insertTab(int, const class QString &)
func (this *QTabBar) InsertTab(index int, text unsafe.Pointer) {
	// 0: (, int index, const QString & text), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9insertTabEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:100
// index:1
// int insertTab(int, const class QIcon &, const class QString &)
func (this *QTabBar) InsertTab_1(index int, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, int index, const QIcon & icon, const QString & text), (&index, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9insertTabEiRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:102
// index:0
// void removeTab(int)
func (this *QTabBar) RemoveTab(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9removeTabEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:103
// index:0
// void moveTab(int, int)
func (this *QTabBar) MoveTab(from int, to int) {
	// 0: (, int from, int to), (&from, &to)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar7moveTabEii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:105
// index:0
// bool isTabEnabled(int)
func (this *QTabBar) IsTabEnabled(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12isTabEnabledEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:106
// index:0
// void setTabEnabled(int, _Bool)
func (this *QTabBar) SetTabEnabled(index int, arg1 bool) {
	// 0: (, int index, bool arg1), (&index, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13setTabEnabledEib", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:108
// index:0
// QString tabText(int)
func (this *QTabBar) TabText(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabTextEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:109
// index:0
// void setTabText(int, const class QString &)
func (this *QTabBar) SetTabText(index int, text unsafe.Pointer) {
	// 0: (, int index, const QString & text), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setTabTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:111
// index:0
// QColor tabTextColor(int)
func (this *QTabBar) TabTextColor(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12tabTextColorEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:112
// index:0
// void setTabTextColor(int, const class QColor &)
func (this *QTabBar) SetTabTextColor(index int, color unsafe.Pointer) {
	// 0: (, int index, const QColor & color), (&index, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setTabTextColorEiRK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, &index, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:114
// index:0
// QIcon tabIcon(int)
func (this *QTabBar) TabIcon(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabIconEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:115
// index:0
// void setTabIcon(int, const class QIcon &)
func (this *QTabBar) SetTabIcon(index int, icon unsafe.Pointer) {
	// 0: (, int index, const QIcon & icon), (&index, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setTabIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, &index, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:117
// index:0
// Qt::TextElideMode elideMode()
func (this *QTabBar) ElideMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9elideModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:118
// index:0
// void setElideMode(Qt::TextElideMode)
func (this *QTabBar) SetElideMode(arg0 int) {
	// 0: (, Qt::TextElideMode arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar12setElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:121
// index:0
// void setTabToolTip(int, const class QString &)
func (this *QTabBar) SetTabToolTip(index int, tip unsafe.Pointer) {
	// 0: (, int index, const QString & tip), (&index, tip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13setTabToolTipEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, tip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:122
// index:0
// QString tabToolTip(int)
func (this *QTabBar) TabToolTip(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar10tabToolTipEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:126
// index:0
// void setTabWhatsThis(int, const class QString &)
func (this *QTabBar) SetTabWhatsThis(index int, text unsafe.Pointer) {
	// 0: (, int index, const QString & text), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setTabWhatsThisEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:127
// index:0
// QString tabWhatsThis(int)
func (this *QTabBar) TabWhatsThis(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12tabWhatsThisEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:130
// index:0
// void setTabData(int, const class QVariant &)
func (this *QTabBar) SetTabData(index int, data unsafe.Pointer) {
	// 0: (, int index, const QVariant & data), (&index, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setTabDataEiRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &index, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:131
// index:0
// QVariant tabData(int)
func (this *QTabBar) TabData(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabDataEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:133
// index:0
// QRect tabRect(int)
func (this *QTabBar) TabRect(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabRectEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:134
// index:0
// int tabAt(const class QPoint &)
func (this *QTabBar) TabAt(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar5tabAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:136
// index:0
// int currentIndex()
func (this *QTabBar) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:137
// index:0
// int count()
func (this *QTabBar) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:139
// index:0
// virtual
// QSize sizeHint()
func (this *QTabBar) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:140
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QTabBar) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:142
// index:0
// void setDrawBase(_Bool)
func (this *QTabBar) SetDrawBase(drawTheBase bool) {
	// 0: (, bool drawTheBase), (&drawTheBase)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11setDrawBaseEb", ffiqt.FFI_TYPE_VOID, this.cthis, &drawTheBase)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:143
// index:0
// bool drawBase()
func (this *QTabBar) DrawBase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8drawBaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:145
// index:0
// QSize iconSize()
func (this *QTabBar) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:146
// index:0
// void setIconSize(const class QSize &)
func (this *QTabBar) SetIconSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:148
// index:0
// bool usesScrollButtons()
func (this *QTabBar) UsesScrollButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar17usesScrollButtonsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:149
// index:0
// void setUsesScrollButtons(_Bool)
func (this *QTabBar) SetUsesScrollButtons(useButtons bool) {
	// 0: (, bool useButtons), (&useButtons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar20setUsesScrollButtonsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &useButtons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:151
// index:0
// bool tabsClosable()
func (this *QTabBar) TabsClosable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12tabsClosableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:152
// index:0
// void setTabsClosable(_Bool)
func (this *QTabBar) SetTabsClosable(closable bool) {
	// 0: (, bool closable), (&closable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setTabsClosableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &closable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:154
// index:0
// void setTabButton(int, enum QTabBar::ButtonPosition, class QWidget *)
func (this *QTabBar) SetTabButton(index int, position int, widget unsafe.Pointer) {
	// 0: (, int index, QTabBar::ButtonPosition position, QWidget * widget), (&index, &position, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar12setTabButtonEiNS_14ButtonPositionEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &position, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:155
// index:0
// QWidget * tabButton(int, enum QTabBar::ButtonPosition)
func (this *QTabBar) TabButton(index int, position int) {
	// 0: (, int index, QTabBar::ButtonPosition position), (&index, &position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9tabButtonEiNS_14ButtonPositionE", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:157
// index:0
// QTabBar::SelectionBehavior selectionBehaviorOnRemove()
func (this *QTabBar) SelectionBehaviorOnRemove() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar25selectionBehaviorOnRemoveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:158
// index:0
// void setSelectionBehaviorOnRemove(enum QTabBar::SelectionBehavior)
func (this *QTabBar) SetSelectionBehaviorOnRemove(behavior int) {
	// 0: (, QTabBar::SelectionBehavior behavior), (&behavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar28setSelectionBehaviorOnRemoveENS_17SelectionBehaviorE", ffiqt.FFI_TYPE_VOID, this.cthis, &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:160
// index:0
// bool expanding()
func (this *QTabBar) Expanding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9expandingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:161
// index:0
// void setExpanding(_Bool)
func (this *QTabBar) SetExpanding(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar12setExpandingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:163
// index:0
// bool isMovable()
func (this *QTabBar) IsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9isMovableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:164
// index:0
// void setMovable(_Bool)
func (this *QTabBar) SetMovable(movable bool) {
	// 0: (, bool movable), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setMovableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:166
// index:0
// bool documentMode()
func (this *QTabBar) DocumentMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12documentModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:167
// index:0
// void setDocumentMode(_Bool)
func (this *QTabBar) SetDocumentMode(set bool) {
	// 0: (, bool set), (&set)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setDocumentModeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:169
// index:0
// bool autoHide()
func (this *QTabBar) AutoHide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8autoHideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:170
// index:0
// void setAutoHide(_Bool)
func (this *QTabBar) SetAutoHide(hide bool) {
	// 0: (, bool hide), (&hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11setAutoHideEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:172
// index:0
// bool changeCurrentOnDrag()
func (this *QTabBar) ChangeCurrentOnDrag() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar19changeCurrentOnDragEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:173
// index:0
// void setChangeCurrentOnDrag(_Bool)
func (this *QTabBar) SetChangeCurrentOnDrag(change bool) {
	// 0: (, bool change), (&change)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar22setChangeCurrentOnDragEb", ffiqt.FFI_TYPE_VOID, this.cthis, &change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:176
// index:0
// QString accessibleTabName(int)
func (this *QTabBar) AccessibleTabName(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar17accessibleTabNameEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:177
// index:0
// void setAccessibleTabName(int, const class QString &)
func (this *QTabBar) SetAccessibleTabName(index int, name unsafe.Pointer) {
	// 0: (, int index, const QString & name), (&index, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar20setAccessibleTabNameEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &index, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:181
// index:0
// void setCurrentIndex(int)
func (this *QTabBar) SetCurrentIndex(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:184
// index:0
// void currentChanged(int)
func (this *QTabBar) CurrentChanged(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar14currentChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:185
// index:0
// void tabCloseRequested(int)
func (this *QTabBar) TabCloseRequested(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar17tabCloseRequestedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:186
// index:0
// void tabMoved(int, int)
func (this *QTabBar) TabMoved(from int, to int) {
	// 0: (, int from, int to), (&from, &to)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar8tabMovedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:187
// index:0
// void tabBarClicked(int)
func (this *QTabBar) TabBarClicked(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13tabBarClickedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:188
// index:0
// void tabBarDoubleClicked(int)
func (this *QTabBar) TabBarDoubleClicked(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar19tabBarDoubleClickedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

//  body block end
