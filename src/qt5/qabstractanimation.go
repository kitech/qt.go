package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QAbstractAnimation::resume();
extern void _ZN18QAbstractAnimation6resumeEv(void* qthis);
  // proto:  void QAbstractAnimation::QAbstractAnimation(QObject * parent);
extern void* dector_ZN18QAbstractAnimationC1EP7QObject(void* arg0);
extern void _ZN18QAbstractAnimationC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QAbstractAnimation::stop();
extern void _ZN18QAbstractAnimation4stopEv(void* qthis);
  // proto:  void QAbstractAnimation::pause();
extern void _ZN18QAbstractAnimation5pauseEv(void* qthis);
  // proto:  void QAbstractAnimation::QAbstractAnimation(const QAbstractAnimation & );
extern void* dector_ZN18QAbstractAnimationC1ERKS_(void* arg0);
extern void _ZN18QAbstractAnimationC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractAnimation::setLoopCount(int loopCount);
extern void _ZN18QAbstractAnimation12setLoopCountEi(void* qthis, int arg0);
  // proto:  int QAbstractAnimation::currentLoop();
extern void _ZNK18QAbstractAnimation11currentLoopEv(void* qthis);
  // proto:  QAnimationGroup * QAbstractAnimation::group();
extern void _ZNK18QAbstractAnimation5groupEv(void* qthis);
  // proto:  void QAbstractAnimation::setPaused(bool );
extern void _ZN18QAbstractAnimation9setPausedEb(void* qthis, bool arg0);
  // proto:  int QAbstractAnimation::totalDuration();
extern void _ZNK18QAbstractAnimation13totalDurationEv(void* qthis);
  // proto:  int QAbstractAnimation::duration();
extern void _ZNK18QAbstractAnimation8durationEv(void* qthis);
  // proto:  const QMetaObject * QAbstractAnimation::metaObject();
extern void _ZNK18QAbstractAnimation10metaObjectEv(void* qthis);
  // proto:  int QAbstractAnimation::currentLoopTime();
extern void _ZNK18QAbstractAnimation15currentLoopTimeEv(void* qthis);
  // proto:  int QAbstractAnimation::currentTime();
extern void _ZNK18QAbstractAnimation11currentTimeEv(void* qthis);
  // proto:  void QAbstractAnimation::setCurrentTime(int msecs);
extern void _ZN18QAbstractAnimation14setCurrentTimeEi(void* qthis, int arg0);
  // proto:  void QAbstractAnimation::~QAbstractAnimation();
extern void _ZN18QAbstractAnimationD0Ev(void* qthis);
  // proto:  int QAbstractAnimation::loopCount();
extern void _ZNK18QAbstractAnimation9loopCountEv(void* qthis);
  // proto:  void QAnimationDriver::advance();
extern void _ZN16QAnimationDriver7advanceEv(void* qthis);
  // proto:  void QAnimationDriver::~QAnimationDriver();
extern void _ZN16QAnimationDriverD0Ev(void* qthis);
  // proto:  void QAnimationDriver::QAnimationDriver(QObject * parent);
extern void* dector_ZN16QAnimationDriverC1EP7QObject(void* arg0);
extern void _ZN16QAnimationDriverC1EP7QObject(void* qthis, void* arg0);
  // proto:  qint64 QAnimationDriver::elapsed();
extern void _ZNK16QAnimationDriver7elapsedEv(void* qthis);
  // proto:  void QAnimationDriver::install();
extern void _ZN16QAnimationDriver7installEv(void* qthis);
  // proto:  const QMetaObject * QAnimationDriver::metaObject();
extern void _ZNK16QAnimationDriver10metaObjectEv(void* qthis);
  // proto:  void QAnimationDriver::uninstall();
extern void _ZN16QAnimationDriver9uninstallEv(void* qthis);
  // proto:  bool QAnimationDriver::isRunning();
extern void _ZNK16QAnimationDriver9isRunningEv(void* qthis);
  // proto:  qint64 QAnimationDriver::startTime();
extern void _ZNK16QAnimationDriver9startTimeEv(void* qthis);
  // proto:  void QAnimationDriver::setStartTime(qint64 startTime);
extern void _ZN16QAnimationDriver12setStartTimeEx(void* qthis, long long arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _currentLoopChanged QAbstractAnimation_currentLoopChanged_signal;
//  _finished QAbstractAnimation_finished_signal;
//  _stateChanged QAbstractAnimation_stateChanged_signal;
//  _directionChanged QAbstractAnimation_directionChanged_signal;
}

// class sizeof(QAnimationDriver)=1
type QAnimationDriver struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _started QAnimationDriver_started_signal;
//  _stopped QAnimationDriver_stopped_signal;
}

  // proto:  void QAbstractAnimation::resume();
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
    C._ZN18QAbstractAnimation6resumeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "resume", args)
  }

}

  // proto:  void QAbstractAnimation::QAbstractAnimation(QObject * parent);
func NewQAbstractAnimation(args ...interface{}) QAbstractAnimation {
  return QAbstractAnimation{}
}

  // proto:  void QAbstractAnimation::stop();
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
    C._ZN18QAbstractAnimation4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "stop", args)
  }

}

  // proto:  void QAbstractAnimation::pause();
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
    C._ZN18QAbstractAnimation5pauseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "pause", args)
  }

}

  // proto:  void QAbstractAnimation::setLoopCount(int loopCount);
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
    C._ZN18QAbstractAnimation12setLoopCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setLoopCount", args)
  }

}

  // proto:  int QAbstractAnimation::currentLoop();
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
    C._ZNK18QAbstractAnimation11currentLoopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoop", args)
  }

}

  // proto:  QAnimationGroup * QAbstractAnimation::group();
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
    C._ZNK18QAbstractAnimation5groupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "group", args)
  }

}

  // proto:  void QAbstractAnimation::setPaused(bool );
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN18QAbstractAnimation9setPausedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setPaused", args)
  }

}

  // proto:  int QAbstractAnimation::totalDuration();
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
    C._ZNK18QAbstractAnimation13totalDurationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "totalDuration", args)
  }

}

  // proto:  int QAbstractAnimation::duration();
func (this *QAbstractAnimation) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation8durationEv
    // invoke: int duration()
    C._ZNK18QAbstractAnimation8durationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "duration", args)
  }

}

  // proto:  const QMetaObject * QAbstractAnimation::metaObject();
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
    C._ZNK18QAbstractAnimation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "metaObject", args)
  }

}

  // proto:  int QAbstractAnimation::currentLoopTime();
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
    C._ZNK18QAbstractAnimation15currentLoopTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoopTime", args)
  }

}

  // proto:  int QAbstractAnimation::currentTime();
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
    C._ZNK18QAbstractAnimation11currentTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentTime", args)
  }

}

  // proto:  void QAbstractAnimation::setCurrentTime(int msecs);
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
    C._ZN18QAbstractAnimation14setCurrentTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setCurrentTime", args)
  }

}

  // proto:  void QAbstractAnimation::~QAbstractAnimation();
func (this *QAbstractAnimation) FreeQAbstractAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "~QAbstractAnimation", args)
  }

}

  // proto:  int QAbstractAnimation::loopCount();
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
    C._ZNK18QAbstractAnimation9loopCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "loopCount", args)
  }

}

  // proto:  void QAnimationDriver::advance();
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
    C._ZN16QAnimationDriver7advanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "advance", args)
  }

}

  // proto:  void QAnimationDriver::~QAnimationDriver();
func (this *QAnimationDriver) FreeQAnimationDriver(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAnimationDriver", "~QAnimationDriver", args)
  }

}

  // proto:  void QAnimationDriver::QAnimationDriver(QObject * parent);
func NewQAnimationDriver(args ...interface{}) QAnimationDriver {
  return QAnimationDriver{}
}

  // proto:  qint64 QAnimationDriver::elapsed();
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
    C._ZNK16QAnimationDriver7elapsedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "elapsed", args)
  }

}

  // proto:  void QAnimationDriver::install();
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
    C._ZN16QAnimationDriver7installEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "install", args)
  }

}

  // proto:  const QMetaObject * QAnimationDriver::metaObject();
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
    C._ZNK16QAnimationDriver10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "metaObject", args)
  }

}

  // proto:  void QAnimationDriver::uninstall();
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
    C._ZN16QAnimationDriver9uninstallEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "uninstall", args)
  }

}

  // proto:  bool QAnimationDriver::isRunning();
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
    C._ZNK16QAnimationDriver9isRunningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "isRunning", args)
  }

}

  // proto:  qint64 QAnimationDriver::startTime();
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
    C._ZNK16QAnimationDriver9startTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "startTime", args)
  }

}

  // proto:  void QAnimationDriver::setStartTime(qint64 startTime);
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
    C._ZN16QAnimationDriver12setStartTimeEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationDriver", "setStartTime", args)
  }

}

// <= body block end

