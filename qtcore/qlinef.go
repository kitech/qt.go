package qtcore

// /usr/include/qt/QtCore/qline.h
// #include <qline.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QLineF struct {
	*qtrt.CObject
}

func (this *QLineF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLineF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLineFFromPointer(cthis unsafe.Pointer) *QLineF {
	return &QLineF{&qtrt.CObject{cthis}}
}
func (*QLineF) NewFromPointer(cthis unsafe.Pointer) *QLineF {
	return NewQLineFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qline.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLineF()
func NewQLineF() *QLineF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLineF)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:220
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QLineF(const QPointF &, const QPointF &)
func NewQLineF_1(pt1 *QPointF, pt2 *QPointF) *QLineF {
	var convArg0 = pt1.GetCthis()
	var convArg1 = pt2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineFC2ERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLineF)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:221
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QLineF(qreal, qreal, qreal, qreal)
func NewQLineF_2(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineFC2Edddd", qtrt.FFI_TYPE_POINTER, x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLineF)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:222
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QLineF(const QLine &)
func NewQLineF_3(line *QLine) *QLineF {
	var convArg0 = line.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineFC2ERK5QLine", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLineF)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:224
// index:0
// Public static Visibility=Default Availability=Available
// [32] QLineF fromPolar(qreal, qreal)
func (this *QLineF) FromPolar(length float64, angle float64) *QLineF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF9fromPolarEdd", qtrt.FFI_TYPE_POINTER, length, angle)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}
func QLineF_FromPolar(length float64, angle float64) *QLineF /*123*/ {
	var nilthis *QLineF
	rv := nilthis.FromPolar(length, angle)
	return rv
}

// /usr/include/qt/QtCore/qline.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QLineF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qline.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF p1()
func (this *QLineF) P1() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2p1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF p2()
func (this *QLineF) P2() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2p2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x1()
func (this *QLineF) X1() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2x1Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y1()
func (this *QLineF) Y1() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2y1Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:234
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x2()
func (this *QLineF) X2() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2x2Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y2()
func (this *QLineF) Y2() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2y2Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:237
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal dx()
func (this *QLineF) Dx() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2dxEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal dy()
func (this *QLineF) Dy() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF2dyEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:240
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal length()
func (this *QLineF) Length() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLength(qreal)
func (this *QLineF) SetLength(len float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF9setLengthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:243
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal angle()
func (this *QLineF) Angle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF5angleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:254
// index:1
// Public Visibility=Default Availability=Available
// [8] qreal angle(const QLineF &)
func (this *QLineF) Angle_1(l *QLineF) float64 {
	var convArg0 = l.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF5angleERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:244
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAngle(qreal)
func (this *QLineF) SetAngle(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF8setAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:246
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal angleTo(const QLineF &)
func (this *QLineF) AngleTo(l *QLineF) float64 {
	var convArg0 = l.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF7angleToERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qline.h:248
// index:0
// Public Visibility=Default Availability=Available
// [32] QLineF unitVector()
func (this *QLineF) UnitVector() *QLineF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF10unitVectorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:249
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QLineF normalVector()
func (this *QLineF) NormalVector() *QLineF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF12normalVectorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:252
// index:0
// Public Visibility=Default Availability=Available
// [4] QLineF::IntersectType intersect(const QLineF &, QPointF *)
func (this *QLineF) Intersect(l *QLineF, intersectionPoint *QPointF /*777 QPointF **/) int {
	var convArg0 = l.GetCthis()
	var convArg1 = intersectionPoint.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF9intersectERKS_P7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qline.h:256
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF pointAt(qreal)
func (this *QLineF) PointAt(t float64) *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF7pointAtEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)
func (this *QLineF) Translate(p *QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF9translateERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:258
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QLineF) Translate_1(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:260
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QLineF translated(const QPointF &)
func (this *QLineF) Translated(p *QPointF) *QLineF /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF10translatedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:261
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QLineF translated(qreal, qreal)
func (this *QLineF) Translated_1(dx float64, dy float64) *QLineF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF10translatedEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:263
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF center()
func (this *QLineF) Center() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:265
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setP1(const QPointF &)
func (this *QLineF) SetP1(p1 *QPointF) {
	var convArg0 = p1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF5setP1ERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:266
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setP2(const QPointF &)
func (this *QLineF) SetP2(p2 *QPointF) {
	var convArg0 = p2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF5setP2ERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:267
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPoints(const QPointF &, const QPointF &)
func (this *QLineF) SetPoints(p1 *QPointF, p2 *QPointF) {
	var convArg0 = p1.GetCthis()
	var convArg1 = p2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF9setPointsERK7QPointFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:268
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLine(qreal, qreal, qreal, qreal)
func (this *QLineF) SetLine(x1 float64, y1 float64, x2 float64, y2 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineF7setLineEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLine toLine()
func (this *QLineF) ToLine() *QLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QLineF6toLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLine)
	return rv2
}

func DeleteQLineF(this *QLineF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QLineFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QLineF__IntersectType = int

const QLineF__NoIntersection QLineF__IntersectType = 0
const QLineF__BoundedIntersection QLineF__IntersectType = 1
const QLineF__UnboundedIntersection QLineF__IntersectType = 2

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
		qtrt.KeepMe()
	}
}

//  keep block end
