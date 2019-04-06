//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractscrollarea.h
// #include <qabstractscrollarea.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractScrollArea) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)

/*
Constructs a viewport.

The parent argument is sent to the QWidget constructor.
*/
func (*QAbstractScrollArea) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractScrollArea {
	return NewQAbstractScrollArea(parent)
}
func NewQAbstractScrollArea(parent QWidget_ITF /*777 QWidget **/) *QAbstractScrollArea {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollAreaC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractScrollAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)

/*
Constructs a viewport.

The parent argument is sent to the QWidget constructor.
*/
func (*QAbstractScrollArea) NewForInheritp() *QAbstractScrollArea {
	return NewQAbstractScrollAreap()
}
func NewQAbstractScrollAreap() *QAbstractScrollArea {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollAreaC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractScrollAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractScrollArea()

/*

 */
func DeleteQAbstractScrollArea(this *QAbstractScrollArea) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollAreaD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScrollBarPolicy verticalScrollBarPolicy() const

/*

 */
func (this *QAbstractScrollArea) VerticalScrollBarPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea23verticalScrollBarPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalScrollBarPolicy(Qt::ScrollBarPolicy)

/*

 */
func (this *QAbstractScrollArea) SetVerticalScrollBarPolicy(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea26setVerticalScrollBarPolicyEN2Qt15ScrollBarPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QScrollBar * verticalScrollBar() const

/*
Returns the vertical scroll bar.

See also setVerticalScrollBar(), verticalScrollBarPolicy, and horizontalScrollBar().
*/
func (this *QAbstractScrollArea) VerticalScrollBar() *QScrollBar /*777 QScrollBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea17verticalScrollBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalScrollBar(QScrollBar *)

/*
Replaces the existing vertical scroll bar with scrollBar, and sets all the former scroll bar's slider properties on the new scroll bar. The former scroll bar is then deleted.

QAbstractScrollArea already provides vertical and horizontal scroll bars by default. You can call this function to replace the default vertical scroll bar with your own custom scroll bar.

This function was introduced in  Qt 4.2.

See also verticalScrollBar() and setHorizontalScrollBar().
*/
func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar QScrollBar_ITF /*777 QScrollBar **/) {
	var convArg0 unsafe.Pointer
	if scrollbar != nil && scrollbar.QScrollBar_PTR() != nil {
		convArg0 = scrollbar.QScrollBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScrollBarPolicy horizontalScrollBarPolicy() const

/*

 */
func (this *QAbstractScrollArea) HorizontalScrollBarPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea25horizontalScrollBarPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalScrollBarPolicy(Qt::ScrollBarPolicy)

/*

 */
func (this *QAbstractScrollArea) SetHorizontalScrollBarPolicy(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea28setHorizontalScrollBarPolicyEN2Qt15ScrollBarPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QScrollBar * horizontalScrollBar() const

/*
Returns the horizontal scroll bar.

See also setHorizontalScrollBar(), horizontalScrollBarPolicy, and verticalScrollBar().
*/
func (this *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar /*777 QScrollBar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea19horizontalScrollBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalScrollBar(QScrollBar *)

/*
Replaces the existing horizontal scroll bar with scrollBar, and sets all the former scroll bar's slider properties on the new scroll bar. The former scroll bar is then deleted.

QAbstractScrollArea already provides horizontal and vertical scroll bars by default. You can call this function to replace the default horizontal scroll bar with your own custom scroll bar.

This function was introduced in  Qt 4.2.

See also horizontalScrollBar() and setVerticalScrollBar().
*/
func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar QScrollBar_ITF /*777 QScrollBar **/) {
	var convArg0 unsafe.Pointer
	if scrollbar != nil && scrollbar.QScrollBar_PTR() != nil {
		convArg0 = scrollbar.QScrollBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget() const

/*
Returns the widget in the corner between the two scroll bars.

By default, no corner widget is present.

This function was introduced in  Qt 4.2.

See also setCornerWidget().
*/
func (this *QAbstractScrollArea) CornerWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea12cornerWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *)

/*
Sets the widget in the corner between the two scroll bars to be widget.

You will probably also want to set at least one of the scroll bar modes to AlwaysOn.

Passing 0 shows no widget in the corner.

Any previous corner widget is hidden.

You may call setCornerWidget() with the same widget at different times.

All widgets set here will be deleted by the scroll area when it is destroyed unless you separately reparent the widget after setting some other corner widget (or 0).

Any newly set widget should have no current parent.

By default, no corner widget is present.

This function was introduced in  Qt 4.2.

See also cornerWidget(), horizontalScrollBarPolicy, and horizontalScrollBarPolicy.
*/
func (this *QAbstractScrollArea) SetCornerWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addScrollBarWidget(QWidget *, Qt::Alignment)

/*
Adds widget as a scroll bar widget in the location specified by alignment.

Scroll bar widgets are shown next to the horizontal or vertical scroll bar, and can be placed on either side of it. If you want the scroll bar widgets to be always visible, set the scrollBarPolicy for the corresponding scroll bar to AlwaysOn.

alignment must be one of Qt::Alignleft and Qt::AlignRight, which maps to the horizontal scroll bar, or Qt::AlignTop and Qt::AlignBottom, which maps to the vertical scroll bar.

A scroll bar widget can be removed by either re-parenting the widget or deleting it. It's also possible to hide a widget with QWidget::hide()

The scroll bar widget will be resized to fit the scroll bar geometry for the current style. The following describes the case for scroll bar widgets on the horizontal scroll bar:

The height of the widget will be set to match the height of the scroll bar. To control the width of the widget, use QWidget::setMinimumWidth and QWidget::setMaximumWidth, or implement QWidget::sizeHint() and set a horizontal size policy. If you want a square widget, call QStyle::pixelMetric(QStyle::PM_ScrollBarExtent) and set the width to this value.

This function was introduced in  Qt 4.2.

See also scrollBarWidgets().
*/
func (this *QAbstractScrollArea) AddScrollBarWidget(widget QWidget_ITF /*777 QWidget **/, alignment int) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea18addScrollBarWidgetEP7QWidget6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] QWidgetList scrollBarWidgets(Qt::Alignment)

/*
Returns a list of the currently set scroll bar widgets. alignment can be any combination of the four location flags.

This function was introduced in  Qt 4.2.

See also addScrollBarWidget().
*/
func (this *QAbstractScrollArea) ScrollBarWidgets(alignment int) *QWidgetList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea16scrollBarWidgetsE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
	rv2 := /*222*/ NewQWidgetListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * viewport() const

/*
Returns the viewport widget.

Use the QScrollArea::widget() function to retrieve the contents of the viewport widget.

See also setViewport() and QScrollArea::widget().
*/
func (this *QAbstractScrollArea) Viewport() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea8viewportEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewport(QWidget *)

/*
Sets the viewport to be the given widget. The QAbstractScrollArea will take ownership of the given widget.

If widget is 0, QAbstractScrollArea will assign a new QWidget instance for the viewport.

This function was introduced in  Qt 4.2.

See also viewport().
*/
func (this *QAbstractScrollArea) SetViewport(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea11setViewportEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumViewportSize() const

/*
Returns the size of the viewport as if the scroll bars had no valid scrolling range.
*/
func (this *QAbstractScrollArea) MaximumViewportSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea19maximumViewportSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QAbstractScrollArea) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QFrame::sizeHint().

Returns the sizeHint property of the scroll area. The size is determined by using viewportSizeHint() plus some extra space for scroll bars, if needed.
*/
func (this *QAbstractScrollArea) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setupViewport(QWidget *)

/*
This slot is called by QAbstractScrollArea after setViewport(viewport) has been called. Reimplement this function in a subclass of QAbstractScrollArea to initialize the new viewport before it is used.

See also setViewport().
*/
func (this *QAbstractScrollArea) SetupViewport(viewport QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if viewport != nil && viewport.QWidget_PTR() != nil {
		convArg0 = viewport.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea13setupViewportEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractScrollArea::SizeAdjustPolicy sizeAdjustPolicy() const

/*

 */
func (this *QAbstractScrollArea) SizeAdjustPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea16sizeAdjustPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeAdjustPolicy(QAbstractScrollArea::SizeAdjustPolicy)

/*

 */
func (this *QAbstractScrollArea) SetSizeAdjustPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setViewportMargins(int, int, int, int)

/*
Sets the margins around the scrolling area to left, top, right and bottom. This is useful for applications such as spreadsheets with "locked" rows and columns. The marginal space is is left blank; put widgets in the unused area.

Note that this function is frequently called by QTreeView and QTableView, so margins must be implemented by QAbstractScrollArea subclasses. Also, if the subclasses are to be used in item views, they should not call this function.

By default all margins are zero.

See also viewportMargins().
*/
func (this *QAbstractScrollArea) SetViewportMargins(left int, top int, right int, bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea18setViewportMarginsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:106
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void setViewportMargins(const QMargins &)

/*
Sets the margins around the scrolling area to left, top, right and bottom. This is useful for applications such as spreadsheets with "locked" rows and columns. The marginal space is is left blank; put widgets in the unused area.

Note that this function is frequently called by QTreeView and QTableView, so margins must be implemented by QAbstractScrollArea subclasses. Also, if the subclasses are to be used in item views, they should not call this function.

By default all margins are zero.

See also viewportMargins().
*/
func (this *QAbstractScrollArea) SetViewportMargins1(margins qtcore.QMargins_ITF) {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:107
// index:0
// Protected Visibility=Default Availability=Available
// [16] QMargins viewportMargins() const

/*
Returns the margins around the scrolling area. By default all the margins are zero.

This function was introduced in  Qt 5.5.

See also setViewportMargins().
*/
func (this *QAbstractScrollArea) ViewportMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea15viewportMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*

 */
func (this *QAbstractScrollArea) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QFrame::event().

This is the main event handler for the QAbstractScrollArea widget (not the scrolling area viewport()). The specified event is a general event object that may need to be cast to the appropriate class depending on its type.

See also QEvent::type().
*/
func (this *QAbstractScrollArea) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)

/*
The main event handler for the scrolling area (the viewport() widget). It handles the event specified, and can be called by subclasses to provide reasonable default behavior.

Returns true to indicate to the event system that the event has been handled, and needs no further processing; otherwise returns false to indicate that the event should be propagated further.

You can reimplement this function in a subclass, but we recommend using one of the specialized event handlers instead.

Specialized handlers for viewport events are: paintEvent(), mousePressEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), mouseMoveEvent(), wheelEvent(), dragEnterEvent(), dragMoveEvent(), dragLeaveEvent(), dropEvent(), contextMenuEvent(), and resizeEvent().
*/
func (this *QAbstractScrollArea) ViewportEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().

This event handler can be reimplemented in a subclass to receive resize events (passed in event), for the viewport() widget.

When resizeEvent() is called, the viewport already has its new geometry: Its new size is accessible through the QResizeEvent::size() function, and the old size through QResizeEvent::oldSize().

See also QWidget::resizeEvent().
*/
func (this *QAbstractScrollArea) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QFrame::paintEvent().

This event handler can be reimplemented in a subclass to receive paint events (passed in event), for the viewport() widget.

Note: If you open a painter, make sure to open it on the viewport().

See also QWidget::paintEvent().
*/
func (this *QAbstractScrollArea) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:115
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().

This event handler can be reimplemented in a subclass to receive mouse press events for the viewport() widget. The event is passed in e.

See also QWidget::mousePressEvent().
*/
func (this *QAbstractScrollArea) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().

This event handler can be reimplemented in a subclass to receive mouse release events for the viewport() widget. The event is passed in e.

See also QWidget::mouseReleaseEvent().
*/
func (this *QAbstractScrollArea) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().

This event handler can be reimplemented in a subclass to receive mouse double click events for the viewport() widget. The event is passed in e.

See also QWidget::mouseDoubleClickEvent().
*/
func (this *QAbstractScrollArea) MouseDoubleClickEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().

This event handler can be reimplemented in a subclass to receive mouse move events for the viewport() widget. The event is passed in e.

See also QWidget::mouseMoveEvent().
*/
func (this *QAbstractScrollArea) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().

This event handler can be reimplemented in a subclass to receive wheel events for the viewport() widget. The event is passed in e.

See also QWidget::wheelEvent().
*/
func (this *QAbstractScrollArea) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().

This event handler can be reimplemented in a subclass to receive context menu events for the viewport() widget. The event is passed in e.

See also QWidget::contextMenuEvent().
*/
func (this *QAbstractScrollArea) ContextMenuEvent(arg0 qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QContextMenuEvent_PTR() != nil {
		convArg0 = arg0.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:126
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QWidget::dragEnterEvent().

This event handler can be reimplemented in a subclass to receive drag enter events (passed in event), for the viewport() widget.

See also QWidget::dragEnterEvent().
*/
func (this *QAbstractScrollArea) DragEnterEvent(arg0 qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragEnterEvent_PTR() != nil {
		convArg0 = arg0.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:127
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().

This event handler can be reimplemented in a subclass to receive drag move events (passed in event), for the viewport() widget.

See also QWidget::dragMoveEvent().
*/
func (this *QAbstractScrollArea) DragMoveEvent(arg0 qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragMoveEvent_PTR() != nil {
		convArg0 = arg0.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:128
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QWidget::dragLeaveEvent().

This event handler can be reimplemented in a subclass to receive drag leave events (passed in event), for the viewport() widget.

See also QWidget::dragLeaveEvent().
*/
func (this *QAbstractScrollArea) DragLeaveEvent(arg0 qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragLeaveEvent_PTR() != nil {
		convArg0 = arg0.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:129
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().

This event handler can be reimplemented in a subclass to receive drop events (passed in event), for the viewport() widget.

See also QWidget::dropEvent().
*/
func (this *QAbstractScrollArea) DropEvent(arg0 qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDropEvent_PTR() != nil {
		convArg0 = arg0.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11060() {
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
