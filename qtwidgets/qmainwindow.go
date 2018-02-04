package qtwidgets

// /usr/include/qt/QtWidgets/qmainwindow.h
// #include <qmainwindow.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QMainWindow) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// bool event(class QEvent *)
func (this *QMainWindow) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QMainWindow struct {
	*QWidget
}

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
// [8] const QMetaObject * metaObject()
func (this *QMainWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMainWindow(QWidget *, Qt::WindowFlags)
func NewQMainWindow(parent *QWidget /*777 QWidget **/, flags int) *QMainWindow {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMainWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMainWindow()
func DeleteQMainWindow(this *QMainWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
func (this *QMainWindow) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QMainWindow) SetIconSize(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolButtonStyle toolButtonStyle()
func (this *QMainWindow) ToolButtonStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow15toolButtonStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow18setToolButtonStyleEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated()
func (this *QMainWindow) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDockNestingEnabled()
func (this *QMainWindow) IsDockNestingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow20isDockNestingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode()
func (this *QMainWindow) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(_Bool)
func (this *QMainWindow) SetDocumentMode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape()
func (this *QMainWindow) TabShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow8tabShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(QTabWidget::TabShape)
func (this *QMainWindow) SetTabShape(tabShape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow11setTabShapeEN10QTabWidget8TabShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition(Qt::DockWidgetArea)
func (this *QMainWindow) TabPosition(area int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11tabPositionEN2Qt14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(Qt::DockWidgetAreas, QTabWidget::TabPosition)
func (this *QMainWindow) SetTabPosition(areas int, tabPosition int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow14setTabPositionE6QFlagsIN2Qt14DockWidgetAreaEEN10QTabWidget11TabPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), areas, tabPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDockOptions(QMainWindow::DockOptions)
func (this *QMainWindow) SetDockOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow14setDockOptionsE6QFlagsINS_10DockOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] QMainWindow::DockOptions dockOptions()
func (this *QMainWindow) DockOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11dockOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSeparator(const QPoint &)
func (this *QMainWindow) IsSeparator(pos *qtcore.QPoint) bool {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11isSeparatorERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenuBar * menuBar()
func (this *QMainWindow) MenuBar() *QMenuBar /*777 QMenuBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow7menuBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuBar(QMenuBar *)
func (this *QMainWindow) SetMenuBar(menubar *QMenuBar /*777 QMenuBar **/) {
	var convArg0 = menubar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10setMenuBarEP8QMenuBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * menuWidget()
func (this *QMainWindow) MenuWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow10menuWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuWidget(QWidget *)
func (this *QMainWindow) SetMenuWidget(menubar *QWidget /*777 QWidget **/) {
	var convArg0 = menubar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13setMenuWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QStatusBar * statusBar()
func (this *QMainWindow) StatusBar() *QStatusBar /*777 QStatusBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow9statusBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStatusBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatusBar(QStatusBar *)
func (this *QMainWindow) SetStatusBar(statusbar *QStatusBar /*777 QStatusBar **/) {
	var convArg0 = statusbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow12setStatusBarEP10QStatusBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * centralWidget()
func (this *QMainWindow) CentralWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow13centralWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCentralWidget(QWidget *)
func (this *QMainWindow) SetCentralWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16setCentralWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * takeCentralWidget()
func (this *QMainWindow) TakeCentralWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow17takeCentralWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCorner(Qt::Corner, Qt::DockWidgetArea)
func (this *QMainWindow) SetCorner(corner int, area int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow9setCornerEN2Qt6CornerENS0_14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner, area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DockWidgetArea corner(Qt::Corner)
func (this *QMainWindow) Corner(corner int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow6cornerEN2Qt6CornerE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToolBarBreak(Qt::ToolBarArea)
func (this *QMainWindow) AddToolBarBreak(area int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15addToolBarBreakEN2Qt11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertToolBarBreak(QToolBar *)
func (this *QMainWindow) InsertToolBarBreak(before *QToolBar /*777 QToolBar **/) {
	var convArg0 = before.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow18insertToolBarBreakEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToolBar(Qt::ToolBarArea, QToolBar *)
func (this *QMainWindow) AddToolBar(area int, toolbar *QToolBar /*777 QToolBar **/) {
	var convArg1 = toolbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEN2Qt11ToolBarAreaEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:153
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addToolBar(QToolBar *)
func (this *QMainWindow) AddToolBar_1(toolbar *QToolBar /*777 QToolBar **/) {
	var convArg0 = toolbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:154
// index:2
// Public Visibility=Default Availability=Available
// [8] QToolBar * addToolBar(const QString &)
func (this *QMainWindow) AddToolBar_2(title *qtcore.QString) *QToolBar /*777 QToolBar **/ {
	var convArg0 = title.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertToolBar(QToolBar *, QToolBar *)
func (this *QMainWindow) InsertToolBar(before *QToolBar /*777 QToolBar **/, toolbar *QToolBar /*777 QToolBar **/) {
	var convArg0 = before.GetCthis()
	var convArg1 = toolbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13insertToolBarEP8QToolBarS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeToolBar(QToolBar *)
func (this *QMainWindow) RemoveToolBar(toolbar *QToolBar /*777 QToolBar **/) {
	var convArg0 = toolbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13removeToolBarEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeToolBarBreak(QToolBar *)
func (this *QMainWindow) RemoveToolBarBreak(before *QToolBar /*777 QToolBar **/) {
	var convArg0 = before.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow18removeToolBarBreakEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:159
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unifiedTitleAndToolBarOnMac()
func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:161
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ToolBarArea toolBarArea(QToolBar *)
func (this *QMainWindow) ToolBarArea(toolbar *QToolBar /*777 QToolBar **/) int {
	var convArg0 = toolbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow11toolBarAreaEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toolBarBreak(QToolBar *)
func (this *QMainWindow) ToolBarBreak(toolbar *QToolBar /*777 QToolBar **/) bool {
	var convArg0 = toolbar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow12toolBarBreakEP8QToolBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addDockWidget(Qt::DockWidgetArea, QDockWidget *)
func (this *QMainWindow) AddDockWidget(area int, dockwidget *QDockWidget /*777 QDockWidget **/) {
	var convArg1 = dockwidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:166
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addDockWidget(Qt::DockWidgetArea, QDockWidget *, Qt::Orientation)
func (this *QMainWindow) AddDockWidget_1(area int, dockwidget *QDockWidget /*777 QDockWidget **/, orientation int) {
	var convArg1 = dockwidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidgetNS0_11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1, orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void splitDockWidget(QDockWidget *, QDockWidget *, Qt::Orientation)
func (this *QMainWindow) SplitDockWidget(after *QDockWidget /*777 QDockWidget **/, dockwidget *QDockWidget /*777 QDockWidget **/, orientation int) {
	var convArg0 = after.GetCthis()
	var convArg1 = dockwidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15splitDockWidgetEP11QDockWidgetS1_N2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabifyDockWidget(QDockWidget *, QDockWidget *)
func (this *QMainWindow) TabifyDockWidget(first *QDockWidget /*777 QDockWidget **/, second *QDockWidget /*777 QDockWidget **/) {
	var convArg0 = first.GetCthis()
	var convArg1 = second.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeDockWidget(QDockWidget *)
func (this *QMainWindow) RemoveDockWidget(dockwidget *QDockWidget /*777 QDockWidget **/) {
	var convArg0 = dockwidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16removeDockWidgetEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreDockWidget(QDockWidget *)
func (this *QMainWindow) RestoreDockWidget(dockwidget *QDockWidget /*777 QDockWidget **/) bool {
	var convArg0 = dockwidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DockWidgetArea dockWidgetArea(QDockWidget *)
func (this *QMainWindow) DockWidgetArea(dockwidget *QDockWidget /*777 QDockWidget **/) int {
	var convArg0 = dockwidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState(int)
func (this *QMainWindow) SaveState(version int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMainWindow9saveStateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), version)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &, int)
func (this *QMainWindow) RestoreState(state *qtcore.QByteArray, version int) bool {
	var convArg0 = state.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow12restoreStateERK10QByteArrayi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, version)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMenu * createPopupMenu()
func (this *QMainWindow) CreatePopupMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15createPopupMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(_Bool)
func (this *QMainWindow) SetAnimated(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDockNestingEnabled(_Bool)
func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow21setDockNestingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnifiedTitleAndToolBarOnMac(_Bool)
func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconSizeChanged(const QSize &)
func (this *QMainWindow) IconSizeChanged(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow15iconSizeChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabifiedDockWidgetActivated(QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget *QDockWidget /*777 QDockWidget **/) {
	var convArg0 = dockWidget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow27tabifiedDockWidgetActivatedEP11QDockWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QMainWindow) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QMainWindow) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMainWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QMainWindow__DockOption = int

const QMainWindow__AnimatedDocks QMainWindow__DockOption = 1
const QMainWindow__AllowNestedDocks QMainWindow__DockOption = 2
const QMainWindow__AllowTabbedDocks QMainWindow__DockOption = 4
const QMainWindow__ForceTabbedDocks QMainWindow__DockOption = 8
const QMainWindow__VerticalTabs QMainWindow__DockOption = 16
const QMainWindow__GroupedDragging QMainWindow__DockOption = 32

//  body block end
