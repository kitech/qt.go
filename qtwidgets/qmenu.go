package qtwidgets

// /usr/include/qt/QtWidgets/qmenu.h
// #include <qmenu.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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

// int columnCount()
func (this *QMenu) InheritColumnCount(f func() int) {
	qtrt.SetAllInheritCallback(this, "columnCount", f)
}

// void changeEvent(QEvent *)
func (this *QMenu) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QMenu) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QMenu) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QMenu) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QMenu) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QMenu) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void enterEvent(QEvent *)
func (this *QMenu) InheritEnterEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterEvent", f)
}

// void leaveEvent(QEvent *)
func (this *QMenu) InheritLeaveEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QMenu) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QMenu) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void actionEvent(QActionEvent *)
func (this *QMenu) InheritActionEvent(f func(arg0 *qtgui.QActionEvent /*777 QActionEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QMenu) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// bool event(QEvent *)
func (this *QMenu) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool focusNextPrevChild(bool)
func (this *QMenu) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void initStyleOption(QStyleOptionMenuItem *, const QAction *)
func (this *QMenu) InheritInitStyleOption(f func(option *QStyleOptionMenuItem /*777 QStyleOptionMenuItem **/, action *QAction /*777 const QAction **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QMenu struct {
	*QWidget
}
type QMenu_ITF interface {
	QWidget_ITF
	QMenu_PTR() *QMenu
}

func (ptr *QMenu) QMenu_PTR() *QMenu { return ptr }

func (this *QMenu) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMenu) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMenuFromPointer(cthis unsafe.Pointer) *QMenu {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMenu{bcthis0}
}
func (*QMenu) NewFromPointer(cthis unsafe.Pointer) *QMenu {
	return NewQMenuFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmenu.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMenu) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenu(QWidget *)

/*
Constructs a menu with parent parent.

Although a popup menu is always a top-level widget, if a parent is passed the popup menu will be deleted when that parent is destroyed (as with any other QObject).
*/
func (*QMenu) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QMenu {
	return NewQMenu(parent)
}
func NewQMenu(parent QWidget_ITF /*777 QWidget **/) *QMenu {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenuC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMenu(QWidget *)

/*
Constructs a menu with parent parent.

Although a popup menu is always a top-level widget, if a parent is passed the popup menu will be deleted when that parent is destroyed (as with any other QObject).
*/
func (*QMenu) NewForInherit__() *QMenu {
	return NewQMenu__()
}
func NewQMenu__() *QMenu {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenuC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMenu(const QString &, QWidget *)

/*
Constructs a menu with parent parent.

Although a popup menu is always a top-level widget, if a parent is passed the popup menu will be deleted when that parent is destroyed (as with any other QObject).
*/
func (*QMenu) NewForInherit_1(title string, parent QWidget_ITF /*777 QWidget **/) *QMenu {
	return NewQMenu_1(title, parent)
}
func NewQMenu_1(title string, parent QWidget_ITF /*777 QWidget **/) *QMenu {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenuC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMenu(const QString &, QWidget *)

/*
Constructs a menu with parent parent.

Although a popup menu is always a top-level widget, if a parent is passed the popup menu will be deleted when that parent is destroyed (as with any other QObject).
*/
func (*QMenu) NewForInherit_1_(title string) *QMenu {
	return NewQMenu_1_(title)
}
func NewQMenu_1_(title string) *QMenu {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenuC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMenuFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMenu")
	return gothis
}

// /usr/include/qt/QtWidgets/qmenu.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMenu()

/*

 */
func DeleteQMenu(this *QMenu) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenuD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmenu.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)

/*
This is an overloaded function.

This convenience function creates a new action with text. The function adds the newly created action to the menu's list of actions, and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddAction(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:80
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &)

/*
This is an overloaded function.

This convenience function creates a new action with text. The function adds the newly created action to the menu's list of actions, and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddAction_1(icon qtgui.QIcon_ITF, text string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:81
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &, const QObject *, const char *, const QKeySequence &)

/*
This is an overloaded function.

This convenience function creates a new action with text. The function adds the newly created action to the menu's list of actions, and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddAction_2(text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string, shortcut qtgui.QKeySequence_ITF) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	var convArg3 unsafe.Pointer
	if shortcut != nil && shortcut.QKeySequence_PTR() != nil {
		convArg3 = shortcut.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:81
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &, const QObject *, const char *, const QKeySequence &)

/*
This is an overloaded function.

This convenience function creates a new action with text. The function adds the newly created action to the menu's list of actions, and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddAction_2_(text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, const QKeySequence &=LValueReference, QKeySequence=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:82
// index:3
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &, const QObject *, const char *, const QKeySequence &)

/*
This is an overloaded function.

This convenience function creates a new action with text. The function adds the newly created action to the menu's list of actions, and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddAction_3(icon qtgui.QIcon_ITF, text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string, shortcut qtgui.QKeySequence_ITF) *QAction /*777 QAction **/ {
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
	var convArg4 unsafe.Pointer
	if shortcut != nil && shortcut.QKeySequence_PTR() != nil {
		convArg4 = shortcut.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:82
// index:3
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &, const QObject *, const char *, const QKeySequence &)

/*
This is an overloaded function.

This convenience function creates a new action with text. The function adds the newly created action to the menu's list of actions, and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddAction_3_(icon qtgui.QIcon_ITF, text string, receiver qtcore.QObject_ITF /*777 const QObject **/, member string) *QAction /*777 QAction **/ {
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
	// arg: 4, const QKeySequence &=LValueReference, QKeySequence=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addMenu(QMenu *)

/*
This convenience function adds menu as a submenu to this menu. It returns menu's menuAction(). This menu does not take ownership of menu.

See also QWidget::addAction() and QMenu::menuAction().
*/
func (this *QMenu) AddMenu(menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu7addMenuEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:157
// index:1
// Public Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QString &)

/*
This convenience function adds menu as a submenu to this menu. It returns menu's menuAction(). This menu does not take ownership of menu.

See also QWidget::addAction() and QMenu::menuAction().
*/
func (this *QMenu) AddMenu_1(title string) *QMenu /*777 QMenu **/ {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu7addMenuERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:158
// index:2
// Public Visibility=Default Availability=Available
// [8] QMenu * addMenu(const QIcon &, const QString &)

/*
This convenience function adds menu as a submenu to this menu. It returns menu's menuAction(). This menu does not take ownership of menu.

See also QWidget::addAction() and QMenu::menuAction().
*/
func (this *QMenu) AddMenu_2(icon qtgui.QIcon_ITF, title string) *QMenu /*777 QMenu **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(title)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu7addMenuERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSeparator()

/*
This convenience function creates a new separator action, i.e. an action with QAction::isSeparator() returning true, and adds the new action to this menu's list of actions. It returns the newly created action.

QMenu takes ownership of the returned QAction.

See also QWidget::addAction().
*/
func (this *QMenu) AddSeparator() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu12addSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addSection(const QString &)

/*
This convenience function creates a new section action, i.e. an action with QAction::isSeparator() returning true but also having text hint, and adds the new action to this menu's list of actions. It returns the newly created action.

The rendering of the hint is style and platform dependent. Widget styles can use the text information in the rendering for sections, or can choose to ignore it and render sections like simple separators.

QMenu takes ownership of the returned QAction.

This function was introduced in  Qt 5.1.

See also QWidget::addAction().
*/
func (this *QMenu) AddSection(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10addSectionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:163
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addSection(const QIcon &, const QString &)

/*
This convenience function creates a new section action, i.e. an action with QAction::isSeparator() returning true but also having text hint, and adds the new action to this menu's list of actions. It returns the newly created action.

The rendering of the hint is style and platform dependent. Widget styles can use the text information in the rendering for sections, or can choose to ignore it and render sections like simple separators.

QMenu takes ownership of the returned QAction.

This function was introduced in  Qt 5.1.

See also QWidget::addAction().
*/
func (this *QMenu) AddSection_1(icon qtgui.QIcon_ITF, text string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10addSectionERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertMenu(QAction *, QMenu *)

/*
This convenience function inserts menu before action before and returns the menus menuAction().

See also QWidget::insertAction() and addMenu().
*/
func (this *QMenu) InsertMenu(before QAction_ITF /*777 QAction **/, menu QMenu_ITF /*777 QMenu **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg1 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10insertMenuEP7QActionPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSeparator(QAction *)

/*
This convenience function creates a new separator action, i.e. an action with QAction::isSeparator() returning true. The function inserts the newly created action into this menu's list of actions before action before and returns it.

QMenu takes ownership of the returned QAction.

See also QWidget::insertAction() and addSeparator().
*/
func (this *QMenu) InsertSeparator(before QAction_ITF /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu15insertSeparatorEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * insertSection(QAction *, const QString &)

/*
This convenience function creates a new title action, i.e. an action with QAction::isSeparator() returning true but also having text hint. The function inserts the newly created action into this menu's list of actions before action before and returns it.

The rendering of the hint is style and platform dependent. Widget styles can use the text information in the rendering for sections, or can choose to ignore it and render sections like simple separators.

QMenu takes ownership of the returned QAction.

This function was introduced in  Qt 5.1.

See also QWidget::insertAction() and addSection().
*/
func (this *QMenu) InsertSection(before QAction_ITF /*777 QAction **/, text string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:168
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * insertSection(QAction *, const QIcon &, const QString &)

/*
This convenience function creates a new title action, i.e. an action with QAction::isSeparator() returning true but also having text hint. The function inserts the newly created action into this menu's list of actions before action before and returns it.

The rendering of the hint is style and platform dependent. Widget styles can use the text information in the rendering for sections, or can choose to ignore it and render sections like simple separators.

QMenu takes ownership of the returned QAction.

This function was introduced in  Qt 5.1.

See also QWidget::insertAction() and addSection().
*/
func (this *QMenu) InsertSection_1(before QAction_ITF /*777 QAction **/, icon qtgui.QIcon_ITF, text string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if before != nil && before.QAction_PTR() != nil {
		convArg0 = before.QAction_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg1 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if there are no visible actions inserted into the menu, false otherwise.

This function was introduced in  Qt 4.2.

See also QWidget::actions().
*/
func (this *QMenu) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all the menu's actions. Actions owned by the menu and not shown in any other widget are deleted.

See also removeAction().
*/
func (this *QMenu) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTearOffEnabled(bool)

/*

 */
func (this *QMenu) SetTearOffEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu17setTearOffEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTearOffEnabled() const

/*

 */
func (this *QMenu) IsTearOffEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu16isTearOffEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTearOffMenuVisible() const

/*
When a menu is torn off a second menu is shown to display the menu contents in a new window. When the menu is in this mode and the menu is visible returns true; otherwise false.

See also showTearOffMenu(), hideTearOffMenu(), and isTearOffEnabled().
*/
func (this *QMenu) IsTearOffMenuVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu20isTearOffMenuVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showTearOffMenu()

/*
This function will forcibly show the torn off menu making it appear on the user's desktop at the specified global position pos.

This function was introduced in  Qt 5.7.

See also hideTearOffMenu(), isTearOffMenuVisible(), and isTearOffEnabled().
*/
func (this *QMenu) ShowTearOffMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:178
// index:1
// Public Visibility=Default Availability=Available
// [-2] void showTearOffMenu(const QPoint &)

/*
This function will forcibly show the torn off menu making it appear on the user's desktop at the specified global position pos.

This function was introduced in  Qt 5.7.

See also hideTearOffMenu(), isTearOffMenuVisible(), and isTearOffEnabled().
*/
func (this *QMenu) ShowTearOffMenu_1(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu15showTearOffMenuERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideTearOffMenu()

/*
This function will forcibly hide the torn off menu making it disappear from the user's desktop.

See also showTearOffMenu(), isTearOffMenuVisible(), and isTearOffEnabled().
*/
func (this *QMenu) HideTearOffMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu15hideTearOffMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultAction(QAction *)

/*
This sets the default action to act. The default action may have a visual cue, depending on the current QStyle. A default action usually indicates what will happen by default when a drop occurs.

See also defaultAction().
*/
func (this *QMenu) SetDefaultAction(arg0 QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAction_PTR() != nil {
		convArg0 = arg0.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu16setDefaultActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * defaultAction() const

/*
Returns the current default action.

See also setDefaultAction().
*/
func (this *QMenu) DefaultAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu13defaultActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveAction(QAction *)

/*
Sets the currently highlighted action to act.

See also activeAction().
*/
func (this *QMenu) SetActiveAction(act QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if act != nil && act.QAction_PTR() != nil {
		convArg0 = act.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu15setActiveActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * activeAction() const

/*
Returns the currently highlighted action, or 0 if no action is currently highlighted.

See also setActiveAction().
*/
func (this *QMenu) ActiveAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu12activeActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void popup(const QPoint &, QAction *)

/*
Displays the menu so that the action atAction will be at the specified global position p. To translate a widget's local coordinates into global coordinates, use QWidget::mapToGlobal().

When positioning a menu with exec() or popup(), bear in mind that you cannot rely on the menu's current size(). For performance reasons, the menu adapts its size only when necessary, so in many cases, the size before and after the show is different. Instead, use sizeHint() which calculates the proper size depending on the menu's current contents.

See also QWidget::mapToGlobal() and exec().
*/
func (this *QMenu) Popup(pos qtcore.QPoint_ITF, at QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if at != nil && at.QAction_PTR() != nil {
		convArg1 = at.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu5popupERK6QPointP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void popup(const QPoint &, QAction *)

/*
Displays the menu so that the action atAction will be at the specified global position p. To translate a widget's local coordinates into global coordinates, use QWidget::mapToGlobal().

When positioning a menu with exec() or popup(), bear in mind that you cannot rely on the menu's current size(). For performance reasons, the menu adapts its size only when necessary, so in many cases, the size before and after the show is different. Instead, use sizeHint() which calculates the proper size depending on the menu's current contents.

See also QWidget::mapToGlobal() and exec().
*/
func (this *QMenu) Popup__(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	// arg: 1, QAction *=Pointer, QAction=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu5popupERK6QPointP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * exec()

/*
Executes this menu synchronously.

This is equivalent to exec(pos()).

This returns the triggered QAction in either the popup menu or one of its submenus, or 0 if no item was triggered (normally because the user pressed Esc).

In most situations you'll want to specify the position yourself, for example, the current mouse position:


  exec(QCursor::pos());



or aligned to a widget:


  exec(somewidget.mapToGlobal(QPoint(0,0)));



or in reaction to a QMouseEvent *e:


  exec(e->globalPos());
*/
func (this *QMenu) Exec() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu4execEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:189
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * exec(const QPoint &, QAction *)

/*
Executes this menu synchronously.

This is equivalent to exec(pos()).

This returns the triggered QAction in either the popup menu or one of its submenus, or 0 if no item was triggered (normally because the user pressed Esc).

In most situations you'll want to specify the position yourself, for example, the current mouse position:


  exec(QCursor::pos());



or aligned to a widget:


  exec(somewidget.mapToGlobal(QPoint(0,0)));



or in reaction to a QMouseEvent *e:


  exec(e->globalPos());
*/
func (this *QMenu) Exec_1(pos qtcore.QPoint_ITF, at QAction_ITF /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if at != nil && at.QAction_PTR() != nil {
		convArg1 = at.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu4execERK6QPointP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:189
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * exec(const QPoint &, QAction *)

/*
Executes this menu synchronously.

This is equivalent to exec(pos()).

This returns the triggered QAction in either the popup menu or one of its submenus, or 0 if no item was triggered (normally because the user pressed Esc).

In most situations you'll want to specify the position yourself, for example, the current mouse position:


  exec(QCursor::pos());



or aligned to a widget:


  exec(somewidget.mapToGlobal(QPoint(0,0)));



or in reaction to a QMouseEvent *e:


  exec(e->globalPos());
*/
func (this *QMenu) Exec_1_(pos qtcore.QPoint_ITF) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	// arg: 1, QAction *=Pointer, QAction=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu4execERK6QPointP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:197
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QMenu) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:199
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect actionGeometry(QAction *) const

/*
Returns the geometry of action act.
*/
func (this *QMenu) ActionGeometry(arg0 QAction_ITF /*777 QAction **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAction_PTR() != nil {
		convArg0 = arg0.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu14actionGeometryEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * actionAt(const QPoint &) const

/*
Returns the item at pt; returns 0 if there is no item there.
*/
func (this *QMenu) ActionAt(arg0 qtcore.QPoint_ITF) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu8actionAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * menuAction() const

/*
Returns the action associated with this menu.
*/
func (this *QMenu) MenuAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu10menuActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmenu.h:204
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QMenu) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qmenu.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*

 */
func (this *QMenu) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:207
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QMenu) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qmenu.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QMenu) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNoReplayFor(QWidget *)

/*

 */
func (this *QMenu) SetNoReplayFor(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu14setNoReplayForEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:219
// index:0
// Public Visibility=Default Availability=Available
// [1] bool separatorsCollapsible() const

/*

 */
func (this *QMenu) SeparatorsCollapsible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu21separatorsCollapsibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSeparatorsCollapsible(bool)

/*

 */
func (this *QMenu) SetSeparatorsCollapsible(collapse bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu24setSeparatorsCollapsibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), collapse)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:222
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toolTipsVisible() const

/*

 */
func (this *QMenu) ToolTipsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu15toolTipsVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTipsVisible(bool)

/*

 */
func (this *QMenu) SetToolTipsVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu18setToolTipsVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToShow()

/*
This signal is emitted just before the menu is shown to the user.

See also aboutToHide() and show().
*/
func (this *QMenu) AboutToShow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu11aboutToShowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToHide()

/*
This signal is emitted just before the menu is hidden from the user.

This function was introduced in  Qt 4.2.

See also aboutToShow() and hide().
*/
func (this *QMenu) AboutToHide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu11aboutToHideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(QAction *)

/*
This signal is emitted when an action in this menu is triggered.

action is the action that caused the signal to be emitted.

Normally, you connect each menu action's triggered() signal to its own custom slot, but sometimes you will want to connect several actions to a single slot, for example, when you have a group of closely related actions, such as "left justify", "center", "right justify".

Note: This signal is emitted for the main parent menu in a hierarchy. Hence, only the parent menu needs to be connected to a slot; sub-menus need not be connected.

See also hovered() and QAction::triggered().
*/
func (this *QMenu) Triggered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9triggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered(QAction *)

/*
This signal is emitted when a menu action is highlighted; action is the action that caused the signal to be emitted.

Often this is used to update status information.

See also triggered() and QAction::hovered().
*/
func (this *QMenu) Hovered(action QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu7hoveredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:232
// index:0
// Protected Visibility=Default Availability=Available
// [4] int columnCount() const

/*
If a menu does not fit on the screen it lays itself out so that it does fit. It is style dependent what layout means (for example, on Windows it will use multiple columns).

This functions returns the number of columns necessary.
*/
func (this *QMenu) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qmenu.h:234
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QMenu) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:235
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QMenu) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:236
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QMenu) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QMenu) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:238
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QMenu) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:240
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QMenu) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:242
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void enterEvent(QEvent *)

/*
Reimplemented from QWidget::enterEvent().
*/
func (this *QMenu) EnterEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10enterEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:243
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)

/*
Reimplemented from QWidget::leaveEvent().
*/
func (this *QMenu) LeaveEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QMenu) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:245
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QMenu) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:246
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)

/*
Reimplemented from QWidget::actionEvent().
*/
func (this *QMenu) ActionEvent(arg0 qtgui.QActionEvent_ITF /*777 QActionEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QActionEvent_PTR() != nil {
		convArg0 = arg0.QActionEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:247
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QMenu) TimerEvent(arg0 qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTimerEvent_PTR() != nil {
		convArg0 = arg0.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenu.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QMenu) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:249
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QMenu) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QMenu18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmenu.h:250
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionMenuItem *, const QAction *) const

/*
Initialize option with the values from this menu and information from action. This method is useful for subclasses when they need a QStyleOptionMenuItem, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom() and QMenuBar::initStyleOption().
*/
func (this *QMenu) InitStyleOption(option QStyleOptionMenuItem_ITF /*777 QStyleOptionMenuItem **/, action QAction_ITF /*777 const QAction **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionMenuItem_PTR() != nil {
		convArg0 = option.QStyleOptionMenuItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg1 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
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
