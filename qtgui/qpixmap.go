package qtgui

// /usr/include/qt/QtGui/qpixmap.h
// #include <qpixmap.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 103
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
type QPixmap struct {
	*QPaintDevice
}

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
// Public
// void QPixmap()
func NewQPixmap() *QPixmap {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:66
// index:1
// Public
// void QPixmap(int, int)
func NewQPixmap_1(w int, h int) *QPixmap {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2Eii", ffiqt.FFI_TYPE_VOID, cthis, w, h)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:67
// index:2
// Public
// void QPixmap(const class QSize &)
func NewQPixmap_2(arg0 *qtcore.QSize) *QPixmap {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2ERK5QSize", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:3
// Public
// void QPixmap(const class QString &, const char *, Qt::ImageConversionFlags)
func NewQPixmap_3(fileName *qtcore.QString, format string, flags int) *QPixmap {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = fileName.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:70
// index:4
// Public
// void QPixmap(const char *const *)
func NewQPixmap_4(xpm []interface{}) *QPixmap {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2EPKPKc", ffiqt.FFI_TYPE_VOID, cthis, xpm)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:73
// index:0
// Public virtual
// void ~QPixmap()
func DeleteQPixmap(*QPixmap) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:80
// index:0
// Public inline
// void swap(class QPixmap &)
func (this *QPixmap) Swap(other *QPixmap) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:85
// index:0
// Public
// bool isNull()
func (this *QPixmap) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:86
// index:0
// Public virtual
// int devType()
func (this *QPixmap) DevType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap7devTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpixmap.h:88
// index:0
// Public
// int width()
func (this *QPixmap) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpixmap.h:89
// index:0
// Public
// int height()
func (this *QPixmap) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpixmap.h:90
// index:0
// Public
// QSize size()
func (this *QPixmap) Size() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:91
// index:0
// Public
// QRect rect()
func (this *QPixmap) Rect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:92
// index:0
// Public
// int depth()
func (this *QPixmap) Depth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap5depthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpixmap.h:94
// index:0
// Public static
// int defaultDepth()
func (this *QPixmap) DefaultDepth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap12defaultDepthEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QPixmap_DefaultDepth() int {
	var nilthis *QPixmap
	rv := nilthis.DefaultDepth()
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:96
// index:0
// Public
// void fill(const class QColor &)
func (this *QPixmap) Fill(fillColor *QColor) {
	var convArg0 = fillColor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4fillERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:97
// index:1
// Public
// void fill(const class QPaintDevice *, const class QPoint &)
func (this *QPixmap) Fill_1(device *QPaintDevice /*777 const QPaintDevice **/, ofs *qtcore.QPoint) {
	var convArg0 = device.GetCthis()
	var convArg1 = ofs.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:98
// index:2
// Public inline
// void fill(const class QPaintDevice *, int, int)
func (this *QPixmap) Fill_2(device *QPaintDevice /*777 const QPaintDevice **/, xofs int, yofs int) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xofs, yofs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:100
// index:0
// Public
// QBitmap mask()
func (this *QPixmap) Mask() *QBitmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4maskEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:101
// index:0
// Public
// void setMask(const class QBitmap &)
func (this *QPixmap) SetMask(arg0 *QBitmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap7setMaskERK7QBitmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:103
// index:0
// Public
// qreal devicePixelRatio()
func (this *QPixmap) DevicePixelRatio() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap16devicePixelRatioEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpixmap.h:104
// index:0
// Public
// void setDevicePixelRatio(qreal)
func (this *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap19setDevicePixelRatioEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), scaleFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:106
// index:0
// Public
// bool hasAlpha()
func (this *QPixmap) HasAlpha() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap8hasAlphaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:107
// index:0
// Public
// bool hasAlphaChannel()
func (this *QPixmap) HasAlphaChannel() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap15hasAlphaChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:110
// index:0
// Public
// QBitmap createHeuristicMask(_Bool)
func (this *QPixmap) CreateHeuristicMask(clipTight bool) *QBitmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap19createHeuristicMaskEb", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), clipTight)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:112
// index:0
// Public
// QBitmap createMaskFromColor(const class QColor &, Qt::MaskMode)
func (this *QPixmap) CreateMaskFromColor(maskColor *QColor, mode int) *QBitmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = maskColor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap19createMaskFromColorERK6QColorN2Qt8MaskModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// Public static
// QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow(arg0 uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", ffiqt.FFI_TYPE_POINTER, arg0, x, y, w, h)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_GrabWindow(arg0 uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWindow(arg0, x, y, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:115
// index:0
// Public static
// QPixmap grabWidget(class QObject *, const class QRect &)
func (this *QPixmap) GrabWidget(widget *qtcore.QObject /*777 QObject **/, rect *qtcore.QRect) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect", ffiqt.FFI_TYPE_POINTER, widget, rect)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_GrabWidget(widget *qtcore.QObject /*777 QObject **/, rect *qtcore.QRect) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWidget(widget, rect)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// Public static inline
// QPixmap grabWidget(class QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1(widget *qtcore.QObject /*777 QObject **/, x int, y int, w int, h int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", ffiqt.FFI_TYPE_POINTER, widget, x, y, w, h)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_GrabWidget_1(widget *qtcore.QObject /*777 QObject **/, x int, y int, w int, h int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.GrabWidget_1(widget, x, y, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// Public inline
// QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode)
func (this *QPixmap) Scaled(w int, h int, aspectMode int, mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), w, h, aspectMode, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// Public
// QPixmap scaled(const class QSize &, Qt::AspectRatioMode, Qt::TransformationMode)
func (this *QPixmap) Scaled_1(s *qtcore.QSize, aspectMode int, mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, aspectMode, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:124
// index:0
// Public
// QPixmap scaledToWidth(int, Qt::TransformationMode)
func (this *QPixmap) ScaledToWidth(w int, mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap13scaledToWidthEiN2Qt18TransformationModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), w, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:125
// index:0
// Public
// QPixmap scaledToHeight(int, Qt::TransformationMode)
func (this *QPixmap) ScaledToHeight(h int, mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap14scaledToHeightEiN2Qt18TransformationModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), h, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:126
// index:0
// Public
// QPixmap transformed(const class QMatrix &, Qt::TransformationMode)
func (this *QPixmap) Transformed(arg0 *QMatrix, mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK7QMatrixN2Qt18TransformationModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:128
// index:1
// Public
// QPixmap transformed(const class QTransform &, Qt::TransformationMode)
func (this *QPixmap) Transformed_1(arg0 *QTransform, mode int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK10QTransformN2Qt18TransformationModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:127
// index:0
// Public static
// QMatrix trueMatrix(const class QMatrix &, int, int)
func (this *QPixmap) TrueMatrix(m *QMatrix, w int, h int) *QMatrix /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK7QMatrixii", ffiqt.FFI_TYPE_POINTER, m, w, h)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_TrueMatrix(m *QMatrix, w int, h int) *QMatrix /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.TrueMatrix(m, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:129
// index:1
// Public static
// QTransform trueMatrix(const class QTransform &, int, int)
func (this *QPixmap) TrueMatrix_1(m *QTransform, w int, h int) *QTransform /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK10QTransformii", ffiqt.FFI_TYPE_POINTER, m, w, h)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_TrueMatrix_1(m *QTransform, w int, h int) *QTransform /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.TrueMatrix_1(m, w, h)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:131
// index:0
// Public
// QImage toImage()
func (this *QPixmap) ToImage() *QImage /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap7toImageEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:132
// index:0
// Public static
// QPixmap fromImage(const class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage(image *QImage, flags int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, image, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_FromImage(image *QImage, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImage(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:135
// index:1
// Public static inline
// QPixmap fromImage(class QImage &&, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage_1(image unsafe.Pointer /*333*/, flags int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap9fromImageEO6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, image, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_FromImage_1(image unsafe.Pointer /*333*/, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImage_1(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:133
// index:0
// Public static
// QPixmap fromImageReader(class QImageReader *, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageReader(imageReader *QImageReader /*777 QImageReader **/, flags int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap15fromImageReaderEP12QImageReader6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, imageReader, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_FromImageReader(imageReader *QImageReader /*777 QImageReader **/, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImageReader(imageReader, flags)
	return rv
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// Public
// bool load(const class QString &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) Load(fileName *qtcore.QString, format string, flags int) bool {
	var convArg0 = fileName.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// Public
// bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData(buf unsafe.Pointer /*666*/, len uint, format string, flags int) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &buf, len, convArg2, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// Public inline
// bool loadFromData(const class QByteArray &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData_1(data *qtcore.QByteArray, format string, flags int) bool {
	var convArg0 = data.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// Public
// bool save(const class QString &, const char *, int)
func (this *QPixmap) Save(fileName *qtcore.QString, format string, quality int) bool {
	var convArg0 = fileName.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// Public
// bool save(class QIODevice *, const char *, int)
func (this *QPixmap) Save_1(device *qtcore.QIODevice /*777 QIODevice **/, format string, quality int) bool {
	var convArg0 = device.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:147
// index:0
// Public
// bool convertFromImage(const class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) ConvertFromImage(img *QImage, flags int) bool {
	var convArg0 = img.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap16convertFromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:149
// index:0
// Public inline
// QPixmap copy(int, int, int, int)
func (this *QPixmap) Copy(x int, y int, width int, height int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4copyEiiii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, width, height)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:150
// index:1
// Public
// QPixmap copy(const class QRect &)
func (this *QPixmap) Copy_1(rect *qtcore.QRect) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4copyERK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:152
// index:0
// Public inline
// void scroll(int, int, int, int, int, int, class QRegion *)
func (this *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed *QRegion /*777 QRegion **/) {
	var convArg6 = exposed.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiiiiiP7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, x, y, width, height, convArg6)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:153
// index:1
// Public
// void scroll(int, int, const class QRect &, class QRegion *)
func (this *QPixmap) Scroll_1(dx int, dy int, rect *qtcore.QRect, exposed *QRegion /*777 QRegion **/) {
	var convArg2 = rect.GetCthis()
	var convArg3 = exposed.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiRK5QRectP7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:158
// index:0
// Public
// qint64 cacheKey()
func (this *QPixmap) CacheKey() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap8cacheKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qpixmap.h:160
// index:0
// Public
// bool isDetached()
func (this *QPixmap) IsDetached() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:161
// index:0
// Public
// void detach()
func (this *QPixmap) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:163
// index:0
// Public
// bool isQBitmap()
func (this *QPixmap) IsQBitmap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap9isQBitmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpixmap.h:165
// index:0
// Public virtual
// QPaintEngine * paintEngine()
func (this *QPixmap) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpixmap.h:175
// index:0
// Protected virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPixmap) Metric(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpixmap.h:176
// index:0
// Protected static
// QPixmap fromImageInPlace(class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageInPlace(image *QImage, flags int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_POINTER, image, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPixmap_FromImageInPlace(image *QImage, flags int) *QPixmap /*123*/ {
	var nilthis *QPixmap
	rv := nilthis.FromImageInPlace(image, flags)
	return rv
}

//  body block end
