package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qspinbox.h
// dst-file: /src/widgets/qspinbox.go
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
// class sizeof(QSpinBox)=1
type QSpinBox struct {
  /*qbase*/ QAbstractSpinBox;
  qclsinst uint64 /* *mut c_void*/;
//  _valueChanged QSpinBox_valueChanged_signal;
}

// class sizeof(QDoubleSpinBox)=1
type QDoubleSpinBox struct {
  /*qbase*/ QAbstractSpinBox;
  qclsinst uint64 /* *mut c_void*/;
//  _valueChanged QDoubleSpinBox_valueChanged_signal;
}


func (this *QSpinBox) setMinimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox10setMinimumEi
  default:
    qtrt.ErrorResolve("QSpinBox", "setMinimum", args)
  }

}


func (this *QSpinBox) cleanText(args ...interface{}) () {
  // cleanText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox9cleanTextEv
  default:
    qtrt.ErrorResolve("QSpinBox", "cleanText", args)
  }

}


func (this *QSpinBox) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox5valueEv
  default:
    qtrt.ErrorResolve("QSpinBox", "value", args)
  }

}


func (this *QSpinBox) FreeQSpinBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSpinBox", "~QSpinBox", args)
  }

}


func (this *QSpinBox) setMaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox10setMaximumEi
  default:
    qtrt.ErrorResolve("QSpinBox", "setMaximum", args)
  }

}


func (this *QSpinBox) setValue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox8setValueEi
  default:
    qtrt.ErrorResolve("QSpinBox", "setValue", args)
  }

}


func (this *QSpinBox) setDisplayIntegerBase(args ...interface{}) () {
  // setDisplayIntegerBase(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox21setDisplayIntegerBaseEi
  default:
    qtrt.ErrorResolve("QSpinBox", "setDisplayIntegerBase", args)
  }

}


func NewQSpinBox(args ...interface{}) QSpinBox {
  return QSpinBox{}
}


func (this *QSpinBox) singleStep(args ...interface{}) () {
  // singleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox10singleStepEv
  default:
    qtrt.ErrorResolve("QSpinBox", "singleStep", args)
  }

}


func (this *QSpinBox) displayIntegerBase(args ...interface{}) () {
  // displayIntegerBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox18displayIntegerBaseEv
  default:
    qtrt.ErrorResolve("QSpinBox", "displayIntegerBase", args)
  }

}


func (this *QSpinBox) setSuffix(args ...interface{}) () {
  // setSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox9setSuffixERK7QString
  default:
    qtrt.ErrorResolve("QSpinBox", "setSuffix", args)
  }

}


func (this *QSpinBox) maximum(args ...interface{}) () {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox7maximumEv
  default:
    qtrt.ErrorResolve("QSpinBox", "maximum", args)
  }

}


func (this *QSpinBox) setPrefix(args ...interface{}) () {
  // setPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox9setPrefixERK7QString
  default:
    qtrt.ErrorResolve("QSpinBox", "setPrefix", args)
  }

}


func (this *QSpinBox) prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox6prefixEv
  default:
    qtrt.ErrorResolve("QSpinBox", "prefix", args)
  }

}


func (this *QSpinBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QSpinBox", "metaObject", args)
  }

}


func (this *QSpinBox) suffix(args ...interface{}) () {
  // suffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox6suffixEv
  default:
    qtrt.ErrorResolve("QSpinBox", "suffix", args)
  }

}


func (this *QSpinBox) minimum(args ...interface{}) () {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSpinBox7minimumEv
  default:
    qtrt.ErrorResolve("QSpinBox", "minimum", args)
  }

}


func (this *QSpinBox) setSingleStep(args ...interface{}) () {
  // setSingleStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox13setSingleStepEi
  default:
    qtrt.ErrorResolve("QSpinBox", "setSingleStep", args)
  }

}


func (this *QSpinBox) setRange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox8setRangeEii
  default:
    qtrt.ErrorResolve("QSpinBox", "setRange", args)
  }

}


func (this *QDoubleSpinBox) textFromValue(args ...interface{}) () {
  // textFromValue(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox13textFromValueEd
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "textFromValue", args)
  }

}


func (this *QDoubleSpinBox) setSingleStep(args ...interface{}) () {
  // setSingleStep(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox13setSingleStepEd
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSingleStep", args)
  }

}


func (this *QDoubleSpinBox) minimum(args ...interface{}) () {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox7minimumEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "minimum", args)
  }

}


func (this *QDoubleSpinBox) valueFromText(args ...interface{}) () {
  // valueFromText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox13valueFromTextERK7QString
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "valueFromText", args)
  }

}


func (this *QDoubleSpinBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "metaObject", args)
  }

}


func (this *QDoubleSpinBox) setValue(args ...interface{}) () {
  // setValue(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox8setValueEd
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setValue", args)
  }

}


func (this *QDoubleSpinBox) setSuffix(args ...interface{}) () {
  // setSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox9setSuffixERK7QString
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSuffix", args)
  }

}


func (this *QDoubleSpinBox) decimals(args ...interface{}) () {
  // decimals()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox8decimalsEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "decimals", args)
  }

}


func (this *QDoubleSpinBox) prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox6prefixEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "prefix", args)
  }

}


func (this *QDoubleSpinBox) singleStep(args ...interface{}) () {
  // singleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox10singleStepEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "singleStep", args)
  }

}


func (this *QDoubleSpinBox) FreeQDoubleSpinBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "~QDoubleSpinBox", args)
  }

}


func (this *QDoubleSpinBox) fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox5fixupER7QString
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "fixup", args)
  }

}


func NewQDoubleSpinBox(args ...interface{}) QDoubleSpinBox {
  return QDoubleSpinBox{}
}


func (this *QDoubleSpinBox) setPrefix(args ...interface{}) () {
  // setPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox9setPrefixERK7QString
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setPrefix", args)
  }

}


func (this *QDoubleSpinBox) cleanText(args ...interface{}) () {
  // cleanText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox9cleanTextEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "cleanText", args)
  }

}


func (this *QDoubleSpinBox) setMinimum(args ...interface{}) () {
  // setMinimum(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox10setMinimumEd
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMinimum", args)
  }

}


func (this *QDoubleSpinBox) setMaximum(args ...interface{}) () {
  // setMaximum(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox10setMaximumEd
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMaximum", args)
  }

}


func (this *QDoubleSpinBox) setDecimals(args ...interface{}) () {
  // setDecimals(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox11setDecimalsEi
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setDecimals", args)
  }

}


func (this *QDoubleSpinBox) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox5valueEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "value", args)
  }

}


func (this *QDoubleSpinBox) setRange(args ...interface{}) () {
  // setRange(double, double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[0][1] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox8setRangeEdd
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setRange", args)
  }

}


func (this *QDoubleSpinBox) maximum(args ...interface{}) () {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox7maximumEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "maximum", args)
  }

}


func (this *QDoubleSpinBox) suffix(args ...interface{}) () {
  // suffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox6suffixEv
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "suffix", args)
  }

}

// <= body block end

