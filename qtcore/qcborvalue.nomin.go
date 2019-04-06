//  header block begin

// +build !minimal

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

// /usr/include/qt/QtCore/qcborvalue.h:161
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

// /usr/include/qt/QtCore/qcborvalue.h:241
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

// /usr/include/qt/QtCore/qcborvalue.h:241
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

//  body block end

//  keep block begin

func init_unused_10352() {
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
