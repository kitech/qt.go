package qtcore

// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QJsonValue struct {
	*qtrt.CObject
}
type QJsonValue_ITF interface {
	QJsonValue_PTR() *QJsonValue
}

func (ptr *QJsonValue) QJsonValue_PTR() *QJsonValue { return ptr }

func (this *QJsonValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonValue) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonValueFromPointer(cthis unsafe.Pointer) *QJsonValue {
	return &QJsonValue{&qtrt.CObject{cthis}}
}
func (*QJsonValue) NewFromPointer(cthis unsafe.Pointer) *QJsonValue {
	return NewQJsonValueFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonvalue.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(QJsonValue::Type)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue(arg0 int) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(QJsonValue::Type)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue__() *QJsonValue {
	// arg: 0, QJsonValue::Type=Enum, QJsonValue::Type=Enum,
	arg0 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:77
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(bool)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_1(b bool) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2Eb", qtrt.FFI_TYPE_POINTER, b)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:78
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(double)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_2(n float64) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2Ed", qtrt.FFI_TYPE_POINTER, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:79
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(int)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_3(n int) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2Ei", qtrt.FFI_TYPE_POINTER, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:80
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(qint64)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_4(n int64) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2Ex", qtrt.FFI_TYPE_POINTER, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:81
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(const QString &)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_5(s string) *QJsonValue {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:82
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(QLatin1String)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_6(s QLatin1String_ITF /*123*/) *QJsonValue {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2E13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:84
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValue(const char *)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_7(s string) *QJsonValue {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:87
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(const QJsonArray &)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_8(a QJsonArray_ITF) *QJsonValue {
	var convArg0 unsafe.Pointer
	if a != nil && a.QJsonArray_PTR() != nil {
		convArg0 = a.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2ERK10QJsonArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:88
// index:9
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(const QJsonObject &)

/*
Creates a QJsonValue of type type.

The default is to create a Null value.
*/
func NewQJsonValue_9(o QJsonObject_ITF) *QJsonValue {
	var convArg0 unsafe.Pointer
	if o != nil && o.QJsonObject_PTR() != nil {
		convArg0 = o.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2ERK11QJsonObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonValue()

/*

 */
func DeleteQJsonValue(this *QJsonValue) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonvalue.h:93
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue & operator=(const QJsonValue &)

/*

 */
func (this *QJsonValue) Operator_equal(other QJsonValue_ITF) *QJsonValue {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonValue_PTR() != nil {
		convArg0 = other.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:105
// index:1
// Public inline Visibility=Default Availability=Available
// [24] QJsonValue & operator=(QJsonValue &&)

/*

 */
func (this *QJsonValue) Operator_equal_1(other unsafe.Pointer /*333*/) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonValue &)

/*
Swaps the value other with this. This operation is very fast and never fails.

This function was introduced in  Qt 5.10.
*/
func (this *QJsonValue) Swap(other QJsonValue_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonValue_PTR() != nil {
		convArg0 = other.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValue4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [24] QJsonValue fromVariant(const QVariant &)

/*
Converts variant to a QJsonValue and returns it.

The conversion will convert QVariant types as follows:


 Source typeDestination type

QMetaType::Nullptr

QJsonValue::Null

QMetaType::Bool

QJsonValue::Bool

QMetaType::Int
QMetaType::UInt
QMetaType::LongLong
QMetaType::ULongLong
QMetaType::Float
QMetaType::Double

QJsonValue::Double

QMetaType::QString

QJsonValue::String

QMetaType::QStringList
QMetaType::QVariantList

QJsonValue::Array

QMetaType::QVariantMap
QMetaType::QVariantHash

QJsonValue::Object


For all other QVariant types a conversion to a QString will be attempted. If the returned string is empty, a Null QJsonValue will be stored, otherwise a String value using the returned QString.

See also toVariant().
*/
func (this *QJsonValue) FromVariant(variant QVariant_ITF) *QJsonValue /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValue11fromVariantERK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}
func QJsonValue_FromVariant(variant QVariant_ITF) *QJsonValue /*123*/ {
	var nilthis *QJsonValue
	rv := nilthis.FromVariant(variant)
	return rv
}

// /usr/include/qt/QtCore/qjsonvalue.h:119
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant() const

/*
Converts the value to a QVariant().

The QJsonValue types will be converted as follows:

ConstantDescription
NullQMetaType::Nullptr
BoolQMetaType::Bool
DoubleQMetaType::Double
StringQString
ArrayQVariantList
ObjectQVariantMap
UndefinedQVariant()


See also fromVariant().
*/
func (this *QJsonValue) ToVariant() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] QJsonValue::Type type() const

/*
Returns the type of the value.

See also QJsonValue::Type.
*/
func (this *QJsonValue) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the value is null.
*/
func (this *QJsonValue) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool() const

/*
Returns true if the value contains a boolean.

See also toBool().
*/
func (this *QJsonValue) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble() const

/*
Returns true if the value contains a double.

See also toDouble().
*/
func (this *QJsonValue) IsDouble() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8isDoubleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString() const

/*
Returns true if the value contains a string.

See also toString().
*/
func (this *QJsonValue) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray() const

/*
Returns true if the value contains an array.

See also toArray().
*/
func (this *QJsonValue) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isObject() const

/*
Returns true if the value contains an object.

See also toObject().
*/
func (this *QJsonValue) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndefined() const

/*
Returns true if the value is undefined. This can happen in certain error cases as e.g. accessing a non existing key in a QJsonObject.
*/
func (this *QJsonValue) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Converts the value to a bool and returns it.

If type() is not bool, the defaultValue will be returned.
*/
func (this *QJsonValue) ToBool(defaultValue bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Converts the value to a bool and returns it.

If type() is not bool, the defaultValue will be returned.
*/
func (this *QJsonValue) ToBool__() bool {
	// arg: 0, bool=Bool, =Invalid,
	defaultValue := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(int) const

/*
Converts the value to an int and returns it.

If type() is not Double or the value is not a whole number, the defaultValue will be returned.
*/
func (this *QJsonValue) ToInt(defaultValue int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue5toIntEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(int) const

/*
Converts the value to an int and returns it.

If type() is not Double or the value is not a whole number, the defaultValue will be returned.
*/
func (this *QJsonValue) ToInt__() int {
	// arg: 0, int=Int, =Invalid,
	defaultValue := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue5toIntEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Converts the value to a double and returns it.

If type() is not Double, the defaultValue will be returned.
*/
func (this *QJsonValue) ToDouble(defaultValue float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Converts the value to a double and returns it.

If type() is not Double, the defaultValue will be returned.
*/
func (this *QJsonValue) ToDouble__() float64 {
	// arg: 0, double=Double, =Invalid,
	defaultValue := float64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Converts the value to a QString and returns it.

If type() is not String, a null QString will be returned.

See also QString::isNull().
*/
func (this *QJsonValue) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qjsonvalue.h:134
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Converts the value to a QString and returns it.

If type() is not String, a null QString will be returned.

See also QString::isNull().
*/
func (this *QJsonValue) ToString_1(defaultValue string) string {
	var tmpArg0 = NewQString_5(defaultValue)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qjsonvalue.h:135
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray toArray() const

/*
Converts the value to an array and returns it.

If type() is not Array, the defaultValue will be returned.
*/
func (this *QJsonValue) ToArray() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:136
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonArray toArray(const QJsonArray &) const

/*
Converts the value to an array and returns it.

If type() is not Array, the defaultValue will be returned.
*/
func (this *QJsonValue) ToArray_1(defaultValue QJsonArray_ITF) *QJsonArray /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QJsonArray_PTR() != nil {
		convArg0 = defaultValue.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayERK10QJsonArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:137
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toObject() const

/*
Converts the value to an object and returns it.

If type() is not Object, the defaultValue will be returned.
*/
func (this *QJsonValue) ToObject() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:138
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonObject toObject(const QJsonObject &) const

/*
Converts the value to an object and returns it.

If type() is not Object, the defaultValue will be returned.
*/
func (this *QJsonValue) ToObject_1(defaultValue QJsonObject_ITF) *QJsonObject /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QJsonObject_PTR() != nil {
		convArg0 = defaultValue.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectERK11QJsonObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:140
// index:0
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](const QString &) const

/*

 */
func (this *QJsonValue) Operator_get_index(key string) *QJsonValue /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValueixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:141
// index:1
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](QLatin1String) const

/*

 */
func (this *QJsonValue) Operator_get_index_1(key QLatin1String_ITF /*123*/) *QJsonValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValueixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:142
// index:2
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](int) const

/*

 */
func (this *QJsonValue) Operator_get_index_2(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValueixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QJsonValue &) const

/*

 */
func (this *QJsonValue) Operator_equal_equal(other QJsonValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonValue_PTR() != nil {
		convArg0 = other.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValueeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonValue &) const

/*

 */
func (this *QJsonValue) Operator_not_equal(other QJsonValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonValue_PTR() != nil {
		convArg0 = other.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValueneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes the type of the JSON value.


*/
type QJsonValue__Type = int

//
const QJsonValue__Null QJsonValue__Type = 0

//
const QJsonValue__Bool QJsonValue__Type = 1

//
const QJsonValue__Double QJsonValue__Type = 2

//
const QJsonValue__String QJsonValue__Type = 3

//
const QJsonValue__Array QJsonValue__Type = 4

//
const QJsonValue__Object QJsonValue__Type = 5

//
const QJsonValue__Undefined QJsonValue__Type = 128

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
