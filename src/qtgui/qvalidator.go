package qtgui

// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:62
// index:0
// Public
// void QValidator(class QObject *)
func NewQValidator(parent *qtcore.QObject /*777 QObject **/) *QValidator {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QValidatorC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QValidator) Locale() *qtcore.QLocale /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator6localeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:74
// index:0
// Public pure virtual
// QValidator::State validate(class QString &, int &)
func (this *QValidator) Validate(arg0 *qtcore.QString, arg1 int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QValidator8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
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

type QValidator__State = int

const QValidator__Invalid QValidator__State = 0
const QValidator__Intermediate QValidator__State = 1
const QValidator__Acceptable QValidator__State = 2

//  body block end
