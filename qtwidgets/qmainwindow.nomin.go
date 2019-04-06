//  header block begin

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

// /usr/include/qt/QtWidgets/qmainwindow.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated() const

/*

 */
func (this *QMainWindow) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDockNestingEnabled() const

/*

 */
func (this *QMainWindow) IsDockNestingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow20isDockNestingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode() const

/*

 */
func (this *QMainWindow) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(bool)

/*

 */
func (this *QMainWindow) SetDocumentMode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape() const

/*

 */
func (this *QMainWindow) TabShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow8tabShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(QTabWidget::TabShape)

/*

 */
func (this *QMainWindow) SetTabShape(tabShape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow11setTabShapeEN10QTabWidget8TabShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabShape)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition(Qt::DockWidgetArea) const

/*
Returns the tab position for area.

Note: The VerticalTabs dock option overrides the tab positions returned by this function.

This function was introduced in  Qt 4.5.

See also setTabPosition() and tabShape().
*/
func (this *QMainWindow) TabPosition(area int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11tabPositionEN2Qt14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(Qt::DockWidgetAreas, QTabWidget::TabPosition)

/*
Sets the tab position for the given dock widget areas to the specified tabPosition. By default, all dock areas show their tabs at the bottom.

Note: The VerticalTabs dock option overrides the tab positions set by this method.

This function was introduced in  Qt 4.5.

See also tabPosition() and setTabShape().
*/
func (this *QMainWindow) SetTabPosition(areas int, tabPosition int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow14setTabPositionE6QFlagsIN2Qt14DockWidgetAreaEEN10QTabWidget11TabPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), areas, tabPosition)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenuBar * menuBar() const

/*
Returns the menu bar for the main window. This function creates and returns an empty menu bar if the menu bar does not exist.

If you want all windows in a Mac application to share one menu bar, don't use this function to create it, because the menu bar created here will have this QMainWindow as its parent. Instead, you must create a menu bar that does not have a parent, which you can then share among all the Mac windows. Create a parent-less menu bar this way:


  QMenuBar *menuBar = new QMenuBar(0);



See also setMenuBar().
*/
func (this *QMainWindow) MenuBar() *QMenuBar /*777 QMenuBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow7menuBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuBar(QMenuBar *)

/*
Sets the menu bar for the main window to menuBar.

Note: QMainWindow takes ownership of the menuBar pointer and deletes it at the appropriate time.

See also menuBar().
*/
func (this *QMainWindow) SetMenuBar(menubar QMenuBar_ITF /*777 QMenuBar **/) {
	var convArg0 unsafe.Pointer
	if menubar != nil && menubar.QMenuBar_PTR() != nil {
		convArg0 = menubar.QMenuBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10setMenuBarEP8QMenuBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * menuWidget() const

/*
Returns the menu bar for the main window. This function returns null if a menu bar hasn't been constructed yet.

This function was introduced in  Qt 4.2.

See also setMenuWidget().
*/
func (this *QMainWindow) MenuWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow10menuWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuWidget(QWidget *)

/*
Sets the menu bar for the main window to menuBar.

QMainWindow takes ownership of the menuBar pointer and deletes it at the appropriate time.

This function was introduced in  Qt 4.2.

See also menuWidget().
*/
func (this *QMainWindow) SetMenuWidget(menubar QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if menubar != nil && menubar.QWidget_PTR() != nil {
		convArg0 = menubar.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13setMenuWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QStatusBar * statusBar() const

/*
Returns the status bar for the main window. This function creates and returns an empty status bar if the status bar does not exist.

See also setStatusBar().
*/
func (this *QMainWindow) StatusBar() *QStatusBar /*777 QStatusBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow9statusBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStatusBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatusBar(QStatusBar *)

/*
Sets the status bar for the main window to statusbar.

Setting the status bar to 0 will remove it from the main window. Note that QMainWindow takes ownership of the statusbar pointer and deletes it at the appropriate time.

See also statusBar().
*/
func (this *QMainWindow) SetStatusBar(statusbar QStatusBar_ITF /*777 QStatusBar **/) {
	var convArg0 unsafe.Pointer
	if statusbar != nil && statusbar.QStatusBar_PTR() != nil {
		convArg0 = statusbar.QStatusBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow12setStatusBarEP10QStatusBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCorner(Qt::Corner, Qt::DockWidgetArea)

/*
Sets the given dock widget area to occupy the specified corner.

See also corner().
*/
func (this *QMainWindow) SetCorner(corner int, area int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow9setCornerEN2Qt6CornerENS0_14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner, area)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DockWidgetArea corner(Qt::Corner) const

/*
Returns the dock widget area that occupies the specified corner.

See also setCorner().
*/
func (this *QMainWindow) Corner(corner int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow6cornerEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToolBarBreak(Qt::ToolBarArea)

/*
Adds a toolbar break to the given area after all the other objects that are present.
*/
func (this *QMainWindow) AddToolBarBreak(area int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15addToolBarBreakEN2Qt11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToolBarBreak(Qt::ToolBarArea)

/*
Adds a toolbar break to the given area after all the other objects that are present.
*/
func (this *QMainWindow) AddToolBarBreakp() {
	// arg: 0, Qt::ToolBarArea=Elaborated, Qt::ToolBarArea=Enum, , Invalid
	area := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15addToolBarBreakEN2Qt11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertToolBarBreak(QToolBar *)

/*
Inserts a toolbar break before the toolbar specified by before.
*/
func (this *QMainWindow) InsertToolBarBreak(before QToolBar_ITF /*777 QToolBar **/) {
	var convArg0 unsafe.Pointer
	if before != nil && before.QToolBar_PTR() != nil {
		convArg0 = before.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow18insertToolBarBreakEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToolBar(Qt::ToolBarArea, QToolBar *)

/*
Adds the toolbar into the specified area in this main window. The toolbar is placed at the end of the current tool bar block (i.e. line). If the main window already manages toolbar then it will only move the toolbar to area.

See also insertToolBar(), addToolBarBreak(), and insertToolBarBreak().
*/
func (this *QMainWindow) AddToolBar(area int, toolbar QToolBar_ITF /*777 QToolBar **/) {
	var convArg1 unsafe.Pointer
	if toolbar != nil && toolbar.QToolBar_PTR() != nil {
		convArg1 = toolbar.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEN2Qt11ToolBarAreaEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:153
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addToolBar(QToolBar *)

/*
Adds the toolbar into the specified area in this main window. The toolbar is placed at the end of the current tool bar block (i.e. line). If the main window already manages toolbar then it will only move the toolbar to area.

See also insertToolBar(), addToolBarBreak(), and insertToolBarBreak().
*/
func (this *QMainWindow) AddToolBar1(toolbar QToolBar_ITF /*777 QToolBar **/) {
	var convArg0 unsafe.Pointer
	if toolbar != nil && toolbar.QToolBar_PTR() != nil {
		convArg0 = toolbar.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:154
// index:2
// Public Visibility=Default Availability=Available
// [8] QToolBar * addToolBar(const QString &)

/*
Adds the toolbar into the specified area in this main window. The toolbar is placed at the end of the current tool bar block (i.e. line). If the main window already manages toolbar then it will only move the toolbar to area.

See also insertToolBar(), addToolBarBreak(), and insertToolBarBreak().
*/
func (this *QMainWindow) AddToolBar2(title string) *QToolBar /*777 QToolBar **/ {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertToolBar(QToolBar *, QToolBar *)

/*
Inserts the toolbar into the area occupied by the before toolbar so that it appears before it. For example, in normal left-to-right layout operation, this means that toolbar will appear to the left of the toolbar specified by before in a horizontal toolbar area.

See also insertToolBarBreak(), addToolBar(), and addToolBarBreak().
*/
func (this *QMainWindow) InsertToolBar(before QToolBar_ITF /*777 QToolBar **/, toolbar QToolBar_ITF /*777 QToolBar **/) {
	var convArg0 unsafe.Pointer
	if before != nil && before.QToolBar_PTR() != nil {
		convArg0 = before.QToolBar_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if toolbar != nil && toolbar.QToolBar_PTR() != nil {
		convArg1 = toolbar.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13insertToolBarEP8QToolBarS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeToolBar(QToolBar *)

/*
Removes the toolbar from the main window layout and hides it. Note that the toolbar is not deleted.
*/
func (this *QMainWindow) RemoveToolBar(toolbar QToolBar_ITF /*777 QToolBar **/) {
	var convArg0 unsafe.Pointer
	if toolbar != nil && toolbar.QToolBar_PTR() != nil {
		convArg0 = toolbar.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13removeToolBarEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeToolBarBreak(QToolBar *)

/*
Removes a toolbar break previously inserted before the toolbar specified by before.
*/
func (this *QMainWindow) RemoveToolBarBreak(before QToolBar_ITF /*777 QToolBar **/) {
	var convArg0 unsafe.Pointer
	if before != nil && before.QToolBar_PTR() != nil {
		convArg0 = before.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow18removeToolBarBreakEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:159
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unifiedTitleAndToolBarOnMac() const

/*

 */
func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:161
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolBarArea toolBarArea(QToolBar *) const

/*
Returns the Qt::ToolBarArea for toolbar. If toolbar has not been added to the main window, this function returns Qt::NoToolBarArea.

See also addToolBar(), addToolBarBreak(), and Qt::ToolBarArea.
*/
func (this *QMainWindow) ToolBarArea(toolbar QToolBar_ITF /*777 QToolBar **/) int {
	var convArg0 unsafe.Pointer
	if toolbar != nil && toolbar.QToolBar_PTR() != nil {
		convArg0 = toolbar.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11toolBarAreaEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toolBarBreak(QToolBar *) const

/*
Returns whether there is a toolbar break before the toolbar.

See also addToolBarBreak() and insertToolBarBreak().
*/
func (this *QMainWindow) ToolBarBreak(toolbar QToolBar_ITF /*777 QToolBar **/) bool {
	var convArg0 unsafe.Pointer
	if toolbar != nil && toolbar.QToolBar_PTR() != nil {
		convArg0 = toolbar.QToolBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow12toolBarBreakEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addDockWidget(Qt::DockWidgetArea, QDockWidget *)

/*
Adds the given dockwidget to the specified area.
*/
func (this *QMainWindow) AddDockWidget(area int, dockwidget QDockWidget_ITF /*777 QDockWidget **/) {
	var convArg1 unsafe.Pointer
	if dockwidget != nil && dockwidget.QDockWidget_PTR() != nil {
		convArg1 = dockwidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:166
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addDockWidget(Qt::DockWidgetArea, QDockWidget *, Qt::Orientation)

/*
Adds the given dockwidget to the specified area.
*/
func (this *QMainWindow) AddDockWidget1(area int, dockwidget QDockWidget_ITF /*777 QDockWidget **/, orientation int) {
	var convArg1 unsafe.Pointer
	if dockwidget != nil && dockwidget.QDockWidget_PTR() != nil {
		convArg1 = dockwidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidgetNS0_11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1, orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void splitDockWidget(QDockWidget *, QDockWidget *, Qt::Orientation)

/*
Splits the space covered by the first dock widget into two parts, moves the first dock widget into the first part, and moves the second dock widget into the second part.

The orientation specifies how the space is divided: A Qt::Horizontal split places the second dock widget to the right of the first; a Qt::Vertical split places the second dock widget below the first.

Note: if first is currently in a tabbed docked area, second will be added as a new tab, not as a neighbor of first. This is because a single tab can contain only one dock widget.

Note: The Qt::LayoutDirection influences the order of the dock widgets in the two parts of the divided area. When right-to-left layout direction is enabled, the placing of the dock widgets will be reversed.

See also tabifyDockWidget(), addDockWidget(), and removeDockWidget().
*/
func (this *QMainWindow) SplitDockWidget(after QDockWidget_ITF /*777 QDockWidget **/, dockwidget QDockWidget_ITF /*777 QDockWidget **/, orientation int) {
	var convArg0 unsafe.Pointer
	if after != nil && after.QDockWidget_PTR() != nil {
		convArg0 = after.QDockWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if dockwidget != nil && dockwidget.QDockWidget_PTR() != nil {
		convArg1 = dockwidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15splitDockWidgetEP11QDockWidgetS1_N2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabifyDockWidget(QDockWidget *, QDockWidget *)

/*
Moves second dock widget on top of first dock widget, creating a tabbed docked area in the main window.

See also tabifiedDockWidgets().
*/
func (this *QMainWindow) TabifyDockWidget(first QDockWidget_ITF /*777 QDockWidget **/, second QDockWidget_ITF /*777 QDockWidget **/) {
	var convArg0 unsafe.Pointer
	if first != nil && first.QDockWidget_PTR() != nil {
		convArg0 = first.QDockWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if second != nil && second.QDockWidget_PTR() != nil {
		convArg1 = second.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeDockWidget(QDockWidget *)

/*
Removes the dockwidget from the main window layout and hides it. Note that the dockwidget is not deleted.
*/
func (this *QMainWindow) RemoveDockWidget(dockwidget QDockWidget_ITF /*777 QDockWidget **/) {
	var convArg0 unsafe.Pointer
	if dockwidget != nil && dockwidget.QDockWidget_PTR() != nil {
		convArg0 = dockwidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16removeDockWidgetEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreDockWidget(QDockWidget *)

/*
Restores the state of dockwidget if it is created after the call to restoreState(). Returns true if the state was restored; otherwise returns false.

See also restoreState() and saveState().
*/
func (this *QMainWindow) RestoreDockWidget(dockwidget QDockWidget_ITF /*777 QDockWidget **/) bool {
	var convArg0 unsafe.Pointer
	if dockwidget != nil && dockwidget.QDockWidget_PTR() != nil {
		convArg0 = dockwidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DockWidgetArea dockWidgetArea(QDockWidget *) const

/*
Returns the Qt::DockWidgetArea for dockwidget. If dockwidget has not been added to the main window, this function returns Qt::NoDockWidgetArea.

See also addDockWidget(), splitDockWidget(), and Qt::DockWidgetArea.
*/
func (this *QMainWindow) DockWidgetArea(dockwidget QDockWidget_ITF /*777 QDockWidget **/) int {
	var convArg0 unsafe.Pointer
	if dockwidget != nil && dockwidget.QDockWidget_PTR() != nil {
		convArg0 = dockwidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMenu * createPopupMenu()

/*
Returns a popup menu containing checkable entries for the toolbars and dock widgets present in the main window. If there are no toolbars and dock widgets present, this function returns a null pointer.

By default, this function is called by the main window when the user activates a context menu, typically by right-clicking on a toolbar or a dock widget.

If you want to create a custom popup menu, reimplement this function and return a newly-created popup menu. Ownership of the popup menu is transferred to the caller.

See also addDockWidget(), addToolBar(), and menuBar().
*/
func (this *QMainWindow) CreatePopupMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15createPopupMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmainwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(bool)

/*

 */
func (this *QMainWindow) SetAnimated(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDockNestingEnabled(bool)

/*

 */
func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow21setDockNestingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnifiedTitleAndToolBarOnMac(bool)

/*

 */
func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), set)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabifiedDockWidgetActivated(QDockWidget *)

/*
This signal is emitted when the tabified dock widget is activated by selecting the tab. The activated dock widget is passed in dockWidget.

This function was introduced in  Qt 5.8.

See also tabifyDockWidget() and tabifiedDockWidgets().
*/
func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget QDockWidget_ITF /*777 QDockWidget **/) {
	var convArg0 unsafe.Pointer
	if dockWidget != nil && dockWidget.QDockWidget_PTR() != nil {
		convArg0 = dockWidget.QDockWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow27tabifiedDockWidgetActivatedEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11274() {
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
