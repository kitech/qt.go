package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qsplitter.h
// dst-file: /src/widgets/qsplitter.go
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
// class sizeof(QSplitter)=1
type QSplitter struct {
  /*qbase*/ QFrame;
  qclsinst uint64 /* *mut c_void*/;
//  _splitterMoved QSplitter_splitterMoved_signal;
}

// class sizeof(QSplitterHandle)=1
type QSplitterHandle struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSplitter) insertWidget(args ...interface{}) () {
  // insertWidget(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter12insertWidgetEiP7QWidget
  default:
    qtrt.ErrorResolve("QSplitter", "insertWidget", args)
  }

}


func (this *QSplitter) childrenCollapsible(args ...interface{}) () {
  // childrenCollapsible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter19childrenCollapsibleEv
  default:
    qtrt.ErrorResolve("QSplitter", "childrenCollapsible", args)
  }

}


func (this *QSplitter) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter5countEv
  default:
    qtrt.ErrorResolve("QSplitter", "count", args)
  }

}


func (this *QSplitter) saveState(args ...interface{}) () {
  // saveState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter9saveStateEv
  default:
    qtrt.ErrorResolve("QSplitter", "saveState", args)
  }

}


func (this *QSplitter) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter10metaObjectEv
  default:
    qtrt.ErrorResolve("QSplitter", "metaObject", args)
  }

}


func (this *QSplitter) opaqueResize(args ...interface{}) () {
  // opaqueResize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter12opaqueResizeEv
  default:
    qtrt.ErrorResolve("QSplitter", "opaqueResize", args)
  }

}


func (this *QSplitter) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter9addWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QSplitter", "addWidget", args)
  }

}


func NewQSplitter(args ...interface{}) QSplitter {
  return QSplitter{}
}


func (this *QSplitter) setHandleWidth(args ...interface{}) () {
  // setHandleWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter14setHandleWidthEi
  default:
    qtrt.ErrorResolve("QSplitter", "setHandleWidth", args)
  }

}


func (this *QSplitter) setStretchFactor(args ...interface{}) () {
  // setStretchFactor(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter16setStretchFactorEii
  default:
    qtrt.ErrorResolve("QSplitter", "setStretchFactor", args)
  }

}


func (this *QSplitter) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QSplitter", "minimumSizeHint", args)
  }

}


func (this *QSplitter) setOpaqueResize(args ...interface{}) () {
  // setOpaqueResize(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter15setOpaqueResizeEb
  default:
    qtrt.ErrorResolve("QSplitter", "setOpaqueResize", args)
  }

}


func (this *QSplitter) widget(args ...interface{}) () {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter6widgetEi
  default:
    qtrt.ErrorResolve("QSplitter", "widget", args)
  }

}


func (this *QSplitter) sizes(args ...interface{}) () {
  // sizes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter5sizesEv
  default:
    qtrt.ErrorResolve("QSplitter", "sizes", args)
  }

}


func (this *QSplitter) isCollapsible(args ...interface{}) () {
  // isCollapsible(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter13isCollapsibleEi
  default:
    qtrt.ErrorResolve("QSplitter", "isCollapsible", args)
  }

}


func (this *QSplitter) setCollapsible(args ...interface{}) () {
  // setCollapsible(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter14setCollapsibleEib
  default:
    qtrt.ErrorResolve("QSplitter", "setCollapsible", args)
  }

}


func (this *QSplitter) restoreState(args ...interface{}) () {
  // restoreState(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter12restoreStateERK10QByteArray
  default:
    qtrt.ErrorResolve("QSplitter", "restoreState", args)
  }

}


func (this *QSplitter) FreeQSplitter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSplitter", "~QSplitter", args)
  }

}


func (this *QSplitter) setChildrenCollapsible(args ...interface{}) () {
  // setChildrenCollapsible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter22setChildrenCollapsibleEb
  default:
    qtrt.ErrorResolve("QSplitter", "setChildrenCollapsible", args)
  }

}


func (this *QSplitter) handleWidth(args ...interface{}) () {
  // handleWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter11handleWidthEv
  default:
    qtrt.ErrorResolve("QSplitter", "handleWidth", args)
  }

}


func (this *QSplitter) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter7refreshEv
  default:
    qtrt.ErrorResolve("QSplitter", "refresh", args)
  }

}


func (this *QSplitter) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter8sizeHintEv
  default:
    qtrt.ErrorResolve("QSplitter", "sizeHint", args)
  }

}


func (this *QSplitter) indexOf(args ...interface{}) () {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter7indexOfEP7QWidget
  default:
    qtrt.ErrorResolve("QSplitter", "indexOf", args)
  }

}


func (this *QSplitter) getRange(args ...interface{}) () {
  // getRange(int, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter8getRangeEiPiS0_
  default:
    qtrt.ErrorResolve("QSplitter", "getRange", args)
  }

}


func (this *QSplitter) handle(args ...interface{}) () {
  // handle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter6handleEi
  default:
    qtrt.ErrorResolve("QSplitter", "handle", args)
  }

}


func (this *QSplitterHandle) FreeQSplitterHandle(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSplitterHandle", "~QSplitterHandle", args)
  }

}


func NewQSplitterHandle(args ...interface{}) QSplitterHandle {
  return QSplitterHandle{}
}


func (this *QSplitterHandle) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle8sizeHintEv
  default:
    qtrt.ErrorResolve("QSplitterHandle", "sizeHint", args)
  }

}


func (this *QSplitterHandle) opaqueResize(args ...interface{}) () {
  // opaqueResize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle12opaqueResizeEv
  default:
    qtrt.ErrorResolve("QSplitterHandle", "opaqueResize", args)
  }

}


func (this *QSplitterHandle) splitter(args ...interface{}) () {
  // splitter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle8splitterEv
  default:
    qtrt.ErrorResolve("QSplitterHandle", "splitter", args)
  }

}


func (this *QSplitterHandle) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle10metaObjectEv
  default:
    qtrt.ErrorResolve("QSplitterHandle", "metaObject", args)
  }

}

// <= body block end

