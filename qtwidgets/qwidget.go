package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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

// bool event(QEvent *)
func (this *QWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QWidget) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QWidget) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QWidget) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QWidget) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QWidget) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QWidget) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QWidget) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QWidget) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QWidget) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void enterEvent(QEvent *)
func (this *QWidget) InheritEnterEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterEvent", f)
}

// void leaveEvent(QEvent *)
func (this *QWidget) InheritLeaveEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QWidget) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void moveEvent(QMoveEvent *)
func (this *QWidget) InheritMoveEvent(f func(event *qtgui.QMoveEvent /*777 QMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QWidget) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void closeEvent(QCloseEvent *)
func (this *QWidget) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QWidget) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void tabletEvent(QTabletEvent *)
func (this *QWidget) InheritTabletEvent(f func(event *qtgui.QTabletEvent /*777 QTabletEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabletEvent", f)
}

// void actionEvent(QActionEvent *)
func (this *QWidget) InheritActionEvent(f func(event *qtgui.QActionEvent /*777 QActionEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QWidget) InheritDragEnterEvent(f func(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QWidget) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QWidget) InheritDragLeaveEvent(f func(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QWidget) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QWidget) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QWidget) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// bool nativeEvent(const QByteArray &, void *, long *)
func (this *QWidget) InheritNativeEvent(f func(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool) {
	qtrt.SetAllInheritCallback(this, "nativeEvent", f)
}

// void changeEvent(QEvent *)
func (this *QWidget) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QWidget) InheritMetric(f func(arg0 int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// void initPainter(QPainter *)
func (this *QWidget) InheritInitPainter(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initPainter", f)
}

// QPaintDevice * redirected(QPoint *)
func (this *QWidget) InheritRedirected(f func(offset *qtcore.QPoint /*777 QPoint **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "redirected", f)
}

// QPainter * sharedPainter()
func (this *QWidget) InheritSharedPainter(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "sharedPainter", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QWidget) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void updateMicroFocus()
func (this *QWidget) InheritUpdateMicroFocus(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateMicroFocus", f)
}

// void create(WId, bool, bool)
func (this *QWidget) InheritCreate(f func(arg0 uint64, initializeWindow bool, destroyOldWindow bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "create", f)
}

// void destroy(bool, bool)
func (this *QWidget) InheritDestroy(f func(destroyWindow bool, destroySubWindows bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "destroy", f)
}

// bool focusNextPrevChild(bool)
func (this *QWidget) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// bool focusNextChild()
func (this *QWidget) InheritFocusNextChild(f func() bool) {
	qtrt.SetAllInheritCallback(this, "focusNextChild", f)
}

// bool focusPreviousChild()
func (this *QWidget) InheritFocusPreviousChild(f func() bool) {
	qtrt.SetAllInheritCallback(this, "focusPreviousChild", f)
}

/*

 */
type QWidget struct {
	*qtcore.QObject
	*qtgui.QPaintDevice
}
type QWidget_ITF interface {
	qtcore.QObject_ITF
	qtgui.QPaintDevice_ITF
	QWidget_PTR() *QWidget
}

func (ptr *QWidget) QWidget_PTR() *QWidget { return ptr }

func (this *QWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWidget) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QPaintDevice = qtgui.NewQPaintDeviceFromPointer(cthis)
}
func NewQWidgetFromPointer(cthis unsafe.Pointer) *QWidget {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := qtgui.NewQPaintDeviceFromPointer(cthis)
	return &QWidget{bcthis0, bcthis1}
}
func (*QWidget) NewFromPointer(cthis unsafe.Pointer) *QWidget {
	return NewQWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwidget.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)

/*
Constructs a widget which is a child of parent, with widget flags set to f.

If parent is 0, the new widget becomes a window. If parent is another widget, this widget becomes a child window inside parent. The new widget is deleted when its parent is deleted.

The widget flags argument, f, is normally 0, but it can be set to customize the frame of a window (i.e. parent must be 0). To customize the frame, use a value composed from the bitwise OR of any of the window flags.

If you add a child widget to an already visible widget you must explicitly show the child to make it visible.

Note that the X11 version of Qt may not be able to deliver all combinations of style flags on all systems. This is because on X11, Qt can only ask the window manager, and the window manager can override the application's settings. On Windows, Qt can set whatever flags you want.

See also windowFlags.
*/
func (*QWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/, f int) *QWidget {
	return NewQWidget(parent, f)
}
func NewQWidget(parent QWidget_ITF /*777 QWidget **/, f int) *QWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)

/*
Constructs a widget which is a child of parent, with widget flags set to f.

If parent is 0, the new widget becomes a window. If parent is another widget, this widget becomes a child window inside parent. The new widget is deleted when its parent is deleted.

The widget flags argument, f, is normally 0, but it can be set to customize the frame of a window (i.e. parent must be 0). To customize the frame, use a value composed from the bitwise OR of any of the window flags.

If you add a child widget to an already visible widget you must explicitly show the child to make it visible.

Note that the X11 version of Qt may not be able to deliver all combinations of style flags on all systems. This is because on X11, Qt can only ask the window manager, and the window manager can override the application's settings. On Windows, Qt can set whatever flags you want.

See also windowFlags.
*/
func (*QWidget) NewForInheritp() *QWidget {
	return NewQWidgetp()
}
func NewQWidgetp() *QWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)

/*
Constructs a widget which is a child of parent, with widget flags set to f.

If parent is 0, the new widget becomes a window. If parent is another widget, this widget becomes a child window inside parent. The new widget is deleted when its parent is deleted.

The widget flags argument, f, is normally 0, but it can be set to customize the frame of a window (i.e. parent must be 0). To customize the frame, use a value composed from the bitwise OR of any of the window flags.

If you add a child widget to an already visible widget you must explicitly show the child to make it visible.

Note that the X11 version of Qt may not be able to deliver all combinations of style flags on all systems. This is because on X11, Qt can only ask the window manager, and the window manager can override the application's settings. On Windows, Qt can set whatever flags you want.

See also windowFlags.
*/
func (*QWidget) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QWidget {
	return NewQWidgetp1(parent)
}
func NewQWidgetp1(parent QWidget_ITF /*777 QWidget **/) *QWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidget()

/*

 */
func DeleteQWidget(this *QWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qwidget.h:217
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const

/*

 */
func (this *QWidget) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] WId winId() const

/*
Returns the window system identifier of the widget.

Portable in principle, but if you use it you are probably about to do something non-portable. Be careful.

If a widget is non-native (alien) and winId() is invoked on it, that widget will be provided a native handle.

This value may change at run-time. An event with type QEvent::WinIdChange will be sent to the widget following a change in window system identifier.

See also find().
*/
func (this *QWidget) WinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5winIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void createWinId()

/*

 */
func (this *QWidget) CreateWinId() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11createWinIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [8] WId internalWinId() const

/*

 */
func (this *QWidget) InternalWinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13internalWinIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] WId effectiveWinId() const

/*
Returns the effective window system identifier of the widget, i.e. the native parent's window system identifier.

If the widget is native, this function returns the native widget ID. Otherwise, the window ID of the first native parent widget, i.e., the top-level widget that contains this widget, is returned.

Note: We recommend that you do not store this value as it is likely to change at run-time.

This function was introduced in  Qt 4.4.

See also nativeParentWidget().
*/
func (this *QWidget) EffectiveWinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14effectiveWinIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:225
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * style() const

/*
See also QWidget::setStyle(), QApplication::setStyle(), and QApplication::style().
*/
func (this *QWidget) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)

/*
Sets the widget's GUI style to style. The ownership of the style object is not transferred.

If no style is set, the widget uses the application's style, QApplication::style() instead.

Setting a widget's style has no effect on existing or future child widgets.

Warning: This function is particularly useful for demonstration purposes, where you want to show Qt's styling capabilities. Real applications should avoid it and use one consistent GUI style instead.

Warning: Qt style sheets are currently not supported for custom QStyle subclasses. We plan to address this in some future release.

See also style(), QStyle, QApplication::style(), and QApplication::setStyle().
*/
func (this *QWidget) SetStyle(arg0 QStyle_ITF /*777 QStyle **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStyle_PTR() != nil {
		convArg0 = arg0.QStyle_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTopLevel() const

/*

 */
func (this *QWidget) IsTopLevel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10isTopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:230
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWindow() const

/*
Returns true if the widget is an independent window, otherwise returns false.

A window is a widget that isn't visually the child of any other widget and that usually has a frame and a window title.

A window can have a parent widget. It will then be grouped with its parent and deleted when the parent is deleted, minimized when the parent is minimized etc. If supported by the window manager, it will also have a common taskbar entry with its parent.

QDialog and QMainWindow widgets are by default windows, even if a parent widget is specified in the constructor. This behavior is specified by the Qt::Window flag.

See also window(), isModal(), and parentWidget().
*/
func (this *QWidget) IsWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8isWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:232
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModal() const

/*

 */
func (this *QWidget) IsModal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7isModalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:233
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowModality windowModality() const

/*

 */
func (this *QWidget) WindowModality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14windowModalityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowModality(Qt::WindowModality)

/*

 */
func (this *QWidget) SetWindowModality(windowModality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowModalityEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), windowModality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:236
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*

 */
func (this *QWidget) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:237
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabledTo(const QWidget *) const

/*
Returns true if this widget would become enabled if ancestor is enabled; otherwise returns false.

This is the case if neither the widget itself nor every parent up to but excluding ancestor has been explicitly disabled.

isEnabledTo(0) returns false if this widget or any if its ancestors was explicitly disabled.

The word ancestor here means a parent widget within the same window.

Therefore isEnabledTo(0) stops at this widget's window, unlike isEnabled() which also takes parent windows into considerations.

See also setEnabled() and enabled.
*/
func (this *QWidget) IsEnabledTo(arg0 QWidget_ITF /*777 const QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isEnabledToEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:238
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabledToTLW() const

/*

 */
func (this *QWidget) IsEnabledToTLW() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14isEnabledToTLWEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QWidget) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDisabled(bool)

/*
Disables widget input events if disable is true; otherwise enables input events.

See the enabled documentation for more information.

See also isEnabledTo(), QKeyEvent, QMouseEvent, and changeEvent().
*/
func (this *QWidget) SetDisabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowModified(bool)

/*

 */
func (this *QWidget) SetWindowModified(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:248
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameGeometry() const

/*

 */
func (this *QWidget) FrameGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13frameGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:249
// index:0
// Public Visibility=Default Availability=Available
// [16] const QRect & geometry() const

/*

 */
func (this *QWidget) Geometry() *qtcore.QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:250
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect normalGeometry() const

/*

 */
func (this *QWidget) NormalGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14normalGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:252
// index:0
// Public Visibility=Default Availability=Available
// [4] int x() const

/*

 */
func (this *QWidget) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int y() const

/*

 */
func (this *QWidget) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:254
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint pos() const

/*

 */
func (this *QWidget) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:255
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize frameSize() const

/*

 */
func (this *QWidget) FrameSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9frameSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:256
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*

 */
func (this *QWidget) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width() const

/*

 */
func (this *QWidget) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:258
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height() const

/*

 */
func (this *QWidget) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:259
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect rect() const

/*

 */
func (this *QWidget) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:260
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect childrenRect() const

/*

 */
func (this *QWidget) ChildrenRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12childrenRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:261
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion childrenRegion() const

/*

 */
func (this *QWidget) ChildrenRegion() *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14childrenRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*

 */
func (this *QWidget) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*

 */
func (this *QWidget) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:265
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumWidth() const

/*

 */
func (this *QWidget) MinimumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12minimumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:266
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumHeight() const

/*

 */
func (this *QWidget) MinimumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13minimumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumWidth() const

/*

 */
func (this *QWidget) MaximumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12maximumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumHeight() const

/*

 */
func (this *QWidget) MaximumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13maximumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSize &)

/*

 */
func (this *QWidget) SetMinimumSize(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:270
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(int, int)

/*

 */
func (this *QWidget) SetMinimumSize1(minw int, minh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minw, minh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSize &)

/*

 */
func (this *QWidget) SetMaximumSize(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:272
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(int, int)

/*

 */
func (this *QWidget) SetMaximumSize1(maxw int, maxh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxw, maxh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(int)

/*

 */
func (this *QWidget) SetMinimumWidth(minw int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15setMinimumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minw)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(int)

/*

 */
func (this *QWidget) SetMinimumHeight(minh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setMinimumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(int)

/*

 */
func (this *QWidget) SetMaximumWidth(maxw int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15setMaximumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxw)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(int)

/*

 */
func (this *QWidget) SetMaximumHeight(maxh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setMaximumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:282
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeIncrement() const

/*

 */
func (this *QWidget) SizeIncrement() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13sizeIncrementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeIncrement(const QSize &)

/*

 */
func (this *QWidget) SetSizeIncrement(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:284
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSizeIncrement(int, int)

/*

 */
func (this *QWidget) SetSizeIncrement1(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:285
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize baseSize() const

/*

 */
func (this *QWidget) BaseSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8baseSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:286
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseSize(const QSize &)

/*

 */
func (this *QWidget) SetBaseSize(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:287
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBaseSize(int, int)

/*

 */
func (this *QWidget) SetBaseSize1(basew int, baseh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), basew, baseh)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedSize(const QSize &)

/*
Sets both the minimum and maximum sizes of the widget to s, thereby preventing it from ever growing or shrinking.

This will override the default size constraints set by QLayout.

To remove constraints, set the size to QWIDGETSIZE_MAX.

Alternatively, if you want the widget to have a fixed size based on its contents, you can call QLayout::setSizeConstraint(QLayout::SetFixedSize);

See also maximumSize and minimumSize.
*/
func (this *QWidget) SetFixedSize(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:290
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFixedSize(int, int)

/*
Sets both the minimum and maximum sizes of the widget to s, thereby preventing it from ever growing or shrinking.

This will override the default size constraints set by QLayout.

To remove constraints, set the size to QWIDGETSIZE_MAX.

Alternatively, if you want the widget to have a fixed size based on its contents, you can call QLayout::setSizeConstraint(QLayout::SetFixedSize);

See also maximumSize and minimumSize.
*/
func (this *QWidget) SetFixedSize1(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedWidth(int)

/*
Sets both the minimum and maximum width of the widget to w without changing the heights. Provided for convenience.

See also sizeHint(), minimumSize(), maximumSize(), and setFixedSize().
*/
func (this *QWidget) SetFixedWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setFixedWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedHeight(int)

/*
Sets both the minimum and maximum heights of the widget to h without changing the widths. Provided for convenience.

See also sizeHint(), minimumSize(), maximumSize(), and setFixedSize().
*/
func (this *QWidget) SetFixedHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setFixedHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:296
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapToGlobal(const QPoint &) const

/*
Translates the widget coordinate pos to global screen coordinates. For example, mapToGlobal(QPoint(0,0)) would give the global coordinates of the top-left pixel of the widget.

See also mapFromGlobal(), mapTo(), and mapToParent().
*/
func (this *QWidget) MapToGlobal(arg0 qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11mapToGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:297
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromGlobal(const QPoint &) const

/*
Translates the global screen coordinate pos to widget coordinates.

See also mapToGlobal(), mapFrom(), and mapFromParent().
*/
func (this *QWidget) MapFromGlobal(arg0 qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13mapFromGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:298
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapToParent(const QPoint &) const

/*
Translates the widget coordinate pos to a coordinate in the parent widget.

Same as mapToGlobal() if the widget has no parent.

See also mapFromParent(), mapTo(), mapToGlobal(), and underMouse().
*/
func (this *QWidget) MapToParent(arg0 qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11mapToParentERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromParent(const QPoint &) const

/*
Translates the parent widget coordinate pos to widget coordinates.

Same as mapFromGlobal() if the widget has no parent.

See also mapToParent(), mapFrom(), mapFromGlobal(), and underMouse().
*/
func (this *QWidget) MapFromParent(arg0 qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13mapFromParentERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapTo(const QWidget *, const QPoint &) const

/*
Translates the widget coordinate pos to the coordinate system of parent. The parent must not be 0 and must be a parent of the calling widget.

See also mapFrom(), mapToParent(), mapToGlobal(), and underMouse().
*/
func (this *QWidget) MapTo(arg0 QWidget_ITF /*777 const QWidget **/, arg1 qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QPoint_PTR() != nil {
		convArg1 = arg1.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5mapToEPKS_RK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFrom(const QWidget *, const QPoint &) const

/*
Translates the widget coordinate pos from the coordinate system of parent to this widget's coordinate system. The parent must not be 0 and must be a parent of the calling widget.

See also mapTo(), mapFromParent(), mapFromGlobal(), and underMouse().
*/
func (this *QWidget) MapFrom(arg0 QWidget_ITF /*777 const QWidget **/, arg1 qtcore.QPoint_ITF) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QPoint_PTR() != nil {
		convArg1 = arg1.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7mapFromEPKS_RK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * window() const

/*
Returns the window for this widget, i.e. the next ancestor widget that has (or could have) a window-system frame.

If the widget is a window, the widget itself is returned.

Typical usage is changing the window title:


  aWidget->window()->setWindowTitle("New Window Title");



See also isWindow().
*/
func (this *QWidget) Window() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:304
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * nativeParentWidget() const

/*
Returns the native parent for this widget, i.e. the next ancestor widget that has a system identifier, or 0 if it does not have any native parent.

This function was introduced in  Qt 4.4.

See also effectiveWinId().
*/
func (this *QWidget) NativeParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget18nativeParentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:305
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWidget * topLevelWidget() const

/*

 */
func (this *QWidget) TopLevelWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14topLevelWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:308
// index:0
// Public Visibility=Default Availability=Available
// [16] const QPalette & palette() const

/*

 */
func (this *QWidget) Palette() *qtgui.QPalette {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7paletteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:309
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)

/*

 */
func (this *QWidget) SetPalette(arg0 qtgui.QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPalette_PTR() != nil {
		convArg0 = arg0.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundRole(QPalette::ColorRole)

/*
Sets the background role of the widget to role.

The background role defines the brush from the widget's palette that is used to render the background.

If role is QPalette::NoRole, then the widget inherits its parent's background role.

Note that styles are free to choose any color from the palette. You can modify the palette or set a style sheet if you don't achieve the result you want with setBackgroundRole().

See also backgroundRole() and foregroundRole().
*/
func (this *QWidget) SetBackgroundRole(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setBackgroundRoleEN8QPalette9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:312
// index:0
// Public Visibility=Default Availability=Available
// [4] QPalette::ColorRole backgroundRole() const

/*
Returns the background role of the widget.

The background role defines the brush from the widget's palette that is used to render the background.

If no explicit background role is set, the widget inherts its parent widget's background role.

See also setBackgroundRole() and foregroundRole().
*/
func (this *QWidget) BackgroundRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14backgroundRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setForegroundRole(QPalette::ColorRole)

/*
Sets the foreground role of the widget to role.

The foreground role defines the color from the widget's palette that is used to draw the foreground.

If role is QPalette::NoRole, the widget uses a foreground role that contrasts with the background role.

Note that styles are free to choose any color from the palette. You can modify the palette or set a style sheet if you don't achieve the result you want with setForegroundRole().

See also foregroundRole() and backgroundRole().
*/
func (this *QWidget) SetForegroundRole(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setForegroundRoleEN8QPalette9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:315
// index:0
// Public Visibility=Default Availability=Available
// [4] QPalette::ColorRole foregroundRole() const

/*
Returns the foreground role.

The foreground role defines the color from the widget's palette that is used to draw the foreground.

If no explicit foreground role is set, the function returns a role that contrasts with the background role.

See also setForegroundRole() and backgroundRole().
*/
func (this *QWidget) ForegroundRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14foregroundRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:317
// index:0
// Public Visibility=Default Availability=Available
// [16] const QFont & font() const

/*

 */
func (this *QWidget) Font() *qtgui.QFont {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:318
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QWidget) SetFont(arg0 qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetrics fontMetrics() const

/*
Returns the font metrics for the widget's current font. Equivalent to QFontMetrics(widget->font()).

See also font(), fontInfo(), and setFont().
*/
func (this *QWidget) FontMetrics() *qtgui.QFontMetrics /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11fontMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFontMetrics)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:320
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontInfo fontInfo() const

/*
Returns the font info for the widget's current font. Equivalent to QFontInfo(widget->font()).

See also font(), fontMetrics(), and setFont().
*/
func (this *QWidget) FontInfo() *qtgui.QFontInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8fontInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFontInfo)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:323
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor() const

/*

 */
func (this *QWidget) Cursor() *qtgui.QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:324
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)

/*

 */
func (this *QWidget) SetCursor(arg0 qtgui.QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCursor_PTR() != nil {
		convArg0 = arg0.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:325
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()

/*

 */
func (this *QWidget) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMouseTracking(bool)

/*

 */
func (this *QWidget) SetMouseTracking(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setMouseTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:329
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasMouseTracking() const

/*

 */
func (this *QWidget) HasMouseTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16hasMouseTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:330
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underMouse() const

/*
Returns true if the widget is under the mouse cursor; otherwise returns false.

This value is not updated properly during drag and drop operations.

See also enterEvent() and leaveEvent().
*/
func (this *QWidget) UnderMouse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10underMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabletTracking(bool)

/*

 */
func (this *QWidget) SetTabletTracking(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setTabletTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:333
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasTabletTracking() const

/*

 */
func (this *QWidget) HasTabletTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget17hasTabletTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:335
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QBitmap &)

/*
Causes only the pixels of the widget for which bitmap has a corresponding 1 bit to be visible. If the region includes pixels outside the rect() of the widget, window system controls in that area may or may not be visible, depending on the platform.

Note that this effect can be slow if the region is particularly complex.

The following code shows how an image with an alpha channel can be used to generate a mask for a widget:


      QLabel topLevelLabel;
      QPixmap pixmap(":/images/tux.png");
      topLevelLabel.setPixmap(pixmap);
      topLevelLabel.setMask(pixmap.mask());



The label shown by this code is masked using the image it contains, giving the appearance that an irregularly-shaped image is being drawn directly onto the screen.

Masked widgets receive mouse events only on their visible portions.

See also mask(), clearMask(), windowOpacity(), and Shaped Clock Example.
*/
func (this *QWidget) SetMask(arg0 qtgui.QBitmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QBitmap_PTR() != nil {
		convArg0 = arg0.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QBitmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:336
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QRegion &)

/*
Causes only the pixels of the widget for which bitmap has a corresponding 1 bit to be visible. If the region includes pixels outside the rect() of the widget, window system controls in that area may or may not be visible, depending on the platform.

Note that this effect can be slow if the region is particularly complex.

The following code shows how an image with an alpha channel can be used to generate a mask for a widget:


      QLabel topLevelLabel;
      QPixmap pixmap(":/images/tux.png");
      topLevelLabel.setPixmap(pixmap);
      topLevelLabel.setMask(pixmap.mask());



The label shown by this code is masked using the image it contains, giving the appearance that an irregularly-shaped image is being drawn directly onto the screen.

Masked widgets receive mouse events only on their visible portions.

See also mask(), clearMask(), windowOpacity(), and Shaped Clock Example.
*/
func (this *QWidget) SetMask1(arg0 qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:337
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion mask() const

/*
Returns the mask currently set on a widget. If no mask is set the return value will be an empty region.

See also setMask(), clearMask(), QRegion::isEmpty(), and Shaped Clock Example.
*/
func (this *QWidget) Mask() *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMask()

/*
Removes any mask set by setMask().

See also setMask().
*/
func (this *QWidget) ClearMask() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9clearMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPaintDevice *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Render(target qtgui.QPaintDevice_ITF /*777 QPaintDevice **/, targetOffset qtcore.QPoint_ITF, sourceRegion qtgui.QRegion_ITF, renderFlags int) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QPaintDevice_PTR() != nil {
		convArg0 = target.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetOffset != nil && targetOffset.QPoint_PTR() != nil {
		convArg1 = targetOffset.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRegion != nil && sourceRegion.QRegion_PTR() != nil {
		convArg2 = sourceRegion.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP12QPaintDeviceRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPaintDevice *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Renderp(target qtgui.QPaintDevice_ITF /*777 QPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QPaintDevice_PTR() != nil {
		convArg0 = target.QPaintDevice_PTR().GetCthis()
	}
	// arg: 1, const QPoint &=LValueReference, QPoint=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QRegion &=LValueReference, QRegion=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QWidget::RenderFlags=Typedef, QWidget::RenderFlags=Typedef, QFlags<QWidget::RenderFlag>, Unexposed
	renderFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP12QPaintDeviceRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPaintDevice *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Renderp1(target qtgui.QPaintDevice_ITF /*777 QPaintDevice **/, targetOffset qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QPaintDevice_PTR() != nil {
		convArg0 = target.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetOffset != nil && targetOffset.QPoint_PTR() != nil {
		convArg1 = targetOffset.QPoint_PTR().GetCthis()
	}
	// arg: 2, const QRegion &=LValueReference, QRegion=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QWidget::RenderFlags=Typedef, QWidget::RenderFlags=Typedef, QFlags<QWidget::RenderFlag>, Unexposed
	renderFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP12QPaintDeviceRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPaintDevice *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Renderp2(target qtgui.QPaintDevice_ITF /*777 QPaintDevice **/, targetOffset qtcore.QPoint_ITF, sourceRegion qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QPaintDevice_PTR() != nil {
		convArg0 = target.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetOffset != nil && targetOffset.QPoint_PTR() != nil {
		convArg1 = targetOffset.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRegion != nil && sourceRegion.QRegion_PTR() != nil {
		convArg2 = sourceRegion.QRegion_PTR().GetCthis()
	}
	// arg: 3, QWidget::RenderFlags=Typedef, QWidget::RenderFlags=Typedef, QFlags<QWidget::RenderFlag>, Unexposed
	renderFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP12QPaintDeviceRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:344
// index:1
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Render1(painter qtgui.QPainter_ITF /*777 QPainter **/, targetOffset qtcore.QPoint_ITF, sourceRegion qtgui.QRegion_ITF, renderFlags int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetOffset != nil && targetOffset.QPoint_PTR() != nil {
		convArg1 = targetOffset.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRegion != nil && sourceRegion.QRegion_PTR() != nil {
		convArg2 = sourceRegion.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP8QPainterRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:344
// index:1
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Render1p(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 1, const QPoint &=LValueReference, QPoint=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QRegion &=LValueReference, QRegion=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QWidget::RenderFlags=Typedef, QWidget::RenderFlags=Typedef, QFlags<QWidget::RenderFlag>, Unexposed
	renderFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP8QPainterRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:344
// index:1
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Render1p1(painter qtgui.QPainter_ITF /*777 QPainter **/, targetOffset qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetOffset != nil && targetOffset.QPoint_PTR() != nil {
		convArg1 = targetOffset.QPoint_PTR().GetCthis()
	}
	// arg: 2, const QRegion &=LValueReference, QRegion=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QWidget::RenderFlags=Typedef, QWidget::RenderFlags=Typedef, QFlags<QWidget::RenderFlag>, Unexposed
	renderFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP8QPainterRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:344
// index:1
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QPoint &, const QRegion &, QWidget::RenderFlags)

/*
Renders the sourceRegion of this widget into the target using renderFlags to determine how to render. Rendering starts at targetOffset in the target. For example:


  QPixmap pixmap(widget->size());
  widget->render(&pixmap);



If sourceRegion is a null region, this function will use QWidget::rect() as the region, i.e. the entire widget.

Ensure that you call QPainter::end() for the target device's active painter (if any) before rendering. For example:


  QPainter painter(this);
  ...
  painter.end();
  myWidget->render(this);



Note: To obtain the contents of a QOpenGLWidget, use QOpenGLWidget::grabFramebuffer() instead.

Note: To obtain the contents of a QGLWidget (deprecated), use QGLWidget::grabFrameBuffer() or QGLWidget::renderPixmap() instead.

This function was introduced in  Qt 4.3.
*/
func (this *QWidget) Render1p2(painter qtgui.QPainter_ITF /*777 QPainter **/, targetOffset qtcore.QPoint_ITF, sourceRegion qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetOffset != nil && targetOffset.QPoint_PTR() != nil {
		convArg1 = targetOffset.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if sourceRegion != nil && sourceRegion.QRegion_PTR() != nil {
		convArg2 = sourceRegion.QRegion_PTR().GetCthis()
	}
	// arg: 3, QWidget::RenderFlags=Typedef, QWidget::RenderFlags=Typedef, QFlags<QWidget::RenderFlag>, Unexposed
	renderFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP8QPainterRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:348
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grab(const QRect &)

/*
Renders the widget into a pixmap restricted by the given rectangle. If the widget has any children, then they are also painted in the appropriate positions.

If a rectangle with an invalid size is specified (the default), the entire widget is painted.

This function was introduced in  Qt 5.0.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also render() and QPixmap.
*/
func (this *QWidget) Grab(rectangle qtcore.QRect_ITF) *qtgui.QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if rectangle != nil && rectangle.QRect_PTR() != nil {
		convArg0 = rectangle.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4grabERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:348
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grab(const QRect &)

/*
Renders the widget into a pixmap restricted by the given rectangle. If the widget has any children, then they are also painted in the appropriate positions.

If a rectangle with an invalid size is specified (the default), the entire widget is painted.

This function was introduced in  Qt 5.0.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also render() and QPixmap.
*/
func (this *QWidget) Grabp() *qtgui.QPixmap /*123*/ {
	// arg: 0, const QRect &=LValueReference, QRect=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4grabERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:356
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabGesture(Qt::GestureType, Qt::GestureFlags)

/*
Subscribes the widget to a given gesture with specific flags.

This function was introduced in  Qt 4.6.

See also ungrabGesture() and QGestureEvent.
*/
func (this *QWidget) GrabGesture(type_ int, flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11grabGestureEN2Qt11GestureTypeE6QFlagsINS0_11GestureFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:356
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabGesture(Qt::GestureType, Qt::GestureFlags)

/*
Subscribes the widget to a given gesture with specific flags.

This function was introduced in  Qt 4.6.

See also ungrabGesture() and QGestureEvent.
*/
func (this *QWidget) GrabGesturep(type_ int) {
	// arg: 1, Qt::GestureFlags=Elaborated, Qt::GestureFlags=Typedef, QFlags<Qt::GestureFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11grabGestureEN2Qt11GestureTypeE6QFlagsINS0_11GestureFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:357
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabGesture(Qt::GestureType)

/*
Unsubscribes the widget from a given gesture type

This function was introduced in  Qt 4.6.

See also grabGesture() and QGestureEvent.
*/
func (this *QWidget) UngrabGesture(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13ungrabGestureEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:361
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowTitle(const QString &)

/*

 */
func (this *QWidget) SetWindowTitle(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setWindowTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:363
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleSheet(const QString &)

/*

 */
func (this *QWidget) SetStyleSheet(styleSheet string) {
	var tmpArg0 = qtcore.NewQString5(styleSheet)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setStyleSheetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:367
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleSheet() const

/*

 */
func (this *QWidget) StyleSheet() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10styleSheetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:369
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowTitle() const

/*

 */
func (this *QWidget) WindowTitle() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11windowTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:370
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowIcon(const QIcon &)

/*

 */
func (this *QWidget) SetWindowIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:371
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon windowIcon() const

/*

 */
func (this *QWidget) WindowIcon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10windowIconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:372
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowIconText(const QString &)

/*

 */
func (this *QWidget) SetWindowIconText(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowIconTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:373
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowIconText() const

/*

 */
func (this *QWidget) WindowIconText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14windowIconTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:374
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowRole(const QString &)

/*
Sets the window's role to role. This only makes sense for windows on X11.

See also windowRole().
*/
func (this *QWidget) SetWindowRole(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowRoleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:375
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowRole() const

/*
Returns the window's role, or an empty string.

See also setWindowRole(), windowIcon, and windowTitle.
*/
func (this *QWidget) WindowRole() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10windowRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:376
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFilePath(const QString &)

/*

 */
func (this *QWidget) SetWindowFilePath(filePath string) {
	var tmpArg0 = qtcore.NewQString5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:377
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowFilePath() const

/*

 */
func (this *QWidget) WindowFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14windowFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowOpacity(qreal)

/*

 */
func (this *QWidget) SetWindowOpacity(level float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setWindowOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:380
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal windowOpacity() const

/*

 */
func (this *QWidget) WindowOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13windowOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:382
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWindowModified() const

/*

 */
func (this *QWidget) IsWindowModified() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16isWindowModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QWidget) SetToolTip(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:385
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QWidget) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:386
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTipDuration(int)

/*

 */
func (this *QWidget) SetToolTipDuration(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setToolTipDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:387
// index:0
// Public Visibility=Default Availability=Available
// [4] int toolTipDuration() const

/*

 */
func (this *QWidget) ToolTipDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15toolTipDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:398
// index:0
// Public Visibility=Default Availability=Available
// [8] QString accessibleName() const

/*

 */
func (this *QWidget) AccessibleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14accessibleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:399
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccessibleName(const QString &)

/*

 */
func (this *QWidget) SetAccessibleName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setAccessibleNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:400
// index:0
// Public Visibility=Default Availability=Available
// [8] QString accessibleDescription() const

/*

 */
func (this *QWidget) AccessibleDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget21accessibleDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:401
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccessibleDescription(const QString &)

/*

 */
func (this *QWidget) SetAccessibleDescription(description string) {
	var tmpArg0 = qtcore.NewQString5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget24setAccessibleDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:404
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)

/*

 */
func (this *QWidget) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:405
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection() const

/*

 */
func (this *QWidget) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:406
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetLayoutDirection()

/*

 */
func (this *QWidget) UnsetLayoutDirection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget20unsetLayoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:408
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)

/*

 */
func (this *QWidget) SetLocale(locale qtcore.QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:409
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale() const

/*

 */
func (this *QWidget) Locale() *qtcore.QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:410
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetLocale()

/*

 */
func (this *QWidget) UnsetLocale() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11unsetLocaleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:412
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRightToLeft() const

/*

 */
func (this *QWidget) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:413
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLeftToRight() const

/*

 */
func (this *QWidget) IsLeftToRight() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13isLeftToRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:416
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFocus()

/*
Gives the keyboard input focus to this widget (or its focus proxy) if this widget or one of its parents is the active window. The reason argument will be passed into any focus event sent from this function, it is used to give an explanation of what caused the widget to get focus. If the window is not active, the widget will be given the focus when the window becomes active.

First, a focus about to change event is sent to the focus widget (if any) to tell it that it is about to lose the focus. Then focus is changed, a focus out event is sent to the previous focus item and a focus in event is sent to the new item to tell it that it just received the focus. (Nothing happens if the focus in and focus out widgets are the same.)

Note: On embedded platforms, setFocus() will not cause an input panel to be opened by the input method. If you want this to happen, you have to send a QEvent::RequestSoftwareInputPanel event to the widget yourself.

setFocus() gives focus to a widget regardless of its focus policy, but does not clear any keyboard grab (see grabKeyboard()).

Be aware that if the widget is hidden, it will not accept focus until it is shown.

Warning: If you call setFocus() in a function which may itself be called from focusOutEvent() or focusInEvent(), you may get an infinite recursion.

See also hasFocus(), clearFocus(), focusInEvent(), focusOutEvent(), setFocusPolicy(), focusWidget(), QApplication::focusWidget(), grabKeyboard(), grabMouse(), Keyboard Focus in Widgets, and QEvent::RequestSoftwareInputPanel.
*/
func (this *QWidget) SetFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget8setFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:423
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)

/*
Gives the keyboard input focus to this widget (or its focus proxy) if this widget or one of its parents is the active window. The reason argument will be passed into any focus event sent from this function, it is used to give an explanation of what caused the widget to get focus. If the window is not active, the widget will be given the focus when the window becomes active.

First, a focus about to change event is sent to the focus widget (if any) to tell it that it is about to lose the focus. Then focus is changed, a focus out event is sent to the previous focus item and a focus in event is sent to the new item to tell it that it just received the focus. (Nothing happens if the focus in and focus out widgets are the same.)

Note: On embedded platforms, setFocus() will not cause an input panel to be opened by the input method. If you want this to happen, you have to send a QEvent::RequestSoftwareInputPanel event to the widget yourself.

setFocus() gives focus to a widget regardless of its focus policy, but does not clear any keyboard grab (see grabKeyboard()).

Be aware that if the widget is hidden, it will not accept focus until it is shown.

Warning: If you call setFocus() in a function which may itself be called from focusOutEvent() or focusInEvent(), you may get an infinite recursion.

See also hasFocus(), clearFocus(), focusInEvent(), focusOutEvent(), setFocusPolicy(), focusWidget(), QApplication::focusWidget(), grabKeyboard(), grabMouse(), Keyboard Focus in Widgets, and QEvent::RequestSoftwareInputPanel.
*/
func (this *QWidget) SetFocus1(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:419
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActiveWindow() const

/*

 */
func (this *QWidget) IsActiveWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14isActiveWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:420
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activateWindow()

/*
Sets the top-level widget containing this widget to be the active window.

An active window is a visible top-level window that has the keyboard input focus.

This function performs the same operation as clicking the mouse on the title bar of a top-level window. On X11, the result depends on the Window Manager. If you want to ensure that the window is stacked on top as well you should also call raise(). Note that the window must be visible, otherwise activateWindow() has no effect.

On Windows, if you are calling this when the application is not currently the active one then it will not make it the active window. It will change the color of the taskbar entry to indicate that the window has changed in some way. This is because Microsoft does not allow an application to interrupt what the user is currently doing in another application.

See also isActiveWindow(), window(), show(), and QWindowsWindowFunctions::setWindowActivationBehavior().
*/
func (this *QWidget) ActivateWindow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14activateWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:421
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFocus()

/*
Takes keyboard input focus from the widget.

If the widget has active focus, a focus out event is sent to this widget to tell it that it has lost the focus.

This widget must enable focus setting in order to get the keyboard input focus, i.e. it must call setFocusPolicy().

See also hasFocus(), setFocus(), focusInEvent(), focusOutEvent(), setFocusPolicy(), and QApplication::focusWidget().
*/
func (this *QWidget) ClearFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10clearFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:424
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FocusPolicy focusPolicy() const

/*

 */
func (this *QWidget) FocusPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11focusPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:425
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusPolicy(Qt::FocusPolicy)

/*

 */
func (this *QWidget) SetFocusPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setFocusPolicyEN2Qt11FocusPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:426
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus() const

/*

 */
func (this *QWidget) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:427
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTabOrder(QWidget *, QWidget *)

/*
Puts the second widget after the first widget in the focus order.

It effectively removes the second widget from its focus chain and inserts it after the first widget.

Note that since the tab order of the second widget is changed, you should order a chain like this:


  setTabOrder(a, b); // a to b
  setTabOrder(b, c); // a to b to c
  setTabOrder(c, d); // a to b to c to d



not like this:


  // WRONG
  setTabOrder(c, d); // c to d
  setTabOrder(a, b); // a to b AND c to d
  setTabOrder(b, c); // a to b to c, but not c to d



If first or second has a focus proxy, setTabOrder() correctly substitutes the proxy.

Note: Since Qt 5.10: A widget that has a child as focus proxy is understood as a compound widget. When setting a tab order between one or two compound widgets, the local tab order inside each will be preserved. This means that if both widgets are compound widgets, the resulting tab order will be from the last child inside first, to the first child inside second.

See also setFocusPolicy(), setFocusProxy(), and Keyboard Focus in Widgets.
*/
func (this *QWidget) SetTabOrder(arg0 QWidget_ITF /*777 QWidget **/, arg1 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QWidget_PTR() != nil {
		convArg1 = arg1.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setTabOrderEPS_S0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QWidget_SetTabOrder(arg0 QWidget_ITF /*777 QWidget **/, arg1 QWidget_ITF /*777 QWidget **/) {
	var nilthis *QWidget
	nilthis.SetTabOrder(arg0, arg1)
}

// /usr/include/qt/QtWidgets/qwidget.h:428
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusProxy(QWidget *)

/*
Sets the widget's focus proxy to widget w. If w is 0, the function resets this widget to have no focus proxy.

Some widgets can "have focus", but create a child widget, such as QLineEdit, to actually handle the focus. In this case, the widget can set the line edit to be its focus proxy.

setFocusProxy() sets the widget which will actually get focus when "this widget" gets it. If there is a focus proxy, setFocus() and hasFocus() operate on the focus proxy.

See also focusProxy().
*/
func (this *QWidget) SetFocusProxy(arg0 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setFocusProxyEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:429
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * focusProxy() const

/*
Returns the focus proxy, or 0 if there is no focus proxy.

See also setFocusProxy().
*/
func (this *QWidget) FocusProxy() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10focusProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:430
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ContextMenuPolicy contextMenuPolicy() const

/*

 */
func (this *QWidget) ContextMenuPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget17contextMenuPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:431
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextMenuPolicy(Qt::ContextMenuPolicy)

/*

 */
func (this *QWidget) SetContextMenuPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget20setContextMenuPolicyEN2Qt17ContextMenuPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:434
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabMouse()

/*
Grabs the mouse input.

This widget receives all mouse events until releaseMouse() is called; other widgets get no mouse events at all. Keyboard events are not affected. Use grabKeyboard() if you want to grab that.

Warning: Bugs in mouse-grabbing applications very often lock the terminal. Use this function with extreme caution, and consider using the -nograb command line option while debugging.

It is almost never necessary to grab the mouse when using Qt, as Qt grabs and releases it sensibly. In particular, Qt grabs the mouse when a mouse button is pressed and keeps it until the last button is released.

Note: Only visible widgets can grab mouse input. If isVisible() returns false for a widget, that widget cannot call grabMouse().

Note: On Windows, grabMouse() only works when the mouse is inside a window owned by the process. On macOS, grabMouse() only works when the mouse is inside the frame of that widget.

See also releaseMouse(), grabKeyboard(), and releaseKeyboard().
*/
func (this *QWidget) GrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9grabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:436
// index:1
// Public Visibility=Default Availability=Available
// [-2] void grabMouse(const QCursor &)

/*
Grabs the mouse input.

This widget receives all mouse events until releaseMouse() is called; other widgets get no mouse events at all. Keyboard events are not affected. Use grabKeyboard() if you want to grab that.

Warning: Bugs in mouse-grabbing applications very often lock the terminal. Use this function with extreme caution, and consider using the -nograb command line option while debugging.

It is almost never necessary to grab the mouse when using Qt, as Qt grabs and releases it sensibly. In particular, Qt grabs the mouse when a mouse button is pressed and keeps it until the last button is released.

Note: Only visible widgets can grab mouse input. If isVisible() returns false for a widget, that widget cannot call grabMouse().

Note: On Windows, grabMouse() only works when the mouse is inside a window owned by the process. On macOS, grabMouse() only works when the mouse is inside the frame of that widget.

See also releaseMouse(), grabKeyboard(), and releaseKeyboard().
*/
func (this *QWidget) GrabMouse1(arg0 qtgui.QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCursor_PTR() != nil {
		convArg0 = arg0.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9grabMouseERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:438
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseMouse()

/*
Releases the mouse grab.

See also grabMouse(), grabKeyboard(), and releaseKeyboard().
*/
func (this *QWidget) ReleaseMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12releaseMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:439
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabKeyboard()

/*
Grabs the keyboard input.

This widget receives all keyboard events until releaseKeyboard() is called; other widgets get no keyboard events at all. Mouse events are not affected. Use grabMouse() if you want to grab that.

The focus widget is not affected, except that it doesn't receive any keyboard events. setFocus() moves the focus as usual, but the new focus widget receives keyboard events only after releaseKeyboard() is called.

If a different widget is currently grabbing keyboard input, that widget's grab is released first.

See also releaseKeyboard(), grabMouse(), releaseMouse(), and focusWidget().
*/
func (this *QWidget) GrabKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12grabKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:440
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseKeyboard()

/*
Releases the keyboard grab.

See also grabKeyboard(), grabMouse(), and releaseMouse().
*/
func (this *QWidget) ReleaseKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15releaseKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:442
// index:0
// Public Visibility=Default Availability=Available
// [4] int grabShortcut(const QKeySequence &, Qt::ShortcutContext)

/*
Adds a shortcut to Qt's shortcut system that watches for the given key sequence in the given context. If the context is Qt::ApplicationShortcut, the shortcut applies to the application as a whole. Otherwise, it is either local to this widget, Qt::WidgetShortcut, or to the window itself, Qt::WindowShortcut.

If the same key sequence has been grabbed by several widgets, when the key sequence occurs a QEvent::Shortcut event is sent to all the widgets to which it applies in a non-deterministic order, but with the ``ambiguous'' flag set to true.

Warning: You should not normally need to use this function; instead create QActions with the shortcut key sequences you require (if you also want equivalent menu options and toolbar buttons), or create QShortcuts if you just need key sequences. Both QAction and QShortcut handle all the event filtering for you, and provide signals which are triggered when the user triggers the key sequence, so are much easier to use than this low-level function.

See also releaseShortcut() and setShortcutEnabled().
*/
func (this *QWidget) GrabShortcut(key qtgui.QKeySequence_ITF, context int) int {
	var convArg0 unsafe.Pointer
	if key != nil && key.QKeySequence_PTR() != nil {
		convArg0 = key.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:442
// index:0
// Public Visibility=Default Availability=Available
// [4] int grabShortcut(const QKeySequence &, Qt::ShortcutContext)

/*
Adds a shortcut to Qt's shortcut system that watches for the given key sequence in the given context. If the context is Qt::ApplicationShortcut, the shortcut applies to the application as a whole. Otherwise, it is either local to this widget, Qt::WidgetShortcut, or to the window itself, Qt::WindowShortcut.

If the same key sequence has been grabbed by several widgets, when the key sequence occurs a QEvent::Shortcut event is sent to all the widgets to which it applies in a non-deterministic order, but with the ``ambiguous'' flag set to true.

Warning: You should not normally need to use this function; instead create QActions with the shortcut key sequences you require (if you also want equivalent menu options and toolbar buttons), or create QShortcuts if you just need key sequences. Both QAction and QShortcut handle all the event filtering for you, and provide signals which are triggered when the user triggers the key sequence, so are much easier to use than this low-level function.

See also releaseShortcut() and setShortcutEnabled().
*/
func (this *QWidget) GrabShortcutp(key qtgui.QKeySequence_ITF) int {
	var convArg0 unsafe.Pointer
	if key != nil && key.QKeySequence_PTR() != nil {
		convArg0 = key.QKeySequence_PTR().GetCthis()
	}
	// arg: 1, Qt::ShortcutContext=Elaborated, Qt::ShortcutContext=Enum, , Invalid
	context := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:443
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseShortcut(int)

/*
Removes the shortcut with the given id from Qt's shortcut system. The widget will no longer receive QEvent::Shortcut events for the shortcut's key sequence (unless it has other shortcuts with the same key sequence).

Warning: You should not normally need to use this function since Qt's shortcut system removes shortcuts automatically when their parent widget is destroyed. It is best to use QAction or QShortcut to handle shortcuts, since they are easier to use than this low-level function. Note also that this is an expensive operation.

See also grabShortcut() and setShortcutEnabled().
*/
func (this *QWidget) ReleaseShortcut(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15releaseShortcutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:444
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutEnabled(int, bool)

/*
If enable is true, the shortcut with the given id is enabled; otherwise the shortcut is disabled.

Warning: You should not normally need to use this function since Qt's shortcut system enables/disables shortcuts automatically as widgets become hidden/visible and gain or lose focus. It is best to use QAction or QShortcut to handle shortcuts, since they are easier to use than this low-level function.

See also grabShortcut() and releaseShortcut().
*/
func (this *QWidget) SetShortcutEnabled(id int, enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setShortcutEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:444
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutEnabled(int, bool)

/*
If enable is true, the shortcut with the given id is enabled; otherwise the shortcut is disabled.

Warning: You should not normally need to use this function since Qt's shortcut system enables/disables shortcuts automatically as widgets become hidden/visible and gain or lose focus. It is best to use QAction or QShortcut to handle shortcuts, since they are easier to use than this low-level function.

See also grabShortcut() and releaseShortcut().
*/
func (this *QWidget) SetShortcutEnabledp(id int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enable := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setShortcutEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:445
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutAutoRepeat(int, bool)

/*
If enable is true, auto repeat of the shortcut with the given id is enabled; otherwise it is disabled.

This function was introduced in  Qt 4.2.

See also grabShortcut() and releaseShortcut().
*/
func (this *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21setShortcutAutoRepeatEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:445
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutAutoRepeat(int, bool)

/*
If enable is true, auto repeat of the shortcut with the given id is enabled; otherwise it is disabled.

This function was introduced in  Qt 4.2.

See also grabShortcut() and releaseShortcut().
*/
func (this *QWidget) SetShortcutAutoRepeatp(id int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enable := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21setShortcutAutoRepeatEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:447
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * mouseGrabber()

/*
Returns the widget that is currently grabbing the mouse input.

If no widget in this application is currently grabbing the mouse, 0 is returned.

See also grabMouse() and keyboardGrabber().
*/
func (this *QWidget) MouseGrabber() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12mouseGrabberEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWidget_MouseGrabber() *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.MouseGrabber()
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:448
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * keyboardGrabber()

/*
Returns the widget that is currently grabbing the keyboard input.

If no widget in this application is currently grabbing the keyboard, 0 is returned.

See also grabMouse() and mouseGrabber().
*/
func (this *QWidget) KeyboardGrabber() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15keyboardGrabberEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWidget_KeyboardGrabber() *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.KeyboardGrabber()
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:451
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool updatesEnabled() const

/*

 */
func (this *QWidget) UpdatesEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14updatesEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:452
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUpdatesEnabled(bool)

/*

 */
func (this *QWidget) SetUpdatesEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setUpdatesEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:459
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()

/*
Updates the widget unless updates are disabled or the widget is hidden.

This function does not cause an immediate repaint; instead it schedules a paint event for processing when Qt returns to the main event loop. This permits Qt to optimize for more speed and less flicker than a call to repaint() does.

Calling update() several times normally results in just one paintEvent() call.

Qt normally erases the widget's area before the paintEvent() call. If the Qt::WA_OpaquePaintEvent widget attribute is set, the widget is responsible for painting all its pixels with an opaque color.

See also repaint(), paintEvent(), setUpdatesEnabled(), and Analog Clock Example.
*/
func (this *QWidget) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:463
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void update(int, int, int, int)

/*
Updates the widget unless updates are disabled or the widget is hidden.

This function does not cause an immediate repaint; instead it schedules a paint event for processing when Qt returns to the main event loop. This permits Qt to optimize for more speed and less flicker than a call to repaint() does.

Calling update() several times normally results in just one paintEvent() call.

Qt normally erases the widget's area before the paintEvent() call. If the Qt::WA_OpaquePaintEvent widget attribute is set, the widget is responsible for painting all its pixels with an opaque color.

See also repaint(), paintEvent(), setUpdatesEnabled(), and Analog Clock Example.
*/
func (this *QWidget) Update1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:464
// index:2
// Public Visibility=Default Availability=Available
// [-2] void update(const QRect &)

/*
Updates the widget unless updates are disabled or the widget is hidden.

This function does not cause an immediate repaint; instead it schedules a paint event for processing when Qt returns to the main event loop. This permits Qt to optimize for more speed and less flicker than a call to repaint() does.

Calling update() several times normally results in just one paintEvent() call.

Qt normally erases the widget's area before the paintEvent() call. If the Qt::WA_OpaquePaintEvent widget attribute is set, the widget is responsible for painting all its pixels with an opaque color.

See also repaint(), paintEvent(), setUpdatesEnabled(), and Analog Clock Example.
*/
func (this *QWidget) Update2(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:465
// index:3
// Public Visibility=Default Availability=Available
// [-2] void update(const QRegion &)

/*
Updates the widget unless updates are disabled or the widget is hidden.

This function does not cause an immediate repaint; instead it schedules a paint event for processing when Qt returns to the main event loop. This permits Qt to optimize for more speed and less flicker than a call to repaint() does.

Calling update() several times normally results in just one paintEvent() call.

Qt normally erases the widget's area before the paintEvent() call. If the Qt::WA_OpaquePaintEvent widget attribute is set, the widget is responsible for painting all its pixels with an opaque color.

See also repaint(), paintEvent(), setUpdatesEnabled(), and Analog Clock Example.
*/
func (this *QWidget) Update3(arg0 qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:460
// index:0
// Public Visibility=Default Availability=Available
// [-2] void repaint()

/*
Repaints the widget directly by calling paintEvent() immediately, unless updates are disabled or the widget is hidden.

We suggest only using repaint() if you need an immediate repaint, for example during animation. In almost all circumstances update() is better, as it permits Qt to optimize for speed and minimize flicker.

Warning: If you call repaint() in a function which may itself be called from paintEvent(), you may get infinite recursion. The update() function never causes recursion.

See also update(), paintEvent(), and setUpdatesEnabled().
*/
func (this *QWidget) Repaint() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:467
// index:1
// Public Visibility=Default Availability=Available
// [-2] void repaint(int, int, int, int)

/*
Repaints the widget directly by calling paintEvent() immediately, unless updates are disabled or the widget is hidden.

We suggest only using repaint() if you need an immediate repaint, for example during animation. In almost all circumstances update() is better, as it permits Qt to optimize for speed and minimize flicker.

Warning: If you call repaint() in a function which may itself be called from paintEvent(), you may get infinite recursion. The update() function never causes recursion.

See also update(), paintEvent(), and setUpdatesEnabled().
*/
func (this *QWidget) Repaint1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:468
// index:2
// Public Visibility=Default Availability=Available
// [-2] void repaint(const QRect &)

/*
Repaints the widget directly by calling paintEvent() immediately, unless updates are disabled or the widget is hidden.

We suggest only using repaint() if you need an immediate repaint, for example during animation. In almost all circumstances update() is better, as it permits Qt to optimize for speed and minimize flicker.

Warning: If you call repaint() in a function which may itself be called from paintEvent(), you may get infinite recursion. The update() function never causes recursion.

See also update(), paintEvent(), and setUpdatesEnabled().
*/
func (this *QWidget) Repaint2(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:469
// index:3
// Public Visibility=Default Availability=Available
// [-2] void repaint(const QRegion &)

/*
Repaints the widget directly by calling paintEvent() immediately, unless updates are disabled or the widget is hidden.

We suggest only using repaint() if you need an immediate repaint, for example during animation. In almost all circumstances update() is better, as it permits Qt to optimize for speed and minimize flicker.

Warning: If you call repaint() in a function which may itself be called from paintEvent(), you may get infinite recursion. The update() function never causes recursion.

See also update(), paintEvent(), and setUpdatesEnabled().
*/
func (this *QWidget) Repaint3(arg0 qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:474
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QWidget) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:475
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHidden(bool)

/*
Convenience function, equivalent to setVisible(!hidden).

See also isHidden().
*/
func (this *QWidget) SetHidden(hidden bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hidden)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:476
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()

/*
Shows the widget and its child widgets.

This is equivalent to calling showFullScreen(), showMaximized(), or setVisible(true), depending on the platform's default behavior for the window flags.

See also raise(), showEvent(), hide(), setVisible(), showMinimized(), showMaximized(), showNormal(), isVisible(), and windowFlags().
*/
func (this *QWidget) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:477
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()

/*
Hides the widget. This function is equivalent to setVisible(false).

Note: If you are working with QDialog or its subclasses and you invoke the show() function after this function, the dialog will be displayed in its original position.

See also hideEvent(), isHidden(), show(), setVisible(), isVisible(), and close().
*/
func (this *QWidget) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:479
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMinimized()

/*
Shows the widget minimized, as an icon.

Calling this function only affects windows.

See also showNormal(), showMaximized(), show(), hide(), isVisible(), and isMinimized().
*/
func (this *QWidget) ShowMinimized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13showMinimizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:480
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMaximized()

/*
Shows the widget maximized.

Calling this function only affects windows.

On X11, this function may not work properly with certain window managers. See the Window Geometry documentation for an explanation.

See also setWindowState(), showNormal(), showMinimized(), show(), hide(), and isVisible().
*/
func (this *QWidget) ShowMaximized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13showMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:481
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showFullScreen()

/*
Shows the widget in full-screen mode.

Calling this function only affects windows.

To return from full-screen mode, call showNormal().

Full-screen mode works fine under Windows, but has certain problems under X. These problems are due to limitations of the ICCCM protocol that specifies the communication between X11 clients and the window manager. ICCCM simply does not understand the concept of non-decorated full-screen windows. Therefore, the best we can do is to request a borderless window and place and resize it to fill the entire screen. Depending on the window manager, this may or may not work. The borderless window is requested using MOTIF hints, which are at least partially supported by virtually all modern window managers.

An alternative would be to bypass the window manager entirely and create a window with the Qt::X11BypassWindowManagerHint flag. This has other severe problems though, like totally broken keyboard focus and very strange effects on desktop changes or when the user raises other windows.

X11 window managers that follow modern post-ICCCM specifications support full-screen mode properly.

See also showNormal(), showMaximized(), show(), hide(), and isVisible().
*/
func (this *QWidget) ShowFullScreen() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14showFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:482
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNormal()

/*
Restores the widget after it has been maximized or minimized.

Calling this function only affects windows.

See also setWindowState(), showMinimized(), showMaximized(), show(), hide(), and isVisible().
*/
func (this *QWidget) ShowNormal() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10showNormalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:484
// index:0
// Public Visibility=Default Availability=Available
// [1] bool close()

/*
Closes this widget. Returns true if the widget was closed; otherwise returns false.

First it sends the widget a QCloseEvent. The widget is hidden if it accepts the close event. If it ignores the event, nothing happens. The default implementation of QWidget::closeEvent() accepts the close event.

If the widget has the Qt::WA_DeleteOnClose flag, the widget is also deleted. A close events is delivered to the widget no matter if the widget is visible or not.

The QApplication::lastWindowClosed() signal is emitted when the last visible primary window (i.e. window with no parent) with the Qt::WA_QuitOnClose attribute set is closed. By default this attribute is set for all widgets except transient windows such as splash screens, tool windows, and popup menus.
*/
func (this *QWidget) Close() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:485
// index:0
// Public Visibility=Default Availability=Available
// [-2] void raise()

/*
Raises this widget to the top of the parent widget's stack.

After this call the widget will be visually in front of any overlapping sibling widgets.

Note: When using activateWindow(), you can call this function to ensure that the window is stacked on top.

See also lower() and stackUnder().
*/
func (this *QWidget) Raise() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5raiseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:486
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lower()

/*
Lowers the widget to the bottom of the parent widget's stack.

After this call the widget will be visually behind (and therefore obscured by) any overlapping sibling widgets.

See also raise() and stackUnder().
*/
func (this *QWidget) Lower() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5lowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:489
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackUnder(QWidget *)

/*
Places the widget under w in the parent widget's stack.

To make this work, the widget itself and w must be siblings.

See also raise() and lower().
*/
func (this *QWidget) StackUnder(arg0 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10stackUnderEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:490
// index:0
// Public Visibility=Default Availability=Available
// [-2] void move(int, int)

/*

 */
func (this *QWidget) Move(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4moveEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:491
// index:1
// Public Visibility=Default Availability=Available
// [-2] void move(const QPoint &)

/*

 */
func (this *QWidget) Move1(arg0 qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4moveERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:492
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(int, int)

/*

 */
func (this *QWidget) Resize(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6resizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:493
// index:1
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSize &)

/*

 */
func (this *QWidget) Resize1(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6resizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:494
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setGeometry(int, int, int, int)

/*

 */
func (this *QWidget) SetGeometry(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setGeometryEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:495
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*

 */
func (this *QWidget) SetGeometry1(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:496
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveGeometry() const

/*
Saves the current geometry and state for top-level widgets.

To save the geometry when the window closes, you can implement a close event like this:


  void MyWidget::closeEvent(QCloseEvent *event)
  {
      QSettings settings("MyCompany", "MyApp");
      settings.setValue("geometry", saveGeometry());
      QWidget::closeEvent(event);
  }



See the Window Geometry documentation for an overview of geometry issues with windows.

Use QMainWindow::saveState() to save the geometry and the state of toolbars and dock widgets.

This function was introduced in  Qt 4.2.

See also restoreGeometry(), QMainWindow::saveState(), and QMainWindow::restoreState().
*/
func (this *QWidget) SaveGeometry() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12saveGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:497
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreGeometry(const QByteArray &)

/*
Restores the geometry and state of top-level widgets stored in the byte array geometry. Returns true on success; otherwise returns false.

If the restored geometry is off-screen, it will be modified to be inside the available screen geometry.

To restore geometry saved using QSettings, you can use code like this:


  QSettings settings("MyCompany", "MyApp");
  myWidget->restoreGeometry(settings.value("myWidget/geometry").toByteArray());



See the Window Geometry documentation for an overview of geometry issues with windows.

Use QMainWindow::restoreState() to restore the geometry and the state of toolbars and dock widgets.

This function was introduced in  Qt 4.2.

See also saveGeometry(), QSettings, QMainWindow::saveState(), and QMainWindow::restoreState().
*/
func (this *QWidget) RestoreGeometry(geometry qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if geometry != nil && geometry.QByteArray_PTR() != nil {
		convArg0 = geometry.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15restoreGeometryERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:498
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()

/*
Adjusts the size of the widget to fit its contents.

This function uses sizeHint() if it is valid, i.e., the size hint's width and height are >= 0. Otherwise, it sets the size to the children rectangle that covers all child widgets (the union of all child widget rectangles).

For windows, the screen size is also taken into account. If the sizeHint() is less than (200, 100) and the size policy is expanding, the window will be at least (200, 100). The maximum size of a window is 2/3 of the screen's width and height.

See also sizeHint() and childrenRect().
*/
func (this *QWidget) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:499
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QWidget) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:500
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisibleTo(const QWidget *) const

/*
Returns true if this widget would become visible if ancestor is shown; otherwise returns false.

The true case occurs if neither the widget itself nor any parent up to but excluding ancestor has been explicitly hidden.

This function will still return true if the widget is obscured by other windows on the screen, but could be physically visible if it or they were to be moved.

isVisibleTo(0) is identical to isVisible().

See also show(), hide(), and isVisible().
*/
func (this *QWidget) IsVisibleTo(arg0 QWidget_ITF /*777 const QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isVisibleToEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:501
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isHidden() const

/*
Returns true if the widget is hidden, otherwise returns false.

A hidden widget will only become visible when show() is called on it. It will not be automatically shown when the parent is shown.

To check visibility, use !isVisible() instead (notice the exclamation mark).

isHidden() implies !isVisible(), but a widget can be not visible and not hidden at the same time. This is the case for widgets that are children of widgets that are not visible.

Widgets are hidden if:


they were created as independent windows,
they were created as children of visible widgets,
hide() or setVisible(false) was called.
*/
func (this *QWidget) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:503
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMinimized() const

/*

 */
func (this *QWidget) IsMinimized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isMinimizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:504
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMaximized() const

/*

 */
func (this *QWidget) IsMaximized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:505
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFullScreen() const

/*

 */
func (this *QWidget) IsFullScreen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12isFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:507
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowStates windowState() const

/*
Returns the current window state. The window state is a OR'ed combination of Qt::WindowState: Qt::WindowMinimized, Qt::WindowMaximized, Qt::WindowFullScreen, and Qt::WindowActive.

See also Qt::WindowState and setWindowState().
*/
func (this *QWidget) WindowState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11windowStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:508
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowState(Qt::WindowStates)

/*
Sets the window state to windowState. The window state is a OR'ed combination of Qt::WindowState: Qt::WindowMinimized, Qt::WindowMaximized, Qt::WindowFullScreen, and Qt::WindowActive.

If the window is not visible (i.e. isVisible() returns false), the window state will take effect when show() is called. For visible windows, the change is immediate. For example, to toggle between full-screen and normal mode, use the following code:


  w->setWindowState(w->windowState() ^ Qt::WindowFullScreen);



In order to restore and activate a minimized window (while preserving its maximized and/or full-screen state), use the following:


  w->setWindowState((w->windowState() & ~Qt::WindowMinimized) | Qt::WindowActive);



Calling this function will hide the widget. You must call show() to make the widget visible again.

Note: On some window systems Qt::WindowActive is not immediate, and may be ignored in certain cases.

When the window state changes, the widget receives a changeEvent() of type QEvent::WindowStateChange.

See also Qt::WindowState and windowState().
*/
func (this *QWidget) SetWindowState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setWindowStateE6QFlagsIN2Qt11WindowStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:509
// index:0
// Public Visibility=Default Availability=Available
// [-2] void overrideWindowState(Qt::WindowStates)

/*

 */
func (this *QWidget) OverrideWindowState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget19overrideWindowStateE6QFlagsIN2Qt11WindowStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:511
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*

 */
func (this *QWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:512
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*

 */
func (this *QWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:514
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy sizePolicy() const

/*

 */
func (this *QWidget) SizePolicy() *QSizePolicy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10sizePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizePolicy)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:515
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy)

/*

 */
func (this *QWidget) SetSizePolicy(arg0 QSizePolicy_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSizePolicy_PTR() != nil {
		convArg0 = arg0.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyE11QSizePolicy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:516
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy)

/*

 */
func (this *QWidget) SetSizePolicy1(horizontal int, vertical int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyEN11QSizePolicy6PolicyES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:517
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Returns the preferred height for this widget, given the width w.

If this widget has a layout, the default implementation returns the layout's preferred height. if there is no layout, the default implementation returns -1 indicating that the preferred height does not depend on the width.
*/
func (this *QWidget) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:518
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Returns true if the widget's preferred height depends on its width; otherwise returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QWidget) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:520
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion visibleRegion() const

/*
Returns the unobscured region where paint events can occur.

For visible widgets, this is an approximation of the area not covered by other widgets; otherwise, this is an empty region.

The repaint() function calls this function if necessary, so in general you do not need to call it.
*/
func (this *QWidget) VisibleRegion() *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13visibleRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:522
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(int, int, int, int)

/*
Sets the margins around the contents of the widget to have the sizes left, top, right, and bottom. The margins are used by the layout system, and may be used by subclasses to specify the area to draw in (e.g. excluding the frame).

Changing the margins will trigger a resizeEvent().

See also contentsMargins(), contentsRect(), and getContentsMargins().
*/
func (this *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:523
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(const QMargins &)

/*
Sets the margins around the contents of the widget to have the sizes left, top, right, and bottom. The margins are used by the layout system, and may be used by subclasses to specify the area to draw in (e.g. excluding the frame).

Changing the margins will trigger a resizeEvent().

See also contentsMargins(), contentsRect(), and getContentsMargins().
*/
func (this *QWidget) SetContentsMargins1(margins qtcore.QMargins_ITF) {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:524
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getContentsMargins(int *, int *, int *, int *) const

/*
Returns the widget's contents margins for left, top, right, and bottom.

See also setContentsMargins() and contentsRect().
*/
func (this *QWidget) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:525
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins contentsMargins() const

/*
The contentsMargins function returns the widget's contents margins.

This function was introduced in  Qt 4.6.

See also getContentsMargins(), setContentsMargins(), and contentsRect().
*/
func (this *QWidget) ContentsMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15contentsMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:527
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect contentsRect() const

/*
Returns the area inside the widget's margins.

See also setContentsMargins() and getContentsMargins().
*/
func (this *QWidget) ContentsRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12contentsRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:530
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayout * layout() const

/*
Returns the layout manager that is installed on this widget, or 0 if no layout manager is installed.

The layout manager sets the geometry of the widget's children that have been added to the layout.

See also setLayout(), sizePolicy(), and Layout Management.
*/
func (this *QWidget) Layout() *QLayout /*777 QLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:531
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(QLayout *)

/*
Sets the layout manager for this widget to layout.

If there already is a layout manager installed on this widget, QWidget won't let you install another. You must first delete the existing layout manager (returned by layout()) before you can call setLayout() with the new layout.

If layout is the layout manager on a different widget, setLayout() will reparent the layout and make it the layout manager for this widget.

Example:


      QVBoxLayout *layout = new QVBoxLayout;
      layout->addWidget(formWidget);
      setLayout(layout);



An alternative to calling this function is to pass this widget to the layout's constructor.

The QWidget will take ownership of layout.

See also layout() and Layout Management.
*/
func (this *QWidget) SetLayout(arg0 QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QLayout_PTR() != nil {
		convArg0 = arg0.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setLayoutEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:532
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateGeometry()

/*
Notifies the layout system that this widget has changed and may need to change geometry.

Call this function if the sizeHint() or sizePolicy() have changed.

For explicitly hidden widgets, updateGeometry() is a no-op. The layout system will be notified as soon as the widget is shown.
*/
func (this *QWidget) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:534
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParent(QWidget *)

/*
Sets the parent of the widget to parent, and resets the window flags. The widget is moved to position (0, 0) in its new parent.

If the new parent widget is in a different window, the reparented widget and its children are appended to the end of the tab chain of the new parent widget, in the same internal order as before. If one of the moved widgets had keyboard focus, setParent() calls clearFocus() for that widget.

If the new parent widget is in the same window as the old parent, setting the parent doesn't change the tab order or keyboard focus.

If the "new" parent widget is the old parent widget, this function does nothing.

Note: The widget becomes invisible as part of changing its parent, even if it was previously visible. You must call show() to make the widget visible again.

Warning: It is very unlikely that you will ever need this function. If you have a widget that changes its content dynamically, it is far easier to use QStackedWidget.

See also setWindowFlags().
*/
func (this *QWidget) SetParent(parent QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:535
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setParent(QWidget *, Qt::WindowFlags)

/*
Sets the parent of the widget to parent, and resets the window flags. The widget is moved to position (0, 0) in its new parent.

If the new parent widget is in a different window, the reparented widget and its children are appended to the end of the tab chain of the new parent widget, in the same internal order as before. If one of the moved widgets had keyboard focus, setParent() calls clearFocus() for that widget.

If the new parent widget is in the same window as the old parent, setting the parent doesn't change the tab order or keyboard focus.

If the "new" parent widget is the old parent widget, this function does nothing.

Note: The widget becomes invisible as part of changing its parent, even if it was previously visible. You must call show() to make the widget visible again.

Warning: It is very unlikely that you will ever need this function. If you have a widget that changes its content dynamically, it is far easier to use QStackedWidget.

See also setWindowFlags().
*/
func (this *QWidget) SetParent1(parent QWidget_ITF /*777 QWidget **/, f int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, f)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:537
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int)

/*
Scrolls the widget including its children dx pixels to the right and dy downward. Both dx and dy may be negative.

After scrolling, the widgets will receive paint events for the areas that need to be repainted. For widgets that Qt knows to be opaque, this is only the newly exposed parts. For example, if an opaque widget is scrolled 8 pixels to the left, only an 8-pixel wide stripe at the right edge needs updating.

Since widgets propagate the contents of their parents by default, you need to set the autoFillBackground property, or use setAttribute() to set the Qt::WA_OpaquePaintEvent attribute, to make a widget opaque.

For widgets that use contents propagation, a scroll will cause an update of the entire scroll area.

See also Transparency and Double Buffering.
*/
func (this *QWidget) Scroll(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6scrollEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:538
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int, const QRect &)

/*
Scrolls the widget including its children dx pixels to the right and dy downward. Both dx and dy may be negative.

After scrolling, the widgets will receive paint events for the areas that need to be repainted. For widgets that Qt knows to be opaque, this is only the newly exposed parts. For example, if an opaque widget is scrolled 8 pixels to the left, only an 8-pixel wide stripe at the right edge needs updating.

Since widgets propagate the contents of their parents by default, you need to set the autoFillBackground property, or use setAttribute() to set the Qt::WA_OpaquePaintEvent attribute, to make a widget opaque.

For widgets that use contents propagation, a scroll will cause an update of the entire scroll area.

See also Transparency and Double Buffering.
*/
func (this *QWidget) Scroll1(dx int, dy int, arg2 qtcore.QRect_ITF) {
	var convArg2 unsafe.Pointer
	if arg2 != nil && arg2.QRect_PTR() != nil {
		convArg2 = arg2.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6scrollEiiRK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:542
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * focusWidget() const

/*
Returns the last child of this widget that setFocus had been called on. For top level widgets this is the widget that will get focus in case this window gets activated

This is not the same as QApplication::focusWidget(), which returns the focus widget in the currently active window.
*/
func (this *QWidget) FocusWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11focusWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:543
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * nextInFocusChain() const

/*
Returns the next widget in this widget's focus chain.

See also previousInFocusChain().
*/
func (this *QWidget) NextInFocusChain() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16nextInFocusChainEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:544
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * previousInFocusChain() const

/*
The previousInFocusChain function returns the previous widget in this widget's focus chain.

This function was introduced in  Qt 4.6.

See also nextInFocusChain().
*/
func (this *QWidget) PreviousInFocusChain() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget20previousInFocusChainEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:547
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptDrops() const

/*

 */
func (this *QWidget) AcceptDrops() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11acceptDropsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:548
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptDrops(bool)

/*

 */
func (this *QWidget) SetAcceptDrops(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setAcceptDropsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:552
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *)

/*
Appends the action action to this widget's list of actions.

All QWidgets have a list of QActions, however they can be represented graphically in many different ways. The default use of the QAction list (as returned by actions()) is to create a context QMenu.

A QWidget should only have one of each action and adding an action it already has will not cause the same action to be in the widget twice.

The ownership of action is not transferred to this QWidget.

See also removeAction(), insertAction(), actions(), and QMenu.
*/
func (this *QWidget) AddAction(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9addActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:560
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertAction(QAction *, QAction *)

/*
Inserts the action action to this widget's list of actions, before the action before. It appends the action if before is 0 or before is not a valid action for this widget.

A QWidget should only have one of each action.

See also removeAction(), addAction(), QMenu, contextMenuPolicy, and actions().
*/
func (this *QWidget) InsertAction(before QAction_ITF /*777 QAction **/, action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg1 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12insertActionEP7QActionS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:561
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAction(QAction *)

/*
Removes the action action from this widget's list of actions.

See also insertAction(), actions(), and insertAction().
*/
func (this *QWidget) RemoveAction(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12removeActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:565
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * parentWidget() const

/*
Returns the parent of this widget, or 0 if it does not have any parent widget.
*/
func (this *QWidget) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:567
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlags(Qt::WindowFlags)

/*

 */
func (this *QWidget) SetWindowFlags(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:568
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::WindowFlags windowFlags() const

/*

 */
func (this *QWidget) WindowFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11windowFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:569
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlag(Qt::WindowType, bool)

/*
Sets the window flag flag on this widget if on is true; otherwise clears the flag.

This function was introduced in  Qt 5.9.

See also setWindowFlags(), windowFlags(), and windowType().
*/
func (this *QWidget) SetWindowFlag(arg0 int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowFlagEN2Qt10WindowTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:569
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlag(Qt::WindowType, bool)

/*
Sets the window flag flag on this widget if on is true; otherwise clears the flag.

This function was introduced in  Qt 5.9.

See also setWindowFlags(), windowFlags(), and windowType().
*/
func (this *QWidget) SetWindowFlagp(arg0 int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowFlagEN2Qt10WindowTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:570
// index:0
// Public Visibility=Default Availability=Available
// [-2] void overrideWindowFlags(Qt::WindowFlags)

/*
Sets the window flags for the widget to flags, without telling the window system.

Warning: Do not call this function unless you really know what you're doing.

See also setWindowFlags().
*/
func (this *QWidget) OverrideWindowFlags(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget19overrideWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:572
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::WindowType windowType() const

/*
Returns the window type of this widget. This is identical to windowFlags() & Qt::WindowType_Mask.

See also windowFlags.
*/
func (this *QWidget) WindowType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10windowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:574
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * find(WId)

/*
Returns a pointer to the widget with window identifer/handle id.

The window identifier type depends on the underlying window system, see qwindowdefs.h for the actual definition. If there is no widget with this identifier, 0 is returned.
*/
func (this *QWidget) Find(arg0 uint64) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4findEy", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWidget_Find(arg0 uint64) *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.Find(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:575
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWidget * childAt(int, int) const

/*
Returns the visible child widget at the position (x, y) in the widget's coordinate system. If there is no visible child widget at the specified position, the function returns 0.
*/
func (this *QWidget) ChildAt(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7childAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:576
// index:1
// Public Visibility=Default Availability=Available
// [8] QWidget * childAt(const QPoint &) const

/*
Returns the visible child widget at the position (x, y) in the widget's coordinate system. If there is no visible child widget at the specified position, the function returns 0.
*/
func (this *QWidget) ChildAt1(p qtcore.QPoint_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7childAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:578
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::WidgetAttribute, bool)

/*
Sets the attribute attribute on this widget if on is true; otherwise clears the attribute.

See also testAttribute().
*/
func (this *QWidget) SetAttribute(arg0 int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setAttributeEN2Qt15WidgetAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:578
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::WidgetAttribute, bool)

/*
Sets the attribute attribute on this widget if on is true; otherwise clears the attribute.

See also testAttribute().
*/
func (this *QWidget) SetAttributep(arg0 int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setAttributeEN2Qt15WidgetAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:579
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool testAttribute(Qt::WidgetAttribute) const

/*
Returns true if attribute attribute is set on this widget; otherwise returns false.

See also setAttribute().
*/
func (this *QWidget) TestAttribute(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13testAttributeEN2Qt15WidgetAttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:581
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*
Reimplemented from QPaintDevice::paintEngine().

Returns the widget's paint engine.

Note that this function should not be called explicitly by the user, since it's meant for reimplementation purposes only. The function is called by Qt internally, and the default implementation may not always return a valid pointer.
*/
func (this *QWidget) PaintEngine() *qtgui.QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:583
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensurePolished() const

/*
Ensures that the widget and its children have been polished by QStyle (i.e., have a proper font and palette).

QWidget calls this function after it has been fully constructed but before it is shown the very first time. You can call this function if you want to ensure that the widget is polished before doing an operation, e.g., the correct font size might be needed in the widget's sizeHint() reimplementation. Note that this function is called from the default implementation of sizeHint().

Polishing is useful for final initialization that must happen after all constructors (from base classes as well as from subclasses) have been called.

If you need to change some settings when a widget is polished, reimplement event() and handle the QEvent::Polish event type.

Note: The function is declared const so that it can be called from other const functions (e.g., sizeHint()).

See also event().
*/
func (this *QWidget) EnsurePolished() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14ensurePolishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:585
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QWidget *) const

/*
Returns true if this widget is a parent, (or grandparent and so on to any level), of the given child, and both widgets are within the same window; otherwise returns false.
*/
func (this *QWidget) IsAncestorOf(child QWidget_ITF /*777 const QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if child != nil && child.QWidget_PTR() != nil {
		convArg0 = child.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12isAncestorOfEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:592
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoFillBackground() const

/*

 */
func (this *QWidget) AutoFillBackground() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget18autoFillBackgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:593
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFillBackground(bool)

/*

 */
func (this *QWidget) SetAutoFillBackground(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21setAutoFillBackgroundEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:595
// index:0
// Public Visibility=Default Availability=Available
// [8] QBackingStore * backingStore() const

/*
Returns the QBackingStore this widget will be drawn into.

This function was introduced in  Qt 5.0.
*/
func (this *QWidget) BackingStore() *qtgui.QBackingStore /*777 QBackingStore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12backingStoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQBackingStoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:597
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * windowHandle() const

/*
If this is a native widget, return the associated QWindow. Otherwise return null.

Native widgets include toplevel widgets, QGLWidget, and child widgets on which winId() was called.

This function was introduced in  Qt 5.0.

See also winId().
*/
func (this *QWidget) WindowHandle() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12windowHandleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:599
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * createWindowContainer(QWindow *, QWidget *, Qt::WindowFlags)

/*
Creates a QWidget that makes it possible to embed window into a QWidget-based application.

The window container is created as a child of parent and with window flags flags.

Once the window has been embedded into the container, the container will control the window's geometry and visibility. Explicit calls to QWindow::setGeometry(), QWindow::show() or QWindow::hide() on an embedded window is not recommended.

The container takes over ownership of window. The window can be removed from the window container with a call to QWindow::setParent().

The window container is attached as a native child window to the toplevel window it is a child of. When a window container is used as a child of a QAbstractScrollArea or QMdiArea, it will create a native window for every widget in its parent chain to allow for proper stacking and clipping in this use case. Creating a native window for the window container also allows for proper stacking and clipping. This must be done before showing the window container. Applications with many native child windows may suffer from performance issues.

The window container has a number of known limitations:


Stacking order; The embedded window will stack on top of the widget hierarchy as an opaque box. The stacking order of multiple overlapping window container instances is undefined.
Rendering Integration; The window container does not interoperate with QGraphicsProxyWidget, QWidget::render() or similar functionality.
Focus Handling; It is possible to let the window container instance have any focus policy and it will delegate focus to the window via a call to QWindow::requestActivate(). However, returning to the normal focus chain from the QWindow instance will be up to the QWindow instance implementation itself. For instance, when entering a Qt Quick based window with tab focus, it is quite likely that further tab presses will only cycle inside the QML application. Also, whether QWindow::requestActivate() actually gives the window focus, is platform dependent.
Using many window container instances in a QWidget-based application can greatly hurt the overall performance of the application.
*/
func (this *QWidget) CreateWindowContainer(window qtgui.QWindow_ITF /*777 QWindow **/, parent QWidget_ITF /*777 QWidget **/, flags int) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21createWindowContainerEP7QWindowPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWidget_CreateWindowContainer(window qtgui.QWindow_ITF /*777 QWindow **/, parent QWidget_ITF /*777 QWidget **/, flags int) *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.CreateWindowContainer(window, parent, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:599
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * createWindowContainer(QWindow *, QWidget *, Qt::WindowFlags)

/*
Creates a QWidget that makes it possible to embed window into a QWidget-based application.

The window container is created as a child of parent and with window flags flags.

Once the window has been embedded into the container, the container will control the window's geometry and visibility. Explicit calls to QWindow::setGeometry(), QWindow::show() or QWindow::hide() on an embedded window is not recommended.

The container takes over ownership of window. The window can be removed from the window container with a call to QWindow::setParent().

The window container is attached as a native child window to the toplevel window it is a child of. When a window container is used as a child of a QAbstractScrollArea or QMdiArea, it will create a native window for every widget in its parent chain to allow for proper stacking and clipping in this use case. Creating a native window for the window container also allows for proper stacking and clipping. This must be done before showing the window container. Applications with many native child windows may suffer from performance issues.

The window container has a number of known limitations:


Stacking order; The embedded window will stack on top of the widget hierarchy as an opaque box. The stacking order of multiple overlapping window container instances is undefined.
Rendering Integration; The window container does not interoperate with QGraphicsProxyWidget, QWidget::render() or similar functionality.
Focus Handling; It is possible to let the window container instance have any focus policy and it will delegate focus to the window via a call to QWindow::requestActivate(). However, returning to the normal focus chain from the QWindow instance will be up to the QWindow instance implementation itself. For instance, when entering a Qt Quick based window with tab focus, it is quite likely that further tab presses will only cycle inside the QML application. Also, whether QWindow::requestActivate() actually gives the window focus, is platform dependent.
Using many window container instances in a QWidget-based application can greatly hurt the overall performance of the application.
*/
func (this *QWidget) CreateWindowContainerp(window qtgui.QWindow_ITF /*777 QWindow **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21createWindowContainerEP7QWindowPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:599
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * createWindowContainer(QWindow *, QWidget *, Qt::WindowFlags)

/*
Creates a QWidget that makes it possible to embed window into a QWidget-based application.

The window container is created as a child of parent and with window flags flags.

Once the window has been embedded into the container, the container will control the window's geometry and visibility. Explicit calls to QWindow::setGeometry(), QWindow::show() or QWindow::hide() on an embedded window is not recommended.

The container takes over ownership of window. The window can be removed from the window container with a call to QWindow::setParent().

The window container is attached as a native child window to the toplevel window it is a child of. When a window container is used as a child of a QAbstractScrollArea or QMdiArea, it will create a native window for every widget in its parent chain to allow for proper stacking and clipping in this use case. Creating a native window for the window container also allows for proper stacking and clipping. This must be done before showing the window container. Applications with many native child windows may suffer from performance issues.

The window container has a number of known limitations:


Stacking order; The embedded window will stack on top of the widget hierarchy as an opaque box. The stacking order of multiple overlapping window container instances is undefined.
Rendering Integration; The window container does not interoperate with QGraphicsProxyWidget, QWidget::render() or similar functionality.
Focus Handling; It is possible to let the window container instance have any focus policy and it will delegate focus to the window via a call to QWindow::requestActivate(). However, returning to the normal focus chain from the QWindow instance will be up to the QWindow instance implementation itself. For instance, when entering a Qt Quick based window with tab focus, it is quite likely that further tab presses will only cycle inside the QML application. Also, whether QWindow::requestActivate() actually gives the window focus, is platform dependent.
Using many window container instances in a QWidget-based application can greatly hurt the overall performance of the application.
*/
func (this *QWidget) CreateWindowContainerp1(window qtgui.QWindow_ITF /*777 QWindow **/, parent QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21createWindowContainerEP7QWindowPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:604
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowTitleChanged(const QString &)

/*
This signal is emitted when the window's title has changed, with the new title as an argument.

This function was introduced in  Qt 5.2.

Note: Notifier signal for property windowTitle.
*/
func (this *QWidget) WindowTitleChanged(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18windowTitleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:605
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowIconChanged(const QIcon &)

/*
This signal is emitted when the window's icon has changed, with the new icon as an argument.

This function was introduced in  Qt 5.2.

Note: Notifier signal for property windowIcon.
*/
func (this *QWidget) WindowIconChanged(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17windowIconChangedERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:606
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowIconTextChanged(const QString &)

/*

 */
func (this *QWidget) WindowIconTextChanged(iconText string) {
	var tmpArg0 = qtcore.NewQString5(iconText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21windowIconTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:607
// index:0
// Public Visibility=Default Availability=Available
// [-2] void customContextMenuRequested(const QPoint &)

/*
This signal is emitted when the widget's contextMenuPolicy is Qt::CustomContextMenu, and the user has requested a context menu on the widget. The position pos is the position of the context menu event that the widget receives. Normally this is in widget coordinates. The exception to this rule is QAbstractScrollArea and its subclasses that map the context menu event to coordinates of the viewport().

See also mapToGlobal(), QMenu, and contextMenuPolicy.
*/
func (this *QWidget) CustomContextMenuRequested(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget26customContextMenuRequestedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:611
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().

This is the main event handler; it handles event event. You can reimplement this function in a subclass, but we recommend using one of the specialized event handlers instead.

Key press and release events are treated differently from other events. event() checks for Tab and Shift+Tab and tries to move the focus appropriately. If there is no widget to move the focus to (or the key press is not Tab or Shift+Tab), event() calls keyPressEvent().

Mouse and tablet event handling is also slightly special: only when the widget is enabled, event() will call the specialized handlers such as mousePressEvent(); otherwise it will discard the event.

This function returns true if the event was recognized, otherwise it returns false. If the recognized event was accepted (see QEvent::accepted), any further processing such as event propagation to the parent widget stops.

See also closeEvent(), focusInEvent(), focusOutEvent(), enterEvent(), keyPressEvent(), keyReleaseEvent(), leaveEvent(), mouseDoubleClickEvent(), mouseMoveEvent(), mousePressEvent(), mouseReleaseEvent(), moveEvent(), paintEvent(), resizeEvent(), QObject::event(), and QObject::timerEvent().
*/
func (this *QWidget) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:612
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive mouse press events for the widget.

If you create new widgets in the mousePressEvent() the mouseReleaseEvent() may not end up where you expect, depending on the underlying window system (or X11 window manager), the widgets' location and maybe more.

The default implementation implements the closing of popup widgets when you click outside the window. For other widget types it does nothing.

See also mouseReleaseEvent(), mouseDoubleClickEvent(), mouseMoveEvent(), event(), QMouseEvent, and Scribble Example.
*/
func (this *QWidget) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:613
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive mouse release events for the widget.

See also mousePressEvent(), mouseDoubleClickEvent(), mouseMoveEvent(), event(), QMouseEvent, and Scribble Example.
*/
func (this *QWidget) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:614
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive mouse double click events for the widget.

The default implementation calls mousePressEvent().

Note: The widget will also receive mouse press and mouse release events in addition to the double click event. It is up to the developer to ensure that the application interprets these events correctly.

See also mousePressEvent(), mouseReleaseEvent(), mouseMoveEvent(), event(), and QMouseEvent.
*/
func (this *QWidget) MouseDoubleClickEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:615
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive mouse move events for the widget.

If mouse tracking is switched off, mouse move events only occur if a mouse button is pressed while the mouse is being moved. If mouse tracking is switched on, mouse move events occur even if no mouse button is pressed.

QMouseEvent::pos() reports the position of the mouse cursor, relative to this widget. For press and release events, the position is usually the same as the position of the last mouse move event, but it might be different if the user's hand shakes. This is a feature of the underlying window system, not Qt.

If you want to show a tooltip immediately, while the mouse is moving (e.g., to get the mouse coordinates with QMouseEvent::pos() and show them as a tooltip), you must first enable mouse tracking as described above. Then, to ensure that the tooltip is updated immediately, you must call QToolTip::showText() instead of setToolTip() in your implementation of mouseMoveEvent().

See also setMouseTracking(), mousePressEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), event(), QMouseEvent, and Scribble Example.
*/
func (this *QWidget) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:619
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive key press events for the widget.

A widget must call setFocusPolicy() to accept focus initially and have focus in order to receive a key press event.

If you reimplement this handler, it is very important that you call the base class implementation if you do not act upon the key.

The default implementation closes popup widgets if the user presses the key sequence for QKeySequence::Cancel (typically the Escape key). Otherwise the event is ignored, so that the widget's parent can interpret it.

Note that QKeyEvent starts with isAccepted() == true, so you do not need to call QKeyEvent::accept() - just do not call the base class implementation if you act upon the key.

See also keyReleaseEvent(), setFocusPolicy(), focusInEvent(), focusOutEvent(), event(), QKeyEvent, and Tetrix Example.
*/
func (this *QWidget) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:620
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive key release events for the widget.

A widget must accept focus initially and have focus in order to receive a key release event.

If you reimplement this handler, it is very important that you call the base class implementation if you do not act upon the key.

The default implementation ignores the event, so that the widget's parent can interpret it.

Note that QKeyEvent starts with isAccepted() == true, so you do not need to call QKeyEvent::accept() - just do not call the base class implementation if you act upon the key.

See also keyPressEvent(), QEvent::ignore(), setFocusPolicy(), focusInEvent(), focusOutEvent(), event(), and QKeyEvent.
*/
func (this *QWidget) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:621
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
This event handler can be reimplemented in a subclass to receive keyboard focus events (focus received) for the widget. The event is passed in the event parameter

A widget normally must setFocusPolicy() to something other than Qt::NoFocus in order to receive focus events. (Note that the application programmer can call setFocus() on any widget, even those that do not normally accept focus.)

The default implementation updates the widget (except for windows that do not specify a focusPolicy()).

See also focusOutEvent(), setFocusPolicy(), keyPressEvent(), keyReleaseEvent(), event(), and QFocusEvent.
*/
func (this *QWidget) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:622
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
This event handler can be reimplemented in a subclass to receive keyboard focus events (focus lost) for the widget. The events is passed in the event parameter.

A widget normally must setFocusPolicy() to something other than Qt::NoFocus in order to receive focus events. (Note that the application programmer can call setFocus() on any widget, even those that do not normally accept focus.)

The default implementation updates the widget (except for windows that do not specify a focusPolicy()).

See also focusInEvent(), setFocusPolicy(), keyPressEvent(), keyReleaseEvent(), event(), and QFocusEvent.
*/
func (this *QWidget) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:623
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void enterEvent(QEvent *)

/*
This event handler can be reimplemented in a subclass to receive widget enter events which are passed in the event parameter.

An event is sent to the widget when the mouse cursor enters the widget.

See also leaveEvent(), mouseMoveEvent(), and event().
*/
func (this *QWidget) EnterEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10enterEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:624
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)

/*
This event handler can be reimplemented in a subclass to receive widget leave events which are passed in the event parameter.

A leave event is sent to the widget when the mouse cursor leaves the widget.

See also enterEvent(), mouseMoveEvent(), and event().
*/
func (this *QWidget) LeaveEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:625
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
This event handler can be reimplemented in a subclass to receive paint events passed in event.

A paint event is a request to repaint all or part of a widget. It can happen for one of the following reasons:


repaint() or update() was invoked,
the widget was obscured and has now been uncovered, or
many other reasons.


Many widgets can simply repaint their entire surface when asked to, but some slow widgets need to optimize by painting only the requested region: QPaintEvent::region(). This speed optimization does not change the result, as painting is clipped to that region during event processing. QListView and QTableView do this, for example.

Qt also tries to speed up painting by merging multiple paint events into one. When update() is called several times or the window system sends several paint events, Qt merges these events into one event with a larger region (see QRegion::united()). The repaint() function does not permit this optimization, so we suggest using update() whenever possible.

When the paint event occurs, the update region has normally been erased, so you are painting on the widget's background.

The background can be set using setBackgroundRole() and setPalette().

Since Qt 4.0, QWidget automatically double-buffers its painting, so there is no need to write double-buffering code in paintEvent() to avoid flicker.

Note: Generally, you should refrain from calling update() or repaint() inside a paintEvent(). For example, calling update() or repaint() on children inside a paintEvent() results in undefined behavior; the child may or may not get a paint event.

Warning: If you are using a custom paint engine without Qt's backingstore, Qt::WA_PaintOnScreen must be set. Otherwise, QWidget::paintEngine() will never be called; the backingstore will be used instead.

See also event(), repaint(), update(), QPainter, QPixmap, QPaintEvent, and Analog Clock Example.
*/
func (this *QWidget) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:626
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)

/*
This event handler can be reimplemented in a subclass to receive widget move events which are passed in the event parameter. When the widget receives this event, it is already at the new position.

The old position is accessible through QMoveEvent::oldPos().

See also resizeEvent(), event(), move(), and QMoveEvent.
*/
func (this *QWidget) MoveEvent(event qtgui.QMoveEvent_ITF /*777 QMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMoveEvent_PTR() != nil {
		convArg0 = event.QMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:627
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
This event handler can be reimplemented in a subclass to receive widget resize events which are passed in the event parameter. When resizeEvent() is called, the widget already has its new geometry. The old size is accessible through QResizeEvent::oldSize().

The widget will be erased and receive a paint event immediately after processing the resize event. No drawing need be (or should be) done inside this handler.

See also moveEvent(), event(), resize(), QResizeEvent, paintEvent(), and Scribble Example.
*/
func (this *QWidget) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:628
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)

/*
This event handler is called with the given event when Qt receives a window close request for a top-level widget from the window system.

By default, the event is accepted and the widget is closed. You can reimplement this function to change the way the widget responds to window close requests. For example, you can prevent the window from closing by calling ignore() on all events.

Main window applications typically use reimplementations of this function to check whether the user's work has been saved and ask for permission before closing. For example, the Application Example uses a helper function to determine whether or not to close the window:


  void MainWindow::closeEvent(QCloseEvent *event)
  {
      if (maybeSave()) {
          writeSettings();
          event->accept();
      } else {
          event->ignore();
      }
  }



See also event(), hide(), close(), QCloseEvent, and Application Example.
*/
func (this *QWidget) CloseEvent(event qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QCloseEvent_PTR() != nil {
		convArg0 = event.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:630
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive widget context menu events.

The handler is called when the widget's contextMenuPolicy is Qt::DefaultContextMenu.

The default implementation ignores the context event. See the QContextMenuEvent documentation for more details.

See also event(), QContextMenuEvent, and customContextMenuRequested().
*/
func (this *QWidget) ContextMenuEvent(event qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QContextMenuEvent_PTR() != nil {
		convArg0 = event.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:636
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)

/*
This event handler is called with the given event whenever the widget's actions are changed.

See also addAction(), insertAction(), removeAction(), actions(), and QActionEvent.
*/
func (this *QWidget) ActionEvent(event qtgui.QActionEvent_ITF /*777 QActionEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QActionEvent_PTR() != nil {
		convArg0 = event.QActionEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:646
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
This event handler can be reimplemented in a subclass to receive widget show events which are passed in the event parameter.

Non-spontaneous show events are sent to widgets immediately before they are shown. The spontaneous show events of windows are delivered afterwards.

Note: A widget receives spontaneous show and hide events when its mapping status is changed by the window system, e.g. a spontaneous hide event when the user minimizes the window, and a spontaneous show event when the window is restored again. After receiving a spontaneous hide event, a widget is still considered visible in the sense of isVisible().

See also visible, event(), and QShowEvent.
*/
func (this *QWidget) ShowEvent(event qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QShowEvent_PTR() != nil {
		convArg0 = event.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:647
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
This event handler can be reimplemented in a subclass to receive widget hide events. The event is passed in the event parameter.

Hide events are sent to widgets immediately after they have been hidden.

Note: A widget receives spontaneous show and hide events when its mapping status is changed by the window system, e.g. a spontaneous hide event when the user minimizes the window, and a spontaneous show event when the window is restored again. After receiving a spontaneous hide event, a widget is still considered visible in the sense of isVisible().

See also visible, event(), and QHideEvent.
*/
func (this *QWidget) HideEvent(event qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHideEvent_PTR() != nil {
		convArg0 = event.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:648
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool nativeEvent(const QByteArray &, void *, long *)

/*
This special event handler can be reimplemented in a subclass to receive native platform events identified by eventType which are passed in the message parameter.

In your reimplementation of this function, if you want to stop the event being handled by Qt, return true and set result. The result parameter has meaning only on Windows. If you return false, this native event is passed back to Qt, which translates the event into a Qt event and sends it to the widget.

Note: Events are only delivered to this event handler if the widget has a native window handle.

Note: This function superseedes the event filter functions x11Event(), winEvent() and macEvent() of Qt 4.


 PlatformEvent Type IdentifierMessage TypeResult Type
Windows"windows_generic_MSG"MSG *LRESULT
macOS"NSEvent"NSEvent *
XCB"xcb_generic_event_t"xcb_generic_event_t *


See also QAbstractNativeEventFilter.
*/
func (this *QWidget) NativeEvent(eventType qtcore.QByteArray_ITF, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 unsafe.Pointer
	if eventType != nil && eventType.QByteArray_PTR() != nil {
		convArg0 = eventType.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11nativeEventERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:651
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
This event handler can be reimplemented to handle state changes.

The state being changed in this event can be retrieved through the event supplied.

Change events include: QEvent::ToolBarChange, QEvent::ActivationChange, QEvent::EnabledChange, QEvent::FontChange, QEvent::StyleChange, QEvent::PaletteChange, QEvent::WindowTitleChange, QEvent::IconTextChange, QEvent::ModifiedChange, QEvent::MouseTrackingChange, QEvent::ParentChange, QEvent::WindowStateChange, QEvent::LanguageChange, QEvent::LocaleChange, QEvent::LayoutDirectionChange, QEvent::ReadOnlyChange.
*/
func (this *QWidget) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:653
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*
Reimplemented from QPaintDevice::metric().

Internal implementation of the virtual QPaintDevice::metric() function.

m is the metric to get.
*/
func (this *QWidget) Metric(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:654
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initPainter(QPainter *) const

/*
Initializes the painter pen, background and font to the same as the given widget's. This function is called automatically when the painter is opened on a QWidget.
*/
func (this *QWidget) InitPainter(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11initPainterEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:655
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(QPoint *) const

/*

 */
func (this *QWidget) Redirected(offset qtcore.QPoint_ITF /*777 QPoint **/) *qtgui.QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg0 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10redirectedEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:656
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPainter * sharedPainter() const

/*

 */
func (this *QWidget) SharedPainter() *qtgui.QPainter /*777 QPainter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13sharedPainterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:658
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive Input Method composition events. This handler is called when the state of the input method changes.

Note that when creating custom text editing widgets, the Qt::WA_InputMethodEnabled window attribute must be set explicitly (using the setAttribute() function) in order to receive input method events.

The default implementation calls event->ignore(), which rejects the Input Method event. See the QInputMethodEvent documentation for more details.

See also event() and QInputMethodEvent.
*/
func (this *QWidget) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:660
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
This method is only relevant for input widgets. It is used by the input method to query a set of properties of the widget to be able to support complex input method operations as support for surrounding text and reconversions.

query specifies which property is queried.

See also inputMethodEvent(), QInputMethodEvent, QInputMethodQueryEvent, and inputMethodHints.
*/
func (this *QWidget) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:662
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::InputMethodHints inputMethodHints() const

/*

 */
func (this *QWidget) InputMethodHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16inputMethodHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:663
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMethodHints(Qt::InputMethodHints)

/*

 */
func (this *QWidget) SetInputMethodHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget19setInputMethodHintsE6QFlagsIN2Qt15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:666
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateMicroFocus()

/*
Updates the widget's micro focus.
*/
func (this *QWidget) UpdateMicroFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16updateMicroFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:669
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(WId, bool, bool)

/*
Creates a new widget window.

The parameter window is ignored in Qt 5. Please use QWindow::fromWinId() to create a QWindow wrapping a foreign window and pass it to QWidget::createWindowContainer() instead.

Initializes the window (sets the geometry etc.) if initializeWindow is true. If initializeWindow is false, no initialization is performed. This parameter only makes sense if window is a valid window.

Destroys the old window if destroyOldWindow is true. If destroyOldWindow is false, you are responsible for destroying the window yourself (using platform native code).

The QWidget constructor calls create(0,true,true) to create a window for this widget.

See also createWindowContainer() and QWindow::fromWinId().
*/
func (this *QWidget) Create(arg0 uint64, initializeWindow bool, destroyOldWindow bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6createEybb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, initializeWindow, destroyOldWindow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:669
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(WId, bool, bool)

/*
Creates a new widget window.

The parameter window is ignored in Qt 5. Please use QWindow::fromWinId() to create a QWindow wrapping a foreign window and pass it to QWidget::createWindowContainer() instead.

Initializes the window (sets the geometry etc.) if initializeWindow is true. If initializeWindow is false, no initialization is performed. This parameter only makes sense if window is a valid window.

Destroys the old window if destroyOldWindow is true. If destroyOldWindow is false, you are responsible for destroying the window yourself (using platform native code).

The QWidget constructor calls create(0,true,true) to create a window for this widget.

See also createWindowContainer() and QWindow::fromWinId().
*/
func (this *QWidget) Createp() {
	// arg: 0, WId=Typedef, WId=Typedef, ::quintptr, Elaborated
	var arg0 unsafe.Pointer
	// arg: 1, bool=Bool, =Invalid, , Invalid
	initializeWindow := true
	// arg: 2, bool=Bool, =Invalid, , Invalid
	destroyOldWindow := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6createEybb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, initializeWindow, destroyOldWindow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:669
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(WId, bool, bool)

/*
Creates a new widget window.

The parameter window is ignored in Qt 5. Please use QWindow::fromWinId() to create a QWindow wrapping a foreign window and pass it to QWidget::createWindowContainer() instead.

Initializes the window (sets the geometry etc.) if initializeWindow is true. If initializeWindow is false, no initialization is performed. This parameter only makes sense if window is a valid window.

Destroys the old window if destroyOldWindow is true. If destroyOldWindow is false, you are responsible for destroying the window yourself (using platform native code).

The QWidget constructor calls create(0,true,true) to create a window for this widget.

See also createWindowContainer() and QWindow::fromWinId().
*/
func (this *QWidget) Createp1(arg0 uint64) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	initializeWindow := true
	// arg: 2, bool=Bool, =Invalid, , Invalid
	destroyOldWindow := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6createEybb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, initializeWindow, destroyOldWindow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:669
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(WId, bool, bool)

/*
Creates a new widget window.

The parameter window is ignored in Qt 5. Please use QWindow::fromWinId() to create a QWindow wrapping a foreign window and pass it to QWidget::createWindowContainer() instead.

Initializes the window (sets the geometry etc.) if initializeWindow is true. If initializeWindow is false, no initialization is performed. This parameter only makes sense if window is a valid window.

Destroys the old window if destroyOldWindow is true. If destroyOldWindow is false, you are responsible for destroying the window yourself (using platform native code).

The QWidget constructor calls create(0,true,true) to create a window for this widget.

See also createWindowContainer() and QWindow::fromWinId().
*/
func (this *QWidget) Createp2(arg0 uint64, initializeWindow bool) {
	// arg: 2, bool=Bool, =Invalid, , Invalid
	destroyOldWindow := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6createEybb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, initializeWindow, destroyOldWindow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:671
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void destroy(bool, bool)

/*
Frees up window system resources. Destroys the widget window if destroyWindow is true.

destroy() calls itself recursively for all the child widgets, passing destroySubWindows for the destroyWindow parameter. To have more control over destruction of subwidgets, destroy subwidgets selectively first.

This function is usually called from the QWidget destructor.
*/
func (this *QWidget) Destroy(destroyWindow bool, destroySubWindows bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7destroyEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destroyWindow, destroySubWindows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:671
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void destroy(bool, bool)

/*
Frees up window system resources. Destroys the widget window if destroyWindow is true.

destroy() calls itself recursively for all the child widgets, passing destroySubWindows for the destroyWindow parameter. To have more control over destruction of subwidgets, destroy subwidgets selectively first.

This function is usually called from the QWidget destructor.
*/
func (this *QWidget) Destroyp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	destroyWindow := true
	// arg: 1, bool=Bool, =Invalid, , Invalid
	destroySubWindows := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7destroyEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destroyWindow, destroySubWindows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:671
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void destroy(bool, bool)

/*
Frees up window system resources. Destroys the widget window if destroyWindow is true.

destroy() calls itself recursively for all the child widgets, passing destroySubWindows for the destroyWindow parameter. To have more control over destruction of subwidgets, destroy subwidgets selectively first.

This function is usually called from the QWidget destructor.
*/
func (this *QWidget) Destroyp1(destroyWindow bool) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	destroySubWindows := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7destroyEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destroyWindow, destroySubWindows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:675
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Finds a new widget to give the keyboard focus to, as appropriate for Tab and Shift+Tab, and returns true if it can find a new widget, or false if it can't.

If next is true, this function searches forward, if next is false, it searches backward.

Sometimes, you will want to reimplement this function. For example, a web browser might reimplement it to move its "current active link" forward or backward, and call focusNextPrevChild() only when it reaches the last or first link on the "page".

Child widgets call focusNextPrevChild() on their parent widgets, but only the window that contains the child widgets decides where to redirect focus. By reimplementing this function for an object, you thus gain control of focus traversal for all child widgets.

See also focusNextChild() and focusPreviousChild().
*/
func (this *QWidget) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:676
// index:0
// Protected inline Visibility=Default Availability=Available
// [1] bool focusNextChild()

/*
Finds a new widget to give the keyboard focus to, as appropriate for Tab, and returns true if it can find a new widget, or false if it can't.

See also focusPreviousChild().
*/
func (this *QWidget) FocusNextChild() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14focusNextChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:677
// index:0
// Protected inline Visibility=Default Availability=Available
// [1] bool focusPreviousChild()

/*
Finds a new widget to give the keyboard focus to, as appropriate for Shift+Tab, and returns true if it can find a new widget, or false if it can't.

See also focusNextChild().
*/
func (this *QWidget) FocusPreviousChild() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18focusPreviousChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QWidget__RenderFlag = int

//
const QWidget__DrawWindowBackground QWidget__RenderFlag = 1

//
const QWidget__DrawChildren QWidget__RenderFlag = 2

//
const QWidget__IgnoreMask QWidget__RenderFlag = 4

func (this *QWidget) RenderFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWidget_RenderFlagItemName(val int) string {
	var nilthis *QWidget
	return nilthis.RenderFlagItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10983() {
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
