package qtcore

// /usr/include/qt/QtCore/qjsondocument.h
// #include <qjsondocument.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QJsonDocument struct {
	*qtrt.CObject
}
type QJsonDocument_ITF interface {
	QJsonDocument_PTR() *QJsonDocument
}

func (ptr *QJsonDocument) QJsonDocument_PTR() *QJsonDocument { return ptr }

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

/*
Constructs an empty and invalid document.
*/
func (*QJsonDocument) NewForInherit() *QJsonDocument {
	return NewQJsonDocument()
}
func NewQJsonDocument() *QJsonDocument {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument(const QJsonObject &)

/*
Constructs an empty and invalid document.
*/
func (*QJsonDocument) NewForInherit1(object QJsonObject_ITF) *QJsonDocument {
	return NewQJsonDocument1(object)
}
func NewQJsonDocument1(object QJsonObject_ITF) *QJsonDocument {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJsonObject_PTR() != nil {
		convArg0 = object.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK11QJsonObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:90
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QJsonDocument(const QJsonArray &)

/*
Constructs an empty and invalid document.
*/
func (*QJsonDocument) NewForInherit2(array QJsonArray_ITF) *QJsonDocument {
	return NewQJsonDocument2(array)
}
func NewQJsonDocument2(array QJsonArray_ITF) *QJsonDocument {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentC2ERK10QJsonArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonDocument)
	return gothis
}

// /usr/include/qt/QtCore/qjsondocument.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonDocument()

/*

 */
func DeleteQJsonDocument(this *QJsonDocument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsondocument.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QJsonDocument & operator=(const QJsonDocument &)

/*

 */
func (this *QJsonDocument) Operator_equal(other QJsonDocument_ITF) *QJsonDocument {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:102
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QJsonDocument & operator=(QJsonDocument &&)

/*

 */
func (this *QJsonDocument) Operator_equal1(other unsafe.Pointer /*333*/) *QJsonDocument {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocumentaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonDocument &)

/*
Swaps the document other with this. This operation is very fast and never fails.

This function was introduced in  Qt 5.10.
*/
func (this *QJsonDocument) Swap(other QJsonDocument_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromRawData(const char *, int, QJsonDocument::DataValidation)

/*
Creates a QJsonDocument that uses the first size bytes from data. It assumes data contains a binary encoded JSON document. The created document does not take ownership of data and the caller has to guarantee that data will not be deleted or modified as long as any QJsonDocument, QJsonObject or QJsonArray still references the data.

data has to be aligned to a 4 byte boundary.

validation decides whether the data is checked for validity before being used. By default the data is validated. If the data is not valid, the method returns a null document.

Returns a QJsonDocument representing the data.

See also rawData(), fromBinaryData(), isNull(), and DataValidation.
*/
func (this *QJsonDocument) FromRawData(data string, size int, validation int) *QJsonDocument /*123*/ {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, size, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromRawData(data string, size int, validation int) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromRawData(data, size, validation)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromRawData(const char *, int, QJsonDocument::DataValidation)

/*
Creates a QJsonDocument that uses the first size bytes from data. It assumes data contains a binary encoded JSON document. The created document does not take ownership of data and the caller has to guarantee that data will not be deleted or modified as long as any QJsonDocument, QJsonObject or QJsonArray still references the data.

data has to be aligned to a 4 byte boundary.

validation decides whether the data is checked for validity before being used. By default the data is validated. If the data is not valid, the method returns a null document.

Returns a QJsonDocument representing the data.

See also rawData(), fromBinaryData(), isNull(), and DataValidation.
*/
func (this *QJsonDocument) FromRawDatap(data string, size int) *QJsonDocument /*123*/ {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	// arg: 2, QJsonDocument::DataValidation=Enum, QJsonDocument::DataValidation=Enum, , Invalid
	validation := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument11fromRawDataEPKciNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, size, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * rawData(int *) const

/*
Returns the raw binary representation of the data size will contain the size of the returned data.

This method is useful to e.g. stream the JSON document in it's binary form to a file.
*/
func (this *QJsonDocument) RawData(size unsafe.Pointer /*666*/) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument7rawDataEPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromBinaryData(const QByteArray &, QJsonDocument::DataValidation)

/*
Creates a QJsonDocument from data.

validation decides whether the data is checked for validity before being used. By default the data is validated. If the data is not valid, the method returns a null document.

See also toBinaryData(), fromRawData(), isNull(), and DataValidation.
*/
func (this *QJsonDocument) FromBinaryData(data QByteArray_ITF, validation int) *QJsonDocument /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromBinaryData(data QByteArray_ITF, validation int) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromBinaryData(data, validation)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromBinaryData(const QByteArray &, QJsonDocument::DataValidation)

/*
Creates a QJsonDocument from data.

validation decides whether the data is checked for validity before being used. By default the data is validated. If the data is not valid, the method returns a null document.

See also toBinaryData(), fromRawData(), isNull(), and DataValidation.
*/
func (this *QJsonDocument) FromBinaryDatap(data QByteArray_ITF) *QJsonDocument /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QJsonDocument::DataValidation=Enum, QJsonDocument::DataValidation=Enum, , Invalid
	validation := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument14fromBinaryDataERK10QByteArrayNS_14DataValidationE", qtrt.FFI_TYPE_POINTER, convArg0, validation)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toBinaryData() const

/*
Returns a binary representation of the document.

The binary representation is also the native format used internally in Qt, and is very efficient and fast to convert to and from.

The binary format can be stored on disk and interchanged with other applications or computers. fromBinaryData() can be used to convert it back into a JSON document.

See also fromBinaryData().
*/
func (this *QJsonDocument) ToBinaryData() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument12toBinaryDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [8] QJsonDocument fromVariant(const QVariant &)

/*
Creates a QJsonDocument from the QVariant variant.

If the variant contains any other type than a QVariantMap, QVariantHash, QVariantList or QStringList, the returned document is invalid.

See also toVariant().
*/
func (this *QJsonDocument) FromVariant(variant QVariant_ITF) *QJsonDocument /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument11fromVariantERK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}
func QJsonDocument_FromVariant(variant QVariant_ITF) *QJsonDocument /*123*/ {
	var nilthis *QJsonDocument
	rv := nilthis.FromVariant(variant)
	return rv
}

// /usr/include/qt/QtCore/qjsondocument.h:125
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant() const

/*
Returns a QVariant representing the Json document.

The returned variant will be a QVariantList if the document is a QJsonArray and a QVariantMap if the document is a QJsonObject.

See also fromVariant() and QJsonValue::toVariant().
*/
func (this *QJsonDocument) ToVariant() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toJson() const

/*
Converts the QJsonDocument to a UTF-8 encoded JSON document.

See also fromJson().
*/
func (this *QJsonDocument) ToJson() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray toJson(QJsonDocument::JsonFormat) const

/*
Converts the QJsonDocument to a UTF-8 encoded JSON document.

See also fromJson().
*/
func (this *QJsonDocument) ToJson1(format int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6toJsonENS_10JsonFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the document doesn't contain any data.
*/
func (this *QJsonDocument) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isArray() const

/*
Returns true if the document contains an array.

See also array() and isObject().
*/
func (this *QJsonDocument) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObject() const

/*
Returns true if the document contains an object.

See also object() and isArray().
*/
func (this *QJsonDocument) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:143
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject object() const

/*
Returns the QJsonObject contained in the document.

Returns an empty object if the document contains an array.

See also isObject(), array(), and setObject().
*/
func (this *QJsonDocument) Object() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:144
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray array() const

/*
Returns the QJsonArray contained in the document.

Returns an empty array if the document contains an object.

See also isArray(), object(), and setArray().
*/
func (this *QJsonDocument) Array() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument5arrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObject(const QJsonObject &)

/*
Sets object as the main object of this document.

See also setArray() and object().
*/
func (this *QJsonDocument) SetObject(object QJsonObject_ITF) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJsonObject_PTR() != nil {
		convArg0 = object.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument9setObjectERK11QJsonObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArray(const QJsonArray &)

/*
Sets array as the main object of this document.

See also setObject() and array().
*/
func (this *QJsonDocument) SetArray(array QJsonArray_ITF) {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonDocument8setArrayERK10QJsonArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsondocument.h:149
// index:0
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](const QString &) const

/*

 */
func (this *QJsonDocument) Operator_get_index(key string) *QJsonValue /*123*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:150
// index:1
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](QLatin1String) const

/*

 */
func (this *QJsonDocument) Operator_get_index1(key QLatin1String_ITF /*123*/) *QJsonValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:151
// index:2
// Public Visibility=Default Availability=Available
// [24] const QJsonValue operator[](int) const

/*

 */
func (this *QJsonDocument) Operator_get_index2(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsondocument.h:153
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QJsonDocument &) const

/*

 */
func (this *QJsonDocument) Operator_equal_equal(other QJsonDocument_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumenteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonDocument &) const

/*

 */
func (this *QJsonDocument) Operator_not_equal(other QJsonDocument_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonDocument_PTR() != nil {
		convArg0 = other.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocumentneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsondocument.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
returns true if this document is null.

Null documents are documents created through the default constructor.

Documents created from UTF-8 encoded text or the binary format are validated during parsing. If validation fails, the returned document will also be null.
*/
func (this *QJsonDocument) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonDocument6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This value is used to tell QJsonDocument whether to validate the binary data when converting to a QJsonDocument using fromBinaryData() or fromRawData().


*/
type QJsonDocument__DataValidation = int

// Validate the data before using it. This is the default.
const QJsonDocument__Validate QJsonDocument__DataValidation = 0

// Bypasses data validation. Only use if you received the data from a trusted place and know it's valid, as using of invalid data can crash the application.
const QJsonDocument__BypassValidation QJsonDocument__DataValidation = 1

func (this *QJsonDocument) DataValidationItemName(val int) string {
	switch val {
	case QJsonDocument__Validate: // 0
		return "Validate"
	case QJsonDocument__BypassValidation: // 1
		return "BypassValidation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QJsonDocument_DataValidationItemName(val int) string {
	var nilthis *QJsonDocument
	return nilthis.DataValidationItemName(val)
}

/*
This value defines the format of the JSON byte array produced when converting to a QJsonDocument using toJson().

      {
          "Array": [
              true,
              999,
              "string"
          ],
          "Key": "Value",
          "null": null
      }



      {"Array":[true,999,"string"],"Key":"Value","null":null}





This enum was introduced or modified in  Qt 5.1.

*/
type QJsonDocument__JsonFormat = int

// Defines human readable output as follows:
const QJsonDocument__Indented QJsonDocument__JsonFormat = 0

// Defines a compact output as follows:
const QJsonDocument__Compact QJsonDocument__JsonFormat = 1

func (this *QJsonDocument) JsonFormatItemName(val int) string {
	switch val {
	case QJsonDocument__Indented: // 0
		return "Indented"
	case QJsonDocument__Compact: // 1
		return "Compact"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QJsonDocument_JsonFormatItemName(val int) string {
	var nilthis *QJsonDocument
	return nilthis.JsonFormatItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10447() {
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
