package qtgui

// /usr/include/qt/QtGui/qpainterpath.h
// #include <qpainterpath.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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

type QPainterPath struct {
	*qtrt.CObject
}
type QPainterPath_ITF interface {
	QPainterPath_PTR() *QPainterPath
}

func (ptr *QPainterPath) QPainterPath_PTR() *QPainterPath { return ptr }

func (this *QPainterPath) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPainterPath) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPainterPathFromPointer(cthis unsafe.Pointer) *QPainterPath {
	return &QPainterPath{&qtrt.CObject{cthis}}
}
func (*QPainterPath) NewFromPointer(cthis unsafe.Pointer) *QPainterPath {
	return NewQPainterPathFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpainterpath.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPainterPath()
func NewQPainterPath() *QPainterPath {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainterPath)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:92
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPainterPath(const QPointF &)
func NewQPainterPath_1(startPoint qtcore.QPointF_ITF) *QPainterPath {
	var convArg0 unsafe.Pointer
	if startPoint != nil && startPoint.QPointF_PTR() != nil {
		convArg0 = startPoint.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainterPath)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath & operator=(const QPainterPath &)
func (this *QPainterPath) Operator_equal(other QPainterPath_ITF) *QPainterPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:96
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPainterPath & operator=(QPainterPath &&)
func (this *QPainterPath) Operator_equal_1(other unsafe.Pointer /*333*/) *QPainterPath {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPainterPath()
func DeleteQPainterPath(this *QPainterPath) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpainterpath.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPainterPath &)
func (this *QPainterPath) Swap(other QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeSubpath()
func (this *QPainterPath) CloseSubpath() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12closeSubpathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveTo(const QPointF &)
func (this *QPainterPath) MoveTo(p qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6moveToERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:105
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(qreal, qreal)
func (this *QPainterPath) MoveTo_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6moveToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lineTo(const QPointF &)
func (this *QPainterPath) LineTo(p qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6lineToERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:108
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void lineTo(qreal, qreal)
func (this *QPainterPath) LineTo_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6lineToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void arcMoveTo(const QRectF &, qreal)
func (this *QPainterPath) ArcMoveTo(rect qtcore.QRectF_ITF, angle float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9arcMoveToERK6QRectFd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:111
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void arcMoveTo(qreal, qreal, qreal, qreal, qreal)
func (this *QPainterPath) ArcMoveTo_1(x float64, y float64, w float64, h float64, angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9arcMoveToEddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void arcTo(const QRectF &, qreal, qreal)
func (this *QPainterPath) ArcTo(rect qtcore.QRectF_ITF, startAngle float64, arcLength float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath5arcToERK6QRectFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, startAngle, arcLength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:114
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QPainterPath) ArcTo_1(x float64, y float64, w float64, h float64, startAngle float64, arcLength float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath5arcToEdddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, startAngle, arcLength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cubicTo(const QPointF &, const QPointF &, const QPointF &)
func (this *QPainterPath) CubicTo(ctrlPt1 qtcore.QPointF_ITF, ctrlPt2 qtcore.QPointF_ITF, endPt qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if ctrlPt1 != nil && ctrlPt1.QPointF_PTR() != nil {
		convArg0 = ctrlPt1.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if ctrlPt2 != nil && ctrlPt2.QPointF_PTR() != nil {
		convArg1 = ctrlPt2.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if endPt != nil && endPt.QPointF_PTR() != nil {
		convArg2 = endPt.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7cubicToERK7QPointFS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:117
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QPainterPath) CubicTo_1(ctrlPt1x float64, ctrlPt1y float64, ctrlPt2x float64, ctrlPt2y float64, endPtx float64, endPty float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7cubicToEdddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ctrlPt1x, ctrlPt1y, ctrlPt2x, ctrlPt2y, endPtx, endPty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quadTo(const QPointF &, const QPointF &)
func (this *QPainterPath) QuadTo(ctrlPt qtcore.QPointF_ITF, endPt qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if ctrlPt != nil && ctrlPt.QPointF_PTR() != nil {
		convArg0 = ctrlPt.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if endPt != nil && endPt.QPointF_PTR() != nil {
		convArg1 = endPt.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6quadToERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:120
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void quadTo(qreal, qreal, qreal, qreal)
func (this *QPainterPath) QuadTo_1(ctrlPtx float64, ctrlPty float64, endPtx float64, endPty float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6quadToEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ctrlPtx, ctrlPty, endPtx, endPty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:122
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF currentPosition() const
func (this *QPainterPath) CurrentPosition() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath15currentPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRect(const QRectF &)
func (this *QPainterPath) AddRect(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:125
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addRect(qreal, qreal, qreal, qreal)
func (this *QPainterPath) AddRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addEllipse(const QRectF &)
func (this *QPainterPath) AddEllipse(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:127
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addEllipse(qreal, qreal, qreal, qreal)
func (this *QPainterPath) AddEllipse_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:128
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void addEllipse(const QPointF &, qreal, qreal)
func (this *QPainterPath) AddEllipse_2(center qtcore.QPointF_ITF, rx float64, ry float64) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseERK7QPointFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPolygon(const QPolygonF &)
func (this *QPainterPath) AddPolygon(polygon QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addPolygonERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addText(const QPointF &, const QFont &, const QString &)
func (this *QPainterPath) AddText(point qtcore.QPointF_ITF, f QFont_ITF, text string) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg1 = f.QFont_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:131
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addText(qreal, qreal, const QFont &, const QString &)
func (this *QPainterPath) AddText_1(x float64, y float64, f QFont_ITF, text string) {
	var convArg2 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg2 = f.QFont_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addTextEddRK5QFontRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPath(const QPainterPath &)
func (this *QPainterPath) AddPath(path QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addPathERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRegion(const QRegion &)
func (this *QPainterPath) AddRegion(region QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9addRegionERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect(rect qtcore.QRectF_ITF, xRadius float64, yRadius float64, mode int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectERK6QRectFddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect__(rect qtcore.QRectF_ITF, xRadius float64, yRadius float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 3, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectERK6QRectFddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundedRect(qreal, qreal, qreal, qreal, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect_1(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectEddddddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundedRect(qreal, qreal, qreal, qreal, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect_1_(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64) {
	// arg: 6, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectEddddddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRoundRect(const QRectF &, int, int)
func (this *QPainterPath) AddRoundRect(rect qtcore.QRectF_ITF, xRnd int, yRnd int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRnd, yRnd)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:142
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundRect(qreal, qreal, qreal, qreal, int, int)
func (this *QPainterPath) AddRoundRect_1(x float64, y float64, w float64, h float64, xRnd int, yRnd int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRnd, yRnd)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundRect(const QRectF &, int)
func (this *QPainterPath) AddRoundRect_2(rect qtcore.QRectF_ITF, roundness int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectERK6QRectFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, roundness)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:145
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundRect(qreal, qreal, qreal, qreal, int)
func (this *QPainterPath) AddRoundRect_3(x float64, y float64, w float64, h float64, roundness int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectEddddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, roundness)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectPath(const QPainterPath &)
func (this *QPainterPath) ConnectPath(path QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath11connectPathERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const
func (this *QPainterPath) Contains(pt qtcore.QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPointF_PTR() != nil {
		convArg0 = pt.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:151
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRectF &) const
func (this *QPainterPath) Contains_1(rect qtcore.QRectF_ITF) bool {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8containsERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:187
// index:2
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPainterPath &) const
func (this *QPainterPath) Contains_2(p QPainterPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainterPath_PTR() != nil {
		convArg0 = p.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8containsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRectF &) const
func (this *QPainterPath) Intersects(rect qtcore.QRectF_ITF) bool {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10intersectsERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:186
// index:1
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QPainterPath &) const
func (this *QPainterPath) Intersects_1(p QPainterPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainterPath_PTR() != nil {
		convArg0 = p.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QPainterPath) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:155
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)
func (this *QPainterPath) Translate_1(offset qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg0 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath translated(qreal, qreal) const
func (this *QPainterPath) Translated(dx float64, dy float64) *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:158
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPainterPath translated(const QPointF &) const
func (this *QPainterPath) Translated_1(offset qtcore.QPointF_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPointF_PTR() != nil {
		convArg0 = offset.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10translatedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:160
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect() const
func (this *QPainterPath) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:161
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF controlPointRect() const
func (this *QPainterPath) ControlPointRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath16controlPointRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FillRule fillRule() const
func (this *QPainterPath) FillRule() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8fillRuleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFillRule(Qt::FillRule)
func (this *QPainterPath) SetFillRule(fillRule int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath11setFillRuleEN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QPainterPath) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath toReversed() const
func (this *QPainterPath) ToReversed() *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10toReversedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF toFillPolygon(const QMatrix &) const
func (this *QPainterPath) ToFillPolygon(matrix QMatrix_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK7QMatrix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF toFillPolygon(const QMatrix &) const
func (this *QPainterPath) ToFillPolygon__() *QPolygonF /*123*/ {
	// arg: 0, const QMatrix &=LValueReference, QMatrix=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK7QMatrix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:174
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF toFillPolygon(const QTransform &) const
func (this *QPainterPath) ToFillPolygon_1(matrix QTransform_ITF) *QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:176
// index:0
// Public Visibility=Default Availability=Available
// [4] int elementCount() const
func (this *QPainterPath) ElementCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath12elementCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:177
// index:0
// Public Visibility=Default Availability=Available
// [24] QPainterPath::Element elementAt(int) const
func (this *QPainterPath) ElementAt(i int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath9elementAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElementPositionAt(int, qreal, qreal)
func (this *QPainterPath) SetElementPositionAt(i int, x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath20setElementPositionAtEidd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:180
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal length() const
func (this *QPainterPath) Length() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal percentAtLength(qreal) const
func (this *QPainterPath) PercentAtLength(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath15percentAtLengthEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:182
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pointAtPercent(qreal) const
func (this *QPainterPath) PointAtPercent(t float64) *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14pointAtPercentEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal angleAtPercent(qreal) const
func (this *QPainterPath) AngleAtPercent(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14angleAtPercentEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal slopeAtPercent(qreal) const
func (this *QPainterPath) SlopeAtPercent(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14slopeAtPercentEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath united(const QPainterPath &) const
func (this *QPainterPath) United(r QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPainterPath_PTR() != nil {
		convArg0 = r.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath intersected(const QPainterPath &) const
func (this *QPainterPath) Intersected(r QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPainterPath_PTR() != nil {
		convArg0 = r.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:190
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath subtracted(const QPainterPath &) const
func (this *QPainterPath) Subtracted(r QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPainterPath_PTR() != nil {
		convArg0 = r.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10subtractedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:191
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath subtractedInverted(const QPainterPath &) const
func (this *QPainterPath) SubtractedInverted(r QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPainterPath_PTR() != nil {
		convArg0 = r.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath18subtractedInvertedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath simplified() const
func (this *QPainterPath) Simplified() *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:195
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QPainterPath &) const
func (this *QPainterPath) Operator_equal_equal(other QPainterPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPatheqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:196
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QPainterPath &) const
func (this *QPainterPath) Operator_not_equal(other QPainterPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPathneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath operator&(const QPainterPath &) const
func (this *QPainterPath) Operator_and(other QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPathanERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath operator|(const QPainterPath &) const
func (this *QPainterPath) Operator_or(other QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPathorERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath operator+(const QPainterPath &) const
func (this *QPainterPath) Operator_add(other QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPathplERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath operator-(const QPainterPath &) const
func (this *QPainterPath) Operator_minus(other QPainterPath_ITF) *QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPathmiERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath & operator&=(const QPainterPath &)
func (this *QPainterPath) Operator_and_equal(other QPainterPath_ITF) *QPainterPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathaNERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:203
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath & operator|=(const QPainterPath &)
func (this *QPainterPath) Operator_or_equal(other QPainterPath_ITF) *QPainterPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathoRERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:204
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath & operator+=(const QPainterPath &)
func (this *QPainterPath) Operator_add_equal(other QPainterPath_ITF) *QPainterPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:205
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath & operator-=(const QPainterPath &)
func (this *QPainterPath) Operator_minus_equal(other QPainterPath_ITF) *QPainterPath {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPainterPath_PTR() != nil {
		convArg0 = other.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

type QPainterPath__ElementType = int

const QPainterPath__MoveToElement QPainterPath__ElementType = 0
const QPainterPath__LineToElement QPainterPath__ElementType = 1
const QPainterPath__CurveToElement QPainterPath__ElementType = 2
const QPainterPath__CurveToDataElement QPainterPath__ElementType = 3

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
