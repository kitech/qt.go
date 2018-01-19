//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:825
// index:0
// void QGraphicsPixmapItem(class QGraphicsItem *)
func NewQGraphicsPixmapItem(parent unsafe.Pointer) *QGraphicsPixmapItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsPixmapItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:826
// index:1
// void QGraphicsPixmapItem(const class QPixmap &, class QGraphicsItem *)
func NewQGraphicsPixmapItem_1(pixmap unsafe.Pointer, parent unsafe.Pointer) *QGraphicsPixmapItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItemC2ERK7QPixmapP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, pixmap, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsPixmapItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:827
// index:0
// virtual
// void ~QGraphicsPixmapItem()
func DeleteQGraphicsPixmapItem(*QGraphicsPixmapItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:829
// index:0
// QPixmap pixmap()
func (this *QGraphicsPixmapItem) Pixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem6pixmapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:830
// index:0
// void setPixmap(const class QPixmap &)
func (this *QGraphicsPixmapItem) SetPixmap(pixmap unsafe.Pointer) {
	// 0: (, const QPixmap & pixmap), (pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:832
// index:0
// Qt::TransformationMode transformationMode()
func (this *QGraphicsPixmapItem) TransformationMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem18transformationModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:833
// index:0
// void setTransformationMode(Qt::TransformationMode)
func (this *QGraphicsPixmapItem) SetTransformationMode(mode int) {
	// 0: (, Qt::TransformationMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem21setTransformationModeEN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:835
// index:0
// QPointF offset()
func (this *QGraphicsPixmapItem) Offset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem6offsetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:836
// index:0
// void setOffset(const class QPointF &)
func (this *QGraphicsPixmapItem) SetOffset(offset unsafe.Pointer) {
	// 0: (, const QPointF & offset), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem9setOffsetERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:837
// index:1
// inline
// void setOffset(qreal, qreal)
func (this *QGraphicsPixmapItem) SetOffset_1(x float64, y float64) {
	// 1: (, qreal x, qreal y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem9setOffsetEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:839
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsPixmapItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:840
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsPixmapItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:841
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsPixmapItem) Contains(point unsafe.Pointer) {
	// 0: (, const QPointF & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:843
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPixmapItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:845
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPixmapItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:846
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsPixmapItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:849
// index:0
// virtual
// int type()
func (this *QGraphicsPixmapItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:851
// index:0
// QGraphicsPixmapItem::ShapeMode shapeMode()
func (this *QGraphicsPixmapItem) ShapeMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsPixmapItem9shapeModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:852
// index:0
// void setShapeMode(enum QGraphicsPixmapItem::ShapeMode)
func (this *QGraphicsPixmapItem) SetShapeMode(mode int) {
	// 0: (, QGraphicsPixmapItem::ShapeMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsPixmapItem12setShapeModeENS_9ShapeModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

//  body block end
