//  header block begin
// /usr/include/qt/QtWidgets/qaction.h
// #include <qaction.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QAction struct {
	*qtcore.QObject
}

func (this *QAction) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qaction.h:62
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAction) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// void QAction(class QObject *)
func NewQAction(parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}
func NewQActionFromPointer(cthis unsafe.Pointer) *QAction {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAction{bcthis0}
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// void QAction(const class QString &, class QObject *)
func NewQAction_1(text unsafe.Pointer, parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:97
// index:2
// void QAction(const class QIcon &, const class QString &, class QObject *)
func NewQAction_2(icon unsafe.Pointer, text unsafe.Pointer, parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2ERK5QIconRK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, icon, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:184
// index:3
// void QAction(class QActionPrivate &, class QObject *)
func NewQAction_3(dd unsafe.Pointer, parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2ER14QActionPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:99
// index:0
// virtual
// void ~QAction()
func DeleteQAction(*QAction) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:101
// index:0
// void setActionGroup(class QActionGroup *)
func (this *QAction) SetActionGroup(group unsafe.Pointer) {
	// 0: (, group QActionGroup *), (group)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction14setActionGroupEP12QActionGroup", ffiqt.FFI_TYPE_VOID, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:102
// index:0
// QActionGroup * actionGroup()
func (this *QAction) ActionGroup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction11actionGroupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:103
// index:0
// void setIcon(const class QIcon &)
func (this *QAction) SetIcon(icon unsafe.Pointer) {
	// 0: (, icon const QIcon &), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:104
// index:0
// QIcon icon()
func (this *QAction) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4iconEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:106
// index:0
// void setText(const class QString &)
func (this *QAction) SetText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:107
// index:0
// QString text()
func (this *QAction) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:109
// index:0
// void setIconText(const class QString &)
func (this *QAction) SetIconText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setIconTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:110
// index:0
// QString iconText()
func (this *QAction) IconText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8iconTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:112
// index:0
// void setToolTip(const class QString &)
func (this *QAction) SetToolTip(tip unsafe.Pointer) {
	// 0: (, tip const QString &), (tip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), tip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:113
// index:0
// QString toolTip()
func (this *QAction) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction7toolTipEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:115
// index:0
// void setStatusTip(const class QString &)
func (this *QAction) SetStatusTip(statusTip unsafe.Pointer) {
	// 0: (, statusTip const QString &), (statusTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setStatusTipERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), statusTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:116
// index:0
// QString statusTip()
func (this *QAction) StatusTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9statusTipEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:118
// index:0
// void setWhatsThis(const class QString &)
func (this *QAction) SetWhatsThis(what unsafe.Pointer) {
	// 0: (, what const QString &), (what)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setWhatsThisERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), what)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:119
// index:0
// QString whatsThis()
func (this *QAction) WhatsThis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9whatsThisEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:121
// index:0
// void setPriority(enum QAction::Priority)
func (this *QAction) SetPriority(priority int) {
	// 0: (, priority QAction::Priority), (&priority)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setPriorityENS_8PriorityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:122
// index:0
// QAction::Priority priority()
func (this *QAction) Priority() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8priorityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:125
// index:0
// QMenu * menu()
func (this *QAction) Menu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4menuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:126
// index:0
// void setMenu(class QMenu *)
func (this *QAction) SetMenu(menu unsafe.Pointer) {
	// 0: (, menu QMenu *), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:129
// index:0
// void setSeparator(_Bool)
func (this *QAction) SetSeparator(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setSeparatorEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:130
// index:0
// bool isSeparator()
func (this *QAction) IsSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction11isSeparatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:133
// index:0
// void setShortcut(const class QKeySequence &)
func (this *QAction) SetShortcut(shortcut unsafe.Pointer) {
	// 0: (, shortcut const QKeySequence &), (shortcut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setShortcutERK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.GetCthis(), shortcut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:134
// index:0
// QKeySequence shortcut()
func (this *QAction) Shortcut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8shortcutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:137
// index:0
// void setShortcuts(class QKeySequence::StandardKey)
func (this *QAction) SetShortcuts(arg0 int) {
	// 0: (, QKeySequence::StandardKey arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setShortcutsEN12QKeySequence11StandardKeyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:138
// index:0
// QList<QKeySequence> shortcuts()
func (this *QAction) Shortcuts() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9shortcutsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:140
// index:0
// void setShortcutContext(Qt::ShortcutContext)
func (this *QAction) SetShortcutContext(context int) {
	// 0: (, context Qt::ShortcutContext), (&context)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction18setShortcutContextEN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:141
// index:0
// Qt::ShortcutContext shortcutContext()
func (this *QAction) ShortcutContext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction15shortcutContextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:143
// index:0
// void setAutoRepeat(_Bool)
func (this *QAction) SetAutoRepeat(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction13setAutoRepeatEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:144
// index:0
// bool autoRepeat()
func (this *QAction) AutoRepeat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction10autoRepeatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:147
// index:0
// void setFont(const class QFont &)
func (this *QAction) SetFont(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:148
// index:0
// QFont font()
func (this *QAction) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4fontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:150
// index:0
// void setCheckable(_Bool)
func (this *QAction) SetCheckable(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setCheckableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:151
// index:0
// bool isCheckable()
func (this *QAction) IsCheckable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction11isCheckableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:153
// index:0
// QVariant data()
func (this *QAction) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:154
// index:0
// void setData(const class QVariant &)
func (this *QAction) SetData(var_ unsafe.Pointer) {
	// 0: (, var const QVariant &), (var_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setDataERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), var_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:156
// index:0
// bool isChecked()
func (this *QAction) IsChecked() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9isCheckedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:158
// index:0
// bool isEnabled()
func (this *QAction) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:160
// index:0
// bool isVisible()
func (this *QAction) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:163
// index:0
// void activate(enum QAction::ActionEvent)
func (this *QAction) Activate(event int) {
	// 0: (, event QAction::ActionEvent), (&event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction8activateENS_11ActionEventE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:164
// index:0
// bool showStatusText(class QWidget *)
func (this *QAction) ShowStatusText(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction14showStatusTextEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:166
// index:0
// void setMenuRole(enum QAction::MenuRole)
func (this *QAction) SetMenuRole(menuRole int) {
	// 0: (, menuRole QAction::MenuRole), (&menuRole)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setMenuRoleENS_8MenuRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &menuRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:167
// index:0
// QAction::MenuRole menuRole()
func (this *QAction) MenuRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8menuRoleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:169
// index:0
// void setIconVisibleInMenu(_Bool)
func (this *QAction) SetIconVisibleInMenu(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction20setIconVisibleInMenuEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:170
// index:0
// bool isIconVisibleInMenu()
func (this *QAction) IsIconVisibleInMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction19isIconVisibleInMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:172
// index:0
// void setShortcutVisibleInContextMenu(_Bool)
func (this *QAction) SetShortcutVisibleInContextMenu(show bool) {
	// 0: (, show bool), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction31setShortcutVisibleInContextMenuEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:173
// index:0
// bool isShortcutVisibleInContextMenu()
func (this *QAction) IsShortcutVisibleInContextMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction30isShortcutVisibleInContextMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:175
// index:0
// QWidget * parentWidget()
func (this *QAction) ParentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction12parentWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:177
// index:0
// QList<QWidget *> associatedWidgets()
func (this *QAction) AssociatedWidgets() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction17associatedWidgetsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:179
// index:0
// QList<QGraphicsWidget *> associatedGraphicsWidgets()
func (this *QAction) AssociatedGraphicsWidgets() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction25associatedGraphicsWidgetsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:183
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAction) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:187
// index:0
// inline
// void trigger()
func (this *QAction) Trigger() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7triggerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:188
// index:0
// inline
// void hover()
func (this *QAction) Hover() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction5hoverEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:189
// index:0
// void setChecked(_Bool)
func (this *QAction) SetChecked(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setCheckedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:190
// index:0
// void toggle()
func (this *QAction) Toggle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction6toggleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:191
// index:0
// void setEnabled(_Bool)
func (this *QAction) SetEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:192
// index:0
// inline
// void setDisabled(_Bool)
func (this *QAction) SetDisabled(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setDisabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:193
// index:0
// void setVisible(_Bool)
func (this *QAction) SetVisible(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:196
// index:0
// void changed()
func (this *QAction) Changed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7changedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// void triggered(_Bool)
func (this *QAction) Triggered(checked bool) {
	// 0: (, checked bool), (&checked)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction9triggeredEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:198
// index:0
// void hovered()
func (this *QAction) Hovered() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7hoveredEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:199
// index:0
// void toggled(_Bool)
func (this *QAction) Toggled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7toggledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
