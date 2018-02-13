package qtgui

// /usr/include/qt/QtGui/qguiapplication.h
// #include <qguiapplication.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QGuiApplication) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QGuiApplication struct {
	*qtcore.QCoreApplication
}
type QGuiApplication_ITF interface {
	qtcore.QCoreApplication_ITF
	QGuiApplication_PTR() *QGuiApplication
}

func (ptr *QGuiApplication) QGuiApplication_PTR() *QGuiApplication { return ptr }

func (this *QGuiApplication) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QCoreApplication.GetCthis()
	}
}
func (this *QGuiApplication) SetCthis(cthis unsafe.Pointer) {
	this.QCoreApplication = qtcore.NewQCoreApplicationFromPointer(cthis)
}
func NewQGuiApplicationFromPointer(cthis unsafe.Pointer) *QGuiApplication {
	bcthis0 := qtcore.NewQCoreApplicationFromPointer(cthis)
	return &QGuiApplication{bcthis0}
}
func (*QGuiApplication) NewFromPointer(cthis unsafe.Pointer) *QGuiApplication {
	return NewQGuiApplicationFromPointer(cthis)
}

// /usr/include/qt/QtGui/qguiapplication.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGuiApplication) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qguiapplication.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGuiApplication(int &, char **, int)
func NewQGuiApplication(argc int, argv []string, arg2 int) *QGuiApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGuiApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qguiapplication.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGuiApplication()
func DeleteQGuiApplication(this *QGuiApplication) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplicationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qguiapplication.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setApplicationDisplayName(const QString &)
func (this *QGuiApplication) SetApplicationDisplayName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication25setApplicationDisplayNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetApplicationDisplayName(name string) {
	var nilthis *QGuiApplication
	nilthis.SetApplicationDisplayName(name)
}

// /usr/include/qt/QtGui/qguiapplication.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString applicationDisplayName()
func (this *QGuiApplication) ApplicationDisplayName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication22applicationDisplayNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QGuiApplication_ApplicationDisplayName() string {
	var nilthis *QGuiApplication
	rv := nilthis.ApplicationDisplayName()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDesktopFileName(const QString &)
func (this *QGuiApplication) SetDesktopFileName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication18setDesktopFileNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetDesktopFileName(name string) {
	var nilthis *QGuiApplication
	nilthis.SetDesktopFileName(name)
}

// /usr/include/qt/QtGui/qguiapplication.h:95
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString desktopFileName()
func (this *QGuiApplication) DesktopFileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication15desktopFileNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QGuiApplication_DesktopFileName() string {
	var nilthis *QGuiApplication
	rv := nilthis.DesktopFileName()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QWindowList allWindows()
func (this *QGuiApplication) AllWindows() *QWindowList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10allWindowsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWindowListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QGuiApplication_AllWindows() *QWindowList /*667*/ {
	var nilthis *QGuiApplication
	rv := nilthis.AllWindows()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QWindowList topLevelWindows()
func (this *QGuiApplication) TopLevelWindows() *QWindowList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication15topLevelWindowsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWindowListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QGuiApplication_TopLevelWindows() *QWindowList /*667*/ {
	var nilthis *QGuiApplication
	rv := nilthis.TopLevelWindows()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * topLevelAt(const QPoint &)
func (this *QGuiApplication) TopLevelAt(pos *qtcore.QPoint) *QWindow /*777 QWindow **/ {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10topLevelAtERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_TopLevelAt(pos *qtcore.QPoint) *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.TopLevelAt(pos)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setWindowIcon(const QIcon &)
func (this *QGuiApplication) SetWindowIcon(icon *QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13setWindowIconERK5QIcon", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetWindowIcon(icon *QIcon) {
	var nilthis *QGuiApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtGui/qguiapplication.h:102
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon windowIcon()
func (this *QGuiApplication) WindowIcon() *QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10windowIconEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QGuiApplication_WindowIcon() *QIcon /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.WindowIcon()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:104
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString platformName()
func (this *QGuiApplication) PlatformName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication12platformNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QGuiApplication_PlatformName() string {
	var nilthis *QGuiApplication
	rv := nilthis.PlatformName()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * modalWindow()
func (this *QGuiApplication) ModalWindow() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication11modalWindowEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_ModalWindow() *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.ModalWindow()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:108
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * focusWindow()
func (this *QGuiApplication) FocusWindow() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication11focusWindowEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_FocusWindow() *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.FocusWindow()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:109
// index:0
// Public static Visibility=Default Availability=Available
// [8] QObject * focusObject()
func (this *QGuiApplication) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication11focusObjectEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_FocusObject() *qtcore.QObject /*777 QObject **/ {
	var nilthis *QGuiApplication
	rv := nilthis.FocusObject()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:111
// index:0
// Public static Visibility=Default Availability=Available
// [8] QScreen * primaryScreen()
func (this *QGuiApplication) PrimaryScreen() *QScreen /*777 QScreen **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13primaryScreenEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_PrimaryScreen() *QScreen /*777 QScreen **/ {
	var nilthis *QGuiApplication
	rv := nilthis.PrimaryScreen()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [8] QScreen * screenAt(const QPoint &)
func (this *QGuiApplication) ScreenAt(point *qtcore.QPoint) *QScreen /*777 QScreen **/ {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication8screenAtERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_ScreenAt(point *qtcore.QPoint) *QScreen /*777 QScreen **/ {
	var nilthis *QGuiApplication
	rv := nilthis.ScreenAt(point)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio()
func (this *QGuiApplication) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qguiapplication.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QCursor * overrideCursor()
func (this *QGuiApplication) OverrideCursor() *QCursor /*777 QCursor **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication14overrideCursorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_OverrideCursor() *QCursor /*777 QCursor **/ {
	var nilthis *QGuiApplication
	rv := nilthis.OverrideCursor()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:119
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setOverrideCursor(const QCursor &)
func (this *QGuiApplication) SetOverrideCursor(arg0 *QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication17setOverrideCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetOverrideCursor(arg0 *QCursor) {
	var nilthis *QGuiApplication
	nilthis.SetOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void changeOverrideCursor(const QCursor &)
func (this *QGuiApplication) ChangeOverrideCursor(arg0 *QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication20changeOverrideCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_ChangeOverrideCursor(arg0 *QCursor) {
	var nilthis *QGuiApplication
	nilthis.ChangeOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void restoreOverrideCursor()
func (this *QGuiApplication) RestoreOverrideCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication21restoreOverrideCursorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_RestoreOverrideCursor() {
	var nilthis *QGuiApplication
	nilthis.RestoreOverrideCursor()
}

// /usr/include/qt/QtGui/qguiapplication.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont font()
func (this *QGuiApplication) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication4fontEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}
func QGuiApplication_Font() *QFont /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.Font()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QGuiApplication) SetFont(arg0 *QFont) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetFont(arg0 *QFont) {
	var nilthis *QGuiApplication
	nilthis.SetFont(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:128
// index:0
// Public static Visibility=Default Availability=Available
// [8] QClipboard * clipboard()
func (this *QGuiApplication) Clipboard() *QClipboard /*777 QClipboard **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication9clipboardEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQClipboardFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_Clipboard() *QClipboard /*777 QClipboard **/ {
	var nilthis *QGuiApplication
	rv := nilthis.Clipboard()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [16] QPalette palette()
func (this *QGuiApplication) Palette() *QPalette /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication7paletteEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPalette)
	return rv2
}
func QGuiApplication_Palette() *QPalette /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.Palette()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)
func (this *QGuiApplication) SetPalette(pal *QPalette) {
	var convArg0 = pal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetPalette(pal *QPalette) {
	var nilthis *QGuiApplication
	nilthis.SetPalette(pal)
}

// /usr/include/qt/QtGui/qguiapplication.h:134
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers keyboardModifiers()
func (this *QGuiApplication) KeyboardModifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication17keyboardModifiersEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QGuiApplication_KeyboardModifiers() int {
	var nilthis *QGuiApplication
	rv := nilthis.KeyboardModifiers()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:135
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers queryKeyboardModifiers()
func (this *QGuiApplication) QueryKeyboardModifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication22queryKeyboardModifiersEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QGuiApplication_QueryKeyboardModifiers() int {
	var nilthis *QGuiApplication
	rv := nilthis.QueryKeyboardModifiers()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:136
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::MouseButtons mouseButtons()
func (this *QGuiApplication) MouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication12mouseButtonsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QGuiApplication_MouseButtons() int {
	var nilthis *QGuiApplication
	rv := nilthis.MouseButtons()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:138
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)
func (this *QGuiApplication) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, direction)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetLayoutDirection(direction int) {
	var nilthis *QGuiApplication
	nilthis.SetLayoutDirection(direction)
}

// /usr/include/qt/QtGui/qguiapplication.h:139
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection()
func (this *QGuiApplication) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication15layoutDirectionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QGuiApplication_LayoutDirection() int {
	var nilthis *QGuiApplication
	rv := nilthis.LayoutDirection()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:141
// index:0
// Public static inline Visibility=Default Availability=Available
// [1] bool isRightToLeft()
func (this *QGuiApplication) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13isRightToLeftEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGuiApplication_IsRightToLeft() bool {
	var nilthis *QGuiApplication
	rv := nilthis.IsRightToLeft()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:142
// index:0
// Public static inline Visibility=Default Availability=Available
// [1] bool isLeftToRight()
func (this *QGuiApplication) IsLeftToRight() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13isLeftToRightEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGuiApplication_IsLeftToRight() bool {
	var nilthis *QGuiApplication
	rv := nilthis.IsLeftToRight()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:144
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStyleHints * styleHints()
func (this *QGuiApplication) StyleHints() *QStyleHints /*777 QStyleHints **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10styleHintsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleHintsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_StyleHints() *QStyleHints /*777 QStyleHints **/ {
	var nilthis *QGuiApplication
	rv := nilthis.StyleHints()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:145
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDesktopSettingsAware(_Bool)
func (this *QGuiApplication) SetDesktopSettingsAware(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication23setDesktopSettingsAwareEb", qtrt.FFI_TYPE_POINTER, on)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetDesktopSettingsAware(on bool) {
	var nilthis *QGuiApplication
	nilthis.SetDesktopSettingsAware(on)
}

// /usr/include/qt/QtGui/qguiapplication.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool desktopSettingsAware()
func (this *QGuiApplication) DesktopSettingsAware() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication20desktopSettingsAwareEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGuiApplication_DesktopSettingsAware() bool {
	var nilthis *QGuiApplication
	rv := nilthis.DesktopSettingsAware()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:148
// index:0
// Public static Visibility=Default Availability=Available
// [8] QInputMethod * inputMethod()
func (this *QGuiApplication) InputMethod() *QInputMethod /*777 QInputMethod **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication11inputMethodEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQInputMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_InputMethod() *QInputMethod /*777 QInputMethod **/ {
	var nilthis *QGuiApplication
	rv := nilthis.InputMethod()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer platformFunction(const QByteArray &)
func (this *QGuiApplication) PlatformFunction(function *qtcore.QByteArray) unsafe.Pointer /*666*/ {
	var convArg0 = function.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16platformFunctionERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QGuiApplication_PlatformFunction(function *qtcore.QByteArray) unsafe.Pointer /*666*/ {
	var nilthis *QGuiApplication
	rv := nilthis.PlatformFunction(function)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setQuitOnLastWindowClosed(_Bool)
func (this *QGuiApplication) SetQuitOnLastWindowClosed(quit bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication25setQuitOnLastWindowClosedEb", qtrt.FFI_TYPE_POINTER, quit)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetQuitOnLastWindowClosed(quit bool) {
	var nilthis *QGuiApplication
	nilthis.SetQuitOnLastWindowClosed(quit)
}

// /usr/include/qt/QtGui/qguiapplication.h:155
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool quitOnLastWindowClosed()
func (this *QGuiApplication) QuitOnLastWindowClosed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication22quitOnLastWindowClosedEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGuiApplication_QuitOnLastWindowClosed() bool {
	var nilthis *QGuiApplication
	rv := nilthis.QuitOnLastWindowClosed()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:157
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::ApplicationState applicationState()
func (this *QGuiApplication) ApplicationState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16applicationStateEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QGuiApplication_ApplicationState() int {
	var nilthis *QGuiApplication
	rv := nilthis.ApplicationState()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:159
// index:0
// Public static Visibility=Default Availability=Available
// [4] int exec()
func (this *QGuiApplication) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication4execEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QGuiApplication_Exec() int {
	var nilthis *QGuiApplication
	rv := nilthis.Exec()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:160
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool notify(QObject *, QEvent *)
func (this *QGuiApplication) Notify(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication6notifyEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSessionRestored()
func (this *QGuiApplication) IsSessionRestored() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication17isSessionRestoredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionId()
func (this *QGuiApplication) SessionId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication9sessionIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qguiapplication.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionKey()
func (this *QGuiApplication) SessionKey() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication10sessionKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qguiapplication.h:167
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSavingSession()
func (this *QGuiApplication) IsSavingSession() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication15isSavingSessionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:169
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isFallbackSessionManagementEnabled()
func (this *QGuiApplication) IsFallbackSessionManagementEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication34isFallbackSessionManagementEnabledEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGuiApplication_IsFallbackSessionManagementEnabled() bool {
	var nilthis *QGuiApplication
	rv := nilthis.IsFallbackSessionManagementEnabled()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:170
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFallbackSessionManagementEnabled(_Bool)
func (this *QGuiApplication) SetFallbackSessionManagementEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication35setFallbackSessionManagementEnabledEb", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetFallbackSessionManagementEnabled(arg0 bool) {
	var nilthis *QGuiApplication
	nilthis.SetFallbackSessionManagementEnabled(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:173
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sync()
func (this *QGuiApplication) Sync() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication4syncEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_Sync() {
	var nilthis *QGuiApplication
	nilthis.Sync()
}

// /usr/include/qt/QtGui/qguiapplication.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fontDatabaseChanged()
func (this *QGuiApplication) FontDatabaseChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication19fontDatabaseChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenAdded(QScreen *)
func (this *QGuiApplication) ScreenAdded(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication11screenAddedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenRemoved(QScreen *)
func (this *QGuiApplication) ScreenRemoved(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13screenRemovedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void primaryScreenChanged(QScreen *)
func (this *QGuiApplication) PrimaryScreenChanged(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication20primaryScreenChangedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lastWindowClosed()
func (this *QGuiApplication) LastWindowClosed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16lastWindowClosedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusObjectChanged(QObject *)
func (this *QGuiApplication) FocusObjectChanged(focusObject *qtcore.QObject /*777 QObject **/) {
	var convArg0 = focusObject.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication18focusObjectChangedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusWindowChanged(QWindow *)
func (this *QGuiApplication) FocusWindowChanged(focusWindow *QWindow /*777 QWindow **/) {
	var convArg0 = focusWindow.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication18focusWindowChangedEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationStateChanged(Qt::ApplicationState)
func (this *QGuiApplication) ApplicationStateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication23applicationStateChangedEN2Qt16ApplicationStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void layoutDirectionChanged(Qt::LayoutDirection)
func (this *QGuiApplication) LayoutDirectionChanged(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication22layoutDirectionChangedEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commitDataRequest(QSessionManager &)
func (this *QGuiApplication) CommitDataRequest(sessionManager *QSessionManager) {
	var convArg0 = sessionManager.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication17commitDataRequestER15QSessionManager", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void saveStateRequest(QSessionManager &)
func (this *QGuiApplication) SaveStateRequest(sessionManager *QSessionManager) {
	var convArg0 = sessionManager.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16saveStateRequestER15QSessionManager", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paletteChanged(const QPalette &)
func (this *QGuiApplication) PaletteChanged(pal *QPalette) {
	var convArg0 = pal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication14paletteChangedERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationDisplayNameChanged()
func (this *QGuiApplication) ApplicationDisplayNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication29applicationDisplayNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QGuiApplication) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

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
}

//  keep block end
