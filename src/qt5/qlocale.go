package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qlocale.h
// dst-file: /src/core/qlocale.go
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
  // proto:  QString QLocale::pmText();
extern void _ZNK7QLocale6pmTextEv(void* qthis);
  // proto:  QString QLocale::nativeLanguageName();
extern void _ZNK7QLocale18nativeLanguageNameEv(void* qthis);
  // proto:  QString QLocale::toLower(const QString & str);
extern void _ZNK7QLocale7toLowerERK7QString(void* qthis, void* arg0);
  // proto:  QChar QLocale::zeroDigit();
extern void _ZNK7QLocale9zeroDigitEv(void* qthis);
  // proto:  QString QLocale::name();
extern void _ZNK7QLocale4nameEv(void* qthis);
  // proto:  QString QLocale::toCurrencyString(qlonglong , const QString & symbol);
extern void _ZNK7QLocale16toCurrencyStringExRK7QString(void* qthis, int64_t arg0, void* arg1);
  // proto:  float QLocale::toFloat(const QString & s, bool * ok);
extern void _ZNK7QLocale7toFloatERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto: static QLocale QLocale::c();
extern void demth_ZN7QLocale1cEv();
  // proto:  QString QLocale::toCurrencyString(uint , const QString & symbol);
extern void demth_ZNK7QLocale16toCurrencyStringEjRK7QString(void* qthis, int32_t arg0, void* arg1);
  // proto:  QString QLocale::createSeparatedList(const QStringList & strl);
extern void _ZNK7QLocale19createSeparatedListERK11QStringList(void* qthis, void* arg0);
  // proto:  uint QLocale::toUInt(const QString & s, bool * ok);
extern void _ZNK7QLocale6toUIntERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto:  QChar QLocale::decimalPoint();
extern void _ZNK7QLocale12decimalPointEv(void* qthis);
  // proto:  QChar QLocale::positiveSign();
extern void _ZNK7QLocale12positiveSignEv(void* qthis);
  // proto:  qlonglong QLocale::toLongLong(const QString & s, bool * ok);
extern void _ZNK7QLocale10toLongLongERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto:  short QLocale::toShort(const QString & s, bool * ok);
extern void _ZNK7QLocale7toShortERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto:  QString QLocale::toString(float i, char f, int prec);
extern void demth_ZNK7QLocale8toStringEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2);
  // proto:  QString QLocale::toString(const QDateTime & dateTime, const QString & format);
extern void _ZNK7QLocale8toStringERK9QDateTimeRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QDateTime QLocale::toDateTime(const QString & string, const QString & format);
extern void _ZNK7QLocale10toDateTimeERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QString QLocale::toCurrencyString(short , const QString & symbol);
extern void demth_ZNK7QLocale16toCurrencyStringEsRK7QString(void* qthis, int16_t arg0, void* arg1);
  // proto:  QChar QLocale::groupSeparator();
extern void _ZNK7QLocale14groupSeparatorEv(void* qthis);
  // proto:  QString QLocale::toCurrencyString(double , const QString & symbol);
extern void _ZNK7QLocale16toCurrencyStringEdRK7QString(void* qthis, double arg0, void* arg1);
  // proto:  QString QLocale::toCurrencyString(qulonglong , const QString & symbol);
extern void _ZNK7QLocale16toCurrencyStringEyRK7QString(void* qthis, int64_t arg0, void* arg1);
  // proto:  void QLocale::QLocale(const QString & name);
extern void* dector_ZN7QLocaleC1ERK7QString(void* arg0);
extern void _ZN7QLocaleC1ERK7QString(void* qthis, void* arg0);
  // proto:  QString QLocale::toString(const QTime & time, const QString & formatStr);
extern void _ZNK7QLocale8toStringERK5QTimeRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QDate QLocale::toDate(const QString & string, const QString & format);
extern void _ZNK7QLocale6toDateERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QString QLocale::nativeCountryName();
extern void _ZNK7QLocale17nativeCountryNameEv(void* qthis);
  // proto:  QChar QLocale::negativeSign();
extern void _ZNK7QLocale12negativeSignEv(void* qthis);
  // proto:  void QLocale::~QLocale();
extern void _ZN7QLocaleD0Ev(void* qthis);
  // proto:  void QLocale::QLocale(const QLocale & other);
extern void* dector_ZN7QLocaleC1ERKS_(void* arg0);
extern void _ZN7QLocaleC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QLocale::toString(const QDate & date, const QString & formatStr);
extern void _ZNK7QLocale8toStringERK5QDateRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QString QLocale::toUpper(const QString & str);
extern void _ZNK7QLocale7toUpperERK7QString(void* qthis, void* arg0);
  // proto:  QChar QLocale::percent();
extern void _ZNK7QLocale7percentEv(void* qthis);
  // proto:  qulonglong QLocale::toULongLong(const QString & s, bool * ok);
extern void _ZNK7QLocale11toULongLongERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto:  QString QLocale::toString(double i, char f, int prec);
extern void _ZNK7QLocale8toStringEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2);
  // proto:  QStringList QLocale::uiLanguages();
extern void _ZNK7QLocale11uiLanguagesEv(void* qthis);
  // proto:  QString QLocale::bcp47Name();
extern void _ZNK7QLocale9bcp47NameEv(void* qthis);
  // proto:  QTime QLocale::toTime(const QString & string, const QString & format);
extern void _ZNK7QLocale6toTimeERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  ushort QLocale::toUShort(const QString & s, bool * ok);
extern void _ZNK7QLocale8toUShortERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto:  QString QLocale::toCurrencyString(ushort , const QString & symbol);
extern void demth_ZNK7QLocale16toCurrencyStringEtRK7QString(void* qthis, int16_t arg0, void* arg1);
  // proto:  double QLocale::toDouble(const QString & s, bool * ok);
extern void _ZNK7QLocale8toDoubleERK7QStringPb(void* qthis, void* arg0, bool* arg1);
  // proto: static QLocale QLocale::system();
extern void _ZN7QLocale6systemEv();
  // proto: static void QLocale::setDefault(const QLocale & locale);
extern void _ZN7QLocale10setDefaultERKS_(void* arg0);
  // proto:  QChar QLocale::exponential();
extern void _ZNK7QLocale11exponentialEv(void* qthis);
  // proto:  QString QLocale::toCurrencyString(float , const QString & symbol);
extern void demth_ZNK7QLocale16toCurrencyStringEfRK7QString(void* qthis, float arg0, void* arg1);
  // proto:  QString QLocale::toString(int i);
extern void demth_ZNK7QLocale8toStringEi(void* qthis, int32_t arg0);
  // proto:  QString QLocale::toString(uint i);
extern void demth_ZNK7QLocale8toStringEj(void* qthis, int32_t arg0);
  // proto:  QString QLocale::toString(qlonglong i);
extern void _ZNK7QLocale8toStringEx(void* qthis, int64_t arg0);
  // proto:  QString QLocale::toString(qulonglong i);
extern void _ZNK7QLocale8toStringEy(void* qthis, int64_t arg0);
  // proto:  QString QLocale::toString(ushort i);
extern void demth_ZNK7QLocale8toStringEt(void* qthis, int16_t arg0);
  // proto:  QString QLocale::amText();
extern void _ZNK7QLocale6amTextEv(void* qthis);
  // proto:  QString QLocale::toCurrencyString(int , const QString & symbol);
extern void demth_ZNK7QLocale16toCurrencyStringEiRK7QString(void* qthis, int32_t arg0, void* arg1);
  // proto:  QString QLocale::toString(short i);
extern void demth_ZNK7QLocale8toStringEs(void* qthis, int16_t arg0);
  // proto:  void QLocale::QLocale();
extern void* dector_ZN7QLocaleC1Ev();
extern void _ZN7QLocaleC1Ev(void* qthis);
  // proto:  int QLocale::toInt(const QString & s, bool * ok);
extern void _ZNK7QLocale5toIntERK7QStringPb(void* qthis, void* arg0, bool* arg1);
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

// class sizeof(QLocale)=1
type QLocale struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QString QLocale::pmText();
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
    // invoke: QString pmText()
    C._ZNK7QLocale6pmTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "pmText", args)
  }

}

  // proto:  QString QLocale::nativeLanguageName();
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
    // invoke: QString nativeLanguageName()
    C._ZNK7QLocale18nativeLanguageNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "nativeLanguageName", args)
  }

}

  // proto:  QString QLocale::toLower(const QString & str);
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
    // invoke: QString toLower(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QLocale7toLowerERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLocale", "toLower", args)
  }

}

  // proto:  QChar QLocale::zeroDigit();
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
    // invoke: QChar zeroDigit()
    C._ZNK7QLocale9zeroDigitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "zeroDigit", args)
  }

}

  // proto:  QString QLocale::name();
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
    // invoke: QString name()
    C._ZNK7QLocale4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "name", args)
  }

}

  // proto:  QString QLocale::toCurrencyString(qlonglong , const QString & symbol);
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
    // invoke: QString toCurrencyString(qlonglong, const class QString &)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale16toCurrencyStringExRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK7QLocale16toCurrencyStringEjRK7QString
    // invoke: QString toCurrencyString(uint, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QLocale16toCurrencyStringEjRK7QString(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QLocale16toCurrencyStringEsRK7QString
    // invoke: QString toCurrencyString(short, const class QString &)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QLocale16toCurrencyStringEsRK7QString(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK7QLocale16toCurrencyStringEdRK7QString
    // invoke: QString toCurrencyString(double, const class QString &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale16toCurrencyStringEdRK7QString(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZNK7QLocale16toCurrencyStringEyRK7QString
    // invoke: QString toCurrencyString(qulonglong, const class QString &)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale16toCurrencyStringEyRK7QString(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZNK7QLocale16toCurrencyStringEtRK7QString
    // invoke: QString toCurrencyString(ushort, const class QString &)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QLocale16toCurrencyStringEtRK7QString(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZNK7QLocale16toCurrencyStringEfRK7QString
    // invoke: QString toCurrencyString(float, const class QString &)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QLocale16toCurrencyStringEfRK7QString(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZNK7QLocale16toCurrencyStringEiRK7QString
    // invoke: QString toCurrencyString(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QLocale16toCurrencyStringEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toCurrencyString", args)
  }

}

  // proto:  float QLocale::toFloat(const QString & s, bool * ok);
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
    // invoke: float toFloat(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale7toFloatERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toFloat", args)
  }

}

  // proto: static QLocale QLocale::c();
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

  // proto:  QString QLocale::createSeparatedList(const QStringList & strl);
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
    // invoke: QString createSeparatedList(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QLocale19createSeparatedListERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLocale", "createSeparatedList", args)
  }

}

  // proto:  uint QLocale::toUInt(const QString & s, bool * ok);
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
    // invoke: uint toUInt(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale6toUIntERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toUInt", args)
  }

}

  // proto:  QChar QLocale::decimalPoint();
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
    // invoke: QChar decimalPoint()
    C._ZNK7QLocale12decimalPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "decimalPoint", args)
  }

}

  // proto:  QChar QLocale::positiveSign();
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
    // invoke: QChar positiveSign()
    C._ZNK7QLocale12positiveSignEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "positiveSign", args)
  }

}

  // proto:  qlonglong QLocale::toLongLong(const QString & s, bool * ok);
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
    // invoke: qlonglong toLongLong(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale10toLongLongERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toLongLong", args)
  }

}

  // proto:  short QLocale::toShort(const QString & s, bool * ok);
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
    // invoke: _ZNK7QLocale7toShortERK7QStringPb
    // invoke: short toShort(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale7toShortERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toShort", args)
  }

}

  // proto:  QString QLocale::toString(float i, char f, int prec);
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
    // invoke: _ZNK7QLocale8toStringEfci
    // invoke: QString toString(float, char, int)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZNK7QLocale8toStringEfci(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK7QLocale8toStringERK9QDateTimeRK7QString
    // invoke: QString toString(const class QDateTime &, const class QString &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale8toStringERK9QDateTimeRK7QString(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QLocale8toStringERK5QTimeRK7QString
    // invoke: QString toString(const class QTime &, const class QString &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale8toStringERK5QTimeRK7QString(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK7QLocale8toStringERK5QDateRK7QString
    // invoke: QString toString(const class QDate &, const class QString &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale8toStringERK5QDateRK7QString(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZNK7QLocale8toStringEdci
    // invoke: QString toString(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK7QLocale8toStringEdci(this.qclsinst, arg0, arg1, arg2)
  case 5:
    // invoke: _ZNK7QLocale8toStringEi
    // invoke: QString toString(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QLocale8toStringEi(this.qclsinst, arg0)
  case 6:
    // invoke: _ZNK7QLocale8toStringEj
    // invoke: QString toString(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QLocale8toStringEj(this.qclsinst, arg0)
  case 7:
    // invoke: _ZNK7QLocale8toStringEx
    // invoke: QString toString(qlonglong)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZNK7QLocale8toStringEx(this.qclsinst, arg0)
  case 8:
    // invoke: _ZNK7QLocale8toStringEy
    // invoke: QString toString(qulonglong)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZNK7QLocale8toStringEy(this.qclsinst, arg0)
  case 9:
    // invoke: _ZNK7QLocale8toStringEt
    // invoke: QString toString(ushort)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QLocale8toStringEt(this.qclsinst, arg0)
  case 10:
    // invoke: _ZNK7QLocale8toStringEs
    // invoke: QString toString(short)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QLocale8toStringEs(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLocale", "toString", args)
  }

}

  // proto:  QDateTime QLocale::toDateTime(const QString & string, const QString & format);
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
    // invoke: QDateTime toDateTime(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale10toDateTimeERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toDateTime", args)
  }

}

  // proto:  QChar QLocale::groupSeparator();
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
    // invoke: QChar groupSeparator()
    C._ZNK7QLocale14groupSeparatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "groupSeparator", args)
  }

}

  // proto:  void QLocale::QLocale(const QString & name);
func NewQLocale(args ...interface{}) QLocale {
  return QLocale{}
}

  // proto:  QDate QLocale::toDate(const QString & string, const QString & format);
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
    // invoke: _ZNK7QLocale6toDateERK7QStringS2_
    // invoke: QDate toDate(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale6toDateERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toDate", args)
  }

}

  // proto:  QString QLocale::nativeCountryName();
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
    // invoke: QString nativeCountryName()
    C._ZNK7QLocale17nativeCountryNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "nativeCountryName", args)
  }

}

  // proto:  QChar QLocale::negativeSign();
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
    // invoke: QChar negativeSign()
    C._ZNK7QLocale12negativeSignEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "negativeSign", args)
  }

}

  // proto:  void QLocale::~QLocale();
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

  // proto:  QString QLocale::toUpper(const QString & str);
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
    // invoke: QString toUpper(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QLocale7toUpperERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLocale", "toUpper", args)
  }

}

  // proto:  QChar QLocale::percent();
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
    // invoke: QChar percent()
    C._ZNK7QLocale7percentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "percent", args)
  }

}

  // proto:  qulonglong QLocale::toULongLong(const QString & s, bool * ok);
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
    // invoke: _ZNK7QLocale11toULongLongERK7QStringPb
    // invoke: qulonglong toULongLong(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale11toULongLongERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toULongLong", args)
  }

}

  // proto:  QStringList QLocale::uiLanguages();
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
    // invoke: QStringList uiLanguages()
    C._ZNK7QLocale11uiLanguagesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "uiLanguages", args)
  }

}

  // proto:  QString QLocale::bcp47Name();
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
    // invoke: QString bcp47Name()
    C._ZNK7QLocale9bcp47NameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "bcp47Name", args)
  }

}

  // proto:  QTime QLocale::toTime(const QString & string, const QString & format);
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
    // invoke: _ZNK7QLocale6toTimeERK7QStringS2_
    // invoke: QTime toTime(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale6toTimeERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toTime", args)
  }

}

  // proto:  ushort QLocale::toUShort(const QString & s, bool * ok);
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
    // invoke: _ZNK7QLocale8toUShortERK7QStringPb
    // invoke: ushort toUShort(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale8toUShortERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toUShort", args)
  }

}

  // proto:  double QLocale::toDouble(const QString & s, bool * ok);
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
    // invoke: _ZNK7QLocale8toDoubleERK7QStringPb
    // invoke: double toDouble(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale8toDoubleERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toDouble", args)
  }

}

  // proto: static QLocale QLocale::system();
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

  // proto: static void QLocale::setDefault(const QLocale & locale);
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

  // proto:  QChar QLocale::exponential();
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
    // invoke: QChar exponential()
    C._ZNK7QLocale11exponentialEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "exponential", args)
  }

}

  // proto:  QString QLocale::amText();
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
    // invoke: QString amText()
    C._ZNK7QLocale6amTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "amText", args)
  }

}

  // proto:  int QLocale::toInt(const QString & s, bool * ok);
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
    // invoke: _ZNK7QLocale5toIntERK7QStringPb
    // invoke: int toInt(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK7QLocale5toIntERK7QStringPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLocale", "toInt", args)
  }

}

// <= body block end

