//  header block begin
// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>
package qtgui

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
type QDoubleValidator struct {
	*QValidator
}

func (this *QDoubleValidator) GetCthis() unsafe.Pointer {
	return this.QValidator.GetCthis()
}

// /usr/include/qt/QtGui/qvalidator.h:126
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDoubleValidator) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDoubleValidator10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:133
// index:0
// void QDoubleValidator(class QObject *)
func NewQDoubleValidator(parent unsafe.Pointer) *QDoubleValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(cthis)
	return gothis
}
func NewQDoubleValidatorFromPointer(cthis unsafe.Pointer) *QDoubleValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QDoubleValidator{bcthis0}
}

// /usr/include/qt/QtGui/qvalidator.h:134
// index:1
// void QDoubleValidator(double, double, int, class QObject *)
func NewQDoubleValidator_1(bottom float64, top float64, decimals int, parent unsafe.Pointer) *QDoubleValidator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidatorC2EddiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &bottom, &top, &decimals, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDoubleValidatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvalidator.h:135
// index:0
// virtual
// void ~QDoubleValidator()
func DeleteQDoubleValidator(*QDoubleValidator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:142
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QDoubleValidator) Validate(arg0 unsafe.Pointer, arg1 int) {
	// 0: (, QString & arg0, int & arg1), (arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDoubleValidator8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:144
// index:0
// virtual
// void setRange(double, double, int)
func (this *QDoubleValidator) SetRange(bottom float64, top float64, decimals int) {
	// 0: (, bottom double, top double, decimals int), (&bottom, &top, &decimals)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator8setRangeEddi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bottom, &top, &decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:145
// index:0
// void setBottom(double)
func (this *QDoubleValidator) SetBottom(arg0 float64) {
	// 0: (, double arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator9setBottomEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:146
// index:0
// void setTop(double)
func (this *QDoubleValidator) SetTop(arg0 float64) {
	// 0: (, double arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator6setTopEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:147
// index:0
// void setDecimals(int)
func (this *QDoubleValidator) SetDecimals(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator11setDecimalsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:148
// index:0
// void setNotation(enum QDoubleValidator::Notation)
func (this *QDoubleValidator) SetNotation(arg0 int) {
	// 0: (, QDoubleValidator::Notation arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator11setNotationENS_8NotationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:150
// index:0
// inline
// double bottom()
func (this *QDoubleValidator) Bottom() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDoubleValidator6bottomEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:151
// index:0
// inline
// double top()
func (this *QDoubleValidator) Top() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDoubleValidator3topEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:152
// index:0
// inline
// int decimals()
func (this *QDoubleValidator) Decimals() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDoubleValidator8decimalsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:153
// index:0
// QDoubleValidator::Notation notation()
func (this *QDoubleValidator) Notation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDoubleValidator8notationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:156
// index:0
// void bottomChanged(double)
func (this *QDoubleValidator) BottomChanged(bottom float64) {
	// 0: (, bottom double), (&bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator13bottomChangedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:157
// index:0
// void topChanged(double)
func (this *QDoubleValidator) TopChanged(top float64) {
	// 0: (, top double), (&top)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator10topChangedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:158
// index:0
// void decimalsChanged(int)
func (this *QDoubleValidator) DecimalsChanged(decimals int) {
	// 0: (, decimals int), (&decimals)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator15decimalsChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvalidator.h:159
// index:0
// void notationChanged(class QDoubleValidator::Notation)
func (this *QDoubleValidator) NotationChanged(notation int) {
	// 0: (, notation QDoubleValidator::Notation), (&notation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDoubleValidator15notationChangedENS_8NotationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &notation)
	gopp.ErrPrint(err, rv)
}

//  body block end
