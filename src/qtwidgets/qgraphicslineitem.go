//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:780
// index:0
// void QGraphicsLineItem(class QGraphicsItem *)
func NewQGraphicsLineItem(parent unsafe.Pointer) *QGraphicsLineItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsLineItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:781
// index:1
// void QGraphicsLineItem(const class QLineF &, class QGraphicsItem *)
func NewQGraphicsLineItem_1(line unsafe.Pointer, parent unsafe.Pointer) *QGraphicsLineItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemC2ERK6QLineFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, line, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsLineItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:782
// index:2
// void QGraphicsLineItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
func NewQGraphicsLineItem_2(x1 float64, y1 float64, x2 float64, y2 float64, parent unsafe.Pointer) *QGraphicsLineItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, &x1, &y1, &x2, &y2, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsLineItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:783
// index:0
// virtual
// void ~QGraphicsLineItem()
func DeleteQGraphicsLineItem(*QGraphicsLineItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:785
// index:0
// QPen pen()
func (this *QGraphicsLineItem) Pen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem3penEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:786
// index:0
// void setPen(const class QPen &)
func (this *QGraphicsLineItem) SetPen(pen unsafe.Pointer) {
	// 0: (, const QPen & pen), (pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem6setPenERK4QPen", ffiqt.FFI_TYPE_VOID, this.cthis, pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:788
// index:0
// QLineF line()
func (this *QGraphicsLineItem) Line() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem4lineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:789
// index:0
// void setLine(const class QLineF &)
func (this *QGraphicsLineItem) SetLine(line unsafe.Pointer) {
	// 0: (, const QLineF & line), (line)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem7setLineERK6QLineF", ffiqt.FFI_TYPE_VOID, this.cthis, line)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:790
// index:1
// inline
// void setLine(qreal, qreal, qreal, qreal)
func (this *QGraphicsLineItem) SetLine_1(x1 float64, y1 float64, x2 float64, y2 float64) {
	// 1: (, qreal x1, qreal y1, qreal x2, qreal y2), (&x1, &y1, &x2, &y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem7setLineEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:793
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsLineItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:794
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsLineItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:795
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsLineItem) Contains(point unsafe.Pointer) {
	// 0: (, const QPointF & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:797
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsLineItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsLineItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:799
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsLineItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:800
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsLineItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:803
// index:0
// virtual
// int type()
func (this *QGraphicsLineItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsLineItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
