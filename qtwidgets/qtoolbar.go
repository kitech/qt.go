package qtwidgets

// /usr/include/qt/QtWidgets/qtoolbar.h
// #include <qtoolbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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

// void actionEvent(QActionEvent *)
func (this *QToolBar) InheritActionEvent(f func(event *qtgui.QActionEvent /*777 QActionEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void changeEvent(QEvent *)
func (this *QToolBar) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QToolBar) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// bool event(QEvent *)
func (this *QToolBar) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(QStyleOptionToolBar *)
func (this *QToolBar) InheritInitStyleOption(f func(option *QStyleOptionToolBar /*777 QStyleOptionToolBar **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QToolBar struct {
	*QWidget
}
type QToolBar_ITF interface {
	QWidget_ITF
	QToolBar_PTR() *QToolBar
}

func (ptr *QToolBar) QToolBar_PTR() *QToolBar { return ptr }

func (this *QToolBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QToolBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQToolBarFromPointer(cthis unsafe.Pointer) *QToolBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QToolBar{bcthis0}
}
func (*QToolBar) NewFromPointer(cthis unsafe.Pointer) *QToolBar {
	return NewQToolBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QToolBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(const QString &, QWidget *)

/*
Constructs a QToolBar with the given parent.

The given window title identifies the toolbar and is shown in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQToolBar(title string, parent QWidget_ITF /*777 QWidget **/) *QToolBar {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(const QString &, QWidget *)

/*
Constructs a QToolBar with the given parent.

The given window title identifies the toolbar and is shown in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQToolBar__(title string) *QToolBar {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:80
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(QWidget *)

/*
Constructs a QToolBar with the given parent.

The given window title identifies the toolbar and is shown in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQToolBar_1(parent QWidget_ITF /*777 QWidget **/) *QToolBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:80
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QToolBar(QWidget *)

/*
Constructs a QToolBar with the given parent.

The given window title identifies the toolbar and is shown in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQToolBar_1_() *QToolBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QToolBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qtoolbar.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QToolBar()

/*

 */
func DeleteQToolBar(this *QToolBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(bool)

/*

 */
func (this *QToolBar) SetMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar10setMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable() const

/*

 */
func (this *QToolBar) IsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar9isMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAllowedAreas(Qt::ToolBarAreas)

/*

 */
func (this *QToolBar) SetAllowedAreas(areas int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15setAllowedAreasE6QFlagsIN2Qt11ToolBarAreaEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), areas)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolBarAreas allowedAreas() const

/*

 */
func (this *QToolBar) AllowedAreas() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar12allowedAreasEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAreaAllowed(Qt::ToolBarArea) const

/*
Returns true if this toolbar is dockable in the given area; otherwise returns false.
*/
func (this *QToolBar) IsAreaAllowed(area int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar13isAreaAllowedEN2Qt11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*

 */
func (this *QToolBar) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QToolBar) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all actions from the toolbar.

See also removeAction().
*/
func (this *QToolBar) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)

/*
This is an overloaded function.

Creates a new action with the given text. This action is added to the end of the toolbar.
*/
func (this *QToolBar) AddAction(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:99
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &)

/*
This is an overloaded function.

Creates a new action with the given text. This action is added to the end of the toolbar.
*/
func (this *QToolBar) AddAction_1(icon qtgui.QIcon_ITF, text string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:100
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &, const QObject *, const char *)

/*
This is an overloaded function.

Creates a new action with the given text. This action is added to the end of the toolbar.
*/
func (this *QToolBar) AddAction_2(text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK7QStringPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:101
// index:3
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &, const QObject *, const char *)

/*
This is an overloaded function.

Creates a new action with the given text. This action is added to the end of the toolbar.
*/
func (this *QToolBar) AddAction_3(icon qtgui.QIcon_ITF, text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addActionERK5QIconRK7QStringPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSeparator()

/*
Adds a separator to the end of the toolbar.

See also insertSeparator().
*/
func (this *QToolBar) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar12addSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSeparator(QAction *)

/*
Inserts a separator into the toolbar in front of the toolbar item associated with the before action.

See also addSeparator().
*/
func (this *QToolBar) InsertSeparator(before QAction_ITF /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15insertSeparatorEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addWidget(QWidget *)

/*
Adds the given widget to the toolbar as the toolbar's last item.

The toolbar takes ownership of widget.

If you add a QToolButton with this method, the toolbar's Qt::ToolButtonStyle will not be respected.

Note: You should use QAction::setVisible() to change the visibility of the widget. Using QWidget::setVisible(), QWidget::show() and QWidget::hide() does not work.

See also insertWidget().
*/
func (this *QToolBar) AddWidget(widget QWidget_ITF /*777 QWidget **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertWidget(QAction *, QWidget *)

/*
Inserts the given widget in front of the toolbar item associated with the before action.

Note: You should use QAction::setVisible() to change the visibility of the widget. Using QWidget::setVisible(), QWidget::show() and QWidget::hide() does not work.

See also addWidget().
*/
func (this *QToolBar) InsertWidget(before QAction_ITF /*777 QAction **/, widget QWidget_ITF /*777 QWidget **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar12insertWidgetEP7QActionP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:161
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect actionGeometry(QAction *) const

/*

 */
func (this *QToolBar) ActionGeometry(action QAction_ITF /*777 QAction **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar14actionGeometryEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * actionAt(const QPoint &) const

/*
Returns the action at point p. This function returns zero if no action was found.

See also QWidget::childAt().
*/
func (this *QToolBar) ActionAt(p qtcore.QPoint_ITF) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar8actionAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:163
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QAction * actionAt(int, int) const

/*
Returns the action at point p. This function returns zero if no action was found.

See also QWidget::childAt().
*/
func (this *QToolBar) ActionAt_1(x int, y int) *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar8actionAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * toggleViewAction() const

/*
Returns a checkable action that can be used to show or hide this toolbar.

The action's text is set to the toolbar's window title.

See also QAction::text and QWidget::windowTitle.
*/
func (this *QToolBar) ToggleViewAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar16toggleViewActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QToolBar) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtoolbar.h:168
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolButtonStyle toolButtonStyle() const

/*

 */
func (this *QToolBar) ToolButtonStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar15toolButtonStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widgetForAction(QAction *) const

/*
Returns the widget associated with the specified action.

This function was introduced in  Qt 4.2.

See also addWidget().
*/
func (this *QToolBar) WidgetForAction(action QAction_ITF /*777 QAction **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar15widgetForActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtoolbar.h:172
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFloatable() const

/*

 */
func (this *QToolBar) IsFloatable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar11isFloatableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFloatable(bool)

/*

 */
func (this *QToolBar) SetFloatable(floatable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar12setFloatableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), floatable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFloating() const

/*

 */
func (this *QToolBar) IsFloating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar10isFloatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QToolBar) SetIconSize(iconSize qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if iconSize != nil && iconSize.QSize_PTR() != nil {
		convArg0 = iconSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolButtonStyle(Qt::ToolButtonStyle)

/*

 */
func (this *QToolBar) SetToolButtonStyle(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar18setToolButtonStyleEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actionTriggered(QAction *)

/*
This signal is emitted when an action in this toolbar is triggered. This happens when the action's tool button is pressed, or when the action is triggered in some other way outside the toolbar. The parameter holds the triggered action.
*/
func (this *QToolBar) ActionTriggered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15actionTriggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void movableChanged(bool)

/*
This signal is emitted when the toolbar becomes movable or fixed. If the toolbar can be moved, movable is true; otherwise it is false.

Note: Notifier signal for property movable.

See also movable.
*/
func (this *QToolBar) MovableChanged(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar14movableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void allowedAreasChanged(Qt::ToolBarAreas)

/*
This signal is emitted when the collection of allowed areas for the toolbar is changed. The new areas in which the toolbar can be positioned are specified by allowedAreas.

Note: Notifier signal for property allowedAreas.

See also allowedAreas.
*/
func (this *QToolBar) AllowedAreasChanged(allowedAreas int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar19allowedAreasChangedE6QFlagsIN2Qt11ToolBarAreaEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), allowedAreas)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void orientationChanged(Qt::Orientation)

/*
This signal is emitted when the orientation of the toolbar changes. The orientation parameter holds the toolbar's new orientation.

Note: Notifier signal for property orientation.

See also orientation.
*/
func (this *QToolBar) OrientationChanged(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar18orientationChangedEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconSizeChanged(const QSize &)

/*
This signal is emitted when the icon size is changed. The iconSize parameter holds the toolbar's new icon size.

Note: Notifier signal for property iconSize.

See also iconSize and QMainWindow::iconSize.
*/
func (this *QToolBar) IconSizeChanged(iconSize qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if iconSize != nil && iconSize.QSize_PTR() != nil {
		convArg0 = iconSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15iconSizeChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toolButtonStyleChanged(Qt::ToolButtonStyle)

/*
This signal is emitted when the tool button style is changed. The toolButtonStyle parameter holds the toolbar's new tool button style.

Note: Notifier signal for property toolButtonStyle.

See also toolButtonStyle and QMainWindow::toolButtonStyle.
*/
func (this *QToolBar) ToolButtonStyleChanged(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topLevelChanged(bool)

/*
This signal is emitted when the floating property changes. The topLevel parameter is true if the toolbar is now floating; otherwise it is false.

This function was introduced in  Qt 4.6.

See also isWindow().
*/
func (this *QToolBar) TopLevelChanged(topLevel bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar15topLevelChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), topLevel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibilityChanged(bool)

/*
This signal is emitted when the toolbar becomes visible (or invisible). This happens when the widget is hidden or shown.

This function was introduced in  Qt 4.7.
*/
func (this *QToolBar) VisibilityChanged(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar17visibilityChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)

/*
Reimplemented from QWidget::actionEvent().
*/
func (this *QToolBar) ActionEvent(event qtgui.QActionEvent_ITF /*777 QActionEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QActionEvent_PTR() != nil {
		convArg0 = event.QActionEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QToolBar) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:193
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QToolBar) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbar.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QToolBar) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QToolBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtoolbar.h:195
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionToolBar *) const

/*

 */
func (this *QToolBar) InitStyleOption(option QStyleOptionToolBar_ITF /*777 QStyleOptionToolBar **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionToolBar_PTR() != nil {
		convArg0 = option.QStyleOptionToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
