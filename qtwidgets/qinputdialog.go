package qtwidgets

// /usr/include/qt/QtWidgets/qinputdialog.h
// #include <qinputdialog.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QInputDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:94
// index:0
// Public
// void QInputDialog(QWidget *, Qt::WindowFlags)
func NewQInputDialog(parent *QWidget /*777 QWidget **/, flags int) *QInputDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQInputDialogFromPointer(cthis)
	return gothis
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setInputModeENS_9InputModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:98
// index:0
// Public
// QInputDialog::InputMode inputMode()
func (this *QInputDialog) InputMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9inputModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:100
// index:0
// Public
// void setLabelText(const QString &)
func (this *QInputDialog) SetLabelText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setLabelTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:101
// index:0
// Public
// QString labelText()
func (this *QInputDialog) LabelText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9labelTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:103
// index:0
// Public
// void setOption(enum QInputDialog::InputDialogOption, _Bool)
func (this *QInputDialog) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9setOptionENS_17InputDialogOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:104
// index:0
// Public
// bool testOption(enum QInputDialog::InputDialogOption)
func (this *QInputDialog) TestOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10testOptionENS_17InputDialogOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qinputdialog.h:106
// index:0
// Public
// QInputDialog::InputDialogOptions options()
func (this *QInputDialog) Options() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:108
// index:0
// Public
// void setTextValue(const QString &)
func (this *QInputDialog) SetTextValue(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog12setTextValueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:109
// index:0
// Public
// QString textValue()
func (this *QInputDialog) TextValue() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog9textValueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:111
// index:0
// Public
// void setTextEchoMode(QLineEdit::EchoMode)
func (this *QInputDialog) SetTextEchoMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15setTextEchoModeEN9QLineEdit8EchoModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:112
// index:0
// Public
// QLineEdit::EchoMode textEchoMode()
func (this *QInputDialog) TextEchoMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog12textEchoModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:114
// index:0
// Public
// void setComboBoxEditable(_Bool)
func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19setComboBoxEditableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:115
// index:0
// Public
// bool isComboBoxEditable()
func (this *QInputDialog) IsComboBoxEditable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog18isComboBoxEditableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qinputdialog.h:117
// index:0
// Public
// void setComboBoxItems(const QStringList &)
func (this *QInputDialog) SetComboBoxItems(items *qtcore.QStringList) {
	var convArg0 = items.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setComboBoxItemsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:120
// index:0
// Public
// void setIntValue(int)
func (this *QInputDialog) SetIntValue(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog11setIntValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:121
// index:0
// Public
// int intValue()
func (this *QInputDialog) IntValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog8intValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:123
// index:0
// Public
// void setIntMinimum(int)
func (this *QInputDialog) SetIntMinimum(min int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setIntMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:124
// index:0
// Public
// int intMinimum()
func (this *QInputDialog) IntMinimum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10intMinimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:126
// index:0
// Public
// void setIntMaximum(int)
func (this *QInputDialog) SetIntMaximum(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setIntMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:127
// index:0
// Public
// int intMaximum()
func (this *QInputDialog) IntMaximum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10intMaximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:129
// index:0
// Public
// void setIntRange(int, int)
func (this *QInputDialog) SetIntRange(min int, max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog11setIntRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:131
// index:0
// Public
// void setIntStep(int)
func (this *QInputDialog) SetIntStep(step int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setIntStepEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:132
// index:0
// Public
// int intStep()
func (this *QInputDialog) IntStep() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog7intStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:134
// index:0
// Public
// void setDoubleValue(double)
func (this *QInputDialog) SetDoubleValue(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleValueEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:135
// index:0
// Public
// double doubleValue()
func (this *QInputDialog) DoubleValue() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog11doubleValueEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:137
// index:0
// Public
// void setDoubleMinimum(double)
func (this *QInputDialog) SetDoubleMinimum(min float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMinimumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:138
// index:0
// Public
// double doubleMinimum()
func (this *QInputDialog) DoubleMinimum() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMinimumEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:140
// index:0
// Public
// void setDoubleMaximum(double)
func (this *QInputDialog) SetDoubleMaximum(max float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16setDoubleMaximumEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:141
// index:0
// Public
// double doubleMaximum()
func (this *QInputDialog) DoubleMaximum() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog13doubleMaximumEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:143
// index:0
// Public
// void setDoubleRange(double, double)
func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog14setDoubleRangeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:145
// index:0
// Public
// void setDoubleDecimals(int)
func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog17setDoubleDecimalsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), decimals)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:146
// index:0
// Public
// int doubleDecimals()
func (this *QInputDialog) DoubleDecimals() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog14doubleDecimalsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:148
// index:0
// Public
// void setOkButtonText(const QString &)
func (this *QInputDialog) SetOkButtonText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15setOkButtonTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:149
// index:0
// Public
// QString okButtonText()
func (this *QInputDialog) OkButtonText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog12okButtonTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:151
// index:0
// Public
// void setCancelButtonText(const QString &)
func (this *QInputDialog) SetCancelButtonText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19setCancelButtonTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:152
// index:0
// Public
// QString cancelButtonText()
func (this *QInputDialog) CancelButtonText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog16cancelButtonTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:155
// index:0
// Public
// void open(QObject *, const char *)
func (this *QInputDialog) Open(receiver *qtcore.QObject /*777 QObject **/, member string) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:157
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QInputDialog) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK12QInputDialog15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:158
// index:0
// Public virtual
// QSize sizeHint()
func (this *QInputDialog) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK12QInputDialog8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qinputdialog.h:160
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QInputDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:176
// index:0
// Public static
// int getInt(QWidget *, const QString &, const QString &, int, int, int, int, _Bool *, Qt::WindowFlags)
func (this *QInputDialog) GetInt(parent *QWidget /*777 QWidget **/, title *qtcore.QString, label *qtcore.QString, value int, minValue int, maxValue int, step int, ok unsafe.Pointer /*666*/, flags int) int {
	var convArg0 = parent.GetCthis()
	var convArg1 = title.GetCthis()
	var convArg2 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog6getIntEP7QWidgetRK7QStringS4_iiiiPb6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, step, &ok, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QInputDialog_GetInt(parent *QWidget /*777 QWidget **/, title *qtcore.QString, label *qtcore.QString, value int, minValue int, maxValue int, step int, ok unsafe.Pointer /*666*/, flags int) int {
	var nilthis *QInputDialog
	rv := nilthis.GetInt(parent, title, label, value, minValue, maxValue, step, ok, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:179
// index:0
// Public static
// double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, _Bool *, Qt::WindowFlags)
func (this *QInputDialog) GetDouble(parent *QWidget /*777 QWidget **/, title *qtcore.QString, label *qtcore.QString, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int) float64 {
	var convArg0 = parent.GetCthis()
	var convArg1 = title.GetCthis()
	var convArg2 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, &ok, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QInputDialog_GetDouble(parent *QWidget /*777 QWidget **/, title *qtcore.QString, label *qtcore.QString, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int) float64 {
	var nilthis *QInputDialog
	rv := nilthis.GetDouble(parent, title, label, value, minValue, maxValue, decimals, ok, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:183
// index:1
// Public static
// double getDouble(QWidget *, const QString &, const QString &, double, double, double, int, _Bool *, Qt::WindowFlags, double)
func (this *QInputDialog) GetDouble_1(parent *QWidget /*777 QWidget **/, title *qtcore.QString, label *qtcore.QString, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int, step float64) float64 {
	var convArg0 = parent.GetCthis()
	var convArg1 = title.GetCthis()
	var convArg2 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog9getDoubleEP7QWidgetRK7QStringS4_dddiPb6QFlagsIN2Qt10WindowTypeEEd", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, value, minValue, maxValue, decimals, &ok, flags, step)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QInputDialog_GetDouble_1(parent *QWidget /*777 QWidget **/, title *qtcore.QString, label *qtcore.QString, value float64, minValue float64, maxValue float64, decimals int, ok unsafe.Pointer /*666*/, flags int, step float64) float64 {
	var nilthis *QInputDialog
	rv := nilthis.GetDouble_1(parent, title, label, value, minValue, maxValue, decimals, ok, flags, step)
	return rv
}

// /usr/include/qt/QtWidgets/qinputdialog.h:196
// index:0
// Public
// void setDoubleStep(double)
func (this *QInputDialog) SetDoubleStep(step float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog13setDoubleStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:197
// index:0
// Public
// double doubleStep()
func (this *QInputDialog) DoubleStep() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputDialog10doubleStepEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qinputdialog.h:201
// index:0
// Public
// void textValueChanged(const QString &)
func (this *QInputDialog) TextValueChanged(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16textValueChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:202
// index:0
// Public
// void textValueSelected(const QString &)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog15intValueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:204
// index:0
// Public
// void intValueSelected(int)
func (this *QInputDialog) IntValueSelected(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog16intValueSelectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:205
// index:0
// Public
// void doubleValueChanged(double)
func (this *QInputDialog) DoubleValueChanged(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog18doubleValueChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:206
// index:0
// Public
// void doubleValueSelected(double)
func (this *QInputDialog) DoubleValueSelected(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog19doubleValueSelectedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qinputdialog.h:209
// index:0
// Public virtual
// void done(int)
func (this *QInputDialog) Done(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), result)
	gopp.ErrPrint(err, rv)
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
