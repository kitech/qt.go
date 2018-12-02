package qtcore

// /usr/include/qt/QtCore/quuid.h
// #include <quuid.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QUuid struct {
	*qtrt.CObject
}
type QUuid_ITF interface {
	QUuid_PTR() *QUuid
}

func (ptr *QUuid) QUuid_PTR() *QUuid { return ptr }

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

/*
Creates the null UUID. toString() will output the null UUID as "{00000000-0000-0000-0000-000000000000}".
*/
func (*QUuid) NewForInherit() *QUuid {
	return NewQUuid()
}
func NewQUuid() *QUuid {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)

/*
Creates the null UUID. toString() will output the null UUID as "{00000000-0000-0000-0000-000000000000}".
*/
func (*QUuid) NewForInherit_1(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
	return NewQUuid_1(l, w1, w2, b1, b2, b3, b4, b5, b6, b7, b8)
}
func NewQUuid_1(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2Ejtthhhhhhhh", qtrt.FFI_TYPE_POINTER, l, w1, w2, b1, b2, b3, b4, b5, b6, b7, b8)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:119
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QUuid(const QString &)

/*
Creates the null UUID. toString() will output the null UUID as "{00000000-0000-0000-0000-000000000000}".
*/
func (*QUuid) NewForInherit_2(arg0 string) *QUuid {
	return NewQUuid_2(arg0)
}
func NewQUuid_2(arg0 string) *QUuid {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:122
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QUuid(const char *)

/*
Creates the null UUID. toString() will output the null UUID as "{00000000-0000-0000-0000-000000000000}".
*/
func (*QUuid) NewForInherit_3(arg0 string) *QUuid {
	return NewQUuid_3(arg0)
}
func NewQUuid_3(arg0 string) *QUuid {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:124
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QUuid(const QByteArray &)

/*
Creates the null UUID. toString() will output the null UUID as "{00000000-0000-0000-0000-000000000000}".
*/
func (*QUuid) NewForInherit_4(arg0 QByteArray_ITF) *QUuid {
	return NewQUuid_4(arg0)
}
func NewQUuid_4(arg0 QByteArray_ITF) *QUuid {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUuidFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUuid)
	return gothis
}

// /usr/include/qt/QtCore/quuid.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid fromString(QStringView)

/*
Creates a QUuid object from the string text, which must be formatted as five hex fields separated by '-', e.g., "{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}" where each 'x' is a hex digit. The curly braces shown here are optional, but it is normal to include them. If the conversion fails, a null UUID is returned. See toString() for an explanation of how the five hex fields map to the public data members in QUuid.

This function was introduced in  Qt 5.10.

See also toString() and QUuid().
*/
func (this *QUuid) FromString(string QStringView_ITF /*123*/) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if string != nil && string.QStringView_PTR() != nil {
		convArg0 = string.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid10fromStringE11QStringView", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_FromString(string QStringView_ITF /*123*/) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.FromString(string)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:121
// index:1
// Public static Visibility=Default Availability=Available
// [16] QUuid fromString(QLatin1String)

/*
Creates a QUuid object from the string text, which must be formatted as five hex fields separated by '-', e.g., "{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}" where each 'x' is a hex digit. The curly braces shown here are optional, but it is normal to include them. If the conversion fails, a null UUID is returned. See toString() for an explanation of how the five hex fields map to the public data members in QUuid.

This function was introduced in  Qt 5.10.

See also toString() and QUuid().
*/
func (this *QUuid) FromString_1(string QLatin1String_ITF /*123*/) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if string != nil && string.QLatin1String_PTR() != nil {
		convArg0 = string.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid10fromStringE13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_FromString_1(string QLatin1String_ITF /*123*/) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.FromString_1(string)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the string representation of this QUuid. The string is formatted as five hex fields separated by '-' and enclosed in curly braces, i.e., "{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}" where 'x' is a hex digit. From left to right, the five hex fields are obtained from the four public data members in QUuid as follows:


 Field #Source
1data1
2data2
3data3
4data4[0] .. data4[1]
5data4[2] .. data4[7]
*/
func (this *QUuid) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/quuid.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toByteArray() const

/*
Returns the binary representation of this QUuid. The byte array is formatted as five hex fields separated by '-' and enclosed in curly braces, i.e., "{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}" where 'x' is a hex digit. From left to right, the five hex fields are obtained from the four public data members in QUuid as follows:


 Field #Source
1data1
2data2
3data3
4data4[0] .. data4[1]
5data4[2] .. data4[7]


This function was introduced in  Qt 4.8.
*/
func (this *QUuid) ToByteArray() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid11toByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/quuid.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toRfc4122() const

/*
Returns the binary representation of this QUuid. The byte array is in big endian format, and formatted according to RFC 4122, section 4.1.2 - "Layout and byte order".

The order is as follows:


 Field #Source
1data1
2data2
3data3
4data4[0] .. data4[7]


This function was introduced in  Qt 4.8.
*/
func (this *QUuid) ToRfc4122() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid9toRfc4122Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/quuid.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid fromRfc4122(const QByteArray &)

/*
Creates a QUuid object from the binary representation of the UUID, as specified by RFC 4122 section 4.1.2. See toRfc4122() for a further explanation of the order of bytes required.

The byte array accepted is NOT a human readable format.

If the conversion fails, a null UUID is created.

This function was introduced in  Qt 4.8.

See also toRfc4122() and QUuid().
*/
func (this *QUuid) FromRfc4122(arg0 QByteArray_ITF) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid11fromRfc4122ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_FromRfc4122(arg0 QByteArray_ITF) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.FromRfc4122(arg0)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is the null UUID {00000000-0000-0000-0000-000000000000}; otherwise returns false.
*/
func (this *QUuid) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/quuid.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QUuid &) const

/*

 */
func (this *QUuid) Operator_equal_equal(orig QUuid_ITF) bool {
	var convArg0 unsafe.Pointer
	if orig != nil && orig.QUuid_PTR() != nil {
		convArg0 = orig.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuideqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/quuid.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QUuid &) const

/*

 */
func (this *QUuid) Operator_not_equal(orig QUuid_ITF) bool {
	var convArg0 unsafe.Pointer
	if orig != nil && orig.QUuid_PTR() != nil {
		convArg0 = orig.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuidneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/quuid.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QUuid &) const

/*

 */
func (this *QUuid) Operator_less_than(other QUuid_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUuid_PTR() != nil {
		convArg0 = other.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuidltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/quuid.h:149
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator>(const QUuid &) const

/*

 */
func (this *QUuid) Operator_greater_than(other QUuid_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUuid_PTR() != nil {
		convArg0 = other.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuidgtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/quuid.h:192
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid createUuid()

/*
On any platform other than Windows, this function returns a new UUID with variant QUuid::DCE and version QUuid::Random. On Windows, a GUID is generated using the Windows API and will be of the type that the API decides to create.

See also variant() and version().
*/
func (this *QUuid) CreateUuid() *QUuid /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid10createUuidEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

/*
This function returns a new UUID with variant QUuid::DCE and version QUuid::Md5. ns is the namespace and baseData is the basic data as described by RFC 4122.

This function was introduced in  Qt 5.0.

See also variant(), version(), and createUuidV5().
*/
func (this *QUuid) CreateUuidV3(ns QUuid_ITF, baseData QByteArray_ITF) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if ns != nil && ns.QUuid_PTR() != nil {
		convArg0 = ns.QUuid_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if baseData != nil && baseData.QByteArray_PTR() != nil {
		convArg1 = baseData.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV3(ns QUuid_ITF, baseData QByteArray_ITF) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV3(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:196
// index:1
// Public static inline Visibility=Default Availability=Available
// [16] QUuid createUuidV3(const QUuid &, const QString &)

/*
This function returns a new UUID with variant QUuid::DCE and version QUuid::Md5. ns is the namespace and baseData is the basic data as described by RFC 4122.

This function was introduced in  Qt 5.0.

See also variant(), version(), and createUuidV5().
*/
func (this *QUuid) CreateUuidV3_1(ns QUuid_ITF, baseData string) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if ns != nil && ns.QUuid_PTR() != nil {
		convArg0 = ns.QUuid_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(baseData)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV3ERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV3_1(ns QUuid_ITF, baseData string) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV3_1(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:195
// index:0
// Public static Visibility=Default Availability=Available
// [16] QUuid createUuidV5(const QUuid &, const QByteArray &)

/*
This function returns a new UUID with variant QUuid::DCE and version QUuid::Sha1. ns is the namespace and baseData is the basic data as described by RFC 4122.

This function was introduced in  Qt 5.0.

See also variant(), version(), and createUuidV3().
*/
func (this *QUuid) CreateUuidV5(ns QUuid_ITF, baseData QByteArray_ITF) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if ns != nil && ns.QUuid_PTR() != nil {
		convArg0 = ns.QUuid_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if baseData != nil && baseData.QByteArray_PTR() != nil {
		convArg1 = baseData.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV5(ns QUuid_ITF, baseData QByteArray_ITF) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV5(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:201
// index:1
// Public static inline Visibility=Default Availability=Available
// [16] QUuid createUuidV5(const QUuid &, const QString &)

/*
This function returns a new UUID with variant QUuid::DCE and version QUuid::Sha1. ns is the namespace and baseData is the basic data as described by RFC 4122.

This function was introduced in  Qt 5.0.

See also variant(), version(), and createUuidV3().
*/
func (this *QUuid) CreateUuidV5_1(ns QUuid_ITF, baseData string) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if ns != nil && ns.QUuid_PTR() != nil {
		convArg0 = ns.QUuid_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(baseData)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuid12createUuidV5ERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}
func QUuid_CreateUuidV5_1(ns QUuid_ITF, baseData string) *QUuid /*123*/ {
	var nilthis *QUuid
	rv := nilthis.CreateUuidV5_1(ns, baseData)
	return rv
}

// /usr/include/qt/QtCore/quuid.h:208
// index:0
// Public Visibility=Default Availability=Available
// [4] QUuid::Variant variant() const

/*
Returns the value in the variant field of the UUID. If the return value is QUuid::DCE, call version() to see which layout it uses. The null UUID is considered to be of an unknown variant.

See also version().
*/
func (this *QUuid) Variant() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid7variantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/quuid.h:209
// index:0
// Public Visibility=Default Availability=Available
// [4] QUuid::Version version() const

/*
Returns the version field of the UUID, if the UUID's variant field is QUuid::DCE. Otherwise it returns QUuid::VerUnknown.

See also variant().
*/
func (this *QUuid) Version() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QUuid7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQUuid(this *QUuid) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QUuidD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum defines the values used in the variant field of the UUID. The value in the variant field determines the layout of the 128-bit value.


*/
type QUuid__Variant = int

//
const QUuid__VarUnknown QUuid__Variant = -1

// Reserved for NCS (Network Computing System) backward compatibility
const QUuid__NCS QUuid__Variant = 0

// Distributed Computing Environment, the scheme used by QUuid
const QUuid__DCE QUuid__Variant = 2

// Reserved for Microsoft backward compatibility (GUID)
const QUuid__Microsoft QUuid__Variant = 6

// Reserved for future definition
const QUuid__Reserved QUuid__Variant = 7

func (this *QUuid) VariantItemName(val int) string {
	switch val {
	case QUuid__VarUnknown: // -1
		return "VarUnknown"
	case QUuid__NCS: // 0
		return "NCS"
	case QUuid__DCE: // 2
		return "DCE"
	case QUuid__Microsoft: // 6
		return "Microsoft"
	case QUuid__Reserved: // 7
		return "Reserved"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QUuid_VariantItemName(val int) string {
	var nilthis *QUuid
	return nilthis.VariantItemName(val)
}

/*
This enum defines the values used in the version field of the UUID. The version field is meaningful only if the value in the variant field is QUuid::DCE.


*/
type QUuid__Version = int

//
const QUuid__VerUnknown QUuid__Version = -1

// Time-based, by using timestamp, clock sequence, and MAC network card address (if available) for the node sections
const QUuid__Time QUuid__Version = 1

// DCE Security version, with embedded POSIX UUIDs
const QUuid__EmbeddedPOSIX QUuid__Version = 2

// Alias for Name
const QUuid__Md5 QUuid__Version = 3

//
const QUuid__Name QUuid__Version = 3

// Random-based, by using random numbers for all sections
const QUuid__Random QUuid__Version = 4

//
const QUuid__Sha1 QUuid__Version = 5

func (this *QUuid) VersionItemName(val int) string {
	switch val {
	case QUuid__VerUnknown: // -1
		return "VerUnknown"
	case QUuid__Time: // 1
		return "Time"
	case QUuid__EmbeddedPOSIX: // 2
		return "EmbeddedPOSIX"
	case QUuid__Md5: // 3
		return "Md5,Name"
		// case QUuid__Name: // 3
		// return ""
	case QUuid__Random: // 4
		return "Random"
	case QUuid__Sha1: // 5
		return "Sha1"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QUuid_VersionItemName(val int) string {
	var nilthis *QUuid
	return nilthis.VersionItemName(val)
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
