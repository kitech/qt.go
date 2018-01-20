//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsscene.h
// #include <qgraphicsscene.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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
func NewQGraphicsSceneFromPointer(cthis unsafe.Pointer) *QGraphicsScene {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsScene{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:98
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsScene) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:124
// index:0
// Public
// void QGraphicsScene(class QObject *)
func NewQGraphicsScene(parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:125
// index:1
// Public
// void QGraphicsScene(const class QRectF &, class QObject *)
func NewQGraphicsScene_1(sceneRect *qtcore.QRectF, parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sceneRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2ERK6QRectFP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:126
// index:2
// Public
// void QGraphicsScene(qreal, qreal, qreal, qreal, class QObject *)
func NewQGraphicsScene_2(x float64, y float64, width float64, height float64, parent unsafe.Pointer) *QGraphicsScene {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneC2EddddP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &width, &height, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:127
// index:0
// Public virtual
// void ~QGraphicsScene()
func DeleteQGraphicsScene(*QGraphicsScene) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsSceneD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:129
// index:0
// Public
// QRectF sceneRect()
func (this *QGraphicsScene) SceneRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene9sceneRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:130
// index:0
// Public inline
// qreal width()
func (this *QGraphicsScene) Width() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:131
// index:0
// Public inline
// qreal height()
func (this *QGraphicsScene) Height() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:132
// index:0
// Public
// void setSceneRect(const class QRectF &)
func (this *QGraphicsScene) SetSceneRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:133
// index:1
// Public inline
// void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setSceneRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:136
// index:0
// Public
// void render(class QPainter *, const class QRectF &, const class QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsScene) Render(painter unsafe.Pointer, target *qtcore.QRectF, source *qtcore.QRectF, aspectRatioMode int) {
	var convArg1 = target.GetCthis()
	var convArg2 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6renderEP8QPainterRK6QRectFS4_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2, &aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:140
// index:0
// Public
// QGraphicsScene::ItemIndexMethod itemIndexMethod()
func (this *QGraphicsScene) ItemIndexMethod() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15itemIndexMethodEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:141
// index:0
// Public
// void setItemIndexMethod(enum QGraphicsScene::ItemIndexMethod)
func (this *QGraphicsScene) SetItemIndexMethod(method int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setItemIndexMethodENS_15ItemIndexMethodE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &method)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:143
// index:0
// Public
// bool isSortCacheEnabled()
func (this *QGraphicsScene) IsSortCacheEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene18isSortCacheEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:144
// index:0
// Public
// void setSortCacheEnabled(_Bool)
func (this *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene19setSortCacheEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:146
// index:0
// Public
// int bspTreeDepth()
func (this *QGraphicsScene) BspTreeDepth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene12bspTreeDepthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:147
// index:0
// Public
// void setBspTreeDepth(int)
func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15setBspTreeDepthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:149
// index:0
// Public
// QRectF itemsBoundingRect()
func (this *QGraphicsScene) ItemsBoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene17itemsBoundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:151
// index:0
// Public
// QList<QGraphicsItem *> items(Qt::SortOrder)
func (this *QGraphicsScene) Items(order int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsEN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:153
// index:1
// Public
// QList<QGraphicsItem *> items(const class QPointF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_1(pos *qtcore.QPointF, mode int, order int, deviceTransform *qtgui.QTransform) interface{} {
	var convArg0 = pos.GetCthis()
	var convArg3 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK7QPointFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &order, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:154
// index:2
// Public
// QList<QGraphicsItem *> items(const class QRectF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_2(rect *qtcore.QRectF, mode int, order int, deviceTransform *qtgui.QTransform) interface{} {
	var convArg0 = rect.GetCthis()
	var convArg3 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK6QRectFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &order, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:155
// index:3
// Public
// QList<QGraphicsItem *> items(const class QPolygonF &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_3(polygon *qtgui.QPolygonF, mode int, order int, deviceTransform *qtgui.QTransform) interface{} {
	var convArg0 = polygon.GetCthis()
	var convArg3 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK9QPolygonFN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &order, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:156
// index:4
// Public
// QList<QGraphicsItem *> items(const class QPainterPath &, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_4(path *qtgui.QPainterPath, mode int, order int, deviceTransform *qtgui.QTransform) interface{} {
	var convArg0 = path.GetCthis()
	var convArg3 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsERK12QPainterPathN2Qt17ItemSelectionModeENS3_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &order, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:170
// index:5
// Public inline
// QList<QGraphicsItem *> items(qreal, qreal, qreal, qreal, Qt::ItemSelectionMode, Qt::SortOrder, const class QTransform &)
func (this *QGraphicsScene) Items_5(x float64, y float64, w float64, h float64, mode int, order int, deviceTransform *qtgui.QTransform) interface{} {
	var convArg6 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5itemsEddddN2Qt17ItemSelectionModeENS0_9SortOrderERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h, &mode, &order, convArg6)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:158
// index:0
// Public
// QList<QGraphicsItem *> collidingItems(const class QGraphicsItem *, Qt::ItemSelectionMode)
func (this *QGraphicsScene) CollidingItems(item unsafe.Pointer, mode int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene14collidingItemsEPK13QGraphicsItemN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:165
// index:0
// Public
// QGraphicsItem * itemAt(const class QPointF &, const class QTransform &)
func (this *QGraphicsScene) ItemAt(pos *qtcore.QPointF, deviceTransform *qtgui.QTransform) interface{} {
	var convArg0 = pos.GetCthis()
	var convArg1 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:179
// index:1
// Public inline
// QGraphicsItem * itemAt(qreal, qreal, const class QTransform &)
func (this *QGraphicsScene) ItemAt_1(x float64, y float64, deviceTransform *qtgui.QTransform) interface{} {
	var convArg2 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene6itemAtEddRK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:182
// index:0
// Public
// QList<QGraphicsItem *> selectedItems()
func (this *QGraphicsScene) SelectedItems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectedItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:183
// index:0
// Public
// QPainterPath selectionArea()
func (this *QGraphicsScene) SelectionArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene13selectionAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:184
// index:0
// Public
// void setSelectionArea(const class QPainterPath &, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea(path *qtgui.QPainterPath, deviceTransform *qtgui.QTransform) {
	var convArg0 = path.GetCthis()
	var convArg1 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:185
// index:1
// Public
// void setSelectionArea(const class QPainterPath &, Qt::ItemSelectionMode, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea_1(path *qtgui.QPainterPath, mode int, deviceTransform *qtgui.QTransform) {
	var convArg0 = path.GetCthis()
	var convArg2 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:186
// index:2
// Public
// void setSelectionArea(const class QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const class QTransform &)
func (this *QGraphicsScene) SetSelectionArea_2(path *qtgui.QPainterPath, selectionOperation int, mode int, deviceTransform *qtgui.QTransform) {
	var convArg0 = path.GetCthis()
	var convArg3 = deviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &selectionOperation, &mode, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:190
// index:0
// Public
// void destroyItemGroup(class QGraphicsItemGroup *)
func (this *QGraphicsScene) DestroyItemGroup(group unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:192
// index:0
// Public
// void addItem(class QGraphicsItem *)
func (this *QGraphicsScene) AddItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addItemEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:193
// index:0
// Public
// QGraphicsEllipseItem * addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddEllipse(rect *qtcore.QRectF, pen *qtgui.QPen, brush *qtgui.QBrush) interface{} {
	var convArg0 = rect.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:202
// index:1
// Public inline
// QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddEllipse_1(x float64, y float64, w float64, h float64, pen *qtgui.QPen, brush *qtgui.QBrush) interface{} {
	var convArg4 = pen.GetCthis()
	var convArg5 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h, convArg4, convArg5)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:194
// index:0
// Public
// QGraphicsLineItem * addLine(const class QLineF &, const class QPen &)
func (this *QGraphicsScene) AddLine(line *qtcore.QLineF, pen *qtgui.QPen) interface{} {
	var convArg0 = line.GetCthis()
	var convArg1 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineERK6QLineFRK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:204
// index:1
// Public inline
// QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const class QPen &)
func (this *QGraphicsScene) AddLine_1(x1 float64, y1 float64, x2 float64, y2 float64, pen *qtgui.QPen) interface{} {
	var convArg4 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addLineEddddRK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:195
// index:0
// Public
// QGraphicsPathItem * addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddPath(path *qtgui.QPainterPath, pen *qtgui.QPen, brush *qtgui.QBrush) interface{} {
	var convArg0 = path.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:196
// index:0
// Public
// QGraphicsPixmapItem * addPixmap(const class QPixmap &)
func (this *QGraphicsScene) AddPixmap(pixmap *qtgui.QPixmap) interface{} {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9addPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:197
// index:0
// Public
// QGraphicsPolygonItem * addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddPolygon(polygon *qtgui.QPolygonF, pen *qtgui.QPen, brush *qtgui.QBrush) interface{} {
	var convArg0 = polygon.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:198
// index:0
// Public
// QGraphicsRectItem * addRect(const class QRectF &, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddRect(rect *qtcore.QRectF, pen *qtgui.QPen, brush *qtgui.QBrush) interface{} {
	var convArg0 = rect.GetCthis()
	var convArg1 = pen.GetCthis()
	var convArg2 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:206
// index:1
// Public inline
// QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
func (this *QGraphicsScene) AddRect_1(x float64, y float64, w float64, h float64, pen *qtgui.QPen, brush *qtgui.QBrush) interface{} {
	var convArg4 = pen.GetCthis()
	var convArg5 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h, convArg4, convArg5)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:199
// index:0
// Public
// QGraphicsTextItem * addText(const class QString &, const class QFont &)
func (this *QGraphicsScene) AddText(text *qtcore.QString, font *qtgui.QFont) interface{} {
	var convArg0 = text.GetCthis()
	var convArg1 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7addTextERK7QStringRK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:200
// index:0
// Public
// QGraphicsSimpleTextItem * addSimpleText(const class QString &, const class QFont &)
func (this *QGraphicsScene) AddSimpleText(text *qtcore.QString, font *qtgui.QFont) interface{} {
	var convArg0 = text.GetCthis()
	var convArg1 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:208
// index:0
// Public
// void removeItem(class QGraphicsItem *)
func (this *QGraphicsScene) RemoveItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10removeItemEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:210
// index:0
// Public
// QGraphicsItem * focusItem()
func (this *QGraphicsScene) FocusItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene9focusItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:211
// index:0
// Public
// void setFocusItem(class QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) SetFocusItem(item unsafe.Pointer, focusReason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12setFocusItemEP13QGraphicsItemN2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:212
// index:0
// Public
// bool hasFocus()
func (this *QGraphicsScene) HasFocus() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene8hasFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:213
// index:0
// Public
// void setFocus(Qt::FocusReason)
func (this *QGraphicsScene) SetFocus(focusReason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:214
// index:0
// Public
// void clearFocus()
func (this *QGraphicsScene) ClearFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10clearFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:216
// index:0
// Public
// void setStickyFocus(_Bool)
func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14setStickyFocusEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:217
// index:0
// Public
// bool stickyFocus()
func (this *QGraphicsScene) StickyFocus() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene11stickyFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:219
// index:0
// Public
// QGraphicsItem * mouseGrabberItem()
func (this *QGraphicsScene) MouseGrabberItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene16mouseGrabberItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:221
// index:0
// Public
// QBrush backgroundBrush()
func (this *QGraphicsScene) BackgroundBrush() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15backgroundBrushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:222
// index:0
// Public
// void setBackgroundBrush(const class QBrush &)
func (this *QGraphicsScene) SetBackgroundBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setBackgroundBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:224
// index:0
// Public
// QBrush foregroundBrush()
func (this *QGraphicsScene) ForegroundBrush() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene15foregroundBrushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:225
// index:0
// Public
// void setForegroundBrush(const class QBrush &)
func (this *QGraphicsScene) SetForegroundBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18setForegroundBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:227
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsScene) InputMethodQuery(query int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:229
// index:0
// Public
// QList<QGraphicsView *> views()
func (this *QGraphicsScene) Views() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5viewsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:231
// index:0
// Public inline
// void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6updateEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:257
// index:1
// Public
// void update(const class QRectF &)
func (this *QGraphicsScene) Update_1(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene6updateERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:236
// index:0
// Public
// QStyle * style()
func (this *QGraphicsScene) Style() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:237
// index:0
// Public
// void setStyle(class QStyle *)
func (this *QGraphicsScene) SetStyle(style unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene8setStyleEP6QStyle", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:239
// index:0
// Public
// QFont font()
func (this *QGraphicsScene) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:240
// index:0
// Public
// void setFont(const class QFont &)
func (this *QGraphicsScene) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:242
// index:0
// Public
// QPalette palette()
func (this *QGraphicsScene) Palette() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene7paletteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:243
// index:0
// Public
// void setPalette(const class QPalette &)
func (this *QGraphicsScene) SetPalette(palette *qtgui.QPalette) {
	var convArg0 = palette.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10setPaletteERK8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:245
// index:0
// Public
// bool isActive()
func (this *QGraphicsScene) IsActive() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:246
// index:0
// Public
// QGraphicsItem * activePanel()
func (this *QGraphicsScene) ActivePanel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene11activePanelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:247
// index:0
// Public
// void setActivePanel(class QGraphicsItem *)
func (this *QGraphicsScene) SetActivePanel(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:248
// index:0
// Public
// QGraphicsWidget * activeWindow()
func (this *QGraphicsScene) ActiveWindow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene12activeWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:249
// index:0
// Public
// void setActiveWindow(class QGraphicsWidget *)
func (this *QGraphicsScene) SetActiveWindow(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:251
// index:0
// Public
// bool sendEvent(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsScene) SendEvent(item unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:253
// index:0
// Public
// qreal minimumRenderSize()
func (this *QGraphicsScene) MinimumRenderSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScene17minimumRenderSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:254
// index:0
// Public
// void setMinimumRenderSize(qreal)
func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene20setMinimumRenderSizeEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:259
// index:0
// Public
// void advance()
func (this *QGraphicsScene) Advance() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene7advanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:260
// index:0
// Public
// void clearSelection()
func (this *QGraphicsScene) ClearSelection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14clearSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:261
// index:0
// Public
// void clear()
func (this *QGraphicsScene) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:264
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QGraphicsScene) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:265
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QGraphicsScene) EventFilter(watched unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), watched, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:266
// index:0
// Protected virtual
// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsScene) ContextMenuEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:267
// index:0
// Protected virtual
// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragEnterEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:268
// index:0
// Protected virtual
// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:269
// index:0
// Protected virtual
// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DragLeaveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:270
// index:0
// Protected virtual
// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsScene) DropEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:271
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsScene) FocusInEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:272
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsScene) FocusOutEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:273
// index:0
// Protected virtual
// void helpEvent(class QGraphicsSceneHelpEvent *)
func (this *QGraphicsScene) HelpEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:274
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsScene) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:275
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsScene) KeyReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:276
// index:0
// Protected virtual
// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:277
// index:0
// Protected virtual
// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:278
// index:0
// Protected virtual
// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:279
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsScene) MouseDoubleClickEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:280
// index:0
// Protected virtual
// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsScene) WheelEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:281
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsScene) InputMethodEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:283
// index:0
// Protected virtual
// void drawBackground(class QPainter *, const class QRectF &)
func (this *QGraphicsScene) DrawBackground(painter unsafe.Pointer, rect *qtcore.QRectF) {
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:284
// index:0
// Protected virtual
// void drawForeground(class QPainter *, const class QRectF &)
func (this *QGraphicsScene) DrawForeground(painter unsafe.Pointer, rect *qtcore.QRectF) {
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:285
// index:0
// Protected virtual
// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsScene) DrawItems(painter unsafe.Pointer, numItems int, items []interface{}, options []interface{}, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, &numItems, items, options, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:295
// index:0
// Protected
// bool focusNextPrevChild(_Bool)
func (this *QGraphicsScene) FocusNextPrevChild(next bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:299
// index:0
// Public
// void sceneRectChanged(const class QRectF &)
func (this *QGraphicsScene) SceneRectChanged(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16sceneRectChangedERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:300
// index:0
// Public
// void selectionChanged()
func (this *QGraphicsScene) SelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16selectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsscene.h:301
// index:0
// Public
// void focusItemChanged(class QGraphicsItem *, class QGraphicsItem *, Qt::FocusReason)
func (this *QGraphicsScene) FocusItemChanged(newFocus unsafe.Pointer, oldFocus unsafe.Pointer, reason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScene16focusItemChangedEP13QGraphicsItemS1_N2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newFocus, oldFocus, &reason)
	gopp.ErrPrint(err, rv)
}

//  body block end
