package qtgui

// /usr/include/qt/QtGui/qicon.h
// #include <qicon.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

/*

 */
type QIcon struct {
	*qtrt.CObject
}
type QIcon_ITF interface {
	QIcon_PTR() *QIcon
}

func (ptr *QIcon) QIcon_PTR() *QIcon { return ptr }

func (this *QIcon) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIcon) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQIconFromPointer(cthis unsafe.Pointer) *QIcon {
	return &QIcon{&qtrt.CObject{cthis}}
}
func (*QIcon) NewFromPointer(cthis unsafe.Pointer) *QIcon {
	return NewQIconFromPointer(cthis)
}

// /usr/include/qt/QtGui/qicon.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIcon()

/*
Constructs a null icon.
*/
func NewQIcon() *QIcon {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QIcon(const QPixmap &)

/*
Constructs a null icon.
*/
func NewQIcon_1(pixmap QPixmap_ITF) *QIcon {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2ERK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QIcon(const QString &)

/*
Constructs a null icon.
*/
func NewQIcon_2(fileName string) *QIcon {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:69
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QIcon(QIconEngine *)

/*
Constructs a null icon.
*/
func NewQIcon_3(engine QIconEngine_ITF /*777 QIconEngine **/) *QIcon {
	var convArg0 unsafe.Pointer
	if engine != nil && engine.QIconEngine_PTR() != nil {
		convArg0 = engine.QIconEngine_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2EP11QIconEngine", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QIcon()

/*

 */
func DeleteQIcon(this *QIcon) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qicon.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon & operator=(const QIcon &)

/*

 */
func (this *QIcon) Operator_equal(other QIcon_ITF) *QIcon {
	var convArg0 unsafe.Pointer
	if other != nil && other.QIcon_PTR() != nil {
		convArg0 = other.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QIcon & operator=(QIcon &&)

/*

 */
func (this *QIcon) Operator_equal_1(other unsafe.Pointer /*333*/) *QIcon {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QIcon &)

/*
Swaps icon other with this icon. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QIcon) Swap(other QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QIcon_PTR() != nil {
		convArg0 = other.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap(size qtcore.QSize_ITF, mode int, state int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap__(size qtcore.QSize_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	// arg: 1, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap__1(size qtcore.QSize_ITF, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_1(w int, h int, mode int, state int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_1_(w int, h int) *QPixmap /*123*/ {
	// arg: 2, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_1_1(w int, h int, mode int) *QPixmap /*123*/ {
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_2(extent int, mode int, state int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extent, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_2_(extent int) *QPixmap /*123*/ {
	// arg: 1, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extent, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_2_1(extent int, mode int) *QPixmap /*123*/ {
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extent, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:86
// index:3
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_3(window QWindow_ITF /*777 QWindow **/, size qtcore.QSize_ITF, mode int, state int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:86
// index:3
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_3_(window QWindow_ITF /*777 QWindow **/, size qtcore.QSize_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:86
// index:3
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns a pixmap with the requested size, mode, and state, generating one if necessary. The pixmap might be smaller than requested, but never larger.

Setting the Qt::AA_UseHighDpiPixmaps application attribute enables this function to return pixmaps that are larger than the requested size. Such images will have a devicePixelRatio larger than 1.

See also actualSize() and paint().
*/
func (this *QIcon) Pixmap_3_1(window QWindow_ITF /*777 QWindow **/, size qtcore.QSize_ITF, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns the actual size of the icon for the requested size, mode, and state. The result might be smaller than requested, but never larger. The returned size is in device-independent pixels (This is relevant for high-dpi pixmaps.)

See also pixmap() and paint().
*/
func (this *QIcon) ActualSize(size qtcore.QSize_ITF, mode int, state int) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns the actual size of the icon for the requested size, mode, and state. The result might be smaller than requested, but never larger. The returned size is in device-independent pixels (This is relevant for high-dpi pixmaps.)

See also pixmap() and paint().
*/
func (this *QIcon) ActualSize__(size qtcore.QSize_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	// arg: 1, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns the actual size of the icon for the requested size, mode, and state. The result might be smaller than requested, but never larger. The returned size is in device-independent pixels (This is relevant for high-dpi pixmaps.)

See also pixmap() and paint().
*/
func (this *QIcon) ActualSize__1(size qtcore.QSize_ITF, mode int) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:89
// index:1
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns the actual size of the icon for the requested size, mode, and state. The result might be smaller than requested, but never larger. The returned size is in device-independent pixels (This is relevant for high-dpi pixmaps.)

See also pixmap() and paint().
*/
func (this *QIcon) ActualSize_1(window QWindow_ITF /*777 QWindow **/, size qtcore.QSize_ITF, mode int, state int) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:89
// index:1
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns the actual size of the icon for the requested size, mode, and state. The result might be smaller than requested, but never larger. The returned size is in device-independent pixels (This is relevant for high-dpi pixmaps.)

See also pixmap() and paint().
*/
func (this *QIcon) ActualSize_1_(window QWindow_ITF /*777 QWindow **/, size qtcore.QSize_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:89
// index:1
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State) const

/*
Returns the actual size of the icon for the requested size, mode, and state. The result might be smaller than requested, but never larger. The returned size is in device-independent pixels (This is relevant for high-dpi pixmaps.)

See also pixmap() and paint().
*/
func (this *QIcon) ActualSize_1_1(window QWindow_ITF /*777 QWindow **/, size qtcore.QSize_ITF, mode int) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the name used to create the icon, if available.

Depending on the way the icon was created, it may have an associated name. This is the case for icons created with fromTheme() or icons using a QIconEngine which supports the QIconEngine::IconNameHook.

This function was introduced in  Qt 4.7.

See also fromTheme() and QIconEngine.
*/
func (this *QIcon) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qicon.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QRect &, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, alignment int, mode int, state int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainterRK5QRect6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QRect &, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint__(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	// arg: 2, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>
	alignment := 0
	// arg: 3, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 4, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainterRK5QRect6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QRect &, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint__1(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, alignment int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	// arg: 3, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 4, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainterRK5QRect6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QRect &, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint__2(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, alignment int, mode int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	// arg: 4, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainterRK5QRect6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:94
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void paint(QPainter *, int, int, int, int, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint_1(painter QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, alignment int, mode int, state int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainteriiii6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:94
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void paint(QPainter *, int, int, int, int, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint_1_(painter QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 5, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>
	alignment := 0
	// arg: 6, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 7, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainteriiii6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:94
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void paint(QPainter *, int, int, int, int, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint_1_1(painter QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, alignment int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 6, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 7, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainteriiii6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:94
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void paint(QPainter *, int, int, int, int, Qt::Alignment, enum QIcon::Mode, enum QIcon::State) const

/*
Uses the painter to paint the icon with specified alignment, required mode, and state into the rectangle rect.

See also actualSize() and pixmap().
*/
func (this *QIcon) Paint_1_2(painter QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, alignment int, mode int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 7, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainteriiii6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the icon is empty; otherwise returns false.

An icon is empty if it has neither a pixmap nor a filename.

Note: Even a non-null icon might not be able to create valid pixmaps, eg. if the file does not exist or cannot be read.
*/
func (this *QIcon) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QIcon) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QIcon) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey() const

/*
Returns a number that identifies the contents of this QIcon object. Distinct QIcon objects can have the same key if they refer to the same contents.

The cacheKey() will change when the icon is altered via addPixmap() or addFile().

Cache keys are mostly useful in conjunction with caching.

This function was introduced in  Qt 4.3.

See also QPixmap::cacheKey().
*/
func (this *QIcon) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPixmap(const QPixmap &, enum QIcon::Mode, enum QIcon::State)

/*
Adds pixmap to the icon, as a specialization for mode and state.

Custom icon engines are free to ignore additionally added pixmaps.

See also addFile().
*/
func (this *QIcon) AddPixmap(pixmap QPixmap_ITF, mode int, state int) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPixmap(const QPixmap &, enum QIcon::Mode, enum QIcon::State)

/*
Adds pixmap to the icon, as a specialization for mode and state.

Custom icon engines are free to ignore additionally added pixmaps.

See also addFile().
*/
func (this *QIcon) AddPixmap__(pixmap QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	// arg: 1, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPixmap(const QPixmap &, enum QIcon::Mode, enum QIcon::State)

/*
Adds pixmap to the icon, as a specialization for mode and state.

Custom icon engines are free to ignore additionally added pixmaps.

See also addFile().
*/
func (this *QIcon) AddPixmap__1(pixmap QPixmap_ITF, mode int) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	// arg: 2, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, enum QIcon::Mode, enum QIcon::State)

/*
Adds an image from the file with the given fileName to the icon, as a specialization for size, mode and state. The file will be loaded on demand. Note: custom icon engines are free to ignore additionally added pixmaps.

If fileName contains a relative path (e.g. the filename only) the relevant file must be found relative to the runtime working directory.

The file name can refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

Use the QImageReader::supportedImageFormats() and QImageWriter::supportedImageFormats() functions to retrieve a complete list of the supported file formats.

If a high resolution version of the image exists (identified by the suffix @2x on the base name), it is automatically loaded and added with the device pixel ratio set to a value of 2. This can be disabled by setting the environment variable QT_HIGHDPI_DISABLE_2X_IMAGE_LOADING (see QImageReader).

Note: When you add a non-empty filename to a QIcon, the icon becomes non-null, even if the file doesn't exist or points to a corrupt file.

See also addPixmap() and QPixmap::devicePixelRatio().
*/
func (this *QIcon) AddFile(fileName string, size qtcore.QSize_ITF, mode int, state int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, enum QIcon::Mode, enum QIcon::State)

/*
Adds an image from the file with the given fileName to the icon, as a specialization for size, mode and state. The file will be loaded on demand. Note: custom icon engines are free to ignore additionally added pixmaps.

If fileName contains a relative path (e.g. the filename only) the relevant file must be found relative to the runtime working directory.

The file name can refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

Use the QImageReader::supportedImageFormats() and QImageWriter::supportedImageFormats() functions to retrieve a complete list of the supported file formats.

If a high resolution version of the image exists (identified by the suffix @2x on the base name), it is automatically loaded and added with the device pixel ratio set to a value of 2. This can be disabled by setting the environment variable QT_HIGHDPI_DISABLE_2X_IMAGE_LOADING (see QImageReader).

Note: When you add a non-empty filename to a QIcon, the icon becomes non-null, even if the file doesn't exist or points to a corrupt file.

See also addPixmap() and QPixmap::devicePixelRatio().
*/
func (this *QIcon) AddFile__(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QSize &=LValueReference, QSize=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, enum QIcon::Mode, enum QIcon::State)

/*
Adds an image from the file with the given fileName to the icon, as a specialization for size, mode and state. The file will be loaded on demand. Note: custom icon engines are free to ignore additionally added pixmaps.

If fileName contains a relative path (e.g. the filename only) the relevant file must be found relative to the runtime working directory.

The file name can refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

Use the QImageReader::supportedImageFormats() and QImageWriter::supportedImageFormats() functions to retrieve a complete list of the supported file formats.

If a high resolution version of the image exists (identified by the suffix @2x on the base name), it is automatically loaded and added with the device pixel ratio set to a value of 2. This can be disabled by setting the environment variable QT_HIGHDPI_DISABLE_2X_IMAGE_LOADING (see QImageReader).

Note: When you add a non-empty filename to a QIcon, the icon becomes non-null, even if the file doesn't exist or points to a corrupt file.

See also addPixmap() and QPixmap::devicePixelRatio().
*/
func (this *QIcon) AddFile__1(fileName string, size qtcore.QSize_ITF) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 2, QIcon::Mode=Enum, QIcon::Mode=Enum,
	mode := 0
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, enum QIcon::Mode, enum QIcon::State)

/*
Adds an image from the file with the given fileName to the icon, as a specialization for size, mode and state. The file will be loaded on demand. Note: custom icon engines are free to ignore additionally added pixmaps.

If fileName contains a relative path (e.g. the filename only) the relevant file must be found relative to the runtime working directory.

The file name can refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

Use the QImageReader::supportedImageFormats() and QImageWriter::supportedImageFormats() functions to retrieve a complete list of the supported file formats.

If a high resolution version of the image exists (identified by the suffix @2x on the base name), it is automatically loaded and added with the device pixel ratio set to a value of 2. This can be disabled by setting the environment variable QT_HIGHDPI_DISABLE_2X_IMAGE_LOADING (see QImageReader).

Note: When you add a non-empty filename to a QIcon, the icon becomes non-null, even if the file doesn't exist or points to a corrupt file.

See also addPixmap() and QPixmap::devicePixelRatio().
*/
func (this *QIcon) AddFile__2(fileName string, size qtcore.QSize_ITF, mode int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	// arg: 3, QIcon::State=Enum, QIcon::State=Enum,
	state := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIsMask(_Bool)

/*
Indicate that this icon is a mask image(boolean isMask), and hence can potentially be modified based on where it's displayed.

This function was introduced in  Qt 5.6.

See also isMask().
*/
func (this *QIcon) SetIsMask(isMask bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9setIsMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), isMask)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMask() const

/*
Returns true if this icon has been marked as a mask image. Certain platforms render mask icons differently (for example, menu icons on macOS).

This function was introduced in  Qt 5.6.

See also setIsMask().
*/
func (this *QIcon) IsMask() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6isMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon fromTheme(const QString &)

/*
Returns the QIcon corresponding to name in the current icon theme.

The latest version of the freedesktop icon specification and naming specification can be obtained here:


http://standards.freedesktop.org/icon-theme-spec/icon-theme-spec-latest.html
http://standards.freedesktop.org/icon-naming-spec/icon-naming-spec-latest.html


To fetch an icon from the current icon theme:


      QIcon undoicon = QIcon::fromTheme("edit-undo");



Note: By default, only X11 will support themed icons. In order to use themed icons on Mac and Windows, you will have to bundle a compliant theme in one of your themeSearchPaths() and set the appropriate themeName().

Note: Qt will make use of GTK's icon-theme.cache if present to speed up the lookup. These caches can be generated using gtk-update-icon-cache: https://developer.gnome.org/gtk3/stable/gtk-update-icon-cache.html.

This function was introduced in  Qt 4.6.

See also themeName(), setThemeName(), and themeSearchPaths().
*/
func (this *QIcon) FromTheme(name string) *QIcon /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QIcon_FromTheme(name string) *QIcon /*123*/ {
	var nilthis *QIcon
	rv := nilthis.FromTheme(name)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:115
// index:1
// Public static Visibility=Default Availability=Available
// [8] QIcon fromTheme(const QString &, const QIcon &)

/*
Returns the QIcon corresponding to name in the current icon theme.

The latest version of the freedesktop icon specification and naming specification can be obtained here:


http://standards.freedesktop.org/icon-theme-spec/icon-theme-spec-latest.html
http://standards.freedesktop.org/icon-naming-spec/icon-naming-spec-latest.html


To fetch an icon from the current icon theme:


      QIcon undoicon = QIcon::fromTheme("edit-undo");



Note: By default, only X11 will support themed icons. In order to use themed icons on Mac and Windows, you will have to bundle a compliant theme in one of your themeSearchPaths() and set the appropriate themeName().

Note: Qt will make use of GTK's icon-theme.cache if present to speed up the lookup. These caches can be generated using gtk-update-icon-cache: https://developer.gnome.org/gtk3/stable/gtk-update-icon-cache.html.

This function was introduced in  Qt 4.6.

See also themeName(), setThemeName(), and themeSearchPaths().
*/
func (this *QIcon) FromTheme_1(name string, fallback QIcon_ITF) *QIcon /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if fallback != nil && fallback.QIcon_PTR() != nil {
		convArg1 = fallback.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QStringRKS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QIcon_FromTheme_1(name string, fallback QIcon_ITF) *QIcon /*123*/ {
	var nilthis *QIcon
	rv := nilthis.FromTheme_1(name, fallback)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:116
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasThemeIcon(const QString &)

/*
Returns true if there is an icon available for name in the current icon theme, otherwise returns false.

This function was introduced in  Qt 4.6.

See also themeSearchPaths(), fromTheme(), and setThemeName().
*/
func (this *QIcon) HasThemeIcon(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon12hasThemeIconERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QIcon_HasThemeIcon(name string) bool {
	var nilthis *QIcon
	rv := nilthis.HasThemeIcon(name)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList themeSearchPaths()

/*
Returns the search paths for icon themes.

The default value will depend on the platform:

On X11, the search path will use the XDG_DATA_DIRS environment variable if available.

By default all platforms will have the resource directory :\icons as a fallback. You can use "rcc -project" to generate a resource file from your icon theme.

This function was introduced in  Qt 4.6.

See also setThemeSearchPaths(), fromTheme(), and setThemeName().
*/
func (this *QIcon) ThemeSearchPaths() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon16themeSearchPathsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QIcon_ThemeSearchPaths() *qtcore.QStringList /*123*/ {
	var nilthis *QIcon
	rv := nilthis.ThemeSearchPaths()
	return rv
}

// /usr/include/qt/QtGui/qicon.h:119
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setThemeSearchPaths(const QStringList &)

/*
Sets the search paths for icon themes to paths.

This function was introduced in  Qt 4.6.

See also themeSearchPaths(), fromTheme(), and setThemeName().
*/
func (this *QIcon) SetThemeSearchPaths(searchpath qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if searchpath != nil && searchpath.QStringList_PTR() != nil {
		convArg0 = searchpath.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon19setThemeSearchPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QIcon_SetThemeSearchPaths(searchpath qtcore.QStringList_ITF) {
	var nilthis *QIcon
	nilthis.SetThemeSearchPaths(searchpath)
}

// /usr/include/qt/QtGui/qicon.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString themeName()

/*
Returns the name of the current icon theme.

On X11, the current icon theme depends on your desktop settings. On other platforms it is not set by default.

This function was introduced in  Qt 4.6.

See also setThemeName(), themeSearchPaths(), fromTheme(), and hasThemeIcon().
*/
func (this *QIcon) ThemeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9themeNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QIcon_ThemeName() string {
	var nilthis *QIcon
	rv := nilthis.ThemeName()
	return rv
}

// /usr/include/qt/QtGui/qicon.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setThemeName(const QString &)

/*
Sets the current icon theme to name.

The name should correspond to a directory name in the themeSearchPath() containing an index.theme file describing it's contents.

This function was introduced in  Qt 4.6.

See also themeSearchPaths() and themeName().
*/
func (this *QIcon) SetThemeName(path string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon12setThemeNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QIcon_SetThemeName(path string) {
	var nilthis *QIcon
	nilthis.SetThemeName(path)
}

/*
This enum type describes the mode for which a pixmap is intended to be used. The currently defined modes are:


*/
type QIcon__Mode = int

// Display the pixmap when the user is not interacting with the icon, but the functionality represented by the icon is available.
const QIcon__Normal QIcon__Mode = 0

// Display the pixmap when the functionality represented by the icon is not available.
const QIcon__Disabled QIcon__Mode = 1

// Display the pixmap when the functionality represented by the icon is available and the user is interacting with the icon, for example, moving the mouse over it or clicking it.
const QIcon__Active QIcon__Mode = 2

// Display the pixmap when the item represented by the icon is selected.
const QIcon__Selected QIcon__Mode = 3

/*
This enum describes the state for which a pixmap is intended to be used. The state can be:


*/
type QIcon__State = int

// Display the pixmap when the widget is in an "on" state
const QIcon__On QIcon__State = 0

// Display the pixmap when the widget is in an "off" state
const QIcon__Off QIcon__State = 1

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
