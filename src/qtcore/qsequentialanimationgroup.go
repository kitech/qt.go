package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  const QMetaObject * QSequentialAnimationGroup::metaObject();
extern void C_ZNK25QSequentialAnimationGroup10metaObjectEv(void* qthis); // 4
  // proto:  QPauseAnimation * QSequentialAnimationGroup::insertPause(int index, int msecs);
extern void* C_ZN25QSequentialAnimationGroup11insertPauseEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QPauseAnimation * QSequentialAnimationGroup::addPause(int msecs);
extern void* C_ZN25QSequentialAnimationGroup8addPauseEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSequentialAnimationGroup::~QSequentialAnimationGroup();
extern void C_ZN25QSequentialAnimationGroupD2Ev(void* qthis); // 4
  // proto:  void QSequentialAnimationGroup::QSequentialAnimationGroup(QObject * parent);
extern void* C_ZN25QSequentialAnimationGroupC2EP7QObject(void* arg0); // 3
  // proto:  int QSequentialAnimationGroup::duration();
extern int32_t C_ZNK25QSequentialAnimationGroup8durationEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QSequentialAnimationGroup)=1
type QSequentialAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _currentAnimationChanged QSequentialAnimationGroup_currentAnimationChanged_signal;
}

// metaObject()
func (this *QSequentialAnimationGroup) MetaObject(args ...interface{}) () {
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
    C.C_ZNK25QSequentialAnimationGroup10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "metaObject", args)
  }

  return
}

// insertPause(int, int)
func (this *QSequentialAnimationGroup) InsertPause(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN25QSequentialAnimationGroup11insertPauseEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPauseAnimation{}) // "QPauseAnimation *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "insertPause", args)
  }

  return
}

// addPause(int)
func (this *QSequentialAnimationGroup) AddPause(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN25QSequentialAnimationGroup8addPauseEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPauseAnimation{}) // "QPauseAnimation *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "addPause", args)
  }

  return
}

// ~QSequentialAnimationGroup()
func (this *QSequentialAnimationGroup) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN25QSequentialAnimationGroupD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "~QSequentialAnimationGroup", args)
  }

  return
}

// QSequentialAnimationGroup(class QObject *)
func GcfreeQSequentialAnimationGroup(this *QSequentialAnimationGroup) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN25QSequentialAnimationGroupC2EP7QObject(arg0)
    this := &QSequentialAnimationGroup{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQSequentialAnimationGroup)
    return this
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "QSequentialAnimationGroup", args)
  }

  return nil // QSequentialAnimationGroup{Qclsinst:qthis}
}

// duration()
func (this *QSequentialAnimationGroup) Duration(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK25QSequentialAnimationGroup8durationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "duration", args)
  }

  return
}

// currentAnimation()
func (this *QSequentialAnimationGroup) CurrentAnimation(args ...interface{}) () {
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
    C.C_ZNK25QSequentialAnimationGroup16currentAnimationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "currentAnimation", args)
  }

  return
}

// <= body block end

