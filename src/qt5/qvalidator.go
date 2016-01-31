package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK27QRegularExpressionValidator17regularExpressionEv(void* qthis); // 4
  // proto:  void QRegularExpressionValidator::~QRegularExpressionValidator();
extern void C_ZN27QRegularExpressionValidatorD2Ev(void* qthis); // 4
  // proto:  void QRegularExpressionValidator::setRegularExpression(const QRegularExpression & re);
extern void C_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression(void* qthis, void* arg0); // 4
  // proto:  QValidator::State QRegularExpressionValidator::validate(QString & input, int & pos);
extern void C_ZNK27QRegularExpressionValidator8validateER7QStringRi(void* qthis, void* arg0, int32_t* arg1); // 4
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(QObject * parent);
extern void C_ZN27QRegularExpressionValidatorC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(const QRegularExpression & re, QObject * parent);
extern void C_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QDoubleValidator::setBottom(double );
extern void C_ZN16QDoubleValidator9setBottomEd(void* qthis, double arg0); // 4
  // proto:  void QDoubleValidator::setTop(double );
extern void C_ZN16QDoubleValidator6setTopEd(void* qthis, double arg0); // 4
  // proto:  double QDoubleValidator::bottom();
extern void C_ZNK16QDoubleValidator6bottomEv(void* qthis); // 2
  // proto:  void QDoubleValidator::setDecimals(int );
extern void C_ZN16QDoubleValidator11setDecimalsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDoubleValidator::~QDoubleValidator();
extern void C_ZN16QDoubleValidatorD2Ev(void* qthis); // 4
  // proto:  QDoubleValidator::Notation QDoubleValidator::notation();
extern void C_ZNK16QDoubleValidator8notationEv(void* qthis); // 4
  // proto:  QValidator::State QDoubleValidator::validate(QString & , int & );
extern void C_ZNK16QDoubleValidator8validateER7QStringRi(void* qthis, void* arg0, int32_t* arg1); // 4
  // proto:  void QDoubleValidator::setRange(double bottom, double top, int decimals);
extern void C_ZN16QDoubleValidator8setRangeEddi(void* qthis, double arg0, double arg1, int32_t arg2); // 4
  // proto:  const QMetaObject * QDoubleValidator::metaObject();
extern void C_ZNK16QDoubleValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QDoubleValidator::QDoubleValidator(QObject * parent);
extern void C_ZN16QDoubleValidatorC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QDoubleValidator::QDoubleValidator(double bottom, double top, int decimals, QObject * parent);
extern void C_ZN16QDoubleValidatorC2EddiP7QObject(void* qthis, double arg0, double arg1, int32_t arg2, void* arg3); // 3
  // proto:  int QDoubleValidator::decimals();
extern void C_ZNK16QDoubleValidator8decimalsEv(void* qthis); // 2
  // proto:  double QDoubleValidator::top();
extern void C_ZNK16QDoubleValidator3topEv(void* qthis); // 2
  // proto:  void QIntValidator::setTop(int );
extern void C_ZN13QIntValidator6setTopEi(void* qthis, int32_t arg0); // 4
  // proto:  void QIntValidator::QIntValidator(QObject * parent);
extern void C_ZN13QIntValidatorC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QIntValidator::QIntValidator(int bottom, int top, QObject * parent);
extern void C_ZN13QIntValidatorC2EiiP7QObject(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 3
  // proto:  const QMetaObject * QIntValidator::metaObject();
extern void C_ZNK13QIntValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QIntValidator::~QIntValidator();
extern void C_ZN13QIntValidatorD2Ev(void* qthis); // 4
  // proto:  int QIntValidator::bottom();
extern void C_ZNK13QIntValidator6bottomEv(void* qthis); // 2
  // proto:  void QIntValidator::setBottom(int );
extern void C_ZN13QIntValidator9setBottomEi(void* qthis, int32_t arg0); // 4
  // proto:  int QIntValidator::top();
extern void C_ZNK13QIntValidator3topEv(void* qthis); // 2
  // proto:  QValidator::State QIntValidator::validate(QString & , int & );
extern void C_ZNK13QIntValidator8validateER7QStringRi(void* qthis, void* arg0, int32_t* arg1); // 4
  // proto:  void QIntValidator::fixup(QString & input);
extern void C_ZNK13QIntValidator5fixupER7QString(void* qthis, void* arg0); // 4
  // proto:  void QIntValidator::setRange(int bottom, int top);
extern void C_ZN13QIntValidator8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QValidator::metaObject();
extern void C_ZNK10QValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QValidator::setLocale(const QLocale & locale);
extern void C_ZN10QValidator9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QLocale QValidator::locale();
extern void C_ZNK10QValidator6localeEv(void* qthis); // 4
  // proto:  void QValidator::~QValidator();
extern void C_ZN10QValidatorD2Ev(void* qthis); // 4
  // proto:  void QValidator::QValidator(QObject * parent);
extern void C_ZN10QValidatorC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QValidator::fixup(QString & );
extern void C_ZNK10QValidator5fixupER7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QRegExpValidator::metaObject();
extern void C_ZNK16QRegExpValidator10metaObjectEv(void* qthis); // 4
  // proto:  void QRegExpValidator::QRegExpValidator(const QRegExp & rx, QObject * parent);
extern void C_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QRegExpValidator::QRegExpValidator(QObject * parent);
extern void C_ZN16QRegExpValidatorC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QRegExpValidator::setRegExp(const QRegExp & rx);
extern void C_ZN16QRegExpValidator9setRegExpERK7QRegExp(void* qthis, void* arg0); // 4
  // proto:  const QRegExp & QRegExpValidator::regExp();
extern void C_ZNK16QRegExpValidator6regExpEv(void* qthis); // 2
  // proto:  QValidator::State QRegExpValidator::validate(QString & input, int & pos);
extern void C_ZNK16QRegExpValidator8validateER7QStringRi(void* qthis, void* arg0, int32_t* arg1); // 4
  // proto:  void QRegExpValidator::~QRegExpValidator();
extern void C_ZN16QRegExpValidatorD2Ev(void* qthis); // 4
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

// class sizeof(QRegularExpressionValidator)=1
type QRegularExpressionValidator struct {
  /*qbase*/ QValidator;
  qclsinst unsafe.Pointer /* *C.void */;
//  _regularExpressionChanged QRegularExpressionValidator_regularExpressionChanged_signal;
}

// class sizeof(QDoubleValidator)=1
type QDoubleValidator struct {
  /*qbase*/ QValidator;
  qclsinst unsafe.Pointer /* *C.void */;
//  _topChanged QDoubleValidator_topChanged_signal;
//  _notationChanged QDoubleValidator_notationChanged_signal;
//  _bottomChanged QDoubleValidator_bottomChanged_signal;
//  _decimalsChanged QDoubleValidator_decimalsChanged_signal;
}

// class sizeof(QIntValidator)=1
type QIntValidator struct {
  /*qbase*/ QValidator;
  qclsinst unsafe.Pointer /* *C.void */;
//  _topChanged QIntValidator_topChanged_signal;
//  _bottomChanged QIntValidator_bottomChanged_signal;
}

// class sizeof(QValidator)=1
type QValidator struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _changed QValidator_changed_signal;
}

// class sizeof(QRegExpValidator)=1
type QRegExpValidator struct {
  /*qbase*/ QValidator;
  qclsinst unsafe.Pointer /* *C.void */;
//  _regExpChanged QRegExpValidator_regExpChanged_signal;
}

// metaObject()
func (this *QRegularExpressionValidator) metaObject(args ...interface{}) () {
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
    C.C_ZNK27QRegularExpressionValidator10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "metaObject", args)
  }

}

// regularExpression()
func (this *QRegularExpressionValidator) regularExpression(args ...interface{}) () {
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
    C.C_ZNK27QRegularExpressionValidator17regularExpressionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "regularExpression", args)
  }

}

// ~QRegularExpressionValidator()
func (this *QRegularExpressionValidator) FreeQRegularExpressionValidator(args ...interface{}) () {
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
    C.C_ZN27QRegularExpressionValidatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "~QRegularExpressionValidator", args)
  }

}

// setRegularExpression(const class QRegularExpression &)
func (this *QRegularExpressionValidator) setRegularExpression(args ...interface{}) () {
  // setRegularExpression(const class QRegularExpression &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression
    // invoke: void setRegularExpression(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "setRegularExpression", args)
  }

}

// validate(class QString &, int &)
func (this *QRegularExpressionValidator) validate(args ...interface{}) () {
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
    // invoke: _ZNK27QRegularExpressionValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK27QRegularExpressionValidator8validateER7QStringRi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "validate", args)
  }

}

// QRegularExpressionValidator(class QObject *)
func NewQRegularExpressionValidator(args ...interface{}) QRegularExpressionValidator {
  // QRegularExpressionValidator(class QObject *)
  // QRegularExpressionValidator(const class QRegularExpression &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QRegularExpressionValidatorC1EP7QObject
    // invoke: void QRegularExpressionValidator(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN27QRegularExpressionValidatorC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN27QRegularExpressionValidatorC1ERK18QRegularExpressionP7QObject
    // invoke: void QRegularExpressionValidator(const class QRegularExpression &, class QObject *)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "QRegularExpressionValidator", args)
  }

  return QRegularExpressionValidator{}
}

// setBottom(double)
func (this *QDoubleValidator) setBottom(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDoubleValidator9setBottomEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setBottom", args)
  }

}

// setTop(double)
func (this *QDoubleValidator) setTop(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDoubleValidator6setTopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setTop", args)
  }

}

// bottom()
func (this *QDoubleValidator) bottom(args ...interface{}) () {
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
    C.C_ZNK16QDoubleValidator6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "bottom", args)
  }

}

// setDecimals(int)
func (this *QDoubleValidator) setDecimals(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QDoubleValidator11setDecimalsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setDecimals", args)
  }

}

// ~QDoubleValidator()
func (this *QDoubleValidator) FreeQDoubleValidator(args ...interface{}) () {
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
    C.C_ZN16QDoubleValidatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "~QDoubleValidator", args)
  }

}

// notation()
func (this *QDoubleValidator) notation(args ...interface{}) () {
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
    C.C_ZNK16QDoubleValidator8notationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "notation", args)
  }

}

// validate(class QString &, int &)
func (this *QDoubleValidator) validate(args ...interface{}) () {
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
    // invoke: _ZNK16QDoubleValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK16QDoubleValidator8validateER7QStringRi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "validate", args)
  }

}

// setRange(double, double, int)
func (this *QDoubleValidator) setRange(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN16QDoubleValidator8setRangeEddi(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setRange", args)
  }

}

// metaObject()
func (this *QDoubleValidator) metaObject(args ...interface{}) () {
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
    C.C_ZNK16QDoubleValidator10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "metaObject", args)
  }

}

// QDoubleValidator(class QObject *)
func NewQDoubleValidator(args ...interface{}) QDoubleValidator {
  // QDoubleValidator(class QObject *)
  // QDoubleValidator(double, double, int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"
  vtys[1][1] = qtrt.DoubleTy(false) // "double"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDoubleValidatorC1EP7QObject
    // invoke: void QDoubleValidator(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QDoubleValidatorC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN16QDoubleValidatorC1EddiP7QObject
    // invoke: void QDoubleValidator(double, double, int, class QObject *)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QObject).qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QDoubleValidatorC2EddiP7QObject(qthis, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "QDoubleValidator", args)
  }

  return QDoubleValidator{}
}

// decimals()
func (this *QDoubleValidator) decimals(args ...interface{}) () {
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
    C.C_ZNK16QDoubleValidator8decimalsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "decimals", args)
  }

}

// top()
func (this *QDoubleValidator) top(args ...interface{}) () {
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
    C.C_ZNK16QDoubleValidator3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDoubleValidator", "top", args)
  }

}

// setTop(int)
func (this *QIntValidator) setTop(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QIntValidator6setTopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIntValidator", "setTop", args)
  }

}

// QIntValidator(class QObject *)
func NewQIntValidator(args ...interface{}) QIntValidator {
  // QIntValidator(class QObject *)
  // QIntValidator(int, int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QIntValidatorC1EP7QObject
    // invoke: void QIntValidator(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QIntValidatorC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN13QIntValidatorC1EiiP7QObject
    // invoke: void QIntValidator(int, int, class QObject *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QIntValidatorC2EiiP7QObject(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QIntValidator", "QIntValidator", args)
  }

  return QIntValidator{}
}

// metaObject()
func (this *QIntValidator) metaObject(args ...interface{}) () {
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
    C.C_ZNK13QIntValidator10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIntValidator", "metaObject", args)
  }

}

// ~QIntValidator()
func (this *QIntValidator) FreeQIntValidator(args ...interface{}) () {
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
    C.C_ZN13QIntValidatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIntValidator", "~QIntValidator", args)
  }

}

// bottom()
func (this *QIntValidator) bottom(args ...interface{}) () {
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
    C.C_ZNK13QIntValidator6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIntValidator", "bottom", args)
  }

}

// setBottom(int)
func (this *QIntValidator) setBottom(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QIntValidator9setBottomEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIntValidator", "setBottom", args)
  }

}

// top()
func (this *QIntValidator) top(args ...interface{}) () {
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
    C.C_ZNK13QIntValidator3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIntValidator", "top", args)
  }

}

// validate(class QString &, int &)
func (this *QIntValidator) validate(args ...interface{}) () {
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
    // invoke: _ZNK13QIntValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK13QIntValidator8validateER7QStringRi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QIntValidator", "validate", args)
  }

}

// fixup(class QString &)
func (this *QIntValidator) fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QIntValidator5fixupER7QString
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QIntValidator5fixupER7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIntValidator", "fixup", args)
  }

}

// setRange(int, int)
func (this *QIntValidator) setRange(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QIntValidator8setRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QIntValidator", "setRange", args)
  }

}

// metaObject()
func (this *QValidator) metaObject(args ...interface{}) () {
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
    C.C_ZNK10QValidator10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QValidator", "metaObject", args)
  }

}

// setLocale(const class QLocale &)
func (this *QValidator) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QValidator9setLocaleERK7QLocale
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QValidator9setLocaleERK7QLocale(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QValidator", "setLocale", args)
  }

}

// locale()
func (this *QValidator) locale(args ...interface{}) () {
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
    C.C_ZNK10QValidator6localeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QValidator", "locale", args)
  }

}

// ~QValidator()
func (this *QValidator) FreeQValidator(args ...interface{}) () {
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
    C.C_ZN10QValidatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QValidator", "~QValidator", args)
  }

}

// QValidator(class QObject *)
func NewQValidator(args ...interface{}) QValidator {
  // QValidator(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QValidatorC1EP7QObject
    // invoke: void QValidator(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QValidatorC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QValidator", "QValidator", args)
  }

  return QValidator{}
}

// fixup(class QString &)
func (this *QValidator) fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QValidator5fixupER7QString
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK10QValidator5fixupER7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QValidator", "fixup", args)
  }

}

// metaObject()
func (this *QRegExpValidator) metaObject(args ...interface{}) () {
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
    C.C_ZNK16QRegExpValidator10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "metaObject", args)
  }

}

// QRegExpValidator(const class QRegExp &, class QObject *)
func NewQRegExpValidator(args ...interface{}) QRegExpValidator {
  // QRegExpValidator(const class QRegExp &, class QObject *)
  // QRegExpValidator(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QRegExpValidatorC1ERK7QRegExpP7QObject
    // invoke: void QRegExpValidator(const class QRegExp &, class QObject *)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN16QRegExpValidatorC1EP7QObject
    // invoke: void QRegExpValidator(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QRegExpValidatorC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "QRegExpValidator", args)
  }

  return QRegExpValidator{}
}

// setRegExp(const class QRegExp &)
func (this *QRegExpValidator) setRegExp(args ...interface{}) () {
  // setRegExp(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QRegExpValidator9setRegExpERK7QRegExp
    // invoke: void setRegExp(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QRegExpValidator9setRegExpERK7QRegExp(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "setRegExp", args)
  }

}

// regExp()
func (this *QRegExpValidator) regExp(args ...interface{}) () {
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
    C.C_ZNK16QRegExpValidator6regExpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "regExp", args)
  }

}

// validate(class QString &, int &)
func (this *QRegExpValidator) validate(args ...interface{}) () {
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
    // invoke: _ZNK16QRegExpValidator8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK16QRegExpValidator8validateER7QStringRi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "validate", args)
  }

}

// ~QRegExpValidator()
func (this *QRegExpValidator) FreeQRegExpValidator(args ...interface{}) () {
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
    C.C_ZN16QRegExpValidatorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExpValidator", "~QRegExpValidator", args)
  }

}

// <= body block end

