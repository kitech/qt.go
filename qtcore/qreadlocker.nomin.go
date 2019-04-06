//  header block begin

// +build !minimal

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
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qreadwritelock.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QReadLocker(QReadWriteLock *)

/*

 */
func (*QReadLocker) NewForInherit(readWriteLock QReadWriteLock_ITF /*777 QReadWriteLock **/) *QReadLocker {
	return NewQReadLocker(readWriteLock)
}
func NewQReadLocker(readWriteLock QReadWriteLock_ITF /*777 QReadWriteLock **/) *QReadLocker {
	var convArg0 unsafe.Pointer
	if readWriteLock != nil && readWriteLock.QReadWriteLock_PTR() != nil {
		convArg0 = readWriteLock.QReadWriteLock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLockerC2EP14QReadWriteLock", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQReadLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQReadLocker)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QReadLocker()

/*

 */
func DeleteQReadLocker(this *QReadLocker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLockerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qreadwritelock.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unlock()

/*
Unlocks the lock.

Attempting to unlock a lock that is not locked is an error, and will result in program termination.

See also lockForRead(), lockForWrite(), tryLockForRead(), and tryLockForWrite().
*/
func (this *QReadLocker) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLocker6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void relock()

/*

 */
func (this *QReadLocker) Relock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLocker6relockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QReadWriteLock * readWriteLock() const

/*

 */
func (this *QReadLocker) ReadWriteLock() *QReadWriteLock /*777 QReadWriteLock **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QReadLocker13readWriteLockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10516() {
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
