// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qmdiarea.h
// #include <qmdiarea.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
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

// void setupViewport(QWidget *)
func (this *QMdiArea) InheritSetupViewport(f func(viewport *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setupViewport", f)
}

// bool event(QEvent *)
func (this *QMdiArea) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QMdiArea) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void paintEvent(QPaintEvent *)
func (this *QMdiArea) InheritPaintEvent(f func(paintEvent *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void childEvent(QChildEvent *)
func (this *QMdiArea) InheritChildEvent(f func(childEvent *qtcore.QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QMdiArea) InheritResizeEvent(f func(resizeEvent *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QMdiArea) InheritTimerEvent(f func(timerEvent *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QMdiArea) InheritShowEvent(f func(showEvent *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// bool viewportEvent(QEvent *)
func (this *QMdiArea) InheritViewportEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QMdiArea) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

/*

 */
type QMdiArea struct {
	*QAbstractScrollArea
}
type QMdiArea_ITF interface {
	QAbstractScrollArea_ITF
	QMdiArea_PTR() *QMdiArea
}

func (ptr *QMdiArea) QMdiArea_PTR() *QMdiArea { return ptr }

func (this *QMdiArea) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QMdiArea) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQMdiAreaFromPointer(cthis unsafe.Pointer) *QMdiArea {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QMdiArea{bcthis0}
}
func (*QMdiArea) NewFromPointer(cthis unsafe.Pointer) *QMdiArea {
	return NewQMdiAreaFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMdiArea) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdiarea.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMdiArea(QWidget *)

/*
Constructs an empty mdi area. parent is passed to QWidget's constructor.
*/
func (*QMdiArea) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QMdiArea {
	return NewQMdiArea(parent)
}
func NewQMdiArea(parent QWidget_ITF /*777 QWidget **/) *QMdiArea {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiAreaC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMdiAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMdiArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qmdiarea.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMdiArea(QWidget *)

/*
Constructs an empty mdi area. parent is passed to QWidget's constructor.
*/
func (*QMdiArea) NewForInheritp() *QMdiArea {
	return NewQMdiAreap()
}
func NewQMdiAreap() *QMdiArea {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiAreaC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMdiAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMdiArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qmdiarea.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMdiArea()

/*

 */
func DeleteQMdiArea(this *QMdiArea) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiAreaD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QAbstractScrollArea::sizeHint().
*/
func (this *QMdiArea) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdiarea.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QAbstractScrollArea::minimumSizeHint().
*/
func (this *QMdiArea) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdiarea.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QMdiSubWindow * currentSubWindow() const

/*
Returns a pointer to the current subwindow, or 0 if there is no current subwindow.

This function will return the same as activeSubWindow() if the QApplication containing QMdiArea is active.

See also activeSubWindow() and QApplication::activeWindow().
*/
func (this *QMdiArea) CurrentSubWindow() *QMdiSubWindow /*777 QMdiSubWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea16currentSubWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdiarea.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QMdiSubWindow * activeSubWindow() const

/*
Returns a pointer to the current active subwindow. If no window is currently active, 0 is returned.

Subwindows are treated as top-level windows with respect to window state, i.e., if a widget outside the MDI area is the active window, no subwindow will be active. Note that if a widget in the window in which the MDI area lives gains focus, the window will be activated.

See also setActiveSubWindow() and Qt::WindowState.
*/
func (this *QMdiArea) ActiveSubWindow() *QMdiSubWindow /*777 QMdiSubWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea15activeSubWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdiarea.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QMdiSubWindow * addSubWindow(QWidget *, Qt::WindowFlags)

/*
Adds widget as a new subwindow to the MDI area. If windowFlags are non-zero, they will override the flags set on the widget.

The widget can be either a QMdiSubWindow or another QWidget (in which case the MDI area will create a subwindow and set the widget as the internal widget).

Note: Once the subwindow has been added, its parent will be the viewport widget of the QMdiArea.


      QMdiArea mdiArea;
      QMdiSubWindow *subWindow1 = new QMdiSubWindow;
      subWindow1->setWidget(internalWidget1);
      subWindow1->setAttribute(Qt::WA_DeleteOnClose);
      mdiArea.addSubWindow(subWindow1);

      QMdiSubWindow *subWindow2 =
          mdiArea.addSubWindow(internalWidget2);



When you create your own subwindow, you must set the Qt::WA_DeleteOnClose widget attribute if you want the window to be deleted when closed in the MDI area. If not, the window will be hidden and the MDI area will not activate the next subwindow.

Returns the QMdiSubWindow that is added to the MDI area.

See also removeSubWindow().
*/
func (this *QMdiArea) AddSubWindow(widget QWidget_ITF /*777 QWidget **/, flags int) *QMdiSubWindow /*777 QMdiSubWindow **/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea12addSubWindowEP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdiarea.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QMdiSubWindow * addSubWindow(QWidget *, Qt::WindowFlags)

/*
Adds widget as a new subwindow to the MDI area. If windowFlags are non-zero, they will override the flags set on the widget.

The widget can be either a QMdiSubWindow or another QWidget (in which case the MDI area will create a subwindow and set the widget as the internal widget).

Note: Once the subwindow has been added, its parent will be the viewport widget of the QMdiArea.


      QMdiArea mdiArea;
      QMdiSubWindow *subWindow1 = new QMdiSubWindow;
      subWindow1->setWidget(internalWidget1);
      subWindow1->setAttribute(Qt::WA_DeleteOnClose);
      mdiArea.addSubWindow(subWindow1);

      QMdiSubWindow *subWindow2 =
          mdiArea.addSubWindow(internalWidget2);



When you create your own subwindow, you must set the Qt::WA_DeleteOnClose widget attribute if you want the window to be deleted when closed in the MDI area. If not, the window will be hidden and the MDI area will not activate the next subwindow.

Returns the QMdiSubWindow that is added to the MDI area.

See also removeSubWindow().
*/
func (this *QMdiArea) AddSubWindowp(widget QWidget_ITF /*777 QWidget **/) *QMdiSubWindow /*777 QMdiSubWindow **/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea12addSubWindowEP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMdiSubWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmdiarea.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeSubWindow(QWidget *)

/*
Removes widget from the MDI area. The widget must be either a QMdiSubWindow or a widget that is the internal widget of a subwindow. Note widget is never actually deleted by QMdiArea. If a QMdiSubWindow is passed in its parent is set to 0 and it is removed, but if an internal widget is passed in the child widget is set to 0 but the QMdiSubWindow is not removed.

See also addSubWindow().
*/
func (this *QMdiArea) RemoveSubWindow(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea15removeSubWindowEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush background() const

/*

 */
func (this *QMdiArea) Background() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qmdiarea.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)

/*

 */
func (this *QMdiArea) SetBackground(background qtgui.QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if background != nil && background.QBrush_PTR() != nil {
		convArg0 = background.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QMdiArea::WindowOrder activationOrder() const

/*

 */
func (this *QMdiArea) ActivationOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea15activationOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActivationOrder(QMdiArea::WindowOrder)

/*

 */
func (this *QMdiArea) SetActivationOrder(order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea18setActivationOrderENS_11WindowOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QMdiArea::AreaOption, bool)

/*
If on is true, option is enabled on the MDI area; otherwise it is disabled. See AreaOption for the effect of each option.

See also AreaOption and testOption().
*/
func (this *QMdiArea) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea9setOptionENS_10AreaOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QMdiArea::AreaOption, bool)

/*
If on is true, option is enabled on the MDI area; otherwise it is disabled. See AreaOption for the effect of each option.

See also AreaOption and testOption().
*/
func (this *QMdiArea) SetOptionp(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea9setOptionENS_10AreaOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QMdiArea::AreaOption) const

/*
Returns true if option is enabled; otherwise returns false.

See also AreaOption and setOption().
*/
func (this *QMdiArea) TestOption(opton int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea10testOptionENS_10AreaOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opton)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewMode(QMdiArea::ViewMode)

/*

 */
func (this *QMdiArea) SetViewMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea11setViewModeENS_8ViewModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] QMdiArea::ViewMode viewMode() const

/*

 */
func (this *QMdiArea) ViewMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea8viewModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode() const

/*

 */
func (this *QMdiArea) DocumentMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea12documentModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(bool)

/*

 */
func (this *QMdiArea) SetDocumentMode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea15setDocumentModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(bool)

/*

 */
func (this *QMdiArea) SetTabsClosable(closable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea15setTabsClosableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), closable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable() const

/*

 */
func (this *QMdiArea) TabsClosable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea12tabsClosableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsMovable(bool)

/*

 */
func (this *QMdiArea) SetTabsMovable(movable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea14setTabsMovableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsMovable() const

/*

 */
func (this *QMdiArea) TabsMovable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea11tabsMovableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(QTabWidget::TabShape)

/*

 */
func (this *QMdiArea) SetTabShape(shape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea11setTabShapeEN10QTabWidget8TabShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shape)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape() const

/*

 */
func (this *QMdiArea) TabShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea8tabShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(QTabWidget::TabPosition)

/*

 */
func (this *QMdiArea) SetTabPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea14setTabPositionEN10QTabWidget11TabPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition() const

/*

 */
func (this *QMdiArea) TabPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMdiArea11tabPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void subWindowActivated(QMdiSubWindow *)

/*
QMdiArea emits this signal after window has been activated. When window is 0, QMdiArea has just deactivated its last active window, and there are no active windows on the workspace.

See also QMdiArea::activeSubWindow().
*/
func (this *QMdiArea) SubWindowActivated(arg0 QMdiSubWindow_ITF /*777 QMdiSubWindow **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMdiSubWindow_PTR() != nil {
		convArg0 = arg0.QMdiSubWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea18subWindowActivatedEP13QMdiSubWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveSubWindow(QMdiSubWindow *)

/*
Activates the subwindow window. If window is 0, any current active window is deactivated.

See also activeSubWindow().
*/
func (this *QMdiArea) SetActiveSubWindow(window QMdiSubWindow_ITF /*777 QMdiSubWindow **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QMdiSubWindow_PTR() != nil {
		convArg0 = window.QMdiSubWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tileSubWindows()

/*
Arranges all child windows in a tile pattern.

See also cascadeSubWindows().
*/
func (this *QMdiArea) TileSubWindows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea14tileSubWindowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cascadeSubWindows()

/*
Arranges all the child windows in a cascade pattern.

See also tileSubWindows().
*/
func (this *QMdiArea) CascadeSubWindows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea17cascadeSubWindowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeActiveSubWindow()

/*
Closes the active subwindow.

See also closeAllSubWindows().
*/
func (this *QMdiArea) CloseActiveSubWindow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea20closeActiveSubWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeAllSubWindows()

/*
Closes all subwindows by sending a QCloseEvent to each window. You may receive subWindowActivated() signals from subwindows before they are closed (if the MDI area activates the subwindow when another is closing).

Subwindows that ignore the close event will remain open.

See also closeActiveSubWindow().
*/
func (this *QMdiArea) CloseAllSubWindows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea18closeAllSubWindowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activateNextSubWindow()

/*
Gives the keyboard focus to another window in the list of child windows. The window activated will be the next one determined by the current activation order.

See also activatePreviousSubWindow() and QMdiArea::WindowOrder.
*/
func (this *QMdiArea) ActivateNextSubWindow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea21activateNextSubWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activatePreviousSubWindow()

/*
Gives the keyboard focus to another window in the list of child windows. The window activated will be the previous one determined by the current activation order.

See also activateNextSubWindow() and QMdiArea::WindowOrder.
*/
func (this *QMdiArea) ActivatePreviousSubWindow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea25activatePreviousSubWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:146
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setupViewport(QWidget *)

/*
Reimplemented from QAbstractScrollArea::setupViewport().

This slot is called by QAbstractScrollArea after setViewport() has been called. Reimplement this function in a subclass of QMdiArea to initialize the new viewport before it is used.

See also setViewport().
*/
func (this *QMdiArea) SetupViewport(viewport QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if viewport != nil && viewport.QWidget_PTR() != nil {
		convArg0 = viewport.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea13setupViewportEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QAbstractScrollArea::event().
*/
func (this *QMdiArea) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*

 */
func (this *QMdiArea) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QAbstractScrollArea::paintEvent().
*/
func (this *QMdiArea) PaintEvent(paintEvent qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if paintEvent != nil && paintEvent.QPaintEvent_PTR() != nil {
		convArg0 = paintEvent.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)

/*
Reimplemented from QObject::childEvent().
*/
func (this *QMdiArea) ChildEvent(childEvent qtcore.QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if childEvent != nil && childEvent.QChildEvent_PTR() != nil {
		convArg0 = childEvent.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:153
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QAbstractScrollArea::resizeEvent().
*/
func (this *QMdiArea) ResizeEvent(resizeEvent qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if resizeEvent != nil && resizeEvent.QResizeEvent_PTR() != nil {
		convArg0 = resizeEvent.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QMdiArea) TimerEvent(timerEvent qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if timerEvent != nil && timerEvent.QTimerEvent_PTR() != nil {
		convArg0 = timerEvent.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QMdiArea) ShowEvent(showEvent qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if showEvent != nil && showEvent.QShowEvent_PTR() != nil {
		convArg0 = showEvent.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdiarea.h:156
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)

/*
Reimplemented from QAbstractScrollArea::viewportEvent().
*/
func (this *QMdiArea) ViewportEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qmdiarea.h:157
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*
Reimplemented from QAbstractScrollArea::scrollContentsBy().
*/
func (this *QMdiArea) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMdiArea16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QMdiArea__AreaOption = int

//
const QMdiArea__DontMaximizeSubWindowOnActivation QMdiArea__AreaOption = 1

func (this *QMdiArea) AreaOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMdiArea_AreaOptionItemName(val int) string {
	var nilthis *QMdiArea
	return nilthis.AreaOptionItemName(val)
}

/*
Specifies the criteria to use for ordering the list of child windows returned by subWindowList(). The functions cascadeSubWindows() and tileSubWindows() follow this order when arranging the windows.



See also subWindowList().

*/
type QMdiArea__WindowOrder = int

// The windows are returned in the order of their creation.
const QMdiArea__CreationOrder QMdiArea__WindowOrder = 0

// The windows are returned in the order in which they are stacked, with the top-most window being last in the list.
const QMdiArea__StackingOrder QMdiArea__WindowOrder = 1

// The windows are returned in the order in which they were activated.
const QMdiArea__ActivationHistoryOrder QMdiArea__WindowOrder = 2

func (this *QMdiArea) WindowOrderItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMdiArea_WindowOrderItemName(val int) string {
	var nilthis *QMdiArea
	return nilthis.WindowOrderItemName(val)
}

/*
This enum describes the view mode of the area; i.e. how sub-windows will be displayed.



This enum was introduced or modified in  Qt 4.4.

See also setViewMode().

*/
type QMdiArea__ViewMode = int

// Display sub-windows with window frames (default).
const QMdiArea__SubWindowView QMdiArea__ViewMode = 0

// Display sub-windows with tabs in a tab bar.
const QMdiArea__TabbedView QMdiArea__ViewMode = 1

func (this *QMdiArea) ViewModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QMdiArea_ViewModeItemName(val int) string {
	var nilthis *QMdiArea
	return nilthis.ViewModeItemName(val)
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
