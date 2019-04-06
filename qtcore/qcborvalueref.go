package qtcore

// /usr/include/qt/QtCore/qcborvalue.h
// #include <qcborvalue.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
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
type QCborValueRef struct {
	*qtrt.CObject
}
type QCborValueRef_ITF interface {
	QCborValueRef_PTR() *QCborValueRef
}

func (ptr *QCborValueRef) QCborValueRef_PTR() *QCborValueRef { return ptr }

func (this *QCborValueRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborValueRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborValueRefFromPointer(cthis unsafe.Pointer) *QCborValueRef {
	return &QCborValueRef{&qtrt.CObject{cthis}}
}
func (*QCborValueRef) NewFromPointer(cthis unsafe.Pointer) *QCborValueRef {
	return NewQCborValueRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborvalue.h:334
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborValueRef & operator=(const QCborValue &)

/*

 */
func (this *QCborValueRef) Operator_equal(other QCborValue_ITF) *QCborValueRef {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRefaSERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:336
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborValueRef & operator=(QCborValue &&)

/*

 */
func (this *QCborValueRef) Operator_equal1(other unsafe.Pointer /*333*/) *QCborValueRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRefaSEO10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:338
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QCborValueRef & operator=(const QCborValueRef &)

/*

 */
func (this *QCborValueRef) Operator_equal2(other QCborValueRef_ITF) *QCborValueRef {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValueRef_PTR() != nil {
		convArg0 = other.QCborValueRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:341
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QCborValue::Type type() const

/*
Returns the type of this QCborValue. The type can also later be retrieved by one of the "isXxx" functions.

See also isInteger(), isByteArray(), isString(), isArray(), isMap(), isTag(), isFalse(), isTrue(), isBool(), isNull(), isUndefined, isDouble(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValueRef) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:342
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInteger() const

/*
Returns true if this QCborValue is of the integer type. The integer value can be retrieved using toInteger().

See also type() and toInteger().
*/
func (this *QCborValueRef) IsInteger() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef9isIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:343
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isByteArray() const

/*
Returns true if this QCborValue is of the byte array type. The byte array value can be retrieved using toByteArray().

See also type() and toByteArray().
*/
func (this *QCborValueRef) IsByteArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11isByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:344
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString() const

/*
Returns true if this QCborValue is of the string type. The string value can be retrieved using toString().

See also type() and toString().
*/
func (this *QCborValueRef) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:345
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray() const

/*
Returns true if this QCborValue is of the array type. The array value can be retrieved using toArray().

See also type() and toArray().
*/
func (this *QCborValueRef) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:346
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMap() const

/*
Returns true if this QCborValue is of the map type. The map value can be retrieved using toMap().

See also type() and toMap().
*/
func (this *QCborValueRef) IsMap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef5isMapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:347
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTag() const

/*
Returns true if this QCborValue is of the tag type. The tag value can be retrieved using tag() and the tagged value using taggedValue().

This function also returns true for extended types that the API recognizes. For code that handles extended types directly before the Qt API is updated to support them, it is possible to recreate the tag + tagged value pair by using taggedValue().

See also type(), tag(), taggedValue(), and taggedValue().
*/
func (this *QCborValueRef) IsTag() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef5isTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:348
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFalse() const

/*
Returns true if this QCborValue is a boolean with false value. This function exists because, internally, CBOR booleans are stored as two separate types, one for true and one for false.

See also type(), isBool(), isTrue(), and toBool().
*/
func (this *QCborValueRef) IsFalse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef7isFalseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:349
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTrue() const

/*
Returns true if this QCborValue is a boolean with true value. This function exists because, internally, CBOR booleans are stored as two separate types, one for false and one for true.

See also type(), isBool(), isFalse(), and toBool().
*/
func (this *QCborValueRef) IsTrue() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6isTrueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:350
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool() const

/*
Returns true if this QCborValue is a boolean. The value can be retrieved using toBool().

See also type(), toBool(), isTrue(), and isFalse().
*/
func (this *QCborValueRef) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:351
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this QCborValue is of the null type.

CBOR null values are used to indicate optional values that were not provided. They are distinct from undefined values, in that null values are usually not the result of an earlier error or problem.

Null values are distinct from undefined values and from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also type(), isUndefined(), and isInvalid().
*/
func (this *QCborValueRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:352
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndefined() const

/*
Returns true if this QCborValue is of the undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are distinct from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also type(), isNull(), and isInvalid().
*/
func (this *QCborValueRef) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:353
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble() const

/*
Returns true if this QCborValue is of the floating-point type. The value can be retrieved using toDouble().

See also type() and toDouble().
*/
func (this *QCborValueRef) IsDouble() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef8isDoubleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:354
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDateTime() const

/*
Returns true if this QCborValue is of the date/time type. The value can be retrieved using toDateTime(). Date/times are extended types that use the tag DateTime.

Additionally, when decoding from a CBOR stream, QCborValue will interpret tags of value UnixTime_t and convert them to the equivalent date/time.

See also type() and toDateTime().
*/
func (this *QCborValueRef) IsDateTime() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef10isDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUrl() const

/*
Returns true if this QCborValue is of the URL type. The URL value can be retrieved using toUrl().

See also type() and toUrl().
*/
func (this *QCborValueRef) IsUrl() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef5isUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:356
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRegularExpression() const

/*
Returns true if this QCborValue contains a regular expression's pattern. The pattern can be retrieved using toRegularExpression().

See also type() and toRegularExpression().
*/
func (this *QCborValueRef) IsRegularExpression() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef19isRegularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:357
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUuid() const

/*
Returns true if this QCborValue contains a UUID. The value can be retrieved using toUuid().

See also type() and toUuid().
*/
func (this *QCborValueRef) IsUuid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6isUuidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:358
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInvalid() const

/*
Returns true if this QCborValue is not of any valid type. Invalid QCborValues are distinct from those with undefined values and they usually represent a decoding error.

See also isUndefined() and isNull().
*/
func (this *QCborValueRef) IsInvalid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef9isInvalidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:359
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isContainer() const

/*
This convenience function returns true if the QCborValue is either an array or a map.

See also isArray() and isMap().
*/
func (this *QCborValueRef) IsContainer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11isContainerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:360
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSimpleType() const

/*
Returns true if this QCborValue is of one of the CBOR simple types. The type itself can later be retrieved using type(), even for types that don't have an enumeration in the API. They can also be checked with the isSimpleType(QCborSimpleType) overload.

See also QCborSimpleType, isSimpleType(QCborSimpleType), and toSimpleType().
*/
func (this *QCborValueRef) IsSimpleType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef12isSimpleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:364
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool isSimpleType(QCborSimpleType) const

/*
Returns true if this QCborValue is of one of the CBOR simple types. The type itself can later be retrieved using type(), even for types that don't have an enumeration in the API. They can also be checked with the isSimpleType(QCborSimpleType) overload.

See also QCborSimpleType, isSimpleType(QCborSimpleType), and toSimpleType().
*/
func (this *QCborValueRef) IsSimpleType1(st int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef12isSimpleTypeE15QCborSimpleType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), st)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:369
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborTag tag(QCborTag) const

/*
Returns the tag of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that number. To retrieve the representation, use taggedValue().

See also isTag(), taggedValue(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValueRef) Tag(defaultValue int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef3tagE8QCborTag", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:369
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborTag tag(QCborTag) const

/*
Returns the tag of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that number. To retrieve the representation, use taggedValue().

See also isTag(), taggedValue(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValueRef) Tagp() int {
	// arg: 0, QCborTag=Enum, QCborTag=Enum, , Invalid
	defaultValue := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef3tagE8QCborTag", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:371
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue taggedValue(const QCborValue &) const

/*
Returns the tagged value of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that representation. To retrieve the tag, use tag().

See also isTag(), tag(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValueRef) TaggedValue(defaultValue QCborValue_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QCborValue_PTR() != nil {
		convArg0 = defaultValue.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11taggedValueERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:371
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue taggedValue(const QCborValue &) const

/*
Returns the tagged value of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that representation. To retrieve the tag, use tag().

See also isTag(), tag(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValueRef) TaggedValuep() *QCborValue /*123*/ {
	// arg: 0, const QCborValue &=LValueReference, QCborValue=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11taggedValueERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:374
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toInteger(qint64) const

/*
Returns the integer value stored in this QCborValue, if it is of the integer type. If it is of the Double type, this function returns the floating point value converted to integer. In any other case, it returns defaultValue.

See also isInteger(), isDouble(), and toDouble().
*/
func (this *QCborValueRef) ToInteger(defaultValue int64) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef9toIntegerEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborvalue.h:374
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toInteger(qint64) const

/*
Returns the integer value stored in this QCborValue, if it is of the integer type. If it is of the Double type, this function returns the floating point value converted to integer. In any other case, it returns defaultValue.

See also isInteger(), isDouble(), and toDouble().
*/
func (this *QCborValueRef) ToIntegerp() int64 {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long, LongLong
	defaultValue := int64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef9toIntegerEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborvalue.h:376
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Returns the boolean value stored in this QCborValue, if it is of a boolean type. Otherwise, it returns defaultValue.

See also isBool(), isTrue(), and isFalse().
*/
func (this *QCborValueRef) ToBool(defaultValue bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:376
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Returns the boolean value stored in this QCborValue, if it is of a boolean type. Otherwise, it returns defaultValue.

See also isBool(), isTrue(), and isFalse().
*/
func (this *QCborValueRef) ToBoolp() bool {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	defaultValue := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:378
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Returns the floating point value stored in this QCborValue, if it is of the Double type. If it is of the Integer type, this function returns the integer value converted to double. In any other case, it returns defaultValue.

See also isDouble(), isInteger(), and toInteger().
*/
func (this *QCborValueRef) ToDouble(defaultValue float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qcborvalue.h:378
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Returns the floating point value stored in this QCborValue, if it is of the Double type. If it is of the Integer type, this function returns the integer value converted to double. In any other case, it returns defaultValue.

See also isDouble(), isInteger(), and toInteger().
*/
func (this *QCborValueRef) ToDoublep() float64 {
	// arg: 0, double=Double, =Invalid, , Invalid
	defaultValue := float64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qcborvalue.h:381
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toByteArray(const QByteArray &) const

/*
Returns the byte array value stored in this QCborValue, if it is of the byte array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QByteArray.

See also isByteArray(), isString(), and toString().
*/
func (this *QCborValueRef) ToByteArray(defaultValue QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QByteArray_PTR() != nil {
		convArg0 = defaultValue.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11toByteArrayERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:381
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toByteArray(const QByteArray &) const

/*
Returns the byte array value stored in this QCborValue, if it is of the byte array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QByteArray.

See also isByteArray(), isString(), and toString().
*/
func (this *QCborValueRef) ToByteArrayp() *QByteArray /*123*/ {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11toByteArrayERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:383
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Returns the string value stored in this QCborValue, if it is of the string type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QString.

See also isString(), isByteArray(), and toByteArray().
*/
func (this *QCborValueRef) ToString(defaultValue string) string {
	var tmpArg0 = NewQString5(defaultValue)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcborvalue.h:383
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Returns the string value stored in this QCborValue, if it is of the string type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QString.

See also isString(), isByteArray(), and toByteArray().
*/
func (this *QCborValueRef) ToStringp() string {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcborvalue.h:385
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QDateTime &) const

/*
Returns the date/time value stored in this QCborValue, if it is of the date/time extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QDateTime.

See also isDateTime(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToDateTime(defaultValue QDateTime_ITF) *QDateTime /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QDateTime_PTR() != nil {
		convArg0 = defaultValue.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef10toDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:385
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QDateTime &) const

/*
Returns the date/time value stored in this QCborValue, if it is of the date/time extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QDateTime.

See also isDateTime(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToDateTimep() *QDateTime /*123*/ {
	// arg: 0, const QDateTime &=LValueReference, QDateTime=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef10toDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:387
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QUrl toUrl(const QUrl &) const

/*
Returns the URL value stored in this QCborValue, if it is of the URL extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUrl.

See also isUrl(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToUrl(defaultValue QUrl_ITF) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QUrl_PTR() != nil {
		convArg0 = defaultValue.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef5toUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:387
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QUrl toUrl(const QUrl &) const

/*
Returns the URL value stored in this QCborValue, if it is of the URL extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUrl.

See also isUrl(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToUrlp() *QUrl /*123*/ {
	// arg: 0, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg0 = NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef5toUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:393
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QUuid toUuid(const QUuid &) const

/*
Returns the UUID value stored in this QCborValue, if it is of the UUID extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUuid.

See also isUuid(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToUuid(defaultValue QUuid_ITF) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QUuid_PTR() != nil {
		convArg0 = defaultValue.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6toUuidERK5QUuid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:393
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QUuid toUuid(const QUuid &) const

/*
Returns the UUID value stored in this QCborValue, if it is of the UUID extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUuid.

See also isUuid(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToUuidp() *QUuid /*123*/ {
	// arg: 0, const QUuid &=LValueReference, QUuid=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef6toUuidERK5QUuid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:401
// index:0
// Public Visibility=Default Availability=Available
// [8] QCborArray toArray() const

/*
Returns the array value stored in this QCborValue, if it is of the array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QCborArray.

See also isArray(), isByteArray(), isMap(), isContainer(), and toMap().
*/
func (this *QCborValueRef) ToArray() *QCborArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef7toArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:402
// index:1
// Public Visibility=Default Availability=Available
// [8] QCborArray toArray(const QCborArray &) const

/*
Returns the array value stored in this QCborValue, if it is of the array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QCborArray.

See also isArray(), isByteArray(), isMap(), isContainer(), and toMap().
*/
func (this *QCborValueRef) ToArray1(a QCborArray_ITF) *QCborArray /*123*/ {
	var convArg0 unsafe.Pointer
	if a != nil && a.QCborArray_PTR() != nil {
		convArg0 = a.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef7toArrayERK10QCborArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:407
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QCborValue &) const

/*
Compares this value and other, and returns an integer that indicates whether this value should be sorted prior to (if the result is negative) or after other (if the result is positive). If this function returns 0, the two values are equal and hold the same contents.

If each QCborValue contains an array or map, the comparison is recursive to elements contained in them.
*/
func (this *QCborValueRef) Compare(other QCborValue_ITF) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef7compareERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcborvalue.h:418
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QCborValue &) const

/*

 */
func (this *QCborValueRef) Operator_equal_equal(other QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRefeqERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:420
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCborValue &) const

/*

 */
func (this *QCborValueRef) Operator_not_equal(other QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRefneERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:422
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QCborValue &) const

/*

 */
func (this *QCborValueRef) Operator_less_than(other QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRefltERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:426
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVariant toVariant() const

/*
Converts this value to a native Qt type and returns the corresponding QVariant.

The following table lists the mapping performed between QCborValue types and Qt meta types.


 CBOR TypeQt or C++ typeNotes
Integerqint64
Doubledouble
Boolbool
Nullstd::nullptr_t
Undefinedno type (QVariant())
Byte arrayQByteArray
StringQString
ArrayQVariantListRecursively converts all values
MapQVariantMapKey types are "stringified"
Other simple typesQCborSimpleType
DateTimeQDateTime
UrlQUrl
RegularExpressionQRegularExpression
UuidQUuid
Other tagsSpecialThe tag is ignored and the tagged value is converted using this function


Note that values in both CBOR Maps and Arrays are converted recursively using this function too and placed in QVariantMap and QVariantList instead. You will not find QCborMap and QCborArray stored inside the QVariants.

QVariantMaps have string keys, unlike CBOR, so the conversion of a QCborMap to QVariantMap will imply a step of "stringification" of the key values. See QCborMap::toJsonObject() for details.

See also fromVariant(), toJsonValue(), QCborArray::toVariantList(), and QCborMap::toVariantMap().
*/
func (this *QCborValueRef) ToVariant() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:427
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue toJsonValue() const

/*
Converts this QCborValue object to an equivalent representation in JSON and returns it as a QJsonValue.

Please note that CBOR contains a richer and wider type set than JSON, so some information may be lost in this conversion. The following table compares CBOR types to JSON types and indicates whether information may be lost or not.


 CBOR TypeJSON TypeComments
BoolBoolNo data loss possible
DoubleNumberInfinities and NaN will be converted to Null; no data loss for other values
IntegerNumberData loss possible in the conversion if the integer is larger than 253 or smaller than -253.
NullNullNo data loss possible
UndefinedNullType information lost
StringStringNo data loss possible
Byte ArrayStringConverted to a lossless encoding like Base64url, but the distinction between strings and byte arrays is lost
Other simple typesStringType information lost
ArrayArrayConversion applies to each contained value
MapObjectKeys are converted to string; values converted according to this table
Tags and extended typesSpecialThe tag number itself is lost and the tagged value is converted to JSON


For information on the conversion of CBOR map keys to string, see QCborMap::toJsonObject().

If this QCborValue contains the undefined value, this function will return an undefined QJsonValue too. Note that JSON does not support undefined values and undefined QJsonValues are an extension to the specification. They cannot be held in a QJsonArray or QJsonObject, but can be returned from functions to indicate a failure. For all other intents and purposes, they are the same as null.
*/
func (this *QCborValueRef) ToJsonValue() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef11toJsonValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:429
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toCbor(QCborValue::EncodingOptions)

/*
Encodes this QCborValue object to its CBOR representation, using the options specified in opt, and return the byte array containing that representation.

This function will not fail, except if this QCborValue or any of the contained items, if this is a map or array, are invalid. Invalid types are not produced normally by the API, but can result from decoding errors.

By default, this function performs no transformation on the values in the QCborValue, writing all floating point directly as double-precision (double) types. If the UseFloat option is specified, it will use single precision (float) for any floating point value for which there's no loss of precision in using that representation. That includes infinities and NaN values.

Similarly, if UseFloat16 is specified, this function will try to use half-precision (qfloat16) floating point if the conversion to that results in no loss of precision. This is always true for infinities and NaN.

If UseIntegers is specified, it will use integers for any floating point value that contains an actual integer.

See also fromCbor(), fromVariant(), and fromJsonValue().
*/
func (this *QCborValueRef) ToCbor(opt int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRef6toCborE6QFlagsIN10QCborValue14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:429
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toCbor(QCborValue::EncodingOptions)

/*
Encodes this QCborValue object to its CBOR representation, using the options specified in opt, and return the byte array containing that representation.

This function will not fail, except if this QCborValue or any of the contained items, if this is a map or array, are invalid. Invalid types are not produced normally by the API, but can result from decoding errors.

By default, this function performs no transformation on the values in the QCborValue, writing all floating point directly as double-precision (double) types. If the UseFloat option is specified, it will use single precision (float) for any floating point value for which there's no loss of precision in using that representation. That includes infinities and NaN values.

Similarly, if UseFloat16 is specified, this function will try to use half-precision (qfloat16) floating point if the conversion to that results in no loss of precision. This is always true for infinities and NaN.

If UseIntegers is specified, it will use integers for any floating point value that contains an actual integer.

See also fromCbor(), fromVariant(), and fromJsonValue().
*/
func (this *QCborValueRef) ToCborp() *QByteArray /*123*/ {
	// arg: 0, QCborValue::EncodingOptions=Elaborated, QCborValue::EncodingOptions=Typedef, QFlags<QCborValue::EncodingOption>, Unexposed
	opt := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRef6toCborE6QFlagsIN10QCborValue14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:431
// index:1
// Public Visibility=Default Availability=Available
// [-2] void toCbor(QCborStreamWriter &, QCborValue::EncodingOptions)

/*
Encodes this QCborValue object to its CBOR representation, using the options specified in opt, and return the byte array containing that representation.

This function will not fail, except if this QCborValue or any of the contained items, if this is a map or array, are invalid. Invalid types are not produced normally by the API, but can result from decoding errors.

By default, this function performs no transformation on the values in the QCborValue, writing all floating point directly as double-precision (double) types. If the UseFloat option is specified, it will use single precision (float) for any floating point value for which there's no loss of precision in using that representation. That includes infinities and NaN values.

Similarly, if UseFloat16 is specified, this function will try to use half-precision (qfloat16) floating point if the conversion to that results in no loss of precision. This is always true for infinities and NaN.

If UseIntegers is specified, it will use integers for any floating point value that contains an actual integer.

See also fromCbor(), fromVariant(), and fromJsonValue().
*/
func (this *QCborValueRef) ToCbor1(writer QCborStreamWriter_ITF, opt int) {
	var convArg0 unsafe.Pointer
	if writer != nil && writer.QCborStreamWriter_PTR() != nil {
		convArg0 = writer.QCborStreamWriter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRef6toCborER17QCborStreamWriter6QFlagsIN10QCborValue14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, opt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:431
// index:1
// Public Visibility=Default Availability=Available
// [-2] void toCbor(QCborStreamWriter &, QCborValue::EncodingOptions)

/*
Encodes this QCborValue object to its CBOR representation, using the options specified in opt, and return the byte array containing that representation.

This function will not fail, except if this QCborValue or any of the contained items, if this is a map or array, are invalid. Invalid types are not produced normally by the API, but can result from decoding errors.

By default, this function performs no transformation on the values in the QCborValue, writing all floating point directly as double-precision (double) types. If the UseFloat option is specified, it will use single precision (float) for any floating point value for which there's no loss of precision in using that representation. That includes infinities and NaN values.

Similarly, if UseFloat16 is specified, this function will try to use half-precision (qfloat16) floating point if the conversion to that results in no loss of precision. This is always true for infinities and NaN.

If UseIntegers is specified, it will use integers for any floating point value that contains an actual integer.

See also fromCbor(), fromVariant(), and fromJsonValue().
*/
func (this *QCborValueRef) ToCbor1p(writer QCborStreamWriter_ITF) {
	var convArg0 unsafe.Pointer
	if writer != nil && writer.QCborStreamWriter_PTR() != nil {
		convArg0 = writer.QCborStreamWriter_PTR().GetCthis()
	}
	// arg: 1, QCborValue::EncodingOptions=Elaborated, QCborValue::EncodingOptions=Typedef, QFlags<QCborValue::EncodingOption>, Unexposed
	opt := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRef6toCborER17QCborStreamWriter6QFlagsIN10QCborValue14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, opt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:433
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toDiagnosticNotation(QCborValue::DiagnosticNotationOptions)

/*
Creates the diagnostic notation equivalent of this CBOR object and returns it. The opts parameter controls the dialect of the notation. Diagnostic notation is useful in debugging, to aid the developer in understanding what value is stored in the QCborValue or in a CBOR stream. For that reason, the Qt API provides no support for parsing the diagnostic back into the in-memory format or CBOR stream, though the representation is unique and it would be possible.

CBOR diagnostic notation is specified by section 6 of RFC 7049. It is a text representation of the CBOR stream and it is very similar to JSON, but it supports the CBOR types not found in JSON. The extended format enabled by the ExtendedFormat flag is currently in some IETF drafts and its format is subject to change.

This function produces the equivalent representation of the stream that toCbor() would produce, without any transformation option provided there. This also implies this function may not produce a representation of the stream that was used to create the object, if it was created using fromCbor(), as that function may have applied transformations. For a high-fidelity notation of a stream, without transformation, see the cbordump example.

See also toCbor() and QJsonDocument::toJson().
*/
func (this *QCborValueRef) ToDiagnosticNotation(opt int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRef20toDiagnosticNotationE6QFlagsIN10QCborValue24DiagnosticNotationOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcborvalue.h:433
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toDiagnosticNotation(QCborValue::DiagnosticNotationOptions)

/*
Creates the diagnostic notation equivalent of this CBOR object and returns it. The opts parameter controls the dialect of the notation. Diagnostic notation is useful in debugging, to aid the developer in understanding what value is stored in the QCborValue or in a CBOR stream. For that reason, the Qt API provides no support for parsing the diagnostic back into the in-memory format or CBOR stream, though the representation is unique and it would be possible.

CBOR diagnostic notation is specified by section 6 of RFC 7049. It is a text representation of the CBOR stream and it is very similar to JSON, but it supports the CBOR types not found in JSON. The extended format enabled by the ExtendedFormat flag is currently in some IETF drafts and its format is subject to change.

This function produces the equivalent representation of the stream that toCbor() would produce, without any transformation option provided there. This also implies this function may not produce a representation of the stream that was used to create the object, if it was created using fromCbor(), as that function may have applied transformations. For a high-fidelity notation of a stream, without transformation, see the cbordump example.

See also toCbor() and QJsonDocument::toJson().
*/
func (this *QCborValueRef) ToDiagnosticNotationp() string {
	// arg: 0, QCborValue::DiagnosticNotationOptions=Elaborated, QCborValue::DiagnosticNotationOptions=Typedef, QFlags<QCborValue::DiagnosticNotationOption>, Unexposed
	opt := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRef20toDiagnosticNotationE6QFlagsIN10QCborValue24DiagnosticNotationOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

func DeleteQCborValueRef(this *QCborValueRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QCborValueRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10353() {
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
