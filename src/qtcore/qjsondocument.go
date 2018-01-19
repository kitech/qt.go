//  header block begin
// /usr/include/qt/QtCore/qjsondocument.h
// #include <qjsondocument.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qjsondocument.h:88
// index:0
// void QJsonDocument()
func NewQJsonDocument() *QJsonDocument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QJsonDocument{cthis}
}

// /usr/include/qt/QtCore/qjsondocument.h:89
// index:1
// void QJsonDocument(const class QJsonObject &)
func NewQJsonDocument_1(object unsafe.Pointer) *QJsonDocument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK11QJsonObject", ffiqt.FFI_TYPE_VOID, cthis, object)
	gopp.ErrPrint(err, rv)
	return &QJsonDocument{cthis}
}

// /usr/include/qt/QtCore/qjsondocument.h:90
// index:2
// void QJsonDocument(const class QJsonArray &)
func NewQJsonDocument_2(array unsafe.Pointer) *QJsonDocument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK10QJsonArray", ffiqt.FFI_TYPE_VOID, cthis, array)
	gopp.ErrPrint(err, rv)
	return &QJsonDocument{cthis}
}

// /usr/include/qt/QtCore/qjsondocument.h:91
// index:0
// void ~QJsonDocument()
func DeleteQJsonDocument(*QJsonDocument) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:108
// index:0
// inline
// void swap(class QJsonDocument &)
func (this *QJsonDocument) Swap(other unsafe.Pointer) {
	// 0: (, QJsonDocument & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// static
// QJsonDocument fromRawData(const char *, int, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromRawData(data unsafe.Pointer, size int, validation int) {
	// 0: (const char * data, int size, QJsonDocument::DataValidation validation), (data, size, validation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QJsonDocument_FromRawData(data unsafe.Pointer, size int, validation int) {
	// 0: (const char * data, int size, QJsonDocument::DataValidation validation), (data, size, validation)
	var nilthis *QJsonDocument
	nilthis.FromRawData(data, size, validation)
}

// /usr/include/qt/QtCore/qjsondocument.h:119
// index:0
// const char * rawData(int *)
func (this *QJsonDocument) RawData(size unsafe.Pointer) {
	// 0: (, int * size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7rawDataEPi", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// static
// QJsonDocument fromBinaryData(const class QByteArray &, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromBinaryData(data unsafe.Pointer, validation int) {
	// 0: (const QByteArray & data, QJsonDocument::DataValidation validation), (data, validation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QJsonDocument_FromBinaryData(data unsafe.Pointer, validation int) {
	// 0: (const QByteArray & data, QJsonDocument::DataValidation validation), (data, validation)
	var nilthis *QJsonDocument
	nilthis.FromBinaryData(data, validation)
}

// /usr/include/qt/QtCore/qjsondocument.h:122
// index:0
// QByteArray toBinaryData()
func (this *QJsonDocument) ToBinaryData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument12toBinaryDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:124
// index:0
// static
// QJsonDocument fromVariant(const class QVariant &)
func (this *QJsonDocument) FromVariant(variant unsafe.Pointer) {
	// 0: (const QVariant & variant), (variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument11fromVariantERK8QVariant", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QJsonDocument_FromVariant(variant unsafe.Pointer) {
	// 0: (const QVariant & variant), (variant)
	var nilthis *QJsonDocument
	nilthis.FromVariant(variant)
}

// /usr/include/qt/QtCore/qjsondocument.h:125
// index:0
// QVariant toVariant()
func (this *QJsonDocument) ToVariant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument9toVariantEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:132
// index:0
// static
// QJsonDocument fromJson(const class QByteArray &, struct QJsonParseError *)
func (this *QJsonDocument) FromJson(json unsafe.Pointer, error unsafe.Pointer) {
	// 0: (const QByteArray & json, QJsonParseError * error), (json, error)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument8fromJsonERK10QByteArrayP15QJsonParseError", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QJsonDocument_FromJson(json unsafe.Pointer, error unsafe.Pointer) {
	// 0: (const QByteArray & json, QJsonParseError * error), (json, error)
	var nilthis *QJsonDocument
	nilthis.FromJson(json, error)
}

// /usr/include/qt/QtCore/qjsondocument.h:137
// index:0
// QByteArray toJson()
func (this *QJsonDocument) ToJson() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:138
// index:1
// QByteArray toJson(enum QJsonDocument::JsonFormat)
func (this *QJsonDocument) ToJson_1(format int) {
	// 1: (, QJsonDocument::JsonFormat format), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonENS_10JsonFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:141
// index:0
// bool isEmpty()
func (this *QJsonDocument) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:142
// index:0
// bool isArray()
func (this *QJsonDocument) IsArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7isArrayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:143
// index:0
// bool isObject()
func (this *QJsonDocument) IsObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument8isObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:145
// index:0
// QJsonObject object()
func (this *QJsonDocument) Object() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6objectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:146
// index:0
// QJsonArray array()
func (this *QJsonDocument) Array() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument5arrayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:148
// index:0
// void setObject(const class QJsonObject &)
func (this *QJsonDocument) SetObject(object unsafe.Pointer) {
	// 0: (, const QJsonObject & object), (object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument9setObjectERK11QJsonObject", ffiqt.FFI_TYPE_VOID, this.cthis, object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:149
// index:0
// void setArray(const class QJsonArray &)
func (this *QJsonDocument) SetArray(array unsafe.Pointer) {
	// 0: (, const QJsonArray & array), (array)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument8setArrayERK10QJsonArray", ffiqt.FFI_TYPE_VOID, this.cthis, array)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:158
// index:0
// bool isNull()
func (this *QJsonDocument) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
