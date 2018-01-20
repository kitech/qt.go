//  header block begin
// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QImageIOHandler struct {
	*qtrt.CObject
}

func (this *QImageIOHandler) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qimageiohandler.h:62
// index:0
// void QImageIOHandler()
func NewQImageIOHandler() *QImageIOHandler {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandlerC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageIOHandlerFromPointer(cthis)
	return gothis
}
func NewQImageIOHandlerFromPointer(cthis unsafe.Pointer) *QImageIOHandler {
	return &QImageIOHandler{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qimageiohandler.h:129
// index:1
// void QImageIOHandler(class QImageIOHandlerPrivate &)
func NewQImageIOHandler_1(dd unsafe.Pointer) *QImageIOHandler {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandlerC1ER22QImageIOHandlerPrivate", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageIOHandlerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:63
// index:0
// virtual
// void ~QImageIOHandler()
func DeleteQImageIOHandler(*QImageIOHandler) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandlerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:65
// index:0
// void setDevice(class QIODevice *)
func (this *QImageIOHandler) SetDevice(device unsafe.Pointer) {
	// 0: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:66
// index:0
// QIODevice * device()
func (this *QImageIOHandler) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:68
// index:0
// void setFormat(const class QByteArray &)
func (this *QImageIOHandler) SetFormat(format unsafe.Pointer) {
	// 0: (, format const QByteArray &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler9setFormatERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:69
// index:1
// void setFormat(const class QByteArray &)
func (this *QImageIOHandler) SetFormat_1(format unsafe.Pointer) {
	// 1: (, format const QByteArray &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler9setFormatERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:70
// index:0
// QByteArray format()
func (this *QImageIOHandler) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler6formatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:72
// index:0
// virtual
// QByteArray name()
func (this *QImageIOHandler) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:74
// index:0
// pure virtual
// bool canRead()
func (this *QImageIOHandler) CanRead() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler7canReadEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:75
// index:0
// pure virtual
// bool read(class QImage *)
func (this *QImageIOHandler) Read(image unsafe.Pointer) {
	// 0: (, image QImage *), (image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler4readEP6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:76
// index:0
// virtual
// bool write(const class QImage &)
func (this *QImageIOHandler) Write(image unsafe.Pointer) {
	// 0: (, image const QImage &), (image)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler5writeERK6QImage", ffiqt.FFI_TYPE_VOID, this.GetCthis(), image)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:115
// index:0
// virtual
// QVariant option(enum QImageIOHandler::ImageOption)
func (this *QImageIOHandler) Option(option int) {
	// 0: (, option QImageIOHandler::ImageOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler6optionENS_11ImageOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:116
// index:0
// virtual
// void setOption(enum QImageIOHandler::ImageOption, const class QVariant &)
func (this *QImageIOHandler) SetOption(option int, value unsafe.Pointer) {
	// 0: (, option QImageIOHandler::ImageOption, value const QVariant &), (&option, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler9setOptionENS_11ImageOptionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:117
// index:0
// virtual
// bool supportsOption(enum QImageIOHandler::ImageOption)
func (this *QImageIOHandler) SupportsOption(option int) {
	// 0: (, option QImageIOHandler::ImageOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler14supportsOptionENS_11ImageOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:120
// index:0
// virtual
// bool jumpToNextImage()
func (this *QImageIOHandler) JumpToNextImage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler15jumpToNextImageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:121
// index:0
// virtual
// bool jumpToImage(int)
func (this *QImageIOHandler) JumpToImage(imageNumber int) {
	// 0: (, imageNumber int), (&imageNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler11jumpToImageEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &imageNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:122
// index:0
// virtual
// int loopCount()
func (this *QImageIOHandler) LoopCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler9loopCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:123
// index:0
// virtual
// int imageCount()
func (this *QImageIOHandler) ImageCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler10imageCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:124
// index:0
// virtual
// int nextImageDelay()
func (this *QImageIOHandler) NextImageDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler14nextImageDelayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:125
// index:0
// virtual
// int currentImageNumber()
func (this *QImageIOHandler) CurrentImageNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler18currentImageNumberEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:126
// index:0
// virtual
// QRect currentImageRect()
func (this *QImageIOHandler) CurrentImageRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler16currentImageRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
