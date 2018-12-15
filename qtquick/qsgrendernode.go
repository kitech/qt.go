package qtquick

// /usr/include/qt/QtQuick/qsgrendernode.h
// #include <qsgrendernode.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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

/*

 */
type QSGRenderNode struct {
	*QSGNode
}
type QSGRenderNode_ITF interface {
	QSGNode_ITF
	QSGRenderNode_PTR() *QSGRenderNode
}

func (ptr *QSGRenderNode) QSGRenderNode_PTR() *QSGRenderNode { return ptr }

func (this *QSGRenderNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGRenderNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGRenderNodeFromPointer(cthis unsafe.Pointer) *QSGRenderNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGRenderNode{bcthis0}
}
func (*QSGRenderNode) NewFromPointer(cthis unsafe.Pointer) *QSGRenderNode {
	return NewQSGRenderNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGRenderNode()

/*
Default constructs an instance of QSGRenderNode.
*/
func (*QSGRenderNode) NewForInherit() *QSGRenderNode {
	return NewQSGRenderNode()
}
func NewQSGRenderNode() *QSGRenderNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSGRenderNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGRenderNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGRenderNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgrendernode.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGRenderNode()

/*

 */
func DeleteQSGRenderNode(this *QSGRenderNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSGRenderNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSGRenderNode::StateFlags changedStates() const

/*
This function should return a mask where each bit represents graphics states changed by the render() function:


DepthState - depth write mask, depth test enabled, depth comparison function
StencilState - stencil write masks, stencil test enabled, stencil operations, stencil comparison functions
ScissorState - scissor enabled, scissor test enabled
ColorState - clear color, color write mask
BlendState - blend enabled, blend function
CullState - front face, cull face enabled
ViewportState - viewport
RenderTargetState - render target


The function is called by the renderer so it can reset the states after rendering this node. This makes the implementation of render() simpler since it does not have to query and restore these states.

The default implementation returns 0, meaning no relevant state was changed in render().

With APIs other than OpenGL the relevant states are only those that are set via the command list (for example, OMSetRenderTargets, RSSetViewports, RSSetScissorRects, OMSetBlendFactor, OMSetStencilRef in case of D3D12), and only when such commands were added to the scenegraph's command list queried via the QSGRendererInterface::CommandList resource enum. States set in pipeline state objects do not need to be reported here. Similarly, draw call related settings (root signature, descriptor heaps, etc.) are always set again by the scenegraph so render() can freely change them.

The software backend exposes its QPainter and saves and restores before and after invoking render(). Therefore reporting any changed states from here is not necessary.

Note: This function may be called before render().
*/
func (this *QSGRenderNode) ChangedStates() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode13changedStatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void releaseResources()

/*
This function is called when all custom graphics resources allocated by this node have to be freed immediately. In case the node does not directly allocate graphics resources (buffers, textures, render targets, fences, etc.) through the graphics API that is in use, there is nothing to do here.

Failing to release all custom resources can lead to incorrect behavior in graphics device loss scenarios on some systems since subsequent reinitialization of the graphics system may fail.

Note: Some scenegraph backends may choose not to call this function. Therefore it is expected that QSGRenderNode implementations perform cleanup both in their destructor and in releaseResources().

Unlike with the destructor, it is expected that render() can reinitialize all resources it needs when called after a call to releaseResources().

With OpenGL, the scenegraph's OpenGL context will be current both when calling the destructor and this function.
*/
func (this *QSGRenderNode) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSGRenderNode16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSGRenderNode::RenderingFlags flags() const

/*
Returns flags describing the behavior of this render node.

The default implementation returns 0.

See also RenderingFlag and rect().
*/
func (this *QSGRenderNode) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF rect() const

/*
Returns the bounding rectangle in item coordinates for the area render() touches. The value is only in use when flags() includes BoundedRectRendering, ignored otherwise.

Reporting the rectangle in combination with BoundedRectRendering is particularly important with the software backend because otherwise having a rendernode in the scene would trigger fullscreen updates, skipping all partial update optimizations.

For rendernodes covering the entire area of a corresponding QQuickItem the return value will be (0, 0, item->width(), item->height()).

See also flags().
*/
func (this *QSGRenderNode) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendernode.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMatrix4x4 * matrix() const

/*
Returns pointer to the current model-view matrix.
*/
func (this *QSGRenderNode) Matrix() *qtgui.QMatrix4x4 /*777 const QMatrix4x4 **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgrendernode.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] const QSGClipNode * clipList() const

/*
Returns the current clip list.
*/
func (this *QSGRenderNode) ClipList() *QSGClipNode /*777 const QSGClipNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode8clipListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGClipNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgrendernode.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal inheritedOpacity() const

/*
Returns the current effective opacity.
*/
func (this *QSGRenderNode) InheritedOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode16inheritedOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

/*


 */
type QSGRenderNode__StateFlag = int

//
const QSGRenderNode__DepthState QSGRenderNode__StateFlag = 1

//
const QSGRenderNode__StencilState QSGRenderNode__StateFlag = 2

//
const QSGRenderNode__ScissorState QSGRenderNode__StateFlag = 4

//
const QSGRenderNode__ColorState QSGRenderNode__StateFlag = 8

//
const QSGRenderNode__BlendState QSGRenderNode__StateFlag = 16

//
const QSGRenderNode__CullState QSGRenderNode__StateFlag = 32

//
const QSGRenderNode__ViewportState QSGRenderNode__StateFlag = 64

//
const QSGRenderNode__RenderTargetState QSGRenderNode__StateFlag = 128

func (this *QSGRenderNode) StateFlagItemName(val int) string {
	switch val {
	case QSGRenderNode__DepthState: // 1
		return "DepthState"
	case QSGRenderNode__StencilState: // 2
		return "StencilState"
	case QSGRenderNode__ScissorState: // 4
		return "ScissorState"
	case QSGRenderNode__ColorState: // 8
		return "ColorState"
	case QSGRenderNode__BlendState: // 16
		return "BlendState"
	case QSGRenderNode__CullState: // 32
		return "CullState"
	case QSGRenderNode__ViewportState: // 64
		return "ViewportState"
	case QSGRenderNode__RenderTargetState: // 128
		return "RenderTargetState"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRenderNode_StateFlagItemName(val int) string {
	var nilthis *QSGRenderNode
	return nilthis.StateFlagItemName(val)
}

/*


 */
type QSGRenderNode__RenderingFlag = int

//
const QSGRenderNode__BoundedRectRendering QSGRenderNode__RenderingFlag = 1

//
const QSGRenderNode__DepthAwareRendering QSGRenderNode__RenderingFlag = 2

//
const QSGRenderNode__OpaqueRendering QSGRenderNode__RenderingFlag = 4

func (this *QSGRenderNode) RenderingFlagItemName(val int) string {
	switch val {
	case QSGRenderNode__BoundedRectRendering: // 1
		return "BoundedRectRendering"
	case QSGRenderNode__DepthAwareRendering: // 2
		return "DepthAwareRendering"
	case QSGRenderNode__OpaqueRendering: // 4
		return "OpaqueRendering"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGRenderNode_RenderingFlagItemName(val int) string {
	var nilthis *QSGRenderNode
	return nilthis.RenderingFlagItemName(val)
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
