package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbutton.h
// #include <qtoolbutton.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
// bool event(class QEvent *)
func (this *QToolButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QToolButton) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QToolButton) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QToolButton) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void actionEvent(class QActionEvent *)
func (this *QToolButton) InheritActionEvent(f func(arg0 *qtgui.QActionEvent /*777 QActionEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void enterEvent(class QEvent *)
func (this *QToolButton) InheritEnterEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterEvent", f)
}

// void leaveEvent(class QEvent *)
func (this *QToolButton) InheritLeaveEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QToolButton) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QToolButton) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// bool hitButton(const class QPoint &)
func (this *QToolButton) InheritHitButton(f func(pos *qtcore.QPoint) bool) {
	qtrt.SetAllInheritCallback(this, "hitButton", f)
}

// void nextCheckState()
func (this *QToolButton) InheritNextCheckState(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "nextCheckState", f)
}

// void initStyleOption(class QStyleOptionToolButton *)
func (this *QToolButton) InheritInitStyleOption(f func(option *QStyleOptionToolButton /*777 QStyleOptionToolButton **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QToolButton struct {
	*QAbstractButton
}

func (this *QToolButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractButton.GetCthis()
	}
}
func (this *QToolButton) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractButton = NewQAbstractButtonFromPointer(cthis)
}
func NewQToolButtonFromPointer(cthis unsafe.Pointer) *QToolButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QToolButton{bcthis0}
}
func (*QToolButton) NewFromPointer(cthis unsafe.Pointer) *QToolButton {
	return NewQToolButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QToolButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolButton(QWidget *)
func NewQToolButton(parent *QWidget /*777 QWidget **/) *QToolButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QToolButton()
func DeleteQToolButton(this *QToolButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QToolButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QToolButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolButtonStyle toolButtonStyle()
func (this *QToolButton) ToolButtonStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton15toolButtonStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ArrowType arrowType()
func (this *QToolButton) ArrowType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton9arrowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArrowType(Qt::ArrowType)
func (this *QToolButton) SetArrowType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton12setArrowTypeEN2Qt9ArrowTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)
func (this *QToolButton) SetMenu(menu *QMenu /*777 QMenu **/) {
	var convArg0 = menu.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu()
func (this *QToolButton) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPopupMode(enum QToolButton::ToolButtonPopupMode)
func (this *QToolButton) SetPopupMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton12setPopupModeENS_19ToolButtonPopupModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QToolButton::ToolButtonPopupMode popupMode()
func (this *QToolButton) PopupMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton9popupModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * defaultAction()
func (this *QToolButton) DefaultAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton13defaultActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRaise(_Bool)
func (this *QToolButton) SetAutoRaise(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton12setAutoRaiseEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRaise()
func (this *QToolButton) AutoRaise() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton9autoRaiseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMenu()
func (this *QToolButton) ShowMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton8showMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QToolButton) SetToolButtonStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton18setToolButtonStyleEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultAction(QAction *)
func (this *QToolButton) SetDefaultAction(arg0 *QAction /*777 QAction **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton16setDefaultActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(QAction *)
func (this *QToolButton) Triggered(arg0 *QAction /*777 QAction **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton9triggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QToolButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QToolButton) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QToolButton) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QToolButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)
func (this *QToolButton) ActionEvent(arg0 *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:115
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void enterEvent(QEvent *)
func (this *QToolButton) EnterEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton10enterEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)
func (this *QToolButton) LeaveEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QToolButton) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QToolButton) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool hitButton(const QPoint &)
func (this *QToolButton) HitButton(pos *qtcore.QPoint) bool {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton9hitButtonERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void nextCheckState()
func (this *QToolButton) NextCheckState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QToolButton14nextCheckStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:122
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionToolButton *)
func (this *QToolButton) InitStyleOption(option *QStyleOptionToolButton /*777 QStyleOptionToolButton **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QToolButton__ToolButtonPopupMode = int

const QToolButton__DelayedPopup QToolButton__ToolButtonPopupMode = 0
const QToolButton__MenuButtonPopup QToolButton__ToolButtonPopupMode = 1
const QToolButton__InstantPopup QToolButton__ToolButtonPopupMode = 2

//  body block end
