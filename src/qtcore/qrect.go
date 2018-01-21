package qtcore

// /usr/include/qt/QtCore/qrect.h
// #include <qrect.h>
// #include <QtCore>

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
type QRect struct {
	*qtrt.CObject
}

func (this *QRect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQRectFromPointer(cthis unsafe.Pointer) *QRect {
	return &QRect{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qrect.h:60
// index:0
// Public inline
// void QRect()
func NewQRect() *QRect {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRectC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:61
// index:1
// Public inline
// void QRect(const class QPoint &, const class QPoint &)
func NewQRect_1(topleft *QPoint, bottomright *QPoint) *QRect {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = topleft.GetCthis()
	var convArg1 = bottomright.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRectC2ERK6QPointS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:62
// index:2
// Public inline
// void QRect(const class QPoint &, const class QSize &)
func NewQRect_2(topleft *QPoint, size *QSize) *QRect {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = topleft.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRectC2ERK6QPointRK5QSize", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:63
// index:3
// Public inline
// void QRect(int, int, int, int)
func NewQRect_3(left int, top int, width int, height int) *QRect {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRectC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &left, &top, &width, &height)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:65
// index:0
// Public inline
// bool isNull()
func (this *QRect) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:66
// index:0
// Public inline
// bool isEmpty()
func (this *QRect) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:67
// index:0
// Public inline
// bool isValid()
func (this *QRect) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:69
// index:0
// Public inline
// int left()
func (this *QRect) Left() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect4leftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:70
// index:0
// Public inline
// int top()
func (this *QRect) Top() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect3topEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:71
// index:0
// Public inline
// int right()
func (this *QRect) Right() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect5rightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:72
// index:0
// Public inline
// int bottom()
func (this *QRect) Bottom() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect6bottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:73
// index:0
// Public
// QRect normalized()
func (this *QRect) Normalized() *QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:75
// index:0
// Public inline
// int x()
func (this *QRect) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:76
// index:0
// Public inline
// int y()
func (this *QRect) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:77
// index:0
// Public inline
// void setLeft(int)
func (this *QRect) SetLeft(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect7setLeftEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:78
// index:0
// Public inline
// void setTop(int)
func (this *QRect) SetTop(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect6setTopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:79
// index:0
// Public inline
// void setRight(int)
func (this *QRect) SetRight(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect8setRightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:80
// index:0
// Public inline
// void setBottom(int)
func (this *QRect) SetBottom(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect9setBottomEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:81
// index:0
// Public inline
// void setX(int)
func (this *QRect) SetX(x int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect4setXEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:82
// index:0
// Public inline
// void setY(int)
func (this *QRect) SetY(y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect4setYEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:84
// index:0
// Public inline
// void setTopLeft(const class QPoint &)
func (this *QRect) SetTopLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect10setTopLeftERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:85
// index:0
// Public inline
// void setBottomRight(const class QPoint &)
func (this *QRect) SetBottomRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect14setBottomRightERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:86
// index:0
// Public inline
// void setTopRight(const class QPoint &)
func (this *QRect) SetTopRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect11setTopRightERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:87
// index:0
// Public inline
// void setBottomLeft(const class QPoint &)
func (this *QRect) SetBottomLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect13setBottomLeftERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:89
// index:0
// Public inline
// QPoint topLeft()
func (this *QRect) TopLeft() *QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect7topLeftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:90
// index:0
// Public inline
// QPoint bottomRight()
func (this *QRect) BottomRight() *QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect11bottomRightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:91
// index:0
// Public inline
// QPoint topRight()
func (this *QRect) TopRight() *QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect8topRightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:92
// index:0
// Public inline
// QPoint bottomLeft()
func (this *QRect) BottomLeft() *QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect10bottomLeftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:93
// index:0
// Public inline
// QPoint center()
func (this *QRect) Center() *QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect6centerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:95
// index:0
// Public inline
// void moveLeft(int)
func (this *QRect) MoveLeft(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect8moveLeftEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:96
// index:0
// Public inline
// void moveTop(int)
func (this *QRect) MoveTop(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect7moveTopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:97
// index:0
// Public inline
// void moveRight(int)
func (this *QRect) MoveRight(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect9moveRightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:98
// index:0
// Public inline
// void moveBottom(int)
func (this *QRect) MoveBottom(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect10moveBottomEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:99
// index:0
// Public inline
// void moveTopLeft(const class QPoint &)
func (this *QRect) MoveTopLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect11moveTopLeftERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:100
// index:0
// Public inline
// void moveBottomRight(const class QPoint &)
func (this *QRect) MoveBottomRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect15moveBottomRightERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:101
// index:0
// Public inline
// void moveTopRight(const class QPoint &)
func (this *QRect) MoveTopRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect12moveTopRightERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:102
// index:0
// Public inline
// void moveBottomLeft(const class QPoint &)
func (this *QRect) MoveBottomLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect14moveBottomLeftERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:103
// index:0
// Public inline
// void moveCenter(const class QPoint &)
func (this *QRect) MoveCenter(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect10moveCenterERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:105
// index:0
// Public inline
// void translate(int, int)
func (this *QRect) Translate(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect9translateEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:106
// index:1
// Public inline
// void translate(const class QPoint &)
func (this *QRect) Translate_1(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect9translateERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:107
// index:0
// Public inline
// QRect translated(int, int)
func (this *QRect) Translated(dx int, dy int) *QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect10translatedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:108
// index:1
// Public inline
// QRect translated(const class QPoint &)
func (this *QRect) Translated_1(p *QPoint) *QRect /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect10translatedERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:109
// index:0
// Public inline
// QRect transposed()
func (this *QRect) Transposed() *QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect10transposedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:111
// index:0
// Public inline
// void moveTo(int, int)
func (this *QRect) MoveTo(x int, t int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect6moveToEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:112
// index:1
// Public inline
// void moveTo(const class QPoint &)
func (this *QRect) MoveTo_1(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect6moveToERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:114
// index:0
// Public inline
// void setRect(int, int, int, int)
func (this *QRect) SetRect(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect7setRectEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:115
// index:0
// Public inline
// void getRect(int *, int *, int *, int *)
func (this *QRect) GetRect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect7getRectEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:117
// index:0
// Public inline
// void setCoords(int, int, int, int)
func (this *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect9setCoordsEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:118
// index:0
// Public inline
// void getCoords(int *, int *, int *, int *)
func (this *QRect) GetCoords(x1 unsafe.Pointer /*666*/, y1 unsafe.Pointer /*666*/, x2 unsafe.Pointer /*666*/, y2 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect9getCoordsEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:120
// index:0
// Public inline
// void adjust(int, int, int, int)
func (this *QRect) Adjust(x1 int, y1 int, x2 int, y2 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect6adjustEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:121
// index:0
// Public inline
// QRect adjusted(int, int, int, int)
func (this *QRect) Adjusted(x1 int, y1 int, x2 int, y2 int) *QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect8adjustedEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:123
// index:0
// Public inline
// QSize size()
func (this *QRect) Size() *QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:124
// index:0
// Public inline
// int width()
func (this *QRect) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:125
// index:0
// Public inline
// int height()
func (this *QRect) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrect.h:126
// index:0
// Public inline
// void setWidth(int)
func (this *QRect) SetWidth(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect8setWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:127
// index:0
// Public inline
// void setHeight(int)
func (this *QRect) SetHeight(h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect9setHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:128
// index:0
// Public inline
// void setSize(const class QSize &)
func (this *QRect) SetSize(s *QSize) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QRect7setSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:135
// index:0
// Public
// bool contains(const class QRect &, _Bool)
func (this *QRect) Contains(r *QRect, proper bool) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect8containsERKS_b", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &proper)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:136
// index:1
// Public
// bool contains(const class QPoint &, _Bool)
func (this *QRect) Contains_1(p *QPoint, proper bool) bool {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect8containsERK6QPointb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &proper)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:137
// index:2
// Public inline
// bool contains(int, int)
func (this *QRect) Contains_2(x int, y int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect8containsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:138
// index:3
// Public inline
// bool contains(int, int, _Bool)
func (this *QRect) Contains_3(x int, y int, proper bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect8containsEiib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &proper)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:139
// index:0
// Public inline
// QRect united(const class QRect &)
func (this *QRect) United(other *QRect) *QRect /*123*/ {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect6unitedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:140
// index:0
// Public inline
// QRect intersected(const class QRect &)
func (this *QRect) Intersected(other *QRect) *QRect /*123*/ {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:141
// index:0
// Public
// bool intersects(const class QRect &)
func (this *QRect) Intersects(r *QRect) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:143
// index:0
// Public inline
// QRect marginsAdded(const class QMargins &)
func (this *QRect) MarginsAdded(margins *QMargins) *QRect /*123*/ {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect12marginsAddedERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:144
// index:0
// Public inline
// QRect marginsRemoved(const class QMargins &)
func (this *QRect) MarginsRemoved(margins *QMargins) *QRect /*123*/ {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QRect14marginsRemovedERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
