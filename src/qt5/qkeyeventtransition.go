package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qkeyeventtransition.h
// dst-file: /src/widgets/qkeyeventtransition.go
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
  // proto:  const QMetaObject * QKeyEventTransition::metaObject();
extern void C_ZNK19QKeyEventTransition10metaObjectEv(void* qthis); // 4
  // proto:  int QKeyEventTransition::key();
extern int32_t C_ZNK19QKeyEventTransition3keyEv(void* qthis); // 4
  // proto:  Qt::KeyboardModifiers QKeyEventTransition::modifierMask();
extern void C_ZNK19QKeyEventTransition12modifierMaskEv(void* qthis); // 4
  // proto:  void QKeyEventTransition::setKey(int key);
extern void C_ZN19QKeyEventTransition6setKeyEi(void* qthis, int32_t arg0); // 4
  // proto:  void QKeyEventTransition::~QKeyEventTransition();
extern void C_ZN19QKeyEventTransitionD2Ev(void* qthis); // 4
  // proto:  void QKeyEventTransition::QKeyEventTransition(QState * sourceState);
extern void* C_ZN19QKeyEventTransitionC2EP6QState(void* arg0); // 3
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

// class sizeof(QKeyEventTransition)=1
type QKeyEventTransition struct {
  /*qbase*/ QEventTransition;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QKeyEventTransition) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QKeyEventTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "metaObject", args)
  }

  return
}

// key()
func (this *QKeyEventTransition) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition3keyEv
    // invoke: int key()
    var ret0 = C.C_ZNK19QKeyEventTransition3keyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "key", args)
  }

  return
}

// modifierMask()
func (this *QKeyEventTransition) Modifiermask(args ...interface{}) () {
  // modifierMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition12modifierMaskEv
    // invoke: Qt::KeyboardModifiers modifierMask()
    C.C_ZNK19QKeyEventTransition12modifierMaskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "modifierMask", args)
  }

  return
}

// setKey(int)
func (this *QKeyEventTransition) Setkey(args ...interface{}) () {
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
    // invoke: void setKey(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN19QKeyEventTransition6setKeyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "setKey", args)
  }

  return
}

// ~QKeyEventTransition()
func (this *QKeyEventTransition) Freeqkeyeventtransition(args ...interface{}) () {
  // ~QKeyEventTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QKeyEventTransitionD0Ev
    // invoke: void ~QKeyEventTransition()
    C.C_ZN19QKeyEventTransitionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "~QKeyEventTransition", args)
  }

  return
}

// QKeyEventTransition(class QState *)
func NewQKeyEventTransition(args ...interface{}) *QKeyEventTransition {
  // QKeyEventTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QKeyEventTransitionC1EP6QState
    // invoke: void QKeyEventTransition(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QKeyEventTransitionC2EP6QState(arg0)
    return &QKeyEventTransition{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "QKeyEventTransition", args)
  }

  return nil // QKeyEventTransition{qclsinst:qthis}
}

// <= body block end

