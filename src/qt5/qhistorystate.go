package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QHistoryState::metaObject();
extern void _ZNK13QHistoryState10metaObjectEv(void* qthis);
  // proto:  void QHistoryState::QHistoryState(const QHistoryState & );
extern void* dector_ZN13QHistoryStateC1ERKS_(void* arg0);
extern void _ZN13QHistoryStateC1ERKS_(void* qthis, void* arg0);
  // proto:  void QHistoryState::setDefaultState(QAbstractState * state);
extern void _ZN13QHistoryState15setDefaultStateEP14QAbstractState(void* qthis, void* arg0);
  // proto:  void QHistoryState::QHistoryState(QState * parent);
extern void* dector_ZN13QHistoryStateC1EP6QState(void* arg0);
extern void _ZN13QHistoryStateC1EP6QState(void* qthis, void* arg0);
  // proto:  QAbstractState * QHistoryState::defaultState();
extern void _ZNK13QHistoryState12defaultStateEv(void* qthis);
  // proto:  void QHistoryState::~QHistoryState();
extern void _ZN13QHistoryStateD0Ev(void* qthis);
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _defaultStateChanged QHistoryState_defaultStateChanged_signal;
//  _historyTypeChanged QHistoryState_historyTypeChanged_signal;
}

  // proto:  const QMetaObject * QHistoryState::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QHistoryState10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHistoryState", "metaObject", args)
  }

}

  // proto:  void QHistoryState::QHistoryState(const QHistoryState & );
func NewQHistoryState(args ...interface{}) QHistoryState {
  return QHistoryState{}
}

  // proto:  void QHistoryState::setDefaultState(QAbstractState * state);
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
    // invoke: void setDefaultState(class QAbstractState *)
    var arg0 = args[0].(QAbstractState).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QHistoryState15setDefaultStateEP14QAbstractState(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QHistoryState", "setDefaultState", args)
  }

}

  // proto:  QAbstractState * QHistoryState::defaultState();
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
    // invoke: QAbstractState * defaultState()
    C._ZNK13QHistoryState12defaultStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHistoryState", "defaultState", args)
  }

}

  // proto:  void QHistoryState::~QHistoryState();
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

