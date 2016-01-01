package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicsview.h
// dst-file: /src/widgets/qgraphicsview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QGraphicsView)=1
type QGraphicsView struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
//  _rubberBandChanged QGraphicsView_rubberBandChanged_signal;
}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "scale", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QGraphicsView10mapToSceneEii
  case 2:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK6QPoint
  case 3:
    // invoke: _ZNK13QGraphicsView10mapToSceneEiiii
  case 4:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK8QPolygon
  case 5:
    // invoke: _ZNK13QGraphicsView10mapToSceneERK12QPainterPath
  default:
    qtrt.ErrorResolve("QGraphicsView", "mapToScene", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QGraphicsView12mapFromSceneEdd
  case 2:
    // invoke: _ZNK13QGraphicsView12mapFromSceneEdddd
  case 3:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK9QPolygonF
  case 4:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK7QPointF
  case 5:
    // invoke: _ZNK13QGraphicsView12mapFromSceneERK12QPainterPath
  default:
    qtrt.ErrorResolve("QGraphicsView", "mapFromScene", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "translate", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "metaObject", args)
  }

}


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
  case 1:
    // invoke: _ZN13QGraphicsView12setSceneRectERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsView", "setSceneRect", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "rubberBandRect", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "setMatrix", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "isInteractive", args)
  }

}


func NewQGraphicsView(args ...interface{}) QGraphicsView {
  return QGraphicsView{}
}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "setBackgroundBrush", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "isTransformed", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QGraphicsView6itemAtERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsView", "itemAt", args)
  }

}


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
  case 1:
    // invoke: _ZN13QGraphicsView8centerOnEPK13QGraphicsItem
  case 2:
    // invoke: _ZN13QGraphicsView8centerOnEdd
  default:
    qtrt.ErrorResolve("QGraphicsView", "centerOn", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "setForegroundBrush", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "shear", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "foregroundBrush", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "updateSceneRect", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "setInteractive", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "matrix", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "viewportTransform", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetCachedContent", args)
  }

}


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
  case 1:
    // invoke: _ZN13QGraphicsView13ensureVisibleERK6QRectFii
  case 2:
    // invoke: _ZN13QGraphicsView13ensureVisibleEddddii
  default:
    qtrt.ErrorResolve("QGraphicsView", "ensureVisible", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "sceneRect", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "scene", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "sizeHint", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "backgroundBrush", args)
  }

}


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
    // invoke: _ZNK13QGraphicsView5itemsEiiiiN2Qt17ItemSelectionModeE
  case 1:
    // invoke: _ZNK13QGraphicsView5itemsERK5QRectN2Qt17ItemSelectionModeE
  case 2:
    // invoke: _ZNK13QGraphicsView5itemsERK8QPolygonN2Qt17ItemSelectionModeE
  case 3:
    // invoke: _ZNK13QGraphicsView5itemsEv
  case 4:
    // invoke: _ZNK13QGraphicsView5itemsEii
  case 5:
    // invoke: _ZNK13QGraphicsView5itemsERK12QPainterPathN2Qt17ItemSelectionModeE
  case 6:
    // invoke: _ZNK13QGraphicsView5itemsERK6QPoint
  default:
    qtrt.ErrorResolve("QGraphicsView", "items", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "transform", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "rotate", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "setTransform", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetMatrix", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "resetTransform", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsView", "setScene", args)
  }

}

// <= body block end

