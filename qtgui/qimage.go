package qtgui

// /usr/include/qt/QtGui/qimage.h
// #include <qimage.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QImage) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// QImage mirrored_helper(bool, bool)
func (this *QImage) InheritMirrored_helper(f func(horizontal bool, vertical bool) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mirrored_helper", f)
}

// QImage rgbSwapped_helper()
func (this *QImage) InheritRgbSwapped_helper(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "rgbSwapped_helper", f)
}

// void mirrored_inplace(bool, bool)
func (this *QImage) InheritMirrored_inplace(f func(horizontal bool, vertical bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mirrored_inplace", f)
}

// void rgbSwapped_inplace()
func (this *QImage) InheritRgbSwapped_inplace(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "rgbSwapped_inplace", f)
}

// QImage convertToFormat_helper(QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) InheritConvertToFormat_helper(f func(format int, flags int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "convertToFormat_helper", f)
}

// bool convertToFormat_inplace(QImage::Format, Qt::ImageConversionFlags)
func (this *QImage) InheritConvertToFormat_inplace(f func(format int, flags int) bool) {
	qtrt.SetAllInheritCallback(this, "convertToFormat_inplace", f)
}

// QImage smoothScaled(int, int)
func (this *QImage) InheritSmoothScaled(f func(w int, h int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "smoothScaled", f)
}

/*

 */
type QImage struct {
	*QPaintDevice
}
type QImage_ITF interface {
	QPaintDevice_ITF
	QImage_PTR() *QImage
}

func (ptr *QImage) QImage_PTR() *QImage { return ptr }

func (this *QImage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QImage) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQImageFromPointer(cthis unsafe.Pointer) *QImage {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QImage{bcthis0}
}
func (*QImage) NewFromPointer(cthis unsafe.Pointer) *QImage {
	return NewQImageFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimage.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImage()

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit() *QImage {
	return NewQImage()
}
func NewQImage() *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:137
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QSize &, QImage::Format)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_1(size qtcore.QSize_ITF, format int) *QImage {
	return NewQImage_1(size, format)
}
func NewQImage_1(size qtcore.QSize_ITF, format int) *QImage {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2ERK5QSizeNS_6FormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:138
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImage(int, int, QImage::Format)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_2(width int, height int, format int) *QImage {
	return NewQImage_2(width, height, format)
}
func NewQImage_2(width int, height int, format int) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EiiNS_6FormatE", qtrt.FFI_TYPE_POINTER, width, height, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:139
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QImage(uchar *, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_3(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	return NewQImage_3(data, width, height, format, cleanupFunction, cleanupInfo)
}
func NewQImage_3(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPhiiNS_6FormatEPFvPvES2_", qtrt.FFI_TYPE_POINTER, data, width, height, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:139
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QImage(uchar *, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_3_(data unsafe.Pointer /*666*/, width int, height int, format int) *QImage {
	return NewQImage_3_(data, width, height, format)
}
func NewQImage_3_(data unsafe.Pointer /*666*/, width int, height int, format int) *QImage {
	// arg: 4, QImageCleanupFunction=Typedef, QImageCleanupFunction=Typedef, void (*)(void *), Pointer
	var cleanupFunction unsafe.Pointer
	// arg: 5, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPhiiNS_6FormatEPFvPvES2_", qtrt.FFI_TYPE_POINTER, data, width, height, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:139
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QImage(uchar *, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_3_1(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	return NewQImage_3_1(data, width, height, format, cleanupFunction)
}
func NewQImage_3_1(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	// arg: 5, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPhiiNS_6FormatEPFvPvES2_", qtrt.FFI_TYPE_POINTER, data, width, height, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:140
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QImage(const uchar *, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_4(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	return NewQImage_4(data, width, height, format, cleanupFunction, cleanupInfo)
}
func NewQImage_4(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKhiiNS_6FormatEPFvPvES3_", qtrt.FFI_TYPE_POINTER, data, width, height, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:140
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QImage(const uchar *, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_4_(data unsafe.Pointer /*666*/, width int, height int, format int) *QImage {
	return NewQImage_4_(data, width, height, format)
}
func NewQImage_4_(data unsafe.Pointer /*666*/, width int, height int, format int) *QImage {
	// arg: 4, QImageCleanupFunction=Typedef, QImageCleanupFunction=Typedef, void (*)(void *), Pointer
	var cleanupFunction unsafe.Pointer
	// arg: 5, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKhiiNS_6FormatEPFvPvES3_", qtrt.FFI_TYPE_POINTER, data, width, height, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:140
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QImage(const uchar *, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_4_1(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	return NewQImage_4_1(data, width, height, format, cleanupFunction)
}
func NewQImage_4_1(data unsafe.Pointer /*666*/, width int, height int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	// arg: 5, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKhiiNS_6FormatEPFvPvES3_", qtrt.FFI_TYPE_POINTER, data, width, height, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:141
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QImage(uchar *, int, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_5(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	return NewQImage_5(data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
}
func NewQImage_5(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPhiiiNS_6FormatEPFvPvES2_", qtrt.FFI_TYPE_POINTER, data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:141
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QImage(uchar *, int, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_5_(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int) *QImage {
	return NewQImage_5_(data, width, height, bytesPerLine, format)
}
func NewQImage_5_(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int) *QImage {
	// arg: 5, QImageCleanupFunction=Typedef, QImageCleanupFunction=Typedef, void (*)(void *), Pointer
	var cleanupFunction unsafe.Pointer
	// arg: 6, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPhiiiNS_6FormatEPFvPvES2_", qtrt.FFI_TYPE_POINTER, data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:141
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QImage(uchar *, int, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_5_1(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	return NewQImage_5_1(data, width, height, bytesPerLine, format, cleanupFunction)
}
func NewQImage_5_1(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	// arg: 6, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPhiiiNS_6FormatEPFvPvES2_", qtrt.FFI_TYPE_POINTER, data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:142
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QImage(const uchar *, int, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_6(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	return NewQImage_6(data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
}
func NewQImage_6(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/, cleanupInfo unsafe.Pointer /*666*/) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKhiiiNS_6FormatEPFvPvES3_", qtrt.FFI_TYPE_POINTER, data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:142
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QImage(const uchar *, int, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_6_(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int) *QImage {
	return NewQImage_6_(data, width, height, bytesPerLine, format)
}
func NewQImage_6_(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int) *QImage {
	// arg: 5, QImageCleanupFunction=Typedef, QImageCleanupFunction=Typedef, void (*)(void *), Pointer
	var cleanupFunction unsafe.Pointer
	// arg: 6, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKhiiiNS_6FormatEPFvPvES3_", qtrt.FFI_TYPE_POINTER, data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:142
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QImage(const uchar *, int, int, int, QImage::Format, QImageCleanupFunction, void *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_6_1(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	return NewQImage_6_1(data, width, height, bytesPerLine, format, cleanupFunction)
}
func NewQImage_6_1(data unsafe.Pointer /*666*/, width int, height int, bytesPerLine int, format int, cleanupFunction unsafe.Pointer /*666*/) *QImage {
	// arg: 6, void *=Pointer, =Invalid, , Invalid
	var cleanupInfo unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKhiiiNS_6FormatEPFvPvES3_", qtrt.FFI_TYPE_POINTER, data, width, height, bytesPerLine, format, cleanupFunction, cleanupInfo)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:145
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QImage(const char *const *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_7(xpm []string) *QImage {
	return NewQImage_7(xpm)
}
func NewQImage_7(xpm []string) *QImage {
	var convArg0 = qtrt.StringSliceToCCharPP(xpm)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2EPKPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:147
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QString &, const char *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_8(fileName string, format string) *QImage {
	return NewQImage_8(fileName, format)
}
func NewQImage_8(fileName string, format string) *QImage {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2ERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:147
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QImage(const QString &, const char *)

/*
Constructs a null image.

See also isNull().
*/
func (*QImage) NewForInherit_8_(fileName string) *QImage {
	return NewQImage_8_(fileName)
}
func NewQImage_8_(fileName string) *QImage {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageC2ERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImage)
	return gothis
}

// /usr/include/qt/QtGui/qimage.h:155
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImage()

/*

 */
func DeleteQImage(this *QImage) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimage.h:157
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage & operator=(const QImage &)

/*

 */
func (this *QImage) Operator_equal(arg0 QImage_ITF) *QImage {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:159
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage & operator=(QImage &&)

/*

 */
func (this *QImage) Operator_equal_1(other unsafe.Pointer /*333*/) *QImage {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImageaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QImage &)

/*
Swaps image other with this image. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QImage) Swap(other QImage_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QImage_PTR() != nil {
		convArg0 = other.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:165
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if it is a null image, otherwise returns false.

A null image has all parameters set to zero and no allocated data.
*/
func (this *QImage) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const

/*

 */
func (this *QImage) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QImage &) const

/*

 */
func (this *QImage) Operator_equal_equal(arg0 QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImageeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QImage &) const

/*

 */
func (this *QImage) Operator_not_equal(arg0 QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QImage_PTR() != nil {
		convArg0 = arg0.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImageneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QImage) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QImage) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:175
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage copy(const QRect &) const

/*
Returns a sub-area of the image as a new image.

The returned image is copied from the position (rectangle.x(), rectangle.y()) in this image, and will always have the size of the given rectangle.

In areas beyond this image, pixels are set to 0. For 32-bit RGB images, this means black; for 32-bit ARGB images, this means transparent black; for 8-bit images, this means the color with index 0 in the color table which can be anything; for 1-bit images, this means Qt::color0.

If the given rectangle is a null rectangle the entire image is copied.

See also QImage().
*/
func (this *QImage) Copy(rect qtcore.QRect_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:175
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage copy(const QRect &) const

/*
Returns a sub-area of the image as a new image.

The returned image is copied from the position (rectangle.x(), rectangle.y()) in this image, and will always have the size of the given rectangle.

In areas beyond this image, pixels are set to 0. For 32-bit RGB images, this means black; for 32-bit ARGB images, this means transparent black; for 8-bit images, this means the color with index 0 in the color table which can be anything; for 1-bit images, this means Qt::color0.

If the given rectangle is a null rectangle the entire image is copied.

See also QImage().
*/
func (this *QImage) Copy__() *QImage /*123*/ {
	// arg: 0, const QRect &=LValueReference, QRect=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4copyERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:176
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage copy(int, int, int, int) const

/*
Returns a sub-area of the image as a new image.

The returned image is copied from the position (rectangle.x(), rectangle.y()) in this image, and will always have the size of the given rectangle.

In areas beyond this image, pixels are set to 0. For 32-bit RGB images, this means black; for 32-bit ARGB images, this means transparent black; for 8-bit images, this means the color with index 0 in the color table which can be anything; for 1-bit images, this means Qt::color0.

If the given rectangle is a null rectangle the entire image is copied.

See also QImage().
*/
func (this *QImage) Copy_1(x int, y int, w int, h int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4copyEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] QImage::Format format() const

/*
Returns the format of the image.

See also Image Formats.
*/
func (this *QImage) Format() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimage.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(QImage::Format, Qt::ImageConversionFlags) const

/*
Returns a copy of the image in the given format.

The specified image conversion flags control how the image data is handled during the conversion process.

See also Image Formats.
*/
func (this *QImage) ConvertToFormat(f int, flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(QImage::Format, Qt::ImageConversionFlags) const

/*
Returns a copy of the image in the given format.

The specified image conversion flags control how the image data is handled during the conversion process.

See also Image Formats.
*/
func (this *QImage) ConvertToFormat__(f int) *QImage /*123*/ {
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:184
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(QImage::Format, Qt::ImageConversionFlags)

/*
Returns a copy of the image in the given format.

The specified image conversion flags control how the image data is handled during the conversion process.

See also Image Formats.
*/
func (this *QImage) ConvertToFormat_1(f int, flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:184
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage convertToFormat(QImage::Format, Qt::ImageConversionFlags)

/*
Returns a copy of the image in the given format.

The specified image conversion flags control how the image data is handled during the conversion process.

See also Image Formats.
*/
func (this *QImage) ConvertToFormat_1_(f int) *QImage /*123*/ {
	// arg: 1, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage15convertToFormatENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:195
// index:0
// Public Visibility=Default Availability=Available
// [1] bool reinterpretAsFormat(QImage::Format)

/*
Changes the format of the image without changing the data. Only works between formats of the same depth.

Returns true if successful.

This function can be used to change images with alpha-channels to their corresponding opaque formats if the data is known to be opaque-only, or to change the format of a given image buffer before overwriting it with new data.

Warning: The function does not check if the image data is valid in the new format and will still return true if the depths are compatible. Operations on an image with invalid data are undefined.

Warning: If the image is not detached, this will cause the data to be copied.

This function was introduced in  Qt 5.9.

See also hasAlphaChannel() and convertToFormat().
*/
func (this *QImage) ReinterpretAsFormat(f int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage19reinterpretAsFormatENS_6FormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:197
// index:0
// Public Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the width of the image.

See also Image Information.
*/
func (this *QImage) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:198
// index:0
// Public Visibility=Default Availability=Available
// [4] int height() const

/*
Returns the height of the image.

See also Image Information.
*/
func (this *QImage) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the size of the image, i.e. its width() and height().

See also Image Information.
*/
func (this *QImage) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:200
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rect() const

/*
Returns the enclosing rectangle (0, 0, width(), height()) of the image.

See also Image Information.
*/
func (this *QImage) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:202
// index:0
// Public Visibility=Default Availability=Available
// [4] int depth() const

/*
Returns the depth of the image.

The image depth is the number of bits used to store a single pixel, also called bits per pixel (bpp).

The supported depths are 1, 8, 16, 24 and 32.

See also bitPlaneCount(), convertToFormat(), Image Formats, and Image Information.
*/
func (this *QImage) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:203
// index:0
// Public Visibility=Default Availability=Available
// [4] int colorCount() const

/*
Returns the size of the color table for the image.

Notice that colorCount() returns 0 for 32-bpp images because these images do not use color tables, but instead encode pixel values as ARGB quadruplets.

This function was introduced in  Qt 4.6.

See also setColorCount() and Image Information.
*/
func (this *QImage) ColorCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10colorCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:204
// index:0
// Public Visibility=Default Availability=Available
// [4] int bitPlaneCount() const

/*
Returns the number of bit planes in the image.

The number of bit planes is the number of bits of color and transparency information for each pixel. This is different from (i.e. smaller than) the depth when the image format contains unused bits.

This function was introduced in  Qt 4.7.

See also depth(), format(), and Image Formats.
*/
func (this *QImage) BitPlaneCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13bitPlaneCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:206
// index:0
// Public Visibility=Default Availability=Available
// [4] QRgb color(int) const

/*
Returns the color in the color table at index i. The first color is at index 0.

The colors in an image's color table are specified as ARGB quadruplets (QRgb). Use the qAlpha(), qRed(), qGreen(), and qBlue() functions to get the color value components.

See also setColor(), pixelIndex(), and Pixel Manipulation.
*/
func (this *QImage) Color(i int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5colorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(int, QRgb)

/*
Sets the color at the given index in the color table, to the given to colorValue. The color value is an ARGB quadruplet.

If index is outside the current size of the color table, it is expanded with setColorCount().

See also color(), colorCount(), setColorTable(), and Pixel Manipulation.
*/
func (this *QImage) SetColor(i int, c uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8setColorEij", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorCount(int)

/*
Resizes the color table to contain colorCount entries.

If the color table is expanded, all the extra colors will be set to transparent (i.e qRgba(0, 0, 0, 0)).

When the image is used, the color table must be large enough to have entries for all the pixel/index values present in the image, otherwise the results are undefined.

This function was introduced in  Qt 4.6.

See also colorCount(), colorTable(), setColor(), and Image Transformations.
*/
func (this *QImage) SetColorCount(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13setColorCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:210
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allGray() const

/*
Returns true if all the colors in the image are shades of gray (i.e. their red, green and blue components are equal); otherwise false.

Note that this function is slow for images without color table.

See also isGrayscale().
*/
func (this *QImage) AllGray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage7allGrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGrayscale() const

/*
For 32-bit images, this function is equivalent to allGray().

For color indexed images, this function returns true if color(i) is QRgb(i, i, i) for all indexes of the color table; otherwise returns false.

See also allGray() and Image Formats.
*/
func (this *QImage) IsGrayscale() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11isGrayscaleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * bits()

/*
Returns a pointer to the first pixel data. This is equivalent to scanLine(0).

Note that QImage uses implicit data sharing. This function performs a deep copy of the shared pixel data, thus ensuring that this QImage is the only one using the current return value.

See also scanLine(), sizeInBytes(), and constBits().
*/
func (this *QImage) Bits() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4bitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:214
// index:1
// Public Visibility=Default Availability=Available
// [8] const uchar * bits() const

/*
Returns a pointer to the first pixel data. This is equivalent to scanLine(0).

Note that QImage uses implicit data sharing. This function performs a deep copy of the shared pixel data, thus ensuring that this QImage is the only one using the current return value.

See also scanLine(), sizeInBytes(), and constBits().
*/
func (this *QImage) Bits_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4bitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:215
// index:0
// Public Visibility=Default Availability=Available
// [8] const uchar * constBits() const

/*
Returns a pointer to the first pixel data.

Note that QImage uses implicit data sharing, but this function does not perform a deep copy of the shared pixel data, because the returned data is const.

This function was introduced in  Qt 4.7.

See also bits() and constScanLine().
*/
func (this *QImage) ConstBits() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage9constBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:218
// index:0
// Public Visibility=Default Availability=Available
// [4] int byteCount() const

/*

 */
func (this *QImage) ByteCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage9byteCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] qsizetype sizeInBytes() const

/*
Returns the image data size in bytes.

This function was introduced in  Qt 5.10.

See also byteCount(), bytesPerLine(), bits(), and Image Information.
*/
func (this *QImage) SizeInBytes() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11sizeInBytesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] uchar * scanLine(int)

/*
Returns a pointer to the pixel data at the scanline with index i. The first scanline is at index 0.

The scanline data is aligned on a 32-bit boundary.

Warning: If you are accessing 32-bpp image data, cast the returned pointer to QRgb* (QRgb has a 32-bit size) and use it to read/write the pixel value. You cannot use the uchar* pointer directly, because the pixel format depends on the byte order on the underlying platform. Use qRed(), qGreen(), qBlue(), and qAlpha() to access the pixels.

See also bytesPerLine(), bits(), Pixel Manipulation, and constScanLine().
*/
func (this *QImage) ScanLine(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8scanLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:223
// index:1
// Public Visibility=Default Availability=Available
// [8] const uchar * scanLine(int) const

/*
Returns a pointer to the pixel data at the scanline with index i. The first scanline is at index 0.

The scanline data is aligned on a 32-bit boundary.

Warning: If you are accessing 32-bpp image data, cast the returned pointer to QRgb* (QRgb has a 32-bit size) and use it to read/write the pixel value. You cannot use the uchar* pointer directly, because the pixel format depends on the byte order on the underlying platform. Use qRed(), qGreen(), qBlue(), and qAlpha() to access the pixels.

See also bytesPerLine(), bits(), Pixel Manipulation, and constScanLine().
*/
func (this *QImage) ScanLine_1(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage8scanLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] const uchar * constScanLine(int) const

/*
Returns a pointer to the pixel data at the scanline with index i. The first scanline is at index 0.

The scanline data is aligned on a 32-bit boundary.

Note that QImage uses implicit data sharing, but this function does not perform a deep copy of the shared pixel data, because the returned data is const.

This function was introduced in  Qt 4.7.

See also scanLine() and constBits().
*/
func (this *QImage) ConstScanLine(arg0 int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13constScanLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qimage.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int bytesPerLine() const

/*
Returns the number of bytes per image scanline.

This is equivalent to sizeInBytes() / height() if height() is non-zero.

See also scanLine().
*/
func (this *QImage) BytesPerLine() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage12bytesPerLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:227
// index:0
// Public Visibility=Default Availability=Available
// [1] bool valid(int, int) const

/*
Returns true if pos is a valid coordinate pair within the image; otherwise returns false.

See also rect() and QRect::contains().
*/
func (this *QImage) Valid(x int, y int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5validEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:228
// index:1
// Public Visibility=Default Availability=Available
// [1] bool valid(const QPoint &) const

/*
Returns true if pos is a valid coordinate pair within the image; otherwise returns false.

See also rect() and QRect::contains().
*/
func (this *QImage) Valid_1(pt qtcore.QPoint_ITF) bool {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5validERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:230
// index:0
// Public Visibility=Default Availability=Available
// [4] int pixelIndex(int, int) const

/*
Returns the pixel index at the given position.

If position is not valid, or if the image is not a paletted image (depth() > 8), the results are undefined.

See also valid(), depth(), and Pixel Manipulation.
*/
func (this *QImage) PixelIndex(x int, y int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelIndexEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:231
// index:1
// Public Visibility=Default Availability=Available
// [4] int pixelIndex(const QPoint &) const

/*
Returns the pixel index at the given position.

If position is not valid, or if the image is not a paletted image (depth() > 8), the results are undefined.

See also valid(), depth(), and Pixel Manipulation.
*/
func (this *QImage) PixelIndex_1(pt qtcore.QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelIndexERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:233
// index:0
// Public Visibility=Default Availability=Available
// [4] QRgb pixel(int, int) const

/*
Returns the color of the pixel at the given position.

If the position is not valid, the results are undefined.

Warning: This function is expensive when used for massive pixel manipulations. Use constBits() or constScanLine() when many pixels needs to be read.

See also setPixel(), valid(), constBits(), constScanLine(), and Pixel Manipulation.
*/
func (this *QImage) Pixel(x int, y int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5pixelEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:234
// index:1
// Public Visibility=Default Availability=Available
// [4] QRgb pixel(const QPoint &) const

/*
Returns the color of the pixel at the given position.

If the position is not valid, the results are undefined.

Warning: This function is expensive when used for massive pixel manipulations. Use constBits() or constScanLine() when many pixels needs to be read.

See also setPixel(), valid(), constBits(), constScanLine(), and Pixel Manipulation.
*/
func (this *QImage) Pixel_1(pt qtcore.QPoint_ITF) uint {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage5pixelERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixel(int, int, uint)

/*
Sets the pixel index or color at the given position to index_or_rgb.

If the image's format is either monochrome or paletted, the given index_or_rgb value must be an index in the image's color table, otherwise the parameter must be a QRgb value.

If position is not a valid coordinate pair in the image, or if index_or_rgb >= colorCount() in the case of monochrome and paletted images, the result is undefined.

Warning: This function is expensive due to the call of the internal detach() function called within; if performance is a concern, we recommend the use of scanLine() or bits() to access pixel data directly.

See also pixel() and Pixel Manipulation.
*/
func (this *QImage) SetPixel(x int, y int, index_or_rgb uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8setPixelEiij", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, index_or_rgb)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:237
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPixel(const QPoint &, uint)

/*
Sets the pixel index or color at the given position to index_or_rgb.

If the image's format is either monochrome or paletted, the given index_or_rgb value must be an index in the image's color table, otherwise the parameter must be a QRgb value.

If position is not a valid coordinate pair in the image, or if index_or_rgb >= colorCount() in the case of monochrome and paletted images, the result is undefined.

Warning: This function is expensive due to the call of the internal detach() function called within; if performance is a concern, we recommend the use of scanLine() or bits() to access pixel data directly.

See also pixel() and Pixel Manipulation.
*/
func (this *QImage) SetPixel_1(pt qtcore.QPoint_ITF, index_or_rgb uint) {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8setPixelERK6QPointj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, index_or_rgb)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:239
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor pixelColor(int, int) const

/*
Returns the color of the pixel at the given position as a QColor.

If the position is not valid, an invalid QColor is returned.

Warning: This function is expensive when used for massive pixel manipulations. Use constBits() or constScanLine() when many pixels needs to be read.

This function was introduced in  Qt 5.6.

See also setPixelColor(), setPixel(), valid(), constBits(), constScanLine(), and Pixel Manipulation.
*/
func (this *QImage) PixelColor(x int, y int) *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelColorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:240
// index:1
// Public Visibility=Default Availability=Available
// [16] QColor pixelColor(const QPoint &) const

/*
Returns the color of the pixel at the given position as a QColor.

If the position is not valid, an invalid QColor is returned.

Warning: This function is expensive when used for massive pixel manipulations. Use constBits() or constScanLine() when many pixels needs to be read.

This function was introduced in  Qt 5.6.

See also setPixelColor(), setPixel(), valid(), constBits(), constScanLine(), and Pixel Manipulation.
*/
func (this *QImage) PixelColor_1(pt qtcore.QPoint_ITF) *QColor /*123*/ {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage10pixelColorERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelColor(int, int, const QColor &)

/*
Sets the color at the given position to color.

If position is not a valid coordinate pair in the image, or the image's format is either monochrome or paletted, the result is undefined.

Warning: This function is expensive due to the call of the internal detach() function called within; if performance is a concern, we recommend the use of scanLine() or bits() to access pixel data directly.

This function was introduced in  Qt 5.6.

See also pixelColor(), pixel(), bits(), scanLine(), and Pixel Manipulation.
*/
func (this *QImage) SetPixelColor(x int, y int, c QColor_ITF) {
	var convArg2 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg2 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13setPixelColorEiiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:243
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPixelColor(const QPoint &, const QColor &)

/*
Sets the color at the given position to color.

If position is not a valid coordinate pair in the image, or the image's format is either monochrome or paletted, the result is undefined.

Warning: This function is expensive due to the call of the internal detach() function called within; if performance is a concern, we recommend the use of scanLine() or bits() to access pixel data directly.

This function was introduced in  Qt 5.6.

See also pixelColor(), pixel(), bits(), scanLine(), and Pixel Manipulation.
*/
func (this *QImage) SetPixelColor_1(pt qtcore.QPoint_ITF, c QColor_ITF) {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg1 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13setPixelColorERK6QPointRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:252
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const

/*
Returns the device pixel ratio for the image. This is the ratio between device pixels and device independent pixels.

Use this function when calculating layout geometry based on the image size: QSize layoutSize = image.size() / image.devicePixelRatio()

The default value is 1.0.

See also setDevicePixelRatio() and QImageReader.
*/
func (this *QImage) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qimage.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevicePixelRatio(qreal)

/*
Sets the device pixel ratio for the image. This is the ratio between image pixels and device-independent pixels.

The default scaleFactor is 1.0. Setting it to something else has two effects:

QPainters that are opened on the image will be scaled. For example, painting on a 200x200 image if with a ratio of 2.0 will result in effective (device-independent) painting bounds of 100x100.

Code paths in Qt that calculate layout geometry based on the image size will take the ratio into account: QSize layoutSize = image.size() / image.devicePixelRatio() The net effect of this is that the image is displayed as high-DPI image rather than a large image (see Drawing High Resolution Versions of Pixmaps and Images).

See also devicePixelRatio().
*/
func (this *QImage) SetDevicePixelRatio(scaleFactor float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage19setDevicePixelRatioEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scaleFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fill(uint)

/*
Fills the entire image with the given pixelValue.

If the depth of this image is 1, only the lowest bit is used. If you say fill(0), fill(2), etc., the image is filled with 0s. If you say fill(1), fill(3), etc., the image is filled with 1s. If the depth is 8, the lowest 8 bits are used and if the depth is 16 the lowest 16 bits are used.

Note: QImage::pixel() returns the color of the pixel at the given coordinates while QColor::pixel() returns the pixel value of the underlying window system (essentially an index value), so normally you will want to use QImage::pixel() to use a color from an existing image or QColor::rgb() to use a specific color.

See also depth() and Image Transformations.
*/
func (this *QImage) Fill(pixel uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4fillEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pixel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:256
// index:1
// Public Visibility=Default Availability=Available
// [-2] void fill(const QColor &)

/*
Fills the entire image with the given pixelValue.

If the depth of this image is 1, only the lowest bit is used. If you say fill(0), fill(2), etc., the image is filled with 0s. If you say fill(1), fill(3), etc., the image is filled with 1s. If the depth is 8, the lowest 8 bits are used and if the depth is 16 the lowest 16 bits are used.

Note: QImage::pixel() returns the color of the pixel at the given coordinates while QColor::pixel() returns the pixel value of the underlying window system (essentially an index value), so normally you will want to use QImage::pixel() to use a color from an existing image or QColor::rgb() to use a specific color.

See also depth() and Image Transformations.
*/
func (this *QImage) Fill_1(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4fillERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:257
// index:2
// Public Visibility=Default Availability=Available
// [-2] void fill(Qt::GlobalColor)

/*
Fills the entire image with the given pixelValue.

If the depth of this image is 1, only the lowest bit is used. If you say fill(0), fill(2), etc., the image is filled with 0s. If you say fill(1), fill(3), etc., the image is filled with 1s. If the depth is 8, the lowest 8 bits are used and if the depth is 16 the lowest 16 bits are used.

Note: QImage::pixel() returns the color of the pixel at the given coordinates while QColor::pixel() returns the pixel value of the underlying window system (essentially an index value), so normally you will want to use QImage::pixel() to use a color from an existing image or QColor::rgb() to use a specific color.

See also depth() and Image Transformations.
*/
func (this *QImage) Fill_2(color int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4fillEN2Qt11GlobalColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlphaChannel() const

/*
Returns true if the image has a format that respects the alpha channel, otherwise returns false.

See also Image Information.
*/
func (this *QImage) HasAlphaChannel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15hasAlphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlphaChannel(const QImage &)

/*

 */
func (this *QImage) SetAlphaChannel(alphaChannel QImage_ITF) {
	var convArg0 unsafe.Pointer
	if alphaChannel != nil && alphaChannel.QImage_PTR() != nil {
		convArg0 = alphaChannel.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage15setAlphaChannelERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:262
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage alphaChannel() const

/*

 */
func (this *QImage) AlphaChannel() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage12alphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:263
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createAlphaMask(Qt::ImageConversionFlags) const

/*
Builds and returns a 1-bpp mask from the alpha buffer in this image. Returns a null image if the image's format is QImage::Format_RGB32.

The flags argument is a bitwise-OR of the Qt::ImageConversionFlags, and controls the conversion process. Passing 0 for flags sets all the default options.

The returned image has little-endian bit order (i.e. the image's format is QImage::Format_MonoLSB), which you can convert to big-endian (QImage::Format_Mono) using the convertToFormat() function.

See also createHeuristicMask() and Image Transformations.
*/
func (this *QImage) CreateAlphaMask(flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15createAlphaMaskE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:263
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createAlphaMask(Qt::ImageConversionFlags) const

/*
Builds and returns a 1-bpp mask from the alpha buffer in this image. Returns a null image if the image's format is QImage::Format_RGB32.

The flags argument is a bitwise-OR of the Qt::ImageConversionFlags, and controls the conversion process. Passing 0 for flags sets all the default options.

The returned image has little-endian bit order (i.e. the image's format is QImage::Format_MonoLSB), which you can convert to big-endian (QImage::Format_Mono) using the convertToFormat() function.

See also createHeuristicMask() and Image Transformations.
*/
func (this *QImage) CreateAlphaMask__() *QImage /*123*/ {
	// arg: 0, Qt::ImageConversionFlags=Elaborated, Qt::ImageConversionFlags=Typedef, QFlags<Qt::ImageConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15createAlphaMaskE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:265
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createHeuristicMask(bool) const

/*
Creates and returns a 1-bpp heuristic mask for this image.

The function works by selecting a color from one of the corners, then chipping away pixels of that color starting at all the edges. The four corners vote for which color is to be masked away. In case of a draw (this generally means that this function is not applicable to the image), the result is arbitrary.

The returned image has little-endian bit order (i.e. the image's format is QImage::Format_MonoLSB), which you can convert to big-endian (QImage::Format_Mono) using the convertToFormat() function.

If clipTight is true (the default) the mask is just large enough to cover the pixels; otherwise, the mask is larger than the data pixels.

Note that this function disregards the alpha buffer.

See also createAlphaMask() and Image Transformations.
*/
func (this *QImage) CreateHeuristicMask(clipTight bool) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:265
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createHeuristicMask(bool) const

/*
Creates and returns a 1-bpp heuristic mask for this image.

The function works by selecting a color from one of the corners, then chipping away pixels of that color starting at all the edges. The four corners vote for which color is to be masked away. In case of a draw (this generally means that this function is not applicable to the image), the result is arbitrary.

The returned image has little-endian bit order (i.e. the image's format is QImage::Format_MonoLSB), which you can convert to big-endian (QImage::Format_Mono) using the convertToFormat() function.

If clipTight is true (the default) the mask is just large enough to cover the pixels; otherwise, the mask is larger than the data pixels.

Note that this function disregards the alpha buffer.

See also createAlphaMask() and Image Transformations.
*/
func (this *QImage) CreateHeuristicMask__() *QImage /*123*/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	clipTight := true
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createHeuristicMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clipTight)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:267
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createMaskFromColor(QRgb, Qt::MaskMode) const

/*
Creates and returns a mask for this image based on the given color value. If the mode is MaskInColor (the default value), all pixels matching color will be opaque pixels in the mask. If mode is MaskOutColor, all pixels matching the given color will be transparent.

See also createAlphaMask() and createHeuristicMask().
*/
func (this *QImage) CreateMaskFromColor(color uint, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createMaskFromColorEjN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:267
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage createMaskFromColor(QRgb, Qt::MaskMode) const

/*
Creates and returns a mask for this image based on the given color value. If the mode is MaskInColor (the default value), all pixels matching color will be opaque pixels in the mask. If mode is MaskOutColor, all pixels matching the given color will be transparent.

See also createAlphaMask() and createHeuristicMask().
*/
func (this *QImage) CreateMaskFromColor__(color uint) *QImage /*123*/ {
	// arg: 1, Qt::MaskMode=Elaborated, Qt::MaskMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage19createMaskFromColorEjN2Qt8MaskModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), color, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Returns a copy of the image scaled to a rectangle defined by the given size according to the given aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the image is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the image is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the image is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null image.

See also isNull() and Image Transformations.
*/
func (this *QImage) Scaled(w int, h int, aspectMode int, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Returns a copy of the image scaled to a rectangle defined by the given size according to the given aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the image is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the image is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the image is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null image.

See also isNull() and Image Transformations.
*/
func (this *QImage) Scaled__(w int, h int) *QImage /*123*/ {
	// arg: 2, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectMode := 0
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage scaled(int, int, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Returns a copy of the image scaled to a rectangle defined by the given size according to the given aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the image is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the image is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the image is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null image.

See also isNull() and Image Transformations.
*/
func (this *QImage) Scaled__1(w int, h int, aspectMode int) *QImage /*123*/ {
	// arg: 3, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledEiiN2Qt15AspectRatioModeENS0_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Returns a copy of the image scaled to a rectangle defined by the given size according to the given aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the image is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the image is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the image is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null image.

See also isNull() and Image Transformations.
*/
func (this *QImage) Scaled_1(s qtcore.QSize_ITF, aspectMode int, mode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Returns a copy of the image scaled to a rectangle defined by the given size according to the given aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the image is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the image is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the image is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null image.

See also isNull() and Image Transformations.
*/
func (this *QImage) Scaled_1_(s qtcore.QSize_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 1, Qt::AspectRatioMode=Elaborated, Qt::AspectRatioMode=Enum, , Invalid
	aspectMode := 0
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:272
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage scaled(const QSize &, Qt::AspectRatioMode, Qt::TransformationMode) const

/*
Returns a copy of the image scaled to a rectangle defined by the given size according to the given aspectRatioMode and transformMode.




If aspectRatioMode is Qt::IgnoreAspectRatio, the image is scaled to size.
If aspectRatioMode is Qt::KeepAspectRatio, the image is scaled to a rectangle as large as possible inside size, preserving the aspect ratio.
If aspectRatioMode is Qt::KeepAspectRatioByExpanding, the image is scaled to a rectangle as small as possible outside size, preserving the aspect ratio.


If the given size is empty, this function returns a null image.

See also isNull() and Image Transformations.
*/
func (this *QImage) Scaled_1_1(s qtcore.QSize_ITF, aspectMode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg0 = s.QSize_PTR().GetCthis()
	}
	// arg: 2, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6scaledERK5QSizeN2Qt15AspectRatioModeENS3_18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, aspectMode, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:274
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToWidth(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given width using the specified transformation mode.

This function automatically calculates the height of the image so that its aspect ratio is preserved.

If the given width is 0 or negative, a null image is returned.

See also Image Transformations.
*/
func (this *QImage) ScaledToWidth(w int, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:274
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToWidth(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given width using the specified transformation mode.

This function automatically calculates the height of the image so that its aspect ratio is preserved.

If the given width is 0 or negative, a null image is returned.

See also Image Transformations.
*/
func (this *QImage) ScaledToWidth__(w int) *QImage /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13scaledToWidthEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:275
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToHeight(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given height using the specified transformation mode.

This function automatically calculates the width of the image so that the ratio of the image is preserved.

If the given height is 0 or negative, a null image is returned.

See also Image Transformations.
*/
func (this *QImage) ScaledToHeight(h int, mode int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:275
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage scaledToHeight(int, Qt::TransformationMode) const

/*
Returns a scaled copy of the image. The returned image is scaled to the given height using the specified transformation mode.

This function automatically calculates the width of the image so that the ratio of the image is preserved.

If the given height is 0 or negative, a null image is returned.

See also Image Transformations.
*/
func (this *QImage) ScaledToHeight__(h int) *QImage /*123*/ {
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage14scaledToHeightEiN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:276
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QMatrix &, Qt::TransformationMode) const

/*
Returns a copy of the image that is transformed using the given transformation matrix and transformation mode.

The transformation matrix is internally adjusted to compensate for unwanted translation; i.e. the image produced is the smallest image that contains all the transformed points of the original image. Use the trueMatrix() function to retrieve the actual matrix used for transforming an image.

See also trueMatrix() and Image Transformations.
*/
func (this *QImage) Transformed(matrix QMatrix_ITF, mode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:276
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QMatrix &, Qt::TransformationMode) const

/*
Returns a copy of the image that is transformed using the given transformation matrix and transformation mode.

The transformation matrix is internally adjusted to compensate for unwanted translation; i.e. the image produced is the smallest image that contains all the transformed points of the original image. Use the trueMatrix() function to retrieve the actual matrix used for transforming an image.

See also trueMatrix() and Image Transformations.
*/
func (this *QImage) Transformed__(matrix QMatrix_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK7QMatrixN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:278
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QTransform &, Qt::TransformationMode) const

/*
Returns a copy of the image that is transformed using the given transformation matrix and transformation mode.

The transformation matrix is internally adjusted to compensate for unwanted translation; i.e. the image produced is the smallest image that contains all the transformed points of the original image. Use the trueMatrix() function to retrieve the actual matrix used for transforming an image.

See also trueMatrix() and Image Transformations.
*/
func (this *QImage) Transformed_1(matrix QTransform_ITF, mode int) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:278
// index:1
// Public Visibility=Default Availability=Available
// [32] QImage transformed(const QTransform &, Qt::TransformationMode) const

/*
Returns a copy of the image that is transformed using the given transformation matrix and transformation mode.

The transformation matrix is internally adjusted to compensate for unwanted translation; i.e. the image produced is the smallest image that contains all the transformed points of the original image. Use the trueMatrix() function to retrieve the actual matrix used for transforming an image.

See also trueMatrix() and Image Transformations.
*/
func (this *QImage) Transformed_1_(matrix QTransform_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, Qt::TransformationMode=Elaborated, Qt::TransformationMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11transformedERK10QTransformN2Qt18TransformationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [48] QMatrix trueMatrix(const QMatrix &, int, int)

/*
Returns the actual matrix used for transforming an image with the given width, height and matrix.

When transforming an image using the transformed() function, the transformation matrix is internally adjusted to compensate for unwanted translation, i.e. transformed() returns the smallest image containing all transformed points of the original image. This function returns the modified matrix, which maps points correctly from the original image into the new image.

See also transformed() and Image Transformations.
*/
func (this *QImage) TrueMatrix(arg0 QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMatrix_PTR() != nil {
		convArg0 = arg0.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage10trueMatrixERK7QMatrixii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}
func QImage_TrueMatrix(arg0 QMatrix_ITF, w int, h int) *QMatrix /*123*/ {
	var nilthis *QImage
	rv := nilthis.TrueMatrix(arg0, w, h)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:279
// index:1
// Public static Visibility=Default Availability=Available
// [88] QTransform trueMatrix(const QTransform &, int, int)

/*
Returns the actual matrix used for transforming an image with the given width, height and matrix.

When transforming an image using the transformed() function, the transformation matrix is internally adjusted to compensate for unwanted translation, i.e. transformed() returns the smallest image containing all transformed points of the original image. This function returns the modified matrix, which maps points correctly from the original image into the new image.

See also transformed() and Image Transformations.
*/
func (this *QImage) TrueMatrix_1(arg0 QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTransform_PTR() != nil {
		convArg0 = arg0.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage10trueMatrixERK10QTransformii", qtrt.FFI_TYPE_POINTER, convArg0, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QImage_TrueMatrix_1(arg0 QTransform_ITF, w int, h int) *QTransform /*123*/ {
	var nilthis *QImage
	rv := nilthis.TrueMatrix_1(arg0, w, h)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage mirrored(bool, bool) const

/*
Returns a mirror of the image, mirrored in the horizontal and/or the vertical direction depending on whether horizontal and vertical are set to true or false.

Note that the original image is not changed.

See also Image Transformations.
*/
func (this *QImage) Mirrored(horizontally bool, vertically bool) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage mirrored(bool, bool) const

/*
Returns a mirror of the image, mirrored in the horizontal and/or the vertical direction depending on whether horizontal and vertical are set to true or false.

Note that the original image is not changed.

See also Image Transformations.
*/
func (this *QImage) Mirrored__() *QImage /*123*/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	horizontally := false
	// arg: 1, bool=Bool, =Invalid, , Invalid
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage mirrored(bool, bool) const

/*
Returns a mirror of the image, mirrored in the horizontal and/or the vertical direction depending on whether horizontal and vertical are set to true or false.

Note that the original image is not changed.

See also Image Transformations.
*/
func (this *QImage) Mirrored__1(horizontally bool) *QImage /*123*/ {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && mirrored(bool, bool)

/*
Returns a mirror of the image, mirrored in the horizontal and/or the vertical direction depending on whether horizontal and vertical are set to true or false.

Note that the original image is not changed.

See also Image Transformations.
*/
func (this *QImage) Mirrored_1(horizontally bool, vertically bool) unsafe.Pointer /*333*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && mirrored(bool, bool)

/*
Returns a mirror of the image, mirrored in the horizontal and/or the vertical direction depending on whether horizontal and vertical are set to true or false.

Note that the original image is not changed.

See also Image Transformations.
*/
func (this *QImage) Mirrored_1_() unsafe.Pointer /*333*/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	horizontally := false
	// arg: 1, bool=Bool, =Invalid, , Invalid
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && mirrored(bool, bool)

/*
Returns a mirror of the image, mirrored in the horizontal and/or the vertical direction depending on whether horizontal and vertical are set to true or false.

Note that the original image is not changed.

See also Image Transformations.
*/
func (this *QImage) Mirrored_1_1(horizontally bool) unsafe.Pointer /*333*/ {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	vertically := true
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage8mirroredEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontally, vertically)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:285
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QImage rgbSwapped() const

/*
Returns a QImage in which the values of the red and blue components of all pixels have been swapped, effectively converting an RGB image to an BGR image.

The original QImage is not changed.

See also Image Transformations.
*/
func (this *QImage) RgbSwapped() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR6QImage10rgbSwappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:287
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QImage && rgbSwapped()

/*
Returns a QImage in which the values of the red and blue components of all pixels have been swapped, effectively converting an RGB image to an BGR image.

The original QImage is not changed.

See also Image Transformations.
*/
func (this *QImage) RgbSwapped_1() unsafe.Pointer /*333*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO6QImage10rgbSwappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv)) //777
}

// /usr/include/qt/QtGui/qimage.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invertPixels(QImage::InvertMode)

/*
Inverts all pixel values in the image.

The given invert mode only have a meaning when the image's depth is 32. The default mode is InvertRgb, which leaves the alpha channel unchanged. If the mode is InvertRgba, the alpha bits are also inverted.

Inverting an 8-bit image means to replace all pixels using color index i with a pixel using color index 255 minus i. The same is the case for a 1-bit image. Note that the color table is not changed.

If the image has a premultiplied alpha channel, the image is first converted to ARGB32 to be inverted and then converted back.

See also Image Transformations.
*/
func (this *QImage) InvertPixels(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12invertPixelsENS_10InvertModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invertPixels(QImage::InvertMode)

/*
Inverts all pixel values in the image.

The given invert mode only have a meaning when the image's depth is 32. The default mode is InvertRgb, which leaves the alpha channel unchanged. If the mode is InvertRgba, the alpha bits are also inverted.

Inverting an 8-bit image means to replace all pixels using color index i with a pixel using color index 255 minus i. The same is the case for a 1-bit image. Note that the color table is not changed.

If the image has a premultiplied alpha channel, the image is first converted to ARGB32 to be inverted and then converted back.

See also Image Transformations.
*/
func (this *QImage) InvertPixels__() {
	// arg: 0, QImage::InvertMode=Enum, QImage::InvertMode=Enum, , Invalid
	arg0 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12invertPixelsENS_10InvertModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:296
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(QIODevice *, const char *)

/*
Loads an image from the file with the given fileName. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), it is auto-detected based on the file's suffix and header. For details, see {QImageReader::setAutoDetectImageFormat()}{QImageReader}.

The file name can either refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

See also Reading and Writing Image Files.
*/
func (this *QImage) Load(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:297
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)

/*
Loads an image from the file with the given fileName. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), it is auto-detected based on the file's suffix and header. For details, see {QImageReader::setAutoDetectImageFormat()}{QImageReader}.

The file name can either refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

See also Reading and Writing Image Files.
*/
func (this *QImage) Load_1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:297
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)

/*
Loads an image from the file with the given fileName. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), it is auto-detected based on the file's suffix and header. For details, see {QImageReader::setAutoDetectImageFormat()}{QImageReader}.

The file name can either refer to an actual file on disk or to one of the application's embedded resources. See the Resource System overview for details on how to embed images and other resource files in the application's executable.

See also Reading and Writing Image Files.
*/
func (this *QImage) Load_1_(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:298
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, int, const char *)

/*
Loads an image from the first len bytes of the given binary data. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), the loader probes the file for a header to guess the file format.

See also Reading and Writing Image Files.
*/
func (this *QImage) LoadFromData(buf unsafe.Pointer /*666*/, len_ int, format string) bool {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:298
// index:0
// Public Visibility=Default Availability=Available
// [1] bool loadFromData(const uchar *, int, const char *)

/*
Loads an image from the first len bytes of the given binary data. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), the loader probes the file for a header to guess the file format.

See also Reading and Writing Image Files.
*/
func (this *QImage) LoadFromData__(buf unsafe.Pointer /*666*/, len_ int) bool {
	// arg: 2, const char *=Pointer, =Invalid, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buf, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:299
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *)

/*
Loads an image from the first len bytes of the given binary data. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), the loader probes the file for a header to guess the file format.

See also Reading and Writing Image Files.
*/
func (this *QImage) LoadFromData_1(data qtcore.QByteArray_ITF, aformat string) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(aformat)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:299
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool loadFromData(const QByteArray &, const char *)

/*
Loads an image from the first len bytes of the given binary data. Returns true if the image was successfully loaded; otherwise invalidates the image and returns false.

The loader attempts to read the image using the specified format, e.g., PNG or JPG. If format is not specified (which is the default), the loader probes the file for a header to guess the file format.

See also Reading and Writing Image Files.
*/
func (this *QImage) LoadFromData_1_(data qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage12loadFromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const

/*
Saves the image to the file with the given fileName, using the given image file format and quality factor. If format is 0, QImage will attempt to guess the format by looking at fileName's suffix.

The quality factor must be in the range 0 to 100 or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 (the default) to use the default settings.

Returns true if the image was successfully saved; otherwise returns false.

See also Reading and Writing Image Files.
*/
func (this *QImage) Save(fileName string, format string, quality int) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const

/*
Saves the image to the file with the given fileName, using the given image file format and quality factor. If format is 0, QImage will attempt to guess the format by looking at fileName's suffix.

The quality factor must be in the range 0 to 100 or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 (the default) to use the default settings.

Returns true if the image was successfully saved; otherwise returns false.

See also Reading and Writing Image Files.
*/
func (this *QImage) Save__(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *, int) const

/*
Saves the image to the file with the given fileName, using the given image file format and quality factor. If format is 0, QImage will attempt to guess the format by looking at fileName's suffix.

The quality factor must be in the range 0 to 100 or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 (the default) to use the default settings.

Returns true if the image was successfully saved; otherwise returns false.

See also Reading and Writing Image Files.
*/
func (this *QImage) Save__1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveERK7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const

/*
Saves the image to the file with the given fileName, using the given image file format and quality factor. If format is 0, QImage will attempt to guess the format by looking at fileName's suffix.

The quality factor must be in the range 0 to 100 or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 (the default) to use the default settings.

Returns true if the image was successfully saved; otherwise returns false.

See also Reading and Writing Image Files.
*/
func (this *QImage) Save_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string, quality int) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const

/*
Saves the image to the file with the given fileName, using the given image file format and quality factor. If format is 0, QImage will attempt to guess the format by looking at fileName's suffix.

The quality factor must be in the range 0 to 100 or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 (the default) to use the default settings.

Returns true if the image was successfully saved; otherwise returns false.

See also Reading and Writing Image Files.
*/
func (this *QImage) Save_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:303
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *, int) const

/*
Saves the image to the file with the given fileName, using the given image file format and quality factor. If format is 0, QImage will attempt to guess the format by looking at fileName's suffix.

The quality factor must be in the range 0 to 100 or -1. Specify 0 to obtain small compressed files, 100 for large uncompressed files, and -1 (the default) to use the default settings.

Returns true if the image was successfully saved; otherwise returns false.

See also Reading and Writing Image Files.
*/
func (this *QImage) Save_1_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, int=Int, =Invalid, , Invalid
	quality := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4saveEP9QIODevicePKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, quality)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:305
// index:0
// Public static Visibility=Default Availability=Available
// [32] QImage fromData(const uchar *, int, const char *)

/*
Constructs a QImage from the first size bytes of the given binary data. The loader attempts to read the image using the specified format. If format is not specified (which is the default), the loader probes the data for a header to guess the file format.

If format is specified, it must be one of the values returned by QImageReader::supportedImageFormats().

If the loading of the image fails, the image returned will be a null image.

See also load(), save(), and Reading and Writing Image Files.
*/
func (this *QImage) FromData(data unsafe.Pointer /*666*/, size int, format string) *QImage /*123*/ {
	var convArg2 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, data, size, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}
func QImage_FromData(data unsafe.Pointer /*666*/, size int, format string) *QImage /*123*/ {
	var nilthis *QImage
	rv := nilthis.FromData(data, size, format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:305
// index:0
// Public static Visibility=Default Availability=Available
// [32] QImage fromData(const uchar *, int, const char *)

/*
Constructs a QImage from the first size bytes of the given binary data. The loader attempts to read the image using the specified format. If format is not specified (which is the default), the loader probes the data for a header to guess the file format.

If format is specified, it must be one of the values returned by QImageReader::supportedImageFormats().

If the loading of the image fails, the image returned will be a null image.

See also load(), save(), and Reading and Writing Image Files.
*/
func (this *QImage) FromData__(data unsafe.Pointer /*666*/, size int) *QImage /*123*/ {
	// arg: 2, const char *=Pointer, =Invalid, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataEPKhiPKc", qtrt.FFI_TYPE_POINTER, data, size, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:306
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QImage fromData(const QByteArray &, const char *)

/*
Constructs a QImage from the first size bytes of the given binary data. The loader attempts to read the image using the specified format. If format is not specified (which is the default), the loader probes the data for a header to guess the file format.

If format is specified, it must be one of the values returned by QImageReader::supportedImageFormats().

If the loading of the image fails, the image returned will be a null image.

See also load(), save(), and Reading and Writing Image Files.
*/
func (this *QImage) FromData_1(data qtcore.QByteArray_ITF, format string) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}
func QImage_FromData_1(data qtcore.QByteArray_ITF, format string) *QImage /*123*/ {
	var nilthis *QImage
	rv := nilthis.FromData_1(data, format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:306
// index:1
// Public static inline Visibility=Default Availability=Available
// [32] QImage fromData(const QByteArray &, const char *)

/*
Constructs a QImage from the first size bytes of the given binary data. The loader attempts to read the image using the specified format. If format is not specified (which is the default), the loader probes the data for a header to guess the file format.

If format is specified, it must be one of the values returned by QImageReader::supportedImageFormats().

If the loading of the image fails, the image returned will be a null image.

See also load(), save(), and Reading and Writing Image Files.
*/
func (this *QImage) FromData_1_(data qtcore.QByteArray_ITF) *QImage /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage8fromDataERK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:312
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey() const

/*
Returns a number that identifies the contents of this QImage object. Distinct QImage objects can only have the same key if they refer to the same contents.

The key will change when the image is altered.
*/
func (this *QImage) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qimage.h:314
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*

 */
func (this *QImage) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimage.h:317
// index:0
// Public Visibility=Default Availability=Available
// [4] int dotsPerMeterX() const

/*
Returns the number of pixels that fit horizontally in a physical meter. Together with dotsPerMeterY(), this number defines the intended scale and aspect ratio of the image.

See also setDotsPerMeterX() and Image Information.
*/
func (this *QImage) DotsPerMeterX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13dotsPerMeterXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:318
// index:0
// Public Visibility=Default Availability=Available
// [4] int dotsPerMeterY() const

/*
Returns the number of pixels that fit vertically in a physical meter. Together with dotsPerMeterX(), this number defines the intended scale and aspect ratio of the image.

See also setDotsPerMeterY() and Image Information.
*/
func (this *QImage) DotsPerMeterY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage13dotsPerMeterYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDotsPerMeterX(int)

/*
Sets the number of pixels that fit horizontally in a physical meter, to x.

Together with dotsPerMeterY(), this number defines the intended scale and aspect ratio of the image, and determines the scale at which QPainter will draw graphics on the image. It does not change the scale or aspect ratio of the image when it is rendered on other paint devices.

See also dotsPerMeterX() and Image Information.
*/
func (this *QImage) SetDotsPerMeterX(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage16setDotsPerMeterXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDotsPerMeterY(int)

/*
Sets the number of pixels that fit vertically in a physical meter, to y.

Together with dotsPerMeterX(), this number defines the intended scale and aspect ratio of the image, and determines the scale at which QPainter will draw graphics on the image. It does not change the scale or aspect ratio of the image when it is rendered on other paint devices.

See also dotsPerMeterY() and Image Information.
*/
func (this *QImage) SetDotsPerMeterY(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage16setDotsPerMeterYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:321
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint offset() const

/*
Returns the number of pixels by which the image is intended to be offset by when positioning relative to other images.

See also setOffset() and Image Information.
*/
func (this *QImage) Offset() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6offsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffset(const QPoint &)

/*
Sets the number of pixels by which the image is intended to be offset by when positioning relative to other images, to offset.

See also offset() and Image Information.
*/
func (this *QImage) SetOffset(arg0 qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage9setOffsetERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:324
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList textKeys() const

/*
Returns the text keys for this image.

You can use these keys with text() to list the image text for a certain key.

See also text().
*/
func (this *QImage) TextKeys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage8textKeysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(const QString &) const

/*
Returns the image text associated with the given key. If the specified key is an empty string, the whole image text is returned, with each key-text pair separated by a newline.

See also setText() and textKeys().
*/
func (this *QImage) Text(key string) string {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4textERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimage.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(const QString &) const

/*
Returns the image text associated with the given key. If the specified key is an empty string, the whole image text is returned, with each key-text pair separated by a newline.

See also setText() and textKeys().
*/
func (this *QImage) Text__() string {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage4textERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimage.h:326
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &, const QString &)

/*
Sets the image text to the given text and associate it with the given key.

If you just want to store a single text block (i.e., a "comment" or just a description), you can either pass an empty key, or use a generic key like "Description".

The image text is embedded into the image data when you call save() or QImageWriter::write().

Not all image formats support embedded text. You can find out if a specific image or format supports embedding text by using QImageWriter::supportsOption(). We give an example:


      QImageWriter writer;
      writer.setFormat("png");
      if (writer.supportsOption(QImageIOHandler::Description))
          qDebug() << "Png supports embedded text";



You can use QImageWriter::supportedImageFormats() to find out which image formats are available to you.

See also text() and textKeys().
*/
func (this *QImage) SetText(key string, value string) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage7setTextERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:328
// index:0
// Public Visibility=Default Availability=Available
// [8] QPixelFormat pixelFormat() const

/*
Returns the QImage::Format as a QPixelFormat
*/
func (this *QImage) PixelFormat() *QPixelFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage11pixelFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixelFormat)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:329
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPixelFormat toPixelFormat(QImage::Format)

/*
Converts format into a QPixelFormat
*/
func (this *QImage) ToPixelFormat(format int) *QPixelFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13toPixelFormatENS_6FormatE", qtrt.FFI_TYPE_POINTER, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixelFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixelFormat)
	return rv2
}
func QImage_ToPixelFormat(format int) *QPixelFormat /*123*/ {
	var nilthis *QImage
	rv := nilthis.ToPixelFormat(format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:330
// index:0
// Public static Visibility=Default Availability=Available
// [4] QImage::Format toImageFormat(QPixelFormat)

/*
Converts format into a QImage::Format
*/
func (this *QImage) ToImageFormat(format QPixelFormat_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if format != nil && format.QPixelFormat_PTR() != nil {
		convArg0 = format.QPixelFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage13toImageFormatE12QPixelFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QImage_ToImageFormat(format QPixelFormat_ITF /*123*/) int {
	var nilthis *QImage
	rv := nilthis.ToImageFormat(format)
	return rv
}

// /usr/include/qt/QtGui/qimage.h:352
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*

 */
func (this *QImage) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimage.h:353
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage mirrored_helper(bool, bool) const

/*

 */
func (this *QImage) Mirrored_helper(horizontal bool, vertical bool) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage15mirrored_helperEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:354
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage rgbSwapped_helper() const

/*

 */
func (this *QImage) RgbSwapped_helper() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage17rgbSwapped_helperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:355
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void mirrored_inplace(bool, bool)

/*

 */
func (this *QImage) Mirrored_inplace(horizontal bool, vertical bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage16mirrored_inplaceEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:356
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rgbSwapped_inplace()

/*

 */
func (this *QImage) RgbSwapped_inplace() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage18rgbSwapped_inplaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:357
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage convertToFormat_helper(QImage::Format, Qt::ImageConversionFlags) const

/*

 */
func (this *QImage) ConvertToFormat_helper(format int, flags int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage22convertToFormat_helperENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qimage.h:358
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool convertToFormat_inplace(QImage::Format, Qt::ImageConversionFlags)

/*

 */
func (this *QImage) ConvertToFormat_inplace(format int, flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QImage23convertToFormat_inplaceENS_6FormatE6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimage.h:359
// index:0
// Protected Visibility=Default Availability=Available
// [32] QImage smoothScaled(int, int) const

/*
Returns a smoothly scaled copy of the image. The returned image has a size of width w by height h pixels.
*/
func (this *QImage) SmoothScaled(w int, h int) *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QImage12smoothScaledEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

/*
This enum type is used to describe how pixel values should be inverted in the invertPixels() function.



See also invertPixels().

*/
type QImage__InvertMode = int

// Invert only the RGB values and leave the alpha channel unchanged.
const QImage__InvertRgb QImage__InvertMode = 0

// Invert all channels, including the alpha channel.
const QImage__InvertRgba QImage__InvertMode = 1

func (this *QImage) InvertModeItemName(val int) string {
	switch val {
	case QImage__InvertRgb: // 0
		return "InvertRgb"
	case QImage__InvertRgba: // 1
		return "InvertRgba"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImage_InvertModeItemName(val int) string {
	var nilthis *QImage
	return nilthis.InvertModeItemName(val)
}

/*
The following image formats are available in Qt. Values from Format_ARGB8565_Premultiplied to Format_ARGB4444_Premultiplied were added in Qt 4.4. Values Format_RGBX8888, Format_RGBA8888 and Format_RGBA8888_Premultiplied were added in Qt 5.2. Values Format_BGR30, Format_A2BGR30_Premultiplied, Format_RGB30, Format_A2RGB30_Premultiplied were added in Qt 5.4. Format_Alpha8 and Format_Grayscale8 were added in Qt 5.5. See the notes after the table.





Note: Formats with more than 8 bit per color channel will only be processed by the raster engine using 8 bit per color.

See also format() and convertToFormat().

*/
type QImage__Format = int

// The image is invalid.
const QImage__Format_Invalid QImage__Format = 0

//
const QImage__Format_Mono QImage__Format = 1

//
const QImage__Format_MonoLSB QImage__Format = 2

//
const QImage__Format_Indexed8 QImage__Format = 3

//
const QImage__Format_RGB32 QImage__Format = 4

//
const QImage__Format_ARGB32 QImage__Format = 5

//
const QImage__Format_ARGB32_Premultiplied QImage__Format = 6

//
const QImage__Format_RGB16 QImage__Format = 7

//
const QImage__Format_ARGB8565_Premultiplied QImage__Format = 8

//
const QImage__Format_RGB666 QImage__Format = 9

//
const QImage__Format_ARGB6666_Premultiplied QImage__Format = 10

//
const QImage__Format_RGB555 QImage__Format = 11

//
const QImage__Format_ARGB8555_Premultiplied QImage__Format = 12

//
const QImage__Format_RGB888 QImage__Format = 13

//
const QImage__Format_RGB444 QImage__Format = 14

//
const QImage__Format_ARGB4444_Premultiplied QImage__Format = 15

//
const QImage__Format_RGBX8888 QImage__Format = 16

//
const QImage__Format_RGBA8888 QImage__Format = 17

//
const QImage__Format_RGBA8888_Premultiplied QImage__Format = 18

//
const QImage__Format_BGR30 QImage__Format = 19

//
const QImage__Format_A2BGR30_Premultiplied QImage__Format = 20

//
const QImage__Format_RGB30 QImage__Format = 21

//
const QImage__Format_A2RGB30_Premultiplied QImage__Format = 22

//
const QImage__Format_Alpha8 QImage__Format = 23

//
const QImage__Format_Grayscale8 QImage__Format = 24

//
const QImage__NImageFormats QImage__Format = 25

func (this *QImage) FormatItemName(val int) string {
	switch val {
	case QImage__Format_Invalid: // 0
		return "Format_Invalid"
	case QImage__Format_Mono: // 1
		return "Format_Mono"
	case QImage__Format_MonoLSB: // 2
		return "Format_MonoLSB"
	case QImage__Format_Indexed8: // 3
		return "Format_Indexed8"
	case QImage__Format_RGB32: // 4
		return "Format_RGB32"
	case QImage__Format_ARGB32: // 5
		return "Format_ARGB32"
	case QImage__Format_ARGB32_Premultiplied: // 6
		return "Format_ARGB32_Premultiplied"
	case QImage__Format_RGB16: // 7
		return "Format_RGB16"
	case QImage__Format_ARGB8565_Premultiplied: // 8
		return "Format_ARGB8565_Premultiplied"
	case QImage__Format_RGB666: // 9
		return "Format_RGB666"
	case QImage__Format_ARGB6666_Premultiplied: // 10
		return "Format_ARGB6666_Premultiplied"
	case QImage__Format_RGB555: // 11
		return "Format_RGB555"
	case QImage__Format_ARGB8555_Premultiplied: // 12
		return "Format_ARGB8555_Premultiplied"
	case QImage__Format_RGB888: // 13
		return "Format_RGB888"
	case QImage__Format_RGB444: // 14
		return "Format_RGB444"
	case QImage__Format_ARGB4444_Premultiplied: // 15
		return "Format_ARGB4444_Premultiplied"
	case QImage__Format_RGBX8888: // 16
		return "Format_RGBX8888"
	case QImage__Format_RGBA8888: // 17
		return "Format_RGBA8888"
	case QImage__Format_RGBA8888_Premultiplied: // 18
		return "Format_RGBA8888_Premultiplied"
	case QImage__Format_BGR30: // 19
		return "Format_BGR30"
	case QImage__Format_A2BGR30_Premultiplied: // 20
		return "Format_A2BGR30_Premultiplied"
	case QImage__Format_RGB30: // 21
		return "Format_RGB30"
	case QImage__Format_A2RGB30_Premultiplied: // 22
		return "Format_A2RGB30_Premultiplied"
	case QImage__Format_Alpha8: // 23
		return "Format_Alpha8"
	case QImage__Format_Grayscale8: // 24
		return "Format_Grayscale8"
	case QImage__NImageFormats: // 25
		return "NImageFormats"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QImage_FormatItemName(val int) string {
	var nilthis *QImage
	return nilthis.FormatItemName(val)
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
