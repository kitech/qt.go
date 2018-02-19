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
func (this *QSGSimpleTextureNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:59
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect() const
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
func (this *QSGSimpleTextureNode) SetSourceRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode13setSourceRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:63
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF sourceRect() const
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
func (this *QSGSimpleTextureNode) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode7textureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltering(QSGTexture::Filtering)
func (this *QSGSimpleTextureNode) SetFiltering(filtering int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode12setFilteringEN10QSGTexture9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::Filtering filtering() const
func (this *QSGSimpleTextureNode) Filtering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode9filteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextureCoordinatesTransform(QSGSimpleTextureNode::TextureCoordinatesTransformMode)
func (this *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode30setTextureCoordinatesTransformE6QFlagsINS_31TextureCoordinatesTransformFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGSimpleTextureNode::TextureCoordinatesTransformMode textureCoordinatesTransform() const
func (this *QSGSimpleTextureNode) TextureCoordinatesTransform() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode27textureCoordinatesTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOwnsTexture(_Bool)
func (this *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode14setOwnsTextureEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), owns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownsTexture() const
func (this *QSGSimpleTextureNode) OwnsTexture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode11ownsTextureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QSGSimpleTextureNode__TextureCoordinatesTransformFlag = int

const QSGSimpleTextureNode__NoTransform QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 0
const QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 1
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
