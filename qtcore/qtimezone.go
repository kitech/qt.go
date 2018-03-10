package qtcore

// /usr/include/qt/QtCore/qtimezone.h
// #include <qtimezone.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QTimeZone struct {
	*qtrt.CObject
}
type QTimeZone_ITF interface {
	QTimeZone_PTR() *QTimeZone
}

func (ptr *QTimeZone) QTimeZone_PTR() *QTimeZone { return ptr }

func (this *QTimeZone) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTimeZone) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTimeZoneFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return &QTimeZone{&qtrt.CObject{cthis}}
}
func (*QTimeZone) NewFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return NewQTimeZoneFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtimezone.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone()

/*
Create a null/invalid time zone instance.
*/
func NewQTimeZone() *QTimeZone {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:93
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(const QByteArray &)

/*
Create a null/invalid time zone instance.
*/
func NewQTimeZone_1(ianaId QByteArray_ITF) *QTimeZone {
	var convArg0 unsafe.Pointer
	if ianaId != nil && ianaId.QByteArray_PTR() != nil {
		convArg0 = ianaId.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:94
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(int)

/*
Create a null/invalid time zone instance.
*/
func NewQTimeZone_2(offsetSeconds int) *QTimeZone {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2Ei", qtrt.FFI_TYPE_POINTER, offsetSeconds)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(const QByteArray &, int, const QString &, const QString &, QLocale::Country, const QString &)

/*
Create a null/invalid time zone instance.
*/
func NewQTimeZone_3(zoneId QByteArray_ITF, offsetSeconds int, name string, abbreviation string, country int, comment string) *QTimeZone {
	var convArg0 unsafe.Pointer
	if zoneId != nil && zoneId.QByteArray_PTR() != nil {
		convArg0 = zoneId.QByteArray_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(name)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(abbreviation)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg5 = NewQString_5(comment)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", qtrt.FFI_TYPE_POINTER, convArg0, offsetSeconds, convArg2, convArg3, country, convArg5)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(const QByteArray &, int, const QString &, const QString &, QLocale::Country, const QString &)

/*
Create a null/invalid time zone instance.
*/
func NewQTimeZone_3_(zoneId QByteArray_ITF, offsetSeconds int, name string, abbreviation string) *QTimeZone {
	var convArg0 unsafe.Pointer
	if zoneId != nil && zoneId.QByteArray_PTR() != nil {
		convArg0 = zoneId.QByteArray_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(name)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(abbreviation)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, QLocale::Country=Elaborated, QLocale::Country=Enum,
	country := 0
	// arg: 5, const QString &=LValueReference, QString=Record,
	var convArg5 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", qtrt.FFI_TYPE_POINTER, convArg0, offsetSeconds, convArg2, convArg3, country, convArg5)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTimeZone(const QByteArray &, int, const QString &, const QString &, QLocale::Country, const QString &)

/*
Create a null/invalid time zone instance.
*/
func NewQTimeZone_3_1(zoneId QByteArray_ITF, offsetSeconds int, name string, abbreviation string, country int) *QTimeZone {
	var convArg0 unsafe.Pointer
	if zoneId != nil && zoneId.QByteArray_PTR() != nil {
		convArg0 = zoneId.QByteArray_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(name)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(abbreviation)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 5, const QString &=LValueReference, QString=Record,
	var convArg5 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", qtrt.FFI_TYPE_POINTER, convArg0, offsetSeconds, convArg2, convArg3, country, convArg5)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimeZone)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTimeZone()

/*

 */
func DeleteQTimeZone(this *QTimeZone) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtimezone.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QTimeZone & operator=(const QTimeZone &)

/*

 */
func (this *QTimeZone) Operator_equal(other QTimeZone_ITF) *QTimeZone {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTimeZone_PTR() != nil {
		convArg0 = other.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:103
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QTimeZone & operator=(QTimeZone &&)

/*

 */
func (this *QTimeZone) Operator_equal_1(other unsafe.Pointer /*333*/) *QTimeZone {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZoneaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QTimeZone &)

/*
Swaps this time zone instance with other. This function is very fast and never fails.
*/
func (this *QTimeZone) Swap(other QTimeZone_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTimeZone_PTR() != nil {
		convArg0 = other.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QTimeZone &) const

/*

 */
func (this *QTimeZone) Operator_equal_equal(other QTimeZone_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTimeZone_PTR() != nil {
		convArg0 = other.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZoneeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QTimeZone &) const

/*

 */
func (this *QTimeZone) Operator_not_equal(other QTimeZone_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTimeZone_PTR() != nil {
		convArg0 = other.QTimeZone_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZoneneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this time zone is valid.
*/
func (this *QTimeZone) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray id() const

/*
Returns the IANA ID for the time zone.

IANA IDs are used on all platforms. On Windows these are translated from the Windows ID into the closest IANA ID for the time zone and country.
*/
func (this *QTimeZone) Id() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone2idEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::Country country() const

/*
Returns the country for the time zone.
*/
func (this *QTimeZone) Country() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone7countryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtimezone.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString comment() const

/*
Returns any comment for the time zone.

A comment may be provided by the host platform to assist users in choosing the correct time zone. Depending on the platform this may not be localized.
*/
func (this *QTimeZone) Comment() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone7commentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayName(const QDateTime &, QTimeZone::NameType, const QLocale &) const

/*
Returns the localized time zone display name at the given atDateTime for the given nameType in the given locale. The nameType and locale requested may not be supported on all platforms, in which case the best available option will be returned.

If the locale is not provided then the application default locale will be used.

The display name may change depending on DST or historical events.

See also abbreviation().
*/
func (this *QTimeZone) DisplayName(atDateTime QDateTime_ITF, nameType int, locale QLocale_ITF) string {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg2 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, nameType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayName(const QDateTime &, QTimeZone::NameType, const QLocale &) const

/*
Returns the localized time zone display name at the given atDateTime for the given nameType in the given locale. The nameType and locale requested may not be supported on all platforms, in which case the best available option will be returned.

If the locale is not provided then the application default locale will be used.

The display name may change depending on DST or historical events.

See also abbreviation().
*/
func (this *QTimeZone) DisplayName__(atDateTime QDateTime_ITF) string {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	// arg: 1, QTimeZone::NameType=Elaborated, QTimeZone::NameType=Enum,
	nameType := 0
	// arg: 2, const QLocale &=LValueReference, QLocale=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, nameType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayName(const QDateTime &, QTimeZone::NameType, const QLocale &) const

/*
Returns the localized time zone display name at the given atDateTime for the given nameType in the given locale. The nameType and locale requested may not be supported on all platforms, in which case the best available option will be returned.

If the locale is not provided then the application default locale will be used.

The display name may change depending on DST or historical events.

See also abbreviation().
*/
func (this *QTimeZone) DisplayName__1(atDateTime QDateTime_ITF, nameType int) string {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	// arg: 2, const QLocale &=LValueReference, QLocale=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, nameType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// Public Visibility=Default Availability=Available
// [8] QString displayName(QTimeZone::TimeType, QTimeZone::NameType, const QLocale &) const

/*
Returns the localized time zone display name at the given atDateTime for the given nameType in the given locale. The nameType and locale requested may not be supported on all platforms, in which case the best available option will be returned.

If the locale is not provided then the application default locale will be used.

The display name may change depending on DST or historical events.

See also abbreviation().
*/
func (this *QTimeZone) DisplayName_1(timeType int, nameType int, locale QLocale_ITF) string {
	var convArg2 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg2 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeType, nameType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// Public Visibility=Default Availability=Available
// [8] QString displayName(QTimeZone::TimeType, QTimeZone::NameType, const QLocale &) const

/*
Returns the localized time zone display name at the given atDateTime for the given nameType in the given locale. The nameType and locale requested may not be supported on all platforms, in which case the best available option will be returned.

If the locale is not provided then the application default locale will be used.

The display name may change depending on DST or historical events.

See also abbreviation().
*/
func (this *QTimeZone) DisplayName_1_(timeType int) string {
	// arg: 1, QTimeZone::NameType=Elaborated, QTimeZone::NameType=Enum,
	nameType := 0
	// arg: 2, const QLocale &=LValueReference, QLocale=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeType, nameType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// Public Visibility=Default Availability=Available
// [8] QString displayName(QTimeZone::TimeType, QTimeZone::NameType, const QLocale &) const

/*
Returns the localized time zone display name at the given atDateTime for the given nameType in the given locale. The nameType and locale requested may not be supported on all platforms, in which case the best available option will be returned.

If the locale is not provided then the application default locale will be used.

The display name may change depending on DST or historical events.

See also abbreviation().
*/
func (this *QTimeZone) DisplayName_1_1(timeType int, nameType int) string {
	// arg: 2, const QLocale &=LValueReference, QLocale=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeType, nameType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QString abbreviation(const QDateTime &) const

/*
Returns the time zone abbreviation at the given atDateTime. The abbreviation may change depending on DST or even historical events.

Note that the abbreviation is not guaranteed to be unique to this time zone and should not be used in place of the ID or display name.

See also displayName().
*/
func (this *QTimeZone) Abbreviation(atDateTime QDateTime_ITF) string {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone12abbreviationERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtimezone.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] int offsetFromUtc(const QDateTime &) const

/*
Returns the total effective offset at the given atDateTime, i.e. the number of seconds to add to UTC to obtain the local time. This includes any DST offset that may be in effect, i.e. it is the sum of standardTimeOffset() and daylightTimeOffset() for the given datetime.

For example, for the time zone "Europe/Berlin" the standard time offset is +3600 seconds and the DST offset is +3600 seconds. During standard time offsetFromUtc() will return +3600 (UTC+01:00), and during DST it will return +7200 (UTC+02:00).

See also standardTimeOffset() and daylightTimeOffset().
*/
func (this *QTimeZone) OffsetFromUtc(atDateTime QDateTime_ITF) int {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone13offsetFromUtcERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimezone.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int standardTimeOffset(const QDateTime &) const

/*
Returns the standard time offset at the given atDateTime, i.e. the number of seconds to add to UTC to obtain the local Standard Time. This excludes any DST offset that may be in effect.

For example, for the time zone "Europe/Berlin" the standard time offset is +3600 seconds. During both standard and DST offsetFromUtc() will return +3600 (UTC+01:00).

See also offsetFromUtc() and daylightTimeOffset().
*/
func (this *QTimeZone) StandardTimeOffset(atDateTime QDateTime_ITF) int {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimezone.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int daylightTimeOffset(const QDateTime &) const

/*
Returns the daylight-saving time offset at the given atDateTime, i.e. the number of seconds to add to the standard time offset to obtain the local daylight-saving time.

For example, for the time zone "Europe/Berlin" the DST offset is +3600 seconds. During standard time daylightTimeOffset() will return 0, and when daylight-saving is in effect it will return +3600.

See also offsetFromUtc() and standardTimeOffset().
*/
func (this *QTimeZone) DaylightTimeOffset(atDateTime QDateTime_ITF) int {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimezone.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasDaylightTime() const

/*
Returns true if the time zone has practiced daylight-saving at any time.

See also isDaylightTime() and daylightTimeOffset().
*/
func (this *QTimeZone) HasDaylightTime() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone15hasDaylightTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDaylightTime(const QDateTime &) const

/*
Returns true if daylight-saving was in effect at the given atDateTime.

See also hasDaylightTime() and daylightTimeOffset().
*/
func (this *QTimeZone) IsDaylightTime(atDateTime QDateTime_ITF) bool {
	var convArg0 unsafe.Pointer
	if atDateTime != nil && atDateTime.QDateTime_PTR() != nil {
		convArg0 = atDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone14isDaylightTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:133
// index:0
// Public Visibility=Default Availability=Available
// [32] QTimeZone::OffsetData offsetData(const QDateTime &) const

/*
Returns the effective offset details at the given forDateTime. This is the equivalent of calling offsetFromUtc(), abbreviation(), etc individually but is more efficient.

See also offsetFromUtc(), standardTimeOffset(), daylightTimeOffset(), and abbreviation().
*/
func (this *QTimeZone) OffsetData(forDateTime QDateTime_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if forDateTime != nil && forDateTime.QDateTime_PTR() != nil {
		convArg0 = forDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone10offsetDataERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasTransitions() const

/*
Returns true if the system backend supports obtaining transitions.

Transitions are changes in the time-zone: these happen when DST turns on or off and when authorities alter the offsets for the time-zone.

See also nextTransition(), previousTransition(), and transitions().
*/
func (this *QTimeZone) HasTransitions() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone14hasTransitionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:136
// index:0
// Public Visibility=Default Availability=Available
// [32] QTimeZone::OffsetData nextTransition(const QDateTime &) const

/*
Returns the first time zone Transition after the given afterDateTime. This is most useful when you have a Transition time and wish to find the Transition after it.

If there is no transition after the given afterDateTime then an invalid OffsetData will be returned with an invalid QDateTime.

The given afterDateTime is exclusive.

See also hasTransitions(), previousTransition(), and transitions().
*/
func (this *QTimeZone) NextTransition(afterDateTime QDateTime_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if afterDateTime != nil && afterDateTime.QDateTime_PTR() != nil {
		convArg0 = afterDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone14nextTransitionERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:137
// index:0
// Public Visibility=Default Availability=Available
// [32] QTimeZone::OffsetData previousTransition(const QDateTime &) const

/*
Returns the first time zone Transition before the given beforeDateTime. This is most useful when you have a Transition time and wish to find the Transition before it.

If there is no transition before the given beforeDateTime then an invalid OffsetData will be returned with an invalid QDateTime.

The given beforeDateTime is exclusive.

See also hasTransitions(), nextTransition(), and transitions().
*/
func (this *QTimeZone) PreviousTransition(beforeDateTime QDateTime_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if beforeDateTime != nil && beforeDateTime.QDateTime_PTR() != nil {
		convArg0 = beforeDateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeZone18previousTransitionERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:140
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray systemTimeZoneId()

/*
Returns the current system time zone IANA ID.

On Windows this ID is translated from the Windows ID using an internal translation table and the user's selected country. As a consequence there is a small chance any Windows install may have IDs not known by Qt, in which case "UTC" will be returned.
*/
func (this *QTimeZone) SystemTimeZoneId() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone16systemTimeZoneIdEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_SystemTimeZoneId() *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.SystemTimeZoneId()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:141
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTimeZone systemTimeZone()

/*
Returns a QTimeZone object that refers to the local system time, as specified by systemTimeZoneId().

This function was introduced in  Qt 5.5.

See also utc().
*/
func (this *QTimeZone) SystemTimeZone() *QTimeZone /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone14systemTimeZoneEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}
func QTimeZone_SystemTimeZone() *QTimeZone /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.SystemTimeZone()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTimeZone utc()

/*
Returns a QTimeZone object that refers to UTC (Universal Time Coordinated).

This function was introduced in  Qt 5.5.

See also systemTimeZone().
*/
func (this *QTimeZone) Utc() *QTimeZone /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone3utcEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTimeZone)
	return rv2
}
func QTimeZone_Utc() *QTimeZone /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.Utc()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:144
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isTimeZoneIdAvailable(const QByteArray &)

/*
Returns true if a given time zone ianaId is available on this system.

See also availableTimeZoneIds().
*/
func (this *QTimeZone) IsTimeZoneIdAvailable(ianaId QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if ianaId != nil && ianaId.QByteArray_PTR() != nil {
		convArg0 = ianaId.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QTimeZone_IsTimeZoneIdAvailable(ianaId QByteArray_ITF) bool {
	var nilthis *QTimeZone
	rv := nilthis.IsTimeZoneIdAvailable(ianaId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:150
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray ianaIdToWindowsId(const QByteArray &)

/*
Returns the Windows ID equivalent to the given ianaId.

See also windowsIdToDefaultIanaId() and windowsIdToIanaIds().
*/
func (this *QTimeZone) IanaIdToWindowsId(ianaId QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if ianaId != nil && ianaId.QByteArray_PTR() != nil {
		convArg0 = ianaId.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_IanaIdToWindowsId(ianaId QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.IanaIdToWindowsId(ianaId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray windowsIdToDefaultIanaId(const QByteArray &)

/*
Returns the default IANA ID for a given windowsId.

Because a Windows ID can cover several IANA IDs in several different countries, this function returns the most frequently used IANA ID with no regard for the country and should thus be used with care. It is usually best to request the default for a specific country.

See also ianaIdToWindowsId() and windowsIdToIanaIds().
*/
func (this *QTimeZone) WindowsIdToDefaultIanaId(windowsId QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if windowsId != nil && windowsId.QByteArray_PTR() != nil {
		convArg0 = windowsId.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_WindowsIdToDefaultIanaId(windowsId QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.WindowsIdToDefaultIanaId(windowsId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:152
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray windowsIdToDefaultIanaId(const QByteArray &, QLocale::Country)

/*
Returns the default IANA ID for a given windowsId.

Because a Windows ID can cover several IANA IDs in several different countries, this function returns the most frequently used IANA ID with no regard for the country and should thus be used with care. It is usually best to request the default for a specific country.

See also ianaIdToWindowsId() and windowsIdToIanaIds().
*/
func (this *QTimeZone) WindowsIdToDefaultIanaId_1(windowsId QByteArray_ITF, country int) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if windowsId != nil && windowsId.QByteArray_PTR() != nil {
		convArg0 = windowsId.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArrayN7QLocale7CountryE", qtrt.FFI_TYPE_POINTER, convArg0, country)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QTimeZone_WindowsIdToDefaultIanaId_1(windowsId QByteArray_ITF, country int) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.WindowsIdToDefaultIanaId_1(windowsId, country)
	return rv
}

/*


 */
type QTimeZone__ = int

//
const QTimeZone__MinUtcOffsetSecs QTimeZone__ = -50400

//
const QTimeZone__MaxUtcOffsetSecs QTimeZone__ = 50400

/*
The type of time zone time, for example when requesting the name. In time zones that do not apply DST, all three values may return the same result.


*/
type QTimeZone__TimeType = int

// The standard time in a time zone, i.e. when Daylight-Saving is not in effect. For example when formatting a display name this will show something like "Pacific Standard Time".
const QTimeZone__StandardTime QTimeZone__TimeType = 0

// A time when Daylight-Saving is in effect. For example when formatting a display name this will show something like "Pacific daylight-saving time".
const QTimeZone__DaylightTime QTimeZone__TimeType = 1

// A time which is not specifically Standard or Daylight-Saving time, either an unknown time or a neutral form. For example when formatting a display name this will show something like "Pacific Time".
const QTimeZone__GenericTime QTimeZone__TimeType = 2

/*
The type of time zone name.


*/
type QTimeZone__NameType = int

// The default form of the time zone name, e.g. LongName, ShortName or OffsetName
const QTimeZone__DefaultName QTimeZone__NameType = 0

// The long form of the time zone name, e.g. "Central European Time"
const QTimeZone__LongName QTimeZone__NameType = 1

// The short form of the time zone name, usually an abbreviation, e.g. "CET"
const QTimeZone__ShortName QTimeZone__NameType = 2

//
const QTimeZone__OffsetName QTimeZone__NameType = 3

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
