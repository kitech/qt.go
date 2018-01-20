//  header block begin
// /usr/include/qt/QtCore/qtimezone.h
// #include <qtimezone.h>
// #include <QtCore>
package qtcore

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
type QTimeZone struct {
	*qtrt.CObject
}

func (this *QTimeZone) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTimeZoneFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return &QTimeZone{&qtrt.CObject{cthis}}
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
// void QTimeZone(const class QByteArray &)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &offsetSeconds)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// Public
// void QTimeZone(const class QByteArray &, int, const class QString &, const class QString &, class QLocale::Country, const class QString &)
func NewQTimeZone_3(zoneId *QByteArray, offsetSeconds int, name *QString, abbreviation *QString, country int, comment *QString) *QTimeZone {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = zoneId.GetCthis()
	var convArg2 = name.GetCthis()
	var convArg3 = abbreviation.GetCthis()
	var convArg5 = comment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &offsetSeconds, convArg2, convArg3, &country, convArg5)
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
// void swap(class QTimeZone &)
func (this *QTimeZone) Swap(other *QTimeZone) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:112
// index:0
// Public
// bool isValid()
func (this *QTimeZone) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:114
// index:0
// Public
// QByteArray id()
func (this *QTimeZone) Id() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone2idEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:115
// index:0
// Public
// QLocale::Country country()
func (this *QTimeZone) Country() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7countryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:116
// index:0
// Public
// QString comment()
func (this *QTimeZone) Comment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7commentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// Public
// QString displayName(const class QDateTime &, class QTimeZone::NameType, const class QLocale &)
func (this *QTimeZone) DisplayName(atDateTime *QDateTime, nameType int, locale *QLocale) interface{} {
	var convArg0 = atDateTime.GetCthis()
	var convArg2 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &nameType, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// Public
// QString displayName(class QTimeZone::TimeType, class QTimeZone::NameType, const class QLocale &)
func (this *QTimeZone) DisplayName_1(timeType int, nameType int, locale *QLocale) interface{} {
	var convArg2 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &timeType, &nameType, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:124
// index:0
// Public
// QString abbreviation(const class QDateTime &)
func (this *QTimeZone) Abbreviation(atDateTime *QDateTime) interface{} {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone12abbreviationERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:126
// index:0
// Public
// int offsetFromUtc(const class QDateTime &)
func (this *QTimeZone) OffsetFromUtc(atDateTime *QDateTime) interface{} {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone13offsetFromUtcERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:127
// index:0
// Public
// int standardTimeOffset(const class QDateTime &)
func (this *QTimeZone) StandardTimeOffset(atDateTime *QDateTime) interface{} {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:128
// index:0
// Public
// int daylightTimeOffset(const class QDateTime &)
func (this *QTimeZone) DaylightTimeOffset(atDateTime *QDateTime) interface{} {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:130
// index:0
// Public
// bool hasDaylightTime()
func (this *QTimeZone) HasDaylightTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone15hasDaylightTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:131
// index:0
// Public
// bool isDaylightTime(const class QDateTime &)
func (this *QTimeZone) IsDaylightTime(atDateTime *QDateTime) interface{} {
	var convArg0 = atDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14isDaylightTimeERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:133
// index:0
// Public
// QTimeZone::OffsetData offsetData(const class QDateTime &)
func (this *QTimeZone) OffsetData(forDateTime *QDateTime) interface{} {
	var convArg0 = forDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone10offsetDataERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:135
// index:0
// Public
// bool hasTransitions()
func (this *QTimeZone) HasTransitions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14hasTransitionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:136
// index:0
// Public
// QTimeZone::OffsetData nextTransition(const class QDateTime &)
func (this *QTimeZone) NextTransition(afterDateTime *QDateTime) interface{} {
	var convArg0 = afterDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14nextTransitionERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:137
// index:0
// Public
// QTimeZone::OffsetData previousTransition(const class QDateTime &)
func (this *QTimeZone) PreviousTransition(beforeDateTime *QDateTime) interface{} {
	var convArg0 = beforeDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18previousTransitionERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:138
// index:0
// Public
// QTimeZone::OffsetDataList transitions(const class QDateTime &, const class QDateTime &)
func (this *QTimeZone) Transitions(fromDateTime *QDateTime, toDateTime *QDateTime) interface{} {
	var convArg0 = fromDateTime.GetCthis()
	var convArg1 = toDateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11transitionsERK9QDateTimeS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtimezone.h:140
// index:0
// Public static
// QByteArray systemTimeZoneId()
func (this *QTimeZone) SystemTimeZoneId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone16systemTimeZoneIdEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_SystemTimeZoneId() {
	var nilthis *QTimeZone
	nilthis.SystemTimeZoneId()
}

// /usr/include/qt/QtCore/qtimezone.h:141
// index:0
// Public static
// QTimeZone systemTimeZone()
func (this *QTimeZone) SystemTimeZone() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone14systemTimeZoneEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_SystemTimeZone() {
	var nilthis *QTimeZone
	nilthis.SystemTimeZone()
}

// /usr/include/qt/QtCore/qtimezone.h:142
// index:0
// Public static
// QTimeZone utc()
func (this *QTimeZone) Utc() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone3utcEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_Utc() {
	var nilthis *QTimeZone
	nilthis.Utc()
}

// /usr/include/qt/QtCore/qtimezone.h:144
// index:0
// Public static
// bool isTimeZoneIdAvailable(const class QByteArray &)
func (this *QTimeZone) IsTimeZoneIdAvailable(ianaId *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray", ffiqt.FFI_TYPE_POINTER, ianaId)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_IsTimeZoneIdAvailable(ianaId *QByteArray) {
	var nilthis *QTimeZone
	nilthis.IsTimeZoneIdAvailable(ianaId)
}

// /usr/include/qt/QtCore/qtimezone.h:146
// index:0
// Public static
// QList<QByteArray> availableTimeZoneIds()
func (this *QTimeZone) AvailableTimeZoneIds() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone20availableTimeZoneIdsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_AvailableTimeZoneIds() {
	var nilthis *QTimeZone
	nilthis.AvailableTimeZoneIds()
}

// /usr/include/qt/QtCore/qtimezone.h:147
// index:1
// Public static
// QList<QByteArray> availableTimeZoneIds(class QLocale::Country)
func (this *QTimeZone) AvailableTimeZoneIds_1(country int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone20availableTimeZoneIdsEN7QLocale7CountryE", ffiqt.FFI_TYPE_POINTER, country)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_AvailableTimeZoneIds_1(country int) {
	var nilthis *QTimeZone
	nilthis.AvailableTimeZoneIds_1(country)
}

// /usr/include/qt/QtCore/qtimezone.h:148
// index:2
// Public static
// QList<QByteArray> availableTimeZoneIds(int)
func (this *QTimeZone) AvailableTimeZoneIds_2(offsetSeconds int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone20availableTimeZoneIdsEi", ffiqt.FFI_TYPE_POINTER, offsetSeconds)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_AvailableTimeZoneIds_2(offsetSeconds int) {
	var nilthis *QTimeZone
	nilthis.AvailableTimeZoneIds_2(offsetSeconds)
}

// /usr/include/qt/QtCore/qtimezone.h:150
// index:0
// Public static
// QByteArray ianaIdToWindowsId(const class QByteArray &)
func (this *QTimeZone) IanaIdToWindowsId(ianaId *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray", ffiqt.FFI_TYPE_POINTER, ianaId)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_IanaIdToWindowsId(ianaId *QByteArray) {
	var nilthis *QTimeZone
	nilthis.IanaIdToWindowsId(ianaId)
}

// /usr/include/qt/QtCore/qtimezone.h:151
// index:0
// Public static
// QByteArray windowsIdToDefaultIanaId(const class QByteArray &)
func (this *QTimeZone) WindowsIdToDefaultIanaId(windowsId *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray", ffiqt.FFI_TYPE_POINTER, windowsId)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_WindowsIdToDefaultIanaId(windowsId *QByteArray) {
	var nilthis *QTimeZone
	nilthis.WindowsIdToDefaultIanaId(windowsId)
}

// /usr/include/qt/QtCore/qtimezone.h:152
// index:1
// Public static
// QByteArray windowsIdToDefaultIanaId(const class QByteArray &, class QLocale::Country)
func (this *QTimeZone) WindowsIdToDefaultIanaId_1(windowsId *QByteArray, country int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArrayN7QLocale7CountryE", ffiqt.FFI_TYPE_POINTER, windowsId, country)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_WindowsIdToDefaultIanaId_1(windowsId *QByteArray, country int) {
	var nilthis *QTimeZone
	nilthis.WindowsIdToDefaultIanaId_1(windowsId, country)
}

// /usr/include/qt/QtCore/qtimezone.h:154
// index:0
// Public static
// QList<QByteArray> windowsIdToIanaIds(const class QByteArray &)
func (this *QTimeZone) WindowsIdToIanaIds(windowsId *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray", ffiqt.FFI_TYPE_POINTER, windowsId)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_WindowsIdToIanaIds(windowsId *QByteArray) {
	var nilthis *QTimeZone
	nilthis.WindowsIdToIanaIds(windowsId)
}

// /usr/include/qt/QtCore/qtimezone.h:155
// index:1
// Public static
// QList<QByteArray> windowsIdToIanaIds(const class QByteArray &, class QLocale::Country)
func (this *QTimeZone) WindowsIdToIanaIds_1(windowsId *QByteArray, country int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArrayN7QLocale7CountryE", ffiqt.FFI_TYPE_POINTER, windowsId, country)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTimeZone_WindowsIdToIanaIds_1(windowsId *QByteArray, country int) {
	var nilthis *QTimeZone
	nilthis.WindowsIdToIanaIds_1(windowsId, country)
}

//  body block end
