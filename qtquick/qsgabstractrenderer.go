package qtquick

// /usr/include/qt/QtQuick/qsgabstractrenderer.h
// #include <qsgabstractrenderer.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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

// void nodeChanged(QSGNode *, QSGNode::DirtyState)
func (this *QSGAbstractRenderer) InheritNodeChanged(f func(node *QSGNode /*777 QSGNode **/, state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "nodeChanged", f)
}

/*

 */
type QSGAbstractRenderer struct {
	*qtcore.QObject
}
type QSGAbstractRenderer_ITF interface {
	qtcore.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func (ptr *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer { return ptr }

func (this *QSGAbstractRenderer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSGAbstractRenderer) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSGAbstractRendererFromPointer(cthis unsafe.Pointer) *QSGAbstractRenderer {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSGAbstractRenderer{bcthis0}
}
func (*QSGAbstractRenderer) NewFromPointer(cthis unsafe.Pointer) *QSGAbstractRenderer {
	return NewQSGAbstractRendererFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSGAbstractRenderer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGAbstractRenderer()

/*

 */
func DeleteQSGAbstractRenderer(this *QSGAbstractRenderer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRendererD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootNode(QSGRootNode *)

/*
Sets the node as the root of the QSGNode scene that you want to render. You need to provide a node before trying to render the scene.

Note: This doesn't take ownership of node.

See also rootNode().
*/
func (this *QSGAbstractRenderer) SetRootNode(node QSGRootNode_ITF /*777 QSGRootNode **/) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGRootNode_PTR() != nil {
		convArg0 = node.QSGRootNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer11setRootNodeEP11QSGRootNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRootNode * rootNode() const

/*
Returns the root of the QSGNode scene.

See also setRootNode().
*/
func (this *QSGAbstractRenderer) RootNode() *QSGRootNode /*777 QSGRootNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer8rootNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRootNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDeviceRect(const QRect &)

/*
Sets rect as the geometry of the surface being rendered to.

See also deviceRect().
*/
func (this *QSGAbstractRenderer) SetDeviceRect(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer13setDeviceRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:70
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setDeviceRect(const QSize &)

/*
Sets rect as the geometry of the surface being rendered to.

See also deviceRect().
*/
func (this *QSGAbstractRenderer) SetDeviceRect_1(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer13setDeviceRectERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect deviceRect() const

/*
Returns the device rect of the surface being rendered to.

See also setDeviceRect().
*/
func (this *QSGAbstractRenderer) DeviceRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer10deviceRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewportRect(const QRect &)

/*
Sets rect as the geometry of the viewport to render on the surface.

See also viewportRect().
*/
func (this *QSGAbstractRenderer) SetViewportRect(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer15setViewportRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setViewportRect(const QSize &)

/*
Sets rect as the geometry of the viewport to render on the surface.

See also viewportRect().
*/
func (this *QSGAbstractRenderer) SetViewportRect_1(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer15setViewportRectERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:75
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewportRect() const

/*
Returns the rect of the viewport to render.

See also setViewportRect().
*/
func (this *QSGAbstractRenderer) ViewportRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer12viewportRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProjectionMatrixToRect(const QRectF &)

/*
Convenience method that calls setProjectionMatrix() with an orthographic matrix generated from rect.

See also setProjectionMatrix() and projectionMatrix().
*/
func (this *QSGAbstractRenderer) SetProjectionMatrixToRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer25setProjectionMatrixToRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProjectionMatrix(const QMatrix4x4 &)

/*
Use matrix to project the QSGNode coordinates onto surface pixels.

See also projectionMatrix() and setProjectionMatrixToRect().
*/
func (this *QSGAbstractRenderer) SetProjectionMatrix(matrix qtgui.QMatrix4x4_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix4x4_PTR() != nil {
		convArg0 = matrix.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer19setProjectionMatrixERK10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:79
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 projectionMatrix() const

/*
Returns the projection matrix

See also setProjectionMatrix() and setProjectionMatrixToRect().
*/
func (this *QSGAbstractRenderer) ProjectionMatrix() *qtgui.QMatrix4x4 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer16projectionMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearColor(const QColor &)

/*
Use color to clear the framebuffer when clearMode() is set to QSGAbstractRenderer::ClearColorBuffer.

See also clearColor() and setClearMode().
*/
func (this *QSGAbstractRenderer) SetClearColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer13setClearColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor clearColor() const

/*
Returns the color that clears the framebuffer at the beginning of the rendering.

See also setClearColor() and clearMode().
*/
func (this *QSGAbstractRenderer) ClearColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer10clearColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearMode(QSGAbstractRenderer::ClearMode)

/*
Defines which attachment of the framebuffer should be cleared before each scene render with the mode flag.

See also clearMode() and setClearColor().
*/
func (this *QSGAbstractRenderer) SetClearMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer12setClearModeE6QFlagsINS_12ClearModeBitEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGAbstractRenderer::ClearMode clearMode() const

/*
Flags defining which attachment of the framebuffer will be cleared before each scene render.

See also setClearMode() and clearColor().
*/
func (this *QSGAbstractRenderer) ClearMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer9clearModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void renderScene(uint)

/*
Render the scene to the specified fboId

If fboId isn't specified, the scene graph will be rendered to the default framebuffer. You will have to call QOpenGLContext::swapBuffers() yourself afterward.

The framebuffer specified by fboId will be bound automatically.

See also QOpenGLContext::swapBuffers() and QOpenGLFramebufferObject::handle().
*/
func (this *QSGAbstractRenderer) RenderScene(fboId uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer11renderSceneEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fboId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void renderScene(uint)

/*
Render the scene to the specified fboId

If fboId isn't specified, the scene graph will be rendered to the default framebuffer. You will have to call QOpenGLContext::swapBuffers() yourself afterward.

The framebuffer specified by fboId will be bound automatically.

See also QOpenGLContext::swapBuffers() and QOpenGLFramebufferObject::handle().
*/
func (this *QSGAbstractRenderer) RenderScene__() {
	// arg: 0, uint=Typedef, uint=Typedef, unsigned int, UInt
	fboId := uint(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer11renderSceneEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fboId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphChanged()

/*
This signal is emitted on the first modification of a node in the tree after the last scene render.
*/
func (this *QSGAbstractRenderer) SceneGraphChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer17sceneGraphChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:93
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSGAbstractRenderer(QObject *)

/*

 */
func NewQSGAbstractRenderer(parent qtcore.QObject_ITF /*777 QObject **/) *QSGAbstractRenderer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRendererC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGAbstractRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSGAbstractRenderer")
	return gothis
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:93
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSGAbstractRenderer(QObject *)

/*

 */
func NewQSGAbstractRenderer__() *QSGAbstractRenderer {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRendererC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGAbstractRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSGAbstractRenderer")
	return gothis
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:94
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void nodeChanged(QSGNode *, QSGNode::DirtyState)

/*

 */
func (this *QSGAbstractRenderer) NodeChanged(node QSGNode_ITF /*777 QSGNode **/, state int) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QSGAbstractRenderer11nodeChangedEP7QSGNode6QFlagsINS0_13DirtyStateBitEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, state)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QSGAbstractRenderer__ClearModeBit = int

//
const QSGAbstractRenderer__ClearColorBuffer QSGAbstractRenderer__ClearModeBit = 1

//
const QSGAbstractRenderer__ClearDepthBuffer QSGAbstractRenderer__ClearModeBit = 2

//
const QSGAbstractRenderer__ClearStencilBuffer QSGAbstractRenderer__ClearModeBit = 4

func (this *QSGAbstractRenderer) ClearModeBitItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSGAbstractRenderer_ClearModeBitItemName(val int) string {
	var nilthis *QSGAbstractRenderer
	return nilthis.ClearModeBitItemName(val)
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
