package qtcore

// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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

type QJsonValueRef struct {
	*qtrt.CObject
}

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
func NewQJsonValueRef(array *QJsonArray /*777 QJsonArray **/, idx int) *QJsonValueRef {
	var convArg0 = array.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValueRefC2EP10QJsonArrayi", ffiqt.FFI_TYPE_POINTER, convArg0, idx)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRef)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:176
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRef(QJsonObject *, int)
func NewQJsonValueRef_1(object *QJsonObject /*777 QJsonObject **/, idx int) *QJsonValueRef {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValueRefC2EP11QJsonObjecti", ffiqt.FFI_TYPE_POINTER, convArg0, idx)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRef)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:183
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant()
func (this *QJsonValueRef) ToVariant() *QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef9toVariantEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QJsonValue::Type type()
func (this *QJsonValueRef) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:185
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QJsonValueRef) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool()
func (this *QJsonValueRef) IsBool() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6isBoolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble()
func (this *QJsonValueRef) IsDouble() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8isDoubleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString()
func (this *QJsonValueRef) IsString() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8isStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray()
func (this *QJsonValueRef) IsArray() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef7isArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:190
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isObject()
func (this *QJsonValueRef) IsObject() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8isObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:191
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndefined()
func (this *QJsonValueRef) IsUndefined() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef11isUndefinedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool()
func (this *QJsonValueRef) ToBool() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6toBoolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:201
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool toBool(_Bool)
func (this *QJsonValueRef) ToBool_1(defaultValue bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6toBoolEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonvalue.h:194
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int toInt()
func (this *QJsonValueRef) ToInt() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef5toIntEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int toInt(int)
func (this *QJsonValueRef) ToInt_1(defaultValue int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef5toIntEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble()
func (this *QJsonValueRef) ToDouble() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toDoubleEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:203
// index:1
// Public inline Visibility=Default Availability=Available
// [8] double toDouble(double)
func (this *QJsonValueRef) ToDouble_1(defaultValue float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toDoubleEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qjsonvalue.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toString()
func (this *QJsonValueRef) ToString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString toString(const QString &)
func (this *QJsonValueRef) ToString_1(defaultValue *QString) *QString /*123*/ {
	var convArg0 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toStringERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:197
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray toArray()
func (this *QJsonValueRef) ToArray() *QJsonArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef7toArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:198
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toObject()
func (this *QJsonValueRef) ToObject() *QJsonObject /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

func DeleteQJsonValueRef(this *QJsonValueRef) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValueRefD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
