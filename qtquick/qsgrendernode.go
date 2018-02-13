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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

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
func DeleteQSGRenderNode(this *QSGRenderNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSGRenderNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSGRenderNode::StateFlags changedStates()
func (this *QSGRenderNode) ChangedStates() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode13changedStatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void releaseResources()
func (this *QSGRenderNode) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSGRenderNode16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSGRenderNode::RenderingFlags flags()
func (this *QSGRenderNode) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF rect()
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
// [8] const QMatrix4x4 * matrix()
func (this *QSGRenderNode) Matrix() *qtgui.QMatrix4x4 /*777 const QMatrix4x4 **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgrendernode.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] const QSGClipNode * clipList()
func (this *QSGRenderNode) ClipList() *QSGClipNode /*777 const QSGClipNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode8clipListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGClipNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgrendernode.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal inheritedOpacity()
func (this *QSGRenderNode) InheritedOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSGRenderNode16inheritedOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

type QSGRenderNode__StateFlag = int

const QSGRenderNode__DepthState QSGRenderNode__StateFlag = 1
const QSGRenderNode__StencilState QSGRenderNode__StateFlag = 2
const QSGRenderNode__ScissorState QSGRenderNode__StateFlag = 4
const QSGRenderNode__ColorState QSGRenderNode__StateFlag = 8
const QSGRenderNode__BlendState QSGRenderNode__StateFlag = 16
const QSGRenderNode__CullState QSGRenderNode__StateFlag = 32
const QSGRenderNode__ViewportState QSGRenderNode__StateFlag = 64
const QSGRenderNode__RenderTargetState QSGRenderNode__StateFlag = 128

type QSGRenderNode__RenderingFlag = int

const QSGRenderNode__BoundedRectRendering QSGRenderNode__RenderingFlag = 1
const QSGRenderNode__DepthAwareRendering QSGRenderNode__RenderingFlag = 2
const QSGRenderNode__OpaqueRendering QSGRenderNode__RenderingFlag = 4

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
