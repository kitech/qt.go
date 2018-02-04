package qtcore

// /usr/include/qt/QtCore/qjsondocument.h
// #include <qjsondocument.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonDocument) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonDocumentFromPointer(cthis unsafe.Pointer) *QJsonDocument {
	return &QJsonDocument{&qtrt.CObject{cthis}}
}
func (*QJsonDocument) NewFromPointer(cthis unsafe.Pointer) *QJsonDocument {
	return NewQJsonDocumentFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsondocument.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument()
func NewQJsonDocument() *QJsonDocument {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument(const QJsonObject &)
func NewQJsonDocument_1(object *QJsonObject) *QJsonDocument {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK11QJsonObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:90
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument(const QJsonArray &)
func NewQJsonDocument_2(array *QJsonArray) *QJsonDocument {
	var convArg0 = array.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK10QJsonArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonDocument()
func DeleteQJsonDocument(this *QJsonDocument) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocumentD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsondocument.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonDocument &)
func (this *QJsonDocument) Swap(other *QJsonDocument) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromRawData(const char *, int, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromRawData(data string, size int, validation int) *QJsonDocument /*123*/ {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", ffiqt.FFI_TYPE_POINTER, convArg0, size, validation)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromRawData(data string, size int, validation int) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromRawData(data, size, validation)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * rawData(int *)
func (this *QJsonDocument) RawData(size unsafe.Pointer /*666*/) string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7rawDataEPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromBinaryData(const QByteArray &, enum QJsonDocument::DataValidation)
func (this *QJsonDocument) FromBinaryData(data *QByteArray, validation int) *QJsonDocument /*123*/ {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", ffiqt.FFI_TYPE_POINTER, convArg0, validation)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromBinaryData(data *QByteArray, validation int) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromBinaryData(data, validation)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toBinaryData()
func (this *QJsonDocument) ToBinaryData() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument12toBinaryDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromVariant(const QVariant &)
func (this *QJsonDocument) FromVariant(variant *QVariant) *QJsonDocument /*123*/ {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument11fromVariantERK8QVariant", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromVariant(variant *QVariant) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromVariant(variant)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:125
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant()
func (this *QJsonDocument) ToVariant() *QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument9toVariantEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toJson()
func (this *QJsonDocument) ToJson() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:138
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray toJson(enum QJsonDocument::JsonFormat)
func (this *QJsonDocument) ToJson_1(format int) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonENS_10JsonFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QJsonDocument) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isArray()
func (this *QJsonDocument) IsArray() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument7isArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObject()
func (this *QJsonDocument) IsObject() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument8isObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:145
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject object()
func (this *QJsonDocument) Object() *QJsonObject /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:146
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray array()
func (this *QJsonDocument) Array() *QJsonArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument5arrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObject(const QJsonObject &)
func (this *QJsonDocument) SetObject(object *QJsonObject) {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument9setObjectERK11QJsonObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArray(const QJsonArray &)
func (this *QJsonDocument) SetArray(array *QJsonArray) {
	var convArg0 = array.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonDocument8setArrayERK10QJsonArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:158
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QJsonDocument) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonDocument6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QJsonDocument__DataValidation = int

const QJsonDocument__Validate QJsonDocument__DataValidation = 0
const QJsonDocument__BypassValidation QJsonDocument__DataValidation = 1

type QJsonDocument__JsonFormat = int

const QJsonDocument__Indented QJsonDocument__JsonFormat = 0
const QJsonDocument__Compact QJsonDocument__JsonFormat = 1

//  body block end
