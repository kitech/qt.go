//  header block begin
// /usr/include/qt/QtGui/qcursor.h
// #include <qcursor.h>
// #include <QtGui>
package qtgui

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
type QCursor struct {
	*qtrt.CObject
}

func (this *QCursor) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQCursorFromPointer(cthis unsafe.Pointer) *QCursor {
	return &QCursor{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qcursor.h:82
// index:0
// Public
// void QCursor()
func NewQCursor() *QCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:83
// index:1
// Public
// void QCursor(Qt::CursorShape)
func NewQCursor_1(shape int) *QCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2EN2Qt11CursorShapeE", ffiqt.FFI_TYPE_VOID, cthis, &shape)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:84
// index:2
// Public
// void QCursor(const class QBitmap &, const class QBitmap &, int, int)
func NewQCursor_2(bitmap *QBitmap, mask *QBitmap, hotX int, hotY int) *QCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = bitmap.GetCthis()
	var convArg1 = mask.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2ERK7QBitmapS2_ii", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, &hotX, &hotY)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:85
// index:3
// Public
// void QCursor(const class QPixmap &, int, int)
func NewQCursor_3(pixmap *QPixmap, hotX int, hotY int) *QCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2ERK7QPixmapii", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &hotX, &hotY)
	gopp.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:87
// index:0
// Public
// void ~QCursor()
func DeleteQCursor(*QCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:95
// index:0
// Public inline
// void swap(class QCursor &)
func (this *QCursor) Swap(other *QCursor) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:99
// index:0
// Public
// Qt::CursorShape shape()
func (this *QCursor) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:100
// index:0
// Public
// void setShape(Qt::CursorShape)
func (this *QCursor) SetShape(newShape int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor8setShapeEN2Qt11CursorShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &newShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:102
// index:0
// Public
// const QBitmap * bitmap()
func (this *QCursor) Bitmap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor6bitmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:103
// index:0
// Public
// const QBitmap * mask()
func (this *QCursor) Mask() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor4maskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:104
// index:0
// Public
// QPixmap pixmap()
func (this *QCursor) Pixmap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor6pixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:105
// index:0
// Public
// QPoint hotSpot()
func (this *QCursor) HotSpot() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor7hotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:107
// index:0
// Public static
// QPoint pos()
func (this *QCursor) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor3posEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCursor_Pos() {
	var nilthis *QCursor
	nilthis.Pos()
}

// /usr/include/qt/QtGui/qcursor.h:108
// index:1
// Public static
// QPoint pos(const class QScreen *)
func (this *QCursor) Pos_1(screen unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor3posEPK7QScreen", ffiqt.FFI_TYPE_POINTER, screen)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCursor_Pos_1(screen unsafe.Pointer) {
	var nilthis *QCursor
	nilthis.Pos_1(screen)
}

// /usr/include/qt/QtGui/qcursor.h:109
// index:0
// Public static
// void setPos(int, int)
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
// Public static
// void setPos(class QScreen *, int, int)
func (this *QCursor) SetPos_1(screen unsafe.Pointer, x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenii", ffiqt.FFI_TYPE_POINTER, screen, x, y)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_1(screen unsafe.Pointer, x int, y int) {
	var nilthis *QCursor
	nilthis.SetPos_1(screen, x, y)
}

// /usr/include/qt/QtGui/qcursor.h:111
// index:2
// Public static inline
// void setPos(const class QPoint &)
func (this *QCursor) SetPos_2(p *qtcore.QPoint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, p)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_2(p *qtcore.QPoint) {
	var nilthis *QCursor
	nilthis.SetPos_2(p)
}

// /usr/include/qt/QtGui/qcursor.h:112
// index:3
// Public static inline
// void setPos(class QScreen *, const class QPoint &)
func (this *QCursor) SetPos_3(screen unsafe.Pointer, p *qtcore.QPoint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenRK6QPoint", ffiqt.FFI_TYPE_POINTER, screen, p)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_3(screen unsafe.Pointer, p *qtcore.QPoint) {
	var nilthis *QCursor
	nilthis.SetPos_3(screen, p)
}

//  body block end
