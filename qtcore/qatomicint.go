package qtcore

// /usr/include/qt/QtCore/qatomic.h
// #include <qatomic.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QAtomicInt struct {
	*qtrt.CObject
}
type QAtomicInt_ITF interface {
	QAtomicInt_PTR() *QAtomicInt
}

func (ptr *QAtomicInt) QAtomicInt_PTR() *QAtomicInt { return ptr }

func (this *QAtomicInt) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAtomicInt) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAtomicIntFromPointer(cthis unsafe.Pointer) *QAtomicInt {
	return &QAtomicInt{&qtrt.CObject{cthis}}
}
func (*QAtomicInt) NewFromPointer(cthis unsafe.Pointer) *QAtomicInt {
	return NewQAtomicIntFromPointer(cthis)
}

// /usr/include/qt/QtCore/qatomic.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAtomicInt(int)

/*

 */
func (*QAtomicInt) NewForInherit(value int) *QAtomicInt {
	return NewQAtomicInt(value)
}
func NewQAtomicInt(value int) *QAtomicInt {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QAtomicIntC2Ei", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAtomicIntFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAtomicInt)
	return gothis
}

// /usr/include/qt/QtCore/qatomic.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAtomicInt(int)

/*

 */
func (*QAtomicInt) NewForInherit__() *QAtomicInt {
	return NewQAtomicInt__()
}
func NewQAtomicInt__() *QAtomicInt {
	// arg: 0, int=Int, =Invalid, , Invalid
	value := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QAtomicIntC2Ei", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAtomicIntFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAtomicInt)
	return gothis
}

func DeleteQAtomicInt(this *QAtomicInt) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QAtomicIntD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
