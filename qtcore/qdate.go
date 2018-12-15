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
// extern C begin: 49
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
type QDate struct {
	*qtrt.CObject
}
type QDate_ITF interface {
	QDate_PTR() *QDate
}

func (ptr *QDate) QDate_PTR() *QDate { return ptr }

func (this *QDate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDate) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDateFromPointer(cthis unsafe.Pointer) *QDate {
	return &QDate{&qtrt.CObject{cthis}}
}
func (*QDate) NewFromPointer(cthis unsafe.Pointer) *QDate {
	return NewQDateFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdatetime.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QDate()

/*

 */
func (*QDate) NewForInherit() *QDate {
	return NewQDate()
}
func NewQDate() *QDate {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDateC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDate)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDate(int, int, int)

/*

 */
func (*QDate) NewForInherit1(y int, m int, d int) *QDate {
	return NewQDate1(y, m, d)
}
func NewQDate1(y int, m int, d int) *QDate {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDateC2Eiii", qtrt.FFI_TYPE_POINTER, y, m, d)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDate)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if both the date and the time are null; otherwise returns false. A null datetime is invalid.

See also QDate::isNull(), QTime::isNull(), and isValid().
*/
func (this *QDate) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if both the date and the time are valid and they are valid in the current Qt::TimeSpec, otherwise returns false.

If the timeSpec() is Qt::LocalTime or Qt::TimeZone then the date and time are checked to see if they fall in the Standard Time to Daylight-Saving Time transition hour, i.e. if the transition is at 2am and the clock goes forward to 3am then the time from 02:00:00 to 02:59:59.999 is considered to be invalid.

See also QDate::isValid() and QTime::isValid().
*/
func (this *QDate) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:123
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isValid(int, int, int)

/*
Returns true if both the date and the time are valid and they are valid in the current Qt::TimeSpec, otherwise returns false.

If the timeSpec() is Qt::LocalTime or Qt::TimeZone then the date and time are checked to see if they fall in the Standard Time to Daylight-Saving Time transition hour, i.e. if the transition is at 2am and the clock goes forward to 3am then the time from 02:00:00 to 02:59:59.999 is considered to be invalid.

See also QDate::isValid() and QTime::isValid().
*/
func (this *QDate) IsValid1(y int, m int, d int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate7isValidEiii", qtrt.FFI_TYPE_POINTER, y, m, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDate_IsValid1(y int, m int, d int) bool {
	var nilthis *QDate
	rv := nilthis.IsValid1(y, m, d)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int year() const

/*

 */
func (this *QDate) Year() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate4yearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int month() const

/*

 */
func (this *QDate) Month() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate5monthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int day() const

/*

 */
func (this *QDate) Day() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate3dayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int dayOfWeek() const

/*

 */
func (this *QDate) DayOfWeek() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate9dayOfWeekEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int dayOfYear() const

/*

 */
func (this *QDate) DayOfYear() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate9dayOfYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] int daysInMonth() const

/*

 */
func (this *QDate) DaysInMonth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate11daysInMonthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int daysInYear() const

/*

 */
func (this *QDate) DaysInYear() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate10daysInYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int weekNumber(int *) const

/*

 */
func (this *QDate) WeekNumber(yearNum unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate10weekNumberEPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), yearNum)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int weekNumber(int *) const

/*

 */
func (this *QDate) WeekNumberp() int {
	// arg: 0, int *=Pointer, =Invalid, , Invalid
	var yearNum unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate10weekNumberEPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), yearNum)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:85
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString shortMonthName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) ShortMonthName(month int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate14shortMonthNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, month, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_ShortMonthName(month int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.ShortMonthName(month, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:85
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString shortMonthName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) ShortMonthNamep(month int) string {
	// arg: 1, QDate::MonthNameType=Enum, QDate::MonthNameType=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate14shortMonthNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, month, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString shortDayName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) ShortDayName(weekday int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate12shortDayNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, weekday, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_ShortDayName(weekday int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.ShortDayName(weekday, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString shortDayName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) ShortDayNamep(weekday int) string {
	// arg: 1, QDate::MonthNameType=Enum, QDate::MonthNameType=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate12shortDayNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, weekday, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString longMonthName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) LongMonthName(month int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate13longMonthNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, month, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_LongMonthName(month int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.LongMonthName(month, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString longMonthName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) LongMonthNamep(month int) string {
	// arg: 1, QDate::MonthNameType=Enum, QDate::MonthNameType=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate13longMonthNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, month, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:88
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString longDayName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) LongDayName(weekday int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate11longDayNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, weekday, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_LongDayName(weekday int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.LongDayName(weekday, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:88
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString longDayName(int, QDate::MonthNameType)

/*

 */
func (this *QDate) LongDayNamep(weekday int) string {
	// arg: 1, QDate::MonthNameType=Enum, QDate::MonthNameType=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate11longDayNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, weekday, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:91
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
func (this *QDate) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:91
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
func (this *QDate) ToStringp() string {
	// arg: 0, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum, , Invalid
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:92
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
func (this *QDate) ToString1(format string) string {
	var tmpArg0 = NewQString5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setDate(int, int, int)

/*
Sets the date part of this datetime to date. If no time is set yet, it is set to midnight. If date is invalid, this QDateTime becomes invalid.

See also date(), setTime(), and setTimeSpec().
*/
func (this *QDate) SetDate(year int, month int, day int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate7setDateEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month, day)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getDate(int *, int *, int *)

/*

 */
func (this *QDate) GetDate(year unsafe.Pointer /*666*/, month unsafe.Pointer /*666*/, day unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate7getDateEPiS0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month, day)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:104
// index:1
// Public Visibility=Default Availability=Available
// [-2] void getDate(int *, int *, int *) const

/*

 */
func (this *QDate) GetDate1(year unsafe.Pointer /*666*/, month unsafe.Pointer /*666*/, day unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate7getDateEPiS0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month, day)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate addDays(qint64) const

/*
Returns a QDateTime object containing a datetime ndays days later than the datetime of this object (or earlier if ndays is negative).

If the timeSpec() is Qt::LocalTime and the resulting date and time fall in the Standard Time to Daylight-Saving Time transition hour then the result will be adjusted accordingly, i.e. if the transition is at 2am and the clock goes forward to 3am and the result falls between 2am and 3am then the result will be adjusted to fall after 3am.

See also daysTo(), addMonths(), addYears(), and addSecs().
*/
func (this *QDate) AddDays(days int64) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate7addDaysEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), days)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate addMonths(int) const

/*
Returns a QDateTime object containing a datetime nmonths months later than the datetime of this object (or earlier if nmonths is negative).

If the timeSpec() is Qt::LocalTime and the resulting date and time fall in the Standard Time to Daylight-Saving Time transition hour then the result will be adjusted accordingly, i.e. if the transition is at 2am and the clock goes forward to 3am and the result falls between 2am and 3am then the result will be adjusted to fall after 3am.

See also daysTo(), addDays(), addYears(), and addSecs().
*/
func (this *QDate) AddMonths(months int) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate9addMonthsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), months)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate addYears(int) const

/*
Returns a QDateTime object containing a datetime nyears years later than the datetime of this object (or earlier if nyears is negative).

If the timeSpec() is Qt::LocalTime and the resulting date and time fall in the Standard Time to Daylight-Saving Time transition hour then the result will be adjusted accordingly, i.e. if the transition is at 2am and the clock goes forward to 3am and the result falls between 2am and 3am then the result will be adjusted to fall after 3am.

See also daysTo(), addDays(), addMonths(), and addSecs().
*/
func (this *QDate) AddYears(years int) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8addYearsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), years)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 daysTo(const QDate &) const

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
func (this *QDate) DaysTo(arg0 QDate_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDate_PTR() != nil {
		convArg0 = arg0.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate6daysToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QDate &) const

/*

 */
func (this *QDate) Operator_equal_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QDate &) const

/*

 */
func (this *QDate) Operator_not_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QDate &) const

/*

 */
func (this *QDate) Operator_less_than(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QDate &) const

/*

 */
func (this *QDate) Operator_less_than_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QDate &) const

/*

 */
func (this *QDate) Operator_greater_than(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDategtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QDate &) const

/*

 */
func (this *QDate) Operator_greater_than_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDategeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDate currentDate()

/*

 */
func (this *QDate) CurrentDate() *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate11currentDateEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}
func QDate_CurrentDate() *QDate /*123*/ {
	var nilthis *QDate
	rv := nilthis.CurrentDate()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDate fromString(const QString &, Qt::DateFormat)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QDate) FromString(s string, f int) *QDate /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}
func QDate_FromString(s string, f int) *QDate /*123*/ {
	var nilthis *QDate
	rv := nilthis.FromString(s, f)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDate fromString(const QString &, Qt::DateFormat)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QDate) FromStringp(s string) *QDate /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum, , Invalid
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:121
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDate fromString(const QString &, const QString &)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QDate) FromString1(s string, format string) *QDate /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate10fromStringERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}
func QDate_FromString1(s string, format string) *QDate /*123*/ {
	var nilthis *QDate
	rv := nilthis.FromString1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isLeapYear(int)

/*

 */
func (this *QDate) IsLeapYear(year int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate10isLeapYearEi", qtrt.FFI_TYPE_POINTER, year)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDate_IsLeapYear(year int) bool {
	var nilthis *QDate
	rv := nilthis.IsLeapYear(year)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:126
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDate fromJulianDay(qint64)

/*

 */
func (this *QDate) FromJulianDay(jd_ int64) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate13fromJulianDayEx", qtrt.FFI_TYPE_POINTER, jd_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}
func QDate_FromJulianDay(jd_ int64) *QDate /*123*/ {
	var nilthis *QDate
	rv := nilthis.FromJulianDay(jd_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toJulianDay() const

/*

 */
func (this *QDate) ToJulianDay() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate11toJulianDayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

func DeleteQDate(this *QDate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QDate__MonthNameType = int

//
const QDate__DateFormat QDate__MonthNameType = 0

//
const QDate__StandaloneFormat QDate__MonthNameType = 1

func (this *QDate) MonthNameTypeItemName(val int) string {
	switch val {
	case QDate__DateFormat: // 0
		return "DateFormat"
	case QDate__StandaloneFormat: // 1
		return "StandaloneFormat"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDate_MonthNameTypeItemName(val int) string {
	var nilthis *QDate
	return nilthis.MonthNameTypeItemName(val)
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
