//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:743
// index:0
// void QGraphicsPolygonItem(class QGraphicsItem *)
func NewQGraphicsPolygonItem(parent unsafe.Pointer) *QGraphicsPolygonItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsPolygonItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:744
// index:1
// void QGraphicsPolygonItem(const class QPolygonF &, class QGraphicsItem *)
func NewQGraphicsPolygonItem_1(polygon unsafe.Pointer, parent unsafe.Pointer) *QGraphicsPolygonItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, polygon, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsPolygonItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:746
// index:0
// virtual
// void ~QGraphicsPolygonItem()
func DeleteQGraphicsPolygonItem(*QGraphicsPolygonItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:748
// index:0
// QPolygonF polygon()
func (this *QGraphicsPolygonItem) Polygon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem7polygonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:749
// index:0
// void setPolygon(const class QPolygonF &)
func (this *QGraphicsPolygonItem) SetPolygon(polygon unsafe.Pointer) {
	// 0: (, const QPolygonF & polygon), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.cthis, polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:751
// index:0
// Qt::FillRule fillRule()
func (this *QGraphicsPolygonItem) FillRule() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem8fillRuleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:752
// index:0
// void setFillRule(Qt::FillRule)
func (this *QGraphicsPolygonItem) SetFillRule(rule int) {
	// 0: (, Qt::FillRule rule), (&rule)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem11setFillRuleEN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.cthis, &rule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:754
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsPolygonItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:755
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsPolygonItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:756
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsPolygonItem) Contains(point unsafe.Pointer) {
	// 0: (, const QPointF & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:758
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPolygonItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:760
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPolygonItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:761
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsPolygonItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:764
// index:0
// virtual
// int type()
func (this *QGraphicsPolygonItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
