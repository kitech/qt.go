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
// extern C begin: 38
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
type QJsonValueRef struct {
	*qtrt.CObject
}
type QJsonValueRef_ITF interface {
	QJsonValueRef_PTR() *QJsonValueRef
}

func (ptr *QJsonValueRef) QJsonValueRef_PTR() *QJsonValueRef { return ptr }

func (this *QJsonValueRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonValueRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonValueRefFromPointer(cthis unsafe.Pointer) *QJsonValueRef {
	return &QJsonValueRef{&qtrt.CObject{cthis}}
}
func (*QJsonValueRef) NewFromPointer(cthis unsafe.Pointer) *QJsonValueRef {
	return NewQJsonValueRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonvalue.h:174
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRef(QJsonArray *, int)

/*

 */
func (*QJsonValueRef) NewForInherit(array QJsonArray_ITF /*777 QJsonArray **/, idx int) *QJsonValueRef {
	return NewQJsonValueRef(array, idx)
}
func NewQJsonValueRef(array QJsonArray_ITF /*777 QJsonArray **/, idx int) *QJsonValueRef {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValueRefC2EP10QJsonArrayi", qtrt.FFI_TYPE_POINTER, convArg0, idx)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRef)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:176
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRef(QJsonObject *, int)

/*

 */
func (*QJsonValueRef) NewForInherit1(object QJsonObject_ITF /*777 QJsonObject **/, idx int) *QJsonValueRef {
	return NewQJsonValueRef1(object, idx)
}
func NewQJsonValueRef1(object QJsonObject_ITF /*777 QJsonObject **/, idx int) *QJsonValueRef {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJsonObject_PTR() != nil {
		convArg0 = object.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValueRefC2EP11QJsonObjecti", qtrt.FFI_TYPE_POINTER, convArg0, idx)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRef)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:180
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonValueRef & operator=(const QJsonValue &)

/*

 */
func (this *QJsonValueRef) Operator_equal(val QJsonValue_ITF) *QJsonValueRef {
	var convArg0 unsafe.Pointer
	if val != nil && val.QJsonValue_PTR() != nil {
		convArg0 = val.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValueRefaSERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:181
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonValueRef & operator=(const QJsonValueRef &)

/*

 */
func (this *QJsonValueRef) Operator_equal1(val QJsonValueRef_ITF) *QJsonValueRef {
	var convArg0 unsafe.Pointer
	if val != nil && val.QJsonValueRef_PTR() != nil {
		convArg0 = val.QJsonValueRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValueRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:183
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
func (this *QJsonValueRef) ToVariant() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QJsonValue::Type type() const

/*
Returns the type of the value.

See also QJsonValue::Type.
*/
func (this *QJsonValueRef) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:185
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the value is null.
*/
func (this *QJsonValueRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool() const

/*
Returns true if the value contains a boolean.

See also toBool().
*/
func (this *QJsonValueRef) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble() const

/*
Returns true if the value contains a double.

See also toDouble().
*/
func (this *QJsonValueRef) IsDouble() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8isDoubleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString() const

/*
Returns true if the value contains a string.

See also toString().
*/
func (this *QJsonValueRef) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray() const

/*
Returns true if the value contains an array.

See also toArray().
*/
func (this *QJsonValueRef) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:190
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isObject() const

/*
Returns true if the value contains an object.

See also toObject().
*/
func (this *QJsonValueRef) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:191
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndefined() const

/*
Returns true if the value is undefined. This can happen in certain error cases as e.g. accessing a non existing key in a QJsonObject.
*/
func (this *QJsonValueRef) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool() const

/*
Converts the value to a bool and returns it.

If type() is not bool, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef6toBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:201
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Converts the value to a bool and returns it.

If type() is not bool, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToBool1(defaultValue bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:194
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int toInt() const

/*
Converts the value to an int and returns it.

If type() is not Double or the value is not a whole number, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToInt() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef5toIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int toInt(int) const

/*
Converts the value to an int and returns it.

If type() is not Double or the value is not a whole number, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToInt1(defaultValue int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef5toIntEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble() const

/*
Converts the value to a double and returns it.

If type() is not Double, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToDouble() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8toDoubleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:203
// index:1
// Public inline Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Converts the value to a double and returns it.

If type() is not Double, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToDouble1(defaultValue float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toString() const

/*
Converts the value to a QString and returns it.

If type() is not String, a null QString will be returned.

See also QString::isNull().
*/
func (this *QJsonValueRef) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qjsonvalue.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Converts the value to a QString and returns it.

If type() is not String, a null QString will be returned.

See also QString::isNull().
*/
func (this *QJsonValueRef) ToString1(defaultValue string) string {
	var tmpArg0 = NewQString5(defaultValue)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qjsonvalue.h:197
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray toArray() const

/*
Converts the value to an array and returns it.

If type() is not Array, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToArray() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef7toArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:198
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toObject() const

/*
Converts the value to an object and returns it.

If type() is not Object, the defaultValue will be returned.
*/
func (this *QJsonValueRef) ToObject() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRef8toObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QJsonValue &) const

/*

 */
func (this *QJsonValueRef) Operator_equal_equal(other QJsonValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonValue_PTR() != nil {
		convArg0 = other.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRefeqERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonValue &) const

/*

 */
func (this *QJsonValueRef) Operator_not_equal(other QJsonValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonValue_PTR() != nil {
		convArg0 = other.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QJsonValueRefneERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQJsonValueRef(this *QJsonValueRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValueRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
