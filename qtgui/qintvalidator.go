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
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QIntValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIntValidator(QObject *)
func NewQIntValidator(parent *qtcore.QObject /*777 QObject **/) *QIntValidator {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:97
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QIntValidator(int, int, QObject *)
func NewQIntValidator_1(bottom int, top int, parent *qtcore.QObject /*777 QObject **/) *QIntValidator {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorC2EiiP7QObject", qtrt.FFI_TYPE_POINTER, bottom, top, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIntValidator()
func DeleteQIntValidator(this *QIntValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QIntValidator) Validate(arg0 string, arg1 int) int {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &)
func (this *QIntValidator) Fixup(input string) {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottom(int)
func (this *QIntValidator) SetBottom(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator9setBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTop(int)
func (this *QIntValidator) SetTop(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator6setTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRange(int, int)
func (this *QIntValidator) SetRange(bottom int, top int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom, top)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom()
func (this *QIntValidator) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top()
func (this *QIntValidator) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bottomChanged(int)
func (this *QIntValidator) BottomChanged(bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator13bottomChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topChanged(int)
func (this *QIntValidator) TopChanged(top int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator10topChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), top)
	qtrt.ErrPrint(err, rv)
}

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
