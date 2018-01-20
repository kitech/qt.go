//  header block begin
// /usr/include/qt/QtWidgets/qmdiarea.h
// #include <qmdiarea.h>
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
type QMdiArea struct {
	*QAbstractScrollArea
}

func (this *QMdiArea) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}

// /usr/include/qt/QtWidgets/qmdiarea.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMdiArea) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:90
// index:0
// void QMdiArea(class QWidget *)
func NewQMdiArea(parent unsafe.Pointer) *QMdiArea {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMdiAreaFromPointer(cthis)
	return gothis
}
func NewQMdiAreaFromPointer(cthis unsafe.Pointer) *QMdiArea {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QMdiArea{bcthis0}
}

// /usr/include/qt/QtWidgets/qmdiarea.h:91
// index:0
// virtual
// void ~QMdiArea()
func DeleteQMdiArea(*QMdiArea) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiAreaD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:93
// index:0
// virtual
// QSize sizeHint()
func (this *QMdiArea) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:94
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QMdiArea) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:96
// index:0
// QMdiSubWindow * currentSubWindow()
func (this *QMdiArea) CurrentSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea16currentSubWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:97
// index:0
// QMdiSubWindow * activeSubWindow()
func (this *QMdiArea) ActiveSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15activeSubWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:98
// index:0
// QList<QMdiSubWindow *> subWindowList(enum QMdiArea::WindowOrder)
func (this *QMdiArea) SubWindowList(order int) {
	// 0: (, order QMdiArea::WindowOrder), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea13subWindowListENS_11WindowOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:100
// index:0
// QMdiSubWindow * addSubWindow(class QWidget *, Qt::WindowFlags)
func (this *QMdiArea) AddSubWindow(widget unsafe.Pointer, flags int) {
	// 0: (, widget QWidget *, QFlags<Qt::WindowType> flags), (widget, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea12addSubWindowEP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:101
// index:0
// void removeSubWindow(class QWidget *)
func (this *QMdiArea) RemoveSubWindow(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15removeSubWindowEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:103
// index:0
// QBrush background()
func (this *QMdiArea) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10backgroundEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:104
// index:0
// void setBackground(const class QBrush &)
func (this *QMdiArea) SetBackground(background unsafe.Pointer) {
	// 0: (, background const QBrush &), (background)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), background)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:106
// index:0
// QMdiArea::WindowOrder activationOrder()
func (this *QMdiArea) ActivationOrder() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15activationOrderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:107
// index:0
// void setActivationOrder(enum QMdiArea::WindowOrder)
func (this *QMdiArea) SetActivationOrder(order int) {
	// 0: (, order QMdiArea::WindowOrder), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18setActivationOrderENS_11WindowOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:109
// index:0
// void setOption(enum QMdiArea::AreaOption, _Bool)
func (this *QMdiArea) SetOption(option int, on bool) {
	// 0: (, option QMdiArea::AreaOption, on bool), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea9setOptionENS_10AreaOptionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:110
// index:0
// bool testOption(enum QMdiArea::AreaOption)
func (this *QMdiArea) TestOption(opton int) {
	// 0: (, opton QMdiArea::AreaOption), (&opton)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10testOptionENS_10AreaOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &opton)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:112
// index:0
// void setViewMode(enum QMdiArea::ViewMode)
func (this *QMdiArea) SetViewMode(mode int) {
	// 0: (, mode QMdiArea::ViewMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:113
// index:0
// QMdiArea::ViewMode viewMode()
func (this *QMdiArea) ViewMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8viewModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:116
// index:0
// bool documentMode()
func (this *QMdiArea) DocumentMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea12documentModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:117
// index:0
// void setDocumentMode(_Bool)
func (this *QMdiArea) SetDocumentMode(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15setDocumentModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:119
// index:0
// void setTabsClosable(_Bool)
func (this *QMdiArea) SetTabsClosable(closable bool) {
	// 0: (, closable bool), (&closable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15setTabsClosableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &closable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:120
// index:0
// bool tabsClosable()
func (this *QMdiArea) TabsClosable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea12tabsClosableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:122
// index:0
// void setTabsMovable(_Bool)
func (this *QMdiArea) SetTabsMovable(movable bool) {
	// 0: (, movable bool), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14setTabsMovableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:123
// index:0
// bool tabsMovable()
func (this *QMdiArea) TabsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea11tabsMovableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:126
// index:0
// void setTabShape(class QTabWidget::TabShape)
func (this *QMdiArea) SetTabShape(shape int) {
	// 0: (, shape QTabWidget::TabShape), (&shape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11setTabShapeEN10QTabWidget8TabShapeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:127
// index:0
// QTabWidget::TabShape tabShape()
func (this *QMdiArea) TabShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8tabShapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:129
// index:0
// void setTabPosition(class QTabWidget::TabPosition)
func (this *QMdiArea) SetTabPosition(position int) {
	// 0: (, position QTabWidget::TabPosition), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14setTabPositionEN10QTabWidget11TabPositionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:130
// index:0
// QTabWidget::TabPosition tabPosition()
func (this *QMdiArea) TabPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea11tabPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:134
// index:0
// void subWindowActivated(class QMdiSubWindow *)
func (this *QMdiArea) SubWindowActivated(arg0 unsafe.Pointer) {
	// 0: (, QMdiSubWindow * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18subWindowActivatedEP13QMdiSubWindow", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:137
// index:0
// void setActiveSubWindow(class QMdiSubWindow *)
func (this *QMdiArea) SetActiveSubWindow(window unsafe.Pointer) {
	// 0: (, window QMdiSubWindow *), (window)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow", ffiqt.FFI_TYPE_VOID, this.GetCthis(), window)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:138
// index:0
// void tileSubWindows()
func (this *QMdiArea) TileSubWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14tileSubWindowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:139
// index:0
// void cascadeSubWindows()
func (this *QMdiArea) CascadeSubWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea17cascadeSubWindowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:140
// index:0
// void closeActiveSubWindow()
func (this *QMdiArea) CloseActiveSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea20closeActiveSubWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:141
// index:0
// void closeAllSubWindows()
func (this *QMdiArea) CloseAllSubWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18closeAllSubWindowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:142
// index:0
// void activateNextSubWindow()
func (this *QMdiArea) ActivateNextSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea21activateNextSubWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:143
// index:0
// void activatePreviousSubWindow()
func (this *QMdiArea) ActivatePreviousSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea25activatePreviousSubWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:146
// index:0
// virtual
// void setupViewport(class QWidget *)
func (this *QMdiArea) SetupViewport(viewport unsafe.Pointer) {
	// 0: (, viewport QWidget *), (viewport)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13setupViewportEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), viewport)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:149
// index:0
// virtual
// bool event(class QEvent *)
func (this *QMdiArea) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:150
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QMdiArea) EventFilter(object unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, object QObject *, event QEvent *), (object, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:151
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QMdiArea) PaintEvent(paintEvent unsafe.Pointer) {
	// 0: (, paintEvent QPaintEvent *), (paintEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), paintEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:152
// index:0
// virtual
// void childEvent(class QChildEvent *)
func (this *QMdiArea) ChildEvent(childEvent unsafe.Pointer) {
	// 0: (, childEvent QChildEvent *), (childEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea10childEventEP11QChildEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), childEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:153
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMdiArea) ResizeEvent(resizeEvent unsafe.Pointer) {
	// 0: (, resizeEvent QResizeEvent *), (resizeEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), resizeEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:154
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QMdiArea) TimerEvent(timerEvent unsafe.Pointer) {
	// 0: (, timerEvent QTimerEvent *), (timerEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), timerEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:155
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QMdiArea) ShowEvent(showEvent unsafe.Pointer) {
	// 0: (, showEvent QShowEvent *), (showEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), showEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:156
// index:0
// virtual
// bool viewportEvent(class QEvent *)
func (this *QMdiArea) ViewportEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13viewportEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:157
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QMdiArea) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

//  body block end
