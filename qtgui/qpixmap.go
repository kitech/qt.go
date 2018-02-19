package qtgui

// /usr/include/qt/QtGui/qpixmap.h
// #include <qpixmap.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 108
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

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPixmap) InheritMetric(f func(arg0 int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// QPixmap fromImageInPlace(class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) InheritFromImageInPlace(f func(image *QImage, flags int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "fromImageInPlace", f)
}

type QPixmap struct {
	*QPaintDevice
}
type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap { return ptr }

func (this *QPixmap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QPixmap) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQPixmapFromPointer(cthis unsafe.Pointer) *QPixmap {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPixmap{bcthis0}
}
func (*QPixmap) NewFromPointer(cthis unsafe.Pointer) *QPixmap {
	return NewQPixmapFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpixmap.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPixmap()
func NewQPixmap() *QPixmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(int, int)
func NewQPixmap_1(w int, h int) *QPixmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2Eii", qtrt.FFI_TYPE_POINTER, w, h)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:67
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QSize &)
func NewQPixmap_2(arg0 qtcore.QSize_ITF) *QPixmap {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QString &, const char *, Qt::ImageConversionFlags)
func NewQPixmap_3(fileName string, format string, flags int) *QPixmap {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QString &, const char *, Qt::ImageConversionFlags)
func NewQPixmap_3_(fileName string) *QPixmap {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const QString &, const char *, Qt::ImageConversionFlags)
func NewQPixmap_3_1(fileName string, format string) *QPixmap {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:70
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QPixmap(const char *const *)
func NewQPixmap_4(xpm []string) *QPixmap {
	var convArg0 = qtrt.StringSliceToCCharPP(xpm)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapC2EPKPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPixmap()
func DeleteQPixmap(this *QPixmap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpixmap.h:75
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap & operator=(const QPixmap &)
func (this *QPixmap) Operator_equal(arg0 QPixmap_ITF) *QPixmap {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:77
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap & operator=(QPixmap &&)
func (this *QPixmap) Operator_equal_1(other unsafe.Pointer /*333*/) *QPixmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmapaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPixmap &)
func (this *QPixmap) Swap(other QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPixmap_PTR() != nil {
		convArg0 = other.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QPixmap) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const
func (this *QPixmap) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int width() const
func (this *QPixmap) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int height() const
func (this *QPixmap) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const
func (this *QPixmap) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:91
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rect() const
func (this *QPixmap) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int depth() const
func (this *QPixmap) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [4] int defaultDepth()
func (this *QPixmap) DefaultDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12defaultDepthEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPixmap_DefaultDepth() int {
	var nilthis *QPixmap
	rv := nilthis.DefaultDepth()
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fill(const QColor &)
func (this *QPixmap) Fill(fillColor QColor_ITF) {
	var convArg0 unsafe.Pointer
	if fillColor != nil && fillColor.QColor_PTR() != nil {
		convArg0 = fillColor.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fill(const QColor &)
func (this *QPixmap) Fill__() {
	// arg: 0, const QColor &=LValueReference, QColor=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:97
// index:1
// Public Visibility=Default Availability=Available
// [-2] void fill(const QPaintDevice *, const QPoint &)
func (this *QPixmap) Fill_1(device QPaintDevice_ITF /*777 const QPaintDevice **/, ofs qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if ofs != nil && ofs.QPoint_PTR() != nil {
		convArg1 = ofs.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:98
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void fill(const QPaintDevice *, int, int)
func (this *QPixmap) Fill_2(device QPaintDevice_ITF /*777 const QPaintDevice **/, xofs int, yofs int) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xofs, yofs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:100
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap mask() const
func (this *QPixmap) Mask() *QBitmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QBitmap &)
func (this *QPixmap) SetMask(arg0 QBitmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QBitmap_PTR() != nil {
		convArg0 = arg0.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap7setMaskERK7QBitmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const
func (this *QPixmap) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevicePixelRatio(qreal)
func (this *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap19setDevicePixelRatioEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scaleFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlpha() const
func (this *QPixmap) HasAlpha() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap8hasAlphaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlphaChannel() const
func (this *QPixmap) HasAlphaChannel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap15hasAlphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:110
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createHeuristicMask(_Bool) const
func (this *QPixmap) CreateHeuristicMask(clipTight bool) *QBitmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:110
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createHeuristicMask(_Bool) const
func (this *QPixmap) CreateHeuristicMask__() *QBitmap /*123*/ {
	// arg: 0, bool=Bool, =Invalid,
	clipTight := true
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:112
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createMaskFromColor(const QColor &, Qt::MaskMode) const
func (this *QPixmap) CreateMaskFromColor(maskColor QColor_ITF, mode int) *QBitmap /*123*/ {
	var convArg0 unsafe.Pointer
	if maskColor != nil && maskColor.QColor_PTR() != nil {
		convArg0 = maskColor.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createMaskFromColorERK6QColorN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:112
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap createMaskFromColor(const QColor &, Qt::MaskMode) const
func (this *QPixmap) CreateMaskFromColor__(maskColor QColor_ITF) *QBitmap /*123*/ {
	var convArg0 unsafe.Pointer
	if maskColor != nil && maskColor.QColor_PTR() != nil {
		convArg0 = maskColor.QColor_PTR().GetCthis()
	}
	// arg: 1, Qt::MaskMode=Elaborated, Qt::MaskMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap19createMaskFromColorERK6QColorN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow(arg0 uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_GrabWindow(arg0 uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWindow(arg0, x, y, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow__(arg0 uint64) *QPixmap /*123*/ {
	// arg: 1, int=Int, =Invalid,
	x := 0
	// arg: 2, int=Int, =Invalid,
	y := 0
	// arg: 3, int=Int, =Invalid,
	w := -1
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow__1(arg0 uint64, x int) *QPixmap /*123*/ {
	// arg: 2, int=Int, =Invalid,
	y := 0
	// arg: 3, int=Int, =Invalid,
	w := -1
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow__2(arg0 uint64, x int, y int) *QPixmap /*123*/ {
	// arg: 3, int=Int, =Invalid,
	w := -1
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow__3(arg0 uint64, x int, y int, w int) *QPixmap /*123*/ {
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:115
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, const QRect &)
func (this *QPixmap) GrabWidget(widget qtcore.QObject_ITF /*777 QObject **/, rect qtcore.QRect_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_GrabWidget(widget qtcore.QObject_ITF /*777 QObject **/, rect qtcore.QRect_ITF) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWidget(widget, rect)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int, w int, h int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_GrabWidget_1(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int, w int, h int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWidget_1(widget, x, y, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1_(widget qtcore.QObject_ITF /*777 QObject **/) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	x := 0
	// arg: 2, int=Int, =Invalid,
	y := 0
	// arg: 3, int=Int, =Invalid,
	w := -1
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1_1(widget qtcore.QObject_ITF /*777 QObject **/, x int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	y := 0
	// arg: 3, int=Int, =Invalid,
	w := -1
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1_2(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	w := -1
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap grabWidget(QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1_3(widget qtcore.QObject_ITF /*777 QObject **/, x int, y int, w int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QObject_PTR() != nil {
		convArg0 = widget.QObject_PTR().GetCthis()
	}
	// arg: 4, int=Int, =Invalid,
	h := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QPixmap) Scaled(w int, h int, aspectMode int, mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QPixmap) Scaled__(w int, h int) *QPixmap /*123*/ {
	// arg: 2, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectMode := 0
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QPixmap) Scaled__1(w int, h int, aspectMode int) *QPixmap /*123*/ {
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QPixmap) Scaled_1(s qtcore.QSize_ITF, aspectMode int, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QPixmap) Scaled_1_(s qtcore.QSize_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum,
	aspectMode := 0
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const
func (this *QPixmap) Scaled_1_1(s qtcore.QSize_ITF, aspectMode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:124
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToWidth(int, Qt::TransformationMode) const
func (this *QPixmap) ScaledToWidth(w int, mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:124
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToWidth(int, Qt::TransformationMode) const
func (this *QPixmap) ScaledToWidth__(w int) *QPixmap /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToHeight(int, Qt::TransformationMode) const
func (this *QPixmap) ScaledToHeight(h int, mode int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledToHeight(int, Qt::TransformationMode) const
func (this *QPixmap) ScaledToHeight__(h int) *QPixmap /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:126
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QMatrix &, Qt::TransformationMode) const
func (this *QPixmap) Transformed(arg0 QMatrix_ITF, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:126
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QMatrix &, Qt::TransformationMode) const
func (this *QPixmap) Transformed__(arg0 QMatrix_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:128
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QTransform &, Qt::TransformationMode) const
func (this *QPixmap) Transformed_1(arg0 QTransform_ITF, mode int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:128
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap transformed(const QTransform &, Qt::TransformationMode) const
func (this *QPixmap) Transformed_1_(arg0 QTransform_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [48] QMatrix trueMatrix(const QMatrix &, int, int)
func (this *QPixmap) TrueMatrix(m QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var convArg0 unsafe.Pointer
	if m != nil && m.QMatrix_PTR() != nil {
		convArg0 = m.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK7QMatrixii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}
func QPixmap_TrueMatrix(m QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.TrueMatrix(m, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:129
// index:1
// Public static Visibility=Default Availability=Available
// [88] QTransform trueMatrix(const QTransform &, int, int)
func (this *QPixmap) TrueMatrix_1(m QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if m != nil && m.QTransform_PTR() != nil {
		convArg0 = m.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK10QTransformii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QPixmap_TrueMatrix_1(m QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.TrueMatrix_1(m, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:131
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage toImage() const
func (this *QPixmap) ToImage() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap7toImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImage(const QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImage(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImage(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImage(const QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage__(image QImage_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:135
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap fromImage(QImage &&, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage_1(image unsafe.Pointer /*333*/, flags int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageEO6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, image, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImage_1(image unsafe.Pointer /*333*/, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImage_1(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:135
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QPixmap fromImage(QImage &&, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage_1_(image unsafe.Pointer /*333*/) *QPixmap /*123*/ {
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap9fromImageEO6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, image, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImageReader(QImageReader *, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageReader(imageReader QImageReader_ITF /*777 QImageReader **/, flags int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if imageReader != nil && imageReader.QImageReader_PTR() != nil {
		convArg0 = imageReader.QImageReader_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap15fromImageReaderEP12QImageReader6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImageReader(imageReader QImageReader_ITF /*777 QImageReader **/, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImageReader(imageReader, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [32] QPixmap fromImageReader(QImageReader *, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageReader__(imageReader QImageReader_ITF /*777 QImageReader **/) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if imageReader != nil && imageReader.QImageReader_PTR() != nil {
		convArg0 = imageReader.QImageReader_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap15fromImageReaderEP12QImageReader6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) Load(fileName string, format string, flags int) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) Load__(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) Load__1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData(buf unsafe.Pointer /*666*/, len_ uint, format string, flags int) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData__(buf unsafe.Pointer /*666*/, len_ uint) bool {
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData__1(buf unsafe.Pointer /*666*/, len_ uint, format string) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData_1(data qtcore.QByteArray_ITF, format string, flags int) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData_1_(data qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData_1_1(data qtcore.QByteArray_ITF, format string) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const
func (this *QPixmap) Save(fileName string, format string, quality int) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const
func (this *QPixmap) Save__(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid,
	quality := -1
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const
func (this *QPixmap) Save__1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid,
	quality := -1
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const
func (this *QPixmap) Save_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string, quality int) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const
func (this *QPixmap) Save_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid,
	quality := -1
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const
func (this *QPixmap) Save_1_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid,
	quality := -1
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool convertFromImage(const QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) ConvertFromImage(img QImage_ITF, flags int) bool {
	var convArg0 unsafe.Pointer
	if img != nil && img.QImage_PTR() != nil {
		convArg0 = img.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16convertFromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool convertFromImage(const QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) ConvertFromImage__(img QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if img != nil && img.QImage_PTR() != nil {
		convArg0 = img.QImage_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16convertFromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QPixmap copy(int, int, int, int) const
func (this *QPixmap) Copy(x int, y int, width int, height int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4copyEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, width, height)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:150
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap copy(const QRect &) const
func (this *QPixmap) Copy_1(rect qtcore.QRect_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:150
// index:1
// Public Visibility=Default Availability=Available
// [32] QPixmap copy(const QRect &) const
func (this *QPixmap) Copy_1_() *QPixmap /*123*/ {
	// arg: 0, const QRect &=LValueReference, QRect=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scroll(int, int, int, int, int, int, QRegion *)
func (this *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed QRegion_ITF /*777 QRegion **/) {
	var convArg6 unsafe.Pointer
	if exposed != nil && exposed.QRegion_PTR() != nil {
		convArg6 = exposed.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiiiiiP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, x, y, width, height, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scroll(int, int, int, int, int, int, QRegion *)
func (this *QPixmap) Scroll__(dx int, dy int, x int, y int, width int, height int) {
	// arg: 6, QRegion *=Pointer, QRegion=Record,
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiiiiiP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, x, y, width, height, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:153
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int, const QRect &, QRegion *)
func (this *QPixmap) Scroll_1(dx int, dy int, rect qtcore.QRect_ITF, exposed QRegion_ITF /*777 QRegion **/) {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if exposed != nil && exposed.QRegion_PTR() != nil {
		convArg3 = exposed.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiRK5QRectP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:153
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int, const QRect &, QRegion *)
func (this *QPixmap) Scroll_1_(dx int, dy int, rect qtcore.QRect_ITF) {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	// arg: 3, QRegion *=Pointer, QRegion=Record,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiRK5QRectP7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey() const
func (this *QPixmap) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qpixmap.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const
func (this *QPixmap) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QPixmap) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isQBitmap() const
func (this *QPixmap) IsQBitmap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap9isQBitmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:165
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const
func (this *QPixmap) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpixmap.h:167
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!() const
func (this *QPixmap) Operator_not() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmapntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:175
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric) const
func (this *QPixmap) Metric(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmap.h:176
// index:0
// Protected static Visibility=Default Availability=Available
// [32] QPixmap fromImageInPlace(QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageInPlace(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}
func QPixmap_FromImageInPlace(image QImage_ITF, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImageInPlace(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:176
// index:0
// Protected static Visibility=Default Availability=Available
// [32] QPixmap fromImageInPlace(QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageInPlace__(image QImage_ITF) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
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
