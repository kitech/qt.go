//  header block begin
// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QTextStream struct {
	*qtrt.CObject
}

func (this *QTextStream) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qtextstream.h:93
// index:0
// void QTextStream()
func NewQTextStream() *QTextStream {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}
func NewQTextStreamFromPointer(cthis unsafe.Pointer) *QTextStream {
	return &QTextStream{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextstream.h:94
// index:1
// void QTextStream(class QIODevice *)
func NewQTextStream_1(device unsafe.Pointer) *QTextStream {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, device)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:95
// index:2
// void QTextStream(FILE *, class QIODevice::OpenMode)
func NewQTextStream_2(fileHandle unsafe.Pointer, openMode int) *QTextStream {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2EP8_IO_FILE6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, cthis, fileHandle, &openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:96
// index:3
// void QTextStream(class QString *, class QIODevice::OpenMode)
func NewQTextStream_3(string unsafe.Pointer, openMode int) *QTextStream {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2EP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, cthis, string, &openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:97
// index:4
// void QTextStream(class QByteArray *, class QIODevice::OpenMode)
func NewQTextStream_4(array unsafe.Pointer, openMode int) *QTextStream {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, cthis, array, &openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:98
// index:5
// void QTextStream(const class QByteArray &, class QIODevice::OpenMode)
func NewQTextStream_5(array unsafe.Pointer, openMode int) *QTextStream {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2ERK10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, cthis, array, &openMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:99
// index:0
// virtual
// void ~QTextStream()
func DeleteQTextStream(*QTextStream) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:102
// index:0
// void setCodec(class QTextCodec *)
func (this *QTextStream) SetCodec(codec unsafe.Pointer) {
	// 0: (, codec QTextCodec *), (codec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_VOID, this.GetCthis(), codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:103
// index:1
// void setCodec(const char *)
func (this *QTextStream) SetCodec_1(codecName unsafe.Pointer) {
	// 1: (, codecName const char *), (codecName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream8setCodecEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), codecName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:104
// index:0
// QTextCodec * codec()
func (this *QTextStream) Codec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream5codecEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:105
// index:0
// void setAutoDetectUnicode(_Bool)
func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream20setAutoDetectUnicodeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:106
// index:0
// bool autoDetectUnicode()
func (this *QTextStream) AutoDetectUnicode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream17autoDetectUnicodeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:107
// index:0
// void setGenerateByteOrderMark(_Bool)
func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	// 0: (, generate bool), (&generate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream24setGenerateByteOrderMarkEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &generate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:108
// index:0
// bool generateByteOrderMark()
func (this *QTextStream) GenerateByteOrderMark() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream21generateByteOrderMarkEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:111
// index:0
// void setLocale(const class QLocale &)
func (this *QTextStream) SetLocale(locale unsafe.Pointer) {
	// 0: (, locale const QLocale &), (locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setLocaleERK7QLocale", ffiqt.FFI_TYPE_VOID, this.GetCthis(), locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:112
// index:0
// QLocale locale()
func (this *QTextStream) Locale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6localeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:114
// index:0
// void setDevice(class QIODevice *)
func (this *QTextStream) SetDevice(device unsafe.Pointer) {
	// 0: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:115
// index:0
// QIODevice * device()
func (this *QTextStream) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:117
// index:0
// void setString(class QString *, class QIODevice::OpenMode)
func (this *QTextStream) SetString(string unsafe.Pointer, openMode int) {
	// 0: (, string QString *, QFlags<QIODevice::OpenModeFlag> openMode), (string, &openMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setStringEP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), string, &openMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:118
// index:0
// QString * string()
func (this *QTextStream) String() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6stringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:120
// index:0
// QTextStream::Status status()
func (this *QTextStream) Status() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6statusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:121
// index:0
// void setStatus(enum QTextStream::Status)
func (this *QTextStream) SetStatus(status int) {
	// 0: (, status QTextStream::Status), (&status)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setStatusENS_6StatusE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &status)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:122
// index:0
// void resetStatus()
func (this *QTextStream) ResetStatus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream11resetStatusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:124
// index:0
// bool atEnd()
func (this *QTextStream) AtEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream5atEndEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:125
// index:0
// void reset()
func (this *QTextStream) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:126
// index:0
// void flush()
func (this *QTextStream) Flush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream5flushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:127
// index:0
// bool seek(qint64)
func (this *QTextStream) Seek(pos int64) {
	// 0: (, pos qint64), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream4seekEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:128
// index:0
// qint64 pos()
func (this *QTextStream) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:130
// index:0
// void skipWhiteSpace()
func (this *QTextStream) SkipWhiteSpace() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream14skipWhiteSpaceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:132
// index:0
// QString readLine(qint64)
func (this *QTextStream) ReadLine(maxlen int64) {
	// 0: (, maxlen qint64), (&maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream8readLineEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:133
// index:0
// bool readLineInto(class QString *, qint64)
func (this *QTextStream) ReadLineInto(line unsafe.Pointer, maxlen int64) {
	// 0: (, line QString *, maxlen qint64), (line, &maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream12readLineIntoEP7QStringx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), line, &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:134
// index:0
// QString readAll()
func (this *QTextStream) ReadAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream7readAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:135
// index:0
// QString read(qint64)
func (this *QTextStream) Read(maxlen int64) {
	// 0: (, maxlen qint64), (&maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream4readEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:137
// index:0
// void setFieldAlignment(enum QTextStream::FieldAlignment)
func (this *QTextStream) SetFieldAlignment(alignment int) {
	// 0: (, alignment QTextStream::FieldAlignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream17setFieldAlignmentENS_14FieldAlignmentE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:138
// index:0
// QTextStream::FieldAlignment fieldAlignment()
func (this *QTextStream) FieldAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream14fieldAlignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:140
// index:0
// void setPadChar(class QChar)
func (this *QTextStream) SetPadChar(ch unsafe.Pointer) {
	// 0: (, ch QChar), (ch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream10setPadCharE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:141
// index:0
// QChar padChar()
func (this *QTextStream) PadChar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream7padCharEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:143
// index:0
// void setFieldWidth(int)
func (this *QTextStream) SetFieldWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream13setFieldWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:144
// index:0
// int fieldWidth()
func (this *QTextStream) FieldWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream10fieldWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:146
// index:0
// void setNumberFlags(QTextStream::NumberFlags)
func (this *QTextStream) SetNumberFlags(flags int) {
	// 0: (, QFlags<QTextStream::NumberFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream14setNumberFlagsE6QFlagsINS_10NumberFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:147
// index:0
// QTextStream::NumberFlags numberFlags()
func (this *QTextStream) NumberFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream11numberFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:149
// index:0
// void setIntegerBase(int)
func (this *QTextStream) SetIntegerBase(base int) {
	// 0: (, base int), (&base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream14setIntegerBaseEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:150
// index:0
// int integerBase()
func (this *QTextStream) IntegerBase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream11integerBaseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:152
// index:0
// void setRealNumberNotation(enum QTextStream::RealNumberNotation)
func (this *QTextStream) SetRealNumberNotation(notation int) {
	// 0: (, notation QTextStream::RealNumberNotation), (&notation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream21setRealNumberNotationENS_18RealNumberNotationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &notation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:153
// index:0
// QTextStream::RealNumberNotation realNumberNotation()
func (this *QTextStream) RealNumberNotation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream18realNumberNotationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:155
// index:0
// void setRealNumberPrecision(int)
func (this *QTextStream) SetRealNumberPrecision(precision int) {
	// 0: (, precision int), (&precision)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream22setRealNumberPrecisionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:156
// index:0
// int realNumberPrecision()
func (this *QTextStream) RealNumberPrecision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream19realNumberPrecisionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
