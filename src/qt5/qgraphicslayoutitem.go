package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qgraphicslayoutitem.h
// dst-file: /src/widgets/qgraphicslayoutitem.go
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
// class sizeof(QGraphicsLayoutItem)=1
type QGraphicsLayoutItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsLayoutItem) setSizePolicy(args ...interface{}) () {
  // setSizePolicy(const class QSizePolicy &)
  // setSizePolicy(class QSizePolicy::Policy, class QSizePolicy::Policy, class QSizePolicy::ControlType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizePolicy{}) // "const QSizePolicy &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QSizePolicy::Policy"
  vtys[1][1] = qtrt.Int32Ty(false) // "QSizePolicy::Policy"
  vtys[1][2] = qtrt.Int32Ty(false) // "QSizePolicy::ControlType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setSizePolicy", args)
 }

}


func (this *QGraphicsLayoutItem) parentLayoutItem(args ...interface{}) () {
  // parentLayoutItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem16parentLayoutItemEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "parentLayoutItem", args)
 }

}


func (this *QGraphicsLayoutItem) minimumWidth(args ...interface{}) () {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12minimumWidthEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumWidth", args)
 }

}


func (this *QGraphicsLayoutItem) graphicsItem(args ...interface{}) () {
  // graphicsItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12graphicsItemEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "graphicsItem", args)
 }

}


func (this *QGraphicsLayoutItem) preferredWidth(args ...interface{}) () {
  // preferredWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem14preferredWidthEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredWidth", args)
 }

}


func (this *QGraphicsLayoutItem) ownedByLayout(args ...interface{}) () {
  // ownedByLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13ownedByLayoutEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "ownedByLayout", args)
 }

}


func (this *QGraphicsLayoutItem) preferredSize(args ...interface{}) () {
  // preferredSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13preferredSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredSize", args)
 }

}


func (this *QGraphicsLayoutItem) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem8geometryEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "geometry", args)
 }

}


func NewQGraphicsLayoutItem(args ...interface{})() {
}


func (this *QGraphicsLayoutItem) minimumHeight(args ...interface{}) () {
  // minimumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13minimumHeightEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumHeight", args)
 }

}


func (this *QGraphicsLayoutItem) preferredHeight(args ...interface{}) () {
  // preferredHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem15preferredHeightEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "preferredHeight", args)
 }

}


func (this *QGraphicsLayoutItem) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem11maximumSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumSize", args)
 }

}


func (this *QGraphicsLayoutItem) sizePolicy(args ...interface{}) () {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem10sizePolicyEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "sizePolicy", args)
 }

}


func (this *QGraphicsLayoutItem) maximumHeight(args ...interface{}) () {
  // maximumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem13maximumHeightEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumHeight", args)
 }

}


func (this *QGraphicsLayoutItem) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem11setGeometryERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setGeometry", args)
 }

}


func (this *QGraphicsLayoutItem) setPreferredWidth(args ...interface{}) () {
  // setPreferredWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem17setPreferredWidthEd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredWidth", args)
 }

}


func (this *QGraphicsLayoutItem) setMaximumSize(args ...interface{}) () {
  // setMaximumSize(const class QSizeF &)
  // setMaximumSize(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMaximumSizeEdd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumSize", args)
 }

}


func (this *QGraphicsLayoutItem) maximumWidth(args ...interface{}) () {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12maximumWidthEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "maximumWidth", args)
 }

}


func (this *QGraphicsLayoutItem) setMinimumSize(args ...interface{}) () {
  // setMinimumSize(qreal, qreal)
  // setMinimumSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem14setMinimumSizeEdd
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumSize", args)
 }

}


func (this *QGraphicsLayoutItem) setMaximumHeight(args ...interface{}) () {
  // setMaximumHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem16setMaximumHeightEd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumHeight", args)
 }

}


func (this *QGraphicsLayoutItem) setPreferredSize(args ...interface{}) () {
  // setPreferredSize(const class QSizeF &)
  // setPreferredSize(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF
  case 1:
    // invoke: _ZN19QGraphicsLayoutItem16setPreferredSizeEdd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredSize", args)
 }

}


func (this *QGraphicsLayoutItem) getContentsMargins(args ...interface{}) () {
  // getContentsMargins(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "getContentsMargins", args)
 }

}


func (this *QGraphicsLayoutItem) setParentLayoutItem(args ...interface{}) () {
  // setParentLayoutItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setParentLayoutItem", args)
 }

}


func (this *QGraphicsLayoutItem) setMinimumWidth(args ...interface{}) () {
  // setMinimumWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem15setMinimumWidthEd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumWidth", args)
 }

}


func (this *QGraphicsLayoutItem) setMaximumWidth(args ...interface{}) () {
  // setMaximumWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem15setMaximumWidthEd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMaximumWidth", args)
 }

}


func (this *QGraphicsLayoutItem) updateGeometry(args ...interface{}) () {
  // updateGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem14updateGeometryEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "updateGeometry", args)
 }

}


func (this *QGraphicsLayoutItem) setPreferredHeight(args ...interface{}) () {
  // setPreferredHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem18setPreferredHeightEd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setPreferredHeight", args)
 }

}


func (this *QGraphicsLayoutItem) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem11minimumSizeEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "minimumSize", args)
 }

}


func (this *QGraphicsLayoutItem) contentsRect(args ...interface{}) () {
  // contentsRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem12contentsRectEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "contentsRect", args)
 }

}


func (this *QGraphicsLayoutItem) isLayout(args ...interface{}) () {
  // isLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsLayoutItem8isLayoutEv
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "isLayout", args)
 }

}


func (this *QGraphicsLayoutItem) FreeQGraphicsLayoutItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "~QGraphicsLayoutItem", args)
 }

}


func (this *QGraphicsLayoutItem) setMinimumHeight(args ...interface{}) () {
  // setMinimumHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsLayoutItem16setMinimumHeightEd
  default:
    qtrt.ErrorResolve("QGraphicsLayoutItem", "setMinimumHeight", args)
 }

}

// <= body block end

