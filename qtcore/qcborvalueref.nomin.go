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

// /usr/include/qt/QtCore/qcborvalue.h:390
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression(const QRegularExpression &) const

/*
Returns the regular expression value stored in this QCborValue, if it is of the regular expression pattern extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QRegularExpression.

See also isRegularExpression(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToRegularExpression(defaultValue QRegularExpression_ITF) *QRegularExpression /*123*/ {
	var convArg0 unsafe.Pointer
	if defaultValue != nil && defaultValue.QRegularExpression_PTR() != nil {
		convArg0 = defaultValue.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef19toRegularExpressionERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qcborvalue.h:390
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression(const QRegularExpression &) const

/*
Returns the regular expression value stored in this QCborValue, if it is of the regular expression pattern extended type. Otherwise, it returns defaultValue.

Note that this function performs no conversion from other types to QRegularExpression.

See also isRegularExpression(), isTag(), and taggedValue().
*/
func (this *QCborValueRef) ToRegularExpressionp() *QRegularExpression /*123*/ {
	// arg: 0, const QRegularExpression &=LValueReference, QRegularExpression=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QCborValueRef19toRegularExpressionERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10354() {
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
