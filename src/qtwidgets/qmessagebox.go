//  header block begin
// /usr/include/qt/QtWidgets/qmessagebox.h
// #include <qmessagebox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
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
type QMessageBox struct {
	*QDialog
}

func (this *QMessageBox) GetCthis() unsafe.Pointer {
	return this.QDialog.GetCthis()
}
func NewQMessageBoxFromPointer(cthis unsafe.Pointer) *QMessageBox {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QMessageBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qmessagebox.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMessageBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:135
// index:0
// Public
// void QMessageBox(class QWidget *)
func NewQMessageBox(parent unsafe.Pointer) *QMessageBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:139
// index:0
// Public virtual
// void ~QMessageBox()
func DeleteQMessageBox(*QMessageBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:141
// index:0
// Public
// void addButton(class QAbstractButton *, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton(button unsafe.Pointer, role int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:142
// index:1
// Public
// QPushButton * addButton(const class QString &, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton_1(text *qtcore.QString, role int) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonERK7QStringNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:143
// index:2
// Public
// QPushButton * addButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) AddButton_2(button int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:144
// index:0
// Public
// void removeButton(class QAbstractButton *)
func (this *QMessageBox) RemoveButton(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:147
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QMessageBox) Open(receiver unsafe.Pointer, member string) {
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), receiver, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:149
// index:0
// Public
// QList<QAbstractButton *> buttons()
func (this *QMessageBox) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:150
// index:0
// Public
// QMessageBox::ButtonRole buttonRole(class QAbstractButton *)
func (this *QMessageBox) ButtonRole(button unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10buttonRoleEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:153
// index:0
// Public
// QMessageBox::StandardButtons standardButtons()
func (this *QMessageBox) StandardButtons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox15standardButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:154
// index:0
// Public
// QMessageBox::StandardButton standardButton(class QAbstractButton *)
func (this *QMessageBox) StandardButton(button unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox14standardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:155
// index:0
// Public
// QAbstractButton * button(enum QMessageBox::StandardButton)
func (this *QMessageBox) Button(which int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox6buttonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:157
// index:0
// Public
// QPushButton * defaultButton()
func (this *QMessageBox) DefaultButton() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox13defaultButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:158
// index:0
// Public
// void setDefaultButton(class QPushButton *)
func (this *QMessageBox) SetDefaultButton(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonEP11QPushButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:159
// index:1
// Public
// void setDefaultButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetDefaultButton_1(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:161
// index:0
// Public
// QAbstractButton * escapeButton()
func (this *QMessageBox) EscapeButton() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox12escapeButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:162
// index:0
// Public
// void setEscapeButton(class QAbstractButton *)
func (this *QMessageBox) SetEscapeButton(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:163
// index:1
// Public
// void setEscapeButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetEscapeButton_1(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:165
// index:0
// Public
// QAbstractButton * clickedButton()
func (this *QMessageBox) ClickedButton() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox13clickedButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:167
// index:0
// Public
// QString text()
func (this *QMessageBox) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:168
// index:0
// Public
// void setText(const class QString &)
func (this *QMessageBox) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:170
// index:0
// Public
// QMessageBox::Icon icon()
func (this *QMessageBox) Icon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:171
// index:0
// Public
// void setIcon(enum QMessageBox::Icon)
func (this *QMessageBox) SetIcon(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7setIconENS_4IconE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:173
// index:0
// Public
// QPixmap iconPixmap()
func (this *QMessageBox) IconPixmap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10iconPixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:174
// index:0
// Public
// void setIconPixmap(const class QPixmap &)
func (this *QMessageBox) SetIconPixmap(pixmap *qtgui.QPixmap) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setIconPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:176
// index:0
// Public
// Qt::TextFormat textFormat()
func (this *QMessageBox) TextFormat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10textFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:177
// index:0
// Public
// void setTextFormat(Qt::TextFormat)
func (this *QMessageBox) SetTextFormat(format int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:180
// index:0
// Public
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QMessageBox) TextInteractionFlags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox20textInteractionFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:182
// index:0
// Public
// void setCheckBox(class QCheckBox *)
func (this *QMessageBox) SetCheckBox(cb unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11setCheckBoxEP9QCheckBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cb)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:183
// index:0
// Public
// QCheckBox * checkBox()
func (this *QMessageBox) CheckBox() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox8checkBoxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:197
// index:0
// Public static
// void about(class QWidget *, const class QString &, const class QString &)
func (this *QMessageBox) About(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_", ffiqt.FFI_TYPE_POINTER, parent, title, text)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_About(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString) {
	var nilthis *QMessageBox
	nilthis.About(parent, title, text)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:198
// index:0
// Public static
// void aboutQt(class QWidget *, const class QString &)
func (this *QMessageBox) AboutQt(parent unsafe.Pointer, title *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, parent, title)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_AboutQt(parent unsafe.Pointer, title *qtcore.QString) {
	var nilthis *QMessageBox
	nilthis.AboutQt(parent, title)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:207
// index:0
// Public static
// int information(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Information(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Information(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) {
	var nilthis *QMessageBox
	nilthis.Information(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:210
// index:1
// Public static
// int information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Information_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Information_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) {
	var nilthis *QMessageBox
	nilthis.Information_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:217
// index:2
// Public static inline
// QMessageBox::StandardButton information(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Information_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Information_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) {
	var nilthis *QMessageBox
	nilthis.Information_2(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:222
// index:0
// Public static
// int question(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Question(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Question(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) {
	var nilthis *QMessageBox
	nilthis.Question(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:225
// index:1
// Public static
// int question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Question_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Question_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) {
	var nilthis *QMessageBox
	nilthis.Question_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:232
// index:2
// Public static inline
// int question(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Question_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Question_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) {
	var nilthis *QMessageBox
	nilthis.Question_2(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:237
// index:0
// Public static
// int warning(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Warning(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Warning(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) {
	var nilthis *QMessageBox
	nilthis.Warning(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:240
// index:1
// Public static
// int warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Warning_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Warning_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) {
	var nilthis *QMessageBox
	nilthis.Warning_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:247
// index:2
// Public static inline
// int warning(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Warning_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Warning_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) {
	var nilthis *QMessageBox
	nilthis.Warning_2(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:252
// index:0
// Public static
// int critical(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Critical(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Critical(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) {
	var nilthis *QMessageBox
	nilthis.Critical(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:255
// index:1
// Public static
// int critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Critical_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Critical_1(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) {
	var nilthis *QMessageBox
	nilthis.Critical_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:262
// index:2
// Public static inline
// int critical(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Critical_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_Critical_2(parent unsafe.Pointer, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) {
	var nilthis *QMessageBox
	nilthis.Critical_2(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:267
// index:0
// Public
// QString buttonText(int)
func (this *QMessageBox) ButtonText(button int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10buttonTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:268
// index:0
// Public
// void setButtonText(int, const class QString &)
func (this *QMessageBox) SetButtonText(button int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setButtonTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:270
// index:0
// Public
// QString informativeText()
func (this *QMessageBox) InformativeText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox15informativeTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:271
// index:0
// Public
// void setInformativeText(const class QString &)
func (this *QMessageBox) SetInformativeText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox18setInformativeTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:274
// index:0
// Public
// QString detailedText()
func (this *QMessageBox) DetailedText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox12detailedTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:275
// index:0
// Public
// void setDetailedText(const class QString &)
func (this *QMessageBox) SetDetailedText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setDetailedTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:278
// index:0
// Public
// void setWindowTitle(const class QString &)
func (this *QMessageBox) SetWindowTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox14setWindowTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:279
// index:0
// Public
// void setWindowModality(Qt::WindowModality)
func (this *QMessageBox) SetWindowModality(windowModality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox17setWindowModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:282
// index:0
// Public static
// QPixmap standardIcon(enum QMessageBox::Icon)
func (this *QMessageBox) StandardIcon(icon int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox12standardIconENS_4IconE", ffiqt.FFI_TYPE_POINTER, icon)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMessageBox_StandardIcon(icon int) {
	var nilthis *QMessageBox
	nilthis.StandardIcon(icon)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:285
// index:0
// Public
// void buttonClicked(class QAbstractButton *)
func (this *QMessageBox) ButtonClicked(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13buttonClickedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:293
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QMessageBox) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:294
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMessageBox) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:295
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QMessageBox) ShowEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:296
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QMessageBox) CloseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:297
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QMessageBox) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:298
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QMessageBox) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
