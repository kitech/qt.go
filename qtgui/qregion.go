package qtgui

// /usr/include/qt/QtGui/qregion.h
// #include <qregion.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 106
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
type QRegion struct {
	*qtrt.CObject
}
type QRegion_ITF interface {
	QRegion_PTR() *QRegion
}

func (ptr *QRegion) QRegion_PTR() *QRegion { return ptr }

func (this *QRegion) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegion) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegionFromPointer(cthis unsafe.Pointer) *QRegion {
	return &QRegion{&qtrt.CObject{cthis}}
}
func (*QRegion) NewFromPointer(cthis unsafe.Pointer) *QRegion {
	return NewQRegionFromPointer(cthis)
}

// /usr/include/qt/QtGui/qregion.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegion()

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit() *QRegion {
	return NewQRegion()
}
func NewQRegion() *QRegion {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegion(int, int, int, int, QRegion::RegionType)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit1(x int, y int, w int, h int, t int) *QRegion {
	return NewQRegion1(x, y, w, h, t)
}
func NewQRegion1(x int, y int, w int, h int, t int) *QRegion {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2EiiiiNS_10RegionTypeE", qtrt.FFI_TYPE_POINTER, x, y, w, h, t)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegion(int, int, int, int, QRegion::RegionType)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit1p(x int, y int, w int, h int) *QRegion {
	return NewQRegion1p(x, y, w, h)
}
func NewQRegion1p(x int, y int, w int, h int) *QRegion {
	// arg: 4, QRegion::RegionType=Enum, QRegion::RegionType=Enum, , Invalid
	t := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2EiiiiNS_10RegionTypeE", qtrt.FFI_TYPE_POINTER, x, y, w, h, t)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QRect &, QRegion::RegionType)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit2(r qtcore.QRect_ITF, t int) *QRegion {
	return NewQRegion2(r, t)
}
func NewQRegion2(r qtcore.QRect_ITF, t int) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2ERK5QRectNS_10RegionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, t)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QRect &, QRegion::RegionType)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit2p(r qtcore.QRect_ITF) *QRegion {
	return NewQRegion2p(r)
}
func NewQRegion2p(r qtcore.QRect_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	// arg: 1, QRegion::RegionType=Enum, QRegion::RegionType=Enum, , Invalid
	t := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2ERK5QRectNS_10RegionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, t)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:70
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QPolygon &, Qt::FillRule)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit3(pa QPolygon_ITF, fillRule int) *QRegion {
	return NewQRegion3(pa, fillRule)
}
func NewQRegion3(pa QPolygon_ITF, fillRule int) *QRegion {
	var convArg0 unsafe.Pointer
	if pa != nil && pa.QPolygon_PTR() != nil {
		convArg0 = pa.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2ERK8QPolygonN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:70
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QPolygon &, Qt::FillRule)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit3p(pa QPolygon_ITF) *QRegion {
	return NewQRegion3p(pa)
}
func NewQRegion3p(pa QPolygon_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if pa != nil && pa.QPolygon_PTR() != nil {
		convArg0 = pa.QPolygon_PTR().GetCthis()
	}
	// arg: 1, Qt::FillRule=Elaborated, Qt::FillRule=Enum, , Invalid
	fillRule := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2ERK8QPolygonN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:74
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QBitmap &)

/*
Constructs an empty region.

See also isEmpty().
*/
func (*QRegion) NewForInherit4(bitmap QBitmap_ITF) *QRegion {
	return NewQRegion4(bitmap)
}
func NewQRegion4(bitmap QBitmap_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if bitmap != nil && bitmap.QBitmap_PTR() != nil {
		convArg0 = bitmap.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2ERK7QBitmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegion()

/*

 */
func DeleteQRegion(this *QRegion) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qregion.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion & operator=(const QRegion &)

/*

 */
func (this *QRegion) Operator_equal(arg0 QRegion_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:78
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QRegion & operator=(QRegion &&)

/*

 */
func (this *QRegion) Operator_equal1(other unsafe.Pointer /*333*/) *QRegion {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegion &)

/*
Swaps region other with this region. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QRegion) Swap(other QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRegion_PTR() != nil {
		convArg0 = other.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegion4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the region is empty; otherwise returns false. An empty region is a region that contains no points.

Example:


  QRegion r1(10, 10, 20, 20);
  r1.isEmpty();               // false

  QRegion r3;
  r3.isEmpty();               // true

  QRegion r2(40, 40, 20, 20);
  r3 = r1.intersected(r2);    // r3: intersection of r1 and r2
  r3.isEmpty();               // true

  r3 = r1.united(r2);         // r3: union of r1 and r2
  r3.isEmpty();               // false
*/
func (this *QRegion) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the region is empty; otherwise returns false. An empty region is a region that contains no points. This function is the same as isEmpty

This function was introduced in  Qt 5.0.

See also isEmpty().
*/
func (this *QRegion) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion::const_iterator begin() const

/*
Returns a const_iterator pointing to the beginning of the range of rectangles that make up this range, in the order in which rects() returns them.

This function was introduced in  Qt 5.8.

See also rbegin(), cbegin(), and end().
*/
func (this *QRegion) Begin() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRegion::const_iterator cbegin() const

/*
Same as begin().

This function was introduced in  Qt 5.8.
*/
func (this *QRegion) Cbegin() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion::const_iterator end() const

/*
Returns a const_iterator pointing to one past the end of the range of rectangles that make up this range, in the order in which rects() returns them.

This function was introduced in  Qt 5.8.

See also rend(), cend(), and begin().
*/
func (this *QRegion) End() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRegion::const_iterator cend() const

/*
Same as end().

This function was introduced in  Qt 5.8.
*/
func (this *QRegion) Cend() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPoint &) const

/*
Returns true if the region contains the point p; otherwise returns false.
*/
func (this *QRegion) Contains(p qtcore.QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion8containsERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:98
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRect &) const

/*
Returns true if the region contains the point p; otherwise returns false.
*/
func (this *QRegion) Contains1(r qtcore.QRect_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion8containsERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(int, int)

/*
Translates (moves) the region dx along the X axis and dy along the Y axis.
*/
func (this *QRegion) Translate(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegion9translateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:101
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)

/*
Translates (moves) the region dx along the X axis and dy along the Y axis.
*/
func (this *QRegion) Translate1(p qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegion9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion translated(int, int) const

/*
Returns a copy of the region that is translated dx along the x axis and dy along the y axis, relative to the current position. Positive values move the region to the right and down.

This function was introduced in  Qt 4.1.

See also translate().
*/
func (this *QRegion) Translated(dx int, dy int) *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion10translatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:103
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QRegion translated(const QPoint &) const

/*
Returns a copy of the region that is translated dx along the x axis and dy along the y axis, relative to the current position. Positive values move the region to the right and down.

This function was introduced in  Qt 4.1.

See also translate().
*/
func (this *QRegion) Translated1(p qtcore.QPoint_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion10translatedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion united(const QRegion &) const

/*
Returns a region which is the union of this region and r.



The figure shows the union of two elliptical regions.

This function was introduced in  Qt 4.2.

See also intersected(), subtracted(), and xored().
*/
func (this *QRegion) United(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:106
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegion united(const QRect &) const

/*
Returns a region which is the union of this region and r.



The figure shows the union of two elliptical regions.

This function was introduced in  Qt 4.2.

See also intersected(), subtracted(), and xored().
*/
func (this *QRegion) United1(r qtcore.QRect_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion6unitedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion intersected(const QRegion &) const

/*
Returns a region which is the intersection of this region and r.



The figure shows the intersection of two elliptical regions.

This function was introduced in  Qt 4.2.

See also subtracted(), united(), and xored().
*/
func (this *QRegion) Intersected(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:108
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegion intersected(const QRect &) const

/*
Returns a region which is the intersection of this region and r.



The figure shows the intersection of two elliptical regions.

This function was introduced in  Qt 4.2.

See also subtracted(), united(), and xored().
*/
func (this *QRegion) Intersected1(r qtcore.QRect_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion11intersectedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion subtracted(const QRegion &) const

/*
Returns a region which is r subtracted from this region.



The figure shows the result when the ellipse on the right is subtracted from the ellipse on the left (left - right).

This function was introduced in  Qt 4.2.

See also intersected(), united(), and xored().
*/
func (this *QRegion) Subtracted(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion10subtractedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion xored(const QRegion &) const

/*
Returns a region which is the exclusive or (XOR) of this region and r.



The figure shows the exclusive or of two elliptical regions.

This function was introduced in  Qt 4.2.

See also intersected(), united(), and subtracted().
*/
func (this *QRegion) Xored(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion5xoredERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRegion &) const

/*
Returns true if this region intersects with region, otherwise returns false.

This function was introduced in  Qt 4.2.
*/
func (this *QRegion) Intersects(r QRegion_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:122
// index:1
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRect &) const

/*
Returns true if this region intersects with region, otherwise returns false.

This function was introduced in  Qt 4.2.
*/
func (this *QRegion) Intersects1(r qtcore.QRect_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion10intersectsERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:124
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect() const

/*
Returns the bounding rectangle of this region. An empty region gives a rectangle that is QRect::isNull().
*/
func (this *QRegion) BoundingRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRects(const QRect *, int)

/*
Sets the region using the array of rectangles specified by rects and number. The rectangles must be optimally Y-X sorted and follow these restrictions:


The rectangles must not intersect.
All rectangles with a given top coordinate must have the same height.
No two rectangles may abut horizontally (they should be combined into a single wider rectangle in that case).
The rectangles must be sorted in ascending order, with Y as the major sort key and X as the minor sort key.


See also rects().
*/
func (this *QRegion) SetRects(rect qtcore.QRect_ITF /*777 const QRect **/, num int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegion8setRectsEPK5QRecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int rectCount() const

/*
Returns the number of rectangles that will be returned in rects().

This function was introduced in  Qt 4.6.
*/
func (this *QRegion) RectCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion9rectCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qregion.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion operator|(const QRegion &) const

/*

 */
func (this *QRegion) Operator_or(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionorERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion operator+(const QRegion &) const

/*

 */
func (this *QRegion) Operator_add(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionplERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:140
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegion operator+(const QRect &) const

/*

 */
func (this *QRegion) Operator_add1(r qtcore.QRect_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionplERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion operator&(const QRegion &) const

/*

 */
func (this *QRegion) Operator_and(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionanERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:142
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegion operator&(const QRect &) const

/*

 */
func (this *QRegion) Operator_and1(r qtcore.QRect_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionanERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion operator-(const QRegion &) const

/*
Applies the subtracted() function to this region and r. r1-r2 is equivalent to r1.subtracted(r2).

See also subtracted().
*/
func (this *QRegion) Operator_minus(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionmiERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion operator^(const QRegion &) const

/*

 */
func (this *QRegion) Operator_caret(r QRegion_ITF) *QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegioneoERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion & operator|=(const QRegion &)

/*

 */
func (this *QRegion) Operator_or_equal(r QRegion_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionoRERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion & operator+=(const QRegion &)

/*

 */
func (this *QRegion) Operator_add_equal(r QRegion_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:148
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegion & operator+=(const QRect &)

/*

 */
func (this *QRegion) Operator_add_equal1(r qtcore.QRect_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionpLERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion & operator&=(const QRegion &)

/*

 */
func (this *QRegion) Operator_and_equal(r QRegion_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionaNERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:150
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegion & operator&=(const QRect &)

/*

 */
func (this *QRegion) Operator_and_equal1(r qtcore.QRect_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionaNERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion & operator-=(const QRegion &)

/*

 */
func (this *QRegion) Operator_minus_equal(r QRegion_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion & operator^=(const QRegion &)

/*

 */
func (this *QRegion) Operator_caret_equal(r QRegion_ITF) *QRegion {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegioneOERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QRegion &) const

/*

 */
func (this *QRegion) Operator_equal_equal(r QRegion_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegioneqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:155
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QRegion &) const

/*

 */
func (this *QRegion) Operator_not_equal(r QRegion_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
Specifies the shape of the region to be created.


*/
type QRegion__RegionType = int

// the region covers the entire rectangle.
const QRegion__Rectangle QRegion__RegionType = 0

// the region is an ellipse inside the rectangle.
const QRegion__Ellipse QRegion__RegionType = 1

func (this *QRegion) RegionTypeItemName(val int) string {
	switch val {
	case QRegion__Rectangle: // 0
		return "Rectangle"
	case QRegion__Ellipse: // 1
		return "Ellipse"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QRegion_RegionTypeItemName(val int) string {
	var nilthis *QRegion
	return nilthis.RegionTypeItemName(val)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
