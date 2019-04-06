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
// extern C begin: 13
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
type QDoubleValidator struct {
	*QValidator
}
type QDoubleValidator_ITF interface {
	QValidator_ITF
	QDoubleValidator_PTR() *QDoubleValidator
}

func (ptr *QDoubleValidator) QDoubleValidator_PTR() *QDoubleValidator { return ptr }

func (this *QDoubleValidator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QValidator.GetCthis()
	}
}
func (this *QDoubleValidator) SetCthis(cthis unsafe.Pointer) {
	this.QValidator = NewQValidatorFromPointer(cthis)
}
func NewQDoubleValidatorFromPointer(cthis unsafe.Pointer) *QDoubleValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QDoubleValidator{bcthis0}
}
func (*QDoubleValidator) NewFromPointer(cthis unsafe.Pointer) *QDoubleValidator {
	return NewQDoubleValidatorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvalidator.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDoubleValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qvalidator.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleValidator(QObject *)

/*

 */
func (*QDoubleValidator) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QDoubleValidator {
	return NewQDoubleValidator(parent)
}
func NewQDoubleValidator(parent qtcore.QObject_ITF /*777 QObject **/) *QDoubleValidator {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDoubleValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleValidator(QObject *)

/*

 */
func (*QDoubleValidator) NewForInheritp() *QDoubleValidator {
	return NewQDoubleValidatorp()
}
func NewQDoubleValidatorp() *QDoubleValidator {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDoubleValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:136
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDoubleValidator(double, double, int, QObject *)

/*

 */
func (*QDoubleValidator) NewForInherit1(bottom float64, top float64, decimals int, parent qtcore.QObject_ITF /*777 QObject **/) *QDoubleValidator {
	return NewQDoubleValidator1(bottom, top, decimals, parent)
}
func NewQDoubleValidator1(bottom float64, top float64, decimals int, parent qtcore.QObject_ITF /*777 QObject **/) *QDoubleValidator {
	var convArg3 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg3 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EddiP7QObject", qtrt.FFI_TYPE_POINTER, bottom, top, decimals, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDoubleValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:136
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDoubleValidator(double, double, int, QObject *)

/*

 */
func (*QDoubleValidator) NewForInherit1p(bottom float64, top float64, decimals int) *QDoubleValidator {
	return NewQDoubleValidator1p(bottom, top, decimals)
}
func NewQDoubleValidator1p(bottom float64, top float64, decimals int) *QDoubleValidator {
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EddiP7QObject", qtrt.FFI_TYPE_POINTER, bottom, top, decimals, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDoubleValidator")
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:137
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDoubleValidator()

/*

 */
func DeleteQDoubleValidator(this *QDoubleValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const

/*
This virtual function returns Invalid if input is invalid according to this validator's rules, Intermediate if it is likely that a little more editing will make the input acceptable (e.g. the user types "4" into a widget which accepts integers between 10 and 99), and Acceptable if the input is valid.

The function can change both input and pos (the cursor position) if required.
*/
func (this *QDoubleValidator) Validate(arg0 string, arg1 int) int {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:146
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRange(double, double, int)

/*

 */
func (this *QDoubleValidator) SetRange(bottom float64, top float64, decimals int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator8setRangeEddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom, top, decimals)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:146
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRange(double, double, int)

/*

 */
func (this *QDoubleValidator) SetRangep(bottom float64, top float64) {
	// arg: 2, int=Int, =Invalid, , Invalid
	decimals := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator8setRangeEddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom, top, decimals)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottom(double)

/*

 */
func (this *QDoubleValidator) SetBottom(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator9setBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTop(double)

/*

 */
func (this *QDoubleValidator) SetTop(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator6setTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecimals(int)

/*

 */
func (this *QDoubleValidator) SetDecimals(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator11setDecimalsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotation(QDoubleValidator::Notation)

/*

 */
func (this *QDoubleValidator) SetNotation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator11setNotationENS_8NotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double bottom() const

/*

 */
func (this *QDoubleValidator) Bottom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator6bottomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double top() const

/*

 */
func (this *QDoubleValidator) Top() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator3topEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int decimals() const

/*

 */
func (this *QDoubleValidator) Decimals() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator8decimalsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:155
// index:0
// Public Visibility=Default Availability=Available
// [4] QDoubleValidator::Notation notation() const

/*

 */
func (this *QDoubleValidator) Notation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator8notationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bottomChanged(double)

/*

 */
func (this *QDoubleValidator) BottomChanged(bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator13bottomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topChanged(double)

/*

 */
func (this *QDoubleValidator) TopChanged(top float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator10topChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), top)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void decimalsChanged(int)

/*

 */
func (this *QDoubleValidator) DecimalsChanged(decimals int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator15decimalsChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), decimals)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notationChanged(QDoubleValidator::Notation)

/*

 */
func (this *QDoubleValidator) NotationChanged(notation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator15notationChangedENS_8NotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), notation)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QDoubleValidator__Notation = int

//
const QDoubleValidator__StandardNotation QDoubleValidator__Notation = 0

//
const QDoubleValidator__ScientificNotation QDoubleValidator__Notation = 1

func (this *QDoubleValidator) NotationItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDoubleValidator_NotationItemName(val int) string {
	var nilthis *QDoubleValidator
	return nilthis.NotationItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10973() {
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
