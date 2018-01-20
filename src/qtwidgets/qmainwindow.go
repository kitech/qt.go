//  header block begin
// /usr/include/qt/QtWidgets/qmainwindow.h
// #include <qmainwindow.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 59
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
	*QWidget
}

func (this *QMainWindow) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qmainwindow.h:62
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMainWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// void QMainWindow(class QWidget *, Qt::WindowFlags)
func NewQMainWindow(parent unsafe.Pointer, flags int) *QMainWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMainWindowFromPointer(cthis)
	return gothis
}
func NewQMainWindowFromPointer(cthis unsafe.Pointer) *QMainWindow {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMainWindow{bcthis0}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:98
// index:0
// void setIconSize(const class QSize &)
func (this *QMainWindow) SetIconSize(iconSize unsafe.Pointer) {
	// 0: (, iconSize const QSize &), (iconSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), iconSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:100
// index:0
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QMainWindow) ToolButtonStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow15toolButtonStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:101
// index:0
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle int) {
	// 0: (, toolButtonStyle Qt::ToolButtonStyle), (&toolButtonStyle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:104
// index:0
// bool isAnimated()
func (this *QMainWindow) IsAnimated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10isAnimatedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:105
// index:0
// bool isDockNestingEnabled()
func (this *QMainWindow) IsDockNestingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow20isDockNestingEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:109
// index:0
// bool documentMode()
func (this *QMainWindow) DocumentMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow12documentModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:110
// index:0
// void setDocumentMode(_Bool)
func (this *QMainWindow) SetDocumentMode(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15setDocumentModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:114
// index:0
// QTabWidget::TabShape tabShape()
func (this *QMainWindow) TabShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow8tabShapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:115
// index:0
// void setTabShape(class QTabWidget::TabShape)
func (this *QMainWindow) SetTabShape(tabShape int) {
	// 0: (, tabShape QTabWidget::TabShape), (&tabShape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setTabShapeEN10QTabWidget8TabShapeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tabShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:116
// index:0
// QTabWidget::TabPosition tabPosition(Qt::DockWidgetArea)
func (this *QMainWindow) TabPosition(area int) {
	// 0: (, area Qt::DockWidgetArea), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11tabPositionEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:117
// index:0
// void setTabPosition(Qt::DockWidgetAreas, class QTabWidget::TabPosition)
func (this *QMainWindow) SetTabPosition(areas int, tabPosition int) {
	// 0: (, QFlags<Qt::DockWidgetArea> areas, tabPosition QTabWidget::TabPosition), (&areas, &tabPosition)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow14setTabPositionE6QFlagsIN2Qt14DockWidgetAreaEEN10QTabWidget11TabPositionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &areas, &tabPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:120
// index:0
// void setDockOptions(QMainWindow::DockOptions)
func (this *QMainWindow) SetDockOptions(options int) {
	// 0: (, QFlags<QMainWindow::DockOption> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow14setDockOptionsE6QFlagsINS_10DockOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:121
// index:0
// QMainWindow::DockOptions dockOptions()
func (this *QMainWindow) DockOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11dockOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:123
// index:0
// bool isSeparator(const class QPoint &)
func (this *QMainWindow) IsSeparator(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11isSeparatorERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:126
// index:0
// QMenuBar * menuBar()
func (this *QMainWindow) MenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow7menuBarEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:127
// index:0
// void setMenuBar(class QMenuBar *)
func (this *QMainWindow) SetMenuBar(menubar unsafe.Pointer) {
	// 0: (, menubar QMenuBar *), (menubar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10setMenuBarEP8QMenuBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), menubar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:129
// index:0
// QWidget * menuWidget()
func (this *QMainWindow) MenuWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10menuWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:130
// index:0
// void setMenuWidget(class QWidget *)
func (this *QMainWindow) SetMenuWidget(menubar unsafe.Pointer) {
	// 0: (, menubar QWidget *), (menubar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13setMenuWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), menubar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:134
// index:0
// QStatusBar * statusBar()
func (this *QMainWindow) StatusBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow9statusBarEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:135
// index:0
// void setStatusBar(class QStatusBar *)
func (this *QMainWindow) SetStatusBar(statusbar unsafe.Pointer) {
	// 0: (, statusbar QStatusBar *), (statusbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow12setStatusBarEP10QStatusBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), statusbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:138
// index:0
// QWidget * centralWidget()
func (this *QMainWindow) CentralWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow13centralWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:139
// index:0
// void setCentralWidget(class QWidget *)
func (this *QMainWindow) SetCentralWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16setCentralWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:141
// index:0
// QWidget * takeCentralWidget()
func (this *QMainWindow) TakeCentralWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow17takeCentralWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:144
// index:0
// void setCorner(Qt::Corner, Qt::DockWidgetArea)
func (this *QMainWindow) SetCorner(corner int, area int) {
	// 0: (, corner Qt::Corner, area Qt::DockWidgetArea), (&corner, &area)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow9setCornerEN2Qt6CornerENS0_14DockWidgetAreaE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &corner, &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:145
// index:0
// Qt::DockWidgetArea corner(Qt::Corner)
func (this *QMainWindow) Corner(corner int) {
	// 0: (, corner Qt::Corner), (&corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow6cornerEN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:149
// index:0
// void addToolBarBreak(Qt::ToolBarArea)
func (this *QMainWindow) AddToolBarBreak(area int) {
	// 0: (, area Qt::ToolBarArea), (&area)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15addToolBarBreakEN2Qt11ToolBarAreaE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:150
// index:0
// void insertToolBarBreak(class QToolBar *)
func (this *QMainWindow) InsertToolBarBreak(before unsafe.Pointer) {
	// 0: (, before QToolBar *), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18insertToolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:152
// index:0
// void addToolBar(Qt::ToolBarArea, class QToolBar *)
func (this *QMainWindow) AddToolBar(area int, toolbar unsafe.Pointer) {
	// 0: (, area Qt::ToolBarArea, toolbar QToolBar *), (&area, toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEN2Qt11ToolBarAreaEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:153
// index:1
// void addToolBar(class QToolBar *)
func (this *QMainWindow) AddToolBar_1(toolbar unsafe.Pointer) {
	// 1: (, toolbar QToolBar *), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:154
// index:2
// QToolBar * addToolBar(const class QString &)
func (this *QMainWindow) AddToolBar_2(title unsafe.Pointer) {
	// 2: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:155
// index:0
// void insertToolBar(class QToolBar *, class QToolBar *)
func (this *QMainWindow) InsertToolBar(before unsafe.Pointer, toolbar unsafe.Pointer) {
	// 0: (, before QToolBar *, toolbar QToolBar *), (before, toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13insertToolBarEP8QToolBarS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:156
// index:0
// void removeToolBar(class QToolBar *)
func (this *QMainWindow) RemoveToolBar(toolbar unsafe.Pointer) {
	// 0: (, toolbar QToolBar *), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13removeToolBarEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:157
// index:0
// void removeToolBarBreak(class QToolBar *)
func (this *QMainWindow) RemoveToolBarBreak(before unsafe.Pointer) {
	// 0: (, before QToolBar *), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18removeToolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:159
// index:0
// bool unifiedTitleAndToolBarOnMac()
func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:161
// index:0
// Qt::ToolBarArea toolBarArea(class QToolBar *)
func (this *QMainWindow) ToolBarArea(toolbar unsafe.Pointer) {
	// 0: (, toolbar QToolBar *), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11toolBarAreaEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:162
// index:0
// bool toolBarBreak(class QToolBar *)
func (this *QMainWindow) ToolBarBreak(toolbar unsafe.Pointer) {
	// 0: (, toolbar QToolBar *), (toolbar)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow12toolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toolbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:165
// index:0
// void addDockWidget(Qt::DockWidgetArea, class QDockWidget *)
func (this *QMainWindow) AddDockWidget(area int, dockwidget unsafe.Pointer) {
	// 0: (, area Qt::DockWidgetArea, dockwidget QDockWidget *), (&area, dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area, dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:166
// index:1
// void addDockWidget(Qt::DockWidgetArea, class QDockWidget *, Qt::Orientation)
func (this *QMainWindow) AddDockWidget_1(area int, dockwidget unsafe.Pointer, orientation int) {
	// 1: (, area Qt::DockWidgetArea, dockwidget QDockWidget *, orientation Qt::Orientation), (&area, dockwidget, &orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidgetNS0_11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &area, dockwidget, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:168
// index:0
// void splitDockWidget(class QDockWidget *, class QDockWidget *, Qt::Orientation)
func (this *QMainWindow) SplitDockWidget(after unsafe.Pointer, dockwidget unsafe.Pointer, orientation int) {
	// 0: (, after QDockWidget *, dockwidget QDockWidget *, orientation Qt::Orientation), (after, dockwidget, &orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15splitDockWidgetEP11QDockWidgetS1_N2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), after, dockwidget, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:170
// index:0
// void tabifyDockWidget(class QDockWidget *, class QDockWidget *)
func (this *QMainWindow) TabifyDockWidget(first unsafe.Pointer, second unsafe.Pointer) {
	// 0: (, first QDockWidget *, second QDockWidget *), (first, second)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), first, second)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:171
// index:0
// QList<QDockWidget *> tabifiedDockWidgets(class QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgets(dockwidget unsafe.Pointer) {
	// 0: (, dockwidget QDockWidget *), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow19tabifiedDockWidgetsEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:172
// index:0
// void removeDockWidget(class QDockWidget *)
func (this *QMainWindow) RemoveDockWidget(dockwidget unsafe.Pointer) {
	// 0: (, dockwidget QDockWidget *), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16removeDockWidgetEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:173
// index:0
// bool restoreDockWidget(class QDockWidget *)
func (this *QMainWindow) RestoreDockWidget(dockwidget unsafe.Pointer) {
	// 0: (, dockwidget QDockWidget *), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:175
// index:0
// Qt::DockWidgetArea dockWidgetArea(class QDockWidget *)
func (this *QMainWindow) DockWidgetArea(dockwidget unsafe.Pointer) {
	// 0: (, dockwidget QDockWidget *), (dockwidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dockwidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:181
// index:0
// QByteArray saveState(int)
func (this *QMainWindow) SaveState(version int) {
	// 0: (, version int), (&version)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow9saveStateEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:182
// index:0
// bool restoreState(const class QByteArray &, int)
func (this *QMainWindow) RestoreState(state unsafe.Pointer, version int) {
	// 0: (, state const QByteArray &, version int), (state, &version)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow12restoreStateERK10QByteArrayi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), state, &version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:185
// index:0
// virtual
// QMenu * createPopupMenu()
func (this *QMainWindow) CreatePopupMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15createPopupMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:190
// index:0
// void setAnimated(_Bool)
func (this *QMainWindow) SetAnimated(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setAnimatedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:191
// index:0
// void setDockNestingEnabled(_Bool)
func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow21setDockNestingEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:194
// index:0
// void setUnifiedTitleAndToolBarOnMac(_Bool)
func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	// 0: (, set bool), (&set)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:198
// index:0
// void iconSizeChanged(const class QSize &)
func (this *QMainWindow) IconSizeChanged(iconSize unsafe.Pointer) {
	// 0: (, iconSize const QSize &), (iconSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), iconSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:199
// index:0
// void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle int) {
	// 0: (, toolButtonStyle Qt::ToolButtonStyle), (&toolButtonStyle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:201
// index:0
// void tabifiedDockWidgetActivated(class QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget unsafe.Pointer) {
	// 0: (, dockWidget QDockWidget *), (dockWidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow27tabifiedDockWidgetActivatedEP11QDockWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dockWidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:206
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QMainWindow) ContextMenuEvent(event unsafe.Pointer) {
	// 0: (, event QContextMenuEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:208
// index:0
// virtual
// bool event(class QEvent *)
func (this *QMainWindow) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
