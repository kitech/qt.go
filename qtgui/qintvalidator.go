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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QIntValidator struct {
	*QValidator
}
type QIntValidator_ITF interface {
	QValidator_ITF
	QIntValidator_PTR() *QIntValidator
}

func (ptr *QIntValidator) QIntValidator_PTR() *QIntValidator { return ptr }

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

// /usr/include/qt/QtGui/qvalidator.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QIntValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIntValidator(QObject *)

/*

 */
func (*QIntValidator) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QIntValidator {
	return NewQIntValidator(parent)
}
func NewQIntValidator(parent qtcore.QObject_ITF /*777 QObject **/) *QIntValidator {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIntValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIntValidator(QObject *)

/*

 */
func (*QIntValidator) NewForInheritp() *QIntValidator {
	return NewQIntValidatorp()
}
func NewQIntValidatorp() *QIntValidator {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIntValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:99
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QIntValidator(int, int, QObject *)

/*

 */
func (*QIntValidator) NewForInherit1(bottom int, top int, parent qtcore.QObject_ITF /*777 QObject **/) *QIntValidator {
	return NewQIntValidator1(bottom, top, parent)
}
func NewQIntValidator1(bottom int, top int, parent qtcore.QObject_ITF /*777 QObject **/) *QIntValidator {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorC2EiiP7QObject", qtrt.FFI_TYPE_POINTER, bottom, top, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIntValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:99
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QIntValidator(int, int, QObject *)

/*

 */
func (*QIntValidator) NewForInherit1p(bottom int, top int) *QIntValidator {
	return NewQIntValidator1p(bottom, top)
}
func NewQIntValidator1p(bottom int, top int) *QIntValidator {
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorC2EiiP7QObject", qtrt.FFI_TYPE_POINTER, bottom, top, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIntValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIntValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIntValidator()

/*

 */
func DeleteQIntValidator(this *QIntValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const

/*
This virtual function returns Invalid if input is invalid according to this validator's rules, Intermediate if it is likely that a little more editing will make the input acceptable (e.g. the user types "4" into a widget which accepts integers between 10 and 99), and Acceptable if the input is valid.

The function can change both input and pos (the cursor position) if required.
*/
func (this *QIntValidator) Validate(arg0 string, arg1 int) int {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &) const

/*
This function attempts to change input to be valid according to this validator's rules. It need not result in a valid string: callers of this function must re-test afterwards; the default does nothing.

Reimplementations of this function can change input even if they do not produce a valid string. For example, an ISBN validator might want to delete every character except digits and "-", even if the result is still not a valid ISBN; a surname validator might want to remove whitespace from the start and end of the string, even if the resulting string is not in the list of accepted surnames.
*/
func (this *QIntValidator) Fixup(input string) {
	var tmpArg0 = qtcore.NewQString5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottom(int)

/*

 */
func (this *QIntValidator) SetBottom(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator9setBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTop(int)

/*

 */
func (this *QIntValidator) SetTop(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator6setTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRange(int, int)

/*

 */
func (this *QIntValidator) SetRange(bottom int, top int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom, top)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom() const

/*

 */
func (this *QIntValidator) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top() const

/*

 */
func (this *QIntValidator) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QIntValidator3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bottomChanged(int)

/*

 */
func (this *QIntValidator) BottomChanged(bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator13bottomChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topChanged(int)

/*

 */
func (this *QIntValidator) TopChanged(top int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QIntValidator10topChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), top)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10971() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
