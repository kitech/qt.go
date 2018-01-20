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
	*qtrt.CObject
}

func (this *QVersionNumber) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQVersionNumberFromPointer(cthis unsafe.Pointer) *QVersionNumber {
	return &QVersionNumber{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qversionnumber.h:221
// index:0
// Public inline
// void QVersionNumber()
func NewQVersionNumber() *QVersionNumber {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:242
// index:1
// Public inline
// void QVersionNumber(int)
func NewQVersionNumber_1(maj int) *QVersionNumber {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &maj)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:245
// index:2
// Public inline
// void QVersionNumber(int, int)
func NewQVersionNumber_2(maj int, min int) *QVersionNumber {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &maj, &min)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:248
// index:3
// Public inline
// void QVersionNumber(int, int, int)
func NewQVersionNumber_3(maj int, min int, mic int) *QVersionNumber {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumberC2Eiii", ffiqt.FFI_TYPE_VOID, cthis, &maj, &min, &mic)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:251
// index:0
// Public inline
// bool isNull()
func (this *QVersionNumber) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:254
// index:0
// Public inline
// bool isNormalized()
func (this *QVersionNumber) IsNormalized() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12isNormalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:257
// index:0
// Public inline
// int majorVersion()
func (this *QVersionNumber) MajorVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12majorVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:260
// index:0
// Public inline
// int minorVersion()
func (this *QVersionNumber) MinorVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12minorVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:263
// index:0
// Public inline
// int microVersion()
func (this *QVersionNumber) MicroVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12microVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:266
// index:0
// Public
// QVersionNumber normalized()
func (this *QVersionNumber) Normalized() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:268
// index:0
// Public
// QVector<int> segments()
func (this *QVersionNumber) Segments() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber8segmentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:270
// index:0
// Public inline
// int segmentAt(int)
func (this *QVersionNumber) SegmentAt(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber9segmentAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:273
// index:0
// Public inline
// int segmentCount()
func (this *QVersionNumber) SegmentCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber12segmentCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:276
// index:0
// Public
// bool isPrefixOf(const class QVersionNumber &)
func (this *QVersionNumber) IsPrefixOf(other *QVersionNumber) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber10isPrefixOfERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:278
// index:0
// Public static
// int compare(const class QVersionNumber &, const class QVersionNumber &)
func (this *QVersionNumber) Compare(v1 *QVersionNumber, v2 *QVersionNumber) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber7compareERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVersionNumber_Compare(v1 *QVersionNumber, v2 *QVersionNumber) {
	var nilthis *QVersionNumber
	nilthis.Compare(v1, v2)
}

// /usr/include/qt/QtCore/qversionnumber.h:280
// index:0
// Public static
// QVersionNumber commonPrefix(const class QVersionNumber &, const class QVersionNumber &)
func (this *QVersionNumber) CommonPrefix(v1 *QVersionNumber, v2 *QVersionNumber) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber12commonPrefixERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVersionNumber_CommonPrefix(v1 *QVersionNumber, v2 *QVersionNumber) {
	var nilthis *QVersionNumber
	nilthis.CommonPrefix(v1, v2)
}

// /usr/include/qt/QtCore/qversionnumber.h:282
// index:0
// Public
// QString toString()
func (this *QVersionNumber) ToString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QVersionNumber8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:284
// index:0
// Public static
// QVersionNumber fromString(const class QString &, int *)
func (this *QVersionNumber) FromString(string *QString, suffixIndex unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringERK7QStringPi", ffiqt.FFI_TYPE_POINTER, string, suffixIndex)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVersionNumber_FromString(string *QString, suffixIndex unsafe.Pointer) {
	var nilthis *QVersionNumber
	nilthis.FromString(string, suffixIndex)
}

// /usr/include/qt/QtCore/qversionnumber.h:286
// index:1
// Public static
// QVersionNumber fromString(class QLatin1String, int *)
func (this *QVersionNumber) FromString_1(string *QLatin1String, suffixIndex unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE13QLatin1StringPi", ffiqt.FFI_TYPE_POINTER, string, suffixIndex)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVersionNumber_FromString_1(string *QLatin1String, suffixIndex unsafe.Pointer) {
	var nilthis *QVersionNumber
	nilthis.FromString_1(string, suffixIndex)
}

// /usr/include/qt/QtCore/qversionnumber.h:287
// index:2
// Public static
// QVersionNumber fromString(class QStringView, int *)
func (this *QVersionNumber) FromString_2(string *QStringView, suffixIndex unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE11QStringViewPi", ffiqt.FFI_TYPE_POINTER, string, suffixIndex)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVersionNumber_FromString_2(string *QStringView, suffixIndex unsafe.Pointer) {
	var nilthis *QVersionNumber
	nilthis.FromString_2(string, suffixIndex)
}

//  body block end
