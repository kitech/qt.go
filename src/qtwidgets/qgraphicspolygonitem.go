//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QGraphicsPolygonItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsPolygonItem) GetCthis() unsafe.Pointer {
	return this.QAbstractGraphicsShapeItem.GetCthis()
}
func NewQGraphicsPolygonItemFromPointer(cthis unsafe.Pointer) *QGraphicsPolygonItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsPolygonItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:743
// index:0
// Public
// void QGraphicsPolygonItem(class QGraphicsItem *)
func NewQGraphicsPolygonItem(parent unsafe.Pointer) *QGraphicsPolygonItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsPolygonItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:744
// index:1
// Public
// void QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
func NewQGraphicsPolygonItem_1(polygon *qtgui.QPolygonF, parent unsafe.Pointer) *QGraphicsPolygonItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsPolygonItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:746
// index:0
// Public virtual
// void ~QGraphicsPolygonItem()
func DeleteQGraphicsPolygonItem(*QGraphicsPolygonItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:748
// index:0
// Public
// QPolygonF polygon()
func (this *QGraphicsPolygonItem) Polygon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem7polygonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:749
// index:0
// Public
// void setPolygon(const class QPolygonF &)
func (this *QGraphicsPolygonItem) SetPolygon(polygon *qtgui.QPolygonF) {
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:751
// index:0
// Public
// Qt::FillRule fillRule()
func (this *QGraphicsPolygonItem) FillRule() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem8fillRuleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:752
// index:0
// Public
// void setFillRule(Qt::FillRule)
func (this *QGraphicsPolygonItem) SetFillRule(rule int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem11setFillRuleEN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &rule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:754
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsPolygonItem) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:755
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsPolygonItem) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:756
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsPolygonItem) Contains(point *qtcore.QPointF) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:758
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPolygonItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:760
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPolygonItem) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:761
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsPolygonItem) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:764
// index:0
// Public virtual
// int type()
func (this *QGraphicsPolygonItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:767
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsPolygonItem) SupportsExtension(extension int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:768
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsPolygonItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:769
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsPolygonItem) Extension(variant *qtcore.QVariant) interface{} {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
