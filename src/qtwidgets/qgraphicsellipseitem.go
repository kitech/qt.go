//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QGraphicsEllipseItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsEllipseItem) GetCthis() unsafe.Pointer {
	return this.QAbstractGraphicsShapeItem.GetCthis()
}
func NewQGraphicsEllipseItemFromPointer(cthis unsafe.Pointer) *QGraphicsEllipseItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsEllipseItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:699
// index:0
// Public
// void QGraphicsEllipseItem(class QGraphicsItem *)
func NewQGraphicsEllipseItem(parent unsafe.Pointer) *QGraphicsEllipseItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEllipseItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:700
// index:1
// Public
// void QGraphicsEllipseItem(const class QRectF &, class QGraphicsItem *)
func NewQGraphicsEllipseItem_1(rect *qtcore.QRectF, parent unsafe.Pointer) *QGraphicsEllipseItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEllipseItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:701
// index:2
// Public
// void QGraphicsEllipseItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
func NewQGraphicsEllipseItem_2(x float64, y float64, w float64, h float64, parent unsafe.Pointer) *QGraphicsEllipseItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &w, &h, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEllipseItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:702
// index:0
// Public virtual
// void ~QGraphicsEllipseItem()
func DeleteQGraphicsEllipseItem(*QGraphicsEllipseItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:704
// index:0
// Public
// QRectF rect()
func (this *QGraphicsEllipseItem) Rect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:705
// index:0
// Public
// void setRect(const class QRectF &)
func (this *QGraphicsEllipseItem) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:706
// index:1
// Public inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsEllipseItem) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:708
// index:0
// Public
// int startAngle()
func (this *QGraphicsEllipseItem) StartAngle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem10startAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:709
// index:0
// Public
// void setStartAngle(int)
func (this *QGraphicsEllipseItem) SetStartAngle(angle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem13setStartAngleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:711
// index:0
// Public
// int spanAngle()
func (this *QGraphicsEllipseItem) SpanAngle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem9spanAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:712
// index:0
// Public
// void setSpanAngle(int)
func (this *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem12setSpanAngleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:714
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsEllipseItem) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:715
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsEllipseItem) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:716
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsEllipseItem) Contains(point *qtcore.QPointF) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:718
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsEllipseItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:720
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsEllipseItem) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:721
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsEllipseItem) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:724
// index:0
// Public virtual
// int type()
func (this *QGraphicsEllipseItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:727
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsEllipseItem) SupportsExtension(extension int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:728
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsEllipseItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:729
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsEllipseItem) Extension(variant *qtcore.QVariant) interface{} {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
