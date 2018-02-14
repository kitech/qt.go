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
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QValidator struct {
	*qtcore.QObject
}
type QValidator_ITF interface {
	qtcore.QObject_ITF
	QValidator_PTR() *QValidator
}

func (ptr *QValidator) QValidator_PTR() *QValidator { return ptr }

func (this *QValidator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QValidator) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQValidatorFromPointer(cthis unsafe.Pointer) *QValidator {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QValidator{bcthis0}
}
func (*QValidator) NewFromPointer(cthis unsafe.Pointer) *QValidator {
	return NewQValidatorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvalidator.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QValidator(QObject *)
func NewQValidator(parent qtcore.QObject_ITF /*777 QObject **/) *QValidator {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QValidator()
func DeleteQValidator(this *QValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)
func (this *QValidator) SetLocale(locale qtcore.QLocale_ITF) {
	var convArg0 = locale.QLocale_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidator9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QValidator) Locale() *qtcore.QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QValidator) Validate(arg0 string, arg1 int) int {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &)
func (this *QValidator) Fixup(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed()
func (this *QValidator) Changed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidator7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QValidator__State = int

const QValidator__Invalid QValidator__State = 0
const QValidator__Intermediate QValidator__State = 1
const QValidator__Acceptable QValidator__State = 2

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
