// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qscrollarea.h
// #include <qscrollarea.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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

// bool event(QEvent *)
func (this *QScrollArea) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QScrollArea) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QScrollArea) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QScrollArea) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// QSize viewportSizeHint()
func (this *QScrollArea) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

/*

 */
type QScrollArea struct {
	*QAbstractScrollArea
}
type QScrollArea_ITF interface {
	QAbstractScrollArea_ITF
	QScrollArea_PTR() *QScrollArea
}

func (ptr *QScrollArea) QScrollArea_PTR() *QScrollArea { return ptr }

func (this *QScrollArea) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QScrollArea) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQScrollAreaFromPointer(cthis unsafe.Pointer) *QScrollArea {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QScrollArea{bcthis0}
}
func (*QScrollArea) NewFromPointer(cthis unsafe.Pointer) *QScrollArea {
	return NewQScrollAreaFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QScrollArea) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QScrollArea10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscrollarea.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollArea(QWidget *)

/*
Constructs an empty scroll area with the given parent.

See also setWidget().
*/
func (*QScrollArea) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QScrollArea {
	return NewQScrollArea(parent)
}
func NewQScrollArea(parent QWidget_ITF /*777 QWidget **/) *QScrollArea {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollAreaC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollarea.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollArea(QWidget *)

/*
Constructs an empty scroll area with the given parent.

See also setWidget().
*/
func (*QScrollArea) NewForInheritp() *QScrollArea {
	return NewQScrollAreap()
}
func NewQScrollAreap() *QScrollArea {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollAreaC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollarea.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScrollArea()

/*

 */
func DeleteQScrollArea(this *QScrollArea) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollAreaD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the scroll area's widget, or 0 if there is none.

See also setWidget().
*/
func (this *QScrollArea) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QScrollArea6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscrollarea.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*
Sets the scroll area's widget.

The widget becomes a child of the scroll area, and will be destroyed when the scroll area is deleted or when a new widget is set.

The widget's autoFillBackground property will be set to true.

If the scroll area is visible when the widget is added, you must show() it explicitly.

Note that You must add the layout of widget before you call this function; if you add it later, the widget will not be visible - regardless of when you show() the scroll area. In this case, you can also not show() the widget later.

See also widget().
*/
func (this *QScrollArea) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * takeWidget()

/*
Removes the scroll area's widget, and passes ownership of the widget to the caller.

See also widget().
*/
func (this *QScrollArea) TakeWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea10takeWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscrollarea.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool widgetResizable() const

/*

 */
func (this *QScrollArea) WidgetResizable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QScrollArea15widgetResizableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidgetResizable(bool)

/*

 */
func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea18setWidgetResizableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resizable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QAbstractScrollArea::sizeHint().
*/
func (this *QScrollArea) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QScrollArea8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollarea.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QScrollArea) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QScrollArea) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QScrollArea9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QScrollArea) SetAlignment(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(int, int, int, int)

/*
Scrolls the contents of the scroll area so that the point (x, y) is visible inside the region of the viewport with margins specified in pixels by xmargin and ymargin. If the specified point cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.
*/
func (this *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea13ensureVisibleEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(int, int, int, int)

/*
Scrolls the contents of the scroll area so that the point (x, y) is visible inside the region of the viewport with margins specified in pixels by xmargin and ymargin. If the specified point cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.
*/
func (this *QScrollArea) EnsureVisiblep(x int, y int) {
	// arg: 2, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 3, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea13ensureVisibleEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(int, int, int, int)

/*
Scrolls the contents of the scroll area so that the point (x, y) is visible inside the region of the viewport with margins specified in pixels by xmargin and ymargin. If the specified point cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.
*/
func (this *QScrollArea) EnsureVisiblep1(x int, y int, xmargin int) {
	// arg: 3, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea13ensureVisibleEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureWidgetVisible(QWidget *, int, int)

/*
Scrolls the contents of the scroll area so that the childWidget of QScrollArea::widget() is visible inside the viewport with margins specified in pixels by xmargin and ymargin. If the specified point cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

This function was introduced in  Qt 4.2.
*/
func (this *QScrollArea) EnsureWidgetVisible(childWidget QWidget_ITF /*777 QWidget **/, xmargin int, ymargin int) {
	var convArg0 unsafe.Pointer
	if childWidget != nil && childWidget.QWidget_PTR() != nil {
		convArg0 = childWidget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureWidgetVisible(QWidget *, int, int)

/*
Scrolls the contents of the scroll area so that the childWidget of QScrollArea::widget() is visible inside the viewport with margins specified in pixels by xmargin and ymargin. If the specified point cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

This function was introduced in  Qt 4.2.
*/
func (this *QScrollArea) EnsureWidgetVisiblep(childWidget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if childWidget != nil && childWidget.QWidget_PTR() != nil {
		convArg0 = childWidget.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureWidgetVisible(QWidget *, int, int)

/*
Scrolls the contents of the scroll area so that the childWidget of QScrollArea::widget() is visible inside the viewport with margins specified in pixels by xmargin and ymargin. If the specified point cannot be reached, the contents are scrolled to the nearest valid position. The default value for both margins is 50 pixels.

This function was introduced in  Qt 4.2.
*/
func (this *QScrollArea) EnsureWidgetVisiblep1(childWidget QWidget_ITF /*777 QWidget **/, xmargin int) {
	var convArg0 unsafe.Pointer
	if childWidget != nil && childWidget.QWidget_PTR() != nil {
		convArg0 = childWidget.QWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:81
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QAbstractScrollArea::event().
*/
func (this *QScrollArea) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*

 */
func (this *QScrollArea) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QAbstractScrollArea::resizeEvent().
*/
func (this *QScrollArea) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*
Reimplemented from QAbstractScrollArea::scrollContentsBy().
*/
func (this *QScrollArea) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QScrollArea16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:86
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint() const

/*
Reimplemented from QAbstractScrollArea::viewportSizeHint().
*/
func (this *QScrollArea) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QScrollArea16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_11241() {
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
