//  header block begin
// /usr/include/qt/QtGui/qguiapplication.h
// #include <qguiapplication.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QGuiApplication struct {
	*qtcore.QCoreApplication
}

func (this *QGuiApplication) GetCthis() unsafe.Pointer {
	return this.QCoreApplication.GetCthis()
}

// /usr/include/qt/QtGui/qguiapplication.h:74
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGuiApplication) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:87
// index:0
// void QGuiApplication(int &, char **, int)
func NewQGuiApplication(argc int, argv []string, arg2 int) *QGuiApplication {
	cthis := qtrt.Calloc(1, 256)
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, &arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQGuiApplicationFromPointer(cthis)
	return gothis
}
func NewQGuiApplicationFromPointer(cthis unsafe.Pointer) *QGuiApplication {
	bcthis0 := qtcore.NewQCoreApplicationFromPointer(cthis)
	return &QGuiApplication{bcthis0}
}

// /usr/include/qt/QtGui/qguiapplication.h:195
// index:1
// void QGuiApplication(class QGuiApplicationPrivate &)
func NewQGuiApplication_1(p unsafe.Pointer) *QGuiApplication {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplicationC2ER22QGuiApplicationPrivate", ffiqt.FFI_TYPE_VOID, cthis, p)
	gopp.ErrPrint(err, rv)
	gothis := NewQGuiApplicationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qguiapplication.h:89
// index:0
// virtual
// void ~QGuiApplication()
func DeleteQGuiApplication(*QGuiApplication) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplicationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:91
// index:0
// static
// void setApplicationDisplayName(const class QString &)
func (this *QGuiApplication) SetApplicationDisplayName(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication25setApplicationDisplayNameERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetApplicationDisplayName(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	var nilthis *QGuiApplication
	nilthis.SetApplicationDisplayName(name)
}

// /usr/include/qt/QtGui/qguiapplication.h:92
// index:0
// static
// QString applicationDisplayName()
func (this *QGuiApplication) ApplicationDisplayName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22applicationDisplayNameEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_ApplicationDisplayName() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.ApplicationDisplayName()
}

// /usr/include/qt/QtGui/qguiapplication.h:94
// index:0
// static
// void setDesktopFileName(const class QString &)
func (this *QGuiApplication) SetDesktopFileName(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18setDesktopFileNameERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetDesktopFileName(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	var nilthis *QGuiApplication
	nilthis.SetDesktopFileName(name)
}

// /usr/include/qt/QtGui/qguiapplication.h:95
// index:0
// static
// QString desktopFileName()
func (this *QGuiApplication) DesktopFileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15desktopFileNameEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_DesktopFileName() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.DesktopFileName()
}

// /usr/include/qt/QtGui/qguiapplication.h:97
// index:0
// static
// QWindowList allWindows()
func (this *QGuiApplication) AllWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10allWindowsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_AllWindows() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.AllWindows()
}

// /usr/include/qt/QtGui/qguiapplication.h:98
// index:0
// static
// QWindowList topLevelWindows()
func (this *QGuiApplication) TopLevelWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15topLevelWindowsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_TopLevelWindows() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.TopLevelWindows()
}

// /usr/include/qt/QtGui/qguiapplication.h:99
// index:0
// static
// QWindow * topLevelAt(const class QPoint &)
func (this *QGuiApplication) TopLevelAt(pos unsafe.Pointer) {
	// 0: (pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10topLevelAtERK6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_TopLevelAt(pos unsafe.Pointer) {
	// 0: (pos const QPoint &), (pos)
	var nilthis *QGuiApplication
	nilthis.TopLevelAt(pos)
}

// /usr/include/qt/QtGui/qguiapplication.h:101
// index:0
// static
// void setWindowIcon(const class QIcon &)
func (this *QGuiApplication) SetWindowIcon(icon unsafe.Pointer) {
	// 0: (icon const QIcon &), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetWindowIcon(icon unsafe.Pointer) {
	// 0: (icon const QIcon &), (icon)
	var nilthis *QGuiApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtGui/qguiapplication.h:102
// index:0
// static
// QIcon windowIcon()
func (this *QGuiApplication) WindowIcon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10windowIconEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_WindowIcon() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.WindowIcon()
}

// /usr/include/qt/QtGui/qguiapplication.h:104
// index:0
// static
// QString platformName()
func (this *QGuiApplication) PlatformName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication12platformNameEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_PlatformName() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.PlatformName()
}

// /usr/include/qt/QtGui/qguiapplication.h:106
// index:0
// static
// QWindow * modalWindow()
func (this *QGuiApplication) ModalWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11modalWindowEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_ModalWindow() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.ModalWindow()
}

// /usr/include/qt/QtGui/qguiapplication.h:108
// index:0
// static
// QWindow * focusWindow()
func (this *QGuiApplication) FocusWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11focusWindowEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_FocusWindow() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.FocusWindow()
}

// /usr/include/qt/QtGui/qguiapplication.h:109
// index:0
// static
// QObject * focusObject()
func (this *QGuiApplication) FocusObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11focusObjectEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_FocusObject() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.FocusObject()
}

// /usr/include/qt/QtGui/qguiapplication.h:111
// index:0
// static
// QScreen * primaryScreen()
func (this *QGuiApplication) PrimaryScreen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13primaryScreenEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_PrimaryScreen() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.PrimaryScreen()
}

// /usr/include/qt/QtGui/qguiapplication.h:112
// index:0
// static
// QList<QScreen *> screens()
func (this *QGuiApplication) Screens() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7screensEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Screens() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.Screens()
}

// /usr/include/qt/QtGui/qguiapplication.h:113
// index:0
// static
// QScreen * screenAt(const class QPoint &)
func (this *QGuiApplication) ScreenAt(point unsafe.Pointer) {
	// 0: (point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication8screenAtERK6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_ScreenAt(point unsafe.Pointer) {
	// 0: (point const QPoint &), (point)
	var nilthis *QGuiApplication
	nilthis.ScreenAt(point)
}

// /usr/include/qt/QtGui/qguiapplication.h:115
// index:0
// qreal devicePixelRatio()
func (this *QGuiApplication) DevicePixelRatio() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication16devicePixelRatioEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:118
// index:0
// static
// QCursor * overrideCursor()
func (this *QGuiApplication) OverrideCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication14overrideCursorEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_OverrideCursor() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.OverrideCursor()
}

// /usr/include/qt/QtGui/qguiapplication.h:119
// index:0
// static
// void setOverrideCursor(const class QCursor &)
func (this *QGuiApplication) SetOverrideCursor(arg0 unsafe.Pointer) {
	// 0: (const QCursor & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17setOverrideCursorERK7QCursor", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetOverrideCursor(arg0 unsafe.Pointer) {
	// 0: (const QCursor & arg0), (arg0)
	var nilthis *QGuiApplication
	nilthis.SetOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:120
// index:0
// static
// void changeOverrideCursor(const class QCursor &)
func (this *QGuiApplication) ChangeOverrideCursor(arg0 unsafe.Pointer) {
	// 0: (const QCursor & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20changeOverrideCursorERK7QCursor", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_ChangeOverrideCursor(arg0 unsafe.Pointer) {
	// 0: (const QCursor & arg0), (arg0)
	var nilthis *QGuiApplication
	nilthis.ChangeOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:121
// index:0
// static
// void restoreOverrideCursor()
func (this *QGuiApplication) RestoreOverrideCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication21restoreOverrideCursorEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_RestoreOverrideCursor() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.RestoreOverrideCursor()
}

// /usr/include/qt/QtGui/qguiapplication.h:124
// index:0
// static
// QFont font()
func (this *QGuiApplication) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4fontEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Font() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.Font()
}

// /usr/include/qt/QtGui/qguiapplication.h:125
// index:0
// static
// void setFont(const class QFont &)
func (this *QGuiApplication) SetFont(arg0 unsafe.Pointer) {
	// 0: (const QFont & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7setFontERK5QFont", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetFont(arg0 unsafe.Pointer) {
	// 0: (const QFont & arg0), (arg0)
	var nilthis *QGuiApplication
	nilthis.SetFont(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:128
// index:0
// static
// QClipboard * clipboard()
func (this *QGuiApplication) Clipboard() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication9clipboardEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Clipboard() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.Clipboard()
}

// /usr/include/qt/QtGui/qguiapplication.h:131
// index:0
// static
// QPalette palette()
func (this *QGuiApplication) Palette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7paletteEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Palette() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.Palette()
}

// /usr/include/qt/QtGui/qguiapplication.h:132
// index:0
// static
// void setPalette(const class QPalette &)
func (this *QGuiApplication) SetPalette(pal unsafe.Pointer) {
	// 0: (pal const QPalette &), (pal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10setPaletteERK8QPalette", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetPalette(pal unsafe.Pointer) {
	// 0: (pal const QPalette &), (pal)
	var nilthis *QGuiApplication
	nilthis.SetPalette(pal)
}

// /usr/include/qt/QtGui/qguiapplication.h:134
// index:0
// static
// Qt::KeyboardModifiers keyboardModifiers()
func (this *QGuiApplication) KeyboardModifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17keyboardModifiersEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_KeyboardModifiers() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.KeyboardModifiers()
}

// /usr/include/qt/QtGui/qguiapplication.h:135
// index:0
// static
// Qt::KeyboardModifiers queryKeyboardModifiers()
func (this *QGuiApplication) QueryKeyboardModifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22queryKeyboardModifiersEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_QueryKeyboardModifiers() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.QueryKeyboardModifiers()
}

// /usr/include/qt/QtGui/qguiapplication.h:136
// index:0
// static
// Qt::MouseButtons mouseButtons()
func (this *QGuiApplication) MouseButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication12mouseButtonsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_MouseButtons() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.MouseButtons()
}

// /usr/include/qt/QtGui/qguiapplication.h:138
// index:0
// static
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QGuiApplication) SetLayoutDirection(direction int) {
	// 0: (direction Qt::LayoutDirection), (direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetLayoutDirection(direction int) {
	// 0: (direction Qt::LayoutDirection), (direction)
	var nilthis *QGuiApplication
	nilthis.SetLayoutDirection(direction)
}

// /usr/include/qt/QtGui/qguiapplication.h:139
// index:0
// static
// Qt::LayoutDirection layoutDirection()
func (this *QGuiApplication) LayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15layoutDirectionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_LayoutDirection() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.LayoutDirection()
}

// /usr/include/qt/QtGui/qguiapplication.h:141
// index:0
// static inline
// bool isRightToLeft()
func (this *QGuiApplication) IsRightToLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13isRightToLeftEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_IsRightToLeft() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.IsRightToLeft()
}

// /usr/include/qt/QtGui/qguiapplication.h:142
// index:0
// static inline
// bool isLeftToRight()
func (this *QGuiApplication) IsLeftToRight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13isLeftToRightEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_IsLeftToRight() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.IsLeftToRight()
}

// /usr/include/qt/QtGui/qguiapplication.h:144
// index:0
// static
// QStyleHints * styleHints()
func (this *QGuiApplication) StyleHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10styleHintsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_StyleHints() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.StyleHints()
}

// /usr/include/qt/QtGui/qguiapplication.h:145
// index:0
// static
// void setDesktopSettingsAware(_Bool)
func (this *QGuiApplication) SetDesktopSettingsAware(on bool) {
	// 0: (on bool), (on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23setDesktopSettingsAwareEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetDesktopSettingsAware(on bool) {
	// 0: (on bool), (on)
	var nilthis *QGuiApplication
	nilthis.SetDesktopSettingsAware(on)
}

// /usr/include/qt/QtGui/qguiapplication.h:146
// index:0
// static
// bool desktopSettingsAware()
func (this *QGuiApplication) DesktopSettingsAware() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20desktopSettingsAwareEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_DesktopSettingsAware() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.DesktopSettingsAware()
}

// /usr/include/qt/QtGui/qguiapplication.h:148
// index:0
// static
// QInputMethod * inputMethod()
func (this *QGuiApplication) InputMethod() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11inputMethodEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_InputMethod() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.InputMethod()
}

// /usr/include/qt/QtGui/qguiapplication.h:150
// index:0
// static
// QPlatformNativeInterface * platformNativeInterface()
func (this *QGuiApplication) PlatformNativeInterface() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23platformNativeInterfaceEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_PlatformNativeInterface() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.PlatformNativeInterface()
}

// /usr/include/qt/QtGui/qguiapplication.h:152
// index:0
// static
// QFunctionPointer platformFunction(const class QByteArray &)
func (this *QGuiApplication) PlatformFunction(function unsafe.Pointer) {
	// 0: (function const QByteArray &), (function)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16platformFunctionERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_PlatformFunction(function unsafe.Pointer) {
	// 0: (function const QByteArray &), (function)
	var nilthis *QGuiApplication
	nilthis.PlatformFunction(function)
}

// /usr/include/qt/QtGui/qguiapplication.h:154
// index:0
// static
// void setQuitOnLastWindowClosed(_Bool)
func (this *QGuiApplication) SetQuitOnLastWindowClosed(quit bool) {
	// 0: (quit bool), (quit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication25setQuitOnLastWindowClosedEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetQuitOnLastWindowClosed(quit bool) {
	// 0: (quit bool), (quit)
	var nilthis *QGuiApplication
	nilthis.SetQuitOnLastWindowClosed(quit)
}

// /usr/include/qt/QtGui/qguiapplication.h:155
// index:0
// static
// bool quitOnLastWindowClosed()
func (this *QGuiApplication) QuitOnLastWindowClosed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22quitOnLastWindowClosedEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_QuitOnLastWindowClosed() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.QuitOnLastWindowClosed()
}

// /usr/include/qt/QtGui/qguiapplication.h:157
// index:0
// static
// Qt::ApplicationState applicationState()
func (this *QGuiApplication) ApplicationState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16applicationStateEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_ApplicationState() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.ApplicationState()
}

// /usr/include/qt/QtGui/qguiapplication.h:159
// index:0
// static
// int exec()
func (this *QGuiApplication) Exec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4execEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Exec() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.Exec()
}

// /usr/include/qt/QtGui/qguiapplication.h:160
// index:0
// virtual
// bool notify(class QObject *, class QEvent *)
func (this *QGuiApplication) Notify(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:164
// index:0
// bool isSessionRestored()
func (this *QGuiApplication) IsSessionRestored() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication17isSessionRestoredEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:165
// index:0
// QString sessionId()
func (this *QGuiApplication) SessionId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication9sessionIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:166
// index:0
// QString sessionKey()
func (this *QGuiApplication) SessionKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication10sessionKeyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:167
// index:0
// bool isSavingSession()
func (this *QGuiApplication) IsSavingSession() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication15isSavingSessionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:169
// index:0
// static
// bool isFallbackSessionManagementEnabled()
func (this *QGuiApplication) IsFallbackSessionManagementEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication34isFallbackSessionManagementEnabledEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_IsFallbackSessionManagementEnabled() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.IsFallbackSessionManagementEnabled()
}

// /usr/include/qt/QtGui/qguiapplication.h:170
// index:0
// static
// void setFallbackSessionManagementEnabled(_Bool)
func (this *QGuiApplication) SetFallbackSessionManagementEnabled(arg0 bool) {
	// 0: (bool arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication35setFallbackSessionManagementEnabledEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetFallbackSessionManagementEnabled(arg0 bool) {
	// 0: (bool arg0), (arg0)
	var nilthis *QGuiApplication
	nilthis.SetFallbackSessionManagementEnabled(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:173
// index:0
// static
// void sync()
func (this *QGuiApplication) Sync() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4syncEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Sync() {
	// 0: (), ()
	var nilthis *QGuiApplication
	nilthis.Sync()
}

// /usr/include/qt/QtGui/qguiapplication.h:175
// index:0
// void fontDatabaseChanged()
func (this *QGuiApplication) FontDatabaseChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication19fontDatabaseChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:176
// index:0
// void screenAdded(class QScreen *)
func (this *QGuiApplication) ScreenAdded(screen unsafe.Pointer) {
	// 0: (, screen QScreen *), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11screenAddedEP7QScreen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:177
// index:0
// void screenRemoved(class QScreen *)
func (this *QGuiApplication) ScreenRemoved(screen unsafe.Pointer) {
	// 0: (, screen QScreen *), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13screenRemovedEP7QScreen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:178
// index:0
// void primaryScreenChanged(class QScreen *)
func (this *QGuiApplication) PrimaryScreenChanged(screen unsafe.Pointer) {
	// 0: (, screen QScreen *), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20primaryScreenChangedEP7QScreen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:179
// index:0
// void lastWindowClosed()
func (this *QGuiApplication) LastWindowClosed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16lastWindowClosedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:180
// index:0
// void focusObjectChanged(class QObject *)
func (this *QGuiApplication) FocusObjectChanged(focusObject unsafe.Pointer) {
	// 0: (, focusObject QObject *), (focusObject)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18focusObjectChangedEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), focusObject)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:181
// index:0
// void focusWindowChanged(class QWindow *)
func (this *QGuiApplication) FocusWindowChanged(focusWindow unsafe.Pointer) {
	// 0: (, focusWindow QWindow *), (focusWindow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18focusWindowChangedEP7QWindow", ffiqt.FFI_TYPE_VOID, this.GetCthis(), focusWindow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:182
// index:0
// void applicationStateChanged(Qt::ApplicationState)
func (this *QGuiApplication) ApplicationStateChanged(state int) {
	// 0: (, state Qt::ApplicationState), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23applicationStateChangedEN2Qt16ApplicationStateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:183
// index:0
// void layoutDirectionChanged(Qt::LayoutDirection)
func (this *QGuiApplication) LayoutDirectionChanged(direction int) {
	// 0: (, direction Qt::LayoutDirection), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22layoutDirectionChangedEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:185
// index:0
// void commitDataRequest(class QSessionManager &)
func (this *QGuiApplication) CommitDataRequest(sessionManager unsafe.Pointer) {
	// 0: (, sessionManager QSessionManager &), (sessionManager)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17commitDataRequestER15QSessionManager", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sessionManager)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:186
// index:0
// void saveStateRequest(class QSessionManager &)
func (this *QGuiApplication) SaveStateRequest(sessionManager unsafe.Pointer) {
	// 0: (, sessionManager QSessionManager &), (sessionManager)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16saveStateRequestER15QSessionManager", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sessionManager)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:188
// index:0
// void paletteChanged(const class QPalette &)
func (this *QGuiApplication) PaletteChanged(pal unsafe.Pointer) {
	// 0: (, pal const QPalette &), (pal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication14paletteChangedERK8QPalette", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:189
// index:0
// void applicationDisplayNameChanged()
func (this *QGuiApplication) ApplicationDisplayNameChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication29applicationDisplayNameChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:192
// index:0
// virtual
// bool event(class QEvent *)
func (this *QGuiApplication) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:193
// index:0
// virtual
// bool compressEvent(class QEvent *, class QObject *, class QPostEventList *)
func (this *QGuiApplication) CompressEvent(arg0 unsafe.Pointer, receiver unsafe.Pointer, arg2 unsafe.Pointer) {
	// 0: (, QEvent * arg0, receiver QObject *, QPostEventList * arg2), (arg0, receiver, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13compressEventEP6QEventP7QObjectP14QPostEventList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, receiver, arg2)
	gopp.ErrPrint(err, rv)
}

//  body block end
