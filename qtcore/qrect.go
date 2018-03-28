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
// extern C begin: 25
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
type QRect struct {
	*qtrt.CObject
}
type QRect_ITF interface {
	QRect_PTR() *QRect
}

func (ptr *QRect) QRect_PTR() *QRect { return ptr }

func (this *QRect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRect) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRectFromPointer(cthis unsafe.Pointer) *QRect {
	return &QRect{&qtrt.CObject{cthis}}
}
func (*QRect) NewFromPointer(cthis unsafe.Pointer) *QRect {
	return NewQRectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrect.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRect()

/*
Constructs a null rectangle.

See also isNull().
*/
func NewQRect() *QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QRect(const QPoint &, const QPoint &)

/*
Constructs a null rectangle.

See also isNull().
*/
func NewQRect_1(topleft QPoint_ITF, bottomright QPoint_ITF) *QRect {
	var convArg0 unsafe.Pointer
	if topleft != nil && topleft.QPoint_PTR() != nil {
		convArg0 = topleft.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bottomright != nil && bottomright.QPoint_PTR() != nil {
		convArg1 = bottomright.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2ERK6QPointS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:62
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QRect(const QPoint &, const QSize &)

/*
Constructs a null rectangle.

See also isNull().
*/
func NewQRect_2(topleft QPoint_ITF, size QSize_ITF) *QRect {
	var convArg0 unsafe.Pointer
	if topleft != nil && topleft.QPoint_PTR() != nil {
		convArg0 = topleft.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2ERK6QPointRK5QSize", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:63
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QRect(int, int, int, int)

/*
Constructs a null rectangle.

See also isNull().
*/
func NewQRect_3(left int, top int, width int, height int) *QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2Eiiii", qtrt.FFI_TYPE_POINTER, left, top, width, height)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the rectangle is a null rectangle, otherwise returns false.

A null rectangle has both the width and the height set to 0 (i.e., right() == left() - 1 and bottom() == top() - 1). A null rectangle is also empty, and hence is not valid.

See also isEmpty() and isValid().
*/
func (this *QRect) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the rectangle is empty, otherwise returns false.

An empty rectangle has a left() > right() or top() > bottom(). An empty rectangle is not valid (i.e., isEmpty() == !isValid()).

Use the normalized() function to retrieve a rectangle where the corners are swapped.

See also isNull(), isValid(), and normalized().
*/
func (this *QRect) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the rectangle is valid, otherwise returns false.

A valid rectangle has a left() <= right() and top() <= bottom(). Note that non-trivial operations like intersections are not defined for invalid rectangles. A valid rectangle is not empty (i.e., isValid() == !isEmpty()).

See also isNull(), isEmpty(), and normalized().
*/
func (this *QRect) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int left() const

/*
Returns the x-coordinate of the rectangle's left edge. Equivalent to x().

See also setLeft(), topLeft(), and bottomLeft().
*/
func (this *QRect) Left() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect4leftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top() const

/*
Returns the y-coordinate of the rectangle's top edge. Equivalent to y().

See also setTop(), topLeft(), and topRight().
*/
func (this *QRect) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int right() const

/*
Returns the x-coordinate of the rectangle's right edge.

Note that for historical reasons this function returns left() + width() - 1; use x() + width() to retrieve the true x-coordinate.

See also setRight(), topRight(), and bottomRight().
*/
func (this *QRect) Right() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect5rightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom() const

/*
Returns the y-coordinate of the rectangle's bottom edge.

Note that for historical reasons this function returns top() + height() - 1; use y() + height() to retrieve the true y-coordinate.

See also setBottom(), bottomLeft(), and bottomRight().
*/
func (this *QRect) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:73
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect normalized() const

/*
Returns a normalized rectangle; i.e., a rectangle that has a non-negative width and height.

If width() < 0 the function swaps the left and right corners, and it swaps the top and bottom corners if height() < 0.

See also isValid() and isEmpty().
*/
func (this *QRect) Normalized() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x() const

/*
Returns the x-coordinate of the rectangle's left edge. Equivalent to left().

See also setX(), y(), and topLeft().
*/
func (this *QRect) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*
Returns the y-coordinate of the rectangle's top edge. Equivalent to top().

See also setY(), x(), and topLeft().
*/
func (this *QRect) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeft(int)

/*
Sets the left edge of the rectangle to the given x coordinate. May change the width, but will never change the right edge of the rectangle.

Equivalent to setX().

See also left() and moveLeft().
*/
func (this *QRect) SetLeft(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7setLeftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTop(int)

/*
Sets the top edge of the rectangle to the given y coordinate. May change the height, but will never change the bottom edge of the rectangle.

Equivalent to setY().

See also top() and moveTop().
*/
func (this *QRect) SetTop(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6setTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRight(int)

/*
Sets the right edge of the rectangle to the given x coordinate. May change the width, but will never change the left edge of the rectangle.

See also right() and moveRight().
*/
func (this *QRect) SetRight(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect8setRightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottom(int)

/*
Sets the bottom edge of the rectangle to the given y coordinate. May change the height, but will never change the top edge of the rectangle.

See also bottom() and moveBottom().
*/
func (this *QRect) SetBottom(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9setBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(int)

/*
Sets the left edge of the rectangle to the given x coordinate. May change the width, but will never change the right edge of the rectangle.

Equivalent to setLeft().

See also x(), setY(), and setTopLeft().
*/
func (this *QRect) SetX(x int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect4setXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(int)

/*
Sets the top edge of the rectangle to the given y coordinate. May change the height, but will never change the bottom edge of the rectangle.

Equivalent to setTop().

See also y(), setX(), and setTopLeft().
*/
func (this *QRect) SetY(y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect4setYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopLeft(const QPoint &)

/*
Set the top-left corner of the rectangle to the given position. May change the size, but will never change the bottom-right corner of the rectangle.

See also topLeft() and moveTopLeft().
*/
func (this *QRect) SetTopLeft(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect10setTopLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomRight(const QPoint &)

/*
Set the bottom-right corner of the rectangle to the given position. May change the size, but will never change the top-left corner of the rectangle.

See also bottomRight() and moveBottomRight().
*/
func (this *QRect) SetBottomRight(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect14setBottomRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopRight(const QPoint &)

/*
Set the top-right corner of the rectangle to the given position. May change the size, but will never change the bottom-left corner of the rectangle.

See also topRight() and moveTopRight().
*/
func (this *QRect) SetTopRight(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect11setTopRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomLeft(const QPoint &)

/*
Set the bottom-left corner of the rectangle to the given position. May change the size, but will never change the top-right corner of the rectangle.

See also bottomLeft() and moveBottomLeft().
*/
func (this *QRect) SetBottomLeft(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect13setBottomLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint topLeft() const

/*
Returns the position of the rectangle's top-left corner.

See also setTopLeft(), top(), and left().
*/
func (this *QRect) TopLeft() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7topLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint bottomRight() const

/*
Returns the position of the rectangle's bottom-right corner.

Note that for historical reasons this function returns QPoint(left() + width() -1, top() + height() - 1).

See also setBottomRight(), bottom(), and right().
*/
func (this *QRect) BottomRight() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect11bottomRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint topRight() const

/*
Returns the position of the rectangle's top-right corner.

Note that for historical reasons this function returns QPoint(left() + width() -1, top()).

See also setTopRight(), top(), and right().
*/
func (this *QRect) TopRight() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8topRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint bottomLeft() const

/*
Returns the position of the rectangle's bottom-left corner. Note that for historical reasons this function returns QPoint(left(), top() + height() - 1).

See also setBottomLeft(), bottom(), and left().
*/
func (this *QRect) BottomLeft() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10bottomLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint center() const

/*
Returns the center point of the rectangle.

See also moveCenter().
*/
func (this *QRect) Center() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveLeft(int)

/*
Moves the rectangle horizontally, leaving the rectangle's left edge at the given x coordinate. The rectangle's size is unchanged.

See also left(), setLeft(), and moveRight().
*/
func (this *QRect) MoveLeft(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect8moveLeftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTop(int)

/*
Moves the rectangle vertically, leaving the rectangle's top edge at the given y coordinate. The rectangle's size is unchanged.

See also top(), setTop(), and moveBottom().
*/
func (this *QRect) MoveTop(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7moveTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveRight(int)

/*
Moves the rectangle horizontally, leaving the rectangle's right edge at the given x coordinate. The rectangle's size is unchanged.

See also right(), setRight(), and moveLeft().
*/
func (this *QRect) MoveRight(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9moveRightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottom(int)

/*
Moves the rectangle vertically, leaving the rectangle's bottom edge at the given y coordinate. The rectangle's size is unchanged.

See also bottom(), setBottom(), and moveTop().
*/
func (this *QRect) MoveBottom(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect10moveBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopLeft(const QPoint &)

/*
Moves the rectangle, leaving the top-left corner at the given position. The rectangle's size is unchanged.

See also setTopLeft(), moveTop(), and moveLeft().
*/
func (this *QRect) MoveTopLeft(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect11moveTopLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottomRight(const QPoint &)

/*
Moves the rectangle, leaving the bottom-right corner at the given position. The rectangle's size is unchanged.

See also setBottomRight(), moveRight(), and moveBottom().
*/
func (this *QRect) MoveBottomRight(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect15moveBottomRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopRight(const QPoint &)

/*
Moves the rectangle, leaving the top-right corner at the given position. The rectangle's size is unchanged.

See also setTopRight(), moveTop(), and moveRight().
*/
func (this *QRect) MoveTopRight(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect12moveTopRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottomLeft(const QPoint &)

/*
Moves the rectangle, leaving the bottom-left corner at the given position. The rectangle's size is unchanged.

See also setBottomLeft(), moveBottom(), and moveLeft().
*/
func (this *QRect) MoveBottomLeft(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect14moveBottomLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveCenter(const QPoint &)

/*
Moves the rectangle, leaving the center point at the given position. The rectangle's size is unchanged.

See also center().
*/
func (this *QRect) MoveCenter(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect10moveCenterERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(int, int)

/*
Moves the rectangle dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also moveTopLeft(), moveTo(), and translated().
*/
func (this *QRect) Translate(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9translateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:106
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)

/*
Moves the rectangle dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also moveTopLeft(), moveTo(), and translated().
*/
func (this *QRect) Translate_1(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect translated(int, int) const

/*
Returns a copy of the rectangle that is translated dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also translate().
*/
func (this *QRect) Translated(dx int, dy int) *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10translatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:108
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QRect translated(const QPoint &) const

/*
Returns a copy of the rectangle that is translated dx along the x axis and dy along the y axis, relative to the current position. Positive values move the rectangle to the right and down.

See also translate().
*/
func (this *QRect) Translated_1(p QPoint_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10translatedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect transposed() const

/*
Returns a copy of the rectangle that has its width and height exchanged:


  QRect r = {15, 51, 42, 24};
  r = r.transposed(); // r == {15, 51, 24, 42}



This function was introduced in  Qt 5.7.

See also QSize::transposed().
*/
func (this *QRect) Transposed() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(int, int)

/*
Moves the rectangle, leaving the top-left corner at the given position (x, y). The rectangle's size is unchanged.

See also translate() and moveTopLeft().
*/
func (this *QRect) MoveTo(x int, t int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6moveToEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, t)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:112
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(const QPoint &)

/*
Moves the rectangle, leaving the top-left corner at the given position (x, y). The rectangle's size is unchanged.

See also translate() and moveTopLeft().
*/
func (this *QRect) MoveTo_1(p QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6moveToERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(int, int, int, int)

/*
Sets the coordinates of the rectangle's top-left corner to (x, y), and its size to the given width and height.

See also getRect() and setCoords().
*/
func (this *QRect) SetRect(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7setRectEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getRect(int *, int *, int *, int *) const

/*
Extracts the position of the rectangle's top-left corner to *x and *y, and its dimensions to *width and *height.

See also setRect() and getCoords().
*/
func (this *QRect) GetRect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7getRectEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCoords(int, int, int, int)

/*
Sets the coordinates of the rectangle's top-left corner to (x1, y1), and the coordinates of its bottom-right corner to (x2, y2).

See also getCoords() and setRect().
*/
func (this *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9setCoordsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getCoords(int *, int *, int *, int *) const

/*
Extracts the position of the rectangle's top-left corner to *x1 and *y1, and the position of the bottom-right corner to *x2 and *y2.

See also setCoords() and getRect().
*/
func (this *QRect) GetCoords(x1 unsafe.Pointer /*666*/, y1 unsafe.Pointer /*666*/, x2 unsafe.Pointer /*666*/, y2 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect9getCoordsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void adjust(int, int, int, int)

/*
Adds dx1, dy1, dx2 and dy2 respectively to the existing coordinates of the rectangle.

See also adjusted() and setRect().
*/
func (this *QRect) Adjust(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6adjustEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect adjusted(int, int, int, int) const

/*
Returns a new rectangle with dx1, dy1, dx2 and dy2 added respectively to the existing coordinates of this rectangle.

See also adjust().
*/
func (this *QRect) Adjusted(x1 int, y1 int, x2 int, y2 int) *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8adjustedEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the size of the rectangle.

See also setSize(), width(), and height().
*/
func (this *QRect) Size() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the width of the rectangle.

See also setWidth(), height(), and size().
*/
func (this *QRect) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height() const

/*
Returns the height of the rectangle.

See also setHeight(), width(), and size().
*/
func (this *QRect) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(int)

/*
Sets the width of the rectangle to the given width. The right edge is changed, but not the left one.

See also width() and setSize().
*/
func (this *QRect) SetWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(int)

/*
Sets the height of the rectangle to the given height. The bottom edge is changed, but not the top one.

See also height() and setSize().
*/
func (this *QRect) SetHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9setHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSize(const QSize &)

/*
Sets the size of the rectangle to the given size. The top-left corner is not moved.

See also size(), setWidth(), and setHeight().
*/
func (this *QRect) SetSize(s QSize_ITF) {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7setSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect operator|(const QRect &) const

/*

 */
func (this *QRect) Operator_or(r QRect_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRectorERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:131
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect operator&(const QRect &) const

/*

 */
func (this *QRect) Operator_and(r QRect_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRectanERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect & operator|=(const QRect &)

/*

 */
func (this *QRect) Operator_or_equal(r QRect_ITF) *QRect {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectoRERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect & operator&=(const QRect &)

/*

 */
func (this *QRect) Operator_and_equal(r QRect_ITF) *QRect {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectaNERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRect &, bool) const

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRect) Contains(r QRect_ITF, proper bool) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsERKS_b", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, proper)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRect &, bool) const

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRect) Contains__(r QRect_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	proper := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsERKS_b", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, proper)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:136
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPoint &, bool) const

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRect) Contains_1(p QPoint_ITF, proper bool) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsERK6QPointb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, proper)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:136
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPoint &, bool) const

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRect) Contains_1_(p QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	proper := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsERK6QPointb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, proper)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:137
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(int, int) const

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRect) Contains_2(x int, y int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:138
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(int, int, bool) const

/*
Returns true if the given point is inside or on the edge of the rectangle, otherwise returns false. If proper is true, this function only returns true if the given point is inside the rectangle (i.e., not on the edge).

See also intersects().
*/
func (this *QRect) Contains_3(x int, y int, proper bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsEiib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, proper)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect united(const QRect &) const

/*
Returns the bounding rectangle of this rectangle and the given rectangle.



This function was introduced in  Qt 4.2.

See also intersected().
*/
func (this *QRect) United(other QRect_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRect_PTR() != nil {
		convArg0 = other.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect intersected(const QRect &) const

/*
Returns the intersection of this rectangle and the given rectangle. Note that r.intersected(s) is equivalent to r & s.



This function was introduced in  Qt 4.2.

See also intersects(), united(), and operator&=().
*/
func (this *QRect) Intersected(other QRect_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRect_PTR() != nil {
		convArg0 = other.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRect &) const

/*
Returns true if this rectangle intersects with the given rectangle (i.e., there is at least one pixel that is within both rectangles), otherwise returns false.

The intersection rectangle can be retrieved using the intersected() function.

See also contains().
*/
func (this *QRect) Intersects(r QRect_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect marginsAdded(const QMargins &) const

/*
Returns a rectangle grown by the margins.

This function was introduced in  Qt 5.1.

See also operator+=(), marginsRemoved(), and operator-=().
*/
func (this *QRect) MarginsAdded(margins QMargins_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect12marginsAddedERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect marginsRemoved(const QMargins &) const

/*
Removes the margins from the rectangle, shrinking it.

This function was introduced in  Qt 5.1.

See also marginsAdded(), operator+=(), and operator-=().
*/
func (this *QRect) MarginsRemoved(margins QMargins_ITF) *QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect14marginsRemovedERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:145
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect & operator+=(const QMargins &)

/*

 */
func (this *QRect) Operator_add_equal(margins QMargins_ITF) *QRect {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectpLERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect & operator-=(const QMargins &)

/*

 */
func (this *QRect) Operator_minus_equal(margins QMargins_ITF) *QRect {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectmIERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

func DeleteQRect(this *QRect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
