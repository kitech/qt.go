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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPen struct {
	*qtrt.CObject
}

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
func NewQPen_2(color *QColor) *QPen {
	var convArg0 = color.GetCthis()
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
func NewQPen_3(brush *QBrush, width float64, s int, c int, j int) *QPen {
	var convArg0 = brush.GetCthis()
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
func DeleteQPen(this *QPen) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPenD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpen.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPen &)
func (this *QPen) Swap(other *QPen) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenStyle style()
func (this *QPen) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(Qt::PenStyle)
func (this *QPen) SetStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setStyleEN2Qt8PenStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dashOffset()
func (this *QPen) DashOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen10dashOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDashOffset(qreal)
func (this *QPen) SetDashOffset(doffset float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen13setDashOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), doffset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal miterLimit()
func (this *QPen) MiterLimit() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen10miterLimitEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMiterLimit(qreal)
func (this *QPen) SetMiterLimit(limit float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen13setMiterLimitEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal widthF()
func (this *QPen) WidthF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen6widthFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidthF(qreal)
func (this *QPen) SetWidthF(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen9setWidthFEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int width()
func (this *QPen) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpen.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(int)
func (this *QPen) SetWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color()
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
func (this *QPen) SetColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brush()
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
func (this *QPen) SetBrush(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen8setBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSolid()
func (this *QPen) IsSolid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen7isSolidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenCapStyle capStyle()
func (this *QPen) CapStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen8capStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapStyle(Qt::PenCapStyle)
func (this *QPen) SetCapStyle(pcs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen11setCapStyleEN2Qt11PenCapStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pcs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenJoinStyle joinStyle()
func (this *QPen) JoinStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen9joinStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setJoinStyle(Qt::PenJoinStyle)
func (this *QPen) SetJoinStyle(pcs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen12setJoinStyleEN2Qt12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pcs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCosmetic()
func (this *QPen) IsCosmetic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QPen10isCosmeticEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCosmetic(_Bool)
func (this *QPen) SetCosmetic(cosmetic bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QPen11setCosmeticEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cosmetic)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
