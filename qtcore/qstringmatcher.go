package qtcore

// /usr/include/qt/QtCore/qstringmatcher.h
// #include <qstringmatcher.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QStringMatcher struct {
	*qtrt.CObject
}
type QStringMatcher_ITF interface {
	QStringMatcher_PTR() *QStringMatcher
}

func (ptr *QStringMatcher) QStringMatcher_PTR() *QStringMatcher { return ptr }

func (this *QStringMatcher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringMatcher) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringMatcherFromPointer(cthis unsafe.Pointer) *QStringMatcher {
	return &QStringMatcher{&qtrt.CObject{cthis}}
}
func (*QStringMatcher) NewFromPointer(cthis unsafe.Pointer) *QStringMatcher {
	return NewQStringMatcherFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringmatcher.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStringMatcher()

/*
Constructs an empty string matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QStringMatcher) NewForInherit() *QStringMatcher {
	return NewQStringMatcher()
}
func NewQStringMatcher() *QStringMatcher {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcherC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:54
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStringMatcher(const QString &, Qt::CaseSensitivity)

/*
Constructs an empty string matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QStringMatcher) NewForInherit1(pattern string, cs int) *QStringMatcher {
	return NewQStringMatcher1(pattern, cs)
}
func NewQStringMatcher1(pattern string, cs int) *QStringMatcher {
	var tmpArg0 = NewQString5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcherC2ERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, cs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:54
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStringMatcher(const QString &, Qt::CaseSensitivity)

/*
Constructs an empty string matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QStringMatcher) NewForInherit1p(pattern string) *QStringMatcher {
	return NewQStringMatcher1p(pattern)
}
func NewQStringMatcher1p(pattern string) *QStringMatcher {
	var tmpArg0 = NewQString5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcherC2ERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, cs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:56
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QStringMatcher(const QChar *, int, Qt::CaseSensitivity)

/*
Constructs an empty string matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QStringMatcher) NewForInherit2(uc QChar_ITF /*777 const QChar **/, len_ int, cs int) *QStringMatcher {
	return NewQStringMatcher2(uc, len_, cs)
}
func NewQStringMatcher2(uc QChar_ITF /*777 const QChar **/, len_ int, cs int) *QStringMatcher {
	var convArg0 unsafe.Pointer
	if uc != nil && uc.QChar_PTR() != nil {
		convArg0 = uc.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcherC2EPK5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, len_, cs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:56
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QStringMatcher(const QChar *, int, Qt::CaseSensitivity)

/*
Constructs an empty string matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QStringMatcher) NewForInherit2p(uc QChar_ITF /*777 const QChar **/, len_ int) *QStringMatcher {
	return NewQStringMatcher2p(uc, len_)
}
func NewQStringMatcher2p(uc QChar_ITF /*777 const QChar **/, len_ int) *QStringMatcher {
	var convArg0 unsafe.Pointer
	if uc != nil && uc.QChar_PTR() != nil {
		convArg0 = uc.QChar_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcherC2EPK5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, len_, cs)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStringMatcher()

/*

 */
func DeleteQStringMatcher(this *QStringMatcher) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcherD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1048)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstringmatcher.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1048] QStringMatcher & operator=(const QStringMatcher &)

/*

 */
func (this *QStringMatcher) Operator_equal(other QStringMatcher_ITF) *QStringMatcher {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStringMatcher_PTR() != nil {
		convArg0 = other.QStringMatcher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcheraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringMatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringMatcher)
	return rv2
}

// /usr/include/qt/QtCore/qstringmatcher.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QString &)

/*
Sets the string that this string matcher will search for to pattern.

See also pattern(), setCaseSensitivity(), and indexIn().
*/
func (this *QStringMatcher) SetPattern(pattern string) {
	var tmpArg0 = NewQString5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcher10setPatternERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)

/*
Sets the case sensitivity setting of this string matcher to cs.

See also caseSensitivity(), setPattern(), and indexIn().
*/
func (this *QStringMatcher) SetCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringMatcher18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QString &, int) const

/*
Searches the string str from character position from (default 0, i.e. from the first character), for the string pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in str, or -1 if no match was found.

See also setPattern() and setCaseSensitivity().
*/
func (this *QStringMatcher) IndexIn(str string, from int) int {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringmatcher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QString &, int) const

/*
Searches the string str from character position from (default 0, i.e. from the first character), for the string pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in str, or -1 if no match was found.

See also setPattern() and setCaseSensitivity().
*/
func (this *QStringMatcher) IndexInp(str string) int {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringmatcher.h:67
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QChar *, int, int) const

/*
Searches the string str from character position from (default 0, i.e. from the first character), for the string pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in str, or -1 if no match was found.

See also setPattern() and setCaseSensitivity().
*/
func (this *QStringMatcher) IndexIn1(str QChar_ITF /*777 const QChar **/, length int, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QChar_PTR() != nil {
		convArg0 = str.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInEPK5QCharii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringmatcher.h:67
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QChar *, int, int) const

/*
Searches the string str from character position from (default 0, i.e. from the first character), for the string pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in str, or -1 if no match was found.

See also setPattern() and setCaseSensitivity().
*/
func (this *QStringMatcher) IndexIn1p(str QChar_ITF /*777 const QChar **/, length int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QChar_PTR() != nil {
		convArg0 = str.QChar_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInEPK5QCharii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringmatcher.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pattern() const

/*
Returns the string pattern that this string matcher will search for.

See also setPattern().
*/
func (this *QStringMatcher) Pattern() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStringMatcher7patternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstringmatcher.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity() const

/*
Returns the case sensitivity setting for this string matcher.

See also setCaseSensitivity().
*/
func (this *QStringMatcher) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStringMatcher15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
