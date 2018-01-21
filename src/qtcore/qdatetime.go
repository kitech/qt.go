package qtcore

// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQDateTimeFromPointer(cthis unsafe.Pointer) *QDateTime {
	return &QDateTime{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qdatetime.h:261
// index:0
// Public
// void QDateTime()
func NewQDateTime() *QDateTime {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:262
// index:1
// Public
// void QDateTime(const class QDate &)
func NewQDateTime_1(arg0 *QDate) *QDateTime {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDate", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:263
// index:2
// Public
// void QDateTime(const class QDate &, const class QTime &, Qt::TimeSpec)
func NewQDateTime_2(arg0 *QDate, arg1 *QTime, spec int) *QDateTime {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, &spec)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:265
// index:3
// Public
// void QDateTime(const class QDate &, const class QTime &, Qt::TimeSpec, int)
func NewQDateTime_3(date *QDate, time *QTime, spec int, offsetSeconds int) *QDateTime {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = date.GetCthis()
	var convArg1 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, &spec, &offsetSeconds)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:267
// index:4
// Public
// void QDateTime(const class QDate &, const class QTime &, const class QTimeZone &)
func NewQDateTime_4(date *QDate, time *QTime, timeZone *QTimeZone) *QDateTime {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = date.GetCthis()
	var convArg1 = time.GetCthis()
	var convArg2 = timeZone.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeRK9QTimeZone", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:271
// index:0
// Public
// void ~QDateTime()
func DeleteQDateTime(*QDateTime) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTimeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:278
// index:0
// Public inline
// void swap(class QDateTime &)
func (this *QDateTime) Swap(other *QDateTime) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:280
// index:0
// Public
// bool isNull()
func (this *QDateTime) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:281
// index:0
// Public
// bool isValid()
func (this *QDateTime) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:283
// index:0
// Public
// QDate date()
func (this *QDateTime) Date() *QDate /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime4dateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:284
// index:0
// Public
// QTime time()
func (this *QDateTime) Time() *QTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime4timeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:285
// index:0
// Public
// Qt::TimeSpec timeSpec()
func (this *QDateTime) TimeSpec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8timeSpecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatetime.h:286
// index:0
// Public
// int offsetFromUtc()
func (this *QDateTime) OffsetFromUtc() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime13offsetFromUtcEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:288
// index:0
// Public
// QTimeZone timeZone()
func (this *QDateTime) TimeZone() *QTimeZone /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8timeZoneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:290
// index:0
// Public
// QString timeZoneAbbreviation()
func (this *QDateTime) TimeZoneAbbreviation() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime20timeZoneAbbreviationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:291
// index:0
// Public
// bool isDaylightTime()
func (this *QDateTime) IsDaylightTime() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime14isDaylightTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:293
// index:0
// Public
// qint64 toMSecsSinceEpoch()
func (this *QDateTime) ToMSecsSinceEpoch() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime17toMSecsSinceEpochEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:294
// index:0
// Public
// qint64 toSecsSinceEpoch()
func (this *QDateTime) ToSecsSinceEpoch() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime16toSecsSinceEpochEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:296
// index:0
// Public
// void setDate(const class QDate &)
func (this *QDateTime) SetDate(date *QDate) {
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime7setDateERK5QDate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:297
// index:0
// Public
// void setTime(const class QTime &)
func (this *QDateTime) SetTime(time *QTime) {
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime7setTimeERK5QTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:298
// index:0
// Public
// void setTimeSpec(Qt::TimeSpec)
func (this *QDateTime) SetTimeSpec(spec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime11setTimeSpecEN2Qt8TimeSpecE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:299
// index:0
// Public
// void setOffsetFromUtc(int)
func (this *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime16setOffsetFromUtcEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offsetSeconds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:301
// index:0
// Public
// void setTimeZone(const class QTimeZone &)
func (this *QDateTime) SetTimeZone(toZone *QTimeZone) {
	var convArg0 = toZone.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime11setTimeZoneERK9QTimeZone", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:303
// index:0
// Public
// void setMSecsSinceEpoch(qint64)
func (this *QDateTime) SetMSecsSinceEpoch(msecs int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18setMSecsSinceEpochEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:304
// index:0
// Public
// void setSecsSinceEpoch(qint64)
func (this *QDateTime) SetSecsSinceEpoch(secs int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime17setSecsSinceEpochEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &secs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:307
// index:0
// Public
// QString toString(Qt::DateFormat)
func (this *QDateTime) ToString(f int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toStringEN2Qt10DateFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:309
// index:1
// Public
// QString toString(const class QString &)
func (this *QDateTime) ToString_1(format *QString) *QString /*123*/ {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toStringERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:311
// index:2
// Public
// QString toString(class QStringView)
func (this *QDateTime) ToString_2(format *QStringView /*123*/) *QString /*123*/ {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toStringE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:313
// index:0
// Public
// QDateTime addDays(qint64)
func (this *QDateTime) AddDays(days int64) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7addDaysEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &days)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:314
// index:0
// Public
// QDateTime addMonths(int)
func (this *QDateTime) AddMonths(months int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime9addMonthsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &months)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:315
// index:0
// Public
// QDateTime addYears(int)
func (this *QDateTime) AddYears(years int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8addYearsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &years)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:316
// index:0
// Public
// QDateTime addSecs(qint64)
func (this *QDateTime) AddSecs(secs int64) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7addSecsEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &secs)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:317
// index:0
// Public
// QDateTime addMSecs(qint64)
func (this *QDateTime) AddMSecs(msecs int64) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8addMSecsEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:319
// index:0
// Public
// QDateTime toTimeSpec(Qt::TimeSpec)
func (this *QDateTime) ToTimeSpec(spec int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime10toTimeSpecEN2Qt8TimeSpecE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spec)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:320
// index:0
// Public inline
// QDateTime toLocalTime()
func (this *QDateTime) ToLocalTime() *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime11toLocalTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:321
// index:0
// Public inline
// QDateTime toUTC()
func (this *QDateTime) ToUTC() *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime5toUTCEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:322
// index:0
// Public
// QDateTime toOffsetFromUtc(int)
func (this *QDateTime) ToOffsetFromUtc(offsetSeconds int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime15toOffsetFromUtcEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offsetSeconds)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:324
// index:0
// Public
// QDateTime toTimeZone(const class QTimeZone &)
func (this *QDateTime) ToTimeZone(toZone *QTimeZone) *QDateTime /*123*/ {
	var convArg0 = toZone.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime10toTimeZoneERK9QTimeZone", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:327
// index:0
// Public
// qint64 daysTo(const class QDateTime &)
func (this *QDateTime) DaysTo(arg0 *QDateTime) int64 {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime6daysToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:328
// index:0
// Public
// qint64 secsTo(const class QDateTime &)
func (this *QDateTime) SecsTo(arg0 *QDateTime) int64 {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime6secsToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:329
// index:0
// Public
// qint64 msecsTo(const class QDateTime &)
func (this *QDateTime) MsecsTo(arg0 *QDateTime) int64 {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime7msecsToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:339
// index:0
// Public
// void setUtcOffset(int)
func (this *QDateTime) SetUtcOffset(seconds int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime12setUtcOffsetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &seconds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:340
// index:0
// Public
// int utcOffset()
func (this *QDateTime) UtcOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime9utcOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:343
// index:0
// Public static
// QDateTime currentDateTime()
func (this *QDateTime) CurrentDateTime() *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime15currentDateTimeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_CurrentDateTime() *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.CurrentDateTime()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:344
// index:0
// Public static
// QDateTime currentDateTimeUtc()
func (this *QDateTime) CurrentDateTimeUtc() *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18currentDateTimeUtcEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_CurrentDateTimeUtc() *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.CurrentDateTimeUtc()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:346
// index:0
// Public static
// QDateTime fromString(const class QString &, Qt::DateFormat)
func (this *QDateTime) FromString(s *QString, f int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringN2Qt10DateFormatE", ffiqt.FFI_TYPE_POINTER, s, f)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromString(s *QString, f int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromString(s, f)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:347
// index:1
// Public static
// QDateTime fromString(const class QString &, const class QString &)
func (this *QDateTime) FromString_1(s *QString, format *QString) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, s, format)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromString_1(s *QString, format *QString) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromString_1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:351
// index:0
// Public
// uint toTime_t()
func (this *QDateTime) ToTime_t() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateTime8toTime_tEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:352
// index:0
// Public
// void setTime_t(uint)
func (this *QDateTime) SetTime_t(secsSince1Jan1970UTC uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime9setTime_tEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &secsSince1Jan1970UTC)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:353
// index:0
// Public static
// QDateTime fromTime_t(uint)
func (this *QDateTime) FromTime_t(secsSince1Jan1970UTC uint) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEj", ffiqt.FFI_TYPE_POINTER, secsSince1Jan1970UTC)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromTime_t(secsSince1Jan1970UTC uint) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t(secsSince1Jan1970UTC)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:354
// index:1
// Public static
// QDateTime fromTime_t(uint, Qt::TimeSpec, int)
func (this *QDateTime) FromTime_t_1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_POINTER, secsSince1Jan1970UTC, spec, offsetFromUtc)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromTime_t_1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t_1(secsSince1Jan1970UTC, spec, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:356
// index:2
// Public static
// QDateTime fromTime_t(uint, const class QTimeZone &)
func (this *QDateTime) FromTime_t_2(secsSince1Jan1970UTC uint, timeZone *QTimeZone) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjRK9QTimeZone", ffiqt.FFI_TYPE_POINTER, secsSince1Jan1970UTC, timeZone)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromTime_t_2(secsSince1Jan1970UTC uint, timeZone *QTimeZone) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t_2(secsSince1Jan1970UTC, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:359
// index:0
// Public static
// QDateTime fromMSecsSinceEpoch(qint64)
func (this *QDateTime) FromMSecsSinceEpoch(msecs int64) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochEx", ffiqt.FFI_TYPE_POINTER, msecs)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromMSecsSinceEpoch(msecs int64) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch(msecs)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:361
// index:1
// Public static
// QDateTime fromMSecsSinceEpoch(qint64, Qt::TimeSpec, int)
func (this *QDateTime) FromMSecsSinceEpoch_1(msecs int64, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_POINTER, msecs, spec, offsetFromUtc)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromMSecsSinceEpoch_1(msecs int64, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch_1(msecs, spec, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:365
// index:2
// Public static
// QDateTime fromMSecsSinceEpoch(qint64, const class QTimeZone &)
func (this *QDateTime) FromMSecsSinceEpoch_2(msecs int64, timeZone *QTimeZone) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone", ffiqt.FFI_TYPE_POINTER, msecs, timeZone)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromMSecsSinceEpoch_2(msecs int64, timeZone *QTimeZone) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch_2(msecs, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:362
// index:0
// Public static
// QDateTime fromSecsSinceEpoch(qint64, Qt::TimeSpec, int)
func (this *QDateTime) FromSecsSinceEpoch(secs int64, spe int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExN2Qt8TimeSpecEi", ffiqt.FFI_TYPE_POINTER, secs, spe, offsetFromUtc)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromSecsSinceEpoch(secs int64, spe int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromSecsSinceEpoch(secs, spe, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:366
// index:1
// Public static
// QDateTime fromSecsSinceEpoch(qint64, const class QTimeZone &)
func (this *QDateTime) FromSecsSinceEpoch_1(secs int64, timeZone *QTimeZone) *QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExRK9QTimeZone", ffiqt.FFI_TYPE_POINTER, secs, timeZone)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDateTime_FromSecsSinceEpoch_1(secs int64, timeZone *QTimeZone) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromSecsSinceEpoch_1(secs, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:369
// index:0
// Public static
// qint64 currentMSecsSinceEpoch()
func (this *QDateTime) CurrentMSecsSinceEpoch() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime22currentMSecsSinceEpochEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int64(rv) // 222
}
func QDateTime_CurrentMSecsSinceEpoch() int64 {
	var nilthis *QDateTime
	rv := nilthis.CurrentMSecsSinceEpoch()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:370
// index:0
// Public static
// qint64 currentSecsSinceEpoch()
func (this *QDateTime) CurrentSecsSinceEpoch() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateTime21currentSecsSinceEpochEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int64(rv) // 222
}
func QDateTime_CurrentSecsSinceEpoch() int64 {
	var nilthis *QDateTime
	rv := nilthis.CurrentSecsSinceEpoch()
	return rv
}

//  body block end
