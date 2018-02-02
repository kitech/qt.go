package qtgui

// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QRegularExpressionValidator struct {
	*QValidator
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
// [8] const QMetaObject * metaObject()
func (this *QRegularExpressionValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionValidator(QObject *)
func NewQRegularExpressionValidator(parent *qtcore.QObject /*777 QObject **/) *QRegularExpressionValidator {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:208
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionValidator(const QRegularExpression &, QObject *)
func NewQRegularExpressionValidator_1(re *qtcore.QRegularExpression, parent *qtcore.QObject /*777 QObject **/) *QRegularExpressionValidator {
	var convArg0 = re.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRegularExpressionValidator()
func DeleteQRegularExpressionValidator(this *QRegularExpressionValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:211
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QRegularExpressionValidator) Validate(input *qtcore.QString, pos int) int {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression regularExpression()
func (this *QRegularExpressionValidator) RegularExpression() *qtcore.QRegularExpression /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator17regularExpressionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRegularExpression(const QRegularExpression &)
func (this *QRegularExpressionValidator) SetRegularExpression(re *qtcore.QRegularExpression) {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void regularExpressionChanged(const QRegularExpression &)
func (this *QRegularExpressionValidator) RegularExpressionChanged(re *qtcore.QRegularExpression) {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidator24regularExpressionChangedERK18QRegularExpression", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
