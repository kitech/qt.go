//  header block begin
// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
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
type QRegularExpression struct {
	*qtrt.CObject
}

func (this *QRegularExpression) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qregularexpression.h:81
// index:0
// QRegularExpression::PatternOptions patternOptions()
func (this *QRegularExpression) PatternOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression14patternOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:82
// index:0
// void setPatternOptions(QRegularExpression::PatternOptions)
func (this *QRegularExpression) SetPatternOptions(options int) {
	// 0: (, QFlags<QRegularExpression::PatternOption> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression17setPatternOptionsE6QFlagsINS_13PatternOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:84
// index:0
// void QRegularExpression()
func NewQRegularExpression() *QRegularExpression {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(cthis)
	return gothis
}
func NewQRegularExpressionFromPointer(cthis unsafe.Pointer) *QRegularExpression {
	return &QRegularExpression{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregularexpression.h:85
// index:1
// void QRegularExpression(const class QString &, QRegularExpression::PatternOptions)
func NewQRegularExpression_1(pattern unsafe.Pointer, options int) *QRegularExpression {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionC2ERK7QString6QFlagsINS_13PatternOptionEE", ffiqt.FFI_TYPE_VOID, cthis, pattern, options)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:87
// index:0
// void ~QRegularExpression()
func DeleteQRegularExpression(*QRegularExpression) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:95
// index:0
// inline
// void swap(class QRegularExpression &)
func (this *QRegularExpression) Swap(other unsafe.Pointer) {
	// 0: (, other QRegularExpression &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:97
// index:0
// QString pattern()
func (this *QRegularExpression) Pattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7patternEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:98
// index:0
// void setPattern(const class QString &)
func (this *QRegularExpression) SetPattern(pattern unsafe.Pointer) {
	// 0: (, pattern const QString &), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression10setPatternERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:100
// index:0
// bool isValid()
func (this *QRegularExpression) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:101
// index:0
// int patternErrorOffset()
func (this *QRegularExpression) PatternErrorOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression18patternErrorOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:102
// index:0
// QString errorString()
func (this *QRegularExpression) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:104
// index:0
// int captureCount()
func (this *QRegularExpression) CaptureCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression12captureCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:105
// index:0
// QStringList namedCaptureGroups()
func (this *QRegularExpression) NamedCaptureGroups() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression18namedCaptureGroupsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:121
// index:0
// QRegularExpressionMatch match(const class QString &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) Match(subject unsafe.Pointer, offset int, matchType int, matchOptions int) {
	// 0: (, subject const QString &, offset int, matchType QRegularExpression::MatchType, QFlags<QRegularExpression::MatchOption> matchOptions), (subject, &offset, &matchType, matchOptions)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), subject, &offset, &matchType, matchOptions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:126
// index:1
// QRegularExpressionMatch match(const class QStringRef &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) Match_1(subjectRef unsafe.Pointer, offset int, matchType int, matchOptions int) {
	// 1: (, subjectRef const QStringRef &, offset int, matchType QRegularExpression::MatchType, QFlags<QRegularExpression::MatchOption> matchOptions), (subjectRef, &offset, &matchType, matchOptions)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), subjectRef, &offset, &matchType, matchOptions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:131
// index:0
// QRegularExpressionMatchIterator globalMatch(const class QString &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) GlobalMatch(subject unsafe.Pointer, offset int, matchType int, matchOptions int) {
	// 0: (, subject const QString &, offset int, matchType QRegularExpression::MatchType, QFlags<QRegularExpression::MatchOption> matchOptions), (subject, &offset, &matchType, matchOptions)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), subject, &offset, &matchType, matchOptions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:136
// index:1
// QRegularExpressionMatchIterator globalMatch(const class QStringRef &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) GlobalMatch_1(subjectRef unsafe.Pointer, offset int, matchType int, matchOptions int) {
	// 1: (, subjectRef const QStringRef &, offset int, matchType QRegularExpression::MatchType, QFlags<QRegularExpression::MatchOption> matchOptions), (subjectRef, &offset, &matchType, matchOptions)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), subjectRef, &offset, &matchType, matchOptions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:141
// index:0
// void optimize()
func (this *QRegularExpression) Optimize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression8optimizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:143
// index:0
// static
// QString escape(const class QString &)
func (this *QRegularExpression) Escape(str unsafe.Pointer) {
	// 0: (str const QString &), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression6escapeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRegularExpression_Escape(str unsafe.Pointer) {
	// 0: (str const QString &), (str)
	var nilthis *QRegularExpression
	nilthis.Escape(str)
}

//  body block end
