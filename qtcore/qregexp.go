package qtcore

// /usr/include/qt/QtCore/qregexp.h
// #include <qregexp.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QRegExp struct {
	*qtrt.CObject
}
type QRegExp_ITF interface {
	QRegExp_PTR() *QRegExp
}

func (ptr *QRegExp) QRegExp_PTR() *QRegExp { return ptr }

func (this *QRegExp) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegExp) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegExpFromPointer(cthis unsafe.Pointer) *QRegExp {
	return &QRegExp{&qtrt.CObject{cthis}}
}
func (*QRegExp) NewFromPointer(cthis unsafe.Pointer) *QRegExp {
	return NewQRegExpFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregexp.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegExp()

/*
Constructs an empty regexp.

See also isValid() and errorString().
*/
func (*QRegExp) NewForInherit() *QRegExp {
	return NewQRegExp()
}
func NewQRegExp() *QRegExp {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegExp)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegExp(const QString &, Qt::CaseSensitivity, QRegExp::PatternSyntax)

/*
Constructs an empty regexp.

See also isValid() and errorString().
*/
func (*QRegExp) NewForInherit_1(pattern string, cs int, syntax int) *QRegExp {
	return NewQRegExp_1(pattern, cs, syntax)
}
func NewQRegExp_1(pattern string, cs int, syntax int) *QRegExp {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpC2ERK7QStringN2Qt15CaseSensitivityENS_13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, cs, syntax)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegExp)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegExp(const QString &, Qt::CaseSensitivity, QRegExp::PatternSyntax)

/*
Constructs an empty regexp.

See also isValid() and errorString().
*/
func (*QRegExp) NewForInherit_1_(pattern string) *QRegExp {
	return NewQRegExp_1_(pattern)
}
func NewQRegExp_1_(pattern string) *QRegExp {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	// arg: 2, QRegExp::PatternSyntax=Enum, QRegExp::PatternSyntax=Enum, , Invalid
	syntax := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpC2ERK7QStringN2Qt15CaseSensitivityENS_13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, cs, syntax)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegExp)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegExp(const QString &, Qt::CaseSensitivity, QRegExp::PatternSyntax)

/*
Constructs an empty regexp.

See also isValid() and errorString().
*/
func (*QRegExp) NewForInherit_1_1(pattern string, cs int) *QRegExp {
	return NewQRegExp_1_1(pattern, cs)
}
func NewQRegExp_1_1(pattern string, cs int) *QRegExp {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegExp::PatternSyntax=Enum, QRegExp::PatternSyntax=Enum, , Invalid
	syntax := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpC2ERK7QStringN2Qt15CaseSensitivityENS_13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, cs, syntax)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegExp)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegExp()

/*

 */
func DeleteQRegExp(this *QRegExp) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregexp.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegExp & operator=(const QRegExp &)

/*

 */
func (this *QRegExp) Operator_equal(rx QRegExp_ITF) *QRegExp {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegExp)
	return rv2
}

// /usr/include/qt/QtCore/qregexp.h:77
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QRegExp & operator=(QRegExp &&)

/*

 */
func (this *QRegExp) Operator_equal_1(other unsafe.Pointer /*333*/) *QRegExp {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegExp)
	return rv2
}

// /usr/include/qt/QtCore/qregexp.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegExp &)

/*
Swaps regular expression other with this regular expression. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QRegExp) Swap(other QRegExp_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRegExp_PTR() != nil {
		convArg0 = other.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QRegExp &) const

/*

 */
func (this *QRegExp) Operator_equal_equal(rx QRegExp_ITF) bool {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExpeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QRegExp &) const

/*

 */
func (this *QRegExp) Operator_not_equal(rx QRegExp_ITF) bool {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExpneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the pattern string is empty; otherwise returns false.

If you call exactMatch() with an empty pattern on an empty string it will return true; otherwise it returns false since it operates over the whole string. If you call indexIn() with an empty pattern on any string it will return the start offset (0 by default) because the empty pattern matches the 'emptiness' at the start of the string. In this case the length of the match returned by matchedLength() will be 0.

See QString::isEmpty().
*/
func (this *QRegExp) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the regular expression is valid; otherwise returns false. An invalid regular expression never matches.

The pattern [a-z is an example of an invalid pattern, since it lacks a closing square bracket.

Note that the validity of a regexp may also depend on the setting of the wildcard flag, for example *.html is a valid wildcard regexp but an invalid full regexp.

See also errorString().
*/
func (this *QRegExp) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pattern() const

/*
Returns the pattern string of the regular expression. The pattern has either regular expression syntax or wildcard syntax, depending on patternSyntax().

See also setPattern(), patternSyntax(), and caseSensitivity().
*/
func (this *QRegExp) Pattern() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7patternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QString &)

/*
Sets the pattern string to pattern. The case sensitivity, wildcard, and minimal matching options are not changed.

See also pattern(), setPatternSyntax(), and setCaseSensitivity().
*/
func (this *QRegExp) SetPattern(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp10setPatternERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity() const

/*
Returns Qt::CaseSensitive if the regexp is matched case sensitively; otherwise returns Qt::CaseInsensitive.

See also setCaseSensitivity(), patternSyntax(), pattern(), and isMinimal().
*/
func (this *QRegExp) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregexp.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)

/*
Sets case sensitive matching to cs.

If cs is Qt::CaseSensitive, \.txt$ matches readme.txt but not README.TXT.

See also caseSensitivity(), setPatternSyntax(), setPattern(), and setMinimal().
*/
func (this *QRegExp) SetCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegExp::PatternSyntax patternSyntax() const

/*
Returns the syntax used by the regular expression. The default is QRegExp::RegExp.

See also setPatternSyntax(), pattern(), and caseSensitivity().
*/
func (this *QRegExp) PatternSyntax() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp13patternSyntaxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregexp.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPatternSyntax(QRegExp::PatternSyntax)

/*
Sets the syntax mode for the regular expression. The default is QRegExp::RegExp.

Setting syntax to QRegExp::Wildcard enables simple shell-like QRegExp wildcard matching. For example, r*.txt matches the string readme.txt in wildcard mode, but does not match readme.

Setting syntax to QRegExp::FixedString means that the pattern is interpreted as a plain string. Special characters (e.g., backslash) don't need to be escaped then.

See also patternSyntax(), setPattern(), setCaseSensitivity(), and escape().
*/
func (this *QRegExp) SetPatternSyntax(syntax int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp16setPatternSyntaxENS_13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), syntax)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMinimal() const

/*
Returns true if minimal (non-greedy) matching is enabled; otherwise returns false.

See also caseSensitivity() and setMinimal().
*/
func (this *QRegExp) IsMinimal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp9isMinimalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimal(bool)

/*
Enables or disables minimal matching. If minimal is false, matching is greedy (maximal) which is the default.

For example, suppose we have the input string "We must be <b>bold</b>, very <b>bold</b>!" and the pattern <b>.*</b>. With the default greedy (maximal) matching, the match is "We must be <b>bold</b>, very <b>bold</b>!". But with minimal (non-greedy) matching, the first match is: "We must be <b>bold</b>, very <b>bold</b>!" and the second match is "We must be <b>bold</b>, very <b>bold</b>!". In practice we might use the pattern <b>[^<]*</b> instead, although this will still fail for nested tags.

See also isMinimal() and setCaseSensitivity().
*/
func (this *QRegExp) SetMinimal(minimal bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp10setMinimalEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exactMatch(const QString &) const

/*
Returns true if str is matched exactly by this regular expression; otherwise returns false. You can determine how much of the string was matched by calling matchedLength().

For a given regexp string R, exactMatch("R") is the equivalent of indexIn("^R$") since exactMatch() effectively encloses the regexp in the start of string and end of string anchors, except that it sets matchedLength() differently.

For example, if the regular expression is blue, then exactMatch() returns true only for input blue. For inputs bluebell, blutak and lightblue, exactMatch() returns false and matchedLength() will return 4, 3 and 0 respectively.

Although const, this function sets matchedLength(), capturedTexts(), and pos().

See also indexIn() and lastIndexIn().
*/
func (this *QRegExp) ExactMatch(str string) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp10exactMatchERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QString &, int, QRegExp::CaretMode) const

/*
Attempts to find a match in str from position offset (0 by default). If offset is -1, the search starts at the last character; if -2, at the next to last character; etc.

Returns the position of the first match, or -1 if there was no match.

The caretMode parameter can be used to instruct whether ^ should match at index 0 or at offset.

You might prefer to use QString::indexOf(), QString::contains(), or even QStringList::filter(). To replace matches use QString::replace().

Example:


  QString str = "offsets: 1.23 .50 71.00 6.00";
  QRegExp rx("\\d*\\.\\d+");    // primitive floating point matching
  int count = 0;
  int pos = 0;
  while ((pos = rx.indexIn(str, pos)) != -1) {
      ++count;
      pos += rx.matchedLength();
  }
  // pos will be 9, 14, 18 and finally 24; count will end up as 4



Although const, this function sets matchedLength(), capturedTexts() and pos().

If the QRegExp is a wildcard expression (see setPatternSyntax()) and want to test a string against the whole wildcard expression, use exactMatch() instead of this function.

See also lastIndexIn() and exactMatch().
*/
func (this *QRegExp) IndexIn(str string, offset int, caretMode int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7indexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QString &, int, QRegExp::CaretMode) const

/*
Attempts to find a match in str from position offset (0 by default). If offset is -1, the search starts at the last character; if -2, at the next to last character; etc.

Returns the position of the first match, or -1 if there was no match.

The caretMode parameter can be used to instruct whether ^ should match at index 0 or at offset.

You might prefer to use QString::indexOf(), QString::contains(), or even QStringList::filter(). To replace matches use QString::replace().

Example:


  QString str = "offsets: 1.23 .50 71.00 6.00";
  QRegExp rx("\\d*\\.\\d+");    // primitive floating point matching
  int count = 0;
  int pos = 0;
  while ((pos = rx.indexIn(str, pos)) != -1) {
      ++count;
      pos += rx.matchedLength();
  }
  // pos will be 9, 14, 18 and finally 24; count will end up as 4



Although const, this function sets matchedLength(), capturedTexts() and pos().

If the QRegExp is a wildcard expression (see setPatternSyntax()) and want to test a string against the whole wildcard expression, use exactMatch() instead of this function.

See also lastIndexIn() and exactMatch().
*/
func (this *QRegExp) IndexIn__(str string) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	offset := int(0)
	// arg: 2, QRegExp::CaretMode=Enum, QRegExp::CaretMode=Enum, , Invalid
	caretMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7indexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QString &, int, QRegExp::CaretMode) const

/*
Attempts to find a match in str from position offset (0 by default). If offset is -1, the search starts at the last character; if -2, at the next to last character; etc.

Returns the position of the first match, or -1 if there was no match.

The caretMode parameter can be used to instruct whether ^ should match at index 0 or at offset.

You might prefer to use QString::indexOf(), QString::contains(), or even QStringList::filter(). To replace matches use QString::replace().

Example:


  QString str = "offsets: 1.23 .50 71.00 6.00";
  QRegExp rx("\\d*\\.\\d+");    // primitive floating point matching
  int count = 0;
  int pos = 0;
  while ((pos = rx.indexIn(str, pos)) != -1) {
      ++count;
      pos += rx.matchedLength();
  }
  // pos will be 9, 14, 18 and finally 24; count will end up as 4



Although const, this function sets matchedLength(), capturedTexts() and pos().

If the QRegExp is a wildcard expression (see setPatternSyntax()) and want to test a string against the whole wildcard expression, use exactMatch() instead of this function.

See also lastIndexIn() and exactMatch().
*/
func (this *QRegExp) IndexIn__1(str string, offset int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegExp::CaretMode=Enum, QRegExp::CaretMode=Enum, , Invalid
	caretMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7indexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexIn(const QString &, int, QRegExp::CaretMode) const

/*
Attempts to find a match backwards in str from position offset. If offset is -1 (the default), the search starts at the last character; if -2, at the next to last character; etc.

Returns the position of the first match, or -1 if there was no match.

The caretMode parameter can be used to instruct whether ^ should match at index 0 or at offset.

Although const, this function sets matchedLength(), capturedTexts() and pos().

Warning: Searching backwards is much slower than searching forwards.

See also indexIn() and exactMatch().
*/
func (this *QRegExp) LastIndexIn(str string, offset int, caretMode int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp11lastIndexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexIn(const QString &, int, QRegExp::CaretMode) const

/*
Attempts to find a match backwards in str from position offset. If offset is -1 (the default), the search starts at the last character; if -2, at the next to last character; etc.

Returns the position of the first match, or -1 if there was no match.

The caretMode parameter can be used to instruct whether ^ should match at index 0 or at offset.

Although const, this function sets matchedLength(), capturedTexts() and pos().

Warning: Searching backwards is much slower than searching forwards.

See also indexIn() and exactMatch().
*/
func (this *QRegExp) LastIndexIn__(str string) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	offset := int(-1)
	// arg: 2, QRegExp::CaretMode=Enum, QRegExp::CaretMode=Enum, , Invalid
	caretMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp11lastIndexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexIn(const QString &, int, QRegExp::CaretMode) const

/*
Attempts to find a match backwards in str from position offset. If offset is -1 (the default), the search starts at the last character; if -2, at the next to last character; etc.

Returns the position of the first match, or -1 if there was no match.

The caretMode parameter can be used to instruct whether ^ should match at index 0 or at offset.

Although const, this function sets matchedLength(), capturedTexts() and pos().

Warning: Searching backwards is much slower than searching forwards.

See also indexIn() and exactMatch().
*/
func (this *QRegExp) LastIndexIn__1(str string, offset int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegExp::CaretMode=Enum, QRegExp::CaretMode=Enum, , Invalid
	caretMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp11lastIndexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int matchedLength() const

/*
Returns the length of the last matched string, or -1 if there was no match.

See also exactMatch(), indexIn(), and lastIndexIn().
*/
func (this *QRegExp) MatchedLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp13matchedLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int captureCount() const

/*
Returns the number of captures contained in the regular expression.

This function was introduced in  Qt 4.6.
*/
func (this *QRegExp) CaptureCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp12captureCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList capturedTexts() const

/*
Returns a list of the captured text strings.

The first string in the list is the entire matched string. Each subsequent list element contains a string that matched a (capturing) subexpression of the regexp.

For example:


  QRegExp rx("(\\d+)(\\s*)(cm|inch(es)?)");
  int pos = rx.indexIn("Length: 36 inches");
  QStringList list = rx.capturedTexts();
  // list is now ("36 inches", "36", " ", "inches", "es")



The above example also captures elements that may be present but which we have no interest in. This problem can be solved by using non-capturing parentheses:


  QRegExp rx("(\\d+)(?:\\s*)(cm|inch(?:es)?)");
  int pos = rx.indexIn("Length: 36 inches");
  QStringList list = rx.capturedTexts();
  // list is now ("36 inches", "36", "inches")



Note that if you want to iterate over the list, you should iterate over a copy, e.g.


  QStringList list = rx.capturedTexts();
  QStringList::iterator it = list.begin();
  while (it != list.end()) {
      myProcessing(*it);
      ++it;
  }



Some regexps can match an indeterminate number of times. For example if the input string is "Offsets: 12 14 99 231 7" and the regexp, rx, is (\d+)+, we would hope to get a list of all the numbers matched. However, after calling rx.indexIn(str), capturedTexts() will return the list ("12", "12"), i.e. the entire match was "12" and the first subexpression matched was "12". The correct approach is to use cap() in a loop.

The order of elements in the string list is as follows. The first element is the entire matching string. Each subsequent element corresponds to the next capturing open left parentheses. Thus capturedTexts()[1] is the text of the first capturing parentheses, capturedTexts()[2] is the text of the second and so on (corresponding to $1, $2, etc., in some other regexp languages).

See also cap() and pos().
*/
func (this *QRegExp) CapturedTexts() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp13capturedTextsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qregexp.h:104
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList capturedTexts()

/*
Returns a list of the captured text strings.

The first string in the list is the entire matched string. Each subsequent list element contains a string that matched a (capturing) subexpression of the regexp.

For example:


  QRegExp rx("(\\d+)(\\s*)(cm|inch(es)?)");
  int pos = rx.indexIn("Length: 36 inches");
  QStringList list = rx.capturedTexts();
  // list is now ("36 inches", "36", " ", "inches", "es")



The above example also captures elements that may be present but which we have no interest in. This problem can be solved by using non-capturing parentheses:


  QRegExp rx("(\\d+)(?:\\s*)(cm|inch(?:es)?)");
  int pos = rx.indexIn("Length: 36 inches");
  QStringList list = rx.capturedTexts();
  // list is now ("36 inches", "36", "inches")



Note that if you want to iterate over the list, you should iterate over a copy, e.g.


  QStringList list = rx.capturedTexts();
  QStringList::iterator it = list.begin();
  while (it != list.end()) {
      myProcessing(*it);
      ++it;
  }



Some regexps can match an indeterminate number of times. For example if the input string is "Offsets: 12 14 99 231 7" and the regexp, rx, is (\d+)+, we would hope to get a list of all the numbers matched. However, after calling rx.indexIn(str), capturedTexts() will return the list ("12", "12"), i.e. the entire match was "12" and the first subexpression matched was "12". The correct approach is to use cap() in a loop.

The order of elements in the string list is as follows. The first element is the entire matching string. Each subsequent element corresponds to the next capturing open left parentheses. Thus capturedTexts()[1] is the text of the first capturing parentheses, capturedTexts()[2] is the text of the second and so on (corresponding to $1, $2, etc., in some other regexp languages).

See also cap() and pos().
*/
func (this *QRegExp) CapturedTexts_1() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp13capturedTextsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qregexp.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cap(int) const

/*
Returns the text captured by the nth subexpression. The entire match has index 0 and the parenthesized subexpressions have indexes starting from 1 (excluding non-capturing parentheses).


  QRegExp rxlen("(\\d+)(?:\\s*)(cm|inch)");
  int pos = rxlen.indexIn("Length: 189cm");
  if (pos > -1) {
      QString value = rxlen.cap(1); // "189"
      QString unit = rxlen.cap(2);  // "cm"
      // ...
  }



The order of elements matched by cap() is as follows. The first element, cap(0), is the entire matching string. Each subsequent element corresponds to the next capturing open left parentheses. Thus cap(1) is the text of the first capturing parentheses, cap(2) is the text of the second, and so on.

See also capturedTexts() and pos().
*/
func (this *QRegExp) Cap(nth int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp3capEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cap(int) const

/*
Returns the text captured by the nth subexpression. The entire match has index 0 and the parenthesized subexpressions have indexes starting from 1 (excluding non-capturing parentheses).


  QRegExp rxlen("(\\d+)(?:\\s*)(cm|inch)");
  int pos = rxlen.indexIn("Length: 189cm");
  if (pos > -1) {
      QString value = rxlen.cap(1); // "189"
      QString unit = rxlen.cap(2);  // "cm"
      // ...
  }



The order of elements matched by cap() is as follows. The first element, cap(0), is the entire matching string. Each subsequent element corresponds to the next capturing open left parentheses. Thus cap(1) is the text of the first capturing parentheses, cap(2) is the text of the second, and so on.

See also capturedTexts() and pos().
*/
func (this *QRegExp) Cap__() string {
	// arg: 0, int=Int, =Invalid, , Invalid
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp3capEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:106
// index:1
// Public Visibility=Default Availability=Available
// [8] QString cap(int)

/*
Returns the text captured by the nth subexpression. The entire match has index 0 and the parenthesized subexpressions have indexes starting from 1 (excluding non-capturing parentheses).


  QRegExp rxlen("(\\d+)(?:\\s*)(cm|inch)");
  int pos = rxlen.indexIn("Length: 189cm");
  if (pos > -1) {
      QString value = rxlen.cap(1); // "189"
      QString unit = rxlen.cap(2);  // "cm"
      // ...
  }



The order of elements matched by cap() is as follows. The first element, cap(0), is the entire matching string. Each subsequent element corresponds to the next capturing open left parentheses. Thus cap(1) is the text of the first capturing parentheses, cap(2) is the text of the second, and so on.

See also capturedTexts() and pos().
*/
func (this *QRegExp) Cap_1(nth int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp3capEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:106
// index:1
// Public Visibility=Default Availability=Available
// [8] QString cap(int)

/*
Returns the text captured by the nth subexpression. The entire match has index 0 and the parenthesized subexpressions have indexes starting from 1 (excluding non-capturing parentheses).


  QRegExp rxlen("(\\d+)(?:\\s*)(cm|inch)");
  int pos = rxlen.indexIn("Length: 189cm");
  if (pos > -1) {
      QString value = rxlen.cap(1); // "189"
      QString unit = rxlen.cap(2);  // "cm"
      // ...
  }



The order of elements matched by cap() is as follows. The first element, cap(0), is the entire matching string. Each subsequent element corresponds to the next capturing open left parentheses. Thus cap(1) is the text of the first capturing parentheses, cap(2) is the text of the second, and so on.

See also capturedTexts() and pos().
*/
func (this *QRegExp) Cap_1_() string {
	// arg: 0, int=Int, =Invalid, , Invalid
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp3capEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int pos(int) const

/*
Returns the position of the nth captured text in the searched string. If nth is 0 (the default), pos() returns the position of the whole match.

Example:


  QRegExp rx("/([a-z]+)/([a-z]+)");
  rx.indexIn("Output /dev/null");   // returns 7 (position of /dev/null)
  rx.pos(0);                        // returns 7 (position of /dev/null)
  rx.pos(1);                        // returns 8 (position of dev)
  rx.pos(2);                        // returns 12 (position of null)



For zero-length matches, pos() always returns -1. (For example, if cap(4) would return an empty string, pos(4) returns -1.) This is a feature of the implementation.

See also cap() and capturedTexts().
*/
func (this *QRegExp) Pos(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp3posEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int pos(int) const

/*
Returns the position of the nth captured text in the searched string. If nth is 0 (the default), pos() returns the position of the whole match.

Example:


  QRegExp rx("/([a-z]+)/([a-z]+)");
  rx.indexIn("Output /dev/null");   // returns 7 (position of /dev/null)
  rx.pos(0);                        // returns 7 (position of /dev/null)
  rx.pos(1);                        // returns 8 (position of dev)
  rx.pos(2);                        // returns 12 (position of null)



For zero-length matches, pos() always returns -1. (For example, if cap(4) would return an empty string, pos(4) returns -1.) This is a feature of the implementation.

See also cap() and capturedTexts().
*/
func (this *QRegExp) Pos__() int {
	// arg: 0, int=Int, =Invalid, , Invalid
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp3posEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:108
// index:1
// Public Visibility=Default Availability=Available
// [4] int pos(int)

/*
Returns the position of the nth captured text in the searched string. If nth is 0 (the default), pos() returns the position of the whole match.

Example:


  QRegExp rx("/([a-z]+)/([a-z]+)");
  rx.indexIn("Output /dev/null");   // returns 7 (position of /dev/null)
  rx.pos(0);                        // returns 7 (position of /dev/null)
  rx.pos(1);                        // returns 8 (position of dev)
  rx.pos(2);                        // returns 12 (position of null)



For zero-length matches, pos() always returns -1. (For example, if cap(4) would return an empty string, pos(4) returns -1.) This is a feature of the implementation.

See also cap() and capturedTexts().
*/
func (this *QRegExp) Pos_1(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp3posEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:108
// index:1
// Public Visibility=Default Availability=Available
// [4] int pos(int)

/*
Returns the position of the nth captured text in the searched string. If nth is 0 (the default), pos() returns the position of the whole match.

Example:


  QRegExp rx("/([a-z]+)/([a-z]+)");
  rx.indexIn("Output /dev/null");   // returns 7 (position of /dev/null)
  rx.pos(0);                        // returns 7 (position of /dev/null)
  rx.pos(1);                        // returns 8 (position of dev)
  rx.pos(2);                        // returns 12 (position of null)



For zero-length matches, pos() always returns -1. (For example, if cap(4) would return an empty string, pos(4) returns -1.) This is a feature of the implementation.

See also cap() and capturedTexts().
*/
func (this *QRegExp) Pos_1_() int {
	// arg: 0, int=Int, =Invalid, , Invalid
	nth := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp3posEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a text string that explains why a regexp pattern is invalid the case being; otherwise returns "no error occurred".

See also isValid().
*/
func (this *QRegExp) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:110
// index:1
// Public Visibility=Default Availability=Available
// [8] QString errorString()

/*
Returns a text string that explains why a regexp pattern is invalid the case being; otherwise returns "no error occurred".

See also isValid().
*/
func (this *QRegExp) ErrorString_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString escape(const QString &)

/*
Returns the string str with every regexp special character escaped with a backslash. The special characters are $, (,), *, +, ., ?, [, ,], ^, {, | and }.

Example:


  s1 = QRegExp::escape("bingo");   // s1 == "bingo"
  s2 = QRegExp::escape("f(x)");    // s2 == "f\\(x\\)"



This function is useful to construct regexp patterns dynamically:


  QRegExp rx("(" + QRegExp::escape(name) +
             "|" + QRegExp::escape(alias) + ")");



See also setPatternSyntax().
*/
func (this *QRegExp) Escape(str string) string {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp6escapeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QRegExp_Escape(str string) string {
	var nilthis *QRegExp
	rv := nilthis.Escape(str)
	return rv
}

/*
The syntax used to interpret the meaning of the pattern.



See also setPatternSyntax().

*/
type QRegExp__PatternSyntax = int

// A rich Perl-like pattern matching syntax. This is the default.
const QRegExp__RegExp QRegExp__PatternSyntax = 0

// This provides a simple pattern matching syntax similar to that used by shells (command interpreters) for "file globbing". See QRegExp wildcard matching.
const QRegExp__Wildcard QRegExp__PatternSyntax = 1

// The pattern is a fixed string. This is equivalent to using the RegExp pattern on a string in which all metacharacters are escaped using escape().
const QRegExp__FixedString QRegExp__PatternSyntax = 2

//
const QRegExp__RegExp2 QRegExp__PatternSyntax = 3

// This is similar to Wildcard but with the behavior of a Unix shell. The wildcard characters can be escaped with the character "\".
const QRegExp__WildcardUnix QRegExp__PatternSyntax = 4

//
const QRegExp__W3CXmlSchema11 QRegExp__PatternSyntax = 5

func (this *QRegExp) PatternSyntaxItemName(val int) string {
	switch val {
	case QRegExp__RegExp: // 0
		return "RegExp"
	case QRegExp__Wildcard: // 1
		return "Wildcard"
	case QRegExp__FixedString: // 2
		return "FixedString"
	case QRegExp__RegExp2: // 3
		return "RegExp2"
	case QRegExp__WildcardUnix: // 4
		return "WildcardUnix"
	case QRegExp__W3CXmlSchema11: // 5
		return "W3CXmlSchema11"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QRegExp_PatternSyntaxItemName(val int) string {
	var nilthis *QRegExp
	return nilthis.PatternSyntaxItemName(val)
}

/*
The CaretMode enum defines the different meanings of the caret (^) in a regular expression. The possible values are:


*/
type QRegExp__CaretMode = int

//
const QRegExp__CaretAtZero QRegExp__CaretMode = 0

// The caret corresponds to the start offset of the search.
const QRegExp__CaretAtOffset QRegExp__CaretMode = 1

// The caret never matches.
const QRegExp__CaretWontMatch QRegExp__CaretMode = 2

func (this *QRegExp) CaretModeItemName(val int) string {
	switch val {
	case QRegExp__CaretAtZero: // 0
		return "CaretAtZero"
	case QRegExp__CaretAtOffset: // 1
		return "CaretAtOffset"
	case QRegExp__CaretWontMatch: // 2
		return "CaretWontMatch"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QRegExp_CaretModeItemName(val int) string {
	var nilthis *QRegExp
	return nilthis.CaretModeItemName(val)
}

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
