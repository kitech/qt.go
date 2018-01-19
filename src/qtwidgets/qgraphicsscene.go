//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsscene.h
// #include <qgraphicsscene.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QGraphicsScene struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:98
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsScene) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:124
// index:0
// void QGraphicsScene(class QObject *)
func NewQGraphicsScene(parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsScene{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:125
// index:1
// void QGraphicsScene(const class QRectF &, class QObject *)
func NewQGraphicsScene_1(sceneRect unsafe.Pointer, parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2ERK6QRectFP7QObject", ffiqt.FFI_TYPE_VOID, cthis, sceneRect, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsScene{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:126
// index:2
// void QGraphicsScene(qreal, qreal, qreal, qreal, class QObject *)
func NewQGraphicsScene_2(x float64, y float64, width float64, height float64, parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EddddP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &width, &height, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsScene{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:127
// index:0
// virtual
// void ~QGraphicsScene()
func DeleteQGraphicsScene(*QGraphicsScene) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:129
// index:0
// QRectF sceneRect()
func (this *QGraphicsScene) SceneRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene9sceneRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:130
// index:0
// inline
// qreal width()
func (this *QGraphicsScene) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:131
// index:0
// inline
// qreal height()
func (this *QGraphicsScene) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6heightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:132
// index:0
// void setSceneRect(const class QRectF &)
func (this *QGraphicsScene) SetSceneRect(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:133
// index:1
// inline
// void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, qreal x, qreal y, qreal w, qreal h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// void render(class QPainter *, const class QRectF &, const class QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsScene) Render(painter unsafe.Pointer, target unsafe.Pointer, source unsafe.Pointer, aspectRatioMode int) {
	// 0: (, QPainter * painter, const QRectF & target, const QRectF & source, Qt::AspectRatioMode aspectRatioMode), (painter, target, source, &aspectRatioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.cthis, painter, target, source, &aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:140
// index:0
// QGraphicsScene::ItemIndexMethod itemIndexMethod()
func (this *QGraphicsScene) ItemIndexMethod() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15itemIndexMethodEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:141
// index:0
// void setItemIndexMethod(enum QGraphicsScene::ItemIndexMethod)
func (this *QGraphicsScene) SetItemIndexMethod(method int) {
	// 0: (, QGraphicsScene::ItemIndexMethod method), (&method)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setItemIndexMethodENS_15ItemIndexMethodE", ffiqt.FFI_TYPE_VOID, this.cthis, &method)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:143
// index:0
// bool isSortCacheEnabled()
func (this *QGraphicsScene) IsSortCacheEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene18isSortCacheEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:144
// index:0
// void setSortCacheEnabled(_Bool)
func (this *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene19setSortCacheEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:146
// index:0
// int bspTreeDepth()
func (this *QGraphicsScene) BspTreeDepth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene12bspTreeDepthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:147
// index:0
// void setBspTreeDepth(int)
func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	// 0: (, int depth), (&depth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15setBspTreeDepthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:149
// index:0
// QRectF itemsBoundingRect()
func (this *QGraphicsScene) ItemsBoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene17itemsBoundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:151
// index:0
// QList<QGraphicsItem *> items(Qt::SortOrder)
func (this *QGraphicsScene) Items(order int) {
	// 0: (, Qt::SortOrder order), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsEN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:153
// index:1
// QList<QGraphicsItem *> items(const class QPointF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_1(pos unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 1: (, const QPointF & pos, Qt::ItemSelectionMode mode, Qt::SortOrder order, const QTransform & deviceTransform), (pos, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK7QPointFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, pos, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:154
// index:2
// QList<QGraphicsItem *> items(const class QRectF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_2(rect unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 2: (, const QRectF & rect, Qt::ItemSelectionMode mode, Qt::SortOrder order, const QTransform & deviceTransform), (rect, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK6QRectFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, rect, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:155
// index:3
// QList<QGraphicsItem *> items(const class QPolygonF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_3(polygon unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 3: (, const QPolygonF & polygon, Qt::ItemSelectionMode mode, Qt::SortOrder order, const QTransform & deviceTransform), (polygon, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK9QPolygonFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, polygon, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:156
// index:4
// QList<QGraphicsItem *> items(const class QPainterPath &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_4(path unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 4: (, const QPainterPath & path, Qt::ItemSelectionMode mode, Qt::SortOrder order, const QTransform & deviceTransform), (path, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK12QPainterPathN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, path, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:170
// index:5
// inline
// QList<QGraphicsItem *> items(qreal, qreal, qreal, qreal, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_5(x float64, y float64, w float64, h float64, mode int, order int, deviceTransform unsafe.Pointer) {
	// 5: (, qreal x, qreal y, qreal w, qreal h, Qt::ItemSelectionMode mode, Qt::SortOrder order, const QTransform & deviceTransform), (&x, &y, &w, &h, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsEddddN2Qt17ItemSelectionModeENS0_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:158
// index:0
// QList<QGraphicsItem *> collidingItems(const class QGraphicsItem *, Qt::ItemSelectionMode)
func (this *QGraphicsScene) CollidingItems(item unsafe.Pointer, mode int) {
	// 0: (, const QGraphicsItem * item, Qt::ItemSelectionMode mode), (item, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene14collidingItemsEPK13QGraphicsItemN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, item, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:165
// index:0
// QGraphicsItem * itemAt(const class QPointF &, const class QTransform &)
func (this *QGraphicsScene) ItemAt(pos unsafe.Pointer, deviceTransform unsafe.Pointer) {
	// 0: (, const QPointF & pos, const QTransform & deviceTransform), (pos, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, pos, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:179
// index:1
// inline
// QGraphicsItem * itemAt(qreal, qreal, const class QTransform &)
func (this *QGraphicsScene) ItemAt_1(x float64, y float64, deviceTransform unsafe.Pointer) {
	// 1: (, qreal x, qreal y, const QTransform & deviceTransform), (&x, &y, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtEddRK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:182
// index:0
// QList<QGraphicsItem *> selectedItems()
func (this *QGraphicsScene) SelectedItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectedItemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:183
// index:0
// QPainterPath selectionArea()
func (this *QGraphicsScene) SelectionArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectionAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:184
// index:0
// void setSelectionArea(const class QPainterPath &, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea(path unsafe.Pointer, deviceTransform unsafe.Pointer) {
	// 0: (, const QPainterPath & path, const QTransform & deviceTransform), (path, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, path, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// void setSelectionArea(const class QPainterPath &, Qt::ItemSelectionMode, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea_1(path unsafe.Pointer, mode int, deviceTransform unsafe.Pointer) {
	// 1: (, const QPainterPath & path, Qt::ItemSelectionMode mode, const QTransform & deviceTransform), (path, &mode, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, path, &mode, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// void setSelectionArea(const class QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea_2(path unsafe.Pointer, selectionOperation int, mode int, deviceTransform unsafe.Pointer) {
	// 2: (, const QPainterPath & path, Qt::ItemSelectionOperation selectionOperation, Qt::ItemSelectionMode mode, const QTransform & deviceTransform), (path, &selectionOperation, &mode, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, path, &selectionOperation, &mode, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:190
// index:0
// void destroyItemGroup(class QGraphicsItemGroup *)
func (this *QGraphicsScene) DestroyItemGroup(group unsafe.Pointer) {
	// 0: (, QGraphicsItemGroup * group), (group)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup", ffiqt.FFI_TYPE_VOID, this.cthis, group)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:192
// index:0
// void addItem(class QGraphicsItem *)
func (this *QGraphicsScene) AddItem(item unsafe.Pointer) {
	// 0: (, QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addItemEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// QGraphicsEllipseItem * addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddEllipse(rect unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, const QRectF & rect, const QPen & pen, const QBrush & brush), (rect, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, rect, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// inline
// QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddEllipse_1(x float64, y float64, w float64, h float64, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 1: (, qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush), (&x, &y, &w, &h, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:194
// index:0
// QGraphicsLineItem * addLine(const class QLineF &, const class QPen &)
func (this *QGraphicsScene) AddLine(line unsafe.Pointer, pen unsafe.Pointer) {
	// 0: (, const QLineF & line, const QPen & pen), (line, pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen", ffiqt.FFI_TYPE_VOID, this.cthis, line, pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:204
// index:1
// inline
// QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const class QPen &)
func (this *QGraphicsScene) AddLine_1(x1 float64, y1 float64, x2 float64, y2 float64, pen unsafe.Pointer) {
	// 1: (, qreal x1, qreal y1, qreal x2, qreal y2, const QPen & pen), (&x1, &y1, &x2, &y2, pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineEddddRK4QPen", ffiqt.FFI_TYPE_VOID, this.cthis, &x1, &y1, &x2, &y2, pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// QGraphicsPathItem * addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddPath(path unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, const QPainterPath & path, const QPen & pen, const QBrush & brush), (path, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, path, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:196
// index:0
// QGraphicsPixmapItem * addPixmap(const class QPixmap &)
func (this *QGraphicsScene) AddPixmap(pixmap unsafe.Pointer) {
	// 0: (, const QPixmap & pixmap), (pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9addPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// QGraphicsPolygonItem * addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddPolygon(polygon unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, const QPolygonF & polygon, const QPen & pen, const QBrush & brush), (polygon, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, polygon, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// QGraphicsRectItem * addRect(const class QRectF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddRect(rect unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, const QRectF & rect, const QPen & pen, const QBrush & brush), (rect, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, rect, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// inline
// QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddRect_1(x float64, y float64, w float64, h float64, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 1: (, qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush), (&x, &y, &w, &h, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:199
// index:0
// QGraphicsTextItem * addText(const class QString &, const class QFont &)
func (this *QGraphicsScene) AddText(text unsafe.Pointer, font unsafe.Pointer) {
	// 0: (, const QString & text, const QFont & font), (text, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addTextERK7QStringRK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, text, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:200
// index:0
// QGraphicsSimpleTextItem * addSimpleText(const class QString &, const class QFont &)
func (this *QGraphicsScene) AddSimpleText(text unsafe.Pointer, font unsafe.Pointer) {
	// 0: (, const QString & text, const QFont & font), (text, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, text, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:208
// index:0
// void removeItem(class QGraphicsItem *)
func (this *QGraphicsScene) RemoveItem(item unsafe.Pointer) {
	// 0: (, QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10removeItemEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:210
// index:0
// QGraphicsItem * focusItem()
func (this *QGraphicsScene) FocusItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene9focusItemEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:211
// index:0
// void setFocusItem(class QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) SetFocusItem(item unsafe.Pointer, focusReason int) {
	// 0: (, QGraphicsItem * item, Qt::FocusReason focusReason), (item, &focusReason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setFocusItemEP13QGraphicsItemN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.cthis, item, &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:212
// index:0
// bool hasFocus()
func (this *QGraphicsScene) HasFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene8hasFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:213
// index:0
// void setFocus(Qt::FocusReason)
func (this *QGraphicsScene) SetFocus(focusReason int) {
	// 0: (, Qt::FocusReason focusReason), (&focusReason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.cthis, &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:214
// index:0
// void clearFocus()
func (this *QGraphicsScene) ClearFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10clearFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:216
// index:0
// void setStickyFocus(_Bool)
func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14setStickyFocusEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:217
// index:0
// bool stickyFocus()
func (this *QGraphicsScene) StickyFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene11stickyFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:219
// index:0
// QGraphicsItem * mouseGrabberItem()
func (this *QGraphicsScene) MouseGrabberItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene16mouseGrabberItemEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:221
// index:0
// QBrush backgroundBrush()
func (this *QGraphicsScene) BackgroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15backgroundBrushEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:222
// index:0
// void setBackgroundBrush(const class QBrush &)
func (this *QGraphicsScene) SetBackgroundBrush(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:224
// index:0
// QBrush foregroundBrush()
func (this *QGraphicsScene) ForegroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15foregroundBrushEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:225
// index:0
// void setForegroundBrush(const class QBrush &)
func (this *QGraphicsScene) SetForegroundBrush(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setForegroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:227
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsScene) InputMethodQuery(query int) {
	// 0: (, Qt::InputMethodQuery query), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:229
// index:0
// QList<QGraphicsView *> views()
func (this *QGraphicsScene) Views() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5viewsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:231
// index:0
// inline
// void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	// 0: (, qreal x, qreal y, qreal w, qreal h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6updateEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:257
// index:1
// void update(const class QRectF &)
func (this *QGraphicsScene) Update_1(rect unsafe.Pointer) {
	// 1: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6updateERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:236
// index:0
// QStyle * style()
func (this *QGraphicsScene) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5styleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:237
// index:0
// void setStyle(class QStyle *)
func (this *QGraphicsScene) SetStyle(style unsafe.Pointer) {
	// 0: (, QStyle * style), (style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene8setStyleEP6QStyle", ffiqt.FFI_TYPE_VOID, this.cthis, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:239
// index:0
// QFont font()
func (this *QGraphicsScene) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:240
// index:0
// void setFont(const class QFont &)
func (this *QGraphicsScene) SetFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:242
// index:0
// QPalette palette()
func (this *QGraphicsScene) Palette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene7paletteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:243
// index:0
// void setPalette(const class QPalette &)
func (this *QGraphicsScene) SetPalette(palette unsafe.Pointer) {
	// 0: (, const QPalette & palette), (palette)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10setPaletteERK8QPalette", ffiqt.FFI_TYPE_VOID, this.cthis, palette)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:245
// index:0
// bool isActive()
func (this *QGraphicsScene) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene8isActiveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:246
// index:0
// QGraphicsItem * activePanel()
func (this *QGraphicsScene) ActivePanel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene11activePanelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:247
// index:0
// void setActivePanel(class QGraphicsItem *)
func (this *QGraphicsScene) SetActivePanel(item unsafe.Pointer) {
	// 0: (, QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:248
// index:0
// QGraphicsWidget * activeWindow()
func (this *QGraphicsScene) ActiveWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene12activeWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:249
// index:0
// void setActiveWindow(class QGraphicsWidget *)
func (this *QGraphicsScene) SetActiveWindow(widget unsafe.Pointer) {
	// 0: (, QGraphicsWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:251
// index:0
// bool sendEvent(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsScene) SendEvent(item unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, QGraphicsItem * item, QEvent * event), (item, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, item, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:253
// index:0
// qreal minimumRenderSize()
func (this *QGraphicsScene) MinimumRenderSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene17minimumRenderSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:254
// index:0
// void setMinimumRenderSize(qreal)
func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	// 0: (, qreal minSize), (&minSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene20setMinimumRenderSizeEd", ffiqt.FFI_TYPE_VOID, this.cthis, &minSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:259
// index:0
// void advance()
func (this *QGraphicsScene) Advance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7advanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:260
// index:0
// void clearSelection()
func (this *QGraphicsScene) ClearSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14clearSelectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:261
// index:0
// void clear()
func (this *QGraphicsScene) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:299
// index:0
// void sceneRectChanged(const class QRectF &)
func (this *QGraphicsScene) SceneRectChanged(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16sceneRectChangedERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:300
// index:0
// void selectionChanged()
func (this *QGraphicsScene) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:301
// index:0
// void focusItemChanged(class QGraphicsItem *, class QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) FocusItemChanged(newFocus unsafe.Pointer, oldFocus unsafe.Pointer, reason int) {
	// 0: (, QGraphicsItem * newFocus, QGraphicsItem * oldFocus, Qt::FocusReason reason), (newFocus, oldFocus, &reason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16focusItemChangedEP13QGraphicsItemS1_N2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.cthis, newFocus, oldFocus, &reason)
	gopp.ErrPrint(err, rv)
}

//  body block end
