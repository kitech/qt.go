//  header block begin
// /usr/include/qt/QtCore/qlocale.h
// #include <qlocale.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QLocale struct {
	*qtrt.CObject
}

func (this *QLocale) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qlocale.h:929
// index:0
// void QLocale()
func NewQLocale() *QLocale {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}
func NewQLocaleFromPointer(cthis unsafe.Pointer) *QLocale {
	return &QLocale{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qlocale.h:930
// index:1
// void QLocale(const class QString &)
func NewQLocale_1(name unsafe.Pointer) *QLocale {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:931
// index:2
// void QLocale(enum QLocale::Language, enum QLocale::Country)
func NewQLocale_2(language int, country int) *QLocale {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_7CountryE", ffiqt.FFI_TYPE_VOID, cthis, &language, &country)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:932
// index:3
// void QLocale(enum QLocale::Language, enum QLocale::Script, enum QLocale::Country)
func NewQLocale_3(language int, script int, country int) *QLocale {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_6ScriptENS_7CountryE", ffiqt.FFI_TYPE_VOID, cthis, &language, &script, &country)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:938
// index:0
// void ~QLocale()
func DeleteQLocale(*QLocale) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:940
// index:0
// inline
// void swap(class QLocale &)
func (this *QLocale) Swap(other unsafe.Pointer) {
	// 0: (, other QLocale &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:942
// index:0
// QLocale::Language language()
func (this *QLocale) Language() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8languageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:943
// index:0
// QLocale::Script script()
func (this *QLocale) Script() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6scriptEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:944
// index:0
// QLocale::Country country()
func (this *QLocale) Country() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7countryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:945
// index:0
// QString name()
func (this *QLocale) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:947
// index:0
// QString bcp47Name()
func (this *QLocale) Bcp47Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9bcp47NameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:948
// index:0
// QString nativeLanguageName()
func (this *QLocale) NativeLanguageName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale18nativeLanguageNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:949
// index:0
// QString nativeCountryName()
func (this *QLocale) NativeCountryName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17nativeCountryNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:952
// index:0
// short toShort(const class QString &, _Bool *)
func (this *QLocale) ToShort(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:961
// index:1
// short toShort(const class QStringRef &, _Bool *)
func (this *QLocale) ToShort_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:971
// index:2
// short toShort(class QStringView, _Bool *)
func (this *QLocale) ToShort_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:953
// index:0
// ushort toUShort(const class QString &, _Bool *)
func (this *QLocale) ToUShort(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:962
// index:1
// ushort toUShort(const class QStringRef &, _Bool *)
func (this *QLocale) ToUShort_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:972
// index:2
// ushort toUShort(class QStringView, _Bool *)
func (this *QLocale) ToUShort_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:954
// index:0
// int toInt(const class QString &, _Bool *)
func (this *QLocale) ToInt(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:963
// index:1
// int toInt(const class QStringRef &, _Bool *)
func (this *QLocale) ToInt_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:973
// index:2
// int toInt(class QStringView, _Bool *)
func (this *QLocale) ToInt_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:955
// index:0
// uint toUInt(const class QString &, _Bool *)
func (this *QLocale) ToUInt(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:964
// index:1
// uint toUInt(const class QStringRef &, _Bool *)
func (this *QLocale) ToUInt_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:974
// index:2
// uint toUInt(class QStringView, _Bool *)
func (this *QLocale) ToUInt_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:956
// index:0
// qlonglong toLongLong(const class QString &, _Bool *)
func (this *QLocale) ToLongLong(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:965
// index:1
// qlonglong toLongLong(const class QStringRef &, _Bool *)
func (this *QLocale) ToLongLong_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:975
// index:2
// qlonglong toLongLong(class QStringView, _Bool *)
func (this *QLocale) ToLongLong_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:957
// index:0
// qulonglong toULongLong(const class QString &, _Bool *)
func (this *QLocale) ToULongLong(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:966
// index:1
// qulonglong toULongLong(const class QStringRef &, _Bool *)
func (this *QLocale) ToULongLong_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:976
// index:2
// qulonglong toULongLong(class QStringView, _Bool *)
func (this *QLocale) ToULongLong_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:958
// index:0
// float toFloat(const class QString &, _Bool *)
func (this *QLocale) ToFloat(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:967
// index:1
// float toFloat(const class QStringRef &, _Bool *)
func (this *QLocale) ToFloat_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:977
// index:2
// float toFloat(class QStringView, _Bool *)
func (this *QLocale) ToFloat_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:959
// index:0
// double toDouble(const class QString &, _Bool *)
func (this *QLocale) ToDouble(s unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, s const QString &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK7QStringPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:968
// index:1
// double toDouble(const class QStringRef &, _Bool *)
func (this *QLocale) ToDouble_1(s unsafe.Pointer, ok unsafe.Pointer) {
	// 1: (, s const QStringRef &, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK10QStringRefPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:978
// index:2
// double toDouble(class QStringView, _Bool *)
func (this *QLocale) ToDouble_2(s unsafe.Pointer, ok unsafe.Pointer) {
	// 2: (, s QStringView, ok bool *), (s, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleE11QStringViewPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:980
// index:0
// QString toString(qlonglong)
func (this *QLocale) ToString(i int64) {
	// 0: (, i qlonglong), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:981
// index:1
// QString toString(qulonglong)
func (this *QLocale) ToString_1(i uint64) {
	// 1: (, i qulonglong), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEy", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:982
// index:2
// inline
// QString toString(short)
func (this *QLocale) ToString_2(i int16) {
	// 2: (, i short), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEs", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:983
// index:3
// inline
// QString toString(ushort)
func (this *QLocale) ToString_3(i uint16) {
	// 3: (, i ushort), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEt", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:984
// index:4
// inline
// QString toString(int)
func (this *QLocale) ToString_4(i int) {
	// 4: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:985
// index:5
// inline
// QString toString(uint)
func (this *QLocale) ToString_5(i uint) {
	// 5: (, i uint), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:986
// index:6
// QString toString(double, char, int)
func (this *QLocale) ToString_6(i float64, f byte, prec int) {
	// 6: (, i double, f char, prec int), (&i, &f, &prec)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEdci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, &f, &prec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:987
// index:7
// inline
// QString toString(float, char, int)
func (this *QLocale) ToString_7(i float32, f byte, prec int) {
	// 7: (, i float, f char, prec int), (&i, &f, &prec)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEfci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, &f, &prec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:990
// index:8
// QString toString(const class QDate &, const class QString &)
func (this *QLocale) ToString_8(date unsafe.Pointer, formatStr unsafe.Pointer) {
	// 8: (, date const QDate &, formatStr const QString &), (date, formatStr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date, formatStr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:991
// index:9
// QString toString(const class QTime &, const class QString &)
func (this *QLocale) ToString_9(time unsafe.Pointer, formatStr unsafe.Pointer) {
	// 9: (, time const QTime &, formatStr const QString &), (time, formatStr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), time, formatStr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:992
// index:10
// QString toString(const class QDateTime &, const class QString &)
func (this *QLocale) ToString_10(dateTime unsafe.Pointer, format unsafe.Pointer) {
	// 10: (, dateTime const QDateTime &, format const QString &), (dateTime, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dateTime, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:994
// index:11
// QString toString(const class QDate &, class QStringView)
func (this *QLocale) ToString_11(date unsafe.Pointer, formatStr unsafe.Pointer) {
	// 11: (, date const QDate &, formatStr QStringView), (date, formatStr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDate11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date, formatStr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:995
// index:12
// QString toString(const class QTime &, class QStringView)
func (this *QLocale) ToString_12(time unsafe.Pointer, formatStr unsafe.Pointer) {
	// 12: (, time const QTime &, formatStr QStringView), (time, formatStr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTime11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), time, formatStr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:996
// index:13
// QString toString(const class QDateTime &, class QStringView)
func (this *QLocale) ToString_13(dateTime unsafe.Pointer, format unsafe.Pointer) {
	// 13: (, dateTime const QDateTime &, format QStringView), (dateTime, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTime11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dateTime, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:997
// index:14
// QString toString(const class QDate &, enum QLocale::FormatType)
func (this *QLocale) ToString_14(date unsafe.Pointer, format int) {
	// 14: (, date const QDate &, format QLocale::FormatType), (date, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:998
// index:15
// QString toString(const class QTime &, enum QLocale::FormatType)
func (this *QLocale) ToString_15(time unsafe.Pointer, format int) {
	// 15: (, time const QTime &, format QLocale::FormatType), (time, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), time, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:999
// index:16
// QString toString(const class QDateTime &, enum QLocale::FormatType)
func (this *QLocale) ToString_16(dateTime unsafe.Pointer, format int) {
	// 16: (, dateTime const QDateTime &, format QLocale::FormatType), (dateTime, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dateTime, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1001
// index:0
// QString dateFormat(enum QLocale::FormatType)
func (this *QLocale) DateFormat(format int) {
	// 0: (, format QLocale::FormatType), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10dateFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1002
// index:0
// QString timeFormat(enum QLocale::FormatType)
func (this *QLocale) TimeFormat(format int) {
	// 0: (, format QLocale::FormatType), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10timeFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1003
// index:0
// QString dateTimeFormat(enum QLocale::FormatType)
func (this *QLocale) DateTimeFormat(format int) {
	// 0: (, format QLocale::FormatType), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14dateTimeFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1005
// index:0
// QDate toDate(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToDate(string unsafe.Pointer, arg1 int) {
	// 0: (, string const QString &, QLocale::FormatType arg1), (string, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1008
// index:1
// QDate toDate(const class QString &, const class QString &)
func (this *QLocale) ToDate_1(string unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, string const QString &, format const QString &), (string, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1006
// index:0
// QTime toTime(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToTime(string unsafe.Pointer, arg1 int) {
	// 0: (, string const QString &, QLocale::FormatType arg1), (string, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1009
// index:1
// QTime toTime(const class QString &, const class QString &)
func (this *QLocale) ToTime_1(string unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, string const QString &, format const QString &), (string, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1007
// index:0
// QDateTime toDateTime(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToDateTime(string unsafe.Pointer, format int) {
	// 0: (, string const QString &, format QLocale::FormatType), (string, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1010
// index:1
// QDateTime toDateTime(const class QString &, const class QString &)
func (this *QLocale) ToDateTime_1(string unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, string const QString &, format const QString &), (string, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1015
// index:0
// QChar decimalPoint()
func (this *QLocale) DecimalPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12decimalPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1016
// index:0
// QChar groupSeparator()
func (this *QLocale) GroupSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14groupSeparatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1017
// index:0
// QChar percent()
func (this *QLocale) Percent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7percentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1018
// index:0
// QChar zeroDigit()
func (this *QLocale) ZeroDigit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9zeroDigitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1019
// index:0
// QChar negativeSign()
func (this *QLocale) NegativeSign() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12negativeSignEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1020
// index:0
// QChar positiveSign()
func (this *QLocale) PositiveSign() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12positiveSignEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1021
// index:0
// QChar exponential()
func (this *QLocale) Exponential() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11exponentialEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1023
// index:0
// QString monthName(int, enum QLocale::FormatType)
func (this *QLocale) MonthName(arg0 int, format int) {
	// 0: (, int arg0, format QLocale::FormatType), (&arg0, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9monthNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1024
// index:0
// QString standaloneMonthName(int, enum QLocale::FormatType)
func (this *QLocale) StandaloneMonthName(arg0 int, format int) {
	// 0: (, int arg0, format QLocale::FormatType), (&arg0, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale19standaloneMonthNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1025
// index:0
// QString dayName(int, enum QLocale::FormatType)
func (this *QLocale) DayName(arg0 int, format int) {
	// 0: (, int arg0, format QLocale::FormatType), (&arg0, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7dayNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1026
// index:0
// QString standaloneDayName(int, enum QLocale::FormatType)
func (this *QLocale) StandaloneDayName(arg0 int, format int) {
	// 0: (, int arg0, format QLocale::FormatType), (&arg0, &format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17standaloneDayNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1028
// index:0
// Qt::DayOfWeek firstDayOfWeek()
func (this *QLocale) FirstDayOfWeek() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14firstDayOfWeekEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1029
// index:0
// QList<Qt::DayOfWeek> weekdays()
func (this *QLocale) Weekdays() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8weekdaysEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1031
// index:0
// QString amText()
func (this *QLocale) AmText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6amTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1032
// index:0
// QString pmText()
func (this *QLocale) PmText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6pmTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1034
// index:0
// QLocale::MeasurementSystem measurementSystem()
func (this *QLocale) MeasurementSystem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17measurementSystemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1036
// index:0
// Qt::LayoutDirection textDirection()
func (this *QLocale) TextDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale13textDirectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1038
// index:0
// QString toUpper(const class QString &)
func (this *QLocale) ToUpper(str unsafe.Pointer) {
	// 0: (, str const QString &), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toUpperERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1039
// index:0
// QString toLower(const class QString &)
func (this *QLocale) ToLower(str unsafe.Pointer) {
	// 0: (, str const QString &), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toLowerERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1041
// index:0
// QString currencySymbol(enum QLocale::CurrencySymbolFormat)
func (this *QLocale) CurrencySymbol(arg0 int) {
	// 0: (, QLocale::CurrencySymbolFormat arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14currencySymbolENS_20CurrencySymbolFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1042
// index:0
// QString toCurrencyString(qlonglong, const class QString &)
func (this *QLocale) ToCurrencyString(arg0 int64, symbol unsafe.Pointer) {
	// 0: (, qlonglong arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringExRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1043
// index:1
// QString toCurrencyString(qulonglong, const class QString &)
func (this *QLocale) ToCurrencyString_1(arg0 uint64, symbol unsafe.Pointer) {
	// 1: (, qulonglong arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEyRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1044
// index:2
// inline
// QString toCurrencyString(short, const class QString &)
func (this *QLocale) ToCurrencyString_2(arg0 int16, symbol unsafe.Pointer) {
	// 2: (, short arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEsRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1045
// index:3
// inline
// QString toCurrencyString(ushort, const class QString &)
func (this *QLocale) ToCurrencyString_3(arg0 uint16, symbol unsafe.Pointer) {
	// 3: (, ushort arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEtRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1046
// index:4
// inline
// QString toCurrencyString(int, const class QString &)
func (this *QLocale) ToCurrencyString_4(arg0 int, symbol unsafe.Pointer) {
	// 4: (, int arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1047
// index:5
// inline
// QString toCurrencyString(uint, const class QString &)
func (this *QLocale) ToCurrencyString_5(arg0 uint, symbol unsafe.Pointer) {
	// 5: (, uint arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEjRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1053
// index:6
// QString toCurrencyString(double, const class QString &)
func (this *QLocale) ToCurrencyString_6(arg0 float64, symbol unsafe.Pointer) {
	// 6: (, double arg0, symbol const QString &), (&arg0, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1054
// index:7
// QString toCurrencyString(double, const class QString &, int)
func (this *QLocale) ToCurrencyString_7(arg0 float64, symbol unsafe.Pointer, precision int) {
	// 7: (, double arg0, symbol const QString &, precision int), (&arg0, symbol, &precision)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, symbol, &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1055
// index:8
// inline
// QString toCurrencyString(float, const class QString &)
func (this *QLocale) ToCurrencyString_8(i float32, symbol unsafe.Pointer) {
	// 8: (, i float, symbol const QString &), (&i, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1057
// index:9
// inline
// QString toCurrencyString(float, const class QString &, int)
func (this *QLocale) ToCurrencyString_9(i float32, symbol unsafe.Pointer, precision int) {
	// 9: (, i float, symbol const QString &, precision int), (&i, symbol, &precision)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, symbol, &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1061
// index:0
// QString formattedDataSize(qint64, int, QLocale::DataSizeFormats)
func (this *QLocale) FormattedDataSize(bytes int64, precision int, format int) {
	// 0: (, bytes qint64, precision int, QFlags<QLocale::DataSizeFormat> format), (&bytes, &precision, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale17formattedDataSizeExi6QFlagsINS_14DataSizeFormatEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bytes, &precision, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1063
// index:0
// QStringList uiLanguages()
func (this *QLocale) UiLanguages() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11uiLanguagesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1068
// index:0
// static
// QString languageToString(enum QLocale::Language)
func (this *QLocale) LanguageToString(language int) {
	// 0: (language QLocale::Language), (language)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale16languageToStringENS_8LanguageE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_LanguageToString(language int) {
	// 0: (language QLocale::Language), (language)
	var nilthis *QLocale
	nilthis.LanguageToString(language)
}

// /usr/include/qt/QtCore/qlocale.h:1069
// index:0
// static
// QString countryToString(enum QLocale::Country)
func (this *QLocale) CountryToString(country int) {
	// 0: (country QLocale::Country), (country)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale15countryToStringENS_7CountryE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_CountryToString(country int) {
	// 0: (country QLocale::Country), (country)
	var nilthis *QLocale
	nilthis.CountryToString(country)
}

// /usr/include/qt/QtCore/qlocale.h:1070
// index:0
// static
// QString scriptToString(enum QLocale::Script)
func (this *QLocale) ScriptToString(script int) {
	// 0: (script QLocale::Script), (script)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale14scriptToStringENS_6ScriptE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_ScriptToString(script int) {
	// 0: (script QLocale::Script), (script)
	var nilthis *QLocale
	nilthis.ScriptToString(script)
}

// /usr/include/qt/QtCore/qlocale.h:1071
// index:0
// static
// void setDefault(const class QLocale &)
func (this *QLocale) SetDefault(locale unsafe.Pointer) {
	// 0: (locale const QLocale &), (locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale10setDefaultERKS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_SetDefault(locale unsafe.Pointer) {
	// 0: (locale const QLocale &), (locale)
	var nilthis *QLocale
	nilthis.SetDefault(locale)
}

// /usr/include/qt/QtCore/qlocale.h:1073
// index:0
// static inline
// QLocale c()
func (this *QLocale) C() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale1cEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_C() {
	// 0: (), ()
	var nilthis *QLocale
	nilthis.C()
}

// /usr/include/qt/QtCore/qlocale.h:1074
// index:0
// static
// QLocale system()
func (this *QLocale) System() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale6systemEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_System() {
	// 0: (), ()
	var nilthis *QLocale
	nilthis.System()
}

// /usr/include/qt/QtCore/qlocale.h:1076
// index:0
// static
// QList<QLocale> matchingLocales(class QLocale::Language, class QLocale::Script, class QLocale::Country)
func (this *QLocale) MatchingLocales(language int, script int, country int) {
	// 0: (language QLocale::Language, script QLocale::Script, country QLocale::Country), (language, script, country)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale15matchingLocalesENS_8LanguageENS_6ScriptENS_7CountryE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_MatchingLocales(language int, script int, country int) {
	// 0: (language QLocale::Language, script QLocale::Script, country QLocale::Country), (language, script, country)
	var nilthis *QLocale
	nilthis.MatchingLocales(language, script, country)
}

// /usr/include/qt/QtCore/qlocale.h:1077
// index:0
// static
// QList<QLocale::Country> countriesForLanguage(enum QLocale::Language)
func (this *QLocale) CountriesForLanguage(lang int) {
	// 0: (lang QLocale::Language), (lang)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale20countriesForLanguageENS_8LanguageE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLocale_CountriesForLanguage(lang int) {
	// 0: (lang QLocale::Language), (lang)
	var nilthis *QLocale
	nilthis.CountriesForLanguage(lang)
}

// /usr/include/qt/QtCore/qlocale.h:1079
// index:0
// void setNumberOptions(QLocale::NumberOptions)
func (this *QLocale) SetNumberOptions(options int) {
	// 0: (, QFlags<QLocale::NumberOption> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale16setNumberOptionsE6QFlagsINS_12NumberOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1080
// index:0
// QLocale::NumberOptions numberOptions()
func (this *QLocale) NumberOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale13numberOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1083
// index:0
// QString quoteString(const class QString &, enum QLocale::QuotationStyle)
func (this *QLocale) QuoteString(str unsafe.Pointer, style int) {
	// 0: (, str const QString &, style QLocale::QuotationStyle), (str, &style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK7QStringNS_14QuotationStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1084
// index:1
// QString quoteString(const class QStringRef &, enum QLocale::QuotationStyle)
func (this *QLocale) QuoteString_1(str unsafe.Pointer, style int) {
	// 1: (, str const QStringRef &, style QLocale::QuotationStyle), (str, &style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK10QStringRefNS_14QuotationStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1086
// index:0
// QString createSeparatedList(const class QStringList &)
func (this *QLocale) CreateSeparatedList(strl unsafe.Pointer) {
	// 0: (, strl const QStringList &), (strl)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale19createSeparatedListERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), strl)
	gopp.ErrPrint(err, rv)
}

//  body block end
