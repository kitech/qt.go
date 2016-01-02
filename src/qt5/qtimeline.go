package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qtimeline.h
// dst-file: /src/core/qtimeline.go
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
  // proto:  void QTimeLine::start();
extern void _ZN9QTimeLine5startEv(void* qthis);
  // proto:  void QTimeLine::QTimeLine(const QTimeLine & );
extern void* dector_ZN9QTimeLineC1ERKS_(void* arg0);
extern void _ZN9QTimeLineC1ERKS_(void* qthis, void* arg0);
  // proto:  int QTimeLine::duration();
extern void _ZNK9QTimeLine8durationEv(void* qthis);
  // proto:  int QTimeLine::currentFrame();
extern void _ZNK9QTimeLine12currentFrameEv(void* qthis);
  // proto:  const QMetaObject * QTimeLine::metaObject();
extern void _ZNK9QTimeLine10metaObjectEv(void* qthis);
  // proto:  void QTimeLine::stop();
extern void _ZN9QTimeLine4stopEv(void* qthis);
  // proto:  void QTimeLine::~QTimeLine();
extern void _ZN9QTimeLineD0Ev(void* qthis);
  // proto:  void QTimeLine::setUpdateInterval(int interval);
extern void _ZN9QTimeLine17setUpdateIntervalEi(void* qthis, int arg0);
  // proto:  QEasingCurve QTimeLine::easingCurve();
extern void _ZNK9QTimeLine11easingCurveEv(void* qthis);
  // proto:  int QTimeLine::loopCount();
extern void _ZNK9QTimeLine9loopCountEv(void* qthis);
  // proto:  void QTimeLine::setStartFrame(int frame);
extern void _ZN9QTimeLine13setStartFrameEi(void* qthis, int arg0);
  // proto:  void QTimeLine::QTimeLine(int duration, QObject * parent);
extern void* dector_ZN9QTimeLineC1EiP7QObject(int arg0, void* arg1);
extern void _ZN9QTimeLineC1EiP7QObject(void* qthis, int arg0, void* arg1);
  // proto:  void QTimeLine::resume();
extern void _ZN9QTimeLine6resumeEv(void* qthis);
  // proto:  void QTimeLine::setEasingCurve(const QEasingCurve & curve);
extern void _ZN9QTimeLine14setEasingCurveERK12QEasingCurve(void* qthis, void* arg0);
  // proto:  int QTimeLine::startFrame();
extern void _ZNK9QTimeLine10startFrameEv(void* qthis);
  // proto:  void QTimeLine::setEndFrame(int frame);
extern void _ZN9QTimeLine11setEndFrameEi(void* qthis, int arg0);
  // proto:  int QTimeLine::updateInterval();
extern void _ZNK9QTimeLine14updateIntervalEv(void* qthis);
  // proto:  void QTimeLine::setLoopCount(int count);
extern void _ZN9QTimeLine12setLoopCountEi(void* qthis, int arg0);
  // proto:  void QTimeLine::setCurrentTime(int msec);
extern void _ZN9QTimeLine14setCurrentTimeEi(void* qthis, int arg0);
  // proto:  int QTimeLine::currentTime();
extern void _ZNK9QTimeLine11currentTimeEv(void* qthis);
  // proto:  void QTimeLine::setDuration(int duration);
extern void _ZN9QTimeLine11setDurationEi(void* qthis, int arg0);
  // proto:  void QTimeLine::toggleDirection();
extern void _ZN9QTimeLine15toggleDirectionEv(void* qthis);
  // proto:  int QTimeLine::endFrame();
extern void _ZNK9QTimeLine8endFrameEv(void* qthis);
  // proto:  void QTimeLine::setPaused(bool paused);
extern void _ZN9QTimeLine9setPausedEb(void* qthis, bool arg0);
  // proto:  int QTimeLine::frameForTime(int msec);
extern void _ZNK9QTimeLine12frameForTimeEi(void* qthis, int arg0);
  // proto:  void QTimeLine::setFrameRange(int startFrame, int endFrame);
extern void _ZN9QTimeLine13setFrameRangeEii(void* qthis, int arg0, int arg1);
  // proto:  qreal QTimeLine::valueForTime(int msec);
extern void _ZNK9QTimeLine12valueForTimeEi(void* qthis, int arg0);
  // proto:  qreal QTimeLine::currentValue();
extern void _ZNK9QTimeLine12currentValueEv(void* qthis);
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

// class sizeof(QTimeLine)=1
type QTimeLine struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _valueChanged QTimeLine_valueChanged_signal;
//  _finished QTimeLine_finished_signal;
//  _frameChanged QTimeLine_frameChanged_signal;
//  _stateChanged QTimeLine_stateChanged_signal;
}

  // proto:  void QTimeLine::start();
func (this *QTimeLine) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine5startEv
  default:
    qtrt.ErrorResolve("QTimeLine", "start", args)
  }

}

  // proto:  void QTimeLine::QTimeLine(const QTimeLine & );
func NewQTimeLine(args ...interface{}) QTimeLine {
  return QTimeLine{}
}

  // proto:  int QTimeLine::duration();
func (this *QTimeLine) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine8durationEv
  default:
    qtrt.ErrorResolve("QTimeLine", "duration", args)
  }

}

  // proto:  int QTimeLine::currentFrame();
func (this *QTimeLine) currentFrame(args ...interface{}) () {
  // currentFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine12currentFrameEv
  default:
    qtrt.ErrorResolve("QTimeLine", "currentFrame", args)
  }

}

  // proto:  const QMetaObject * QTimeLine::metaObject();
func (this *QTimeLine) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine10metaObjectEv
  default:
    qtrt.ErrorResolve("QTimeLine", "metaObject", args)
  }

}

  // proto:  void QTimeLine::stop();
func (this *QTimeLine) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine4stopEv
  default:
    qtrt.ErrorResolve("QTimeLine", "stop", args)
  }

}

  // proto:  void QTimeLine::~QTimeLine();
func (this *QTimeLine) FreeQTimeLine(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeLine", "~QTimeLine", args)
  }

}

  // proto:  void QTimeLine::setUpdateInterval(int interval);
func (this *QTimeLine) setUpdateInterval(args ...interface{}) () {
  // setUpdateInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine17setUpdateIntervalEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setUpdateInterval", args)
  }

}

  // proto:  QEasingCurve QTimeLine::easingCurve();
func (this *QTimeLine) easingCurve(args ...interface{}) () {
  // easingCurve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine11easingCurveEv
  default:
    qtrt.ErrorResolve("QTimeLine", "easingCurve", args)
  }

}

  // proto:  int QTimeLine::loopCount();
func (this *QTimeLine) loopCount(args ...interface{}) () {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine9loopCountEv
  default:
    qtrt.ErrorResolve("QTimeLine", "loopCount", args)
  }

}

  // proto:  void QTimeLine::setStartFrame(int frame);
func (this *QTimeLine) setStartFrame(args ...interface{}) () {
  // setStartFrame(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine13setStartFrameEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setStartFrame", args)
  }

}

  // proto:  void QTimeLine::resume();
func (this *QTimeLine) resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine6resumeEv
  default:
    qtrt.ErrorResolve("QTimeLine", "resume", args)
  }

}

  // proto:  void QTimeLine::setEasingCurve(const QEasingCurve & curve);
func (this *QTimeLine) setEasingCurve(args ...interface{}) () {
  // setEasingCurve(const class QEasingCurve &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEasingCurve{}) // "const QEasingCurve &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine14setEasingCurveERK12QEasingCurve
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setEasingCurve", args)
  }

}

  // proto:  int QTimeLine::startFrame();
func (this *QTimeLine) startFrame(args ...interface{}) () {
  // startFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine10startFrameEv
  default:
    qtrt.ErrorResolve("QTimeLine", "startFrame", args)
  }

}

  // proto:  void QTimeLine::setEndFrame(int frame);
func (this *QTimeLine) setEndFrame(args ...interface{}) () {
  // setEndFrame(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine11setEndFrameEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setEndFrame", args)
  }

}

  // proto:  int QTimeLine::updateInterval();
func (this *QTimeLine) updateInterval(args ...interface{}) () {
  // updateInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine14updateIntervalEv
  default:
    qtrt.ErrorResolve("QTimeLine", "updateInterval", args)
  }

}

  // proto:  void QTimeLine::setLoopCount(int count);
func (this *QTimeLine) setLoopCount(args ...interface{}) () {
  // setLoopCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine12setLoopCountEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setLoopCount", args)
  }

}

  // proto:  void QTimeLine::setCurrentTime(int msec);
func (this *QTimeLine) setCurrentTime(args ...interface{}) () {
  // setCurrentTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine14setCurrentTimeEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setCurrentTime", args)
  }

}

  // proto:  int QTimeLine::currentTime();
func (this *QTimeLine) currentTime(args ...interface{}) () {
  // currentTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine11currentTimeEv
  default:
    qtrt.ErrorResolve("QTimeLine", "currentTime", args)
  }

}

  // proto:  void QTimeLine::setDuration(int duration);
func (this *QTimeLine) setDuration(args ...interface{}) () {
  // setDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine11setDurationEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setDuration", args)
  }

}

  // proto:  void QTimeLine::toggleDirection();
func (this *QTimeLine) toggleDirection(args ...interface{}) () {
  // toggleDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine15toggleDirectionEv
  default:
    qtrt.ErrorResolve("QTimeLine", "toggleDirection", args)
  }

}

  // proto:  int QTimeLine::endFrame();
func (this *QTimeLine) endFrame(args ...interface{}) () {
  // endFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine8endFrameEv
  default:
    qtrt.ErrorResolve("QTimeLine", "endFrame", args)
  }

}

  // proto:  void QTimeLine::setPaused(bool paused);
func (this *QTimeLine) setPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine9setPausedEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setPaused", args)
  }

}

  // proto:  int QTimeLine::frameForTime(int msec);
func (this *QTimeLine) frameForTime(args ...interface{}) () {
  // frameForTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine12frameForTimeEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "frameForTime", args)
  }

}

  // proto:  void QTimeLine::setFrameRange(int startFrame, int endFrame);
func (this *QTimeLine) setFrameRange(args ...interface{}) () {
  // setFrameRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine13setFrameRangeEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTimeLine", "setFrameRange", args)
  }

}

  // proto:  qreal QTimeLine::valueForTime(int msec);
func (this *QTimeLine) valueForTime(args ...interface{}) () {
  // valueForTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine12valueForTimeEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTimeLine", "valueForTime", args)
  }

}

  // proto:  qreal QTimeLine::currentValue();
func (this *QTimeLine) currentValue(args ...interface{}) () {
  // currentValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine12currentValueEv
  default:
    qtrt.ErrorResolve("QTimeLine", "currentValue", args)
  }

}

// <= body block end

