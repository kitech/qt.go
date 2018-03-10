package qtcore

// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QCollator struct {
	*qtrt.CObject
}
type QCollator_ITF interface {
	QCollator_PTR() *QCollator
}

func (ptr *QCollator) QCollator_PTR() *QCollator { return ptr }

func (this *QCollator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCollator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCollatorFromPointer(cthis unsafe.Pointer) *QCollator {
	return &QCollator{&qtrt.CObject{cthis}}
}
func (*QCollator) NewFromPointer(cthis unsafe.Pointer) *QCollator {
	return NewQCollatorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcollator.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCollator(const QLocale &)

/*
Constructs a QCollator from locale. If locale is not specified the system's default locale is used.

See also setLocale().
*/
func NewQCollator(locale QLocale_ITF) *QCollator {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatorC2ERK7QLocale", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCollatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCollator)
	return gothis
}

// /usr/include/qt/QtCore/qcollator.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCollator(const QLocale &)

/*
Constructs a QCollator from locale. If locale is not specified the system's default locale is used.

See also setLocale().
*/
func NewQCollator__() *QCollator {
	// arg: 0, const QLocale &=LValueReference, QLocale=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatorC2ERK7QLocale", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCollatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCollator)
	return gothis
}

// /usr/include/qt/QtCore/qcollator.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCollator()

/*

 */
func DeleteQCollator(this *QCollator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcollator.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QCollator & operator=(const QCollator &)

/*

 */
func (this *QCollator) Operator_equal(arg0 QCollator_ITF) *QCollator {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCollator_PTR() != nil {
		convArg0 = arg0.QCollator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCollatorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCollator)
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QCollator & operator=(QCollator &&)

/*

 */
func (this *QCollator) Operator_equal_1(other unsafe.Pointer /*333*/) *QCollator {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatoraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCollatorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCollator)
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCollator &)

/*
Swaps this collator with other. This function is very fast and never fails.
*/
func (this *QCollator) Swap(other QCollator_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCollator_PTR() != nil {
		convArg0 = other.QCollator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)

/*
Sets the locale of the collator to locale.

See also locale().
*/
func (this *QCollator) SetLocale(locale QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale() const

/*
Returns the locale of the collator.

See also setLocale().
*/
func (this *QCollator) Locale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity() const

/*
Returns case sensitivity of the collator.

See also setCaseSensitivity().
*/
func (this *QCollator) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcollator.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)

/*
Sets the case sensitivity of the collator.

See also caseSensitivity().
*/
func (this *QCollator) SetCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumericMode(_Bool)

/*
Enables numeric sorting mode when on is set to true.

This will enable proper sorting of numeric digits, so that e.g. 100 sorts after 99.

By default this mode is off.

Note: On Windows, this functionality makes use of the ICU library. If Qt was compiled without ICU support, it falls back to code using native Windows API, which only works from Windows 7 onwards. On older versions of Windows, it will not work and a warning will be emitted at runtime.

See also numericMode().
*/
func (this *QCollator) SetNumericMode(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator14setNumericModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool numericMode() const

/*
Returns true if numeric sorting is enabled, false otherwise.

See also setNumericMode().
*/
func (this *QCollator) NumericMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator11numericModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIgnorePunctuation(_Bool)

/*
If on is set to true, punctuation characters and symbols are ignored when determining sort order.

The default is locale dependent.

Note: This method is not currently supported if Qt is configured to not use ICU on Linux.

See also ignorePunctuation().
*/
func (this *QCollator) SetIgnorePunctuation(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator20setIgnorePunctuationEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ignorePunctuation() const

/*
Returns true if punctuation characters and symbols are ignored when determining sort order.

See also setIgnorePunctuation().
*/
func (this *QCollator) IgnorePunctuation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator17ignorePunctuationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, const QString &) const

/*
Compares s1 with s2. Returns an integer less than, equal to, or greater than zero depending on whether s1 is smaller, equal or larger than s2.
*/
func (this *QCollator) Compare(s1 string, s2 string) int {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7compareERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcollator.h:113
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &) const

/*
Compares s1 with s2. Returns an integer less than, equal to, or greater than zero depending on whether s1 is smaller, equal or larger than s2.
*/
func (this *QCollator) Compare_1(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7compareERK10QStringRefS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcollator.h:114
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(const QChar *, int, const QChar *, int) const

/*
Compares s1 with s2. Returns an integer less than, equal to, or greater than zero depending on whether s1 is smaller, equal or larger than s2.
*/
func (this *QCollator) Compare_2(s1 QChar_ITF /*777 const QChar **/, len1 int, s2 QChar_ITF /*777 const QChar **/, len2 int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QChar_PTR() != nil {
		convArg0 = s1.QChar_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if s2 != nil && s2.QChar_PTR() != nil {
		convArg2 = s2.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7compareEPK5QChariS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len1, convArg2, len2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcollator.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator()(const QString &, const QString &) const

/*

 */
func (this *QCollator) Operator_fncall(s1 string, s2 string) bool {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollatorclERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QCollatorSortKey sortKey(const QString &) const

/*
Returns a sortKey for string.

Creating the sort key is usually somewhat slower, than using the compare() methods directly. But if the string is compared repeatedly (e.g. when sorting a whole list of strings), it's usually faster to create the sort keys for each string and then sort using the keys.
*/
func (this *QCollator) SortKey(string string) *QCollatorSortKey /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7sortKeyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCollatorSortKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCollatorSortKey)
	return rv2
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
