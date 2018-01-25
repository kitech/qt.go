package qtwidgets

// /usr/include/qt/QtWidgets/qmessagebox.h
// #include <qmessagebox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QMessageBox) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQMessageBoxFromPointer(cthis unsafe.Pointer) *QMessageBox {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QMessageBox{bcthis0}
}
func (*QMessageBox) NewFromPointer(cthis unsafe.Pointer) *QMessageBox {
	return NewQMessageBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMessageBox) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:135
// index:0
// Public
// void QMessageBox(class QWidget *)
func NewQMessageBox(parent *QWidget /*444 QWidget **/) *QMessageBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmessagebox.h:202
// index:1
// Public
// void QMessageBox(const class QString &, const class QString &, enum QMessageBox::Icon, int, int, int, class QWidget *, Qt::WindowFlags)
func NewQMessageBox_1(title *qtcore.QString, text *qtcore.QString, icon int, button0 int, button1 int, button2 int, parent *QWidget /*444 QWidget **/, f int) *QMessageBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = title.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg6 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBoxC2ERK7QStringS2_NS_4IconEiiiP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, icon, button0, button1, button2, convArg6, f)
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
func (this *QMessageBox) AddButton(button *QAbstractButton /*444 QAbstractButton **/, role int) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:142
// index:1
// Public
// QPushButton * addButton(const class QString &, enum QMessageBox::ButtonRole)
func (this *QMessageBox) AddButton_1(text *qtcore.QString, role int) *QPushButton /*444 QPushButton **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonERK7QStringNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:143
// index:2
// Public
// QPushButton * addButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) AddButton_2(button int) *QPushButton /*444 QPushButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9addButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:144
// index:0
// Public
// void removeButton(class QAbstractButton *)
func (this *QMessageBox) RemoveButton(button *QAbstractButton /*444 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:147
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QMessageBox) Open(receiver *qtcore.QObject /*444 QObject **/, member string) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:150
// index:0
// Public
// QMessageBox::ButtonRole buttonRole(class QAbstractButton *)
func (this *QMessageBox) ButtonRole(button *QAbstractButton /*444 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10buttonRoleEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:153
// index:0
// Public
// QMessageBox::StandardButtons standardButtons()
func (this *QMessageBox) StandardButtons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox15standardButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:154
// index:0
// Public
// QMessageBox::StandardButton standardButton(class QAbstractButton *)
func (this *QMessageBox) StandardButton(button *QAbstractButton /*444 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox14standardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:155
// index:0
// Public
// QAbstractButton * button(enum QMessageBox::StandardButton)
func (this *QMessageBox) Button(which int) *QAbstractButton /*444 QAbstractButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox6buttonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:157
// index:0
// Public
// QPushButton * defaultButton()
func (this *QMessageBox) DefaultButton() *QPushButton /*444 QPushButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox13defaultButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:158
// index:0
// Public
// void setDefaultButton(class QPushButton *)
func (this *QMessageBox) SetDefaultButton(button *QPushButton /*444 QPushButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonEP11QPushButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:159
// index:1
// Public
// void setDefaultButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetDefaultButton_1(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox16setDefaultButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:161
// index:0
// Public
// QAbstractButton * escapeButton()
func (this *QMessageBox) EscapeButton() *QAbstractButton /*444 QAbstractButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox12escapeButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:162
// index:0
// Public
// void setEscapeButton(class QAbstractButton *)
func (this *QMessageBox) SetEscapeButton(button *QAbstractButton /*444 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:163
// index:1
// Public
// void setEscapeButton(enum QMessageBox::StandardButton)
func (this *QMessageBox) SetEscapeButton_1(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox15setEscapeButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:165
// index:0
// Public
// QAbstractButton * clickedButton()
func (this *QMessageBox) ClickedButton() *QAbstractButton /*444 QAbstractButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox13clickedButtonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:167
// index:0
// Public
// QString text()
func (this *QMessageBox) Text() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QMessageBox) Icon() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:171
// index:0
// Public
// void setIcon(enum QMessageBox::Icon)
func (this *QMessageBox) SetIcon(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7setIconENS_4IconE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:173
// index:0
// Public
// QPixmap iconPixmap()
func (this *QMessageBox) IconPixmap() *qtgui.QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10iconPixmapEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QMessageBox) TextFormat() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10textFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:177
// index:0
// Public
// void setTextFormat(Qt::TextFormat)
func (this *QMessageBox) SetTextFormat(format int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:179
// index:0
// Public
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QMessageBox) SetTextInteractionFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:180
// index:0
// Public
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QMessageBox) TextInteractionFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox20textInteractionFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:182
// index:0
// Public
// void setCheckBox(class QCheckBox *)
func (this *QMessageBox) SetCheckBox(cb *QCheckBox /*444 QCheckBox **/) {
	var convArg0 = cb.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11setCheckBoxEP9QCheckBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:183
// index:0
// Public
// QCheckBox * checkBox()
func (this *QMessageBox) CheckBox() *QCheckBox /*444 QCheckBox **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox8checkBoxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCheckBoxFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:197
// index:0
// Public static
// void about(class QWidget *, const class QString &, const class QString &)
func (this *QMessageBox) About(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_", ffiqt.FFI_TYPE_POINTER, parent, title, text)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_About(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString) {
	var nilthis *QMessageBox
	nilthis.About(parent, title, text)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:198
// index:0
// Public static
// void aboutQt(class QWidget *, const class QString &)
func (this *QMessageBox) AboutQt(parent *QWidget /*444 QWidget **/, title *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7aboutQtEP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, parent, title)
	gopp.ErrPrint(err, rv)
}
func QMessageBox_AboutQt(parent *QWidget /*444 QWidget **/, title *qtcore.QString) {
	var nilthis *QMessageBox
	nilthis.AboutQt(parent, title)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:207
// index:0
// Public static
// int information(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Information(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Information(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:210
// index:1
// Public static
// int information(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Information_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Information_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:217
// index:2
// Public static inline
// QMessageBox::StandardButton information(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Information_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QMessageBox_Information_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Information_2(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:222
// index:0
// Public static
// int question(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Question(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Question(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:225
// index:1
// Public static
// int question(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Question_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Question_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:232
// index:2
// Public static inline
// int question(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Question_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Question_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Question_2(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:237
// index:0
// Public static
// int warning(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Warning(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Warning(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:240
// index:1
// Public static
// int warning(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Warning_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Warning_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:247
// index:2
// Public static inline
// int warning(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Warning_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Warning_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Warning_2(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:252
// index:0
// Public static
// int critical(class QWidget *, const class QString &, const class QString &, int, int, int)
func (this *QMessageBox) Critical(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1, button2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Critical(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int, button2 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical(parent, title, text, button0, button1, button2)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:255
// index:1
// Public static
// int critical(class QWidget *, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, int, int)
func (this *QMessageBox) Critical_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Critical_1(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0Text *qtcore.QString, button1Text *qtcore.QString, button2Text *qtcore.QString, defaultButtonNumber int, escapeButtonNumber int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical_1(parent, title, text, button0Text, button1Text, button2Text, defaultButtonNumber, escapeButtonNumber)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:262
// index:2
// Public static inline
// int critical(class QWidget *, const class QString &, const class QString &, enum QMessageBox::StandardButton, enum QMessageBox::StandardButton)
func (this *QMessageBox) Critical_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_NS_14StandardButtonES5_", ffiqt.FFI_TYPE_POINTER, parent, title, text, button0, button1)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMessageBox_Critical_2(parent *QWidget /*444 QWidget **/, title *qtcore.QString, text *qtcore.QString, button0 int, button1 int) int {
	var nilthis *QMessageBox
	rv := nilthis.Critical_2(parent, title, text, button0, button1)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:267
// index:0
// Public
// QString buttonText(int)
func (this *QMessageBox) ButtonText(button int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox10buttonTextEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmessagebox.h:268
// index:0
// Public
// void setButtonText(int, const class QString &)
func (this *QMessageBox) SetButtonText(button int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13setButtonTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:270
// index:0
// Public
// QString informativeText()
func (this *QMessageBox) InformativeText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox15informativeTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QMessageBox) DetailedText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMessageBox12detailedTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox17setWindowModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:282
// index:0
// Public static
// QPixmap standardIcon(enum QMessageBox::Icon)
func (this *QMessageBox) StandardIcon(icon int) *qtgui.QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox12standardIconENS_4IconE", ffiqt.FFI_TYPE_POINTER, icon)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QMessageBox_StandardIcon(icon int) *qtgui.QPixmap /*123*/ {
	var nilthis *QMessageBox
	rv := nilthis.StandardIcon(icon)
	return rv
}

// /usr/include/qt/QtWidgets/qmessagebox.h:285
// index:0
// Public
// void buttonClicked(class QAbstractButton *)
func (this *QMessageBox) ButtonClicked(button *QAbstractButton /*444 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13buttonClickedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:293
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QMessageBox) Event(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmessagebox.h:294
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMessageBox) ResizeEvent(event *qtgui.QResizeEvent /*444 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:295
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QMessageBox) ShowEvent(event *qtgui.QShowEvent /*444 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:296
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QMessageBox) CloseEvent(event *qtgui.QCloseEvent /*444 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:297
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QMessageBox) KeyPressEvent(event *qtgui.QKeyEvent /*444 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmessagebox.h:298
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QMessageBox) ChangeEvent(event *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMessageBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QMessageBox__Icon = int

const QMessageBox__NoIcon QMessageBox__Icon = 0
const QMessageBox__Information QMessageBox__Icon = 1
const QMessageBox__Warning QMessageBox__Icon = 2
const QMessageBox__Critical QMessageBox__Icon = 3
const QMessageBox__Question QMessageBox__Icon = 4

type QMessageBox__ButtonRole = int

const QMessageBox__InvalidRole QMessageBox__ButtonRole = 4294967295
const QMessageBox__AcceptRole QMessageBox__ButtonRole = 0
const QMessageBox__RejectRole QMessageBox__ButtonRole = 1
const QMessageBox__DestructiveRole QMessageBox__ButtonRole = 2
const QMessageBox__ActionRole QMessageBox__ButtonRole = 3
const QMessageBox__HelpRole QMessageBox__ButtonRole = 4
const QMessageBox__YesRole QMessageBox__ButtonRole = 5
const QMessageBox__NoRole QMessageBox__ButtonRole = 6
const QMessageBox__ResetRole QMessageBox__ButtonRole = 7
const QMessageBox__ApplyRole QMessageBox__ButtonRole = 8
const QMessageBox__NRoles QMessageBox__ButtonRole = 9

type QMessageBox__StandardButton = int

const QMessageBox__NoButton QMessageBox__StandardButton = 0
const QMessageBox__Ok QMessageBox__StandardButton = 1024
const QMessageBox__Save QMessageBox__StandardButton = 2048
const QMessageBox__SaveAll QMessageBox__StandardButton = 4096
const QMessageBox__Open QMessageBox__StandardButton = 8192
const QMessageBox__Yes QMessageBox__StandardButton = 16384
const QMessageBox__YesToAll QMessageBox__StandardButton = 32768
const QMessageBox__No QMessageBox__StandardButton = 65536
const QMessageBox__NoToAll QMessageBox__StandardButton = 131072
const QMessageBox__Abort QMessageBox__StandardButton = 262144
const QMessageBox__Retry QMessageBox__StandardButton = 524288
const QMessageBox__Ignore QMessageBox__StandardButton = 1048576
const QMessageBox__Close QMessageBox__StandardButton = 2097152
const QMessageBox__Cancel QMessageBox__StandardButton = 4194304
const QMessageBox__Discard QMessageBox__StandardButton = 8388608
const QMessageBox__Help QMessageBox__StandardButton = 16777216
const QMessageBox__Apply QMessageBox__StandardButton = 33554432
const QMessageBox__Reset QMessageBox__StandardButton = 67108864
const QMessageBox__RestoreDefaults QMessageBox__StandardButton = 134217728
const QMessageBox__FirstButton QMessageBox__StandardButton = 1024
const QMessageBox__LastButton QMessageBox__StandardButton = 134217728
const QMessageBox__YesAll QMessageBox__StandardButton = 32768
const QMessageBox__NoAll QMessageBox__StandardButton = 131072
const QMessageBox__Default QMessageBox__StandardButton = 256
const QMessageBox__Escape QMessageBox__StandardButton = 512
const QMessageBox__FlagMask QMessageBox__StandardButton = 768
const QMessageBox__ButtonMask QMessageBox__StandardButton = 4294966527

//  body block end
