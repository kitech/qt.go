//  header block begin
// /usr/include/qt/QtGui/qtextdocumentwriter.h
// #include <qtextdocumentwriter.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QTextDocumentWriter struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:57
// index:0
// void QTextDocumentWriter()
func NewQTextDocumentWriter() *QTextDocumentWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextDocumentWriter{cthis}
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:58
// index:1
// void QTextDocumentWriter(class QIODevice *, const class QByteArray &)
func NewQTextDocumentWriter_1(device unsafe.Pointer, format unsafe.Pointer) *QTextDocumentWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, device, format)
	gopp.ErrPrint(err, rv)
	return &QTextDocumentWriter{cthis}
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:59
// index:2
// void QTextDocumentWriter(const class QString &, const class QByteArray &)
func NewQTextDocumentWriter_2(fileName unsafe.Pointer, format unsafe.Pointer) *QTextDocumentWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, fileName, format)
	gopp.ErrPrint(err, rv)
	return &QTextDocumentWriter{cthis}
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:60
// index:0
// void ~QTextDocumentWriter()
func DeleteQTextDocumentWriter(*QTextDocumentWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:62
// index:0
// void setFormat(const class QByteArray &)
func (this *QTextDocumentWriter) SetFormat(format unsafe.Pointer) {
	// 0: (, const QByteArray & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setFormatERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:63
// index:0
// QByteArray format()
func (this *QTextDocumentWriter) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:65
// index:0
// void setDevice(class QIODevice *)
func (this *QTextDocumentWriter) SetDevice(device unsafe.Pointer) {
	// 0: (, QIODevice * device), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.cthis, device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:66
// index:0
// QIODevice * device()
func (this *QTextDocumentWriter) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6deviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:67
// index:0
// void setFileName(const class QString &)
func (this *QTextDocumentWriter) SetFileName(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:68
// index:0
// QString fileName()
func (this *QTextDocumentWriter) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:70
// index:0
// bool write(const class QTextDocument *)
func (this *QTextDocumentWriter) Write(document unsafe.Pointer) {
	// 0: (, const QTextDocument * document), (document)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeEPK13QTextDocument", ffiqt.FFI_TYPE_VOID, this.cthis, document)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:71
// index:1
// bool write(const class QTextDocumentFragment &)
func (this *QTextDocumentWriter) Write_1(fragment unsafe.Pointer) {
	// 1: (, const QTextDocumentFragment & fragment), (fragment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment", ffiqt.FFI_TYPE_VOID, this.cthis, fragment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:74
// index:0
// void setCodec(class QTextCodec *)
func (this *QTextDocumentWriter) SetCodec(codec unsafe.Pointer) {
	// 0: (, QTextCodec * codec), (codec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_VOID, this.cthis, codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:75
// index:0
// QTextCodec * codec()
func (this *QTextDocumentWriter) Codec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter5codecEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:78
// index:0
// static
// QList<QByteArray> supportedDocumentFormats()
func (this *QTextDocumentWriter) SupportedDocumentFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter24supportedDocumentFormatsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextDocumentWriter_SupportedDocumentFormats() {
	// 0: (), ()
	var nilthis *QTextDocumentWriter
	nilthis.SupportedDocumentFormats()
}

//  body block end
