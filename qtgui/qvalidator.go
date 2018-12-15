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
// extern C begin: 22
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
type QValidator struct {
	*qtcore.QObject
}
type QValidator_ITF interface {
	qtcore.QObject_ITF
	QValidator_PTR() *QValidator
}

func (ptr *QValidator) QValidator_PTR() *QValidator { return ptr }

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

// /usr/include/qt/QtGui/qvalidator.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QValidator(QObject *)

/*
Sets up the validator. The parent parameter is passed on to the QObject constructor.
*/
func (*QValidator) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QValidator {
	return NewQValidator(parent)
}
func NewQValidator(parent qtcore.QObject_ITF /*777 QObject **/) *QValidator {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QValidator(QObject *)

/*
Sets up the validator. The parent parameter is passed on to the QObject constructor.
*/
func (*QValidator) NewForInheritp() *QValidator {
	return NewQValidatorp()
}
func NewQValidatorp() *QValidator {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QValidator()

/*

 */
func DeleteQValidator(this *QValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)

/*
Sets the locale that will be used for the validator. Unless setLocale has been called, the validator will use the default locale set with QLocale::setDefault(). If a default locale has not been set, it is the operating system's locale.

See also locale() and QLocale::setDefault().
*/
func (this *QValidator) SetLocale(locale qtcore.QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidator9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale() const

/*
Returns the locale for the validator. The locale is by default initialized to the same as QLocale().

See also setLocale() and QLocale::QLocale().
*/
func (this *QValidator) Locale() *qtcore.QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const

/*
This virtual function returns Invalid if input is invalid according to this validator's rules, Intermediate if it is likely that a little more editing will make the input acceptable (e.g. the user types "4" into a widget which accepts integers between 10 and 99), and Acceptable if the input is valid.

The function can change both input and pos (the cursor position) if required.
*/
func (this *QValidator) Validate(arg0 string, arg1 int) int {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &) const

/*
This function attempts to change input to be valid according to this validator's rules. It need not result in a valid string: callers of this function must re-test afterwards; the default does nothing.

Reimplementations of this function can change input even if they do not produce a valid string. For example, an ISBN validator might want to delete every character except digits and "-", even if the result is still not a valid ISBN; a surname validator might want to remove whitespace from the start and end of the string, even if the resulting string is not in the list of accepted surnames.
*/
func (this *QValidator) Fixup(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QValidator5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed()

/*
This signal is emitted when any property that may affect the validity of a string has changed.
*/
func (this *QValidator) Changed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QValidator7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
This enum type defines the states in which a validated string can exist.


*/
type QValidator__State = int

// The string is clearly invalid.
const QValidator__Invalid QValidator__State = 0

// The string is a plausible intermediate value.
const QValidator__Intermediate QValidator__State = 1

// The string is acceptable as a final result; i.e. it is valid.
const QValidator__Acceptable QValidator__State = 2

func (this *QValidator) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QValidator_StateItemName(val int) string {
	var nilthis *QValidator
	return nilthis.StateItemName(val)
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
