//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsview.h
// #include <qgraphicsview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QGraphicsView struct {
	*QAbstractScrollArea
}

func (this *QGraphicsView) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:117
// index:0
// void QGraphicsView(class QWidget *)
func NewQGraphicsView(parent unsafe.Pointer) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(cthis)
	return gothis
}
func NewQGraphicsViewFromPointer(cthis unsafe.Pointer) *QGraphicsView {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QGraphicsView{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:118
// index:1
// void QGraphicsView(class QGraphicsScene *, class QWidget *)
func NewQGraphicsView_1(scene unsafe.Pointer, parent unsafe.Pointer) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, scene, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:239
// index:2
// void QGraphicsView(class QGraphicsViewPrivate &, class QWidget *)
func NewQGraphicsView_2(arg0 unsafe.Pointer, parent unsafe.Pointer) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2ER20QGraphicsViewPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:119
// index:0
// virtual
// void ~QGraphicsView()
func DeleteQGraphicsView(*QGraphicsView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:121
// index:0
// virtual
// QSize sizeHint()
func (this *QGraphicsView) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:123
// index:0
// QPainter::RenderHints renderHints()
func (this *QGraphicsView) RenderHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView11renderHintsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// void setRenderHint(class QPainter::RenderHint, _Bool)
func (this *QGraphicsView) SetRenderHint(hint int, enabled bool) {
	// 0: (, hint QPainter::RenderHint, enabled bool), (&hint, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hint, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:125
// index:0
// void setRenderHints(class QPainter::RenderHints)
func (this *QGraphicsView) SetRenderHints(hints int) {
	// 0: (, QFlags<QPainter::RenderHint> hints), (&hints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14setRenderHintsE6QFlagsIN8QPainter10RenderHintEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:127
// index:0
// Qt::Alignment alignment()
func (this *QGraphicsView) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:128
// index:0
// void setAlignment(Qt::Alignment)
func (this *QGraphicsView) SetAlignment(alignment int) {
	// 0: (, QFlags<Qt::AlignmentFlag> alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:130
// index:0
// QGraphicsView::ViewportAnchor transformationAnchor()
func (this *QGraphicsView) TransformationAnchor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView20transformationAnchorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:131
// index:0
// void setTransformationAnchor(enum QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetTransformationAnchor(anchor int) {
	// 0: (, anchor QGraphicsView::ViewportAnchor), (&anchor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView23setTransformationAnchorENS_14ViewportAnchorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:133
// index:0
// QGraphicsView::ViewportAnchor resizeAnchor()
func (this *QGraphicsView) ResizeAnchor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12resizeAnchorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:134
// index:0
// void setResizeAnchor(enum QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetResizeAnchor(anchor int) {
	// 0: (, anchor QGraphicsView::ViewportAnchor), (&anchor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15setResizeAnchorENS_14ViewportAnchorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:136
// index:0
// QGraphicsView::ViewportUpdateMode viewportUpdateMode()
func (this *QGraphicsView) ViewportUpdateMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView18viewportUpdateModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:137
// index:0
// void setViewportUpdateMode(enum QGraphicsView::ViewportUpdateMode)
func (this *QGraphicsView) SetViewportUpdateMode(mode int) {
	// 0: (, mode QGraphicsView::ViewportUpdateMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView21setViewportUpdateModeENS_18ViewportUpdateModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:139
// index:0
// QGraphicsView::OptimizationFlags optimizationFlags()
func (this *QGraphicsView) OptimizationFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView17optimizationFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// void setOptimizationFlag(enum QGraphicsView::OptimizationFlag, _Bool)
func (this *QGraphicsView) SetOptimizationFlag(flag int, enabled bool) {
	// 0: (, flag QGraphicsView::OptimizationFlag, enabled bool), (&flag, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flag, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:141
// index:0
// void setOptimizationFlags(QGraphicsView::OptimizationFlags)
func (this *QGraphicsView) SetOptimizationFlags(flags int) {
	// 0: (, QFlags<QGraphicsView::OptimizationFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView20setOptimizationFlagsE6QFlagsINS_16OptimizationFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:143
// index:0
// QGraphicsView::DragMode dragMode()
func (this *QGraphicsView) DragMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView8dragModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:144
// index:0
// void setDragMode(enum QGraphicsView::DragMode)
func (this *QGraphicsView) SetDragMode(mode int) {
	// 0: (, mode QGraphicsView::DragMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11setDragModeENS_8DragModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:147
// index:0
// Qt::ItemSelectionMode rubberBandSelectionMode()
func (this *QGraphicsView) RubberBandSelectionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView23rubberBandSelectionModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:148
// index:0
// void setRubberBandSelectionMode(Qt::ItemSelectionMode)
func (this *QGraphicsView) SetRubberBandSelectionMode(mode int) {
	// 0: (, mode Qt::ItemSelectionMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView26setRubberBandSelectionModeEN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:149
// index:0
// QRect rubberBandRect()
func (this *QGraphicsView) RubberBandRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView14rubberBandRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:152
// index:0
// QGraphicsView::CacheMode cacheMode()
func (this *QGraphicsView) CacheMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9cacheModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:153
// index:0
// void setCacheMode(QGraphicsView::CacheMode)
func (this *QGraphicsView) SetCacheMode(mode int) {
	// 0: (, QFlags<QGraphicsView::CacheModeFlag> mode), (mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setCacheModeE6QFlagsINS_13CacheModeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:154
// index:0
// void resetCachedContent()
func (this *QGraphicsView) ResetCachedContent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18resetCachedContentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:156
// index:0
// bool isInteractive()
func (this *QGraphicsView) IsInteractive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView13isInteractiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:157
// index:0
// void setInteractive(_Bool)
func (this *QGraphicsView) SetInteractive(allowed bool) {
	// 0: (, allowed bool), (&allowed)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14setInteractiveEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &allowed)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:159
// index:0
// QGraphicsScene * scene()
func (this *QGraphicsView) Scene() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5sceneEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:160
// index:0
// void setScene(class QGraphicsScene *)
func (this *QGraphicsView) SetScene(scene unsafe.Pointer) {
	// 0: (, scene QGraphicsScene *), (scene)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8setSceneEP14QGraphicsScene", ffiqt.FFI_TYPE_VOID, this.GetCthis(), scene)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:162
// index:0
// QRectF sceneRect()
func (this *QGraphicsView) SceneRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9sceneRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:163
// index:0
// void setSceneRect(const class QRectF &)
func (this *QGraphicsView) SetSceneRect(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:164
// index:1
// inline
// void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:166
// index:0
// QMatrix matrix()
func (this *QGraphicsView) Matrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6matrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:167
// index:0
// void setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsView) SetMatrix(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QMatrix &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:168
// index:0
// void resetMatrix()
func (this *QGraphicsView) ResetMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11resetMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:169
// index:0
// QTransform transform()
func (this *QGraphicsView) Transform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9transformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:170
// index:0
// QTransform viewportTransform()
func (this *QGraphicsView) ViewportTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView17viewportTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:171
// index:0
// bool isTransformed()
func (this *QGraphicsView) IsTransformed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView13isTransformedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:172
// index:0
// void setTransform(const class QTransform &, _Bool)
func (this *QGraphicsView) SetTransform(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QTransform &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setTransformERK10QTransformb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:173
// index:0
// void resetTransform()
func (this *QGraphicsView) ResetTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14resetTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:174
// index:0
// void rotate(qreal)
func (this *QGraphicsView) Rotate(angle float64) {
	// 0: (, angle qreal), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView6rotateEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:175
// index:0
// void scale(qreal, qreal)
func (this *QGraphicsView) Scale(sx float64, sy float64) {
	// 0: (, sx qreal, sy qreal), (&sx, &sy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5scaleEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:176
// index:0
// void shear(qreal, qreal)
func (this *QGraphicsView) Shear(sh float64, sv float64) {
	// 0: (, sh qreal, sv qreal), (&sh, &sv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5shearEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sh, &sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:177
// index:0
// void translate(qreal, qreal)
func (this *QGraphicsView) Translate(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9translateEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:179
// index:0
// void centerOn(const class QPointF &)
func (this *QGraphicsView) CenterOn(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:180
// index:1
// inline
// void centerOn(qreal, qreal)
func (this *QGraphicsView) CenterOn_1(x float64, y float64) {
	// 1: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:181
// index:2
// void centerOn(const class QGraphicsItem *)
func (this *QGraphicsView) CenterOn_2(item unsafe.Pointer) {
	// 2: (, item const QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// void ensureVisible(const class QRectF &, int, int)
func (this *QGraphicsView) EnsureVisible(rect unsafe.Pointer, xmargin int, ymargin int) {
	// 0: (, rect const QRectF &, xmargin int, ymargin int), (rect, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// inline
// void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsView) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, xmargin int, ymargin int), (&x, &y, &w, &h, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// void ensureVisible(const class QGraphicsItem *, int, int)
func (this *QGraphicsView) EnsureVisible_2(item unsafe.Pointer, xmargin int, ymargin int) {
	// 2: (, item const QGraphicsItem *, xmargin int, ymargin int), (item, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:185
// index:0
// void fitInView(const class QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView(rect unsafe.Pointer, aspectRadioMode int) {
	// 0: (, rect const QRectF &, aspectRadioMode Qt::AspectRatioMode), (rect, &aspectRadioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewERK6QRectFN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// inline
// void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_1(x float64, y float64, w float64, h float64, aspectRadioMode int) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, aspectRadioMode Qt::AspectRatioMode), (&x, &y, &w, &h, &aspectRadioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:188
// index:2
// void fitInView(const class QGraphicsItem *, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_2(item unsafe.Pointer, aspectRadioMode int) {
	// 2: (, item const QGraphicsItem *, aspectRadioMode Qt::AspectRatioMode), (item, &aspectRadioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEPK13QGraphicsItemN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// void render(class QPainter *, const class QRectF &, const class QRect &, Qt::AspectRatioMode)
func (this *QGraphicsView) Render(painter unsafe.Pointer, target unsafe.Pointer, source unsafe.Pointer, aspectRatioMode int) {
	// 0: (, painter QPainter *, target const QRectF &, source const QRect &, aspectRatioMode Qt::AspectRatioMode), (painter, target, source, &aspectRatioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, target, source, &aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:194
// index:0
// QList<QGraphicsItem *> items()
func (this *QGraphicsView) Items() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:195
// index:1
// QList<QGraphicsItem *> items(const class QPoint &)
func (this *QGraphicsView) Items_1(pos unsafe.Pointer) {
	// 1: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:196
// index:2
// inline
// QList<QGraphicsItem *> items(int, int)
func (this *QGraphicsView) Items_2(x int, y int) {
	// 2: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:197
// index:3
// QList<QGraphicsItem *> items(const class QRect &, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_3(rect unsafe.Pointer, mode int) {
	// 3: (, rect const QRect &, mode Qt::ItemSelectionMode), (rect, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK5QRectN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:198
// index:4
// inline
// QList<QGraphicsItem *> items(int, int, int, int, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_4(x int, y int, w int, h int, mode int) {
	// 4: (, x int, y int, w int, h int, mode Qt::ItemSelectionMode), (&x, &y, &w, &h, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEiiiiN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:199
// index:5
// QList<QGraphicsItem *> items(const class QPolygon &, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_5(polygon unsafe.Pointer, mode int) {
	// 5: (, polygon const QPolygon &, mode Qt::ItemSelectionMode), (polygon, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK8QPolygonN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:200
// index:6
// QList<QGraphicsItem *> items(const class QPainterPath &, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_6(path unsafe.Pointer, mode int) {
	// 6: (, path const QPainterPath &, mode Qt::ItemSelectionMode), (path, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK12QPainterPathN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:201
// index:0
// QGraphicsItem * itemAt(const class QPoint &)
func (this *QGraphicsView) ItemAt(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:202
// index:1
// inline
// QGraphicsItem * itemAt(int, int)
func (this *QGraphicsView) ItemAt_1(x int, y int) {
	// 1: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:204
// index:0
// QPointF mapToScene(const class QPoint &)
func (this *QGraphicsView) MapToScene(point unsafe.Pointer) {
	// 0: (, point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:205
// index:1
// QPolygonF mapToScene(const class QRect &)
func (this *QGraphicsView) MapToScene_1(rect unsafe.Pointer) {
	// 1: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:206
// index:2
// QPolygonF mapToScene(const class QPolygon &)
func (this *QGraphicsView) MapToScene_2(polygon unsafe.Pointer) {
	// 2: (, polygon const QPolygon &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:207
// index:3
// QPainterPath mapToScene(const class QPainterPath &)
func (this *QGraphicsView) MapToScene_3(path unsafe.Pointer) {
	// 3: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:212
// index:4
// inline
// QPointF mapToScene(int, int)
func (this *QGraphicsView) MapToScene_4(x int, y int) {
	// 4: (, x int, y int), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:213
// index:5
// inline
// QPolygonF mapToScene(int, int, int, int)
func (this *QGraphicsView) MapToScene_5(x int, y int, w int, h int) {
	// 5: (, x int, y int, w int, h int), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:208
// index:0
// QPoint mapFromScene(const class QPointF &)
func (this *QGraphicsView) MapFromScene(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:209
// index:1
// QPolygon mapFromScene(const class QRectF &)
func (this *QGraphicsView) MapFromScene_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:210
// index:2
// QPolygon mapFromScene(const class QPolygonF &)
func (this *QGraphicsView) MapFromScene_2(polygon unsafe.Pointer) {
	// 2: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:211
// index:3
// QPainterPath mapFromScene(const class QPainterPath &)
func (this *QGraphicsView) MapFromScene_3(path unsafe.Pointer) {
	// 3: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:214
// index:4
// inline
// QPoint mapFromScene(qreal, qreal)
func (this *QGraphicsView) MapFromScene_4(x float64, y float64) {
	// 4: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:215
// index:5
// inline
// QPolygon mapFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) MapFromScene_5(x float64, y float64, w float64, h float64) {
	// 5: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:217
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsView) InputMethodQuery(query int) {
	// 0: (, query Qt::InputMethodQuery), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:219
// index:0
// QBrush backgroundBrush()
func (this *QGraphicsView) BackgroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView15backgroundBrushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:220
// index:0
// void setBackgroundBrush(const class QBrush &)
func (this *QGraphicsView) SetBackgroundBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18setBackgroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:222
// index:0
// QBrush foregroundBrush()
func (this *QGraphicsView) ForegroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView15foregroundBrushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:223
// index:0
// void setForegroundBrush(const class QBrush &)
func (this *QGraphicsView) SetForegroundBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18setForegroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:227
// index:0
// void invalidateScene(const class QRectF &, class QGraphicsScene::SceneLayers)
func (this *QGraphicsView) InvalidateScene(rect unsafe.Pointer, layers int) {
	// 0: (, rect const QRectF &, QFlags<QGraphicsScene::SceneLayer> layers), (rect, &layers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15invalidateSceneERK6QRectF6QFlagsIN14QGraphicsScene10SceneLayerEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &layers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:228
// index:0
// void updateSceneRect(const class QRectF &)
func (this *QGraphicsView) UpdateSceneRect(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15updateSceneRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:232
// index:0
// void rubberBandChanged(class QRect, class QPointF, class QPointF)
func (this *QGraphicsView) RubberBandChanged(viewportRect unsafe.Pointer, fromScenePoint unsafe.Pointer, toScenePoint unsafe.Pointer) {
	// 0: (, viewportRect QRect, fromScenePoint QPointF, toScenePoint QPointF), (viewportRect, fromScenePoint, toScenePoint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView17rubberBandChangedE5QRect7QPointFS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), viewportRect, fromScenePoint, toScenePoint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:236
// index:0
// virtual
// void setupViewport(class QWidget *)
func (this *QGraphicsView) SetupViewport(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13setupViewportEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:240
// index:0
// virtual
// bool event(class QEvent *)
func (this *QGraphicsView) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:241
// index:0
// virtual
// bool viewportEvent(class QEvent *)
func (this *QGraphicsView) ViewportEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:244
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QGraphicsView) ContextMenuEvent(event unsafe.Pointer) {
	// 0: (, event QContextMenuEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:247
// index:0
// virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QGraphicsView) DragEnterEvent(event unsafe.Pointer) {
	// 0: (, event QDragEnterEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:248
// index:0
// virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QGraphicsView) DragLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QDragLeaveEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:249
// index:0
// virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QGraphicsView) DragMoveEvent(event unsafe.Pointer) {
	// 0: (, event QDragMoveEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:250
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QGraphicsView) DropEvent(event unsafe.Pointer) {
	// 0: (, event QDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:252
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsView) FocusInEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:253
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QGraphicsView) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:254
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsView) FocusOutEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:255
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsView) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:256
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsView) KeyReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:257
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QGraphicsView) MouseDoubleClickEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:258
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QGraphicsView) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:259
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QGraphicsView) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:260
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QGraphicsView) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:262
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QGraphicsView) WheelEvent(event unsafe.Pointer) {
	// 0: (, event QWheelEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:264
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QGraphicsView) PaintEvent(event unsafe.Pointer) {
	// 0: (, event QPaintEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:265
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QGraphicsView) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:266
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QGraphicsView) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:267
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QGraphicsView) ShowEvent(event unsafe.Pointer) {
	// 0: (, event QShowEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:268
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsView) InputMethodEvent(event unsafe.Pointer) {
	// 0: (, event QInputMethodEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:270
// index:0
// virtual
// void drawBackground(class QPainter *, const class QRectF &)
func (this *QGraphicsView) DrawBackground(painter unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRectF &), (painter, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:271
// index:0
// virtual
// void drawForeground(class QPainter *, const class QRectF &)
func (this *QGraphicsView) DrawForeground(painter unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRectF &), (painter, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:272
// index:0
// virtual
// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *)
func (this *QGraphicsView) DrawItems(painter unsafe.Pointer, numItems int, items []interface{}, options []interface{}) {
	// 0: (, painter QPainter *, numItems int, items unsafe.Pointer, options unsafe.Pointer), (painter, &numItems, items, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, &numItems, items, options)
	gopp.ErrPrint(err, rv)
}

//  body block end
