package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQXmlStreamWriterFromPointer(cthis unsafe.Pointer) *QXmlStreamWriter {
	return &QXmlStreamWriter{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:472
// index:0
// Public
// void QXmlStreamWriter()
func NewQXmlStreamWriter() *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:473
// index:1
// Public
// void QXmlStreamWriter(class QIODevice *)
func NewQXmlStreamWriter_1(device *QIODevice /*444 QIODevice **/) *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:474
// index:2
// Public
// void QXmlStreamWriter(class QByteArray *)
func NewQXmlStreamWriter_2(array *QByteArray /*444 QByteArray **/) *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = array.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:475
// index:3
// Public
// void QXmlStreamWriter(class QString *)
func NewQXmlStreamWriter_3(string *QString /*444 QString **/) *QXmlStreamWriter {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterC2EP7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:476
// index:0
// Public
// void ~QXmlStreamWriter()
func DeleteQXmlStreamWriter(*QXmlStreamWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:478
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QXmlStreamWriter) SetDevice(device *QIODevice /*444 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:479
// index:0
// Public
// QIODevice * device()
func (this *QXmlStreamWriter) Device() *QIODevice /*444 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:482
// index:0
// Public
// void setCodec(class QTextCodec *)
func (this *QXmlStreamWriter) SetCodec(codec *QTextCodec /*444 QTextCodec **/) {
	var convArg0 = codec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:483
// index:1
// Public
// void setCodec(const char *)
func (this *QXmlStreamWriter) SetCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter8setCodecEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:484
// index:0
// Public
// QTextCodec * codec()
func (this *QXmlStreamWriter) Codec() *QTextCodec /*444 QTextCodec **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter5codecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:487
// index:0
// Public
// void setAutoFormatting(_Bool)
func (this *QXmlStreamWriter) SetAutoFormatting(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17setAutoFormattingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:488
// index:0
// Public
// bool autoFormatting()
func (this *QXmlStreamWriter) AutoFormatting() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter14autoFormattingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:490
// index:0
// Public
// void setAutoFormattingIndent(int)
func (this *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter23setAutoFormattingIndentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacesOrTabs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:491
// index:0
// Public
// int autoFormattingIndent()
func (this *QXmlStreamWriter) AutoFormattingIndent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter20autoFormattingIndentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qxmlstream.h:493
// index:0
// Public
// void writeAttribute(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteAttribute(qualifiedName *QString, value *QString) {
	var convArg0 = qualifiedName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:494
// index:1
// Public
// void writeAttribute(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteAttribute_1(namespaceUri *QString, name *QString, value *QString) {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:495
// index:2
// Public
// void writeAttribute(const class QXmlStreamAttribute &)
func (this *QXmlStreamWriter) WriteAttribute_2(attribute *QXmlStreamAttribute) {
	var convArg0 = attribute.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:496
// index:0
// Public
// void writeAttributes(const class QXmlStreamAttributes &)
func (this *QXmlStreamWriter) WriteAttributes(attributes *QXmlStreamAttributes) {
	var convArg0 = attributes.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:498
// index:0
// Public
// void writeCDATA(const class QString &)
func (this *QXmlStreamWriter) WriteCDATA(text *QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter10writeCDATAERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:499
// index:0
// Public
// void writeCharacters(const class QString &)
func (this *QXmlStreamWriter) WriteCharacters(text *QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeCharactersERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:500
// index:0
// Public
// void writeComment(const class QString &)
func (this *QXmlStreamWriter) WriteComment(text *QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter12writeCommentERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:502
// index:0
// Public
// void writeDTD(const class QString &)
func (this *QXmlStreamWriter) WriteDTD(dtd *QString) {
	var convArg0 = dtd.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter8writeDTDERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:504
// index:0
// Public
// void writeEmptyElement(const class QString &)
func (this *QXmlStreamWriter) WriteEmptyElement(qualifiedName *QString) {
	var convArg0 = qualifiedName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeEmptyElementERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:505
// index:1
// Public
// void writeEmptyElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteEmptyElement_1(namespaceUri *QString, name *QString) {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:507
// index:0
// Public
// void writeTextElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteTextElement(qualifiedName *QString, text *QString) {
	var convArg0 = qualifiedName.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:508
// index:1
// Public
// void writeTextElement(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteTextElement_1(namespaceUri *QString, name *QString, text *QString) {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:510
// index:0
// Public
// void writeEndDocument()
func (this *QXmlStreamWriter) WriteEndDocument() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter16writeEndDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:511
// index:0
// Public
// void writeEndElement()
func (this *QXmlStreamWriter) WriteEndElement() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter15writeEndElementEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:513
// index:0
// Public
// void writeEntityReference(const class QString &)
func (this *QXmlStreamWriter) WriteEntityReference(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter20writeEntityReferenceERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:514
// index:0
// Public
// void writeNamespace(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteNamespace(namespaceUri *QString, prefix *QString) {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:515
// index:0
// Public
// void writeDefaultNamespace(const class QString &)
func (this *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri *QString) {
	var convArg0 = namespaceUri.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:516
// index:0
// Public
// void writeProcessingInstruction(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteProcessingInstruction(target *QString, data *QString) {
	var convArg0 = target.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:518
// index:0
// Public
// void writeStartDocument()
func (this *QXmlStreamWriter) WriteStartDocument() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:519
// index:1
// Public
// void writeStartDocument(const class QString &)
func (this *QXmlStreamWriter) WriteStartDocument_1(version *QString) {
	var convArg0 = version.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:520
// index:2
// Public
// void writeStartDocument(const class QString &, _Bool)
func (this *QXmlStreamWriter) WriteStartDocument_2(version *QString, standalone bool) {
	var convArg0 = version.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &standalone)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:521
// index:0
// Public
// void writeStartElement(const class QString &)
func (this *QXmlStreamWriter) WriteStartElement(qualifiedName *QString) {
	var convArg0 = qualifiedName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeStartElementERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:522
// index:1
// Public
// void writeStartElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) WriteStartElement_1(namespaceUri *QString, name *QString) {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:525
// index:0
// Public
// void writeCurrentToken(const class QXmlStreamReader &)
func (this *QXmlStreamWriter) WriteCurrentToken(reader *QXmlStreamReader) {
	var convArg0 = reader.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:528
// index:0
// Public
// bool hasError()
func (this *QXmlStreamWriter) HasError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamWriter8hasErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
