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
	*qtrt.CObject
}

func (this *QTextDocumentWriter) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextDocumentWriterFromPointer(cthis unsafe.Pointer) *QTextDocumentWriter {
	return &QTextDocumentWriter{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:57
// index:0
// Public
// void QTextDocumentWriter()
func NewQTextDocumentWriter() *QTextDocumentWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:58
// index:1
// Public
// void QTextDocumentWriter(class QIODevice *, const class QByteArray &)
func NewQTextDocumentWriter_1(device unsafe.Pointer, format *qtcore.QByteArray) *QTextDocumentWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, device, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:59
// index:2
// Public
// void QTextDocumentWriter(const class QString &, const class QByteArray &)
func NewQTextDocumentWriter_2(fileName *qtcore.QString, format *qtcore.QByteArray) *QTextDocumentWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:60
// index:0
// Public
// void ~QTextDocumentWriter()
func DeleteQTextDocumentWriter(*QTextDocumentWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:62
// index:0
// Public
// void setFormat(const class QByteArray &)
func (this *QTextDocumentWriter) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:63
// index:0
// Public
// QByteArray format()
func (this *QTextDocumentWriter) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:65
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QTextDocumentWriter) SetDevice(device unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:66
// index:0
// Public
// QIODevice * device()
func (this *QTextDocumentWriter) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:67
// index:0
// Public
// void setFileName(const class QString &)
func (this *QTextDocumentWriter) SetFileName(fileName *qtcore.QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:68
// index:0
// Public
// QString fileName()
func (this *QTextDocumentWriter) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:70
// index:0
// Public
// bool write(const class QTextDocument *)
func (this *QTextDocumentWriter) Write(document unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeEPK13QTextDocument", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), document)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:71
// index:1
// Public
// bool write(const class QTextDocumentFragment &)
func (this *QTextDocumentWriter) Write_1(fragment *QTextDocumentFragment) interface{} {
	var convArg0 = fragment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:74
// index:0
// Public
// void setCodec(class QTextCodec *)
func (this *QTextDocumentWriter) SetCodec(codec unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:75
// index:0
// Public
// QTextCodec * codec()
func (this *QTextDocumentWriter) Codec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter5codecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:78
// index:0
// Public static
// QList<QByteArray> supportedDocumentFormats()
func (this *QTextDocumentWriter) SupportedDocumentFormats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter24supportedDocumentFormatsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextDocumentWriter_SupportedDocumentFormats() {
	var nilthis *QTextDocumentWriter
	nilthis.SupportedDocumentFormats()
}

//  body block end
