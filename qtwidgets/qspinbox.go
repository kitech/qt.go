package qtwidgets

// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QSpinBox) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// QValidator::State validate(QString &, int &)
func (this *QSpinBox) InheritValidate(f func(input string, pos int) int) {
	qtrt.SetAllInheritCallback(this, "validate", f)
}

// int valueFromText(const QString &)
func (this *QSpinBox) InheritValueFromText(f func(text string) int) {
	qtrt.SetAllInheritCallback(this, "valueFromText", f)
}

// QString textFromValue(int)
func (this *QSpinBox) InheritTextFromValue(f func(val int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "textFromValue", f)
}

// void fixup(QString &)
func (this *QSpinBox) InheritFixup(f func(str string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "fixup", f)
}

/*

 */
type QSpinBox struct {
	*QAbstractSpinBox
}
type QSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QSpinBox_PTR() *QSpinBox
}

func (ptr *QSpinBox) QSpinBox_PTR() *QSpinBox { return ptr }

func (this *QSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSpinBox.GetCthis()
	}
}
func (this *QSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSpinBox = NewQAbstractSpinBoxFromPointer(cthis)
}
func NewQSpinBoxFromPointer(cthis unsafe.Pointer) *QSpinBox {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QSpinBox{bcthis0}
}
func (*QSpinBox) NewFromPointer(cthis unsafe.Pointer) *QSpinBox {
	return NewQSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qspinbox.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSpinBox(QWidget *)

/*
Constructs a spin box with 0 as minimum value and 99 as maximum value, a step value of 1. The value is initially set to 0. It is parented to parent.

See also setMinimum(), setMaximum(), and setSingleStep().
*/
func (*QSpinBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QSpinBox {
	return NewQSpinBox(parent)
}
func NewQSpinBox(parent QWidget_ITF /*777 QWidget **/) *QSpinBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSpinBox(QWidget *)

/*
Constructs a spin box with 0 as minimum value and 99 as maximum value, a step value of 1. The value is initially set to 0. It is parented to parent.

See also setMinimum(), setMaximum(), and setSingleStep().
*/
func (*QSpinBox) NewForInheritp() *QSpinBox {
	return NewQSpinBoxp()
}
func NewQSpinBoxp() *QSpinBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSpinBox()

/*

 */
func DeleteQSpinBox(this *QSpinBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qspinbox.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int value() const

/*

 */
func (this *QSpinBox) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString prefix() const

/*

 */
func (this *QSpinBox) Prefix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox6prefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrefix(const QString &)

/*

 */
func (this *QSpinBox) SetPrefix(prefix string) {
	var tmpArg0 = qtcore.NewQString5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox9setPrefixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffix() const

/*

 */
func (this *QSpinBox) Suffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox6suffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSuffix(const QString &)

/*

 */
func (this *QSpinBox) SetSuffix(suffix string) {
	var tmpArg0 = qtcore.NewQString5(suffix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox9setSuffixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cleanText() const

/*

 */
func (this *QSpinBox) CleanText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox9cleanTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int singleStep() const

/*

 */
func (this *QSpinBox) SingleStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox10singleStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleStep(int)

/*

 */
func (this *QSpinBox) SetSingleStep(val int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox13setSingleStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum() const

/*

 */
func (this *QSpinBox) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)

/*

 */
func (this *QSpinBox) SetMinimum(min int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum() const

/*

 */
func (this *QSpinBox) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)

/*

 */
func (this *QSpinBox) SetMaximum(max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)

/*
Convenience function to set the minimum, and maximum values with a single function call.


  setRange(minimum, maximum);



is equivalent to:


  setMinimum(minimum);
  setMaximum(maximum);



See also minimum and maximum.
*/
func (this *QSpinBox) SetRange(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int displayIntegerBase() const

/*

 */
func (this *QSpinBox) DisplayIntegerBase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox18displayIntegerBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDisplayIntegerBase(int)

/*

 */
func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox21setDisplayIntegerBaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), base)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSpinBox) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qspinbox.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const

/*
Reimplemented from QAbstractSpinBox::validate().
*/
func (this *QSpinBox) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int valueFromText(const QString &) const

/*
This virtual function is used by the spin box whenever it needs to interpret text entered by the user as a value.

Subclasses that need to display spin box values in a non-numeric way need to reimplement this function.

Note: QSpinBox handles specialValueText() separately; this function is only concerned with the other values.

See also textFromValue() and validate().
*/
func (this *QSpinBox) ValueFromText(text string) int {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox13valueFromTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QString textFromValue(int) const

/*
This virtual function is used by the spin box whenever it needs to display the given value. The default implementation returns a string containing value printed in the standard way using QWidget::locale().toString(), but with the thousand separator removed unless setGroupSeparatorShown() is set. Reimplementations may return anything. (See the example in the detailed description.)

Note: QSpinBox does not call this function for specialValueText() and that neither prefix() nor suffix() should be included in the return value.

If you reimplement this, you may also need to reimplement valueFromText() and validate()

See also valueFromText(), validate(), and QLocale::groupSeparator().
*/
func (this *QSpinBox) TextFromValue(val int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox13textFromValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &) const

/*
Reimplemented from QAbstractSpinBox::fixup().
*/
func (this *QSpinBox) Fixup(str string) {
	var tmpArg0 = qtcore.NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSpinBox5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)

/*

 */
func (this *QSpinBox) SetValue(val int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(int)

/*
This signal is emitted whenever the spin box's value is changed. The new value's integer value is passed in i.

Note: Signal valueChanged is overloaded in this class. To connect to this one using the function pointer syntax, you must specify the signal type in a static cast, as shown in this example:


  connect(spinBox, static_cast<void(QSpinBox::*)(int)>(&QSpinBox::valueChanged),
      [=](int i){ /-* ... *-/ });



Note: Notifier signal for property value.
*/
func (this *QSpinBox) ValueChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:106
// index:1
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QString &)

/*
This signal is emitted whenever the spin box's value is changed. The new value's integer value is passed in i.

Note: Signal valueChanged is overloaded in this class. To connect to this one using the function pointer syntax, you must specify the signal type in a static cast, as shown in this example:


  connect(spinBox, static_cast<void(QSpinBox::*)(int)>(&QSpinBox::valueChanged),
      [=](int i){ /-* ... *-/ });



Note: Notifier signal for property value.
*/
func (this *QSpinBox) ValueChanged1(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
