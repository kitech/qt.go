package qtwidgets

// /usr/include/qt/QtWidgets/qsystemtrayicon.h
// #include <qsystemtrayicon.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// bool event(class QEvent *)
func (this *QSystemTrayIcon) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QSystemTrayIcon struct {
	*qtcore.QObject
}

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
// [8] const QMetaObject * metaObject()
func (this *QSystemTrayIcon) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(QObject *)
func NewQSystemTrayIcon(parent *qtcore.QObject /*777 QObject **/) *QSystemTrayIcon {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(const QIcon &, QObject *)
func NewQSystemTrayIcon_1(icon *qtgui.QIcon, parent *qtcore.QObject /*777 QObject **/) *QSystemTrayIcon {
	var convArg0 = icon.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconC2ERK5QIconP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSystemTrayIcon()
func DeleteQSystemTrayIcon(this *QSystemTrayIcon) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIconD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextMenu(QMenu *)
func (this *QSystemTrayIcon) SetContextMenu(menu *QMenu /*777 QMenu **/) {
	var convArg0 = menu.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon14setContextMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * contextMenu()
func (this *QSystemTrayIcon) ContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon11contextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QSystemTrayIcon) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QSystemTrayIcon) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QSystemTrayIcon) ToolTip() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QSystemTrayIcon) SetToolTip(tip *qtcore.QString) {
	var convArg0 = tip.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isSystemTrayAvailable()
func (this *QSystemTrayIcon) IsSystemTrayAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon21isSystemTrayAvailableEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
func (this *QSystemTrayIcon) SupportsMessages() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon16supportsMessagesEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
// [16] QRect geometry()
func (this *QSystemTrayIcon) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QSystemTrayIcon) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSystemTrayIcon9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QSystemTrayIcon) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void show()
func (this *QSystemTrayIcon) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hide()
func (this *QSystemTrayIcon) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, const QIcon &, int)
func (this *QSystemTrayIcon) ShowMessage(title *qtcore.QString, msg *qtcore.QString, icon *qtgui.QIcon, msecs int) {
	var convArg0 = title.GetCthis()
	var convArg1 = msg.GetCthis()
	var convArg2 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)
func (this *QSystemTrayIcon) ShowMessage_1(title *qtcore.QString, msg *qtcore.QString, icon int, msecs int) {
	var convArg0 = title.GetCthis()
	var convArg1 = msg.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, icon, msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(QSystemTrayIcon::ActivationReason)
func (this *QSystemTrayIcon) Activated(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon9activatedENS_16ActivationReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageClicked()
func (this *QSystemTrayIcon) MessageClicked() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon14messageClickedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSystemTrayIcon) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSystemTrayIcon5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QSystemTrayIcon__ActivationReason = int

const QSystemTrayIcon__Unknown QSystemTrayIcon__ActivationReason = 0
const QSystemTrayIcon__Context QSystemTrayIcon__ActivationReason = 1
const QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = 2
const QSystemTrayIcon__Trigger QSystemTrayIcon__ActivationReason = 3
const QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = 4

type QSystemTrayIcon__MessageIcon = int

const QSystemTrayIcon__NoIcon QSystemTrayIcon__MessageIcon = 0
const QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = 1
const QSystemTrayIcon__Warning QSystemTrayIcon__MessageIcon = 2
const QSystemTrayIcon__Critical QSystemTrayIcon__MessageIcon = 3

//  body block end
