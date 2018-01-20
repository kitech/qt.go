//  header block begin
// /usr/include/qt/QtGui/qmovie.h
// #include <qmovie.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 53
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
type QMovie struct {
	*qtcore.QObject
}

func (this *QMovie) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQMovieFromPointer(cthis unsafe.Pointer) *QMovie {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMovie{bcthis0}
}

// /usr/include/qt/QtGui/qmovie.h:65
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMovie) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:82
// index:0
// Public
// void QMovie(class QObject *)
func NewQMovie(parent unsafe.Pointer) *QMovie {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:83
// index:1
// Public
// void QMovie(class QIODevice *, const class QByteArray &, class QObject *)
func NewQMovie_1(device unsafe.Pointer, format *qtcore.QByteArray, parent unsafe.Pointer) *QMovie {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, device, convArg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:84
// index:2
// Public
// void QMovie(const class QString &, const class QByteArray &, class QObject *)
func NewQMovie_2(fileName *qtcore.QString, format *qtcore.QByteArray, parent unsafe.Pointer) *QMovie {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = fileName.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQMovieFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qmovie.h:85
// index:0
// Public virtual
// void ~QMovie()
func DeleteQMovie(*QMovie) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:87
// index:0
// Public static
// QList<QByteArray> supportedFormats()
func (this *QMovie) SupportedFormats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie16supportedFormatsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QMovie_SupportedFormats() {
	var nilthis *QMovie
	nilthis.SupportedFormats()
}

// /usr/include/qt/QtGui/qmovie.h:89
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QMovie) SetDevice(device unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:90
// index:0
// Public
// QIODevice * device()
func (this *QMovie) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:92
// index:0
// Public
// void setFileName(const class QString &)
func (this *QMovie) SetFileName(fileName *qtcore.QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:93
// index:0
// Public
// QString fileName()
func (this *QMovie) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:95
// index:0
// Public
// void setFormat(const class QByteArray &)
func (this *QMovie) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:96
// index:0
// Public
// QByteArray format()
func (this *QMovie) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:98
// index:0
// Public
// void setBackgroundColor(const class QColor &)
func (this *QMovie) SetBackgroundColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:99
// index:0
// Public
// QColor backgroundColor()
func (this *QMovie) BackgroundColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:101
// index:0
// Public
// QMovie::MovieState state()
func (this *QMovie) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:103
// index:0
// Public
// QRect frameRect()
func (this *QMovie) FrameRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9frameRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:104
// index:0
// Public
// QImage currentImage()
func (this *QMovie) CurrentImage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie12currentImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:105
// index:0
// Public
// QPixmap currentPixmap()
func (this *QMovie) CurrentPixmap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie13currentPixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:107
// index:0
// Public
// bool isValid()
func (this *QMovie) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:108
// index:0
// Public
// QImageReader::ImageReaderError lastError()
func (this *QMovie) LastError() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9lastErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:109
// index:0
// Public
// QString lastErrorString()
func (this *QMovie) LastErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie15lastErrorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:111
// index:0
// Public
// bool jumpToFrame(int)
func (this *QMovie) JumpToFrame(frameNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie11jumpToFrameEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &frameNumber)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:112
// index:0
// Public
// int loopCount()
func (this *QMovie) LoopCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9loopCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:113
// index:0
// Public
// int frameCount()
func (this *QMovie) FrameCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie10frameCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:114
// index:0
// Public
// int nextFrameDelay()
func (this *QMovie) NextFrameDelay() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie14nextFrameDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:115
// index:0
// Public
// int currentFrameNumber()
func (this *QMovie) CurrentFrameNumber() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie18currentFrameNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:117
// index:0
// Public
// int speed()
func (this *QMovie) Speed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie5speedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:119
// index:0
// Public
// QSize scaledSize()
func (this *QMovie) ScaledSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie10scaledSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:120
// index:0
// Public
// void setScaledSize(const class QSize &)
func (this *QMovie) SetScaledSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie13setScaledSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:122
// index:0
// Public
// QMovie::CacheMode cacheMode()
func (this *QMovie) CacheMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9cacheModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:123
// index:0
// Public
// void setCacheMode(enum QMovie::CacheMode)
func (this *QMovie) SetCacheMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie12setCacheModeENS_9CacheModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:126
// index:0
// Public
// void started()
func (this *QMovie) Started() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie7startedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:127
// index:0
// Public
// void resized(const class QSize &)
func (this *QMovie) Resized(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie7resizedERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:128
// index:0
// Public
// void updated(const class QRect &)
func (this *QMovie) Updated(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie7updatedERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:129
// index:0
// Public
// void stateChanged(class QMovie::MovieState)
func (this *QMovie) StateChanged(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie12stateChangedENS_10MovieStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:130
// index:0
// Public
// void error(class QImageReader::ImageReaderError)
func (this *QMovie) Error(error int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie5errorEN12QImageReader16ImageReaderErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:131
// index:0
// Public
// void finished()
func (this *QMovie) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:132
// index:0
// Public
// void frameChanged(int)
func (this *QMovie) FrameChanged(frameNumber int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie12frameChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &frameNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:135
// index:0
// Public
// void start()
func (this *QMovie) Start() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:136
// index:0
// Public
// bool jumpToNextFrame()
func (this *QMovie) JumpToNextFrame() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie15jumpToNextFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qmovie.h:137
// index:0
// Public
// void setPaused(_Bool)
func (this *QMovie) SetPaused(paused bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie9setPausedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:138
// index:0
// Public
// void stop()
func (this *QMovie) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:139
// index:0
// Public
// void setSpeed(int)
func (this *QMovie) SetSpeed(percentSpeed int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie8setSpeedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &percentSpeed)
	gopp.ErrPrint(err, rv)
}

//  body block end
