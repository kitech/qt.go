package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qparallelanimationgroup.h
// dst-file: /src/core/qparallelanimationgroup.go
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
  // proto:  void QParallelAnimationGroup::QParallelAnimationGroup(QObject * parent);
extern void* C_ZN23QParallelAnimationGroupC2EP7QObject(void* arg0); // 3
  // proto:  void QParallelAnimationGroup::~QParallelAnimationGroup();
extern void C_ZN23QParallelAnimationGroupD2Ev(void* qthis); // 4
  // proto:  int QParallelAnimationGroup::duration();
extern int32_t C_ZNK23QParallelAnimationGroup8durationEv(void* qthis); // 4
  // proto:  const QMetaObject * QParallelAnimationGroup::metaObject();
extern void C_ZNK23QParallelAnimationGroup10metaObjectEv(void* qthis); // 4
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

// class sizeof(QParallelAnimationGroup)=1
type QParallelAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QParallelAnimationGroup(class QObject *)
func NewQParallelAnimationGroup(args ...interface{}) *QParallelAnimationGroup {
  // QParallelAnimationGroup(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QParallelAnimationGroupC1EP7QObject
    // invoke: void QParallelAnimationGroup(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QParallelAnimationGroupC2EP7QObject(arg0)
    return &QParallelAnimationGroup{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "QParallelAnimationGroup", args)
  }

  return nil // QParallelAnimationGroup{qclsinst:qthis}
}

// ~QParallelAnimationGroup()
func (this *QParallelAnimationGroup) Freeqparallelanimationgroup(args ...interface{}) () {
  // ~QParallelAnimationGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QParallelAnimationGroupD0Ev
    // invoke: void ~QParallelAnimationGroup()
    C.C_ZN23QParallelAnimationGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "~QParallelAnimationGroup", args)
  }

  return
}

// duration()
func (this *QParallelAnimationGroup) Duration(args ...interface{}) (ret interface{}) {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QParallelAnimationGroup8durationEv
    // invoke: int duration()
    var ret0 = C.C_ZNK23QParallelAnimationGroup8durationEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "duration", args)
  }

  return
}

// metaObject()
func (this *QParallelAnimationGroup) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QParallelAnimationGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK23QParallelAnimationGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "metaObject", args)
  }

  return
}

// <= body block end

