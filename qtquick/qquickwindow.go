package qtquick

// /usr/include/qt/QtQuick/qquickwindow.h
// #include <qquickwindow.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// void exposeEvent(QExposeEvent *)
func (this *QQuickWindow) InheritExposeEvent(f func(arg0 *qtgui.QExposeEvent /*777 QExposeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "exposeEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QQuickWindow) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QQuickWindow) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QQuickWindow) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QQuickWindow) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QQuickWindow) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool event(QEvent *)
func (this *QQuickWindow) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QQuickWindow) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QQuickWindow) InheritKeyReleaseEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QQuickWindow) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickWindow) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickWindow) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QQuickWindow) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QQuickWindow) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

/*

 */
type QQuickWindow struct {
	*qtgui.QWindow
}
type QQuickWindow_ITF interface {
	qtgui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow { return ptr }

func (this *QQuickWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWindow.GetCthis()
	}
}
func (this *QQuickWindow) SetCthis(cthis unsafe.Pointer) {
	this.QWindow = qtgui.NewQWindowFromPointer(cthis)
}
func NewQQuickWindowFromPointer(cthis unsafe.Pointer) *QQuickWindow {
	bcthis0 := qtgui.NewQWindowFromPointer(cthis)
	return &QQuickWindow{bcthis0}
}
func (*QQuickWindow) NewFromPointer(cthis unsafe.Pointer) *QQuickWindow {
	return NewQQuickWindowFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickwindow.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QWindow *)

/*
Constructs a window for displaying a QML scene with parent window parent.
*/
func (*QQuickWindow) NewForInherit(parent qtgui.QWindow_ITF /*777 QWindow **/) *QQuickWindow {
	return NewQQuickWindow(parent)
}
func NewQQuickWindow(parent qtgui.QWindow_ITF /*777 QWindow **/) *QQuickWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWindow")
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QWindow *)

/*
Constructs a window for displaying a QML scene with parent window parent.
*/
func (*QQuickWindow) NewForInherit__() *QQuickWindow {
	return NewQQuickWindow__()
}
func NewQQuickWindow__() *QQuickWindow {
	// arg: 0, QWindow *=Pointer, QWindow=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWindow")
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:110
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QQuickRenderControl *)

/*
Constructs a window for displaying a QML scene with parent window parent.
*/
func (*QQuickWindow) NewForInherit_1(renderControl QQuickRenderControl_ITF /*777 QQuickRenderControl **/) *QQuickWindow {
	return NewQQuickWindow_1(renderControl)
}
func NewQQuickWindow_1(renderControl QQuickRenderControl_ITF /*777 QQuickRenderControl **/) *QQuickWindow {
	var convArg0 unsafe.Pointer
	if renderControl != nil && renderControl.QQuickRenderControl_PTR() != nil {
		convArg0 = renderControl.QQuickRenderControl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowC2EP19QQuickRenderControl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWindow")
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickWindow()

/*

 */
func DeleteQQuickWindow(this *QQuickWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickwindow.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * contentItem() const

/*

 */
func (this *QQuickWindow) ContentItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow11contentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * activeFocusItem() const

/*

 */
func (this *QQuickWindow) ActiveFocusItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow15activeFocusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * focusObject() const

/*

 */
func (this *QQuickWindow) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow11focusObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * mouseGrabberItem() const

/*
Returns the item which currently has the mouse grab.
*/
func (this *QQuickWindow) MouseGrabberItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow16mouseGrabberItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sendEvent(QQuickItem *, QEvent *)

/*

 */
func (this *QQuickWindow) SendEvent(arg0 QQuickItem_ITF /*777 QQuickItem **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow9sendEventEP10QQuickItemP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage grabWindow()

/*
Grabs the contents of the window and returns it as an image.

It is possible to call the grabWindow() function when the window is not visible. This requires that the window is created and has a valid size and that no other QQuickWindow instances are rendering in the same process.

Warning: Calling this function will cause performance problems.

Warning: This function can only be called from the GUI thread.
*/
func (this *QQuickWindow) GrabWindow() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow10grabWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderTarget(uint, const QSize &)

/*
Sets the render target for this window to be fbo.

The specified fbo must be created in the context of the window or one that shares with it.

Note: This function only has an effect when using the default OpenGL scene graph adaptation.

Warning: This function can only be called from the thread doing the rendering.

See also renderTarget().
*/
func (this *QQuickWindow) SetRenderTarget(fboId uint, size qtcore.QSize_ITF) {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15setRenderTargetEjRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fboId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] uint renderTargetId() const

/*
Returns the FBO id of the render target when set; otherwise returns 0.
*/
func (this *QQuickWindow) RenderTargetId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow14renderTargetIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtQuick/qquickwindow.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize renderTargetSize() const

/*
Returns the size of the currently set render target; otherwise returns an empty size.
*/
func (this *QQuickWindow) RenderTargetSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow16renderTargetSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetOpenGLState()

/*
Call this function to reset the OpenGL context its default state.

The scene graph uses the OpenGL context and will both rely on and clobber its state. When mixing raw OpenGL commands with scene graph rendering, this function provides a convenient way of resetting the OpenGL context state back to its default values.

This function does not touch state in the fixed-function pipeline.

This function does not clear the color, depth and stencil buffers. Use QQuickWindow::setClearBeforeRendering to control clearing of the color buffer. The depth and stencil buffer might be clobbered by the scene graph renderer. Clear these manually on demand.

Note: This function only has an effect when using the default OpenGL scene graph adaptation.

This function was introduced in  Qt 5.2.

See also QQuickWindow::beforeRendering().
*/
func (this *QQuickWindow) ResetOpenGLState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow16resetOpenGLStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlIncubationController * incubationController() const

/*
Returns an incubation controller that splices incubation between frames for this window. QQuickView automatically installs this controller for you, otherwise you will need to install it yourself using QQmlEngine::setIncubationController().

The controller is owned by the window and will be destroyed when the window is deleted.
*/
func (this *QQuickWindow) IncubationController() *qtqml.QQmlIncubationController /*777 QQmlIncubationController **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow20incubationControllerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtqml.NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleRoot() const

/*
Returns an accessibility interface for this window, or 0 if such an interface cannot be created.
*/
func (this *QQuickWindow) AccessibleRoot() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow14accessibleRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromImage(const QImage &) const

/*
Creates a new QSGTexture from the supplied image. If the image has an alpha channel, the corresponding texture will have an alpha channel.

The caller of the function is responsible for deleting the returned texture. For example whe using the OpenGL adaptation the actual OpenGL texture will be deleted when the texture object is deleted.

When options contains TextureCanUseAtlas, the engine may put the image into a texture atlas. Textures in an atlas need to rely on QSGTexture::normalizedTextureSubRect() for their geometry and will not support QSGTexture::Repeat. Other values from CreateTextureOption are ignored.

When options contains TextureIsOpaque, the engine will create an RGB texture which returns false for QSGTexture::hasAlphaChannel(). Opaque textures will in most cases be faster to render. When this flag is not set, the texture will have an alpha channel based on the image's format.

When options contains TextureHasMipmaps, the engine will create a texture which can use mipmap filtering. Mipmapped textures can not be in an atlas.

When using the OpenGL adaptation, the returned texture will be using GL_TEXTURE_2D as texture target and GL_RGBA as internal format. Reimplement QSGTexture to create textures with different parameters.

Warning: This function will return 0 if the scene graph has not yet been initialized.

Warning: The returned texture is not memory managed by the scene graph and must be explicitly deleted by the caller on the rendering thread. This is achieved by deleting the texture from a QSGNode destructor or by using deleteLater() in the case where the texture already has affinity to the rendering thread.

This function can be called from any thread.

See also sceneGraphInitialized() and QSGTexture.
*/
func (this *QQuickWindow) CreateTextureFromImage(image qtgui.QImage_ITF) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow22createTextureFromImageERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:144
// index:1
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromImage(const QImage &, QQuickWindow::CreateTextureOptions) const

/*
Creates a new QSGTexture from the supplied image. If the image has an alpha channel, the corresponding texture will have an alpha channel.

The caller of the function is responsible for deleting the returned texture. For example whe using the OpenGL adaptation the actual OpenGL texture will be deleted when the texture object is deleted.

When options contains TextureCanUseAtlas, the engine may put the image into a texture atlas. Textures in an atlas need to rely on QSGTexture::normalizedTextureSubRect() for their geometry and will not support QSGTexture::Repeat. Other values from CreateTextureOption are ignored.

When options contains TextureIsOpaque, the engine will create an RGB texture which returns false for QSGTexture::hasAlphaChannel(). Opaque textures will in most cases be faster to render. When this flag is not set, the texture will have an alpha channel based on the image's format.

When options contains TextureHasMipmaps, the engine will create a texture which can use mipmap filtering. Mipmapped textures can not be in an atlas.

When using the OpenGL adaptation, the returned texture will be using GL_TEXTURE_2D as texture target and GL_RGBA as internal format. Reimplement QSGTexture to create textures with different parameters.

Warning: This function will return 0 if the scene graph has not yet been initialized.

Warning: The returned texture is not memory managed by the scene graph and must be explicitly deleted by the caller on the rendering thread. This is achieved by deleting the texture from a QSGNode destructor or by using deleteLater() in the case where the texture already has affinity to the rendering thread.

This function can be called from any thread.

See also sceneGraphInitialized() and QSGTexture.
*/
func (this *QQuickWindow) CreateTextureFromImage_1(image qtgui.QImage_ITF, options int) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow22createTextureFromImageERK6QImage6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromId(uint, const QSize &, QQuickWindow::CreateTextureOptions) const

/*
Creates a new QSGTexture object from an existing OpenGL texture id and size.

The caller of the function is responsible for deleting the returned texture.

The returned texture will be using GL_TEXTURE_2D as texture target and assumes that internal format is GL_RGBA. Reimplement QSGTexture to create textures with different parameters.

Use options to customize the texture attributes. The TextureUsesAtlas option is ignored.

Warning: This function will return null if the scenegraph has not yet been initialized or OpenGL is not in use.

Note: This function only has an effect when using the default OpenGL scene graph adaptation.

See also sceneGraphInitialized() and QSGTexture.
*/
func (this *QQuickWindow) CreateTextureFromId(id uint, size qtcore.QSize_ITF, options int) *QSGTexture /*777 QSGTexture **/ {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createTextureFromIdEjRK5QSize6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromId(uint, const QSize &, QQuickWindow::CreateTextureOptions) const

/*
Creates a new QSGTexture object from an existing OpenGL texture id and size.

The caller of the function is responsible for deleting the returned texture.

The returned texture will be using GL_TEXTURE_2D as texture target and assumes that internal format is GL_RGBA. Reimplement QSGTexture to create textures with different parameters.

Use options to customize the texture attributes. The TextureUsesAtlas option is ignored.

Warning: This function will return null if the scenegraph has not yet been initialized or OpenGL is not in use.

Note: This function only has an effect when using the default OpenGL scene graph adaptation.

See also sceneGraphInitialized() and QSGTexture.
*/
func (this *QQuickWindow) CreateTextureFromId__(id uint, size qtcore.QSize_ITF) *QSGTexture /*777 QSGTexture **/ {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QQuickWindow::CreateTextureOptions=Typedef, QQuickWindow::CreateTextureOptions=Typedef, QFlags<QQuickWindow::CreateTextureOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createTextureFromIdEjRK5QSize6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearBeforeRendering(bool)

/*
Sets whether the scene graph rendering of QML should clear the color buffer before it starts rendering to enabled.

By disabling clearing of the color buffer, it is possible to render OpengGL content under the scene graph.

The color buffer is cleared by default.

See also clearBeforeRendering() and beforeRendering().
*/
func (this *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow23setClearBeforeRenderingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clearBeforeRendering() const

/*
Returns whether clearing of the color buffer is done before rendering or not.

See also setClearBeforeRendering().
*/
func (this *QQuickWindow) ClearBeforeRendering() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow20clearBeforeRenderingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)

/*

 */
func (this *QQuickWindow) SetColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:151
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color() const

/*

 */
func (this *QQuickWindow) Color() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasDefaultAlphaBuffer()

/*
Returns whether to use alpha transparency on newly created windows.

This function was introduced in  Qt 5.1.

See also setDefaultAlphaBuffer().
*/
func (this *QQuickWindow) HasDefaultAlphaBuffer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21hasDefaultAlphaBufferEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQuickWindow_HasDefaultAlphaBuffer() bool {
	var nilthis *QQuickWindow
	rv := nilthis.HasDefaultAlphaBuffer()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultAlphaBuffer(bool)

/*
useAlpha specifies whether to use alpha transparency on newly created windows.

In any application which expects to create translucent windows, it's necessary to set this to true before creating the first QQuickWindow. The default value is false.

This function was introduced in  Qt 5.1.

See also hasDefaultAlphaBuffer().
*/
func (this *QQuickWindow) SetDefaultAlphaBuffer(useAlpha bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21setDefaultAlphaBufferEb", qtrt.FFI_TYPE_POINTER, useAlpha)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	var nilthis *QQuickWindow
	nilthis.SetDefaultAlphaBuffer(useAlpha)
}

// /usr/include/qt/QtQuick/qquickwindow.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentOpenGLContext(bool)

/*
Sets whether the OpenGL context should be preserved, and cannot be released until the last window is deleted, to persistent. The default value is true.

The OpenGL context can be released to free up graphics resources when the window is obscured, hidden or not rendering. When this happens is implementation specific.

The QOpenGLContext::aboutToBeDestroyed() signal is emitted from the QQuickWindow::openglContext() when the OpenGL context is about to be released. The QQuickWindow::sceneGraphInitialized() signal is emitted when a new OpenGL context is created for this window. Make a Qt::DirectConnection to these signals to be notified.

The OpenGL context is still released when the last QQuickWindow is deleted.

Note: This only has an effect when using the default OpenGL scene graph adaptation.

See also isPersistentOpenGLContext(), setPersistentSceneGraph(), QOpenGLContext::aboutToBeDestroyed(), and sceneGraphInitialized().
*/
func (this *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow26setPersistentOpenGLContextEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), persistent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:157
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentOpenGLContext() const

/*
Returns whether the OpenGL context can be released during the lifetime of the QQuickWindow.

Note: This is a hint. When and how this happens is implementation specific. It also only has an effect when using the default OpenGL scene graph adaptation
*/
func (this *QQuickWindow) IsPersistentOpenGLContext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow25isPersistentOpenGLContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentSceneGraph(bool)

/*
Sets whether the scene graph nodes and resources can be released to persistent. The default value is true.

The scene graph nodes and resources can be released to free up graphics resources when the window is obscured, hidden or not rendering. When this happens is implementation specific.

The QQuickWindow::sceneGraphInvalidated() signal is emitted when cleanup occurs. The QQuickWindow::sceneGraphInitialized() signal is emitted when a new scene graph is recreated for this window. Make a Qt::DirectConnection to these signals to be notified.

The scene graph nodes and resources are still released when the last QQuickWindow is deleted.

See also isPersistentSceneGraph(), setPersistentOpenGLContext(), sceneGraphInvalidated(), and sceneGraphInitialized().
*/
func (this *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow23setPersistentSceneGraphEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), persistent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentSceneGraph() const

/*
Returns whether the scene graph nodes and resources can be released during the lifetime of this QQuickWindow.

Note: This is a hint. When and how this happens is implementation specific.
*/
func (this *QQuickWindow) IsPersistentSceneGraph() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow22isPersistentSceneGraphEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSceneGraphInitialized() const

/*
Returns true if the scene graph has been initialized; otherwise returns false.
*/
func (this *QQuickWindow) IsSceneGraphInitialized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow23isSceneGraphInitializedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scheduleRenderJob(QRunnable *, QQuickWindow::RenderStage)

/*
Schedules job to run when the rendering of this window reaches the given stage.

This is a convenience to the equivalent signals in QQuickWindow for "one shot" tasks.

The window takes ownership over job and will delete it when the job is completed.

If rendering is shut down before job has a chance to run, the job will be run and then deleted as part of the scene graph cleanup. If the window is never shown and no rendering happens before the QQuickWindow is destroyed, all pending jobs will be destroyed without their run() method being called.

If the rendering is happening on a different thread, then the job will happen on the rendering thread.

If stage is NoStage, job will be run at the earliest opportunity whenever the render thread is not busy rendering a frame. If there is no OpenGL context available or the window is not exposed at the time the job is either posted or handled, it is deleted without executing the run() method. If a non-threaded renderer is in use, the run() method of the job is executed synchronously. The OpenGL context is changed to the renderer context before executing a NoStage job.

Note: This function does not trigger rendering; the jobs targeting any other stage than NoStage will be stored run until rendering is triggered elsewhere. To force the job to run earlier, call QQuickWindow::update();

This function was introduced in  Qt 5.4.

See also beforeRendering(), afterRendering(), beforeSynchronizing(), afterSynchronizing(), frameSwapped(), and sceneGraphInvalidated().
*/
func (this *QQuickWindow) ScheduleRenderJob(job qtcore.QRunnable_ITF /*777 QRunnable **/, schedule int) {
	var convArg0 unsafe.Pointer
	if job != nil && job.QRunnable_PTR() != nil {
		convArg0 = job.QRunnable_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17scheduleRenderJobEP9QRunnableNS_11RenderStageE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, schedule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal effectiveDevicePixelRatio() const

/*
Returns the device pixel ratio for this window.

This is different from QWindow::devicePixelRatio() in that it supports redirected rendering via QQuickRenderControl. When using a QQuickRenderControl, the QQuickWindow is often not created, meaning it is never shown and there is no underlying native window created in the windowing system. As a result, querying properties like the device pixel ratio cannot give correct results. Use this function instead.

See also QWindow::devicePixelRatio().
*/
func (this *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow25effectiveDevicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickwindow.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRendererInterface * rendererInterface() const

/*
Returns the current renderer interface. The value is always valid and is never null.

Note: This function can be called at any time after constructing the QQuickWindow, even while isSceneGraphInitialized() is still false. However, some renderer interface functions, in particular QSGRendererInterface::getResource() will not be functional until the scenegraph is up and running. Backend queries, like QSGRendererInterface::graphicsApi() or QSGRendererInterface::shaderType(), will always be functional on the other hand.

Note: The ownership of the returned pointer stays with Qt. The returned instance may or may not be shared between different QQuickWindow instances, depending on the scenegraph backend in use. Therefore applications are expected to query the interface object for each QQuickWindow instead of reusing the already queried pointer.

This function was introduced in  Qt 5.8.

See also QSGRenderNode and QSGRendererInterface.
*/
func (this *QQuickWindow) RendererInterface() *QSGRendererInterface /*777 QSGRendererInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow17rendererInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRendererInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSceneGraphBackend(QSGRendererInterface::GraphicsApi)

/*
Requests a Qt Quick scenegraph backend for the specified graphics api. Backends can either be built-in or be installed in form of dynamically loaded plugins.

Note: The call to the function must happen before constructing the first QQuickWindow in the application. It cannot be changed afterwards.

If the selected backend is invalid or an error occurs, the default backend (OpenGL or software, depending on the Qt configuration) is used.

This function was introduced in  Qt 5.8.

See also sceneGraphBackend().
*/
func (this *QQuickWindow) SetSceneGraphBackend(api int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow20setSceneGraphBackendEN20QSGRendererInterface11GraphicsApiE", qtrt.FFI_TYPE_POINTER, api)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetSceneGraphBackend(api int) {
	var nilthis *QQuickWindow
	nilthis.SetSceneGraphBackend(api)
}

// /usr/include/qt/QtQuick/qquickwindow.h:172
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void setSceneGraphBackend(const QString &)

/*
Requests a Qt Quick scenegraph backend for the specified graphics api. Backends can either be built-in or be installed in form of dynamically loaded plugins.

Note: The call to the function must happen before constructing the first QQuickWindow in the application. It cannot be changed afterwards.

If the selected backend is invalid or an error occurs, the default backend (OpenGL or software, depending on the Qt configuration) is used.

This function was introduced in  Qt 5.8.

See also sceneGraphBackend().
*/
func (this *QQuickWindow) SetSceneGraphBackend_1(backend string) {
	var tmpArg0 = qtcore.NewQString_5(backend)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow20setSceneGraphBackendERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetSceneGraphBackend_1(backend string) {
	var nilthis *QQuickWindow
	nilthis.SetSceneGraphBackend_1(backend)
}

// /usr/include/qt/QtQuick/qquickwindow.h:173
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sceneGraphBackend()

/*
Returns the requested Qt Quick scenegraph backend.

Note: The return value of this function may still be outdated by subsequent calls to setSceneGraphBackend() until the first QQuickWindow in the application has been constructed.

This function was introduced in  Qt 5.9.

See also setSceneGraphBackend().
*/
func (this *QQuickWindow) SceneGraphBackend() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17sceneGraphBackendEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QQuickWindow_SceneGraphBackend() string {
	var nilthis *QQuickWindow
	rv := nilthis.SceneGraphBackend()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRectangleNode * createRectangleNode() const

/*
Creates a simple rectangle node. When the scenegraph is not initialized, the return value is null.

This is cross-backend alternative to constructing a QSGSimpleRectNode directly.

This function was introduced in  Qt 5.8.

See also QSGRectangleNode.
*/
func (this *QQuickWindow) CreateRectangleNode() *QSGRectangleNode /*777 QSGRectangleNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createRectangleNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRectangleNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGImageNode * createImageNode() const

/*
Creates a simple image node. When the scenegraph is not initialized, the return value is null.

This is cross-backend alternative to constructing a QSGSimpleTextureNode directly.

This function was introduced in  Qt 5.8.

See also QSGImageNode.
*/
func (this *QQuickWindow) CreateImageNode() *QSGImageNode /*777 QSGImageNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow15createImageNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGImageNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGNinePatchNode * createNinePatchNode() const

/*
Creates a nine patch node. When the scenegraph is not initialized, the return value is null.

This function was introduced in  Qt 5.8.
*/
func (this *QQuickWindow) CreateNinePatchNode() *QSGNinePatchNode /*777 QSGNinePatchNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createNinePatchNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNinePatchNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [4] QQuickWindow::TextRenderType textRenderType()

/*
Returns the render type of text-like elements in Qt Quick. The default is QQuickWindow::QtTextRendering.

This function was introduced in  Qt 5.10.

See also setTextRenderType().
*/
func (this *QQuickWindow) TextRenderType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14textRenderTypeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QQuickWindow_TextRenderType() int {
	var nilthis *QQuickWindow
	rv := nilthis.TextRenderType()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:180
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTextRenderType(QQuickWindow::TextRenderType)

/*
Sets the default render type of text-like elements in Qt Quick to renderType.

Note: setting the render type will only affect elements created afterwards; the render type of existing elements will not be modified.

This function was introduced in  Qt 5.10.

See also textRenderType().
*/
func (this *QQuickWindow) SetTextRenderType(renderType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17setTextRenderTypeENS_14TextRenderTypeE", qtrt.FFI_TYPE_POINTER, renderType)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetTextRenderType(renderType int) {
	var nilthis *QQuickWindow
	nilthis.SetTextRenderType(renderType)
}

// /usr/include/qt/QtQuick/qquickwindow.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frameSwapped()

/*
This signal is emitted when a frame has been queued for presenting. With vertical synchronization enabled the signal is emitted at most once per vsync interval in a continuously animating scene.

This signal will be emitted from the scene graph rendering thread.
*/
func (this *QQuickWindow) FrameSwapped() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow12frameSwappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphInitialized()

/*
This signal is emitted when the scene graph has been initialized.

This signal will be emitted from the scene graph rendering thread.
*/
func (this *QQuickWindow) SceneGraphInitialized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphInitializedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphInvalidated()

/*
This signal is emitted when the scene graph has been invalidated.

This signal implies that the graphics rendering context used has been invalidated and all user resources tied to that context should be released.

In the case of the default OpenGL adaptation the context of this window will be bound when this function is called. The only exception is if the native OpenGL has been destroyed outside Qt's control, for instance through EGL_CONTEXT_LOST.

This signal will be emitted from the scene graph rendering thread.
*/
func (this *QQuickWindow) SceneGraphInvalidated() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphInvalidatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beforeSynchronizing()

/*
This signal is emitted before the scene graph is synchronized with the QML state.

This signal can be used to do any preparation required before calls to QQuickItem::updatePaintNode().

The OpenGL context used for rendering the scene graph will be bound at this point.

Warning: This signal is emitted from the scene graph rendering thread. If your slot function needs to finish before execution continues, you must make sure that the connection is direct (see Qt::ConnectionType).

Warning: Make very sure that a signal handler for beforeSynchronizing leaves the GL context in the same state as it was when the signal handler was entered. Failing to do so can result in the scene not rendering properly.

See also resetOpenGLState().
*/
func (this *QQuickWindow) BeforeSynchronizing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow19beforeSynchronizingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterSynchronizing()

/*
This signal is emitted after the scene graph is synchronized with the QML state.

This signal can be used to do preparation required after calls to QQuickItem::updatePaintNode(), while the GUI thread is still locked.

The graphics context used for rendering the scene graph will be bound at this point.

Warning: This signal is emitted from the scene graph rendering thread. If your slot function needs to finish before execution continues, you must make sure that the connection is direct (see Qt::ConnectionType).

Warning: When using the OpenGL adaptation, make sure that a signal handler for afterSynchronizing leaves the OpenGL context in the same state as it was when the signal handler was entered. Failing to do so can result in the scene not rendering properly.

This function was introduced in  Qt 5.3.

See also resetOpenGLState().
*/
func (this *QQuickWindow) AfterSynchronizing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow18afterSynchronizingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beforeRendering()

/*
This signal is emitted before the scene starts rendering.

Combined with the modes for clearing the background, this option can be used to paint using raw OpenGL under QML content.

The OpenGL context used for rendering the scene graph will be bound at this point.

Warning: This signal is emitted from the scene graph rendering thread. If your slot function needs to finish before execution continues, you must make sure that the connection is direct (see Qt::ConnectionType).

Warning: Make very sure that a signal handler for beforeRendering leaves the OpenGL context in the same state as it was when the signal handler was entered. Failing to do so can result in the scene not rendering properly.

See also resetOpenGLState().
*/
func (this *QQuickWindow) BeforeRendering() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15beforeRenderingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterRendering()

/*
This signal is emitted after the scene has completed rendering, before swapbuffers is called.

This signal can be used to paint using raw OpenGL on top of QML content, or to do screen scraping of the current frame buffer.

The OpenGL context used for rendering the scene graph will be bound at this point.

Warning: This signal is emitted from the scene graph rendering thread. If your slot function needs to finish before execution continues, you must make sure that the connection is direct (see Qt::ConnectionType).

Warning: Make very sure that a signal handler for afterRendering() leaves the OpenGL context in the same state as it was when the signal handler was entered. Failing to do so can result in the scene not rendering properly.

See also resetOpenGLState().
*/
func (this *QQuickWindow) AfterRendering() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14afterRenderingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterAnimating()

/*
This signal is emitted on the gui thread before requesting the render thread to perform the synchronization of the scene graph.

Unlike the other similar signals, this one is emitted on the gui thread instead of the render thread. It can be used to synchronize external animation systems with the QML content.

This function was introduced in  Qt 5.3.
*/
func (this *QQuickWindow) AfterAnimating() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14afterAnimatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphAboutToStop()

/*
This signal is emitted on the render thread when the scene graph is about to stop rendering. This happens usually because the window has been hidden.

Applications may use this signal to release resources, but should be prepared to reinstantiated them again fast. The scene graph and the graphics context are not released at this time.

Warning: This signal is emitted from the scene graph rendering thread. If your slot function needs to finish before execution continues, you must make sure that the connection is direct (see Qt::ConnectionType).

Warning: Make very sure that a signal handler for sceneGraphAboutToStop() leaves the graphics context in the same state as it was when the signal handler was entered. Failing to do so can result in the scene not rendering properly.

This function was introduced in  Qt 5.3.

See also sceneGraphInvalidated() and resetOpenGLState().
*/
func (this *QQuickWindow) SceneGraphAboutToStop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphAboutToStopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void colorChanged(const QColor &)

/*

 */
func (this *QQuickWindow) ColorChanged(arg0 qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QColor_PTR() != nil {
		convArg0 = arg0.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow12colorChangedERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusItemChanged()

/*

 */
func (this *QQuickWindow) ActiveFocusItemChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow22activeFocusItemChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphError(QQuickWindow::SceneGraphError, const QString &)

/*
This signal is emitted when an error occurred during scene graph initialization.

Applications should connect to this signal if they wish to handle errors, like graphics context creation failures, in a custom way. When no slot is connected to the signal, the behavior will be different: Quick will print the message, or show a message box, and terminate the application.

This signal will be emitted from the gui thread.

This function was introduced in  Qt 5.3.
*/
func (this *QQuickWindow) SceneGraphError(error int, message string) {
	var tmpArg1 = qtcore.NewQString_5(message)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15sceneGraphErrorENS_15SceneGraphErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()

/*
Schedules the window to render another frame.

Calling QQuickWindow::update() differs from QQuickItem::update() in that it always triggers a repaint, regardless of changes in the underlying scene graph or not.
*/
func (this *QQuickWindow) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseResources()

/*
This function tries to release redundant resources currently held by the QML scene.

Calling this function might result in the scene graph and the OpenGL context used for rendering being released to release graphics memory. If this happens, the sceneGraphInvalidated() signal will be called, allowing users to clean up their own graphics resources. The setPersistentOpenGLContext() and setPersistentSceneGraph() functions can be used to prevent this from happening, if handling the cleanup is not feasible in the application, at the cost of higher memory usage.

See also sceneGraphInvalidated(), setPersistentOpenGLContext(), and setPersistentSceneGraph().
*/
func (this *QQuickWindow) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:207
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void exposeEvent(QExposeEvent *)

/*
Reimplemented from QWindow::exposeEvent().
*/
func (this *QQuickWindow) ExposeEvent(arg0 qtgui.QExposeEvent_ITF /*777 QExposeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QExposeEvent_PTR() != nil {
		convArg0 = arg0.QExposeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow11exposeEventEP12QExposeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWindow::resizeEvent().
*/
func (this *QQuickWindow) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWindow::showEvent().
*/
func (this *QQuickWindow) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:211
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWindow::hideEvent().
*/
func (this *QQuickWindow) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWindow::focusInEvent().
*/
func (this *QQuickWindow) FocusInEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWindow::focusOutEvent().
*/
func (this *QQuickWindow) FocusOutEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QWindow::event().
*/
func (this *QQuickWindow) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWindow::keyPressEvent().
*/
func (this *QQuickWindow) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWindow::keyReleaseEvent().
*/
func (this *QQuickWindow) KeyReleaseEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mousePressEvent().
*/
func (this *QQuickWindow) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mouseReleaseEvent().
*/
func (this *QQuickWindow) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mouseDoubleClickEvent().
*/
func (this *QQuickWindow) MouseDoubleClickEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mouseMoveEvent().
*/
func (this *QQuickWindow) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:225
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWindow::wheelEvent().
*/
func (this *QQuickWindow) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QQuickWindow__CreateTextureOption = int

//
const QQuickWindow__TextureHasAlphaChannel QQuickWindow__CreateTextureOption = 1

//
const QQuickWindow__TextureHasMipmaps QQuickWindow__CreateTextureOption = 2

//
const QQuickWindow__TextureOwnsGLTexture QQuickWindow__CreateTextureOption = 4

//
const QQuickWindow__TextureCanUseAtlas QQuickWindow__CreateTextureOption = 8

//
const QQuickWindow__TextureIsOpaque QQuickWindow__CreateTextureOption = 16

func (this *QQuickWindow) CreateTextureOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWindow_CreateTextureOptionItemName(val int) string {
	var nilthis *QQuickWindow
	return nilthis.CreateTextureOptionItemName(val)
}

/*


This enum was introduced or modified in  Qt 5.4.

See also Scene Graph and Rendering.

*/
type QQuickWindow__RenderStage = int

// Before synchronization.
const QQuickWindow__BeforeSynchronizingStage QQuickWindow__RenderStage = 0

// After synchronization.
const QQuickWindow__AfterSynchronizingStage QQuickWindow__RenderStage = 1

// Before rendering.
const QQuickWindow__BeforeRenderingStage QQuickWindow__RenderStage = 2

// After rendering.
const QQuickWindow__AfterRenderingStage QQuickWindow__RenderStage = 3

// After the frame is swapped.
const QQuickWindow__AfterSwapStage QQuickWindow__RenderStage = 4

//
const QQuickWindow__NoStage QQuickWindow__RenderStage = 5

func (this *QQuickWindow) RenderStageItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWindow_RenderStageItemName(val int) string {
	var nilthis *QQuickWindow
	return nilthis.RenderStageItemName(val)
}

/*
This enum describes the error in a sceneGraphError() signal.



This enum was introduced or modified in  Qt 5.3.

*/
type QQuickWindow__SceneGraphError = int

//
const QQuickWindow__ContextNotAvailable QQuickWindow__SceneGraphError = 1

func (this *QQuickWindow) SceneGraphErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWindow_SceneGraphErrorItemName(val int) string {
	var nilthis *QQuickWindow
	return nilthis.SceneGraphErrorItemName(val)
}

/*
This enum describes the default render type of text-like elements in Qt Quick (Text, TextInput, etc.).

Select NativeTextRendering if you prefer text to look native on the target platform and do not require advanced features such as transformation of the text. Using such features in combination with the NativeTextRendering render type will lend poor and sometimes pixelated results.



This enum was introduced or modified in  Qt 5.10.

*/
type QQuickWindow__TextRenderType = int

// Use Qt's own rasterization algorithm.
const QQuickWindow__QtTextRendering QQuickWindow__TextRenderType = 0

// Use the operating system's native rasterizer for text.
const QQuickWindow__NativeTextRendering QQuickWindow__TextRenderType = 1

func (this *QQuickWindow) TextRenderTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWindow_TextRenderTypeItemName(val int) string {
	var nilthis *QQuickWindow
	return nilthis.TextRenderTypeItemName(val)
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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
