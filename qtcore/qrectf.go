package qtcore

// /usr/include/qt/QtCore/qrect.h
// #include <qrect.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRectF) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQRectFFromPointer(cthis unsafe.Pointer) *QRectF {
	return &QRectF{&qtrt.CObject{cthis}}
}
func (*QRectF) NewFromPointer(cthis unsafe.Pointer) *QRectF {
	return NewQRectFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrect.h:514
// index:0
// Public inline
// void QRectF()
func NewQRectF() *QRectF {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:515
// index:1
// Public inline
// void QRectF(const class QPointF &, const class QSizeF &)
func NewQRectF_1(topleft *QPointF, size *QSizeF) *QRectF {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = topleft.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2ERK7QPointFRK6QSizeF", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:516
// index:2
// Public inline
// void QRectF(const class QPointF &, const class QPointF &)
func NewQRectF_2(topleft *QPointF, bottomRight *QPointF) *QRectF {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = topleft.GetCthis()
	var convArg1 = bottomRight.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2ERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:517
// index:3
// Public inline
// void QRectF(qreal, qreal, qreal, qreal)
func NewQRectF_3(left float64, top float64, width float64, height float64) *QRectF {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2Edddd", ffiqt.FFI_TYPE_VOID, cthis, left, top, width, height)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:518
// index:4
// Public inline
// void QRectF(const class QRect &)
func NewQRectF_4(rect *QRect) *QRectF {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectFC2ERK5QRect", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:520
// index:0
// Public inline
// bool isNull()
func (this *QRectF) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:521
// index:0
// Public inline
// bool isEmpty()
func (this *QRectF) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:522
// index:0
// Public inline
// bool isValid()
func (this *QRectF) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:523
// index:0
// Public
// QRectF normalized()
func (this *QRectF) Normalized() *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10normalizedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:525
// index:0
// Public inline
// qreal left()
func (this *QRectF) Left() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF4leftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:526
// index:0
// Public inline
// qreal top()
func (this *QRectF) Top() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF3topEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:527
// index:0
// Public inline
// qreal right()
func (this *QRectF) Right() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF5rightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:528
// index:0
// Public inline
// qreal bottom()
func (this *QRectF) Bottom() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6bottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:530
// index:0
// Public inline
// qreal x()
func (this *QRectF) X() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:531
// index:0
// Public inline
// qreal y()
func (this *QRectF) Y() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:532
// index:0
// Public inline
// void setLeft(qreal)
func (this *QRectF) SetLeft(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7setLeftEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:533
// index:0
// Public inline
// void setTop(qreal)
func (this *QRectF) SetTop(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6setTopEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:534
// index:0
// Public inline
// void setRight(qreal)
func (this *QRectF) SetRight(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF8setRightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:535
// index:0
// Public inline
// void setBottom(qreal)
func (this *QRectF) SetBottom(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9setBottomEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:536
// index:0
// Public inline
// void setX(qreal)
func (this *QRectF) SetX(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF4setXEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:537
// index:0
// Public inline
// void setY(qreal)
func (this *QRectF) SetY(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF4setYEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:539
// index:0
// Public inline
// QPointF topLeft()
func (this *QRectF) TopLeft() *QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7topLeftEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:540
// index:0
// Public inline
// QPointF bottomRight()
func (this *QRectF) BottomRight() *QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF11bottomRightEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:541
// index:0
// Public inline
// QPointF topRight()
func (this *QRectF) TopRight() *QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8topRightEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:542
// index:0
// Public inline
// QPointF bottomLeft()
func (this *QRectF) BottomLeft() *QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10bottomLeftEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:543
// index:0
// Public inline
// QPointF center()
func (this *QRectF) Center() *QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6centerEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:545
// index:0
// Public inline
// void setTopLeft(const class QPointF &)
func (this *QRectF) SetTopLeft(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF10setTopLeftERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:546
// index:0
// Public inline
// void setBottomRight(const class QPointF &)
func (this *QRectF) SetBottomRight(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF14setBottomRightERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:547
// index:0
// Public inline
// void setTopRight(const class QPointF &)
func (this *QRectF) SetTopRight(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF11setTopRightERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:548
// index:0
// Public inline
// void setBottomLeft(const class QPointF &)
func (this *QRectF) SetBottomLeft(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF13setBottomLeftERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:550
// index:0
// Public inline
// void moveLeft(qreal)
func (this *QRectF) MoveLeft(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF8moveLeftEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:551
// index:0
// Public inline
// void moveTop(qreal)
func (this *QRectF) MoveTop(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7moveTopEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:552
// index:0
// Public inline
// void moveRight(qreal)
func (this *QRectF) MoveRight(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9moveRightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:553
// index:0
// Public inline
// void moveBottom(qreal)
func (this *QRectF) MoveBottom(pos float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF10moveBottomEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:554
// index:0
// Public inline
// void moveTopLeft(const class QPointF &)
func (this *QRectF) MoveTopLeft(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF11moveTopLeftERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:555
// index:0
// Public inline
// void moveBottomRight(const class QPointF &)
func (this *QRectF) MoveBottomRight(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF15moveBottomRightERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:556
// index:0
// Public inline
// void moveTopRight(const class QPointF &)
func (this *QRectF) MoveTopRight(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF12moveTopRightERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:557
// index:0
// Public inline
// void moveBottomLeft(const class QPointF &)
func (this *QRectF) MoveBottomLeft(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF14moveBottomLeftERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:558
// index:0
// Public inline
// void moveCenter(const class QPointF &)
func (this *QRectF) MoveCenter(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF10moveCenterERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:560
// index:0
// Public inline
// void translate(qreal, qreal)
func (this *QRectF) Translate(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9translateEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:561
// index:1
// Public inline
// void translate(const class QPointF &)
func (this *QRectF) Translate_1(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9translateERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:563
// index:0
// Public inline
// QRectF translated(qreal, qreal)
func (this *QRectF) Translated(dx float64, dy float64) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10translatedEdd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:564
// index:1
// Public inline
// QRectF translated(const class QPointF &)
func (this *QRectF) Translated_1(p *QPointF) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10translatedERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:566
// index:0
// Public inline
// QRectF transposed()
func (this *QRectF) Transposed() *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10transposedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:568
// index:0
// Public inline
// void moveTo(qreal, qreal)
func (this *QRectF) MoveTo(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6moveToEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:569
// index:1
// Public inline
// void moveTo(const class QPointF &)
func (this *QRectF) MoveTo_1(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6moveToERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:571
// index:0
// Public inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QRectF) SetRect(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:572
// index:0
// Public inline
// void getRect(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) GetRect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF7getRectEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:574
// index:0
// Public inline
// void setCoords(qreal, qreal, qreal, qreal)
func (this *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9setCoordsEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:575
// index:0
// Public inline
// void getCoords(qreal *, qreal *, qreal *, qreal *)
func (this *QRectF) GetCoords(x1 unsafe.Pointer /*666*/, y1 unsafe.Pointer /*666*/, x2 unsafe.Pointer /*666*/, y2 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF9getCoordsEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:577
// index:0
// Public inline
// void adjust(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjust(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF6adjustEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:578
// index:0
// Public inline
// QRectF adjusted(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjusted(x1 float64, y1 float64, x2 float64, y2 float64) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8adjustedEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:580
// index:0
// Public inline
// QSizeF size()
func (this *QRectF) Size() *QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:581
// index:0
// Public inline
// qreal width()
func (this *QRectF) Width() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:582
// index:0
// Public inline
// qreal height()
func (this *QRectF) Height() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qrect.h:583
// index:0
// Public inline
// void setWidth(qreal)
func (this *QRectF) SetWidth(w float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:584
// index:0
// Public inline
// void setHeight(qreal)
func (this *QRectF) SetHeight(h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF9setHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:585
// index:0
// Public inline
// void setSize(const class QSizeF &)
func (this *QRectF) SetSize(s *QSizeF) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QRectF7setSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:592
// index:0
// Public
// bool contains(const class QRectF &)
func (this *QRectF) Contains(r *QRectF) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8containsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:593
// index:1
// Public
// bool contains(const class QPointF &)
func (this *QRectF) Contains_1(p *QPointF) bool {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:594
// index:2
// Public inline
// bool contains(qreal, qreal)
func (this *QRectF) Contains_2(x float64, y float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF8containsEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:595
// index:0
// Public inline
// QRectF united(const class QRectF &)
func (this *QRectF) United(other *QRectF) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6unitedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:596
// index:0
// Public inline
// QRectF intersected(const class QRectF &)
func (this *QRectF) Intersected(other *QRectF) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:597
// index:0
// Public
// bool intersects(const class QRectF &)
func (this *QRectF) Intersects(r *QRectF) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:599
// index:0
// Public inline
// QRectF marginsAdded(const class QMarginsF &)
func (this *QRectF) MarginsAdded(margins *QMarginsF) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF12marginsAddedERK9QMarginsF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:600
// index:0
// Public inline
// QRectF marginsRemoved(const class QMarginsF &)
func (this *QRectF) MarginsRemoved(margins *QMarginsF) *QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF14marginsRemovedERK9QMarginsF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:612
// index:0
// Public inline
// QRect toRect()
func (this *QRectF) ToRect() *QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF6toRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:613
// index:0
// Public
// QRect toAlignedRect()
func (this *QRectF) ToAlignedRect() *QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QRectF13toAlignedRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
