//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsscene.h
// #include <qgraphicsscene.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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
	*qtcore.QObject
}

func (this *QGraphicsScene) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:98
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsScene) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:124
// index:0
// void QGraphicsScene(class QObject *)
func NewQGraphicsScene(parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneFromPointer(cthis unsafe.Pointer) *QGraphicsScene {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsScene{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:125
// index:1
// void QGraphicsScene(const class QRectF &, class QObject *)
func NewQGraphicsScene_1(sceneRect unsafe.Pointer, parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2ERK6QRectFP7QObject", ffiqt.FFI_TYPE_VOID, cthis, sceneRect, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:126
// index:2
// void QGraphicsScene(qreal, qreal, qreal, qreal, class QObject *)
func NewQGraphicsScene_2(x float64, y float64, width float64, height float64, parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EddddP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &width, &height, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(cthis)
	return gothis
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene9sceneRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:130
// index:0
// inline
// qreal width()
func (this *QGraphicsScene) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:131
// index:0
// inline
// qreal height()
func (this *QGraphicsScene) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:132
// index:0
// void setSceneRect(const class QRectF &)
func (this *QGraphicsScene) SetSceneRect(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:133
// index:1
// inline
// void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// void render(class QPainter *, const class QRectF &, const class QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsScene) Render(painter unsafe.Pointer, target unsafe.Pointer, source unsafe.Pointer, aspectRatioMode int) {
	// 0: (, painter QPainter *, target const QRectF &, source const QRectF &, aspectRatioMode Qt::AspectRatioMode), (painter, target, source, &aspectRatioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, target, source, &aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:140
// index:0
// QGraphicsScene::ItemIndexMethod itemIndexMethod()
func (this *QGraphicsScene) ItemIndexMethod() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15itemIndexMethodEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:141
// index:0
// void setItemIndexMethod(enum QGraphicsScene::ItemIndexMethod)
func (this *QGraphicsScene) SetItemIndexMethod(method int) {
	// 0: (, method QGraphicsScene::ItemIndexMethod), (&method)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setItemIndexMethodENS_15ItemIndexMethodE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &method)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:143
// index:0
// bool isSortCacheEnabled()
func (this *QGraphicsScene) IsSortCacheEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene18isSortCacheEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:144
// index:0
// void setSortCacheEnabled(_Bool)
func (this *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene19setSortCacheEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:146
// index:0
// int bspTreeDepth()
func (this *QGraphicsScene) BspTreeDepth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene12bspTreeDepthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:147
// index:0
// void setBspTreeDepth(int)
func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	// 0: (, depth int), (&depth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15setBspTreeDepthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:149
// index:0
// QRectF itemsBoundingRect()
func (this *QGraphicsScene) ItemsBoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene17itemsBoundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:151
// index:0
// QList<QGraphicsItem *> items(Qt::SortOrder)
func (this *QGraphicsScene) Items(order int) {
	// 0: (, order Qt::SortOrder), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsEN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:153
// index:1
// QList<QGraphicsItem *> items(const class QPointF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_1(pos unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 1: (, pos const QPointF &, mode Qt::ItemSelectionMode, order Qt::SortOrder, deviceTransform const QTransform &), (pos, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK7QPointFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:154
// index:2
// QList<QGraphicsItem *> items(const class QRectF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_2(rect unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 2: (, rect const QRectF &, mode Qt::ItemSelectionMode, order Qt::SortOrder, deviceTransform const QTransform &), (rect, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK6QRectFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:155
// index:3
// QList<QGraphicsItem *> items(const class QPolygonF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_3(polygon unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 3: (, polygon const QPolygonF &, mode Qt::ItemSelectionMode, order Qt::SortOrder, deviceTransform const QTransform &), (polygon, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK9QPolygonFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:156
// index:4
// QList<QGraphicsItem *> items(const class QPainterPath &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_4(path unsafe.Pointer, mode int, order int, deviceTransform unsafe.Pointer) {
	// 4: (, path const QPainterPath &, mode Qt::ItemSelectionMode, order Qt::SortOrder, deviceTransform const QTransform &), (path, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK12QPainterPathN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:170
// index:5
// inline
// QList<QGraphicsItem *> items(qreal, qreal, qreal, qreal, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_5(x float64, y float64, w float64, h float64, mode int, order int, deviceTransform unsafe.Pointer) {
	// 5: (, x qreal, y qreal, w qreal, h qreal, mode Qt::ItemSelectionMode, order Qt::SortOrder, deviceTransform const QTransform &), (&x, &y, &w, &h, &mode, &order, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsEddddN2Qt17ItemSelectionModeENS0_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &mode, &order, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:158
// index:0
// QList<QGraphicsItem *> collidingItems(const class QGraphicsItem *, Qt::ItemSelectionMode)
func (this *QGraphicsScene) CollidingItems(item unsafe.Pointer, mode int) {
	// 0: (, item const QGraphicsItem *, mode Qt::ItemSelectionMode), (item, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene14collidingItemsEPK13QGraphicsItemN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:165
// index:0
// QGraphicsItem * itemAt(const class QPointF &, const class QTransform &)
func (this *QGraphicsScene) ItemAt(pos unsafe.Pointer, deviceTransform unsafe.Pointer) {
	// 0: (, pos const QPointF &, deviceTransform const QTransform &), (pos, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:179
// index:1
// inline
// QGraphicsItem * itemAt(qreal, qreal, const class QTransform &)
func (this *QGraphicsScene) ItemAt_1(x float64, y float64, deviceTransform unsafe.Pointer) {
	// 1: (, x qreal, y qreal, deviceTransform const QTransform &), (&x, &y, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtEddRK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:182
// index:0
// QList<QGraphicsItem *> selectedItems()
func (this *QGraphicsScene) SelectedItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectedItemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:183
// index:0
// QPainterPath selectionArea()
func (this *QGraphicsScene) SelectionArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectionAreaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:184
// index:0
// void setSelectionArea(const class QPainterPath &, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea(path unsafe.Pointer, deviceTransform unsafe.Pointer) {
	// 0: (, path const QPainterPath &, deviceTransform const QTransform &), (path, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// void setSelectionArea(const class QPainterPath &, Qt::ItemSelectionMode, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea_1(path unsafe.Pointer, mode int, deviceTransform unsafe.Pointer) {
	// 1: (, path const QPainterPath &, mode Qt::ItemSelectionMode, deviceTransform const QTransform &), (path, &mode, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &mode, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// void setSelectionArea(const class QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea_2(path unsafe.Pointer, selectionOperation int, mode int, deviceTransform unsafe.Pointer) {
	// 2: (, path const QPainterPath &, selectionOperation Qt::ItemSelectionOperation, mode Qt::ItemSelectionMode, deviceTransform const QTransform &), (path, &selectionOperation, &mode, deviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &selectionOperation, &mode, deviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:190
// index:0
// void destroyItemGroup(class QGraphicsItemGroup *)
func (this *QGraphicsScene) DestroyItemGroup(group unsafe.Pointer) {
	// 0: (, group QGraphicsItemGroup *), (group)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup", ffiqt.FFI_TYPE_VOID, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:192
// index:0
// void addItem(class QGraphicsItem *)
func (this *QGraphicsScene) AddItem(item unsafe.Pointer) {
	// 0: (, item QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addItemEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// QGraphicsEllipseItem * addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddEllipse(rect unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, rect const QRectF &, pen const QPen &, brush const QBrush &), (rect, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// inline
// QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddEllipse_1(x float64, y float64, w float64, h float64, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, pen const QPen &, brush const QBrush &), (&x, &y, &w, &h, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:194
// index:0
// QGraphicsLineItem * addLine(const class QLineF &, const class QPen &)
func (this *QGraphicsScene) AddLine(line unsafe.Pointer, pen unsafe.Pointer) {
	// 0: (, line const QLineF &, pen const QPen &), (line, pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), line, pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:204
// index:1
// inline
// QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const class QPen &)
func (this *QGraphicsScene) AddLine_1(x1 float64, y1 float64, x2 float64, y2 float64, pen unsafe.Pointer) {
	// 1: (, x1 qreal, y1 qreal, x2 qreal, y2 qreal, pen const QPen &), (&x1, &y1, &x2, &y2, pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineEddddRK4QPen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x1, &y1, &x2, &y2, pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// QGraphicsPathItem * addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddPath(path unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, path const QPainterPath &, pen const QPen &, brush const QBrush &), (path, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:196
// index:0
// QGraphicsPixmapItem * addPixmap(const class QPixmap &)
func (this *QGraphicsScene) AddPixmap(pixmap unsafe.Pointer) {
	// 0: (, pixmap const QPixmap &), (pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9addPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// QGraphicsPolygonItem * addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddPolygon(polygon unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, polygon const QPolygonF &, pen const QPen &, brush const QBrush &), (polygon, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// QGraphicsRectItem * addRect(const class QRectF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddRect(rect unsafe.Pointer, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 0: (, rect const QRectF &, pen const QPen &, brush const QBrush &), (rect, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// inline
// QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddRect_1(x float64, y float64, w float64, h float64, pen unsafe.Pointer, brush unsafe.Pointer) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, pen const QPen &, brush const QBrush &), (&x, &y, &w, &h, pen, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, pen, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:199
// index:0
// QGraphicsTextItem * addText(const class QString &, const class QFont &)
func (this *QGraphicsScene) AddText(text unsafe.Pointer, font unsafe.Pointer) {
	// 0: (, text const QString &, font const QFont &), (text, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addTextERK7QStringRK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:200
// index:0
// QGraphicsSimpleTextItem * addSimpleText(const class QString &, const class QFont &)
func (this *QGraphicsScene) AddSimpleText(text unsafe.Pointer, font unsafe.Pointer) {
	// 0: (, text const QString &, font const QFont &), (text, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:201
// index:0
// QGraphicsProxyWidget * addWidget(class QWidget *, Qt::WindowFlags)
func (this *QGraphicsScene) AddWidget(widget unsafe.Pointer, wFlags int) {
	// 0: (, widget QWidget *, QFlags<Qt::WindowType> wFlags), (widget, &wFlags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9addWidgetEP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, &wFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:208
// index:0
// void removeItem(class QGraphicsItem *)
func (this *QGraphicsScene) RemoveItem(item unsafe.Pointer) {
	// 0: (, item QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10removeItemEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:210
// index:0
// QGraphicsItem * focusItem()
func (this *QGraphicsScene) FocusItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene9focusItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:211
// index:0
// void setFocusItem(class QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) SetFocusItem(item unsafe.Pointer, focusReason int) {
	// 0: (, item QGraphicsItem *, focusReason Qt::FocusReason), (item, &focusReason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setFocusItemEP13QGraphicsItemN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:212
// index:0
// bool hasFocus()
func (this *QGraphicsScene) HasFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene8hasFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:213
// index:0
// void setFocus(Qt::FocusReason)
func (this *QGraphicsScene) SetFocus(focusReason int) {
	// 0: (, focusReason Qt::FocusReason), (&focusReason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:214
// index:0
// void clearFocus()
func (this *QGraphicsScene) ClearFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10clearFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:216
// index:0
// void setStickyFocus(_Bool)
func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14setStickyFocusEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:217
// index:0
// bool stickyFocus()
func (this *QGraphicsScene) StickyFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene11stickyFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:219
// index:0
// QGraphicsItem * mouseGrabberItem()
func (this *QGraphicsScene) MouseGrabberItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene16mouseGrabberItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:221
// index:0
// QBrush backgroundBrush()
func (this *QGraphicsScene) BackgroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15backgroundBrushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:222
// index:0
// void setBackgroundBrush(const class QBrush &)
func (this *QGraphicsScene) SetBackgroundBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:224
// index:0
// QBrush foregroundBrush()
func (this *QGraphicsScene) ForegroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15foregroundBrushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:225
// index:0
// void setForegroundBrush(const class QBrush &)
func (this *QGraphicsScene) SetForegroundBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setForegroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:227
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsScene) InputMethodQuery(query int) {
	// 0: (, query Qt::InputMethodQuery), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:229
// index:0
// QList<QGraphicsView *> views()
func (this *QGraphicsScene) Views() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5viewsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:231
// index:0
// inline
// void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	// 0: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6updateEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:257
// index:1
// void update(const class QRectF &)
func (this *QGraphicsScene) Update_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6updateERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:233
// index:0
// inline
// void invalidate(qreal, qreal, qreal, qreal, QGraphicsScene::SceneLayers)
func (this *QGraphicsScene) Invalidate(x float64, y float64, w float64, h float64, layers int) {
	// 0: (, x qreal, y qreal, w qreal, h qreal, QFlags<QGraphicsScene::SceneLayer> layers), (&x, &y, &w, &h, layers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateEdddd6QFlagsINS_10SceneLayerEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, layers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:258
// index:1
// void invalidate(const class QRectF &, QGraphicsScene::SceneLayers)
func (this *QGraphicsScene) Invalidate_1(rect unsafe.Pointer, layers int) {
	// 1: (, rect const QRectF &, QFlags<QGraphicsScene::SceneLayer> layers), (rect, layers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10invalidateERK6QRectF6QFlagsINS_10SceneLayerEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, layers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:236
// index:0
// QStyle * style()
func (this *QGraphicsScene) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5styleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:237
// index:0
// void setStyle(class QStyle *)
func (this *QGraphicsScene) SetStyle(style unsafe.Pointer) {
	// 0: (, style QStyle *), (style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene8setStyleEP6QStyle", ffiqt.FFI_TYPE_VOID, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:239
// index:0
// QFont font()
func (this *QGraphicsScene) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene4fontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:240
// index:0
// void setFont(const class QFont &)
func (this *QGraphicsScene) SetFont(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:242
// index:0
// QPalette palette()
func (this *QGraphicsScene) Palette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene7paletteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:243
// index:0
// void setPalette(const class QPalette &)
func (this *QGraphicsScene) SetPalette(palette unsafe.Pointer) {
	// 0: (, palette const QPalette &), (palette)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10setPaletteERK8QPalette", ffiqt.FFI_TYPE_VOID, this.GetCthis(), palette)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:245
// index:0
// bool isActive()
func (this *QGraphicsScene) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene8isActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:246
// index:0
// QGraphicsItem * activePanel()
func (this *QGraphicsScene) ActivePanel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene11activePanelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:247
// index:0
// void setActivePanel(class QGraphicsItem *)
func (this *QGraphicsScene) SetActivePanel(item unsafe.Pointer) {
	// 0: (, item QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:248
// index:0
// QGraphicsWidget * activeWindow()
func (this *QGraphicsScene) ActiveWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene12activeWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:249
// index:0
// void setActiveWindow(class QGraphicsWidget *)
func (this *QGraphicsScene) SetActiveWindow(widget unsafe.Pointer) {
	// 0: (, widget QGraphicsWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:251
// index:0
// bool sendEvent(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsScene) SendEvent(item unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, item QGraphicsItem *, event QEvent *), (item, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:253
// index:0
// qreal minimumRenderSize()
func (this *QGraphicsScene) MinimumRenderSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene17minimumRenderSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:254
// index:0
// void setMinimumRenderSize(qreal)
func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	// 0: (, minSize qreal), (&minSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene20setMinimumRenderSizeEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &minSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:259
// index:0
// void advance()
func (this *QGraphicsScene) Advance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7advanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:260
// index:0
// void clearSelection()
func (this *QGraphicsScene) ClearSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14clearSelectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:261
// index:0
// void clear()
func (this *QGraphicsScene) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:264
// index:0
// virtual
// bool event(class QEvent *)
func (this *QGraphicsScene) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:265
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QGraphicsScene) EventFilter(watched unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, watched QObject *, event QEvent *), (watched, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), watched, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:266
// index:0
// virtual
// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsScene) ContextMenuEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneContextMenuEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:267
// index:0
// virtual
// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragEnterEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:268
// index:0
// virtual
// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:269
// index:0
// virtual
// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:270
// index:0
// virtual
// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DropEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:271
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsScene) FocusInEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:272
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsScene) FocusOutEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:273
// index:0
// virtual
// void helpEvent(class QGraphicsSceneHelpEvent *)
func (this *QGraphicsScene) HelpEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHelpEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:274
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsScene) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:275
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsScene) KeyReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:276
// index:0
// virtual
// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:277
// index:0
// virtual
// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:278
// index:0
// virtual
// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:279
// index:0
// virtual
// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseDoubleClickEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:280
// index:0
// virtual
// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsScene) WheelEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneWheelEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:281
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsScene) InputMethodEvent(event unsafe.Pointer) {
	// 0: (, event QInputMethodEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:283
// index:0
// virtual
// void drawBackground(class QPainter *, const class QRectF &)
func (this *QGraphicsScene) DrawBackground(painter unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRectF &), (painter, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:284
// index:0
// virtual
// void drawForeground(class QPainter *, const class QRectF &)
func (this *QGraphicsScene) DrawForeground(painter unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRectF &), (painter, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:285
// index:0
// virtual
// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsScene) DrawItems(painter unsafe.Pointer, numItems int, items []interface{}, options []interface{}, widget unsafe.Pointer) {
	// 0: (, painter QPainter *, numItems int, items unsafe.Pointer, options unsafe.Pointer, widget QWidget *), (painter, &numItems, items, options, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, &numItems, items, options, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:295
// index:0
// bool focusNextPrevChild(_Bool)
func (this *QGraphicsScene) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:299
// index:0
// void sceneRectChanged(const class QRectF &)
func (this *QGraphicsScene) SceneRectChanged(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16sceneRectChangedERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:300
// index:0
// void selectionChanged()
func (this *QGraphicsScene) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:301
// index:0
// void focusItemChanged(class QGraphicsItem *, class QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) FocusItemChanged(newFocus unsafe.Pointer, oldFocus unsafe.Pointer, reason int) {
	// 0: (, newFocus QGraphicsItem *, oldFocus QGraphicsItem *, reason Qt::FocusReason), (newFocus, oldFocus, &reason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16focusItemChangedEP13QGraphicsItemS1_N2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newFocus, oldFocus, &reason)
	gopp.ErrPrint(err, rv)
}

//  body block end
