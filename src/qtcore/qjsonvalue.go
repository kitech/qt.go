package qtcore

// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QJsonValue struct {
	*qtrt.CObject
}

func (this *QJsonValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonValue) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQJsonValueFromPointer(cthis unsafe.Pointer) *QJsonValue {
	return &QJsonValue{&qtrt.CObject{cthis}}
}
func (*QJsonValue) NewFromPointer(cthis unsafe.Pointer) *QJsonValue {
	return NewQJsonValueFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonvalue.h:76
// index:0
// Public
// void QJsonValue(enum QJsonValue::Type)
func NewQJsonValue(arg0 int) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:77
// index:1
// Public
// void QJsonValue(_Bool)
func NewQJsonValue_1(b bool) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Eb", ffiqt.FFI_TYPE_VOID, cthis, b)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:78
// index:2
// Public
// void QJsonValue(double)
func NewQJsonValue_2(n float64) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Ed", ffiqt.FFI_TYPE_VOID, cthis, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:79
// index:3
// Public
// void QJsonValue(int)
func NewQJsonValue_3(n int) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Ei", ffiqt.FFI_TYPE_VOID, cthis, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:80
// index:4
// Public
// void QJsonValue(qint64)
func NewQJsonValue_4(n int64) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Ex", ffiqt.FFI_TYPE_VOID, cthis, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:81
// index:5
// Public
// void QJsonValue(const class QString &)
func NewQJsonValue_5(s *QString) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:82
// index:6
// Public
// void QJsonValue(class QLatin1String)
func NewQJsonValue_6(s *QLatin1String /*123*/) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2E13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:84
// index:7
// Public inline
// void QJsonValue(const char *)
func NewQJsonValue_7(s string) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:87
// index:8
// Public
// void QJsonValue(const class QJsonArray &)
func NewQJsonValue_8(a *QJsonArray) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ERK10QJsonArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:88
// index:9
// Public
// void QJsonValue(const class QJsonObject &)
func NewQJsonValue_9(o *QJsonObject) *QJsonValue {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = o.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ERK11QJsonObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:90
// index:0
// Public
// void ~QJsonValue()
func DeleteQJsonValue(*QJsonValue) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:111
// index:0
// Public inline
// void swap(class QJsonValue &)
func (this *QJsonValue) Swap(other *QJsonValue) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValue4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:118
// index:0
// Public static
// QJsonValue fromVariant(const class QVariant &)
func (this *QJsonValue) FromVariant(variant *QVariant) *QJsonValue /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValue11fromVariantERK8QVariant", ffiqt.FFI_TYPE_POINTER, variant)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QJsonValue_FromVariant(variant *QVariant) *QJsonValue /*123*/ {
	var nilthis *QJsonValue
	rv := nilthis.FromVariant(variant)
	return rv
}

// /usr/include/qt/QtCore/qjsonvalue.h:119
// index:0
// Public
// QVariant toVariant()
func (this *QJsonValue) ToVariant() *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue9toVariantEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:121
// index:0
// Public
// QJsonValue::Type type()
func (this *QJsonValue) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:122
// index:0
// Public inline
// bool isNull()
func (this *QJsonValue) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:123
// index:0
// Public inline
// bool isBool()
func (this *QJsonValue) IsBool() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue6isBoolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:124
// index:0
// Public inline
// bool isDouble()
func (this *QJsonValue) IsDouble() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8isDoubleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:125
// index:0
// Public inline
// bool isString()
func (this *QJsonValue) IsString() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8isStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:126
// index:0
// Public inline
// bool isArray()
func (this *QJsonValue) IsArray() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue7isArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:127
// index:0
// Public inline
// bool isObject()
func (this *QJsonValue) IsObject() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8isObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:128
// index:0
// Public inline
// bool isUndefined()
func (this *QJsonValue) IsUndefined() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue11isUndefinedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:130
// index:0
// Public
// bool toBool(_Bool)
func (this *QJsonValue) ToBool(defaultValue bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue6toBoolEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:131
// index:0
// Public
// int toInt(int)
func (this *QJsonValue) ToInt(defaultValue int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue5toIntEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qjsonvalue.h:132
// index:0
// Public
// double toDouble(double)
func (this *QJsonValue) ToDouble(defaultValue float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toDoubleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qjsonvalue.h:133
// index:0
// Public
// QString toString()
func (this *QJsonValue) ToString() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:134
// index:1
// Public
// QString toString(const class QString &)
func (this *QJsonValue) ToString_1(defaultValue *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toStringERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:135
// index:0
// Public
// QJsonArray toArray()
func (this *QJsonValue) ToArray() *QJsonArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:136
// index:1
// Public
// QJsonArray toArray(const class QJsonArray &)
func (this *QJsonValue) ToArray_1(defaultValue *QJsonArray) *QJsonArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayERK10QJsonArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:137
// index:0
// Public
// QJsonObject toObject()
func (this *QJsonValue) ToObject() *QJsonObject /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:138
// index:1
// Public
// QJsonObject toObject(const class QJsonObject &)
func (this *QJsonValue) ToObject_1(defaultValue *QJsonObject) *QJsonObject /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectERK11QJsonObject", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
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
