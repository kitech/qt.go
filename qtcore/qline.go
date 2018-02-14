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
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QLine struct {
	*qtrt.CObject
}
type QLine_ITF interface {
	QLine_PTR() *QLine
}

func (ptr *QLine) QLine_PTR() *QLine { return ptr }

func (this *QLine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLineFromPointer(cthis unsafe.Pointer) *QLine {
	return &QLine{&qtrt.CObject{cthis}}
}
func (*QLine) NewFromPointer(cthis unsafe.Pointer) *QLine {
	return NewQLineFromPointer(cthis)
}

// /usr/include/qt/QtCore/qline.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLine()
func NewQLine() *QLine {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLineC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLine)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:56
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QLine(const QPoint &, const QPoint &)
func NewQLine_1(pt1 QPoint_ITF, pt2 QPoint_ITF) *QLine {
	var convArg0 = pt1.QPoint_PTR().GetCthis()
	var convArg1 = pt2.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLineC2ERK6QPointS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLine)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:57
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QLine(int, int, int, int)
func NewQLine_2(x1 int, y1 int, x2 int, y2 int) *QLine {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLineC2Eiiii", qtrt.FFI_TYPE_POINTER, x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLine)
	return gothis
}

// /usr/include/qt/QtCore/qline.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QLine) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qline.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint p1()
func (this *QLine) P1() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2p1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint p2()
func (this *QLine) P2() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2p2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x1()
func (this *QLine) X1() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2x1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qline.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y1()
func (this *QLine) Y1() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2y1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qline.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x2()
func (this *QLine) X2() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2x2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qline.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y2()
func (this *QLine) Y2() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2y2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qline.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int dx()
func (this *QLine) Dx() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2dxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qline.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int dy()
func (this *QLine) Dy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine2dyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qline.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)
func (this *QLine) Translate(p QPoint_ITF) {
	var convArg0 = p.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLine9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(int, int)
func (this *QLine) Translate_1(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLine9translateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLine translated(const QPoint &)
func (this *QLine) Translated(p QPoint_ITF) *QLine /*123*/ {
	var convArg0 = p.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine10translatedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLine)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:77
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QLine translated(int, int)
func (this *QLine) Translated_1(dx int, dy int) *QLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine10translatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLine)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint center()
func (this *QLine) Center() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QLine6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qline.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setP1(const QPoint &)
func (this *QLine) SetP1(p1 QPoint_ITF) {
	var convArg0 = p1.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLine5setP1ERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setP2(const QPoint &)
func (this *QLine) SetP2(p2 QPoint_ITF) {
	var convArg0 = p2.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLine5setP2ERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPoints(const QPoint &, const QPoint &)
func (this *QLine) SetPoints(p1 QPoint_ITF, p2 QPoint_ITF) {
	var convArg0 = p1.QPoint_PTR().GetCthis()
	var convArg1 = p2.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLine9setPointsERK6QPointS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qline.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLine(int, int, int, int)
func (this *QLine) SetLine(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLine7setLineEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	qtrt.ErrPrint(err, rv)
}

func DeleteQLine(this *QLine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QLineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
}

//  keep block end
