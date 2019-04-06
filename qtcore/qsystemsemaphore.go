package qtcore

// /usr/include/qt/QtCore/qsystemsemaphore.h
// #include <qsystemsemaphore.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QSystemSemaphore struct {
	*qtrt.CObject
}
type QSystemSemaphore_ITF interface {
	QSystemSemaphore_PTR() *QSystemSemaphore
}

func (ptr *QSystemSemaphore) QSystemSemaphore_PTR() *QSystemSemaphore { return ptr }

func (this *QSystemSemaphore) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSystemSemaphore) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSystemSemaphoreFromPointer(cthis unsafe.Pointer) *QSystemSemaphore {
	return &QSystemSemaphore{&qtrt.CObject{cthis}}
}
func (*QSystemSemaphore) NewFromPointer(cthis unsafe.Pointer) *QSystemSemaphore {
	return NewQSystemSemaphoreFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemSemaphore(const QString &, int, QSystemSemaphore::AccessMode)

/*
Requests a system semaphore for the specified key. The parameters initialValue and mode are used according to the following rules, which are system dependent.

In Unix, if the mode is Open and the system already has a semaphore identified by key, that semaphore is used, and the semaphore's resource count is not changed, i.e., initialValue is ignored. But if the system does not already have a semaphore identified by key, it creates a new semaphore for that key and sets its resource count to initialValue.

In Unix, if the mode is Create and the system already has a semaphore identified by key, that semaphore is used, and its resource count is set to initialValue. If the system does not already have a semaphore identified by key, it creates a new semaphore for that key and sets its resource count to initialValue.

In Windows, mode is ignored, and the system always tries to create a semaphore for the specified key. If the system does not already have a semaphore identified as key, it creates the semaphore and sets its resource count to initialValue. But if the system already has a semaphore identified as key it uses that semaphore and ignores initialValue.

The mode parameter is only used in Unix systems to handle the case where a semaphore survives a process crash. In that case, the next process to allocate a semaphore with the same key will get the semaphore that survived the crash, and unless mode is Create, the resource count will not be reset to initialValue but will retain the initial value it had been given by the crashed process.

See also acquire() and key().
*/
func (*QSystemSemaphore) NewForInherit(key string, initialValue int, mode int) *QSystemSemaphore {
	return NewQSystemSemaphore(key, initialValue, mode)
}
func NewQSystemSemaphore(key string, initialValue int, mode int) *QSystemSemaphore {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemSemaphore(const QString &, int, QSystemSemaphore::AccessMode)

/*
Requests a system semaphore for the specified key. The parameters initialValue and mode are used according to the following rules, which are system dependent.

In Unix, if the mode is Open and the system already has a semaphore identified by key, that semaphore is used, and the semaphore's resource count is not changed, i.e., initialValue is ignored. But if the system does not already have a semaphore identified by key, it creates a new semaphore for that key and sets its resource count to initialValue.

In Unix, if the mode is Create and the system already has a semaphore identified by key, that semaphore is used, and its resource count is set to initialValue. If the system does not already have a semaphore identified by key, it creates a new semaphore for that key and sets its resource count to initialValue.

In Windows, mode is ignored, and the system always tries to create a semaphore for the specified key. If the system does not already have a semaphore identified as key, it creates the semaphore and sets its resource count to initialValue. But if the system already has a semaphore identified as key it uses that semaphore and ignores initialValue.

The mode parameter is only used in Unix systems to handle the case where a semaphore survives a process crash. In that case, the next process to allocate a semaphore with the same key will get the semaphore that survived the crash, and unless mode is Create, the resource count will not be reset to initialValue but will retain the initial value it had been given by the crashed process.

See also acquire() and key().
*/
func (*QSystemSemaphore) NewForInheritp(key string) *QSystemSemaphore {
	return NewQSystemSemaphorep(key)
}
func NewQSystemSemaphorep(key string) *QSystemSemaphore {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	initialValue := int(0)
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemSemaphore(const QString &, int, QSystemSemaphore::AccessMode)

/*
Requests a system semaphore for the specified key. The parameters initialValue and mode are used according to the following rules, which are system dependent.

In Unix, if the mode is Open and the system already has a semaphore identified by key, that semaphore is used, and the semaphore's resource count is not changed, i.e., initialValue is ignored. But if the system does not already have a semaphore identified by key, it creates a new semaphore for that key and sets its resource count to initialValue.

In Unix, if the mode is Create and the system already has a semaphore identified by key, that semaphore is used, and its resource count is set to initialValue. If the system does not already have a semaphore identified by key, it creates a new semaphore for that key and sets its resource count to initialValue.

In Windows, mode is ignored, and the system always tries to create a semaphore for the specified key. If the system does not already have a semaphore identified as key, it creates the semaphore and sets its resource count to initialValue. But if the system already has a semaphore identified as key it uses that semaphore and ignores initialValue.

The mode parameter is only used in Unix systems to handle the case where a semaphore survives a process crash. In that case, the next process to allocate a semaphore with the same key will get the semaphore that survived the crash, and unless mode is Create, the resource count will not be reset to initialValue but will retain the initial value it had been given by the crashed process.

See also acquire() and key().
*/
func (*QSystemSemaphore) NewForInheritp1(key string, initialValue int) *QSystemSemaphore {
	return NewQSystemSemaphorep1(key, initialValue)
}
func NewQSystemSemaphorep1(key string, initialValue int) *QSystemSemaphore {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSystemSemaphore()

/*

 */
func DeleteQSystemSemaphore(this *QSystemSemaphore) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, QSystemSemaphore::AccessMode)

/*
This function works the same as the constructor. It reconstructs this QSystemSemaphore object. If the new key is different from the old key, calling this function is like calling the destructor of the semaphore with the old key, then calling the constructor to create a new semaphore with the new key. The initialValue and mode parameters are as defined for the constructor.

See also QSystemSemaphore() and key().
*/
func (this *QSystemSemaphore) SetKey(key string, initialValue int, mode int) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, QSystemSemaphore::AccessMode)

/*
This function works the same as the constructor. It reconstructs this QSystemSemaphore object. If the new key is different from the old key, calling this function is like calling the destructor of the semaphore with the old key, then calling the constructor to create a new semaphore with the new key. The initialValue and mode parameters are as defined for the constructor.

See also QSystemSemaphore() and key().
*/
func (this *QSystemSemaphore) SetKeyp(key string) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	initialValue := int(0)
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, QSystemSemaphore::AccessMode)

/*
This function works the same as the constructor. It reconstructs this QSystemSemaphore object. If the new key is different from the old key, calling this function is like calling the destructor of the semaphore with the old key, then calling the constructor to create a new semaphore with the new key. The initialValue and mode parameters are as defined for the constructor.

See also QSystemSemaphore() and key().
*/
func (this *QSystemSemaphore) SetKeyp1(key string, initialValue int) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key() const

/*
Returns the key assigned to this system semaphore. The key is the name by which the semaphore can be accessed from other processes.

See also setKey().
*/
func (this *QSystemSemaphore) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSystemSemaphore3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acquire()

/*
Acquires one of the resources guarded by this semaphore, if there is one available, and returns true. If all the resources guarded by this semaphore have already been acquired, the call blocks until one of them is released by another process or thread having a semaphore with the same key.

If false is returned, a system error has occurred. Call error() to get a value of QSystemSemaphore::SystemSemaphoreError that indicates which error occurred.

See also release().
*/
func (this *QSystemSemaphore) Acquire() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore7acquireEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool release(int)

/*
Releases n resources guarded by the semaphore. Returns true unless there is a system error.

Example: Create a system semaphore having five resources; acquire them all and then release them all.


  QSystemSemaphore sem("market", 5, QSystemSemaphore::Create);
  for (int i = 0; i < 5; ++i)  // acquire all 5 resources
      sem.acquire();
  sem.release(5);              // release the 5 resources



This function can also "create" resources. For example, immediately following the sequence of statements above, suppose we add the statement:


  sem.release(10);          // "create" 10 new resources



Ten new resources are now guarded by the semaphore, in addition to the five that already existed. You would not normally use this function to create more resources.

See also acquire().
*/
func (this *QSystemSemaphore) Release(n int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool release(int)

/*
Releases n resources guarded by the semaphore. Returns true unless there is a system error.

Example: Create a system semaphore having five resources; acquire them all and then release them all.


  QSystemSemaphore sem("market", 5, QSystemSemaphore::Create);
  for (int i = 0; i < 5; ++i)  // acquire all 5 resources
      sem.acquire();
  sem.release(5);              // release the 5 resources



This function can also "create" resources. For example, immediately following the sequence of statements above, suppose we add the statement:


  sem.release(10);          // "create" 10 new resources



Ten new resources are now guarded by the semaphore, in addition to the five that already existed. You would not normally use this function to create more resources.

See also acquire().
*/
func (this *QSystemSemaphore) Releasep() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QSystemSemaphore::SystemSemaphoreError error() const

/*
Returns a value indicating whether an error occurred, and, if so, which error it was.

See also errorString().
*/
func (this *QSystemSemaphore) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSystemSemaphore5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a text description of the last error that occurred. If error() returns an error value, call this function to get a text string that describes the error.

See also error().
*/
func (this *QSystemSemaphore) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSystemSemaphore11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*
This enum is used by the constructor and setKey(). Its purpose is to enable handling the problem in Unix implementations of semaphores that survive a crash. In Unix, when a semaphore survives a crash, we need a way to force it to reset its resource count, when the system reuses the semaphore. In Windows, where semaphores can't survive a crash, this enum has no effect.


*/
type QSystemSemaphore__AccessMode = int

// If the semaphore already exists, its initial resource count is not reset. If the semaphore does not already exist, it is created and its initial resource count set.
const QSystemSemaphore__Open QSystemSemaphore__AccessMode = 0

// QSystemSemaphore takes ownership of the semaphore and sets its resource count to the requested value, regardless of whether the semaphore already exists by having survived a crash. This value should be passed to the constructor, when the first semaphore for a particular key is constructed and you know that if the semaphore already exists it could only be because of a crash. In Windows, where a semaphore can't survive a crash, Create and Open have the same behavior.
const QSystemSemaphore__Create QSystemSemaphore__AccessMode = 1

func (this *QSystemSemaphore) AccessModeItemName(val int) string {
	switch val {
	case QSystemSemaphore__Open: // 0
		return "Open"
	case QSystemSemaphore__Create: // 1
		return "Create"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSystemSemaphore_AccessModeItemName(val int) string {
	var nilthis *QSystemSemaphore
	return nilthis.AccessModeItemName(val)
}

/*

 */
type QSystemSemaphore__SystemSemaphoreError = int

// No error occurred.
const QSystemSemaphore__NoError QSystemSemaphore__SystemSemaphoreError = 0

// The operation failed because the caller didn't have the required permissions.
const QSystemSemaphore__PermissionDenied QSystemSemaphore__SystemSemaphoreError = 1

// The operation failed because of an invalid key.
const QSystemSemaphore__KeyError QSystemSemaphore__SystemSemaphoreError = 2

// The operation failed because a system semaphore with the specified key already existed.
const QSystemSemaphore__AlreadyExists QSystemSemaphore__SystemSemaphoreError = 3

// The operation failed because a system semaphore with the specified key could not be found.
const QSystemSemaphore__NotFound QSystemSemaphore__SystemSemaphoreError = 4

// The operation failed because there was not enough memory available to fill the request.
const QSystemSemaphore__OutOfResources QSystemSemaphore__SystemSemaphoreError = 5

// Something else happened and it was bad.
const QSystemSemaphore__UnknownError QSystemSemaphore__SystemSemaphoreError = 6

func (this *QSystemSemaphore) SystemSemaphoreErrorItemName(val int) string {
	switch val {
	case QSystemSemaphore__NoError: // 0
		return "NoError"
	case QSystemSemaphore__PermissionDenied: // 1
		return "PermissionDenied"
	case QSystemSemaphore__KeyError: // 2
		return "KeyError"
	case QSystemSemaphore__AlreadyExists: // 3
		return "AlreadyExists"
	case QSystemSemaphore__NotFound: // 4
		return "NotFound"
	case QSystemSemaphore__OutOfResources: // 5
		return "OutOfResources"
	case QSystemSemaphore__UnknownError: // 6
		return "UnknownError"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSystemSemaphore_SystemSemaphoreErrorItemName(val int) string {
	var nilthis *QSystemSemaphore
	return nilthis.SystemSemaphoreErrorItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10561() {
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
