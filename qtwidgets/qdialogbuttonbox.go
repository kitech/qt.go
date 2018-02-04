package qtwidgets

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h
// #include <qdialogbuttonbox.h>
// #include <QtWidgets>

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
// void changeEvent(class QEvent *)
func (this *QDialogButtonBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// bool event(class QEvent *)
func (this *QDialogButtonBox) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

type QDialogButtonBox struct {
	*QWidget
}

func (this *QDialogButtonBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDialogButtonBox) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDialogButtonBoxFromPointer(cthis unsafe.Pointer) *QDialogButtonBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialogButtonBox{bcthis0}
}
func (*QDialogButtonBox) NewFromPointer(cthis unsafe.Pointer) *QDialogButtonBox {
	return NewQDialogButtonBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDialogButtonBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QWidget *)
func NewQDialogButtonBox(parent *QWidget /*777 QWidget **/) *QDialogButtonBox {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:121
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(Qt::Orientation, QWidget *)
func NewQDialogButtonBox_1(orientation int, parent *QWidget /*777 QWidget **/) *QDialogButtonBox {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_POINTER, orientation, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:122
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QDialogButtonBox::StandardButtons, QWidget *)
func NewQDialogButtonBox_2(buttons int, parent *QWidget /*777 QWidget **/) *QDialogButtonBox {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEP7QWidget", ffiqt.FFI_TYPE_POINTER, buttons, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:123
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QDialogButtonBox::StandardButtons, Qt::Orientation, QWidget *)
func NewQDialogButtonBox_3(buttons int, orientation int, parent *QWidget /*777 QWidget **/) *QDialogButtonBox {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_POINTER, buttons, orientation, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDialogButtonBox()
func DeleteQDialogButtonBox(this *QDialogButtonBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBoxD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QDialogButtonBox) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QDialogButtonBox) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addButton(QAbstractButton *, enum QDialogButtonBox::ButtonRole)
func (this *QDialogButtonBox) AddButton(button *QAbstractButton /*777 QAbstractButton **/, role int) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:131
// index:1
// Public Visibility=Default Availability=Available
// [8] QPushButton * addButton(const QString &, enum QDialogButtonBox::ButtonRole)
func (this *QDialogButtonBox) AddButton_1(text *qtcore.QString, role int) *QPushButton /*777 QPushButton **/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonERK7QStringNS_10ButtonRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:132
// index:2
// Public Visibility=Default Availability=Available
// [8] QPushButton * addButton(enum QDialogButtonBox::StandardButton)
func (this *QDialogButtonBox) AddButton_2(button int) *QPushButton /*777 QPushButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeButton(QAbstractButton *)
func (this *QDialogButtonBox) RemoveButton(button *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox12removeButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QDialogButtonBox) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QDialogButtonBox::ButtonRole buttonRole(QAbstractButton *)
func (this *QDialogButtonBox) ButtonRole(button *QAbstractButton /*777 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardButtons(QDialogButtonBox::StandardButtons)
func (this *QDialogButtonBox) SetStandardButtons(buttons int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox18setStandardButtonsE6QFlagsINS_14StandardButtonEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QDialogButtonBox::StandardButtons standardButtons()
func (this *QDialogButtonBox) StandardButtons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox15standardButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] QDialogButtonBox::StandardButton standardButton(QAbstractButton *)
func (this *QDialogButtonBox) StandardButton(button *QAbstractButton /*777 QAbstractButton **/) int {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QPushButton * button(enum QDialogButtonBox::StandardButton)
func (this *QDialogButtonBox) Button(which int) *QPushButton /*777 QPushButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox6buttonENS_14StandardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenterButtons(_Bool)
func (this *QDialogButtonBox) SetCenterButtons(center bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox16setCenterButtonsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), center)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool centerButtons()
func (this *QDialogButtonBox) CenterButtons() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QDialogButtonBox13centerButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(QAbstractButton *)
func (this *QDialogButtonBox) Clicked(button *QAbstractButton /*777 QAbstractButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox7clickedEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accepted()
func (this *QDialogButtonBox) Accepted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox8acceptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void helpRequested()
func (this *QDialogButtonBox) HelpRequested() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox13helpRequestedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rejected()
func (this *QDialogButtonBox) Rejected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox8rejectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QDialogButtonBox) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QDialogButtonBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDialogButtonBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QDialogButtonBox__ButtonRole = int

const QDialogButtonBox__InvalidRole QDialogButtonBox__ButtonRole = 4294967295
const QDialogButtonBox__AcceptRole QDialogButtonBox__ButtonRole = 0
const QDialogButtonBox__RejectRole QDialogButtonBox__ButtonRole = 1
const QDialogButtonBox__DestructiveRole QDialogButtonBox__ButtonRole = 2
const QDialogButtonBox__ActionRole QDialogButtonBox__ButtonRole = 3
const QDialogButtonBox__HelpRole QDialogButtonBox__ButtonRole = 4
const QDialogButtonBox__YesRole QDialogButtonBox__ButtonRole = 5
const QDialogButtonBox__NoRole QDialogButtonBox__ButtonRole = 6
const QDialogButtonBox__ResetRole QDialogButtonBox__ButtonRole = 7
const QDialogButtonBox__ApplyRole QDialogButtonBox__ButtonRole = 8
const QDialogButtonBox__NRoles QDialogButtonBox__ButtonRole = 9

type QDialogButtonBox__StandardButton = int

const QDialogButtonBox__NoButton QDialogButtonBox__StandardButton = 0
const QDialogButtonBox__Ok QDialogButtonBox__StandardButton = 1024
const QDialogButtonBox__Save QDialogButtonBox__StandardButton = 2048
const QDialogButtonBox__SaveAll QDialogButtonBox__StandardButton = 4096
const QDialogButtonBox__Open QDialogButtonBox__StandardButton = 8192
const QDialogButtonBox__Yes QDialogButtonBox__StandardButton = 16384
const QDialogButtonBox__YesToAll QDialogButtonBox__StandardButton = 32768
const QDialogButtonBox__No QDialogButtonBox__StandardButton = 65536
const QDialogButtonBox__NoToAll QDialogButtonBox__StandardButton = 131072
const QDialogButtonBox__Abort QDialogButtonBox__StandardButton = 262144
const QDialogButtonBox__Retry QDialogButtonBox__StandardButton = 524288
const QDialogButtonBox__Ignore QDialogButtonBox__StandardButton = 1048576
const QDialogButtonBox__Close QDialogButtonBox__StandardButton = 2097152
const QDialogButtonBox__Cancel QDialogButtonBox__StandardButton = 4194304
const QDialogButtonBox__Discard QDialogButtonBox__StandardButton = 8388608
const QDialogButtonBox__Help QDialogButtonBox__StandardButton = 16777216
const QDialogButtonBox__Apply QDialogButtonBox__StandardButton = 33554432
const QDialogButtonBox__Reset QDialogButtonBox__StandardButton = 67108864
const QDialogButtonBox__RestoreDefaults QDialogButtonBox__StandardButton = 134217728
const QDialogButtonBox__FirstButton QDialogButtonBox__StandardButton = 1024
const QDialogButtonBox__LastButton QDialogButtonBox__StandardButton = 134217728

type QDialogButtonBox__ButtonLayout = int

const QDialogButtonBox__WinLayout QDialogButtonBox__ButtonLayout = 0
const QDialogButtonBox__MacLayout QDialogButtonBox__ButtonLayout = 1
const QDialogButtonBox__KdeLayout QDialogButtonBox__ButtonLayout = 2
const QDialogButtonBox__GnomeLayout QDialogButtonBox__ButtonLayout = 3
const QDialogButtonBox__AndroidLayout QDialogButtonBox__ButtonLayout = 5

//  body block end
