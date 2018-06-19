package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QReadWriteLock struct {
	*qtrt.CObject
}
type QReadWriteLock_ITF interface {
	QReadWriteLock_PTR() *QReadWriteLock
}

func (ptr *QReadWriteLock) QReadWriteLock_PTR() *QReadWriteLock { return ptr }

func (this *QReadWriteLock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QReadWriteLock) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQReadWriteLockFromPointer(cthis unsafe.Pointer) *QReadWriteLock {
	return &QReadWriteLock{&qtrt.CObject{cthis}}
}
func (*QReadWriteLock) NewFromPointer(cthis unsafe.Pointer) *QReadWriteLock {
	return NewQReadWriteLockFromPointer(cthis)
}

// /usr/include/qt/QtCore/qreadwritelock.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QReadWriteLock(QReadWriteLock::RecursionMode)

/*
Constructs a QReadWriteLock object in the given recursionMode.

The default recursion mode is NonRecursive.

This function was introduced in  Qt 4.4.

See also lockForRead(), lockForWrite(), and RecursionMode.
*/
func NewQReadWriteLock(recursionMode int) *QReadWriteLock {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLockC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, recursionMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQReadWriteLock)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QReadWriteLock(QReadWriteLock::RecursionMode)

/*
Constructs a QReadWriteLock object in the given recursionMode.

The default recursion mode is NonRecursive.

This function was introduced in  Qt 4.4.

See also lockForRead(), lockForWrite(), and RecursionMode.
*/
func NewQReadWriteLock__() *QReadWriteLock {
	// arg: 0, QReadWriteLock::RecursionMode=Enum, QReadWriteLock::RecursionMode=Enum, , Invalid
	recursionMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLockC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, recursionMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQReadWriteLock)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QReadWriteLock()

/*

 */
func DeleteQReadWriteLock(this *QReadWriteLock) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLockD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qreadwritelock.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockForRead()

/*
Locks the lock for reading. This function will block the current thread if another thread has locked for writing.

It is not possible to lock for read if the thread already has locked for write.

See also unlock(), lockForWrite(), and tryLockForRead().
*/
func (this *QReadWriteLock) LockForRead() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock11lockForReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLockForRead()

/*
Attempts to lock for reading. If the lock was obtained, this function returns true, otherwise it returns false instead of waiting for the lock to become available, i.e. it does not block.

The lock attempt will fail if another thread has locked for writing.

If the lock was obtained, the lock must be unlocked with unlock() before another thread can successfully lock it for writing.

It is not possible to lock for read if the thread already has locked for write.

See also unlock() and lockForRead().
*/
func (this *QReadWriteLock) TryLockForRead() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock14tryLockForReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:62
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryLockForRead(int)

/*
Attempts to lock for reading. If the lock was obtained, this function returns true, otherwise it returns false instead of waiting for the lock to become available, i.e. it does not block.

The lock attempt will fail if another thread has locked for writing.

If the lock was obtained, the lock must be unlocked with unlock() before another thread can successfully lock it for writing.

It is not possible to lock for read if the thread already has locked for write.

See also unlock() and lockForRead().
*/
func (this *QReadWriteLock) TryLockForRead_1(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock14tryLockForReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockForWrite()

/*
Locks the lock for writing. This function will block the current thread if another thread (including the current) has locked for reading or writing (unless the lock has been created using the QReadWriteLock::Recursive mode).

It is not possible to lock for write if the thread already has locked for read.

See also unlock(), lockForRead(), and tryLockForWrite().
*/
func (this *QReadWriteLock) LockForWrite() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock12lockForWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLockForWrite()

/*
Attempts to lock for writing. If the lock was obtained, this function returns true; otherwise, it returns false immediately.

The lock attempt will fail if another thread has locked for reading or writing.

If the lock was obtained, the lock must be unlocked with unlock() before another thread can successfully lock it.

It is not possible to lock for write if the thread already has locked for read.

See also unlock() and lockForWrite().
*/
func (this *QReadWriteLock) TryLockForWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock15tryLockForWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:66
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryLockForWrite(int)

/*
Attempts to lock for writing. If the lock was obtained, this function returns true; otherwise, it returns false immediately.

The lock attempt will fail if another thread has locked for reading or writing.

If the lock was obtained, the lock must be unlocked with unlock() before another thread can successfully lock it.

It is not possible to lock for write if the thread already has locked for read.

See also unlock() and lockForWrite().
*/
func (this *QReadWriteLock) TryLockForWrite_1(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock15tryLockForWriteEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()

/*
Unlocks the lock.

Attempting to unlock a lock that is not locked is an error, and will result in program termination.

See also lockForRead(), lockForWrite(), tryLockForRead(), and tryLockForWrite().
*/
func (this *QReadWriteLock) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


This enum was introduced or modified in  Qt 4.4.

See also QReadWriteLock().

*/
type QReadWriteLock__RecursionMode = int

// In this mode, a thread may only lock a QReadWriteLock once.
const QReadWriteLock__NonRecursive QReadWriteLock__RecursionMode = 0

// In this mode, a thread can lock the same QReadWriteLock multiple times. The QReadWriteLock won't be unlocked until a corresponding number of unlock() calls have been made.
const QReadWriteLock__Recursive QReadWriteLock__RecursionMode = 1

/*


 */
type QReadWriteLock__StateForWaitCondition = int

//
const QReadWriteLock__LockedForRead QReadWriteLock__StateForWaitCondition = 0

//
const QReadWriteLock__LockedForWrite QReadWriteLock__StateForWaitCondition = 1

//
const QReadWriteLock__Unlocked QReadWriteLock__StateForWaitCondition = 2

//
const QReadWriteLock__RecursivelyLocked QReadWriteLock__StateForWaitCondition = 3

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
