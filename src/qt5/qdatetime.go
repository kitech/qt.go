package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZN5QTime10fromStringERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  void QTime::QTime(int h, int m, int s, int ms);
extern void* C_ZN5QTimeC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 3
  // proto:  void QTime::QTime();
extern void* C_ZN5QTimeC2Ev(); // 1
  // proto:  int QTime::msec();
extern void C_ZNK5QTime4msecEv(void* qthis); // 4
  // proto:  int QTime::second();
extern void C_ZNK5QTime6secondEv(void* qthis); // 4
  // proto:  int QTime::msecsSinceStartOfDay();
extern void C_ZNK5QTime20msecsSinceStartOfDayEv(void* qthis); // 2
  // proto:  bool QTime::setHMS(int h, int m, int s, int ms);
extern void C_ZN5QTime6setHMSEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QTime QTime::addMSecs(int ms);
extern void C_ZNK5QTime8addMSecsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTime::start();
extern void C_ZN5QTime5startEv(void* qthis); // 4
  // proto: static QTime QTime::fromMSecsSinceStartOfDay(int msecs);
extern void C_ZN5QTime24fromMSecsSinceStartOfDayEi(int32_t arg0); // 2
  // proto:  int QTime::minute();
extern void C_ZNK5QTime6minuteEv(void* qthis); // 4
  // proto:  QString QTime::toString(const QString & format);
extern void C_ZNK5QTime8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto: static bool QTime::isValid(int h, int m, int s, int ms);
extern void C_ZN5QTime7isValidEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  bool QTime::isValid();
extern void C_ZNK5QTime7isValidEv(void* qthis); // 4
  // proto:  int QTime::secsTo(const QTime & );
extern void C_ZNK5QTime6secsToERKS_(void* qthis, void* arg0); // 4
  // proto:  int QTime::elapsed();
extern void C_ZNK5QTime7elapsedEv(void* qthis); // 4
  // proto:  QTime QTime::addSecs(int secs);
extern void C_ZNK5QTime7addSecsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTime::msecsTo(const QTime & );
extern void C_ZNK5QTime7msecsToERKS_(void* qthis, void* arg0); // 4
  // proto:  int QTime::restart();
extern void C_ZN5QTime7restartEv(void* qthis); // 4
  // proto:  int QTime::hour();
extern void C_ZNK5QTime4hourEv(void* qthis); // 4
  // proto: static QTime QTime::currentTime();
extern void C_ZN5QTime11currentTimeEv(); // 4
  // proto:  bool QTime::isNull();
extern void C_ZNK5QTime6isNullEv(void* qthis); // 2
  // proto: static QDateTime QDateTime::fromString(const QString & s, const QString & format);
extern void C_ZN9QDateTime10fromStringERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  void QDateTime::setOffsetFromUtc(int offsetSeconds);
extern void C_ZN9QDateTime16setOffsetFromUtcEi(void* qthis, int32_t arg0); // 4
  // proto:  uint QDateTime::toTime_t();
extern void C_ZNK9QDateTime8toTime_tEv(void* qthis); // 4
  // proto:  QDateTime QDateTime::toTimeZone(const QTimeZone & toZone);
extern void C_ZNK9QDateTime10toTimeZoneERK9QTimeZone(void* qthis, void* arg0); // 4
  // proto:  QDateTime QDateTime::toUTC();
extern void C_ZNK9QDateTime5toUTCEv(void* qthis); // 2
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
extern void C_ZNK9QDateTime15toOffsetFromUtcEi(void* qthis, int32_t arg0); // 4
  // proto:  QDateTime QDateTime::addMSecs(qint64 msecs);
extern void C_ZNK9QDateTime8addMSecsEx(void* qthis, int64_t arg0); // 4
  // proto:  QDateTime QDateTime::addMonths(int months);
extern void C_ZNK9QDateTime9addMonthsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QDateTime::offsetFromUtc();
extern void C_ZNK9QDateTime13offsetFromUtcEv(void* qthis); // 4
  // proto:  QString QDateTime::toString(const QString & format);
extern void C_ZNK9QDateTime8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QDateTime::swap(QDateTime & other);
extern void C_ZN9QDateTime4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QDateTime::setDate(const QDate & date);
extern void C_ZN9QDateTime7setDateERK5QDate(void* qthis, void* arg0); // 4
  // proto: static QDateTime QDateTime::currentDateTime();
extern void C_ZN9QDateTime15currentDateTimeEv(); // 4
  // proto: static QDateTime QDateTime::currentDateTimeUtc();
extern void C_ZN9QDateTime18currentDateTimeUtcEv(); // 4
  // proto: static qint64 QDateTime::currentMSecsSinceEpoch();
extern void C_ZN9QDateTime22currentMSecsSinceEpochEv(); // 4
  // proto:  bool QDateTime::isValid();
extern void C_ZNK9QDateTime7isValidEv(void* qthis); // 4
  // proto:  qint64 QDateTime::secsTo(const QDateTime & );
extern void C_ZNK9QDateTime6secsToERKS_(void* qthis, void* arg0); // 4
  // proto:  void QDateTime::setMSecsSinceEpoch(qint64 msecs);
extern void C_ZN9QDateTime18setMSecsSinceEpochEx(void* qthis, int64_t arg0); // 4
  // proto:  QDateTime QDateTime::addDays(qint64 days);
extern void C_ZNK9QDateTime7addDaysEx(void* qthis, int64_t arg0); // 4
  // proto:  void QDateTime::setTime_t(uint secsSince1Jan1970UTC);
extern void C_ZN9QDateTime9setTime_tEj(void* qthis, int32_t arg0); // 4
  // proto:  QDateTime QDateTime::addSecs(qint64 secs);
extern void C_ZNK9QDateTime7addSecsEx(void* qthis, int64_t arg0); // 4
  // proto:  QString QDateTime::timeZoneAbbreviation();
extern void C_ZNK9QDateTime20timeZoneAbbreviationEv(void* qthis); // 4
  // proto:  void QDateTime::setTime(const QTime & time);
extern void C_ZN9QDateTime7setTimeERK5QTime(void* qthis, void* arg0); // 4
  // proto:  qint64 QDateTime::toMSecsSinceEpoch();
extern void C_ZNK9QDateTime17toMSecsSinceEpochEv(void* qthis); // 4
  // proto:  qint64 QDateTime::msecsTo(const QDateTime & );
extern void C_ZNK9QDateTime7msecsToERKS_(void* qthis, void* arg0); // 4
  // proto:  QDate QDateTime::date();
extern void C_ZNK9QDateTime4dateEv(void* qthis); // 4
  // proto:  qint64 QDateTime::daysTo(const QDateTime & );
extern void C_ZNK9QDateTime6daysToERKS_(void* qthis, void* arg0); // 4
  // proto:  QDateTime QDateTime::toLocalTime();
extern void C_ZNK9QDateTime11toLocalTimeEv(void* qthis); // 2
  // proto:  Qt::TimeSpec QDateTime::timeSpec();
extern void C_ZNK9QDateTime8timeSpecEv(void* qthis); // 4
  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC, const QTimeZone & timeZone);
extern void C_ZN9QDateTime10fromTime_tEjRK9QTimeZone(int32_t arg0, void* arg1); // 4
  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC);
extern void C_ZN9QDateTime10fromTime_tEj(int32_t arg0); // 4
  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs);
extern void C_ZN9QDateTime19fromMSecsSinceEpochEx(int64_t arg0); // 4
  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs, const QTimeZone & timeZone);
extern void C_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone(int64_t arg0, void* arg1); // 4
  // proto:  bool QDateTime::isNull();
extern void C_ZNK9QDateTime6isNullEv(void* qthis); // 4
  // proto:  int QDateTime::utcOffset();
extern void C_ZNK9QDateTime9utcOffsetEv(void* qthis); // 4
  // proto:  QTime QDateTime::time();
extern void C_ZNK9QDateTime4timeEv(void* qthis); // 4
  // proto:  QTimeZone QDateTime::timeZone();
extern void C_ZNK9QDateTime8timeZoneEv(void* qthis); // 4
  // proto:  bool QDateTime::isDaylightTime();
extern void C_ZNK9QDateTime14isDaylightTimeEv(void* qthis); // 4
  // proto:  QDateTime QDateTime::addYears(int years);
extern void C_ZNK9QDateTime8addYearsEi(void* qthis, int32_t arg0); // 4
  // proto: static QDate QDate::fromString(const QString & s, const QString & format);
extern void C_ZN5QDate10fromStringERK7QStringS2_(void* arg0, void* arg1); // 4
  // proto:  int QDate::month();
extern void C_ZNK5QDate5monthEv(void* qthis); // 4
  // proto:  int QDate::year();
extern void C_ZNK5QDate4yearEv(void* qthis); // 4
  // proto:  int QDate::daysInMonth();
extern void C_ZNK5QDate11daysInMonthEv(void* qthis); // 4
  // proto:  qint64 QDate::daysTo(const QDate & );
extern void C_ZNK5QDate6daysToERKS_(void* qthis, void* arg0); // 4
  // proto:  int QDate::weekNumber(int * yearNum);
extern void C_ZNK5QDate10weekNumberEPi(void* qthis, int32_t* arg0); // 4
  // proto:  int QDate::daysInYear();
extern void C_ZNK5QDate10daysInYearEv(void* qthis); // 4
  // proto:  QDate QDate::addMonths(int months);
extern void C_ZNK5QDate9addMonthsEi(void* qthis, int32_t arg0); // 4
  // proto:  qint64 QDate::toJulianDay();
extern void C_ZNK5QDate11toJulianDayEv(void* qthis); // 2
  // proto:  QString QDate::toString(const QString & format);
extern void C_ZNK5QDate8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto: static QDate QDate::fromJulianDay(qint64 jd);
extern void C_ZN5QDate13fromJulianDayEx(int64_t arg0); // 2
  // proto:  bool QDate::setDate(int year, int month, int day);
extern void C_ZN5QDate7setDateEiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QDate::getDate(int * year, int * month, int * day);
extern void C_ZN5QDate7getDateEPiS0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2); // 4
  // proto:  int QDate::dayOfWeek();
extern void C_ZNK5QDate9dayOfWeekEv(void* qthis); // 4
  // proto: static bool QDate::isLeapYear(int year);
extern void C_ZN5QDate10isLeapYearEi(int32_t arg0); // 4
  // proto:  void QDate::QDate();
extern void* C_ZN5QDateC2Ev(); // 1
  // proto:  void QDate::QDate(int y, int m, int d);
extern void* C_ZN5QDateC2Eiii(int32_t arg0, int32_t arg1, int32_t arg2); // 3
  // proto:  bool QDate::isValid();
extern void C_ZNK5QDate7isValidEv(void* qthis); // 2
  // proto: static bool QDate::isValid(int y, int m, int d);
extern void C_ZN5QDate7isValidEiii(int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  QDate QDate::addDays(qint64 days);
extern void C_ZNK5QDate7addDaysEx(void* qthis, int64_t arg0); // 4
  // proto:  int QDate::day();
extern void C_ZNK5QDate3dayEv(void* qthis); // 4
  // proto:  QDate QDate::addYears(int years);
extern void C_ZNK5QDate8addYearsEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QDate::isNull();
extern void C_ZNK5QDate6isNullEv(void* qthis); // 2
  // proto:  int QDate::dayOfYear();
extern void C_ZNK5QDate9dayOfYearEv(void* qthis); // 4
  // proto: static QDate QDate::currentDate();
extern void C_ZN5QDate11currentDateEv(); // 4
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
func (this *QTime) fromString_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QTime10fromStringERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "fromString", args)
  }

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
func (this *QTime) msec(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime4msecEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "msec", args)
  }

}

// second()
func (this *QTime) second(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime6secondEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "second", args)
  }

}

// msecsSinceStartOfDay()
func (this *QTime) msecsSinceStartOfDay(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime20msecsSinceStartOfDayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "msecsSinceStartOfDay", args)
  }

}

// setHMS(int, int, int, int)
func (this *QTime) setHMS(args ...interface{}) () {
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
    var ret = C.C_ZN5QTime6setHMSEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "setHMS", args)
  }

}

// addMSecs(int)
func (this *QTime) addMSecs(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime8addMSecsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "addMSecs", args)
  }

}

// start()
func (this *QTime) start(args ...interface{}) () {
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

}

// fromMSecsSinceStartOfDay(int)
func (this *QTime) fromMSecsSinceStartOfDay_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QTime24fromMSecsSinceStartOfDayEi(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "fromMSecsSinceStartOfDay", args)
  }

}

// minute()
func (this *QTime) minute(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime6minuteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "minute", args)
  }

}

// toString(const class QString &)
func (this *QTime) toString(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "toString", args)
  }

}

// isValid(int, int, int, int)
func (this *QTime) isValid_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QTime7isValidEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK5QTime7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK5QTime7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "isValid", args)
  }

}

// secsTo(const class QTime &)
func (this *QTime) secsTo(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime6secsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "secsTo", args)
  }

}

// elapsed()
func (this *QTime) elapsed(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime7elapsedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "elapsed", args)
  }

}

// addSecs(int)
func (this *QTime) addSecs(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime7addSecsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "addSecs", args)
  }

}

// msecsTo(const class QTime &)
func (this *QTime) msecsTo(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime7msecsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "msecsTo", args)
  }

}

// restart()
func (this *QTime) restart(args ...interface{}) () {
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
    var ret = C.C_ZN5QTime7restartEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "restart", args)
  }

}

// hour()
func (this *QTime) hour(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime4hourEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "hour", args)
  }

}

// currentTime()
func (this *QTime) currentTime_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QTime11currentTimeEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "currentTime", args)
  }

}

// isNull()
func (this *QTime) isNull(args ...interface{}) () {
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
    var ret = C.C_ZNK5QTime6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTime", "isNull", args)
  }

}

// fromString(const class QString &, const class QString &)
func (this *QDateTime) fromString_s(args ...interface{}) () {
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
    var ret = C.C_ZN9QDateTime10fromStringERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "fromString", args)
  }

}

// setOffsetFromUtc(int)
func (this *QDateTime) setOffsetFromUtc(args ...interface{}) () {
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

}

// toTime_t()
func (this *QDateTime) toTime_t(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime8toTime_tEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toTime_t", args)
  }

}

// toTimeZone(const class QTimeZone &)
func (this *QDateTime) toTimeZone(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime10toTimeZoneERK9QTimeZone(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toTimeZone", args)
  }

}

// toUTC()
func (this *QDateTime) toUTC(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime5toUTCEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toUTC", args)
  }

}

// setUtcOffset(int)
func (this *QDateTime) setUtcOffset(args ...interface{}) () {
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

}

// setTimeZone(const class QTimeZone &)
func (this *QDateTime) setTimeZone(args ...interface{}) () {
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
func (this *QDateTime) FreeQDateTime(args ...interface{}) () {
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

}

// toOffsetFromUtc(int)
func (this *QDateTime) toOffsetFromUtc(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime15toOffsetFromUtcEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toOffsetFromUtc", args)
  }

}

// addMSecs(qint64)
func (this *QDateTime) addMSecs(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime8addMSecsEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "addMSecs", args)
  }

}

// addMonths(int)
func (this *QDateTime) addMonths(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime9addMonthsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "addMonths", args)
  }

}

// offsetFromUtc()
func (this *QDateTime) offsetFromUtc(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime13offsetFromUtcEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "offsetFromUtc", args)
  }

}

// toString(const class QString &)
func (this *QDateTime) toString(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toString", args)
  }

}

// swap(class QDateTime &)
func (this *QDateTime) swap(args ...interface{}) () {
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

}

// setDate(const class QDate &)
func (this *QDateTime) setDate(args ...interface{}) () {
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

}

// currentDateTime()
func (this *QDateTime) currentDateTime_s(args ...interface{}) () {
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
    var ret = C.C_ZN9QDateTime15currentDateTimeEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTime", args)
  }

}

// currentDateTimeUtc()
func (this *QDateTime) currentDateTimeUtc_s(args ...interface{}) () {
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
    var ret = C.C_ZN9QDateTime18currentDateTimeUtcEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTimeUtc", args)
  }

}

// currentMSecsSinceEpoch()
func (this *QDateTime) currentMSecsSinceEpoch_s(args ...interface{}) () {
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
    var ret = C.C_ZN9QDateTime22currentMSecsSinceEpochEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "currentMSecsSinceEpoch", args)
  }

}

// isValid()
func (this *QDateTime) isValid(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "isValid", args)
  }

}

// secsTo(const class QDateTime &)
func (this *QDateTime) secsTo(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime6secsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "secsTo", args)
  }

}

// setMSecsSinceEpoch(qint64)
func (this *QDateTime) setMSecsSinceEpoch(args ...interface{}) () {
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

}

// addDays(qint64)
func (this *QDateTime) addDays(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime7addDaysEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "addDays", args)
  }

}

// setTime_t(uint)
func (this *QDateTime) setTime_t(args ...interface{}) () {
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

}

// addSecs(qint64)
func (this *QDateTime) addSecs(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime7addSecsEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "addSecs", args)
  }

}

// timeZoneAbbreviation()
func (this *QDateTime) timeZoneAbbreviation(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime20timeZoneAbbreviationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "timeZoneAbbreviation", args)
  }

}

// setTime(const class QTime &)
func (this *QDateTime) setTime(args ...interface{}) () {
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

}

// toMSecsSinceEpoch()
func (this *QDateTime) toMSecsSinceEpoch(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime17toMSecsSinceEpochEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toMSecsSinceEpoch", args)
  }

}

// msecsTo(const class QDateTime &)
func (this *QDateTime) msecsTo(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime7msecsToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "msecsTo", args)
  }

}

// date()
func (this *QDateTime) date(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime4dateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "date", args)
  }

}

// daysTo(const class QDateTime &)
func (this *QDateTime) daysTo(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime6daysToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "daysTo", args)
  }

}

// toLocalTime()
func (this *QDateTime) toLocalTime(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime11toLocalTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "toLocalTime", args)
  }

}

// timeSpec()
func (this *QDateTime) timeSpec(args ...interface{}) () {
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

}

// fromTime_t(uint, const class QTimeZone &)
func (this *QDateTime) fromTime_t_s(args ...interface{}) () {
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
    var ret = C.C_ZN9QDateTime10fromTime_tEjRK9QTimeZone(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QDateTime10fromTime_tEj
    // invoke: QDateTime fromTime_t(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QDateTime10fromTime_tEj(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "fromTime_t", args)
  }

}

// fromMSecsSinceEpoch(qint64)
func (this *QDateTime) fromMSecsSinceEpoch_s(args ...interface{}) () {
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
    var ret = C.C_ZN9QDateTime19fromMSecsSinceEpochEx(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone
    // invoke: QDateTime fromMSecsSinceEpoch(qint64, const class QTimeZone &)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTimeZone).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "fromMSecsSinceEpoch", args)
  }

}

// isNull()
func (this *QDateTime) isNull(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "isNull", args)
  }

}

// utcOffset()
func (this *QDateTime) utcOffset(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime9utcOffsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "utcOffset", args)
  }

}

// time()
func (this *QDateTime) time(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime4timeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "time", args)
  }

}

// timeZone()
func (this *QDateTime) timeZone(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime8timeZoneEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "timeZone", args)
  }

}

// isDaylightTime()
func (this *QDateTime) isDaylightTime(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime14isDaylightTimeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "isDaylightTime", args)
  }

}

// addYears(int)
func (this *QDateTime) addYears(args ...interface{}) () {
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
    var ret = C.C_ZNK9QDateTime8addYearsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDateTime", "addYears", args)
  }

}

// fromString(const class QString &, const class QString &)
func (this *QDate) fromString_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QDate10fromStringERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "fromString", args)
  }

}

// month()
func (this *QDate) month(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate5monthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "month", args)
  }

}

// year()
func (this *QDate) year(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate4yearEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "year", args)
  }

}

// daysInMonth()
func (this *QDate) daysInMonth(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate11daysInMonthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "daysInMonth", args)
  }

}

// daysTo(const class QDate &)
func (this *QDate) daysTo(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate6daysToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "daysTo", args)
  }

}

// weekNumber(int *)
func (this *QDate) weekNumber(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK5QDate10weekNumberEPi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "weekNumber", args)
  }

}

// daysInYear()
func (this *QDate) daysInYear(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate10daysInYearEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "daysInYear", args)
  }

}

// addMonths(int)
func (this *QDate) addMonths(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate9addMonthsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "addMonths", args)
  }

}

// toJulianDay()
func (this *QDate) toJulianDay(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate11toJulianDayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "toJulianDay", args)
  }

}

// toString(const class QString &)
func (this *QDate) toString(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "toString", args)
  }

}

// fromJulianDay(qint64)
func (this *QDate) fromJulianDay_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QDate13fromJulianDayEx(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "fromJulianDay", args)
  }

}

// setDate(int, int, int)
func (this *QDate) setDate(args ...interface{}) () {
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
    var ret = C.C_ZN5QDate7setDateEiii(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "setDate", args)
  }

}

// getDate(int *, int *, int *)
func (this *QDate) getDate(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C.C_ZN5QDate7getDateEPiS0_S0_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDate", "getDate", args)
  }

}

// dayOfWeek()
func (this *QDate) dayOfWeek(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate9dayOfWeekEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "dayOfWeek", args)
  }

}

// isLeapYear(int)
func (this *QDate) isLeapYear_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QDate10isLeapYearEi(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "isLeapYear", args)
  }

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
func (this *QDate) isValid(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QDate7isValidEiii
    // invoke: bool isValid(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN5QDate7isValidEiii(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "isValid", args)
  }

}

// addDays(qint64)
func (this *QDate) addDays(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate7addDaysEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "addDays", args)
  }

}

// day()
func (this *QDate) day(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate3dayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "day", args)
  }

}

// addYears(int)
func (this *QDate) addYears(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate8addYearsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "addYears", args)
  }

}

// isNull()
func (this *QDate) isNull(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "isNull", args)
  }

}

// dayOfYear()
func (this *QDate) dayOfYear(args ...interface{}) () {
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
    var ret = C.C_ZNK5QDate9dayOfYearEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "dayOfYear", args)
  }

}

// currentDate()
func (this *QDate) currentDate_s(args ...interface{}) () {
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
    var ret = C.C_ZN5QDate11currentDateEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDate", "currentDate", args)
  }

}

// <= body block end

