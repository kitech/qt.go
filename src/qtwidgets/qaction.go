//  header block begin
// /usr/include/qt/QtWidgets/qaction.h
// #include <qaction.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
func NewQActionFromPointer(cthis unsafe.Pointer) *QAction {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAction{bcthis0}
}

// /usr/include/qt/QtWidgets/qaction.h:62
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAction) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// Public
// void QAction(class QObject *)
func NewQAction(parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// Public
// void QAction(const class QString &, class QObject *)
func NewQAction_1(text *qtcore.QString, parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:97
// index:2
// Public
// void QAction(const class QIcon &, const class QString &, class QObject *)
func NewQAction_2(icon *qtgui.QIcon, text *qtcore.QString, parent unsafe.Pointer) *QAction {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionC2ERK5QIconRK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:99
// index:0
// Public virtual
// void ~QAction()
func DeleteQAction(*QAction) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QActionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:101
// index:0
// Public
// void setActionGroup(class QActionGroup *)
func (this *QAction) SetActionGroup(group unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction14setActionGroupEP12QActionGroup", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:102
// index:0
// Public
// QActionGroup * actionGroup()
func (this *QAction) ActionGroup() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction11actionGroupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:103
// index:0
// Public
// void setIcon(const class QIcon &)
func (this *QAction) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:104
// index:0
// Public
// QIcon icon()
func (this *QAction) Icon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:106
// index:0
// Public
// void setText(const class QString &)
func (this *QAction) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:107
// index:0
// Public
// QString text()
func (this *QAction) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:109
// index:0
// Public
// void setIconText(const class QString &)
func (this *QAction) SetIconText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setIconTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:110
// index:0
// Public
// QString iconText()
func (this *QAction) IconText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8iconTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:112
// index:0
// Public
// void setToolTip(const class QString &)
func (this *QAction) SetToolTip(tip *qtcore.QString) {
	var convArg0 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:113
// index:0
// Public
// QString toolTip()
func (this *QAction) ToolTip() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:115
// index:0
// Public
// void setStatusTip(const class QString &)
func (this *QAction) SetStatusTip(statusTip *qtcore.QString) {
	var convArg0 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:116
// index:0
// Public
// QString statusTip()
func (this *QAction) StatusTip() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9statusTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:118
// index:0
// Public
// void setWhatsThis(const class QString &)
func (this *QAction) SetWhatsThis(what *qtcore.QString) {
	var convArg0 = what.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:119
// index:0
// Public
// QString whatsThis()
func (this *QAction) WhatsThis() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:121
// index:0
// Public
// void setPriority(enum QAction::Priority)
func (this *QAction) SetPriority(priority int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setPriorityENS_8PriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:122
// index:0
// Public
// QAction::Priority priority()
func (this *QAction) Priority() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8priorityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:125
// index:0
// Public
// QMenu * menu()
func (this *QAction) Menu() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4menuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:126
// index:0
// Public
// void setMenu(class QMenu *)
func (this *QAction) SetMenu(menu unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:129
// index:0
// Public
// void setSeparator(_Bool)
func (this *QAction) SetSeparator(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setSeparatorEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:130
// index:0
// Public
// bool isSeparator()
func (this *QAction) IsSeparator() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction11isSeparatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:133
// index:0
// Public
// void setShortcut(const class QKeySequence &)
func (this *QAction) SetShortcut(shortcut *qtgui.QKeySequence) {
	var convArg0 = shortcut.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setShortcutERK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:134
// index:0
// Public
// QKeySequence shortcut()
func (this *QAction) Shortcut() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8shortcutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:137
// index:0
// Public
// void setShortcuts(class QKeySequence::StandardKey)
func (this *QAction) SetShortcuts(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setShortcutsEN12QKeySequence11StandardKeyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:138
// index:0
// Public
// QList<QKeySequence> shortcuts()
func (this *QAction) Shortcuts() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9shortcutsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:140
// index:0
// Public
// void setShortcutContext(Qt::ShortcutContext)
func (this *QAction) SetShortcutContext(context int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction18setShortcutContextEN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:141
// index:0
// Public
// Qt::ShortcutContext shortcutContext()
func (this *QAction) ShortcutContext() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction15shortcutContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:143
// index:0
// Public
// void setAutoRepeat(_Bool)
func (this *QAction) SetAutoRepeat(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction13setAutoRepeatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:144
// index:0
// Public
// bool autoRepeat()
func (this *QAction) AutoRepeat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction10autoRepeatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:147
// index:0
// Public
// void setFont(const class QFont &)
func (this *QAction) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:148
// index:0
// Public
// QFont font()
func (this *QAction) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:150
// index:0
// Public
// void setCheckable(_Bool)
func (this *QAction) SetCheckable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction12setCheckableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:151
// index:0
// Public
// bool isCheckable()
func (this *QAction) IsCheckable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction11isCheckableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:153
// index:0
// Public
// QVariant data()
func (this *QAction) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:154
// index:0
// Public
// void setData(const class QVariant &)
func (this *QAction) SetData(var_ *qtcore.QVariant) {
	var convArg0 = var_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7setDataERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:156
// index:0
// Public
// bool isChecked()
func (this *QAction) IsChecked() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9isCheckedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:158
// index:0
// Public
// bool isEnabled()
func (this *QAction) IsEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:160
// index:0
// Public
// bool isVisible()
func (this *QAction) IsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:163
// index:0
// Public
// void activate(enum QAction::ActionEvent)
func (this *QAction) Activate(event int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction8activateENS_11ActionEventE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:164
// index:0
// Public
// bool showStatusText(class QWidget *)
func (this *QAction) ShowStatusText(widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction14showStatusTextEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:166
// index:0
// Public
// void setMenuRole(enum QAction::MenuRole)
func (this *QAction) SetMenuRole(menuRole int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setMenuRoleENS_8MenuRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &menuRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:167
// index:0
// Public
// QAction::MenuRole menuRole()
func (this *QAction) MenuRole() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction8menuRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:169
// index:0
// Public
// void setIconVisibleInMenu(_Bool)
func (this *QAction) SetIconVisibleInMenu(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction20setIconVisibleInMenuEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:170
// index:0
// Public
// bool isIconVisibleInMenu()
func (this *QAction) IsIconVisibleInMenu() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction19isIconVisibleInMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:172
// index:0
// Public
// void setShortcutVisibleInContextMenu(_Bool)
func (this *QAction) SetShortcutVisibleInContextMenu(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction31setShortcutVisibleInContextMenuEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:173
// index:0
// Public
// bool isShortcutVisibleInContextMenu()
func (this *QAction) IsShortcutVisibleInContextMenu() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction30isShortcutVisibleInContextMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:175
// index:0
// Public
// QWidget * parentWidget()
func (this *QAction) ParentWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction12parentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:177
// index:0
// Public
// QList<QWidget *> associatedWidgets()
func (this *QAction) AssociatedWidgets() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction17associatedWidgetsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:179
// index:0
// Public
// QList<QGraphicsWidget *> associatedGraphicsWidgets()
func (this *QAction) AssociatedGraphicsWidgets() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QAction25associatedGraphicsWidgetsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:183
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAction) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaction.h:187
// index:0
// Public inline
// void trigger()
func (this *QAction) Trigger() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7triggerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:188
// index:0
// Public inline
// void hover()
func (this *QAction) Hover() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction5hoverEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:189
// index:0
// Public
// void setChecked(_Bool)
func (this *QAction) SetChecked(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setCheckedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:190
// index:0
// Public
// void toggle()
func (this *QAction) Toggle() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction6toggleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:191
// index:0
// Public
// void setEnabled(_Bool)
func (this *QAction) SetEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:192
// index:0
// Public inline
// void setDisabled(_Bool)
func (this *QAction) SetDisabled(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction11setDisabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:193
// index:0
// Public
// void setVisible(_Bool)
func (this *QAction) SetVisible(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:196
// index:0
// Public
// void changed()
func (this *QAction) Changed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7changedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// Public
// void triggered(_Bool)
func (this *QAction) Triggered(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction9triggeredEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:198
// index:0
// Public
// void hovered()
func (this *QAction) Hovered() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7hoveredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:199
// index:0
// Public
// void toggled(_Bool)
func (this *QAction) Toggled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QAction7toggledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
