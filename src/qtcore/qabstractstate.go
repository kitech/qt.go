package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern void* C_ZNK14QAbstractState7machineEv(void* qthis); // 4
  // proto:  bool QAbstractState::active();
extern bool C_ZNK14QAbstractState6activeEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractState::metaObject();
extern void C_ZNK14QAbstractState10metaObjectEv(void* qthis); // 4
  // proto:  QState * QAbstractState::parentState();
extern void* C_ZNK14QAbstractState11parentStateEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _entered QAbstractState_entered_signal;
//  _exited QAbstractState_exited_signal;
//  _activeChanged QAbstractState_activeChanged_signal;
}

// machine()
func (this *QAbstractState) Machine(args ...interface{}) (ret interface{}) {
  // machine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState7machineEv
    // invoke: QStateMachine * machine()
    var ret0 = C.C_ZNK14QAbstractState7machineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStateMachine{}) // "QStateMachine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractState", "machine", args)
  }

  return
}

// active()
func (this *QAbstractState) Active(args ...interface{}) (ret interface{}) {
  // active()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState6activeEv
    // invoke: bool active()
    var ret0 = C.C_ZNK14QAbstractState6activeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractState", "active", args)
  }

  return
}

// metaObject()
func (this *QAbstractState) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QAbstractState10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractState", "metaObject", args)
  }

  return
}

// parentState()
func (this *QAbstractState) Parentstate(args ...interface{}) (ret interface{}) {
  // parentState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QAbstractState11parentStateEv
    // invoke: QState * parentState()
    var ret0 = C.C_ZNK14QAbstractState11parentStateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QState{}) // "QState *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractState", "parentState", args)
  }

  return
}

// ~QAbstractState()
func (this *QAbstractState) Freeqabstractstate(args ...interface{}) () {
  // ~QAbstractState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QAbstractStateD0Ev
    // invoke: void ~QAbstractState()
    C.C_ZN14QAbstractStateD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractState", "~QAbstractState", args)
  }

  return
}

// <= body block end

