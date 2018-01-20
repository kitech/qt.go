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
	*qtrt.CObject
}

func (this *QLineF) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQLineFFromPointer(cthis unsafe.Pointer) *QLineF {
	return &QLineF{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qline.h:219
// index:0
// Public inline
// void QLineF()
func NewQLineF() *QLineF {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:220
// index:1
// Public inline
// void QLineF(const class QPointF &, const class QPointF &)
func NewQLineF_1(pt1 *QPointF, pt2 *QPointF) *QLineF {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = pt1.GetCthis()
	var convArg1 = pt2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2ERK7QPointFS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:221
// index:2
// Public inline
// void QLineF(qreal, qreal, qreal, qreal)
func NewQLineF_2(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2Edddd", ffiqt.FFI_TYPE_VOID, cthis, &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:222
// index:3
// Public inline
// void QLineF(const class QLine &)
func NewQLineF_3(line *QLine) *QLineF {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineFC2ERK5QLine", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:224
// index:0
// Public static
// QLineF fromPolar(qreal, qreal)
func (this *QLineF) FromPolar(length float64, angle float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9fromPolarEdd", ffiqt.FFI_TYPE_POINTER, length, angle)
	gopp.ErrPrint(err, rv)
	return rv
}
func QLineF_FromPolar(length float64, angle float64) {
	var nilthis *QLineF
	nilthis.FromPolar(length, angle)
}

// /usr/include/qt/QtCore/qline.h:226
// index:0
// Public inline
// bool isNull()
func (this *QLineF) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:228
// index:0
// Public inline
// QPointF p1()
func (this *QLineF) P1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2p1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:229
// index:0
// Public inline
// QPointF p2()
func (this *QLineF) P2() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2p2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:231
// index:0
// Public inline
// qreal x1()
func (this *QLineF) X1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2x1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:232
// index:0
// Public inline
// qreal y1()
func (this *QLineF) Y1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2y1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:234
// index:0
// Public inline
// qreal x2()
func (this *QLineF) X2() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2x2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:235
// index:0
// Public inline
// qreal y2()
func (this *QLineF) Y2() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2y2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:237
// index:0
// Public inline
// qreal dx()
func (this *QLineF) Dx() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2dxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:238
// index:0
// Public inline
// qreal dy()
func (this *QLineF) Dy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF2dyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:240
// index:0
// Public
// qreal length()
func (this *QLineF) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:241
// index:0
// Public
// void setLength(qreal)
func (this *QLineF) SetLength(len float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9setLengthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:243
// index:0
// Public
// qreal angle()
func (this *QLineF) Angle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF5angleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:254
// index:1
// Public
// qreal angle(const class QLineF &)
func (this *QLineF) Angle_1(l *QLineF) interface{} {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF5angleERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:244
// index:0
// Public
// void setAngle(qreal)
func (this *QLineF) SetAngle(angle float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF8setAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:246
// index:0
// Public
// qreal angleTo(const class QLineF &)
func (this *QLineF) AngleTo(l *QLineF) interface{} {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF7angleToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:248
// index:0
// Public
// QLineF unitVector()
func (this *QLineF) UnitVector() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF10unitVectorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:249
// index:0
// Public inline
// QLineF normalVector()
func (this *QLineF) NormalVector() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF12normalVectorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:252
// index:0
// Public
// QLineF::IntersectType intersect(const class QLineF &, class QPointF *)
func (this *QLineF) Intersect(l *QLineF, intersectionPoint unsafe.Pointer) interface{} {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF9intersectERKS_P7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, intersectionPoint)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:256
// index:0
// Public inline
// QPointF pointAt(qreal)
func (this *QLineF) PointAt(t float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF7pointAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:257
// index:0
// Public inline
// void translate(const class QPointF &)
func (this *QLineF) Translate(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9translateERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:258
// index:1
// Public inline
// void translate(qreal, qreal)
func (this *QLineF) Translate_1(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9translateEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:260
// index:0
// Public inline
// QLineF translated(const class QPointF &)
func (this *QLineF) Translated(p *QPointF) interface{} {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF10translatedERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:261
// index:1
// Public inline
// QLineF translated(qreal, qreal)
func (this *QLineF) Translated_1(dx float64, dy float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF10translatedEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:263
// index:0
// Public inline
// QPointF center()
func (this *QLineF) Center() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6centerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qline.h:265
// index:0
// Public inline
// void setP1(const class QPointF &)
func (this *QLineF) SetP1(p1 *QPointF) {
	var convArg0 = p1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF5setP1ERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:266
// index:0
// Public inline
// void setP2(const class QPointF &)
func (this *QLineF) SetP2(p2 *QPointF) {
	var convArg0 = p2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF5setP2ERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:267
// index:0
// Public inline
// void setPoints(const class QPointF &, const class QPointF &)
func (this *QLineF) SetPoints(p1 *QPointF, p2 *QPointF) {
	var convArg0 = p1.GetCthis()
	var convArg1 = p2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF9setPointsERK7QPointFS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:268
// index:0
// Public inline
// void setLine(qreal, qreal, qreal, qreal)
func (this *QLineF) SetLine(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QLineF7setLineEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:273
// index:0
// Public inline
// QLine toLine()
func (this *QLineF) ToLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QLineF6toLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
