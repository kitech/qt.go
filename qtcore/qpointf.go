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
// extern C begin: 17
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
type QPointF struct {
	*qtrt.CObject
}
type QPointF_ITF interface {
	QPointF_PTR() *QPointF
}

func (ptr *QPointF) QPointF_PTR() *QPointF { return ptr }

func (this *QPointF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPointF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPointFFromPointer(cthis unsafe.Pointer) *QPointF {
	return &QPointF{&qtrt.CObject{cthis}}
}
func (*QPointF) NewFromPointer(cthis unsafe.Pointer) *QPointF {
	return NewQPointFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpoint.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPointF()

/*

 */
func (*QPointF) NewForInherit() *QPointF {
	return NewQPointF()
}
func NewQPointF() *QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointF)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:223
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPointF(const QPoint &)

/*

 */
func (*QPointF) NewForInherit1(p QPoint_ITF) *QPointF {
	return NewQPointF1(p)
}
func NewQPointF1(p QPoint_ITF) *QPointF {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointF)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:224
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QPointF(qreal, qreal)

/*

 */
func (*QPointF) NewForInherit2(xpos float64, ypos float64) *QPointF {
	return NewQPointF2(xpos, ypos)
}
func NewQPointF2(xpos float64, ypos float64) *QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFC2Edd", qtrt.FFI_TYPE_POINTER, xpos, ypos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointF)
	return gothis
}

// /usr/include/qt/QtCore/qpoint.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal manhattanLength() const

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
func (this *QPointF) ManhattanLength() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF15manhattanLengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if both the x and y coordinates are set to 0, otherwise returns false.
*/
func (this *QPointF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpoint.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x() const

/*
Returns the x coordinate of this point.

See also setX() and rx().
*/
func (this *QPointF) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y() const

/*
Returns the y coordinate of this point.

See also setY() and ry().
*/
func (this *QPointF) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qpoint.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(qreal)

/*
Sets the x coordinate of this point to the given x coordinate.

See also x() and setY().
*/
func (this *QPointF) SetX(x float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(qreal)

/*
Sets the y coordinate of this point to the given y coordinate.

See also y() and setX().
*/
func (this *QPointF) SetY(y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpoint.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & rx()

/*
Returns a reference to the x coordinate of this point.

Using a reference makes it possible to directly manipulate x. For example:


  QPoint p(1, 2);
  p.rx()--;   // p becomes (0, 2)



See also x() and setX().
*/
func (this *QPointF) Rx() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF2rxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & ry()

/*
Returns a reference to the y coordinate of this point.

Using a reference makes it possible to directly manipulate y. For example:


  QPoint p(1, 2);
  p.ry()++;   // p becomes (1, 3)



See also y() and setY().
*/
func (this *QPointF) Ry() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF2ryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qpoint.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF & operator+=(const QPointF &)

/*

 */
func (this *QPointF) Operator_add_equal(p QPointF_ITF) *QPointF {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF & operator-=(const QPointF &)

/*

 */
func (this *QPointF) Operator_minus_equal(p QPointF_ITF) *QPointF {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF & operator*=(qreal)

/*

 */
func (this *QPointF) Operator_mul_equal(c float64) *QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFmLEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:241
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF & operator/=(qreal)

/*

 */
func (this *QPointF) Operator_div_equal(c float64) *QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFdVEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qpoint.h:243
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] qreal dotProduct(const QPointF &, const QPointF &)

/*
QPoint p( 3, 7);
  QPoint q(-1, 4);
  int lengthSquared = QPoint::dotProduct(p, q);   // lengthSquared becomes 25



Returns the dot product of p1 and p2.

This function was introduced in  Qt 5.1.
*/
func (this *QPointF) DotProduct(p1 QPointF_ITF, p2 QPointF_ITF) float64 {
	var convArg0 unsafe.Pointer
	if p1 != nil && p1.QPointF_PTR() != nil {
		convArg0 = p1.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if p2 != nil && p2.QPointF_PTR() != nil {
		convArg1 = p2.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointF10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QPointF_DotProduct(p1 QPointF_ITF, p2 QPointF_ITF) float64 {
	var nilthis *QPointF
	rv := nilthis.DotProduct(p1, p2)
	return rv
}

// /usr/include/qt/QtCore/qpoint.h:256
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint() const

/*

 */
func (this *QPointF) ToPoint() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPointF7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

func DeleteQPointF(this *QPointF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPointFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10457() {
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
