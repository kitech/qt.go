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
func NewQDate_1(y int, m int, d int) *QDate {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDateC2Eiii", qtrt.FFI_TYPE_POINTER, y, m, d)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDate)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QDate) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QDate) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:130
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isValid(int, int, int)
func (this *QDate) IsValid_1(y int, m int, d int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate7isValidEiii", qtrt.FFI_TYPE_POINTER, y, m, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDate_IsValid_1(y int, m int, d int) bool {
	var nilthis *QDate
	rv := nilthis.IsValid_1(y, m, d)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int year()
func (this *QDate) Year() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate4yearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int month()
func (this *QDate) Month() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate5monthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int day()
func (this *QDate) Day() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate3dayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int dayOfWeek()
func (this *QDate) DayOfWeek() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate9dayOfWeekEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int dayOfYear()
func (this *QDate) DayOfYear() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate9dayOfYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] int daysInMonth()
func (this *QDate) DaysInMonth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate11daysInMonthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int daysInYear()
func (this *QDate) DaysInYear() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate10daysInYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int weekNumber(int *)
func (this *QDate) WeekNumber(yearNum unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate10weekNumberEPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), yearNum)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatetime.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString shortMonthName(int, enum QDate::MonthNameType)
func (this *QDate) ShortMonthName(month int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate14shortMonthNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, month, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_ShortMonthName(month int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.ShortMonthName(month, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:88
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString shortDayName(int, enum QDate::MonthNameType)
func (this *QDate) ShortDayName(weekday int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate12shortDayNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, weekday, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_ShortDayName(weekday int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.ShortDayName(weekday, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString longMonthName(int, enum QDate::MonthNameType)
func (this *QDate) LongMonthName(month int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate13longMonthNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, month, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_LongMonthName(month int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.LongMonthName(month, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:92
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString longDayName(int, enum QDate::MonthNameType)
func (this *QDate) LongDayName(weekday int, type_ int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate11longDayNameEiNS_13MonthNameTypeE", qtrt.FFI_TYPE_POINTER, weekday, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDate_LongDayName(weekday int, type_ int) string {
	var nilthis *QDate
	rv := nilthis.LongDayName(weekday, type_)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat)
func (this *QDate) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:97
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &)
func (this *QDate) ToString_1(format string) string {
	var tmpArg0 = NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:99
// index:2
// Public Visibility=Default Availability=Available
// [8] QString toString(QStringView)
func (this *QDate) ToString_2(format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg0 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setDate(int, int, int)
func (this *QDate) SetDate(year int, month int, day int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate7setDateEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month, day)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getDate(int *, int *, int *)
func (this *QDate) GetDate(year unsafe.Pointer /*666*/, month unsafe.Pointer /*666*/, day unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate7getDateEPiS0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month, day)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:111
// index:1
// Public Visibility=Default Availability=Available
// [-2] void getDate(int *, int *, int *)
func (this *QDate) GetDate_1(year unsafe.Pointer /*666*/, month unsafe.Pointer /*666*/, day unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate7getDateEPiS0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month, day)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate addDays(qint64)
func (this *QDate) AddDays(days int64) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate7addDaysEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), days)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate addMonths(int)
func (this *QDate) AddMonths(months int) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate9addMonthsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), months)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate addYears(int)
func (this *QDate) AddYears(years int) *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8addYearsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), years)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 daysTo(const QDate &)
func (this *QDate) DaysTo(arg0 QDate_ITF) int64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDate_PTR() != nil {
		convArg0 = arg0.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate6daysToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QDate &)
func (this *QDate) Operator_equal_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QDate &)
func (this *QDate) Operator_not_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QDate &)
func (this *QDate) Operator_less_than(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QDate &)
func (this *QDate) Operator_less_than_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDateleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QDate &)
func (this *QDate) Operator_greater_than(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDategtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QDate &)
func (this *QDate) Operator_greater_than_equal(other QDate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDate_PTR() != nil {
		convArg0 = other.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDategeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDate currentDate()
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

// /usr/include/qt/QtCore/qdatetime.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDate fromString(const QString &, Qt::DateFormat)
func (this *QDate) FromString(s string, f int) *QDate /*123*/ {
	var tmpArg0 = NewQString_5(s)
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

// /usr/include/qt/QtCore/qdatetime.h:128
// index:1
// Public static Visibility=Default Availability=Available
// [8] QDate fromString(const QString &, const QString &)
func (this *QDate) FromString_1(s string, format string) *QDate /*123*/ {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDate10fromStringERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}
func QDate_FromString_1(s string, format string) *QDate /*123*/ {
	var nilthis *QDate
	rv := nilthis.FromString_1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isLeapYear(int)
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

// /usr/include/qt/QtCore/qdatetime.h:133
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDate fromJulianDay(qint64)
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

// /usr/include/qt/QtCore/qdatetime.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toJulianDay()
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

type QDate__MonthNameType = int

const QDate__DateFormat QDate__MonthNameType = 0
const QDate__StandaloneFormat QDate__MonthNameType = 1

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
