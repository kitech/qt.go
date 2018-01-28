package qtquick

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h
// #include <qsgsimpletexturenode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSGSimpleTextureNode struct {
	*QSGGeometryNode
}

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
// Public
// void QSGSimpleTextureNode()
func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	cthis := qtrt.Calloc(1, 256) // 384
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNodeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGSimpleTextureNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:55
// index:0
// Public virtual
// void ~QSGSimpleTextureNode()
func DeleteQSGSimpleTextureNode(*QSGSimpleTextureNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:57
// index:0
// Public
// void setRect(const QRectF &)
func (this *QSGSimpleTextureNode) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:58
// index:1
// Public inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QSGSimpleTextureNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:59
// index:0
// Public
// QRectF rect()
func (this *QSGSimpleTextureNode) Rect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:61
// index:0
// Public
// void setSourceRect(const QRectF &)
func (this *QSGSimpleTextureNode) SetSourceRect(r *qtcore.QRectF) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode13setSourceRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:62
// index:1
// Public inline
// void setSourceRect(qreal, qreal, qreal, qreal)
func (this *QSGSimpleTextureNode) SetSourceRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode13setSourceRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:63
// index:0
// Public
// QRectF sourceRect()
func (this *QSGSimpleTextureNode) SourceRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode10sourceRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:65
// index:0
// Public
// void setTexture(QSGTexture *)
func (this *QSGSimpleTextureNode) SetTexture(texture *QSGTexture /*777 QSGTexture **/) {
	var convArg0 = texture.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode10setTextureEP10QSGTexture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:66
// index:0
// Public
// QSGTexture * texture()
func (this *QSGSimpleTextureNode) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode7textureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:68
// index:0
// Public
// void setFiltering(QSGTexture::Filtering)
func (this *QSGSimpleTextureNode) SetFiltering(filtering int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode12setFilteringEN10QSGTexture9FilteringE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:69
// index:0
// Public
// QSGTexture::Filtering filtering()
func (this *QSGSimpleTextureNode) Filtering() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode9filteringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:79
// index:0
// Public
// QSGSimpleTextureNode::TextureCoordinatesTransformMode textureCoordinatesTransform()
func (this *QSGSimpleTextureNode) TextureCoordinatesTransform() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode27textureCoordinatesTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:81
// index:0
// Public
// void setOwnsTexture(_Bool)
func (this *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGSimpleTextureNode14setOwnsTextureEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), owns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:82
// index:0
// Public
// bool ownsTexture()
func (this *QSGSimpleTextureNode) OwnsTexture() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGSimpleTextureNode11ownsTextureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QSGSimpleTextureNode__TextureCoordinatesTransformFlag = int

const QSGSimpleTextureNode__NoTransform QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 0
const QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 1
const QSGSimpleTextureNode__MirrorVertically QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 2

//  body block end
