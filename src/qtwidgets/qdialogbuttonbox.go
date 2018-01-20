//  header block begin
// /usr/include/qt/QtWidgets/qdialogbuttonbox.h
// #include <qdialogbuttonbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QDialogButtonBox struct {
	*QWidget
}

func (this *QDialogButtonBox) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDialogButtonBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:120
// index:0
// void QDialogButtonBox(class QWidget *)
func NewQDialogButtonBox(parent unsafe.Pointer) *QDialogButtonBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(cthis)
	return gothis
}
func NewQDialogButtonBoxFromPointer(cthis unsafe.Pointer) *QDialogButtonBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialogButtonBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:121
// index:1
// void QDialogButtonBox(Qt::Orientation, class QWidget *)
func NewQDialogButtonBox_1(orientation int, parent unsafe.Pointer) *QDialogButtonBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:122
// index:2
// void QDialogButtonBox(QDialogButtonBox::StandardButtons, class QWidget *)
func NewQDialogButtonBox_2(buttons int, parent unsafe.Pointer) *QDialogButtonBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, buttons, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:123
// index:3
// void QDialogButtonBox(QDialogButtonBox::StandardButtons, Qt::Orientation, class QWidget *)
func NewQDialogButtonBox_3(buttons int, orientation int, parent unsafe.Pointer) *QDialogButtonBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, buttons, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:125
// index:0
// virtual
// void ~QDialogButtonBox()
func DeleteQDialogButtonBox(*QDialogButtonBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:127
// index:0
// void setOrientation(Qt::Orientation)
func (this *QDialogButtonBox) SetOrientation(orientation int) {
	// 0: (, orientation Qt::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:128
// index:0
// Qt::Orientation orientation()
func (this *QDialogButtonBox) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:130
// index:0
// void addButton(class QAbstractButton *, enum QDialogButtonBox::ButtonRole)
func (this *QDialogButtonBox) AddButton(button unsafe.Pointer, role int) {
	// 0: (, button QAbstractButton *, role QDialogButtonBox::ButtonRole), (button, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:131
// index:1
// QPushButton * addButton(const class QString &, enum QDialogButtonBox::ButtonRole)
func (this *QDialogButtonBox) AddButton_1(text unsafe.Pointer, role int) {
	// 1: (, text const QString &, role QDialogButtonBox::ButtonRole), (text, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonERK7QStringNS_10ButtonRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:132
// index:2
// QPushButton * addButton(enum QDialogButtonBox::StandardButton)
func (this *QDialogButtonBox) AddButton_2(button int) {
	// 2: (, button QDialogButtonBox::StandardButton), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:133
// index:0
// void removeButton(class QAbstractButton *)
func (this *QDialogButtonBox) RemoveButton(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:134
// index:0
// void clear()
func (this *QDialogButtonBox) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:136
// index:0
// QList<QAbstractButton *> buttons()
func (this *QDialogButtonBox) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:137
// index:0
// QDialogButtonBox::ButtonRole buttonRole(class QAbstractButton *)
func (this *QDialogButtonBox) ButtonRole(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:139
// index:0
// void setStandardButtons(QDialogButtonBox::StandardButtons)
func (this *QDialogButtonBox) SetStandardButtons(buttons int) {
	// 0: (, QFlags<QDialogButtonBox::StandardButton> buttons), (buttons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox18setStandardButtonsE6QFlagsINS_14StandardButtonEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:140
// index:0
// QDialogButtonBox::StandardButtons standardButtons()
func (this *QDialogButtonBox) StandardButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox15standardButtonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:141
// index:0
// QDialogButtonBox::StandardButton standardButton(class QAbstractButton *)
func (this *QDialogButtonBox) StandardButton(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:142
// index:0
// QPushButton * button(enum QDialogButtonBox::StandardButton)
func (this *QDialogButtonBox) Button(which int) {
	// 0: (, which QDialogButtonBox::StandardButton), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox6buttonENS_14StandardButtonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:144
// index:0
// void setCenterButtons(_Bool)
func (this *QDialogButtonBox) SetCenterButtons(center bool) {
	// 0: (, center bool), (&center)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox16setCenterButtonsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &center)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:145
// index:0
// bool centerButtons()
func (this *QDialogButtonBox) CenterButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox13centerButtonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:148
// index:0
// void clicked(class QAbstractButton *)
func (this *QDialogButtonBox) Clicked(button unsafe.Pointer) {
	// 0: (, button QAbstractButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox7clickedEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:149
// index:0
// void accepted()
func (this *QDialogButtonBox) Accepted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox8acceptedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:150
// index:0
// void helpRequested()
func (this *QDialogButtonBox) HelpRequested() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox13helpRequestedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:151
// index:0
// void rejected()
func (this *QDialogButtonBox) Rejected() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox8rejectedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:154
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QDialogButtonBox) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:155
// index:0
// virtual
// bool event(class QEvent *)
func (this *QDialogButtonBox) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
