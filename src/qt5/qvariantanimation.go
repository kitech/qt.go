package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qvariantanimation.h
// dst-file: /src/core/qvariantanimation.go
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
// class sizeof(QVariantAnimation)=1
type QVariantAnimation struct {
  /*qbase*/ QAbstractAnimation;
  qclsinst uint64 /* *mut c_void*/;
//  _valueChanged QVariantAnimation_valueChanged_signal;
}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setDuration", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setKeyValueAt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "endValue", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "keyValueAt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "currentValue", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "duration", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "keyValues", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setStartValue", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "startValue", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setEndValue", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "easingCurve", args)
  }

}


func NewQVariantAnimation(args ...interface{}) QVariantAnimation {
  return QVariantAnimation{}
}


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
  default:
    qtrt.ErrorResolve("QVariantAnimation", "setEasingCurve", args)
  }

}

// <= body block end

