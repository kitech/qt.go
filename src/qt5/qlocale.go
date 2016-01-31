package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  ushort QLocale::toUShort(const QString & s, bool * ok);
extern int16_t C_ZNK7QLocale8toUShortERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QTime QLocale::toTime(const QString & string, const QString & format);
extern void* C_ZNK7QLocale6toTimeERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QLocale::uiLanguages();
extern void C_ZNK7QLocale11uiLanguagesEv(void* qthis); // 4
  // proto:  uint QLocale::toUInt(const QString & s, bool * ok);
extern int32_t C_ZNK7QLocale6toUIntERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  float QLocale::toFloat(const QString & s, bool * ok);
extern float C_ZNK7QLocale7toFloatERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QLocale QLocale::system();
extern void* C_ZN7QLocale6systemEv(); // 4
  // proto:  short QLocale::toShort(const QString & s, bool * ok);
extern int16_t C_ZNK7QLocale7toShortERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QLocale::amText();
extern void* C_ZNK7QLocale6amTextEv(void* qthis); // 4
  // proto: static void QLocale::setDefault(const QLocale & locale);
extern void C_ZN7QLocale10setDefaultERKS_(void* arg0); // 4
  // proto:  QDate QLocale::toDate(const QString & string, const QString & format);
extern void* C_ZNK7QLocale6toDateERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QLocale::toLower(const QString & str);
extern void* C_ZNK7QLocale7toLowerERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QLocale::pmText();
extern void* C_ZNK7QLocale6pmTextEv(void* qthis); // 4
  // proto:  Qt::DayOfWeek QLocale::firstDayOfWeek();
extern void C_ZNK7QLocale14firstDayOfWeekEv(void* qthis); // 4
  // proto:  QChar QLocale::percent();
extern void* C_ZNK7QLocale7percentEv(void* qthis); // 4
  // proto:  QList<Qt::DayOfWeek> QLocale::weekdays();
extern void C_ZNK7QLocale8weekdaysEv(void* qthis); // 4
  // proto:  qlonglong QLocale::toLongLong(const QString & s, bool * ok);
extern int64_t C_ZNK7QLocale10toLongLongERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QChar QLocale::negativeSign();
extern void* C_ZNK7QLocale12negativeSignEv(void* qthis); // 4
  // proto:  QString QLocale::toString(int i);
extern void* C_ZNK7QLocale8toStringEi(void* qthis, int32_t arg0); // 2
  // proto:  QString QLocale::toString(uint i);
extern void* C_ZNK7QLocale8toStringEj(void* qthis, int32_t arg0); // 2
  // proto:  QString QLocale::toString(const QTime & time, const QString & formatStr);
extern void* C_ZNK7QLocale8toStringERK5QTimeRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QLocale::toString(double i, char f, int prec);
extern void* C_ZNK7QLocale8toStringEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QString QLocale::toString(float i, char f, int prec);
extern void* C_ZNK7QLocale8toStringEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2); // 2
  // proto:  QString QLocale::toString(ushort i);
extern void* C_ZNK7QLocale8toStringEt(void* qthis, int16_t arg0); // 2
  // proto:  QString QLocale::toString(const QDateTime & dateTime, const QString & format);
extern void* C_ZNK7QLocale8toStringERK9QDateTimeRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QLocale::toString(qlonglong i);
extern void* C_ZNK7QLocale8toStringEx(void* qthis, int64_t arg0); // 4
  // proto:  QString QLocale::toString(qulonglong i);
extern void* C_ZNK7QLocale8toStringEy(void* qthis, int64_t arg0); // 4
  // proto:  QString QLocale::toString(const QDate & date, const QString & formatStr);
extern void* C_ZNK7QLocale8toStringERK5QDateRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QLocale::toString(short i);
extern void* C_ZNK7QLocale8toStringEs(void* qthis, int16_t arg0); // 2
  // proto:  void QLocale::~QLocale();
extern void C_ZN7QLocaleD2Ev(void* qthis); // 4
  // proto:  QString QLocale::createSeparatedList(const QStringList & strl);
extern void* C_ZNK7QLocale19createSeparatedListERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QString QLocale::nativeLanguageName();
extern void* C_ZNK7QLocale18nativeLanguageNameEv(void* qthis); // 4
  // proto:  QChar QLocale::groupSeparator();
extern void* C_ZNK7QLocale14groupSeparatorEv(void* qthis); // 4
  // proto:  void QLocale::QLocale(const QLocale & other);
extern void* C_ZN7QLocaleC2ERKS_(void* arg0); // 3
  // proto:  void QLocale::QLocale(const QString & name);
extern void* C_ZN7QLocaleC2ERK7QString(void* arg0); // 3
  // proto:  void QLocale::QLocale();
extern void* C_ZN7QLocaleC2Ev(); // 3
  // proto:  QChar QLocale::positiveSign();
extern void* C_ZNK7QLocale12positiveSignEv(void* qthis); // 4
  // proto:  QString QLocale::toCurrencyString(double , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEdRK7QString(void* qthis, double arg0, void* arg1); // 4
  // proto:  QString QLocale::toCurrencyString(uint , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEjRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QString QLocale::toCurrencyString(short , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEsRK7QString(void* qthis, int16_t arg0, void* arg1); // 2
  // proto:  QString QLocale::toCurrencyString(qulonglong , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEyRK7QString(void* qthis, int64_t arg0, void* arg1); // 4
  // proto:  QString QLocale::toCurrencyString(float , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEfRK7QString(void* qthis, float arg0, void* arg1); // 2
  // proto:  QString QLocale::toCurrencyString(ushort , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEtRK7QString(void* qthis, int16_t arg0, void* arg1); // 2
  // proto:  QString QLocale::toCurrencyString(qlonglong , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringExRK7QString(void* qthis, int64_t arg0, void* arg1); // 4
  // proto:  QString QLocale::toCurrencyString(int , const QString & symbol);
extern void* C_ZNK7QLocale16toCurrencyStringEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QString QLocale::nativeCountryName();
extern void* C_ZNK7QLocale17nativeCountryNameEv(void* qthis); // 4
  // proto:  QLocale::Language QLocale::language();
extern void C_ZNK7QLocale8languageEv(void* qthis); // 4
  // proto:  QString QLocale::bcp47Name();
extern void* C_ZNK7QLocale9bcp47NameEv(void* qthis); // 4
  // proto:  NumberOptions QLocale::numberOptions();
extern void C_ZNK7QLocale13numberOptionsEv(void* qthis); // 4
  // proto:  qulonglong QLocale::toULongLong(const QString & s, bool * ok);
extern int64_t C_ZNK7QLocale11toULongLongERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QLocale QLocale::c();
extern void* C_ZN7QLocale1cEv(); // 2
  // proto:  QString QLocale::name();
extern void* C_ZNK7QLocale4nameEv(void* qthis); // 4
  // proto:  QChar QLocale::exponential();
extern void* C_ZNK7QLocale11exponentialEv(void* qthis); // 4
  // proto:  QLocale::Country QLocale::country();
extern void C_ZNK7QLocale7countryEv(void* qthis); // 4
  // proto:  int QLocale::toInt(const QString & s, bool * ok);
extern int32_t C_ZNK7QLocale5toIntERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  Qt::LayoutDirection QLocale::textDirection();
extern void C_ZNK7QLocale13textDirectionEv(void* qthis); // 4
  // proto:  QLocale::Script QLocale::script();
extern void C_ZNK7QLocale6scriptEv(void* qthis); // 4
  // proto:  double QLocale::toDouble(const QString & s, bool * ok);
extern double C_ZNK7QLocale8toDoubleERK7QStringPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QDateTime QLocale::toDateTime(const QString & string, const QString & format);
extern void* C_ZNK7QLocale10toDateTimeERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QChar QLocale::decimalPoint();
extern void* C_ZNK7QLocale12decimalPointEv(void* qthis); // 4
  // proto:  QString QLocale::toUpper(const QString & str);
extern void* C_ZNK7QLocale7toUpperERK7QString(void* qthis, void* arg0); // 4
  // proto:  QChar QLocale::zeroDigit();
extern void* C_ZNK7QLocale9zeroDigitEv(void* qthis); // 4
  // proto:  QLocale::MeasurementSystem QLocale::measurementSystem();
extern void C_ZNK7QLocale17measurementSystemEv(void* qthis); // 4
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

// toUShort(const class QString &, _Bool *)
func (this *QLocale) Toushort(args ...interface{}) (ret interface{}) {
  // toUShort(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8toUShortERK7QStringPb
    // invoke: ushort toUShort(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale8toUShortERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toUShort", args)
  }

  return
}

// toTime(const class QString &, const class QString &)
func (this *QLocale) Totime(args ...interface{}) (ret interface{}) {
  // toTime(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

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
    var ret0 = C.C_ZNK7QLocale6toTimeERK7QStringS2_(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toTime", args)
  }

  return
}

// uiLanguages()
func (this *QLocale) Uilanguages(args ...interface{}) () {
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
    C.C_ZNK7QLocale11uiLanguagesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "uiLanguages", args)
  }

  return
}

// toUInt(const class QString &, _Bool *)
func (this *QLocale) Touint(args ...interface{}) (ret interface{}) {
  // toUInt(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6toUIntERK7QStringPb
    // invoke: uint toUInt(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale6toUIntERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toUInt", args)
  }

  return
}

// toFloat(const class QString &, _Bool *)
func (this *QLocale) Tofloat(args ...interface{}) (ret interface{}) {
  // toFloat(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7toFloatERK7QStringPb
    // invoke: float toFloat(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale7toFloatERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toFloat", args)
  }

  return
}

// system()
func (this *QLocale) System_S(args ...interface{}) (ret interface{}) {
  // system()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLocale6systemEv
    // invoke: QLocale system()
    var ret0 = C.C_ZN7QLocale6systemEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLocale{}) // "QLocale"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "system", args)
  }

  return
}

// toShort(const class QString &, _Bool *)
func (this *QLocale) Toshort(args ...interface{}) (ret interface{}) {
  // toShort(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7toShortERK7QStringPb
    // invoke: short toShort(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale7toShortERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "short"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toShort", args)
  }

  return
}

// amText()
func (this *QLocale) Amtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale6amTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "amText", args)
  }

  return
}

// setDefault(const class QLocale &)
func (this *QLocale) Setdefault_S(args ...interface{}) () {
  // setDefault(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLocale10setDefaultERKS_
    // invoke: void setDefault(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLocale10setDefaultERKS_(arg0)
  default:
    qtrt.ErrorResolve("QLocale", "setDefault", args)
  }

  return
}

// toDate(const class QString &, const class QString &)
func (this *QLocale) Todate(args ...interface{}) (ret interface{}) {
  // toDate(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

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
    var ret0 = C.C_ZNK7QLocale6toDateERK7QStringS2_(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toDate", args)
  }

  return
}

// toLower(const class QString &)
func (this *QLocale) Tolower(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale7toLowerERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toLower", args)
  }

  return
}

// pmText()
func (this *QLocale) Pmtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale6pmTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "pmText", args)
  }

  return
}

// firstDayOfWeek()
func (this *QLocale) Firstdayofweek(args ...interface{}) () {
  // firstDayOfWeek()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale14firstDayOfWeekEv
    // invoke: Qt::DayOfWeek firstDayOfWeek()
    C.C_ZNK7QLocale14firstDayOfWeekEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "firstDayOfWeek", args)
  }

  return
}

// percent()
func (this *QLocale) Percent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale7percentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "percent", args)
  }

  return
}

// weekdays()
func (this *QLocale) Weekdays(args ...interface{}) () {
  // weekdays()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8weekdaysEv
    // invoke: QList<Qt::DayOfWeek> weekdays()
    C.C_ZNK7QLocale8weekdaysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "weekdays", args)
  }

  return
}

// toLongLong(const class QString &, _Bool *)
func (this *QLocale) Tolonglong(args ...interface{}) (ret interface{}) {
  // toLongLong(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale10toLongLongERK7QStringPb
    // invoke: qlonglong toLongLong(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale10toLongLongERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qlonglong"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toLongLong", args)
  }

  return
}

// negativeSign()
func (this *QLocale) Negativesign(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale12negativeSignEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "negativeSign", args)
  }

  return
}

// toString(int)
func (this *QLocale) Tostring(args ...interface{}) (ret interface{}) {
  // toString(int)
  // toString(uint)
  // toString(const class QTime &, const class QString &)
  // toString(double, char, int)
  // toString(float, char, int)
  // toString(ushort)
  // toString(const class QDateTime &, const class QString &)
  // toString(qlonglong)
  // toString(qulonglong)
  // toString(const class QDate &, const class QString &)
  // toString(short)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "double"
  vtys[3][1] = qtrt.ByteTy(false) // "char"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.FloatTy(false) // "float"
  vtys[4][1] = qtrt.ByteTy(false) // "char"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[9][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int16Ty(false) // "short"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8toStringEi
    // invoke: QString toString(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLocale8toStringEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK7QLocale8toStringEj
    // invoke: QString toString(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLocale8toStringEj(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK7QLocale8toStringERK5QTimeRK7QString
    // invoke: QString toString(const class QTime &, const class QString &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale8toStringERK5QTimeRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 3:
    // invoke: _ZNK7QLocale8toStringEdci
    // invoke: QString toString(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QLocale8toStringEdci(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 4:
    // invoke: _ZNK7QLocale8toStringEfci
    // invoke: QString toString(float, char, int)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QLocale8toStringEfci(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 5:
    // invoke: _ZNK7QLocale8toStringEt
    // invoke: QString toString(ushort)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLocale8toStringEt(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 6:
    // invoke: _ZNK7QLocale8toStringERK9QDateTimeRK7QString
    // invoke: QString toString(const class QDateTime &, const class QString &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale8toStringERK9QDateTimeRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 7:
    // invoke: _ZNK7QLocale8toStringEx
    // invoke: QString toString(qlonglong)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLocale8toStringEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 8:
    // invoke: _ZNK7QLocale8toStringEy
    // invoke: QString toString(qulonglong)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLocale8toStringEy(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 9:
    // invoke: _ZNK7QLocale8toStringERK5QDateRK7QString
    // invoke: QString toString(const class QDate &, const class QString &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale8toStringERK5QDateRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 10:
    // invoke: _ZNK7QLocale8toStringEs
    // invoke: QString toString(short)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLocale8toStringEs(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toString", args)
  }

  return
}

// ~QLocale()
func (this *QLocale) Freeqlocale(args ...interface{}) () {
  // ~QLocale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLocaleD0Ev
    // invoke: void ~QLocale()
    C.C_ZN7QLocaleD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "~QLocale", args)
  }

  return
}

// createSeparatedList(const class QStringList &)
func (this *QLocale) Createseparatedlist(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale19createSeparatedListERK11QStringList(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "createSeparatedList", args)
  }

  return
}

// nativeLanguageName()
func (this *QLocale) Nativelanguagename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale18nativeLanguageNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "nativeLanguageName", args)
  }

  return
}

// groupSeparator()
func (this *QLocale) Groupseparator(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale14groupSeparatorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "groupSeparator", args)
  }

  return
}

// QLocale(const class QLocale &)
func NewQLocale(args ...interface{}) *QLocale {
  // QLocale(const class QLocale &)
  // QLocale(const class QString &)
  // QLocale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLocaleC1ERKS_
    // invoke: void QLocale(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QLocaleC2ERKS_(arg0)
    return &QLocale{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QLocaleC1ERK7QString
    // invoke: void QLocale(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QLocaleC2ERK7QString(arg0)
    return &QLocale{qclsinst:qthis}
  case 2:
    // invoke: _ZN7QLocaleC1Ev
    // invoke: void QLocale()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QLocaleC2Ev()
    return &QLocale{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLocale", "QLocale", args)
  }

  return nil // QLocale{qclsinst:qthis}
}

// positiveSign()
func (this *QLocale) Positivesign(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale12positiveSignEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "positiveSign", args)
  }

  return
}

// toCurrencyString(double, const class QString &)
func (this *QLocale) Tocurrencystring(args ...interface{}) (ret interface{}) {
  // toCurrencyString(double, const class QString &)
  // toCurrencyString(uint, const class QString &)
  // toCurrencyString(short, const class QString &)
  // toCurrencyString(qulonglong, const class QString &)
  // toCurrencyString(float, const class QString &)
  // toCurrencyString(ushort, const class QString &)
  // toCurrencyString(qlonglong, const class QString &)
  // toCurrencyString(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int16Ty(false) // "short"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.FloatTy(false) // "float"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[5][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale16toCurrencyStringEdRK7QString
    // invoke: QString toCurrencyString(double, const class QString &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEdRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK7QLocale16toCurrencyStringEjRK7QString
    // invoke: QString toCurrencyString(uint, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEjRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK7QLocale16toCurrencyStringEsRK7QString
    // invoke: QString toCurrencyString(short, const class QString &)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEsRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 3:
    // invoke: _ZNK7QLocale16toCurrencyStringEyRK7QString
    // invoke: QString toCurrencyString(qulonglong, const class QString &)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEyRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 4:
    // invoke: _ZNK7QLocale16toCurrencyStringEfRK7QString
    // invoke: QString toCurrencyString(float, const class QString &)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEfRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 5:
    // invoke: _ZNK7QLocale16toCurrencyStringEtRK7QString
    // invoke: QString toCurrencyString(ushort, const class QString &)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEtRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 6:
    // invoke: _ZNK7QLocale16toCurrencyStringExRK7QString
    // invoke: QString toCurrencyString(qlonglong, const class QString &)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringExRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 7:
    // invoke: _ZNK7QLocale16toCurrencyStringEiRK7QString
    // invoke: QString toCurrencyString(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale16toCurrencyStringEiRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toCurrencyString", args)
  }

  return
}

// nativeCountryName()
func (this *QLocale) Nativecountryname(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale17nativeCountryNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "nativeCountryName", args)
  }

  return
}

// language()
func (this *QLocale) Language(args ...interface{}) () {
  // language()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8languageEv
    // invoke: QLocale::Language language()
    C.C_ZNK7QLocale8languageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "language", args)
  }

  return
}

// bcp47Name()
func (this *QLocale) Bcp47Name(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale9bcp47NameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "bcp47Name", args)
  }

  return
}

// numberOptions()
func (this *QLocale) Numberoptions(args ...interface{}) () {
  // numberOptions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale13numberOptionsEv
    // invoke: NumberOptions numberOptions()
    C.C_ZNK7QLocale13numberOptionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "numberOptions", args)
  }

  return
}

// toULongLong(const class QString &, _Bool *)
func (this *QLocale) Toulonglong(args ...interface{}) (ret interface{}) {
  // toULongLong(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale11toULongLongERK7QStringPb
    // invoke: qulonglong toULongLong(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale11toULongLongERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qulonglong"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toULongLong", args)
  }

  return
}

// c()
func (this *QLocale) C_S(args ...interface{}) (ret interface{}) {
  // c()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLocale1cEv
    // invoke: QLocale c()
    var ret0 = C.C_ZN7QLocale1cEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLocale{}) // "QLocale"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "c", args)
  }

  return
}

// name()
func (this *QLocale) Name(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "name", args)
  }

  return
}

// exponential()
func (this *QLocale) Exponential(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale11exponentialEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "exponential", args)
  }

  return
}

// country()
func (this *QLocale) Country(args ...interface{}) () {
  // country()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale7countryEv
    // invoke: QLocale::Country country()
    C.C_ZNK7QLocale7countryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "country", args)
  }

  return
}

// toInt(const class QString &, _Bool *)
func (this *QLocale) Toint(args ...interface{}) (ret interface{}) {
  // toInt(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale5toIntERK7QStringPb
    // invoke: int toInt(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale5toIntERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toInt", args)
  }

  return
}

// textDirection()
func (this *QLocale) Textdirection(args ...interface{}) () {
  // textDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale13textDirectionEv
    // invoke: Qt::LayoutDirection textDirection()
    C.C_ZNK7QLocale13textDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "textDirection", args)
  }

  return
}

// script()
func (this *QLocale) Script(args ...interface{}) () {
  // script()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale6scriptEv
    // invoke: QLocale::Script script()
    C.C_ZNK7QLocale6scriptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "script", args)
  }

  return
}

// toDouble(const class QString &, _Bool *)
func (this *QLocale) Todouble(args ...interface{}) (ret interface{}) {
  // toDouble(const class QString &, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale8toDoubleERK7QStringPb
    // invoke: double toDouble(const class QString &, _Bool *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QLocale8toDoubleERK7QStringPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toDouble", args)
  }

  return
}

// toDateTime(const class QString &, const class QString &)
func (this *QLocale) Todatetime(args ...interface{}) (ret interface{}) {
  // toDateTime(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

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
    var ret0 = C.C_ZNK7QLocale10toDateTimeERK7QStringS2_(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toDateTime", args)
  }

  return
}

// decimalPoint()
func (this *QLocale) Decimalpoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale12decimalPointEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "decimalPoint", args)
  }

  return
}

// toUpper(const class QString &)
func (this *QLocale) Toupper(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale7toUpperERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "toUpper", args)
  }

  return
}

// zeroDigit()
func (this *QLocale) Zerodigit(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QLocale9zeroDigitEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLocale", "zeroDigit", args)
  }

  return
}

// measurementSystem()
func (this *QLocale) Measurementsystem(args ...interface{}) () {
  // measurementSystem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLocale17measurementSystemEv
    // invoke: QLocale::MeasurementSystem measurementSystem()
    C.C_ZNK7QLocale17measurementSystemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLocale", "measurementSystem", args)
  }

  return
}

// <= body block end

