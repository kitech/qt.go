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
// extern C begin: 109
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
// [-2] void QRegion(int, int, int, int, enum QRegion::RegionType)
func NewQRegion_1(x int, y int, w int, h int, t int) *QRegion {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegionC2EiiiiNS_10RegionTypeE", qtrt.FFI_TYPE_POINTER, x, y, w, h, t)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegion)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QRect &, enum QRegion::RegionType)
func NewQRegion_2(r qtcore.QRect_ITF, t int) *QRegion {
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

// /usr/include/qt/QtGui/qregion.h:70
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QPolygon &, Qt::FillRule)
func NewQRegion_3(pa QPolygon_ITF, fillRule int) *QRegion {
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

// /usr/include/qt/QtGui/qregion.h:74
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QRegion(const QBitmap &)
func NewQRegion_4(bitmap QBitmap_ITF) *QRegion {
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
func (this *QRegion) Operator_equal_1(other unsafe.Pointer /*333*/) *QRegion {
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
// [1] bool isEmpty()
func (this *QRegion) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QRegion) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion::const_iterator begin()
func (this *QRegion) Begin() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRegion::const_iterator cbegin()
func (this *QRegion) Cbegin() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion::const_iterator end()
func (this *QRegion) End() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRegion::const_iterator cend()
func (this *QRegion) Cend() *qtcore.QRect /*777 const QRect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPoint &)
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
// [1] bool contains(const QRect &)
func (this *QRegion) Contains_1(r qtcore.QRect_ITF) bool {
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
func (this *QRegion) Translate(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegion9translateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:101
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)
func (this *QRegion) Translate_1(p qtcore.QPoint_ITF) {
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
// [8] QRegion translated(int, int)
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
// [8] QRegion translated(const QPoint &)
func (this *QRegion) Translated_1(p qtcore.QPoint_ITF) *QRegion /*123*/ {
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
// [8] QRegion united(const QRegion &)
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
// [8] QRegion united(const QRect &)
func (this *QRegion) United_1(r qtcore.QRect_ITF) *QRegion /*123*/ {
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
// [8] QRegion intersected(const QRegion &)
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
// [8] QRegion intersected(const QRect &)
func (this *QRegion) Intersected_1(r qtcore.QRect_ITF) *QRegion /*123*/ {
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
// [8] QRegion subtracted(const QRegion &)
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
// [8] QRegion xored(const QRegion &)
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
// [1] bool intersects(const QRegion &)
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
// [1] bool intersects(const QRect &)
func (this *QRegion) Intersects_1(r qtcore.QRect_ITF) bool {
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
// [16] QRect boundingRect()
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
// [4] int rectCount()
func (this *QRegion) RectCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegion9rectCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qregion.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion operator|(const QRegion &)
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
// [8] QRegion operator+(const QRegion &)
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
// [8] QRegion operator+(const QRect &)
func (this *QRegion) Operator_add_1(r qtcore.QRect_ITF) *QRegion /*123*/ {
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
// [8] QRegion operator&(const QRegion &)
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
// [8] QRegion operator&(const QRect &)
func (this *QRegion) Operator_and_1(r qtcore.QRect_ITF) *QRegion /*123*/ {
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
// [8] QRegion operator-(const QRegion &)
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
// [8] QRegion operator^(const QRegion &)
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
func (this *QRegion) Operator_add_equal_1(r qtcore.QRect_ITF) *QRegion {
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
func (this *QRegion) Operator_and_equal_1(r qtcore.QRect_ITF) *QRegion {
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
// [1] bool operator==(const QRegion &)
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
// [1] bool operator!=(const QRegion &)
func (this *QRegion) Operator_not_equal(r QRegion_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRegion_PTR() != nil {
		convArg0 = r.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegionneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QRegion__RegionType = int

const QRegion__Rectangle QRegion__RegionType = 0
const QRegion__Ellipse QRegion__RegionType = 1

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
