//  header block begin
// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 290
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

// /usr/include/qt/QtWidgets/qabstractbutton.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// void QAbstractButton(class QWidget *)
func NewQAbstractButton(parent unsafe.Pointer) *QAbstractButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButtonC1EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractButtonFromPointer(cthis)
	return gothis
}
func NewQAbstractButtonFromPointer(cthis unsafe.Pointer) *QAbstractButton {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:147
// index:1
// void QAbstractButton(class QAbstractButtonPrivate &, class QWidget *)
func NewQAbstractButton_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButtonC1ER22QAbstractButtonPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:76
// index:0
// virtual
// void ~QAbstractButton()
func DeleteQAbstractButton(*QAbstractButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:78
// index:0
// void setText(const class QString &)
func (this *QAbstractButton) SetText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:79
// index:0
// QString text()
func (this *QAbstractButton) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:81
// index:0
// void setIcon(const class QIcon &)
func (this *QAbstractButton) SetIcon(icon unsafe.Pointer) {
	// 0: (, icon const QIcon &), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:82
// index:0
// QIcon icon()
func (this *QAbstractButton) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton4iconEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:84
// index:0
// QSize iconSize()
func (this *QAbstractButton) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:87
// index:0
// void setShortcut(const class QKeySequence &)
func (this *QAbstractButton) SetShortcut(key unsafe.Pointer) {
	// 0: (, key const QKeySequence &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton11setShortcutERK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:88
// index:0
// QKeySequence shortcut()
func (this *QAbstractButton) Shortcut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton8shortcutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:91
// index:0
// void setCheckable(_Bool)
func (this *QAbstractButton) SetCheckable(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton12setCheckableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:92
// index:0
// bool isCheckable()
func (this *QAbstractButton) IsCheckable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton11isCheckableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:94
// index:0
// bool isChecked()
func (this *QAbstractButton) IsChecked() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton9isCheckedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:96
// index:0
// void setDown(_Bool)
func (this *QAbstractButton) SetDown(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7setDownEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:97
// index:0
// bool isDown()
func (this *QAbstractButton) IsDown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton6isDownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:99
// index:0
// void setAutoRepeat(_Bool)
func (this *QAbstractButton) SetAutoRepeat(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13setAutoRepeatEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:100
// index:0
// bool autoRepeat()
func (this *QAbstractButton) AutoRepeat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton10autoRepeatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:102
// index:0
// void setAutoRepeatDelay(int)
func (this *QAbstractButton) SetAutoRepeatDelay(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton18setAutoRepeatDelayEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:103
// index:0
// int autoRepeatDelay()
func (this *QAbstractButton) AutoRepeatDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton15autoRepeatDelayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:105
// index:0
// void setAutoRepeatInterval(int)
func (this *QAbstractButton) SetAutoRepeatInterval(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton21setAutoRepeatIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:106
// index:0
// int autoRepeatInterval()
func (this *QAbstractButton) AutoRepeatInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton18autoRepeatIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:108
// index:0
// void setAutoExclusive(_Bool)
func (this *QAbstractButton) SetAutoExclusive(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton16setAutoExclusiveEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:109
// index:0
// bool autoExclusive()
func (this *QAbstractButton) AutoExclusive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton13autoExclusiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:112
// index:0
// QButtonGroup * group()
func (this *QAbstractButton) Group() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton5groupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:116
// index:0
// void setIconSize(const class QSize &)
func (this *QAbstractButton) SetIconSize(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:117
// index:0
// void animateClick(int)
func (this *QAbstractButton) AnimateClick(msec int) {
	// 0: (, msec int), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton12animateClickEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:118
// index:0
// void click()
func (this *QAbstractButton) Click() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton5clickEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:119
// index:0
// void toggle()
func (this *QAbstractButton) Toggle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton6toggleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:120
// index:0
// void setChecked(_Bool)
func (this *QAbstractButton) SetChecked(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton10setCheckedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:123
// index:0
// void pressed()
func (this *QAbstractButton) Pressed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7pressedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:124
// index:0
// void released()
func (this *QAbstractButton) Released() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton8releasedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// void clicked(_Bool)
func (this *QAbstractButton) Clicked(checked bool) {
	// 0: (, checked bool), (&checked)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7clickedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:126
// index:0
// void toggled(_Bool)
func (this *QAbstractButton) Toggled(checked bool) {
	// 0: (, checked bool), (&checked)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton7toggledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:129
// index:0
// pure virtual
// void paintEvent(class QPaintEvent *)
func (this *QAbstractButton) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:130
// index:0
// virtual
// bool hitButton(const class QPoint &)
func (this *QAbstractButton) HitButton(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractButton9hitButtonERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:131
// index:0
// virtual
// void checkStateSet()
func (this *QAbstractButton) CheckStateSet() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13checkStateSetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:132
// index:0
// virtual
// void nextCheckState()
func (this *QAbstractButton) NextCheckState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton14nextCheckStateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:134
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAbstractButton) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:135
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractButton) KeyPressEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:136
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractButton) KeyReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:137
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractButton) MousePressEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:138
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractButton) MouseReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:139
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractButton) MouseMoveEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:140
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractButton) FocusInEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:141
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractButton) FocusOutEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:142
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QAbstractButton) ChangeEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:143
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractButton) TimerEvent(e unsafe.Pointer) {
	// 0: (, e QTimerEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractButton10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
