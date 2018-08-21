package qtcore

// /usr/include/qt/QtCore/qlockfile.h
// #include <qlockfile.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QLockFile struct {
	*qtrt.CObject
}
type QLockFile_ITF interface {
	QLockFile_PTR() *QLockFile
}

func (ptr *QLockFile) QLockFile_PTR() *QLockFile { return ptr }

func (this *QLockFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLockFile) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLockFileFromPointer(cthis unsafe.Pointer) *QLockFile {
	return &QLockFile{&qtrt.CObject{cthis}}
}
func (*QLockFile) NewFromPointer(cthis unsafe.Pointer) *QLockFile {
	return NewQLockFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlockfile.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLockFile(const QString &)

/*
Constructs a new lock file object. The object is created in an unlocked state. When calling lock() or tryLock(), a lock file named fileName will be created, if it doesn't already exist.

See also lock() and unlock().
*/
func NewQLockFile(fileName string) *QLockFile {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLockFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLockFile)
	return gothis
}

// /usr/include/qt/QtCore/qlockfile.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QLockFile()

/*

 */
func DeleteQLockFile(this *QLockFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qlockfile.h:56
// index:0
// Public Visibility=Default Availability=Available
// [1] bool lock()

/*
Creates the lock file.

If another process (or another thread) has created the lock file already, this function will block until that process (or thread) releases it.

Calling this function multiple times on the same lock from the same thread without unlocking first is not allowed. This function will dead-lock when the file is locked recursively.

Returns true if the lock was acquired, false if it could not be acquired due to an unrecoverable error, such as no permissions in the parent directory.

See also unlock() and tryLock().
*/
func (this *QLockFile) Lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:57
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)

/*
Attempts to create the lock file. This function returns true if the lock was obtained; otherwise it returns false. If another process (or another thread) has created the lock file already, this function will wait for at most timeout milliseconds for the lock file to become available.

Note: Passing a negative number as the timeout is equivalent to calling lock(), i.e. this function will wait forever until the lock file can be locked if timeout is negative.

If the lock was obtained, it must be released with unlock() before another process (or thread) can successfully lock it.

Calling this function multiple times on the same lock from the same thread without unlocking first is not allowed, this function will always return false when attempting to lock the file recursively.

See also lock() and unlock().
*/
func (this *QLockFile) TryLock(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:57
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)

/*
Attempts to create the lock file. This function returns true if the lock was obtained; otherwise it returns false. If another process (or another thread) has created the lock file already, this function will wait for at most timeout milliseconds for the lock file to become available.

Note: Passing a negative number as the timeout is equivalent to calling lock(), i.e. this function will wait forever until the lock file can be locked if timeout is negative.

If the lock was obtained, it must be released with unlock() before another process (or thread) can successfully lock it.

Calling this function multiple times on the same lock from the same thread without unlocking first is not allowed, this function will always return false when attempting to lock the file recursively.

See also lock() and unlock().
*/
func (this *QLockFile) TryLock__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()

/*
Releases the lock, by deleting the lock file.

Calling unlock() without locking the file first, does nothing.

See also lock() and tryLock().
*/
func (this *QLockFile) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStaleLockTime(int)

/*
Sets staleLockTime to be the time in milliseconds after which a lock file is considered stale. The default value is 30000, i.e. 30 seconds. If your application typically keeps the file locked for more than 30 seconds (for instance while saving megabytes of data for 2 minutes), you should set a bigger value using setStaleLockTime().

The value of staleLockTime is used by lock() and tryLock() in order to determine when an existing lock file is considered stale, i.e. left over by a crashed process. This is useful for the case where the PID got reused meanwhile, so one way to detect a stale lock file is by the fact that it has been around for a long time.

See also staleLockTime().
*/
func (this *QLockFile) SetStaleLockTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile16setStaleLockTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:61
// index:0
// Public Visibility=Default Availability=Available
// [4] int staleLockTime() const

/*
Returns the time in milliseconds after which a lock file is considered stale.

See also setStaleLockTime().
*/
func (this *QLockFile) StaleLockTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile13staleLockTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlockfile.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLocked() const

/*
Returns true if the lock was acquired by this QLockFile instance, otherwise returns false.

See also lock(), unlock(), and tryLock().
*/
func (this *QLockFile) IsLocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile8isLockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool getLockInfo(qint64 *, QString *, QString *) const

/*
Retrieves information about the current owner of the lock file.

If tryLock() returns false, and error() returns LockFailedError, this function can be called to find out more information about the existing lock file:


the PID of the application (returned in pid)
the hostname it's running on (useful in case of networked filesystems),
the name of the application which created it (returned in appname),


Note that tryLock() automatically deleted the file if there is no running application with this PID, so LockFailedError can only happen if there is an application with this PID (it could be unrelated though).

This can be used to inform users about the existing lock file and give them the choice to delete it. After removing the file using removeStaleLockFile(), the application can call tryLock() again.

This function returns true if the information could be successfully retrieved, false if the lock file doesn't exist or doesn't contain the expected data. This can happen if the lock file was deleted between the time where tryLock() failed and the call to this function. Simply call tryLock() again if this happens.
*/
func (this *QLockFile) GetLockInfo(pid unsafe.Pointer /*666*/, hostname string, appname string) bool {
	var tmpArg1 = NewQString_5(hostname)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(appname)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile11getLockInfoEPxP7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pid, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removeStaleLockFile()

/*
Attempts to forcefully remove an existing lock file.

Calling this is not recommended when protecting a short-lived operation: QLockFile already takes care of removing lock files after they are older than staleLockTime().

This method should only be called when protecting a resource for a long time, i.e. with staleLockTime(0), and after tryLock() returned LockFailedError, and the user agreed on removing the lock file.

Returns true on success, false if the lock file couldn't be removed. This happens on Windows, when the application owning the lock is still running.
*/
func (this *QLockFile) RemoveStaleLockFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile19removeStaleLockFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] QLockFile::LockError error() const

/*
Returns the lock file error status.

If tryLock() returns false, this function can be called to find out the reason why the locking failed.
*/
func (this *QLockFile) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*
This enum describes the result of the last call to lock() or tryLock().


*/
type QLockFile__LockError = int

// The lock was acquired successfully.
const QLockFile__NoError QLockFile__LockError = 0

// The lock could not be acquired because another process holds it.
const QLockFile__LockFailedError QLockFile__LockError = 1

// The lock file could not be created, for lack of permissions in the parent directory.
const QLockFile__PermissionError QLockFile__LockError = 2

// Another error happened, for instance a full partition prevented writing out the lock file.
const QLockFile__UnknownError QLockFile__LockError = 3

func (this *QLockFile) LockErrorItemName(val int) string {
	switch val {
	case QLockFile__NoError: // 0
		return "NoError"
	case QLockFile__LockFailedError: // 1
		return "LockFailedError"
	case QLockFile__PermissionError: // 2
		return "PermissionError"
	case QLockFile__UnknownError: // 3
		return "UnknownError"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLockFile_LockErrorItemName(val int) string {
	var nilthis *QLockFile
	return nilthis.LockErrorItemName(val)
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
