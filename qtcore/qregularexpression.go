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
// extern C begin: 74
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
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
// [4] QRegularExpression::PatternOptions patternOptions() const

/*
Returns the pattern options for the regular expression.

See also setPatternOptions() and pattern().
*/
func (this *QRegularExpression) PatternOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression14patternOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPatternOptions(QRegularExpression::PatternOptions)

/*
Sets the given options as the pattern options of the regular expression. The pattern string is left unchanged.

See also patternOptions() and setPattern().
*/
func (this *QRegularExpression) SetPatternOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression17setPatternOptionsE6QFlagsINS_13PatternOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpression()

/*
Constructs a QRegularExpression object with an empty pattern and no pattern options.

See also setPattern() and setPatternOptions().
*/
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

/*
Constructs a QRegularExpression object with an empty pattern and no pattern options.

See also setPattern() and setPatternOptions().
*/
func NewQRegularExpression_1(pattern string, options int) *QRegularExpression {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionC2ERK7QString6QFlagsINS_13PatternOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, options)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpression)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpression(const QString &, QRegularExpression::PatternOptions)

/*
Constructs a QRegularExpression object with an empty pattern and no pattern options.

See also setPattern() and setPatternOptions().
*/
func NewQRegularExpression_1_(pattern string) *QRegularExpression {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QRegularExpression::PatternOptions=Typedef, QRegularExpression::PatternOptions=Typedef, QFlags<QRegularExpression::PatternOption>
	options := 0
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

/*

 */
func DeleteQRegularExpression(this *QRegularExpression) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression & operator=(const QRegularExpression &)

/*

 */
func (this *QRegularExpression) Operator_equal(re QRegularExpression_ITF) *QRegularExpression {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QRegularExpression & operator=(QRegularExpression &&)

/*

 */
func (this *QRegularExpression) Operator_equal_1(re unsafe.Pointer /*333*/) *QRegularExpression {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpressionaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), re)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpression &)

/*
Swaps the regular expression other with this regular expression. This operation is very fast and never fails.
*/
func (this *QRegularExpression) Swap(other QRegularExpression_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRegularExpression_PTR() != nil {
		convArg0 = other.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pattern() const

/*
Returns the pattern string of the regular expression.

See also setPattern() and patternOptions().
*/
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

/*
Sets the pattern string of the regular expression to pattern. The pattern options are left unchanged.

See also pattern() and setPatternOptions().
*/
func (this *QRegularExpression) SetPattern(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRegularExpression10setPatternERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the regular expression is a valid regular expression (that is, it contains no syntax errors, etc.), or false otherwise. Use errorString() to obtain a textual description of the error.

See also errorString() and patternErrorOffset().
*/
func (this *QRegularExpression) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int patternErrorOffset() const

/*
Returns the offset, inside the pattern string, at which an error was found when checking the validity of the regular expression. If no error was found, then -1 is returned.

See also pattern(), isValid(), and errorString().
*/
func (this *QRegularExpression) PatternErrorOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression18patternErrorOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a textual description of the error found when checking the validity of the regular expression, or "no error" if no error was found.

See also isValid() and patternErrorOffset().
*/
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
// [4] int captureCount() const

/*
Returns the number of capturing groups inside the pattern string, or -1 if the regular expression is not valid.

Note: The implicit capturing group 0 is not included in the returned number.

See also isValid().
*/
func (this *QRegularExpression) CaptureCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression12captureCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList namedCaptureGroups() const

/*
Returns a list of captureCount() + 1 elements, containing the names of the named capturing groups in the pattern string. The list is sorted such that the element of the list at position i is the name of the i-th capturing group, if it has a name, or an empty string if that capturing group is unnamed.

For instance, given the regular expression


  (?<day>\d\d)-(?<month>\d\d)-(?<year>\d\d\d\d) (\w+) (?<name>\w+)



namedCaptureGroups() will return the following list:


  ("", "day", "month", "year", "", "name")



which corresponds to the fact that the capturing group #0 (corresponding to the whole match) has no name, the capturing group #1 has name "day", the capturing group #2 has name "month", etc.

If the regular expression is not valid, returns an empty list.

This function was introduced in  Qt 5.1.

See also isValid(), QRegularExpressionMatch::captured(), and QString::isEmpty().
*/
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
// [8] QRegularExpressionMatch match(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match(subject string, offset int, matchType int, matchOptions int) *QRegularExpressionMatch /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match__(subject string) *QRegularExpressionMatch /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	offset := int(0)
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match__1(subject string, offset int) *QRegularExpressionMatch /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match__2(subject string, offset int, matchType int) *QRegularExpressionMatch /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:126
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match_1(subjectRef QStringRef_ITF, offset int, matchType int, matchOptions int) *QRegularExpressionMatch /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:126
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match_1_(subjectRef QStringRef_ITF) *QRegularExpressionMatch /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	offset := int(0)
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:126
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match_1_1(subjectRef QStringRef_ITF, offset int) *QRegularExpressionMatch /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:126
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch match(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to match the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatch object contains the results of the match.

See also QRegularExpressionMatch and normal matching.
*/
func (this *QRegularExpression) Match_1_2(subjectRef QStringRef_ITF, offset int, matchType int) *QRegularExpressionMatch /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression5matchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch(subject string, offset int, matchType int, matchOptions int) *QRegularExpressionMatchIterator /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch__(subject string) *QRegularExpressionMatchIterator /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	offset := int(0)
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch__1(subject string, offset int) *QRegularExpressionMatchIterator /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QString &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch__2(subject string, offset int, matchType int) *QRegularExpressionMatchIterator /*123*/ {
	var tmpArg0 = NewQString_5(subject)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK7QStringiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch_1(subjectRef QStringRef_ITF, offset int, matchType int, matchOptions int) *QRegularExpressionMatchIterator /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch_1_(subjectRef QStringRef_ITF) *QRegularExpressionMatchIterator /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	offset := int(0)
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch_1_1(subjectRef QStringRef_ITF, offset int) *QRegularExpressionMatchIterator /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	// arg: 2, QRegularExpression::MatchType=Enum, QRegularExpression::MatchType=Enum,
	matchType := 0
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator globalMatch(const QStringRef &, int, QRegularExpression::MatchType, QRegularExpression::MatchOptions) const

/*
Attempts to perform a global match of the regular expression against the given subject string, starting at the position offset inside the subject, using a match of type matchType and honoring the given matchOptions.

The returned QRegularExpressionMatchIterator is positioned before the first match result (if any).

See also QRegularExpressionMatchIterator and global matching.
*/
func (this *QRegularExpression) GlobalMatch_1_2(subjectRef QStringRef_ITF, offset int, matchType int) *QRegularExpressionMatchIterator /*123*/ {
	var convArg0 unsafe.Pointer
	if subjectRef != nil && subjectRef.QStringRef_PTR() != nil {
		convArg0 = subjectRef.QStringRef_PTR().GetCthis()
	}
	// arg: 3, QRegularExpression::MatchOptions=Typedef, QRegularExpression::MatchOptions=Typedef, QFlags<QRegularExpression::MatchOption>
	matchOptions := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression11globalMatchERK10QStringRefiNS_9MatchTypeE6QFlagsINS_11MatchOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, matchType, matchOptions)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void optimize() const

/*
Forces an immediate optimization of the pattern, including JIT-compiling it (if the JIT compiler is enabled).

Patterns are normally optimized only after a certain number of usages. If you can predict that this QRegularExpression object is going to be used for several matches, it may be convenient to optimize it in advance by calling this function.

This function was introduced in  Qt 5.4.

See also QRegularExpression::OptimizeOnFirstUsageOption.
*/
func (this *QRegularExpression) Optimize() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpression8optimizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:143
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString escape(const QString &)

/*
Escapes all characters of str so that they no longer have any special meaning when used as a regular expression pattern string, and returns the escaped string. For instance:


  QString escaped = QRegularExpression::escape("a(x) = f(x) + g(x)");
  // escaped == "a\\(x\\)\\ \\=\\ f\\(x\\)\\ \\+\\ g\\(x\\)"



This is very convenient in order to build patterns from arbitrary strings:


  QString pattern = "(" + QRegularExpression::escape(name) +
                    "|" + QRegularExpression::escape(nickname) + ")";
  QRegularExpression re(pattern);



Note: This function implements Perl's quotemeta algorithm and escapes with a backslash all characters in str, except for the characters in the [A-Z], [a-z] and [0-9] ranges, as well as the underscore (_) character. The only difference with Perl is that a literal NUL inside str is escaped with the sequence "\\0" (backslash + '0'), instead of "\\\0" (backslash + NUL).
*/
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

// /usr/include/qt/QtCore/qregularexpression.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QRegularExpression &) const

/*

 */
func (this *QRegularExpression) Operator_equal_equal(re QRegularExpression_ITF) bool {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpressioneqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QRegularExpression &) const

/*

 */
func (this *QRegularExpression) Operator_not_equal(re QRegularExpression_ITF) bool {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRegularExpressionneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QRegularExpression__PatternOption = int

//
const QRegularExpression__NoPatternOption QRegularExpression__PatternOption = 0

//
const QRegularExpression__CaseInsensitiveOption QRegularExpression__PatternOption = 1

//
const QRegularExpression__DotMatchesEverythingOption QRegularExpression__PatternOption = 2

//
const QRegularExpression__MultilineOption QRegularExpression__PatternOption = 4

//
const QRegularExpression__ExtendedPatternSyntaxOption QRegularExpression__PatternOption = 8

//
const QRegularExpression__InvertedGreedinessOption QRegularExpression__PatternOption = 16

//
const QRegularExpression__DontCaptureOption QRegularExpression__PatternOption = 32

//
const QRegularExpression__UseUnicodePropertiesOption QRegularExpression__PatternOption = 64

//
const QRegularExpression__OptimizeOnFirstUsageOption QRegularExpression__PatternOption = 128

//
const QRegularExpression__DontAutomaticallyOptimizeOption QRegularExpression__PatternOption = 256

/*
The MatchType enum defines the type of the match that should be attempted against the subject string.


*/
type QRegularExpression__MatchType = int

// A normal match is done.
const QRegularExpression__NormalMatch QRegularExpression__MatchType = 0

// The pattern string is matched partially against the subject string. If a partial match is found, then it is recorded, and other matching alternatives are tried as usual. If a complete match is then found, then it's preferred to the partial match; in this case only the complete match is reported. If instead no complete match is found (but only the partial one), then the partial one is reported.
const QRegularExpression__PartialPreferCompleteMatch QRegularExpression__MatchType = 1

// The pattern string is matched partially against the subject string. If a partial match is found, then matching stops and the partial match is reported. In this case, other matching alternatives (potentially leading to a complete match) are not tried. Moreover, this match type assumes that the subject string only a substring of a larger text, and that (in this text) there are other characters beyond the end of the subject string. This can lead to surprising results; see the discussion in the partial matching section for more details.
const QRegularExpression__PartialPreferFirstMatch QRegularExpression__MatchType = 2

//
const QRegularExpression__NoMatch QRegularExpression__MatchType = 3

/*


 */
type QRegularExpression__MatchOption = int

//
const QRegularExpression__NoMatchOption QRegularExpression__MatchOption = 0

//
const QRegularExpression__AnchoredMatchOption QRegularExpression__MatchOption = 1

//
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
