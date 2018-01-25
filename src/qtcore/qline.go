package qtcore

// /usr/include/qt/QtCore/qline.h
// #include <qline.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QLine struct {
	*qtrt.CObject
}

func (this *QLine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLine) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLineFromPointer(cthis unsafe.Pointer) *QLine {
	return &QLine{&qtrt.CObject{cthis}}
}
func (*QLine) NewFromPointer(cthis unsafe.Pointer) *QLine {
	return NewQLineFromPointer(cthis)
}

// /usr/include/qt/QtCore/qline.h:55
// index:0
// Public inline
// void QLine()
func NewQLine() *QLine {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLineC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:56
// index:1
// Public inline
// void QLine(const class QPoint &, const class QPoint &)
func NewQLine_1(pt1 *QPoint, pt2 *QPoint) *QLine {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = pt1.GetCthis()
	var convArg1 = pt2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLineC2ERK6QPointS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:57
// index:2
// Public inline
// void QLine(int, int, int, int)
func NewQLine_2(x1 int, y1 int, x2 int, y2 int) *QLine {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLineC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:59
// index:0
// Public inline
// bool isNull()
func (this *QLine) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qline.h:61
// index:0
// Public inline
// QPoint p1()
func (this *QLine) P1() *QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2p1Ev", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qline.h:62
// index:0
// Public inline
// QPoint p2()
func (this *QLine) P2() *QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2p2Ev", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qline.h:64
// index:0
// Public inline
// int x1()
func (this *QLine) X1() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2x1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qline.h:65
// index:0
// Public inline
// int y1()
func (this *QLine) Y1() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2y1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qline.h:67
// index:0
// Public inline
// int x2()
func (this *QLine) X2() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2x2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qline.h:68
// index:0
// Public inline
// int y2()
func (this *QLine) Y2() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2y2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qline.h:70
// index:0
// Public inline
// int dx()
func (this *QLine) Dx() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2dxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qline.h:71
// index:0
// Public inline
// int dy()
func (this *QLine) Dy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine2dyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qline.h:73
// index:0
// Public inline
// void translate(const class QPoint &)
func (this *QLine) Translate(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLine9translateERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:74
// index:1
// Public inline
// void translate(int, int)
func (this *QLine) Translate_1(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLine9translateEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:76
// index:0
// Public inline
// QLine translated(const class QPoint &)
func (this *QLine) Translated(p *QPoint) *QLine /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine10translatedERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qline.h:77
// index:1
// Public inline
// QLine translated(int, int)
func (this *QLine) Translated_1(dx int, dy int) *QLine /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine10translatedEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qline.h:79
// index:0
// Public inline
// QPoint center()
func (this *QLine) Center() *QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QLine6centerEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qline.h:81
// index:0
// Public inline
// void setP1(const class QPoint &)
func (this *QLine) SetP1(p1 *QPoint) {
	var convArg0 = p1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLine5setP1ERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:82
// index:0
// Public inline
// void setP2(const class QPoint &)
func (this *QLine) SetP2(p2 *QPoint) {
	var convArg0 = p2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLine5setP2ERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:83
// index:0
// Public inline
// void setPoints(const class QPoint &, const class QPoint &)
func (this *QLine) SetPoints(p1 *QPoint, p2 *QPoint) {
	var convArg0 = p1.GetCthis()
	var convArg1 = p2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLine9setPointsERK6QPointS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:84
// index:0
// Public inline
// void setLine(int, int, int, int)
func (this *QLine) SetLine(x1 int, y1 int, x2 int, y2 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QLine7setLineEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

//  body block end
