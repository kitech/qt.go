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
type QGraphicsEllipseItem struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:699
// index:0
// void QGraphicsEllipseItem(class QGraphicsItem *)
func NewQGraphicsEllipseItem(parent unsafe.Pointer) *QGraphicsEllipseItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsEllipseItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:700
// index:1
// void QGraphicsEllipseItem(const class QRectF &, class QGraphicsItem *)
func NewQGraphicsEllipseItem_1(rect unsafe.Pointer, parent unsafe.Pointer) *QGraphicsEllipseItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, rect, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsEllipseItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:701
// index:2
// void QGraphicsEllipseItem(qreal, qreal, qreal, qreal, class QGraphicsItem *)
func NewQGraphicsEllipseItem_2(x float64, y float64, w float64, h float64, parent unsafe.Pointer) *QGraphicsEllipseItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &w, &h, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsEllipseItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:702
// index:0
// virtual
// void ~QGraphicsEllipseItem()
func DeleteQGraphicsEllipseItem(*QGraphicsEllipseItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:704
// index:0
// QRectF rect()
func (this *QGraphicsEllipseItem) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:705
// index:0
// void setRect(const class QRectF &)
func (this *QGraphicsEllipseItem) SetRect(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem7setRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:706
// index:1
// inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsEllipseItem) SetRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, qreal x, qreal y, qreal w, qreal h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem7setRectEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:708
// index:0
// int startAngle()
func (this *QGraphicsEllipseItem) StartAngle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem10startAngleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:709
// index:0
// void setStartAngle(int)
func (this *QGraphicsEllipseItem) SetStartAngle(angle int) {
	// 0: (, int angle), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem13setStartAngleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:711
// index:0
// int spanAngle()
func (this *QGraphicsEllipseItem) SpanAngle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem9spanAngleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:712
// index:0
// void setSpanAngle(int)
func (this *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	// 0: (, int angle), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem12setSpanAngleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:714
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsEllipseItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:715
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsEllipseItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:716
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsEllipseItem) Contains(point unsafe.Pointer) {
	// 0: (, const QPointF & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:718
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsEllipseItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:720
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsEllipseItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:721
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsEllipseItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:724
// index:0
// virtual
// int type()
func (this *QGraphicsEllipseItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
