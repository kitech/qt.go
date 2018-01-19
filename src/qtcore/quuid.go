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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/quuid.h:89
// index:0
// inline
// void QUuid()
func NewQUuid() *QUuid {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QUuid{cthis}
}

// /usr/include/qt/QtCore/quuid.h:91
// index:1
// inline
// void QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)
func NewQUuid_1(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2Ejtthhhhhhhh", ffiqt.FFI_TYPE_VOID, cthis, &l, &w1, &w2, &b1, &b2, &b3, &b4, &b5, &b6, &b7, &b8)
	gopp.ErrPrint(err, rv)
	return &QUuid{cthis}
}

// /usr/include/qt/QtCore/quuid.h:119
// index:2
// void QUuid(const class QString &)
func NewQUuid_2(arg0 unsafe.Pointer) *QUuid {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	return &QUuid{cthis}
}

// /usr/include/qt/QtCore/quuid.h:122
// index:3
// void QUuid(const char *)
func NewQUuid_3(arg0 unsafe.Pointer) *QUuid {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	return &QUuid{cthis}
}

// /usr/include/qt/QtCore/quuid.h:124
// index:4
// void QUuid(const class QByteArray &)
func NewQUuid_4(arg0 unsafe.Pointer) *QUuid {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuidC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	return &QUuid{cthis}
}

// /usr/include/qt/QtCore/quuid.h:120
// index:0
// static
// QUuid fromString(class QStringView)
func (this *QUuid) FromString(string unsafe.Pointer) {
	// 0: (QStringView string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid10fromStringE11QStringView", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_FromString(string unsafe.Pointer) {
	// 0: (QStringView string), (string)
	var nilthis *QUuid
	nilthis.FromString(string)
}

// /usr/include/qt/QtCore/quuid.h:121
// index:1
// static
// QUuid fromString(class QLatin1String)
func (this *QUuid) FromString_1(string unsafe.Pointer) {
	// 1: (QLatin1String string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid10fromStringE13QLatin1String", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_FromString_1(string unsafe.Pointer) {
	// 1: (QLatin1String string), (string)
	var nilthis *QUuid
	nilthis.FromString_1(string)
}

// /usr/include/qt/QtCore/quuid.h:123
// index:0
// QString toString()
func (this *QUuid) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid8toStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/quuid.h:125
// index:0
// QByteArray toByteArray()
func (this *QUuid) ToByteArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid11toByteArrayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/quuid.h:126
// index:0
// QByteArray toRfc4122()
func (this *QUuid) ToRfc4122() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid9toRfc4122Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/quuid.h:127
// index:0
// static
// QUuid fromRfc4122(const class QByteArray &)
func (this *QUuid) FromRfc4122(arg0 unsafe.Pointer) {
	// 0: (const QByteArray & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid11fromRfc4122ERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_FromRfc4122(arg0 unsafe.Pointer) {
	// 0: (const QByteArray & arg0), (arg0)
	var nilthis *QUuid
	nilthis.FromRfc4122(arg0)
}

// /usr/include/qt/QtCore/quuid.h:128
// index:0
// bool isNull()
func (this *QUuid) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/quuid.h:192
// index:0
// static
// QUuid createUuid()
func (this *QUuid) CreateUuid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid10createUuidEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_CreateUuid() {
	// 0: (), ()
	var nilthis *QUuid
	nilthis.CreateUuid()
}

// /usr/include/qt/QtCore/quuid.h:194
// index:0
// static
// QUuid createUuidV3(const class QUuid &, const class QByteArray &)
func (this *QUuid) CreateUuidV3(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 0: (const QUuid & ns, const QByteArray & baseData), (ns, baseData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_CreateUuidV3(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 0: (const QUuid & ns, const QByteArray & baseData), (ns, baseData)
	var nilthis *QUuid
	nilthis.CreateUuidV3(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:196
// index:1
// static inline
// QUuid createUuidV3(const class QUuid &, const class QString &)
func (this *QUuid) CreateUuidV3_1(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 1: (const QUuid & ns, const QString & baseData), (ns, baseData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_CreateUuidV3_1(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 1: (const QUuid & ns, const QString & baseData), (ns, baseData)
	var nilthis *QUuid
	nilthis.CreateUuidV3_1(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:195
// index:0
// static
// QUuid createUuidV5(const class QUuid &, const class QByteArray &)
func (this *QUuid) CreateUuidV5(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 0: (const QUuid & ns, const QByteArray & baseData), (ns, baseData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_CreateUuidV5(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 0: (const QUuid & ns, const QByteArray & baseData), (ns, baseData)
	var nilthis *QUuid
	nilthis.CreateUuidV5(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:201
// index:1
// static inline
// QUuid createUuidV5(const class QUuid &, const class QString &)
func (this *QUuid) CreateUuidV5_1(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 1: (const QUuid & ns, const QString & baseData), (ns, baseData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUuid_CreateUuidV5_1(ns unsafe.Pointer, baseData unsafe.Pointer) {
	// 1: (const QUuid & ns, const QString & baseData), (ns, baseData)
	var nilthis *QUuid
	nilthis.CreateUuidV5_1(ns, baseData)
}

// /usr/include/qt/QtCore/quuid.h:208
// index:0
// QUuid::Variant variant()
func (this *QUuid) Variant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid7variantEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/quuid.h:209
// index:0
// QUuid::Version version()
func (this *QUuid) Version() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QUuid7versionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
