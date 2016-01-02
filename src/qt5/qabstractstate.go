package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qabstractstate.h
// dst-file: /src/core/qabstractstate.go
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
  // proto:  void QAbstractState::~QAbstractState();
extern void _ZN14QAbstractStateD0Ev(void* qthis);
  // proto:  void QAbstractState::QAbstractState(const QAbstractState & );
extern void* dector_ZN14QAbstractStateC1ERKS_(void* arg0);
extern void _ZN14QAbstractStateC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractState::QAbstractState(QState * parent);
extern void* dector_ZN14QAbstractStateC1EP6QState(void* arg0);
extern void _ZN14QAbstractStateC1EP6QState(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractState::metaObject();
extern void _ZNK14QAbstractState10metaObjectEv(void* qthis);
  // proto:  QState * QAbstractState::parentState();
extern void _ZNK14QAbstractState11parentStateEv(void* qthis);
  // proto:  QStateMachine * QAbstractState::machine();
extern void _ZNK14QAbstractState7machineEv(void* qthis);
  // proto:  bool QAbstractState::active();
extern void _ZNK14QAbstractState6activeEv(void* qthis);
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

// class sizeof(QAbstractState)=1
type QAbstractState struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _entered QAbstractState_entered_signal;
//  _exited QAbstractState_exited_signal;
//  _activeChanged QAbstractState_activeChanged_signal;
}

  // proto:  void QAbstractState::~QAbstractState();
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

  // proto:  void QAbstractState::QAbstractState(const QAbstractState & );
func NewQAbstractState(args ...interface{}) QAbstractState {
  return QAbstractState{}
}

  // proto:  const QMetaObject * QAbstractState::metaObject();
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

  // proto:  QState * QAbstractState::parentState();
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

  // proto:  QStateMachine * QAbstractState::machine();
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

  // proto:  bool QAbstractState::active();
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

