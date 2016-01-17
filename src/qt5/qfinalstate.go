package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtCore/qfinalstate.h
// dst-file: /src/core/qfinalstate.go
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
  // proto:  void QFinalState::QFinalState(QState * parent);
extern void _ZN11QFinalStateC2EP6QState(void* qthis, void* arg0); // 3
  // proto:  void QFinalState::~QFinalState();
extern void _ZN11QFinalStateD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QFinalState::metaObject();
extern void _ZNK11QFinalState10metaObjectEv(void* qthis); // 4
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

// class sizeof(QFinalState)=1
type QFinalState struct {
  /*qbase*/ QAbstractState;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QFinalState(class QState *)
func NewQFinalState(args ...interface{}) QFinalState {
  // QFinalState(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFinalStateC1EP6QState
    // invoke: void QFinalState(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QFinalStateC2EP6QState(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFinalState", "QFinalState", args)
  }

  return QFinalState{}
}

// ~QFinalState()
func (this *QFinalState) FreeQFinalState(args ...interface{}) () {
  // ~QFinalState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFinalStateD0Ev
    // invoke: void ~QFinalState()
    C._ZN11QFinalStateD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFinalState", "~QFinalState", args)
  }

}

// metaObject()
func (this *QFinalState) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFinalState10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QFinalState10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFinalState", "metaObject", args)
  }

}

// <= body block end

