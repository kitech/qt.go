//  header block begin
// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QPolygonF struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpolygon.h:144
// index:0
// inline
// void QPolygonF()
func NewQPolygonF() *QPolygonF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPolygonF{cthis}
}

// /usr/include/qt/QtGui/qpolygon.h:146
// index:1
// inline
// void QPolygonF(int)
func NewQPolygonF_1(size int) *QPolygonF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &size)
	gopp.ErrPrint(err, rv)
	return &QPolygonF{cthis}
}

// /usr/include/qt/QtGui/qpolygon.h:151
// index:2
// void QPolygonF(const class QRectF &)
func NewQPolygonF_2(r unsafe.Pointer) *QPolygonF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2ERK6QRectF", ffiqt.FFI_TYPE_VOID, cthis, r)
	gopp.ErrPrint(err, rv)
	return &QPolygonF{cthis}
}

// /usr/include/qt/QtGui/qpolygon.h:152
// index:3
// void QPolygonF(const class QPolygon &)
func NewQPolygonF_3(a unsafe.Pointer) *QPolygonF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2ERK8QPolygon", ffiqt.FFI_TYPE_VOID, cthis, a)
	gopp.ErrPrint(err, rv)
	return &QPolygonF{cthis}
}

// /usr/include/qt/QtGui/qpolygon.h:145
// index:0
// inline
// void ~QPolygonF()
func DeleteQPolygonF(*QPolygonF) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:159
// index:0
// inline
// void swap(class QPolygonF &)
func (this *QPolygonF) Swap(other unsafe.Pointer) {
	// 0: (, QPolygonF & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonF4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:163
// index:0
// inline
// void translate(qreal, qreal)
func (this *QPolygonF) Translate(dx float64, dy float64) {
	// 0: (, qreal dx, qreal dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonF9translateEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:164
// index:1
// void translate(const class QPointF &)
func (this *QPolygonF) Translate_1(offset unsafe.Pointer) {
	// 1: (, const QPointF & offset), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonF9translateERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:166
// index:0
// inline
// QPolygonF translated(qreal, qreal)
func (this *QPolygonF) Translated(dx float64, dy float64) {
	// 0: (, qreal dx, qreal dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10translatedEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:167
// index:1
// QPolygonF translated(const class QPointF &)
func (this *QPolygonF) Translated_1(offset unsafe.Pointer) {
	// 1: (, const QPointF & offset), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10translatedERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:169
// index:0
// QPolygon toPolygon()
func (this *QPolygonF) ToPolygon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF9toPolygonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:171
// index:0
// inline
// bool isClosed()
func (this *QPolygonF) IsClosed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF8isClosedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:173
// index:0
// QRectF boundingRect()
func (this *QPolygonF) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:175
// index:0
// bool containsPoint(const class QPointF &, Qt::FillRule)
func (this *QPolygonF) ContainsPoint(pt unsafe.Pointer, fillRule int) {
	// 0: (, const QPointF & pt, Qt::FillRule fillRule), (pt, &fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF13containsPointERK7QPointFN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.cthis, pt, &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:177
// index:0
// QPolygonF united(const class QPolygonF &)
func (this *QPolygonF) United(r unsafe.Pointer) {
	// 0: (, const QPolygonF & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF6unitedERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:178
// index:0
// QPolygonF intersected(const class QPolygonF &)
func (this *QPolygonF) Intersected(r unsafe.Pointer) {
	// 0: (, const QPolygonF & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF11intersectedERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:179
// index:0
// QPolygonF subtracted(const class QPolygonF &)
func (this *QPolygonF) Subtracted(r unsafe.Pointer) {
	// 0: (, const QPolygonF & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10subtractedERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:181
// index:0
// bool intersects(const class QPolygonF &)
func (this *QPolygonF) Intersects(r unsafe.Pointer) {
	// 0: (, const QPolygonF & r), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10intersectsERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, r)
	gopp.ErrPrint(err, rv)
}

//  body block end
