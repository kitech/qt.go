package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qtimer.h
// dst-file: /src/core/qtimer.go
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
  // proto:  void QTimer::~QTimer();
extern void _ZN6QTimerD0Ev(void* qthis);
  // proto:  void QTimer::stop();
extern void _ZN6QTimer4stopEv(void* qthis);
  // proto:  int QTimer::timerId();
extern void _ZNK6QTimer7timerIdEv(void* qthis);
  // proto:  void QTimer::setSingleShot(bool singleShot);
extern void demth_ZN6QTimer13setSingleShotEb(void* qthis, bool arg0);
  // proto: static void QTimer::singleShot(int msec, const QObject * receiver, const char * member);
extern void _ZN6QTimer10singleShotEiPK7QObjectPKc(int arg0, void* arg1, char* arg2);
  // proto:  void QTimer::start();
extern void _ZN6QTimer5startEv(void* qthis);
  // proto:  int QTimer::interval();
extern void _ZNK6QTimer8intervalEv(void* qthis);
  // proto:  void QTimer::setInterval(int msec);
extern void _ZN6QTimer11setIntervalEi(void* qthis, int arg0);
  // proto:  void QTimer::start(int msec);
extern void _ZN6QTimer5startEi(void* qthis, int arg0);
  // proto:  void QTimer::QTimer(const QTimer & );
extern void* dector_ZN6QTimerC1ERKS_(void* arg0);
extern void demth_ZN6QTimerC1ERKS_(void* qthis, void* arg0);
  // proto:  int QTimer::remainingTime();
extern void _ZNK6QTimer13remainingTimeEv(void* qthis);
  // proto:  bool QTimer::isSingleShot();
extern void demth_ZNK6QTimer12isSingleShotEv(void* qthis);
  // proto:  bool QTimer::isActive();
extern void demth_ZNK6QTimer8isActiveEv(void* qthis);
  // proto:  const QMetaObject * QTimer::metaObject();
extern void _ZNK6QTimer10metaObjectEv(void* qthis);
  // proto:  void QTimer::QTimer(QObject * parent);
extern void* dector_ZN6QTimerC1EP7QObject(void* arg0);
extern void _ZN6QTimerC1EP7QObject(void* qthis, void* arg0);
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

// class sizeof(QTimer)=1
type QTimer struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _timeout QTimer_timeout_signal;
}

  // proto:  void QTimer::~QTimer();
func (this *QTimer) FreeQTimer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTimer", "~QTimer", args)
  }

}

  // proto:  void QTimer::stop();
func (this *QTimer) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimer4stopEv
    // invoke: void stop()
    C._ZN6QTimer4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "stop", args)
  }

}

  // proto:  int QTimer::timerId();
func (this *QTimer) timerId(args ...interface{}) () {
  // timerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer7timerIdEv
    // invoke: int timerId()
    C._ZNK6QTimer7timerIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "timerId", args)
  }

}

  // proto:  void QTimer::setSingleShot(bool singleShot);
func (this *QTimer) setSingleShot(args ...interface{}) () {
  // setSingleShot(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimer13setSingleShotEb
    // invoke: void setSingleShot(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C.demth_ZN6QTimer13setSingleShotEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimer", "setSingleShot", args)
  }

}

  // proto: static void QTimer::singleShot(int msec, const QObject * receiver, const char * member);
func (this *QTimer) singleShot_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTimer", "singleShot", args)
  }

}

  // proto:  void QTimer::start();
func (this *QTimer) start(args ...interface{}) () {
  // start()
  // start(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimer5startEv
    // invoke: void start()
    C._ZN6QTimer5startEv(this.qclsinst)
  case 1:
    // invoke: _ZN6QTimer5startEi
    // invoke: void start(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QTimer5startEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimer", "start", args)
  }

}

  // proto:  int QTimer::interval();
func (this *QTimer) interval(args ...interface{}) () {
  // interval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer8intervalEv
    // invoke: int interval()
    C._ZNK6QTimer8intervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "interval", args)
  }

}

  // proto:  void QTimer::setInterval(int msec);
func (this *QTimer) setInterval(args ...interface{}) () {
  // setInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimer11setIntervalEi
    // invoke: void setInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QTimer11setIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimer", "setInterval", args)
  }

}

  // proto:  void QTimer::QTimer(const QTimer & );
func NewQTimer(args ...interface{}) QTimer {
  return QTimer{}
}

  // proto:  int QTimer::remainingTime();
func (this *QTimer) remainingTime(args ...interface{}) () {
  // remainingTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer13remainingTimeEv
    // invoke: int remainingTime()
    C._ZNK6QTimer13remainingTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "remainingTime", args)
  }

}

  // proto:  bool QTimer::isSingleShot();
func (this *QTimer) isSingleShot(args ...interface{}) () {
  // isSingleShot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer12isSingleShotEv
    // invoke: bool isSingleShot()
    C.demth_ZNK6QTimer12isSingleShotEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "isSingleShot", args)
  }

}

  // proto:  bool QTimer::isActive();
func (this *QTimer) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer8isActiveEv
    // invoke: bool isActive()
    C.demth_ZNK6QTimer8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "isActive", args)
  }

}

  // proto:  const QMetaObject * QTimer::metaObject();
func (this *QTimer) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK6QTimer10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "metaObject", args)
  }

}

// <= body block end

