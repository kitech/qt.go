package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QBasicTimer::~QBasicTimer();
extern void demth_ZN11QBasicTimerD0Ev(void* qthis);
  // proto:  void QBasicTimer::stop();
extern void _ZN11QBasicTimer4stopEv(void* qthis);
  // proto:  int QBasicTimer::timerId();
extern void demth_ZNK11QBasicTimer7timerIdEv(void* qthis);
  // proto:  bool QBasicTimer::isActive();
extern void demth_ZNK11QBasicTimer8isActiveEv(void* qthis);
  // proto:  void QBasicTimer::QBasicTimer();
extern void* dector_ZN11QBasicTimerC1Ev();
extern void demth_ZN11QBasicTimerC1Ev(void* qthis);
  // proto:  void QBasicTimer::start(int msec, QObject * obj);
extern void _ZN11QBasicTimer5startEiP7QObject(void* qthis, int arg0, void* arg1);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QBasicTimer::~QBasicTimer();
func (this *QBasicTimer) FreeQBasicTimer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBasicTimer", "~QBasicTimer", args)
  }

}

  // proto:  void QBasicTimer::stop();
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
  default:
    qtrt.ErrorResolve("QBasicTimer", "stop", args)
  }

}

  // proto:  int QBasicTimer::timerId();
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
  default:
    qtrt.ErrorResolve("QBasicTimer", "timerId", args)
  }

}

  // proto:  bool QBasicTimer::isActive();
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
  default:
    qtrt.ErrorResolve("QBasicTimer", "isActive", args)
  }

}

  // proto:  void QBasicTimer::QBasicTimer();
func NewQBasicTimer(args ...interface{}) QBasicTimer {
  return QBasicTimer{}
}

  // proto:  void QBasicTimer::start(int msec, QObject * obj);
func (this *QBasicTimer) start(args ...interface{}) () {
  // start(int, Qt::TimerType, class QObject *)
  // start(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::TimerType"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QBasicTimer5startEiN2Qt9TimerTypeEP7QObject
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN11QBasicTimer5startEiP7QObject
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QBasicTimer", "start", args)
  }

}

// <= body block end

