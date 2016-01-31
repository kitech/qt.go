package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTimeLine::toggleDirection();
extern void C_ZN9QTimeLine15toggleDirectionEv(void* qthis); // 4
  // proto:  void QTimeLine::setCurrentTime(int msec);
extern void C_ZN9QTimeLine14setCurrentTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QTimeLine::valueForTime(int msec);
extern double C_ZNK9QTimeLine12valueForTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTimeLine::setEasingCurve(const QEasingCurve & curve);
extern void C_ZN9QTimeLine14setEasingCurveERK12QEasingCurve(void* qthis, void* arg0); // 4
  // proto:  qreal QTimeLine::currentValue();
extern double C_ZNK9QTimeLine12currentValueEv(void* qthis); // 4
  // proto:  void QTimeLine::QTimeLine(int duration, QObject * parent);
extern void* C_ZN9QTimeLineC2EiP7QObject(int32_t arg0, void* arg1); // 3
  // proto:  void QTimeLine::setFrameRange(int startFrame, int endFrame);
extern void C_ZN9QTimeLine13setFrameRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTimeLine::duration();
extern int32_t C_ZNK9QTimeLine8durationEv(void* qthis); // 4
  // proto:  int QTimeLine::loopCount();
extern int32_t C_ZNK9QTimeLine9loopCountEv(void* qthis); // 4
  // proto:  void QTimeLine::resume();
extern void C_ZN9QTimeLine6resumeEv(void* qthis); // 4
  // proto:  void QTimeLine::start();
extern void C_ZN9QTimeLine5startEv(void* qthis); // 4
  // proto:  QTimeLine::State QTimeLine::state();
extern void C_ZNK9QTimeLine5stateEv(void* qthis); // 4
  // proto:  int QTimeLine::frameForTime(int msec);
extern int32_t C_ZNK9QTimeLine12frameForTimeEi(void* qthis, int32_t arg0); // 4
  // proto:  QTimeLine::Direction QTimeLine::direction();
extern void C_ZNK9QTimeLine9directionEv(void* qthis); // 4
  // proto:  void QTimeLine::setLoopCount(int count);
extern void C_ZN9QTimeLine12setLoopCountEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTimeLine::endFrame();
extern int32_t C_ZNK9QTimeLine8endFrameEv(void* qthis); // 4
  // proto:  void QTimeLine::stop();
extern void C_ZN9QTimeLine4stopEv(void* qthis); // 4
  // proto:  int QTimeLine::startFrame();
extern int32_t C_ZNK9QTimeLine10startFrameEv(void* qthis); // 4
  // proto:  int QTimeLine::updateInterval();
extern int32_t C_ZNK9QTimeLine14updateIntervalEv(void* qthis); // 4
  // proto:  int QTimeLine::currentFrame();
extern int32_t C_ZNK9QTimeLine12currentFrameEv(void* qthis); // 4
  // proto:  const QMetaObject * QTimeLine::metaObject();
extern void C_ZNK9QTimeLine10metaObjectEv(void* qthis); // 4
  // proto:  void QTimeLine::setUpdateInterval(int interval);
extern void C_ZN9QTimeLine17setUpdateIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  QTimeLine::CurveShape QTimeLine::curveShape();
extern void C_ZNK9QTimeLine10curveShapeEv(void* qthis); // 4
  // proto:  int QTimeLine::currentTime();
extern int32_t C_ZNK9QTimeLine11currentTimeEv(void* qthis); // 4
  // proto:  QEasingCurve QTimeLine::easingCurve();
extern void* C_ZNK9QTimeLine11easingCurveEv(void* qthis); // 4
  // proto:  void QTimeLine::setEndFrame(int frame);
extern void C_ZN9QTimeLine11setEndFrameEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTimeLine::setDuration(int duration);
extern void C_ZN9QTimeLine11setDurationEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTimeLine::~QTimeLine();
extern void C_ZN9QTimeLineD2Ev(void* qthis); // 4
  // proto:  void QTimeLine::setStartFrame(int frame);
extern void C_ZN9QTimeLine13setStartFrameEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTimeLine::setPaused(bool paused);
extern void C_ZN9QTimeLine9setPausedEb(void* qthis, bool arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QTimeLine_valueChanged_signal;
//  _finished QTimeLine_finished_signal;
//  _frameChanged QTimeLine_frameChanged_signal;
//  _stateChanged QTimeLine_stateChanged_signal;
}

// toggleDirection()
func (this *QTimeLine) Toggledirection(args ...interface{}) () {
  // toggleDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine15toggleDirectionEv
    // invoke: void toggleDirection()
    C.C_ZN9QTimeLine15toggleDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "toggleDirection", args)
  }

  return
}

// setCurrentTime(int)
func (this *QTimeLine) Setcurrenttime(args ...interface{}) () {
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
    // invoke: void setCurrentTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine14setCurrentTimeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setCurrentTime", args)
  }

  return
}

// valueForTime(int)
func (this *QTimeLine) Valuefortime(args ...interface{}) (ret interface{}) {
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
    // invoke: qreal valueForTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeLine12valueForTimeEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "valueForTime", args)
  }

  return
}

// setEasingCurve(const class QEasingCurve &)
func (this *QTimeLine) Seteasingcurve(args ...interface{}) () {
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
    // invoke: void setEasingCurve(const class QEasingCurve &)
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine14setEasingCurveERK12QEasingCurve(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setEasingCurve", args)
  }

  return
}

// currentValue()
func (this *QTimeLine) Currentvalue(args ...interface{}) (ret interface{}) {
  // currentValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine12currentValueEv
    // invoke: qreal currentValue()
    var ret0 = C.C_ZNK9QTimeLine12currentValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "currentValue", args)
  }

  return
}

// QTimeLine(int, class QObject *)
func NewQTimeLine(args ...interface{}) *QTimeLine {
  // QTimeLine(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLineC1EiP7QObject
    // invoke: void QTimeLine(int, class QObject *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeLineC2EiP7QObject(arg0, arg1)
    return &QTimeLine{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTimeLine", "QTimeLine", args)
  }

  return nil // QTimeLine{qclsinst:qthis}
}

// setFrameRange(int, int)
func (this *QTimeLine) Setframerange(args ...interface{}) () {
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
    // invoke: void setFrameRange(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QTimeLine13setFrameRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTimeLine", "setFrameRange", args)
  }

  return
}

// duration()
func (this *QTimeLine) Duration(args ...interface{}) (ret interface{}) {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine8durationEv
    // invoke: int duration()
    var ret0 = C.C_ZNK9QTimeLine8durationEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "duration", args)
  }

  return
}

// loopCount()
func (this *QTimeLine) Loopcount(args ...interface{}) (ret interface{}) {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine9loopCountEv
    // invoke: int loopCount()
    var ret0 = C.C_ZNK9QTimeLine9loopCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "loopCount", args)
  }

  return
}

// resume()
func (this *QTimeLine) Resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine6resumeEv
    // invoke: void resume()
    C.C_ZN9QTimeLine6resumeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "resume", args)
  }

  return
}

// start()
func (this *QTimeLine) Start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine5startEv
    // invoke: void start()
    C.C_ZN9QTimeLine5startEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "start", args)
  }

  return
}

// state()
func (this *QTimeLine) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine5stateEv
    // invoke: QTimeLine::State state()
    C.C_ZNK9QTimeLine5stateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "state", args)
  }

  return
}

// frameForTime(int)
func (this *QTimeLine) Framefortime(args ...interface{}) (ret interface{}) {
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
    // invoke: int frameForTime(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeLine12frameForTimeEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "frameForTime", args)
  }

  return
}

// direction()
func (this *QTimeLine) Direction(args ...interface{}) () {
  // direction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine9directionEv
    // invoke: QTimeLine::Direction direction()
    C.C_ZNK9QTimeLine9directionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "direction", args)
  }

  return
}

// setLoopCount(int)
func (this *QTimeLine) Setloopcount(args ...interface{}) () {
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
    // invoke: void setLoopCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine12setLoopCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setLoopCount", args)
  }

  return
}

// endFrame()
func (this *QTimeLine) Endframe(args ...interface{}) (ret interface{}) {
  // endFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine8endFrameEv
    // invoke: int endFrame()
    var ret0 = C.C_ZNK9QTimeLine8endFrameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "endFrame", args)
  }

  return
}

// stop()
func (this *QTimeLine) Stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLine4stopEv
    // invoke: void stop()
    C.C_ZN9QTimeLine4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "stop", args)
  }

  return
}

// startFrame()
func (this *QTimeLine) Startframe(args ...interface{}) (ret interface{}) {
  // startFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine10startFrameEv
    // invoke: int startFrame()
    var ret0 = C.C_ZNK9QTimeLine10startFrameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "startFrame", args)
  }

  return
}

// updateInterval()
func (this *QTimeLine) Updateinterval(args ...interface{}) (ret interface{}) {
  // updateInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine14updateIntervalEv
    // invoke: int updateInterval()
    var ret0 = C.C_ZNK9QTimeLine14updateIntervalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "updateInterval", args)
  }

  return
}

// currentFrame()
func (this *QTimeLine) Currentframe(args ...interface{}) (ret interface{}) {
  // currentFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine12currentFrameEv
    // invoke: int currentFrame()
    var ret0 = C.C_ZNK9QTimeLine12currentFrameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "currentFrame", args)
  }

  return
}

// metaObject()
func (this *QTimeLine) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QTimeLine10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "metaObject", args)
  }

  return
}

// setUpdateInterval(int)
func (this *QTimeLine) Setupdateinterval(args ...interface{}) () {
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
    // invoke: void setUpdateInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine17setUpdateIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setUpdateInterval", args)
  }

  return
}

// curveShape()
func (this *QTimeLine) Curveshape(args ...interface{}) () {
  // curveShape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine10curveShapeEv
    // invoke: QTimeLine::CurveShape curveShape()
    C.C_ZNK9QTimeLine10curveShapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "curveShape", args)
  }

  return
}

// currentTime()
func (this *QTimeLine) Currenttime(args ...interface{}) (ret interface{}) {
  // currentTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine11currentTimeEv
    // invoke: int currentTime()
    var ret0 = C.C_ZNK9QTimeLine11currentTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "currentTime", args)
  }

  return
}

// easingCurve()
func (this *QTimeLine) Easingcurve(args ...interface{}) (ret interface{}) {
  // easingCurve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeLine11easingCurveEv
    // invoke: QEasingCurve easingCurve()
    var ret0 = C.C_ZNK9QTimeLine11easingCurveEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QEasingCurve{}) // "QEasingCurve"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimeLine", "easingCurve", args)
  }

  return
}

// setEndFrame(int)
func (this *QTimeLine) Setendframe(args ...interface{}) () {
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
    // invoke: void setEndFrame(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine11setEndFrameEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setEndFrame", args)
  }

  return
}

// setDuration(int)
func (this *QTimeLine) Setduration(args ...interface{}) () {
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
    // invoke: void setDuration(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine11setDurationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setDuration", args)
  }

  return
}

// ~QTimeLine()
func (this *QTimeLine) Freeqtimeline(args ...interface{}) () {
  // ~QTimeLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeLineD0Ev
    // invoke: void ~QTimeLine()
    C.C_ZN9QTimeLineD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeLine", "~QTimeLine", args)
  }

  return
}

// setStartFrame(int)
func (this *QTimeLine) Setstartframe(args ...interface{}) () {
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
    // invoke: void setStartFrame(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine13setStartFrameEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setStartFrame", args)
  }

  return
}

// setPaused(_Bool)
func (this *QTimeLine) Setpaused(args ...interface{}) () {
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
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeLine9setPausedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeLine", "setPaused", args)
  }

  return
}

// <= body block end

