package qtcore

// /usr/include/qt/QtCore/qversionnumber.h
// #include <qversionnumber.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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

type QVersionNumber struct {
	*qtrt.CObject
}

func (this *QVersionNumber) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVersionNumber) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVersionNumberFromPointer(cthis unsafe.Pointer) *QVersionNumber {
	return &QVersionNumber{&qtrt.CObject{cthis}}
}
func (*QVersionNumber) NewFromPointer(cthis unsafe.Pointer) *QVersionNumber {
	return NewQVersionNumberFromPointer(cthis)
}

// /usr/include/qt/QtCore/qversionnumber.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber()
func NewQVersionNumber() *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:242
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber(int)
func NewQVersionNumber_1(maj int) *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Ei", qtrt.FFI_TYPE_POINTER, maj)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:245
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber(int, int)
func NewQVersionNumber_2(maj int, min int) *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Eii", qtrt.FFI_TYPE_POINTER, maj, min)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:248
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber(int, int, int)
func NewQVersionNumber_3(maj int, min int, mic int) *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Eiii", qtrt.FFI_TYPE_POINTER, maj, min, mic)
	gopp.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:251
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QVersionNumber) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qversionnumber.h:254
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNormalized()
func (this *QVersionNumber) IsNormalized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12isNormalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qversionnumber.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int majorVersion()
func (this *QVersionNumber) MajorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12majorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:260
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minorVersion()
func (this *QVersionNumber) MinorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12minorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:263
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int microVersion()
func (this *QVersionNumber) MicroVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12microVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:266
// index:0
// Public Visibility=Default Availability=Available
// [8] QVersionNumber normalized()
func (this *QVersionNumber) Normalized() *QVersionNumber /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}

// /usr/include/qt/QtCore/qversionnumber.h:270
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int segmentAt(int)
func (this *QVersionNumber) SegmentAt(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber9segmentAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int segmentCount()
func (this *QVersionNumber) SegmentCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12segmentCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:276
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPrefixOf(const QVersionNumber &)
func (this *QVersionNumber) IsPrefixOf(other *QVersionNumber) bool {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber10isPrefixOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qversionnumber.h:278
// index:0
// Public static Visibility=Default Availability=Available
// [4] int compare(const QVersionNumber &, const QVersionNumber &)
func (this *QVersionNumber) Compare(v1 *QVersionNumber, v2 *QVersionNumber) int {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber7compareERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QVersionNumber_Compare(v1 *QVersionNumber, v2 *QVersionNumber) int {
	var nilthis *QVersionNumber
	rv := nilthis.Compare(v1, v2)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:280
// index:0
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber commonPrefix(const QVersionNumber &, const QVersionNumber &)
func (this *QVersionNumber) CommonPrefix(v1 *QVersionNumber, v2 *QVersionNumber) *QVersionNumber /*123*/ {
	var convArg0 = v1.GetCthis()
	var convArg1 = v2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber12commonPrefixERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_CommonPrefix(v1 *QVersionNumber, v2 *QVersionNumber) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.CommonPrefix(v1, v2)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:282
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QVersionNumber) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qversionnumber.h:284
// index:0
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(const QString &, int *)
func (this *QVersionNumber) FromString(string string, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringERK7QStringPi", qtrt.FFI_TYPE_POINTER, convArg0, &suffixIndex)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_FromString(string string, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.FromString(string, suffixIndex)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:286
// index:1
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(QLatin1String, int *)
func (this *QVersionNumber) FromString_1(string *QLatin1String /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE13QLatin1StringPi", qtrt.FFI_TYPE_POINTER, convArg0, &suffixIndex)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_FromString_1(string *QLatin1String /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.FromString_1(string, suffixIndex)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:287
// index:2
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(QStringView, int *)
func (this *QVersionNumber) FromString_2(string *QStringView /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE11QStringViewPi", qtrt.FFI_TYPE_POINTER, convArg0, &suffixIndex)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_FromString_2(string *QStringView /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.FromString_2(string, suffixIndex)
	return rv
}

func DeleteQVersionNumber(this *QVersionNumber) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QVersionNumber__ = int

const QVersionNumber__InlineSegmentMarker QVersionNumber__ = 0
const QVersionNumber__InlineSegmentStartIdx QVersionNumber__ = 1
const QVersionNumber__InlineSegmentCount QVersionNumber__ = 7

//  body block end
