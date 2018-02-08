package qtcore

// /usr/include/qt/QtCore/quuid.h
// #include <quuid.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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

type QUuid struct {
	*qtrt.CObject
}

func (this *QUuid) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QUuid) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQUuidFromPointer(cthis unsafe.Pointer) *QUuid {
	return &QUuid{&qtrt.CObject{cthis}}
}
func (*QUuid) NewFromPointer(cthis unsafe.Pointer) *QUuid {
	return NewQUuidFromPointer(cthis)
}

// /usr/include/qt/QtCore/quuid.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QUuid()
func NewQUuid() *QUuid {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)
func NewQUuid_1(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2Ejtthhhhhhhh", qtrt.FFI_TYPE_POINTER, l, w1, w2, b1, b2, b3, b4, b5, b6, b7, b8)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:119
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUuid(const QString &)
func NewQUuid_2(arg0 string) *QUuid {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:122
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QUuid(const char *)
func NewQUuid_3(arg0 string) *QUuid {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:124
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QUuid(const QByteArray &)
func NewQUuid_4(arg0 *QByteArray) *QUuid {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid fromString(QStringView)
func (this *QUuid) FromString(string *QStringView /*123*/) *QUuid /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid10fromStringE11QStringView", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_FromString(string *QStringView /*123*/) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.FromString(string)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:121
// index:1
// Public static Visibility=Default Availability=Available
// [16] QUuid fromString(QLatin1String)
func (this *QUuid) FromString_1(string *QLatin1String /*123*/) *QUuid /*123*/ {
	var convArg0 = string.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid10fromStringE13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_FromString_1(string *QLatin1String /*123*/) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.FromString_1(string)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QUuid) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/quuid.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toByteArray()
func (this *QUuid) ToByteArray() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid11toByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/quuid.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toRfc4122()
func (this *QUuid) ToRfc4122() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid9toRfc4122Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/quuid.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid fromRfc4122(const QByteArray &)
func (this *QUuid) FromRfc4122(arg0 *QByteArray) *QUuid /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid11fromRfc4122ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_FromRfc4122(arg0 *QByteArray) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.FromRfc4122(arg0)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QUuid) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/quuid.h:192
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid createUuid()
func (this *QUuid) CreateUuid() *QUuid /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid10createUuidEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuid() *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuid()
	return rv
}

// /usr/include/qt/QtCore/quuid.h:194
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid createUuidV3(const QUuid &, const QByteArray &)
func (this *QUuid) CreateUuidV3(ns *QUuid, baseData *QByteArray) *QUuid /*123*/ {
	var convArg0 = ns.GetCthis()
	var convArg1 = baseData.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV3(ns *QUuid, baseData *QByteArray) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV3(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:196
// index:1
// Public static inline Visibility=Default Availability=Available
// [16] QUuid createUuidV3(const QUuid &, const QString &)
func (this *QUuid) CreateUuidV3_1(ns *QUuid, baseData string) *QUuid /*123*/ {
	var convArg0 = ns.GetCthis()
	var tmpArg1 = NewQString_5(baseData)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV3_1(ns *QUuid, baseData string) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV3_1(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:195
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid createUuidV5(const QUuid &, const QByteArray &)
func (this *QUuid) CreateUuidV5(ns *QUuid, baseData *QByteArray) *QUuid /*123*/ {
	var convArg0 = ns.GetCthis()
	var convArg1 = baseData.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV5(ns *QUuid, baseData *QByteArray) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV5(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:201
// index:1
// Public static inline Visibility=Default Availability=Available
// [16] QUuid createUuidV5(const QUuid &, const QString &)
func (this *QUuid) CreateUuidV5_1(ns *QUuid, baseData string) *QUuid /*123*/ {
	var convArg0 = ns.GetCthis()
	var tmpArg1 = NewQString_5(baseData)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV5_1(ns *QUuid, baseData string) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV5_1(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:208
// index:0
// Public Visibility=Default Availability=Available
// [4] QUuid::Variant variant()
func (this *QUuid) Variant() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid7variantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/quuid.h:209
// index:0
// Public Visibility=Default Availability=Available
// [4] QUuid::Version version()
func (this *QUuid) Version() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQUuid(this *QUuid) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QUuid__Variant = int

const QUuid__VarUnknown QUuid__Variant = -1
const QUuid__NCS QUuid__Variant = 0
const QUuid__DCE QUuid__Variant = 2
const QUuid__Microsoft QUuid__Variant = 6
const QUuid__Reserved QUuid__Variant = 7

type QUuid__Version = int

const QUuid__VerUnknown QUuid__Version = -1
const QUuid__Time QUuid__Version = 1
const QUuid__EmbeddedPOSIX QUuid__Version = 2
const QUuid__Md5 QUuid__Version = 3
const QUuid__Name QUuid__Version = 3
const QUuid__Random QUuid__Version = 4
const QUuid__Sha1 QUuid__Version = 5

//  body block end
