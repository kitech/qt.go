package qtwidgets

// /usr/include/qt/QtWidgets/qscrollbar.h
// #include <qscrollbar.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QScrollBar struct {
	*QAbstractSlider
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QScrollBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollbar.h:59
// index:0
// Public
// void QScrollBar(QWidget *)
func NewQScrollBar(parent *QWidget /*777 QWidget **/) *QScrollBar {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:60
// index:1
// Public
// void QScrollBar(Qt::Orientation, QWidget *)
func NewQScrollBar_1(arg0 int, parent *QWidget /*777 QWidget **/) *QScrollBar {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:61
// index:0
// Public virtual
// void ~QScrollBar()
func DeleteQScrollBar(*QScrollBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:63
// index:0
// Public virtual
// QSize sizeHint()
func (this *QScrollBar) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollbar.h:64
// index:0
// Public virtual
// bool event(QEvent *)
func (this *QScrollBar) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollbar.h:68
// index:0
// Protected virtual
// void wheelEvent(QWheelEvent *)
func (this *QScrollBar) WheelEvent(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:70
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QScrollBar) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:71
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QScrollBar) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:72
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QScrollBar) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:73
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QScrollBar) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:74
// index:0
// Protected virtual
// void hideEvent(QHideEvent *)
func (this *QScrollBar) HideEvent(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:75
// index:0
// Protected virtual
// void sliderChange(QAbstractSlider::SliderChange)
func (this *QScrollBar) SliderChange(change int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:77
// index:0
// Protected virtual
// void contextMenuEvent(QContextMenuEvent *)
func (this *QScrollBar) ContextMenuEvent(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:79
// index:0
// Protected
// void initStyleOption(QStyleOptionSlider *)
func (this *QScrollBar) InitStyleOption(option *QStyleOptionSlider /*777 QStyleOptionSlider **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
