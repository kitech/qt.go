package qtwidgets

// /usr/include/qt/QtWidgets/qsizegrip.h
// #include <qsizegrip.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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

// void paintEvent(QPaintEvent *)
func (this *QSizeGrip) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QSizeGrip) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QSizeGrip) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QSizeGrip) InheritMouseReleaseEvent(f func(mouseEvent *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void moveEvent(QMoveEvent *)
func (this *QSizeGrip) InheritMoveEvent(f func(moveEvent *qtgui.QMoveEvent /*777 QMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QSizeGrip) InheritShowEvent(f func(showEvent *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QSizeGrip) InheritHideEvent(f func(hideEvent *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QSizeGrip) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(QEvent *)
func (this *QSizeGrip) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QSizeGrip struct {
	*QWidget
}
type QSizeGrip_ITF interface {
	QWidget_ITF
	QSizeGrip_PTR() *QSizeGrip
}

func (ptr *QSizeGrip) QSizeGrip_PTR() *QSizeGrip { return ptr }

func (this *QSizeGrip) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QSizeGrip) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQSizeGripFromPointer(cthis unsafe.Pointer) *QSizeGrip {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSizeGrip{bcthis0}
}
func (*QSizeGrip) NewFromPointer(cthis unsafe.Pointer) *QSizeGrip {
	return NewQSizeGripFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSizeGrip) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSizeGrip10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsizegrip.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSizeGrip(QWidget *)

/*
Constructs a resize corner as a child widget of the given parent.
*/
func (*QSizeGrip) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QSizeGrip {
	return NewQSizeGrip(parent)
}
func NewQSizeGrip(parent QWidget_ITF /*777 QWidget **/) *QSizeGrip {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGripC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeGripFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSizeGrip")
	return gothis
}

// /usr/include/qt/QtWidgets/qsizegrip.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSizeGrip()

/*

 */
func DeleteQSizeGrip(this *QSizeGrip) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGripD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QSizeGrip) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSizeGrip8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qsizegrip.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().
*/
func (this *QSizeGrip) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:62
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().

Paints the resize grip.

Resize grips are usually rendered as small diagonal textured lines in the lower-right corner. The paint event is passed in the event parameter.
*/
func (this *QSizeGrip) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:63
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().

Receives the mouse press events for the widget, and primes the resize operation. The mouse press event is passed in the event parameter.
*/
func (this *QSizeGrip) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:64
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().

Resizes the top-level widget containing this widget. The mouse move event is passed in the event parameter.
*/
func (this *QSizeGrip) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:65
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QSizeGrip) MouseReleaseEvent(mouseEvent qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if mouseEvent != nil && mouseEvent.QMouseEvent_PTR() != nil {
		convArg0 = mouseEvent.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)

/*
Reimplemented from QWidget::moveEvent().
*/
func (this *QSizeGrip) MoveEvent(moveEvent qtgui.QMoveEvent_ITF /*777 QMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if moveEvent != nil && moveEvent.QMoveEvent_PTR() != nil {
		convArg0 = moveEvent.QMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:67
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QSizeGrip) ShowEvent(showEvent qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if showEvent != nil && showEvent.QShowEvent_PTR() != nil {
		convArg0 = showEvent.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QSizeGrip) HideEvent(hideEvent qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if hideEvent != nil && hideEvent.QHideEvent_PTR() != nil {
		convArg0 = hideEvent.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizegrip.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QSizeGrip) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizegrip.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSizeGrip) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSizeGrip5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
