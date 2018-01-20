//  header block begin
// /usr/include/qt/QtCore/quuid.h
// #include <quuid.h>
// #include <QtCore>
package qtcore

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
type QUuid struct {
	*qtrt.CObject
}

func (this *QUuid) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQUuidFromPointer(cthis unsafe.Pointer) *QUuid {
	return &QUuid{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/quuid.h:89
// index:0
// Public inline
// void QUuid()
func NewQUuid() *QUuid {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:91
// index:1
// Public inline
// void QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)
func NewQUuid_1(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2Ejtthhhhhhhh", ffiqt.FFI_TYPE_VOID, cthis, &l, &w1, &w2, &b1, &b2, &b3, &b4, &b5, &b6, &b7, &b8)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:119
// index:2
// Public
// void QUuid(const class QString &)
func NewQUuid_2(arg0 *QString) *QUuid {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:122
// index:3
// Public
// void QUuid(const char *)
func NewQUuid_3(arg0 string) *QUuid {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:124
// index:4
// Public
// void QUuid(const class QByteArray &)
func NewQUuid_4(arg0 *QByteArray) *QUuid {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:120
// index:0
// Public static
// QUuid fromString(class QStringView)
func (this *QUuid) FromString(string *QStringView) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid10fromStringE11QStringView", ffiqt.FFI_TYPE_POINTER, string)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_FromString(string *QStringView) {
	var nilthis *QUuid
	nilthis.FromString(string)
}

// /usr/include/qt/QtCore/quuid.h:121
// index:1
// Public static
// QUuid fromString(class QLatin1String)
func (this *QUuid) FromString_1(string *QLatin1String) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid10fromStringE13QLatin1String", ffiqt.FFI_TYPE_POINTER, string)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_FromString_1(string *QLatin1String) {
	var nilthis *QUuid
	nilthis.FromString_1(string)
}

// /usr/include/qt/QtCore/quuid.h:123
// index:0
// Public
// QString toString()
func (this *QUuid) ToString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:125
// index:0
// Public
// QByteArray toByteArray()
func (this *QUuid) ToByteArray() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid11toByteArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:126
// index:0
// Public
// QByteArray toRfc4122()
func (this *QUuid) ToRfc4122() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid9toRfc4122Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:127
// index:0
// Public static
// QUuid fromRfc4122(const class QByteArray &)
func (this *QUuid) FromRfc4122(arg0 *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid11fromRfc4122ERK10QByteArray", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_FromRfc4122(arg0 *QByteArray) {
	var nilthis *QUuid
	nilthis.FromRfc4122(arg0)
}

// /usr/include/qt/QtCore/quuid.h:128
// index:0
// Public
// bool isNull()
func (this *QUuid) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:192
// index:0
// Public static
// QUuid createUuid()
func (this *QUuid) CreateUuid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid10createUuidEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_CreateUuid() {
	var nilthis *QUuid
	nilthis.CreateUuid()
}

// /usr/include/qt/QtCore/quuid.h:194
// index:0
// Public static
// QUuid createUuidV3(const class QUuid &, const class QByteArray &)
func (this *QUuid) CreateUuidV3(ns *QUuid, baseData *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK10QByteArray", ffiqt.FFI_TYPE_POINTER, ns, baseData)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_CreateUuidV3(ns *QUuid, baseData *QByteArray) {
	var nilthis *QUuid
	nilthis.CreateUuidV3(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:196
// index:1
// Public static inline
// QUuid createUuidV3(const class QUuid &, const class QString &)
func (this *QUuid) CreateUuidV3_1(ns *QUuid, baseData *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK7QString", ffiqt.FFI_TYPE_POINTER, ns, baseData)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_CreateUuidV3_1(ns *QUuid, baseData *QString) {
	var nilthis *QUuid
	nilthis.CreateUuidV3_1(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:195
// index:0
// Public static
// QUuid createUuidV5(const class QUuid &, const class QByteArray &)
func (this *QUuid) CreateUuidV5(ns *QUuid, baseData *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK10QByteArray", ffiqt.FFI_TYPE_POINTER, ns, baseData)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_CreateUuidV5(ns *QUuid, baseData *QByteArray) {
	var nilthis *QUuid
	nilthis.CreateUuidV5(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:201
// index:1
// Public static inline
// QUuid createUuidV5(const class QUuid &, const class QString &)
func (this *QUuid) CreateUuidV5_1(ns *QUuid, baseData *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK7QString", ffiqt.FFI_TYPE_POINTER, ns, baseData)
	gopp.ErrPrint(err, rv)
	return rv
}
func QUuid_CreateUuidV5_1(ns *QUuid, baseData *QString) {
	var nilthis *QUuid
	nilthis.CreateUuidV5_1(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:208
// index:0
// Public
// QUuid::Variant variant()
func (this *QUuid) Variant() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid7variantEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:209
// index:0
// Public
// QUuid::Version version()
func (this *QUuid) Version() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid7versionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
