package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicslayout.h
// dst-file: /src/widgets/qgraphicslayout.go
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
// class sizeof(QGraphicsLayout)=1
type QGraphicsLayout struct {
  /*qbase*/ QGraphicsLayoutItem;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsLayout) updateGeometry(args ...interface{}) () {
  // updateGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout14updateGeometryEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "updateGeometry", args)
  }

}


func (this *QGraphicsLayout) isActivated(args ...interface{}) () {
  // isActivated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout11isActivatedEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "isActivated", args)
  }

}


func (this *QGraphicsLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "invalidate", args)
  }

}


func (this *QGraphicsLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout8removeAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "removeAt", args)
  }

}


func (this *QGraphicsLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "itemAt", args)
  }

}


func NewQGraphicsLayout(args ...interface{}) QGraphicsLayout {
  return QGraphicsLayout{}
}


func (this *QGraphicsLayout) getContentsMargins(args ...interface{}) () {
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
    // invoke: _ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "getContentsMargins", args)
  }

}


func (this *QGraphicsLayout) setContentsMargins(args ...interface{}) () {
  // setContentsMargins(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout18setContentsMarginsEdddd
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setContentsMargins", args)
  }

}


func (this *QGraphicsLayout) widgetEvent(args ...interface{}) () {
  // widgetEvent(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout11widgetEventEP6QEvent
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "widgetEvent", args)
  }

}


func (this *QGraphicsLayout) instantInvalidatePropagation_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "instantInvalidatePropagation", args)
  }

}


func (this *QGraphicsLayout) setInstantInvalidatePropagation_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "setInstantInvalidatePropagation", args)
  }

}


func (this *QGraphicsLayout) FreeQGraphicsLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "~QGraphicsLayout", args)
  }

}


func (this *QGraphicsLayout) activate(args ...interface{}) () {
  // activate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsLayout8activateEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "activate", args)
  }

}


func (this *QGraphicsLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsLayout5countEv
  default:
    qtrt.ErrorResolve("QGraphicsLayout", "count", args)
  }

}

// <= body block end

