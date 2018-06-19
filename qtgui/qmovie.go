package qtgui

// /usr/include/qt/QtGui/qmovie.h
// #include <qmovie.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 61
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
type QMovie struct {
	*qtcore.QObject
}
type QMovie_ITF interface {
	qtcore.QObject_ITF
	QMovie_PTR() *QMovie
}

func (ptr *QMovie) QMovie_PTR() *QMovie { return ptr }

func (this *QMovie) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMovie) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMovieFromPointer(cthis unsafe.Pointer) *QMovie {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMovie{bcthis0}
}
func (*QMovie) NewFromPointer(cthis unsafe.Pointer) *QMovie {
	return NewQMovieFromPointer(cthis)
}

// /usr/include/qt/QtGui/qmovie.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMovie) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qmovie.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie(parent qtcore.QObject_ITF /*777 QObject **/) *QMovie {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie__() *QMovie {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QIODevice *, const QByteArray &, QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QMovie {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QIODevice *, const QByteArray &, QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QMovie {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QIODevice *, const QByteArray &, QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie_1_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QMovie {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMovie(const QString &, const QByteArray &, QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie_2(fileName string, format qtcore.QByteArray_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QMovie {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMovie(const QString &, const QByteArray &, QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie_2_(fileName string) *QMovie {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMovie(const QString &, const QByteArray &, QObject *)

/*
Constructs a QMovie object, passing the parent object to QObject's constructor.

See also setFileName(), setDevice(), and setFormat().
*/
func NewQMovie_2_1(fileName string, format qtcore.QByteArray_ITF) *QMovie {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMovie")
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMovie()

/*

 */
func DeleteQMovie(this *QMovie) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qmovie.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets the current device to device. QMovie will read image data from this device when the movie is running.

See also device() and setFormat().
*/
func (this *QMovie) SetDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the device QMovie reads image data from. If no device has currently been assigned, 0 is returned.

See also setDevice() and fileName().
*/
func (this *QMovie) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qmovie.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*
Sets the name of the file that QMovie reads image data from, to fileName.

See also fileName(), setDevice(), and setFormat().
*/
func (this *QMovie) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Returns the name of the file that QMovie reads image data from. If no file name has been assigned, or if the assigned device is not a file, an empty QString is returned.

See also setFileName() and device().
*/
func (this *QMovie) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qmovie.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)

/*
Sets the format that QMovie will use when decoding image data, to format. By default, QMovie will attempt to guess the format of the image data.

You can call supportedFormats() for the full list of formats QMovie supports.

See also format() and QImageReader::supportedImageFormats().
*/
func (this *QMovie) SetFormat(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format() const

/*
Returns the format that QMovie uses when decoding image data. If no format has been assigned, an empty QByteArray() is returned.

See also setFormat().
*/
func (this *QMovie) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qmovie.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)

/*
For image formats that support it, this function sets the background color to color.

See also backgroundColor().
*/
func (this *QMovie) SetBackgroundColor(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie18setBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor backgroundColor() const

/*
Returns the background color of the movie. If no background color has been assigned, an invalid QColor is returned.

See also setBackgroundColor().
*/
func (this *QMovie) BackgroundColor() *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qmovie.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] QMovie::MovieState state() const

/*
Returns the current state of QMovie.

See also MovieState and stateChanged().
*/
func (this *QMovie) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qmovie.h:103
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameRect() const

/*
Returns the rect of the last frame. If no frame has yet been updated, an invalid QRect is returned.

See also currentImage() and currentPixmap().
*/
func (this *QMovie) FrameRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9frameRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qmovie.h:104
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage currentImage() const

/*
Returns the current frame as a QImage.

See also currentPixmap() and updated().
*/
func (this *QMovie) CurrentImage() *QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie12currentImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qmovie.h:105
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap currentPixmap() const

/*
Returns the current frame as a QPixmap.

See also currentImage() and updated().
*/
func (this *QMovie) CurrentPixmap() *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie13currentPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qmovie.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the movie is valid (e.g., the image data is readable and the image format is supported); otherwise returns false.

For information about why the movie is not valid, see lastError().
*/
func (this *QMovie) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmovie.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] QImageReader::ImageReaderError lastError() const

/*
Returns the most recent error that occurred while attempting to read image data.

See also lastErrorString().
*/
func (this *QMovie) LastError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9lastErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qmovie.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString lastErrorString() const

/*
Returns a human-readable representation of the most recent error that occurred while attempting to read image data.

See also lastError().
*/
func (this *QMovie) LastErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie15lastErrorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qmovie.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToFrame(int)

/*
Jumps to frame number frameNumber. Returns true on success; otherwise returns false.
*/
func (this *QMovie) JumpToFrame(frameNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie11jumpToFrameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frameNumber)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmovie.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount() const

/*
Returns the number of times the movie will loop before it finishes. If the movie will only play once (no looping), loopCount returns 0. If the movie loops forever, loopCount returns -1.

Note that, if the image data comes from a sequential device (e.g. a socket), QMovie can only loop the movie if the cacheMode is set to QMovie::CacheAll.
*/
func (this *QMovie) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameCount() const

/*
Returns the number of frames in the movie.

Certain animation formats do not support this feature, in which case 0 is returned.
*/
func (this *QMovie) FrameCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie10frameCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextFrameDelay() const

/*
Returns the number of milliseconds QMovie will wait before updating the next frame in the animation.
*/
func (this *QMovie) NextFrameDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie14nextFrameDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentFrameNumber() const

/*
Returns the sequence number of the current frame. The number of the first frame in the movie is 0.
*/
func (this *QMovie) CurrentFrameNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie18currentFrameNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int speed() const

/*

 */
func (this *QMovie) Speed() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie5speedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize scaledSize()

/*
Returns the scaled size of frames.

This function was introduced in  Qt 4.1.

See also setScaledSize() and QImageReader::scaledSize().
*/
func (this *QMovie) ScaledSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie10scaledSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qmovie.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledSize(const QSize &)

/*
Sets the scaled frame size to size.

This function was introduced in  Qt 4.1.

See also scaledSize() and QImageReader::setScaledSize().
*/
func (this *QMovie) SetScaledSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie13setScaledSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] QMovie::CacheMode cacheMode() const

/*

 */
func (this *QMovie) CacheMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9cacheModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qmovie.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(QMovie::CacheMode)

/*

 */
func (this *QMovie) SetCacheMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie12setCacheModeENS_9CacheModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void started()

/*
This signal is emitted after QMovie::start() has been called, and QMovie has entered QMovie::Running state.
*/
func (this *QMovie) Started() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie7startedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resized(const QSize &)

/*
This signal is emitted when the current frame has been resized to size. This effect is sometimes used in animations as an alternative to replacing the frame. You can call currentImage() or currentPixmap() to get a copy of the updated frame.
*/
func (this *QMovie) Resized(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie7resizedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updated(const QRect &)

/*
This signal is emitted when the rect rect in the current frame has been updated. You can call currentImage() or currentPixmap() to get a copy of the updated frame.
*/
func (this *QMovie) Updated(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie7updatedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QMovie::MovieState)

/*
This signal is emitted every time the state of the movie changes. The new state is specified by state.

See also QMovie::state().
*/
func (this *QMovie) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie12stateChangedENS_10MovieStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(QImageReader::ImageReaderError)

/*
This signal is emitted by QMovie when the error error occurred during playback. QMovie will stop the movie, and enter QMovie::NotRunning state.

See also lastError() and lastErrorString().
*/
func (this *QMovie) Error(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie5errorEN12QImageReader16ImageReaderErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*
This signal is emitted when the movie has finished.

See also QMovie::stop().
*/
func (this *QMovie) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frameChanged(int)

/*
This signal is emitted when the frame number has changed to frameNumber. You can call currentImage() or currentPixmap() to get a copy of the frame.

This function was introduced in  Qt 4.1.
*/
func (this *QMovie) FrameChanged(frameNumber int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie12frameChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frameNumber)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()

/*
Starts the movie. QMovie will enter Running state, and start emitting updated() and resized() as the movie progresses.

If QMovie is in the Paused state, this function is equivalent to calling setPaused(false). If QMovie is already in the Running state, this function does nothing.

See also stop() and setPaused().
*/
func (this *QMovie) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToNextFrame()

/*
Jumps to the next frame. Returns true on success; otherwise returns false.
*/
func (this *QMovie) JumpToNextFrame() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie15jumpToNextFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmovie.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(bool)

/*
If paused is true, QMovie will enter Paused state and emit stateChanged(Paused); otherwise it will enter Running state and emit stateChanged(Running).

See also state().
*/
func (this *QMovie) SetPaused(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the movie. QMovie enters NotRunning state, and stops emitting updated() and resized(). If start() is called again, the movie will restart from the beginning.

If QMovie is already in the NotRunning state, this function does nothing.

See also start() and setPaused().
*/
func (this *QMovie) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpeed(int)

/*

 */
func (this *QMovie) SetSpeed(percentSpeed int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie8setSpeedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), percentSpeed)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the different states of QMovie.


*/
type QMovie__MovieState = int

// The movie is not running. This is QMovie's initial state, and the state it enters after stop() has been called or the movie is finished.
const QMovie__NotRunning QMovie__MovieState = 0

// The movie is paused, and QMovie stops emitting updated() or resized(). This state is entered after calling pause() or setPaused(true). The current frame number it kept, and the movie will continue with the next frame when unpause() or setPaused(false) is called.
const QMovie__Paused QMovie__MovieState = 1

// The movie is running.
const QMovie__Running QMovie__MovieState = 2

/*
This enum describes the different cache modes of QMovie.


*/
type QMovie__CacheMode = int

// No frames are cached (the default).
const QMovie__CacheNone QMovie__CacheMode = 0

// All frames are cached.
const QMovie__CacheAll QMovie__CacheMode = 1

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
