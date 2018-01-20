//  header block begin
// /usr/include/qt/QtWidgets/qsystemtrayicon.h
// #include <qsystemtrayicon.h>
// #include <QtWidgets>
package qtwidgets

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
type QSystemTrayIcon struct {
	*qtcore.QObject
}

func (this *QSystemTrayIcon) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQSystemTrayIconFromPointer(cthis unsafe.Pointer) *QSystemTrayIcon {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSystemTrayIcon{bcthis0}
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSystemTrayIcon) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// Public
// void QSystemTrayIcon(class QObject *)
func NewQSystemTrayIcon(parent unsafe.Pointer) *QSystemTrayIcon {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIconC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:70
// index:1
// Public
// void QSystemTrayIcon(const class QIcon &, class QObject *)
func NewQSystemTrayIcon_1(icon *qtgui.QIcon, parent unsafe.Pointer) *QSystemTrayIcon {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIconC2ERK5QIconP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSystemTrayIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:71
// index:0
// Public virtual
// void ~QSystemTrayIcon()
func DeleteQSystemTrayIcon(*QSystemTrayIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIconD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:82
// index:0
// Public
// void setContextMenu(class QMenu *)
func (this *QSystemTrayIcon) SetContextMenu(menu unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon14setContextMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:83
// index:0
// Public
// QMenu * contextMenu()
func (this *QSystemTrayIcon) ContextMenu() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon11contextMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:86
// index:0
// Public
// QIcon icon()
func (this *QSystemTrayIcon) Icon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:87
// index:0
// Public
// void setIcon(const class QIcon &)
func (this *QSystemTrayIcon) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:89
// index:0
// Public
// QString toolTip()
func (this *QSystemTrayIcon) ToolTip() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:90
// index:0
// Public
// void setToolTip(const class QString &)
func (this *QSystemTrayIcon) SetToolTip(tip *qtcore.QString) {
	var convArg0 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:92
// index:0
// Public static
// bool isSystemTrayAvailable()
func (this *QSystemTrayIcon) IsSystemTrayAvailable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon21isSystemTrayAvailableEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSystemTrayIcon_IsSystemTrayAvailable() {
	var nilthis *QSystemTrayIcon
	nilthis.IsSystemTrayAvailable()
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:93
// index:0
// Public static
// bool supportsMessages()
func (this *QSystemTrayIcon) SupportsMessages() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon16supportsMessagesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSystemTrayIcon_SupportsMessages() {
	var nilthis *QSystemTrayIcon
	nilthis.SupportsMessages()
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:97
// index:0
// Public
// QRect geometry()
func (this *QSystemTrayIcon) Geometry() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:98
// index:0
// Public
// bool isVisible()
func (this *QSystemTrayIcon) IsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:101
// index:0
// Public
// void setVisible(_Bool)
func (this *QSystemTrayIcon) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:102
// index:0
// Public inline
// void show()
func (this *QSystemTrayIcon) Show() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon4showEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:103
// index:0
// Public inline
// void hide()
func (this *QSystemTrayIcon) Hide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon4hideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// Public
// void showMessage(const class QString &, const class QString &, const class QIcon &, int)
func (this *QSystemTrayIcon) ShowMessage(title *qtcore.QString, msg *qtcore.QString, icon *qtgui.QIcon, msecs int) {
	var convArg0 = title.GetCthis()
	var convArg1 = msg.GetCthis()
	var convArg2 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public
// void showMessage(const class QString &, const class QString &, class QSystemTrayIcon::MessageIcon, int)
func (this *QSystemTrayIcon) ShowMessage_1(title *qtcore.QString, msg *qtcore.QString, icon int, msecs int) {
	var convArg0 = title.GetCthis()
	var convArg1 = msg.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &icon, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:109
// index:0
// Public
// void activated(class QSystemTrayIcon::ActivationReason)
func (this *QSystemTrayIcon) Activated(reason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon9activatedENS_16ActivationReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:110
// index:0
// Public
// void messageClicked()
func (this *QSystemTrayIcon) MessageClicked() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon14messageClickedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:113
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSystemTrayIcon) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
