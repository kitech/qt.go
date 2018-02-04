package qtwidgets

// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 288
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
// void paintEvent(class QPaintEvent *)
func (this *QAbstractButton) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// bool hitButton(const class QPoint &)
func (this *QAbstractButton) InheritHitButton(f func(pos *qtcore.QPoint) bool) {
	qtrt.SetAllInheritCallback(this, "hitButton", f)
}

// void checkStateSet()
func (this *QAbstractButton) InheritCheckStateSet(f func()) {
	qtrt.SetAllInheritCallback(this, "checkStateSet", f)
}

// void nextCheckState()
func (this *QAbstractButton) InheritNextCheckState(f func()) {
	qtrt.SetAllInheritCallback(this, "nextCheckState", f)
}

// bool event(class QEvent *)
func (this *QAbstractButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractButton) InheritKeyPressEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractButton) InheritKeyReleaseEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractButton) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractButton) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractButton) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QAbstractButton) InheritFocusInEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractButton) InheritFocusOutEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QAbstractButton) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QAbstractButton) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/)) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

type QAbstractButton struct {
	*QWidget
}

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
// [8] const QMetaObject * metaObject()
func (this *QAbstractButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractButton(QWidget *)
func NewQAbstractButton(parent *QWidget /*777 QWidget **/) *QAbstractButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButtonC1EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractButton()
func DeleteQAbstractButton(this *QAbstractButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QAbstractButton) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QAbstractButton) Text() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QAbstractButton) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QAbstractButton) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
func (this *QAbstractButton) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcut(const QKeySequence &)
func (this *QAbstractButton) SetShortcut(key *qtgui.QKeySequence) {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton11setShortcutERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence shortcut()
func (this *QAbstractButton) Shortcut() *qtgui.QKeySequence /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton8shortcutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(_Bool)
func (this *QAbstractButton) SetCheckable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12setCheckableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCheckable()
func (this *QAbstractButton) IsCheckable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton11isCheckableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isChecked()
func (this *QAbstractButton) IsChecked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton9isCheckedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDown(_Bool)
func (this *QAbstractButton) SetDown(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7setDownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDown()
func (this *QAbstractButton) IsDown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton6isDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeat(_Bool)
func (this *QAbstractButton) SetAutoRepeat(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13setAutoRepeatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRepeat()
func (this *QAbstractButton) AutoRepeat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton10autoRepeatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeatDelay(int)
func (this *QAbstractButton) SetAutoRepeatDelay(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton18setAutoRepeatDelayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoRepeatDelay()
func (this *QAbstractButton) AutoRepeatDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton15autoRepeatDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeatInterval(int)
func (this *QAbstractButton) SetAutoRepeatInterval(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton21setAutoRepeatIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoRepeatInterval()
func (this *QAbstractButton) AutoRepeatInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton18autoRepeatIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoExclusive(_Bool)
func (this *QAbstractButton) SetAutoExclusive(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton16setAutoExclusiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoExclusive()
func (this *QAbstractButton) AutoExclusive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton13autoExclusiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QButtonGroup * group()
func (this *QAbstractButton) Group() *QButtonGroup /*777 QButtonGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQButtonGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QAbstractButton) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void animateClick(int)
func (this *QAbstractButton) AnimateClick(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12animateClickEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void click()
func (this *QAbstractButton) Click() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton5clickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggle()
func (this *QAbstractButton) Toggle() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton6toggleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChecked(_Bool)
func (this *QAbstractButton) SetChecked(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton10setCheckedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pressed()
func (this *QAbstractButton) Pressed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7pressedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void released()
func (this *QAbstractButton) Released() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton8releasedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(_Bool)
func (this *QAbstractButton) Clicked(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7clickedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggled(_Bool)
func (this *QAbstractButton) Toggled(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton7toggledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:129
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QAbstractButton) PaintEvent(e *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:130
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool hitButton(const QPoint &)
func (this *QAbstractButton) HitButton(pos *qtcore.QPoint) bool {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton9hitButtonERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void checkStateSet()
func (this *QAbstractButton) CheckStateSet() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13checkStateSetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void nextCheckState()
func (this *QAbstractButton) NextCheckState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton14nextCheckStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:135
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QAbstractButton) KeyPressEvent(e *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QAbstractButton) KeyReleaseEvent(e *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:137
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QAbstractButton) MousePressEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractButton) MouseReleaseEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:139
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractButton) MouseMoveEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QAbstractButton) FocusInEvent(e *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QAbstractButton) FocusOutEvent(e *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QAbstractButton) ChangeEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QAbstractButton) TimerEvent(e *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractButton10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
