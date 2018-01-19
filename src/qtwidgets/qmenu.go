//  header block begin
// /usr/include/qt/QtWidgets/qmenu.h
// #include <qmenu.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qmenu.h:64
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMenu) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// void QMenu(class QWidget *)
func NewQMenu(parent unsafe.Pointer) *QMenu {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QMenu{cthis}
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// void QMenu(const class QString &, class QWidget *)
func NewQMenu_1(title unsafe.Pointer, parent unsafe.Pointer) *QMenu {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenuC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, title, parent)
	gopp.ErrPrint(err, rv)
	return &QMenu{cthis}
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
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:80
// index:1
// QAction * addAction(const class QIcon &, const class QString &)
func (this *QMenu) AddAction_1(icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, const QIcon & icon, const QString & text), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:81
// index:2
// QAction * addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
func (this *QMenu) AddAction_2(text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer, shortcut unsafe.Pointer) {
	// 2: (, const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut), (text, receiver, member, shortcut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.cthis, text, receiver, member, shortcut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:82
// index:3
// QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
func (this *QMenu) AddAction_3(icon unsafe.Pointer, text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer, shortcut unsafe.Pointer) {
	// 3: (, const QIcon & icon, const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut), (icon, text, receiver, member, shortcut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.cthis, icon, text, receiver, member, shortcut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:156
// index:0
// QAction * addMenu(class QMenu *)
func (this *QMenu) AddMenu(menu unsafe.Pointer) {
	// 0: (, QMenu * menu), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:157
// index:1
// QMenu * addMenu(const class QString &)
func (this *QMenu) AddMenu_1(title unsafe.Pointer) {
	// 1: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:158
// index:2
// QMenu * addMenu(const class QIcon &, const class QString &)
func (this *QMenu) AddMenu_2(icon unsafe.Pointer, title unsafe.Pointer) {
	// 2: (, const QIcon & icon, const QString & title), (icon, title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7addMenuERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, icon, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:160
// index:0
// QAction * addSeparator()
func (this *QMenu) AddSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu12addSeparatorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:162
// index:0
// QAction * addSection(const class QString &)
func (this *QMenu) AddSection(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10addSectionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:163
// index:1
// QAction * addSection(const class QIcon &, const class QString &)
func (this *QMenu) AddSection_1(icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, const QIcon & icon, const QString & text), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10addSectionERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:165
// index:0
// QAction * insertMenu(class QAction *, class QMenu *)
func (this *QMenu) InsertMenu(before unsafe.Pointer, menu unsafe.Pointer) {
	// 0: (, QAction * before, QMenu * menu), (before, menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu10insertMenuEP7QActionPS_", ffiqt.FFI_TYPE_VOID, this.cthis, before, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:166
// index:0
// QAction * insertSeparator(class QAction *)
func (this *QMenu) InsertSeparator(before unsafe.Pointer) {
	// 0: (, QAction * before), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:167
// index:0
// QAction * insertSection(class QAction *, const class QString &)
func (this *QMenu) InsertSection(before unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, QAction * before, const QString & text), (before, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, before, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:168
// index:1
// QAction * insertSection(class QAction *, const class QIcon &, const class QString &)
func (this *QMenu) InsertSection_1(before unsafe.Pointer, icon unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, QAction * before, const QIcon & icon, const QString & text), (before, icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, before, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:170
// index:0
// bool isEmpty()
func (this *QMenu) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:171
// index:0
// void clear()
func (this *QMenu) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:173
// index:0
// void setTearOffEnabled(_Bool)
func (this *QMenu) SetTearOffEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu17setTearOffEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:174
// index:0
// bool isTearOffEnabled()
func (this *QMenu) IsTearOffEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu16isTearOffEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:176
// index:0
// bool isTearOffMenuVisible()
func (this *QMenu) IsTearOffMenuVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu20isTearOffMenuVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:177
// index:0
// void showTearOffMenu()
func (this *QMenu) ShowTearOffMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:178
// index:1
// void showTearOffMenu(const class QPoint &)
func (this *QMenu) ShowTearOffMenu_1(pos unsafe.Pointer) {
	// 1: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:179
// index:0
// void hideTearOffMenu()
func (this *QMenu) HideTearOffMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15hideTearOffMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:181
// index:0
// void setDefaultAction(class QAction *)
func (this *QMenu) SetDefaultAction(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu16setDefaultActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:182
// index:0
// QAction * defaultAction()
func (this *QMenu) DefaultAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu13defaultActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:184
// index:0
// void setActiveAction(class QAction *)
func (this *QMenu) SetActiveAction(act unsafe.Pointer) {
	// 0: (, QAction * act), (act)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15setActiveActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, act)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:185
// index:0
// QAction * activeAction()
func (this *QMenu) ActiveAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu12activeActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:187
// index:0
// void popup(const class QPoint &, class QAction *)
func (this *QMenu) Popup(pos unsafe.Pointer, at unsafe.Pointer) {
	// 0: (, const QPoint & pos, QAction * at), (pos, at)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu5popupERK6QPointP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, pos, at)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:188
// index:0
// QAction * exec()
func (this *QMenu) Exec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu4execEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:189
// index:1
// QAction * exec(const class QPoint &, class QAction *)
func (this *QMenu) Exec_1(pos unsafe.Pointer, at unsafe.Pointer) {
	// 1: (, const QPoint & pos, QAction * at), (pos, at)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu4execERK6QPointP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, pos, at)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:197
// index:0
// virtual
// QSize sizeHint()
func (this *QMenu) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:199
// index:0
// QRect actionGeometry(class QAction *)
func (this *QMenu) ActionGeometry(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu14actionGeometryEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:200
// index:0
// QAction * actionAt(const class QPoint &)
func (this *QMenu) ActionAt(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu8actionAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:202
// index:0
// QAction * menuAction()
func (this *QMenu) MenuAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu10menuActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:204
// index:0
// QString title()
func (this *QMenu) Title() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu5titleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:205
// index:0
// void setTitle(const class QString &)
func (this *QMenu) SetTitle(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu8setTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:207
// index:0
// QIcon icon()
func (this *QMenu) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu4iconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:208
// index:0
// void setIcon(const class QIcon &)
func (this *QMenu) SetIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:210
// index:0
// void setNoReplayFor(class QWidget *)
func (this *QMenu) SetNoReplayFor(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu14setNoReplayForEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:211
// index:0
// QPlatformMenu * platformMenu()
func (this *QMenu) PlatformMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu12platformMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:212
// index:0
// void setPlatformMenu(class QPlatformMenu *)
func (this *QMenu) SetPlatformMenu(platformMenu unsafe.Pointer) {
	// 0: (, QPlatformMenu * platformMenu), (platformMenu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu15setPlatformMenuEP13QPlatformMenu", ffiqt.FFI_TYPE_VOID, this.cthis, platformMenu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:219
// index:0
// bool separatorsCollapsible()
func (this *QMenu) SeparatorsCollapsible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu21separatorsCollapsibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:220
// index:0
// void setSeparatorsCollapsible(_Bool)
func (this *QMenu) SetSeparatorsCollapsible(collapse bool) {
	// 0: (, bool collapse), (&collapse)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu24setSeparatorsCollapsibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &collapse)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:222
// index:0
// bool toolTipsVisible()
func (this *QMenu) ToolTipsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QMenu15toolTipsVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:223
// index:0
// void setToolTipsVisible(_Bool)
func (this *QMenu) SetToolTipsVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu18setToolTipsVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:226
// index:0
// void aboutToShow()
func (this *QMenu) AboutToShow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11aboutToShowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:227
// index:0
// void aboutToHide()
func (this *QMenu) AboutToHide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu11aboutToHideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:228
// index:0
// void triggered(class QAction *)
func (this *QMenu) Triggered(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu9triggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:229
// index:0
// void hovered(class QAction *)
func (this *QMenu) Hovered(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QMenu7hoveredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

//  body block end
