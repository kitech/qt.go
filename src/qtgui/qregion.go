//  header block begin
// /usr/include/qt/QtGui/qregion.h
// #include <qregion.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 104
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
	return this.Cthis
}

// /usr/include/qt/QtGui/qregion.h:67
// index:0
// void QRegion()
func NewQRegion() *QRegion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}
func NewQRegionFromPointer(cthis unsafe.Pointer) *QRegion {
	return &QRegion{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qregion.h:68
// index:1
// void QRegion(int, int, int, int, enum QRegion::RegionType)
func NewQRegion_1(x int, y int, w int, h int, t int) *QRegion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2EiiiiNS_10RegionTypeE", ffiqt.FFI_TYPE_VOID, cthis, &x, &y, &w, &h, &t)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:69
// index:2
// void QRegion(const class QRect &, enum QRegion::RegionType)
func NewQRegion_2(r unsafe.Pointer, t int) *QRegion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2ERK5QRectNS_10RegionTypeE", ffiqt.FFI_TYPE_VOID, cthis, r, &t)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:70
// index:3
// void QRegion(const class QPolygon &, Qt::FillRule)
func NewQRegion_3(pa unsafe.Pointer, fillRule int) *QRegion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2ERK8QPolygonN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, cthis, pa, &fillRule)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:74
// index:4
// void QRegion(const class QBitmap &)
func NewQRegion_4(bitmap unsafe.Pointer) *QRegion {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionC2ERK7QBitmap", ffiqt.FFI_TYPE_VOID, cthis, bitmap)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qregion.h:75
// index:0
// void ~QRegion()
func DeleteQRegion(*QRegion) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:81
// index:0
// inline
// void swap(class QRegion &)
func (this *QRegion) Swap(other unsafe.Pointer) {
	// 0: (, other QRegion &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:82
// index:0
// bool isEmpty()
func (this *QRegion) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:83
// index:0
// bool isNull()
func (this *QRegion) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:88
// index:0
// QRegion::const_iterator begin()
func (this *QRegion) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:89
// index:0
// inline
// QRegion::const_iterator cbegin()
func (this *QRegion) Cbegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6cbeginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:90
// index:0
// QRegion::const_iterator end()
func (this *QRegion) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:91
// index:0
// inline
// QRegion::const_iterator cend()
func (this *QRegion) Cend() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion4cendEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:97
// index:0
// bool contains(const class QPoint &)
func (this *QRegion) Contains(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion8containsERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:98
// index:1
// bool contains(const class QRect &)
func (this *QRegion) Contains_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion8containsERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:100
// index:0
// void translate(int, int)
func (this *QRegion) Translate(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion9translateEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:101
// index:1
// inline
// void translate(const class QPoint &)
func (this *QRegion) Translate_1(p unsafe.Pointer) {
	// 1: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion9translateERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:102
// index:0
// QRegion translated(int, int)
func (this *QRegion) Translated(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10translatedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:103
// index:1
// inline
// QRegion translated(const class QPoint &)
func (this *QRegion) Translated_1(p unsafe.Pointer) {
	// 1: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10translatedERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:105
// index:0
// QRegion united(const class QRegion &)
func (this *QRegion) United(r unsafe.Pointer) {
	// 0: (, r const QRegion &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6unitedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:106
// index:1
// QRegion united(const class QRect &)
func (this *QRegion) United_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion6unitedERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:107
// index:0
// QRegion intersected(const class QRegion &)
func (this *QRegion) Intersected(r unsafe.Pointer) {
	// 0: (, r const QRegion &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion11intersectedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:108
// index:1
// QRegion intersected(const class QRect &)
func (this *QRegion) Intersected_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion11intersectedERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:109
// index:0
// QRegion subtracted(const class QRegion &)
func (this *QRegion) Subtracted(r unsafe.Pointer) {
	// 0: (, r const QRegion &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10subtractedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:110
// index:0
// QRegion xored(const class QRegion &)
func (this *QRegion) Xored(r unsafe.Pointer) {
	// 0: (, r const QRegion &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion5xoredERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:121
// index:0
// bool intersects(const class QRegion &)
func (this *QRegion) Intersects(r unsafe.Pointer) {
	// 0: (, r const QRegion &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10intersectsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:122
// index:1
// bool intersects(const class QRect &)
func (this *QRegion) Intersects_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion10intersectsERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:124
// index:0
// QRect boundingRect()
func (this *QRegion) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:125
// index:0
// QVector<QRect> rects()
func (this *QRegion) Rects() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion5rectsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:126
// index:0
// void setRects(const class QRect *, int)
func (this *QRegion) SetRects(rect unsafe.Pointer, num int) {
	// 0: (, rect const QRect *, num int), (rect, &num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegion8setRectsEPK5QRecti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:127
// index:0
// int rectCount()
func (this *QRegion) RectCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegion9rectCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
