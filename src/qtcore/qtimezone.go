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

// /usr/include/qt/QtCore/qtimezone.h:92
// index:0
// void QTimeZone()
func NewQTimeZone() *QTimeZone {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}
func NewQTimeZoneFromPointer(cthis unsafe.Pointer) *QTimeZone {
	return &QTimeZone{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtimezone.h:93
// index:1
// void QTimeZone(const class QByteArray &)
func NewQTimeZone_1(ianaId unsafe.Pointer) *QTimeZone {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, ianaId)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:94
// index:2
// void QTimeZone(int)
func NewQTimeZone_2(offsetSeconds int) *QTimeZone {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &offsetSeconds)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:95
// index:3
// void QTimeZone(const class QByteArray &, int, const class QString &, const class QString &, class QLocale::Country, const class QString &)
func NewQTimeZone_3(zoneId unsafe.Pointer, offsetSeconds int, name unsafe.Pointer, abbreviation unsafe.Pointer, country int, comment unsafe.Pointer) *QTimeZone {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneC2ERK10QByteArrayiRK7QStringS5_N7QLocale7CountryES5_", ffiqt.FFI_TYPE_VOID, cthis, zoneId, &offsetSeconds, name, abbreviation, &country, comment)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeZoneFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimezone.h:99
// index:0
// void ~QTimeZone()
func DeleteQTimeZone(*QTimeZone) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZoneD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:106
// index:0
// inline
// void swap(class QTimeZone &)
func (this *QTimeZone) Swap(other unsafe.Pointer) {
	// 0: (, other QTimeZone &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:112
// index:0
// bool isValid()
func (this *QTimeZone) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:114
// index:0
// QByteArray id()
func (this *QTimeZone) Id() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone2idEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:115
// index:0
// QLocale::Country country()
func (this *QTimeZone) Country() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7countryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:116
// index:0
// QString comment()
func (this *QTimeZone) Comment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone7commentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:118
// index:0
// QString displayName(const class QDateTime &, class QTimeZone::NameType, const class QLocale &)
func (this *QTimeZone) DisplayName(atDateTime unsafe.Pointer, nameType int, locale unsafe.Pointer) {
	// 0: (, atDateTime const QDateTime &, nameType QTimeZone::NameType, locale const QLocale &), (atDateTime, &nameType, locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameERK9QDateTimeNS_8NameTypeERK7QLocale", ffiqt.FFI_TYPE_VOID, this.GetCthis(), atDateTime, &nameType, locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:121
// index:1
// QString displayName(class QTimeZone::TimeType, class QTimeZone::NameType, const class QLocale &)
func (this *QTimeZone) DisplayName_1(timeType int, nameType int, locale unsafe.Pointer) {
	// 1: (, timeType QTimeZone::TimeType, nameType QTimeZone::NameType, locale const QLocale &), (&timeType, &nameType, locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11displayNameENS_8TimeTypeENS_8NameTypeERK7QLocale", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &timeType, &nameType, locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:124
// index:0
// QString abbreviation(const class QDateTime &)
func (this *QTimeZone) Abbreviation(atDateTime unsafe.Pointer) {
	// 0: (, atDateTime const QDateTime &), (atDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone12abbreviationERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), atDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:126
// index:0
// int offsetFromUtc(const class QDateTime &)
func (this *QTimeZone) OffsetFromUtc(atDateTime unsafe.Pointer) {
	// 0: (, atDateTime const QDateTime &), (atDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone13offsetFromUtcERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), atDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:127
// index:0
// int standardTimeOffset(const class QDateTime &)
func (this *QTimeZone) StandardTimeOffset(atDateTime unsafe.Pointer) {
	// 0: (, atDateTime const QDateTime &), (atDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), atDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:128
// index:0
// int daylightTimeOffset(const class QDateTime &)
func (this *QTimeZone) DaylightTimeOffset(atDateTime unsafe.Pointer) {
	// 0: (, atDateTime const QDateTime &), (atDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), atDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:130
// index:0
// bool hasDaylightTime()
func (this *QTimeZone) HasDaylightTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone15hasDaylightTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:131
// index:0
// bool isDaylightTime(const class QDateTime &)
func (this *QTimeZone) IsDaylightTime(atDateTime unsafe.Pointer) {
	// 0: (, atDateTime const QDateTime &), (atDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14isDaylightTimeERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), atDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:133
// index:0
// QTimeZone::OffsetData offsetData(const class QDateTime &)
func (this *QTimeZone) OffsetData(forDateTime unsafe.Pointer) {
	// 0: (, forDateTime const QDateTime &), (forDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone10offsetDataERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), forDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:135
// index:0
// bool hasTransitions()
func (this *QTimeZone) HasTransitions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14hasTransitionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:136
// index:0
// QTimeZone::OffsetData nextTransition(const class QDateTime &)
func (this *QTimeZone) NextTransition(afterDateTime unsafe.Pointer) {
	// 0: (, afterDateTime const QDateTime &), (afterDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone14nextTransitionERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), afterDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:137
// index:0
// QTimeZone::OffsetData previousTransition(const class QDateTime &)
func (this *QTimeZone) PreviousTransition(beforeDateTime unsafe.Pointer) {
	// 0: (, beforeDateTime const QDateTime &), (beforeDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone18previousTransitionERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.GetCthis(), beforeDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:138
// index:0
// QTimeZone::OffsetDataList transitions(const class QDateTime &, const class QDateTime &)
func (this *QTimeZone) Transitions(fromDateTime unsafe.Pointer, toDateTime unsafe.Pointer) {
	// 0: (, fromDateTime const QDateTime &, toDateTime const QDateTime &), (fromDateTime, toDateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeZone11transitionsERK9QDateTimeS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fromDateTime, toDateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimezone.h:140
// index:0
// static
// QByteArray systemTimeZoneId()
func (this *QTimeZone) SystemTimeZoneId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone16systemTimeZoneIdEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_SystemTimeZoneId() {
	// 0: (), ()
	var nilthis *QTimeZone
	nilthis.SystemTimeZoneId()
}

// /usr/include/qt/QtCore/qtimezone.h:141
// index:0
// static
// QTimeZone systemTimeZone()
func (this *QTimeZone) SystemTimeZone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone14systemTimeZoneEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_SystemTimeZone() {
	// 0: (), ()
	var nilthis *QTimeZone
	nilthis.SystemTimeZone()
}

// /usr/include/qt/QtCore/qtimezone.h:142
// index:0
// static
// QTimeZone utc()
func (this *QTimeZone) Utc() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone3utcEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_Utc() {
	// 0: (), ()
	var nilthis *QTimeZone
	nilthis.Utc()
}

// /usr/include/qt/QtCore/qtimezone.h:144
// index:0
// static
// bool isTimeZoneIdAvailable(const class QByteArray &)
func (this *QTimeZone) IsTimeZoneIdAvailable(ianaId unsafe.Pointer) {
	// 0: (ianaId const QByteArray &), (ianaId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_IsTimeZoneIdAvailable(ianaId unsafe.Pointer) {
	// 0: (ianaId const QByteArray &), (ianaId)
	var nilthis *QTimeZone
	nilthis.IsTimeZoneIdAvailable(ianaId)
}

// /usr/include/qt/QtCore/qtimezone.h:146
// index:0
// static
// QList<QByteArray> availableTimeZoneIds()
func (this *QTimeZone) AvailableTimeZoneIds() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone20availableTimeZoneIdsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_AvailableTimeZoneIds() {
	// 0: (), ()
	var nilthis *QTimeZone
	nilthis.AvailableTimeZoneIds()
}

// /usr/include/qt/QtCore/qtimezone.h:147
// index:1
// static
// QList<QByteArray> availableTimeZoneIds(class QLocale::Country)
func (this *QTimeZone) AvailableTimeZoneIds_1(country int) {
	// 1: (country QLocale::Country), (country)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone20availableTimeZoneIdsEN7QLocale7CountryE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_AvailableTimeZoneIds_1(country int) {
	// 1: (country QLocale::Country), (country)
	var nilthis *QTimeZone
	nilthis.AvailableTimeZoneIds_1(country)
}

// /usr/include/qt/QtCore/qtimezone.h:148
// index:2
// static
// QList<QByteArray> availableTimeZoneIds(int)
func (this *QTimeZone) AvailableTimeZoneIds_2(offsetSeconds int) {
	// 2: (offsetSeconds int), (offsetSeconds)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone20availableTimeZoneIdsEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_AvailableTimeZoneIds_2(offsetSeconds int) {
	// 2: (offsetSeconds int), (offsetSeconds)
	var nilthis *QTimeZone
	nilthis.AvailableTimeZoneIds_2(offsetSeconds)
}

// /usr/include/qt/QtCore/qtimezone.h:150
// index:0
// static
// QByteArray ianaIdToWindowsId(const class QByteArray &)
func (this *QTimeZone) IanaIdToWindowsId(ianaId unsafe.Pointer) {
	// 0: (ianaId const QByteArray &), (ianaId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_IanaIdToWindowsId(ianaId unsafe.Pointer) {
	// 0: (ianaId const QByteArray &), (ianaId)
	var nilthis *QTimeZone
	nilthis.IanaIdToWindowsId(ianaId)
}

// /usr/include/qt/QtCore/qtimezone.h:151
// index:0
// static
// QByteArray windowsIdToDefaultIanaId(const class QByteArray &)
func (this *QTimeZone) WindowsIdToDefaultIanaId(windowsId unsafe.Pointer) {
	// 0: (windowsId const QByteArray &), (windowsId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_WindowsIdToDefaultIanaId(windowsId unsafe.Pointer) {
	// 0: (windowsId const QByteArray &), (windowsId)
	var nilthis *QTimeZone
	nilthis.WindowsIdToDefaultIanaId(windowsId)
}

// /usr/include/qt/QtCore/qtimezone.h:152
// index:1
// static
// QByteArray windowsIdToDefaultIanaId(const class QByteArray &, class QLocale::Country)
func (this *QTimeZone) WindowsIdToDefaultIanaId_1(windowsId unsafe.Pointer, country int) {
	// 1: (windowsId const QByteArray &, country QLocale::Country), (windowsId, country)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArrayN7QLocale7CountryE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_WindowsIdToDefaultIanaId_1(windowsId unsafe.Pointer, country int) {
	// 1: (windowsId const QByteArray &, country QLocale::Country), (windowsId, country)
	var nilthis *QTimeZone
	nilthis.WindowsIdToDefaultIanaId_1(windowsId, country)
}

// /usr/include/qt/QtCore/qtimezone.h:154
// index:0
// static
// QList<QByteArray> windowsIdToIanaIds(const class QByteArray &)
func (this *QTimeZone) WindowsIdToIanaIds(windowsId unsafe.Pointer) {
	// 0: (windowsId const QByteArray &), (windowsId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_WindowsIdToIanaIds(windowsId unsafe.Pointer) {
	// 0: (windowsId const QByteArray &), (windowsId)
	var nilthis *QTimeZone
	nilthis.WindowsIdToIanaIds(windowsId)
}

// /usr/include/qt/QtCore/qtimezone.h:155
// index:1
// static
// QList<QByteArray> windowsIdToIanaIds(const class QByteArray &, class QLocale::Country)
func (this *QTimeZone) WindowsIdToIanaIds_1(windowsId unsafe.Pointer, country int) {
	// 1: (windowsId const QByteArray &, country QLocale::Country), (windowsId, country)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArrayN7QLocale7CountryE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimeZone_WindowsIdToIanaIds_1(windowsId unsafe.Pointer, country int) {
	// 1: (windowsId const QByteArray &, country QLocale::Country), (windowsId, country)
	var nilthis *QTimeZone
	nilthis.WindowsIdToIanaIds_1(windowsId, country)
}

//  body block end
