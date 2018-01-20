//  header block begin
// /usr/include/qt/QtGui/qpixmap.h
// #include <qpixmap.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 105
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
	return this.QPaintDevice.GetCthis()
}

// /usr/include/qt/QtGui/qpixmap.h:64
// index:0
// void QPixmap()
func NewQPixmap() *QPixmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}
func NewQPixmapFromPointer(cthis unsafe.Pointer) *QPixmap {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPixmap{bcthis0}
}

// /usr/include/qt/QtGui/qpixmap.h:65
// index:1
// void QPixmap(class QPlatformPixmap *)
func NewQPixmap_1(data unsafe.Pointer) *QPixmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2EP15QPlatformPixmap", ffiqt.FFI_TYPE_VOID, cthis, data)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:66
// index:2
// void QPixmap(int, int)
func NewQPixmap_2(w int, h int) *QPixmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &w, &h)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:67
// index:3
// void QPixmap(const class QSize &)
func NewQPixmap_3(arg0 unsafe.Pointer) *QPixmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2ERK5QSize", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:68
// index:4
// void QPixmap(const class QString &, const char *, Qt::ImageConversionFlags)
func NewQPixmap_4(fileName unsafe.Pointer, format unsafe.Pointer, flags int) *QPixmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2ERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, cthis, fileName, format, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:70
// index:5
// void QPixmap(const char *const *)
func NewQPixmap_5(xpm []interface{}) *QPixmap {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapC2EPKPKc", ffiqt.FFI_TYPE_VOID, cthis, xpm)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixmapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixmap.h:73
// index:0
// virtual
// void ~QPixmap()
func DeleteQPixmap(*QPixmap) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmapD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:80
// index:0
// inline
// void swap(class QPixmap &)
func (this *QPixmap) Swap(other unsafe.Pointer) {
	// 0: (, other QPixmap &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:85
// index:0
// bool isNull()
func (this *QPixmap) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:86
// index:0
// virtual
// int devType()
func (this *QPixmap) DevType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap7devTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:88
// index:0
// int width()
func (this *QPixmap) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:89
// index:0
// int height()
func (this *QPixmap) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:90
// index:0
// QSize size()
func (this *QPixmap) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:91
// index:0
// QRect rect()
func (this *QPixmap) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4rectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:92
// index:0
// int depth()
func (this *QPixmap) Depth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap5depthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:94
// index:0
// static
// int defaultDepth()
func (this *QPixmap) DefaultDepth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap12defaultDepthEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_DefaultDepth() {
	// 0: (), ()
	var nilthis *QPixmap
	nilthis.DefaultDepth()
}

// /usr/include/qt/QtGui/qpixmap.h:96
// index:0
// void fill(const class QColor &)
func (this *QPixmap) Fill(fillColor unsafe.Pointer) {
	// 0: (, fillColor const QColor &), (fillColor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4fillERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fillColor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:97
// index:1
// void fill(const class QPaintDevice *, const class QPoint &)
func (this *QPixmap) Fill_1(device unsafe.Pointer, ofs unsafe.Pointer) {
	// 1: (, device const QPaintDevice *, ofs const QPoint &), (device, ofs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device, ofs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:98
// index:2
// inline
// void fill(const class QPaintDevice *, int, int)
func (this *QPixmap) Fill_2(device unsafe.Pointer, xofs int, yofs int) {
	// 2: (, device const QPaintDevice *, xofs int, yofs int), (device, &xofs, &yofs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4fillEPK12QPaintDeviceii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device, &xofs, &yofs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:100
// index:0
// QBitmap mask()
func (this *QPixmap) Mask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4maskEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:101
// index:0
// void setMask(const class QBitmap &)
func (this *QPixmap) SetMask(arg0 unsafe.Pointer) {
	// 0: (, const QBitmap & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap7setMaskERK7QBitmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:103
// index:0
// qreal devicePixelRatio()
func (this *QPixmap) DevicePixelRatio() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap16devicePixelRatioEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:104
// index:0
// void setDevicePixelRatio(qreal)
func (this *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	// 0: (, scaleFactor qreal), (&scaleFactor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap19setDevicePixelRatioEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &scaleFactor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:106
// index:0
// bool hasAlpha()
func (this *QPixmap) HasAlpha() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap8hasAlphaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:107
// index:0
// bool hasAlphaChannel()
func (this *QPixmap) HasAlphaChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap15hasAlphaChannelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:110
// index:0
// QBitmap createHeuristicMask(_Bool)
func (this *QPixmap) CreateHeuristicMask(clipTight bool) {
	// 0: (, clipTight bool), (&clipTight)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap19createHeuristicMaskEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &clipTight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:112
// index:0
// QBitmap createMaskFromColor(const class QColor &, Qt::MaskMode)
func (this *QPixmap) CreateMaskFromColor(maskColor unsafe.Pointer, mode int) {
	// 0: (, maskColor const QColor &, mode Qt::MaskMode), (maskColor, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap19createMaskFromColorERK6QColorN2Qt8MaskModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), maskColor, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:114
// index:0
// static
// QPixmap grabWindow(WId, int, int, int, int)
func (this *QPixmap) GrabWindow(arg0 uint64, x int, y int, w int, h int) {
	// 0: (WId arg0, x int, y int, w int, h int), (arg0, x, y, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10grabWindowEyiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_GrabWindow(arg0 uint64, x int, y int, w int, h int) {
	// 0: (WId arg0, x int, y int, w int, h int), (arg0, x, y, w, h)
	var nilthis *QPixmap
	nilthis.GrabWindow(arg0, x, y, w, h)
}

// /usr/include/qt/QtGui/qpixmap.h:115
// index:0
// static
// QPixmap grabWidget(class QObject *, const class QRect &)
func (this *QPixmap) GrabWidget(widget unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (widget QObject *, rect const QRect &), (widget, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_GrabWidget(widget unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (widget QObject *, rect const QRect &), (widget, rect)
	var nilthis *QPixmap
	nilthis.GrabWidget(widget, rect)
}

// /usr/include/qt/QtGui/qpixmap.h:116
// index:1
// static inline
// QPixmap grabWidget(class QObject *, int, int, int, int)
func (this *QPixmap) GrabWidget_1(widget unsafe.Pointer, x int, y int, w int, h int) {
	// 1: (widget QObject *, x int, y int, w int, h int), (widget, x, y, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10grabWidgetEP7QObjectiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_GrabWidget_1(widget unsafe.Pointer, x int, y int, w int, h int) {
	// 1: (widget QObject *, x int, y int, w int, h int), (widget, x, y, w, h)
	var nilthis *QPixmap
	nilthis.GrabWidget_1(widget, x, y, w, h)
}

// /usr/include/qt/QtGui/qpixmap.h:119
// index:0
// inline
// QPixmap scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode)
func (this *QPixmap) Scaled(w int, h int, aspectMode int, mode int) {
	// 0: (, w int, h int, aspectMode Qt::AspectRatioMode, mode Qt::TransformationMode), (&w, &h, &aspectMode, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h, &aspectMode, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:122
// index:1
// QPixmap scaled(const class QSize &, Qt::AspectRatioMode, Qt::TransformationMode)
func (this *QPixmap) Scaled_1(s unsafe.Pointer, aspectMode int, mode int) {
	// 1: (, s const QSize &, aspectMode Qt::AspectRatioMode, mode Qt::TransformationMode), (s, &aspectMode, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &aspectMode, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:124
// index:0
// QPixmap scaledToWidth(int, Qt::TransformationMode)
func (this *QPixmap) ScaledToWidth(w int, mode int) {
	// 0: (, w int, mode Qt::TransformationMode), (&w, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap13scaledToWidthEiN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:125
// index:0
// QPixmap scaledToHeight(int, Qt::TransformationMode)
func (this *QPixmap) ScaledToHeight(h int, mode int) {
	// 0: (, h int, mode Qt::TransformationMode), (&h, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap14scaledToHeightEiN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:126
// index:0
// QPixmap transformed(const class QMatrix &, Qt::TransformationMode)
func (this *QPixmap) Transformed(arg0 unsafe.Pointer, mode int) {
	// 0: (, const QMatrix & arg0, mode Qt::TransformationMode), (arg0, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK7QMatrixN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:128
// index:1
// QPixmap transformed(const class QTransform &, Qt::TransformationMode)
func (this *QPixmap) Transformed_1(arg0 unsafe.Pointer, mode int) {
	// 1: (, const QTransform & arg0, mode Qt::TransformationMode), (arg0, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap11transformedERK10QTransformN2Qt18TransformationModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:127
// index:0
// static
// QMatrix trueMatrix(const class QMatrix &, int, int)
func (this *QPixmap) TrueMatrix(m unsafe.Pointer, w int, h int) {
	// 0: (m const QMatrix &, w int, h int), (m, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK7QMatrixii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_TrueMatrix(m unsafe.Pointer, w int, h int) {
	// 0: (m const QMatrix &, w int, h int), (m, w, h)
	var nilthis *QPixmap
	nilthis.TrueMatrix(m, w, h)
}

// /usr/include/qt/QtGui/qpixmap.h:129
// index:1
// static
// QTransform trueMatrix(const class QTransform &, int, int)
func (this *QPixmap) TrueMatrix_1(m unsafe.Pointer, w int, h int) {
	// 1: (m const QTransform &, w int, h int), (m, w, h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap10trueMatrixERK10QTransformii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_TrueMatrix_1(m unsafe.Pointer, w int, h int) {
	// 1: (m const QTransform &, w int, h int), (m, w, h)
	var nilthis *QPixmap
	nilthis.TrueMatrix_1(m, w, h)
}

// /usr/include/qt/QtGui/qpixmap.h:131
// index:0
// QImage toImage()
func (this *QPixmap) ToImage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap7toImageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:132
// index:0
// static
// QPixmap fromImage(const class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage(image unsafe.Pointer, flags int) {
	// 0: (image const QImage &, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_FromImage(image unsafe.Pointer, flags int) {
	// 0: (image const QImage &, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	var nilthis *QPixmap
	nilthis.FromImage(image, flags)
}

// /usr/include/qt/QtGui/qpixmap.h:135
// index:1
// static inline
// QPixmap fromImage(class QImage &&, Qt::ImageConversionFlags)
func (this *QPixmap) FromImage_1(image unsafe.Pointer, flags int) {
	// 1: (image QImage &&, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap9fromImageEO6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_FromImage_1(image unsafe.Pointer, flags int) {
	// 1: (image QImage &&, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	var nilthis *QPixmap
	nilthis.FromImage_1(image, flags)
}

// /usr/include/qt/QtGui/qpixmap.h:133
// index:0
// static
// QPixmap fromImageReader(class QImageReader *, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageReader(imageReader unsafe.Pointer, flags int) {
	// 0: (imageReader QImageReader *, QFlags<Qt::ImageConversionFlag> flags), (imageReader, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap15fromImageReaderEP12QImageReader6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_FromImageReader(imageReader unsafe.Pointer, flags int) {
	// 0: (imageReader QImageReader *, QFlags<Qt::ImageConversionFlag> flags), (imageReader, flags)
	var nilthis *QPixmap
	nilthis.FromImageReader(imageReader, flags)
}

// /usr/include/qt/QtGui/qpixmap.h:141
// index:0
// bool load(const class QString &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) Load(fileName unsafe.Pointer, format unsafe.Pointer, flags int) {
	// 0: (, fileName const QString &, format const char *, QFlags<Qt::ImageConversionFlag> flags), (fileName, format, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap4loadERK7QStringPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, format, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:142
// index:0
// bool loadFromData(const uchar *, uint, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData(buf unsafe.Pointer, len uint, format unsafe.Pointer, flags int) {
	// 0: (, buf const uchar *, len uint, format const char *, QFlags<Qt::ImageConversionFlag> flags), (buf, &len, format, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataEPKhjPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), buf, &len, format, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:143
// index:1
// inline
// bool loadFromData(const class QByteArray &, const char *, Qt::ImageConversionFlags)
func (this *QPixmap) LoadFromData_1(data unsafe.Pointer, format unsafe.Pointer, flags int) {
	// 1: (, data const QByteArray &, format const char *, QFlags<Qt::ImageConversionFlag> flags), (data, format, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap12loadFromDataERK10QByteArrayPKc6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, format, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:144
// index:0
// bool save(const class QString &, const char *, int)
func (this *QPixmap) Save(fileName unsafe.Pointer, format unsafe.Pointer, quality int) {
	// 0: (, fileName const QString &, format const char *, quality int), (fileName, format, &quality)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4saveERK7QStringPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, format, &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:145
// index:1
// bool save(class QIODevice *, const char *, int)
func (this *QPixmap) Save_1(device unsafe.Pointer, format unsafe.Pointer, quality int) {
	// 1: (, device QIODevice *, format const char *, quality int), (device, format, &quality)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4saveEP9QIODevicePKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device, format, &quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:147
// index:0
// bool convertFromImage(const class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) ConvertFromImage(img unsafe.Pointer, flags int) {
	// 0: (, img const QImage &, QFlags<Qt::ImageConversionFlag> flags), (img, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap16convertFromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), img, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:149
// index:0
// inline
// QPixmap copy(int, int, int, int)
func (this *QPixmap) Copy(x int, y int, width int, height int) {
	// 0: (, x int, y int, width int, height int), (&x, &y, &width, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4copyEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &width, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:150
// index:1
// QPixmap copy(const class QRect &)
func (this *QPixmap) Copy_1(rect unsafe.Pointer) {
	// 1: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap4copyERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:152
// index:0
// inline
// void scroll(int, int, int, int, int, int, class QRegion *)
func (this *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed unsafe.Pointer) {
	// 0: (, dx int, dy int, x int, y int, width int, height int, exposed QRegion *), (&dx, &dy, &x, &y, &width, &height, exposed)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiiiiiP7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy, &x, &y, &width, &height, exposed)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:153
// index:1
// void scroll(int, int, const class QRect &, class QRegion *)
func (this *QPixmap) Scroll_1(dx int, dy int, rect unsafe.Pointer, exposed unsafe.Pointer) {
	// 1: (, dx int, dy int, rect const QRect &, exposed QRegion *), (&dx, &dy, rect, exposed)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap6scrollEiiRK5QRectP7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy, rect, exposed)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:158
// index:0
// qint64 cacheKey()
func (this *QPixmap) CacheKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap8cacheKeyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:160
// index:0
// bool isDetached()
func (this *QPixmap) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:161
// index:0
// void detach()
func (this *QPixmap) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap6detachEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:163
// index:0
// bool isQBitmap()
func (this *QPixmap) IsQBitmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap9isQBitmapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:165
// index:0
// virtual
// QPaintEngine * paintEngine()
func (this *QPixmap) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:175
// index:0
// virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPixmap) Metric(arg0 int) {
	// 0: (, QPaintDevice::PaintDeviceMetric arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:176
// index:0
// static
// QPixmap fromImageInPlace(class QImage &, Qt::ImageConversionFlags)
func (this *QPixmap) FromImageInPlace(image unsafe.Pointer, flags int) {
	// 0: (image QImage &, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QPixmap16fromImageInPlaceER6QImage6QFlagsIN2Qt19ImageConversionFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmap_FromImageInPlace(image unsafe.Pointer, flags int) {
	// 0: (image QImage &, QFlags<Qt::ImageConversionFlag> flags), (image, flags)
	var nilthis *QPixmap
	nilthis.FromImageInPlace(image, flags)
}

// /usr/include/qt/QtGui/qpixmap.h:198
// index:0
// QPlatformPixmap * handle()
func (this *QPixmap) Handle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QPixmap6handleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
