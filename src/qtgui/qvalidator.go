//  header block begin
// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QValidator struct {
	*qtcore.QObject
}

func (this *QValidator) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qvalidator.h:60
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QValidator) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:62
// index:0
// void QValidator(class QObject *)
func NewQValidator(parent unsafe.Pointer) *QValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(cthis)
	return gothis
}
func NewQValidatorFromPointer(cthis unsafe.Pointer) *QValidator {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QValidator{bcthis0}
}

// /usr/include/qt/QtGui/qvalidator.h:81
// index:1
// void QValidator(class QObjectPrivate &, class QObject *)
func NewQValidator_1(d unsafe.Pointer, parent unsafe.Pointer) *QValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorC1ER14QObjectPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, d, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:82
// index:2
// void QValidator(class QValidatorPrivate &, class QObject *)
func NewQValidator_2(d unsafe.Pointer, parent unsafe.Pointer) *QValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorC1ER17QValidatorPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, d, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:63
// index:0
// virtual
// void ~QValidator()
func DeleteQValidator(*QValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:71
// index:0
// void setLocale(const class QLocale &)
func (this *QValidator) SetLocale(locale unsafe.Pointer) {
	// 0: (, locale const QLocale &), (locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidator9setLocaleERK7QLocale", ffiqt.FFI_TYPE_VOID, this.GetCthis(), locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:72
// index:0
// QLocale locale()
func (this *QValidator) Locale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator6localeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:74
// index:0
// pure virtual
// QValidator::State validate(class QString &, int &)
func (this *QValidator) Validate(arg0 unsafe.Pointer, arg1 int) {
	// 0: (, QString & arg0, int & arg1), (arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:75
// index:0
// virtual
// void fixup(class QString &)
func (this *QValidator) Fixup(arg0 unsafe.Pointer) {
	// 0: (, QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator5fixupER7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:78
// index:0
// void changed()
func (this *QValidator) Changed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidator7changedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
