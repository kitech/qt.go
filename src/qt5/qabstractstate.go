package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qabstractstate.h
// dst-file: /src/core/qabstractstate.go
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
// class sizeof(QAbstractState)=1
type QAbstractState struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _entered QAbstractState_entered_signal;
//  _exited QAbstractState_exited_signal;
//  _activeChanged QAbstractState_activeChanged_signal;
}


func (this *QAbstractState) FreeQAbstractState(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractState", "~QAbstractState", args)
  }

}


func NewQAbstractState(args ...interface{}) QAbstractState {
  return QAbstractState{}
}


func (this *QAbstractState) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractState", "metaObject", args)
  }

}


func (this *QAbstractState) parentState(args ...interface{}) () {
  // parentState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState11parentStateEv
  default:
    qtrt.ErrorResolve("QAbstractState", "parentState", args)
  }

}


func (this *QAbstractState) machine(args ...interface{}) () {
  // machine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState7machineEv
  default:
    qtrt.ErrorResolve("QAbstractState", "machine", args)
  }

}


func (this *QAbstractState) active(args ...interface{}) () {
  // active()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState6activeEv
  default:
    qtrt.ErrorResolve("QAbstractState", "active", args)
  }

}

// <= body block end

