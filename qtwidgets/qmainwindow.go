// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmainwindow.h
// #include <qmainwindow.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 57
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

// void contextMenuEvent(QContextMenuEvent *)
func (this *QMainWindow) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// bool event(QEvent *)
func (this *QMainWindow) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QMainWindow struct {
	*QWidget
}
type QMainWindow_ITF interface {
	QWidget_ITF
	QMainWindow_PTR() *QMainWindow
}

func (ptr *QMainWindow) QMainWindow_PTR() *QMainWindow { return ptr }

func (this *QMainWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QMainWindow) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQMainWindowFromPointer(cthis unsafe.Pointer) *QMainWindow {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMainWindow{bcthis0}
}
func (*QMainWindow) NewFromPointer(cthis unsafe.Pointer) *QMainWindow {
	return NewQMainWindowFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMainWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)

/*
Constructs a QMainWindow with the given parent and the specified widget flags.

QMainWindow sets the Qt::Window flag itself, and will hence always be created as a top-level widget.
*/
func (*QMainWindow) NewForInherit(parent QWidget_ITF /*777 QWidget **/, flags int) *QMainWindow {
	return NewQMainWindow(parent, flags)
}
func NewQMainWindow(parent QWidget_ITF /*777 QWidget **/, flags int) *QMainWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMainWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMainWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)

/*
Constructs a QMainWindow with the given parent and the specified widget flags.

QMainWindow sets the Qt::Window flag itself, and will hence always be created as a top-level widget.
*/
func (*QMainWindow) NewForInheritp() *QMainWindow {
	return NewQMainWindowp()
}
func NewQMainWindowp() *QMainWindow {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMainWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMainWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)

/*
Constructs a QMainWindow with the given parent and the specified widget flags.

QMainWindow sets the Qt::Window flag itself, and will hence always be created as a top-level widget.
*/
func (*QMainWindow) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QMainWindow {
	return NewQMainWindowp1(parent)
}
func NewQMainWindowp1(parent QWidget_ITF /*777 QWidget **/) *QMainWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMainWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMainWindow")
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMainWindow()

/*

 */
func DeleteQMainWindow(this *QMainWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QMainWindow) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QMainWindow) SetIconSize(iconSize qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if iconSize != nil && iconSize.QSize_PTR() != nil {
		convArg0 = iconSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolButtonStyle toolButtonStyle() const

/*

 */
func (this *QMainWindow) ToolButtonStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow15toolButtonStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolButtonStyle(Qt::ToolButtonStyle)

/*

 */
func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow18setToolButtonStyleEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDockOptions(QMainWindow::DockOptions)

/*

 */
func (this *QMainWindow) SetDockOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow14setDockOptionsE6QFlagsINS_10DockOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] QMainWindow::DockOptions dockOptions() const

/*

 */
func (this *QMainWindow) DockOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11dockOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSeparator(const QPoint &) const

/*

 */
func (this *QMainWindow) IsSeparator(pos qtcore.QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11isSeparatorERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * centralWidget() const

/*
Returns the central widget for the main window. This function returns zero if the central widget has not been set.

See also setCentralWidget().
*/
func (this *QMainWindow) CentralWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow13centralWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCentralWidget(QWidget *)

/*
Sets the given widget to be the main window's central widget.

Note: QMainWindow takes ownership of the widget pointer and deletes it at the appropriate time.

See also centralWidget().
*/
func (this *QMainWindow) SetCentralWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16setCentralWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * takeCentralWidget()

/*
Removes the central widget from this main window.

The ownership of the removed widget is passed to the caller.

This function was introduced in  Qt 5.2.
*/
func (this *QMainWindow) TakeCentralWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow17takeCentralWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState(int) const

/*
Saves the current state of this mainwindow's toolbars and dockwidgets. This includes the corner settings which can be set with setCorner(). The version number is stored as part of the data.

The objectName property is used to identify each QToolBar and QDockWidget. You should make sure that this property is unique for each QToolBar and QDockWidget you add to the QMainWindow

To restore the saved state, pass the return value and version number to restoreState().

To save the geometry when the window closes, you can implement a close event like this:


  void MyMainWindow::closeEvent(QCloseEvent *event)
  {
      QSettings settings("MyCompany", "MyApp");
      settings.setValue("geometry", saveGeometry());
      settings.setValue("windowState", saveState());
      QMainWindow::closeEvent(event);
  }



See also restoreState(), QWidget::saveGeometry(), and QWidget::restoreGeometry().
*/
func (this *QMainWindow) SaveState(version int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow9saveStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), version)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState(int) const

/*
Saves the current state of this mainwindow's toolbars and dockwidgets. This includes the corner settings which can be set with setCorner(). The version number is stored as part of the data.

The objectName property is used to identify each QToolBar and QDockWidget. You should make sure that this property is unique for each QToolBar and QDockWidget you add to the QMainWindow

To restore the saved state, pass the return value and version number to restoreState().

To save the geometry when the window closes, you can implement a close event like this:


  void MyMainWindow::closeEvent(QCloseEvent *event)
  {
      QSettings settings("MyCompany", "MyApp");
      settings.setValue("geometry", saveGeometry());
      settings.setValue("windowState", saveState());
      QMainWindow::closeEvent(event);
  }



See also restoreState(), QWidget::saveGeometry(), and QWidget::restoreGeometry().
*/
func (this *QMainWindow) SaveStatep() *qtcore.QByteArray /*123*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	version := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow9saveStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), version)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &, int)

/*
Restores the state of this mainwindow's toolbars and dockwidgets. Also restores the corner settings too. The version number is compared with that stored in state. If they do not match, the mainwindow's state is left unchanged, and this function returns false; otherwise, the state is restored, and this function returns true.

To restore geometry saved using QSettings, you can use code like this:


  void MainWindow::readSettings()
  {
      QSettings settings("MyCompany", "MyApp");
      restoreGeometry(settings.value("myWidget/geometry").toByteArray());
      restoreState(settings.value("myWidget/windowState").toByteArray());
  }



See also saveState(), QWidget::saveGeometry(), QWidget::restoreGeometry(), and restoreDockWidget().
*/
func (this *QMainWindow) RestoreState(state qtcore.QByteArray_ITF, version int) bool {
	var convArg0 unsafe.Pointer
	if state != nil && state.QByteArray_PTR() != nil {
		convArg0 = state.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow12restoreStateERK10QByteArrayi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, version)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &, int)

/*
Restores the state of this mainwindow's toolbars and dockwidgets. Also restores the corner settings too. The version number is compared with that stored in state. If they do not match, the mainwindow's state is left unchanged, and this function returns false; otherwise, the state is restored, and this function returns true.

To restore geometry saved using QSettings, you can use code like this:


  void MainWindow::readSettings()
  {
      QSettings settings("MyCompany", "MyApp");
      restoreGeometry(settings.value("myWidget/geometry").toByteArray());
      restoreState(settings.value("myWidget/windowState").toByteArray());
  }



See also saveState(), QWidget::saveGeometry(), QWidget::restoreGeometry(), and restoreDockWidget().
*/
func (this *QMainWindow) RestoreStatep(state qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if state != nil && state.QByteArray_PTR() != nil {
		convArg0 = state.QByteArray_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	version := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow12restoreStateERK10QByteArrayi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, version)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconSizeChanged(const QSize &)

/*
This signal is emitted when the size of the icons used in the window is changed. The new icon size is passed in iconSize.

You can connect this signal to other components to help maintain a consistent appearance for your application.

See also setIconSize().
*/
func (this *QMainWindow) IconSizeChanged(iconSize qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if iconSize != nil && iconSize.QSize_PTR() != nil {
		convArg0 = iconSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15iconSizeChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toolButtonStyleChanged(Qt::ToolButtonStyle)

/*
This signal is emitted when the style used for tool buttons in the window is changed. The new style is passed in toolButtonStyle.

You can connect this signal to other components to help maintain a consistent appearance for your application.

See also setToolButtonStyle().
*/
func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QMainWindow) ContextMenuEvent(event qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QContextMenuEvent_PTR() != nil {
		convArg0 = event.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QWidget::event().
*/
func (this *QMainWindow) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QMainWindow__DockOption = int

//
const QMainWindow__AnimatedDocks QMainWindow__DockOption = 1

//
const QMainWindow__AllowNestedDocks QMainWindow__DockOption = 2

//
const QMainWindow__AllowTabbedDocks QMainWindow__DockOption = 4

//
const QMainWindow__ForceTabbedDocks QMainWindow__DockOption = 8

//
const QMainWindow__VerticalTabs QMainWindow__DockOption = 16

//
const QMainWindow__GroupedDragging QMainWindow__DockOption = 32

func (this *QMainWindow) DockOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMainWindow_DockOptionItemName(val int) string {
	var nilthis *QMainWindow
	return nilthis.DockOptionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11273() {
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
