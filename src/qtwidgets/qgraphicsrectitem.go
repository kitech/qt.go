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

// /usr/include/qt/QtWidgets/qgraphicsitem.h:661
// index:0
// void QGraphicsRectItem(class QGraphicsItem *)
func NewQGraphicsRectItem(parent unsafe.Pointer) *QGraphicsRectItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(cthis)
	return gothis
}
func NewQGraphicsRectItemFromPointer(cthis unsafe.Pointer) *QGraphicsRectItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsRectItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:662
// index:1
// void QGraphicsRectItem(const class QRectF &, class QGraphicsItem *)
func NewQGraphicsRectItem_1(rect unsafe.Pointer, parent unsafe.Pointer) *QGraphicsRectItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, rect, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:663
// index:2
// void QGraphicsRectItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
func NewQGraphicsRectItem_2(x float64, y float64, w float64, h float64, parent unsafe.Pointer) *QGraphicsRectItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &w, &h, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:664
// index:0
// virtual
// void ~QGraphicsRectItem()
func DeleteQGraphicsRectItem(*QGraphicsRectItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:666
// index:0
// QRectF rect()
func (this *QGraphicsRectItem) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4rectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:667
// index:0
// void setRect(const class QRectF &)
func (this *QGraphicsRectItem) SetRect(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:668
// index:1
// inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsRectItem) SetRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:670
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsRectItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:671
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsRectItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:672
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsRectItem) Contains(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:674
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsRectItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionGraphicsItem *, widget QWidget *), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:676
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsRectItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:677
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsRectItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:680
// index:0
// virtual
// int type()
func (this *QGraphicsRectItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:683
// index:0
// virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsRectItem) SupportsExtension(extension int) {
	// 0: (, extension QGraphicsItem::Extension), (&extension)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:684
// index:0
// virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsRectItem) SetExtension(extension int, variant unsafe.Pointer) {
	// 0: (, extension QGraphicsItem::Extension, variant const QVariant &), (&extension, variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extension, variant)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:685
// index:0
// virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsRectItem) Extension(variant unsafe.Pointer) {
	// 0: (, variant const QVariant &), (variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRectItem9extensionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), variant)
	gopp.ErrPrint(err, rv)
}

//  body block end
