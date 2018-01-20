//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QGraphicsRectItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsRectItem) GetCthis() unsafe.Pointer {
	return this.QAbstractGraphicsShapeItem.GetCthis()
}
func NewQGraphicsRectItemFromPointer(cthis unsafe.Pointer) *QGraphicsRectItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsRectItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:661
// index:0
// Public
// void QGraphicsRectItem(class QGraphicsItem *)
func NewQGraphicsRectItem(parent unsafe.Pointer) *QGraphicsRectItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:662
// index:1
// Public
// void QGraphicsRectItem(const class QRectF &, class QGraphicsItem *)
func NewQGraphicsRectItem_1(rect *qtcore.QRectF, parent unsafe.Pointer) *QGraphicsRectItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:663
// index:2
// Public
// void QGraphicsRectItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
func NewQGraphicsRectItem_2(x float64, y float64, w float64, h float64, parent unsafe.Pointer) *QGraphicsRectItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &w, &h, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:664
// index:0
// Public virtual
// void ~QGraphicsRectItem()
func DeleteQGraphicsRectItem(*QGraphicsRectItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:666
// index:0
// Public
// QRectF rect()
func (this *QGraphicsRectItem) Rect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:667
// index:0
// Public
// void setRect(const class QRectF &)
func (this *QGraphicsRectItem) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:668
// index:1
// Public inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsRectItem) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:670
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsRectItem) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:671
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsRectItem) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:672
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsRectItem) Contains(point *qtcore.QPointF) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:674
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsRectItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:676
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsRectItem) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:677
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsRectItem) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:680
// index:0
// Public virtual
// int type()
func (this *QGraphicsRectItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:683
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsRectItem) SupportsExtension(extension int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:684
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsRectItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:685
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsRectItem) Extension(variant *qtcore.QVariant) interface{} {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
