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
func NewQByteArrayMatcher_1(pattern QByteArray_ITF) *QByteArrayMatcher {
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
func NewQByteArrayMatcher_2(pattern string, length int) *QByteArrayMatcher {
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
// [4] int indexIn(const QByteArray &, int)
func (this *QByteArrayMatcher) IndexIn(ba QByteArray_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:64
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexIn(const char *, int, int)
func (this *QByteArrayMatcher) IndexIn_1(str string, len int, from int) int {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInEPKcii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray pattern()
func (this *QByteArrayMatcher) Pattern() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7patternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
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
