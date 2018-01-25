package qtgui

// /usr/include/qt/QtGui/qregion.h
// #include <qregion.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 103
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QRegion struct {
	*qtrt.CObject
}

func (this *QRegion) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegion) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQRegionFromPointer(cthis unsafe.Pointer) *QRegion {
	return &QRegion{&qtrt.CObject{cthis}}
}
func (*QRegion) NewFromPointer(cthis unsafe.Pointer) *QRegion {
	return NewQRegionFromPointer(cthis)
}

// /usr/include/qt/QtGui/qregion.h:67
// index:0
// Public
// void QRegion()
func NewQRegion() *QRegion {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:68
// index:1
// Public
// void QRegion(int, int, int, int, enum QRegion::RegionType)
func NewQRegion_1(x int, y int, w int, h int, t int) *QRegion {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2EiiiiNS_10RegionTypeE", ffiqt.FFI_TYPE_VOID, cthis, x, y, w, h, t)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:69
// index:2
// Public
// void QRegion(const class QRect &, enum QRegion::RegionType)
func NewQRegion_2(r *qtcore.QRect, t int) *QRegion {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2ERK5QRectNS_10RegionTypeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, t)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:70
// index:3
// Public
// void QRegion(const class QPolygon &, Qt::FillRule)
func NewQRegion_3(pa *QPolygon, fillRule int) *QRegion {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pa.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2ERK8QPolygonN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, fillRule)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:74
// index:4
// Public
// void QRegion(const class QBitmap &)
func NewQRegion_4(bitmap *QBitmap) *QRegion {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = bitmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2ERK7QBitmap", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:75
// index:0
// Public
// void ~QRegion()
func DeleteQRegion(*QRegion) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:81
// index:0
// Public inline
// void swap(class QRegion &)
func (this *QRegion) Swap(other *QRegion) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:82
// index:0
// Public
// bool isEmpty()
func (this *QRegion) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:83
// index:0
// Public
// bool isNull()
func (this *QRegion) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:88
// index:0
// Public
// QRegion::const_iterator begin()
func (this *QRegion) Begin() *qtcore.QRect /*444 const QRect **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:89
// index:0
// Public inline
// QRegion::const_iterator cbegin()
func (this *QRegion) Cbegin() *qtcore.QRect /*444 const QRect **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6cbeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:90
// index:0
// Public
// QRegion::const_iterator end()
func (this *QRegion) End() *qtcore.QRect /*444 const QRect **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:91
// index:0
// Public inline
// QRegion::const_iterator cend()
func (this *QRegion) Cend() *qtcore.QRect /*444 const QRect **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion4cendEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:97
// index:0
// Public
// bool contains(const class QPoint &)
func (this *QRegion) Contains(p *qtcore.QPoint) bool {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion8containsERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:98
// index:1
// Public
// bool contains(const class QRect &)
func (this *QRegion) Contains_1(r *qtcore.QRect) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion8containsERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:100
// index:0
// Public
// void translate(int, int)
func (this *QRegion) Translate(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion9translateEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:101
// index:1
// Public inline
// void translate(const class QPoint &)
func (this *QRegion) Translate_1(p *qtcore.QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion9translateERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:102
// index:0
// Public
// QRegion translated(int, int)
func (this *QRegion) Translated(dx int, dy int) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10translatedEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:103
// index:1
// Public inline
// QRegion translated(const class QPoint &)
func (this *QRegion) Translated_1(p *qtcore.QPoint) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10translatedERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:105
// index:0
// Public
// QRegion united(const class QRegion &)
func (this *QRegion) United(r *QRegion) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6unitedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:106
// index:1
// Public
// QRegion united(const class QRect &)
func (this *QRegion) United_1(r *qtcore.QRect) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6unitedERK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:107
// index:0
// Public
// QRegion intersected(const class QRegion &)
func (this *QRegion) Intersected(r *QRegion) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:108
// index:1
// Public
// QRegion intersected(const class QRect &)
func (this *QRegion) Intersected_1(r *qtcore.QRect) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion11intersectedERK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:109
// index:0
// Public
// QRegion subtracted(const class QRegion &)
func (this *QRegion) Subtracted(r *QRegion) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10subtractedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:110
// index:0
// Public
// QRegion xored(const class QRegion &)
func (this *QRegion) Xored(r *QRegion) *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion5xoredERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:121
// index:0
// Public
// bool intersects(const class QRegion &)
func (this *QRegion) Intersects(r *QRegion) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:122
// index:1
// Public
// bool intersects(const class QRect &)
func (this *QRegion) Intersects_1(r *qtcore.QRect) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10intersectsERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qregion.h:124
// index:0
// Public
// QRect boundingRect()
func (this *QRegion) BoundingRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion12boundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qregion.h:126
// index:0
// Public
// void setRects(const class QRect *, int)
func (this *QRegion) SetRects(rect *qtcore.QRect /*444 const QRect **/, num int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion8setRectsEPK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:127
// index:0
// Public
// int rectCount()
func (this *QRegion) RectCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion9rectCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

type QRegion__RegionType = int

const QRegion__Rectangle QRegion__RegionType = 0
const QRegion__Ellipse QRegion__RegionType = 1

//  body block end
