package qtgui

// /usr/include/qt/QtGui/qpen.h
// #include <qpen.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPen(Qt::PenStyle)
func NewQPen_1(arg0 int) *QPen {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2EN2Qt8PenStyleE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2ERK6QColor", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2ERK6QBrushdN2Qt8PenStyleENS3_11PenCapStyleENS3_12PenJoinStyleE", ffiqt.FFI_TYPE_POINTER, convArg0, width, s, c, j)
	gopp.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPen)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPen()
func DeleteQPen(this *QPen) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpen.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPen &)
func (this *QPen) Swap(other *QPen) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenStyle style()
func (this *QPen) Style() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(Qt::PenStyle)
func (this *QPen) SetStyle(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setStyleEN2Qt8PenStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dashOffset()
func (this *QPen) DashOffset() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen10dashOffsetEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDashOffset(qreal)
func (this *QPen) SetDashOffset(doffset float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen13setDashOffsetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), doffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal miterLimit()
func (this *QPen) MiterLimit() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen10miterLimitEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMiterLimit(qreal)
func (this *QPen) SetMiterLimit(limit float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen13setMiterLimitEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal widthF()
func (this *QPen) WidthF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen6widthFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpen.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidthF(qreal)
func (this *QPen) SetWidthF(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen9setWidthFEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int width()
func (this *QPen) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpen.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(int)
func (this *QPen) SetWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color()
func (this *QPen) Color() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5colorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brush()
func (this *QPen) Brush() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5brushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSolid()
func (this *QPen) IsSolid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen7isSolidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenCapStyle capStyle()
func (this *QPen) CapStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen8capStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapStyle(Qt::PenCapStyle)
func (this *QPen) SetCapStyle(pcs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen11setCapStyleEN2Qt11PenCapStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pcs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::PenJoinStyle joinStyle()
func (this *QPen) JoinStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen9joinStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpen.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setJoinStyle(Qt::PenJoinStyle)
func (this *QPen) SetJoinStyle(pcs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen12setJoinStyleEN2Qt12PenJoinStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pcs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCosmetic()
func (this *QPen) IsCosmetic() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen10isCosmeticEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpen.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCosmetic(_Bool)
func (this *QPen) SetCosmetic(cosmetic bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen11setCosmeticEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cosmetic)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QPen) IsDetached() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
