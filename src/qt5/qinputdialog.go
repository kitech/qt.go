package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qinputdialog.h
// dst-file: /src/widgets/qinputdialog.go
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
  // proto:  QInputDialog::InputMode QInputDialog::inputMode();
extern void C_ZNK12QInputDialog9inputModeEv(void* qthis); // 4
  // proto:  void QInputDialog::setDoubleDecimals(int decimals);
extern void C_ZN12QInputDialog17setDoubleDecimalsEi(void* qthis, int32_t arg0); // 4
  // proto:  double QInputDialog::doubleMinimum();
extern double C_ZNK12QInputDialog13doubleMinimumEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntMinimum(int min);
extern void C_ZN12QInputDialog13setIntMinimumEi(void* qthis, int32_t arg0); // 4
  // proto:  int QInputDialog::intMinimum();
extern int32_t C_ZNK12QInputDialog10intMinimumEv(void* qthis); // 4
  // proto:  int QInputDialog::intMaximum();
extern int32_t C_ZNK12QInputDialog10intMaximumEv(void* qthis); // 4
  // proto:  double QInputDialog::doubleMaximum();
extern double C_ZNK12QInputDialog13doubleMaximumEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntStep(int step);
extern void C_ZN12QInputDialog10setIntStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QInputDialog::done(int result);
extern void C_ZN12QInputDialog4doneEi(void* qthis, int32_t arg0); // 4
  // proto:  void QInputDialog::setTextValue(const QString & text);
extern void C_ZN12QInputDialog12setTextValueERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QInputDialog::setComboBoxEditable(bool editable);
extern void C_ZN12QInputDialog19setComboBoxEditableEb(void* qthis, bool arg0); // 4
  // proto:  void QInputDialog::open(QObject * receiver, const char * member);
extern void C_ZN12QInputDialog4openEP7QObjectPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QInputDialog::setDoubleRange(double min, double max);
extern void C_ZN12QInputDialog14setDoubleRangeEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QInputDialog::setLabelText(const QString & text);
extern void C_ZN12QInputDialog12setLabelTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QInputDialog::labelText();
extern void* C_ZNK12QInputDialog9labelTextEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntMaximum(int max);
extern void C_ZN12QInputDialog13setIntMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QInputDialog::setDoubleMaximum(double max);
extern void C_ZN12QInputDialog16setDoubleMaximumEd(void* qthis, double arg0); // 4
  // proto:  double QInputDialog::doubleValue();
extern double C_ZNK12QInputDialog11doubleValueEv(void* qthis); // 4
  // proto:  void QInputDialog::setCancelButtonText(const QString & text);
extern void C_ZN12QInputDialog19setCancelButtonTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  InputDialogOptions QInputDialog::options();
extern void C_ZNK12QInputDialog7optionsEv(void* qthis); // 4
  // proto:  void QInputDialog::setDoubleMinimum(double min);
extern void C_ZN12QInputDialog16setDoubleMinimumEd(void* qthis, double arg0); // 4
  // proto:  void QInputDialog::setComboBoxItems(const QStringList & items);
extern void C_ZN12QInputDialog16setComboBoxItemsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QInputDialog::setDoubleValue(double value);
extern void C_ZN12QInputDialog14setDoubleValueEd(void* qthis, double arg0); // 4
  // proto:  QStringList QInputDialog::comboBoxItems();
extern void C_ZNK12QInputDialog13comboBoxItemsEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntRange(int min, int max);
extern void C_ZN12QInputDialog11setIntRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QInputDialog::intValue();
extern int32_t C_ZNK12QInputDialog8intValueEv(void* qthis); // 4
  // proto:  QString QInputDialog::cancelButtonText();
extern void* C_ZNK12QInputDialog16cancelButtonTextEv(void* qthis); // 4
  // proto:  void QInputDialog::setOkButtonText(const QString & text);
extern void C_ZN12QInputDialog15setOkButtonTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QSize QInputDialog::sizeHint();
extern void* C_ZNK12QInputDialog8sizeHintEv(void* qthis); // 4
  // proto:  QString QInputDialog::textValue();
extern void* C_ZNK12QInputDialog9textValueEv(void* qthis); // 4
  // proto:  void QInputDialog::~QInputDialog();
extern void C_ZN12QInputDialogD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QInputDialog::metaObject();
extern void C_ZNK12QInputDialog10metaObjectEv(void* qthis); // 4
  // proto:  QSize QInputDialog::minimumSizeHint();
extern void* C_ZNK12QInputDialog15minimumSizeHintEv(void* qthis); // 4
  // proto:  QLineEdit::EchoMode QInputDialog::textEchoMode();
extern void C_ZNK12QInputDialog12textEchoModeEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntValue(int value);
extern void C_ZN12QInputDialog11setIntValueEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QInputDialog::isComboBoxEditable();
extern bool C_ZNK12QInputDialog18isComboBoxEditableEv(void* qthis); // 4
  // proto:  QString QInputDialog::okButtonText();
extern void* C_ZNK12QInputDialog12okButtonTextEv(void* qthis); // 4
  // proto:  int QInputDialog::doubleDecimals();
extern int32_t C_ZNK12QInputDialog14doubleDecimalsEv(void* qthis); // 4
  // proto:  void QInputDialog::setVisible(bool visible);
extern void C_ZN12QInputDialog10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  int QInputDialog::intStep();
extern int32_t C_ZNK12QInputDialog7intStepEv(void* qthis); // 4
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

// class sizeof(QInputDialog)=1
type QInputDialog struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _doubleValueChanged QInputDialog_doubleValueChanged_signal;
//  _intValueChanged QInputDialog_intValueChanged_signal;
//  _textValueChanged QInputDialog_textValueChanged_signal;
//  _intValueSelected QInputDialog_intValueSelected_signal;
//  _textValueSelected QInputDialog_textValueSelected_signal;
//  _doubleValueSelected QInputDialog_doubleValueSelected_signal;
}

// inputMode()
func (this *QInputDialog) Inputmode(args ...interface{}) () {
  // inputMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog9inputModeEv
    // invoke: QInputDialog::InputMode inputMode()
    C.C_ZNK12QInputDialog9inputModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "inputMode", args)
  }

  return
}

// setDoubleDecimals(int)
func (this *QInputDialog) Setdoubledecimals(args ...interface{}) () {
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
    // invoke: void setDoubleDecimals(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog17setDoubleDecimalsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleDecimals", args)
  }

  return
}

// doubleMinimum()
func (this *QInputDialog) Doubleminimum(args ...interface{}) (ret interface{}) {
  // doubleMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog13doubleMinimumEv
    // invoke: double doubleMinimum()
    var ret0 = C.C_ZNK12QInputDialog13doubleMinimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleMinimum", args)
  }

  return
}

// setIntMinimum(int)
func (this *QInputDialog) Setintminimum(args ...interface{}) () {
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
    // invoke: void setIntMinimum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog13setIntMinimumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntMinimum", args)
  }

  return
}

// intMinimum()
func (this *QInputDialog) Intminimum(args ...interface{}) (ret interface{}) {
  // intMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog10intMinimumEv
    // invoke: int intMinimum()
    var ret0 = C.C_ZNK12QInputDialog10intMinimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "intMinimum", args)
  }

  return
}

// intMaximum()
func (this *QInputDialog) Intmaximum(args ...interface{}) (ret interface{}) {
  // intMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog10intMaximumEv
    // invoke: int intMaximum()
    var ret0 = C.C_ZNK12QInputDialog10intMaximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "intMaximum", args)
  }

  return
}

// doubleMaximum()
func (this *QInputDialog) Doublemaximum(args ...interface{}) (ret interface{}) {
  // doubleMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog13doubleMaximumEv
    // invoke: double doubleMaximum()
    var ret0 = C.C_ZNK12QInputDialog13doubleMaximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleMaximum", args)
  }

  return
}

// setIntStep(int)
func (this *QInputDialog) Setintstep(args ...interface{}) () {
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
    // invoke: void setIntStep(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog10setIntStepEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntStep", args)
  }

  return
}

// done(int)
func (this *QInputDialog) Done(args ...interface{}) () {
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
    // invoke: void done(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog4doneEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "done", args)
  }

  return
}

// setTextValue(const class QString &)
func (this *QInputDialog) Settextvalue(args ...interface{}) () {
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
    // invoke: void setTextValue(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog12setTextValueERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setTextValue", args)
  }

  return
}

// setComboBoxEditable(_Bool)
func (this *QInputDialog) Setcomboboxeditable(args ...interface{}) () {
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
    // invoke: void setComboBoxEditable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog19setComboBoxEditableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setComboBoxEditable", args)
  }

  return
}

// open(class QObject *, const char *)
func (this *QInputDialog) Open(args ...interface{}) () {
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
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN12QInputDialog4openEP7QObjectPKc(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QInputDialog", "open", args)
  }

  return
}

// setDoubleRange(double, double)
func (this *QInputDialog) Setdoublerange(args ...interface{}) () {
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
    // invoke: void setDoubleRange(double, double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN12QInputDialog14setDoubleRangeEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleRange", args)
  }

  return
}

// setLabelText(const class QString &)
func (this *QInputDialog) Setlabeltext(args ...interface{}) () {
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
    // invoke: void setLabelText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog12setLabelTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setLabelText", args)
  }

  return
}

// labelText()
func (this *QInputDialog) Labeltext(args ...interface{}) (ret interface{}) {
  // labelText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog9labelTextEv
    // invoke: QString labelText()
    var ret0 = C.C_ZNK12QInputDialog9labelTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "labelText", args)
  }

  return
}

// setIntMaximum(int)
func (this *QInputDialog) Setintmaximum(args ...interface{}) () {
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
    // invoke: void setIntMaximum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog13setIntMaximumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntMaximum", args)
  }

  return
}

// setDoubleMaximum(double)
func (this *QInputDialog) Setdoublemaximum(args ...interface{}) () {
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
    // invoke: void setDoubleMaximum(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog16setDoubleMaximumEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleMaximum", args)
  }

  return
}

// doubleValue()
func (this *QInputDialog) Doublevalue(args ...interface{}) (ret interface{}) {
  // doubleValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog11doubleValueEv
    // invoke: double doubleValue()
    var ret0 = C.C_ZNK12QInputDialog11doubleValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleValue", args)
  }

  return
}

// setCancelButtonText(const class QString &)
func (this *QInputDialog) Setcancelbuttontext(args ...interface{}) () {
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
    // invoke: void setCancelButtonText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog19setCancelButtonTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setCancelButtonText", args)
  }

  return
}

// options()
func (this *QInputDialog) Options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog7optionsEv
    // invoke: InputDialogOptions options()
    C.C_ZNK12QInputDialog7optionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "options", args)
  }

  return
}

// setDoubleMinimum(double)
func (this *QInputDialog) Setdoubleminimum(args ...interface{}) () {
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
    // invoke: void setDoubleMinimum(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog16setDoubleMinimumEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleMinimum", args)
  }

  return
}

// setComboBoxItems(const class QStringList &)
func (this *QInputDialog) Setcomboboxitems(args ...interface{}) () {
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
    // invoke: void setComboBoxItems(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog16setComboBoxItemsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setComboBoxItems", args)
  }

  return
}

// setDoubleValue(double)
func (this *QInputDialog) Setdoublevalue(args ...interface{}) () {
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
    // invoke: void setDoubleValue(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog14setDoubleValueEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleValue", args)
  }

  return
}

// comboBoxItems()
func (this *QInputDialog) Comboboxitems(args ...interface{}) () {
  // comboBoxItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog13comboBoxItemsEv
    // invoke: QStringList comboBoxItems()
    C.C_ZNK12QInputDialog13comboBoxItemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "comboBoxItems", args)
  }

  return
}

// setIntRange(int, int)
func (this *QInputDialog) Setintrange(args ...interface{}) () {
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
    // invoke: void setIntRange(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QInputDialog11setIntRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntRange", args)
  }

  return
}

// intValue()
func (this *QInputDialog) Intvalue(args ...interface{}) (ret interface{}) {
  // intValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog8intValueEv
    // invoke: int intValue()
    var ret0 = C.C_ZNK12QInputDialog8intValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "intValue", args)
  }

  return
}

// cancelButtonText()
func (this *QInputDialog) Cancelbuttontext(args ...interface{}) (ret interface{}) {
  // cancelButtonText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog16cancelButtonTextEv
    // invoke: QString cancelButtonText()
    var ret0 = C.C_ZNK12QInputDialog16cancelButtonTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "cancelButtonText", args)
  }

  return
}

// setOkButtonText(const class QString &)
func (this *QInputDialog) Setokbuttontext(args ...interface{}) () {
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
    // invoke: void setOkButtonText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog15setOkButtonTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setOkButtonText", args)
  }

  return
}

// sizeHint()
func (this *QInputDialog) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK12QInputDialog8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "sizeHint", args)
  }

  return
}

// textValue()
func (this *QInputDialog) Textvalue(args ...interface{}) (ret interface{}) {
  // textValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog9textValueEv
    // invoke: QString textValue()
    var ret0 = C.C_ZNK12QInputDialog9textValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "textValue", args)
  }

  return
}

// ~QInputDialog()
func (this *QInputDialog) Freeqinputdialog(args ...interface{}) () {
  // ~QInputDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputDialogD0Ev
    // invoke: void ~QInputDialog()
    C.C_ZN12QInputDialogD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "~QInputDialog", args)
  }

  return
}

// metaObject()
func (this *QInputDialog) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QInputDialog10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QInputDialog) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK12QInputDialog15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "minimumSizeHint", args)
  }

  return
}

// textEchoMode()
func (this *QInputDialog) Textechomode(args ...interface{}) () {
  // textEchoMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog12textEchoModeEv
    // invoke: QLineEdit::EchoMode textEchoMode()
    C.C_ZNK12QInputDialog12textEchoModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "textEchoMode", args)
  }

  return
}

// setIntValue(int)
func (this *QInputDialog) Setintvalue(args ...interface{}) () {
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
    // invoke: void setIntValue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog11setIntValueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntValue", args)
  }

  return
}

// isComboBoxEditable()
func (this *QInputDialog) Iscomboboxeditable(args ...interface{}) (ret interface{}) {
  // isComboBoxEditable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog18isComboBoxEditableEv
    // invoke: bool isComboBoxEditable()
    var ret0 = C.C_ZNK12QInputDialog18isComboBoxEditableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "isComboBoxEditable", args)
  }

  return
}

// okButtonText()
func (this *QInputDialog) Okbuttontext(args ...interface{}) (ret interface{}) {
  // okButtonText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog12okButtonTextEv
    // invoke: QString okButtonText()
    var ret0 = C.C_ZNK12QInputDialog12okButtonTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "okButtonText", args)
  }

  return
}

// doubleDecimals()
func (this *QInputDialog) Doubledecimals(args ...interface{}) (ret interface{}) {
  // doubleDecimals()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog14doubleDecimalsEv
    // invoke: int doubleDecimals()
    var ret0 = C.C_ZNK12QInputDialog14doubleDecimalsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleDecimals", args)
  }

  return
}

// setVisible(_Bool)
func (this *QInputDialog) Setvisible(args ...interface{}) () {
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setVisible", args)
  }

  return
}

// intStep()
func (this *QInputDialog) Intstep(args ...interface{}) (ret interface{}) {
  // intStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputDialog7intStepEv
    // invoke: int intStep()
    var ret0 = C.C_ZNK12QInputDialog7intStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QInputDialog", "intStep", args)
  }

  return
}

// <= body block end

