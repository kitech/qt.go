package qtgui

// /usr/include/qt/QtGui/qcursor.h
// #include <qcursor.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QCursor struct {
	*qtrt.CObject
}
type QCursor_ITF interface {
	QCursor_PTR() *QCursor
}

func (ptr *QCursor) QCursor_PTR() *QCursor { return ptr }

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
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCursor(Qt::CursorShape)
func NewQCursor_1(shape int) *QCursor {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursorC2EN2Qt11CursorShapeE", qtrt.FFI_TYPE_POINTER, shape)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCursor(const QBitmap &, const QBitmap &, int, int)
func NewQCursor_2(bitmap QBitmap_ITF, mask QBitmap_ITF, hotX int, hotY int) *QCursor {
	var convArg0 unsafe.Pointer
	if bitmap != nil && bitmap.QBitmap_PTR() != nil {
		convArg0 = bitmap.QBitmap_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if mask != nil && mask.QBitmap_PTR() != nil {
		convArg1 = mask.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursorC2ERK7QBitmapS2_ii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, hotX, hotY)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:85
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCursor(const QPixmap &, int, int)
func NewQCursor_3(pixmap QPixmap_ITF, hotX int, hotY int) *QCursor {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursorC2ERK7QPixmapii", qtrt.FFI_TYPE_POINTER, convArg0, hotX, hotY)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCursor)
	return gothis
}

// /usr/include/qt/QtGui/qcursor.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCursor()
func DeleteQCursor(this *QCursor) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qcursor.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor & operator=(const QCursor &)
func (this *QCursor) Operator_equal(cursor QCursor_ITF) *QCursor {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QCursor_PTR() != nil {
		convArg0 = cursor.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QCursor & operator=(QCursor &&)
func (this *QCursor) Operator_equal_1(other unsafe.Pointer /*333*/) *QCursor {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursoraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCursor &)
func (this *QCursor) Swap(other QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCursor_PTR() != nil {
		convArg0 = other.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorShape shape()
func (this *QCursor) Shape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCursor5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qcursor.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShape(Qt::CursorShape)
func (this *QCursor) SetShape(newShape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor8setShapeEN2Qt11CursorShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newShape)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBitmap * bitmap()
func (this *QCursor) Bitmap() *QBitmap /*777 const QBitmap **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCursor6bitmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qcursor.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] const QBitmap * mask()
func (this *QCursor) Mask() *QBitmap /*777 const QBitmap **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCursor4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qcursor.h:104
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap()
func (this *QCursor) Pixmap() *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCursor6pixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint hotSpot()
func (this *QCursor) HotSpot() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QCursor7hotSpotEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qcursor.h:107
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QCursor) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor3posEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
func (this *QCursor) Pos_1(screen QScreen_ITF /*777 const QScreen **/) *qtcore.QPoint /*123*/ {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor3posEPK7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}
func QCursor_Pos_1(screen QScreen_ITF /*777 const QScreen **/) *qtcore.QPoint /*123*/ {
	var nilthis *QCursor
	rv := nilthis.Pos_1(screen)
	return rv
}

// /usr/include/qt/QtGui/qcursor.h:109
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPos(int, int)
func (this *QCursor) SetPos(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor6setPosEii", qtrt.FFI_TYPE_POINTER, x, y)
	qtrt.ErrPrint(err, rv)
}
func QCursor_SetPos(x int, y int) {
	var nilthis *QCursor
	nilthis.SetPos(x, y)
}

// /usr/include/qt/QtGui/qcursor.h:110
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void setPos(QScreen *, int, int)
func (this *QCursor) SetPos_1(screen QScreen_ITF /*777 QScreen **/, x int, y int) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenii", qtrt.FFI_TYPE_POINTER, convArg0, x, y)
	qtrt.ErrPrint(err, rv)
}
func QCursor_SetPos_1(screen QScreen_ITF /*777 QScreen **/, x int, y int) {
	var nilthis *QCursor
	nilthis.SetPos_1(screen, x, y)
}

// /usr/include/qt/QtGui/qcursor.h:111
// index:2
// Public static inline Visibility=Default Availability=Available
// [-2] void setPos(const QPoint &)
func (this *QCursor) SetPos_2(p qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor6setPosERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCursor_SetPos_2(p qtcore.QPoint_ITF) {
	var nilthis *QCursor
	nilthis.SetPos_2(p)
}

// /usr/include/qt/QtGui/qcursor.h:112
// index:3
// Public static inline Visibility=Default Availability=Available
// [-2] void setPos(QScreen *, const QPoint &)
func (this *QCursor) SetPos_3(screen QScreen_ITF /*777 QScreen **/, p qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg1 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QCursor6setPosEP7QScreenRK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QCursor_SetPos_3(screen QScreen_ITF /*777 QScreen **/, p qtcore.QPoint_ITF) {
	var nilthis *QCursor
	nilthis.SetPos_3(screen, p)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
