package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QVariant QVariantAnimation::keyValueAt(qreal step);
extern void* C_ZNK17QVariantAnimation10keyValueAtEd(void* qthis, double arg0); // 4
  // proto:  QVariant QVariantAnimation::currentValue();
extern void* C_ZNK17QVariantAnimation12currentValueEv(void* qthis); // 4
  // proto:  void QVariantAnimation::setEasingCurve(const QEasingCurve & easing);
extern void C_ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve(void* qthis, void* arg0); // 4
  // proto:  void QVariantAnimation::QVariantAnimation(QObject * parent);
extern void* C_ZN17QVariantAnimationC2EP7QObject(void* arg0); // 3
  // proto:  int QVariantAnimation::duration();
extern int32_t C_ZNK17QVariantAnimation8durationEv(void* qthis); // 4
  // proto:  void QVariantAnimation::setStartValue(const QVariant & value);
extern void C_ZN17QVariantAnimation13setStartValueERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  KeyValues QVariantAnimation::keyValues();
extern void C_ZNK17QVariantAnimation9keyValuesEv(void* qthis); // 4
  // proto:  void QVariantAnimation::~QVariantAnimation();
extern void C_ZN17QVariantAnimationD2Ev(void* qthis); // 4
  // proto:  QVariant QVariantAnimation::endValue();
extern void* C_ZNK17QVariantAnimation8endValueEv(void* qthis); // 4
  // proto:  const QMetaObject * QVariantAnimation::metaObject();
extern void C_ZNK17QVariantAnimation10metaObjectEv(void* qthis); // 4
  // proto:  void QVariantAnimation::setKeyValueAt(qreal step, const QVariant & value);
extern void C_ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant(void* qthis, double arg0, void* arg1); // 4
  // proto:  QEasingCurve QVariantAnimation::easingCurve();
extern void* C_ZNK17QVariantAnimation11easingCurveEv(void* qthis); // 4
  // proto:  void QVariantAnimation::setEndValue(const QVariant & value);
extern void C_ZN17QVariantAnimation11setEndValueERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  void QVariantAnimation::setDuration(int msecs);
extern void C_ZN17QVariantAnimation11setDurationEi(void* qthis, int32_t arg0); // 4
  // proto:  QVariant QVariantAnimation::startValue();
extern void* C_ZNK17QVariantAnimation10startValueEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QVariantAnimation_valueChanged_signal;
}

// keyValueAt(qreal)
func (this *QVariantAnimation) Keyvalueat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QVariantAnimation10keyValueAtEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QVariantAnimation", "keyValueAt", args)
  }

  return
}

// currentValue()
func (this *QVariantAnimation) Currentvalue(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QVariantAnimation12currentValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QVariantAnimation", "currentValue", args)
  }

  return
}

// setEasingCurve(const class QEasingCurve &)
func (this *QVariantAnimation) Seteasingcurve(args ...interface{}) () {
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
    C.C_ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setEasingCurve", args)
  }

  return
}

// QVariantAnimation(class QObject *)
func NewQVariantAnimation(args ...interface{}) *QVariantAnimation {
  // QVariantAnimation(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimationC1EP7QObject
    // invoke: void QVariantAnimation(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QVariantAnimationC2EP7QObject(arg0)
    return &QVariantAnimation{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVariantAnimation", "QVariantAnimation", args)
  }

  return nil // QVariantAnimation{qclsinst:qthis}
}

// duration()
func (this *QVariantAnimation) Duration(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QVariantAnimation8durationEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QVariantAnimation", "duration", args)
  }

  return
}

// setStartValue(const class QVariant &)
func (this *QVariantAnimation) Setstartvalue(args ...interface{}) () {
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
    C.C_ZN17QVariantAnimation13setStartValueERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setStartValue", args)
  }

  return
}

// keyValues()
func (this *QVariantAnimation) Keyvalues(args ...interface{}) () {
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
    C.C_ZNK17QVariantAnimation9keyValuesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "keyValues", args)
  }

  return
}

// ~QVariantAnimation()
func (this *QVariantAnimation) Freeqvariantanimation(args ...interface{}) () {
  // ~QVariantAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QVariantAnimationD0Ev
    // invoke: void ~QVariantAnimation()
    C.C_ZN17QVariantAnimationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "~QVariantAnimation", args)
  }

  return
}

// endValue()
func (this *QVariantAnimation) Endvalue(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QVariantAnimation8endValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QVariantAnimation", "endValue", args)
  }

  return
}

// metaObject()
func (this *QVariantAnimation) Metaobject(args ...interface{}) () {
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
    C.C_ZNK17QVariantAnimation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "metaObject", args)
  }

  return
}

// setKeyValueAt(qreal, const class QVariant &)
func (this *QVariantAnimation) Setkeyvalueat(args ...interface{}) () {
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
    C.C_ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setKeyValueAt", args)
  }

  return
}

// easingCurve()
func (this *QVariantAnimation) Easingcurve(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QVariantAnimation11easingCurveEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QEasingCurve{}) // "QEasingCurve"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QVariantAnimation", "easingCurve", args)
  }

  return
}

// setEndValue(const class QVariant &)
func (this *QVariantAnimation) Setendvalue(args ...interface{}) () {
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
    C.C_ZN17QVariantAnimation11setEndValueERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setEndValue", args)
  }

  return
}

// setDuration(int)
func (this *QVariantAnimation) Setduration(args ...interface{}) () {
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
    C.C_ZN17QVariantAnimation11setDurationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setDuration", args)
  }

  return
}

// startValue()
func (this *QVariantAnimation) Startvalue(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QVariantAnimation10startValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QVariantAnimation", "startValue", args)
  }

  return
}

// <= body block end

