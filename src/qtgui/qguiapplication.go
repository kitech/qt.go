//  header block begin
// /usr/include/qt/QtGui/qguiapplication.h
// #include <qguiapplication.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
func NewQGuiApplicationFromPointer(cthis unsafe.Pointer) *QGuiApplication {
	bcthis0 := qtcore.NewQCoreApplicationFromPointer(cthis)
	return &QGuiApplication{bcthis0}
}

// /usr/include/qt/QtGui/qguiapplication.h:74
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGuiApplication) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:87
// index:0
// Public
// void QGuiApplication(int &, char **, int)
func NewQGuiApplication(argc int, argv []string, arg2 int) *QGuiApplication {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, &arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQGuiApplicationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qguiapplication.h:89
// index:0
// Public virtual
// void ~QGuiApplication()
func DeleteQGuiApplication(*QGuiApplication) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplicationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:91
// index:0
// Public static
// void setApplicationDisplayName(const class QString &)
func (this *QGuiApplication) SetApplicationDisplayName(name *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication25setApplicationDisplayNameERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetApplicationDisplayName(name *qtcore.QString) {
	var nilthis *QGuiApplication
	nilthis.SetApplicationDisplayName(name)
}

// /usr/include/qt/QtGui/qguiapplication.h:92
// index:0
// Public static
// QString applicationDisplayName()
func (this *QGuiApplication) ApplicationDisplayName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22applicationDisplayNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_ApplicationDisplayName() {
	var nilthis *QGuiApplication
	nilthis.ApplicationDisplayName()
}

// /usr/include/qt/QtGui/qguiapplication.h:94
// index:0
// Public static
// void setDesktopFileName(const class QString &)
func (this *QGuiApplication) SetDesktopFileName(name *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18setDesktopFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetDesktopFileName(name *qtcore.QString) {
	var nilthis *QGuiApplication
	nilthis.SetDesktopFileName(name)
}

// /usr/include/qt/QtGui/qguiapplication.h:95
// index:0
// Public static
// QString desktopFileName()
func (this *QGuiApplication) DesktopFileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15desktopFileNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_DesktopFileName() {
	var nilthis *QGuiApplication
	nilthis.DesktopFileName()
}

// /usr/include/qt/QtGui/qguiapplication.h:97
// index:0
// Public static
// QWindowList allWindows()
func (this *QGuiApplication) AllWindows() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10allWindowsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_AllWindows() {
	var nilthis *QGuiApplication
	nilthis.AllWindows()
}

// /usr/include/qt/QtGui/qguiapplication.h:98
// index:0
// Public static
// QWindowList topLevelWindows()
func (this *QGuiApplication) TopLevelWindows() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15topLevelWindowsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_TopLevelWindows() {
	var nilthis *QGuiApplication
	nilthis.TopLevelWindows()
}

// /usr/include/qt/QtGui/qguiapplication.h:99
// index:0
// Public static
// QWindow * topLevelAt(const class QPoint &)
func (this *QGuiApplication) TopLevelAt(pos *qtcore.QPoint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10topLevelAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, pos)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_TopLevelAt(pos *qtcore.QPoint) {
	var nilthis *QGuiApplication
	nilthis.TopLevelAt(pos)
}

// /usr/include/qt/QtGui/qguiapplication.h:101
// index:0
// Public static
// void setWindowIcon(const class QIcon &)
func (this *QGuiApplication) SetWindowIcon(icon *QIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, icon)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetWindowIcon(icon *QIcon) {
	var nilthis *QGuiApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtGui/qguiapplication.h:102
// index:0
// Public static
// QIcon windowIcon()
func (this *QGuiApplication) WindowIcon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10windowIconEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_WindowIcon() {
	var nilthis *QGuiApplication
	nilthis.WindowIcon()
}

// /usr/include/qt/QtGui/qguiapplication.h:104
// index:0
// Public static
// QString platformName()
func (this *QGuiApplication) PlatformName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication12platformNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_PlatformName() {
	var nilthis *QGuiApplication
	nilthis.PlatformName()
}

// /usr/include/qt/QtGui/qguiapplication.h:106
// index:0
// Public static
// QWindow * modalWindow()
func (this *QGuiApplication) ModalWindow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11modalWindowEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_ModalWindow() {
	var nilthis *QGuiApplication
	nilthis.ModalWindow()
}

// /usr/include/qt/QtGui/qguiapplication.h:108
// index:0
// Public static
// QWindow * focusWindow()
func (this *QGuiApplication) FocusWindow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11focusWindowEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_FocusWindow() {
	var nilthis *QGuiApplication
	nilthis.FocusWindow()
}

// /usr/include/qt/QtGui/qguiapplication.h:109
// index:0
// Public static
// QObject * focusObject()
func (this *QGuiApplication) FocusObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11focusObjectEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_FocusObject() {
	var nilthis *QGuiApplication
	nilthis.FocusObject()
}

// /usr/include/qt/QtGui/qguiapplication.h:111
// index:0
// Public static
// QScreen * primaryScreen()
func (this *QGuiApplication) PrimaryScreen() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13primaryScreenEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_PrimaryScreen() {
	var nilthis *QGuiApplication
	nilthis.PrimaryScreen()
}

// /usr/include/qt/QtGui/qguiapplication.h:112
// index:0
// Public static
// QList<QScreen *> screens()
func (this *QGuiApplication) Screens() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7screensEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_Screens() {
	var nilthis *QGuiApplication
	nilthis.Screens()
}

// /usr/include/qt/QtGui/qguiapplication.h:113
// index:0
// Public static
// QScreen * screenAt(const class QPoint &)
func (this *QGuiApplication) ScreenAt(point *qtcore.QPoint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication8screenAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, point)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_ScreenAt(point *qtcore.QPoint) {
	var nilthis *QGuiApplication
	nilthis.ScreenAt(point)
}

// /usr/include/qt/QtGui/qguiapplication.h:115
// index:0
// Public
// qreal devicePixelRatio()
func (this *QGuiApplication) DevicePixelRatio() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication16devicePixelRatioEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:118
// index:0
// Public static
// QCursor * overrideCursor()
func (this *QGuiApplication) OverrideCursor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication14overrideCursorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_OverrideCursor() {
	var nilthis *QGuiApplication
	nilthis.OverrideCursor()
}

// /usr/include/qt/QtGui/qguiapplication.h:119
// index:0
// Public static
// void setOverrideCursor(const class QCursor &)
func (this *QGuiApplication) SetOverrideCursor(arg0 *QCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17setOverrideCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetOverrideCursor(arg0 *QCursor) {
	var nilthis *QGuiApplication
	nilthis.SetOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:120
// index:0
// Public static
// void changeOverrideCursor(const class QCursor &)
func (this *QGuiApplication) ChangeOverrideCursor(arg0 *QCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20changeOverrideCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_ChangeOverrideCursor(arg0 *QCursor) {
	var nilthis *QGuiApplication
	nilthis.ChangeOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:121
// index:0
// Public static
// void restoreOverrideCursor()
func (this *QGuiApplication) RestoreOverrideCursor() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication21restoreOverrideCursorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_RestoreOverrideCursor() {
	var nilthis *QGuiApplication
	nilthis.RestoreOverrideCursor()
}

// /usr/include/qt/QtGui/qguiapplication.h:124
// index:0
// Public static
// QFont font()
func (this *QGuiApplication) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4fontEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_Font() {
	var nilthis *QGuiApplication
	nilthis.Font()
}

// /usr/include/qt/QtGui/qguiapplication.h:125
// index:0
// Public static
// void setFont(const class QFont &)
func (this *QGuiApplication) SetFont(arg0 *QFont) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetFont(arg0 *QFont) {
	var nilthis *QGuiApplication
	nilthis.SetFont(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:128
// index:0
// Public static
// QClipboard * clipboard()
func (this *QGuiApplication) Clipboard() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication9clipboardEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_Clipboard() {
	var nilthis *QGuiApplication
	nilthis.Clipboard()
}

// /usr/include/qt/QtGui/qguiapplication.h:131
// index:0
// Public static
// QPalette palette()
func (this *QGuiApplication) Palette() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7paletteEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_Palette() {
	var nilthis *QGuiApplication
	nilthis.Palette()
}

// /usr/include/qt/QtGui/qguiapplication.h:132
// index:0
// Public static
// void setPalette(const class QPalette &)
func (this *QGuiApplication) SetPalette(pal *QPalette) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10setPaletteERK8QPalette", ffiqt.FFI_TYPE_POINTER, pal)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetPalette(pal *QPalette) {
	var nilthis *QGuiApplication
	nilthis.SetPalette(pal)
}

// /usr/include/qt/QtGui/qguiapplication.h:134
// index:0
// Public static
// Qt::KeyboardModifiers keyboardModifiers()
func (this *QGuiApplication) KeyboardModifiers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17keyboardModifiersEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_KeyboardModifiers() {
	var nilthis *QGuiApplication
	nilthis.KeyboardModifiers()
}

// /usr/include/qt/QtGui/qguiapplication.h:135
// index:0
// Public static
// Qt::KeyboardModifiers queryKeyboardModifiers()
func (this *QGuiApplication) QueryKeyboardModifiers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22queryKeyboardModifiersEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_QueryKeyboardModifiers() {
	var nilthis *QGuiApplication
	nilthis.QueryKeyboardModifiers()
}

// /usr/include/qt/QtGui/qguiapplication.h:136
// index:0
// Public static
// Qt::MouseButtons mouseButtons()
func (this *QGuiApplication) MouseButtons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication12mouseButtonsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_MouseButtons() {
	var nilthis *QGuiApplication
	nilthis.MouseButtons()
}

// /usr/include/qt/QtGui/qguiapplication.h:138
// index:0
// Public static
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QGuiApplication) SetLayoutDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, direction)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetLayoutDirection(direction int) {
	var nilthis *QGuiApplication
	nilthis.SetLayoutDirection(direction)
}

// /usr/include/qt/QtGui/qguiapplication.h:139
// index:0
// Public static
// Qt::LayoutDirection layoutDirection()
func (this *QGuiApplication) LayoutDirection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_LayoutDirection() {
	var nilthis *QGuiApplication
	nilthis.LayoutDirection()
}

// /usr/include/qt/QtGui/qguiapplication.h:141
// index:0
// Public static inline
// bool isRightToLeft()
func (this *QGuiApplication) IsRightToLeft() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13isRightToLeftEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_IsRightToLeft() {
	var nilthis *QGuiApplication
	nilthis.IsRightToLeft()
}

// /usr/include/qt/QtGui/qguiapplication.h:142
// index:0
// Public static inline
// bool isLeftToRight()
func (this *QGuiApplication) IsLeftToRight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13isLeftToRightEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_IsLeftToRight() {
	var nilthis *QGuiApplication
	nilthis.IsLeftToRight()
}

// /usr/include/qt/QtGui/qguiapplication.h:144
// index:0
// Public static
// QStyleHints * styleHints()
func (this *QGuiApplication) StyleHints() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10styleHintsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_StyleHints() {
	var nilthis *QGuiApplication
	nilthis.StyleHints()
}

// /usr/include/qt/QtGui/qguiapplication.h:145
// index:0
// Public static
// void setDesktopSettingsAware(_Bool)
func (this *QGuiApplication) SetDesktopSettingsAware(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23setDesktopSettingsAwareEb", ffiqt.FFI_TYPE_POINTER, on)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetDesktopSettingsAware(on bool) {
	var nilthis *QGuiApplication
	nilthis.SetDesktopSettingsAware(on)
}

// /usr/include/qt/QtGui/qguiapplication.h:146
// index:0
// Public static
// bool desktopSettingsAware()
func (this *QGuiApplication) DesktopSettingsAware() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20desktopSettingsAwareEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_DesktopSettingsAware() {
	var nilthis *QGuiApplication
	nilthis.DesktopSettingsAware()
}

// /usr/include/qt/QtGui/qguiapplication.h:148
// index:0
// Public static
// QInputMethod * inputMethod()
func (this *QGuiApplication) InputMethod() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11inputMethodEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_InputMethod() {
	var nilthis *QGuiApplication
	nilthis.InputMethod()
}

// /usr/include/qt/QtGui/qguiapplication.h:150
// index:0
// Public static
// QPlatformNativeInterface * platformNativeInterface()
func (this *QGuiApplication) PlatformNativeInterface() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23platformNativeInterfaceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_PlatformNativeInterface() {
	var nilthis *QGuiApplication
	nilthis.PlatformNativeInterface()
}

// /usr/include/qt/QtGui/qguiapplication.h:152
// index:0
// Public static
// QFunctionPointer platformFunction(const class QByteArray &)
func (this *QGuiApplication) PlatformFunction(function *qtcore.QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16platformFunctionERK10QByteArray", ffiqt.FFI_TYPE_POINTER, function)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_PlatformFunction(function *qtcore.QByteArray) {
	var nilthis *QGuiApplication
	nilthis.PlatformFunction(function)
}

// /usr/include/qt/QtGui/qguiapplication.h:154
// index:0
// Public static
// void setQuitOnLastWindowClosed(_Bool)
func (this *QGuiApplication) SetQuitOnLastWindowClosed(quit bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication25setQuitOnLastWindowClosedEb", ffiqt.FFI_TYPE_POINTER, quit)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetQuitOnLastWindowClosed(quit bool) {
	var nilthis *QGuiApplication
	nilthis.SetQuitOnLastWindowClosed(quit)
}

// /usr/include/qt/QtGui/qguiapplication.h:155
// index:0
// Public static
// bool quitOnLastWindowClosed()
func (this *QGuiApplication) QuitOnLastWindowClosed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22quitOnLastWindowClosedEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_QuitOnLastWindowClosed() {
	var nilthis *QGuiApplication
	nilthis.QuitOnLastWindowClosed()
}

// /usr/include/qt/QtGui/qguiapplication.h:157
// index:0
// Public static
// Qt::ApplicationState applicationState()
func (this *QGuiApplication) ApplicationState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16applicationStateEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_ApplicationState() {
	var nilthis *QGuiApplication
	nilthis.ApplicationState()
}

// /usr/include/qt/QtGui/qguiapplication.h:159
// index:0
// Public static
// int exec()
func (this *QGuiApplication) Exec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4execEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_Exec() {
	var nilthis *QGuiApplication
	nilthis.Exec()
}

// /usr/include/qt/QtGui/qguiapplication.h:160
// index:0
// Public virtual
// bool notify(class QObject *, class QEvent *)
func (this *QGuiApplication) Notify(arg0 unsafe.Pointer, arg1 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:164
// index:0
// Public
// bool isSessionRestored()
func (this *QGuiApplication) IsSessionRestored() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication17isSessionRestoredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:165
// index:0
// Public
// QString sessionId()
func (this *QGuiApplication) SessionId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication9sessionIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:166
// index:0
// Public
// QString sessionKey()
func (this *QGuiApplication) SessionKey() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication10sessionKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:167
// index:0
// Public
// bool isSavingSession()
func (this *QGuiApplication) IsSavingSession() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication15isSavingSessionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:169
// index:0
// Public static
// bool isFallbackSessionManagementEnabled()
func (this *QGuiApplication) IsFallbackSessionManagementEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication34isFallbackSessionManagementEnabledEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGuiApplication_IsFallbackSessionManagementEnabled() {
	var nilthis *QGuiApplication
	nilthis.IsFallbackSessionManagementEnabled()
}

// /usr/include/qt/QtGui/qguiapplication.h:170
// index:0
// Public static
// void setFallbackSessionManagementEnabled(_Bool)
func (this *QGuiApplication) SetFallbackSessionManagementEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication35setFallbackSessionManagementEnabledEb", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetFallbackSessionManagementEnabled(arg0 bool) {
	var nilthis *QGuiApplication
	nilthis.SetFallbackSessionManagementEnabled(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:173
// index:0
// Public static
// void sync()
func (this *QGuiApplication) Sync() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4syncEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_Sync() {
	var nilthis *QGuiApplication
	nilthis.Sync()
}

// /usr/include/qt/QtGui/qguiapplication.h:175
// index:0
// Public
// void fontDatabaseChanged()
func (this *QGuiApplication) FontDatabaseChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication19fontDatabaseChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:176
// index:0
// Public
// void screenAdded(class QScreen *)
func (this *QGuiApplication) ScreenAdded(screen unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11screenAddedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:177
// index:0
// Public
// void screenRemoved(class QScreen *)
func (this *QGuiApplication) ScreenRemoved(screen unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13screenRemovedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:178
// index:0
// Public
// void primaryScreenChanged(class QScreen *)
func (this *QGuiApplication) PrimaryScreenChanged(screen unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20primaryScreenChangedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:179
// index:0
// Public
// void lastWindowClosed()
func (this *QGuiApplication) LastWindowClosed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16lastWindowClosedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:180
// index:0
// Public
// void focusObjectChanged(class QObject *)
func (this *QGuiApplication) FocusObjectChanged(focusObject unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18focusObjectChangedEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), focusObject)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:181
// index:0
// Public
// void focusWindowChanged(class QWindow *)
func (this *QGuiApplication) FocusWindowChanged(focusWindow unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18focusWindowChangedEP7QWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), focusWindow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:182
// index:0
// Public
// void applicationStateChanged(Qt::ApplicationState)
func (this *QGuiApplication) ApplicationStateChanged(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23applicationStateChangedEN2Qt16ApplicationStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:183
// index:0
// Public
// void layoutDirectionChanged(Qt::LayoutDirection)
func (this *QGuiApplication) LayoutDirectionChanged(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22layoutDirectionChangedEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:185
// index:0
// Public
// void commitDataRequest(class QSessionManager &)
func (this *QGuiApplication) CommitDataRequest(sessionManager *QSessionManager) {
	var convArg0 = sessionManager.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17commitDataRequestER15QSessionManager", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:186
// index:0
// Public
// void saveStateRequest(class QSessionManager &)
func (this *QGuiApplication) SaveStateRequest(sessionManager *QSessionManager) {
	var convArg0 = sessionManager.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16saveStateRequestER15QSessionManager", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:188
// index:0
// Public
// void paletteChanged(const class QPalette &)
func (this *QGuiApplication) PaletteChanged(pal *QPalette) {
	var convArg0 = pal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication14paletteChangedERK8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:189
// index:0
// Public
// void applicationDisplayNameChanged()
func (this *QGuiApplication) ApplicationDisplayNameChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication29applicationDisplayNameChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:192
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QGuiApplication) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:193
// index:0
// Protected virtual
// bool compressEvent(class QEvent *, class QObject *, class QPostEventList *)
func (this *QGuiApplication) CompressEvent(arg0 unsafe.Pointer, receiver unsafe.Pointer, arg2 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13compressEventEP6QEventP7QObjectP14QPostEventList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, receiver, arg2)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
