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
// extern C begin: 52
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

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
// [8] const QMetaObject * metaObject()
func (this *QMovie) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qmovie.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QObject *)
func NewQMovie(parent *qtcore.QObject /*777 QObject **/) *QMovie {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMovie(QIODevice *, const QByteArray &, QObject *)
func NewQMovie_1(device *qtcore.QIODevice /*777 QIODevice **/, format *qtcore.QByteArray, parent *qtcore.QObject /*777 QObject **/) *QMovie {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMovie(const QString &, const QByteArray &, QObject *)
func NewQMovie_2(fileName string, format *qtcore.QByteArray, parent *qtcore.QObject /*777 QObject **/) *QMovie {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = format.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMovie()
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
func (this *QMovie) SetDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QMovie) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qmovie.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QMovie) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
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
func (this *QMovie) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format()
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
func (this *QMovie) SetBackgroundColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie18setBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor backgroundColor()
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
// [4] QMovie::MovieState state()
func (this *QMovie) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qmovie.h:103
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameRect()
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
// [32] QImage currentImage()
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
// [32] QPixmap currentPixmap()
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
// [1] bool isValid()
func (this *QMovie) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmovie.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] QImageReader::ImageReaderError lastError()
func (this *QMovie) LastError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9lastErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qmovie.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString lastErrorString()
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
func (this *QMovie) JumpToFrame(frameNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie11jumpToFrameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frameNumber)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmovie.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount()
func (this *QMovie) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameCount()
func (this *QMovie) FrameCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie10frameCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextFrameDelay()
func (this *QMovie) NextFrameDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie14nextFrameDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentFrameNumber()
func (this *QMovie) CurrentFrameNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie18currentFrameNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int speed()
func (this *QMovie) Speed() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie5speedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qmovie.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize scaledSize()
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
func (this *QMovie) SetScaledSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie13setScaledSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] QMovie::CacheMode cacheMode()
func (this *QMovie) CacheMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMovie9cacheModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qmovie.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(enum QMovie::CacheMode)
func (this *QMovie) SetCacheMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie12setCacheModeENS_9CacheModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void started()
func (this *QMovie) Started() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie7startedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resized(const QSize &)
func (this *QMovie) Resized(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie7resizedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updated(const QRect &)
func (this *QMovie) Updated(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie7updatedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QMovie::MovieState)
func (this *QMovie) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie12stateChangedENS_10MovieStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(QImageReader::ImageReaderError)
func (this *QMovie) Error(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie5errorEN12QImageReader16ImageReaderErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()
func (this *QMovie) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frameChanged(int)
func (this *QMovie) FrameChanged(frameNumber int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie12frameChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frameNumber)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()
func (this *QMovie) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToNextFrame()
func (this *QMovie) JumpToNextFrame() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie15jumpToNextFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qmovie.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QMovie) SetPaused(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QMovie) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpeed(int)
func (this *QMovie) SetSpeed(percentSpeed int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMovie8setSpeedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), percentSpeed)
	qtrt.ErrPrint(err, rv)
}

type QMovie__MovieState = int

const QMovie__NotRunning QMovie__MovieState = 0
const QMovie__Paused QMovie__MovieState = 1
const QMovie__Running QMovie__MovieState = 2

type QMovie__CacheMode = int

const QMovie__CacheNone QMovie__CacheMode = 0
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
