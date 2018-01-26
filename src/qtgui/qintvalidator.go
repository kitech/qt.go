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
	if this == nil {
		return nil
	} else {
		return this.QValidator.GetCthis()
	}
}
func (this *QIntValidator) SetCthis(cthis unsafe.Pointer) {
	this.QValidator = NewQValidatorFromPointer(cthis)
}
func NewQIntValidatorFromPointer(cthis unsafe.Pointer) *QIntValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QIntValidator{bcthis0}
}
func (*QIntValidator) NewFromPointer(cthis unsafe.Pointer) *QIntValidator {
	return NewQIntValidatorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvalidator.h:91
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QIntValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:96
// index:0
// Public
// void QIntValidator(class QObject *)
func NewQIntValidator(parent *qtcore.QObject /*777 QObject **/) *QIntValidator {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:97
// index:1
// Public
// void QIntValidator(int, int, class QObject *)
func NewQIntValidator_1(bottom int, top int, parent *qtcore.QObject /*777 QObject **/) *QIntValidator {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidatorC2EiiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, bottom, top, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:98
// index:0
// Public virtual
// void ~QIntValidator()
func DeleteQIntValidator(*QIntValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:100
// index:0
// Public virtual
// QValidator::State validate(class QString &, int &)
func (this *QIntValidator) Validate(arg0 *qtcore.QString, arg1 int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:101
// index:0
// Public virtual
// void fixup(class QString &)
func (this *QIntValidator) Fixup(input *qtcore.QString) {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:103
// index:0
// Public
// void setBottom(int)
func (this *QIntValidator) SetBottom(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator9setBottomEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:104
// index:0
// Public
// void setTop(int)
func (this *QIntValidator) SetTop(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator6setTopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:105
// index:0
// Public virtual
// void setRange(int, int)
func (this *QIntValidator) SetRange(bottom int, top int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator8setRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bottom, top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:107
// index:0
// Public inline
// int bottom()
func (this *QIntValidator) Bottom() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator6bottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qvalidator.h:108
// index:0
// Public inline
// int top()
func (this *QIntValidator) Top() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QIntValidator3topEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qvalidator.h:110
// index:0
// Public
// void bottomChanged(int)
func (this *QIntValidator) BottomChanged(bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator13bottomChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:111
// index:0
// Public
// void topChanged(int)
func (this *QIntValidator) TopChanged(top int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QIntValidator10topChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), top)
	gopp.ErrPrint(err, rv)
}

//  body block end
