//  header block begin
// /usr/include/qt/QtWidgets/qmenu.h
// #include <qmenu.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
type QMenu struct {
	*QWidget
}

func (this *QMenu) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qmenu.h:64
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMenu) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// void QMenu(class QWidget *)
func NewQMenu(parent unsafe.Pointer) *QMenu {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(cthis)
	return gothis
}
func NewQMenuFromPointer(cthis unsafe.Pointer) *QMenu {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMenu{bcthis0}
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// void QMenu(const class QString &, class QWidget *)
func NewQMenu_1(title unsafe.Pointer, parent unsafe.Pointer) *QMenu {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, title, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:262
// index:2
// void QMenu(class QMenuPrivate &, class QWidget *)
func NewQMenu_2(dd unsafe.Pointer, parent unsafe.Pointer) *QMenu {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2ER12QMenuPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:76
// index:0
// virtual
// void ~QMenu()
func DeleteQMenu(*QMenu) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:79
// index:0
// QAction * addAction(const class QString &)
func (this *QMenu) AddAction(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:80
// index:1
// QAction * addAction(const class QIcon &, const class QString &)
func (this *QMenu) AddAction_1(icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, icon const QIcon &, text const QString &), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:81
// index:2
// QAction * addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
func (this *QMenu) AddAction_2(text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer, shortcut unsafe.Pointer) {
	// 2: (, text const QString &, receiver const QObject *, member const char *, shortcut const QKeySequence &), (text, receiver, member, shortcut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, receiver, member, shortcut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:82
// index:3
// QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
func (this *QMenu) AddAction_3(icon unsafe.Pointer, text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer, shortcut unsafe.Pointer) {
	// 3: (, icon const QIcon &, text const QString &, receiver const QObject *, member const char *, shortcut const QKeySequence &), (icon, text, receiver, member, shortcut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, text, receiver, member, shortcut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:156
// index:0
// QAction * addMenu(class QMenu *)
func (this *QMenu) AddMenu(menu unsafe.Pointer) {
	// 0: (, menu QMenu *), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:157
// index:1
// QMenu * addMenu(const class QString &)
func (this *QMenu) AddMenu_1(title unsafe.Pointer) {
	// 1: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:158
// index:2
// QMenu * addMenu(const class QIcon &, const class QString &)
func (this *QMenu) AddMenu_2(icon unsafe.Pointer, title unsafe.Pointer) {
	// 2: (, icon const QIcon &, title const QString &), (icon, title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:160
// index:0
// QAction * addSeparator()
func (this *QMenu) AddSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu12addSeparatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:162
// index:0
// QAction * addSection(const class QString &)
func (this *QMenu) AddSection(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10addSectionERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:163
// index:1
// QAction * addSection(const class QIcon &, const class QString &)
func (this *QMenu) AddSection_1(icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, icon const QIcon &, text const QString &), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10addSectionERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:165
// index:0
// QAction * insertMenu(class QAction *, class QMenu *)
func (this *QMenu) InsertMenu(before unsafe.Pointer, menu unsafe.Pointer) {
	// 0: (, before QAction *, menu QMenu *), (before, menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10insertMenuEP7QActionPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:166
// index:0
// QAction * insertSeparator(class QAction *)
func (this *QMenu) InsertSeparator(before unsafe.Pointer) {
	// 0: (, before QAction *), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:167
// index:0
// QAction * insertSection(class QAction *, const class QString &)
func (this *QMenu) InsertSection(before unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, before QAction *, text const QString &), (before, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:168
// index:1
// QAction * insertSection(class QAction *, const class QIcon &, const class QString &)
func (this *QMenu) InsertSection_1(before unsafe.Pointer, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, before QAction *, icon const QIcon &, text const QString &), (before, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:170
// index:0
// bool isEmpty()
func (this *QMenu) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:171
// index:0
// void clear()
func (this *QMenu) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:173
// index:0
// void setTearOffEnabled(_Bool)
func (this *QMenu) SetTearOffEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu17setTearOffEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:174
// index:0
// bool isTearOffEnabled()
func (this *QMenu) IsTearOffEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu16isTearOffEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:176
// index:0
// bool isTearOffMenuVisible()
func (this *QMenu) IsTearOffMenuVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu20isTearOffMenuVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:177
// index:0
// void showTearOffMenu()
func (this *QMenu) ShowTearOffMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:178
// index:1
// void showTearOffMenu(const class QPoint &)
func (this *QMenu) ShowTearOffMenu_1(pos unsafe.Pointer) {
	// 1: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:179
// index:0
// void hideTearOffMenu()
func (this *QMenu) HideTearOffMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15hideTearOffMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:181
// index:0
// void setDefaultAction(class QAction *)
func (this *QMenu) SetDefaultAction(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu16setDefaultActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:182
// index:0
// QAction * defaultAction()
func (this *QMenu) DefaultAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu13defaultActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:184
// index:0
// void setActiveAction(class QAction *)
func (this *QMenu) SetActiveAction(act unsafe.Pointer) {
	// 0: (, act QAction *), (act)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15setActiveActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), act)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:185
// index:0
// QAction * activeAction()
func (this *QMenu) ActiveAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu12activeActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:187
// index:0
// void popup(const class QPoint &, class QAction *)
func (this *QMenu) Popup(pos unsafe.Pointer, at unsafe.Pointer) {
	// 0: (, pos const QPoint &, at QAction *), (pos, at)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5popupERK6QPointP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos, at)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:188
// index:0
// QAction * exec()
func (this *QMenu) Exec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu4execEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:189
// index:1
// QAction * exec(const class QPoint &, class QAction *)
func (this *QMenu) Exec_1(pos unsafe.Pointer, at unsafe.Pointer) {
	// 1: (, pos const QPoint &, at QAction *), (pos, at)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu4execERK6QPointP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos, at)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:197
// index:0
// virtual
// QSize sizeHint()
func (this *QMenu) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:199
// index:0
// QRect actionGeometry(class QAction *)
func (this *QMenu) ActionGeometry(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu14actionGeometryEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:200
// index:0
// QAction * actionAt(const class QPoint &)
func (this *QMenu) ActionAt(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu8actionAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:202
// index:0
// QAction * menuAction()
func (this *QMenu) MenuAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu10menuActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:204
// index:0
// QString title()
func (this *QMenu) Title() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu5titleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:205
// index:0
// void setTitle(const class QString &)
func (this *QMenu) SetTitle(title unsafe.Pointer) {
	// 0: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu8setTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:207
// index:0
// QIcon icon()
func (this *QMenu) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu4iconEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:208
// index:0
// void setIcon(const class QIcon &)
func (this *QMenu) SetIcon(icon unsafe.Pointer) {
	// 0: (, icon const QIcon &), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:210
// index:0
// void setNoReplayFor(class QWidget *)
func (this *QMenu) SetNoReplayFor(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu14setNoReplayForEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:211
// index:0
// QPlatformMenu * platformMenu()
func (this *QMenu) PlatformMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu12platformMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:212
// index:0
// void setPlatformMenu(class QPlatformMenu *)
func (this *QMenu) SetPlatformMenu(platformMenu unsafe.Pointer) {
	// 0: (, platformMenu QPlatformMenu *), (platformMenu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15setPlatformMenuEP13QPlatformMenu", ffiqt.FFI_TYPE_VOID, this.GetCthis(), platformMenu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:219
// index:0
// bool separatorsCollapsible()
func (this *QMenu) SeparatorsCollapsible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu21separatorsCollapsibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:220
// index:0
// void setSeparatorsCollapsible(_Bool)
func (this *QMenu) SetSeparatorsCollapsible(collapse bool) {
	// 0: (, collapse bool), (&collapse)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu24setSeparatorsCollapsibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &collapse)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:222
// index:0
// bool toolTipsVisible()
func (this *QMenu) ToolTipsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu15toolTipsVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:223
// index:0
// void setToolTipsVisible(_Bool)
func (this *QMenu) SetToolTipsVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu18setToolTipsVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:226
// index:0
// void aboutToShow()
func (this *QMenu) AboutToShow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11aboutToShowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:227
// index:0
// void aboutToHide()
func (this *QMenu) AboutToHide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11aboutToHideEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:228
// index:0
// void triggered(class QAction *)
func (this *QMenu) Triggered(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9triggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:229
// index:0
// void hovered(class QAction *)
func (this *QMenu) Hovered(action unsafe.Pointer) {
	// 0: (, action QAction *), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7hoveredEP7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:232
// index:0
// int columnCount()
func (this *QMenu) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu11columnCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:234
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QMenu) ChangeEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:235
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QMenu) KeyPressEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:236
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QMenu) MouseReleaseEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:237
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QMenu) MousePressEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:238
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QMenu) MouseMoveEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:240
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QMenu) WheelEvent(arg0 unsafe.Pointer) {
	// 0: (, QWheelEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:242
// index:0
// virtual
// void enterEvent(class QEvent *)
func (this *QMenu) EnterEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10enterEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:243
// index:0
// virtual
// void leaveEvent(class QEvent *)
func (this *QMenu) LeaveEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10leaveEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:244
// index:0
// virtual
// void hideEvent(class QHideEvent *)
func (this *QMenu) HideEvent(arg0 unsafe.Pointer) {
	// 0: (, QHideEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:245
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QMenu) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:246
// index:0
// virtual
// void actionEvent(class QActionEvent *)
func (this *QMenu) ActionEvent(arg0 unsafe.Pointer) {
	// 0: (, QActionEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:247
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QMenu) TimerEvent(arg0 unsafe.Pointer) {
	// 0: (, QTimerEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:248
// index:0
// virtual
// bool event(class QEvent *)
func (this *QMenu) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:249
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QMenu) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:250
// index:0
// void initStyleOption(class QStyleOptionMenuItem *, const class QAction *)
func (this *QMenu) InitStyleOption(option unsafe.Pointer, action unsafe.Pointer) {
	// 0: (, option QStyleOptionMenuItem *, action const QAction *), (option, action)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, action)
	gopp.ErrPrint(err, rv)
}

//  body block end
