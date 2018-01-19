//  header block begin
// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>
package qtcore

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qcollator.h:86
// index:0
// void QCollator(const class QLocale &)
func NewQCollator(locale unsafe.Pointer) *QCollator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollatorC2ERK7QLocale", ffiqt.FFI_TYPE_VOID, cthis, locale)
	gopp.ErrPrint(err, rv)
	return &QCollator{cthis}
}

// /usr/include/qt/QtCore/qcollator.h:88
// index:0
// void ~QCollator()
func DeleteQCollator(*QCollator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:97
// index:0
// inline
// void swap(class QCollator &)
func (this *QCollator) Swap(other unsafe.Pointer) {
	// 0: (, QCollator & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:100
// index:0
// void setLocale(const class QLocale &)
func (this *QCollator) SetLocale(locale unsafe.Pointer) {
	// 0: (, const QLocale & locale), (locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator9setLocaleERK7QLocale", ffiqt.FFI_TYPE_VOID, this.cthis, locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:101
// index:0
// QLocale locale()
func (this *QCollator) Locale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator6localeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:103
// index:0
// Qt::CaseSensitivity caseSensitivity()
func (this *QCollator) CaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator15caseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:104
// index:0
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QCollator) SetCaseSensitivity(cs int) {
	// 0: (, Qt::CaseSensitivity cs), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:106
// index:0
// void setNumericMode(_Bool)
func (this *QCollator) SetNumericMode(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator14setNumericModeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:107
// index:0
// bool numericMode()
func (this *QCollator) NumericMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator11numericModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:109
// index:0
// void setIgnorePunctuation(_Bool)
func (this *QCollator) SetIgnorePunctuation(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QCollator20setIgnorePunctuationEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:110
// index:0
// bool ignorePunctuation()
func (this *QCollator) IgnorePunctuation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator17ignorePunctuationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:112
// index:0
// int compare(const class QString &, const class QString &)
func (this *QCollator) Compare(s1 unsafe.Pointer, s2 unsafe.Pointer) {
	// 0: (, const QString & s1, const QString & s2), (s1, s2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7compareERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, s1, s2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:113
// index:1
// int compare(const class QStringRef &, const class QStringRef &)
func (this *QCollator) Compare_1(s1 unsafe.Pointer, s2 unsafe.Pointer) {
	// 1: (, const QStringRef & s1, const QStringRef & s2), (s1, s2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7compareERK10QStringRefS2_", ffiqt.FFI_TYPE_VOID, this.cthis, s1, s2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:114
// index:2
// int compare(const class QChar *, int, const class QChar *, int)
func (this *QCollator) Compare_2(s1 unsafe.Pointer, len1 int, s2 unsafe.Pointer, len2 int) {
	// 2: (, const QChar * s1, int len1, const QChar * s2, int len2), (s1, &len1, s2, &len2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7compareEPK5QChariS2_i", ffiqt.FFI_TYPE_VOID, this.cthis, s1, &len1, s2, &len2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:119
// index:0
// QCollatorSortKey sortKey(const class QString &)
func (this *QCollator) SortKey(string unsafe.Pointer) {
	// 0: (, const QString & string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QCollator7sortKeyERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, string)
	gopp.ErrPrint(err, rv)
}

//  body block end
