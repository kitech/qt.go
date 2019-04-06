package qtwidgets

// /usr/include/qt/QtWidgets/qsystemtrayicon.h
// #include <qsystemtrayicon.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
func (this *QSystemTrayIcon) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QSystemTrayIcon struct {
	*qtcore.QObject
}
type QSystemTrayIcon_ITF interface {
	qtcore.QObject_ITF
	QSystemTrayIcon_PTR() *QSystemTrayIcon
}

func (ptr *QSystemTrayIcon) QSystemTrayIcon_PTR() *QSystemTrayIcon { return ptr }

func (this *QSystemTrayIcon) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSystemTrayIcon) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSystemTrayIconFromPointer(cthis unsafe.Pointer) *QSystemTrayIcon {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSystemTrayIcon{bcthis0}
}
func (*QSystemTrayIcon) NewFromPointer(cthis unsafe.Pointer) *QSystemTrayIcon {
	return NewQSystemTrayIconFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSystemTrayIcon) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(QObject *)

/*
Constructs a QSystemTrayIcon object with the given parent.

The icon is initially invisible.

See also visible.
*/
func (*QSystemTrayIcon) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QSystemTrayIcon {
	return NewQSystemTrayIcon(parent)
}
func NewQSystemTrayIcon(parent qtcore.QObject_ITF /*777 QObject **/) *QSystemTrayIcon {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSystemTrayIcon")
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(QObject *)

/*
Constructs a QSystemTrayIcon object with the given parent.

The icon is initially invisible.

See also visible.
*/
func (*QSystemTrayIcon) NewForInheritp() *QSystemTrayIcon {
	return NewQSystemTrayIconp()
}
func NewQSystemTrayIconp() *QSystemTrayIcon {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSystemTrayIcon")
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(const QIcon &, QObject *)

/*
Constructs a QSystemTrayIcon object with the given parent.

The icon is initially invisible.

See also visible.
*/
func (*QSystemTrayIcon) NewForInherit1(icon qtgui.QIcon_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QSystemTrayIcon {
	return NewQSystemTrayIcon1(icon, parent)
}
func NewQSystemTrayIcon1(icon qtgui.QIcon_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QSystemTrayIcon {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconC2ERK5QIconP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSystemTrayIcon")
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(const QIcon &, QObject *)

/*
Constructs a QSystemTrayIcon object with the given parent.

The icon is initially invisible.

See also visible.
*/
func (*QSystemTrayIcon) NewForInherit1p(icon qtgui.QIcon_ITF) *QSystemTrayIcon {
	return NewQSystemTrayIcon1p(icon)
}
func NewQSystemTrayIcon1p(icon qtgui.QIcon_ITF) *QSystemTrayIcon {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconC2ERK5QIconP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSystemTrayIcon")
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSystemTrayIcon()

/*

 */
func DeleteQSystemTrayIcon(this *QSystemTrayIcon) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QSystemTrayIcon) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QSystemTrayIcon) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QSystemTrayIcon) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QSystemTrayIcon) SetToolTip(tip string) {
	var tmpArg0 = qtcore.NewQString5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isSystemTrayAvailable()

/*
Returns true if the system tray is available; otherwise returns false.

If the system tray is currently unavailable but becomes available later, QSystemTrayIcon will automatically add an entry in the system tray if it is visible.
*/
func (this *QSystemTrayIcon) IsSystemTrayAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon21isSystemTrayAvailableEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	var nilthis *QSystemTrayIcon
	rv := nilthis.IsSystemTrayAvailable()
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:93
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool supportsMessages()

/*
Returns true if the system tray supports balloon messages; otherwise returns false.

See also showMessage().
*/
func (this *QSystemTrayIcon) SupportsMessages() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon16supportsMessagesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QSystemTrayIcon_SupportsMessages() bool {
	var nilthis *QSystemTrayIcon
	rv := nilthis.SupportsMessages()
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:97
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect geometry() const

/*
Returns the geometry of the system tray icon in screen coordinates.

This function was introduced in  Qt 4.3.

See also visible.
*/
func (this *QSystemTrayIcon) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QSystemTrayIcon) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QSystemTrayIcon) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void show()

/*
Shows the icon in the system tray.

See also hide() and visible.
*/
func (this *QSystemTrayIcon) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hide()

/*
Hides the system tray entry.

See also show() and visible.
*/
func (this *QSystemTrayIcon) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, const QIcon &, int)

/*
Shows a balloon message for the entry with the given title, message and icon for the time specified in millisecondsTimeoutHint. title and message must be plain text strings.

Message can be clicked by the user; the messageClicked() signal will emitted when this occurs.

Note that display of messages are dependent on the system configuration and user preferences, and that messages may not appear at all. Hence, it should not be relied upon as the sole means for providing critical information.

On Windows, the millisecondsTimeoutHint is usually ignored by the system when the application has focus.

Has been turned into a slot in Qt 5.2.

This function was introduced in  Qt 4.3.

See also show() and supportsMessages().
*/
func (this *QSystemTrayIcon) ShowMessage(title string, msg string, icon qtgui.QIcon_ITF, msecs int) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg2 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, const QIcon &, int)

/*
Shows a balloon message for the entry with the given title, message and icon for the time specified in millisecondsTimeoutHint. title and message must be plain text strings.

Message can be clicked by the user; the messageClicked() signal will emitted when this occurs.

Note that display of messages are dependent on the system configuration and user preferences, and that messages may not appear at all. Hence, it should not be relied upon as the sole means for providing critical information.

On Windows, the millisecondsTimeoutHint is usually ignored by the system when the application has focus.

Has been turned into a slot in Qt 5.2.

This function was introduced in  Qt 4.3.

See also show() and supportsMessages().
*/
func (this *QSystemTrayIcon) ShowMessagep(title string, msg string, icon qtgui.QIcon_ITF) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg2 = icon.QIcon_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	msecs := int(10000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)

/*
Shows a balloon message for the entry with the given title, message and icon for the time specified in millisecondsTimeoutHint. title and message must be plain text strings.

Message can be clicked by the user; the messageClicked() signal will emitted when this occurs.

Note that display of messages are dependent on the system configuration and user preferences, and that messages may not appear at all. Hence, it should not be relied upon as the sole means for providing critical information.

On Windows, the millisecondsTimeoutHint is usually ignored by the system when the application has focus.

Has been turned into a slot in Qt 5.2.

This function was introduced in  Qt 4.3.

See also show() and supportsMessages().
*/
func (this *QSystemTrayIcon) ShowMessage1(title string, msg string, icon int, msecs int) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, icon, msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)

/*
Shows a balloon message for the entry with the given title, message and icon for the time specified in millisecondsTimeoutHint. title and message must be plain text strings.

Message can be clicked by the user; the messageClicked() signal will emitted when this occurs.

Note that display of messages are dependent on the system configuration and user preferences, and that messages may not appear at all. Hence, it should not be relied upon as the sole means for providing critical information.

On Windows, the millisecondsTimeoutHint is usually ignored by the system when the application has focus.

Has been turned into a slot in Qt 5.2.

This function was introduced in  Qt 4.3.

See also show() and supportsMessages().
*/
func (this *QSystemTrayIcon) ShowMessage1p(title string, msg string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QSystemTrayIcon::MessageIcon=Elaborated, QSystemTrayIcon::MessageIcon=Enum, , Invalid
	icon := 0
	// arg: 3, int=Int, =Invalid, , Invalid
	msecs := int(10000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, icon, msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)

/*
Shows a balloon message for the entry with the given title, message and icon for the time specified in millisecondsTimeoutHint. title and message must be plain text strings.

Message can be clicked by the user; the messageClicked() signal will emitted when this occurs.

Note that display of messages are dependent on the system configuration and user preferences, and that messages may not appear at all. Hence, it should not be relied upon as the sole means for providing critical information.

On Windows, the millisecondsTimeoutHint is usually ignored by the system when the application has focus.

Has been turned into a slot in Qt 5.2.

This function was introduced in  Qt 4.3.

See also show() and supportsMessages().
*/
func (this *QSystemTrayIcon) ShowMessage1p1(title string, msg string, icon int) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, int=Int, =Invalid, , Invalid
	msecs := int(10000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, icon, msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(QSystemTrayIcon::ActivationReason)

/*
This signal is emitted when the user activates the system tray icon. reason specifies the reason for activation. QSystemTrayIcon::ActivationReason enumerates the various reasons.

See also QSystemTrayIcon::ActivationReason.
*/
func (this *QSystemTrayIcon) Activated(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon9activatedENS_16ActivationReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageClicked()

/*
This signal is emitted when the message displayed using showMessage() was clicked by the user.

Currently this signal is not sent on macOS.

Note: We follow Microsoft Windows behavior, so the signal is also emitted when the user clicks on a tray icon with a balloon message displayed.

See also activated().
*/
func (this *QSystemTrayIcon) MessageClicked() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon14messageClickedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSystemTrayIcon) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes the reason the system tray was activated.



Note: On macOS, a double click will only be emitted if no context menu is set, since the menu opens on mouse press



See also activated().

*/
type QSystemTrayIcon__ActivationReason = int

// Unknown reason
const QSystemTrayIcon__Unknown QSystemTrayIcon__ActivationReason = 0

// The context menu for the system tray entry was requested
const QSystemTrayIcon__Context QSystemTrayIcon__ActivationReason = 1

// The system tray entry was double clicked.
const QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = 2

// The system tray entry was clicked
const QSystemTrayIcon__Trigger QSystemTrayIcon__ActivationReason = 3

// The system tray entry was clicked with the middle mouse button
const QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = 4

func (this *QSystemTrayIcon) ActivationReasonItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSystemTrayIcon_ActivationReasonItemName(val int) string {
	var nilthis *QSystemTrayIcon
	return nilthis.ActivationReasonItemName(val)
}

/*
This enum describes the icon that is shown when a balloon message is displayed.



See also QMessageBox.

*/
type QSystemTrayIcon__MessageIcon = int

// No icon is shown.
const QSystemTrayIcon__NoIcon QSystemTrayIcon__MessageIcon = 0

// An information icon is shown.
const QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = 1

// A standard warning icon is shown.
const QSystemTrayIcon__Warning QSystemTrayIcon__MessageIcon = 2

// A critical warning icon is shown.
const QSystemTrayIcon__Critical QSystemTrayIcon__MessageIcon = 3

func (this *QSystemTrayIcon) MessageIconItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSystemTrayIcon_MessageIconItemName(val int) string {
	var nilthis *QSystemTrayIcon
	return nilthis.MessageIconItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11335() {
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
