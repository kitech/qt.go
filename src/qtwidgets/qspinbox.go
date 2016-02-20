package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QString QSpinBox::suffix();
extern void* C_ZNK8QSpinBox6suffixEv(void* qthis); // 4
  // proto:  void QSpinBox::QSpinBox(QWidget * parent);
extern void* C_ZN8QSpinBoxC2EP7QWidget(void* arg0); // 3
  // proto:  int QSpinBox::singleStep();
extern int32_t C_ZNK8QSpinBox10singleStepEv(void* qthis); // 4
  // proto:  QString QSpinBox::prefix();
extern void* C_ZNK8QSpinBox6prefixEv(void* qthis); // 4
  // proto:  int QSpinBox::minimum();
extern int32_t C_ZNK8QSpinBox7minimumEv(void* qthis); // 4
  // proto:  QString QSpinBox::cleanText();
extern void* C_ZNK8QSpinBox9cleanTextEv(void* qthis); // 4
  // proto:  void QSpinBox::setSingleStep(int val);
extern void C_ZN8QSpinBox13setSingleStepEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSpinBox::setDisplayIntegerBase(int base);
extern void C_ZN8QSpinBox21setDisplayIntegerBaseEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSpinBox::setMaximum(int max);
extern void C_ZN8QSpinBox10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSpinBox::setValue(int val);
extern void C_ZN8QSpinBox8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSpinBox::displayIntegerBase();
extern int32_t C_ZNK8QSpinBox18displayIntegerBaseEv(void* qthis); // 4
  // proto:  void QSpinBox::setRange(int min, int max);
extern void C_ZN8QSpinBox8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QSpinBox::metaObject();
extern void C_ZNK8QSpinBox10metaObjectEv(void* qthis); // 4
  // proto:  void QSpinBox::~QSpinBox();
extern void C_ZN8QSpinBoxD2Ev(void* qthis); // 4
  // proto:  void QSpinBox::setSuffix(const QString & suffix);
extern void C_ZN8QSpinBox9setSuffixERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QSpinBox::maximum();
extern int32_t C_ZNK8QSpinBox7maximumEv(void* qthis); // 4
  // proto:  int QSpinBox::value();
extern int32_t C_ZNK8QSpinBox5valueEv(void* qthis); // 4
  // proto:  void QSpinBox::setPrefix(const QString & prefix);
extern void C_ZN8QSpinBox9setPrefixERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSpinBox::setMinimum(int min);
extern void C_ZN8QSpinBox10setMinimumEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QDoubleSpinBox::suffix();
extern void* C_ZNK14QDoubleSpinBox6suffixEv(void* qthis); // 4
  // proto:  double QDoubleSpinBox::singleStep();
extern double C_ZNK14QDoubleSpinBox10singleStepEv(void* qthis); // 4
  // proto:  QString QDoubleSpinBox::prefix();
extern void* C_ZNK14QDoubleSpinBox6prefixEv(void* qthis); // 4
  // proto:  double QDoubleSpinBox::minimum();
extern double C_ZNK14QDoubleSpinBox7minimumEv(void* qthis); // 4
  // proto:  QString QDoubleSpinBox::textFromValue(double val);
extern void* C_ZNK14QDoubleSpinBox13textFromValueEd(void* qthis, double arg0); // 4
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
extern double C_ZNK14QDoubleSpinBox13valueFromTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QValidator::State QDoubleSpinBox::validate(QString & input, int & pos);
extern void C_ZNK14QDoubleSpinBox8validateER7QStringRi(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QDoubleSpinBox::cleanText();
extern void* C_ZNK14QDoubleSpinBox9cleanTextEv(void* qthis); // 4
  // proto:  void QDoubleSpinBox::setDecimals(int prec);
extern void C_ZN14QDoubleSpinBox11setDecimalsEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QDoubleSpinBox::metaObject();
extern void C_ZNK14QDoubleSpinBox10metaObjectEv(void* qthis); // 4
  // proto:  void QDoubleSpinBox::~QDoubleSpinBox();
extern void C_ZN14QDoubleSpinBoxD2Ev(void* qthis); // 4
  // proto:  void QDoubleSpinBox::setSuffix(const QString & suffix);
extern void C_ZN14QDoubleSpinBox9setSuffixERK7QString(void* qthis, void* arg0); // 4
  // proto:  double QDoubleSpinBox::maximum();
extern double C_ZNK14QDoubleSpinBox7maximumEv(void* qthis); // 4
  // proto:  double QDoubleSpinBox::value();
extern double C_ZNK14QDoubleSpinBox5valueEv(void* qthis); // 4
  // proto:  int QDoubleSpinBox::decimals();
extern int32_t C_ZNK14QDoubleSpinBox8decimalsEv(void* qthis); // 4
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
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QSpinBox)=1
type QSpinBox struct {
  /*qbase*/ QAbstractSpinBox;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QSpinBox_valueChanged_signal;
}

// class sizeof(QDoubleSpinBox)=1
type QDoubleSpinBox struct {
  /*qbase*/ QAbstractSpinBox;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QDoubleSpinBox_valueChanged_signal;
}

// suffix()
func (this *QSpinBox) Suffix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox6suffixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "suffix", args)
  }

  return
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QSpinBoxC2EP7QWidget(arg0)
    return &QSpinBox{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSpinBox", "QSpinBox", args)
  }

  return nil // QSpinBox{Qclsinst:qthis}
}

// singleStep()
func (this *QSpinBox) Singlestep(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox10singleStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "singleStep", args)
  }

  return
}

// prefix()
func (this *QSpinBox) Prefix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox6prefixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "prefix", args)
  }

  return
}

// minimum()
func (this *QSpinBox) Minimum(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox7minimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "minimum", args)
  }

  return
}

// cleanText()
func (this *QSpinBox) Cleantext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox9cleanTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "cleanText", args)
  }

  return
}

// setSingleStep(int)
func (this *QSpinBox) Setsinglestep(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox13setSingleStepEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setSingleStep", args)
  }

  return
}

// setDisplayIntegerBase(int)
func (this *QSpinBox) Setdisplayintegerbase(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox21setDisplayIntegerBaseEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setDisplayIntegerBase", args)
  }

  return
}

// setMaximum(int)
func (this *QSpinBox) Setmaximum(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox10setMaximumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setMaximum", args)
  }

  return
}

// setValue(int)
func (this *QSpinBox) Setvalue(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox8setValueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setValue", args)
  }

  return
}

// displayIntegerBase()
func (this *QSpinBox) Displayintegerbase(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox18displayIntegerBaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "displayIntegerBase", args)
  }

  return
}

// setRange(int, int)
func (this *QSpinBox) Setrange(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QSpinBox8setRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSpinBox", "setRange", args)
  }

  return
}

// metaObject()
func (this *QSpinBox) Metaobject(args ...interface{}) () {
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
    C.C_ZNK8QSpinBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "metaObject", args)
  }

  return
}

// ~QSpinBox()
func (this *QSpinBox) Freeqspinbox(args ...interface{}) () {
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
    C.C_ZN8QSpinBoxD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSpinBox", "~QSpinBox", args)
  }

  return
}

// setSuffix(const class QString &)
func (this *QSpinBox) Setsuffix(args ...interface{}) () {
  // setSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox9setSuffixERK7QString
    // invoke: void setSuffix(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox9setSuffixERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setSuffix", args)
  }

  return
}

// maximum()
func (this *QSpinBox) Maximum(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox7maximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "maximum", args)
  }

  return
}

// value()
func (this *QSpinBox) Value(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QSpinBox5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSpinBox", "value", args)
  }

  return
}

// setPrefix(const class QString &)
func (this *QSpinBox) Setprefix(args ...interface{}) () {
  // setPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSpinBox9setPrefixERK7QString
    // invoke: void setPrefix(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox9setPrefixERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setPrefix", args)
  }

  return
}

// setMinimum(int)
func (this *QSpinBox) Setminimum(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QSpinBox10setMinimumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSpinBox", "setMinimum", args)
  }

  return
}

// suffix()
func (this *QDoubleSpinBox) Suffix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox6suffixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "suffix", args)
  }

  return
}

// singleStep()
func (this *QDoubleSpinBox) Singlestep(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox10singleStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "singleStep", args)
  }

  return
}

// prefix()
func (this *QDoubleSpinBox) Prefix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox6prefixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "prefix", args)
  }

  return
}

// minimum()
func (this *QDoubleSpinBox) Minimum(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox7minimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "minimum", args)
  }

  return
}

// textFromValue(double)
func (this *QDoubleSpinBox) Textfromvalue(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDoubleSpinBox13textFromValueEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "textFromValue", args)
  }

  return
}

// fixup(class QString &)
func (this *QDoubleSpinBox) Fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox5fixupER7QString
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK14QDoubleSpinBox5fixupER7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "fixup", args)
  }

  return
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QDoubleSpinBoxC2EP7QWidget(arg0)
    return &QDoubleSpinBox{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "QDoubleSpinBox", args)
  }

  return nil // QDoubleSpinBox{Qclsinst:qthis}
}

// setSingleStep(double)
func (this *QDoubleSpinBox) Setsinglestep(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox13setSingleStepEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSingleStep", args)
  }

  return
}

// setRange(double, double)
func (this *QDoubleSpinBox) Setrange(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN14QDoubleSpinBox8setRangeEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setRange", args)
  }

  return
}

// setMaximum(double)
func (this *QDoubleSpinBox) Setmaximum(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox10setMaximumEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMaximum", args)
  }

  return
}

// setValue(double)
func (this *QDoubleSpinBox) Setvalue(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox8setValueEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setValue", args)
  }

  return
}

// valueFromText(const class QString &)
func (this *QDoubleSpinBox) Valuefromtext(args ...interface{}) (ret interface{}) {
  // valueFromText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox13valueFromTextERK7QString
    // invoke: double valueFromText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDoubleSpinBox13valueFromTextERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "valueFromText", args)
  }

  return
}

// validate(class QString &, int &)
func (this *QDoubleSpinBox) Validate(args ...interface{}) () {
  // validate(class QString &, int &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDoubleSpinBox8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK14QDoubleSpinBox8validateER7QStringRi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "validate", args)
  }

  return
}

// cleanText()
func (this *QDoubleSpinBox) Cleantext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox9cleanTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "cleanText", args)
  }

  return
}

// setDecimals(int)
func (this *QDoubleSpinBox) Setdecimals(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox11setDecimalsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setDecimals", args)
  }

  return
}

// metaObject()
func (this *QDoubleSpinBox) Metaobject(args ...interface{}) () {
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
    C.C_ZNK14QDoubleSpinBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "metaObject", args)
  }

  return
}

// ~QDoubleSpinBox()
func (this *QDoubleSpinBox) Freeqdoublespinbox(args ...interface{}) () {
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
    C.C_ZN14QDoubleSpinBoxD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "~QDoubleSpinBox", args)
  }

  return
}

// setSuffix(const class QString &)
func (this *QDoubleSpinBox) Setsuffix(args ...interface{}) () {
  // setSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox9setSuffixERK7QString
    // invoke: void setSuffix(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox9setSuffixERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setSuffix", args)
  }

  return
}

// maximum()
func (this *QDoubleSpinBox) Maximum(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox7maximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "maximum", args)
  }

  return
}

// value()
func (this *QDoubleSpinBox) Value(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "value", args)
  }

  return
}

// decimals()
func (this *QDoubleSpinBox) Decimals(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QDoubleSpinBox8decimalsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "decimals", args)
  }

  return
}

// setPrefix(const class QString &)
func (this *QDoubleSpinBox) Setprefix(args ...interface{}) () {
  // setPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDoubleSpinBox9setPrefixERK7QString
    // invoke: void setPrefix(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox9setPrefixERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setPrefix", args)
  }

  return
}

// setMinimum(double)
func (this *QDoubleSpinBox) Setminimum(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN14QDoubleSpinBox10setMinimumEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleSpinBox", "setMinimum", args)
  }

  return
}

// <= body block end

