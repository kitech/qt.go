package qtcore

// /usr/include/qt/QtCore/qtimezone.h
// #include <qtimezone.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
type QTimeZone struct {
	*qtrt.CObject
}

func (this *QTimeZone) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTimeZone) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTimeZoneFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return &QTimeZone{&qtrt.CObject{cthis}}
}
func (*QTimeZone) NewFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return NewQTimeZoneFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtimezone.h:92
// index:0
// Public
// void QTimeZone()
func NewQTimeZone() *QTimeZone {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:93
// index:1
// Public
// void QTimeZone(const QByteArray &)
func NewQTimeZone_1(ianaId *QByteArray) *QTimeZone {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = ianaId.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:94
// index:2
// Public
// void QTimeZone(int)
func NewQTimeZone_2(offsetSeconds int) *QTimeZone {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2Ei", ffiqt.FFI_TYPE_VOID, cthis, offsetSeconds)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// Public
// void QTimeZone(const QByteArray &, int, const QString &, const QString &, QLocale::Country, const QString &)
func NewQTimeZone_3(zoneId *QByteArray, offsetSeconds int, name *QString, abbreviation *QString, country int, comment *QString) *QTimeZone {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = zoneId.GetCthis()
	var convArg2 = name.GetCthis()
	var convArg3 = abbreviation.GetCthis()
	var convArg5 = comment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, offsetSeconds, convArg2, convArg3, country, convArg5)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:99
// index:0
// Public
// void ~QTimeZone()
func DeleteQTimeZone(*QTimeZone) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:106
// index:0
// Public inline
// void swap(QTimeZone &)
func (this *QTimeZone) Swap(other *QTimeZone) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:112
// index:0
// Public
// bool isValid()
func (this *QTimeZone) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:114
// index:0
// Public
// QByteArray id()
func (this *QTimeZone) Id() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone2idEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:115
// index:0
// Public
// QLocale::Country country()
func (this *QTimeZone) Country() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7countryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimezone.h:116
// index:0
// Public
// QString comment()
func (this *QTimeZone) Comment() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7commentEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// Public
// QString displayName(const QDateTime &, QTimeZone::NameType, const QLocale &)
func (this *QTimeZone) DisplayName(atDateTime *QDateTime, nameType int, locale *QLocale) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = atDateTime.GetCthis()
	var convArg2 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, nameType, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// Public
// QString displayName(QTimeZone::TimeType, QTimeZone::NameType, const QLocale &)
func (this *QTimeZone) DisplayName_1(timeType int, nameType int, locale *QLocale) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), timeType, nameType, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:124
// index:0
// Public
// QString abbreviation(const QDateTime &)
func (this *QTimeZone) Abbreviation(atDateTime *QDateTime) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone12abbreviationERK9QDateTime", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtimezone.h:126
// index:0
// Public
// int offsetFromUtc(const QDateTime &)
func (this *QTimeZone) OffsetFromUtc(atDateTime *QDateTime) int {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone13offsetFromUtcERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimezone.h:127
// index:0
// Public
// int standardTimeOffset(const QDateTime &)
func (this *QTimeZone) StandardTimeOffset(atDateTime *QDateTime) int {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimezone.h:128
// index:0
// Public
// int daylightTimeOffset(const QDateTime &)
func (this *QTimeZone) DaylightTimeOffset(atDateTime *QDateTime) int {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimezone.h:130
// index:0
// Public
// bool hasDaylightTime()
func (this *QTimeZone) HasDaylightTime() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone15hasDaylightTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:131
// index:0
// Public
// bool isDaylightTime(const QDateTime &)
func (this *QTimeZone) IsDaylightTime(atDateTime *QDateTime) bool {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14isDaylightTimeERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:133
// index:0
// Public
// QTimeZone::OffsetData offsetData(const QDateTime &)
func (this *QTimeZone) OffsetData(forDateTime *QDateTime) unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = forDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone10offsetDataERK9QDateTime", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:135
// index:0
// Public
// bool hasTransitions()
func (this *QTimeZone) HasTransitions() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14hasTransitionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimezone.h:136
// index:0
// Public
// QTimeZone::OffsetData nextTransition(const QDateTime &)
func (this *QTimeZone) NextTransition(afterDateTime *QDateTime) unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = afterDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14nextTransitionERK9QDateTime", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:137
// index:0
// Public
// QTimeZone::OffsetData previousTransition(const QDateTime &)
func (this *QTimeZone) PreviousTransition(beforeDateTime *QDateTime) unsafe.Pointer /*444*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = beforeDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18previousTransitionERK9QDateTime", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qtimezone.h:140
// index:0
// Public static
// QByteArray systemTimeZoneId()
func (this *QTimeZone) SystemTimeZoneId() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone16systemTimeZoneIdEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTimeZone_SystemTimeZoneId() *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.SystemTimeZoneId()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:141
// index:0
// Public static
// QTimeZone systemTimeZone()
func (this *QTimeZone) SystemTimeZone() *QTimeZone /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone14systemTimeZoneEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTimeZone_SystemTimeZone() *QTimeZone /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.SystemTimeZone()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:142
// index:0
// Public static
// QTimeZone utc()
func (this *QTimeZone) Utc() *QTimeZone /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone3utcEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTimeZoneFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTimeZone_Utc() *QTimeZone /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.Utc()
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:144
// index:0
// Public static
// bool isTimeZoneIdAvailable(const QByteArray &)
func (this *QTimeZone) IsTimeZoneIdAvailable(ianaId *QByteArray) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray", ffiqt.FFI_TYPE_POINTER, ianaId)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QTimeZone_IsTimeZoneIdAvailable(ianaId *QByteArray) bool {
	var nilthis *QTimeZone
	rv := nilthis.IsTimeZoneIdAvailable(ianaId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:150
// index:0
// Public static
// QByteArray ianaIdToWindowsId(const QByteArray &)
func (this *QTimeZone) IanaIdToWindowsId(ianaId *QByteArray) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray", ffiqt.FFI_TYPE_POINTER, ianaId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTimeZone_IanaIdToWindowsId(ianaId *QByteArray) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.IanaIdToWindowsId(ianaId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:151
// index:0
// Public static
// QByteArray windowsIdToDefaultIanaId(const QByteArray &)
func (this *QTimeZone) WindowsIdToDefaultIanaId(windowsId *QByteArray) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray", ffiqt.FFI_TYPE_POINTER, windowsId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTimeZone_WindowsIdToDefaultIanaId(windowsId *QByteArray) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.WindowsIdToDefaultIanaId(windowsId)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:152
// index:1
// Public static
// QByteArray windowsIdToDefaultIanaId(const QByteArray &, QLocale::Country)
func (this *QTimeZone) WindowsIdToDefaultIanaId_1(windowsId *QByteArray, country int) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArrayN7QLocale7CountryE", ffiqt.FFI_TYPE_POINTER, windowsId, country)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTimeZone_WindowsIdToDefaultIanaId_1(windowsId *QByteArray, country int) *QByteArray /*123*/ {
	var nilthis *QTimeZone
	rv := nilthis.WindowsIdToDefaultIanaId_1(windowsId, country)
	return rv
}

type QTimeZone__ = int

const QTimeZone__MinUtcOffsetSecs QTimeZone__ = 4294916896
const QTimeZone__MaxUtcOffsetSecs QTimeZone__ = 50400

type QTimeZone__TimeType = int

const QTimeZone__StandardTime QTimeZone__TimeType = 0
const QTimeZone__DaylightTime QTimeZone__TimeType = 1
const QTimeZone__GenericTime QTimeZone__TimeType = 2

type QTimeZone__NameType = int

const QTimeZone__DefaultName QTimeZone__NameType = 0
const QTimeZone__LongName QTimeZone__NameType = 1
const QTimeZone__ShortName QTimeZone__NameType = 2
const QTimeZone__OffsetName QTimeZone__NameType = 3

//  body block end
