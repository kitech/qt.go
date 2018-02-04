package qtwidgets

// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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
// bool eventFilter(class QObject *, class QEvent *)
func (this *QMdiSubWindow) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(class QEvent *)
func (this *QMdiSubWindow) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void showEvent(class QShowEvent *)
func (this *QMdiSubWindow) InheritShowEvent(f func(showEvent *qtgui.QShowEvent /*777 QShowEvent **/)) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QMdiSubWindow) InheritHideEvent(f func(hideEvent *qtgui.QHideEvent /*777 QHideEvent **/)) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QMdiSubWindow) InheritChangeEvent(f func(changeEvent *qtcore.QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QMdiSubWindow) InheritCloseEvent(f func(closeEvent *qtgui.QCloseEvent /*777 QCloseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void leaveEvent(class QEvent *)
func (this *QMdiSubWindow) InheritLeaveEvent(f func(leaveEvent *qtcore.QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QMdiSubWindow) InheritResizeEvent(f func(resizeEvent *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QMdiSubWindow) InheritTimerEvent(f func(timerEvent *qtcore.QTimerEvent /*777 QTimerEvent **/)) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void moveEvent(class QMoveEvent *)
func (this *QMdiSubWindow) InheritMoveEvent(f func(moveEvent *qtgui.QMoveEvent /*777 QMoveEvent **/)) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QMdiSubWindow) InheritPaintEvent(f func(paintEvent *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QMdiSubWindow) InheritMousePressEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QMdiSubWindow) InheritMouseDoubleClickEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QMdiSubWindow) InheritMouseReleaseEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QMdiSubWindow) InheritMouseMoveEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QMdiSubWindow) InheritKeyPressEvent(f func(keyEvent *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QMdiSubWindow) InheritContextMenuEvent(f func(contextMenuEvent *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QMdiSubWindow) InheritFocusInEvent(f func(focusInEvent *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QMdiSubWindow) InheritFocusOutEvent(f func(focusOutEvent *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void childEvent(class QChildEvent *)
func (this *QMdiSubWindow) InheritChildEvent(f func(childEvent *qtcore.QChildEvent /*777 QChildEvent **/)) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

type QMdiSubWindow struct {
	*QWidget
}

func (this *QMdiSubWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMdiSubWindow) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMdiSubWindowFromPointer(cthis unsafe.Pointer) *QMdiSubWindow {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMdiSubWindow{bcthis0}
}
func (*QMdiSubWindow) NewFromPointer(cthis unsafe.Pointer) *QMdiSubWindow {
	return NewQMdiSubWindowFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMdiSubWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMdiSubWindow(QWidget *, Qt::WindowFlags)
func NewQMdiSubWindow(parent *QWidget /*777 QWidget **/, flags int) *QMdiSubWindow {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMdiSubWindow()
func DeleteQMdiSubWindow(this *QMdiSubWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QMdiSubWindow) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QMdiSubWindow) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)
func (this *QMdiSubWindow) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QMdiSubWindow) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * maximizedButtonsWidget()
func (this *QMdiSubWindow) MaximizedButtonsWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * maximizedSystemMenuIconWidget()
func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isShaded()
func (this *QMdiSubWindow) IsShaded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow8isShadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QMdiSubWindow::SubWindowOption, _Bool)
func (this *QMdiSubWindow) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9setOptionENS_15SubWindowOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(enum QMdiSubWindow::SubWindowOption)
func (this *QMdiSubWindow) TestOption(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10testOptionENS_15SubWindowOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardSingleStep(int)
func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow21setKeyboardSingleStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyboardSingleStep()
func (this *QMdiSubWindow) KeyboardSingleStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow18keyboardSingleStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardPageStep(int)
func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow19setKeyboardPageStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyboardPageStep()
func (this *QMdiSubWindow) KeyboardPageStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow16keyboardPageStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSystemMenu(QMenu *)
func (this *QMdiSubWindow) SetSystemMenu(systemMenu *QMenu /*777 QMenu **/) {
	var convArg0 = systemMenu.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13setSystemMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * systemMenu()
func (this *QMdiSubWindow) SystemMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10systemMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QMdiArea * mdiArea()
func (this *QMdiSubWindow) MdiArea() *QMdiArea /*777 QMdiArea **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow7mdiAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMdiAreaFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowStateChanged(Qt::WindowStates, Qt::WindowStates)
func (this *QMdiSubWindow) WindowStateChanged(oldState int, newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow18windowStateChangedE6QFlagsIN2Qt11WindowStateEES3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldState, newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToActivate()
func (this *QMdiSubWindow) AboutToActivate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow15aboutToActivateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showSystemMenu()
func (this *QMdiSubWindow) ShowSystemMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow14showSystemMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showShaded()
func (this *QMdiSubWindow) ShowShaded() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10showShadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QMdiSubWindow) EventFilter(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QMdiSubWindow) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QMdiSubWindow) ShowEvent(showEvent *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = showEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QMdiSubWindow) HideEvent(hideEvent *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = hideEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QMdiSubWindow) ChangeEvent(changeEvent *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = changeEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:115
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QMdiSubWindow) CloseEvent(closeEvent *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = closeEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)
func (this *QMdiSubWindow) LeaveEvent(leaveEvent *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = leaveEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QMdiSubWindow) ResizeEvent(resizeEvent *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = resizeEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QMdiSubWindow) TimerEvent(timerEvent *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = timerEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)
func (this *QMdiSubWindow) MoveEvent(moveEvent *qtgui.QMoveEvent /*777 QMoveEvent **/) {
	var convArg0 = moveEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QMdiSubWindow) PaintEvent(paintEvent *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = paintEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QMdiSubWindow) MousePressEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QMdiSubWindow) MouseDoubleClickEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QMdiSubWindow) MouseReleaseEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QMdiSubWindow) MouseMoveEvent(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = mouseEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QMdiSubWindow) KeyPressEvent(keyEvent *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = keyEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:127
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QMdiSubWindow) ContextMenuEvent(contextMenuEvent *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = contextMenuEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:129
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QMdiSubWindow) FocusInEvent(focusInEvent *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = focusInEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:130
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QMdiSubWindow) FocusOutEvent(focusOutEvent *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = focusOutEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)
func (this *QMdiSubWindow) ChildEvent(childEvent *qtcore.QChildEvent /*777 QChildEvent **/) {
	var convArg0 = childEvent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QMdiSubWindow__SubWindowOption = int

const QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = 1
const QMdiSubWindow__AllowOutsideAreaVertically QMdiSubWindow__SubWindowOption = 2
const QMdiSubWindow__RubberBandResize QMdiSubWindow__SubWindowOption = 4
const QMdiSubWindow__RubberBandMove QMdiSubWindow__SubWindowOption = 8

//  body block end
