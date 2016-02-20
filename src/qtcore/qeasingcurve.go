package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern double C_ZNK12QEasingCurve9overshootEv(void* qthis); // 4
  // proto:  qreal QEasingCurve::period();
extern double C_ZNK12QEasingCurve6periodEv(void* qthis); // 4
  // proto:  qreal QEasingCurve::valueForProgress(qreal progress);
extern double C_ZNK12QEasingCurve16valueForProgressEd(void* qthis, double arg0); // 4
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
extern double C_ZNK12QEasingCurve9amplitudeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// overshoot()
func (this *QEasingCurve) Overshoot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QEasingCurve9overshootEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEasingCurve", "overshoot", args)
  }

  return
}

// period()
func (this *QEasingCurve) Period(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QEasingCurve6periodEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEasingCurve", "period", args)
  }

  return
}

// valueForProgress(qreal)
func (this *QEasingCurve) Valueforprogress(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QEasingCurve16valueForProgressEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEasingCurve", "valueForProgress", args)
  }

  return
}

// toCubicSpline()
func (this *QEasingCurve) Tocubicspline(args ...interface{}) () {
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
    C.C_ZNK12QEasingCurve13toCubicSplineEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "toCubicSpline", args)
  }

  return
}

// addCubicBezierSegment(const class QPointF &, const class QPointF &, const class QPointF &)
func (this *QEasingCurve) Addcubicbeziersegment(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPointF).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QEasingCurve21addCubicBezierSegmentERK7QPointFS2_S2_(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QEasingCurve", "addCubicBezierSegment", args)
  }

  return
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
    var arg0 = args[0].(*QEasingCurve).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QEasingCurveC2ERKS_(arg0)
    return &QEasingCurve{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QEasingCurve", "QEasingCurve", args)
  }

  return nil // QEasingCurve{Qclsinst:qthis}
}

// type()
func (this *QEasingCurve) Type_(args ...interface{}) () {
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
    C.C_ZNK12QEasingCurve4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "type", args)
  }

  return
}

// swap(class QEasingCurve &)
func (this *QEasingCurve) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QEasingCurve).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "swap", args)
  }

  return
}

// addTCBSegment(const class QPointF &, qreal, qreal, qreal)
func (this *QEasingCurve) Addtcbsegment(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN12QEasingCurve13addTCBSegmentERK7QPointFddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QEasingCurve", "addTCBSegment", args)
  }

  return
}

// setAmplitude(qreal)
func (this *QEasingCurve) Setamplitude(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve12setAmplitudeEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "setAmplitude", args)
  }

  return
}

// ~QEasingCurve()
func (this *QEasingCurve) Freeqeasingcurve(args ...interface{}) () {
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
    C.C_ZN12QEasingCurveD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "~QEasingCurve", args)
  }

  return
}

// setPeriod(qreal)
func (this *QEasingCurve) Setperiod(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve9setPeriodEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "setPeriod", args)
  }

  return
}

// customType()
func (this *QEasingCurve) Customtype(args ...interface{}) () {
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
    C.C_ZNK12QEasingCurve10customTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEasingCurve", "customType", args)
  }

  return
}

// setOvershoot(qreal)
func (this *QEasingCurve) Setovershoot(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QEasingCurve12setOvershootEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEasingCurve", "setOvershoot", args)
  }

  return
}

// amplitude()
func (this *QEasingCurve) Amplitude(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QEasingCurve9amplitudeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEasingCurve", "amplitude", args)
  }

  return
}

// <= body block end

