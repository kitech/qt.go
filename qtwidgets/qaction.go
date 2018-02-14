package qtwidgets

// /usr/include/qt/QtWidgets/qaction.h
// #include <qaction.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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

// bool event(class QEvent *)
func (this *QAction) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QAction struct {
	*qtcore.QObject
}
type QAction_ITF interface {
	qtcore.QObject_ITF
	QAction_PTR() *QAction
}

func (ptr *QAction) QAction_PTR() *QAction { return ptr }

func (this *QAction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAction) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQActionFromPointer(cthis unsafe.Pointer) *QAction {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAction{bcthis0}
}
func (*QAction) NewFromPointer(cthis unsafe.Pointer) *QAction {
	return NewQActionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qaction.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAction) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAction(QObject *)
func NewQAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QString &, QObject *)
func NewQAction_1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:97
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QIcon &, const QString &, QObject *)
func NewQAction_2(icon qtgui.QIcon_ITF, text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var convArg0 = icon.QIcon_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2ERK5QIconRK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAction()
func DeleteQAction(this *QAction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qaction.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActionGroup(QActionGroup *)
func (this *QAction) SetActionGroup(group QActionGroup_ITF /*777 QActionGroup **/) {
	var convArg0 = group.QActionGroup_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction14setActionGroupEP12QActionGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QActionGroup * actionGroup()
func (this *QAction) ActionGroup() *QActionGroup /*777 QActionGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction11actionGroupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QAction) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 = icon.QIcon_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QAction) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qaction.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QAction) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QAction) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconText(const QString &)
func (this *QAction) SetIconText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setIconTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString iconText()
func (this *QAction) IconText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8iconTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QAction) SetToolTip(tip string) {
	var tmpArg0 = qtcore.NewQString_5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QAction) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)
func (this *QAction) SetStatusTip(statusTip string) {
	var tmpArg0 = qtcore.NewQString_5(statusTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString statusTip()
func (this *QAction) StatusTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QAction) SetWhatsThis(what string) {
	var tmpArg0 = qtcore.NewQString_5(what)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QAction) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPriority(enum QAction::Priority)
func (this *QAction) SetPriority(priority int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setPriorityENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] QAction::Priority priority()
func (this *QAction) Priority() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaction.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu()
func (this *QAction) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)
func (this *QAction) SetMenu(menu QMenu_ITF /*777 QMenu **/) {
	var convArg0 = menu.QMenu_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSeparator(_Bool)
func (this *QAction) SetSeparator(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setSeparatorEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSeparator()
func (this *QAction) IsSeparator() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction11isSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcut(const QKeySequence &)
func (this *QAction) SetShortcut(shortcut qtgui.QKeySequence_ITF) {
	var convArg0 = shortcut.QKeySequence_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setShortcutERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence shortcut()
func (this *QAction) Shortcut() *qtgui.QKeySequence /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8shortcutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtWidgets/qaction.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcuts(QKeySequence::StandardKey)
func (this *QAction) SetShortcuts(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setShortcutsEN12QKeySequence11StandardKeyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutContext(Qt::ShortcutContext)
func (this *QAction) SetShortcutContext(context int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction18setShortcutContextEN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), context)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ShortcutContext shortcutContext()
func (this *QAction) ShortcutContext() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction15shortcutContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaction.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeat(_Bool)
func (this *QAction) SetAutoRepeat(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction13setAutoRepeatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRepeat()
func (this *QAction) AutoRepeat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction10autoRepeatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QAction) SetFont(font qtgui.QFont_ITF) {
	var convArg0 = font.QFont_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:148
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QAction) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qaction.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(_Bool)
func (this *QAction) SetCheckable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setCheckableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCheckable()
func (this *QAction) IsCheckable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction11isCheckableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:153
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data()
func (this *QAction) Data() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qaction.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QVariant &)
func (this *QAction) SetData(var_ qtcore.QVariant_ITF) {
	var convArg0 = var_.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setDataERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isChecked()
func (this *QAction) IsChecked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9isCheckedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:158
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QAction) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QAction) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activate(enum QAction::ActionEvent)
func (this *QAction) Activate(event int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction8activateENS_11ActionEventE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), event)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showStatusText(QWidget *)
func (this *QAction) ShowStatusText(widget QWidget_ITF /*777 QWidget **/) bool {
	var convArg0 = widget.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction14showStatusTextEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuRole(enum QAction::MenuRole)
func (this *QAction) SetMenuRole(menuRole int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setMenuRoleENS_8MenuRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), menuRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] QAction::MenuRole menuRole()
func (this *QAction) MenuRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8menuRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaction.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconVisibleInMenu(_Bool)
func (this *QAction) SetIconVisibleInMenu(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction20setIconVisibleInMenuEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isIconVisibleInMenu()
func (this *QAction) IsIconVisibleInMenu() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction19isIconVisibleInMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutVisibleInContextMenu(_Bool)
func (this *QAction) SetShortcutVisibleInContextMenu(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction31setShortcutVisibleInContextMenuEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isShortcutVisibleInContextMenu()
func (this *QAction) IsShortcutVisibleInContextMenu() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction30isShortcutVisibleInContextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * parentWidget()
func (this *QAction) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAction) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = arg0.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void trigger()
func (this *QAction) Trigger() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7triggerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hover()
func (this *QAction) Hover() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction5hoverEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChecked(_Bool)
func (this *QAction) SetChecked(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setCheckedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggle()
func (this *QAction) Toggle() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction6toggleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QAction) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDisabled(_Bool)
func (this *QAction) SetDisabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QAction) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed()
func (this *QAction) Changed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(_Bool)
func (this *QAction) Triggered(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction9triggeredEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered()
func (this *QAction) Hovered() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7hoveredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggled(_Bool)
func (this *QAction) Toggled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7toggledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

type QAction__MenuRole = int

const QAction__NoRole QAction__MenuRole = 0
const QAction__TextHeuristicRole QAction__MenuRole = 1
const QAction__ApplicationSpecificRole QAction__MenuRole = 2
const QAction__AboutQtRole QAction__MenuRole = 3
const QAction__AboutRole QAction__MenuRole = 4
const QAction__PreferencesRole QAction__MenuRole = 5
const QAction__QuitRole QAction__MenuRole = 6

type QAction__Priority = int

const QAction__LowPriority QAction__Priority = 0
const QAction__NormalPriority QAction__Priority = 128
const QAction__HighPriority QAction__Priority = 256

type QAction__ActionEvent = int

const QAction__Trigger QAction__ActionEvent = 0
const QAction__Hover QAction__ActionEvent = 1

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
