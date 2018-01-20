//  header block begin
// /usr/include/qt/QtWidgets/qdial.h
// #include <qdial.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QDial struct {
	*QAbstractSlider
}

func (this *QDial) GetCthis() unsafe.Pointer {
	return this.QAbstractSlider.GetCthis()
}

// /usr/include/qt/QtWidgets/qdial.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDial) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:64
// index:0
// void QDial(class QWidget *)
func NewQDial(parent unsafe.Pointer) *QDial {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDialC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialFromPointer(cthis)
	return gothis
}
func NewQDialFromPointer(cthis unsafe.Pointer) *QDial {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QDial{bcthis0}
}

// /usr/include/qt/QtWidgets/qdial.h:66
// index:0
// virtual
// void ~QDial()
func DeleteQDial(*QDial) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDialD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:68
// index:0
// bool wrapping()
func (this *QDial) Wrapping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial8wrappingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:70
// index:0
// int notchSize()
func (this *QDial) NotchSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial9notchSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:72
// index:0
// void setNotchTarget(double)
func (this *QDial) SetNotchTarget(target float64) {
	// 0: (, target double), (&target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial14setNotchTargetEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:73
// index:0
// qreal notchTarget()
func (this *QDial) NotchTarget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial11notchTargetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:74
// index:0
// bool notchesVisible()
func (this *QDial) NotchesVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial14notchesVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:76
// index:0
// virtual
// QSize sizeHint()
func (this *QDial) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:77
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QDial) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:80
// index:0
// void setNotchesVisible(_Bool)
func (this *QDial) SetNotchesVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial17setNotchesVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:81
// index:0
// void setWrapping(_Bool)
func (this *QDial) SetWrapping(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial11setWrappingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:84
// index:0
// virtual
// bool event(class QEvent *)
func (this *QDial) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:85
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QDial) ResizeEvent(re unsafe.Pointer) {
	// 0: (, re QResizeEvent *), (re)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), re)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:86
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QDial) PaintEvent(pe unsafe.Pointer) {
	// 0: (, pe QPaintEvent *), (pe)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pe)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:88
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QDial) MousePressEvent(me unsafe.Pointer) {
	// 0: (, me QMouseEvent *), (me)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), me)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:89
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QDial) MouseReleaseEvent(me unsafe.Pointer) {
	// 0: (, me QMouseEvent *), (me)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), me)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:90
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QDial) MouseMoveEvent(me unsafe.Pointer) {
	// 0: (, me QMouseEvent *), (me)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), me)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:92
// index:0
// virtual
// void sliderChange(enum QAbstractSlider::SliderChange)
func (this *QDial) SliderChange(change int) {
	// 0: (, change QAbstractSlider::SliderChange), (&change)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:93
// index:0
// void initStyleOption(class QStyleOptionSlider *)
func (this *QDial) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionSlider *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
