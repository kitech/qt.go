package qtgui

// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 70
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
type QPolygon struct {
	*qtrt.CObject
}

func (this *QPolygon) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPolygon) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPolygonFromPointer(cthis unsafe.Pointer) *QPolygon {
	return &QPolygon{&qtrt.CObject{cthis}}
}
func (*QPolygon) NewFromPointer(cthis unsafe.Pointer) *QPolygon {
	return NewQPolygonFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpolygon.h:59
// index:0
// Public inline
// void QPolygon()
func NewQPolygon() *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:61
// index:1
// Public inline
// void QPolygon(int)
func NewQPolygon_1(size int) *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2Ei", ffiqt.FFI_TYPE_VOID, cthis, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:66
// index:2
// Public
// void QPolygon(const QRect &, _Bool)
func NewQPolygon_2(r *qtcore.QRect, closed bool) *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2ERK5QRectb", ffiqt.FFI_TYPE_VOID, cthis, convArg0, closed)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:67
// index:3
// Public
// void QPolygon(int, const int *)
func NewQPolygon_3(nPoints int, points unsafe.Pointer /*666*/) *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2EiPKi", ffiqt.FFI_TYPE_VOID, cthis, nPoints, &points)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:60
// index:0
// Public inline
// void ~QPolygon()
func DeleteQPolygon(*QPolygon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:74
// index:0
// Public inline
// void swap(QPolygon &)
func (this *QPolygon) Swap(other *QPolygon) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:78
// index:0
// Public
// void translate(int, int)
func (this *QPolygon) Translate(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9translateEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:79
// index:1
// Public
// void translate(const QPoint &)
func (this *QPolygon) Translate_1(offset *qtcore.QPoint) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9translateERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:81
// index:0
// Public
// QPolygon translated(int, int)
func (this *QPolygon) Translated(dx int, dy int) *QPolygon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10translatedEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:82
// index:1
// Public inline
// QPolygon translated(const QPoint &)
func (this *QPolygon) Translated_1(offset *qtcore.QPoint) *QPolygon /*123*/ {
	var convArg0 = offset.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10translatedERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:84
// index:0
// Public
// QRect boundingRect()
func (this *QPolygon) BoundingRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon12boundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:86
// index:0
// Public
// void point(int, int *, int *)
func (this *QPolygon) Point(i int, x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon5pointEiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:87
// index:1
// Public
// QPoint point(int)
func (this *QPolygon) Point_1(i int) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon5pointEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:88
// index:0
// Public
// void setPoint(int, int, int)
func (this *QPolygon) SetPoint(index int, x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon8setPointEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:89
// index:1
// Public
// void setPoint(int, const QPoint &)
func (this *QPolygon) SetPoint_1(index int, p *qtcore.QPoint) {
	var convArg1 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon8setPointEiRK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:90
// index:0
// Public
// void setPoints(int, const int *)
func (this *QPolygon) SetPoints(nPoints int, points unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9setPointsEiPKi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nPoints, &points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:92
// index:0
// Public
// void putPoints(int, int, const int *)
func (this *QPolygon) PutPoints(index int, nPoints int, points unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiPKi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, nPoints, &points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:94
// index:1
// Public
// void putPoints(int, int, const QPolygon &, int)
func (this *QPolygon) PutPoints_1(index int, nPoints int, from *QPolygon, fromIndex int) {
	var convArg2 = from.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiRKS_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, nPoints, convArg2, fromIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:96
// index:0
// Public
// bool containsPoint(const QPoint &, Qt::FillRule)
func (this *QPolygon) ContainsPoint(pt *qtcore.QPoint, fillRule int) bool {
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon13containsPointERK6QPointN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpolygon.h:98
// index:0
// Public
// QPolygon united(const QPolygon &)
func (this *QPolygon) United(r *QPolygon) *QPolygon /*123*/ {
	var convArg0 = r.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon6unitedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:99
// index:0
// Public
// QPolygon intersected(const QPolygon &)
func (this *QPolygon) Intersected(r *QPolygon) *QPolygon /*123*/ {
	var convArg0 = r.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:100
// index:0
// Public
// QPolygon subtracted(const QPolygon &)
func (this *QPolygon) Subtracted(r *QPolygon) *QPolygon /*123*/ {
	var convArg0 = r.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10subtractedERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:102
// index:0
// Public
// bool intersects(const QPolygon &)
func (this *QPolygon) Intersects(r *QPolygon) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
