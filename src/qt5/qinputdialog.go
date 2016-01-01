package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qinputdialog.h
// dst-file: /src/widgets/qinputdialog.go
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
// class sizeof(QInputDialog)=1
type QInputDialog struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _doubleValueChanged QInputDialog_doubleValueChanged_signal;
//  _intValueChanged QInputDialog_intValueChanged_signal;
//  _textValueChanged QInputDialog_textValueChanged_signal;
//  _intValueSelected QInputDialog_intValueSelected_signal;
//  _textValueSelected QInputDialog_textValueSelected_signal;
//  _doubleValueSelected QInputDialog_doubleValueSelected_signal;
}


func (this *QInputDialog) doubleMaximum(args ...interface{}) () {
  // doubleMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog13doubleMaximumEv
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleMaximum", args)
  }

}


func (this *QInputDialog) setIntMaximum(args ...interface{}) () {
  // setIntMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog13setIntMaximumEi
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntMaximum", args)
  }

}


func (this *QInputDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog10metaObjectEv
  default:
    qtrt.ErrorResolve("QInputDialog", "metaObject", args)
  }

}


func (this *QInputDialog) setIntStep(args ...interface{}) () {
  // setIntStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog10setIntStepEi
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntStep", args)
  }

}


func (this *QInputDialog) intMaximum(args ...interface{}) () {
  // intMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog10intMaximumEv
  default:
    qtrt.ErrorResolve("QInputDialog", "intMaximum", args)
  }

}


func (this *QInputDialog) intStep(args ...interface{}) () {
  // intStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog7intStepEv
  default:
    qtrt.ErrorResolve("QInputDialog", "intStep", args)
  }

}


func (this *QInputDialog) doubleDecimals(args ...interface{}) () {
  // doubleDecimals()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog14doubleDecimalsEv
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleDecimals", args)
  }

}


func (this *QInputDialog) setDoubleDecimals(args ...interface{}) () {
  // setDoubleDecimals(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog17setDoubleDecimalsEi
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleDecimals", args)
  }

}


func (this *QInputDialog) setIntMinimum(args ...interface{}) () {
  // setIntMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog13setIntMinimumEi
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntMinimum", args)
  }

}


func (this *QInputDialog) setTextValue(args ...interface{}) () {
  // setTextValue(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog12setTextValueERK7QString
  default:
    qtrt.ErrorResolve("QInputDialog", "setTextValue", args)
  }

}


func (this *QInputDialog) done(args ...interface{}) () {
  // done(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog4doneEi
  default:
    qtrt.ErrorResolve("QInputDialog", "done", args)
  }

}


func (this *QInputDialog) FreeQInputDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QInputDialog", "~QInputDialog", args)
  }

}


func NewQInputDialog(args ...interface{}) QInputDialog {
  return QInputDialog{}
}


func (this *QInputDialog) setLabelText(args ...interface{}) () {
  // setLabelText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog12setLabelTextERK7QString
  default:
    qtrt.ErrorResolve("QInputDialog", "setLabelText", args)
  }

}


func (this *QInputDialog) labelText(args ...interface{}) () {
  // labelText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog9labelTextEv
  default:
    qtrt.ErrorResolve("QInputDialog", "labelText", args)
  }

}


func (this *QInputDialog) setOkButtonText(args ...interface{}) () {
  // setOkButtonText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog15setOkButtonTextERK7QString
  default:
    qtrt.ErrorResolve("QInputDialog", "setOkButtonText", args)
  }

}


func (this *QInputDialog) comboBoxItems(args ...interface{}) () {
  // comboBoxItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog13comboBoxItemsEv
  default:
    qtrt.ErrorResolve("QInputDialog", "comboBoxItems", args)
  }

}


func (this *QInputDialog) intMinimum(args ...interface{}) () {
  // intMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog10intMinimumEv
  default:
    qtrt.ErrorResolve("QInputDialog", "intMinimum", args)
  }

}


func (this *QInputDialog) setComboBoxEditable(args ...interface{}) () {
  // setComboBoxEditable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog19setComboBoxEditableEb
  default:
    qtrt.ErrorResolve("QInputDialog", "setComboBoxEditable", args)
  }

}


func (this *QInputDialog) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog10setVisibleEb
  default:
    qtrt.ErrorResolve("QInputDialog", "setVisible", args)
  }

}


func (this *QInputDialog) setDoubleMinimum(args ...interface{}) () {
  // setDoubleMinimum(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog16setDoubleMinimumEd
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleMinimum", args)
  }

}


func (this *QInputDialog) doubleMinimum(args ...interface{}) () {
  // doubleMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog13doubleMinimumEv
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleMinimum", args)
  }

}


func (this *QInputDialog) cancelButtonText(args ...interface{}) () {
  // cancelButtonText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog16cancelButtonTextEv
  default:
    qtrt.ErrorResolve("QInputDialog", "cancelButtonText", args)
  }

}


func (this *QInputDialog) setComboBoxItems(args ...interface{}) () {
  // setComboBoxItems(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog16setComboBoxItemsERK11QStringList
  default:
    qtrt.ErrorResolve("QInputDialog", "setComboBoxItems", args)
  }

}


func (this *QInputDialog) isComboBoxEditable(args ...interface{}) () {
  // isComboBoxEditable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog18isComboBoxEditableEv
  default:
    qtrt.ErrorResolve("QInputDialog", "isComboBoxEditable", args)
  }

}


func (this *QInputDialog) open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog4openEP7QObjectPKc
  default:
    qtrt.ErrorResolve("QInputDialog", "open", args)
  }

}


func (this *QInputDialog) okButtonText(args ...interface{}) () {
  // okButtonText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog12okButtonTextEv
  default:
    qtrt.ErrorResolve("QInputDialog", "okButtonText", args)
  }

}


func (this *QInputDialog) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog8sizeHintEv
  default:
    qtrt.ErrorResolve("QInputDialog", "sizeHint", args)
  }

}


func (this *QInputDialog) setIntRange(args ...interface{}) () {
  // setIntRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog11setIntRangeEii
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntRange", args)
  }

}


func (this *QInputDialog) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QInputDialog", "minimumSizeHint", args)
  }

}


func (this *QInputDialog) setDoubleValue(args ...interface{}) () {
  // setDoubleValue(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog14setDoubleValueEd
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleValue", args)
  }

}


func (this *QInputDialog) setIntValue(args ...interface{}) () {
  // setIntValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog11setIntValueEi
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntValue", args)
  }

}


func (this *QInputDialog) doubleValue(args ...interface{}) () {
  // doubleValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog11doubleValueEv
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleValue", args)
  }

}


func (this *QInputDialog) setCancelButtonText(args ...interface{}) () {
  // setCancelButtonText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog19setCancelButtonTextERK7QString
  default:
    qtrt.ErrorResolve("QInputDialog", "setCancelButtonText", args)
  }

}


func (this *QInputDialog) textValue(args ...interface{}) () {
  // textValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog9textValueEv
  default:
    qtrt.ErrorResolve("QInputDialog", "textValue", args)
  }

}


func (this *QInputDialog) setDoubleMaximum(args ...interface{}) () {
  // setDoubleMaximum(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog16setDoubleMaximumEd
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleMaximum", args)
  }

}


func (this *QInputDialog) intValue(args ...interface{}) () {
  // intValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog8intValueEv
  default:
    qtrt.ErrorResolve("QInputDialog", "intValue", args)
  }

}


func (this *QInputDialog) setDoubleRange(args ...interface{}) () {
  // setDoubleRange(double, double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[0][1] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialog14setDoubleRangeEdd
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleRange", args)
  }

}

// <= body block end

