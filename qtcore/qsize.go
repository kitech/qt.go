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
// extern C begin: 5
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
type QSize struct {
	*qtrt.CObject
}
type QSize_ITF interface {
	QSize_PTR() *QSize
}

func (ptr *QSize) QSize_PTR() *QSize { return ptr }

func (this *QSize) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSize) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSizeFromPointer(cthis unsafe.Pointer) *QSize {
	return &QSize{&qtrt.CObject{cthis}}
}
func (*QSize) NewFromPointer(cthis unsafe.Pointer) *QSize {
	return NewQSizeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsize.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSize()

/*
Constructs a size with an invalid width and height (i.e., isValid() returns false).

See also isValid().
*/
func (*QSize) NewForInherit() *QSize {
	return NewQSize()
}
func NewQSize() *QSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSize)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:56
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSize(int, int)

/*
Constructs a size with an invalid width and height (i.e., isValid() returns false).

See also isValid().
*/
func (*QSize) NewForInherit1(w int, h int) *QSize {
	return NewQSize1(w, h)
}
func NewQSize1(w int, h int) *QSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizeC2Eii", qtrt.FFI_TYPE_POINTER, w, h)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSize)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if both the width and height is 0; otherwise returns false.

See also isValid() and isEmpty().
*/
func (this *QSize) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if either of the width and height is less than or equal to 0; otherwise returns false.

See also isNull() and isValid().
*/
func (this *QSize) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if both the width and height is equal to or greater than 0; otherwise returns false.

See also isNull() and isEmpty().
*/
func (this *QSize) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the width.

See also height() and setWidth().
*/
func (this *QSize) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsize.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height() const

/*
Returns the height.

See also width() and setHeight().
*/
func (this *QSize) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsize.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(int)

/*
Sets the width to the given width.

See also rwidth(), width(), and setHeight().
*/
func (this *QSize) SetWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(int)

/*
Sets the height to the given height.

See also rheight(), height(), and setWidth().
*/
func (this *QSize) SetHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize9setHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void transpose()

/*
Swaps the width and height values.

See also setWidth(), setHeight(), and transposed().
*/
func (this *QSize) Transpose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize9transposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize transposed() const

/*
Returns a QSize with width and height swapped.

This function was introduced in  Qt 5.0.

See also transpose().
*/
func (this *QSize) Transposed() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scale(int, int, Qt::AspectRatioMode)

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
func (this *QSize) Scale(w int, h int, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize5scaleEiiN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:70
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void scale(const QSize &, Qt::AspectRatioMode)

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
func (this *QSize) Scale1(s QSize_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize5scaleERKS_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize scaled(int, int, Qt::AspectRatioMode) const

/*
Return a size scaled to a rectangle with the given width and height, according to the specified mode.

This function was introduced in  Qt 5.0.

See also scale().
*/
func (this *QSize) Scaled(w int, h int, mode int) *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize6scaledEiiN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:72
// index:1
// Public Visibility=Default Availability=Available
// [8] QSize scaled(const QSize &, Qt::AspectRatioMode) const

/*
Return a size scaled to a rectangle with the given width and height, according to the specified mode.

This function was introduced in  Qt 5.0.

See also scale().
*/
func (this *QSize) Scaled1(s QSize_ITF, mode int) *QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize6scaledERKS_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize expandedTo(const QSize &) const

/*
Returns a size holding the maximum width and height of this size and the given otherSize.

See also boundedTo() and scale().
*/
func (this *QSize) ExpandedTo(arg0 QSize_ITF) *QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize10expandedToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize boundedTo(const QSize &) const

/*
Returns a size holding the minimum width and height of this size and the given otherSize.

See also expandedTo() and scale().
*/
func (this *QSize) BoundedTo(arg0 QSize_ITF) *QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QSize9boundedToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int & rwidth()

/*
Returns a reference to the width.

Using a reference makes it possible to manipulate the width directly. For example:


  QSize size(100, 10);
  size.rwidth() += 20;

  // size becomes (120,10)



See also rheight() and setWidth().
*/
func (this *QSize) Rwidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize6rwidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qsize.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int & rheight()

/*
Returns a reference to the height.

Using a reference makes it possible to manipulate the height directly. For example:


  QSize size(100, 10);
  size.rheight() += 5;

  // size becomes (100,15)



See also rwidth() and setHeight().
*/
func (this *QSize) Rheight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSize7rheightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qsize.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize & operator+=(const QSize &)

/*

 */
func (this *QSize) Operator_add_equal(arg0 QSize_ITF) *QSize {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizepLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize & operator-=(const QSize &)

/*

 */
func (this *QSize) Operator_minus_equal(arg0 QSize_ITF) *QSize {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizemIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize & operator*=(qreal)

/*

 */
func (this *QSize) Operator_mul_equal(c float64) *QSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizemLEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize & operator/=(qreal)

/*

 */
func (this *QSize) Operator_div_equal(c float64) *QSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizedVEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

func DeleteQSize(this *QSize) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QSizeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10519() {
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
