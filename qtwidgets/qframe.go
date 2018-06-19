package qtwidgets

// /usr/include/qt/QtWidgets/qframe.h
// #include <qframe.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

// bool event(QEvent *)
func (this *QFrame) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(QPaintEvent *)
func (this *QFrame) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void changeEvent(QEvent *)
func (this *QFrame) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void drawFrame(QPainter *)
func (this *QFrame) InheritDrawFrame(f func(arg0 *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawFrame", f)
}

// void initStyleOption(QStyleOptionFrame *)
func (this *QFrame) InheritInitStyleOption(f func(option *QStyleOptionFrame /*777 QStyleOptionFrame **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QFrame struct {
	*QWidget
}
type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func (ptr *QFrame) QFrame_PTR() *QFrame { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFrame) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)

/*
Constructs a frame widget with frame style NoFrame and a 1-pixel frame width.

The parent and f arguments are passed to the QWidget constructor.
*/
func NewQFrame(parent QWidget_ITF /*777 QWidget **/, f int) *QFrame {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)

/*
Constructs a frame widget with frame style NoFrame and a 1-pixel frame width.

The parent and f arguments are passed to the QWidget constructor.
*/
func NewQFrame__() *QFrame {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)

/*
Constructs a frame widget with frame style NoFrame and a 1-pixel frame width.

The parent and f arguments are passed to the QWidget constructor.
*/
func NewQFrame__1(parent QWidget_ITF /*777 QWidget **/) *QFrame {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFrame()

/*

 */
func DeleteQFrame(this *QFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qframe.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameStyle() const

/*
Returns the frame style.

The default value is QFrame::Plain.

See also setFrameStyle(), frameShape(), and frameShadow().
*/
func (this *QFrame) FrameStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10frameStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameStyle(int)

/*
Sets the frame style to style.

The style is the bitwise OR between a frame shape and a frame shadow style. See the picture of the frames in the main class documentation.

The frame shapes are given in QFrame::Shape and the shadow styles in QFrame::Shadow.

If a mid-line width greater than 0 is specified, an additional line is drawn for Raised or Sunken Box, HLine, and VLine frames. The mid-color of the current color group is used for drawing middle lines.

See also frameStyle().
*/
func (this *QFrame) SetFrameStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame13setFrameStyleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameWidth() const

/*

 */
func (this *QFrame) FrameWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10frameWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QFrame) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] QFrame::Shape frameShape() const

/*

 */
func (this *QFrame) FrameShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame10frameShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qframe.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameShape(QFrame::Shape)

/*

 */
func (this *QFrame) SetFrameShape(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame13setFrameShapeENS_5ShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QFrame::Shadow frameShadow() const

/*

 */
func (this *QFrame) FrameShadow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame11frameShadowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qframe.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameShadow(QFrame::Shadow)

/*

 */
func (this *QFrame) SetFrameShadow(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame14setFrameShadowENS_6ShadowE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineWidth() const

/*

 */
func (this *QFrame) LineWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame9lineWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWidth(int)

/*

 */
func (this *QFrame) SetLineWidth(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame12setLineWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int midLineWidth() const

/*

 */
func (this *QFrame) MidLineWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame12midLineWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMidLineWidth(int)

/*

 */
func (this *QFrame) SetMidLineWidth(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame15setMidLineWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:107
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameRect() const

/*

 */
func (this *QFrame) FrameRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame9frameRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qframe.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameRect(const QRect &)

/*

 */
func (this *QFrame) SetFrameRect(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame12setFrameRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QFrame) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qframe.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QFrame) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QFrame) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:114
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawFrame(QPainter *)

/*

 */
func (this *QFrame) DrawFrame(arg0 qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPainter_PTR() != nil {
		convArg0 = arg0.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QFrame9drawFrameEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:119
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionFrame *) const

/*
Initializes option with the values from this QFrame. This method is useful for subclasses when they need a QStyleOptionFrame but don't want to fill in all the information themselves.

This function was introduced in  Qt 5.5.

See also QStyleOption::initFrom().
*/
func (this *QFrame) InitStyleOption(option QStyleOptionFrame_ITF /*777 QStyleOptionFrame **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionFrame_PTR() != nil {
		convArg0 = option.QStyleOptionFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum type defines the shapes of frame available.



When it does not call QStyle, Shape interacts with QFrame::Shadow, the lineWidth() and the midLineWidth() to create the total result. See the picture of the frames in the main class documentation.

See also QFrame::Shadow, QFrame::style(), and QStyle::drawPrimitive().

*/
type QFrame__Shape = int

// QFrame draws nothing
const QFrame__NoFrame QFrame__Shape = 0

//
const QFrame__Box QFrame__Shape = 1

//
const QFrame__Panel QFrame__Shape = 2

//
const QFrame__WinPanel QFrame__Shape = 3

//
const QFrame__HLine QFrame__Shape = 4

//
const QFrame__VLine QFrame__Shape = 5

//
const QFrame__StyledPanel QFrame__Shape = 6

/*
This enum type defines the types of shadow that are used to give a 3D effect to frames.



Shadow interacts with QFrame::Shape, the lineWidth() and the midLineWidth(). See the picture of the frames in the main class documentation.

See also QFrame::Shape, lineWidth(), and midLineWidth().

*/
type QFrame__Shadow = int

//
const QFrame__Plain QFrame__Shadow = 16

//
const QFrame__Raised QFrame__Shadow = 32

//
const QFrame__Sunken QFrame__Shadow = 48

/*
This enum defines two constants that can be used to extract the two components of frameStyle():



Normally, you don't need to use these, since frameShadow() and frameShape() already extract the Shadow and the Shape parts of frameStyle().

See also frameStyle() and setFrameStyle().

*/
type QFrame__StyleMask = int

//
const QFrame__Shadow_Mask QFrame__StyleMask = 240

//
const QFrame__Shape_Mask QFrame__StyleMask = 15

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
