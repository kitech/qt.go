package qtgui

// /usr/include/qt/QtGui/qtextdocumentwriter.h
// #include <qtextdocumentwriter.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextDocumentWriter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextDocumentWriterFromPointer(cthis unsafe.Pointer) *QTextDocumentWriter {
	return &QTextDocumentWriter{&qtrt.CObject{cthis}}
}
func (*QTextDocumentWriter) NewFromPointer(cthis unsafe.Pointer) *QTextDocumentWriter {
	return NewQTextDocumentWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentWriter()
func NewQTextDocumentWriter() *QTextDocumentWriter {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentWriter(QIODevice *, const QByteArray &)
func NewQTextDocumentWriter_1(device *qtcore.QIODevice /*777 QIODevice **/, format *qtcore.QByteArray) *QTextDocumentWriter {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:59
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentWriter(const QString &, const QByteArray &)
func NewQTextDocumentWriter_2(fileName *qtcore.QString, format *qtcore.QByteArray) *QTextDocumentWriter {
	var convArg0 = fileName.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextDocumentWriter()
func DeleteQTextDocumentWriter(this *QTextDocumentWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriterD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)
func (this *QTextDocumentWriter) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format()
func (this *QTextDocumentWriter) Format() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QTextDocumentWriter) SetDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QTextDocumentWriter) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QTextDocumentWriter) SetFileName(fileName *qtcore.QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QTextDocumentWriter) FileName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(const QTextDocument *)
func (this *QTextDocumentWriter) Write(document *QTextDocument /*777 const QTextDocument **/) bool {
	var convArg0 = document.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeEPK13QTextDocument", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:71
// index:1
// Public Visibility=Default Availability=Available
// [1] bool write(const QTextDocumentFragment &)
func (this *QTextDocumentWriter) Write_1(fragment *QTextDocumentFragment) bool {
	var convArg0 = fragment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)
func (this *QTextDocumentWriter) SetCodec(codec *qtcore.QTextCodec /*777 QTextCodec **/) {
	var convArg0 = codec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QTextDocumentWriter8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec()
func (this *QTextDocumentWriter) Codec() *qtcore.QTextCodec /*777 QTextCodec **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QTextDocumentWriter5codecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
