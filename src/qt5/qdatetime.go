package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qdatetime.h
// dst-file: /src/core/qdatetime.go
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
  // proto: static QTime QTime::fromString(const QString & s, const QString & format);
extern void* C_ZN5QTime10fromStringERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  void QTime::QTime(int h, int m, int s, int ms);
extern void* C_ZN5QTimeC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 3
  // proto:  void QTime::QTime();
extern void* C_ZN5QTimeC2Ev(); // 1
  // proto:  int QTime::msec();
extern int32_t C_ZNK5QTime4msecEv(void* qthis); // 4
  // proto:  int QTime::second();
extern int32_t C_ZNK5QTime6secondEv(void* qthis); // 4
  // proto:  int QTime::msecsSinceStartOfDay();
extern int32_t C_ZNK5QTime20msecsSinceStartOfDayEv(void* qthis); // 2
  // proto:  bool QTime::setHMS(int h, int m, int s, int ms);
extern bool C_ZN5QTime6setHMSEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QTime QTime::addMSecs(int ms);
extern void* C_ZNK5QTime8addMSecsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTime::start();
extern void C_ZN5QTime5startEv(void* qthis); // 4
  // proto: static QTime QTime::fromMSecsSinceStartOfDay(int msecs);
extern void* C_ZN5QTime24fromMSecsSinceStartOfDayEi(int32_t arg0); // 2
  // proto:  int QTime::minute();
extern int32_t C_ZNK5QTime6minuteEv(void* qthis); // 4
  // proto:  QString QTime::toString(const QString & format);
extern void* C_ZNK5QTime8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QTime::isValid(int h, int m, int s, int ms);
extern bool C_ZN5QTime7isValidEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  bool QTime::isValid();
extern bool C_ZNK5QTime7isValidEv(void* qthis); // 4
  // proto:  int QTime::secsTo(const QTime & );
extern int32_t C_ZNK5QTime6secsToERKS_(void* qthis, void* arg0); // 4
  // proto:  int QTime::elapsed();
extern int32_t C_ZNK5QTime7elapsedEv(void* qthis); // 4
  // proto:  QTime QTime::addSecs(int secs);
extern void* C_ZNK5QTime7addSecsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTime::msecsTo(const QTime & );
extern int32_t C_ZNK5QTime7msecsToERKS_(void* qthis, void* arg0); // 4
  // proto:  int QTime::restart();
extern int32_t C_ZN5QTime7restartEv(void* qthis); // 4
  // proto:  int QTime::hour();
extern int32_t C_ZNK5QTime4hourEv(void* qthis); // 4
  // proto: static QTime QTime::currentTime();
extern void* C_ZN5QTime11currentTimeEv(); // 4
  // proto:  bool QTime::isNull();
extern bool C_ZNK5QTime6isNullEv(void* qthis); // 2
  // proto: static QDateTime QDateTime::fromString(const QString & s, const QString & format);
extern void* C_ZN9QDateTime10fromStringERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  void QDateTime::setOffsetFromUtc(int offsetSeconds);
extern void C_ZN9QDateTime16setOffsetFromUtcEi(void* qthis, int32_t arg0); // 4
  // proto:  uint QDateTime::toTime_t();
extern int32_t C_ZNK9QDateTime8toTime_tEv(void* qthis); // 4
  // proto:  QDateTime QDateTime::toTimeZone(const QTimeZone & toZone);
extern void* C_ZNK9QDateTime10toTimeZoneERK9QTimeZone(void* qthis, void* arg0); // 4
  // proto:  QDateTime QDateTime::toUTC();
extern void* C_ZNK9QDateTime5toUTCEv(void* qthis); // 2
  // proto:  void QDateTime::setUtcOffset(int seconds);
extern void C_ZN9QDateTime12setUtcOffsetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDateTime::setTimeZone(const QTimeZone & toZone);
extern void C_ZN9QDateTime11setTimeZoneERK9QTimeZone(void* qthis, void* arg0); // 4
  // proto:  void QDateTime::QDateTime(const QDateTime & other);
extern void* C_ZN9QDateTimeC2ERKS_(void* arg0); // 3
  // proto:  void QDateTime::QDateTime(const QDate & );
extern void* C_ZN9QDateTimeC2ERK5QDate(void* arg0); // 3
  // proto:  void QDateTime::QDateTime();
extern void* C_ZN9QDateTimeC2Ev(); // 3
  // proto:  void QDateTime::QDateTime(const QDate & date, const QTime & time, const QTimeZone & timeZone);
extern void* C_ZN9QDateTimeC2ERK5QDateRK5QTimeRK9QTimeZone(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QDateTime::~QDateTime();
extern void C_ZN9QDateTimeD2Ev(void* qthis); // 4
  // proto:  QDateTime QDateTime::toOffsetFromUtc(int offsetSeconds);
extern void* C_ZNK9QDateTime15toOffsetFromUtcEi(void* qthis, int32_t arg0); // 4
  // proto:  QDateTime QDateTime::addMSecs(qint64 msecs);
extern void* C_ZNK9QDateTime8addMSecsEx(void* qthis, int64_t arg0); // 4
  // proto:  QDateTime QDateTime::addMonths(int months);
extern void* C_ZNK9QDateTime9addMonthsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QDateTime::offsetFromUtc();
extern int32_t C_ZNK9QDateTime13offsetFromUtcEv(void* qthis); // 4
  // proto:  QString QDateTime::toString(const QString & format);
extern void* C_ZNK9QDateTime8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QDateTime::swap(QDateTime & other);
extern void C_ZN9QDateTime4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QDateTime::setDate(const QDate & date);
extern void C_ZN9QDateTime7setDateERK5QDate(void* qthis, void* arg0); // 4
  // proto: static QDateTime QDateTime::currentDateTime();
extern void* C_ZN9QDateTime15currentDateTimeEv(); // 4
  // proto: static QDateTime QDateTime::currentDateTimeUtc();
extern void* C_ZN9QDateTime18currentDateTimeUtcEv(); // 4
  // proto: static qint64 QDateTime::currentMSecsSinceEpoch();
extern int64_t C_ZN9QDateTime22currentMSecsSinceEpochEv(); // 4
  // proto:  bool QDateTime::isValid();
extern bool C_ZNK9QDateTime7isValidEv(void* qthis); // 4
  // proto:  qint64 QDateTime::secsTo(const QDateTime & );
extern int64_t C_ZNK9QDateTime6secsToERKS_(void* qthis, void* arg0); // 4
  // proto:  void QDateTime::setMSecsSinceEpoch(qint64 msecs);
extern void C_ZN9QDateTime18setMSecsSinceEpochEx(void* qthis, int64_t arg0); // 4
  // proto:  QDateTime QDateTime::addDays(qint64 days);
extern void* C_ZNK9QDateTime7addDaysEx(void* qthis, int64_t arg0); // 4
  // proto:  void QDateTime::setTime_t(uint secsSince1Jan1970UTC);
extern void C_ZN9QDateTime9setTime_tEj(void* qthis, int32_t arg0); // 4
  // proto:  QDateTime QDateTime::addSecs(qint64 secs);
extern void* C_ZNK9QDateTime7addSecsEx(void* qthis, int64_t arg0); // 4
  // proto:  QString QDateTime::timeZoneAbbreviation();
extern void* C_ZNK9QDateTime20timeZoneAbbreviationEv(void* qthis); // 4
  // proto:  void QDateTime::setTime(const QTime & time);
extern void C_ZN9QDateTime7setTimeERK5QTime(void* qthis, void* arg0); // 4
  // proto:  qint64 QDateTime::toMSecsSinceEpoch();
extern int64_t C_ZNK9QDateTime17toMSecsSinceEpochEv(void* qthis); // 4
  // proto:  qint64 QDateTime::msecsTo(const QDateTime & );
extern int64_t C_ZNK9QDateTime7msecsToERKS_(void* qthis, void* arg0); // 4
  // proto:  QDate QDateTime::date();
extern void* C_ZNK9QDateTime4dateEv(void* qthis); // 4
  // proto:  qint64 QDateTime::daysTo(const QDateTime & );
extern int64_t C_ZNK9QDateTime6daysToERKS_(void* qthis, void* arg0); // 4
  // proto:  QDateTime QDateTime::toLocalTime();
extern void* C_ZNK9QDateTime11toLocalTimeEv(void* qthis); // 2
  // proto:  Qt::TimeSpec QDateTime::timeSpec();
extern void C_ZNK9QDateTime8timeSpecEv(void* qthis); // 4
  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC, const QTimeZone & timeZone);
extern void* C_ZN9QDateTime10fromTime_tEjRK9QTimeZone(int32_t arg0, void* arg1); // 4
  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC);
extern void* C_ZN9QDateTime10fromTime_tEj(int32_t arg0); // 4
  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs);
extern void* C_ZN9QDateTime19fromMSecsSinceEpochEx(int64_t arg0); // 4
  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs, const QTimeZone & timeZone);
extern void* C_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone(int64_t arg0, void* arg1); // 4
  // proto:  bool QDateTime::isNull();
extern bool C_ZNK9QDateTime6isNullEv(void* qthis); // 4
  // proto:  int QDateTime::utcOffset();
extern int32_t C_ZNK9QDateTime9utcOffsetEv(void* qthis); // 4
  // proto:  QTime QDateTime::time();
extern void* C_ZNK9QDateTime4timeEv(void* qthis); // 4
  // proto:  QTimeZone QDateTime::timeZone();
extern void* C_ZNK9QDateTime8timeZoneEv(void* qthis); // 4
  // proto:  bool QDateTime::isDaylightTime();
extern bool C_ZNK9QDateTime14isDaylightTimeEv(void* qthis); // 4
  // proto:  QDateTime QDateTime::addYears(int years);
extern void* C_ZNK9QDateTime8addYearsEi(void* qthis, int32_t arg0); // 4
  // proto: static QDate QDate::fromString(const QString & s, const QString & format);
extern void* C_ZN5QDate10fromStringERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  int QDate::month();
extern int32_t C_ZNK5QDate5monthEv(void* qthis); // 4
  // proto:  int QDate::year();
extern int32_t C_ZNK5QDate4yearEv(void* qthis); // 4
  // proto:  int QDate::daysInMonth();
extern int32_t C_ZNK5QDate11daysInMonthEv(void* qthis); // 4
  // proto:  qint64 QDate::daysTo(const QDate & );
extern int64_t C_ZNK5QDate6daysToERKS_(void* qthis, void* arg0); // 4
  // proto:  int QDate::weekNumber(int * yearNum);
extern int32_t C_ZNK5QDate10weekNumberEPi(void* qthis, void* arg0); // 4
  // proto:  int QDate::daysInYear();
extern int32_t C_ZNK5QDate10daysInYearEv(void* qthis); // 4
  // proto:  QDate QDate::addMonths(int months);
extern void* C_ZNK5QDate9addMonthsEi(void* qthis, int32_t arg0); // 4
  // proto:  qint64 QDate::toJulianDay();
extern int64_t C_ZNK5QDate11toJulianDayEv(void* qthis); // 2
  // proto:  QString QDate::toString(const QString & format);
extern void* C_ZNK5QDate8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto: static QDate QDate::fromJulianDay(qint64 jd);
extern void* C_ZN5QDate13fromJulianDayEx(int64_t arg0); // 2
  // proto:  bool QDate::setDate(int year, int month, int day);
extern bool C_ZN5QDate7setDateEiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QDate::getDate(int * year, int * month, int * day);
extern void C_ZN5QDate7getDateEPiS0_S0_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  int QDate::dayOfWeek();
extern int32_t C_ZNK5QDate9dayOfWeekEv(void* qthis); // 4
  // proto: static bool QDate::isLeapYear(int year);
extern bool C_ZN5QDate10isLeapYearEi(int32_t arg0); // 4
  // proto:  void QDate::QDate();
extern void* C_ZN5QDateC2Ev(); // 1
  // proto:  void QDate::QDate(int y, int m, int d);
extern void* C_ZN5QDateC2Eiii(int32_t arg0, int32_t arg1, int32_t arg2); // 3
  // proto:  bool QDate::isValid();
extern bool C_ZNK5QDate7isValidEv(void* qthis); // 2
  // proto: static bool QDate::isValid(int y, int m, int d);
extern bool C_ZN5QDate7isValidEiii(int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  QDate QDate::addDays(qint64 days);
extern void* C_ZNK5QDate7addDaysEx(void* qthis, int64_t arg0); // 4
  // proto:  int QDate::day();
extern int32_t C_ZNK5QDate3dayEv(void* qthis); // 4
  // proto:  QDate QDate::addYears(int years);
extern void* C_ZNK5QDate8addYearsEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QDate::isNull();
extern bool C_ZNK5QDate6isNullEv(void* qthis); // 2
  // proto:  int QDate::dayOfYear();
extern int32_t C_ZNK5QDate9dayOfYearEv(void* qthis); // 4
  // proto: static QDate QDate::currentDate();
extern void* C_ZN5QDate11currentDateEv(); // 4
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

// class sizeof(QTime)=4
type QTime struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDateTime)=1
type QDateTime struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDate)=8
type QDate struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// fromString(const class QString &, const class QString &)
func (this *QTime) Fromstring_S(args ...interface{}) (ret interface{}) {
  // fromString(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime10fromStringERK7QStringS2_
    // invoke: QTime fromString(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QTime10fromStringERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "fromString", args)
  }

  return
}

// QTime(int, int, int, int)
func NewQTime(args ...interface{}) *QTime {
  // QTime(int, int, int, int)
  // QTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTimeC1Eiiii
    // invoke: void QTime(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QTimeC2Eiiii(arg0, arg1, arg2, arg3)
    return &QTime{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QTimeC1Ev
    // invoke: void QTime()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QTimeC2Ev()
    return &QTime{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTime", "QTime", args)
  }

  return nil // QTime{qclsinst:qthis}
}

// msec()
func (this *QTime) Msec(args ...interface{}) (ret interface{}) {
  // msec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime4msecEv
    // invoke: int msec()
    var ret0 = C.C_ZNK5QTime4msecEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "msec", args)
  }

  return
}

// second()
func (this *QTime) Second(args ...interface{}) (ret interface{}) {
  // second()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6secondEv
    // invoke: int second()
    var ret0 = C.C_ZNK5QTime6secondEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "second", args)
  }

  return
}

// msecsSinceStartOfDay()
func (this *QTime) Msecssincestartofday(args ...interface{}) (ret interface{}) {
  // msecsSinceStartOfDay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime20msecsSinceStartOfDayEv
    // invoke: int msecsSinceStartOfDay()
    var ret0 = C.C_ZNK5QTime20msecsSinceStartOfDayEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "msecsSinceStartOfDay", args)
  }

  return
}

// setHMS(int, int, int, int)
func (this *QTime) Sethms(args ...interface{}) (ret interface{}) {
  // setHMS(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime6setHMSEiiii
    // invoke: bool setHMS(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN5QTime6setHMSEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "setHMS", args)
  }

  return
}

// addMSecs(int)
func (this *QTime) Addmsecs(args ...interface{}) (ret interface{}) {
  // addMSecs(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime8addMSecsEi
    // invoke: QTime addMSecs(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QTime8addMSecsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "addMSecs", args)
  }

  return
}

// start()
func (this *QTime) Start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime5startEv
    // invoke: void start()
    C.C_ZN5QTime5startEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "start", args)
  }

  return
}

// fromMSecsSinceStartOfDay(int)
func (this *QTime) Frommsecssincestartofday_S(args ...interface{}) (ret interface{}) {
  // fromMSecsSinceStartOfDay(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime24fromMSecsSinceStartOfDayEi
    // invoke: QTime fromMSecsSinceStartOfDay(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QTime24fromMSecsSinceStartOfDayEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "fromMSecsSinceStartOfDay", args)
  }

  return
}

// minute()
func (this *QTime) Minute(args ...interface{}) (ret interface{}) {
  // minute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6minuteEv
    // invoke: int minute()
    var ret0 = C.C_ZNK5QTime6minuteEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "minute", args)
  }

  return
}

// toString(const class QString &)
func (this *QTime) Tostring(args ...interface{}) (ret interface{}) {
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QTime8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "toString", args)
  }

  return
}

// isValid(int, int, int, int)
func (this *QTime) Isvalid_S(args ...interface{}) (ret interface{}) {
  // isValid(int, int, int, int)
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime7isValidEiiii
    // invoke: bool isValid(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN5QTime7isValidEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK5QTime7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK5QTime7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "isValid", args)
  }

  return
}

// secsTo(const class QTime &)
func (this *QTime) Secsto(args ...interface{}) (ret interface{}) {
  // secsTo(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6secsToERKS_
    // invoke: int secsTo(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QTime6secsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "secsTo", args)
  }

  return
}

// elapsed()
func (this *QTime) Elapsed(args ...interface{}) (ret interface{}) {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7elapsedEv
    // invoke: int elapsed()
    var ret0 = C.C_ZNK5QTime7elapsedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "elapsed", args)
  }

  return
}

// addSecs(int)
func (this *QTime) Addsecs(args ...interface{}) (ret interface{}) {
  // addSecs(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7addSecsEi
    // invoke: QTime addSecs(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QTime7addSecsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "addSecs", args)
  }

  return
}

// msecsTo(const class QTime &)
func (this *QTime) Msecsto(args ...interface{}) (ret interface{}) {
  // msecsTo(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7msecsToERKS_
    // invoke: int msecsTo(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QTime7msecsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "msecsTo", args)
  }

  return
}

// restart()
func (this *QTime) Restart(args ...interface{}) (ret interface{}) {
  // restart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime7restartEv
    // invoke: int restart()
    var ret0 = C.C_ZN5QTime7restartEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "restart", args)
  }

  return
}

// hour()
func (this *QTime) Hour(args ...interface{}) (ret interface{}) {
  // hour()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime4hourEv
    // invoke: int hour()
    var ret0 = C.C_ZNK5QTime4hourEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "hour", args)
  }

  return
}

// currentTime()
func (this *QTime) Currenttime_S(args ...interface{}) (ret interface{}) {
  // currentTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime11currentTimeEv
    // invoke: QTime currentTime()
    var ret0 = C.C_ZN5QTime11currentTimeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "currentTime", args)
  }

  return
}

// isNull()
func (this *QTime) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK5QTime6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTime", "isNull", args)
  }

  return
}

// fromString(const class QString &, const class QString &)
func (this *QDateTime) Fromstring_S(args ...interface{}) (ret interface{}) {
  // fromString(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime10fromStringERK7QStringS2_
    // invoke: QDateTime fromString(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QDateTime10fromStringERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "fromString", args)
  }

  return
}

// setOffsetFromUtc(int)
func (this *QDateTime) Setoffsetfromutc(args ...interface{}) () {
  // setOffsetFromUtc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime16setOffsetFromUtcEi
    // invoke: void setOffsetFromUtc(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime16setOffsetFromUtcEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setOffsetFromUtc", args)
  }

  return
}

// toTime_t()
func (this *QDateTime) Totime_T(args ...interface{}) (ret interface{}) {
  // toTime_t()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8toTime_tEv
    // invoke: uint toTime_t()
    var ret0 = C.C_ZNK9QDateTime8toTime_tEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toTime_t", args)
  }

  return
}

// toTimeZone(const class QTimeZone &)
func (this *QDateTime) Totimezone(args ...interface{}) (ret interface{}) {
  // toTimeZone(const class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime10toTimeZoneERK9QTimeZone
    // invoke: QDateTime toTimeZone(const class QTimeZone &)
    var arg0 = args[0].(QTimeZone).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime10toTimeZoneERK9QTimeZone(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toTimeZone", args)
  }

  return
}

// toUTC()
func (this *QDateTime) Toutc(args ...interface{}) (ret interface{}) {
  // toUTC()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime5toUTCEv
    // invoke: QDateTime toUTC()
    var ret0 = C.C_ZNK9QDateTime5toUTCEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toUTC", args)
  }

  return
}

// setUtcOffset(int)
func (this *QDateTime) Setutcoffset(args ...interface{}) () {
  // setUtcOffset(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime12setUtcOffsetEi
    // invoke: void setUtcOffset(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime12setUtcOffsetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setUtcOffset", args)
  }

  return
}

// setTimeZone(const class QTimeZone &)
func (this *QDateTime) Settimezone(args ...interface{}) () {
  // setTimeZone(const class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime11setTimeZoneERK9QTimeZone
    // invoke: void setTimeZone(const class QTimeZone &)
    var arg0 = args[0].(QTimeZone).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime11setTimeZoneERK9QTimeZone(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setTimeZone", args)
  }

  return
}

// QDateTime(const class QDateTime &)
func NewQDateTime(args ...interface{}) *QDateTime {
  // QDateTime(const class QDateTime &)
  // QDateTime(const class QDate &)
  // QDateTime()
  // QDateTime(const class QDate &, const class QTime &, const class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[3][1] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[3][2] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTimeC1ERKS_
    // invoke: void QDateTime(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDateTimeC2ERKS_(arg0)
    return &QDateTime{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QDateTimeC1ERK5QDate
    // invoke: void QDateTime(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDateTimeC2ERK5QDate(arg0)
    return &QDateTime{qclsinst:qthis}
  case 2:
    // invoke: _ZN9QDateTimeC1Ev
    // invoke: void QDateTime()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDateTimeC2Ev()
    return &QDateTime{qclsinst:qthis}
  case 3:
    // invoke: _ZN9QDateTimeC1ERK5QDateRK5QTimeRK9QTimeZone
    // invoke: void QDateTime(const class QDate &, const class QTime &, const class QTimeZone &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTime).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTimeZone).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDateTimeC2ERK5QDateRK5QTimeRK9QTimeZone(arg0, arg1, arg2)
    return &QDateTime{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDateTime", "QDateTime", args)
  }

  return nil // QDateTime{qclsinst:qthis}
}

// ~QDateTime()
func (this *QDateTime) Freeqdatetime(args ...interface{}) () {
  // ~QDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTimeD0Ev
    // invoke: void ~QDateTime()
    C.C_ZN9QDateTimeD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "~QDateTime", args)
  }

  return
}

// toOffsetFromUtc(int)
func (this *QDateTime) Tooffsetfromutc(args ...interface{}) (ret interface{}) {
  // toOffsetFromUtc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime15toOffsetFromUtcEi
    // invoke: QDateTime toOffsetFromUtc(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime15toOffsetFromUtcEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toOffsetFromUtc", args)
  }

  return
}

// addMSecs(qint64)
func (this *QDateTime) Addmsecs(args ...interface{}) (ret interface{}) {
  // addMSecs(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8addMSecsEx
    // invoke: QDateTime addMSecs(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime8addMSecsEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "addMSecs", args)
  }

  return
}

// addMonths(int)
func (this *QDateTime) Addmonths(args ...interface{}) (ret interface{}) {
  // addMonths(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime9addMonthsEi
    // invoke: QDateTime addMonths(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime9addMonthsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "addMonths", args)
  }

  return
}

// offsetFromUtc()
func (this *QDateTime) Offsetfromutc(args ...interface{}) (ret interface{}) {
  // offsetFromUtc()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime13offsetFromUtcEv
    // invoke: int offsetFromUtc()
    var ret0 = C.C_ZNK9QDateTime13offsetFromUtcEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "offsetFromUtc", args)
  }

  return
}

// toString(const class QString &)
func (this *QDateTime) Tostring(args ...interface{}) (ret interface{}) {
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toString", args)
  }

  return
}

// swap(class QDateTime &)
func (this *QDateTime) Swap(args ...interface{}) () {
  // swap(class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime4swapERS_
    // invoke: void swap(class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "swap", args)
  }

  return
}

// setDate(const class QDate &)
func (this *QDateTime) Setdate(args ...interface{}) () {
  // setDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime7setDateERK5QDate
    // invoke: void setDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime7setDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setDate", args)
  }

  return
}

// currentDateTime()
func (this *QDateTime) Currentdatetime_S(args ...interface{}) (ret interface{}) {
  // currentDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime15currentDateTimeEv
    // invoke: QDateTime currentDateTime()
    var ret0 = C.C_ZN9QDateTime15currentDateTimeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTime", args)
  }

  return
}

// currentDateTimeUtc()
func (this *QDateTime) Currentdatetimeutc_S(args ...interface{}) (ret interface{}) {
  // currentDateTimeUtc()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime18currentDateTimeUtcEv
    // invoke: QDateTime currentDateTimeUtc()
    var ret0 = C.C_ZN9QDateTime18currentDateTimeUtcEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTimeUtc", args)
  }

  return
}

// currentMSecsSinceEpoch()
func (this *QDateTime) Currentmsecssinceepoch_S(args ...interface{}) (ret interface{}) {
  // currentMSecsSinceEpoch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime22currentMSecsSinceEpochEv
    // invoke: qint64 currentMSecsSinceEpoch()
    var ret0 = C.C_ZN9QDateTime22currentMSecsSinceEpochEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "currentMSecsSinceEpoch", args)
  }

  return
}

// isValid()
func (this *QDateTime) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK9QDateTime7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "isValid", args)
  }

  return
}

// secsTo(const class QDateTime &)
func (this *QDateTime) Secsto(args ...interface{}) (ret interface{}) {
  // secsTo(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime6secsToERKS_
    // invoke: qint64 secsTo(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime6secsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "secsTo", args)
  }

  return
}

// setMSecsSinceEpoch(qint64)
func (this *QDateTime) Setmsecssinceepoch(args ...interface{}) () {
  // setMSecsSinceEpoch(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime18setMSecsSinceEpochEx
    // invoke: void setMSecsSinceEpoch(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime18setMSecsSinceEpochEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setMSecsSinceEpoch", args)
  }

  return
}

// addDays(qint64)
func (this *QDateTime) Adddays(args ...interface{}) (ret interface{}) {
  // addDays(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7addDaysEx
    // invoke: QDateTime addDays(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime7addDaysEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "addDays", args)
  }

  return
}

// setTime_t(uint)
func (this *QDateTime) Settime_T(args ...interface{}) () {
  // setTime_t(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime9setTime_tEj
    // invoke: void setTime_t(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime9setTime_tEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setTime_t", args)
  }

  return
}

// addSecs(qint64)
func (this *QDateTime) Addsecs(args ...interface{}) (ret interface{}) {
  // addSecs(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7addSecsEx
    // invoke: QDateTime addSecs(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime7addSecsEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "addSecs", args)
  }

  return
}

// timeZoneAbbreviation()
func (this *QDateTime) Timezoneabbreviation(args ...interface{}) (ret interface{}) {
  // timeZoneAbbreviation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime20timeZoneAbbreviationEv
    // invoke: QString timeZoneAbbreviation()
    var ret0 = C.C_ZNK9QDateTime20timeZoneAbbreviationEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "timeZoneAbbreviation", args)
  }

  return
}

// setTime(const class QTime &)
func (this *QDateTime) Settime(args ...interface{}) () {
  // setTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime7setTimeERK5QTime
    // invoke: void setTime(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDateTime7setTimeERK5QTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setTime", args)
  }

  return
}

// toMSecsSinceEpoch()
func (this *QDateTime) Tomsecssinceepoch(args ...interface{}) (ret interface{}) {
  // toMSecsSinceEpoch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime17toMSecsSinceEpochEv
    // invoke: qint64 toMSecsSinceEpoch()
    var ret0 = C.C_ZNK9QDateTime17toMSecsSinceEpochEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toMSecsSinceEpoch", args)
  }

  return
}

// msecsTo(const class QDateTime &)
func (this *QDateTime) Msecsto(args ...interface{}) (ret interface{}) {
  // msecsTo(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7msecsToERKS_
    // invoke: qint64 msecsTo(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime7msecsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "msecsTo", args)
  }

  return
}

// date()
func (this *QDateTime) Date(args ...interface{}) (ret interface{}) {
  // date()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime4dateEv
    // invoke: QDate date()
    var ret0 = C.C_ZNK9QDateTime4dateEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "date", args)
  }

  return
}

// daysTo(const class QDateTime &)
func (this *QDateTime) Daysto(args ...interface{}) (ret interface{}) {
  // daysTo(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime6daysToERKS_
    // invoke: qint64 daysTo(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime6daysToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "daysTo", args)
  }

  return
}

// toLocalTime()
func (this *QDateTime) Tolocaltime(args ...interface{}) (ret interface{}) {
  // toLocalTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime11toLocalTimeEv
    // invoke: QDateTime toLocalTime()
    var ret0 = C.C_ZNK9QDateTime11toLocalTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "toLocalTime", args)
  }

  return
}

// timeSpec()
func (this *QDateTime) Timespec(args ...interface{}) () {
  // timeSpec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8timeSpecEv
    // invoke: Qt::TimeSpec timeSpec()
    C.C_ZNK9QDateTime8timeSpecEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "timeSpec", args)
  }

  return
}

// fromTime_t(uint, const class QTimeZone &)
func (this *QDateTime) Fromtime_T_S(args ...interface{}) (ret interface{}) {
  // fromTime_t(uint, const class QTimeZone &)
  // fromTime_t(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[0][1] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime10fromTime_tEjRK9QTimeZone
    // invoke: QDateTime fromTime_t(uint, const class QTimeZone &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTimeZone).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QDateTime10fromTime_tEjRK9QTimeZone(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN9QDateTime10fromTime_tEj
    // invoke: QDateTime fromTime_t(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QDateTime10fromTime_tEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "fromTime_t", args)
  }

  return
}

// fromMSecsSinceEpoch(qint64)
func (this *QDateTime) Frommsecssinceepoch_S(args ...interface{}) (ret interface{}) {
  // fromMSecsSinceEpoch(qint64)
  // fromMSecsSinceEpoch(qint64, const class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qint64"
  vtys[1][1] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime19fromMSecsSinceEpochEx
    // invoke: QDateTime fromMSecsSinceEpoch(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QDateTime19fromMSecsSinceEpochEx(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone
    // invoke: QDateTime fromMSecsSinceEpoch(qint64, const class QTimeZone &)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTimeZone).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "fromMSecsSinceEpoch", args)
  }

  return
}

// isNull()
func (this *QDateTime) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK9QDateTime6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "isNull", args)
  }

  return
}

// utcOffset()
func (this *QDateTime) Utcoffset(args ...interface{}) (ret interface{}) {
  // utcOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime9utcOffsetEv
    // invoke: int utcOffset()
    var ret0 = C.C_ZNK9QDateTime9utcOffsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "utcOffset", args)
  }

  return
}

// time()
func (this *QDateTime) Time(args ...interface{}) (ret interface{}) {
  // time()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime4timeEv
    // invoke: QTime time()
    var ret0 = C.C_ZNK9QDateTime4timeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "time", args)
  }

  return
}

// timeZone()
func (this *QDateTime) Timezone(args ...interface{}) (ret interface{}) {
  // timeZone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8timeZoneEv
    // invoke: QTimeZone timeZone()
    var ret0 = C.C_ZNK9QDateTime8timeZoneEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTimeZone{}) // "QTimeZone"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "timeZone", args)
  }

  return
}

// isDaylightTime()
func (this *QDateTime) Isdaylighttime(args ...interface{}) (ret interface{}) {
  // isDaylightTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime14isDaylightTimeEv
    // invoke: bool isDaylightTime()
    var ret0 = C.C_ZNK9QDateTime14isDaylightTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "isDaylightTime", args)
  }

  return
}

// addYears(int)
func (this *QDateTime) Addyears(args ...interface{}) (ret interface{}) {
  // addYears(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8addYearsEi
    // invoke: QDateTime addYears(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDateTime8addYearsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDateTime", "addYears", args)
  }

  return
}

// fromString(const class QString &, const class QString &)
func (this *QDate) Fromstring_S(args ...interface{}) (ret interface{}) {
  // fromString(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate10fromStringERK7QStringS2_
    // invoke: QDate fromString(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QDate10fromStringERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "fromString", args)
  }

  return
}

// month()
func (this *QDate) Month(args ...interface{}) (ret interface{}) {
  // month()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate5monthEv
    // invoke: int month()
    var ret0 = C.C_ZNK5QDate5monthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "month", args)
  }

  return
}

// year()
func (this *QDate) Year(args ...interface{}) (ret interface{}) {
  // year()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate4yearEv
    // invoke: int year()
    var ret0 = C.C_ZNK5QDate4yearEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "year", args)
  }

  return
}

// daysInMonth()
func (this *QDate) Daysinmonth(args ...interface{}) (ret interface{}) {
  // daysInMonth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate11daysInMonthEv
    // invoke: int daysInMonth()
    var ret0 = C.C_ZNK5QDate11daysInMonthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "daysInMonth", args)
  }

  return
}

// daysTo(const class QDate &)
func (this *QDate) Daysto(args ...interface{}) (ret interface{}) {
  // daysTo(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate6daysToERKS_
    // invoke: qint64 daysTo(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QDate6daysToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "daysTo", args)
  }

  return
}

// weekNumber(int *)
func (this *QDate) Weeknumber(args ...interface{}) (ret interface{}) {
  // weekNumber(int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate10weekNumberEPi
    // invoke: int weekNumber(int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QDate10weekNumberEPi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "weekNumber", args)
  }

  return
}

// daysInYear()
func (this *QDate) Daysinyear(args ...interface{}) (ret interface{}) {
  // daysInYear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate10daysInYearEv
    // invoke: int daysInYear()
    var ret0 = C.C_ZNK5QDate10daysInYearEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "daysInYear", args)
  }

  return
}

// addMonths(int)
func (this *QDate) Addmonths(args ...interface{}) (ret interface{}) {
  // addMonths(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate9addMonthsEi
    // invoke: QDate addMonths(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QDate9addMonthsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "addMonths", args)
  }

  return
}

// toJulianDay()
func (this *QDate) Tojulianday(args ...interface{}) (ret interface{}) {
  // toJulianDay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate11toJulianDayEv
    // invoke: qint64 toJulianDay()
    var ret0 = C.C_ZNK5QDate11toJulianDayEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "toJulianDay", args)
  }

  return
}

// toString(const class QString &)
func (this *QDate) Tostring(args ...interface{}) (ret interface{}) {
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QDate8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "toString", args)
  }

  return
}

// fromJulianDay(qint64)
func (this *QDate) Fromjulianday_S(args ...interface{}) (ret interface{}) {
  // fromJulianDay(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate13fromJulianDayEx
    // invoke: QDate fromJulianDay(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QDate13fromJulianDayEx(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "fromJulianDay", args)
  }

  return
}

// setDate(int, int, int)
func (this *QDate) Setdate(args ...interface{}) (ret interface{}) {
  // setDate(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate7setDateEiii
    // invoke: bool setDate(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN5QDate7setDateEiii(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "setDate", args)
  }

  return
}

// getDate(int *, int *, int *)
func (this *QDate) Getdate(args ...interface{}) () {
  // getDate(int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate7getDateEPiS0_S0_
    // invoke: void getDate(int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C.C_ZN5QDate7getDateEPiS0_S0_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDate", "getDate", args)
  }

  return
}

// dayOfWeek()
func (this *QDate) Dayofweek(args ...interface{}) (ret interface{}) {
  // dayOfWeek()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate9dayOfWeekEv
    // invoke: int dayOfWeek()
    var ret0 = C.C_ZNK5QDate9dayOfWeekEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "dayOfWeek", args)
  }

  return
}

// isLeapYear(int)
func (this *QDate) Isleapyear_S(args ...interface{}) (ret interface{}) {
  // isLeapYear(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate10isLeapYearEi
    // invoke: bool isLeapYear(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QDate10isLeapYearEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "isLeapYear", args)
  }

  return
}

// QDate()
func NewQDate(args ...interface{}) *QDate {
  // QDate()
  // QDate(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDateC1Ev
    // invoke: void QDate()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QDateC2Ev()
    return &QDate{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QDateC1Eiii
    // invoke: void QDate(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QDateC2Eiii(arg0, arg1, arg2)
    return &QDate{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDate", "QDate", args)
  }

  return nil // QDate{qclsinst:qthis}
}

// isValid()
func (this *QDate) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  // isValid(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK5QDate7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN5QDate7isValidEiii
    // invoke: bool isValid(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN5QDate7isValidEiii(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "isValid", args)
  }

  return
}

// addDays(qint64)
func (this *QDate) Adddays(args ...interface{}) (ret interface{}) {
  // addDays(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate7addDaysEx
    // invoke: QDate addDays(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QDate7addDaysEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "addDays", args)
  }

  return
}

// day()
func (this *QDate) Day(args ...interface{}) (ret interface{}) {
  // day()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate3dayEv
    // invoke: int day()
    var ret0 = C.C_ZNK5QDate3dayEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "day", args)
  }

  return
}

// addYears(int)
func (this *QDate) Addyears(args ...interface{}) (ret interface{}) {
  // addYears(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate8addYearsEi
    // invoke: QDate addYears(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QDate8addYearsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "addYears", args)
  }

  return
}

// isNull()
func (this *QDate) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK5QDate6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "isNull", args)
  }

  return
}

// dayOfYear()
func (this *QDate) Dayofyear(args ...interface{}) (ret interface{}) {
  // dayOfYear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate9dayOfYearEv
    // invoke: int dayOfYear()
    var ret0 = C.C_ZNK5QDate9dayOfYearEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "dayOfYear", args)
  }

  return
}

// currentDate()
func (this *QDate) Currentdate_S(args ...interface{}) (ret interface{}) {
  // currentDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate11currentDateEv
    // invoke: QDate currentDate()
    var ret0 = C.C_ZN5QDate11currentDateEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDate", "currentDate", args)
  }

  return
}

// <= body block end

