package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qspinbox.h
// dst-file: /src/widgets/qspinbox.go
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
  // proto:  QString QSpinBox::suffix();
extern void C_ZNK8QSpinBox6suffixEv(void* qthis); // 4
  // proto:  void QSpinBox::QSpinBox(QWidget * parent);
extern void* C_ZN8QSpinBoxC2EP7QWidget(void* arg0); // 3
  // proto:  int QSpinBox::singleStep();
extern void C_ZNK8QSpinBox10singleStepEv(void* qthis); // 4
  // proto:  QString QSpinBox::prefix();
extern void C_ZNK8QSpinBox6prefixEv(void* qthis); // 4
  // proto:  int QSpinBox::minimum();
extern void C_ZNK8QSpinBox7minimumEv(void* qthis); // 4
  // proto:  QString QSpinBox::cleanText();
extern void C_ZNK8QSpinBox9cleanTextEv(void* qthis); // 4
  // proto:  void QSpinBox::setSingleStep(int val);
extern void C_ZN8QSpinBox13setSingleStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSpinBox::setDisplayIntegerBase(int base);
extern void C_ZN8QSpinBox21setDisplayIntegerBaseEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSpinBox::setMaximum(int max);
extern void C_ZN8QSpinBox10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSpinBox::setValue(int val);
extern void C_ZN8QSpinBox8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSpinBox::displayIntegerBase();
extern void C_ZNK8QSpinBox18displayIntegerBaseEv(void* qthis); // 4
  // proto:  void QSpinBox::setRange(int min, int max);
extern void C_ZN8QSpinBox8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QSpinBox::metaObject();
extern void C_ZNK8QSpinBox10metaObjectEv(void* qthis); // 4
  // proto:  void QSpinBox::~QSpinBox();
extern void C_ZN8QSpinBoxD2Ev(void* qthis); // 4
  // proto:  void QSpinBox::setSuffix(const QString & suffix);
extern void C_ZN8QSpinBox9setSuffixERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QSpinBox::maximum();
extern void C_ZNK8QSpinBox7maximumEv(void* qthis); // 4
  // proto:  int QSpinBox::value();
extern void C_ZNK8QSpinBox5valueEv(void* qthis); // 4
  // proto:  void QSpinBox::setPrefix(const QString & prefix);
extern void C_ZN8QSpinBox9setPrefixERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSpinBox::setMinimum(int min);
extern void C_ZN8QSpinBox10setMinimumEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QDoubleSpinBox::suffix();
extern void C_ZNK14QDoubleSpinBox6suffixEv(void* qthis); // 4
  // proto:  double QDoubleSpinBox::singleStep();
extern void C_ZNK14QDoubleSpinBox10singleStepEv(void* qthis); // 4
  // proto:  QString QDoubleSpinBox::prefix();
extern void C_ZNK14QDoubleSpinBox6prefixEv(void* qthis); // 4
  // proto:  double QDoubleSpinBox::minimum();
extern void C_ZNK14QDoubleSpinBox7minimumEv(void* qthis); // 4
  // proto:  QString QDoubleSpinBox::textFromValue(double val);
extern void C_ZNK14QDoubleSpinBox13textFromValueEd(void* qthis, double arg0); // 4
  // proto:  void QDoubleSpinBox::fixup(QString & str);
extern void C_ZNK14QDoubleSpinBox5fixupER7QString(void* qthis, void* arg0); // 4
  // proto:  void QDoubleSpinBox::QDoubleSpinBox(QWidget * parent);
extern void* C_ZN14QDoubleSpinBoxC2EP7QWidget(void* arg0); // 3
  // proto:  void QDoubleSpinBox::setSingleStep(double val);
extern void C_ZN14QDoubleSpinBox13setSingleStepEd(void* qthis, double arg0); // 4
  // proto:  void QDoubleSpinBox::setRange(double min, double max);
extern void C_ZN14QDoubleSpinBox8setRangeEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QDoubleSpinBox::setMaximum(double max);
extern void C_ZN14QDoubleSpinBox10setMaximumEd(void* qthis, double arg0); // 4
  // proto:  void QDoubleSpinBox::setValue(double val);
extern void C_ZN14QDoubleSpinBox8setValueEd(void* qthis, double arg0); // 4
  // proto:  double QDoubleSpinBox::valueFromText(const QString & text);
extern void C_ZNK14QDoubleSpinBox13valueFromTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QValidator::State QDoubleSpinBox::validate(QString & input, int & pos);
extern void C_ZNK14QDoubleSpinBox8validateER7QStringRi(void* qthis, void* arg0, int32_t* arg1); // 4
  // proto:  QString QDoubleSpinBox::cleanText();
extern void C_ZNK14QDoubleSpinBox9cleanTextEv(void* qthis); // 4
  // proto:  void QDoubleSpinBox::setDecimals(int prec);
extern void C_ZN14QDoubleSpinBox11setDecimalsEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QDoubleSpinBox::metaObject();
extern void C_ZNK14QDoubleSpinBox10metaObjectEv(void* qthis); // 4
  // proto:  void QDoubleSpinBox::~QDoubleSpinBox();
extern void C_ZN14QDoubleSpinBoxD2Ev(void* qthis); // 4
  // proto:  void QDoubleSpinBox::setSuffix(const QString & suffix);
extern void C_ZN14QDoubleSpinBox9setSuffixERK7QString(void* qthis, void* arg0); // 4
  // proto:  double QDoubleSpinBox::maximum();
extern void C_ZNK14QDoubleSpinBox7maximumEv(void* qthis); // 4
  // proto:  double QDoubleSpinBox::value();
extern void C_ZNK14QDoubleSpinBox5valueEv(void* qthis); // 4
  // proto:  int QDoubleSpinBox::decimals();
extern void C_ZNK14QDoubleSpinBox8decimalsEv(void* qthis); // 4
  // proto:  void QDoubleSpinBox::setPrefix(const QString & prefix);
extern void C_ZN14QDoubleSpinBox9setPrefixERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QDoubleSpinBox::setMinimum(double min);
extern void C_ZN14QDoubleSpinBox10setMinimumEd(void* qthis, double arg0); // 4
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

// class sizeof(QSpinBox)=1
type QSpinBox struct {
  /*qbase*/ QAbstractSpinBox;
  qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QSpinBox_valueChanged_signal;
}

// class sizeof(QDoubleSpinBox)=1
type QDoubleSpinBox struct {
  /*qbase*/ QAbstractSpinBox;
  qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QDoubleSpinBox_valueChanged_signal;
}

// suffix()
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
    // invoke: QString suffix()
    var ret = C.C_ZNK8QSpinBox6suffixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "suffix", args)
  }

}

// QSpinBox(class QWidget *)
func NewQSpinBox(args ...interface{}) *QSpinBox {
  // QSpinBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBoxC1EP7QWidget
    // invoke: void QSpinBox(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QSpinBoxC2EP7QWidget(arg0)
    return &QSpinBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSpinBox", "QSpinBox", args)
  }

  return nil // QSpinBox{qclsinst:qthis}
}

// singleStep()
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
    // invoke: int singleStep()
    var ret = C.C_ZNK8QSpinBox10singleStepEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "singleStep", args)
  }

}

// prefix()
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
    // invoke: QString prefix()
    var ret = C.C_ZNK8QSpinBox6prefixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "prefix", args)
  }

}

// minimum()
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
    // invoke: int minimum()
    var ret = C.C_ZNK8QSpinBox7minimumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "minimum", args)
  }

}

// cleanText()
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
    // invoke: QString cleanText()
    var ret = C.C_ZNK8QSpinBox9cleanTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "cleanText", args)
  }

}

// setSingleStep(int)
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
    // invoke: void setSingleStep(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox13setSingleStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setSingleStep", args)
  }

}

// setDisplayIntegerBase(int)
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
    // invoke: void setDisplayIntegerBase(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox21setDisplayIntegerBaseEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setDisplayIntegerBase", args)
  }

}

// setMaximum(int)
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
    // invoke: void setMaximum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox10setMaximumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setMaximum", args)
  }

}

// setValue(int)
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
    // invoke: void setValue(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox8setValueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setValue", args)
  }

}

// displayIntegerBase()
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
    // invoke: int displayIntegerBase()
    var ret = C.C_ZNK8QSpinBox18displayIntegerBaseEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "displayIntegerBase", args)
  }

}

// setRange(int, int)
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
    // invoke: void setRange(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QSpinBox8setRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSpinBox", "setRange", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK8QSpinBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "metaObject", args)
  }

}

// ~QSpinBox()
func (this *QSpinBox) FreeQSpinBox(args ...interface{}) () {
  // ~QSpinBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBoxD0Ev
    // invoke: void ~QSpinBox()
    C.C_ZN8QSpinBoxD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "~QSpinBox", args)
  }

}

// setSuffix(const class QString &)
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
    // invoke: void setSuffix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox9setSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setSuffix", args)
  }

}

// maximum()
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
    // invoke: int maximum()
    var ret = C.C_ZNK8QSpinBox7maximumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "maximum", args)
  }

}

// value()
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
    // invoke: int value()
    var ret = C.C_ZNK8QSpinBox5valueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSpinBox", "value", args)
  }

}

// setPrefix(const class QString &)
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
    // invoke: void setPrefix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox9setPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setPrefix", args)
  }

}

// setMinimum(int)
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
    // invoke: void setMinimum(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox10setMinimumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setMinimum", args)
  }

}

// suffix()
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
    // invoke: QString suffix()
    var ret = C.C_ZNK14QDoubleSpinBox6suffixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "suffix", args)
  }

}

// singleStep()
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
    // invoke: double singleStep()
    var ret = C.C_ZNK14QDoubleSpinBox10singleStepEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "singleStep", args)
  }

}

// prefix()
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
    // invoke: QString prefix()
    var ret = C.C_ZNK14QDoubleSpinBox6prefixEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "prefix", args)
  }

}

// minimum()
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
    // invoke: double minimum()
    var ret = C.C_ZNK14QDoubleSpinBox7minimumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "minimum", args)
  }

}

// textFromValue(double)
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
    // invoke: QString textFromValue(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK14QDoubleSpinBox13textFromValueEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "textFromValue", args)
  }

}

// fixup(class QString &)
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
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK14QDoubleSpinBox5fixupER7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "fixup", args)
  }

}

// QDoubleSpinBox(class QWidget *)
func NewQDoubleSpinBox(args ...interface{}) *QDoubleSpinBox {
  // QDoubleSpinBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBoxC1EP7QWidget
    // invoke: void QDoubleSpinBox(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QDoubleSpinBoxC2EP7QWidget(arg0)
    return &QDoubleSpinBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "QDoubleSpinBox", args)
  }

  return nil // QDoubleSpinBox{qclsinst:qthis}
}

// setSingleStep(double)
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
    // invoke: void setSingleStep(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox13setSingleStepEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSingleStep", args)
  }

}

// setRange(double, double)
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
    // invoke: void setRange(double, double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN14QDoubleSpinBox8setRangeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setRange", args)
  }

}

// setMaximum(double)
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
    // invoke: void setMaximum(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox10setMaximumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMaximum", args)
  }

}

// setValue(double)
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
    // invoke: void setValue(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox8setValueEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setValue", args)
  }

}

// valueFromText(const class QString &)
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
    // invoke: double valueFromText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK14QDoubleSpinBox13valueFromTextERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "valueFromText", args)
  }

}

// validate(class QString &, int &)
func (this *QDoubleSpinBox) validate(args ...interface{}) () {
  // validate(class QString &, int &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK14QDoubleSpinBox8validateER7QStringRi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "validate", args)
  }

}

// cleanText()
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
    // invoke: QString cleanText()
    var ret = C.C_ZNK14QDoubleSpinBox9cleanTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "cleanText", args)
  }

}

// setDecimals(int)
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
    // invoke: void setDecimals(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox11setDecimalsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setDecimals", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QDoubleSpinBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "metaObject", args)
  }

}

// ~QDoubleSpinBox()
func (this *QDoubleSpinBox) FreeQDoubleSpinBox(args ...interface{}) () {
  // ~QDoubleSpinBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBoxD0Ev
    // invoke: void ~QDoubleSpinBox()
    C.C_ZN14QDoubleSpinBoxD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "~QDoubleSpinBox", args)
  }

}

// setSuffix(const class QString &)
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
    // invoke: void setSuffix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox9setSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSuffix", args)
  }

}

// maximum()
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
    // invoke: double maximum()
    var ret = C.C_ZNK14QDoubleSpinBox7maximumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "maximum", args)
  }

}

// value()
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
    // invoke: double value()
    var ret = C.C_ZNK14QDoubleSpinBox5valueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "value", args)
  }

}

// decimals()
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
    // invoke: int decimals()
    var ret = C.C_ZNK14QDoubleSpinBox8decimalsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "decimals", args)
  }

}

// setPrefix(const class QString &)
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
    // invoke: void setPrefix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox9setPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setPrefix", args)
  }

}

// setMinimum(double)
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
    // invoke: void setMinimum(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox10setMinimumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMinimum", args)
  }

}

// <= body block end

