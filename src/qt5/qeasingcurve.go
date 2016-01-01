package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qeasingcurve.h
// dst-file: /src/core/qeasingcurve.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QEasingCurve)=8
type QEasingCurve struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQEasingCurve(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "setPeriod", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "addTCBSegment", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "addCubicBezierSegment", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "valueForProgress", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "setAmplitude", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "swap", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QEasingCurve", "setOvershoot", args)
 }

}


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

