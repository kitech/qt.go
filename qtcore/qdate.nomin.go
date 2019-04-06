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
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qdatetime.h:86
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

// /usr/include/qt/QtCore/qdatetime.h:86
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

// /usr/include/qt/QtCore/qdatetime.h:88
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

// /usr/include/qt/QtCore/qdatetime.h:88
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

// /usr/include/qt/QtCore/qdatetime.h:90
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

// /usr/include/qt/QtCore/qdatetime.h:90
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

// /usr/include/qt/QtCore/qdatetime.h:92
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

// /usr/include/qt/QtCore/qdatetime.h:92
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

// /usr/include/qt/QtCore/qdatetime.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*

 */
func (this *QDate) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*

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

// /usr/include/qt/QtCore/qdatetime.h:97
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*

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
// index:2
// Public Visibility=Default Availability=Available
// [8] QString toString(QStringView) const

/*

 */
func (this *QDate) ToString2(format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg0 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDate8toStringE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:127
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

// /usr/include/qt/QtCore/qdatetime.h:127
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

// /usr/include/qt/QtCore/qdatetime.h:128
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

//  body block end

//  keep block begin

func init_unused_10320() {
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
