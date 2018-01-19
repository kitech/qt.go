//  header block begin
// /usr/include/qt/QtCore/qversionnumber.h
// #include <qversionnumber.h>
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
type QVersionNumber struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qversionnumber.h:221
// index:0
// inline
// void QVersionNumber()
func NewQVersionNumber() *QVersionNumber {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QVersionNumber{cthis}
}

// /usr/include/qt/QtCore/qversionnumber.h:242
// index:1
// inline
// void QVersionNumber(int)
func NewQVersionNumber_1(maj int) *QVersionNumber {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &maj)
	gopp.ErrPrint(err, rv)
	return &QVersionNumber{cthis}
}

// /usr/include/qt/QtCore/qversionnumber.h:245
// index:2
// inline
// void QVersionNumber(int, int)
func NewQVersionNumber_2(maj int, min int) *QVersionNumber {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &maj, &min)
	gopp.ErrPrint(err, rv)
	return &QVersionNumber{cthis}
}

// /usr/include/qt/QtCore/qversionnumber.h:248
// index:3
// inline
// void QVersionNumber(int, int, int)
func NewQVersionNumber_3(maj int, min int, mic int) *QVersionNumber {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Eiii", ffiqt.FFI_TYPE_VOID, cthis, &maj, &min, &mic)
	gopp.ErrPrint(err, rv)
	return &QVersionNumber{cthis}
}

// /usr/include/qt/QtCore/qversionnumber.h:251
// index:0
// inline
// bool isNull()
func (this *QVersionNumber) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:254
// index:0
// inline
// bool isNormalized()
func (this *QVersionNumber) IsNormalized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12isNormalizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:257
// index:0
// inline
// int majorVersion()
func (this *QVersionNumber) MajorVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12majorVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:260
// index:0
// inline
// int minorVersion()
func (this *QVersionNumber) MinorVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12minorVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:263
// index:0
// inline
// int microVersion()
func (this *QVersionNumber) MicroVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12microVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:266
// index:0
// QVersionNumber normalized()
func (this *QVersionNumber) Normalized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber10normalizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:268
// index:0
// QVector<int> segments()
func (this *QVersionNumber) Segments() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber8segmentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:270
// index:0
// inline
// int segmentAt(int)
func (this *QVersionNumber) SegmentAt(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber9segmentAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:273
// index:0
// inline
// int segmentCount()
func (this *QVersionNumber) SegmentCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12segmentCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:276
// index:0
// bool isPrefixOf(const class QVersionNumber &)
func (this *QVersionNumber) IsPrefixOf(other unsafe.Pointer) {
	// 0: (, const QVersionNumber & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber10isPrefixOfERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:278
// index:0
// static
// int compare(const class QVersionNumber &, const class QVersionNumber &)
func (this *QVersionNumber) Compare(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (const QVersionNumber & v1, const QVersionNumber & v2), (v1, v2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber7compareERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVersionNumber_Compare(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (const QVersionNumber & v1, const QVersionNumber & v2), (v1, v2)
	var nilthis *QVersionNumber
	nilthis.Compare(v1, v2)
}

// /usr/include/qt/QtCore/qversionnumber.h:280
// index:0
// static
// QVersionNumber commonPrefix(const class QVersionNumber &, const class QVersionNumber &)
func (this *QVersionNumber) CommonPrefix(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (const QVersionNumber & v1, const QVersionNumber & v2), (v1, v2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber12commonPrefixERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVersionNumber_CommonPrefix(v1 unsafe.Pointer, v2 unsafe.Pointer) {
	// 0: (const QVersionNumber & v1, const QVersionNumber & v2), (v1, v2)
	var nilthis *QVersionNumber
	nilthis.CommonPrefix(v1, v2)
}

// /usr/include/qt/QtCore/qversionnumber.h:282
// index:0
// QString toString()
func (this *QVersionNumber) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber8toStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qversionnumber.h:284
// index:0
// static
// QVersionNumber fromString(const class QString &, int *)
func (this *QVersionNumber) FromString(string unsafe.Pointer, suffixIndex unsafe.Pointer) {
	// 0: (const QString & string, int * suffixIndex), (string, suffixIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringERK7QStringPi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVersionNumber_FromString(string unsafe.Pointer, suffixIndex unsafe.Pointer) {
	// 0: (const QString & string, int * suffixIndex), (string, suffixIndex)
	var nilthis *QVersionNumber
	nilthis.FromString(string, suffixIndex)
}

// /usr/include/qt/QtCore/qversionnumber.h:286
// index:1
// static
// QVersionNumber fromString(class QLatin1String, int *)
func (this *QVersionNumber) FromString_1(string unsafe.Pointer, suffixIndex unsafe.Pointer) {
	// 1: (QLatin1String string, int * suffixIndex), (string, suffixIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE13QLatin1StringPi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVersionNumber_FromString_1(string unsafe.Pointer, suffixIndex unsafe.Pointer) {
	// 1: (QLatin1String string, int * suffixIndex), (string, suffixIndex)
	var nilthis *QVersionNumber
	nilthis.FromString_1(string, suffixIndex)
}

// /usr/include/qt/QtCore/qversionnumber.h:287
// index:2
// static
// QVersionNumber fromString(class QStringView, int *)
func (this *QVersionNumber) FromString_2(string unsafe.Pointer, suffixIndex unsafe.Pointer) {
	// 2: (QStringView string, int * suffixIndex), (string, suffixIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE11QStringViewPi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVersionNumber_FromString_2(string unsafe.Pointer, suffixIndex unsafe.Pointer) {
	// 2: (QStringView string, int * suffixIndex), (string, suffixIndex)
	var nilthis *QVersionNumber
	nilthis.FromString_2(string, suffixIndex)
}

//  body block end
