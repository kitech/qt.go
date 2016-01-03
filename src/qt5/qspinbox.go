package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QSpinBox::setMinimum(int min);
extern void _ZN8QSpinBox10setMinimumEi(void* qthis, int32_t arg0);
  // proto:  QString QSpinBox::cleanText();
extern void _ZNK8QSpinBox9cleanTextEv(void* qthis);
  // proto:  int QSpinBox::value();
extern void _ZNK8QSpinBox5valueEv(void* qthis);
  // proto:  void QSpinBox::~QSpinBox();
extern void _ZN8QSpinBoxD0Ev(void* qthis);
  // proto:  void QSpinBox::setMaximum(int max);
extern void _ZN8QSpinBox10setMaximumEi(void* qthis, int32_t arg0);
  // proto:  void QSpinBox::setValue(int val);
extern void _ZN8QSpinBox8setValueEi(void* qthis, int32_t arg0);
  // proto:  void QSpinBox::setDisplayIntegerBase(int base);
extern void _ZN8QSpinBox21setDisplayIntegerBaseEi(void* qthis, int32_t arg0);
  // proto:  void QSpinBox::QSpinBox(QWidget * parent);
extern void* dector_ZN8QSpinBoxC1EP7QWidget(void* arg0);
extern void _ZN8QSpinBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  int QSpinBox::singleStep();
extern void _ZNK8QSpinBox10singleStepEv(void* qthis);
  // proto:  void QSpinBox::QSpinBox(const QSpinBox & );
extern void* dector_ZN8QSpinBoxC1ERKS_(void* arg0);
extern void _ZN8QSpinBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  int QSpinBox::displayIntegerBase();
extern void _ZNK8QSpinBox18displayIntegerBaseEv(void* qthis);
  // proto:  void QSpinBox::setSuffix(const QString & suffix);
extern void _ZN8QSpinBox9setSuffixERK7QString(void* qthis, void* arg0);
  // proto:  int QSpinBox::maximum();
extern void _ZNK8QSpinBox7maximumEv(void* qthis);
  // proto:  void QSpinBox::setPrefix(const QString & prefix);
extern void _ZN8QSpinBox9setPrefixERK7QString(void* qthis, void* arg0);
  // proto:  QString QSpinBox::prefix();
extern void _ZNK8QSpinBox6prefixEv(void* qthis);
  // proto:  const QMetaObject * QSpinBox::metaObject();
extern void _ZNK8QSpinBox10metaObjectEv(void* qthis);
  // proto:  QString QSpinBox::suffix();
extern void _ZNK8QSpinBox6suffixEv(void* qthis);
  // proto:  int QSpinBox::minimum();
extern void _ZNK8QSpinBox7minimumEv(void* qthis);
  // proto:  void QSpinBox::setSingleStep(int val);
extern void _ZN8QSpinBox13setSingleStepEi(void* qthis, int32_t arg0);
  // proto:  void QSpinBox::setRange(int min, int max);
extern void _ZN8QSpinBox8setRangeEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QString QDoubleSpinBox::textFromValue(double val);
extern void _ZNK14QDoubleSpinBox13textFromValueEd(void* qthis, double arg0);
  // proto:  void QDoubleSpinBox::setSingleStep(double val);
extern void _ZN14QDoubleSpinBox13setSingleStepEd(void* qthis, double arg0);
  // proto:  double QDoubleSpinBox::minimum();
extern void _ZNK14QDoubleSpinBox7minimumEv(void* qthis);
  // proto:  double QDoubleSpinBox::valueFromText(const QString & text);
extern void _ZNK14QDoubleSpinBox13valueFromTextERK7QString(void* qthis, void* arg0);
  // proto:  const QMetaObject * QDoubleSpinBox::metaObject();
extern void _ZNK14QDoubleSpinBox10metaObjectEv(void* qthis);
  // proto:  void QDoubleSpinBox::setValue(double val);
extern void _ZN14QDoubleSpinBox8setValueEd(void* qthis, double arg0);
  // proto:  void QDoubleSpinBox::setSuffix(const QString & suffix);
extern void _ZN14QDoubleSpinBox9setSuffixERK7QString(void* qthis, void* arg0);
  // proto:  int QDoubleSpinBox::decimals();
extern void _ZNK14QDoubleSpinBox8decimalsEv(void* qthis);
  // proto:  QString QDoubleSpinBox::prefix();
extern void _ZNK14QDoubleSpinBox6prefixEv(void* qthis);
  // proto:  double QDoubleSpinBox::singleStep();
extern void _ZNK14QDoubleSpinBox10singleStepEv(void* qthis);
  // proto:  void QDoubleSpinBox::~QDoubleSpinBox();
extern void _ZN14QDoubleSpinBoxD0Ev(void* qthis);
  // proto:  void QDoubleSpinBox::fixup(QString & str);
extern void _ZNK14QDoubleSpinBox5fixupER7QString(void* qthis, void* arg0);
  // proto:  void QDoubleSpinBox::QDoubleSpinBox(const QDoubleSpinBox & );
extern void* dector_ZN14QDoubleSpinBoxC1ERKS_(void* arg0);
extern void _ZN14QDoubleSpinBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDoubleSpinBox::setPrefix(const QString & prefix);
extern void _ZN14QDoubleSpinBox9setPrefixERK7QString(void* qthis, void* arg0);
  // proto:  QString QDoubleSpinBox::cleanText();
extern void _ZNK14QDoubleSpinBox9cleanTextEv(void* qthis);
  // proto:  void QDoubleSpinBox::setMinimum(double min);
extern void _ZN14QDoubleSpinBox10setMinimumEd(void* qthis, double arg0);
  // proto:  void QDoubleSpinBox::setMaximum(double max);
extern void _ZN14QDoubleSpinBox10setMaximumEd(void* qthis, double arg0);
  // proto:  void QDoubleSpinBox::setDecimals(int prec);
extern void _ZN14QDoubleSpinBox11setDecimalsEi(void* qthis, int32_t arg0);
  // proto:  double QDoubleSpinBox::value();
extern void _ZNK14QDoubleSpinBox5valueEv(void* qthis);
  // proto:  void QDoubleSpinBox::setRange(double min, double max);
extern void _ZN14QDoubleSpinBox8setRangeEdd(void* qthis, double arg0, double arg1);
  // proto:  void QDoubleSpinBox::QDoubleSpinBox(QWidget * parent);
extern void* dector_ZN14QDoubleSpinBoxC1EP7QWidget(void* arg0);
extern void _ZN14QDoubleSpinBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  double QDoubleSpinBox::maximum();
extern void _ZNK14QDoubleSpinBox7maximumEv(void* qthis);
  // proto:  QString QDoubleSpinBox::suffix();
extern void _ZNK14QDoubleSpinBox6suffixEv(void* qthis);
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

  // proto:  void QSpinBox::setMinimum(int min);
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
    C._ZN8QSpinBox10setMinimumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setMinimum", args)
  }

}

  // proto:  QString QSpinBox::cleanText();
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
    C._ZNK8QSpinBox9cleanTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "cleanText", args)
  }

}

  // proto:  int QSpinBox::value();
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
    C._ZNK8QSpinBox5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "value", args)
  }

}

  // proto:  void QSpinBox::~QSpinBox();
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

  // proto:  void QSpinBox::setMaximum(int max);
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
    C._ZN8QSpinBox10setMaximumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setMaximum", args)
  }

}

  // proto:  void QSpinBox::setValue(int val);
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
    C._ZN8QSpinBox8setValueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setValue", args)
  }

}

  // proto:  void QSpinBox::setDisplayIntegerBase(int base);
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
    C._ZN8QSpinBox21setDisplayIntegerBaseEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setDisplayIntegerBase", args)
  }

}

  // proto:  void QSpinBox::QSpinBox(QWidget * parent);
func NewQSpinBox(args ...interface{}) QSpinBox {
  return QSpinBox{}
}

  // proto:  int QSpinBox::singleStep();
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
    C._ZNK8QSpinBox10singleStepEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "singleStep", args)
  }

}

  // proto:  int QSpinBox::displayIntegerBase();
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
    C._ZNK8QSpinBox18displayIntegerBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "displayIntegerBase", args)
  }

}

  // proto:  void QSpinBox::setSuffix(const QString & suffix);
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
    C._ZN8QSpinBox9setSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setSuffix", args)
  }

}

  // proto:  int QSpinBox::maximum();
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
    C._ZNK8QSpinBox7maximumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "maximum", args)
  }

}

  // proto:  void QSpinBox::setPrefix(const QString & prefix);
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
    C._ZN8QSpinBox9setPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setPrefix", args)
  }

}

  // proto:  QString QSpinBox::prefix();
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
    C._ZNK8QSpinBox6prefixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "prefix", args)
  }

}

  // proto:  const QMetaObject * QSpinBox::metaObject();
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
    C._ZNK8QSpinBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "metaObject", args)
  }

}

  // proto:  QString QSpinBox::suffix();
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
    C._ZNK8QSpinBox6suffixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "suffix", args)
  }

}

  // proto:  int QSpinBox::minimum();
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
    C._ZNK8QSpinBox7minimumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "minimum", args)
  }

}

  // proto:  void QSpinBox::setSingleStep(int val);
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
    C._ZN8QSpinBox13setSingleStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setSingleStep", args)
  }

}

  // proto:  void QSpinBox::setRange(int min, int max);
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
    C._ZN8QSpinBox8setRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSpinBox", "setRange", args)
  }

}

  // proto:  QString QDoubleSpinBox::textFromValue(double val);
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
    C._ZNK14QDoubleSpinBox13textFromValueEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "textFromValue", args)
  }

}

  // proto:  void QDoubleSpinBox::setSingleStep(double val);
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
    C._ZN14QDoubleSpinBox13setSingleStepEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSingleStep", args)
  }

}

  // proto:  double QDoubleSpinBox::minimum();
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
    C._ZNK14QDoubleSpinBox7minimumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "minimum", args)
  }

}

  // proto:  double QDoubleSpinBox::valueFromText(const QString & text);
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
    C._ZNK14QDoubleSpinBox13valueFromTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "valueFromText", args)
  }

}

  // proto:  const QMetaObject * QDoubleSpinBox::metaObject();
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
    C._ZNK14QDoubleSpinBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "metaObject", args)
  }

}

  // proto:  void QDoubleSpinBox::setValue(double val);
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
    C._ZN14QDoubleSpinBox8setValueEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setValue", args)
  }

}

  // proto:  void QDoubleSpinBox::setSuffix(const QString & suffix);
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
    C._ZN14QDoubleSpinBox9setSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSuffix", args)
  }

}

  // proto:  int QDoubleSpinBox::decimals();
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
    C._ZNK14QDoubleSpinBox8decimalsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "decimals", args)
  }

}

  // proto:  QString QDoubleSpinBox::prefix();
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
    C._ZNK14QDoubleSpinBox6prefixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "prefix", args)
  }

}

  // proto:  double QDoubleSpinBox::singleStep();
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
    C._ZNK14QDoubleSpinBox10singleStepEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "singleStep", args)
  }

}

  // proto:  void QDoubleSpinBox::~QDoubleSpinBox();
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

  // proto:  void QDoubleSpinBox::fixup(QString & str);
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
    C._ZNK14QDoubleSpinBox5fixupER7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "fixup", args)
  }

}

  // proto:  void QDoubleSpinBox::QDoubleSpinBox(const QDoubleSpinBox & );
func NewQDoubleSpinBox(args ...interface{}) QDoubleSpinBox {
  return QDoubleSpinBox{}
}

  // proto:  void QDoubleSpinBox::setPrefix(const QString & prefix);
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
    C._ZN14QDoubleSpinBox9setPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setPrefix", args)
  }

}

  // proto:  QString QDoubleSpinBox::cleanText();
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
    C._ZNK14QDoubleSpinBox9cleanTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "cleanText", args)
  }

}

  // proto:  void QDoubleSpinBox::setMinimum(double min);
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
    C._ZN14QDoubleSpinBox10setMinimumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMinimum", args)
  }

}

  // proto:  void QDoubleSpinBox::setMaximum(double max);
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
    C._ZN14QDoubleSpinBox10setMaximumEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMaximum", args)
  }

}

  // proto:  void QDoubleSpinBox::setDecimals(int prec);
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
    C._ZN14QDoubleSpinBox11setDecimalsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setDecimals", args)
  }

}

  // proto:  double QDoubleSpinBox::value();
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
    C._ZNK14QDoubleSpinBox5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "value", args)
  }

}

  // proto:  void QDoubleSpinBox::setRange(double min, double max);
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
    C._ZN14QDoubleSpinBox8setRangeEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setRange", args)
  }

}

  // proto:  double QDoubleSpinBox::maximum();
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
    C._ZNK14QDoubleSpinBox7maximumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "maximum", args)
  }

}

  // proto:  QString QDoubleSpinBox::suffix();
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
    C._ZNK14QDoubleSpinBox6suffixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "suffix", args)
  }

}

// <= body block end

