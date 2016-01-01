package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qhistorystate.h
// dst-file: /src/core/qhistorystate.go
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
// class sizeof(QHistoryState)=1
type QHistoryState struct {
  /*qbase*/ QAbstractState;
  qclsinst uint64 /* *mut c_void*/;
//  _defaultStateChanged QHistoryState_defaultStateChanged_signal;
//  _historyTypeChanged QHistoryState_historyTypeChanged_signal;
}


func (this *QHistoryState) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QHistoryState10metaObjectEv
  default:
    qtrt.ErrorResolve("QHistoryState", "metaObject", args)
 }

}


func NewQHistoryState(args ...interface{})() {
}


func (this *QHistoryState) setDefaultState(args ...interface{}) () {
  // setDefaultState(class QAbstractState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractState{}) // "QAbstractState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QHistoryState15setDefaultStateEP14QAbstractState
  default:
    qtrt.ErrorResolve("QHistoryState", "setDefaultState", args)
 }

}


func (this *QHistoryState) defaultState(args ...interface{}) () {
  // defaultState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QHistoryState12defaultStateEv
  default:
    qtrt.ErrorResolve("QHistoryState", "defaultState", args)
 }

}


func (this *QHistoryState) FreeQHistoryState(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QHistoryState", "~QHistoryState", args)
 }

}

// <= body block end

