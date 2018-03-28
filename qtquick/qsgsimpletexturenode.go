package qtquick

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h
// #include <qsgsimpletexturenode.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QSGSimpleTextureNode struct {
	*QSGGeometryNode
}
type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func (ptr *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode { return ptr }

func (this *QSGSimpleTextureNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGGeometryNode.GetCthis()
	}
}
func (this *QSGSimpleTextureNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGGeometryNode = NewQSGGeometryNodeFromPointer(cthis)
}
func NewQSGSimpleTextureNodeFromPointer(cthis unsafe.Pointer) *QSGSimpleTextureNode {
	bcthis0 := NewQSGGeometryNodeFromPointer(cthis)
	return &QSGSimpleTextureNode{bcthis0}
}
func (*QSGSimpleTextureNode) NewFromPointer(cthis unsafe.Pointer) *QSGSimpleTextureNode {
	return NewQSGSimpleTextureNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGSimpleTextureNode()

/*
Constructs a new simple texture node
*/
func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGSimpleTextureNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGSimpleTextureNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGSimpleTextureNode()

/*

 */
func DeleteQSGSimpleTextureNode(this *QSGSimpleTextureNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 384)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)

/*
Sets the target rect of this texture node to r.

See also rect().
*/
func (this *QSGSimpleTextureNode) SetRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode7setRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:58
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)

/*
Sets the target rect of this texture node to r.

See also rect().
*/
func (this *QSGSimpleTextureNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:59
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect() const

/*
Returns the target rect of this texture node.

See also setRect().
*/
func (this *QSGSimpleTextureNode) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceRect(const QRectF &)

/*
Sets the source rect of this texture node to r.

This function was introduced in  Qt 5.5.

See also sourceRect().
*/
func (this *QSGSimpleTextureNode) SetSourceRect(r qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode13setSourceRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:62
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSourceRect(qreal, qreal, qreal, qreal)

/*
Sets the source rect of this texture node to r.

This function was introduced in  Qt 5.5.

See also sourceRect().
*/
func (this *QSGSimpleTextureNode) SetSourceRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode13setSourceRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:63
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF sourceRect() const

/*
Returns the source rect of this texture node.

This function was introduced in  Qt 5.5.

See also setSourceRect().
*/
func (this *QSGSimpleTextureNode) SourceRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode10sourceRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTexture(QSGTexture *)

/*
Sets the texture of this texture node to texture.

Use setOwnsTexture() to set whether the node should take ownership of the texture. By default, the node does not take ownership.

Warning: A texture node must have a texture before being added to the scenegraph to be rendered.

See also texture().
*/
func (this *QSGSimpleTextureNode) SetTexture(texture QSGTexture_ITF /*777 QSGTexture **/) {
	var convArg0 unsafe.Pointer
	if texture != nil && texture.QSGTexture_PTR() != nil {
		convArg0 = texture.QSGTexture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode10setTextureEP10QSGTexture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * texture() const

/*
Returns the texture for this texture node

See also setTexture().
*/
func (this *QSGSimpleTextureNode) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode7textureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltering(QSGTexture::Filtering)

/*
Sets the filtering to be used for this texture node to filtering.

For smooth scaling, use QSGTexture::Linear; for normal scaling, use QSGTexture::Nearest.

See also filtering().
*/
func (this *QSGSimpleTextureNode) SetFiltering(filtering int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode12setFilteringEN10QSGTexture9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::Filtering filtering() const

/*
Returns the filtering currently set on this texture node

See also setFiltering().
*/
func (this *QSGSimpleTextureNode) Filtering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode9filteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextureCoordinatesTransform(QSGSimpleTextureNode::TextureCoordinatesTransformMode)

/*
Sets the method used to generate texture coordinates to mode. This can be used to obtain correct orientation of the texture. This is commonly needed when using a third party OpenGL library to render to texture as OpenGL has an inverted y-axis relative to Qt Quick.

See also textureCoordinatesTransform().
*/
func (this *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode30setTextureCoordinatesTransformE6QFlagsINS_31TextureCoordinatesTransformFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGSimpleTextureNode::TextureCoordinatesTransformMode textureCoordinatesTransform() const

/*
Returns the mode used to generate texture coordinates for this node.

See also setTextureCoordinatesTransform().
*/
func (this *QSGSimpleTextureNode) TextureCoordinatesTransform() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode27textureCoordinatesTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOwnsTexture(bool)

/*
Sets whether the node takes ownership of the texture to owns.

By default, the node does not take ownership of the texture.

This function was introduced in  Qt 5.4.

See also ownsTexture() and setTexture().
*/
func (this *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode14setOwnsTextureEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), owns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsTexture() const

/*
Returns true if the node takes ownership of the texture; otherwise returns false.

This function was introduced in  Qt 5.4.

See also setOwnsTexture().
*/
func (this *QSGSimpleTextureNode) OwnsTexture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode11ownsTextureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag = int

//
const QSGSimpleTextureNode__NoTransform QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 0

//
const QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 1

//
const QSGSimpleTextureNode__MirrorVertically QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 2

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
