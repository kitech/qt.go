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

/*

 */
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

/*
Constructs an empty QPainterPath object.
*/
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

/*
Constructs an empty QPainterPath object.
*/
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

/*

 */
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

/*

 */
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

/*

 */
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

/*
Swaps painter path other with this painter path. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
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

/*
Closes the current subpath by drawing a line to the beginning of the subpath, automatically starting a new path. The current point of the new path is (0, 0).

If the subpath does not contain any elements, this function does nothing.

See also moveTo() and Composing a QPainterPath.
*/
func (this *QPainterPath) CloseSubpath() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12closeSubpathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveTo(const QPointF &)

/*
Moves the current point to the given point, implicitly starting a new subpath and closing the previous one.

See also closeSubpath() and Composing a QPainterPath.
*/
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

/*
Moves the current point to the given point, implicitly starting a new subpath and closing the previous one.

See also closeSubpath() and Composing a QPainterPath.
*/
func (this *QPainterPath) MoveTo_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6moveToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lineTo(const QPointF &)

/*
Adds a straight line from the current position to the given endPoint. After the line is drawn, the current position is updated to be at the end point of the line.

See also addPolygon(), addRect(), and Composing a QPainterPath.
*/
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

/*
Adds a straight line from the current position to the given endPoint. After the line is drawn, the current position is updated to be at the end point of the line.

See also addPolygon(), addRect(), and Composing a QPainterPath.
*/
func (this *QPainterPath) LineTo_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6lineToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void arcMoveTo(const QRectF &, qreal)

/*
Creates a move to that lies on the arc that occupies the given rectangle at angle.

Angles are specified in degrees. Clockwise arcs can be specified using negative angles.

This function was introduced in  Qt 4.2.

See also moveTo() and arcTo().
*/
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

/*
Creates a move to that lies on the arc that occupies the given rectangle at angle.

Angles are specified in degrees. Clockwise arcs can be specified using negative angles.

This function was introduced in  Qt 4.2.

See also moveTo() and arcTo().
*/
func (this *QPainterPath) ArcMoveTo_1(x float64, y float64, w float64, h float64, angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9arcMoveToEddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void arcTo(const QRectF &, qreal, qreal)

/*
Creates an arc that occupies the given rectangle, beginning at the specified startAngle and extending sweepLength degrees counter-clockwise.

Angles are specified in degrees. Clockwise arcs can be specified using negative angles.

Note that this function connects the starting point of the arc to the current position if they are not already connected. After the arc has been added, the current position is the last point in arc. To draw a line back to the first point, use the closeSubpath() function.



  QLinearGradient myGradient;
  QPen myPen;

  QPointF center, startPoint;

  QPainterPath myPath;
  myPath.moveTo(center);
  myPath.arcTo(boundingRect, startAngle,
               sweepLength);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also arcMoveTo(), addEllipse(), QPainter::drawArc(), QPainter::drawPie(), and Composing a QPainterPath.
*/
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

/*
Creates an arc that occupies the given rectangle, beginning at the specified startAngle and extending sweepLength degrees counter-clockwise.

Angles are specified in degrees. Clockwise arcs can be specified using negative angles.

Note that this function connects the starting point of the arc to the current position if they are not already connected. After the arc has been added, the current position is the last point in arc. To draw a line back to the first point, use the closeSubpath() function.



  QLinearGradient myGradient;
  QPen myPen;

  QPointF center, startPoint;

  QPainterPath myPath;
  myPath.moveTo(center);
  myPath.arcTo(boundingRect, startAngle,
               sweepLength);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also arcMoveTo(), addEllipse(), QPainter::drawArc(), QPainter::drawPie(), and Composing a QPainterPath.
*/
func (this *QPainterPath) ArcTo_1(x float64, y float64, w float64, h float64, startAngle float64, arcLength float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath5arcToEdddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, startAngle, arcLength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cubicTo(const QPointF &, const QPointF &, const QPointF &)

/*
Adds a cubic Bezier curve between the current position and the given endPoint using the control points specified by c1, and c2.

After the curve is added, the current position is updated to be at the end point of the curve.



  QLinearGradient myGradient;
  QPen myPen;

  QPainterPath myPath;
  myPath.cubicTo(c1, c2, endPoint);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also quadTo() and Composing a QPainterPath.
*/
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

/*
Adds a cubic Bezier curve between the current position and the given endPoint using the control points specified by c1, and c2.

After the curve is added, the current position is updated to be at the end point of the curve.



  QLinearGradient myGradient;
  QPen myPen;

  QPainterPath myPath;
  myPath.cubicTo(c1, c2, endPoint);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also quadTo() and Composing a QPainterPath.
*/
func (this *QPainterPath) CubicTo_1(ctrlPt1x float64, ctrlPt1y float64, ctrlPt2x float64, ctrlPt2y float64, endPtx float64, endPty float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7cubicToEdddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ctrlPt1x, ctrlPt1y, ctrlPt2x, ctrlPt2y, endPtx, endPty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quadTo(const QPointF &, const QPointF &)

/*
Adds a quadratic Bezier curve between the current position and the given endPoint with the control point specified by c.

After the curve is added, the current point is updated to be at the end point of the curve.

See also cubicTo() and Composing a QPainterPath.
*/
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

/*
Adds a quadratic Bezier curve between the current position and the given endPoint with the control point specified by c.

After the curve is added, the current point is updated to be at the end point of the curve.

See also cubicTo() and Composing a QPainterPath.
*/
func (this *QPainterPath) QuadTo_1(ctrlPtx float64, ctrlPty float64, endPtx float64, endPty float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath6quadToEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ctrlPtx, ctrlPty, endPtx, endPty)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:122
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF currentPosition() const

/*
Returns the current position of the path.
*/
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

/*
Adds the given rectangle to this path as a closed subpath.

The rectangle is added as a clockwise set of lines. The painter path's current position after the rectangle has been added is at the top-left corner of the rectangle.



  QLinearGradient myGradient;
  QPen myPen;
  QRectF myRectangle;

  QPainterPath myPath;
  myPath.addRect(myRectangle);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also addRegion(), lineTo(), and Composing a QPainterPath.
*/
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

/*
Adds the given rectangle to this path as a closed subpath.

The rectangle is added as a clockwise set of lines. The painter path's current position after the rectangle has been added is at the top-left corner of the rectangle.



  QLinearGradient myGradient;
  QPen myPen;
  QRectF myRectangle;

  QPainterPath myPath;
  myPath.addRect(myRectangle);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also addRegion(), lineTo(), and Composing a QPainterPath.
*/
func (this *QPainterPath) AddRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath7addRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addEllipse(const QRectF &)

/*
Creates an ellipse within the specified boundingRectangle and adds it to the painter path as a closed subpath.

The ellipse is composed of a clockwise curve, starting and finishing at zero degrees (the 3 o'clock position).



  QLinearGradient myGradient;
  QPen myPen;
  QRectF boundingRectangle;

  QPainterPath myPath;
  myPath.addEllipse(boundingRectangle);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also arcTo(), QPainter::drawEllipse(), and Composing a QPainterPath.
*/
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

/*
Creates an ellipse within the specified boundingRectangle and adds it to the painter path as a closed subpath.

The ellipse is composed of a clockwise curve, starting and finishing at zero degrees (the 3 o'clock position).



  QLinearGradient myGradient;
  QPen myPen;
  QRectF boundingRectangle;

  QPainterPath myPath;
  myPath.addEllipse(boundingRectangle);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also arcTo(), QPainter::drawEllipse(), and Composing a QPainterPath.
*/
func (this *QPainterPath) AddEllipse_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath10addEllipseEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:128
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void addEllipse(const QPointF &, qreal, qreal)

/*
Creates an ellipse within the specified boundingRectangle and adds it to the painter path as a closed subpath.

The ellipse is composed of a clockwise curve, starting and finishing at zero degrees (the 3 o'clock position).



  QLinearGradient myGradient;
  QPen myPen;
  QRectF boundingRectangle;

  QPainterPath myPath;
  myPath.addEllipse(boundingRectangle);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also arcTo(), QPainter::drawEllipse(), and Composing a QPainterPath.
*/
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

/*
Adds the given polygon to the path as an (unclosed) subpath.

Note that the current position after the polygon has been added, is the last point in polygon. To draw a line back to the first point, use the closeSubpath() function.



  QLinearGradient myGradient;
  QPen myPen;
  QPolygonF myPolygon;

  QPainterPath myPath;
  myPath.addPolygon(myPolygon);

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also lineTo() and Composing a QPainterPath.
*/
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

/*
Adds the given text to this path as a set of closed subpaths created from the font supplied. The subpaths are positioned so that the left end of the text's baseline lies at the specified point.



  QLinearGradient myGradient;
  QPen myPen;
  QFont myFont;
  QPointF baseline(x, y);

  QPainterPath myPath;
  myPath.addText(baseline, myFont, tr("Qt"));

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also QPainter::drawText() and Composing a QPainterPath.
*/
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

/*
Adds the given text to this path as a set of closed subpaths created from the font supplied. The subpaths are positioned so that the left end of the text's baseline lies at the specified point.



  QLinearGradient myGradient;
  QPen myPen;
  QFont myFont;
  QPointF baseline(x, y);

  QPainterPath myPath;
  myPath.addText(baseline, myFont, tr("Qt"));

  QPainter painter(this);
  painter.setBrush(myGradient);
  painter.setPen(myPen);
  painter.drawPath(myPath);





See also QPainter::drawText() and Composing a QPainterPath.
*/
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

/*
Adds the given path to this path as a closed subpath.

See also connectPath() and Composing a QPainterPath.
*/
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

/*
Adds the given region to the path by adding each rectangle in the region as a separate closed subpath.

See also addRect() and Composing a QPainterPath.
*/
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

/*
Adds the given rectangle rect with rounded corners to the path.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

This function was introduced in  Qt 4.4.

See also addRect().
*/
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

/*
Adds the given rectangle rect with rounded corners to the path.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

This function was introduced in  Qt 4.4.

See also addRect().
*/
func (this *QPainterPath) AddRoundedRect__(rect qtcore.QRectF_ITF, xRadius float64, yRadius float64) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 3, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectERK6QRectFddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundedRect(qreal, qreal, qreal, qreal, qreal, qreal, Qt::SizeMode)

/*
Adds the given rectangle rect with rounded corners to the path.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

This function was introduced in  Qt 4.4.

See also addRect().
*/
func (this *QPainterPath) AddRoundedRect_1(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectEddddddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundedRect(qreal, qreal, qreal, qreal, qreal, qreal, Qt::SizeMode)

/*
Adds the given rectangle rect with rounded corners to the path.

The xRadius and yRadius arguments specify the radii of the ellipses defining the corners of the rounded rectangle. When mode is Qt::RelativeSize, xRadius and yRadius are specified in percentage of half the rectangle's width and height respectively, and should be in the range 0.0 to 100.0.

This function was introduced in  Qt 4.4.

See also addRect().
*/
func (this *QPainterPath) AddRoundedRect_1_(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64) {
	// arg: 6, Qt::SizeMode=Elaborated, Qt::SizeMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath14addRoundedRectEddddddN2Qt8SizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRadius, yRadius, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addRoundRect(const QRectF &, int, int)

/*

 */
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

/*

 */
func (this *QPainterPath) AddRoundRect_1(x float64, y float64, w float64, h float64, xRnd int, yRnd int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xRnd, yRnd)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void addRoundRect(const QRectF &, int)

/*

 */
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

/*

 */
func (this *QPainterPath) AddRoundRect_3(x float64, y float64, w float64, h float64, roundness int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath12addRoundRectEddddi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, roundness)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectPath(const QPainterPath &)

/*
Connects the given path to this path by adding a line from the last element of this path to the first element of the given path.

See also addPath() and Composing a QPainterPath.
*/
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

/*
Returns true if the given point is inside the path, otherwise returns false.

See also intersects().
*/
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

/*
Returns true if the given point is inside the path, otherwise returns false.

See also intersects().
*/
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

/*
Returns true if the given point is inside the path, otherwise returns false.

See also intersects().
*/
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

/*
Returns true if any point in the given rectangle intersects the path; otherwise returns false.

There is an intersection if any of the lines making up the rectangle crosses a part of the path or if any part of the rectangle overlaps with any area enclosed by the path. This function respects the current fillRule to determine what is considered inside the path.

See also contains().
*/
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

/*
Returns true if any point in the given rectangle intersects the path; otherwise returns false.

There is an intersection if any of the lines making up the rectangle crosses a part of the path or if any part of the rectangle overlaps with any area enclosed by the path. This function respects the current fillRule to determine what is considered inside the path.

See also contains().
*/
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

/*
Translates all elements in the path by (dx, dy).

This function was introduced in  Qt 4.6.

See also translated().
*/
func (this *QPainterPath) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:155
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)

/*
Translates all elements in the path by (dx, dy).

This function was introduced in  Qt 4.6.

See also translated().
*/
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

/*
Returns a copy of the path that is translated by (dx, dy).

This function was introduced in  Qt 4.6.

See also translate().
*/
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

/*
Returns a copy of the path that is translated by (dx, dy).

This function was introduced in  Qt 4.6.

See also translate().
*/
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

/*
Returns the bounding rectangle of this painter path as a rectangle with floating point precision.

See also controlPointRect().
*/
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

/*
Returns the rectangle containing all the points and control points in this path.

This function is significantly faster to compute than the exact boundingRect(), and the returned rectangle is always a superset of the rectangle returned by boundingRect().

See also boundingRect().
*/
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

/*
Returns the painter path's currently set fill rule.

See also setFillRule().
*/
func (this *QPainterPath) FillRule() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath8fillRuleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFillRule(Qt::FillRule)

/*
Sets the fill rule of the painter path to the given fillRule. Qt provides two methods for filling paths:


 Qt::OddEvenFill (default)Qt::WindingFill



See also fillRule().
*/
func (this *QPainterPath) SetFillRule(fillRule int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath11setFillRuleEN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fillRule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if either there are no elements in this path, or if the only element is a MoveToElement; otherwise returns false.

See also elementCount().
*/
func (this *QPainterPath) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpainterpath.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath toReversed() const

/*
Creates and returns a reversed copy of the path.

It is the order of the elements that is reversed: If a QPainterPath is composed by calling the moveTo(), lineTo() and cubicTo() functions in the specified order, the reversed copy is composed by calling cubicTo(), lineTo() and moveTo().
*/
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

/*
Converts the path into a polygon using the QTransform matrix, and returns the polygon.

The polygon is created by first converting all subpaths to polygons, then using a rewinding technique to make sure that overlapping subpaths can be filled using the correct fill rule.

Note that rewinding inserts addition lines in the polygon so the outline of the fill polygon does not match the outline of the path.

See also toSubpathPolygons(), toFillPolygons(), and QPainterPath Conversion.
*/
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

/*
Converts the path into a polygon using the QTransform matrix, and returns the polygon.

The polygon is created by first converting all subpaths to polygons, then using a rewinding technique to make sure that overlapping subpaths can be filled using the correct fill rule.

Note that rewinding inserts addition lines in the polygon so the outline of the fill polygon does not match the outline of the path.

See also toSubpathPolygons(), toFillPolygons(), and QPainterPath Conversion.
*/
func (this *QPainterPath) ToFillPolygon__() *QPolygonF /*123*/ {
	// arg: 0, const QMatrix &=LValueReference, QMatrix=Record, , Invalid
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

/*
Converts the path into a polygon using the QTransform matrix, and returns the polygon.

The polygon is created by first converting all subpaths to polygons, then using a rewinding technique to make sure that overlapping subpaths can be filled using the correct fill rule.

Note that rewinding inserts addition lines in the polygon so the outline of the fill polygon does not match the outline of the path.

See also toSubpathPolygons(), toFillPolygons(), and QPainterPath Conversion.
*/
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

/*
Returns the number of path elements in the painter path.

See also ElementType, elementAt(), and isEmpty().
*/
func (this *QPainterPath) ElementCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath12elementCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:177
// index:0
// Public Visibility=Default Availability=Available
// [24] QPainterPath::Element elementAt(int) const

/*
Returns the element at the given index in the painter path.

See also ElementType, elementCount(), and isEmpty().
*/
func (this *QPainterPath) ElementAt(i int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath9elementAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElementPositionAt(int, qreal, qreal)

/*
Sets the x and y coordinate of the element at index index to x and y.

This function was introduced in  Qt 4.2.
*/
func (this *QPainterPath) SetElementPositionAt(i int, x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPainterPath20setElementPositionAtEidd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:180
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal length() const

/*
Returns the length of the current path.
*/
func (this *QPainterPath) Length() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal percentAtLength(qreal) const

/*
Returns percentage of the whole path at the specified length len.

Note that similarly to other percent methods, the percentage measurement is not linear with regards to the length, if curves are present in the path. When curves are present the percentage argument is mapped to the t parameter of the Bezier equations.
*/
func (this *QPainterPath) PercentAtLength(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath15percentAtLengthEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:182
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pointAtPercent(qreal) const

/*
Returns the point at at the percentage t of the current path. The argument t has to be between 0 and 1.

Note that similarly to other percent methods, the percentage measurement is not linear with regards to the length, if curves are present in the path. When curves are present the percentage argument is mapped to the t parameter of the Bezier equations.
*/
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

/*
Returns the angle of the path tangent at the percentage t. The argument t has to be between 0 and 1.

Positive values for the angles mean counter-clockwise while negative values mean the clockwise direction. Zero degrees is at the 3 o'clock position.

Note that similarly to the other percent methods, the percentage measurement is not linear with regards to the length if curves are present in the path. When curves are present the percentage argument is mapped to the t parameter of the Bezier equations.
*/
func (this *QPainterPath) AngleAtPercent(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14angleAtPercentEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal slopeAtPercent(qreal) const

/*
Returns the slope of the path at the percentage t. The argument t has to be between 0 and 1.

Note that similarly to other percent methods, the percentage measurement is not linear with regards to the length, if curves are present in the path. When curves are present the percentage argument is mapped to the t parameter of the Bezier equations.
*/
func (this *QPainterPath) SlopeAtPercent(t float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPainterPath14slopeAtPercentEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpainterpath.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath united(const QPainterPath &) const

/*
Returns a path which is the union of this path's fill area and p's fill area.

Set operations on paths will treat the paths as areas. Non-closed paths will be treated as implicitly closed. Bezier curves may be flattened to line segments due to numerical instability of doing bezier curve intersections.

This function was introduced in  Qt 4.3.

See also intersected() and subtracted().
*/
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

/*
Returns a path which is the intersection of this path's fill area and p's fill area. Bezier curves may be flattened to line segments due to numerical instability of doing bezier curve intersections.

This function was introduced in  Qt 4.3.
*/
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

/*
Returns a path which is p's fill area subtracted from this path's fill area.

Set operations on paths will treat the paths as areas. Non-closed paths will be treated as implicitly closed. Bezier curves may be flattened to line segments due to numerical instability of doing bezier curve intersections.

This function was introduced in  Qt 4.3.
*/
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

/*

 */
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

/*
Returns a simplified version of this path. This implies merging all subpaths that intersect, and returning a path containing no intersecting edges. Consecutive parallel lines will also be merged. The simplified path will always use the default fill rule, Qt::OddEvenFill. Bezier curves may be flattened to line segments due to numerical instability of doing bezier curve intersections.

This function was introduced in  Qt 4.4.
*/
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*
Subtracts the other path from a copy of this path, and returns the copy.

This function was introduced in  Qt 4.5.

See also subtracted(), operator-=(), and operator+().
*/
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*
This enum describes the types of elements used to connect vertices in subpaths.

Note that elements added as closed subpaths using the addEllipse(), addPath(), addPolygon(), addRect(), addRegion() and addText() convenience functions, is actually added to the path as a collection of separate elements using the moveTo(), lineTo() and cubicTo() functions.



See also elementAt() and elementCount().

*/
type QPainterPath__ElementType = int

// A new subpath. See also moveTo().
const QPainterPath__MoveToElement QPainterPath__ElementType = 0

// A line. See also lineTo().
const QPainterPath__LineToElement QPainterPath__ElementType = 1

// A curve. See also cubicTo() and quadTo().
const QPainterPath__CurveToElement QPainterPath__ElementType = 2

// The extra data required to describe a curve in a CurveToElement element.
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
