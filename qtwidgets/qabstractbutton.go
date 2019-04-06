// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 288
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void paintEvent(QPaintEvent *)
func (this *QAbstractButton) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// bool hitButton(const QPoint &)
func (this *QAbstractButton) InheritHitButton(f func(pos *qtcore.QPoint) bool) {
	qtrt.SetAllInheritCallback(this, "hitButton", f)
}

// void checkStateSet()
func (this *QAbstractButton) InheritCheckStateSet(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "checkStateSet", f)
}

// void nextCheckState()
func (this *QAbstractButton) InheritNextCheckState(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "nextCheckState", f)
}

// bool event(QEvent *)
func (this *QAbstractButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QAbstractButton) InheritKeyPressEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QAbstractButton) InheritKeyReleaseEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QAbstractButton) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractButton) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractButton) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QAbstractButton) InheritFocusInEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QAbstractButton) InheritFocusOutEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void changeEvent(QEvent *)
func (this *QAbstractButton) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QAbstractButton) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

/*

 */
type QAbstractButton struct {
	*QWidget
}
type QAbstractButton_ITF interface {
	QWidget_ITF
	QAbstractButton_PTR() *QAbstractButton
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton { return ptr }

func (this *QAbstractButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QAbstractButton) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQAbstractButtonFromPointer(cthis unsafe.Pointer) *QAbstractButton {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractButton{bcthis0}
}
func (*QAbstractButton) NewFromPointer(cthis unsafe.Pointer) *QAbstractButton {
	return NewQAbstractButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractButton(QWidget *)

/*
Constructs an abstract button with a parent.
*/
func (*QAbstractButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractButton {
	return NewQAbstractButton(parent)
}
func NewQAbstractButton(parent QWidget_ITF /*777 QWidget **/) *QAbstractButton {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractButton(QWidget *)

/*
Constructs an abstract button with a parent.
*/
func (*QAbstractButton) NewForInheritp() *QAbstractButton {
	return NewQAbstractButtonp()
}
func NewQAbstractButtonp() *QAbstractButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractButton()

/*

 */
func DeleteQAbstractButton(this *QAbstractButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*

 */
func (this *QAbstractButton) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QAbstractButton) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QAbstractButton) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QAbstractButton) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QAbstractButton) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcut(const QKeySequence &)

/*

 */
func (this *QAbstractButton) SetShortcut(key qtgui.QKeySequence_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QKeySequence_PTR() != nil {
		convArg0 = key.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton11setShortcutERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence shortcut() const

/*

 */
func (this *QAbstractButton) Shortcut() *qtgui.QKeySequence /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton8shortcutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(bool)

/*

 */
func (this *QAbstractButton) SetCheckable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12setCheckableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*

 */
func (this *QAbstractButton) IsCheckable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton11isCheckableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isChecked() const

/*

 */
func (this *QAbstractButton) IsChecked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton9isCheckedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDown(bool)

/*

 */
func (this *QAbstractButton) SetDown(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7setDownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDown() const

/*

 */
func (this *QAbstractButton) IsDown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton6isDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeat(bool)

/*

 */
func (this *QAbstractButton) SetAutoRepeat(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13setAutoRepeatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRepeat() const

/*

 */
func (this *QAbstractButton) AutoRepeat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton10autoRepeatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeatDelay(int)

/*

 */
func (this *QAbstractButton) SetAutoRepeatDelay(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton18setAutoRepeatDelayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoRepeatDelay() const

/*

 */
func (this *QAbstractButton) AutoRepeatDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton15autoRepeatDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeatInterval(int)

/*

 */
func (this *QAbstractButton) SetAutoRepeatInterval(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton21setAutoRepeatIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoRepeatInterval() const

/*

 */
func (this *QAbstractButton) AutoRepeatInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton18autoRepeatIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoExclusive(bool)

/*

 */
func (this *QAbstractButton) SetAutoExclusive(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton16setAutoExclusiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoExclusive() const

/*

 */
func (this *QAbstractButton) AutoExclusive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton13autoExclusiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QAbstractButton) SetIconSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void animateClick(int)

/*
Performs an animated click: the button is pressed immediately, and released msec milliseconds later (the default is 100 ms).

Calling this function again before the button is released resets the release timer.

All signals associated with a click are emitted as appropriate.

This function does nothing if the button is disabled.

See also click().
*/
func (this *QAbstractButton) AnimateClick(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12animateClickEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void animateClick(int)

/*
Performs an animated click: the button is pressed immediately, and released msec milliseconds later (the default is 100 ms).

Calling this function again before the button is released resets the release timer.

All signals associated with a click are emitted as appropriate.

This function does nothing if the button is disabled.

See also click().
*/
func (this *QAbstractButton) AnimateClickp() {
	// arg: 0, int=Int, =Invalid, , Invalid
	msec := int(100)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12animateClickEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void click()

/*
Performs a click.

All the usual signals associated with a click are emitted as appropriate. If the button is checkable, the state of the button is toggled.

This function does nothing if the button is disabled.

See also animateClick().
*/
func (this *QAbstractButton) Click() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton5clickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggle()

/*
Toggles the state of a checkable button.

See also checked.
*/
func (this *QAbstractButton) Toggle() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton6toggleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChecked(bool)

/*

 */
func (this *QAbstractButton) SetChecked(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton10setCheckedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pressed()

/*
This signal is emitted when the button is pressed down.

See also released() and clicked().
*/
func (this *QAbstractButton) Pressed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7pressedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void released()

/*
This signal is emitted when the button is released.

See also pressed(), clicked(), and toggled().
*/
func (this *QAbstractButton) Released() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton8releasedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
This signal is emitted when the button is activated (i.e., pressed down then released while the mouse cursor is inside the button), when the shortcut key is typed, or when click() or animateClick() is called. Notably, this signal is not emitted if you call setDown(), setChecked() or toggle().

If the button is checkable, checked is true if the button is checked, or false if the button is unchecked.

See also pressed(), released(), and toggled().
*/
func (this *QAbstractButton) Clicked(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7clickedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
This signal is emitted when the button is activated (i.e., pressed down then released while the mouse cursor is inside the button), when the shortcut key is typed, or when click() or animateClick() is called. Notably, this signal is not emitted if you call setDown(), setChecked() or toggle().

If the button is checkable, checked is true if the button is checked, or false if the button is unchecked.

See also pressed(), released(), and toggled().
*/
func (this *QAbstractButton) Clickedp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7clickedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggled(bool)

/*
This signal is emitted whenever a checkable button changes its state. checked is true if the button is checked, or false if the button is unchecked.

This may be the result of a user action, click() slot activation, or because setChecked() is called.

The states of buttons in exclusive button groups are updated before this signal is emitted. This means that slots can act on either the "off" signal or the "on" signal emitted by the buttons in the group whose states have changed.

For example, a slot that reacts to signals emitted by newly checked buttons but which ignores signals from buttons that have been unchecked can be implemented using the following pattern:


  void MyWidget::reactToToggle(bool checked)
  {
     if (checked) {
        // Examine the new button states.
        ...
     }
  }



Button groups can be created using the QButtonGroup class, and updates to the button states monitored with the QButtonGroup::buttonClicked() signal.

Note: Notifier signal for property checked.

See also checked and clicked().
*/
func (this *QAbstractButton) Toggled(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7toggledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:129
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QAbstractButton) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:130
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool hitButton(const QPoint &) const

/*
Returns true if pos is inside the clickable button rectangle; otherwise returns false.

By default, the clickable area is the entire widget. Subclasses may reimplement this function to provide support for clickable areas of different shapes and sizes.
*/
func (this *QAbstractButton) HitButton(pos qtcore.QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton9hitButtonERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void checkStateSet()

/*
This virtual handler is called when setChecked() is used, unless it is called from within nextCheckState(). It allows subclasses to reset their intermediate button states.

See also nextCheckState().
*/
func (this *QAbstractButton) CheckStateSet() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13checkStateSetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void nextCheckState()

/*
This virtual handler is called when a button is clicked. The default implementation calls setChecked(!isChecked()) if the button isCheckable(). It allows subclasses to implement intermediate button states.

See also checkStateSet().
*/
func (this *QAbstractButton) NextCheckState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton14nextCheckStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QWidget::event().
*/
func (this *QAbstractButton) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:135
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QAbstractButton) KeyPressEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyReleaseEvent().
*/
func (this *QAbstractButton) KeyReleaseEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:137
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QAbstractButton) MousePressEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QAbstractButton) MouseReleaseEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:139
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QAbstractButton) MouseMoveEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QAbstractButton) FocusInEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QAbstractButton) FocusOutEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QAbstractButton) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QAbstractButton) TimerEvent(e qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QTimerEvent_PTR() != nil {
		convArg0 = e.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10985() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
