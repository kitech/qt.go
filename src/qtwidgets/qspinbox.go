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
func NewQSpinBoxFromPointer(cthis unsafe.Pointer) *QSpinBox {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QSpinBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qspinbox.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSpinBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:65
// index:0
// Public
// void QSpinBox(class QWidget *)
func NewQSpinBox(parent unsafe.Pointer) *QSpinBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSpinBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// Public virtual
// void ~QSpinBox()
func DeleteQSpinBox(*QSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:68
// index:0
// Public
// int value()
func (this *QSpinBox) Value() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:70
// index:0
// Public
// QString prefix()
func (this *QSpinBox) Prefix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox6prefixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:71
// index:0
// Public
// void setPrefix(const class QString &)
func (this *QSpinBox) SetPrefix(prefix *qtcore.QString) {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox9setPrefixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:73
// index:0
// Public
// QString suffix()
func (this *QSpinBox) Suffix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox6suffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:74
// index:0
// Public
// void setSuffix(const class QString &)
func (this *QSpinBox) SetSuffix(suffix *qtcore.QString) {
	var convArg0 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox9setSuffixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:76
// index:0
// Public
// QString cleanText()
func (this *QSpinBox) CleanText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox9cleanTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:78
// index:0
// Public
// int singleStep()
func (this *QSpinBox) SingleStep() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox10singleStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:79
// index:0
// Public
// void setSingleStep(int)
func (this *QSpinBox) SetSingleStep(val int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox13setSingleStepEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:81
// index:0
// Public
// int minimum()
func (this *QSpinBox) Minimum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox7minimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:82
// index:0
// Public
// void setMinimum(int)
func (this *QSpinBox) SetMinimum(min int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox10setMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:84
// index:0
// Public
// int maximum()
func (this *QSpinBox) Maximum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox7maximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:85
// index:0
// Public
// void setMaximum(int)
func (this *QSpinBox) SetMaximum(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox10setMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:87
// index:0
// Public
// void setRange(int, int)
func (this *QSpinBox) SetRange(min int, max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox8setRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:89
// index:0
// Public
// int displayIntegerBase()
func (this *QSpinBox) DisplayIntegerBase() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox18displayIntegerBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:90
// index:0
// Public
// void setDisplayIntegerBase(int)
func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox21setDisplayIntegerBaseEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:93
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSpinBox) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:94
// index:0
// Protected virtual
// QValidator::State validate(class QString &, int &)
func (this *QSpinBox) Validate(input *qtcore.QString, pos int) interface{} {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:95
// index:0
// Protected virtual
// int valueFromText(const class QString &)
func (this *QSpinBox) ValueFromText(text *qtcore.QString) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox13valueFromTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:96
// index:0
// Protected virtual
// QString textFromValue(int)
func (this *QSpinBox) TextFromValue(val int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox13textFromValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &val)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qspinbox.h:97
// index:0
// Protected virtual
// void fixup(class QString &)
func (this *QSpinBox) Fixup(str *qtcore.QString) {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:101
// index:0
// Public
// void setValue(int)
func (this *QSpinBox) SetValue(val int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox8setValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:104
// index:0
// Public
// void valueChanged(int)
func (this *QSpinBox) ValueChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:105
// index:1
// Public
// void valueChanged(const class QString &)
func (this *QSpinBox) ValueChanged_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
