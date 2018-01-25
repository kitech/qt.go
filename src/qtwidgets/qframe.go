package qtwidgets

// /usr/include/qt/QtWidgets/qframe.h
// #include <qframe.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QFrame struct {
	*QWidget
}

func (this *QFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QFrame) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQFrameFromPointer(cthis unsafe.Pointer) *QFrame {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QFrame{bcthis0}
}
func (*QFrame) NewFromPointer(cthis unsafe.Pointer) *QFrame {
	return NewQFrameFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qframe.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFrame) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public
// void QFrame(class QWidget *, Qt::WindowFlags)
func NewQFrame(parent *QWidget /*444 QWidget **/, f int) *QFrame {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:65
// index:0
// Public virtual
// void ~QFrame()
func DeleteQFrame(*QFrame) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrameD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:67
// index:0
// Public
// int frameStyle()
func (this *QFrame) FrameStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10frameStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qframe.h:68
// index:0
// Public
// void setFrameStyle(int)
func (this *QFrame) SetFrameStyle(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame13setFrameStyleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:70
// index:0
// Public
// int frameWidth()
func (this *QFrame) FrameWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10frameWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qframe.h:72
// index:0
// Public virtual
// QSize sizeHint()
func (this *QFrame) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:96
// index:0
// Public
// QFrame::Shape frameShape()
func (this *QFrame) FrameShape() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame10frameShapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qframe.h:97
// index:0
// Public
// void setFrameShape(enum QFrame::Shape)
func (this *QFrame) SetFrameShape(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame13setFrameShapeENS_5ShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:98
// index:0
// Public
// QFrame::Shadow frameShadow()
func (this *QFrame) FrameShadow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame11frameShadowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qframe.h:99
// index:0
// Public
// void setFrameShadow(enum QFrame::Shadow)
func (this *QFrame) SetFrameShadow(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame14setFrameShadowENS_6ShadowE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:101
// index:0
// Public
// int lineWidth()
func (this *QFrame) LineWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame9lineWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qframe.h:102
// index:0
// Public
// void setLineWidth(int)
func (this *QFrame) SetLineWidth(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame12setLineWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:104
// index:0
// Public
// int midLineWidth()
func (this *QFrame) MidLineWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame12midLineWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qframe.h:105
// index:0
// Public
// void setMidLineWidth(int)
func (this *QFrame) SetMidLineWidth(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame15setMidLineWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:107
// index:0
// Public
// QRect frameRect()
func (this *QFrame) FrameRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame9frameRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:108
// index:0
// Public
// void setFrameRect(const class QRect &)
func (this *QFrame) SetFrameRect(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame12setFrameRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:111
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QFrame) Event(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qframe.h:112
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QFrame) PaintEvent(arg0 *qtgui.QPaintEvent /*444 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:113
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QFrame) ChangeEvent(arg0 *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:114
// index:0
// Protected
// void drawFrame(class QPainter *)
func (this *QFrame) DrawFrame(arg0 *qtgui.QPainter /*444 QPainter **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QFrame9drawFrameEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:119
// index:0
// Protected
// void initStyleOption(class QStyleOptionFrame *)
func (this *QFrame) InitStyleOption(option *QStyleOptionFrame /*444 QStyleOptionFrame **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QFrame__Shape = int

const QFrame__NoFrame QFrame__Shape = 0
const QFrame__Box QFrame__Shape = 1
const QFrame__Panel QFrame__Shape = 2
const QFrame__WinPanel QFrame__Shape = 3
const QFrame__HLine QFrame__Shape = 4
const QFrame__VLine QFrame__Shape = 5
const QFrame__StyledPanel QFrame__Shape = 6

type QFrame__Shadow = int

const QFrame__Plain QFrame__Shadow = 16
const QFrame__Raised QFrame__Shadow = 32
const QFrame__Sunken QFrame__Shadow = 48

type QFrame__StyleMask = int

const QFrame__Shadow_Mask QFrame__StyleMask = 240
const QFrame__Shape_Mask QFrame__StyleMask = 15

//  body block end
