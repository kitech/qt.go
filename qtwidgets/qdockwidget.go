package qtwidgets

// /usr/include/qt/QtWidgets/qdockwidget.h
// #include <qdockwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
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

// void changeEvent(QEvent *)
func (this *QDockWidget) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(QCloseEvent *)
func (this *QDockWidget) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QDockWidget) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// bool event(QEvent *)
func (this *QDockWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(QStyleOptionDockWidget *)
func (this *QDockWidget) InheritInitStyleOption(f func(option *QStyleOptionDockWidget /*777 QStyleOptionDockWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QDockWidget struct {
	*QWidget
}
type QDockWidget_ITF interface {
	QWidget_ITF
	QDockWidget_PTR() *QDockWidget
}

func (ptr *QDockWidget) QDockWidget_PTR() *QDockWidget { return ptr }

func (this *QDockWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDockWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDockWidgetFromPointer(cthis unsafe.Pointer) *QDockWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDockWidget{bcthis0}
}
func (*QDockWidget) NewFromPointer(cthis unsafe.Pointer) *QDockWidget {
	return NewQDockWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDockWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdockwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDockWidget(const QString &, QWidget *, Qt::WindowFlags)

/*
Constructs a QDockWidget with parent parent and window flags flags. The dock widget will be placed in the left dock widget area.

The window title is set to title. This title is used when the QDockWidget is docked and undocked. It is also used in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQDockWidget(title string, parent QWidget_ITF /*777 QWidget **/, flags int) *QDockWidget {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDockWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDockWidget(const QString &, QWidget *, Qt::WindowFlags)

/*
Constructs a QDockWidget with parent parent and window flags flags. The dock widget will be placed in the left dock widget area.

The window title is set to title. This title is used when the QDockWidget is docked and undocked. It is also used in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQDockWidget__(title string) *QDockWidget {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDockWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDockWidget(const QString &, QWidget *, Qt::WindowFlags)

/*
Constructs a QDockWidget with parent parent and window flags flags. The dock widget will be placed in the left dock widget area.

The window title is set to title. This title is used when the QDockWidget is docked and undocked. It is also used in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQDockWidget__1(title string, parent QWidget_ITF /*777 QWidget **/) *QDockWidget {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDockWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDockWidget(QWidget *, Qt::WindowFlags)

/*
Constructs a QDockWidget with parent parent and window flags flags. The dock widget will be placed in the left dock widget area.

The window title is set to title. This title is used when the QDockWidget is docked and undocked. It is also used in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQDockWidget_1(parent QWidget_ITF /*777 QWidget **/, flags int) *QDockWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDockWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDockWidget(QWidget *, Qt::WindowFlags)

/*
Constructs a QDockWidget with parent parent and window flags flags. The dock widget will be placed in the left dock widget area.

The window title is set to title. This title is used when the QDockWidget is docked and undocked. It is also used in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQDockWidget_1_() *QDockWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDockWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDockWidget(QWidget *, Qt::WindowFlags)

/*
Constructs a QDockWidget with parent parent and window flags flags. The dock widget will be placed in the left dock widget area.

The window title is set to title. This title is used when the QDockWidget is docked and undocked. It is also used in the context menu provided by QMainWindow.

See also setWindowTitle().
*/
func NewQDockWidget_1_1(parent QWidget_ITF /*777 QWidget **/) *QDockWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDockWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdockwidget.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDockWidget()

/*

 */
func DeleteQDockWidget(this *QDockWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the widget for the dock widget. This function returns zero if the widget has not been set.

See also setWidget().
*/
func (this *QDockWidget) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdockwidget.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*
Sets the widget for the dock widget to widget.

If the dock widget is visible when widget is added, you must show() it explicitly.

Note that you must add the layout of the widget before you call this function; if not, the widget will not be visible.

See also widget().
*/
func (this *QDockWidget) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFeatures(QDockWidget::DockWidgetFeatures)

/*

 */
func (this *QDockWidget) SetFeatures(features int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget11setFeaturesE6QFlagsINS_17DockWidgetFeatureEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), features)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QDockWidget::DockWidgetFeatures features() const

/*

 */
func (this *QDockWidget) Features() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget8featuresEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFloating(bool)

/*

 */
func (this *QDockWidget) SetFloating(floating bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget11setFloatingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), floating)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFloating() const

/*

 */
func (this *QDockWidget) IsFloating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget10isFloatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdockwidget.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAllowedAreas(Qt::DockWidgetAreas)

/*

 */
func (this *QDockWidget) SetAllowedAreas(areas int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget15setAllowedAreasE6QFlagsIN2Qt14DockWidgetAreaEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), areas)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DockWidgetAreas allowedAreas() const

/*

 */
func (this *QDockWidget) AllowedAreas() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget12allowedAreasEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitleBarWidget(QWidget *)

/*
Sets an arbitrary widget as the dock widget's title bar. If widget is 0, any custom title bar widget previously set on the dock widget is removed, but not deleted, and the default title bar will be used instead.

If a title bar widget is set, QDockWidget will not use native window decorations when it is floated.

Here are some tips for implementing custom title bars:


Mouse events that are not explicitly handled by the title bar widget must be ignored by calling QMouseEvent::ignore(). These events then propagate to the QDockWidget parent, which handles them in the usual manner, moving when the title bar is dragged, docking and undocking when it is double-clicked, etc.
When DockWidgetVerticalTitleBar is set on QDockWidget, the title bar widget is repositioned accordingly. In resizeEvent(), the title bar should check what orientation it should assume:
  QDockWidget *dockWidget = qobject_cast<QDockWidget*>(parentWidget());
  if (dockWidget->features() & QDockWidget::DockWidgetVerticalTitleBar) {
      // I need to be vertical
  } else {
      // I need to be horizontal
  }



The title bar widget must have a valid QWidget::sizeHint() and QWidget::minimumSizeHint(). These functions should take into account the current orientation of the title bar.
It is not possible to remove a title bar from a dock widget. However, a similar effect can be achieved by setting a default constructed QWidget as the title bar widget.


Using qobject_cast() as shown above, the title bar widget has full access to its parent QDockWidget. Hence it can perform such operations as docking and hiding in response to user actions.

This function was introduced in  Qt 4.3.

See also titleBarWidget() and DockWidgetVerticalTitleBar.
*/
func (this *QDockWidget) SetTitleBarWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget17setTitleBarWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * titleBarWidget() const

/*
Returns the custom title bar widget set on the QDockWidget, or 0 if no custom title bar has been set.

This function was introduced in  Qt 4.3.

See also setTitleBarWidget().
*/
func (this *QDockWidget) TitleBarWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget14titleBarWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdockwidget.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAreaAllowed(Qt::DockWidgetArea) const

/*
Returns true if this dock widget can be placed in the given area; otherwise returns false.
*/
func (this *QDockWidget) IsAreaAllowed(area int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget13isAreaAllowedEN2Qt14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdockwidget.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * toggleViewAction() const

/*
Returns a checkable action that can be used to show or close this dock widget.

The action's text is set to the dock widget's window title.

See also QAction::text and QWidget::windowTitle.
*/
func (this *QDockWidget) ToggleViewAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget16toggleViewActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdockwidget.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void featuresChanged(QDockWidget::DockWidgetFeatures)

/*
This signal is emitted when the features property changes. The features parameter gives the new value of the property.

Note: Notifier signal for property features.
*/
func (this *QDockWidget) FeaturesChanged(features int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget15featuresChangedE6QFlagsINS_17DockWidgetFeatureEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), features)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void topLevelChanged(bool)

/*
This signal is emitted when the floating property changes. The topLevel parameter is true if the dock widget is now floating; otherwise it is false.

See also isWindow().
*/
func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget15topLevelChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), topLevel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void allowedAreasChanged(Qt::DockWidgetAreas)

/*
This signal is emitted when the allowedAreas property changes. The allowedAreas parameter gives the new value of the property.

Note: Notifier signal for property allowedAreas.
*/
func (this *QDockWidget) AllowedAreasChanged(allowedAreas int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget19allowedAreasChangedE6QFlagsIN2Qt14DockWidgetAreaEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), allowedAreas)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibilityChanged(bool)

/*
This signal is emitted when the dock widget becomes visible (or invisible). This happens when the widget is hidden or shown, as well as when it is docked in a tabbed dock area and its tab becomes selected or unselected.

This function was introduced in  Qt 4.3.
*/
func (this *QDockWidget) VisibilityChanged(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget17visibilityChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dockLocationChanged(Qt::DockWidgetArea)

/*
This signal is emitted when the dock widget is moved to another dock area, or is moved to a different location in its current dock area. This happens when the dock widget is moved programmatically or is dragged to a new location by the user.

This function was introduced in  Qt 4.3.
*/
func (this *QDockWidget) DockLocationChanged(area int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget19dockLocationChangedEN2Qt14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QDockWidget) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)

/*
Reimplemented from QWidget::closeEvent().
*/
func (this *QDockWidget) CloseEvent(event qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QCloseEvent_PTR() != nil {
		convArg0 = event.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QDockWidget) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdockwidget.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QDockWidget) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDockWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdockwidget.h:120
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionDockWidget *) const

/*
Initialize option with the values from this QDockWidget. This method is useful for subclasses when they need a QStyleOptionDockWidget, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QDockWidget) InitStyleOption(option QStyleOptionDockWidget_ITF /*777 QStyleOptionDockWidget **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionDockWidget_PTR() != nil {
		convArg0 = option.QStyleOptionDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QDockWidget__DockWidgetFeature = int

//
const QDockWidget__DockWidgetClosable QDockWidget__DockWidgetFeature = 1

//
const QDockWidget__DockWidgetMovable QDockWidget__DockWidgetFeature = 2

//
const QDockWidget__DockWidgetFloatable QDockWidget__DockWidgetFeature = 4

//
const QDockWidget__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = 8

//
const QDockWidget__DockWidgetFeatureMask QDockWidget__DockWidgetFeature = 15

//
const QDockWidget__AllDockWidgetFeatures QDockWidget__DockWidgetFeature = 7

//
const QDockWidget__NoDockWidgetFeatures QDockWidget__DockWidgetFeature = 0

//
const QDockWidget__Reserved QDockWidget__DockWidgetFeature = 255

func (this *QDockWidget) DockWidgetFeatureItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDockWidget_DockWidgetFeatureItemName(val int) string {
	var nilthis *QDockWidget
	return nilthis.DockWidgetFeatureItemName(val)
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
