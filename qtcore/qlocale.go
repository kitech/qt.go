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

/*

 */
type QLocale struct {
	*qtrt.CObject
}
type QLocale_ITF interface {
	QLocale_PTR() *QLocale
}

func (ptr *QLocale) QLocale_PTR() *QLocale { return ptr }

func (this *QLocale) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLocale) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLocaleFromPointer(cthis unsafe.Pointer) *QLocale {
	return &QLocale{&qtrt.CObject{cthis}}
}
func (*QLocale) NewFromPointer(cthis unsafe.Pointer) *QLocale {
	return NewQLocaleFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlocale.h:929
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocale()

/*
Constructs a QLocale object initialized with the default locale. If no default locale was set using setDefault(), this locale will be the same as the one returned by system().

See also setDefault().
*/
func (*QLocale) NewForInherit() *QLocale {
	return NewQLocale()
}
func NewQLocale() *QLocale {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLocale)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:930
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLocale(const QString &)

/*
Constructs a QLocale object initialized with the default locale. If no default locale was set using setDefault(), this locale will be the same as the one returned by system().

See also setDefault().
*/
func (*QLocale) NewForInherit1(name string) *QLocale {
	return NewQLocale1(name)
}
func NewQLocale1(name string) *QLocale {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLocale)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:931
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLocale(QLocale::Language, QLocale::Country)

/*
Constructs a QLocale object initialized with the default locale. If no default locale was set using setDefault(), this locale will be the same as the one returned by system().

See also setDefault().
*/
func (*QLocale) NewForInherit2(language int, country int) *QLocale {
	return NewQLocale2(language, country)
}
func NewQLocale2(language int, country int) *QLocale {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_7CountryE", qtrt.FFI_TYPE_POINTER, language, country)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLocale)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:931
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLocale(QLocale::Language, QLocale::Country)

/*
Constructs a QLocale object initialized with the default locale. If no default locale was set using setDefault(), this locale will be the same as the one returned by system().

See also setDefault().
*/
func (*QLocale) NewForInherit2p(language int) *QLocale {
	return NewQLocale2p(language)
}
func NewQLocale2p(language int) *QLocale {
	// arg: 1, QLocale::Country=Enum, QLocale::Country=Enum, , Invalid
	country := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_7CountryE", qtrt.FFI_TYPE_POINTER, language, country)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLocale)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:932
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QLocale(QLocale::Language, QLocale::Script, QLocale::Country)

/*
Constructs a QLocale object initialized with the default locale. If no default locale was set using setDefault(), this locale will be the same as the one returned by system().

See also setDefault().
*/
func (*QLocale) NewForInherit3(language int, script int, country int) *QLocale {
	return NewQLocale3(language, script, country)
}
func NewQLocale3(language int, script int, country int) *QLocale {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleC2ENS_8LanguageENS_6ScriptENS_7CountryE", qtrt.FFI_TYPE_POINTER, language, script, country)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLocale)
	return gothis
}

// /usr/include/qt/QtCore/qlocale.h:935
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QLocale & operator=(QLocale &&)

/*

 */
func (this *QLocale) Operator_equal(other unsafe.Pointer /*333*/) *QLocale {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:937
// index:1
// Public Visibility=Default Availability=Available
// [8] QLocale & operator=(const QLocale &)

/*

 */
func (this *QLocale) Operator_equal1(other QLocale_ITF) *QLocale {
	var convArg0 unsafe.Pointer
	if other != nil && other.QLocale_PTR() != nil {
		convArg0 = other.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:938
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QLocale()

/*

 */
func DeleteQLocale(this *QLocale) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocaleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qlocale.h:940
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QLocale &)

/*
Swaps locale other with this locale. This operation is very fast and never fails.

This function was introduced in  Qt 5.6.
*/
func (this *QLocale) Swap(other QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QLocale_PTR() != nil {
		convArg0 = other.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:942
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::Language language() const

/*
Returns the language of this locale.

See also script(), country(), languageToString(), and bcp47Name().
*/
func (this *QLocale) Language() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8languageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:943
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::Script script() const

/*
Returns the script of this locale.

This function was introduced in  Qt 4.8.

See also language(), country(), languageToString(), scriptToString(), and bcp47Name().
*/
func (this *QLocale) Script() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6scriptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:944
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::Country country() const

/*
Returns the country of this locale.

See also language(), script(), countryToString(), and bcp47Name().
*/
func (this *QLocale) Country() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7countryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:945
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the language and country of this locale as a string of the form "language_country", where language is a lowercase, two-letter ISO 639 language code, and country is an uppercase, two- or three-letter ISO 3166 country code.

Note that even if QLocale object was constructed with an explicit script, name() will not contain it for compatibility reasons. Use bcp47Name() instead if you need a full locale name.

See also QLocale(), language(), script(), country(), and bcp47Name().
*/
func (this *QLocale) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:947
// index:0
// Public Visibility=Default Availability=Available
// [8] QString bcp47Name() const

/*
Returns the dash-separated language, script and country (and possibly other BCP47 fields) of this locale as a string.

Unlike the uiLanguages() the returned value of the bcp47Name() represents the locale name of the QLocale data but not the language the user-interface should be in.

This function tries to conform the locale name to BCP47.

This function was introduced in  Qt 4.8.

See also language(), country(), script(), and uiLanguages().
*/
func (this *QLocale) Bcp47Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale9bcp47NameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:948
// index:0
// Public Visibility=Default Availability=Available
// [8] QString nativeLanguageName() const

/*
Returns a native name of the language for the locale. For example "Schwiizertüütsch" for Swiss-German locale.

This function was introduced in  Qt 4.8.

See also nativeCountryName() and languageToString().
*/
func (this *QLocale) NativeLanguageName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale18nativeLanguageNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:949
// index:0
// Public Visibility=Default Availability=Available
// [8] QString nativeCountryName() const

/*
Returns a native name of the country for the locale. For example "España" for Spanish/Spain locale.

This function was introduced in  Qt 4.8.

See also nativeLanguageName() and countryToString().
*/
func (this *QLocale) NativeCountryName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale17nativeCountryNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:952
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(const QString &, bool *) const

/*
Returns the short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUShort() and toString().
*/
func (this *QLocale) ToShort(s string, ok *bool) int16 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toShortERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:952
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(const QString &, bool *) const

/*
Returns the short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUShort() and toString().
*/
func (this *QLocale) ToShortp(s string) int16 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toShortERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:961
// index:1
// Public Visibility=Default Availability=Available
// [2] short toShort(const QStringRef &, bool *) const

/*
Returns the short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUShort() and toString().
*/
func (this *QLocale) ToShort1(s QStringRef_ITF, ok *bool) int16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toShortERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:961
// index:1
// Public Visibility=Default Availability=Available
// [2] short toShort(const QStringRef &, bool *) const

/*
Returns the short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUShort() and toString().
*/
func (this *QLocale) ToShort1p(s QStringRef_ITF) int16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toShortERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:971
// index:2
// Public Visibility=Default Availability=Available
// [2] short toShort(QStringView, bool *) const

/*
Returns the short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUShort() and toString().
*/
func (this *QLocale) ToShort2(s QStringView_ITF /*123*/, ok *bool) int16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toShortE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:971
// index:2
// Public Visibility=Default Availability=Available
// [2] short toShort(QStringView, bool *) const

/*
Returns the short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUShort() and toString().
*/
func (this *QLocale) ToShort2p(s QStringView_ITF /*123*/) int16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toShortE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:953
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(const QString &, bool *) const

/*
Returns the unsigned short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toShort() and toString().
*/
func (this *QLocale) ToUShort(s string, ok *bool) uint16 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:953
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(const QString &, bool *) const

/*
Returns the unsigned short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toShort() and toString().
*/
func (this *QLocale) ToUShortp(s string) uint16 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:962
// index:1
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(const QStringRef &, bool *) const

/*
Returns the unsigned short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toShort() and toString().
*/
func (this *QLocale) ToUShort1(s QStringRef_ITF, ok *bool) uint16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:962
// index:1
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(const QStringRef &, bool *) const

/*
Returns the unsigned short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toShort() and toString().
*/
func (this *QLocale) ToUShort1p(s QStringRef_ITF) uint16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toUShortERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:972
// index:2
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(QStringView, bool *) const

/*
Returns the unsigned short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toShort() and toString().
*/
func (this *QLocale) ToUShort2(s QStringView_ITF /*123*/, ok *bool) uint16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toUShortE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:972
// index:2
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(QStringView, bool *) const

/*
Returns the unsigned short int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toShort() and toString().
*/
func (this *QLocale) ToUShort2p(s QStringView_ITF /*123*/) uint16 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toUShortE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:954
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(const QString &, bool *) const

/*
Returns the int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUInt() and toString().
*/
func (this *QLocale) ToInt(s string, ok *bool) int {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale5toIntERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:954
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(const QString &, bool *) const

/*
Returns the int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUInt() and toString().
*/
func (this *QLocale) ToIntp(s string) int {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale5toIntERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:963
// index:1
// Public Visibility=Default Availability=Available
// [4] int toInt(const QStringRef &, bool *) const

/*
Returns the int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUInt() and toString().
*/
func (this *QLocale) ToInt1(s QStringRef_ITF, ok *bool) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale5toIntERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:963
// index:1
// Public Visibility=Default Availability=Available
// [4] int toInt(const QStringRef &, bool *) const

/*
Returns the int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUInt() and toString().
*/
func (this *QLocale) ToInt1p(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale5toIntERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:973
// index:2
// Public Visibility=Default Availability=Available
// [4] int toInt(QStringView, bool *) const

/*
Returns the int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUInt() and toString().
*/
func (this *QLocale) ToInt2(s QStringView_ITF /*123*/, ok *bool) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale5toIntE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:973
// index:2
// Public Visibility=Default Availability=Available
// [4] int toInt(QStringView, bool *) const

/*
Returns the int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toUInt() and toString().
*/
func (this *QLocale) ToInt2p(s QStringView_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale5toIntE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:955
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(const QString &, bool *) const

/*
Returns the unsigned int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt() and toString().
*/
func (this *QLocale) ToUInt(s string, ok *bool) uint {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:955
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(const QString &, bool *) const

/*
Returns the unsigned int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt() and toString().
*/
func (this *QLocale) ToUIntp(s string) uint {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:964
// index:1
// Public Visibility=Default Availability=Available
// [4] uint toUInt(const QStringRef &, bool *) const

/*
Returns the unsigned int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt() and toString().
*/
func (this *QLocale) ToUInt1(s QStringRef_ITF, ok *bool) uint {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:964
// index:1
// Public Visibility=Default Availability=Available
// [4] uint toUInt(const QStringRef &, bool *) const

/*
Returns the unsigned int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt() and toString().
*/
func (this *QLocale) ToUInt1p(s QStringRef_ITF) uint {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toUIntERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:974
// index:2
// Public Visibility=Default Availability=Available
// [4] uint toUInt(QStringView, bool *) const

/*
Returns the unsigned int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt() and toString().
*/
func (this *QLocale) ToUInt2(s QStringView_ITF /*123*/, ok *bool) uint {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toUIntE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:974
// index:2
// Public Visibility=Default Availability=Available
// [4] uint toUInt(QStringView, bool *) const

/*
Returns the unsigned int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt() and toString().
*/
func (this *QLocale) ToUInt2p(s QStringView_ITF /*123*/) uint {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6toUIntE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:956
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(const QString &, bool *) const

/*
Returns the long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt(), toULongLong(), toDouble(), and toString().
*/
func (this *QLocale) ToLongLong(s string, ok *bool) int64 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:956
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(const QString &, bool *) const

/*
Returns the long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt(), toULongLong(), toDouble(), and toString().
*/
func (this *QLocale) ToLongLongp(s string) int64 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:965
// index:1
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(const QStringRef &, bool *) const

/*
Returns the long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt(), toULongLong(), toDouble(), and toString().
*/
func (this *QLocale) ToLongLong1(s QStringRef_ITF, ok *bool) int64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:965
// index:1
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(const QStringRef &, bool *) const

/*
Returns the long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt(), toULongLong(), toDouble(), and toString().
*/
func (this *QLocale) ToLongLong1p(s QStringRef_ITF) int64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toLongLongERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:975
// index:2
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(QStringView, bool *) const

/*
Returns the long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt(), toULongLong(), toDouble(), and toString().
*/
func (this *QLocale) ToLongLong2(s QStringView_ITF /*123*/, ok *bool) int64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toLongLongE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:975
// index:2
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(QStringView, bool *) const

/*
Returns the long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toInt(), toULongLong(), toDouble(), and toString().
*/
func (this *QLocale) ToLongLong2p(s QStringView_ITF /*123*/) int64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10toLongLongE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:957
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(const QString &, bool *) const

/*
Returns the unsigned long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toLongLong(), toInt(), toDouble(), and toString().
*/
func (this *QLocale) ToULongLong(s string, ok *bool) uint64 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:957
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(const QString &, bool *) const

/*
Returns the unsigned long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toLongLong(), toInt(), toDouble(), and toString().
*/
func (this *QLocale) ToULongLongp(s string) uint64 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK7QStringPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:966
// index:1
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(const QStringRef &, bool *) const

/*
Returns the unsigned long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toLongLong(), toInt(), toDouble(), and toString().
*/
func (this *QLocale) ToULongLong1(s QStringRef_ITF, ok *bool) uint64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:966
// index:1
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(const QStringRef &, bool *) const

/*
Returns the unsigned long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toLongLong(), toInt(), toDouble(), and toString().
*/
func (this *QLocale) ToULongLong1p(s QStringRef_ITF) uint64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11toULongLongERK10QStringRefPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:976
// index:2
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(QStringView, bool *) const

/*
Returns the unsigned long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toLongLong(), toInt(), toDouble(), and toString().
*/
func (this *QLocale) ToULongLong2(s QStringView_ITF /*123*/, ok *bool) uint64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11toULongLongE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:976
// index:2
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(QStringView, bool *) const

/*
Returns the unsigned long long int represented by the localized string s.

If the conversion fails the function returns 0.

If ok is not 0, failure is reported by setting *ok to false, and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toLongLong(), toInt(), toDouble(), and toString().
*/
func (this *QLocale) ToULongLong2p(s QStringView_ITF /*123*/) uint64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11toULongLongE11QStringViewPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:958
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(const QString &, bool *) const

/*
Returns the float represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toDouble(), toInt(), and toString().
*/
func (this *QLocale) ToFloat(s string, ok *bool) float32 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK7QStringPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:958
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(const QString &, bool *) const

/*
Returns the float represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toDouble(), toInt(), and toString().
*/
func (this *QLocale) ToFloatp(s string) float32 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK7QStringPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:967
// index:1
// Public Visibility=Default Availability=Available
// [4] float toFloat(const QStringRef &, bool *) const

/*
Returns the float represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toDouble(), toInt(), and toString().
*/
func (this *QLocale) ToFloat1(s QStringRef_ITF, ok *bool) float32 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK10QStringRefPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:967
// index:1
// Public Visibility=Default Availability=Available
// [4] float toFloat(const QStringRef &, bool *) const

/*
Returns the float represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toDouble(), toInt(), and toString().
*/
func (this *QLocale) ToFloat1p(s QStringRef_ITF) float32 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toFloatERK10QStringRefPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:977
// index:2
// Public Visibility=Default Availability=Available
// [4] float toFloat(QStringView, bool *) const

/*
Returns the float represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toDouble(), toInt(), and toString().
*/
func (this *QLocale) ToFloat2(s QStringView_ITF /*123*/, ok *bool) float32 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toFloatE11QStringViewPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:977
// index:2
// Public Visibility=Default Availability=Available
// [4] float toFloat(QStringView, bool *) const

/*
Returns the float represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

This function ignores leading and trailing whitespace.

See also toDouble(), toInt(), and toString().
*/
func (this *QLocale) ToFloat2p(s QStringView_ITF /*123*/) float32 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toFloatE11QStringViewPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:959
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(const QString &, bool *) const

/*
Returns the double represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

Unlike QString::toDouble(), this function does not use the 'C' locale if the string cannot be interpreted in this locale.


  bool ok;
  double d;

  QLocale c(QLocale::C);
  d = c.toDouble( "1234.56", &ok );  // ok == true, d == 1234.56
  d = c.toDouble( "1,234.56", &ok ); // ok == true, d == 1234.56
  d = c.toDouble( "1234,56", &ok );  // ok == false

  QLocale german(QLocale::German);
  d = german.toDouble( "1234,56", &ok );  // ok == true, d == 1234.56
  d = german.toDouble( "1.234,56", &ok ); // ok == true, d == 1234.56
  d = german.toDouble( "1234.56", &ok );  // ok == false

  d = german.toDouble( "1.234", &ok );    // ok == true, d == 1234.0



Notice that the last conversion returns 1234.0, because '.' is the thousands group separator in the German locale.

This function ignores leading and trailing whitespace.

See also toFloat(), toInt(), and toString().
*/
func (this *QLocale) ToDouble(s string, ok *bool) float64 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK7QStringPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:959
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(const QString &, bool *) const

/*
Returns the double represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

Unlike QString::toDouble(), this function does not use the 'C' locale if the string cannot be interpreted in this locale.


  bool ok;
  double d;

  QLocale c(QLocale::C);
  d = c.toDouble( "1234.56", &ok );  // ok == true, d == 1234.56
  d = c.toDouble( "1,234.56", &ok ); // ok == true, d == 1234.56
  d = c.toDouble( "1234,56", &ok );  // ok == false

  QLocale german(QLocale::German);
  d = german.toDouble( "1234,56", &ok );  // ok == true, d == 1234.56
  d = german.toDouble( "1.234,56", &ok ); // ok == true, d == 1234.56
  d = german.toDouble( "1234.56", &ok );  // ok == false

  d = german.toDouble( "1.234", &ok );    // ok == true, d == 1234.0



Notice that the last conversion returns 1234.0, because '.' is the thousands group separator in the German locale.

This function ignores leading and trailing whitespace.

See also toFloat(), toInt(), and toString().
*/
func (this *QLocale) ToDoublep(s string) float64 {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK7QStringPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:968
// index:1
// Public Visibility=Default Availability=Available
// [8] double toDouble(const QStringRef &, bool *) const

/*
Returns the double represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

Unlike QString::toDouble(), this function does not use the 'C' locale if the string cannot be interpreted in this locale.


  bool ok;
  double d;

  QLocale c(QLocale::C);
  d = c.toDouble( "1234.56", &ok );  // ok == true, d == 1234.56
  d = c.toDouble( "1,234.56", &ok ); // ok == true, d == 1234.56
  d = c.toDouble( "1234,56", &ok );  // ok == false

  QLocale german(QLocale::German);
  d = german.toDouble( "1234,56", &ok );  // ok == true, d == 1234.56
  d = german.toDouble( "1.234,56", &ok ); // ok == true, d == 1234.56
  d = german.toDouble( "1234.56", &ok );  // ok == false

  d = german.toDouble( "1.234", &ok );    // ok == true, d == 1234.0



Notice that the last conversion returns 1234.0, because '.' is the thousands group separator in the German locale.

This function ignores leading and trailing whitespace.

See also toFloat(), toInt(), and toString().
*/
func (this *QLocale) ToDouble1(s QStringRef_ITF, ok *bool) float64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK10QStringRefPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:968
// index:1
// Public Visibility=Default Availability=Available
// [8] double toDouble(const QStringRef &, bool *) const

/*
Returns the double represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

Unlike QString::toDouble(), this function does not use the 'C' locale if the string cannot be interpreted in this locale.


  bool ok;
  double d;

  QLocale c(QLocale::C);
  d = c.toDouble( "1234.56", &ok );  // ok == true, d == 1234.56
  d = c.toDouble( "1,234.56", &ok ); // ok == true, d == 1234.56
  d = c.toDouble( "1234,56", &ok );  // ok == false

  QLocale german(QLocale::German);
  d = german.toDouble( "1234,56", &ok );  // ok == true, d == 1234.56
  d = german.toDouble( "1.234,56", &ok ); // ok == true, d == 1234.56
  d = german.toDouble( "1234.56", &ok );  // ok == false

  d = german.toDouble( "1.234", &ok );    // ok == true, d == 1234.0



Notice that the last conversion returns 1234.0, because '.' is the thousands group separator in the German locale.

This function ignores leading and trailing whitespace.

See also toFloat(), toInt(), and toString().
*/
func (this *QLocale) ToDouble1p(s QStringRef_ITF) float64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toDoubleERK10QStringRefPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:978
// index:2
// Public Visibility=Default Availability=Available
// [8] double toDouble(QStringView, bool *) const

/*
Returns the double represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

Unlike QString::toDouble(), this function does not use the 'C' locale if the string cannot be interpreted in this locale.


  bool ok;
  double d;

  QLocale c(QLocale::C);
  d = c.toDouble( "1234.56", &ok );  // ok == true, d == 1234.56
  d = c.toDouble( "1,234.56", &ok ); // ok == true, d == 1234.56
  d = c.toDouble( "1234,56", &ok );  // ok == false

  QLocale german(QLocale::German);
  d = german.toDouble( "1234,56", &ok );  // ok == true, d == 1234.56
  d = german.toDouble( "1.234,56", &ok ); // ok == true, d == 1234.56
  d = german.toDouble( "1234.56", &ok );  // ok == false

  d = german.toDouble( "1.234", &ok );    // ok == true, d == 1234.0



Notice that the last conversion returns 1234.0, because '.' is the thousands group separator in the German locale.

This function ignores leading and trailing whitespace.

See also toFloat(), toInt(), and toString().
*/
func (this *QLocale) ToDouble2(s QStringView_ITF /*123*/, ok *bool) float64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toDoubleE11QStringViewPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:978
// index:2
// Public Visibility=Default Availability=Available
// [8] double toDouble(QStringView, bool *) const

/*
Returns the double represented by the localized string s, or 0.0 if the conversion failed.

If ok is not 0, reports failure by setting *ok to false and success by setting *ok to true.

Unlike QString::toDouble(), this function does not use the 'C' locale if the string cannot be interpreted in this locale.


  bool ok;
  double d;

  QLocale c(QLocale::C);
  d = c.toDouble( "1234.56", &ok );  // ok == true, d == 1234.56
  d = c.toDouble( "1,234.56", &ok ); // ok == true, d == 1234.56
  d = c.toDouble( "1234,56", &ok );  // ok == false

  QLocale german(QLocale::German);
  d = german.toDouble( "1234,56", &ok );  // ok == true, d == 1234.56
  d = german.toDouble( "1.234,56", &ok ); // ok == true, d == 1234.56
  d = german.toDouble( "1234.56", &ok );  // ok == false

  d = german.toDouble( "1.234", &ok );    // ok == true, d == 1234.0



Notice that the last conversion returns 1234.0, because '.' is the thousands group separator in the German locale.

This function ignores leading and trailing whitespace.

See also toFloat(), toInt(), and toString().
*/
func (this *QLocale) ToDouble2p(s QStringView_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toDoubleE11QStringViewPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlocale.h:980
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(qlonglong) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString(i int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:981
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(qulonglong) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString1(i uint64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:982
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QString toString(short) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString2(i int16) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:983
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QString toString(ushort) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString3(i uint16) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:984
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QString toString(int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString4(i int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:985
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QString toString(uint) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString5(i uint) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:986
// index:6
// Public Visibility=Default Availability=Available
// [8] QString toString(double, char, int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString6(i float64, f byte, prec int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:986
// index:6
// Public Visibility=Default Availability=Available
// [8] QString toString(double, char, int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString6p(i float64) string {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	f := 'g'
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:986
// index:6
// Public Visibility=Default Availability=Available
// [8] QString toString(double, char, int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString6p1(i float64, f byte) string {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:987
// index:7
// Public inline Visibility=Default Availability=Available
// [8] QString toString(float, char, int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString7(i float32, f byte, prec int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:987
// index:7
// Public inline Visibility=Default Availability=Available
// [8] QString toString(float, char, int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString7p(i float32) string {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	f := 'g'
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:987
// index:7
// Public inline Visibility=Default Availability=Available
// [8] QString toString(float, char, int) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString7p1(i float32, f byte) string {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:990
// index:8
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDate &, const QString &) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString8(date QDate_ITF, formatStr string) string {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(formatStr)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:991
// index:9
// Public Visibility=Default Availability=Available
// [8] QString toString(const QTime &, const QString &) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString9(time QTime_ITF, formatStr string) string {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(formatStr)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:992
// index:10
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDateTime &, const QString &) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString10(dateTime QDateTime_ITF, format string) string {
	var convArg0 unsafe.Pointer
	if dateTime != nil && dateTime.QDateTime_PTR() != nil {
		convArg0 = dateTime.QDateTime_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(format)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:994
// index:11
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDate &, QStringView) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString11(date QDate_ITF, formatStr QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if formatStr != nil && formatStr.QStringView_PTR() != nil {
		convArg1 = formatStr.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDate11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:995
// index:12
// Public Visibility=Default Availability=Available
// [8] QString toString(const QTime &, QStringView) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString12(time QTime_ITF, formatStr QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if formatStr != nil && formatStr.QStringView_PTR() != nil {
		convArg1 = formatStr.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTime11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:996
// index:13
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDateTime &, QStringView) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString13(dateTime QDateTime_ITF, format QStringView_ITF /*123*/) string {
	var convArg0 unsafe.Pointer
	if dateTime != nil && dateTime.QDateTime_PTR() != nil {
		convArg0 = dateTime.QDateTime_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QStringView_PTR() != nil {
		convArg1 = format.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTime11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:997
// index:14
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDate &, QLocale::FormatType) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString14(date QDate_ITF, format int) string {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:997
// index:14
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDate &, QLocale::FormatType) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString14p(date QDate_ITF) string {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QDateNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:998
// index:15
// Public Visibility=Default Availability=Available
// [8] QString toString(const QTime &, QLocale::FormatType) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString15(time QTime_ITF, format int) string {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:998
// index:15
// Public Visibility=Default Availability=Available
// [8] QString toString(const QTime &, QLocale::FormatType) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString15p(time QTime_ITF) string {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK5QTimeNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:999
// index:16
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDateTime &, QLocale::FormatType) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString16(dateTime QDateTime_ITF, format int) string {
	var convArg0 unsafe.Pointer
	if dateTime != nil && dateTime.QDateTime_PTR() != nil {
		convArg0 = dateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:999
// index:16
// Public Visibility=Default Availability=Available
// [8] QString toString(const QDateTime &, QLocale::FormatType) const

/*
Returns a localized string representation of i.

See also toLongLong().
*/
func (this *QLocale) ToString16p(dateTime QDateTime_ITF) string {
	var convArg0 unsafe.Pointer
	if dateTime != nil && dateTime.QDateTime_PTR() != nil {
		convArg0 = dateTime.QDateTime_PTR().GetCthis()
	}
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale8toStringERK9QDateTimeNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1001
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dateFormat(QLocale::FormatType) const

/*
Returns the date format used for the current locale.

If format is LongFormat the format will be a long version. Otherwise it uses a shorter version.

This function was introduced in  Qt 4.1.

See also QDate::toString() and QDate::fromString().
*/
func (this *QLocale) DateFormat(format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10dateFormatENS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1001
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dateFormat(QLocale::FormatType) const

/*
Returns the date format used for the current locale.

If format is LongFormat the format will be a long version. Otherwise it uses a shorter version.

This function was introduced in  Qt 4.1.

See also QDate::toString() and QDate::fromString().
*/
func (this *QLocale) DateFormatp() string {
	// arg: 0, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10dateFormatENS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1002
// index:0
// Public Visibility=Default Availability=Available
// [8] QString timeFormat(QLocale::FormatType) const

/*
Returns the time format used for the current locale.

If format is LongFormat the format will be a long version. Otherwise it uses a shorter version.

This function was introduced in  Qt 4.1.

See also QTime::toString() and QTime::fromString().
*/
func (this *QLocale) TimeFormat(format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10timeFormatENS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1002
// index:0
// Public Visibility=Default Availability=Available
// [8] QString timeFormat(QLocale::FormatType) const

/*
Returns the time format used for the current locale.

If format is LongFormat the format will be a long version. Otherwise it uses a shorter version.

This function was introduced in  Qt 4.1.

See also QTime::toString() and QTime::fromString().
*/
func (this *QLocale) TimeFormatp() string {
	// arg: 0, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale10timeFormatENS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1003
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dateTimeFormat(QLocale::FormatType) const

/*
Returns the date time format used for the current locale.

If format is ShortFormat the format will be a short version. Otherwise it uses a longer version.

This function was introduced in  Qt 4.4.

See also QDateTime::toString() and QDateTime::fromString().
*/
func (this *QLocale) DateTimeFormat(format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale14dateTimeFormatENS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1003
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dateTimeFormat(QLocale::FormatType) const

/*
Returns the date time format used for the current locale.

If format is ShortFormat the format will be a short version. Otherwise it uses a longer version.

This function was introduced in  Qt 4.4.

See also QDateTime::toString() and QDateTime::fromString().
*/
func (this *QLocale) DateTimeFormatp() string {
	// arg: 0, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale14dateTimeFormatENS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1005
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

// /usr/include/qt/QtCore/qlocale.h:1005
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

// /usr/include/qt/QtCore/qlocale.h:1008
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

// /usr/include/qt/QtCore/qlocale.h:1006
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

// /usr/include/qt/QtCore/qlocale.h:1006
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

// /usr/include/qt/QtCore/qlocale.h:1009
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

// /usr/include/qt/QtCore/qlocale.h:1007
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

// /usr/include/qt/QtCore/qlocale.h:1007
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

// /usr/include/qt/QtCore/qlocale.h:1010
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

// /usr/include/qt/QtCore/qlocale.h:1015
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar decimalPoint() const

/*
Returns the decimal point character of this locale.

This function was introduced in  Qt 4.1.
*/
func (this *QLocale) DecimalPoint() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale12decimalPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1016
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar groupSeparator() const

/*
Returns the group separator character of this locale.

This function was introduced in  Qt 4.1.
*/
func (this *QLocale) GroupSeparator() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale14groupSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1017
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar percent() const

/*
Returns the percent character of this locale.

This function was introduced in  Qt 4.1.
*/
func (this *QLocale) Percent() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7percentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1018
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar zeroDigit() const

/*
Returns the zero digit character of this locale.

This function was introduced in  Qt 4.1.
*/
func (this *QLocale) ZeroDigit() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale9zeroDigitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1019
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar negativeSign() const

/*
Returns the negative sign character of this locale.

This function was introduced in  Qt 4.1.
*/
func (this *QLocale) NegativeSign() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale12negativeSignEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1020
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar positiveSign() const

/*
Returns the positive sign character of this locale.

This function was introduced in  Qt 4.5.
*/
func (this *QLocale) PositiveSign() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale12positiveSignEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1021
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar exponential() const

/*
Returns the exponential character of this locale.

This function was introduced in  Qt 4.1.
*/
func (this *QLocale) Exponential() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11exponentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1023
// index:0
// Public Visibility=Default Availability=Available
// [8] QString monthName(int, QLocale::FormatType) const

/*
Returns the localized name of month, in the format specified by type.

This function was introduced in  Qt 4.2.

See also dayName() and standaloneMonthName().
*/
func (this *QLocale) MonthName(arg0 int, format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale9monthNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1023
// index:0
// Public Visibility=Default Availability=Available
// [8] QString monthName(int, QLocale::FormatType) const

/*
Returns the localized name of month, in the format specified by type.

This function was introduced in  Qt 4.2.

See also dayName() and standaloneMonthName().
*/
func (this *QLocale) MonthNamep(arg0 int) string {
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale9monthNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1024
// index:0
// Public Visibility=Default Availability=Available
// [8] QString standaloneMonthName(int, QLocale::FormatType) const

/*
Returns the localized name of month that is used as a standalone text, in the format specified by type.

If the locale information doesn't specify the standalone month name then return value is the same as in monthName().

This function was introduced in  Qt 4.5.

See also monthName() and standaloneDayName().
*/
func (this *QLocale) StandaloneMonthName(arg0 int, format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale19standaloneMonthNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1024
// index:0
// Public Visibility=Default Availability=Available
// [8] QString standaloneMonthName(int, QLocale::FormatType) const

/*
Returns the localized name of month that is used as a standalone text, in the format specified by type.

If the locale information doesn't specify the standalone month name then return value is the same as in monthName().

This function was introduced in  Qt 4.5.

See also monthName() and standaloneDayName().
*/
func (this *QLocale) StandaloneMonthNamep(arg0 int) string {
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale19standaloneMonthNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1025
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dayName(int, QLocale::FormatType) const

/*
Returns the localized name of the day (where 1 represents Monday, 2 represents Tuesday and so on), in the format specified by type.

This function was introduced in  Qt 4.2.

See also monthName() and standaloneDayName().
*/
func (this *QLocale) DayName(arg0 int, format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7dayNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1025
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dayName(int, QLocale::FormatType) const

/*
Returns the localized name of the day (where 1 represents Monday, 2 represents Tuesday and so on), in the format specified by type.

This function was introduced in  Qt 4.2.

See also monthName() and standaloneDayName().
*/
func (this *QLocale) DayNamep(arg0 int) string {
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7dayNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1026
// index:0
// Public Visibility=Default Availability=Available
// [8] QString standaloneDayName(int, QLocale::FormatType) const

/*
Returns the localized name of the day (where 1 represents Monday, 2 represents Tuesday and so on) that is used as a standalone text, in the format specified by type.

If the locale information does not specify the standalone day name then return value is the same as in dayName().

This function was introduced in  Qt 4.5.

See also dayName() and standaloneMonthName().
*/
func (this *QLocale) StandaloneDayName(arg0 int, format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale17standaloneDayNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1026
// index:0
// Public Visibility=Default Availability=Available
// [8] QString standaloneDayName(int, QLocale::FormatType) const

/*
Returns the localized name of the day (where 1 represents Monday, 2 represents Tuesday and so on) that is used as a standalone text, in the format specified by type.

If the locale information does not specify the standalone day name then return value is the same as in dayName().

This function was introduced in  Qt 4.5.

See also dayName() and standaloneMonthName().
*/
func (this *QLocale) StandaloneDayNamep(arg0 int) string {
	// arg: 1, QLocale::FormatType=Enum, QLocale::FormatType=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale17standaloneDayNameEiNS_10FormatTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1028
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DayOfWeek firstDayOfWeek() const

/*
Returns the first day of the week according to the current locale.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) FirstDayOfWeek() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale14firstDayOfWeekEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1031
// index:0
// Public Visibility=Default Availability=Available
// [8] QString amText() const

/*
Returns the localized name of the "AM" suffix for times specified using the conventions of the 12-hour clock.

This function was introduced in  Qt 4.5.

See also pmText().
*/
func (this *QLocale) AmText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6amTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1032
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pmText() const

/*
Returns the localized name of the "PM" suffix for times specified using the conventions of the 12-hour clock.

This function was introduced in  Qt 4.5.

See also amText().
*/
func (this *QLocale) PmText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale6pmTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1034
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::MeasurementSystem measurementSystem() const

/*
Returns the measurement system for the locale.

This function was introduced in  Qt 4.4.
*/
func (this *QLocale) MeasurementSystem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale17measurementSystemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1036
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection textDirection() const

/*
Returns the text direction of the language.

This function was introduced in  Qt 4.7.
*/
func (this *QLocale) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1038
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toUpper(const QString &) const

/*
Returns an uppercase copy of str.

If Qt Core is using the ICU libraries, they will be used to perform the transformation according to the rules of the current locale. Otherwise the conversion may be done in a platform-dependent manner, with QString::toUpper() as a generic fallback.

This function was introduced in  Qt 4.8.

See also QString::toUpper().
*/
func (this *QLocale) ToUpper(str string) string {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toUpperERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1039
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toLower(const QString &) const

/*
Returns a lowercase copy of str.

If Qt Core is using the ICU libraries, they will be used to perform the transformation according to the rules of the current locale. Otherwise the conversion may be done in a platform-dependent manner, with QString::toLower() as a generic fallback.

This function was introduced in  Qt 4.8.

See also QString::toLower().
*/
func (this *QLocale) ToLower(str string) string {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale7toLowerERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1041
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currencySymbol(QLocale::CurrencySymbolFormat) const

/*
Returns a currency symbol according to the format.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) CurrencySymbol(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale14currencySymbolENS_20CurrencySymbolFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1041
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currencySymbol(QLocale::CurrencySymbolFormat) const

/*
Returns a currency symbol according to the format.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) CurrencySymbolp() string {
	// arg: 0, QLocale::CurrencySymbolFormat=Enum, QLocale::CurrencySymbolFormat=Enum, , Invalid
	arg0 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale14currencySymbolENS_20CurrencySymbolFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1042
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(qlonglong, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString(arg0 int64, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringExRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1042
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(qlonglong, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyStringp(arg0 int64) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringExRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1043
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(qulonglong, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString1(arg0 uint64, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEyRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1043
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(qulonglong, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString1p(arg0 uint64) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEyRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1044
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(short, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString2(arg0 int16, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEsRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1044
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(short, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString2p(arg0 int16) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEsRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1045
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(ushort, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString3(arg0 uint16, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEtRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1045
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(ushort, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString3p(arg0 uint16) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEtRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1046
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(int, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString4(arg0 int, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1046
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(int, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString4p(arg0 int) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1047
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(uint, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString5(arg0 uint, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEjRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1047
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(uint, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString5p(arg0 uint) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEjRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1053
// index:6
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(double, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString6(arg0 float64, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1053
// index:6
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(double, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString6p(arg0 float64) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1054
// index:7
// Public Visibility=Default Availability=Available
// [8] QString toCurrencyString(double, const QString &, int) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString7(arg0 float64, symbol string, precision int) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEdRK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1, precision)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1055
// index:8
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(float, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString8(i float32, symbol string) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1055
// index:8
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(float, const QString &) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString8p(i float32) string {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1057
// index:9
// Public inline Visibility=Default Availability=Available
// [8] QString toCurrencyString(float, const QString &, int) const

/*
Returns a localized string representation of value as a currency. If the symbol is provided it is used instead of the default currency symbol.

This function was introduced in  Qt 4.8.

See also currencySymbol().
*/
func (this *QLocale) ToCurrencyString9(i float32, symbol string, precision int) string {
	var tmpArg1 = NewQString5(symbol)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale16toCurrencyStringEfRK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, convArg1, precision)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1061
// index:0
// Public Visibility=Default Availability=Available
// [8] QString formattedDataSize(qint64, int, QLocale::DataSizeFormats)

/*
Converts a size in bytes to a human-readable localized string, comprising a number and a quantified unit. The quantifier is chosen such that the number is at least one, and as small as possible. For example if bytes is 16384, precision is 2, and format is DataSizeIecFormat (the default), this function returns "16.00 KiB"; for 1330409069609 bytes it returns "1.21 GiB"; and so on. If format is DataSizeIecFormat or DataSizeTraditionalFormat, the given number of bytes is divided by a power of 1024, with result less than 1024; for DataSizeSIFormat, it is divided by a power of 1000, with result less than 1000. DataSizeIecFormat uses the new IEC standard quantifiers Ki, Mi and so on, whereas DataSizeSIFormat uses the older SI quantifiers k, M, etc., and DataSizeTraditionalFormat abuses them.

This function was introduced in  Qt 5.10.
*/
func (this *QLocale) FormattedDataSize(bytes int64, precision int, format int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale17formattedDataSizeExi6QFlagsINS_14DataSizeFormatEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytes, precision, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1061
// index:0
// Public Visibility=Default Availability=Available
// [8] QString formattedDataSize(qint64, int, QLocale::DataSizeFormats)

/*
Converts a size in bytes to a human-readable localized string, comprising a number and a quantified unit. The quantifier is chosen such that the number is at least one, and as small as possible. For example if bytes is 16384, precision is 2, and format is DataSizeIecFormat (the default), this function returns "16.00 KiB"; for 1330409069609 bytes it returns "1.21 GiB"; and so on. If format is DataSizeIecFormat or DataSizeTraditionalFormat, the given number of bytes is divided by a power of 1024, with result less than 1024; for DataSizeSIFormat, it is divided by a power of 1000, with result less than 1000. DataSizeIecFormat uses the new IEC standard quantifiers Ki, Mi and so on, whereas DataSizeSIFormat uses the older SI quantifiers k, M, etc., and DataSizeTraditionalFormat abuses them.

This function was introduced in  Qt 5.10.
*/
func (this *QLocale) FormattedDataSizep(bytes int64) string {
	// arg: 1, int=Int, =Invalid, , Invalid
	precision := int(2)
	// arg: 2, QLocale::DataSizeFormats=Typedef, QLocale::DataSizeFormats=Typedef, QFlags<QLocale::DataSizeFormat>, Unexposed
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale17formattedDataSizeExi6QFlagsINS_14DataSizeFormatEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytes, precision, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1061
// index:0
// Public Visibility=Default Availability=Available
// [8] QString formattedDataSize(qint64, int, QLocale::DataSizeFormats)

/*
Converts a size in bytes to a human-readable localized string, comprising a number and a quantified unit. The quantifier is chosen such that the number is at least one, and as small as possible. For example if bytes is 16384, precision is 2, and format is DataSizeIecFormat (the default), this function returns "16.00 KiB"; for 1330409069609 bytes it returns "1.21 GiB"; and so on. If format is DataSizeIecFormat or DataSizeTraditionalFormat, the given number of bytes is divided by a power of 1024, with result less than 1024; for DataSizeSIFormat, it is divided by a power of 1000, with result less than 1000. DataSizeIecFormat uses the new IEC standard quantifiers Ki, Mi and so on, whereas DataSizeSIFormat uses the older SI quantifiers k, M, etc., and DataSizeTraditionalFormat abuses them.

This function was introduced in  Qt 5.10.
*/
func (this *QLocale) FormattedDataSizep1(bytes int64, precision int) string {
	// arg: 2, QLocale::DataSizeFormats=Typedef, QLocale::DataSizeFormats=Typedef, QFlags<QLocale::DataSizeFormat>, Unexposed
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale17formattedDataSizeExi6QFlagsINS_14DataSizeFormatEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytes, precision, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1063
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList uiLanguages() const

/*
Returns an ordered list of locale names for translation purposes in preference order (like "en-Latn-US", "en-US", "en").

The return value represents locale names that the user expects to see the UI translation in.

Most like you do not need to use this function directly, but just pass the QLocale object to the QTranslator::load() function.

The first item in the list is the most preferred one.

This function was introduced in  Qt 4.8.

See also QTranslator and bcp47Name().
*/
func (this *QLocale) UiLanguages() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11uiLanguagesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qlocale.h:1065
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QLocale &) const

/*

 */
func (this *QLocale) Operator_equal_equal(other QLocale_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QLocale_PTR() != nil {
		convArg0 = other.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocaleeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlocale.h:1066
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QLocale &) const

/*

 */
func (this *QLocale) Operator_not_equal(other QLocale_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QLocale_PTR() != nil {
		convArg0 = other.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocaleneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlocale.h:1068
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString languageToString(QLocale::Language)

/*
Returns a QString containing the name of language.

See also countryToString(), scriptToString(), and bcp47Name().
*/
func (this *QLocale) LanguageToString(language int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale16languageToStringENS_8LanguageE", qtrt.FFI_TYPE_POINTER, language)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QLocale_LanguageToString(language int) string {
	var nilthis *QLocale
	rv := nilthis.LanguageToString(language)
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1069
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString countryToString(QLocale::Country)

/*
Returns a QString containing the name of country.

See also languageToString(), scriptToString(), country(), and bcp47Name().
*/
func (this *QLocale) CountryToString(country int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale15countryToStringENS_7CountryE", qtrt.FFI_TYPE_POINTER, country)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QLocale_CountryToString(country int) string {
	var nilthis *QLocale
	rv := nilthis.CountryToString(country)
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1070
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString scriptToString(QLocale::Script)

/*
Returns a QString containing the name of script.

This function was introduced in  Qt 4.8.

See also languageToString(), countryToString(), script(), and bcp47Name().
*/
func (this *QLocale) ScriptToString(script int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale14scriptToStringENS_6ScriptE", qtrt.FFI_TYPE_POINTER, script)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QLocale_ScriptToString(script int) string {
	var nilthis *QLocale
	rv := nilthis.ScriptToString(script)
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1071
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefault(const QLocale &)

/*
Sets the global default locale to locale. These values are used when a QLocale object is constructed with no arguments. If this function is not called, the system's locale is used.

Warning: In a multithreaded application, the default locale should be set at application startup, before any non-GUI threads are created.

Warning: This function is not reentrant.

See also system() and c().
*/
func (this *QLocale) SetDefault(locale QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale10setDefaultERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QLocale_SetDefault(locale QLocale_ITF) {
	var nilthis *QLocale
	nilthis.SetDefault(locale)
}

// /usr/include/qt/QtCore/qlocale.h:1073
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QLocale c()

/*
Returns a QLocale object initialized to the "C" locale.

See also system().
*/
func (this *QLocale) C() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale1cEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}
func QLocale_C() *QLocale /*123*/ {
	var nilthis *QLocale
	rv := nilthis.C()
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1074
// index:0
// Public static Visibility=Default Availability=Available
// [8] QLocale system()

/*
Returns a QLocale object initialized to the system locale.

On Windows and Mac, this locale will use the decimal/grouping characters and date/time formats specified in the system configuration panel.

See also c().
*/
func (this *QLocale) System() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale6systemEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}
func QLocale_System() *QLocale /*123*/ {
	var nilthis *QLocale
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qlocale.h:1079
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumberOptions(QLocale::NumberOptions)

/*
Sets the options related to number conversions for this QLocale instance.

This function was introduced in  Qt 4.2.

See also numberOptions().
*/
func (this *QLocale) SetNumberOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLocale16setNumberOptionsE6QFlagsINS_12NumberOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlocale.h:1080
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocale::NumberOptions numberOptions() const

/*
Returns the options related to number conversions for this QLocale instance.

By default, no options are set for the standard locales.

This function was introduced in  Qt 4.2.

See also setNumberOptions().
*/
func (this *QLocale) NumberOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale13numberOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qlocale.h:1083
// index:0
// Public Visibility=Default Availability=Available
// [8] QString quoteString(const QString &, QLocale::QuotationStyle) const

/*
Returns str quoted according to the current locale using the given quotation style.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) QuoteString(str string, style int) string {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK7QStringNS_14QuotationStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1083
// index:0
// Public Visibility=Default Availability=Available
// [8] QString quoteString(const QString &, QLocale::QuotationStyle) const

/*
Returns str quoted according to the current locale using the given quotation style.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) QuoteStringp(str string) string {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QLocale::QuotationStyle=Enum, QLocale::QuotationStyle=Enum, , Invalid
	style := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK7QStringNS_14QuotationStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1084
// index:1
// Public Visibility=Default Availability=Available
// [8] QString quoteString(const QStringRef &, QLocale::QuotationStyle) const

/*
Returns str quoted according to the current locale using the given quotation style.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) QuoteString1(str QStringRef_ITF, style int) string {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK10QStringRefNS_14QuotationStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1084
// index:1
// Public Visibility=Default Availability=Available
// [8] QString quoteString(const QStringRef &, QLocale::QuotationStyle) const

/*
Returns str quoted according to the current locale using the given quotation style.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) QuoteString1p(str QStringRef_ITF) string {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, QLocale::QuotationStyle=Enum, QLocale::QuotationStyle=Enum, , Invalid
	style := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale11quoteStringERK10QStringRefNS_14QuotationStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, style)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlocale.h:1086
// index:0
// Public Visibility=Default Availability=Available
// [8] QString createSeparatedList(const QStringList &) const

/*
Returns a string that represents a join of a given list of strings with a separator defined by the locale.

This function was introduced in  Qt 4.8.
*/
func (this *QLocale) CreateSeparatedList(strl QStringList_ITF) string {
	var convArg0 unsafe.Pointer
	if strl != nil && strl.QStringList_PTR() != nil {
		convArg0 = strl.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLocale19createSeparatedListERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*
This enumerated type is used to specify a language.

QLocale::AfanOromoObsolete, please use Oromo
QLocale::BhutaniDzongkhaObsolete, please use Dzongkha
QLocale::ByelorussianBelarusianObsolete, please use Belarusian
QLocale::CambodianKhmerObsolete, please use Khmer
QLocale::FrisianWesternFrisiansame as WesternFrisian
QLocale::KurundiRundiObsolete, please use Rundi
QLocale::MoldavianRomanianObsolete, please use Romanian
QLocale::NorwegianNorwegianBokmalsame as NorwegianBokmal
QLocale::RhaetoRomanceRomanshObsolete, please use Romansh
QLocale::SerboCroatianSerbianObsolete, please use Serbian
QLocale::TagalogFilipinoObsolete, please use Filipino
QLocale::TwiAkanObsolete, please use Akan
QLocale::UigurUighurObsolete, please use Uighur
QLocale::ChewaNyanjaObsolete, please use Nyanja


See also language() and languageToString().

*/
type QLocale__Language = int

//
const QLocale__AnyLanguage QLocale__Language = 0

// The "C" locale is identical in behavior to English/UnitedStates.
const QLocale__C QLocale__Language = 1

//
const QLocale__Abkhazian QLocale__Language = 2

//
const QLocale__Oromo QLocale__Language = 3

//
const QLocale__Afar QLocale__Language = 4

//
const QLocale__Afrikaans QLocale__Language = 5

//
const QLocale__Albanian QLocale__Language = 6

//
const QLocale__Amharic QLocale__Language = 7

//
const QLocale__Arabic QLocale__Language = 8

//
const QLocale__Armenian QLocale__Language = 9

//
const QLocale__Assamese QLocale__Language = 10

//
const QLocale__Aymara QLocale__Language = 11

//
const QLocale__Azerbaijani QLocale__Language = 12

//
const QLocale__Bashkir QLocale__Language = 13

//
const QLocale__Basque QLocale__Language = 14

//
const QLocale__Bengali QLocale__Language = 15

//
const QLocale__Dzongkha QLocale__Language = 16

//
const QLocale__Bihari QLocale__Language = 17

//
const QLocale__Bislama QLocale__Language = 18

//
const QLocale__Breton QLocale__Language = 19

//
const QLocale__Bulgarian QLocale__Language = 20

//
const QLocale__Burmese QLocale__Language = 21

//
const QLocale__Belarusian QLocale__Language = 22

//
const QLocale__Khmer QLocale__Language = 23

//
const QLocale__Catalan QLocale__Language = 24

//
const QLocale__Chinese QLocale__Language = 25

//
const QLocale__Corsican QLocale__Language = 26

//
const QLocale__Croatian QLocale__Language = 27

//
const QLocale__Czech QLocale__Language = 28

//
const QLocale__Danish QLocale__Language = 29

//
const QLocale__Dutch QLocale__Language = 30

//
const QLocale__English QLocale__Language = 31

//
const QLocale__Esperanto QLocale__Language = 32

//
const QLocale__Estonian QLocale__Language = 33

//
const QLocale__Faroese QLocale__Language = 34

//
const QLocale__Fijian QLocale__Language = 35

//
const QLocale__Finnish QLocale__Language = 36

//
const QLocale__French QLocale__Language = 37

//
const QLocale__WesternFrisian QLocale__Language = 38

//
const QLocale__Gaelic QLocale__Language = 39

//
const QLocale__Galician QLocale__Language = 40

//
const QLocale__Georgian QLocale__Language = 41

//
const QLocale__German QLocale__Language = 42

//
const QLocale__Greek QLocale__Language = 43

//
const QLocale__Greenlandic QLocale__Language = 44

//
const QLocale__Guarani QLocale__Language = 45

//
const QLocale__Gujarati QLocale__Language = 46

//
const QLocale__Hausa QLocale__Language = 47

//
const QLocale__Hebrew QLocale__Language = 48

//
const QLocale__Hindi QLocale__Language = 49

//
const QLocale__Hungarian QLocale__Language = 50

//
const QLocale__Icelandic QLocale__Language = 51

//
const QLocale__Indonesian QLocale__Language = 52

//
const QLocale__Interlingua QLocale__Language = 53

//
const QLocale__Interlingue QLocale__Language = 54

//
const QLocale__Inuktitut QLocale__Language = 55

//
const QLocale__Inupiak QLocale__Language = 56

//
const QLocale__Irish QLocale__Language = 57

//
const QLocale__Italian QLocale__Language = 58

//
const QLocale__Japanese QLocale__Language = 59

//
const QLocale__Javanese QLocale__Language = 60

//
const QLocale__Kannada QLocale__Language = 61

//
const QLocale__Kashmiri QLocale__Language = 62

//
const QLocale__Kazakh QLocale__Language = 63

//
const QLocale__Kinyarwanda QLocale__Language = 64

//
const QLocale__Kirghiz QLocale__Language = 65

//
const QLocale__Korean QLocale__Language = 66

//
const QLocale__Kurdish QLocale__Language = 67

//
const QLocale__Rundi QLocale__Language = 68

//
const QLocale__Lao QLocale__Language = 69

//
const QLocale__Latin QLocale__Language = 70

//
const QLocale__Latvian QLocale__Language = 71

//
const QLocale__Lingala QLocale__Language = 72

//
const QLocale__Lithuanian QLocale__Language = 73

//
const QLocale__Macedonian QLocale__Language = 74

//
const QLocale__Malagasy QLocale__Language = 75

//
const QLocale__Malay QLocale__Language = 76

//
const QLocale__Malayalam QLocale__Language = 77

//
const QLocale__Maltese QLocale__Language = 78

//
const QLocale__Maori QLocale__Language = 79

//
const QLocale__Marathi QLocale__Language = 80

//
const QLocale__Marshallese QLocale__Language = 81

//
const QLocale__Mongolian QLocale__Language = 82

//
const QLocale__NauruLanguage QLocale__Language = 83

//
const QLocale__Nepali QLocale__Language = 84

//
const QLocale__NorwegianBokmal QLocale__Language = 85

//
const QLocale__Occitan QLocale__Language = 86

//
const QLocale__Oriya QLocale__Language = 87

//
const QLocale__Pashto QLocale__Language = 88

//
const QLocale__Persian QLocale__Language = 89

//
const QLocale__Polish QLocale__Language = 90

//
const QLocale__Portuguese QLocale__Language = 91

//
const QLocale__Punjabi QLocale__Language = 92

//
const QLocale__Quechua QLocale__Language = 93

//
const QLocale__Romansh QLocale__Language = 94

//
const QLocale__Romanian QLocale__Language = 95

//
const QLocale__Russian QLocale__Language = 96

//
const QLocale__Samoan QLocale__Language = 97

//
const QLocale__Sango QLocale__Language = 98

//
const QLocale__Sanskrit QLocale__Language = 99

//
const QLocale__Serbian QLocale__Language = 100

//
const QLocale__Ossetic QLocale__Language = 101

//
const QLocale__SouthernSotho QLocale__Language = 102

//
const QLocale__Tswana QLocale__Language = 103

//
const QLocale__Shona QLocale__Language = 104

//
const QLocale__Sindhi QLocale__Language = 105

//
const QLocale__Sinhala QLocale__Language = 106

//
const QLocale__Swati QLocale__Language = 107

//
const QLocale__Slovak QLocale__Language = 108

//
const QLocale__Slovenian QLocale__Language = 109

//
const QLocale__Somali QLocale__Language = 110

//
const QLocale__Spanish QLocale__Language = 111

//
const QLocale__Sundanese QLocale__Language = 112

//
const QLocale__Swahili QLocale__Language = 113

//
const QLocale__Swedish QLocale__Language = 114

//
const QLocale__Sardinian QLocale__Language = 115

//
const QLocale__Tajik QLocale__Language = 116

//
const QLocale__Tamil QLocale__Language = 117

//
const QLocale__Tatar QLocale__Language = 118

//
const QLocale__Telugu QLocale__Language = 119

//
const QLocale__Thai QLocale__Language = 120

//
const QLocale__Tibetan QLocale__Language = 121

//
const QLocale__Tigrinya QLocale__Language = 122

//
const QLocale__Tongan QLocale__Language = 123

//
const QLocale__Tsonga QLocale__Language = 124

//
const QLocale__Turkish QLocale__Language = 125

//
const QLocale__Turkmen QLocale__Language = 126

//
const QLocale__Tahitian QLocale__Language = 127

//
const QLocale__Uighur QLocale__Language = 128

//
const QLocale__Ukrainian QLocale__Language = 129

//
const QLocale__Urdu QLocale__Language = 130

//
const QLocale__Uzbek QLocale__Language = 131

//
const QLocale__Vietnamese QLocale__Language = 132

//
const QLocale__Volapuk QLocale__Language = 133

//
const QLocale__Welsh QLocale__Language = 134

//
const QLocale__Wolof QLocale__Language = 135

//
const QLocale__Xhosa QLocale__Language = 136

//
const QLocale__Yiddish QLocale__Language = 137

//
const QLocale__Yoruba QLocale__Language = 138

//
const QLocale__Zhuang QLocale__Language = 139

//
const QLocale__Zulu QLocale__Language = 140

//
const QLocale__NorwegianNynorsk QLocale__Language = 141

//
const QLocale__Bosnian QLocale__Language = 142

//
const QLocale__Divehi QLocale__Language = 143

//
const QLocale__Manx QLocale__Language = 144

//
const QLocale__Cornish QLocale__Language = 145

//
const QLocale__Akan QLocale__Language = 146

//
const QLocale__Konkani QLocale__Language = 147

//
const QLocale__Ga QLocale__Language = 148

//
const QLocale__Igbo QLocale__Language = 149

//
const QLocale__Kamba QLocale__Language = 150

//
const QLocale__Syriac QLocale__Language = 151

//
const QLocale__Blin QLocale__Language = 152

//
const QLocale__Geez QLocale__Language = 153

//
const QLocale__Koro QLocale__Language = 154

//
const QLocale__Sidamo QLocale__Language = 155

//
const QLocale__Atsam QLocale__Language = 156

//
const QLocale__Tigre QLocale__Language = 157

//
const QLocale__Jju QLocale__Language = 158

//
const QLocale__Friulian QLocale__Language = 159

//
const QLocale__Venda QLocale__Language = 160

//
const QLocale__Ewe QLocale__Language = 161

//
const QLocale__Walamo QLocale__Language = 162

//
const QLocale__Hawaiian QLocale__Language = 163

//
const QLocale__Tyap QLocale__Language = 164

//
const QLocale__Nyanja QLocale__Language = 165

//
const QLocale__Filipino QLocale__Language = 166

//
const QLocale__SwissGerman QLocale__Language = 167

//
const QLocale__SichuanYi QLocale__Language = 168

//
const QLocale__Kpelle QLocale__Language = 169

//
const QLocale__LowGerman QLocale__Language = 170

//
const QLocale__SouthNdebele QLocale__Language = 171

//
const QLocale__NorthernSotho QLocale__Language = 172

//
const QLocale__NorthernSami QLocale__Language = 173

//
const QLocale__Taroko QLocale__Language = 174

//
const QLocale__Gusii QLocale__Language = 175

//
const QLocale__Taita QLocale__Language = 176

//
const QLocale__Fulah QLocale__Language = 177

//
const QLocale__Kikuyu QLocale__Language = 178

//
const QLocale__Samburu QLocale__Language = 179

//
const QLocale__Sena QLocale__Language = 180

//
const QLocale__NorthNdebele QLocale__Language = 181

//
const QLocale__Rombo QLocale__Language = 182

//
const QLocale__Tachelhit QLocale__Language = 183

//
const QLocale__Kabyle QLocale__Language = 184

//
const QLocale__Nyankole QLocale__Language = 185

//
const QLocale__Bena QLocale__Language = 186

//
const QLocale__Vunjo QLocale__Language = 187

//
const QLocale__Bambara QLocale__Language = 188

//
const QLocale__Embu QLocale__Language = 189

//
const QLocale__Cherokee QLocale__Language = 190

//
const QLocale__Morisyen QLocale__Language = 191

//
const QLocale__Makonde QLocale__Language = 192

//
const QLocale__Langi QLocale__Language = 193

//
const QLocale__Ganda QLocale__Language = 194

//
const QLocale__Bemba QLocale__Language = 195

//
const QLocale__Kabuverdianu QLocale__Language = 196

//
const QLocale__Meru QLocale__Language = 197

//
const QLocale__Kalenjin QLocale__Language = 198

//
const QLocale__Nama QLocale__Language = 199

//
const QLocale__Machame QLocale__Language = 200

//
const QLocale__Colognian QLocale__Language = 201

//
const QLocale__Masai QLocale__Language = 202

//
const QLocale__Soga QLocale__Language = 203

//
const QLocale__Luyia QLocale__Language = 204

//
const QLocale__Asu QLocale__Language = 205

//
const QLocale__Teso QLocale__Language = 206

//
const QLocale__Saho QLocale__Language = 207

//
const QLocale__KoyraChiini QLocale__Language = 208

//
const QLocale__Rwa QLocale__Language = 209

//
const QLocale__Luo QLocale__Language = 210

//
const QLocale__Chiga QLocale__Language = 211

//
const QLocale__CentralMoroccoTamazight QLocale__Language = 212

//
const QLocale__KoyraboroSenni QLocale__Language = 213

//
const QLocale__Shambala QLocale__Language = 214

//
const QLocale__Bodo QLocale__Language = 215

//
const QLocale__Avaric QLocale__Language = 216

//
const QLocale__Chamorro QLocale__Language = 217

//
const QLocale__Chechen QLocale__Language = 218

//
const QLocale__Church QLocale__Language = 219

//
const QLocale__Chuvash QLocale__Language = 220

//
const QLocale__Cree QLocale__Language = 221

//
const QLocale__Haitian QLocale__Language = 222

//
const QLocale__Herero QLocale__Language = 223

//
const QLocale__HiriMotu QLocale__Language = 224

//
const QLocale__Kanuri QLocale__Language = 225

//
const QLocale__Komi QLocale__Language = 226

//
const QLocale__Kongo QLocale__Language = 227

//
const QLocale__Kwanyama QLocale__Language = 228

//
const QLocale__Limburgish QLocale__Language = 229

//
const QLocale__LubaKatanga QLocale__Language = 230

//
const QLocale__Luxembourgish QLocale__Language = 231

//
const QLocale__Navaho QLocale__Language = 232

//
const QLocale__Ndonga QLocale__Language = 233

//
const QLocale__Ojibwa QLocale__Language = 234

//
const QLocale__Pali QLocale__Language = 235

//
const QLocale__Walloon QLocale__Language = 236

//
const QLocale__Aghem QLocale__Language = 237

//
const QLocale__Basaa QLocale__Language = 238

//
const QLocale__Zarma QLocale__Language = 239

//
const QLocale__Duala QLocale__Language = 240

//
const QLocale__JolaFonyi QLocale__Language = 241

//
const QLocale__Ewondo QLocale__Language = 242

//
const QLocale__Bafia QLocale__Language = 243

//
const QLocale__MakhuwaMeetto QLocale__Language = 244

//
const QLocale__Mundang QLocale__Language = 245

//
const QLocale__Kwasio QLocale__Language = 246

//
const QLocale__Nuer QLocale__Language = 247

//
const QLocale__Sakha QLocale__Language = 248

//
const QLocale__Sangu QLocale__Language = 249

//
const QLocale__CongoSwahili QLocale__Language = 250

//
const QLocale__Tasawaq QLocale__Language = 251

//
const QLocale__Vai QLocale__Language = 252

//
const QLocale__Walser QLocale__Language = 253

//
const QLocale__Yangben QLocale__Language = 254

//
const QLocale__Avestan QLocale__Language = 255

//
const QLocale__Asturian QLocale__Language = 256

//
const QLocale__Ngomba QLocale__Language = 257

//
const QLocale__Kako QLocale__Language = 258

//
const QLocale__Meta QLocale__Language = 259

//
const QLocale__Ngiemboon QLocale__Language = 260

//
const QLocale__Aragonese QLocale__Language = 261

//
const QLocale__Akkadian QLocale__Language = 262

//
const QLocale__AncientEgyptian QLocale__Language = 263

//
const QLocale__AncientGreek QLocale__Language = 264

//
const QLocale__Aramaic QLocale__Language = 265

//
const QLocale__Balinese QLocale__Language = 266

//
const QLocale__Bamun QLocale__Language = 267

//
const QLocale__BatakToba QLocale__Language = 268

//
const QLocale__Buginese QLocale__Language = 269

//
const QLocale__Buhid QLocale__Language = 270

//
const QLocale__Carian QLocale__Language = 271

//
const QLocale__Chakma QLocale__Language = 272

//
const QLocale__ClassicalMandaic QLocale__Language = 273

//
const QLocale__Coptic QLocale__Language = 274

//
const QLocale__Dogri QLocale__Language = 275

//
const QLocale__EasternCham QLocale__Language = 276

//
const QLocale__EasternKayah QLocale__Language = 277

//
const QLocale__Etruscan QLocale__Language = 278

//
const QLocale__Gothic QLocale__Language = 279

//
const QLocale__Hanunoo QLocale__Language = 280

//
const QLocale__Ingush QLocale__Language = 281

//
const QLocale__LargeFloweryMiao QLocale__Language = 282

//
const QLocale__Lepcha QLocale__Language = 283

//
const QLocale__Limbu QLocale__Language = 284

//
const QLocale__Lisu QLocale__Language = 285

//
const QLocale__Lu QLocale__Language = 286

//
const QLocale__Lycian QLocale__Language = 287

//
const QLocale__Lydian QLocale__Language = 288

//
const QLocale__Mandingo QLocale__Language = 289

//
const QLocale__Manipuri QLocale__Language = 290

//
const QLocale__Meroitic QLocale__Language = 291

//
const QLocale__NorthernThai QLocale__Language = 292

//
const QLocale__OldIrish QLocale__Language = 293

//
const QLocale__OldNorse QLocale__Language = 294

//
const QLocale__OldPersian QLocale__Language = 295

//
const QLocale__OldTurkish QLocale__Language = 296

//
const QLocale__Pahlavi QLocale__Language = 297

//
const QLocale__Parthian QLocale__Language = 298

//
const QLocale__Phoenician QLocale__Language = 299

//
const QLocale__PrakritLanguage QLocale__Language = 300

//
const QLocale__Rejang QLocale__Language = 301

//
const QLocale__Sabaean QLocale__Language = 302

//
const QLocale__Samaritan QLocale__Language = 303

//
const QLocale__Santali QLocale__Language = 304

//
const QLocale__Saurashtra QLocale__Language = 305

//
const QLocale__Sora QLocale__Language = 306

//
const QLocale__Sylheti QLocale__Language = 307

//
const QLocale__Tagbanwa QLocale__Language = 308

//
const QLocale__TaiDam QLocale__Language = 309

//
const QLocale__TaiNua QLocale__Language = 310

//
const QLocale__Ugaritic QLocale__Language = 311

//
const QLocale__Akoose QLocale__Language = 312

//
const QLocale__Lakota QLocale__Language = 313

//
const QLocale__StandardMoroccanTamazight QLocale__Language = 314

//
const QLocale__Mapuche QLocale__Language = 315

//
const QLocale__CentralKurdish QLocale__Language = 316

//
const QLocale__LowerSorbian QLocale__Language = 317

//
const QLocale__UpperSorbian QLocale__Language = 318

//
const QLocale__Kenyang QLocale__Language = 319

//
const QLocale__Mohawk QLocale__Language = 320

//
const QLocale__Nko QLocale__Language = 321

//
const QLocale__Prussian QLocale__Language = 322

//
const QLocale__Kiche QLocale__Language = 323

//
const QLocale__SouthernSami QLocale__Language = 324

//
const QLocale__LuleSami QLocale__Language = 325

//
const QLocale__InariSami QLocale__Language = 326

//
const QLocale__SkoltSami QLocale__Language = 327

//
const QLocale__Warlpiri QLocale__Language = 328

//
const QLocale__ManichaeanMiddlePersian QLocale__Language = 329

//
const QLocale__Mende QLocale__Language = 330

//
const QLocale__AncientNorthArabian QLocale__Language = 331

//
const QLocale__LinearA QLocale__Language = 332

//
const QLocale__HmongNjua QLocale__Language = 333

//
const QLocale__Ho QLocale__Language = 334

//
const QLocale__Lezghian QLocale__Language = 335

//
const QLocale__Bassa QLocale__Language = 336

//
const QLocale__Mono QLocale__Language = 337

//
const QLocale__TedimChin QLocale__Language = 338

//
const QLocale__Maithili QLocale__Language = 339

//
const QLocale__Ahom QLocale__Language = 340

//
const QLocale__AmericanSignLanguage QLocale__Language = 341

//
const QLocale__ArdhamagadhiPrakrit QLocale__Language = 342

//
const QLocale__Bhojpuri QLocale__Language = 343

//
const QLocale__HieroglyphicLuwian QLocale__Language = 344

//
const QLocale__LiteraryChinese QLocale__Language = 345

//
const QLocale__Mazanderani QLocale__Language = 346

//
const QLocale__Mru QLocale__Language = 347

//
const QLocale__Newari QLocale__Language = 348

//
const QLocale__NorthernLuri QLocale__Language = 349

//
const QLocale__Palauan QLocale__Language = 350

//
const QLocale__Papiamento QLocale__Language = 351

//
const QLocale__Saraiki QLocale__Language = 352

//
const QLocale__TokelauLanguage QLocale__Language = 353

//
const QLocale__TokPisin QLocale__Language = 354

//
const QLocale__TuvaluLanguage QLocale__Language = 355

//
const QLocale__UncodedLanguages QLocale__Language = 356

//
const QLocale__Cantonese QLocale__Language = 357

//
const QLocale__Osage QLocale__Language = 358

//
const QLocale__Tangut QLocale__Language = 359

//
const QLocale__Norwegian QLocale__Language = 85

//
const QLocale__Moldavian QLocale__Language = 95

//
const QLocale__SerboCroatian QLocale__Language = 100

//
const QLocale__Tagalog QLocale__Language = 166

//
const QLocale__Twi QLocale__Language = 146

//
const QLocale__Afan QLocale__Language = 3

//
const QLocale__Byelorussian QLocale__Language = 22

//
const QLocale__Bhutani QLocale__Language = 16

//
const QLocale__Cambodian QLocale__Language = 23

//
const QLocale__Kurundi QLocale__Language = 68

//
const QLocale__RhaetoRomance QLocale__Language = 94

//
const QLocale__Chewa QLocale__Language = 165

//
const QLocale__Frisian QLocale__Language = 38

//
const QLocale__Uigur QLocale__Language = 128

//
const QLocale__LastLanguage QLocale__Language = 359

func (this *QLocale) LanguageItemName(val int) string {
	switch val {
	case QLocale__AnyLanguage: // 0
		return "AnyLanguage"
	case QLocale__C: // 1
		return "C"
	case QLocale__Abkhazian: // 2
		return "Abkhazian"
	case QLocale__Oromo: // 3
		return "Oromo,Afan"
	case QLocale__Afar: // 4
		return "Afar"
	case QLocale__Afrikaans: // 5
		return "Afrikaans"
	case QLocale__Albanian: // 6
		return "Albanian"
	case QLocale__Amharic: // 7
		return "Amharic"
	case QLocale__Arabic: // 8
		return "Arabic"
	case QLocale__Armenian: // 9
		return "Armenian"
	case QLocale__Assamese: // 10
		return "Assamese"
	case QLocale__Aymara: // 11
		return "Aymara"
	case QLocale__Azerbaijani: // 12
		return "Azerbaijani"
	case QLocale__Bashkir: // 13
		return "Bashkir"
	case QLocale__Basque: // 14
		return "Basque"
	case QLocale__Bengali: // 15
		return "Bengali"
	case QLocale__Dzongkha: // 16
		return "Dzongkha,Bhutani"
	case QLocale__Bihari: // 17
		return "Bihari"
	case QLocale__Bislama: // 18
		return "Bislama"
	case QLocale__Breton: // 19
		return "Breton"
	case QLocale__Bulgarian: // 20
		return "Bulgarian"
	case QLocale__Burmese: // 21
		return "Burmese"
	case QLocale__Belarusian: // 22
		return "Belarusian,Byelorussian"
	case QLocale__Khmer: // 23
		return "Khmer,Cambodian"
	case QLocale__Catalan: // 24
		return "Catalan"
	case QLocale__Chinese: // 25
		return "Chinese"
	case QLocale__Corsican: // 26
		return "Corsican"
	case QLocale__Croatian: // 27
		return "Croatian"
	case QLocale__Czech: // 28
		return "Czech"
	case QLocale__Danish: // 29
		return "Danish"
	case QLocale__Dutch: // 30
		return "Dutch"
	case QLocale__English: // 31
		return "English"
	case QLocale__Esperanto: // 32
		return "Esperanto"
	case QLocale__Estonian: // 33
		return "Estonian"
	case QLocale__Faroese: // 34
		return "Faroese"
	case QLocale__Fijian: // 35
		return "Fijian"
	case QLocale__Finnish: // 36
		return "Finnish"
	case QLocale__French: // 37
		return "French"
	case QLocale__WesternFrisian: // 38
		return "WesternFrisian,Frisian"
	case QLocale__Gaelic: // 39
		return "Gaelic"
	case QLocale__Galician: // 40
		return "Galician"
	case QLocale__Georgian: // 41
		return "Georgian"
	case QLocale__German: // 42
		return "German"
	case QLocale__Greek: // 43
		return "Greek"
	case QLocale__Greenlandic: // 44
		return "Greenlandic"
	case QLocale__Guarani: // 45
		return "Guarani"
	case QLocale__Gujarati: // 46
		return "Gujarati"
	case QLocale__Hausa: // 47
		return "Hausa"
	case QLocale__Hebrew: // 48
		return "Hebrew"
	case QLocale__Hindi: // 49
		return "Hindi"
	case QLocale__Hungarian: // 50
		return "Hungarian"
	case QLocale__Icelandic: // 51
		return "Icelandic"
	case QLocale__Indonesian: // 52
		return "Indonesian"
	case QLocale__Interlingua: // 53
		return "Interlingua"
	case QLocale__Interlingue: // 54
		return "Interlingue"
	case QLocale__Inuktitut: // 55
		return "Inuktitut"
	case QLocale__Inupiak: // 56
		return "Inupiak"
	case QLocale__Irish: // 57
		return "Irish"
	case QLocale__Italian: // 58
		return "Italian"
	case QLocale__Japanese: // 59
		return "Japanese"
	case QLocale__Javanese: // 60
		return "Javanese"
	case QLocale__Kannada: // 61
		return "Kannada"
	case QLocale__Kashmiri: // 62
		return "Kashmiri"
	case QLocale__Kazakh: // 63
		return "Kazakh"
	case QLocale__Kinyarwanda: // 64
		return "Kinyarwanda"
	case QLocale__Kirghiz: // 65
		return "Kirghiz"
	case QLocale__Korean: // 66
		return "Korean"
	case QLocale__Kurdish: // 67
		return "Kurdish"
	case QLocale__Rundi: // 68
		return "Rundi,Kurundi"
	case QLocale__Lao: // 69
		return "Lao"
	case QLocale__Latin: // 70
		return "Latin"
	case QLocale__Latvian: // 71
		return "Latvian"
	case QLocale__Lingala: // 72
		return "Lingala"
	case QLocale__Lithuanian: // 73
		return "Lithuanian"
	case QLocale__Macedonian: // 74
		return "Macedonian"
	case QLocale__Malagasy: // 75
		return "Malagasy"
	case QLocale__Malay: // 76
		return "Malay"
	case QLocale__Malayalam: // 77
		return "Malayalam"
	case QLocale__Maltese: // 78
		return "Maltese"
	case QLocale__Maori: // 79
		return "Maori"
	case QLocale__Marathi: // 80
		return "Marathi"
	case QLocale__Marshallese: // 81
		return "Marshallese"
	case QLocale__Mongolian: // 82
		return "Mongolian"
	case QLocale__NauruLanguage: // 83
		return "NauruLanguage"
	case QLocale__Nepali: // 84
		return "Nepali"
	case QLocale__NorwegianBokmal: // 85
		return "NorwegianBokmal,Norwegian"
	case QLocale__Occitan: // 86
		return "Occitan"
	case QLocale__Oriya: // 87
		return "Oriya"
	case QLocale__Pashto: // 88
		return "Pashto"
	case QLocale__Persian: // 89
		return "Persian"
	case QLocale__Polish: // 90
		return "Polish"
	case QLocale__Portuguese: // 91
		return "Portuguese"
	case QLocale__Punjabi: // 92
		return "Punjabi"
	case QLocale__Quechua: // 93
		return "Quechua"
	case QLocale__Romansh: // 94
		return "Romansh,RhaetoRomance"
	case QLocale__Romanian: // 95
		return "Romanian,Moldavian"
	case QLocale__Russian: // 96
		return "Russian"
	case QLocale__Samoan: // 97
		return "Samoan"
	case QLocale__Sango: // 98
		return "Sango"
	case QLocale__Sanskrit: // 99
		return "Sanskrit"
	case QLocale__Serbian: // 100
		return "Serbian,SerboCroatian"
	case QLocale__Ossetic: // 101
		return "Ossetic"
	case QLocale__SouthernSotho: // 102
		return "SouthernSotho"
	case QLocale__Tswana: // 103
		return "Tswana"
	case QLocale__Shona: // 104
		return "Shona"
	case QLocale__Sindhi: // 105
		return "Sindhi"
	case QLocale__Sinhala: // 106
		return "Sinhala"
	case QLocale__Swati: // 107
		return "Swati"
	case QLocale__Slovak: // 108
		return "Slovak"
	case QLocale__Slovenian: // 109
		return "Slovenian"
	case QLocale__Somali: // 110
		return "Somali"
	case QLocale__Spanish: // 111
		return "Spanish"
	case QLocale__Sundanese: // 112
		return "Sundanese"
	case QLocale__Swahili: // 113
		return "Swahili"
	case QLocale__Swedish: // 114
		return "Swedish"
	case QLocale__Sardinian: // 115
		return "Sardinian"
	case QLocale__Tajik: // 116
		return "Tajik"
	case QLocale__Tamil: // 117
		return "Tamil"
	case QLocale__Tatar: // 118
		return "Tatar"
	case QLocale__Telugu: // 119
		return "Telugu"
	case QLocale__Thai: // 120
		return "Thai"
	case QLocale__Tibetan: // 121
		return "Tibetan"
	case QLocale__Tigrinya: // 122
		return "Tigrinya"
	case QLocale__Tongan: // 123
		return "Tongan"
	case QLocale__Tsonga: // 124
		return "Tsonga"
	case QLocale__Turkish: // 125
		return "Turkish"
	case QLocale__Turkmen: // 126
		return "Turkmen"
	case QLocale__Tahitian: // 127
		return "Tahitian"
	case QLocale__Uighur: // 128
		return "Uighur,Uigur"
	case QLocale__Ukrainian: // 129
		return "Ukrainian"
	case QLocale__Urdu: // 130
		return "Urdu"
	case QLocale__Uzbek: // 131
		return "Uzbek"
	case QLocale__Vietnamese: // 132
		return "Vietnamese"
	case QLocale__Volapuk: // 133
		return "Volapuk"
	case QLocale__Welsh: // 134
		return "Welsh"
	case QLocale__Wolof: // 135
		return "Wolof"
	case QLocale__Xhosa: // 136
		return "Xhosa"
	case QLocale__Yiddish: // 137
		return "Yiddish"
	case QLocale__Yoruba: // 138
		return "Yoruba"
	case QLocale__Zhuang: // 139
		return "Zhuang"
	case QLocale__Zulu: // 140
		return "Zulu"
	case QLocale__NorwegianNynorsk: // 141
		return "NorwegianNynorsk"
	case QLocale__Bosnian: // 142
		return "Bosnian"
	case QLocale__Divehi: // 143
		return "Divehi"
	case QLocale__Manx: // 144
		return "Manx"
	case QLocale__Cornish: // 145
		return "Cornish"
	case QLocale__Akan: // 146
		return "Akan,Twi"
	case QLocale__Konkani: // 147
		return "Konkani"
	case QLocale__Ga: // 148
		return "Ga"
	case QLocale__Igbo: // 149
		return "Igbo"
	case QLocale__Kamba: // 150
		return "Kamba"
	case QLocale__Syriac: // 151
		return "Syriac"
	case QLocale__Blin: // 152
		return "Blin"
	case QLocale__Geez: // 153
		return "Geez"
	case QLocale__Koro: // 154
		return "Koro"
	case QLocale__Sidamo: // 155
		return "Sidamo"
	case QLocale__Atsam: // 156
		return "Atsam"
	case QLocale__Tigre: // 157
		return "Tigre"
	case QLocale__Jju: // 158
		return "Jju"
	case QLocale__Friulian: // 159
		return "Friulian"
	case QLocale__Venda: // 160
		return "Venda"
	case QLocale__Ewe: // 161
		return "Ewe"
	case QLocale__Walamo: // 162
		return "Walamo"
	case QLocale__Hawaiian: // 163
		return "Hawaiian"
	case QLocale__Tyap: // 164
		return "Tyap"
	case QLocale__Nyanja: // 165
		return "Nyanja,Chewa"
	case QLocale__Filipino: // 166
		return "Filipino,Tagalog"
	case QLocale__SwissGerman: // 167
		return "SwissGerman"
	case QLocale__SichuanYi: // 168
		return "SichuanYi"
	case QLocale__Kpelle: // 169
		return "Kpelle"
	case QLocale__LowGerman: // 170
		return "LowGerman"
	case QLocale__SouthNdebele: // 171
		return "SouthNdebele"
	case QLocale__NorthernSotho: // 172
		return "NorthernSotho"
	case QLocale__NorthernSami: // 173
		return "NorthernSami"
	case QLocale__Taroko: // 174
		return "Taroko"
	case QLocale__Gusii: // 175
		return "Gusii"
	case QLocale__Taita: // 176
		return "Taita"
	case QLocale__Fulah: // 177
		return "Fulah"
	case QLocale__Kikuyu: // 178
		return "Kikuyu"
	case QLocale__Samburu: // 179
		return "Samburu"
	case QLocale__Sena: // 180
		return "Sena"
	case QLocale__NorthNdebele: // 181
		return "NorthNdebele"
	case QLocale__Rombo: // 182
		return "Rombo"
	case QLocale__Tachelhit: // 183
		return "Tachelhit"
	case QLocale__Kabyle: // 184
		return "Kabyle"
	case QLocale__Nyankole: // 185
		return "Nyankole"
	case QLocale__Bena: // 186
		return "Bena"
	case QLocale__Vunjo: // 187
		return "Vunjo"
	case QLocale__Bambara: // 188
		return "Bambara"
	case QLocale__Embu: // 189
		return "Embu"
	case QLocale__Cherokee: // 190
		return "Cherokee"
	case QLocale__Morisyen: // 191
		return "Morisyen"
	case QLocale__Makonde: // 192
		return "Makonde"
	case QLocale__Langi: // 193
		return "Langi"
	case QLocale__Ganda: // 194
		return "Ganda"
	case QLocale__Bemba: // 195
		return "Bemba"
	case QLocale__Kabuverdianu: // 196
		return "Kabuverdianu"
	case QLocale__Meru: // 197
		return "Meru"
	case QLocale__Kalenjin: // 198
		return "Kalenjin"
	case QLocale__Nama: // 199
		return "Nama"
	case QLocale__Machame: // 200
		return "Machame"
	case QLocale__Colognian: // 201
		return "Colognian"
	case QLocale__Masai: // 202
		return "Masai"
	case QLocale__Soga: // 203
		return "Soga"
	case QLocale__Luyia: // 204
		return "Luyia"
	case QLocale__Asu: // 205
		return "Asu"
	case QLocale__Teso: // 206
		return "Teso"
	case QLocale__Saho: // 207
		return "Saho"
	case QLocale__KoyraChiini: // 208
		return "KoyraChiini"
	case QLocale__Rwa: // 209
		return "Rwa"
	case QLocale__Luo: // 210
		return "Luo"
	case QLocale__Chiga: // 211
		return "Chiga"
	case QLocale__CentralMoroccoTamazight: // 212
		return "CentralMoroccoTamazight"
	case QLocale__KoyraboroSenni: // 213
		return "KoyraboroSenni"
	case QLocale__Shambala: // 214
		return "Shambala"
	case QLocale__Bodo: // 215
		return "Bodo"
	case QLocale__Avaric: // 216
		return "Avaric"
	case QLocale__Chamorro: // 217
		return "Chamorro"
	case QLocale__Chechen: // 218
		return "Chechen"
	case QLocale__Church: // 219
		return "Church"
	case QLocale__Chuvash: // 220
		return "Chuvash"
	case QLocale__Cree: // 221
		return "Cree"
	case QLocale__Haitian: // 222
		return "Haitian"
	case QLocale__Herero: // 223
		return "Herero"
	case QLocale__HiriMotu: // 224
		return "HiriMotu"
	case QLocale__Kanuri: // 225
		return "Kanuri"
	case QLocale__Komi: // 226
		return "Komi"
	case QLocale__Kongo: // 227
		return "Kongo"
	case QLocale__Kwanyama: // 228
		return "Kwanyama"
	case QLocale__Limburgish: // 229
		return "Limburgish"
	case QLocale__LubaKatanga: // 230
		return "LubaKatanga"
	case QLocale__Luxembourgish: // 231
		return "Luxembourgish"
	case QLocale__Navaho: // 232
		return "Navaho"
	case QLocale__Ndonga: // 233
		return "Ndonga"
	case QLocale__Ojibwa: // 234
		return "Ojibwa"
	case QLocale__Pali: // 235
		return "Pali"
	case QLocale__Walloon: // 236
		return "Walloon"
	case QLocale__Aghem: // 237
		return "Aghem"
	case QLocale__Basaa: // 238
		return "Basaa"
	case QLocale__Zarma: // 239
		return "Zarma"
	case QLocale__Duala: // 240
		return "Duala"
	case QLocale__JolaFonyi: // 241
		return "JolaFonyi"
	case QLocale__Ewondo: // 242
		return "Ewondo"
	case QLocale__Bafia: // 243
		return "Bafia"
	case QLocale__MakhuwaMeetto: // 244
		return "MakhuwaMeetto"
	case QLocale__Mundang: // 245
		return "Mundang"
	case QLocale__Kwasio: // 246
		return "Kwasio"
	case QLocale__Nuer: // 247
		return "Nuer"
	case QLocale__Sakha: // 248
		return "Sakha"
	case QLocale__Sangu: // 249
		return "Sangu"
	case QLocale__CongoSwahili: // 250
		return "CongoSwahili"
	case QLocale__Tasawaq: // 251
		return "Tasawaq"
	case QLocale__Vai: // 252
		return "Vai"
	case QLocale__Walser: // 253
		return "Walser"
	case QLocale__Yangben: // 254
		return "Yangben"
	case QLocale__Avestan: // 255
		return "Avestan"
	case QLocale__Asturian: // 256
		return "Asturian"
	case QLocale__Ngomba: // 257
		return "Ngomba"
	case QLocale__Kako: // 258
		return "Kako"
	case QLocale__Meta: // 259
		return "Meta"
	case QLocale__Ngiemboon: // 260
		return "Ngiemboon"
	case QLocale__Aragonese: // 261
		return "Aragonese"
	case QLocale__Akkadian: // 262
		return "Akkadian"
	case QLocale__AncientEgyptian: // 263
		return "AncientEgyptian"
	case QLocale__AncientGreek: // 264
		return "AncientGreek"
	case QLocale__Aramaic: // 265
		return "Aramaic"
	case QLocale__Balinese: // 266
		return "Balinese"
	case QLocale__Bamun: // 267
		return "Bamun"
	case QLocale__BatakToba: // 268
		return "BatakToba"
	case QLocale__Buginese: // 269
		return "Buginese"
	case QLocale__Buhid: // 270
		return "Buhid"
	case QLocale__Carian: // 271
		return "Carian"
	case QLocale__Chakma: // 272
		return "Chakma"
	case QLocale__ClassicalMandaic: // 273
		return "ClassicalMandaic"
	case QLocale__Coptic: // 274
		return "Coptic"
	case QLocale__Dogri: // 275
		return "Dogri"
	case QLocale__EasternCham: // 276
		return "EasternCham"
	case QLocale__EasternKayah: // 277
		return "EasternKayah"
	case QLocale__Etruscan: // 278
		return "Etruscan"
	case QLocale__Gothic: // 279
		return "Gothic"
	case QLocale__Hanunoo: // 280
		return "Hanunoo"
	case QLocale__Ingush: // 281
		return "Ingush"
	case QLocale__LargeFloweryMiao: // 282
		return "LargeFloweryMiao"
	case QLocale__Lepcha: // 283
		return "Lepcha"
	case QLocale__Limbu: // 284
		return "Limbu"
	case QLocale__Lisu: // 285
		return "Lisu"
	case QLocale__Lu: // 286
		return "Lu"
	case QLocale__Lycian: // 287
		return "Lycian"
	case QLocale__Lydian: // 288
		return "Lydian"
	case QLocale__Mandingo: // 289
		return "Mandingo"
	case QLocale__Manipuri: // 290
		return "Manipuri"
	case QLocale__Meroitic: // 291
		return "Meroitic"
	case QLocale__NorthernThai: // 292
		return "NorthernThai"
	case QLocale__OldIrish: // 293
		return "OldIrish"
	case QLocale__OldNorse: // 294
		return "OldNorse"
	case QLocale__OldPersian: // 295
		return "OldPersian"
	case QLocale__OldTurkish: // 296
		return "OldTurkish"
	case QLocale__Pahlavi: // 297
		return "Pahlavi"
	case QLocale__Parthian: // 298
		return "Parthian"
	case QLocale__Phoenician: // 299
		return "Phoenician"
	case QLocale__PrakritLanguage: // 300
		return "PrakritLanguage"
	case QLocale__Rejang: // 301
		return "Rejang"
	case QLocale__Sabaean: // 302
		return "Sabaean"
	case QLocale__Samaritan: // 303
		return "Samaritan"
	case QLocale__Santali: // 304
		return "Santali"
	case QLocale__Saurashtra: // 305
		return "Saurashtra"
	case QLocale__Sora: // 306
		return "Sora"
	case QLocale__Sylheti: // 307
		return "Sylheti"
	case QLocale__Tagbanwa: // 308
		return "Tagbanwa"
	case QLocale__TaiDam: // 309
		return "TaiDam"
	case QLocale__TaiNua: // 310
		return "TaiNua"
	case QLocale__Ugaritic: // 311
		return "Ugaritic"
	case QLocale__Akoose: // 312
		return "Akoose"
	case QLocale__Lakota: // 313
		return "Lakota"
	case QLocale__StandardMoroccanTamazight: // 314
		return "StandardMoroccanTamazight"
	case QLocale__Mapuche: // 315
		return "Mapuche"
	case QLocale__CentralKurdish: // 316
		return "CentralKurdish"
	case QLocale__LowerSorbian: // 317
		return "LowerSorbian"
	case QLocale__UpperSorbian: // 318
		return "UpperSorbian"
	case QLocale__Kenyang: // 319
		return "Kenyang"
	case QLocale__Mohawk: // 320
		return "Mohawk"
	case QLocale__Nko: // 321
		return "Nko"
	case QLocale__Prussian: // 322
		return "Prussian"
	case QLocale__Kiche: // 323
		return "Kiche"
	case QLocale__SouthernSami: // 324
		return "SouthernSami"
	case QLocale__LuleSami: // 325
		return "LuleSami"
	case QLocale__InariSami: // 326
		return "InariSami"
	case QLocale__SkoltSami: // 327
		return "SkoltSami"
	case QLocale__Warlpiri: // 328
		return "Warlpiri"
	case QLocale__ManichaeanMiddlePersian: // 329
		return "ManichaeanMiddlePersian"
	case QLocale__Mende: // 330
		return "Mende"
	case QLocale__AncientNorthArabian: // 331
		return "AncientNorthArabian"
	case QLocale__LinearA: // 332
		return "LinearA"
	case QLocale__HmongNjua: // 333
		return "HmongNjua"
	case QLocale__Ho: // 334
		return "Ho"
	case QLocale__Lezghian: // 335
		return "Lezghian"
	case QLocale__Bassa: // 336
		return "Bassa"
	case QLocale__Mono: // 337
		return "Mono"
	case QLocale__TedimChin: // 338
		return "TedimChin"
	case QLocale__Maithili: // 339
		return "Maithili"
	case QLocale__Ahom: // 340
		return "Ahom"
	case QLocale__AmericanSignLanguage: // 341
		return "AmericanSignLanguage"
	case QLocale__ArdhamagadhiPrakrit: // 342
		return "ArdhamagadhiPrakrit"
	case QLocale__Bhojpuri: // 343
		return "Bhojpuri"
	case QLocale__HieroglyphicLuwian: // 344
		return "HieroglyphicLuwian"
	case QLocale__LiteraryChinese: // 345
		return "LiteraryChinese"
	case QLocale__Mazanderani: // 346
		return "Mazanderani"
	case QLocale__Mru: // 347
		return "Mru"
	case QLocale__Newari: // 348
		return "Newari"
	case QLocale__NorthernLuri: // 349
		return "NorthernLuri"
	case QLocale__Palauan: // 350
		return "Palauan"
	case QLocale__Papiamento: // 351
		return "Papiamento"
	case QLocale__Saraiki: // 352
		return "Saraiki"
	case QLocale__TokelauLanguage: // 353
		return "TokelauLanguage"
	case QLocale__TokPisin: // 354
		return "TokPisin"
	case QLocale__TuvaluLanguage: // 355
		return "TuvaluLanguage"
	case QLocale__UncodedLanguages: // 356
		return "UncodedLanguages"
	case QLocale__Cantonese: // 357
		return "Cantonese"
	case QLocale__Osage: // 358
		return "Osage"
	case QLocale__Tangut: // 359
		return "Tangut,LastLanguage"
		// case QLocale__Norwegian: // 85
		// return ""
		// case QLocale__Moldavian: // 95
		// return ""
		// case QLocale__SerboCroatian: // 100
		// return ""
		// case QLocale__Tagalog: // 166
		// return ""
		// case QLocale__Twi: // 146
		// return ""
		// case QLocale__Afan: // 3
		// return ""
		// case QLocale__Byelorussian: // 22
		// return ""
		// case QLocale__Bhutani: // 16
		// return ""
		// case QLocale__Cambodian: // 23
		// return ""
		// case QLocale__Kurundi: // 68
		// return ""
		// case QLocale__RhaetoRomance: // 94
		// return ""
		// case QLocale__Chewa: // 165
		// return ""
		// case QLocale__Frisian: // 38
		// return ""
		// case QLocale__Uigur: // 128
		// return ""
		// case QLocale__LastLanguage: // 359
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_LanguageItemName(val int) string {
	var nilthis *QLocale
	return nilthis.LanguageItemName(val)
}

/*
This enumerated type is used to specify a script.

QLocale::SimplifiedChineseScriptSimplifiedHanScriptsame as SimplifiedHanScript
QLocale::TraditionalChineseScriptTraditionalHanScriptsame as TraditionalHanScript


See also script(), scriptToString(), and languageToString().

*/
type QLocale__Script = int

//
const QLocale__AnyScript QLocale__Script = 0

//
const QLocale__ArabicScript QLocale__Script = 1

//
const QLocale__CyrillicScript QLocale__Script = 2

//
const QLocale__DeseretScript QLocale__Script = 3

//
const QLocale__GurmukhiScript QLocale__Script = 4

// same as SimplifiedChineseScript
const QLocale__SimplifiedHanScript QLocale__Script = 5

// same as TraditionalChineseScript
const QLocale__TraditionalHanScript QLocale__Script = 6

//
const QLocale__LatinScript QLocale__Script = 7

//
const QLocale__MongolianScript QLocale__Script = 8

//
const QLocale__TifinaghScript QLocale__Script = 9

//
const QLocale__ArmenianScript QLocale__Script = 10

//
const QLocale__BengaliScript QLocale__Script = 11

//
const QLocale__CherokeeScript QLocale__Script = 12

//
const QLocale__DevanagariScript QLocale__Script = 13

//
const QLocale__EthiopicScript QLocale__Script = 14

//
const QLocale__GeorgianScript QLocale__Script = 15

//
const QLocale__GreekScript QLocale__Script = 16

//
const QLocale__GujaratiScript QLocale__Script = 17

//
const QLocale__HebrewScript QLocale__Script = 18

//
const QLocale__JapaneseScript QLocale__Script = 19

//
const QLocale__KhmerScript QLocale__Script = 20

//
const QLocale__KannadaScript QLocale__Script = 21

//
const QLocale__KoreanScript QLocale__Script = 22

//
const QLocale__LaoScript QLocale__Script = 23

//
const QLocale__MalayalamScript QLocale__Script = 24

//
const QLocale__MyanmarScript QLocale__Script = 25

//
const QLocale__OriyaScript QLocale__Script = 26

//
const QLocale__TamilScript QLocale__Script = 27

//
const QLocale__TeluguScript QLocale__Script = 28

//
const QLocale__ThaanaScript QLocale__Script = 29

//
const QLocale__ThaiScript QLocale__Script = 30

//
const QLocale__TibetanScript QLocale__Script = 31

//
const QLocale__SinhalaScript QLocale__Script = 32

//
const QLocale__SyriacScript QLocale__Script = 33

//
const QLocale__YiScript QLocale__Script = 34

//
const QLocale__VaiScript QLocale__Script = 35

//
const QLocale__AvestanScript QLocale__Script = 36

//
const QLocale__BalineseScript QLocale__Script = 37

//
const QLocale__BamumScript QLocale__Script = 38

//
const QLocale__BatakScript QLocale__Script = 39

//
const QLocale__BopomofoScript QLocale__Script = 40

//
const QLocale__BrahmiScript QLocale__Script = 41

//
const QLocale__BugineseScript QLocale__Script = 42

//
const QLocale__BuhidScript QLocale__Script = 43

//
const QLocale__CanadianAboriginalScript QLocale__Script = 44

//
const QLocale__CarianScript QLocale__Script = 45

//
const QLocale__ChakmaScript QLocale__Script = 46

//
const QLocale__ChamScript QLocale__Script = 47

//
const QLocale__CopticScript QLocale__Script = 48

//
const QLocale__CypriotScript QLocale__Script = 49

//
const QLocale__EgyptianHieroglyphsScript QLocale__Script = 50

//
const QLocale__FraserScript QLocale__Script = 51

//
const QLocale__GlagoliticScript QLocale__Script = 52

//
const QLocale__GothicScript QLocale__Script = 53

//
const QLocale__HanScript QLocale__Script = 54

//
const QLocale__HangulScript QLocale__Script = 55

//
const QLocale__HanunooScript QLocale__Script = 56

//
const QLocale__ImperialAramaicScript QLocale__Script = 57

//
const QLocale__InscriptionalPahlaviScript QLocale__Script = 58

//
const QLocale__InscriptionalParthianScript QLocale__Script = 59

//
const QLocale__JavaneseScript QLocale__Script = 60

//
const QLocale__KaithiScript QLocale__Script = 61

//
const QLocale__KatakanaScript QLocale__Script = 62

//
const QLocale__KayahLiScript QLocale__Script = 63

//
const QLocale__KharoshthiScript QLocale__Script = 64

//
const QLocale__LannaScript QLocale__Script = 65

//
const QLocale__LepchaScript QLocale__Script = 66

//
const QLocale__LimbuScript QLocale__Script = 67

//
const QLocale__LinearBScript QLocale__Script = 68

//
const QLocale__LycianScript QLocale__Script = 69

//
const QLocale__LydianScript QLocale__Script = 70

//
const QLocale__MandaeanScript QLocale__Script = 71

//
const QLocale__MeiteiMayekScript QLocale__Script = 72

//
const QLocale__MeroiticScript QLocale__Script = 73

//
const QLocale__MeroiticCursiveScript QLocale__Script = 74

//
const QLocale__NkoScript QLocale__Script = 75

//
const QLocale__NewTaiLueScript QLocale__Script = 76

//
const QLocale__OghamScript QLocale__Script = 77

//
const QLocale__OlChikiScript QLocale__Script = 78

//
const QLocale__OldItalicScript QLocale__Script = 79

//
const QLocale__OldPersianScript QLocale__Script = 80

//
const QLocale__OldSouthArabianScript QLocale__Script = 81

//
const QLocale__OrkhonScript QLocale__Script = 82

//
const QLocale__OsmanyaScript QLocale__Script = 83

//
const QLocale__PhagsPaScript QLocale__Script = 84

//
const QLocale__PhoenicianScript QLocale__Script = 85

//
const QLocale__PollardPhoneticScript QLocale__Script = 86

//
const QLocale__RejangScript QLocale__Script = 87

//
const QLocale__RunicScript QLocale__Script = 88

//
const QLocale__SamaritanScript QLocale__Script = 89

//
const QLocale__SaurashtraScript QLocale__Script = 90

//
const QLocale__SharadaScript QLocale__Script = 91

//
const QLocale__ShavianScript QLocale__Script = 92

//
const QLocale__SoraSompengScript QLocale__Script = 93

//
const QLocale__CuneiformScript QLocale__Script = 94

//
const QLocale__SundaneseScript QLocale__Script = 95

//
const QLocale__SylotiNagriScript QLocale__Script = 96

//
const QLocale__TagalogScript QLocale__Script = 97

//
const QLocale__TagbanwaScript QLocale__Script = 98

//
const QLocale__TaiLeScript QLocale__Script = 99

//
const QLocale__TaiVietScript QLocale__Script = 100

//
const QLocale__TakriScript QLocale__Script = 101

//
const QLocale__UgariticScript QLocale__Script = 102

//
const QLocale__BrailleScript QLocale__Script = 103

//
const QLocale__HiraganaScript QLocale__Script = 104

//
const QLocale__CaucasianAlbanianScript QLocale__Script = 105

//
const QLocale__BassaVahScript QLocale__Script = 106

//
const QLocale__DuployanScript QLocale__Script = 107

//
const QLocale__ElbasanScript QLocale__Script = 108

//
const QLocale__GranthaScript QLocale__Script = 109

//
const QLocale__PahawhHmongScript QLocale__Script = 110

//
const QLocale__KhojkiScript QLocale__Script = 111

//
const QLocale__LinearAScript QLocale__Script = 112

//
const QLocale__MahajaniScript QLocale__Script = 113

//
const QLocale__ManichaeanScript QLocale__Script = 114

//
const QLocale__MendeKikakuiScript QLocale__Script = 115

//
const QLocale__ModiScript QLocale__Script = 116

//
const QLocale__MroScript QLocale__Script = 117

//
const QLocale__OldNorthArabianScript QLocale__Script = 118

//
const QLocale__NabataeanScript QLocale__Script = 119

//
const QLocale__PalmyreneScript QLocale__Script = 120

//
const QLocale__PauCinHauScript QLocale__Script = 121

//
const QLocale__OldPermicScript QLocale__Script = 122

//
const QLocale__PsalterPahlaviScript QLocale__Script = 123

//
const QLocale__SiddhamScript QLocale__Script = 124

//
const QLocale__KhudawadiScript QLocale__Script = 125

//
const QLocale__TirhutaScript QLocale__Script = 126

//
const QLocale__VarangKshitiScript QLocale__Script = 127

//
const QLocale__AhomScript QLocale__Script = 128

//
const QLocale__AnatolianHieroglyphsScript QLocale__Script = 129

//
const QLocale__HatranScript QLocale__Script = 130

//
const QLocale__MultaniScript QLocale__Script = 131

//
const QLocale__OldHungarianScript QLocale__Script = 132

//
const QLocale__SignWritingScript QLocale__Script = 133

//
const QLocale__AdlamScript QLocale__Script = 134

//
const QLocale__BhaiksukiScript QLocale__Script = 135

//
const QLocale__MarchenScript QLocale__Script = 136

//
const QLocale__NewaScript QLocale__Script = 137

//
const QLocale__OsageScript QLocale__Script = 138

//
const QLocale__TangutScript QLocale__Script = 139

//
const QLocale__HanWithBopomofoScript QLocale__Script = 140

//
const QLocale__JamoScript QLocale__Script = 141

//
const QLocale__SimplifiedChineseScript QLocale__Script = 5

//
const QLocale__TraditionalChineseScript QLocale__Script = 6

//
const QLocale__LastScript QLocale__Script = 141

func (this *QLocale) ScriptItemName(val int) string {
	switch val {
	case QLocale__AnyScript: // 0
		return "AnyScript"
	case QLocale__ArabicScript: // 1
		return "ArabicScript"
	case QLocale__CyrillicScript: // 2
		return "CyrillicScript"
	case QLocale__DeseretScript: // 3
		return "DeseretScript"
	case QLocale__GurmukhiScript: // 4
		return "GurmukhiScript"
	case QLocale__SimplifiedHanScript: // 5
		return "SimplifiedHanScript,SimplifiedChineseScript"
	case QLocale__TraditionalHanScript: // 6
		return "TraditionalHanScript,TraditionalChineseScript"
	case QLocale__LatinScript: // 7
		return "LatinScript"
	case QLocale__MongolianScript: // 8
		return "MongolianScript"
	case QLocale__TifinaghScript: // 9
		return "TifinaghScript"
	case QLocale__ArmenianScript: // 10
		return "ArmenianScript"
	case QLocale__BengaliScript: // 11
		return "BengaliScript"
	case QLocale__CherokeeScript: // 12
		return "CherokeeScript"
	case QLocale__DevanagariScript: // 13
		return "DevanagariScript"
	case QLocale__EthiopicScript: // 14
		return "EthiopicScript"
	case QLocale__GeorgianScript: // 15
		return "GeorgianScript"
	case QLocale__GreekScript: // 16
		return "GreekScript"
	case QLocale__GujaratiScript: // 17
		return "GujaratiScript"
	case QLocale__HebrewScript: // 18
		return "HebrewScript"
	case QLocale__JapaneseScript: // 19
		return "JapaneseScript"
	case QLocale__KhmerScript: // 20
		return "KhmerScript"
	case QLocale__KannadaScript: // 21
		return "KannadaScript"
	case QLocale__KoreanScript: // 22
		return "KoreanScript"
	case QLocale__LaoScript: // 23
		return "LaoScript"
	case QLocale__MalayalamScript: // 24
		return "MalayalamScript"
	case QLocale__MyanmarScript: // 25
		return "MyanmarScript"
	case QLocale__OriyaScript: // 26
		return "OriyaScript"
	case QLocale__TamilScript: // 27
		return "TamilScript"
	case QLocale__TeluguScript: // 28
		return "TeluguScript"
	case QLocale__ThaanaScript: // 29
		return "ThaanaScript"
	case QLocale__ThaiScript: // 30
		return "ThaiScript"
	case QLocale__TibetanScript: // 31
		return "TibetanScript"
	case QLocale__SinhalaScript: // 32
		return "SinhalaScript"
	case QLocale__SyriacScript: // 33
		return "SyriacScript"
	case QLocale__YiScript: // 34
		return "YiScript"
	case QLocale__VaiScript: // 35
		return "VaiScript"
	case QLocale__AvestanScript: // 36
		return "AvestanScript"
	case QLocale__BalineseScript: // 37
		return "BalineseScript"
	case QLocale__BamumScript: // 38
		return "BamumScript"
	case QLocale__BatakScript: // 39
		return "BatakScript"
	case QLocale__BopomofoScript: // 40
		return "BopomofoScript"
	case QLocale__BrahmiScript: // 41
		return "BrahmiScript"
	case QLocale__BugineseScript: // 42
		return "BugineseScript"
	case QLocale__BuhidScript: // 43
		return "BuhidScript"
	case QLocale__CanadianAboriginalScript: // 44
		return "CanadianAboriginalScript"
	case QLocale__CarianScript: // 45
		return "CarianScript"
	case QLocale__ChakmaScript: // 46
		return "ChakmaScript"
	case QLocale__ChamScript: // 47
		return "ChamScript"
	case QLocale__CopticScript: // 48
		return "CopticScript"
	case QLocale__CypriotScript: // 49
		return "CypriotScript"
	case QLocale__EgyptianHieroglyphsScript: // 50
		return "EgyptianHieroglyphsScript"
	case QLocale__FraserScript: // 51
		return "FraserScript"
	case QLocale__GlagoliticScript: // 52
		return "GlagoliticScript"
	case QLocale__GothicScript: // 53
		return "GothicScript"
	case QLocale__HanScript: // 54
		return "HanScript"
	case QLocale__HangulScript: // 55
		return "HangulScript"
	case QLocale__HanunooScript: // 56
		return "HanunooScript"
	case QLocale__ImperialAramaicScript: // 57
		return "ImperialAramaicScript"
	case QLocale__InscriptionalPahlaviScript: // 58
		return "InscriptionalPahlaviScript"
	case QLocale__InscriptionalParthianScript: // 59
		return "InscriptionalParthianScript"
	case QLocale__JavaneseScript: // 60
		return "JavaneseScript"
	case QLocale__KaithiScript: // 61
		return "KaithiScript"
	case QLocale__KatakanaScript: // 62
		return "KatakanaScript"
	case QLocale__KayahLiScript: // 63
		return "KayahLiScript"
	case QLocale__KharoshthiScript: // 64
		return "KharoshthiScript"
	case QLocale__LannaScript: // 65
		return "LannaScript"
	case QLocale__LepchaScript: // 66
		return "LepchaScript"
	case QLocale__LimbuScript: // 67
		return "LimbuScript"
	case QLocale__LinearBScript: // 68
		return "LinearBScript"
	case QLocale__LycianScript: // 69
		return "LycianScript"
	case QLocale__LydianScript: // 70
		return "LydianScript"
	case QLocale__MandaeanScript: // 71
		return "MandaeanScript"
	case QLocale__MeiteiMayekScript: // 72
		return "MeiteiMayekScript"
	case QLocale__MeroiticScript: // 73
		return "MeroiticScript"
	case QLocale__MeroiticCursiveScript: // 74
		return "MeroiticCursiveScript"
	case QLocale__NkoScript: // 75
		return "NkoScript"
	case QLocale__NewTaiLueScript: // 76
		return "NewTaiLueScript"
	case QLocale__OghamScript: // 77
		return "OghamScript"
	case QLocale__OlChikiScript: // 78
		return "OlChikiScript"
	case QLocale__OldItalicScript: // 79
		return "OldItalicScript"
	case QLocale__OldPersianScript: // 80
		return "OldPersianScript"
	case QLocale__OldSouthArabianScript: // 81
		return "OldSouthArabianScript"
	case QLocale__OrkhonScript: // 82
		return "OrkhonScript"
	case QLocale__OsmanyaScript: // 83
		return "OsmanyaScript"
	case QLocale__PhagsPaScript: // 84
		return "PhagsPaScript"
	case QLocale__PhoenicianScript: // 85
		return "PhoenicianScript"
	case QLocale__PollardPhoneticScript: // 86
		return "PollardPhoneticScript"
	case QLocale__RejangScript: // 87
		return "RejangScript"
	case QLocale__RunicScript: // 88
		return "RunicScript"
	case QLocale__SamaritanScript: // 89
		return "SamaritanScript"
	case QLocale__SaurashtraScript: // 90
		return "SaurashtraScript"
	case QLocale__SharadaScript: // 91
		return "SharadaScript"
	case QLocale__ShavianScript: // 92
		return "ShavianScript"
	case QLocale__SoraSompengScript: // 93
		return "SoraSompengScript"
	case QLocale__CuneiformScript: // 94
		return "CuneiformScript"
	case QLocale__SundaneseScript: // 95
		return "SundaneseScript"
	case QLocale__SylotiNagriScript: // 96
		return "SylotiNagriScript"
	case QLocale__TagalogScript: // 97
		return "TagalogScript"
	case QLocale__TagbanwaScript: // 98
		return "TagbanwaScript"
	case QLocale__TaiLeScript: // 99
		return "TaiLeScript"
	case QLocale__TaiVietScript: // 100
		return "TaiVietScript"
	case QLocale__TakriScript: // 101
		return "TakriScript"
	case QLocale__UgariticScript: // 102
		return "UgariticScript"
	case QLocale__BrailleScript: // 103
		return "BrailleScript"
	case QLocale__HiraganaScript: // 104
		return "HiraganaScript"
	case QLocale__CaucasianAlbanianScript: // 105
		return "CaucasianAlbanianScript"
	case QLocale__BassaVahScript: // 106
		return "BassaVahScript"
	case QLocale__DuployanScript: // 107
		return "DuployanScript"
	case QLocale__ElbasanScript: // 108
		return "ElbasanScript"
	case QLocale__GranthaScript: // 109
		return "GranthaScript"
	case QLocale__PahawhHmongScript: // 110
		return "PahawhHmongScript"
	case QLocale__KhojkiScript: // 111
		return "KhojkiScript"
	case QLocale__LinearAScript: // 112
		return "LinearAScript"
	case QLocale__MahajaniScript: // 113
		return "MahajaniScript"
	case QLocale__ManichaeanScript: // 114
		return "ManichaeanScript"
	case QLocale__MendeKikakuiScript: // 115
		return "MendeKikakuiScript"
	case QLocale__ModiScript: // 116
		return "ModiScript"
	case QLocale__MroScript: // 117
		return "MroScript"
	case QLocale__OldNorthArabianScript: // 118
		return "OldNorthArabianScript"
	case QLocale__NabataeanScript: // 119
		return "NabataeanScript"
	case QLocale__PalmyreneScript: // 120
		return "PalmyreneScript"
	case QLocale__PauCinHauScript: // 121
		return "PauCinHauScript"
	case QLocale__OldPermicScript: // 122
		return "OldPermicScript"
	case QLocale__PsalterPahlaviScript: // 123
		return "PsalterPahlaviScript"
	case QLocale__SiddhamScript: // 124
		return "SiddhamScript"
	case QLocale__KhudawadiScript: // 125
		return "KhudawadiScript"
	case QLocale__TirhutaScript: // 126
		return "TirhutaScript"
	case QLocale__VarangKshitiScript: // 127
		return "VarangKshitiScript"
	case QLocale__AhomScript: // 128
		return "AhomScript"
	case QLocale__AnatolianHieroglyphsScript: // 129
		return "AnatolianHieroglyphsScript"
	case QLocale__HatranScript: // 130
		return "HatranScript"
	case QLocale__MultaniScript: // 131
		return "MultaniScript"
	case QLocale__OldHungarianScript: // 132
		return "OldHungarianScript"
	case QLocale__SignWritingScript: // 133
		return "SignWritingScript"
	case QLocale__AdlamScript: // 134
		return "AdlamScript"
	case QLocale__BhaiksukiScript: // 135
		return "BhaiksukiScript"
	case QLocale__MarchenScript: // 136
		return "MarchenScript"
	case QLocale__NewaScript: // 137
		return "NewaScript"
	case QLocale__OsageScript: // 138
		return "OsageScript"
	case QLocale__TangutScript: // 139
		return "TangutScript"
	case QLocale__HanWithBopomofoScript: // 140
		return "HanWithBopomofoScript"
	case QLocale__JamoScript: // 141
		return "JamoScript,LastScript"
		// case QLocale__SimplifiedChineseScript: // 5
		// return ""
		// case QLocale__TraditionalChineseScript: // 6
		// return ""
		// case QLocale__LastScript: // 141
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_ScriptItemName(val int) string {
	var nilthis *QLocale
	return nilthis.ScriptItemName(val)
}

/*
This enumerated type is used to specify a country.

QLocale::DemocraticRepublicOfCongoCongoKinshasaObsolete, please use CongoKinshasa
QLocale::PeoplesRepublicOfCongoCongoBrazzavilleObsolete, please use CongoBrazzaville
QLocale::DemocraticRepublicOfKoreaNorthKoreaObsolete, please use NorthKorea
QLocale::RepublicOfKoreaSouthKoreaObsolete, please use SouthKorea
QLocale::RussianFederationRussiasame as Russia
QLocale::SyrianArabRepublicSyriaObsolete, please use Syria
QLocale::TokelauTokelauCountryObsolete, please use TokelauCountry
QLocale::TuvaluTuvaluCountryObsolete, please use TuvaluCountry


See also country() and countryToString().

*/
type QLocale__Country = int

//
const QLocale__AnyCountry QLocale__Country = 0

//
const QLocale__Afghanistan QLocale__Country = 1

//
const QLocale__Albania QLocale__Country = 2

//
const QLocale__Algeria QLocale__Country = 3

//
const QLocale__AmericanSamoa QLocale__Country = 4

//
const QLocale__Andorra QLocale__Country = 5

//
const QLocale__Angola QLocale__Country = 6

//
const QLocale__Anguilla QLocale__Country = 7

//
const QLocale__Antarctica QLocale__Country = 8

//
const QLocale__AntiguaAndBarbuda QLocale__Country = 9

//
const QLocale__Argentina QLocale__Country = 10

//
const QLocale__Armenia QLocale__Country = 11

//
const QLocale__Aruba QLocale__Country = 12

//
const QLocale__Australia QLocale__Country = 13

//
const QLocale__Austria QLocale__Country = 14

//
const QLocale__Azerbaijan QLocale__Country = 15

//
const QLocale__Bahamas QLocale__Country = 16

//
const QLocale__Bahrain QLocale__Country = 17

//
const QLocale__Bangladesh QLocale__Country = 18

//
const QLocale__Barbados QLocale__Country = 19

//
const QLocale__Belarus QLocale__Country = 20

//
const QLocale__Belgium QLocale__Country = 21

//
const QLocale__Belize QLocale__Country = 22

//
const QLocale__Benin QLocale__Country = 23

//
const QLocale__Bermuda QLocale__Country = 24

//
const QLocale__Bhutan QLocale__Country = 25

//
const QLocale__Bolivia QLocale__Country = 26

//
const QLocale__BosniaAndHerzegowina QLocale__Country = 27

//
const QLocale__Botswana QLocale__Country = 28

//
const QLocale__BouvetIsland QLocale__Country = 29

//
const QLocale__Brazil QLocale__Country = 30

//
const QLocale__BritishIndianOceanTerritory QLocale__Country = 31

//
const QLocale__Brunei QLocale__Country = 32

//
const QLocale__Bulgaria QLocale__Country = 33

//
const QLocale__BurkinaFaso QLocale__Country = 34

//
const QLocale__Burundi QLocale__Country = 35

//
const QLocale__Cambodia QLocale__Country = 36

//
const QLocale__Cameroon QLocale__Country = 37

//
const QLocale__Canada QLocale__Country = 38

//
const QLocale__CapeVerde QLocale__Country = 39

//
const QLocale__CaymanIslands QLocale__Country = 40

//
const QLocale__CentralAfricanRepublic QLocale__Country = 41

//
const QLocale__Chad QLocale__Country = 42

//
const QLocale__Chile QLocale__Country = 43

//
const QLocale__China QLocale__Country = 44

//
const QLocale__ChristmasIsland QLocale__Country = 45

//
const QLocale__CocosIslands QLocale__Country = 46

//
const QLocale__Colombia QLocale__Country = 47

//
const QLocale__Comoros QLocale__Country = 48

//
const QLocale__CongoKinshasa QLocale__Country = 49

//
const QLocale__CongoBrazzaville QLocale__Country = 50

//
const QLocale__CookIslands QLocale__Country = 51

//
const QLocale__CostaRica QLocale__Country = 52

//
const QLocale__IvoryCoast QLocale__Country = 53

//
const QLocale__Croatia QLocale__Country = 54

//
const QLocale__Cuba QLocale__Country = 55

//
const QLocale__Cyprus QLocale__Country = 56

//
const QLocale__CzechRepublic QLocale__Country = 57

//
const QLocale__Denmark QLocale__Country = 58

//
const QLocale__Djibouti QLocale__Country = 59

//
const QLocale__Dominica QLocale__Country = 60

//
const QLocale__DominicanRepublic QLocale__Country = 61

//
const QLocale__EastTimor QLocale__Country = 62

//
const QLocale__Ecuador QLocale__Country = 63

//
const QLocale__Egypt QLocale__Country = 64

//
const QLocale__ElSalvador QLocale__Country = 65

//
const QLocale__EquatorialGuinea QLocale__Country = 66

//
const QLocale__Eritrea QLocale__Country = 67

//
const QLocale__Estonia QLocale__Country = 68

//
const QLocale__Ethiopia QLocale__Country = 69

//
const QLocale__FalklandIslands QLocale__Country = 70

//
const QLocale__FaroeIslands QLocale__Country = 71

//
const QLocale__Fiji QLocale__Country = 72

//
const QLocale__Finland QLocale__Country = 73

//
const QLocale__France QLocale__Country = 74

//
const QLocale__Guernsey QLocale__Country = 75

//
const QLocale__FrenchGuiana QLocale__Country = 76

//
const QLocale__FrenchPolynesia QLocale__Country = 77

//
const QLocale__FrenchSouthernTerritories QLocale__Country = 78

//
const QLocale__Gabon QLocale__Country = 79

//
const QLocale__Gambia QLocale__Country = 80

//
const QLocale__Georgia QLocale__Country = 81

//
const QLocale__Germany QLocale__Country = 82

//
const QLocale__Ghana QLocale__Country = 83

//
const QLocale__Gibraltar QLocale__Country = 84

//
const QLocale__Greece QLocale__Country = 85

//
const QLocale__Greenland QLocale__Country = 86

//
const QLocale__Grenada QLocale__Country = 87

//
const QLocale__Guadeloupe QLocale__Country = 88

//
const QLocale__Guam QLocale__Country = 89

//
const QLocale__Guatemala QLocale__Country = 90

//
const QLocale__Guinea QLocale__Country = 91

//
const QLocale__GuineaBissau QLocale__Country = 92

//
const QLocale__Guyana QLocale__Country = 93

//
const QLocale__Haiti QLocale__Country = 94

//
const QLocale__HeardAndMcDonaldIslands QLocale__Country = 95

//
const QLocale__Honduras QLocale__Country = 96

//
const QLocale__HongKong QLocale__Country = 97

//
const QLocale__Hungary QLocale__Country = 98

//
const QLocale__Iceland QLocale__Country = 99

//
const QLocale__India QLocale__Country = 100

//
const QLocale__Indonesia QLocale__Country = 101

//
const QLocale__Iran QLocale__Country = 102

//
const QLocale__Iraq QLocale__Country = 103

//
const QLocale__Ireland QLocale__Country = 104

//
const QLocale__Israel QLocale__Country = 105

//
const QLocale__Italy QLocale__Country = 106

//
const QLocale__Jamaica QLocale__Country = 107

//
const QLocale__Japan QLocale__Country = 108

//
const QLocale__Jordan QLocale__Country = 109

//
const QLocale__Kazakhstan QLocale__Country = 110

//
const QLocale__Kenya QLocale__Country = 111

//
const QLocale__Kiribati QLocale__Country = 112

//
const QLocale__NorthKorea QLocale__Country = 113

//
const QLocale__SouthKorea QLocale__Country = 114

//
const QLocale__Kuwait QLocale__Country = 115

//
const QLocale__Kyrgyzstan QLocale__Country = 116

//
const QLocale__Laos QLocale__Country = 117

//
const QLocale__Latvia QLocale__Country = 118

//
const QLocale__Lebanon QLocale__Country = 119

//
const QLocale__Lesotho QLocale__Country = 120

//
const QLocale__Liberia QLocale__Country = 121

//
const QLocale__Libya QLocale__Country = 122

//
const QLocale__Liechtenstein QLocale__Country = 123

//
const QLocale__Lithuania QLocale__Country = 124

//
const QLocale__Luxembourg QLocale__Country = 125

//
const QLocale__Macau QLocale__Country = 126

//
const QLocale__Macedonia QLocale__Country = 127

//
const QLocale__Madagascar QLocale__Country = 128

//
const QLocale__Malawi QLocale__Country = 129

//
const QLocale__Malaysia QLocale__Country = 130

//
const QLocale__Maldives QLocale__Country = 131

//
const QLocale__Mali QLocale__Country = 132

//
const QLocale__Malta QLocale__Country = 133

//
const QLocale__MarshallIslands QLocale__Country = 134

//
const QLocale__Martinique QLocale__Country = 135

//
const QLocale__Mauritania QLocale__Country = 136

//
const QLocale__Mauritius QLocale__Country = 137

//
const QLocale__Mayotte QLocale__Country = 138

//
const QLocale__Mexico QLocale__Country = 139

//
const QLocale__Micronesia QLocale__Country = 140

//
const QLocale__Moldova QLocale__Country = 141

//
const QLocale__Monaco QLocale__Country = 142

//
const QLocale__Mongolia QLocale__Country = 143

//
const QLocale__Montserrat QLocale__Country = 144

//
const QLocale__Morocco QLocale__Country = 145

//
const QLocale__Mozambique QLocale__Country = 146

//
const QLocale__Myanmar QLocale__Country = 147

//
const QLocale__Namibia QLocale__Country = 148

//
const QLocale__NauruCountry QLocale__Country = 149

//
const QLocale__Nepal QLocale__Country = 150

//
const QLocale__Netherlands QLocale__Country = 151

//
const QLocale__CuraSao QLocale__Country = 152

//
const QLocale__NewCaledonia QLocale__Country = 153

//
const QLocale__NewZealand QLocale__Country = 154

//
const QLocale__Nicaragua QLocale__Country = 155

//
const QLocale__Niger QLocale__Country = 156

//
const QLocale__Nigeria QLocale__Country = 157

//
const QLocale__Niue QLocale__Country = 158

//
const QLocale__NorfolkIsland QLocale__Country = 159

//
const QLocale__NorthernMarianaIslands QLocale__Country = 160

//
const QLocale__Norway QLocale__Country = 161

//
const QLocale__Oman QLocale__Country = 162

//
const QLocale__Pakistan QLocale__Country = 163

//
const QLocale__Palau QLocale__Country = 164

//
const QLocale__PalestinianTerritories QLocale__Country = 165

//
const QLocale__Panama QLocale__Country = 166

//
const QLocale__PapuaNewGuinea QLocale__Country = 167

//
const QLocale__Paraguay QLocale__Country = 168

//
const QLocale__Peru QLocale__Country = 169

//
const QLocale__Philippines QLocale__Country = 170

//
const QLocale__Pitcairn QLocale__Country = 171

//
const QLocale__Poland QLocale__Country = 172

//
const QLocale__Portugal QLocale__Country = 173

//
const QLocale__PuertoRico QLocale__Country = 174

//
const QLocale__Qatar QLocale__Country = 175

//
const QLocale__Reunion QLocale__Country = 176

//
const QLocale__Romania QLocale__Country = 177

//
const QLocale__Russia QLocale__Country = 178

//
const QLocale__Rwanda QLocale__Country = 179

//
const QLocale__SaintKittsAndNevis QLocale__Country = 180

//
const QLocale__SaintLucia QLocale__Country = 181

//
const QLocale__SaintVincentAndTheGrenadines QLocale__Country = 182

//
const QLocale__Samoa QLocale__Country = 183

//
const QLocale__SanMarino QLocale__Country = 184

//
const QLocale__SaoTomeAndPrincipe QLocale__Country = 185

//
const QLocale__SaudiArabia QLocale__Country = 186

//
const QLocale__Senegal QLocale__Country = 187

//
const QLocale__Seychelles QLocale__Country = 188

//
const QLocale__SierraLeone QLocale__Country = 189

//
const QLocale__Singapore QLocale__Country = 190

//
const QLocale__Slovakia QLocale__Country = 191

//
const QLocale__Slovenia QLocale__Country = 192

//
const QLocale__SolomonIslands QLocale__Country = 193

//
const QLocale__Somalia QLocale__Country = 194

//
const QLocale__SouthAfrica QLocale__Country = 195

//
const QLocale__SouthGeorgiaAndTheSouthSandwichIslands QLocale__Country = 196

//
const QLocale__Spain QLocale__Country = 197

//
const QLocale__SriLanka QLocale__Country = 198

//
const QLocale__SaintHelena QLocale__Country = 199

//
const QLocale__SaintPierreAndMiquelon QLocale__Country = 200

//
const QLocale__Sudan QLocale__Country = 201

//
const QLocale__Suriname QLocale__Country = 202

//
const QLocale__SvalbardAndJanMayenIslands QLocale__Country = 203

//
const QLocale__Swaziland QLocale__Country = 204

//
const QLocale__Sweden QLocale__Country = 205

//
const QLocale__Switzerland QLocale__Country = 206

//
const QLocale__Syria QLocale__Country = 207

//
const QLocale__Taiwan QLocale__Country = 208

//
const QLocale__Tajikistan QLocale__Country = 209

//
const QLocale__Tanzania QLocale__Country = 210

//
const QLocale__Thailand QLocale__Country = 211

//
const QLocale__Togo QLocale__Country = 212

//
const QLocale__TokelauCountry QLocale__Country = 213

//
const QLocale__Tonga QLocale__Country = 214

//
const QLocale__TrinidadAndTobago QLocale__Country = 215

//
const QLocale__Tunisia QLocale__Country = 216

//
const QLocale__Turkey QLocale__Country = 217

//
const QLocale__Turkmenistan QLocale__Country = 218

//
const QLocale__TurksAndCaicosIslands QLocale__Country = 219

//
const QLocale__TuvaluCountry QLocale__Country = 220

//
const QLocale__Uganda QLocale__Country = 221

//
const QLocale__Ukraine QLocale__Country = 222

//
const QLocale__UnitedArabEmirates QLocale__Country = 223

//
const QLocale__UnitedKingdom QLocale__Country = 224

//
const QLocale__UnitedStates QLocale__Country = 225

//
const QLocale__UnitedStatesMinorOutlyingIslands QLocale__Country = 226

//
const QLocale__Uruguay QLocale__Country = 227

//
const QLocale__Uzbekistan QLocale__Country = 228

//
const QLocale__Vanuatu QLocale__Country = 229

//
const QLocale__VaticanCityState QLocale__Country = 230

//
const QLocale__Venezuela QLocale__Country = 231

//
const QLocale__Vietnam QLocale__Country = 232

//
const QLocale__BritishVirginIslands QLocale__Country = 233

//
const QLocale__UnitedStatesVirginIslands QLocale__Country = 234

//
const QLocale__WallisAndFutunaIslands QLocale__Country = 235

//
const QLocale__WesternSahara QLocale__Country = 236

//
const QLocale__Yemen QLocale__Country = 237

//
const QLocale__CanaryIslands QLocale__Country = 238

//
const QLocale__Zambia QLocale__Country = 239

//
const QLocale__Zimbabwe QLocale__Country = 240

//
const QLocale__ClippertonIsland QLocale__Country = 241

//
const QLocale__Montenegro QLocale__Country = 242

//
const QLocale__Serbia QLocale__Country = 243

//
const QLocale__SaintBarthelemy QLocale__Country = 244

//
const QLocale__SaintMartin QLocale__Country = 245

//
const QLocale__LatinAmericaAndTheCaribbean QLocale__Country = 246

//
const QLocale__AscensionIsland QLocale__Country = 247

//
const QLocale__AlandIslands QLocale__Country = 248

//
const QLocale__DiegoGarcia QLocale__Country = 249

//
const QLocale__CeutaAndMelilla QLocale__Country = 250

//
const QLocale__IsleOfMan QLocale__Country = 251

//
const QLocale__Jersey QLocale__Country = 252

//
const QLocale__TristanDaCunha QLocale__Country = 253

//
const QLocale__SouthSudan QLocale__Country = 254

//
const QLocale__Bonaire QLocale__Country = 255

//
const QLocale__SintMaarten QLocale__Country = 256

//
const QLocale__Kosovo QLocale__Country = 257

//
const QLocale__EuropeanUnion QLocale__Country = 258

//
const QLocale__OutlyingOceania QLocale__Country = 259

//
const QLocale__Tokelau QLocale__Country = 213

//
const QLocale__Tuvalu QLocale__Country = 220

//
const QLocale__DemocraticRepublicOfCongo QLocale__Country = 49

//
const QLocale__PeoplesRepublicOfCongo QLocale__Country = 50

//
const QLocale__DemocraticRepublicOfKorea QLocale__Country = 113

//
const QLocale__RepublicOfKorea QLocale__Country = 114

//
const QLocale__RussianFederation QLocale__Country = 178

//
const QLocale__SyrianArabRepublic QLocale__Country = 207

//
const QLocale__LastCountry QLocale__Country = 259

func (this *QLocale) CountryItemName(val int) string {
	switch val {
	case QLocale__AnyCountry: // 0
		return "AnyCountry"
	case QLocale__Afghanistan: // 1
		return "Afghanistan"
	case QLocale__Albania: // 2
		return "Albania"
	case QLocale__Algeria: // 3
		return "Algeria"
	case QLocale__AmericanSamoa: // 4
		return "AmericanSamoa"
	case QLocale__Andorra: // 5
		return "Andorra"
	case QLocale__Angola: // 6
		return "Angola"
	case QLocale__Anguilla: // 7
		return "Anguilla"
	case QLocale__Antarctica: // 8
		return "Antarctica"
	case QLocale__AntiguaAndBarbuda: // 9
		return "AntiguaAndBarbuda"
	case QLocale__Argentina: // 10
		return "Argentina"
	case QLocale__Armenia: // 11
		return "Armenia"
	case QLocale__Aruba: // 12
		return "Aruba"
	case QLocale__Australia: // 13
		return "Australia"
	case QLocale__Austria: // 14
		return "Austria"
	case QLocale__Azerbaijan: // 15
		return "Azerbaijan"
	case QLocale__Bahamas: // 16
		return "Bahamas"
	case QLocale__Bahrain: // 17
		return "Bahrain"
	case QLocale__Bangladesh: // 18
		return "Bangladesh"
	case QLocale__Barbados: // 19
		return "Barbados"
	case QLocale__Belarus: // 20
		return "Belarus"
	case QLocale__Belgium: // 21
		return "Belgium"
	case QLocale__Belize: // 22
		return "Belize"
	case QLocale__Benin: // 23
		return "Benin"
	case QLocale__Bermuda: // 24
		return "Bermuda"
	case QLocale__Bhutan: // 25
		return "Bhutan"
	case QLocale__Bolivia: // 26
		return "Bolivia"
	case QLocale__BosniaAndHerzegowina: // 27
		return "BosniaAndHerzegowina"
	case QLocale__Botswana: // 28
		return "Botswana"
	case QLocale__BouvetIsland: // 29
		return "BouvetIsland"
	case QLocale__Brazil: // 30
		return "Brazil"
	case QLocale__BritishIndianOceanTerritory: // 31
		return "BritishIndianOceanTerritory"
	case QLocale__Brunei: // 32
		return "Brunei"
	case QLocale__Bulgaria: // 33
		return "Bulgaria"
	case QLocale__BurkinaFaso: // 34
		return "BurkinaFaso"
	case QLocale__Burundi: // 35
		return "Burundi"
	case QLocale__Cambodia: // 36
		return "Cambodia"
	case QLocale__Cameroon: // 37
		return "Cameroon"
	case QLocale__Canada: // 38
		return "Canada"
	case QLocale__CapeVerde: // 39
		return "CapeVerde"
	case QLocale__CaymanIslands: // 40
		return "CaymanIslands"
	case QLocale__CentralAfricanRepublic: // 41
		return "CentralAfricanRepublic"
	case QLocale__Chad: // 42
		return "Chad"
	case QLocale__Chile: // 43
		return "Chile"
	case QLocale__China: // 44
		return "China"
	case QLocale__ChristmasIsland: // 45
		return "ChristmasIsland"
	case QLocale__CocosIslands: // 46
		return "CocosIslands"
	case QLocale__Colombia: // 47
		return "Colombia"
	case QLocale__Comoros: // 48
		return "Comoros"
	case QLocale__CongoKinshasa: // 49
		return "CongoKinshasa,DemocraticRepublicOfCongo"
	case QLocale__CongoBrazzaville: // 50
		return "CongoBrazzaville,PeoplesRepublicOfCongo"
	case QLocale__CookIslands: // 51
		return "CookIslands"
	case QLocale__CostaRica: // 52
		return "CostaRica"
	case QLocale__IvoryCoast: // 53
		return "IvoryCoast"
	case QLocale__Croatia: // 54
		return "Croatia"
	case QLocale__Cuba: // 55
		return "Cuba"
	case QLocale__Cyprus: // 56
		return "Cyprus"
	case QLocale__CzechRepublic: // 57
		return "CzechRepublic"
	case QLocale__Denmark: // 58
		return "Denmark"
	case QLocale__Djibouti: // 59
		return "Djibouti"
	case QLocale__Dominica: // 60
		return "Dominica"
	case QLocale__DominicanRepublic: // 61
		return "DominicanRepublic"
	case QLocale__EastTimor: // 62
		return "EastTimor"
	case QLocale__Ecuador: // 63
		return "Ecuador"
	case QLocale__Egypt: // 64
		return "Egypt"
	case QLocale__ElSalvador: // 65
		return "ElSalvador"
	case QLocale__EquatorialGuinea: // 66
		return "EquatorialGuinea"
	case QLocale__Eritrea: // 67
		return "Eritrea"
	case QLocale__Estonia: // 68
		return "Estonia"
	case QLocale__Ethiopia: // 69
		return "Ethiopia"
	case QLocale__FalklandIslands: // 70
		return "FalklandIslands"
	case QLocale__FaroeIslands: // 71
		return "FaroeIslands"
	case QLocale__Fiji: // 72
		return "Fiji"
	case QLocale__Finland: // 73
		return "Finland"
	case QLocale__France: // 74
		return "France"
	case QLocale__Guernsey: // 75
		return "Guernsey"
	case QLocale__FrenchGuiana: // 76
		return "FrenchGuiana"
	case QLocale__FrenchPolynesia: // 77
		return "FrenchPolynesia"
	case QLocale__FrenchSouthernTerritories: // 78
		return "FrenchSouthernTerritories"
	case QLocale__Gabon: // 79
		return "Gabon"
	case QLocale__Gambia: // 80
		return "Gambia"
	case QLocale__Georgia: // 81
		return "Georgia"
	case QLocale__Germany: // 82
		return "Germany"
	case QLocale__Ghana: // 83
		return "Ghana"
	case QLocale__Gibraltar: // 84
		return "Gibraltar"
	case QLocale__Greece: // 85
		return "Greece"
	case QLocale__Greenland: // 86
		return "Greenland"
	case QLocale__Grenada: // 87
		return "Grenada"
	case QLocale__Guadeloupe: // 88
		return "Guadeloupe"
	case QLocale__Guam: // 89
		return "Guam"
	case QLocale__Guatemala: // 90
		return "Guatemala"
	case QLocale__Guinea: // 91
		return "Guinea"
	case QLocale__GuineaBissau: // 92
		return "GuineaBissau"
	case QLocale__Guyana: // 93
		return "Guyana"
	case QLocale__Haiti: // 94
		return "Haiti"
	case QLocale__HeardAndMcDonaldIslands: // 95
		return "HeardAndMcDonaldIslands"
	case QLocale__Honduras: // 96
		return "Honduras"
	case QLocale__HongKong: // 97
		return "HongKong"
	case QLocale__Hungary: // 98
		return "Hungary"
	case QLocale__Iceland: // 99
		return "Iceland"
	case QLocale__India: // 100
		return "India"
	case QLocale__Indonesia: // 101
		return "Indonesia"
	case QLocale__Iran: // 102
		return "Iran"
	case QLocale__Iraq: // 103
		return "Iraq"
	case QLocale__Ireland: // 104
		return "Ireland"
	case QLocale__Israel: // 105
		return "Israel"
	case QLocale__Italy: // 106
		return "Italy"
	case QLocale__Jamaica: // 107
		return "Jamaica"
	case QLocale__Japan: // 108
		return "Japan"
	case QLocale__Jordan: // 109
		return "Jordan"
	case QLocale__Kazakhstan: // 110
		return "Kazakhstan"
	case QLocale__Kenya: // 111
		return "Kenya"
	case QLocale__Kiribati: // 112
		return "Kiribati"
	case QLocale__NorthKorea: // 113
		return "NorthKorea,DemocraticRepublicOfKorea"
	case QLocale__SouthKorea: // 114
		return "SouthKorea,RepublicOfKorea"
	case QLocale__Kuwait: // 115
		return "Kuwait"
	case QLocale__Kyrgyzstan: // 116
		return "Kyrgyzstan"
	case QLocale__Laos: // 117
		return "Laos"
	case QLocale__Latvia: // 118
		return "Latvia"
	case QLocale__Lebanon: // 119
		return "Lebanon"
	case QLocale__Lesotho: // 120
		return "Lesotho"
	case QLocale__Liberia: // 121
		return "Liberia"
	case QLocale__Libya: // 122
		return "Libya"
	case QLocale__Liechtenstein: // 123
		return "Liechtenstein"
	case QLocale__Lithuania: // 124
		return "Lithuania"
	case QLocale__Luxembourg: // 125
		return "Luxembourg"
	case QLocale__Macau: // 126
		return "Macau"
	case QLocale__Macedonia: // 127
		return "Macedonia"
	case QLocale__Madagascar: // 128
		return "Madagascar"
	case QLocale__Malawi: // 129
		return "Malawi"
	case QLocale__Malaysia: // 130
		return "Malaysia"
	case QLocale__Maldives: // 131
		return "Maldives"
	case QLocale__Mali: // 132
		return "Mali"
	case QLocale__Malta: // 133
		return "Malta"
	case QLocale__MarshallIslands: // 134
		return "MarshallIslands"
	case QLocale__Martinique: // 135
		return "Martinique"
	case QLocale__Mauritania: // 136
		return "Mauritania"
	case QLocale__Mauritius: // 137
		return "Mauritius"
	case QLocale__Mayotte: // 138
		return "Mayotte"
	case QLocale__Mexico: // 139
		return "Mexico"
	case QLocale__Micronesia: // 140
		return "Micronesia"
	case QLocale__Moldova: // 141
		return "Moldova"
	case QLocale__Monaco: // 142
		return "Monaco"
	case QLocale__Mongolia: // 143
		return "Mongolia"
	case QLocale__Montserrat: // 144
		return "Montserrat"
	case QLocale__Morocco: // 145
		return "Morocco"
	case QLocale__Mozambique: // 146
		return "Mozambique"
	case QLocale__Myanmar: // 147
		return "Myanmar"
	case QLocale__Namibia: // 148
		return "Namibia"
	case QLocale__NauruCountry: // 149
		return "NauruCountry"
	case QLocale__Nepal: // 150
		return "Nepal"
	case QLocale__Netherlands: // 151
		return "Netherlands"
	case QLocale__CuraSao: // 152
		return "CuraSao"
	case QLocale__NewCaledonia: // 153
		return "NewCaledonia"
	case QLocale__NewZealand: // 154
		return "NewZealand"
	case QLocale__Nicaragua: // 155
		return "Nicaragua"
	case QLocale__Niger: // 156
		return "Niger"
	case QLocale__Nigeria: // 157
		return "Nigeria"
	case QLocale__Niue: // 158
		return "Niue"
	case QLocale__NorfolkIsland: // 159
		return "NorfolkIsland"
	case QLocale__NorthernMarianaIslands: // 160
		return "NorthernMarianaIslands"
	case QLocale__Norway: // 161
		return "Norway"
	case QLocale__Oman: // 162
		return "Oman"
	case QLocale__Pakistan: // 163
		return "Pakistan"
	case QLocale__Palau: // 164
		return "Palau"
	case QLocale__PalestinianTerritories: // 165
		return "PalestinianTerritories"
	case QLocale__Panama: // 166
		return "Panama"
	case QLocale__PapuaNewGuinea: // 167
		return "PapuaNewGuinea"
	case QLocale__Paraguay: // 168
		return "Paraguay"
	case QLocale__Peru: // 169
		return "Peru"
	case QLocale__Philippines: // 170
		return "Philippines"
	case QLocale__Pitcairn: // 171
		return "Pitcairn"
	case QLocale__Poland: // 172
		return "Poland"
	case QLocale__Portugal: // 173
		return "Portugal"
	case QLocale__PuertoRico: // 174
		return "PuertoRico"
	case QLocale__Qatar: // 175
		return "Qatar"
	case QLocale__Reunion: // 176
		return "Reunion"
	case QLocale__Romania: // 177
		return "Romania"
	case QLocale__Russia: // 178
		return "Russia,RussianFederation"
	case QLocale__Rwanda: // 179
		return "Rwanda"
	case QLocale__SaintKittsAndNevis: // 180
		return "SaintKittsAndNevis"
	case QLocale__SaintLucia: // 181
		return "SaintLucia"
	case QLocale__SaintVincentAndTheGrenadines: // 182
		return "SaintVincentAndTheGrenadines"
	case QLocale__Samoa: // 183
		return "Samoa"
	case QLocale__SanMarino: // 184
		return "SanMarino"
	case QLocale__SaoTomeAndPrincipe: // 185
		return "SaoTomeAndPrincipe"
	case QLocale__SaudiArabia: // 186
		return "SaudiArabia"
	case QLocale__Senegal: // 187
		return "Senegal"
	case QLocale__Seychelles: // 188
		return "Seychelles"
	case QLocale__SierraLeone: // 189
		return "SierraLeone"
	case QLocale__Singapore: // 190
		return "Singapore"
	case QLocale__Slovakia: // 191
		return "Slovakia"
	case QLocale__Slovenia: // 192
		return "Slovenia"
	case QLocale__SolomonIslands: // 193
		return "SolomonIslands"
	case QLocale__Somalia: // 194
		return "Somalia"
	case QLocale__SouthAfrica: // 195
		return "SouthAfrica"
	case QLocale__SouthGeorgiaAndTheSouthSandwichIslands: // 196
		return "SouthGeorgiaAndTheSouthSandwichIslands"
	case QLocale__Spain: // 197
		return "Spain"
	case QLocale__SriLanka: // 198
		return "SriLanka"
	case QLocale__SaintHelena: // 199
		return "SaintHelena"
	case QLocale__SaintPierreAndMiquelon: // 200
		return "SaintPierreAndMiquelon"
	case QLocale__Sudan: // 201
		return "Sudan"
	case QLocale__Suriname: // 202
		return "Suriname"
	case QLocale__SvalbardAndJanMayenIslands: // 203
		return "SvalbardAndJanMayenIslands"
	case QLocale__Swaziland: // 204
		return "Swaziland"
	case QLocale__Sweden: // 205
		return "Sweden"
	case QLocale__Switzerland: // 206
		return "Switzerland"
	case QLocale__Syria: // 207
		return "Syria,SyrianArabRepublic"
	case QLocale__Taiwan: // 208
		return "Taiwan"
	case QLocale__Tajikistan: // 209
		return "Tajikistan"
	case QLocale__Tanzania: // 210
		return "Tanzania"
	case QLocale__Thailand: // 211
		return "Thailand"
	case QLocale__Togo: // 212
		return "Togo"
	case QLocale__TokelauCountry: // 213
		return "TokelauCountry,Tokelau"
	case QLocale__Tonga: // 214
		return "Tonga"
	case QLocale__TrinidadAndTobago: // 215
		return "TrinidadAndTobago"
	case QLocale__Tunisia: // 216
		return "Tunisia"
	case QLocale__Turkey: // 217
		return "Turkey"
	case QLocale__Turkmenistan: // 218
		return "Turkmenistan"
	case QLocale__TurksAndCaicosIslands: // 219
		return "TurksAndCaicosIslands"
	case QLocale__TuvaluCountry: // 220
		return "TuvaluCountry,Tuvalu"
	case QLocale__Uganda: // 221
		return "Uganda"
	case QLocale__Ukraine: // 222
		return "Ukraine"
	case QLocale__UnitedArabEmirates: // 223
		return "UnitedArabEmirates"
	case QLocale__UnitedKingdom: // 224
		return "UnitedKingdom"
	case QLocale__UnitedStates: // 225
		return "UnitedStates"
	case QLocale__UnitedStatesMinorOutlyingIslands: // 226
		return "UnitedStatesMinorOutlyingIslands"
	case QLocale__Uruguay: // 227
		return "Uruguay"
	case QLocale__Uzbekistan: // 228
		return "Uzbekistan"
	case QLocale__Vanuatu: // 229
		return "Vanuatu"
	case QLocale__VaticanCityState: // 230
		return "VaticanCityState"
	case QLocale__Venezuela: // 231
		return "Venezuela"
	case QLocale__Vietnam: // 232
		return "Vietnam"
	case QLocale__BritishVirginIslands: // 233
		return "BritishVirginIslands"
	case QLocale__UnitedStatesVirginIslands: // 234
		return "UnitedStatesVirginIslands"
	case QLocale__WallisAndFutunaIslands: // 235
		return "WallisAndFutunaIslands"
	case QLocale__WesternSahara: // 236
		return "WesternSahara"
	case QLocale__Yemen: // 237
		return "Yemen"
	case QLocale__CanaryIslands: // 238
		return "CanaryIslands"
	case QLocale__Zambia: // 239
		return "Zambia"
	case QLocale__Zimbabwe: // 240
		return "Zimbabwe"
	case QLocale__ClippertonIsland: // 241
		return "ClippertonIsland"
	case QLocale__Montenegro: // 242
		return "Montenegro"
	case QLocale__Serbia: // 243
		return "Serbia"
	case QLocale__SaintBarthelemy: // 244
		return "SaintBarthelemy"
	case QLocale__SaintMartin: // 245
		return "SaintMartin"
	case QLocale__LatinAmericaAndTheCaribbean: // 246
		return "LatinAmericaAndTheCaribbean"
	case QLocale__AscensionIsland: // 247
		return "AscensionIsland"
	case QLocale__AlandIslands: // 248
		return "AlandIslands"
	case QLocale__DiegoGarcia: // 249
		return "DiegoGarcia"
	case QLocale__CeutaAndMelilla: // 250
		return "CeutaAndMelilla"
	case QLocale__IsleOfMan: // 251
		return "IsleOfMan"
	case QLocale__Jersey: // 252
		return "Jersey"
	case QLocale__TristanDaCunha: // 253
		return "TristanDaCunha"
	case QLocale__SouthSudan: // 254
		return "SouthSudan"
	case QLocale__Bonaire: // 255
		return "Bonaire"
	case QLocale__SintMaarten: // 256
		return "SintMaarten"
	case QLocale__Kosovo: // 257
		return "Kosovo"
	case QLocale__EuropeanUnion: // 258
		return "EuropeanUnion"
	case QLocale__OutlyingOceania: // 259
		return "OutlyingOceania,LastCountry"
		// case QLocale__Tokelau: // 213
		// return ""
		// case QLocale__Tuvalu: // 220
		// return ""
		// case QLocale__DemocraticRepublicOfCongo: // 49
		// return ""
		// case QLocale__PeoplesRepublicOfCongo: // 50
		// return ""
		// case QLocale__DemocraticRepublicOfKorea: // 113
		// return ""
		// case QLocale__RepublicOfKorea: // 114
		// return ""
		// case QLocale__RussianFederation: // 178
		// return ""
		// case QLocale__SyrianArabRepublic: // 207
		// return ""
		// case QLocale__LastCountry: // 259
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_CountryItemName(val int) string {
	var nilthis *QLocale
	return nilthis.CountryItemName(val)
}

/*
This enum defines which units are used for measurement.

QLocale::ImperialSystemImperialUSSystemProvided for compatibility. Same as ImperialUSSystem


This enum was introduced or modified in  Qt 4.4.

*/
type QLocale__MeasurementSystem = int

// This value indicates metric units, such as meters, centimeters and millimeters.
const QLocale__MetricSystem QLocale__MeasurementSystem = 0

// This value indicates imperial units, such as inches and miles as they are used in the United States.
const QLocale__ImperialUSSystem QLocale__MeasurementSystem = 1

// This value indicates imperial units, such as inches and miles as they are used in the United Kingdom.
const QLocale__ImperialUKSystem QLocale__MeasurementSystem = 2

//
const QLocale__ImperialSystem QLocale__MeasurementSystem = 1

func (this *QLocale) MeasurementSystemItemName(val int) string {
	switch val {
	case QLocale__MetricSystem: // 0
		return "MetricSystem"
	case QLocale__ImperialUSSystem: // 1
		return "ImperialUSSystem,ImperialSystem"
	case QLocale__ImperialUKSystem: // 2
		return "ImperialUKSystem"
		// case QLocale__ImperialSystem: // 1
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_MeasurementSystemItemName(val int) string {
	var nilthis *QLocale
	return nilthis.MeasurementSystemItemName(val)
}

/*
This enum describes the types of format that can be used when converting QDate and QTime objects to strings.


*/
type QLocale__FormatType = int

// The long version of day and month names; for example, returning "January" as a month name.
const QLocale__LongFormat QLocale__FormatType = 0

// The short version of day and month names; for example, returning "Jan" as a month name.
const QLocale__ShortFormat QLocale__FormatType = 1

// A special version of day and month names for use when space is limited; for example, returning "J" as a month name. Note that the narrow format might contain the same text for different months and days or it can even be an empty string if the locale doesn't support narrow names, so you should avoid using it for date formatting. Also, for the system locale this format is the same as ShortFormat.
const QLocale__NarrowFormat QLocale__FormatType = 2

func (this *QLocale) FormatTypeItemName(val int) string {
	switch val {
	case QLocale__LongFormat: // 0
		return "LongFormat"
	case QLocale__ShortFormat: // 1
		return "ShortFormat"
	case QLocale__NarrowFormat: // 2
		return "NarrowFormat"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_FormatTypeItemName(val int) string {
	var nilthis *QLocale
	return nilthis.FormatTypeItemName(val)
}

/*


 */
type QLocale__NumberOption = int

//
const QLocale__DefaultNumberOptions QLocale__NumberOption = 0

//
const QLocale__OmitGroupSeparator QLocale__NumberOption = 1

//
const QLocale__RejectGroupSeparator QLocale__NumberOption = 2

//
const QLocale__OmitLeadingZeroInExponent QLocale__NumberOption = 4

//
const QLocale__RejectLeadingZeroInExponent QLocale__NumberOption = 8

//
const QLocale__IncludeTrailingZeroesAfterDot QLocale__NumberOption = 16

//
const QLocale__RejectTrailingZeroesAfterDot QLocale__NumberOption = 32

func (this *QLocale) NumberOptionItemName(val int) string {
	switch val {
	case QLocale__DefaultNumberOptions: // 0
		return "DefaultNumberOptions"
	case QLocale__OmitGroupSeparator: // 1
		return "OmitGroupSeparator"
	case QLocale__RejectGroupSeparator: // 2
		return "RejectGroupSeparator"
	case QLocale__OmitLeadingZeroInExponent: // 4
		return "OmitLeadingZeroInExponent"
	case QLocale__RejectLeadingZeroInExponent: // 8
		return "RejectLeadingZeroInExponent"
	case QLocale__IncludeTrailingZeroesAfterDot: // 16
		return "IncludeTrailingZeroesAfterDot"
	case QLocale__RejectTrailingZeroesAfterDot: // 32
		return "RejectTrailingZeroesAfterDot"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_NumberOptionItemName(val int) string {
	var nilthis *QLocale
	return nilthis.NumberOptionItemName(val)
}

/*
This enum defines constants that can be given as precision to QString::number(), QByteArray::number(), and QLocale::toString() when converting floats or doubles, in order to express a variable number of digits as precision.



This enum was introduced or modified in  Qt 5.7.

See also toString(), QString, and QByteArray.

*/
type QLocale__FloatingPointPrecisionOption = int

//
const QLocale__FloatingPointShortest QLocale__FloatingPointPrecisionOption = -128

func (this *QLocale) FloatingPointPrecisionOptionItemName(val int) string {
	switch val {
	case QLocale__FloatingPointShortest: // -128
		return "FloatingPointShortest"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_FloatingPointPrecisionOptionItemName(val int) string {
	var nilthis *QLocale
	return nilthis.FloatingPointPrecisionOptionItemName(val)
}

/*
Specifies the format of the currency symbol.



This enum was introduced or modified in  Qt 4.8.

*/
type QLocale__CurrencySymbolFormat = int

//
const QLocale__CurrencyIsoCode QLocale__CurrencySymbolFormat = 0

// a currency symbol.
const QLocale__CurrencySymbol QLocale__CurrencySymbolFormat = 1

// a user readable name of the currency.
const QLocale__CurrencyDisplayName QLocale__CurrencySymbolFormat = 2

func (this *QLocale) CurrencySymbolFormatItemName(val int) string {
	switch val {
	case QLocale__CurrencyIsoCode: // 0
		return "CurrencyIsoCode"
	case QLocale__CurrencySymbol: // 1
		return "CurrencySymbol"
	case QLocale__CurrencyDisplayName: // 2
		return "CurrencyDisplayName"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_CurrencySymbolFormatItemName(val int) string {
	var nilthis *QLocale
	return nilthis.CurrencySymbolFormatItemName(val)
}

/*


 */
type QLocale__DataSizeFormat = int

//
const QLocale__DataSizeBase1000 QLocale__DataSizeFormat = 1

//
const QLocale__DataSizeSIQuantifiers QLocale__DataSizeFormat = 2

//
const QLocale__DataSizeIecFormat QLocale__DataSizeFormat = 0

//
const QLocale__DataSizeTraditionalFormat QLocale__DataSizeFormat = 2

//
const QLocale__DataSizeSIFormat QLocale__DataSizeFormat = 3

func (this *QLocale) DataSizeFormatItemName(val int) string {
	switch val {
	case QLocale__DataSizeBase1000: // 1
		return "DataSizeBase1000"
	case QLocale__DataSizeSIQuantifiers: // 2
		return "DataSizeSIQuantifiers,DataSizeTraditionalFormat"
	case QLocale__DataSizeIecFormat: // 0
		return "DataSizeIecFormat"
		// case QLocale__DataSizeTraditionalFormat: // 2
		// return ""
	case QLocale__DataSizeSIFormat: // 3
		return "DataSizeSIFormat"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_DataSizeFormatItemName(val int) string {
	var nilthis *QLocale
	return nilthis.DataSizeFormatItemName(val)
}

/*
This enum defines a set of possible styles for locale specific quotation.



This enum was introduced or modified in  Qt 4.8.

See also quoteString().

*/
type QLocale__QuotationStyle = int

// If this option is set, the standard quotation marks will be used to quote strings.
const QLocale__StandardQuotation QLocale__QuotationStyle = 0

// If this option is set, the alternate quotation marks will be used to quote strings.
const QLocale__AlternateQuotation QLocale__QuotationStyle = 1

func (this *QLocale) QuotationStyleItemName(val int) string {
	switch val {
	case QLocale__StandardQuotation: // 0
		return "StandardQuotation"
	case QLocale__AlternateQuotation: // 1
		return "AlternateQuotation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLocale_QuotationStyleItemName(val int) string {
	var nilthis *QLocale
	return nilthis.QuotationStyleItemName(val)
}

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
