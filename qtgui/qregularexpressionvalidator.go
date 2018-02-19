package qtgui

// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QRegularExpressionValidator struct {
	*QValidator
}
type QRegularExpressionValidator_ITF interface {
	QValidator_ITF
	QRegularExpressionValidator_PTR() *QRegularExpressionValidator
}

func (ptr *QRegularExpressionValidator) QRegularExpressionValidator_PTR() *QRegularExpressionValidator {
	return ptr
}

func (this *QRegularExpressionValidator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QValidator.GetCthis()
	}
}
func (this *QRegularExpressionValidator) SetCthis(cthis unsafe.Pointer) {
	this.QValidator = NewQValidatorFromPointer(cthis)
}
func NewQRegularExpressionValidatorFromPointer(cthis unsafe.Pointer) *QRegularExpressionValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QRegularExpressionValidator{bcthis0}
}
func (*QRegularExpressionValidator) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionValidator {
	return NewQRegularExpressionValidatorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvalidator.h:203
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QRegularExpressionValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionValidator(QObject *)
func NewQRegularExpressionValidator(parent qtcore.QObject_ITF /*777 QObject **/) *QRegularExpressionValidator {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegularExpressionValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionValidator(QObject *)
func NewQRegularExpressionValidator__() *QRegularExpressionValidator {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegularExpressionValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:208
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionValidator(const QRegularExpression &, QObject *)
func NewQRegularExpressionValidator_1(re qtcore.QRegularExpression_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QRegularExpressionValidator {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegularExpressionValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:208
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionValidator(const QRegularExpression &, QObject *)
func NewQRegularExpressionValidator_1_(re qtcore.QRegularExpression_ITF) *QRegularExpressionValidator {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegularExpressionValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRegularExpressionValidator()
func DeleteQRegularExpressionValidator(this *QRegularExpressionValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:211
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const
func (this *QRegularExpressionValidator) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression regularExpression() const
func (this *QRegularExpressionValidator) RegularExpression() *qtcore.QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator17regularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRegularExpression(const QRegularExpression &)
func (this *QRegularExpressionValidator) SetRegularExpression(re qtcore.QRegularExpression_ITF) {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void regularExpressionChanged(const QRegularExpression &)
func (this *QRegularExpressionValidator) RegularExpressionChanged(re qtcore.QRegularExpression_ITF) {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QRegularExpressionValidator24regularExpressionChangedERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
