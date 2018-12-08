package qtwidgets

// /usr/include/qt/QtWidgets/qscrollbar.h
// #include <qscrollbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

// void wheelEvent(QWheelEvent *)
func (this *QScrollBar) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QScrollBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QScrollBar) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QScrollBar) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QScrollBar) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QScrollBar) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void sliderChange(QAbstractSlider::SliderChange)
func (this *QScrollBar) InheritSliderChange(f func(change int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sliderChange", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QScrollBar) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void initStyleOption(QStyleOptionSlider *)
func (this *QScrollBar) InheritInitStyleOption(f func(option *QStyleOptionSlider /*777 QStyleOptionSlider **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QScrollBar struct {
	*QAbstractSlider
}
type QScrollBar_ITF interface {
	QAbstractSlider_ITF
	QScrollBar_PTR() *QScrollBar
}

func (ptr *QScrollBar) QScrollBar_PTR() *QScrollBar { return ptr }

func (this *QScrollBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSlider.GetCthis()
	}
}
func (this *QScrollBar) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSlider = NewQAbstractSliderFromPointer(cthis)
}
func NewQScrollBarFromPointer(cthis unsafe.Pointer) *QScrollBar {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QScrollBar{bcthis0}
}
func (*QScrollBar) NewFromPointer(cthis unsafe.Pointer) *QScrollBar {
	return NewQScrollBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QScrollBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QScrollBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qscrollbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollBar(QWidget *)

/*
Constructs a vertical scroll bar.

The parent argument is sent to the QWidget constructor.

The minimum defaults to 0, the maximum to 99, with a singleStep size of 1 and a pageStep size of 10, and an initial value of 0.
*/
func (*QScrollBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QScrollBar {
	return NewQScrollBar(parent)
}
func NewQScrollBar(parent QWidget_ITF /*777 QWidget **/) *QScrollBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QScrollBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollBar(QWidget *)

/*
Constructs a vertical scroll bar.

The parent argument is sent to the QWidget constructor.

The minimum defaults to 0, the maximum to 99, with a singleStep size of 1 and a pageStep size of 10, and an initial value of 0.
*/
func (*QScrollBar) NewForInheritp() *QScrollBar {
	return NewQScrollBarp()
}
func NewQScrollBarp() *QScrollBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QScrollBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QScrollBar(Qt::Orientation, QWidget *)

/*
Constructs a vertical scroll bar.

The parent argument is sent to the QWidget constructor.

The minimum defaults to 0, the maximum to 99, with a singleStep size of 1 and a pageStep size of 10, and an initial value of 0.
*/
func (*QScrollBar) NewForInherit1(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QScrollBar {
	return NewQScrollBar1(arg0, parent)
}
func NewQScrollBar1(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QScrollBar {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBarC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QScrollBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QScrollBar(Qt::Orientation, QWidget *)

/*
Constructs a vertical scroll bar.

The parent argument is sent to the QWidget constructor.

The minimum defaults to 0, the maximum to 99, with a singleStep size of 1 and a pageStep size of 10, and an initial value of 0.
*/
func (*QScrollBar) NewForInherit1p(arg0 int) *QScrollBar {
	return NewQScrollBar1p(arg0)
}
func NewQScrollBar1p(arg0 int) *QScrollBar {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBarC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QScrollBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScrollBar()

/*

 */
func DeleteQScrollBar(this *QScrollBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QScrollBar) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QScrollBar8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollbar.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QScrollBar) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollbar.h:68
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QScrollBar) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QScrollBar) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QScrollBar) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:72
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QScrollBar) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QScrollBar) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:74
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QScrollBar) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:75
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sliderChange(QAbstractSlider::SliderChange)

/*
Reimplemented from QAbstractSlider::sliderChange().
*/
func (this *QScrollBar) SliderChange(change int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().
*/
func (this *QScrollBar) ContextMenuEvent(arg0 qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QContextMenuEvent_PTR() != nil {
		convArg0 = arg0.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:79
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSlider *) const

/*
Initialize option with the values from this QScrollBar. This method is useful for subclasses when they need a QStyleOptionSlider, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QScrollBar) InitStyleOption(option QStyleOptionSlider_ITF /*777 QStyleOptionSlider **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionSlider_PTR() != nil {
		convArg0 = option.QStyleOptionSlider_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
