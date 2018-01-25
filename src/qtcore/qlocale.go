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
func (this *QLocale) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLocaleFromPointer(cthis unsafe.Pointer) *QLocale {
	return &QLocale{&qtrt.CObject{cthis}}
}
func (*QLocale) NewFromPointer(cthis unsafe.Pointer) *QLocale {
	return NewQLocaleFromPointer(cthis)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_7CountryE", ffiqt.FFI_TYPE_VOID, cthis, language, country)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_6ScriptENS_7CountryE", ffiqt.FFI_TYPE_VOID, cthis, language, script, country)
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
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:947
// index:0
// Public
// QString bcp47Name()
func (this *QLocale) Bcp47Name() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9bcp47NameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:948
// index:0
// Public
// QString nativeLanguageName()
func (this *QLocale) NativeLanguageName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale18nativeLanguageNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:949
// index:0
// Public
// QString nativeCountryName()
func (this *QLocale) NativeCountryName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17nativeCountryNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:952
// index:0
// Public
// short toShort(const class QString &, _Bool *)
func (this *QLocale) ToShort(s *QString, ok unsafe.Pointer /*666*/) int16 {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toShortE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toUShortE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale5toIntE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toUIntE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toLongLongE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11toULongLongE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toFloatE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK7QStringPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK10QStringRefPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toDoubleE11QStringViewPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qlocale.h:980
// index:0
// Public
// QString toString(qlonglong)
func (this *QLocale) ToString(i int64) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEx", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:981
// index:1
// Public
// QString toString(qulonglong)
func (this *QLocale) ToString_1(i uint64) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEy", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:982
// index:2
// Public inline
// QString toString(short)
func (this *QLocale) ToString_2(i int16) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEs", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:983
// index:3
// Public inline
// QString toString(ushort)
func (this *QLocale) ToString_3(i uint16) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEt", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:984
// index:4
// Public inline
// QString toString(int)
func (this *QLocale) ToString_4(i int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:985
// index:5
// Public inline
// QString toString(uint)
func (this *QLocale) ToString_5(i uint) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEj", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:986
// index:6
// Public
// QString toString(double, char, int)
func (this *QLocale) ToString_6(i float64, f byte, prec int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEdci", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i, f, prec)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:987
// index:7
// Public inline
// QString toString(float, char, int)
func (this *QLocale) ToString_7(i float32, f byte, prec int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringEfci", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i, f, prec)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:990
// index:8
// Public
// QString toString(const class QDate &, const class QString &)
func (this *QLocale) ToString_8(date *QDate, formatStr *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = date.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:991
// index:9
// Public
// QString toString(const class QTime &, const class QString &)
func (this *QLocale) ToString_9(time *QTime, formatStr *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = time.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:992
// index:10
// Public
// QString toString(const class QDateTime &, const class QString &)
func (this *QLocale) ToString_10(dateTime *QDateTime, format *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = dateTime.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:994
// index:11
// Public
// QString toString(const class QDate &, class QStringView)
func (this *QLocale) ToString_11(date *QDate, formatStr *QStringView /*123*/) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = date.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDate11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:995
// index:12
// Public
// QString toString(const class QTime &, class QStringView)
func (this *QLocale) ToString_12(time *QTime, formatStr *QStringView /*123*/) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = time.GetCthis()
	var convArg1 = formatStr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTime11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:996
// index:13
// Public
// QString toString(const class QDateTime &, class QStringView)
func (this *QLocale) ToString_13(dateTime *QDateTime, format *QStringView /*123*/) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = dateTime.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTime11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:997
// index:14
// Public
// QString toString(const class QDate &, enum QLocale::FormatType)
func (this *QLocale) ToString_14(date *QDate, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:998
// index:15
// Public
// QString toString(const class QTime &, enum QLocale::FormatType)
func (this *QLocale) ToString_15(time *QTime, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:999
// index:16
// Public
// QString toString(const class QDateTime &, enum QLocale::FormatType)
func (this *QLocale) ToString_16(dateTime *QDateTime, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = dateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1001
// index:0
// Public
// QString dateFormat(enum QLocale::FormatType)
func (this *QLocale) DateFormat(format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10dateFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1002
// index:0
// Public
// QString timeFormat(enum QLocale::FormatType)
func (this *QLocale) TimeFormat(format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10timeFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1003
// index:0
// Public
// QString dateTimeFormat(enum QLocale::FormatType)
func (this *QLocale) DateTimeFormat(format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14dateTimeFormatENS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1005
// index:0
// Public
// QDate toDate(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToDate(string *QString, arg1 int) *QDate /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1008
// index:1
// Public
// QDate toDate(const class QString &, const class QString &)
func (this *QLocale) ToDate_1(string *QString, format *QString) *QDate /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toDateERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1006
// index:0
// Public
// QTime toTime(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToTime(string *QString, arg1 int) *QTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1009
// index:1
// Public
// QTime toTime(const class QString &, const class QString &)
func (this *QLocale) ToTime_1(string *QString, format *QString) *QTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6toTimeERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1007
// index:0
// Public
// QDateTime toDateTime(const class QString &, enum QLocale::FormatType)
func (this *QLocale) ToDateTime(string *QString, format int) *QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1010
// index:1
// Public
// QDateTime toDateTime(const class QString &, const class QString &)
func (this *QLocale) ToDateTime_1(string *QString, format *QString) *QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale10toDateTimeERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1015
// index:0
// Public
// QChar decimalPoint()
func (this *QLocale) DecimalPoint() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12decimalPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1016
// index:0
// Public
// QChar groupSeparator()
func (this *QLocale) GroupSeparator() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14groupSeparatorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1017
// index:0
// Public
// QChar percent()
func (this *QLocale) Percent() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7percentEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1018
// index:0
// Public
// QChar zeroDigit()
func (this *QLocale) ZeroDigit() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9zeroDigitEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1019
// index:0
// Public
// QChar negativeSign()
func (this *QLocale) NegativeSign() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12negativeSignEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1020
// index:0
// Public
// QChar positiveSign()
func (this *QLocale) PositiveSign() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale12positiveSignEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1021
// index:0
// Public
// QChar exponential()
func (this *QLocale) Exponential() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11exponentialEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1023
// index:0
// Public
// QString monthName(int, enum QLocale::FormatType)
func (this *QLocale) MonthName(arg0 int, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale9monthNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1024
// index:0
// Public
// QString standaloneMonthName(int, enum QLocale::FormatType)
func (this *QLocale) StandaloneMonthName(arg0 int, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale19standaloneMonthNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1025
// index:0
// Public
// QString dayName(int, enum QLocale::FormatType)
func (this *QLocale) DayName(arg0 int, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7dayNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1026
// index:0
// Public
// QString standaloneDayName(int, enum QLocale::FormatType)
func (this *QLocale) StandaloneDayName(arg0 int, format int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale17standaloneDayNameEiNS_10FormatTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
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
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6amTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1032
// index:0
// Public
// QString pmText()
func (this *QLocale) PmText() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale6pmTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
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
	mv := qtrt.Calloc(1, 256)
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toUpperERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1039
// index:0
// Public
// QString toLower(const class QString &)
func (this *QLocale) ToLower(str *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale7toLowerERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1041
// index:0
// Public
// QString currencySymbol(enum QLocale::CurrencySymbolFormat)
func (this *QLocale) CurrencySymbol(arg0 int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale14currencySymbolENS_20CurrencySymbolFormatE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1042
// index:0
// Public
// QString toCurrencyString(qlonglong, const class QString &)
func (this *QLocale) ToCurrencyString(arg0 int64, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringExRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1043
// index:1
// Public
// QString toCurrencyString(qulonglong, const class QString &)
func (this *QLocale) ToCurrencyString_1(arg0 uint64, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEyRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1044
// index:2
// Public inline
// QString toCurrencyString(short, const class QString &)
func (this *QLocale) ToCurrencyString_2(arg0 int16, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEsRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1045
// index:3
// Public inline
// QString toCurrencyString(ushort, const class QString &)
func (this *QLocale) ToCurrencyString_3(arg0 uint16, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEtRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1046
// index:4
// Public inline
// QString toCurrencyString(int, const class QString &)
func (this *QLocale) ToCurrencyString_4(arg0 int, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEiRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1047
// index:5
// Public inline
// QString toCurrencyString(uint, const class QString &)
func (this *QLocale) ToCurrencyString_5(arg0 uint, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEjRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1053
// index:6
// Public
// QString toCurrencyString(double, const class QString &)
func (this *QLocale) ToCurrencyString_6(arg0 float64, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1054
// index:7
// Public
// QString toCurrencyString(double, const class QString &, int)
func (this *QLocale) ToCurrencyString_7(arg0 float64, symbol *QString, precision int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QStringi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0, convArg1, precision)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1055
// index:8
// Public inline
// QString toCurrencyString(float, const class QString &)
func (this *QLocale) ToCurrencyString_8(i float32, symbol *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1057
// index:9
// Public inline
// QString toCurrencyString(float, const class QString &, int)
func (this *QLocale) ToCurrencyString_9(i float32, symbol *QString, precision int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = symbol.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QStringi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i, convArg1, precision)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
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
	mv := qtrt.Calloc(1, 256)
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK7QStringNS_14QuotationStyleE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, style)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1084
// index:1
// Public
// QString quoteString(const class QStringRef &, enum QLocale::QuotationStyle)
func (this *QLocale) QuoteString_1(str *QStringRef, style int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK10QStringRefNS_14QuotationStyleE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, style)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1086
// index:0
// Public
// QString createSeparatedList(const class QStringList &)
func (this *QLocale) CreateSeparatedList(strl *QStringList) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = strl.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLocale19createSeparatedListERK11QStringList", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QLocale__Language = int

const QLocale__AnyLanguage QLocale__Language = 0
const QLocale__C QLocale__Language = 1
const QLocale__Abkhazian QLocale__Language = 2
const QLocale__Oromo QLocale__Language = 3
const QLocale__Afar QLocale__Language = 4
const QLocale__Afrikaans QLocale__Language = 5
const QLocale__Albanian QLocale__Language = 6
const QLocale__Amharic QLocale__Language = 7
const QLocale__Arabic QLocale__Language = 8
const QLocale__Armenian QLocale__Language = 9
const QLocale__Assamese QLocale__Language = 10
const QLocale__Aymara QLocale__Language = 11
const QLocale__Azerbaijani QLocale__Language = 12
const QLocale__Bashkir QLocale__Language = 13
const QLocale__Basque QLocale__Language = 14
const QLocale__Bengali QLocale__Language = 15
const QLocale__Dzongkha QLocale__Language = 16
const QLocale__Bihari QLocale__Language = 17
const QLocale__Bislama QLocale__Language = 18
const QLocale__Breton QLocale__Language = 19
const QLocale__Bulgarian QLocale__Language = 20
const QLocale__Burmese QLocale__Language = 21
const QLocale__Belarusian QLocale__Language = 22
const QLocale__Khmer QLocale__Language = 23
const QLocale__Catalan QLocale__Language = 24
const QLocale__Chinese QLocale__Language = 25
const QLocale__Corsican QLocale__Language = 26
const QLocale__Croatian QLocale__Language = 27
const QLocale__Czech QLocale__Language = 28
const QLocale__Danish QLocale__Language = 29
const QLocale__Dutch QLocale__Language = 30
const QLocale__English QLocale__Language = 31
const QLocale__Esperanto QLocale__Language = 32
const QLocale__Estonian QLocale__Language = 33
const QLocale__Faroese QLocale__Language = 34
const QLocale__Fijian QLocale__Language = 35
const QLocale__Finnish QLocale__Language = 36
const QLocale__French QLocale__Language = 37
const QLocale__WesternFrisian QLocale__Language = 38
const QLocale__Gaelic QLocale__Language = 39
const QLocale__Galician QLocale__Language = 40
const QLocale__Georgian QLocale__Language = 41
const QLocale__German QLocale__Language = 42
const QLocale__Greek QLocale__Language = 43
const QLocale__Greenlandic QLocale__Language = 44
const QLocale__Guarani QLocale__Language = 45
const QLocale__Gujarati QLocale__Language = 46
const QLocale__Hausa QLocale__Language = 47
const QLocale__Hebrew QLocale__Language = 48
const QLocale__Hindi QLocale__Language = 49
const QLocale__Hungarian QLocale__Language = 50
const QLocale__Icelandic QLocale__Language = 51
const QLocale__Indonesian QLocale__Language = 52
const QLocale__Interlingua QLocale__Language = 53
const QLocale__Interlingue QLocale__Language = 54
const QLocale__Inuktitut QLocale__Language = 55
const QLocale__Inupiak QLocale__Language = 56
const QLocale__Irish QLocale__Language = 57
const QLocale__Italian QLocale__Language = 58
const QLocale__Japanese QLocale__Language = 59
const QLocale__Javanese QLocale__Language = 60
const QLocale__Kannada QLocale__Language = 61
const QLocale__Kashmiri QLocale__Language = 62
const QLocale__Kazakh QLocale__Language = 63
const QLocale__Kinyarwanda QLocale__Language = 64
const QLocale__Kirghiz QLocale__Language = 65
const QLocale__Korean QLocale__Language = 66
const QLocale__Kurdish QLocale__Language = 67
const QLocale__Rundi QLocale__Language = 68
const QLocale__Lao QLocale__Language = 69
const QLocale__Latin QLocale__Language = 70
const QLocale__Latvian QLocale__Language = 71
const QLocale__Lingala QLocale__Language = 72
const QLocale__Lithuanian QLocale__Language = 73
const QLocale__Macedonian QLocale__Language = 74
const QLocale__Malagasy QLocale__Language = 75
const QLocale__Malay QLocale__Language = 76
const QLocale__Malayalam QLocale__Language = 77
const QLocale__Maltese QLocale__Language = 78
const QLocale__Maori QLocale__Language = 79
const QLocale__Marathi QLocale__Language = 80
const QLocale__Marshallese QLocale__Language = 81
const QLocale__Mongolian QLocale__Language = 82
const QLocale__NauruLanguage QLocale__Language = 83
const QLocale__Nepali QLocale__Language = 84
const QLocale__NorwegianBokmal QLocale__Language = 85
const QLocale__Occitan QLocale__Language = 86
const QLocale__Oriya QLocale__Language = 87
const QLocale__Pashto QLocale__Language = 88
const QLocale__Persian QLocale__Language = 89
const QLocale__Polish QLocale__Language = 90
const QLocale__Portuguese QLocale__Language = 91
const QLocale__Punjabi QLocale__Language = 92
const QLocale__Quechua QLocale__Language = 93
const QLocale__Romansh QLocale__Language = 94
const QLocale__Romanian QLocale__Language = 95
const QLocale__Russian QLocale__Language = 96
const QLocale__Samoan QLocale__Language = 97
const QLocale__Sango QLocale__Language = 98
const QLocale__Sanskrit QLocale__Language = 99
const QLocale__Serbian QLocale__Language = 100
const QLocale__Ossetic QLocale__Language = 101
const QLocale__SouthernSotho QLocale__Language = 102
const QLocale__Tswana QLocale__Language = 103
const QLocale__Shona QLocale__Language = 104
const QLocale__Sindhi QLocale__Language = 105
const QLocale__Sinhala QLocale__Language = 106
const QLocale__Swati QLocale__Language = 107
const QLocale__Slovak QLocale__Language = 108
const QLocale__Slovenian QLocale__Language = 109
const QLocale__Somali QLocale__Language = 110
const QLocale__Spanish QLocale__Language = 111
const QLocale__Sundanese QLocale__Language = 112
const QLocale__Swahili QLocale__Language = 113
const QLocale__Swedish QLocale__Language = 114
const QLocale__Sardinian QLocale__Language = 115
const QLocale__Tajik QLocale__Language = 116
const QLocale__Tamil QLocale__Language = 117
const QLocale__Tatar QLocale__Language = 118
const QLocale__Telugu QLocale__Language = 119
const QLocale__Thai QLocale__Language = 120
const QLocale__Tibetan QLocale__Language = 121
const QLocale__Tigrinya QLocale__Language = 122
const QLocale__Tongan QLocale__Language = 123
const QLocale__Tsonga QLocale__Language = 124
const QLocale__Turkish QLocale__Language = 125
const QLocale__Turkmen QLocale__Language = 126
const QLocale__Tahitian QLocale__Language = 127
const QLocale__Uighur QLocale__Language = 128
const QLocale__Ukrainian QLocale__Language = 129
const QLocale__Urdu QLocale__Language = 130
const QLocale__Uzbek QLocale__Language = 131
const QLocale__Vietnamese QLocale__Language = 132
const QLocale__Volapuk QLocale__Language = 133
const QLocale__Welsh QLocale__Language = 134
const QLocale__Wolof QLocale__Language = 135
const QLocale__Xhosa QLocale__Language = 136
const QLocale__Yiddish QLocale__Language = 137
const QLocale__Yoruba QLocale__Language = 138
const QLocale__Zhuang QLocale__Language = 139
const QLocale__Zulu QLocale__Language = 140
const QLocale__NorwegianNynorsk QLocale__Language = 141
const QLocale__Bosnian QLocale__Language = 142
const QLocale__Divehi QLocale__Language = 143
const QLocale__Manx QLocale__Language = 144
const QLocale__Cornish QLocale__Language = 145
const QLocale__Akan QLocale__Language = 146
const QLocale__Konkani QLocale__Language = 147
const QLocale__Ga QLocale__Language = 148
const QLocale__Igbo QLocale__Language = 149
const QLocale__Kamba QLocale__Language = 150
const QLocale__Syriac QLocale__Language = 151
const QLocale__Blin QLocale__Language = 152
const QLocale__Geez QLocale__Language = 153
const QLocale__Koro QLocale__Language = 154
const QLocale__Sidamo QLocale__Language = 155
const QLocale__Atsam QLocale__Language = 156
const QLocale__Tigre QLocale__Language = 157
const QLocale__Jju QLocale__Language = 158
const QLocale__Friulian QLocale__Language = 159
const QLocale__Venda QLocale__Language = 160
const QLocale__Ewe QLocale__Language = 161
const QLocale__Walamo QLocale__Language = 162
const QLocale__Hawaiian QLocale__Language = 163
const QLocale__Tyap QLocale__Language = 164
const QLocale__Nyanja QLocale__Language = 165
const QLocale__Filipino QLocale__Language = 166
const QLocale__SwissGerman QLocale__Language = 167
const QLocale__SichuanYi QLocale__Language = 168
const QLocale__Kpelle QLocale__Language = 169
const QLocale__LowGerman QLocale__Language = 170
const QLocale__SouthNdebele QLocale__Language = 171
const QLocale__NorthernSotho QLocale__Language = 172
const QLocale__NorthernSami QLocale__Language = 173
const QLocale__Taroko QLocale__Language = 174
const QLocale__Gusii QLocale__Language = 175
const QLocale__Taita QLocale__Language = 176
const QLocale__Fulah QLocale__Language = 177
const QLocale__Kikuyu QLocale__Language = 178
const QLocale__Samburu QLocale__Language = 179
const QLocale__Sena QLocale__Language = 180
const QLocale__NorthNdebele QLocale__Language = 181
const QLocale__Rombo QLocale__Language = 182
const QLocale__Tachelhit QLocale__Language = 183
const QLocale__Kabyle QLocale__Language = 184
const QLocale__Nyankole QLocale__Language = 185
const QLocale__Bena QLocale__Language = 186
const QLocale__Vunjo QLocale__Language = 187
const QLocale__Bambara QLocale__Language = 188
const QLocale__Embu QLocale__Language = 189
const QLocale__Cherokee QLocale__Language = 190
const QLocale__Morisyen QLocale__Language = 191
const QLocale__Makonde QLocale__Language = 192
const QLocale__Langi QLocale__Language = 193
const QLocale__Ganda QLocale__Language = 194
const QLocale__Bemba QLocale__Language = 195
const QLocale__Kabuverdianu QLocale__Language = 196
const QLocale__Meru QLocale__Language = 197
const QLocale__Kalenjin QLocale__Language = 198
const QLocale__Nama QLocale__Language = 199
const QLocale__Machame QLocale__Language = 200
const QLocale__Colognian QLocale__Language = 201
const QLocale__Masai QLocale__Language = 202
const QLocale__Soga QLocale__Language = 203
const QLocale__Luyia QLocale__Language = 204
const QLocale__Asu QLocale__Language = 205
const QLocale__Teso QLocale__Language = 206
const QLocale__Saho QLocale__Language = 207
const QLocale__KoyraChiini QLocale__Language = 208
const QLocale__Rwa QLocale__Language = 209
const QLocale__Luo QLocale__Language = 210
const QLocale__Chiga QLocale__Language = 211
const QLocale__CentralMoroccoTamazight QLocale__Language = 212
const QLocale__KoyraboroSenni QLocale__Language = 213
const QLocale__Shambala QLocale__Language = 214
const QLocale__Bodo QLocale__Language = 215
const QLocale__Avaric QLocale__Language = 216
const QLocale__Chamorro QLocale__Language = 217
const QLocale__Chechen QLocale__Language = 218
const QLocale__Church QLocale__Language = 219
const QLocale__Chuvash QLocale__Language = 220
const QLocale__Cree QLocale__Language = 221
const QLocale__Haitian QLocale__Language = 222
const QLocale__Herero QLocale__Language = 223
const QLocale__HiriMotu QLocale__Language = 224
const QLocale__Kanuri QLocale__Language = 225
const QLocale__Komi QLocale__Language = 226
const QLocale__Kongo QLocale__Language = 227
const QLocale__Kwanyama QLocale__Language = 228
const QLocale__Limburgish QLocale__Language = 229
const QLocale__LubaKatanga QLocale__Language = 230
const QLocale__Luxembourgish QLocale__Language = 231
const QLocale__Navaho QLocale__Language = 232
const QLocale__Ndonga QLocale__Language = 233
const QLocale__Ojibwa QLocale__Language = 234
const QLocale__Pali QLocale__Language = 235
const QLocale__Walloon QLocale__Language = 236
const QLocale__Aghem QLocale__Language = 237
const QLocale__Basaa QLocale__Language = 238
const QLocale__Zarma QLocale__Language = 239
const QLocale__Duala QLocale__Language = 240
const QLocale__JolaFonyi QLocale__Language = 241
const QLocale__Ewondo QLocale__Language = 242
const QLocale__Bafia QLocale__Language = 243
const QLocale__MakhuwaMeetto QLocale__Language = 244
const QLocale__Mundang QLocale__Language = 245
const QLocale__Kwasio QLocale__Language = 246
const QLocale__Nuer QLocale__Language = 247
const QLocale__Sakha QLocale__Language = 248
const QLocale__Sangu QLocale__Language = 249
const QLocale__CongoSwahili QLocale__Language = 250
const QLocale__Tasawaq QLocale__Language = 251
const QLocale__Vai QLocale__Language = 252
const QLocale__Walser QLocale__Language = 253
const QLocale__Yangben QLocale__Language = 254
const QLocale__Avestan QLocale__Language = 255
const QLocale__Asturian QLocale__Language = 256
const QLocale__Ngomba QLocale__Language = 257
const QLocale__Kako QLocale__Language = 258
const QLocale__Meta QLocale__Language = 259
const QLocale__Ngiemboon QLocale__Language = 260
const QLocale__Aragonese QLocale__Language = 261
const QLocale__Akkadian QLocale__Language = 262
const QLocale__AncientEgyptian QLocale__Language = 263
const QLocale__AncientGreek QLocale__Language = 264
const QLocale__Aramaic QLocale__Language = 265
const QLocale__Balinese QLocale__Language = 266
const QLocale__Bamun QLocale__Language = 267
const QLocale__BatakToba QLocale__Language = 268
const QLocale__Buginese QLocale__Language = 269
const QLocale__Buhid QLocale__Language = 270
const QLocale__Carian QLocale__Language = 271
const QLocale__Chakma QLocale__Language = 272
const QLocale__ClassicalMandaic QLocale__Language = 273
const QLocale__Coptic QLocale__Language = 274
const QLocale__Dogri QLocale__Language = 275
const QLocale__EasternCham QLocale__Language = 276
const QLocale__EasternKayah QLocale__Language = 277
const QLocale__Etruscan QLocale__Language = 278
const QLocale__Gothic QLocale__Language = 279
const QLocale__Hanunoo QLocale__Language = 280
const QLocale__Ingush QLocale__Language = 281
const QLocale__LargeFloweryMiao QLocale__Language = 282
const QLocale__Lepcha QLocale__Language = 283
const QLocale__Limbu QLocale__Language = 284
const QLocale__Lisu QLocale__Language = 285
const QLocale__Lu QLocale__Language = 286
const QLocale__Lycian QLocale__Language = 287
const QLocale__Lydian QLocale__Language = 288
const QLocale__Mandingo QLocale__Language = 289
const QLocale__Manipuri QLocale__Language = 290
const QLocale__Meroitic QLocale__Language = 291
const QLocale__NorthernThai QLocale__Language = 292
const QLocale__OldIrish QLocale__Language = 293
const QLocale__OldNorse QLocale__Language = 294
const QLocale__OldPersian QLocale__Language = 295
const QLocale__OldTurkish QLocale__Language = 296
const QLocale__Pahlavi QLocale__Language = 297
const QLocale__Parthian QLocale__Language = 298
const QLocale__Phoenician QLocale__Language = 299
const QLocale__PrakritLanguage QLocale__Language = 300
const QLocale__Rejang QLocale__Language = 301
const QLocale__Sabaean QLocale__Language = 302
const QLocale__Samaritan QLocale__Language = 303
const QLocale__Santali QLocale__Language = 304
const QLocale__Saurashtra QLocale__Language = 305
const QLocale__Sora QLocale__Language = 306
const QLocale__Sylheti QLocale__Language = 307
const QLocale__Tagbanwa QLocale__Language = 308
const QLocale__TaiDam QLocale__Language = 309
const QLocale__TaiNua QLocale__Language = 310
const QLocale__Ugaritic QLocale__Language = 311
const QLocale__Akoose QLocale__Language = 312
const QLocale__Lakota QLocale__Language = 313
const QLocale__StandardMoroccanTamazight QLocale__Language = 314
const QLocale__Mapuche QLocale__Language = 315
const QLocale__CentralKurdish QLocale__Language = 316
const QLocale__LowerSorbian QLocale__Language = 317
const QLocale__UpperSorbian QLocale__Language = 318
const QLocale__Kenyang QLocale__Language = 319
const QLocale__Mohawk QLocale__Language = 320
const QLocale__Nko QLocale__Language = 321
const QLocale__Prussian QLocale__Language = 322
const QLocale__Kiche QLocale__Language = 323
const QLocale__SouthernSami QLocale__Language = 324
const QLocale__LuleSami QLocale__Language = 325
const QLocale__InariSami QLocale__Language = 326
const QLocale__SkoltSami QLocale__Language = 327
const QLocale__Warlpiri QLocale__Language = 328
const QLocale__ManichaeanMiddlePersian QLocale__Language = 329
const QLocale__Mende QLocale__Language = 330
const QLocale__AncientNorthArabian QLocale__Language = 331
const QLocale__LinearA QLocale__Language = 332
const QLocale__HmongNjua QLocale__Language = 333
const QLocale__Ho QLocale__Language = 334
const QLocale__Lezghian QLocale__Language = 335
const QLocale__Bassa QLocale__Language = 336
const QLocale__Mono QLocale__Language = 337
const QLocale__TedimChin QLocale__Language = 338
const QLocale__Maithili QLocale__Language = 339
const QLocale__Ahom QLocale__Language = 340
const QLocale__AmericanSignLanguage QLocale__Language = 341
const QLocale__ArdhamagadhiPrakrit QLocale__Language = 342
const QLocale__Bhojpuri QLocale__Language = 343
const QLocale__HieroglyphicLuwian QLocale__Language = 344
const QLocale__LiteraryChinese QLocale__Language = 345
const QLocale__Mazanderani QLocale__Language = 346
const QLocale__Mru QLocale__Language = 347
const QLocale__Newari QLocale__Language = 348
const QLocale__NorthernLuri QLocale__Language = 349
const QLocale__Palauan QLocale__Language = 350
const QLocale__Papiamento QLocale__Language = 351
const QLocale__Saraiki QLocale__Language = 352
const QLocale__TokelauLanguage QLocale__Language = 353
const QLocale__TokPisin QLocale__Language = 354
const QLocale__TuvaluLanguage QLocale__Language = 355
const QLocale__UncodedLanguages QLocale__Language = 356
const QLocale__Cantonese QLocale__Language = 357
const QLocale__Osage QLocale__Language = 358
const QLocale__Tangut QLocale__Language = 359
const QLocale__Norwegian QLocale__Language = 85
const QLocale__Moldavian QLocale__Language = 95
const QLocale__SerboCroatian QLocale__Language = 100
const QLocale__Tagalog QLocale__Language = 166
const QLocale__Twi QLocale__Language = 146
const QLocale__Afan QLocale__Language = 3
const QLocale__Byelorussian QLocale__Language = 22
const QLocale__Bhutani QLocale__Language = 16
const QLocale__Cambodian QLocale__Language = 23
const QLocale__Kurundi QLocale__Language = 68
const QLocale__RhaetoRomance QLocale__Language = 94
const QLocale__Chewa QLocale__Language = 165
const QLocale__Frisian QLocale__Language = 38
const QLocale__Uigur QLocale__Language = 128
const QLocale__LastLanguage QLocale__Language = 359

type QLocale__Script = int

const QLocale__AnyScript QLocale__Script = 0
const QLocale__ArabicScript QLocale__Script = 1
const QLocale__CyrillicScript QLocale__Script = 2
const QLocale__DeseretScript QLocale__Script = 3
const QLocale__GurmukhiScript QLocale__Script = 4
const QLocale__SimplifiedHanScript QLocale__Script = 5
const QLocale__TraditionalHanScript QLocale__Script = 6
const QLocale__LatinScript QLocale__Script = 7
const QLocale__MongolianScript QLocale__Script = 8
const QLocale__TifinaghScript QLocale__Script = 9
const QLocale__ArmenianScript QLocale__Script = 10
const QLocale__BengaliScript QLocale__Script = 11
const QLocale__CherokeeScript QLocale__Script = 12
const QLocale__DevanagariScript QLocale__Script = 13
const QLocale__EthiopicScript QLocale__Script = 14
const QLocale__GeorgianScript QLocale__Script = 15
const QLocale__GreekScript QLocale__Script = 16
const QLocale__GujaratiScript QLocale__Script = 17
const QLocale__HebrewScript QLocale__Script = 18
const QLocale__JapaneseScript QLocale__Script = 19
const QLocale__KhmerScript QLocale__Script = 20
const QLocale__KannadaScript QLocale__Script = 21
const QLocale__KoreanScript QLocale__Script = 22
const QLocale__LaoScript QLocale__Script = 23
const QLocale__MalayalamScript QLocale__Script = 24
const QLocale__MyanmarScript QLocale__Script = 25
const QLocale__OriyaScript QLocale__Script = 26
const QLocale__TamilScript QLocale__Script = 27
const QLocale__TeluguScript QLocale__Script = 28
const QLocale__ThaanaScript QLocale__Script = 29
const QLocale__ThaiScript QLocale__Script = 30
const QLocale__TibetanScript QLocale__Script = 31
const QLocale__SinhalaScript QLocale__Script = 32
const QLocale__SyriacScript QLocale__Script = 33
const QLocale__YiScript QLocale__Script = 34
const QLocale__VaiScript QLocale__Script = 35
const QLocale__AvestanScript QLocale__Script = 36
const QLocale__BalineseScript QLocale__Script = 37
const QLocale__BamumScript QLocale__Script = 38
const QLocale__BatakScript QLocale__Script = 39
const QLocale__BopomofoScript QLocale__Script = 40
const QLocale__BrahmiScript QLocale__Script = 41
const QLocale__BugineseScript QLocale__Script = 42
const QLocale__BuhidScript QLocale__Script = 43
const QLocale__CanadianAboriginalScript QLocale__Script = 44
const QLocale__CarianScript QLocale__Script = 45
const QLocale__ChakmaScript QLocale__Script = 46
const QLocale__ChamScript QLocale__Script = 47
const QLocale__CopticScript QLocale__Script = 48
const QLocale__CypriotScript QLocale__Script = 49
const QLocale__EgyptianHieroglyphsScript QLocale__Script = 50
const QLocale__FraserScript QLocale__Script = 51
const QLocale__GlagoliticScript QLocale__Script = 52
const QLocale__GothicScript QLocale__Script = 53
const QLocale__HanScript QLocale__Script = 54
const QLocale__HangulScript QLocale__Script = 55
const QLocale__HanunooScript QLocale__Script = 56
const QLocale__ImperialAramaicScript QLocale__Script = 57
const QLocale__InscriptionalPahlaviScript QLocale__Script = 58
const QLocale__InscriptionalParthianScript QLocale__Script = 59
const QLocale__JavaneseScript QLocale__Script = 60
const QLocale__KaithiScript QLocale__Script = 61
const QLocale__KatakanaScript QLocale__Script = 62
const QLocale__KayahLiScript QLocale__Script = 63
const QLocale__KharoshthiScript QLocale__Script = 64
const QLocale__LannaScript QLocale__Script = 65
const QLocale__LepchaScript QLocale__Script = 66
const QLocale__LimbuScript QLocale__Script = 67
const QLocale__LinearBScript QLocale__Script = 68
const QLocale__LycianScript QLocale__Script = 69
const QLocale__LydianScript QLocale__Script = 70
const QLocale__MandaeanScript QLocale__Script = 71
const QLocale__MeiteiMayekScript QLocale__Script = 72
const QLocale__MeroiticScript QLocale__Script = 73
const QLocale__MeroiticCursiveScript QLocale__Script = 74
const QLocale__NkoScript QLocale__Script = 75
const QLocale__NewTaiLueScript QLocale__Script = 76
const QLocale__OghamScript QLocale__Script = 77
const QLocale__OlChikiScript QLocale__Script = 78
const QLocale__OldItalicScript QLocale__Script = 79
const QLocale__OldPersianScript QLocale__Script = 80
const QLocale__OldSouthArabianScript QLocale__Script = 81
const QLocale__OrkhonScript QLocale__Script = 82
const QLocale__OsmanyaScript QLocale__Script = 83
const QLocale__PhagsPaScript QLocale__Script = 84
const QLocale__PhoenicianScript QLocale__Script = 85
const QLocale__PollardPhoneticScript QLocale__Script = 86
const QLocale__RejangScript QLocale__Script = 87
const QLocale__RunicScript QLocale__Script = 88
const QLocale__SamaritanScript QLocale__Script = 89
const QLocale__SaurashtraScript QLocale__Script = 90
const QLocale__SharadaScript QLocale__Script = 91
const QLocale__ShavianScript QLocale__Script = 92
const QLocale__SoraSompengScript QLocale__Script = 93
const QLocale__CuneiformScript QLocale__Script = 94
const QLocale__SundaneseScript QLocale__Script = 95
const QLocale__SylotiNagriScript QLocale__Script = 96
const QLocale__TagalogScript QLocale__Script = 97
const QLocale__TagbanwaScript QLocale__Script = 98
const QLocale__TaiLeScript QLocale__Script = 99
const QLocale__TaiVietScript QLocale__Script = 100
const QLocale__TakriScript QLocale__Script = 101
const QLocale__UgariticScript QLocale__Script = 102
const QLocale__BrailleScript QLocale__Script = 103
const QLocale__HiraganaScript QLocale__Script = 104
const QLocale__CaucasianAlbanianScript QLocale__Script = 105
const QLocale__BassaVahScript QLocale__Script = 106
const QLocale__DuployanScript QLocale__Script = 107
const QLocale__ElbasanScript QLocale__Script = 108
const QLocale__GranthaScript QLocale__Script = 109
const QLocale__PahawhHmongScript QLocale__Script = 110
const QLocale__KhojkiScript QLocale__Script = 111
const QLocale__LinearAScript QLocale__Script = 112
const QLocale__MahajaniScript QLocale__Script = 113
const QLocale__ManichaeanScript QLocale__Script = 114
const QLocale__MendeKikakuiScript QLocale__Script = 115
const QLocale__ModiScript QLocale__Script = 116
const QLocale__MroScript QLocale__Script = 117
const QLocale__OldNorthArabianScript QLocale__Script = 118
const QLocale__NabataeanScript QLocale__Script = 119
const QLocale__PalmyreneScript QLocale__Script = 120
const QLocale__PauCinHauScript QLocale__Script = 121
const QLocale__OldPermicScript QLocale__Script = 122
const QLocale__PsalterPahlaviScript QLocale__Script = 123
const QLocale__SiddhamScript QLocale__Script = 124
const QLocale__KhudawadiScript QLocale__Script = 125
const QLocale__TirhutaScript QLocale__Script = 126
const QLocale__VarangKshitiScript QLocale__Script = 127
const QLocale__AhomScript QLocale__Script = 128
const QLocale__AnatolianHieroglyphsScript QLocale__Script = 129
const QLocale__HatranScript QLocale__Script = 130
const QLocale__MultaniScript QLocale__Script = 131
const QLocale__OldHungarianScript QLocale__Script = 132
const QLocale__SignWritingScript QLocale__Script = 133
const QLocale__AdlamScript QLocale__Script = 134
const QLocale__BhaiksukiScript QLocale__Script = 135
const QLocale__MarchenScript QLocale__Script = 136
const QLocale__NewaScript QLocale__Script = 137
const QLocale__OsageScript QLocale__Script = 138
const QLocale__TangutScript QLocale__Script = 139
const QLocale__HanWithBopomofoScript QLocale__Script = 140
const QLocale__JamoScript QLocale__Script = 141
const QLocale__SimplifiedChineseScript QLocale__Script = 5
const QLocale__TraditionalChineseScript QLocale__Script = 6
const QLocale__LastScript QLocale__Script = 141

type QLocale__Country = int

const QLocale__AnyCountry QLocale__Country = 0
const QLocale__Afghanistan QLocale__Country = 1
const QLocale__Albania QLocale__Country = 2
const QLocale__Algeria QLocale__Country = 3
const QLocale__AmericanSamoa QLocale__Country = 4
const QLocale__Andorra QLocale__Country = 5
const QLocale__Angola QLocale__Country = 6
const QLocale__Anguilla QLocale__Country = 7
const QLocale__Antarctica QLocale__Country = 8
const QLocale__AntiguaAndBarbuda QLocale__Country = 9
const QLocale__Argentina QLocale__Country = 10
const QLocale__Armenia QLocale__Country = 11
const QLocale__Aruba QLocale__Country = 12
const QLocale__Australia QLocale__Country = 13
const QLocale__Austria QLocale__Country = 14
const QLocale__Azerbaijan QLocale__Country = 15
const QLocale__Bahamas QLocale__Country = 16
const QLocale__Bahrain QLocale__Country = 17
const QLocale__Bangladesh QLocale__Country = 18
const QLocale__Barbados QLocale__Country = 19
const QLocale__Belarus QLocale__Country = 20
const QLocale__Belgium QLocale__Country = 21
const QLocale__Belize QLocale__Country = 22
const QLocale__Benin QLocale__Country = 23
const QLocale__Bermuda QLocale__Country = 24
const QLocale__Bhutan QLocale__Country = 25
const QLocale__Bolivia QLocale__Country = 26
const QLocale__BosniaAndHerzegowina QLocale__Country = 27
const QLocale__Botswana QLocale__Country = 28
const QLocale__BouvetIsland QLocale__Country = 29
const QLocale__Brazil QLocale__Country = 30
const QLocale__BritishIndianOceanTerritory QLocale__Country = 31
const QLocale__Brunei QLocale__Country = 32
const QLocale__Bulgaria QLocale__Country = 33
const QLocale__BurkinaFaso QLocale__Country = 34
const QLocale__Burundi QLocale__Country = 35
const QLocale__Cambodia QLocale__Country = 36
const QLocale__Cameroon QLocale__Country = 37
const QLocale__Canada QLocale__Country = 38
const QLocale__CapeVerde QLocale__Country = 39
const QLocale__CaymanIslands QLocale__Country = 40
const QLocale__CentralAfricanRepublic QLocale__Country = 41
const QLocale__Chad QLocale__Country = 42
const QLocale__Chile QLocale__Country = 43
const QLocale__China QLocale__Country = 44
const QLocale__ChristmasIsland QLocale__Country = 45
const QLocale__CocosIslands QLocale__Country = 46
const QLocale__Colombia QLocale__Country = 47
const QLocale__Comoros QLocale__Country = 48
const QLocale__CongoKinshasa QLocale__Country = 49
const QLocale__CongoBrazzaville QLocale__Country = 50
const QLocale__CookIslands QLocale__Country = 51
const QLocale__CostaRica QLocale__Country = 52
const QLocale__IvoryCoast QLocale__Country = 53
const QLocale__Croatia QLocale__Country = 54
const QLocale__Cuba QLocale__Country = 55
const QLocale__Cyprus QLocale__Country = 56
const QLocale__CzechRepublic QLocale__Country = 57
const QLocale__Denmark QLocale__Country = 58
const QLocale__Djibouti QLocale__Country = 59
const QLocale__Dominica QLocale__Country = 60
const QLocale__DominicanRepublic QLocale__Country = 61
const QLocale__EastTimor QLocale__Country = 62
const QLocale__Ecuador QLocale__Country = 63
const QLocale__Egypt QLocale__Country = 64
const QLocale__ElSalvador QLocale__Country = 65
const QLocale__EquatorialGuinea QLocale__Country = 66
const QLocale__Eritrea QLocale__Country = 67
const QLocale__Estonia QLocale__Country = 68
const QLocale__Ethiopia QLocale__Country = 69
const QLocale__FalklandIslands QLocale__Country = 70
const QLocale__FaroeIslands QLocale__Country = 71
const QLocale__Fiji QLocale__Country = 72
const QLocale__Finland QLocale__Country = 73
const QLocale__France QLocale__Country = 74
const QLocale__Guernsey QLocale__Country = 75
const QLocale__FrenchGuiana QLocale__Country = 76
const QLocale__FrenchPolynesia QLocale__Country = 77
const QLocale__FrenchSouthernTerritories QLocale__Country = 78
const QLocale__Gabon QLocale__Country = 79
const QLocale__Gambia QLocale__Country = 80
const QLocale__Georgia QLocale__Country = 81
const QLocale__Germany QLocale__Country = 82
const QLocale__Ghana QLocale__Country = 83
const QLocale__Gibraltar QLocale__Country = 84
const QLocale__Greece QLocale__Country = 85
const QLocale__Greenland QLocale__Country = 86
const QLocale__Grenada QLocale__Country = 87
const QLocale__Guadeloupe QLocale__Country = 88
const QLocale__Guam QLocale__Country = 89
const QLocale__Guatemala QLocale__Country = 90
const QLocale__Guinea QLocale__Country = 91
const QLocale__GuineaBissau QLocale__Country = 92
const QLocale__Guyana QLocale__Country = 93
const QLocale__Haiti QLocale__Country = 94
const QLocale__HeardAndMcDonaldIslands QLocale__Country = 95
const QLocale__Honduras QLocale__Country = 96
const QLocale__HongKong QLocale__Country = 97
const QLocale__Hungary QLocale__Country = 98
const QLocale__Iceland QLocale__Country = 99
const QLocale__India QLocale__Country = 100
const QLocale__Indonesia QLocale__Country = 101
const QLocale__Iran QLocale__Country = 102
const QLocale__Iraq QLocale__Country = 103
const QLocale__Ireland QLocale__Country = 104
const QLocale__Israel QLocale__Country = 105
const QLocale__Italy QLocale__Country = 106
const QLocale__Jamaica QLocale__Country = 107
const QLocale__Japan QLocale__Country = 108
const QLocale__Jordan QLocale__Country = 109
const QLocale__Kazakhstan QLocale__Country = 110
const QLocale__Kenya QLocale__Country = 111
const QLocale__Kiribati QLocale__Country = 112
const QLocale__NorthKorea QLocale__Country = 113
const QLocale__SouthKorea QLocale__Country = 114
const QLocale__Kuwait QLocale__Country = 115
const QLocale__Kyrgyzstan QLocale__Country = 116
const QLocale__Laos QLocale__Country = 117
const QLocale__Latvia QLocale__Country = 118
const QLocale__Lebanon QLocale__Country = 119
const QLocale__Lesotho QLocale__Country = 120
const QLocale__Liberia QLocale__Country = 121
const QLocale__Libya QLocale__Country = 122
const QLocale__Liechtenstein QLocale__Country = 123
const QLocale__Lithuania QLocale__Country = 124
const QLocale__Luxembourg QLocale__Country = 125
const QLocale__Macau QLocale__Country = 126
const QLocale__Macedonia QLocale__Country = 127
const QLocale__Madagascar QLocale__Country = 128
const QLocale__Malawi QLocale__Country = 129
const QLocale__Malaysia QLocale__Country = 130
const QLocale__Maldives QLocale__Country = 131
const QLocale__Mali QLocale__Country = 132
const QLocale__Malta QLocale__Country = 133
const QLocale__MarshallIslands QLocale__Country = 134
const QLocale__Martinique QLocale__Country = 135
const QLocale__Mauritania QLocale__Country = 136
const QLocale__Mauritius QLocale__Country = 137
const QLocale__Mayotte QLocale__Country = 138
const QLocale__Mexico QLocale__Country = 139
const QLocale__Micronesia QLocale__Country = 140
const QLocale__Moldova QLocale__Country = 141
const QLocale__Monaco QLocale__Country = 142
const QLocale__Mongolia QLocale__Country = 143
const QLocale__Montserrat QLocale__Country = 144
const QLocale__Morocco QLocale__Country = 145
const QLocale__Mozambique QLocale__Country = 146
const QLocale__Myanmar QLocale__Country = 147
const QLocale__Namibia QLocale__Country = 148
const QLocale__NauruCountry QLocale__Country = 149
const QLocale__Nepal QLocale__Country = 150
const QLocale__Netherlands QLocale__Country = 151
const QLocale__CuraSao QLocale__Country = 152
const QLocale__NewCaledonia QLocale__Country = 153
const QLocale__NewZealand QLocale__Country = 154
const QLocale__Nicaragua QLocale__Country = 155
const QLocale__Niger QLocale__Country = 156
const QLocale__Nigeria QLocale__Country = 157
const QLocale__Niue QLocale__Country = 158
const QLocale__NorfolkIsland QLocale__Country = 159
const QLocale__NorthernMarianaIslands QLocale__Country = 160
const QLocale__Norway QLocale__Country = 161
const QLocale__Oman QLocale__Country = 162
const QLocale__Pakistan QLocale__Country = 163
const QLocale__Palau QLocale__Country = 164
const QLocale__PalestinianTerritories QLocale__Country = 165
const QLocale__Panama QLocale__Country = 166
const QLocale__PapuaNewGuinea QLocale__Country = 167
const QLocale__Paraguay QLocale__Country = 168
const QLocale__Peru QLocale__Country = 169
const QLocale__Philippines QLocale__Country = 170
const QLocale__Pitcairn QLocale__Country = 171
const QLocale__Poland QLocale__Country = 172
const QLocale__Portugal QLocale__Country = 173
const QLocale__PuertoRico QLocale__Country = 174
const QLocale__Qatar QLocale__Country = 175
const QLocale__Reunion QLocale__Country = 176
const QLocale__Romania QLocale__Country = 177
const QLocale__Russia QLocale__Country = 178
const QLocale__Rwanda QLocale__Country = 179
const QLocale__SaintKittsAndNevis QLocale__Country = 180
const QLocale__SaintLucia QLocale__Country = 181
const QLocale__SaintVincentAndTheGrenadines QLocale__Country = 182
const QLocale__Samoa QLocale__Country = 183
const QLocale__SanMarino QLocale__Country = 184
const QLocale__SaoTomeAndPrincipe QLocale__Country = 185
const QLocale__SaudiArabia QLocale__Country = 186
const QLocale__Senegal QLocale__Country = 187
const QLocale__Seychelles QLocale__Country = 188
const QLocale__SierraLeone QLocale__Country = 189
const QLocale__Singapore QLocale__Country = 190
const QLocale__Slovakia QLocale__Country = 191
const QLocale__Slovenia QLocale__Country = 192
const QLocale__SolomonIslands QLocale__Country = 193
const QLocale__Somalia QLocale__Country = 194
const QLocale__SouthAfrica QLocale__Country = 195
const QLocale__SouthGeorgiaAndTheSouthSandwichIslands QLocale__Country = 196
const QLocale__Spain QLocale__Country = 197
const QLocale__SriLanka QLocale__Country = 198
const QLocale__SaintHelena QLocale__Country = 199
const QLocale__SaintPierreAndMiquelon QLocale__Country = 200
const QLocale__Sudan QLocale__Country = 201
const QLocale__Suriname QLocale__Country = 202
const QLocale__SvalbardAndJanMayenIslands QLocale__Country = 203
const QLocale__Swaziland QLocale__Country = 204
const QLocale__Sweden QLocale__Country = 205
const QLocale__Switzerland QLocale__Country = 206
const QLocale__Syria QLocale__Country = 207
const QLocale__Taiwan QLocale__Country = 208
const QLocale__Tajikistan QLocale__Country = 209
const QLocale__Tanzania QLocale__Country = 210
const QLocale__Thailand QLocale__Country = 211
const QLocale__Togo QLocale__Country = 212
const QLocale__TokelauCountry QLocale__Country = 213
const QLocale__Tonga QLocale__Country = 214
const QLocale__TrinidadAndTobago QLocale__Country = 215
const QLocale__Tunisia QLocale__Country = 216
const QLocale__Turkey QLocale__Country = 217
const QLocale__Turkmenistan QLocale__Country = 218
const QLocale__TurksAndCaicosIslands QLocale__Country = 219
const QLocale__TuvaluCountry QLocale__Country = 220
const QLocale__Uganda QLocale__Country = 221
const QLocale__Ukraine QLocale__Country = 222
const QLocale__UnitedArabEmirates QLocale__Country = 223
const QLocale__UnitedKingdom QLocale__Country = 224
const QLocale__UnitedStates QLocale__Country = 225
const QLocale__UnitedStatesMinorOutlyingIslands QLocale__Country = 226
const QLocale__Uruguay QLocale__Country = 227
const QLocale__Uzbekistan QLocale__Country = 228
const QLocale__Vanuatu QLocale__Country = 229
const QLocale__VaticanCityState QLocale__Country = 230
const QLocale__Venezuela QLocale__Country = 231
const QLocale__Vietnam QLocale__Country = 232
const QLocale__BritishVirginIslands QLocale__Country = 233
const QLocale__UnitedStatesVirginIslands QLocale__Country = 234
const QLocale__WallisAndFutunaIslands QLocale__Country = 235
const QLocale__WesternSahara QLocale__Country = 236
const QLocale__Yemen QLocale__Country = 237
const QLocale__CanaryIslands QLocale__Country = 238
const QLocale__Zambia QLocale__Country = 239
const QLocale__Zimbabwe QLocale__Country = 240
const QLocale__ClippertonIsland QLocale__Country = 241
const QLocale__Montenegro QLocale__Country = 242
const QLocale__Serbia QLocale__Country = 243
const QLocale__SaintBarthelemy QLocale__Country = 244
const QLocale__SaintMartin QLocale__Country = 245
const QLocale__LatinAmericaAndTheCaribbean QLocale__Country = 246
const QLocale__AscensionIsland QLocale__Country = 247
const QLocale__AlandIslands QLocale__Country = 248
const QLocale__DiegoGarcia QLocale__Country = 249
const QLocale__CeutaAndMelilla QLocale__Country = 250
const QLocale__IsleOfMan QLocale__Country = 251
const QLocale__Jersey QLocale__Country = 252
const QLocale__TristanDaCunha QLocale__Country = 253
const QLocale__SouthSudan QLocale__Country = 254
const QLocale__Bonaire QLocale__Country = 255
const QLocale__SintMaarten QLocale__Country = 256
const QLocale__Kosovo QLocale__Country = 257
const QLocale__EuropeanUnion QLocale__Country = 258
const QLocale__OutlyingOceania QLocale__Country = 259
const QLocale__Tokelau QLocale__Country = 213
const QLocale__Tuvalu QLocale__Country = 220
const QLocale__DemocraticRepublicOfCongo QLocale__Country = 49
const QLocale__PeoplesRepublicOfCongo QLocale__Country = 50
const QLocale__DemocraticRepublicOfKorea QLocale__Country = 113
const QLocale__RepublicOfKorea QLocale__Country = 114
const QLocale__RussianFederation QLocale__Country = 178
const QLocale__SyrianArabRepublic QLocale__Country = 207
const QLocale__LastCountry QLocale__Country = 259

type QLocale__MeasurementSystem = int

const QLocale__MetricSystem QLocale__MeasurementSystem = 0
const QLocale__ImperialUSSystem QLocale__MeasurementSystem = 1
const QLocale__ImperialUKSystem QLocale__MeasurementSystem = 2
const QLocale__ImperialSystem QLocale__MeasurementSystem = 1

type QLocale__FormatType = int

const QLocale__LongFormat QLocale__FormatType = 0
const QLocale__ShortFormat QLocale__FormatType = 1
const QLocale__NarrowFormat QLocale__FormatType = 2

type QLocale__NumberOption = int

const QLocale__DefaultNumberOptions QLocale__NumberOption = 0
const QLocale__OmitGroupSeparator QLocale__NumberOption = 1
const QLocale__RejectGroupSeparator QLocale__NumberOption = 2
const QLocale__OmitLeadingZeroInExponent QLocale__NumberOption = 4
const QLocale__RejectLeadingZeroInExponent QLocale__NumberOption = 8
const QLocale__IncludeTrailingZeroesAfterDot QLocale__NumberOption = 16
const QLocale__RejectTrailingZeroesAfterDot QLocale__NumberOption = 32

type QLocale__FloatingPointPrecisionOption = int

const QLocale__FloatingPointShortest QLocale__FloatingPointPrecisionOption = 4294967168

type QLocale__CurrencySymbolFormat = int

const QLocale__CurrencyIsoCode QLocale__CurrencySymbolFormat = 0
const QLocale__CurrencySymbol QLocale__CurrencySymbolFormat = 1
const QLocale__CurrencyDisplayName QLocale__CurrencySymbolFormat = 2

type QLocale__DataSizeFormat = int

const QLocale__DataSizeBase1000 QLocale__DataSizeFormat = 1
const QLocale__DataSizeSIQuantifiers QLocale__DataSizeFormat = 2
const QLocale__DataSizeIecFormat QLocale__DataSizeFormat = 0
const QLocale__DataSizeTraditionalFormat QLocale__DataSizeFormat = 2
const QLocale__DataSizeSIFormat QLocale__DataSizeFormat = 3

type QLocale__QuotationStyle = int

const QLocale__StandardQuotation QLocale__QuotationStyle = 0
const QLocale__AlternateQuotation QLocale__QuotationStyle = 1

//  body block end
