package qtwidgets

// /usr/include/qt/QtWidgets/qslider.h
// #include <qslider.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
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

// void paintEvent(class QPaintEvent *)
func (this *QSlider) InheritPaintEvent(f func(ev *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QSlider) InheritMousePressEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSlider) InheritMouseReleaseEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QSlider) InheritMouseMoveEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void initStyleOption(class QStyleOptionSlider *)
func (this *QSlider) InheritInitStyleOption(f func(option *QStyleOptionSlider /*777 QStyleOptionSlider **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QSlider struct {
	*QAbstractSlider
}
type QSlider_ITF interface {
	QAbstractSlider_ITF
	QSlider_PTR() *QSlider
}

func (ptr *QSlider) QSlider_PTR() *QSlider { return ptr }

func (this *QSlider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSlider.GetCthis()
	}
}
func (this *QSlider) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSlider = NewQAbstractSliderFromPointer(cthis)
}
func NewQSliderFromPointer(cthis unsafe.Pointer) *QSlider {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QSlider{bcthis0}
}
func (*QSlider) NewFromPointer(cthis unsafe.Pointer) *QSlider {
	return NewQSliderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qslider.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QSlider) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSlider10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qslider.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSlider(QWidget *)
func NewQSlider(parent QWidget_ITF /*777 QWidget **/) *QSlider {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSliderC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSlider")
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSlider(QWidget *)
func NewQSlider__() *QSlider {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSliderC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSlider")
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSlider(Qt::Orientation, QWidget *)
func NewQSlider_1(orientation int, parent QWidget_ITF /*777 QWidget **/) *QSlider {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSliderC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSlider")
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSlider(Qt::Orientation, QWidget *)
func NewQSlider_1_(orientation int) *QSlider {
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSliderC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSlider")
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSlider()
func DeleteQSlider(this *QSlider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSliderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qslider.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const
func (this *QSlider) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSlider8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qslider.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const
func (this *QSlider) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSlider15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qslider.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTickPosition(enum QSlider::TickPosition)
func (this *QSlider) SetTickPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider15setTickPositionENS_12TickPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] QSlider::TickPosition tickPosition() const
func (this *QSlider) TickPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSlider12tickPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qslider.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTickInterval(int)
func (this *QSlider) SetTickInterval(ti int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider15setTickIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ti)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int tickInterval() const
func (this *QSlider) TickInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSlider12tickIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qslider.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSlider) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qslider.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QSlider) PaintEvent(ev qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QPaintEvent_PTR() != nil {
		convArg0 = ev.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QSlider) MousePressEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QSlider) MouseReleaseEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QSlider) MouseMoveEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSlider14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:92
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSlider *) const
func (this *QSlider) InitStyleOption(option QStyleOptionSlider_ITF /*777 QStyleOptionSlider **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionSlider_PTR() != nil {
		convArg0 = option.QStyleOptionSlider_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QSlider__TickPosition = int

const QSlider__NoTicks QSlider__TickPosition = 0
const QSlider__TicksAbove QSlider__TickPosition = 1
const QSlider__TicksLeft QSlider__TickPosition = 1
const QSlider__TicksBelow QSlider__TickPosition = 2
const QSlider__TicksRight QSlider__TickPosition = 2
const QSlider__TicksBothSides QSlider__TickPosition = 3

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
