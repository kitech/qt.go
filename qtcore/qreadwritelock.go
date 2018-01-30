package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
type QReadWriteLock struct {
	*qtrt.CObject
}

func (this *QReadWriteLock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QReadWriteLock) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
// [-2] void QReadWriteLock(enum QReadWriteLock::RecursionMode)
func NewQReadWriteLock(recursionMode int) *QReadWriteLock {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLockC2ENS_13RecursionModeE", ffiqt.FFI_TYPE_POINTER, recursionMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QReadWriteLock()
func DeleteQReadWriteLock(*QReadWriteLock) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLockD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockForRead()
func (this *QReadWriteLock) LockForRead() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock11lockForReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLockForRead()
func (this *QReadWriteLock) TryLockForRead() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock14tryLockForReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:62
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryLockForRead(int)
func (this *QReadWriteLock) TryLockForRead_1(timeout int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock14tryLockForReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockForWrite()
func (this *QReadWriteLock) LockForWrite() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock12lockForWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLockForWrite()
func (this *QReadWriteLock) TryLockForWrite() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock15tryLockForWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:66
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryLockForWrite(int)
func (this *QReadWriteLock) TryLockForWrite_1(timeout int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock15tryLockForWriteEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QReadWriteLock) Unlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QReadWriteLock6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QReadWriteLock__RecursionMode = int

const QReadWriteLock__NonRecursive QReadWriteLock__RecursionMode = 0
const QReadWriteLock__Recursive QReadWriteLock__RecursionMode = 1

type QReadWriteLock__StateForWaitCondition = int

const QReadWriteLock__LockedForRead QReadWriteLock__StateForWaitCondition = 0
const QReadWriteLock__LockedForWrite QReadWriteLock__StateForWaitCondition = 1
const QReadWriteLock__Unlocked QReadWriteLock__StateForWaitCondition = 2
const QReadWriteLock__RecursivelyLocked QReadWriteLock__StateForWaitCondition = 3

//  body block end
