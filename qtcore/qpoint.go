package qtcore

// /usr/include/qt/QtCore/qpoint.h
// #include <qpoint.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QPoint struct {
	*qtrt.CObject
}
type QPoint_ITF interface {
	QPoint_PTR() *QPoint
}

func (ptr *QPoint) QPoint_PTR() *QPoint { return ptr }

func (this *QPoint) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPoint) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPointFromPointer(cthis unsafe.Pointer) *QPoint {
	return &QPoint{&qtrt.CObject{cthis}}
}
func (*QPoint) NewFromPointer(cthis unsafe.Pointer) *QPoint {
	return NewQPointFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPoint()

/*
Constructs a null point, i.e. with coordinates (0, 0)

See also isNull().
*/
func NewQPoint() *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPoint)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:56
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPoint(int, int)

/*
Constructs a null point, i.e. with coordinates (0, 0)

See also isNull().
*/
func NewQPoint_1(xpos int, ypos int) *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointC2Eii", qtrt.FFI_TYPE_POINTER, xpos, ypos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPoint)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if both the x and y coordinates are set to 0, otherwise returns false.
*/
func (this *QPoint) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpoint.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x() const

/*
Returns the x coordinate of this point.

See also setX() and rx().
*/
func (this *QPoint) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*
Returns the y coordinate of this point.

See also setY() and ry().
*/
func (this *QPoint) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(int)

/*
Sets the x coordinate of this point to the given x coordinate.

See also x() and setY().
*/
func (this *QPoint) SetX(x int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint4setXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(int)

/*
Sets the y coordinate of this point to the given y coordinate.

See also y() and setX().
*/
func (this *QPoint) SetY(y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint4setYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int manhattanLength() const

/*
Returns the sum of the absolute values of x() and y(), traditionally known as the "Manhattan length" of the vector from the origin to the point. For example:


  QPoint oldPosition;

  MyWidget::mouseMoveEvent(QMouseEvent *event)
  {
      QPoint point = event->pos() - oldPosition;
      if (point.manhattanLength() > 3)
          // the mouse has moved more than 3 pixels since the oldPosition
  }



This is a useful, and quick to calculate, approximation to the true length:


  double trueLength = std::sqrt(std::pow(x(), 2) + std::pow(y(), 2));



The tradition of "Manhattan length" arises because such distances apply to travelers who can only travel on a rectangular grid, like the streets of Manhattan.
*/
func (this *QPoint) ManhattanLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QPoint15manhattanLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int & rx()

/*
Returns a reference to the x coordinate of this point.

Using a reference makes it possible to directly manipulate x. For example:


  QPoint p(1, 2);
  p.rx()--;   // p becomes (0, 2)



See also x() and setX().
*/
func (this *QPoint) Rx() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint2rxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int & ry()

/*
Returns a reference to the y coordinate of this point.

Using a reference makes it possible to directly manipulate y. For example:


  QPoint p(1, 2);
  p.ry()++;   // p becomes (1, 3)



See also y() and setY().
*/
func (this *QPoint) Ry() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint2ryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint & operator+=(const QPoint &)

/*

 */
func (this *QPoint) Operator_add_equal(p QPoint_ITF) *QPoint {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint & operator-=(const QPoint &)

/*

 */
func (this *QPoint) Operator_minus_equal(p QPoint_ITF) *QPoint {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint & operator*=(float)

/*

 */
func (this *QPoint) Operator_mul_equal(factor float32) *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointmLEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPoint & operator*=(double)

/*

 */
func (this *QPoint) Operator_mul_equal_1(factor float64) *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointmLEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:75
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QPoint & operator*=(int)

/*

 */
func (this *QPoint) Operator_mul_equal_2(factor int) *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointmLEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint & operator/=(qreal)

/*

 */
func (this *QPoint) Operator_div_equal(divisor float64) *QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointdVEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), divisor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:79
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] int dotProduct(const QPoint &, const QPoint &)

/*
QPoint p( 3, 7);
  QPoint q(-1, 4);
  int lengthSquared = QPoint::dotProduct(p, q);   // lengthSquared becomes 25



Returns the dot product of p1 and p2.

This function was introduced in  Qt 5.1.
*/
func (this *QPoint) DotProduct(p1 QPoint_ITF, p2 QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if p1 != nil && p1.QPoint_PTR() != nil {
		convArg0 = p1.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if p2 != nil && p2.QPoint_PTR() != nil {
		convArg1 = p2.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPoint10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPoint_DotProduct(p1 QPoint_ITF, p2 QPoint_ITF) int {
	var nilthis *QPoint
	rv := nilthis.DotProduct(p1, p2)
	return rv
}

func DeleteQPoint(this *QPoint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QPointD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
