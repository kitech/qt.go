package qtwidgets

// /usr/include/qt/QtWidgets/qinputdialog.h
// #include <qinputdialog.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 96
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

/*

 */
type QInputDialog struct {
	*QDialog
}
type QInputDialog_ITF interface {
	QDialog_ITF
	QInputDialog_PTR() *QInputDialog
}

func (ptr *QInputDialog) QInputDialog_PTR() *QInputDialog { return ptr }

func (this *QInputDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QInputDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQInputDialogFromPointer(cthis unsafe.Pointer) *QInputDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QInputDialog{bcthis0}
}
func (*QInputDialog) NewFromPointer(cthis unsafe.Pointer) *QInputDialog {
	return NewQInputDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QInputDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qinputdialog.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a new input dialog with the given parent and window flags.

This function was introduced in  Qt 4.5.
*/
func (*QInputDialog) NewForInherit(parent QWidget_ITF /*777 QWidget **/, flags int) *QInputDialog {
	return NewQInputDialog(parent, flags)
}
func NewQInputDialog(parent QWidget_ITF /*777 QWidget **/, flags int) *QInputDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQInputDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QInputDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qinputdialog.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a new input dialog with the given parent and window flags.

This function was introduced in  Qt 4.5.
*/
func (*QInputDialog) NewForInherit__() *QInputDialog {
	return NewQInputDialog__()
}
func NewQInputDialog__() *QInputDialog {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQInputDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QInputDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qinputdialog.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a new input dialog with the given parent and window flags.

This function was introduced in  Qt 4.5.
*/
func (*QInputDialog) NewForInherit__1(parent QWidget_ITF /*777 QWidget **/) *QInputDialog {
	return NewQInputDialog__1(parent)
}
func NewQInputDialog__1(parent QWidget_ITF /*777 QWidget **/) *QInputDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQInputDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QInputDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qinputdialog.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QInputDialog()

/*

 */
func DeleteQInputDialog(this *QInputDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMode(QInputDialog::InputMode)

/*

 */
func (this *QInputDialog) SetInputMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog12setInputModeENS_9InputModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QInputDialog::InputMode inputMode() const

/*

 */
func (this *QInputDialog) InputMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog9inputModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelText(const QString &)

/*

 */
func (this *QInputDialog) SetLabelText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog12setLabelTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString labelText() const

/*

 */
func (this *QInputDialog) LabelText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog9labelTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QInputDialog::InputDialogOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options and testOption().
*/
func (this *QInputDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9setOptionENS_17InputDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QInputDialog::InputDialogOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options and testOption().
*/
func (this *QInputDialog) SetOption__(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9setOptionENS_17InputDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QInputDialog::InputDialogOption) const

/*
Returns true if the given option is enabled; otherwise, returns false.

See also options and setOption().
*/
func (this *QInputDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10testOptionENS_17InputDialogOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qinputdialog.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QInputDialog::InputDialogOptions)

/*

 */
func (this *QInputDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog10setOptionsE6QFlagsINS_17InputDialogOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QInputDialog::InputDialogOptions options() const

/*

 */
func (this *QInputDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextValue(const QString &)

/*

 */
func (this *QInputDialog) SetTextValue(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog12setTextValueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString textValue() const

/*

 */
func (this *QInputDialog) TextValue() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog9textValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextEchoMode(QLineEdit::EchoMode)

/*

 */
func (this *QInputDialog) SetTextEchoMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog15setTextEchoModeEN9QLineEdit8EchoModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QLineEdit::EchoMode textEchoMode() const

/*

 */
func (this *QInputDialog) TextEchoMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog12textEchoModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setComboBoxEditable(bool)

/*

 */
func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog19setComboBoxEditableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isComboBoxEditable() const

/*

 */
func (this *QInputDialog) IsComboBoxEditable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog18isComboBoxEditableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qinputdialog.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setComboBoxItems(const QStringList &)

/*

 */
func (this *QInputDialog) SetComboBoxItems(items qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg0 = items.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16setComboBoxItemsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList comboBoxItems() const

/*

 */
func (this *QInputDialog) ComboBoxItems() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog13comboBoxItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntValue(int)

/*

 */
func (this *QInputDialog) SetIntValue(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog11setIntValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] int intValue() const

/*

 */
func (this *QInputDialog) IntValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog8intValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntMinimum(int)

/*

 */
func (this *QInputDialog) SetIntMinimum(min int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog13setIntMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] int intMinimum() const

/*

 */
func (this *QInputDialog) IntMinimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10intMinimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntMaximum(int)

/*

 */
func (this *QInputDialog) SetIntMaximum(max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog13setIntMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int intMaximum() const

/*

 */
func (this *QInputDialog) IntMaximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10intMaximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntRange(int, int)

/*
Sets the range of integer values accepted by the dialog when used in IntInput mode, with minimum and maximum values specified by min and max respectively.
*/
func (this *QInputDialog) SetIntRange(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog11setIntRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntStep(int)

/*

 */
func (this *QInputDialog) SetIntStep(step int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog10setIntStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int intStep() const

/*

 */
func (this *QInputDialog) IntStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog7intStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleValue(double)

/*

 */
func (this *QInputDialog) SetDoubleValue(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleValueEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleValue() const

/*

 */
func (this *QInputDialog) DoubleValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog11doubleValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleMinimum(double)

/*

 */
func (this *QInputDialog) SetDoubleMinimum(min float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMinimumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleMinimum() const

/*

 */
func (this *QInputDialog) DoubleMinimum() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMinimumEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleMaximum(double)

/*

 */
func (this *QInputDialog) SetDoubleMaximum(max float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMaximumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleMaximum() const

/*

 */
func (this *QInputDialog) DoubleMaximum() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMaximumEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleRange(double, double)

/*
Sets the range of double precision floating point values accepted by the dialog when used in DoubleInput mode, with minimum and maximum values specified by min and max respectively.
*/
func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleRangeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleDecimals(int)

/*

 */
func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog17setDoubleDecimalsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), decimals)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] int doubleDecimals() const

/*

 */
func (this *QInputDialog) DoubleDecimals() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog14doubleDecimalsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOkButtonText(const QString &)

/*

 */
func (this *QInputDialog) SetOkButtonText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog15setOkButtonTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QString okButtonText() const

/*

 */
func (this *QInputDialog) OkButtonText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog12okButtonTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCancelButtonText(const QString &)

/*

 */
func (this *QInputDialog) SetCancelButtonText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog19setCancelButtonTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cancelButtonText() const

/*

 */
func (this *QInputDialog) CancelButtonText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog16cancelButtonTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)

/*
This is an overloaded function.

This function connects one of its signals to the slot specified by receiver and member. The specific signal depends on the arguments that are specified in member. These are:


textValueSelected() if member has a QString for its first argument.
intValueSelected() if member has an int for its first argument.
doubleValueSelected() if member has a double for its first argument.
accepted() if member has NO arguments.


The signal will be disconnected from the slot when the dialog is closed.

This function was introduced in  Qt 4.5.
*/
func (this *QInputDialog) Open(receiver qtcore.QObject_ITF /*777 QObject **/, member string) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QInputDialog) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QInputDialog) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:160
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().
*/
func (this *QInputDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the line edit. mode is the echo mode the line edit will use. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's line edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getText(this, tr("QInputDialog::getText()"),
                                           tr("User name:"), QLineEdit::Normal,
                                           QDir::home().dirName(), &ok);
      if (ok && !text.isEmpty())
          textLabel->setText(text);



See also getInt(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetText(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string, ok *bool, flags int, inputMethodHints int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(text)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QInputDialog_GetText(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string, ok *bool, flags int, inputMethodHints int) string {
	var nilthis *QInputDialog
	rv := nilthis.GetText(parent, title, label, echo, text, ok, flags, inputMethodHints)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the line edit. mode is the echo mode the line edit will use. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's line edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getText(this, tr("QInputDialog::getText()"),
                                           tr("User name:"), QLineEdit::Normal,
                                           QDir::home().dirName(), &ok);
      if (ok && !text.isEmpty())
          textLabel->setText(text);



See also getInt(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetText__(parent QWidget_ITF /*777 QWidget **/, title string, label string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QLineEdit::EchoMode=Elaborated, QLineEdit::EchoMode=Enum, , Invalid
	echo := 0
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 6, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 7, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the line edit. mode is the echo mode the line edit will use. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's line edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getText(this, tr("QInputDialog::getText()"),
                                           tr("User name:"), QLineEdit::Normal,
                                           QDir::home().dirName(), &ok);
      if (ok && !text.isEmpty())
          textLabel->setText(text);



See also getInt(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetText__1(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 6, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 7, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the line edit. mode is the echo mode the line edit will use. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's line edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getText(this, tr("QInputDialog::getText()"),
                                           tr("User name:"), QLineEdit::Normal,
                                           QDir::home().dirName(), &ok);
      if (ok && !text.isEmpty())
          textLabel->setText(text);



See also getInt(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetText__2(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(text)
	var convArg4 = tmpArg4.GetCthis()
	// arg: 5, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 6, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 7, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the line edit. mode is the echo mode the line edit will use. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's line edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getText(this, tr("QInputDialog::getText()"),
                                           tr("User name:"), QLineEdit::Normal,
                                           QDir::home().dirName(), &ok);
      if (ok && !text.isEmpty())
          textLabel->setText(text);



See also getInt(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetText__3(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string, ok *bool) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(text)
	var convArg4 = tmpArg4.GetCthis()
	// arg: 6, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 7, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the line edit. mode is the echo mode the line edit will use. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's line edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getText(this, tr("QInputDialog::getText()"),
                                           tr("User name:"), QLineEdit::Normal,
                                           QDir::home().dirName(), &ok);
      if (ok && !text.isEmpty())
          textLabel->setText(text);



See also getInt(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetText__4(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string, ok *bool, flags int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(text)
	var convArg4 = tmpArg4.GetCthis()
	// arg: 7, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getMultiLineText(QWidget *, const QString &, const QString &, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a multiline string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the plain text edit. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's plain text edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getMultiLineText(this, tr("QInputDialog::getMultiLineText()"),
                                                    tr("Address:"), "John Doe\nFreedom Street", &ok);
      if (ok && !text.isEmpty())
          multiLineTextLabel->setText(text);



This function was introduced in  Qt 5.2.

See also getInt(), getDouble(), getItem(), and getText().
*/
func (this *QInputDialog) GetMultiLineText(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string, ok *bool, flags int, inputMethodHints int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QInputDialog_GetMultiLineText(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string, ok *bool, flags int, inputMethodHints int) string {
	var nilthis *QInputDialog
	rv := nilthis.GetMultiLineText(parent, title, label, text, ok, flags, inputMethodHints)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getMultiLineText(QWidget *, const QString &, const QString &, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a multiline string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the plain text edit. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's plain text edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getMultiLineText(this, tr("QInputDialog::getMultiLineText()"),
                                                    tr("Address:"), "John Doe\nFreedom Street", &ok);
      if (ok && !text.isEmpty())
          multiLineTextLabel->setText(text);



This function was introduced in  Qt 5.2.

See also getInt(), getDouble(), getItem(), and getText().
*/
func (this *QInputDialog) GetMultiLineText__(parent QWidget_ITF /*777 QWidget **/, title string, label string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 5, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 6, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getMultiLineText(QWidget *, const QString &, const QString &, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a multiline string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the plain text edit. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's plain text edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getMultiLineText(this, tr("QInputDialog::getMultiLineText()"),
                                                    tr("Address:"), "John Doe\nFreedom Street", &ok);
      if (ok && !text.isEmpty())
          multiLineTextLabel->setText(text);



This function was introduced in  Qt 5.2.

See also getInt(), getDouble(), getItem(), and getText().
*/
func (this *QInputDialog) GetMultiLineText__1(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 5, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 6, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getMultiLineText(QWidget *, const QString &, const QString &, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a multiline string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the plain text edit. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's plain text edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getMultiLineText(this, tr("QInputDialog::getMultiLineText()"),
                                                    tr("Address:"), "John Doe\nFreedom Street", &ok);
      if (ok && !text.isEmpty())
          multiLineTextLabel->setText(text);



This function was introduced in  Qt 5.2.

See also getInt(), getDouble(), getItem(), and getText().
*/
func (this *QInputDialog) GetMultiLineText__2(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string, ok *bool) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 5, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 6, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getMultiLineText(QWidget *, const QString &, const QString &, const QString &, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to get a multiline string from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). text is the default text which is placed in the plain text edit. inputMethodHints is the input method hints that will be used in the edit widget if an input method is active.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the specified widget flags.

If the dialog is accepted, this function returns the text in the dialog's plain text edit. If the dialog is rejected, a null QString is returned.

Use this static function like this:


      bool ok;
      QString text = QInputDialog::getMultiLineText(this, tr("QInputDialog::getMultiLineText()"),
                                                    tr("Address:"), "John Doe\nFreedom Street", &ok);
      if (ok && !text.isEmpty())
          multiLineTextLabel->setText(text);



This function was introduced in  Qt 5.2.

See also getInt(), getDouble(), getItem(), and getText().
*/
func (this *QInputDialog) GetMultiLineText__3(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string, ok *bool, flags int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 6, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, bool, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to let the user select an item from a string list.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). items is the string list which is inserted into the combo box. current is the number of the item which should be the current item. inputMethodHints is the input method hints that will be used if the combo box is editable and an input method is active.

If editable is true the user can enter their own text; otherwise, the user may only select one of the existing items.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the text of the current item, or if editable is true, the current text of the combo box.

Use this static function like this:


      QStringList items;
      items << tr("Spring") << tr("Summer") << tr("Fall") << tr("Winter");

      bool ok;
      QString item = QInputDialog::getItem(this, tr("QInputDialog::getItem()"),
                                           tr("Season:"), items, 0, false, &ok);
      if (ok && !item.isEmpty())
          itemLabel->setText(item);



See also getText(), getInt(), getDouble(), and getMultiLineText().
*/
func (this *QInputDialog) GetItem(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool, ok *bool, flags int, inputMethodHints int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg3 = items.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QInputDialog_GetItem(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool, ok *bool, flags int, inputMethodHints int) string {
	var nilthis *QInputDialog
	rv := nilthis.GetItem(parent, title, label, items, current, editable, ok, flags, inputMethodHints)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, bool, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to let the user select an item from a string list.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). items is the string list which is inserted into the combo box. current is the number of the item which should be the current item. inputMethodHints is the input method hints that will be used if the combo box is editable and an input method is active.

If editable is true the user can enter their own text; otherwise, the user may only select one of the existing items.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the text of the current item, or if editable is true, the current text of the combo box.

Use this static function like this:


      QStringList items;
      items << tr("Spring") << tr("Summer") << tr("Fall") << tr("Winter");

      bool ok;
      QString item = QInputDialog::getItem(this, tr("QInputDialog::getItem()"),
                                           tr("Season:"), items, 0, false, &ok);
      if (ok && !item.isEmpty())
          itemLabel->setText(item);



See also getText(), getInt(), getDouble(), and getMultiLineText().
*/
func (this *QInputDialog) GetItem__(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg3 = items.QStringList_PTR().GetCthis()
	}
	// arg: 4, int=Int, =Invalid, , Invalid
	current := int(0)
	// arg: 5, bool=Bool, =Invalid, , Invalid
	editable := true
	// arg: 6, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 7, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 8, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, bool, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to let the user select an item from a string list.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). items is the string list which is inserted into the combo box. current is the number of the item which should be the current item. inputMethodHints is the input method hints that will be used if the combo box is editable and an input method is active.

If editable is true the user can enter their own text; otherwise, the user may only select one of the existing items.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the text of the current item, or if editable is true, the current text of the combo box.

Use this static function like this:


      QStringList items;
      items << tr("Spring") << tr("Summer") << tr("Fall") << tr("Winter");

      bool ok;
      QString item = QInputDialog::getItem(this, tr("QInputDialog::getItem()"),
                                           tr("Season:"), items, 0, false, &ok);
      if (ok && !item.isEmpty())
          itemLabel->setText(item);



See also getText(), getInt(), getDouble(), and getMultiLineText().
*/
func (this *QInputDialog) GetItem__1(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg3 = items.QStringList_PTR().GetCthis()
	}
	// arg: 5, bool=Bool, =Invalid, , Invalid
	editable := true
	// arg: 6, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 7, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 8, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, bool, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to let the user select an item from a string list.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). items is the string list which is inserted into the combo box. current is the number of the item which should be the current item. inputMethodHints is the input method hints that will be used if the combo box is editable and an input method is active.

If editable is true the user can enter their own text; otherwise, the user may only select one of the existing items.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the text of the current item, or if editable is true, the current text of the combo box.

Use this static function like this:


      QStringList items;
      items << tr("Spring") << tr("Summer") << tr("Fall") << tr("Winter");

      bool ok;
      QString item = QInputDialog::getItem(this, tr("QInputDialog::getItem()"),
                                           tr("Season:"), items, 0, false, &ok);
      if (ok && !item.isEmpty())
          itemLabel->setText(item);



See also getText(), getInt(), getDouble(), and getMultiLineText().
*/
func (this *QInputDialog) GetItem__2(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg3 = items.QStringList_PTR().GetCthis()
	}
	// arg: 6, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 7, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 8, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, bool, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to let the user select an item from a string list.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). items is the string list which is inserted into the combo box. current is the number of the item which should be the current item. inputMethodHints is the input method hints that will be used if the combo box is editable and an input method is active.

If editable is true the user can enter their own text; otherwise, the user may only select one of the existing items.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the text of the current item, or if editable is true, the current text of the combo box.

Use this static function like this:


      QStringList items;
      items << tr("Spring") << tr("Summer") << tr("Fall") << tr("Winter");

      bool ok;
      QString item = QInputDialog::getItem(this, tr("QInputDialog::getItem()"),
                                           tr("Season:"), items, 0, false, &ok);
      if (ok && !item.isEmpty())
          itemLabel->setText(item);



See also getText(), getInt(), getDouble(), and getMultiLineText().
*/
func (this *QInputDialog) GetItem__3(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool, ok *bool) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg3 = items.QStringList_PTR().GetCthis()
	}
	// arg: 7, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	// arg: 8, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, bool, bool *, Qt::WindowFlags, Qt::InputMethodHints)

/*
Static convenience function to let the user select an item from a string list.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). items is the string list which is inserted into the combo box. current is the number of the item which should be the current item. inputMethodHints is the input method hints that will be used if the combo box is editable and an input method is active.

If editable is true the user can enter their own text; otherwise, the user may only select one of the existing items.

If ok is nonnull *a ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the text of the current item, or if editable is true, the current text of the combo box.

Use this static function like this:


      QStringList items;
      items << tr("Spring") << tr("Summer") << tr("Fall") << tr("Winter");

      bool ok;
      QString item = QInputDialog::getItem(this, tr("QInputDialog::getItem()"),
                                           tr("Season:"), items, 0, false, &ok);
      if (ok && !item.isEmpty())
          itemLabel->setText(item);



See also getText(), getInt(), getDouble(), and getMultiLineText().
*/
func (this *QInputDialog) GetItem__4(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool, ok *bool, flags int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if items != nil && items.QStringList_PTR() != nil {
		convArg3 = items.QStringList_PTR().GetCthis()
	}
	// arg: 8, Qt::InputMethodHints=Elaborated, Qt::InputMethodHints=Typedef, QFlags<Qt::InputMethodHint>, Unexposed
	inputMethodHints := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int, step int, ok *bool, flags int) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QInputDialog_GetInt(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int, step int, ok *bool, flags int) int {
	var nilthis *QInputDialog
	rv := nilthis.GetInt(parent, title, label, value, minValue, maxValue, step, ok, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt__(parent QWidget_ITF /*777 QWidget **/, title string, label string) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, int=Int, =Invalid, , Invalid
	value := int(0)
	// arg: 4, int=Int, =Invalid, , Invalid
	minValue := int(-2147483647)
	// arg: 5, int=Int, =Invalid, , Invalid
	maxValue := int(2147483647)
	// arg: 6, int=Int, =Invalid, , Invalid
	step := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt__1(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, int=Int, =Invalid, , Invalid
	minValue := int(-2147483647)
	// arg: 5, int=Int, =Invalid, , Invalid
	maxValue := int(2147483647)
	// arg: 6, int=Int, =Invalid, , Invalid
	step := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt__2(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 5, int=Int, =Invalid, , Invalid
	maxValue := int(2147483647)
	// arg: 6, int=Int, =Invalid, , Invalid
	step := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt__3(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 6, int=Int, =Invalid, , Invalid
	step := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt__4(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int, step int) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get an integer input from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default integer which the spinbox will be set to. min and max are the minimum and maximum values the user may choose. step is the amount by which the values change as the user presses the arrow buttons to increment or decrement the value.

If ok is nonnull *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

On success, this function returns the integer which has been entered by the user; on failure, it returns the initial value.

Use this static function like this:


      bool ok;
      int i = QInputDialog::getInt(this, tr("QInputDialog::getInteger()"),
                                   tr("Percentage:"), 25, 0, 100, 1, &ok);
      if (ok)
          integerLabel->setText(tr("%1%").arg(i));



This function was introduced in  Qt 4.5.

See also getText(), getDouble(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetInt__5(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int, step int, ok *bool) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool, flags int) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QInputDialog_GetDouble(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool, flags int) float64 {
	var nilthis *QInputDialog
	rv := nilthis.GetDouble(parent, title, label, value, minValue, maxValue, decimals, ok, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble__(parent QWidget_ITF /*777 QWidget **/, title string, label string) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, double=Double, =Invalid, , Invalid
	value := float64(0)
	// arg: 4, double=Double, =Invalid, , Invalid
	minValue := float64(-2147483647)
	// arg: 5, double=Double, =Invalid, , Invalid
	maxValue := float64(2147483647)
	// arg: 6, int=Int, =Invalid, , Invalid
	decimals := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble__1(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, double=Double, =Invalid, , Invalid
	minValue := float64(-2147483647)
	// arg: 5, double=Double, =Invalid, , Invalid
	maxValue := float64(2147483647)
	// arg: 6, int=Int, =Invalid, , Invalid
	decimals := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble__2(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 5, double=Double, =Invalid, , Invalid
	maxValue := float64(2147483647)
	// arg: 6, int=Int, =Invalid, , Invalid
	decimals := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble__3(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 6, int=Int, =Invalid, , Invalid
	decimals := int(1)
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble__4(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 7, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble__5(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 8, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:183
// index:1
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, bool *, Qt::WindowFlags, double)

/*
Static convenience function to get a floating point number from the user.

title is the text which is displayed in the title bar of the dialog. label is the text which is shown to the user (it should say what should be entered). value is the default floating point number that the line edit will be set to. min and max are the minimum and maximum values the user may choose. decimals is the maximum number of decimal places the number may have.

If ok is nonnull, *ok will be set to true if the user pressed OK and to false if the user pressed Cancel. The dialog's parent is parent. The dialog will be modal and uses the widget flags.

This function returns the floating point number which has been entered by the user.

Use this static function like this:


      bool ok;
      double d = QInputDialog::getDouble(this, tr("QInputDialog::getDouble()"),
                                         tr("Amount:"), 37.56, -10000, 10000, 2, &ok);
      if (ok)
          doubleLabel->setText(QString("$%1").arg(d));



See also getText(), getInt(), getItem(), and getMultiLineText().
*/
func (this *QInputDialog) GetDouble_1(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool, flags int, step float64) float64 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEEd", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, ok, flags, step)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QInputDialog_GetDouble_1(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool, flags int, step float64) float64 {
	var nilthis *QInputDialog
	rv := nilthis.GetDouble_1(parent, title, label, value, minValue, maxValue, decimals, ok, flags, step)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleStep(double)

/*

 */
func (this *QInputDialog) SetDoubleStep(step float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog13setDoubleStepEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:197
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleStep() const

/*

 */
func (this *QInputDialog) DoubleStep() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10doubleStepEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textValueChanged(const QString &)

/*
This signal is emitted whenever the text string changes in the dialog. The current string is specified by text.

This signal is only relevant when the input dialog is used in TextInput mode.

Note: Notifier signal for property textValue.
*/
func (this *QInputDialog) TextValueChanged(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16textValueChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textValueSelected(const QString &)

/*
This signal is emitted whenever the user selects a text string by accepting the dialog; for example, by clicking the OK button. The selected string is specified by text.

This signal is only relevant when the input dialog is used in TextInput mode.
*/
func (this *QInputDialog) TextValueSelected(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog17textValueSelectedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void intValueChanged(int)

/*
This signal is emitted whenever the integer value changes in the dialog. The current value is specified by value.

This signal is only relevant when the input dialog is used in IntInput mode.

Note: Notifier signal for property intValue.
*/
func (this *QInputDialog) IntValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog15intValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void intValueSelected(int)

/*
This signal is emitted whenever the user selects a integer value by accepting the dialog; for example, by clicking the OK button. The selected value is specified by value.

This signal is only relevant when the input dialog is used in IntInput mode.
*/
func (this *QInputDialog) IntValueSelected(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16intValueSelectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void doubleValueChanged(double)

/*
This signal is emitted whenever the double value changes in the dialog. The current value is specified by value.

This signal is only relevant when the input dialog is used in DoubleInput mode.

Note: Notifier signal for property doubleValue.
*/
func (this *QInputDialog) DoubleValueChanged(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog18doubleValueChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void doubleValueSelected(double)

/*
This signal is emitted whenever the user selects a double value by accepting the dialog; for example, by clicking the OK button. The selected value is specified by value.

This signal is only relevant when the input dialog is used in DoubleInput mode.
*/
func (this *QInputDialog) DoubleValueSelected(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog19doubleValueSelectedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void done(int)

/*
Reimplemented from QDialog::done().

Closes the dialog and sets its result code to result. If this dialog is shown with exec(), done() causes the local event loop to finish, and exec() to return result.

See also QDialog::done().
*/
func (this *QInputDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QInputDialog__InputDialogOption = int

//
const QInputDialog__NoButtons QInputDialog__InputDialogOption = 1

//
const QInputDialog__UseListViewForComboBoxItems QInputDialog__InputDialogOption = 2

//
const QInputDialog__UsePlainTextEditForTextInput QInputDialog__InputDialogOption = 4

func (this *QInputDialog) InputDialogOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QInputDialog_InputDialogOptionItemName(val int) string {
	var nilthis *QInputDialog
	return nilthis.InputDialogOptionItemName(val)
}

/*
This enum describes the different modes of input that can be selected for the dialog.



This enum was introduced or modified in  Qt 4.5.

See also inputMode.

*/
type QInputDialog__InputMode = int

// Used to input text strings.
const QInputDialog__TextInput QInputDialog__InputMode = 0

// Used to input integers.
const QInputDialog__IntInput QInputDialog__InputMode = 1

// Used to input floating point numbers with double precision accuracy.
const QInputDialog__DoubleInput QInputDialog__InputMode = 2

func (this *QInputDialog) InputModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QInputDialog_InputModeItemName(val int) string {
	var nilthis *QInputDialog
	return nilthis.InputModeItemName(val)
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
