//  header block begin
// /usr/include/qt/QtWidgets/qinputdialog.h
// #include <qinputdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 96
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

// /usr/include/qt/QtWidgets/qinputdialog.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QInputDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:94
// index:0
// void QInputDialog(class QWidget *, Qt::WindowFlags)
func NewQInputDialog(parent unsafe.Pointer, flags int) *QInputDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQInputDialogFromPointer(cthis)
	return gothis
}
func NewQInputDialogFromPointer(cthis unsafe.Pointer) *QInputDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QInputDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qinputdialog.h:95
// index:0
// virtual
// void ~QInputDialog()
func DeleteQInputDialog(*QInputDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:97
// index:0
// void setInputMode(enum QInputDialog::InputMode)
func (this *QInputDialog) SetInputMode(mode int) {
	// 0: (, mode QInputDialog::InputMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setInputModeENS_9InputModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:98
// index:0
// QInputDialog::InputMode inputMode()
func (this *QInputDialog) InputMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9inputModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:100
// index:0
// void setLabelText(const class QString &)
func (this *QInputDialog) SetLabelText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setLabelTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:101
// index:0
// QString labelText()
func (this *QInputDialog) LabelText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9labelTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:103
// index:0
// void setOption(enum QInputDialog::InputDialogOption, _Bool)
func (this *QInputDialog) SetOption(option int, on bool) {
	// 0: (, option QInputDialog::InputDialogOption, on bool), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9setOptionENS_17InputDialogOptionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:104
// index:0
// bool testOption(enum QInputDialog::InputDialogOption)
func (this *QInputDialog) TestOption(option int) {
	// 0: (, option QInputDialog::InputDialogOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10testOptionENS_17InputDialogOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:105
// index:0
// void setOptions(QInputDialog::InputDialogOptions)
func (this *QInputDialog) SetOptions(options int) {
	// 0: (, QFlags<QInputDialog::InputDialogOption> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setOptionsE6QFlagsINS_17InputDialogOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:106
// index:0
// QInputDialog::InputDialogOptions options()
func (this *QInputDialog) Options() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog7optionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:108
// index:0
// void setTextValue(const class QString &)
func (this *QInputDialog) SetTextValue(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setTextValueERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:109
// index:0
// QString textValue()
func (this *QInputDialog) TextValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9textValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:111
// index:0
// void setTextEchoMode(class QLineEdit::EchoMode)
func (this *QInputDialog) SetTextEchoMode(mode int) {
	// 0: (, mode QLineEdit::EchoMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15setTextEchoModeEN9QLineEdit8EchoModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:112
// index:0
// QLineEdit::EchoMode textEchoMode()
func (this *QInputDialog) TextEchoMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog12textEchoModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:114
// index:0
// void setComboBoxEditable(_Bool)
func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	// 0: (, editable bool), (&editable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19setComboBoxEditableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:115
// index:0
// bool isComboBoxEditable()
func (this *QInputDialog) IsComboBoxEditable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog18isComboBoxEditableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:117
// index:0
// void setComboBoxItems(const class QStringList &)
func (this *QInputDialog) SetComboBoxItems(items unsafe.Pointer) {
	// 0: (, items const QStringList &), (items)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setComboBoxItemsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), items)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:118
// index:0
// QStringList comboBoxItems()
func (this *QInputDialog) ComboBoxItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13comboBoxItemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:120
// index:0
// void setIntValue(int)
func (this *QInputDialog) SetIntValue(value int) {
	// 0: (, value int), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog11setIntValueEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:121
// index:0
// int intValue()
func (this *QInputDialog) IntValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog8intValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:123
// index:0
// void setIntMinimum(int)
func (this *QInputDialog) SetIntMinimum(min int) {
	// 0: (, min int), (&min)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setIntMinimumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:124
// index:0
// int intMinimum()
func (this *QInputDialog) IntMinimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10intMinimumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:126
// index:0
// void setIntMaximum(int)
func (this *QInputDialog) SetIntMaximum(max int) {
	// 0: (, max int), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setIntMaximumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:127
// index:0
// int intMaximum()
func (this *QInputDialog) IntMaximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10intMaximumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:129
// index:0
// void setIntRange(int, int)
func (this *QInputDialog) SetIntRange(min int, max int) {
	// 0: (, min int, max int), (&min, &max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog11setIntRangeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:131
// index:0
// void setIntStep(int)
func (this *QInputDialog) SetIntStep(step int) {
	// 0: (, step int), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setIntStepEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:132
// index:0
// int intStep()
func (this *QInputDialog) IntStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog7intStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:134
// index:0
// void setDoubleValue(double)
func (this *QInputDialog) SetDoubleValue(value float64) {
	// 0: (, value double), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleValueEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:135
// index:0
// double doubleValue()
func (this *QInputDialog) DoubleValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog11doubleValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:137
// index:0
// void setDoubleMinimum(double)
func (this *QInputDialog) SetDoubleMinimum(min float64) {
	// 0: (, min double), (&min)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMinimumEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:138
// index:0
// double doubleMinimum()
func (this *QInputDialog) DoubleMinimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMinimumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:140
// index:0
// void setDoubleMaximum(double)
func (this *QInputDialog) SetDoubleMaximum(max float64) {
	// 0: (, max double), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMaximumEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:141
// index:0
// double doubleMaximum()
func (this *QInputDialog) DoubleMaximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMaximumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:143
// index:0
// void setDoubleRange(double, double)
func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	// 0: (, min double, max double), (&min, &max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleRangeEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &min, &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:145
// index:0
// void setDoubleDecimals(int)
func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	// 0: (, decimals int), (&decimals)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog17setDoubleDecimalsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:146
// index:0
// int doubleDecimals()
func (this *QInputDialog) DoubleDecimals() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog14doubleDecimalsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:148
// index:0
// void setOkButtonText(const class QString &)
func (this *QInputDialog) SetOkButtonText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15setOkButtonTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:149
// index:0
// QString okButtonText()
func (this *QInputDialog) OkButtonText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog12okButtonTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:151
// index:0
// void setCancelButtonText(const class QString &)
func (this *QInputDialog) SetCancelButtonText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19setCancelButtonTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:152
// index:0
// QString cancelButtonText()
func (this *QInputDialog) CancelButtonText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog16cancelButtonTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:155
// index:0
// void open(class QObject *, const char *)
func (this *QInputDialog) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, receiver QObject *, member const char *), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:157
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QInputDialog) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:158
// index:0
// virtual
// QSize sizeHint()
func (this *QInputDialog) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:160
// index:0
// virtual
// void setVisible(_Bool)
func (this *QInputDialog) SetVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:162
// index:0
// static
// QString getText(class QWidget *, const class QString &, const class QString &, class QLineEdit::EchoMode, const class QString &, _Bool *, Qt::WindowFlags, Qt::InputMethodHints)
func (this *QInputDialog) GetText(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, echo int, text unsafe.Pointer, ok unsafe.Pointer, flags int, inputMethodHints int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, echo QLineEdit::EchoMode, text const QString &, ok bool *, QFlags<Qt::WindowType> flags, QFlags<Qt::InputMethodHint> inputMethodHints), (parent, title, label, echo, text, ok, flags, inputMethodHints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog7getTextEP7QWidgetRK7QStringS4_N9QLineEdit8EchoModeES4_Pb6QFlagsIN2Qt10WindowTypeEES8_INS9_15InputMethodHintEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputDialog_GetText(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, echo int, text unsafe.Pointer, ok unsafe.Pointer, flags int, inputMethodHints int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, echo QLineEdit::EchoMode, text const QString &, ok bool *, QFlags<Qt::WindowType> flags, QFlags<Qt::InputMethodHint> inputMethodHints), (parent, title, label, echo, text, ok, flags, inputMethodHints)
	var nilthis *QInputDialog
	nilthis.GetText(parent, title, label, echo, text, ok, flags, inputMethodHints)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:167
// index:0
// static
// QString getMultiLineText(class QWidget *, const class QString &, const class QString &, const class QString &, _Bool *, Qt::WindowFlags, Qt::InputMethodHints)
func (this *QInputDialog) GetMultiLineText(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, text unsafe.Pointer, ok unsafe.Pointer, flags int, inputMethodHints int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, text const QString &, ok bool *, QFlags<Qt::WindowType> flags, QFlags<Qt::InputMethodHint> inputMethodHints), (parent, title, label, text, ok, flags, inputMethodHints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16getMultiLineTextEP7QWidgetRK7QStringS4_S4_Pb6QFlagsIN2Qt10WindowTypeEES6_INS7_15InputMethodHintEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputDialog_GetMultiLineText(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, text unsafe.Pointer, ok unsafe.Pointer, flags int, inputMethodHints int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, text const QString &, ok bool *, QFlags<Qt::WindowType> flags, QFlags<Qt::InputMethodHint> inputMethodHints), (parent, title, label, text, ok, flags, inputMethodHints)
	var nilthis *QInputDialog
	nilthis.GetMultiLineText(parent, title, label, text, ok, flags, inputMethodHints)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:171
// index:0
// static
// QString getItem(class QWidget *, const class QString &, const class QString &, const class QStringList &, int, _Bool, _Bool *, Qt::WindowFlags, Qt::InputMethodHints)
func (this *QInputDialog) GetItem(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, items unsafe.Pointer, current int, editable bool, ok unsafe.Pointer, flags int, inputMethodHints int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, items const QStringList &, current int, editable bool, ok bool *, QFlags<Qt::WindowType> flags, QFlags<Qt::InputMethodHint> inputMethodHints), (parent, title, label, items, current, editable, ok, flags, inputMethodHints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog7getItemEP7QWidgetRK7QStringS4_RK11QStringListibPb6QFlagsIN2Qt10WindowTypeEES9_INSA_15InputMethodHintEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputDialog_GetItem(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, items unsafe.Pointer, current int, editable bool, ok unsafe.Pointer, flags int, inputMethodHints int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, items const QStringList &, current int, editable bool, ok bool *, QFlags<Qt::WindowType> flags, QFlags<Qt::InputMethodHint> inputMethodHints), (parent, title, label, items, current, editable, ok, flags, inputMethodHints)
	var nilthis *QInputDialog
	nilthis.GetItem(parent, title, label, items, current, editable, ok, flags, inputMethodHints)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// static
// int getInt(class QWidget *, const class QString &, const class QString &, int, int, int, int, _Bool *, Qt::WindowFlags)
func (this *QInputDialog) GetInt(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, value int, minValue int, maxValue int, step int, ok unsafe.Pointer, flags int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, value int, minValue int, maxValue int, step int, ok bool *, QFlags<Qt::WindowType> flags), (parent, title, label, value, minValue, maxValue, step, ok, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputDialog_GetInt(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, value int, minValue int, maxValue int, step int, ok unsafe.Pointer, flags int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, value int, minValue int, maxValue int, step int, ok bool *, QFlags<Qt::WindowType> flags), (parent, title, label, value, minValue, maxValue, step, ok, flags)
	var nilthis *QInputDialog
	nilthis.GetInt(parent, title, label, value, minValue, maxValue, step, ok, flags)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// static
// double getDouble(class QWidget *, const class QString &, const class QString &, double, double, double, int, _Bool *, Qt::WindowFlags)
func (this *QInputDialog) GetDouble(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer, flags int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, value double, minValue double, maxValue double, decimals int, ok bool *, QFlags<Qt::WindowType> flags), (parent, title, label, value, minValue, maxValue, decimals, ok, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputDialog_GetDouble(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer, flags int) {
	// 0: (parent QWidget *, title const QString &, label const QString &, value double, minValue double, maxValue double, decimals int, ok bool *, QFlags<Qt::WindowType> flags), (parent, title, label, value, minValue, maxValue, decimals, ok, flags)
	var nilthis *QInputDialog
	nilthis.GetDouble(parent, title, label, value, minValue, maxValue, decimals, ok, flags)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:183
// index:1
// static
// double getDouble(class QWidget *, const class QString &, const class QString &, double, double, double, int, _Bool *, Qt::WindowFlags, double)
func (this *QInputDialog) GetDouble_1(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer, flags int, step float64) {
	// 1: (parent QWidget *, title const QString &, label const QString &, value double, minValue double, maxValue double, decimals int, ok bool *, QFlags<Qt::WindowType> flags, step double), (parent, title, label, value, minValue, maxValue, decimals, ok, flags, step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEEd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QInputDialog_GetDouble_1(parent unsafe.Pointer, title unsafe.Pointer, label unsafe.Pointer, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer, flags int, step float64) {
	// 1: (parent QWidget *, title const QString &, label const QString &, value double, minValue double, maxValue double, decimals int, ok bool *, QFlags<Qt::WindowType> flags, step double), (parent, title, label, value, minValue, maxValue, decimals, ok, flags, step)
	var nilthis *QInputDialog
	nilthis.GetDouble_1(parent, title, label, value, minValue, maxValue, decimals, ok, flags, step)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:196
// index:0
// void setDoubleStep(double)
func (this *QInputDialog) SetDoubleStep(step float64) {
	// 0: (, step double), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setDoubleStepEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:197
// index:0
// double doubleStep()
func (this *QInputDialog) DoubleStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10doubleStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:201
// index:0
// void textValueChanged(const class QString &)
func (this *QInputDialog) TextValueChanged(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16textValueChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:202
// index:0
// void textValueSelected(const class QString &)
func (this *QInputDialog) TextValueSelected(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog17textValueSelectedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:203
// index:0
// void intValueChanged(int)
func (this *QInputDialog) IntValueChanged(value int) {
	// 0: (, value int), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15intValueChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:204
// index:0
// void intValueSelected(int)
func (this *QInputDialog) IntValueSelected(value int) {
	// 0: (, value int), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16intValueSelectedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:205
// index:0
// void doubleValueChanged(double)
func (this *QInputDialog) DoubleValueChanged(value float64) {
	// 0: (, value double), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog18doubleValueChangedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:206
// index:0
// void doubleValueSelected(double)
func (this *QInputDialog) DoubleValueSelected(value float64) {
	// 0: (, value double), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19doubleValueSelectedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:209
// index:0
// virtual
// void done(int)
func (this *QInputDialog) Done(result int) {
	// 0: (, result int), (&result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog4doneEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

//  body block end
