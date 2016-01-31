package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QEasingCurve::overshoot();
extern void C_ZNK12QEasingCurve9overshootEv(void* qthis); // 4
  // proto:  qreal QEasingCurve::period();
extern void C_ZNK12QEasingCurve6periodEv(void* qthis); // 4
  // proto:  qreal QEasingCurve::valueForProgress(qreal progress);
extern void C_ZNK12QEasingCurve16valueForProgressEd(void* qthis, double arg0); // 4
  // proto:  QVector<QPointF> QEasingCurve::toCubicSpline();
extern void C_ZNK12QEasingCurve13toCubicSplineEv(void* qthis); // 4
  // proto:  void QEasingCurve::addCubicBezierSegment(const QPointF & c1, const QPointF & c2, const QPointF & endPoint);
extern void C_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QEasingCurve::QEasingCurve(const QEasingCurve & other);
extern void* C_ZN12QEasingCurveC2ERKS_(void* arg0); // 3
  // proto:  QEasingCurve::Type QEasingCurve::type();
extern void C_ZNK12QEasingCurve4typeEv(void* qthis); // 4
  // proto:  void QEasingCurve::swap(QEasingCurve & other);
extern void C_ZN12QEasingCurve4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QEasingCurve::addTCBSegment(const QPointF & nextPoint, qreal t, qreal c, qreal b);
extern void C_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd(void* qthis, void* arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QEasingCurve::setAmplitude(qreal amplitude);
extern void C_ZN12QEasingCurve12setAmplitudeEd(void* qthis, double arg0); // 4
  // proto:  void QEasingCurve::~QEasingCurve();
extern void C_ZN12QEasingCurveD2Ev(void* qthis); // 4
  // proto:  void QEasingCurve::setPeriod(qreal period);
extern void C_ZN12QEasingCurve9setPeriodEd(void* qthis, double arg0); // 4
  // proto:  EasingFunction QEasingCurve::customType();
extern void C_ZNK12QEasingCurve10customTypeEv(void* qthis); // 4
  // proto:  void QEasingCurve::setOvershoot(qreal overshoot);
extern void C_ZN12QEasingCurve12setOvershootEd(void* qthis, double arg0); // 4
  // proto:  qreal QEasingCurve::amplitude();
extern void C_ZNK12QEasingCurve9amplitudeEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// overshoot()
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
    // invoke: qreal overshoot()
    var ret = C.C_ZNK12QEasingCurve9overshootEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "overshoot", args)
  }

}

// period()
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
    // invoke: qreal period()
    var ret = C.C_ZNK12QEasingCurve6periodEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "period", args)
  }

}

// valueForProgress(qreal)
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
    // invoke: qreal valueForProgress(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QEasingCurve16valueForProgressEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "valueForProgress", args)
  }

}

// toCubicSpline()
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
    // invoke: QVector<QPointF> toCubicSpline()
    C.C_ZNK12QEasingCurve13toCubicSplineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "toCubicSpline", args)
  }

}

// addCubicBezierSegment(const class QPointF &, const class QPointF &, const class QPointF &)
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
    // invoke: void addCubicBezierSegment(const class QPointF &, const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QEasingCurve", "addCubicBezierSegment", args)
  }

}

// QEasingCurve(const class QEasingCurve &)
func NewQEasingCurve(args ...interface{}) *QEasingCurve {
  // QEasingCurve(const class QEasingCurve &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEasingCurve{}) // "const QEasingCurve &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurveC1ERKS_
    // invoke: void QEasingCurve(const class QEasingCurve &)
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QEasingCurveC2ERKS_(arg0)
    return &QEasingCurve{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QEasingCurve", "QEasingCurve", args)
  }

  return nil // QEasingCurve{qclsinst:qthis}
}

// type()
func (this *QEasingCurve) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QEasingCurve4typeEv
    // invoke: QEasingCurve::Type type()
    C.C_ZNK12QEasingCurve4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "type", args)
  }

}

// swap(class QEasingCurve &)
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
    // invoke: void swap(class QEasingCurve &)
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "swap", args)
  }

}

// addTCBSegment(const class QPointF &, qreal, qreal, qreal)
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
    // invoke: void addTCBSegment(const class QPointF &, qreal, qreal, qreal)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QEasingCurve", "addTCBSegment", args)
  }

}

// setAmplitude(qreal)
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
    // invoke: void setAmplitude(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve12setAmplitudeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "setAmplitude", args)
  }

}

// ~QEasingCurve()
func (this *QEasingCurve) FreeQEasingCurve(args ...interface{}) () {
  // ~QEasingCurve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QEasingCurveD0Ev
    // invoke: void ~QEasingCurve()
    C.C_ZN12QEasingCurveD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "~QEasingCurve", args)
  }

}

// setPeriod(qreal)
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
    // invoke: void setPeriod(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve9setPeriodEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "setPeriod", args)
  }

}

// customType()
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
    // invoke: EasingFunction customType()
    C.C_ZNK12QEasingCurve10customTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "customType", args)
  }

}

// setOvershoot(qreal)
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
    // invoke: void setOvershoot(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve12setOvershootEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "setOvershoot", args)
  }

}

// amplitude()
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
    // invoke: qreal amplitude()
    var ret = C.C_ZNK12QEasingCurve9amplitudeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QEasingCurve", "amplitude", args)
  }

}

// <= body block end

