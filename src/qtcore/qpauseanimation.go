package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qpauseanimation.h
// dst-file: /src/core/qpauseanimation.go
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
  // proto:  void QPauseAnimation::~QPauseAnimation();
extern void C_ZN15QPauseAnimationD2Ev(void* qthis); // 4
  // proto:  void QPauseAnimation::QPauseAnimation(int msecs, QObject * parent);
extern void* C_ZN15QPauseAnimationC2EiP7QObject(int32_t arg0, void* arg1); // 3
  // proto:  void QPauseAnimation::QPauseAnimation(QObject * parent);
extern void* C_ZN15QPauseAnimationC2EP7QObject(void* arg0); // 3
  // proto:  int QPauseAnimation::duration();
extern int32_t C_ZNK15QPauseAnimation8durationEv(void* qthis); // 4
  // proto:  const QMetaObject * QPauseAnimation::metaObject();
extern void C_ZNK15QPauseAnimation10metaObjectEv(void* qthis); // 4
  // proto:  void QPauseAnimation::setDuration(int msecs);
extern void C_ZN15QPauseAnimation11setDurationEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QPauseAnimation)=1
type QPauseAnimation struct {
  /*qbase*/ QAbstractAnimation;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QPauseAnimation()
func (this *QPauseAnimation) Freeqpauseanimation(args ...interface{}) () {
  // ~QPauseAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QPauseAnimationD0Ev
    // invoke: void ~QPauseAnimation()
    C.C_ZN15QPauseAnimationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPauseAnimation", "~QPauseAnimation", args)
  }

  return
}

// QPauseAnimation(int, class QObject *)
func NewQPauseAnimation(args ...interface{}) *QPauseAnimation {
  // QPauseAnimation(int, class QObject *)
  // QPauseAnimation(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QPauseAnimationC1EiP7QObject
    // invoke: void QPauseAnimation(int, class QObject *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QPauseAnimationC2EiP7QObject(arg0, arg1)
    return &QPauseAnimation{Qclsinst:qthis}
  case 1:
    // invoke: _ZN15QPauseAnimationC1EP7QObject
    // invoke: void QPauseAnimation(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QPauseAnimationC2EP7QObject(arg0)
    return &QPauseAnimation{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPauseAnimation", "QPauseAnimation", args)
  }

  return nil // QPauseAnimation{Qclsinst:qthis}
}

// duration()
func (this *QPauseAnimation) Duration(args ...interface{}) (ret interface{}) {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QPauseAnimation8durationEv
    // invoke: int duration()
    var ret0 = C.C_ZNK15QPauseAnimation8durationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPauseAnimation", "duration", args)
  }

  return
}

// metaObject()
func (this *QPauseAnimation) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QPauseAnimation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QPauseAnimation10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPauseAnimation", "metaObject", args)
  }

  return
}

// setDuration(int)
func (this *QPauseAnimation) Setduration(args ...interface{}) () {
  // setDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QPauseAnimation11setDurationEi
    // invoke: void setDuration(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QPauseAnimation11setDurationEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPauseAnimation", "setDuration", args)
  }

  return
}

// <= body block end

