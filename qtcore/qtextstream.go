package qtcore

// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QTextStream struct {
	*qtrt.CObject
}
type QTextStream_ITF interface {
	QTextStream_PTR() *QTextStream
}

func (ptr *QTextStream) QTextStream_PTR() *QTextStream { return ptr }

func (this *QTextStream) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextStream) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextStreamFromPointer(cthis unsafe.Pointer) *QTextStream {
	return &QTextStream{&qtrt.CObject{cthis}}
}
func (*QTextStream) NewFromPointer(cthis unsafe.Pointer) *QTextStream {
	return NewQTextStreamFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextstream.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextStream()
func NewQTextStream() *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QIODevice *)
func NewQTextStream_1(device QIODevice_ITF /*777 QIODevice **/) *QTextStream {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QString *, QIODevice::OpenMode)
func NewQTextStream_2(string string, openMode int) *QTextStream {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:97
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QByteArray *, QIODevice::OpenMode)
func NewQTextStream_3(array QByteArray_ITF /*777 QByteArray **/, openMode int) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:98
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(const QByteArray &, QIODevice::OpenMode)
func NewQTextStream_4(array QByteArray_ITF, openMode int) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2ERK10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextStream()
func DeleteQTextStream(this *QTextStream) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextstream.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)
func (this *QTextStream) SetCodec(codec QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8setCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const char *)
func (this *QTextStream) SetCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8setCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec()
func (this *QTextStream) Codec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextstream.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDetectUnicode(_Bool)
func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream20setAutoDetectUnicodeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDetectUnicode()
func (this *QTextStream) AutoDetectUnicode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream17autoDetectUnicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGenerateByteOrderMark(_Bool)
func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream24setGenerateByteOrderMarkEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), generate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool generateByteOrderMark()
func (this *QTextStream) GenerateByteOrderMark() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream21generateByteOrderMarkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)
func (this *QTextStream) SetLocale(locale QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QTextStream) Locale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QTextStream) SetDevice(device QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QTextStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextstream.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setString(QString *, QIODevice::OpenMode)
func (this *QTextStream) SetString(string string, openMode int) {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStringEP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString * string()
func (this *QTextStream) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::Status status()
func (this *QTextStream) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(enum QTextStream::Status)
func (this *QTextStream) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStatusENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetStatus()
func (this *QTextStream) ResetStatus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream11resetStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QTextStream) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QTextStream) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()
func (this *QTextStream) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool seek(qint64)
func (this *QTextStream) Seek(pos int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream4seekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 pos()
func (this *QTextStream) Pos() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qtextstream.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void skipWhiteSpace()
func (this *QTextStream) SkipWhiteSpace() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14skipWhiteSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLine(qint64)
func (this *QTextStream) ReadLine(maxlen int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8readLineEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool readLineInto(QString *, qint64)
func (this *QTextStream) ReadLineInto(line string, maxlen int64) bool {
	var tmpArg0 = NewQString_5(line)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream12readLineIntoEP7QStringx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readAll()
func (this *QTextStream) ReadAll() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream7readAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QString read(qint64)
func (this *QTextStream) Read(maxlen int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream4readEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldAlignment(enum QTextStream::FieldAlignment)
func (this *QTextStream) SetFieldAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream17setFieldAlignmentENS_14FieldAlignmentE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::FieldAlignment fieldAlignment()
func (this *QTextStream) FieldAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream14fieldAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPadChar(QChar)
func (this *QTextStream) SetPadChar(ch QChar_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream10setPadCharE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:141
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar padChar()
func (this *QTextStream) PadChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream7padCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldWidth(int)
func (this *QTextStream) SetFieldWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream13setFieldWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] int fieldWidth()
func (this *QTextStream) FieldWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream10fieldWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumberFlags(QTextStream::NumberFlags)
func (this *QTextStream) SetNumberFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14setNumberFlagsE6QFlagsINS_10NumberFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::NumberFlags numberFlags()
func (this *QTextStream) NumberFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream11numberFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntegerBase(int)
func (this *QTextStream) SetIntegerBase(base int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14setIntegerBaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), base)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] int integerBase()
func (this *QTextStream) IntegerBase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream11integerBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealNumberNotation(enum QTextStream::RealNumberNotation)
func (this *QTextStream) SetRealNumberNotation(notation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream21setRealNumberNotationENS_18RealNumberNotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), notation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::RealNumberNotation realNumberNotation()
func (this *QTextStream) RealNumberNotation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream18realNumberNotationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealNumberPrecision(int)
func (this *QTextStream) SetRealNumberPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream22setRealNumberPrecisionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int realNumberPrecision()
func (this *QTextStream) RealNumberPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream19realNumberPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:158
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(QChar &)
func (this *QTextStream) Operator_right_shift(ch QChar_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsER5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:159
// index:1
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(char &)
func (this *QTextStream) Operator_right_shift_1(ch byte) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ch)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:160
// index:2
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(short &)
func (this *QTextStream) Operator_right_shift_2(i int16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:161
// index:3
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(unsigned short &)
func (this *QTextStream) Operator_right_shift_3(i uint16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:162
// index:4
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(int &)
func (this *QTextStream) Operator_right_shift_4(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:163
// index:5
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(unsigned int &)
func (this *QTextStream) Operator_right_shift_5(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:164
// index:6
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(long &)
func (this *QTextStream) Operator_right_shift_6(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:165
// index:7
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(unsigned long &)
func (this *QTextStream) Operator_right_shift_7(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:166
// index:8
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(qlonglong &)
func (this *QTextStream) Operator_right_shift_8(i int64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:167
// index:9
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(qulonglong &)
func (this *QTextStream) Operator_right_shift_9(i uint64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:168
// index:10
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(float &)
func (this *QTextStream) Operator_right_shift_10(f float32) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:169
// index:11
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(double &)
func (this *QTextStream) Operator_right_shift_11(f float64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:170
// index:12
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(QString &)
func (this *QTextStream) Operator_right_shift_12(s string) *QTextStream {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:171
// index:13
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(QByteArray &)
func (this *QTextStream) Operator_right_shift_13(array QByteArray_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsER10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:172
// index:14
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(char *)
func (this *QTextStream) Operator_right_shift_14(c string) *QTextStream {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsEPc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(QChar)
func (this *QTextStream) Operator_left_shift(ch QChar_ITF /*123*/) *QTextStream {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:175
// index:1
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(char)
func (this *QTextStream) Operator_left_shift_1(ch byte) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ch)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:176
// index:2
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(short)
func (this *QTextStream) Operator_left_shift_2(i int16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:177
// index:3
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(unsigned short)
func (this *QTextStream) Operator_left_shift_3(i uint16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:178
// index:4
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(int)
func (this *QTextStream) Operator_left_shift_4(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:179
// index:5
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(unsigned int)
func (this *QTextStream) Operator_left_shift_5(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:180
// index:6
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(long)
func (this *QTextStream) Operator_left_shift_6(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:181
// index:7
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(unsigned long)
func (this *QTextStream) Operator_left_shift_7(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:182
// index:8
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(qlonglong)
func (this *QTextStream) Operator_left_shift_8(i int64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:183
// index:9
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(qulonglong)
func (this *QTextStream) Operator_left_shift_9(i uint64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:184
// index:10
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(float)
func (this *QTextStream) Operator_left_shift_10(f float32) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:185
// index:11
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(double)
func (this *QTextStream) Operator_left_shift_11(f float64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:186
// index:12
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const QString &)
func (this *QTextStream) Operator_left_shift_12(s string) *QTextStream {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:187
// index:13
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(QLatin1String)
func (this *QTextStream) Operator_left_shift_13(s QLatin1String_ITF /*123*/) *QTextStream {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:188
// index:14
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const QStringRef &)
func (this *QTextStream) Operator_left_shift_14(s QStringRef_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsERK10QStringRef", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:189
// index:15
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const QByteArray &)
func (this *QTextStream) Operator_left_shift_15(array QByteArray_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:190
// index:16
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const char *)
func (this *QTextStream) Operator_left_shift_16(c string) *QTextStream {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:191
// index:17
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const void *)
func (this *QTextStream) Operator_left_shift_17(ptr unsafe.Pointer /*666*/) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ptr)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

type QTextStream__RealNumberNotation = int

const QTextStream__SmartNotation QTextStream__RealNumberNotation = 0
const QTextStream__FixedNotation QTextStream__RealNumberNotation = 1
const QTextStream__ScientificNotation QTextStream__RealNumberNotation = 2

type QTextStream__FieldAlignment = int

const QTextStream__AlignLeft QTextStream__FieldAlignment = 0
const QTextStream__AlignRight QTextStream__FieldAlignment = 1
const QTextStream__AlignCenter QTextStream__FieldAlignment = 2
const QTextStream__AlignAccountingStyle QTextStream__FieldAlignment = 3

type QTextStream__Status = int

const QTextStream__Ok QTextStream__Status = 0
const QTextStream__ReadPastEnd QTextStream__Status = 1
const QTextStream__ReadCorruptData QTextStream__Status = 2
const QTextStream__WriteFailed QTextStream__Status = 3

type QTextStream__NumberFlag = int

const QTextStream__ShowBase QTextStream__NumberFlag = 1
const QTextStream__ForcePoint QTextStream__NumberFlag = 2
const QTextStream__ForceSign QTextStream__NumberFlag = 4
const QTextStream__UppercaseBase QTextStream__NumberFlag = 8
const QTextStream__UppercaseDigits QTextStream__NumberFlag = 16

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
}

//  keep block end
