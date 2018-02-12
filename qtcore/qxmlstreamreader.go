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
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamReaderFromPointer(cthis unsafe.Pointer) *QXmlStreamReader {
	return &QXmlStreamReader{&qtrt.CObject{cthis}}
}
func (*QXmlStreamReader) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamReader {
	return NewQXmlStreamReaderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:360
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamReader()
func NewQXmlStreamReader() *QXmlStreamReader {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:361
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamReader(QIODevice *)
func NewQXmlStreamReader_1(device *QIODevice /*777 QIODevice **/) *QXmlStreamReader {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:362
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamReader(const QByteArray &)
func NewQXmlStreamReader_2(data *QByteArray) *QXmlStreamReader {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:363
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamReader(const QString &)
func NewQXmlStreamReader_3(data string) *QXmlStreamReader {
	var tmpArg0 = NewQString_5(data)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:364
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamReader(const char *)
func NewQXmlStreamReader_4(data string) *QXmlStreamReader {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReaderC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamReader)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:365
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamReader()
func DeleteQXmlStreamReader(this *QXmlStreamReader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:367
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QXmlStreamReader) SetDevice(device *QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:368
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QXmlStreamReader) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qxmlstream.h:369
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addData(const QByteArray &)
func (this *QXmlStreamReader) AddData(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:370
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addData(const QString &)
func (this *QXmlStreamReader) AddData_1(data string) {
	var tmpArg0 = NewQString_5(data)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:371
// index:2
// Public Visibility=Default Availability=Available
// [-2] void addData(const char *)
func (this *QXmlStreamReader) AddData_2(data string) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader7addDataEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:372
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QXmlStreamReader) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:375
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QXmlStreamReader) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:376
// index:0
// Public Visibility=Default Availability=Available
// [4] QXmlStreamReader::TokenType readNext()
func (this *QXmlStreamReader) ReadNext() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader8readNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:378
// index:0
// Public Visibility=Default Availability=Available
// [1] bool readNextStartElement()
func (this *QXmlStreamReader) ReadNextStartElement() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader20readNextStartElementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void skipCurrentElement()
func (this *QXmlStreamReader) SkipCurrentElement() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader18skipCurrentElementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:381
// index:0
// Public Visibility=Default Availability=Available
// [4] QXmlStreamReader::TokenType tokenType()
func (this *QXmlStreamReader) TokenType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader9tokenTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:382
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tokenString()
func (this *QXmlStreamReader) TokenString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader11tokenStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qxmlstream.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNamespaceProcessing(_Bool)
func (this *QXmlStreamReader) SetNamespaceProcessing(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader22setNamespaceProcessingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:385
// index:0
// Public Visibility=Default Availability=Available
// [1] bool namespaceProcessing()
func (this *QXmlStreamReader) NamespaceProcessing() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader19namespaceProcessingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:387
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isStartDocument()
func (this *QXmlStreamReader) IsStartDocument() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader15isStartDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:388
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEndDocument()
func (this *QXmlStreamReader) IsEndDocument() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader13isEndDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:389
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isStartElement()
func (this *QXmlStreamReader) IsStartElement() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader14isStartElementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:390
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEndElement()
func (this *QXmlStreamReader) IsEndElement() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isEndElementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:391
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCharacters()
func (this *QXmlStreamReader) IsCharacters() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isCharactersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:392
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWhitespace()
func (this *QXmlStreamReader) IsWhitespace() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader12isWhitespaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:393
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCDATA()
func (this *QXmlStreamReader) IsCDATA() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader7isCDATAEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:394
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isComment()
func (this *QXmlStreamReader) IsComment() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader9isCommentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:395
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDTD()
func (this *QXmlStreamReader) IsDTD() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader5isDTDEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:396
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEntityReference()
func (this *QXmlStreamReader) IsEntityReference() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader17isEntityReferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:397
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isProcessingInstruction()
func (this *QXmlStreamReader) IsProcessingInstruction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader23isProcessingInstructionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:399
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStandaloneDocument()
func (this *QXmlStreamReader) IsStandaloneDocument() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader20isStandaloneDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:400
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef documentVersion()
func (this *QXmlStreamReader) DocumentVersion() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader15documentVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:401
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef documentEncoding()
func (this *QXmlStreamReader) DocumentEncoding() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader16documentEncodingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:403
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 lineNumber()
func (this *QXmlStreamReader) LineNumber() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader10lineNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qxmlstream.h:404
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 columnNumber()
func (this *QXmlStreamReader) ColumnNumber() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader12columnNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qxmlstream.h:405
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 characterOffset()
func (this *QXmlStreamReader) CharacterOffset() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader15characterOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qxmlstream.h:407
// index:0
// Public Visibility=Default Availability=Available
// [8] QXmlStreamAttributes attributes()
func (this *QXmlStreamReader) Attributes() *QXmlStreamAttributes /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader10attributesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamAttributesFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamAttributes)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:414
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readElementText(enum QXmlStreamReader::ReadElementTextBehaviour)
func (this *QXmlStreamReader) ReadElementText(behaviour int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader15readElementTextENS_24ReadElementTextBehaviourE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behaviour)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qxmlstream.h:416
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef name()
func (this *QXmlStreamReader) Name() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:417
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef namespaceUri()
func (this *QXmlStreamReader) NamespaceUri() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader12namespaceUriEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:418
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef qualifiedName()
func (this *QXmlStreamReader) QualifiedName() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader13qualifiedNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:419
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef prefix()
func (this *QXmlStreamReader) Prefix() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader6prefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:421
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef processingInstructionTarget()
func (this *QXmlStreamReader) ProcessingInstructionTarget() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader27processingInstructionTargetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:422
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef processingInstructionData()
func (this *QXmlStreamReader) ProcessingInstructionData() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader25processingInstructionDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:424
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef text()
func (this *QXmlStreamReader) Text() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:427
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addExtraNamespaceDeclaration(const QXmlStreamNamespaceDeclaration &)
func (this *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaraction *QXmlStreamNamespaceDeclaration) {
	var convArg0 = extraNamespaceDeclaraction.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader28addExtraNamespaceDeclarationERK30QXmlStreamNamespaceDeclaration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:431
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef dtdName()
func (this *QXmlStreamReader) DtdName() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader7dtdNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:432
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef dtdPublicId()
func (this *QXmlStreamReader) DtdPublicId() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader11dtdPublicIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:433
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef dtdSystemId()
func (this *QXmlStreamReader) DtdSystemId() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader11dtdSystemIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:443
// index:0
// Public Visibility=Default Availability=Available
// [-2] void raiseError(const QString &)
func (this *QXmlStreamReader) RaiseError(message string) {
	var tmpArg0 = NewQString_5(message)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader10raiseErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:444
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QXmlStreamReader) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qxmlstream.h:445
// index:0
// Public Visibility=Default Availability=Available
// [4] QXmlStreamReader::Error error()
func (this *QXmlStreamReader) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:447
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasError()
func (this *QXmlStreamReader) HasError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader8hasErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:452
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEntityResolver(QXmlStreamEntityResolver *)
func (this *QXmlStreamReader) SetEntityResolver(resolver *QXmlStreamEntityResolver /*777 QXmlStreamEntityResolver **/) {
	var convArg0 = resolver.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QXmlStreamReader17setEntityResolverEP24QXmlStreamEntityResolver", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:453
// index:0
// Public Visibility=Default Availability=Available
// [8] QXmlStreamEntityResolver * entityResolver()
func (this *QXmlStreamReader) EntityResolver() *QXmlStreamEntityResolver /*777 QXmlStreamEntityResolver **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QXmlStreamReader14entityResolverEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQXmlStreamEntityResolverFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
