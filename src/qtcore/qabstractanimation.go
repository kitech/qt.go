package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qabstractanimation.h
// dst-file: /src/core/qabstractanimation.go
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
  // proto:  int QAbstractAnimation::currentLoop();
extern int32_t C_ZNK18QAbstractAnimation11currentLoopEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::currentLoopTime();
extern int32_t C_ZNK18QAbstractAnimation15currentLoopTimeEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::setCurrentTime(int msecs);
extern void C_ZN18QAbstractAnimation14setCurrentTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractAnimation::~QAbstractAnimation();
extern void C_ZN18QAbstractAnimationD2Ev(void* qthis); // 4
  // proto:  void QAbstractAnimation::pause();
extern void C_ZN18QAbstractAnimation5pauseEv(void* qthis); // 4
  // proto:  QAnimationGroup * QAbstractAnimation::group();
extern void* C_ZNK18QAbstractAnimation5groupEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::loopCount();
extern int32_t C_ZNK18QAbstractAnimation9loopCountEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::QAbstractAnimation(QObject * parent);
extern void* C_ZN18QAbstractAnimationC2EP7QObject(void* arg0); // 3
  // proto:  QAbstractAnimation::State QAbstractAnimation::state();
extern void C_ZNK18QAbstractAnimation5stateEv(void* qthis); // 4
  // proto:  QAbstractAnimation::Direction QAbstractAnimation::direction();
extern void C_ZNK18QAbstractAnimation9directionEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::setLoopCount(int loopCount);
extern void C_ZN18QAbstractAnimation12setLoopCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractAnimation::resume();
extern void C_ZN18QAbstractAnimation6resumeEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::stop();
extern void C_ZN18QAbstractAnimation4stopEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractAnimation::metaObject();
extern void C_ZNK18QAbstractAnimation10metaObjectEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::currentTime();
extern int32_t C_ZNK18QAbstractAnimation11currentTimeEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::totalDuration();
extern int32_t C_ZNK18QAbstractAnimation13totalDurationEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::setPaused(bool );
extern void C_ZN18QAbstractAnimation9setPausedEb(void* qthis, bool arg0); // 4
  // proto:  void QAnimationDriver::advance();
extern void C_ZN16QAnimationDriver7advanceEv(void* qthis); // 4
  // proto:  void QAnimationDriver::QAnimationDriver(QObject * parent);
extern void* C_ZN16QAnimationDriverC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QAnimationDriver::metaObject();
extern void C_ZNK16QAnimationDriver10metaObjectEv(void* qthis); // 4
  // proto:  qint64 QAnimationDriver::startTime();
extern int64_t C_ZNK16QAnimationDriver9startTimeEv(void* qthis); // 4
  // proto:  qint64 QAnimationDriver::elapsed();
extern int64_t C_ZNK16QAnimationDriver7elapsedEv(void* qthis); // 4
  // proto:  void QAnimationDriver::setStartTime(qint64 startTime);
extern void C_ZN16QAnimationDriver12setStartTimeEx(void* qthis, int64_t arg0); // 4
  // proto:  void QAnimationDriver::install();
extern void C_ZN16QAnimationDriver7installEv(void* qthis); // 4
  // proto:  void QAnimationDriver::~QAnimationDriver();
extern void C_ZN16QAnimationDriverD2Ev(void* qthis); // 4
  // proto:  bool QAnimationDriver::isRunning();
extern bool C_ZNK16QAnimationDriver9isRunningEv(void* qthis); // 4
  // proto:  void QAnimationDriver::uninstall();
extern void C_ZN16QAnimationDriver9uninstallEv(void* qthis); // 4
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

// class sizeof(QAbstractAnimation)=1
type QAbstractAnimation struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _currentLoopChanged QAbstractAnimation_currentLoopChanged_signal;
//  _finished QAbstractAnimation_finished_signal;
//  _stateChanged QAbstractAnimation_stateChanged_signal;
//  _directionChanged QAbstractAnimation_directionChanged_signal;
}

// class sizeof(QAnimationDriver)=1
type QAnimationDriver struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _started QAnimationDriver_started_signal;
//  _stopped QAnimationDriver_stopped_signal;
}

// currentLoop()
func (this *QAbstractAnimation) Currentloop(args ...interface{}) (ret interface{}) {
  // currentLoop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation11currentLoopEv
    // invoke: int currentLoop()
    var ret0 = C.C_ZNK18QAbstractAnimation11currentLoopEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoop", args)
  }

  return
}

// currentLoopTime()
func (this *QAbstractAnimation) Currentlooptime(args ...interface{}) (ret interface{}) {
  // currentLoopTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation15currentLoopTimeEv
    // invoke: int currentLoopTime()
    var ret0 = C.C_ZNK18QAbstractAnimation15currentLoopTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoopTime", args)
  }

  return
}

// setCurrentTime(int)
func (this *QAbstractAnimation) Setcurrenttime(args ...interface{}) () {
  // setCurrentTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation14setCurrentTimeEi
    // invoke: void setCurrentTime(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QAbstractAnimation14setCurrentTimeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setCurrentTime", args)
  }

  return
}

// ~QAbstractAnimation()
func (this *QAbstractAnimation) Freeqabstractanimation(args ...interface{}) () {
  // ~QAbstractAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimationD0Ev
    // invoke: void ~QAbstractAnimation()
    C.C_ZN18QAbstractAnimationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "~QAbstractAnimation", args)
  }

  return
}

// pause()
func (this *QAbstractAnimation) Pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation5pauseEv
    // invoke: void pause()
    C.C_ZN18QAbstractAnimation5pauseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "pause", args)
  }

  return
}

// group()
func (this *QAbstractAnimation) Group(args ...interface{}) (ret interface{}) {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation5groupEv
    // invoke: QAnimationGroup * group()
    var ret0 = C.C_ZNK18QAbstractAnimation5groupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAnimationGroup{}) // "QAnimationGroup *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "group", args)
  }

  return
}

// loopCount()
func (this *QAbstractAnimation) Loopcount(args ...interface{}) (ret interface{}) {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation9loopCountEv
    // invoke: int loopCount()
    var ret0 = C.C_ZNK18QAbstractAnimation9loopCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "loopCount", args)
  }

  return
}

// QAbstractAnimation(class QObject *)
func NewQAbstractAnimation(args ...interface{}) *QAbstractAnimation {
  // QAbstractAnimation(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimationC1EP7QObject
    // invoke: void QAbstractAnimation(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QAbstractAnimationC2EP7QObject(arg0)
    return &QAbstractAnimation{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "QAbstractAnimation", args)
  }

  return nil // QAbstractAnimation{Qclsinst:qthis}
}

// state()
func (this *QAbstractAnimation) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation5stateEv
    // invoke: QAbstractAnimation::State state()
    C.C_ZNK18QAbstractAnimation5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "state", args)
  }

  return
}

// direction()
func (this *QAbstractAnimation) Direction(args ...interface{}) () {
  // direction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation9directionEv
    // invoke: QAbstractAnimation::Direction direction()
    C.C_ZNK18QAbstractAnimation9directionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "direction", args)
  }

  return
}

// setLoopCount(int)
func (this *QAbstractAnimation) Setloopcount(args ...interface{}) () {
  // setLoopCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation12setLoopCountEi
    // invoke: void setLoopCount(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QAbstractAnimation12setLoopCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setLoopCount", args)
  }

  return
}

// resume()
func (this *QAbstractAnimation) Resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation6resumeEv
    // invoke: void resume()
    C.C_ZN18QAbstractAnimation6resumeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "resume", args)
  }

  return
}

// stop()
func (this *QAbstractAnimation) Stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation4stopEv
    // invoke: void stop()
    C.C_ZN18QAbstractAnimation4stopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "stop", args)
  }

  return
}

// metaObject()
func (this *QAbstractAnimation) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QAbstractAnimation10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "metaObject", args)
  }

  return
}

// currentTime()
func (this *QAbstractAnimation) Currenttime(args ...interface{}) (ret interface{}) {
  // currentTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation11currentTimeEv
    // invoke: int currentTime()
    var ret0 = C.C_ZNK18QAbstractAnimation11currentTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentTime", args)
  }

  return
}

// totalDuration()
func (this *QAbstractAnimation) Totalduration(args ...interface{}) (ret interface{}) {
  // totalDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation13totalDurationEv
    // invoke: int totalDuration()
    var ret0 = C.C_ZNK18QAbstractAnimation13totalDurationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "totalDuration", args)
  }

  return
}

// setPaused(_Bool)
func (this *QAbstractAnimation) Setpaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation9setPausedEb
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN18QAbstractAnimation9setPausedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setPaused", args)
  }

  return
}

// advance()
func (this *QAnimationDriver) Advance(args ...interface{}) () {
  // advance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver7advanceEv
    // invoke: void advance()
    C.C_ZN16QAnimationDriver7advanceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "advance", args)
  }

  return
}

// QAnimationDriver(class QObject *)
func NewQAnimationDriver(args ...interface{}) *QAnimationDriver {
  // QAnimationDriver(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriverC1EP7QObject
    // invoke: void QAnimationDriver(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QAnimationDriverC2EP7QObject(arg0)
    return &QAnimationDriver{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAnimationDriver", "QAnimationDriver", args)
  }

  return nil // QAnimationDriver{Qclsinst:qthis}
}

// metaObject()
func (this *QAnimationDriver) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QAnimationDriver10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "metaObject", args)
  }

  return
}

// startTime()
func (this *QAnimationDriver) Starttime(args ...interface{}) (ret interface{}) {
  // startTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver9startTimeEv
    // invoke: qint64 startTime()
    var ret0 = C.C_ZNK16QAnimationDriver9startTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAnimationDriver", "startTime", args)
  }

  return
}

// elapsed()
func (this *QAnimationDriver) Elapsed(args ...interface{}) (ret interface{}) {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver7elapsedEv
    // invoke: qint64 elapsed()
    var ret0 = C.C_ZNK16QAnimationDriver7elapsedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAnimationDriver", "elapsed", args)
  }

  return
}

// setStartTime(qint64)
func (this *QAnimationDriver) Setstarttime(args ...interface{}) () {
  // setStartTime(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver12setStartTimeEx
    // invoke: void setStartTime(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAnimationDriver12setStartTimeEx(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "setStartTime", args)
  }

  return
}

// install()
func (this *QAnimationDriver) Install(args ...interface{}) () {
  // install()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver7installEv
    // invoke: void install()
    C.C_ZN16QAnimationDriver7installEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "install", args)
  }

  return
}

// ~QAnimationDriver()
func (this *QAnimationDriver) Freeqanimationdriver(args ...interface{}) () {
  // ~QAnimationDriver()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriverD0Ev
    // invoke: void ~QAnimationDriver()
    C.C_ZN16QAnimationDriverD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "~QAnimationDriver", args)
  }

  return
}

// isRunning()
func (this *QAnimationDriver) Isrunning(args ...interface{}) (ret interface{}) {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver9isRunningEv
    // invoke: bool isRunning()
    var ret0 = C.C_ZNK16QAnimationDriver9isRunningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAnimationDriver", "isRunning", args)
  }

  return
}

// uninstall()
func (this *QAnimationDriver) Uninstall(args ...interface{}) () {
  // uninstall()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver9uninstallEv
    // invoke: void uninstall()
    C.C_ZN16QAnimationDriver9uninstallEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "uninstall", args)
  }

  return
}

// <= body block end

