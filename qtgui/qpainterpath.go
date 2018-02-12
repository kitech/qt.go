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
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPainterPath struct {
	*qtrt.CObject
}

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
func NewQPainterPath_1(startPoint *qtcore.QPointF) *QPainterPath {
	var convArg0 = startPoint.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPathC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPainterPath)
	return gothis
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
func (this *QPainterPath) Swap(other *QPainterPath) {
	var convArg0 = other.GetCthis()
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
func (this *QPainterPath) MoveTo(p *qtcore.QPointF) {
	var convArg0 = p.GetCthis()
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
func (this *QPainterPath) LineTo(p *qtcore.QPointF) {
	var convArg0 = p.GetCthis()
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
func (this *QPainterPath) ArcMoveTo(rect *qtcore.QRectF, angle float64) {
	var convArg0 = rect.GetCthis()
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
func (this *QPainterPath) ArcTo(rect *qtcore.QRectF, startAngle float64, arcLength float64) {
	var convArg0 = rect.GetCthis()
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
func (this *QPainterPath) CubicTo(ctrlPt1 *qtcore.QPointF, ctrlPt2 *qtcore.QPointF, endPt *qtcore.QPointF) {
	var convArg0 = ctrlPt1.GetCthis()
	var convArg1 = ctrlPt2.GetCthis()
	var convArg2 = endPt.GetCthis()
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
func (this *QPainterPath) QuadTo(ctrlPt *qtcore.QPointF, endPt *qtcore.QPointF) {
	var convArg0 = ctrlPt.GetCthis()
	var convArg1 = endPt.GetCthis()
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
// [16] QPointF currentPosition()
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
func (this *QPainterPath) AddRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
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
func (this *QPainterPath) AddEllipse(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
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
func (this *QPainterPath) AddEllipse_2(center *qtcore.QPointF, rx float64, ry float64) {
	var convArg0 = center.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseERK7QPointFdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, rx, ry)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPolygon(const QPolygonF &)
func (this *QPainterPath) AddPolygon(polygon *QPolygonF) {
	var convArg0 = polygon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addPolygonERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addText(const QPointF &, const QFont &, const QString &)
func (this *QPainterPath) AddText(point *qtcore.QPointF, f *QFont, text string) {
	var convArg0 = point.GetCthis()
	var convArg1 = f.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:131
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addText(qreal, qreal, const QFont &, const QString &)
func (this *QPainterPath) AddText_1(x float64, y float64, f *QFont, text string) {
	var convArg2 = f.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addTextEddRK5QFontRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPath(const QPainterPath &)
func (this *QPainterPath) AddPath(path *QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addPathERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRegion(const QRegion &)
func (this *QPainterPath) AddRegion(region *QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9addRegionERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRoundedRect(const QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect(rect *qtcore.QRectF, xRadius float64, yRadius float64, mode int) {
	var convArg0 = rect.GetCthis()
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

// /usr/include/qt/QtGui/qpainterpath.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRoundRect(const QRectF &, int, int)
func (this *QPainterPath) AddRoundRect(rect *qtcore.QRectF, xRnd int, yRnd int) {
	var convArg0 = rect.GetCthis()
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
func (this *QPainterPath) AddRoundRect_2(rect *qtcore.QRectF, roundness int) {
	var convArg0 = rect.GetCthis()
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
func (this *QPainterPath) ConnectPath(path *QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath11connectPathERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QPainterPath) Contains(pt *qtcore.QPointF) bool {
	var convArg0 = pt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:151
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRectF &)
func (this *QPainterPath) Contains_1(rect *qtcore.QRectF) bool {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8containsERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:187
// index:2
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPainterPath &)
func (this *QPainterPath) Contains_2(p *QPainterPath) bool {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8containsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRectF &)
func (this *QPainterPath) Intersects(rect *qtcore.QRectF) bool {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10intersectsERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:186
// index:1
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QPainterPath &)
func (this *QPainterPath) Intersects_1(p *QPainterPath) bool {
	var convArg0 = p.GetCthis()
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
func (this *QPainterPath) Translate_1(offset *qtcore.QPointF) {
	var convArg0 = offset.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath translated(qreal, qreal)
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
// [8] QPainterPath translated(const QPointF &)
func (this *QPainterPath) Translated_1(offset *qtcore.QPointF) *QPainterPath /*123*/ {
	var convArg0 = offset.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10translatedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:160
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
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
// [32] QRectF controlPointRect()
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
// [4] Qt::FillRule fillRule()
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
// [1] bool isEmpty()
func (this *QPainterPath) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath toReversed()
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
// [8] QPolygonF toFillPolygon(const QMatrix &)
func (this *QPainterPath) ToFillPolygon(matrix *QMatrix) *QPolygonF /*123*/ {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK7QMatrix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:174
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF toFillPolygon(const QTransform &)
func (this *QPainterPath) ToFillPolygon_1(matrix *QTransform) *QPolygonF /*123*/ {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:176
// index:0
// Public Visibility=Default Availability=Available
// [4] int elementCount()
func (this *QPainterPath) ElementCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath12elementCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:177
// index:0
// Public Visibility=Default Availability=Available
// [24] QPainterPath::Element elementAt(int)
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
// [8] qreal length()
func (this *QPainterPath) Length() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal percentAtLength(qreal)
func (this *QPainterPath) PercentAtLength(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath15percentAtLengthEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:182
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pointAtPercent(qreal)
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
// [8] qreal angleAtPercent(qreal)
func (this *QPainterPath) AngleAtPercent(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14angleAtPercentEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal slopeAtPercent(qreal)
func (this *QPainterPath) SlopeAtPercent(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14slopeAtPercentEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath united(const QPainterPath &)
func (this *QPainterPath) United(r *QPainterPath) *QPainterPath /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath intersected(const QPainterPath &)
func (this *QPainterPath) Intersected(r *QPainterPath) *QPainterPath /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:190
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath subtracted(const QPainterPath &)
func (this *QPainterPath) Subtracted(r *QPainterPath) *QPainterPath /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10subtractedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:191
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath subtractedInverted(const QPainterPath &)
func (this *QPainterPath) SubtractedInverted(r *QPainterPath) *QPainterPath /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath18subtractedInvertedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qpainterpath.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath simplified()
func (this *QPainterPath) Simplified() *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
