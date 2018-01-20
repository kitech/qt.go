//  header block begin
// /usr/include/qt/QtGui/qpen.h
// #include <qpen.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.Cthis
}

// /usr/include/qt/QtGui/qpen.h:63
// index:0
// void QPen()
func NewQPen() *QPen {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(cthis)
	return gothis
}
func NewQPenFromPointer(cthis unsafe.Pointer) *QPen {
	return &QPen{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpen.h:64
// index:1
// void QPen(Qt::PenStyle)
func NewQPen_1(arg0 int) *QPen {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2EN2Qt8PenStyleE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:65
// index:2
// void QPen(const class QColor &)
func NewQPen_2(color unsafe.Pointer) *QPen {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2ERK6QColor", ffiqt.FFI_TYPE_VOID, cthis, color)
	gopp.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:66
// index:3
// void QPen(const class QBrush &, qreal, Qt::PenStyle, Qt::PenCapStyle, Qt::PenJoinStyle)
func NewQPen_3(brush unsafe.Pointer, width float64, s int, c int, j int) *QPen {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenC2ERK6QBrushdN2Qt8PenStyleENS3_11PenCapStyleENS3_12PenJoinStyleE", ffiqt.FFI_TYPE_VOID, cthis, brush, &width, &s, &c, &j)
	gopp.ErrPrint(err, rv)
	gothis := NewQPenFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpen.h:70
// index:0
// void ~QPen()
func DeleteQPen(*QPen) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPenD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:79
// index:0
// inline
// void swap(class QPen &)
func (this *QPen) Swap(other unsafe.Pointer) {
	// 0: (, other QPen &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:81
// index:0
// Qt::PenStyle style()
func (this *QPen) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5styleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:82
// index:0
// void setStyle(Qt::PenStyle)
func (this *QPen) SetStyle(arg0 int) {
	// 0: (, Qt::PenStyle arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setStyleEN2Qt8PenStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:84
// index:0
// QVector<qreal> dashPattern()
func (this *QPen) DashPattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen11dashPatternEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:87
// index:0
// qreal dashOffset()
func (this *QPen) DashOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen10dashOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:88
// index:0
// void setDashOffset(qreal)
func (this *QPen) SetDashOffset(doffset float64) {
	// 0: (, doffset qreal), (&doffset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen13setDashOffsetEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &doffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:90
// index:0
// qreal miterLimit()
func (this *QPen) MiterLimit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen10miterLimitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:91
// index:0
// void setMiterLimit(qreal)
func (this *QPen) SetMiterLimit(limit float64) {
	// 0: (, limit qreal), (&limit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen13setMiterLimitEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:93
// index:0
// qreal widthF()
func (this *QPen) WidthF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen6widthFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:94
// index:0
// void setWidthF(qreal)
func (this *QPen) SetWidthF(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen9setWidthFEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:96
// index:0
// int width()
func (this *QPen) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:97
// index:0
// void setWidth(int)
func (this *QPen) SetWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:99
// index:0
// QColor color()
func (this *QPen) Color() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5colorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:100
// index:0
// void setColor(const class QColor &)
func (this *QPen) SetColor(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:102
// index:0
// QBrush brush()
func (this *QPen) Brush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen5brushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:103
// index:0
// void setBrush(const class QBrush &)
func (this *QPen) SetBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen8setBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:105
// index:0
// bool isSolid()
func (this *QPen) IsSolid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen7isSolidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:107
// index:0
// Qt::PenCapStyle capStyle()
func (this *QPen) CapStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen8capStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:108
// index:0
// void setCapStyle(Qt::PenCapStyle)
func (this *QPen) SetCapStyle(pcs int) {
	// 0: (, pcs Qt::PenCapStyle), (&pcs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen11setCapStyleEN2Qt11PenCapStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pcs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:110
// index:0
// Qt::PenJoinStyle joinStyle()
func (this *QPen) JoinStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen9joinStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:111
// index:0
// void setJoinStyle(Qt::PenJoinStyle)
func (this *QPen) SetJoinStyle(pcs int) {
	// 0: (, pcs Qt::PenJoinStyle), (&pcs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen12setJoinStyleEN2Qt12PenJoinStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pcs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:113
// index:0
// bool isCosmetic()
func (this *QPen) IsCosmetic() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QPen10isCosmeticEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:114
// index:0
// void setCosmetic(_Bool)
func (this *QPen) SetCosmetic(cosmetic bool) {
	// 0: (, cosmetic bool), (&cosmetic)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen11setCosmeticEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cosmetic)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:120
// index:0
// bool isDetached()
func (this *QPen) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QPen10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
