package qtcore

// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QSizeF struct {
	*qtrt.CObject
}
type QSizeF_ITF interface {
	QSizeF_PTR() *QSizeF
}

func (ptr *QSizeF) QSizeF_PTR() *QSizeF { return ptr }

func (this *QSizeF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSizeF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSizeFFromPointer(cthis unsafe.Pointer) *QSizeF {
	return &QSizeF{&qtrt.CObject{cthis}}
}
func (*QSizeF) NewFromPointer(cthis unsafe.Pointer) *QSizeF {
	return NewQSizeFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsize.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSizeF()

/*

 */
func (*QSizeF) NewForInherit() *QSizeF {
	return NewQSizeF()
}
func NewQSizeF() *QSizeF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizeF)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:219
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSizeF(const QSize &)

/*

 */
func (*QSizeF) NewForInherit1(sz QSize_ITF) *QSizeF {
	return NewQSizeF1(sz)
}
func NewQSizeF1(sz QSize_ITF) *QSizeF {
	var convArg0 unsafe.Pointer
	if sz != nil && sz.QSize_PTR() != nil {
		convArg0 = sz.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizeF)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:220
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QSizeF(qreal, qreal)

/*

 */
func (*QSizeF) NewForInherit2(w float64, h float64) *QSizeF {
	return NewQSizeF2(w, h)
}
func NewQSizeF2(w float64, h float64) *QSizeF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFC2Edd", qtrt.FFI_TYPE_POINTER, w, h)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizeF)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if both the width and height is 0; otherwise returns false.

See also isValid() and isEmpty().
*/
func (this *QSizeF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if either of the width and height is less than or equal to 0; otherwise returns false.

See also isNull() and isValid().
*/
func (this *QSizeF) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if both the width and height is equal to or greater than 0; otherwise returns false.

See also isNull() and isEmpty().
*/
func (this *QSizeF) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width() const

/*
Returns the width.

See also height() and setWidth().
*/
func (this *QSizeF) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qsize.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height() const

/*
Returns the height.

See also width() and setHeight().
*/
func (this *QSizeF) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qsize.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)

/*
Sets the width to the given width.

See also rwidth(), width(), and setHeight().
*/
func (this *QSizeF) SetWidth(w float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)

/*
Sets the height to the given height.

See also rheight(), height(), and setWidth().
*/
func (this *QSizeF) SetHeight(h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void transpose()

/*
Swaps the width and height values.

See also setWidth(), setHeight(), and transposed().
*/
func (this *QSizeF) Transpose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF9transposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF transposed() const

/*
Returns a QSize with width and height swapped.

This function was introduced in  Qt 5.0.

See also transpose().
*/
func (this *QSizeF) Transposed() *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scale(qreal, qreal, Qt::AspectRatioMode)

/*
Scales the size to a rectangle with the given width and height, according to the specified mode:


If mode is Qt::IgnoreAspectRatio, the size is set to (width, height).
If mode is Qt::KeepAspectRatio, the current size is scaled to a rectangle as large as possible inside (width, height), preserving the aspect ratio.
If mode is Qt::KeepAspectRatioByExpanding, the current size is scaled to a rectangle as small as possible outside (width, height), preserving the aspect ratio.


Example:


  QSize t1(10, 12);
  t1.scale(60, 60, Qt::IgnoreAspectRatio);
  // t1 is (60, 60)

  QSize t2(10, 12);
  t2.scale(60, 60, Qt::KeepAspectRatio);
  // t2 is (50, 60)

  QSize t3(10, 12);
  t3.scale(60, 60, Qt::KeepAspectRatioByExpanding);
  // t3 is (60, 72)



See also setWidth(), setHeight(), and scaled().
*/
func (this *QSizeF) Scale(w float64, h float64, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF5scaleEddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:234
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void scale(const QSizeF &, Qt::AspectRatioMode)

/*
Scales the size to a rectangle with the given width and height, according to the specified mode:


If mode is Qt::IgnoreAspectRatio, the size is set to (width, height).
If mode is Qt::KeepAspectRatio, the current size is scaled to a rectangle as large as possible inside (width, height), preserving the aspect ratio.
If mode is Qt::KeepAspectRatioByExpanding, the current size is scaled to a rectangle as small as possible outside (width, height), preserving the aspect ratio.


Example:


  QSize t1(10, 12);
  t1.scale(60, 60, Qt::IgnoreAspectRatio);
  // t1 is (60, 60)

  QSize t2(10, 12);
  t2.scale(60, 60, Qt::KeepAspectRatio);
  // t2 is (50, 60)

  QSize t3(10, 12);
  t3.scale(60, 60, Qt::KeepAspectRatioByExpanding);
  // t3 is (60, 72)



See also setWidth(), setHeight(), and scaled().
*/
func (this *QSizeF) Scale1(s QSizeF_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizeF_PTR() != nil {
		convArg0 = s.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF5scaleERKS_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:235
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF scaled(qreal, qreal, Qt::AspectRatioMode) const

/*
Return a size scaled to a rectangle with the given width and height, according to the specified mode.

This function was introduced in  Qt 5.0.

See also scale().
*/
func (this *QSizeF) Scaled(w float64, h float64, mode int) *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6scaledEddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:236
// index:1
// Public Visibility=Default Availability=Available
// [16] QSizeF scaled(const QSizeF &, Qt::AspectRatioMode) const

/*
Return a size scaled to a rectangle with the given width and height, according to the specified mode.

This function was introduced in  Qt 5.0.

See also scale().
*/
func (this *QSizeF) Scaled1(s QSizeF_ITF, mode int) *QSizeF /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizeF_PTR() != nil {
		convArg0 = s.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6scaledERKS_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF expandedTo(const QSizeF &) const

/*
Returns a size holding the maximum width and height of this size and the given otherSize.

See also boundedTo() and scale().
*/
func (this *QSizeF) ExpandedTo(arg0 QSizeF_ITF) *QSizeF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSizeF_PTR() != nil {
		convArg0 = arg0.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF10expandedToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF boundedTo(const QSizeF &) const

/*
Returns a size holding the minimum width and height of this size and the given otherSize.

See also expandedTo() and scale().
*/
func (this *QSizeF) BoundedTo(arg0 QSizeF_ITF) *QSizeF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSizeF_PTR() != nil {
		convArg0 = arg0.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF9boundedToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:241
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & rwidth()

/*
Returns a reference to the width.

Using a reference makes it possible to manipulate the width directly. For example:


  QSize size(100, 10);
  size.rwidth() += 20;

  // size becomes (120,10)



See also rheight() and setWidth().
*/
func (this *QSizeF) Rwidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF6rwidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qsize.h:242
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & rheight()

/*
Returns a reference to the height.

Using a reference makes it possible to manipulate the height directly. For example:


  QSize size(100, 10);
  size.rheight() += 5;

  // size becomes (100,15)



See also rwidth() and setHeight().
*/
func (this *QSizeF) Rheight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF7rheightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qsize.h:244
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF & operator+=(const QSizeF &)

/*

 */
func (this *QSizeF) Operator_add_equal(arg0 QSizeF_ITF) *QSizeF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSizeF_PTR() != nil {
		convArg0 = arg0.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:245
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF & operator-=(const QSizeF &)

/*

 */
func (this *QSizeF) Operator_minus_equal(arg0 QSizeF_ITF) *QSizeF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSizeF_PTR() != nil {
		convArg0 = arg0.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF & operator*=(qreal)

/*

 */
func (this *QSizeF) Operator_mul_equal(c float64) *QSizeF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFmLEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:247
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF & operator/=(qreal)

/*

 */
func (this *QSizeF) Operator_div_equal(c float64) *QSizeF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFdVEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize toSize() const

/*

 */
func (this *QSizeF) ToSize() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6toSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

func DeleteQSizeF(this *QSizeF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
