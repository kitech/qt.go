//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 59
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QXmlStreamWriter struct {
	*qtrt.CObject
}

func (this *QXmlStreamWriter) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qxmlstream.h:472
// index:0
// void QXmlStreamWriter()
func NewQXmlStreamWriter() *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}
func NewQXmlStreamWriterFromPointer(cthis unsafe.Pointer) *QXmlStreamWriter {
	return &QXmlStreamWriter{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:473
// index:1
// void QXmlStreamWriter(class QIODevice *)
func NewQXmlStreamWriter_1(device unsafe.Pointer) *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, device)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:474
// index:2
// void QXmlStreamWriter(class QByteArray *)
func NewQXmlStreamWriter_2(array unsafe.Pointer) *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, array)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:475
// index:3
// void QXmlStreamWriter(class QString *)
func NewQXmlStreamWriter_3(string unsafe.Pointer) *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP7QString", ffiqt.FFI_TYPE_VOID, cthis, string)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:476
// index:0
// void ~QXmlStreamWriter()
func DeleteQXmlStreamWriter(*QXmlStreamWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:478
// index:0
// void setDevice(class QIODevice *)
func (this *QXmlStreamWriter) SetDevice(device unsafe.Pointer) {
	// 0: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:479
// index:0
// QIODevice * device()
func (this *QXmlStreamWriter) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:482
// index:0
// void setCodec(class QTextCodec *)
func (this *QXmlStreamWriter) SetCodec(codec unsafe.Pointer) {
	// 0: (, codec QTextCodec *), (codec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_VOID, this.GetCthis(), codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:483
// index:1
// void setCodec(const char *)
func (this *QXmlStreamWriter) SetCodec_1(codecName unsafe.Pointer) {
	// 1: (, codecName const char *), (codecName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter8setCodecEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), codecName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:484
// index:0
// QTextCodec * codec()
func (this *QXmlStreamWriter) Codec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter5codecEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:487
// index:0
// void setAutoFormatting(_Bool)
func (this *QXmlStreamWriter) SetAutoFormatting(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17setAutoFormattingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:488
// index:0
// bool autoFormatting()
func (this *QXmlStreamWriter) AutoFormatting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter14autoFormattingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:490
// index:0
// void setAutoFormattingIndent(int)
func (this *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	// 0: (, spacesOrTabs int), (&spacesOrTabs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter23setAutoFormattingIndentEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacesOrTabs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:491
// index:0
// int autoFormattingIndent()
func (this *QXmlStreamWriter) AutoFormattingIndent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter20autoFormattingIndentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:493
// index:0
// void writeAttribute(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteAttribute(qualifiedName unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, qualifiedName const QString &, value const QString &), (qualifiedName, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:494
// index:1
// void writeAttribute(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteAttribute_1(namespaceUri unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	// 1: (, namespaceUri const QString &, name const QString &, value const QString &), (namespaceUri, name, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:495
// index:2
// void writeAttribute(const class QXmlStreamAttribute &)
func (this *QXmlStreamWriter) WriteAttribute_2(attribute unsafe.Pointer) {
	// 2: (, attribute const QXmlStreamAttribute &), (attribute)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute", ffiqt.FFI_TYPE_VOID, this.GetCthis(), attribute)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:496
// index:0
// void writeAttributes(const class QXmlStreamAttributes &)
func (this *QXmlStreamWriter) WriteAttributes(attributes unsafe.Pointer) {
	// 0: (, attributes const QXmlStreamAttributes &), (attributes)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes", ffiqt.FFI_TYPE_VOID, this.GetCthis(), attributes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:498
// index:0
// void writeCDATA(const class QString &)
func (this *QXmlStreamWriter) WriteCDATA(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter10writeCDATAERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:499
// index:0
// void writeCharacters(const class QString &)
func (this *QXmlStreamWriter) WriteCharacters(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeCharactersERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:500
// index:0
// void writeComment(const class QString &)
func (this *QXmlStreamWriter) WriteComment(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter12writeCommentERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:502
// index:0
// void writeDTD(const class QString &)
func (this *QXmlStreamWriter) WriteDTD(dtd unsafe.Pointer) {
	// 0: (, dtd const QString &), (dtd)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter8writeDTDERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), dtd)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:504
// index:0
// void writeEmptyElement(const class QString &)
func (this *QXmlStreamWriter) WriteEmptyElement(qualifiedName unsafe.Pointer) {
	// 0: (, qualifiedName const QString &), (qualifiedName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeEmptyElementERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:505
// index:1
// void writeEmptyElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteEmptyElement_1(namespaceUri unsafe.Pointer, name unsafe.Pointer) {
	// 1: (, namespaceUri const QString &, name const QString &), (namespaceUri, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:507
// index:0
// void writeTextElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteTextElement(qualifiedName unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, qualifiedName const QString &, text const QString &), (qualifiedName, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:508
// index:1
// void writeTextElement(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteTextElement_1(namespaceUri unsafe.Pointer, name unsafe.Pointer, text unsafe.Pointer) {
	// 1: (, namespaceUri const QString &, name const QString &, text const QString &), (namespaceUri, name, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:510
// index:0
// void writeEndDocument()
func (this *QXmlStreamWriter) WriteEndDocument() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeEndDocumentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:511
// index:0
// void writeEndElement()
func (this *QXmlStreamWriter) WriteEndElement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeEndElementEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:513
// index:0
// void writeEntityReference(const class QString &)
func (this *QXmlStreamWriter) WriteEntityReference(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter20writeEntityReferenceERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:514
// index:0
// void writeNamespace(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteNamespace(namespaceUri unsafe.Pointer, prefix unsafe.Pointer) {
	// 0: (, namespaceUri const QString &, prefix const QString &), (namespaceUri, prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:515
// index:0
// void writeDefaultNamespace(const class QString &)
func (this *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri unsafe.Pointer) {
	// 0: (, namespaceUri const QString &), (namespaceUri)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:516
// index:0
// void writeProcessingInstruction(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteProcessingInstruction(target unsafe.Pointer, data unsafe.Pointer) {
	// 0: (, target const QString &, data const QString &), (target, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), target, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:518
// index:0
// void writeStartDocument()
func (this *QXmlStreamWriter) WriteStartDocument() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:519
// index:1
// void writeStartDocument(const class QString &)
func (this *QXmlStreamWriter) WriteStartDocument_1(version unsafe.Pointer) {
	// 1: (, version const QString &), (version)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:520
// index:2
// void writeStartDocument(const class QString &, _Bool)
func (this *QXmlStreamWriter) WriteStartDocument_2(version unsafe.Pointer, standalone bool) {
	// 2: (, version const QString &, standalone bool), (version, &standalone)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), version, &standalone)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:521
// index:0
// void writeStartElement(const class QString &)
func (this *QXmlStreamWriter) WriteStartElement(qualifiedName unsafe.Pointer) {
	// 0: (, qualifiedName const QString &), (qualifiedName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeStartElementERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:522
// index:1
// void writeStartElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteStartElement_1(namespaceUri unsafe.Pointer, name unsafe.Pointer) {
	// 1: (, namespaceUri const QString &, name const QString &), (namespaceUri, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:525
// index:0
// void writeCurrentToken(const class QXmlStreamReader &)
func (this *QXmlStreamWriter) WriteCurrentToken(reader unsafe.Pointer) {
	// 0: (, reader const QXmlStreamReader &), (reader)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader", ffiqt.FFI_TYPE_VOID, this.GetCthis(), reader)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:528
// index:0
// bool hasError()
func (this *QXmlStreamWriter) HasError() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter8hasErrorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
