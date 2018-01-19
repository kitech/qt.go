//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsview.h
// #include <qgraphicsview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:117
// index:0
// void QGraphicsView(class QWidget *)
func NewQGraphicsView(parent unsafe.Pointer) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsView{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:118
// index:1
// void QGraphicsView(class QGraphicsScene *, class QWidget *)
func NewQGraphicsView_1(scene unsafe.Pointer, parent unsafe.Pointer) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, scene, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsView{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:123
// index:0
// QPainter::RenderHints renderHints()
func (this *QGraphicsView) RenderHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView11renderHintsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// void setRenderHint(class QPainter::RenderHint, _Bool)
func (this *QGraphicsView) SetRenderHint(hint int, enabled bool) {
	// 0: (, QPainter::RenderHint hint, bool enabled), (&hint, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hint, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:127
// index:0
// Qt::Alignment alignment()
func (this *QGraphicsView) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:130
// index:0
// QGraphicsView::ViewportAnchor transformationAnchor()
func (this *QGraphicsView) TransformationAnchor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView20transformationAnchorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:131
// index:0
// void setTransformationAnchor(enum QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetTransformationAnchor(anchor int) {
	// 0: (, QGraphicsView::ViewportAnchor anchor), (&anchor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView23setTransformationAnchorENS_14ViewportAnchorE", ffiqt.FFI_TYPE_VOID, this.cthis, &anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:133
// index:0
// QGraphicsView::ViewportAnchor resizeAnchor()
func (this *QGraphicsView) ResizeAnchor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12resizeAnchorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:134
// index:0
// void setResizeAnchor(enum QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetResizeAnchor(anchor int) {
	// 0: (, QGraphicsView::ViewportAnchor anchor), (&anchor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15setResizeAnchorENS_14ViewportAnchorE", ffiqt.FFI_TYPE_VOID, this.cthis, &anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:136
// index:0
// QGraphicsView::ViewportUpdateMode viewportUpdateMode()
func (this *QGraphicsView) ViewportUpdateMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView18viewportUpdateModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:137
// index:0
// void setViewportUpdateMode(enum QGraphicsView::ViewportUpdateMode)
func (this *QGraphicsView) SetViewportUpdateMode(mode int) {
	// 0: (, QGraphicsView::ViewportUpdateMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView21setViewportUpdateModeENS_18ViewportUpdateModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:139
// index:0
// QGraphicsView::OptimizationFlags optimizationFlags()
func (this *QGraphicsView) OptimizationFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView17optimizationFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// void setOptimizationFlag(enum QGraphicsView::OptimizationFlag, _Bool)
func (this *QGraphicsView) SetOptimizationFlag(flag int, enabled bool) {
	// 0: (, QGraphicsView::OptimizationFlag flag, bool enabled), (&flag, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", ffiqt.FFI_TYPE_VOID, this.cthis, &flag, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:143
// index:0
// QGraphicsView::DragMode dragMode()
func (this *QGraphicsView) DragMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView8dragModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:144
// index:0
// void setDragMode(enum QGraphicsView::DragMode)
func (this *QGraphicsView) SetDragMode(mode int) {
	// 0: (, QGraphicsView::DragMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11setDragModeENS_8DragModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:147
// index:0
// Qt::ItemSelectionMode rubberBandSelectionMode()
func (this *QGraphicsView) RubberBandSelectionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView23rubberBandSelectionModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:148
// index:0
// void setRubberBandSelectionMode(Qt::ItemSelectionMode)
func (this *QGraphicsView) SetRubberBandSelectionMode(mode int) {
	// 0: (, Qt::ItemSelectionMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView26setRubberBandSelectionModeEN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:149
// index:0
// QRect rubberBandRect()
func (this *QGraphicsView) RubberBandRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView14rubberBandRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:152
// index:0
// QGraphicsView::CacheMode cacheMode()
func (this *QGraphicsView) CacheMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9cacheModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:154
// index:0
// void resetCachedContent()
func (this *QGraphicsView) ResetCachedContent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18resetCachedContentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:156
// index:0
// bool isInteractive()
func (this *QGraphicsView) IsInteractive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView13isInteractiveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:157
// index:0
// void setInteractive(_Bool)
func (this *QGraphicsView) SetInteractive(allowed bool) {
	// 0: (, bool allowed), (&allowed)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14setInteractiveEb", ffiqt.FFI_TYPE_VOID, this.cthis, &allowed)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:159
// index:0
// QGraphicsScene * scene()
func (this *QGraphicsView) Scene() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5sceneEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:160
// index:0
// void setScene(class QGraphicsScene *)
func (this *QGraphicsView) SetScene(scene unsafe.Pointer) {
	// 0: (, QGraphicsScene * scene), (scene)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8setSceneEP14QGraphicsScene", ffiqt.FFI_TYPE_VOID, this.cthis, scene)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:162
// index:0
// QRectF sceneRect()
func (this *QGraphicsView) SceneRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9sceneRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:163
// index:0
// void setSceneRect(const class QRectF &)
func (this *QGraphicsView) SetSceneRect(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:164
// index:1
// inline
// void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, qreal x, qreal y, qreal w, qreal h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:166
// index:0
// QMatrix matrix()
func (this *QGraphicsView) Matrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6matrixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:167
// index:0
// void setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsView) SetMatrix(matrix unsafe.Pointer, combine bool) {
	// 0: (, const QMatrix & matrix, bool combine), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_VOID, this.cthis, matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:168
// index:0
// void resetMatrix()
func (this *QGraphicsView) ResetMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11resetMatrixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:169
// index:0
// QTransform transform()
func (this *QGraphicsView) Transform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9transformEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:170
// index:0
// QTransform viewportTransform()
func (this *QGraphicsView) ViewportTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView17viewportTransformEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:171
// index:0
// bool isTransformed()
func (this *QGraphicsView) IsTransformed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView13isTransformedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:172
// index:0
// void setTransform(const class QTransform &, _Bool)
func (this *QGraphicsView) SetTransform(matrix unsafe.Pointer, combine bool) {
	// 0: (, const QTransform & matrix, bool combine), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setTransformERK10QTransformb", ffiqt.FFI_TYPE_VOID, this.cthis, matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:173
// index:0
// void resetTransform()
func (this *QGraphicsView) ResetTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14resetTransformEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:174
// index:0
// void rotate(qreal)
func (this *QGraphicsView) Rotate(angle float64) {
	// 0: (, qreal angle), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView6rotateEd", ffiqt.FFI_TYPE_VOID, this.cthis, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:175
// index:0
// void scale(qreal, qreal)
func (this *QGraphicsView) Scale(sx float64, sy float64) {
	// 0: (, qreal sx, qreal sy), (&sx, &sy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5scaleEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:176
// index:0
// void shear(qreal, qreal)
func (this *QGraphicsView) Shear(sh float64, sv float64) {
	// 0: (, qreal sh, qreal sv), (&sh, &sv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5shearEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &sh, &sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:177
// index:0
// void translate(qreal, qreal)
func (this *QGraphicsView) Translate(dx float64, dy float64) {
	// 0: (, qreal dx, qreal dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9translateEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:179
// index:0
// void centerOn(const class QPointF &)
func (this *QGraphicsView) CenterOn(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:180
// index:1
// inline
// void centerOn(qreal, qreal)
func (this *QGraphicsView) CenterOn_1(x float64, y float64) {
	// 1: (, qreal x, qreal y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:181
// index:2
// void centerOn(const class QGraphicsItem *)
func (this *QGraphicsView) CenterOn_2(item unsafe.Pointer) {
	// 2: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// void ensureVisible(const class QRectF &, int, int)
func (this *QGraphicsView) EnsureVisible(rect unsafe.Pointer, xmargin int, ymargin int) {
	// 0: (, const QRectF & rect, int xmargin, int ymargin), (rect, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.cthis, rect, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// inline
// void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsView) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	// 1: (, qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin), (&x, &y, &w, &h, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// void ensureVisible(const class QGraphicsItem *, int, int)
func (this *QGraphicsView) EnsureVisible_2(item unsafe.Pointer, xmargin int, ymargin int) {
	// 2: (, const QGraphicsItem * item, int xmargin, int ymargin), (item, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", ffiqt.FFI_TYPE_VOID, this.cthis, item, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:185
// index:0
// void fitInView(const class QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView(rect unsafe.Pointer, aspectRadioMode int) {
	// 0: (, const QRectF & rect, Qt::AspectRatioMode aspectRadioMode), (rect, &aspectRadioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewERK6QRectFN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.cthis, rect, &aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// inline
// void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_1(x float64, y float64, w float64, h float64, aspectRadioMode int) {
	// 1: (, qreal x, qreal y, qreal w, qreal h, Qt::AspectRatioMode aspectRadioMode), (&x, &y, &w, &h, &aspectRadioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h, &aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:188
// index:2
// void fitInView(const class QGraphicsItem *, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_2(item unsafe.Pointer, aspectRadioMode int) {
	// 2: (, const QGraphicsItem * item, Qt::AspectRatioMode aspectRadioMode), (item, &aspectRadioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEPK13QGraphicsItemN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.cthis, item, &aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// void render(class QPainter *, const class QRectF &, const class QRect &, Qt::AspectRatioMode)
func (this *QGraphicsView) Render(painter unsafe.Pointer, target unsafe.Pointer, source unsafe.Pointer, aspectRatioMode int) {
	// 0: (, QPainter * painter, const QRectF & target, const QRect & source, Qt::AspectRatioMode aspectRatioMode), (painter, target, source, &aspectRatioMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.cthis, painter, target, source, &aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:194
// index:0
// QList<QGraphicsItem *> items()
func (this *QGraphicsView) Items() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:195
// index:1
// QList<QGraphicsItem *> items(const class QPoint &)
func (this *QGraphicsView) Items_1(pos unsafe.Pointer) {
	// 1: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:196
// index:2
// inline
// QList<QGraphicsItem *> items(int, int)
func (this *QGraphicsView) Items_2(x int, y int) {
	// 2: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:197
// index:3
// QList<QGraphicsItem *> items(const class QRect &, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_3(rect unsafe.Pointer, mode int) {
	// 3: (, const QRect & rect, Qt::ItemSelectionMode mode), (rect, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK5QRectN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, rect, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:198
// index:4
// inline
// QList<QGraphicsItem *> items(int, int, int, int, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_4(x int, y int, w int, h int, mode int) {
	// 4: (, int x, int y, int w, int h, Qt::ItemSelectionMode mode), (&x, &y, &w, &h, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsEiiiiN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:199
// index:5
// QList<QGraphicsItem *> items(const class QPolygon &, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_5(polygon unsafe.Pointer, mode int) {
	// 5: (, const QPolygon & polygon, Qt::ItemSelectionMode mode), (polygon, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK8QPolygonN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, polygon, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:200
// index:6
// QList<QGraphicsItem *> items(const class QPainterPath &, Qt::ItemSelectionMode)
func (this *QGraphicsView) Items_6(path unsafe.Pointer, mode int) {
	// 6: (, const QPainterPath & path, Qt::ItemSelectionMode mode), (path, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5itemsERK12QPainterPathN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, path, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:201
// index:0
// QGraphicsItem * itemAt(const class QPoint &)
func (this *QGraphicsView) ItemAt(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:202
// index:1
// inline
// QGraphicsItem * itemAt(int, int)
func (this *QGraphicsView) ItemAt_1(x int, y int) {
	// 1: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:204
// index:0
// QPointF mapToScene(const class QPoint &)
func (this *QGraphicsView) MapToScene(point unsafe.Pointer) {
	// 0: (, const QPoint & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:205
// index:1
// QPolygonF mapToScene(const class QRect &)
func (this *QGraphicsView) MapToScene_1(rect unsafe.Pointer) {
	// 1: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:206
// index:2
// QPolygonF mapToScene(const class QPolygon &)
func (this *QGraphicsView) MapToScene_2(polygon unsafe.Pointer) {
	// 2: (, const QPolygon & polygon), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK8QPolygon", ffiqt.FFI_TYPE_VOID, this.cthis, polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:207
// index:3
// QPainterPath mapToScene(const class QPainterPath &)
func (this *QGraphicsView) MapToScene_3(path unsafe.Pointer) {
	// 3: (, const QPainterPath & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:212
// index:4
// inline
// QPointF mapToScene(int, int)
func (this *QGraphicsView) MapToScene_4(x int, y int) {
	// 4: (, int x, int y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:213
// index:5
// inline
// QPolygonF mapToScene(int, int, int, int)
func (this *QGraphicsView) MapToScene_5(x int, y int, w int, h int) {
	// 5: (, int x, int y, int w, int h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:208
// index:0
// QPoint mapFromScene(const class QPointF &)
func (this *QGraphicsView) MapFromScene(point unsafe.Pointer) {
	// 0: (, const QPointF & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:209
// index:1
// QPolygon mapFromScene(const class QRectF &)
func (this *QGraphicsView) MapFromScene_1(rect unsafe.Pointer) {
	// 1: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:210
// index:2
// QPolygon mapFromScene(const class QPolygonF &)
func (this *QGraphicsView) MapFromScene_2(polygon unsafe.Pointer) {
	// 2: (, const QPolygonF & polygon), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.cthis, polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:211
// index:3
// QPainterPath mapFromScene(const class QPainterPath &)
func (this *QGraphicsView) MapFromScene_3(path unsafe.Pointer) {
	// 3: (, const QPainterPath & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:214
// index:4
// inline
// QPoint mapFromScene(qreal, qreal)
func (this *QGraphicsView) MapFromScene_4(x float64, y float64) {
	// 4: (, qreal x, qreal y), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:215
// index:5
// inline
// QPolygon mapFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) MapFromScene_5(x float64, y float64, w float64, h float64) {
	// 5: (, qreal x, qreal y, qreal w, qreal h), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:217
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsView) InputMethodQuery(query int) {
	// 0: (, Qt::InputMethodQuery query), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:219
// index:0
// QBrush backgroundBrush()
func (this *QGraphicsView) BackgroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView15backgroundBrushEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:220
// index:0
// void setBackgroundBrush(const class QBrush &)
func (this *QGraphicsView) SetBackgroundBrush(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18setBackgroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:222
// index:0
// QBrush foregroundBrush()
func (this *QGraphicsView) ForegroundBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView15foregroundBrushEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:223
// index:0
// void setForegroundBrush(const class QBrush &)
func (this *QGraphicsView) SetForegroundBrush(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18setForegroundBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:228
// index:0
// void updateSceneRect(const class QRectF &)
func (this *QGraphicsView) UpdateSceneRect(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15updateSceneRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:232
// index:0
// void rubberBandChanged(class QRect, class QPointF, class QPointF)
func (this *QGraphicsView) RubberBandChanged(viewportRect unsafe.Pointer, fromScenePoint unsafe.Pointer, toScenePoint unsafe.Pointer) {
	// 0: (, QRect viewportRect, QPointF fromScenePoint, QPointF toScenePoint), (viewportRect, fromScenePoint, toScenePoint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView17rubberBandChangedE5QRect7QPointFS1_", ffiqt.FFI_TYPE_VOID, this.cthis, viewportRect, fromScenePoint, toScenePoint)
	gopp.ErrPrint(err, rv)
}

//  body block end
