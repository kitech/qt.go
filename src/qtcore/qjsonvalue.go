//  header block begin
// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
	return this.Cthis
}

// /usr/include/qt/QtCore/qjsonvalue.h:76
// index:0
// void QJsonValue(enum QJsonValue::Type)
func NewQJsonValue(arg0 int) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}
func NewQJsonValueFromPointer(cthis unsafe.Pointer) *QJsonValue {
	return &QJsonValue{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qjsonvalue.h:77
// index:1
// void QJsonValue(_Bool)
func NewQJsonValue_1(b bool) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Eb", ffiqt.FFI_TYPE_VOID, cthis, &b)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:78
// index:2
// void QJsonValue(double)
func NewQJsonValue_2(n float64) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Ed", ffiqt.FFI_TYPE_VOID, cthis, &n)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:79
// index:3
// void QJsonValue(int)
func NewQJsonValue_3(n int) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &n)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:80
// index:4
// void QJsonValue(qint64)
func NewQJsonValue_4(n int64) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2Ex", ffiqt.FFI_TYPE_VOID, cthis, &n)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:81
// index:5
// void QJsonValue(const class QString &)
func NewQJsonValue_5(s unsafe.Pointer) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:82
// index:6
// void QJsonValue(class QLatin1String)
func NewQJsonValue_6(s unsafe.Pointer) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2E13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:84
// index:7
// inline
// void QJsonValue(const char *)
func NewQJsonValue_7(s unsafe.Pointer) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:87
// index:8
// void QJsonValue(const class QJsonArray &)
func NewQJsonValue_8(a unsafe.Pointer) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ERK10QJsonArray", ffiqt.FFI_TYPE_VOID, cthis, a)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:88
// index:9
// void QJsonValue(const class QJsonObject &)
func NewQJsonValue_9(o unsafe.Pointer) *QJsonValue {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueC2ERK11QJsonObject", ffiqt.FFI_TYPE_VOID, cthis, o)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:90
// index:0
// void ~QJsonValue()
func DeleteQJsonValue(*QJsonValue) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValueD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:111
// index:0
// inline
// void swap(class QJsonValue &)
func (this *QJsonValue) Swap(other unsafe.Pointer) {
	// 0: (, other QJsonValue &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValue4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:118
// index:0
// static
// QJsonValue fromVariant(const class QVariant &)
func (this *QJsonValue) FromVariant(variant unsafe.Pointer) {
	// 0: (variant const QVariant &), (variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonValue11fromVariantERK8QVariant", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QJsonValue_FromVariant(variant unsafe.Pointer) {
	// 0: (variant const QVariant &), (variant)
	var nilthis *QJsonValue
	nilthis.FromVariant(variant)
}

// /usr/include/qt/QtCore/qjsonvalue.h:119
// index:0
// QVariant toVariant()
func (this *QJsonValue) ToVariant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue9toVariantEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:121
// index:0
// QJsonValue::Type type()
func (this *QJsonValue) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:122
// index:0
// inline
// bool isNull()
func (this *QJsonValue) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:123
// index:0
// inline
// bool isBool()
func (this *QJsonValue) IsBool() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue6isBoolEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:124
// index:0
// inline
// bool isDouble()
func (this *QJsonValue) IsDouble() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8isDoubleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:125
// index:0
// inline
// bool isString()
func (this *QJsonValue) IsString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8isStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:126
// index:0
// inline
// bool isArray()
func (this *QJsonValue) IsArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue7isArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:127
// index:0
// inline
// bool isObject()
func (this *QJsonValue) IsObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8isObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:128
// index:0
// inline
// bool isUndefined()
func (this *QJsonValue) IsUndefined() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue11isUndefinedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:130
// index:0
// bool toBool(_Bool)
func (this *QJsonValue) ToBool(defaultValue bool) {
	// 0: (, defaultValue bool), (&defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue6toBoolEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:131
// index:0
// int toInt(int)
func (this *QJsonValue) ToInt(defaultValue int) {
	// 0: (, defaultValue int), (&defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue5toIntEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:132
// index:0
// double toDouble(double)
func (this *QJsonValue) ToDouble(defaultValue float64) {
	// 0: (, defaultValue double), (&defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toDoubleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:133
// index:0
// QString toString()
func (this *QJsonValue) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:134
// index:1
// QString toString(const class QString &)
func (this *QJsonValue) ToString_1(defaultValue unsafe.Pointer) {
	// 1: (, defaultValue const QString &), (defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toStringERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:135
// index:0
// QJsonArray toArray()
func (this *QJsonValue) ToArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:136
// index:1
// QJsonArray toArray(const class QJsonArray &)
func (this *QJsonValue) ToArray_1(defaultValue unsafe.Pointer) {
	// 1: (, defaultValue const QJsonArray &), (defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue7toArrayERK10QJsonArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:137
// index:0
// QJsonObject toObject()
func (this *QJsonValue) ToObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:138
// index:1
// QJsonObject toObject(const class QJsonObject &)
func (this *QJsonValue) ToObject_1(defaultValue unsafe.Pointer) {
	// 1: (, defaultValue const QJsonObject &), (defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonValue8toObjectERK11QJsonObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
}

//  body block end
