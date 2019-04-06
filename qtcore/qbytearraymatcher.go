package qtcore

// /usr/include/qt/QtCore/qbytearraymatcher.h
// #include <qbytearraymatcher.h>
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

/*

 */
type QByteArrayMatcher struct {
	*qtrt.CObject
}
type QByteArrayMatcher_ITF interface {
	QByteArrayMatcher_PTR() *QByteArrayMatcher
}

func (ptr *QByteArrayMatcher) QByteArrayMatcher_PTR() *QByteArrayMatcher { return ptr }

func (this *QByteArrayMatcher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteArrayMatcher) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQByteArrayMatcherFromPointer(cthis unsafe.Pointer) *QByteArrayMatcher {
	return &QByteArrayMatcher{&qtrt.CObject{cthis}}
}
func (*QByteArrayMatcher) NewFromPointer(cthis unsafe.Pointer) *QByteArrayMatcher {
	return NewQByteArrayMatcherFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QByteArrayMatcher()

/*
Constructs an empty byte array matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QByteArrayMatcher) NewForInherit() *QByteArrayMatcher {
	return NewQByteArrayMatcher()
}
func NewQByteArrayMatcher() *QByteArrayMatcher {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArrayMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:54
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QByteArrayMatcher(const QByteArray &)

/*
Constructs an empty byte array matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QByteArrayMatcher) NewForInherit1(pattern QByteArray_ITF) *QByteArrayMatcher {
	return NewQByteArrayMatcher1(pattern)
}
func NewQByteArrayMatcher1(pattern QByteArray_ITF) *QByteArrayMatcher {
	var convArg0 unsafe.Pointer
	if pattern != nil && pattern.QByteArray_PTR() != nil {
		convArg0 = pattern.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArrayMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:55
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QByteArrayMatcher(const char *, int)

/*
Constructs an empty byte array matcher that won't match anything. Call setPattern() to give it a pattern to match.
*/
func (*QByteArrayMatcher) NewForInherit2(pattern string, length int) *QByteArrayMatcher {
	return NewQByteArrayMatcher2(pattern, length)
}
func NewQByteArrayMatcher2(pattern string, length int) *QByteArrayMatcher {
	var convArg0 = qtrt.CString(pattern)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2EPKci", qtrt.FFI_TYPE_POINTER, convArg0, length)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArrayMatcher)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QByteArrayMatcher()

/*

 */
func DeleteQByteArrayMatcher(this *QByteArrayMatcher) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayMatcherD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1040)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1040] QByteArrayMatcher & operator=(const QByteArrayMatcher &)

/*

 */
func (this *QByteArrayMatcher) Operator_equal(other QByteArrayMatcher_ITF) *QByteArrayMatcher {
	var convArg0 unsafe.Pointer
	if other != nil && other.QByteArrayMatcher_PTR() != nil {
		convArg0 = other.QByteArrayMatcher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayMatcheraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArrayMatcher)
	return rv2
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QByteArray &)

/*
Sets the byte array that this byte array matcher will search for to pattern.

See also pattern() and indexIn().
*/
func (this *QByteArrayMatcher) SetPattern(pattern QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if pattern != nil && pattern.QByteArray_PTR() != nil {
		convArg0 = pattern.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayMatcher10setPatternERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QByteArray &, int) const

/*
Searches the byte array ba, from byte position from (default 0, i.e. from the first byte), for the byte array pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in ba, or -1 if no match was found.
*/
func (this *QByteArrayMatcher) IndexIn(ba QByteArray_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QByteArray &, int) const

/*
Searches the byte array ba, from byte position from (default 0, i.e. from the first byte), for the byte array pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in ba, or -1 if no match was found.
*/
func (this *QByteArrayMatcher) IndexInp(ba QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:64
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexIn(const char *, int, int) const

/*
Searches the byte array ba, from byte position from (default 0, i.e. from the first byte), for the byte array pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in ba, or -1 if no match was found.
*/
func (this *QByteArrayMatcher) IndexIn1(str string, len_ int, from int) int {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInEPKcii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:64
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexIn(const char *, int, int) const

/*
Searches the byte array ba, from byte position from (default 0, i.e. from the first byte), for the byte array pattern() that was set in the constructor or in the most recent call to setPattern(). Returns the position where the pattern() matched in ba, or -1 if no match was found.
*/
func (this *QByteArrayMatcher) IndexIn1p(str string, len_ int) int {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	// arg: 2, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInEPKcii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray pattern() const

/*
Returns the byte array pattern that this byte array matcher will search for.

See also setPattern().
*/
func (this *QByteArrayMatcher) Pattern() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7patternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10313() {
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
