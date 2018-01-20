//  header block begin
// /usr/include/qt/QtWidgets/qslider.h
// #include <qslider.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QSlider struct {
	*QAbstractSlider
}

func (this *QSlider) GetCthis() unsafe.Pointer {
	return this.QAbstractSlider.GetCthis()
}

// /usr/include/qt/QtWidgets/qslider.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSlider) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:71
// index:0
// void QSlider(class QWidget *)
func NewQSlider(parent unsafe.Pointer) *QSlider {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSliderC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(cthis)
	return gothis
}
func NewQSliderFromPointer(cthis unsafe.Pointer) *QSlider {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QSlider{bcthis0}
}

// /usr/include/qt/QtWidgets/qslider.h:72
// index:1
// void QSlider(Qt::Orientation, class QWidget *)
func NewQSlider_1(orientation int, parent unsafe.Pointer) *QSlider {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSliderC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:74
// index:0
// virtual
// void ~QSlider()
func DeleteQSlider(*QSlider) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSliderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:76
// index:0
// virtual
// QSize sizeHint()
func (this *QSlider) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:77
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QSlider) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:79
// index:0
// void setTickPosition(enum QSlider::TickPosition)
func (this *QSlider) SetTickPosition(position int) {
	// 0: (, position QSlider::TickPosition), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider15setTickPositionENS_12TickPositionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:80
// index:0
// QSlider::TickPosition tickPosition()
func (this *QSlider) TickPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider12tickPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:82
// index:0
// void setTickInterval(int)
func (this *QSlider) SetTickInterval(ti int) {
	// 0: (, ti int), (&ti)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider15setTickIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ti)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:83
// index:0
// int tickInterval()
func (this *QSlider) TickInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider12tickIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:85
// index:0
// virtual
// bool event(class QEvent *)
func (this *QSlider) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:88
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QSlider) PaintEvent(ev unsafe.Pointer) {
	// 0: (, ev QPaintEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:89
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSlider) MousePressEvent(ev unsafe.Pointer) {
	// 0: (, ev QMouseEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:90
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSlider) MouseReleaseEvent(ev unsafe.Pointer) {
	// 0: (, ev QMouseEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:91
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QSlider) MouseMoveEvent(ev unsafe.Pointer) {
	// 0: (, ev QMouseEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:92
// index:0
// void initStyleOption(class QStyleOptionSlider *)
func (this *QSlider) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionSlider *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
