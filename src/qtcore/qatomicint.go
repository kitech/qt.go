//  header block begin
// /usr/include/qt/QtCore/qatomic.h
// #include <qatomic.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QAtomicInt struct {
	*qtrt.CObject
}

func (this *QAtomicInt) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qatomic.h:162
// index:0
// inline
// void QAtomicInt(int)
func NewQAtomicInt(value int) *QAtomicInt {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QAtomicIntC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &value)
	gopp.ErrPrint(err, rv)
	gothis := NewQAtomicIntFromPointer(cthis)
	return gothis
}
func NewQAtomicIntFromPointer(cthis unsafe.Pointer) *QAtomicInt {
	return &QAtomicInt{&qtrt.CObject{cthis}}
}

//  body block end
