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
import "gopp"
import "qt.go/qtrt"

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
func (this *QRect) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRectFromPointer(cthis unsafe.Pointer) *QRect {
	return &QRect{&qtrt.CObject{cthis}}
}
func (*QRect) NewFromPointer(cthis unsafe.Pointer) *QRect {
	return NewQRectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrect.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRect()
func NewQRect() *QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QRect(const QPoint &, const QPoint &)
func NewQRect_1(topleft *QPoint, bottomright *QPoint) *QRect {
	var convArg0 = topleft.GetCthis()
	var convArg1 = bottomright.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2ERK6QPointS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:62
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QRect(const QPoint &, const QSize &)
func NewQRect_2(topleft *QPoint, size *QSize) *QRect {
	var convArg0 = topleft.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2ERK6QPointRK5QSize", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:63
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QRect(int, int, int, int)
func NewQRect_3(left int, top int, width int, height int) *QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectC2Eiiii", qtrt.FFI_TYPE_POINTER, left, top, width, height)
	gopp.ErrPrint(err, rv)
	gothis := NewQRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRect)
	return gothis
}

// /usr/include/qt/QtCore/qrect.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QRect) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QRect) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRect) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int left()
func (this *QRect) Left() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect4leftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top()
func (this *QRect) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int right()
func (this *QRect) Right() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect5rightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom()
func (this *QRect) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:73
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect normalized()
func (this *QRect) Normalized() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QRect) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QRect) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeft(int)
func (this *QRect) SetLeft(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7setLeftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTop(int)
func (this *QRect) SetTop(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6setTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRight(int)
func (this *QRect) SetRight(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect8setRightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottom(int)
func (this *QRect) SetBottom(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9setBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setX(int)
func (this *QRect) SetX(x int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect4setXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setY(int)
func (this *QRect) SetY(y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect4setYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopLeft(const QPoint &)
func (this *QRect) SetTopLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect10setTopLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomRight(const QPoint &)
func (this *QRect) SetBottomRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect14setBottomRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopRight(const QPoint &)
func (this *QRect) SetTopRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect11setTopRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomLeft(const QPoint &)
func (this *QRect) SetBottomLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect13setBottomLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint topLeft()
func (this *QRect) TopLeft() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7topLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint bottomRight()
func (this *QRect) BottomRight() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect11bottomRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint topRight()
func (this *QRect) TopRight() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8topRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint bottomLeft()
func (this *QRect) BottomLeft() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10bottomLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint center()
func (this *QRect) Center() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveLeft(int)
func (this *QRect) MoveLeft(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect8moveLeftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTop(int)
func (this *QRect) MoveTop(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7moveTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveRight(int)
func (this *QRect) MoveRight(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9moveRightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottom(int)
func (this *QRect) MoveBottom(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect10moveBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopLeft(const QPoint &)
func (this *QRect) MoveTopLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect11moveTopLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottomRight(const QPoint &)
func (this *QRect) MoveBottomRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect15moveBottomRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTopRight(const QPoint &)
func (this *QRect) MoveTopRight(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect12moveTopRightERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBottomLeft(const QPoint &)
func (this *QRect) MoveBottomLeft(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect14moveBottomLeftERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveCenter(const QPoint &)
func (this *QRect) MoveCenter(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect10moveCenterERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(int, int)
func (this *QRect) Translate(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9translateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:106
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)
func (this *QRect) Translate_1(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect translated(int, int)
func (this *QRect) Translated(dx int, dy int) *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10translatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:108
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QRect translated(const QPoint &)
func (this *QRect) Translated_1(p *QPoint) *QRect /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10translatedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect transposed()
func (this *QRect) Transposed() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(int, int)
func (this *QRect) MoveTo(x int, t int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6moveToEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:112
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void moveTo(const QPoint &)
func (this *QRect) MoveTo_1(p *QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6moveToERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(int, int, int, int)
func (this *QRect) SetRect(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7setRectEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getRect(int *, int *, int *, int *)
func (this *QRect) GetRect(x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, w unsafe.Pointer /*666*/, h unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect7getRectEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCoords(int, int, int, int)
func (this *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9setCoordsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void getCoords(int *, int *, int *, int *)
func (this *QRect) GetCoords(x1 unsafe.Pointer /*666*/, y1 unsafe.Pointer /*666*/, x2 unsafe.Pointer /*666*/, y2 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect9getCoordsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &x1, &y1, &x2, &y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void adjust(int, int, int, int)
func (this *QRect) Adjust(x1 int, y1 int, x2 int, y2 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect6adjustEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect adjusted(int, int, int, int)
func (this *QRect) Adjusted(x1 int, y1 int, x2 int, y2 int) *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8adjustedEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x1, y1, x2, y2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize size()
func (this *QRect) Size() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width()
func (this *QRect) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height()
func (this *QRect) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrect.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(int)
func (this *QRect) SetWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(int)
func (this *QRect) SetHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect9setHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSize(const QSize &)
func (this *QRect) SetSize(s *QSize) {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRect7setSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrect.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRect &, _Bool)
func (this *QRect) Contains(r *QRect, proper bool) bool {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsERKS_b", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, proper)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:136
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(const QPoint &, _Bool)
func (this *QRect) Contains_1(p *QPoint, proper bool) bool {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsERK6QPointb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, proper)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:137
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(int, int)
func (this *QRect) Contains_2(x int, y int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:138
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(int, int, _Bool)
func (this *QRect) Contains_3(x int, y int, proper bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect8containsEiib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, proper)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect united(const QRect &)
func (this *QRect) United(other *QRect) *QRect /*123*/ {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect intersected(const QRect &)
func (this *QRect) Intersected(other *QRect) *QRect /*123*/ {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QRect &)
func (this *QRect) Intersects(r *QRect) bool {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrect.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect marginsAdded(const QMargins &)
func (this *QRect) MarginsAdded(margins *QMargins) *QRect /*123*/ {
	var convArg0 = margins.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect12marginsAddedERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qrect.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect marginsRemoved(const QMargins &)
func (this *QRect) MarginsRemoved(margins *QMargins) *QRect /*123*/ {
	var convArg0 = margins.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QRect14marginsRemovedERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

func DeleteQRect(this *QRect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QRectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
