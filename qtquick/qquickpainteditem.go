package qtquick

// /usr/include/qt/QtQuick/qquickpainteditem.h
// #include <qquickpainteditem.h>
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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// void releaseResources()
func (this *QQuickPaintedItem) InheritReleaseResources(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "releaseResources", f)
}

type QQuickPaintedItem struct {
	*QQuickItem
}
type QQuickPaintedItem_ITF interface {
	QQuickItem_ITF
	QQuickPaintedItem_PTR() *QQuickPaintedItem
}

func (ptr *QQuickPaintedItem) QQuickPaintedItem_PTR() *QQuickPaintedItem { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickPaintedItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickPaintedItem(QQuickItem *)
func NewQQuickPaintedItem(parent QQuickItem_ITF /*777 QQuickItem **/) *QQuickPaintedItem {
	var convArg0 = parent.QQuickItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItemC2EP10QQuickItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickPaintedItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickPaintedItem()
func DeleteQQuickPaintedItem(this *QQuickPaintedItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRect &)
func (this *QQuickPaintedItem) Update(rect qtcore.QRect_ITF) {
	var convArg0 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem6updateERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool opaquePainting()
func (this *QQuickPaintedItem) OpaquePainting() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem14opaquePaintingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaquePainting(_Bool)
func (this *QQuickPaintedItem) SetOpaquePainting(opaque bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem17setOpaquePaintingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool antialiasing()
func (this *QQuickPaintedItem) Antialiasing() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem12antialiasingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAntialiasing(_Bool)
func (this *QQuickPaintedItem) SetAntialiasing(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem15setAntialiasingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mipmap()
func (this *QQuickPaintedItem) Mipmap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem6mipmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMipmap(_Bool)
func (this *QQuickPaintedItem) SetMipmap(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem9setMipmapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickPaintedItem::PerformanceHints performanceHints()
func (this *QQuickPaintedItem) PerformanceHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem16performanceHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPerformanceHint(enum QQuickPaintedItem::PerformanceHint, _Bool)
func (this *QQuickPaintedItem) SetPerformanceHint(hint int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem18setPerformanceHintENS_15PerformanceHintEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPerformanceHints(QQuickPaintedItem::PerformanceHints)
func (this *QQuickPaintedItem) SetPerformanceHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem19setPerformanceHintsE6QFlagsINS_15PerformanceHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:90
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF contentsBoundingRect()
func (this *QQuickPaintedItem) ContentsBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem20contentsBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize contentsSize()
func (this *QQuickPaintedItem) ContentsSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem12contentsSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsSize(const QSize &)
func (this *QQuickPaintedItem) SetContentsSize(arg0 qtcore.QSize_ITF) {
	var convArg0 = arg0.QSize_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem15setContentsSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetContentsSize()
func (this *QQuickPaintedItem) ResetContentsSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem17resetContentsSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal contentsScale()
func (this *QQuickPaintedItem) ContentsScale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem13contentsScaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsScale(qreal)
func (this *QQuickPaintedItem) SetContentsScale(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem16setContentsScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize textureSize()
func (this *QQuickPaintedItem) TextureSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem11textureSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextureSize(const QSize &)
func (this *QQuickPaintedItem) SetTextureSize(size qtcore.QSize_ITF) {
	var convArg0 = size.QSize_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem14setTextureSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:102
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor fillColor()
func (this *QQuickPaintedItem) FillColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem9fillColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFillColor(const QColor &)
func (this *QQuickPaintedItem) SetFillColor(arg0 qtgui.QColor_ITF) {
	var convArg0 = arg0.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem12setFillColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickPaintedItem::RenderTarget renderTarget()
func (this *QQuickPaintedItem) RenderTarget() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem12renderTargetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderTarget(enum QQuickPaintedItem::RenderTarget)
func (this *QQuickPaintedItem) SetRenderTarget(target int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem15setRenderTargetENS_12RenderTargetE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), target)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:108
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *)
func (this *QQuickPaintedItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem5paintEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:110
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isTextureProvider()
func (this *QQuickPaintedItem) IsTextureProvider() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem17isTextureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGTextureProvider * textureProvider()
func (this *QQuickPaintedItem) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QQuickPaintedItem15textureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fillColorChanged()
func (this *QQuickPaintedItem) FillColorChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem16fillColorChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsSizeChanged()
func (this *QQuickPaintedItem) ContentsSizeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem19contentsSizeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsScaleChanged()
func (this *QQuickPaintedItem) ContentsScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem20contentsScaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void renderTargetChanged()
func (this *QQuickPaintedItem) RenderTargetChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem19renderTargetChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textureSizeChanged()
func (this *QQuickPaintedItem) TextureSizeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem18textureSizeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void releaseResources()
func (this *QQuickPaintedItem) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QQuickPaintedItem16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QQuickPaintedItem__RenderTarget = int

const QQuickPaintedItem__Image QQuickPaintedItem__RenderTarget = 0
const QQuickPaintedItem__FramebufferObject QQuickPaintedItem__RenderTarget = 1
const QQuickPaintedItem__InvertedYFramebufferObject QQuickPaintedItem__RenderTarget = 2

type QQuickPaintedItem__PerformanceHint = int

const QQuickPaintedItem__FastFBOResizing QQuickPaintedItem__PerformanceHint = 1

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
