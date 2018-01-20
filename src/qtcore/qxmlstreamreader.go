//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QXmlStreamReader struct {
	*qtrt.CObject
}

func (this *QXmlStreamReader) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qxmlstream.h:360
// index:0
// void QXmlStreamReader()
func NewQXmlStreamReader() *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}
func NewQXmlStreamReaderFromPointer(cthis unsafe.Pointer) *QXmlStreamReader {
	return &QXmlStreamReader{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:361
// index:1
// void QXmlStreamReader(class QIODevice *)
func NewQXmlStreamReader_1(device unsafe.Pointer) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, device)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:362
// index:2
// void QXmlStreamReader(const class QByteArray &)
func NewQXmlStreamReader_2(data unsafe.Pointer) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, data)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:363
// index:3
// void QXmlStreamReader(const class QString &)
func NewQXmlStreamReader_3(data unsafe.Pointer) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, data)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:364
// index:4
// void QXmlStreamReader(const char *)
func NewQXmlStreamReader_4(data unsafe.Pointer) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, data)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:365
// index:0
// void ~QXmlStreamReader()
func DeleteQXmlStreamReader(*QXmlStreamReader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:367
// index:0
// void setDevice(class QIODevice *)
func (this *QXmlStreamReader) SetDevice(device unsafe.Pointer) {
	// 0: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:368
// index:0
// QIODevice * device()
func (this *QXmlStreamReader) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:369
// index:0
// void addData(const class QByteArray &)
func (this *QXmlStreamReader) AddData(data unsafe.Pointer) {
	// 0: (, data const QByteArray &), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:370
// index:1
// void addData(const class QString &)
func (this *QXmlStreamReader) AddData_1(data unsafe.Pointer) {
	// 1: (, data const QString &), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:371
// index:2
// void addData(const char *)
func (this *QXmlStreamReader) AddData_2(data unsafe.Pointer) {
	// 2: (, data const char *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:372
// index:0
// void clear()
func (this *QXmlStreamReader) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:375
// index:0
// bool atEnd()
func (this *QXmlStreamReader) AtEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader5atEndEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:376
// index:0
// QXmlStreamReader::TokenType readNext()
func (this *QXmlStreamReader) ReadNext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader8readNextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:378
// index:0
// bool readNextStartElement()
func (this *QXmlStreamReader) ReadNextStartElement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader20readNextStartElementEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:379
// index:0
// void skipCurrentElement()
func (this *QXmlStreamReader) SkipCurrentElement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader18skipCurrentElementEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:381
// index:0
// QXmlStreamReader::TokenType tokenType()
func (this *QXmlStreamReader) TokenType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader9tokenTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:382
// index:0
// QString tokenString()
func (this *QXmlStreamReader) TokenString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11tokenStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:384
// index:0
// void setNamespaceProcessing(_Bool)
func (this *QXmlStreamReader) SetNamespaceProcessing(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader22setNamespaceProcessingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:385
// index:0
// bool namespaceProcessing()
func (this *QXmlStreamReader) NamespaceProcessing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader19namespaceProcessingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:387
// index:0
// inline
// bool isStartDocument()
func (this *QXmlStreamReader) IsStartDocument() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader15isStartDocumentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:388
// index:0
// inline
// bool isEndDocument()
func (this *QXmlStreamReader) IsEndDocument() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader13isEndDocumentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:389
// index:0
// inline
// bool isStartElement()
func (this *QXmlStreamReader) IsStartElement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader14isStartElementEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:390
// index:0
// inline
// bool isEndElement()
func (this *QXmlStreamReader) IsEndElement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isEndElementEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:391
// index:0
// inline
// bool isCharacters()
func (this *QXmlStreamReader) IsCharacters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isCharactersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:392
// index:0
// bool isWhitespace()
func (this *QXmlStreamReader) IsWhitespace() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isWhitespaceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:393
// index:0
// bool isCDATA()
func (this *QXmlStreamReader) IsCDATA() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader7isCDATAEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:394
// index:0
// inline
// bool isComment()
func (this *QXmlStreamReader) IsComment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader9isCommentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:395
// index:0
// inline
// bool isDTD()
func (this *QXmlStreamReader) IsDTD() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader5isDTDEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:396
// index:0
// inline
// bool isEntityReference()
func (this *QXmlStreamReader) IsEntityReference() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader17isEntityReferenceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:397
// index:0
// inline
// bool isProcessingInstruction()
func (this *QXmlStreamReader) IsProcessingInstruction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader23isProcessingInstructionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:399
// index:0
// bool isStandaloneDocument()
func (this *QXmlStreamReader) IsStandaloneDocument() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader20isStandaloneDocumentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:400
// index:0
// QStringRef documentVersion()
func (this *QXmlStreamReader) DocumentVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader15documentVersionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:401
// index:0
// QStringRef documentEncoding()
func (this *QXmlStreamReader) DocumentEncoding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader16documentEncodingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:403
// index:0
// qint64 lineNumber()
func (this *QXmlStreamReader) LineNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader10lineNumberEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:404
// index:0
// qint64 columnNumber()
func (this *QXmlStreamReader) ColumnNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12columnNumberEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:405
// index:0
// qint64 characterOffset()
func (this *QXmlStreamReader) CharacterOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader15characterOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:407
// index:0
// QXmlStreamAttributes attributes()
func (this *QXmlStreamReader) Attributes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader10attributesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:414
// index:0
// QString readElementText(enum QXmlStreamReader::ReadElementTextBehaviour)
func (this *QXmlStreamReader) ReadElementText(behaviour int) {
	// 0: (, behaviour QXmlStreamReader::ReadElementTextBehaviour), (&behaviour)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader15readElementTextENS_24ReadElementTextBehaviourE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &behaviour)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:416
// index:0
// QStringRef name()
func (this *QXmlStreamReader) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:417
// index:0
// QStringRef namespaceUri()
func (this *QXmlStreamReader) NamespaceUri() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12namespaceUriEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:418
// index:0
// QStringRef qualifiedName()
func (this *QXmlStreamReader) QualifiedName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader13qualifiedNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:419
// index:0
// QStringRef prefix()
func (this *QXmlStreamReader) Prefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader6prefixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:421
// index:0
// QStringRef processingInstructionTarget()
func (this *QXmlStreamReader) ProcessingInstructionTarget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader27processingInstructionTargetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:422
// index:0
// QStringRef processingInstructionData()
func (this *QXmlStreamReader) ProcessingInstructionData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader25processingInstructionDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:424
// index:0
// QStringRef text()
func (this *QXmlStreamReader) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:426
// index:0
// QXmlStreamNamespaceDeclarations namespaceDeclarations()
func (this *QXmlStreamReader) NamespaceDeclarations() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader21namespaceDeclarationsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:427
// index:0
// void addExtraNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
func (this *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaraction unsafe.Pointer) {
	// 0: (, extraNamespaceDeclaraction const QXmlStreamNamespaceDeclaration &), (extraNamespaceDeclaraction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration", ffiqt.FFI_TYPE_VOID, this.GetCthis(), extraNamespaceDeclaraction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:429
// index:0
// QXmlStreamNotationDeclarations notationDeclarations()
func (this *QXmlStreamReader) NotationDeclarations() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader20notationDeclarationsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:430
// index:0
// QXmlStreamEntityDeclarations entityDeclarations()
func (this *QXmlStreamReader) EntityDeclarations() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader18entityDeclarationsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:431
// index:0
// QStringRef dtdName()
func (this *QXmlStreamReader) DtdName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader7dtdNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:432
// index:0
// QStringRef dtdPublicId()
func (this *QXmlStreamReader) DtdPublicId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11dtdPublicIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:433
// index:0
// QStringRef dtdSystemId()
func (this *QXmlStreamReader) DtdSystemId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11dtdSystemIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:443
// index:0
// void raiseError(const class QString &)
func (this *QXmlStreamReader) RaiseError(message unsafe.Pointer) {
	// 0: (, message const QString &), (message)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader10raiseErrorERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), message)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:444
// index:0
// QString errorString()
func (this *QXmlStreamReader) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:445
// index:0
// QXmlStreamReader::Error error()
func (this *QXmlStreamReader) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader5errorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:447
// index:0
// inline
// bool hasError()
func (this *QXmlStreamReader) HasError() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader8hasErrorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:452
// index:0
// void setEntityResolver(class QXmlStreamEntityResolver *)
func (this *QXmlStreamReader) SetEntityResolver(resolver unsafe.Pointer) {
	// 0: (, resolver QXmlStreamEntityResolver *), (resolver)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver", ffiqt.FFI_TYPE_VOID, this.GetCthis(), resolver)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:453
// index:0
// QXmlStreamEntityResolver * entityResolver()
func (this *QXmlStreamReader) EntityResolver() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader14entityResolverEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
