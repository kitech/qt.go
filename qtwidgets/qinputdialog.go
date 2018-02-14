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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

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
// [8] const QMetaObject * metaObject()
func (this *QInputDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qinputdialog.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputDialog(QWidget *, Qt::WindowFlags)
func NewQInputDialog(parent QWidget_ITF /*777 QWidget **/, flags int) *QInputDialog {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQInputDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qinputdialog.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QInputDialog()
func DeleteQInputDialog(this *QInputDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMode(enum QInputDialog::InputMode)
func (this *QInputDialog) SetInputMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog12setInputModeENS_9InputModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QInputDialog::InputMode inputMode()
func (this *QInputDialog) InputMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog9inputModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelText(const QString &)
func (this *QInputDialog) SetLabelText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog12setLabelTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString labelText()
func (this *QInputDialog) LabelText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog9labelTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QInputDialog::InputDialogOption, _Bool)
func (this *QInputDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9setOptionENS_17InputDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(enum QInputDialog::InputDialogOption)
func (this *QInputDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10testOptionENS_17InputDialogOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qinputdialog.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QInputDialog::InputDialogOptions)
func (this *QInputDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog10setOptionsE6QFlagsINS_17InputDialogOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QInputDialog::InputDialogOptions options()
func (this *QInputDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextValue(const QString &)
func (this *QInputDialog) SetTextValue(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog12setTextValueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString textValue()
func (this *QInputDialog) TextValue() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog9textValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextEchoMode(QLineEdit::EchoMode)
func (this *QInputDialog) SetTextEchoMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog15setTextEchoModeEN9QLineEdit8EchoModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QLineEdit::EchoMode textEchoMode()
func (this *QInputDialog) TextEchoMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog12textEchoModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setComboBoxEditable(_Bool)
func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog19setComboBoxEditableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isComboBoxEditable()
func (this *QInputDialog) IsComboBoxEditable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog18isComboBoxEditableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qinputdialog.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setComboBoxItems(const QStringList &)
func (this *QInputDialog) SetComboBoxItems(items qtcore.QStringList_ITF) {
	var convArg0 = items.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16setComboBoxItemsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList comboBoxItems()
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
func (this *QInputDialog) SetIntValue(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog11setIntValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] int intValue()
func (this *QInputDialog) IntValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog8intValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntMinimum(int)
func (this *QInputDialog) SetIntMinimum(min int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog13setIntMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] int intMinimum()
func (this *QInputDialog) IntMinimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10intMinimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntMaximum(int)
func (this *QInputDialog) SetIntMaximum(max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog13setIntMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int intMaximum()
func (this *QInputDialog) IntMaximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10intMaximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntRange(int, int)
func (this *QInputDialog) SetIntRange(min int, max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog11setIntRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntStep(int)
func (this *QInputDialog) SetIntStep(step int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog10setIntStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int intStep()
func (this *QInputDialog) IntStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog7intStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleValue(double)
func (this *QInputDialog) SetDoubleValue(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleValueEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleValue()
func (this *QInputDialog) DoubleValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog11doubleValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleMinimum(double)
func (this *QInputDialog) SetDoubleMinimum(min float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMinimumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleMinimum()
func (this *QInputDialog) DoubleMinimum() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMinimumEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleMaximum(double)
func (this *QInputDialog) SetDoubleMaximum(max float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMaximumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleMaximum()
func (this *QInputDialog) DoubleMaximum() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMaximumEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleRange(double, double)
func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleRangeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleDecimals(int)
func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog17setDoubleDecimalsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), decimals)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] int doubleDecimals()
func (this *QInputDialog) DoubleDecimals() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog14doubleDecimalsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOkButtonText(const QString &)
func (this *QInputDialog) SetOkButtonText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog15setOkButtonTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QString okButtonText()
func (this *QInputDialog) OkButtonText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog12okButtonTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCancelButtonText(const QString &)
func (this *QInputDialog) SetCancelButtonText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog19setCancelButtonTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cancelButtonText()
func (this *QInputDialog) CancelButtonText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog16cancelButtonTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qinputdialog.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)
func (this *QInputDialog) Open(receiver qtcore.QObject_ITF /*777 QObject **/, member string) {
	var convArg0 = receiver.QObject_PTR().GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
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
// [8] QSize sizeHint()
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
// [-2] void setVisible(_Bool)
func (this *QInputDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getText(QWidget *, const QString &, const QString &, QLineEdit::EchoMode, const QString &, _Bool *, Qt::WindowFlags, Qt::InputMethodHints)
func (this *QInputDialog) GetText(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string, ok unsafe.Pointer /*666*/, flags int, inputMethodHints int) string {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(text)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, echo, convArg4, &ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QInputDialog_GetText(parent QWidget_ITF /*777 QWidget **/, title string, label string, echo int, text string, ok unsafe.Pointer /*666*/, flags int, inputMethodHints int) string {
	var nilthis *QInputDialog
	rv := nilthis.GetText(parent, title, label, echo, text, ok, flags, inputMethodHints)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getMultiLineText(QWidget *, const QString &, const QString &, const QString &, _Bool *, Qt::WindowFlags, Qt::InputMethodHints)
func (this *QInputDialog) GetMultiLineText(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string, ok unsafe.Pointer /*666*/, flags int, inputMethodHints int) string {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, &ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QInputDialog_GetMultiLineText(parent QWidget_ITF /*777 QWidget **/, title string, label string, text string, ok unsafe.Pointer /*666*/, flags int, inputMethodHints int) string {
	var nilthis *QInputDialog
	rv := nilthis.GetMultiLineText(parent, title, label, text, ok, flags, inputMethodHints)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getItem(QWidget *, const QString &, const QString &, const QStringList &, int, _Bool, _Bool *, Qt::WindowFlags, Qt::InputMethodHints)
func (this *QInputDialog) GetItem(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool, ok unsafe.Pointer /*666*/, flags int, inputMethodHints int) string {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 = items.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, current, editable, &ok, flags, inputMethodHints)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QInputDialog_GetItem(parent QWidget_ITF /*777 QWidget **/, title string, label string, items qtcore.QStringList_ITF, current int, editable bool, ok unsafe.Pointer /*666*/, flags int, inputMethodHints int) string {
	var nilthis *QInputDialog
	rv := nilthis.GetItem(parent, title, label, items, current, editable, ok, flags, inputMethodHints)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [4] int getInt(QWidget *, const QString &, const QString &, int, int, int, int, _Bool *, Qt::WindowFlags)
func (this *QInputDialog) GetInt(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int, step int, ok unsafe.Pointer /*666*/, flags int) int {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, &ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QInputDialog_GetInt(parent QWidget_ITF /*777 QWidget **/, title string, label string, value int, minValue int, maxValue int, step int, ok unsafe.Pointer /*666*/, flags int) int {
	var nilthis *QInputDialog
	rv := nilthis.GetInt(parent, title, label, value, minValue, maxValue, step, ok, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, _Bool *, Qt::WindowFlags)
func (this *QInputDialog) GetDouble(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int) float64 {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, &ok, flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QInputDialog_GetDouble(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int) float64 {
	var nilthis *QInputDialog
	rv := nilthis.GetDouble(parent, title, label, value, minValue, maxValue, decimals, ok, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:183
// index:1
// Public static Visibility=Default Availability=Available
// [8] double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, _Bool *, Qt::WindowFlags, double)
func (this *QInputDialog) GetDouble_1(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int, step float64) float64 {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(label)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEEd", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, &ok, flags, step)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QInputDialog_GetDouble_1(parent QWidget_ITF /*777 QWidget **/, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int, step float64) float64 {
	var nilthis *QInputDialog
	rv := nilthis.GetDouble_1(parent, title, label, value, minValue, maxValue, decimals, ok, flags, step)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDoubleStep(double)
func (this *QInputDialog) SetDoubleStep(step float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog13setDoubleStepEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:197
// index:0
// Public Visibility=Default Availability=Available
// [8] double doubleStep()
func (this *QInputDialog) DoubleStep() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QInputDialog10doubleStepEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textValueChanged(const QString &)
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
func (this *QInputDialog) IntValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog15intValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void intValueSelected(int)
func (this *QInputDialog) IntValueSelected(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog16intValueSelectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void doubleValueChanged(double)
func (this *QInputDialog) DoubleValueChanged(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog18doubleValueChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void doubleValueSelected(double)
func (this *QInputDialog) DoubleValueSelected(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog19doubleValueSelectedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QInputDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QInputDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

type QInputDialog__InputDialogOption = int

const QInputDialog__NoButtons QInputDialog__InputDialogOption = 1
const QInputDialog__UseListViewForComboBoxItems QInputDialog__InputDialogOption = 2
const QInputDialog__UsePlainTextEditForTextInput QInputDialog__InputDialogOption = 4

type QInputDialog__InputMode = int

const QInputDialog__TextInput QInputDialog__InputMode = 0
const QInputDialog__IntInput QInputDialog__InputMode = 1
const QInputDialog__DoubleInput QInputDialog__InputMode = 2

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
