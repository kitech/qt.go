//  header block begin
// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDoubleSpinBox struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qspinbox.h:115
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDoubleSpinBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:126
// index:0
// void QDoubleSpinBox(class QWidget *)
func NewQDoubleSpinBox(parent unsafe.Pointer) *QDoubleSpinBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QDoubleSpinBox{cthis}
}

// /usr/include/qt/QtWidgets/qspinbox.h:127
// index:0
// virtual
// void ~QDoubleSpinBox()
func DeleteQDoubleSpinBox(*QDoubleSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:129
// index:0
// double value()
func (this *QDoubleSpinBox) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:131
// index:0
// QString prefix()
func (this *QDoubleSpinBox) Prefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox6prefixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:132
// index:0
// void setPrefix(const class QString &)
func (this *QDoubleSpinBox) SetPrefix(prefix unsafe.Pointer) {
	// 0: (, const QString & prefix), (prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox9setPrefixERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:134
// index:0
// QString suffix()
func (this *QDoubleSpinBox) Suffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox6suffixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:135
// index:0
// void setSuffix(const class QString &)
func (this *QDoubleSpinBox) SetSuffix(suffix unsafe.Pointer) {
	// 0: (, const QString & suffix), (suffix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox9setSuffixERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, suffix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:137
// index:0
// QString cleanText()
func (this *QDoubleSpinBox) CleanText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox9cleanTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:139
// index:0
// double singleStep()
func (this *QDoubleSpinBox) SingleStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox10singleStepEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:140
// index:0
// void setSingleStep(double)
func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	// 0: (, double val), (&val)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox13setSingleStepEd", ffiqt.FFI_TYPE_VOID, this.cthis, &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:142
// index:0
// double minimum()
func (this *QDoubleSpinBox) Minimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox7minimumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:143
// index:0
// void setMinimum(double)
func (this *QDoubleSpinBox) SetMinimum(min float64) {
	// 0: (, double min), (&min)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox10setMinimumEd", ffiqt.FFI_TYPE_VOID, this.cthis, &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:145
// index:0
// double maximum()
func (this *QDoubleSpinBox) Maximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox7maximumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:146
// index:0
// void setMaximum(double)
func (this *QDoubleSpinBox) SetMaximum(max float64) {
	// 0: (, double max), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox10setMaximumEd", ffiqt.FFI_TYPE_VOID, this.cthis, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:148
// index:0
// void setRange(double, double)
func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	// 0: (, double min, double max), (&min, &max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox8setRangeEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:150
// index:0
// int decimals()
func (this *QDoubleSpinBox) Decimals() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox8decimalsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:151
// index:0
// void setDecimals(int)
func (this *QDoubleSpinBox) SetDecimals(prec int) {
	// 0: (, int prec), (&prec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox11setDecimalsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &prec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:153
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QDoubleSpinBox) Validate(input unsafe.Pointer, pos int) {
	// 0: (, QString & input, int & pos), (input, &pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.cthis, input, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:154
// index:0
// virtual
// double valueFromText(const class QString &)
func (this *QDoubleSpinBox) ValueFromText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox13valueFromTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:155
// index:0
// virtual
// QString textFromValue(double)
func (this *QDoubleSpinBox) TextFromValue(val float64) {
	// 0: (, double val), (&val)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox13textFromValueEd", ffiqt.FFI_TYPE_VOID, this.cthis, &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:156
// index:0
// virtual
// void fixup(class QString &)
func (this *QDoubleSpinBox) Fixup(str unsafe.Pointer) {
	// 0: (, QString & str), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDoubleSpinBox5fixupER7QString", ffiqt.FFI_TYPE_VOID, this.cthis, str)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:159
// index:0
// void setValue(double)
func (this *QDoubleSpinBox) SetValue(val float64) {
	// 0: (, double val), (&val)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox8setValueEd", ffiqt.FFI_TYPE_VOID, this.cthis, &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:162
// index:0
// void valueChanged(double)
func (this *QDoubleSpinBox) ValueChanged(arg0 float64) {
	// 0: (, double arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox12valueChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:163
// index:1
// void valueChanged(const class QString &)
func (this *QDoubleSpinBox) ValueChanged_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDoubleSpinBox12valueChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
