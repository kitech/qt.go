package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qlcdnumber.h
// dst-file: /src/widgets/qlcdnumber.go
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
  // proto:  void QLCDNumber::display(int num);
extern void _ZN10QLCDNumber7displayEi(void* qthis, int32_t arg0);
  // proto:  void QLCDNumber::setHexMode();
extern void _ZN10QLCDNumber10setHexModeEv(void* qthis);
  // proto:  void QLCDNumber::display(double num);
extern void _ZN10QLCDNumber7displayEd(void* qthis, double arg0);
  // proto:  const QMetaObject * QLCDNumber::metaObject();
extern void _ZNK10QLCDNumber10metaObjectEv(void* qthis);
  // proto:  void QLCDNumber::QLCDNumber(const QLCDNumber & );
extern void* dector_ZN10QLCDNumberC1ERKS_(void* arg0);
extern void _ZN10QLCDNumberC1ERKS_(void* qthis, void* arg0);
  // proto:  int QLCDNumber::digitCount();
extern void _ZNK10QLCDNumber10digitCountEv(void* qthis);
  // proto:  void QLCDNumber::~QLCDNumber();
extern void _ZN10QLCDNumberD0Ev(void* qthis);
  // proto:  bool QLCDNumber::checkOverflow(int num);
extern void _ZNK10QLCDNumber13checkOverflowEi(void* qthis, int32_t arg0);
  // proto:  void QLCDNumber::setDecMode();
extern void _ZN10QLCDNumber10setDecModeEv(void* qthis);
  // proto:  void QLCDNumber::QLCDNumber(uint numDigits, QWidget * parent);
extern void* dector_ZN10QLCDNumberC1EjP7QWidget(int32_t arg0, void* arg1);
extern void _ZN10QLCDNumberC1EjP7QWidget(void* qthis, int32_t arg0, void* arg1);
  // proto:  bool QLCDNumber::checkOverflow(double num);
extern void _ZNK10QLCDNumber13checkOverflowEd(void* qthis, double arg0);
  // proto:  QSize QLCDNumber::sizeHint();
extern void _ZNK10QLCDNumber8sizeHintEv(void* qthis);
  // proto:  void QLCDNumber::display(const QString & str);
extern void _ZN10QLCDNumber7displayERK7QString(void* qthis, void* arg0);
  // proto:  void QLCDNumber::QLCDNumber(QWidget * parent);
extern void* dector_ZN10QLCDNumberC1EP7QWidget(void* arg0);
extern void _ZN10QLCDNumberC1EP7QWidget(void* qthis, void* arg0);
  // proto:  double QLCDNumber::value();
extern void _ZNK10QLCDNumber5valueEv(void* qthis);
  // proto:  void QLCDNumber::setBinMode();
extern void _ZN10QLCDNumber10setBinModeEv(void* qthis);
  // proto:  int QLCDNumber::intValue();
extern void _ZNK10QLCDNumber8intValueEv(void* qthis);
  // proto:  void QLCDNumber::setDigitCount(int nDigits);
extern void _ZN10QLCDNumber13setDigitCountEi(void* qthis, int32_t arg0);
  // proto:  void QLCDNumber::setSmallDecimalPoint(bool );
extern void _ZN10QLCDNumber20setSmallDecimalPointEb(void* qthis, bool arg0);
  // proto:  bool QLCDNumber::smallDecimalPoint();
extern void _ZNK10QLCDNumber17smallDecimalPointEv(void* qthis);
  // proto:  void QLCDNumber::setOctMode();
extern void _ZN10QLCDNumber10setOctModeEv(void* qthis);
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

// class sizeof(QLCDNumber)=1
type QLCDNumber struct {
  /*qbase*/ QFrame;
  qclsinst unsafe.Pointer /* *C.void */;
//  _overflow QLCDNumber_overflow_signal;
}

  // proto:  void QLCDNumber::display(int num);
func (this *QLCDNumber) display(args ...interface{}) () {
  // display(int)
  // display(double)
  // display(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber7displayEi
    // invoke: void display(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QLCDNumber7displayEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QLCDNumber7displayEd
    // invoke: void display(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN10QLCDNumber7displayEd(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QLCDNumber7displayERK7QString
    // invoke: void display(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QLCDNumber7displayERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "display", args)
  }

}

  // proto:  void QLCDNumber::setHexMode();
func (this *QLCDNumber) setHexMode(args ...interface{}) () {
  // setHexMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setHexModeEv
    // invoke: void setHexMode()
    C._ZN10QLCDNumber10setHexModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setHexMode", args)
  }

}

  // proto:  const QMetaObject * QLCDNumber::metaObject();
func (this *QLCDNumber) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QLCDNumber10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "metaObject", args)
  }

}

  // proto:  void QLCDNumber::QLCDNumber(const QLCDNumber & );
func NewQLCDNumber(args ...interface{}) QLCDNumber {
  return QLCDNumber{}
}

  // proto:  int QLCDNumber::digitCount();
func (this *QLCDNumber) digitCount(args ...interface{}) () {
  // digitCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber10digitCountEv
    // invoke: int digitCount()
    C._ZNK10QLCDNumber10digitCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "digitCount", args)
  }

}

  // proto:  void QLCDNumber::~QLCDNumber();
func (this *QLCDNumber) FreeQLCDNumber(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLCDNumber", "~QLCDNumber", args)
  }

}

  // proto:  bool QLCDNumber::checkOverflow(int num);
func (this *QLCDNumber) checkOverflow(args ...interface{}) () {
  // checkOverflow(int)
  // checkOverflow(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber13checkOverflowEi
    // invoke: bool checkOverflow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QLCDNumber13checkOverflowEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QLCDNumber13checkOverflowEd
    // invoke: bool checkOverflow(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZNK10QLCDNumber13checkOverflowEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "checkOverflow", args)
  }

}

  // proto:  void QLCDNumber::setDecMode();
func (this *QLCDNumber) setDecMode(args ...interface{}) () {
  // setDecMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setDecModeEv
    // invoke: void setDecMode()
    C._ZN10QLCDNumber10setDecModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setDecMode", args)
  }

}

  // proto:  QSize QLCDNumber::sizeHint();
func (this *QLCDNumber) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK10QLCDNumber8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "sizeHint", args)
  }

}

  // proto:  double QLCDNumber::value();
func (this *QLCDNumber) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber5valueEv
    // invoke: double value()
    C._ZNK10QLCDNumber5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "value", args)
  }

}

  // proto:  void QLCDNumber::setBinMode();
func (this *QLCDNumber) setBinMode(args ...interface{}) () {
  // setBinMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setBinModeEv
    // invoke: void setBinMode()
    C._ZN10QLCDNumber10setBinModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setBinMode", args)
  }

}

  // proto:  int QLCDNumber::intValue();
func (this *QLCDNumber) intValue(args ...interface{}) () {
  // intValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber8intValueEv
    // invoke: int intValue()
    C._ZNK10QLCDNumber8intValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "intValue", args)
  }

}

  // proto:  void QLCDNumber::setDigitCount(int nDigits);
func (this *QLCDNumber) setDigitCount(args ...interface{}) () {
  // setDigitCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber13setDigitCountEi
    // invoke: void setDigitCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QLCDNumber13setDigitCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setDigitCount", args)
  }

}

  // proto:  void QLCDNumber::setSmallDecimalPoint(bool );
func (this *QLCDNumber) setSmallDecimalPoint(args ...interface{}) () {
  // setSmallDecimalPoint(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber20setSmallDecimalPointEb
    // invoke: void setSmallDecimalPoint(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QLCDNumber20setSmallDecimalPointEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setSmallDecimalPoint", args)
  }

}

  // proto:  bool QLCDNumber::smallDecimalPoint();
func (this *QLCDNumber) smallDecimalPoint(args ...interface{}) () {
  // smallDecimalPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber17smallDecimalPointEv
    // invoke: bool smallDecimalPoint()
    C._ZNK10QLCDNumber17smallDecimalPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "smallDecimalPoint", args)
  }

}

  // proto:  void QLCDNumber::setOctMode();
func (this *QLCDNumber) setOctMode(args ...interface{}) () {
  // setOctMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setOctModeEv
    // invoke: void setOctMode()
    C._ZN10QLCDNumber10setOctModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setOctMode", args)
  }

}

// <= body block end

