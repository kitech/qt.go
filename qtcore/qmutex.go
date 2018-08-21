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
type QMutex struct {
	*QBasicMutex
}
type QMutex_ITF interface {
	QBasicMutex_ITF
	QMutex_PTR() *QMutex
}

func (ptr *QMutex) QMutex_PTR() *QMutex { return ptr }

func (this *QMutex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QBasicMutex.GetCthis()
	}
}
func (this *QMutex) SetCthis(cthis unsafe.Pointer) {
	this.QBasicMutex = NewQBasicMutexFromPointer(cthis)
}
func NewQMutexFromPointer(cthis unsafe.Pointer) *QMutex {
	bcthis0 := NewQBasicMutexFromPointer(cthis)
	return &QMutex{bcthis0}
}
func (*QMutex) NewFromPointer(cthis unsafe.Pointer) *QMutex {
	return NewQMutexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMutex(QMutex::RecursionMode)

/*
Constructs a new mutex. The mutex is created in an unlocked state.

If mode is QMutex::Recursive, a thread can lock the same mutex multiple times and the mutex won't be unlocked until a corresponding number of unlock() calls have been made. Otherwise a thread may only lock a mutex once. The default is QMutex::NonRecursive.

Recursive mutexes are slower and take more memory than non-recursive ones.

See also lock() and unlock().
*/
func NewQMutex(mode int) *QMutex {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMutexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMutex)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMutex(QMutex::RecursionMode)

/*
Constructs a new mutex. The mutex is created in an unlocked state.

If mode is QMutex::Recursive, a thread can lock the same mutex multiple times and the mutex won't be unlocked until a corresponding number of unlock() calls have been made. Otherwise a thread may only lock a mutex once. The default is QMutex::NonRecursive.

Recursive mutexes are slower and take more memory than non-recursive ones.

See also lock() and unlock().
*/
func NewQMutex__() *QMutex {
	// arg: 0, QMutex::RecursionMode=Enum, QMutex::RecursionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMutexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMutex)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMutex()

/*

 */
func DeleteQMutex(this *QMutex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmutex.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lock()

/*
Locks the mutex. If another thread has locked the mutex then this call will block until that thread has unlocked it.

Calling this function multiple times on the same mutex from the same thread is allowed if this mutex is a recursive mutex. If this mutex is a non-recursive mutex, this function will dead-lock when the mutex is locked recursively.

See also unlock().
*/
func (this *QMutex) Lock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)

/*
Attempts to lock the mutex. This function returns true if the lock was obtained; otherwise it returns false. If another thread has locked the mutex, this function will wait for at most timeout milliseconds for the mutex to become available.

Note: Passing a negative number as the timeout is equivalent to calling lock(), i.e. this function will wait forever until mutex can be locked if timeout is negative.

If the lock was obtained, the mutex must be unlocked with unlock() before another thread can successfully lock it.

Calling this function multiple times on the same mutex from the same thread is allowed if this mutex is a recursive mutex. If this mutex is a non-recursive mutex, this function will always return false when attempting to lock the mutex recursively.

See also lock() and unlock().
*/
func (this *QMutex) TryLock(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)

/*
Attempts to lock the mutex. This function returns true if the lock was obtained; otherwise it returns false. If another thread has locked the mutex, this function will wait for at most timeout milliseconds for the mutex to become available.

Note: Passing a negative number as the timeout is equivalent to calling lock(), i.e. this function will wait forever until mutex can be locked if timeout is negative.

If the lock was obtained, the mutex must be unlocked with unlock() before another thread can successfully lock it.

Calling this function multiple times on the same mutex from the same thread is allowed if this mutex is a recursive mutex. If this mutex is a non-recursive mutex, this function will always return false when attempting to lock the mutex recursively.

See also lock() and unlock().
*/
func (this *QMutex) TryLock__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()

/*
Unlocks the mutex. Attempting to unlock a mutex in a different thread to the one that locked it results in an error. Unlocking a mutex that is not locked results in undefined behavior.

See also lock().
*/
func (this *QMutex) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool try_lock()

/*
Attempts to lock the mutex. This function returns true if the lock was obtained; otherwise it returns false.

This function is provided for compatibility with the Standard Library concept Lockable. It is equivalent to tryLock().

The function returns true if the lock was obtained; otherwise it returns false

This function was introduced in  Qt 5.8.
*/
func (this *QMutex) Try_lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex8try_lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRecursive() const

/*
Returns true if the mutex is recursive.

This function was introduced in  Qt 5.7.
*/
func (this *QMutex) IsRecursive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


See also QMutex().

*/
type QMutex__RecursionMode = int

// In this mode, a thread may only lock a mutex once.
const QMutex__NonRecursive QMutex__RecursionMode = 0

// In this mode, a thread can lock the same mutex multiple times and the mutex won't be unlocked until a corresponding number of unlock() calls have been made.
const QMutex__Recursive QMutex__RecursionMode = 1

func (this *QMutex) RecursionModeItemName(val int) string {
	switch val {
	case QMutex__NonRecursive: // 0
		return "NonRecursive"
	case QMutex__Recursive: // 1
		return "Recursive"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMutex_RecursionModeItemName(val int) string {
	var nilthis *QMutex
	return nilthis.RecursionModeItemName(val)
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
