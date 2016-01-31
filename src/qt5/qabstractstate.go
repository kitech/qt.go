package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QStateMachine * QAbstractState::machine();
extern void C_ZNK14QAbstractState7machineEv(void* qthis); // 4
  // proto:  bool QAbstractState::active();
extern void C_ZNK14QAbstractState6activeEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractState::metaObject();
extern void C_ZNK14QAbstractState10metaObjectEv(void* qthis); // 4
  // proto:  QState * QAbstractState::parentState();
extern void C_ZNK14QAbstractState11parentStateEv(void* qthis); // 4
  // proto:  void QAbstractState::~QAbstractState();
extern void C_ZN14QAbstractStateD2Ev(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _entered QAbstractState_entered_signal;
//  _exited QAbstractState_exited_signal;
//  _activeChanged QAbstractState_activeChanged_signal;
}

// machine()
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
    // invoke: QStateMachine * machine()
    var ret = C.C_ZNK14QAbstractState7machineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractState", "machine", args)
  }

}

// active()
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
    // invoke: bool active()
    var ret = C.C_ZNK14QAbstractState6activeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractState", "active", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QAbstractState10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractState", "metaObject", args)
  }

}

// parentState()
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
    // invoke: QState * parentState()
    var ret = C.C_ZNK14QAbstractState11parentStateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractState", "parentState", args)
  }

}

// ~QAbstractState()
func (this *QAbstractState) FreeQAbstractState(args ...interface{}) () {
  // ~QAbstractState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QAbstractStateD0Ev
    // invoke: void ~QAbstractState()
    C.C_ZN14QAbstractStateD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractState", "~QAbstractState", args)
  }

}

// <= body block end

