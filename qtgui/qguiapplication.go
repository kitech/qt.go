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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QGuiApplication) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGuiApplication) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qguiapplication.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGuiApplication(int &, char **, int)

/*
Initializes the window system and constructs an application object with argc command line arguments in argv.

Warning: The data referred to by argc and argv must stay valid for the entire lifetime of the QGuiApplication object. In addition, argc must be greater than zero and argv must contain at least one valid character string.

The global qApp pointer refers to this application object. Only one application object should be created.

This application object must be constructed before any paint devices (including pixmaps, bitmaps etc.).

Note: argc and argv might be changed as Qt removes command line arguments that it recognizes.
*/
func NewQGuiApplication(argc int, argv []string, arg2 int) *QGuiApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGuiApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGuiApplication")
	return gothis
}

// /usr/include/qt/QtGui/qguiapplication.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGuiApplication(int &, char **, int)

/*
Initializes the window system and constructs an application object with argc command line arguments in argv.

Warning: The data referred to by argc and argv must stay valid for the entire lifetime of the QGuiApplication object. In addition, argc must be greater than zero and argv must contain at least one valid character string.

The global qApp pointer refers to this application object. Only one application object should be created.

This application object must be constructed before any paint devices (including pixmaps, bitmaps etc.).

Note: argc and argv might be changed as Qt removes command line arguments that it recognizes.
*/
func NewQGuiApplication__(argc int, argv []string) *QGuiApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	// arg: 2, int=Int, =Invalid,
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGuiApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGuiApplication")
	return gothis
}

// /usr/include/qt/QtGui/qguiapplication.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGuiApplication()

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*
Returns a list of all the windows in the application.

The list is empty if there are no windows.

See also topLevelWindows().
*/
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

/*
Returns a list of the top-level windows in the application.

See also allWindows().
*/
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

/*
Returns the top level window at the given position pos, if any.
*/
func (this *QGuiApplication) TopLevelAt(pos qtcore.QPoint_ITF) *QWindow /*777 QWindow **/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10topLevelAtERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_TopLevelAt(pos qtcore.QPoint_ITF) *QWindow /*777 QWindow **/ {
	var nilthis *QGuiApplication
	rv := nilthis.TopLevelAt(pos)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setWindowIcon(const QIcon &)

/*

 */
func (this *QGuiApplication) SetWindowIcon(icon QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13setWindowIconERK5QIcon", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetWindowIcon(icon QIcon_ITF) {
	var nilthis *QGuiApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtGui/qguiapplication.h:102
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon windowIcon()

/*

 */
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

/*

 */
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

/*
Returns the most recently shown modal window. If no modal windows are visible, this function returns zero.

A modal window is a window which has its modality property set to Qt::WindowModal or Qt::ApplicationModal. A modal window must be closed before the user can continue with other parts of the program.

Modal window are organized in a stack. This function returns the modal window at the top of the stack.

See also Qt::WindowModality and QWindow::setModality().
*/
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

/*
Returns the QWindow that receives events tied to focus, such as key events.
*/
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

/*
Returns the QObject in currently active window that will be final receiver of events tied to focus, such as key events.
*/
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

/*

 */
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

/*
Returns the screen at point, or nullptr if outside of any screen.

The point is in relation to the virtualGeometry() of each set of virtual siblings. If the point maps to more than one set of virtual siblings the first match is returned.

This function was introduced in  Qt 5.10.
*/
func (this *QGuiApplication) ScreenAt(point qtcore.QPoint_ITF) *QScreen /*777 QScreen **/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication8screenAtERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGuiApplication_ScreenAt(point qtcore.QPoint_ITF) *QScreen /*777 QScreen **/ {
	var nilthis *QGuiApplication
	rv := nilthis.ScreenAt(point)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const

/*
Returns the highest screen device pixel ratio found on the system. This is the ratio between physical pixels and device-independent pixels.

Use this function only when you don't know which window you are targeting. If you do know the target window, use QWindow::devicePixelRatio() instead.

See also QWindow::devicePixelRatio().
*/
func (this *QGuiApplication) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qguiapplication.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QCursor * overrideCursor()

/*
Returns the active application override cursor.

This function returns 0 if no application cursor has been defined (i.e. the internal cursor stack is empty).

See also setOverrideCursor() and restoreOverrideCursor().
*/
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

/*
Sets the application override cursor to cursor.

Application override cursors are intended for showing the user that the application is in a special state, for example during an operation that might take some time.

This cursor will be displayed in all the application's widgets until restoreOverrideCursor() or another setOverrideCursor() is called.

Application cursors are stored on an internal stack. setOverrideCursor() pushes the cursor onto the stack, and restoreOverrideCursor() pops the active cursor off the stack. changeOverrideCursor() changes the curently active application override cursor.

Every setOverrideCursor() must eventually be followed by a corresponding restoreOverrideCursor(), otherwise the stack will never be emptied.

Example:


  QApplication::setOverrideCursor(QCursor(Qt::WaitCursor));
  calculateHugeMandelbrot();              // lunch time...
  QApplication::restoreOverrideCursor();



See also overrideCursor(), restoreOverrideCursor(), changeOverrideCursor(), and QWidget::setCursor().
*/
func (this *QGuiApplication) SetOverrideCursor(arg0 QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCursor_PTR() != nil {
		convArg0 = arg0.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication17setOverrideCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetOverrideCursor(arg0 QCursor_ITF) {
	var nilthis *QGuiApplication
	nilthis.SetOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void changeOverrideCursor(const QCursor &)

/*
Changes the currently active application override cursor to cursor.

This function has no effect if setOverrideCursor() was not called.

See also setOverrideCursor(), overrideCursor(), restoreOverrideCursor(), and QWidget::setCursor().
*/
func (this *QGuiApplication) ChangeOverrideCursor(arg0 QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCursor_PTR() != nil {
		convArg0 = arg0.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication20changeOverrideCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_ChangeOverrideCursor(arg0 QCursor_ITF) {
	var nilthis *QGuiApplication
	nilthis.ChangeOverrideCursor(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void restoreOverrideCursor()

/*
Undoes the last setOverrideCursor().

If setOverrideCursor() has been called twice, calling restoreOverrideCursor() will activate the first cursor set. Calling this function a second time restores the original widgets' cursors.

See also setOverrideCursor() and overrideCursor().
*/
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

/*
Returns the default application font.

See also setFont().
*/
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

/*
Changes the default application font to font.

See also font().
*/
func (this *QGuiApplication) SetFont(arg0 QFont_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetFont(arg0 QFont_ITF) {
	var nilthis *QGuiApplication
	nilthis.SetFont(arg0)
}

// /usr/include/qt/QtGui/qguiapplication.h:128
// index:0
// Public static Visibility=Default Availability=Available
// [8] QClipboard * clipboard()

/*
Returns the object for interacting with the clipboard.
*/
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

/*
Returns the default application palette.

See also setPalette().
*/
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

/*
Changes the default application palette to pal.

See also palette().
*/
func (this *QGuiApplication) SetPalette(pal QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg0 = pal.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QGuiApplication_SetPalette(pal QPalette_ITF) {
	var nilthis *QGuiApplication
	nilthis.SetPalette(pal)
}

// /usr/include/qt/QtGui/qguiapplication.h:134
// index:0
// Public static Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers keyboardModifiers()

/*
Returns the current state of the modifier keys on the keyboard. The current state is updated sychronously as the event queue is emptied of events that will spontaneously change the keyboard state (QEvent::KeyPress and QEvent::KeyRelease events).

It should be noted this may not reflect the actual keys held on the input device at the time of calling but rather the modifiers as last reported in one of the above events. If no keys are being held Qt::NoModifier is returned.

See also mouseButtons() and queryKeyboardModifiers().
*/
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

/*
Queries and returns the state of the modifier keys on the keyboard. Unlike keyboardModifiers, this method returns the actual keys held on the input device at the time of calling the method.

It does not rely on the keypress events having been received by this process, which makes it possible to check the modifiers while moving a window, for instance. Note that in most cases, you should use keyboardModifiers(), which is faster and more accurate since it contains the state of the modifiers as they were when the currently processed event was received.

See also keyboardModifiers().
*/
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

/*
Returns the current state of the buttons on the mouse. The current state is updated syncronously as the event queue is emptied of events that will spontaneously change the mouse state (QEvent::MouseButtonPress and QEvent::MouseButtonRelease events).

It should be noted this may not reflect the actual buttons held on the input device at the time of calling but rather the mouse buttons as last reported in one of the above events. If no mouse buttons are being held Qt::NoButton is returned.

See also keyboardModifiers().
*/
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

/*

 */
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

/*

 */
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

/*
Returns true if the application's layout direction is Qt::RightToLeft; otherwise returns false.

See also layoutDirection() and isLeftToRight().
*/
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

/*
Returns true if the application's layout direction is Qt::LeftToRight; otherwise returns false.

See also layoutDirection() and isRightToLeft().
*/
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

/*
Returns the application's style hints.

The style hints encapsulate a set of platform dependent properties such as double click intervals, full width selection and others.

The hints can be used to integrate tighter with the underlying platform.

See also QStyleHints.
*/
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

/*
Sets whether Qt should use the system's standard colors, fonts, etc., to on. By default, this is true.

This function must be called before creating the QGuiApplication object, like this:


  int main(int argc, char *argv[])
  {
      QApplication::setDesktopSettingsAware(false);
      QApplication app(argc, argv);
      ...
      return app.exec();
  }



See also desktopSettingsAware().
*/
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

/*
Returns true if Qt is set to use the system's standard colors, fonts, etc.; otherwise returns false. The default is true.

See also setDesktopSettingsAware().
*/
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

/*
returns the input method.

The input method returns properties about the state and position of the virtual keyboard. It also provides information about the position of the current focused input element.

See also QInputMethod.
*/
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

/*
Returns a function pointer from the platformplugin matching function
*/
func (this *QGuiApplication) PlatformFunction(function qtcore.QByteArray_ITF) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if function != nil && function.QByteArray_PTR() != nil {
		convArg0 = function.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16platformFunctionERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QGuiApplication_PlatformFunction(function qtcore.QByteArray_ITF) unsafe.Pointer /*666*/ {
	var nilthis *QGuiApplication
	rv := nilthis.PlatformFunction(function)
	return rv
}

// /usr/include/qt/QtGui/qguiapplication.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setQuitOnLastWindowClosed(_Bool)

/*

 */
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

/*

 */
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

/*
Returns the current state of the application.

You can react to application state changes to perform actions such as stopping/resuming CPU-intensive tasks, freeing/loading resources or saving/restoring application data.

This function was introduced in  Qt 5.2.
*/
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

/*
Enters the main event loop and waits until exit() is called, and then returns the value that was set to exit() (which is 0 if exit() is called via quit()).

It is necessary to call this function to start event handling. The main event loop receives events from the window system and dispatches these to the application widgets.

Generally, no user interaction can take place before calling exec().

To make your application perform idle processing, e.g., executing a special function whenever there are no pending events, use a QTimer with 0 timeout. More advanced idle processing schemes can be achieved using processEvents().

We recommend that you connect clean-up code to the aboutToQuit() signal, instead of putting it in your application's main() function. This is because, on some platforms, the QApplication::exec() call may not return.

See also quitOnLastWindowClosed, quit(), exit(), processEvents(), and QCoreApplication::exec().
*/
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

/*
Reimplemented from QCoreApplication::notify().
*/
func (this *QGuiApplication) Notify(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication6notifyEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSessionRestored() const

/*
Returns true if the application has been restored from an earlier session; otherwise returns false.

See also sessionId(), commitDataRequest(), and saveStateRequest().
*/
func (this *QGuiApplication) IsSessionRestored() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication17isSessionRestoredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionId() const

/*
Returns the current session's identifier.

If the application has been restored from an earlier session, this identifier is the same as it was in that previous session. The session identifier is guaranteed to be unique both for different applications and for different instances of the same application.

See also isSessionRestored(), sessionKey(), commitDataRequest(), and saveStateRequest().
*/
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
// [8] QString sessionKey() const

/*
Returns the session key in the current session.

If the application has been restored from an earlier session, this key is the same as it was when the previous session ended.

The session key changes every time the session is saved. If the shutdown process is cancelled, another session key will be used when shutting down again.

See also isSessionRestored(), sessionId(), commitDataRequest(), and saveStateRequest().
*/
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
// [1] bool isSavingSession() const

/*
Returns true if the application is currently saving the session; otherwise returns false.

This is true when commitDataRequest() and saveStateRequest() are emitted, but also when the windows are closed afterwards by session management.

This function was introduced in  Qt 5.0.

See also sessionId(), commitDataRequest(), and saveStateRequest().
*/
func (this *QGuiApplication) IsSavingSession() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGuiApplication15isSavingSessionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qguiapplication.h:169
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isFallbackSessionManagementEnabled()

/*
Returns whether QGuiApplication will use fallback session management.

The default is true.

If this is true and the session manager allows user interaction, QGuiApplication will try to close toplevel windows after commitDataRequest() has been emitted. If a window cannot be closed, session shutdown will be canceled and the application will keep running.

Fallback session management only benefits applications that have an "are you sure you want to close this window?" feature or other logic that prevents closing a toplevel window depending on certain conditions, and that do nothing to explicitly implement session management. In applications that do implement session management using the proper session management API, fallback session management interferes and may break session management logic.

Warning: If all windows are closed due to fallback session management and quitOnLastWindowClosed() is true, the application will quit before it is explicitly instructed to quit through the platform's session management protocol. That violation of protocol may prevent the platform session manager from saving application state.

This function was introduced in  Qt 5.6.

See also setFallbackSessionManagementEnabled(), QSessionManager::allowsInteraction(), saveStateRequest(), commitDataRequest(), and Session Management.
*/
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

/*
Sets whether QGuiApplication will use fallback session management to enabled.

This function was introduced in  Qt 5.6.

See also isFallbackSessionManagementEnabled().
*/
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

/*
Function that can be used to sync Qt state with the Window Systems state.

This function will first empty Qts events by calling QCoreApplication::processEvents(), then the platform plugin will sync up with the windowsystem, and finally Qts events will be delived by another call to QCoreApplication::processEvents();

This function is timeconsuming and its use is discouraged.

This function was introduced in  Qt 5.2.
*/
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

/*
This signal is emitted when application fonts are loaded or removed.

See also QFontDatabase::addApplicationFont(), QFontDatabase::addApplicationFontFromData(), QFontDatabase::removeAllApplicationFonts(), and QFontDatabase::removeApplicationFont().
*/
func (this *QGuiApplication) FontDatabaseChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication19fontDatabaseChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenAdded(QScreen *)

/*
This signal is emitted whenever a new screen screen has been added to the system.

See also screens(), primaryScreen, and screenRemoved().
*/
func (this *QGuiApplication) ScreenAdded(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication11screenAddedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenRemoved(QScreen *)

/*
This signal is emitted whenever a screen is removed from the system. It provides an opportunity to manage the windows on the screen before Qt falls back to moving them to the primary screen.

This function was introduced in  Qt 5.4.

See also screens(), screenAdded(), QObject::destroyed(), and QWindow::setScreen().
*/
func (this *QGuiApplication) ScreenRemoved(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication13screenRemovedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void primaryScreenChanged(QScreen *)

/*

 */
func (this *QGuiApplication) PrimaryScreenChanged(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication20primaryScreenChangedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lastWindowClosed()

/*
This signal is emitted from exec() when the last visible primary window (i.e. window with no parent) is closed.

By default, QGuiApplication quits after this signal is emitted. This feature can be turned off by setting quitOnLastWindowClosed to false.

See also QWindow::close() and QWindow::isTopLevel().
*/
func (this *QGuiApplication) LastWindowClosed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16lastWindowClosedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusObjectChanged(QObject *)

/*
This signal is emitted when final receiver of events tied to focus is changed. focusObject is the new receiver.

See also focusObject().
*/
func (this *QGuiApplication) FocusObjectChanged(focusObject qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if focusObject != nil && focusObject.QObject_PTR() != nil {
		convArg0 = focusObject.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication18focusObjectChangedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusWindowChanged(QWindow *)

/*
This signal is emitted when the focused window changes. focusWindow is the new focused window.

See also focusWindow().
*/
func (this *QGuiApplication) FocusWindowChanged(focusWindow QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if focusWindow != nil && focusWindow.QWindow_PTR() != nil {
		convArg0 = focusWindow.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication18focusWindowChangedEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationStateChanged(Qt::ApplicationState)

/*
This signal is emitted when the state of the application changes.

This function was introduced in  Qt 5.2.

See also applicationState().
*/
func (this *QGuiApplication) ApplicationStateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication23applicationStateChangedEN2Qt16ApplicationStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void layoutDirectionChanged(Qt::LayoutDirection)

/*

 */
func (this *QGuiApplication) LayoutDirectionChanged(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication22layoutDirectionChangedEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commitDataRequest(QSessionManager &)

/*
This signal deals with session management. It is emitted when the QSessionManager wants the application to commit all its data.

Usually this means saving all open files, after getting permission from the user. Furthermore you may want to provide a means by which the user can cancel the shutdown.

You should not exit the application within this signal. Instead, the session manager may or may not do this afterwards, depending on the context.

Warning: Within this signal, no user interaction is possible, unless you ask the manager for explicit permission. See QSessionManager::allowsInteraction() and QSessionManager::allowsErrorInteraction() for details and example usage.

Note: You should use Qt::DirectConnection when connecting to this signal.

This function was introduced in  Qt 4.2.

See also setFallbackSessionManagementEnabled(), isSessionRestored(), sessionId(), saveStateRequest(), and Session Management.
*/
func (this *QGuiApplication) CommitDataRequest(sessionManager QSessionManager_ITF) {
	var convArg0 unsafe.Pointer
	if sessionManager != nil && sessionManager.QSessionManager_PTR() != nil {
		convArg0 = sessionManager.QSessionManager_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication17commitDataRequestER15QSessionManager", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void saveStateRequest(QSessionManager &)

/*
This signal deals with session management. It is invoked when the session manager wants the application to preserve its state for a future session.

For example, a text editor would create a temporary file that includes the current contents of its edit buffers, the location of the cursor and other aspects of the current editing session.

You should never exit the application within this signal. Instead, the session manager may or may not do this afterwards, depending on the context. Futhermore, most session managers will very likely request a saved state immediately after the application has been started. This permits the session manager to learn about the application's restart policy.

Warning: Within this signal, no user interaction is possible, unless you ask the manager for explicit permission. See QSessionManager::allowsInteraction() and QSessionManager::allowsErrorInteraction() for details.

Note: You should use Qt::DirectConnection when connecting to this signal.

This function was introduced in  Qt 4.2.

See also isSessionRestored(), sessionId(), commitDataRequest(), and Session Management.
*/
func (this *QGuiApplication) SaveStateRequest(sessionManager QSessionManager_ITF) {
	var convArg0 unsafe.Pointer
	if sessionManager != nil && sessionManager.QSessionManager_PTR() != nil {
		convArg0 = sessionManager.QSessionManager_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication16saveStateRequestER15QSessionManager", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paletteChanged(const QPalette &)

/*
This signal is emitted when the palette of the application changes.

This function was introduced in  Qt 5.4.

See also palette().
*/
func (this *QGuiApplication) PaletteChanged(pal QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg0 = pal.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication14paletteChangedERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationDisplayNameChanged()

/*

 */
func (this *QGuiApplication) ApplicationDisplayNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGuiApplication29applicationDisplayNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qguiapplication.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QGuiApplication) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
