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
func NewQDialogButtonBoxFromPointer(cthis unsafe.Pointer) *QDialogButtonBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialogButtonBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDialogButtonBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:120
// index:0
// Public
// void QDialogButtonBox(class QWidget *)
func NewQDialogButtonBox(parent unsafe.Pointer) *QDialogButtonBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:121
// index:1
// Public
// void QDialogButtonBox(Qt::Orientation, class QWidget *)
func NewQDialogButtonBox_1(orientation int, parent unsafe.Pointer) *QDialogButtonBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:125
// index:0
// Public virtual
// void ~QDialogButtonBox()
func DeleteQDialogButtonBox(*QDialogButtonBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:127
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QDialogButtonBox) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:128
// index:0
// Public
// Qt::Orientation orientation()
func (this *QDialogButtonBox) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:130
// index:0
// Public
// void addButton(class QAbstractButton *, enum QDialogButtonBox::ButtonRole)
func (this *QDialogButtonBox) AddButton(button unsafe.Pointer, role int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:131
// index:1
// Public
// QPushButton * addButton(const class QString &, enum QDialogButtonBox::ButtonRole)
func (this *QDialogButtonBox) AddButton_1(text *qtcore.QString, role int) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonERK7QStringNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:132
// index:2
// Public
// QPushButton * addButton(enum QDialogButtonBox::StandardButton)
func (this *QDialogButtonBox) AddButton_2(button int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:133
// index:0
// Public
// void removeButton(class QAbstractButton *)
func (this *QDialogButtonBox) RemoveButton(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:134
// index:0
// Public
// void clear()
func (this *QDialogButtonBox) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:136
// index:0
// Public
// QList<QAbstractButton *> buttons()
func (this *QDialogButtonBox) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:137
// index:0
// Public
// QDialogButtonBox::ButtonRole buttonRole(class QAbstractButton *)
func (this *QDialogButtonBox) ButtonRole(button unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:140
// index:0
// Public
// QDialogButtonBox::StandardButtons standardButtons()
func (this *QDialogButtonBox) StandardButtons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox15standardButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:141
// index:0
// Public
// QDialogButtonBox::StandardButton standardButton(class QAbstractButton *)
func (this *QDialogButtonBox) StandardButton(button unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:142
// index:0
// Public
// QPushButton * button(enum QDialogButtonBox::StandardButton)
func (this *QDialogButtonBox) Button(which int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox6buttonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:144
// index:0
// Public
// void setCenterButtons(_Bool)
func (this *QDialogButtonBox) SetCenterButtons(center bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox16setCenterButtonsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &center)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:145
// index:0
// Public
// bool centerButtons()
func (this *QDialogButtonBox) CenterButtons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox13centerButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:148
// index:0
// Public
// void clicked(class QAbstractButton *)
func (this *QDialogButtonBox) Clicked(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox7clickedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:149
// index:0
// Public
// void accepted()
func (this *QDialogButtonBox) Accepted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox8acceptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:150
// index:0
// Public
// void helpRequested()
func (this *QDialogButtonBox) HelpRequested() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox13helpRequestedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:151
// index:0
// Public
// void rejected()
func (this *QDialogButtonBox) Rejected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox8rejectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:154
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QDialogButtonBox) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:155
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QDialogButtonBox) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
