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
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
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

// /usr/include/qt/QtCore/qdatetime.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateTime()

/*
Constructs a null datetime (i.e. null date and null time). A null datetime is invalid, since the date is invalid.

See also isValid().
*/
func (*QDateTime) NewForInherit() *QDateTime {
	return NewQDateTime()
}
func NewQDateTime() *QDateTime {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:252
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &)

/*
Constructs a null datetime (i.e. null date and null time). A null datetime is invalid, since the date is invalid.

See also isValid().
*/
func (*QDateTime) NewForInherit1(arg0 QDate_ITF) *QDateTime {
	return NewQDateTime1(arg0)
}
func NewQDateTime1(arg0 QDate_ITF) *QDateTime {
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

// /usr/include/qt/QtCore/qdatetime.h:253
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, Qt::TimeSpec)

/*
Constructs a null datetime (i.e. null date and null time). A null datetime is invalid, since the date is invalid.

See also isValid().
*/
func (*QDateTime) NewForInherit2(arg0 QDate_ITF, arg1 QTime_ITF, spec int) *QDateTime {
	return NewQDateTime2(arg0, arg1, spec)
}
func NewQDateTime2(arg0 QDate_ITF, arg1 QTime_ITF, spec int) *QDateTime {
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

// /usr/include/qt/QtCore/qdatetime.h:253
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, Qt::TimeSpec)

/*
Constructs a null datetime (i.e. null date and null time). A null datetime is invalid, since the date is invalid.

See also isValid().
*/
func (*QDateTime) NewForInherit2p(arg0 QDate_ITF, arg1 QTime_ITF) *QDateTime {
	return NewQDateTime2p(arg0, arg1)
}
func NewQDateTime2p(arg0 QDate_ITF, arg1 QTime_ITF) *QDateTime {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDate_PTR() != nil {
		convArg0 = arg0.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QTime_PTR() != nil {
		convArg1 = arg1.QTime_PTR().GetCthis()
	}
	// arg: 2, Qt::TimeSpec=Elaborated, Qt::TimeSpec=Enum, , Invalid
	spec := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeC2ERK5QDateRK5QTimeN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, spec)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDateTime)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:255
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, Qt::TimeSpec, int)

/*
Constructs a null datetime (i.e. null date and null time). A null datetime is invalid, since the date is invalid.

See also isValid().
*/
func (*QDateTime) NewForInherit3(date QDate_ITF, time QTime_ITF, spec int, offsetSeconds int) *QDateTime {
	return NewQDateTime3(date, time, spec, offsetSeconds)
}
func NewQDateTime3(date QDate_ITF, time QTime_ITF, spec int, offsetSeconds int) *QDateTime {
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

// /usr/include/qt/QtCore/qdatetime.h:257
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QDateTime(const QDate &, const QTime &, const QTimeZone &)

/*
Constructs a null datetime (i.e. null date and null time). A null datetime is invalid, since the date is invalid.

See also isValid().
*/
func (*QDateTime) NewForInherit4(date QDate_ITF, time QTime_ITF, timeZone QTimeZone_ITF) *QDateTime {
	return NewQDateTime4(date, time, timeZone)
}
func NewQDateTime4(date QDate_ITF, time QTime_ITF, timeZone QTimeZone_ITF) *QDateTime {
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

// /usr/include/qt/QtCore/qdatetime.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDateTime()

/*

 */
func DeleteQDateTime(this *QDateTime) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdatetime.h:264
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime & operator=(QDateTime &&)

/*

 */
func (this *QDateTime) Operator_equal(other unsafe.Pointer /*333*/) *QDateTime {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTimeaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:266
// index:1
// Public Visibility=Default Availability=Available
// [8] QDateTime & operator=(const QDateTime &)

/*

 */
func (this *QDateTime) Operator_equal1(other QDateTime_ITF) *QDateTime {
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

// /usr/include/qt/QtCore/qdatetime.h:268
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDateTime &)

/*
Swaps this datetime with other. This operation is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QDateTime) Swap(other QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:270
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if both the date and the time are null; otherwise returns false. A null datetime is invalid.

See also QDate::isNull(), QTime::isNull(), and isValid().
*/
func (this *QDateTime) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:271
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if both the date and the time are valid and they are valid in the current Qt::TimeSpec, otherwise returns false.

If the timeSpec() is Qt::LocalTime or Qt::TimeZone then the date and time are checked to see if they fall in the Standard Time to Daylight-Saving Time transition hour, i.e. if the transition is at 2am and the clock goes forward to 3am then the time from 02:00:00 to 02:59:59.999 is considered to be invalid.

See also QDate::isValid() and QTime::isValid().
*/
func (this *QDateTime) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:273
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate date() const

/*
Returns the date part of the datetime.

See also setDate(), time(), and timeSpec().
*/
func (this *QDateTime) Date() *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime4dateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:274
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime time() const

/*
Returns the time part of the datetime.

See also setTime(), date(), and timeSpec().
*/
func (this *QDateTime) Time() *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime4timeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:275
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TimeSpec timeSpec() const

/*
Returns the time specification of the datetime.

See also setTimeSpec(), date(), time(), and Qt::TimeSpec.
*/
func (this *QDateTime) TimeSpec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8timeSpecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatetime.h:276
// index:0
// Public Visibility=Default Availability=Available
// [4] int offsetFromUtc() const

/*
Returns the current Offset From UTC in seconds.

If the timeSpec() is Qt::OffsetFromUTC this will be the value originally set.

If the timeSpec() is Qt::TimeZone this will be the offset effective in the Time Zone including any Daylight-Saving Offset.

If the timeSpec() is Qt::LocalTime this will be the difference between the Local Time and UTC including any Daylight-Saving Offset.

If the timeSpec() is Qt::UTC this will be 0.

This function was introduced in  Qt 5.2.

See also setOffsetFromUtc().
*/
func (this *QDateTime) OffsetFromUtc() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime13offsetFromUtcEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:278
// index:0
// Public Visibility=Default Availability=Available
// [8] QTimeZone timeZone() const

/*
Returns the time zone of the datetime.

If the timeSpec() is Qt::LocalTime then an instance of the current system time zone will be returned. Note however that if you copy this time zone the instance will not remain in sync if the system time zone changes.

This function was introduced in  Qt 5.2.

See also setTimeZone() and Qt::TimeSpec.
*/
func (this *QDateTime) TimeZone() *QTimeZone /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8timeZoneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:280
// index:0
// Public Visibility=Default Availability=Available
// [8] QString timeZoneAbbreviation() const

/*
Returns the Time Zone Abbreviation for the datetime.

If the timeSpec() is Qt::UTC this will be "UTC".

If the timeSpec() is Qt::OffsetFromUTC this will be in the format "UTC[+-]00:00".

If the timeSpec() is Qt::LocalTime then the host system is queried for the correct abbreviation.

Note that abbreviations may or may not be localized.

Note too that the abbreviation is not guaranteed to be a unique value, i.e. different time zones may have the same abbreviation.

This function was introduced in  Qt 5.2.

See also timeSpec().
*/
func (this *QDateTime) TimeZoneAbbreviation() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime20timeZoneAbbreviationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:281
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDaylightTime() const

/*
Returns if this datetime falls in Daylight-Saving Time.

If the Qt::TimeSpec is not Qt::LocalTime or Qt::TimeZone then will always return false.

This function was introduced in  Qt 5.2.

See also timeSpec().
*/
func (this *QDateTime) IsDaylightTime() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime14isDaylightTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:283
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 toMSecsSinceEpoch() const

/*
Returns the datetime as the number of milliseconds that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC).

On systems that do not support time zones, this function will behave as if local time were Qt::UTC.

The behavior for this function is undefined if the datetime stored in this object is not valid. However, for all valid dates, this function returns a unique value.

This function was introduced in  Qt 4.7.

See also toSecsSinceEpoch() and setMSecsSinceEpoch().
*/
func (this *QDateTime) ToMSecsSinceEpoch() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime17toMSecsSinceEpochEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:284
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 toSecsSinceEpoch() const

/*
Returns the datetime as the number of seconds that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC).

On systems that do not support time zones, this function will behave as if local time were Qt::UTC.

The behavior for this function is undefined if the datetime stored in this object is not valid. However, for all valid dates, this function returns a unique value.

This function was introduced in  Qt 5.8.

See also toMSecsSinceEpoch() and setSecsSinceEpoch().
*/
func (this *QDateTime) ToSecsSinceEpoch() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime16toSecsSinceEpochEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:286
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDate(const QDate &)

/*
Sets the date part of this datetime to date. If no time is set yet, it is set to midnight. If date is invalid, this QDateTime becomes invalid.

See also date(), setTime(), and setTimeSpec().
*/
func (this *QDateTime) SetDate(date QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime7setDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTime(const QTime &)

/*
Sets the time part of this datetime to time. If time is not valid, this function sets it to midnight. Therefore, it's possible to clear any set time in a QDateTime by setting it to a default QTime:


  QDateTime dt = QDateTime::currentDateTime();
  dt.setTime(QTime());



See also time(), setDate(), and setTimeSpec().
*/
func (this *QDateTime) SetTime(time QTime_ITF) {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime7setTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeSpec(Qt::TimeSpec)

/*
Sets the time specification used in this datetime to spec. The datetime will refer to a different point in time.

If spec is Qt::OffsetFromUTC then the timeSpec() will be set to Qt::UTC, i.e. an effective offset of 0.

If spec is Qt::TimeZone then the spec will be set to Qt::LocalTime, i.e. the current system time zone.

Example:


  QDateTime local(QDateTime::currentDateTime());
  qDebug() << "Local time is:" << local;

  QDateTime UTC(local);
  UTC.setTimeSpec(Qt::UTC);
  qDebug() << "UTC time is:" << UTC;

  qDebug() << "There are" << local.secsTo(UTC) << "seconds difference between the datetimes.";



See also timeSpec(), setDate(), setTime(), setTimeZone(), and Qt::TimeSpec.
*/
func (this *QDateTime) SetTimeSpec(spec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime11setTimeSpecEN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffsetFromUtc(int)

/*
Sets the timeSpec() to Qt::OffsetFromUTC and the offset to offsetSeconds. The datetime will refer to a different point in time.

The maximum and minimum offset is 14 positive or negative hours. If offsetSeconds is larger or smaller than that, then the result is undefined.

If offsetSeconds is 0 then the timeSpec() will be set to Qt::UTC.

This function was introduced in  Qt 5.2.

See also isValid() and offsetFromUtc().
*/
func (this *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime16setOffsetFromUtcEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offsetSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeZone(const QTimeZone &)

/*
Sets the time zone used in this datetime to toZone. The datetime will refer to a different point in time.

If toZone is invalid then the datetime will be invalid.

This function was introduced in  Qt 5.2.

See also timeZone() and Qt::TimeSpec.
*/
func (this *QDateTime) SetTimeZone(toZone QTimeZone_ITF) {
	var convArg0 unsafe.Pointer
	if toZone != nil && toZone.QTimeZone_PTR() != nil {
		convArg0 = toZone.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime11setTimeZoneERK9QTimeZone", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMSecsSinceEpoch(qint64)

/*
Sets the date and time given the number of milliseconds msecs that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC). On systems that do not support time zones this function will behave as if local time were Qt::UTC.

Note that passing the minimum of qint64 (std::numeric_limits<qint64>::min()) to msecs will result in undefined behavior.

This function was introduced in  Qt 4.7.

See also toMSecsSinceEpoch() and setSecsSinceEpoch().
*/
func (this *QDateTime) SetMSecsSinceEpoch(msecs int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18setMSecsSinceEpochEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSecsSinceEpoch(qint64)

/*
Sets the date and time given the number of seconds secs that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC). On systems that do not support time zones this function will behave as if local time were Qt::UTC.

This function was introduced in  Qt 5.8.

See also toSecsSinceEpoch() and setMSecsSinceEpoch().
*/
func (this *QDateTime) SetSecsSinceEpoch(secs int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime17setSecsSinceEpochEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:297
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*
Returns the datetime as a string. The format parameter determines the format of the result string.

These expressions may be used for the date:


 ExpressionOutput
dthe day as number without a leading zero (1 to 31)
ddthe day as number with a leading zero (01 to 31)
dddthe abbreviated localized day name (e.g. 'Mon' to 'Sun'). Uses the system locale to localize the name, i.e. QLocale::system().
ddddthe long localized day name (e.g. 'Monday' to 'Sunday'). Uses the system locale to localize the name, i.e. QLocale::system().
Mthe month as number without a leading zero (1-12)
MMthe month as number with a leading zero (01-12)
MMMthe abbreviated localized month name (e.g. 'Jan' to 'Dec'). Uses the system locale to localize the name, i.e. QLocale::system().
MMMMthe long localized month name (e.g. 'January' to 'December'). Uses the system locale to localize the name, i.e. QLocale::system().
yythe year as two digit number (00-99)
yyyythe year as four digit number


These expressions may be used for the time:


 ExpressionOutput
hthe hour without a leading zero (0 to 23 or 1 to 12 if AM/PM display)
hhthe hour with a leading zero (00 to 23 or 01 to 12 if AM/PM display)
Hthe hour without a leading zero (0 to 23, even with AM/PM display)
HHthe hour with a leading zero (00 to 23, even with AM/PM display)
mthe minute without a leading zero (0 to 59)
mmthe minute with a leading zero (00 to 59)
sthe second without a leading zero (0 to 59)
ssthe second with a leading zero (00 to 59)
zthe milliseconds without leading zeroes (0 to 999)
zzzthe milliseconds with leading zeroes (000 to 999)
AP or Ause AM/PM display. A/AP will be replaced by either "AM" or "PM".
ap or ause am/pm display. a/ap will be replaced by either "am" or "pm".
tthe timezone (for example "CEST")


All other input characters will be ignored. Any sequence of characters that are enclosed in single quotes will be treated as text and not be used as an expression. Two consecutive single quotes ("''") are replaced by a singlequote in the output. Formats without separators (e.g. "HHmm") are currently not supported.

Example format strings (assumed that the QDateTime is 21 May 2001 14:13:09):


 FormatResult
dd.MM.yyyy21.05.2001
ddd MMMM d yyTue May 21 01
hh:mm:ss.zzz14:13:09.042
h:m:s ap2:13:9 pm


If the datetime is invalid, an empty string will be returned.

See also fromString(), QDate::toString(), QTime::toString(), and QLocale::toString().
*/
func (this *QDateTime) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:297
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*
Returns the datetime as a string. The format parameter determines the format of the result string.

These expressions may be used for the date:


 ExpressionOutput
dthe day as number without a leading zero (1 to 31)
ddthe day as number with a leading zero (01 to 31)
dddthe abbreviated localized day name (e.g. 'Mon' to 'Sun'). Uses the system locale to localize the name, i.e. QLocale::system().
ddddthe long localized day name (e.g. 'Monday' to 'Sunday'). Uses the system locale to localize the name, i.e. QLocale::system().
Mthe month as number without a leading zero (1-12)
MMthe month as number with a leading zero (01-12)
MMMthe abbreviated localized month name (e.g. 'Jan' to 'Dec'). Uses the system locale to localize the name, i.e. QLocale::system().
MMMMthe long localized month name (e.g. 'January' to 'December'). Uses the system locale to localize the name, i.e. QLocale::system().
yythe year as two digit number (00-99)
yyyythe year as four digit number


These expressions may be used for the time:


 ExpressionOutput
hthe hour without a leading zero (0 to 23 or 1 to 12 if AM/PM display)
hhthe hour with a leading zero (00 to 23 or 01 to 12 if AM/PM display)
Hthe hour without a leading zero (0 to 23, even with AM/PM display)
HHthe hour with a leading zero (00 to 23, even with AM/PM display)
mthe minute without a leading zero (0 to 59)
mmthe minute with a leading zero (00 to 59)
sthe second without a leading zero (0 to 59)
ssthe second with a leading zero (00 to 59)
zthe milliseconds without leading zeroes (0 to 999)
zzzthe milliseconds with leading zeroes (000 to 999)
AP or Ause AM/PM display. A/AP will be replaced by either "AM" or "PM".
ap or ause am/pm display. a/ap will be replaced by either "am" or "pm".
tthe timezone (for example "CEST")


All other input characters will be ignored. Any sequence of characters that are enclosed in single quotes will be treated as text and not be used as an expression. Two consecutive single quotes ("''") are replaced by a singlequote in the output. Formats without separators (e.g. "HHmm") are currently not supported.

Example format strings (assumed that the QDateTime is 21 May 2001 14:13:09):


 FormatResult
dd.MM.yyyy21.05.2001
ddd MMMM d yyTue May 21 01
hh:mm:ss.zzz14:13:09.042
h:m:s ap2:13:9 pm


If the datetime is invalid, an empty string will be returned.

See also fromString(), QDate::toString(), QTime::toString(), and QLocale::toString().
*/
func (this *QDateTime) ToStringp() string {
	// arg: 0, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum, , Invalid
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:298
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Returns the datetime as a string. The format parameter determines the format of the result string.

These expressions may be used for the date:


 ExpressionOutput
dthe day as number without a leading zero (1 to 31)
ddthe day as number with a leading zero (01 to 31)
dddthe abbreviated localized day name (e.g. 'Mon' to 'Sun'). Uses the system locale to localize the name, i.e. QLocale::system().
ddddthe long localized day name (e.g. 'Monday' to 'Sunday'). Uses the system locale to localize the name, i.e. QLocale::system().
Mthe month as number without a leading zero (1-12)
MMthe month as number with a leading zero (01-12)
MMMthe abbreviated localized month name (e.g. 'Jan' to 'Dec'). Uses the system locale to localize the name, i.e. QLocale::system().
MMMMthe long localized month name (e.g. 'January' to 'December'). Uses the system locale to localize the name, i.e. QLocale::system().
yythe year as two digit number (00-99)
yyyythe year as four digit number


These expressions may be used for the time:


 ExpressionOutput
hthe hour without a leading zero (0 to 23 or 1 to 12 if AM/PM display)
hhthe hour with a leading zero (00 to 23 or 01 to 12 if AM/PM display)
Hthe hour without a leading zero (0 to 23, even with AM/PM display)
HHthe hour with a leading zero (00 to 23, even with AM/PM display)
mthe minute without a leading zero (0 to 59)
mmthe minute with a leading zero (00 to 59)
sthe second without a leading zero (0 to 59)
ssthe second with a leading zero (00 to 59)
zthe milliseconds without leading zeroes (0 to 999)
zzzthe milliseconds with leading zeroes (000 to 999)
AP or Ause AM/PM display. A/AP will be replaced by either "AM" or "PM".
ap or ause am/pm display. a/ap will be replaced by either "am" or "pm".
tthe timezone (for example "CEST")


All other input characters will be ignored. Any sequence of characters that are enclosed in single quotes will be treated as text and not be used as an expression. Two consecutive single quotes ("''") are replaced by a singlequote in the output. Formats without separators (e.g. "HHmm") are currently not supported.

Example format strings (assumed that the QDateTime is 21 May 2001 14:13:09):


 FormatResult
dd.MM.yyyy21.05.2001
ddd MMMM d yyTue May 21 01
hh:mm:ss.zzz14:13:09.042
h:m:s ap2:13:9 pm


If the datetime is invalid, an empty string will be returned.

See also fromString(), QDate::toString(), QTime::toString(), and QLocale::toString().
*/
func (this *QDateTime) ToString1(format string) string {
	var tmpArg0 = NewQString5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:300
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addDays(qint64) const

/*
Returns a QDateTime object containing a datetime ndays days later than the datetime of this object (or earlier if ndays is negative).

If the timeSpec() is Qt::LocalTime and the resulting date and time fall in the Standard Time to Daylight-Saving Time transition hour then the result will be adjusted accordingly, i.e. if the transition is at 2am and the clock goes forward to 3am and the result falls between 2am and 3am then the result will be adjusted to fall after 3am.

See also daysTo(), addMonths(), addYears(), and addSecs().
*/
func (this *QDateTime) AddDays(days int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7addDaysEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), days)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:301
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addMonths(int) const

/*
Returns a QDateTime object containing a datetime nmonths months later than the datetime of this object (or earlier if nmonths is negative).

If the timeSpec() is Qt::LocalTime and the resulting date and time fall in the Standard Time to Daylight-Saving Time transition hour then the result will be adjusted accordingly, i.e. if the transition is at 2am and the clock goes forward to 3am and the result falls between 2am and 3am then the result will be adjusted to fall after 3am.

See also daysTo(), addDays(), addYears(), and addSecs().
*/
func (this *QDateTime) AddMonths(months int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime9addMonthsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), months)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:302
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addYears(int) const

/*
Returns a QDateTime object containing a datetime nyears years later than the datetime of this object (or earlier if nyears is negative).

If the timeSpec() is Qt::LocalTime and the resulting date and time fall in the Standard Time to Daylight-Saving Time transition hour then the result will be adjusted accordingly, i.e. if the transition is at 2am and the clock goes forward to 3am and the result falls between 2am and 3am then the result will be adjusted to fall after 3am.

See also daysTo(), addDays(), addMonths(), and addSecs().
*/
func (this *QDateTime) AddYears(years int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8addYearsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), years)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:303
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addSecs(qint64) const

/*
Returns a QDateTime object containing a datetime s seconds later than the datetime of this object (or earlier if s is negative).

If this datetime is invalid, an invalid datetime will be returned.

See also addMSecs(), secsTo(), addDays(), addMonths(), and addYears().
*/
func (this *QDateTime) AddSecs(secs int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7addSecsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:304
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime addMSecs(qint64) const

/*
Returns a QDateTime object containing a datetime msecs miliseconds later than the datetime of this object (or earlier if msecs is negative).

If this datetime is invalid, an invalid datetime will be returned.

See also addSecs(), msecsTo(), addDays(), addMonths(), and addYears().
*/
func (this *QDateTime) AddMSecs(msecs int64) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8addMSecsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:306
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toTimeSpec(Qt::TimeSpec) const

/*
Returns a copy of this datetime converted to the given time spec.

If spec is Qt::OffsetFromUTC then it is set to Qt::UTC. To set to a spec of Qt::OffsetFromUTC use toOffsetFromUtc().

If spec is Qt::TimeZone then it is set to Qt::LocalTime, i.e. the local Time Zone.

Example:


  QDateTime local(QDateTime::currentDateTime());
  QDateTime UTC(local.toTimeSpec(Qt::UTC));
  qDebug() << "Local time is:" << local;
  qDebug() << "UTC time is:" << UTC;
  qDebug() << "No difference between times:" << local.secsTo(UTC);



See also timeSpec(), toTimeZone(), toUTC(), and toLocalTime().
*/
func (this *QDateTime) ToTimeSpec(spec int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime10toTimeSpecEN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime toLocalTime() const

/*
Returns a datetime containing the date and time information in this datetime, but specified using the Qt::LocalTime definition.

Example:


  QDateTime UTC(QDateTime::currentDateTimeUtc());
  QDateTime local(UTC.toLocalTime());
  qDebug() << "UTC time is:" << UTC;
  qDebug() << "Local time is:" << local;
  qDebug() << "No difference between times:" << UTC.secsTo(local);



See also toTimeSpec().
*/
func (this *QDateTime) ToLocalTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime11toLocalTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:308
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime toUTC() const

/*
Returns a datetime containing the date and time information in this datetime, but specified using the Qt::UTC definition.

Example:


  QDateTime local(QDateTime::currentDateTime());
  QDateTime UTC(local.toUTC());
  qDebug() << "Local time is:" << local;
  qDebug() << "UTC time is:" << UTC;
  qDebug() << "No difference between times:" << local.secsTo(UTC);



See also toTimeSpec().
*/
func (this *QDateTime) ToUTC() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime5toUTCEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:309
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toOffsetFromUtc(int) const

/*
Returns a copy of this datetime converted to a spec of Qt::OffsetFromUTC with the given offsetSeconds.

If the offsetSeconds equals 0 then a UTC datetime will be returned

This function was introduced in  Qt 5.2.

See also setOffsetFromUtc(), offsetFromUtc(), and toTimeSpec().
*/
func (this *QDateTime) ToOffsetFromUtc(offsetSeconds int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime15toOffsetFromUtcEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offsetSeconds)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:311
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toTimeZone(const QTimeZone &) const

/*
Returns a copy of this datetime converted to the given timeZone

This function was introduced in  Qt 5.2.

See also timeZone() and toTimeSpec().
*/
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

// /usr/include/qt/QtCore/qdatetime.h:314
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 daysTo(const QDateTime &) const

/*
Returns the number of days from this datetime to the other datetime. The number of days is counted as the number of times midnight is reached between this datetime to the other datetime. This means that a 10 minute difference from 23:55 to 0:05 the next day counts as one day.

If the other datetime is earlier than this datetime, the value returned is negative.

Example:


  QDateTime startDate(QDate(2012, 7, 6), QTime(8, 30, 0));
  QDateTime endDate(QDate(2012, 7, 7), QTime(16, 30, 0));
  qDebug() << "Days from startDate to endDate: " << startDate.daysTo(endDate);

  startDate = QDateTime(QDate(2012, 7, 6), QTime(23, 55, 0));
  endDate = QDateTime(QDate(2012, 7, 7), QTime(0, 5, 0));
  qDebug() << "Days from startDate to endDate: " << startDate.daysTo(endDate);

  qSwap(startDate, endDate); // Make endDate before startDate.
  qDebug() << "Days from startDate to endDate: " << startDate.daysTo(endDate);



See also addDays(), secsTo(), and msecsTo().
*/
func (this *QDateTime) DaysTo(arg0 QDateTime_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDateTime_PTR() != nil {
		convArg0 = arg0.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime6daysToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:315
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 secsTo(const QDateTime &) const

/*
Returns the number of seconds from this datetime to the other datetime. If the other datetime is earlier than this datetime, the value returned is negative.

Before performing the comparison, the two datetimes are converted to Qt::UTC to ensure that the result is correct if daylight-saving (DST) applies to one of the two datetimes but not the other.

Returns 0 if either datetime is invalid.

Example:


  QDateTime now = QDateTime::currentDateTime();
  QDateTime xmas(QDate(now.date().year(), 12, 25), QTime(0, 0));
  qDebug("There are %d seconds to Christmas", now.secsTo(xmas));



See also addSecs(), daysTo(), and QTime::secsTo().
*/
func (this *QDateTime) SecsTo(arg0 QDateTime_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDateTime_PTR() != nil {
		convArg0 = arg0.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime6secsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:316
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 msecsTo(const QDateTime &) const

/*
Returns the number of milliseconds from this datetime to the other datetime. If the other datetime is earlier than this datetime, the value returned is negative.

Before performing the comparison, the two datetimes are converted to Qt::UTC to ensure that the result is correct if daylight-saving (DST) applies to one of the two datetimes and but not the other.

Returns 0 if either datetime is invalid.

See also addMSecs(), daysTo(), and QTime::msecsTo().
*/
func (this *QDateTime) MsecsTo(arg0 QDateTime_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDateTime_PTR() != nil {
		convArg0 = arg0.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime7msecsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:318
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QDateTime &) const

/*

 */
func (this *QDateTime) Operator_equal_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:319
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QDateTime &) const

/*

 */
func (this *QDateTime) Operator_not_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:320
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QDateTime &) const

/*

 */
func (this *QDateTime) Operator_less_than(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:321
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QDateTime &) const

/*

 */
func (this *QDateTime) Operator_less_than_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimeleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:322
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QDateTime &) const

/*

 */
func (this *QDateTime) Operator_greater_than(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimegtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:323
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QDateTime &) const

/*

 */
func (this *QDateTime) Operator_greater_than_equal(other QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDateTime_PTR() != nil {
		convArg0 = other.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTimegeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:326
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUtcOffset(int)

/*

 */
func (this *QDateTime) SetUtcOffset(seconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime12setUtcOffsetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), seconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:327
// index:0
// Public Visibility=Default Availability=Available
// [4] int utcOffset() const

/*

 */
func (this *QDateTime) UtcOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime9utcOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:330
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime currentDateTime()

/*
Returns the current datetime, as reported by the system clock, in the local time zone.

See also currentDateTimeUtc(), QDate::currentDate(), QTime::currentTime(), and toTimeSpec().
*/
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

// /usr/include/qt/QtCore/qdatetime.h:331
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime currentDateTimeUtc()

/*
Returns the current datetime, as reported by the system clock, in UTC.

This function was introduced in  Qt 4.7.

See also currentDateTime(), QDate::currentDate(), QTime::currentTime(), and toTimeSpec().
*/
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

// /usr/include/qt/QtCore/qdatetime.h:333
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromString(const QString &, Qt::DateFormat)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QDateTime) FromString(s string, f int) *QDateTime /*123*/ {
	var tmpArg0 = NewQString5(s)
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

// /usr/include/qt/QtCore/qdatetime.h:333
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromString(const QString &, Qt::DateFormat)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QDateTime) FromStringp(s string) *QDateTime /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum, , Invalid
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:334
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromString(const QString &, const QString &)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QDateTime) FromString1(s string, format string) *QDateTime /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromStringERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromString1(s string, format string) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromString1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:338
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toTime_t() const

/*

 */
func (this *QDateTime) ToTime_t() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toTime_tEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:339
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTime_t(uint)

/*

 */
func (this *QDateTime) SetTime_t(secsSince1Jan1970UTC uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime9setTime_tEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secsSince1Jan1970UTC)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:340
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint)

/*

 */
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

// /usr/include/qt/QtCore/qdatetime.h:341
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint, Qt::TimeSpec, int)

/*

 */
func (this *QDateTime) FromTime_t1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, secsSince1Jan1970UTC, spec, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromTime_t1(secsSince1Jan1970UTC uint, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t1(secsSince1Jan1970UTC, spec, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:341
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint, Qt::TimeSpec, int)

/*

 */
func (this *QDateTime) FromTime_t1p(secsSince1Jan1970UTC uint, spec int) *QDateTime /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	offsetFromUtc := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime10fromTime_tEjN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, secsSince1Jan1970UTC, spec, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:343
// index:2
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromTime_t(uint, const QTimeZone &)

/*

 */
func (this *QDateTime) FromTime_t2(secsSince1Jan1970UTC uint, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
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
func QDateTime_FromTime_t2(secsSince1Jan1970UTC uint, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromTime_t2(secsSince1Jan1970UTC, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:346
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64)

/*
Returns a datetime whose date and time are the number of milliseconds, msecs, that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC), and converted to Qt::LocalTime. On systems that do not support time zones, the time will be set as if local time were Qt::UTC.

Note that there are possible values for msecs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

This function was introduced in  Qt 4.7.

See also toMSecsSinceEpoch() and setMSecsSinceEpoch().
*/
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

// /usr/include/qt/QtCore/qdatetime.h:348
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64, Qt::TimeSpec, int)

/*
Returns a datetime whose date and time are the number of milliseconds, msecs, that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC), and converted to Qt::LocalTime. On systems that do not support time zones, the time will be set as if local time were Qt::UTC.

Note that there are possible values for msecs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

This function was introduced in  Qt 4.7.

See also toMSecsSinceEpoch() and setMSecsSinceEpoch().
*/
func (this *QDateTime) FromMSecsSinceEpoch1(msecs int64, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, msecs, spec, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}
func QDateTime_FromMSecsSinceEpoch1(msecs int64, spec int, offsetFromUtc int) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch1(msecs, spec, offsetFromUtc)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:348
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64, Qt::TimeSpec, int)

/*
Returns a datetime whose date and time are the number of milliseconds, msecs, that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC), and converted to Qt::LocalTime. On systems that do not support time zones, the time will be set as if local time were Qt::UTC.

Note that there are possible values for msecs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

This function was introduced in  Qt 4.7.

See also toMSecsSinceEpoch() and setMSecsSinceEpoch().
*/
func (this *QDateTime) FromMSecsSinceEpoch1p(msecs int64, spec int) *QDateTime /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	offsetFromUtc := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime19fromMSecsSinceEpochExN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, msecs, spec, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:352
// index:2
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromMSecsSinceEpoch(qint64, const QTimeZone &)

/*
Returns a datetime whose date and time are the number of milliseconds, msecs, that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC), and converted to Qt::LocalTime. On systems that do not support time zones, the time will be set as if local time were Qt::UTC.

Note that there are possible values for msecs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

This function was introduced in  Qt 4.7.

See also toMSecsSinceEpoch() and setMSecsSinceEpoch().
*/
func (this *QDateTime) FromMSecsSinceEpoch2(msecs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
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
func QDateTime_FromMSecsSinceEpoch2(msecs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromMSecsSinceEpoch2(msecs, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:349
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromSecsSinceEpoch(qint64, Qt::TimeSpec, int)

/*
Returns a datetime whose date and time are the number of seconds secs that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC) and converted to the given spec.

Note that there are possible values for secs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

If the spec is not Qt::OffsetFromUTC then the offsetSeconds will be ignored. If the spec is Qt::OffsetFromUTC and the offsetSeconds is 0 then the spec will be set to Qt::UTC, i.e. an offset of 0 seconds.

If spec is Qt::TimeZone then the spec will be set to Qt::LocalTime, i.e. the current system time zone.

This function was introduced in  Qt 5.8.

See also toSecsSinceEpoch() and setSecsSinceEpoch().
*/
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

// /usr/include/qt/QtCore/qdatetime.h:349
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromSecsSinceEpoch(qint64, Qt::TimeSpec, int)

/*
Returns a datetime whose date and time are the number of seconds secs that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC) and converted to the given spec.

Note that there are possible values for secs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

If the spec is not Qt::OffsetFromUTC then the offsetSeconds will be ignored. If the spec is Qt::OffsetFromUTC and the offsetSeconds is 0 then the spec will be set to Qt::UTC, i.e. an offset of 0 seconds.

If spec is Qt::TimeZone then the spec will be set to Qt::LocalTime, i.e. the current system time zone.

This function was introduced in  Qt 5.8.

See also toSecsSinceEpoch() and setSecsSinceEpoch().
*/
func (this *QDateTime) FromSecsSinceEpochp(secs int64) *QDateTime /*123*/ {
	// arg: 1, Qt::TimeSpec=Elaborated, Qt::TimeSpec=Enum, , Invalid
	spe := 0
	// arg: 2, int=Int, =Invalid, , Invalid
	offsetFromUtc := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, secs, spe, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:349
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromSecsSinceEpoch(qint64, Qt::TimeSpec, int)

/*
Returns a datetime whose date and time are the number of seconds secs that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC) and converted to the given spec.

Note that there are possible values for secs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

If the spec is not Qt::OffsetFromUTC then the offsetSeconds will be ignored. If the spec is Qt::OffsetFromUTC and the offsetSeconds is 0 then the spec will be set to Qt::UTC, i.e. an offset of 0 seconds.

If spec is Qt::TimeZone then the spec will be set to Qt::LocalTime, i.e. the current system time zone.

This function was introduced in  Qt 5.8.

See also toSecsSinceEpoch() and setSecsSinceEpoch().
*/
func (this *QDateTime) FromSecsSinceEpochp1(secs int64, spe int) *QDateTime /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	offsetFromUtc := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateTime18fromSecsSinceEpochExN2Qt8TimeSpecEi", qtrt.FFI_TYPE_POINTER, secs, spe, offsetFromUtc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:353
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDateTime fromSecsSinceEpoch(qint64, const QTimeZone &)

/*
Returns a datetime whose date and time are the number of seconds secs that have passed since 1970-01-01T00:00:00.000, Coordinated Universal Time (Qt::UTC) and converted to the given spec.

Note that there are possible values for secs that lie outside the valid range of QDateTime, both negative and positive. The behavior of this function is undefined for those values.

If the spec is not Qt::OffsetFromUTC then the offsetSeconds will be ignored. If the spec is Qt::OffsetFromUTC and the offsetSeconds is 0 then the spec will be set to Qt::UTC, i.e. an offset of 0 seconds.

If spec is Qt::TimeZone then the spec will be set to Qt::LocalTime, i.e. the current system time zone.

This function was introduced in  Qt 5.8.

See also toSecsSinceEpoch() and setSecsSinceEpoch().
*/
func (this *QDateTime) FromSecsSinceEpoch1(secs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
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
func QDateTime_FromSecsSinceEpoch1(secs int64, timeZone QTimeZone_ITF) *QDateTime /*123*/ {
	var nilthis *QDateTime
	rv := nilthis.FromSecsSinceEpoch1(secs, timeZone)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:356
// index:0
// Public static Visibility=Default Availability=Available
// [8] qint64 currentMSecsSinceEpoch()

/*
Returns the number of milliseconds since 1970-01-01T00:00:00 Universal Coordinated Time. This number is like the POSIX time_t variable, but expressed in milliseconds instead.

This function was introduced in  Qt 4.7.

See also currentDateTime(), currentDateTimeUtc(), toTime_t(), and toTimeSpec().
*/
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

// /usr/include/qt/QtCore/qdatetime.h:357
// index:0
// Public static Visibility=Default Availability=Available
// [8] qint64 currentSecsSinceEpoch()

/*
Returns the number of seconds since 1970-01-01T00:00:00 Universal Coordinated Time.

This function was introduced in  Qt 5.8.

See also currentMSecsSinceEpoch().
*/
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
