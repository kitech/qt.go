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
func NewQValidatorFromPointer(cthis unsafe.Pointer) *QValidator {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QValidator{bcthis0}
}

// /usr/include/qt/QtGui/qvalidator.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QValidator) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvalidator.h:62
// index:0
// Public
// void QValidator(class QObject *)
func NewQValidator(parent unsafe.Pointer) *QValidator {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:63
// index:0
// Public virtual
// void ~QValidator()
func DeleteQValidator(*QValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:71
// index:0
// Public
// void setLocale(const class QLocale &)
func (this *QValidator) SetLocale(locale *qtcore.QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidator9setLocaleERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:72
// index:0
// Public
// QLocale locale()
func (this *QValidator) Locale() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator6localeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvalidator.h:74
// index:0
// Public pure virtual
// QValidator::State validate(class QString &, int &)
func (this *QValidator) Validate(arg0 *qtcore.QString, arg1 int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvalidator.h:75
// index:0
// Public virtual
// void fixup(class QString &)
func (this *QValidator) Fixup(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:78
// index:0
// Public
// void changed()
func (this *QValidator) Changed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidator7changedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
