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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qcursor.h:82
// index:0
// void QCursor()
func NewQCursor() *QCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QCursor{cthis}
}

// /usr/include/qt/QtGui/qcursor.h:83
// index:1
// void QCursor(Qt::CursorShape)
func NewQCursor_1(shape int) *QCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2EN2Qt11CursorShapeE", ffiqt.FFI_TYPE_VOID, cthis, &shape)
	gopp.ErrPrint(err, rv)
	return &QCursor{cthis}
}

// /usr/include/qt/QtGui/qcursor.h:84
// index:2
// void QCursor(const class QBitmap &, const class QBitmap &, int, int)
func NewQCursor_2(bitmap unsafe.Pointer, mask unsafe.Pointer, hotX int, hotY int) *QCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2ERK7QBitmapS2_ii", ffiqt.FFI_TYPE_VOID, cthis, bitmap, mask, &hotX, &hotY)
	gopp.ErrPrint(err, rv)
	return &QCursor{cthis}
}

// /usr/include/qt/QtGui/qcursor.h:85
// index:3
// void QCursor(const class QPixmap &, int, int)
func NewQCursor_3(pixmap unsafe.Pointer, hotX int, hotY int) *QCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorC2ERK7QPixmapii", ffiqt.FFI_TYPE_VOID, cthis, pixmap, &hotX, &hotY)
	gopp.ErrPrint(err, rv)
	return &QCursor{cthis}
}

// /usr/include/qt/QtGui/qcursor.h:87
// index:0
// void ~QCursor()
func DeleteQCursor(*QCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:95
// index:0
// inline
// void swap(class QCursor &)
func (this *QCursor) Swap(other unsafe.Pointer) {
	// 0: (, QCursor & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:99
// index:0
// Qt::CursorShape shape()
func (this *QCursor) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:100
// index:0
// void setShape(Qt::CursorShape)
func (this *QCursor) SetShape(newShape int) {
	// 0: (, Qt::CursorShape newShape), (&newShape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor8setShapeEN2Qt11CursorShapeE", ffiqt.FFI_TYPE_VOID, this.cthis, &newShape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:102
// index:0
// const QBitmap * bitmap()
func (this *QCursor) Bitmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor6bitmapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:103
// index:0
// const QBitmap * mask()
func (this *QCursor) Mask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor4maskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:104
// index:0
// QPixmap pixmap()
func (this *QCursor) Pixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor6pixmapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:105
// index:0
// QPoint hotSpot()
func (this *QCursor) HotSpot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QCursor7hotSpotEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:107
// index:0
// static
// QPoint pos()
func (this *QCursor) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor3posEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCursor_Pos() {
	// 0: (), ()
	var nilthis *QCursor
	nilthis.Pos()
}

// /usr/include/qt/QtGui/qcursor.h:108
// index:1
// static
// QPoint pos(const class QScreen *)
func (this *QCursor) Pos_1(screen unsafe.Pointer) {
	// 1: (const QScreen * screen), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor3posEPK7QScreen", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCursor_Pos_1(screen unsafe.Pointer) {
	// 1: (const QScreen * screen), (screen)
	var nilthis *QCursor
	nilthis.Pos_1(screen)
}

// /usr/include/qt/QtGui/qcursor.h:109
// index:0
// static
// void setPos(int, int)
func (this *QCursor) SetPos(x int, y int) {
	// 0: (int x, int y), (x, y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos(x int, y int) {
	// 0: (int x, int y), (x, y)
	var nilthis *QCursor
	nilthis.SetPos(x, y)
}

// /usr/include/qt/QtGui/qcursor.h:110
// index:1
// static
// void setPos(class QScreen *, int, int)
func (this *QCursor) SetPos_1(screen unsafe.Pointer, x int, y int) {
	// 1: (QScreen * screen, int x, int y), (screen, x, y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_1(screen unsafe.Pointer, x int, y int) {
	// 1: (QScreen * screen, int x, int y), (screen, x, y)
	var nilthis *QCursor
	nilthis.SetPos_1(screen, x, y)
}

// /usr/include/qt/QtGui/qcursor.h:111
// index:2
// static inline
// void setPos(const class QPoint &)
func (this *QCursor) SetPos_2(p unsafe.Pointer) {
	// 2: (const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosERK6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_2(p unsafe.Pointer) {
	// 2: (const QPoint & p), (p)
	var nilthis *QCursor
	nilthis.SetPos_2(p)
}

// /usr/include/qt/QtGui/qcursor.h:112
// index:3
// static inline
// void setPos(class QScreen *, const class QPoint &)
func (this *QCursor) SetPos_3(screen unsafe.Pointer, p unsafe.Pointer) {
	// 3: (QScreen * screen, const QPoint & p), (screen, p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenRK6QPoint", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCursor_SetPos_3(screen unsafe.Pointer, p unsafe.Pointer) {
	// 3: (QScreen * screen, const QPoint & p), (screen, p)
	var nilthis *QCursor
	nilthis.SetPos_3(screen, p)
}

//  body block end
