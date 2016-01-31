package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK12QInputDialog13doubleMinimumEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntMinimum(int min);
extern void C_ZN12QInputDialog13setIntMinimumEi(void* qthis, int32_t arg0); // 4
  // proto:  int QInputDialog::intMinimum();
extern void C_ZNK12QInputDialog10intMinimumEv(void* qthis); // 4
  // proto:  int QInputDialog::intMaximum();
extern void C_ZNK12QInputDialog10intMaximumEv(void* qthis); // 4
  // proto:  double QInputDialog::doubleMaximum();
extern void C_ZNK12QInputDialog13doubleMaximumEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntStep(int step);
extern void C_ZN12QInputDialog10setIntStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QInputDialog::done(int result);
extern void C_ZN12QInputDialog4doneEi(void* qthis, int32_t arg0); // 4
  // proto:  void QInputDialog::setTextValue(const QString & text);
extern void C_ZN12QInputDialog12setTextValueERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QInputDialog::setComboBoxEditable(bool editable);
extern void C_ZN12QInputDialog19setComboBoxEditableEb(void* qthis, bool arg0); // 4
  // proto:  void QInputDialog::open(QObject * receiver, const char * member);
extern void C_ZN12QInputDialog4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  void QInputDialog::setDoubleRange(double min, double max);
extern void C_ZN12QInputDialog14setDoubleRangeEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QInputDialog::setLabelText(const QString & text);
extern void C_ZN12QInputDialog12setLabelTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QInputDialog::labelText();
extern void C_ZNK12QInputDialog9labelTextEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntMaximum(int max);
extern void C_ZN12QInputDialog13setIntMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QInputDialog::setDoubleMaximum(double max);
extern void C_ZN12QInputDialog16setDoubleMaximumEd(void* qthis, double arg0); // 4
  // proto:  double QInputDialog::doubleValue();
extern void C_ZNK12QInputDialog11doubleValueEv(void* qthis); // 4
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
extern void C_ZNK12QInputDialog8intValueEv(void* qthis); // 4
  // proto:  QString QInputDialog::cancelButtonText();
extern void C_ZNK12QInputDialog16cancelButtonTextEv(void* qthis); // 4
  // proto:  void QInputDialog::setOkButtonText(const QString & text);
extern void C_ZN12QInputDialog15setOkButtonTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QSize QInputDialog::sizeHint();
extern void C_ZNK12QInputDialog8sizeHintEv(void* qthis); // 4
  // proto:  QString QInputDialog::textValue();
extern void C_ZNK12QInputDialog9textValueEv(void* qthis); // 4
  // proto:  void QInputDialog::~QInputDialog();
extern void C_ZN12QInputDialogD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QInputDialog::metaObject();
extern void C_ZNK12QInputDialog10metaObjectEv(void* qthis); // 4
  // proto:  QSize QInputDialog::minimumSizeHint();
extern void C_ZNK12QInputDialog15minimumSizeHintEv(void* qthis); // 4
  // proto:  QLineEdit::EchoMode QInputDialog::textEchoMode();
extern void C_ZNK12QInputDialog12textEchoModeEv(void* qthis); // 4
  // proto:  void QInputDialog::setIntValue(int value);
extern void C_ZN12QInputDialog11setIntValueEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QInputDialog::isComboBoxEditable();
extern void C_ZNK12QInputDialog18isComboBoxEditableEv(void* qthis); // 4
  // proto:  QString QInputDialog::okButtonText();
extern void C_ZNK12QInputDialog12okButtonTextEv(void* qthis); // 4
  // proto:  int QInputDialog::doubleDecimals();
extern void C_ZNK12QInputDialog14doubleDecimalsEv(void* qthis); // 4
  // proto:  void QInputDialog::setVisible(bool visible);
extern void C_ZN12QInputDialog10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  int QInputDialog::intStep();
extern void C_ZNK12QInputDialog7intStepEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _doubleValueChanged QInputDialog_doubleValueChanged_signal;
//  _intValueChanged QInputDialog_intValueChanged_signal;
//  _textValueChanged QInputDialog_textValueChanged_signal;
//  _intValueSelected QInputDialog_intValueSelected_signal;
//  _textValueSelected QInputDialog_textValueSelected_signal;
//  _doubleValueSelected QInputDialog_doubleValueSelected_signal;
}

// inputMode()
func (this *QInputDialog) inputMode(args ...interface{}) () {
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
    C.C_ZNK12QInputDialog9inputModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "inputMode", args)
  }

}

// setDoubleDecimals(int)
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
    // invoke: void setDoubleDecimals(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog17setDoubleDecimalsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleDecimals", args)
  }

}

// doubleMinimum()
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
    // invoke: double doubleMinimum()
    C.C_ZNK12QInputDialog13doubleMinimumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleMinimum", args)
  }

}

// setIntMinimum(int)
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
    // invoke: void setIntMinimum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog13setIntMinimumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntMinimum", args)
  }

}

// intMinimum()
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
    // invoke: int intMinimum()
    C.C_ZNK12QInputDialog10intMinimumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "intMinimum", args)
  }

}

// intMaximum()
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
    // invoke: int intMaximum()
    C.C_ZNK12QInputDialog10intMaximumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "intMaximum", args)
  }

}

// doubleMaximum()
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
    // invoke: double doubleMaximum()
    C.C_ZNK12QInputDialog13doubleMaximumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleMaximum", args)
  }

}

// setIntStep(int)
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
    // invoke: void setIntStep(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog10setIntStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntStep", args)
  }

}

// done(int)
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
    // invoke: void done(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog4doneEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "done", args)
  }

}

// setTextValue(const class QString &)
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
    // invoke: void setTextValue(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog12setTextValueERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setTextValue", args)
  }

}

// setComboBoxEditable(_Bool)
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
    // invoke: void setComboBoxEditable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog19setComboBoxEditableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setComboBoxEditable", args)
  }

}

// open(class QObject *, const char *)
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
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN12QInputDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QInputDialog", "open", args)
  }

}

// setDoubleRange(double, double)
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
    // invoke: void setDoubleRange(double, double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN12QInputDialog14setDoubleRangeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleRange", args)
  }

}

// setLabelText(const class QString &)
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
    // invoke: void setLabelText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog12setLabelTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setLabelText", args)
  }

}

// labelText()
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
    // invoke: QString labelText()
    C.C_ZNK12QInputDialog9labelTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "labelText", args)
  }

}

// setIntMaximum(int)
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
    // invoke: void setIntMaximum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog13setIntMaximumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntMaximum", args)
  }

}

// setDoubleMaximum(double)
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
    // invoke: void setDoubleMaximum(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog16setDoubleMaximumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleMaximum", args)
  }

}

// doubleValue()
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
    // invoke: double doubleValue()
    C.C_ZNK12QInputDialog11doubleValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleValue", args)
  }

}

// setCancelButtonText(const class QString &)
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
    // invoke: void setCancelButtonText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog19setCancelButtonTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setCancelButtonText", args)
  }

}

// options()
func (this *QInputDialog) options(args ...interface{}) () {
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
    C.C_ZNK12QInputDialog7optionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "options", args)
  }

}

// setDoubleMinimum(double)
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
    // invoke: void setDoubleMinimum(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog16setDoubleMinimumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleMinimum", args)
  }

}

// setComboBoxItems(const class QStringList &)
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
    // invoke: void setComboBoxItems(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog16setComboBoxItemsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setComboBoxItems", args)
  }

}

// setDoubleValue(double)
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
    // invoke: void setDoubleValue(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog14setDoubleValueEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setDoubleValue", args)
  }

}

// comboBoxItems()
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
    // invoke: QStringList comboBoxItems()
    C.C_ZNK12QInputDialog13comboBoxItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "comboBoxItems", args)
  }

}

// setIntRange(int, int)
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
    // invoke: void setIntRange(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QInputDialog11setIntRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntRange", args)
  }

}

// intValue()
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
    // invoke: int intValue()
    C.C_ZNK12QInputDialog8intValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "intValue", args)
  }

}

// cancelButtonText()
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
    // invoke: QString cancelButtonText()
    C.C_ZNK12QInputDialog16cancelButtonTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "cancelButtonText", args)
  }

}

// setOkButtonText(const class QString &)
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
    // invoke: void setOkButtonText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog15setOkButtonTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setOkButtonText", args)
  }

}

// sizeHint()
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
    // invoke: QSize sizeHint()
    C.C_ZNK12QInputDialog8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "sizeHint", args)
  }

}

// textValue()
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
    // invoke: QString textValue()
    C.C_ZNK12QInputDialog9textValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "textValue", args)
  }

}

// ~QInputDialog()
func (this *QInputDialog) FreeQInputDialog(args ...interface{}) () {
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
    C.C_ZN12QInputDialogD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "~QInputDialog", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QInputDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "metaObject", args)
  }

}

// minimumSizeHint()
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
    // invoke: QSize minimumSizeHint()
    C.C_ZNK12QInputDialog15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "minimumSizeHint", args)
  }

}

// textEchoMode()
func (this *QInputDialog) textEchoMode(args ...interface{}) () {
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
    C.C_ZNK12QInputDialog12textEchoModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "textEchoMode", args)
  }

}

// setIntValue(int)
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
    // invoke: void setIntValue(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog11setIntValueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setIntValue", args)
  }

}

// isComboBoxEditable()
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
    // invoke: bool isComboBoxEditable()
    C.C_ZNK12QInputDialog18isComboBoxEditableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "isComboBoxEditable", args)
  }

}

// okButtonText()
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
    // invoke: QString okButtonText()
    C.C_ZNK12QInputDialog12okButtonTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "okButtonText", args)
  }

}

// doubleDecimals()
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
    // invoke: int doubleDecimals()
    C.C_ZNK12QInputDialog14doubleDecimalsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "doubleDecimals", args)
  }

}

// setVisible(_Bool)
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputDialog10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputDialog", "setVisible", args)
  }

}

// intStep()
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
    // invoke: int intStep()
    C.C_ZNK12QInputDialog7intStepEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputDialog", "intStep", args)
  }

}

// <= body block end

