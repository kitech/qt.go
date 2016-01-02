package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qelapsedtimer.h
// dst-file: /src/core/qelapsedtimer.go
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
  // proto:  void QElapsedTimer::start();
extern void _ZN13QElapsedTimer5startEv(void* qthis);
  // proto:  qint64 QElapsedTimer::nsecsElapsed();
extern void _ZNK13QElapsedTimer12nsecsElapsedEv(void* qthis);
  // proto:  void QElapsedTimer::invalidate();
extern void _ZN13QElapsedTimer10invalidateEv(void* qthis);
  // proto: static bool QElapsedTimer::isMonotonic();
extern void _ZN13QElapsedTimer11isMonotonicEv();
  // proto:  void QElapsedTimer::QElapsedTimer();
extern void* dector_ZN13QElapsedTimerC1Ev();
extern void _ZN13QElapsedTimerC1Ev(void* qthis);
  // proto:  qint64 QElapsedTimer::msecsTo(const QElapsedTimer & other);
extern void _ZNK13QElapsedTimer7msecsToERKS_(void* qthis, void* arg0);
  // proto:  qint64 QElapsedTimer::msecsSinceReference();
extern void _ZNK13QElapsedTimer19msecsSinceReferenceEv(void* qthis);
  // proto:  bool QElapsedTimer::hasExpired(qint64 timeout);
extern void _ZNK13QElapsedTimer10hasExpiredEx(void* qthis, long long arg0);
  // proto:  qint64 QElapsedTimer::restart();
extern void _ZN13QElapsedTimer7restartEv(void* qthis);
  // proto:  bool QElapsedTimer::isValid();
extern void _ZNK13QElapsedTimer7isValidEv(void* qthis);
  // proto:  qint64 QElapsedTimer::secsTo(const QElapsedTimer & other);
extern void _ZNK13QElapsedTimer6secsToERKS_(void* qthis, void* arg0);
  // proto:  qint64 QElapsedTimer::elapsed();
extern void _ZNK13QElapsedTimer7elapsedEv(void* qthis);
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

// class sizeof(QElapsedTimer)=16
type QElapsedTimer struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QElapsedTimer::start();
func (this *QElapsedTimer) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer5startEv
    // invoke: void start()
    C._ZN13QElapsedTimer5startEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "start", args)
  }

}

  // proto:  qint64 QElapsedTimer::nsecsElapsed();
func (this *QElapsedTimer) nsecsElapsed(args ...interface{}) () {
  // nsecsElapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer12nsecsElapsedEv
    // invoke: qint64 nsecsElapsed()
    C._ZNK13QElapsedTimer12nsecsElapsedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "nsecsElapsed", args)
  }

}

  // proto:  void QElapsedTimer::invalidate();
func (this *QElapsedTimer) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer10invalidateEv
    // invoke: void invalidate()
    C._ZN13QElapsedTimer10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "invalidate", args)
  }

}

  // proto: static bool QElapsedTimer::isMonotonic();
func (this *QElapsedTimer) isMonotonic_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isMonotonic", args)
  }

}

  // proto:  void QElapsedTimer::QElapsedTimer();
func NewQElapsedTimer(args ...interface{}) QElapsedTimer {
  return QElapsedTimer{}
}

  // proto:  qint64 QElapsedTimer::msecsTo(const QElapsedTimer & other);
func (this *QElapsedTimer) msecsTo(args ...interface{}) () {
  // msecsTo(const class QElapsedTimer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QElapsedTimer{}) // "const QElapsedTimer &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7msecsToERKS_
    // invoke: qint64 msecsTo(const class QElapsedTimer &)
    var arg0 = args[0].(QElapsedTimer).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QElapsedTimer7msecsToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "msecsTo", args)
  }

}

  // proto:  qint64 QElapsedTimer::msecsSinceReference();
func (this *QElapsedTimer) msecsSinceReference(args ...interface{}) () {
  // msecsSinceReference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer19msecsSinceReferenceEv
    // invoke: qint64 msecsSinceReference()
    C._ZNK13QElapsedTimer19msecsSinceReferenceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "msecsSinceReference", args)
  }

}

  // proto:  bool QElapsedTimer::hasExpired(qint64 timeout);
func (this *QElapsedTimer) hasExpired(args ...interface{}) () {
  // hasExpired(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer10hasExpiredEx
    // invoke: bool hasExpired(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZNK13QElapsedTimer10hasExpiredEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "hasExpired", args)
  }

}

  // proto:  qint64 QElapsedTimer::restart();
func (this *QElapsedTimer) restart(args ...interface{}) () {
  // restart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer7restartEv
    // invoke: qint64 restart()
    C._ZN13QElapsedTimer7restartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "restart", args)
  }

}

  // proto:  bool QElapsedTimer::isValid();
func (this *QElapsedTimer) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7isValidEv
    // invoke: bool isValid()
    C._ZNK13QElapsedTimer7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isValid", args)
  }

}

  // proto:  qint64 QElapsedTimer::secsTo(const QElapsedTimer & other);
func (this *QElapsedTimer) secsTo(args ...interface{}) () {
  // secsTo(const class QElapsedTimer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QElapsedTimer{}) // "const QElapsedTimer &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer6secsToERKS_
    // invoke: qint64 secsTo(const class QElapsedTimer &)
    var arg0 = args[0].(QElapsedTimer).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QElapsedTimer6secsToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "secsTo", args)
  }

}

  // proto:  qint64 QElapsedTimer::elapsed();
func (this *QElapsedTimer) elapsed(args ...interface{}) () {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7elapsedEv
    // invoke: qint64 elapsed()
    C._ZNK13QElapsedTimer7elapsedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "elapsed", args)
  }

}

// <= body block end

