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
type QGraphicsLineItem struct {
	*QGraphicsItem
}

func (this *QGraphicsLineItem) GetCthis() unsafe.Pointer {
	return this.QGraphicsItem.GetCthis()
}
func NewQGraphicsLineItemFromPointer(cthis unsafe.Pointer) *QGraphicsLineItem {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QGraphicsLineItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:780
// index:0
// Public
// void QGraphicsLineItem(class QGraphicsItem *)
func NewQGraphicsLineItem(parent unsafe.Pointer) *QGraphicsLineItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLineItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:781
// index:1
// Public
// void QGraphicsLineItem(const class QLineF &, class QGraphicsItem *)
func NewQGraphicsLineItem_1(line *qtcore.QLineF, parent unsafe.Pointer) *QGraphicsLineItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemC2ERK6QLineFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLineItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:782
// index:2
// Public
// void QGraphicsLineItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
func NewQGraphicsLineItem_2(x1 float64, y1 float64, x2 float64, y2 float64, parent unsafe.Pointer) *QGraphicsLineItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, &x1, &y1, &x2, &y2, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLineItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:783
// index:0
// Public virtual
// void ~QGraphicsLineItem()
func DeleteQGraphicsLineItem(*QGraphicsLineItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:785
// index:0
// Public
// QPen pen()
func (this *QGraphicsLineItem) Pen() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem3penEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:786
// index:0
// Public
// void setPen(const class QPen &)
func (this *QGraphicsLineItem) SetPen(pen *qtgui.QPen) {
	var convArg0 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem6setPenERK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:788
// index:0
// Public
// QLineF line()
func (this *QGraphicsLineItem) Line() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem4lineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:789
// index:0
// Public
// void setLine(const class QLineF &)
func (this *QGraphicsLineItem) SetLine(line *qtcore.QLineF) {
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem7setLineERK6QLineF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:790
// index:1
// Public inline
// void setLine(qreal, qreal, qreal, qreal)
func (this *QGraphicsLineItem) SetLine_1(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem7setLineEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:793
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsLineItem) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:794
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsLineItem) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:795
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsLineItem) Contains(point *qtcore.QPointF) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:797
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsLineItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:799
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsLineItem) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:800
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsLineItem) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:803
// index:0
// Public virtual
// int type()
func (this *QGraphicsLineItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:806
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsLineItem) SupportsExtension(extension int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:807
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsLineItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:808
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsLineItem) Extension(variant *qtcore.QVariant) interface{} {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
