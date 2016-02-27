package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
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
extern void* C_ZNK21QMouseEventTransition11hitTestPathEv(void* qthis); // 4
  // proto:  void QMouseEventTransition::~QMouseEventTransition();
extern void C_ZN21QMouseEventTransitionD2Ev(void* qthis); // 4
  // proto:  Qt::MouseButton QMouseEventTransition::button();
extern void C_ZNK21QMouseEventTransition6buttonEv(void* qthis); // 4
  // proto:  void QMouseEventTransition::QMouseEventTransition(QState * sourceState);
extern void* C_ZN21QMouseEventTransitionC2EP6QState(void* arg0); // 3
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QMouseEventTransition)=1
type QMouseEventTransition struct {
  /*qbase*/ qtcore.QEventTransition;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QMouseEventTransition) MetaObject(args ...interface{}) () {
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
    C.C_ZNK21QMouseEventTransition10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "metaObject", args)
  }

  return
}

// setHitTestPath(const class QPainterPath &)
func (this *QMouseEventTransition) SetHitTestPath(args ...interface{}) () {
  // setHitTestPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath
    // invoke: void setHitTestPath(const class QPainterPath &)
    var arg0 = args[0].(*qtgui.QPainterPath).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "setHitTestPath", args)
  }

  return
}

// modifierMask()
func (this *QMouseEventTransition) ModifierMask(args ...interface{}) () {
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
    C.C_ZNK21QMouseEventTransition12modifierMaskEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "modifierMask", args)
  }

  return
}

// hitTestPath()
func (this *QMouseEventTransition) HitTestPath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QMouseEventTransition11hitTestPathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPainterPath{}) // "QPainterPath"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "hitTestPath", args)
  }

  return
}

// ~QMouseEventTransition()
func (this *QMouseEventTransition) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN21QMouseEventTransitionD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "~QMouseEventTransition", args)
  }

  return
}

// button()
func (this *QMouseEventTransition) Button(args ...interface{}) () {
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
    C.C_ZNK21QMouseEventTransition6buttonEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "button", args)
  }

  return
}

// QMouseEventTransition(class QState *)
func GcfreeQMouseEventTransition(this *QMouseEventTransition) {
  qtrt.UniverseFree(this)
}
func NewQMouseEventTransition(args ...interface{}) *QMouseEventTransition {
  // QMouseEventTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QMouseEventTransitionC1EP6QState
    // invoke: void QMouseEventTransition(class QState *)
    var arg0 = args[0].(*qtcore.QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QMouseEventTransitionC2EP6QState(arg0)
    this := &QMouseEventTransition{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQMouseEventTransition)
    return this
  default:
    qtrt.ErrorResolve("QMouseEventTransition", "QMouseEventTransition", args)
  }

  return nil // QMouseEventTransition{Qclsinst:qthis}
}

// <= body block end

