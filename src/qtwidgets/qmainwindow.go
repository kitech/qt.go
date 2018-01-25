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
// Public virtual
// const QMetaObject * metaObject()
func (this *QMainWindow) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:94
// index:0
// Public
// void QMainWindow(class QWidget *, Qt::WindowFlags)
func NewQMainWindow(parent *QWidget /*444 QWidget **/, flags int) *QMainWindow {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMainWindowFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmainwindow.h:95
// index:0
// Public virtual
// void ~QMainWindow()
func DeleteQMainWindow(*QMainWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:97
// index:0
// Public
// QSize iconSize()
func (this *QMainWindow) IconSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow8iconSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:98
// index:0
// Public
// void setIconSize(const class QSize &)
func (this *QMainWindow) SetIconSize(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:100
// index:0
// Public
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QMainWindow) ToolButtonStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow15toolButtonStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:101
// index:0
// Public
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:104
// index:0
// Public
// bool isAnimated()
func (this *QMainWindow) IsAnimated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10isAnimatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:105
// index:0
// Public
// bool isDockNestingEnabled()
func (this *QMainWindow) IsDockNestingEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow20isDockNestingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:109
// index:0
// Public
// bool documentMode()
func (this *QMainWindow) DocumentMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow12documentModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:110
// index:0
// Public
// void setDocumentMode(_Bool)
func (this *QMainWindow) SetDocumentMode(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15setDocumentModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:114
// index:0
// Public
// QTabWidget::TabShape tabShape()
func (this *QMainWindow) TabShape() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow8tabShapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:115
// index:0
// Public
// void setTabShape(class QTabWidget::TabShape)
func (this *QMainWindow) SetTabShape(tabShape int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setTabShapeEN10QTabWidget8TabShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), tabShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:116
// index:0
// Public
// QTabWidget::TabPosition tabPosition(Qt::DockWidgetArea)
func (this *QMainWindow) TabPosition(area int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11tabPositionEN2Qt14DockWidgetAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:121
// index:0
// Public
// QMainWindow::DockOptions dockOptions()
func (this *QMainWindow) DockOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11dockOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:123
// index:0
// Public
// bool isSeparator(const class QPoint &)
func (this *QMainWindow) IsSeparator(pos *qtcore.QPoint) bool {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11isSeparatorERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:126
// index:0
// Public
// QMenuBar * menuBar()
func (this *QMainWindow) MenuBar() *QMenuBar /*444 QMenuBar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow7menuBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:127
// index:0
// Public
// void setMenuBar(class QMenuBar *)
func (this *QMainWindow) SetMenuBar(menubar *QMenuBar /*444 QMenuBar **/) {
	var convArg0 = menubar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10setMenuBarEP8QMenuBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:129
// index:0
// Public
// QWidget * menuWidget()
func (this *QMainWindow) MenuWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow10menuWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:130
// index:0
// Public
// void setMenuWidget(class QWidget *)
func (this *QMainWindow) SetMenuWidget(menubar *QWidget /*444 QWidget **/) {
	var convArg0 = menubar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13setMenuWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:134
// index:0
// Public
// QStatusBar * statusBar()
func (this *QMainWindow) StatusBar() *QStatusBar /*444 QStatusBar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow9statusBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStatusBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:135
// index:0
// Public
// void setStatusBar(class QStatusBar *)
func (this *QMainWindow) SetStatusBar(statusbar *QStatusBar /*444 QStatusBar **/) {
	var convArg0 = statusbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow12setStatusBarEP10QStatusBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:138
// index:0
// Public
// QWidget * centralWidget()
func (this *QMainWindow) CentralWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow13centralWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:139
// index:0
// Public
// void setCentralWidget(class QWidget *)
func (this *QMainWindow) SetCentralWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16setCentralWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:141
// index:0
// Public
// QWidget * takeCentralWidget()
func (this *QMainWindow) TakeCentralWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow17takeCentralWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:144
// index:0
// Public
// void setCorner(Qt::Corner, Qt::DockWidgetArea)
func (this *QMainWindow) SetCorner(corner int, area int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow9setCornerEN2Qt6CornerENS0_14DockWidgetAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), corner, area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:145
// index:0
// Public
// Qt::DockWidgetArea corner(Qt::Corner)
func (this *QMainWindow) Corner(corner int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow6cornerEN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:149
// index:0
// Public
// void addToolBarBreak(Qt::ToolBarArea)
func (this *QMainWindow) AddToolBarBreak(area int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15addToolBarBreakEN2Qt11ToolBarAreaE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:150
// index:0
// Public
// void insertToolBarBreak(class QToolBar *)
func (this *QMainWindow) InsertToolBarBreak(before *QToolBar /*444 QToolBar **/) {
	var convArg0 = before.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18insertToolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:152
// index:0
// Public
// void addToolBar(Qt::ToolBarArea, class QToolBar *)
func (this *QMainWindow) AddToolBar(area int, toolbar *QToolBar /*444 QToolBar **/) {
	var convArg1 = toolbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEN2Qt11ToolBarAreaEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:153
// index:1
// Public
// void addToolBar(class QToolBar *)
func (this *QMainWindow) AddToolBar_1(toolbar *QToolBar /*444 QToolBar **/) {
	var convArg0 = toolbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:154
// index:2
// Public
// QToolBar * addToolBar(const class QString &)
func (this *QMainWindow) AddToolBar_2(title *qtcore.QString) *QToolBar /*444 QToolBar **/ {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow10addToolBarERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQToolBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:155
// index:0
// Public
// void insertToolBar(class QToolBar *, class QToolBar *)
func (this *QMainWindow) InsertToolBar(before *QToolBar /*444 QToolBar **/, toolbar *QToolBar /*444 QToolBar **/) {
	var convArg0 = before.GetCthis()
	var convArg1 = toolbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13insertToolBarEP8QToolBarS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:156
// index:0
// Public
// void removeToolBar(class QToolBar *)
func (this *QMainWindow) RemoveToolBar(toolbar *QToolBar /*444 QToolBar **/) {
	var convArg0 = toolbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13removeToolBarEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:157
// index:0
// Public
// void removeToolBarBreak(class QToolBar *)
func (this *QMainWindow) RemoveToolBarBreak(before *QToolBar /*444 QToolBar **/) {
	var convArg0 = before.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow18removeToolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:159
// index:0
// Public
// bool unifiedTitleAndToolBarOnMac()
func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow27unifiedTitleAndToolBarOnMacEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:161
// index:0
// Public
// Qt::ToolBarArea toolBarArea(class QToolBar *)
func (this *QMainWindow) ToolBarArea(toolbar *QToolBar /*444 QToolBar **/) int {
	var convArg0 = toolbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow11toolBarAreaEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:162
// index:0
// Public
// bool toolBarBreak(class QToolBar *)
func (this *QMainWindow) ToolBarBreak(toolbar *QToolBar /*444 QToolBar **/) bool {
	var convArg0 = toolbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow12toolBarBreakEP8QToolBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:165
// index:0
// Public
// void addDockWidget(Qt::DockWidgetArea, class QDockWidget *)
func (this *QMainWindow) AddDockWidget(area int, dockwidget *QDockWidget /*444 QDockWidget **/) {
	var convArg1 = dockwidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:166
// index:1
// Public
// void addDockWidget(Qt::DockWidgetArea, class QDockWidget *, Qt::Orientation)
func (this *QMainWindow) AddDockWidget_1(area int, dockwidget *QDockWidget /*444 QDockWidget **/, orientation int) {
	var convArg1 = dockwidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow13addDockWidgetEN2Qt14DockWidgetAreaEP11QDockWidgetNS0_11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), area, convArg1, orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:168
// index:0
// Public
// void splitDockWidget(class QDockWidget *, class QDockWidget *, Qt::Orientation)
func (this *QMainWindow) SplitDockWidget(after *QDockWidget /*444 QDockWidget **/, dockwidget *QDockWidget /*444 QDockWidget **/, orientation int) {
	var convArg0 = after.GetCthis()
	var convArg1 = dockwidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15splitDockWidgetEP11QDockWidgetS1_N2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:170
// index:0
// Public
// void tabifyDockWidget(class QDockWidget *, class QDockWidget *)
func (this *QMainWindow) TabifyDockWidget(first *QDockWidget /*444 QDockWidget **/, second *QDockWidget /*444 QDockWidget **/) {
	var convArg0 = first.GetCthis()
	var convArg1 = second.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16tabifyDockWidgetEP11QDockWidgetS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:172
// index:0
// Public
// void removeDockWidget(class QDockWidget *)
func (this *QMainWindow) RemoveDockWidget(dockwidget *QDockWidget /*444 QDockWidget **/) {
	var convArg0 = dockwidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16removeDockWidgetEP11QDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:173
// index:0
// Public
// bool restoreDockWidget(class QDockWidget *)
func (this *QMainWindow) RestoreDockWidget(dockwidget *QDockWidget /*444 QDockWidget **/) bool {
	var convArg0 = dockwidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow17restoreDockWidgetEP11QDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:175
// index:0
// Public
// Qt::DockWidgetArea dockWidgetArea(class QDockWidget *)
func (this *QMainWindow) DockWidgetArea(dockwidget *QDockWidget /*444 QDockWidget **/) int {
	var convArg0 = dockwidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow14dockWidgetAreaEP11QDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:181
// index:0
// Public
// QByteArray saveState(int)
func (this *QMainWindow) SaveState(version int) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMainWindow9saveStateEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), version)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:182
// index:0
// Public
// bool restoreState(const class QByteArray &, int)
func (this *QMainWindow) RestoreState(state *qtcore.QByteArray, version int) bool {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow12restoreStateERK10QByteArrayi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, version)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmainwindow.h:185
// index:0
// Public virtual
// QMenu * createPopupMenu()
func (this *QMainWindow) CreatePopupMenu() *QMenu /*444 QMenu **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15createPopupMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qmainwindow.h:190
// index:0
// Public
// void setAnimated(_Bool)
func (this *QMainWindow) SetAnimated(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow11setAnimatedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:191
// index:0
// Public
// void setDockNestingEnabled(_Bool)
func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow21setDockNestingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:194
// index:0
// Public
// void setUnifiedTitleAndToolBarOnMac(_Bool)
func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow30setUnifiedTitleAndToolBarOnMacEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:198
// index:0
// Public
// void iconSizeChanged(const class QSize &)
func (this *QMainWindow) IconSizeChanged(iconSize *qtcore.QSize) {
	var convArg0 = iconSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:199
// index:0
// Public
// void toolButtonStyleChanged(Qt::ToolButtonStyle)
func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow22toolButtonStyleChangedEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), toolButtonStyle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:201
// index:0
// Public
// void tabifiedDockWidgetActivated(class QDockWidget *)
func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget *QDockWidget /*444 QDockWidget **/) {
	var convArg0 = dockWidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow27tabifiedDockWidgetActivatedEP11QDockWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:206
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QMainWindow) ContextMenuEvent(event *qtgui.QContextMenuEvent /*444 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmainwindow.h:208
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QMainWindow) Event(event *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMainWindow5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
