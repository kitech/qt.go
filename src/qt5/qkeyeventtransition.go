package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QKeyEventTransition::setKey(int key);
extern void _ZN19QKeyEventTransition6setKeyEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QKeyEventTransition::metaObject();
extern void _ZNK19QKeyEventTransition10metaObjectEv(void* qthis);
  // proto:  void QKeyEventTransition::~QKeyEventTransition();
extern void _ZN19QKeyEventTransitionD0Ev(void* qthis);
  // proto:  int QKeyEventTransition::key();
extern void _ZNK19QKeyEventTransition3keyEv(void* qthis);
  // proto:  void QKeyEventTransition::QKeyEventTransition(QState * sourceState);
extern void* dector_ZN19QKeyEventTransitionC1EP6QState(void* arg0);
extern void _ZN19QKeyEventTransitionC1EP6QState(void* qthis, void* arg0);
  // proto:  void QKeyEventTransition::QKeyEventTransition(const QKeyEventTransition & );
extern void* dector_ZN19QKeyEventTransitionC1ERKS_(void* arg0);
extern void _ZN19QKeyEventTransitionC1ERKS_(void* qthis, void* arg0);
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

  // proto:  void QKeyEventTransition::setKey(int key);
func (this *QKeyEventTransition) setKey(args ...interface{}) () {
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
    C._ZN19QKeyEventTransition6setKeyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "setKey", args)
  }

}

  // proto:  const QMetaObject * QKeyEventTransition::metaObject();
func (this *QKeyEventTransition) metaObject(args ...interface{}) () {
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
    C._ZNK19QKeyEventTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "metaObject", args)
  }

}

  // proto:  void QKeyEventTransition::~QKeyEventTransition();
func (this *QKeyEventTransition) FreeQKeyEventTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "~QKeyEventTransition", args)
  }

}

  // proto:  int QKeyEventTransition::key();
func (this *QKeyEventTransition) key(args ...interface{}) () {
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
    C._ZNK19QKeyEventTransition3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "key", args)
  }

}

  // proto:  void QKeyEventTransition::QKeyEventTransition(QState * sourceState);
func NewQKeyEventTransition(args ...interface{}) QKeyEventTransition {
  return QKeyEventTransition{}
}

// <= body block end

