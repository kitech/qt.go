package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGraphicsView::scale(qreal sx, qreal sy);
extern void _ZN13QGraphicsView5scaleEdd(void* qthis, double arg0, double arg1);
  // proto:  QPolygonF QGraphicsView::mapToScene(const QRect & rect);
extern void _ZNK13QGraphicsView10mapToSceneERK5QRect(void* qthis, void* arg0);
  // proto:  QPolygon QGraphicsView::mapFromScene(const QRectF & rect);
extern void _ZNK13QGraphicsView12mapFromSceneERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsView::translate(qreal dx, qreal dy);
extern void _ZN13QGraphicsView9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  QPointF QGraphicsView::mapToScene(int x, int y);
extern void demth_ZNK13QGraphicsView10mapToSceneEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  const QMetaObject * QGraphicsView::metaObject();
extern void _ZNK13QGraphicsView10metaObjectEv(void* qthis);
  // proto:  void QGraphicsView::setSceneRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN13QGraphicsView12setSceneRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QRect QGraphicsView::rubberBandRect();
extern void _ZNK13QGraphicsView14rubberBandRectEv(void* qthis);
  // proto:  void QGraphicsView::setMatrix(const QMatrix & matrix, bool combine);
extern void _ZN13QGraphicsView9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1);
  // proto:  bool QGraphicsView::isInteractive();
extern void _ZNK13QGraphicsView13isInteractiveEv(void* qthis);
  // proto:  void QGraphicsView::QGraphicsView(const QGraphicsView & );
extern void* dector_ZN13QGraphicsViewC1ERKS_(void* arg0);
extern void _ZN13QGraphicsViewC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsView::setBackgroundBrush(const QBrush & brush);
extern void _ZN13QGraphicsView18setBackgroundBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  bool QGraphicsView::isTransformed();
extern void _ZNK13QGraphicsView13isTransformedEv(void* qthis);
  // proto:  QGraphicsItem * QGraphicsView::itemAt(int x, int y);
extern void demth_ZNK13QGraphicsView6itemAtEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QGraphicsView::centerOn(const QPointF & pos);
extern void _ZN13QGraphicsView8centerOnERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsView::setForegroundBrush(const QBrush & brush);
extern void _ZN13QGraphicsView18setForegroundBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  void QGraphicsView::shear(qreal sh, qreal sv);
extern void _ZN13QGraphicsView5shearEdd(void* qthis, double arg0, double arg1);
  // proto:  QBrush QGraphicsView::foregroundBrush();
extern void _ZNK13QGraphicsView15foregroundBrushEv(void* qthis);
  // proto:  QGraphicsItem * QGraphicsView::itemAt(const QPoint & pos);
extern void _ZNK13QGraphicsView6itemAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsView::updateSceneRect(const QRectF & rect);
extern void _ZN13QGraphicsView15updateSceneRectERK6QRectF(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsView::mapToScene(const QPoint & point);
extern void _ZNK13QGraphicsView10mapToSceneERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsView::setInteractive(bool allowed);
extern void _ZN13QGraphicsView14setInteractiveEb(void* qthis, bool arg0);
  // proto:  QMatrix QGraphicsView::matrix();
extern void _ZNK13QGraphicsView6matrixEv(void* qthis);
  // proto:  QPolygonF QGraphicsView::mapToScene(int x, int y, int w, int h);
extern void demth_ZNK13QGraphicsView10mapToSceneEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QGraphicsView::QGraphicsView(QGraphicsScene * scene, QWidget * parent);
extern void* dector_ZN13QGraphicsViewC1EP14QGraphicsSceneP7QWidget(void* arg0, void* arg1);
extern void _ZN13QGraphicsViewC1EP14QGraphicsSceneP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsView::QGraphicsView(QWidget * parent);
extern void* dector_ZN13QGraphicsViewC1EP7QWidget(void* arg0);
extern void _ZN13QGraphicsViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QGraphicsView::centerOn(const QGraphicsItem * item);
extern void _ZN13QGraphicsView8centerOnEPK13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QTransform QGraphicsView::viewportTransform();
extern void _ZNK13QGraphicsView17viewportTransformEv(void* qthis);
  // proto:  void QGraphicsView::resetCachedContent();
extern void _ZN13QGraphicsView18resetCachedContentEv(void* qthis);
  // proto:  QPolygonF QGraphicsView::mapToScene(const QPolygon & polygon);
extern void _ZNK13QGraphicsView10mapToSceneERK8QPolygon(void* qthis, void* arg0);
  // proto:  void QGraphicsView::ensureVisible(const QGraphicsItem * item, int xmargin, int ymargin);
extern void _ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  QRectF QGraphicsView::sceneRect();
extern void _ZNK13QGraphicsView9sceneRectEv(void* qthis);
  // proto:  QGraphicsScene * QGraphicsView::scene();
extern void _ZNK13QGraphicsView5sceneEv(void* qthis);
  // proto:  QSize QGraphicsView::sizeHint();
extern void _ZNK13QGraphicsView8sizeHintEv(void* qthis);
  // proto:  QBrush QGraphicsView::backgroundBrush();
extern void _ZNK13QGraphicsView15backgroundBrushEv(void* qthis);
  // proto:  QPoint QGraphicsView::mapFromScene(qreal x, qreal y);
extern void demth_ZNK13QGraphicsView12mapFromSceneEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsView::ensureVisible(const QRectF & rect, int xmargin, int ymargin);
extern void _ZN13QGraphicsView13ensureVisibleERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  QList<QGraphicsItem *> QGraphicsView::items();
extern void _ZNK13QGraphicsView5itemsEv(void* qthis);
  // proto:  QTransform QGraphicsView::transform();
extern void _ZNK13QGraphicsView9transformEv(void* qthis);
  // proto:  QList<QGraphicsItem *> QGraphicsView::items(int x, int y);
extern void demth_ZNK13QGraphicsView5itemsEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QGraphicsView::centerOn(qreal x, qreal y);
extern void demth_ZN13QGraphicsView8centerOnEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsView::ensureVisible(qreal x, qreal y, qreal w, qreal h, int xmargin, int ymargin);
extern void demth_ZN13QGraphicsView13ensureVisibleEddddii(void* qthis, double arg0, double arg1, double arg2, double arg3, int32_t arg4, int32_t arg5);
  // proto:  void QGraphicsView::rotate(qreal angle);
extern void _ZN13QGraphicsView6rotateEd(void* qthis, double arg0);
  // proto:  QPolygon QGraphicsView::mapFromScene(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZNK13QGraphicsView12mapFromSceneEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsView::setTransform(const QTransform & matrix, bool combine);
extern void _ZN13QGraphicsView12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1);
  // proto:  void QGraphicsView::setSceneRect(const QRectF & rect);
extern void _ZN13QGraphicsView12setSceneRectERK6QRectF(void* qthis, void* arg0);
  // proto:  QPolygon QGraphicsView::mapFromScene(const QPolygonF & polygon);
extern void _ZNK13QGraphicsView12mapFromSceneERK9QPolygonF(void* qthis, void* arg0);
  // proto:  QPoint QGraphicsView::mapFromScene(const QPointF & point);
extern void _ZNK13QGraphicsView12mapFromSceneERK7QPointF(void* qthis, void* arg0);
  // proto:  QPainterPath QGraphicsView::mapFromScene(const QPainterPath & path);
extern void _ZNK13QGraphicsView12mapFromSceneERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QGraphicsView::~QGraphicsView();
extern void _ZN13QGraphicsViewD0Ev(void* qthis);
  // proto:  QList<QGraphicsItem *> QGraphicsView::items(const QPoint & pos);
extern void _ZNK13QGraphicsView5itemsERK6QPoint(void* qthis, void* arg0);
  // proto:  void QGraphicsView::resetMatrix();
extern void _ZN13QGraphicsView11resetMatrixEv(void* qthis);
  // proto:  void QGraphicsView::resetTransform();
extern void _ZN13QGraphicsView14resetTransformEv(void* qthis);
  // proto:  QPainterPath QGraphicsView::mapToScene(const QPainterPath & path);
extern void _ZNK13QGraphicsView10mapToSceneERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QGraphicsView::setScene(QGraphicsScene * scene);
extern void _ZN13QGraphicsView8setSceneEP14QGraphicsScene(void* qthis, void* arg0);
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _rubberBandChanged QGraphicsView_rubberBandChanged_signal;
}

  // proto:  void QGraphicsView::scale(qreal sx, qreal sy);
func (this *QGraphicsView) scale(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN13QGraphicsView5scaleEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "scale", args)
  }

}

  // proto:  QPolygonF QGraphicsView::mapToScene(const QRect & rect);
func (this *QGraphicsView) mapToScene(args ...interface{}) () {
  // mapToScene(const class QRect &)
  // mapToScene(int, int)
  // mapToScene(const class QPoint &)
  // mapToScene(int, int, int, int)
  // mapToScene(const class QPolygon &)
  // mapToScene(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK5QRect
    // invoke: QPolygonF mapToScene(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView10mapToSceneERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QGraphicsView10mapToSceneEii
    // invoke: QPointF mapToScene(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK13QGraphicsView10mapToSceneEii(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK6QPoint
    // invoke: QPointF mapToScene(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView10mapToSceneERK6QPoint(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK13QGraphicsView10mapToSceneEiiii
    // invoke: QPolygonF mapToScene(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZNK13QGraphicsView10mapToSceneEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK8QPolygon
    // invoke: QPolygonF mapToScene(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView10mapToSceneERK8QPolygon(this.qclsinst, arg0)
  case 5:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK12QPainterPath
    // invoke: QPainterPath mapToScene(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView10mapToSceneERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "mapToScene", args)
  }

}

  // proto:  QPolygon QGraphicsView::mapFromScene(const QRectF & rect);
func (this *QGraphicsView) mapFromScene(args ...interface{}) () {
  // mapFromScene(const class QRectF &)
  // mapFromScene(qreal, qreal)
  // mapFromScene(qreal, qreal, qreal, qreal)
  // mapFromScene(const class QPolygonF &)
  // mapFromScene(const class QPointF &)
  // mapFromScene(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK6QRectF
    // invoke: QPolygon mapFromScene(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView12mapFromSceneERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QGraphicsView12mapFromSceneEdd
    // invoke: QPoint mapFromScene(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZNK13QGraphicsView12mapFromSceneEdd(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK13QGraphicsView12mapFromSceneEdddd
    // invoke: QPolygon mapFromScene(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.demth_ZNK13QGraphicsView12mapFromSceneEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK9QPolygonF
    // invoke: QPolygon mapFromScene(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView12mapFromSceneERK9QPolygonF(this.qclsinst, arg0)
  case 4:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK7QPointF
    // invoke: QPoint mapFromScene(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView12mapFromSceneERK7QPointF(this.qclsinst, arg0)
  case 5:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK12QPainterPath
    // invoke: QPainterPath mapFromScene(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView12mapFromSceneERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "mapFromScene", args)
  }

}

  // proto:  void QGraphicsView::translate(qreal dx, qreal dy);
func (this *QGraphicsView) translate(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN13QGraphicsView9translateEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "translate", args)
  }

}

  // proto:  const QMetaObject * QGraphicsView::metaObject();
func (this *QGraphicsView) metaObject(args ...interface{}) () {
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
    C._ZNK13QGraphicsView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "metaObject", args)
  }

}

  // proto:  void QGraphicsView::setSceneRect(qreal x, qreal y, qreal w, qreal h);
func (this *QGraphicsView) setSceneRect(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.demth_ZN13QGraphicsView12setSceneRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN13QGraphicsView12setSceneRectERK6QRectF
    // invoke: void setSceneRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView12setSceneRectERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setSceneRect", args)
  }

}

  // proto:  QRect QGraphicsView::rubberBandRect();
func (this *QGraphicsView) rubberBandRect(args ...interface{}) () {
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
    C._ZNK13QGraphicsView14rubberBandRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "rubberBandRect", args)
  }

}

  // proto:  void QGraphicsView::setMatrix(const QMatrix & matrix, bool combine);
func (this *QGraphicsView) setMatrix(args ...interface{}) () {
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
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN13QGraphicsView9setMatrixERK7QMatrixb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setMatrix", args)
  }

}

  // proto:  bool QGraphicsView::isInteractive();
func (this *QGraphicsView) isInteractive(args ...interface{}) () {
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
    C._ZNK13QGraphicsView13isInteractiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "isInteractive", args)
  }

}

  // proto:  void QGraphicsView::QGraphicsView(const QGraphicsView & );
func NewQGraphicsView(args ...interface{}) QGraphicsView {
  return QGraphicsView{}
}

  // proto:  void QGraphicsView::setBackgroundBrush(const QBrush & brush);
func (this *QGraphicsView) setBackgroundBrush(args ...interface{}) () {
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
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView18setBackgroundBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setBackgroundBrush", args)
  }

}

  // proto:  bool QGraphicsView::isTransformed();
func (this *QGraphicsView) isTransformed(args ...interface{}) () {
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
    C._ZNK13QGraphicsView13isTransformedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "isTransformed", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsView::itemAt(int x, int y);
func (this *QGraphicsView) itemAt(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK13QGraphicsView6itemAtEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK13QGraphicsView6itemAtERK6QPoint
    // invoke: QGraphicsItem * itemAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView6itemAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "itemAt", args)
  }

}

  // proto:  void QGraphicsView::centerOn(const QPointF & pos);
func (this *QGraphicsView) centerOn(args ...interface{}) () {
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView8centerOnERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN13QGraphicsView8centerOnEPK13QGraphicsItem
    // invoke: void centerOn(const class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView8centerOnEPK13QGraphicsItem(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN13QGraphicsView8centerOnEdd
    // invoke: void centerOn(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZN13QGraphicsView8centerOnEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "centerOn", args)
  }

}

  // proto:  void QGraphicsView::setForegroundBrush(const QBrush & brush);
func (this *QGraphicsView) setForegroundBrush(args ...interface{}) () {
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
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView18setForegroundBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setForegroundBrush", args)
  }

}

  // proto:  void QGraphicsView::shear(qreal sh, qreal sv);
func (this *QGraphicsView) shear(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN13QGraphicsView5shearEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "shear", args)
  }

}

  // proto:  QBrush QGraphicsView::foregroundBrush();
func (this *QGraphicsView) foregroundBrush(args ...interface{}) () {
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
    C._ZNK13QGraphicsView15foregroundBrushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "foregroundBrush", args)
  }

}

  // proto:  void QGraphicsView::updateSceneRect(const QRectF & rect);
func (this *QGraphicsView) updateSceneRect(args ...interface{}) () {
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
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView15updateSceneRectERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "updateSceneRect", args)
  }

}

  // proto:  void QGraphicsView::setInteractive(bool allowed);
func (this *QGraphicsView) setInteractive(args ...interface{}) () {
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
    C._ZN13QGraphicsView14setInteractiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setInteractive", args)
  }

}

  // proto:  QMatrix QGraphicsView::matrix();
func (this *QGraphicsView) matrix(args ...interface{}) () {
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
    C._ZNK13QGraphicsView6matrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "matrix", args)
  }

}

  // proto:  QTransform QGraphicsView::viewportTransform();
func (this *QGraphicsView) viewportTransform(args ...interface{}) () {
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
    C._ZNK13QGraphicsView17viewportTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "viewportTransform", args)
  }

}

  // proto:  void QGraphicsView::resetCachedContent();
func (this *QGraphicsView) resetCachedContent(args ...interface{}) () {
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
    C._ZN13QGraphicsView18resetCachedContentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetCachedContent", args)
  }

}

  // proto:  void QGraphicsView::ensureVisible(const QGraphicsItem * item, int xmargin, int ymargin);
func (this *QGraphicsView) ensureVisible(args ...interface{}) () {
  // ensureVisible(const class QGraphicsItem *, int, int)
  // ensureVisible(const class QRectF &, int, int)
  // ensureVisible(qreal, qreal, qreal, qreal, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "const QGraphicsItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii
    // invoke: void ensureVisible(const class QGraphicsItem *, int, int)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN13QGraphicsView13ensureVisibleEPK13QGraphicsItemii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN13QGraphicsView13ensureVisibleERK6QRectFii
    // invoke: void ensureVisible(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN13QGraphicsView13ensureVisibleERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN13QGraphicsView13ensureVisibleEddddii
    // invoke: void ensureVisible(qreal, qreal, qreal, qreal, int, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.demth_ZN13QGraphicsView13ensureVisibleEddddii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QGraphicsView", "ensureVisible", args)
  }

}

  // proto:  QRectF QGraphicsView::sceneRect();
func (this *QGraphicsView) sceneRect(args ...interface{}) () {
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
    C._ZNK13QGraphicsView9sceneRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "sceneRect", args)
  }

}

  // proto:  QGraphicsScene * QGraphicsView::scene();
func (this *QGraphicsView) scene(args ...interface{}) () {
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
    C._ZNK13QGraphicsView5sceneEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "scene", args)
  }

}

  // proto:  QSize QGraphicsView::sizeHint();
func (this *QGraphicsView) sizeHint(args ...interface{}) () {
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
    C._ZNK13QGraphicsView8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "sizeHint", args)
  }

}

  // proto:  QBrush QGraphicsView::backgroundBrush();
func (this *QGraphicsView) backgroundBrush(args ...interface{}) () {
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
    C._ZNK13QGraphicsView15backgroundBrushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "backgroundBrush", args)
  }

}

  // proto:  QList<QGraphicsItem *> QGraphicsView::items();
func (this *QGraphicsView) items(args ...interface{}) () {
  // items(int, int, int, int, Qt::ItemSelectionMode)
  // items(const class QRect &, Qt::ItemSelectionMode)
  // items(const class QPolygon &, Qt::ItemSelectionMode)
  // items()
  // items(int, int)
  // items(const class QPainterPath &, Qt::ItemSelectionMode)
  // items(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[2][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[5][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGraphicsView5itemsEv
    // invoke: QList<QGraphicsItem *> items()
    C._ZNK13QGraphicsView5itemsEv(this.qclsinst)
  case 1:
    // invoke: _ZNK13QGraphicsView5itemsEii
    // invoke: QList<QGraphicsItem *> items(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK13QGraphicsView5itemsEii(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK13QGraphicsView5itemsERK6QPoint
    // invoke: QList<QGraphicsItem *> items(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGraphicsView5itemsERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "items", args)
  }

}

  // proto:  QTransform QGraphicsView::transform();
func (this *QGraphicsView) transform(args ...interface{}) () {
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
    C._ZNK13QGraphicsView9transformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "transform", args)
  }

}

  // proto:  void QGraphicsView::rotate(qreal angle);
func (this *QGraphicsView) rotate(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView6rotateEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "rotate", args)
  }

}

  // proto:  void QGraphicsView::setTransform(const QTransform & matrix, bool combine);
func (this *QGraphicsView) setTransform(args ...interface{}) () {
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
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN13QGraphicsView12setTransformERK10QTransformb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setTransform", args)
  }

}

  // proto:  void QGraphicsView::~QGraphicsView();
func (this *QGraphicsView) FreeQGraphicsView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsView", "~QGraphicsView", args)
  }

}

  // proto:  void QGraphicsView::resetMatrix();
func (this *QGraphicsView) resetMatrix(args ...interface{}) () {
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
    C._ZN13QGraphicsView11resetMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetMatrix", args)
  }

}

  // proto:  void QGraphicsView::resetTransform();
func (this *QGraphicsView) resetTransform(args ...interface{}) () {
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
    C._ZN13QGraphicsView14resetTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetTransform", args)
  }

}

  // proto:  void QGraphicsView::setScene(QGraphicsScene * scene);
func (this *QGraphicsView) setScene(args ...interface{}) () {
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
    var arg0 = args[0].(QGraphicsScene).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGraphicsView8setSceneEP14QGraphicsScene(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsView", "setScene", args)
  }

}

// <= body block end

