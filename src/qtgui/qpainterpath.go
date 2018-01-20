//  header block begin
// /usr/include/qt/QtGui/qpainterpath.h
// #include <qpainterpath.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QPainterPath struct {
	*qtrt.CObject
}

func (this *QPainterPath) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qpainterpath.h:91
// index:0
// void QPainterPath()
func NewQPainterPath() *QPainterPath {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPathC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterPathFromPointer(cthis)
	return gothis
}
func NewQPainterPathFromPointer(cthis unsafe.Pointer) *QPainterPath {
	return &QPainterPath{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpainterpath.h:92
// index:1
// void QPainterPath(const class QPointF &)
func NewQPainterPath_1(startPoint unsafe.Pointer) *QPainterPath {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPathC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, startPoint)
	gopp.ErrPrint(err, rv)
	gothis := NewQPainterPathFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpainterpath.h:99
// index:0
// void ~QPainterPath()
func DeleteQPainterPath(*QPainterPath) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPathD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:100
// index:0
// inline
// void swap(class QPainterPath &)
func (this *QPainterPath) Swap(other unsafe.Pointer) {
	// 0: (, other QPainterPath &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:102
// index:0
// void closeSubpath()
func (this *QPainterPath) CloseSubpath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath12closeSubpathEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:104
// index:0
// void moveTo(const class QPointF &)
func (this *QPainterPath) MoveTo(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath6moveToERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:105
// index:1
// inline
// void moveTo(qreal, qreal)
func (this *QPainterPath) MoveTo_1(x float64, y float64) {
	// 1: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath6moveToEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:107
// index:0
// void lineTo(const class QPointF &)
func (this *QPainterPath) LineTo(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath6lineToERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:108
// index:1
// inline
// void lineTo(qreal, qreal)
func (this *QPainterPath) LineTo_1(x float64, y float64) {
	// 1: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath6lineToEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:110
// index:0
// void arcMoveTo(const class QRectF &, qreal)
func (this *QPainterPath) ArcMoveTo(rect unsafe.Pointer, angle float64) {
	// 0: (, rect const QRectF &, angle qreal), (rect, &angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath9arcMoveToERK6QRectFd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:111
// index:1
// inline
// void arcMoveTo(qreal, qreal, qreal, qreal, qreal)
func (this *QPainterPath) ArcMoveTo_1(x float64, y float64, w float64, h float64, angle float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, angle qreal), (&x, &y, &w, &h, &angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath9arcMoveToEddddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:113
// index:0
// void arcTo(const class QRectF &, qreal, qreal)
func (this *QPainterPath) ArcTo(rect unsafe.Pointer, startAngle float64, arcLength float64) {
	// 0: (, rect const QRectF &, startAngle qreal, arcLength qreal), (rect, &startAngle, &arcLength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath5arcToERK6QRectFdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &startAngle, &arcLength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:114
// index:1
// inline
// void arcTo(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QPainterPath) ArcTo_1(x float64, y float64, w float64, h float64, startAngle float64, arcLength float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, startAngle qreal, arcLength qreal), (&x, &y, &w, &h, &startAngle, &arcLength)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath5arcToEdddddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &startAngle, &arcLength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:116
// index:0
// void cubicTo(const class QPointF &, const class QPointF &, const class QPointF &)
func (this *QPainterPath) CubicTo(ctrlPt1 unsafe.Pointer, ctrlPt2 unsafe.Pointer, endPt unsafe.Pointer) {
	// 0: (, ctrlPt1 const QPointF &, ctrlPt2 const QPointF &, endPt const QPointF &), (ctrlPt1, ctrlPt2, endPt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7cubicToERK7QPointFS2_S2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ctrlPt1, ctrlPt2, endPt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:117
// index:1
// inline
// void cubicTo(qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QPainterPath) CubicTo_1(ctrlPt1x float64, ctrlPt1y float64, ctrlPt2x float64, ctrlPt2y float64, endPtx float64, endPty float64) {
	// 1: (, ctrlPt1x qreal, ctrlPt1y qreal, ctrlPt2x qreal, ctrlPt2y qreal, endPtx qreal, endPty qreal), (&ctrlPt1x, &ctrlPt1y, &ctrlPt2x, &ctrlPt2y, &endPtx, &endPty)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7cubicToEdddddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ctrlPt1x, &ctrlPt1y, &ctrlPt2x, &ctrlPt2y, &endPtx, &endPty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:119
// index:0
// void quadTo(const class QPointF &, const class QPointF &)
func (this *QPainterPath) QuadTo(ctrlPt unsafe.Pointer, endPt unsafe.Pointer) {
	// 0: (, ctrlPt const QPointF &, endPt const QPointF &), (ctrlPt, endPt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath6quadToERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ctrlPt, endPt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:120
// index:1
// inline
// void quadTo(qreal, qreal, qreal, qreal)
func (this *QPainterPath) QuadTo_1(ctrlPtx float64, ctrlPty float64, endPtx float64, endPty float64) {
	// 1: (, ctrlPtx qreal, ctrlPty qreal, endPtx qreal, endPty qreal), (&ctrlPtx, &ctrlPty, &endPtx, &endPty)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath6quadToEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ctrlPtx, &ctrlPty, &endPtx, &endPty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:122
// index:0
// QPointF currentPosition()
func (this *QPainterPath) CurrentPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath15currentPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:124
// index:0
// void addRect(const class QRectF &)
func (this *QPainterPath) AddRect(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7addRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:125
// index:1
// inline
// void addRect(qreal, qreal, qreal, qreal)
func (this *QPainterPath) AddRect_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7addRectEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:126
// index:0
// void addEllipse(const class QRectF &)
func (this *QPainterPath) AddEllipse(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:127
// index:1
// inline
// void addEllipse(qreal, qreal, qreal, qreal)
func (this *QPainterPath) AddEllipse_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:128
// index:2
// inline
// void addEllipse(const class QPointF &, qreal, qreal)
func (this *QPainterPath) AddEllipse_2(center unsafe.Pointer, rx float64, ry float64) {
	// 2: (, center const QPointF &, rx qreal, ry qreal), (center, &rx, &ry)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseERK7QPointFdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), center, &rx, &ry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:129
// index:0
// void addPolygon(const class QPolygonF &)
func (this *QPainterPath) AddPolygon(polygon unsafe.Pointer) {
	// 0: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath10addPolygonERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:130
// index:0
// void addText(const class QPointF &, const class QFont &, const class QString &)
func (this *QPainterPath) AddText(point unsafe.Pointer, f unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, point const QPointF &, f const QFont &, text const QString &), (point, f, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7addTextERK7QPointFRK5QFontRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point, f, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:131
// index:1
// inline
// void addText(qreal, qreal, const class QFont &, const class QString &)
func (this *QPainterPath) AddText_1(x float64, y float64, f unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, x qreal, y qreal, f const QFont &, text const QString &), (&x, &y, f, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7addTextEddRK5QFontRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, f, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:132
// index:0
// void addPath(const class QPainterPath &)
func (this *QPainterPath) AddPath(path unsafe.Pointer) {
	// 0: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath7addPathERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:133
// index:0
// void addRegion(const class QRegion &)
func (this *QPainterPath) AddRegion(region unsafe.Pointer) {
	// 0: (, region const QRegion &), (region)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath9addRegionERK7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:135
// index:0
// void addRoundedRect(const class QRectF &, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect(rect unsafe.Pointer, xRadius float64, yRadius float64, mode int) {
	// 0: (, rect const QRectF &, xRadius qreal, yRadius qreal, mode Qt::SizeMode), (rect, &xRadius, &yRadius, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectERK6QRectFddN2Qt8SizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xRadius, &yRadius, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:137
// index:1
// inline
// void addRoundedRect(qreal, qreal, qreal, qreal, qreal, qreal, Qt::SizeMode)
func (this *QPainterPath) AddRoundedRect_1(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode int) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, xRadius qreal, yRadius qreal, mode Qt::SizeMode), (&x, &y, &w, &h, &xRadius, &yRadius, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectEddddddN2Qt8SizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &xRadius, &yRadius, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:141
// index:0
// void addRoundRect(const class QRectF &, int, int)
func (this *QPainterPath) AddRoundRect(rect unsafe.Pointer, xRnd int, yRnd int) {
	// 0: (, rect const QRectF &, xRnd int, yRnd int), (rect, &xRnd, &yRnd)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xRnd, &yRnd)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:142
// index:1
// inline
// void addRoundRect(qreal, qreal, qreal, qreal, int, int)
func (this *QPainterPath) AddRoundRect_1(x float64, y float64, w float64, h float64, xRnd int, yRnd int) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, xRnd int, yRnd int), (&x, &y, &w, &h, &xRnd, &yRnd)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectEddddii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &xRnd, &yRnd)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:144
// index:2
// inline
// void addRoundRect(const class QRectF &, int)
func (this *QPainterPath) AddRoundRect_2(rect unsafe.Pointer, roundness int) {
	// 2: (, rect const QRectF &, roundness int), (rect, &roundness)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectERK6QRectFi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &roundness)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:145
// index:3
// inline
// void addRoundRect(qreal, qreal, qreal, qreal, int)
func (this *QPainterPath) AddRoundRect_3(x float64, y float64, w float64, h float64, roundness int) {
	// 3: (, x qreal, y qreal, w qreal, h qreal, roundness int), (&x, &y, &w, &h, &roundness)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectEddddi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &roundness)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:148
// index:0
// void connectPath(const class QPainterPath &)
func (this *QPainterPath) ConnectPath(path unsafe.Pointer) {
	// 0: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath11connectPathERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:150
// index:0
// bool contains(const class QPointF &)
func (this *QPainterPath) Contains(pt unsafe.Pointer) {
	// 0: (, pt const QPointF &), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:151
// index:1
// bool contains(const class QRectF &)
func (this *QPainterPath) Contains_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath8containsERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:187
// index:2
// bool contains(const class QPainterPath &)
func (this *QPainterPath) Contains_2(p unsafe.Pointer) {
	// 2: (, p const QPainterPath &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath8containsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:152
// index:0
// bool intersects(const class QRectF &)
func (this *QPainterPath) Intersects(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10intersectsERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:186
// index:1
// bool intersects(const class QPainterPath &)
func (this *QPainterPath) Intersects_1(p unsafe.Pointer) {
	// 1: (, p const QPainterPath &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10intersectsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:154
// index:0
// void translate(qreal, qreal)
func (this *QPainterPath) Translate(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath9translateEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:155
// index:1
// inline
// void translate(const class QPointF &)
func (this *QPainterPath) Translate_1(offset unsafe.Pointer) {
	// 1: (, offset const QPointF &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath9translateERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:157
// index:0
// QPainterPath translated(qreal, qreal)
func (this *QPainterPath) Translated(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10translatedEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:158
// index:1
// inline
// QPainterPath translated(const class QPointF &)
func (this *QPainterPath) Translated_1(offset unsafe.Pointer) {
	// 1: (, offset const QPointF &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10translatedERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:160
// index:0
// QRectF boundingRect()
func (this *QPainterPath) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:161
// index:0
// QRectF controlPointRect()
func (this *QPainterPath) ControlPointRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath16controlPointRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:163
// index:0
// Qt::FillRule fillRule()
func (this *QPainterPath) FillRule() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath8fillRuleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:164
// index:0
// void setFillRule(Qt::FillRule)
func (this *QPainterPath) SetFillRule(fillRule int) {
	// 0: (, fillRule Qt::FillRule), (&fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath11setFillRuleEN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:166
// index:0
// bool isEmpty()
func (this *QPainterPath) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:168
// index:0
// QPainterPath toReversed()
func (this *QPainterPath) ToReversed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10toReversedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:169
// index:0
// QList<QPolygonF> toSubpathPolygons(const class QMatrix &)
func (this *QPainterPath) ToSubpathPolygons(matrix unsafe.Pointer) {
	// 0: (, matrix const QMatrix &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath17toSubpathPolygonsERK7QMatrix", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:172
// index:1
// QList<QPolygonF> toSubpathPolygons(const class QTransform &)
func (this *QPainterPath) ToSubpathPolygons_1(matrix unsafe.Pointer) {
	// 1: (, matrix const QTransform &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath17toSubpathPolygonsERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:170
// index:0
// QList<QPolygonF> toFillPolygons(const class QMatrix &)
func (this *QPainterPath) ToFillPolygons(matrix unsafe.Pointer) {
	// 0: (, matrix const QMatrix &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath14toFillPolygonsERK7QMatrix", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:173
// index:1
// QList<QPolygonF> toFillPolygons(const class QTransform &)
func (this *QPainterPath) ToFillPolygons_1(matrix unsafe.Pointer) {
	// 1: (, matrix const QTransform &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath14toFillPolygonsERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:171
// index:0
// QPolygonF toFillPolygon(const class QMatrix &)
func (this *QPainterPath) ToFillPolygon(matrix unsafe.Pointer) {
	// 0: (, matrix const QMatrix &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK7QMatrix", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:174
// index:1
// QPolygonF toFillPolygon(const class QTransform &)
func (this *QPainterPath) ToFillPolygon_1(matrix unsafe.Pointer) {
	// 1: (, matrix const QTransform &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath13toFillPolygonERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:176
// index:0
// int elementCount()
func (this *QPainterPath) ElementCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath12elementCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:177
// index:0
// QPainterPath::Element elementAt(int)
func (this *QPainterPath) ElementAt(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath9elementAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:178
// index:0
// void setElementPositionAt(int, qreal, qreal)
func (this *QPainterPath) SetElementPositionAt(i int, x float64, y float64) {
	// 0: (, i int, x qreal, y qreal), (&i, &x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPainterPath20setElementPositionAtEidd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:180
// index:0
// qreal length()
func (this *QPainterPath) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath6lengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:181
// index:0
// qreal percentAtLength(qreal)
func (this *QPainterPath) PercentAtLength(t float64) {
	// 0: (, t qreal), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath15percentAtLengthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:182
// index:0
// QPointF pointAtPercent(qreal)
func (this *QPainterPath) PointAtPercent(t float64) {
	// 0: (, t qreal), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath14pointAtPercentEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:183
// index:0
// qreal angleAtPercent(qreal)
func (this *QPainterPath) AngleAtPercent(t float64) {
	// 0: (, t qreal), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath14angleAtPercentEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:184
// index:0
// qreal slopeAtPercent(qreal)
func (this *QPainterPath) SlopeAtPercent(t float64) {
	// 0: (, t qreal), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath14slopeAtPercentEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:188
// index:0
// QPainterPath united(const class QPainterPath &)
func (this *QPainterPath) United(r unsafe.Pointer) {
	// 0: (, r const QPainterPath &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath6unitedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:189
// index:0
// QPainterPath intersected(const class QPainterPath &)
func (this *QPainterPath) Intersected(r unsafe.Pointer) {
	// 0: (, r const QPainterPath &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath11intersectedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:190
// index:0
// QPainterPath subtracted(const class QPainterPath &)
func (this *QPainterPath) Subtracted(r unsafe.Pointer) {
	// 0: (, r const QPainterPath &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10subtractedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:191
// index:0
// QPainterPath subtractedInverted(const class QPainterPath &)
func (this *QPainterPath) SubtractedInverted(r unsafe.Pointer) {
	// 0: (, r const QPainterPath &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath18subtractedInvertedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:193
// index:0
// QPainterPath simplified()
func (this *QPainterPath) Simplified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPainterPath10simplifiedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
