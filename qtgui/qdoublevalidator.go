package qtgui

// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QDoubleValidator struct {
	*QValidator
}

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

// /usr/include/qt/QtGui/qvalidator.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDoubleValidator) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qvalidator.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleValidator(QObject *)
func NewQDoubleValidator(parent *qtcore.QObject /*777 QObject **/) *QDoubleValidator {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:134
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDoubleValidator(double, double, int, QObject *)
func NewQDoubleValidator_1(bottom float64, top float64, decimals int, parent *qtcore.QObject /*777 QObject **/) *QDoubleValidator {
	var convArg3 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EddiP7QObject", qtrt.FFI_TYPE_POINTER, bottom, top, decimals, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:135
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDoubleValidator()
func DeleteQDoubleValidator(this *QDoubleValidator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qvalidator.h:142
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QDoubleValidator) Validate(arg0 *qtcore.QString, arg1 int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRange(double, double, int)
func (this *QDoubleValidator) SetRange(bottom float64, top float64, decimals int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator8setRangeEddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom, top, decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottom(double)
func (this *QDoubleValidator) SetBottom(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator9setBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTop(double)
func (this *QDoubleValidator) SetTop(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator6setTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecimals(int)
func (this *QDoubleValidator) SetDecimals(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator11setDecimalsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotation(enum QDoubleValidator::Notation)
func (this *QDoubleValidator) SetNotation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator11setNotationENS_8NotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double bottom()
func (this *QDoubleValidator) Bottom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator6bottomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double top()
func (this *QDoubleValidator) Top() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator3topEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int decimals()
func (this *QDoubleValidator) Decimals() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator8decimalsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qvalidator.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QDoubleValidator::Notation notation()
func (this *QDoubleValidator) Notation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDoubleValidator8notationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qvalidator.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bottomChanged(double)
func (this *QDoubleValidator) BottomChanged(bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator13bottomChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topChanged(double)
func (this *QDoubleValidator) TopChanged(top float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator10topChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void decimalsChanged(int)
func (this *QDoubleValidator) DecimalsChanged(decimals int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator15decimalsChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notationChanged(QDoubleValidator::Notation)
func (this *QDoubleValidator) NotationChanged(notation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDoubleValidator15notationChangedENS_8NotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), notation)
	gopp.ErrPrint(err, rv)
}

type QDoubleValidator__Notation = int

const QDoubleValidator__StandardNotation QDoubleValidator__Notation = 0
const QDoubleValidator__ScientificNotation QDoubleValidator__Notation = 1

//  body block end
