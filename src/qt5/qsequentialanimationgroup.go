package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qsequentialanimationgroup.h
// dst-file: /src/core/qsequentialanimationgroup.go
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
  // proto:  const QMetaObject * QSequentialAnimationGroup::metaObject();
extern void C_ZNK25QSequentialAnimationGroup10metaObjectEv(void* qthis); // 4
  // proto:  QPauseAnimation * QSequentialAnimationGroup::insertPause(int index, int msecs);
extern void C_ZN25QSequentialAnimationGroup11insertPauseEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QPauseAnimation * QSequentialAnimationGroup::addPause(int msecs);
extern void C_ZN25QSequentialAnimationGroup8addPauseEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSequentialAnimationGroup::~QSequentialAnimationGroup();
extern void C_ZN25QSequentialAnimationGroupD2Ev(void* qthis); // 4
  // proto:  void QSequentialAnimationGroup::QSequentialAnimationGroup(QObject * parent);
extern void* C_ZN25QSequentialAnimationGroupC2EP7QObject(void* arg0); // 3
  // proto:  int QSequentialAnimationGroup::duration();
extern void C_ZNK25QSequentialAnimationGroup8durationEv(void* qthis); // 4
  // proto:  QAbstractAnimation * QSequentialAnimationGroup::currentAnimation();
extern void C_ZNK25QSequentialAnimationGroup16currentAnimationEv(void* qthis); // 4
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

// class sizeof(QSequentialAnimationGroup)=1
type QSequentialAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentAnimationChanged QSequentialAnimationGroup_currentAnimationChanged_signal;
}

// metaObject()
func (this *QSequentialAnimationGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK25QSequentialAnimationGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "metaObject", args)
  }

}

// insertPause(int, int)
func (this *QSequentialAnimationGroup) insertPause(args ...interface{}) () {
  // insertPause(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroup11insertPauseEii
    // invoke: QPauseAnimation * insertPause(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN25QSequentialAnimationGroup11insertPauseEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "insertPause", args)
  }

}

// addPause(int)
func (this *QSequentialAnimationGroup) addPause(args ...interface{}) () {
  // addPause(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroup8addPauseEi
    // invoke: QPauseAnimation * addPause(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN25QSequentialAnimationGroup8addPauseEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "addPause", args)
  }

}

// ~QSequentialAnimationGroup()
func (this *QSequentialAnimationGroup) FreeQSequentialAnimationGroup(args ...interface{}) () {
  // ~QSequentialAnimationGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroupD0Ev
    // invoke: void ~QSequentialAnimationGroup()
    C.C_ZN25QSequentialAnimationGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "~QSequentialAnimationGroup", args)
  }

}

// QSequentialAnimationGroup(class QObject *)
func NewQSequentialAnimationGroup(args ...interface{}) *QSequentialAnimationGroup {
  // QSequentialAnimationGroup(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroupC1EP7QObject
    // invoke: void QSequentialAnimationGroup(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN25QSequentialAnimationGroupC2EP7QObject(arg0)
    return &QSequentialAnimationGroup{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "QSequentialAnimationGroup", args)
  }

  return nil // QSequentialAnimationGroup{qclsinst:qthis}
}

// duration()
func (this *QSequentialAnimationGroup) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup8durationEv
    // invoke: int duration()
    var ret = C.C_ZNK25QSequentialAnimationGroup8durationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "duration", args)
  }

}

// currentAnimation()
func (this *QSequentialAnimationGroup) currentAnimation(args ...interface{}) () {
  // currentAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup16currentAnimationEv
    // invoke: QAbstractAnimation * currentAnimation()
    C.C_ZNK25QSequentialAnimationGroup16currentAnimationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "currentAnimation", args)
  }

}

// <= body block end

