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

/*

 */
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

/*

 */
func (*QRectF) NewForInherit() *QRectF {
	return NewQRectF()
}
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

/*

 */
func (*QRectF) NewForInherit_1(topleft QPointF_ITF, size QSizeF_ITF) *QRectF {
	return NewQRectF_1(topleft, size)
}
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

/*

 */
func (*QRectF) NewForInherit_2(topleft QPointF_ITF, bottomRight QPointF_ITF) *QRectF {
	return NewQRectF_2(topleft, bottomRight)
}
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

/*

 */
func (*QRectF) NewForInherit_3(left float64, top float64, width float64, height float64) *QRectF {
	return NewQRectF_3(left, top, width, height)
}
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

/*

 */
func (*QRectF) NewForInherit_4(rect QRect_ITF) *QRectF {
	return NewQRectF_4(rect)
}
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

/*
Returns true if the rectangle is a null rectangle, otherwise returns false.

A null rectangle has both the width and the height set to 0 (i.e., right() == left() - 1 and bottom() == top() - 1). A null rectangle is also empty, and hence is not valid.

See also isEmpty() and isValid().
*/
func (this *QRectF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:521
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the rectangle is empty, otherwise returns false.

An empty rectangle has a left() > right() or top() > bottom(). An empty rectangle is not valid (i.e., isEmpty() == !isValid()).

Use the normalized() function to retrieve a rectangle where the corners are swapped.

See also isNull(), isValid(), and normalized().
*/
func (this *QRectF) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:522
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the rectangle is valid, otherwise returns false.

A valid rectangle has a left() <= right() and top() <= bottom(). Note that non-trivial operations like intersections are not defined for invalid rectangles. A valid rectangle is not empty (i.e., isValid() == !isEmpty()).

See also isNull(), isEmpty(), and normalized().
*/
func (this *QRectF) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:523
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF normalized() const

/*
Returns a normalized rectangle; i.e., a rectangle that has a non-negative width and height.

If width() < 0 the function swaps the left and right corners, and it swaps the top and bottom corners if height() < 0.

See also isValid() and isEmpty().
*/
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

/*
Returns the x-coordinate of the rectangle's left edge. Equivalent to x().

See also setLeft(), topLeft(), and bottomLeft().
*/
func (this *QRectF) Left() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF4leftEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:526
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal top() const

/*
Returns the y-coordinate of the rectangle's top edge. Equivalent to y().

See also setTop(), topLeft(), and topRight().
*/
func (this *QRectF) Top() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF3topEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:527
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal right() const

/*
Returns the x-coordinate of the rectangle's right edge.

Note that for historical reasons this function returns left() + width() - 1; use x() + width() to retrieve the true x-coordinate.

See also setRight(), topRight(), and bottomRight().
*/
func (this *QRectF) Right() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF5rightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:528
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottom() const

/*
Returns the y-coordinate of the rectangle's bottom edge.

Note that for historical reasons this function returns top() + height() - 1; use y() + height() to retrieve the true y-coordinate.

See also setBottom(), bottomLeft(), and bottomRight().
*/
func (this *QRectF) Bottom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6bottomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:530
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x() const

/*
Returns the x-coordinate of the rectangle's left edge. Equivalent to left().

See also setX(), y(), and topLeft().
*/
func (this *QRectF) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:531
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y() const

/*
Returns the y-coordinate of the rectangle's top edge. Equivalent to top().

See also setY(), x(), and topLeft().
*/
func (this *QRectF) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:532
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeft(qreal)

/*
Sets the left edge of the rectangle to the given x coordinate. May change the width, but will never change the right edge of the rectangle.

Equivalent to setX().

See also left() and moveLeft().
*/
func (this *QRectF) SetLeft(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7setLeftEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:533
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTop(qreal)

/*
Sets the top edge of the rectangle to the given y coordinate. May change the height, but will never change the bottom edge of the rectangle.

Equivalent to setY().

See also top() and moveTop().
*/
func (this *QRectF) SetTop(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6setTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:534
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRight(qreal)

/*
Sets the right edge of the rectangle to the given x coordinate. May change the width, but will never change the left edge of the rectangle.

See also right() and moveRight().
*/
func (this *QRectF) SetRight(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF8setRightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:535
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottom(qreal)

/*
Sets the bottom edge of the rectangle to the given y coordinate. May change the height, but will never change the top edge of the rectangle.

See also bottom() and moveBottom().
*/
func (this *QRectF) SetBottom(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9setBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:536
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(qreal)

/*
Sets the left edge of the rectangle to the given x coordinate. May change the width, but will never change the right edge of the rectangle.

Equivalent to setLeft().

See also x(), setY(), and setTopLeft().
*/
func (this *QRectF) SetX(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:537
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(qreal)

/*
Sets the top edge of the rectangle to the given y coordinate. May change the height, but will never change the bottom edge of the rectangle.

Equivalent to setTop().

See also y(), setX(), and setTopLeft().
*/
func (this *QRectF) SetY(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:539
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF topLeft() const

/*
Returns the position of the rectangle's top-left corner.

See also setTopLeft(), top(), and left().
*/
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

/*
Returns the position of the rectangle's bottom-right corner.

Note that for historical reasons this function returns QPoint(left() + width() -1, top() + height() - 1).

See also setBottomRight(), bottom(), and right().
*/
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

/*
Returns the position of the rectangle's top-right corner.

Note that for historical reasons this function returns QPoint(left() + width() -1, top()).

See also setTopRight(), top(), and right().
*/
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

/*
Returns the position of the rectangle's bottom-left corner. Note that for historical reasons this function returns QPoint(left(), top() + height() - 1).

See also setBottomLeft(), bottom(), and left().
*/
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

/*
Returns the center point of the rectangle.

See also moveCenter().
*/
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

/*
Set the top-left corner of the rectangle to the given position. May change the size, but will never change the bottom-right corner of the rectangle.

See also topLeft() and moveTopLeft().
*/
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

/*
Set the bottom-right corner of the rectangle to the given position. May change the size, but will never change the top-left corner of the rectangle.

See also bottomRight() and moveBottomRight().
*/
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

/*
Set the top-right corner of the rectangle to the given position. May change the size, but will never change the bottom-left corner of the rectangle.

See also topRight() and moveTopRight().
*/
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

/*
Set the bottom-left corner of the rectangle to the given position. May change the size, but will never change the top-right corner of the rectangle.

See also bottomLeft() and moveBottomLeft().
*/
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

/*
Moves the rectangle horizontally, leaving the rectangle's left edge at the given x coordinate. The rectangle's size is unchanged.

See also left(), setLeft(), and moveRight().
*/
func (this *QRectF) MoveLeft(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF8moveLeftEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:551
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTop(qreal)

/*
Moves the rectangle vertically, leaving the rectangle's top edge at the given y coordinate. The rectangle's size is unchanged.

See also top(), setTop(), and moveBottom().
*/
func (this *QRectF) MoveTop(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7moveTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:552
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveRight(qreal)

/*
Moves the rectangle horizontally, leaving the rectangle's right edge at the given x coordinate. The rectangle's size is unchanged.

See also right(), setRight(), and moveLeft().
*/
func (this *QRectF) MoveRight(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9moveRightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:553
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottom(qreal)

/*
Moves the rectangle vertically, leaving the rectangle's bottom edge at the given y coordinate. The rectangle's size is unchanged.

See also bottom(), setBottom(), and moveTop().
*/
func (this *QRectF) MoveBottom(pos float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF10moveBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:554
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopLeft(const QPointF &)

/*
Moves the rectangle, leaving the top-left corner at the given position. The rectangle's size is unchanged.

See also setTopLeft(), moveTop(), and moveLeft().
*/
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

/*
Moves the rectangle, leaving the bottom-right corner at the given position. The rectangle's size is unchanged.

See also setBottomRight(), moveRight(), and moveBottom().
*/
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

/*
Moves the rectangle, leaving the top-right corner at the given position. The rectangle's size is unchanged.

See also setTopRight(), moveTop(), and moveRight().
*/
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

/*
Moves the rectangle, leaving the bottom-left corner at the given position. The rectangle's size is unchanged.

See also setBottomLeft(), moveBottom(), and moveLeft().
*/
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

/*
Moves the rectangle, leaving the center point at the given position. The rectangle's size is unchanged.

See also center().
*/
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

/*
Moves the rectangle dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also moveTopLeft(), moveTo(), and translated().
*/
func (this *QRectF) Translate(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:561
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)

/*
Moves the rectangle dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also moveTopLeft(), moveTo(), and translated().
*/
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

/*
Returns a copy of the rectangle that is translated dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also translate().
*/
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

/*
Returns a copy of the rectangle that is translated dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also translate().
*/
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

/*
Returns a copy of the rectangle that has its width and height exchanged:


  QRect r = {15, 51, 42, 24};
  r = r.transposed(); // r == {15, 51, 24, 42}



This function was introduced in  Qt 5.7.

See also QSize::transposed().
*/
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

/*
Moves the rectangle, leaving the top-left corner at the given position (x, y). The rectangle's size is unchanged.

See also translate() and moveTopLeft().
*/
func (this *QRectF) MoveTo(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6moveToEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:569
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(const QPointF &)

/*
Moves the rectangle, leaving the top-left corner at the given position (x, y). The rectangle's size is unchanged.

See also translate() and moveTopLeft().
*/
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

/*
Sets the coordinates of the rectangle's top-left corner to (x, y), and its size to the given width and height.

See also getRect() and setCoords().
*/
func (this *QRectF) SetRect(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:572
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getRect(qreal *, qreal *, qreal *, qreal *) const

/*
Extracts the position of the rectangle's top-left corner to *x and *y, and its dimensions to *width and *height.

See also setRect() and getCoords().
*/
func (this *QRectF) GetRect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF7getRectEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:574
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCoords(qreal, qreal, qreal, qreal)

/*
Sets the coordinates of the rectangle's top-left corner to (x1, y1), and the coordinates of its bottom-right corner to (x2, y2).

See also getCoords() and setRect().
*/
func (this *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9setCoordsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:575
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getCoords(qreal *, qreal *, qreal *, qreal *) const

/*
Extracts the position of the rectangle's top-left corner to *x1 and *y1, and the position of the bottom-right corner to *x2 and *y2.

See also setCoords() and getRect().
*/
func (this *QRectF) GetCoords(x1 unsafe.Pointer /*666*/, y1 unsafe.Pointer /*666*/, x2 unsafe.Pointer /*666*/, y2 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF9getCoordsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:577
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void adjust(qreal, qreal, qreal, qreal)

/*
Adds dx1, dy1, dx2 and dy2 respectively to the existing coordinates of the rectangle.

See also adjusted() and setRect().
*/
func (this *QRectF) Adjust(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF6adjustEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:578
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF adjusted(qreal, qreal, qreal, qreal) const

/*
Returns a new rectangle with dx1, dy1, dx2 and dy2 added respectively to the existing coordinates of this rectangle.

See also adjust().
*/
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

/*
Returns the size of the rectangle.

See also setSize(), width(), and height().
*/
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

/*
Returns the width of the rectangle.

See also setWidth(), height(), and size().
*/
func (this *QRectF) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:582
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height() const

/*
Returns the height of the rectangle.

See also setHeight(), width(), and size().
*/
func (this *QRectF) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrect.h:583
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)

/*
Sets the width of the rectangle to the given width. The right edge is changed, but not the left one.

See also width() and setSize().
*/
func (this *QRectF) SetWidth(w float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:584
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)

/*
Sets the height of the rectangle to the given height. The bottom edge is changed, but not the top one.

See also height() and setSize().
*/
func (this *QRectF) SetHeight(h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QRectF9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:585
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSize(const QSizeF &)

/*
Sets the size of the rectangle to the given size. The top-left corner is not moved.

See also size(), setWidth(), and setHeight().
*/
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
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

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
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

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRectF) Contains_2(x float64, y float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QRectF8containsEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:595
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF united(const QRectF &) const

/*
Returns the bounding rectangle of this rectangle and the given rectangle.



This function was introduced in  Qt 4.2.

See also intersected().
*/
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

/*
Returns the intersection of this rectangle and the given rectangle. Note that r.intersected(s) is equivalent to r & s.



This function was introduced in  Qt 4.2.

See also intersects(), united(), and operator&=().
*/
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

/*
Returns true if this rectangle intersects with the given rectangle (i.e., there is at least one pixel that is within both rectangles), otherwise returns false.

The intersection rectangle can be retrieved using the intersected() function.

See also contains().
*/
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

/*
Returns a rectangle grown by the margins.

This function was introduced in  Qt 5.1.

See also operator+=(), marginsRemoved(), and operator-=().
*/
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

/*
Removes the margins from the rectangle, shrinking it.

This function was introduced in  Qt 5.1.

See also marginsAdded(), operator+=(), and operator-=().
*/
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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
