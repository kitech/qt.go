package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicslinearlayout.h
// dst-file: /src/widgets/qgraphicslinearlayout.go
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
// class sizeof(QGraphicsLinearLayout)=1
type QGraphicsLinearLayout struct {
  /*qbase*/ QGraphicsLayout;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsLinearLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout7spacingEv
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "spacing", args)
  }

}


func NewQGraphicsLinearLayout(args ...interface{}) QGraphicsLinearLayout {
  return QGraphicsLinearLayout{}
}


func (this *QGraphicsLinearLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "itemAt", args)
  }

}


func (this *QGraphicsLinearLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "invalidate", args)
  }

}


func (this *QGraphicsLinearLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout11setGeometryERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setGeometry", args)
  }

}


func (this *QGraphicsLinearLayout) addStretch(args ...interface{}) () {
  // addStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10addStretchEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "addStretch", args)
  }

}


func (this *QGraphicsLinearLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout5countEv
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "count", args)
  }

}


func (this *QGraphicsLinearLayout) setSpacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10setSpacingEd
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setSpacing", args)
  }

}


func (this *QGraphicsLinearLayout) insertItem(args ...interface{}) () {
  // insertItem(int, class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "insertItem", args)
  }

}


func (this *QGraphicsLinearLayout) FreeQGraphicsLinearLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "~QGraphicsLinearLayout", args)
  }

}


func (this *QGraphicsLinearLayout) dump(args ...interface{}) () {
  // dump(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout4dumpEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "dump", args)
  }

}


func (this *QGraphicsLinearLayout) setStretchFactor(args ...interface{}) () {
  // setStretchFactor(class QGraphicsLayoutItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setStretchFactor", args)
  }

}


func (this *QGraphicsLinearLayout) addItem(args ...interface{}) () {
  // addItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "addItem", args)
  }

}


func (this *QGraphicsLinearLayout) itemSpacing(args ...interface{}) () {
  // itemSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout11itemSpacingEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "itemSpacing", args)
  }

}


func (this *QGraphicsLinearLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout8removeAtEi
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "removeAt", args)
  }

}


func (this *QGraphicsLinearLayout) insertStretch(args ...interface{}) () {
  // insertStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout13insertStretchEii
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "insertStretch", args)
  }

}


func (this *QGraphicsLinearLayout) setItemSpacing(args ...interface{}) () {
  // setItemSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout14setItemSpacingEid
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "setItemSpacing", args)
  }

}


func (this *QGraphicsLinearLayout) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "removeItem", args)
  }

}


func (this *QGraphicsLinearLayout) stretchFactor(args ...interface{}) () {
  // stretchFactor(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsLinearLayout", "stretchFactor", args)
  }

}

// <= body block end

