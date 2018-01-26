package qtquick

// /usr/include/qt/QtQuick/qquickpainteditem.h
// #include <qquickpainteditem.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QQuickPaintedItem struct {
	*QQuickItem
}

func (this *QQuickPaintedItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQuickItem.GetCthis()
	}
}
func (this *QQuickPaintedItem) SetCthis(cthis unsafe.Pointer) {
	this.QQuickItem = NewQQuickItemFromPointer(cthis)
}
func NewQQuickPaintedItemFromPointer(cthis unsafe.Pointer) *QQuickPaintedItem {
	bcthis0 := NewQQuickItemFromPointer(cthis)
	return &QQuickPaintedItem{bcthis0}
}
func (*QQuickPaintedItem) NewFromPointer(cthis unsafe.Pointer) *QQuickPaintedItem {
	return NewQQuickPaintedItemFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:51
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickPaintedItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:60
// index:0
// Public
// void QQuickPaintedItem(class QQuickItem *)
func NewQQuickPaintedItem(parent *QQuickItem /*777 QQuickItem **/) *QQuickPaintedItem {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItemC1EP10QQuickItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickPaintedItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:61
// index:0
// Public virtual
// void ~QQuickPaintedItem()
func DeleteQQuickPaintedItem(*QQuickPaintedItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:75
// index:0
// Public
// void update(const class QRect &)
func (this *QQuickPaintedItem) Update(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem6updateERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:77
// index:0
// Public
// bool opaquePainting()
func (this *QQuickPaintedItem) OpaquePainting() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem14opaquePaintingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:78
// index:0
// Public
// void setOpaquePainting(_Bool)
func (this *QQuickPaintedItem) SetOpaquePainting(opaque bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem17setOpaquePaintingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:80
// index:0
// Public
// bool antialiasing()
func (this *QQuickPaintedItem) Antialiasing() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem12antialiasingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:81
// index:0
// Public
// void setAntialiasing(_Bool)
func (this *QQuickPaintedItem) SetAntialiasing(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem15setAntialiasingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:83
// index:0
// Public
// bool mipmap()
func (this *QQuickPaintedItem) Mipmap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem6mipmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:84
// index:0
// Public
// void setMipmap(_Bool)
func (this *QQuickPaintedItem) SetMipmap(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem9setMipmapEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:86
// index:0
// Public
// QQuickPaintedItem::PerformanceHints performanceHints()
func (this *QQuickPaintedItem) PerformanceHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem16performanceHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:87
// index:0
// Public
// void setPerformanceHint(enum QQuickPaintedItem::PerformanceHint, _Bool)
func (this *QQuickPaintedItem) SetPerformanceHint(hint int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem18setPerformanceHintENS_15PerformanceHintEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:90
// index:0
// Public
// QRectF contentsBoundingRect()
func (this *QQuickPaintedItem) ContentsBoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem20contentsBoundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:92
// index:0
// Public
// QSize contentsSize()
func (this *QQuickPaintedItem) ContentsSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem12contentsSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:93
// index:0
// Public
// void setContentsSize(const class QSize &)
func (this *QQuickPaintedItem) SetContentsSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem15setContentsSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:94
// index:0
// Public
// void resetContentsSize()
func (this *QQuickPaintedItem) ResetContentsSize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem17resetContentsSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:96
// index:0
// Public
// qreal contentsScale()
func (this *QQuickPaintedItem) ContentsScale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem13contentsScaleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:97
// index:0
// Public
// void setContentsScale(qreal)
func (this *QQuickPaintedItem) SetContentsScale(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem16setContentsScaleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:99
// index:0
// Public
// QSize textureSize()
func (this *QQuickPaintedItem) TextureSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem11textureSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:100
// index:0
// Public
// void setTextureSize(const class QSize &)
func (this *QQuickPaintedItem) SetTextureSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem14setTextureSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:102
// index:0
// Public
// QColor fillColor()
func (this *QQuickPaintedItem) FillColor() *qtgui.QColor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem9fillColorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:103
// index:0
// Public
// void setFillColor(const class QColor &)
func (this *QQuickPaintedItem) SetFillColor(arg0 *qtgui.QColor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem12setFillColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:105
// index:0
// Public
// QQuickPaintedItem::RenderTarget renderTarget()
func (this *QQuickPaintedItem) RenderTarget() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem12renderTargetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:106
// index:0
// Public
// void setRenderTarget(enum QQuickPaintedItem::RenderTarget)
func (this *QQuickPaintedItem) SetRenderTarget(target int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem15setRenderTargetENS_12RenderTargetE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:108
// index:0
// Public pure virtual
// void paint(class QPainter *)
func (this *QQuickPaintedItem) Paint(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem5paintEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:110
// index:0
// Public virtual
// bool isTextureProvider()
func (this *QQuickPaintedItem) IsTextureProvider() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem17isTextureProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:111
// index:0
// Public virtual
// QSGTextureProvider * textureProvider()
func (this *QQuickPaintedItem) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QQuickPaintedItem15textureProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:114
// index:0
// Public
// void fillColorChanged()
func (this *QQuickPaintedItem) FillColorChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem16fillColorChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:115
// index:0
// Public
// void contentsSizeChanged()
func (this *QQuickPaintedItem) ContentsSizeChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem19contentsSizeChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:116
// index:0
// Public
// void contentsScaleChanged()
func (this *QQuickPaintedItem) ContentsScaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem20contentsScaleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:117
// index:0
// Public
// void renderTargetChanged()
func (this *QQuickPaintedItem) RenderTargetChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem19renderTargetChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:118
// index:0
// Public
// void textureSizeChanged()
func (this *QQuickPaintedItem) TextureSizeChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem18textureSizeChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:123
// index:0
// Protected virtual
// void releaseResources()
func (this *QQuickPaintedItem) ReleaseResources() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QQuickPaintedItem16releaseResourcesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QQuickPaintedItem__RenderTarget = int

const QQuickPaintedItem__Image QQuickPaintedItem__RenderTarget = 0
const QQuickPaintedItem__FramebufferObject QQuickPaintedItem__RenderTarget = 1
const QQuickPaintedItem__InvertedYFramebufferObject QQuickPaintedItem__RenderTarget = 2

type QQuickPaintedItem__PerformanceHint = int

const QQuickPaintedItem__FastFBOResizing QQuickPaintedItem__PerformanceHint = 1

//  body block end
