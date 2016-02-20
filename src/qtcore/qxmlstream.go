package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qxmlstream.h
// dst-file: /src/core/qxmlstream.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QString * QXmlStreamStringRef::string();
extern void* C_ZNK19QXmlStreamStringRef6stringEv(void* qthis); // 2
  // proto:  void QXmlStreamStringRef::~QXmlStreamStringRef();
extern void C_ZN19QXmlStreamStringRefD2Ev(void* qthis); // 2
  // proto:  void QXmlStreamStringRef::clear();
extern void C_ZN19QXmlStreamStringRef5clearEv(void* qthis); // 2
  // proto:  void QXmlStreamStringRef::QXmlStreamStringRef();
extern void* C_ZN19QXmlStreamStringRefC2Ev(); // 1
  // proto:  void QXmlStreamStringRef::QXmlStreamStringRef(const QString & aString);
extern void* C_ZN19QXmlStreamStringRefC2ERK7QString(void* arg0); // 1
  // proto:  int QXmlStreamStringRef::position();
extern int32_t C_ZNK19QXmlStreamStringRef8positionEv(void* qthis); // 2
  // proto:  int QXmlStreamStringRef::size();
extern int32_t C_ZNK19QXmlStreamStringRef4sizeEv(void* qthis); // 2
  // proto:  bool QXmlStreamReader::isEntityReference();
extern bool C_ZNK16QXmlStreamReader17isEntityReferenceEv(void* qthis); // 2
  // proto:  bool QXmlStreamReader::isDTD();
extern bool C_ZNK16QXmlStreamReader5isDTDEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamReader::qualifiedName();
extern void C_ZNK16QXmlStreamReader13qualifiedNameEv(void* qthis); // 4
  // proto:  QXmlStreamReader::TokenType QXmlStreamReader::tokenType();
extern void C_ZNK16QXmlStreamReader9tokenTypeEv(void* qthis); // 4
  // proto:  QXmlStreamReader::TokenType QXmlStreamReader::readNext();
extern void C_ZN16QXmlStreamReader8readNextEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::text();
extern void C_ZNK16QXmlStreamReader4textEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::dtdPublicId();
extern void C_ZNK16QXmlStreamReader11dtdPublicIdEv(void* qthis); // 4
  // proto:  QXmlStreamEntityDeclarations QXmlStreamReader::entityDeclarations();
extern void C_ZNK16QXmlStreamReader18entityDeclarationsEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::setNamespaceProcessing(bool );
extern void C_ZN16QXmlStreamReader22setNamespaceProcessingEb(void* qthis, bool arg0); // 4
  // proto:  QXmlStreamNamespaceDeclarations QXmlStreamReader::namespaceDeclarations();
extern void C_ZNK16QXmlStreamReader21namespaceDeclarationsEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isCharacters();
extern bool C_ZNK16QXmlStreamReader12isCharactersEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamReader::prefix();
extern void C_ZNK16QXmlStreamReader6prefixEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::addExtraNamespaceDeclaration(const QXmlStreamNamespaceDeclaration & extraNamespaceDeclaraction);
extern void C_ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration(void* qthis, void* arg0); // 4
  // proto:  QStringRef QXmlStreamReader::documentEncoding();
extern void C_ZNK16QXmlStreamReader16documentEncodingEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::processingInstructionData();
extern void C_ZNK16QXmlStreamReader25processingInstructionDataEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::processingInstructionTarget();
extern void C_ZNK16QXmlStreamReader27processingInstructionTargetEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::setEntityResolver(QXmlStreamEntityResolver * resolver);
extern void C_ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver(void* qthis, void* arg0); // 4
  // proto:  bool QXmlStreamReader::isProcessingInstruction();
extern bool C_ZNK16QXmlStreamReader23isProcessingInstructionEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamReader::documentVersion();
extern void C_ZNK16QXmlStreamReader15documentVersionEv(void* qthis); // 4
  // proto:  QXmlStreamNotationDeclarations QXmlStreamReader::notationDeclarations();
extern void C_ZNK16QXmlStreamReader20notationDeclarationsEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::namespaceUri();
extern void C_ZNK16QXmlStreamReader12namespaceUriEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::setDevice(QIODevice * device);
extern void C_ZN16QXmlStreamReader9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamReader::raiseError(const QString & message);
extern void C_ZN16QXmlStreamReader10raiseErrorERK7QString(void* qthis, void* arg0); // 4
  // proto:  QXmlStreamEntityResolver * QXmlStreamReader::entityResolver();
extern void* C_ZNK16QXmlStreamReader14entityResolverEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isStartElement();
extern bool C_ZNK16QXmlStreamReader14isStartElementEv(void* qthis); // 2
  // proto:  bool QXmlStreamReader::isWhitespace();
extern bool C_ZNK16QXmlStreamReader12isWhitespaceEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::QXmlStreamReader();
extern void* C_ZN16QXmlStreamReaderC2Ev(); // 3
  // proto:  void QXmlStreamReader::QXmlStreamReader(const QString & data);
extern void* C_ZN16QXmlStreamReaderC2ERK7QString(void* arg0); // 3
  // proto:  void QXmlStreamReader::QXmlStreamReader(const char * data);
extern void* C_ZN16QXmlStreamReaderC2EPKc(void* arg0); // 3
  // proto:  void QXmlStreamReader::QXmlStreamReader(const QByteArray & data);
extern void* C_ZN16QXmlStreamReaderC2ERK10QByteArray(void* arg0); // 3
  // proto:  void QXmlStreamReader::QXmlStreamReader(QIODevice * device);
extern void* C_ZN16QXmlStreamReaderC2EP9QIODevice(void* arg0); // 3
  // proto:  bool QXmlStreamReader::hasError();
extern bool C_ZNK16QXmlStreamReader8hasErrorEv(void* qthis); // 2
  // proto:  qint64 QXmlStreamReader::characterOffset();
extern int64_t C_ZNK16QXmlStreamReader15characterOffsetEv(void* qthis); // 4
  // proto:  QString QXmlStreamReader::errorString();
extern void* C_ZNK16QXmlStreamReader11errorStringEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isEndElement();
extern bool C_ZNK16QXmlStreamReader12isEndElementEv(void* qthis); // 2
  // proto:  bool QXmlStreamReader::isCDATA();
extern bool C_ZNK16QXmlStreamReader7isCDATAEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isStandaloneDocument();
extern bool C_ZNK16QXmlStreamReader20isStandaloneDocumentEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::skipCurrentElement();
extern void C_ZN16QXmlStreamReader18skipCurrentElementEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isComment();
extern bool C_ZNK16QXmlStreamReader9isCommentEv(void* qthis); // 2
  // proto:  qint64 QXmlStreamReader::lineNumber();
extern int64_t C_ZNK16QXmlStreamReader10lineNumberEv(void* qthis); // 4
  // proto:  QIODevice * QXmlStreamReader::device();
extern void* C_ZNK16QXmlStreamReader6deviceEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isEndDocument();
extern bool C_ZNK16QXmlStreamReader13isEndDocumentEv(void* qthis); // 2
  // proto:  qint64 QXmlStreamReader::columnNumber();
extern int64_t C_ZNK16QXmlStreamReader12columnNumberEv(void* qthis); // 4
  // proto:  QString QXmlStreamReader::tokenString();
extern void* C_ZNK16QXmlStreamReader11tokenStringEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::name();
extern void C_ZNK16QXmlStreamReader4nameEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::namespaceProcessing();
extern bool C_ZNK16QXmlStreamReader19namespaceProcessingEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::isStartDocument();
extern bool C_ZNK16QXmlStreamReader15isStartDocumentEv(void* qthis); // 2
  // proto:  void QXmlStreamReader::clear();
extern void C_ZN16QXmlStreamReader5clearEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::atEnd();
extern bool C_ZNK16QXmlStreamReader5atEndEv(void* qthis); // 4
  // proto:  QStringRef QXmlStreamReader::dtdName();
extern void C_ZNK16QXmlStreamReader7dtdNameEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::addData(const QByteArray & data);
extern void C_ZN16QXmlStreamReader7addDataERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamReader::addData(const QString & data);
extern void C_ZN16QXmlStreamReader7addDataERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamReader::addData(const char * data);
extern void C_ZN16QXmlStreamReader7addDataEPKc(void* qthis, void* arg0); // 4
  // proto:  QStringRef QXmlStreamReader::dtdSystemId();
extern void C_ZNK16QXmlStreamReader11dtdSystemIdEv(void* qthis); // 4
  // proto:  bool QXmlStreamReader::readNextStartElement();
extern bool C_ZN16QXmlStreamReader20readNextStartElementEv(void* qthis); // 4
  // proto:  QXmlStreamReader::Error QXmlStreamReader::error();
extern void C_ZNK16QXmlStreamReader5errorEv(void* qthis); // 4
  // proto:  QXmlStreamAttributes QXmlStreamReader::attributes();
extern void* C_ZNK16QXmlStreamReader10attributesEv(void* qthis); // 4
  // proto:  void QXmlStreamReader::~QXmlStreamReader();
extern void C_ZN16QXmlStreamReaderD2Ev(void* qthis); // 4
  // proto:  QString QXmlStreamEntityResolver::resolveEntity(const QString & publicId, const QString & systemId);
extern void* C_ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QXmlStreamEntityResolver::resolveUndeclaredEntity(const QString & name);
extern void* C_ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamEntityResolver::~QXmlStreamEntityResolver();
extern void C_ZN24QXmlStreamEntityResolverD2Ev(void* qthis); // 4
  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration();
extern void* C_ZN30QXmlStreamNamespaceDeclarationC2Ev(); // 3
  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration(const QString & prefix, const QString & namespaceUri);
extern void* C_ZN30QXmlStreamNamespaceDeclarationC2ERK7QStringS2_(void* arg0, void* arg1); // 3
  // proto:  void QXmlStreamNamespaceDeclaration::QXmlStreamNamespaceDeclaration(const QXmlStreamNamespaceDeclaration & );
extern void* C_ZN30QXmlStreamNamespaceDeclarationC2ERKS_(void* arg0); // 3
  // proto:  void QXmlStreamNamespaceDeclaration::~QXmlStreamNamespaceDeclaration();
extern void C_ZN30QXmlStreamNamespaceDeclarationD2Ev(void* qthis); // 4
  // proto:  QStringRef QXmlStreamNamespaceDeclaration::namespaceUri();
extern void C_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamNamespaceDeclaration::prefix();
extern void C_ZNK30QXmlStreamNamespaceDeclaration6prefixEv(void* qthis); // 2
  // proto:  void QXmlStreamEntityDeclaration::~QXmlStreamEntityDeclaration();
extern void C_ZN27QXmlStreamEntityDeclarationD2Ev(void* qthis); // 4
  // proto:  QStringRef QXmlStreamEntityDeclaration::publicId();
extern void C_ZNK27QXmlStreamEntityDeclaration8publicIdEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamEntityDeclaration::name();
extern void C_ZNK27QXmlStreamEntityDeclaration4nameEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamEntityDeclaration::value();
extern void C_ZNK27QXmlStreamEntityDeclaration5valueEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamEntityDeclaration::systemId();
extern void C_ZNK27QXmlStreamEntityDeclaration8systemIdEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamEntityDeclaration::notationName();
extern void C_ZNK27QXmlStreamEntityDeclaration12notationNameEv(void* qthis); // 2
  // proto:  void QXmlStreamEntityDeclaration::QXmlStreamEntityDeclaration(const QXmlStreamEntityDeclaration & );
extern void* C_ZN27QXmlStreamEntityDeclarationC2ERKS_(void* arg0); // 3
  // proto:  void QXmlStreamEntityDeclaration::QXmlStreamEntityDeclaration();
extern void* C_ZN27QXmlStreamEntityDeclarationC2Ev(); // 3
  // proto:  bool QXmlStreamAttributes::hasAttribute(const QString & namespaceUri, const QString & name);
extern bool C_ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  bool QXmlStreamAttributes::hasAttribute(const QString & qualifiedName);
extern bool C_ZNK20QXmlStreamAttributes12hasAttributeERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QXmlStreamAttributes::append(const QString & namespaceUri, const QString & name, const QString & value);
extern void C_ZN20QXmlStreamAttributes6appendERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QXmlStreamAttributes::append(const QString & qualifiedName, const QString & value);
extern void C_ZN20QXmlStreamAttributes6appendERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringRef QXmlStreamAttributes::value(const QString & namespaceUri, const QString & name);
extern void C_ZNK20QXmlStreamAttributes5valueERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringRef QXmlStreamAttributes::value(const QString & qualifiedName);
extern void C_ZNK20QXmlStreamAttributes5valueERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamAttributes::QXmlStreamAttributes();
extern void* C_ZN20QXmlStreamAttributesC2Ev(); // 1
  // proto:  void QXmlStreamWriter::writeTextElement(const QString & namespaceUri, const QString & name, const QString & text);
extern void C_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QXmlStreamWriter::writeTextElement(const QString & qualifiedName, const QString & text);
extern void C_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QXmlStreamWriter::writeNamespace(const QString & namespaceUri, const QString & prefix);
extern void C_ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QXmlStreamWriter::autoFormatting();
extern bool C_ZNK16QXmlStreamWriter14autoFormattingEv(void* qthis); // 4
  // proto:  QIODevice * QXmlStreamWriter::device();
extern void* C_ZNK16QXmlStreamWriter6deviceEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::writeEntityReference(const QString & name);
extern void C_ZN16QXmlStreamWriter20writeEntityReferenceERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeProcessingInstruction(const QString & target, const QString & data);
extern void C_ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QXmlStreamWriter::writeAttributes(const QXmlStreamAttributes & attributes);
extern void C_ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes(void* qthis, void* arg0); // 4
  // proto:  int QXmlStreamWriter::autoFormattingIndent();
extern int32_t C_ZNK16QXmlStreamWriter20autoFormattingIndentEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::setCodec(const char * codecName);
extern void C_ZN16QXmlStreamWriter8setCodecEPKc(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::setCodec(QTextCodec * codec);
extern void C_ZN16QXmlStreamWriter8setCodecEP10QTextCodec(void* qthis, void* arg0); // 4
  // proto:  QTextCodec * QXmlStreamWriter::codec();
extern void* C_ZNK16QXmlStreamWriter5codecEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::writeComment(const QString & text);
extern void C_ZN16QXmlStreamWriter12writeCommentERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeCDATA(const QString & text);
extern void C_ZN16QXmlStreamWriter10writeCDATAERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::setDevice(QIODevice * device);
extern void C_ZN16QXmlStreamWriter9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeAttribute(const QString & namespaceUri, const QString & name, const QString & value);
extern void C_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QXmlStreamWriter::writeAttribute(const QString & qualifiedName, const QString & value);
extern void C_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QXmlStreamWriter::writeAttribute(const QXmlStreamAttribute & attribute);
extern void C_ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::setAutoFormatting(bool );
extern void C_ZN16QXmlStreamWriter17setAutoFormattingEb(void* qthis, bool arg0); // 4
  // proto:  bool QXmlStreamWriter::hasError();
extern bool C_ZNK16QXmlStreamWriter8hasErrorEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::writeDTD(const QString & dtd);
extern void C_ZN16QXmlStreamWriter8writeDTDERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(QByteArray * array);
extern void* C_ZN16QXmlStreamWriterC2EP10QByteArray(void* arg0); // 3
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(QString * string);
extern void* C_ZN16QXmlStreamWriterC2EP7QString(void* arg0); // 3
  // proto:  void QXmlStreamWriter::QXmlStreamWriter();
extern void* C_ZN16QXmlStreamWriterC2Ev(); // 3
  // proto:  void QXmlStreamWriter::QXmlStreamWriter(QIODevice * device);
extern void* C_ZN16QXmlStreamWriterC2EP9QIODevice(void* arg0); // 3
  // proto:  void QXmlStreamWriter::writeStartElement(const QString & namespaceUri, const QString & name);
extern void C_ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QXmlStreamWriter::writeStartElement(const QString & qualifiedName);
extern void C_ZN16QXmlStreamWriter17writeStartElementERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeDefaultNamespace(const QString & namespaceUri);
extern void C_ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeStartDocument();
extern void C_ZN16QXmlStreamWriter18writeStartDocumentEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::writeStartDocument(const QString & version, bool standalone);
extern void C_ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QXmlStreamWriter::writeStartDocument(const QString & version);
extern void C_ZN16QXmlStreamWriter18writeStartDocumentERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeEndDocument();
extern void C_ZN16QXmlStreamWriter16writeEndDocumentEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::~QXmlStreamWriter();
extern void C_ZN16QXmlStreamWriterD2Ev(void* qthis); // 4
  // proto:  void QXmlStreamWriter::writeEndElement();
extern void C_ZN16QXmlStreamWriter15writeEndElementEv(void* qthis); // 4
  // proto:  void QXmlStreamWriter::writeCurrentToken(const QXmlStreamReader & reader);
extern void C_ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::setAutoFormattingIndent(int spacesOrTabs);
extern void C_ZN16QXmlStreamWriter23setAutoFormattingIndentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QXmlStreamWriter::writeCharacters(const QString & text);
extern void C_ZN16QXmlStreamWriter15writeCharactersERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QXmlStreamWriter::writeEmptyElement(const QString & namespaceUri, const QString & name);
extern void C_ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QXmlStreamWriter::writeEmptyElement(const QString & qualifiedName);
extern void C_ZN16QXmlStreamWriter17writeEmptyElementERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringRef QXmlStreamNotationDeclaration::name();
extern void C_ZNK29QXmlStreamNotationDeclaration4nameEv(void* qthis); // 2
  // proto:  void QXmlStreamNotationDeclaration::QXmlStreamNotationDeclaration(const QXmlStreamNotationDeclaration & );
extern void* C_ZN29QXmlStreamNotationDeclarationC2ERKS_(void* arg0); // 3
  // proto:  void QXmlStreamNotationDeclaration::QXmlStreamNotationDeclaration();
extern void* C_ZN29QXmlStreamNotationDeclarationC2Ev(); // 3
  // proto:  QStringRef QXmlStreamNotationDeclaration::systemId();
extern void C_ZNK29QXmlStreamNotationDeclaration8systemIdEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamNotationDeclaration::publicId();
extern void C_ZNK29QXmlStreamNotationDeclaration8publicIdEv(void* qthis); // 2
  // proto:  void QXmlStreamNotationDeclaration::~QXmlStreamNotationDeclaration();
extern void C_ZN29QXmlStreamNotationDeclarationD2Ev(void* qthis); // 4
  // proto:  QStringRef QXmlStreamAttribute::qualifiedName();
extern void C_ZNK19QXmlStreamAttribute13qualifiedNameEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamAttribute::name();
extern void C_ZNK19QXmlStreamAttribute4nameEv(void* qthis); // 2
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QString & qualifiedName, const QString & value);
extern void* C_ZN19QXmlStreamAttributeC2ERK7QStringS2_(void* arg0, void* arg1); // 3
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute();
extern void* C_ZN19QXmlStreamAttributeC2Ev(); // 3
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QXmlStreamAttribute & );
extern void* C_ZN19QXmlStreamAttributeC2ERKS_(void* arg0); // 3
  // proto:  void QXmlStreamAttribute::QXmlStreamAttribute(const QString & namespaceUri, const QString & name, const QString & value);
extern void* C_ZN19QXmlStreamAttributeC2ERK7QStringS2_S2_(void* arg0, void* arg1, void* arg2); // 3
  // proto:  QStringRef QXmlStreamAttribute::value();
extern void C_ZNK19QXmlStreamAttribute5valueEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamAttribute::namespaceUri();
extern void C_ZNK19QXmlStreamAttribute12namespaceUriEv(void* qthis); // 2
  // proto:  QStringRef QXmlStreamAttribute::prefix();
extern void C_ZNK19QXmlStreamAttribute6prefixEv(void* qthis); // 2
  // proto:  void QXmlStreamAttribute::~QXmlStreamAttribute();
extern void C_ZN19QXmlStreamAttributeD2Ev(void* qthis); // 4
  // proto:  bool QXmlStreamAttribute::isDefault();
extern bool C_ZNK19QXmlStreamAttribute9isDefaultEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QXmlStreamStringRef)=16
type QXmlStreamStringRef struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamReader)=1
type QXmlStreamReader struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamEntityResolver)=8
type QXmlStreamEntityResolver struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamNamespaceDeclaration)=40
type QXmlStreamNamespaceDeclaration struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamEntityDeclaration)=88
type QXmlStreamEntityDeclaration struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamAttributes)=1
type QXmlStreamAttributes struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamWriter)=1
type QXmlStreamWriter struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamNotationDeclaration)=56
type QXmlStreamNotationDeclaration struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QXmlStreamAttribute)=80
type QXmlStreamAttribute struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// string()
func (this *QXmlStreamStringRef) String(args ...interface{}) (ret interface{}) {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamStringRef6stringEv
    // invoke: const QString * string()
    var ret0 = C.C_ZNK19QXmlStreamStringRef6stringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "const QString *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "string", args)
  }

  return
}

// ~QXmlStreamStringRef()
func (this *QXmlStreamStringRef) Freeqxmlstreamstringref(args ...interface{}) () {
  // ~QXmlStreamStringRef()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QXmlStreamStringRefD0Ev
    // invoke: void ~QXmlStreamStringRef()
    C.C_ZN19QXmlStreamStringRefD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "~QXmlStreamStringRef", args)
  }

  return
}

// clear()
func (this *QXmlStreamStringRef) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QXmlStreamStringRef5clearEv
    // invoke: void clear()
    C.C_ZN19QXmlStreamStringRef5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "clear", args)
  }

  return
}

// QXmlStreamStringRef()
func NewQXmlStreamStringRef(args ...interface{}) *QXmlStreamStringRef {
  // QXmlStreamStringRef()
  // QXmlStreamStringRef(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QXmlStreamStringRefC1Ev
    // invoke: void QXmlStreamStringRef()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QXmlStreamStringRefC2Ev()
    return &QXmlStreamStringRef{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QXmlStreamStringRefC1ERK7QString
    // invoke: void QXmlStreamStringRef(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QXmlStreamStringRefC2ERK7QString(arg0)
    return &QXmlStreamStringRef{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "QXmlStreamStringRef", args)
  }

  return nil // QXmlStreamStringRef{Qclsinst:qthis}
}

// position()
func (this *QXmlStreamStringRef) Position(args ...interface{}) (ret interface{}) {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamStringRef8positionEv
    // invoke: int position()
    var ret0 = C.C_ZNK19QXmlStreamStringRef8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "position", args)
  }

  return
}

// size()
func (this *QXmlStreamStringRef) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamStringRef4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK19QXmlStreamStringRef4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamStringRef", "size", args)
  }

  return
}

// isEntityReference()
func (this *QXmlStreamReader) Isentityreference(args ...interface{}) (ret interface{}) {
  // isEntityReference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader17isEntityReferenceEv
    // invoke: bool isEntityReference()
    var ret0 = C.C_ZNK16QXmlStreamReader17isEntityReferenceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isEntityReference", args)
  }

  return
}

// isDTD()
func (this *QXmlStreamReader) Isdtd(args ...interface{}) (ret interface{}) {
  // isDTD()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader5isDTDEv
    // invoke: bool isDTD()
    var ret0 = C.C_ZNK16QXmlStreamReader5isDTDEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isDTD", args)
  }

  return
}

// qualifiedName()
func (this *QXmlStreamReader) Qualifiedname(args ...interface{}) () {
  // qualifiedName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader13qualifiedNameEv
    // invoke: QStringRef qualifiedName()
    C.C_ZNK16QXmlStreamReader13qualifiedNameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "qualifiedName", args)
  }

  return
}

// tokenType()
func (this *QXmlStreamReader) Tokentype(args ...interface{}) () {
  // tokenType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader9tokenTypeEv
    // invoke: QXmlStreamReader::TokenType tokenType()
    C.C_ZNK16QXmlStreamReader9tokenTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "tokenType", args)
  }

  return
}

// readNext()
func (this *QXmlStreamReader) Readnext(args ...interface{}) () {
  // readNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader8readNextEv
    // invoke: QXmlStreamReader::TokenType readNext()
    C.C_ZN16QXmlStreamReader8readNextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "readNext", args)
  }

  return
}

// text()
func (this *QXmlStreamReader) Text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader4textEv
    // invoke: QStringRef text()
    C.C_ZNK16QXmlStreamReader4textEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "text", args)
  }

  return
}

// dtdPublicId()
func (this *QXmlStreamReader) Dtdpublicid(args ...interface{}) () {
  // dtdPublicId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11dtdPublicIdEv
    // invoke: QStringRef dtdPublicId()
    C.C_ZNK16QXmlStreamReader11dtdPublicIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "dtdPublicId", args)
  }

  return
}

// entityDeclarations()
func (this *QXmlStreamReader) Entitydeclarations(args ...interface{}) () {
  // entityDeclarations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader18entityDeclarationsEv
    // invoke: QXmlStreamEntityDeclarations entityDeclarations()
    C.C_ZNK16QXmlStreamReader18entityDeclarationsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "entityDeclarations", args)
  }

  return
}

// setNamespaceProcessing(_Bool)
func (this *QXmlStreamReader) Setnamespaceprocessing(args ...interface{}) () {
  // setNamespaceProcessing(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader22setNamespaceProcessingEb
    // invoke: void setNamespaceProcessing(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader22setNamespaceProcessingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "setNamespaceProcessing", args)
  }

  return
}

// namespaceDeclarations()
func (this *QXmlStreamReader) Namespacedeclarations(args ...interface{}) () {
  // namespaceDeclarations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader21namespaceDeclarationsEv
    // invoke: QXmlStreamNamespaceDeclarations namespaceDeclarations()
    C.C_ZNK16QXmlStreamReader21namespaceDeclarationsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "namespaceDeclarations", args)
  }

  return
}

// isCharacters()
func (this *QXmlStreamReader) Ischaracters(args ...interface{}) (ret interface{}) {
  // isCharacters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12isCharactersEv
    // invoke: bool isCharacters()
    var ret0 = C.C_ZNK16QXmlStreamReader12isCharactersEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isCharacters", args)
  }

  return
}

// prefix()
func (this *QXmlStreamReader) Prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader6prefixEv
    // invoke: QStringRef prefix()
    C.C_ZNK16QXmlStreamReader6prefixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "prefix", args)
  }

  return
}

// addExtraNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
func (this *QXmlStreamReader) Addextranamespacedeclaration(args ...interface{}) () {
  // addExtraNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamNamespaceDeclaration{}) // "const QXmlStreamNamespaceDeclaration &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration
    // invoke: void addExtraNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
    var arg0 = args[0].(*QXmlStreamNamespaceDeclaration).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "addExtraNamespaceDeclaration", args)
  }

  return
}

// documentEncoding()
func (this *QXmlStreamReader) Documentencoding(args ...interface{}) () {
  // documentEncoding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader16documentEncodingEv
    // invoke: QStringRef documentEncoding()
    C.C_ZNK16QXmlStreamReader16documentEncodingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "documentEncoding", args)
  }

  return
}

// processingInstructionData()
func (this *QXmlStreamReader) Processinginstructiondata(args ...interface{}) () {
  // processingInstructionData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader25processingInstructionDataEv
    // invoke: QStringRef processingInstructionData()
    C.C_ZNK16QXmlStreamReader25processingInstructionDataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "processingInstructionData", args)
  }

  return
}

// processingInstructionTarget()
func (this *QXmlStreamReader) Processinginstructiontarget(args ...interface{}) () {
  // processingInstructionTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader27processingInstructionTargetEv
    // invoke: QStringRef processingInstructionTarget()
    C.C_ZNK16QXmlStreamReader27processingInstructionTargetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "processingInstructionTarget", args)
  }

  return
}

// setEntityResolver(class QXmlStreamEntityResolver *)
func (this *QXmlStreamReader) Setentityresolver(args ...interface{}) () {
  // setEntityResolver(class QXmlStreamEntityResolver *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamEntityResolver{}) // "QXmlStreamEntityResolver *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver
    // invoke: void setEntityResolver(class QXmlStreamEntityResolver *)
    var arg0 = args[0].(*QXmlStreamEntityResolver).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "setEntityResolver", args)
  }

  return
}

// isProcessingInstruction()
func (this *QXmlStreamReader) Isprocessinginstruction(args ...interface{}) (ret interface{}) {
  // isProcessingInstruction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader23isProcessingInstructionEv
    // invoke: bool isProcessingInstruction()
    var ret0 = C.C_ZNK16QXmlStreamReader23isProcessingInstructionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isProcessingInstruction", args)
  }

  return
}

// documentVersion()
func (this *QXmlStreamReader) Documentversion(args ...interface{}) () {
  // documentVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader15documentVersionEv
    // invoke: QStringRef documentVersion()
    C.C_ZNK16QXmlStreamReader15documentVersionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "documentVersion", args)
  }

  return
}

// notationDeclarations()
func (this *QXmlStreamReader) Notationdeclarations(args ...interface{}) () {
  // notationDeclarations()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader20notationDeclarationsEv
    // invoke: QXmlStreamNotationDeclarations notationDeclarations()
    C.C_ZNK16QXmlStreamReader20notationDeclarationsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "notationDeclarations", args)
  }

  return
}

// namespaceUri()
func (this *QXmlStreamReader) Namespaceuri(args ...interface{}) () {
  // namespaceUri()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12namespaceUriEv
    // invoke: QStringRef namespaceUri()
    C.C_ZNK16QXmlStreamReader12namespaceUriEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "namespaceUri", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QXmlStreamReader) Setdevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader9setDeviceEP9QIODevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "setDevice", args)
  }

  return
}

// raiseError(const class QString &)
func (this *QXmlStreamReader) Raiseerror(args ...interface{}) () {
  // raiseError(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader10raiseErrorERK7QString
    // invoke: void raiseError(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader10raiseErrorERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "raiseError", args)
  }

  return
}

// entityResolver()
func (this *QXmlStreamReader) Entityresolver(args ...interface{}) (ret interface{}) {
  // entityResolver()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader14entityResolverEv
    // invoke: QXmlStreamEntityResolver * entityResolver()
    var ret0 = C.C_ZNK16QXmlStreamReader14entityResolverEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QXmlStreamEntityResolver{}) // "QXmlStreamEntityResolver *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "entityResolver", args)
  }

  return
}

// isStartElement()
func (this *QXmlStreamReader) Isstartelement(args ...interface{}) (ret interface{}) {
  // isStartElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader14isStartElementEv
    // invoke: bool isStartElement()
    var ret0 = C.C_ZNK16QXmlStreamReader14isStartElementEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isStartElement", args)
  }

  return
}

// isWhitespace()
func (this *QXmlStreamReader) Iswhitespace(args ...interface{}) (ret interface{}) {
  // isWhitespace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12isWhitespaceEv
    // invoke: bool isWhitespace()
    var ret0 = C.C_ZNK16QXmlStreamReader12isWhitespaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isWhitespace", args)
  }

  return
}

// QXmlStreamReader()
func NewQXmlStreamReader(args ...interface{}) *QXmlStreamReader {
  // QXmlStreamReader()
  // QXmlStreamReader(const class QString &)
  // QXmlStreamReader(const char *)
  // QXmlStreamReader(const class QByteArray &)
  // QXmlStreamReader(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReaderC1Ev
    // invoke: void QXmlStreamReader()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamReaderC2Ev()
    return &QXmlStreamReader{Qclsinst:qthis}
  case 1:
    // invoke: _ZN16QXmlStreamReaderC1ERK7QString
    // invoke: void QXmlStreamReader(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamReaderC2ERK7QString(arg0)
    return &QXmlStreamReader{Qclsinst:qthis}
  case 2:
    // invoke: _ZN16QXmlStreamReaderC1EPKc
    // invoke: void QXmlStreamReader(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamReaderC2EPKc(arg0)
    return &QXmlStreamReader{Qclsinst:qthis}
  case 3:
    // invoke: _ZN16QXmlStreamReaderC1ERK10QByteArray
    // invoke: void QXmlStreamReader(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamReaderC2ERK10QByteArray(arg0)
    return &QXmlStreamReader{Qclsinst:qthis}
  case 4:
    // invoke: _ZN16QXmlStreamReaderC1EP9QIODevice
    // invoke: void QXmlStreamReader(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamReaderC2EP9QIODevice(arg0)
    return &QXmlStreamReader{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "QXmlStreamReader", args)
  }

  return nil // QXmlStreamReader{Qclsinst:qthis}
}

// hasError()
func (this *QXmlStreamReader) Haserror(args ...interface{}) (ret interface{}) {
  // hasError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader8hasErrorEv
    // invoke: bool hasError()
    var ret0 = C.C_ZNK16QXmlStreamReader8hasErrorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "hasError", args)
  }

  return
}

// characterOffset()
func (this *QXmlStreamReader) Characteroffset(args ...interface{}) (ret interface{}) {
  // characterOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader15characterOffsetEv
    // invoke: qint64 characterOffset()
    var ret0 = C.C_ZNK16QXmlStreamReader15characterOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "characterOffset", args)
  }

  return
}

// errorString()
func (this *QXmlStreamReader) Errorstring(args ...interface{}) (ret interface{}) {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11errorStringEv
    // invoke: QString errorString()
    var ret0 = C.C_ZNK16QXmlStreamReader11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "errorString", args)
  }

  return
}

// isEndElement()
func (this *QXmlStreamReader) Isendelement(args ...interface{}) (ret interface{}) {
  // isEndElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12isEndElementEv
    // invoke: bool isEndElement()
    var ret0 = C.C_ZNK16QXmlStreamReader12isEndElementEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isEndElement", args)
  }

  return
}

// isCDATA()
func (this *QXmlStreamReader) Iscdata(args ...interface{}) (ret interface{}) {
  // isCDATA()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader7isCDATAEv
    // invoke: bool isCDATA()
    var ret0 = C.C_ZNK16QXmlStreamReader7isCDATAEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isCDATA", args)
  }

  return
}

// isStandaloneDocument()
func (this *QXmlStreamReader) Isstandalonedocument(args ...interface{}) (ret interface{}) {
  // isStandaloneDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader20isStandaloneDocumentEv
    // invoke: bool isStandaloneDocument()
    var ret0 = C.C_ZNK16QXmlStreamReader20isStandaloneDocumentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isStandaloneDocument", args)
  }

  return
}

// skipCurrentElement()
func (this *QXmlStreamReader) Skipcurrentelement(args ...interface{}) () {
  // skipCurrentElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader18skipCurrentElementEv
    // invoke: void skipCurrentElement()
    C.C_ZN16QXmlStreamReader18skipCurrentElementEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "skipCurrentElement", args)
  }

  return
}

// isComment()
func (this *QXmlStreamReader) Iscomment(args ...interface{}) (ret interface{}) {
  // isComment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader9isCommentEv
    // invoke: bool isComment()
    var ret0 = C.C_ZNK16QXmlStreamReader9isCommentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isComment", args)
  }

  return
}

// lineNumber()
func (this *QXmlStreamReader) Linenumber(args ...interface{}) (ret interface{}) {
  // lineNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader10lineNumberEv
    // invoke: qint64 lineNumber()
    var ret0 = C.C_ZNK16QXmlStreamReader10lineNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "lineNumber", args)
  }

  return
}

// device()
func (this *QXmlStreamReader) Device(args ...interface{}) (ret interface{}) {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader6deviceEv
    // invoke: QIODevice * device()
    var ret0 = C.C_ZNK16QXmlStreamReader6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "device", args)
  }

  return
}

// isEndDocument()
func (this *QXmlStreamReader) Isenddocument(args ...interface{}) (ret interface{}) {
  // isEndDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader13isEndDocumentEv
    // invoke: bool isEndDocument()
    var ret0 = C.C_ZNK16QXmlStreamReader13isEndDocumentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isEndDocument", args)
  }

  return
}

// columnNumber()
func (this *QXmlStreamReader) Columnnumber(args ...interface{}) (ret interface{}) {
  // columnNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader12columnNumberEv
    // invoke: qint64 columnNumber()
    var ret0 = C.C_ZNK16QXmlStreamReader12columnNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "columnNumber", args)
  }

  return
}

// tokenString()
func (this *QXmlStreamReader) Tokenstring(args ...interface{}) (ret interface{}) {
  // tokenString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11tokenStringEv
    // invoke: QString tokenString()
    var ret0 = C.C_ZNK16QXmlStreamReader11tokenStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "tokenString", args)
  }

  return
}

// name()
func (this *QXmlStreamReader) Name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader4nameEv
    // invoke: QStringRef name()
    C.C_ZNK16QXmlStreamReader4nameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "name", args)
  }

  return
}

// namespaceProcessing()
func (this *QXmlStreamReader) Namespaceprocessing(args ...interface{}) (ret interface{}) {
  // namespaceProcessing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader19namespaceProcessingEv
    // invoke: bool namespaceProcessing()
    var ret0 = C.C_ZNK16QXmlStreamReader19namespaceProcessingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "namespaceProcessing", args)
  }

  return
}

// isStartDocument()
func (this *QXmlStreamReader) Isstartdocument(args ...interface{}) (ret interface{}) {
  // isStartDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader15isStartDocumentEv
    // invoke: bool isStartDocument()
    var ret0 = C.C_ZNK16QXmlStreamReader15isStartDocumentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "isStartDocument", args)
  }

  return
}

// clear()
func (this *QXmlStreamReader) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader5clearEv
    // invoke: void clear()
    C.C_ZN16QXmlStreamReader5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "clear", args)
  }

  return
}

// atEnd()
func (this *QXmlStreamReader) Atend(args ...interface{}) (ret interface{}) {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader5atEndEv
    // invoke: bool atEnd()
    var ret0 = C.C_ZNK16QXmlStreamReader5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "atEnd", args)
  }

  return
}

// dtdName()
func (this *QXmlStreamReader) Dtdname(args ...interface{}) () {
  // dtdName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader7dtdNameEv
    // invoke: QStringRef dtdName()
    C.C_ZNK16QXmlStreamReader7dtdNameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "dtdName", args)
  }

  return
}

// addData(const class QByteArray &)
func (this *QXmlStreamReader) Adddata(args ...interface{}) () {
  // addData(const class QByteArray &)
  // addData(const class QString &)
  // addData(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader7addDataERK10QByteArray
    // invoke: void addData(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader7addDataERK10QByteArray(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN16QXmlStreamReader7addDataERK7QString
    // invoke: void addData(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamReader7addDataERK7QString(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN16QXmlStreamReader7addDataEPKc
    // invoke: void addData(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN16QXmlStreamReader7addDataEPKc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "addData", args)
  }

  return
}

// dtdSystemId()
func (this *QXmlStreamReader) Dtdsystemid(args ...interface{}) () {
  // dtdSystemId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader11dtdSystemIdEv
    // invoke: QStringRef dtdSystemId()
    C.C_ZNK16QXmlStreamReader11dtdSystemIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "dtdSystemId", args)
  }

  return
}

// readNextStartElement()
func (this *QXmlStreamReader) Readnextstartelement(args ...interface{}) (ret interface{}) {
  // readNextStartElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReader20readNextStartElementEv
    // invoke: bool readNextStartElement()
    var ret0 = C.C_ZN16QXmlStreamReader20readNextStartElementEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "readNextStartElement", args)
  }

  return
}

// error()
func (this *QXmlStreamReader) Error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader5errorEv
    // invoke: QXmlStreamReader::Error error()
    C.C_ZNK16QXmlStreamReader5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "error", args)
  }

  return
}

// attributes()
func (this *QXmlStreamReader) Attributes(args ...interface{}) (ret interface{}) {
  // attributes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamReader10attributesEv
    // invoke: QXmlStreamAttributes attributes()
    var ret0 = C.C_ZNK16QXmlStreamReader10attributesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QXmlStreamAttributes{}) // "QXmlStreamAttributes"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "attributes", args)
  }

  return
}

// ~QXmlStreamReader()
func (this *QXmlStreamReader) Freeqxmlstreamreader(args ...interface{}) () {
  // ~QXmlStreamReader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamReaderD0Ev
    // invoke: void ~QXmlStreamReader()
    C.C_ZN16QXmlStreamReaderD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamReader", "~QXmlStreamReader", args)
  }

  return
}

// resolveEntity(const class QString &, const class QString &)
func (this *QXmlStreamEntityResolver) Resolveentity(args ...interface{}) (ret interface{}) {
  // resolveEntity(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_
    // invoke: QString resolveEntity(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN24QXmlStreamEntityResolver13resolveEntityERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamEntityResolver", "resolveEntity", args)
  }

  return
}

// resolveUndeclaredEntity(const class QString &)
func (this *QXmlStreamEntityResolver) Resolveundeclaredentity(args ...interface{}) (ret interface{}) {
  // resolveUndeclaredEntity(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString
    // invoke: QString resolveUndeclaredEntity(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN24QXmlStreamEntityResolver23resolveUndeclaredEntityERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamEntityResolver", "resolveUndeclaredEntity", args)
  }

  return
}

// ~QXmlStreamEntityResolver()
func (this *QXmlStreamEntityResolver) Freeqxmlstreamentityresolver(args ...interface{}) () {
  // ~QXmlStreamEntityResolver()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QXmlStreamEntityResolverD0Ev
    // invoke: void ~QXmlStreamEntityResolver()
    C.C_ZN24QXmlStreamEntityResolverD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityResolver", "~QXmlStreamEntityResolver", args)
  }

  return
}

// QXmlStreamNamespaceDeclaration()
func NewQXmlStreamNamespaceDeclaration(args ...interface{}) *QXmlStreamNamespaceDeclaration {
  // QXmlStreamNamespaceDeclaration()
  // QXmlStreamNamespaceDeclaration(const class QString &, const class QString &)
  // QXmlStreamNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QXmlStreamNamespaceDeclaration{}) // "const QXmlStreamNamespaceDeclaration &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QXmlStreamNamespaceDeclarationC1Ev
    // invoke: void QXmlStreamNamespaceDeclaration()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QXmlStreamNamespaceDeclarationC2Ev()
    return &QXmlStreamNamespaceDeclaration{Qclsinst:qthis}
  case 1:
    // invoke: _ZN30QXmlStreamNamespaceDeclarationC1ERK7QStringS2_
    // invoke: void QXmlStreamNamespaceDeclaration(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QXmlStreamNamespaceDeclarationC2ERK7QStringS2_(arg0, arg1)
    return &QXmlStreamNamespaceDeclaration{Qclsinst:qthis}
  case 2:
    // invoke: _ZN30QXmlStreamNamespaceDeclarationC1ERKS_
    // invoke: void QXmlStreamNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
    var arg0 = args[0].(*QXmlStreamNamespaceDeclaration).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QXmlStreamNamespaceDeclarationC2ERKS_(arg0)
    return &QXmlStreamNamespaceDeclaration{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "QXmlStreamNamespaceDeclaration", args)
  }

  return nil // QXmlStreamNamespaceDeclaration{Qclsinst:qthis}
}

// ~QXmlStreamNamespaceDeclaration()
func (this *QXmlStreamNamespaceDeclaration) Freeqxmlstreamnamespacedeclaration(args ...interface{}) () {
  // ~QXmlStreamNamespaceDeclaration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QXmlStreamNamespaceDeclarationD0Ev
    // invoke: void ~QXmlStreamNamespaceDeclaration()
    C.C_ZN30QXmlStreamNamespaceDeclarationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "~QXmlStreamNamespaceDeclaration", args)
  }

  return
}

// namespaceUri()
func (this *QXmlStreamNamespaceDeclaration) Namespaceuri(args ...interface{}) () {
  // namespaceUri()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv
    // invoke: QStringRef namespaceUri()
    C.C_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "namespaceUri", args)
  }

  return
}

// prefix()
func (this *QXmlStreamNamespaceDeclaration) Prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QXmlStreamNamespaceDeclaration6prefixEv
    // invoke: QStringRef prefix()
    C.C_ZNK30QXmlStreamNamespaceDeclaration6prefixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNamespaceDeclaration", "prefix", args)
  }

  return
}

// ~QXmlStreamEntityDeclaration()
func (this *QXmlStreamEntityDeclaration) Freeqxmlstreamentitydeclaration(args ...interface{}) () {
  // ~QXmlStreamEntityDeclaration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QXmlStreamEntityDeclarationD0Ev
    // invoke: void ~QXmlStreamEntityDeclaration()
    C.C_ZN27QXmlStreamEntityDeclarationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "~QXmlStreamEntityDeclaration", args)
  }

  return
}

// publicId()
func (this *QXmlStreamEntityDeclaration) Publicid(args ...interface{}) () {
  // publicId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration8publicIdEv
    // invoke: QStringRef publicId()
    C.C_ZNK27QXmlStreamEntityDeclaration8publicIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "publicId", args)
  }

  return
}

// name()
func (this *QXmlStreamEntityDeclaration) Name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration4nameEv
    // invoke: QStringRef name()
    C.C_ZNK27QXmlStreamEntityDeclaration4nameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "name", args)
  }

  return
}

// value()
func (this *QXmlStreamEntityDeclaration) Value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration5valueEv
    // invoke: QStringRef value()
    C.C_ZNK27QXmlStreamEntityDeclaration5valueEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "value", args)
  }

  return
}

// systemId()
func (this *QXmlStreamEntityDeclaration) Systemid(args ...interface{}) () {
  // systemId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration8systemIdEv
    // invoke: QStringRef systemId()
    C.C_ZNK27QXmlStreamEntityDeclaration8systemIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "systemId", args)
  }

  return
}

// notationName()
func (this *QXmlStreamEntityDeclaration) Notationname(args ...interface{}) () {
  // notationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QXmlStreamEntityDeclaration12notationNameEv
    // invoke: QStringRef notationName()
    C.C_ZNK27QXmlStreamEntityDeclaration12notationNameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "notationName", args)
  }

  return
}

// QXmlStreamEntityDeclaration(const class QXmlStreamEntityDeclaration &)
func NewQXmlStreamEntityDeclaration(args ...interface{}) *QXmlStreamEntityDeclaration {
  // QXmlStreamEntityDeclaration(const class QXmlStreamEntityDeclaration &)
  // QXmlStreamEntityDeclaration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamEntityDeclaration{}) // "const QXmlStreamEntityDeclaration &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QXmlStreamEntityDeclarationC1ERKS_
    // invoke: void QXmlStreamEntityDeclaration(const class QXmlStreamEntityDeclaration &)
    var arg0 = args[0].(*QXmlStreamEntityDeclaration).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QXmlStreamEntityDeclarationC2ERKS_(arg0)
    return &QXmlStreamEntityDeclaration{Qclsinst:qthis}
  case 1:
    // invoke: _ZN27QXmlStreamEntityDeclarationC1Ev
    // invoke: void QXmlStreamEntityDeclaration()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QXmlStreamEntityDeclarationC2Ev()
    return &QXmlStreamEntityDeclaration{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamEntityDeclaration", "QXmlStreamEntityDeclaration", args)
  }

  return nil // QXmlStreamEntityDeclaration{Qclsinst:qthis}
}

// hasAttribute(const class QString &, const class QString &)
func (this *QXmlStreamAttributes) Hasattribute(args ...interface{}) (ret interface{}) {
  // hasAttribute(const class QString &, const class QString &)
  // hasAttribute(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_
    // invoke: bool hasAttribute(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK20QXmlStreamAttributes12hasAttributeERK7QString
    // invoke: bool hasAttribute(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QXmlStreamAttributes12hasAttributeERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "hasAttribute", args)
  }

  return
}

// append(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamAttributes) Append(args ...interface{}) () {
  // append(const class QString &, const class QString &, const class QString &)
  // append(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QXmlStreamAttributes6appendERK7QStringS2_S2_
    // invoke: void append(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN20QXmlStreamAttributes6appendERK7QStringS2_S2_(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN20QXmlStreamAttributes6appendERK7QStringS2_
    // invoke: void append(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QXmlStreamAttributes6appendERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "append", args)
  }

  return
}

// value(const class QString &, const class QString &)
func (this *QXmlStreamAttributes) Value(args ...interface{}) () {
  // value(const class QString &, const class QString &)
  // value(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QXmlStreamAttributes5valueERK7QStringS2_
    // invoke: QStringRef value(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK20QXmlStreamAttributes5valueERK7QStringS2_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK20QXmlStreamAttributes5valueERK7QString
    // invoke: QStringRef value(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK20QXmlStreamAttributes5valueERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "value", args)
  }

  return
}

// QXmlStreamAttributes()
func NewQXmlStreamAttributes(args ...interface{}) *QXmlStreamAttributes {
  // QXmlStreamAttributes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QXmlStreamAttributesC1Ev
    // invoke: void QXmlStreamAttributes()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QXmlStreamAttributesC2Ev()
    return &QXmlStreamAttributes{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamAttributes", "QXmlStreamAttributes", args)
  }

  return nil // QXmlStreamAttributes{Qclsinst:qthis}
}

// writeTextElement(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamWriter) Writetextelement(args ...interface{}) () {
  // writeTextElement(const class QString &, const class QString &, const class QString &)
  // writeTextElement(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_
    // invoke: void writeTextElement(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_S2_(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_
    // invoke: void writeTextElement(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter16writeTextElementERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeTextElement", args)
  }

  return
}

// writeNamespace(const class QString &, const class QString &)
func (this *QXmlStreamWriter) Writenamespace(args ...interface{}) () {
  // writeNamespace(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_
    // invoke: void writeNamespace(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter14writeNamespaceERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeNamespace", args)
  }

  return
}

// autoFormatting()
func (this *QXmlStreamWriter) Autoformatting(args ...interface{}) (ret interface{}) {
  // autoFormatting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter14autoFormattingEv
    // invoke: bool autoFormatting()
    var ret0 = C.C_ZNK16QXmlStreamWriter14autoFormattingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "autoFormatting", args)
  }

  return
}

// device()
func (this *QXmlStreamWriter) Device(args ...interface{}) (ret interface{}) {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter6deviceEv
    // invoke: QIODevice * device()
    var ret0 = C.C_ZNK16QXmlStreamWriter6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "device", args)
  }

  return
}

// writeEntityReference(const class QString &)
func (this *QXmlStreamWriter) Writeentityreference(args ...interface{}) () {
  // writeEntityReference(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter20writeEntityReferenceERK7QString
    // invoke: void writeEntityReference(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter20writeEntityReferenceERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEntityReference", args)
  }

  return
}

// writeProcessingInstruction(const class QString &, const class QString &)
func (this *QXmlStreamWriter) Writeprocessinginstruction(args ...interface{}) () {
  // writeProcessingInstruction(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_
    // invoke: void writeProcessingInstruction(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter26writeProcessingInstructionERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeProcessingInstruction", args)
  }

  return
}

// writeAttributes(const class QXmlStreamAttributes &)
func (this *QXmlStreamWriter) Writeattributes(args ...interface{}) () {
  // writeAttributes(const class QXmlStreamAttributes &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamAttributes{}) // "const QXmlStreamAttributes &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes
    // invoke: void writeAttributes(const class QXmlStreamAttributes &)
    var arg0 = args[0].(*QXmlStreamAttributes).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter15writeAttributesERK20QXmlStreamAttributes(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeAttributes", args)
  }

  return
}

// autoFormattingIndent()
func (this *QXmlStreamWriter) Autoformattingindent(args ...interface{}) (ret interface{}) {
  // autoFormattingIndent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter20autoFormattingIndentEv
    // invoke: int autoFormattingIndent()
    var ret0 = C.C_ZNK16QXmlStreamWriter20autoFormattingIndentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "autoFormattingIndent", args)
  }

  return
}

// setCodec(const char *)
func (this *QXmlStreamWriter) Setcodec(args ...interface{}) () {
  // setCodec(const char *)
  // setCodec(class QTextCodec *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter8setCodecEPKc
    // invoke: void setCodec(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN16QXmlStreamWriter8setCodecEPKc(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN16QXmlStreamWriter8setCodecEP10QTextCodec
    // invoke: void setCodec(class QTextCodec *)
    var arg0 = args[0].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter8setCodecEP10QTextCodec(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setCodec", args)
  }

  return
}

// codec()
func (this *QXmlStreamWriter) Codec(args ...interface{}) (ret interface{}) {
  // codec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter5codecEv
    // invoke: QTextCodec * codec()
    var ret0 = C.C_ZNK16QXmlStreamWriter5codecEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "codec", args)
  }

  return
}

// writeComment(const class QString &)
func (this *QXmlStreamWriter) Writecomment(args ...interface{}) () {
  // writeComment(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter12writeCommentERK7QString
    // invoke: void writeComment(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter12writeCommentERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeComment", args)
  }

  return
}

// writeCDATA(const class QString &)
func (this *QXmlStreamWriter) Writecdata(args ...interface{}) () {
  // writeCDATA(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter10writeCDATAERK7QString
    // invoke: void writeCDATA(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter10writeCDATAERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeCDATA", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QXmlStreamWriter) Setdevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter9setDeviceEP9QIODevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setDevice", args)
  }

  return
}

// writeAttribute(const class QString &, const class QString &, const class QString &)
func (this *QXmlStreamWriter) Writeattribute(args ...interface{}) () {
  // writeAttribute(const class QString &, const class QString &, const class QString &)
  // writeAttribute(const class QString &, const class QString &)
  // writeAttribute(const class QXmlStreamAttribute &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QXmlStreamAttribute{}) // "const QXmlStreamAttribute &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_
    // invoke: void writeAttribute(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_S2_(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_
    // invoke: void writeAttribute(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter14writeAttributeERK7QStringS2_(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute
    // invoke: void writeAttribute(const class QXmlStreamAttribute &)
    var arg0 = args[0].(*QXmlStreamAttribute).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter14writeAttributeERK19QXmlStreamAttribute(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeAttribute", args)
  }

  return
}

// setAutoFormatting(_Bool)
func (this *QXmlStreamWriter) Setautoformatting(args ...interface{}) () {
  // setAutoFormatting(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17setAutoFormattingEb
    // invoke: void setAutoFormatting(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter17setAutoFormattingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setAutoFormatting", args)
  }

  return
}

// hasError()
func (this *QXmlStreamWriter) Haserror(args ...interface{}) (ret interface{}) {
  // hasError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QXmlStreamWriter8hasErrorEv
    // invoke: bool hasError()
    var ret0 = C.C_ZNK16QXmlStreamWriter8hasErrorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "hasError", args)
  }

  return
}

// writeDTD(const class QString &)
func (this *QXmlStreamWriter) Writedtd(args ...interface{}) () {
  // writeDTD(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter8writeDTDERK7QString
    // invoke: void writeDTD(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter8writeDTDERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeDTD", args)
  }

  return
}

// QXmlStreamWriter(class QByteArray *)
func NewQXmlStreamWriter(args ...interface{}) *QXmlStreamWriter {
  // QXmlStreamWriter(class QByteArray *)
  // QXmlStreamWriter(class QString *)
  // QXmlStreamWriter()
  // QXmlStreamWriter(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriterC1EP10QByteArray
    // invoke: void QXmlStreamWriter(class QByteArray *)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamWriterC2EP10QByteArray(arg0)
    return &QXmlStreamWriter{Qclsinst:qthis}
  case 1:
    // invoke: _ZN16QXmlStreamWriterC1EP7QString
    // invoke: void QXmlStreamWriter(class QString *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamWriterC2EP7QString(arg0)
    return &QXmlStreamWriter{Qclsinst:qthis}
  case 2:
    // invoke: _ZN16QXmlStreamWriterC1Ev
    // invoke: void QXmlStreamWriter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamWriterC2Ev()
    return &QXmlStreamWriter{Qclsinst:qthis}
  case 3:
    // invoke: _ZN16QXmlStreamWriterC1EP9QIODevice
    // invoke: void QXmlStreamWriter(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QXmlStreamWriterC2EP9QIODevice(arg0)
    return &QXmlStreamWriter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "QXmlStreamWriter", args)
  }

  return nil // QXmlStreamWriter{Qclsinst:qthis}
}

// writeStartElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) Writestartelement(args ...interface{}) () {
  // writeStartElement(const class QString &, const class QString &)
  // writeStartElement(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_
    // invoke: void writeStartElement(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter17writeStartElementERK7QStringS2_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN16QXmlStreamWriter17writeStartElementERK7QString
    // invoke: void writeStartElement(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter17writeStartElementERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeStartElement", args)
  }

  return
}

// writeDefaultNamespace(const class QString &)
func (this *QXmlStreamWriter) Writedefaultnamespace(args ...interface{}) () {
  // writeDefaultNamespace(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString
    // invoke: void writeDefaultNamespace(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter21writeDefaultNamespaceERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeDefaultNamespace", args)
  }

  return
}

// writeStartDocument()
func (this *QXmlStreamWriter) Writestartdocument(args ...interface{}) () {
  // writeStartDocument()
  // writeStartDocument(const class QString &, _Bool)
  // writeStartDocument(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter18writeStartDocumentEv
    // invoke: void writeStartDocument()
    C.C_ZN16QXmlStreamWriter18writeStartDocumentEv(this.Qclsinst)
  case 1:
    // invoke: _ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb
    // invoke: void writeStartDocument(const class QString &, _Bool)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter18writeStartDocumentERK7QStringb(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN16QXmlStreamWriter18writeStartDocumentERK7QString
    // invoke: void writeStartDocument(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter18writeStartDocumentERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeStartDocument", args)
  }

  return
}

// writeEndDocument()
func (this *QXmlStreamWriter) Writeenddocument(args ...interface{}) () {
  // writeEndDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter16writeEndDocumentEv
    // invoke: void writeEndDocument()
    C.C_ZN16QXmlStreamWriter16writeEndDocumentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEndDocument", args)
  }

  return
}

// ~QXmlStreamWriter()
func (this *QXmlStreamWriter) Freeqxmlstreamwriter(args ...interface{}) () {
  // ~QXmlStreamWriter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriterD0Ev
    // invoke: void ~QXmlStreamWriter()
    C.C_ZN16QXmlStreamWriterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "~QXmlStreamWriter", args)
  }

  return
}

// writeEndElement()
func (this *QXmlStreamWriter) Writeendelement(args ...interface{}) () {
  // writeEndElement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter15writeEndElementEv
    // invoke: void writeEndElement()
    C.C_ZN16QXmlStreamWriter15writeEndElementEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEndElement", args)
  }

  return
}

// writeCurrentToken(const class QXmlStreamReader &)
func (this *QXmlStreamWriter) Writecurrenttoken(args ...interface{}) () {
  // writeCurrentToken(const class QXmlStreamReader &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamReader{}) // "const QXmlStreamReader &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader
    // invoke: void writeCurrentToken(const class QXmlStreamReader &)
    var arg0 = args[0].(*QXmlStreamReader).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter17writeCurrentTokenERK16QXmlStreamReader(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeCurrentToken", args)
  }

  return
}

// setAutoFormattingIndent(int)
func (this *QXmlStreamWriter) Setautoformattingindent(args ...interface{}) () {
  // setAutoFormattingIndent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter23setAutoFormattingIndentEi
    // invoke: void setAutoFormattingIndent(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter23setAutoFormattingIndentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "setAutoFormattingIndent", args)
  }

  return
}

// writeCharacters(const class QString &)
func (this *QXmlStreamWriter) Writecharacters(args ...interface{}) () {
  // writeCharacters(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter15writeCharactersERK7QString
    // invoke: void writeCharacters(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter15writeCharactersERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeCharacters", args)
  }

  return
}

// writeEmptyElement(const class QString &, const class QString &)
func (this *QXmlStreamWriter) Writeemptyelement(args ...interface{}) () {
  // writeEmptyElement(const class QString &, const class QString &)
  // writeEmptyElement(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_
    // invoke: void writeEmptyElement(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QXmlStreamWriter17writeEmptyElementERK7QStringS2_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN16QXmlStreamWriter17writeEmptyElementERK7QString
    // invoke: void writeEmptyElement(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QXmlStreamWriter17writeEmptyElementERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QXmlStreamWriter", "writeEmptyElement", args)
  }

  return
}

// name()
func (this *QXmlStreamNotationDeclaration) Name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QXmlStreamNotationDeclaration4nameEv
    // invoke: QStringRef name()
    C.C_ZNK29QXmlStreamNotationDeclaration4nameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "name", args)
  }

  return
}

// QXmlStreamNotationDeclaration(const class QXmlStreamNotationDeclaration &)
func NewQXmlStreamNotationDeclaration(args ...interface{}) *QXmlStreamNotationDeclaration {
  // QXmlStreamNotationDeclaration(const class QXmlStreamNotationDeclaration &)
  // QXmlStreamNotationDeclaration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QXmlStreamNotationDeclaration{}) // "const QXmlStreamNotationDeclaration &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QXmlStreamNotationDeclarationC1ERKS_
    // invoke: void QXmlStreamNotationDeclaration(const class QXmlStreamNotationDeclaration &)
    var arg0 = args[0].(*QXmlStreamNotationDeclaration).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN29QXmlStreamNotationDeclarationC2ERKS_(arg0)
    return &QXmlStreamNotationDeclaration{Qclsinst:qthis}
  case 1:
    // invoke: _ZN29QXmlStreamNotationDeclarationC1Ev
    // invoke: void QXmlStreamNotationDeclaration()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN29QXmlStreamNotationDeclarationC2Ev()
    return &QXmlStreamNotationDeclaration{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "QXmlStreamNotationDeclaration", args)
  }

  return nil // QXmlStreamNotationDeclaration{Qclsinst:qthis}
}

// systemId()
func (this *QXmlStreamNotationDeclaration) Systemid(args ...interface{}) () {
  // systemId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QXmlStreamNotationDeclaration8systemIdEv
    // invoke: QStringRef systemId()
    C.C_ZNK29QXmlStreamNotationDeclaration8systemIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "systemId", args)
  }

  return
}

// publicId()
func (this *QXmlStreamNotationDeclaration) Publicid(args ...interface{}) () {
  // publicId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QXmlStreamNotationDeclaration8publicIdEv
    // invoke: QStringRef publicId()
    C.C_ZNK29QXmlStreamNotationDeclaration8publicIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "publicId", args)
  }

  return
}

// ~QXmlStreamNotationDeclaration()
func (this *QXmlStreamNotationDeclaration) Freeqxmlstreamnotationdeclaration(args ...interface{}) () {
  // ~QXmlStreamNotationDeclaration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QXmlStreamNotationDeclarationD0Ev
    // invoke: void ~QXmlStreamNotationDeclaration()
    C.C_ZN29QXmlStreamNotationDeclarationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamNotationDeclaration", "~QXmlStreamNotationDeclaration", args)
  }

  return
}

// qualifiedName()
func (this *QXmlStreamAttribute) Qualifiedname(args ...interface{}) () {
  // qualifiedName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute13qualifiedNameEv
    // invoke: QStringRef qualifiedName()
    C.C_ZNK19QXmlStreamAttribute13qualifiedNameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "qualifiedName", args)
  }

  return
}

// name()
func (this *QXmlStreamAttribute) Name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute4nameEv
    // invoke: QStringRef name()
    C.C_ZNK19QXmlStreamAttribute4nameEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "name", args)
  }

  return
}

// QXmlStreamAttribute(const class QString &, const class QString &)
func NewQXmlStreamAttribute(args ...interface{}) *QXmlStreamAttribute {
  // QXmlStreamAttribute(const class QString &, const class QString &)
  // QXmlStreamAttribute()
  // QXmlStreamAttribute(const class QXmlStreamAttribute &)
  // QXmlStreamAttribute(const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QXmlStreamAttribute{}) // "const QXmlStreamAttribute &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QXmlStreamAttributeC1ERK7QStringS2_
    // invoke: void QXmlStreamAttribute(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QXmlStreamAttributeC2ERK7QStringS2_(arg0, arg1)
    return &QXmlStreamAttribute{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QXmlStreamAttributeC1Ev
    // invoke: void QXmlStreamAttribute()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QXmlStreamAttributeC2Ev()
    return &QXmlStreamAttribute{Qclsinst:qthis}
  case 2:
    // invoke: _ZN19QXmlStreamAttributeC1ERKS_
    // invoke: void QXmlStreamAttribute(const class QXmlStreamAttribute &)
    var arg0 = args[0].(*QXmlStreamAttribute).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QXmlStreamAttributeC2ERKS_(arg0)
    return &QXmlStreamAttribute{Qclsinst:qthis}
  case 3:
    // invoke: _ZN19QXmlStreamAttributeC1ERK7QStringS2_S2_
    // invoke: void QXmlStreamAttribute(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QXmlStreamAttributeC2ERK7QStringS2_S2_(arg0, arg1, arg2)
    return &QXmlStreamAttribute{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "QXmlStreamAttribute", args)
  }

  return nil // QXmlStreamAttribute{Qclsinst:qthis}
}

// value()
func (this *QXmlStreamAttribute) Value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute5valueEv
    // invoke: QStringRef value()
    C.C_ZNK19QXmlStreamAttribute5valueEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "value", args)
  }

  return
}

// namespaceUri()
func (this *QXmlStreamAttribute) Namespaceuri(args ...interface{}) () {
  // namespaceUri()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute12namespaceUriEv
    // invoke: QStringRef namespaceUri()
    C.C_ZNK19QXmlStreamAttribute12namespaceUriEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "namespaceUri", args)
  }

  return
}

// prefix()
func (this *QXmlStreamAttribute) Prefix(args ...interface{}) () {
  // prefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute6prefixEv
    // invoke: QStringRef prefix()
    C.C_ZNK19QXmlStreamAttribute6prefixEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "prefix", args)
  }

  return
}

// ~QXmlStreamAttribute()
func (this *QXmlStreamAttribute) Freeqxmlstreamattribute(args ...interface{}) () {
  // ~QXmlStreamAttribute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QXmlStreamAttributeD0Ev
    // invoke: void ~QXmlStreamAttribute()
    C.C_ZN19QXmlStreamAttributeD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "~QXmlStreamAttribute", args)
  }

  return
}

// isDefault()
func (this *QXmlStreamAttribute) Isdefault(args ...interface{}) (ret interface{}) {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QXmlStreamAttribute9isDefaultEv
    // invoke: bool isDefault()
    var ret0 = C.C_ZNK19QXmlStreamAttribute9isDefaultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QXmlStreamAttribute", "isDefault", args)
  }

  return
}

// <= body block end

