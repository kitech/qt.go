package qtcore

// /usr/include/qt/QtCore/qlocale.h
// #include <qlocale.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQLocaleFromPointer(cthis unsafe.Pointer) *QLocale {
	return &QLocale{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qlocale.h:929
// index:0
// Public
// void QLocale()
func NewQLocale() *QLocale {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:930
// index:1
// Public
// void QLocale(const class QString &)
func NewQLocale_1(name *QString) *QLocale {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:931
// index:2
// Public
// void QLocale(enum QLocale::Language, enum QLocale::Country)
func NewQLocale_2(language int, country int) *QLocale {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_7CountryE", ffiqt.FFI_TYPE_VOID, cthis, &language, &country)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:932
// index:3
// Public
// void QLocale(enum QLocale::Language, enum QLocale::Script, enum QLocale::Country)
func NewQLocale_3(language int, script int, country int) *QLocale {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_6ScriptENS_7CountryE", ffiqt.FFI_TYPE_VOID, cthis, &language, &script, &country)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:938
// index:0
// Public
// void ~QLocale()
func DeleteQLocale(*QLocale) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:940
// index:0
// Public inline
// void swap(class QLocale &)
func (this *QLocale) Swap(other *QLocale) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:942
// index:0
// Public
// QLocale::Language language()
func (this *QLocale) Language() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8languageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:943
// index:0
// Public
// QLocale::Script script()
func (this *QLocale) Script() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6scriptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:944
// index:0
// Public
// QLocale::Country country()
func (this *QLocale) Country() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7countryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:945
// index:0
// Public
// QString name()
func (this *QLocale) Name() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:947
// index:0
// Public
// QString bcp47Name()
func (this *QLocale) Bcp47Name() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9bcp47NameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:948
// index:0
// Public
// QString nativeLanguageName()
func (this *QLocale) NativeLanguageName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale18nativeLanguageNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:949
// index:0
// Public
// QString nativeCountryName()
func (this *QLocale) NativeCountryName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17nativeCountryNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:952
// index:0
// Public
// short toShort(const class QString &, _Bool *)
func (this *QLocale) ToShort(s *QString, ok unsafe.Pointer /*666*/) int16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int16(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:961
// index:1
// Public
// short toShort(const class QStringRef &, _Bool *)
func (this *QLocale) ToShort_1(s *QStringRef, ok unsafe.Pointer /*666*/) int16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int16(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:971
// index:2
// Public
// short toShort(class QStringView, _Bool *)
func (this *QLocale) ToShort_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) int16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int16(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:953
// index:0
// Public
// ushort toUShort(const class QString &, _Bool *)
func (this *QLocale) ToUShort(s *QString, ok unsafe.Pointer /*666*/) uint16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:962
// index:1
// Public
// ushort toUShort(const class QStringRef &, _Bool *)
func (this *QLocale) ToUShort_1(s *QStringRef, ok unsafe.Pointer /*666*/) uint16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:972
// index:2
// Public
// ushort toUShort(class QStringView, _Bool *)
func (this *QLocale) ToUShort_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) uint16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:954
// index:0
// Public
// int toInt(const class QString &, _Bool *)
func (this *QLocale) ToInt(s *QString, ok unsafe.Pointer /*666*/) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:963
// index:1
// Public
// int toInt(const class QStringRef &, _Bool *)
func (this *QLocale) ToInt_1(s *QStringRef, ok unsafe.Pointer /*666*/) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:973
// index:2
// Public
// int toInt(class QStringView, _Bool *)
func (this *QLocale) ToInt_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:955
// index:0
// Public
// uint toUInt(const class QString &, _Bool *)
func (this *QLocale) ToUInt(s *QString, ok unsafe.Pointer /*666*/) uint {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:964
// index:1
// Public
// uint toUInt(const class QStringRef &, _Bool *)
func (this *QLocale) ToUInt_1(s *QStringRef, ok unsafe.Pointer /*666*/) uint {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:974
// index:2
// Public
// uint toUInt(class QStringView, _Bool *)
func (this *QLocale) ToUInt_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) uint {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:956
// index:0
// Public
// qlonglong toLongLong(const class QString &, _Bool *)
func (this *QLocale) ToLongLong(s *QString, ok unsafe.Pointer /*666*/) int64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:965
// index:1
// Public
// qlonglong toLongLong(const class QStringRef &, _Bool *)
func (this *QLocale) ToLongLong_1(s *QStringRef, ok unsafe.Pointer /*666*/) int64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:975
// index:2
// Public
// qlonglong toLongLong(class QStringView, _Bool *)
func (this *QLocale) ToLongLong_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) int64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:957
// index:0
// Public
// qulonglong toULongLong(const class QString &, _Bool *)
func (this *QLocale) ToULongLong(s *QString, ok unsafe.Pointer /*666*/) uint64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:966
// index:1
// Public
// qulonglong toULongLong(const class QStringRef &, _Bool *)
func (this *QLocale) ToULongLong_1(s *QStringRef, ok unsafe.Pointer /*666*/) uint64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:976
// index:2
// Public
// qulonglong toULongLong(class QStringView, _Bool *)
func (this *QLocale) ToULongLong_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) uint64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:958
// index:0
// Public
// float toFloat(const class QString &, _Bool *)
func (this *QLocale) ToFloat(s *QString, ok unsafe.Pointer /*666*/) float32 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:967
// index:1
// Public
// float toFloat(const class QStringRef &, _Bool *)
func (this *QLocale) ToFloat_1(s *QStringRef, ok unsafe.Pointer /*666*/) float32 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:977
// index:2
// Public
// float toFloat(class QStringView, _Bool *)
func (this *QLocale) ToFloat_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) float32 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:959
// index:0
// Public
// double toDouble(const class QString &, _Bool *)
func (this *QLocale) ToDouble(s *QString, ok unsafe.Pointer /*666*/) float64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:968
// index:1
// Public
// double toDouble(const class QStringRef &, _Bool *)
func (this *QLocale) ToDouble_1(s *QStringRef, ok unsafe.Pointer /*666*/) float64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:978
// index:2
// Public
// double toDouble(class QStringView, _Bool *)
func (this *QLocale) ToDouble_2(s *QStringView /*123*/, ok unsafe.Pointer /*666*/) float64 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:980
// index:0
// Public
// QString toString(qlonglong)
func (this *QLocale) ToString(i int64) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:981
// index:1
// Public
// QString toString(qulonglong)
func (this *QLocale) ToString_1(i uint64) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:982
// index:2
// Public inline
// QString toString(short)
func (this *QLocale) ToString_2(i int16) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEs", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:983
// index:3
// Public inline
// QString toString(ushort)
func (this *QLocale) ToString_3(i uint16) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:984
// index:4
// Public inline
// QString toString(int)
func (this *QLocale) ToString_4(i int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:985
// index:5
// Public inline
// QString toString(uint)
func (this *QLocale) ToString_5(i uint) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:986
// index:6
// Public
// QString toString(double, char, int)
func (this *QLocale) ToString_6(i float64, f byte, prec int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEdci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &f, &prec)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:987
// index:7
// Public inline
// QString toString(float, char, int)
func (this *QLocale) ToString_7(i float32, f byte, prec int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEfci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &f, &prec)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:990
// index:8
// Public
// QString toString(const class QDate &, const class QString &)
func (this *QLocale) ToString_8(date *QDate, formatStr *QString) *QString /*123*/ {
	var convArg0 = date.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:991
// index:9
// Public
// QString toString(const class QTime &, const class QString &)
func (this *QLocale) ToString_9(time *QTime, formatStr *QString) *QString /*123*/ {
	var convArg0 = time.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:992
// index:10
// Public
// QString toString(const class QDateTime &, const class QString &)
func (this *QLocale) ToString_10(dateTime *QDateTime, format *QString) *QString /*123*/ {
	var convArg0 = dateTime.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:994
// index:11
// Public
// QString toString(const class QDate &, class QStringView)
func (this *QLocale) ToString_11(date *QDate, formatStr *QStringView /*123*/) *QString /*123*/ {
	var convArg0 = date.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDate11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:995
// index:12
// Public
// QString toString(const class QTime &, class QStringView)
func (this *QLocale) ToString_12(time *QTime, formatStr *QStringView /*123*/) *QString /*123*/ {
	var convArg0 = time.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTime11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:996
// index:13
// Public
// QString toString(const class QDateTime &, class QStringView)
func (this *QLocale) ToString_13(dateTime *QDateTime, format *QStringView /*123*/) *QString /*123*/ {
	var convArg0 = dateTime.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTime11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:997
// index:14
// Public
// QString toString(const class QDate &, enum QLocale::FormatType)
func (this *QLocale) ToString_14(date *QDate, format int) *QString /*123*/ {
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:998
// index:15
// Public
// QString toString(const class QTime &, enum QLocale::FormatType)
func (this *QLocale) ToString_15(time *QTime, format int) *QString /*123*/ {
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:999
// index:16
// Public
// QString toString(const class QDateTime &, enum QLocale::FormatType)
func (this *QLocale) ToString_16(dateTime *QDateTime, format int) *QString /*123*/ {
	var convArg0 = dateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1001
// index:0
// Public
// QString dateFormat(enum QLocale::FormatType)
func (this *QLocale) DateFormat(format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10dateFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1002
// index:0
// Public
// QString timeFormat(enum QLocale::FormatType)
func (this *QLocale) TimeFormat(format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10timeFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1003
// index:0
// Public
// QString dateTimeFormat(enum QLocale::FormatType)
func (this *QLocale) DateTimeFormat(format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14dateTimeFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1005
// index:0
// Public
// QDate toDate(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToDate(string *QString, arg1 int) *QDate /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1008
// index:1
// Public
// QDate toDate(const class QString &, const class QString &)
func (this *QLocale) ToDate_1(string *QString, format *QString) *QDate /*123*/ {
	var convArg0 = string.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1006
// index:0
// Public
// QTime toTime(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToTime(string *QString, arg1 int) *QTime /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1009
// index:1
// Public
// QTime toTime(const class QString &, const class QString &)
func (this *QLocale) ToTime_1(string *QString, format *QString) *QTime /*123*/ {
	var convArg0 = string.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1007
// index:0
// Public
// QDateTime toDateTime(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToDateTime(string *QString, format int) *QDateTime /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1010
// index:1
// Public
// QDateTime toDateTime(const class QString &, const class QString &)
func (this *QLocale) ToDateTime_1(string *QString, format *QString) *QDateTime /*123*/ {
	var convArg0 = string.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1015
// index:0
// Public
// QChar decimalPoint()
func (this *QLocale) DecimalPoint() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12decimalPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1016
// index:0
// Public
// QChar groupSeparator()
func (this *QLocale) GroupSeparator() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14groupSeparatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1017
// index:0
// Public
// QChar percent()
func (this *QLocale) Percent() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7percentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1018
// index:0
// Public
// QChar zeroDigit()
func (this *QLocale) ZeroDigit() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9zeroDigitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1019
// index:0
// Public
// QChar negativeSign()
func (this *QLocale) NegativeSign() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12negativeSignEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1020
// index:0
// Public
// QChar positiveSign()
func (this *QLocale) PositiveSign() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12positiveSignEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1021
// index:0
// Public
// QChar exponential()
func (this *QLocale) Exponential() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11exponentialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1023
// index:0
// Public
// QString monthName(int, enum QLocale::FormatType)
func (this *QLocale) MonthName(arg0 int, format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9monthNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1024
// index:0
// Public
// QString standaloneMonthName(int, enum QLocale::FormatType)
func (this *QLocale) StandaloneMonthName(arg0 int, format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale19standaloneMonthNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1025
// index:0
// Public
// QString dayName(int, enum QLocale::FormatType)
func (this *QLocale) DayName(arg0 int, format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7dayNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1026
// index:0
// Public
// QString standaloneDayName(int, enum QLocale::FormatType)
func (this *QLocale) StandaloneDayName(arg0 int, format int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17standaloneDayNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1028
// index:0
// Public
// Qt::DayOfWeek firstDayOfWeek()
func (this *QLocale) FirstDayOfWeek() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14firstDayOfWeekEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1031
// index:0
// Public
// QString amText()
func (this *QLocale) AmText() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6amTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1032
// index:0
// Public
// QString pmText()
func (this *QLocale) PmText() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6pmTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1034
// index:0
// Public
// QLocale::MeasurementSystem measurementSystem()
func (this *QLocale) MeasurementSystem() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17measurementSystemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1036
// index:0
// Public
// Qt::LayoutDirection textDirection()
func (this *QLocale) TextDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale13textDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1038
// index:0
// Public
// QString toUpper(const class QString &)
func (this *QLocale) ToUpper(str *QString) *QString /*123*/ {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toUpperERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1039
// index:0
// Public
// QString toLower(const class QString &)
func (this *QLocale) ToLower(str *QString) *QString /*123*/ {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toLowerERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1041
// index:0
// Public
// QString currencySymbol(enum QLocale::CurrencySymbolFormat)
func (this *QLocale) CurrencySymbol(arg0 int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14currencySymbolENS_20CurrencySymbolFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1042
// index:0
// Public
// QString toCurrencyString(qlonglong, const class QString &)
func (this *QLocale) ToCurrencyString(arg0 int64, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringExRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1043
// index:1
// Public
// QString toCurrencyString(qulonglong, const class QString &)
func (this *QLocale) ToCurrencyString_1(arg0 uint64, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEyRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1044
// index:2
// Public inline
// QString toCurrencyString(short, const class QString &)
func (this *QLocale) ToCurrencyString_2(arg0 int16, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEsRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1045
// index:3
// Public inline
// QString toCurrencyString(ushort, const class QString &)
func (this *QLocale) ToCurrencyString_3(arg0 uint16, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEtRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1046
// index:4
// Public inline
// QString toCurrencyString(int, const class QString &)
func (this *QLocale) ToCurrencyString_4(arg0 int, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1047
// index:5
// Public inline
// QString toCurrencyString(uint, const class QString &)
func (this *QLocale) ToCurrencyString_5(arg0 uint, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEjRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1053
// index:6
// Public
// QString toCurrencyString(double, const class QString &)
func (this *QLocale) ToCurrencyString_6(arg0 float64, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1054
// index:7
// Public
// QString toCurrencyString(double, const class QString &, int)
func (this *QLocale) ToCurrencyString_7(arg0 float64, symbol *QString, precision int) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, convArg1, &precision)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1055
// index:8
// Public inline
// QString toCurrencyString(float, const class QString &)
func (this *QLocale) ToCurrencyString_8(i float32, symbol *QString) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1057
// index:9
// Public inline
// QString toCurrencyString(float, const class QString &, int)
func (this *QLocale) ToCurrencyString_9(i float32, symbol *QString, precision int) *QString /*123*/ {
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, convArg1, &precision)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1068
// index:0
// Public static
// QString languageToString(enum QLocale::Language)
func (this *QLocale) LanguageToString(language int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale16languageToStringENS_8LanguageE", ffiqt.FFI_TYPE_POINTER, language)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QLocale_LanguageToString(language int) *QString /*123*/ {
	var nilthis *QLocale
	rv := nilthis.LanguageToString(language)
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1069
// index:0
// Public static
// QString countryToString(enum QLocale::Country)
func (this *QLocale) CountryToString(country int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale15countryToStringENS_7CountryE", ffiqt.FFI_TYPE_POINTER, country)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QLocale_CountryToString(country int) *QString /*123*/ {
	var nilthis *QLocale
	rv := nilthis.CountryToString(country)
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1070
// index:0
// Public static
// QString scriptToString(enum QLocale::Script)
func (this *QLocale) ScriptToString(script int) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale14scriptToStringENS_6ScriptE", ffiqt.FFI_TYPE_POINTER, script)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QLocale_ScriptToString(script int) *QString /*123*/ {
	var nilthis *QLocale
	rv := nilthis.ScriptToString(script)
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1071
// index:0
// Public static
// void setDefault(const class QLocale &)
func (this *QLocale) SetDefault(locale *QLocale) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale10setDefaultERKS_", ffiqt.FFI_TYPE_POINTER, locale)
	gopp.ErrPrint(err, rv)
}
func QLocale_SetDefault(locale *QLocale) {
	var nilthis *QLocale
	nilthis.SetDefault(locale)
}

// /usr/include/qt/QtCore/qlocale.h:1073
// index:0
// Public static inline
// QLocale c()
func (this *QLocale) C() *QLocale /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale1cEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QLocale_C() *QLocale /*123*/ {
	var nilthis *QLocale
	rv := nilthis.C()
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1074
// index:0
// Public static
// QLocale system()
func (this *QLocale) System() *QLocale /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocale6systemEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QLocale_System() *QLocale /*123*/ {
	var nilthis *QLocale
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1080
// index:0
// Public
// QLocale::NumberOptions numberOptions()
func (this *QLocale) NumberOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale13numberOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1083
// index:0
// Public
// QString quoteString(const class QString &, enum QLocale::QuotationStyle)
func (this *QLocale) QuoteString(str *QString, style int) *QString /*123*/ {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK7QStringNS_14QuotationStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &style)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1084
// index:1
// Public
// QString quoteString(const class QStringRef &, enum QLocale::QuotationStyle)
func (this *QLocale) QuoteString_1(str *QStringRef, style int) *QString /*123*/ {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK10QStringRefNS_14QuotationStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &style)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1086
// index:0
// Public
// QString createSeparatedList(const class QStringList &)
func (this *QLocale) CreateSeparatedList(strl *QStringList) *QString /*123*/ {
	var convArg0 = strl.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale19createSeparatedListERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
