//  header block begin
// /usr/include/qt/QtCore/qline.h
// #include <qline.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QLineF struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qline.h:219
// index:0
// inline
// void QLineF()
func NewQLineF() *QLineF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QLineF{cthis}
}

// /usr/include/qt/QtCore/qline.h:220
// index:1
// inline
// void QLineF(const class QPointF &, const class QPointF &)
func NewQLineF_1(pt1 unsafe.Pointer, pt2 unsafe.Pointer) *QLineF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2ERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, cthis, pt1, pt2)
	gopp.ErrPrint(err, rv)
	return &QLineF{cthis}
}

// /usr/include/qt/QtCore/qline.h:221
// index:2
// inline
// void QLineF(qreal, qreal, qreal, qreal)
func NewQLineF_2(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2Edddd", ffiqt.FFI_TYPE_VOID, cthis, &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
	return &QLineF{cthis}
}

// /usr/include/qt/QtCore/qline.h:222
// index:3
// inline
// void QLineF(const class QLine &)
func NewQLineF_3(line unsafe.Pointer) *QLineF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2ERK5QLine", ffiqt.FFI_TYPE_VOID, cthis, line)
	gopp.ErrPrint(err, rv)
	return &QLineF{cthis}
}

// /usr/include/qt/QtCore/qline.h:224
// index:0
// static
// QLineF fromPolar(qreal, qreal)
func (this *QLineF) FromPolar(length float64, angle float64) {
	// 0: (qreal length, qreal angle), (length, angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9fromPolarEdd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLineF_FromPolar(length float64, angle float64) {
	// 0: (qreal length, qreal angle), (length, angle)
	var nilthis *QLineF
	nilthis.FromPolar(length, angle)
}

// /usr/include/qt/QtCore/qline.h:226
// index:0
// inline
// bool isNull()
func (this *QLineF) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:228
// index:0
// inline
// QPointF p1()
func (this *QLineF) P1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2p1Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:229
// index:0
// inline
// QPointF p2()
func (this *QLineF) P2() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2p2Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:231
// index:0
// inline
// qreal x1()
func (this *QLineF) X1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2x1Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:232
// index:0
// inline
// qreal y1()
func (this *QLineF) Y1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2y1Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:234
// index:0
// inline
// qreal x2()
func (this *QLineF) X2() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2x2Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:235
// index:0
// inline
// qreal y2()
func (this *QLineF) Y2() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2y2Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:237
// index:0
// inline
// qreal dx()
func (this *QLineF) Dx() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2dxEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:238
// index:0
// inline
// qreal dy()
func (this *QLineF) Dy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2dyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:240
// index:0
// qreal length()
func (this *QLineF) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6lengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:241
// index:0
// void setLength(qreal)
func (this *QLineF) SetLength(len float64) {
	// 0: (, qreal len), (&len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9setLengthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:243
// index:0
// qreal angle()
func (this *QLineF) Angle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF5angleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:254
// index:1
// qreal angle(const class QLineF &)
func (this *QLineF) Angle_1(l unsafe.Pointer) {
	// 1: (, const QLineF & l), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF5angleERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:244
// index:0
// void setAngle(qreal)
func (this *QLineF) SetAngle(angle float64) {
	// 0: (, qreal angle), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF8setAngleEd", ffiqt.FFI_TYPE_VOID, this.cthis, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:246
// index:0
// qreal angleTo(const class QLineF &)
func (this *QLineF) AngleTo(l unsafe.Pointer) {
	// 0: (, const QLineF & l), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF7angleToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:248
// index:0
// QLineF unitVector()
func (this *QLineF) UnitVector() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF10unitVectorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:249
// index:0
// inline
// QLineF normalVector()
func (this *QLineF) NormalVector() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF12normalVectorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:252
// index:0
// QLineF::IntersectType intersect(const class QLineF &, class QPointF *)
func (this *QLineF) Intersect(l unsafe.Pointer, intersectionPoint unsafe.Pointer) {
	// 0: (, const QLineF & l, QPointF * intersectionPoint), (l, intersectionPoint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF9intersectERKS_P7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, l, intersectionPoint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:256
// index:0
// inline
// QPointF pointAt(qreal)
func (this *QLineF) PointAt(t float64) {
	// 0: (, qreal t), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF7pointAtEd", ffiqt.FFI_TYPE_VOID, this.cthis, &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:257
// index:0
// inline
// void translate(const class QPointF &)
func (this *QLineF) Translate(p unsafe.Pointer) {
	// 0: (, const QPointF & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9translateERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:258
// index:1
// inline
// void translate(qreal, qreal)
func (this *QLineF) Translate_1(dx float64, dy float64) {
	// 1: (, qreal dx, qreal dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9translateEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:260
// index:0
// inline
// QLineF translated(const class QPointF &)
func (this *QLineF) Translated(p unsafe.Pointer) {
	// 0: (, const QPointF & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF10translatedERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:261
// index:1
// inline
// QLineF translated(qreal, qreal)
func (this *QLineF) Translated_1(dx float64, dy float64) {
	// 1: (, qreal dx, qreal dy), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF10translatedEdd", ffiqt.FFI_TYPE_VOID, this.cthis, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:263
// index:0
// inline
// QPointF center()
func (this *QLineF) Center() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6centerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:265
// index:0
// inline
// void setP1(const class QPointF &)
func (this *QLineF) SetP1(p1 unsafe.Pointer) {
	// 0: (, const QPointF & p1), (p1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF5setP1ERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, p1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:266
// index:0
// inline
// void setP2(const class QPointF &)
func (this *QLineF) SetP2(p2 unsafe.Pointer) {
	// 0: (, const QPointF & p2), (p2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF5setP2ERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, p2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:267
// index:0
// inline
// void setPoints(const class QPointF &, const class QPointF &)
func (this *QLineF) SetPoints(p1 unsafe.Pointer, p2 unsafe.Pointer) {
	// 0: (, const QPointF & p1, const QPointF & p2), (p1, p2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9setPointsERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, this.cthis, p1, p2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:268
// index:0
// inline
// void setLine(qreal, qreal, qreal, qreal)
func (this *QLineF) SetLine(x1 float64, y1 float64, x2 float64, y2 float64) {
	// 0: (, qreal x1, qreal y1, qreal x2, qreal y2), (&x1, &y1, &x2, &y2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF7setLineEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:273
// index:0
// inline
// QLine toLine()
func (this *QLineF) ToLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6toLineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
