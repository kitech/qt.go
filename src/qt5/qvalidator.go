package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qvalidator.h
// dst-file: /src/gui/qvalidator.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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


func NewQRegularExpressionValidator(args ...interface{})() {
}


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


func NewQDoubleValidator(args ...interface{})() {
}


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


func NewQIntValidator(args ...interface{})() {
}


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


func NewQValidator(args ...interface{})() {
}


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


func NewQRegExpValidator(args ...interface{})() {
}


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

