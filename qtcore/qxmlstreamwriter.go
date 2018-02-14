package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QXmlStreamWriter struct {
	*qtrt.CObject
}
type QXmlStreamWriter_ITF interface {
	QXmlStreamWriter_PTR() *QXmlStreamWriter
}

func (ptr *QXmlStreamWriter) QXmlStreamWriter_PTR() *QXmlStreamWriter { return ptr }

func (this *QXmlStreamWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamWriter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamWriterFromPointer(cthis unsafe.Pointer) *QXmlStreamWriter {
	return &QXmlStreamWriter{&qtrt.CObject{cthis}}
}
func (*QXmlStreamWriter) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamWriter {
	return NewQXmlStreamWriterFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:472
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamWriter()
func NewQXmlStreamWriter() *QXmlStreamWriter {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamWriter)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:473
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamWriter(QIODevice *)
func NewQXmlStreamWriter_1(device QIODevice_ITF /*777 QIODevice **/) *QXmlStreamWriter {
	var convArg0 = device.QIODevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamWriter)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:474
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamWriter(QByteArray *)
func NewQXmlStreamWriter_2(array QByteArray_ITF /*777 QByteArray **/) *QXmlStreamWriter {
	var convArg0 = array.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamWriter)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:475
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamWriter(QString *)
func NewQXmlStreamWriter_3(string string) *QXmlStreamWriter {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamWriter)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:476
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamWriter()
func DeleteQXmlStreamWriter(this *QXmlStreamWriter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:478
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QXmlStreamWriter) SetDevice(device QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 = device.QIODevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:479
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QXmlStreamWriter) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamWriter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qxmlstream.h:482
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)
func (this *QXmlStreamWriter) SetCodec(codec QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 = codec.QTextCodec_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter8setCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:483
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const char *)
func (this *QXmlStreamWriter) SetCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter8setCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:484
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec()
func (this *QXmlStreamWriter) Codec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamWriter5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qxmlstream.h:487
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFormatting(_Bool)
func (this *QXmlStreamWriter) SetAutoFormatting(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter17setAutoFormattingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:488
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoFormatting()
func (this *QXmlStreamWriter) AutoFormatting() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamWriter14autoFormattingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:490
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFormattingIndent(int)
func (this *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter23setAutoFormattingIndentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacesOrTabs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:491
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoFormattingIndent()
func (this *QXmlStreamWriter) AutoFormattingIndent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamWriter20autoFormattingIndentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qxmlstream.h:493
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeAttribute(const QString &, const QString &)
func (this *QXmlStreamWriter) WriteAttribute(qualifiedName string, value string) {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:494
// index:1
// Public Visibility=Default Availability=Available
// [-2] void writeAttribute(const QString &, const QString &, const QString &)
func (this *QXmlStreamWriter) WriteAttribute_1(namespaceUri string, name string, value string) {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(value)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:495
// index:2
// Public Visibility=Default Availability=Available
// [-2] void writeAttribute(const QXmlStreamAttribute &)
func (this *QXmlStreamWriter) WriteAttribute_2(attribute QXmlStreamAttribute_ITF) {
	var convArg0 = attribute.QXmlStreamAttribute_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:496
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeAttributes(const QXmlStreamAttributes &)
func (this *QXmlStreamWriter) WriteAttributes(attributes QXmlStreamAttributes_ITF) {
	var convArg0 = attributes.QXmlStreamAttributes_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:498
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeCDATA(const QString &)
func (this *QXmlStreamWriter) WriteCDATA(text string) {
	var tmpArg0 = NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter10writeCDATAERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:499
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeCharacters(const QString &)
func (this *QXmlStreamWriter) WriteCharacters(text string) {
	var tmpArg0 = NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeCharactersERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:500
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeComment(const QString &)
func (this *QXmlStreamWriter) WriteComment(text string) {
	var tmpArg0 = NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter12writeCommentERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:502
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeDTD(const QString &)
func (this *QXmlStreamWriter) WriteDTD(dtd string) {
	var tmpArg0 = NewQString_5(dtd)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter8writeDTDERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:504
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeEmptyElement(const QString &)
func (this *QXmlStreamWriter) WriteEmptyElement(qualifiedName string) {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeEmptyElementERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:505
// index:1
// Public Visibility=Default Availability=Available
// [-2] void writeEmptyElement(const QString &, const QString &)
func (this *QXmlStreamWriter) WriteEmptyElement_1(namespaceUri string, name string) {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:507
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeTextElement(const QString &, const QString &)
func (this *QXmlStreamWriter) WriteTextElement(qualifiedName string, text string) {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:508
// index:1
// Public Visibility=Default Availability=Available
// [-2] void writeTextElement(const QString &, const QString &, const QString &)
func (this *QXmlStreamWriter) WriteTextElement_1(namespaceUri string, name string, text string) {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(text)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:510
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeEndDocument()
func (this *QXmlStreamWriter) WriteEndDocument() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeEndDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:511
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeEndElement()
func (this *QXmlStreamWriter) WriteEndElement() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeEndElementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:513
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeEntityReference(const QString &)
func (this *QXmlStreamWriter) WriteEntityReference(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter20writeEntityReferenceERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:514
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeNamespace(const QString &, const QString &)
func (this *QXmlStreamWriter) WriteNamespace(namespaceUri string, prefix string) {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:515
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeDefaultNamespace(const QString &)
func (this *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri string) {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:516
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeProcessingInstruction(const QString &, const QString &)
func (this *QXmlStreamWriter) WriteProcessingInstruction(target string, data string) {
	var tmpArg0 = NewQString_5(target)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(data)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:518
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeStartDocument()
func (this *QXmlStreamWriter) WriteStartDocument() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:519
// index:1
// Public Visibility=Default Availability=Available
// [-2] void writeStartDocument(const QString &)
func (this *QXmlStreamWriter) WriteStartDocument_1(version string) {
	var tmpArg0 = NewQString_5(version)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:520
// index:2
// Public Visibility=Default Availability=Available
// [-2] void writeStartDocument(const QString &, _Bool)
func (this *QXmlStreamWriter) WriteStartDocument_2(version string, standalone bool) {
	var tmpArg0 = NewQString_5(version)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, standalone)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:521
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeStartElement(const QString &)
func (this *QXmlStreamWriter) WriteStartElement(qualifiedName string) {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeStartElementERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:522
// index:1
// Public Visibility=Default Availability=Available
// [-2] void writeStartElement(const QString &, const QString &)
func (this *QXmlStreamWriter) WriteStartElement_1(namespaceUri string, name string) {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:525
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeCurrentToken(const QXmlStreamReader &)
func (this *QXmlStreamWriter) WriteCurrentToken(reader QXmlStreamReader_ITF) {
	var convArg0 = reader.QXmlStreamReader_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:528
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasError()
func (this *QXmlStreamWriter) HasError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamWriter8hasErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
		qtrt.KeepMe()
	}
}

//  keep block end
