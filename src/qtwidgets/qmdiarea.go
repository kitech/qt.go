//  header block begin
// /usr/include/qt/QtWidgets/qmdiarea.h
// #include <qmdiarea.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
type QMdiArea struct {
	*QAbstractScrollArea
}

func (this *QMdiArea) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}
func NewQMdiAreaFromPointer(cthis unsafe.Pointer) *QMdiArea {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QMdiArea{bcthis0}
}

// /usr/include/qt/QtWidgets/qmdiarea.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMdiArea) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:90
// index:0
// Public
// void QMdiArea(class QWidget *)
func NewQMdiArea(parent unsafe.Pointer) *QMdiArea {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMdiAreaFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmdiarea.h:91
// index:0
// Public virtual
// void ~QMdiArea()
func DeleteQMdiArea(*QMdiArea) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiAreaD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:93
// index:0
// Public virtual
// QSize sizeHint()
func (this *QMdiArea) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:94
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QMdiArea) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:96
// index:0
// Public
// QMdiSubWindow * currentSubWindow()
func (this *QMdiArea) CurrentSubWindow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea16currentSubWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:97
// index:0
// Public
// QMdiSubWindow * activeSubWindow()
func (this *QMdiArea) ActiveSubWindow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15activeSubWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:98
// index:0
// Public
// QList<QMdiSubWindow *> subWindowList(enum QMdiArea::WindowOrder)
func (this *QMdiArea) SubWindowList(order int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea13subWindowListENS_11WindowOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:101
// index:0
// Public
// void removeSubWindow(class QWidget *)
func (this *QMdiArea) RemoveSubWindow(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15removeSubWindowEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:103
// index:0
// Public
// QBrush background()
func (this *QMdiArea) Background() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:104
// index:0
// Public
// void setBackground(const class QBrush &)
func (this *QMdiArea) SetBackground(background *qtgui.QBrush) {
	var convArg0 = background.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:106
// index:0
// Public
// QMdiArea::WindowOrder activationOrder()
func (this *QMdiArea) ActivationOrder() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15activationOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:107
// index:0
// Public
// void setActivationOrder(enum QMdiArea::WindowOrder)
func (this *QMdiArea) SetActivationOrder(order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18setActivationOrderENS_11WindowOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:109
// index:0
// Public
// void setOption(enum QMdiArea::AreaOption, _Bool)
func (this *QMdiArea) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea9setOptionENS_10AreaOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:110
// index:0
// Public
// bool testOption(enum QMdiArea::AreaOption)
func (this *QMdiArea) TestOption(opton int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10testOptionENS_10AreaOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &opton)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:112
// index:0
// Public
// void setViewMode(enum QMdiArea::ViewMode)
func (this *QMdiArea) SetViewMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:113
// index:0
// Public
// QMdiArea::ViewMode viewMode()
func (this *QMdiArea) ViewMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8viewModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:116
// index:0
// Public
// bool documentMode()
func (this *QMdiArea) DocumentMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea12documentModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:117
// index:0
// Public
// void setDocumentMode(_Bool)
func (this *QMdiArea) SetDocumentMode(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15setDocumentModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:119
// index:0
// Public
// void setTabsClosable(_Bool)
func (this *QMdiArea) SetTabsClosable(closable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15setTabsClosableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &closable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:120
// index:0
// Public
// bool tabsClosable()
func (this *QMdiArea) TabsClosable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea12tabsClosableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:122
// index:0
// Public
// void setTabsMovable(_Bool)
func (this *QMdiArea) SetTabsMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14setTabsMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:123
// index:0
// Public
// bool tabsMovable()
func (this *QMdiArea) TabsMovable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea11tabsMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:126
// index:0
// Public
// void setTabShape(class QTabWidget::TabShape)
func (this *QMdiArea) SetTabShape(shape int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11setTabShapeEN10QTabWidget8TabShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:127
// index:0
// Public
// QTabWidget::TabShape tabShape()
func (this *QMdiArea) TabShape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8tabShapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:129
// index:0
// Public
// void setTabPosition(class QTabWidget::TabPosition)
func (this *QMdiArea) SetTabPosition(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14setTabPositionEN10QTabWidget11TabPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:130
// index:0
// Public
// QTabWidget::TabPosition tabPosition()
func (this *QMdiArea) TabPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea11tabPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:134
// index:0
// Public
// void subWindowActivated(class QMdiSubWindow *)
func (this *QMdiArea) SubWindowActivated(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18subWindowActivatedEP13QMdiSubWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:137
// index:0
// Public
// void setActiveSubWindow(class QMdiSubWindow *)
func (this *QMdiArea) SetActiveSubWindow(window unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), window)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:138
// index:0
// Public
// void tileSubWindows()
func (this *QMdiArea) TileSubWindows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14tileSubWindowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:139
// index:0
// Public
// void cascadeSubWindows()
func (this *QMdiArea) CascadeSubWindows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea17cascadeSubWindowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:140
// index:0
// Public
// void closeActiveSubWindow()
func (this *QMdiArea) CloseActiveSubWindow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea20closeActiveSubWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:141
// index:0
// Public
// void closeAllSubWindows()
func (this *QMdiArea) CloseAllSubWindows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18closeAllSubWindowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:142
// index:0
// Public
// void activateNextSubWindow()
func (this *QMdiArea) ActivateNextSubWindow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea21activateNextSubWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:143
// index:0
// Public
// void activatePreviousSubWindow()
func (this *QMdiArea) ActivatePreviousSubWindow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea25activatePreviousSubWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:146
// index:0
// Protected virtual
// void setupViewport(class QWidget *)
func (this *QMdiArea) SetupViewport(viewport unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13setupViewportEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), viewport)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:149
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QMdiArea) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:150
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QMdiArea) EventFilter(object unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:151
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QMdiArea) PaintEvent(paintEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), paintEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:152
// index:0
// Protected virtual
// void childEvent(class QChildEvent *)
func (this *QMdiArea) ChildEvent(childEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), childEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:153
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMdiArea) ResizeEvent(resizeEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resizeEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:154
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QMdiArea) TimerEvent(timerEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timerEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:155
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QMdiArea) ShowEvent(showEvent unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), showEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:156
// index:0
// Protected virtual
// bool viewportEvent(class QEvent *)
func (this *QMdiArea) ViewportEvent(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmdiarea.h:157
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QMdiArea) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

//  body block end
