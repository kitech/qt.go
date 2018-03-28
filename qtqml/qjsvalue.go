package qtqml

// /usr/include/qt/QtQml/qjsvalue.h
// #include <qjsvalue.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QJSValue struct {
	*qtrt.CObject
}
type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue { return ptr }

func (this *QJSValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJSValue) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJSValueFromPointer(cthis unsafe.Pointer) *QJSValue {
	return &QJSValue{&qtrt.CObject{cthis}}
}
func (*QJSValue) NewFromPointer(cthis unsafe.Pointer) *QJSValue {
	return NewQJSValueFromPointer(cthis)
}

// /usr/include/qt/QtQml/qjsvalue.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(QJSValue::SpecialValue)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue(value int) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ENS_12SpecialValueE", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(QJSValue::SpecialValue)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue__() *QJSValue {
	// arg: 0, QJSValue::SpecialValue=Enum, QJSValue::SpecialValue=Enum,
	value := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ENS_12SpecialValueE", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:82
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(bool)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_1(value bool) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Eb", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:83
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(int)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_2(value int) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Ei", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:84
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(uint)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_3(value uint) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Ej", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:85
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(double)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_4(value float64) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Ed", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:86
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(const QString &)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_5(value string) *QJSValue {
	var tmpArg0 = qtcore.NewQString_5(value)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:87
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(const QLatin1String &)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_6(value qtcore.QLatin1String_ITF) *QJSValue {
	var convArg0 unsafe.Pointer
	if value != nil && value.QLatin1String_PTR() != nil {
		convArg0 = value.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ERK13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:89
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(const char *)

/*
Constructs a new QJSValue with a special value.
*/
func NewQJSValue_7(str string) *QJSValue {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJSValue()

/*

 */
func DeleteQJSValue(this *QJSValue) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qjsvalue.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QJSValue & operator=(QJSValue &&)

/*

 */
func (this *QJSValue) Operator_equal(other unsafe.Pointer /*333*/) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:92
// index:1
// Public Visibility=Default Availability=Available
// [8] QJSValue & operator=(const QJSValue &)

/*

 */
func (this *QJSValue) Operator_equal_1(other QJSValue_ITF) *QJSValue {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJSValue_PTR() != nil {
		convArg0 = other.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBool() const

/*
Returns true if this QJSValue is of the primitive type Boolean; otherwise returns false.

See also toBool().
*/
func (this *QJSValue) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNumber() const

/*
Returns true if this QJSValue is of the primitive type Number; otherwise returns false.

See also toNumber().
*/
func (this *QJSValue) IsNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this QJSValue is of the primitive type Null; otherwise returns false.
*/
func (this *QJSValue) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isString() const

/*
Returns true if this QJSValue is of the primitive type String; otherwise returns false.

See also toString().
*/
func (this *QJSValue) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndefined() const

/*
Returns true if this QJSValue is of the primitive type Undefined; otherwise returns false.
*/
func (this *QJSValue) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVariant() const

/*
Returns true if this QJSValue is a variant value; otherwise returns false.

See also toVariant().
*/
func (this *QJSValue) IsVariant() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9isVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isQObject() const

/*
Returns true if this QJSValue is a QObject; otherwise returns false.

Note: This function returns true even if the QObject that this QJSValue wraps has been deleted.

See also toQObject() and QJSEngine::newQObject().
*/
func (this *QJSValue) IsQObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9isQObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isQMetaObject() const

/*
Returns true if this QJSValue is a QMetaObject; otherwise returns false.

This function was introduced in  Qt 5.8.

See also toQMetaObject() and QJSEngine::newQMetaObject().
*/
func (this *QJSValue) IsQMetaObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue13isQMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObject() const

/*
Returns true if this QJSValue is of the Object type; otherwise returns false.

Note that function values, variant values, and QObject values are objects, so this function returns true for such values.

See also QJSEngine::newObject().
*/
func (this *QJSValue) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDate() const

/*
Returns true if this QJSValue is an object of the Date class; otherwise returns false.
*/
func (this *QJSValue) IsDate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6isDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRegExp() const

/*
Returns true if this QJSValue is an object of the RegExp class; otherwise returns false.
*/
func (this *QJSValue) IsRegExp() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isRegExpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isArray() const

/*
Returns true if this QJSValue is an object of the Array class; otherwise returns false.

See also QJSEngine::newArray().
*/
func (this *QJSValue) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isError() const

/*
Returns true if this QJSValue is an object of the Error class; otherwise returns false.

See also QJSEngine - Script Exceptions.
*/
func (this *QJSValue) IsError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue7isErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the string value of this QJSValue, as defined in ECMA-262 section 9.8, "ToString".

Note that if this QJSValue is an object, calling this function has side effects on the script engine, since the engine will call the object's toString() function (and possibly valueOf()) in an attempt to convert the object to a primitive value (possibly resulting in an uncaught script exception).

See also isString().
*/
func (this *QJSValue) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qjsvalue.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] double toNumber() const

/*
Returns the number value of this QJSValue, as defined in ECMA-262 section 9.3, "ToNumber".

Note that if this QJSValue is an object, calling this function has side effects on the script engine, since the engine will call the object's valueOf() function (and possibly toString()) in an attempt to convert the object to a primitive value (possibly resulting in an uncaught script exception).

See also isNumber(), toInt(), and toUInt().
*/
func (this *QJSValue) ToNumber() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8toNumberEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qjsvalue.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] qint32 toInt() const

/*
Returns the signed 32-bit integer value of this QJSValue, using the conversion rules described in ECMA-262 section 9.5, "ToInt32".

Note that if this QJSValue is an object, calling this function has side effects on the script engine, since the engine will call the object's valueOf() function (and possibly toString()) in an attempt to convert the object to a primitive value (possibly resulting in an uncaught script exception).

See also toNumber() and toUInt().
*/
func (this *QJSValue) ToInt() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue5toIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv) // 222
}

// /usr/include/qt/QtQml/qjsvalue.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 toUInt() const

/*
Returns the unsigned 32-bit integer value of this QJSValue, using the conversion rules described in ECMA-262 section 9.6, "ToUint32".

Note that if this QJSValue is an object, calling this function has side effects on the script engine, since the engine will call the object's valueOf() function (and possibly toString()) in an attempt to convert the object to a primitive value (possibly resulting in an uncaught script exception).

See also toNumber() and toInt().
*/
func (this *QJSValue) ToUInt() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6toUIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtQml/qjsvalue.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool() const

/*
Returns the boolean value of this QJSValue, using the conversion rules described in ECMA-262 section 9.2, "ToBoolean".

Note that if this QJSValue is an object, calling this function has side effects on the script engine, since the engine will call the object's valueOf() function (and possibly toString()) in an attempt to convert the object to a primitive value (possibly resulting in an uncaught script exception).

See also isBool().
*/
func (this *QJSValue) ToBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6toBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:113
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant() const

/*
Returns the QVariant value of this QJSValue, if it can be converted to a QVariant; otherwise returns an invalid QVariant. The conversion is performed according to the following table:


 Input TypeResult
UndefinedAn invalid QVariant.
NullA QVariant containing a null pointer (QMetaType::Nullptr).
BooleanA QVariant containing the value of the boolean.
NumberA QVariant containing the value of the number.
StringA QVariant containing the value of the string.
QVariant ObjectThe result is the QVariant value of the object (no conversion).
QObject ObjectA QVariant containing a pointer to the QObject.
Date ObjectA QVariant containing the date value (toDateTime()).
RegExp ObjectA QVariant containing the regular expression value.
Array ObjectThe array is converted to a QVariantList. Each element is converted to a QVariant, recursively; cyclic references are not followed.
ObjectThe object is converted to a QVariantMap. Each property is converted to a QVariant, recursively; cyclic references are not followed.


See also isVariant().
*/
func (this *QJSValue) ToVariant() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * toQObject() const

/*
If this QJSValue is a QObject, returns the QObject pointer that the QJSValue represents; otherwise, returns 0.

If the QObject that this QJSValue wraps has been deleted, this function returns 0 (i.e. it is possible for toQObject() to return 0 even when isQObject() returns true).

See also isQObject().
*/
func (this *QJSValue) ToQObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9toQObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qjsvalue.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMetaObject * toQMetaObject() const

/*
* If this QJSValue is a QMetaObject, returns the QMetaObject pointer * that the QJSValue represents; otherwise, returns 0. * *

This function was introduced in  Qt 5.8.

See also isQMetaObject().
*/
func (this *QJSValue) ToQMetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue13toQMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qjsvalue.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime() const

/*
Returns a QDateTime representation of this value, in local time. If this QJSValue is not a date, or the value of the date is NaN (Not-a-Number), an invalid QDateTime is returned.

See also isDate().
*/
func (this *QJSValue) ToDateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue10toDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool equals(const QJSValue &) const

/*
Returns true if this QJSValue is equal to other, otherwise returns false. The comparison follows the behavior described in ECMA-262 section 11.9.3, "The Abstract Equality Comparison Algorithm".

This function can return true even if the type of this QJSValue is different from the type of the other value; i.e. the comparison is not strict. For example, comparing the number 9 to the string "9" returns true; comparing an undefined value to a null value returns true; comparing a Number object whose primitive value is 6 to a String object whose primitive value is "6" returns true; and comparing the number 1 to the boolean value true returns true. If you want to perform a comparison without such implicit value conversion, use strictlyEquals().

Note that if this QJSValue or the other value are objects, calling this function has side effects on the script engine, since the engine will call the object's valueOf() function (and possibly toString()) in an attempt to convert the object to a primitive value (possibly resulting in an uncaught script exception).

See also strictlyEquals().
*/
func (this *QJSValue) Equals(other QJSValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJSValue_PTR() != nil {
		convArg0 = other.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6equalsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strictlyEquals(const QJSValue &) const

/*
Returns true if this QJSValue is equal to other using strict comparison (no conversion), otherwise returns false. The comparison follows the behavior described in ECMA-262 section 11.9.6, "The Strict Equality Comparison Algorithm".

If the type of this QJSValue is different from the type of the other value, this function returns false. If the types are equal, the result depends on the type, as shown in the following table:


 TypeResult
Undefinedtrue
Nulltrue
Booleantrue if both values are true, false otherwise
Numberfalse if either value is NaN (Not-a-Number); true if values are equal, false otherwise
Stringtrue if both values are exactly the same sequence of characters, false otherwise
Objecttrue if both values refer to the same object, false otherwise


See also equals().
*/
func (this *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJSValue_PTR() != nil {
		convArg0 = other.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue14strictlyEqualsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue prototype() const

/*
If this QJSValue is an object, returns the internal prototype (__proto__ property) of this object; otherwise returns an undefined QJSValue.

See also setPrototype() and isObject().
*/
func (this *QJSValue) Prototype() *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9prototypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrototype(const QJSValue &)

/*
If this QJSValue is an object, sets the internal prototype (__proto__ property) of this object to be prototype; if the QJSValue is null, it sets the prototype to null; otherwise does nothing.

The internal prototype should not be confused with the public property with name "prototype"; the public prototype is usually only set on functions that act as constructors.

See also prototype() and isObject().
*/
func (this *QJSValue) SetPrototype(prototype QJSValue_ITF) {
	var convArg0 unsafe.Pointer
	if prototype != nil && prototype.QJSValue_PTR() != nil {
		convArg0 = prototype.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue12setPrototypeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue property(const QString &) const

/*
Returns the value of this QJSValue's property with the given name. If no such property exists, an undefined QJSValue is returned.

If the property is implemented using a getter function (i.e. has the PropertyGetter flag set), calling property() has side-effects on the script engine, since the getter function will be called (possibly resulting in an uncaught script exception). If an exception occurred, property() returns the value that was thrown (typically an Error object).

See also setProperty(), hasProperty(), and QJSValueIterator.
*/
func (this *QJSValue) Property(name string) *QJSValue /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8propertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:130
// index:1
// Public Visibility=Default Availability=Available
// [8] QJSValue property(quint32) const

/*
Returns the value of this QJSValue's property with the given name. If no such property exists, an undefined QJSValue is returned.

If the property is implemented using a getter function (i.e. has the PropertyGetter flag set), calling property() has side-effects on the script engine, since the getter function will be called (possibly resulting in an uncaught script exception). If an exception occurred, property() returns the value that was thrown (typically an Error object).

See also setProperty(), hasProperty(), and QJSValueIterator.
*/
func (this *QJSValue) Property_1(arrayIndex uint) *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8propertyEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arrayIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProperty(const QString &, const QJSValue &)

/*
Sets the value of this QJSValue's property with the given name to the given value.

If this QJSValue is not an object, this function does nothing.

If this QJSValue does not already have a property with name name, a new property is created.

See also property() and deleteProperty().
*/
func (this *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QJSValue_PTR() != nil {
		convArg1 = value.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue11setPropertyERK7QStringRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:131
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setProperty(quint32, const QJSValue &)

/*
Sets the value of this QJSValue's property with the given name to the given value.

If this QJSValue is not an object, this function does nothing.

If this QJSValue does not already have a property with name name, a new property is created.

See also property() and deleteProperty().
*/
func (this *QJSValue) SetProperty_1(arrayIndex uint, value QJSValue_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QJSValue_PTR() != nil {
		convArg1 = value.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue11setPropertyEjRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arrayIndex, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasProperty(const QString &) const

/*
Returns true if this object has a property of the given name, otherwise returns false.

See also property() and hasOwnProperty().
*/
func (this *QJSValue) HasProperty(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue11hasPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasOwnProperty(const QString &) const

/*
Returns true if this object has an own (not prototype-inherited) property of the given name, otherwise returns false.

See also property() and hasProperty().
*/
func (this *QJSValue) HasOwnProperty(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue14hasOwnPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool deleteProperty(const QString &)

/*
Attempts to delete this object's property of the given name. Returns true if the property was deleted, otherwise returns false.

The behavior of this function is consistent with the JavaScript delete operator. In particular:


Non-configurable properties cannot be deleted.
This function will return true even if this object doesn't have a property of the given name (i.e., non-existent properties are "trivially deletable").
If this object doesn't have an own property of the given name, but an object in the prototype() chain does, the prototype object's property is not deleted, and this function returns true.


See also setProperty() and hasOwnProperty().
*/
func (this *QJSValue) DeleteProperty(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue14deletePropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCallable() const

/*
Returns true if this QJSValue can be called a function, otherwise returns false.

See also call().
*/
func (this *QJSValue) IsCallable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue10isCallableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSEngine * engine() const

/*

 */
func (this *QJSValue) Engine() *QJSEngine /*777 QJSEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

/*
This enum is used to specify a single-valued type.


*/
type QJSValue__SpecialValue = int

// A null value.
const QJSValue__NullValue QJSValue__SpecialValue = 0

// An undefined value.
const QJSValue__UndefinedValue QJSValue__SpecialValue = 1

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
