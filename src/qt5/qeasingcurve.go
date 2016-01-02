package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qeasingcurve.h
// dst-file: /src/core/qeasingcurve.go
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
  // proto:  void QEasingCurve::QEasingCurve(const QEasingCurve & other);
extern void* dector_ZN12QEasingCurveC1ERKS_(void* arg0);
extern void _ZN12QEasingCurveC1ERKS_(void* qthis, void* arg0);
  // proto:  void QEasingCurve::~QEasingCurve();
extern void _ZN12QEasingCurveD0Ev(void* qthis);
  // proto:  EasingFunction QEasingCurve::customType();
extern void _ZNK12QEasingCurve10customTypeEv(void* qthis);
  // proto:  qreal QEasingCurve::overshoot();
extern void _ZNK12QEasingCurve9overshootEv(void* qthis);
  // proto:  void QEasingCurve::setPeriod(qreal period);
extern void _ZN12QEasingCurve9setPeriodEd(void* qthis, double arg0);
  // proto:  void QEasingCurve::addTCBSegment(const QPointF & nextPoint, qreal t, qreal c, qreal b);
extern void _ZN12QEasingCurve13addTCBSegmentERK7QPointFddd(void* qthis, void* arg0, double arg1, double arg2, double arg3);
  // proto:  void QEasingCurve::addCubicBezierSegment(const QPointF & c1, const QPointF & c2, const QPointF & endPoint);
extern void _ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  qreal QEasingCurve::period();
extern void _ZNK12QEasingCurve6periodEv(void* qthis);
  // proto:  qreal QEasingCurve::valueForProgress(qreal progress);
extern void _ZNK12QEasingCurve16valueForProgressEd(void* qthis, double arg0);
  // proto:  void QEasingCurve::setAmplitude(qreal amplitude);
extern void _ZN12QEasingCurve12setAmplitudeEd(void* qthis, double arg0);
  // proto:  void QEasingCurve::swap(QEasingCurve & other);
extern void demth_ZN12QEasingCurve4swapERS_(void* qthis, void* arg0);
  // proto:  void QEasingCurve::setOvershoot(qreal overshoot);
extern void _ZN12QEasingCurve12setOvershootEd(void* qthis, double arg0);
  // proto:  QVector<QPointF> QEasingCurve::toCubicSpline();
extern void _ZNK12QEasingCurve13toCubicSplineEv(void* qthis);
  // proto:  qreal QEasingCurve::amplitude();
extern void _ZNK12QEasingCurve9amplitudeEv(void* qthis);
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

// class sizeof(QEasingCurve)=8
type QEasingCurve struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QEasingCurve::QEasingCurve(const QEasingCurve & other);
func NewQEasingCurve(args ...interface{}) QEasingCurve {
  return QEasingCurve{}
}

  // proto:  void QEasingCurve::~QEasingCurve();
func (this *QEasingCurve) FreeQEasingCurve(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEasingCurve", "~QEasingCurve", args)
  }

}

  // proto:  EasingFunction QEasingCurve::customType();
func (this *QEasingCurve) customType(args ...interface{}) () {
  // customType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve10customTypeEv
  default:
    qtrt.ErrorResolve("QEasingCurve", "customType", args)
  }

}

  // proto:  qreal QEasingCurve::overshoot();
func (this *QEasingCurve) overshoot(args ...interface{}) () {
  // overshoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve9overshootEv
  default:
    qtrt.ErrorResolve("QEasingCurve", "overshoot", args)
  }

}

  // proto:  void QEasingCurve::setPeriod(qreal period);
func (this *QEasingCurve) setPeriod(args ...interface{}) () {
  // setPeriod(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurve9setPeriodEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "setPeriod", args)
  }

}

  // proto:  void QEasingCurve::addTCBSegment(const QPointF & nextPoint, qreal t, qreal c, qreal b);
func (this *QEasingCurve) addTCBSegment(args ...interface{}) () {
  // addTCBSegment(const class QPointF &, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurve13addTCBSegmentERK7QPointFddd
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "addTCBSegment", args)
  }

}

  // proto:  void QEasingCurve::addCubicBezierSegment(const QPointF & c1, const QPointF & c2, const QPointF & endPoint);
func (this *QEasingCurve) addCubicBezierSegment(args ...interface{}) () {
  // addCubicBezierSegment(const class QPointF &, const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "addCubicBezierSegment", args)
  }

}

  // proto:  qreal QEasingCurve::period();
func (this *QEasingCurve) period(args ...interface{}) () {
  // period()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve6periodEv
  default:
    qtrt.ErrorResolve("QEasingCurve", "period", args)
  }

}

  // proto:  qreal QEasingCurve::valueForProgress(qreal progress);
func (this *QEasingCurve) valueForProgress(args ...interface{}) () {
  // valueForProgress(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve16valueForProgressEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "valueForProgress", args)
  }

}

  // proto:  void QEasingCurve::setAmplitude(qreal amplitude);
func (this *QEasingCurve) setAmplitude(args ...interface{}) () {
  // setAmplitude(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurve12setAmplitudeEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "setAmplitude", args)
  }

}

  // proto:  void QEasingCurve::swap(QEasingCurve & other);
func (this *QEasingCurve) swap(args ...interface{}) () {
  // swap(class QEasingCurve &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEasingCurve{}) // "QEasingCurve &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurve4swapERS_
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "swap", args)
  }

}

  // proto:  void QEasingCurve::setOvershoot(qreal overshoot);
func (this *QEasingCurve) setOvershoot(args ...interface{}) () {
  // setOvershoot(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurve12setOvershootEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "setOvershoot", args)
  }

}

  // proto:  QVector<QPointF> QEasingCurve::toCubicSpline();
func (this *QEasingCurve) toCubicSpline(args ...interface{}) () {
  // toCubicSpline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve13toCubicSplineEv
  default:
    qtrt.ErrorResolve("QEasingCurve", "toCubicSpline", args)
  }

}

  // proto:  qreal QEasingCurve::amplitude();
func (this *QEasingCurve) amplitude(args ...interface{}) () {
  // amplitude()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve9amplitudeEv
  default:
    qtrt.ErrorResolve("QEasingCurve", "amplitude", args)
  }

}

// <= body block end

