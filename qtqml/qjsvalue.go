package qtqml

// /usr/include/qt/QtQml/qjsvalue.h
// #include <qjsvalue.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QJSValue struct {
	*qtrt.CObject
}
type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue { return ptr }

func (this *QJSValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJSValue) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJSValueFromPointer(cthis unsafe.Pointer) *QJSValue {
	return &QJSValue{&qtrt.CObject{cthis}}
}
func (*QJSValue) NewFromPointer(cthis unsafe.Pointer) *QJSValue {
	return NewQJSValueFromPointer(cthis)
}

// /usr/include/qt/QtQml/qjsvalue.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(enum QJSValue::SpecialValue)
func NewQJSValue(value int) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ENS_12SpecialValueE", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(enum QJSValue::SpecialValue)
func NewQJSValue__() *QJSValue {
	// arg: 0, QJSValue::SpecialValue=Enum, QJSValue::SpecialValue=Enum,
	value := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ENS_12SpecialValueE", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:82
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(_Bool)
func NewQJSValue_1(value bool) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Eb", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:83
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(int)
func NewQJSValue_2(value int) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Ei", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:84
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(uint)
func NewQJSValue_3(value uint) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Ej", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:85
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(double)
func NewQJSValue_4(value float64) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2Ed", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:86
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(const QString &)
func NewQJSValue_5(value string) *QJSValue {
	var tmpArg0 = qtcore.NewQString_5(value)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:87
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(const QLatin1String &)
func NewQJSValue_6(value qtcore.QLatin1String_ITF) *QJSValue {
	var convArg0 unsafe.Pointer
	if value != nil && value.QLatin1String_PTR() != nil {
		convArg0 = value.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2ERK13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:89
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QJSValue(const char *)
func NewQJSValue_7(str string) *QJSValue {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValue)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJSValue()
func DeleteQJSValue(this *QJSValue) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qjsvalue.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QJSValue & operator=(QJSValue &&)
func (this *QJSValue) Operator_equal(other unsafe.Pointer /*333*/) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:92
// index:1
// Public Visibility=Default Availability=Available
// [8] QJSValue & operator=(const QJSValue &)
func (this *QJSValue) Operator_equal_1(other QJSValue_ITF) *QJSValue {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJSValue_PTR() != nil {
		convArg0 = other.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValueaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBool() const
func (this *QJSValue) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNumber() const
func (this *QJSValue) IsNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QJSValue) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isString() const
func (this *QJSValue) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndefined() const
func (this *QJSValue) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVariant() const
func (this *QJSValue) IsVariant() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9isVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isQObject() const
func (this *QJSValue) IsQObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9isQObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isQMetaObject() const
func (this *QJSValue) IsQMetaObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue13isQMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObject() const
func (this *QJSValue) IsObject() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDate() const
func (this *QJSValue) IsDate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6isDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRegExp() const
func (this *QJSValue) IsRegExp() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8isRegExpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isArray() const
func (this *QJSValue) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isError() const
func (this *QJSValue) IsError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue7isErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const
func (this *QJSValue) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qjsvalue.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] double toNumber() const
func (this *QJSValue) ToNumber() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8toNumberEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qjsvalue.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] qint32 toInt() const
func (this *QJSValue) ToInt() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue5toIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv) // 222
}

// /usr/include/qt/QtQml/qjsvalue.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 toUInt() const
func (this *QJSValue) ToUInt() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6toUIntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtQml/qjsvalue.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool() const
func (this *QJSValue) ToBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6toBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:113
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant toVariant() const
func (this *QJSValue) ToVariant() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * toQObject() const
func (this *QJSValue) ToQObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9toQObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qjsvalue.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMetaObject * toQMetaObject() const
func (this *QJSValue) ToQMetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue13toQMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qjsvalue.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime() const
func (this *QJSValue) ToDateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue10toDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool equals(const QJSValue &) const
func (this *QJSValue) Equals(other QJSValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJSValue_PTR() != nil {
		convArg0 = other.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6equalsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strictlyEquals(const QJSValue &) const
func (this *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJSValue_PTR() != nil {
		convArg0 = other.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue14strictlyEqualsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue prototype() const
func (this *QJSValue) Prototype() *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue9prototypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrototype(const QJSValue &)
func (this *QJSValue) SetPrototype(prototype QJSValue_ITF) {
	var convArg0 unsafe.Pointer
	if prototype != nil && prototype.QJSValue_PTR() != nil {
		convArg0 = prototype.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue12setPrototypeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue property(const QString &) const
func (this *QJSValue) Property(name string) *QJSValue /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8propertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:130
// index:1
// Public Visibility=Default Availability=Available
// [8] QJSValue property(quint32) const
func (this *QJSValue) Property_1(arrayIndex uint) *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue8propertyEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arrayIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProperty(const QString &, const QJSValue &)
func (this *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QJSValue_PTR() != nil {
		convArg1 = value.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue11setPropertyERK7QStringRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:131
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setProperty(quint32, const QJSValue &)
func (this *QJSValue) SetProperty_1(arrayIndex uint, value QJSValue_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QJSValue_PTR() != nil {
		convArg1 = value.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue11setPropertyEjRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arrayIndex, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasProperty(const QString &) const
func (this *QJSValue) HasProperty(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue11hasPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasOwnProperty(const QString &) const
func (this *QJSValue) HasOwnProperty(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue14hasOwnPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool deleteProperty(const QString &)
func (this *QJSValue) DeleteProperty(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QJSValue14deletePropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCallable() const
func (this *QJSValue) IsCallable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue10isCallableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSEngine * engine() const
func (this *QJSValue) Engine() *QJSEngine /*777 QJSEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QJSValue6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

type QJSValue__SpecialValue = int

const QJSValue__NullValue QJSValue__SpecialValue = 0
const QJSValue__UndefinedValue QJSValue__SpecialValue = 1

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
