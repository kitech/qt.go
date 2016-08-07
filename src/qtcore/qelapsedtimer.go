package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern void C_ZN13QElapsedTimer10invalidateEv(void* qthis); // 4
  // proto:  bool QElapsedTimer::isValid();
extern bool C_ZNK13QElapsedTimer7isValidEv(void* qthis); // 4
  // proto:  qint64 QElapsedTimer::secsTo(const QElapsedTimer & other);
extern int64_t C_ZNK13QElapsedTimer6secsToERKS_(void* qthis, void* arg0); // 4
  // proto: static bool QElapsedTimer::isMonotonic();
extern bool C_ZN13QElapsedTimer11isMonotonicEv(); // 4
  // proto:  bool QElapsedTimer::hasExpired(qint64 timeout);
extern bool C_ZNK13QElapsedTimer10hasExpiredEx(void* qthis, int64_t arg0); // 4
  // proto:  void QElapsedTimer::QElapsedTimer();
extern void* C_ZN13QElapsedTimerC2Ev(); // 1
  // proto:  qint64 QElapsedTimer::elapsed();
extern int64_t C_ZNK13QElapsedTimer7elapsedEv(void* qthis); // 4
  // proto:  void QElapsedTimer::start();
extern void C_ZN13QElapsedTimer5startEv(void* qthis); // 4
  // proto:  qint64 QElapsedTimer::msecsTo(const QElapsedTimer & other);
extern int64_t C_ZNK13QElapsedTimer7msecsToERKS_(void* qthis, void* arg0); // 4
  // proto:  qint64 QElapsedTimer::msecsSinceReference();
extern int64_t C_ZNK13QElapsedTimer19msecsSinceReferenceEv(void* qthis); // 4
  // proto: static QElapsedTimer::ClockType QElapsedTimer::clockType();
extern void C_ZN13QElapsedTimer9clockTypeEv(); // 4
  // proto:  qint64 QElapsedTimer::nsecsElapsed();
extern int64_t C_ZNK13QElapsedTimer12nsecsElapsedEv(void* qthis); // 4
  // proto:  qint64 QElapsedTimer::restart();
extern int64_t C_ZN13QElapsedTimer7restartEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
func (this *QElapsedTimer) Invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer10invalidateEv
    // invoke: void invalidate()
    C.C_ZN13QElapsedTimer10invalidateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "invalidate", args)
  }

  return
}

// isValid()
func (this *QElapsedTimer) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK13QElapsedTimer7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isValid", args)
  }

  return
}

// secsTo(const class QElapsedTimer &)
func (this *QElapsedTimer) Secsto(args ...interface{}) (ret interface{}) {
  // secsTo(const class QElapsedTimer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QElapsedTimer{}) // "const QElapsedTimer &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer6secsToERKS_
    // invoke: qint64 secsTo(const class QElapsedTimer &)
    var arg0 = args[0].(*QElapsedTimer).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QElapsedTimer6secsToERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "secsTo", args)
  }

  return
}

// isMonotonic()
func (this *QElapsedTimer) Ismonotonic_S(args ...interface{}) (ret interface{}) {
  // isMonotonic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer11isMonotonicEv
    // invoke: bool isMonotonic()
    var ret0 = C.C_ZN13QElapsedTimer11isMonotonicEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isMonotonic", args)
  }

  return
}

// hasExpired(qint64)
func (this *QElapsedTimer) Hasexpired(args ...interface{}) (ret interface{}) {
  // hasExpired(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer10hasExpiredEx
    // invoke: bool hasExpired(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QElapsedTimer10hasExpiredEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "hasExpired", args)
  }

  return
}

// QElapsedTimer()
func NewQElapsedTimer(args ...interface{}) *QElapsedTimer {
  // QElapsedTimer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimerC1Ev
    // invoke: void QElapsedTimer()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QElapsedTimerC2Ev()
    return &QElapsedTimer{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QElapsedTimer", "QElapsedTimer", args)
  }

  return nil // QElapsedTimer{Qclsinst:qthis}
}

// elapsed()
func (this *QElapsedTimer) Elapsed(args ...interface{}) (ret interface{}) {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7elapsedEv
    // invoke: qint64 elapsed()
    var ret0 = C.C_ZNK13QElapsedTimer7elapsedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "elapsed", args)
  }

  return
}

// start()
func (this *QElapsedTimer) Start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer5startEv
    // invoke: void start()
    C.C_ZN13QElapsedTimer5startEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QElapsedTimer", "start", args)
  }

  return
}

// msecsTo(const class QElapsedTimer &)
func (this *QElapsedTimer) Msecsto(args ...interface{}) (ret interface{}) {
  // msecsTo(const class QElapsedTimer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QElapsedTimer{}) // "const QElapsedTimer &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7msecsToERKS_
    // invoke: qint64 msecsTo(const class QElapsedTimer &)
    var arg0 = args[0].(*QElapsedTimer).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QElapsedTimer7msecsToERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "msecsTo", args)
  }

  return
}

// msecsSinceReference()
func (this *QElapsedTimer) Msecssincereference(args ...interface{}) (ret interface{}) {
  // msecsSinceReference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer19msecsSinceReferenceEv
    // invoke: qint64 msecsSinceReference()
    var ret0 = C.C_ZNK13QElapsedTimer19msecsSinceReferenceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "msecsSinceReference", args)
  }

  return
}

// clockType()
func (this *QElapsedTimer) Clocktype_S(args ...interface{}) () {
  // clockType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer9clockTypeEv
    // invoke: QElapsedTimer::ClockType clockType()
    C.C_ZN13QElapsedTimer9clockTypeEv()
  default:
    qtrt.ErrorResolve("QElapsedTimer", "clockType", args)
  }

  return
}

// nsecsElapsed()
func (this *QElapsedTimer) Nsecselapsed(args ...interface{}) (ret interface{}) {
  // nsecsElapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer12nsecsElapsedEv
    // invoke: qint64 nsecsElapsed()
    var ret0 = C.C_ZNK13QElapsedTimer12nsecsElapsedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "nsecsElapsed", args)
  }

  return
}

// restart()
func (this *QElapsedTimer) Restart(args ...interface{}) (ret interface{}) {
  // restart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer7restartEv
    // invoke: qint64 restart()
    var ret0 = C.C_ZN13QElapsedTimer7restartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QElapsedTimer", "restart", args)
  }

  return
}

// <= body block end

