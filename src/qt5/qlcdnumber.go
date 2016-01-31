package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QLCDNumber::setSmallDecimalPoint(bool );
extern void C_ZN10QLCDNumber20setSmallDecimalPointEb(void* qthis, bool arg0); // 4
  // proto:  void QLCDNumber::setDecMode();
extern void C_ZN10QLCDNumber10setDecModeEv(void* qthis); // 4
  // proto:  bool QLCDNumber::smallDecimalPoint();
extern bool C_ZNK10QLCDNumber17smallDecimalPointEv(void* qthis); // 4
  // proto:  void QLCDNumber::setHexMode();
extern void C_ZN10QLCDNumber10setHexModeEv(void* qthis); // 4
  // proto:  void QLCDNumber::~QLCDNumber();
extern void C_ZN10QLCDNumberD2Ev(void* qthis); // 4
  // proto:  QLCDNumber::SegmentStyle QLCDNumber::segmentStyle();
extern void C_ZNK10QLCDNumber12segmentStyleEv(void* qthis); // 4
  // proto:  void QLCDNumber::QLCDNumber(uint numDigits, QWidget * parent);
extern void* C_ZN10QLCDNumberC2EjP7QWidget(int32_t arg0, void* arg1); // 3
  // proto:  void QLCDNumber::QLCDNumber(QWidget * parent);
extern void* C_ZN10QLCDNumberC2EP7QWidget(void* arg0); // 3
  // proto:  int QLCDNumber::intValue();
extern int32_t C_ZNK10QLCDNumber8intValueEv(void* qthis); // 4
  // proto:  int QLCDNumber::digitCount();
extern int32_t C_ZNK10QLCDNumber10digitCountEv(void* qthis); // 4
  // proto:  QSize QLCDNumber::sizeHint();
extern void* C_ZNK10QLCDNumber8sizeHintEv(void* qthis); // 4
  // proto:  void QLCDNumber::setOctMode();
extern void C_ZN10QLCDNumber10setOctModeEv(void* qthis); // 4
  // proto:  const QMetaObject * QLCDNumber::metaObject();
extern void C_ZNK10QLCDNumber10metaObjectEv(void* qthis); // 4
  // proto:  void QLCDNumber::setBinMode();
extern void C_ZN10QLCDNumber10setBinModeEv(void* qthis); // 4
  // proto:  double QLCDNumber::value();
extern double C_ZNK10QLCDNumber5valueEv(void* qthis); // 4
  // proto:  void QLCDNumber::setDigitCount(int nDigits);
extern void C_ZN10QLCDNumber13setDigitCountEi(void* qthis, int32_t arg0); // 4
  // proto:  QLCDNumber::Mode QLCDNumber::mode();
extern void C_ZNK10QLCDNumber4modeEv(void* qthis); // 4
  // proto:  bool QLCDNumber::checkOverflow(int num);
extern bool C_ZNK10QLCDNumber13checkOverflowEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QLCDNumber::checkOverflow(double num);
extern bool C_ZNK10QLCDNumber13checkOverflowEd(void* qthis, double arg0); // 4
  // proto:  void QLCDNumber::display(int num);
extern void C_ZN10QLCDNumber7displayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLCDNumber::display(double num);
extern void C_ZN10QLCDNumber7displayEd(void* qthis, double arg0); // 4
  // proto:  void QLCDNumber::display(const QString & str);
extern void C_ZN10QLCDNumber7displayERK7QString(void* qthis, void* arg0); // 4
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

// setSmallDecimalPoint(_Bool)
func (this *QLCDNumber) Setsmalldecimalpoint(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber20setSmallDecimalPointEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setSmallDecimalPoint", args)
  }

  return
}

// setDecMode()
func (this *QLCDNumber) Setdecmode(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber10setDecModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setDecMode", args)
  }

  return
}

// smallDecimalPoint()
func (this *QLCDNumber) Smalldecimalpoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QLCDNumber17smallDecimalPointEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLCDNumber", "smallDecimalPoint", args)
  }

  return
}

// setHexMode()
func (this *QLCDNumber) Sethexmode(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber10setHexModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setHexMode", args)
  }

  return
}

// ~QLCDNumber()
func (this *QLCDNumber) Freeqlcdnumber(args ...interface{}) () {
  // ~QLCDNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumberD0Ev
    // invoke: void ~QLCDNumber()
    C.C_ZN10QLCDNumberD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "~QLCDNumber", args)
  }

  return
}

// segmentStyle()
func (this *QLCDNumber) Segmentstyle(args ...interface{}) () {
  // segmentStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber12segmentStyleEv
    // invoke: QLCDNumber::SegmentStyle segmentStyle()
    C.C_ZNK10QLCDNumber12segmentStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "segmentStyle", args)
  }

  return
}

// QLCDNumber(uint, class QWidget *)
func NewQLCDNumber(args ...interface{}) *QLCDNumber {
  // QLCDNumber(uint, class QWidget *)
  // QLCDNumber(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumberC1EjP7QWidget
    // invoke: void QLCDNumber(uint, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QLCDNumberC2EjP7QWidget(arg0, arg1)
    return &QLCDNumber{qclsinst:qthis}
  case 1:
    // invoke: _ZN10QLCDNumberC1EP7QWidget
    // invoke: void QLCDNumber(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QLCDNumberC2EP7QWidget(arg0)
    return &QLCDNumber{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLCDNumber", "QLCDNumber", args)
  }

  return nil // QLCDNumber{qclsinst:qthis}
}

// intValue()
func (this *QLCDNumber) Intvalue(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QLCDNumber8intValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLCDNumber", "intValue", args)
  }

  return
}

// digitCount()
func (this *QLCDNumber) Digitcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QLCDNumber10digitCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLCDNumber", "digitCount", args)
  }

  return
}

// sizeHint()
func (this *QLCDNumber) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QLCDNumber8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLCDNumber", "sizeHint", args)
  }

  return
}

// setOctMode()
func (this *QLCDNumber) Setoctmode(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber10setOctModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setOctMode", args)
  }

  return
}

// metaObject()
func (this *QLCDNumber) Metaobject(args ...interface{}) () {
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
    C.C_ZNK10QLCDNumber10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "metaObject", args)
  }

  return
}

// setBinMode()
func (this *QLCDNumber) Setbinmode(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber10setBinModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setBinMode", args)
  }

  return
}

// value()
func (this *QLCDNumber) Value(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QLCDNumber5valueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLCDNumber", "value", args)
  }

  return
}

// setDigitCount(int)
func (this *QLCDNumber) Setdigitcount(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber13setDigitCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "setDigitCount", args)
  }

  return
}

// mode()
func (this *QLCDNumber) Mode(args ...interface{}) () {
  // mode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber4modeEv
    // invoke: QLCDNumber::Mode mode()
    C.C_ZNK10QLCDNumber4modeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLCDNumber", "mode", args)
  }

  return
}

// checkOverflow(int)
func (this *QLCDNumber) Checkoverflow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QLCDNumber13checkOverflowEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK10QLCDNumber13checkOverflowEd
    // invoke: bool checkOverflow(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QLCDNumber13checkOverflowEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLCDNumber", "checkOverflow", args)
  }

  return
}

// display(int)
func (this *QLCDNumber) Display(args ...interface{}) () {
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
    C.C_ZN10QLCDNumber7displayEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QLCDNumber7displayEd
    // invoke: void display(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN10QLCDNumber7displayEd(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QLCDNumber7displayERK7QString
    // invoke: void display(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QLCDNumber7displayERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLCDNumber", "display", args)
  }

  return
}

// <= body block end

