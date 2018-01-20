//  header block begin
// /usr/include/qt/QtCore/qsemaphore.h
// #include <qsemaphore.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QSemaphore struct {
	*qtrt.CObject
}

func (this *QSemaphore) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSemaphoreFromPointer(cthis unsafe.Pointer) *QSemaphore {
	return &QSemaphore{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsemaphore.h:55
// index:0
// Public
// void QSemaphore(int)
func NewQSemaphore(n int) *QSemaphore {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSemaphoreC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &n)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:56
// index:0
// Public
// void ~QSemaphore()
func DeleteQSemaphore(*QSemaphore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSemaphoreD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:58
// index:0
// Public
// void acquire(int)
func (this *QSemaphore) Acquire(n int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSemaphore7acquireEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:59
// index:0
// Public
// bool tryAcquire(int)
func (this *QSemaphore) TryAcquire(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsemaphore.h:60
// index:1
// Public
// bool tryAcquire(int, int)
func (this *QSemaphore) TryAcquire_1(n int, timeout int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n, &timeout)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsemaphore.h:62
// index:0
// Public
// void release(int)
func (this *QSemaphore) Release(n int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSemaphore7releaseEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:64
// index:0
// Public
// int available()
func (this *QSemaphore) Available() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSemaphore9availableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
