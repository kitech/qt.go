package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK18QAbstractAnimation11currentLoopEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::currentLoopTime();
extern void C_ZNK18QAbstractAnimation15currentLoopTimeEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::setCurrentTime(int msecs);
extern void C_ZN18QAbstractAnimation14setCurrentTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractAnimation::~QAbstractAnimation();
extern void C_ZN18QAbstractAnimationD2Ev(void* qthis); // 4
  // proto:  void QAbstractAnimation::pause();
extern void C_ZN18QAbstractAnimation5pauseEv(void* qthis); // 4
  // proto:  QAnimationGroup * QAbstractAnimation::group();
extern void C_ZNK18QAbstractAnimation5groupEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::loopCount();
extern void C_ZNK18QAbstractAnimation9loopCountEv(void* qthis); // 4
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
extern void C_ZNK18QAbstractAnimation11currentTimeEv(void* qthis); // 4
  // proto:  int QAbstractAnimation::totalDuration();
extern void C_ZNK18QAbstractAnimation13totalDurationEv(void* qthis); // 4
  // proto:  void QAbstractAnimation::setPaused(bool );
extern void C_ZN18QAbstractAnimation9setPausedEb(void* qthis, bool arg0); // 4
  // proto:  void QAnimationDriver::advance();
extern void C_ZN16QAnimationDriver7advanceEv(void* qthis); // 4
  // proto:  void QAnimationDriver::QAnimationDriver(QObject * parent);
extern void* C_ZN16QAnimationDriverC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QAnimationDriver::metaObject();
extern void C_ZNK16QAnimationDriver10metaObjectEv(void* qthis); // 4
  // proto:  qint64 QAnimationDriver::startTime();
extern void C_ZNK16QAnimationDriver9startTimeEv(void* qthis); // 4
  // proto:  qint64 QAnimationDriver::elapsed();
extern void C_ZNK16QAnimationDriver7elapsedEv(void* qthis); // 4
  // proto:  void QAnimationDriver::setStartTime(qint64 startTime);
extern void C_ZN16QAnimationDriver12setStartTimeEx(void* qthis, int64_t arg0); // 4
  // proto:  void QAnimationDriver::install();
extern void C_ZN16QAnimationDriver7installEv(void* qthis); // 4
  // proto:  void QAnimationDriver::~QAnimationDriver();
extern void C_ZN16QAnimationDriverD2Ev(void* qthis); // 4
  // proto:  bool QAnimationDriver::isRunning();
extern void C_ZNK16QAnimationDriver9isRunningEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentLoopChanged QAbstractAnimation_currentLoopChanged_signal;
//  _finished QAbstractAnimation_finished_signal;
//  _stateChanged QAbstractAnimation_stateChanged_signal;
//  _directionChanged QAbstractAnimation_directionChanged_signal;
}

// class sizeof(QAnimationDriver)=1
type QAnimationDriver struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _started QAnimationDriver_started_signal;
//  _stopped QAnimationDriver_stopped_signal;
}

// currentLoop()
func (this *QAbstractAnimation) currentLoop(args ...interface{}) () {
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
    var ret = C.C_ZNK18QAbstractAnimation11currentLoopEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoop", args)
  }

}

// currentLoopTime()
func (this *QAbstractAnimation) currentLoopTime(args ...interface{}) () {
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
    var ret = C.C_ZNK18QAbstractAnimation15currentLoopTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoopTime", args)
  }

}

// setCurrentTime(int)
func (this *QAbstractAnimation) setCurrentTime(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QAbstractAnimation14setCurrentTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setCurrentTime", args)
  }

}

// ~QAbstractAnimation()
func (this *QAbstractAnimation) FreeQAbstractAnimation(args ...interface{}) () {
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
    C.C_ZN18QAbstractAnimationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "~QAbstractAnimation", args)
  }

}

// pause()
func (this *QAbstractAnimation) pause(args ...interface{}) () {
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
    C.C_ZN18QAbstractAnimation5pauseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "pause", args)
  }

}

// group()
func (this *QAbstractAnimation) group(args ...interface{}) () {
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
    var ret = C.C_ZNK18QAbstractAnimation5groupEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "group", args)
  }

}

// loopCount()
func (this *QAbstractAnimation) loopCount(args ...interface{}) () {
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
    var ret = C.C_ZNK18QAbstractAnimation9loopCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "loopCount", args)
  }

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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QAbstractAnimationC2EP7QObject(arg0)
    return &QAbstractAnimation{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "QAbstractAnimation", args)
  }

  return nil // QAbstractAnimation{qclsinst:qthis}
}

// state()
func (this *QAbstractAnimation) state(args ...interface{}) () {
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
    C.C_ZNK18QAbstractAnimation5stateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "state", args)
  }

}

// direction()
func (this *QAbstractAnimation) direction(args ...interface{}) () {
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
    C.C_ZNK18QAbstractAnimation9directionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "direction", args)
  }

}

// setLoopCount(int)
func (this *QAbstractAnimation) setLoopCount(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QAbstractAnimation12setLoopCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setLoopCount", args)
  }

}

// resume()
func (this *QAbstractAnimation) resume(args ...interface{}) () {
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
    C.C_ZN18QAbstractAnimation6resumeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "resume", args)
  }

}

// stop()
func (this *QAbstractAnimation) stop(args ...interface{}) () {
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
    C.C_ZN18QAbstractAnimation4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "stop", args)
  }

}

// metaObject()
func (this *QAbstractAnimation) metaObject(args ...interface{}) () {
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
    C.C_ZNK18QAbstractAnimation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "metaObject", args)
  }

}

// currentTime()
func (this *QAbstractAnimation) currentTime(args ...interface{}) () {
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
    var ret = C.C_ZNK18QAbstractAnimation11currentTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentTime", args)
  }

}

// totalDuration()
func (this *QAbstractAnimation) totalDuration(args ...interface{}) () {
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
    var ret = C.C_ZNK18QAbstractAnimation13totalDurationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "totalDuration", args)
  }

}

// setPaused(_Bool)
func (this *QAbstractAnimation) setPaused(args ...interface{}) () {
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
    C.C_ZN18QAbstractAnimation9setPausedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setPaused", args)
  }

}

// advance()
func (this *QAnimationDriver) advance(args ...interface{}) () {
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
    C.C_ZN16QAnimationDriver7advanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "advance", args)
  }

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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QAnimationDriverC2EP7QObject(arg0)
    return &QAnimationDriver{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAnimationDriver", "QAnimationDriver", args)
  }

  return nil // QAnimationDriver{qclsinst:qthis}
}

// metaObject()
func (this *QAnimationDriver) metaObject(args ...interface{}) () {
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
    C.C_ZNK16QAnimationDriver10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "metaObject", args)
  }

}

// startTime()
func (this *QAnimationDriver) startTime(args ...interface{}) () {
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
    var ret = C.C_ZNK16QAnimationDriver9startTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAnimationDriver", "startTime", args)
  }

}

// elapsed()
func (this *QAnimationDriver) elapsed(args ...interface{}) () {
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
    var ret = C.C_ZNK16QAnimationDriver7elapsedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAnimationDriver", "elapsed", args)
  }

}

// setStartTime(qint64)
func (this *QAnimationDriver) setStartTime(args ...interface{}) () {
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAnimationDriver12setStartTimeEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "setStartTime", args)
  }

}

// install()
func (this *QAnimationDriver) install(args ...interface{}) () {
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
    C.C_ZN16QAnimationDriver7installEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "install", args)
  }

}

// ~QAnimationDriver()
func (this *QAnimationDriver) FreeQAnimationDriver(args ...interface{}) () {
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
    C.C_ZN16QAnimationDriverD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "~QAnimationDriver", args)
  }

}

// isRunning()
func (this *QAnimationDriver) isRunning(args ...interface{}) () {
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
    var ret = C.C_ZNK16QAnimationDriver9isRunningEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAnimationDriver", "isRunning", args)
  }

}

// uninstall()
func (this *QAnimationDriver) uninstall(args ...interface{}) () {
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
    C.C_ZN16QAnimationDriver9uninstallEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "uninstall", args)
  }

}

// <= body block end

