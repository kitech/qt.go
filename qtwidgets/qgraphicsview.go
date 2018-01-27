package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsview.h
// #include <qgraphicsview.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QGraphicsView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQGraphicsViewFromPointer(cthis unsafe.Pointer) *QGraphicsView {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QGraphicsView{bcthis0}
}
func (*QGraphicsView) NewFromPointer(cthis unsafe.Pointer) *QGraphicsView {
	return NewQGraphicsViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:61
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:117
// index:0
// Public
// void QGraphicsView(QWidget *)
func NewQGraphicsView(parent *QWidget /*777 QWidget **/) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:118
// index:1
// Public
// void QGraphicsView(QGraphicsScene *, QWidget *)
func NewQGraphicsView_1(scene *QGraphicsScene /*777 QGraphicsScene **/, parent *QWidget /*777 QWidget **/) *QGraphicsView {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = scene.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:119
// index:0
// Public virtual
// void ~QGraphicsView()
func DeleteQGraphicsView(*QGraphicsView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:121
// index:0
// Public virtual
// QSize sizeHint()
func (this *QGraphicsView) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:123
// index:0
// Public
// QPainter::RenderHints renderHints()
func (this *QGraphicsView) RenderHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView11renderHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:124
// index:0
// Public
// void setRenderHint(QPainter::RenderHint, bool)
func (this *QGraphicsView) SetRenderHint(hint int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13setRenderHintEN8QPainter10RenderHintEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hint, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:127
// index:0
// Public
// Qt::Alignment alignment()
func (this *QGraphicsView) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:130
// index:0
// Public
// QGraphicsView::ViewportAnchor transformationAnchor()
func (this *QGraphicsView) TransformationAnchor() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView20transformationAnchorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:131
// index:0
// Public
// void setTransformationAnchor(QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetTransformationAnchor(anchor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView23setTransformationAnchorENS_14ViewportAnchorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:133
// index:0
// Public
// QGraphicsView::ViewportAnchor resizeAnchor()
func (this *QGraphicsView) ResizeAnchor() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12resizeAnchorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:134
// index:0
// Public
// void setResizeAnchor(QGraphicsView::ViewportAnchor)
func (this *QGraphicsView) SetResizeAnchor(anchor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15setResizeAnchorENS_14ViewportAnchorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:136
// index:0
// Public
// QGraphicsView::ViewportUpdateMode viewportUpdateMode()
func (this *QGraphicsView) ViewportUpdateMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView18viewportUpdateModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:137
// index:0
// Public
// void setViewportUpdateMode(QGraphicsView::ViewportUpdateMode)
func (this *QGraphicsView) SetViewportUpdateMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView21setViewportUpdateModeENS_18ViewportUpdateModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:139
// index:0
// Public
// QGraphicsView::OptimizationFlags optimizationFlags()
func (this *QGraphicsView) OptimizationFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView17optimizationFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:140
// index:0
// Public
// void setOptimizationFlag(QGraphicsView::OptimizationFlag, bool)
func (this *QGraphicsView) SetOptimizationFlag(flag int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView19setOptimizationFlagENS_16OptimizationFlagEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:141
// index:0
// Public
// void setOptimizationFlags(QGraphicsView::OptimizationFlags)
func (this *QGraphicsView) SetOptimizationFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView20setOptimizationFlagsE6QFlagsINS_16OptimizationFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:143
// index:0
// Public
// QGraphicsView::DragMode dragMode()
func (this *QGraphicsView) DragMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView8dragModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:144
// index:0
// Public
// void setDragMode(QGraphicsView::DragMode)
func (this *QGraphicsView) SetDragMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11setDragModeENS_8DragModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:147
// index:0
// Public
// Qt::ItemSelectionMode rubberBandSelectionMode()
func (this *QGraphicsView) RubberBandSelectionMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView23rubberBandSelectionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:148
// index:0
// Public
// void setRubberBandSelectionMode(Qt::ItemSelectionMode)
func (this *QGraphicsView) SetRubberBandSelectionMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView26setRubberBandSelectionModeEN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:149
// index:0
// Public
// QRect rubberBandRect()
func (this *QGraphicsView) RubberBandRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView14rubberBandRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:152
// index:0
// Public
// QGraphicsView::CacheMode cacheMode()
func (this *QGraphicsView) CacheMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9cacheModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:154
// index:0
// Public
// void resetCachedContent()
func (this *QGraphicsView) ResetCachedContent() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18resetCachedContentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:156
// index:0
// Public
// bool isInteractive()
func (this *QGraphicsView) IsInteractive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView13isInteractiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:157
// index:0
// Public
// void setInteractive(bool)
func (this *QGraphicsView) SetInteractive(allowed bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14setInteractiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), allowed)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:159
// index:0
// Public
// QGraphicsScene * scene()
func (this *QGraphicsView) Scene() *QGraphicsScene /*777 QGraphicsScene **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView5sceneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:160
// index:0
// Public
// void setScene(QGraphicsScene *)
func (this *QGraphicsView) SetScene(scene *QGraphicsScene /*777 QGraphicsScene **/) {
	var convArg0 = scene.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8setSceneEP14QGraphicsScene", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:162
// index:0
// Public
// QRectF sceneRect()
func (this *QGraphicsView) SceneRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9sceneRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:163
// index:0
// Public
// void setSceneRect(const QRectF &)
func (this *QGraphicsView) SetSceneRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:164
// index:1
// Public inline
// void setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) SetSceneRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setSceneRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:166
// index:0
// Public
// QMatrix matrix()
func (this *QGraphicsView) Matrix() *qtgui.QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6matrixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:167
// index:0
// Public
// void setMatrix(const QMatrix &, bool)
func (this *QGraphicsView) SetMatrix(matrix *qtgui.QMatrix, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:168
// index:0
// Public
// void resetMatrix()
func (this *QGraphicsView) ResetMatrix() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11resetMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:169
// index:0
// Public
// QTransform transform()
func (this *QGraphicsView) Transform() *qtgui.QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView9transformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:170
// index:0
// Public
// QTransform viewportTransform()
func (this *QGraphicsView) ViewportTransform() *qtgui.QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView17viewportTransformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:171
// index:0
// Public
// bool isTransformed()
func (this *QGraphicsView) IsTransformed() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView13isTransformedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:172
// index:0
// Public
// void setTransform(const QTransform &, bool)
func (this *QGraphicsView) SetTransform(matrix *qtgui.QTransform, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12setTransformERK10QTransformb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:173
// index:0
// Public
// void resetTransform()
func (this *QGraphicsView) ResetTransform() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14resetTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:174
// index:0
// Public
// void rotate(qreal)
func (this *QGraphicsView) Rotate(angle float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView6rotateEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:175
// index:0
// Public
// void scale(qreal, qreal)
func (this *QGraphicsView) Scale(sx float64, sy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5scaleEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:176
// index:0
// Public
// void shear(qreal, qreal)
func (this *QGraphicsView) Shear(sh float64, sv float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5shearEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:177
// index:0
// Public
// void translate(qreal, qreal)
func (this *QGraphicsView) Translate(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9translateEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:179
// index:0
// Public
// void centerOn(const QPointF &)
func (this *QGraphicsView) CenterOn(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:180
// index:1
// Public inline
// void centerOn(qreal, qreal)
func (this *QGraphicsView) CenterOn_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:181
// index:2
// Public
// void centerOn(const QGraphicsItem *)
func (this *QGraphicsView) CenterOn_2(item *QGraphicsItem /*777 const QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView8centerOnEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:182
// index:0
// Public
// void ensureVisible(const QRectF &, int, int)
func (this *QGraphicsView) EnsureVisible(rect *qtcore.QRectF, xmargin int, ymargin int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleERK6QRectFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:183
// index:1
// Public inline
// void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsView) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEddddii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:184
// index:2
// Public
// void ensureVisible(const QGraphicsItem *, int, int)
func (this *QGraphicsView) EnsureVisible_2(item *QGraphicsItem /*777 const QGraphicsItem **/, xmargin int, ymargin int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:185
// index:0
// Public
// void fitInView(const QRectF &, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView(rect *qtcore.QRectF, aspectRadioMode int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewERK6QRectFN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:186
// index:1
// Public inline
// void fitInView(qreal, qreal, qreal, qreal, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_1(x float64, y float64, w float64, h float64, aspectRadioMode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEddddN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:188
// index:2
// Public
// void fitInView(const QGraphicsItem *, Qt::AspectRatioMode)
func (this *QGraphicsView) FitInView_2(item *QGraphicsItem /*777 const QGraphicsItem **/, aspectRadioMode int) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9fitInViewEPK13QGraphicsItemN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectRadioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:191
// index:0
// Public
// void render(QPainter *, const QRectF &, const QRect &, Qt::AspectRatioMode)
func (this *QGraphicsView) Render(painter *qtgui.QPainter /*777 QPainter **/, target *qtcore.QRectF, source *qtcore.QRect, aspectRatioMode int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = target.GetCthis()
	var convArg2 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView6renderEP8QPainterRK6QRectFRK5QRectN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, aspectRatioMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:201
// index:0
// Public
// QGraphicsItem * itemAt(const QPoint &)
func (this *QGraphicsView) ItemAt(pos *qtcore.QPoint) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:202
// index:1
// Public inline
// QGraphicsItem * itemAt(int, int)
func (this *QGraphicsView) ItemAt_1(x int, y int) *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView6itemAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:204
// index:0
// Public
// QPointF mapToScene(const QPoint &)
func (this *QGraphicsView) MapToScene(point *qtcore.QPoint) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:205
// index:1
// Public
// QPolygonF mapToScene(const QRect &)
func (this *QGraphicsView) MapToScene_1(rect *qtcore.QRect) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:206
// index:2
// Public
// QPolygonF mapToScene(const QPolygon &)
func (this *QGraphicsView) MapToScene_2(polygon *qtgui.QPolygon) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK8QPolygon", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:207
// index:3
// Public
// QPainterPath mapToScene(const QPainterPath &)
func (this *QGraphicsView) MapToScene_3(path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:212
// index:4
// Public inline
// QPointF mapToScene(int, int)
func (this *QGraphicsView) MapToScene_4(x int, y int) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:213
// index:5
// Public inline
// QPolygonF mapToScene(int, int, int, int)
func (this *QGraphicsView) MapToScene_5(x int, y int, w int, h int) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView10mapToSceneEiiii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:208
// index:0
// Public
// QPoint mapFromScene(const QPointF &)
func (this *QGraphicsView) MapFromScene(point *qtcore.QPointF) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:209
// index:1
// Public
// QPolygon mapFromScene(const QRectF &)
func (this *QGraphicsView) MapFromScene_1(rect *qtcore.QRectF) *qtgui.QPolygon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:210
// index:2
// Public
// QPolygon mapFromScene(const QPolygonF &)
func (this *QGraphicsView) MapFromScene_2(polygon *qtgui.QPolygonF) *qtgui.QPolygon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:211
// index:3
// Public
// QPainterPath mapFromScene(const QPainterPath &)
func (this *QGraphicsView) MapFromScene_3(path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:214
// index:4
// Public inline
// QPoint mapFromScene(qreal, qreal)
func (this *QGraphicsView) MapFromScene_4(x float64, y float64) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:215
// index:5
// Public inline
// QPolygon mapFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) MapFromScene_5(x float64, y float64, w float64, h float64) *qtgui.QPolygon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView12mapFromSceneEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:217
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsView) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:219
// index:0
// Public
// QBrush backgroundBrush()
func (this *QGraphicsView) BackgroundBrush() *qtgui.QBrush /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView15backgroundBrushEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:220
// index:0
// Public
// void setBackgroundBrush(const QBrush &)
func (this *QGraphicsView) SetBackgroundBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18setBackgroundBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:222
// index:0
// Public
// QBrush foregroundBrush()
func (this *QGraphicsView) ForegroundBrush() *qtgui.QBrush /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsView15foregroundBrushEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:223
// index:0
// Public
// void setForegroundBrush(const QBrush &)
func (this *QGraphicsView) SetForegroundBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18setForegroundBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:228
// index:0
// Public
// void updateSceneRect(const QRectF &)
func (this *QGraphicsView) UpdateSceneRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15updateSceneRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:232
// index:0
// Public
// void rubberBandChanged(QRect, QPointF, QPointF)
func (this *QGraphicsView) RubberBandChanged(viewportRect *qtcore.QRect /*123*/, fromScenePoint *qtcore.QPointF /*123*/, toScenePoint *qtcore.QPointF /*123*/) {
	var convArg0 = viewportRect.GetCthis()
	var convArg1 = fromScenePoint.GetCthis()
	var convArg2 = toScenePoint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView17rubberBandChangedE5QRect7QPointFS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:236
// index:0
// Protected virtual
// void setupViewport(QWidget *)
func (this *QGraphicsView) SetupViewport(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13setupViewportEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:240
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QGraphicsView) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:241
// index:0
// Protected virtual
// bool viewportEvent(QEvent *)
func (this *QGraphicsView) ViewportEvent(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:244
// index:0
// Protected virtual
// void contextMenuEvent(QContextMenuEvent *)
func (this *QGraphicsView) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:247
// index:0
// Protected virtual
// void dragEnterEvent(QDragEnterEvent *)
func (this *QGraphicsView) DragEnterEvent(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:248
// index:0
// Protected virtual
// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QGraphicsView) DragLeaveEvent(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:249
// index:0
// Protected virtual
// void dragMoveEvent(QDragMoveEvent *)
func (this *QGraphicsView) DragMoveEvent(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:250
// index:0
// Protected virtual
// void dropEvent(QDropEvent *)
func (this *QGraphicsView) DropEvent(event *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:252
// index:0
// Protected virtual
// void focusInEvent(QFocusEvent *)
func (this *QGraphicsView) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:253
// index:0
// Protected virtual
// bool focusNextPrevChild(bool)
func (this *QGraphicsView) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:254
// index:0
// Protected virtual
// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsView) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:255
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QGraphicsView) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:256
// index:0
// Protected virtual
// void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsView) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:257
// index:0
// Protected virtual
// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QGraphicsView) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:258
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QGraphicsView) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:259
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QGraphicsView) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:260
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QGraphicsView) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:262
// index:0
// Protected virtual
// void wheelEvent(QWheelEvent *)
func (this *QGraphicsView) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:264
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QGraphicsView) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:265
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QGraphicsView) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:266
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QGraphicsView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:267
// index:0
// Protected virtual
// void showEvent(QShowEvent *)
func (this *QGraphicsView) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:268
// index:0
// Protected virtual
// void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsView) InputMethodEvent(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:270
// index:0
// Protected virtual
// void drawBackground(QPainter *, const QRectF &)
func (this *QGraphicsView) DrawBackground(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:271
// index:0
// Protected virtual
// void drawForeground(QPainter *, const QRectF &)
func (this *QGraphicsView) DrawForeground(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRectF) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsview.h:272
// index:0
// Protected virtual
// void drawItems(QPainter *, int, QGraphicsItem **, const QStyleOptionGraphicsItem *)
func (this *QGraphicsView) DrawItems(painter *qtgui.QPainter /*777 QPainter **/, numItems int, items []interface{}, options []interface{}) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numItems, items, options)
	gopp.ErrPrint(err, rv)
}

type QGraphicsView__ViewportAnchor = int

const QGraphicsView__NoAnchor QGraphicsView__ViewportAnchor = 0
const QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = 1
const QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = 2

type QGraphicsView__CacheModeFlag = int

const QGraphicsView__CacheNone QGraphicsView__CacheModeFlag = 0
const QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = 1

type QGraphicsView__DragMode = int

const QGraphicsView__NoDrag QGraphicsView__DragMode = 0
const QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = 1
const QGraphicsView__RubberBandDrag QGraphicsView__DragMode = 2

type QGraphicsView__ViewportUpdateMode = int

const QGraphicsView__FullViewportUpdate QGraphicsView__ViewportUpdateMode = 0
const QGraphicsView__MinimalViewportUpdate QGraphicsView__ViewportUpdateMode = 1
const QGraphicsView__SmartViewportUpdate QGraphicsView__ViewportUpdateMode = 2
const QGraphicsView__NoViewportUpdate QGraphicsView__ViewportUpdateMode = 3
const QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = 4

type QGraphicsView__OptimizationFlag = int

const QGraphicsView__DontClipPainter QGraphicsView__OptimizationFlag = 1
const QGraphicsView__DontSavePainterState QGraphicsView__OptimizationFlag = 2
const QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = 4
const QGraphicsView__IndirectPainting QGraphicsView__OptimizationFlag = 8

//  body block end
