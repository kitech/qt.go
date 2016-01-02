package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QRegularExpression QRegularExpressionValidator::regularExpression();
extern void _ZNK27QRegularExpressionValidator17regularExpressionEv(void* qthis);
  // proto:  void QRegularExpressionValidator::~QRegularExpressionValidator();
extern void _ZN27QRegularExpressionValidatorD0Ev(void* qthis);
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(const QRegularExpression & re, QObject * parent);
extern void* dector_ZN27QRegularExpressionValidatorC1ERK18QRegularExpressionP7QObject(void* arg0, void* arg1);
extern void _ZN27QRegularExpressionValidatorC1ERK18QRegularExpressionP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  const QMetaObject * QRegularExpressionValidator::metaObject();
extern void _ZNK27QRegularExpressionValidator10metaObjectEv(void* qthis);
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(const QRegularExpressionValidator & );
extern void* dector_ZN27QRegularExpressionValidatorC1ERKS_(void* arg0);
extern void _ZN27QRegularExpressionValidatorC1ERKS_(void* qthis, void* arg0);
  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(QObject * parent);
extern void* dector_ZN27QRegularExpressionValidatorC1EP7QObject(void* arg0);
extern void _ZN27QRegularExpressionValidatorC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QRegularExpressionValidator::setRegularExpression(const QRegularExpression & re);
extern void _ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression(void* qthis, void* arg0);
  // proto:  int QDoubleValidator::decimals();
extern void _ZNK16QDoubleValidator8decimalsEv(void* qthis);
  // proto:  void QDoubleValidator::~QDoubleValidator();
extern void _ZN16QDoubleValidatorD0Ev(void* qthis);
  // proto:  double QDoubleValidator::top();
extern void _ZNK16QDoubleValidator3topEv(void* qthis);
  // proto:  double QDoubleValidator::bottom();
extern void _ZNK16QDoubleValidator6bottomEv(void* qthis);
  // proto:  void QDoubleValidator::setDecimals(int );
extern void _ZN16QDoubleValidator11setDecimalsEi(void* qthis, int arg0);
  // proto:  void QDoubleValidator::QDoubleValidator(const QDoubleValidator & );
extern void* dector_ZN16QDoubleValidatorC1ERKS_(void* arg0);
extern void _ZN16QDoubleValidatorC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDoubleValidator::setBottom(double );
extern void _ZN16QDoubleValidator9setBottomEd(void* qthis, double arg0);
  // proto:  void QDoubleValidator::setRange(double bottom, double top, int decimals);
extern void _ZN16QDoubleValidator8setRangeEddi(void* qthis, double arg0, double arg1, int arg2);
  // proto:  void QDoubleValidator::QDoubleValidator(QObject * parent);
extern void* dector_ZN16QDoubleValidatorC1EP7QObject(void* arg0);
extern void _ZN16QDoubleValidatorC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QDoubleValidator::QDoubleValidator(double bottom, double top, int decimals, QObject * parent);
extern void* dector_ZN16QDoubleValidatorC1EddiP7QObject(double arg0, double arg1, int arg2, void* arg3);
extern void _ZN16QDoubleValidatorC1EddiP7QObject(void* qthis, double arg0, double arg1, int arg2, void* arg3);
  // proto:  const QMetaObject * QDoubleValidator::metaObject();
extern void _ZNK16QDoubleValidator10metaObjectEv(void* qthis);
  // proto:  void QDoubleValidator::setTop(double );
extern void _ZN16QDoubleValidator6setTopEd(void* qthis, double arg0);
  // proto:  void QIntValidator::QIntValidator(QObject * parent);
extern void* dector_ZN13QIntValidatorC1EP7QObject(void* arg0);
extern void _ZN13QIntValidatorC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QIntValidator::setBottom(int );
extern void _ZN13QIntValidator9setBottomEi(void* qthis, int arg0);
  // proto:  void QIntValidator::setRange(int bottom, int top);
extern void _ZN13QIntValidator8setRangeEii(void* qthis, int arg0, int arg1);
  // proto:  const QMetaObject * QIntValidator::metaObject();
extern void _ZNK13QIntValidator10metaObjectEv(void* qthis);
  // proto:  void QIntValidator::QIntValidator(const QIntValidator & );
extern void* dector_ZN13QIntValidatorC1ERKS_(void* arg0);
extern void _ZN13QIntValidatorC1ERKS_(void* qthis, void* arg0);
  // proto:  int QIntValidator::top();
extern void _ZNK13QIntValidator3topEv(void* qthis);
  // proto:  void QIntValidator::fixup(QString & input);
extern void _ZNK13QIntValidator5fixupER7QString(void* qthis, void* arg0);
  // proto:  void QIntValidator::~QIntValidator();
extern void _ZN13QIntValidatorD0Ev(void* qthis);
  // proto:  void QIntValidator::setTop(int );
extern void _ZN13QIntValidator6setTopEi(void* qthis, int arg0);
  // proto:  int QIntValidator::bottom();
extern void _ZNK13QIntValidator6bottomEv(void* qthis);
  // proto:  void QIntValidator::QIntValidator(int bottom, int top, QObject * parent);
extern void* dector_ZN13QIntValidatorC1EiiP7QObject(int arg0, int arg1, void* arg2);
extern void _ZN13QIntValidatorC1EiiP7QObject(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  const QMetaObject * QValidator::metaObject();
extern void _ZNK10QValidator10metaObjectEv(void* qthis);
  // proto:  void QValidator::QValidator(const QValidator & );
extern void* dector_ZN10QValidatorC1ERKS_(void* arg0);
extern void _ZN10QValidatorC1ERKS_(void* qthis, void* arg0);
  // proto:  void QValidator::setLocale(const QLocale & locale);
extern void _ZN10QValidator9setLocaleERK7QLocale(void* qthis, void* arg0);
  // proto:  void QValidator::fixup(QString & );
extern void _ZNK10QValidator5fixupER7QString(void* qthis, void* arg0);
  // proto:  void QValidator::~QValidator();
extern void _ZN10QValidatorD0Ev(void* qthis);
  // proto:  void QValidator::QValidator(QObject * parent);
extern void* dector_ZN10QValidatorC1EP7QObject(void* arg0);
extern void _ZN10QValidatorC1EP7QObject(void* qthis, void* arg0);
  // proto:  QLocale QValidator::locale();
extern void _ZNK10QValidator6localeEv(void* qthis);
  // proto:  void QRegExpValidator::~QRegExpValidator();
extern void _ZN16QRegExpValidatorD0Ev(void* qthis);
  // proto:  const QRegExp & QRegExpValidator::regExp();
extern void _ZNK16QRegExpValidator6regExpEv(void* qthis);
  // proto:  const QMetaObject * QRegExpValidator::metaObject();
extern void _ZNK16QRegExpValidator10metaObjectEv(void* qthis);
  // proto:  void QRegExpValidator::QRegExpValidator(const QRegExp & rx, QObject * parent);
extern void* dector_ZN16QRegExpValidatorC1ERK7QRegExpP7QObject(void* arg0, void* arg1);
extern void _ZN16QRegExpValidatorC1ERK7QRegExpP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  void QRegExpValidator::setRegExp(const QRegExp & rx);
extern void _ZN16QRegExpValidator9setRegExpERK7QRegExp(void* qthis, void* arg0);
  // proto:  void QRegExpValidator::QRegExpValidator(const QRegExpValidator & );
extern void* dector_ZN16QRegExpValidatorC1ERKS_(void* arg0);
extern void _ZN16QRegExpValidatorC1ERKS_(void* qthis, void* arg0);
  // proto:  void QRegExpValidator::QRegExpValidator(QObject * parent);
extern void* dector_ZN16QRegExpValidatorC1EP7QObject(void* arg0);
extern void _ZN16QRegExpValidatorC1EP7QObject(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _regularExpressionChanged QRegularExpressionValidator_regularExpressionChanged_signal;
}

// class sizeof(QDoubleValidator)=1
type QDoubleValidator struct {
  /*qbase*/ QValidator;
  qclsinst uint64 /* *mut c_void*/;
//  _topChanged QDoubleValidator_topChanged_signal;
//  _notationChanged QDoubleValidator_notationChanged_signal;
//  _bottomChanged QDoubleValidator_bottomChanged_signal;
//  _decimalsChanged QDoubleValidator_decimalsChanged_signal;
}

// class sizeof(QIntValidator)=1
type QIntValidator struct {
  /*qbase*/ QValidator;
  qclsinst uint64 /* *mut c_void*/;
//  _topChanged QIntValidator_topChanged_signal;
//  _bottomChanged QIntValidator_bottomChanged_signal;
}

// class sizeof(QValidator)=1
type QValidator struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _changed QValidator_changed_signal;
}

// class sizeof(QRegExpValidator)=1
type QRegExpValidator struct {
  /*qbase*/ QValidator;
  qclsinst uint64 /* *mut c_void*/;
//  _regExpChanged QRegExpValidator_regExpChanged_signal;
}

  // proto:  QRegularExpression QRegularExpressionValidator::regularExpression();
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
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "regularExpression", args)
  }

}

  // proto:  void QRegularExpressionValidator::~QRegularExpressionValidator();
func (this *QRegularExpressionValidator) FreeQRegularExpressionValidator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "~QRegularExpressionValidator", args)
  }

}

  // proto:  void QRegularExpressionValidator::QRegularExpressionValidator(const QRegularExpression & re, QObject * parent);
func NewQRegularExpressionValidator(args ...interface{}) QRegularExpressionValidator {
  return QRegularExpressionValidator{}
}

  // proto:  const QMetaObject * QRegularExpressionValidator::metaObject();
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
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "metaObject", args)
  }

}

  // proto:  void QRegularExpressionValidator::setRegularExpression(const QRegularExpression & re);
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
  default:
    qtrt.ErrorResolve("QRegularExpressionValidator", "setRegularExpression", args)
  }

}

  // proto:  int QDoubleValidator::decimals();
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "decimals", args)
  }

}

  // proto:  void QDoubleValidator::~QDoubleValidator();
func (this *QDoubleValidator) FreeQDoubleValidator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDoubleValidator", "~QDoubleValidator", args)
  }

}

  // proto:  double QDoubleValidator::top();
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "top", args)
  }

}

  // proto:  double QDoubleValidator::bottom();
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "bottom", args)
  }

}

  // proto:  void QDoubleValidator::setDecimals(int );
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setDecimals", args)
  }

}

  // proto:  void QDoubleValidator::QDoubleValidator(const QDoubleValidator & );
func NewQDoubleValidator(args ...interface{}) QDoubleValidator {
  return QDoubleValidator{}
}

  // proto:  void QDoubleValidator::setBottom(double );
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setBottom", args)
  }

}

  // proto:  void QDoubleValidator::setRange(double bottom, double top, int decimals);
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setRange", args)
  }

}

  // proto:  const QMetaObject * QDoubleValidator::metaObject();
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "metaObject", args)
  }

}

  // proto:  void QDoubleValidator::setTop(double );
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
  default:
    qtrt.ErrorResolve("QDoubleValidator", "setTop", args)
  }

}

  // proto:  void QIntValidator::QIntValidator(QObject * parent);
func NewQIntValidator(args ...interface{}) QIntValidator {
  return QIntValidator{}
}

  // proto:  void QIntValidator::setBottom(int );
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
  default:
    qtrt.ErrorResolve("QIntValidator", "setBottom", args)
  }

}

  // proto:  void QIntValidator::setRange(int bottom, int top);
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
  default:
    qtrt.ErrorResolve("QIntValidator", "setRange", args)
  }

}

  // proto:  const QMetaObject * QIntValidator::metaObject();
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
  default:
    qtrt.ErrorResolve("QIntValidator", "metaObject", args)
  }

}

  // proto:  int QIntValidator::top();
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
  default:
    qtrt.ErrorResolve("QIntValidator", "top", args)
  }

}

  // proto:  void QIntValidator::fixup(QString & input);
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
  default:
    qtrt.ErrorResolve("QIntValidator", "fixup", args)
  }

}

  // proto:  void QIntValidator::~QIntValidator();
func (this *QIntValidator) FreeQIntValidator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIntValidator", "~QIntValidator", args)
  }

}

  // proto:  void QIntValidator::setTop(int );
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
  default:
    qtrt.ErrorResolve("QIntValidator", "setTop", args)
  }

}

  // proto:  int QIntValidator::bottom();
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
  default:
    qtrt.ErrorResolve("QIntValidator", "bottom", args)
  }

}

  // proto:  const QMetaObject * QValidator::metaObject();
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
  default:
    qtrt.ErrorResolve("QValidator", "metaObject", args)
  }

}

  // proto:  void QValidator::QValidator(const QValidator & );
func NewQValidator(args ...interface{}) QValidator {
  return QValidator{}
}

  // proto:  void QValidator::setLocale(const QLocale & locale);
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
  default:
    qtrt.ErrorResolve("QValidator", "setLocale", args)
  }

}

  // proto:  void QValidator::fixup(QString & );
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
  default:
    qtrt.ErrorResolve("QValidator", "fixup", args)
  }

}

  // proto:  void QValidator::~QValidator();
func (this *QValidator) FreeQValidator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QValidator", "~QValidator", args)
  }

}

  // proto:  QLocale QValidator::locale();
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
  default:
    qtrt.ErrorResolve("QValidator", "locale", args)
  }

}

  // proto:  void QRegExpValidator::~QRegExpValidator();
func (this *QRegExpValidator) FreeQRegExpValidator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegExpValidator", "~QRegExpValidator", args)
  }

}

  // proto:  const QRegExp & QRegExpValidator::regExp();
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
  default:
    qtrt.ErrorResolve("QRegExpValidator", "regExp", args)
  }

}

  // proto:  const QMetaObject * QRegExpValidator::metaObject();
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
  default:
    qtrt.ErrorResolve("QRegExpValidator", "metaObject", args)
  }

}

  // proto:  void QRegExpValidator::QRegExpValidator(const QRegExp & rx, QObject * parent);
func NewQRegExpValidator(args ...interface{}) QRegExpValidator {
  return QRegExpValidator{}
}

  // proto:  void QRegExpValidator::setRegExp(const QRegExp & rx);
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
  default:
    qtrt.ErrorResolve("QRegExpValidator", "setRegExp", args)
  }

}

// <= body block end

