package qtquick

// /usr/include/qt/QtQuick/qsgimagenode.h
// #include <qsgimagenode.h>
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
type QSGImageNode struct {
	*QSGGeometryNode
}

func (this *QSGImageNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGGeometryNode.GetCthis()
	}
}
func (this *QSGImageNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGGeometryNode = NewQSGGeometryNodeFromPointer(cthis)
}
func NewQSGImageNodeFromPointer(cthis unsafe.Pointer) *QSGImageNode {
	bcthis0 := NewQSGGeometryNodeFromPointer(cthis)
	return &QSGImageNode{bcthis0}
}
func (*QSGImageNode) NewFromPointer(cthis unsafe.Pointer) *QSGImageNode {
	return NewQSGImageNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:51
// index:0
// Public inline virtual
// void ~QSGImageNode()
func DeleteQSGImageNode(*QSGImageNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:53
// index:0
// Public pure virtual
// void setRect(const QRectF &)
func (this *QSGImageNode) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:54
// index:1
// Public inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QSGImageNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:55
// index:0
// Public pure virtual
// QRectF rect()
func (this *QSGImageNode) Rect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgimagenode.h:57
// index:0
// Public pure virtual
// void setSourceRect(const QRectF &)
func (this *QSGImageNode) SetSourceRect(r *qtcore.QRectF) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode13setSourceRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:58
// index:1
// Public inline
// void setSourceRect(qreal, qreal, qreal, qreal)
func (this *QSGImageNode) SetSourceRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode13setSourceRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:59
// index:0
// Public pure virtual
// QRectF sourceRect()
func (this *QSGImageNode) SourceRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode10sourceRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgimagenode.h:61
// index:0
// Public pure virtual
// void setTexture(QSGTexture *)
func (this *QSGImageNode) SetTexture(texture *QSGTexture /*777 QSGTexture **/) {
	var convArg0 = texture.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode10setTextureEP10QSGTexture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:62
// index:0
// Public pure virtual
// QSGTexture * texture()
func (this *QSGImageNode) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode7textureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgimagenode.h:64
// index:0
// Public pure virtual
// void setFiltering(QSGTexture::Filtering)
func (this *QSGImageNode) SetFiltering(filtering int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode12setFilteringEN10QSGTexture9FilteringE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:65
// index:0
// Public pure virtual
// QSGTexture::Filtering filtering()
func (this *QSGImageNode) Filtering() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode9filteringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:67
// index:0
// Public pure virtual
// void setMipmapFiltering(QSGTexture::Filtering)
func (this *QSGImageNode) SetMipmapFiltering(filtering int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode18setMipmapFilteringEN10QSGTexture9FilteringE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filtering)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:68
// index:0
// Public pure virtual
// QSGTexture::Filtering mipmapFiltering()
func (this *QSGImageNode) MipmapFiltering() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode15mipmapFilteringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:80
// index:0
// Public pure virtual
// QSGImageNode::TextureCoordinatesTransformMode textureCoordinatesTransform()
func (this *QSGImageNode) TextureCoordinatesTransform() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode27textureCoordinatesTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:82
// index:0
// Public pure virtual
// void setOwnsTexture(bool)
func (this *QSGImageNode) SetOwnsTexture(owns bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QSGImageNode14setOwnsTextureEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), owns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgimagenode.h:83
// index:0
// Public pure virtual
// bool ownsTexture()
func (this *QSGImageNode) OwnsTexture() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QSGImageNode11ownsTextureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QSGImageNode__TextureCoordinatesTransformFlag = int

const QSGImageNode__NoTransform QSGImageNode__TextureCoordinatesTransformFlag = 0
const QSGImageNode__MirrorHorizontally QSGImageNode__TextureCoordinatesTransformFlag = 1
const QSGImageNode__MirrorVertically QSGImageNode__TextureCoordinatesTransformFlag = 2

//  body block end
