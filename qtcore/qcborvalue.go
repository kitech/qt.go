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
type QCborValue struct {
	*qtrt.CObject
}
type QCborValue_ITF interface {
	QCborValue_PTR() *QCborValue
}

func (ptr *QCborValue) QCborValue_PTR() *QCborValue { return ptr }

func (this *QCborValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborValue) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborValueFromPointer(cthis unsafe.Pointer) *QCborValue {
	return &QCborValue{&qtrt.CObject{cthis}}
}
func (*QCborValue) NewFromPointer(cthis unsafe.Pointer) *QCborValue {
	return NewQCborValueFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborvalue.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue()

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit() *QCborValue {
	return NewQCborValue()
}
func NewQCborValue() *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:130
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(QCborValue::Type)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit1(t_ int) *QCborValue {
	return NewQCborValue1(t_)
}
func NewQCborValue1(t_ int) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, t_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:132
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(bool)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit2(b_ bool) *QCborValue {
	return NewQCborValue2(b_)
}
func NewQCborValue2(b_ bool) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2Eb", qtrt.FFI_TYPE_POINTER, b_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:134
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(int)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit3(i int) *QCborValue {
	return NewQCborValue3(i)
}
func NewQCborValue3(i int) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2Ei", qtrt.FFI_TYPE_POINTER, i)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:135
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(unsigned int)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit4(u uint) *QCborValue {
	return NewQCborValue4(u)
}
func NewQCborValue4(u uint) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2Ej", qtrt.FFI_TYPE_POINTER, u)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:137
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(qint64)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit5(i int64) *QCborValue {
	return NewQCborValue5(i)
}
func NewQCborValue5(i int64) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2Ex", qtrt.FFI_TYPE_POINTER, i)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:138
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(double)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit6(v float64) *QCborValue {
	return NewQCborValue6(v)
}
func NewQCborValue6(v float64) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2Ed", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:139
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(QCborSimpleType)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit7(st int) *QCborValue {
	return NewQCborValue7(st)
}
func NewQCborValue7(st int) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2E15QCborSimpleType", qtrt.FFI_TYPE_POINTER, st)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:141
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QByteArray &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit8(ba QByteArray_ITF) *QCborValue {
	return NewQCborValue8(ba)
}
func NewQCborValue8(ba QByteArray_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:142
// index:9
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QString &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit9(s string) *QCborValue {
	return NewQCborValue9(s)
}
func NewQCborValue9(s string) *QCborValue {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:143
// index:10
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(QLatin1String)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit10(s QLatin1String_ITF /*123*/) *QCborValue {
	return NewQCborValue10(s)
}
func NewQCborValue10(s QLatin1String_ITF /*123*/) *QCborValue {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2E13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:145
// index:11
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(const char *)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit11(s string) *QCborValue {
	return NewQCborValue11(s)
}
func NewQCborValue11(s string) *QCborValue {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:147
// index:12
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QCborArray &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit12(a QCborArray_ITF) *QCborValue {
	return NewQCborValue12(a)
}
func NewQCborValue12(a QCborArray_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if a != nil && a.QCborArray_PTR() != nil {
		convArg0 = a.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK10QCborArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:148
// index:13
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(QCborArray &&)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit13(a unsafe.Pointer /*333*/) *QCborValue {
	return NewQCborValue13(a)
}
func NewQCborValue13(a unsafe.Pointer /*333*/) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2EO10QCborArray", qtrt.FFI_TYPE_POINTER, a)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:149
// index:14
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QCborMap &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit14(m QCborMap_ITF) *QCborValue {
	return NewQCborValue14(m)
}
func NewQCborValue14(m QCborMap_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if m != nil && m.QCborMap_PTR() != nil {
		convArg0 = m.QCborMap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK8QCborMap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:150
// index:15
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(QCborMap &&)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit15(m unsafe.Pointer /*333*/) *QCborValue {
	return NewQCborValue15(m)
}
func NewQCborValue15(m unsafe.Pointer /*333*/) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2EO8QCborMap", qtrt.FFI_TYPE_POINTER, m)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:151
// index:16
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(QCborTag, const QCborValue &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit16(tag int, taggedValue QCborValue_ITF) *QCborValue {
	return NewQCborValue16(tag, taggedValue)
}
func NewQCborValue16(tag int, taggedValue QCborValue_ITF) *QCborValue {
	var convArg1 unsafe.Pointer
	if taggedValue != nil && taggedValue.QCborValue_PTR() != nil {
		convArg1 = taggedValue.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2E8QCborTagRKS_", qtrt.FFI_TYPE_POINTER, tag, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:151
// index:16
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(QCborTag, const QCborValue &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit16p(tag int) *QCborValue {
	return NewQCborValue16p(tag)
}
func NewQCborValue16p(tag int) *QCborValue {
	// arg: 1, const QCborValue &=LValueReference, QCborValue=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2E8QCborTagRKS_", qtrt.FFI_TYPE_POINTER, tag, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:152
// index:17
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(QCborKnownTags, const QCborValue &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit17(t_ int, tv QCborValue_ITF) *QCborValue {
	return NewQCborValue17(t_, tv)
}
func NewQCborValue17(t_ int, tv QCborValue_ITF) *QCborValue {
	var convArg1 unsafe.Pointer
	if tv != nil && tv.QCborValue_PTR() != nil {
		convArg1 = tv.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2E14QCborKnownTagsRKS_", qtrt.FFI_TYPE_POINTER, t_, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:152
// index:17
// Public inline Visibility=Default Availability=Available
// [-2] void QCborValue(QCborKnownTags, const QCborValue &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit17p(t_ int) *QCborValue {
	return NewQCborValue17p(t_)
}
func NewQCborValue17p(t_ int) *QCborValue {
	// arg: 1, const QCborValue &=LValueReference, QCborValue=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2E14QCborKnownTagsRKS_", qtrt.FFI_TYPE_POINTER, t_, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:156
// index:18
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QDateTime &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit18(dt QDateTime_ITF) *QCborValue {
	return NewQCborValue18(dt)
}
func NewQCborValue18(dt QDateTime_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDateTime_PTR() != nil {
		convArg0 = dt.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK9QDateTime", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:157
// index:19
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QUrl &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit19(url QUrl_ITF) *QCborValue {
	return NewQCborValue19(url)
}
func NewQCborValue19(url QUrl_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:158
// index:20
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QRegularExpression &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit20(rx QRegularExpression_ITF) *QCborValue {
	return NewQCborValue20(rx)
}
func NewQCborValue20(rx QRegularExpression_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegularExpression_PTR() != nil {
		convArg0 = rx.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:159
// index:21
// Public Visibility=Default Availability=Available
// [-2] void QCborValue(const QUuid &)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit21(uuid QUuid_ITF) *QCborValue {
	return NewQCborValue21(uuid)
}
func NewQCborValue21(uuid QUuid_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if uuid != nil && uuid.QUuid_PTR() != nil {
		convArg0 = uuid.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2ERK5QUuid", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:164
// index:22
// Public inline Visibility=Default Availability=NotAvailable
// [-2] void QCborValue(const void *)

/*
Creates a QCborValue of the Undefined type.

CBOR undefined values are used to indicate missing information, usually as a result of a previous operation that did not complete as expected. They are also used by the QCborArray and QCborMap API to indicate the searched item was not found.

Undefined values are represented by the Undefined simple type. Because of that, QCborValues with undefined values will also return true for isSimpleType() and isSimpleType(QCborSimpleType::Undefined).

Undefined values are different from null values.

QCborValue objects with undefined values are also different from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also isUndefined(), isNull(), and isSimpleType().
*/
func (*QCborValue) NewForInherit22(arg0 unsafe.Pointer /*666*/) *QCborValue {
	return NewQCborValue22(arg0)
}
func NewQCborValue22(arg0 unsafe.Pointer /*666*/) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueC2EPKv", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborValue)
	return gothis
}

// /usr/include/qt/QtCore/qcborvalue.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QCborValue()

/*

 */
func DeleteQCborValue(this *QCborValue) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcborvalue.h:173
// index:0
// Public Visibility=Default Availability=Available
// [24] QCborValue & operator=(const QCborValue &)

/*

 */
func (this *QCborValue) Operator_equal(other QCborValue_ITF) *QCborValue {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:174
// index:1
// Public inline Visibility=Default Availability=Available
// [24] QCborValue & operator=(QCborValue &&)

/*

 */
func (this *QCborValue) Operator_equal1(other unsafe.Pointer /*333*/) *QCborValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValueaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCborValue &)

/*
Swaps the contents of this QCborValue object and other.
*/
func (this *QCborValue) Swap(other QCborValue_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QCborValue::Type type() const

/*
Returns the type of this QCborValue. The type can also later be retrieved by one of the "isXxx" functions.

See also isInteger(), isByteArray(), isString(), isArray(), isMap(), isTag(), isFalse(), isTrue(), isBool(), isNull(), isUndefined, isDouble(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValue) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:190
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInteger() const

/*
Returns true if this QCborValue is of the integer type. The integer value can be retrieved using toInteger().

See also type() and toInteger().
*/
func (this *QCborValue) IsInteger() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue9isIntegerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:191
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isByteArray() const

/*
Returns true if this QCborValue is of the byte array type. The byte array value can be retrieved using toByteArray().

See also type() and toByteArray().
*/
func (this *QCborValue) IsByteArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11isByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:192
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isString() const

/*
Returns true if this QCborValue is of the string type. The string value can be retrieved using toString().

See also type() and toString().
*/
func (this *QCborValue) IsString() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue8isStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isArray() const

/*
Returns true if this QCborValue is of the array type. The array value can be retrieved using toArray().

See also type() and toArray().
*/
func (this *QCborValue) IsArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue7isArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:194
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMap() const

/*
Returns true if this QCborValue is of the map type. The map value can be retrieved using toMap().

See also type() and toMap().
*/
func (this *QCborValue) IsMap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue5isMapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTag() const

/*
Returns true if this QCborValue is of the tag type. The tag value can be retrieved using tag() and the tagged value using taggedValue().

This function also returns true for extended types that the API recognizes. For code that handles extended types directly before the Qt API is updated to support them, it is possible to recreate the tag + tagged value pair by using taggedValue().

See also type(), tag(), taggedValue(), and taggedValue().
*/
func (this *QCborValue) IsTag() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue5isTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFalse() const

/*
Returns true if this QCborValue is a boolean with false value. This function exists because, internally, CBOR booleans are stored as two separate types, one for true and one for false.

See also type(), isBool(), isTrue(), and toBool().
*/
func (this *QCborValue) IsFalse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue7isFalseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:197
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTrue() const

/*
Returns true if this QCborValue is a boolean with true value. This function exists because, internally, CBOR booleans are stored as two separate types, one for false and one for true.

See also type(), isBool(), isFalse(), and toBool().
*/
func (this *QCborValue) IsTrue() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6isTrueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:198
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBool() const

/*
Returns true if this QCborValue is a boolean. The value can be retrieved using toBool().

See also type(), toBool(), isTrue(), and isFalse().
*/
func (this *QCborValue) IsBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6isBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:199
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this QCborValue is of the null type.

CBOR null values are used to indicate optional values that were not provided. They are distinct from undefined values, in that null values are usually not the result of an earlier error or problem.

Null values are distinct from undefined values and from invalid QCborValue objects. The API will not create invalid QCborValues, but they may exist as a result of a parsing error.

See also type(), isUndefined(), and isInvalid().
*/
func (this *QCborValue) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:200
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
func (this *QCborValue) IsUndefined() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11isUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:201
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDouble() const

/*
Returns true if this QCborValue is of the floating-point type. The value can be retrieved using toDouble().

See also type() and toDouble().
*/
func (this *QCborValue) IsDouble() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue8isDoubleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDateTime() const

/*
Returns true if this QCborValue is of the date/time type. The value can be retrieved using toDateTime(). Date/times are extended types that use the tag DateTime.

Additionally, when decoding from a CBOR stream, QCborValue will interpret tags of value UnixTime_t and convert them to the equivalent date/time.

See also type() and toDateTime().
*/
func (this *QCborValue) IsDateTime() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue10isDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUrl() const

/*
Returns true if this QCborValue is of the URL type. The URL value can be retrieved using toUrl().

See also type() and toUrl().
*/
func (this *QCborValue) IsUrl() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue5isUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:204
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRegularExpression() const

/*
Returns true if this QCborValue contains a regular expression's pattern. The pattern can be retrieved using toRegularExpression().

See also type() and toRegularExpression().
*/
func (this *QCborValue) IsRegularExpression() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue19isRegularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:205
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUuid() const

/*
Returns true if this QCborValue contains a UUID. The value can be retrieved using toUuid().

See also type() and toUuid().
*/
func (this *QCborValue) IsUuid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6isUuidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isInvalid() const

/*
Returns true if this QCborValue is not of any valid type. Invalid QCborValues are distinct from those with undefined values and they usually represent a decoding error.

See also isUndefined() and isNull().
*/
func (this *QCborValue) IsInvalid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue9isInvalidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isContainer() const

/*
This convenience function returns true if the QCborValue is either an array or a map.

See also isArray() and isMap().
*/
func (this *QCborValue) IsContainer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11isContainerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSimpleType() const

/*
Returns true if this QCborValue is of one of the CBOR simple types. The type itself can later be retrieved using type(), even for types that don't have an enumeration in the API. They can also be checked with the isSimpleType(QCborSimpleType) overload.

See also QCborSimpleType, isSimpleType(QCborSimpleType), and toSimpleType().
*/
func (this *QCborValue) IsSimpleType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue12isSimpleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:213
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool isSimpleType(QCborSimpleType) const

/*
Returns true if this QCborValue is of one of the CBOR simple types. The type itself can later be retrieved using type(), even for types that don't have an enumeration in the API. They can also be checked with the isSimpleType(QCborSimpleType) overload.

See also QCborSimpleType, isSimpleType(QCborSimpleType), and toSimpleType().
*/
func (this *QCborValue) IsSimpleType1(st int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue12isSimpleTypeE15QCborSimpleType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), st)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QCborSimpleType toSimpleType(QCborSimpleType) const

/*
Returns the simple type this QCborValue is of, if it is a simple type. If it is not a simple type, it returns defaultValue.

The following types are simple types and this function will return the listed values:


 QCborValue::FalseQCborSimpleType::False
QCborValue::TrueQCborSimpleType::True
QCborValue::NullQCborSimpleType::Null
QCborValue::UndefinedQCborSimpleType::Undefined


See also type(), isSimpleType(), isBool(), isTrue(), isFalse(), isTrue(), isNull(), and isUndefined().
*/
func (this *QCborValue) ToSimpleType(defaultValue int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue12toSimpleTypeE15QCborSimpleType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QCborSimpleType toSimpleType(QCborSimpleType) const

/*
Returns the simple type this QCborValue is of, if it is a simple type. If it is not a simple type, it returns defaultValue.

The following types are simple types and this function will return the listed values:


 QCborValue::FalseQCborSimpleType::False
QCborValue::TrueQCborSimpleType::True
QCborValue::NullQCborSimpleType::Null
QCborValue::UndefinedQCborSimpleType::Undefined


See also type(), isSimpleType(), isBool(), isTrue(), isFalse(), isTrue(), isNull(), and isUndefined().
*/
func (this *QCborValue) ToSimpleTypep() int {
	// arg: 0, QCborSimpleType=Enum, QCborSimpleType=Enum, , Invalid
	defaultValue := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue12toSimpleTypeE15QCborSimpleType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toInteger(qint64) const

/*
Returns the integer value stored in this QCborValue, if it is of the integer type. If it is of the Double type, this function returns the floating point value converted to integer. In any other case, it returns defaultValue.

See also isInteger(), isDouble(), and toDouble().
*/
func (this *QCborValue) ToInteger(defaultValue int64) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue9toIntegerEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborvalue.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qint64 toInteger(qint64) const

/*
Returns the integer value stored in this QCborValue, if it is of the integer type. If it is of the Double type, this function returns the floating point value converted to integer. In any other case, it returns defaultValue.

See also isInteger(), isDouble(), and toDouble().
*/
func (this *QCborValue) ToIntegerp() int64 {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long, LongLong
	defaultValue := int64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue9toIntegerEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborvalue.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Returns the boolean value stored in this QCborValue, if it is of a boolean type. Otherwise, it returns defaultValue.

See also isBool(), isTrue(), and isFalse().
*/
func (this *QCborValue) ToBool(defaultValue bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toBool(bool) const

/*
Returns the boolean value stored in this QCborValue, if it is of a boolean type. Otherwise, it returns defaultValue.

See also isBool(), isTrue(), and isFalse().
*/
func (this *QCborValue) ToBoolp() bool {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	defaultValue := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6toBoolEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Returns the floating point value stored in this QCborValue, if it is of the Double type. If it is of the Integer type, this function returns the integer value converted to double. In any other case, it returns defaultValue.

See also isDouble(), isInteger(), and toInteger().
*/
func (this *QCborValue) ToDouble(defaultValue float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qcborvalue.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double toDouble(double) const

/*
Returns the floating point value stored in this QCborValue, if it is of the Double type. If it is of the Integer type, this function returns the integer value converted to double. In any other case, it returns defaultValue.

See also isDouble(), isInteger(), and toInteger().
*/
func (this *QCborValue) ToDoublep() float64 {
	// arg: 0, double=Double, =Invalid, , Invalid
	defaultValue := float64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue8toDoubleEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qcborvalue.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] QCborTag tag(QCborTag) const

/*
Returns the tag of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that number. To retrieve the representation, use taggedValue().

See also isTag(), taggedValue(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValue) Tag(defaultValue int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue3tagE8QCborTag", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] QCborTag tag(QCborTag) const

/*
Returns the tag of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that number. To retrieve the representation, use taggedValue().

See also isTag(), taggedValue(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValue) Tagp() int {
	// arg: 0, QCborTag=Enum, QCborTag=Enum, , Invalid
	defaultValue := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue3tagE8QCborTag", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultValue)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:230
// index:0
// Public Visibility=Default Availability=Available
// [24] QCborValue taggedValue(const QCborValue &) const

/*
Returns the tagged value of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that representation. To retrieve the tag, use tag().

See also isTag(), tag(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValue) TaggedValue(defaultValue QCborValue_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QCborValue_PTR() != nil {
		convArg0 = defaultValue.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11taggedValueERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:230
// index:0
// Public Visibility=Default Availability=Available
// [24] QCborValue taggedValue(const QCborValue &) const

/*
Returns the tagged value of this extended QCborValue object, if it is of the tag type, defaultValue otherwise.

CBOR represents extended types by associating a number (the tag) with a stored representation. This function returns that representation. To retrieve the tag, use tag().

See also isTag(), tag(), isDateTime(), isUrl(), isRegularExpression(), and isUuid().
*/
func (this *QCborValue) TaggedValuep() *QCborValue /*123*/ {
	// arg: 0, const QCborValue &=LValueReference, QCborValue=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11taggedValueERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toByteArray(const QByteArray &) const

/*
Returns the byte array value stored in this QCborValue, if it is of the byte array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QByteArray.

See also isByteArray(), isString(), and toString().
*/
func (this *QCborValue) ToByteArray(defaultValue QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QByteArray_PTR() != nil {
		convArg0 = defaultValue.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11toByteArrayERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toByteArray(const QByteArray &) const

/*
Returns the byte array value stored in this QCborValue, if it is of the byte array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QByteArray.

See also isByteArray(), isString(), and toString().
*/
func (this *QCborValue) ToByteArrayp() *QByteArray /*123*/ {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11toByteArrayERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Returns the string value stored in this QCborValue, if it is of the string type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QString.

See also isString(), isByteArray(), and toByteArray().
*/
func (this *QCborValue) ToString(defaultValue string) string {
	var tmpArg0 = NewQString5(defaultValue)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcborvalue.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString(const QString &) const

/*
Returns the string value stored in this QCborValue, if it is of the string type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QString.

See also isString(), isByteArray(), and toByteArray().
*/
func (this *QCborValue) ToStringp() string {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue8toStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcborvalue.h:234
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QDateTime &) const

/*
Returns the date/time value stored in this QCborValue, if it is of the date/time extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QDateTime.

See also isDateTime(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToDateTime(defaultValue QDateTime_ITF) *QDateTime /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QDateTime_PTR() != nil {
		convArg0 = defaultValue.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue10toDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:234
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime(const QDateTime &) const

/*
Returns the date/time value stored in this QCborValue, if it is of the date/time extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QDateTime.

See also isDateTime(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToDateTimep() *QDateTime /*123*/ {
	// arg: 0, const QDateTime &=LValueReference, QDateTime=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue10toDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl toUrl(const QUrl &) const

/*
Returns the URL value stored in this QCborValue, if it is of the URL extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUrl.

See also isUrl(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToUrl(defaultValue QUrl_ITF) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QUrl_PTR() != nil {
		convArg0 = defaultValue.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue5toUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl toUrl(const QUrl &) const

/*
Returns the URL value stored in this QCborValue, if it is of the URL extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUrl.

See also isUrl(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToUrlp() *QUrl /*123*/ {
	// arg: 0, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg0 = NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue5toUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression(const QRegularExpression &) const

/*
Returns the regular expression value stored in this QCborValue, if it is of the regular expression pattern extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QRegularExpression.

See also isRegularExpression(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToRegularExpression(defaultValue QRegularExpression_ITF) *QRegularExpression /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QRegularExpression_PTR() != nil {
		convArg0 = defaultValue.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue19toRegularExpressionERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression(const QRegularExpression &) const

/*
Returns the regular expression value stored in this QCborValue, if it is of the regular expression pattern extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QRegularExpression.

See also isRegularExpression(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToRegularExpressionp() *QRegularExpression /*123*/ {
	// arg: 0, const QRegularExpression &=LValueReference, QRegularExpression=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue19toRegularExpressionERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:237
// index:0
// Public Visibility=Default Availability=Available
// [16] QUuid toUuid(const QUuid &) const

/*
Returns the UUID value stored in this QCborValue, if it is of the UUID extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUuid.

See also isUuid(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToUuid(defaultValue QUuid_ITF) *QUuid /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QUuid_PTR() != nil {
		convArg0 = defaultValue.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6toUuidERK5QUuid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:237
// index:0
// Public Visibility=Default Availability=Available
// [16] QUuid toUuid(const QUuid &) const

/*
Returns the UUID value stored in this QCborValue, if it is of the UUID extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QUuid.

See also isUuid(), isTag(), and taggedValue().
*/
func (this *QCborValue) ToUuidp() *QUuid /*123*/ {
	// arg: 0, const QUuid &=LValueReference, QUuid=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue6toUuidERK5QUuid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:244
// index:0
// Public Visibility=Default Availability=Available
// [8] QCborArray toArray() const

/*
Returns the array value stored in this QCborValue, if it is of the array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QCborArray.

See also isArray(), isByteArray(), isMap(), isContainer(), and toMap().
*/
func (this *QCborValue) ToArray() *QCborArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue7toArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:245
// index:1
// Public Visibility=Default Availability=Available
// [8] QCborArray toArray(const QCborArray &) const

/*
Returns the array value stored in this QCborValue, if it is of the array type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QCborArray.

See also isArray(), isByteArray(), isMap(), isContainer(), and toMap().
*/
func (this *QCborValue) ToArray1(defaultValue QCborArray_ITF) *QCborArray /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QCborArray_PTR() != nil {
		convArg0 = defaultValue.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue7toArrayERK10QCborArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:250
// index:0
// Public Visibility=Default Availability=Available
// [24] const QCborValue operator[](const QString &) const

/*

 */
func (this *QCborValue) Operator_get_index(key string) *QCborValue /*123*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValueixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:251
// index:1
// Public Visibility=Default Availability=Available
// [24] const QCborValue operator[](QLatin1String) const

/*

 */
func (this *QCborValue) Operator_get_index1(key QLatin1String_ITF /*123*/) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValueixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:252
// index:2
// Public Visibility=Default Availability=Available
// [24] const QCborValue operator[](qint64) const

/*

 */
func (this *QCborValue) Operator_get_index2(key int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValueixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:254
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QCborValue &) const

/*
Compares this value and other, and returns an integer that indicates whether this value should be sorted prior to (if the result is negative) or after other (if the result is positive). If this function returns 0, the two values are equal and hold the same contents.

If each QCborValue contains an array or map, the comparison is recursive to elements contained in them.
*/
func (this *QCborValue) Compare(other QCborValue_ITF) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcborvalue.h:264
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QCborValue &) const

/*

 */
func (this *QCborValue) Operator_equal_equal(other QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValueeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:266
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCborValue &) const

/*

 */
func (this *QCborValue) Operator_not_equal(other QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValueneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:268
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QCborValue &) const

/*

 */
func (this *QCborValue) Operator_less_than(other QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborValue_PTR() != nil {
		convArg0 = other.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValueltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborvalue.h:272
// index:0
// Public static Visibility=Default Availability=Available
// [24] QCborValue fromVariant(const QVariant &)

/*
Converts the QVariant variant into QCborValue and returns it.

QVariants may contain a large list of different meta types, many of which have no corresponding representation in CBOR. That includes all user-defined meta types. When preparing transmission using CBOR, it is suggested to encode carefully each value to prevent loss of representation.

The following table lists the conversion this function will apply:


 Qt (C++) typeCBOR type
invalid (QVariant())Undefined
boolBool
std::nullptr_tNull
short, ushort, int, uint, qint64Integer
quint64Integer, but they are cast to qint64 first so values higher than 263-1 (INT64_MAX) will be wrapped to negative
float, doubleDouble
QByteArrayByteArray
QDateTimeDateTime
QCborSimpleTypeSimple type
QJsonArrayArray, converted using QCborArray::formJsonArray()
QJsonDocumentArray or Map
QJsonObjectMap, converted using QCborMap::fromJsonObject()
QJsonValueconverted using fromJsonValue()
QRegularExpressionRegularExpression
QStringString
QStringListArray
QVariantHashMap
QVariantListArray
QVariantMapMap
QUrlUrl
QUuidUuid


For any other types, this function will return Null if the QVariant itself is null, and otherwise will try to convert to string using QVariant::toString(). If the conversion to string fails, this function returns Undefined.

Please note that the conversions via QVariant::toString() are subject to change at any time. QCborValue may be extended in the future to support more types, which will result in a change in how this function performs conversions.

See also toVariant(), fromJsonValue(), QCborArray::toVariantList(), and QCborMap::toVariantMap().
*/
func (this *QCborValue) FromVariant(variant QVariant_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue11fromVariantERK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}
func QCborValue_FromVariant(variant QVariant_ITF) *QCborValue /*123*/ {
	var nilthis *QCborValue
	rv := nilthis.FromVariant(variant)
	return rv
}

// /usr/include/qt/QtCore/qcborvalue.h:273
// index:0
// Public Visibility=Default Availability=Available
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
func (this *QCborValue) ToVariant() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue9toVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:274
// index:0
// Public static Visibility=Default Availability=Available
// [24] QCborValue fromJsonValue(const QJsonValue &)

/*
Converts the JSON value contained in v into its corresponding CBOR value and returns it. There is no data loss in converting from JSON to CBOR, as the CBOR type set is richer than JSON's. Additionally, values converted to CBOR using this function can be converted back to JSON using toJsonValue() with no data loss.

The following table lists the mapping of JSON types to CBOR types:


 JSON TypeCBOR Type
BoolBool
NumberInteger (if the number has no fraction and is in the qint64 range) or Double
StringString
ArrayArray
ObjectMap
NullNull


QJsonValue can also be undefined, indicating a previous operation that failed to complete (for example, searching for a key not present in an object). Undefined values are not JSON types and may not appear in JSON arrays and objects, but this function does return the QCborValue undefined value if the corresponding QJsonValue is undefined.

See also toJsonValue(), fromVariant(), QCborArray::fromJsonArray(), and QCborMap::fromJsonObject().
*/
func (this *QCborValue) FromJsonValue(v QJsonValue_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue13fromJsonValueERK10QJsonValue", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}
func QCborValue_FromJsonValue(v QJsonValue_ITF) *QCborValue /*123*/ {
	var nilthis *QCborValue
	rv := nilthis.FromJsonValue(v)
	return rv
}

// /usr/include/qt/QtCore/qcborvalue.h:275
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
func (this *QCborValue) ToJsonValue() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue11toJsonValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [24] QCborValue fromCbor(QCborStreamReader &)

/*
Decodes one item from the CBOR stream found in reader and returns the equivalent representation. This function is recursive: if the item is a map or array, it will decode all items found in that map or array, until the outermost object is finished.

This function need not be used on the root element of a QCborStreamReader. For example, the following code illustrates how to skip the CBOR signature tag from the beginning of a file:


      if (reader.isTag() && reader.toTag() == QCborKnownTags::Signature)
          reader.next();

      QCborValue contents = QCborValue::fromCbor(reader);



The returned value may be partially complete and indistinguishable from a valid QCborValue even if the decoding failed. To determine if there was an error, check if reader.lastError() is indicating an error condition. This function stops decoding immediately after the first error.

See also toCbor(), toDiagnosticNotation(), toVariant(), and toJsonValue().
*/
func (this *QCborValue) FromCbor(reader QCborStreamReader_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if reader != nil && reader.QCborStreamReader_PTR() != nil {
		convArg0 = reader.QCborStreamReader_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue8fromCborER17QCborStreamReader", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}
func QCborValue_FromCbor(reader QCborStreamReader_ITF) *QCborValue /*123*/ {
	var nilthis *QCborValue
	rv := nilthis.FromCbor(reader)
	return rv
}

// /usr/include/qt/QtCore/qcborvalue.h:283
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toCbor(QCborValue::EncodingOptions)

/*
Encodes this QCborValue object to its CBOR representation, using the options specified in opt, and return the byte array containing that representation.

This function will not fail, except if this QCborValue or any of the contained items, if this is a map or array, are invalid. Invalid types are not produced normally by the API, but can result from decoding errors.

By default, this function performs no transformation on the values in the QCborValue, writing all floating point directly as double-precision (double) types. If the UseFloat option is specified, it will use single precision (float) for any floating point value for which there's no loss of precision in using that representation. That includes infinities and NaN values.

Similarly, if UseFloat16 is specified, this function will try to use half-precision (qfloat16) floating point if the conversion to that results in no loss of precision. This is always true for infinities and NaN.

If UseIntegers is specified, it will use integers for any floating point value that contains an actual integer.

See also fromCbor(), fromVariant(), and fromJsonValue().
*/
func (this *QCborValue) ToCbor(opt int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue6toCborE6QFlagsINS_14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:283
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toCbor(QCborValue::EncodingOptions)

/*
Encodes this QCborValue object to its CBOR representation, using the options specified in opt, and return the byte array containing that representation.

This function will not fail, except if this QCborValue or any of the contained items, if this is a map or array, are invalid. Invalid types are not produced normally by the API, but can result from decoding errors.

By default, this function performs no transformation on the values in the QCborValue, writing all floating point directly as double-precision (double) types. If the UseFloat option is specified, it will use single precision (float) for any floating point value for which there's no loss of precision in using that representation. That includes infinities and NaN values.

Similarly, if UseFloat16 is specified, this function will try to use half-precision (qfloat16) floating point if the conversion to that results in no loss of precision. This is always true for infinities and NaN.

If UseIntegers is specified, it will use integers for any floating point value that contains an actual integer.

See also fromCbor(), fromVariant(), and fromJsonValue().
*/
func (this *QCborValue) ToCborp() *QByteArray /*123*/ {
	// arg: 0, QCborValue::EncodingOptions=Typedef, QCborValue::EncodingOptions=Typedef, QFlags<QCborValue::EncodingOption>, Unexposed
	opt := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue6toCborE6QFlagsINS_14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:284
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
func (this *QCborValue) ToCbor1(writer QCborStreamWriter_ITF, opt int) {
	var convArg0 unsafe.Pointer
	if writer != nil && writer.QCborStreamWriter_PTR() != nil {
		convArg0 = writer.QCborStreamWriter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue6toCborER17QCborStreamWriter6QFlagsINS_14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, opt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:284
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
func (this *QCborValue) ToCbor1p(writer QCborStreamWriter_ITF) {
	var convArg0 unsafe.Pointer
	if writer != nil && writer.QCborStreamWriter_PTR() != nil {
		convArg0 = writer.QCborStreamWriter_PTR().GetCthis()
	}
	// arg: 1, QCborValue::EncodingOptions=Typedef, QCborValue::EncodingOptions=Typedef, QFlags<QCborValue::EncodingOption>, Unexposed
	opt := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborValue6toCborER17QCborStreamWriter6QFlagsINS_14EncodingOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, opt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborvalue.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toDiagnosticNotation(QCborValue::DiagnosticNotationOptions) const

/*
Creates the diagnostic notation equivalent of this CBOR object and returns it. The opts parameter controls the dialect of the notation. Diagnostic notation is useful in debugging, to aid the developer in understanding what value is stored in the QCborValue or in a CBOR stream. For that reason, the Qt API provides no support for parsing the diagnostic back into the in-memory format or CBOR stream, though the representation is unique and it would be possible.

CBOR diagnostic notation is specified by section 6 of RFC 7049. It is a text representation of the CBOR stream and it is very similar to JSON, but it supports the CBOR types not found in JSON. The extended format enabled by the ExtendedFormat flag is currently in some IETF drafts and its format is subject to change.

This function produces the equivalent representation of the stream that toCbor() would produce, without any transformation option provided there. This also implies this function may not produce a representation of the stream that was used to create the object, if it was created using fromCbor(), as that function may have applied transformations. For a high-fidelity notation of a stream, without transformation, see the cbordump example.

See also toCbor() and QJsonDocument::toJson().
*/
func (this *QCborValue) ToDiagnosticNotation(opts int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue20toDiagnosticNotationE6QFlagsINS_24DiagnosticNotationOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opts)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcborvalue.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toDiagnosticNotation(QCborValue::DiagnosticNotationOptions) const

/*
Creates the diagnostic notation equivalent of this CBOR object and returns it. The opts parameter controls the dialect of the notation. Diagnostic notation is useful in debugging, to aid the developer in understanding what value is stored in the QCborValue or in a CBOR stream. For that reason, the Qt API provides no support for parsing the diagnostic back into the in-memory format or CBOR stream, though the representation is unique and it would be possible.

CBOR diagnostic notation is specified by section 6 of RFC 7049. It is a text representation of the CBOR stream and it is very similar to JSON, but it supports the CBOR types not found in JSON. The extended format enabled by the ExtendedFormat flag is currently in some IETF drafts and its format is subject to change.

This function produces the equivalent representation of the stream that toCbor() would produce, without any transformation option provided there. This also implies this function may not produce a representation of the stream that was used to create the object, if it was created using fromCbor(), as that function may have applied transformations. For a high-fidelity notation of a stream, without transformation, see the cbordump example.

See also toCbor() and QJsonDocument::toJson().
*/
func (this *QCborValue) ToDiagnosticNotationp() string {
	// arg: 0, QCborValue::DiagnosticNotationOptions=Typedef, QCborValue::DiagnosticNotationOptions=Typedef, QFlags<QCborValue::DiagnosticNotationOption>, Unexposed
	opts := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborValue20toDiagnosticNotationE6QFlagsINS_24DiagnosticNotationOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opts)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*


 */
type QCborValue__EncodingOption = int

//
const QCborValue__SortKeysInMaps QCborValue__EncodingOption = 1

//
const QCborValue__UseFloat QCborValue__EncodingOption = 2

//
const QCborValue__UseFloat16 QCborValue__EncodingOption = 6

//
const QCborValue__UseIntegers QCborValue__EncodingOption = 8

//
const QCborValue__NoTransformation QCborValue__EncodingOption = 0

func (this *QCborValue) EncodingOptionItemName(val int) string {
	switch val {
	case QCborValue__SortKeysInMaps: // 1
		return "SortKeysInMaps"
	case QCborValue__UseFloat: // 2
		return "UseFloat"
	case QCborValue__UseFloat16: // 6
		return "UseFloat16"
	case QCborValue__UseIntegers: // 8
		return "UseIntegers"
	case QCborValue__NoTransformation: // 0
		return "NoTransformation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCborValue_EncodingOptionItemName(val int) string {
	var nilthis *QCborValue
	return nilthis.EncodingOptionItemName(val)
}

/*


 */
type QCborValue__DiagnosticNotationOption = int

//
const QCborValue__Compact QCborValue__DiagnosticNotationOption = 0

//
const QCborValue__LineWrapped QCborValue__DiagnosticNotationOption = 1

//
const QCborValue__ExtendedFormat QCborValue__DiagnosticNotationOption = 2

func (this *QCborValue) DiagnosticNotationOptionItemName(val int) string {
	switch val {
	case QCborValue__Compact: // 0
		return "Compact"
	case QCborValue__LineWrapped: // 1
		return "LineWrapped"
	case QCborValue__ExtendedFormat: // 2
		return "ExtendedFormat"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCborValue_DiagnosticNotationOptionItemName(val int) string {
	var nilthis *QCborValue
	return nilthis.DiagnosticNotationOptionItemName(val)
}

/*
This enum represents the QCborValue type. It is returned by the type() function.

The CBOR built-in types are:

QCborValue::FalseSimpleType + int(QCborSimpleType::False)bool: the simple type for value false
QCborValue::TrueSimpleType + int(QCborSimpleType::True)bool: the simple type for value true
QCborValue::NullSimpleType + int(QCborSimpleType::Null)std::nullptr_t: the simple type for the null value
QCborValue::UndefinedSimpleType + int(QCborSimpleType::Undefined)(no type) the simple type for the undefined value


Additionally, QCborValue can represent extended types:



See also type().

*/
type QCborValue__Type = int

//
const QCborValue__Integer QCborValue__Type = 0

//
const QCborValue__ByteArray QCborValue__Type = 64

//
const QCborValue__String QCborValue__Type = 96

//
const QCborValue__Array QCborValue__Type = 128

//
const QCborValue__Map QCborValue__Type = 160

//
const QCborValue__Tag QCborValue__Type = 192

//
const QCborValue__SimpleType QCborValue__Type = 256

//
const QCborValue__False QCborValue__Type = 276

//
const QCborValue__True QCborValue__Type = 277

//
const QCborValue__Null QCborValue__Type = 278

//
const QCborValue__Undefined QCborValue__Type = 279

//
const QCborValue__Double QCborValue__Type = 514

//
const QCborValue__DateTime QCborValue__Type = 65536

//
const QCborValue__Url QCborValue__Type = 65568

//
const QCborValue__RegularExpression QCborValue__Type = 65571

//
const QCborValue__Uuid QCborValue__Type = 65573

//
const QCborValue__Invalid QCborValue__Type = -1

func (this *QCborValue) TypeItemName(val int) string {
	switch val {
	case QCborValue__Integer: // 0
		return "Integer"
	case QCborValue__ByteArray: // 64
		return "ByteArray"
	case QCborValue__String: // 96
		return "String"
	case QCborValue__Array: // 128
		return "Array"
	case QCborValue__Map: // 160
		return "Map"
	case QCborValue__Tag: // 192
		return "Tag"
	case QCborValue__SimpleType: // 256
		return "SimpleType"
	case QCborValue__False: // 276
		return "False"
	case QCborValue__True: // 277
		return "True"
	case QCborValue__Null: // 278
		return "Null"
	case QCborValue__Undefined: // 279
		return "Undefined"
	case QCborValue__Double: // 514
		return "Double"
	case QCborValue__DateTime: // 65536
		return "DateTime"
	case QCborValue__Url: // 65568
		return "Url"
	case QCborValue__RegularExpression: // 65571
		return "RegularExpression"
	case QCborValue__Uuid: // 65573
		return "Uuid"
	case QCborValue__Invalid: // -1
		return "Invalid"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCborValue_TypeItemName(val int) string {
	var nilthis *QCborValue
	return nilthis.TypeItemName(val)
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
