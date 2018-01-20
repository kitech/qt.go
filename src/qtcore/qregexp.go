//  header block begin
// /usr/include/qt/QtCore/qregexp.h
// #include <qregexp.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 86
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
type QRegExp struct {
	*qtrt.CObject
}

func (this *QRegExp) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qregexp.h:70
// index:0
// void QRegExp()
func NewQRegExp() *QRegExp {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExpC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(cthis)
	return gothis
}
func NewQRegExpFromPointer(cthis unsafe.Pointer) *QRegExp {
	return &QRegExp{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregexp.h:71
// index:1
// void QRegExp(const class QString &, Qt::CaseSensitivity, enum QRegExp::PatternSyntax)
func NewQRegExp_1(pattern unsafe.Pointer, cs int, syntax int) *QRegExp {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExpC2ERK7QStringN2Qt15CaseSensitivityENS_13PatternSyntaxE", ffiqt.FFI_TYPE_VOID, cthis, pattern, &cs, &syntax)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:74
// index:0
// void ~QRegExp()
func DeleteQRegExp(*QRegExp) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExpD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:79
// index:0
// inline
// void swap(class QRegExp &)
func (this *QRegExp) Swap(other unsafe.Pointer) {
	// 0: (, other QRegExp &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:84
// index:0
// bool isEmpty()
func (this *QRegExp) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:85
// index:0
// bool isValid()
func (this *QRegExp) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:86
// index:0
// QString pattern()
func (this *QRegExp) Pattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7patternEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:87
// index:0
// void setPattern(const class QString &)
func (this *QRegExp) SetPattern(pattern unsafe.Pointer) {
	// 0: (, pattern const QString &), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp10setPatternERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:88
// index:0
// Qt::CaseSensitivity caseSensitivity()
func (this *QRegExp) CaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp15caseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:89
// index:0
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QRegExp) SetCaseSensitivity(cs int) {
	// 0: (, cs Qt::CaseSensitivity), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:90
// index:0
// QRegExp::PatternSyntax patternSyntax()
func (this *QRegExp) PatternSyntax() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp13patternSyntaxEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:91
// index:0
// void setPatternSyntax(enum QRegExp::PatternSyntax)
func (this *QRegExp) SetPatternSyntax(syntax int) {
	// 0: (, syntax QRegExp::PatternSyntax), (&syntax)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp16setPatternSyntaxENS_13PatternSyntaxE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &syntax)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:93
// index:0
// bool isMinimal()
func (this *QRegExp) IsMinimal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp9isMinimalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:94
// index:0
// void setMinimal(_Bool)
func (this *QRegExp) SetMinimal(minimal bool) {
	// 0: (, minimal bool), (&minimal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp10setMinimalEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &minimal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:96
// index:0
// bool exactMatch(const class QString &)
func (this *QRegExp) ExactMatch(str unsafe.Pointer) {
	// 0: (, str const QString &), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp10exactMatchERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:98
// index:0
// int indexIn(const class QString &, int, enum QRegExp::CaretMode)
func (this *QRegExp) IndexIn(str unsafe.Pointer, offset int, caretMode int) {
	// 0: (, str const QString &, offset int, caretMode QRegExp::CaretMode), (str, &offset, &caretMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7indexInERK7QStringiNS_9CaretModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str, &offset, &caretMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:99
// index:0
// int lastIndexIn(const class QString &, int, enum QRegExp::CaretMode)
func (this *QRegExp) LastIndexIn(str unsafe.Pointer, offset int, caretMode int) {
	// 0: (, str const QString &, offset int, caretMode QRegExp::CaretMode), (str, &offset, &caretMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp11lastIndexInERK7QStringiNS_9CaretModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str, &offset, &caretMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:100
// index:0
// int matchedLength()
func (this *QRegExp) MatchedLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp13matchedLengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:102
// index:0
// int captureCount()
func (this *QRegExp) CaptureCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp12captureCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:103
// index:0
// QStringList capturedTexts()
func (this *QRegExp) CapturedTexts() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp13capturedTextsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:104
// index:1
// QStringList capturedTexts()
func (this *QRegExp) CapturedTexts_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp13capturedTextsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:105
// index:0
// QString cap(int)
func (this *QRegExp) Cap(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp3capEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:106
// index:1
// QString cap(int)
func (this *QRegExp) Cap_1(nth int) {
	// 1: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp3capEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:107
// index:0
// int pos(int)
func (this *QRegExp) Pos(nth int) {
	// 0: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp3posEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:108
// index:1
// int pos(int)
func (this *QRegExp) Pos_1(nth int) {
	// 1: (, nth int), (&nth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp3posEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:109
// index:0
// QString errorString()
func (this *QRegExp) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:110
// index:1
// QString errorString()
func (this *QRegExp) ErrorString_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:113
// index:0
// static
// QString escape(const class QString &)
func (this *QRegExp) Escape(str unsafe.Pointer) {
	// 0: (str const QString &), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp6escapeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRegExp_Escape(str unsafe.Pointer) {
	// 0: (str const QString &), (str)
	var nilthis *QRegExp
	nilthis.Escape(str)
}

//  body block end
