package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  Qt::TimerType QTimer::timerType();
extern void C_ZNK6QTimer9timerTypeEv(void* qthis); // 2
  // proto:  void QTimer::~QTimer();
extern void C_ZN6QTimerD2Ev(void* qthis); // 4
  // proto:  void QTimer::setSingleShot(bool singleShot);
extern void C_ZN6QTimer13setSingleShotEb(void* qthis, bool arg0); // 2
  // proto:  bool QTimer::isSingleShot();
extern bool C_ZNK6QTimer12isSingleShotEv(void* qthis); // 2
  // proto:  void QTimer::start(int msec);
extern void C_ZN6QTimer5startEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTimer::start();
extern void C_ZN6QTimer5startEv(void* qthis); // 4
  // proto:  void QTimer::QTimer(QObject * parent);
extern void* C_ZN6QTimerC2EP7QObject(void* arg0); // 3
  // proto:  int QTimer::timerId();
extern int32_t C_ZNK6QTimer7timerIdEv(void* qthis); // 2
  // proto:  void QTimer::setInterval(int msec);
extern void C_ZN6QTimer11setIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTimer::stop();
extern void C_ZN6QTimer4stopEv(void* qthis); // 4
  // proto:  int QTimer::remainingTime();
extern int32_t C_ZNK6QTimer13remainingTimeEv(void* qthis); // 4
  // proto:  bool QTimer::isActive();
extern bool C_ZNK6QTimer8isActiveEv(void* qthis); // 2
  // proto:  const QMetaObject * QTimer::metaObject();
extern void C_ZNK6QTimer10metaObjectEv(void* qthis); // 4
  // proto:  int QTimer::interval();
extern int32_t C_ZNK6QTimer8intervalEv(void* qthis); // 2
  // proto: static void QTimer::singleShot(int msec, const QObject * receiver, const char * member);
extern void C_ZN6QTimer10singleShotEiPK7QObjectPKc(int32_t arg0, void* arg1, void* arg2); // 4
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

// class sizeof(QTimer)=1
type QTimer struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _timeout QTimer_timeout_signal;
}

// timerType()
func (this *QTimer) TimerType(args ...interface{}) () {
  // timerType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QTimer9timerTypeEv
    // invoke: Qt::TimerType timerType()
    C.C_ZNK6QTimer9timerTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "timerType", args)
  }

  return
}

// ~QTimer()
func (this *QTimer) Free(args ...interface{}) () {
  // ~QTimer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimerD0Ev
    // invoke: void ~QTimer()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN6QTimerD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTimer", "~QTimer", args)
  }

  return
}

// setSingleShot(_Bool)
func (this *QTimer) SetSingleShot(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QTimer13setSingleShotEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimer", "setSingleShot", args)
  }

  return
}

// isSingleShot()
func (this *QTimer) IsSingleShot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QTimer12isSingleShotEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimer", "isSingleShot", args)
  }

  return
}

// start(int)
func (this *QTimer) Start(args ...interface{}) () {
  // start(int)
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimer5startEi
    // invoke: void start(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QTimer5startEi(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QTimer5startEv
    // invoke: void start()
    C.C_ZN6QTimer5startEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "start", args)
  }

  return
}

// QTimer(class QObject *)
func GcfreeQTimer(this *QTimer) {
  qtrt.UniverseFree(this)
}
func NewQTimer(args ...interface{}) *QTimer {
  // QTimer(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimerC1EP7QObject
    // invoke: void QTimer(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QTimerC2EP7QObject(arg0)
    this := &QTimer{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTimer)
    return this
  default:
    qtrt.ErrorResolve("QTimer", "QTimer", args)
  }

  return nil // QTimer{Qclsinst:qthis}
}

// timerId()
func (this *QTimer) TimerId(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QTimer7timerIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimer", "timerId", args)
  }

  return
}

// setInterval(int)
func (this *QTimer) SetInterval(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QTimer11setIntervalEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimer", "setInterval", args)
  }

  return
}

// stop()
func (this *QTimer) Stop(args ...interface{}) () {
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
    C.C_ZN6QTimer4stopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "stop", args)
  }

  return
}

// remainingTime()
func (this *QTimer) RemainingTime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QTimer13remainingTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimer", "remainingTime", args)
  }

  return
}

// isActive()
func (this *QTimer) IsActive(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QTimer8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimer", "isActive", args)
  }

  return
}

// metaObject()
func (this *QTimer) MetaObject(args ...interface{}) () {
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
    C.C_ZNK6QTimer10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimer", "metaObject", args)
  }

  return
}

// interval()
func (this *QTimer) Interval(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QTimer8intervalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimer", "interval", args)
  }

  return
}

// singleShot(int, const class QObject *, const char *)
func (this *QTimer) SingleShot_s(args ...interface{}) () {
  // singleShot(int, const class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QTimer10singleShotEiPK7QObjectPKc
    // invoke: void singleShot(int, const class QObject *, const char *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    C.C_ZN6QTimer10singleShotEiPK7QObjectPKc(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTimer", "singleShot", args)
  }

  return
}

// <= body block end
