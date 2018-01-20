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

// /usr/include/qt/QtWidgets/qmessagebox.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMessageBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:135
// index:0
// void QMessageBox(class QWidget *)
func NewQMessageBox(parent unsafe.Pointer) *QMessageBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(cthis)
	return gothis
}
func NewQMessageBoxFromPointer(cthis unsafe.Pointer) *QMessageBox {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QMessageBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qmessagebox.h:136
// index:1
// void QMessageBox(enum QMessageBox::Icon, const class QString &, const class QString &, QMessageBox::StandardButtons, class QWidget *, Qt::WindowFlags)
func NewQMessageBox_1(icon int, title unsafe.Pointer, text unsafe.Pointer, buttons int, parent unsafe.Pointer, flags int) *QMessageBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxC2ENS_4IconERK7QStringS3_6QFlagsINS_14StandardButtonEEP7QWidgetS4_IN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, &icon, title, text, buttons, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:202
// index:2
// void QMessageBox(const class QString &, const class QString &, enum QMessageBox::Icon, int, int, int, class QWidget *, Qt::WindowFlags)
func NewQMessageBox_2(title unsafe.Pointer, text unsafe.Pointer, icon int, button0 int, button1 int, button2 int, parent unsafe.Pointer, f int) *QMessageBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxC2ERK7QStringS2_NS_4IconEiiiP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, title, text, &icon, &button0, &button1, &button2, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:139
// index:0
// virtual
// void ~QMessageBox()
func DeleteQMessageBox(*QMessageBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:141
// index:0
// void addButton(class QAbstractButton *, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton(button unsafe.Pointer, role int) {
	// 0: (, button QAbstractButton *, role QMessageBox::ButtonRole), (button, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:142
// index:1
// QPushButton * addButton(const class QString &, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton_1(text unsafe.Pointer, role int) {
	// 1: (, text const QString &, role QMessageBox::ButtonRole), (text, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonERK7QStringNS_10ButtonRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:143
// index:2
// QPushButton * addButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) AddButton_2(button int) {
	// 2: (, button QMessageBox::StandardButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:144
// index:0
// void removeButton(class QAbstractButton *)
func (this *QMessageBox) RemoveButton(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:147
// index:0
// void open(class QObject *, const char *)
func (this *QMessageBox) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, receiver QObject *, member const char *), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:149
// index:0
// QList<QAbstractButton *> buttons()
func (this *QMessageBox) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:150
// index:0
// QMessageBox::ButtonRole buttonRole(class QAbstractButton *)
func (this *QMessageBox) ButtonRole(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10buttonRoleEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:152
// index:0
// void setStandardButtons(QMessageBox::StandardButtons)
func (this *QMessageBox) SetStandardButtons(buttons int) {
	// 0: (, QFlags<QMessageBox::StandardButton> buttons), (buttons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox18setStandardButtonsE6QFlagsINS_14StandardButtonEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:153
// index:0
// QMessageBox::StandardButtons standardButtons()
func (this *QMessageBox) StandardButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox15standardButtonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:154
// index:0
// QMessageBox::StandardButton standardButton(class QAbstractButton *)
func (this *QMessageBox) StandardButton(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox14standardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:155
// index:0
// QAbstractButton * button(enum QMessageBox::StandardButton)
func (this *QMessageBox) Button(which int) {
	// 0: (, which QMessageBox::StandardButton), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox6buttonENS_14StandardButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:157
// index:0
// QPushButton * defaultButton()
func (this *QMessageBox) DefaultButton() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox13defaultButtonEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:158
// index:0
// void setDefaultButton(class QPushButton *)
func (this *QMessageBox) SetDefaultButton(button unsafe.Pointer) {
	// 0: (, button QPushButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonEP11QPushButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:159
// index:1
// void setDefaultButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetDefaultButton_1(button int) {
	// 1: (, button QMessageBox::StandardButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:161
// index:0
// QAbstractButton * escapeButton()
func (this *QMessageBox) EscapeButton() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox12escapeButtonEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:162
// index:0
// void setEscapeButton(class QAbstractButton *)
func (this *QMessageBox) SetEscapeButton(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:163
// index:1
// void setEscapeButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetEscapeButton_1(button int) {
	// 1: (, button QMessageBox::StandardButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:165
// index:0
// QAbstractButton * clickedButton()
func (this *QMessageBox) ClickedButton() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox13clickedButtonEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:167
// index:0
// QString text()
func (this *QMessageBox) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:168
// index:0
// void setText(const class QString &)
func (this *QMessageBox) SetText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:170
// index:0
// QMessageBox::Icon icon()
func (this *QMessageBox) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox4iconEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:171
// index:0
// void setIcon(enum QMessageBox::Icon)
func (this *QMessageBox) SetIcon(arg0 int) {
	// 0: (, QMessageBox::Icon arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7setIconENS_4IconE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:173
// index:0
// QPixmap iconPixmap()
func (this *QMessageBox) IconPixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10iconPixmapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:174
// index:0
// void setIconPixmap(const class QPixmap &)
func (this *QMessageBox) SetIconPixmap(pixmap unsafe.Pointer) {
	// 0: (, pixmap const QPixmap &), (pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setIconPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:176
// index:0
// Qt::TextFormat textFormat()
func (this *QMessageBox) TextFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10textFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:177
// index:0
// void setTextFormat(Qt::TextFormat)
func (this *QMessageBox) SetTextFormat(format int) {
	// 0: (, format Qt::TextFormat), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:179
// index:0
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QMessageBox) SetTextInteractionFlags(flags int) {
	// 0: (, QFlags<Qt::TextInteractionFlag> flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:180
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QMessageBox) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:182
// index:0
// void setCheckBox(class QCheckBox *)
func (this *QMessageBox) SetCheckBox(cb unsafe.Pointer) {
	// 0: (, cb QCheckBox *), (cb)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11setCheckBoxEP9QCheckBox", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cb)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:183
// index:0
// QCheckBox * checkBox()
func (this *QMessageBox) CheckBox() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox8checkBoxEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:185
// index:0
// static
// QMessageBox::StandardButton information(class QWidget *, const class QString &, const class QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Information(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Information(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	var nilthis *QMessageBox
	nilthis.Information(parent, title, text, buttons, defaultButton)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:207
// index:1
// static
// int information(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Information_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Information_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	var nilthis *QMessageBox
	nilthis.Information_1(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:210
// index:2
// static
// int information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Information_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Information_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	var nilthis *QMessageBox
	nilthis.Information_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:217
// index:3
// static inline
// QMessageBox::StandardButton information(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Information_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Information_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	var nilthis *QMessageBox
	nilthis.Information_3(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:188
// index:0
// static
// QMessageBox::StandardButton question(class QWidget *, const class QString &, const class QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Question(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Question(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	var nilthis *QMessageBox
	nilthis.Question(parent, title, text, buttons, defaultButton)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:222
// index:1
// static
// int question(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Question_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Question_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	var nilthis *QMessageBox
	nilthis.Question_1(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:225
// index:2
// static
// int question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Question_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Question_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	var nilthis *QMessageBox
	nilthis.Question_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:232
// index:3
// static inline
// int question(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Question_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Question_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	var nilthis *QMessageBox
	nilthis.Question_3(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:191
// index:0
// static
// QMessageBox::StandardButton warning(class QWidget *, const class QString &, const class QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Warning(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Warning(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	var nilthis *QMessageBox
	nilthis.Warning(parent, title, text, buttons, defaultButton)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:237
// index:1
// static
// int warning(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Warning_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Warning_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	var nilthis *QMessageBox
	nilthis.Warning_1(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:240
// index:2
// static
// int warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Warning_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Warning_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	var nilthis *QMessageBox
	nilthis.Warning_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:247
// index:3
// static inline
// int warning(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Warning_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Warning_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	var nilthis *QMessageBox
	nilthis.Warning_3(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:194
// index:0
// static
// QMessageBox::StandardButton critical(class QWidget *, const class QString &, const class QString &, QMessageBox::StandardButtons, enum QMessageBox::StandardButton)
func (this *QMessageBox) Critical(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_6QFlagsINS_14StandardButtonEES6_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Critical(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, buttons int, defaultButton int) {
	// 0: (parent QWidget *, title const QString &, text const QString &, QFlags<QMessageBox::StandardButton> buttons, defaultButton QMessageBox::StandardButton), (parent, title, text, buttons, defaultButton)
	var nilthis *QMessageBox
	nilthis.Critical(parent, title, text, buttons, defaultButton)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:252
// index:1
// static
// int critical(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Critical_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Critical_1(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int, button2 int) {
	// 1: (parent QWidget *, title const QString &, text const QString &, button0 int, button1 int, button2 int), (parent, title, text, button0, button1, button2)
	var nilthis *QMessageBox
	nilthis.Critical_1(parent, title, text, button0, button1, button2)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:255
// index:2
// static
// int critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Critical_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Critical_2(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0Text unsafe.Pointer, button1Text unsafe.Pointer, button2Text unsafe.Pointer, defaultButtonNumber int, escapeButtonNumber int) {
	// 2: (parent QWidget *, title const QString &, text const QString &, button0Text const QString &, button1Text const QString &, button2Text const QString &, defaultButtonNumber int, escapeButtonNumber int), (parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	var nilthis *QMessageBox
	nilthis.Critical_2(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:262
// index:3
// static inline
// int critical(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Critical_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_Critical_3(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer, button0 int, button1 int) {
	// 3: (parent QWidget *, title const QString &, text const QString &, button0 QMessageBox::StandardButton, button1 QMessageBox::StandardButton), (parent, title, text, button0, button1)
	var nilthis *QMessageBox
	nilthis.Critical_3(parent, title, text, button0, button1)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:197
// index:0
// static
// void about(class QWidget *, const class QString &, const class QString &)
func (this *QMessageBox) About(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer) {
	// 0: (parent QWidget *, title const QString &, text const QString &), (parent, title, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_About(parent unsafe.Pointer, title unsafe.Pointer, text unsafe.Pointer) {
	// 0: (parent QWidget *, title const QString &, text const QString &), (parent, title, text)
	var nilthis *QMessageBox
	nilthis.About(parent, title, text)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:198
// index:0
// static
// void aboutQt(class QWidget *, const class QString &)
func (this *QMessageBox) AboutQt(parent unsafe.Pointer, title unsafe.Pointer) {
	// 0: (parent QWidget *, title const QString &), (parent, title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_AboutQt(parent unsafe.Pointer, title unsafe.Pointer) {
	// 0: (parent QWidget *, title const QString &), (parent, title)
	var nilthis *QMessageBox
	nilthis.AboutQt(parent, title)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:267
// index:0
// QString buttonText(int)
func (this *QMessageBox) ButtonText(button int) {
	// 0: (, button int), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10buttonTextEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:268
// index:0
// void setButtonText(int, const class QString &)
func (this *QMessageBox) SetButtonText(button int, text unsafe.Pointer) {
	// 0: (, button int, text const QString &), (&button, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setButtonTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:270
// index:0
// QString informativeText()
func (this *QMessageBox) InformativeText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox15informativeTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:271
// index:0
// void setInformativeText(const class QString &)
func (this *QMessageBox) SetInformativeText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox18setInformativeTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:274
// index:0
// QString detailedText()
func (this *QMessageBox) DetailedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox12detailedTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:275
// index:0
// void setDetailedText(const class QString &)
func (this *QMessageBox) SetDetailedText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setDetailedTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:278
// index:0
// void setWindowTitle(const class QString &)
func (this *QMessageBox) SetWindowTitle(title unsafe.Pointer) {
	// 0: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox14setWindowTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:279
// index:0
// void setWindowModality(Qt::WindowModality)
func (this *QMessageBox) SetWindowModality(windowModality int) {
	// 0: (, windowModality Qt::WindowModality), (&windowModality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox17setWindowModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:282
// index:0
// static
// QPixmap standardIcon(enum QMessageBox::Icon)
func (this *QMessageBox) StandardIcon(icon int) {
	// 0: (icon QMessageBox::Icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox12standardIconENS_4IconE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_StandardIcon(icon int) {
	// 0: (icon QMessageBox::Icon), (icon)
	var nilthis *QMessageBox
	nilthis.StandardIcon(icon)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:285
// index:0
// void buttonClicked(class QAbstractButton *)
func (this *QMessageBox) ButtonClicked(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13buttonClickedEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:293
// index:0
// virtual
// bool event(class QEvent *)
func (this *QMessageBox) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:294
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMessageBox) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:295
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QMessageBox) ShowEvent(event unsafe.Pointer) {
	// 0: (, event QShowEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:296
// index:0
// virtual
// void closeEvent(class QCloseEvent *)
func (this *QMessageBox) CloseEvent(event unsafe.Pointer) {
	// 0: (, event QCloseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:297
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QMessageBox) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:298
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QMessageBox) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
