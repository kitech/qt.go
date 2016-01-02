package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qvariantanimation.h
// dst-file: /src/core/qvariantanimation.go
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
  // proto:  void QVariantAnimation::setDuration(int msecs);
extern void _ZN17QVariantAnimation11setDurationEi(void* qthis, int arg0);
  // proto:  void QVariantAnimation::setKeyValueAt(qreal step, const QVariant & value);
extern void _ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant(void* qthis, double arg0, void* arg1);
  // proto:  QVariant QVariantAnimation::endValue();
extern void _ZNK17QVariantAnimation8endValueEv(void* qthis);
  // proto:  QVariant QVariantAnimation::keyValueAt(qreal step);
extern void _ZNK17QVariantAnimation10keyValueAtEd(void* qthis, double arg0);
  // proto:  void QVariantAnimation::~QVariantAnimation();
extern void _ZN17QVariantAnimationD0Ev(void* qthis);
  // proto:  QVariant QVariantAnimation::currentValue();
extern void _ZNK17QVariantAnimation12currentValueEv(void* qthis);
  // proto:  int QVariantAnimation::duration();
extern void _ZNK17QVariantAnimation8durationEv(void* qthis);
  // proto:  KeyValues QVariantAnimation::keyValues();
extern void _ZNK17QVariantAnimation9keyValuesEv(void* qthis);
  // proto:  void QVariantAnimation::setStartValue(const QVariant & value);
extern void _ZN17QVariantAnimation13setStartValueERK8QVariant(void* qthis, void* arg0);
  // proto:  QVariant QVariantAnimation::startValue();
extern void _ZNK17QVariantAnimation10startValueEv(void* qthis);
  // proto:  const QMetaObject * QVariantAnimation::metaObject();
extern void _ZNK17QVariantAnimation10metaObjectEv(void* qthis);
  // proto:  void QVariantAnimation::setEndValue(const QVariant & value);
extern void _ZN17QVariantAnimation11setEndValueERK8QVariant(void* qthis, void* arg0);
  // proto:  QEasingCurve QVariantAnimation::easingCurve();
extern void _ZNK17QVariantAnimation11easingCurveEv(void* qthis);
  // proto:  void QVariantAnimation::QVariantAnimation(const QVariantAnimation & );
extern void* dector_ZN17QVariantAnimationC1ERKS_(void* arg0);
extern void _ZN17QVariantAnimationC1ERKS_(void* qthis, void* arg0);
  // proto:  void QVariantAnimation::QVariantAnimation(QObject * parent);
extern void* dector_ZN17QVariantAnimationC1EP7QObject(void* arg0);
extern void _ZN17QVariantAnimationC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QVariantAnimation::setEasingCurve(const QEasingCurve & easing);
extern void _ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve(void* qthis, void* arg0);
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

// class sizeof(QVariantAnimation)=1
type QVariantAnimation struct {
  /*qbase*/ QAbstractAnimation;
  qclsinst uint64 /* *mut c_void*/;
//  _valueChanged QVariantAnimation_valueChanged_signal;
}

  // proto:  void QVariantAnimation::setDuration(int msecs);
func (this *QVariantAnimation) setDuration(args ...interface{}) () {
  // setDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimation11setDurationEi
    // invoke: void setDuration(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN17QVariantAnimation11setDurationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setDuration", args)
  }

}

  // proto:  void QVariantAnimation::setKeyValueAt(qreal step, const QVariant & value);
func (this *QVariantAnimation) setKeyValueAt(args ...interface{}) () {
  // setKeyValueAt(qreal, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant
    // invoke: void setKeyValueAt(qreal, const class QVariant &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setKeyValueAt", args)
  }

}

  // proto:  QVariant QVariantAnimation::endValue();
func (this *QVariantAnimation) endValue(args ...interface{}) () {
  // endValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation8endValueEv
    // invoke: QVariant endValue()
    C._ZNK17QVariantAnimation8endValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "endValue", args)
  }

}

  // proto:  QVariant QVariantAnimation::keyValueAt(qreal step);
func (this *QVariantAnimation) keyValueAt(args ...interface{}) () {
  // keyValueAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation10keyValueAtEd
    // invoke: QVariant keyValueAt(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZNK17QVariantAnimation10keyValueAtEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "keyValueAt", args)
  }

}

  // proto:  void QVariantAnimation::~QVariantAnimation();
func (this *QVariantAnimation) FreeQVariantAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVariantAnimation", "~QVariantAnimation", args)
  }

}

  // proto:  QVariant QVariantAnimation::currentValue();
func (this *QVariantAnimation) currentValue(args ...interface{}) () {
  // currentValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation12currentValueEv
    // invoke: QVariant currentValue()
    C._ZNK17QVariantAnimation12currentValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "currentValue", args)
  }

}

  // proto:  int QVariantAnimation::duration();
func (this *QVariantAnimation) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation8durationEv
    // invoke: int duration()
    C._ZNK17QVariantAnimation8durationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "duration", args)
  }

}

  // proto:  KeyValues QVariantAnimation::keyValues();
func (this *QVariantAnimation) keyValues(args ...interface{}) () {
  // keyValues()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation9keyValuesEv
    // invoke: KeyValues keyValues()
    C._ZNK17QVariantAnimation9keyValuesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "keyValues", args)
  }

}

  // proto:  void QVariantAnimation::setStartValue(const QVariant & value);
func (this *QVariantAnimation) setStartValue(args ...interface{}) () {
  // setStartValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimation13setStartValueERK8QVariant
    // invoke: void setStartValue(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QVariantAnimation13setStartValueERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setStartValue", args)
  }

}

  // proto:  QVariant QVariantAnimation::startValue();
func (this *QVariantAnimation) startValue(args ...interface{}) () {
  // startValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation10startValueEv
    // invoke: QVariant startValue()
    C._ZNK17QVariantAnimation10startValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "startValue", args)
  }

}

  // proto:  const QMetaObject * QVariantAnimation::metaObject();
func (this *QVariantAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK17QVariantAnimation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "metaObject", args)
  }

}

  // proto:  void QVariantAnimation::setEndValue(const QVariant & value);
func (this *QVariantAnimation) setEndValue(args ...interface{}) () {
  // setEndValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimation11setEndValueERK8QVariant
    // invoke: void setEndValue(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QVariantAnimation11setEndValueERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setEndValue", args)
  }

}

  // proto:  QEasingCurve QVariantAnimation::easingCurve();
func (this *QVariantAnimation) easingCurve(args ...interface{}) () {
  // easingCurve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QVariantAnimation11easingCurveEv
    // invoke: QEasingCurve easingCurve()
    C._ZNK17QVariantAnimation11easingCurveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "easingCurve", args)
  }

}

  // proto:  void QVariantAnimation::QVariantAnimation(const QVariantAnimation & );
func NewQVariantAnimation(args ...interface{}) QVariantAnimation {
  return QVariantAnimation{}
}

  // proto:  void QVariantAnimation::setEasingCurve(const QEasingCurve & easing);
func (this *QVariantAnimation) setEasingCurve(args ...interface{}) () {
  // setEasingCurve(const class QEasingCurve &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEasingCurve{}) // "const QEasingCurve &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve
    // invoke: void setEasingCurve(const class QEasingCurve &)
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setEasingCurve", args)
  }

}

// <= body block end

