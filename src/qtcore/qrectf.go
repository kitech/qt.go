//  header block begin
// /usr/include/qt/QtCore/qrect.h
// #include <qrect.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QRectF struct {
	*qtrt.CObject
}

func (this *QRectF) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qrect.h:514
// index:0
// inline
// void QRectF()
func NewQRectF() *QRectF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}
func NewQRectFFromPointer(cthis unsafe.Pointer) *QRectF {
	return &QRectF{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qrect.h:515
// index:1
// inline
// void QRectF(const class QPointF &, const class QSizeF &)
func NewQRectF_1(topleft unsafe.Pointer, size unsafe.Pointer) *QRectF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2ERK7QPointFRK6QSizeF", ffiqt.FFI_TYPE_VOID, cthis, topleft, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:516
// index:2
// inline
// void QRectF(const class QPointF &, const class QPointF &)
func NewQRectF_2(topleft unsafe.Pointer, bottomRight unsafe.Pointer) *QRectF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2ERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, cthis, topleft, bottomRight)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:517
// index:3
// inline
// void QRectF(qreal, qreal, qreal, qreal)
func NewQRectF_3(left float64, top float64, width float64, height float64) *QRectF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2Edddd", ffiqt.FFI_TYPE_VOID, cthis, &left, &top, &width, &height)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:518
// index:4
// inline
// void QRectF(const class QRect &)
func NewQRectF_4(rect unsafe.Pointer) *QRectF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2ERK5QRect", ffiqt.FFI_TYPE_VOID, cthis, rect)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:520
// index:0
// inline
// bool isNull()
func (this *QRectF) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:521
// index:0
// inline
// bool isEmpty()
func (this *QRectF) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:522
// index:0
// inline
// bool isValid()
func (this *QRectF) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:523
// index:0
// QRectF normalized()
func (this *QRectF) Normalized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10normalizedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:525
// index:0
// inline
// qreal left()
func (this *QRectF) Left() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF4leftEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:526
// index:0
// inline
// qreal top()
func (this *QRectF) Top() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF3topEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:527
// index:0
// inline
// qreal right()
func (this *QRectF) Right() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF5rightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:528
// index:0
// inline
// qreal bottom()
func (this *QRectF) Bottom() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6bottomEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:530
// index:0
// inline
// qreal x()
func (this *QRectF) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF1xEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:531
// index:0
// inline
// qreal y()
func (this *QRectF) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF1yEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:532
// index:0
// inline
// void setLeft(qreal)
func (this *QRectF) SetLeft(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7setLeftEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:533
// index:0
// inline
// void setTop(qreal)
func (this *QRectF) SetTop(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6setTopEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:534
// index:0
// inline
// void setRight(qreal)
func (this *QRectF) SetRight(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF8setRightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:535
// index:0
// inline
// void setBottom(qreal)
func (this *QRectF) SetBottom(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9setBottomEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:536
// index:0
// inline
// void setX(qreal)
func (this *QRectF) SetX(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF4setXEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:537
// index:0
// inline
// void setY(qreal)
func (this *QRectF) SetY(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF4setYEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:539
// index:0
// inline
// QPointF topLeft()
func (this *QRectF) TopLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7topLeftEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:540
// index:0
// inline
// QPointF bottomRight()
func (this *QRectF) BottomRight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF11bottomRightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:541
// index:0
// inline
// QPointF topRight()
func (this *QRectF) TopRight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8topRightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:542
// index:0
// inline
// QPointF bottomLeft()
func (this *QRectF) BottomLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10bottomLeftEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:543
// index:0
// inline
// QPointF center()
func (this *QRectF) Center() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6centerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:545
// index:0
// inline
// void setTopLeft(const class QPointF &)
func (this *QRectF) SetTopLeft(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF10setTopLeftERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:546
// index:0
// inline
// void setBottomRight(const class QPointF &)
func (this *QRectF) SetBottomRight(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF14setBottomRightERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:547
// index:0
// inline
// void setTopRight(const class QPointF &)
func (this *QRectF) SetTopRight(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF11setTopRightERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:548
// index:0
// inline
// void setBottomLeft(const class QPointF &)
func (this *QRectF) SetBottomLeft(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF13setBottomLeftERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:550
// index:0
// inline
// void moveLeft(qreal)
func (this *QRectF) MoveLeft(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF8moveLeftEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:551
// index:0
// inline
// void moveTop(qreal)
func (this *QRectF) MoveTop(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7moveTopEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:552
// index:0
// inline
// void moveRight(qreal)
func (this *QRectF) MoveRight(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9moveRightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:553
// index:0
// inline
// void moveBottom(qreal)
func (this *QRectF) MoveBottom(pos float64) {
	// 0: (, pos qreal), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF10moveBottomEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:554
// index:0
// inline
// void moveTopLeft(const class QPointF &)
func (this *QRectF) MoveTopLeft(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF11moveTopLeftERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:555
// index:0
// inline
// void moveBottomRight(const class QPointF &)
func (this *QRectF) MoveBottomRight(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF15moveBottomRightERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:556
// index:0
// inline
// void moveTopRight(const class QPointF &)
func (this *QRectF) MoveTopRight(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF12moveTopRightERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:557
// index:0
// inline
// void moveBottomLeft(const class QPointF &)
func (this *QRectF) MoveBottomLeft(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF14moveBottomLeftERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:558
// index:0
// inline
// void moveCenter(const class QPointF &)
func (this *QRectF) MoveCenter(p unsafe.Pointer) {
	// 0: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF10moveCenterERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:560
// index:0
// inline
// void translate(qreal, qreal)
func (this *QRectF) Translate(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9translateEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:561
// index:1
// inline
// void translate(const class QPointF &)
func (this *QRectF) Translate_1(p unsafe.Pointer) {
	// 1: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9translateERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:563
// index:0
// inline
// QRectF translated(qreal, qreal)
func (this *QRectF) Translated(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10translatedEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:564
// index:1
// inline
// QRectF translated(const class QPointF &)
func (this *QRectF) Translated_1(p unsafe.Pointer) {
	// 1: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10translatedERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:566
// index:0
// inline
// QRectF transposed()
func (this *QRectF) Transposed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10transposedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:568
// index:0
// inline
// void moveTo(qreal, qreal)
func (this *QRectF) MoveTo(x float64, y float64) {
	// 0: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6moveToEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:569
// index:1
// inline
// void moveTo(const class QPointF &)
func (this *QRectF) MoveTo_1(p unsafe.Pointer) {
	// 1: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6moveToERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:571
// index:0
// inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QRectF) SetRect(x float64, y float64, w float64, h float64) {
	// 0: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7setRectEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:572
// index:0
// inline
// void getRect(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) GetRect(x unsafe.Pointer, y unsafe.Pointer, w unsafe.Pointer, h unsafe.Pointer) {
	// 0: (, x qreal *, y qreal *, w qreal *, h qreal *), (x, y, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7getRectEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:574
// index:0
// inline
// void setCoords(qreal, qreal, qreal, qreal)
func (this *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	// 0: (, x1 qreal, y1 qreal, x2 qreal, y2 qreal), (&x1, &y1, &x2, &y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9setCoordsEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:575
// index:0
// inline
// void getCoords(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) GetCoords(x1 unsafe.Pointer, y1 unsafe.Pointer, x2 unsafe.Pointer, y2 unsafe.Pointer) {
	// 0: (, x1 qreal *, y1 qreal *, x2 qreal *, y2 qreal *), (x1, y1, x2, y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF9getCoordsEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:577
// index:0
// inline
// void adjust(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjust(x1 float64, y1 float64, x2 float64, y2 float64) {
	// 0: (, x1 qreal, y1 qreal, x2 qreal, y2 qreal), (&x1, &y1, &x2, &y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6adjustEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:578
// index:0
// inline
// QRectF adjusted(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjusted(x1 float64, y1 float64, x2 float64, y2 float64) {
	// 0: (, x1 qreal, y1 qreal, x2 qreal, y2 qreal), (&x1, &y1, &x2, &y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8adjustedEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:580
// index:0
// inline
// QSizeF size()
func (this *QRectF) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:581
// index:0
// inline
// qreal width()
func (this *QRectF) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:582
// index:0
// inline
// qreal height()
func (this *QRectF) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:583
// index:0
// inline
// void setWidth(qreal)
func (this *QRectF) SetWidth(w float64) {
	// 0: (, w qreal), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF8setWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:584
// index:0
// inline
// void setHeight(qreal)
func (this *QRectF) SetHeight(h float64) {
	// 0: (, h qreal), (&h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9setHeightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:585
// index:0
// inline
// void setSize(const class QSizeF &)
func (this *QRectF) SetSize(s unsafe.Pointer) {
	// 0: (, s const QSizeF &), (s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7setSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:592
// index:0
// bool contains(const class QRectF &)
func (this *QRectF) Contains(r unsafe.Pointer) {
	// 0: (, r const QRectF &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8containsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:593
// index:1
// bool contains(const class QPointF &)
func (this *QRectF) Contains_1(p unsafe.Pointer) {
	// 1: (, p const QPointF &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:594
// index:2
// inline
// bool contains(qreal, qreal)
func (this *QRectF) Contains_2(x float64, y float64) {
	// 2: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8containsEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:595
// index:0
// inline
// QRectF united(const class QRectF &)
func (this *QRectF) United(other unsafe.Pointer) {
	// 0: (, other const QRectF &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6unitedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:596
// index:0
// inline
// QRectF intersected(const class QRectF &)
func (this *QRectF) Intersected(other unsafe.Pointer) {
	// 0: (, other const QRectF &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF11intersectedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:597
// index:0
// bool intersects(const class QRectF &)
func (this *QRectF) Intersects(r unsafe.Pointer) {
	// 0: (, r const QRectF &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10intersectsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:599
// index:0
// inline
// QRectF marginsAdded(const class QMarginsF &)
func (this *QRectF) MarginsAdded(margins unsafe.Pointer) {
	// 0: (, margins const QMarginsF &), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF12marginsAddedERK9QMarginsF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:600
// index:0
// inline
// QRectF marginsRemoved(const class QMarginsF &)
func (this *QRectF) MarginsRemoved(margins unsafe.Pointer) {
	// 0: (, margins const QMarginsF &), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF14marginsRemovedERK9QMarginsF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:612
// index:0
// inline
// QRect toRect()
func (this *QRectF) ToRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6toRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:613
// index:0
// QRect toAlignedRect()
func (this *QRectF) ToAlignedRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF13toAlignedRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
