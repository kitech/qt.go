package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

// void exposeEvent(QExposeEvent *)
func (this *QWindow) InheritExposeEvent(f func(arg0 *QExposeEvent /*777 QExposeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "exposeEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QWindow) InheritResizeEvent(f func(arg0 *QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void moveEvent(QMoveEvent *)
func (this *QWindow) InheritMoveEvent(f func(arg0 *QMoveEvent /*777 QMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QWindow) InheritFocusInEvent(f func(arg0 *QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QWindow) InheritFocusOutEvent(f func(arg0 *QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QWindow) InheritShowEvent(f func(arg0 *QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QWindow) InheritHideEvent(f func(arg0 *QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// bool event(QEvent *)
func (this *QWindow) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QWindow) InheritKeyPressEvent(f func(arg0 *QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QWindow) InheritKeyReleaseEvent(f func(arg0 *QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QWindow) InheritMousePressEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QWindow) InheritMouseReleaseEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QWindow) InheritMouseDoubleClickEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QWindow) InheritMouseMoveEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QWindow) InheritWheelEvent(f func(arg0 *QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void touchEvent(QTouchEvent *)
func (this *QWindow) InheritTouchEvent(f func(arg0 *QTouchEvent /*777 QTouchEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "touchEvent", f)
}

// void tabletEvent(QTabletEvent *)
func (this *QWindow) InheritTabletEvent(f func(arg0 *QTabletEvent /*777 QTabletEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabletEvent", f)
}

// bool nativeEvent(const QByteArray &, void *, long *)
func (this *QWindow) InheritNativeEvent(f func(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool) {
	qtrt.SetAllInheritCallback(this, "nativeEvent", f)
}

/*

 */
type QWindow struct {
	*qtcore.QObject
	*QSurface
}
type QWindow_ITF interface {
	qtcore.QObject_ITF
	QSurface_ITF
	QWindow_PTR() *QWindow
}

func (ptr *QWindow) QWindow_PTR() *QWindow { return ptr }

func (this *QWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWindow) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QSurface = NewQSurfaceFromPointer(cthis)
}
func NewQWindowFromPointer(cthis unsafe.Pointer) *QWindow {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQSurfaceFromPointer(cthis)
	return &QWindow{bcthis0, bcthis1}
}
func (*QWindow) NewFromPointer(cthis unsafe.Pointer) *QWindow {
	return NewQWindowFromPointer(cthis)
}

// /usr/include/qt/QtGui/qwindow.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWindow(QScreen *)

/*
Creates a window as a top level on the targetScreen.

The window is not shown until setVisible(true), show(), or similar is called.

See also setScreen().
*/
func (*QWindow) NewForInherit(screen QScreen_ITF /*777 QScreen **/) *QWindow {
	return NewQWindow(screen)
}
func NewQWindow(screen QScreen_ITF /*777 QScreen **/) *QWindow {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowC2EP7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWindow")
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWindow(QScreen *)

/*
Creates a window as a top level on the targetScreen.

The window is not shown until setVisible(true), show(), or similar is called.

See also setScreen().
*/
func (*QWindow) NewForInheritp() *QWindow {
	return NewQWindowp()
}
func NewQWindowp() *QWindow {
	// arg: 0, QScreen *=Pointer, QScreen=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowC2EP7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWindow")
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:145
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWindow(QWindow *)

/*
Creates a window as a top level on the targetScreen.

The window is not shown until setVisible(true), show(), or similar is called.

See also setScreen().
*/
func (*QWindow) NewForInherit1(parent QWindow_ITF /*777 QWindow **/) *QWindow {
	return NewQWindow1(parent)
}
func NewQWindow1(parent QWindow_ITF /*777 QWindow **/) *QWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWindow")
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:146
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWindow()

/*

 */
func DeleteQWindow(this *QWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qwindow.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSurfaceType(QSurface::SurfaceType)

/*
Sets the surfaceType of the window.

Specifies whether the window is meant for raster rendering with QBackingStore, or OpenGL rendering with QOpenGLContext.

The surfaceType will be used when the native surface is created in the create() function. Calling this function after the native surface has been created requires calling destroy() and create() to release the old native surface and create a new one.

See also surfaceType(), QBackingStore, QOpenGLContext, create(), and destroy().
*/
func (this *QWindow) SetSurfaceType(surfaceType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setSurfaceTypeEN8QSurface11SurfaceTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), surfaceType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType() const

/*
Reimplemented from QSurface::surfaceType().

Returns the surface type of the window.

See also setSurfaceType().
*/
func (this *QWindow) SurfaceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11surfaceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QWindow) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QWindow::Visibility visibility() const

/*

 */
func (this *QWindow) Visibility() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow10visibilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisibility(QWindow::Visibility)

/*

 */
func (this *QWindow) SetVisibility(v int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13setVisibilityENS_10VisibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), v)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void create()

/*
Allocates the platform resources associated with the window.

It is at this point that the surface format set using setFormat() gets resolved into an actual native surface. However, the window remains hidden until setVisible() is called.

Note that it is not usually necessary to call this function directly, as it will be implicitly called by show(), setVisible(), and other functions that require access to the platform resources.

Call destroy() to free the platform resources if necessary.

See also destroy().
*/
func (this *QWindow) Create() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow6createEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] WId winId() const

/*
Returns the window's platform id.

For platforms where this id might be useful, the value returned will uniquely represent the window inside the corresponding screen.

See also screen().
*/
func (this *QWindow) WinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5winIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtGui/qwindow.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * parent(QWindow::AncestorMode) const

/*
Returns the parent window, if any.

If mode is IncludeTransients, then the transient parent is returned if there is no parent.

A window without a parent is known as a top level window.

This function was introduced in  Qt 5.9.

See also setParent().
*/
func (this *QWindow) Parent(mode int) *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6parentENS_12AncestorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:161
// index:1
// Public Visibility=Default Availability=Available
// [8] QWindow * parent() const

/*
Returns the parent window, if any.

If mode is IncludeTransients, then the transient parent is returned if there is no parent.

A window without a parent is known as a top level window.

This function was introduced in  Qt 5.9.

See also setParent().
*/
func (this *QWindow) Parent1() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParent(QWindow *)

/*
Sets the parent Window. This will lead to the windowing system managing the clip of the window, so it will be clipped to the parent window.

Setting parent to be 0 will make the window become a top level window.

If parent is a window created by fromWinId(), then the current window will be embedded inside parent, if the platform supports it.

See also parent().
*/
func (this *QWindow) SetParent(parent QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTopLevel() const

/*
Returns whether the window is top level, i.e. has no parent window.
*/
func (this *QWindow) IsTopLevel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow10isTopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModal() const

/*
Returns whether the window is modal.

A modal window prevents other windows from getting any input.

See also QWindow::modality.
*/
func (this *QWindow) IsModal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow7isModalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowModality modality() const

/*

 */
func (this *QWindow) Modality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8modalityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModality(Qt::WindowModality)

/*

 */
func (this *QWindow) SetModality(modality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setModalityEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QSurfaceFormat &)

/*
Sets the window's surface format.

The format determines properties such as color depth, alpha, depth and stencil buffer size, etc. For example, to give a window a transparent background (provided that the window system supports compositing, and provided that other content in the window does not make it opaque again):


  QSurfaceFormat format;
  format.setAlphaBufferSize(8);
  window.setFormat(format);



The surface format will be resolved in the create() function. Calling this function after create() has been called will not re-resolve the surface format of the native surface.

When the format is not explicitly set via this function, the format returned by QSurfaceFormat::defaultFormat() will be used. This means that when having multiple windows, individual calls to this function can be replaced by one single call to QSurfaceFormat::setDefaultFormat() before creating the first window.

See also format(), create(), destroy(), and QSurfaceFormat::setDefaultFormat().
*/
func (this *QWindow) SetFormat(format QSurfaceFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QSurfaceFormat_PTR() != nil {
		convArg0 = format.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setFormatERK14QSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:171
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format() const

/*
Reimplemented from QSurface::format().

Returns the actual format of this window.

After the window has been created, this function will return the actual surface format of the window. It might differ from the requested format if the requested format could not be fulfilled by the platform. It might also be a superset, for example certain buffer sizes may be larger than requested.

Note: Depending on the platform, certain values in this surface format may still contain the requested values, that is, the values that have been passed to setFormat(). Typical examples are the OpenGL version, profile and options. These may not get updated during create() since these are context specific and a single window may be used together with multiple contexts over its lifetime. Use the QOpenGLContext's format() instead to query such values.

See also setFormat(), create(), requestedFormat(), and QOpenGLContext::format().
*/
func (this *QWindow) Format() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QSurfaceFormat requestedFormat() const

/*
Returns the requested surface format of this window.

If the requested format was not supported by the platform implementation, the requestedFormat will differ from the actual window format.

This is the value set with setFormat().

See also setFormat() and format().
*/
func (this *QWindow) RequestedFormat() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow15requestedFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::WindowFlags)

/*

 */
func (this *QWindow) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8setFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowFlags flags() const

/*

 */
func (this *QWindow) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(Qt::WindowType, bool)

/*
Sets the window flag flag on this window if on is true; otherwise clears the flag.

This function was introduced in  Qt 5.9.

See also setFlags(), flags(), and type().
*/
func (this *QWindow) SetFlag(arg0 int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setFlagEN2Qt10WindowTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(Qt::WindowType, bool)

/*
Sets the window flag flag on this window if on is true; otherwise clears the flag.

This function was introduced in  Qt 5.9.

See also setFlags(), flags(), and type().
*/
func (this *QWindow) SetFlagp(arg0 int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setFlagEN2Qt10WindowTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowType type() const

/*
Returns the type of the window.

This returns the part of the window flags that represents whether the window is a dialog, tooltip, popup, regular window, etc.

See also flags() and setFlags().
*/
func (this *QWindow) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QWindow) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qwindow.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)

/*

 */
func (this *QWindow) SetOpacity(level float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity() const

/*

 */
func (this *QWindow) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QRegion &)

/*
Sets the mask of the window.

The mask is a hint to the windowing system that the application does not want to receive mouse or touch input outside the given region.

The window manager may or may not choose to display any areas of the window not included in the mask, thus it is the application's responsibility to clear to transparent the areas that are not part of the mask.

See also mask().
*/
func (this *QWindow) SetMask(region QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setMaskERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion mask() const

/*
Returns the mask set on the window.

The mask is a hint to the windowing system that the application does not want to receive mouse or touch input outside the given region.

See also setMask().
*/
func (this *QWindow) Mask() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:187
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if the window should appear active from a style perspective.

This is the case for the window that has input focus as well as windows that are in the same parent / transient parent chain as the focus window.

To get the window that currently has focus, use QGuiApplication::focusWindow().

Note: Getter function for property active.
*/
func (this *QWindow) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reportContentOrientationChange(Qt::ScreenOrientation)

/*

 */
func (this *QWindow) ReportContentOrientationChange(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow30reportContentOrientationChangeEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation contentOrientation() const

/*

 */
func (this *QWindow) ContentOrientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow18contentOrientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const

/*
Returns the ratio between physical pixels and device-independent pixels for the window. This value is dependent on the screen the window is on, and may change when the window is moved.

Common values are 1.0 on normal displays and 2.0 on Apple "retina" displays.

Note: For windows not backed by a platform window, meaning that create() was not called, the function will fall back to the associated QScreen's device pixel ratio.

See also QScreen::devicePixelRatio().
*/
func (this *QWindow) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:194
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowState windowState() const

/*
the screen-occupation state of the window

See also setWindowState() and windowStates().
*/
func (this *QWindow) WindowState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11windowStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:195
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowStates windowStates() const

/*
the screen-occupation state of the window

The window can be in a combination of several states. For example, if the window is both minimized and maximized, the window will appear minimized, but clicking on the task bar entry will restore it to the maximized state.

This function was introduced in  Qt 5.10.

See also setWindowStates().
*/
func (this *QWindow) WindowStates() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12windowStatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowState(Qt::WindowState)

/*
set the screen-occupation state of the window

The window state represents whether the window appears in the windowing system as maximized, minimized, fullscreen, or normal.

The enum value Qt::WindowActive is not an accepted parameter.

See also windowState(), showNormal(), showFullScreen(), showMinimized(), showMaximized(), and setWindowStates().
*/
func (this *QWindow) SetWindowState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setWindowStateEN2Qt11WindowStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowStates(Qt::WindowStates)

/*
set the screen-occupation state of the window

The window state represents whether the window appears in the windowing system as maximized, minimized and/or fullscreen.

The window can be in a combination of several states. For example, if the window is both minimized and maximized, the window will appear minimized, but clicking on the task bar entry will restore it to the maximized state.

The enum value Qt::WindowActive should not be set.

This function was introduced in  Qt 5.10.

See also windowStates(), showNormal(), showFullScreen(), showMinimized(), and showMaximized().
*/
func (this *QWindow) SetWindowStates(states int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15setWindowStatesE6QFlagsIN2Qt11WindowStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), states)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransientParent(QWindow *)

/*
Sets the transient parent

This is a hint to the window manager that this window is a dialog or pop-up on behalf of the given window.

In order to cause the window to be centered above its transient parent by default, depending on the window manager, it may also be necessary to call setFlags() with a suitable Qt::WindowType (such as Qt::Dialog).

See also transientParent() and parent().
*/
func (this *QWindow) SetTransientParent(parent QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18setTransientParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * transientParent() const

/*
Returns the transient parent of the window.

See also setTransientParent() and parent().
*/
func (this *QWindow) TransientParent() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow15transientParentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:202
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QWindow *, QWindow::AncestorMode) const

/*
Returns true if the window is an ancestor of the given child. If mode is IncludeTransients, then transient parents are also considered ancestors.
*/
func (this *QWindow) IsAncestorOf(child QWindow_ITF /*777 const QWindow **/, mode int) bool {
	var convArg0 unsafe.Pointer
	if child != nil && child.QWindow_PTR() != nil {
		convArg0 = child.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12isAncestorOfEPKS_NS_12AncestorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:202
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QWindow *, QWindow::AncestorMode) const

/*
Returns true if the window is an ancestor of the given child. If mode is IncludeTransients, then transient parents are also considered ancestors.
*/
func (this *QWindow) IsAncestorOfp(child QWindow_ITF /*777 const QWindow **/) bool {
	var convArg0 unsafe.Pointer
	if child != nil && child.QWindow_PTR() != nil {
		convArg0 = child.QWindow_PTR().GetCthis()
	}
	// arg: 1, QWindow::AncestorMode=Enum, QWindow::AncestorMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12isAncestorOfEPKS_NS_12AncestorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:204
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExposed() const

/*
Returns if this window is exposed in the windowing system.

When the window is not exposed, it is shown by the application but it is still not showing in the windowing system, so the application should minimize rendering and other graphical activities.

An exposeEvent() is sent every time this value changes.

See also exposeEvent().
*/
func (this *QWindow) IsExposed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow9isExposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minimumWidth() const

/*

 */
func (this *QWindow) MinimumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12minimumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minimumHeight() const

/*

 */
func (this *QWindow) MinimumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13minimumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int maximumWidth() const

/*

 */
func (this *QWindow) MaximumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12maximumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int maximumHeight() const

/*

 */
func (this *QWindow) MaximumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13maximumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:211
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Returns the minimum size of the window.

See also setMinimumSize().
*/
func (this *QWindow) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:212
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
Returns the maximum size of the window.

See also setMaximumSize().
*/
func (this *QWindow) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize baseSize() const

/*
Returns the base size of the window.

See also setBaseSize().
*/
func (this *QWindow) BaseSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8baseSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeIncrement() const

/*
Returns the size increment of the window.

See also setSizeIncrement().
*/
func (this *QWindow) SizeIncrement() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13sizeIncrementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSize &)

/*
Sets the minimum size of the window.

This is a hint to the window manager to prevent resizing below the specified size.

See also setMaximumSize() and minimumSize().
*/
func (this *QWindow) SetMinimumSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setMinimumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSize &)

/*
Sets the maximum size of the window.

This is a hint to the window manager to prevent resizing above the specified size.

See also setMinimumSize() and maximumSize().
*/
func (this *QWindow) SetMaximumSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setMaximumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseSize(const QSize &)

/*
Sets the base size of the window.

The base size is used to calculate a proper window size if the window defines sizeIncrement().

See also setMinimumSize(), setMaximumSize(), setSizeIncrement(), and baseSize().
*/
func (this *QWindow) SetBaseSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setBaseSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeIncrement(const QSize &)

/*
Sets the size increment (size) of the window.

When the user resizes the window, the size will move in steps of sizeIncrement().width() pixels horizontally and sizeIncrement().height() pixels vertically, with baseSize() as the basis.

By default, this property contains a size with zero width and height.

The windowing system might not support size increments.

See also sizeIncrement(), setBaseSize(), setMinimumSize(), and setMaximumSize().
*/
func (this *QWindow) SetSizeIncrement(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setSizeIncrementERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:221
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect geometry() const

/*
Returns the geometry of the window, excluding its window frame.

The geometry is in relation to the virtualGeometry() of its screen.

See also setGeometry(), frameMargins(), and frameGeometry().
*/
func (this *QWindow) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:223
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins frameMargins() const

/*
Returns the window frame margins surrounding the window.

See also geometry() and frameGeometry().
*/
func (this *QWindow) FrameMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12frameMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:224
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameGeometry() const

/*
Returns the geometry of the window, including its window frame.

The geometry is in relation to the virtualGeometry() of its screen.

See also geometry() and frameMargins().
*/
func (this *QWindow) FrameGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13frameGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint framePosition() const

/*
Returns the top left position of the window, including its window frame.

This returns the same value as frameGeometry().topLeft().

See also setFramePosition(), geometry(), and frameGeometry().
*/
func (this *QWindow) FramePosition() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13framePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFramePosition(const QPoint &)

/*
Sets the upper left position of the window (point) including its window frame.

The position is in relation to the virtualGeometry() of its screen.

See also framePosition(), setGeometry(), and frameGeometry().
*/
func (this *QWindow) SetFramePosition(point qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setFramePositionERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width() const

/*

 */
func (this *QWindow) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height() const

/*

 */
func (this *QWindow) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x() const

/*

 */
func (this *QWindow) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*

 */
func (this *QWindow) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:234
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] QSize size() const

/*
Reimplemented from QSurface::size().

Returns the size of the window excluding any window frame

See also resize().
*/
func (this *QWindow) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint position() const

/*
Returns the position of the window on the desktop excluding any window frame

See also setPosition().
*/
func (this *QWindow) Position() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPoint &)

/*
set the position of the window on the desktop to pt

The position is in relation to the virtualGeometry() of its screen.

See also position().
*/
func (this *QWindow) SetPosition(pt qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setPositionERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:238
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int, int)

/*
set the position of the window on the desktop to pt

The position is in relation to the virtualGeometry() of its screen.

See also position().
*/
func (this *QWindow) SetPosition1(posx int, posy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setPositionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), posx, posy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSize &)

/*
set the size of the window, excluding any window frame, to newSize

See also size() and geometry().
*/
func (this *QWindow) Resize(newSize qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if newSize != nil && newSize.QSize_PTR() != nil {
		convArg0 = newSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow6resizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:241
// index:1
// Public Visibility=Default Availability=Available
// [-2] void resize(int, int)

/*
set the size of the window, excluding any window frame, to newSize

See also size() and geometry().
*/
func (this *QWindow) Resize1(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow6resizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilePath(const QString &)

/*
set the file name this window is representing.

The windowing system might use filePath to display the path of the document this window is representing in the tile bar.

See also filePath().
*/
func (this *QWindow) SetFilePath(filePath string) {
	var tmpArg0 = qtcore.NewQString5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:244
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath() const

/*
the file name this window is representing.

See also setFilePath().
*/
func (this *QWindow) FilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8filePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qwindow.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*
Sets the window's icon in the windowing system

The window icon might be used by the windowing system for example to decorate the window, and/or in the task switcher.

Note: On macOS, the window title bar icon is meant for windows representing documents, and will only show up if a file path is also set.

See also icon() and setFilePath().
*/
func (this *QWindow) SetIcon(icon QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:247
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*
Returns the window's icon in the windowing system

See also setIcon().
*/
func (this *QWindow) Icon() *QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroy()

/*
Releases the native platform resources associated with this window.

See also create().
*/
func (this *QWindow) Destroy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7destroyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:253
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setKeyboardGrabEnabled(bool)

/*
Sets whether keyboard grab should be enabled or not (grab).

If the return value is true, the window receives all key events until setKeyboardGrabEnabled(false) is called; other windows get no key events at all. Mouse events are not affected. Use setMouseGrabEnabled() if you want to grab that.

See also setMouseGrabEnabled().
*/
func (this *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow22setKeyboardGrabEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), grab)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setMouseGrabEnabled(bool)

/*
Sets whether mouse grab should be enabled or not (grab).

If the return value is true, the window receives all mouse events until setMouseGrabEnabled(false) is called; other windows get no mouse events at all. Keyboard events are not affected. Use setKeyboardGrabEnabled() if you want to grab that.

See also setKeyboardGrabEnabled().
*/
func (this *QWindow) SetMouseGrabEnabled(grab bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow19setMouseGrabEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), grab)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:256
// index:0
// Public Visibility=Default Availability=Available
// [8] QScreen * screen() const

/*
Returns the screen on which the window is shown, or null if there is none.

For child windows, this returns the screen of the corresponding top level window.

See also setScreen() and QScreen::virtualSiblings().
*/
func (this *QWindow) Screen() *QScreen /*777 QScreen **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6screenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreen(QScreen *)

/*
Sets the screen on which the window should be shown.

If the window has been created, it will be recreated on the newScreen.

Note: If the screen is part of a virtual desktop of multiple screens, the window will not move automatically to newScreen. To place the window relative to the screen, use the screen's topLeft() position.

This function only works for top level windows.

See also screen() and QScreen::virtualSiblings().
*/
func (this *QWindow) SetScreen(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setScreenEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:259
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleRoot() const

/*

 */
func (this *QWindow) AccessibleRoot() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow14accessibleRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:260
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * focusObject() const

/*
Returns the QObject that will be the final receiver of events tied focus, such as key events.
*/
func (this *QWindow) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11focusObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:262
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapToGlobal(const QPoint &) const

/*
Translates the window coordinate pos to global screen coordinates. For example, mapToGlobal(QPoint(0,0)) would give the global coordinates of the top-left pixel of the window.

See also mapFromGlobal().
*/
func (this *QWindow) MapToGlobal(pos qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11mapToGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromGlobal(const QPoint &) const

/*
Translates the global screen coordinate pos to window coordinates.

See also mapToGlobal().
*/
func (this *QWindow) MapFromGlobal(pos qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13mapFromGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:266
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor() const

/*
the cursor shape for this window

See also setCursor() and unsetCursor().
*/
func (this *QWindow) Cursor() *QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)

/*
set the cursor shape for this window

The mouse cursor will assume this shape when it is over this window, unless an override cursor is set. See the list of predefined cursor objects for a range of useful shapes.

If no cursor has been set, or after a call to unsetCursor(), the parent window's cursor is used.

By default, the cursor has the Qt::ArrowCursor shape.

Some underlying window implementations will reset the cursor if it leaves a window even if the mouse is grabbed. If you want to have a cursor set for all windows, even when outside the window, consider QGuiApplication::setOverrideCursor().

See also cursor() and QGuiApplication::setOverrideCursor().
*/
func (this *QWindow) SetCursor(arg0 QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCursor_PTR() != nil {
		convArg0 = arg0.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()

/*
Restores the default arrow cursor for this window.
*/
func (this *QWindow) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:271
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * fromWinId(WId)

/*
Creates a local representation of a window created by another process or by using native libraries below Qt.

Given the handle id to a native window, this method creates a QWindow object which can be used to represent the window when invoking methods like setParent() and setTransientParent().

This can be used, on platforms which support it, to embed a QWindow inside a native window, or to embed a native window inside a QWindow.

If foreign windows are not supported or embedding the native window failed in the platform plugin, this function returns 0.

Note: The resulting QWindow should not be used to manipulate the underlying native window (besides re-parenting), or to observe state changes of the native window. Any support for these kind of operations is incidental, highly platform dependent and untested.

See also setParent() and setTransientParent().
*/
func (this *QWindow) FromWinId(id uint64) *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9fromWinIdEy", qtrt.FFI_TYPE_POINTER, id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWindow_FromWinId(id uint64) *QWindow /*777 QWindow **/ {
	var nilthis *QWindow
	rv := nilthis.FromWinId(id)
	return rv
}

// /usr/include/qt/QtGui/qwindow.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestActivate()

/*
Requests the window to be activated, i.e. receive keyboard focus.

See also isActive(), QGuiApplication::focusWindow(), and QWindowsWindowFunctions::setWindowActivationBehavior().
*/
func (this *QWindow) RequestActivate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15requestActivateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QWindow) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()

/*
Shows the window.

This is equivalent to calling showFullScreen(), showMaximized(), or showNormal(), depending on the platform's default behavior for the window type and flags.

See also showFullScreen(), showMaximized(), showNormal(), hide(), QStyleHints::showIsFullScreen(), and flags().
*/
func (this *QWindow) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()

/*
Hides the window.

Equivalent to calling setVisible(false).

See also show() and setVisible().
*/
func (this *QWindow) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:286
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMinimized()

/*
Shows the window as minimized.

Equivalent to calling setWindowStates(Qt::WindowMinimized) and then setVisible(true).

See also setWindowStates() and setVisible().
*/
func (this *QWindow) ShowMinimized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13showMinimizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMaximized()

/*
Shows the window as maximized.

Equivalent to calling setWindowStates(Qt::WindowMaximized) and then setVisible(true).

See also setWindowStates() and setVisible().
*/
func (this *QWindow) ShowMaximized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13showMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showFullScreen()

/*
Shows the window as fullscreen.

Equivalent to calling setWindowStates(Qt::WindowFullScreen) and then setVisible(true).

See also setWindowStates() and setVisible().
*/
func (this *QWindow) ShowFullScreen() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14showFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNormal()

/*
Shows the window as normal, i.e. neither maximized, minimized, nor fullscreen.

Equivalent to calling setWindowStates(Qt::WindowNoState) and then setVisible(true).

See also setWindowStates() and setVisible().
*/
func (this *QWindow) ShowNormal() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10showNormalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:291
// index:0
// Public Visibility=Default Availability=Available
// [1] bool close()

/*
Close the window.

This closes the window, effectively calling destroy(), and potentially quitting the application. Returns true on success, false if it has a parent window (in which case the top level window should be closed instead).

See also destroy() and QGuiApplication::quitOnLastWindowClosed().
*/
func (this *QWindow) Close() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void raise()

/*
Raise the window in the windowing system.

Requests that the window be raised to appear above other windows.
*/
func (this *QWindow) Raise() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5raiseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lower()

/*
Lower the window in the windowing system.

Requests that the window be lowered to appear below other windows.
*/
func (this *QWindow) Lower() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5lowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*

 */
func (this *QWindow) SetTitle(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(int)

/*

 */
func (this *QWindow) SetX(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4setXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(int)

/*

 */
func (this *QWindow) SetY(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4setYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(int)

/*

 */
func (this *QWindow) SetWidth(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeight(int)

/*

 */
func (this *QWindow) SetHeight(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(int, int, int, int)

/*
Sets the geometry of the window, excluding its window frame, to a rectangle constructed from posx, posy, w and h.

The geometry is in relation to the virtualGeometry() of its screen.

See also geometry().
*/
func (this *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setGeometryEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), posx, posy, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:302
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Sets the geometry of the window, excluding its window frame, to a rectangle constructed from posx, posy, w and h.

The geometry is in relation to the virtualGeometry() of its screen.

See also geometry().
*/
func (this *QWindow) SetGeometry1(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(int)

/*

 */
func (this *QWindow) SetMinimumWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15setMinimumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(int)

/*

 */
func (this *QWindow) SetMinimumHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setMinimumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:306
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(int)

/*

 */
func (this *QWindow) SetMaximumWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15setMaximumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(int)

/*

 */
func (this *QWindow) SetMaximumHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setMaximumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:309
// index:0
// Public Visibility=Default Availability=Available
// [-2] void alert(int)

/*
Causes an alert to be shown for msec miliseconds. If msec is 0 (the default), then the alert is shown indefinitely until the window becomes active again. This function has no effect on an active window.

In alert state, the window indicates that it demands attention, for example by flashing or bouncing the taskbar entry.

This function was introduced in  Qt 5.1.
*/
func (this *QWindow) Alert(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5alertEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestUpdate()

/*
Schedules a QEvent::UpdateRequest event to be delivered to this window.

The event is delivered in sync with the display vsync on platforms where this is possible. Otherwise, the event is delivered after a delay of 5 ms. The additional time is there to give the event loop a bit of idle time to gather system events, and can be overridden using the QT_QPA_UPDATE_IDLE_TIME environment variable.

When driving animations, this function should be called once after drawing has completed. Calling this function multiple times will result in a single event being delivered to the window.

Subclasses of QWindow should reimplement event(), intercept the event and call the application's rendering code, then call the base class implementation.

Note: The subclass' reimplementation of event() must invoke the base class implementation, unless it is absolutely sure that the event does not need to be handled by the base class. For example, the default implementation of this function relies on QEvent::Timer events. Filtering them away would therefore break the delivery of the update events.

This function was introduced in  Qt 5.5.
*/
func (this *QWindow) RequestUpdate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13requestUpdateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenChanged(QScreen *)

/*
This signal is emitted when a window's screen changes, either by being set explicitly with setScreen(), or automatically when the window's screen is removed.
*/
func (this *QWindow) ScreenChanged(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13screenChangedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modalityChanged(Qt::WindowModality)

/*
This signal is emitted when the Qwindow::modality property changes to modality.

Note: Notifier signal for property modality.
*/
func (this *QWindow) ModalityChanged(modality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15modalityChangedEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowStateChanged(Qt::WindowState)

/*
This signal is emitted when the windowState changes, either by being set explicitly with setWindowStates(), or automatically when the user clicks one of the titlebar buttons or by other means.
*/
func (this *QWindow) WindowStateChanged(windowState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18windowStateChangedEN2Qt11WindowStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), windowState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:317
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowTitleChanged(const QString &)

/*

 */
func (this *QWindow) WindowTitleChanged(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18windowTitleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void xChanged(int)

/*

 */
func (this *QWindow) XChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8xChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void yChanged(int)

/*

 */
func (this *QWindow) YChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8yChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widthChanged(int)

/*

 */
func (this *QWindow) WidthChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow12widthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:323
// index:0
// Public Visibility=Default Availability=Available
// [-2] void heightChanged(int)

/*

 */
func (this *QWindow) HeightChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13heightChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:325
// index:0
// Public Visibility=Default Availability=Available
// [-2] void minimumWidthChanged(int)

/*

 */
func (this *QWindow) MinimumWidthChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow19minimumWidthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:326
// index:0
// Public Visibility=Default Availability=Available
// [-2] void minimumHeightChanged(int)

/*

 */
func (this *QWindow) MinimumHeightChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow20minimumHeightChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:327
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumWidthChanged(int)

/*

 */
func (this *QWindow) MaximumWidthChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow19maximumWidthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumHeightChanged(int)

/*

 */
func (this *QWindow) MaximumHeightChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow20maximumHeightChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChanged(bool)

/*

 */
func (this *QWindow) VisibleChanged(arg bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14visibleChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:331
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibilityChanged(QWindow::Visibility)

/*

 */
func (this *QWindow) VisibilityChanged(visibility int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow17visibilityChangedENS_10VisibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visibility)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeChanged()

/*

 */
func (this *QWindow) ActiveChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13activeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:333
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentOrientationChanged(Qt::ScreenOrientation)

/*

 */
func (this *QWindow) ContentOrientationChanged(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow25contentOrientationChangedEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:335
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusObjectChanged(QObject *)

/*
This signal is emitted when the final receiver of events tied to focus is changed to object.

See also focusObject().
*/
func (this *QWindow) FocusObjectChanged(object qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18focusObjectChangedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityChanged(qreal)

/*

 */
func (this *QWindow) OpacityChanged(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14opacityChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:340
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void exposeEvent(QExposeEvent *)

/*
The expose event (ev) is sent by the window system whenever the window's exposure on screen changes.

The application can start rendering into the window with QBackingStore and QOpenGLContext as soon as it gets an exposeEvent() such that isExposed() is true.

If the window is moved off screen, is made totally obscured by another window, iconified or similar, this function might be called and the value of isExposed() might change to false. When this happens, an application should stop its rendering as it is no longer visible to the user.

A resize event will always be sent before the expose event the first time a window is shown.

See also isExposed().
*/
func (this *QWindow) ExposeEvent(arg0 QExposeEvent_ITF /*777 QExposeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QExposeEvent_PTR() != nil {
		convArg0 = arg0.QExposeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11exposeEventEP12QExposeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:341
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Override this to handle resize events (ev).

The resize event is called whenever the window is resized in the windowing system, either directly through the windowing system acknowledging a setGeometry() or resize() request, or indirectly through the user resizing the window manually.
*/
func (this *QWindow) ResizeEvent(arg0 QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:342
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)

/*
Override this to handle window move events (ev).
*/
func (this *QWindow) MoveEvent(arg0 QMoveEvent_ITF /*777 QMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMoveEvent_PTR() != nil {
		convArg0 = arg0.QMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:343
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Override this to handle focus in events (ev).

Focus in events are sent when the window receives keyboard focus.

See also focusOutEvent().
*/
func (this *QWindow) FocusInEvent(arg0 QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:344
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Override this to handle focus out events (ev).

Focus out events are sent when the window loses keyboard focus.

See also focusInEvent().
*/
func (this *QWindow) FocusOutEvent(arg0 QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:346
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Override this to handle show events (ev).

The function is called when the window has requested becoming visible.

If the window is successfully shown by the windowing system, this will be followed by a resize and an expose event.
*/
func (this *QWindow) ShowEvent(arg0 QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:347
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Override this to handle hide events (ev).

The function is called when the window has requested being hidden in the windowing system.
*/
func (this *QWindow) HideEvent(arg0 QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:350
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().

Override this to handle any event (ev) sent to the window. Return true if the event was recognized and processed.

Remember to call the base class version if you wish for mouse events, key events, resize events, etc to be dispatched as usual.
*/
func (this *QWindow) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:351
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Override this to handle key press events (ev).

See also keyReleaseEvent().
*/
func (this *QWindow) KeyPressEvent(arg0 QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:352
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Override this to handle key release events (ev).

See also keyPressEvent().
*/
func (this *QWindow) KeyReleaseEvent(arg0 QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:353
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Override this to handle mouse press events (ev).

See also mouseReleaseEvent().
*/
func (this *QWindow) MousePressEvent(arg0 QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:354
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Override this to handle mouse release events (ev).

See also mousePressEvent().
*/
func (this *QWindow) MouseReleaseEvent(arg0 QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:355
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Override this to handle mouse double click events (ev).

See also mousePressEvent() and QStyleHints::mouseDoubleClickInterval().
*/
func (this *QWindow) MouseDoubleClickEvent(arg0 QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:356
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Override this to handle mouse move events (ev).
*/
func (this *QWindow) MouseMoveEvent(arg0 QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:358
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Override this to handle mouse wheel or other wheel events (ev).
*/
func (this *QWindow) WheelEvent(arg0 QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:360
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void touchEvent(QTouchEvent *)

/*
Override this to handle touch events (ev).
*/
func (this *QWindow) TouchEvent(arg0 QTouchEvent_ITF /*777 QTouchEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTouchEvent_PTR() != nil {
		convArg0 = arg0.QTouchEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10touchEventEP11QTouchEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:362
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabletEvent(QTabletEvent *)

/*
Override this to handle tablet press, move, and release events (ev).

Proximity enter and leave events are not sent to windows, they are delivered to the application instance.
*/
func (this *QWindow) TabletEvent(arg0 QTabletEvent_ITF /*777 QTabletEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTabletEvent_PTR() != nil {
		convArg0 = arg0.QTabletEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11tabletEventEP12QTabletEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:364
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool nativeEvent(const QByteArray &, void *, long *)

/*
Override this to handle platform dependent events. Will be given eventType, message and result.

This might make your application non-portable.

Should return true only if the event was handled.
*/
func (this *QWindow) NativeEvent(eventType qtcore.QByteArray_ITF, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 unsafe.Pointer
	if eventType != nil && eventType.QByteArray_PTR() != nil {
		convArg0 = eventType.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11nativeEventERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes what part of the screen the window occupies or should occupy.



This enum was introduced or modified in  Qt 5.1.

*/
type QWindow__Visibility = int

// The window is not visible in any way, however it may remember a latent visibility which can be restored by setting AutomaticVisibility.
const QWindow__Hidden QWindow__Visibility = 0

// This means to give the window a default visible state, which might be fullscreen or windowed depending on the platform. It can be given as a parameter to setVisibility but will never be read back from the visibility accessor.
const QWindow__AutomaticVisibility QWindow__Visibility = 1

// The window occupies part of the screen, but not necessarily the entire screen. This state will occur only on windowing systems which support showing multiple windows simultaneously. In this state it is possible for the user to move and resize the window manually, if WindowFlags permit it and if it is supported by the windowing system.
const QWindow__Windowed QWindow__Visibility = 2

// The window is reduced to an entry or icon on the task bar, dock, task list or desktop, depending on how the windowing system handles minimized windows.
const QWindow__Minimized QWindow__Visibility = 3

// The window occupies one entire screen, and the titlebar is still visible. On most windowing systems this is the state achieved by clicking the maximize button on the toolbar.
const QWindow__Maximized QWindow__Visibility = 4

// The window occupies one entire screen, is not resizable, and there is no titlebar. On some platforms which do not support showing multiple simultaneous windows, this can be the usual visibility when the window is not hidden.
const QWindow__FullScreen QWindow__Visibility = 5

func (this *QWindow) VisibilityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWindow_VisibilityItemName(val int) string {
	var nilthis *QWindow
	return nilthis.VisibilityItemName(val)
}

/*
This enum is used to control whether or not transient parents should be considered ancestors.


*/
type QWindow__AncestorMode = int

// Transient parents are not considered ancestors.
const QWindow__ExcludeTransients QWindow__AncestorMode = 0

// Transient parents are considered ancestors.
const QWindow__IncludeTransients QWindow__AncestorMode = 1

func (this *QWindow) AncestorModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWindow_AncestorModeItemName(val int) string {
	var nilthis *QWindow
	return nilthis.AncestorModeItemName(val)
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
