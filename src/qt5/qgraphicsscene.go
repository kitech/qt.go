package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicsscene.h
// dst-file: /src/widgets/qgraphicsscene.go
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
// class sizeof(QGraphicsScene)=1
type QGraphicsScene struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _changed QGraphicsScene_changed_signal;
//  _sceneRectChanged QGraphicsScene_sceneRectChanged_signal;
//  _selectionChanged QGraphicsScene_selectionChanged_signal;
//  _focusItemChanged QGraphicsScene_focusItemChanged_signal;
}


func (this *QGraphicsScene) setForegroundBrush(args ...interface{}) () {
  // setForegroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setForegroundBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setForegroundBrush", args)
  }

}


func (this *QGraphicsScene) setSceneRect(args ...interface{}) () {
  // setSceneRect(const class QRectF &)
  // setSceneRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene12setSceneRectERK6QRectF
  case 1:
    // invoke: _ZN14QGraphicsScene12setSceneRectEdddd
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSceneRect", args)
  }

}


func (this *QGraphicsScene) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8isActiveEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isActive", args)
  }

}


func (this *QGraphicsScene) hasFocus(args ...interface{}) () {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8hasFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "hasFocus", args)
  }

}


func (this *QGraphicsScene) itemsBoundingRect(args ...interface{}) () {
  // itemsBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17itemsBoundingRectEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemsBoundingRect", args)
  }

}


func (this *QGraphicsScene) sendEvent(args ...interface{}) () {
  // sendEvent(class QGraphicsItem *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sendEvent", args)
  }

}


func (this *QGraphicsScene) minimumRenderSize(args ...interface{}) () {
  // minimumRenderSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17minimumRenderSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "minimumRenderSize", args)
  }

}


func (this *QGraphicsScene) selectionArea(args ...interface{}) () {
  // selectionArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectionAreaEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectionArea", args)
  }

}


func (this *QGraphicsScene) update(args ...interface{}) () {
  // update(const class QRectF &)
  // update(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene6updateERK6QRectF
  case 1:
    // invoke: _ZN14QGraphicsScene6updateEdddd
  default:
    qtrt.ErrorResolve("QGraphicsScene", "update", args)
  }

}


func (this *QGraphicsScene) addPolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPolygon", args)
  }

}


func (this *QGraphicsScene) addLine(args ...interface{}) () {
  // addLine(const class QLineF &, const class QPen &)
  // addLine(qreal, qreal, qreal, qreal, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addLineERK6QLineFRK4QPen
  case 1:
    // invoke: _ZN14QGraphicsScene7addLineEddddRK4QPen
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addLine", args)
  }

}


func (this *QGraphicsScene) palette(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene7paletteEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "palette", args)
  }

}


func (this *QGraphicsScene) isSortCacheEnabled(args ...interface{}) () {
  // isSortCacheEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene18isSortCacheEnabledEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isSortCacheEnabled", args)
  }

}


func NewQGraphicsScene(args ...interface{}) QGraphicsScene {
  return QGraphicsScene{}
}


func (this *QGraphicsScene) clearFocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10clearFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearFocus", args)
  }

}


func (this *QGraphicsScene) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "metaObject", args)
  }

}


func (this *QGraphicsScene) addSimpleText(args ...interface{}) () {
  // addSimpleText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addSimpleText", args)
  }

}


func (this *QGraphicsScene) setBspTreeDepth(args ...interface{}) () {
  // setBspTreeDepth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setBspTreeDepthEi
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBspTreeDepth", args)
  }

}


func (this *QGraphicsScene) sceneRect(args ...interface{}) () {
  // sceneRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9sceneRectEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sceneRect", args)
  }

}


func (this *QGraphicsScene) activeWindow(args ...interface{}) () {
  // activeWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12activeWindowEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activeWindow", args)
  }

}


func (this *QGraphicsScene) backgroundBrush(args ...interface{}) () {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15backgroundBrushEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "backgroundBrush", args)
  }

}


func (this *QGraphicsScene) itemAt(args ...interface{}) () {
  // itemAt(qreal, qreal, const class QTransform &)
  // itemAt(const class QPointF &, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6itemAtEddRK10QTransform
  case 1:
    // invoke: _ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemAt", args)
  }

}


func (this *QGraphicsScene) advance(args ...interface{}) () {
  // advance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7advanceEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "advance", args)
  }

}


func (this *QGraphicsScene) setStickyFocus(args ...interface{}) () {
  // setStickyFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setStickyFocusEb
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStickyFocus", args)
  }

}


func (this *QGraphicsScene) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectedItemsEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectedItems", args)
  }

}


func (this *QGraphicsScene) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene5clearEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clear", args)
  }

}


func (this *QGraphicsScene) setActivePanel(args ...interface{}) () {
  // setActivePanel(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActivePanel", args)
  }

}


func (this *QGraphicsScene) addPixmap(args ...interface{}) () {
  // addPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9addPixmapERK7QPixmap
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPixmap", args)
  }

}


func (this *QGraphicsScene) foregroundBrush(args ...interface{}) () {
  // foregroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15foregroundBrushEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "foregroundBrush", args)
  }

}


func (this *QGraphicsScene) views(args ...interface{}) () {
  // views()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5viewsEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "views", args)
  }

}


func (this *QGraphicsScene) FreeQGraphicsScene(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsScene", "~QGraphicsScene", args)
  }

}


func (this *QGraphicsScene) addRect(args ...interface{}) () {
  // addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  // addRect(const class QRectF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush
  case 1:
    // invoke: _ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addRect", args)
  }

}


func (this *QGraphicsScene) bspTreeDepth(args ...interface{}) () {
  // bspTreeDepth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12bspTreeDepthEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "bspTreeDepth", args)
  }

}


func (this *QGraphicsScene) setStyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene8setStyleEP6QStyle
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStyle", args)
  }

}


func (this *QGraphicsScene) setPalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10setPaletteERK8QPalette
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setPalette", args)
  }

}


func (this *QGraphicsScene) setMinimumRenderSize(args ...interface{}) () {
  // setMinimumRenderSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene20setMinimumRenderSizeEd
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setMinimumRenderSize", args)
  }

}


func (this *QGraphicsScene) mouseGrabberItem(args ...interface{}) () {
  // mouseGrabberItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene16mouseGrabberItemEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "mouseGrabberItem", args)
  }

}


func (this *QGraphicsScene) addEllipse(args ...interface{}) () {
  // addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
  // addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush
  case 1:
    // invoke: _ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addEllipse", args)
  }

}


func (this *QGraphicsScene) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6heightEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "height", args)
  }

}


func (this *QGraphicsScene) setSelectionArea(args ...interface{}) () {
  // setSelectionArea(const class QPainterPath &, Qt::ItemSelectionMode, const class QTransform &)
  // setSelectionArea(const class QPainterPath &, const class QTransform &)
  // setSelectionArea(const class QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionOperation"
  vtys[2][2] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[2][3] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt17ItemSelectionModeERK10QTransform
  case 1:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform
  case 2:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathN2Qt22ItemSelectionOperationENS3_17ItemSelectionModeERK10QTransform
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSelectionArea", args)
  }

}


func (this *QGraphicsScene) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene4fontEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "font", args)
  }

}


func (this *QGraphicsScene) clearSelection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14clearSelectionEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearSelection", args)
  }

}


func (this *QGraphicsScene) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10removeItemEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsScene", "removeItem", args)
  }

}


func (this *QGraphicsScene) setActiveWindow(args ...interface{}) () {
  // setActiveWindow(class QGraphicsWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsWidget{}) // "QGraphicsWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActiveWindow", args)
  }

}


func (this *QGraphicsScene) focusItem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9focusItemEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "focusItem", args)
  }

}


func (this *QGraphicsScene) addText(args ...interface{}) () {
  // addText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addTextERK7QStringRK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addText", args)
  }

}


func (this *QGraphicsScene) setSortCacheEnabled(args ...interface{}) () {
  // setSortCacheEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene19setSortCacheEnabledEb
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSortCacheEnabled", args)
  }

}


func (this *QGraphicsScene) destroyItemGroup(args ...interface{}) () {
  // destroyItemGroup(class QGraphicsItemGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItemGroup{}) // "QGraphicsItemGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup
  default:
    qtrt.ErrorResolve("QGraphicsScene", "destroyItemGroup", args)
  }

}


func (this *QGraphicsScene) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5widthEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "width", args)
  }

}


func (this *QGraphicsScene) addItem(args ...interface{}) () {
  // addItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addItemEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addItem", args)
  }

}


func (this *QGraphicsScene) setBackgroundBrush(args ...interface{}) () {
  // setBackgroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setBackgroundBrushERK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBackgroundBrush", args)
  }

}


func (this *QGraphicsScene) activePanel(args ...interface{}) () {
  // activePanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11activePanelEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activePanel", args)
  }

}


func (this *QGraphicsScene) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5styleEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "style", args)
  }

}


func (this *QGraphicsScene) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setFont", args)
  }

}


func (this *QGraphicsScene) addPath(args ...interface{}) () {
  // addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPath", args)
  }

}


func (this *QGraphicsScene) stickyFocus(args ...interface{}) () {
  // stickyFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11stickyFocusEv
  default:
    qtrt.ErrorResolve("QGraphicsScene", "stickyFocus", args)
  }

}

// <= body block end

