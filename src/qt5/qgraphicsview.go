package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qgraphicsview.h
// dst-file: /src/widgets/qgraphicsview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGraphicsView::setBackgroundBrush(const QBrush & brush);
extern void C_ZN13QGraphicsView18setBackgroundBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  Qt::Alignment QGraphicsView::alignment();
extern void C_ZNK13QGraphicsView9alignmentEv(void* qthis); // 4
  // proto:  void QGraphicsView::resetCachedContent();
extern void C_ZN13QGraphicsView18resetCachedContentEv(void* qthis); // 4
  // proto:  QPainter::RenderHints QGraphicsView::renderHints();
extern void C_ZNK13QGraphicsView11renderHintsEv(void* qthis); // 4
  // proto:  void QGraphicsView::setMatrix(const QMatrix & matrix, bool combine);
extern void C_ZN13QGraphicsView9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QGraphicsView::shear(qreal sh, qreal sv);
extern void C_ZN13QGraphicsView5shearEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QGraphicsView::resetTransform();
extern void C_ZN13QGraphicsView14resetTransformEv(void* qthis); // 4
  // proto:  void QGraphicsView::rotate(qreal angle);
extern void C_ZN13QGraphicsView6rotateEd(void* qthis, double arg0); // 4
  // proto:  const QMetaObject * QGraphicsView::metaObject();
extern void C_ZNK13QGraphicsView10metaObjectEv(void* qthis); // 4
  // proto:  QGraphicsItem * QGraphicsView::itemAt(int x, int y);
extern void C_ZNK13QGraphicsView6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QGraphicsItem * QGraphicsView::itemAt(const QPoint & pos);
extern void C_ZNK13QGraphicsView6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  OptimizationFlags QGraphicsView::optimizationFlags();
extern void C_ZNK13QGraphicsView17optimizationFlagsEv(void* qthis); // 4
  // proto:  void QGraphicsView::updateSceneRect(const QRectF & rect);
extern void C_ZN13QGraphicsView15updateSceneRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsView::setTransform(const QTransform & matrix, bool combine);
extern void C_ZN13QGraphicsView12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QGraphicsView::ViewportAnchor QGraphicsView::resizeAnchor();
extern void C_ZNK13QGraphicsView12resizeAnchorEv(void* qthis); // 4
  // proto:  void QGraphicsView::scale(qreal sx, qreal sy);
extern void C_ZN13QGraphicsView5scaleEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QMatrix QGraphicsView::matrix();
extern void* C_ZNK13QGraphicsView6matrixEv(void* qthis); // 4
  // proto:  void QGraphicsView::setScene(QGraphicsScene * scene);
extern void C_ZN13QGraphicsView8setSceneEP14QGraphicsScene(void* qthis, void* arg0); // 4
  // proto:  QSize QGraphicsView::sizeHint();
extern void* C_ZNK13QGraphicsView8sizeHintEv(void* qthis); // 4
  // proto:  QList<QGraphicsItem *> QGraphicsView::items(int x, int y);
extern void C_ZNK13QGraphicsView5itemsEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QList<QGraphicsItem *> QGraphicsView::items(const QPoint & pos);
extern void C_ZNK13QGraphicsView5itemsERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QList<QGraphicsItem *> QGraphicsView::items();
extern void C_ZNK13QGraphicsView5itemsEv(void* qthis); // 4
  // proto:  Qt::ItemSelectionMode QGraphicsView::rubberBandSelectionMode();
extern void C_ZNK13QGraphicsView23rubberBandSelectionModeEv(void* qthis); // 4
  // proto:  void QGraphicsView::setSceneRect(qreal x, qreal y, qreal w, qreal h);
extern void C_ZN13QGraphicsView12setSceneRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  void QGraphicsView::setSceneRect(const QRectF & rect);
extern void C_ZN13QGraphicsView12setSceneRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPolygon QGraphicsView::mapFromScene(const QPolygonF & polygon);
extern void* C_ZNK13QGraphicsView12mapFromSceneERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QGraphicsView::mapFromScene(const QPainterPath & path);
extern void* C_ZNK13QGraphicsView12mapFromSceneERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPoint QGraphicsView::mapFromScene(qreal x, qreal y);
extern void* C_ZNK13QGraphicsView12mapFromSceneEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  QPolygon QGraphicsView::mapFromScene(const QRectF & rect);
extern void* C_ZNK13QGraphicsView12mapFromSceneERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPoint QGraphicsView::mapFromScene(const QPointF & point);
extern void* C_ZNK13QGraphicsView12mapFromSceneERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPolygon QGraphicsView::mapFromScene(qreal x, qreal y, qreal w, qreal h);
extern void* C_ZNK13QGraphicsView12mapFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 2
  // proto:  QGraphicsView::DragMode QGraphicsView::dragMode();
extern void C_ZNK13QGraphicsView8dragModeEv(void* qthis); // 4
  // proto:  QGraphicsScene * QGraphicsView::scene();
extern void C_ZNK13QGraphicsView5sceneEv(void* qthis); // 4
  // proto:  void QGraphicsView::~QGraphicsView();
extern void C_ZN13QGraphicsViewD2Ev(void* qthis); // 4
  // proto:  QGraphicsView::ViewportAnchor QGraphicsView::transformationAnchor();
extern void C_ZNK13QGraphicsView20transformationAnchorEv(void* qthis); // 4
  // proto:  void QGraphicsView::QGraphicsView(QWidget * parent);
extern void* C_ZN13QGraphicsViewC2EP7QWidget(void* arg0); // 3
  // proto:  void QGraphicsView::QGraphicsView(QGraphicsScene * scene, QWidget * parent);
extern void* C_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget(void* arg0, void* arg1); // 3
  // proto:  QRect QGraphicsView::rubberBandRect();
extern void* C_ZNK13QGraphicsView14rubberBandRectEv(void* qthis); // 4
  // proto:  QTransform QGraphicsView::transform();
extern void* C_ZNK13QGraphicsView9transformEv(void* qthis); // 4
  // proto:  QBrush QGraphicsView::foregroundBrush();
extern void* C_ZNK13QGraphicsView15foregroundBrushEv(void* qthis); // 4
  // proto:  void QGraphicsView::setForegroundBrush(const QBrush & brush);
extern void C_ZN13QGraphicsView18setForegroundBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsView::ensureVisible(const QRectF & rect, int xmargin, int ymargin);
extern void C_ZN13QGraphicsView13ensureVisibleERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QGraphicsView::ensureVisible(qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin);
extern void C_ZN13QGraphicsView13ensureVisibleEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4, int32_t arg5); // 2
  // proto:  void QGraphicsView::ensureVisible(const QGraphicsItem * item, int xmargin, int ymargin);
extern void C_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QGraphicsView::setInteractive(bool allowed);
extern void C_ZN13QGraphicsView14setInteractiveEb(void* qthis, bool arg0); // 4
  // proto:  bool QGraphicsView::isTransformed();
extern bool C_ZNK13QGraphicsView13isTransformedEv(void* qthis); // 4
  // proto:  bool QGraphicsView::isInteractive();
extern bool C_ZNK13QGraphicsView13isInteractiveEv(void* qthis); // 4
  // proto:  QBrush QGraphicsView::backgroundBrush();
extern void* C_ZNK13QGraphicsView15backgroundBrushEv(void* qthis); // 4
  // proto:  QPainterPath QGraphicsView::mapToScene(const QPainterPath & path);
extern void* C_ZNK13QGraphicsView10mapToSceneERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QPointF QGraphicsView::mapToScene(int x, int y);
extern void* C_ZNK13QGraphicsView10mapToSceneEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QPolygonF QGraphicsView::mapToScene(const QRect & rect);
extern void* C_ZNK13QGraphicsView10mapToSceneERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsView::mapToScene(const QPolygon & polygon);
extern void* C_ZNK13QGraphicsView10mapToSceneERK8QPolygon(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QGraphicsView::mapToScene(int x, int y, int w, int h);
extern void* C_ZNK13QGraphicsView10mapToSceneEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  QPointF QGraphicsView::mapToScene(const QPoint & point);
extern void* C_ZNK13QGraphicsView10mapToSceneERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsView::resetMatrix();
extern void C_ZN13QGraphicsView11resetMatrixEv(void* qthis); // 4
  // proto:  void QGraphicsView::centerOn(const QPointF & pos);
extern void C_ZN13QGraphicsView8centerOnERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsView::centerOn(const QGraphicsItem * item);
extern void C_ZN13QGraphicsView8centerOnEPK13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsView::centerOn(qreal x, qreal y);
extern void C_ZN13QGraphicsView8centerOnEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsView::translate(qreal dx, qreal dy);
extern void C_ZN13QGraphicsView9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QGraphicsView::ViewportUpdateMode QGraphicsView::viewportUpdateMode();
extern void C_ZNK13QGraphicsView18viewportUpdateModeEv(void* qthis); // 4
  // proto:  QRectF QGraphicsView::sceneRect();
extern void* C_ZNK13QGraphicsView9sceneRectEv(void* qthis); // 4
  // proto:  CacheMode QGraphicsView::cacheMode();
extern void C_ZNK13QGraphicsView9cacheModeEv(void* qthis); // 4
  // proto:  QTransform QGraphicsView::viewportTransform();
extern void* C_ZNK13QGraphicsView17viewportTransformEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QGraphicsView)=1
type QGraphicsView struct {
  /*qbase*/ QAbstractScrollArea;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _rubberBandChanged QGraphicsView_rubberBandChanged_signal;
}

// setBackgroundBrush(const class QBrush &)
func (this *QGraphicsView) Setbackgroundbrush(args ...interface{}) () {
  // setBackgroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView18setBackgroundBrushERK6QBrush
    // invoke: void setBackgroundBrush(const class QBrush &)
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView18setBackgroundBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setBackgroundBrush", args)
  }

  return
}

// alignment()
func (this *QGraphicsView) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK13QGraphicsView9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "alignment", args)
  }

  return
}

// resetCachedContent()
func (this *QGraphicsView) Resetcachedcontent(args ...interface{}) () {
  // resetCachedContent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView18resetCachedContentEv
    // invoke: void resetCachedContent()
    C.C_ZN13QGraphicsView18resetCachedContentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetCachedContent", args)
  }

  return
}

// renderHints()
func (this *QGraphicsView) Renderhints(args ...interface{}) () {
  // renderHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView11renderHintsEv
    // invoke: QPainter::RenderHints renderHints()
    C.C_ZNK13QGraphicsView11renderHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "renderHints", args)
  }

  return
}

// setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsView) Setmatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView9setMatrixERK7QMatrixb
    // invoke: void setMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsView9setMatrixERK7QMatrixb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setMatrix", args)
  }

  return
}

// shear(qreal, qreal)
func (this *QGraphicsView) Shear(args ...interface{}) () {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView5shearEdd
    // invoke: void shear(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsView5shearEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "shear", args)
  }

  return
}

// resetTransform()
func (this *QGraphicsView) Resettransform(args ...interface{}) () {
  // resetTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView14resetTransformEv
    // invoke: void resetTransform()
    C.C_ZN13QGraphicsView14resetTransformEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetTransform", args)
  }

  return
}

// rotate(qreal)
func (this *QGraphicsView) Rotate(args ...interface{}) () {
  // rotate(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView6rotateEd
    // invoke: void rotate(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView6rotateEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "rotate", args)
  }

  return
}

// metaObject()
func (this *QGraphicsView) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QGraphicsView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "metaObject", args)
  }

  return
}

// itemAt(int, int)
func (this *QGraphicsView) Itemat(args ...interface{}) () {
  // itemAt(int, int)
  // itemAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView6itemAtEii
    // invoke: QGraphicsItem * itemAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK13QGraphicsView6itemAtEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK13QGraphicsView6itemAtERK6QPoint
    // invoke: QGraphicsItem * itemAt(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QGraphicsView6itemAtERK6QPoint(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "itemAt", args)
  }

  return
}

// optimizationFlags()
func (this *QGraphicsView) Optimizationflags(args ...interface{}) () {
  // optimizationFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView17optimizationFlagsEv
    // invoke: OptimizationFlags optimizationFlags()
    C.C_ZNK13QGraphicsView17optimizationFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "optimizationFlags", args)
  }

  return
}

// updateSceneRect(const class QRectF &)
func (this *QGraphicsView) Updatescenerect(args ...interface{}) () {
  // updateSceneRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView15updateSceneRectERK6QRectF
    // invoke: void updateSceneRect(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView15updateSceneRectERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "updateSceneRect", args)
  }

  return
}

// setTransform(const class QTransform &, _Bool)
func (this *QGraphicsView) Settransform(args ...interface{}) () {
  // setTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView12setTransformERK10QTransformb
    // invoke: void setTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsView12setTransformERK10QTransformb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setTransform", args)
  }

  return
}

// resizeAnchor()
func (this *QGraphicsView) Resizeanchor(args ...interface{}) () {
  // resizeAnchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView12resizeAnchorEv
    // invoke: QGraphicsView::ViewportAnchor resizeAnchor()
    C.C_ZNK13QGraphicsView12resizeAnchorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resizeAnchor", args)
  }

  return
}

// scale(qreal, qreal)
func (this *QGraphicsView) Scale(args ...interface{}) () {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView5scaleEdd
    // invoke: void scale(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsView5scaleEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "scale", args)
  }

  return
}

// matrix()
func (this *QGraphicsView) Matrix(args ...interface{}) (ret interface{}) {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView6matrixEv
    // invoke: QMatrix matrix()
    var ret0 = C.C_ZNK13QGraphicsView6matrixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "matrix", args)
  }

  return
}

// setScene(class QGraphicsScene *)
func (this *QGraphicsView) Setscene(args ...interface{}) () {
  // setScene(class QGraphicsScene *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsScene{}) // "QGraphicsScene *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView8setSceneEP14QGraphicsScene
    // invoke: void setScene(class QGraphicsScene *)
    var arg0 = args[0].(*QGraphicsScene).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView8setSceneEP14QGraphicsScene(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setScene", args)
  }

  return
}

// sizeHint()
func (this *QGraphicsView) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK13QGraphicsView8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "sizeHint", args)
  }

  return
}

// items(int, int)
func (this *QGraphicsView) Items(args ...interface{}) () {
  // items(int, int)
  // items(const class QPoint &)
  // items()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView5itemsEii
    // invoke: QList<QGraphicsItem *> items(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK13QGraphicsView5itemsEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK13QGraphicsView5itemsERK6QPoint
    // invoke: QList<QGraphicsItem *> items(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QGraphicsView5itemsERK6QPoint(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZNK13QGraphicsView5itemsEv
    // invoke: QList<QGraphicsItem *> items()
    C.C_ZNK13QGraphicsView5itemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "items", args)
  }

  return
}

// rubberBandSelectionMode()
func (this *QGraphicsView) Rubberbandselectionmode(args ...interface{}) () {
  // rubberBandSelectionMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView23rubberBandSelectionModeEv
    // invoke: Qt::ItemSelectionMode rubberBandSelectionMode()
    C.C_ZNK13QGraphicsView23rubberBandSelectionModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "rubberBandSelectionMode", args)
  }

  return
}

// setSceneRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsView) Setscenerect(args ...interface{}) () {
  // setSceneRect(qreal, qreal, qreal, qreal)
  // setSceneRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView12setSceneRectEdddd
    // invoke: void setSceneRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN13QGraphicsView12setSceneRectEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN13QGraphicsView12setSceneRectERK6QRectF
    // invoke: void setSceneRect(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView12setSceneRectERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setSceneRect", args)
  }

  return
}

// mapFromScene(const class QPolygonF &)
func (this *QGraphicsView) Mapfromscene(args ...interface{}) (ret interface{}) {
  // mapFromScene(const class QPolygonF &)
  // mapFromScene(const class QPainterPath &)
  // mapFromScene(qreal, qreal)
  // mapFromScene(const class QRectF &)
  // mapFromScene(const class QPointF &)
  // mapFromScene(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK9QPolygonF
    // invoke: QPolygon mapFromScene(const class QPolygonF &)
    var arg0 = args[0].(*QPolygonF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView12mapFromSceneERK9QPolygonF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK12QPainterPath
    // invoke: QPainterPath mapFromScene(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView12mapFromSceneERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsView12mapFromSceneEdd
    // invoke: QPoint mapFromScene(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsView12mapFromSceneEdd(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK6QRectF
    // invoke: QPolygon mapFromScene(const class QRectF &)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView12mapFromSceneERK6QRectF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK7QPointF
    // invoke: QPoint mapFromScene(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView12mapFromSceneERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsView12mapFromSceneEdddd
    // invoke: QPolygon mapFromScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsView12mapFromSceneEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygon{}) // "QPolygon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "mapFromScene", args)
  }

  return
}

// dragMode()
func (this *QGraphicsView) Dragmode(args ...interface{}) () {
  // dragMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView8dragModeEv
    // invoke: QGraphicsView::DragMode dragMode()
    C.C_ZNK13QGraphicsView8dragModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "dragMode", args)
  }

  return
}

// scene()
func (this *QGraphicsView) Scene(args ...interface{}) () {
  // scene()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView5sceneEv
    // invoke: QGraphicsScene * scene()
    C.C_ZNK13QGraphicsView5sceneEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "scene", args)
  }

  return
}

// ~QGraphicsView()
func (this *QGraphicsView) Freeqgraphicsview(args ...interface{}) () {
  // ~QGraphicsView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsViewD0Ev
    // invoke: void ~QGraphicsView()
    C.C_ZN13QGraphicsViewD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "~QGraphicsView", args)
  }

  return
}

// transformationAnchor()
func (this *QGraphicsView) Transformationanchor(args ...interface{}) () {
  // transformationAnchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView20transformationAnchorEv
    // invoke: QGraphicsView::ViewportAnchor transformationAnchor()
    C.C_ZNK13QGraphicsView20transformationAnchorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "transformationAnchor", args)
  }

  return
}

// QGraphicsView(class QWidget *)
func NewQGraphicsView(args ...interface{}) *QGraphicsView {
  // QGraphicsView(class QWidget *)
  // QGraphicsView(class QGraphicsScene *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsScene{}) // "QGraphicsScene *"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsViewC1EP7QWidget
    // invoke: void QGraphicsView(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QGraphicsViewC2EP7QWidget(arg0)
    return &QGraphicsView{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QGraphicsViewC1EP14QGraphicsSceneP7QWidget
    // invoke: void QGraphicsView(class QGraphicsScene *, class QWidget *)
    var arg0 = args[0].(*QGraphicsScene).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QGraphicsViewC2EP14QGraphicsSceneP7QWidget(arg0, arg1)
    return &QGraphicsView{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsView", "QGraphicsView", args)
  }

  return nil // QGraphicsView{Qclsinst:qthis}
}

// rubberBandRect()
func (this *QGraphicsView) Rubberbandrect(args ...interface{}) (ret interface{}) {
  // rubberBandRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView14rubberBandRectEv
    // invoke: QRect rubberBandRect()
    var ret0 = C.C_ZNK13QGraphicsView14rubberBandRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "rubberBandRect", args)
  }

  return
}

// transform()
func (this *QGraphicsView) Transform(args ...interface{}) (ret interface{}) {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView9transformEv
    // invoke: QTransform transform()
    var ret0 = C.C_ZNK13QGraphicsView9transformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "transform", args)
  }

  return
}

// foregroundBrush()
func (this *QGraphicsView) Foregroundbrush(args ...interface{}) (ret interface{}) {
  // foregroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView15foregroundBrushEv
    // invoke: QBrush foregroundBrush()
    var ret0 = C.C_ZNK13QGraphicsView15foregroundBrushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "foregroundBrush", args)
  }

  return
}

// setForegroundBrush(const class QBrush &)
func (this *QGraphicsView) Setforegroundbrush(args ...interface{}) () {
  // setForegroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView18setForegroundBrushERK6QBrush
    // invoke: void setForegroundBrush(const class QBrush &)
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView18setForegroundBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setForegroundBrush", args)
  }

  return
}

// ensureVisible(const class QRectF &, int, int)
func (this *QGraphicsView) Ensurevisible(args ...interface{}) () {
  // ensureVisible(const class QRectF &, int, int)
  // ensureVisible(qreal, qreal, qreal, qreal, int, int)
  // ensureVisible(const class QGraphicsItem *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView13ensureVisibleERK6QRectFii
    // invoke: void ensureVisible(const class QRectF &, int, int)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN13QGraphicsView13ensureVisibleERK6QRectFii(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN13QGraphicsView13ensureVisibleEddddii
    // invoke: void ensureVisible(qreal, qreal, qreal, qreal, int, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN13QGraphicsView13ensureVisibleEddddii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 2:
    // invoke: _ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii
    // invoke: void ensureVisible(const class QGraphicsItem *, int, int)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsView", "ensureVisible", args)
  }

  return
}

// setInteractive(_Bool)
func (this *QGraphicsView) Setinteractive(args ...interface{}) () {
  // setInteractive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView14setInteractiveEb
    // invoke: void setInteractive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView14setInteractiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setInteractive", args)
  }

  return
}

// isTransformed()
func (this *QGraphicsView) Istransformed(args ...interface{}) (ret interface{}) {
  // isTransformed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView13isTransformedEv
    // invoke: bool isTransformed()
    var ret0 = C.C_ZNK13QGraphicsView13isTransformedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "isTransformed", args)
  }

  return
}

// isInteractive()
func (this *QGraphicsView) Isinteractive(args ...interface{}) (ret interface{}) {
  // isInteractive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView13isInteractiveEv
    // invoke: bool isInteractive()
    var ret0 = C.C_ZNK13QGraphicsView13isInteractiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "isInteractive", args)
  }

  return
}

// backgroundBrush()
func (this *QGraphicsView) Backgroundbrush(args ...interface{}) (ret interface{}) {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView15backgroundBrushEv
    // invoke: QBrush backgroundBrush()
    var ret0 = C.C_ZNK13QGraphicsView15backgroundBrushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "backgroundBrush", args)
  }

  return
}

// mapToScene(const class QPainterPath &)
func (this *QGraphicsView) Maptoscene(args ...interface{}) (ret interface{}) {
  // mapToScene(const class QPainterPath &)
  // mapToScene(int, int)
  // mapToScene(const class QRect &)
  // mapToScene(const class QPolygon &)
  // mapToScene(int, int, int, int)
  // mapToScene(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK12QPainterPath
    // invoke: QPainterPath mapToScene(const class QPainterPath &)
    var arg0 = args[0].(*QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView10mapToSceneERK12QPainterPath(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QGraphicsView10mapToSceneEii
    // invoke: QPointF mapToScene(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QGraphicsView10mapToSceneEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK5QRect
    // invoke: QPolygonF mapToScene(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView10mapToSceneERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK8QPolygon
    // invoke: QPolygonF mapToScene(const class QPolygon &)
    var arg0 = args[0].(*QPolygon).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView10mapToSceneERK8QPolygon(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK13QGraphicsView10mapToSceneEiiii
    // invoke: QPolygonF mapToScene(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QGraphicsView10mapToSceneEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPolygonF{}) // "QPolygonF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK6QPoint
    // invoke: QPointF mapToScene(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGraphicsView10mapToSceneERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "mapToScene", args)
  }

  return
}

// resetMatrix()
func (this *QGraphicsView) Resetmatrix(args ...interface{}) () {
  // resetMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView11resetMatrixEv
    // invoke: void resetMatrix()
    C.C_ZN13QGraphicsView11resetMatrixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetMatrix", args)
  }

  return
}

// centerOn(const class QPointF &)
func (this *QGraphicsView) Centeron(args ...interface{}) () {
  // centerOn(const class QPointF &)
  // centerOn(const class QGraphicsItem *)
  // centerOn(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView8centerOnERK7QPointF
    // invoke: void centerOn(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView8centerOnERK7QPointF(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN13QGraphicsView8centerOnEPK13QGraphicsItem
    // invoke: void centerOn(const class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGraphicsView8centerOnEPK13QGraphicsItem(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN13QGraphicsView8centerOnEdd
    // invoke: void centerOn(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsView8centerOnEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "centerOn", args)
  }

  return
}

// translate(qreal, qreal)
func (this *QGraphicsView) Translate(args ...interface{}) () {
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGraphicsView9translateEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "translate", args)
  }

  return
}

// viewportUpdateMode()
func (this *QGraphicsView) Viewportupdatemode(args ...interface{}) () {
  // viewportUpdateMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView18viewportUpdateModeEv
    // invoke: QGraphicsView::ViewportUpdateMode viewportUpdateMode()
    C.C_ZNK13QGraphicsView18viewportUpdateModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "viewportUpdateMode", args)
  }

  return
}

// sceneRect()
func (this *QGraphicsView) Scenerect(args ...interface{}) (ret interface{}) {
  // sceneRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView9sceneRectEv
    // invoke: QRectF sceneRect()
    var ret0 = C.C_ZNK13QGraphicsView9sceneRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "sceneRect", args)
  }

  return
}

// cacheMode()
func (this *QGraphicsView) Cachemode(args ...interface{}) () {
  // cacheMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView9cacheModeEv
    // invoke: CacheMode cacheMode()
    C.C_ZNK13QGraphicsView9cacheModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "cacheMode", args)
  }

  return
}

// viewportTransform()
func (this *QGraphicsView) Viewporttransform(args ...interface{}) (ret interface{}) {
  // viewportTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView17viewportTransformEv
    // invoke: QTransform viewportTransform()
    var ret0 = C.C_ZNK13QGraphicsView17viewportTransformEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsView", "viewportTransform", args)
  }

  return
}

// <= body block end

