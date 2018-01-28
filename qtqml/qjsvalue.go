package qtqml

// /usr/include/qt/QtQml/qjsvalue.h
// #include <qjsvalue.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QJSValue struct {
	*qtrt.CObject
}

func (this *QJSValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJSValue) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQJSValueFromPointer(cthis unsafe.Pointer) *QJSValue {
	return &QJSValue{&qtrt.CObject{cthis}}
}
func (*QJSValue) NewFromPointer(cthis unsafe.Pointer) *QJSValue {
	return NewQJSValueFromPointer(cthis)
}

// /usr/include/qt/QtQml/qjsvalue.h:72
// index:0
// Public
// void QJSValue(enum QJSValue::SpecialValue)
func NewQJSValue(value int) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2ENS_12SpecialValueE", ffiqt.FFI_TYPE_VOID, cthis, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:82
// index:1
// Public
// void QJSValue(_Bool)
func NewQJSValue_1(value bool) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2Eb", ffiqt.FFI_TYPE_VOID, cthis, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:83
// index:2
// Public
// void QJSValue(int)
func NewQJSValue_2(value int) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2Ei", ffiqt.FFI_TYPE_VOID, cthis, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:84
// index:3
// Public
// void QJSValue(uint)
func NewQJSValue_3(value uint) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2Ej", ffiqt.FFI_TYPE_VOID, cthis, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:85
// index:4
// Public
// void QJSValue(double)
func NewQJSValue_4(value float64) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2Ed", ffiqt.FFI_TYPE_VOID, cthis, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:86
// index:5
// Public
// void QJSValue(const QString &)
func NewQJSValue_5(value *qtcore.QString) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:87
// index:6
// Public
// void QJSValue(const QLatin1String &)
func NewQJSValue_6(value *qtcore.QLatin1String) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2ERK13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:89
// index:7
// Public
// void QJSValue(const char *)
func NewQJSValue_7(str string) *QJSValue {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalue.h:73
// index:0
// Public
// void ~QJSValue()
func DeleteQJSValue(*QJSValue) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValueD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:94
// index:0
// Public
// bool isBool()
func (this *QJSValue) IsBool() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6isBoolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:95
// index:0
// Public
// bool isNumber()
func (this *QJSValue) IsNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8isNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:96
// index:0
// Public
// bool isNull()
func (this *QJSValue) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:97
// index:0
// Public
// bool isString()
func (this *QJSValue) IsString() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8isStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:98
// index:0
// Public
// bool isUndefined()
func (this *QJSValue) IsUndefined() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue11isUndefinedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:99
// index:0
// Public
// bool isVariant()
func (this *QJSValue) IsVariant() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue9isVariantEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:100
// index:0
// Public
// bool isQObject()
func (this *QJSValue) IsQObject() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue9isQObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:101
// index:0
// Public
// bool isQMetaObject()
func (this *QJSValue) IsQMetaObject() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue13isQMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:102
// index:0
// Public
// bool isObject()
func (this *QJSValue) IsObject() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8isObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:103
// index:0
// Public
// bool isDate()
func (this *QJSValue) IsDate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6isDateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:104
// index:0
// Public
// bool isRegExp()
func (this *QJSValue) IsRegExp() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8isRegExpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:105
// index:0
// Public
// bool isArray()
func (this *QJSValue) IsArray() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue7isArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:106
// index:0
// Public
// bool isError()
func (this *QJSValue) IsError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue7isErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:108
// index:0
// Public
// QString toString()
func (this *QJSValue) ToString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8toStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:109
// index:0
// Public
// double toNumber()
func (this *QJSValue) ToNumber() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8toNumberEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qjsvalue.h:110
// index:0
// Public
// qint32 toInt()
func (this *QJSValue) ToInt() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue5toIntEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 222
}

// /usr/include/qt/QtQml/qjsvalue.h:111
// index:0
// Public
// quint32 toUInt()
func (this *QJSValue) ToUInt() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6toUIntEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtQml/qjsvalue.h:112
// index:0
// Public
// bool toBool()
func (this *QJSValue) ToBool() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6toBoolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:113
// index:0
// Public
// QVariant toVariant()
func (this *QJSValue) ToVariant() *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue9toVariantEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:114
// index:0
// Public
// QObject * toQObject()
func (this *QJSValue) ToQObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue9toQObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:115
// index:0
// Public
// const QMetaObject * toQMetaObject()
func (this *QJSValue) ToQMetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue13toQMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:116
// index:0
// Public
// QDateTime toDateTime()
func (this *QJSValue) ToDateTime() *qtcore.QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue10toDateTimeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:118
// index:0
// Public
// bool equals(const QJSValue &)
func (this *QJSValue) Equals(other *QJSValue) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6equalsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:119
// index:0
// Public
// bool strictlyEquals(const QJSValue &)
func (this *QJSValue) StrictlyEquals(other *QJSValue) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue14strictlyEqualsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:121
// index:0
// Public
// QJSValue prototype()
func (this *QJSValue) Prototype() *QJSValue /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue9prototypeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:122
// index:0
// Public
// void setPrototype(const QJSValue &)
func (this *QJSValue) SetPrototype(prototype *QJSValue) {
	var convArg0 = prototype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValue12setPrototypeERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:124
// index:0
// Public
// QJSValue property(const QString &)
func (this *QJSValue) Property(name *qtcore.QString) *QJSValue /*123*/ {
	var convArg0 = name.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8propertyERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:130
// index:1
// Public
// QJSValue property(quint32)
func (this *QJSValue) Property_1(arrayIndex uint) *QJSValue /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue8propertyEj", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arrayIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalue.h:125
// index:0
// Public
// void setProperty(const QString &, const QJSValue &)
func (this *QJSValue) SetProperty(name *qtcore.QString, value *QJSValue) {
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValue11setPropertyERK7QStringRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:131
// index:1
// Public
// void setProperty(quint32, const QJSValue &)
func (this *QJSValue) SetProperty_1(arrayIndex uint, value *QJSValue) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValue11setPropertyEjRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arrayIndex, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalue.h:127
// index:0
// Public
// bool hasProperty(const QString &)
func (this *QJSValue) HasProperty(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue11hasPropertyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:128
// index:0
// Public
// bool hasOwnProperty(const QString &)
func (this *QJSValue) HasOwnProperty(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue14hasOwnPropertyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:133
// index:0
// Public
// bool deleteProperty(const QString &)
func (this *QJSValue) DeleteProperty(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QJSValue14deletePropertyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:135
// index:0
// Public
// bool isCallable()
func (this *QJSValue) IsCallable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue10isCallableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalue.h:141
// index:0
// Public
// QJSEngine * engine()
func (this *QJSValue) Engine() *QJSEngine /*777 QJSEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QJSValue6engineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJSEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QJSValue__SpecialValue = int

const QJSValue__NullValue QJSValue__SpecialValue = 0
const QJSValue__UndefinedValue QJSValue__SpecialValue = 1

//  body block end
