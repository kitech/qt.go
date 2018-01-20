//  header block begin
// /usr/include/qt/QtWidgets/qtoolbutton.h
// #include <qtoolbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QToolButton struct {
	*QAbstractButton
}

func (this *QToolButton) GetCthis() unsafe.Pointer {
	return this.QAbstractButton.GetCthis()
}
func NewQToolButtonFromPointer(cthis unsafe.Pointer) *QToolButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QToolButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QToolButton) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:74
// index:0
// Public
// void QToolButton(class QWidget *)
func NewQToolButton(parent unsafe.Pointer) *QToolButton {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:75
// index:0
// Public virtual
// void ~QToolButton()
func DeleteQToolButton(*QToolButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:77
// index:0
// Public virtual
// QSize sizeHint()
func (this *QToolButton) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:78
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QToolButton) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:80
// index:0
// Public
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QToolButton) ToolButtonStyle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton15toolButtonStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:82
// index:0
// Public
// Qt::ArrowType arrowType()
func (this *QToolButton) ArrowType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9arrowTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:83
// index:0
// Public
// void setArrowType(Qt::ArrowType)
func (this *QToolButton) SetArrowType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton12setArrowTypeEN2Qt9ArrowTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:86
// index:0
// Public
// void setMenu(class QMenu *)
func (this *QToolButton) SetMenu(menu unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton7setMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:87
// index:0
// Public
// QMenu * menu()
func (this *QToolButton) Menu() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton4menuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:89
// index:0
// Public
// void setPopupMode(enum QToolButton::ToolButtonPopupMode)
func (this *QToolButton) SetPopupMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton12setPopupModeENS_19ToolButtonPopupModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:90
// index:0
// Public
// QToolButton::ToolButtonPopupMode popupMode()
func (this *QToolButton) PopupMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9popupModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:93
// index:0
// Public
// QAction * defaultAction()
func (this *QToolButton) DefaultAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton13defaultActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:95
// index:0
// Public
// void setAutoRaise(_Bool)
func (this *QToolButton) SetAutoRaise(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton12setAutoRaiseEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:96
// index:0
// Public
// bool autoRaise()
func (this *QToolButton) AutoRaise() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9autoRaiseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:100
// index:0
// Public
// void showMenu()
func (this *QToolButton) ShowMenu() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton8showMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:102
// index:0
// Public
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QToolButton) SetToolButtonStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:103
// index:0
// Public
// void setDefaultAction(class QAction *)
func (this *QToolButton) SetDefaultAction(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton16setDefaultActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:106
// index:0
// Public
// void triggered(class QAction *)
func (this *QToolButton) Triggered(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton9triggeredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:109
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QToolButton) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:110
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QToolButton) MousePressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:111
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QToolButton) MouseReleaseEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:112
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QToolButton) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:113
// index:0
// Protected virtual
// void actionEvent(class QActionEvent *)
func (this *QToolButton) ActionEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:115
// index:0
// Protected virtual
// void enterEvent(class QEvent *)
func (this *QToolButton) EnterEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton10enterEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:116
// index:0
// Protected virtual
// void leaveEvent(class QEvent *)
func (this *QToolButton) LeaveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton10leaveEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:117
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QToolButton) TimerEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:118
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QToolButton) ChangeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:120
// index:0
// Protected virtual
// bool hitButton(const class QPoint &)
func (this *QToolButton) HitButton(pos *qtcore.QPoint) interface{} {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9hitButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:121
// index:0
// Protected virtual
// void nextCheckState()
func (this *QToolButton) NextCheckState() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton14nextCheckStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:122
// index:0
// Protected
// void initStyleOption(class QStyleOptionToolButton *)
func (this *QToolButton) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
