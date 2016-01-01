package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qkeyeventtransition.h
// dst-file: /src/widgets/qkeyeventtransition.go
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
// class sizeof(QKeyEventTransition)=1
type QKeyEventTransition struct {
  /*qbase*/ QEventTransition;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QKeyEventTransition) setKey(args ...interface{}) () {
  // setKey(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QKeyEventTransition6setKeyEi
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "setKey", args)
  }

}


func (this *QKeyEventTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition10metaObjectEv
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "metaObject", args)
  }

}


func (this *QKeyEventTransition) FreeQKeyEventTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "~QKeyEventTransition", args)
  }

}


func (this *QKeyEventTransition) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition3keyEv
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "key", args)
  }

}


func NewQKeyEventTransition(args ...interface{}) QKeyEventTransition {
  return QKeyEventTransition{}
}

// <= body block end

