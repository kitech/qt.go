package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QElapsedTimer::invalidate();
extern void _ZN13QElapsedTimer10invalidateEv(void* qthis); // 4
  // proto:  bool QElapsedTimer::isValid();
extern void _ZNK13QElapsedTimer7isValidEv(void* qthis); // 4
  // proto:  qint64 QElapsedTimer::secsTo(const QElapsedTimer & other);
extern void _ZNK13QElapsedTimer6secsToERKS_(void* qthis, void* arg0); // 4
  // proto: static bool QElapsedTimer::isMonotonic();
extern void _ZN13QElapsedTimer11isMonotonicEv(); // 4
  // proto:  bool QElapsedTimer::hasExpired(qint64 timeout);
extern void _ZNK13QElapsedTimer10hasExpiredEx(void* qthis, int64_t arg0); // 4
  // proto:  void QElapsedTimer::QElapsedTimer();
extern void _ZN13QElapsedTimerC2Ev(void* qthis); // 1
  // proto:  qint64 QElapsedTimer::elapsed();
extern void _ZNK13QElapsedTimer7elapsedEv(void* qthis); // 4
  // proto:  void QElapsedTimer::start();
extern void _ZN13QElapsedTimer5startEv(void* qthis); // 4
  // proto:  qint64 QElapsedTimer::msecsTo(const QElapsedTimer & other);
extern void _ZNK13QElapsedTimer7msecsToERKS_(void* qthis, void* arg0); // 4
  // proto:  qint64 QElapsedTimer::msecsSinceReference();
extern void _ZNK13QElapsedTimer19msecsSinceReferenceEv(void* qthis); // 4
  // proto: static QElapsedTimer::ClockType QElapsedTimer::clockType();
extern void _ZN13QElapsedTimer9clockTypeEv(); // 4
  // proto:  qint64 QElapsedTimer::nsecsElapsed();
extern void _ZNK13QElapsedTimer12nsecsElapsedEv(void* qthis); // 4
  // proto:  qint64 QElapsedTimer::restart();
extern void _ZN13QElapsedTimer7restartEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
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

// isValid()
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

// secsTo(const class QElapsedTimer &)
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

// isMonotonic()
func (this *QElapsedTimer) isMonotonic_s(args ...interface{}) () {
  // isMonotonic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer11isMonotonicEv
    // invoke: bool isMonotonic()
    C._ZN13QElapsedTimer11isMonotonicEv()
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isMonotonic", args)
  }

}

// hasExpired(qint64)
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

// QElapsedTimer()
func NewQElapsedTimer(args ...interface{}) QElapsedTimer {
  // QElapsedTimer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimerC1Ev
    // invoke: void QElapsedTimer()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QElapsedTimerC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "QElapsedTimer", args)
  }

  return QElapsedTimer{}
}

// elapsed()
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

// start()
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

// msecsTo(const class QElapsedTimer &)
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

// msecsSinceReference()
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

// clockType()
func (this *QElapsedTimer) clockType_s(args ...interface{}) () {
  // clockType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer9clockTypeEv
    // invoke: QElapsedTimer::ClockType clockType()
    C._ZN13QElapsedTimer9clockTypeEv()
  default:
    qtrt.ErrorResolve("QElapsedTimer", "clockType", args)
  }

}

// nsecsElapsed()
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

// restart()
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

// <= body block end

