package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qhistorystate.h
// dst-file: /src/core/qhistorystate.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QHistoryState::HistoryType QHistoryState::historyType();
extern void C_ZNK13QHistoryState11historyTypeEv(void* qthis); // 4
  // proto:  void QHistoryState::~QHistoryState();
extern void C_ZN13QHistoryStateD2Ev(void* qthis); // 4
  // proto:  QAbstractState * QHistoryState::defaultState();
extern void C_ZNK13QHistoryState12defaultStateEv(void* qthis); // 4
  // proto:  void QHistoryState::QHistoryState(QState * parent);
extern void* C_ZN13QHistoryStateC2EP6QState(void* arg0); // 3
  // proto:  const QMetaObject * QHistoryState::metaObject();
extern void C_ZNK13QHistoryState10metaObjectEv(void* qthis); // 4
  // proto:  void QHistoryState::setDefaultState(QAbstractState * state);
extern void C_ZN13QHistoryState15setDefaultStateEP14QAbstractState(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QHistoryState)=1
type QHistoryState struct {
  /*qbase*/ QAbstractState;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _defaultStateChanged QHistoryState_defaultStateChanged_signal;
//  _historyTypeChanged QHistoryState_historyTypeChanged_signal;
}

// historyType()
func (this *QHistoryState) Historytype(args ...interface{}) () {
  // historyType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QHistoryState11historyTypeEv
    // invoke: QHistoryState::HistoryType historyType()
    C.C_ZNK13QHistoryState11historyTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHistoryState", "historyType", args)
  }

  return
}

// ~QHistoryState()
func (this *QHistoryState) Freeqhistorystate(args ...interface{}) () {
  // ~QHistoryState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QHistoryStateD0Ev
    // invoke: void ~QHistoryState()
    C.C_ZN13QHistoryStateD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHistoryState", "~QHistoryState", args)
  }

  return
}

// defaultState()
func (this *QHistoryState) Defaultstate(args ...interface{}) () {
  // defaultState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QHistoryState12defaultStateEv
    // invoke: QAbstractState * defaultState()
    C.C_ZNK13QHistoryState12defaultStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHistoryState", "defaultState", args)
  }

  return
}

// QHistoryState(class QState *)
func NewQHistoryState(args ...interface{}) *QHistoryState {
  // QHistoryState(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QHistoryStateC1EP6QState
    // invoke: void QHistoryState(class QState *)
    var arg0 = args[0].(*QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QHistoryStateC2EP6QState(arg0)
    return &QHistoryState{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QHistoryState", "QHistoryState", args)
  }

  return nil // QHistoryState{Qclsinst:qthis}
}

// metaObject()
func (this *QHistoryState) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QHistoryState10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QHistoryState10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QHistoryState", "metaObject", args)
  }

  return
}

// setDefaultState(class QAbstractState *)
func (this *QHistoryState) Setdefaultstate(args ...interface{}) () {
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
    // invoke: void setDefaultState(class QAbstractState *)
    var arg0 = args[0].(*QAbstractState).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QHistoryState15setDefaultStateEP14QAbstractState(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHistoryState", "setDefaultState", args)
  }

  return
}

// <= body block end

