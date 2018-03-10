package qtgui

// /usr/include/qt/QtGui/qtextdocumentwriter.h
// #include <qtextdocumentwriter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QTextDocumentWriter struct {
	*qtrt.CObject
}
type QTextDocumentWriter_ITF interface {
	QTextDocumentWriter_PTR() *QTextDocumentWriter
}

func (ptr *QTextDocumentWriter) QTextDocumentWriter_PTR() *QTextDocumentWriter { return ptr }

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

/*
Constructs an empty QTextDocumentWriter object. Before writing, you must call setFormat() to set a document format, then setDevice() or setFileName().
*/
func NewQTextDocumentWriter() *QTextDocumentWriter {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentWriter(QIODevice *, const QByteArray &)

/*
Constructs an empty QTextDocumentWriter object. Before writing, you must call setFormat() to set a document format, then setDevice() or setFileName().
*/
func NewQTextDocumentWriter_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QTextDocumentWriter {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2EP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:59
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentWriter(const QString &, const QByteArray &)

/*
Constructs an empty QTextDocumentWriter object. Before writing, you must call setFormat() to set a document format, then setDevice() or setFileName().
*/
func NewQTextDocumentWriter_2(fileName string, format qtcore.QByteArray_ITF) *QTextDocumentWriter {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:59
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentWriter(const QString &, const QByteArray &)

/*
Constructs an empty QTextDocumentWriter object. Before writing, you must call setFormat() to set a document format, then setDevice() or setFileName().
*/
func NewQTextDocumentWriter_2_(fileName string) *QTextDocumentWriter {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentWriter)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextDocumentWriter()

/*

 */
func DeleteQTextDocumentWriter(this *QTextDocumentWriter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)

/*
Sets the format used to write documents to the format specified. format is a case insensitive text string. For example:


          QTextDocumentWriter writer;
          writer.setFormat("odf"); // same as writer.setFormat("ODF");



You can call supportedDocumentFormats() for the full list of formats QTextDocumentWriter supports.

See also format().
*/
func (this *QTextDocumentWriter) SetFormat(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format() const

/*
Returns the format used for writing documents.

See also setFormat().
*/
func (this *QTextDocumentWriter) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets the writer's device to the device specified. If a device has already been set, the old device is removed but otherwise left unchanged.

If the device is not already open, QTextDocumentWriter will attempt to open the device in QIODevice::WriteOnly mode by calling open().

Note: This will not work for certain devices, such as QProcess, QTcpSocket and QUdpSocket, where some configuration is required before the device can be opened.

See also device() and setFileName().
*/
func (this *QTextDocumentWriter) SetDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the device currently assigned, or 0 if no device has been assigned.

See also setDevice().
*/
func (this *QTextDocumentWriter) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextDocumentWriter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*
Sets the name of the file to be written to fileName. Internally, QTextDocumentWriter will create a QFile and open it in QIODevice::WriteOnly mode, and use this file when writing the document.

See also fileName() and setDevice().
*/
func (this *QTextDocumentWriter) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
If the currently assigned device is a QFile, or if setFileName() has been called, this function returns the name of the file to be written to. In all other cases, it returns an empty string.

See also setFileName() and setDevice().
*/
func (this *QTextDocumentWriter) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextDocumentWriter8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(const QTextDocument *)

/*
Writes the given document to the assigned device or file and returns true if successful; otherwise returns false.
*/
func (this *QTextDocumentWriter) Write(document QTextDocument_ITF /*777 const QTextDocument **/) bool {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeEPK13QTextDocument", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:71
// index:1
// Public Visibility=Default Availability=Available
// [1] bool write(const QTextDocumentFragment &)

/*
Writes the given document to the assigned device or file and returns true if successful; otherwise returns false.
*/
func (this *QTextDocumentWriter) Write_1(fragment QTextDocumentFragment_ITF) bool {
	var convArg0 unsafe.Pointer
	if fragment != nil && fragment.QTextDocumentFragment_PTR() != nil {
		convArg0 = fragment.QTextDocumentFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)

/*
Sets the codec for this stream to codec. The codec is used for encoding any data that is written. By default, QTextDocumentWriter uses UTF-8.

See also codec().
*/
func (this *QTextDocumentWriter) SetCodec(codec qtcore.QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter8setCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec() const

/*
Returns the codec that is currently assigned to the writer.

See also setCodec().
*/
func (this *QTextDocumentWriter) Codec() *qtcore.QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextDocumentWriter5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
