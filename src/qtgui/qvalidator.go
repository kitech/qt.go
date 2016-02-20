package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtGui/qvalidator.h
// dst-file: /src/gui/qvalidator.go
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QRegularExpressionValidator::metaObject();
extern void C_ZNK27QRegularExpressionValidator10metaObjectEv(void* qthis); // 4
  // proto:  QRegularExpression QRegularExpressionValidator::regularExpression();
extern void* C_ZNK27QRegularExpressionValidator17regularExpressionEv(void* qthis); // 4
  // proto:  void QRegularExpressionValidator::~QRegularExpressionValidator();
extern void C_ZN27QRegularExpressionValidatorD2Ev(void* qthis); // 4
  // proto:  void QRegularExpressionValidator::setRegularExpression(const QRegularExpression & re);
extern void C_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression(void* qthis, void* arg0); // 4
  // proto:  QValidator::State QRegularExpressionValidator::validate(QString & input, int & pos);
extern void C_ZNK27QRegularExpressionValidator8validateER7QStringRi(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(QObject * parent);
extern void* C_ZN27QRegularExpressionValidatorC2EP7QObject(void* arg0); // 3
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(const QRegularExpression & re, QObject * parent);
extern void* C_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QDoubleValidator::setBottom(double );
extern void C_ZN16QDoubleValidator9setBottomEd(void* qthis, double arg0); // 4
  // proto:  void QDoubleValidator::setTop(double );
extern void C_ZN16QDoubleValidator6setTopEd(void* qthis, double arg0); // 4
  // proto:  double QDoubleValidator::bottom();
extern double C_ZNK16QDoubleValidator6bottomEv(void* qthis); // 2
  // proto:  void QDoubleValidator::setDecimals(int );
extern void C_ZN16QDoubleValidator11setDecimalsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDoubleValidator::~QDoubleValidator();
extern void C_ZN16QDoubleValidatorD2Ev(void* qthis); // 4
  // proto:  QDoubleValidator::Notation QDoubleValidator::notation();
extern void C_ZNK16QDoubleValidator8notationEv(void* qthis); // 4
  // proto:  QValidator::State QDoubleValidator::validate(QString & , int & );
extern void C_ZNK16QDoubleValidator8validateER7QStringRi(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QDoubleValidator::setRange(double bottom, double top, int decimals);
extern void C_ZN16QDoubleValidator8setRangeEddi(void* qthis, double arg0, double arg1, int32_t arg2); // 4
  // proto:  const QMetaObject * QDoubleValidator::metaObject();
extern void C_ZNK16QDoubleValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QDoubleValidator::QDoubleValidator(QObject * parent);
extern void* C_ZN16QDoubleValidatorC2EP7QObject(void* arg0); // 3
  // proto:  void QDoubleValidator::QDoubleValidator(double bottom, double top, int decimals, QObject * parent);
extern void* C_ZN16QDoubleValidatorC2EddiP7QObject(double arg0, double arg1, int32_t arg2, void* arg3); // 3
  // proto:  int QDoubleValidator::decimals();
extern int32_t C_ZNK16QDoubleValidator8decimalsEv(void* qthis); // 2
  // proto:  double QDoubleValidator::top();
extern double C_ZNK16QDoubleValidator3topEv(void* qthis); // 2
  // proto:  void QIntValidator::setTop(int );
extern void C_ZN13QIntValidator6setTopEi(void* qthis, int32_t arg0); // 4
  // proto:  void QIntValidator::QIntValidator(QObject * parent);
extern void* C_ZN13QIntValidatorC2EP7QObject(void* arg0); // 3
  // proto:  void QIntValidator::QIntValidator(int bottom, int top, QObject * parent);
extern void* C_ZN13QIntValidatorC2EiiP7QObject(int32_t arg0, int32_t arg1, void* arg2); // 3
  // proto:  const QMetaObject * QIntValidator::metaObject();
extern void C_ZNK13QIntValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QIntValidator::~QIntValidator();
extern void C_ZN13QIntValidatorD2Ev(void* qthis); // 4
  // proto:  int QIntValidator::bottom();
extern int32_t C_ZNK13QIntValidator6bottomEv(void* qthis); // 2
  // proto:  void QIntValidator::setBottom(int );
extern void C_ZN13QIntValidator9setBottomEi(void* qthis, int32_t arg0); // 4
  // proto:  int QIntValidator::top();
extern int32_t C_ZNK13QIntValidator3topEv(void* qthis); // 2
  // proto:  QValidator::State QIntValidator::validate(QString & , int & );
extern void C_ZNK13QIntValidator8validateER7QStringRi(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QIntValidator::fixup(QString & input);
extern void C_ZNK13QIntValidator5fixupER7QString(void* qthis, void* arg0); // 4
  // proto:  void QIntValidator::setRange(int bottom, int top);
extern void C_ZN13QIntValidator8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QValidator::metaObject();
extern void C_ZNK10QValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QValidator::setLocale(const QLocale & locale);
extern void C_ZN10QValidator9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QLocale QValidator::locale();
extern void* C_ZNK10QValidator6localeEv(void* qthis); // 4
  // proto:  void QValidator::~QValidator();
extern void C_ZN10QValidatorD2Ev(void* qthis); // 4
  // proto:  void QValidator::QValidator(QObject * parent);
extern void* C_ZN10QValidatorC2EP7QObject(void* arg0); // 3
  // proto:  void QValidator::fixup(QString & );
extern void C_ZNK10QValidator5fixupER7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QRegExpValidator::metaObject();
extern void C_ZNK16QRegExpValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QRegExpValidator::QRegExpValidator(const QRegExp & rx, QObject * parent);
extern void* C_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QRegExpValidator::QRegExpValidator(QObject * parent);
extern void* C_ZN16QRegExpValidatorC2EP7QObject(void* arg0); // 3
  // proto:  void QRegExpValidator::setRegExp(const QRegExp & rx);
extern void C_ZN16QRegExpValidator9setRegExpERK7QRegExp(void* qthis, void* arg0); // 4
  // proto:  const QRegExp & QRegExpValidator::regExp();
extern void* C_ZNK16QRegExpValidator6regExpEv(void* qthis); // 2
  // proto:  QValidator::State QRegExpValidator::validate(QString & input, int & pos);
extern void C_ZNK16QRegExpValidator8validateER7QStringRi(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QRegExpValidator::~QRegExpValidator();
extern void C_ZN16QRegExpValidatorD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QRegularExpressionValidator)=1
type QRegularExpressionValidator struct {
  /*qbase*/ QValidator;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _regularExpressionChanged QRegularExpressionValidator_regularExpressionChanged_signal;
}

// class sizeof(QDoubleValidator)=1
type QDoubleValidator struct {
  /*qbase*/ QValidator;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _topChanged QDoubleValidator_topChanged_signal;
//  _notationChanged QDoubleValidator_notationChanged_signal;
//  _bottomChanged QDoubleValidator_bottomChanged_signal;
//  _decimalsChanged QDoubleValidator_decimalsChanged_signal;
}

// class sizeof(QIntValidator)=1
type QIntValidator struct {
  /*qbase*/ QValidator;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _topChanged QIntValidator_topChanged_signal;
//  _bottomChanged QIntValidator_bottomChanged_signal;
}

// class sizeof(QValidator)=1
type QValidator struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _changed QValidator_changed_signal;
}

// class sizeof(QRegExpValidator)=1
type QRegExpValidator struct {
  /*qbase*/ QValidator;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _regExpChanged QRegExpValidator_regExpChanged_signal;
}

// metaObject()
func (this *QRegularExpressionValidator) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QRegularExpressionValidator10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK27QRegularExpressionValidator10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "metaObject", args)
  }

  return
}

// regularExpression()
func (this *QRegularExpressionValidator) Regularexpression(args ...interface{}) (ret interface{}) {
  // regularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QRegularExpressionValidator17regularExpressionEv
    // invoke: QRegularExpression regularExpression()
    var ret0 = C.C_ZNK27QRegularExpressionValidator17regularExpressionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRegularExpression{}) // "QRegularExpression"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "regularExpression", args)
  }

  return
}

// ~QRegularExpressionValidator()
func (this *QRegularExpressionValidator) Freeqregularexpressionvalidator(args ...interface{}) () {
  // ~QRegularExpressionValidator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QRegularExpressionValidatorD0Ev
    // invoke: void ~QRegularExpressionValidator()
    C.C_ZN27QRegularExpressionValidatorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "~QRegularExpressionValidator", args)
  }

  return
}

// setRegularExpression(const class QRegularExpression &)
func (this *QRegularExpressionValidator) Setregularexpression(args ...interface{}) () {
  // setRegularExpression(const class QRegularExpression &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRegularExpression{}) // "const QRegularExpression &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression
    // invoke: void setRegularExpression(const class QRegularExpression &)
    var arg0 = args[0].(*qtcore.QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "setRegularExpression", args)
  }

  return
}

// validate(class QString &, int &)
func (this *QRegularExpressionValidator) Validate(args ...interface{}) () {
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
    // invoke: _ZNK27QRegularExpressionValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK27QRegularExpressionValidator8validateER7QStringRi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "validate", args)
  }

  return
}

// QRegularExpressionValidator(class QObject *)
func NewQRegularExpressionValidator(args ...interface{}) *QRegularExpressionValidator {
  // QRegularExpressionValidator(class QObject *)
  // QRegularExpressionValidator(const class QRegularExpression &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QRegularExpressionValidatorC1EP7QObject
    // invoke: void QRegularExpressionValidator(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QRegularExpressionValidatorC2EP7QObject(arg0)
    return &QRegularExpressionValidator{Qclsinst:qthis}
  case 1:
    // invoke: _ZN27QRegularExpressionValidatorC1ERK18QRegularExpressionP7QObject
    // invoke: void QRegularExpressionValidator(const class QRegularExpression &, class QObject *)
    var arg0 = args[0].(*qtcore.QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject(arg0, arg1)
    return &QRegularExpressionValidator{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "QRegularExpressionValidator", args)
  }

  return nil // QRegularExpressionValidator{Qclsinst:qthis}
}

// setBottom(double)
func (this *QDoubleValidator) Setbottom(args ...interface{}) () {
  // setBottom(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidator9setBottomEd
    // invoke: void setBottom(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDoubleValidator9setBottomEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setBottom", args)
  }

  return
}

// setTop(double)
func (this *QDoubleValidator) Settop(args ...interface{}) () {
  // setTop(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidator6setTopEd
    // invoke: void setTop(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDoubleValidator6setTopEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setTop", args)
  }

  return
}

// bottom()
func (this *QDoubleValidator) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDoubleValidator6bottomEv
    // invoke: double bottom()
    var ret0 = C.C_ZNK16QDoubleValidator6bottomEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleValidator", "bottom", args)
  }

  return
}

// setDecimals(int)
func (this *QDoubleValidator) Setdecimals(args ...interface{}) () {
  // setDecimals(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidator11setDecimalsEi
    // invoke: void setDecimals(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDoubleValidator11setDecimalsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setDecimals", args)
  }

  return
}

// ~QDoubleValidator()
func (this *QDoubleValidator) Freeqdoublevalidator(args ...interface{}) () {
  // ~QDoubleValidator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidatorD0Ev
    // invoke: void ~QDoubleValidator()
    C.C_ZN16QDoubleValidatorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "~QDoubleValidator", args)
  }

  return
}

// notation()
func (this *QDoubleValidator) Notation(args ...interface{}) () {
  // notation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDoubleValidator8notationEv
    // invoke: QDoubleValidator::Notation notation()
    C.C_ZNK16QDoubleValidator8notationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "notation", args)
  }

  return
}

// validate(class QString &, int &)
func (this *QDoubleValidator) Validate(args ...interface{}) () {
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
    // invoke: _ZNK16QDoubleValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK16QDoubleValidator8validateER7QStringRi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "validate", args)
  }

  return
}

// setRange(double, double, int)
func (this *QDoubleValidator) Setrange(args ...interface{}) () {
  // setRange(double, double, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[0][1] = qtrt.DoubleTy(false) // "double"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidator8setRangeEddi
    // invoke: void setRange(double, double, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN16QDoubleValidator8setRangeEddi(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setRange", args)
  }

  return
}

// metaObject()
func (this *QDoubleValidator) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDoubleValidator10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QDoubleValidator10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "metaObject", args)
  }

  return
}

// QDoubleValidator(class QObject *)
func NewQDoubleValidator(args ...interface{}) *QDoubleValidator {
  // QDoubleValidator(class QObject *)
  // QDoubleValidator(double, double, int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"
  vtys[1][1] = qtrt.DoubleTy(false) // "double"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidatorC1EP7QObject
    // invoke: void QDoubleValidator(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QDoubleValidatorC2EP7QObject(arg0)
    return &QDoubleValidator{Qclsinst:qthis}
  case 1:
    // invoke: _ZN16QDoubleValidatorC1EddiP7QObject
    // invoke: void QDoubleValidator(double, double, int, class QObject *)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QDoubleValidatorC2EddiP7QObject(arg0, arg1, arg2, arg3)
    return &QDoubleValidator{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDoubleValidator", "QDoubleValidator", args)
  }

  return nil // QDoubleValidator{Qclsinst:qthis}
}

// decimals()
func (this *QDoubleValidator) Decimals(args ...interface{}) (ret interface{}) {
  // decimals()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDoubleValidator8decimalsEv
    // invoke: int decimals()
    var ret0 = C.C_ZNK16QDoubleValidator8decimalsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleValidator", "decimals", args)
  }

  return
}

// top()
func (this *QDoubleValidator) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QDoubleValidator3topEv
    // invoke: double top()
    var ret0 = C.C_ZNK16QDoubleValidator3topEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDoubleValidator", "top", args)
  }

  return
}

// setTop(int)
func (this *QIntValidator) Settop(args ...interface{}) () {
  // setTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QIntValidator6setTopEi
    // invoke: void setTop(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QIntValidator6setTopEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIntValidator", "setTop", args)
  }

  return
}

// QIntValidator(class QObject *)
func NewQIntValidator(args ...interface{}) *QIntValidator {
  // QIntValidator(class QObject *)
  // QIntValidator(int, int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QIntValidatorC1EP7QObject
    // invoke: void QIntValidator(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QIntValidatorC2EP7QObject(arg0)
    return &QIntValidator{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QIntValidatorC1EiiP7QObject
    // invoke: void QIntValidator(int, int, class QObject *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QIntValidatorC2EiiP7QObject(arg0, arg1, arg2)
    return &QIntValidator{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QIntValidator", "QIntValidator", args)
  }

  return nil // QIntValidator{Qclsinst:qthis}
}

// metaObject()
func (this *QIntValidator) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QIntValidator10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QIntValidator10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIntValidator", "metaObject", args)
  }

  return
}

// ~QIntValidator()
func (this *QIntValidator) Freeqintvalidator(args ...interface{}) () {
  // ~QIntValidator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QIntValidatorD0Ev
    // invoke: void ~QIntValidator()
    C.C_ZN13QIntValidatorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIntValidator", "~QIntValidator", args)
  }

  return
}

// bottom()
func (this *QIntValidator) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QIntValidator6bottomEv
    // invoke: int bottom()
    var ret0 = C.C_ZNK13QIntValidator6bottomEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIntValidator", "bottom", args)
  }

  return
}

// setBottom(int)
func (this *QIntValidator) Setbottom(args ...interface{}) () {
  // setBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QIntValidator9setBottomEi
    // invoke: void setBottom(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QIntValidator9setBottomEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIntValidator", "setBottom", args)
  }

  return
}

// top()
func (this *QIntValidator) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QIntValidator3topEv
    // invoke: int top()
    var ret0 = C.C_ZNK13QIntValidator3topEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIntValidator", "top", args)
  }

  return
}

// validate(class QString &, int &)
func (this *QIntValidator) Validate(args ...interface{}) () {
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
    // invoke: _ZNK13QIntValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK13QIntValidator8validateER7QStringRi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QIntValidator", "validate", args)
  }

  return
}

// fixup(class QString &)
func (this *QIntValidator) Fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QIntValidator5fixupER7QString
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QIntValidator5fixupER7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIntValidator", "fixup", args)
  }

  return
}

// setRange(int, int)
func (this *QIntValidator) Setrange(args ...interface{}) () {
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
    // invoke: _ZN13QIntValidator8setRangeEii
    // invoke: void setRange(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QIntValidator8setRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QIntValidator", "setRange", args)
  }

  return
}

// metaObject()
func (this *QValidator) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QValidator10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QValidator10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QValidator", "metaObject", args)
  }

  return
}

// setLocale(const class QLocale &)
func (this *QValidator) Setlocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QValidator9setLocaleERK7QLocale
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(*qtcore.QLocale).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QValidator9setLocaleERK7QLocale(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QValidator", "setLocale", args)
  }

  return
}

// locale()
func (this *QValidator) Locale(args ...interface{}) (ret interface{}) {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QValidator6localeEv
    // invoke: QLocale locale()
    var ret0 = C.C_ZNK10QValidator6localeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QLocale{}) // "QLocale"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QValidator", "locale", args)
  }

  return
}

// ~QValidator()
func (this *QValidator) Freeqvalidator(args ...interface{}) () {
  // ~QValidator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QValidatorD0Ev
    // invoke: void ~QValidator()
    C.C_ZN10QValidatorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QValidator", "~QValidator", args)
  }

  return
}

// QValidator(class QObject *)
func NewQValidator(args ...interface{}) *QValidator {
  // QValidator(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QValidatorC1EP7QObject
    // invoke: void QValidator(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QValidatorC2EP7QObject(arg0)
    return &QValidator{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QValidator", "QValidator", args)
  }

  return nil // QValidator{Qclsinst:qthis}
}

// fixup(class QString &)
func (this *QValidator) Fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QValidator5fixupER7QString
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK10QValidator5fixupER7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QValidator", "fixup", args)
  }

  return
}

// metaObject()
func (this *QRegExpValidator) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QRegExpValidator10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QRegExpValidator10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "metaObject", args)
  }

  return
}

// QRegExpValidator(const class QRegExp &, class QObject *)
func NewQRegExpValidator(args ...interface{}) *QRegExpValidator {
  // QRegExpValidator(const class QRegExp &, class QObject *)
  // QRegExpValidator(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRegExp{}) // "const QRegExp &"
  vtys[0][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QRegExpValidatorC1ERK7QRegExpP7QObject
    // invoke: void QRegExpValidator(const class QRegExp &, class QObject *)
    var arg0 = args[0].(*qtcore.QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject(arg0, arg1)
    return &QRegExpValidator{Qclsinst:qthis}
  case 1:
    // invoke: _ZN16QRegExpValidatorC1EP7QObject
    // invoke: void QRegExpValidator(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QRegExpValidatorC2EP7QObject(arg0)
    return &QRegExpValidator{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRegExpValidator", "QRegExpValidator", args)
  }

  return nil // QRegExpValidator{Qclsinst:qthis}
}

// setRegExp(const class QRegExp &)
func (this *QRegExpValidator) Setregexp(args ...interface{}) () {
  // setRegExp(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QRegExpValidator9setRegExpERK7QRegExp
    // invoke: void setRegExp(const class QRegExp &)
    var arg0 = args[0].(*qtcore.QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QRegExpValidator9setRegExpERK7QRegExp(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "setRegExp", args)
  }

  return
}

// regExp()
func (this *QRegExpValidator) Regexp(args ...interface{}) (ret interface{}) {
  // regExp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QRegExpValidator6regExpEv
    // invoke: const QRegExp & regExp()
    var ret0 = C.C_ZNK16QRegExpValidator6regExpEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRegExp{}) // "const QRegExp &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRegExpValidator", "regExp", args)
  }

  return
}

// validate(class QString &, int &)
func (this *QRegExpValidator) Validate(args ...interface{}) () {
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
    // invoke: _ZNK16QRegExpValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK16QRegExpValidator8validateER7QStringRi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "validate", args)
  }

  return
}

// ~QRegExpValidator()
func (this *QRegExpValidator) Freeqregexpvalidator(args ...interface{}) () {
  // ~QRegExpValidator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QRegExpValidatorD0Ev
    // invoke: void ~QRegExpValidator()
    C.C_ZN16QRegExpValidatorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "~QRegExpValidator", args)
  }

  return
}

// <= body block end

