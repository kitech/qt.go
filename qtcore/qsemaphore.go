package qtcore

// /usr/include/qt/QtCore/qsemaphore.h
// #include <qsemaphore.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QSemaphore struct {
	*qtrt.CObject
}
type QSemaphore_ITF interface {
	QSemaphore_PTR() *QSemaphore
}

func (ptr *QSemaphore) QSemaphore_PTR() *QSemaphore { return ptr }

func (this *QSemaphore) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSemaphore) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSemaphoreFromPointer(cthis unsafe.Pointer) *QSemaphore {
	return &QSemaphore{&qtrt.CObject{cthis}}
}
func (*QSemaphore) NewFromPointer(cthis unsafe.Pointer) *QSemaphore {
	return NewQSemaphoreFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsemaphore.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSemaphore(int)
func NewQSemaphore(n int) *QSemaphore {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphoreC2Ei", qtrt.FFI_TYPE_POINTER, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSemaphore()
func DeleteQSemaphore(this *QSemaphore) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphoreD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsemaphore.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acquire(int)
func (this *QSemaphore) Acquire(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore7acquireEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryAcquire(int)
func (this *QSemaphore) TryAcquire(n int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsemaphore.h:60
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryAcquire(int, int)
func (this *QSemaphore) TryAcquire_1(n int, timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n, timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsemaphore.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void release(int)
func (this *QSemaphore) Release(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int available()
func (this *QSemaphore) Available() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSemaphore9availableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
