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
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

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
// [-2] void QJsonValue(enum QJsonValue::Type)
func NewQJsonValue(arg0 int) *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValue)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:77
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJsonValue(_Bool)
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
func NewQJsonValue_6(s QLatin1String_ITF /*123*/) *QJsonValue {
	var convArg0 = s.QLatin1String_PTR().GetCthis()
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
func NewQJsonValue_8(a QJsonArray_ITF) *QJsonValue {
	var convArg0 = a.QJsonArray_PTR().GetCthis()
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
func NewQJsonValue_9(o QJsonObject_ITF) *QJsonValue {
	var convArg0 = o.QJsonObject_PTR().GetCthis()
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
func DeleteQJsonValue(this *QJsonValue) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValueD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonvalue.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonValue &)
func (this *QJsonValue) Swap(other QJsonValue_ITF) {
	var convArg0 = other.QJsonValue_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonValue4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [24] QJsonValue fromVariant(const QVariant &)
func (this *QJsonValue) FromVariant(variant QVariant_ITF) *QJsonValue /*123*/ {
	var convArg0 = variant.QVariant_PTR().GetCthis()
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
// [16] QVariant toVariant()
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
// [4] QJsonValue::Type type()
func (this *QJsonValue) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QJsonValue) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool()
func (this *QJsonValue) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble()
func (this *QJsonValue) IsDouble() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8isDoubleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString()
func (this *QJsonValue) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray()
func (this *QJsonValue) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isObject()
func (this *QJsonValue) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndefined()
func (this *QJsonValue) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool(_Bool)
func (this *QJsonValue) ToBool(defaultValue bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(int)
func (this *QJsonValue) ToInt(defaultValue int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue5toIntEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(double)
func (this *QJsonValue) ToDouble(defaultValue float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
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
// [8] QString toString(const QString &)
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
// [16] QJsonArray toArray()
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
// [16] QJsonArray toArray(const QJsonArray &)
func (this *QJsonValue) ToArray_1(defaultValue QJsonArray_ITF) *QJsonArray /*123*/ {
	var convArg0 = defaultValue.QJsonArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayERK10QJsonArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:137
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toObject()
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
// [16] QJsonObject toObject(const QJsonObject &)
func (this *QJsonValue) ToObject_1(defaultValue QJsonObject_ITF) *QJsonObject /*123*/ {
	var convArg0 = defaultValue.QJsonObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectERK11QJsonObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

type QJsonValue__Type = int

const QJsonValue__Null QJsonValue__Type = 0
const QJsonValue__Bool QJsonValue__Type = 1
const QJsonValue__Double QJsonValue__Type = 2
const QJsonValue__String QJsonValue__Type = 3
const QJsonValue__Array QJsonValue__Type = 4
const QJsonValue__Object QJsonValue__Type = 5
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
		qtrt.KeepMe()
	}
}

//  keep block end
