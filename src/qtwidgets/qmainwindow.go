//  header block begin
// /usr/include/qt/QtWidgets/qmainwindow.h
// #include <qmainwindow.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 49
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
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
type QMainWindow struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qmainwindow.h:62
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMainWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:95
// index:0
// virtual
// void ~QMainWindow()
func DeleteQMainWindow(*QMainWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:97
// index:0
// QSize iconSize()
func (this *QMainWindow) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:98
// index:0
// void setIconSize(const class QSize &)
func (this *QMainWindow) SetIconSize(iconSize unsafe.Pointer) {
	// 0: (, const QSize & iconSize), (iconSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, iconSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:100
// index:0
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QMainWindow) ToolButtonStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow15toolButtonStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:101
// index:0
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle int) {
	// 0: (, Qt::ToolButtonStyle toolButtonStyle), (&toolButtonStyle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:104
// index:0
// bool isAnimated()
func (this *QMainWindow) IsAnimated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10isAnimatedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:105
// index:0
// bool isDockNestingEnabled()
func (this *QMainWindow) IsDockNestingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow20isDockNestingEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:109
// index:0
// bool documentMode()
func (this *QMainWindow) DocumentMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow12documentModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:110
// index:0
// void setDocumentMode(_Bool)
func (this *QMainWindow) SetDocumentMode(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15setDocumentModeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:114
// index:0
// QTabWidget::TabShape tabShape()
func (this *QMainWindow) TabShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow8tabShapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:115
// index:0
// void setTabShape(class QTabWidget::TabShape)
func (this *QMainWindow) SetTabShape(tabShape int) {
	// 0: (, QTabWidget::TabShape tabShape), (&tabShape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setTabShapeEN10QTabWidget8TabShapeE", ffiqt.FFI_TYPE_VOID, this.cthis, &tabShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:116
// index:0
// QTabWidget::TabPosition tabPosition(Qt::DockWidgetArea)
func (this *QMainWindow) TabPosition(area int) {
	// 0: (, Qt::DockWidgetArea area), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11tabPositionEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.cthis, &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:121
// index:0
// QMainWindow::DockOptions dockOptions()
func (this *QMainWindow) DockOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11dockOptionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:123
// index:0
// bool isSeparator(const class QPoint &)
func (this *QMainWindow) IsSeparator(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11isSeparatorERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:126
// index:0
// QMenuBar * menuBar()
func (this *QMainWindow) MenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow7menuBarEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:127
// index:0
// void setMenuBar(class QMenuBar *)
func (this *QMainWindow) SetMenuBar(menubar unsafe.Pointer) {
	// 0: (, QMenuBar * menubar), (menubar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10setMenuBarEP8QMenuBar", ffiqt.FFI_TYPE_VOID, this.cthis, menubar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:129
// index:0
// QWidget * menuWidget()
func (this *QMainWindow) MenuWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10menuWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:130
// index:0
// void setMenuWidget(class QWidget *)
func (this *QMainWindow) SetMenuWidget(menubar unsafe.Pointer) {
	// 0: (, QWidget * menubar), (menubar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13setMenuWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, menubar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:134
// index:0
// QStatusBar * statusBar()
func (this *QMainWindow) StatusBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow9statusBarEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:135
// index:0
// void setStatusBar(class QStatusBar *)
func (this *QMainWindow) SetStatusBar(statusbar unsafe.Pointer) {
	// 0: (, QStatusBar * statusbar), (statusbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow12setStatusBarEP10QStatusBar", ffiqt.FFI_TYPE_VOID, this.cthis, statusbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:138
// index:0
// QWidget * centralWidget()
func (this *QMainWindow) CentralWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow13centralWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:139
// index:0
// void setCentralWidget(class QWidget *)
func (this *QMainWindow) SetCentralWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16setCentralWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:141
// index:0
// QWidget * takeCentralWidget()
func (this *QMainWindow) TakeCentralWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow17takeCentralWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:144
// index:0
// void setCorner(Qt::Corner, Qt::DockWidgetArea)
func (this *QMainWindow) SetCorner(corner int, area int) {
	// 0: (, Qt::Corner corner, Qt::DockWidgetArea area), (&corner, &area)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow9setCornerEN2Qt6CornerENS0_14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.cthis, &corner, &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:145
// index:0
// Qt::DockWidgetArea corner(Qt::Corner)
func (this *QMainWindow) Corner(corner int) {
	// 0: (, Qt::Corner corner), (&corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow6cornerEN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.cthis, &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:149
// index:0
// void addToolBarBreak(Qt::ToolBarArea)
func (this *QMainWindow) AddToolBarBreak(area int) {
	// 0: (, Qt::ToolBarArea area), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15addToolBarBreakEN2Qt11ToolBarAreaE", ffiqt.FFI_TYPE_VOID, this.cthis, &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:150
// index:0
// void insertToolBarBreak(class QToolBar *)
func (this *QMainWindow) InsertToolBarBreak(before unsafe.Pointer) {
	// 0: (, QToolBar * before), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18insertToolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:152
// index:0
// void addToolBar(Qt::ToolBarArea, class QToolBar *)
func (this *QMainWindow) AddToolBar(area int, toolbar unsafe.Pointer) {
	// 0: (, Qt::ToolBarArea area, QToolBar * toolbar), (&area, toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEN2Qt11ToolBarAreaEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, &area, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:153
// index:1
// void addToolBar(class QToolBar *)
func (this *QMainWindow) AddToolBar_1(toolbar unsafe.Pointer) {
	// 1: (, QToolBar * toolbar), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:154
// index:2
// QToolBar * addToolBar(const class QString &)
func (this *QMainWindow) AddToolBar_2(title unsafe.Pointer) {
	// 2: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:155
// index:0
// void insertToolBar(class QToolBar *, class QToolBar *)
func (this *QMainWindow) InsertToolBar(before unsafe.Pointer, toolbar unsafe.Pointer) {
	// 0: (, QToolBar * before, QToolBar * toolbar), (before, toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13insertToolBarEP8QToolBarS1_", ffiqt.FFI_TYPE_VOID, this.cthis, before, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:156
// index:0
// void removeToolBar(class QToolBar *)
func (this *QMainWindow) RemoveToolBar(toolbar unsafe.Pointer) {
	// 0: (, QToolBar * toolbar), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13removeToolBarEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:157
// index:0
// void removeToolBarBreak(class QToolBar *)
func (this *QMainWindow) RemoveToolBarBreak(before unsafe.Pointer) {
	// 0: (, QToolBar * before), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18removeToolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:159
// index:0
// bool unifiedTitleAndToolBarOnMac()
func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:161
// index:0
// Qt::ToolBarArea toolBarArea(class QToolBar *)
func (this *QMainWindow) ToolBarArea(toolbar unsafe.Pointer) {
	// 0: (, QToolBar * toolbar), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11toolBarAreaEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:162
// index:0
// bool toolBarBreak(class QToolBar *)
func (this *QMainWindow) ToolBarBreak(toolbar unsafe.Pointer) {
	// 0: (, QToolBar * toolbar), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow12toolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.cthis, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:165
// index:0
// void addDockWidget(Qt::DockWidgetArea, class QDockWidget *)
func (this *QMainWindow) AddDockWidget(area int, dockwidget unsafe.Pointer) {
	// 0: (, Qt::DockWidgetArea area, QDockWidget * dockwidget), (&area, dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &area, dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:166
// index:1
// void addDockWidget(Qt::DockWidgetArea, class QDockWidget *, Qt::Orientation)
func (this *QMainWindow) AddDockWidget_1(area int, dockwidget unsafe.Pointer, orientation int) {
	// 1: (, Qt::DockWidgetArea area, QDockWidget * dockwidget, Qt::Orientation orientation), (&area, dockwidget, &orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidgetNS0_11OrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &area, dockwidget, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:168
// index:0
// void splitDockWidget(class QDockWidget *, class QDockWidget *, Qt::Orientation)
func (this *QMainWindow) SplitDockWidget(after unsafe.Pointer, dockwidget unsafe.Pointer, orientation int) {
	// 0: (, QDockWidget * after, QDockWidget * dockwidget, Qt::Orientation orientation), (after, dockwidget, &orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15splitDockWidgetEP11QDockWidgetS1_N2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, after, dockwidget, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:170
// index:0
// void tabifyDockWidget(class QDockWidget *, class QDockWidget *)
func (this *QMainWindow) TabifyDockWidget(first unsafe.Pointer, second unsafe.Pointer) {
	// 0: (, QDockWidget * first, QDockWidget * second), (first, second)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_", ffiqt.FFI_TYPE_VOID, this.cthis, first, second)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:171
// index:0
// QList<QDockWidget *> tabifiedDockWidgets(class QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgets(dockwidget unsafe.Pointer) {
	// 0: (, QDockWidget * dockwidget), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.cthis, dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:172
// index:0
// void removeDockWidget(class QDockWidget *)
func (this *QMainWindow) RemoveDockWidget(dockwidget unsafe.Pointer) {
	// 0: (, QDockWidget * dockwidget), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16removeDockWidgetEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.cthis, dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:173
// index:0
// bool restoreDockWidget(class QDockWidget *)
func (this *QMainWindow) RestoreDockWidget(dockwidget unsafe.Pointer) {
	// 0: (, QDockWidget * dockwidget), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.cthis, dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:175
// index:0
// Qt::DockWidgetArea dockWidgetArea(class QDockWidget *)
func (this *QMainWindow) DockWidgetArea(dockwidget unsafe.Pointer) {
	// 0: (, QDockWidget * dockwidget), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.cthis, dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:181
// index:0
// QByteArray saveState(int)
func (this *QMainWindow) SaveState(version int) {
	// 0: (, int version), (&version)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow9saveStateEi", ffiqt.FFI_TYPE_VOID, this.cthis, &version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:182
// index:0
// bool restoreState(const class QByteArray &, int)
func (this *QMainWindow) RestoreState(state unsafe.Pointer, version int) {
	// 0: (, const QByteArray & state, int version), (state, &version)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow12restoreStateERK10QByteArrayi", ffiqt.FFI_TYPE_VOID, this.cthis, state, &version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:185
// index:0
// virtual
// QMenu * createPopupMenu()
func (this *QMainWindow) CreatePopupMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15createPopupMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:190
// index:0
// void setAnimated(_Bool)
func (this *QMainWindow) SetAnimated(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setAnimatedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:191
// index:0
// void setDockNestingEnabled(_Bool)
func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow21setDockNestingEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:194
// index:0
// void setUnifiedTitleAndToolBarOnMac(_Bool)
func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	// 0: (, bool set), (&set)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb", ffiqt.FFI_TYPE_VOID, this.cthis, &set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:198
// index:0
// void iconSizeChanged(const class QSize &)
func (this *QMainWindow) IconSizeChanged(iconSize unsafe.Pointer) {
	// 0: (, const QSize & iconSize), (iconSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, iconSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:199
// index:0
// void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle int) {
	// 0: (, Qt::ToolButtonStyle toolButtonStyle), (&toolButtonStyle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:201
// index:0
// void tabifiedDockWidgetActivated(class QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget unsafe.Pointer) {
	// 0: (, QDockWidget * dockWidget), (dockWidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow27tabifiedDockWidgetActivatedEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.cthis, dockWidget)
	gopp.ErrPrint(err, rv)
}

//  body block end
