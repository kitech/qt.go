//  header block begin
// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>
package qtcore

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
type QJsonValueRef struct {
	*qtrt.CObject
}

func (this *QJsonValueRef) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qjsonvalue.h:174
// index:0
// inline
// void QJsonValueRef(class QJsonArray *, int)
func NewQJsonValueRef(array unsafe.Pointer, idx int) *QJsonValueRef {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValueRefC2EP10QJsonArrayi", ffiqt.FFI_TYPE_VOID, cthis, array, &idx)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueRefFromPointer(cthis)
	return gothis
}
func NewQJsonValueRefFromPointer(cthis unsafe.Pointer) *QJsonValueRef {
	return &QJsonValueRef{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qjsonvalue.h:176
// index:1
// inline
// void QJsonValueRef(class QJsonObject *, int)
func NewQJsonValueRef_1(object unsafe.Pointer, idx int) *QJsonValueRef {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValueRefC2EP11QJsonObjecti", ffiqt.FFI_TYPE_VOID, cthis, object, &idx)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValueRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:183
// index:0
// QVariant toVariant()
func (this *QJsonValueRef) ToVariant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef9toVariantEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:184
// index:0
// inline
// QJsonValue::Type type()
func (this *QJsonValueRef) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:185
// index:0
// inline
// bool isNull()
func (this *QJsonValueRef) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:186
// index:0
// inline
// bool isBool()
func (this *QJsonValueRef) IsBool() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6isBoolEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:187
// index:0
// inline
// bool isDouble()
func (this *QJsonValueRef) IsDouble() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8isDoubleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:188
// index:0
// inline
// bool isString()
func (this *QJsonValueRef) IsString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8isStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:189
// index:0
// inline
// bool isArray()
func (this *QJsonValueRef) IsArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef7isArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:190
// index:0
// inline
// bool isObject()
func (this *QJsonValueRef) IsObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8isObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:191
// index:0
// inline
// bool isUndefined()
func (this *QJsonValueRef) IsUndefined() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef11isUndefinedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:193
// index:0
// inline
// bool toBool()
func (this *QJsonValueRef) ToBool() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6toBoolEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:201
// index:1
// inline
// bool toBool(_Bool)
func (this *QJsonValueRef) ToBool_1(defaultValue bool) {
	// 1: (, defaultValue bool), (&defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef6toBoolEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:194
// index:0
// inline
// int toInt()
func (this *QJsonValueRef) ToInt() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef5toIntEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:202
// index:1
// inline
// int toInt(int)
func (this *QJsonValueRef) ToInt_1(defaultValue int) {
	// 1: (, defaultValue int), (&defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef5toIntEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:195
// index:0
// inline
// double toDouble()
func (this *QJsonValueRef) ToDouble() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toDoubleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:203
// index:1
// inline
// double toDouble(double)
func (this *QJsonValueRef) ToDouble_1(defaultValue float64) {
	// 1: (, defaultValue double), (&defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toDoubleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:196
// index:0
// inline
// QString toString()
func (this *QJsonValueRef) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:204
// index:1
// inline
// QString toString(const class QString &)
func (this *QJsonValueRef) ToString_1(defaultValue unsafe.Pointer) {
	// 1: (, defaultValue const QString &), (defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toStringERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:197
// index:0
// QJsonArray toArray()
func (this *QJsonValueRef) ToArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef7toArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonvalue.h:198
// index:0
// QJsonObject toObject()
func (this *QJsonValueRef) ToObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QJsonValueRef8toObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
