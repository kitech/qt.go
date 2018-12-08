package qtwidgets

// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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

// bool eventFilter(QObject *, QEvent *)
func (this *QMdiSubWindow) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(QEvent *)
func (this *QMdiSubWindow) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void showEvent(QShowEvent *)
func (this *QMdiSubWindow) InheritShowEvent(f func(showEvent *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QMdiSubWindow) InheritHideEvent(f func(hideEvent *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void changeEvent(QEvent *)
func (this *QMdiSubWindow) InheritChangeEvent(f func(changeEvent *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(QCloseEvent *)
func (this *QMdiSubWindow) InheritCloseEvent(f func(closeEvent *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void leaveEvent(QEvent *)
func (this *QMdiSubWindow) InheritLeaveEvent(f func(leaveEvent *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QMdiSubWindow) InheritResizeEvent(f func(resizeEvent *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QMdiSubWindow) InheritTimerEvent(f func(timerEvent *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void moveEvent(QMoveEvent *)
func (this *QMdiSubWindow) InheritMoveEvent(f func(moveEvent *qtgui.QMoveEvent /*777 QMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QMdiSubWindow) InheritPaintEvent(f func(paintEvent *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QMdiSubWindow) InheritMousePressEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QMdiSubWindow) InheritMouseDoubleClickEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QMdiSubWindow) InheritMouseReleaseEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QMdiSubWindow) InheritMouseMoveEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QMdiSubWindow) InheritKeyPressEvent(f func(keyEvent *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QMdiSubWindow) InheritContextMenuEvent(f func(contextMenuEvent *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QMdiSubWindow) InheritFocusInEvent(f func(focusInEvent *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QMdiSubWindow) InheritFocusOutEvent(f func(focusOutEvent *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void childEvent(QChildEvent *)
func (this *QMdiSubWindow) InheritChildEvent(f func(childEvent *qtcore.QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

/*

 */
type QMdiSubWindow struct {
	*QWidget
}
type QMdiSubWindow_ITF interface {
	QWidget_ITF
	QMdiSubWindow_PTR() *QMdiSubWindow
}

func (ptr *QMdiSubWindow) QMdiSubWindow_PTR() *QMdiSubWindow { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMdiSubWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMdiSubWindow(QWidget *, Qt::WindowFlags)

/*
Constructs a new QMdiSubWindow widget. The parent and flags arguments are passed to QWidget's constructor.

Instead of using addSubWindow(), it is also simply possible to use setParent() when you add the subwindow to a QMdiArea.

Note that only QMdiSubWindows can be set as children of QMdiArea; you cannot, for instance, write:


  //bad code
  QMdiArea mdiArea;
  QTextEdit editor(&mdiArea); // invalid child widget



See also QMdiArea::addSubWindow().
*/
func (*QMdiSubWindow) NewForInherit(parent QWidget_ITF /*777 QWidget **/, flags int) *QMdiSubWindow {
	return NewQMdiSubWindow(parent, flags)
}
func NewQMdiSubWindow(parent QWidget_ITF /*777 QWidget **/, flags int) *QMdiSubWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMdiSubWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMdiSubWindow(QWidget *, Qt::WindowFlags)

/*
Constructs a new QMdiSubWindow widget. The parent and flags arguments are passed to QWidget's constructor.

Instead of using addSubWindow(), it is also simply possible to use setParent() when you add the subwindow to a QMdiArea.

Note that only QMdiSubWindows can be set as children of QMdiArea; you cannot, for instance, write:


  //bad code
  QMdiArea mdiArea;
  QTextEdit editor(&mdiArea); // invalid child widget



See also QMdiArea::addSubWindow().
*/
func (*QMdiSubWindow) NewForInheritp() *QMdiSubWindow {
	return NewQMdiSubWindowp()
}
func NewQMdiSubWindowp() *QMdiSubWindow {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMdiSubWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMdiSubWindow(QWidget *, Qt::WindowFlags)

/*
Constructs a new QMdiSubWindow widget. The parent and flags arguments are passed to QWidget's constructor.

Instead of using addSubWindow(), it is also simply possible to use setParent() when you add the subwindow to a QMdiArea.

Note that only QMdiSubWindows can be set as children of QMdiArea; you cannot, for instance, write:


  //bad code
  QMdiArea mdiArea;
  QTextEdit editor(&mdiArea); // invalid child widget



See also QMdiArea::addSubWindow().
*/
func (*QMdiSubWindow) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QMdiSubWindow {
	return NewQMdiSubWindowp1(parent)
}
func NewQMdiSubWindowp1(parent QWidget_ITF /*777 QWidget **/) *QMdiSubWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMdiSubWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMdiSubWindow()

/*

 */
func DeleteQMdiSubWindow(this *QMdiSubWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QMdiSubWindow) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QMdiSubWindow) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*
Sets widget as the internal widget of this subwindow. The internal widget is displayed in the center of the subwindow beneath the title bar.

QMdiSubWindow takes temporary ownership of widget; you do not have to delete it. Any existing internal widget will be removed and reparented to the root window.

See also widget().
*/
func (this *QMdiSubWindow) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the current internal widget.

See also setWidget().
*/
func (this *QMdiSubWindow) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * maximizedButtonsWidget() const

/*

 */
func (this *QMdiSubWindow) MaximizedButtonsWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * maximizedSystemMenuIconWidget() const

/*

 */
func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isShaded() const

/*
Returns true if this window is shaded; otherwise returns false.

A window is shaded if it is collapsed so that only the title bar is visible.
*/
func (this *QMdiSubWindow) IsShaded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow8isShadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QMdiSubWindow::SubWindowOption, bool)

/*
If on is true, option is enabled on the subwindow; otherwise it is disabled. See SubWindowOption for the effect of each option.

See also SubWindowOption and testOption().
*/
func (this *QMdiSubWindow) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9setOptionENS_15SubWindowOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QMdiSubWindow::SubWindowOption, bool)

/*
If on is true, option is enabled on the subwindow; otherwise it is disabled. See SubWindowOption for the effect of each option.

See also SubWindowOption and testOption().
*/
func (this *QMdiSubWindow) SetOptionp(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9setOptionENS_15SubWindowOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QMdiSubWindow::SubWindowOption) const

/*
Returns true if option is enabled; otherwise returns false.

See also SubWindowOption and setOption().
*/
func (this *QMdiSubWindow) TestOption(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10testOptionENS_15SubWindowOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardSingleStep(int)

/*

 */
func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow21setKeyboardSingleStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyboardSingleStep() const

/*

 */
func (this *QMdiSubWindow) KeyboardSingleStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow18keyboardSingleStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardPageStep(int)

/*

 */
func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow19setKeyboardPageStepEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyboardPageStep() const

/*

 */
func (this *QMdiSubWindow) KeyboardPageStep() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow16keyboardPageStepEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSystemMenu(QMenu *)

/*
Sets systemMenu as the current system menu for this subwindow.

By default, each QMdiSubWindow has a standard system menu.

QActions for the system menu created by QMdiSubWindow will automatically be updated depending on the current window state; e.g., the minimize action will be disabled after the window is minimized.

QActions added by the user are not updated by QMdiSubWindow.

QMdiSubWindow takes ownership of systemMenu; you do not have to delete it. Any existing menus will be deleted.

See also systemMenu() and showSystemMenu().
*/
func (this *QMdiSubWindow) SetSystemMenu(systemMenu QMenu_ITF /*777 QMenu **/) {
	var convArg0 unsafe.Pointer
	if systemMenu != nil && systemMenu.QMenu_PTR() != nil {
		convArg0 = systemMenu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13setSystemMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * systemMenu() const

/*
Returns a pointer to the current system menu, or zero if no system menu is set. QMdiSubWindow provides a default system menu, but you can also set the menu with setSystemMenu().

See also setSystemMenu() and showSystemMenu().
*/
func (this *QMdiSubWindow) SystemMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow10systemMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QMdiArea * mdiArea() const

/*
Returns the area containing this sub-window, or 0 if there is none.

This function was introduced in  Qt 4.4.

See also QMdiArea::addSubWindow().
*/
func (this *QMdiSubWindow) MdiArea() *QMdiArea /*777 QMdiArea **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMdiSubWindow7mdiAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMdiAreaFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowStateChanged(Qt::WindowStates, Qt::WindowStates)

/*
QMdiSubWindow emits this signal after the window state changes. oldState is the window state before it changed, and newState is the new, current state.
*/
func (this *QMdiSubWindow) WindowStateChanged(oldState int, newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow18windowStateChangedE6QFlagsIN2Qt11WindowStateEES3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldState, newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToActivate()

/*
QMdiSubWindow emits this signal immediately before it is activated. After the subwindow has been activated, the QMdiArea that manages the subwindow will also emit the subWindowActivated() signal.

See also QMdiArea::subWindowActivated().
*/
func (this *QMdiSubWindow) AboutToActivate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow15aboutToActivateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showSystemMenu()

/*
Shows the system menu below the system menu icon in the title bar.

See also setSystemMenu() and systemMenu().
*/
func (this *QMdiSubWindow) ShowSystemMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow14showSystemMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showShaded()

/*
Calling this function makes the subwindow enter the shaded mode. When the subwindow is shaded, only the title bar is visible.

Although shading is not supported by all styles, this function will still show the subwindow as shaded, regardless of whether support for shading is available. However, when used with styles without shading support, the user will be unable to return from shaded mode through the user interface (e.g., through a shade button in the title bar).

See also isShaded().
*/
func (this *QMdiSubWindow) ShowShaded() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10showShadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QMdiSubWindow) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QMdiSubWindow) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QMdiSubWindow) ShowEvent(showEvent qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if showEvent != nil && showEvent.QShowEvent_PTR() != nil {
		convArg0 = showEvent.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QMdiSubWindow) HideEvent(hideEvent qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if hideEvent != nil && hideEvent.QHideEvent_PTR() != nil {
		convArg0 = hideEvent.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QMdiSubWindow) ChangeEvent(changeEvent qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if changeEvent != nil && changeEvent.QEvent_PTR() != nil {
		convArg0 = changeEvent.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:115
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)

/*
Reimplemented from QWidget::closeEvent().
*/
func (this *QMdiSubWindow) CloseEvent(closeEvent qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if closeEvent != nil && closeEvent.QCloseEvent_PTR() != nil {
		convArg0 = closeEvent.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)

/*
Reimplemented from QWidget::leaveEvent().
*/
func (this *QMdiSubWindow) LeaveEvent(leaveEvent qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if leaveEvent != nil && leaveEvent.QEvent_PTR() != nil {
		convArg0 = leaveEvent.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().

Warning: When maximizing or restoring a subwindow, the resulting call to this function may have an invalid QResizeEvent::oldSize().
*/
func (this *QMdiSubWindow) ResizeEvent(resizeEvent qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if resizeEvent != nil && resizeEvent.QResizeEvent_PTR() != nil {
		convArg0 = resizeEvent.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QMdiSubWindow) TimerEvent(timerEvent qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if timerEvent != nil && timerEvent.QTimerEvent_PTR() != nil {
		convArg0 = timerEvent.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)

/*
Reimplemented from QWidget::moveEvent().
*/
func (this *QMdiSubWindow) MoveEvent(moveEvent qtgui.QMoveEvent_ITF /*777 QMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if moveEvent != nil && moveEvent.QMoveEvent_PTR() != nil {
		convArg0 = moveEvent.QMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QMdiSubWindow) PaintEvent(paintEvent qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if paintEvent != nil && paintEvent.QPaintEvent_PTR() != nil {
		convArg0 = paintEvent.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QMdiSubWindow) MousePressEvent(mouseEvent qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if mouseEvent != nil && mouseEvent.QMouseEvent_PTR() != nil {
		convArg0 = mouseEvent.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().
*/
func (this *QMdiSubWindow) MouseDoubleClickEvent(mouseEvent qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if mouseEvent != nil && mouseEvent.QMouseEvent_PTR() != nil {
		convArg0 = mouseEvent.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QMdiSubWindow) MouseReleaseEvent(mouseEvent qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if mouseEvent != nil && mouseEvent.QMouseEvent_PTR() != nil {
		convArg0 = mouseEvent.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QMdiSubWindow) MouseMoveEvent(mouseEvent qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if mouseEvent != nil && mouseEvent.QMouseEvent_PTR() != nil {
		convArg0 = mouseEvent.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QMdiSubWindow) KeyPressEvent(keyEvent qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if keyEvent != nil && keyEvent.QKeyEvent_PTR() != nil {
		convArg0 = keyEvent.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:127
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QMdiSubWindow) ContextMenuEvent(contextMenuEvent qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if contextMenuEvent != nil && contextMenuEvent.QContextMenuEvent_PTR() != nil {
		convArg0 = contextMenuEvent.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:129
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QMdiSubWindow) FocusInEvent(focusInEvent qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if focusInEvent != nil && focusInEvent.QFocusEvent_PTR() != nil {
		convArg0 = focusInEvent.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:130
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QMdiSubWindow) FocusOutEvent(focusOutEvent qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if focusOutEvent != nil && focusOutEvent.QFocusEvent_PTR() != nil {
		convArg0 = focusOutEvent.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)

/*
Reimplemented from QObject::childEvent().
*/
func (this *QMdiSubWindow) ChildEvent(childEvent qtcore.QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if childEvent != nil && childEvent.QChildEvent_PTR() != nil {
		convArg0 = childEvent.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMdiSubWindow10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QMdiSubWindow__SubWindowOption = int

//
const QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = 1

//
const QMdiSubWindow__AllowOutsideAreaVertically QMdiSubWindow__SubWindowOption = 2

//
const QMdiSubWindow__RubberBandResize QMdiSubWindow__SubWindowOption = 4

//
const QMdiSubWindow__RubberBandMove QMdiSubWindow__SubWindowOption = 8

func (this *QMdiSubWindow) SubWindowOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMdiSubWindow_SubWindowOptionItemName(val int) string {
	var nilthis *QMdiSubWindow
	return nilthis.SubWindowOptionItemName(val)
}

//  body block end

//  keep block begin

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
