//  header block begin
// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QSpinBox struct {
	*QAbstractSpinBox
}

func (this *QSpinBox) GetCthis() unsafe.Pointer {
	return this.QAbstractSpinBox.GetCthis()
}

// /usr/include/qt/QtWidgets/qspinbox.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSpinBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:65
// index:0
// void QSpinBox(class QWidget *)
func NewQSpinBox(parent unsafe.Pointer) *QSpinBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSpinBoxFromPointer(cthis)
	return gothis
}
func NewQSpinBoxFromPointer(cthis unsafe.Pointer) *QSpinBox {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QSpinBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// virtual
// void ~QSpinBox()
func DeleteQSpinBox(*QSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:68
// index:0
// int value()
func (this *QSpinBox) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox5valueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:70
// index:0
// QString prefix()
func (this *QSpinBox) Prefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox6prefixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:71
// index:0
// void setPrefix(const class QString &)
func (this *QSpinBox) SetPrefix(prefix unsafe.Pointer) {
	// 0: (, prefix const QString &), (prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox9setPrefixERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:73
// index:0
// QString suffix()
func (this *QSpinBox) Suffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox6suffixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:74
// index:0
// void setSuffix(const class QString &)
func (this *QSpinBox) SetSuffix(suffix unsafe.Pointer) {
	// 0: (, suffix const QString &), (suffix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox9setSuffixERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), suffix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:76
// index:0
// QString cleanText()
func (this *QSpinBox) CleanText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox9cleanTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:78
// index:0
// int singleStep()
func (this *QSpinBox) SingleStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox10singleStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:79
// index:0
// void setSingleStep(int)
func (this *QSpinBox) SetSingleStep(val int) {
	// 0: (, val int), (&val)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox13setSingleStepEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:81
// index:0
// int minimum()
func (this *QSpinBox) Minimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox7minimumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:82
// index:0
// void setMinimum(int)
func (this *QSpinBox) SetMinimum(min int) {
	// 0: (, min int), (&min)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox10setMinimumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:84
// index:0
// int maximum()
func (this *QSpinBox) Maximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox7maximumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:85
// index:0
// void setMaximum(int)
func (this *QSpinBox) SetMaximum(max int) {
	// 0: (, max int), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox10setMaximumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:87
// index:0
// void setRange(int, int)
func (this *QSpinBox) SetRange(min int, max int) {
	// 0: (, min int, max int), (&min, &max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox8setRangeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:89
// index:0
// int displayIntegerBase()
func (this *QSpinBox) DisplayIntegerBase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox18displayIntegerBaseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:90
// index:0
// void setDisplayIntegerBase(int)
func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	// 0: (, base int), (&base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox21setDisplayIntegerBaseEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:93
// index:0
// virtual
// bool event(class QEvent *)
func (this *QSpinBox) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:94
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QSpinBox) Validate(input unsafe.Pointer, pos int) {
	// 0: (, input QString &, pos int &), (input, &pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), input, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:95
// index:0
// virtual
// int valueFromText(const class QString &)
func (this *QSpinBox) ValueFromText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox13valueFromTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:96
// index:0
// virtual
// QString textFromValue(int)
func (this *QSpinBox) TextFromValue(val int) {
	// 0: (, val int), (&val)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox13textFromValueEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:97
// index:0
// virtual
// void fixup(class QString &)
func (this *QSpinBox) Fixup(str unsafe.Pointer) {
	// 0: (, str QString &), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox5fixupER7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:101
// index:0
// void setValue(int)
func (this *QSpinBox) SetValue(val int) {
	// 0: (, val int), (&val)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox8setValueEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:104
// index:0
// void valueChanged(int)
func (this *QSpinBox) ValueChanged(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:105
// index:1
// void valueChanged(const class QString &)
func (this *QSpinBox) ValueChanged_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
