//  header block begin
// /usr/include/qt/QtCore/qjsondocument.h
// #include <qjsondocument.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QJsonDocument struct {
	*qtrt.CObject
}

func (this *QJsonDocument) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQJsonDocumentFromPointer(cthis unsafe.Pointer) *QJsonDocument {
	return &QJsonDocument{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qjsondocument.h:88
// index:0
// Public
// void QJsonDocument()
func NewQJsonDocument() *QJsonDocument {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:89
// index:1
// Public
// void QJsonDocument(const class QJsonObject &)
func NewQJsonDocument_1(object *QJsonObject) *QJsonDocument {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK11QJsonObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:90
// index:2
// Public
// void QJsonDocument(const class QJsonArray &)
func NewQJsonDocument_2(array *QJsonArray) *QJsonDocument {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = array.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK10QJsonArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:91
// index:0
// Public
// void ~QJsonDocument()
func DeleteQJsonDocument(*QJsonDocument) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:108
// index:0
// Public inline
// void swap(class QJsonDocument &)
func (this *QJsonDocument) Swap(other *QJsonDocument) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// Public static
// QJsonDocument fromRawData(const char *, int, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromRawData(data string, size int, validation int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", ffiqt.FFI_TYPE_POINTER, data, size, validation)
	gopp.ErrPrint(err, rv)
	return rv
}
func QJsonDocument_FromRawData(data string, size int, validation int) {
	var nilthis *QJsonDocument
	nilthis.FromRawData(data, size, validation)
}

// /usr/include/qt/QtCore/qjsondocument.h:119
// index:0
// Public
// const char * rawData(int *)
func (this *QJsonDocument) RawData(size unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7rawDataEPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// Public static
// QJsonDocument fromBinaryData(const class QByteArray &, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromBinaryData(data *QByteArray, validation int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", ffiqt.FFI_TYPE_POINTER, data, validation)
	gopp.ErrPrint(err, rv)
	return rv
}
func QJsonDocument_FromBinaryData(data *QByteArray, validation int) {
	var nilthis *QJsonDocument
	nilthis.FromBinaryData(data, validation)
}

// /usr/include/qt/QtCore/qjsondocument.h:122
// index:0
// Public
// QByteArray toBinaryData()
func (this *QJsonDocument) ToBinaryData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument12toBinaryDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:124
// index:0
// Public static
// QJsonDocument fromVariant(const class QVariant &)
func (this *QJsonDocument) FromVariant(variant *QVariant) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument11fromVariantERK8QVariant", ffiqt.FFI_TYPE_POINTER, variant)
	gopp.ErrPrint(err, rv)
	return rv
}
func QJsonDocument_FromVariant(variant *QVariant) {
	var nilthis *QJsonDocument
	nilthis.FromVariant(variant)
}

// /usr/include/qt/QtCore/qjsondocument.h:125
// index:0
// Public
// QVariant toVariant()
func (this *QJsonDocument) ToVariant() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument9toVariantEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:132
// index:0
// Public static
// QJsonDocument fromJson(const class QByteArray &, struct QJsonParseError *)
func (this *QJsonDocument) FromJson(json *QByteArray, error unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument8fromJsonERK10QByteArrayP15QJsonParseError", ffiqt.FFI_TYPE_POINTER, json, error)
	gopp.ErrPrint(err, rv)
	return rv
}
func QJsonDocument_FromJson(json *QByteArray, error unsafe.Pointer) {
	var nilthis *QJsonDocument
	nilthis.FromJson(json, error)
}

// /usr/include/qt/QtCore/qjsondocument.h:137
// index:0
// Public
// QByteArray toJson()
func (this *QJsonDocument) ToJson() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:138
// index:1
// Public
// QByteArray toJson(enum QJsonDocument::JsonFormat)
func (this *QJsonDocument) ToJson_1(format int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonENS_10JsonFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:141
// index:0
// Public
// bool isEmpty()
func (this *QJsonDocument) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:142
// index:0
// Public
// bool isArray()
func (this *QJsonDocument) IsArray() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7isArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:143
// index:0
// Public
// bool isObject()
func (this *QJsonDocument) IsObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument8isObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:145
// index:0
// Public
// QJsonObject object()
func (this *QJsonDocument) Object() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:146
// index:0
// Public
// QJsonArray array()
func (this *QJsonDocument) Array() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument5arrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:148
// index:0
// Public
// void setObject(const class QJsonObject &)
func (this *QJsonDocument) SetObject(object *QJsonObject) {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument9setObjectERK11QJsonObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:149
// index:0
// Public
// void setArray(const class QJsonArray &)
func (this *QJsonDocument) SetArray(array *QJsonArray) {
	var convArg0 = array.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument8setArrayERK10QJsonArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:158
// index:0
// Public
// bool isNull()
func (this *QJsonDocument) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
