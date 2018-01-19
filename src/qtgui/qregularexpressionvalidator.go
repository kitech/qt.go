//  header block begin
// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>
package qtgui

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
type QRegularExpressionValidator struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qvalidator.h:203
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QRegularExpressionValidator) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:207
// index:0
// void QRegularExpressionValidator(class QObject *)
func NewQRegularExpressionValidator(parent unsafe.Pointer) *QRegularExpressionValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QRegularExpressionValidator{cthis}
}

// /usr/include/qt/QtGui/qvalidator.h:208
// index:1
// void QRegularExpressionValidator(const class QRegularExpression &, class QObject *)
func NewQRegularExpressionValidator_1(re unsafe.Pointer, parent unsafe.Pointer) *QRegularExpressionValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorC2ERK18QRegularExpressionP7QObject", ffiqt.FFI_TYPE_VOID, cthis, re, parent)
	gopp.ErrPrint(err, rv)
	return &QRegularExpressionValidator{cthis}
}

// /usr/include/qt/QtGui/qvalidator.h:209
// index:0
// virtual
// void ~QRegularExpressionValidator()
func DeleteQRegularExpressionValidator(*QRegularExpressionValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:211
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QRegularExpressionValidator) Validate(input unsafe.Pointer, pos int) {
	// 0: (, QString & input, int & pos), (input, &pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.cthis, input, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:213
// index:0
// QRegularExpression regularExpression()
func (this *QRegularExpressionValidator) RegularExpression() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QRegularExpressionValidator17regularExpressionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:216
// index:0
// void setRegularExpression(const class QRegularExpression &)
func (this *QRegularExpressionValidator) SetRegularExpression(re unsafe.Pointer) {
	// 0: (, const QRegularExpression & re), (re)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidator20setRegularExpressionERK18QRegularExpression", ffiqt.FFI_TYPE_VOID, this.cthis, re)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:219
// index:0
// void regularExpressionChanged(const class QRegularExpression &)
func (this *QRegularExpressionValidator) RegularExpressionChanged(re unsafe.Pointer) {
	// 0: (, const QRegularExpression & re), (re)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QRegularExpressionValidator24regularExpressionChangedERK18QRegularExpression", ffiqt.FFI_TYPE_VOID, this.cthis, re)
	gopp.ErrPrint(err, rv)
}

//  body block end
