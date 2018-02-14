package qtcore

// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QDateTime struct {
	*qtrt.CObject
}
type QDateTime_ITF interface {
	QDateTime_PTR() *QDateTime
}

func (ptr *QDateTime) QDateTime_PTR() *QDateTime { return ptr }

func (this *QDateTime) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDateTime) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDateTimeFromPointer(cthis unsafe.Pointer) *QDateTime {
	return &QDateTime{&qtrt.CObject{cthis}}
}
func (*QDateTime) NewFromPointer(cthis unsafe.Pointer) *QDateTime {
	return NewQDateTimeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdatetime.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateTime()
func NewQDateTime() *QDateTime {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:262
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &)
func NewQDateTime_1(arg0 QDate_ITF) *QDateTime {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDate_PTR() != nil {
		convArg0 = arg0.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDate", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:263
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, Qt::TimeSpec)
func NewQDateTime_2(arg0 QDate_ITF, arg1 QTime_ITF, spec int) *QDateTime {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDate_PTR() != nil {
		convArg0 = arg0.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QTime_PTR() != nil {
		convArg1 = arg1.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, spec)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:265
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, Qt::TimeSpec, int)
func NewQDateTime_3(date QDate_ITF, time QTime_ITF, spec int, offsetSeconds int) *QDateTime {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg1 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, spec, offsetSeconds)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:267
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, const QTimeZone &)
func NewQDateTime_4(date QDate_ITF, time QTime_ITF, timeZone QTimeZone_ITF) *QDateTime {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg1 = time.QTime_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if timeZone != nil && timeZone.QTimeZone_PTR() != nil {
		convArg2 = timeZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeRK9QTimeZone", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDateTime()
func DeleteQDateTime(this *QDateTime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdatetime.h:274
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime & operator=(QDateTime &&)
func (this *QDateTime) Operator_equal(other unsafe.Pointer /*333*/) *QDateTime {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:276
// index:1
// Public Visibility=Default Availability=Available
// [8] QDateTime & operator=(const QDateTime &)
func (this *QDateTime) Operator_equal_1(other QDateTime_ITF) *QDateTime {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:278
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDateTime &)
func (this *QDateTime) Swap(other QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:280
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QDateTime) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:281
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QDateTime) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:283
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate date()
func (this *QDateTime) Date() *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime4dateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:284
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime time()
func (this *QDateTime) Time() *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime4timeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:285
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TimeSpec timeSpec()
func (this *QDateTime) TimeSpec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8timeSpecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatetime.h:286
// index:0
// Public Visibility=Default Availability=Available
// [4] int offsetFromUtc()
func (this *QDateTime) OffsetFromUtc() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime13offsetFromUtcEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:288
// index:0
// Public Visibility=Default Availability=Available
// [8] QTimeZone timeZone()
func (this *QDateTime) TimeZone() *QTimeZone /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8timeZoneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:290
// index:0
// Public Visibility=Default Availability=Available
// [8] QString timeZoneAbbreviation()
func (this *QDateTime) TimeZoneAbbreviation() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime20timeZoneAbbreviationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:291
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDaylightTime()
func (this *QDateTime) IsDaylightTime() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime14isDaylightTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:293
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 toMSecsSinceEpoch()
func (this *QDateTime) ToMSecsSinceEpoch() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime17toMSecsSinceEpochEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:294
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 toSecsSinceEpoch()
func (this *QDateTime) ToSecsSinceEpoch() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime16toSecsSinceEpochEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:296
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDate(const QDate &)
func (this *QDateTime) SetDate(date QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime7setDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTime(const QTime &)
func (this *QDateTime) SetTime(time QTime_ITF) {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime7setTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeSpec(Qt::TimeSpec)
func (this *QDateTime) SetTimeSpec(spec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime11setTimeSpecEN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffsetFromUtc(int)
func (this *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime16setOffsetFromUtcEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offsetSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeZone(const QTimeZone &)
func (this *QDateTime) SetTimeZone(toZone QTimeZone_ITF) {
	var convArg0 unsafe.Pointer
	if toZone != nil && toZone.QTimeZone_PTR() != nil {
		convArg0 = toZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime11setTimeZoneERK9QTimeZone", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:303
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMSecsSinceEpoch(qint64)
func (this *QDateTime) SetMSecsSinceEpoch(msecs int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18setMSecsSinceEpochEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSecsSinceEpoch(qint64)
func (this *QDateTime) SetSecsSinceEpoch(secs int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime17setSecsSinceEpochEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:307
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat)
func (this *QDateTime) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:309
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &)
func (this *QDateTime) ToString_1(format string) string {
	var tmpArg0 = NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:311
// index:2
// Public Visibility=Default Availability=Available
// [8] QString toString(QStringView)
func (this *QDateTime) ToString_2(format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg0 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:313
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addDays(qint64)
func (this *QDateTime) AddDays(days int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7addDaysEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), days)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:314
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addMonths(int)
func (this *QDateTime) AddMonths(months int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime9addMonthsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), months)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:315
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addYears(int)
func (this *QDateTime) AddYears(years int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8addYearsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), years)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:316
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addSecs(qint64)
func (this *QDateTime) AddSecs(secs int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7addSecsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:317
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addMSecs(qint64)
func (this *QDateTime) AddMSecs(msecs int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8addMSecsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toTimeSpec(Qt::TimeSpec)
func (this *QDateTime) ToTimeSpec(spec int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime10toTimeSpecEN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:320
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime toLocalTime()
func (this *QDateTime) ToLocalTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime11toLocalTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:321
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime toUTC()
func (this *QDateTime) ToUTC() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime5toUTCEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:322
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toOffsetFromUtc(int)
func (this *QDateTime) ToOffsetFromUtc(offsetSeconds int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime15toOffsetFromUtcEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offsetSeconds)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:324
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toTimeZone(const QTimeZone &)
func (this *QDateTime) ToTimeZone(toZone QTimeZone_ITF) *QDateTime /*123*/ {
	var convArg0 unsafe.Pointer
	if toZone != nil && toZone.QTimeZone_PTR() != nil {
		convArg0 = toZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime10toTimeZoneERK9QTimeZone", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:327
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 daysTo(const QDateTime &)
func (this *QDateTime) DaysTo(arg0 QDateTime_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDateTime_PTR() != nil {
		convArg0 = arg0.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime6daysToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:328
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 secsTo(const QDateTime &)
func (this *QDateTime) SecsTo(arg0 QDateTime_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDateTime_PTR() != nil {
		convArg0 = arg0.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime6secsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:329
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 msecsTo(const QDateTime &)
func (this *QDateTime) MsecsTo(arg0 QDateTime_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDateTime_PTR() != nil {
		convArg0 = arg0.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7msecsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:331
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QDateTime &)
func (this *QDateTime) Operator_equal_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:332
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QDateTime &)
func (this *QDateTime) Operator_not_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:333
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QDateTime &)
func (this *QDateTime) Operator_less_than(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:334
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QDateTime &)
func (this *QDateTime) Operator_less_than_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:335
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QDateTime &)
func (this *QDateTime) Operator_greater_than(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimegtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:336
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QDateTime &)
func (this *QDateTime) Operator_greater_than_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimegeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:339
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUtcOffset(int)
func (this *QDateTime) SetUtcOffset(seconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime12setUtcOffsetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), seconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:340
// index:0
// Public Visibility=Default Availability=Available
// [4] int utcOffset()
func (this *QDateTime) UtcOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime9utcOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:343
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime currentDateTime()
func (this *QDateTime) CurrentDateTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime15currentDateTimeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_CurrentDateTime() *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.CurrentDateTime()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:344
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime currentDateTimeUtc()
func (this *QDateTime) CurrentDateTimeUtc() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18currentDateTimeUtcEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_CurrentDateTimeUtc() *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.CurrentDateTimeUtc()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:346
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromString(const QString &, Qt::DateFormat)
func (this *QDateTime) FromString(s string, f int) *QDateTime /*123*/ {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromString(s string, f int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromString(s, f)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:347
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromString(const QString &, const QString &)
func (this *QDateTime) FromString_1(s string, format string) *QDateTime /*123*/ {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromString_1(s string, format string) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromString_1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:351
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toTime_t()
func (this *QDateTime) ToTime_t() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toTime_tEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:352
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTime_t(uint)
func (this *QDateTime) SetTime_t(secsSince1Jan1970UTC uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime9setTime_tEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secsSince1Jan1970UTC)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:353
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint)
func (this *QDateTime) FromTime_t(secsSince1Jan1970UTC uint) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEj", qtrt.FFI_TYPE_POINTER, secsSince1Jan1970UTC)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromTime_t(secsSince1Jan1970UTC uint) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t(secsSince1Jan1970UTC)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:354
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint, Qt::TimeSpec, int)
func (this *QDateTime) FromTime_t_1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, secsSince1Jan1970UTC, spec, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromTime_t_1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t_1(secsSince1Jan1970UTC, spec, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:356
// index:2
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint, const QTimeZone &)
func (this *QDateTime) FromTime_t_2(secsSince1Jan1970UTC uint, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var convArg1 unsafe.Pointer
	if timeZone != nil && timeZone.QTimeZone_PTR() != nil {
		convArg1 = timeZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjRK9QTimeZone", qtrt.FFI_TYPE_POINTER, secsSince1Jan1970UTC, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromTime_t_2(secsSince1Jan1970UTC uint, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t_2(secsSince1Jan1970UTC, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:359
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64)
func (this *QDateTime) FromMSecsSinceEpoch(msecs int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochEx", qtrt.FFI_TYPE_POINTER, msecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromMSecsSinceEpoch(msecs int64) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch(msecs)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:361
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64, Qt::TimeSpec, int)
func (this *QDateTime) FromMSecsSinceEpoch_1(msecs int64, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, msecs, spec, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromMSecsSinceEpoch_1(msecs int64, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch_1(msecs, spec, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:365
// index:2
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64, const QTimeZone &)
func (this *QDateTime) FromMSecsSinceEpoch_2(msecs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var convArg1 unsafe.Pointer
	if timeZone != nil && timeZone.QTimeZone_PTR() != nil {
		convArg1 = timeZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExRK9QTimeZone", qtrt.FFI_TYPE_POINTER, msecs, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromMSecsSinceEpoch_2(msecs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch_2(msecs, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:362
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromSecsSinceEpoch(qint64, Qt::TimeSpec, int)
func (this *QDateTime) FromSecsSinceEpoch(secs int64, spe int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, secs, spe, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromSecsSinceEpoch(secs int64, spe int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromSecsSinceEpoch(secs, spe, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:366
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromSecsSinceEpoch(qint64, const QTimeZone &)
func (this *QDateTime) FromSecsSinceEpoch_1(secs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var convArg1 unsafe.Pointer
	if timeZone != nil && timeZone.QTimeZone_PTR() != nil {
		convArg1 = timeZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExRK9QTimeZone", qtrt.FFI_TYPE_POINTER, secs, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromSecsSinceEpoch_1(secs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromSecsSinceEpoch_1(secs, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:369
// index:0
// Public static Visibility=Default Availability=Available
// [8] qint64 currentMSecsSinceEpoch()
func (this *QDateTime) CurrentMSecsSinceEpoch() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime22currentMSecsSinceEpochEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}
func QDateTime_CurrentMSecsSinceEpoch() int64 {
	var nilthis *QDateTime
	rv := nilthis.CurrentMSecsSinceEpoch()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:370
// index:0
// Public static Visibility=Default Availability=Available
// [8] qint64 currentSecsSinceEpoch()
func (this *QDateTime) CurrentSecsSinceEpoch() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime21currentSecsSinceEpochEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}
func QDateTime_CurrentSecsSinceEpoch() int64 {
	var nilthis *QDateTime
	rv := nilthis.CurrentSecsSinceEpoch()
	return rv
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
