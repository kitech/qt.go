//  header block begin

// +build !minimal

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

// /usr/include/qt/QtCore/qdatetime.h:267
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

// /usr/include/qt/QtCore/qdatetime.h:288
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

// /usr/include/qt/QtCore/qdatetime.h:301
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

// /usr/include/qt/QtCore/qdatetime.h:307
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*

 */
func (this *QDateTime) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:307
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*

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

// /usr/include/qt/QtCore/qdatetime.h:309
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*

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

// /usr/include/qt/QtCore/qdatetime.h:311
// index:2
// Public Visibility=Default Availability=Available
// [8] QString toString(QStringView) const

/*

 */
func (this *QDateTime) ToString2(format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg0 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateTime8toStringE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:324
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

// /usr/include/qt/QtCore/qdatetime.h:346
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

// /usr/include/qt/QtCore/qdatetime.h:346
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

// /usr/include/qt/QtCore/qdatetime.h:347
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

// /usr/include/qt/QtCore/qdatetime.h:365
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

// /usr/include/qt/QtCore/qdatetime.h:366
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

//  body block end

//  keep block begin

func init_unused_10324() {
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
