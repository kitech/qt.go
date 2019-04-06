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
// extern C begin: 39
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qdatetime.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*

 */
func (this *QTime) ToString(f int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(Qt::DateFormat) const

/*

 */
func (this *QTime) ToStringp() string {
	// arg: 0, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum, , Invalid
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringEN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:173
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*

 */
func (this *QTime) ToString1(format string) string {
	var tmpArg0 = NewQString5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:175
// index:2
// Public Visibility=Default Availability=Available
// [8] QString toString(QStringView) const

/*

 */
func (this *QTime) ToString2(format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg0 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QTime8toStringE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdatetime.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [4] QTime fromString(const QString &, Qt::DateFormat)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QTime) FromString(s string, f int) *QTime /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}
func QTime_FromString(s string, f int) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromString(s, f)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [4] QTime fromString(const QString &, Qt::DateFormat)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QTime) FromStringp(s string) *QTime /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::DateFormat=Elaborated, Qt::DateFormat=Enum, , Invalid
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringN2Qt10DateFormatE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:197
// index:1
// Public static Visibility=Default Availability=Available
// [4] QTime fromString(const QString &, const QString &)

/*
Returns the QDateTime represented by the string, using the format given, or an invalid datetime if this is not possible.

Note for Qt::TextDate: It is recommended that you use the English short month names (e.g. "Jan"). Although localized month names can also be used, they depend on the user's locale settings.

See also toString() and QLocale::toDateTime().
*/
func (this *QTime) FromString1(s string, format string) *QTime /*123*/ {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}
func QTime_FromString1(s string, format string) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromString1(s, format)
	return rv
}

//  body block end

//  keep block begin

func init_unused_10322() {
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
