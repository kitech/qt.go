//  header block begin
// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>
package qtgui

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
	return this.QValidator.GetCthis()
}

// /usr/include/qt/QtGui/qvalidator.h:173
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QRegExpValidator) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QRegExpValidator10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:177
// index:0
// void QRegExpValidator(class QObject *)
func NewQRegExpValidator(parent unsafe.Pointer) *QRegExpValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(cthis)
	return gothis
}
func NewQRegExpValidatorFromPointer(cthis unsafe.Pointer) *QRegExpValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QRegExpValidator{bcthis0}
}

// /usr/include/qt/QtGui/qvalidator.h:178
// index:1
// void QRegExpValidator(const class QRegExp &, class QObject *)
func NewQRegExpValidator_1(rx unsafe.Pointer, parent unsafe.Pointer) *QRegExpValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject", ffiqt.FFI_TYPE_VOID, cthis, rx, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:179
// index:0
// virtual
// void ~QRegExpValidator()
func DeleteQRegExpValidator(*QRegExpValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:181
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QRegExpValidator) Validate(input unsafe.Pointer, pos int) {
	// 0: (, input QString &, pos int &), (input, &pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QRegExpValidator8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), input, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:183
// index:0
// void setRegExp(const class QRegExp &)
func (this *QRegExpValidator) SetRegExp(rx unsafe.Pointer) {
	// 0: (, rx const QRegExp &), (rx)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidator9setRegExpERK7QRegExp", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:184
// index:0
// inline
// const QRegExp & regExp()
func (this *QRegExpValidator) RegExp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QRegExpValidator6regExpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:187
// index:0
// void regExpChanged(const class QRegExp &)
func (this *QRegExpValidator) RegExpChanged(regExp unsafe.Pointer) {
	// 0: (, regExp const QRegExp &), (regExp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRegExpValidator13regExpChangedERK7QRegExp", ffiqt.FFI_TYPE_VOID, this.GetCthis(), regExp)
	gopp.ErrPrint(err, rv)
}

//  body block end
