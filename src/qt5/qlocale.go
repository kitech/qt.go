package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qlocale.h
// dst-file: /src/core/qlocale.go
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
// class sizeof(QLocale)=1
type QLocale struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QLocale) pmText(args ...interface{}) () {
  // pmText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6pmTextEv
  default:
    qtrt.ErrorResolve("QLocale", "pmText", args)
 }

}


func (this *QLocale) nativeLanguageName(args ...interface{}) () {
  // nativeLanguageName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale18nativeLanguageNameEv
  default:
    qtrt.ErrorResolve("QLocale", "nativeLanguageName", args)
 }

}


func (this *QLocale) toLower(args ...interface{}) () {
  // toLower(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7toLowerERK7QString
  default:
    qtrt.ErrorResolve("QLocale", "toLower", args)
 }

}


func (this *QLocale) zeroDigit(args ...interface{}) () {
  // zeroDigit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale9zeroDigitEv
  default:
    qtrt.ErrorResolve("QLocale", "zeroDigit", args)
 }

}


func (this *QLocale) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale4nameEv
  default:
    qtrt.ErrorResolve("QLocale", "name", args)
 }

}


func (this *QLocale) toCurrencyString(args ...interface{}) () {
  // toCurrencyString(qlonglong, const class QString &)
  // toCurrencyString(uint, const class QString &)
  // toCurrencyString(short, const class QString &)
  // toCurrencyString(double, const class QString &)
  // toCurrencyString(qulonglong, const class QString &)
  // toCurrencyString(ushort, const class QString &)
  // toCurrencyString(float, const class QString &)
  // toCurrencyString(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int16Ty(false) // "short"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "double"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[5][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.FloatTy(false) // "float"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale16toCurrencyStringExRK7QString
  case 1:
    // invoke: _ZNK7QLocale16toCurrencyStringEjRK7QString
  case 2:
    // invoke: _ZNK7QLocale16toCurrencyStringEsRK7QString
  case 3:
    // invoke: _ZNK7QLocale16toCurrencyStringEdRK7QString
  case 4:
    // invoke: _ZNK7QLocale16toCurrencyStringEyRK7QString
  case 5:
    // invoke: _ZNK7QLocale16toCurrencyStringEtRK7QString
  case 6:
    // invoke: _ZNK7QLocale16toCurrencyStringEfRK7QString
  case 7:
    // invoke: _ZNK7QLocale16toCurrencyStringEiRK7QString
  default:
    qtrt.ErrorResolve("QLocale", "toCurrencyString", args)
 }

}


func (this *QLocale) toFloat(args ...interface{}) () {
  // toFloat(const class QString &, _Bool *)
  // toFloat(const class QStringRef &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7toFloatERK7QStringPb
  case 1:
    // invoke: _ZNK7QLocale7toFloatERK10QStringRefPb
  default:
    qtrt.ErrorResolve("QLocale", "toFloat", args)
 }

}


func (this *QLocale) c_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLocale", "c", args)
 }

}


func (this *QLocale) createSeparatedList(args ...interface{}) () {
  // createSeparatedList(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale19createSeparatedListERK11QStringList
  default:
    qtrt.ErrorResolve("QLocale", "createSeparatedList", args)
 }

}


func (this *QLocale) toUInt(args ...interface{}) () {
  // toUInt(const class QString &, _Bool *)
  // toUInt(const class QStringRef &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6toUIntERK7QStringPb
  case 1:
    // invoke: _ZNK7QLocale6toUIntERK10QStringRefPb
  default:
    qtrt.ErrorResolve("QLocale", "toUInt", args)
 }

}


func (this *QLocale) decimalPoint(args ...interface{}) () {
  // decimalPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale12decimalPointEv
  default:
    qtrt.ErrorResolve("QLocale", "decimalPoint", args)
 }

}


func (this *QLocale) positiveSign(args ...interface{}) () {
  // positiveSign()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale12positiveSignEv
  default:
    qtrt.ErrorResolve("QLocale", "positiveSign", args)
 }

}


func (this *QLocale) toLongLong(args ...interface{}) () {
  // toLongLong(const class QString &, _Bool *)
  // toLongLong(const class QStringRef &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale10toLongLongERK7QStringPb
  case 1:
    // invoke: _ZNK7QLocale10toLongLongERK10QStringRefPb
  default:
    qtrt.ErrorResolve("QLocale", "toLongLong", args)
 }

}


func (this *QLocale) toShort(args ...interface{}) () {
  // toShort(const class QStringRef &, _Bool *)
  // toShort(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7toShortERK10QStringRefPb
  case 1:
    // invoke: _ZNK7QLocale7toShortERK7QStringPb
  default:
    qtrt.ErrorResolve("QLocale", "toShort", args)
 }

}


func (this *QLocale) toString(args ...interface{}) () {
  // toString(const class QTime &, enum QLocale::FormatType)
  // toString(float, char, int)
  // toString(const class QDateTime &, enum QLocale::FormatType)
  // toString(const class QDateTime &, const class QString &)
  // toString(const class QTime &, const class QString &)
  // toString(const class QDate &, const class QString &)
  // toString(double, char, int)
  // toString(int)
  // toString(uint)
  // toString(const class QDate &, enum QLocale::FormatType)
  // toString(qlonglong)
  // toString(qulonglong)
  // toString(ushort)
  // toString(short)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[0][1] = qtrt.Int32Ty(false) // "QLocale::FormatType"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.ByteTy(false) // "char"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[2][1] = qtrt.Int32Ty(false) // "QLocale::FormatType"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[5][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.DoubleTy(false) // "double"
  vtys[6][1] = qtrt.ByteTy(false) // "char"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "uint"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[9][1] = qtrt.Int32Ty(false) // "QLocale::FormatType"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.Int16Ty(false) // "short"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8toStringERK5QTimeNS_10FormatTypeE
  case 1:
    // invoke: _ZNK7QLocale8toStringEfci
  case 2:
    // invoke: _ZNK7QLocale8toStringERK9QDateTimeNS_10FormatTypeE
  case 3:
    // invoke: _ZNK7QLocale8toStringERK9QDateTimeRK7QString
  case 4:
    // invoke: _ZNK7QLocale8toStringERK5QTimeRK7QString
  case 5:
    // invoke: _ZNK7QLocale8toStringERK5QDateRK7QString
  case 6:
    // invoke: _ZNK7QLocale8toStringEdci
  case 7:
    // invoke: _ZNK7QLocale8toStringEi
  case 8:
    // invoke: _ZNK7QLocale8toStringEj
  case 9:
    // invoke: _ZNK7QLocale8toStringERK5QDateNS_10FormatTypeE
  case 10:
    // invoke: _ZNK7QLocale8toStringEx
  case 11:
    // invoke: _ZNK7QLocale8toStringEy
  case 12:
    // invoke: _ZNK7QLocale8toStringEt
  case 13:
    // invoke: _ZNK7QLocale8toStringEs
  default:
    qtrt.ErrorResolve("QLocale", "toString", args)
 }

}


func (this *QLocale) toDateTime(args ...interface{}) () {
  // toDateTime(const class QString &, const class QString &)
  // toDateTime(const class QString &, enum QLocale::FormatType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "QLocale::FormatType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale10toDateTimeERK7QStringS2_
  case 1:
    // invoke: _ZNK7QLocale10toDateTimeERK7QStringNS_10FormatTypeE
  default:
    qtrt.ErrorResolve("QLocale", "toDateTime", args)
 }

}


func (this *QLocale) groupSeparator(args ...interface{}) () {
  // groupSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale14groupSeparatorEv
  default:
    qtrt.ErrorResolve("QLocale", "groupSeparator", args)
 }

}


func NewQLocale(args ...interface{})() {
}


func (this *QLocale) toDate(args ...interface{}) () {
  // toDate(const class QString &, enum QLocale::FormatType)
  // toDate(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "QLocale::FormatType"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6toDateERK7QStringNS_10FormatTypeE
  case 1:
    // invoke: _ZNK7QLocale6toDateERK7QStringS2_
  default:
    qtrt.ErrorResolve("QLocale", "toDate", args)
 }

}


func (this *QLocale) nativeCountryName(args ...interface{}) () {
  // nativeCountryName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale17nativeCountryNameEv
  default:
    qtrt.ErrorResolve("QLocale", "nativeCountryName", args)
 }

}


func (this *QLocale) negativeSign(args ...interface{}) () {
  // negativeSign()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale12negativeSignEv
  default:
    qtrt.ErrorResolve("QLocale", "negativeSign", args)
 }

}


func (this *QLocale) FreeQLocale(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLocale", "~QLocale", args)
 }

}


func (this *QLocale) toUpper(args ...interface{}) () {
  // toUpper(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7toUpperERK7QString
  default:
    qtrt.ErrorResolve("QLocale", "toUpper", args)
 }

}


func (this *QLocale) percent(args ...interface{}) () {
  // percent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7percentEv
  default:
    qtrt.ErrorResolve("QLocale", "percent", args)
 }

}


func (this *QLocale) toULongLong(args ...interface{}) () {
  // toULongLong(const class QStringRef &, _Bool *)
  // toULongLong(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale11toULongLongERK10QStringRefPb
  case 1:
    // invoke: _ZNK7QLocale11toULongLongERK7QStringPb
  default:
    qtrt.ErrorResolve("QLocale", "toULongLong", args)
 }

}


func (this *QLocale) uiLanguages(args ...interface{}) () {
  // uiLanguages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale11uiLanguagesEv
  default:
    qtrt.ErrorResolve("QLocale", "uiLanguages", args)
 }

}


func (this *QLocale) bcp47Name(args ...interface{}) () {
  // bcp47Name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale9bcp47NameEv
  default:
    qtrt.ErrorResolve("QLocale", "bcp47Name", args)
 }

}


func (this *QLocale) toTime(args ...interface{}) () {
  // toTime(const class QString &, enum QLocale::FormatType)
  // toTime(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "QLocale::FormatType"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6toTimeERK7QStringNS_10FormatTypeE
  case 1:
    // invoke: _ZNK7QLocale6toTimeERK7QStringS2_
  default:
    qtrt.ErrorResolve("QLocale", "toTime", args)
 }

}


func (this *QLocale) toUShort(args ...interface{}) () {
  // toUShort(const class QStringRef &, _Bool *)
  // toUShort(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8toUShortERK10QStringRefPb
  case 1:
    // invoke: _ZNK7QLocale8toUShortERK7QStringPb
  default:
    qtrt.ErrorResolve("QLocale", "toUShort", args)
 }

}


func (this *QLocale) toDouble(args ...interface{}) () {
  // toDouble(const class QStringRef &, _Bool *)
  // toDouble(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8toDoubleERK10QStringRefPb
  case 1:
    // invoke: _ZNK7QLocale8toDoubleERK7QStringPb
  default:
    qtrt.ErrorResolve("QLocale", "toDouble", args)
 }

}


func (this *QLocale) system_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLocale", "system", args)
 }

}


func (this *QLocale) setDefault_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLocale", "setDefault", args)
 }

}


func (this *QLocale) exponential(args ...interface{}) () {
  // exponential()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale11exponentialEv
  default:
    qtrt.ErrorResolve("QLocale", "exponential", args)
 }

}


func (this *QLocale) amText(args ...interface{}) () {
  // amText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6amTextEv
  default:
    qtrt.ErrorResolve("QLocale", "amText", args)
 }

}


func (this *QLocale) toInt(args ...interface{}) () {
  // toInt(const class QStringRef &, _Bool *)
  // toInt(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale5toIntERK10QStringRefPb
  case 1:
    // invoke: _ZNK7QLocale5toIntERK7QStringPb
  default:
    qtrt.ErrorResolve("QLocale", "toInt", args)
 }

}

// <= body block end

