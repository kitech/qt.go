package qtcore

// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QBasicMutex struct {
	*qtrt.CObject
}
type QBasicMutex_ITF interface {
	QBasicMutex_PTR() *QBasicMutex
}

func (ptr *QBasicMutex) QBasicMutex_PTR() *QBasicMutex { return ptr }

func (this *QBasicMutex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBasicMutex) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBasicMutexFromPointer(cthis unsafe.Pointer) *QBasicMutex {
	return &QBasicMutex{&qtrt.CObject{cthis}}
}
func (*QBasicMutex) NewFromPointer(cthis unsafe.Pointer) *QBasicMutex {
	return NewQBasicMutexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void lock()

/*
Locks the mutex. If another thread has locked the mutex then this call will block until that thread has unlocked it.

Calling this function multiple times on the same mutex from the same thread is allowed if this mutex is a recursive mutex. If this mutex is a non-recursive mutex, this function will dead-lock when the mutex is locked recursively.

See also unlock().
*/
func (this *QBasicMutex) Lock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unlock()

/*
Unlocks the mutex. Attempting to unlock a mutex in a different thread to the one that locked it results in an error. Unlocking a mutex that is not locked results in undefined behavior.

See also lock().
*/
func (this *QBasicMutex) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool tryLock()

/*
Attempts to lock the mutex. This function returns true if the lock was obtained; otherwise it returns false. If another thread has locked the mutex, this function will wait for at most timeout milliseconds for the mutex to become available.

Note: Passing a negative number as the timeout is equivalent to calling lock(), i.e. this function will wait forever until mutex can be locked if timeout is negative.

If the lock was obtained, the mutex must be unlocked with unlock() before another thread can successfully lock it.

Calling this function multiple times on the same mutex from the same thread is allowed if this mutex is a recursive mutex. If this mutex is a non-recursive mutex, this function will always return false when attempting to lock the mutex recursively.

See also lock() and unlock().
*/
func (this *QBasicMutex) TryLock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex7tryLockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool try_lock()

/*
Attempts to lock the mutex. This function returns true if the lock was obtained; otherwise it returns false.

This function is provided for compatibility with the Standard Library concept Lockable. It is equivalent to tryLock().

The function returns true if the lock was obtained; otherwise it returns false

This function was introduced in  Qt 5.8.
*/
func (this *QBasicMutex) Try_lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex8try_lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRecursive()

/*
Returns true if the mutex is recursive.

This function was introduced in  Qt 5.7.
*/
func (this *QBasicMutex) IsRecursive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:91
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isRecursive() const

/*
Returns true if the mutex is recursive.

This function was introduced in  Qt 5.7.
*/
func (this *QBasicMutex) IsRecursive1() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QBasicMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQBasicMutex(this *QBasicMutex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
