package qtwidgets

// /usr/include/qt/QtWidgets/qdial.h
// #include <qdial.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if this == nil {
		return nil
	} else {
		return this.QAbstractSlider.GetCthis()
	}
}
func (this *QDial) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSlider = NewQAbstractSliderFromPointer(cthis)
}
func NewQDialFromPointer(cthis unsafe.Pointer) *QDial {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QDial{bcthis0}
}
func (*QDial) NewFromPointer(cthis unsafe.Pointer) *QDial {
	return NewQDialFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdial.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDial) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdial.h:64
// index:0
// Public
// void QDial(QWidget *)
func NewQDial(parent *QWidget /*777 QWidget **/) *QDial {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDialC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdial.h:66
// index:0
// Public virtual
// void ~QDial()
func DeleteQDial(*QDial) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDialD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:68
// index:0
// Public
// bool wrapping()
func (this *QDial) Wrapping() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial8wrappingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdial.h:70
// index:0
// Public
// int notchSize()
func (this *QDial) NotchSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial9notchSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qdial.h:72
// index:0
// Public
// void setNotchTarget(double)
func (this *QDial) SetNotchTarget(target float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial14setNotchTargetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:73
// index:0
// Public
// qreal notchTarget()
func (this *QDial) NotchTarget() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial11notchTargetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qdial.h:74
// index:0
// Public
// bool notchesVisible()
func (this *QDial) NotchesVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial14notchesVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdial.h:76
// index:0
// Public virtual
// QSize sizeHint()
func (this *QDial) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdial.h:77
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QDial) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdial.h:80
// index:0
// Public
// void setNotchesVisible(bool)
func (this *QDial) SetNotchesVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial17setNotchesVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:81
// index:0
// Public
// void setWrapping(bool)
func (this *QDial) SetWrapping(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial11setWrappingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:84
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QDial) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdial.h:85
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QDial) ResizeEvent(re *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:86
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QDial) PaintEvent(pe *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = pe.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:88
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QDial) MousePressEvent(me *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = me.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:89
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QDial) MouseReleaseEvent(me *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = me.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:90
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QDial) MouseMoveEvent(me *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = me.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:92
// index:0
// Protected virtual
// void sliderChange(QAbstractSlider::SliderChange)
func (this *QDial) SliderChange(change int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdial.h:93
// index:0
// Protected
// void initStyleOption(QStyleOptionSlider *)
func (this *QDial) InitStyleOption(option *QStyleOptionSlider /*777 QStyleOptionSlider **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
