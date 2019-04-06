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

// void setViewportMargins(int, int, int, int)
func (this *QAbstractScrollArea) InheritSetViewportMargins(f func(left int, top int, right int, bottom int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setViewportMargins", f)
}

// QMargins viewportMargins()
func (this *QAbstractScrollArea) InheritViewportMargins(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportMargins", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QAbstractScrollArea) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(QEvent *)
func (this *QAbstractScrollArea) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool viewportEvent(QEvent *)
func (this *QAbstractScrollArea) InheritViewportEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QAbstractScrollArea) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QAbstractScrollArea) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QAbstractScrollArea) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractScrollArea) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QAbstractScrollArea) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractScrollArea) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QAbstractScrollArea) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QAbstractScrollArea) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QAbstractScrollArea) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QAbstractScrollArea) InheritDragMoveEvent(f func(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QAbstractScrollArea) InheritDragLeaveEvent(f func(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QAbstractScrollArea) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QAbstractScrollArea) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QAbstractScrollArea) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// QSize viewportSizeHint()
func (this *QAbstractScrollArea) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

/*

 */
type QAbstractScrollArea struct {
	*QFrame
}
type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea { return ptr }

func (this *QAbstractScrollArea) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QAbstractScrollArea) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQAbstractScrollAreaFromPointer(cthis unsafe.Pointer) *QAbstractScrollArea {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QAbstractScrollArea{bcthis0}
}
func (*QAbstractScrollArea) NewFromPointer(cthis unsafe.Pointer) *QAbstractScrollArea {
	return NewQAbstractScrollAreaFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().

This function is called with key event e when key presses occur. It handles PageUp, PageDown, Up, Down, Left, and Right, and ignores all other key presses.
*/
func (this *QAbstractScrollArea) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*
This virtual handler is called when the scroll bars are moved by dx, dy, and consequently the viewport's contents should be scrolled accordingly.

The default implementation simply calls update() on the entire viewport(), subclasses can reimplement this handler for optimization purposes, or - like QScrollArea - to move a contents widget. The parameters dx and dy are there for convenience, so that the class knows how much should be scrolled (useful e.g. when doing pixel-shifts). You may just as well ignore these values and scroll directly to the position the scroll bars indicate.

Calling this function in order to scroll programmatically is an error, use the scroll bars instead (e.g. by calling QScrollBar::setValue() directly).
*/
func (this *QAbstractScrollArea) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractScrollArea16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint() const

/*
Returns the recommended size for the viewport. The default implementation returns viewport()->sizeHint(). Note that the size is just the viewport's size, without any scroll bars visible.

This function was introduced in  Qt 5.2.
*/
func (this *QAbstractScrollArea) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractScrollArea16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

/*
This enum specifies how the size hint of the QAbstractScrollArea should adjust when the size of the viewport changes.



This enum was introduced or modified in  Qt 5.2.

*/
type QAbstractScrollArea__SizeAdjustPolicy = int

// The scroll area will behave like before - and not do any adjust.
const QAbstractScrollArea__AdjustIgnored QAbstractScrollArea__SizeAdjustPolicy = 0

// The scroll area will adjust to its viewport the first time it is shown.
const QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = 1

// The scroll area will always adjust to the viewport
const QAbstractScrollArea__AdjustToContents QAbstractScrollArea__SizeAdjustPolicy = 2

func (this *QAbstractScrollArea) SizeAdjustPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractScrollArea_SizeAdjustPolicyItemName(val int) string {
	var nilthis *QAbstractScrollArea
	return nilthis.SizeAdjustPolicyItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11059() {
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
