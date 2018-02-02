package qtgui

// /usr/include/qt/QtGui/qcursor.h
// #include <qcursor.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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

type QCursor struct {
	*qtrt.CObject
}

func (this *QCursor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCursor) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCursorFromPointer(cthis unsafe.Pointer) *QCursor {
	return &QCursor{&qtrt.CObject{cthis}}
}
func (*QCursor) NewFromPointer(cthis unsafe.Pointer) *QCursor {
	return NewQCursorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qcursor.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCursor()
func NewQCursor() *QCursor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCursor(Qt::CursorShape)
func NewQCursor_1(shape int) *QCursor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2EN2Qt11CursorShapeE", ffiqt.FFI_TYPE_POINTER, shape)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCursor(const QBitmap &, const QBitmap &, int, int)
func NewQCursor_2(bitmap *QBitmap, mask *QBitmap, hotX int, hotY int) *QCursor {
	var convArg0 = bitmap.GetCthis()
	var convArg1 = mask.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2ERK7QBitmapS2_ii", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, hotX, hotY)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:85
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCursor(const QPixmap &, int, int)
func NewQCursor_3(pixmap *QPixmap, hotX int, hotY int) *QCursor {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2ERK7QPixmapii", ffiqt.FFI_TYPE_POINTER, convArg0, hotX, hotY)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCursor()
func DeleteQCursor(this *QCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qcursor.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCursor &)
func (this *QCursor) Swap(other *QCursor) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorShape shape()
func (this *QCursor) Shape() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qcursor.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShape(Qt::CursorShape)
func (this *QCursor) SetShape(newShape int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor8setShapeEN2Qt11CursorShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBitmap * bitmap()
func (this *QCursor) Bitmap() *QBitmap /*777 const QBitmap **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor6bitmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBitmap * mask()
func (this *QCursor) Mask() *QBitmap /*777 const QBitmap **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor4maskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:104
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap()
func (this *QCursor) Pixmap() *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor6pixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint hotSpot()
func (this *QCursor) HotSpot() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor7hotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:107
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QCursor) Pos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor3posEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}
func QCursor_Pos() *qtcore.QPoint /*123*/ {
	var nilthis *QCursor
	rv := nilthis.Pos()
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:108
// index:1
// Public static Visibility=Default Availability=Available
// [8] QPoint pos(const QScreen *)
func (this *QCursor) Pos_1(screen *QScreen /*777 const QScreen **/) *qtcore.QPoint /*123*/ {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor3posEPK7QScreen", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}
func QCursor_Pos_1(screen *QScreen /*777 const QScreen **/) *qtcore.QPoint /*123*/ {
	var nilthis *QCursor
	rv := nilthis.Pos_1(screen)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:109
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPos(int, int)
func (this *QCursor) SetPos(x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEii", ffiqt.FFI_TYPE_POINTER, x, y)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos(x int, y int) {
	var nilthis *QCursor
	nilthis.SetPos(x, y)
}

// /usr/include/qt/QtGui/qcursor.h:110
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void setPos(QScreen *, int, int)
func (this *QCursor) SetPos_1(screen *QScreen /*777 QScreen **/, x int, y int) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenii", ffiqt.FFI_TYPE_POINTER, convArg0, x, y)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_1(screen *QScreen /*777 QScreen **/, x int, y int) {
	var nilthis *QCursor
	nilthis.SetPos_1(screen, x, y)
}

// /usr/include/qt/QtGui/qcursor.h:111
// index:2
// Public static inline Visibility=Default Availability=Available
// [-2] void setPos(const QPoint &)
func (this *QCursor) SetPos_2(p *qtcore.QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_2(p *qtcore.QPoint) {
	var nilthis *QCursor
	nilthis.SetPos_2(p)
}

// /usr/include/qt/QtGui/qcursor.h:112
// index:3
// Public static inline Visibility=Default Availability=Available
// [-2] void setPos(QScreen *, const QPoint &)
func (this *QCursor) SetPos_3(screen *QScreen /*777 QScreen **/, p *qtcore.QPoint) {
	var convArg0 = screen.GetCthis()
	var convArg1 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenRK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_3(screen *QScreen /*777 QScreen **/, p *qtcore.QPoint) {
	var nilthis *QCursor
	nilthis.SetPos_3(screen, p)
}

//  body block end
