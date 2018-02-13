package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QRegularExpression struct {
	*qtrt.CObject
}
type QRegularExpression_ITF interface {
	QRegularExpression_PTR() *QRegularExpression
}

func (ptr *QRegularExpression) QRegularExpression_PTR() *QRegularExpression { return ptr }

func (this *QRegularExpression) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegularExpression) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegularExpressionFromPointer(cthis unsafe.Pointer) *QRegularExpression {
	return &QRegularExpression{&qtrt.CObject{cthis}}
}
func (*QRegularExpression) NewFromPointer(cthis unsafe.Pointer) *QRegularExpression {
	return NewQRegularExpressionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::PatternOptions patternOptions()
func (this *QRegularExpression) PatternOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression14patternOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPatternOptions(QRegularExpression::PatternOptions)
func (this *QRegularExpression) SetPatternOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression17setPatternOptionsE6QFlagsINS_13PatternOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpression()
func NewQRegularExpression() *QRegularExpression {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpression)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpression(const QString &, QRegularExpression::PatternOptions)
func NewQRegularExpression_1(pattern string, options int) *QRegularExpression {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionC2ERK7QString6QFlagsINS_13PatternOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, options)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpression)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegularExpression()
func DeleteQRegularExpression(this *QRegularExpression) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpression &)
func (this *QRegularExpression) Swap(other *QRegularExpression) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pattern()
func (this *QRegularExpression) Pattern() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression7patternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QString &)
func (this *QRegularExpression) SetPattern(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression10setPatternERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRegularExpression) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int patternErrorOffset()
func (this *QRegularExpression) PatternErrorOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression18patternErrorOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QRegularExpression) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int captureCount()
func (this *QRegularExpression) CaptureCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression12captureCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList namedCaptureGroups()
func (this *QRegularExpression) NamedCaptureGroups() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression18namedCaptureGroupsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QString &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) Match(subject string, offset int, matchType int, matchOptions int) *QRegularExpressionMatch /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:126
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QStringRef &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) Match_1(subjectRef *QStringRef, offset int, matchType int, matchOptions int) *QRegularExpressionMatch /*123*/ {
	var convArg0 = subjectRef.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QString &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) GlobalMatch(subject string, offset int, matchType int, matchOptions int) *QRegularExpressionMatchIterator /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QStringRef &, int, enum QRegularExpression::MatchType, QRegularExpression::MatchOptions)
func (this *QRegularExpression) GlobalMatch_1(subjectRef *QStringRef, offset int, matchType int, matchOptions int) *QRegularExpressionMatchIterator /*123*/ {
	var convArg0 = subjectRef.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void optimize()
func (this *QRegularExpression) Optimize() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression8optimizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:143
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString escape(const QString &)
func (this *QRegularExpression) Escape(str string) string {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression6escapeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QRegularExpression_Escape(str string) string {
	var nilthis *QRegularExpression
	rv := nilthis.Escape(str)
	return rv
}

type QRegularExpression__PatternOption = int

const QRegularExpression__NoPatternOption QRegularExpression__PatternOption = 0
const QRegularExpression__CaseInsensitiveOption QRegularExpression__PatternOption = 1
const QRegularExpression__DotMatchesEverythingOption QRegularExpression__PatternOption = 2
const QRegularExpression__MultilineOption QRegularExpression__PatternOption = 4
const QRegularExpression__ExtendedPatternSyntaxOption QRegularExpression__PatternOption = 8
const QRegularExpression__InvertedGreedinessOption QRegularExpression__PatternOption = 16
const QRegularExpression__DontCaptureOption QRegularExpression__PatternOption = 32
const QRegularExpression__UseUnicodePropertiesOption QRegularExpression__PatternOption = 64
const QRegularExpression__OptimizeOnFirstUsageOption QRegularExpression__PatternOption = 128
const QRegularExpression__DontAutomaticallyOptimizeOption QRegularExpression__PatternOption = 256

type QRegularExpression__MatchType = int

const QRegularExpression__NormalMatch QRegularExpression__MatchType = 0
const QRegularExpression__PartialPreferCompleteMatch QRegularExpression__MatchType = 1
const QRegularExpression__PartialPreferFirstMatch QRegularExpression__MatchType = 2
const QRegularExpression__NoMatch QRegularExpression__MatchType = 3

type QRegularExpression__MatchOption = int

const QRegularExpression__NoMatchOption QRegularExpression__MatchOption = 0
const QRegularExpression__AnchoredMatchOption QRegularExpression__MatchOption = 1
const QRegularExpression__DontCheckSubjectStringMatchOption QRegularExpression__MatchOption = 2

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
