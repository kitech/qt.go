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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QAction) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAction) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAction(QObject *)

/*
Constructs an action with parent. If parent is an action group the action will be automatically inserted into the group.

Note: The parent argument is optional since Qt 5.7.
*/
func (*QAction) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	return NewQAction(parent)
}
func NewQAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAction(QObject *)

/*
Constructs an action with parent. If parent is an action group the action will be automatically inserted into the group.

Note: The parent argument is optional since Qt 5.7.
*/
func (*QAction) NewForInheritp() *QAction {
	return NewQActionp()
}
func NewQActionp() *QAction {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QString &, QObject *)

/*
Constructs an action with parent. If parent is an action group the action will be automatically inserted into the group.

Note: The parent argument is optional since Qt 5.7.
*/
func (*QAction) NewForInherit1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	return NewQAction1(text, parent)
}
func NewQAction1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QString &, QObject *)

/*
Constructs an action with parent. If parent is an action group the action will be automatically inserted into the group.

Note: The parent argument is optional since Qt 5.7.
*/
func (*QAction) NewForInherit1p(text string) *QAction {
	return NewQAction1p(text)
}
func NewQAction1p(text string) *QAction {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:97
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QIcon &, const QString &, QObject *)

/*
Constructs an action with parent. If parent is an action group the action will be automatically inserted into the group.

Note: The parent argument is optional since Qt 5.7.
*/
func (*QAction) NewForInherit2(icon qtgui.QIcon_ITF, text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	return NewQAction2(icon, text, parent)
}
func NewQAction2(icon qtgui.QIcon_ITF, text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2ERK5QIconRK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:97
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QIcon &, const QString &, QObject *)

/*
Constructs an action with parent. If parent is an action group the action will be automatically inserted into the group.

Note: The parent argument is optional since Qt 5.7.
*/
func (*QAction) NewForInherit2p(icon qtgui.QIcon_ITF, text string) *QAction {
	return NewQAction2p(icon, text)
}
func NewQAction2p(icon qtgui.QIcon_ITF, text string) *QAction {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QActionC2ERK5QIconRK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAction()

/*

 */
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

/*
Sets this action group to group. The action will be automatically added to the group's list of actions.

Actions within the group will be mutually exclusive.

See also QActionGroup and QAction::actionGroup().
*/
func (this *QAction) SetActionGroup(group QActionGroup_ITF /*777 QActionGroup **/) {
	var convArg0 unsafe.Pointer
	if group != nil && group.QActionGroup_PTR() != nil {
		convArg0 = group.QActionGroup_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction14setActionGroupEP12QActionGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QActionGroup * actionGroup() const

/*
Returns the action group for this action. If no action group manages this action then 0 will be returned.

See also QActionGroup and QAction::setActionGroup().
*/
func (this *QAction) ActionGroup() *QActionGroup /*777 QActionGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction11actionGroupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QAction) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
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

/*

 */
func (this *QAction) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QAction) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconText(const QString &)

/*

 */
func (this *QAction) SetIconText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setIconTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString iconText() const

/*

 */
func (this *QAction) IconText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8iconTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QAction) SetToolTip(tip string) {
	var tmpArg0 = qtcore.NewQString5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QAction) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)

/*

 */
func (this *QAction) SetStatusTip(statusTip string) {
	var tmpArg0 = qtcore.NewQString5(statusTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString statusTip() const

/*

 */
func (this *QAction) StatusTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)

/*

 */
func (this *QAction) SetWhatsThis(what string) {
	var tmpArg0 = qtcore.NewQString5(what)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString whatsThis() const

/*

 */
func (this *QAction) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPriority(QAction::Priority)

/*

 */
func (this *QAction) SetPriority(priority int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setPriorityENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] QAction::Priority priority() const

/*

 */
func (this *QAction) Priority() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaction.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSeparator(bool)

/*
If b is true then this action will be considered a separator.

How a separator is represented depends on the widget it is inserted into. Under most circumstances the text, submenu, and icon will be ignored for separator actions.

See also QAction::isSeparator().
*/
func (this *QAction) SetSeparator(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setSeparatorEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSeparator() const

/*
Returns true if this action is a separator action; otherwise it returns false.

See also QAction::setSeparator().
*/
func (this *QAction) IsSeparator() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction11isSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcut(const QKeySequence &)

/*

 */
func (this *QAction) SetShortcut(shortcut qtgui.QKeySequence_ITF) {
	var convArg0 unsafe.Pointer
	if shortcut != nil && shortcut.QKeySequence_PTR() != nil {
		convArg0 = shortcut.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setShortcutERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence shortcut() const

/*
Returns the primary shortcut.

Note: Getter function for property shortcut.

See also setShortcuts().
*/
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

/*
Sets shortcuts as the list of shortcuts that trigger the action. The first element of the list is the primary shortcut.

This function was introduced in  Qt 4.2.

See also shortcuts() and shortcut.
*/
func (this *QAction) SetShortcuts(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setShortcutsEN12QKeySequence11StandardKeyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutContext(Qt::ShortcutContext)

/*

 */
func (this *QAction) SetShortcutContext(context int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction18setShortcutContextEN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), context)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ShortcutContext shortcutContext() const

/*

 */
func (this *QAction) ShortcutContext() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction15shortcutContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaction.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeat(bool)

/*

 */
func (this *QAction) SetAutoRepeat(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction13setAutoRepeatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRepeat() const

/*

 */
func (this *QAction) AutoRepeat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction10autoRepeatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QAction) SetFont(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:148
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
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
// [-2] void setCheckable(bool)

/*

 */
func (this *QAction) SetCheckable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction12setCheckableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*

 */
func (this *QAction) IsCheckable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction11isCheckableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:153
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data() const

/*
Returns the user data as set in QAction::setData.

See also setData().
*/
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

/*
Sets the action's internal data to the given userData.

See also data().
*/
func (this *QAction) SetData(var_ qtcore.QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if var_ != nil && var_.QVariant_PTR() != nil {
		convArg0 = var_.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7setDataERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isChecked() const

/*

 */
func (this *QAction) IsChecked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9isCheckedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:158
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*

 */
func (this *QAction) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QAction) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activate(QAction::ActionEvent)

/*
Sends the relevant signals for ActionEvent event.

Action based widgets use this API to cause the QAction to emit signals as well as emitting their own.
*/
func (this *QAction) Activate(event int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction8activateENS_11ActionEventE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), event)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showStatusText(QWidget *)

/*
Updates the relevant status bar for the widget specified by sending a QStatusTipEvent to its parent widget. Returns true if an event was sent; otherwise returns false.

If a null widget is specified, the event is sent to the action's parent.

See also statusTip.
*/
func (this *QAction) ShowStatusText(widget QWidget_ITF /*777 QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction14showStatusTextEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showStatusText(QWidget *)

/*
Updates the relevant status bar for the widget specified by sending a QStatusTipEvent to its parent widget. Returns true if an event was sent; otherwise returns false.

If a null widget is specified, the event is sent to the action's parent.

See also statusTip.
*/
func (this *QAction) ShowStatusTextp() bool {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction14showStatusTextEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuRole(QAction::MenuRole)

/*

 */
func (this *QAction) SetMenuRole(menuRole int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setMenuRoleENS_8MenuRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), menuRole)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] QAction::MenuRole menuRole() const

/*

 */
func (this *QAction) MenuRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction8menuRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaction.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconVisibleInMenu(bool)

/*

 */
func (this *QAction) SetIconVisibleInMenu(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction20setIconVisibleInMenuEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isIconVisibleInMenu() const

/*

 */
func (this *QAction) IsIconVisibleInMenu() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction19isIconVisibleInMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutVisibleInContextMenu(bool)

/*

 */
func (this *QAction) SetShortcutVisibleInContextMenu(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction31setShortcutVisibleInContextMenuEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isShortcutVisibleInContextMenu() const

/*

 */
func (this *QAction) IsShortcutVisibleInContextMenu() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction30isShortcutVisibleInContextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * parentWidget() const

/*
Returns the parent widget.
*/
func (this *QAction) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QAction12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaction.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QAction) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void trigger()

/*
This is a convenience slot that calls activate(Trigger).
*/
func (this *QAction) Trigger() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7triggerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hover()

/*
This is a convenience slot that calls activate(Hover).
*/
func (this *QAction) Hover() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction5hoverEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChecked(bool)

/*

 */
func (this *QAction) SetChecked(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setCheckedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggle()

/*
This is a convenience function for the checked property. Connect to it to change the checked state to its opposite state.
*/
func (this *QAction) Toggle() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction6toggleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QAction) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDisabled(bool)

/*
This is a convenience function for the enabled property, that is useful for signals--slots connections. If b is true the action is disabled; otherwise it is enabled.
*/
func (this *QAction) SetDisabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QAction) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed()

/*
This signal is emitted when an action has changed. If you are only interested in actions in a given widget, you can watch for QWidget::actionEvent() sent with an QEvent::ActionChanged.

Note: Notifier signal for property autoRepeat. Notifier signal for property checkable. Notifier signal for property enabled. Notifier signal for property font. Notifier signal for property icon. Notifier signal for property iconText. Notifier signal for property iconVisibleInMenu. Notifier signal for property menuRole. Notifier signal for property shortcut. Notifier signal for property shortcutContext. Notifier signal for property shortcutVisibleInContextMenu. Notifier signal for property statusTip. Notifier signal for property text. Notifier signal for property toolTip. Notifier signal for property visible. Notifier signal for property whatsThis.

See also QWidget::actionEvent().
*/
func (this *QAction) Changed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(bool)

/*
This signal is emitted when an action is activated by the user; for example, when the user clicks a menu option, toolbar button, or presses an action's shortcut key combination, or when trigger() was called. Notably, it is not emitted when setChecked() or toggle() is called.

If the action is checkable, checked is true if the action is checked, or false if the action is unchecked.

See also QAction::activate(), QAction::toggled(), and checked.
*/
func (this *QAction) Triggered(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction9triggeredEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(bool)

/*
This signal is emitted when an action is activated by the user; for example, when the user clicks a menu option, toolbar button, or presses an action's shortcut key combination, or when trigger() was called. Notably, it is not emitted when setChecked() or toggle() is called.

If the action is checkable, checked is true if the action is checked, or false if the action is unchecked.

See also QAction::activate(), QAction::toggled(), and checked.
*/
func (this *QAction) Triggeredp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction9triggeredEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered()

/*
This signal is emitted when an action is highlighted by the user; for example, when the user pauses with the cursor over a menu option, toolbar button, or presses an action's shortcut key combination.

See also QAction::activate().
*/
func (this *QAction) Hovered() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7hoveredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggled(bool)

/*
This signal is emitted whenever a checkable action changes its isChecked() status. This can be the result of a user interaction, or because setChecked() was called. As setChecked() changes the QAction, it emits changed() in addition to toggled().

checked is true if the action is checked, or false if the action is unchecked.

Note: Notifier signal for property checked.

See also QAction::activate(), QAction::triggered(), and checked.
*/
func (this *QAction) Toggled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QAction7toggledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes how an action should be moved into the application menu on macOS.



Setting this value only has effect on items that are in the immediate menus of the menubar, not the submenus of those menus. For example, if you have File menu in your menubar and the File menu has a submenu, setting the MenuRole for the actions in that submenu have no effect. They will never be moved.

*/
type QAction__MenuRole = int

// This action should not be put into the application menu
const QAction__NoRole QAction__MenuRole = 0

// This action should be put in the application menu based on the action's text as described in the QMenuBar documentation.
const QAction__TextHeuristicRole QAction__MenuRole = 1

// This action should be put in the application menu with an application specific role
const QAction__ApplicationSpecificRole QAction__MenuRole = 2

// This action handles the "About Qt" menu item.
const QAction__AboutQtRole QAction__MenuRole = 3

// This action should be placed where the "About" menu item is in the application menu. The text of the menu item will be set to "About <application name>". The application name is fetched from the Info.plist file in the application's bundle (See Qt for macOS - Deployment).
const QAction__AboutRole QAction__MenuRole = 4

// This action should be placed where the "Preferences..." menu item is in the application menu.
const QAction__PreferencesRole QAction__MenuRole = 5

// This action should be placed where the Quit menu item is in the application menu.
const QAction__QuitRole QAction__MenuRole = 6

func (this *QAction) MenuRoleItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAction_MenuRoleItemName(val int) string {
	var nilthis *QAction
	return nilthis.MenuRoleItemName(val)
}

/*
This enum defines priorities for actions in user interface.



This enum was introduced or modified in  Qt 4.6.

See also priority.

*/
type QAction__Priority = int

// The action should not be prioritized in the user interface.
const QAction__LowPriority QAction__Priority = 0

//
const QAction__NormalPriority QAction__Priority = 128

//
const QAction__HighPriority QAction__Priority = 256

func (this *QAction) PriorityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAction_PriorityItemName(val int) string {
	var nilthis *QAction
	return nilthis.PriorityItemName(val)
}

/*
This enum type is used when calling QAction::activate()


*/
type QAction__ActionEvent = int

// this will cause the QAction::triggered() signal to be emitted.
const QAction__Trigger QAction__ActionEvent = 0

// this will cause the QAction::hovered() signal to be emitted.
const QAction__Hover QAction__ActionEvent = 1

func (this *QAction) ActionEventItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAction_ActionEventItemName(val int) string {
	var nilthis *QAction
	return nilthis.ActionEventItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11065() {
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
