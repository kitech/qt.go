//  header block begin
// /usr/include/qt/QtGui/qpicture.h
// #include <qpicture.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QPictureIO struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpicture.h:134
// index:0
// void QPictureIO()
func NewQPictureIO() *QPictureIO {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPictureIO{cthis}
}

// /usr/include/qt/QtGui/qpicture.h:135
// index:1
// void QPictureIO(class QIODevice *, const char *)
func NewQPictureIO_1(ioDevice unsafe.Pointer, format unsafe.Pointer) *QPictureIO {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOC2EP9QIODevicePKc", ffiqt.FFI_TYPE_VOID, cthis, ioDevice, format)
	gopp.ErrPrint(err, rv)
	return &QPictureIO{cthis}
}

// /usr/include/qt/QtGui/qpicture.h:136
// index:2
// void QPictureIO(const class QString &, const char *)
func NewQPictureIO_2(fileName unsafe.Pointer, format unsafe.Pointer) *QPictureIO {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOC2ERK7QStringPKc", ffiqt.FFI_TYPE_VOID, cthis, fileName, format)
	gopp.ErrPrint(err, rv)
	return &QPictureIO{cthis}
}

// /usr/include/qt/QtGui/qpicture.h:137
// index:0
// void ~QPictureIO()
func DeleteQPictureIO(*QPictureIO) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:139
// index:0
// const QPicture & picture()
func (this *QPictureIO) Picture() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO7pictureEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:140
// index:0
// int status()
func (this *QPictureIO) Status() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO6statusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:141
// index:0
// const char * format()
func (this *QPictureIO) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:142
// index:0
// QIODevice * ioDevice()
func (this *QPictureIO) IoDevice() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO8ioDeviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:143
// index:0
// QString fileName()
func (this *QPictureIO) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:144
// index:0
// int quality()
func (this *QPictureIO) Quality() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO7qualityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:145
// index:0
// QString description()
func (this *QPictureIO) Description() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO11descriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:146
// index:0
// const char * parameters()
func (this *QPictureIO) Parameters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO10parametersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:147
// index:0
// float gamma()
func (this *QPictureIO) Gamma() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO5gammaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:149
// index:0
// void setPicture(const class QPicture &)
func (this *QPictureIO) SetPicture(arg0 unsafe.Pointer) {
	// 0: (, const QPicture & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO10setPictureERK8QPicture", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:150
// index:0
// void setStatus(int)
func (this *QPictureIO) SetStatus(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO9setStatusEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:151
// index:0
// void setFormat(const char *)
func (this *QPictureIO) SetFormat(arg0 unsafe.Pointer) {
	// 0: (, const char * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO9setFormatEPKc", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:152
// index:0
// void setIODevice(class QIODevice *)
func (this *QPictureIO) SetIODevice(arg0 unsafe.Pointer) {
	// 0: (, QIODevice * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO11setIODeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:153
// index:0
// void setFileName(const class QString &)
func (this *QPictureIO) SetFileName(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:154
// index:0
// void setQuality(int)
func (this *QPictureIO) SetQuality(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO10setQualityEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:155
// index:0
// void setDescription(const class QString &)
func (this *QPictureIO) SetDescription(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO14setDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:156
// index:0
// void setParameters(const char *)
func (this *QPictureIO) SetParameters(arg0 unsafe.Pointer) {
	// 0: (, const char * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13setParametersEPKc", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:157
// index:0
// void setGamma(float)
func (this *QPictureIO) SetGamma(arg0 float32) {
	// 0: (, float arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO8setGammaEf", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:159
// index:0
// bool read()
func (this *QPictureIO) Read() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO4readEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:160
// index:0
// bool write()
func (this *QPictureIO) Write() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO5writeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:162
// index:0
// static
// QByteArray pictureFormat(const class QString &)
func (this *QPictureIO) PictureFormat(fileName unsafe.Pointer) {
	// 0: (const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13pictureFormatERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPictureIO_PictureFormat(fileName unsafe.Pointer) {
	// 0: (const QString & fileName), (fileName)
	var nilthis *QPictureIO
	nilthis.PictureFormat(fileName)
}

// /usr/include/qt/QtGui/qpicture.h:163
// index:1
// static
// QByteArray pictureFormat(class QIODevice *)
func (this *QPictureIO) PictureFormat_1(arg0 unsafe.Pointer) {
	// 1: (QIODevice * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13pictureFormatEP9QIODevice", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPictureIO_PictureFormat_1(arg0 unsafe.Pointer) {
	// 1: (QIODevice * arg0), (arg0)
	var nilthis *QPictureIO
	nilthis.PictureFormat_1(arg0)
}

// /usr/include/qt/QtGui/qpicture.h:164
// index:0
// static
// QList<QByteArray> inputFormats()
func (this *QPictureIO) InputFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO12inputFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPictureIO_InputFormats() {
	// 0: (), ()
	var nilthis *QPictureIO
	nilthis.InputFormats()
}

// /usr/include/qt/QtGui/qpicture.h:165
// index:0
// static
// QList<QByteArray> outputFormats()
func (this *QPictureIO) OutputFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13outputFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPictureIO_OutputFormats() {
	// 0: (), ()
	var nilthis *QPictureIO
	nilthis.OutputFormats()
}

// /usr/include/qt/QtGui/qpicture.h:167
// index:0
// static
// void defineIOHandler(const char *, const char *, const char *, picture_io_handler, picture_io_handler)
func (this *QPictureIO) DefineIOHandler(format unsafe.Pointer, header unsafe.Pointer, flags unsafe.Pointer, read_picture unsafe.Pointer, write_picture unsafe.Pointer) {
	// 0: (const char * format, const char * header, const char * flags, picture_io_handler read_picture, picture_io_handler write_picture), (format, header, flags, read_picture, write_picture)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO15defineIOHandlerEPKcS1_S1_PFvPS_ES4_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPictureIO_DefineIOHandler(format unsafe.Pointer, header unsafe.Pointer, flags unsafe.Pointer, read_picture unsafe.Pointer, write_picture unsafe.Pointer) {
	// 0: (const char * format, const char * header, const char * flags, picture_io_handler read_picture, picture_io_handler write_picture), (format, header, flags, read_picture, write_picture)
	var nilthis *QPictureIO
	nilthis.DefineIOHandler(format, header, flags, read_picture, write_picture)
}

//  body block end
