package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qmouseeventtransition.h
// dst-file: /src/widgets/qmouseeventtransition.go
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
  // proto:  const QMetaObject * QMouseEventTransition::metaObject();
extern void C_ZNK21QMouseEventTransition10metaObjectEv(void* qthis); // 4
  // proto:  void QMouseEventTransition::setHitTestPath(const QPainterPath & path);
extern void C_ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  Qt::KeyboardModifiers QMouseEventTransition::modifierMask();
extern void C_ZNK21QMouseEventTransition12modifierMaskEv(void* qthis); // 4
  // proto:  QPainterPath QMouseEventTransition::hitTestPath();
extern void C_ZNK21QMouseEventTransition11hitTestPathEv(void* qthis); // 4
  // proto:  void QMouseEventTransition::~QMouseEventTransition();
extern void C_ZN21QMouseEventTransitionD2Ev(void* qthis); // 4
  // proto:  Qt::MouseButton QMouseEventTransition::button();
extern void C_ZNK21QMouseEventTransition6buttonEv(void* qthis); // 4
  // proto:  void QMouseEventTransition::QMouseEventTransition(QState * sourceState);
extern void C_ZN21QMouseEventTransitionC2EP6QState(void* qthis, void* arg0); // 3
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

// class sizeof(QMouseEventTransition)=1
type QMouseEventTransition struct {
  /*qbase*/ QEventTransition;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QMouseEventTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QMouseEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK21QMouseEventTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "metaObject", args)
  }

}

// setHitTestPath(const class QPainterPath &)
func (this *QMouseEventTransition) setHitTestPath(args ...interface{}) () {
  // setHitTestPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath
    // invoke: void setHitTestPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "setHitTestPath", args)
  }

}

// modifierMask()
func (this *QMouseEventTransition) modifierMask(args ...interface{}) () {
  // modifierMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QMouseEventTransition12modifierMaskEv
    // invoke: Qt::KeyboardModifiers modifierMask()
    C.C_ZNK21QMouseEventTransition12modifierMaskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "modifierMask", args)
  }

}

// hitTestPath()
func (this *QMouseEventTransition) hitTestPath(args ...interface{}) () {
  // hitTestPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QMouseEventTransition11hitTestPathEv
    // invoke: QPainterPath hitTestPath()
    C.C_ZNK21QMouseEventTransition11hitTestPathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "hitTestPath", args)
  }

}

// ~QMouseEventTransition()
func (this *QMouseEventTransition) FreeQMouseEventTransition(args ...interface{}) () {
  // ~QMouseEventTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QMouseEventTransitionD0Ev
    // invoke: void ~QMouseEventTransition()
    C.C_ZN21QMouseEventTransitionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "~QMouseEventTransition", args)
  }

}

// button()
func (this *QMouseEventTransition) button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QMouseEventTransition6buttonEv
    // invoke: Qt::MouseButton button()
    C.C_ZNK21QMouseEventTransition6buttonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "button", args)
  }

}

// QMouseEventTransition(class QState *)
func NewQMouseEventTransition(args ...interface{}) QMouseEventTransition {
  // QMouseEventTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QMouseEventTransitionC1EP6QState
    // invoke: void QMouseEventTransition(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN21QMouseEventTransitionC2EP6QState(qthis, arg0)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "QMouseEventTransition", args)
  }

  return QMouseEventTransition{}
}

// <= body block end

