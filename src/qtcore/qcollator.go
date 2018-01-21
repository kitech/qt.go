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
func NewQCollatorFromPointer(cthis unsafe.Pointer) *QCollator {
	return &QCollator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qcollator.h:86
// index:0
// Public
// void QCollator(const class QLocale &)
func NewQCollator(locale *QLocale) *QCollator {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollatorC2ERK7QLocale", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCollatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcollator.h:88
// index:0
// Public
// void ~QCollator()
func DeleteQCollator(*QCollator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:97
// index:0
// Public inline
// void swap(class QCollator &)
func (this *QCollator) Swap(other *QCollator) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:100
// index:0
// Public
// void setLocale(const class QLocale &)
func (this *QCollator) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator9setLocaleERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:101
// index:0
// Public
// QLocale locale()
func (this *QCollator) Locale() *QLocale /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator6localeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:103
// index:0
// Public
// Qt::CaseSensitivity caseSensitivity()
func (this *QCollator) CaseSensitivity() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator15caseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qcollator.h:104
// index:0
// Public
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QCollator) SetCaseSensitivity(cs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:106
// index:0
// Public
// void setNumericMode(_Bool)
func (this *QCollator) SetNumericMode(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator14setNumericModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:107
// index:0
// Public
// bool numericMode()
func (this *QCollator) NumericMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator11numericModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:109
// index:0
// Public
// void setIgnorePunctuation(_Bool)
func (this *QCollator) SetIgnorePunctuation(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator20setIgnorePunctuationEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:110
// index:0
// Public
// bool ignorePunctuation()
func (this *QCollator) IgnorePunctuation() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator17ignorePunctuationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcollator.h:112
// index:0
// Public
// int compare(const class QString &, const class QString &)
func (this *QCollator) Compare(s1 *QString, s2 *QString) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7compareERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qcollator.h:113
// index:1
// Public
// int compare(const class QStringRef &, const class QStringRef &)
func (this *QCollator) Compare_1(s1 *QStringRef, s2 *QStringRef) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7compareERK10QStringRefS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qcollator.h:114
// index:2
// Public
// int compare(const class QChar *, int, const class QChar *, int)
func (this *QCollator) Compare_2(s1 *QChar /*444 const QChar **/, len1 int, s2 *QChar /*444 const QChar **/, len2 int) int {
	var convArg0 = s1.GetCthis()
	var convArg2 = s2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7compareEPK5QChariS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len1, convArg2, &len2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qcollator.h:119
// index:0
// Public
// QCollatorSortKey sortKey(const class QString &)
func (this *QCollator) SortKey(string *QString) *QCollatorSortKey /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7sortKeyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCollatorSortKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
