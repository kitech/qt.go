//  header block begin
// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QDate struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qdatetime.h:69
// index:0
// inline
// void QDate()
func NewQDate() *QDate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDateC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QDate{cthis}
}

// /usr/include/qt/QtCore/qdatetime.h:70
// index:1
// void QDate(int, int, int)
func NewQDate_1(y int, m int, d int) *QDate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDateC2Eiii", ffiqt.FFI_TYPE_VOID, cthis, &y, &m, &d)
	gopp.ErrPrint(err, rv)
	return &QDate{cthis}
}

// /usr/include/qt/QtCore/qdatetime.h:72
// index:0
// inline
// bool isNull()
func (this *QDate) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:73
// index:0
// inline
// bool isValid()
func (this *QDate) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:130
// index:1
// static
// bool isValid(int, int, int)
func (this *QDate) IsValid_1(y int, m int, d int) {
	// 1: (int y, int m, int d), (y, m, d)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate7isValidEiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_IsValid_1(y int, m int, d int) {
	// 1: (int y, int m, int d), (y, m, d)
	var nilthis *QDate
	nilthis.IsValid_1(y, m, d)
}

// /usr/include/qt/QtCore/qdatetime.h:75
// index:0
// int year()
func (this *QDate) Year() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate4yearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:76
// index:0
// int month()
func (this *QDate) Month() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate5monthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:77
// index:0
// int day()
func (this *QDate) Day() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate3dayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:78
// index:0
// int dayOfWeek()
func (this *QDate) DayOfWeek() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate9dayOfWeekEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:79
// index:0
// int dayOfYear()
func (this *QDate) DayOfYear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate9dayOfYearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:80
// index:0
// int daysInMonth()
func (this *QDate) DaysInMonth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate11daysInMonthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:81
// index:0
// int daysInYear()
func (this *QDate) DaysInYear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate10daysInYearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:82
// index:0
// int weekNumber(int *)
func (this *QDate) WeekNumber(yearNum unsafe.Pointer) {
	// 0: (, int * yearNum), (yearNum)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate10weekNumberEPi", ffiqt.FFI_TYPE_VOID, this.cthis, yearNum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:86
// index:0
// static
// QString shortMonthName(int, enum QDate::MonthNameType)
func (this *QDate) ShortMonthName(month int, type_ int) {
	// 0: (int month, QDate::MonthNameType type), (month, type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate14shortMonthNameEiNS_13MonthNameTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_ShortMonthName(month int, type_ int) {
	// 0: (int month, QDate::MonthNameType type), (month, type_)
	var nilthis *QDate
	nilthis.ShortMonthName(month, type_)
}

// /usr/include/qt/QtCore/qdatetime.h:88
// index:0
// static
// QString shortDayName(int, enum QDate::MonthNameType)
func (this *QDate) ShortDayName(weekday int, type_ int) {
	// 0: (int weekday, QDate::MonthNameType type), (weekday, type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate12shortDayNameEiNS_13MonthNameTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_ShortDayName(weekday int, type_ int) {
	// 0: (int weekday, QDate::MonthNameType type), (weekday, type_)
	var nilthis *QDate
	nilthis.ShortDayName(weekday, type_)
}

// /usr/include/qt/QtCore/qdatetime.h:90
// index:0
// static
// QString longMonthName(int, enum QDate::MonthNameType)
func (this *QDate) LongMonthName(month int, type_ int) {
	// 0: (int month, QDate::MonthNameType type), (month, type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate13longMonthNameEiNS_13MonthNameTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_LongMonthName(month int, type_ int) {
	// 0: (int month, QDate::MonthNameType type), (month, type_)
	var nilthis *QDate
	nilthis.LongMonthName(month, type_)
}

// /usr/include/qt/QtCore/qdatetime.h:92
// index:0
// static
// QString longDayName(int, enum QDate::MonthNameType)
func (this *QDate) LongDayName(weekday int, type_ int) {
	// 0: (int weekday, QDate::MonthNameType type), (weekday, type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate11longDayNameEiNS_13MonthNameTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_LongDayName(weekday int, type_ int) {
	// 0: (int weekday, QDate::MonthNameType type), (weekday, type_)
	var nilthis *QDate
	nilthis.LongDayName(weekday, type_)
}

// /usr/include/qt/QtCore/qdatetime.h:95
// index:0
// QString toString(Qt::DateFormat)
func (this *QDate) ToString(f int) {
	// 0: (, Qt::DateFormat f), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate8toStringEN2Qt10DateFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:97
// index:1
// QString toString(const class QString &)
func (this *QDate) ToString_1(format unsafe.Pointer) {
	// 1: (, const QString & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate8toStringERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:99
// index:2
// QString toString(class QStringView)
func (this *QDate) ToString_2(format unsafe.Pointer) {
	// 2: (, QStringView format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate8toStringE11QStringView", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:106
// index:0
// bool setDate(int, int, int)
func (this *QDate) SetDate(year int, month int, day int) {
	// 0: (, int year, int month, int day), (&year, &month, &day)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate7setDateEiii", ffiqt.FFI_TYPE_VOID, this.cthis, &year, &month, &day)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:109
// index:0
// void getDate(int *, int *, int *)
func (this *QDate) GetDate(year unsafe.Pointer, month unsafe.Pointer, day unsafe.Pointer) {
	// 0: (, int * year, int * month, int * day), (year, month, day)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate7getDateEPiS0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, year, month, day)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:111
// index:1
// void getDate(int *, int *, int *)
func (this *QDate) GetDate_1(year unsafe.Pointer, month unsafe.Pointer, day unsafe.Pointer) {
	// 1: (, int * year, int * month, int * day), (year, month, day)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate7getDateEPiS0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, year, month, day)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:113
// index:0
// QDate addDays(qint64)
func (this *QDate) AddDays(days int64) {
	// 0: (, qint64 days), (&days)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate7addDaysEx", ffiqt.FFI_TYPE_VOID, this.cthis, &days)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:114
// index:0
// QDate addMonths(int)
func (this *QDate) AddMonths(months int) {
	// 0: (, int months), (&months)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate9addMonthsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &months)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:115
// index:0
// QDate addYears(int)
func (this *QDate) AddYears(years int) {
	// 0: (, int years), (&years)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate8addYearsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &years)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:116
// index:0
// qint64 daysTo(const class QDate &)
func (this *QDate) DaysTo(arg0 unsafe.Pointer) {
	// 0: (, const QDate & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate6daysToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:125
// index:0
// static
// QDate currentDate()
func (this *QDate) CurrentDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate11currentDateEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_CurrentDate() {
	// 0: (), ()
	var nilthis *QDate
	nilthis.CurrentDate()
}

// /usr/include/qt/QtCore/qdatetime.h:127
// index:0
// static
// QDate fromString(const class QString &, Qt::DateFormat)
func (this *QDate) FromString(s unsafe.Pointer, f int) {
	// 0: (const QString & s, Qt::DateFormat f), (s, f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate10fromStringERK7QStringN2Qt10DateFormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_FromString(s unsafe.Pointer, f int) {
	// 0: (const QString & s, Qt::DateFormat f), (s, f)
	var nilthis *QDate
	nilthis.FromString(s, f)
}

// /usr/include/qt/QtCore/qdatetime.h:128
// index:1
// static
// QDate fromString(const class QString &, const class QString &)
func (this *QDate) FromString_1(s unsafe.Pointer, format unsafe.Pointer) {
	// 1: (const QString & s, const QString & format), (s, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate10fromStringERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_FromString_1(s unsafe.Pointer, format unsafe.Pointer) {
	// 1: (const QString & s, const QString & format), (s, format)
	var nilthis *QDate
	nilthis.FromString_1(s, format)
}

// /usr/include/qt/QtCore/qdatetime.h:131
// index:0
// static
// bool isLeapYear(int)
func (this *QDate) IsLeapYear(year int) {
	// 0: (int year), (year)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate10isLeapYearEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_IsLeapYear(year int) {
	// 0: (int year), (year)
	var nilthis *QDate
	nilthis.IsLeapYear(year)
}

// /usr/include/qt/QtCore/qdatetime.h:133
// index:0
// static inline
// QDate fromJulianDay(qint64)
func (this *QDate) FromJulianDay(jd_ int64) {
	// 0: (qint64 jd_), (jd_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDate13fromJulianDayEx", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDate_FromJulianDay(jd_ int64) {
	// 0: (qint64 jd_), (jd_)
	var nilthis *QDate
	nilthis.FromJulianDay(jd_)
}

// /usr/include/qt/QtCore/qdatetime.h:135
// index:0
// inline
// qint64 toJulianDay()
func (this *QDate) ToJulianDay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDate11toJulianDayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
