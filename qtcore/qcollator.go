package qtcore

// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QCollator struct {
	*qtrt.CObject
}

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
func NewQCollator(locale *QLocale) *QCollator {
	var convArg0 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatorC2ERK7QLocale", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCollatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCollator)
	return gothis
}

// /usr/include/qt/QtCore/qcollator.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCollator()
func DeleteQCollator(this *QCollator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcollator.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCollator &)
func (this *QCollator) Swap(other *QCollator) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)
func (this *QCollator) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QCollator) Locale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity()
func (this *QCollator) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qcollator.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QCollator) SetCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumericMode(_Bool)
func (this *QCollator) SetNumericMode(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator14setNumericModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool numericMode()
func (this *QCollator) NumericMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator11numericModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIgnorePunctuation(_Bool)
func (this *QCollator) SetIgnorePunctuation(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QCollator20setIgnorePunctuationEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ignorePunctuation()
func (this *QCollator) IgnorePunctuation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator17ignorePunctuationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, const QString &)
func (this *QCollator) Compare(s1 *QString, s2 *QString) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7compareERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcollator.h:113
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &)
func (this *QCollator) Compare_1(s1 *QStringRef, s2 *QStringRef) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7compareERK10QStringRefS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcollator.h:114
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(const QChar *, int, const QChar *, int)
func (this *QCollator) Compare_2(s1 *QChar /*777 const QChar **/, len1 int, s2 *QChar /*777 const QChar **/, len2 int) int {
	var convArg0 = s1.GetCthis()
	var convArg2 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7compareEPK5QChariS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len1, convArg2, len2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcollator.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QCollatorSortKey sortKey(const QString &)
func (this *QCollator) SortKey(string *QString) *QCollatorSortKey /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QCollator7sortKeyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCollatorSortKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCollatorSortKey)
	return rv2
}

//  body block end
