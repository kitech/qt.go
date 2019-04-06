// +build !minimal

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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
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

// /usr/include/qt/QtCore/qsemaphore.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSemaphore(int)

/*
Creates a new semaphore and initializes the number of resources it guards to n (by default, 0).

See also release() and available().
*/
func (*QSemaphore) NewForInherit(n int) *QSemaphore {
	return NewQSemaphore(n)
}
func NewQSemaphore(n int) *QSemaphore {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphoreC2Ei", qtrt.FFI_TYPE_POINTER, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSemaphore(int)

/*
Creates a new semaphore and initializes the number of resources it guards to n (by default, 0).

See also release() and available().
*/
func (*QSemaphore) NewForInheritp() *QSemaphore {
	return NewQSemaphorep()
}
func NewQSemaphorep() *QSemaphore {
	// arg: 0, int=Int, =Invalid, , Invalid
	n := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphoreC2Ei", qtrt.FFI_TYPE_POINTER, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSemaphore()

/*

 */
func DeleteQSemaphore(this *QSemaphore) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphoreD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsemaphore.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acquire(int)

/*
Tries to acquire n resources guarded by the semaphore. If n > available(), this call will block until enough resources are available.

See also release(), available(), and tryAcquire().
*/
func (this *QSemaphore) Acquire(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore7acquireEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acquire(int)

/*
Tries to acquire n resources guarded by the semaphore. If n > available(), this call will block until enough resources are available.

See also release(), available(), and tryAcquire().
*/
func (this *QSemaphore) Acquirep() {
	// arg: 0, int=Int, =Invalid, , Invalid
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore7acquireEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:58
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryAcquire(int)

/*
Tries to acquire n resources guarded by the semaphore and returns true on success. If available() < n, this call immediately returns false without acquiring any resources.

Example:


  QSemaphore sem(5);      // sem.available() == 5
  sem.tryAcquire(250);    // sem.available() == 5, returns false
  sem.tryAcquire(3);      // sem.available() == 2, returns true



See also acquire().
*/
func (this *QSemaphore) TryAcquire(n int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsemaphore.h:58
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryAcquire(int)

/*
Tries to acquire n resources guarded by the semaphore and returns true on success. If available() < n, this call immediately returns false without acquiring any resources.

Example:


  QSemaphore sem(5);      // sem.available() == 5
  sem.tryAcquire(250);    // sem.available() == 5, returns false
  sem.tryAcquire(3);      // sem.available() == 2, returns true



See also acquire().
*/
func (this *QSemaphore) TryAcquirep() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsemaphore.h:59
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryAcquire(int, int)

/*
Tries to acquire n resources guarded by the semaphore and returns true on success. If available() < n, this call immediately returns false without acquiring any resources.

Example:


  QSemaphore sem(5);      // sem.available() == 5
  sem.tryAcquire(250);    // sem.available() == 5, returns false
  sem.tryAcquire(3);      // sem.available() == 2, returns true



See also acquire().
*/
func (this *QSemaphore) TryAcquire1(n int, timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore10tryAcquireEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n, timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsemaphore.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void release(int)

/*
Releases n resources guarded by the semaphore.

This function can be used to "create" resources as well. For example:


  QSemaphore sem(5);      // a semaphore that guards 5 resources
  sem.acquire(5);         // acquire all 5 resources
  sem.release(5);         // release the 5 resources
  sem.release(10);        // "create" 10 new resources



QSemaphoreReleaser is a RAII wrapper around this function.

See also acquire(), available(), and QSemaphoreReleaser.
*/
func (this *QSemaphore) Release(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void release(int)

/*
Releases n resources guarded by the semaphore.

This function can be used to "create" resources as well. For example:


  QSemaphore sem(5);      // a semaphore that guards 5 resources
  sem.acquire(5);         // acquire all 5 resources
  sem.release(5);         // release the 5 resources
  sem.release(10);        // "create" 10 new resources



QSemaphoreReleaser is a RAII wrapper around this function.

See also acquire(), available(), and QSemaphoreReleaser.
*/
func (this *QSemaphore) Releasep() {
	// arg: 0, int=Int, =Invalid, , Invalid
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int available() const

/*
Returns the number of resources currently available to the semaphore. This number can never be negative.

See also acquire() and release().
*/
func (this *QSemaphore) Available() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSemaphore9availableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

func init_unused_10531() {
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
