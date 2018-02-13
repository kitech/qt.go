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
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// int indexOfIn(const char *, uint, const char *, int, int)
func (this *QStaticByteArrayMatcherBase) InheritIndexOfIn(f func(needle string, nlen uint, haystack string, hlen int, from int) int) {
	qtrt.SetAllInheritCallback(this, "indexOfIn", f)
}

type QStaticByteArrayMatcherBase struct {
	*qtrt.CObject
}
type QStaticByteArrayMatcherBase_ITF interface {
	QStaticByteArrayMatcherBase_PTR() *QStaticByteArrayMatcherBase
}

func (ptr *QStaticByteArrayMatcherBase) QStaticByteArrayMatcherBase_PTR() *QStaticByteArrayMatcherBase {
	return ptr
}

func (this *QStaticByteArrayMatcherBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStaticByteArrayMatcherBase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStaticByteArrayMatcherBaseFromPointer(cthis unsafe.Pointer) *QStaticByteArrayMatcherBase {
	return &QStaticByteArrayMatcherBase{&qtrt.CObject{cthis}}
}
func (*QStaticByteArrayMatcherBase) NewFromPointer(cthis unsafe.Pointer) *QStaticByteArrayMatcherBase {
	return NewQStaticByteArrayMatcherBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:93
// index:0
// Protected inline Visibility=Default Availability=Available
// [-2] void QStaticByteArrayMatcherBase(const char *, uint)
func NewQStaticByteArrayMatcherBase(pattern string, n uint) *QStaticByteArrayMatcherBase {
	var convArg0 = qtrt.CString(pattern)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN27QStaticByteArrayMatcherBaseC2EPKcj", qtrt.FFI_TYPE_POINTER, convArg0, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStaticByteArrayMatcherBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticByteArrayMatcherBase)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:98
// index:0
// Protected Visibility=Default Availability=Available
// [4] int indexOfIn(const char *, uint, const char *, int, int)
func (this *QStaticByteArrayMatcherBase) IndexOfIn(needle string, nlen uint, haystack string, hlen int, from int) int {
	var convArg0 = qtrt.CString(needle)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(haystack)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, nlen, convArg2, hlen, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQStaticByteArrayMatcherBase(this *QStaticByteArrayMatcherBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QStaticByteArrayMatcherBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
}

//  keep block end
