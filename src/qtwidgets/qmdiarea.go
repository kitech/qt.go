//  header block begin
// /usr/include/qt/QtWidgets/qmdiarea.h
// #include <qmdiarea.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qmdiarea.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMdiArea) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:90
// index:0
// void QMdiArea(class QWidget *)
func NewQMdiArea(parent unsafe.Pointer) *QMdiArea {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QMdiArea{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:94
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QMdiArea) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:96
// index:0
// QMdiSubWindow * currentSubWindow()
func (this *QMdiArea) CurrentSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea16currentSubWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:97
// index:0
// QMdiSubWindow * activeSubWindow()
func (this *QMdiArea) ActiveSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15activeSubWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:98
// index:0
// QList<QMdiSubWindow *> subWindowList(enum QMdiArea::WindowOrder)
func (this *QMdiArea) SubWindowList(order int) {
	// 0: (, QMdiArea::WindowOrder order), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea13subWindowListENS_11WindowOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:101
// index:0
// void removeSubWindow(class QWidget *)
func (this *QMdiArea) RemoveSubWindow(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15removeSubWindowEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:103
// index:0
// QBrush background()
func (this *QMdiArea) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10backgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:104
// index:0
// void setBackground(const class QBrush &)
func (this *QMdiArea) SetBackground(background unsafe.Pointer) {
	// 0: (, const QBrush & background), (background)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, background)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:106
// index:0
// QMdiArea::WindowOrder activationOrder()
func (this *QMdiArea) ActivationOrder() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea15activationOrderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:107
// index:0
// void setActivationOrder(enum QMdiArea::WindowOrder)
func (this *QMdiArea) SetActivationOrder(order int) {
	// 0: (, QMdiArea::WindowOrder order), (&order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18setActivationOrderENS_11WindowOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:109
// index:0
// void setOption(enum QMdiArea::AreaOption, _Bool)
func (this *QMdiArea) SetOption(option int, on bool) {
	// 0: (, QMdiArea::AreaOption option, bool on), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea9setOptionENS_10AreaOptionEb", ffiqt.FFI_TYPE_VOID, this.cthis, &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:110
// index:0
// bool testOption(enum QMdiArea::AreaOption)
func (this *QMdiArea) TestOption(opton int) {
	// 0: (, QMdiArea::AreaOption opton), (&opton)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea10testOptionENS_10AreaOptionE", ffiqt.FFI_TYPE_VOID, this.cthis, &opton)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:112
// index:0
// void setViewMode(enum QMdiArea::ViewMode)
func (this *QMdiArea) SetViewMode(mode int) {
	// 0: (, QMdiArea::ViewMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:113
// index:0
// QMdiArea::ViewMode viewMode()
func (this *QMdiArea) ViewMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8viewModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:116
// index:0
// bool documentMode()
func (this *QMdiArea) DocumentMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea12documentModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:117
// index:0
// void setDocumentMode(_Bool)
func (this *QMdiArea) SetDocumentMode(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15setDocumentModeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:119
// index:0
// void setTabsClosable(_Bool)
func (this *QMdiArea) SetTabsClosable(closable bool) {
	// 0: (, bool closable), (&closable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea15setTabsClosableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &closable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:120
// index:0
// bool tabsClosable()
func (this *QMdiArea) TabsClosable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea12tabsClosableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:122
// index:0
// void setTabsMovable(_Bool)
func (this *QMdiArea) SetTabsMovable(movable bool) {
	// 0: (, bool movable), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14setTabsMovableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:123
// index:0
// bool tabsMovable()
func (this *QMdiArea) TabsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea11tabsMovableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:126
// index:0
// void setTabShape(class QTabWidget::TabShape)
func (this *QMdiArea) SetTabShape(shape int) {
	// 0: (, QTabWidget::TabShape shape), (&shape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea11setTabShapeEN10QTabWidget8TabShapeE", ffiqt.FFI_TYPE_VOID, this.cthis, &shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:127
// index:0
// QTabWidget::TabShape tabShape()
func (this *QMdiArea) TabShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea8tabShapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:129
// index:0
// void setTabPosition(class QTabWidget::TabPosition)
func (this *QMdiArea) SetTabPosition(position int) {
	// 0: (, QTabWidget::TabPosition position), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14setTabPositionEN10QTabWidget11TabPositionE", ffiqt.FFI_TYPE_VOID, this.cthis, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:130
// index:0
// QTabWidget::TabPosition tabPosition()
func (this *QMdiArea) TabPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMdiArea11tabPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:134
// index:0
// void subWindowActivated(class QMdiSubWindow *)
func (this *QMdiArea) SubWindowActivated(arg0 unsafe.Pointer) {
	// 0: (, QMdiSubWindow * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18subWindowActivatedEP13QMdiSubWindow", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:137
// index:0
// void setActiveSubWindow(class QMdiSubWindow *)
func (this *QMdiArea) SetActiveSubWindow(window unsafe.Pointer) {
	// 0: (, QMdiSubWindow * window), (window)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow", ffiqt.FFI_TYPE_VOID, this.cthis, window)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:138
// index:0
// void tileSubWindows()
func (this *QMdiArea) TileSubWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea14tileSubWindowsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:139
// index:0
// void cascadeSubWindows()
func (this *QMdiArea) CascadeSubWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea17cascadeSubWindowsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:140
// index:0
// void closeActiveSubWindow()
func (this *QMdiArea) CloseActiveSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea20closeActiveSubWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:141
// index:0
// void closeAllSubWindows()
func (this *QMdiArea) CloseAllSubWindows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea18closeAllSubWindowsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:142
// index:0
// void activateNextSubWindow()
func (this *QMdiArea) ActivateNextSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea21activateNextSubWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:143
// index:0
// void activatePreviousSubWindow()
func (this *QMdiArea) ActivatePreviousSubWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMdiArea25activatePreviousSubWindowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
