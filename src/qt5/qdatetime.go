package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  QTime QTime::addMSecs(int ms);
extern void _ZNK5QTime8addMSecsEi(void* qthis, int arg0);
  // proto: static QTime QTime::fromMSecsSinceStartOfDay(int msecs);
extern void demth_ZN5QTime24fromMSecsSinceStartOfDayEi(int arg0);
  // proto: static QTime QTime::currentTime();
extern void _ZN5QTime11currentTimeEv();
  // proto:  int QTime::second();
extern void _ZNK5QTime6secondEv(void* qthis);
  // proto:  int QTime::restart();
extern void _ZN5QTime7restartEv(void* qthis);
  // proto:  void QTime::start();
extern void _ZN5QTime5startEv(void* qthis);
  // proto:  bool QTime::isNull();
extern void _ZNK5QTime6isNullEv(void* qthis);
  // proto:  int QTime::msecsSinceStartOfDay();
extern void _ZNK5QTime20msecsSinceStartOfDayEv(void* qthis);
  // proto:  int QTime::hour();
extern void _ZNK5QTime4hourEv(void* qthis);
  // proto:  int QTime::elapsed();
extern void _ZNK5QTime7elapsedEv(void* qthis);
  // proto:  QTime QTime::addSecs(int secs);
extern void _ZNK5QTime7addSecsEi(void* qthis, int arg0);
  // proto:  bool QTime::isValid();
extern void _ZNK5QTime7isValidEv(void* qthis);
  // proto:  void QTime::QTime(int ms);
extern void* dector_ZN5QTimeC1Ei(int arg0);
extern void _ZN5QTimeC1Ei(void* qthis, int arg0);
  // proto:  int QTime::msec();
extern void _ZNK5QTime4msecEv(void* qthis);
  // proto:  void QTime::QTime(int h, int m, int s, int ms);
extern void* dector_ZN5QTimeC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN5QTimeC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  int QTime::secsTo(const QTime & );
extern void _ZNK5QTime6secsToERKS_(void* qthis, void* arg0);
  // proto:  void QTime::QTime();
extern void* dector_ZN5QTimeC1Ev();
extern void _ZN5QTimeC1Ev(void* qthis);
  // proto:  bool QTime::setHMS(int h, int m, int s, int ms);
extern void _ZN5QTime6setHMSEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  QString QTime::toString(const QString & format);
extern void _ZNK5QTime8toStringERK7QString(void* qthis, void* arg0);
  // proto:  int QTime::msecsTo(const QTime & );
extern void _ZNK5QTime7msecsToERKS_(void* qthis, void* arg0);
  // proto:  int QTime::minute();
extern void _ZNK5QTime6minuteEv(void* qthis);
  // proto: static bool QTime::isValid(int h, int m, int s, int ms);
extern void _ZN5QTime7isValidEiiii(int arg0, int arg1, int arg2, int arg3);
  // proto: static QTime QTime::fromString(const QString & s, const QString & format);
extern void _ZN5QTime10fromStringERK7QStringS2_(void* arg0, void* arg1);
  // proto:  QDateTime QDateTime::toLocalTime();
extern void demth_ZNK9QDateTime11toLocalTimeEv(void* qthis);
  // proto:  void QDateTime::setOffsetFromUtc(int offsetSeconds);
extern void _ZN9QDateTime16setOffsetFromUtcEi(void* qthis, int arg0);
  // proto:  QTimeZone QDateTime::timeZone();
extern void _ZNK9QDateTime8timeZoneEv(void* qthis);
  // proto:  void QDateTime::setTime(const QTime & time);
extern void _ZN9QDateTime7setTimeERK5QTime(void* qthis, void* arg0);
  // proto:  qint64 QDateTime::toMSecsSinceEpoch();
extern void _ZNK9QDateTime17toMSecsSinceEpochEv(void* qthis);
  // proto:  void QDateTime::setTime_t(uint secsSince1Jan1970UTC);
extern void _ZN9QDateTime9setTime_tEj(void* qthis, unsigned int arg0);
  // proto:  void QDateTime::QDateTime(const QDateTime & other);
extern void* dector_ZN9QDateTimeC1ERKS_(void* arg0);
extern void _ZN9QDateTimeC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDateTime::QDateTime();
extern void* dector_ZN9QDateTimeC1Ev();
extern void _ZN9QDateTimeC1Ev(void* qthis);
  // proto:  bool QDateTime::isDaylightTime();
extern void _ZNK9QDateTime14isDaylightTimeEv(void* qthis);
  // proto:  bool QDateTime::isValid();
extern void _ZNK9QDateTime7isValidEv(void* qthis);
  // proto:  QString QDateTime::toString(const QString & format);
extern void _ZNK9QDateTime8toStringERK7QString(void* qthis, void* arg0);
  // proto:  QDateTime QDateTime::addYears(int years);
extern void _ZNK9QDateTime8addYearsEi(void* qthis, int arg0);
  // proto:  void QDateTime::setMSecsSinceEpoch(qint64 msecs);
extern void _ZN9QDateTime18setMSecsSinceEpochEx(void* qthis, long long arg0);
  // proto:  QDateTime QDateTime::toOffsetFromUtc(int offsetSeconds);
extern void _ZNK9QDateTime15toOffsetFromUtcEi(void* qthis, int arg0);
  // proto:  void QDateTime::setUtcOffset(int seconds);
extern void _ZN9QDateTime12setUtcOffsetEi(void* qthis, int arg0);
  // proto:  QDateTime QDateTime::addSecs(qint64 secs);
extern void _ZNK9QDateTime7addSecsEx(void* qthis, long long arg0);
  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs);
extern void _ZN9QDateTime19fromMSecsSinceEpochEx(long long arg0);
  // proto:  void QDateTime::QDateTime(const QDate & date, const QTime & time, const QTimeZone & timeZone);
extern void* dector_ZN9QDateTimeC1ERK5QDateRK5QTimeRK9QTimeZone(void* arg0, void* arg1, void* arg2);
extern void _ZN9QDateTimeC1ERK5QDateRK5QTimeRK9QTimeZone(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto: static QDateTime QDateTime::fromString(const QString & s, const QString & format);
extern void _ZN9QDateTime10fromStringERK7QStringS2_(void* arg0, void* arg1);
  // proto:  void QDateTime::swap(QDateTime & other);
extern void demth_ZN9QDateTime4swapERS_(void* qthis, void* arg0);
  // proto:  uint QDateTime::toTime_t();
extern void _ZNK9QDateTime8toTime_tEv(void* qthis);
  // proto:  QString QDateTime::timeZoneAbbreviation();
extern void _ZNK9QDateTime20timeZoneAbbreviationEv(void* qthis);
  // proto:  QDateTime QDateTime::toUTC();
extern void demth_ZNK9QDateTime5toUTCEv(void* qthis);
  // proto:  QDate QDateTime::date();
extern void _ZNK9QDateTime4dateEv(void* qthis);
  // proto:  bool QDateTime::isNull();
extern void _ZNK9QDateTime6isNullEv(void* qthis);
  // proto: static qint64 QDateTime::currentMSecsSinceEpoch();
extern void _ZN9QDateTime22currentMSecsSinceEpochEv();
  // proto:  int QDateTime::offsetFromUtc();
extern void _ZNK9QDateTime13offsetFromUtcEv(void* qthis);
  // proto:  void QDateTime::QDateTime(const QDate & );
extern void* dector_ZN9QDateTimeC1ERK5QDate(void* arg0);
extern void _ZN9QDateTimeC1ERK5QDate(void* qthis, void* arg0);
  // proto:  QDateTime QDateTime::addMSecs(qint64 msecs);
extern void _ZNK9QDateTime8addMSecsEx(void* qthis, long long arg0);
  // proto:  qint64 QDateTime::secsTo(const QDateTime & );
extern void _ZNK9QDateTime6secsToERKS_(void* qthis, void* arg0);
  // proto:  void QDateTime::~QDateTime();
extern void _ZN9QDateTimeD0Ev(void* qthis);
  // proto:  QDateTime QDateTime::addMonths(int months);
extern void _ZNK9QDateTime9addMonthsEi(void* qthis, int arg0);
  // proto: static QDateTime QDateTime::currentDateTime();
extern void _ZN9QDateTime15currentDateTimeEv();
  // proto:  QDateTime QDateTime::toTimeZone(const QTimeZone & toZone);
extern void _ZNK9QDateTime10toTimeZoneERK9QTimeZone(void* qthis, void* arg0);
  // proto:  qint64 QDateTime::msecsTo(const QDateTime & );
extern void _ZNK9QDateTime7msecsToERKS_(void* qthis, void* arg0);
  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC);
extern void _ZN9QDateTime10fromTime_tEj(unsigned int arg0);
  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC, const QTimeZone & timeZone);
extern void _ZN9QDateTime10fromTime_tEjRK9QTimeZone(unsigned int arg0, void* arg1);
  // proto:  void QDateTime::setDate(const QDate & date);
extern void _ZN9QDateTime7setDateERK5QDate(void* qthis, void* arg0);
  // proto:  int QDateTime::utcOffset();
extern void _ZNK9QDateTime9utcOffsetEv(void* qthis);
  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs, const QTimeZone & timeZone);
extern void _ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone(long long arg0, void* arg1);
  // proto:  QTime QDateTime::time();
extern void _ZNK9QDateTime4timeEv(void* qthis);
  // proto:  qint64 QDateTime::daysTo(const QDateTime & );
extern void _ZNK9QDateTime6daysToERKS_(void* qthis, void* arg0);
  // proto:  QDateTime QDateTime::addDays(qint64 days);
extern void _ZNK9QDateTime7addDaysEx(void* qthis, long long arg0);
  // proto:  void QDateTime::setTimeZone(const QTimeZone & toZone);
extern void _ZN9QDateTime11setTimeZoneERK9QTimeZone(void* qthis, void* arg0);
  // proto: static QDateTime QDateTime::currentDateTimeUtc();
extern void _ZN9QDateTime18currentDateTimeUtcEv();
  // proto:  qint64 QDate::daysTo(const QDate & );
extern void _ZNK5QDate6daysToERKS_(void* qthis, void* arg0);
  // proto:  QDate QDate::addYears(int years);
extern void _ZNK5QDate8addYearsEi(void* qthis, int arg0);
  // proto:  int QDate::month();
extern void _ZNK5QDate5monthEv(void* qthis);
  // proto:  qint64 QDate::toJulianDay();
extern void _ZNK5QDate11toJulianDayEv(void* qthis);
  // proto:  void QDate::QDate(qint64 julianDay);
extern void* dector_ZN5QDateC1Ex(long long arg0);
extern void _ZN5QDateC1Ex(void* qthis, long long arg0);
  // proto:  void QDate::QDate();
extern void* dector_ZN5QDateC1Ev();
extern void _ZN5QDateC1Ev(void* qthis);
  // proto:  void QDate::getDate(int * year, int * month, int * day);
extern void _ZN5QDate7getDateEPiS0_S0_(void* qthis, int* arg0, int* arg1, int* arg2);
  // proto: static QDate QDate::currentDate();
extern void _ZN5QDate11currentDateEv();
  // proto:  void QDate::QDate(int y, int m, int d);
extern void* dector_ZN5QDateC1Eiii(int arg0, int arg1, int arg2);
extern void _ZN5QDateC1Eiii(void* qthis, int arg0, int arg1, int arg2);
  // proto:  int QDate::weekNumber(int * yearNum);
extern void _ZNK5QDate10weekNumberEPi(void* qthis, int* arg0);
  // proto:  QString QDate::toString(const QString & format);
extern void _ZNK5QDate8toStringERK7QString(void* qthis, void* arg0);
  // proto:  int QDate::dayOfYear();
extern void _ZNK5QDate9dayOfYearEv(void* qthis);
  // proto:  int QDate::day();
extern void _ZNK5QDate3dayEv(void* qthis);
  // proto:  bool QDate::setDate(int year, int month, int day);
extern void _ZN5QDate7setDateEiii(void* qthis, int arg0, int arg1, int arg2);
  // proto:  bool QDate::isNull();
extern void _ZNK5QDate6isNullEv(void* qthis);
  // proto: static QDate QDate::fromJulianDay(qint64 jd);
extern void demth_ZN5QDate13fromJulianDayEx(long long arg0);
  // proto:  bool QDate::isValid();
extern void _ZNK5QDate7isValidEv(void* qthis);
  // proto:  QDate QDate::addDays(qint64 days);
extern void _ZNK5QDate7addDaysEx(void* qthis, long long arg0);
  // proto: static bool QDate::isValid(int y, int m, int d);
extern void _ZN5QDate7isValidEiii(int arg0, int arg1, int arg2);
  // proto:  int QDate::daysInMonth();
extern void _ZNK5QDate11daysInMonthEv(void* qthis);
  // proto: static QDate QDate::fromString(const QString & s, const QString & format);
extern void _ZN5QDate10fromStringERK7QStringS2_(void* arg0, void* arg1);
  // proto: static bool QDate::isLeapYear(int year);
extern void _ZN5QDate10isLeapYearEi(int arg0);
  // proto:  int QDate::daysInYear();
extern void _ZNK5QDate10daysInYearEv(void* qthis);
  // proto:  int QDate::dayOfWeek();
extern void _ZNK5QDate9dayOfWeekEv(void* qthis);
  // proto:  QDate QDate::addMonths(int months);
extern void _ZNK5QDate9addMonthsEi(void* qthis, int arg0);
  // proto:  int QDate::year();
extern void _ZNK5QDate4yearEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDateTime)=1
type QDateTime struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDate)=8
type QDate struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QTime QTime::addMSecs(int ms);
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
    C._ZNK5QTime8addMSecsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTime", "addMSecs", args)
  }

}

  // proto: static QTime QTime::fromMSecsSinceStartOfDay(int msecs);
func (this *QTime) fromMSecsSinceStartOfDay_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "fromMSecsSinceStartOfDay", args)
  }

}

  // proto: static QTime QTime::currentTime();
func (this *QTime) currentTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "currentTime", args)
  }

}

  // proto:  int QTime::second();
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
    C._ZNK5QTime6secondEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "second", args)
  }

}

  // proto:  int QTime::restart();
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
    C._ZN5QTime7restartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "restart", args)
  }

}

  // proto:  void QTime::start();
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
    C._ZN5QTime5startEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "start", args)
  }

}

  // proto:  bool QTime::isNull();
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
    C._ZNK5QTime6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "isNull", args)
  }

}

  // proto:  int QTime::msecsSinceStartOfDay();
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
    C._ZNK5QTime20msecsSinceStartOfDayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "msecsSinceStartOfDay", args)
  }

}

  // proto:  int QTime::hour();
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
    C._ZNK5QTime4hourEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "hour", args)
  }

}

  // proto:  int QTime::elapsed();
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
    C._ZNK5QTime7elapsedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "elapsed", args)
  }

}

  // proto:  QTime QTime::addSecs(int secs);
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
    C._ZNK5QTime7addSecsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTime", "addSecs", args)
  }

}

  // proto:  bool QTime::isValid();
func (this *QTime) isValid(args ...interface{}) () {
  // isValid()
  // isValid(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7isValidEv
    // invoke: bool isValid()
    C._ZNK5QTime7isValidEv(this.qclsinst)
  case 1:
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
    C._ZN5QTime7isValidEiiii(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTime", "isValid", args)
  }

}

  // proto:  void QTime::QTime(int ms);
func NewQTime(args ...interface{}) QTime {
  return QTime{}
}

  // proto:  int QTime::msec();
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
    C._ZNK5QTime4msecEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "msec", args)
  }

}

  // proto:  int QTime::secsTo(const QTime & );
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
    C._ZNK5QTime6secsToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTime", "secsTo", args)
  }

}

  // proto:  bool QTime::setHMS(int h, int m, int s, int ms);
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
    C._ZN5QTime6setHMSEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTime", "setHMS", args)
  }

}

  // proto:  QString QTime::toString(const QString & format);
func (this *QTime) toString(args ...interface{}) () {
  // toString(const class QString &)
  // toString(Qt::DateFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::DateFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK5QTime8toStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTime", "toString", args)
  }

}

  // proto:  int QTime::msecsTo(const QTime & );
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
    C._ZNK5QTime7msecsToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTime", "msecsTo", args)
  }

}

  // proto:  int QTime::minute();
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
    C._ZNK5QTime6minuteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTime", "minute", args)
  }

}

  // proto: static bool QTime::isValid(int h, int m, int s, int ms);
func (this *QTime) isValid_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "isValid", args)
  }

}

  // proto: static QTime QTime::fromString(const QString & s, const QString & format);
func (this *QTime) fromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "fromString", args)
  }

}

  // proto:  QDateTime QDateTime::toLocalTime();
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
    C.demth_ZNK9QDateTime11toLocalTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "toLocalTime", args)
  }

}

  // proto:  void QDateTime::setOffsetFromUtc(int offsetSeconds);
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
    C._ZN9QDateTime16setOffsetFromUtcEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setOffsetFromUtc", args)
  }

}

  // proto:  QTimeZone QDateTime::timeZone();
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
    C._ZNK9QDateTime8timeZoneEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "timeZone", args)
  }

}

  // proto:  void QDateTime::setTime(const QTime & time);
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
    C._ZN9QDateTime7setTimeERK5QTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setTime", args)
  }

}

  // proto:  qint64 QDateTime::toMSecsSinceEpoch();
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
    C._ZNK9QDateTime17toMSecsSinceEpochEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "toMSecsSinceEpoch", args)
  }

}

  // proto:  void QDateTime::setTime_t(uint secsSince1Jan1970UTC);
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
    C._ZN9QDateTime9setTime_tEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setTime_t", args)
  }

}

  // proto:  void QDateTime::QDateTime(const QDateTime & other);
func NewQDateTime(args ...interface{}) QDateTime {
  return QDateTime{}
}

  // proto:  bool QDateTime::isDaylightTime();
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
    C._ZNK9QDateTime14isDaylightTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "isDaylightTime", args)
  }

}

  // proto:  bool QDateTime::isValid();
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
    C._ZNK9QDateTime7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "isValid", args)
  }

}

  // proto:  QString QDateTime::toString(const QString & format);
func (this *QDateTime) toString(args ...interface{}) () {
  // toString(Qt::DateFormat)
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Qt::DateFormat"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QDateTime8toStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "toString", args)
  }

}

  // proto:  QDateTime QDateTime::addYears(int years);
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
    C._ZNK9QDateTime8addYearsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "addYears", args)
  }

}

  // proto:  void QDateTime::setMSecsSinceEpoch(qint64 msecs);
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
    C._ZN9QDateTime18setMSecsSinceEpochEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setMSecsSinceEpoch", args)
  }

}

  // proto:  QDateTime QDateTime::toOffsetFromUtc(int offsetSeconds);
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
    C._ZNK9QDateTime15toOffsetFromUtcEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "toOffsetFromUtc", args)
  }

}

  // proto:  void QDateTime::setUtcOffset(int seconds);
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
    C._ZN9QDateTime12setUtcOffsetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setUtcOffset", args)
  }

}

  // proto:  QDateTime QDateTime::addSecs(qint64 secs);
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
    C._ZNK9QDateTime7addSecsEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "addSecs", args)
  }

}

  // proto: static QDateTime QDateTime::fromMSecsSinceEpoch(qint64 msecs);
func (this *QDateTime) fromMSecsSinceEpoch_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "fromMSecsSinceEpoch", args)
  }

}

  // proto: static QDateTime QDateTime::fromString(const QString & s, const QString & format);
func (this *QDateTime) fromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "fromString", args)
  }

}

  // proto:  void QDateTime::swap(QDateTime & other);
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
    C.demth_ZN9QDateTime4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "swap", args)
  }

}

  // proto:  uint QDateTime::toTime_t();
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
    C._ZNK9QDateTime8toTime_tEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "toTime_t", args)
  }

}

  // proto:  QString QDateTime::timeZoneAbbreviation();
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
    C._ZNK9QDateTime20timeZoneAbbreviationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "timeZoneAbbreviation", args)
  }

}

  // proto:  QDateTime QDateTime::toUTC();
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
    C.demth_ZNK9QDateTime5toUTCEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "toUTC", args)
  }

}

  // proto:  QDate QDateTime::date();
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
    C._ZNK9QDateTime4dateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "date", args)
  }

}

  // proto:  bool QDateTime::isNull();
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
    C._ZNK9QDateTime6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "isNull", args)
  }

}

  // proto: static qint64 QDateTime::currentMSecsSinceEpoch();
func (this *QDateTime) currentMSecsSinceEpoch_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "currentMSecsSinceEpoch", args)
  }

}

  // proto:  int QDateTime::offsetFromUtc();
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
    C._ZNK9QDateTime13offsetFromUtcEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "offsetFromUtc", args)
  }

}

  // proto:  QDateTime QDateTime::addMSecs(qint64 msecs);
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
    C._ZNK9QDateTime8addMSecsEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "addMSecs", args)
  }

}

  // proto:  qint64 QDateTime::secsTo(const QDateTime & );
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
    C._ZNK9QDateTime6secsToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "secsTo", args)
  }

}

  // proto:  void QDateTime::~QDateTime();
func (this *QDateTime) FreeQDateTime(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "~QDateTime", args)
  }

}

  // proto:  QDateTime QDateTime::addMonths(int months);
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
    C._ZNK9QDateTime9addMonthsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "addMonths", args)
  }

}

  // proto: static QDateTime QDateTime::currentDateTime();
func (this *QDateTime) currentDateTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTime", args)
  }

}

  // proto:  QDateTime QDateTime::toTimeZone(const QTimeZone & toZone);
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
    C._ZNK9QDateTime10toTimeZoneERK9QTimeZone(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "toTimeZone", args)
  }

}

  // proto:  qint64 QDateTime::msecsTo(const QDateTime & );
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
    C._ZNK9QDateTime7msecsToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "msecsTo", args)
  }

}

  // proto: static QDateTime QDateTime::fromTime_t(uint secsSince1Jan1970UTC);
func (this *QDateTime) fromTime_t_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "fromTime_t", args)
  }

}

  // proto:  void QDateTime::setDate(const QDate & date);
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
    C._ZN9QDateTime7setDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setDate", args)
  }

}

  // proto:  int QDateTime::utcOffset();
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
    C._ZNK9QDateTime9utcOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "utcOffset", args)
  }

}

  // proto:  QTime QDateTime::time();
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
    C._ZNK9QDateTime4timeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTime", "time", args)
  }

}

  // proto:  qint64 QDateTime::daysTo(const QDateTime & );
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
    C._ZNK9QDateTime6daysToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "daysTo", args)
  }

}

  // proto:  QDateTime QDateTime::addDays(qint64 days);
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
    C._ZNK9QDateTime7addDaysEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "addDays", args)
  }

}

  // proto:  void QDateTime::setTimeZone(const QTimeZone & toZone);
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
    C._ZN9QDateTime11setTimeZoneERK9QTimeZone(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTime", "setTimeZone", args)
  }

}

  // proto: static QDateTime QDateTime::currentDateTimeUtc();
func (this *QDateTime) currentDateTimeUtc_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTimeUtc", args)
  }

}

  // proto:  qint64 QDate::daysTo(const QDate & );
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
    C._ZNK5QDate6daysToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDate", "daysTo", args)
  }

}

  // proto:  QDate QDate::addYears(int years);
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
    C._ZNK5QDate8addYearsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDate", "addYears", args)
  }

}

  // proto:  int QDate::month();
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
    C._ZNK5QDate5monthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "month", args)
  }

}

  // proto:  qint64 QDate::toJulianDay();
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
    C._ZNK5QDate11toJulianDayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "toJulianDay", args)
  }

}

  // proto:  void QDate::QDate(qint64 julianDay);
func NewQDate(args ...interface{}) QDate {
  return QDate{}
}

  // proto:  void QDate::getDate(int * year, int * month, int * day);
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
    C._ZN5QDate7getDateEPiS0_S0_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDate", "getDate", args)
  }

}

  // proto: static QDate QDate::currentDate();
func (this *QDate) currentDate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "currentDate", args)
  }

}

  // proto:  int QDate::weekNumber(int * yearNum);
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
    C._ZNK5QDate10weekNumberEPi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDate", "weekNumber", args)
  }

}

  // proto:  QString QDate::toString(const QString & format);
func (this *QDate) toString(args ...interface{}) () {
  // toString(const class QString &)
  // toString(Qt::DateFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::DateFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK5QDate8toStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDate", "toString", args)
  }

}

  // proto:  int QDate::dayOfYear();
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
    C._ZNK5QDate9dayOfYearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "dayOfYear", args)
  }

}

  // proto:  int QDate::day();
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
    C._ZNK5QDate3dayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "day", args)
  }

}

  // proto:  bool QDate::setDate(int year, int month, int day);
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
    C._ZN5QDate7setDateEiii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDate", "setDate", args)
  }

}

  // proto:  bool QDate::isNull();
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
    C._ZNK5QDate6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "isNull", args)
  }

}

  // proto: static QDate QDate::fromJulianDay(qint64 jd);
func (this *QDate) fromJulianDay_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "fromJulianDay", args)
  }

}

  // proto:  bool QDate::isValid();
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
    C._ZNK5QDate7isValidEv(this.qclsinst)
  case 1:
    // invoke: _ZN5QDate7isValidEiii
    // invoke: bool isValid(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN5QDate7isValidEiii(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDate", "isValid", args)
  }

}

  // proto:  QDate QDate::addDays(qint64 days);
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
    C._ZNK5QDate7addDaysEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDate", "addDays", args)
  }

}

  // proto: static bool QDate::isValid(int y, int m, int d);
func (this *QDate) isValid_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "isValid", args)
  }

}

  // proto:  int QDate::daysInMonth();
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
    C._ZNK5QDate11daysInMonthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "daysInMonth", args)
  }

}

  // proto: static QDate QDate::fromString(const QString & s, const QString & format);
func (this *QDate) fromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "fromString", args)
  }

}

  // proto: static bool QDate::isLeapYear(int year);
func (this *QDate) isLeapYear_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "isLeapYear", args)
  }

}

  // proto:  int QDate::daysInYear();
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
    C._ZNK5QDate10daysInYearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "daysInYear", args)
  }

}

  // proto:  int QDate::dayOfWeek();
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
    C._ZNK5QDate9dayOfWeekEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "dayOfWeek", args)
  }

}

  // proto:  QDate QDate::addMonths(int months);
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
    C._ZNK5QDate9addMonthsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDate", "addMonths", args)
  }

}

  // proto:  int QDate::year();
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
    C._ZNK5QDate4yearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDate", "year", args)
  }

}

// <= body block end

