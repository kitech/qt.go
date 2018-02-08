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
import "gopp"
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
// bool event(class QEvent *)
func (this *QFrame) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QFrame) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QFrame) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void drawFrame(class QPainter *)
func (this *QFrame) InheritDrawFrame(f func(arg0 *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawFrame", f)
}

// void initStyleOption(class QStyleOptionFrame *)
func (this *QFrame) InheritInitStyleOption(f func(option *QStyleOptionFrame /*777 QStyleOptionFrame **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFrame) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)
func NewQFrame(parent *QWidget /*777 QWidget **/, f int) *QFrame {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFrame()
func DeleteQFrame(this *QFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qframe.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameStyle()
func (this *QFrame) FrameStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10frameStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameStyle(int)
func (this *QFrame) SetFrameStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame13setFrameStyleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameWidth()
func (this *QFrame) FrameWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10frameWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QFrame) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] QFrame::Shape frameShape()
func (this *QFrame) FrameShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10frameShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qframe.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameShape(enum QFrame::Shape)
func (this *QFrame) SetFrameShape(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame13setFrameShapeENS_5ShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QFrame::Shadow frameShadow()
func (this *QFrame) FrameShadow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame11frameShadowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qframe.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameShadow(enum QFrame::Shadow)
func (this *QFrame) SetFrameShadow(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame14setFrameShadowENS_6ShadowE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineWidth()
func (this *QFrame) LineWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame9lineWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWidth(int)
func (this *QFrame) SetLineWidth(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame12setLineWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int midLineWidth()
func (this *QFrame) MidLineWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame12midLineWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMidLineWidth(int)
func (this *QFrame) SetMidLineWidth(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame15setMidLineWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:107
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameRect()
func (this *QFrame) FrameRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame9frameRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameRect(const QRect &)
func (this *QFrame) SetFrameRect(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame12setFrameRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFrame) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qframe.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QFrame) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QFrame) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:114
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawFrame(QPainter *)
func (this *QFrame) DrawFrame(arg0 *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame9drawFrameEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:119
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionFrame *)
func (this *QFrame) InitStyleOption(option *QStyleOptionFrame /*777 QStyleOptionFrame **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
