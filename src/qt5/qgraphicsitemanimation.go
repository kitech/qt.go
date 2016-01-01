package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicsitemanimation.h
// dst-file: /src/widgets/qgraphicsitemanimation.go
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
// class sizeof(QGraphicsItemAnimation)=1
type QGraphicsItemAnimation struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsItemAnimation) setPosAt(args ...interface{}) () {
  // setPosAt(qreal, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setPosAt", args)
  }

}


func NewQGraphicsItemAnimation(args ...interface{}) QGraphicsItemAnimation {
  return QGraphicsItemAnimation{}
}


func (this *QGraphicsItemAnimation) xTranslationAt(args ...interface{}) () {
  // xTranslationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation14xTranslationAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "xTranslationAt", args)
  }

}


func (this *QGraphicsItemAnimation) setRotationAt(args ...interface{}) () {
  // setRotationAt(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation13setRotationAtEdd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setRotationAt", args)
  }

}


func (this *QGraphicsItemAnimation) posList(args ...interface{}) () {
  // posList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation7posListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "posList", args)
  }

}


func (this *QGraphicsItemAnimation) verticalScaleAt(args ...interface{}) () {
  // verticalScaleAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15verticalScaleAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "verticalScaleAt", args)
  }

}


func (this *QGraphicsItemAnimation) posAt(args ...interface{}) () {
  // posAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation5posAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "posAt", args)
  }

}


func (this *QGraphicsItemAnimation) horizontalShearAt(args ...interface{}) () {
  // horizontalShearAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation17horizontalShearAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "horizontalShearAt", args)
  }

}


func (this *QGraphicsItemAnimation) yTranslationAt(args ...interface{}) () {
  // yTranslationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation14yTranslationAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "yTranslationAt", args)
  }

}


func (this *QGraphicsItemAnimation) matrixAt(args ...interface{}) () {
  // matrixAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation8matrixAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "matrixAt", args)
  }

}


func (this *QGraphicsItemAnimation) item(args ...interface{}) () {
  // item()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation4itemEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "item", args)
  }

}


func (this *QGraphicsItemAnimation) FreeQGraphicsItemAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "~QGraphicsItemAnimation", args)
  }

}


func (this *QGraphicsItemAnimation) setScaleAt(args ...interface{}) () {
  // setScaleAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation10setScaleAtEddd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setScaleAt", args)
  }

}


func (this *QGraphicsItemAnimation) setTranslationAt(args ...interface{}) () {
  // setTranslationAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation16setTranslationAtEddd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setTranslationAt", args)
  }

}


func (this *QGraphicsItemAnimation) setShearAt(args ...interface{}) () {
  // setShearAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation10setShearAtEddd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setShearAt", args)
  }

}


func (this *QGraphicsItemAnimation) rotationAt(args ...interface{}) () {
  // rotationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation10rotationAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "rotationAt", args)
  }

}


func (this *QGraphicsItemAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "metaObject", args)
  }

}


func (this *QGraphicsItemAnimation) scaleList(args ...interface{}) () {
  // scaleList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation9scaleListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "scaleList", args)
  }

}


func (this *QGraphicsItemAnimation) rotationList(args ...interface{}) () {
  // rotationList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation12rotationListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "rotationList", args)
  }

}


func (this *QGraphicsItemAnimation) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation5resetEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "reset", args)
  }

}


func (this *QGraphicsItemAnimation) setTimeLine(args ...interface{}) () {
  // setTimeLine(class QTimeLine *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeLine{}) // "QTimeLine *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setTimeLine", args)
  }

}


func (this *QGraphicsItemAnimation) shearList(args ...interface{}) () {
  // shearList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation9shearListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "shearList", args)
  }

}


func (this *QGraphicsItemAnimation) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation5clearEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "clear", args)
  }

}


func (this *QGraphicsItemAnimation) translationList(args ...interface{}) () {
  // translationList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15translationListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "translationList", args)
  }

}


func (this *QGraphicsItemAnimation) setItem(args ...interface{}) () {
  // setItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setItem", args)
  }

}


func (this *QGraphicsItemAnimation) setStep(args ...interface{}) () {
  // setStep(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation7setStepEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setStep", args)
  }

}


func (this *QGraphicsItemAnimation) timeLine(args ...interface{}) () {
  // timeLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation8timeLineEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "timeLine", args)
  }

}


func (this *QGraphicsItemAnimation) horizontalScaleAt(args ...interface{}) () {
  // horizontalScaleAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation17horizontalScaleAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "horizontalScaleAt", args)
  }

}


func (this *QGraphicsItemAnimation) verticalShearAt(args ...interface{}) () {
  // verticalShearAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15verticalShearAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "verticalShearAt", args)
  }

}

// <= body block end

