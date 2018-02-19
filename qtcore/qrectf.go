package qtcore

// /usr/include/qt/QtCore/qrect.h
// #include <qrect.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 72
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QRectF struct {
	*qtrt.CObject
}
type QRectF_ITF interface {
	QRectF_PTR() *QRectF
}

func (ptr *QRectF) QRectF_PTR() *QRectF { return ptr }

func (this *QRectF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRectF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRectFFromPointer(cthis unsafe.Pointer) *QRectF {
	return &QRectF{&qtrt.CObject{cthis}}
}
func (*QRectF) NewFromPointer(cthis unsafe.Pointer) *QRectF {
	return NewQRectFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrect.h:514
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRectF()
func NewQRectF() *QRectF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRectF)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:515
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QRectF(const QPointF &, const QSizeF &)
func NewQRectF_1(topleft QPointF_ITF, size QSizeF_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if topleft != nil && topleft.QPointF_PTR() != nil {
		convArg0 = topleft.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg1 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFC2ERK7QPointFRK6QSizeF", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRectF)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:516
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QRectF(const QPointF &, const QPointF &)
func NewQRectF_2(topleft QPointF_ITF, bottomRight QPointF_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if topleft != nil && topleft.QPointF_PTR() != nil {
		convArg0 = topleft.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bottomRight != nil && bottomRight.QPointF_PTR() != nil {
		convArg1 = bottomRight.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFC2ERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRectF)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:517
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QRectF(qreal, qreal, qreal, qreal)
func NewQRectF_3(left float64, top float64, width float64, height float64) *QRectF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFC2Edddd", qtrt.FFI_TYPE_POINTER, left, top, width, height)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRectF)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:518
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QRectF(const QRect &)
func NewQRectF_4(rect QRect_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFC2ERK5QRect", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRectF)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:520
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QRectF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:521
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QRectF) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:522
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QRectF) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:523
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF normalized() const
func (this *QRectF) Normalized() *QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:525
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal left() const
func (this *QRectF) Left() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF4leftEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:526
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal top() const
func (this *QRectF) Top() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF3topEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:527
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal right() const
func (this *QRectF) Right() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF5rightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:528
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottom() const
func (this *QRectF) Bottom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6bottomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:530
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x() const
func (this *QRectF) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:531
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y() const
func (this *QRectF) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:532
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeft(qreal)
func (this *QRectF) SetLeft(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7setLeftEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:533
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTop(qreal)
func (this *QRectF) SetTop(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6setTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:534
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRight(qreal)
func (this *QRectF) SetRight(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF8setRightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:535
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottom(qreal)
func (this *QRectF) SetBottom(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9setBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:536
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(qreal)
func (this *QRectF) SetX(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:537
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(qreal)
func (this *QRectF) SetY(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:539
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF topLeft() const
func (this *QRectF) TopLeft() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7topLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:540
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF bottomRight() const
func (this *QRectF) BottomRight() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF11bottomRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:541
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF topRight() const
func (this *QRectF) TopRight() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF8topRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:542
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF bottomLeft() const
func (this *QRectF) BottomLeft() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF10bottomLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:543
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF center() const
func (this *QRectF) Center() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:545
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopLeft(const QPointF &)
func (this *QRectF) SetTopLeft(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF10setTopLeftERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:546
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomRight(const QPointF &)
func (this *QRectF) SetBottomRight(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF14setBottomRightERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:547
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopRight(const QPointF &)
func (this *QRectF) SetTopRight(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF11setTopRightERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:548
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomLeft(const QPointF &)
func (this *QRectF) SetBottomLeft(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF13setBottomLeftERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:550
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveLeft(qreal)
func (this *QRectF) MoveLeft(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF8moveLeftEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:551
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTop(qreal)
func (this *QRectF) MoveTop(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7moveTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:552
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveRight(qreal)
func (this *QRectF) MoveRight(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9moveRightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:553
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottom(qreal)
func (this *QRectF) MoveBottom(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF10moveBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:554
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopLeft(const QPointF &)
func (this *QRectF) MoveTopLeft(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF11moveTopLeftERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:555
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottomRight(const QPointF &)
func (this *QRectF) MoveBottomRight(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF15moveBottomRightERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:556
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopRight(const QPointF &)
func (this *QRectF) MoveTopRight(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF12moveTopRightERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:557
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottomLeft(const QPointF &)
func (this *QRectF) MoveBottomLeft(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF14moveBottomLeftERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:558
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveCenter(const QPointF &)
func (this *QRectF) MoveCenter(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF10moveCenterERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:560
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QRectF) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:561
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)
func (this *QRectF) Translate_1(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:563
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF translated(qreal, qreal) const
func (this *QRectF) Translated(dx float64, dy float64) *QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:564
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF translated(const QPointF &) const
func (this *QRectF) Translated_1(p QPointF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF10translatedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:566
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF transposed() const
func (this *QRectF) Transposed() *QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:568
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(qreal, qreal)
func (this *QRectF) MoveTo(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6moveToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:569
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(const QPointF &)
func (this *QRectF) MoveTo_1(p QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6moveToERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:571
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QRectF) SetRect(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:572
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getRect(qreal *, qreal *, qreal *, qreal *) const
func (this *QRectF) GetRect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7getRectEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:574
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCoords(qreal, qreal, qreal, qreal)
func (this *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9setCoordsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:575
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getCoords(qreal *, qreal *, qreal *, qreal *) const
func (this *QRectF) GetCoords(x1 unsafe.Pointer /*666*/, y1 unsafe.Pointer /*666*/, x2 unsafe.Pointer /*666*/, y2 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF9getCoordsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:577
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void adjust(qreal, qreal, qreal, qreal)
func (this *QRectF) Adjust(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6adjustEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:578
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF adjusted(qreal, qreal, qreal, qreal) const
func (this *QRectF) Adjusted(x1 float64, y1 float64, x2 float64, y2 float64) *QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF8adjustedEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:580
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF size() const
func (this *QRectF) Size() *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:581
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width() const
func (this *QRectF) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:582
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height() const
func (this *QRectF) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:583
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)
func (this *QRectF) SetWidth(w float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:584
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)
func (this *QRectF) SetHeight(h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:585
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSize(const QSizeF &)
func (this *QRectF) SetSize(s QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizeF_PTR() != nil {
		convArg0 = s.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7setSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:587
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF operator|(const QRectF &) const
func (this *QRectF) Operator_or(r QRectF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectForERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:588
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF operator&(const QRectF &) const
func (this *QRectF) Operator_and(r QRectF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectFanERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:589
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF & operator|=(const QRectF &)
func (this *QRectF) Operator_or_equal(r QRectF_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFoRERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:590
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF & operator&=(const QRectF &)
func (this *QRectF) Operator_and_equal(r QRectF_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFaNERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:592
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRectF &) const
func (this *QRectF) Contains(r QRectF_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF8containsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:593
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const
func (this *QRectF) Contains_1(p QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:594
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(qreal, qreal) const
func (this *QRectF) Contains_2(x float64, y float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF8containsEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:595
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF united(const QRectF &) const
func (this *QRectF) United(other QRectF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRectF_PTR() != nil {
		convArg0 = other.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:596
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF intersected(const QRectF &) const
func (this *QRectF) Intersected(other QRectF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRectF_PTR() != nil {
		convArg0 = other.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:597
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRectF &) const
func (this *QRectF) Intersects(r QRectF_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:599
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF marginsAdded(const QMarginsF &) const
func (this *QRectF) MarginsAdded(margins QMarginsF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF12marginsAddedERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:600
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF marginsRemoved(const QMarginsF &) const
func (this *QRectF) MarginsRemoved(margins QMarginsF_ITF) *QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF14marginsRemovedERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:601
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF & operator+=(const QMarginsF &)
func (this *QRectF) Operator_add_equal(margins QMarginsF_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFpLERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:602
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF & operator-=(const QMarginsF &)
func (this *QRectF) Operator_minus_equal(margins QMarginsF_ITF) *QRectF {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFmIERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:612
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect toRect() const
func (this *QRectF) ToRect() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6toRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:613
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect toAlignedRect() const
func (this *QRectF) ToAlignedRect() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF13toAlignedRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

func DeleteQRectF(this *QRectF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
