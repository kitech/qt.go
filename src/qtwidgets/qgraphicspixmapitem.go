package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
}

//  ext block end

//  body block begin
type QGraphicsPixmapItem struct {
	*QGraphicsItem
}

func (this *QGraphicsPixmapItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsItem.GetCthis()
	}
}
func NewQGraphicsPixmapItemFromPointer(cthis unsafe.Pointer) *QGraphicsPixmapItem {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QGraphicsPixmapItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:825
// index:0
// Public
// void QGraphicsPixmapItem(class QGraphicsItem *)
func NewQGraphicsPixmapItem(parent *QGraphicsItem /*444 QGraphicsItem **/) *QGraphicsPixmapItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsPixmapItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:826
// index:1
// Public
// void QGraphicsPixmapItem(const class QPixmap &, class QGraphicsItem *)
func NewQGraphicsPixmapItem_1(pixmap *qtgui.QPixmap, parent *QGraphicsItem /*444 QGraphicsItem **/) *QGraphicsPixmapItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = pixmap.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItemC2ERK7QPixmapP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsPixmapItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:827
// index:0
// Public virtual
// void ~QGraphicsPixmapItem()
func DeleteQGraphicsPixmapItem(*QGraphicsPixmapItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:829
// index:0
// Public
// QPixmap pixmap()
func (this *QGraphicsPixmapItem) Pixmap() *qtgui.QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem6pixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:830
// index:0
// Public
// void setPixmap(const class QPixmap &)
func (this *QGraphicsPixmapItem) SetPixmap(pixmap *qtgui.QPixmap) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:832
// index:0
// Public
// Qt::TransformationMode transformationMode()
func (this *QGraphicsPixmapItem) TransformationMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem18transformationModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:833
// index:0
// Public
// void setTransformationMode(Qt::TransformationMode)
func (this *QGraphicsPixmapItem) SetTransformationMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem21setTransformationModeEN2Qt18TransformationModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:835
// index:0
// Public
// QPointF offset()
func (this *QGraphicsPixmapItem) Offset() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem6offsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:836
// index:0
// Public
// void setOffset(const class QPointF &)
func (this *QGraphicsPixmapItem) SetOffset(offset *qtcore.QPointF) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem9setOffsetERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:837
// index:1
// Public inline
// void setOffset(qreal, qreal)
func (this *QGraphicsPixmapItem) SetOffset_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem9setOffsetEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:839
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsPixmapItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:840
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsPixmapItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:841
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsPixmapItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:843
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPixmapItem) Paint(painter *qtgui.QPainter /*444 QPainter **/, option *QStyleOptionGraphicsItem /*444 const QStyleOptionGraphicsItem **/, widget *QWidget /*444 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:845
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPixmapItem) IsObscuredBy(item *QGraphicsItem /*444 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:846
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsPixmapItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:849
// index:0
// Public virtual
// int type()
func (this *QGraphicsPixmapItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:851
// index:0
// Public
// QGraphicsPixmapItem::ShapeMode shapeMode()
func (this *QGraphicsPixmapItem) ShapeMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem9shapeModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:852
// index:0
// Public
// void setShapeMode(enum QGraphicsPixmapItem::ShapeMode)
func (this *QGraphicsPixmapItem) SetShapeMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem12setShapeModeENS_9ShapeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:855
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsPixmapItem) SupportsExtension(extension int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:856
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsPixmapItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:857
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsPixmapItem) Extension(variant *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
