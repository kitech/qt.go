//  header block begin
// /usr/include/qt/QtWidgets/qinputdialog.h
// #include <qinputdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 95
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
type QInputDialog struct {
	*QDialog
}

func (this *QInputDialog) GetCthis() unsafe.Pointer {
	return this.QDialog.GetCthis()
}
func NewQInputDialogFromPointer(cthis unsafe.Pointer) *QInputDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QInputDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qinputdialog.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QInputDialog) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:95
// index:0
// Public virtual
// void ~QInputDialog()
func DeleteQInputDialog(*QInputDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:97
// index:0
// Public
// void setInputMode(enum QInputDialog::InputMode)
func (this *QInputDialog) SetInputMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setInputModeENS_9InputModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:98
// index:0
// Public
// QInputDialog::InputMode inputMode()
func (this *QInputDialog) InputMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9inputModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:100
// index:0
// Public
// void setLabelText(const class QString &)
func (this *QInputDialog) SetLabelText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setLabelTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:101
// index:0
// Public
// QString labelText()
func (this *QInputDialog) LabelText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9labelTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:103
// index:0
// Public
// void setOption(enum QInputDialog::InputDialogOption, _Bool)
func (this *QInputDialog) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9setOptionENS_17InputDialogOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:104
// index:0
// Public
// bool testOption(enum QInputDialog::InputDialogOption)
func (this *QInputDialog) TestOption(option int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10testOptionENS_17InputDialogOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:106
// index:0
// Public
// QInputDialog::InputDialogOptions options()
func (this *QInputDialog) Options() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:108
// index:0
// Public
// void setTextValue(const class QString &)
func (this *QInputDialog) SetTextValue(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setTextValueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:109
// index:0
// Public
// QString textValue()
func (this *QInputDialog) TextValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9textValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:111
// index:0
// Public
// void setTextEchoMode(class QLineEdit::EchoMode)
func (this *QInputDialog) SetTextEchoMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15setTextEchoModeEN9QLineEdit8EchoModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:112
// index:0
// Public
// QLineEdit::EchoMode textEchoMode()
func (this *QInputDialog) TextEchoMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog12textEchoModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:114
// index:0
// Public
// void setComboBoxEditable(_Bool)
func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19setComboBoxEditableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:115
// index:0
// Public
// bool isComboBoxEditable()
func (this *QInputDialog) IsComboBoxEditable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog18isComboBoxEditableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:117
// index:0
// Public
// void setComboBoxItems(const class QStringList &)
func (this *QInputDialog) SetComboBoxItems(items *qtcore.QStringList) {
	var convArg0 = items.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setComboBoxItemsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:118
// index:0
// Public
// QStringList comboBoxItems()
func (this *QInputDialog) ComboBoxItems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13comboBoxItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:120
// index:0
// Public
// void setIntValue(int)
func (this *QInputDialog) SetIntValue(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog11setIntValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:121
// index:0
// Public
// int intValue()
func (this *QInputDialog) IntValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog8intValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:123
// index:0
// Public
// void setIntMinimum(int)
func (this *QInputDialog) SetIntMinimum(min int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setIntMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:124
// index:0
// Public
// int intMinimum()
func (this *QInputDialog) IntMinimum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10intMinimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:126
// index:0
// Public
// void setIntMaximum(int)
func (this *QInputDialog) SetIntMaximum(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setIntMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:127
// index:0
// Public
// int intMaximum()
func (this *QInputDialog) IntMaximum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10intMaximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:129
// index:0
// Public
// void setIntRange(int, int)
func (this *QInputDialog) SetIntRange(min int, max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog11setIntRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:131
// index:0
// Public
// void setIntStep(int)
func (this *QInputDialog) SetIntStep(step int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setIntStepEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:132
// index:0
// Public
// int intStep()
func (this *QInputDialog) IntStep() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog7intStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:134
// index:0
// Public
// void setDoubleValue(double)
func (this *QInputDialog) SetDoubleValue(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleValueEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:135
// index:0
// Public
// double doubleValue()
func (this *QInputDialog) DoubleValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog11doubleValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:137
// index:0
// Public
// void setDoubleMinimum(double)
func (this *QInputDialog) SetDoubleMinimum(min float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMinimumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:138
// index:0
// Public
// double doubleMinimum()
func (this *QInputDialog) DoubleMinimum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMinimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:140
// index:0
// Public
// void setDoubleMaximum(double)
func (this *QInputDialog) SetDoubleMaximum(max float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMaximumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:141
// index:0
// Public
// double doubleMaximum()
func (this *QInputDialog) DoubleMaximum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMaximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:143
// index:0
// Public
// void setDoubleRange(double, double)
func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleRangeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:145
// index:0
// Public
// void setDoubleDecimals(int)
func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog17setDoubleDecimalsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:146
// index:0
// Public
// int doubleDecimals()
func (this *QInputDialog) DoubleDecimals() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog14doubleDecimalsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:148
// index:0
// Public
// void setOkButtonText(const class QString &)
func (this *QInputDialog) SetOkButtonText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15setOkButtonTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:149
// index:0
// Public
// QString okButtonText()
func (this *QInputDialog) OkButtonText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog12okButtonTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:151
// index:0
// Public
// void setCancelButtonText(const class QString &)
func (this *QInputDialog) SetCancelButtonText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19setCancelButtonTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:152
// index:0
// Public
// QString cancelButtonText()
func (this *QInputDialog) CancelButtonText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog16cancelButtonTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:155
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QInputDialog) Open(receiver unsafe.Pointer, member string) {
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), receiver, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:157
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QInputDialog) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:158
// index:0
// Public virtual
// QSize sizeHint()
func (this *QInputDialog) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:160
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QInputDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:196
// index:0
// Public
// void setDoubleStep(double)
func (this *QInputDialog) SetDoubleStep(step float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setDoubleStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:197
// index:0
// Public
// double doubleStep()
func (this *QInputDialog) DoubleStep() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10doubleStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:201
// index:0
// Public
// void textValueChanged(const class QString &)
func (this *QInputDialog) TextValueChanged(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16textValueChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:202
// index:0
// Public
// void textValueSelected(const class QString &)
func (this *QInputDialog) TextValueSelected(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog17textValueSelectedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:203
// index:0
// Public
// void intValueChanged(int)
func (this *QInputDialog) IntValueChanged(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15intValueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:204
// index:0
// Public
// void intValueSelected(int)
func (this *QInputDialog) IntValueSelected(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16intValueSelectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:205
// index:0
// Public
// void doubleValueChanged(double)
func (this *QInputDialog) DoubleValueChanged(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog18doubleValueChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:206
// index:0
// Public
// void doubleValueSelected(double)
func (this *QInputDialog) DoubleValueSelected(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19doubleValueSelectedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:209
// index:0
// Public virtual
// void done(int)
func (this *QInputDialog) Done(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

//  body block end
