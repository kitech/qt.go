package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qbasictimer.h
// dst-file: /src/core/qbasictimer.go
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
  // proto:  void QBasicTimer::stop();
extern void C_ZN11QBasicTimer4stopEv(void* qthis); // 4
  // proto:  void QBasicTimer::start(int msec, QObject * obj);
extern void C_ZN11QBasicTimer5startEiP7QObject(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QBasicTimer::QBasicTimer();
extern void* C_ZN11QBasicTimerC2Ev(); // 1
  // proto:  void QBasicTimer::~QBasicTimer();
extern void C_ZN11QBasicTimerD2Ev(void* qthis); // 2
  // proto:  int QBasicTimer::timerId();
extern int32_t C_ZNK11QBasicTimer7timerIdEv(void* qthis); // 2
  // proto:  bool QBasicTimer::isActive();
extern bool C_ZNK11QBasicTimer8isActiveEv(void* qthis); // 2
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

// class sizeof(QBasicTimer)=4
type QBasicTimer struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// stop()
func (this *QBasicTimer) Stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimer4stopEv
    // invoke: void stop()
    C.C_ZN11QBasicTimer4stopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBasicTimer", "stop", args)
  }

  return
}

// start(int, class QObject *)
func (this *QBasicTimer) Start(args ...interface{}) () {
  // start(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimer5startEiP7QObject
    // invoke: void start(int, class QObject *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QBasicTimer5startEiP7QObject(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBasicTimer", "start", args)
  }

  return
}

// QBasicTimer()
func NewQBasicTimer(args ...interface{}) *QBasicTimer {
  // QBasicTimer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimerC1Ev
    // invoke: void QBasicTimer()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QBasicTimerC2Ev()
    return &QBasicTimer{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QBasicTimer", "QBasicTimer", args)
  }

  return nil // QBasicTimer{Qclsinst:qthis}
}

// ~QBasicTimer()
func (this *QBasicTimer) Freeqbasictimer(args ...interface{}) () {
  // ~QBasicTimer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimerD0Ev
    // invoke: void ~QBasicTimer()
    C.C_ZN11QBasicTimerD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBasicTimer", "~QBasicTimer", args)
  }

  return
}

// timerId()
func (this *QBasicTimer) Timerid(args ...interface{}) (ret interface{}) {
  // timerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QBasicTimer7timerIdEv
    // invoke: int timerId()
    var ret0 = C.C_ZNK11QBasicTimer7timerIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBasicTimer", "timerId", args)
  }

  return
}

// isActive()
func (this *QBasicTimer) Isactive(args ...interface{}) (ret interface{}) {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QBasicTimer8isActiveEv
    // invoke: bool isActive()
    var ret0 = C.C_ZNK11QBasicTimer8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBasicTimer", "isActive", args)
  }

  return
}

// <= body block end

