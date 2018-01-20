//  header block begin
// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDateTime struct {
	*qtrt.CObject
}

func (this *QDateTime) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qdatetime.h:261
// index:0
// void QDateTime()
func NewQDateTime() *QDateTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}
func NewQDateTimeFromPointer(cthis unsafe.Pointer) *QDateTime {
	return &QDateTime{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qdatetime.h:262
// index:1
// void QDateTime(const class QDate &)
func NewQDateTime_1(arg0 unsafe.Pointer) *QDateTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDate", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:263
// index:2
// void QDateTime(const class QDate &, const class QTime &, Qt::TimeSpec)
func NewQDateTime_2(arg0 unsafe.Pointer, arg1 unsafe.Pointer, spec int) *QDateTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecE", ffiqt.FFI_TYPE_VOID, cthis, arg0, arg1, &spec)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:265
// index:3
// void QDateTime(const class QDate &, const class QTime &, Qt::TimeSpec, int)
func NewQDateTime_3(date unsafe.Pointer, time unsafe.Pointer, spec int, offsetSeconds int) *QDateTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_VOID, cthis, date, time, &spec, &offsetSeconds)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:267
// index:4
// void QDateTime(const class QDate &, const class QTime &, const class QTimeZone &)
func NewQDateTime_4(date unsafe.Pointer, time unsafe.Pointer, timeZone unsafe.Pointer) *QDateTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeRK9QTimeZone", ffiqt.FFI_TYPE_VOID, cthis, date, time, timeZone)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:271
// index:0
// void ~QDateTime()
func DeleteQDateTime(*QDateTime) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:278
// index:0
// inline
// void swap(class QDateTime &)
func (this *QDateTime) Swap(other unsafe.Pointer) {
	// 0: (, other QDateTime &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:280
// index:0
// bool isNull()
func (this *QDateTime) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:281
// index:0
// bool isValid()
func (this *QDateTime) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:283
// index:0
// QDate date()
func (this *QDateTime) Date() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime4dateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:284
// index:0
// QTime time()
func (this *QDateTime) Time() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime4timeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:285
// index:0
// Qt::TimeSpec timeSpec()
func (this *QDateTime) TimeSpec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8timeSpecEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:286
// index:0
// int offsetFromUtc()
func (this *QDateTime) OffsetFromUtc() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime13offsetFromUtcEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:288
// index:0
// QTimeZone timeZone()
func (this *QDateTime) TimeZone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8timeZoneEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:290
// index:0
// QString timeZoneAbbreviation()
func (this *QDateTime) TimeZoneAbbreviation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime20timeZoneAbbreviationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:291
// index:0
// bool isDaylightTime()
func (this *QDateTime) IsDaylightTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime14isDaylightTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:293
// index:0
// qint64 toMSecsSinceEpoch()
func (this *QDateTime) ToMSecsSinceEpoch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime17toMSecsSinceEpochEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:294
// index:0
// qint64 toSecsSinceEpoch()
func (this *QDateTime) ToSecsSinceEpoch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime16toSecsSinceEpochEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:296
// index:0
// void setDate(const class QDate &)
func (this *QDateTime) SetDate(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime7setDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:297
// index:0
// void setTime(const class QTime &)
func (this *QDateTime) SetTime(time unsafe.Pointer) {
	// 0: (, time const QTime &), (time)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime7setTimeERK5QTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), time)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:298
// index:0
// void setTimeSpec(Qt::TimeSpec)
func (this *QDateTime) SetTimeSpec(spec int) {
	// 0: (, spec Qt::TimeSpec), (&spec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime11setTimeSpecEN2Qt8TimeSpecE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:299
// index:0
// void setOffsetFromUtc(int)
func (this *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	// 0: (, offsetSeconds int), (&offsetSeconds)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime16setOffsetFromUtcEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &offsetSeconds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:301
// index:0
// void setTimeZone(const class QTimeZone &)
func (this *QDateTime) SetTimeZone(toZone unsafe.Pointer) {
	// 0: (, toZone const QTimeZone &), (toZone)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime11setTimeZoneERK9QTimeZone", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toZone)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:303
// index:0
// void setMSecsSinceEpoch(qint64)
func (this *QDateTime) SetMSecsSinceEpoch(msecs int64) {
	// 0: (, msecs qint64), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18setMSecsSinceEpochEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:304
// index:0
// void setSecsSinceEpoch(qint64)
func (this *QDateTime) SetSecsSinceEpoch(secs int64) {
	// 0: (, secs qint64), (&secs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime17setSecsSinceEpochEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &secs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:307
// index:0
// QString toString(Qt::DateFormat)
func (this *QDateTime) ToString(f int) {
	// 0: (, f Qt::DateFormat), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toStringEN2Qt10DateFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:309
// index:1
// QString toString(const class QString &)
func (this *QDateTime) ToString_1(format unsafe.Pointer) {
	// 1: (, format const QString &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toStringERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:311
// index:2
// QString toString(class QStringView)
func (this *QDateTime) ToString_2(format unsafe.Pointer) {
	// 2: (, format QStringView), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toStringE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:313
// index:0
// QDateTime addDays(qint64)
func (this *QDateTime) AddDays(days int64) {
	// 0: (, days qint64), (&days)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7addDaysEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &days)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:314
// index:0
// QDateTime addMonths(int)
func (this *QDateTime) AddMonths(months int) {
	// 0: (, months int), (&months)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime9addMonthsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &months)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:315
// index:0
// QDateTime addYears(int)
func (this *QDateTime) AddYears(years int) {
	// 0: (, years int), (&years)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8addYearsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &years)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:316
// index:0
// QDateTime addSecs(qint64)
func (this *QDateTime) AddSecs(secs int64) {
	// 0: (, secs qint64), (&secs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7addSecsEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &secs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:317
// index:0
// QDateTime addMSecs(qint64)
func (this *QDateTime) AddMSecs(msecs int64) {
	// 0: (, msecs qint64), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8addMSecsEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:319
// index:0
// QDateTime toTimeSpec(Qt::TimeSpec)
func (this *QDateTime) ToTimeSpec(spec int) {
	// 0: (, spec Qt::TimeSpec), (&spec)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime10toTimeSpecEN2Qt8TimeSpecE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:320
// index:0
// inline
// QDateTime toLocalTime()
func (this *QDateTime) ToLocalTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime11toLocalTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:321
// index:0
// inline
// QDateTime toUTC()
func (this *QDateTime) ToUTC() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime5toUTCEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:322
// index:0
// QDateTime toOffsetFromUtc(int)
func (this *QDateTime) ToOffsetFromUtc(offsetSeconds int) {
	// 0: (, offsetSeconds int), (&offsetSeconds)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime15toOffsetFromUtcEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &offsetSeconds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:324
// index:0
// QDateTime toTimeZone(const class QTimeZone &)
func (this *QDateTime) ToTimeZone(toZone unsafe.Pointer) {
	// 0: (, toZone const QTimeZone &), (toZone)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime10toTimeZoneERK9QTimeZone", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toZone)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:327
// index:0
// qint64 daysTo(const class QDateTime &)
func (this *QDateTime) DaysTo(arg0 unsafe.Pointer) {
	// 0: (, const QDateTime & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime6daysToERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:328
// index:0
// qint64 secsTo(const class QDateTime &)
func (this *QDateTime) SecsTo(arg0 unsafe.Pointer) {
	// 0: (, const QDateTime & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime6secsToERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:329
// index:0
// qint64 msecsTo(const class QDateTime &)
func (this *QDateTime) MsecsTo(arg0 unsafe.Pointer) {
	// 0: (, const QDateTime & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7msecsToERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:339
// index:0
// void setUtcOffset(int)
func (this *QDateTime) SetUtcOffset(seconds int) {
	// 0: (, seconds int), (&seconds)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime12setUtcOffsetEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &seconds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:340
// index:0
// int utcOffset()
func (this *QDateTime) UtcOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime9utcOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:343
// index:0
// static
// QDateTime currentDateTime()
func (this *QDateTime) CurrentDateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime15currentDateTimeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_CurrentDateTime() {
	// 0: (), ()
	var nilthis *QDateTime
	nilthis.CurrentDateTime()
}

// /usr/include/qt/QtCore/qdatetime.h:344
// index:0
// static
// QDateTime currentDateTimeUtc()
func (this *QDateTime) CurrentDateTimeUtc() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18currentDateTimeUtcEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_CurrentDateTimeUtc() {
	// 0: (), ()
	var nilthis *QDateTime
	nilthis.CurrentDateTimeUtc()
}

// /usr/include/qt/QtCore/qdatetime.h:346
// index:0
// static
// QDateTime fromString(const class QString &, Qt::DateFormat)
func (this *QDateTime) FromString(s unsafe.Pointer, f int) {
	// 0: (s const QString &, f Qt::DateFormat), (s, f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringN2Qt10DateFormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromString(s unsafe.Pointer, f int) {
	// 0: (s const QString &, f Qt::DateFormat), (s, f)
	var nilthis *QDateTime
	nilthis.FromString(s, f)
}

// /usr/include/qt/QtCore/qdatetime.h:347
// index:1
// static
// QDateTime fromString(const class QString &, const class QString &)
func (this *QDateTime) FromString_1(s unsafe.Pointer, format unsafe.Pointer) {
	// 1: (s const QString &, format const QString &), (s, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromString_1(s unsafe.Pointer, format unsafe.Pointer) {
	// 1: (s const QString &, format const QString &), (s, format)
	var nilthis *QDateTime
	nilthis.FromString_1(s, format)
}

// /usr/include/qt/QtCore/qdatetime.h:351
// index:0
// uint toTime_t()
func (this *QDateTime) ToTime_t() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toTime_tEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:352
// index:0
// void setTime_t(uint)
func (this *QDateTime) SetTime_t(secsSince1Jan1970UTC uint) {
	// 0: (, secsSince1Jan1970UTC uint), (&secsSince1Jan1970UTC)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime9setTime_tEj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &secsSince1Jan1970UTC)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:353
// index:0
// static
// QDateTime fromTime_t(uint)
func (this *QDateTime) FromTime_t(secsSince1Jan1970UTC uint) {
	// 0: (secsSince1Jan1970UTC uint), (secsSince1Jan1970UTC)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromTime_t(secsSince1Jan1970UTC uint) {
	// 0: (secsSince1Jan1970UTC uint), (secsSince1Jan1970UTC)
	var nilthis *QDateTime
	nilthis.FromTime_t(secsSince1Jan1970UTC)
}

// /usr/include/qt/QtCore/qdatetime.h:354
// index:1
// static
// QDateTime fromTime_t(uint, Qt::TimeSpec, int)
func (this *QDateTime) FromTime_t_1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) {
	// 1: (secsSince1Jan1970UTC uint, spec Qt::TimeSpec, offsetFromUtc int), (secsSince1Jan1970UTC, spec, offsetFromUtc)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromTime_t_1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) {
	// 1: (secsSince1Jan1970UTC uint, spec Qt::TimeSpec, offsetFromUtc int), (secsSince1Jan1970UTC, spec, offsetFromUtc)
	var nilthis *QDateTime
	nilthis.FromTime_t_1(secsSince1Jan1970UTC, spec, offsetFromUtc)
}

// /usr/include/qt/QtCore/qdatetime.h:356
// index:2
// static
// QDateTime fromTime_t(uint, const class QTimeZone &)
func (this *QDateTime) FromTime_t_2(secsSince1Jan1970UTC uint, timeZone unsafe.Pointer) {
	// 2: (secsSince1Jan1970UTC uint, timeZone const QTimeZone &), (secsSince1Jan1970UTC, timeZone)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjRK9QTimeZone", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromTime_t_2(secsSince1Jan1970UTC uint, timeZone unsafe.Pointer) {
	// 2: (secsSince1Jan1970UTC uint, timeZone const QTimeZone &), (secsSince1Jan1970UTC, timeZone)
	var nilthis *QDateTime
	nilthis.FromTime_t_2(secsSince1Jan1970UTC, timeZone)
}

// /usr/include/qt/QtCore/qdatetime.h:359
// index:0
// static
// QDateTime fromMSecsSinceEpoch(qint64)
func (this *QDateTime) FromMSecsSinceEpoch(msecs int64) {
	// 0: (msecs qint64), (msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochEx", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromMSecsSinceEpoch(msecs int64) {
	// 0: (msecs qint64), (msecs)
	var nilthis *QDateTime
	nilthis.FromMSecsSinceEpoch(msecs)
}

// /usr/include/qt/QtCore/qdatetime.h:361
// index:1
// static
// QDateTime fromMSecsSinceEpoch(qint64, Qt::TimeSpec, int)
func (this *QDateTime) FromMSecsSinceEpoch_1(msecs int64, spec int, offsetFromUtc int) {
	// 1: (msecs qint64, spec Qt::TimeSpec, offsetFromUtc int), (msecs, spec, offsetFromUtc)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromMSecsSinceEpoch_1(msecs int64, spec int, offsetFromUtc int) {
	// 1: (msecs qint64, spec Qt::TimeSpec, offsetFromUtc int), (msecs, spec, offsetFromUtc)
	var nilthis *QDateTime
	nilthis.FromMSecsSinceEpoch_1(msecs, spec, offsetFromUtc)
}

// /usr/include/qt/QtCore/qdatetime.h:365
// index:2
// static
// QDateTime fromMSecsSinceEpoch(qint64, const class QTimeZone &)
func (this *QDateTime) FromMSecsSinceEpoch_2(msecs int64, timeZone unsafe.Pointer) {
	// 2: (msecs qint64, timeZone const QTimeZone &), (msecs, timeZone)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromMSecsSinceEpoch_2(msecs int64, timeZone unsafe.Pointer) {
	// 2: (msecs qint64, timeZone const QTimeZone &), (msecs, timeZone)
	var nilthis *QDateTime
	nilthis.FromMSecsSinceEpoch_2(msecs, timeZone)
}

// /usr/include/qt/QtCore/qdatetime.h:362
// index:0
// static
// QDateTime fromSecsSinceEpoch(qint64, Qt::TimeSpec, int)
func (this *QDateTime) FromSecsSinceEpoch(secs int64, spe int, offsetFromUtc int) {
	// 0: (secs qint64, spe Qt::TimeSpec, offsetFromUtc int), (secs, spe, offsetFromUtc)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromSecsSinceEpoch(secs int64, spe int, offsetFromUtc int) {
	// 0: (secs qint64, spe Qt::TimeSpec, offsetFromUtc int), (secs, spe, offsetFromUtc)
	var nilthis *QDateTime
	nilthis.FromSecsSinceEpoch(secs, spe, offsetFromUtc)
}

// /usr/include/qt/QtCore/qdatetime.h:366
// index:1
// static
// QDateTime fromSecsSinceEpoch(qint64, const class QTimeZone &)
func (this *QDateTime) FromSecsSinceEpoch_1(secs int64, timeZone unsafe.Pointer) {
	// 1: (secs qint64, timeZone const QTimeZone &), (secs, timeZone)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExRK9QTimeZone", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_FromSecsSinceEpoch_1(secs int64, timeZone unsafe.Pointer) {
	// 1: (secs qint64, timeZone const QTimeZone &), (secs, timeZone)
	var nilthis *QDateTime
	nilthis.FromSecsSinceEpoch_1(secs, timeZone)
}

// /usr/include/qt/QtCore/qdatetime.h:369
// index:0
// static
// qint64 currentMSecsSinceEpoch()
func (this *QDateTime) CurrentMSecsSinceEpoch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime22currentMSecsSinceEpochEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_CurrentMSecsSinceEpoch() {
	// 0: (), ()
	var nilthis *QDateTime
	nilthis.CurrentMSecsSinceEpoch()
}

// /usr/include/qt/QtCore/qdatetime.h:370
// index:0
// static
// qint64 currentSecsSinceEpoch()
func (this *QDateTime) CurrentSecsSinceEpoch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime21currentSecsSinceEpochEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDateTime_CurrentSecsSinceEpoch() {
	// 0: (), ()
	var nilthis *QDateTime
	nilthis.CurrentSecsSinceEpoch()
}

//  body block end
