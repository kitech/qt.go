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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qmovie.h:65
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMovie) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:82
// index:0
// void QMovie(class QObject *)
func NewQMovie(parent unsafe.Pointer) *QMovie {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QMovie{cthis}
}

// /usr/include/qt/QtGui/qmovie.h:83
// index:1
// void QMovie(class QIODevice *, const class QByteArray &, class QObject *)
func NewQMovie_1(device unsafe.Pointer, format unsafe.Pointer, parent unsafe.Pointer) *QMovie {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, device, format, parent)
	gopp.ErrPrint(err, rv)
	return &QMovie{cthis}
}

// /usr/include/qt/QtGui/qmovie.h:84
// index:2
// void QMovie(const class QString &, const class QByteArray &, class QObject *)
func NewQMovie_2(fileName unsafe.Pointer, format unsafe.Pointer, parent unsafe.Pointer) *QMovie {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, fileName, format, parent)
	gopp.ErrPrint(err, rv)
	return &QMovie{cthis}
}

// /usr/include/qt/QtGui/qmovie.h:85
// index:0
// virtual
// void ~QMovie()
func DeleteQMovie(*QMovie) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovieD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:87
// index:0
// static
// QList<QByteArray> supportedFormats()
func (this *QMovie) SupportedFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie16supportedFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMovie_SupportedFormats() {
	// 0: (), ()
	var nilthis *QMovie
	nilthis.SupportedFormats()
}

// /usr/include/qt/QtGui/qmovie.h:89
// index:0
// void setDevice(class QIODevice *)
func (this *QMovie) SetDevice(device unsafe.Pointer) {
	// 0: (, QIODevice * device), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.cthis, device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:90
// index:0
// QIODevice * device()
func (this *QMovie) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie6deviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:92
// index:0
// void setFileName(const class QString &)
func (this *QMovie) SetFileName(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:93
// index:0
// QString fileName()
func (this *QMovie) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:95
// index:0
// void setFormat(const class QByteArray &)
func (this *QMovie) SetFormat(format unsafe.Pointer) {
	// 0: (, const QByteArray & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie9setFormatERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:96
// index:0
// QByteArray format()
func (this *QMovie) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:98
// index:0
// void setBackgroundColor(const class QColor &)
func (this *QMovie) SetBackgroundColor(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:99
// index:0
// QColor backgroundColor()
func (this *QMovie) BackgroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie15backgroundColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:101
// index:0
// QMovie::MovieState state()
func (this *QMovie) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie5stateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:103
// index:0
// QRect frameRect()
func (this *QMovie) FrameRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9frameRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:104
// index:0
// QImage currentImage()
func (this *QMovie) CurrentImage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie12currentImageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:105
// index:0
// QPixmap currentPixmap()
func (this *QMovie) CurrentPixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie13currentPixmapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:107
// index:0
// bool isValid()
func (this *QMovie) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:108
// index:0
// QImageReader::ImageReaderError lastError()
func (this *QMovie) LastError() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9lastErrorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:109
// index:0
// QString lastErrorString()
func (this *QMovie) LastErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie15lastErrorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:111
// index:0
// bool jumpToFrame(int)
func (this *QMovie) JumpToFrame(frameNumber int) {
	// 0: (, int frameNumber), (&frameNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie11jumpToFrameEi", ffiqt.FFI_TYPE_VOID, this.cthis, &frameNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:112
// index:0
// int loopCount()
func (this *QMovie) LoopCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9loopCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:113
// index:0
// int frameCount()
func (this *QMovie) FrameCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie10frameCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:114
// index:0
// int nextFrameDelay()
func (this *QMovie) NextFrameDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie14nextFrameDelayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:115
// index:0
// int currentFrameNumber()
func (this *QMovie) CurrentFrameNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie18currentFrameNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:117
// index:0
// int speed()
func (this *QMovie) Speed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie5speedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:119
// index:0
// QSize scaledSize()
func (this *QMovie) ScaledSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie10scaledSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:120
// index:0
// void setScaledSize(const class QSize &)
func (this *QMovie) SetScaledSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie13setScaledSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:122
// index:0
// QMovie::CacheMode cacheMode()
func (this *QMovie) CacheMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMovie9cacheModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:123
// index:0
// void setCacheMode(enum QMovie::CacheMode)
func (this *QMovie) SetCacheMode(mode int) {
	// 0: (, QMovie::CacheMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie12setCacheModeENS_9CacheModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:126
// index:0
// void started()
func (this *QMovie) Started() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie7startedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:127
// index:0
// void resized(const class QSize &)
func (this *QMovie) Resized(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie7resizedERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:128
// index:0
// void updated(const class QRect &)
func (this *QMovie) Updated(rect unsafe.Pointer) {
	// 0: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie7updatedERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:129
// index:0
// void stateChanged(class QMovie::MovieState)
func (this *QMovie) StateChanged(state int) {
	// 0: (, QMovie::MovieState state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie12stateChangedENS_10MovieStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:130
// index:0
// void error(class QImageReader::ImageReaderError)
func (this *QMovie) Error(error int) {
	// 0: (, QImageReader::ImageReaderError error), (&error)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie5errorEN12QImageReader16ImageReaderErrorE", ffiqt.FFI_TYPE_VOID, this.cthis, &error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:131
// index:0
// void finished()
func (this *QMovie) Finished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie8finishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:132
// index:0
// void frameChanged(int)
func (this *QMovie) FrameChanged(frameNumber int) {
	// 0: (, int frameNumber), (&frameNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie12frameChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &frameNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:135
// index:0
// void start()
func (this *QMovie) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie5startEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:136
// index:0
// bool jumpToNextFrame()
func (this *QMovie) JumpToNextFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie15jumpToNextFrameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:137
// index:0
// void setPaused(_Bool)
func (this *QMovie) SetPaused(paused bool) {
	// 0: (, bool paused), (&paused)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie9setPausedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:138
// index:0
// void stop()
func (this *QMovie) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie4stopEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qmovie.h:139
// index:0
// void setSpeed(int)
func (this *QMovie) SetSpeed(percentSpeed int) {
	// 0: (, int percentSpeed), (&percentSpeed)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMovie8setSpeedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &percentSpeed)
	gopp.ErrPrint(err, rv)
}

//  body block end
