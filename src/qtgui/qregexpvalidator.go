package qtgui

// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QRegExpValidator struct {
	*QValidator
}

func (this *QRegExpValidator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QValidator.GetCthis()
	}
}
func (this *QRegExpValidator) SetCthis(cthis unsafe.Pointer) {
	this.QValidator = NewQValidatorFromPointer(cthis)
}
func NewQRegExpValidatorFromPointer(cthis unsafe.Pointer) *QRegExpValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QRegExpValidator{bcthis0}
}
func (*QRegExpValidator) NewFromPointer(cthis unsafe.Pointer) *QRegExpValidator {
	return NewQRegExpValidatorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvalidator.h:173
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QRegExpValidator) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QRegExpValidator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:177
// index:0
// Public
// void QRegExpValidator(class QObject *)
func NewQRegExpValidator(parent *qtcore.QObject /*444 QObject **/) *QRegExpValidator {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:178
// index:1
// Public
// void QRegExpValidator(const class QRegExp &, class QObject *)
func NewQRegExpValidator_1(rx *qtcore.QRegExp, parent *qtcore.QObject /*444 QObject **/) *QRegExpValidator {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = rx.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:179
// index:0
// Public virtual
// void ~QRegExpValidator()
func DeleteQRegExpValidator(*QRegExpValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:181
// index:0
// Public virtual
// QValidator::State validate(class QString &, int &)
func (this *QRegExpValidator) Validate(input *qtcore.QString, pos int) int {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QRegExpValidator8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:183
// index:0
// Public
// void setRegExp(const class QRegExp &)
func (this *QRegExpValidator) SetRegExp(rx *qtcore.QRegExp) {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidator9setRegExpERK7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:184
// index:0
// Public inline
// const QRegExp & regExp()
func (this *QRegExpValidator) RegExp() *qtcore.QRegExp {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QRegExpValidator6regExpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:187
// index:0
// Public
// void regExpChanged(const class QRegExp &)
func (this *QRegExpValidator) RegExpChanged(regExp *qtcore.QRegExp) {
	var convArg0 = regExp.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidator13regExpChangedERK7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
