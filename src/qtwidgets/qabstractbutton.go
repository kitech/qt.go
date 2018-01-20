//  header block begin
// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 278
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
type QAbstractButton struct {
	*QWidget
}

func (this *QAbstractButton) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQAbstractButtonFromPointer(cthis unsafe.Pointer) *QAbstractButton {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractButton) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// Public
// void QAbstractButton(class QWidget *)
func NewQAbstractButton(parent unsafe.Pointer) *QAbstractButton {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButtonC1EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:76
// index:0
// Public virtual
// void ~QAbstractButton()
func DeleteQAbstractButton(*QAbstractButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:78
// index:0
// Public
// void setText(const class QString &)
func (this *QAbstractButton) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:79
// index:0
// Public
// QString text()
func (this *QAbstractButton) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:81
// index:0
// Public
// void setIcon(const class QIcon &)
func (this *QAbstractButton) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:82
// index:0
// Public
// QIcon icon()
func (this *QAbstractButton) Icon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:84
// index:0
// Public
// QSize iconSize()
func (this *QAbstractButton) IconSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:87
// index:0
// Public
// void setShortcut(const class QKeySequence &)
func (this *QAbstractButton) SetShortcut(key *qtgui.QKeySequence) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton11setShortcutERK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:88
// index:0
// Public
// QKeySequence shortcut()
func (this *QAbstractButton) Shortcut() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton8shortcutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:91
// index:0
// Public
// void setCheckable(_Bool)
func (this *QAbstractButton) SetCheckable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton12setCheckableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:92
// index:0
// Public
// bool isCheckable()
func (this *QAbstractButton) IsCheckable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton11isCheckableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:94
// index:0
// Public
// bool isChecked()
func (this *QAbstractButton) IsChecked() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton9isCheckedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:96
// index:0
// Public
// void setDown(_Bool)
func (this *QAbstractButton) SetDown(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7setDownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:97
// index:0
// Public
// bool isDown()
func (this *QAbstractButton) IsDown() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton6isDownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:99
// index:0
// Public
// void setAutoRepeat(_Bool)
func (this *QAbstractButton) SetAutoRepeat(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13setAutoRepeatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:100
// index:0
// Public
// bool autoRepeat()
func (this *QAbstractButton) AutoRepeat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton10autoRepeatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:102
// index:0
// Public
// void setAutoRepeatDelay(int)
func (this *QAbstractButton) SetAutoRepeatDelay(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton18setAutoRepeatDelayEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:103
// index:0
// Public
// int autoRepeatDelay()
func (this *QAbstractButton) AutoRepeatDelay() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton15autoRepeatDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:105
// index:0
// Public
// void setAutoRepeatInterval(int)
func (this *QAbstractButton) SetAutoRepeatInterval(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton21setAutoRepeatIntervalEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:106
// index:0
// Public
// int autoRepeatInterval()
func (this *QAbstractButton) AutoRepeatInterval() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton18autoRepeatIntervalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:108
// index:0
// Public
// void setAutoExclusive(_Bool)
func (this *QAbstractButton) SetAutoExclusive(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton16setAutoExclusiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:109
// index:0
// Public
// bool autoExclusive()
func (this *QAbstractButton) AutoExclusive() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton13autoExclusiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:112
// index:0
// Public
// QButtonGroup * group()
func (this *QAbstractButton) Group() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:116
// index:0
// Public
// void setIconSize(const class QSize &)
func (this *QAbstractButton) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:117
// index:0
// Public
// void animateClick(int)
func (this *QAbstractButton) AnimateClick(msec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton12animateClickEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:118
// index:0
// Public
// void click()
func (this *QAbstractButton) Click() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton5clickEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:119
// index:0
// Public
// void toggle()
func (this *QAbstractButton) Toggle() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton6toggleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:120
// index:0
// Public
// void setChecked(_Bool)
func (this *QAbstractButton) SetChecked(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton10setCheckedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:123
// index:0
// Public
// void pressed()
func (this *QAbstractButton) Pressed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7pressedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:124
// index:0
// Public
// void released()
func (this *QAbstractButton) Released() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton8releasedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// Public
// void clicked(_Bool)
func (this *QAbstractButton) Clicked(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7clickedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:126
// index:0
// Public
// void toggled(_Bool)
func (this *QAbstractButton) Toggled(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7toggledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:129
// index:0
// Protected pure virtual
// void paintEvent(class QPaintEvent *)
func (this *QAbstractButton) PaintEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:130
// index:0
// Protected virtual
// bool hitButton(const class QPoint &)
func (this *QAbstractButton) HitButton(pos *qtcore.QPoint) interface{} {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton9hitButtonERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:131
// index:0
// Protected virtual
// void checkStateSet()
func (this *QAbstractButton) CheckStateSet() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13checkStateSetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:132
// index:0
// Protected virtual
// void nextCheckState()
func (this *QAbstractButton) NextCheckState() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton14nextCheckStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:134
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAbstractButton) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:135
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractButton) KeyPressEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:136
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractButton) KeyReleaseEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:137
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractButton) MousePressEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:138
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractButton) MouseReleaseEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:139
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractButton) MouseMoveEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:140
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractButton) FocusInEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:141
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractButton) FocusOutEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:142
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QAbstractButton) ChangeEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:143
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractButton) TimerEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
