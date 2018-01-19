//  header block begin
// /usr/include/qt/QtWidgets/qsystemtrayicon.h
// #include <qsystemtrayicon.h>
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
type QSystemTrayIcon struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSystemTrayIcon) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// void QSystemTrayIcon(class QObject *)
func NewQSystemTrayIcon(parent unsafe.Pointer) *QSystemTrayIcon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIconC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSystemTrayIcon{cthis}
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:70
// index:1
// void QSystemTrayIcon(const class QIcon &, class QObject *)
func NewQSystemTrayIcon_1(icon unsafe.Pointer, parent unsafe.Pointer) *QSystemTrayIcon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIconC2ERK5QIconP7QObject", ffiqt.FFI_TYPE_VOID, cthis, icon, parent)
	gopp.ErrPrint(err, rv)
	return &QSystemTrayIcon{cthis}
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:71
// index:0
// virtual
// void ~QSystemTrayIcon()
func DeleteQSystemTrayIcon(*QSystemTrayIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIconD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:82
// index:0
// void setContextMenu(class QMenu *)
func (this *QSystemTrayIcon) SetContextMenu(menu unsafe.Pointer) {
	// 0: (, QMenu * menu), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon14setContextMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.cthis, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:83
// index:0
// QMenu * contextMenu()
func (this *QSystemTrayIcon) ContextMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon11contextMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:86
// index:0
// QIcon icon()
func (this *QSystemTrayIcon) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon4iconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:87
// index:0
// void setIcon(const class QIcon &)
func (this *QSystemTrayIcon) SetIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:89
// index:0
// QString toolTip()
func (this *QSystemTrayIcon) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon7toolTipEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:90
// index:0
// void setToolTip(const class QString &)
func (this *QSystemTrayIcon) SetToolTip(tip unsafe.Pointer) {
	// 0: (, const QString & tip), (tip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, tip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:92
// index:0
// static
// bool isSystemTrayAvailable()
func (this *QSystemTrayIcon) IsSystemTrayAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon21isSystemTrayAvailableEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSystemTrayIcon_IsSystemTrayAvailable() {
	// 0: (), ()
	var nilthis *QSystemTrayIcon
	nilthis.IsSystemTrayAvailable()
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:93
// index:0
// static
// bool supportsMessages()
func (this *QSystemTrayIcon) SupportsMessages() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon16supportsMessagesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSystemTrayIcon_SupportsMessages() {
	// 0: (), ()
	var nilthis *QSystemTrayIcon
	nilthis.SupportsMessages()
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:97
// index:0
// QRect geometry()
func (this *QSystemTrayIcon) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon8geometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:98
// index:0
// bool isVisible()
func (this *QSystemTrayIcon) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSystemTrayIcon9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:101
// index:0
// void setVisible(_Bool)
func (this *QSystemTrayIcon) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:102
// index:0
// inline
// void show()
func (this *QSystemTrayIcon) Show() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon4showEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:103
// index:0
// inline
// void hide()
func (this *QSystemTrayIcon) Hide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon4hideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// void showMessage(const class QString &, const class QString &, const class QIcon &, int)
func (this *QSystemTrayIcon) ShowMessage(title unsafe.Pointer, msg unsafe.Pointer, icon unsafe.Pointer, msecs int) {
	// 0: (, const QString & title, const QString & msg, const QIcon & icon, int msecs), (title, msg, icon, &msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", ffiqt.FFI_TYPE_VOID, this.cthis, title, msg, icon, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// void showMessage(const class QString &, const class QString &, class QSystemTrayIcon::MessageIcon, int)
func (this *QSystemTrayIcon) ShowMessage_1(title unsafe.Pointer, msg unsafe.Pointer, icon int, msecs int) {
	// 1: (, const QString & title, const QString & msg, QSystemTrayIcon::MessageIcon icon, int msecs), (title, msg, &icon, &msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", ffiqt.FFI_TYPE_VOID, this.cthis, title, msg, &icon, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:109
// index:0
// void activated(class QSystemTrayIcon::ActivationReason)
func (this *QSystemTrayIcon) Activated(reason int) {
	// 0: (, QSystemTrayIcon::ActivationReason reason), (&reason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon9activatedENS_16ActivationReasonE", ffiqt.FFI_TYPE_VOID, this.cthis, &reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:110
// index:0
// void messageClicked()
func (this *QSystemTrayIcon) MessageClicked() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSystemTrayIcon14messageClickedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
