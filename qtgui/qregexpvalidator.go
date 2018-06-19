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
// extern C begin: 18
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
type QRegExpValidator struct {
	*QValidator
}
type QRegExpValidator_ITF interface {
	QValidator_ITF
	QRegExpValidator_PTR() *QRegExpValidator
}

func (ptr *QRegExpValidator) QRegExpValidator_PTR() *QRegExpValidator { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRegExpValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QRegExpValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegExpValidator(QObject *)

/*

 */
func NewQRegExpValidator(parent qtcore.QObject_ITF /*777 QObject **/) *QRegExpValidator {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegExpValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegExpValidator(QObject *)

/*

 */
func NewQRegExpValidator__() *QRegExpValidator {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegExpValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:178
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegExpValidator(const QRegExp &, QObject *)

/*

 */
func NewQRegExpValidator_1(rx qtcore.QRegExp_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QRegExpValidator {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegExpValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:178
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegExpValidator(const QRegExp &, QObject *)

/*

 */
func NewQRegExpValidator_1_(rx qtcore.QRegExp_ITF) *QRegExpValidator {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidatorC2ERK7QRegExpP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRegExpValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:179
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRegExpValidator()

/*

 */
func DeleteQRegExpValidator(this *QRegExpValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:181
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const

/*
This virtual function returns Invalid if input is invalid according to this validator's rules, Intermediate if it is likely that a little more editing will make the input acceptable (e.g. the user types "4" into a widget which accepts integers between 10 and 99), and Acceptable if the input is valid.

The function can change both input and pos (the cursor position) if required.
*/
func (this *QRegExpValidator) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QRegExpValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRegExp(const QRegExp &)

/*

 */
func (this *QRegExpValidator) SetRegExp(rx qtcore.QRegExp_ITF) {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidator9setRegExpERK7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QRegExp & regExp() const

/*

 */
func (this *QRegExpValidator) RegExp() *qtcore.QRegExp {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QRegExpValidator6regExpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRegExp)
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void regExpChanged(const QRegExp &)

/*

 */
func (this *QRegExpValidator) RegExpChanged(regExp qtcore.QRegExp_ITF) {
	var convArg0 unsafe.Pointer
	if regExp != nil && regExp.QRegExp_PTR() != nil {
		convArg0 = regExp.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRegExpValidator13regExpChangedERK7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
