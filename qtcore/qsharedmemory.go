package qtcore

// /usr/include/qt/QtCore/qsharedmemory.h
// #include <qsharedmemory.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
type QSharedMemory struct {
	*QObject
}
type QSharedMemory_ITF interface {
	QObject_ITF
	QSharedMemory_PTR() *QSharedMemory
}

func (ptr *QSharedMemory) QSharedMemory_PTR() *QSharedMemory { return ptr }

func (this *QSharedMemory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSharedMemory) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQSharedMemoryFromPointer(cthis unsafe.Pointer) *QSharedMemory {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSharedMemory{bcthis0}
}
func (*QSharedMemory) NewFromPointer(cthis unsafe.Pointer) *QSharedMemory {
	return NewQSharedMemoryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsharedmemory.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSharedMemory) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsharedmemory.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSharedMemory(QObject *)

/*
Constructs a shared memory object with the given parent and with its key set to key. Because its key is set, its create() and attach() functions can be called.

See also setKey(), create(), and attach().
*/
func NewQSharedMemory(parent QObject_ITF /*777 QObject **/) *QSharedMemory {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSharedMemory")
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSharedMemory(QObject *)

/*
Constructs a shared memory object with the given parent and with its key set to key. Because its key is set, its create() and attach() functions can be called.

See also setKey(), create(), and attach().
*/
func NewQSharedMemory__() *QSharedMemory {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSharedMemory")
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:78
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSharedMemory(const QString &, QObject *)

/*
Constructs a shared memory object with the given parent and with its key set to key. Because its key is set, its create() and attach() functions can be called.

See also setKey(), create(), and attach().
*/
func NewQSharedMemory_1(key string, parent QObject_ITF /*777 QObject **/) *QSharedMemory {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSharedMemory")
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:78
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSharedMemory(const QString &, QObject *)

/*
Constructs a shared memory object with the given parent and with its key set to key. Because its key is set, its create() and attach() functions can be called.

See also setKey(), create(), and attach().
*/
func NewQSharedMemory_1_(key string) *QSharedMemory {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedMemoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSharedMemory")
	return gothis
}

// /usr/include/qt/QtCore/qsharedmemory.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSharedMemory()

/*

 */
func DeleteQSharedMemory(this *QSharedMemory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsharedmemory.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &)

/*
Sets the platform independent key for this shared memory object. If key is the same as the current key, the function returns without doing anything.

You can call key() to retrieve the platform independent key. Internally, QSharedMemory converts this key into a platform specific key. If you instead call nativeKey(), you will get the platform specific, converted key.

If the shared memory object is attached to an underlying shared memory segment, it will detach from it before setting the new key. This function does not do an attach().

See also key(), nativeKey(), and isAttached().
*/
func (this *QSharedMemory) SetKey(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6setKeyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key() const

/*
Returns the key assigned with setKey() to this shared memory, or a null key if no key has been assigned, or if the segment is using a nativeKey(). The key is the identifier used by Qt applications to identify the shared memory segment.

You can find the native, platform specific, key used by the operating system by calling nativeKey().

See also setKey() and setNativeKey().
*/
func (this *QSharedMemory) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsharedmemory.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNativeKey(const QString &)

/*
Sets the native, platform specific, key for this shared memory object. If key is the same as the current native key, the function returns without doing anything. If all you want is to assign a key to a segment, you should call setKey() instead.

You can call nativeKey() to retrieve the native key. If a native key has been assigned, calling key() will return a null string.

If the shared memory object is attached to an underlying shared memory segment, it will detach from it before setting the new key. This function does not do an attach().

The application will not be portable if you set a native key.

This function was introduced in  Qt 4.8.

See also nativeKey(), key(), and isAttached().
*/
func (this *QSharedMemory) SetNativeKey(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory12setNativeKeyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString nativeKey() const

/*
Returns the native, platform specific, key for this shared memory object. The native key is the identifier used by the operating system to identify the shared memory segment.

You can use the native key to access shared memory segments that have not been created by Qt, or to grant shared memory access to non-Qt applications.

This function was introduced in  Qt 4.8.

See also setKey() and setNativeKey().
*/
func (this *QSharedMemory) NativeKey() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory9nativeKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsharedmemory.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool create(int, enum QSharedMemory::AccessMode)

/*
Creates a shared memory segment of size bytes with the key passed to the constructor, set with setKey() or set with setNativeKey(), then attaches to the new shared memory segment with the given access mode and returns true. If a shared memory segment identified by the key already exists, the attach operation is not performed and false is returned. When the return value is false, call error() to determine which error occurred.

See also error().
*/
func (this *QSharedMemory) Create(size int, mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6createEiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool create(int, enum QSharedMemory::AccessMode)

/*
Creates a shared memory segment of size bytes with the key passed to the constructor, set with setKey() or set with setNativeKey(), then attaches to the new shared memory segment with the given access mode and returns true. If a shared memory segment identified by the key already exists, the attach operation is not performed and false is returned. When the return value is false, call error() to determine which error occurred.

See also error().
*/
func (this *QSharedMemory) Create__(size int) bool {
	// arg: 1, QSharedMemory::AccessMode=Enum, QSharedMemory::AccessMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6createEiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the size of the attached shared memory segment. If no shared memory segment is attached, 0 is returned.

Note: The size of the segment may be larger than the requested size that was passed to create().

See also create() and attach().
*/
func (this *QSharedMemory) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsharedmemory.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool attach(enum QSharedMemory::AccessMode)

/*
Attempts to attach the process to the shared memory segment identified by the key that was passed to the constructor or to a call to setKey() or setNativeKey(). The access mode is ReadWrite by default. It can also be ReadOnly. Returns true if the attach operation is successful. If false is returned, call error() to determine which error occurred. After attaching the shared memory segment, a pointer to the shared memory can be obtained by calling data().

See also isAttached(), detach(), and create().
*/
func (this *QSharedMemory) Attach(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6attachENS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool attach(enum QSharedMemory::AccessMode)

/*
Attempts to attach the process to the shared memory segment identified by the key that was passed to the constructor or to a call to setKey() or setNativeKey(). The access mode is ReadWrite by default. It can also be ReadOnly. Returns true if the attach operation is successful. If false is returned, call error() to determine which error occurred. After attaching the shared memory segment, a pointer to the shared memory can be obtained by calling data().

See also isAttached(), detach(), and create().
*/
func (this *QSharedMemory) Attach__() bool {
	// arg: 0, QSharedMemory::AccessMode=Enum, QSharedMemory::AccessMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6attachENS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAttached() const

/*
Returns true if this process is attached to the shared memory segment.

See also attach() and detach().
*/
func (this *QSharedMemory) IsAttached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory10isAttachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool detach()

/*
Detaches the process from the shared memory segment. If this was the last process attached to the shared memory segment, then the shared memory segment is released by the system, i.e., the contents are destroyed. The function returns true if it detaches the shared memory segment. If it returns false, it usually means the segment either isn't attached, or it is locked by another process.

See also attach() and isAttached().
*/
func (this *QSharedMemory) Detach() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] void * data()

/*
Returns a pointer to the contents of the shared memory segment, if one is attached. Otherwise it returns null. Remember to lock the shared memory with lock() before reading from or writing to the shared memory, and remember to release the lock with unlock() after you are done.

See also attach().
*/
func (this *QSharedMemory) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qsharedmemory.h:95
// index:1
// Public Visibility=Default Availability=Available
// [8] const void * data() const

/*
Returns a pointer to the contents of the shared memory segment, if one is attached. Otherwise it returns null. Remember to lock the shared memory with lock() before reading from or writing to the shared memory, and remember to release the lock with unlock() after you are done.

See also attach().
*/
func (this *QSharedMemory) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qsharedmemory.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] const void * constData() const

/*
Returns a const pointer to the contents of the shared memory segment, if one is attached. Otherwise it returns null. Remember to lock the shared memory with lock() before reading from or writing to the shared memory, and remember to release the lock with unlock() after you are done.

See also attach() and create().
*/
func (this *QSharedMemory) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qsharedmemory.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool lock()

/*
This is a semaphore that locks the shared memory segment for access by this process and returns true. If another process has locked the segment, this function blocks until the lock is released. Then it acquires the lock and returns true. If this function returns false, it means that you have ignored a false return from create() or attach(), that you have set the key with setNativeKey() or that QSystemSemaphore::acquire() failed due to an unknown system error.

See also unlock(), data(), and QSystemSemaphore::acquire().
*/
func (this *QSharedMemory) Lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unlock()

/*
Releases the lock on the shared memory segment and returns true, if the lock is currently held by this process. If the segment is not locked, or if the lock is held by another process, nothing happens and false is returned.

See also lock().
*/
func (this *QSharedMemory) Unlock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSharedMemory6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsharedmemory.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QSharedMemory::SharedMemoryError error() const

/*
Returns a value indicating whether an error occurred, and, if so, which error it was.

See also errorString().
*/
func (this *QSharedMemory) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a text description of the last error that occurred. If error() returns an error value, call this function to get a text string that describes the error.

See also error().
*/
func (this *QSharedMemory) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSharedMemory11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*

 */
type QSharedMemory__AccessMode = int

// The shared memory segment is read-only. Writing to the shared memory segment is not allowed. An attempt to write to a shared memory segment created with ReadOnly causes the program to abort.
const QSharedMemory__ReadOnly QSharedMemory__AccessMode = 0

// Reading and writing the shared memory segment are both allowed.
const QSharedMemory__ReadWrite QSharedMemory__AccessMode = 1

/*

 */
type QSharedMemory__SharedMemoryError = int

// No error occurred.
const QSharedMemory__NoError QSharedMemory__SharedMemoryError = 0

// The operation failed because the caller didn't have the required permissions.
const QSharedMemory__PermissionDenied QSharedMemory__SharedMemoryError = 1

// A create operation failed because the requested size was invalid.
const QSharedMemory__InvalidSize QSharedMemory__SharedMemoryError = 2

// The operation failed because of an invalid key.
const QSharedMemory__KeyError QSharedMemory__SharedMemoryError = 3

// A create() operation failed because a shared memory segment with the specified key already existed.
const QSharedMemory__AlreadyExists QSharedMemory__SharedMemoryError = 4

// An attach() failed because a shared memory segment with the specified key could not be found.
const QSharedMemory__NotFound QSharedMemory__SharedMemoryError = 5

// The attempt to lock() the shared memory segment failed because create() or attach() failed and returned false, or because a system error occurred in QSystemSemaphore::acquire().
const QSharedMemory__LockError QSharedMemory__SharedMemoryError = 6

// A create() operation failed because there was not enough memory available to fill the request.
const QSharedMemory__OutOfResources QSharedMemory__SharedMemoryError = 7

// Something else happened and it was bad.
const QSharedMemory__UnknownError QSharedMemory__SharedMemoryError = 8

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
