//  header block begin
// /usr/include/qt/QtGui/qbitmap.h
// #include <qbitmap.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QBitmap struct {
	*QPixmap
}

func (this *QBitmap) GetCthis() unsafe.Pointer {
	return this.QPixmap.GetCthis()
}

// /usr/include/qt/QtGui/qbitmap.h:54
// index:0
// void QBitmap()
func NewQBitmap() *QBitmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmapC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(cthis)
	return gothis
}
func NewQBitmapFromPointer(cthis unsafe.Pointer) *QBitmap {
	bcthis0 := NewQPixmapFromPointer(cthis)
	return &QBitmap{bcthis0}
}

// /usr/include/qt/QtGui/qbitmap.h:55
// index:1
// void QBitmap(const class QPixmap &)
func NewQBitmap_1(arg0 unsafe.Pointer) *QBitmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmapC2ERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:56
// index:2
// void QBitmap(int, int)
func NewQBitmap_2(w int, h int) *QBitmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmapC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &w, &h)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:57
// index:3
// void QBitmap(const class QSize &)
func NewQBitmap_3(arg0 unsafe.Pointer) *QBitmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmapC2ERK5QSize", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:58
// index:4
// void QBitmap(const class QString &, const char *)
func NewQBitmap_4(fileName unsafe.Pointer, format unsafe.Pointer) *QBitmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmapC2ERK7QStringPKc", ffiqt.FFI_TYPE_VOID, cthis, fileName, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:65
// index:0
// virtual
// void ~QBitmap()
func DeleteQBitmap(*QBitmap) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmapD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:69
// index:0
// inline
// void swap(class QBitmap &)
func (this *QBitmap) Swap(other unsafe.Pointer) {
	// 0: (, other QBitmap &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmap4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:72
// index:0
// inline
// void clear()
func (this *QBitmap) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmap5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:74
// index:0
// static
// QBitmap fromImage(const class QImage &, Qt::ImageConversionFlags)
func (this *QBitmap) FromImage(image unsafe.Pointer, flags int) {
	// 0: (image const QImage &, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QBitmap_FromImage(image unsafe.Pointer, flags int) {
	// 0: (image const QImage &, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	var nilthis *QBitmap
	nilthis.FromImage(image, flags)
}

// /usr/include/qt/QtGui/qbitmap.h:75
// index:0
// static
// QBitmap fromData(const class QSize &, const uchar *, class QImage::Format)
func (this *QBitmap) FromData(size unsafe.Pointer, bits unsafe.Pointer, monoFormat int) {
	// 0: (size const QSize &, bits const uchar *, monoFormat QImage::Format), (size, bits, monoFormat)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBitmap8fromDataERK5QSizePKhN6QImage6FormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QBitmap_FromData(size unsafe.Pointer, bits unsafe.Pointer, monoFormat int) {
	// 0: (size const QSize &, bits const uchar *, monoFormat QImage::Format), (size, bits, monoFormat)
	var nilthis *QBitmap
	nilthis.FromData(size, bits, monoFormat)
}

// /usr/include/qt/QtGui/qbitmap.h:78
// index:0
// QBitmap transformed(const class QMatrix &)
func (this *QBitmap) Transformed(arg0 unsafe.Pointer) {
	// 0: (, const QMatrix & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBitmap11transformedERK7QMatrix", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:79
// index:1
// QBitmap transformed(const class QTransform &)
func (this *QBitmap) Transformed_1(matrix unsafe.Pointer) {
	// 1: (, matrix const QTransform &), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBitmap11transformedERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

//  body block end
