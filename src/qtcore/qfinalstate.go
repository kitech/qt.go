package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern void* C_ZN11QFinalStateC2EP6QState(void* arg0); // 3
  // proto:  void QFinalState::~QFinalState();
extern void C_ZN11QFinalStateD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QFinalState::metaObject();
extern void C_ZNK11QFinalState10metaObjectEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QFinalState)=1
type QFinalState struct {
  /*qbase*/ QAbstractState;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QFinalState(class QState *)
func GcfreeQFinalState(this *QFinalState) {
  qtrt.UniverseFree(this)
}
func NewQFinalState(args ...interface{}) *QFinalState {
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
    var arg0 = args[0].(*QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QFinalStateC2EP6QState(arg0)
    this := &QFinalState{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFinalState)
    return this
  default:
    qtrt.ErrorResolve("QFinalState", "QFinalState", args)
  }

  return nil // QFinalState{Qclsinst:qthis}
}

// ~QFinalState()
func (this *QFinalState) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QFinalStateD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QFinalState", "~QFinalState", args)
  }

  return
}

// metaObject()
func (this *QFinalState) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QFinalState10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFinalState", "metaObject", args)
  }

  return
}

// <= body block end

