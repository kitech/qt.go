//  header block begin
// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QIntValidator struct {
	*QValidator
}

func (this *QIntValidator) GetCthis() unsafe.Pointer {
	return this.QValidator.GetCthis()
}

// /usr/include/qt/QtGui/qvalidator.h:91
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QIntValidator) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:96
// index:0
// void QIntValidator(class QObject *)
func NewQIntValidator(parent unsafe.Pointer) *QIntValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(cthis)
	return gothis
}
func NewQIntValidatorFromPointer(cthis unsafe.Pointer) *QIntValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QIntValidator{bcthis0}
}

// /usr/include/qt/QtGui/qvalidator.h:97
// index:1
// void QIntValidator(int, int, class QObject *)
func NewQIntValidator_1(bottom int, top int, parent unsafe.Pointer) *QIntValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidatorC2EiiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &bottom, &top, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:98
// index:0
// virtual
// void ~QIntValidator()
func DeleteQIntValidator(*QIntValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:100
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QIntValidator) Validate(arg0 unsafe.Pointer, arg1 int) {
	// 0: (, QString & arg0, int & arg1), (arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:101
// index:0
// virtual
// void fixup(class QString &)
func (this *QIntValidator) Fixup(input unsafe.Pointer) {
	// 0: (, input QString &), (input)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator5fixupER7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), input)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:103
// index:0
// void setBottom(int)
func (this *QIntValidator) SetBottom(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator9setBottomEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:104
// index:0
// void setTop(int)
func (this *QIntValidator) SetTop(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator6setTopEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:105
// index:0
// virtual
// void setRange(int, int)
func (this *QIntValidator) SetRange(bottom int, top int) {
	// 0: (, bottom int, top int), (&bottom, &top)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator8setRangeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bottom, &top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:107
// index:0
// inline
// int bottom()
func (this *QIntValidator) Bottom() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator6bottomEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:108
// index:0
// inline
// int top()
func (this *QIntValidator) Top() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator3topEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:110
// index:0
// void bottomChanged(int)
func (this *QIntValidator) BottomChanged(bottom int) {
	// 0: (, bottom int), (&bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator13bottomChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:111
// index:0
// void topChanged(int)
func (this *QIntValidator) TopChanged(top int) {
	// 0: (, top int), (&top)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator10topChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &top)
	gopp.ErrPrint(err, rv)
}

//  body block end
