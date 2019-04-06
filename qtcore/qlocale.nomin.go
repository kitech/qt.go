//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qlocale.h
// #include <qlocale.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qlocale.h:1008
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate toDate(const QString &, QLocale::FormatType) const

/*
Parses the date string given in string and returns the date. The format of the date string is chosen according to the format parameter (see dateFormat()).

If the date could not be parsed, returns an invalid date.

This function was introduced in  Qt 4.4.

See also dateFormat(), toTime(), toDateTime(), and QDate::fromString().
*/
func (this *QLocale) ToDate(string string, arg1 int) *QDate /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1008
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate toDate(const QString &, QLocale::FormatType) const

/*
Parses the date string given in string and returns the date. The format of the date string is chosen according to the format parameter (see dateFormat()).

If the date could not be parsed, returns an invalid date.

This function was introduced in  Qt 4.4.

See also dateFormat(), toTime(), toDateTime(), and QDate::fromString().
*/
func (this *QLocale) ToDatep(string string) *QDate /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	arg1 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1011
// index:1
// Public Visibility=Default Availability=Available
// [8] QDate toDate(const QString &, const QString &) const

/*
Parses the date string given in string and returns the date. The format of the date string is chosen according to the format parameter (see dateFormat()).

If the date could not be parsed, returns an invalid date.

This function was introduced in  Qt 4.4.

See also dateFormat(), toTime(), toDateTime(), and QDate::fromString().
*/
func (this *QLocale) ToDate1(string string, format string) *QDate /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1009
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime toTime(const QString &, QLocale::FormatType) const

/*
Parses the time string given in string and returns the time. The format of the time string is chosen according to the format parameter (see timeFormat()).

If the time could not be parsed, returns an invalid time.

This function was introduced in  Qt 4.4.

See also timeFormat(), toDate(), toDateTime(), and QTime::fromString().
*/
func (this *QLocale) ToTime(string string, arg1 int) *QTime /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1009
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime toTime(const QString &, QLocale::FormatType) const

/*
Parses the time string given in string and returns the time. The format of the time string is chosen according to the format parameter (see timeFormat()).

If the time could not be parsed, returns an invalid time.

This function was introduced in  Qt 4.4.

See also timeFormat(), toDate(), toDateTime(), and QTime::fromString().
*/
func (this *QLocale) ToTimep(string string) *QTime /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	arg1 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1012
// index:1
// Public Visibility=Default Availability=Available
// [4] QTime toTime(const QString &, const QString &) const

/*
Parses the time string given in string and returns the time. The format of the time string is chosen according to the format parameter (see timeFormat()).

If the time could not be parsed, returns an invalid time.

This function was introduced in  Qt 4.4.

See also timeFormat(), toDate(), toDateTime(), and QTime::fromString().
*/
func (this *QLocale) ToTime1(string string, format string) *QTime /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1010
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QString &, QLocale::FormatType) const

/*
Parses the date/time string given in string and returns the time. The format of the date/time string is chosen according to the format parameter (see dateTimeFormat()).

If the string could not be parsed, returns an invalid QDateTime.

This function was introduced in  Qt 4.4.

See also dateTimeFormat(), toTime(), toDate(), and QDateTime::fromString().
*/
func (this *QLocale) ToDateTime(string string, format int) *QDateTime /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1010
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QString &, QLocale::FormatType) const

/*
Parses the date/time string given in string and returns the time. The format of the date/time string is chosen according to the format parameter (see dateTimeFormat()).

If the string could not be parsed, returns an invalid QDateTime.

This function was introduced in  Qt 4.4.

See also dateTimeFormat(), toTime(), toDate(), and QDateTime::fromString().
*/
func (this *QLocale) ToDateTimep(string string) *QDateTime /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1013
// index:1
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QString &, const QString &) const

/*
Parses the date/time string given in string and returns the time. The format of the date/time string is chosen according to the format parameter (see dateTimeFormat()).

If the string could not be parsed, returns an invalid QDateTime.

This function was introduced in  Qt 4.4.

See also dateTimeFormat(), toTime(), toDate(), and QDateTime::fromString().
*/
func (this *QLocale) ToDateTime1(string string, format string) *QDateTime /*123*/ {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10326() {
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
