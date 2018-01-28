package qtgui

// /usr/include/qt/QtGui/qguiapplication.h
// #include <qguiapplication.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QGuiApplication) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qguiapplication.h:87
// index:0
// Public
// void QGuiApplication(int &, char **, int)
func NewQGuiApplication(argc int, argv []string, arg2 int) *QGuiApplication {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, arg2)
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
// void setApplicationDisplayName(const QString &)
func (this *QGuiApplication) SetApplicationDisplayName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication25setApplicationDisplayNameERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QGuiApplication) ApplicationDisplayName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22applicationDisplayNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QGuiApplication_ApplicationDisplayName() *qtcore.QString /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.ApplicationDisplayName()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:94
// index:0
// Public static
// void setDesktopFileName(const QString &)
func (this *QGuiApplication) SetDesktopFileName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18setDesktopFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QGuiApplication) DesktopFileName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15desktopFileNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QGuiApplication_DesktopFileName() *qtcore.QString /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.DesktopFileName()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:99
// index:0
// Public static
// QWindow * topLevelAt(const QPoint &)
func (this *QGuiApplication) TopLevelAt(pos *qtcore.QPoint) *QWindow /*777 QWindow **/ {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10topLevelAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_TopLevelAt(pos *qtcore.QPoint) *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.TopLevelAt(pos)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:101
// index:0
// Public static
// void setWindowIcon(const QIcon &)
func (this *QGuiApplication) SetWindowIcon(icon *QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QGuiApplication) WindowIcon() *QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10windowIconEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QGuiApplication_WindowIcon() *QIcon /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.WindowIcon()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:104
// index:0
// Public static
// QString platformName()
func (this *QGuiApplication) PlatformName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication12platformNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QGuiApplication_PlatformName() *qtcore.QString /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.PlatformName()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:106
// index:0
// Public static
// QWindow * modalWindow()
func (this *QGuiApplication) ModalWindow() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11modalWindowEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_ModalWindow() *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.ModalWindow()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:108
// index:0
// Public static
// QWindow * focusWindow()
func (this *QGuiApplication) FocusWindow() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11focusWindowEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_FocusWindow() *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.FocusWindow()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:109
// index:0
// Public static
// QObject * focusObject()
func (this *QGuiApplication) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11focusObjectEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_FocusObject() *qtcore.QObject /*777 QObject **/ {
	var nilthis *QGuiApplication
	rv := nilthis.FocusObject()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:111
// index:0
// Public static
// QScreen * primaryScreen()
func (this *QGuiApplication) PrimaryScreen() *QScreen /*777 QScreen **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13primaryScreenEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_PrimaryScreen() *QScreen /*777 QScreen **/ {
	var nilthis *QGuiApplication
	rv := nilthis.PrimaryScreen()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:113
// index:0
// Public static
// QScreen * screenAt(const QPoint &)
func (this *QGuiApplication) ScreenAt(point *qtcore.QPoint) *QScreen /*777 QScreen **/ {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication8screenAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_ScreenAt(point *qtcore.QPoint) *QScreen /*777 QScreen **/ {
	var nilthis *QGuiApplication
	rv := nilthis.ScreenAt(point)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:115
// index:0
// Public
// qreal devicePixelRatio()
func (this *QGuiApplication) DevicePixelRatio() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication16devicePixelRatioEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qguiapplication.h:118
// index:0
// Public static
// QCursor * overrideCursor()
func (this *QGuiApplication) OverrideCursor() *QCursor /*777 QCursor **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication14overrideCursorEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_OverrideCursor() *QCursor /*777 QCursor **/ {
	var nilthis *QGuiApplication
	rv := nilthis.OverrideCursor()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:119
// index:0
// Public static
// void setOverrideCursor(const QCursor &)
func (this *QGuiApplication) SetOverrideCursor(arg0 *QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17setOverrideCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QGuiApplication_SetOverrideCursor(arg0 *QCursor) {
	var nilthis *QGuiApplication
	nilthis.SetOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:120
// index:0
// Public static
// void changeOverrideCursor(const QCursor &)
func (this *QGuiApplication) ChangeOverrideCursor(arg0 *QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20changeOverrideCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QGuiApplication) Font() *QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4fontEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QGuiApplication_Font() *QFont /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.Font()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:125
// index:0
// Public static
// void setFont(const QFont &)
func (this *QGuiApplication) SetFont(arg0 *QFont) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QGuiApplication) Clipboard() *QClipboard /*777 QClipboard **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication9clipboardEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQClipboardFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_Clipboard() *QClipboard /*777 QClipboard **/ {
	var nilthis *QGuiApplication
	rv := nilthis.Clipboard()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:131
// index:0
// Public static
// QPalette palette()
func (this *QGuiApplication) Palette() *QPalette /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication7paletteEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QGuiApplication_Palette() *QPalette /*123*/ {
	var nilthis *QGuiApplication
	rv := nilthis.Palette()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:132
// index:0
// Public static
// void setPalette(const QPalette &)
func (this *QGuiApplication) SetPalette(pal *QPalette) {
	var convArg0 = pal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10setPaletteERK8QPalette", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QGuiApplication) KeyboardModifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17keyboardModifiersEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QGuiApplication_KeyboardModifiers() int {
	var nilthis *QGuiApplication
	rv := nilthis.KeyboardModifiers()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:135
// index:0
// Public static
// Qt::KeyboardModifiers queryKeyboardModifiers()
func (this *QGuiApplication) QueryKeyboardModifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22queryKeyboardModifiersEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QGuiApplication_QueryKeyboardModifiers() int {
	var nilthis *QGuiApplication
	rv := nilthis.QueryKeyboardModifiers()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:136
// index:0
// Public static
// Qt::MouseButtons mouseButtons()
func (this *QGuiApplication) MouseButtons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication12mouseButtonsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QGuiApplication_MouseButtons() int {
	var nilthis *QGuiApplication
	rv := nilthis.MouseButtons()
	return rv
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
func (this *QGuiApplication) LayoutDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QGuiApplication_LayoutDirection() int {
	var nilthis *QGuiApplication
	rv := nilthis.LayoutDirection()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:141
// index:0
// Public static inline
// bool isRightToLeft()
func (this *QGuiApplication) IsRightToLeft() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13isRightToLeftEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QGuiApplication_IsRightToLeft() bool {
	var nilthis *QGuiApplication
	rv := nilthis.IsRightToLeft()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:142
// index:0
// Public static inline
// bool isLeftToRight()
func (this *QGuiApplication) IsLeftToRight() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13isLeftToRightEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QGuiApplication_IsLeftToRight() bool {
	var nilthis *QGuiApplication
	rv := nilthis.IsLeftToRight()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:144
// index:0
// Public static
// QStyleHints * styleHints()
func (this *QGuiApplication) StyleHints() *QStyleHints /*777 QStyleHints **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication10styleHintsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStyleHintsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_StyleHints() *QStyleHints /*777 QStyleHints **/ {
	var nilthis *QGuiApplication
	rv := nilthis.StyleHints()
	return rv
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
func (this *QGuiApplication) DesktopSettingsAware() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20desktopSettingsAwareEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QGuiApplication_DesktopSettingsAware() bool {
	var nilthis *QGuiApplication
	rv := nilthis.DesktopSettingsAware()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:148
// index:0
// Public static
// QInputMethod * inputMethod()
func (this *QGuiApplication) InputMethod() *QInputMethod /*777 QInputMethod **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11inputMethodEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQInputMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGuiApplication_InputMethod() *QInputMethod /*777 QInputMethod **/ {
	var nilthis *QGuiApplication
	rv := nilthis.InputMethod()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:152
// index:0
// Public static
// QFunctionPointer platformFunction(const QByteArray &)
func (this *QGuiApplication) PlatformFunction(function *qtcore.QByteArray) unsafe.Pointer /*666*/ {
	var convArg0 = function.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16platformFunctionERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func QGuiApplication_PlatformFunction(function *qtcore.QByteArray) unsafe.Pointer /*666*/ {
	var nilthis *QGuiApplication
	rv := nilthis.PlatformFunction(function)
	return rv
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
func (this *QGuiApplication) QuitOnLastWindowClosed() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22quitOnLastWindowClosedEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QGuiApplication_QuitOnLastWindowClosed() bool {
	var nilthis *QGuiApplication
	rv := nilthis.QuitOnLastWindowClosed()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:157
// index:0
// Public static
// Qt::ApplicationState applicationState()
func (this *QGuiApplication) ApplicationState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16applicationStateEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QGuiApplication_ApplicationState() int {
	var nilthis *QGuiApplication
	rv := nilthis.ApplicationState()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:159
// index:0
// Public static
// int exec()
func (this *QGuiApplication) Exec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication4execEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QGuiApplication_Exec() int {
	var nilthis *QGuiApplication
	rv := nilthis.Exec()
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:160
// index:0
// Public virtual
// bool notify(QObject *, QEvent *)
func (this *QGuiApplication) Notify(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:164
// index:0
// Public
// bool isSessionRestored()
func (this *QGuiApplication) IsSessionRestored() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication17isSessionRestoredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:165
// index:0
// Public
// QString sessionId()
func (this *QGuiApplication) SessionId() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication9sessionIdEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qguiapplication.h:166
// index:0
// Public
// QString sessionKey()
func (this *QGuiApplication) SessionKey() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication10sessionKeyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qguiapplication.h:167
// index:0
// Public
// bool isSavingSession()
func (this *QGuiApplication) IsSavingSession() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGuiApplication15isSavingSessionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:169
// index:0
// Public static
// bool isFallbackSessionManagementEnabled()
func (this *QGuiApplication) IsFallbackSessionManagementEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication34isFallbackSessionManagementEnabledEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QGuiApplication_IsFallbackSessionManagementEnabled() bool {
	var nilthis *QGuiApplication
	rv := nilthis.IsFallbackSessionManagementEnabled()
	return rv
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
// void screenAdded(QScreen *)
func (this *QGuiApplication) ScreenAdded(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication11screenAddedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:177
// index:0
// Public
// void screenRemoved(QScreen *)
func (this *QGuiApplication) ScreenRemoved(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication13screenRemovedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:178
// index:0
// Public
// void primaryScreenChanged(QScreen *)
func (this *QGuiApplication) PrimaryScreenChanged(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication20primaryScreenChangedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
// void focusObjectChanged(QObject *)
func (this *QGuiApplication) FocusObjectChanged(focusObject *qtcore.QObject /*777 QObject **/) {
	var convArg0 = focusObject.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18focusObjectChangedEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:181
// index:0
// Public
// void focusWindowChanged(QWindow *)
func (this *QGuiApplication) FocusWindowChanged(focusWindow *QWindow /*777 QWindow **/) {
	var convArg0 = focusWindow.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication18focusWindowChangedEP7QWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:182
// index:0
// Public
// void applicationStateChanged(Qt::ApplicationState)
func (this *QGuiApplication) ApplicationStateChanged(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication23applicationStateChangedEN2Qt16ApplicationStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:183
// index:0
// Public
// void layoutDirectionChanged(Qt::LayoutDirection)
func (this *QGuiApplication) LayoutDirectionChanged(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication22layoutDirectionChangedEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:185
// index:0
// Public
// void commitDataRequest(QSessionManager &)
func (this *QGuiApplication) CommitDataRequest(sessionManager *QSessionManager) {
	var convArg0 = sessionManager.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication17commitDataRequestER15QSessionManager", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:186
// index:0
// Public
// void saveStateRequest(QSessionManager &)
func (this *QGuiApplication) SaveStateRequest(sessionManager *QSessionManager) {
	var convArg0 = sessionManager.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication16saveStateRequestER15QSessionManager", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:188
// index:0
// Public
// void paletteChanged(const QPalette &)
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
// bool event(QEvent *)
func (this *QGuiApplication) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGuiApplication5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
