package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamReader) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQXmlStreamReaderFromPointer(cthis unsafe.Pointer) *QXmlStreamReader {
	return &QXmlStreamReader{&qtrt.CObject{cthis}}
}
func (*QXmlStreamReader) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamReader {
	return NewQXmlStreamReaderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:360
// index:0
// Public
// void QXmlStreamReader()
func NewQXmlStreamReader() *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:361
// index:1
// Public
// void QXmlStreamReader(class QIODevice *)
func NewQXmlStreamReader_1(device *QIODevice /*444 QIODevice **/) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:362
// index:2
// Public
// void QXmlStreamReader(const class QByteArray &)
func NewQXmlStreamReader_2(data *QByteArray) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:363
// index:3
// Public
// void QXmlStreamReader(const class QString &)
func NewQXmlStreamReader_3(data *QString) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:364
// index:4
// Public
// void QXmlStreamReader(const char *)
func NewQXmlStreamReader_4(data string) *QXmlStreamReader {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:365
// index:0
// Public
// void ~QXmlStreamReader()
func DeleteQXmlStreamReader(*QXmlStreamReader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:367
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QXmlStreamReader) SetDevice(device *QIODevice /*444 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:368
// index:0
// Public
// QIODevice * device()
func (this *QXmlStreamReader) Device() *QIODevice /*444 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:369
// index:0
// Public
// void addData(const class QByteArray &)
func (this *QXmlStreamReader) AddData(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:370
// index:1
// Public
// void addData(const class QString &)
func (this *QXmlStreamReader) AddData_1(data *QString) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:371
// index:2
// Public
// void addData(const char *)
func (this *QXmlStreamReader) AddData_2(data string) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:372
// index:0
// Public
// void clear()
func (this *QXmlStreamReader) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:375
// index:0
// Public
// bool atEnd()
func (this *QXmlStreamReader) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:376
// index:0
// Public
// QXmlStreamReader::TokenType readNext()
func (this *QXmlStreamReader) ReadNext() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader8readNextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:378
// index:0
// Public
// bool readNextStartElement()
func (this *QXmlStreamReader) ReadNextStartElement() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader20readNextStartElementEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:379
// index:0
// Public
// void skipCurrentElement()
func (this *QXmlStreamReader) SkipCurrentElement() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader18skipCurrentElementEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:381
// index:0
// Public
// QXmlStreamReader::TokenType tokenType()
func (this *QXmlStreamReader) TokenType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader9tokenTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:382
// index:0
// Public
// QString tokenString()
func (this *QXmlStreamReader) TokenString() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11tokenStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:384
// index:0
// Public
// void setNamespaceProcessing(_Bool)
func (this *QXmlStreamReader) SetNamespaceProcessing(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader22setNamespaceProcessingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:385
// index:0
// Public
// bool namespaceProcessing()
func (this *QXmlStreamReader) NamespaceProcessing() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader19namespaceProcessingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:387
// index:0
// Public inline
// bool isStartDocument()
func (this *QXmlStreamReader) IsStartDocument() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader15isStartDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:388
// index:0
// Public inline
// bool isEndDocument()
func (this *QXmlStreamReader) IsEndDocument() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader13isEndDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:389
// index:0
// Public inline
// bool isStartElement()
func (this *QXmlStreamReader) IsStartElement() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader14isStartElementEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:390
// index:0
// Public inline
// bool isEndElement()
func (this *QXmlStreamReader) IsEndElement() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isEndElementEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:391
// index:0
// Public inline
// bool isCharacters()
func (this *QXmlStreamReader) IsCharacters() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isCharactersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:392
// index:0
// Public
// bool isWhitespace()
func (this *QXmlStreamReader) IsWhitespace() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isWhitespaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:393
// index:0
// Public
// bool isCDATA()
func (this *QXmlStreamReader) IsCDATA() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader7isCDATAEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:394
// index:0
// Public inline
// bool isComment()
func (this *QXmlStreamReader) IsComment() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader9isCommentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:395
// index:0
// Public inline
// bool isDTD()
func (this *QXmlStreamReader) IsDTD() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader5isDTDEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:396
// index:0
// Public inline
// bool isEntityReference()
func (this *QXmlStreamReader) IsEntityReference() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader17isEntityReferenceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:397
// index:0
// Public inline
// bool isProcessingInstruction()
func (this *QXmlStreamReader) IsProcessingInstruction() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader23isProcessingInstructionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:399
// index:0
// Public
// bool isStandaloneDocument()
func (this *QXmlStreamReader) IsStandaloneDocument() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader20isStandaloneDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:400
// index:0
// Public
// QStringRef documentVersion()
func (this *QXmlStreamReader) DocumentVersion() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader15documentVersionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:401
// index:0
// Public
// QStringRef documentEncoding()
func (this *QXmlStreamReader) DocumentEncoding() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader16documentEncodingEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:403
// index:0
// Public
// qint64 lineNumber()
func (this *QXmlStreamReader) LineNumber() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader10lineNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qxmlstream.h:404
// index:0
// Public
// qint64 columnNumber()
func (this *QXmlStreamReader) ColumnNumber() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12columnNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qxmlstream.h:405
// index:0
// Public
// qint64 characterOffset()
func (this *QXmlStreamReader) CharacterOffset() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader15characterOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qxmlstream.h:407
// index:0
// Public
// QXmlStreamAttributes attributes()
func (this *QXmlStreamReader) Attributes() *QXmlStreamAttributes /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader10attributesEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQXmlStreamAttributesFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:414
// index:0
// Public
// QString readElementText(enum QXmlStreamReader::ReadElementTextBehaviour)
func (this *QXmlStreamReader) ReadElementText(behaviour int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader15readElementTextENS_24ReadElementTextBehaviourE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), behaviour)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:416
// index:0
// Public
// QStringRef name()
func (this *QXmlStreamReader) Name() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:417
// index:0
// Public
// QStringRef namespaceUri()
func (this *QXmlStreamReader) NamespaceUri() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader12namespaceUriEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:418
// index:0
// Public
// QStringRef qualifiedName()
func (this *QXmlStreamReader) QualifiedName() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader13qualifiedNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:419
// index:0
// Public
// QStringRef prefix()
func (this *QXmlStreamReader) Prefix() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader6prefixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:421
// index:0
// Public
// QStringRef processingInstructionTarget()
func (this *QXmlStreamReader) ProcessingInstructionTarget() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader27processingInstructionTargetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:422
// index:0
// Public
// QStringRef processingInstructionData()
func (this *QXmlStreamReader) ProcessingInstructionData() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader25processingInstructionDataEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:424
// index:0
// Public
// QStringRef text()
func (this *QXmlStreamReader) Text() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:427
// index:0
// Public
// void addExtraNamespaceDeclaration(const class QXmlStreamNamespaceDeclaration &)
func (this *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaraction *QXmlStreamNamespaceDeclaration) {
	var convArg0 = extraNamespaceDeclaraction.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:431
// index:0
// Public
// QStringRef dtdName()
func (this *QXmlStreamReader) DtdName() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader7dtdNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:432
// index:0
// Public
// QStringRef dtdPublicId()
func (this *QXmlStreamReader) DtdPublicId() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11dtdPublicIdEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:433
// index:0
// Public
// QStringRef dtdSystemId()
func (this *QXmlStreamReader) DtdSystemId() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11dtdSystemIdEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:443
// index:0
// Public
// void raiseError(const class QString &)
func (this *QXmlStreamReader) RaiseError(message *QString) {
	var convArg0 = message.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader10raiseErrorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:444
// index:0
// Public
// QString errorString()
func (this *QXmlStreamReader) ErrorString() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:445
// index:0
// Public
// QXmlStreamReader::Error error()
func (this *QXmlStreamReader) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:447
// index:0
// Public inline
// bool hasError()
func (this *QXmlStreamReader) HasError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader8hasErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:452
// index:0
// Public
// void setEntityResolver(class QXmlStreamEntityResolver *)
func (this *QXmlStreamReader) SetEntityResolver(resolver *QXmlStreamEntityResolver /*444 QXmlStreamEntityResolver **/) {
	var convArg0 = resolver.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:453
// index:0
// Public
// QXmlStreamEntityResolver * entityResolver()
func (this *QXmlStreamReader) EntityResolver() *QXmlStreamEntityResolver /*444 QXmlStreamEntityResolver **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QXmlStreamReader14entityResolverEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQXmlStreamEntityResolverFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QXmlStreamReader__TokenType = int

const QXmlStreamReader__NoToken QXmlStreamReader__TokenType = 0
const QXmlStreamReader__Invalid QXmlStreamReader__TokenType = 1
const QXmlStreamReader__StartDocument QXmlStreamReader__TokenType = 2
const QXmlStreamReader__EndDocument QXmlStreamReader__TokenType = 3
const QXmlStreamReader__StartElement QXmlStreamReader__TokenType = 4
const QXmlStreamReader__EndElement QXmlStreamReader__TokenType = 5
const QXmlStreamReader__Characters QXmlStreamReader__TokenType = 6
const QXmlStreamReader__Comment QXmlStreamReader__TokenType = 7
const QXmlStreamReader__DTD QXmlStreamReader__TokenType = 8
const QXmlStreamReader__EntityReference QXmlStreamReader__TokenType = 9
const QXmlStreamReader__ProcessingInstruction QXmlStreamReader__TokenType = 10

type QXmlStreamReader__ReadElementTextBehaviour = int

const QXmlStreamReader__ErrorOnUnexpectedElement QXmlStreamReader__ReadElementTextBehaviour = 0
const QXmlStreamReader__IncludeChildElements QXmlStreamReader__ReadElementTextBehaviour = 1
const QXmlStreamReader__SkipChildElements QXmlStreamReader__ReadElementTextBehaviour = 2

type QXmlStreamReader__Error = int

const QXmlStreamReader__NoError QXmlStreamReader__Error = 0
const QXmlStreamReader__UnexpectedElementError QXmlStreamReader__Error = 1
const QXmlStreamReader__CustomError QXmlStreamReader__Error = 2
const QXmlStreamReader__NotWellFormedError QXmlStreamReader__Error = 3
const QXmlStreamReader__PrematureEndOfDocumentError QXmlStreamReader__Error = 4

//  body block end
