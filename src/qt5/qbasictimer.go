package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZN11QBasicTimer4stopEv(void* qthis); // 4
  // proto:  void QBasicTimer::start(int msec, QObject * obj);
extern void _ZN11QBasicTimer5startEiP7QObject(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QBasicTimer::QBasicTimer();
extern void _ZN11QBasicTimerC2Ev(void* qthis); // 1
  // proto:  void QBasicTimer::~QBasicTimer();
extern void _ZN11QBasicTimerD2Ev(void* qthis); // 2
  // proto:  int QBasicTimer::timerId();
extern void _ZNK11QBasicTimer7timerIdEv(void* qthis); // 2
  // proto:  bool QBasicTimer::isActive();
extern void _ZNK11QBasicTimer8isActiveEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// stop()
func (this *QBasicTimer) stop(args ...interface{}) () {
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
    C._ZN11QBasicTimer4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicTimer", "stop", args)
  }

}

// start(int, class QObject *)
func (this *QBasicTimer) start(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN11QBasicTimer5startEiP7QObject(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QBasicTimer", "start", args)
  }

}

// QBasicTimer()
func NewQBasicTimer(args ...interface{}) QBasicTimer {
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
    C._ZN11QBasicTimerC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QBasicTimer", "QBasicTimer", args)
  }

  return QBasicTimer{}
}

// ~QBasicTimer()
func (this *QBasicTimer) FreeQBasicTimer(args ...interface{}) () {
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
    C._ZN11QBasicTimerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicTimer", "~QBasicTimer", args)
  }

}

// timerId()
func (this *QBasicTimer) timerId(args ...interface{}) () {
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
    C._ZNK11QBasicTimer7timerIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicTimer", "timerId", args)
  }

}

// isActive()
func (this *QBasicTimer) isActive(args ...interface{}) () {
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
    C._ZNK11QBasicTimer8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBasicTimer", "isActive", args)
  }

}

// <= body block end

