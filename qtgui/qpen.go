package qtgui

// /usr/include/qt/QtGui/qpen.h
// #include <qpen.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QPen struct {
	*qtrt.CObject
}
type QPen_ITF interface {
	QPen_PTR() *QPen
}

func (ptr *QPen) QPen_PTR() *QPen { return ptr }

func (this *QPen) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPen) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPenFromPointer(cthis unsafe.Pointer) *QPen {
	return &QPen{&qtrt.CObject{cthis}}
}
func (*QPen) NewFromPointer(cthis unsafe.Pointer) *QPen {
	return NewQPenFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpen.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPen()

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit() *QPen {
	return NewQPen()
}
func NewQPen() *QPen {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPen(Qt::PenStyle)

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit_1(arg0 int) *QPen {
	return NewQPen_1(arg0)
}
func NewQPen_1(arg0 int) *QPen {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2EN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:65
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPen(const QColor &)

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit_2(color QColor_ITF) *QPen {
	return NewQPen_2(color)
}
func NewQPen_2(color QColor_ITF) *QPen {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2ERK6QColor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPen(const QBrush &, qreal, Qt::PenStyle, Qt::PenCapStyle, Qt::PenJoinStyle)

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit_3(brush QBrush_ITF, width float64, s int, c int, j int) *QPen {
	return NewQPen_3(brush, width, s, c, j)
}
func NewQPen_3(brush QBrush_ITF, width float64, s int, c int, j int) *QPen {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2ERK6QBrushdN2Qt8PenStyleENS3_11PenCapStyleENS3_12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, convArg0, width, s, c, j)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPen(const QBrush &, qreal, Qt::PenStyle, Qt::PenCapStyle, Qt::PenJoinStyle)

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit_3_(brush QBrush_ITF, width float64) *QPen {
	return NewQPen_3_(brush, width)
}
func NewQPen_3_(brush QBrush_ITF, width float64) *QPen {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	// arg: 2, Qt::PenStyle=Elaborated, Qt::PenStyle=Enum, , Invalid
	s := 0
	// arg: 3, Qt::PenCapStyle=Elaborated, Qt::PenCapStyle=Enum, , Invalid
	c := 0
	// arg: 4, Qt::PenJoinStyle=Elaborated, Qt::PenJoinStyle=Enum, , Invalid
	j := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2ERK6QBrushdN2Qt8PenStyleENS3_11PenCapStyleENS3_12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, convArg0, width, s, c, j)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPen(const QBrush &, qreal, Qt::PenStyle, Qt::PenCapStyle, Qt::PenJoinStyle)

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit_3_1(brush QBrush_ITF, width float64, s int) *QPen {
	return NewQPen_3_1(brush, width, s)
}
func NewQPen_3_1(brush QBrush_ITF, width float64, s int) *QPen {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	// arg: 3, Qt::PenCapStyle=Elaborated, Qt::PenCapStyle=Enum, , Invalid
	c := 0
	// arg: 4, Qt::PenJoinStyle=Elaborated, Qt::PenJoinStyle=Enum, , Invalid
	j := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2ERK6QBrushdN2Qt8PenStyleENS3_11PenCapStyleENS3_12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, convArg0, width, s, c, j)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPen(const QBrush &, qreal, Qt::PenStyle, Qt::PenCapStyle, Qt::PenJoinStyle)

/*
Constructs a default black solid line pen with 1 width.
*/
func (*QPen) NewForInherit_3_2(brush QBrush_ITF, width float64, s int, c int) *QPen {
	return NewQPen_3_2(brush, width, s, c)
}
func NewQPen_3_2(brush QBrush_ITF, width float64, s int, c int) *QPen {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	// arg: 4, Qt::PenJoinStyle=Elaborated, Qt::PenJoinStyle=Enum, , Invalid
	j := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenC2ERK6QBrushdN2Qt8PenStyleENS3_11PenCapStyleENS3_12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, convArg0, width, s, c, j)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPen()

/*

 */
func DeleteQPen(this *QPen) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpen.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QPen & operator=(const QPen &)

/*

 */
func (this *QPen) Operator_equal(pen QPen_ITF) *QPen {
	var convArg0 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg0 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qpen.h:76
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPen & operator=(QPen &&)

/*

 */
func (this *QPen) Operator_equal_1(other unsafe.Pointer /*333*/) *QPen {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qpen.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPen &)

/*
Swaps pen other with this pen. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QPen) Swap(other QPen_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPen_PTR() != nil {
		convArg0 = other.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenStyle style() const

/*
Returns the pen style.

See also setStyle() and Pen Style.
*/
func (this *QPen) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(Qt::PenStyle)

/*
Sets the pen style to the given style.

See the Qt::PenStyle documentation for a list of the available styles. Since Qt 4.1 it is also possible to specify a custom dash pattern using the setDashPattern() function which implicitly converts the style of the pen to Qt::CustomDashLine.

Note: This function resets the dash offset to zero.

See also style() and Pen Style.
*/
func (this *QPen) SetStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setStyleEN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dashOffset() const

/*
Returns the dash offset for the pen.

See also setDashOffset().
*/
func (this *QPen) DashOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen10dashOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDashOffset(qreal)

/*
Sets the dash offset (the starting point on the dash pattern) for this pen to the offset specified. The offset is measured in terms of the units used to specify the dash pattern.


 For example, a pattern where each stroke is four units long, followed by a gap of two units, will begin with the stroke when drawn as a line.However, if the dash offset is set to 4.0, any line drawn will begin with the gap. Values of the offset up to 4.0 will cause part of the stroke to be drawn first, and values of the offset between 4.0 and 6.0 will cause the line to begin with part of the gap.



Note: This implicitly converts the style of the pen to Qt::CustomDashLine.

See also dashOffset().
*/
func (this *QPen) SetDashOffset(doffset float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen13setDashOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), doffset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal miterLimit() const

/*
Returns the miter limit of the pen. The miter limit is only relevant when the join style is set to Qt::MiterJoin.

See also setMiterLimit() and Join Style.
*/
func (this *QPen) MiterLimit() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen10miterLimitEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMiterLimit(qreal)

/*
Sets the miter limit of this pen to the given limit.



The miter limit describes how far a miter join can extend from the join point. This is used to reduce artifacts between line joins where the lines are close to parallel.

This value does only have effect when the pen style is set to Qt::MiterJoin. The value is specified in units of the pen's width, e.g. a miter limit of 5 in width 10 is 50 pixels long. The default miter limit is 2, i.e. twice the pen width in pixels.

See also miterLimit(), setJoinStyle(), and Join Style.
*/
func (this *QPen) SetMiterLimit(limit float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen13setMiterLimitEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal widthF() const

/*
Returns the pen width with floating point precision.

See also setWidthF() and width().
*/
func (this *QPen) WidthF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen6widthFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidthF(qreal)

/*
Sets the pen width to the given width in pixels with floating point precision.

A line width of zero indicates a cosmetic pen. This means that the pen width is always drawn one pixel wide, independent of the transformation on the painter.

Setting a pen width with a negative value is not supported.

See also setWidth() and widthF().
*/
func (this *QPen) SetWidthF(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen9setWidthFEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the pen width with integer precision.

See also setWidth() and widthF().
*/
func (this *QPen) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpen.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(int)

/*
Sets the pen width to the given width in pixels with integer precision.

A line width of zero indicates a cosmetic pen. This means that the pen width is always drawn one pixel wide, independent of the transformation set on the painter.

Setting a pen width with a negative value is not supported.

See also setWidthF() and width().
*/
func (this *QPen) SetWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color() const

/*
Returns the color of this pen's brush.

See also brush() and setColor().
*/
func (this *QPen) Color() *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qpen.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)

/*
Sets the color of this pen's brush to the given color.

See also setBrush() and color().
*/
func (this *QPen) SetColor(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brush() const

/*
Returns the brush used to fill strokes generated with this pen.

See also setBrush().
*/
func (this *QPen) Brush() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen5brushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qpen.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBrush(const QBrush &)

/*
Sets the brush used to fill strokes generated with this pen to the given brush.

See also brush() and setColor().
*/
func (this *QPen) SetBrush(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSolid() const

/*
Returns true if the pen has a solid fill, otherwise false.

See also style() and dashPattern().
*/
func (this *QPen) IsSolid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen7isSolidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenCapStyle capStyle() const

/*
Returns the pen's cap style.

See also setCapStyle() and Cap Style.
*/
func (this *QPen) CapStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen8capStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapStyle(Qt::PenCapStyle)

/*
Sets the pen's cap style to the given style. The default value is Qt::SquareCap.

See also capStyle() and Cap Style.
*/
func (this *QPen) SetCapStyle(pcs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen11setCapStyleEN2Qt11PenCapStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pcs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenJoinStyle joinStyle() const

/*
Returns the pen's join style.

See also setJoinStyle() and Join Style.
*/
func (this *QPen) JoinStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen9joinStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setJoinStyle(Qt::PenJoinStyle)

/*
Sets the pen's join style to the given style. The default value is Qt::BevelJoin.

See also joinStyle() and Join Style.
*/
func (this *QPen) SetJoinStyle(pcs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen12setJoinStyleEN2Qt12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pcs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCosmetic() const

/*
Returns true if the pen is cosmetic; otherwise returns false.

Cosmetic pens are used to draw strokes that have a constant width regardless of any transformations applied to the QPainter they are used with. Drawing a shape with a cosmetic pen ensures that its outline will have the same thickness at different scale factors.

A zero width pen is cosmetic by default.

See also setCosmetic() and widthF().
*/
func (this *QPen) IsCosmetic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen10isCosmeticEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCosmetic(bool)

/*
Sets this pen to cosmetic or non-cosmetic, depending on the value of cosmetic.

See also isCosmetic().
*/
func (this *QPen) SetCosmetic(cosmetic bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen11setCosmeticEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cosmetic)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QPen &) const

/*

 */
func (this *QPen) Operator_equal_equal(p QPen_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPen_PTR() != nil {
		convArg0 = p.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPeneqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QPen &) const

/*

 */
func (this *QPen) Operator_not_equal(p QPen_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPen_PTR() != nil {
		convArg0 = p.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPenneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()

/*

 */
func (this *QPen) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
}

//  keep block end
